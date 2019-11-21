####Presenter:Ben the Elder


####Topic: Kind

####Date & time: 11/18/2019 2:00PM

####Notes

Kind allows you to start a docker image with all the Kubernetes binaries
basic workflow is a command that builds Kubernetes and puts it in a docker image.
After that its a Kubernetes cluster you just run tests because whats running is a full Kubernetes cluster.

Generally speaking you want to run the upstream Kubernetes tests.
You probably want to set a regex focus on test for your areas.

This runs on presubmit, and release blocking signal. 
Take the pr code build into an image 
create cluster from image
run battery of test after this

kind = automated Kubernetes building for testing strategies.

Does not test block device and it doesn't isolate the block device. 

Ignore privileged containers.

Not everything in Kubernets can be namespaces because of linux.

kind does not
    doesn't test block devices
    doesn't test the ingress 

Kubernetes uses klog package which is a port from google

kind build node-image

plan is to build a pre-submit test sweet to run a single command and pass a single test.

Kind only builds specific things it needs versus all of the binaries.

default is kind create cluster --image=/kindtest/default --name cluster_name

writes to default kubeconfig
you can change this by specifying --kubeconfig=config name

kubectl cluster-info --context kube-kubectl

kind delete cluster --name cluster_name

some parts of the builds are cached so after the initial build it will use cached items unless things have changed.

there is a config where you can specify granularly how you want your cluster to be set up.

storage class to handle multiple kind nodes but it's not available yet.

current options only put all workloads on the daemon doing the work

no true multi-node yet 

with sig-storage they are hoping there is a local storage driver that comes with it. The problem is the deployment mechanism for CSI. PR pending.

on docker with mac containers aren't access via ip.



Questions:

Has anyone tried to launch kind cluster in pod for testing?
You should not do this.
Kubernetes does do this but you shouldn't do this.
    secrets don't nest well
    you have to make sure you clean up
    you have to run a privileged container.
    
Is there any way to use it for training for spinning up multiple clusters.
no, and yes. possible with magic.


####Key learnings / takeaways

Kinda can be used to create cluster for testing and a project is in the works to do automated testing build into kind for release to check if prs that are merged are valid and will deploy fine.


####Action items
- [none] 
- []
- []
