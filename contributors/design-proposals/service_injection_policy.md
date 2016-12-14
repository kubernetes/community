# Service Injection Policy

## Abstract

Describes a policy resource that allows loose coupling from services to the pods
that consume them.

## Motivation

Consuming a service involves more than just connectivity.  In addition to
coordinates to reach the service, credentials and non-secret configuration
parameters are typically needed to use the service.  The primitives for this
already exist, but a gap exists where loose coupling is desired: it should be
possible to inject pods with the information they need to use a service on a
service-by-service basis, without the pod authors having to incorporate the
information into every pod spec where it is needed.

## Constraints and Assumptions

1.  New mechanisms must be made to work with controllers such as deployments and
    replicasets that create pods

## Use Cases

1.  As a user, I want to be able to describe a way that pods should be injected
    with the information to consume a particular service in a loosely-coupled
    way, so that I can concisely model the information about how the service
    should be consumed without altering every consuming pod spec

<!--
2.  As a user, I want a controller that manage pods to create a new generation
    of pods when the pods that controller's pods should be injected with
    information about a new set of services
-->

### Loose coupling between services and their consumers



## Proposed Changes

### ServiceInjectionPolicy API object

```go
type ServiceInjectionPolicy struct {
	unversioned.TypeMeta
	ObjectMeta

	Spec ServiceInjectionPolicySpec
}

type ServiceInjectionPolicySpec struct {
	LabelSelector *unversioned.LabelSelector
	Env           []EnvVar
	EnvFrom       []EnvVarFrom
	Volumes       []Volume
	VolumeMounts  []VolumeMount
}
```

## Examples

```yaml

```