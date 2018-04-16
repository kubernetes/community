# Scheduling Policy

_Status: Draft_

_Authors: @arnaudmz, @yastij_

_Reviewers: @bsalamat, @liggitt_

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
-  placing pods on specific nodes (master roles for instance).
-  using specific priority classes.
-  expressing pod affinity or anti-affinity rules.

Also Mandatory (and optionally default) values must also be enforced by scheduling policies in case of:

-  mutli-arch (amd64, arm64) of multi-os (Linux, Windows) clusters
-  multi-az / region / failure domain clusters

# Overview

### syntaxic standards

```yaml
apiVersion: policy/v1alpha1
kind: <POLICY_KIND>
metadata:
  name: <POLICY_NAME>
spec:
  bindingMode: <BINDING_MODE>            # Describes a bindingMode (any or all)
  namespaces:                            # a list of namespaces that the policy targets (optional)
    - ns1
    - ns1
  namespaceSelector:                     # a selector to match a set of namespaces (optional)
    key: value
  rules:                                 # rules that must be satisfied (optional)
    fieldA:                              # name of the field (optional)
      - match:                             # Describes how the rule is matched (required)
        - elt1                           # elements here could be objects like tolerations or strings like SchedulerName
        action:                            # Describes the action that should be taken (allowed, denied or required)

```

### policy composition

`bindingMode` Describes how policies should be composed:

- any: any policy that has its rules satisfied
- all: all policies MUST have its rules satisfied

if we have a heterogenous bindingMode across policies (i.e. some policies with any and others with all).
Then the most restrictive one (i.e the all bindingMode) is applied.

### unset fields:

if a field was not specified in the policy no rule should apply (i.e. everything is allowed).

### empty match:

Usually policies should distinguish between empty matches, since it depends on the action:

- require/deny: no rule apply.
- allow: allows everything.


### inside the matches:

The match field can express pretty much any structure you want to, but there's some things that should be considered:

- the structure should match as much as possible what you try to take action on.
- match elements are ANDed.
- the `*` wildcard is usually needed in the allow section. It should be represented with an empty value (e.g. empty array).


### conflict handling for policies

There is two kinds of conflict in policies:

- structural conflicts: these exists when any rule has the same matchingExpression and opposed actions (deny and require)

- semantical conflicts: these exists due to semantic of the fields (e.g. require a NodeSelector `master=true` and also require `master=false`)


Structural conflicts must handled at creation time, on the other hand, semantical conflicts should be handled at runtime: detect that there's a conflict and emit an event stating that policies couldn't be satisfied due to a conflict.


### required vs allowed vs denied

Policies may have overlapping rules, to handle this policies are computed in the following order:

- compute what is denied.
- compute what is required.
- compute what was allowed.

they should also obey to the following rules:

- everything that is required is by definition allowed.
- everything that is not denied is not automatically allowed: to be allowed a rule must not be denied AND must be allowed.


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
  bindingMode: all
  namespaces:
    - default
  rules:
    schedulerNames:
      - match: []
        action:
    priorityClassNames:
      - match: []
        action:
    tolerations:
      - match: []
        action:
    nodeSelectors:
      - match: []
        action:
    nodeAffinities:
      - match: []
        action:
    podAntiAffinities:
      - match: []
        action:
    podAffinities:
      - match: []
        action:
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
  bindingMode: all
  namespaces:
    - default
  rules:
    schedulerNames:
      - match: ["green-scheduler","my-scheduler"]
        action: required
```

An empty list of schedulerNames has no effect, as pod use the `defaultScheduler` if no scheduler is specified:

```yaml
apiVersion: policy/v1alpha1
kind: SchedulingPolicy
metadata:
  name: mySchedulingPolicy
spec:
  bindingMode: all
  namespaces:
    - default
  rules:
    schedulerNames:
      - match: []
        action: required
```


#### allowed
Allow pods to use either the green-scheduler (which is used by specifying `spec.schedulerName` in pod definition) or the `my-scheduler` scheduler (by specifying `spec.schedulerName: "my-scheduler"`) in the namespace `default`:

```yaml
apiVersion: policy/v1alpha1
kind: SchedulingPolicy
metadata:
  name: mySchedulingPolicy
spec:
  bindingMode: all
  namespaces:
    - default
  rules:
    schedulerNames:
      - match: ["green-scheduler","my-scheduler"]
        action: allowed
```

An empty list of schedulerNames will allow usage of all schedulers:

```yaml
apiVersion: policy/v1alpha1
kind: SchedulingPolicy
metadata:
  name: mySchedulingPolicy
spec:
  bindingMode: all
  namespaces:
    - default
  rules:
    schedulerNames:
      - match: []
        action: allowed
