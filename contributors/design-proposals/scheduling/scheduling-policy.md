# Scheduling Policy

_Status: Draft_

_Authors: @arnaudmz, @yastij_

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

During real-life Kubernetes architecturing we encountered contexts where role-isolation (between administration and simple namespace usage in a multi-tenant context) could be improved. So far, no restriction is possible on toleration, priority class usage, nodeSelector, anti-affinity depending on user permissions (RBAC).

Identified use-cases aim to ensure that administrators have a way to restrict users or namepace when
-  using schedulers,
-  placing pods on specific nodes (master roles for instance),
-  using specific priority classes,
-  expressing pod affinity or anti-affinity rules.

Mandatory (and optionally default) values must also be enforced by scheduling policies in case of:
-  mutli-arch (amd64, arm64) of multi-os (Linux, Windows) clusters
-  multi-az / region / failure domain clusters

# Overview

Implementing SchedulingPolicy implies:
-  Creating a new resource named **SchedulingPolicy** (schedpol)
-  Creating an **AdmissionController** that dehaves on a deny-all-but basis
-  Allow SchedulingPolicy to be used by pods using RoleBindings or ClusterRoleBindings

# Detailed Design

SchedulingPolicy resource specs are composed of several main attributes:
-  **Required** scheduling components. These list the manadory NodeSelectors, Scheduler-Name, Anti-Affinity rules, Priority class, and Tolerations and optionally valid values that have to be provided to allow scheduling.
-  **Allowed** scheduling components. These list the optional components that can be specified in pods definition.
-  **Default** scheduling components. These list default values to set unless specified in pods definition.

SchedulingPolicy resources are supposed to apply in a deny-all-except approach.

An AdmissionController must be added to the mutating phase to
-  add default values if unspecified for NodeSelectors, Scheduler-Name, Anti-Affinity rules, Priority class, and Tolerations,
-  reject pod scheduling if the serviceaccount running the pod is not allowed to specify requested NodeSelectors, Scheduler-Name, Anti-Affinity rules, Priority class, and Tolerations.

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
  required:
    schedulerNames:      # Describes schedulers names that are required
    priorityClasseNames: # Describes priority class names that are required
    nodeSelectors:       # Describes node selectors that must be used
    affinities:          # Describes affinities that must be used
  allowed:
    schedulerNames:      # Describes schedulers names that are allowed
    priorityClasseNames: # Describes priority class names that are allowed
    nodeSelectors:       # Describes node selectors that can be used
    tolerations:         # Describes tolerations that can be used
    affinities:          # Describes affinities that can be used
  default:
    schedulerName:       # Describes default scheduler name
    priorityClasseName:  # Describes default priority class name
    nodeSelector:        # Describes default node selector
    tolerations:         # Describes default tolerations
    affinity:            # Describes default affinity
```

### required
Elements here are required, pods won't schedule if they aren't present. Also note that if something is required it is also allowed.

### allowed
Elements here are allowed, the policy allows the presence of these elements. From Pod's perspective, a pod can use one or N of the allowed items.

### default
If pods do not specify one, the elements here will be added


### Scheduler name

If `schedulerNames` is absent from `allowed`, `default` or `required`, no scheduler is allowed by this specific policy.

#### required

Require that pods use either the green-scheduler (which is used by specifying `spec.schedulerName` in pod definition) or the `my-scheduler` scheduler (by specifying `spec.schedulerName: "my-scheduler"`):
```yaml
Kind: SchedulingPolicy
spec:
  required:
    SchedulerNames: ["green-scheduler", "my-scheduler"]
```

An empty list of schedulerNames here is not a valid syntax:
```yaml
Kind: SchedulingPolicy
spec:
  required:
    SchedulerNames: [] # equivalent to ["default-scheduler"] or to not specifying this item
```


#### allowed
Allow pods to use either the green-scheduler (which is used by specifying `spec.schedulerName` in pod definition) or the `my-scheduler` scheduler (by specifying `spec.schedulerName: "my-scheduler"`):
```yaml
Kind: SchedulingPolicy
spec:
  allowed:
    SchedulerNames: ["green-scheduler", "my-scheduler"]
