# Webhooks Beta


## PUBLIC
Authors: @erictune, @caesarxuchao, @enisoc
Thanks to: {@dbsmith,  @smarterclayton, @deads2k, @cheftako, @jpbetz, @mbohlool, @mml, @janetkuo} for comments, data, prior designs, etc.  


[TOC]


**Discussion and Decision for this Beta plan desired at 11 October SIG-API-Machinery meeting, due to short release cycle in 1.9.**


# Summary 

This document proposes a detailed plan for bringing Webhooks to Beta. Highlights include (incomplete, see rest of doc for complete list) :



*   Adding the ability for webhooks to mutate.  
*   Bootstrapping
*   Monitoring
*   Versioned rather than Internal data sent on hook
*   Ordering behavior within webhooks, and with other admission phases, is better defined

This plan is compatible with the [original design doc]( https://github.com/kubernetes/community/blob/master/contributors/design-proposals/api-machinery/admission_control_extension.md).


# Definitions 

**Mutating Webhook**: Webhook that can change a request as well as accept/reject.  

**Non-Mutating Webhook**: Webhook that cannot change request, but can accept or reject.

**Webhook**: encompasses both Mutating Webhook and/or Non-mutating Webhook.

**Validating Webhook**: synonym for Non-Mutating Webhook

**Static Admission Controller**: Compiled-in Admission Controllers, (in plugin/pkg/admission).

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
*   Internal Go interface refactor (e.g. along the lines suggested  #[1137](https://github.com/kubernetes/community/pull/1137)).  Nice-to-have but not gating beta.   


## Work Items 



*   Add API for registering mutating webhooks.  See 

<p id="gdcalert1" ><span style="color: red; font-weight: bold">>>>>  GDC alert: undefined internal link (link text: "API Changes"). Did you generate a TOC? </span><br>(<a href="#">Back to top</a>)(<a href="#gdcalert2">Next alert</a>)<br><span style="color: red; font-weight: bold">>>>> </span></p>

[API Changes](#heading=h.5xu31klaqsa0)
*   Copy the non-mutating webhook admission controller code and rename it to be for mutating. (Splitting into two registration APIs make ordering clear.) Add changes to handle mutating responses.  See 

<p id="gdcalert2" ><span style="color: red; font-weight: bold">>>>>  GDC alert: undefined internal link (link text: "Responses for Mutations"). Did you generate a TOC? </span><br>(<a href="#">Back to top</a>)(<a href="#gdcalert3">Next alert</a>)<br><span style="color: red; font-weight: bold">>>>> </span></p>

[Responses for Mutations](#heading=h.a9rxps4ka8y3).
*   Document recommended flag order for admission plugins.  See [Order of Admission](#order-of-admission).
*   In kube-up.sh and other installers, change flag per previous item. 
*   Ensure able to monitor latency and rejection from webhooks. See [Monitorability](#monitorability).
*   Don't send internal objects.   See [#49733](https://github.com/kubernetes/kubernetes/issues/49733)
*   Serialize mutating Webhooks into order in the apiregistration. Leave non-mutating in parallel.
*   Good Error Messages	.  See [Good Error Messages](#good-error-messages)	
*   Conversion logic in GenericWebhook to send converted resource to webhook.  See [Conversion](#conversion) and [#49733](https://github.com/kubernetes/kubernetes/issues/49733).
*   Schedule discussion around resiliency to down webhooks and bootstrapping


# Design Discussion 


## Why Webhooks First  

We will do webhooks beta before initializers beta because:



1.  **Serves Most Use Cases**: We reviewed code of all current use cases, namely: Kubernetes Built-in Admission Controllers, OpenShift Admission Controllers, Istio & Service Catalog.  (See also [Use Cases Detailed Descriptions](#use-cases-detailed-descriptions).)  All of those use cases are well served by mutating and non-mutating webhooks.  (See also [Comparison of Webhooks and Initializers](#comparison-of-webhooks-and-initializers)).
1.  **Less Work**: An engineer quite experienced with both code bases estimated that it is less work to adding Mutating Webhooks and bring both kinds of webhooks to beta; than to bring non-mutating webhooks and initializers to Beta.  Some open issues with Initializers with long expected development time include quota replenishment bug, and controller awareness of uninitialized objects.
1.  **API Consistency**: Prefer completing one related pair of interfaces (both kinds of webhooks) at the same time.


## Why Support Mutation for Beta

Based on experience and feedback from the alpha phase of both Webhooks and Initializers, we believe Webhooks Beta should support mutation because:  



1.   we have lots of use cases to inform this (both from Initializers, and Admission Controllers) to ensure we have needed features
1.  We have experience with Webhooks API already to give confidence in the API.  The registration API will be quite similar except in the responses.
1.  There is a strong community demand for something that satisfies a mutating case.


## Plan for Existing Initializer-Clients 

After the release of 1.9, we will advise users who currently use initializers to:



*   Move to Webhooks if their use case fits that model well.
*   Provide SIG-API-Machinery with feedback if Initializers is a better fit. 

We will continue to support Initializers as an Alpha API in 1.9.

We will point out https://github.com/caesarxuchao/example-webhook-admission-controller and extend it to support mutating webhooks, and as that Initializer docs point to it (e.g. https://github.com/kelseyhightower/kubernetes-initializer-tutorial)


## Monitorability 

There should be prometheus variables to show:



*   API operation latency
    *   Overall
    *   By webhook name
*   API response codes
    *   Overall
    *   By webhook name.

Adding a webhook dynamically adds a key to a map-valued prometheus metric.


## API Changes 

GenericAdmissionWebhook Admission Controller is duplicated.



*   One is called `MutatingAdmissionWebhook`
*   The other is called `ValidatingAdmissionWebhook`
*   Splitting them allows them to appear in different places in the `--admission-control` flag's order.

ExternalAdmissionHookConfiguration API is duplicated.



*   One is called  `MutatingAdmissionWebhookConfiguration`
*   The other is called `ValidatingAdmissionWebhookConfiguration`
*   Splitting them:
    *   makes it clear what the order is when some items don't have both flavors,
    *   enforces mutate-before-validate,
    *   better allows declarative update of the config than one big list with an implied partition point

The `ValidatingAdmissionWebhookConfiguration` stays the same as `ExternalAdmissionHookConfiguration` except it moves to v1beta1.

The `MutatingAdmissionWebhookConfiguration` is the same API as `ValidatingAdmissionWebhookConfiguration`.  It is only visible via the v1beta1 version.

The Group/Version called 

`admissionregistration.k8s.io/v1alpha1` with kinds

InitializerConfiguration and ExternalAdmissionHookConfiguration.

InitializerConfiguration will not join `admissionregistration.k8s.io/v1beta1` at this time.

Any webhooks that register with v1alpha1 may or may not be surprised when they start getting Versioned data.  But we don't make any promises for Alpha, and detecting which API was used to register is too complex.


## Order of Admission 

The recommended order on kubernetes.io for `--admission-control` will be: `MutatingAdmissionWebhook,NamespaceLifecycle,LimitRanger,ServiceAccount,PersistentVolumeLabel,DefaultStorageClass,DefaultTolerationSeconds,ValidatingAdmissionWebhook,ResourceQuota`

 This order has the following properties:



*   ResourceQuota comes last, so that if prior ones reject a request, it won't increment quota.
*   All other Static ones are in the order recommended by [the docs](https://kubernetes.io/docs/admin/admission-controllers/#is-there-a-recommended-set-of-plug-ins-to-use).  (which variously do mutation and validation) Preserves the behavior when there are no webhooks.
*   Ensures dynamic mutations happen before all validations.
*   Ensures dynamic validations happen after all mutations.
*   Users don't need to reason about the static ones, just the ones they add.

Validation continues to happen after all the admission controllers (e.g. after mutating webhooks, static admission controllers, and non-mutating admission controllers.)

**TODO**: we should move ResourceQuota after Validation, e.g. as described in #1137.  However, this is a longstanding bug and likely a larger change than can be done in 1.9, and should be revisted along with a larger quota redesign.  Therefore, it is not a requirement for Beta,


## Parallel vs Serial 

The main reason for parallel is reducing latency due to round trip and converstion.  We think this can often mitigated by consolidating multiple webhooks shared by the same project into one.

Reasons not to allow parallel are complexity of reasoning about concurrent patches, and CRD not supporting PATCH.

`ValidatingAdmissionWebhook `is already parallel,  and there are no responses to merge.  Therefore, it stays parallel.

`MutatingAdmissionWebhook `will run in serial, to ensure conflicts are resolved deterministically.

The order is the sort order of all the WebhookConfigs, by name, and by index within the Webhooks list.  

If a strong need arises for parallel mutating webhooks, the API will allow adding a partial order later as follows:  Add a "parallelStage" field to WebhookConfig object.  Any objects for which this field is unset all run in serial first.  Then ones which have the phase set are run.  All the ones with the same value run in parallel, and different ones in the order based on the numeric (or lexicographical) order by values).


## Good Error Messages	 

When a webhook is persistently failing to allow e.g. pods to be created, then the error message from the apiserver must show which webhook failed.

When a core controller, e.g. ReplicaSet, fails to make a resources, it must send a helpful event that is visible in `kubectl describe` for the controlling resources, saying the reason create failed.


## Conversion 

The apiserver will send versioned objects to the webhook to avoid churns in the webhook code. The current registration API has an Rule.[APIVersions](https://github.com/kubernetes/kubernetes/blob/master/pkg/apis/admissionregistration/types.go#L86) field, which is used to match against the version in the endpoint url, and only matching requests will be forwarded to the webhook. The current API isn't able to specify what version the APIServer should convert the request body to before forwarding the it to the webhook. So we propose two ways to modify the API.

Plan A:

We change the semantic of the APIVersions field. The field means the external versions the webhooks is able to decode. The APIServer forwards the webhook any requests that are stored at the same etcd path as the one appears in the registration API.

Plan B:

The semantic of the APIVersions field is unchanged. We add a ConvertToGroup and ConvertToVersion field to express the versions the webhooks understand. A straw-man design is [here](https://docs.google.com/document/d/1CPcRQ12dIxIE1_T8BYx41V2w5WEGP0vHAWCa7HpqsVg/edit#heading=h.8emrnnxujf1k).

It's TBD if we go with Plan A or B. 

It is possible that two different webhooks want to access the same stored-resource via different GVKs.  In this case there could be extra conversions.  Conversion do take some CPU.  However, we don't expect this to be a significant issue for beta because:



*   Pods are by far the most common core object to webhook, based on the  [Use Cases](#use-cases-detailed-descriptions), and pods do not have multiple versions.  Most or all other resources we know we plan to Webhook also do not have multiple versions.  So extra conversion is not needed in practice.
*   Reads are ~10x more common that writes in a typical cluster, so any increase in compute cost for writes to multi-version objects is significantly mitigated.

If the administrator chooses to install webhooks that all use the same GVKs for a given internal type, then it will be more efficient, since the apiserver may not need to convert every time: it just needs to patch the json, and then send it back out again without converting to internal.  

For `MutatingAdmissionWebhook`, if there is a "*" match then the apiserver is allowed to send whatever GVK it wants.  The webhook needs to handle any type.  For star-match, we think the webhook will usually only be looking at the metadata (e.g. the gc admission controller), or looking at event only (e.g. object count quota).  


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

We will add a WebhookConfig.Webhooks.exemptNamespace (or  an WebhookConfig.Webhooks.Rule.exemptNamespace?  This can exempt  some namespaces from webhooks (See #51290) for bootstrapping.

TODO: detail the bootstrapping procedure for a webhook that is not released with Kubernetes core.  Using a special namespace that is exempt from initializers.


## Support for Custom Resources 

Webhooks should work with Custom Resources created by CRDs.

They are particularly needed for Custom Resources, where they can supplement the validation and defaulting provided by OpenAPI.  Therefore, the webhooks will be moved or copied to genericapiserver for 1.9.


## Support for Aggregated API Servers

Webhooks should work with Custom Resources on Aggregated API Servers.

Aggregated API Servers should watch apiregistraton on the main APIserver, and should identify webhooks with rules that match any of their resources, and call those webhooks.

For example a user might install a WEbhook that adds a certain annotation to every single object.  Aggregated APIs need to support this use case.


## Moving Built-in Admission Controllers 

This section summarizes recommendations for Posting static admission controllers to Webhooks.

See also [Details of Porting Admission Controllers](#details-of-porting-admission-controllers) and this [Backup Document](https://docs.google.com/spreadsheets/d/1zyCABnIzE7GiGensn-KXneWrkSJ6zfeJWeLaUY-ZmM4/edit#gid=0).

Here is an estimate of how each kind of admission controller would be moved (or not).



*   Leave static:
    *   OwnerReferencesPermissionEnforcement 
        *   GC is a core feature of Kubernetes.  Move to required.
    *   ResourceQuota 
        *   May [redesign](https://github.com/kubernetes/kubernetes/issues/51820)
        *   Original design doc says it remains static.
*   Divide into Mutating and non-mutating  Webhooks
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
    *   Discussed [here](https://github.com/kubernetes/community/blob/master/contributors/design-proposals/api-machinery/admission_control_extension.md)
    *   All are highly reliable.  Most are simple.  No external deps.
    *   Many need update checks.
    *   Can be separated into mutation and validate phases.
*   OpenShift static Admission Controllers
    *   Discussed [here](https://github.com/kubernetes/community/blob/master/contributors/design-proposals/api-machinery/admission_control_extension.md)
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

Good further discussion of use cases [here](https://github.com/kubernetes/community/blob/master/contributors/design-proposals/api-machinery/admission_control_extension.md)


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
 <li>Should be very low latecny</li> </ul>
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
