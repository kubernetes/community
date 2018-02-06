## Refactor Cloud Provider out of Kubernetes Core

As kubernetes has evolved tremendously, it has become difficult for different cloudproviders (currently 7) to make changes and iterate quickly. Moreover, the cloudproviders are constrained by the kubernetes build/release life-cycle. This proposal aims to move towards a kubernetes code base where cloud providers specific code will move out of the core repository and into "official" repositories, where it will be maintained by the cloud providers themselves.

### 1. Current use of Cloud Provider

The following components have cloudprovider dependencies

    1. kube-controller-manager
    2. kubelet
    3. kube-apiserver

#### Cloud Provider in Kube-Controller-Manager

The kube-controller-manager has many controller loops

 - nodeController
 - volumeController
 - routeController
 - serviceController
 - replicationController
 - endpointController
 - resourceQuotaController
 - namespaceController
 - deploymentController
 - etc..

Among these controller loops, the following are cloud provider dependent.

 - nodeController
 - volumeController
 - routeController
 - serviceController

The nodeController uses the cloudprovider to check if a node has been deleted from the cloud. If cloud provider reports a node as deleted, then this controller immediately deletes the node from kubernetes. This check removes the need to wait for a specific amount of time to conclude that an inactive node is actually dead.

The volumeController uses the cloudprovider to create, delete, attach and detach volumes to nodes. For instance, the logic for provisioning, attaching, and detaching a EBS volume resides in the AWS cloudprovider. The volumeController uses this code to perform its operations.

The routeController configures routes for hosts in the cloud provider.

The serviceController maintains a list of currently active nodes, and is responsible for creating and deleting LoadBalancers in the underlying cloud.

#### Cloud Provider in Kubelet

Moving on to the kubelet, the following cloud provider dependencies exist in kubelet.

 - Find the cloud nodename of the host that kubelet is running on for the following reasons :
      1. To obtain the config map for the kubelet, if one already exists
      2. To uniquely identify current node using nodeInformer
      3. To instantiate a reference to the current node object
 - Find the InstanceID, ProviderID, ExternalID, Zone Info of the node object while initializing it
 - Periodically poll the cloud provider to figure out if the node has any new IP addresses associated with it
 - It sets a condition that makes the node unschedulable until cloud routes are configured.
 - It allows the cloud provider to post process DNS settings

#### Cloud Provider in Kube-apiserver

Finally, in the kube-apiserver, the cloud provider is used for transferring SSH keys to all of the nodes, and within an admission controller for setting labels on persistent volumes.

### 2. Strategy for refactoring Kube-Controller-Manager

In order to create a 100% cloud independent controller manager, the controller-manager will be split into multiple binaries.

1. Cloud dependent controller-manager binaries
2. Cloud independent controller-manager binaries - This is the existing `kube-controller-manager` that is being shipped with kubernetes releases.

The cloud dependent binaries will run those loops that rely on cloudprovider as a kubernetes system service. The rest of the controllers will be run in the cloud independent controller manager. 

The decision to run entire controller loops, rather than only the very minute parts that rely on cloud provider was made because it makes the implementation simple. Otherwise, the shared datastructures and utility functions have to be disentangled, and carefully separated to avoid any concurrency issues. This approach among other things, prevents code duplication and improves development velocity. 

Note that the controller loop implementation will continue to reside in the core repository. It takes in cloudprovider.Interface as an input in its constructor. Vendor maintained cloud-controller-manager binary could link these controllers in, as it serves as a reference form of the controller implementation. 

There are four controllers that rely on cloud provider specific code. These are node controller, service controller, route controller and attach detach controller. Copies of each of these controllers have been bundled them together into one binary. The cloud dependent binary registers itself as a controller, and runs the cloud specific controller loops with the user-agent named "external-controller-manager".

RouteController and serviceController are entirely cloud specific. Therefore, it is really simple to move these two controller loops out of the cloud-independent binary and into the cloud dependent binary.

NodeController does a lot more than just talk to the cloud. It does the following operations -

1. CIDR management
2. Monitor Node Status
3. Node Pod Eviction

While Monitoring Node status, if the status reported by kubelet is either 'ConditionUnknown' or 'ConditionFalse', then the controller checks if the node has been deleted from the cloud provider. If it has already been deleted from the cloud provider, then it deletes the nodeobject without waiting for the `monitorGracePeriod` amount of time. This is the only operation that needs to be moved into the cloud dependent controller manager.

Finally, The attachDetachController is tricky, and it is not simple to disentangle it from the controller-manager easily, therefore, this will be addressed with Flex Volumes (Discussed under a separate section below)

### 3. Strategy for refactoring Kubelet

The majority of the calls by the kubelet to the cloud is done during the initialization of the Node Object. The other uses are for configuring Routes (in case of GCE), scrubbing DNS, and periodically polling for IP addresses.

All of the above steps, except the Node initialization step can be moved into a controller. Specifically, IP address polling, and configuration of Routes can be moved into the cloud dependent controller manager.

Scrubbing DNS, after discussing with @thockin, was found to be redundant. So, it can be disregarded. It is being removed.

