---
kep-number: draft-20180426
title: Namespace Population
authors:
  - "@easeway"
owning-sig: sig-auth
participating-sigs:
  - sig-api-machinery
  - sig-cluster-lifecycle
reviewers:
  - "@davidopp"
  - "@tallclair"
  - "@ericchiang"
  - "@liggitt"
  - "@roberthbailey"  
approvers:
  - "@tallclair"
  - "@ericchiang"
  - "@liggitt"
  - "@roberthbailey"
editor: "@easeway"
creation-date: 2018-04-26
status: provisional
---

# Namespace Population

## Table of Contents

* [Table of Contents](#table-of-contents)
* [Summary](#summary)
* [Motivation](#motivation)
    * [Goals](#goals)
    * [Non-Goals](#non-goals)
* [Proposal](#proposal)
    * [User Stories](#user-stories)
      * [Security Defaults](#security-defaults)
    * [Implementation Details](#implementation-details)
      * [Namespace Creation Flow](#namespace-creation-flow)
      * [The Controller](#the-controller)
      * [Namespace Match](#namespace-match)
      * [Apply the template](#apply-the-template)
      * [Schema validation](#schema-validation)
      * [Opt-out](#opt-out)
      * [Namespace initialization](#namespace-initialization)
      * [Extending Initializers](#extending-initializers)
      * [Manual Update of Objects](#manual-update-of-objects)
* [Graduation Criteria](#graduation-criteria)

## Summary

Namespace Population is an automated mechanism to make sure the predefined policy objects
(e.g. NetworkPolicy, Role, RoleBinding) are present in selected namespaces.

## Motivation

It's a common problem that Kubernetes cluster administrators want to apply pre-defined
policy objects into namespaces, for example, setting up default ResourceQuota,
binding `default` service account with a PodSecurityPolicy,
restricting network ingress/egress using a NetworkPolicy etc.

Today, people can achieve that in a few ways:

1. Use scripts or automation tools (e.g. Terraform) to create a namespace and then populate
  these policy objects;
2. Adopt a namespace request process: users request namespaces and cluster admins create for
  them with all policy objects populated;
3. Build some in-house services/APIs that users can use to create a namespace by themselves.

The approach `#1` is very basic,
and assumes all users will collaboratively use these scripts in a good way;
`#2` introduces more controls over `#1` at the cost of more lengthy process;
`#3` works pretty automated, but requires significant development effort,
and result in a non-standard way that users must learn before they can use.

This proposal demonstrates a solution that achieves the same result with standard approach
that the users already know:

```
kubectl create namespace my-namespace
```

The above command will create the namespace and also populated all policy objects pre-defined
by cluster administrators.

Please read forward for more details in [User Stories](#user-stories) section.

### Goals

- The namespace population mechanism described in this proposal is effective to
populate identical configurations into large number of namespaces;
- It's also effective to populate different sets of predefined configurations into
different sets of namespaces;
- The mechanism is agnostic to object types as long as they are namespace scoped;
- The mechanism works in a single cluster;
- This proposal doesn't require any changes in core Kubernetes.

#### MVP Scope

- `NamespaceTemplate` is applied into newly created namespaces; the existing namespaces are left as is;
- All policy objects are defined inline in `NamespaceTemplate` custom resource, no external sources of policies supported;
- Not support namespace matching - apply to all newly created namespaces;
- NO opt-out mechanism;
- Not watching changes to policy objects already created.

### Non-Goals

- The namespace population mechanism is NOT effective for customizing namespaces individually;
- Limit the number of namespaces to be created by users/groups (defend against namespace exhaustion or malicious namespace creators);
- Conflicts of namespace names;
- No dynamic objects (generating the object definition during population,
e.g. string substitution) is supported - all objects are defined statically, 
except `$(CREATOR)` and `$(NAMESPACE)` (explained in [Apply the template](#apply-the-template));
- No validation is performed over static object definitions, until they are applied;
- The mechanism does NOT work across clusters.
- This proposal relies on initializers mechanism in core Kubernetes, 
  however, the deprecation/change of initializers in core Kubernetes is out-of-scope and should be discussed in a separate proposal,
  as this proposal doesn't require any core Kubernetes changes.

## Proposal

The namespace population is performed by defining one or more cluster-scope
Custom Resource with kind `NamespaceTemplate`,
and deploy a controller in the cluster.

The Custom Resource Definition of `NamespaceTemplate`:

```yaml
apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  name: namespacetemplates.policy
spec:
  group: policy
  version: v1alpha1
  scope: Cluster
  names:
    plural: namespacetemplates
    singular: namespacetemplate
    kind: NamespaceTemplate
    shortNames:
    - nstpl
```

The example of `NamespaceTemplate`:

```yaml
apiVersion: policy/v1alpha1
kind: NamespaceTemplate
metadata:
  name: default
spec:
  namespaces:
    # labelSelector selects namespaces by labels to populate defined objects
    labelSelector: {}
  templates:
    # list of objects to be populated into namespaces defined here
    - apiVersion: networking.k8s.io/v1
      kind: NetworkPolicy
      metadata:
        name: default
      spec:
        podSelector: {}
        policyTypes: ["Ingress", "Egress"]
    - apiVersion: rbac.authorization.k8s.io/v1
      kind: RoleBinding
      metadata:
        name: use-podsecuritypolicy
      roleRef:
        apiGroup: rbac.authorization.k8s.io
        kind: ClusterRole
        name: use-psp-default
      subjects:
        - apiGroup: rbac.authorization.k8s.io
          kind: Group
          name: system:serviceaccounts
    - apiVersion: rbac.authorization.k8s.io/v1
      kind: RoleBinding
      metadata:
        name: creator
      roleRef:
        apiGroup: rbac.authorization.k8s.io
        kind: ClusterRole
        name: admin
      subjects:
        - apiGroup: rbac.authorization.k8s.io
          kind: User
          name: '$(CREATOR)'
```

There's a size limit of any object created inside Kubernetes (due to etcd).
If there are too many policy objects inlined in the template, the creation of `NamespaceTemplate` will fail.
In practice, it's uncommon to hit the limit.
If that's the case, it's recommended the cluster admin redesign the `NamespaceTemplate` into multiple objects
and separate policy objects into different tiers of groups.

Alternatively, the policy objects can be defined in some other places, for example:

- As real policy objects in a real namespace which serves as a template only.
  This approach has advantages that all policy objects are handled through Kubernetes APIs with all supported features, like schema validation.
  And is more intuitive for cluster admin to understand.
  Easier to control the access via RBAC.
  However it's difficult to dynamically create a policy object (e.g. RoleBinding for the creator) according to
  the creator and namespace name.
- Remote file server, git etc. and use some tools like kinflator to compose on the fly
  Out-of-scope in this proposal.

The `NamespaceTemplate` is potentially extensive to support arbitrary source of policy objects,
by introduce a property `source` to define how to locate the policy objects.
This is out-of-scope in this proposal.

### User Stories

#### Security Defaults

The cluster administrator manages a shared Kubernetes clusters and allows
developers self-service to create namespaces by themselves, with predefined
NetworkPolicy, RoleBinding using PodSecurityPolicy, ResourceQuota,
ServiceAccounts and Roles present in the namespace, as security defaults.

The cluster administrator creates `NamespaceTemplate` with policy objects in `templates`.
With the help of `NamespaceTemplate` enforcement controller, these objects are
created automatically when a user creates a new namespace.

### Implementation Details

#### Namespace Creation Flow

The overall namespace creation flow with `NamespaceTemplate` is:

1. A namespace creator creates a namespace, e.g. with CLI `kubectl create namespace ns1`
2. The _MutatingAdmissionWebhook_ intercepts the creation request, and
   - add label `authorization.k8s.io/creator=username`
   - add `namespace-template` to `metadata/initializers.pending` list
3. The `NamespaceTemplate` controller is notified of creation of namespace, it
   - find all matching `NamespaceTemplate` objects by labels and sort them by name
   - Concate policy objects defined in these objects in the order
   - Create these policy objects into the namespace
   - Remove `namespace-template` from `metadata/initializers.pending` list
4. The client gets the response that the namespace is created

#### The Controller

A controller is deployed in the cluster (recommended in the `kube-system` namespace).
It runs with a service account with privilege to

- get/list/watch Namespace objects
- get/list/watch NamespaceTemplate objects
- get/list/create/update/patch/delete namespace-scope objects in all namespaces, except `kube-system`

The controller watches Namespace and NamespaceTemplate objects,
when a namespace is added/updated, it applies the objects in namespace templates matching the namespace;
when a namespace template is added/updated, it applies objects in the template to all matching namespaces.

#### Namespace Match

The `NamespaceTemplate` uses `labelSelector` to efficiently select namespaces.
This design is not proposing extra selection mechanism, 
even though it's almost impossible to select/deselect a namespace without labels.
This problem is out-of-scope here.

When multiple `NamespaceTemplate` objects are matched, the items in the templates are concated
according to the order of names of `NamespaceTemplate` objects (ascend).
It's cluster admin's responsibility to design the policy objects carefully into multiple `NamespaceTemplate` objects,
the population mechanism is not aware of any conflicts introduced by the policy objects.

Using `labelSelector` imposes a potential security concern that 
a malicious namespace creator creates a namespace with carefully selected labels to pick the  `NamespaceTemplate` which is different from cluster admin's intention.
And it's also possible the owner of namespace later updates the labels to matching a different set of
`NamespaceTemplate` from those at creation time.

To mitigate the problem, the RBAC system can be utilized to define the permission that a namespace creator must have to use certain `NamespaceTemplate`.
The mechanism introduces the `use` verb in RBAC ClusterRole to express the permission to access a `NamespaceTemplate`. (Similar to PodSecurityPolicy)

Here's an example:

```yaml
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: use-namespace-template-default
rules:
  - apiGroups: ["poliy"]
    resources: ["namespacetemplates"]
    resourceNames: ["default"]
    verbs: ["use"]
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: namespace-template-usage
roleRef:
  kind: ClusterRole
  name: use-namespace-template-default
  apiGroup: rbac.authorization.k8s.io
subjects:
- kind: User
  name: "user1"
  apiGroup: rbac.authorization.k8s.io
```

The namespace creation flow becomes:

1. A namespace creator creates a namespace, e.g. with CLI `kubectl create namespace ns1` with or without certain labels
2. The _MutatingAdmissionWebhook_ intercepts the creation request, and
   - add label `authorization.k8s.io/creator=username`
   - add `namespace-template` to `metadata/initializers.pending` list
3. The _ValidatingAdmissionWebhook_ 
   - find out all matched `NamespaceTemplate` objects
   - verify if the creator is allowed to apply **ALL** of the matched `NamespaceTemplate`. It fails the request if the creator doesn't have `use` permission on **ANY** of the mached `NamespaceTemplate`.
4. Same as the original flow

Using this mechanism, the cluster admin is able to design `NamespaceTemplate` intended to be used by a limited number of users.

For example

```yaml
apiVersion: policy/v1alpha1
kind: NamespaceTemplate
metadata:
  name: nonprivileged
spec:
  namespaces:
    # matches namespaces without being labeled as privileged
    labelSelector:
      matchExpressions:
      - key: namespace-class
        operator: NotIn
        values: [privileged]
  templates: [] # ... details omitted
---
apiVersion: policy/v1alpha1
kind: NamespaceTemplate
metadata:
  name: privileged
  annotations:
    policy/namespace-template/require-permission: true
spec:
  namespaces:
    # matches namespaces being labeled as privileged
    labelSelector:
      matchLabels:
        namespace-class: privileged
  templates: [] # ... details omitted
```

To create a namespace allowing privileged workloads, 
the creator must put `namespace-class=privileged` in labels of namespace.
To limit the users who are allowed to create _privileged_ namespaces, 
the cluster admin sets additional RBAC rules:

```yaml
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: privileged-namespace-creator
rules:
  - apiGroups: ["policy"]
    resources: ["namespacetemplates"]
    resourceNames: ["privileged"]
    verbs: ["use"]
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: privileged-namespace-creator
roleRef:
  kind: ClusterRole
  name: privileged-namespace-creator
  apiGroup: rbac.authorization.k8s.io
subjects:
- kind: User
  name: "user1"
  apiGroup: rbac.authorization.k8s.io
```

Also note, the annotation `policy/namespace-template/require-permission=true` on a `NamespaceTemplate`
require permission from authorizer (RBAC).
Without this annotation, the `NamespaceTemplate` can be used by anyone, thus serving as `default` template.

It's not a risk if a namespace creator creates a namespace with labels not matching any `NamespaceTemplate`,
as the creator won't get further permission inside the namespace if the RBAC is properly setup.

Alternatively, the _ValidatingAdmissionWebhook_ may not be necessary.
As a result, a namespace can always be created, 
while the controller populating the namespace will find out a certain `NamespaceTemplate` is not accessible to the creator.
It may simply skip all policy objects in that `NamespaceTemplate`.
This approach is not preferred, as it introduces some level of implicitness that is not obvious to users when they make mistakes.

#### Apply the template

The list of `templates` in `NamespaceTemplate` is extracted and concatenated as multi-document YAML
(separated by `---` line between two objects).
There's no limit of the number of items defined in `templates` but the total size of the object.
As long as the object can be created in the cluster (it's out-of-scope how the object is created), 
it can be handled by the controller.

The additional label is injected

```
policy/namespace-template-name=<name of NamespaceTemplate>
```

To indicate which namespace template creates this object.

Then `kubectl apply` is used to apply the template.
The reason to use `kubectl` is that the apply logic (3-way merge) is very complicated,
and currently performed by `kubectl` (not in API server).
Once the logic is available in API server, the APIs will be used directly.
`--prune` options is used together with `-l` to clean up old objects no longer defined when namespace template changes.

When multiple templates matches the same namespace,
the templates are concatenated in the alphabetic order of the name of `NamespaceTemplate`.

A `NamespaceTemplate` can be disabled by an explicit annotation:

```
policy/namespace-template-apply=disable
```

Inside `templates`, two special tokens will be substituted:

- `$(CREATOR)`: the name of the user created the namespace
- `$(NAMESPACE)`: the name of current namespace

Chose of syntax `$()` is based on considerations to avoid conflict with environment variable
substitutions `${}` which is highly possible in templated objects.

These tokens are only allowed in field values.
A more strict implementation could also check the templated resource type, 
i.e. only allow `$(CREATOR)` to be used inside RBAC RoleBinding.
A simpler implementation treats the whole YAML as a string and perform substitution from there,
and leaves the checks and validations to `kubectl apply` process.
This doesn't introduce security concerns because the value of these two tokens are under control.

#### Schema validation

As `NamespaceTemplate` is defined as CRD (Custom Resource Definition),
currently there's no effective way to validate the schema of CR (Custom Resource)
if it contains complex types, or encapsulating existing object types.
This problem is out-of-scope here.
As a result, the `templates` are not validated during the creation of `NamespaceTemplate`,
and it will fail at population time (via `kubectl apply`) if anything is written incorrectly.
The failure will be reported as Kubernetes Events in the target namespace.

An alternative is to create a template namespace,
and put all these objects into the namespace as real resources.
When a new namespace is created, the objects in the template namespace are copied over.
As real objects are created, they are effective, and may cause side effects,
though within known object types, there's no impact of the effectiveness of these objects.

There are other alternatives that the objects can be defined somewhere else from `NamespaceTemplate`.
The implementation can be pluggable (e.g. adding a `source` property) to support different sources of templates.
While the initial proposal will focus on inlined objects.

The MVP scope covers the inlining templates only.

#### Opt-out

Cluster administrator is able to opt-out some namespaces from being populated by the controller.
`kube-system` is always opt-out.
For other namespaces, there are two options to opt-out from automated population:

1. Put a label that is excluded in the `labelSelector` of the `NamespaceTemplate`
2. Annotate the namespace with `policy/namespace-template-opt-out=true`

Note: the objects previously populated by the controller will be left as-is in the namespace when it's opted out.

Normally, there's no _exclusive_ condition in `labelSelector` of the `NamespaceTemplate`, so option 2 is recommended.

Some ephemeral namespaces created by privileged controllers (e.g. CI/CD pipelines) may need to be opt-out
as they are not accessibly by users and will be managed completely by the controllers.
As these controllers don't have knowledge about `NamespaceTemplate`, 
or it's almost impossible to customize the way how they create namespaces,
an automated opt-out mechanism is required.
The mechanism is based on the service account running the controllers.
With a `MutatingAdmissionWebhook`, 
a namespace created by these controllers will be automatically annotated for opt-out.

_Opt-out_ can be possibly abused by a user who puts the annotation when creates a namespace.
The impact is minimal as with correct RBAC rules, 
the user won't get further permissions inside the namespace.
However the situation can be handled in a better way with a `ValidatingAdmissionWebhook` checking
a specific RBAC binding:

```yaml
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: namespace-template-optout
rules:
  - apiGroups: ["policy"]
    resources: ["namespacetemplates"]
    verbs: ["optout"]
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: namespace-template-optout
roleRef:
  kind: ClusterRole
  name: namespace-template-optout
  apiGroup: rbac.authorization.k8s.io
subjects:
- kind: ServiceAccount
  name: default
  namespace: 'cicd-namespace'
  apiGroup: rbac.authorization.k8s.io
```

The webhook will fail the request if the opt-out annotation is present without the RBAC binding for the requestor.

_Note: This mechanism is out of MVP scope._

#### Namespace initialization

When a namespace is created, 
the `NamespaceTemplate` enforcement controller is populating the namespace using policy objects
defined in `templates`.
During this period, no other users should be able to create objects into the namespace.
Once the namespace is populated, 
the creator of the namespace will gain permission (as a RoleBinding defined in `templates`) to 
create further objects into the namespace.

There are a few issues in this process:

1. The user creates the namespace should not be able to access the namespace before it's fully populated;
2. The `NamespaceTemplate` enforcement controller has no context about who created the namespace, thus
have problem fill in the value for `$(CREATOR)`;
3. The old clients implementation may break as they see the namespace and move forward to create objects but get forbidden error.

Issue `#1` can be solved with properly configured RBAC rules: 
an extra `ClusterRole` is created to allow _create_ verb of namespaces:

```yaml
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: namespace-creator
rules:
  - apiGroups: [""]
    resources: ["namespaces"]
    verbs: ["create"]
```

After the namespace is created, 
the creator should be granted further permissions to manage the namespace,
with the following objects included in the `templates` of `NamespaceTemplate`:

```yaml
apiVersion: rbac.authorization.k8s.io/v1
kind: Role
metadata:
  name: namespace-owner
rules:
  - apiGroups: [""]
    resources: ["namespaces"]
    resourceNames: ["$(NAMESPACE)"]
    verbs: ["get", "list", "watch", "update", "patch", "delete"]
---
apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  name: admin
roleRef:
  kind: ClusterRole
  name: admin
  apiGroup: rbac.authorization.k8s.io
subjects:
- kind: User
  name: "$(CREATOR)"
  apiGroup: rbac.authorization.k8s.io
---
apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  name: namespace-owner
roleRef:
  kind: Role
  name: namespace-owner
  apiGroup: rbac.authorization.k8s.io
subjects:
- kind: User
  name: "$(CREATOR)"
  apiGroup: rbac.authorization.k8s.io
```

Issue `#2` can be solved using a `MutatingAdmissionWebhook` which adds an annotation

```
authorization.k8s.io/creator=name
```

during the API request to create a namespace.
Then the controller is able to fill `$(CREATOR)` from that annoation.

This annotation must be made immutable.
The `MutatingAdmissionWebhook` can easily achieve this in further _update/patch_ requests.
Alternatively, a `ValidatingAdmissionWebhook` can be used to reject requests which attempts to change the annotation,
however it may break existing client implementations which don't expect immutable annotations as they don't exist in Kubernetes today.

Issue `#3` can be solved by using a `MutatingAdmissionWebhook` to adding `NamespaceTemplate` enforcement controller into `metadata.initializers` during the creation of the namespace.
By default, the namespace creator will hold the response for a period of time until the namespace is fully initialized,
(ref: `WaitForInitialized` in 
[k8s.io/apiserver/pkg/registry/generic/registry/store.go](https://github.com/kubernetes/kubernetes/blob/master/staging/src/k8s.io/apiserver/pkg/registry/generic/registry/store.go))
unless `includeUninitialized` is set to `true` in the request from the client.

Note: this proposal relies on initializers mechanism in core Kubernetes. The implementation will have to change if the generic initializers mechanism is deprecated and move to be namespace specific (using `Namespace.spec.initializers` instead of `metadata.initializers`), though the change will be trivial.

For further protection,
to prevent other privileged users/service accounts to create objects into an uninitialized namespace,
A `ValidatingAdmissionWebhook` can be used to validate with authorization interface if the requestor
has the `initialize` permission on the current namespace if it's uninitialized:

```yaml
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: namespace-initialization
rules:
  - apiGroups: [""]
    resources: ["namespaces"]
    verbs: ["initialize"]
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: namespace-initializer
roleRef:
  kind: ClusterRole
  name: namespace-initialization
  apiGroup: rbac.authorization.k8s.io
subjects:
- kind: ServiceAccount
  name: "system:serviceaccount:kube-system:name"
  apiGroup: rbac.authorization.k8s.io
```

#### Extending Initializers

Kubernetes provides resource initialization mechanism
(injecting items into `metadata/initializers.pending` list) in a _MutatingAdmissionWebhook_ to
allow controllers to finish initialization work before the creation request is completed and response is 
delivered to the client.

In this proposal, the _MutatingAdmissionWebhook_ leverages this mechanism by adding `namespace-template` into
`metadata/initializers.pending` list.
This can be effectively extended to allow pluggable controllers to perform additional initialization tasks.
The same approach as Initializers admission controller (which is being deprecated) is proposed,
by creating a Custom Resource `NamespaceInitializationConfiguration` and list the names of initializers there.
The generic implementation of _Initializers_ can be moved here and specifialized for namespaces.

There's a scalability issue with current initialization mechanism in core Kubernetes.
To maximize the compatibility, namespaces are not allowed additional _phase_ beyond `active` and `terminating`,
and client is assuming synchronous behavior of the creation API.
So the initialization mechanism holds up the request for a short period of time (e.g. 60s).
That causes scalability issue, in the case there are large number of namespaces, and the initialization task is non-trivil.
This is a known issue and the proposal is not going to solve it.

#### Manual Update of objects

The objects created by namespace population may be altered manually or through
other ways.
The recommended way is setting up RBAC properly to mitigate the risk (not allowing users to update these objects).
The proposal doesn't watch update of objects already created.
The update may be reverted in the case triggers template re-apply (namespace change or template change).

This topic is not in MVP scope.

## Graduation Criteria

### Alpha

- Namespace can be populated without break of existing clients
- Creator of namespace is automatically granted permissions in the namespace
- Uninitialized namespace is only accessible to service accounts that have permission for initialization

### Beta

- This feature is used in real clusters

### GA

- User demands enabling this feature in production clusters
