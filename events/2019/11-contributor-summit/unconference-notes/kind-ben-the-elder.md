## Presenter: Ben the Elder


#### Topic: Kind

#### Date & time: 11/18/2019 2:00PM

#### Notes

Kind allows you to start a docker image with all the Kubernetes binaries
The basic workflow is a command that builds Kubernetes and puts it in a docker image.
After that it's a Kubernetes cluster you just run tests because what's running is a full kubernetes cluster.

Generally speaking you want to run the upstream kubernetes tests.
You probably want to set a regex focus on tests for your areas.

This runs on pre-submit and release blocking signal. 
Take the pr code build into an image, create cluster from image and run battery of test

kind = automated kubernetes building for testing strategies.

Does not test block device and it doesn't isolate the block device. 

Ignore privileged containers.

Not everything in Kubernetes can be namespaced because of linux.

kind does not

    * Test block devices
    * Test the ingress 

Kubernetes uses klog package which is a port from google

kind build node-image

Plan is to build a pre-submit test suite to run a single command and pass a single test.

Kind only builds specific things it needs versus all of the binaries.

Default is kind create cluster --image=/kindtest/default --name cluster_name

Writes to default kube config and you can change this by specifying --kubeconfig=config name

kubectl cluster-info --context kube-kubectl

kind delete cluster --name cluster_name

Some parts of the builds are cached so after the initial build it will use cached items unless things have changed.

There is a config where you can specify how you want your cluster to be set up.

Storage class to handle multiple kind nodes but its not available yet.

Current options only put all workloads on the daemon doing the work

Not true multi-node yet 

With sig-storage they are hoping there is a local storage driver that comes with it. The problem is the deployment mechanism for CSI, PR pending.

Docker with mac containers aren't access via ip.



Questions:

Has anyone tried to launch kind cluster in pod for testing?

    * You should not do this.

Kubernetes does do this but you shouldn't do this.
    
    * secrets don't nest well
    * you have to make sure you clean up
    * you have to run a privileged container.
    
Is there any way to use it for training for spinning up multiple clusters.
    
    * no, and yes. possible with magic.


#### Key Learning/Takeaways

Kind can be used to create cluster for testing and a project is in the works to do automated testing build into kind for release to check if prs that are merged are valid and will deploy fine.


#### Action items
- N/A