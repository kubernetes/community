---
kep-number: 13
title: Removing In-Tree Cloud Providers
authors:
  - "@andrewsykim"
  - "@cheftako"
owning-sig: sig-cloudprovider
participating-sigs:
  - sig-apps
  - sig-aws
  - sig-azure
  - sig-apimachinery
  - sig-gcp
  - sig-network
  - sig-openstack
  - sig-storage
reviewers:
  - "@andrewsykim"
  - "@cheftako"
  - "@d-nishi"
  - "@dims"
  - "@hogepodge"
  - "@mcrute"
  - "@steward-yu"
approvers:
  - "@thockin"
editor: TBD
status: implementable

---

# Removing In-Tree Cloud Provider Code

## Table of Contents

   * [Removing In-Tree Cloud Provider Code](#removing-in-tree-cloud-provider-code)
      * [Table of Contents](#table-of-contents)
      * [Terms](#terms)
      * [Summary](#summary)
      * [Motivation](#motivation)
         * [Goals](#goals)
         * [Non-Goals](#non-goals)
      * [Proposal](#proposal)
         * [Approach](#approach)
            * [Phase 1 - Moving Cloud Provider Code to Staging](#phase-1---moving-cloud-provider-code-to-staging)
            * [Phase 2 - Building CCM from Provider Repos](#phase-2---building-ccm-from-provider-repos)
            * [Phase 3 - Migrating Provider Code to Provider Repos](#phase-3---migrating-provider-code-to-provider-repos)
         * [Staging Directory](#staging-directory)
            * [Cloud Provider Instances](#cloud-provider-instances)
      * [Alternatives](#alternatives)
         * [Staging Alternatives](#staging-alternatives)
            * [Git Filter-Branch](#git-filter-branch)
         * [Build Location Alternatives](#build-location-alternatives)
            * [Build K8s/K8s from within K8s/Cloud-provider](#build-k8sk8s-from-within-k8scloud-provider)
            * [Build K8s/Cloud-provider within K8s/K8s](#build-k8scloud-provider-within-k8sk8s)
         * [Config Alternatives](#config-alternatives)
            * [Use component config to determine where controllers run](#use-component-config-to-determine-where-controllers-run)

## Terms

- **CCM**: Cloud Controller Manager - The controller manager responsible for running cloud provider dependent logic,
such as the service and route controllers.
- **KCM**: Kubernetes Controller Manager - The controller manager responsible for running generic Kubernetes logic,
such as job and node_lifecycle controllers.
- **KAS**: Kubernetes API Server - The core api server responsible for handling all API requests for the Kubernetes
control plane. This includes things like namespace, node, pod and job resources.
- **K8s/K8s**: The core kubernetes github repository.
- **K8s/cloud-provider**: Any or all of the repos for each cloud provider. Examples include [cloud-provider-gcp](https://github.com/kubernetes/cloud-provider-gcp),
[cloud-provider-aws](https://github.com/kubernetes/cloud-provider-aws) and [cloud-provider-azure](https://github.com/kubernetes/cloud-provider-azure).
We have created these repos for each of the in-tree cloud providers. This document assumes in various places that the
cloud providers will place the relevant code in these repos. Whether this is a long-term solution to which additional
cloud providers will be added, or an incremental step toward moving out of the Kubernetes org is out of scope of this
document, and merits discussion in a broader forum and input from SIG-Architecture and Steering Committee.
- **K8s SIGs/library**: Any SIG owned repository.
- **Staging**: Staging: Separate repositories which are currently visible under the K8s/K8s repo, which contain code
considered to be safe to be vendored outside of the K8s/K8s repo and which should eventually be fully separated from
the K8s/K8s repo. Contents of Staging are prevented from depending on code in K8s/K8s which are not in Staging.
Controlled by [publishing kubernetes-rules-configmap](https://github.com/kubernetes/publishing-bot/blob/master/configs/kubernetes-rules-configmap.yaml)


## Summary

This is a proposal outlining steps to remove "in-tree" cloud provider code from the kubernetes/kubernetes repo.

## Motivation

Motiviation behind this effort is to allow cloud providers to develop and make releases independent from the core
Kubernetes release cycle. The de-coupling of provider code allows for separation of concern between the core Kubernetes
and the initial cloud providers that are currently built into the project. Additionally, one of the goals of SIG Cloud
Provider is to move all providers out-of-tree to promote a vendor neutral ecosystem.

### Goals

- Remove all cloud provider specific code from kubernetes/kubernetes.
- Move the cloud provider specific controller manager logic into repos appropriate for those cloud providers.

### Non-Goals

- Testing/validating out-of-tree provider code for all cloud providers, this should be done by providers

## Proposal

### Approach

In order to remove cloud provider code from kubernetes/kubernetes. A 3 phase approach will be taken.

1. Move all code in `kubernetes/kubernetes/pkg/cloudprovider/providers` to `kubernetes/kubernetes/staging/cloud-provider-<provider>/`. This requires removing all dependencies in each cloud provider to `kubernetes/kubernetes`.
2. Begin to build CCM from external repos (`kubernetes/cloud-provider-<provider>`) via code synced from `kubernetes/kubernetes/staging/cloud-provider-<provider>`. This allows us to develop providers "out-of-tree" while being able to build both in-tree components like kube-controller-manager and out-of-tree components like cloud-controller-manager. Development is still done in `kubernetes/kubernetes/staging/cloud-provider-<provider>`.
3. Delete all code in `kubernetes/kubernetes/staging/cloud-provider-<provider>` and shift main development to `kubernetes/cloud-provider-<provider>`.

#### Phase 1 - Moving Cloud Provider Code to Staging

In Phase 1, all cloud provider code in `k8s.io/kubernetes/pkg/cloudprovider/providers` will be moved to `k8s.io/kubernetes/staging/k8s.io/cloud-provider/providers`. Moving code to a staging directory still allows us to build core Kubernetes components (`kube-controller-manager`, `kubelet`, etc) without interruption while also signalling to the community that in-tree provider code is slated for removal.

The biggest challenge of this phase is to remove dependences to k8s.io/kubernetes in all the providers. This is required to avoid circular dependencies within the various Kubernetes repos. All other repos "staged" (`client-go`, `apimachinery`, `api`, etc) in Kubernetes follow the same pattern. Below are the current list of packges in `k8s.io/kubernetes` that we need to removed. Note that _some_ dependencies may need to be duplicated for the duration of this migration as other critical components of Kubernetes may share the same code but removing the dependency may not be trivial. In general, we will avoid duplicating code and only use it as a last resort to resolve circular dependencies to `kubernetes/kubernetes`.

| Dependency                        | Proposed Location            | Is duplicated? | Partial or Full? |
|-----------------------------------|------------------------------|----------------|------------------|
| pkg/api/service                   |                              |                |                  |
| pkg/api/v1/service                |                              |                |                  |
| pkg/apis/core/v1/helper           |                              |                |                  |
| pkg/controller                    |                              |                |                  |
| pkg/credentialprovider/aws        |                              |                |                  |
| pkg/master/ports                  |                              |                |                  |
| pkg/cloudprovider                 | k8s.io/cloud-provider        | No             | Full             |
| pkg/features                      |                              |                |                  |
| pkg/kubelet/apis                  |                              |                |                  |
| pkg/util/file                     | k8s.io/utils/file            | No             | Full             |
| pkg/util/keymutex                 | k8s.io/utils/keymutex        | No             | Full             |
| pkg/util/io                       | k8s.io/utils/io              | No             | Full             |
| pkg/util/mount                    |                              |                |                  |
| pkg/util/net/sets                 | k8s.io/utils/netsets         | No             | Full             |
| pkg/util/node                     | k8s.io/cloud-provider/node   | No             | Partial          |
| pkg/util/nsenter                  | k8s.io/utils/nsenter         |                |                  |
| pkg/util/parsers                  |                              |                |                  |
| pkg/util/strings                  | k8s.io/utils/strings         | No             | Full             |
| pkg/version                       |                              |                |                  |
| pkg/volume                        | k8s.io/cloud-provider/volume | Yes            | Partial          |
| pkg/volume/util                   | k8s.io/cloud-provider/volume | Yes            | Partial          |
| pkg/volume/util/volumepathhandler | k8s.io/cloud-provider/volume | Yes            | Partial          |

#### Phase 2 - Building CCM from Provider Repos

In Phase 2, providers will be expected to build cloud controller manager from their respective provider repos (`kubernetes/cloud-provider-<provider>`). Provider repos will be frequently synced from their respective staging repos by the kubernetes publishing bot. Development of providers are still done in `kubernetes/kubernetes/staging/cloudprovider/providers`.

#### Phase 3 - Migrating Provider Code to Provider Repos

In Phase 3, all code in `kubernetes/kubernetes/staging/cloud-provider-<provider>` will be removed and development of Kubernetes cloud providers should be done in provider specific repos. It's important that by this phase, both in-tree and out-of-tree cloud providers are tested and production ready. Ideally most Kubernetes clusters in production should be using the out-of-tree provider before in-tree support is removed.

### Staging Directory

There are several sections of code which need to be shared between the K8s/K8s repo and the K8s/Cloud-provider repos.
The plan for doing that sharing is to move the relevant code into the Staging directory as that is where we share code
today. The current Staging repo has the following packages in it.
- Api
- Apiextensions-apiserver
- Apimachinery
- Apiserver
- Client-go
- Code-generator
- Kube-aggregator
- Metrics
- Sample-apiserver
- Sample-Controller

With the additions needed in the short term to make this work; the Staging area would now need to look as follows.
- Api
- Apiextensions-apiserver
- Apimachinery
- Apiserver
- Client-go
- **cloud-provider-aws**
- **cloud-provider-azure**
- **cloud-provider-cloudstack**
- **cloud-provider-gce**
- **cloud-provider-openstack**
- **cloud-provider-ovirt**
- **cloud-provider-photon**
- **cloud-provider-vsphere**
- Code-generator
- Kube-aggregator
- Metrics
- Sample-apiserver
- Sample-Controller

#### Cloud Provider Instances

Currently in K8s/K8s the cloud providers are actually included by including [providers.go](https://github.com/kubernetes/kubernetes/blob/master/pkg/cloudprovider/providers/providers.go)
file which then includes each of the in-tree cloud providers. In the short term we would leave that file where it is
and adjust it to point at the new homes under Staging. For the K8s/cloud-provider repo, would have the following CCM
wrapper file. (Essentially a modified copy of cmd/cloud-controller-manager/controller-manager.go) The wrapper for each
cloud provider would import just their vendored cloud-provider implementation rather than providers.go file.

k8s/k8s: pkg/cloudprovider/providers/providers.go
```package cloudprovider

import (
  // Prior to cloud providers having been moved to Staging
  _ "k8s.io/cloudprovider-aws"
  _ "k8s.io/cloudprovider-azure"
  _ "k8s.io/cloudprovider-cloudstack"
  _ "k8s.io/cloudprovider-gce"
  _ "k8s.io/cloudprovider-openstack"
  _ "k8s.io/cloudprovider-ovirt"
  _ "k8s.io/cloudprovider-photon"
  _ "k8s.io/cloudprovider-vsphere"
)
```

k8s/cloud-provider-gcp: pkg/cloudprovider/providers/providers.go
```package cloudprovider

import (
  // Cloud providers
  _ "k8s.io/cloud-provider-aws"
  _ "k8s.io/cloud-provider-azure"
  _ "k8s.io/cloud-provider-gce"
  _ "k8s.io/cloud-provider-openstack"
  _ "k8s.io/cloud-provider-vsphere"
)
```

## Alternatives

### Staging Alternatives

#### Git Filter-Branch

One possible alternative is to make use of a Git Filter Branch to extract a sub-directory into a virtual repo. The repo
needs to be sync'd in an ongoing basis with K8s/K8s as we want one source of truth until K8s/K8s does not pull in the
code. This has issues such as not giving K8s/K8s developers any indications of what the dependencies various
K8s/Cloud-providers have. Without that information it becomes very easy to accidentally break various cloud providers
and time you change dependencies in the K8s/K8s repo. With staging the dependency line is simple and [automatically
enforced](https://github.com/kubernetes/kubernetes/blob/master/hack/verify-no-vendor-cycles.sh). Things in Staging are
not allowed to depend on things outside of Staging. If you want to add such a dependency you need to add the dependent
code to Staging. The act of doing this means that code should get synced and solve the problem. In addition the usage
of a second different library and repo movement mechanism will make things more difficult for everyone.

“Trying to share code through the git filter will not provide this protection. In addition it means that we now have
two sharing code mechanisms which increases complexity on the community and build tooling. As such I think it is better
to continue to use the Staging mechanisms. ”

### Build Location Alternatives

#### Build K8s/K8s from within K8s/Cloud-provider

The idea here is to avoid having to add a new build target to K8s/K8s. The K8s/Cloud-provider could have their own
custom targets for building things like KAS without other cloud-providers implementations linked in. It would also
allow other customizations of the standard binaries to be created. While a powerful tool, this mechanism seems to
encourage customization of these core binaries and as such to be discouraged. Providing the appropriate generic
binaries cuts down on the need to duplicate build logic for these core components and allow each optimization of build.
Download prebuilt images at a version and then just build the appropriate addons.

#### Build K8s/Cloud-provider within K8s/K8s

The idea here would be to treat the various K8s/Cloud-provider repos as libraries. You would specify a build flavor and
we would pull in the relevant code based on what you specified. This would put tight restrictions on how the
K8s/Cloud-provider repos would work as they would need to be consumed by the K8s/K8s build system. This seems less
extensible and removes the nice loose coupling which the other systems have. It also makes it difficult for the cloud
providers to control their release cadence.

### Config Alternatives

#### Use component config to determine where controllers run

Currently KCM and CCM have their configuration passed in as command line flags. If their configuration were obtained
from a configuration server (component config) then we could have a single source of truth about where each controller
should be run. This both solves the HA migration issue and other concerns about making sure that a controller only runs
in 1 controller manager. Rather than having the controllers as on or off, controllers would now be configured to state
where they should run, KCM, CCM, Nowhere, … If the KCM could handle this as a run-time change nothing would need to
change. Otherwise it becomes a slight variant of the proposed solution. This is probably the correct long term
solution. However for the timeline we are currently working with we should use the proposed solution.
