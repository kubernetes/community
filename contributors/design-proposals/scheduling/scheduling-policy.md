# Scheduling Policy

_Status: Draft_

_Authors: @arnaudmz, @yastij_

_Reviewers: @bsalamat, @tallclair, @liggitt_

# Objectives

-  Define the concept of scheduling policies
-  Propose their initial design and scope

## Non-Goals

-  How taints / tolerations work.
-  How NodeSelector works.
-  How node / pod affinity / anti-affinity rules work.
-  How several schedulers can be used within a single cluster.
-  How priority classes work.
-  How to set defaults in Kubernetes.

# Background

During real-life Kubernetes architecturing we encountered contexts where role-isolation (between administration and simple namespace usage in a multi-tenant context) could be improved. So far, no restriction is possible on toleration, priority class usage, nodeSelector, anti-affinity depending on user permissions (RBAC).

Identified use-cases aim to ensure that administrators have a way to restrict users or namespaces, it allows administrators to:

-  Restrict execution for specific applications (which are namespace scoped) into certain nodes
-  Create policies that prevent users from even attempting to schedule workloads onto masters to maximize security
-  require that a pods under a namespace run on dedicated nodes
-  Restrict usage of some `PriorityClass`
-  Restrict usage to a specific set of schedulers.
-  enforcing pod affinity or anti-affinity rules on some particular namespace.

Also Mandatory values must also be enforced by scheduling policies in case of:

-  mutli-arch (amd64, arm64) of multi-os (Linux, Windows) clusters (will also be handled later by the [RuntimeClass]() KEP)
-  multi-az / region / failure domain clusters

# Overview

The schedulingPolicy will live out-of-tree under kubernetes-sigs org. It will use a CRD-based approach.

## syntaxic standards

### empty match and unset fields:

an empty field and an unset one matches everything.

### how policies are computed

Policies are computed using the following algorithm:

```
sortedPolicies = sort_by_priority(sort_deny_first(policies))
for policy in sortedPolicies:
  if policy matches pod: // all specified policy rules match
    return policy.action
```
 note that:
- rules of policies with higher priority supersede lower priority rules if they both match.
- matching is done statically, i.e. we don't interpret logical operators (see nodeAffinity section for more details).
- matching is considered true if a subset of a set-based field is matched.



# Detailed Design


## SchedulingPolicy

Proposed API group: `policy.k8s.io/v1alpha1`


### SchedulingPolicy content

SchedulingPolicy spec is composed of optional fields that allow scheduling rules. If a field is absent from a `SchedulingPolicy` it is automatically allowed.

```yaml
apiVersion: policy.k8s.io/v1alpha1
kind: SchedulingPolicy
metadata:
  name: mySchedulingPolicy
spec:
  priority: <exception,cluster,default>
  action:  <allow,deny>     
  rules:
    namespaceSelector:
      - key: value
        operator: operator
    podSelector:
        key: value
    schedulerNames:
       match: []
    priorityClassNames:
       match: []
    tolerations:
       match: []
    nodeSelectors:
       match: []
    nodeAffinities:
       match: []
    podAntiAffinities:
       match: []
    podAffinities:
       match: []
```

- allow: When a policy matches the pod is allowed
- deny: when a policy matches the pod is denied


### Scoping

The `spec.rules.namespaceSelector` and `spec.rules.podSelector` attributes scope the policy application range.

Pods must match both `spec.namespaceSelector` and `spec.podSelector` matching rules to be constrained by a policy.
Both `spec.namespaceSelector` and `spec.podSelector` are optional. If absent, all pods and all namespaces are targeted by the policy.

### Scheduler name

#### Allowed
Allow pods to use either the `green-scheduler` (which is used by specifying `spec.schedulerName` in pod definition) or the `my-scheduler` scheduler (by specifying `spec.schedulerName: "my-scheduler"`) in namespaces labeled `team`:

```yaml
apiVersion: policy.k8s.io/v1alpha1
kind: SchedulingPolicy
metadata:
  name: mySchedulingPolicy
spec:
  action: allow    
  rules:
    namespaces:
      - key: team
        operator: Exists
    podSelector: {}
    schedulerNames:
       match: ["green-scheduler","my-scheduler"]
```


```yaml
apiVersion: policy.k8s.io/v1alpha1
kind: SchedulingPolicy
metadata:
  name: mySchedulingPolicy
spec:
  action: allow
  rules:
    namespaces:
      - key: team
        operator: Exists
    podSelector: {}    
    schedulerNames:
       match: []
```

