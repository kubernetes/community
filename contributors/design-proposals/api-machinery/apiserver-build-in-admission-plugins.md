# Build some Admission Controllers into the Generic API server library

**Related PR:**

| Topic | Link |
| ----- | ---- |
| Admission Control | https://git.k8s.io/community/contributors/design-proposals/api-machinery/admission_control.md |

## Introduction

An admission controller is a piece of code that intercepts requests to the Kubernetes API - think a middleware.
The API server lets you have a whole chain of them. Each is run in sequence before a request is accepted
into the cluster. If any of the plugins in the sequence rejects the request, the entire request is rejected
immediately and an error is returned to the user.

Many features in Kubernetes require an admission control plugin to be enabled in order to properly support the feature.
In fact in the [documentation](https://kubernetes.io/docs/admin/admission-controllers/#is-there-a-recommended-set-of-plug-ins-to-use) you will find
a recommended set of them to use.

At the moment admission controllers are implemented as plugins and they have to be compiled into the
final binary in order to be used at a later time. Some even require an access to cache, an authorizer etc.
This is where an admission plugin initializer kicks in. An admission plugin initializer is used to pass additional
configuration and runtime references to a cache, a client and an authorizer.

To streamline the process of adding new plugins especially for aggregated API servers we would like to build some plugins
into the generic API server library and provide a plugin initializer. While anyone can author and register one, having a known set of
provided references let's people focus on what they need their admission plugin to do instead of paying attention to wiring.

## Implementation

The first step would involve creating a "standard" plugin initializer that would be part of the
generic API server. It would use kubeconfig to populate
[external clients](https://git.k8s.io/kubernetes/pkg/kubeapiserver/admission/initializer.go#L29)
and [external informers](https://git.k8s.io/kubernetes/pkg/kubeapiserver/admission/initializer.go#L35).
By default for servers that would be run on the kubernetes cluster in-cluster config would be used.
The standard initializer would also provide a client config for connecting to the core kube-apiserver.
Some API servers might be started as static pods, which don't have in-cluster configs.
In that case the config could be easily populated form the file.

The second step would be to move some plugins from [admission pkg](https://git.k8s.io/kubernetes/plugin/pkg/admission)
to the generic API server library. Some admission plugins are used to ensure consistent user expectations.
These plugins should be moved. One example is the Namespace Lifecycle plugin which prevents users
from creating resources in non-existent namespaces.

*Note*:
For loading in-cluster configuration [visit](https://git.k8s.io/kubernetes/staging/src/k8s.io/client-go/examples/in-cluster-client-configuration/main.go)
 For loading the configuration directly from a file [visit](https://git.k8s.io/kubernetes/staging/src/k8s.io/client-go/examples/out-of-cluster-client-configuration/main.go)

## How to add an admission plugin ?
 At this point adding an admission plugin is very simple and boils down to performing the
following series of steps:
 1. Write an admission plugin
 2. Register the plugin
 3. Reference the plugin in the admission chain

## An example
The sample apiserver provides an example admission plugin that makes meaningful use of the "standard" plugin initializer.
The admission plugin ensures that a resource name is not on the list of banned names.
The source code of the plugin can be found [here](https://github.com/kubernetes/kubernetes/blob/2f00e6d72c9d58fe3edc3488a91948cf4bfcc6d9/staging/src/k8s.io/sample-apiserver/pkg/admission/plugin/banflunder/admission.go).

Having the plugin, the next step is the registration. [AdmissionOptions](https://github.com/kubernetes/kubernetes/blob/2f00e6d72c9d58fe3edc3488a91948cf4bfcc6d9/staging/src/k8s.io/apiserver/pkg/server/options/admission.go)
provides two important things. Firstly it exposes [a register](https://github.com/kubernetes/kubernetes/blob/2f00e6d72c9d58fe3edc3488a91948cf4bfcc6d9/staging/src/k8s.io/apiserver/pkg/server/options/admission.go#L43)
under which all admission plugins are registered. In fact, that's exactly what the [Register](https://github.com/kubernetes/kubernetes/blob/2f00e6d72c9d58fe3edc3488a91948cf4bfcc6d9/staging/src/k8s.io/sample-apiserver/pkg/admission/plugin/banflunder/admission.go#L33)
method does from our example admission plugin. It accepts a global registry as a parameter and then simply registers itself in that registry.
Secondly, it adds an admission chain to the server configuration via [ApplyTo](https://github.com/kubernetes/kubernetes/blob/2f00e6d72c9d58fe3edc3488a91948cf4bfcc6d9/staging/src/k8s.io/apiserver/pkg/server/options/admission.go#L66) method.
The method accepts optional parameters in the form of `pluginInitalizers`. This is useful when admission plugins need custom configuration that is not provided by the generic initializer.

The following code has been extracted from the sample server and illustrates how to register and wire an admission plugin:

```go
  // register admission plugins
  banflunder.Register(o.Admission.Plugins)

  // create custom plugin initializer
  informerFactory := informers.NewSharedInformerFactory(client, serverConfig.LoopbackClientConfig.Timeout)
  admissionInitializer, _ := wardleinitializer.New(informerFactory)

  // add admission chain to the server configuration
  o.Admission.ApplyTo(serverConfig, admissionInitializer)
```
