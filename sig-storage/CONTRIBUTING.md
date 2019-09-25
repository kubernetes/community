### Ramping up on Kubernetes Storage

We recommend the following presentations, docs, and videos to help get familiar with Kubernetes Storage concepts.

| Date | Title | Link | Description |
| --- | --- | --- | --- |
| - | Persistent Volume Framework | [Doc](http://kubernetes.io/docs/user-guide/persistent-volumes/) | Public user docs for Kubernetes Persistent Volume framework.
| 2018 May 03 | SIG Storage Intro | [Video](https://www.youtube.com/watch?v=GvrTl2T-Tts&list=PLj6h78yzYM2N8GdbjmhVU65KYm_68qBmo&index=164&t=0s) | An overview of SIG Storage By Saad Ali at KubeCon/CloudNativeCon EU 2018. |
| 2018 May 04 | Kubernetes Storage Lingo 101 | [Video](https://www.youtube.com/watch?v=uSxlgK1bCuA&t=0s&index=300&list=PLj6h78yzYM2N8GdbjmhVU65KYm_68qBmo) | An overview of various terms used in Kubernetes storage and what they mean by Saad Ali at KubeCon/CloudNativeCon EU 2018.|
| 2017 May 18 | Storage Classes & Dynamic Provisioning in Kubernetes |[Video](https://youtu.be/qktFhjJmFhg)| Intro to the basic Kubernetes storage concepts for users (direct volume reference, PV/PVC, and dynamic provisioning). |
| 2017 March 29 | Dynamic Provisioning and Storage Classes in Kubernetes |[Blog post](https://kubernetes.io/blog/2017/03/dynamic-provisioning-and-storage-classes-kubernetes/)| Overview of Dynamic Provisioning and Storage Classes in Kubernetes at GA. |
| 2017 March 29 | How Kubernetes Storage Works | [Slides](https://docs.google.com/presentation/d/1Yl5JKifcncn0gSZf3e1dWspd8iFaWObLm9LxCaXZJIk/edit?usp=sharing) | Overview for developers on how Kubernetes storage works for KubeCon/CloudNativeCon EU 2017 by Saad Ali
| 2017 February 17 | Overview of Dynamic Provisioning for SIG Apps | [Video](https://youtu.be/NXUHmxXytUQ?t=10m33s) | Overview of Storage Classes and Dynamic Provisioning for SIG Apps
| 2016 October 7 | Dynamic Provisioning and Storage Classes in Kubernetes |[Blog post](https://kubernetes.io/blog/2016/10/dynamic-provisioning-and-storage-in-kubernetes/)| Overview of Dynamic Provisioning and Storage Classes in Kubernetes at Beta. |
| 2016 July 26 | Overview of Basic Volume for SIG Apps | [Video](https://youtu.be/DrLGxkFdDNc?t=11m19s) | Overview of Basic Volume for SIG Apps
| 2016 March 25 | The State of State | [Video](https://www.youtube.com/watch?v=jsTQ24CLRhI&index=6&list=PLosInM-8doqcBy3BirmLM4S_pmox6qTw3) | The State of State at KubeCon/CloudNativeCon EU 2016 by Matthew Bates
| 2016 March 25 | Kubernetes Storage 101 | [Video](https://www.youtube.com/watch?v=ZqTHe6Xj0Ek&list=PLosInM-8doqcBy3BirmLM4S_pmox6qTw3&index=38) | Kubernetes Storage 101 at KubeCon/CloudNativeCon EU 2016 by Erin Boyd

Keep in mind that these artifacts reflect the state of the art at the time they were created. In Kubernetes we try very hard to maintain backwards compatibility, but Kubernetes is a fast moving project and we do add features going forward and attending the Storage SIG meetings and the Storage SIG Google group are both good ways of continually staying up to speed. 

### How to help

We love having folks help in any capacity! We recommend you start by reading the overall [Kubernetes contributor's guide](/contributors/guide)

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
A great way to get involved is to pick an issue and help address it. We would love help here. Storage related issues are [listed here](https://github.com/kubernetes/kubernetes/issues?utf8=%E2%9C%93&q=is%3Aopen+is%3Aissue+label%3Asig%2Fstorage+).

### Adding support for a new storage platform in Kubernetes
For folks looking to add support for a new storage platform in Kubernetes, you have several options:
- Write an in-tree volume plugin or provisioner: You can contribute a new in-tree volume plugin or provisioner, that gets built and ships with Kubernetes, for use within the Persistent Volume Framework. 
[See the Ceph RBD volume plugin example](https://git.k8s.io/kubernetes/pkg/volume/rbd) or [the AWS Provisioner example](https://github.com/kubernetes/kubernetes/pull/29006)
- Write a FlexVolume plugin: This is an out-of-tree volume plugin which you develop and build separately outside of Kubernetes. 
You then install the plugin on every Kubernetes host within your cluster and then [configure the plugin in Kubernetes as a FlexVolume](https://git.k8s.io/examples/staging/volumes/flexvolume)
- Write a Provisioner Controller: You can write a separate controller that watches for pending claims with a specific selector label on them. 
Once an appropriate claim is discovered, the controller then provisions the appropriate storage intended for the claim and creates a corresponding 
persistent volume for the claim that includes the same label used in the original claim selector. This will ensure that the PV for the new 
storage provisioned gets bound to the original claim.