note: this policy has no effect as we allow all when no policy is set.

### Tolerations

Toleration usage can be regulated using fine-grain rules with `tolerations` field. If specifying multiple `tolerations`, pod will be scheduled if one of the tolerations is satisfied.


#### Allowed

This allows pods with tolerations the following:
- tolerations that tolerates taints with key named `mykey` that has a value `value` and with a `NoSchedule` effect.
- tolerations that tolerates taints with key `other_key` that has a `NoExecute` effect.

```yaml
apiVersion: policy.k8s.io/v1alpha1
kind: SchedulingPolicy
metadata:
  name: mySchedulingPolicy
spec:
  action: allow   
  rules:
    namespaces:
      - key: team
        operator: Exists
    podSelector: {}    
    tolerations:
       match:
        - key: "mykey"
          operator: "Equal"
          value: "value"
          effects: "NoSchedule"
        - key: "other_key"
          operator: "Exists"
          effect: "NoExecute"
```


### Priority classes

Priority class usage can be regulated using fine-grain rules with `priorityClasseName` field.

##### Allow

In this example, we only allow the `critical-job` priority.


```yaml
apiVersion: policy.k8s.io/v1alpha1
kind: SchedulingPolicy
metadata:
  name: mySchedulingPolicy
spec:
  action: allow
  rules:
    namespaces:
      - key: team
        operator: Exists
    podSelector: {}
    priorityClasseNames:
       match: "critical-priority"
```


### Node Selector

The `nodeSelector` field makes it possible to specify what pods are allowed or denied based on their nodeSelectors.

#### Examples

##### Complete policy

```yaml
apiVersion: policy.k8s.io/v1alpha1
kind: SchedulingPolicy
metadata:
  name: mySchedulingPolicy
spec:
  action: deny   
  rules:
    namespaces:
      - key: team
        operator: Exists
    podSelector: {}
    nodeSelectors:
      match:
        - failure-domain.beta.kubernetes.io/region: ""
        - disk: "ssd"
        - disk: "premium"
```


In this example, pods cannot be scheduled if they have all of the following at the same time:
-  `disk: ssd` nodeSelector
-  `disk: hdd` nodeSelector
-  `failure-domain.beta.kubernetes.io/region` nodeSelector with any value.


### Affinity rules

As anti-affinity rules are really time-consuming, we must be able to restrict their usage for each type (`nodeAffinities`, `podAffinities`, `podAntiAffinities`) a schedulingpolicy can list allowed/denied constraints (`requiredDuringSchedulingIgnoredDuringExecution`
or `requiredDuringSchedulingIgnoredDuringExecution`).

#### Examples

##### Basic policy



```yaml
apiVersion: policy.k8s.io/v1alpha1
kind: SchedulingPolicy
metadata:
  name: mySchedulingPolicy
spec:
  action: allow   
  rules:
    namespaces:
      - key: team
        operator: Exists
    podSelector: {}
    nodeAffinities:
      requiredDuringSchedulingIgnoredDuringExecution:
         match:
          - key: "authorized-region"
            operator: "NotIn"
            values: ["eu-2", "us-1"]
          - key: "PCI-region"
            operator: "Exists"
            values: []          
      preferredDuringSchedulingIgnoredDuringExecution:
         match:
          - key: "flavor"
            operator: In
            values: ["m1.small", "m1.medium"]

```

In this example, we allow pods with nodeAffinity to select nodes having `authorized-region` without `eu-1` or `us-1` values, or nodes having `PCI-region` label set. On those filtered nodes we require the pod to prefer nodes with the lowest compute capabilities (`m1.small` or `m1.medium`). The matching is done when a pod has:

- All the "required" and "preferred" sections.
- Each section has the same keys and the same operators.
- Values must be the same or subset of those of the pod.



# References
-  [Pod affinity/anti-affinity](https://kubernetes.io/docs/concepts/configuration/assign-pod-node/#affinity-and-anti-affinity)
-  [Pod priorities](https://kubernetes.io/docs/concepts/configuration/pod-priority-preemption/)
-  [Taints and tolerations](https://kubernetes.io/docs/concepts/configuration/taint-and-toleration/)
-  [Using multiple schedulers](https://kubernetes.io/docs/tasks/administer-cluster/configure-multiple-schedulers/)