```

An empty list of schedulerNames will allow usage of all schedulers:
```yaml
Kind: SchedulingPolicy
spec:
  allowed:
    SchedulerNames: []
```

#### default
pods will default use either the my-scheduler if nothing is specified in `spec.schedulerName`:
```yaml
Kind: SchedulingPolicy
spec:
  default:
    SchedulerName: "my-scheduler"
```


### Tolerations

Toleration usage can be regulated using fine-grain rules with `tolerations` field. If specifying multiple `tolerations`, pod will be scheduled if one of the tolerations is satisfied.


#### Allowed

This allows requires tolerations in the following forms:
- tolerations that tolerates taints with key named `mykey` that has a value `value` and with a `NoSchedule` effect.
- tolerations that tolerates taints with key `other_key` that has a `NoExecute` effect.

##### Fine-grain allowed tolerations

```yaml
Kind: SchedulingPolicy
spec:
  allowed:
    tolerations:
    - keys: ["mykey"]
      operators: ["Equal"]
      values: ["value"]
      effects: ["NoSchedule"]
    - keys: ["other_key"]
      operators: ["Exists"]
      effects: ["NoExecute"]
```

Here we allow tolerations in the following forms:
- tolerations that tolerates all `PreferNoSchedule` taints with any value.
- tolerations that tolerates taints based on any key existence with effect `NoSchedule`.
Also note that this SchedulingPolicy does not allow tolerating NoExecute taints.

##### Coarse-grain allowed tolerations

```yaml
Kind: SchedulingPolicy
spec:
  allowed:
    tolerations:
    - keys: []
      operators: []
      values: []
      effects: ["PreferNoSchedule"]
    - keys: []
      operators: ["Exists"]
      effects: ["NoSchedule"]
```
an empty list of toleration allows all types of tolerations:

```yaml
Kind: SchedulingPolicy
spec:
  allowed:
    tolerations: []
```

Which is equivalent to:

```yaml
Kind: SchedulingPolicy
spec:
  allowed:
    tolerations:
    - keys: []
      operators: []
      values: []
      effects: []
```

#### default

if no toleration is not specified, the following `SchedulingPolicy` will add a:
- toleration for <mykey,value,NoSchedule> and <mykey,other_value,NoSchedule> Taints.
- toleration for <other_key,NoExecute> Taint.

```yaml
Kind: SchedulingPolicy
spec:
  default:
    tolerations:
    - key: "mykey"
      operator: "Equal"
      values: ["value","other_value"]
      effect: "NoSchedule"
    - key: "other_key"
      operator: "Exists"
      effect: "NoExecute"
```
note: an empty array of toleration is not a valid syntax for default toleration.

### Priority classes

Priority class usage can be regulated using fine-grain rules with `priorityClasseName` field.

##### Only allow a single priority class

```yaml
Kind: SchedulingPolicy
spec:
  required:
    priorityClasseNames: ["high-priority"]
  default:
    priorityClasseName: "high-priority"
```
In this example, only the `high-priority` PriorityClass is enforced by default.

##### Allow all priorities

```yaml
Kind: SchedulingPolicy
spec:
  allowed:
    priorityClasseNames: []
```
In this example, all priority classes are allowed, but not mandatory.

Note: an empty list of required priorityClasseNames is considered as invalid
```yaml
Kind: SchedulingPolicy
spec:
  required:
    priorityClasseNames: []
```


### Node Selector

The `nodeSelector` fields in `required`, `default` and `allowed` sections allow to precise what nodeSelectors are mandatory, possible and may provide default values if not set. As for other components, `required` nodeSelectors are automatically considered as allowed.

#### Examples

##### Complete policy

```yaml
Kind: SchedulingPolicy
spec:
  required:
    nodeSelectors:
      beta.kubernetes.io/arch: ["amd64", "arm64"] # pick one
  default: # if not set, inject this to pods definitions
    nodeSelector:
      beta.kubernetes.io/arch: "amd64"
  allowed: # other optional allowed nodeSelectors
    nodeSelectors:
      disk: ["ssd", "hdd"]
      failure-domain.beta.kubernetes.io/region: [] # means any value
