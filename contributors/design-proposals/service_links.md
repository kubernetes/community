# Explicit service links

## Abstract

Proposes an API for explicit links to services in the pod spec.

## Motivation

Currently, all pods are automatically injected with docker-link-style
environment variables for each service in their namespaces.  Making service
dependencies explicit in the PodSpec has a number of advantages:

1.  Reduce accidental coupling
2.  Reduce the number of IP tables rules that have to be created by the
    kube-proxy
3.  Identify relationships between pods and services, allowing user interfaces
    to display additional information for end users.

## Constraints and Assumptions

1.  The old behavior must remain the default in the v1 API
2.  It must be possible to migrate off the default to using explicit links
3.  This proposal does not capture any possible future changes to the service
    proxy
4.  This proposal does not capture any concerns around 

## Use Cases

1.  As a user, I want to be able to explicitly define which pods are injected
    with information about a service so that I can ensure my pods receive
    information only about services they actually need to use
2.  As a user, I want to control how services are injected into pods and
    containers
3.  As a user, I want to prevent my pods from starting if services they depend
    on do not exist

### Explicit links between pods and services

Currently in Kubernetes a pod is injected with docker- link style environment
variables and simpler `{SERVICE_NAME}_SERVICE_HOST}` and
`{SERVICE_NAME}_SERVICE_PORT`
[environment variables](http://kubernetes.io/docs/user-guide/services/#environment-variables)
for every service in its namespace.  This is disadvantageous for a number of
reasons:

1.  It is not possible to tell from the pod's spec which services it consumes
2.  The proxy must generate rules on a node for every service in every namespace
3.  There is no way to prevent accidental coupling

Adding explicit links to services addresses all three of the above concerns,
once the existing automatic injection behavior is turned off:

1.  Explicit links will appear in the pod's spec
2.  The proxy can be modified to generate rules only for the services which are
    explicitly consumed on a host

### Controlling how pods are injected with information

It is common for legacy applications to expect custom environment variable names
that do not fit the docker-link style that is generated today. Custom
environment variables for service links are possible, but they are cumbersome
and lead to duplication in the environment.  To get custom environment vars for
service links, a pod must define the custom variables and use variable expansion
to inject the value of the default variable.

Explicit links offer an opportunity to ease the experience and make custom
variables easier to reason about.  Instead of referencing variables that will be
present at runtime, but are not present otherwise in the pod spec, a pod author
could simply specify a format or prefix for the variables to be created.

### Gating pod start on existence of services

Today, the `kubectl` creates resources in a directory in an arbitrary order
and concurrently.  This means that no assumption about creation order of API
resources is a safe one to make.  Therefore, it is desirable to express that a
pod should not be started until a service it depends on exists.

Additionally, a user might want to create a partial set of resources from a
template or chart and have automation asynchronously supply other resources.
Having the ability to describe a dependency relationship between pods and
services facilitates this.

When assessing what the decision plane should be for allowing a pod with a
dependency on a service to start, it is natural to begin thinking of readiness
and availability of the services.  However, these assessments are inherently
racey and ephemeral.  A service could have many endpoints when the pod was
started, and by the time one of the containers in the pod is running, zero
endpoints might be available.  Container processes should be robust in the
face of services becoming unreachable, but gating on existence of a service is
useful because containers cannot be injected with environment variables after
being started.


## Analysis

Determining the right API is a complex topic that warrants some analysis.

### What degree of specificity is desired?

The desired degree of specificity for an explicit service link is important to
consider.  There is a continuum of options of different complexities.  On the
simple end of the spectrum is an API that manifests environment variables in the
exact same way as the current implicit mechanism, but allows a service to be
explicitly referenced.  On the complex end of the spectrum is an API that allows
individual aspects of a service to be projected into individual environment
variables.

The right API is probably not either of those options, but it is useful to
consider the pros and cons of both as a way to discover the right API.

#### Pro/con: Simple API

Pros:

1.  Very small practical and conceptual delta from current behavior
2.  Users can employ the same strategies for custom env vars
3.  Easiest to explain
4.  Easiest to support

Cons:

1.  Inflexible
2.  Users have the same usability issues when custom env var names are required
3.  Does not extend well to services in other namespaces

#### Pro/con: Most complex API

Pros:

1.  Offers highest degree of flexibility
2.  Custom env var names are cheap at fine-grained level of control

Cons:

1.  Most difficult to explain and reason about
2.  Most difficult to support
3.  Higher implementation cost
4.  Large conceptual delta
5.  Likely to add significant complexity to pod spec
6.  Likely to be very verbose

#### Lessons Learned from Prior APIs

One fundamental lesson of Kubernetes API design is that we must make backward
compatible API changes except in cases of special exception.  This means that we
need to consider backward compatibility when making new API changes.  As a
requirement, backward compatibility imposes a burden that you have to live with
the APIs you release.  This means that we should attempt to be conservative with
new API features and avoid unnecessary complexity.  Therefore, a simpler API is
advantageous because complexity can be added later **if desired and warranted**.

#### Middleground

A conservative middleground might be an API that has behavior extremely close to
the current implicit behavior but allows specifying custom variable names.  This
approach has the advantage of being close to what users already expect, and can
be implemented with a fairly simple API.

### Non-existent Services

When a service referenced using this API does not exist, the Kubelet should
avoid starting the pod.

## Proposed Changes

### API changes

#### Expressing dependencies on services

```go
type PodSpec {
  // other fields omitted

  // ServiceDependencies is a list of services in this pod's namespace that the
  // pod depends on in order to start.
  ServiceDependencies []LocalObjectReference `json:"serviceDependencies,omitempty"`
}
```

#### Environment variable injection

Pull request [#148](https://github.com/kubernetes/community/pull/148)
proposes an `EnvFrom` addition to the PodSpec to handle creating multiple
environment variables from a single subject.  This API surface is ideal for the
service use-cases, so we will extend it here:

```go
type Container struct {
	// other fields omitted

	EnvFrom []EnvFromSource `json:"envFrom,omitempty"`
}

type EnvVarFromSource struct {
	// other fields omitted

	// Prefix is an optional prefix for generated environment variables
	Prefix string `json:"prefix,omitempty"`

	// Can be refactored to ObjectReference at a later time while
	// maintaining backward compatibility.
	Service *ServiceEnvSource `json:"service,omitempty"`
}

// ServiceEnvVarSource selects a Service to populate environment
// variables with.
//
// The following environment variables will be generated for the
// target service:
//
// <service name>_SERVICE_HOST
// <service name>_SERVICE_PORT
// <service name>_PORT=tcp://10.0.0.11:6379
// <service name>_PORT_6379_TCP=tcp://10.0.0.11:6379
// <service name>_PORT_6379_TCP_PROTO=tcp
// <service name>_PORT_6379_TCP_PORT=6379
// <service name>_PORT_6379_TCP_ADDR=10.0.0.11
type ServiceEnvSource struct {
  // The Service to generate environment variables for.
  LocalObjectReference
}
```

When a pod uses this API feature, it indicates that the automatic injection
behavior is not desired for that pod.  No automatic injection will occur when
any container in a pod uses an explicit link to a service.

Notes:

1. The `kubernetes` and `kubernetes-ro` services will always be injected,
   regardless of the state of the pod's spec.
2. The `ServiceDependencies` and `EnvFrom` API features are orthogonal
   and can be used independently

#### Validations

Validations should be added as follows:

1. The `Name` field of `ServiceEnvSource` must be set
2. The names of services listed in `ServiceDependencies` must be valid
   Service names

There are no expectations that if a container consumes a service in
environment variables via `EnvFrom`, that this service must also be named as a
dependency or vice versa.  For example, a pod author may want to wait until a
service exists for the pod to start, but may not care at all about environment
variables.

### Kubelet Changes

The Kubelet must be modified to:

1. Support creation of environment variables via the new API
2. Disable the automatic injection for a pod whenever an `EnvFrom` for a
   service appears in the pod's spec
3. Start the containers in a pod only when all of the pod's
   `ServiceDependencies` exist
4. Create an event for the pod when pod start is delayed due to waiting for a
   service to exist

## Examples

### No custom variable names

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: simple-service-reference-example
spec:
  containers:
  - image: busybox
    name: example-container
    envFrom:
    - service:
        name: redis-master
```

Yields the environment variables:

```
REDIS_MASTER_SERVICE_HOST=10.0.0.11
REDIS_MASTER_SERVICE_PORT=6379
REDIS_MASTER_PORT=tcp://10.0.0.11:6379
REDIS_MASTER_PORT_6379_TCP=tcp://10.0.0.11:6379
REDIS_MASTER_PORT_6379_TCP_PROTO=tcp
REDIS_MASTER_PORT_6379_TCP_PORT=6379
REDIS_MASTER_PORT_6379_TCP_ADDR=10.0.0.11
```

### Custom variable names with prefix

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: simple-service-reference-example
spec:
  containers:
  - image: busybox
    name: example-container
    envFrom:
    - prefix: "PREFIX_"
      service:
        name: redis-master
```

Yields the environment variables:

```
PREFIX_REDIS_MASTER_SERVICE_HOST=10.0.0.11
PREFIX_REDIS_MASTER_SERVICE_PORT=6379
PREFIX_REDIS_MASTER_PORT=tcp://10.0.0.11:6379
PREFIX_REDIS_MASTER_PORT_6379_TCP=tcp://10.0.0.11:6379
PREFIX_REDIS_MASTER_PORT_6379_TCP_PROTO=tcp
PREFIX_REDIS_MASTER_PORT_6379_TCP_PORT=6379
PREFIX_REDIS_MASTER_PORT_6379_TCP_ADDR=10.0.0.11
```
