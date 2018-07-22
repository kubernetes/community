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

### syntaxic standards

```yaml
apiVersion: policy/v1alpha1
kind: <POLICY_KIND>
metadata:
  name: <POLICY_NAME>
spec:
  priority: <exception,cluster,user,default>
  namespaces:
    matchNames:
      - ns1
      - ns2
    matchLabels:
      key: value
  action:                            # Describes the action that should be taken (allowed, denied or required)      
  rules:                                 # rules that must be satisfied (optional)
    fieldA:                              # name of the field (optional)
      - match:                             # Describes how the rule is matched (required)
        - elt1                           # elements here could be objects like tolerations or strings like SchedulerName

```


### empty match and unset fields:

if a field is set to empty in the policy, except when the action is `required`, it should match everything then the corresponding action will apply.

### unset fields:

When a field is not specified it is automatically allowed. This makes it easy to rollout the feature for existing clusters, as it makes everything allowed in the cluster when no policy is created.

### inside the matches:

At the policy level, the match field can express pretty much any structure you want to, but there's some things that should be considered:

- the structure should match as much as possible what you try to take action on.
- match elements are combined.

example :

```
- match:
  - keys: ["projectA-dedicated","projectB-dedicated"]
    operators: ["Exists"]
    effects: []
```

- This matches the structure of toleration
- This match means the following : t
   - The toleration `projectA-dedicated` with the operator `Exists`
   - The toleration `projectB-dedicated` with the operator `Exists`

### policy composition and conflict handling

Policies are composed by ANDing them, note that rules of policies from a lower priority are superseded by ones from a higher priority if there is a conflict.

There is two kinds of conflict to handle in policies of the same priority level: Structural and semantical conflicts.


Structural conflicts must handled at creation time as much as possible, on the other hand, semantical conflicts should be handled at runtime: detect that there's a conflict and emit an event stating that policies couldn't be satisfied due to a conflict.


### how policies are computed

Policies may have overlapping rules, to handle this policies are computed in the following order:

- compute policies at `exception` priority.
- compute policies at `cluster` priority.
- compute policies at `user` priority.
- compute policies at `default` priority.

If a policy doesn't specify a priority the default priority applies.

they should also obey to the following rules:

- everything that is required is by definition allowed.


# Detailed Design


## SchedulingPolicy

Proposed API group: `policy/v1alpha1`


### SchedulingPolicy content

SchedulingPolicy spec is composed of optional fields that allow scheduling rules. If a field is absent from a SchedulingPolicy, this `SchedulingPolicy` won't allow any item from the missing field.

```yaml
apiVersion: policy/v1alpha1
kind: SchedulingPolicy
metadata:
  name: mySchedulingPolicy
spec:
  priority: <exception,cluster,user,default>
  namespaces:
    matchNames:
      - ns1
      - ns2
    matchLabels:
      key: value
  action:  <required,allowed,denied>     
  rules:
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


### required
Elements here are required, pods won't schedule if they aren't present. Also note that if something is required it is also allowed.

### allowed
Elements here are allowed, the policy allows the presence of these elements. From Pod's perspective, a pod can use one or N of the allowed items.

### deny
If pods specify one of these, the pod won't schedule to a node, we won't dive into deny as it is the exact opposite of required.

### Scheduler name

If `schedulerNames` is absent from `allowed`, `default` or `required`, no scheduler is allowed by this specific policy.

#### required

Require that pods use either the green-scheduler (which is used by specifying `spec.schedulerName` in pod definition) or the `my-scheduler` scheduler (by specifying `spec.schedulerName: "my-scheduler"`):

```yaml
apiVersion: policy/v1alpha1
kind: SchedulingPolicy
metadata:
  name: mySchedulingPolicy
spec:
  namespaces:
    matchNames:
      - default
  action: required     
  rules:
    schedulerNames:
       match: ["green-scheduler","my-scheduler"]
```

An empty list of schedulerNames has no effect, as pod use the `defaultScheduler` if no scheduler is specified:

```yaml
apiVersion: policy/v1alpha1
kind: SchedulingPolicy
metadata:
  name: mySchedulingPolicy
spec:
  namespaces:
    matchNames:
      - default
  action: required     
  rules:
    schedulerNames:
       match: []
```


#### allowed
Allow pods to use either the `green-scheduler` (which is used by specifying `spec.schedulerName` in pod definition) or the `my-scheduler` scheduler (by specifying `spec.schedulerName: "my-scheduler"`) in the namespace `default`:

```yaml
apiVersion: policy/v1alpha1
kind: SchedulingPolicy
metadata:
  name: mySchedulingPolicy
spec:
  namespaces:
    matchNames:
      - default
  action: allowed     
  rules:
    schedulerNames:
       match: ["green-scheduler","my-scheduler"]
```

An empty list of schedulerNames will allow usage of all schedulers:

```yaml
apiVersion: policy/v1alpha1
kind: SchedulingPolicy
metadata:
  name: mySchedulingPolicy
spec:
  namespaces:
    matchNames:
      - default
  action: allowed     
  rules:
    schedulerNames:
       match: []
```

note: this policy has no effect as we allow all when no policy is set.

### Tolerations

Toleration usage can be regulated using fine-grain rules with `tolerations` field. If specifying multiple `tolerations`, pod will be scheduled if one of the tolerations is satisfied.

#### required

This requires toleration in the following forms of:

- tolerations that tolerates taints with key named `projectA-dedicated` with all effects.
- tolerations that tolerates taints with key named `node-misc` with `NoSchedule` effect.

```yaml
apiVersion: policy/v1alpha1
kind: SchedulingPolicy
metadata:
  name: mySchedulingPolicy
