# Source code tour guide

## Architecture

<details>

### Kubernetes architecture

![k8s](https://raw.githubusercontent.com/qas/diagrams/master/k8s/architecture_detailed.svg)
(Click the image to view links)

[architecture details](https://github.com/kubernetes/community/blob/master/contributors/design-proposals/architecture/architecture.md)

[architecture roadmap](https://github.com/kubernetes/community/blob/master/contributors/design-proposals/architecture/architectural-roadmap.md)
### Concepts

**Kubernetes components**
- control plane (master)
	- kube-apiserver node
	- etcd
- user plane
	- worker nodes
</details>

## API Machinery

<details>

### Codebase

- Super high level

  - things favor a modular packages and reusable code approach
    - **apimachinery** are tools to create api runtimes, things like working with JSON and YAML
    - **apiserver** is a library for making RESTful servers that store data in etcd
  - a user request flow
    - a happy user like you sitting in front of a computer > [cmd] > [pkg] > [some apiserver] > [registry] > [storage] > [etcd]
  - a state management and state handling flow
    - Controllers are for running logic for resources. Controllers watch etcd state updates and handle them.

- Overview
  - uses Go as the programming language
  - implements RESTful concepts at the core of working with object storage and everything is stored in etcd, a.k.a. the only database K8s uses
  - usually a `doc.go` file should include a description of the the module's purpose
  - `/cmd` houses the entrypoints
  - `/pkg` houses things that not only includes reusable packages, but also "internal" code that `/cmd` commands call. The code in here also uses modules found in `/staging/src/k8s`.
  - `/staging/src/k8s` external repos usually with their own SIG (special interest groups, which are basically a subset of a larger group of people or community) for modules, things like the following
    - tools for parsing JSON/YAML using `apimachinery` library primarily for building JSON/YAML runtimes
    - building api servers using the `apiserver` library
    - logic for storage and registry
    - logic for commands
    - and even dedicated example apps using the tools and libraries used by K8s
      - sample-apiserver (data handling stuff)
      - sample-cli-plugin
      - sample-controller (logic handling stuff)

- Kubernetes architecture & design
  - `/staging` hosts code that is split into separate repos
    - **Tools & libraries**
      - **storage** handles etcd integration
      - **registry** allows RESTful style model for interacting with etcd storage
        - [store.go](https://sourcegraph.com/github.com/kubernetes/kubernetes@master/-/blob/staging/src/k8s.io/apiserver/pkg/registry/generic/registry/store.go) looks like it defines the storage manipulation model with RESTful support
      - **apimachinery** handles runtime logic, i.e. encoding/decoding
      - **apiserver** library to create a RESTful server for a given resource

- Kubernetes entrypoints
  - `/cmd` hosts all entrypoints and uses hosted packages in `/pkg`

- Kubernetes logic
  - `/pkg/apis` hosts api resources made using the **apiserver** library
  - `/pkg/controller` hosts logic for K8s resources
  - `/pkg/registry` hosts db strategies using the `/staging/src/k8s.io/apiserver/registry` package

- Kubernetes database
  - etcd is the only database used related files are usually found in the following places
    - **apiserver** is where it's actually integrated
      - `/staging/src/k8s.io/apiserver/registry`
      - `/staging/src/k8s.io/apiserver/storage`
    - usage is in
      - `/pkg`
      - `/staging/src/k8s.io/apiextensions-apiserver`
      - `/staging/src/k8s.io/sample-apiserver` nice reusable example to look at how **apiserver** integration works
    - other notes
      - `/staging/src/k8s.io/apimachinery` is used to serialize/deserialize the registry data

- Kubernetes dependencies
  - [spf13/cobra](https://github.com/spf13/cobra) - for building command line apps like **kubectl** found in `/cmd`

### Kubernetes' framework structures and examples

<details>
<summary>
	<a href="https://sourcegraph.com/github.com/kubernetes/kubernetes@master/-/tree/staging/src/k8s.io/apiserver/pkg/server">HTTP server framework</a>
</summary>
<details>
	<summary>
	<a href="https://sourcegraph.com/github.com/kubernetes/kubernetes@master/-/blob/staging/src/k8s.io/apiserver/pkg/server/genericapiserver.go#L92">GenericAPIServer</a> - GenericAPIServer contains state for a Kubernetes cluster api server.
	</summary>
</details>
</details>

<details>
<summary>
<a href="https://sourcegraph.com/github.com/kubernetes/kubernetes@master/-/blob/staging/src/k8s.io/apiserver/pkg/storage/storagebackend/config.go#L50">storage config for creating storage backend</a>
</summary>
</details>

<details>
<summary>
<a href="https://sourcegraph.com/github.com/kubernetes/kubernetes@master/-/blob/staging/src/k8s.io/apiserver/pkg/registry/generic/registry/store.go">generic registry store</a>
</summary>
</details>

<details>
<summary>
<a href="">runtime scheme to serializing/deserializing API objects</a>
</summary>
</details>

<details>
<summary>
<a href="https://sourcegraph.com/github.com/kubernetes/kubernetes@master/-/blob/staging/src/k8s.io/apimachinery/pkg/runtime/types.go#L36">runtime type TypeMeta</a>
</summary>
</details>

### Kubernetes specific resource examples

<details>
<summary>
<a href="https://sourcegraph.com/github.com/kubernetes/kubernetes@master/-/blob/pkg/controller/deployment/deployment_controller.go#L68">deployment_controller.go</a>
</summary>
</details>

<details>
<summary>
<a href="https://github.com/kubernetes/kubernetes/blob/master/pkg/registry/core/pod/storage/storage.go#L78">example usage of genericregistry.Store</a>
</summary>
</details>

<details>
<summary>
<a href="https://sourcegraph.com/github.com/kubernetes/kubernetes@master/-/blob/pkg/registry/core/pod/storage/storage.go#L74">PodStorage</a>
</summary>
</details>
</details>

## Apps
TODO

## Auth
TODO

## Autoscaling
TODO

## CLI
TODO

## Cluster Lifecycle
TODO

## Instrumentation Tools
TODO

## Network
TODO

## Node
TODO

## Scheduling
TODO

## Storage
TODO

## Testing
TODO

## UI
TODO