```


### Tolerations

Toleration usage can be regulated using fine-grain rules with `tolerations` field. If specifying multiple `tolerations`, pod will be scheduled if one of the tolerations is satisfied.

#### required

This allows to require toleration in the following forms of:

- tolerations that tolerates taints with key named `projectA-dedicated` with all effects.

##### Fine-grain allowed tolerations

```yaml
apiVersion: policy/v1alpha1
kind: SchedulingPolicy
metadata:
  name: mySchedulingPolicy
spec:
  bindingMode: all
  namespaces:
    - projectA
  rules:
    tolerations:
      - match:
        - keys: ["projectA-dedicated"]
          operators: ["Exists"]
          effects: []
        action: required
```

 an empty list of matches has no effect (i.e. do not require anything).

#### Allowed

This allows requires tolerations in the following forms:
- tolerations that tolerates taints with key named `mykey` that has a value `value` and with a `NoSchedule` effect.
- tolerations that tolerates taints with key `other_key` that has a `NoExecute` effect.

##### Fine-grain allowed tolerations

```yaml
apiVersion: policy/v1alpha1
kind: SchedulingPolicy
metadata:
  name: mySchedulingPolicy
spec:
  bindingMode: all
  namespaces:
    - default
  rules:
    tolerations:
      - match:
        - keys: ["mykey"]
          operators: ["Equal"]
          values: ["value"]
          effects: ["NoSchedule"]
        - keys: ["other_key"]
          operators: ["Exists"]
          effects: ["NoExecute"]
        action: allowed
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
  bindingMode: all
  namespaces:
    - default
  rules:
    tolerations:
      - match: []
        action: allowed
```

Which is equivalent to:


```yaml
apiVersion: policy/v1alpha1
kind: SchedulingPolicy
metadata:
  name: mySchedulingPolicy
spec:
  bindingMode: all
  namespaces:
    - default
  rules:
    tolerations:
      - match:
        - keys: []
          operators: []
          values: []
          effects: []
        action: allowed
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
  bindingMode: all
  namespaces:
    - default
  rules:
    priorityClasseNames:
      - match: ["high-priority","critical-job"]
        action: required
```



##### Allow

In this example, we only allow the `critical-job` priority.


```yaml
apiVersion: policy/v1alpha1
kind: SchedulingPolicy
metadata:
  name: mySchedulingPolicy
spec:
  bindingMode: all
  namespaces:
    - default
  rules:
    priorityClasseNames:
      - match: ["critical-priority"]
        action: allowed
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
  bindingMode: all
  namespaces:
    - default
  rules:
    nodeSelectors:
      - match:
        - beta.kubernetes.io/arch: ["amd64", "arm64"]
        - team: []
        action: required
      - match:
        - failure-domain.beta.kubernetes.io/region: []
        - disk: ["ssd", "hdd"]
        action: allowed

```


In this example, pods can be scheduled if they match the two following:
-  a nodeSelector `beta.kubernetes.io/arch=amd64` or `beta.kubernetes.io/arch=arm64`.
- a nodeSelector with the key `team`.

They can also optionally specify:
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
  bindingMode: all
  namespaces:
    - default
  rules:
    nodeAffinities:
      - match:
        - keys: ["failure-domain.beta.kubernetes.io/region","authorized-region"]
          operators: ["In","NotIn"]
          values: ["eu-2", "us-1"]
          type: "requiredDuringSchedulingIgnoredDuringExecution"
        action: allowed
      - match:
        - keys: ["beta.kubernetes.io/arch"]
          operators: ["In"]
          values: ["amd64", "arm64"]
          type: "requiredDuringSchedulingIgnoredDuringExecution"
        action: required
    podAntiAffinities:
      - match: []
        action: allowed

```

In this example, we allow:
- hard NodeAffinity based on
  -  `beta.kubernetes.io/arch` if value is `amd64` or `arm64`.
  -  any combination of <keys, operators, values> specified in `allowed` rule.
- All podAntiAffinities
- No podAffinities


##### By default behavior

by default no `SchedulingPolicy` is created, so any workload will be running as expected (i.e. no restriction apply).


# References
-  [Pod affinity/anti-affinity](https://kubernetes.io/docs/concepts/configuration/assign-pod-node/#affinity-and-anti-affinity)
-  [Pod priorities](https://kubernetes.io/docs/concepts/configuration/pod-priority-preemption/)
-  [Taints and tolerations](https://kubernetes.io/docs/concepts/configuration/taint-and-toleration/)
-  [Using multiple schedulers](https://kubernetes.io/docs/tasks/administer-cluster/configure-multiple-schedulers/)
