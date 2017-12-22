# Webhooks Beta


## PUBLIC
Authors: @erictune, @caesarxuchao, @enisoc
Thanks to: {@dbsmith,  @smarterclayton, @deads2k, @cheftako, @jpbetz, @mbohlool, @mml, @janetkuo} for comments, data, prior designs, etc.  


[TOC]


# Summary 

This document proposes a detailed plan for bringing Webhooks to Beta. Highlights include (incomplete, see rest of doc for complete list) :



*   Adding the ability for webhooks to mutate.  
*   Bootstrapping
*   Monitoring
*   Versioned rather than Internal data sent on hook
*   Ordering behavior within webhooks, and with other admission phases, is better defined

This plan is compatible with the [original design doc](/contributors/design-proposals/api-machinery/admission_control_extension.md).


# Definitions 

**Mutating Webhook**: Webhook that can change a request as well as accept/reject.  

**Non-Mutating Webhook**: Webhook that cannot change request, but can accept or reject.

**Webhook**: encompasses both Mutating Webhook and/or Non-mutating Webhook.

**Validating Webhook**: synonym for Non-Mutating Webhook

**Static Admission Controller**: Compiled-in Admission Controllers, (in plugin/pkg/admission).

**Webhook Host**: a process / binary hosting a webhook.

# Naming 

Many names were considered before settling on mutating. None of the names
considered were completely satisfactory. The following are the names which were
considered and a brief explanation of the perspectives on each.

* Mutating: Well defined meaning related to mutable and immutable. Some
  negative connotations related to genetic mutation. Might be too specifically
  as CS term.
* Defaulting: Clearly indicates a create use case. However implies a lack of
  functionality for the update use case.
* Modifying: Similar issues to mutating but not as well defined.
* Revising: Less clear what it does. Does it imply it works only on updates?
* Transforming: Some concern that it might have more to do with changing the
  type or shape of the related object.
* Adjusting: Same general perspective as modifying.
* Updating: Nice clear meaning. However it seems to easy to confuse update with
  the update operation and intuit it does not apply to the create operation.  


# Development Plan 

Google able to staff development, test, review, and documentation.  Community help welcome, too, esp. Reviewing.

Intent is Beta of Webhooks (**both** kinds) in 1.9.  

Not in scope:



