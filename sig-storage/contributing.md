### Ramping up on Kubernetes Storage
For folks that prefer reading the docs first, we recommend reading our Storage Docs
- [The Persistent Volume Framework](http://kubernetes.io/docs/user-guide/persistent-volumes/) 
- [The new Dynamic Provisioning Proposal](https://github.com/pmorie/kubernetes/blob/7aa61dd0ff3908784acb4fa300713f02e62119af/docs/proposals/volume-provisioning.md) and [implementation](https://github.com/kubernetes/kubernetes/pull/29006)

For folks that prefer a video overview, we recommend watching the following videos:
- [The state of state](https://www.youtube.com/watch?v=jsTQ24CLRhI&index=6&list=PLosInM-8doqcBy3BirmLM4S_pmox6qTw3) 
- [Kubernetes Storage 101](https://www.youtube.com/watch?v=ZqTHe6Xj0Ek&list=PLosInM-8doqcBy3BirmLM4S_pmox6qTw3&index=38)
- [Overview of Basic Volume for SIG Apps](https://youtu.be/DrLGxkFdDNc?t=11m19s)
- [Overview of Dynamic Provisioning for SIG Apps](https://youtu.be/NXUHmxXytUQ?t=10m33s)

Keep in mind that the video overviews reflect the state of the art at the time they were created. In Kubernetes we try very hard to maintain backwards compatibility but Kubernetes is a fast moving project and we do add features going forward and attending the Storage SIG meetings and the Storage SIG Google group are both good ways of continually staying up to speed. 

### How to help

We love having folks help in any capacity! We recommend you start by reading the overall [Kubernetes contributors guide](https://github.com/GoogleCloudPlatform/continuous-deployment-on-kubernetes/blob/master/CONTRIBUTING.md)

### Helping with Features
If you have a feature idea, please submit a feature proposal PR first and put it on the [Storage SIG Meeting Agenda](https://docs.google.com/document/d/1-8KEG8AjAgKznS9NFm3qWqkGyCHmvU6HVl0sk5hwoAE/edit#heading=h.bag869lp4lyz). 
Our PR review bandwidth is fairly small, as such, we strongly recommend that you do not start writing the implementation before you've 
discussed the feature with the community. This helps the community understand what you're trying to do with the proposal and helps the 
community and you work through the approach until there is consensus. The community then will also be able to communicate with you how 
soon they will be able to review your proposal PR, to set expectations. However, generally speaking once the your proposal PR is merged, 
your implementation PR review and merge should go fairly quickly as the review is  focused on the implementation quality and not 
what you are proposing. We are really trying to improve our test coverage and documentation, so please include functional tests, e2e tests 
and documentation in your implementation PR.

### Helping with Issues
A great way to get involved is to pick an issue and help address it. We would love help here. Storage related issues are [listed here](https://github.com/kubernetes/kubernetes/labels/sig%2Fstorage)

### Adding support for a new storage platform in Kubernetes
For folks looking to add support for a new storage platform in Kubernetes, you have several options:
- Write an in-tree volume plugin or provisioner: You can contribute a new in-tree volume plugin or provisioner, that gets built and ships with Kubernetes, for use within the Persistent Volume Framework. 
[See the Ceph RBD volume plugin example](https://github.com/kubernetes/kubernetes/tree/master/pkg/volume/rbd) or [the AWS Provisioner example](https://github.com/kubernetes/kubernetes/pull/29006)
- Write a FlexVolume plugin: This is an out-of-tree volume plugin which you develop and build separately outside of Kubernetes. 
You then install the plugin on every Kubernetes host within your cluster and then [configure the plugin in Kubernetes as a FlexVolume](https://github.com/kubernetes/kubernetes/tree/master/examples/volumes/flexvolume)
- Write a Provisioner Controller: You can write a separate controller that watches for pending claims with a specific selector label on them. 
Once an appropriate claim is discovered, the controller then provisions the appropriate storage intended for the claim and creates a corresponding 
persistent volume for the claim that includes the same label used in the original claim selector. This will ensure that the PV for the new 
storage provisioned gets bound to the original claim.