```

In this example, pods can be scheduled if they:
-  have no nodeSelector at all. The default `beta.kubernetes.io/arch=amd64` will then be assigned.
-  have a nodeSelector `beta.kubernetes.io/arch=amd64` or `beta.kubernetes.io/arch=arm64`

They can also optionally have:
-  `disk: ssd` nodeSelector,
-  `disk: hdd` nodeSelector,
-  `failure-domain.beta.kubernetes.io/region` nodeSelector with any value.

##### Allowed-only policy

```yaml
Kind: SchedulingPolicy
spec:
  allowed: # other optional allowed nodeSelectors
    nodeSelectors:
      failure-domain.beta.kubernetes.io/zone: ["eu-west-1a", "eu-west-1b", "eu-west-1c"]
```

In this example, pods can be scheduled if they:
-  have no nodeSelector at all.
-  `failure-domain.beta.kubernetes.io/zone` nodeSelector with a value in the three listed: `eu-west-1a`, `eu-west-1b` or `eu-west-1c`.

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
  required:
    affinities:
      nodeAffinities:
        requiredDuringSchedulingIgnoredDuringExecution:
          nodeSelectorTerms:
          - matchExpressions:
            - keys: ["beta.kubernetes.io/arch"]
              operators: ["In"]
              values: ["amd64", "arm64"]
  allowed:
    affinities:
      nodeAffinities:
        requiredDuringSchedulingIgnoredDuringExecution:
          nodeSelectorTerms:
          - matchExpressions:
            - keys: ["failure-domain.beta.kubernetes.io/region","kubernetes.io/authorized-region"]
              operators: ["In","NotIn"]
              values: ["eu-2", "us-1"]
      podAntiAffinities: {}
  default:
    affinity:
      nodeAffinity:
        requiredDuringSchedulingIgnoredDuringExecution:
          nodeSelectorTerms:
          - matchExpressions:
            - keys: ["beta.kubernetes.io/arch"]
              operators: ["In"]
              values: ["amd64"]
```

In this example, we allow:
- hard NodeAffinity based on
  -  `beta.kubernetes.io/arch` if value is `amd64` or `arm64`. Defaults to `amd64` if unspecified.
  -  (optionally) any combination of <keys, operators, values> specified in `allowed` secion.
- All podAntiAffinities
- No podAffinities

```yaml
Kind: SchedulingPolicy
spec:
  allowed:
    affinities:
      nodeAffinities:
        requiredDuringSchedulingIgnoredDuringExecution:
          nodeSelectorTerms:
          - matchExpressions:
            - keys: ["failure-domain.beta.kubernetes.io/region","kubernetes.io/authorized-region"]
              operators: ["In","NotIn"]
              values: ["eu-2", "us-1"]
          - matchExpressions:
            - keys: ["failure-domain.beta.kubernetes.io/zone"]
              operators: ["NotIn"]
              values: ["dc1", "dc2"]
          podAntiAffinities: {}
```

This example highlights the case where you don't want a full combinatory, here we allow the same affinities as the previous example, in addition,
we allow a match expression that checks that some zone labels are not in the specified values.


##### Allow-all policy
In this example, all affinities are allowed:

```yaml
Kind: SchedulingPolicy
spec:
  allowed:
    affinities: {}
```

Which is equivalent to:

```yaml
Kind: SchedulingPolicy
spec:
  allowed:
    affinities:
      nodeAffinities: {}
      podAffinities: {}
      podAntiAffinities: {}
```

If a sub-item of allowedAffinities is absent from SchedulingPolicy, it is not allowed e.g:

```yaml
Kind: SchedulingPolicy
spec:
  allowedAffinities:
    nodeAffinities: {}
```

In this example, only nodeAffinities (required and preferred) are allowed but no podAffinities nor podAntiAffinities.

