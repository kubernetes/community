---
kep-number: 13
title: Switching To Cloud Provider Repo And Builds
authors:
  - "@cheftako"
  - "@calebamiles"
  - "@nckturner"
owning-sig: sig-apimachinery
participating-sigs:
  - sig-apps
  - sig-aws
  - sig-azure
  - sig-cloud-provider
  - sig-gcp
  - sig-network
  - sig-openstack
  - sig-storage
reviewers:
  - "@andrewsykim"
  - "@calebamiles"
  - "@nckturner"
approvers:
  - "@thockin"
editor: TBD
status: provisional
---

# Switching To Cloud Provider Repo And Builds 

## How To Remove Cloud Provider Code From Kubernetes Core

## Table of Contents

- [Switching To Cloud Provider Repo And Builds](#switching-to-cloud-provider-repo-and-builds)
   - [Table of Contents](#table-of-contents)
   - [Terms](#terms)
   - [Summary](#summary)
   - [Motivation](#motivation)
      - [Goals](#goals)
      - [Intermediary Goals](#intermediary-goals)
      - [Non-Goals](#non-goals)
   - [Proposal](#proposal)
   - [Building CCM For Cloud Providers](#building-ccm-for-cloud-providers)
      - [Background](#background)
      - [Design Options](#design-options)
          - [Staging](#staging)
          - [Cloud Provider Instances](#cloud-provider-instances)
          - [Build Targets](#build-targets)
          - [K8s/K8s Releases and K8s/CP Releases](#k8s/k8s-releases-and-k8s/cp-releases)
          - [Migrating to a CCM Build](#migrating-to-a-ccm-build)
          - [Flags, Service Accounts, etc](#flags,-service-accounts,-etc)
          - [Deployment Scripts](#deployment-scripts)
          - [kubeadm](#kubeadm)
          - [CI, TestGrid and Other Testing Issues](#ci,-testgrid-and-other-testing-issues)
   - [Alternatives](#alternatives)
      - [Staging Alternatives](#staging-alternatives)
          - [Git Filter-Branch](#Git Filter-Branch)
      - [Build Location Alternatives](#Build Location Alternatives)
          - [Build K8s/K8s from within K8s/Cloud-provider](#build-k8s/k8s-from-within-k8s/cloud-provider)
          - [Build K8s/Cloud-provider within K8s/K8s](#build-k8s/cloud-provider-within-k8s/k8s)
      - [Config Alternatives](#config-alternatives)
          - [Use component config to determine where controllers run](#use-component-config-to-determine-where-controllers-run)

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

We want to remove any cloud provider specific logic from the kubernetes/kubernetes repo. We want to restructure the code
to make it easy for any cloud provider to extend the kubernetes core in a consistent manner for their cloud. New cloud
providers should look at the [Creating a Custom Cluster from Scratch](https://kubernetes.io/docs/getting-started-guides/scratch/#cloud-provider)
and the [cloud provider interface](https://github.com/kubernetes/kubernetes/blob/master/pkg/cloudprovider/cloud.go#L31)
which will need to be implemented.

## Motivation

We are trying to remove any dependencies from Kubernetes Core to any specific cloud provider. Currently we have seven
such dependencies. To prevent this number from growing we have locked Kubernetes Core to the addition of any new
dependencies. This means all new cloud providers have to implement all their pieces outside of the Core.
However everyone still ends up consuming the current set of seven in repo dependencies. For the seven in repo cloud
providers any changes to their specific cloud provider code requires OSS PR approvals and a deployment to get those
changes in to an official build. The relevant dependencies require changes in the following areas.

- [Kube Controller Manager](https://kubernetes.io/docs/reference/generated/kube-controller-manager/) - Track usages of [CMServer.CloudProvider](https://github.com/kubernetes/kubernetes/blob/master/cmd/kube-controller-manager/app/options/options.go)
- [API Server](https://kubernetes.io/docs/reference/generated/kube-apiserver/) - Track usages of [ServerRunOptions.CloudProvider](https://github.com/kubernetes/kubernetes/blob/master/cmd/kube-apiserver/app/options/options.go)
- [kubelet](https://kubernetes.io/docs/reference/generated/kubelet/) - Track usages of [KubeletFlags.CloudProvider](https://github.com/kubernetes/kubernetes/blob/master/cmd/kubelet/app/options/options.go)
- [How Cloud Provider Functionality is deployed to and enabled in the cluster](https://kubernetes.io/docs/setup/pick-right-solution/#hosted-solutions) - Track usage from [PROVIDER_UTILS](https://github.com/kubernetes/kubernetes/blob/master/cluster/kube-util.sh)

For the cloud providers who are in repo, moving out would allow them to more quickly iterate on their solution and
decouple cloud provider fixes from open source releases. Moving the cloud provider code out of the open source
processes means that these processes do not need to load/run unnecessary code for the environment they are in.
We would like to abstract a core controller manager library so help standardize the behavior of the cloud
controller managers produced by each cloud provider. We would like to minimize the number and scope of controllers
running in the cloud controller manager so as to minimize the surface area for per cloud provider deviation.

### Goals

- Get to a point where we do not load the cloud interface for any of kubernetes core processes.
- Remove all cloud provider specific code from kubernetes/kubernetes.
- Have a generic controller manager library available for use by the per cloud provider controller managers.
- Move the cloud provider specific controller manager logic into repos appropriate for those cloud providers.

### Intermediary Goals

Have a cloud controller manager in the kubernetes main repo which hosts all of
the controller loops for the in repo cloud providers.
Do not run any cloud provider logic in the kube controller manager, the kube apiserver or the kubelet.
At intermediary points we may just move some of the cloud specific controllers out. (Eg. volumes may be later than the rest)

### Non-Goals

Forcing cloud providers to use the generic cloud manager.

## Proposal

## Building CCM For Cloud Providers

### Background

The CCM in K8s/K8s links in all eight of the in-tree providers (aws, azure, cloudstack, gce, openstack, ovirt, photon 
and vsphere). Each of these providers has an implementation of the Cloud Provider Interface in the K8s/K8s repo. CNCF 
has created a repo for each of the cloud providers to extract their cloud specific code into. The assumption is that 
this will include the CCM executable, their Cloud Provider Implementation and various build and deploy scripts. Until 
we have extracted every in-tree provider and removed cloud provider dependencies from other binaries, we need to 
maintain the Cloud Provider Implementation in the K8s/K8s repo. After the Cloud Provider specific code has been 
extracted Golang imports for things like the service controller, prometheus code, utilities etc will still require each 
cloud provider specific repository to vendor in a significant portion of the code in k8s/k8s. 

We need a solution which meets the objective in both the short and long term. In the short term we cannot delete CCM or 
cloud provider code from the K8s/K8s repo. We need to keep this code in K8s/K8s while we still support cloud provider 
deployments from K8s/K8s. In the long term this code should be part of each cloud providers repo and that code should 
be removed from K8s/K8s. This suggests that in the short term that code should have one source of truth. However it 
should probably not end up in the vendor directory as that is not its intended final home. Other code such as the 
controllers should end up in the vendor directory. Additionally each provider will need their own copy of 
pkg/cloudprovider/providers/providers.go and related build file to properly control which Cloud Provider Implementation 
get linked in.

We also need to be able to package a combination of binaries from K8s/K8s and K8s/cloud-provider-<cp> into a deployable 
package. The code for this will need to accommodate things like differing side cars for each cloud provider’s CSI 
implementation and possible desire to run additional controller managers or extension api servers. As such it seems 
better to have this code live in the cloud provider specific repo. This also allows this code to be simpler as it does 
not have to attempt to support all the different cloud provider configurations. This separate from things like the 
local deployment option, which K8s/K8s should continue to support.

Lastly there are specific flags which need to be set on various binaries for this to work. Kubernetes API-Server, 
Kubernetes Controller-Manager and kubelet should all have the --cloud-provider flag set to external. For the Cloud 
Controller-Manager the --cloud-provider flag should be set appropriately for that cloud provider. In addition we need 
to set the set of controllers running in the Kubernetes Controller-Manager. More on that later.

For further background material please look at [running cloud controller](https://kubernetes.io/docs/tasks/administer-cluster/running-cloud-controller/).

### Design Options

For the short term problem of sharing code between K8s/K8s repo, the K8s/cloud-provider repos and K8s SIGs/library repos; 
there are 2 fundamental solutions. We can push more code into Staging to make the code available to the 
K8s/cloud-provider repos. Currently code which needs to be shared between K8s/K8s itself and other projects is put 
the code in Staging. This allows the build machinery to detect things like attempts to make the shared code depend on 
code which is not shared (i.e. disallowing code in Staging from depending Iton code in K8s/K8s but not in Staging) 
In addition using Staging means that we can benefit from work to properly break out the libraries/code which exists in 
Staging. Most of the repositories we are adding to staging should end up as K8s SIGs/library repos. The K8s/cloud-provider 
repos should be the top of any build dependency trees. Code which needs to be consumed by other repos (Eg CSI plugins, 
shared controllers, ...) should not be in the K8s/cloud-provider repo but in an appropriate K8s SIGs/library repo.
 
The code which needs to be shared can be broken into several types. 

There is code which properly belongs in the various cloud-provider repos. The classic example of this would be the 
implementations of the cloud provider interface. (Eg. [gce](https://github.com/kubernetes/kubernetes/tree/master/pkg/cloudprovider/providers/gce) 
or [aws](https://github.com/kubernetes/kubernetes/tree/master/pkg/cloudprovider/providers/aws)) These sections of code 
need to be shared until we can remove all cloud provider dependencies from K8s/K8s. (i.e. When KAS, KCM and kubelet no 
longer contain a cloud-provider flag and no longer depend on either the cloud provider interface or any cloud provider 
implementations) At that point they should be permanently moved to the individual provider repos. I would suggest that 
as long as the code is shared it be in vendor for K8s/cloud-provider. We would want to create a separate Staging repo 
in K8s/K8s for each cloud-provider. 

The majority of the controller-manager framework and its dependencies also needs to be shared. This is code that would 
be shared even after cloud providers are removed from K8s/K8s. As such it probably makes sense to make the controller 
manager framework its own K8s/K8s Staging repo. 

It should be generally possible for cloud providers to determine where a controller runs and even over-ride specific 
controller functionality. Please note that if a cloud provider exercises this possibility it is up to that cloud provider 
to keep their custom controller conformant to the K8s/K8s standard. This means any controllers may be run in either KCM 
or CCM. As an example the NodeIpamController, will be shared across K8s/K8s and K8s/cloud-provider-gce, both in the 
short and long term. Currently it needs to take a cloud provider to allow it to do GCE CIDR management. We could handle 
this by leaving the cloud provider interface with the controller manager framework code. The GCE controller manager could 
then inject the cloud provider for that controller. For everyone else (especially the KCM) NodeIpamController is 
interesting because currently everyone needs its generic behavior for things like ranges. However Google then wires in 
the cloud provider to provide custom functionality on things like CIDRs. The thought then is that in the short term we 
allow it to be run in either KCM or CCM. Most cloud providers would run it in the KCM, Google would run it in the CCM. 
When we are ready to move cloud provider code out of K8s/K8s, we remove the cloud provider code from the version which 
is in K8s/K8s and continue to have flags to control if it runs. K8s/Cloud-Provider-Google could then have an enhanced 
version which is run in the CCM. Other controllers such as Route and Service needs to run in either the KCM or CCM. For 
things like K8s/K8s e2e tests we will always want these controllers in the K8s/K8s repo. Having it in the K8s/K8s repo 
is also useful for keeping the behavior of these sort of core systems consistent.

#### Staging

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
- **Controller**
  - **Cloud**
  - **Service**
  - **NodeIpam**
  - **Route**
  - **?Volume?**
- **Controller-manager**
- **Cloud-provider-aws**
- **Cloud-provider-azure**
- **Cloud-provider-cloudstack**
- **Cloud-provider-gce**
- **Cloud-provider-openstack**
- **Cloud-provider-ovirt**
- **Cloud-provider-photon**
- **Cloud-provider-vsphere**
- Code-generator
- Kube-aggregator
- **Kube-utilities**
- Metrics
- Sample-apiserver
- Sample-Controller

When we complete the cloud provider work, several of the new modules in staging should be moving to their permanent new 
home in the appropriate K8s/Cloud-provider repos they will no longer be needed in the K8s/K8s repo. There are however 
other new modules we will add which continue to be needed by both K8s/K8s and K8s/Cloud-provider. Those modules will 
remain in Staging until the Staging initiative completes and they are moved into some other Kubernetes shared code repo.
- Api
- Apiextensions-apiserver
- Apimachinery
- Apiserver
- Client-go
- **Controller**
  - **Cloud**
  - **Service**
  - **NodeIpam**
  - **Route**
  - **?Volume?**
- **Controller-manager**
- ~~Cloud-provider-aws~~ 
- ~~Cloud-provider-azure~~
- ~~Cloud-provider-cloudstack~~
- ~~Cloud-provider-gce~~
- ~~Cloud-provider-openstack~~
- ~~Cloud-provider-ovirt~~
- ~~Cloud-provider-photon~~
- ~~Cloud-provider-vsphere~~
- Code-generator
- Kube-aggregator
- **Kube-utilities**
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
  _ "k8s.io/cloudprovider-gce"
)
```

#### Build Targets

We then get to the issue of creating a deployable artifact for each cloud provider. There are several artifacts beyond 
the CCM which are cloud provider specific. These artifacts include things like the deployment scripts themselves, the 
contents of the add-on manager and the sidecar needed to get CSI/cloud-specific persistent storage to work. Ideally 
these would then be packaged with a version of the Kubernetes core components (KAS, KCM, kubelet, …) which have not 
been statically linked against the cloud provider libraries. However in the short term the K8s/K8s deployable builds 
will still need to link these binaries against all of the in-tree plugins. We need a way for the K8s/cloud-provider 
repo to consume artifacts generated by the K8s/K8s repo. For official releases these artifacts should be published to a 
package repository. From there the cloud providers can pull the cloud agnostic Kubernetes artifact and decorate it 
appropriate to their cloud. The K8s/Cloud-Provider should be using some sort of manifest file to determine which 
official K8s/K8s artifact to pull. This allows for things like repeatability in builds, for hot fix builds and also a 
mechanism for rolling out changes which span K8s/K8s and K8s/Cloud-Provider. For local cloud releases we can use a 
build convention. We can expect the K8s/K8s and K8s/Cloud-Provider repos to be checked out under the same GOPATH. 
The K8s/Cloud-Provider can then have a local build target which looks for the K8s/K8s artifacts to be in the 
appropriate place under the GOPATH (or at a location pointed to be a K8s home environment variable) This allows for 
both official builds and for developers to easily work at running changes which span K8s/K8s and K8s/Cloud-Provider.

#### K8s/K8s Releases and K8s/CP Releases

One of the goals for this project is to free cloud providers to generate releases as they want. This implies that 
K8s/CP releases are largely decoupled from K8s/K8s releases. The cadence of K8s/K8s releases should not change. We are 
assuming that K8s/CP releases will be on a similar or faster cadence as needed. It is desirable for the community for 
all cloud providers to support recent releases. As such it would be good for them to minimize the lag between K8s/K8s 
releases and K8s/CP releases. A K8s/CP cannot however release a Kubernetes version prior to that version having been 
released by K8s/K8s. (So for example the cloud provider cannot make a release with a 1.13 Kubernetes core, prior to 
K8s/K8s having released 1.13) The ownership and responsibility of publishing K8s/K8s will not change with this project. 
The K8s/CP releases must necessarily move to being the responsibility of that cloud provider or a set of owners 
delegated by that cloud provider.

#### Migrating to a CCM Build

Migrating to a CCM build requires some thought. When dealing with a new cluster, things are fairly simple; we install 
everything at the same version and then the customer begins to customize. Migrating an existing cluster which is a 
generic K8s/K8s cluster with the cloud-provider flag set to a K8s/Cloud-Provider built is a bit trickier. The system 
needs to work during the migration where disparate pieces can be on different versions. While we specify that the exact 
upgrade steps are cloud provider specific, we do provide guidance that the control plane (master) should be upgraded 
first (master version >= kubelet version) and that the system should be able to handle up to a 2 revision difference 
between the control plane and the kubelets. In addition with disruptions budgets etc, there will not be a consistent 
version of the kubelets, until the upgrade completes. So we need to ensure that our cloud provider/CCM builds work with 
existing clusters. This means that we need to account for things like older kubelets having the cloud provider enabled 
and using it for things like direct volume mount/unmount, IP discovery, … We can even expect that scaling events such 
as increases in the size of a replica set may cause us to deploy old kubelet images which directly use the cloud 
provider implementation in clusters which are controlled by a CCM build. We need to make sure we test these sort of 
scenarios and ensure they work (get their IP, can mount cloud specific volume types, …) 

HA migrations presents some special issues. HA masters are composed of multiple master nodes in which components like 
the controller managers use leader election to determine which is the currently operating instance. Before the 
migration begins there is no guarantee which instances will be leading or that all the leaders will be on the same 
master instance. So we will begin by taking down 1 master instance and upgrading it. At this point there will be some 
well known conditions. The Kube-Controller-Master in the lead will be one of the old build instances. If the leader had 
been the instance running on that master it will lose its lease to one of the older instances. At the same time there 
will only be one Cloud-Controller-Manager which is the running on the master running the new code. This implies that we 
will have a few controllers running in both the new lead cloud-controller-manager and the old kube-controller-manager. 
I would suggest that as an initial part of a HA upgrade we disable these controllers in the kube-controller-managers. 
This can be accomplished by providing the controllers flag. If the only controller we wanted to disable were service 
and route then we would set the flag as follows
 ```
kube-controller-manager --controllers=\"*,-service,-route\"
``` 
This assumes that upgrading the first master instance can be accomplished inside of the SLO for these controllers being 
down.

#### Flags, Service Accounts, etc

The correct set of flags, service accounts etc, which will be needed for each cloud provider, is expected to be 
different and is at some level, left as an exercise for each cloud provider. That having been said, there are a few 
common guidelines which are worth mentioning. It is expected that all the core components (kube-apiserver, 
kube-controller-manager, kubelet) should have their --cloud-provider flag set to “external”. Cloud-providers, who have 
their own volume type (eg. gce-pd) but do not yet have the CSI plugin (& side car) enabled in their CCM build, will 
need to set the --external-cloud-volume-plugin to their cloud provider implementation key. (eg. gce) There are also a 
considerable number of roles and bindings which are needed to get the CCM working. For core components this is handled 
through a bootstrapping process inside of the kube-apiserver. However the CCM and cloud-provider pieces are not 
considered core components. The expectation then is that the appropriate objects will be created by deploying yaml 
files for them in the addons directory. The add on manager (or cloud provider equivalent system) will then cause the 
objects to be created as the system comes up. The set of objects which we know we will need include :-
- ServiceAccount
  - cloud-controller-manager
- User
  - system:cloud-controller-manager
- Role
  - system::leader-locking-cloud-controller-manager
- ClusterRole
  - system:controller:cloud-node-controller
  - system:cloud-controller-manager
  - system:controller:pvl-controller
- RoleBinding
  - cloud-controller-manager to system::leader-locking-cloud-controller-manager
- ClusterRoleBinding
  - cloud-node-controller to system:controller:cloud-node-controller

#### Deployment Scripts

Currently there is a lot of common code in the K8s/K8s cluster directory. Each of the cloud providers today build on 
top of that common code to do their deployment. The code running in K8s/Cloud-provider will necessarily be different. 
We have new files (addon) and executables (CCM and CSI) to be deployed. The existing executables need to be started 
with different flags (--cloud-provider=external). We also have additional executables which need to be started. This 
may then result in different resource requirements for the master and the kubelets. So it is clear that there will need 
to be at least some changes between the deployment scripts going from K8s/K8s to K8s/Cloud-provider. There is also 
likely to be some desire to stream-line and simplify these scripts in K8s/Cloud-provider. A lot of the generic handling 
and flexibility in K8s/K8s is not needed in K8s/Cloud-provider. It is also worth looking at 
[CCM Repo Requirements](#repository-requirements) for some suggestions on common K8s/Cloud-provider 
layout suggestions. These include an installer directory for custom installer code.

#### kubeadm [WIP]

kubeadm is a tool for creating clusters. For reference see [creating cluster with kubeadm](https://kubernetes.io/docs/setup/independent/create-cluster-kubeadm/). 
Need to determine how kubeadm and K8s/Cloud-providers should interact. More planning clearly needs to be done on 
cloud-provider and kubeadm planning.

#### CI, TestGrid and Other Testing Issues [WIP]

As soon as we have more than one repo involved in a product/build we get some interesting problems in running tests. A 
K8s/Cloud-provider deployment is the result of a combination of code from K8s/K8s and K8s/Cloud-provider. If something 
is broken it could be the result of a code change in either, and modifications by an engineer testing against one cloud 
provider could have unintended consequences in another. 

To address this issue, SIG-Testing has invested in a process by which participating cloud providers can report results 
from their own CI testing processes back to TestGrid, the Kubernetes community tool for tracking project health. 
Results become part of a central dashboard to display test results across multiple cloud providers. 

More thought should be put into managing multiple, independently evolving components. 

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
