---
kep-number: 2
title: Cloud Provider Controller Manager
authors:
  - "@cheftako"
  - "@calebamiles"
  - "@hogepodge"
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
  - "@hogepodge"
  - "@jagosan"
approvers:
  - "@thockin"
editor: TBD
status: provisional
replaces:
  - contributors/design-proposals/cloud-provider/cloud-provider-refactoring.md
---

# Remove Cloud Provider Code From Kubernetes Core

## Table of Contents

- [Remove Cloud Provider Code From Kubernetes Core](#remove-cloud-provider-code-from-kubernetes-core)
   - [Table of Contents](#table-of-contents)
   - [Summary](#summary)
   - [Motivation](#motivation)
      - [Goals](#goals)
      - [Intermediary Goals](#intermediary-goals)
      - [Non-Goals](#non-goals)
   - [Proposal](#proposal)
      - [Controller Manager Changes](#controller-manager-changes)
      - [Kubelet Changes](#kubelet-changes)
      - [API Server Changes](#api-server-changes)
      - [Volume Management Changes](#volume-management-changes)
      - [Deployment Changes](#deployment-changes)
      - [Implementation Details/Notes/Constraints](#implementation-detailsnotesconstraints)
          - [Repository Requirements](#repository-requirements)
              - [Notes for Repository Requirements](#notes-for-repository-requirements)
              - [Repository Timeline](#repository-timeline)
      - [Security Considerations](#security-considerations)
   - [Graduation Criteria](#graduation-criteria)
      - [Graduation to Beta](#graduation-to-beta)
         - [Process Goals](#process-goals)
   - [Implementation History](#implementation-history)
   - [Alternatives](#alternatives)
   
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
- [Kubelet](https://kubernetes.io/docs/reference/generated/kubelet/) - Track usages of [KubeletFlags.CloudProvider](https://github.com/kubernetes/kubernetes/blob/master/cmd/kubelet/app/options/options.go)
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

### Controller Manager Changes

For the controller manager we would like to create a set of common code which can be used by both the cloud controller
manager and the kube controller manager. The cloud controller manager would then be responsible for running controllers
whose function is specific to cloud provider functionality. The kube controller manager would then be responsible
for running all controllers whose function was not related to a cloud provider.

In order to create a 100% cloud independent controller manager, the controller-manager will be split into multiple binaries.

1. Cloud dependent controller-manager binaries
2. Cloud independent controller-manager binaries - This is the existing `kube-controller-manager` that is being shipped
with kubernetes releases.

The cloud dependent binaries will run those loops that rely on cloudprovider in a separate process(es) within the kubernetes control plane.
The rest of the controllers will be run in the cloud independent controller manager.
The decision to run entire controller loops, rather than only the very minute parts that rely on cloud provider was made
because it makes the implementation simple. Otherwise, the shared data structures and utility functions have to be
disentangled, and carefully separated to avoid any concurrency issues. This approach among other things, prevents code
duplication and improves development velocity.

Note that the controller loop implementation will continue to reside in the core repository. It takes in
cloudprovider.Interface as an input in its constructor. Vendor maintained cloud-controller-manager binary could link
these controllers in, as it serves as a reference form of the controller implementation.

There are four controllers that rely on cloud provider specific code. These are node controller, service controller,
route controller and attach detach controller. Copies of each of these controllers have been bundled together into
one binary. The cloud dependent binary registers itself as a controller, and runs the cloud specific controller loops
with the user-agent named "external-controller-manager".

RouteController and serviceController are entirely cloud specific. Therefore, it is really simple to move these two
controller loops out of the cloud-independent binary and into the cloud dependent binary.

NodeController does a lot more than just talk to the cloud. It does the following operations -

1. CIDR management
2. Monitor Node Status
3. Node Pod Eviction

While Monitoring Node status, if the status reported by kubelet is either 'ConditionUnknown' or 'ConditionFalse', then
the controller checks if the node has been deleted from the cloud provider. If it has already been deleted from the
cloud provider, then it deletes the nodeobject without waiting for the `monitorGracePeriod` amount of time. This is the
only operation that needs to be moved into the cloud dependent controller manager.

Finally, The attachDetachController is tricky, and it is not simple to disentangle it from the controller-manager
easily, therefore, this will be addressed with Flex Volumes (Discussed under a separate section below)


The kube-controller-manager has many controller loops. [See NewControllerInitializers](https://github.com/kubernetes/kubernetes/blob/release-1.9/cmd/kube-controller-manager/app/controllermanager.go#L332)

 - [nodeIpamController](https://github.com/kubernetes/kubernetes/tree/release-1.10/pkg/controller/nodeipam)
 - [nodeLifecycleController](https://github.com/kubernetes/kubernetes/tree/release-1.10/pkg/controller/nodelifecycle)
 - [volumeController](https://github.com/kubernetes/kubernetes/tree/release-1.9/pkg/controller/volume)
 - [routeController](https://github.com/kubernetes/kubernetes/tree/release-1.9/pkg/controller/route)
 - [serviceController](https://github.com/kubernetes/kubernetes/tree/release-1.9/pkg/controller/service)
 - replicationController
 - endpointController
 - resourceQuotaController
 - namespaceController
 - deploymentController
 - etc..

Among these controller loops, the following are cloud provider dependent.

 - [nodeIpamController](https://github.com/kubernetes/kubernetes/tree/release-1.10/pkg/controller/nodeipam)
 - [nodeLifecycleController](https://github.com/kubernetes/kubernetes/tree/release-1.10/pkg/controller/nodelifecycle)
 - [volumeController](https://github.com/kubernetes/kubernetes/tree/release-1.9/pkg/controller/volume)
 - [routeController](https://github.com/kubernetes/kubernetes/tree/release-1.9/pkg/controller/route)
 - [serviceController](https://github.com/kubernetes/kubernetes/tree/release-1.9/pkg/controller/service)

The nodeIpamController uses the cloudprovider to handle cloud specific CIDR assignment of a node. Currently the only
cloud provider using this functionality is GCE. So the current plan is to break this functionality out of the common 
verion of the nodeIpamController. Most cloud providers can just run the default version of this controller. However any
cloud provider which needs cloud specific version of this functionality and disable the default version running in the 
KCM and run their own version in the CCM.

The nodeLifecycleController uses the cloudprovider to check if a node has been deleted from/exists in the cloud. 
If cloud provider reports a node as deleted, then this controller immediately deletes the node from kubernetes. 
This check removes the need to wait for a specific amount of time to conclude that an inactive node is actually dead.
The current plan is to move this functionality into its own controller, allowing the nodeIpamController to remain in
K8s/K8s and the Kube Controller Manager.

The volumeController uses the cloudprovider to create, delete, attach and detach volumes to nodes. For instance, the
logic for provisioning, attaching, and detaching a EBS volume resides in the AWS cloudprovider. The volumeController
uses this code to perform its operations.

The routeController configures routes for hosts in the cloud provider.

The serviceController maintains a list of currently active nodes, and is responsible for creating and deleting
LoadBalancers in the underlying cloud.

### Kubelet Changes

Moving on to the kubelet, the following cloud provider dependencies exist in kubelet.

 - Find the cloud nodename of the host that kubelet is running on for the following reasons :
      1. To obtain the config map for the kubelet, if one already exists
      2. To uniquely identify current node using nodeInformer
      3. To instantiate a reference to the current node object
 - Find the InstanceID, ProviderID, ExternalID, Zone Info of the node object while initializing it
 - Periodically poll the cloud provider to figure out if the node has any new IP addresses associated with it
 - It sets a condition that makes the node unschedulable until cloud routes are configured.
 - It allows the cloud provider to post process DNS settings

The majority of the calls by the kubelet to the cloud is done during the initialization of the Node Object. The other
uses are for configuring Routes (in case of GCE), scrubbing DNS, and periodically polling for IP addresses.

All of the above steps, except the Node initialization step can be moved into a controller. Specifically, IP address
polling, and configuration of Routes can be moved into the cloud dependent controller manager.

[Scrubbing DNS was found to be redundant](https://github.com/kubernetes/kubernetes/pull/36785). So, it can be disregarded. It is being removed.

Finally, Node initialization needs to be addressed. This is the trickiest part. Pods will be scheduled even on
uninitialized nodes. This can lead to scheduling pods on incompatible zones, and other weird errors. Therefore, an
approach is needed where kubelet can create a Node, but mark it as "NotReady". Then, some asynchronous process can
update it and mark it as ready. This is now possible because of the concept of Taints.

This approach requires kubelet to be started with known taints. This will make the node unschedulable until these
taints are removed. The external cloud controller manager will asynchronously update the node objects and remove the
taints.

### API Server Changes

Finally, in the kube-apiserver, the cloud provider is used for transferring SSH keys to all of the nodes, and within an
admission controller for setting labels on persistent volumes.

Kube-apiserver uses the cloud provider for two purposes

1. Distribute SSH Keys - This can be moved to the cloud dependent controller manager
2. Admission Controller for PV - This can be refactored using the taints approach used in Kubelet

### Volume Management Changes

Volumes need cloud providers, but they only need **specific** cloud providers. The majority of volume management logic
resides in the controller manager. These controller loops need to be moved into the cloud-controller manager. The cloud
controller manager also needs a mechanism to read parameters for initialization from cloud config. This can be done via
config maps.

There are two entirely different approach to refactoring volumes -
[Flex Volumes](https://github.com/kubernetes/community/blob/master/contributors/devel/flexvolume.md) and
[CSI Container Storage Interface](https://github.com/kubernetes/community/blob/master/contributors/design-proposals/storage/container-storage-interface.md). There is an undergoing effort to move all
of the volume logic from the controller-manager into plugins called Flex Volumes. In the Flex volumes world, all of the
vendor specific code will be packaged in a separate binary as a plugin. After discussing with @thockin, this was
decidedly the best approach to remove all cloud provider dependency for volumes out of kubernetes core. Some of the discovery
information for this can be found at [https://goo.gl/CtzpVm](https://goo.gl/CtzpVm).

### Deployment Changes

This change will introduce new binaries to the list of binaries required to run kubernetes. The change will be designed
such that these binaries can be installed via `kubectl apply -f` and the appropriate instances of the binaries will be
running.

Issues such as monitoring, configuring the new binaries will generally be left to cloud provider. However they should
ensure that test runs upload the logs for these new processes to [test grid](https://k8s-testgrid.appspot.com/).

Applying the cloud controller manager is the only step that is different in the upgrade process.
In order to complete the upgrade process, you need to apply the cloud-controller-manager deployment to the setup.
A deployment descriptor file will be provided with this change. You need to apply this change using

```
kubectl apply -f cloud-controller-manager.yml
```

This will start the cloud specific controller manager in your kubernetes setup.

The downgrade steps are also the same as before for all the components except the cloud-controller-manager.
In case of the cloud-controller-manager, the deployment should be deleted using

```
kubectl delete -f cloud-controller-manager.yml
```

### Implementation Details/Notes/Constraints

#### Repository Requirements

**This is a proposed structure, and may change during the 1.11 release cycle.
WG-Cloud-Provider will work with individual sigs to refine these requirements
to maintain consistency while meeting the technical needs of the provider
maintainers**

Each cloud provider hosted within the `kubernetes` organization shall have a
single repository named `kubernetes/cloud-provider-<provider_name>`. Those
repositories shall have the following structure:

* A `cloud-controller-manager` subdirectory that contains the implementation
  of the provider-specific cloud controller.
* A `docs` subdirectory.
* A `docs/cloud-controller-manager.md` file that describes the options and
  usage of the cloud controller manager code.
* A `docs/testing.md` file that describes how the provider code is tested.
* A `Makefile` with a `test` entrypoint to run the provider tests.

Additionally, the repository should have:

* A `docs/getting-started.md` file that describes the installation and basic
  operation of the cloud controller manager code.

Where the provider has additional capabilities, the repository should have
the following subdirectories that contain the common features:

* `dns` for DNS provider code.
* `cni` for the Container Network Interface (CNI) driver.
* `csi` for the Container Storage Interface (CSI) driver.
* `flex` for the Flex Volume driver.
* `installer` for custom installer code.

Each repository may have additional directories and files that are used for
additional feature that include but are not limited to:

* Other provider specific testing.
* Additional documentation, including examples and developer documentation.
* Dependencies on provider-hosted or other external code.


##### Notes for Repository Requirements

This purpose of these requirements is to define a common structure for the
cloud provider repositories owned by current and future cloud provider SIGs.
In accordance with the 
[WG-Cloud-Provider Charter](https://docs.google.com/document/d/1m4Kvnh_u_9cENEE9n1ifYowQEFSgiHnbw43urGJMB64/edit#)
to "define a set of common expected behaviors across cloud providers", this
proposal defines the location and structure of commonly expected code.

As each provider can and will have additional features that go beyond expected
common code, requirements only apply to the location of the
following code:

* Cloud Controller Manager implementations.
* Documentation.

This document may be amended with additional locations that relate to enabling
consistent upstream testing, independent storage drivers, and other code with
common integration hooks may be added

The development of the 
[Cloud Controller Manager](https://github.com/kubernetes/kubernetes/tree/master/cmd/cloud-controller-manager)
and
[Cloud Provider Interface](https://github.com/kubernetes/kubernetes/blob/master/pkg/cloudprovider/cloud.go)
has enabled the provider SIGs to develop external providers that
capture the core functionality of the upstream providers. By defining the
expected locations and naming conventions of where the external provider code
is, we will create a consistent experience for:

* Users of the providers, who will have easily understandable conventions for
  discovering and using all of the providers.
* SIG-Docs, who will have a common hook for building or linking to externally
  managed documentation
* SIG-Testing, who will be able to use common entry points for enabling
  provider-specific e2e testing.
* Future cloud provider authors, who will have a common framework and examples
  from which to build and share their code base.

##### Repository Timeline

To facilitate community development, providers named in the 
[Makes SIGs responsible for implementations of `CloudProvider`](https://github.com/kubernetes/community/pull/1862)
patch can immediately migrate their external provider work into their named
repositories.

Each provider will work to implement the required structure during the
Kubernetes 1.11 development cycle, with conformance by the 1.11 release.
WG-Cloud-Provider may actively change repository requirements during the
1.11 release cycle to respond to collective SIG technical needs.

After the 1.11 release all current and new provider implementations must
conform with the requirements outlined in this document.

### Security Considerations

Make sure that you consider the impact of this feature from the point of view of Security.

## Graduation Criteria

How will we know that this has succeeded?
Gathering user feedback is crucial for building high quality experiences and SIGs have the important responsibility of
setting milestones for stability and completeness.
Hopefully the content previously contained in [umbrella issues][] will be tracked in the `Graduation Criteria` section.

[umbrella issues]: https://github.com/kubernetes/kubernetes/issues/42752

### Graduation to Beta

As part of the graduation to `stable` or General Availability (GA), we have set
both process and technical goals.

#### Process Goals

- 

We propose the following repository structure for the cloud providers which
currently live in `kubernetes/pkg/cloudprovider/providers/*`

```
git@github.com:kubernetes/cloud-provider-wg
git@github.com:kubernetes/cloud-provider-aws
git@github.com:kubernetes/cloud-provider-azure
git@github.com:kubernetes/cloud-provider-gcp
git@github.com:kubernetes/cloud-provider-openstack
```

We propose this structure in order to obtain

- ease of contributor on boarding and off boarding by creating repositories under
  the existing `kubernetes` GitHub organization
- ease of automation turn up using existing tooling
- unambiguous ownership of assets by the CNCF

The use of a tracking repository `git@github.com:kubernetes/wg-cloud-provider`
is proposed to

- create an index of all cloud providers which WG Cloud Provider believes
  should be highlighted based on defined criteria for quality, usage, and other
  requirements deemed necessary by the working group
- serve as a location for tracking issues which affect all Cloud Providers
- serve as a repository for user experience reports related to Cloud Providers
  which live within the Kubernetes GitHub organization or desire to do so

Major milestones:

- March 18, 2018: Accepted proposal for repository requirements.

*Major milestones in the life cycle of a KEP should be tracked in `Implementation History`.
Major milestones might include

- the `Summary` and `Motivation` sections being merged signaling SIG acceptance
- the `Proposal` section being merged signaling agreement on a proposed design
- the date implementation started
- the first Kubernetes release where an initial version of the KEP was available
- the version of Kubernetes where the KEP graduated to general availability
- when the KEP was retired or superseded*

The ultimate intention of WG Cloud Provider is to prevent multiple classes
of software purporting to be an implementation of the Cloud Provider interface
from fracturing the Kubernetes Community while also ensuring that new Cloud
Providers adhere to standards of quality and whose management follow Kubernetes
Community norms.

## Alternatives

One alternate to consider is the use of a side-car. The cloud-interface in tree could then be a [GRPC](https://github.com/grpc/grpc-go)
call out to that side-car. We could then leave the Kube API Server, Kube Controller Manager and Kubelet pretty much as is.
We would still need separate repos to hold the code for the side care and to handle cluster setup for the cloud provider.
However we believe that different cloud providers will (already) want different control loops. As such we are likely to need
something like the cloud controller manager anyway. From the perspective it seems easier to centralize the effort in that
direction. In addition it should limit the proliferation of new processes across the entire cluster.