## Multiple SchedulingPolicies considerations

NOTE: here a merge is the set of resulting authorizations after going through the available policies (i.e. we do not aggregate policies into a newly created `SchedulingPolicy`)

several merging strategies are being considered.

### Option1: smart deep merge

If RBAC permissions provide a serviceaccount a way to use several schedpols, conflict resolution must occur. The proposed behaviour is:

-  `required` fields use an alphabetic order and then a first-seen-wins strategy for each sub-keys
-  `allowed` fields are additive
-  `default` fields use an alphabetic order and then a first-seen-wins strategy for each sub-keys

for instance, if we have the following two schedpols that apply:

```yaml
Kind: SchedulingPolicy
metadata:
  name: schedpol-a # first in alphabetic order
spec:
  required:
    nodeSelectors:
      beta.kubernetes.io/arch: ["amd64", "arm64"]
  allowed:
    nodeSelectors:
      disk: ["ssd"]
  default:
    nodeSelector:
      beta.kubernetes.io/arch: "amd64"
---
Kind: SchedulingPolicy
metadata:
  name: schedpol-b # second in alphabetic order
spec:
  required:
    nodeSelectors:
      beta.kubernetes.io/arch: ["amd64", "arm64", "i386"]
      beta.kubernetes.io/os: ["Linux", "Windows"]
    priorityClasseNames: ["bronze", "gold", "silver"]
  allowed:
    nodeSelectors:
      disk: ["sata"]
  default:
    nodeSelector:
      beta.kubernetes.io/arch: "i386"
      beta.kubernetes.io/os: "Linux"
    priorityClasseName: "bronze"
```

The merged applied schedpol will be:

```yaml
Kind: SchedulingPolicy
spec:
  required:
    nodeSelectors:
      beta.kubernetes.io/arch: ["amd64", "arm64"]
      beta.kubernetes.io/os: ["Linux", "Windows"]
    priorityClasseNames: ["bronze", "gold", "silver"]
  allowed:
    nodeSelectors:
      disk: ["ssd", "sata"]
  default:
    nodeSelector:
      beta.kubernetes.io/arch: "amd64"
      beta.kubernetes.io/os: "Linux"
    priorityClasseName: "bronze"
```

### Option2: first-seen wins, ever
In this option the strategy is at SchedulingPolicy level.

for instance, if we have the following two schedpols that apply:

```yaml
Kind: SchedulingPolicy
metadata:
  name: schedpol-a # first in alphabetic order
spec:
  required:
    nodeSelectors:
      beta.kubernetes.io/arch: ["amd64", "arm64"]
  allowed:
    nodeSelectors:
      disk: ["ssd"]
  default:
    nodeSelector:
      beta.kubernetes.io/arch: "amd64"
---
Kind: SchedulingPolicy
metadata:
  name: schedpol-b # second in alphabetic order
spec:
  required:
    nodeSelectors:
      beta.kubernetes.io/arch: ["amd64", "arm64", "i386"]
      beta.kubernetes.io/os: ["Linux", "Windows"]
    priorityClasseNames: ["bronze", "gold", "silver"]
  allowed:
    nodeSelectors:
      disk: ["sata"]
  default:
    nodeSelector:
      beta.kubernetes.io/arch: "i386"
      beta.kubernetes.io/os: "Linux"
    priorityClasseName: "bronze"
```

The merged applied schedpol will be:

```yaml
Kind: SchedulingPolicy
metadata:
  name: schedpol-a # first in alphabetic order
spec:
  required:
    nodeSelectors:
      beta.kubernetes.io/arch: ["amd64", "arm64"]
  allowed:
    nodeSelectors:
      disk: ["ssd"]
  default:
    nodeSelector:
      beta.kubernetes.io/arch: "amd64"
```

### Option3: simple merge

In this strategy, merge is performed on a first-seen-wins on second-level entries.
for instance, if we have the following two schedpols that apply:

