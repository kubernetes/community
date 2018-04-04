# PodRestrictions

## Overview

This is a proposal for a new policy object (and associated admission controller) that adds basic
restrictions to most fields on a Pod. This is motivated by a desire to unify the various policies we
have now (which all behave in slightly different ways), and to unblock expanding the set of
restricted fields. Furthermore, it aims to solve some of the key usability problems with
PodSecurityPolicies, such as the inability to compose policies and the difficulty in binding
policies.

### Existing Policies

There are multiple policies that constrain pod objects today. PodRestriction would supersede these
policies, and they would eventually be deprecated.

[AlwaysPullImages](https://kubernetes.io/docs/admin/admission-controllers/#alwayspullimages) -
Admission controller that sets the PullAlways ImagePullPolicy to be set on every container.

[PodSecurityPolicy](https://kubernetes.io/docs/concepts/policy/pod-security-policy/) - Restricts
PodSecurityContext, SecurityContext, and some volume options. Bound to users and service accounts.

[LimitRanger](https://kubernetes.io/docs/admin/admission-controllers/#limitranger) - Limits the
range of resource limits that can be placed on a Pod. Intersection of LimitRanges in a namespace.

[Toleration](https://kubernetes.io/docs/admin/admission-controllers/#podtolerationrestriction) &
[NodeSelector](https://kubernetes.io/docs/admin/admission-controllers/#podnodeselector)
restrictions - Restrict scheduling options through namespace annotations. Several proposals are out
for new policies to supercede these ([#1937](https://github.com/kubernetes/community/pull/1937),
[#48041](https://github.com/kubernetes/kubernetes/issues/48041),
[#61185](https://github.com/kubernetes/kubernetes/issues/61185)).

There have been other proposals for policies restricting fields like labels, and annotations.

## Scope

This proposal is about building a new policy type that can place restrictions on any or all pod
fields (see [Monolith vs. Modular](#monolith-vs-modular)). However, the goal is not to combine all
admission controllers into a single one. The following controls are explicitly out of scope:

-   Defaulting (and mutations in general) - See [No Defaulting](#no-defaulting).

-   External dependencies - The policy doesn't express any inter-object relationships, and doesn't
    call out to external servers or have plugin points.

-   Complex logic or interfield relationships - The goal is to keep the policy simple, both to
    implement and reason about. The only exception is fields that explicitly reference others like
    VolumeMounts and EnvVarSource, but most fields should be able to be validated without knowledge
    of any other Pod fields.

## Design

### Policy Binding

PodRestrictions are applied at the namespace level. Any Pod (or PodTemplate object) must conform to
the policies matching the namespace it is created in.

```go
Binding:
    Namespaces        []string
    NamespaceSelector *metav1.LabelSelector
```

#### Justification

There are a few different approaches taken by our current policies (see
[#60001](https://github.com/kubernetes/kubernetes/issues/60001)).

**Bind to Users**: Binding policy to users is the obvious first choice, but has a number of
problems. First, users don't typically create pods, they create resources like Deployments that
create pods on their behalf. That "on behalf" ends up being problematic to implement, especially
with issues like confused deputy attacks, third party controllers, and changing policies.

**Bind to ServiceAccounts**: To get around the user problem, PodSecurityPolicy adds the ability to
bind to the target objects service account. ServiceAccount usage isn't actually authorized
separately from namespace usage, so this just ends up being namespace-level authorization. However,
because PodSecurityPolicy is authorized by either the source or target user, this ends up being less
secure, since the Pod can use it's service account to create pods in other namespaces with the same
access level. For instance, using PodSecurityPolicy you cannot have a controller that runs as
privileged without opening up all pods it creates to running as privileged.

**Namespaced Objects**: This approach has its advantages, but it doesn't scale well to large
organizations with many namespaces (without a system on top managing all the policies across
namespaces). See extensions below for how we might use namespaced objects with PodRestrictions.

**Bind to Pod labels**: Pod labels don't have any authorization of usage without something like
PodRestriction, so binding to labels isn't sufficient no it's own.

**Bind to Namespaces**: The approach I'm recommending. Cluster level objects use a namespace
selector to select the namespaces it applies to. It is consistent with the authorization of create
requests, which are authorized at the namespace level. A small number of policies can be managed at
the cluster level, and mapped to the namespaces they should apply to.

#### Optional Extensions

-   Namespaced PodRestrictions - PodRestriction objects can be created in a namespace, so that
    namespace admins can place additional restrictions within their namespace.

-   Additional Bindings - The NamespaceSelector is not a top-level field so that additional binding
    approaches may be added in the future. Examples may include: ServiceAccounts, LabelSelector,
    NodeSelector, etc.

### Policy Matching - Union or Intersection?

**Intersection** refers to the model where a Pod must be valid under all PodRestrictions that
pertain to it. The advantage of intersecting is that restrictions can be composed. For example, I
could have a policy that frontend & backend Pods cannot share a node, and a separate policy that all
production Pods must drop privileges. Without the ability to intersect those policies, I need 3
different versions (either, or, both). A problem with this approach is that running a privileged
namespace (e.g. kube-system) means that every restriction needs an exception in the
NamespaceSelector.

**Union** refers to the model where a Pod must be valid under any PodRestriction that is bound to
it. This approach makes it easier to reason about what the restrictions on a given namespace are,
and easier to add exceptions for namespaces, but can lead to a combinatorial explosion of different
policies that can become unmanageable.

**Proposal**: Why not both? The default behavior is intersection, but you can punch a hole through
for a specific namespace with a union approach. An optional `BindingMode` is added to the Binding
which specifies the behavior. A pod must be valid under either:

-   ANY `BindingModeAccept` policy that is bound to it, or

-   ALL `BindingModeDrop` policies that are bound to it

In other words, the Pod is first checked against Accept restrictions, and is accepted if any
match. Then, it is checked against the Drop restrictions, and rejected if any don't match. The
default behavior at the end is to accept the pod.

```
Binding:
    Mode BindingMode

BindingModeAccept
BindingModeDrop
```

### Admission Handling

The PodRestriction controller handles all create, update, and patch requests for Pod resources. It
is also applied to all (native) resources that create pods via a PodTemplate, including ReplicaSets,
ReplicationControllers, DaemonSets, StatefulSets, Deployments, Jobs, etc.

#### Justification

From a security perspective, it should be sufficient to apply the restriction to Pod requests, so
why validate other objects? Fail-fast. By surfacing an error before the controller object is
created, the user gets immediate feedback that they are attempting a forbidden operation. The
alternative experience is that the controller object creation succeeds, but then the user must
follow debug steps to determine why Pods are not being created.

#### Optional Extensions

New CRDs that hold a PodTemplate and create Pods should be able to be registered against the
PodRestriction controller. The registration should include a path (auto-discovery?) to extract the
PodTemplate from the CRD.

### API Design

#### Monolith vs. Modular

Our current approach to policies has been "modular", i.e. a different controller and API for each
"area" of features. This has led to some inconsistencies and gaps in our policies, and makes adding
new restrictions needlessly complex. I propose we take a "monolithic" approach, with a single policy
(PodRestriction) for all Pod fields.

A monolithic restriction policy (i.e. single resource type) also allows the user to express more
complex inter-relationships, such as:

-   Pods using a certain volume type must run with reduced privileges

-   Pods can only run as privileged if they run with a specific AntiAffinity

-   Pods labeled as staging must not run on production nodes.

#### No Defaulting

There are 2 problems with mutating Pod fields: 1) it becomes hard to compose policies (what happens
when the defaults conflict?) and 2) you need a policy ordering mechanism for when a Pod matches
multiple policies. Instead, PodRestriction opts out of mutating defaults entirely, and leaves that
for a different component to solve ([PodPreset?](../service-catalog/pod-preset.md)).

## API Sketch

I don't want to hash out all the details of the API (such as which fields are constrained, and what
the naming is) in this proposal, but there are a couple high level design goals and takeaways:

1.  The PodRestriction API layout mostly mirrors that of the Pod object. This makes the API more
    predictable, and easier to manage (i.e. I want to restrict Pod field X, I know how to do that in
    the PodRestriction).

2.  Restrictions are self-contained. There isn't complex interdependent logic here. If you want more
    than basic whitelist / blacklist / pattern matching, then implement a new policy or controller
    for that.

3.  There are a lot of fields in the PodSpec, and thus this API gets large (see [Monolith
    vs. Modular](#monolith-vs-modular)). We don't need to build it all at once, but I think over
    time it will eventually grow into the full version.

4.  Fields are unrestricted by default. This is important due to the size of the API. If I want to
    restrict 1 field, I shouldn't need to write 500 lines to open up all the other fields.

```go
// This API sketch is partial, and only intended to indicate the the flavor of
// the API, not to completely specify it.

type PodRestriction struct {
    metav1.TypeMeta
    metav1.ObjectMeta

    Binding *PodRestrictionBinding // Default to cluster-wide

    Spec PodRestrictionSpec
}

type PodRestrictionBinding struct {
    Mode              BindingMode  // Default = Drop

    Namespaces        []string
    NamespaceSelector *meta.LabelSelector
}

type BindingMode string

const (
    BindingModeAccept BindingMode = "Accept"
    BindingModeDrop   BindingMode = "Drop"
)

type PodRestrictionSpec struct {
    Metadata *ObjectMetaRestriction

    Spec *PodSpecRestriction
}

type ObjectMetaRestriction struct {
	Name        *StringRestriction
	Labels      *StringMapRestriction
	Annotations *StringMapRestriction
}

type PodSpecRestriction struct {
	Volumes []VolumeRestriction

	InitContainers []ContainerRestriction // Allow any; default to same as containers?
	Containers     []ContainerRestriction // Allow any

	RestartPolicy                 *StringRestriction
	TerminationGracePeriodSeconds *NumericRestriction
	ActiveDeadlineSeconds         *NumericRestriction
	DNSPolicy                     *StringRestriction
	NodeSelector                  *StringMapRestriction
	AutomountServiceAccountToken  *BoolRestriction
	SecurityContext               *PodSecurityContextRestriction

    // TODO: Scheduling restrictions are deeply nested structures and require
    // more thought.
	// Affinity      *AffinityRestriction
	// SchedulerName *StringRestriction
	// Tolerations   []TolerationRestriction

    // ...
}

type ContainerRestriction struct {
	Image []ImageRestriction  // Allow Any

	ImagePullPolicy          *StringRestriction
	SecurityContext          *SecurityContextRestriction

    // ...
}

// QUESTION: Are image restrictions injecting too much login?
type ImageRestriction struct {
	Registry  *StringRestriction
	Image     *StringRestriction
	Tag       *StringRestriction
}

type SecurityContextRestriction struct {
	Capabilities             *CapabilitiesRestriction
	Privileged               *BoolRestriction
	SELinuxOptions           *SELinuxOptionsRestriction
	RunAsUser                *NumericRestriction
	RunAsGroup               *NumericRestriction
	RunAsNonRoot             *BoolRestriction
	ReadOnlyRootFilesystem   *BoolRestriction
	AllowPrivilegeEscalation *BoolRestriction
}

type SELinuxOptionsRestriction struct {
	PointerRestriction
	User  *StringRestriction
	Role  *StringRestriction
	Type  *StringRestriction
	Level *StringRestriction
}

type CapabilitiesRestriction struct {
	PointerRestriction
	Add  *StringListRestriction
	Drop *StringListRestriction
}

type PodSecurityContextRestriction struct {
	HostNetwork           *BoolRestriction
	HostPid               *BoolRestriction
	HostIPC               *BoolRestriction
	ShareProcessNamespace *BoolRestriction

    // RunAs* options left to container restrictions.
	SupplementalGroups    *NumericRestriction
	FSGroup               *NumericRestriction
}

type VolumeRestriction struct {
    CollectionRestriction

    // Mount restrictions apply to the volumes included in this restriction
	Mount *MountRestriction

	VolumeType *StringListRestriction  // Allowed volume types

    // Type-specific restrictions
	HostPath *HostPathVolumeRestriction
    // ...
}

type MountRestriction struct {
	ReadOnly         *BoolRestriction
	Mountpropagation *StringRestriction
}

type HostPathVolumeRestriction struct {
	Path *StringRestriction // Prefix match whitelist / blacklist
	Type *StringRestriction
}

type PointerRestriction struct {
	ForbidNil  bool
	RequireNil bool
}

type CollectionRestriction struct {
	ForbidEmpty  bool
	RequireEmpty bool
}

type StringRestriction struct {
	PointerRestriction

	Whitelist []string
	Blacklist []string
	Regex     string
	// Default string
}

// This type needs more thought.
type StringMapRestriction struct {
	CollectionRestriction
	KeyBlacklist      []string
    KeyWhitelist      []string
    RequiredKeys      []string
	Values            map[string]StringRestriction
}

type StringListRestriction struct {
    CollectionRestriction
    Values          *StringRestriction
    RequiredValues  []string
}

type NumericRestriction struct {
	PointerRestriction
	Ranges []NumericRange // Allow Any
}

type NumericRange struct {
	// Allow all values in range, inclusive.
	// Unset means open on that side, e.g.
	// {Max: 10} allows all values <= 10.
	Min, Max *int64
}

type BoolRestriction struct {
	PointerRestriction
	Require *bool
}
```

## Examples

A [restricted PodSecurityPolicy](https://kubernetes.io/docs/concepts/policy/pod-security-policy/#example-policies) translated into a PodRestriction:

```yaml
apiVersion: policy/v1alpha1
kind: PodRestriction
metadata:
  name: restricted
# No binding == cluster-wide
spec:
  metadata:
    annotations:
      requireKeys:
        - seccomp.security.alpha.kubernetes.io/pod
      valueRestrictions:
        apparmor.security.beta.kubernetes.io/pod:
          whitelist: ['runtime/default']
        seccomp.security.alpha.kubernetes.io/pod:
          whitelist: ['docker/default']
  spec:
    containers:
    - securityContext:
        privileged:
          allowNil: true
          require:  false
        allowPrivilegeEscalation:
          require:  false
        capabilities:
          add:
            requireEmpty: true
          drop:
            requiredValues: ['ALL']
        runAsUser:
          allowNil: false
          ranges:
            - min: 1
        runAsGroup:
          allowNil: false
          ranges:
            - min: 1
    securityContext:
      hostNetwork:
        require: false
      hostIPC:
        require: false
      hostPID:
        require: false
      supplementalGroups:
        ranges:
          - min: 1
      fsGroup:
        ranges:
          - min: 1
    volumes:
      volumeType:
        whitelist:
          - 'configMap'
          - 'emptyDir'
          - 'projected'
          - 'secret'
          - 'downwardAPI'
          - 'persistentVolumeClaim'
```

Allowing all (and only) addons in kube-system:

```yaml
apiVersion: policy/v1alpha1
kind: PodRestriction
metadata:
  name: kube-system
binding:
  mode: 'Allow'
  namespaces: ['kube-system']
spec:
  metadata:
    labels:
      requireKeys:
        - k8s-app
        - addonmanager.kubernetes.io/mode: Reconcile
      valueRestrictions:
        addonmanager.kubernetes.io/mode
          whitelist: ['Reconcile', 'EnsureExists']
  # No restrictions on PodSpec
```