spec:
  namespaces:
    matchNames:
      - projectA
  action: required     
  rules:
    tolerations:
      match:
        - keys: ["projectA-dedicated"]
          operators: ["Exists"]
          effects: []
        - keys: ["node-misc"]
          operators: ["Exists"]
          effects: ["NoSchedule"]
```

 note: an empty list of matches has no effect (i.e. do not require anything).

#### Allowed

This allows requires tolerations in the following forms:
- tolerations that tolerates taints with key named `mykey` that has a value `value` and with a `NoSchedule` effect.
- tolerations that tolerates taints with key `other_key` that has a `NoExecute` effect.

```yaml
apiVersion: policy/v1alpha1
kind: SchedulingPolicy
metadata:
  name: mySchedulingPolicy
spec:
  namespaces:
    matchNames:
      - projectA
  action: allowed   
  rules:
    tolerations:
       match:
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


an empty list of toleration allows all types of tolerations:

```yaml
apiVersion: policy/v1alpha1
kind: SchedulingPolicy
metadata:
  name: mySchedulingPolicy
spec:
  namespaces:
    matchNames:
      - projectA
  action: allowed   
  rules:
    tolerations:
      match: []
```

Which is equivalent to:


```yaml
apiVersion: policy/v1alpha1
kind: SchedulingPolicy
metadata:
  name: mySchedulingPolicy
spec:
  namespaces:
    matchNames:
      - projectA
  action: allowed   
  rules:
    tolerations:
       match:
        - keys: []
          operators: []
          values: []
          effects: []
```



### Priority classes

Priority class usage can be regulated using fine-grain rules with `priorityClasseName` field.

##### required

this example requires a `priorityClass` that is either `high-priority` or `critical-job`

```yaml
apiVersion: policy/v1alpha1
kind: SchedulingPolicy
metadata:
  name: mySchedulingPolicy
spec:
  namespaces:
    matchNames:
      - default
  action: required   
  rules:
    priorityClasseNames:
       match: ["high-priority","critical-job"]
```



##### Allow

In this example, we only allow the `critical-job` priority.


```yaml
apiVersion: policy/v1alpha1
kind: SchedulingPolicy
metadata:
  name: mySchedulingPolicy
spec:
  namespaces:
    matchNames:
      - default
  action: allowed   
  rules:
    priorityClasseNames:
      - match: ["critical-priority"]
```




### Node Selector

The `nodeSelector` field makes it possible to specify what nodeSelectors are required, allowed and denied. As for other components, `required` nodeSelectors are automatically considered as allowed.

#### Examples

##### Complete policy

```yaml
apiVersion: policy/v1alpha1
kind: SchedulingPolicy
metadata:
  name: mySchedulingPolicy
spec:
  namespaces:
    matchNames:
      - default
  action: required   
  rules:
    nodeSelectors:
      match:
        - failure-domain.beta.kubernetes.io/region: ["northeurope"]
        - disk: ["ssd", "premium"]
```


In this example, pods cannot be scheduled if they match one of the following:
-  `disk: ssd` nodeSelector,
-  `disk: hdd` nodeSelector,
-  `failure-domain.beta.kubernetes.io/region` nodeSelector with any value.


### Affinity rules

As anti-affinity rules are really time-consuming, we must be able to restrict their usage with `allowedAffinities`.
`allowedAffinities` is supposed to keep a coarse-grained approach in allowing affinities. For each type (`nodeAffinities`, `podAffinities`, `podAntiAffinities`) a schedulingpolicy can list allowed constraints (`requiredDuringSchedulingIgnoredDuringExecution`
or `requiredDuringSchedulingIgnoredDuringExecution`).

If `allowedAffinities` is totally absent from the spec, no affinity is allowed whatever its kind.

#### Examples

##### Basic policy



```yaml
apiVersion: policy/v1alpha1
kind: SchedulingPolicy
metadata:
  name: mySchedulingPolicy
spec:
  namespaces:
    matchNames:
      - default
  action: required   
  rules:
    nodeAffinities:
      requiredDuringSchedulingIgnoredDuringExecution:
         match:
          - keys: ["failure-domain.beta.kubernetes.io/region","authorized-region"]
            operator: "NotIn"
            values: ["eu-2", "us-1"]
          - keys: ["PCI-region"]
            operator: "Exists"
            values: []          
      preferredDuringSchedulingIgnoredDuringExecution:
         match:
          - keys: ["flavor"]
            operator: In
            values: ["m1.small", "m1.medium"]

```

In this example, we require pods to use nodeAffinity to select nodes having `failure-domain.beta.kubernetes.io/region` or `authorized-region` without `eu-1` or `us-1` values, or nodes having `PCI-region` label set. On those filtered nodes we require the pod to prefer nodes with the lowest compute capabilities (`m1.small` or `m1.medium`)


# References
-  [Pod affinity/anti-affinity](https://kubernetes.io/docs/concepts/configuration/assign-pod-node/#affinity-and-anti-affinity)
-  [Pod priorities](https://kubernetes.io/docs/concepts/configuration/pod-priority-preemption/)
-  [Taints and tolerations](https://kubernetes.io/docs/concepts/configuration/taint-and-toleration/)
-  [Using multiple schedulers](https://kubernetes.io/docs/tasks/administer-cluster/configure-multiple-schedulers/)
