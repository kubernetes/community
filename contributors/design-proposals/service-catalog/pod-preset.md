# Pod Preset

  * [Abstract](#abstract)
  * [Motivation](#motivation)
  * [Constraints and Assumptions](#constraints-and-assumptions)
  * [Use Cases](#use-cases)
    * [Summary](#summary)
    * [Prior Art](#prior-art)
    * [Objectives](#objectives)
  * [Proposed Changes](#proposed-changes)
    * [PodPreset API object](#podpreset-api-object)
      * [Validations](#validations)
    * [AdmissionControl Plug-in: PodPreset](#admissioncontrol-plug-in-podpreset)
      * [Behavior](#behavior)
      * [PodPreset Exclude Annotation](#podpreset-exclude-annotation)
  * [Examples](#examples)
    * [Simple Pod Spec Example](#simple-pod-spec-example)
    * [Pod Spec with `ConfigMap` Example](#pod-spec-with-configmap-example)
    * [ReplicaSet with Pod Spec Example](#replicaset-with-pod-spec-example)
    * [Multiple PodPreset Example](#multiple-podpreset-example)
    * [Conflict Example](#conflict-example)


## Abstract

Describes a policy resource that allows for the loose coupling of a Pod's
definition from additional runtime requirements for that Pod. For example,
mounting of Secrets, or setting additional environment variables,
may not be known at Pod deployment time, but may be required at Pod creation
time.

## Motivation

Consuming a service involves more than just connectivity. In addition to
coordinates to reach the service, credentials and non-secret configuration
parameters are typically needed to use the service.  The primitives for this
already exist, but a gap exists where loose coupling is desired: it should be
possible to inject pods with the information they need to use a service on a
service-by-service basis, without the pod authors having to incorporate the
information into every pod spec where it is needed.

## Constraints and Assumptions

1.  Future work might require new mechanisms to be made to work with existing
    controllers such as deployments and replicasets that create pods. Existing
    controllers that create pods should recreate their pods when a new Pod Injection
    Policy is added that would effect them.

## Use Cases

- As a user, I want to be able to provision a new pod
  without needing to know the application configuration primitives the
  services my pod will consume.
- As a cluster admin, I want specific configuration items of a service to be
  withheld visibly from a developer deploying a service, but not to block the
  developer from shipping.
- As an app developer, I want to provision a Cloud Spanner instance and then
  access it from within my Kubernetes cluster.
- As an app developer, I want the Cloud Spanner provisioning process to
  configure my Kubernetes cluster so the endpoints and credentials for my
  Cloud Spanner instance are implicitly injected into Pods matching a label
  selector (without me having to modify the PodSpec to add the specific
  Configmap/Secret containing the endpoint/credential data).


**Specific Example:**

1. Database Administrator provisions a MySQL service for their cluster.
2. Database Administrator creates secrets for the cluster containing the
   database name, username, and password.
3. Database Administrator creates a `PodPreset`  defining the database
   port as an environment variable, as well as the secrets. See
   [Examples](#examples) below for various examples.
4. Developer of an application can now label their pod with the specified
   `Selector` the Database Administrator tells them, and consume the MySQL
   database without needing to know any of the details from step 2 and 3.

### Summary

The use case we are targeting is to automatically inject into Pods the
information required to access non-Kubernetes-Services, such as accessing an
instances of Cloud Spanner. Accessing external services such as Cloud Spanner
may require the Pods to have specific credential and endpoint data.

Using a Pod Preset allows pod template authors to not have to explicitly
set information for every pod. This way authors of pod templates consuming a
specific service do not need to know all the details about that service.

### Prior Art

Internally for Kubernetes we already support accessing the Kubernetes api from
all Pods by injecting the credentials and endpoint data automatically - e.g.
injecting the serviceaccount credentials into a volume (via secret) using an
[admission controller](https://github.com/kubernetes/kubernetes/blob/97212f5b3a2961d0b58a20bdb6bda3ccfa159bd7/plugin/pkg/admission/serviceaccount/admission.go),
and injecting the Service endpoints into environment
variables. This is done without the Pod explicitly mounting the serviceaccount
secret.

### Objectives

The goal of this proposal is to generalize these capabilities so we can introduce
similar support for accessing Services running external to the Kubernetes cluster.
We can assume that an appropriate Secret and Configmap have already been created
as part of the provisioning process of the external service. The need then is to
provide a mechanism for injecting the Secret and Configmap into Pods automatically.

The [ExplicitServiceLinks proposal](https://github.com/kubernetes/community/pull/176),
will allow us to decouple where a Service's credential and endpoint information
is stored in the Kubernetes cluster from a Pod's intent to access that Service
(e.g. in declaring it wants to access a Service, a Pod is automatically injected
with the credential and endpoint data required to do so).

## Proposed Changes

### PodPreset API object

This resource is alpha. The policy itself is immutable. The API group will be
added to new group `settings` and the version is `v1alpha1`.

```go
// PodPreset is a policy resource that defines additional runtime
// requirements for a Pod.
type PodPreset struct {
    unversioned.TypeMeta
    ObjectMeta

    // +optional
    Spec PodPresetSpec
}

// PodPresetSpec is a description of a pod preset.
type PodPresetSpec struct {
    // Selector is a label query over a set of resources, in this case pods.
    // Required.
    Selector      unversioned.LabelSelector
    // Env defines the collection of EnvVar to inject into containers.
    // +optional
    Env           []EnvVar
    // EnvFrom defines the collection of EnvFromSource to inject into
    // containers.
    // +optional
    EnvFrom       []EnvFromSource
    // Volumes defines the collection of Volume to inject into the pod.
    // +optional
    Volumes       []Volume `json:omitempty`
    // VolumeMounts defines the collection of VolumeMount to inject into
    // containers.
    // +optional
    VolumeMounts  []VolumeMount
}
```

#### Validations

In order for the Pod Preset to be valid it must fulfill the
following constraints:

- The `Selector` field must be defined. This is how we know which pods
  to inject so therefore it is required and cannot be empty.
- The policy must define _at least_ 1 of `Env`, `EnvFrom`, or `Volumes` with
  corresponding `VolumeMounts`.
- If you define a `Volume`, it has to define a `VolumeMount`.
- For `Env`, `EnvFrom`, `Volumes`, and `VolumeMounts` all existing API
  validations are applied.

This resource will be immutable, if you want to change something you can delete
the old policy and recreate a new one. We can change this to be mutable in the
future but by disallowing it now, we will not break people in the future.

#### Conflicts

There are a number of edge conditions that might occur at the time of
injection. These are as follows:

- Merging lists with no conflicts: if a pod already has a `Volume`,
  `VolumeMount` or `EnvVar` defined **exactly** as defined in the
  PodPreset. No error will occur since they are the exact same. The
  motivation behind this is if services have no quite converted to using pod
  injection policies yet and have duplicated information and an error should
  obviously not be thrown if the items that need to be injected already exist
  and are exactly the same.
- Merging lists with conflicts: if a PIP redefines an `EnvVar` or a `Volume`,
  an event on the pod showing the error on the conflict will be thrown and
  nothing will be injected.
- Conflicts between `Env` and `EnvFrom`: this would throw an error with an
  event on the pod showing the error on the conflict. Nothing would be
  injected.

> **Note:** In the case of a conflict nothing will be injected. The entire
> policy is ignored and an event is thrown on the pod detailing the conflict.

### AdmissionControl Plug-in: PodPreset

The **PodPreset** plug-in introspects all incoming pod creation
requests and injects the pod based off a `Selector` with the desired
attributes, except when the [PodPreset Exclude Annotation](#podpreset-exclude-annotation)
is set to true.

For the initial alpha, the order of precedence for applying multiple
`PodPreset` specs is from oldest to newest. All Pod Injection
Policies in a namespace should be order agnostic; the order of application is
unspecified. Users should ensure that policies do not overlap.
However we can use merge keys to detect some of the conflicts that may occur.

This will not be enabled by default for all clusters, but once GA will be
a part of the set of strongly recommended plug-ins documented
[here](https://kubernetes.io/docs/admin/admission-controllers/#is-there-a-recommended-set-of-plug-ins-to-use).

**Why not an Initializer?**

This will be first implemented as an AdmissionControl plug-in then can be
converted to an Initializer once that is fully ready. The proposal for
Initializers can be found at [kubernetes/community#132](https://github.com/kubernetes/community/pull/132).

#### PodPreset Exclude Annotation
There may be instances where you wish for a pod to not be altered by any pod
preset mutations. For these events, one can add an annotation in the pod spec
of the form: `podpreset.admission.kubernetes.io/exclude: "true"`.

#### Behavior

This will modify the pod spec. The supported changes to
`Env`, `EnvFrom`, and `VolumeMounts` apply to the container spec for
all containers in the pod with the specified matching `Selector`. The
changes to `Volumes` apply to the pod spec for all pods matching `Selector`.

The resultant modified pod spec will be annotated to show that it was modified by
the `PodPreset`. This will be of the form
`podpreset.admission.kubernetes.io/podpreset-<pip name>": "<resource version>"`.

*Why modify all containers in a pod?*

Currently there is no concept of labels on specific containers in a pod which
would be necessary for per-container pod injections. We could add labels
for specific containers which would allow this and be the best solution to not
injecting all. Container labels have been discussed various times through
multiple issues and proposals, which all congregate to this thread on the
[kubernetes-sig-node mailing
list](https://groups.google.com/forum/#!topic/kubernetes-sig-node/gijxbYC7HT8).
In the future, even if container labels were added, we would need to be careful
about not making breaking changes to the current behavior.

Other solutions include basing the container to inject based off
matching its name to another field in the `PodPreset` spec, but
this would not scale well and would cause annoyance with configuration
management.

In the future we might question whether we need or want containers to express
that they expect injection. At this time we are deferring this issue.

## Examples

### Simple Pod Spec Example

This is a simple example to show how a Pod spec is modified by the Pod
Injection Policy.

**User submitted pod spec:**

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: website
  labels:
    app: website
    role: frontend
spec:
  containers:
    - name: website
      image: ecorp/website
      ports:
        - containerPort: 80
```

**Example Pod Preset:**

```yaml
kind: PodPreset
apiVersion: settings/v1alpha1
metadata:
  name: allow-database
  namespace: myns
spec:
  selector:
    matchLabels:
      role: frontend
  env:
    - name: DB_PORT
      value: 6379
  volumeMounts:
    - mountPath: /cache
      name: cache-volume
  volumes:
    - name: cache-volume
      emptyDir: {}
```

**Pod spec after admission controller:**

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: website
  labels:
    app: website
    role: frontend
  annotations:
    podpreset.admission.kubernetes.io/allow-database: "resource version"
spec:
  containers:
    - name: website
      image: ecorp/website
      volumeMounts:
        - mountPath: /cache
          name: cache-volume
      ports:
        - containerPort: 80
      env:
        - name: DB_PORT
          value: 6379
  volumes:
    - name: cache-volume
      emptyDir: {}
```

### Pod Spec with `ConfigMap` Example

This is an example to show how a Pod spec is modified by the Pod Injection
Policy that defines a `ConfigMap` for Environment Variables.

**User submitted pod spec:**

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: website
  labels:
    app: website
    role: frontend
spec:
  containers:
    - name: website
      image: ecorp/website
      ports:
        - containerPort: 80
```

**User submitted `ConfigMap`:**

```yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: etcd-env-config
data:
  number_of_members: "1"
  initial_cluster_state: new
  initial_cluster_token: DUMMY_ETCD_INITIAL_CLUSTER_TOKEN
  discovery_token: DUMMY_ETCD_DISCOVERY_TOKEN
  discovery_url: http://etcd_discovery:2379
  etcdctl_peers: http://etcd:2379
  duplicate_key: FROM_CONFIG_MAP
  REPLACE_ME: "a value"
```

**Example Pod Preset:**

```yaml
kind: PodPreset
apiVersion: settings/v1alpha1
metadata:
  name: allow-database
  namespace: myns
spec:
  selector:
    matchLabels:
      role: frontend
  env:
    - name: DB_PORT
      value: 6379
    - name: duplicate_key
      value: FROM_ENV
    - name: expansion
      value: $(REPLACE_ME)
  envFrom:
    - configMapRef:
        name: etcd-env-config
  volumeMounts:
    - mountPath: /cache
      name: cache-volume
    - mountPath: /etc/app/config.json
      readOnly: true
      name: secret-volume
  volumes:
    - name: cache-volume
      emptyDir: {}
    - name: secret-volume
      secretName: config-details
```

**Pod spec after admission controller:**

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: website
  labels:
    app: website
    role: frontend
  annotations:
    podpreset.admission.kubernetes.io/allow-database: "resource version"
spec:
  containers:
    - name: website
      image: ecorp/website
      volumeMounts:
        - mountPath: /cache
          name: cache-volume
        - mountPath: /etc/app/config.json
          readOnly: true
          name: secret-volume
      ports:
        - containerPort: 80
      env:
        - name: DB_PORT
          value: 6379
        - name: duplicate_key
          value: FROM_ENV
        - name: expansion
          value: $(REPLACE_ME)
      envFrom:
        - configMapRef:
          name: etcd-env-config
  volumes:
    - name: cache-volume
      emptyDir: {}
    - name: secret-volume
      secretName: config-details
```

### ReplicaSet with Pod Spec Example

The following example shows that only the pod spec is modified by the Pod
Injection Policy.

**User submitted ReplicaSet:**

```yaml
apiVersion: settings/v1alpha1
kind: ReplicaSet
metadata:
  name: frontend
spec:
  replicas: 3
  selector:
    matchLabels:
      tier: frontend
    matchExpressions:
      - {key: tier, operator: In, values: [frontend]}
  template:
    metadata:
      labels:
        app: guestbook
        tier: frontend
    spec:
      containers:
      - name: php-redis
        image: gcr.io/google_samples/gb-frontend:v3
        resources:
          requests:
            cpu: 100m
            memory: 100Mi
        env:
          - name: GET_HOSTS_FROM
            value: dns
        ports:
          - containerPort: 80
```

**Example Pod Preset:**

```yaml
kind: PodPreset
apiVersion: settings/v1alpha1
metadata:
  name: allow-database
  namespace: myns
spec:
  selector:
    matchLabels:
      tier: frontend
  env:
    - name: DB_PORT
      value: 6379
  volumeMounts:
    - mountPath: /cache
      name: cache-volume
  volumes:
    - name: cache-volume
      emptyDir: {}
```

**Pod spec after admission controller:**

```yaml
kind: Pod
  metadata:
    labels:
      app: guestbook
      tier: frontend
    annotations:
    podpreset.admission.kubernetes.io/allow-database: "resource version"
  spec:
    containers:
      - name: php-redis
        image: gcr.io/google_samples/gb-frontend:v3
        resources:
          requests:
            cpu: 100m
            memory: 100Mi
        volumeMounts:
          - mountPath: /cache
            name: cache-volume
        env:
          - name: GET_HOSTS_FROM
            value: dns
          - name: DB_PORT
            value: 6379
        ports:
          - containerPort: 80
    volumes:
      - name: cache-volume
        emptyDir: {}
```

### Multiple PodPreset Example

This is an example to show how a Pod spec is modified by multiple Pod
Injection Policies.

**User submitted pod spec:**

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: website
  labels:
    app: website
    role: frontend
spec:
  containers:
    - name: website
      image: ecorp/website
      ports:
        - containerPort: 80
```

**Example Pod Preset:**

```yaml
kind: PodPreset
apiVersion: settings/v1alpha1
metadata:
  name: allow-database
  namespace: myns
spec:
  selector:
    matchLabels:
      role: frontend
  env:
    - name: DB_PORT
      value: 6379
  volumeMounts:
    - mountPath: /cache
      name: cache-volume
  volumes:
    - name: cache-volume
      emptyDir: {}
```

**Another Pod Preset:**

```yaml
kind: PodPreset
apiVersion: settings/v1alpha1
metadata:
  name: proxy
  namespace: myns
spec:
  selector:
    matchLabels:
      role: frontend
  volumeMounts:
    - mountPath: /etc/proxy/configs
      name: proxy-volume
  volumes:
    - name: proxy-volume
      emptyDir: {}
```

**Pod spec after admission controller:**

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: website
  labels:
    app: website
    role: frontend
  annotations:
    podpreset.admission.kubernetes.io/allow-database: "resource version"
    podpreset.admission.kubernetes.io/proxy: "resource version"
spec:
  containers:
    - name: website
      image: ecorp/website
      volumeMounts:
        - mountPath: /cache
          name: cache-volume
        - mountPath: /etc/proxy/configs
          name: proxy-volume
      ports:
        - containerPort: 80
      env:
        - name: DB_PORT
          value: 6379
  volumes:
    - name: cache-volume
      emptyDir: {}
    - name: proxy-volume
      emptyDir: {}
```

### Conflict Example

This is a example to show how a Pod spec is not modified by the Pod Injection
Policy when there is a conflict.

**User submitted pod spec:**

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: website
  labels:
    app: website
    role: frontend
spec:
  containers:
    - name: website
      image: ecorp/website
      volumeMounts:
        - mountPath: /cache
          name: cache-volume
      ports:
  volumes:
    - name: cache-volume
      emptyDir: {}
        - containerPort: 80
```

**Example Pod Preset:**

```yaml
kind: PodPreset
apiVersion: settings/v1alpha1
metadata:
  name: allow-database
  namespace: myns
spec:
  selector:
    matchLabels:
      role: frontend
  env:
    - name: DB_PORT
      value: 6379
  volumeMounts:
    - mountPath: /cache
      name: other-volume
  volumes:
    - name: other-volume
      emptyDir: {}
```

**Pod spec after admission controller will not change because of the conflict:**

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: website
  labels:
    app: website
    role: frontend
spec:
  containers:
    - name: website
      image: ecorp/website
      volumeMounts:
        - mountPath: /cache
          name: cache-volume
      ports:
  volumes:
    - name: cache-volume
      emptyDir: {}
        - containerPort: 80
```

**If we run `kubectl describe...` we can see the event:**

```
$ kubectl describe ...
....
Events:
  FirstSeen             LastSeen            Count   From                    SubobjectPath               Reason      Message
  Tue, 07 Feb 2017 16:56:12 -0700   Tue, 07 Feb 2017 16:56:12 -0700 1   {podpreset.admission.kubernetes.io/allow-database }    conflict  Conflict on pod preset. Duplicate mountPath /cache.
```
