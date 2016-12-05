# Extension of Admission Control via Initializers and External Admission Enforcement

Admission control is the primary business-logic policy and enforcement subsystem in Kubernetes. It provides synchronous
hooks for all API operations and allows an integrator to impose additional controls on the system - rejecting, altering,
or reacting to changes to core objects. Today each of these plugins must be compiled into Kubernetes. As Kubernetes grows,
the requirement that all policy enforcement beyond coarse grained access control be done through in-tree compilation and
distribution becomes unweildy and limits administrators and the growth of the ecosystem.

This proposal covers changes to the admission control subsystem that allow extension of admission without recompilation
and dynamic adminission control configuration in ways that resemble existing controller behavior.


## Background

The four core systems in Kubernetes are:

1. API servers with persistent storage, providing basic object validation and CRUD operations
2. Authentication and authorization layers that identify an actor and constrain the coarse actions that actor can take on API objects
3. [Admission controller layers](admission_control.md) that can control and limit the CRUD operations clients perform synchronously.
4. Controllers which watch the API and react to changes made by other users asynchronously (scheduler, replication controller, kubelet, kube-proxy, and ingress are all examples of controllers).

Admission control supports a wide range of policy and behavior enforcement for administrators.


### Types of Admission Control

In Kubernetes 1.5 and OpenShift 1.4, the following types of functionality have been implemented through admission
(all file references are relative to `plugin/pkg/admission`, or simply identified by name for OpenShift). Many of the
Kubernetes admission controllers originated in OpenShift and are listed in both for history.

#### Resource Control

These admission controllers take resource usage for pods into account to ensure namespaces cannot abuse the cluster
by consuming more than their fair share of resources. These perform security or defaulting type roles.

##### Kubernetes

Name | Code | Description
---- | ---- | -----------
InitialResources | initialresources/admission.go | Default the resources for a container based on past usage
LimitRanger | limitranger/admission.go | Set defaults for container requests and limits, or enforce upper bounds on certain resources (no more than 2GB of memory, default to 512MB)
ResourceQuota | resourcequota/admission.go | Calculate and deny number of objects (pods, rc, service load balancers) or total consumed resources (cpu, memory, disk) in a namespace

##### OpenShift

Name | Code | Description
---- | ---- | -----------
ClusterResourceOverride | clusterresourceoverride/admission.go | Allows administrators to override the user's container request for CPU or memory as a percentage of their request (the administrator's target overcommit number), or to default a limit based on a request. Allows cluster administrators to control overcommit on a cluster.
ClusterResourceQuota | clusterresourcequota/admission.go | Performs quota calculations over a set of namespaces with a shared quota. Can be used in conjunction with resource quota for hard and soft limits.
ExternalIPRanger | externalip_admission.go | Prevents users from creating services with externalIPs inside of fixed CIDR ranges, including the pod network, service network, or node network CIDRs to prevent hijacking of connections.
ImageLimitRange | admission.go | Performs LimitRanging on images that are pushed into the integrated image registry
OriginResourceQuota | resourcequota/admission.go | Performs quota calculations for API resources exposed by OpenShift. Demonstrates how quota would be implemented for API extensions.
ProjectRequestLimit | requestlimit/admission.go | A quota on how many namespaces may be created by any individual user. Has a global default and also a per user override.
RunOnceDuration | runonceduration/admission.go | Enforces a maximum ActiveDeadlineSeconds value on all RestartNever pods in a namespace. This ensures that users are defaulted to have a deadline if they did not request it (which prevents pathological resource consumption)

Quota is typically last in the admission chain, to give all other components a chance to reject or modify the resource.


#### Security

These controllers defend against specific actions within a resource that might be dangerous that the authorization
system cannot enforce.

##### Kubernetes