Finally, Node initialization needs to be addressed. This is the trickiest part. Pods will be scheduled even on uninitialized nodes. This can lead to scheduling pods on incompatible zones, and other weird errors. Therefore, an approach is needed where kubelet can create a Node, but mark it as "NotReady". Then, some asynchronous process can update it and mark it as ready. This is now possible because of the concept of Taints. 

This approach requires kubelet to be started with known taints. This will make the node unschedulable until these taints are removed. The external cloud controller manager will asynchronously update the node objects and remove the taints. 

### 4. Strategy for refactoring Kube-ApiServer

Kube-apiserver uses the cloud provider for two purposes

1. Distribute SSH Keys - This can be moved to the cloud dependent controller manager
2. Admission Controller for PV - This can be refactored using the taints approach used in Kubelet

### 5. Strategy for refactoring Volumes

Volumes need cloud providers, but they only need SPECIFIC cloud providers. The majority of volume management logic resides in the controller manager. These controller loops need to be moved into the cloud-controller manager. The cloud controller manager also needs a mechanism to read parameters for initialization from cloud config. This can be done via config maps.

There is an entirely different approach to refactoring volumes - Flex Volumes. There is an undergoing effort to move all of the volume logic from the controller-manager into plugins called Flex Volumes. In the Flex volumes world, all of the vendor specific code will be packaged in a separate binary as a plugin. After discussing with @thockin, this was decidedly the best approach to remove all cloud provider dependency for volumes out of kubernetes core.

### 6. Deployment, Upgrades and Downgrades

This change will introduce new binaries to the list of binaries required to run kubernetes. The change will be designed such that these binaries can be installed via `kubectl apply -f` and the appropriate instances of the binaries will be running. 

##### 6.1 Upgrading kubelet and proxy

The kubelet and proxy runs on every node in the kubernetes cluster. Based on your setup (systemd/other), you can follow the normal upgrade steps for it. This change does not affect the kubelet and proxy upgrade steps for your setup.

##### 6.2 Upgrading plugins

Plugins such as cni, flex volumes can be upgraded just as you normally upgrade them. This change does not affect the plugin upgrade steps for your setup.

###### 6.3 Upgrading kubernetes core services

The master node components (kube-controller-manager,kube-scheduler, kube-apiserver etc.) can be upgraded just as you normally upgrade them. This change does not affect the plugin upgrade steps for your setup.

##### 6.4 Applying the cloud-controller-manager

This is the only step that is different in the upgrade process. In order to complete the upgrade process, you need to apply the cloud-controller-manager deployment to the setup. A deployment descriptor file will be provided with this change. You need to apply this change using

```
kubectl apply -f cloud-controller-manager.yml
```

This will start the cloud specific controller manager in your kubernetes setup.

The downgrade steps are also the same as before for all the components except the cloud-controller-manager. In case of the cloud-controller-manager, the deployment should be deleted using

```
kubectl delete -f cloud-controller-manager.yml
```

### 7. Roadmap

##### 7.1 Transition plan

Release 1.6: Add the first implementation of the cloud-controller-manager binary. This binary's purpose is to let users run two controller managers and address any issues that they uncover, that we might have missed. It also doubles as a reference implementation to the external cloud controller manager for the future. Since the cloud-controller-manager runs cloud specific controller loops, it is important to ensure that the kube-controller-manager does not run these loops as well. This is done by leaving the `--cloud-provider` flag unset in the kube-controller-manager. At this stage, the cloud-controller-manager will still be in "beta" stage and optional.

Release 1.7: In this release, all of the supported turnups will be converted to use cloud controller by default. At this point users will still be allowed to opt-out. Users will be expected run the monolithic cloud controller binary. The cloud controller manager will still continue to use the existing library, but code will be factored out to reduce literal duplication between the controller-manager and the cloud-controller-manager. A deprecation announcement will be made to inform users to switch to the cloud-controller-manager.

Release 1.8: The main change aimed for this release is to break up the various cloud providers into individual binaries. Users will still be allowed to opt-out. There will be a second warning to inform users about the deprecation of the `--cloud-provider` option in the controller-manager.

Release 1.9: All of the legacy cloud providers will be completely removed in this version

##### 7.2 Code/Library Evolution

* Break controller-manager into 2 binaries. One binary will be the existing controller-manager, and the other will only run the cloud specific loops with no other changes. The new cloud-controller-manager will still load all the cloudprovider libraries, and therefore will allow the users to choose which cloud-provider to use.
* Move the cloud specific parts of kubelet out using the external admission controller pattern mentioned in the previous sections above.
* The cloud controller will then be made into a library. It will take the cloudprovider.Interface as an argument to its constructor. Individual cloudprovider binaries will be created using this library.
*  Cloud specific operations will be moved out of kube-apiserver using the external admission controller pattern mentioned above. 
* All cloud specific volume controller loops (attach, detach, provision operation controllers) will be switched to using flex volumes. Flex volumes do not need in-tree cloud specific calls. 
* As the final step, all of the cloud provider specific code will be moved out of tree. 
