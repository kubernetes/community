# Scheduling Policy

_Status: Draft_
_Author: @arnaudmz, @yastij_
_Reviewers: @bsalamat, @liggitt_

# Objectives

-  Define the concept of scheduling policies
-  Propose their initial design and scope

## Non-Goals

-  How taints / tolerations work
-  How NodeSelector works
-  How node / pod affinity / anti-affinity rules work
-  How several schedulers can be used within a single cluster
-  How priority classes work

# Background

During real-life Kubernetes architecting we encountered contexts where role-isolation (between administration and simple namespace usage in a multi-tenant context) could be improved. So far, no restriction is possible on toleration, priority class usage, nodeSelector, anti-affinity depending on user permissions (RBAC).

Identified use-cases aim to ensure that administrators have a way to restrict users or namepace when
-  using schedulers,
-  placing pods on specific nodes (master roles for instance),
-  using specific priority classes,
-  expressing pod affinity or anti-affinity rules.

# Overview

Implementing SchedulingPolicy implies:
-  Creating a new resource named **SchedulingPolicy** (schedpol)
-  Creating an **AdmissionController** that dehaves on a deny-all-but basis
-  Allow SchedulingPolicy to be used by pods using RoleBindings or ClusterRoleBindings

# Detailed Design

SchedulingPolicy resources are supposed to apply in a deny-all-except approach. They are designed to apply in an additive way (i.e and'ed). From Pod's perspective, a pod can use one or N of the allowed items.

An AdmissionController must be added to the validating phase and must reject pod scheduling if the serviceaccount running the pod is not allowed to specify requested NodeSelectors, Scheduler-Name, Anti-Affinity rules, Priority class, and Tolerations.

All usable scheduling policies (allowed by RBAC) are merged before evaluating if scheduling constraints defined in pods are allowed.

## SchedulingPolicy

Proposed API group: `extensions/v1alpha1`

SchedulingPolicy is a cluster-scoped resource (not namespaced).

### SchedulingPolicy content

SchedulingPolicy spec is composed of optional fields that allow scheduling rules. If a field is absent from a SchedulingPolicy, this schedpol won't allow any item from the missing field.

```yaml
apiVersion: extensions/valpha1
kind: SchedulingPolicy
metadata:
  name: my-schedpol
spec:
  allowedSchedulerNames:      # Describes schedulers names that are allowed
  allowedPriorityClasseNames: # Describes priority classe names that are allowed
  allowedNodeSelectors:       # Describes node selectors that can be used
  allowedTolerations:         # Describes tolerations that can be used
  allowedAffinities:          # Describes affinities that can be used
```

### Scheduler name

It should be possible to allow users to use only specific schedulers using `allowedSchedulerNames` field.

If `allowedSchedulerNames` is absent from SchedulingPolicy, no scheduler is allowed by this specific policy.

#### Examples

Allow serviceaccounts to use either the default-scheduler (which is used by specifying `spec.schedulerName` in pod definition) or the `my-scheduler` scheduler (by specifying `spec.schedulerName: "my-scheduler"`):
```yaml
Kind: SchedulingPolicy
spec:
  allowedSchedulerNames:
  - default-scheduler
  - my-scheduler
```


Allow all schedulers:
```yaml
Kind: SchedulingPolicy
spec:
  allowedSchedulerNames: []
```


### Tolerations

Toleration usage can be allowed using fine-grain rules with `allowedTolerations` field. If specifying multiple `allowedTolerations`, pod will be scheduled if one of the allowedTolerations is satisfied.

If `allowedTolerations` is absent from SchedulingPolicy, no toleration is allowed.

#### Examples

##### Fine-grain allowedTolerations
```yaml
Kind: SchedulingPolicy
spec:
  allowedTolerations:
  - keys: ["mykey"]
    operators: ["Equal"]
    values: ["value"]
    effects: ["NoSchedule"]
  - keys: ["other_key"]
    operators: ["Exists"]
    effects: ["NoExecute"]
```
This example allows tolerations in the following forms:
- tolerations that tolerates taints with key named `mykey` that has a value `value` and with a `NoSchedule` effect.
- tolerations that tolerates taints with key `other_key` that has a `NoExecute` effect.

##### Coarse-grain allowedTolerations
```yaml
Kind: SchedulingPolicy
spec:
  allowedTolerations:
  - keys: []
    operators: []
    values: []
    effects: ["PreferNoSchedule"]
  - keys: []
    operators: ["Exists"]
    effects: ["NoSchedule"]
```
This example allows tolerations in the following forms:
- tolerations that tolerates all `PreferNoSchedule` taints with any value.
- tolerations that tolerates taints based on any key existence with effect `NoSchedule`.
Also note that this SchedulingPolicy does not allow tolerating NoExecute taints.


### Priority classes

We must be able to enforce users to use specific priority classes by using the `allowedPriorityClasseNames` field.

If `allowedPriorityClasseNames` is absent from SchedulingPolicy, no priority class is allowed.

#### Examples

##### Only allow a single priority class
```yaml
Kind: SchedulingPolicy
spec:
  allowedPriorityClasseNames:
  - high-priority
```
In this example, only the `high-priority` PriorityClass is allowed.


##### Allow all priorities

```yaml
Kind: SchedulingPolicy
spec:
  allowedPriorityClasseNames: []
```
In this example, all priority classes are allowed.

### Node Selector

As anti-affinity rules are really time-consuming, we must be able to restrict their usage with `allowedNodeSelectors`.

If `allowedNodeSelectors` is totally absent from the spec, no node selector is allowed.

#### Examples

##### Fine-grained policy

```yaml
Kind: SchedulingPolicy
spec:
  allowedNodeSelectors:
    disk: ["ssd"]
    region: [] # means any value
```
In this example, pods can be scheduled only if they either:
-  have no nodeSelector
-  or have a `disk: ssd` nodeSelector
-  and / or have a `region` key nodeSelector with any value

### Affinity rules

As anti-affinity rules are really time-consuming, we must be able to restrict their usage with `allowedAffinities`.
`allowedAffinities` is supposed to keep a coarse-grained approach in allowing affinities. For each type (`nodeAffinities`, `podAffinities`, `podAntiAffinities`) a schedulingpolicy can list allowed constraints (`requiredDuringSchedulingIgnoredDuringExecution`
or `requiredDuringSchedulingIgnoredDuringExecution`).

If `allowedAffinities` is totally absent from the spec, no affinity is allowed whatever its kind.

#### Examples

##### Basic policy
```yaml
Kind: SchedulingPolicy
spec:
  allowedAffinities:
    nodeAffinities:
    - requiredDuringSchedulingIgnoredDuringExecution
    podAntiAffinities:
    - requiredDuringSchedulingIgnoredDuringExecution
    - preferredDuringSchedulingIgnoredDuringExecution
```

##### Allow-all policy
In this example, all affinities are allowed:
```yaml
Kind: SchedulingPolicy
spec:
  allowedAffinities:
    nodeAffinities: []
    podAffinities: []
    podAntiAffinities: []
```

If a sub-item of allowedAffinities is absent from SchedulingPolicy, it is not allowed e.g:
```yaml
Kind: SchedulingPolicy
spec:
  allowedAffinities:
    nodeAffinities: []
```
In this example, only soft and hard nodeAffinities are allowed.

### When both `allowedNodeSelectors` and `nodeAffinities` are specified

Use of both `allowedNodeSelectors` and `nodeAffinities` is not recommended as the latter being way more permissive.

## Default SchedulingPolicies

### Restricted policy
Here is a reasonable policy that might be allowed for any cluster without specific needs:
```yaml
apiVersion: extensions/valpha1
kind: SchedulingPolicy
metadata:
  name: restricted
spec:
  allowedSchedulerNames: ["default-scheduler"]
```
It only allows usage of the default scheduler, no tolerations, nodeSelectors nor affinities.

Multi-archi (x86_64, arm) or multi-OS (Linux, Windows) clusters might also allow the following nodeSelectors:
```yaml
apiVersion: extensions/valpha1
kind: SchedulingPolicy
metadata:
  name: restricted
spec:
  allowedSchedulerNames: ["default-scheduler"]
  allowedNodeSelectors:
    beta.kubernetes.io/arch: []
    beta.kubernetes.io/os: []
```

### Privileged Policy

This is the privileged SchedulingPolicy, it allows usage of all schedulers, priority classes, nodeSelectors, affinities and tolerations.

```yaml
apiVersion: extensions/valpha1
kind: SchedulingPolicy
metadata:
  name: privileged
spec:
  allowedSchedulerNames: []
  allowedPriorityClasseNames: []
  allowedNodeSelectors: {}
  allowedTolerations:
  - keys: [] # any keys
    operators: [] # => Equivalent to ["Exists", "Equals"]
    values: [] # any values
    effects: [] # => Equivalent to ["PreferNoSchedule", "NoSchedule", "NoExecute"]
  allowedAffinities:
    nodeAffinities: []
    podAffinities: []
    podAntiAffinities: []
```

## RBAC
SchedulingPolicy are supposed to be allowed using the verb `use` to apply at pod runtime

the following default ClusterRoles / ClusterRoleBindings are supposed to be provisioned to ensure at least the default-scheduler can be used.

RBAC objects are going to be auto-provisioned at cluster creation / upgrade.


This ClusterRole allows the use of the default scheduler:
```yaml
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  annotations:
    rbac.authorization.kubernetes.io/autoupdate: "true"
  labels:
    kubernetes.io/bootstrapping: rbac-defaults
  name: sp:restricted
rules:
- apiGroups: ['extensions']
  resources: ['schedulingpolicies']
  verbs:     ['use']
  resourceNames:
  - restricted
```

This ClusterRoleBinding ensures any serviceaccount can use the default-scheduler:
```yaml
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  annotations:
    rbac.authorization.kubernetes.io/autoupdate: "true"
  labels:
    kubernetes.io/bootstrapping: rbac-defaults
  name: sp:restricted
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: sp:restricted
subjects:
- kind: Group
  name: system:authenticated
  apiGroup: rbac.authorization.k8s.io
```

This RoleBinding ensures that kube-system pods can run with no scheduling restriction:
```yaml
apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  annotations:
    rbac.authorization.kubernetes.io/autoupdate: "true"
  labels:
    kubernetes.io/bootstrapping: rbac-defaults
  name: sp:kube-system-privileged
  namespace: kube-system
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: sp:privileged
subjects:
- kind: Group
  name: system:serviceaccounts:kube-system
  apiGroup: rbac.authorization.k8s.io
```
# References
-  [Pod affinity/anti-affinity](https://kubernetes.io/docs/concepts/configuration/assign-pod-node/#affinity-and-anti-affinity)
-  [Pod priorities](https://kubernetes.io/docs/concepts/configuration/pod-priority-preemption/)
-  [Taints and tolerations](https://kubernetes.io/docs/concepts/configuration/taint-and-toleration/)
-  [RBAC](https://kubernetes.io/docs/admin/authorization/rbac/)
-  [Using multiple schedulers](https://kubernetes.io/docs/tasks/administer-cluster/configure-multiple-schedulers/)