Name | Code | Description
---- | ---- | -----------
AlwaysPullImages | alwayspullimages/admission.go | Forces the Kubelet to pull images to prevent pods from from accessing private images that another user with credentials has already pulled to the node.
LimitPodHardAntiAffinityTopology | antiaffinity/admission.go | Defended the cluster against abusive anti-affinity topology rules that might hang the scheduler.
DenyEscalatingExec | exec/admission.go | Prevent users from executing into pods that have higher privileges via their service account than allowed by their policy (regular users can't exec into admin pods).
DenyExecOnPrivileged | exec/admission.go | Blanket ban exec access to pods with host level security. Superceded by DenyEscalatingExec
OwnerReferencesPermissionEnforcement | gc/gc_admission.go | Require that a user who sets a owner reference (which could result in garbage collection) has permission to delete the object, to prevent abuse.
ImagePolicyWebhook | imagepolicy/admission.go | Invoke a remote API to determine whether an image is allowed to run on the cluster.
PodNodeSelector | podnodeselector/admission.go | Default and limit what node selectors may be used within a namespace by reading a namespace annotation and a global configuration.
PodSecurityPolicy | security/podsecuritypolicy/admission.go | Control what security features pods are allowed to run as based on the end user launching the pod or the service account. Sophisticated policy rules.
SecurityContextDeny | securitycontext/scdeny/admission.go | Blanket deny setting any security context settings on a pod.

##### OpenShift

Name | Code | Description
---- | ---- | -----------
BuildByStrategy | strategyrestrictions/admission.go | Control which types of image builds a user can create by checking for a specific virtual authorization rule (field level authorization), since some build types have security implications.
OriginPodNodeEnvironment | nodeenv/admission.go | Predecessor to PodNodeSelector.
PodNodeConstraints | podnodeconstraints/admission.go | Prevent users from setting nodeName directly unless they can invoke the `bind` resource on pods (same as a scheduler). This prevents users from attacking nodes by repeatedly creating pods that target a specific node and forcing it to reject those pods. (field level authorization)
RestrictedEndpointsAdmission | endpoint_admission.go | In a multitenant network setup where namespaces are isolated like OpenShift SDN, service endpoints must not allow a user to probe other namespaces. If a user edits the endpoints object and sets IPs that fall within the pod network CIDR, the user must have `create` permission on a virtual resource `endpoints/restricted`. The service controller is granted this permission by default.
SecurityContextConstraint | admission.go | Predecessor to PodSecurityPolicy.
SCCExecRestrictions | scc_exec.go | Predecessor to DenyEscalatingExec.

Many other controllers have been proposed, including but not limited to:

* Control over what taints and tolerations a user can set on a pod
* Control over which labels and annotations can be set or changed
* Generic control over which fields certain users may set (field level access control)


#### Defaulting / Injection

These controllers inject namespace or cluster context into pods and other resources at runtime to decouple
application config from runtime config (separate the user's pod settings from environmental controls)

##### Kubernetes

Name | Code | Description
---- | ---- | -----------
ServiceAccount | serviceaccount/admission.go | Bind mount the service account token for a pod into the pod at a specific location.
PersistentVolumeLabel | persistentvolume/label/admission.go | Lazily bind persistent volume claims to a given zone when a pod is scheduled.
DefaultStorageClass | storageclass/default/admission.go | Set a default storage class on any PVC created without a storage class.

Many other controllers have been proposed, including but not limited to:

* ServiceInjectionPolicy to inject environment, configmaps, and secrets into pods that reference those services
* Namespace level environment injection (all pods in this namespace should have env var `ENV=PROD`)
* Label selector based resource defaults (all pods with these labels get these default resources)


#### Referential Consistency

These controllers enforce that certain guarantees of the system related to integrity.

##### Kubernetes

Name | Code | Description
---- | ---- | -----------
NamespaceAutoProvision | namespace/autoprovision/admission.go | When users create resources in a namespace that does not exist, ensure the namespace is created so it can be seen with `kubectl get namespaces`
NamespaceExists | namespace/exists/admission.go | Require that a namespace object exist prior to a resource being created.
NamespaceLifecycle | namespace/lifecycle/admission.go | More powerful and flexible version of NamespaceExists.

##### OpenShift

Name | Code | Description
---- | ---- | -----------
JenkinsBootstrapper | jenkinsbootstrapper/admission.go | Spawn a Jenkins instance in any project where a Build is defined that references a Jenkins pipeline. Checks that the creating user has permission to act-as an editor in the project to prevent escalation within a namespace.
ImagePolicy | imagepolicy/imagepolicy.go | Performs policy functions like ImagePolicyWebhook, but also is able to mutate the image reference from a tag to a digest (fully qualified spec), look up additional information about the image from the OpenShift Image API and potentially enforce resource consumption or placement decisions based on the image. May also be used to deny images from being used that don't resolve to image metadata that OpenShift tracks.
OriginNamespaceLifecycle | lifecycle/admission.go | Controls accepting resources for namespaces.


### Patterns

In a study of all known admission controllers, the following patterns were seen most often:

1. Defaulting on creation
2. Synchronous validation on creation
3. Synchronous validation on update - side-effect free

Other patterns seen less frequently include:

1. Defaulting on update
2. Resolving / further specifying values on update (ImagePolicy)
3. Creating resources in response to user action with the correct permission check (JenkinsBootstrapper)
4. Policy decisions based on *who* is doing the action (JenkinsBootstrapper)
5. Synchronous validation on update - with side effects (quota)

While admission controllers can operate on all verbs, resources, and sub resource types, in practice they
mostly deal with create and update on primary resources. Most sub resources are highly privileged operations
and so are typically covered by policy. Other controllers like quota tend to be per apiserver and therefore
are not required to be extensible.


### Building enforcement

In order to implement custom admission, an admin, integrator, or distribution of Kubernetes must compile their
admission controller(s) into the Kubernetes `kube-apiserver` binary. As Kubernetes is intended to be a
modular layered system this means core components must be upgraded to effect policy changes and only a fixed
list of plugins can be used. It also prevents experimentation and prototyping of policy, or "quick fix"
solutions applied on site. As we add additional APIs that are not hosted in the main binary (either as third
party resources or API extension servers), these APIs have many of the same security and policy needs that
the core resources do, but must compile in their own subsets of admission.

Further, distributions of Kubernetes like OpenShift that wish to offer complete solutions (such as OpenShift's
multi-tenancy model) have no mechanism for running on top of Kubernetes without recompilation of the core or
for extending the core with additional policy. This prevents the formation of an open ecosystem for tools
*around* Kubernetes, forcing all changes to policy to go through the Kubernetes codebase review gate (when
such review is unnecessary or disruptive to Kubernetes itself).


## Design

It should be possible to perform holistic policy enforcement in Kubernetes without the recompilation of the
core project as plugins that can be added and removed to a stock Kubernetes release. That extension
of admission control should leverage similar mechanisms to our existing controller frameworks where possible
and otherwise be performant and reliable.


### Requirements

1.  Easy Initialization
    Privileged components should be able to easily participate in the **initialization** of a new object.
2.  Synchronous Validation
    Synchronous rejection of initialized objects or mutations must be possible outside of the kube-apiserver binary
3.  Backwards Compatible
    Existing API clients must see no change in behavior to external admission other than increased latency
4.  Easy Installation
    Administrators should be able to easily write a new admission plugin and deploy it in the cluster
5.  Performant
    External admission must not significantly regress performance in large and dense clusters
6.  Reliable
    External admission should be capable of being "production-grade" for deployment in an extremely large and dense cluster
7.  Internally Consistent
    Developing an admission controller should reuse as much infrastructure and tools as possible from building custom controllers so as to reduce the cost of extension.


### Specification

Based on observation of the actual admission control implementations the majority of mutation
occurs as part of creation, and a large chunk of the remaining controllers are for side-effect free
validation of creation and updates. Therefore we propose the following changes to Kubernetes:

1.  Allow some controllers to act as "initializers" - watching the API and mutating the object before it is visible to normal clients.
    This would reuse the majority of the infrastructure in place for controllers. Because creation is
    one-way, the object can be "revealed" to regular clients once a set list of initializers is consumed. These
    controllers could run on the cluster as pods.
2.  Add a generic **external admission webhook** controller that is non-mutating (thus parallelizable)
    This generic webhook API would resemble `admission.Interface` and be given the input object (for create) and the
    previous object (for update). After initialization or on any update, these hooks would be invoked in parallel
    against the remote servers and any rejection would reject the mutation.
3.  Make the registration of both initializers and admission webhooks dynamic via the API (a configmap or cluster scoped resource)
    Administrators should be able to dynamically add or remove hooks and initializers on demand to the cluster.
    Configuration would be similar to registering new API group versions and include config like "fail open" or
    "fail closed".

Some admission controller types would not be possible for these extensions:

* Mutating admission webhooks could be a later addition
* Admission controllers that need access to the acting user can receive that via the external webhook.
* Admission controllers that "react" to the acting user can couple the information received via a webhook and then act if they observe mutation succeed (tuple combining resource UID and resource generation).
* Quota will continue to be a core plugin per API server, so extension is not critical.



#### Implications:

1. Initializers and generic admission controllers are highly privileged, so while some separation is valuable they are effectively cluster scoped
2. This mechanism would allow dedicated infrastructure to host admission for multiple clusters, and allow some expensive admission to be centralized (like quota which is hard to performantly distribute)
3. There is no way to build initializers for updates without a much more complicated model, but we anticipate initializers to work best on creation.
4. Ordering will probably be necessary on initializers because defaulting in the wild requires ordering. Non-mutating validation on the other hand can be fully parallel.
5. Some admission depends on knowing the identity of the actor - we will likely need to include the **creator** as information to initializers.
6. Quota must still run after all validators are invoked. We may need to make quota extensible in the future.


### TODO all implementation details


### Alternatives considered

The following are all viable alternatives to this specification, but have some downsides against the requirements above.
There should be no reason these could not be implemented for specific use cases.

1.  Admission controller that can run shell commands inside its context to mutate objects.
    * Limits on performance and reliability
    * Requires the masters be updated (can't be done dynamically)
2. Admission controller that can run a scripting language like Lua or JavaScript in process.
    * Limits on performance and reliability
    * Not consistent with existing tools and infrastructure
    * Requires that masters be updated and has limits on dynamic behavior
3. Direct external call outs for object mutation (RPC to initialize objects)
    * Requires a new programming model
    * Duplicates our create - watch - update logic from controllers
4. Make it easy to recompile Kubernetes to have new admission controllers
    * Limits administrators to using Go
    * Prevents easy installation and dynamic reconfiguration


<!-- BEGIN MUNGE: GENERATED_ANALYTICS -->
[![Analytics](https://kubernetes-site.appspot.com/UA-36037335-10/GitHub/docs/proposals/admission_control_extension.md?pixel)]()
<!-- END MUNGE: GENERATED_ANALYTICS -->
