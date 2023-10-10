# SIG Node Meeting Notes

# Future {#future}



*   Monitoring pipeline proposal
*   [Virtual Kubelet (@rbitia)](https://docs.google.com/document/d/1MAn_HMZScni89hDwI4nQMk_SWTx9oi16PuOREhpVJJI/edit?usp=sharing)
*   PR Review Request - [#63170](https://github.com/kubernetes/kubernetes/pull/63170) - @micahhausler
*   Image name inconsistency in node status - [#64413](https://github.com/kubernetes/kubernetes/issues/64413) - @resouer
*   [Jess Frazelle / Kent Rancourt]: Proposal for kubelet feature to freeze pods placed in a hypothetical "frozen" state by the replication controller in response to a scale-down event. Enable pods to be thawed by socket activation.


# Dec 18



*   Graduate SupportPodPidsLimit to beta (Derek Carr)
    *   https://github.com/kubernetes/kubernetes/pull/72076
*   UserNS remapping: [updated proposal](https://github.com/kubernetes/community/pull/2595),[ implementation PR](https://github.com/kubernetes/kubernetes/pull/64005) (vikasc):

    Question raised at Zoom Chat:

*   From Mike Danese to Everyone: (10:42 AM)

     I would like to reconcile the "Motivation" section with how we want people to use users and groups in general even without remapping. We want tenants to run workloads as different users and groups to segment their disk access to improve security, and we want userns to support compatibility with images that expect to be running as a given uid/gid. The current proposal uses a file in /etc/... which makes getting both hard. Any path towards both these goals? yup exactly 

*   From Patrick Lang to Everyone: (10:44 AM)

     https://github.com/kubernetes/community/blob/bf48175c42fb71141f83071ce42f178d475b0bad/contributors/design-proposals/node/node-usernamespace-remapping.md#sandbox-type-runtimes  - this proposes using nodeselectors to select nodes that can support this. Should this be designed to work with RuntimeClass instead to avoid using node selectors?

*   From Mike Danese to Everyone: (10:45 AM)

     SIngle mapping on a node seems problematic IMO. Namespace would be better. Per pod would be best. I'd like to see that explored thanks 

*   From 刘澜涛 to Everyone: (10:46 AM)

     Just some background, it is possible to support per pod user namespace mapping in containerd.   It is already supported, we just need to add corresponding CRI implementation and kubelet support.



# Dec 11

Cancelled due to KubeCon Seattle


# Dec 4



*   Proposal: node shut down handling (Jing Xu)
    *   https://docs.google.com/document/d/1V7y9Boagev2LwQTSI1SaatMOdRC2gSQ9PBcESaITVEA
*   Mechanism to apply/monitor quotas for ephemeral storage: https://docs.google.com/presentation/d/1I9yYACVSBOO0SGB0ohvpTImJSCc2AaOTuO6Kgsg4IIU/edit?usp=sharing


# Nov 27



*   **Adding Termination Reason (OOMKilled/etc) Event or Counter (Brian)**
    *   Following discussion in sig-instrumentation, discuss either
        *   A. generating a logline / event for Pod termination reason or
        *   B. adding a termination reason counter in cadvisor/kubelet that can export to Prometheus
    *   Want to get better metrics when a container is killed.
    *   Want to get a count of all the termination reasons.
    *   Pod has last state, which has a termination reason
    *   kube-state-metrics has a termination reason, but as soon as the pod restarts it's gone
    *   would be nice to have a counter for how many containers have been killed for each possible reason (OOM, etc.)
    *   at minimum would be nice to have an event with the termination reason
        *   on top of this would be nice if cadvisor or kubelet could export a counter (e.g. for these events)
    *   Relevant Issue is: _"Log something about OOMKilled containers"_ https://github.com/kubernetes/kubernetes/issues/69676
*   **Node-Problem-Detector updates (wangzhen127)**
    *   adding a few new plugins:
        *   kubelet, docker, containerd crashlooping issues
        *   more checks on filesystem and docker overlay2 issues
    *   plans on making NPD easier to configure within k8s clusters
*   **wg-lts a new working group. [**Dhawal Yogesh Bhanushali dbhanushali@vmware.com**] \
**Mailing list: https://groups.google.com/forum/#!forum/kubernetes-wg-lts \
PR: https://github.com/kubernetes/community/pull/2911 \
Survey: https://docs.google.com/document/d/13lOGSSY7rz3yHMjrPq6f1kkRnzONxQeqUEbLOOOG-f4/edit

    slack: #wg-lts

*


# Nov 20

Meeting canceled due to USA holiday (Thanksgiving)


# Nov 13

no agenda


# Nov 6



*   Windows progress update [@patricklang]
    *   Tracking remaining v1.13 [stable] work at https://github.com/PatrickLang/k8s-project-management/projects/1

            As of 11/6, down to test & docs

    *   Release criteria doc, includes test case list: https://docs.google.com/document/d/1YkLZIYYLMQhxdI2esN5PuTkhQHhO0joNvnbHpW68yg8/edit#
*   Change of node label name(space) of NFD labels (@marquiz) \
https://github.com/kubernetes-incubator/node-feature-discovery/issues/176#issuecomment-436166692 \
(related to NFD repo migration: https://github.com/kubernetes-incubator/node-feature-discovery/issues/175)


# Oct 30



*   Node GPU Monitoring Demo ([dashpole@google.com](mailto:dashpole@google.com)) https://github.com/dashpole/example-gpu-monitor#example-gpu-monitor
    *   Feedback from Swati from Nvidia: the initial proposal with gRPC is very difficult to integrate. The new version with socket is doble.
    *   Next step: dashpole: Open PR to add and test the new socket endpoint in 1.13 \
Swati from Nvidia: work on integrating [DCGM exporter](https://github.com/NVIDIA/gpu-monitoring-tools/blob/master/exporters/prometheus-dcgm/dcgm-exporter/dcgm-exporter) with the endpoint.
*   Topology Manager (formerly known as NUMA Manager) proposal [cdoyle] \
https://github.com/kubernetes/community/pull/1680
    *   _TL;DR: Align topology-dependent resource binding within Kubelet._
    *   Agreed to move forward with the proposal.


# Oct 9



*   Move Node Feature Discovery to kubernetes-sigs (Markus Lehtonen) \
https://github.com/kubernetes-incubator/node-feature-discovery/issues/175
*   [RuntimeClass scheduling](https://docs.google.com/document/d/1W51yBNTvp0taeEss56GTk8jczqFJ2d6jBeN6sCSlYZU/edit#) (tallclair)


# Oct 02



*   1.13 release
*   Q4 planning discussion:

    https://docs.google.com/document/d/1HU6Ytm378IIw_ES3_6oQBoKxzYlt02DvOoOcQFkBcy4/edit?usp=sharing



#  \
Sept 18



*   Discuss release notes for 1.12 (Derek Carr)
    *   https://docs.google.com/document/d/1ZZkcIqDwUZiC77rhjA_XUGGsmAa9hyQ6nc5PXSX3LGI/edit
*   NUMA Manager Proposal Demo & Update (Louise Daly, Connor Doyle, Balaji Subramaniam)
    *   https://github.com/kubernetes/community/blob/4793277c981e7c7a5d9cbf1b2ab1003fc68384d3/contributors/design-proposals/node/numa-manager.md
    *   https://github.com/lmdaly/kubernetes/tree/dev/numa_manager - Demo Code


# Sept 11



*   Kubelet Devices Endpoint (device monitoring v2)
    *   [Slides](https://docs.google.com/presentation/d/1xz-iHs8Ec6PqtZGzsmG1e68aLGCX576j_WRptd2114g/edit?usp=sharing)
    *   [KEP](https://github.com/kubernetes/community/pull/2454)
*   Fix port-forward fo non-namespaced Pods (@Xu)
    *   https://docs.google.com/presentation/d/1x1B2_DFZ9VI2E_-pB2pzeYKtUxrK4XLXd7I0JWD1Z40/edit#slide=id.g4189217af3_0_58
    *   Related to: containerd shimv2 + kata etc
    *   some comments in the meeting:
        *   Kubelet doesn't want to see any containers/images visible to but not managed by kubelet. So if we want a solution like method 3, it should not visible to kubelet at all. And method 1 looks good to kubelet.
        *   There is a debug-container under discussion, which works quite similar to the method 3.
        *   Sig-node likes to revisit the port-forward method itself from the architecture level(?), however, the feature is required by many use cases, such as OpenShift, and it is essential.
*   RFC Improve node health tracking - https://github.com/kubernetes/community/pull/2640
    *   No discussion necessary at the moment


# Sept 4



*   Ephemeral storage (part 2)
    *   Discussion of moving ephemeral storage & volume management to the CRI: \
https://groups.google.com/d/topic/kubernetes-sig-storage/v2DKu8kNIgo/discussion
*   Interested mentor/mentees post 1.12?


# Aug 28



*   Discuss ephemeral storage quota enhancement (Robert Krawitz @Red Hat) https://docs.google.com/document/d/1ETuraEnA4UcMezNxSaEvxc_ZNF3ow-WyWlosAsqGsW0/edit?usp=sharing_eil&ts=5b7effb9


# Aug 21 proposed agenda



*   Windows GA updates (Patrick Lang @Microsoft)
    *   Just finished sig-windows , and meeting notes is at https://docs.google.com/document/d/1Tjxzjjuy4SQsFSUVXZbvqVb64hjNAG5CQX8bK7Yda9w/edit#
    *   Discussed with sig-network and sig-storage as suggested at sig-node several weeks ago
    *   Based on the current quality and testing, plan to have GA for 1.13.
*   Sidecar-container proposal: https://github.com/kubernetes/community/pull/2148 (Joseph)
*   Kata & Container shim v2 Integration updates ([Xu@hyper.sh](mailto:Xu@hyper.sh), https://docs.google.com/presentation/d/1icEJ77idnXrRSj-mSpAkmy9cCpD1rIefPFrbKQnyB3Q/edit?usp=sharing)
    *   Shim V2 Proposal: https://github.com/containerd/containerd/issues/2426
*   Device Assignment Proposal: https://github.com/kubernetes/community/pull/2454


# Aug 14



*   Summary on the [New Resource API](https://github.com/kubernetes/community/pull/2265) [Follow-up](https://docs.google.com/document/d/1iWlfyYG781UVXCLEnpcFGNecutD4kALVcUJ0jqY7-_k) offline discussions (@vikaschoudhary16, @jiayingz, @connor, @renaud):
    *   We will start with multiple-matching model without priority field.
    *   We will start with allowing ResourceClass mutation.
*   Summary on the [New Resource API](https://github.com/kubernetes/community/pull/2265) [user feedbacks](https://docs.google.com/document/d/1syxE8dwsUde5BuHuVibrJIxH6GfhTp2JYOM_cgNVrAI/) (@vikaschoudhary16, @jiayingz):
    *   tldr: we received feedbacks from multiple HW vendors, admins who manage large enterprise clusters, and k8s providers that represent large customer sets that the new Resource API a useful feature that will help unblock some of their important use cases


# Aug 07



*   [Follow-up](https://docs.google.com/document/d/1iWlfyYG781UVXCLEnpcFGNecutD4kALVcUJ0jqY7-_k) [New Resource API ](https://github.com/kubernetes/community/pull/2265)KEP proposal (@vikaschoudhary16 and @jiayingz):
    *   Using `Priority` field for handling overlapping res classes
    *   Should res classes with same priority be allowed?
    *   non-mutable or mutable
*   Device plugin 1.12 enhancement plan: (@jiayingz)
    *   https://docs.google.com/document/d/1evJDu6H-LowS5saVmwGKOpEWA8DtDn0QpbSmzX3iihY
*   Add list of CSI drivers to Node.Status (@jsafrane)
    *   https://github.com/kubernetes/community/pull/2487
    *   sig-node will review
    *   Major concern right away: node object is already too big. Jan will get API approval +  sig-architecture approval to remove the old annotation (to save space).
    *   Will continue next week.


# July 31



*   [tstclair & jsb] CRI versions and validation
    *   Docker 18.03-ce is the only installable version for Ubuntu 18.04
    *   Validate against a docker API version: https://github.com/kubernetes/kubernetes/issues/53221
    *   Sig-cluster-lifecycle requirements:
        *   Test dashboard to show the container runtime status:
            *   Docker: https://k8s-testgrid.appspot.com/sig-node-kubelet, https://k8s-testgrid.appspot.com/sig-node-cri
            *   Containerd: https://k8s-testgrid.appspot.com/sig-node-containerd
            *   CRI-O: https://k8s-testgrid.appspot.com/sig-node-cri-o
            *   Pending work to move to the newly defined test jobs for CRI
        *   A central place of document to tell users how to configure each container runtime.
            *   Follow up here - https://github.com/kubernetes/website/issues/9692
*   1.12 Feature Freeze Review (Dawn/Derek)
    *   Planning doc for discussion
        *   https://docs.google.com/document/d/1m4Jzcd2p364kt2aMLSHZys7_02_oI3uwNkObrRXBO3k/edit
    *   List of feature issues opened
        *   https://github.com/kubernetes/features/issues?q=is%3Aissue+is%3Aopen+label%3Asig%2Fnode


# July 24



*   Agenda Topic (owner)
*   [tstclair & jsb] CRI versions and validation
    *   Docker 18.03-ce is the only installable version for Ubuntu 18.04


# July 17



*   [Device "Scheduling](https://docs.google.com/document/d/1Gad4s8BaFmUQ0JeYdJcXr6sek_q-1mjoKkFPaRQL6oA/edit?usp=sharing)" proposal (@dashpole) [slides](https://docs.google.com/presentation/d/1xz-iHs8Ec6PqtZGzsmG1e68aLGCX576j_WRptd2114g/edit?usp=sharing)
*   [New Resource API ](https://github.com/kubernetes/community/pull/2265)KEP proposal (@vikaschoudhary16 and @jiayingz)
    *   Discussed the user stories we would like to enable through this proposal
    *   Discussed goals and non-goals. In particular, discussed why we proposed to use ResourceClass to express resource property matching constraints instead of directly expressing such constraints in Pod/Container spec. These reasons are documented in the non-goal section
    *   Question on whether it is common for people to create a cluster with multiple types of GPU. Or do people usually create multiple clusters, each maps to one GPU type.
        *   Answer: Yes. We have seen such customers from Openshift, GKE, and on-prem.
    *   Bobby: When we start to support group resources, and ComputeResource to ResourceClass matching behavior becomes many-to-many, that would bring quite a lot complexity and scaling concerns on scheduler. People have already seen scaling problems on scheduler when they use very big number of extended resources.
        *   Answer: scaling is definitely a concern. We don't plan to support group resource in the initial phase. We need to collect scaling and performance numbers and carefully evaluate them after the initial phase before moving forward. We should publish performance numbers and best-practice guidelines to advise people from mis-using the building blocks we put in. We should also think about the necessary mechanism to prevent people from getting into "bad" situation.


# July 10



*   Follow-up from June 26
    *   Move cri-tools to kubernetes-sigs. https://github.com/kubernetes-incubator/cri-tools/issues/331
    *   Move cri-o to kubernetes-sigs (Mrunal/Antonio). \
https://github.com/kubernetes-incubator/cri-o/issues/1639
    *   Follow-up
        *   no disagreement in sig to move out of incubator org into sigs org
        *   dawn/derek to ack on issues and initiate transfer
*   Future of Node feature discovery (Markus): \
https://docs.google.com/document/d/1TXvveLiA_ByQoHTlFWtWCxz_kwlXD0EO9TqbRLTiOSQ/edit#
*   [RuntimeClass](https://github.com/kubernetes/community/pull/2290) follow up: s/Parameters/RuntimeHandler/ (tallclair)
*   Float [command refactor](https://github.com/kubernetes/kubernetes/issues/26093#issuecomment-403581393) idea (tallclair)
*   [New Resource API ](https://github.com/kubernetes/community/pull/2265)KEP proposal (@vikaschoudhary16 and @jiayingz)
*   [Device "Scheduling](https://docs.google.com/document/d/1Gad4s8BaFmUQ0JeYdJcXr6sek_q-1mjoKkFPaRQL6oA/edit?usp=sharing)" proposal (@dashpole)


# July 3

Cancelled


# June 26 {#june-26}



*   Move cri-tools to kubernetes-sigs. https://github.com/kubernetes-incubator/cri-tools/issues/331
*   Move cri-o to kubernetes-sigs (Mrunal/Antonio). \
https://github.com/kubernetes-incubator/cri-o/issues/1639
*   RuntimeClass KEP: https://github.com/kubernetes/community/pull/2290
    *   Derek: How will RuntimeClass will be used? What is the scope?
    *   Derek: Can we have a concrete example to prove the api works?
    *   Derek: Should unqualified image name normalization problem solved by RuntimeClass? - If yes, it is not only container runtime need to handle it, both Kubelet and scheduler need to normalize image name based on the RuntimeClass.
    *   How is storage and networking going to work in this sandbox model?
*   Future of Node feature discovery (Markus): \
https://docs.google.com/document/d/1TXvveLiA_ByQoHTlFWtWCxz_kwlXD0EO9TqbRLTiOSQ/edit#


# June 19 {#june-19}



*   [User Capabilities](https://github.com/filbranden/kubernetes-community/blob/usercap1/contributors/design-proposals/node/user-capabilities.md) (using ambient capabilities for non-root in container) [@filbranden](https://github.com/filbranden) \
PR [kubernetes/community#2285](https://github.com/kubernetes/community/pull/2285)


# June 12 {#june-12}



*   Sandboxes API Decision (@tallclair)
    *   Proposed to move forward with RuntimeClass, instead of sandbox boolean
    *   Some rationale behind the proposal
        *   Unblocks runtime extensions beyond sandboxes like kata vs. gVisor, for example,  windows
        *   Provides a clean way of specifying pod overhead and sandbox configuration
    *
*   APIsnoop e2e test to API mapping (@hh and @rohfle)
    *   https://apisnoop.cncf.io
    *   https://github.com/cncf/apisnoop/issues/17
        *   https://github.com/cncf/apisnoop/tree/master/dev/e2e-audit-correlate/results/20180605-e2e-kubetest/
    *   [slides](https://docs.google.com/presentation/d/1wrdBlLtHb_z5qmNwDDPrc9DRDs3Klpac83v8h5iAqjE/edit?usp=sharing)


# June 5 {#june-5}

Cancelled


# May 29 {#may-29}



*   [Sandboxes API follow-up discussion](https://groups.google.com/d/topic/kubernetes-sig-node/rHLECuaXGJs/discussion)


# May 22 {#may-22}



*   [Sandboxes API Proposal](https://docs.google.com/document/d/1WzO_QjJFfedhsiBtfcVB2QzTWRXHEPX1xOyqDGXxO-0/edit#) (@tallclair)
    *   Q: Do we expect cluster admin to label nodes which support sandbox?
        *   It's not clear whether we should require all container runtimes to support sandbox, or this is an optional feature.
        *   This doesn't block alpha.
        *   Ideally each node will only support one sandbox technology, but if the user want to use different sandbox technology, they may need to have different node setup and label the node.
    *   Q: Will this be included in conformance test?
    *   Q: Can the sandbox be a runc implementation which meets the conformance?
        *   It is hard to define a conformance test suite to validate whether the sandbox meets the requirement.
        *   We can only provide guideline for now.
        *   We are still working on the definition, but one probably definition: idependent kernel, and 2 layers of security?
    *   If the sandbox definition is not clear, derek is worried that:
        *   User may never or always set Sandbox=true.
        *   If sandbox is only a boolean, when `sandbox=true`, a workload may work for some sandbox technology, but not work for another sandbox technology.
    *   Q: Why not just expose OCI compatible runtime? So that we can get similar thing with sandbox boolean, and more flexibility.
        *   For example, we can have an admission controller to do the validation and set defaults for an OCI compatible runtime, which is not necessarily to be in core Kubernetes.
    *   Q: What is the criteria to graduate this to beta/GA?
        *   The alpha version is mainly to get user feedback, it is not finalized api.
    *   Q: Is this a Kubernetes 1.13 evolution? Or a longer term project?
        *   A: 1.13 would be good if there is strong signal that this works, but expect this to take longer. Would like to see this gets into alpha in 1.12.
    *   Q: Is Selinux blocked when sandbox=true?
        *   AppArmor, Selinux are not blocked, guestOS can support to provide better isolation between containers inside the sandbox.
    *   Q: Instead of changing all "Sandbox" in CRI, can we rename to a more specific name "Kernel-Isolated" or something, which also makes the api less confusing? "Sandbox" is too vague.


# May 15 {#may-15}



*   [Monitoring proposal update](https://docs.google.com/document/d/1NYnqw-HDQ6Y3L_mk85Q3wkxDtGNWTxpsedsgw4NgWpg/edit?usp=sharing) (@dashpole)
*   CFS quota change discussion : https://github.com/kubernetes/kubernetes/pull/63437
    *   Dawn to sync with Vish, we seemed good with a knob at node level to tune period on call.
*   Cheaper heartbeats: https://github.com/kubernetes/community/pull/2090
*   Update on 1.11 development feature progress (reviews, etc.)
    *   https://docs.google.com/document/d/1rtdcp4n3dTTxjplkNvPDgAW4bbB5Ve4ZsMXBGOYKiP0/edit
    *   sysctls to Beta fields need reviews (Jan)
        *   https://github.com/kubernetes/kubernetes/pull/63717
*   CRI container log path:
    *   Only add pod name and namespace into the log path to avoid regression: https://github.com/kubernetes/kubernetes/issues/58638
    *   For more metadata support e.g. namespace uid, need more thoughts, let's do it in the future.
        *   e.g. We can potentially use fluentd output filter plugin (e.g. https://github.com/bwalex/fluent-plugin-docker-format) + node metadata endpoint proposed by @dashpole.


# May 8 {#may-8}



*   Cancelled due to KubeCon and OpenShift


# May 1 {#may-1}



*   Sysctl move to beta (@ingvagabund) (Seth)
    *   https://github.com/kubernetes/community/pull/2093
    *   [Discuss KEP for user namespace (proposal, @mrunal, @vikas, @adelton)](https://github.com/kubernetes/community/pull/2067/files)
    *   Add probe based mechanism for kubelet plugin discovery: (@vikasc)

        https://github.com/kubernetes/kubernetes/pull/63328



# April 24 {#april-24}



*   [Discuss KEP for user namespace (proposal, @mrunal, @vikas, @adelton)](https://github.com/kubernetes/community/pull/2067/files)
*   Online volume resizing ([proposal](https://github.com/kubernetes/community/pull/1535), @gnufied)
*   Maximum volume limit ([proposal](https://github.com/kubernetes/community/pull/2051), gnufied)
*   Sig-Node Initial Charter: (https://github.com/kubernetes/community/pull/2065, @derekwaynecarr)
*   Add probe based mechanism for kubelet plugin discovery: (@vikasc)

    https://github.com/kubernetes/kubernetes/pull/59963



# April 17 {#april-17}



*   Re-categorize/tag Node E2E tests ([proposal](https://docs.google.com/document/d/1BdNVUGtYO6NDx10x_fueRh_DLT-SVdlPC_SsXjYCHOE/edit?usp=sharing), @yujuhong)
    *   We need to consider whether to include test like eviction test into conformance suite. They are essential features, but the test itself has some special requirement 1) it needs to be run in serial; 2) it does make some assumptions such as the node size.
*   [Virtual Kubelet (@rbitia)](https://docs.google.com/document/d/1MAn_HMZScni89hDwI4nQMk_SWTx9oi16PuOREhpVJJI/edit?usp=sharing) Updates
*   Add probe based mechanism for kubelet plugin discovery: https://github.com/kubernetes/kubernetes/pull/59963


# April 10 {#april-10}



*   Windows Service Accounts in Kubernetes (@patricklang)
    *   https://github.com/kubernetes/kubernetes/issues/62038
    *   How to support Windows Group Managed Service Account (gMSA)
        *   WindowsCredentialSpec is windows specific identity
        *   Option 1: WindowsCredentialSpec references a secret
        *   Option 2: WindowsCredentialSpec contains the JSON passed as-is in the OCI spec
        *   No password information inside, only some information about the host.
*   Sandbox resource management - [discussion questions](https://docs.google.com/document/d/1FtPFc8hhAVBWwTvU0BWrn34Salh2zZk0GXJz8ilN8bw/edit#)
    *   Pod resources design proposal: https://docs.google.com/document/d/1EJKT4gyl58-kzt2bnwkv08MIUZ6lkDpXcxkHqCvvAp4/edit
    *   If user set the overhead, how should runtime controller deal with it?
        *   Reject the pod, may cause problem on upgrade.
        *   Overwrite it.
    *   Derek: Can the node advertise/register the overhead to the control plane? Can node report per pod overhead in the node status? And let scheduler make decision based on that?
    *   Derek: Can we make kata as a node resource? And pod wants to run on kata needs to request that resource?
    *   Derek: Pod level limit is easier to set by user. - Dawn: It may not be true in the future based on the trend of Kubernetes today, e.g. Sidecar containers injected by Istio. - Derek: We can resolve the istio issue by dynamically resizing the limit.
    *   Dawn: Do we expect mixed sandbox and non-sandbox containers on a node? - Tim: Yes. The reason is that we do have system containers can't run in sandbox, e.g. fluentd, kube-proxy etc.
    *   Derek: Do we expect multiple runtimes on a single node? e.g. containerd, cri-o etc. - Dawn: No. Each node should have one container runtime, but it supports running container with different OCI runtimes, e.g. runc, kata etc.
    *   Derek: What happens to the best effort pod? - Dawn: This is an issue for frakti today, so they still set fixed limit.
*   Container runtime stream server authentication https://github.com/kubernetes/kubernetes/issues/36666#issuecomment-378440458
    *
*   Add probe based mechanism for kubelet plugin discovery: https://github.com/kubernetes/kubernetes/pull/59963
*   Node-level checkpointing:

https://github.com/kubernetes/kubernetes/pull/56040



*   sig-node charter discussion
    *   wip: https://github.com/derekwaynecarr/community/blob/sig-node-charter/sig-node/governance.md


# April 3 {#april-3}



*   Resource Management F2F update
    *   [Meeting notes](https://docs.google.com/document/d/1Df6uGCzGleAhRQYZ20v55U1YCBLfxV4PJh0CIWcMHkk/edit#heading=h.lc5c2ccmeh4i)
*   Q2 planning draft (@dawnchen)
    *   https://docs.google.com/document/d/1rtdcp4n3dTTxjplkNvPDgAW4bbB5Ve4ZsMXBGOYKiP0/edit?usp=sharing
*   Node Controller question - https://github.com/kubernetes/kubernetes/issues/61921
    *   Node deduplicates multiple InternalIP addresses on a node
    *   ask Walter Fender (@cheftako)


# March 27 {#march-27}



*   [reduce scope of node on taints/labels (@mikedanese)](https://github.com/kubernetes/community/pull/911)
*   [Node TLS bootstrap via TPM](https://docs.google.com/document/d/12UgErB_46iHBOi0YEKbaFbEXv9E5O6XWQihwPkwB_7s) (@awly)


# March 20 {#march-20}



*   Secure Container Use Cases updates (@tallclair)
    *   Strongly suggesting pod as the security boundary, @tallclair to write up decision (welcoming feedback and/or disagreement)
    *   Extra overhead of sandbox. Charge to who? pod-level resource request / limits? See updates on the [solution space doc](https://docs.google.com/document/d/1QQ5u1RBDLXWvC8K3pscTtTRThsOeBSts_imYEoRyw8A/edit#heading=h.6xmg7pt28mrd)
    *   Need to think more about volume model in the sandboxes to enforce 2 security boundaries (and prevent issues like CVE-2017-100210[12]
    *   sysctl: need to review the support list.
    *   device plugins - needs more consideration, [contributions welcome](https://docs.google.com/document/d/1QQ5u1RBDLXWvC8K3pscTtTRThsOeBSts_imYEoRyw8A/edit#heading=h.uqm3zexmm967)
*   [Proposal: Pod Ready++](https://docs.google.com/document/d/1VFZbc_IqPf_Msd-jul7LKTmGjvQ5qRldYOFV0lGqxf8/edit?ts=5aa1d7b4#heading=h.wvpldvtvh8u4) (@freehan)
*   [containerd f2f meeting update (@lantaol)](https://docs.google.com/document/d/1MrgDYOSTjysMPcc6D7OnaeIDEc48lMjQAB10EIKQ5Go/edit?usp=sharing)
*   [Virtual Kubelet (@rbitia)](https://docs.google.com/document/d/1MAn_HMZScni89hDwI4nQMk_SWTx9oi16PuOREhpVJJI/edit?usp=sharing)
    *   Doc posted during the meeting.  Please read for next week.
    *   https://docs.google.com/document/d/1vhrbB6vytFJxU6CrlMqlH6wC3coQOwtsIj-aFzBbPXI/edit#heading=h.dshaptx6acc from Ben VMWare.
    *   Review the design next week
*   [RawProc options for CRI (@jessfraz)](https://github.com/kubernetes/community/pull/1934)
    *   Add use cases and how this will be incorporated in the kubernetes API


# March 13 {#march-13}



*   Containerd F2F meeting notes: https://docs.google.com/document/d/1MrgDYOSTjysMPcc6D7OnaeIDEc48lMjQAB10EIKQ5Go/edit?usp=sharing


# Mar 6 {#mar-6}



*   CRI-O status updates (@mrunal)
    *   slide update: https://docs.google.com/presentation/d/1TQ6sBo63AXt6QF3LxA3jtnNzT330MpPDODORhLVcDH0/edit?ts=5a9ed52a#slide=id.g3134b94d16_0_0
    *   pr help needed for dashboard: https://github.com/kubernetes/test-infra/pull/5943
    *
*   Official Contributor Experience Reach-Out (@jberkus)
    *   Process for label and PR workflow changes
        *   [Issue Triage](https://github.com/kubernetes/community/blob/master/contributors/guide/issue-triage.md)
    *   Mentoring!
        *   [Meet our contributors](https://github.com/kubernetes/community/blob/master/mentoring/meet-our-contributors.md)
        *   [Group mentoring](https://github.com/kubernetes/community/blob/master/mentoring/group-mentoring.md)
            *   why: https://k8s.devstats.cncf.io/d/QQN85o3zk/pr-workload-table?orgId=1
            *   https://k8s.devstats.cncf.io/d/46/user-reviews?orgId=1
    *   [Contributor Guide](https://github.com/kubernetes/community/tree/master/contributors/guide)
    *   [Governance and charters](https://github.com/kubernetes/community/tree/master/committee-steering/governance)
    *   What can Contribex do for you?
*


# Feb 27 {#feb-27}



*   Virtual Kubelet implementation doc: (@rbitia) https://docs.google.com/document/d/1tu27_BquhUAmYLaJznjbbdLJKYrH5-stm1i0OcxRgE8/edit?usp=sharing


# Feb 20 {#feb-20}



*   Secure isolation discussion continue (@tallclair)
    *   https://docs.google.com/document/d/1QQ5u1RBDLXWvC8K3pscTtTRThsOeBSts_imYEoRyw8A/edit?usp=sharing


# Feb 13 {#feb-13}



*   Secure isolation update
    *   Owner: @tallclair
    *   https://docs.google.com/document/d/1QQ5u1RBDLXWvC8K3pscTtTRThsOeBSts_imYEoRyw8A/edit?usp=sharing
    *   Introduce solution space doc, goals & high-level overview
    *   Discuss in a future meeting, after folks have time to digest the document
*   Collection of options to integrate kata with CRI
    *   Owner:  @resouer, ~20 mins
    *   Doc: https://docs.google.com/document/d/1PgXJpzSfhR_1idkNtcZNuncfUYV-U-syO-HSGxwQrVw/edit?usp=sharing
    *   Introduce the existing & proposed ways of integrating Kata with CRI, a brief evaluation will also be included.


# Feb 6 {#feb-6}



*   CRI: testing policy
    *   Owner: yujuhong@
    *   https://github.com/kubernetes/community/pull/1743
*   CRI: container runtime cgroup detection (or not)
    *   Owner: yujuhong@, lantaol@
    *   https://github.com/kubernetes/kubernetes/issues/30702
    *   The runtime stats is only used for monitoring today.
    *   Do not change CRI for this. Instead, the runtime cgroup can be passed to the kubelet through the existing flag `runtime-cgroups`.
*   CRI: Image Filesystem Storage Identifier
    *   Owner: yujuhong@, lantaol@
    *   https://github.com/kubernetes/kubernetes/issues/57356
    *   Slides: https://docs.google.com/presentation/d/1BbXgmEVhH0p2cgoojN36Q6SMHZ8tl00kfjITmYEf9vM/edit?usp=sharing
    *   Kubelet can use `statfs` to get fs stats. It is simple enough.
    *   Kubelet get image filesystem path through CRI vs. Configure image filesystem path through kubelet flag? The latter seems to be preferred.
*   Release reminder: sig-node's v1.10 marked issues and features need scrubbed.  We're officially past feature freeze and code freeze is coming soon.  [More details in email on kubernetes-dev today](https://groups.google.com/forum/#!topic/kubernetes-dev/Jt4hbwzZrbA) and [new issue list](https://docs.google.com/document/d/11u91ypj8Gt8PlTincWuQ3iB2X3tITBxqn6JMkTduEZw/edit#heading=h.30996wkvo4bo) this week is largely SIG-Node's.
*   Reminder: We are close (hopefully this week) to graduating the Kubelet's componentconfig API to beta status, as our existing TODO list is complete. Please take a look at https://github.com/kubernetes/kubernetes/pull/53833.
*   Reminder: NamespaceOption is changing in the CRI with [#58973](https://github.com/kubernetes/kubernetes/pull/58973), runtimes will break.
    *


# Jan 30 {#jan-30}



*   Revised Windows container resource management (CPU & memory):
    *   **Owner**: Jiangtian Li (jiangtli@microsoft.com), Patrick Lang (plang@microsoft.com)
    *   https://github.com/kubernetes/community/pull/1510
    *   What about container without resource request/limit? - Zero value passed to Docker, and Docker will set an value based on the current node usage.
    *   CPU assignment? - Not supported in windows platform now.
    *   Is in place resource update supported on Windows? - Immutable today.
*   Virtual Kubelet introduction
    *   **Owner:** Ria Bhatia, [ria.bhatia@microsoft.com](mailto:ria.bhatia@microsoft.com)
    *   **Context:**

        [Virtual Kubelet](https://github.com/Virtual-Kubelet/Virtual-Kubelet) is an interface for plugging in _anything, _into a K8 cluster_, _to replicate the lifecycle of a kubernetes pod. I'd like to kick off discussion for the concepts of VK, the design of the project itself, networking scenarios and also come up with a formal provider definition.

*   Node e2e conformance tests: https://github.com/kubernetes/kubernetes/issues/59001
    *   [Spreadsheet](https://docs.google.com/spreadsheets/d/1Af7zSLDEDM3i5g-8MZi0oZ7zPt85rWI1kHMQybhS9H0/edit?usp=sharing) to re-categorize tests: please comment if you think a test should be added to or removed from the conformance suite. Please also suggest feature tags for tests.


# Jan 23 {#jan-23}



*   [Invoke Allocate RPC call at container creation time#58282](https://github.com/kubernetes/kubernetes/pull/58282)
    *   **Owner:** @RenaudWasTaken <rgaubert@nvidia.com>
    *   **Context:** \
Last week the resource-management workgroup tried to tackle a design issue related to the device plugin. \
After much discussion we agreed that we wanted more opinions on the different approaches that we currently have. \
I've created a [document that captures the different approaches as well as the pros and cons](https://docs.google.com/document/d/1xfmakZZ_0Pq6OpLTXZTqOD20tj_hQuQAXUuVujxL1Rw/edit?usp=sharing)
    *   We should have define a separate per container operation in device plugin.
    *   https://github.com/kubernetes/kubernetes/pull/58172 -- yuju to review
*   Logging improvements
    *   Derek Carr, Peter Portante
    *   https://github.com/kubernetes/kubernetes/issues/58638
    *   Namespace name can duplicate overtime, we may need namespace UUID in container log path.
*   CRI container log stats
    *   Slides: https://docs.google.com/presentation/d/1BKbTa7RBVMTjZlk_6CV5SZfV4fen3bzk5PoqiXYIUK4/edit#slide=id.p
    *   Issue: https://github.com/kubernetes/kubernetes/issues/55905
*   Container log rotation
    *   yujuhong@
    *   https://docs.google.com/document/d/1oQe8dFiLln7cGyrRdholMsgogliOtpAzq6-K3068Ncg/edit?usp=sharing
*   Node auto repair repository discussion
    *   Derek Carr
    *   desire is to have an open source node remedy system that watches node conditions and/or taints to coordinate remedy response across multiple clouds (i.e. for example, reboot node)


# Jan 16 {#jan-16}



*   Review Pending v1.Pod API Changes [@verb]
    *   Debug Containers: [Command in PodStatus](https://github.com/verb/community/blob/ff62a6a15f094f00acf5e5dfe972b73197f04042/contributors/design-proposals/node/troubleshoot-running-pods.md#debug-container-status)
        *   Requires keeping track of v1.Container in kubelet
    *   Shared PID: [ShareProcessNamespace in PodSpec](https://github.com/verb/community/blob/080397194824d19d72b8c50f92667e93e7a244af/contributors/design-proposals/node/pod-pid-namespace.md#kubernetes-api-changes)
*   Windows container roadmap and Windows configuration in CRI (Patrick.Lang, jiangtli @microsoft.com)
    *   https://trello.com/b/rjTqrwjl/windows-k8s-roadmap
    *   Windows container configuration in CRI:
        *   https://github.com/kubernetes/community/pull/1510
        *   Questions:
            *   Question: Is there pod level sandbox isolation mechanism on Windows?
            *   Answer: Currently only kernel level isolation for per-container, no hierarchy isolation like cgroups. They will work with containerd and cri-containerd to add the hypervisor level isolation at pod level.
            *   Dawn: Can we add e2e test for the change we make?
            *   Dawn & Yuju: Can we have an overview of how resource management will work on windows?
            *   Dawn: Let's start with CPU and memory, from your current proposal, currently there is no storage resource define there now.
*   Review request: Node-level Checkpointing Manager (Vikas, vichoudh@redhat.com)
    *   https://github.com/kubernetes/kubernetes/pull/56040
    *   Just requires approval
    *   It's good to have a common library instead of having separate implementations in different part. However, we should make it clear that we prefer not to add checkpoint.
    *   Derek: Maybe we can have a document to track all the checkpoint done by kubelet, and whenever people add new checkpoint they need to update the document.
*   Short discussion about the node-fencing progress (bronhaim@, redhat)
    *   https://github.com/kubernetes/community/pull/1416
    *   https://github.com/rootfs/node-fencing


# Jan 9 {#jan-9}



*   crictl v1.0.0-alpha.0 demo. @lantaol
    *   Slides: https://docs.google.com/presentation/d/1jnInpUccKCRxCuXWC141gjipRs0Y5_0XSELDpUEXz5E/edit?usp=sharing
*   [kube-spawn](https://github.com/kinvolk/kube-spawn/) local multi-node cluster tool
    *   Using kubeadm and systemd-nspawn
    *   Useful for: testing CRI or other Kubernetes patches
    *   Demo: https://asciinema.org/a/152314
*   deprecating rktnetes in k8s v1.10 (tracking issue: https://github.com/kubernetes/kubernetes/issues/53601)
*   sig-node Q1, 2018 plan
    *   https://docs.google.com/document/d/15F3nWPPG3keP0pzxgucPjA7UBj3C31VsFElO7KkDU04/edit
    *   much of the work carried from last quarter
    *   draft includes priority and possible owners
    *   please review and suggest changes before sharing with SIG PM
*


# Jan 2 {#jan-2}

Cancelled.
