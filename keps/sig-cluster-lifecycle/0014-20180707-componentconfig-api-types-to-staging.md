---
kep-number: 17
title: Moving ComponentConfig API types to staging repos
status: provisional
authors:
  - "@luxas"
  - "@sttts"
owning-sig: sig-cluster-lifecycle
participating-sigs:
  - sig-api-machinery
  - sig-node
  - sig-network
  - sig-scheduling
  - sig-cloud-provider
reviewers:
  - "@thockin"
  - "@liggitt"
  - "@wojtek-t"
  - "@stewart-yu"
  - "@dixudx"
approvers:
  - "@thockin"
  - "@jbeda"
  - "@deads2k"
editor:
  name: "@luxas"
creation-date: 2018-07-07
last-updated: 2018-07-07
---

# Moving ComponentConfig API types to staging repos

**How we can start supporting reading versioned configuration for all our components after a code move for ComponentConfig to staging**

## Table of Contents

 * [Moving ComponentConfig API types to staging repos](#moving-componentconfig-api-types-to-staging-repos)
    * [Table of Contents](#table-of-contents)
    * [Summary](#summary)
       * [The current state of the world](#the-current-state-of-the-world)
          * [Current kubelet](#current-kubelet)
          * [Current kube-proxy](#current-kube-proxy)
          * [Current kube-scheduler](#current-kube-scheduler)
          * [Current kube-controller-manager](#current-kube-controller-manager)
          * [Current kube-apiserver](#current-kube-apiserver)
          * [Current cloud-controller-manager](#current-cloud-controller-manager)
       * [Goals](#goals)
       * [Non-goals](#non-goals)
       * [Related proposals and further references](#related-proposals-and-further-references)
    * [Proposal](#proposal)
       * [Migration strategy per component or k8s.io repo](#migration-strategy-per-component-or-k8sio-repo)
          * [k8s.io/apimachinery changes](#k8sioapimachinery-changes)
          * [k8s.io/apiserver changes](#k8sioapiserver-changes)
          * [kubelet changes](#kubelet-changes)
          * [kube-proxy changes](#kube-proxy-changes)
          * [kube-scheduler changes](#kube-scheduler-changes)
          * [k8s.io/controller-manager changes](#k8siocontroller-manager-changes)
          * [kube-controller-manager changes](#kube-controller-manager-changes)
          * [cloud-controller-manager changes](#cloud-controller-manager-changes)
          * [kube-apiserver changes](#kube-apiserver-changes)
       * [Timeframe and Implementation Order](#timeframe-and-implementation-order)
       * [OWNERS files for new packages and repos](#owners-files-for-new-packages-and-repos)

## Summary

Currently all ComponentConfiguration API types are in the core Kubernetes repo. This makes them practically inaccessible for any third-party tool. With more and more generated code being removed from the core Kubernetes repo, vendoring gets even more complicated. Last but not least, efforts to move out kubeadm from the core repo are blocked by this.

This KEP is about creating new staging repos, `k8s.io/{component}`, which will host the external
types of the core components’ ComponentConfig in a top-level `config/` package. Internal types will *eventually* be stored in
`k8s.io/{component}/pkg/apis/config` (but a non-goal for this KEP). Shared types will go to `k8s.io/{apimachinery,apiserver,controller-manager}/pkg/apis/config`.

### The current state of the world

#### Current kubelet

* **Package**: [k8s.io/kubernetes/pkg/kubelet/apis/kubeletconfig](https://github.com/kubernetes/kubernetes/blob/release-1.11/pkg/kubelet/apis/kubeletconfig/types.go)
* **GroupVersionKind:** `kubelet.config.k8s.io/v1beta.KubeletConfiguration`
* **Supports** reading **config from file** with **flag precedence**, **well-tested**.

#### Current kube-proxy

* **Package**: [k8s.io/kubernetes/pkg/proxy/apis/kubeproxyconfig](https://github.com/kubernetes/kubernetes/blob/release-1.11/pkg/proxy/apis/kubeproxyconfig/types.go)
* **GroupVersionKind**: `kubeproxy.config.k8s.io/v1alpha1.KubeProxyConfiguration`
* **Supports** reading **config from file**, **without flag precedence**, **not tested**.
* This API group has its own copy of `ClientConnectionConfiguration` instead of a shared type.

#### Current kube-scheduler

* **Package**: [k8s.io/kubernetes/pkg/apis/componentconfig](https://github.com/kubernetes/kubernetes/blob/release-1.11/pkg/apis/componentconfig/types.go)
* **GroupVersionKind**: `componentconfig/v1alpha1.KubeSchedulerConfiguration`
* **Supports** reading **config from file**, **without flag precedence**, **not tested**
* This API group has its own copies of `ClientConnectionConfiguration` & `LeaderElectionConfiguration` instead of shared types.

#### Current kube-controller-manager

* **Package**: [k8s.io/kubernetes/pkg/apis/componentconfig](https://github.com/kubernetes/kubernetes/blob/release-1.11/pkg/apis/componentconfig/types.go)
* **GroupVersionKind**: `componentconfig/v1alpha1.KubeControllerManagerConfiguration`
* **No support for config from file**
* This API group has its own copies of `ClientConnectionConfiguration` & `LeaderElectionConfiguration` instead of shared types.

#### Current kube-apiserver

* **Doesn’t expose component configuration anywhere**
* **No support for config from file**
* The most similar thing to componentconfig for the API server is the `ServerRunOptions` struct in
  [k8s.io/kubernetes/cmd/kube-apiserver/app/options/options.go](https://github.com/kubernetes/kubernetes/blob/release-1.11/cmd/kube-apiserver/app/options/options.go)

#### Current cloud-controller-manager

* **Package**: [k8s.io/kubernetes/pkg/apis/componentconfig](https://github.com/kubernetes/kubernetes/blob/release-1.11/pkg/apis/componentconfig/types.go)
* **GroupVersionKind**: `componentconfig/v1alpha1.CloudControllerManagerConfiguration`
* **No support for config from file**
* This API group has its own copies of `ClientConnectionConfiguration` & `LeaderElectionConfiguration` instead of shared types.

### Goals

* Find a home for the ComponentConfig API types, hosted as a staging repo in the "core" repo that is kubernetes/kubernetes
* Make ComponentConfig API types consumable from *projects outside of kube* and from different parts of kube itself
    * Resolve dependencies from the external ComponentConfig API types so that everything can depend on them
    * The only dependency of the ComponentConfig API types should be `k8s.io/apimachinery`
* Split internal types from versioned types
* Remove the monolithic `componentconfig/v1alpha1` API group
* Enable the staging bot so that a `[https://github.com/kubernetes/](https://github.com/kubernetes/){component}`
  (imported as `k8s.io/{component}`) repos are published regularly.
* The future API server componentconfig code should be compatible with the proposed structure

### Non-goals

* Graduate the API versions
    * For v1.12, we’re working incrementally and will keep the versions of the existing ComponentConfigs.
* Do major refactoring of the ComponentConfigs. This PR is about code moves, not about re-defining the structure. We will do the latter in follow-ups.
* Change the components to support reading a config file, do flag precedence correctly or add e2e testing
    * Further, the "load-versioned-config-from-flag" feature in this proposal *should not* be confused with the
      [Dynamic Kubelet Configuration](https://kubernetes.io/docs/tasks/administer-cluster/reconfigure-kubelet/) feature. There is nothing in this proposal advocating for that
      a component should support a similar feature. This is all about the making the one-off “read bytes from a source and unmarshal into internal config state” possible for
      both the internal components and external consumers of these APIs
    * This is work to be done after this proposal is implemented (for every component but the kubelet which has this implemented already), and might or might not require
      further, more component-specific proposals/KEPs
* Create a net-new ComponentConfiguration struct for the API server
* Publish the internal types to the new `k8s.io/{component}` repo
* Support ComponentConfiguration for the cloud-controller-manager, as it’s a stop-gap for the cloud providers to move out of tree. This effort is in progress.
    * When the currently in-tree cloud providers have move out of tree, e.g. to `k8s.io/cloud-provider-gcp`, they should create their own external and internal types and
      make the command support loading configuration files.
    * The new repo can reuse the generic types from the to-be-created, `k8s.io/controller-manager` repo eventually.
    * Meanwhile, the cloud-controller-manager will reference the parts it needs from the main repo, and live privately in `cmd/cloud-controller-manager`
* Add further defaulting for the external types. The defaulting strategy decision is post-poned, outside of this KEP.

### Related proposals and further references

* [Original Google Docs version of this KEP](https://docs.google.com/document/d/1-u2y03ufX7FzBDWv9dVI_HyiIQZL8iz3o3vabeOpP5Y/edit)
* [Kubernetes Component Configuration](https://docs.google.com/document/d/1arP4T9Qkp2SovlJZ_y790sBeiWXDO6SG10pZ_UUU-Lc/edit) by [@mikedanese](https://github.com/mikedanese)
* [Versioned Component Configuration Files](https://docs.google.com/document/d/1FdaEJUEh091qf5B98HM6_8MS764iXrxxigNIdwHYW9c/edit#) by [@mtaufen](https://github.com/mtaufen)
* [Creating a ComponentConfig struct for the API server](https://docs.google.com/document/d/1fcStTcdS2Foo6dVdI787Dilr0snNbqzXcsYdF74JIrg/edit) by [@luxas](https://github.com/luxas) & [@sttts](https://github.com/sttts)
* Related tracking issues in kubernetes/kubernetes:
    * [Move `KubeControllerManagerConfiguration` to `pkg/controller/apis/`](https://github.com/kubernetes/kubernetes/issues/57618)
    * [kube-proxy config should move out of ComponentConfig apigroup](https://github.com/kubernetes/kubernetes/issues/53577)
    * [Move ClientConnectionConfiguration struct to its own api group](https://github.com/kubernetes/kubernetes/issues/54318)
    * [Move `CloudControllerManagerConfiguration` to `cmd/cloud-controller-manager/apis`](https://github.com/kubernetes/kubernetes/issues/65458)

## Proposal

* for component in [kubelet kubeproxy kubecontrollermanager kubeapiserver kubescheduler]
    * API group name: `{component}.config.k8s.io`
    * Kind name: `{Component}Configuration`
    * Code location:
        * External types: `k8s.io/{component}/config/{version}/types.go`
            * Like `k8s.io/api`
        * Internal types: `k8s.io/kubernetes/pkg/{component}/apis/config`
            * Alternatives, if applicable
                * `k8s.io/{component}/pkg/apis/config` (preferred)
                * `k8s.io/kubernetes/cmd/{component}/app/apis/config`
            * If dependencies allow it, we can move them to `k8s.io/{component}/pkg/apis/config/types.go`. Not having the external types there is intentional because the `pkg/` package tree is considered as "on-your-own-risks / no code compatibility guarantees", while `config/` is considered as a code API.
        * Internal scheme package: `k8s.io/kubernetes/pkg/{component}/apis/config/scheme/scheme.go`
            * The scheme package should expose `Scheme *runtime.Scheme`, `Codecs *serializer.CodecFactory`, and `AddToScheme(*runtime.Scheme)`, and have an `init()` method that runs `AddToScheme(Scheme)`
    * For the move to a staging repo to be possible, the external API package must not depend on the core repo.
        * Hence, all non-staging repo dependencies need to be removed/resolved before the package move.
    * Conversions from the external type to the internal type will be kept in `{internal_api_path}/{external_version}`, like for `k8s.io/api`
        * Defaulting code will be kept in this package, next to the conversions, if it already exists. We keep the decision about the destiny of defaulting open in this KEP.
* Create a "shared types"-package with structs generic to all or many componentconfig API groups, in the `k8s.io/apimachinery`, `k8s.io/apiserver` and `k8s.io/controller-manager` repos, depending on the struct.
    * Location: `k8s.io/{apimachinery,apiserver,controller-manager}/pkg/apis/config/{,v1alpha1}`
    * These aren’t "real" API groups, but have both internal and external versions
    * Conversions and internal types are published to the staging repo.
    * Defaults are moved into the leaf groups using them as much as possible.
* Remove the monolithic `componentconfig/v1alpha1` API group
* Enable the staging bot to create the Github repos
* We do not add further defaulting, but merely keep the minimum defaulting in-place to keep the current behaviour, but remove those defaulting funcs which are not used today.

### Migration strategy per component or k8s.io repo

#### k8s.io/apimachinery changes

* **Not a "real" API group, instead shared packages only with both external and internal types.**
* **External Package with defaulting (where absolutely necessary) & conversions**: `k8s.io/apimachinery/pkg/apis/config/v1alpha1/types.go`
* **Internal Package**: `k8s.io/apimachinery/pkg/apis/config/types.go`
* Structs to be hosted:
    * ClientConnectionConfiguration
* Assignee: @hanxiaoshuai

#### k8s.io/apiserver changes

* **Not a "real" API group, instead shared packages only with both external and internal types.**
* **External Package with defaulting (where absolutely necessary) & conversions**: `k8s.io/apiserver/pkg/apis/config/v1alpha1/types.go`
* **Internal Package**: `k8s.io/apiserver/pkg/apis/config/types.go`
* Structs to be hosted:
    * LeaderElectionConfiguration
    * DebuggingConfiguration
    * later to be created: SecureServingConfiguration, AuthenticationConfiguration, AuthorizationConfiguration, etc.
* Assignee: @hanxiaoshuai

#### kubelet changes

* **GroupVersionKind:** `kubelet.config.k8s.io/v1beta.KubeletConfiguration`
* **External Package:** `k8s.io/kubelet/config/v1beta1/types.go`
* **Internal Package:** `k8s.io/kubernetes/pkg/kubelet/apis/config/types.go`
* **Internal Scheme:** `k8s.io/kubernetes/pkg/kubelet/apis/config/scheme/scheme.go`
* **Conversions & defaulting (where absolutely necessary) Package:** `k8s.io/kubernetes/pkg/kubelet/apis/config/v1beta1`
* **Future Internal Package:** `k8s.io/kubelet/pkg/apis/config/types.go`
* Assignee: @mtaufen

#### kube-proxy changes

* **GroupVersionKind**: `kubeproxy.config.k8s.io/v1alpha1.KubeProxyConfiguration`
* **External Package**: `k8s.io/kube-proxy/config/v1alpha1/types.go`
* **Internal Package**: `k8s.io/kubernetes/pkg/proxy/apis/config/types.go`
* **Internal Scheme**: `k8s.io/kubernetes/pkg/proxy/apis/config/scheme/scheme.go`
* **Conversions & defaulting (where absolutely necessary) Package:** `k8s.io/kubernetes/pkg/proxy/apis/config/v1alpha1`
* **Future Internal Package:** `k8s.io/kube-proxy/pkg/apis/config/types.go`
* Start referencing `ClientConnectionConfiguration` from the generic ComponentConfig packages
* Assignee: @m1093782566

#### kube-scheduler changes

* **GroupVersionKind**: `kubescheduler.config.k8s.io/v1alpha1.KubeSchedulerConfiguration`
* **External Package**: `k8s.io/kube-scheduler/config/v1alpha1/types.go`
* **Internal Package**: `k8s.io/kubernetes/pkg/scheduler/apis/config/types.go`
* **Internal Scheme**: `k8s.io/kubernetes/pkg/scheduler/apis/config/scheme/scheme.go`
* **Conversions & defaulting (where absolutely necessary) Package:** `k8s.io/kubernetes/pkg/scheduler/apis/config/v1alpha1`
* **Future Internal Package:** `k8s.io/kube-scheduler/pkg/apis/config/types.go`
* Start referencing `ClientConnectionConfiguration` & `LeaderElectionConfiguration` from the generic ComponentConfig packages
* Assignee: @dixudx

#### k8s.io/controller-manager changes

* **Not a "real" API group, instead shared packages only with both external and internal types.**
* **External Package with defaulting (where absolutely necessary) & conversions**: `k8s.io/controller-manager/pkg/apis/config/v1alpha1/types.go`
* **Internal Package**: `k8s.io/controller-manager/pkg/apis/config/types.go`
* Will host structs:
    * GenericComponentConfiguration (should be renamed to GenericControllerManagerConfiguration at the same time)
* Assignee: @stewart-yu

#### kube-controller-manager changes

* **GroupVersionKind**: `kubecontrollermanager.config.k8s.io/v1alpha1.KubeControllerManagerConfiguration`
* **External Package**: `k8s.io/kube-controller-manager/config/v1alpha1/types.go`
* **Internal Package**: `k8s.io/kubernetes/pkg/controller/apis/config/types.go`
* **Internal Scheme**: `k8s.io/kubernetes/pkg/controller/apis/config/scheme/scheme.go`
* **Conversions & defaulting (where absolutely necessary) Package:** `k8s.io/kubernetes/pkg/controller/apis/config/v1alpha1`
* **Future Internal Package:** `k8s.io/kube-controller-manager/pkg/apis/config/types.go`
* Start referencing `ClientConnectionConfiguration` & `LeaderElectionConfiguration` from the generic ComponentConfig packages
* Assignee: @stewart-yu

#### cloud-controller-manager changes

* **Not a "real" API group, instead only internal types in `cmd/`.**
* **Internal Package:** `cmd/cloud-controller-manager/app/apis/config/types.go`
* We do not plan to publish any external types for this in a staging repo.
* The internal cloud-controller-manager ComponentConfiguration types will reference both `k8s.io/controller-manager/pkg/apis/config`
  and `k8s.io/kubernetes/pkg/controller/apis/config/`
* Assignee: @stewart-yu

#### kube-apiserver changes

* **Doesn’t have a ComponentConfig struct at the moment, so there is nothing to move around.**
* Eventually, we want to create this ComponentConfig struct, but exactly how to do that is out of scope for this specific proposal.
* See [Creating a ComponentConfig struct for the API server](https://docs.google.com/document/d/1fcStTcdS2Foo6dVdI787Dilr0snNbqzXcsYdF74JIrg/edit) for a proposal on how to refactor the API server code to be able to expose the final ComponentConfig structure.

### Vendoring

#### Vgo

Vgo – as the future standard vendoring mechanism in Golang – supports [vgo modules](https://research.swtch.com/vgo-module) using a `k8s.io/{component}/config/go.mod` file. Tags of the shape `config/vX.Y` on `k8s.io/{component}` will define a version of the component config of that component. Such a tagged module can be imported into a 3rd-party program without inheriting dependencies outside of the `k8s.io/{component}/config` package.

The `k8s.io/{component}/config/go.mod` file will look like this:

```
module "k8s.io/{component}/config"
require (
	"k8s.io/apimachinery" v1.12.0
)
```

The exact vgo semver versioning scheme we will use is out of scope of this document. We will be able to version the config package independently from the main package `k8s.io/{component}` if we want to, e.g. to implement correct semver semantics.

Other 3rd-party code can import the config module as usual. Vgo does not add the dependencies from code outside of `k8s.io/{component}/config` (actually, vgo creates a separate `vgo.sum` for the config package with the transitive dependencies).

Compare http://github.com/sttts/kubeadm for a test project using latest vgo.

#### Dep

Dep supports the import of sub-packages without inheriting dependencies from outside of the sub-package.

### Timeframe and Implementation Order

Objective: Done for v1.12

Implementation order:
* Start with the apiserver and apimachinery shared config repos, copy over the necessary structs as needed and create the external and internal API packages.
* Start with the scheduler as the first component.
* One PR for moving `KubeSchedulerConfiguration` to `staging/src/k8s.io/kube-scheduler/config/v1alpha1/types.go`, and the internal type to `pkg/scheduler/apis/config/types.go`.
    * Set up the conversion for the external type by creating the package `pkg/scheduler/apis/config/v1alpha1`, without `types.go`, like how `k8s.io/api` is set up.
    * This should be a pure code move.
* Set up staging publishing bot (async, non-critical)
* The kubelet, kube-proxy and controller-manager follow, each one independently.

### OWNERS files for new packages and repos

* Approvers:
    * @kubernetes/api-approvers
    * @sttts
    * @luxas
    * @mtaufen
* Reviewers:
    * @kubernetes/api-reviewers
    * @sttts
    * @luxas
    * @mtaufen
    * @dixudx
    * @stewart-yu