```yaml
Kind: SchedulingPolicy
metadata:
  name: schedpol-a # first in alphabetic order
spec:
  required:
    nodeSelectors:
      beta.kubernetes.io/arch: ["amd64", "arm64"]
  allowed:
    nodeSelectors:
      disk: ["ssd"]
  default:
    nodeSelector:
      beta.kubernetes.io/arch: "amd64"
---
Kind: SchedulingPolicy
metadata:
  name: schedpol-b # second in alphabetic order
spec:
  required:
    nodeSelectors:
      beta.kubernetes.io/arch: ["amd64", "arm64", "i386"]
      beta.kubernetes.io/os: ["Linux", "Windows"]
    priorityClasseNames: ["bronze", "gold", "silver"]
  allowed:
    nodeSelectors:
      disk: ["sata"]
  default:
    nodeSelector:
      beta.kubernetes.io/arch: "i386"
      beta.kubernetes.io/os: "Linux"
    priorityClasseName: "bronze"
```

The merged applied schedpol will be:

```yaml
Kind: SchedulingPolicy
spec:
  required:
    nodeSelectors:
      beta.kubernetes.io/arch: ["amd64", "arm64"]
    priorityClasseNames: ["bronze", "gold", "silver"]
  allowed:
    nodeSelectors:
      disk: ["ssd"]
  default:
    nodeSelector:
      beta.kubernetes.io/arch: "amd64"
    priorityClasseName: "bronze"
```
## Default SchedulingPolicies

### Restricted policy
Here is a reasonable policy that might be allowed for any cluster without specific needs:

```yaml
apiVersion: extensions/valpha1
kind: SchedulingPolicy
metadata:
  name: restricted
spec:
  allowed:
    schedulerNames: ["default-scheduler"]
```

It only allows usage of the default scheduler, no tolerations, nodeSelectors nor affinities.

Multi-archi (x86_64, arm) or multi-OS (Linux, Windows) clusters might also allow the following nodeSelectors:

```yaml
apiVersion: extensions/valpha1
kind: SchedulingPolicy
metadata:
  name: restricted-multiarch-by-node-selector
spec:
  required:
    schedulerNames: ["default-scheduler"]
    nodeSelectors:
      beta.kubernetes.io/arch: ["amd64", "arm64"] # pick one in required values
      beta.kubernetes.io/os: ["Linux", "Windows"] # pick one
  default: # if not set, inject to pods definitions
    schedulerName: "default-scheduler"
    nodeSelector:
      beta.kubernetes.io/arch: "amd64"
      beta.kubernetes.io/os: "Linux"
```

```yaml
apiVersion: extensions/valpha1
kind: SchedulingPolicy
metadata:
  name: restricted-multiarch-by-affinity
spec:
  required:
    schedulerNames: ["default-scheduler"]
    affinities:
      nodeAffinities:
        requiredDuringSchedulingIgnoredDuringExecution:
          nodeSelectorTerms:
          - matchExpressions:
            - keys: ["beta.kubernetes.io/arch"]
              operators: ["In"]
              values: ["amd64", "arm64"]
          - matchExpressions:
            - keys: ["beta.kubernetes.io/os"]
              operators: ["In"]
              values: ["Linux", "Windows"]
  default: # if not set, inject to pods definitions
    schedulerName: "default-scheduler"
    affinity:
      nodeAffinity:
        requiredDuringSchedulingIgnoredDuringExecution:
          nodeSelectorTerms:
          - matchExpressions:
            - key: "beta.kubernetes.io/arch"
              operator: "In"
              values: ["amd64"]
          - matchExpressions:
            - key: "beta.kubernetes.io/os"
              operator: "In"
              values: ["Linux"]
```

### Privileged Policy

This is the privileged SchedulingPolicy, it allows usage of all schedulers, priority classes, nodeSelectors, affinities and tolerations.

```yaml
apiVersion: extensions/valpha1
kind: SchedulingPolicy
metadata:
  name: privileged
spec:
  allowed:
    schedulerNames: []
    priorityClasseNames: []
    nodeSelectors: {}
    tolerations: []
    affinities: {}
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
