### Ramping up on Kubernetes Storage

We recommend the following presentations, docs, and videos to help get familiar with Kubernetes Storage concepts.

| Date | Title | Link | Description |
| --- | --- | --- | --- |
| 2023 January 23 | PV/PVC Controller Deep Dive | [Slides](https://docs.google.com/presentation/d/1XI_XzxlAYF6dqRfQQsoBd-UWRHssH2Ts5KwK5NDIMgg/edit) | A walk through of the PV controller source code following the Dynamic Volume Provisioning with CSI scenario |
| 2022 November 9 | Kubernetes SIG Storage Deep Dive | [Video](https://www.youtube.com/watch?v=_XXn3-yDZA0) | An overview and update of SIG Storage by Xing Yang, VMware & Mauricio Poppe, Google at KubeCon NA 2022. |
| 2022 May 16 | Kubernetes SIG Storage Deep Dive | [Video](https://www.youtube.com/watch?v=dsEeQqRSg74) | An overview and update of SIG Storage by Xing Yang, VMware & Jan Šafránek, Red Hat at KubeCon Europe 2022. |
| 2020 November 20 | Intro & Deep Dive: Kubernetes SIG-Storage | [Video](https://www.youtube.com/watch?v=rnCdvWToPPM&t=2s) | An overview and update of SIG Storage by Xing Yang and Michelle Au at KubeCon/CloudNativeCon NA 2020. |
| 2020 November 20 | Intro & Deep Dive: Kubernetes SIG-Storage | [Video](https://www.youtube.com/watch?v=rnCdvWToPPM&t=2s) | An overview and update of SIG Storage by Xing Yang and Michelle Au at KubeCon/CloudNativeCon NA 2020. |
| 2020 November 20 | Intro & Deep Dive: Kubernetes Data Protection WG | [Video](https://www.youtube.com/watch?v=g8HEQnLVo04) | An overview of Data Protection WG by Xing Yang and Xiangqian Yu at KubeCon/CloudNativeCon NA 2020. |
| 2020 November 18 | Beyond File and Block Storage in Kubernetes | [Video](https://www.youtube.com/watch?v=Y3GgJb71Cwo) | An introduction of Container Object Storage Interface (COSI) by Sidhartha Mani at KubeCon/CloudNativeCon NA 2020. |
| 2020 December 10 | Kubernetes 1.20: Kubernetes Volume Snapshot Moves to GA |[Blog post](https://kubernetes.io/blog/2020/12/10/kubernetes-1.20-volume-snapshot-moves-to-ga/)| Overview of Volume Snapshots in Kubernetes at GA. |
| - | Persistent Volume Framework | [Doc](http://kubernetes.io/docs/user-guide/persistent-volumes/) | Public user docs for Kubernetes Persistent Volume framework.
| 2019 May 23 | **Kubernetes Storage 101** | [Video](https://www.youtube.com/watch?v=_qfSzrPn9Cs) | An overview of usage of persistent storage in Kubernetes by David Zhu and Jan Šafránek at KubeCon/CloudNativeCon EU 2019.|
| 2018 May 03 | SIG Storage Intro | [Video](https://www.youtube.com/watch?v=GvrTl2T-Tts&list=PLj6h78yzYM2N8GdbjmhVU65KYm_68qBmo&index=164&t=0s) | An overview of SIG Storage By Saad Ali at KubeCon/CloudNativeCon EU 2018. |
| 2018 May 04 | **Kubernetes Storage Lingo 101** | [Video](https://www.youtube.com/watch?v=uSxlgK1bCuA&t=0s&index=300&list=PLj6h78yzYM2N8GdbjmhVU65KYm_68qBmo) | An overview of various terms used in Kubernetes storage and what they mean by Saad Ali at KubeCon/CloudNativeCon EU 2018.|
| 2017 May 18 | **Storage Classes & Dynamic Provisioning in Kubernetes** |[Video](https://youtu.be/qktFhjJmFhg)| Intro to the basic Kubernetes storage concepts for users (direct volume reference, PV/PVC, and dynamic provisioning). |
| 2017 March 29 | **Dynamic Provisioning and Storage Classes in Kubernetes** |[Blog post](https://kubernetes.io/blog/2017/03/dynamic-provisioning-and-storage-classes-kubernetes/)| Overview of Dynamic Provisioning and Storage Classes in Kubernetes at GA. |
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
For folks looking to add support for a new storage platform in Kubernetes, take a look of the [CSI Drivers Doc](https://kubernetes-csi.github.io/docs/). The CSI Drivers Doc website documents how to develop, deploy, and test a [Container Storage Interface](https://github.com/container-storage-interface/spec/blob/master/spec.md) (CSI) driver on Kubernetes. [CSI Hostpath Driver](https://github.com/kubernetes-csi/csi-driver-host-path/blob/master/README.md) is a demo implementation and is used for CSI testing.

Also see [here](https://github.com/kubernetes/community/blob/master/sig-storage/volume-plugin-faq.md) for deprecation notices regarding in-tree volume plugin and out-of-tree FlexVolume driver.
