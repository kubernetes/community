# Optional service links

Status: Pending

Version: Alpha | Beta | GA

Implementation Owner: @kongslund

## Motivation

Today, a list of all services that were running when a pod's containers are created is automatically injected to those containers as environment variables matching the syntax of Docker links. There is no way to disable this.

Docker links have long been considered as a [deprecated legacy feature](https://docs.docker.com/engine/userguide/networking/default_network/dockerlinks/) of Docker since the introduction of networks and DNS. Likewise, in Kubernetes, DNS is to be preferred over service links.

Possible issues with injected service links are

* Accidental coupling.
* Incompatibilities with container images that no longer utilize service links and explicitly fail at startup time if certain service links are defined.
* Performance penalty in starting up pods [for namespaces with many services](https://github.com/kubernetes/kubernetes/issues/1768#issuecomment-330778184).

## Proposal

Make it possible for a user to disable injection of service links into containers of a pod by adding a disable flag to the pod's spec. Make the default value false in order to stay backwards compatible with the v1 API. Make an exception for the `kubernetes` service in the master namespace so that it will always get injected.

## User Experience

### Use Cases

* As a user, I want to be able to disable service link injection since the injected environment variables interfer with a Docker image that I am trying to run on Kubernetes
* As a user, I want to be able to disable service link injection since I don't need it and it takes increasingly longer time to start pods as services are added to the namespace.

## Implementation

`PodSpec` is extended with an additional field, `disableServiceLinks` of type boolean. Default value is false.

In `kubelet_pods.go`, the value of that field is passed along to the function `getServiceEnvVarMap` where it is used to decide which selector should be used for the `serviceLister`. Current behavior is `labels.Everything()`. In case `disableServiceLinks` is true then only the `kubernetes` service in the `kl.masterServiceNamespace` should be injected.

```
func (kl *Kubelet) getServiceEnvVarMap(ns string, disableServiceLinks bool) (map[string]string, error) {
  ...decide on selector
}
```

### Client/Server Backwards/Forwards compatibility

Pods that do not have the field set will assume a value of false.

## Alternatives considered

An alternative is to add support for explicit service links, e.g. by applying a label selector map with a default behavior of including everything.

Making service links explicit has been discussed in the [Explicit service links](https://github.com/pmorie/community/blob/6239773beb623c0dafc768be8372e5daf605aab0/contributors/design-proposals/service_links.md) design proposal without being merged in.