*   Initializers remains Alpha for 1.9.  (See [Comparison of Webhooks and Initializers](#comparison-of-webhooks-and-initializers) section).  No changes to it.  Will revisit its status post-1.9.
*   Converting static admission controllers is out of scope (but some investigation has been done, see Moving Built-in Admission Controllers section).


## Work Items 

*   Add API for registering mutating webhooks.  See [API Changes](#api-changes)
*   Copy the non-mutating webhook admission controller code and rename it to be for mutating. (Splitting into two registration APIs make ordering clear.) Add changes to handle mutating responses.  See [Responses for Mutations](#responses-for-mutations).
*   Document recommended flag order for admission plugins.  See [Order of Admission](#order-of-admission).
*   In kube-up.sh and other installers, change flag per previous item. 
*   Ensure able to monitor latency and rejection from webhooks. See [Monitorability](#monitorability).
*   Don't send internal objects.   See [#49733](https://github.com/kubernetes/kubernetes/issues/49733)
*   Serialize mutating Webhooks into order in the apiregistration. Leave non-mutating in parallel.
*   Good Error Messages.  See [Good Error Messages](#good-error-messages)	
*   Conversion logic in GenericWebhook to send converted resource to webhook.  See [Conversion](#conversion) and [#49733](https://github.com/kubernetes/kubernetes/issues/49733).
*   Schedule discussion around resiliency to down webhooks and bootstrapping
*   Internal Go interface refactor (e.g. along the lines suggested  #[1137](https://github.com/kubernetes/community/pull/1137)).


# Design Discussion 


## Why Webhooks First  

We will do webhooks beta before initializers beta because:



1.  **Serves Most Use Cases**: We reviewed code of all current use cases, namely: Kubernetes Built-in Admission Controllers, OpenShift Admission Controllers, Istio & Service Catalog.  (See also [Use Cases Detailed Descriptions](#use-cases-detailed-descriptions).)  All of those use cases are well served by mutating and non-mutating webhooks.  (See also [Comparison of Webhooks and Initializers](#comparison-of-webhooks-and-initializers)).
1.  **Less Work**: An engineer quite experienced with both code bases estimated that it is less work to adding Mutating Webhooks and bring both kinds of webhooks to beta; than to bring non-mutating webhooks and initializers to Beta.  Some open issues with Initializers with long expected development time include quota replenishment bug, and controller awareness of uninitialized objects.
1.  **API Consistency**: Prefer completing one related pair of interfaces (both kinds of webhooks) at the same time.


## Why Support Mutation for Beta 

Based on experience and feedback from the alpha phase of both Webhooks and Initializers, we believe Webhooks Beta should support mutation because:  



1.  We have lots of use cases to inform this (both from Initializers, and Admission Controllers) to ensure we have needed features
1.  We have experience with Webhooks API already to give confidence in the API.  The registration API will be quite similar except in the responses.
1.  There is a strong community demand for something that satisfies a mutating case.


## Plan for Existing Initializer-Clients 

After the release of 1.9, we will advise users who currently use initializers to:



*   Move to Webhooks if their use case fits that model well.
*   Provide SIG-API-Machinery with feedback if Initializers is a better fit. 

We will continue to support Initializers as an Alpha API in 1.9.

We will make a user guide and extensively document these webhooks. We will update some existing examples, maybe https://github.com/caesarxuchao/example-webhook-admission-controller (since the initializer docs point to it, e.g. https://github.com/kelseyhightower/kubernetes-initializer-tutorial), or maybe https://github.com/openshift/generic-admission-server.

We will clearly document the reasons for each and how users should decide which to use.


## Monitorability 

There should be prometheus variables to show:



*   API operation latency
    *   Overall
    *   By webhook name
*   API response codes
    *   Overall
    *   By webhook name.

Adding a webhook dynamically adds a key to a map-valued prometheus metric. Webhook host process authors should consider how to make their webhook host monitorable: while eventually we hope to offer a set of best practices around this, for the initial release we won't have requirements here.


## API Changes 

GenericAdmissionWebhook Admission Controller is split and renamed.



*   One is called `MutatingAdmissionWebhook`
*   The other is called `ValidatingAdmissionWebhook`
*   Splitting them allows them to appear in different places in the `--admission-control` flag's order.

ExternalAdmissionHookConfiguration API is split and renamed.



*   One is called  `MutatingAdmissionWebhookConfiguration`
*   The other is called `ValidatingAdmissionWebhookConfiguration`
*   Splitting them:
    *   makes it clear what the order is when some items don't have both flavors,
    *   enforces mutate-before-validate,
    *   better allows declarative update of the config than one big list with an implied partition point

The `ValidatingAdmissionWebhookConfiguration` stays the same as `ExternalAdmissionHookConfiguration` except it moves to v1beta1.

The `MutatingAdmissionWebhookConfiguration` is the same API as `ValidatingAdmissionWebhookConfiguration`.  It is only visible via the v1beta1 version.

We will change from having a Kubernetes service object to just accepting a DNS
name for the location of the webhook.

The Group/Version called 

`admissionregistration.k8s.io/v1alpha1` with kinds

InitializerConfiguration and ExternalAdmissionHookConfiguration.

InitializerConfiguration will not join `admissionregistration.k8s.io/v1beta1` at this time.

Any webhooks that register with v1alpha1 may or may not be surprised when they start getting versioned data.  But we don't make any promises for Alpha, and this is a very important bug to fix.


## Order of Admission 

At kubernetes.io, we will document the ordering requirements or just recommend a particular order for `--admission-control`. A starting point might be `MutatingAdmissionWebhook,NamespaceLifecycle,LimitRanger,ServiceAccount,PersistentVolumeLabel,DefaultStorageClass,DefaultTolerationSeconds,ValidatingAdmissionWebhook,ResourceQuota`.

 There might be other ordering dependencies that we will document clearly, but some important properties of a valid ordering:

*   ResourceQuota comes last, so that if prior ones reject a request, it won't increment quota.
*   All other Static ones are in the order recommended by [the docs](https://kubernetes.io/docs/admin/admission-controllers/#is-there-a-recommended-set-of-plug-ins-to-use).  (which variously do mutation and validation) Preserves the behavior when there are no webhooks.
*   Ensures dynamic mutations happen before all validations.
*   Ensures dynamic validations happen after all mutations.
*   Users don't need to reason about the static ones, just the ones they add.

System administrators will likely need to know something about the webhooks they
intend to run in order to make the best ordering, but we will try to document a
good "first guess".

Validation continues to happen after all the admission controllers (e.g. after mutating webhooks, static admission controllers, and non-mutating admission controllers.)

**TODO**: we should move ResourceQuota after Validation, e.g. as described in #1137.  However, this is a longstanding bug and likely a larger change than can be done in 1.9--a larger quota redesign is out of scope. But we will likely make an improvement in the current ordering.


## Parallel vs Serial 

The main reason for parallel is reducing latency due to round trip and conversion.  We think this can often mitigated by consolidating multiple webhooks shared by the same project into one.

Reasons not to allow parallel are complexity of reasoning about concurrent patches, and CRD not supporting PATCH.

`ValidatingAdmissionWebhook `is already parallel,  and there are no responses to merge.  Therefore, it stays parallel.

`MutatingAdmissionWebhook `will run in serial, to ensure conflicts are resolved deterministically.

The order is the sort order of all the WebhookConfigs, by name, and by index within the Webhooks list.  

We don't plan to make mutating webhooks parallel at this time, but we will revisit the question in the future and decide before going to GA.

## Good Error Messages	 

When a webhook is persistently failing to allow e.g. pods to be created, then the error message from the apiserver must show which webhook failed.

When a core controller, e.g. ReplicaSet, fails to make a resources, it must send a helpful event that is visible in `kubectl describe` for the controlling resources, saying the reason create failed.

## Registering for all possible representations of the same object

Some Kubernetes resources are mounted in the api type system at multiple places
(e.g., during a move between groups). Additionally, some resources have multiple
active versions. There's not currently a way to easily tell which of the exposed
resources map to the same "storage location". We will not try to solve that
problem at the moment: if the system administrator wishes to hook all
deployments, they must (e.g.) make sure their hook is registered for both
deployments.v1beta1.extensions AND deployments.v1.apps.

This is likely to be error-prone, especially over upgrades. For GA, we may
consider mechanisms to make this easier. We expect to gather user feedback
before designing this.


## Conversion and Versioning

Webhooks will receive the admission review subject in the exact version which
the user sent it to the control plane. This may require the webhook to
understand multiple versions of those types.

All communication to webhooks will be JSON formatted, with a request body of
type admission.k8s.io/v1beta1. For GA, we will likely also allow proto, via a
TBD mechanism.

We will not take any particular steps to make it possible to know whether an
apiserver is safe to upgrade, given the webhooks it is running. System
administrators must understand the stack of webhooks they are running, watch the
Kubernetes release notes, and look to the webhook authors for guidance about
whether the webhook supports Kubernetes version N. We may choose to address this
deficency in future betas.

To follow the debate that got us to this position, you can look at this
potential design for the next steps: https://docs.google.com/document/d/1BT8mZaT42jVxtC6l14YMXpUq0vZc6V5MPf_jnzDMMcg/edit


## Mutations 

The Response for  `MutatingAdmissionWebhook`  must have content-type, and it must be one of:

*   `application/json`
*   `application/protobuf`
*   `application/strategic-merge-patch+json`
*   `application/json-patch+json`
*   `application/merge-json-patch+json`

If the response is a patch, it is merged with the versioned response from the previous webhook, where possible without Conversion.

We encourage the use of patch to avoid the "old clients dropping new fields" problem.


## Bootstrapping 

Bootstrapping (both turning on a cluster for the first time and making sure a
cluster can boot from a cold start) is made more difficult by having webhooks,
which are a dependency of the control plane. This is covered in its [own design
doc](./admission-webhook-bootstrapping.md).

## Upgrading the control plane

There are two categories of webhooks: security critical (e.g., scan images for
vulnerabilities) and nice-to-have (set labels).

Security critical webhooks cannot work with Kubernetes types they don't have
built-in knowledge of, because they can't know if e.g. Kubernetes 1.11 adds a
backwards-compatible `v1.Pod.EvilField` which will defeat their functionality.

They therefore need to be updated before any apiserver. It is the responsibility
of the author of such a webhook to release new versions in response to new
Kubernetes versions in a timely manner. Webhooks must support two consecutive
Kubernetes versions so that rollback/forward is possible. When/if Kubernetes
introduces LTS versions, webhook authors will have to also support two
consecutive LTS versions.

Non-security-critical webhooks can either be turned off to perform an upgrade,
or can just continue running the old webhook version as long as a completely new
version of an object they want to hook is not added. If they are metadata-only
hooks, then they should be able to run until we deprecate meta/v1. Such webhooks
should document that they don't consider themselves security critical, aren't
obligated to follow the above requirements for security-critical webhooks, and
therefore do not guarantee to be updated for every Kubernetes release.

It is expected that webhook authors will distribute config for each Kubernetes
version that registers their webhook for all the necessary types, since it would
be unreasonable to make system administrators understand all of the webhooks
they run to that level of detail.

## Support for Custom Resources 

Webhooks should work with Custom Resources created by CRDs.

They are particularly needed for Custom Resources, where they can supplement the validation and defaulting provided by OpenAPI.  Therefore, the webhooks will be moved or copied to genericapiserver for 1.9.


## Support for Aggregated API Servers 

Webhooks should work with Custom Resources on Aggregated API Servers.

Aggregated API Servers should watch apiregistraton on the main APIserver, and should identify webhooks with rules that match any of their resources, and call those webhooks.

For example a user might install a Webhook that adds a certain annotation to every single object.  Aggregated APIs need to support this use case.

We will build the dynamic admission stack into the generic apiserver layer to support this use case.


## Moving Built-in Admission Controllers 

This section summarizes recommendations for Posting static admission controllers to Webhooks.

See also [Details of Porting Admission Controllers](#details-of-porting-admission-controllers) and this [Backup Document](https://docs.google.com/spreadsheets/d/1zyCABnIzE7GiGensn-KXneWrkSJ6zfeJWeLaUY-ZmM4/edit#gid=0).

Here is an estimate of how each kind of admission controller would be moved (or not). This is to see if we can cover the use cases we currently have, not necessarily a promise that all of these will or should be move into another process.



*   Leave static:
    *   OwnerReferencesPermissionEnforcement 
        *   GC is a core feature of Kubernetes. Move to required.
    *   ResourceQuota 
        *   May [redesign](https://github.com/kubernetes/kubernetes/issues/51820)
        *   Original design doc says it remains static.
*   Divide into Mutating and non-mutating Webhooks
    *   PodSecurityPolicy
    *   NamespaceLifecycle
*   Use Mutating Webhook
    *   AlwaysPullImages
    *   ServiceAccount
    *   StorageClass
*   Use non-mutating Webhook
    *   Eventratelimit
    *   DenyEscalatingExec
    *   ImagePolicy 
        *   Need to standardize the webhook format 
    *   NodeRestriction
        *   Needs to be admission to access User.Info
    *   PodNodeSelector
    *   PodTolerationRestriction
*   Move to resource's validation or defaulting
    *   AntiAffinity
    *   DefaultTolerationSeconds
    *   PersistentVolumeClaimResize
    *   Initializers are reasonable to consider moving into the API machinery

For "Divide", the backend may well be different port of same binary, sharing a SharedInformer, so data is not cached twice.

For all Kubernetes built-in webhooks, the backend will likely be compiled into kube-controller-manager and share the SharedInformer.


# Use Case Analysis 


## Use Cases Detailed Descriptions 

Mutating Webhooks, Non-mutating webhooks, Initializers, and Finalizers (collectively, Object Lifecycle Extensions) serve to:



*   allow policy and behavioral changes to be developed independently of the control loops for individual Resources.  These might include company specific rules, or a PaaS that layers on top of Kubernetes.
*   implement business logic for Custom Resource Definitions
*   separate Kubernetes business logic from the core Apiserver logic, which increases reusability, security, and reliability of the core.

Specific Use cases:



*   Kubernetes static Admission Controllers
    *   Documented [here](https://kubernetes.io/docs/admin/admission-controllers/)
    *   Discussed [here](/contributors/design-proposals/api-machinery/admission_control_extension.md)
    *   All are highly reliable.  Most are simple.  No external deps.
    *   Many need update checks.
    *   Can be separated into mutation and validate phases.
*   OpenShift static Admission Controllers
    *   Discussed [here](/contributors/design-proposals/api-machinery/admission_control_extension.md)
    *   Similar to Kubernetes ones.
*   Istio, Case 1: Add Container to all Pods.
    *   Currently uses Initializer but can use Mutating Webhook.
    *   Simple, can be highly reliable and fast.  No external deps.
    *   No current use case for updates.
*   Istio, Case 2: Validate Mixer CRDs
    *   Checking cached values from other CRD objects.
    *   No external deps.
    *   Must check updates.
*   Service Catalog
    *   Watch PodPreset and edit Pods.
    *   Simple, can be highly reliable and fast. No external deps.
    *   No current use case for updates.

Good further discussion of use cases [here](/contributors/design-proposals/api-machinery/admission_control_extension.md)


## Details of Porting Admission Controllers 

This section summarizes which Kubernetes static admission controllers can readily be ported to Object Lifecycle Extensions.


### Static Admission Controllers 


<table>
  <tr>
   <td>Admission Controller
   </td>
   <td>How
   </td>
   <td>Why
   </td>
  </tr>
  <tr>
   <td>PodSecurityPolicy
   </td>
   <td>Use Mutating Webhook and Non-Mutating Webhook.
   </td>
   <td>Requires User.Info, so needs webhook.
<p>
Mutating will set SC from matching PSP.
<p>
Non-Mutating will check again in case any other mutators or initializers try to change it. 
   </td>
  </tr>
  <tr>
   <td>ResourceQuota
   </td>
   <td>Leave static
   </td>
   <td>A Redesign for Resource Quota has been proposed, to allow at least object count quota for other objects as well. This suggests that Quota might need to remain compiled in like authn and authz are.  
   </td>
  </tr>
  <tr>
   <td>AlwaysPullImages
   </td>
   <td>Use Mutating Webhook (could implement using initializer since the thing is it validating is forbidden to change by Update Validation of the object)
   </td>
   <td>Needs to 
   </td>
  </tr>
  <tr>
   <td>AntiAffinity
   </td>
   <td>Move to pod validation
   </td>
   <td>Since this is provided by the core project, which also manages the pod business logic, it isn't clear why this is even an admission controller.  Ask Scheduler people.
   </td>
  </tr>
  <tr>
   <td>DefaultTolerationSeconds
   </td>
   <td>Move to pod defaulting or use a Mutating Webhook.
   </td>
   <td>It is very simple.
   </td>
  </tr>
  <tr>
   <td>eventratelimit
   </td>
   <td>Non-mutating webhook
   </td>
   <td>Simple logic, does not mutate.  Alternatively, have rate limit be a built-in of api server.
   </td>
  </tr>
  <tr>
   <td>DenyEscalatingExec
   </td>
   <td>Non-mutating Webhook.
   </td>
   <td>It is very simple.  It is optional.  
   </td>
  </tr>
  <tr>
   <td>OwnerReferences- PermissionEnforcement (gc)
   </td>
   <td>Leave compiled in
   </td>
   <td>Garbage collection is core to Kubernetes.  Main and all aggregated apiservers should enforce it.
   </td>
  </tr>
  <tr>
   <td>ImagePolicy
   </td>
   <td>Non-mutating webhook
   </td>
   <td>Must use webhook since image can be updated on pod, and that needs to be checked.
   </td>
  </tr>
  <tr>
   <td>LimitRanger
   </td>
   <td>Mutating Webhook
   </td>
   <td>Fast 
   </td>
  </tr>
  <tr>
   <td>NamespaceExists
   </td>
   <td>Leave compiled in
   </td>
   <td>This has been on by default for years, right?
   </td>
  </tr>
  <tr>
   <td>NamespaceLifecycle
   </td>
   <td>Split:
<p>
 
<p>
Cleanup, leave compiled in. 
<p>
 
<p>
Protection of system namespaces: use non-mutating webhook
   </td>
   <td>
   </td>
  </tr>
  <tr>
   <td>NodeRestriction
   </td>
   <td>Use a non-mutating webhook
   </td>
   <td>Needs webhook so it can use User.Info.
   </td>
  </tr>
  <tr>
   <td>PersistentVolumeClaimResize
   </td>
   <td>Move to validation
   </td>
   <td>This should be in the validation logic for storage class.
   </td>
  </tr>
  <tr>
   <td>PodNodeSelector
   </td>
   <td>Move to non-mutating webhook 
   </td>
   <td>Already compiled in, so fast enough to use webhook.  Does not mutate.
   </td>
  </tr>
  <tr>
   <td>podtolerationrestriction
   </td>
   <td>Move to non-mutating webhook 
   </td>
   <td>Already compiled in, so fast enough to use webhook.  Does not mutate.
   </td>
  </tr>
  <tr>
   <td>serviceaccount
   </td>
   <td>Move to mutating webhook.
   </td>
   <td>Already compiled in, so fast enough to use webhook.  Does mutate by defaulting the service account.
   </td>
  </tr>
  <tr>
   <td>storageclass
   </td>
   <td>Move to mutating webhook.
   </td>
   <td>
   </td>
  </tr>
</table>


[Backup Document](https://docs.google.com/spreadsheets/d/1zyCABnIzE7GiGensn-KXneWrkSJ6zfeJWeLaUY-ZmM4/edit#gid=0)


### OpenShift Admission Controllers 


<table>
  <tr>
   <td>Admission Controller
   </td>
   <td>How
   </td>
   <td>Why
   </td>
  </tr>
  <tr>
   <td>pkg/authorization/admission/restrictusers"
   </td>
   <td>Non-mutating Webhook or leave static
   </td>
   <td>Verification only. But uses a few loopback clients to check other resources.
   </td>
  </tr>
  <tr>
   <td>pkg/build/admission/jenkinsbootstrapper
   </td>
   <td>Non-mutating Webhook or leave static
   </td>
   <td>Doesn't mutate Build or BuildConfig, but creates Jenkins instances.
   </td>
  </tr>
  <tr>
   <td>pkg/build/admission/secretinjector
   </td>
   <td>Mutating webhook or leave static
   </td>
   <td>uses a few loopback clients to check other resources.
   </td>
  </tr>
  <tr>
   <td>pkg/build/admission/strategyrestrictions
   </td>
   <td>Non-mutating Webhook or leave static
   </td>
   <td>Verifications only. But uses a few loopback clients, and calls subjectAccessReview
   </td>
  </tr>
  <tr>
   <td>pkg/image/admission
   </td>
   <td>Non-Mutating Webhook
   </td>
   <td>Fast, checks image size
   </td>
  </tr>
  <tr>
   <td>pkg/image/admission/imagepolicy
   </td>
   <td>Mutating and non-mutating webhooks
   </td>
   <td>Rewriting image pull spec is mutating.
<p>
acceptor.Accepts is non-Mutating
   </td>
  </tr>
  <tr>
   <td>pkg/ingress/admission
   </td>
   <td>Non-mutating webhook, or leave static.
   </td>
   <td>Simple, but calls to authorizer.
   </td>
  </tr>
  <tr>
   <td>pkg/project/admission/lifecycle
   </td>
   <td>Initializer or Non-mutating webhook?
   </td>
   <td>Needs to update another resource: Namespace
   </td>
  </tr>
  <tr>
   <td>pkg/project/admission/nodeenv
   </td>
   <td>Mutating webhook
   </td>
   <td>Fast
   </td>
  </tr>
  <tr>
   <td>pkg/project/admission/requestlimit
   </td>
   <td>Non-mutating webhook
   </td>
   <td>Fast, verification only
   </td>
  </tr>
  <tr>
   <td>pkg/quota/admission/clusterresourceoverride
   </td>
   <td>Mutating webhook
   </td>
   <td>Updates container resource request and limit
   </td>
  </tr>
  <tr>
   <td>pkg/quota/admission/clusterresourcequota
   </td>
   <td>Leave static.
   </td>
   <td>Refactor with the k8s quota
   </td>
  </tr>
  <tr>
   <td>pkg/quota/admission/runonceduration
   </td>
   <td>Mutating webhook
   </td>
   <td>Fast. Needs a ProjectCache though. Updates pod.Spec.ActiveDeadlineSeconds
   </td>
  </tr>
  <tr>
   <td>pkg/scheduler/admission/podnodeconstraints
   </td>
   <td>Non-mutating webhook or leave static
   </td>
   <td>Verification only. But calls to authorizer.
   </td>
  </tr>
  <tr>
   <td>pkg/security/admission
   </td>
   <td>Use Mutating Webhook and Non-Mutating Webhook.
   </td>
   <td>Similar to PSP in k8s
   </td>
  </tr>
  <tr>
   <td>pkg/service/admission/externalip
   </td>
   <td>Non-mutating webhook
   </td>
   <td>Fast and verification only
   </td>
  </tr>
  <tr>
   <td>pkg/service/admission/endpoints
   </td>
   <td>Non-mutating webhook or leave static
   </td>
   <td>Verification only. But calls to authorizer.
   </td>
  </tr>
</table>



### Other Projects 

Istio Pod Injector:



*   Injects Sidecar Container, Init Container, adds a volume for Istio config, and changes the Security Context
*   Source:
    *   https://github.com/istio/pilot/blob/master/platform/kube/inject/inject.go#L278 
    *   https://github.com/istio/pilot/blob/master/cmd/sidecar-initializer/main.go 

<table>
  <tr>
   <td>
Function
   </td>
   <td>How
   </td>
   <td>Why
   </td>
  </tr>
  <tr>
   <td>Istio Pod Injector
   </td>
   <td>Mutating Webhook
   </td>
   <td>Containers can only be added at pod creation time.
<p>
Because the change is complex, showing intermediate state may help debugging.
<p>
Fast, so could also use webhook.
   </td>
  </tr>
  <tr>
   <td>Istio Mixer CRD Validation
   </td>
   <td>Non-Mutating Webhook
   </td>
   <td>
   </td>
  </tr>
  <tr>
   <td>Service Catalog PodPreset
   </td>
   <td>Initializer
   </td>
   <td>Containers can only be added at pod creation time.
<p>
Because the change is complex, showing intermediate state may help debugging.
<p>
Fast, so could also use webhook.
   </td>
  </tr>
  <tr>
   <td>Allocate Cert for Service
   </td>
   <td>Initializer 
   </td>
   <td>Longer duration operation which might fail, with external dependency, so don't use webhook.
<p>
Let user see initializing state.
<p>
Don't let controllers that depend on services see the service before it is ready.
   </td>
  </tr>
</table>



## Comparison of Webhooks and Initializers 


<table>
  <tr>
   <td>Mutating and Non-Mutating Webhooks
   </td>
   <td>Initializers (and Finalizers)
   </td>
  </tr>
  <tr>
   <td><ul>

<li>Act on Create, update, or delete
<li>Reject Create, Update or delete</li></ul>

   </td>
   <td><ul>

<li>Act on Create and delete 
<li>Reject Create.</li></ul>

   </td>
  </tr>
  <tr>
   <td><ul>

<li>Clients never see pre-created state. <ul>

 <li>Good for enforcement.
 <li>Simple invariants.</li> </ul>
</li> </ul>

   </td>
   <td><ul>

<li>Clients can see pre-initialized state. <ul>

 <li>Let clients see progress 
 <li>Debuggable</li> </ul>
</li> </ul>

   </td>
  </tr>
  <tr>
   <td><ul>

<li>Admin cannot easily override broken webhook. <ul>

 <li>Must be highly reliable code
 <li>Avoid deps on external systems.</li> </ul>
</li> </ul>

   </td>
   <td><ul>

<li>Admin can easily fix a "stuck" object by "manually" initializing (or finalizing). <ul>

 <li>Can be <em>slightly</em> less reliable.
 <li>Prefer when there are deps on external systems.</li> </ul>
</li> </ul>

   </td>
  </tr>
  <tr>
   <td><ul>

<li>Synchronous <ul>

 <li>Apiserver uses a go routine
 <li>TCP connection open
 <li>Should be very low latency</li> </ul>
</li> </ul>

   </td>
   <td><ul>

<li>Asynchronous <ul>

 <li>Can be somewhat higher latency</li> </ul>
</li> </ul>

   </td>
  </tr>
  <tr>
   <td><ul>

<li>Does not persist intermediate state <ul>

 <li>Should happen very quickly.
 <li>Does not increase etcd traffic.</li> </ul>
</li> </ul>

   </td>
   <td><ul>

<li>Persist intermediate state <ul>

 <li>Longer ops can persist across apiserver upgrades/failures
 <li>Does increase etcd traffic.</li> </ul>
</li> </ul>

   </td>
  </tr>
  <tr>
   <td><ul>

<li>Webhook does not know if later webhooks fail <ul>

 <li>Must not have side effects,
 <li>Or have a really good GC plan.</li> </ul>
</li> </ul>

   </td>
   <td><ul>

<li>Initializer does not know if later initializers fail, but if paired with a finalizer, it could see the resource again. <ul>

 <li>This is not implemented
 <li>TODO: initializers: have a way to ensure finalizer runs even if later initializers reject?</li> </ul>
</li> </ul>

   </td>
  </tr>
  <tr>
   <td>
    Use Examples:<ul>

<li>checking one field on an object,  and setting another field on the same object</li></ul>

   </td>
   <td>
    Use Examples:<ul>

<li>Allocate (and deallocate) external resource in parallel with a Kubernetes resource.</li></ul>

   </td>
  </tr>
</table>


Another [Detailed Comparison of Initializers and Webhooks](https://docs.google.com/document/d/17P_XjXDpxDC5xSD0nMT1W18qE2AlCMkVJcV6jXKlNIs/edit?ts=59d5683b#heading=h.5irk4csrpu0y)
