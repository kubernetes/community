# SIG Node Meeting Notes

# Future

*   [Virtual Kubelet (@rbitia)](https://docs.google.com/document/d/1MAn_HMZScni89hDwI4nQMk_SWTx9oi16PuOREhpVJJI/edit?usp=sharing)
*   [Jess Frazelle / Kent Rancourt]: Proposal for kubelet feature to freeze pods placed in a hypothetical “frozen” state by the replication controller in response to a scale-down event. Enable pods to be thawed by socket activation.
*   `regular resource usage tracking resource tracking for 100 pods per node` => this e2e test has been failing (flaking?) quite consistently in release-blocking dashboards. Should we block the 1.14 release on this? If not, could you help resolve it? [Issue:75039](https://github.com/kubernetes/kubernetes/issues/75039) (@mariantalla)
*   Issue to discuss: [Hardware topology awareness at node level (including NUMA)](https://github.com/kubernetes/kubernetes/issues/49964)
*   OCI Hooks PreStart, PostStop - @alban
    *   Suggestion to add Kubernetes labels in the OCI State ([slide 15](https://docs.google.com/presentation/d/1i8csKAf15j3ZDeHxuUlHBDvHiBI44_UxATdaaAy-pjE/edit#slide=id.g5da9f7933e_0_51))
    *   Use case: [Inspektor Gadget](https://github.com/kinvolk/inspektor-gadget)
    *   Previous presentation at the OCI meeting, July 24th: [notes](https://hackmd.io/El8Dd2xrTlCaCG59ns5cwg?view) ; [slides](https://docs.google.com/presentation/d/1i8csKAf15j3ZDeHxuUlHBDvHiBI44_UxATdaaAy-pjE/edit)
    *   @RenaudWasTaken: [Improving the Device Plugin Mechanisms](https://docs.google.com/document/d/1wPlJL8DsVpHnbVbTaad35ILB-jqoMLkGFLnQpWWNduc/edit?usp=sharing)
*   Hugepages in 1.18([notes](https://docs.google.com/document/d/1Xnb3ioEH_JmWLdFJ-i6Lp7jFc_fDuzzAa2oEXJglBbM/edit#))


# December 24 & 31, 2019

*   Happy Holidays! No meeting 🎉🎉


# December 17, 2019



*   [klueska] Looking for review on **<code>CPUManager</code></strong> change:
    *   Blocking progress of moving <strong><code>TopologyManager</code></strong> into beta
    *   Overall issue: https://github.com/kubernetes/kubernetes/issues/83476
    *   Specific PR: https://github.com/kubernetes/kubernetes/pull/84462
*   [bart0sh] Request for review: multi-size hugepages: [#82820](https://github.com/kubernetes/kubernetes/pull/82820) and [#84051](https://github.com/kubernetes/kubernetes/pull/84051)
    *   Hugepages KEP update is already merged: [#1271](https://github.com/kubernetes/enhancements/pull/1271)
*   [michaelgugino] Discuss/request for review ‘Node Maintenance Lease’ KEP: https://github.com/kubernetes/enhancements/pull/1411
*   Issue: https://github.com/kubernetes/kubernetes/issues/85966


# December 10, 2019



*   [kkmsft,patricklang] Review alternate proposal to #[84486](https://github.com/kubernetes/kubernetes/pull/84486) discussed earlier - annotations in ImageManager with certain fields copied from PodSandboxConfig (runtimehandler). [Link ](https://docs.google.com/document/d/1S30VFhwdKyeuwFuGpCiN4XIIQYkScPqZZ5n_gfwWVy8/edit)
    *   ask: can Lantao review since this was based on in-person discussion at Kubecon?
    *   Will update existing KEP and change API sections to follow this approach.
*   [kad, RenaudWasTaken] WG-Resource-Management future
    *   kad and RenaudWasTaken to come up with written proposal/plan/goals for WG
*   [wojtek-t] Immutable Secrets/ConfigMaps ([KEP](https://github.com/kubernetes/enhancements/pull/1369))
*   [RainbowMango] Request for review and a proposal:
    *   https://github.com/kubernetes/kubernetes/pull/85446
    *   A suggestion for fast forwarding metrics stuff of kubelet.
*   [vpickard] Questions about doing E2E tests on NUMA nodes	 
*   ~~[klueska] Best method to migrate checkpointed state across kubelet upgrade~~
    *   ~~Needed to remove for this week since I can’t make it to the meeting anymore~~
    *   ~~Please see [this](https://github.com/kubernetes/kubernetes/pull/84462#issuecomment-550018667) and comment when you have time~~
*   [jaypipes] Currently working on repro'ing this bug: https://github.com/kubernetes/kubernetes/issues/79159
    *   I can reproduce the bug successfully but am a bit confused as to why I never see Allocatable.cpu for the node never get decremented. Is this number not intended to be decremented ever? i.e. does the scheduler keep the current number of static cores allocatable to each node in memory?


# December 03, 2019



*   [vinaykul, quinton-hoole] [In-Place Vertical Scaling KEP](https://github.com/kubernetes/enhancements/pull/686/commits/533c3c625a49b07fcbbf9449080f1a7b139c92f3) updates
    *   K8s contributor summit [Live API review](https://docs.google.com/presentation/d/16-R9gEpKF6nyzDNtA5JDsKzQzr_vF_ZevN0ezZVYHpk/edit#slide=id.g78e88100d6_0_23) session with @liggitt, @thockin
    *   Next steps
*   [mrunal/giuseppe] Cgroup V2 updates
    *   KEP - https://github.com/kubernetes/enhancements/pull/1370
    *   How does this relate to PR https://github.com/containerd/containerd/pull/3799 ?
*   [bg.chun] CRI Update(for hugepages) review ([slides](https://docs.google.com/presentation/d/1H5xhCjgCZJjdTm-mGUFwqoz8XmoVd8xTCwJ1AfVR_54/edit?usp=sharinghttps://docs.google.com/presentation/d/1H5xhCjgCZJjdTm-mGUFwqoz8XmoVd8xTCwJ1AfVR_54/edit?usp=sharing))
    *   [CRI Update PR](https://github.com/kubernetes/kubernetes/pull/83614), [Hugepages KEP PR(merged)](https://github.com/kubernetes/enhancements/pull/1199), Hugepages in 1.18([notes](https://docs.google.com/document/d/1Xnb3ioEH_JmWLdFJ-i6Lp7jFc_fDuzzAa2oEXJglBbM/edit#))
*   [patricklang] Can we get a review on a test simplification for some SIG-Node owned tests? https://github.com/kubernetes/kubernetes/pull/84788 - Yu-ju reviewed ✔️ Thanks!
*   [bart0sh] Request for review: multi-size hugepages: [#82820](https://github.com/kubernetes/kubernetes/pull/82820) and [#84051](https://github.com/kubernetes/kubernetes/pull/84051)

    Hugepages KEP update is already merged: [#1271](https://github.com/kubernetes/enhancements/pull/1271)



# November 12, 2019



*   [verb] Promoting shareProcessNamespace to GA ([#84356](https://pr.k8s.io/84356))
    *   ready for 1.17
*   [verb] Updates on pod troubleshooting and kubectl debug ([KEP](https://features.k8s.io/1204))
    *   API updated in 1.16
    *   Worked out plan for kubectl debug with sig-cli
        *   (forgot to mention) the current plan is to have functionality to allow node-level debugging, via `kubectl debug node` launching a privileged container in the host namespace
    *   CRI update pending to allow process namespace targeting
    *   Enhancement [#277](https://features.k8s.io/277) has rough roadmap
*   [bg.chun] asking CRI update for hugepages in 1.17 Rel(slides/TBD)
    *   [CRI Update PR](https://github.com/kubernetes/kubernetes/pull/83614), [Hugepages KEP PR(merged)](https://github.com/kubernetes/enhancements/pull/1199)
*   [vinaykul, quinton-hoole] [In-Place Vertical Scaling KEP](https://github.com/kubernetes/enhancements/pull/686/commits/533c3c625a49b07fcbbf9449080f1a7b139c92f3)
    *   Discuss [mini-KEP for Container resources CRI extension/update](https://github.com/kubernetes/enhancements/pull/1342) changes
        *   Remove GetContainerResources, extend ContainerStatus API.
        *   ContainerResources struct will be oneof
    *   Status: Awaiting API review
*   [patricklang, kkmsft, tallclair] - 1.17 action needed - closing on RuntimeHandler updates. What would it take to get to /approve from SIG-Node?
    *   Diagram added to [Discussion Slides](https://docs.google.com/presentation/d/11x6giWCYk1JGt3QFrcbS-m_Q9mcxIz_T0-RGhhSndf4/edit?usp=sharing)
    *   https://github.com/kubernetes/kubernetes/pull/84486
    *   Scope:
        *   Use runtime class to choose hyperV / process isolation;
            *   add PodSandboxConfig if needed to EnsureImage / GetImageRef
        *   Support OS version;
        *   Multiple snapshots will be another KEP.
*   [patricklang] - 1.17 action needed - approver for kubelet Windows version labelling: https://github.com/kubernetes/kubernetes/pull/84472 . Got API feedback/approve from Jordan, will need someone for kubelet
*   
*   [bart0sh] 1.17 action needed: review of pending Hugepages PRs: https://github.com/kubernetes/kubernetes/issues/84526

    https://docs.google.com/document/d/1Xnb3ioEH_JmWLdFJ-i6Lp7jFc_fDuzzAa2oEXJglBbM/edit#



# November 5, 2019



*   [klueska,RenaudWasTaken] [Improving the Device Plugin Mechanisms](https://docs.google.com/document/d/1wPlJL8DsVpHnbVbTaad35ILB-jqoMLkGFLnQpWWNduc/edit?usp=sharing)
*   [kad,klihub,marquiz,bart0sh] [Container Device Interface](https://docs.google.com/document/d/1Tc0Kc4GDWx1gFvGQbBUizudSuND6Kq8GiH7KVm_X5eg/edit#heading=h.5lakm98lya8j)
*   ~~[jirving] proposed changes/rewording of sidecars kep https://github.com/kubernetes/enhancements/pull/1344 - reviewed thanks!~~
*   [patricklang,kkmsft,tallclair] Need decision on adding RuntimeHandler to PodSandboxConfig vs individual methods (Pull/List/RemoveImage)
    *   [Discussion Slides](https://docs.google.com/presentation/d/11x6giWCYk1JGt3QFrcbS-m_Q9mcxIz_T0-RGhhSndf4/edit?usp=sharing) , [KEP](https://github.com/kubernetes/enhancements/blob/master/keps/sig-windows/windows-runtimeclass-support.md)
    *   https://github.com/kubernetes/kubernetes/pull/84486
    *   Q: should runtimeHandler move into PodSandboxConfig, or stay separate?
    *   Also have [PR#84472](https://github.com/kubernetes/kubernetes/pull/84472) ready for review with node label `node.kubernetes.io/windows-build. `Who can review?
*   [vinaykul, quinton-hoole] [In-Place Vertical Scaling KEP](https://github.com/kubernetes/enhancements/pull/686/commits/533c3c625a49b07fcbbf9449080f1a7b139c92f3)
    *   Discuss [mini-KEP for Container resources CRI extension/update](https://github.com/kubernetes/enhancements/pull/1342)
        *   Q: Remove GetContainerResources, extend ContainerStatus API?
        *   ContainerResources struct will be oneof
    *   Awaiting scheduling of API review
*   [klueska] TopologyManager status / reviews
    *   [At risk of not making it for 1.17](https://github.com/kubernetes/kubernetes/issues/83479)
    *   https://github.com/kubernetes/kubernetes/pull/84525
*   


# October 29, 2019



*   [bg.chun] Asking for a review for pending PRs for hugepage enhancements
    *   [Hugepages PRs](https://docs.google.com/document/d/1Xnb3ioEH_JmWLdFJ-i6Lp7jFc_fDuzzAa2oEXJglBbM/edit#)
*   [tallclair,draveness] Node-local API cache [#84248](https://github.com/kubernetes/kubernetes/issues/84248)
    *   notes: https://github.com/kubernetes/kubernetes/issues/84248#issuecomment-547539826
*   [vinaykul, quinton-hoole] [In-Place Vertical Scaling KEP](https://github.com/kubernetes/enhancements/pull/686/commits/533c3c625a49b07fcbbf9449080f1a7b139c92f3)
    *   Kick-off review of [mini-KEP for Container resources CRI extension/update](https://github.com/kubernetes/enhancements/pull/1342)
    *   Quick [prototype code](https://github.com/vinaykul/kubernetes/commit/4f2c9f08d0797b208ea137f6734de5c58230995a) for above mini-KEP.
    *   Schedule API review (opened PR for moving parent KEP to implementable)
*   [patricklang,kkmsft,tallclair] Review approach to runtimeHandler for CRI
    *   https://github.com/kubernetes/kubernetes/pull/84486
    *   Q: should runtimeHandler move into PodSandboxConfig, or stay separate?
*   [jirving] Request review of Sidecars PR https://github.com/kubernetes/kubernetes/pull/80744
    *   Should sidecars/non-sidecars share the TerminationGracePeriod or should they each get TerminationGracePeriod to terminate in?
    *   Should sidecars be included in Pod phase calculation? i.e should a sidecar that exited badly make the pod ‘Failed’ ?
    *   API has been reviewed by @thockin, some small changes were suggested https://github.com/kubernetes/kubernetes/pull/79649
*   @Kevin Klues, @RenaudWasTaken: [Improving the Device Plugin Mechanisms](https://docs.google.com/document/d/1wPlJL8DsVpHnbVbTaad35ILB-jqoMLkGFLnQpWWNduc/edit?usp=sharing)
*   [kad,klihub,marquiz,bart0sh] [Container Device Interface](https://docs.google.com/document/d/1Tc0Kc4GDWx1gFvGQbBUizudSuND6Kq8GiH7KVm_X5eg/edit#heading=h.5lakm98lya8j)


# October 22, 2019



*   [vinaykul] [In-Place Vertical Scaling KEP](https://github.com/kubernetes/enhancements/pull/686/commits/533c3c625a49b07fcbbf9449080f1a7b139c92f3)
    *   Review first-stab at [proposed CRI change](https://github.com/kubernetes/enhancements/issues/1287#issuecomment-544809736) to support Windows containers.
    *   API review prep [document](https://docs.google.com/presentation/d/16-R9gEpKF6nyzDNtA5JDsKzQzr_vF_ZevN0ezZVYHpk/edit?usp=sharing) - any comments?
    *   Review open issues https://github.com/kubernetes/enhancements/issues/1287
    *   Help schedule API review with @thockin
*   TODO for next week: discuss hooks/device plugins improvements. Nvidia+Intel [doc](https://docs.google.com/document/d/1wPlJL8DsVpHnbVbTaad35ILB-jqoMLkGFLnQpWWNduc/edit#)


# October 15, 2019



*   Enhancements (KEP) freeze today
    *   https://github.com/kubernetes/enhancements/pull/1148
    *   https://github.com/kubernetes/enhancements/pull/1243
*   [patricklang, tallclair] RuntimeClass for multi-arch and os versions
    *   enhancement [#1301](https://github.com/kubernetes/enhancements/issues/1301) filed for 1.17, KEP [#1302](https://github.com/kubernetes/enhancements/pull/1302) opened based off last week’s discussion
*   [pjbgf] Seccomp: new built-in profiles, complain mode and ConfigMaps
    *   https://github.com/kubernetes/enhancements/pull/1269
    *   https://github.com/kubernetes/enhancements/pull/1257
*   [vinaykul] [In-Place Vertical Scaling KEP](https://github.com/kubernetes/enhancements/pull/686/commits/533c3c625a49b07fcbbf9449080f1a7b139c92f3) - status update & next steps
    *   Tracking to v1.18 milestone now
    *   API review prep [document](https://docs.google.com/presentation/d/16-R9gEpKF6nyzDNtA5JDsKzQzr_vF_ZevN0ezZVYHpk/edit?usp=sharing)
    *   Outstanding issues https://github.com/kubernetes/enhancements/issues/1287
    *   Help scheduling API review with @thockin


# October 8



*   [mikebrow, kad, klueska, Lantao] Support OCI hooks in Kubernetes (continue ...)
    *   Renaud Gaubert: https://github.com/containerd/cri/pull/1248#issuecomment-537366588: Can we move the device support down to container runtime, and define runtime device plugin? In this way:
        *   The device can be supported in the container runtime with lower level information, and can probably avoid arbitrary OCI hooks, which is hacky;
        *   The same logic can be shared by Kubernetes, Docker and podman etc. instead of having different implementations in each.
        *   Containerd knows better how to support different runtimes, e.g. runc, kata, gvisor etc.
    *   Jiaying:
        *   The idea sounds reasonable to GPU, we may want to think about other plugins and general Kubernetes use cases.
    *   **TODO: Let’s have a written doc about **
        *   https://github.com/containerd/cri/pull/1248#issuecomment-537366588. (Renaud Gaubert)
        *   Participants in document proposal: Renaud, Kevin, Alexander, Ed, Zhipeng (Howard) Huang
*   [klueska] Review TopologyManager extension mentioned by @derekwaynecarr
    *   https://github.com/kubernetes/enhancements/pull/1121
*   [patricklang, tallclair] RuntimeClass for multi-arch and os versions
    *   Approaches to multi-version support
        *   Please review this [doc outlining the problems & approach](https://docs.google.com/document/d/12uZt-KSG8v4CSyUDr0EC6btmzpVOZAWzqYDif3EoeBU/edit?usp=sharing)  first - it outlines the directions we need to decide along with SIG-Node and SIG-Architecture
        *   Once we agree on the path forward, it will be moved out into KEPs such as https://github.com/kubernetes/enhancements/pull/1254
*   [vinaykul] [In-Place Vertical Scaling KEP](https://github.com/kubernetes/enhancements/pull/686/commits/533c3c625a49b07fcbbf9449080f1a7b139c92f3) - status update.
    *   Approved and merged as provisional.
    *   Outstanding issues https://github.com/kubernetes/enhancements/issues/1287
    *   Identify a few time-slots to schedule API review with @thockin
*   [dims] PR for [Moving external facing kubelet APIs to staging](https://github.com/kubernetes/kubernetes/pull/83551) is ready. Need reviews!


# October 1



*   [mikebrow, lantaol] Support OCI hooks in Kubernetes:
    *   Containerd issue: https://github.com/containerd/cri/issues/405
    *   WIP Containerd PR: https://github.com/containerd/cri/pull/1248, basically the current idea is to model OCI hooks similar with seccomp/apparmor profile:
        *   The OCI hook details are defined on the node in hook config files (profiles). The reason is that OCI hooks run in the [host (runtime) namespace](https://github.com/opencontainers/runtime-spec/blob/master/config.md#posix-platform-hooks), so it is very specific to the node environment, e.g. it may directly invoke binaries in the host filesystem.
        *   Each group of OCI hooks is referenced by a name (hook name, similar to seccomp profile name or runtime handler);
        *   What hooks to run is from CRI, and Device Plugin will be the only user right now.
    *   cri-o approach: 
        *   oci-hooks are defined on the node via set of [json files](https://www.mankier.com/5/oci-hooks). 
        *   Hooks are executed when the condition is met. 
        *   Container annotations that are passed down from CRI to OCI can be used as a condition. Device Plugins can set those annotations. 
    *   Use cases:
        *   https://speakerdeck.com/kad/extending-kubernetes-with-intel-r-accelerator-devices
        *   https://docs.google.com/document/d/1VNIDKgRnn1dL9hnNdaaTdiJFrEGDAhROVJAHMOyingI/edit?usp=sharing
            *   Device plugin:
                *   It doesn’t handle container restart
                *   Not enough information
                *   Cleanup the device
        *   [Inspektor Gadget](https://github.com/kinvolk/inspektor-gadget) use OCI hooks and would like OCI annotations with pod uid, container id and pod labels (see [slide 15](https://docs.google.com/presentation/d/1i8csKAf15j3ZDeHxuUlHBDvHiBI44_UxATdaaAy-pjE/edit#slide=id.g5da9f7933e_0_51) and [notes](https://hackmd.io/El8Dd2xrTlCaCG59ns5cwg?view) from OCI meeting) [@alban]
        *   Main feedback we get about the device plugin interface:
            *   Not full container lifecycle is under control, e.g. cleanup the device after container exits; response: (if runc’s hook types are insufficient we should bring that to the runtime spec team as an issue) 

                We should be able to enable all hook types at the CRI runtime layer.

            *   Lack of information. Most of the information can be retrieved by Kubelet, but not passed into device plugin yet.
                *   Previous PR discussing using annotations: https://github.com/kubernetes/kubernetes/pull/61775
*   [mvladev] [Master Service of type ExternalName KEP](https://github.com/kubernetes/enhancements/pull/1216) feedback
    *   kubelet [PR](https://github.com/kubernetes/kubernetes/pull/79312)
*   [yjhong] Beta CoreOS image in node e2e ([issue](https://github.com/kubernetes/test-infra/issues/14465), [PR](https://github.com/kubernetes/test-infra/pull/14466)); current [image config](https://github.com/kubernetes/test-infra/blob/master/jobs/e2e_node/image-config.yaml).
*   [lantaol, patricklang] Plan and status update of Kubernetes containerd Windows support alpha.
    *   https://github.com/kubernetes/enhancements/blob/master/keps/sig-windows/20190424-windows-cri-containerd.md#alpha-release-proposed-117
    *   Draft updates to handle multi os version & runtimeclass
        *   https://github.com/kubernetes/enhancements/pull/1254 
    *   Project board: https://github.com/orgs/kubernetes/projects/34
    *   Recent meeting notes: https://docs.google.com/document/d/1PKlmTIh0-qSzGDPhYmc_5BJjdD9GRdMGIATzvlhUAwM/edit?usp=sharing
    *   Containerd windows CRI validation test: https://testgrid.k8s.io/sig-node-containerd#cri-validation-windows
    *   Containerd windows Kubernetes e2e test: https://testgrid.k8s.io/sig-windows#containerd-l2bridge-windows-master (Still using forked containerd right now, will setup one with upstream containerd soon)
*   [vinaykul] [In-Place Vertical Scaling KEP](https://github.com/kubernetes/enhancements/pull/686/commits/533c3c625a49b07fcbbf9449080f1a7b139c92f3) - just a status update.
    *   Approval from sig-node and sig-scheduling.
    *   Awaiting @kgolab lgtm and @mwielgus approve, API review and next steps.


# September 24



*   [vinaykul] [In-Place Vertical Scaling KEP](https://github.com/kubernetes/enhancements/pull/686/commits/533c3c625a49b07fcbbf9449080f1a7b139c92f3) review follow-up and status update.
    *   Identify approver from sig-node to list in the KEP approvers list.
*   [bg.chun] update hugepages KEP([slide](https://docs.google.com/presentation/d/1H5xhCjgCZJjdTm-mGUFwqoz8XmoVd8xTCwJ1AfVR_54/edit?usp=sharing))
    *   [Update huge pages KEP for NUMA support of huge pages](https://github.com/kubernetes/enhancements/pull/1245)
    *   [Update huge pages KEP for container isolation of huge pages](https://github.com/kubernetes/enhancements/pull/1199)
*   [harche/lumjjb] Update on Image Decryption KEP
    *   https://github.com/kubernetes/enhancements/pull/1066 
*   [klueska] Discuss TopologyManager enhancements for 1.17 / beta
    *   https://docs.google.com/document/d/1YTUvezTLGmVtEF3KyLOX3FzfSDw9w0zB-s2Tj6PNzLA/edit#


# September 17



*   [vinaykul] Review follow-up for [In-Place Vertical Scaling KEP](https://github.com/kubernetes/enhancements/pull/686/commits/533c3c625a49b07fcbbf9449080f1a7b139c92f3). Discuss next steps.

Updates:



    *   Reviewed with sig-scheduling last Thu, Abdullah and Bobby are fine with it.
    *   Updated sig-autoscaling today, requested approval - they will look this week.
    *   No word from sig-arch yet for API review request.


# September 10



*   [vinaykul] Review updates to [In-Place Vertical Scaling KEP](https://github.com/kubernetes/enhancements/pull/686/commits/533c3c625a49b07fcbbf9449080f1a7b139c92f3) from previous week.

Updates:



    *   Restate Kubelet restart resize handling to minimum guarantees.
    *   Clarify scheduler role - requested sig-scheduling review
*   [bg.chun] Discuss [update existing hugepages kep](https://github.com/kubernetes/enhancements/pull/1199) \
key points
    *   Update CRI to support container isolation of huge pages
        *   Issue: [Support Container Isolation of Hugepages](https://github.com/kubernetes/kubernetes/issues/80716)
        *   PR: TBD
    *   Update cAdvisor to discover hugepages per NUMA node
        *   cAdvisor PR: [Add hugepage info to v1 node structure](https://github.com/google/cadvisor/pull/2304)
*   [danbeaulieu] https://github.com/kubernetes/kubernetes/issues/82440


# September 03



*   [vinaykul] Review updates to [In-Place Vertical Scaling KEP](https://github.com/kubernetes/enhancements/pull/686/commits/533c3c625a49b07fcbbf9449080f1a7b139c92f3) from previous week.

Key changes:



    *   Call out details of Kubelet-APIServer interaction in handling Pod resize.
        *   Scheduling a API review with SIG API.
    *   Move eviction of lower priority Pods to Future work section.
    *   Remove Pod departure as a trigger for resize retries.
    *   Call out details of Static CPU manager policy resize handling.
*   [harche/lumjjb] Update on Image Decryption KEP
    *    https://github.com/kubernetes/enhancements/pull/1066 
*   [Seth] Discuss https://github.com/kubernetes/kubernetes/pull/81858
    *   about “extended resources via kubelet flag?" 


# August 27th



*   [vinaykul] Kick-off review of u[pdated Vertical Scaling KEP](https://github.com/kubernetes/enhancements/pull/686#). The design has been updated as per consensus achieved from discussions over the past few weeks.

Key changes in [this commit](https://github.com/kubernetes/enhancements/pull/686/commits/bc9dc2beb57d6d3a3c39fd2987f7714658562dde):



    *   PodSpec holds ResourcesAllocated for Pod’s Containers on Node
    *   ‘resourceallocation’ subresource for Kubelet to set/update ResourcesAllocated
    *   Kubelet restart fault tolerance section from [KEP discussion](https://github.com/kubernetes/enhancements/pull/686#discussion_r311136176) 
    *   Remove Resizing PodCondition 
*   [csadeniyi] Request a status update on: Clear the NodeSelector when loading checkpointed Pods from disk: https://github.com/kubernetes/kubernetes/pull/72202


# August 20th



*   [vinaykul] Continue Vertical Scaling KEP review
    *   Review Kubelet restart handling flow for node-local ResourcesAllocated source-of-truth:  https://github.com/kubernetes/enhancements/pull/686#discussion_r308785033
    *   https://github.com/kubernetes/enhancements/pull/686#discussion_r311136176
        *   Decide whether to store accepted resize [node-locally or in PodSpec](https://github.com/kubernetes/enhancements/pull/686#discussion_r313656371)
        *   Decide if we want to have Resize PodCondition  
    *   dashpole@’s slide deck: https://docs.google.com/presentation/d/1YRFrZT1ISBZa3ZSKeL2eyIQFFGVbtUSXM2Hnfj8SteI/edit#slide=id.p


# August 13th



*   [harche/lumjjb] Update on Image Decryption KEP
    *    https://github.com/kubernetes/enhancements/pull/1066 
    *   Next Step:
        *   Need approval from Derek and Dawn
        *   Lantao will be the reviewer from sig-node.
*   [vinaykul] Continue Vertical Scaling KEP review
    *   Review Kubelet restart handling flow for node-local ResourcesAllocated source-of-truth:  https://github.com/kubernetes/enhancements/pull/686#discussion_r308785033
    *   https://github.com/kubernetes/enhancements/pull/686#discussion_r311136176
*   [cadeniyi] Requesting review

    https://github.com/kubernetes/kubernetes/pull/72202


    AI: find the reviewer: yjhong?

*   [klueska] Limiting the supported socket count when TopologyManager enabled
    *   HintProviders (in the worst case) need to enumerate all possible socket masks and generate hints for them \
	→ results in 2^n -1 masks, where n is the number of sockets \
	→ state explosion for large number of sockets
    *   What is the best way to safeguard against this?
        *   Fail the kubelet if the machine has more than (e.g.) 8 sockets and TopologyManager is enabled?
        *   Wait until a pod with topology aligned resources tries to launch and fail the pod without attempting to enumerate the socket masks?


# August 6th



*   Should pod overheads be considered for eviction handling - see https://github.com/kubernetes/kubernetes/issues/80993
*   handling node shutdown - see KEP https://github.com/kubernetes/enhancements/pull/1116 
*   code organization for kubeadm validators - see https://github.com/kubernetes/kubeadm/issues/1638 (Added by yassine)
*   Requesting Reviews:
        *   Switch to k8s.io/utils/inotify - [PR #80689](https://github.com/kubernetes/kubernetes/pull/80689)  (Added by dims on behalf of sig-arch/code-org efforts)
    *   upgrade vishvananda/netlink to v1.0.0 - [PR #80290](https://github.com/kubernetes/kubernetes/pull/80290) (Added by dims on behalf of sig-arch/code-org efforts)
    *   Add ImageFSInfo, ContainerStats, and ListContainerStats impl for linux to dockershim - [PR #80105](https://github.com/kubernetes/kubernetes/pull/80105) (Added by dims to help with the dockerd CRI split effort)
*   [vinaykul] Continue Vertical Scaling KEP review
    *   Review Kubelet restart handling flow for node-local ResourcesAllocated source-of-truth:  https://github.com/kubernetes/enhancements/pull/686#discussion_r308785033
    *   https://github.com/kubernetes/enhancements/pull/686#discussion_r311136176
    *   


# July 30



*   [jaypipes] Request a status update on [node-level userspace remapping work](https://github.com/kubernetes/enhancements/issues/127) from either @derekwaynecarr or @vikaschoudhary16
    *   Is this stuck waiting on reviews? Stuck on implementation or design debate? Is there anything that EKS team members can contribute to speed this along?
    *   There are resource constraints on RH side, plus it's de-prioritized. Jay will reach out to the customer who is asking about this feature and determine if there is an alternate solution to the problem and whether the user ID remapping solution would be solving the problem fully in a unique way for the customer. EKS team may contribute resources to reignite Vikas' original PR if customer feels user ID remapping is the most viable solution (and EKS team feels the proposed implementation is viable as well)
*   [derekwaynecarr] Need to fix up our mailing list settings
    *   https://github.com/kubernetes/community/issues/3927
    *   Dawn - Fixed the setting.
*   [derekwaynecarr] 1.16 Feature Freeze items
*   [klueska] Blocking TopologyManager PRs:

    (All other PRs currently depend on these and require a rebase once they are merged)


    (Once these are merged, all other changes should be isolated to cm)

    *   https://github.com/kubernetes/kubernetes/pull/74357 \
(approved in cm, needs approval from OWNER in cmd/kubelet, pkg/kubelet)
    *   https://github.com/kubernetes/kubernetes/pull/74423 \
(needs approval from OWNER in pkg/kubelet)
*   [vinaykul] Vertical Scaling KEP
    *   Review Kubelet restart handling flow for node-local ResourcesAllocated source-of-truth:  https://github.com/kubernetes/enhancements/pull/686#discussion_r308785033
*   [verb] [#80645](https://pr.k8s.io/80645): Assigning sig-node as owner of pods integration tests
    *   wants approval from Dawn & Derek 
*   [tallclair] seccomp to GA
    *   https://github.com/kubernetes/enhancements/pull/1148


# July 23



*   [vinaykul] Continue review of updated Vertical Scaling KEP. https://github.com/kubernetes/enhancements/pull/686
    *   Move ResourcesAllocated into PodSpec w/ subresource update by Kubelet only vs. Node-local storage of current container requests/limits info to discover it in case of Kubelet restart.
    *   Keep ResizingPod condition to enable VPA to quickly know if it should evict in case of resize failure due to capacity. This avoids cluster resource waste.
*   [harche/Brandon] image decryption kep discussion/introduction
    *    https://github.com/kubernetes/enhancements/pull/1066  
    *   Kubecon CN Demo: https://www.youtube.com/watch?v=bzHPnlSfM_8
    *   Note: The feature author wants to target Kubernetes 1.16, Dawn thinks that 1.17 might be more reasonable timeline. Feature authors are okay with 1.17 target!
*   [dawnchen] https://github.com/kubernetes/kubernetes/issues/80203 
    *   Agreed on https://github.com/kubernetes/kubernetes/issues/80203#issuecomment-514333789
*   [tallclair] seccomp to GA
    *   https://github.com/kubernetes/enhancements/pull/1148
*   [sashayakovtseva] https://github.com/kubernetes/kubernetes/issues/79803 
*   [klueska] Outstanding TopologyManager PRs
    *   ~~Not a part of the enhancements tracking sheet for 1.16 yet (should we add it?) \
https://docs.google.com/spreadsheets/d/1txj0SuCJGm_uDCcJJDx_IjrfhlSnW-thtfQY91oXRLo/edit#gid=0~~
    *   https://github.com/kubernetes/kubernetes/pull/74357 \
(approved in cm, needs OWNER in kubelet to approve)
    *   ~~https://github.com/kubernetes/kubernetes/pull/80294 \
(follow up / cleanup PR for [74357](https://github.com/kubernetes/kubernetes/pull/74357))~~

        ~~(needs approval from OWNER in cm)~~

    *   ~~https://github.com/kubernetes/kubernetes/pull/80301 \
(approved in cm, needs OWNER in kubelet to approve) ~~
    *   https://github.com/kubernetes/kubernetes/pull/74423 \
(needs approval from OWNER in kubelet)
    *   ~~https://github.com/kubernetes/kubernetes/pull/80315~~

        ~~(needs approval from OWNER in cm)~~

    *   https://github.com/kubernetes/kubernetes/pull/73920 \
(needs rebase onto master once [74357](https://github.com/kubernetes/kubernetes/pull/74357) merges)

        (needs approval from OWNER in cm)

    *   ~~https://github.com/kubernetes/kubernetes/pull/74345 \
(needs rebase onto master once [74357](https://github.com/kubernetes/kubernetes/pull/74357) merges)~~

        ~~(has known bugs → [follow-up PR to fix](https://github.com/klueska/kubernetes/tree/upstream-get-topology-hints-map) and change internal interface) \
(needs approval from OWNER in cm)~~

    *   https://github.com/klueska/kubernetes/commits/upstream-get-topology-hints-map

        (Not a PR yet (only top 6 commits will be in PR)) \
(waiting for all above to merge before creating it)


        (will need approval from OWNER in cm)

*   https://github.com/moshe010/kubernetes/commit/64b4d86e607ae5a532995976915f00a31beb5300

        (Not yet a PR, waiting for all above to merge before creating it)


        (will need approval from OWNER in cm)

*   Review request
    *   https://github.com/kubernetes/kubernetes/pull/80105 (added by dims)
    *   https://github.com/kubernetes/enhancements/pull/866 (added by dims)
    *   https://github.com/kubernetes/kubernetes/pull/75229 (added by andrewsykim)
*   [bg.chun] New Feedback for “strict” policy of Topology Manager(just asking a review)
    *   https://docs.google.com/document/d/1Ytz-Qzwz65MvXOODmz9SDouN6MbyoOiZ_spaz61cUIQ/edit \



# July 16



*   Discuss interest in solving [#79389](https://github.com/kubernetes/kubernetes/issues/79389) via fallbackImages KEP (icelynjennings)
    *   https://docs.google.com/document/d/1ho1SOZ5PzkcKeSpRYVYVOht4pG5758uP_34twOWVabA/edit
    *   Conclusion: Current Container Runtime implementation satisfied the original request, we are closing the request. Thanks for the discussion.
*   [vinaykul] Continue review of updated Vertical Scaling KEP. https://github.com/kubernetes/enhancements/pull/686
    *   Move ResourcesAllocated into PodSpec w/ subresource update by Kubelet only vs. Node-local storage of current container requests/limits info to discover it in case of Kubelet restart.
    *   Keep ResizingPod condition to enable VPA to quickly know if it should evict in case of resize failure due to capacity.
*   [Lmdaly] Topology Manager 1.16 Update and Review remaining PRs:
    *   https://github.com/kubernetes/kubernetes/issues/72828
    *   PRs to be approved (in merge order):
        *   https://github.com/kubernetes/kubernetes/pull/79343
        *   https://github.com/kubernetes/kubernetes/pull/73580
        *   https://github.com/kubernetes/kubernetes/pull/74357
        *   https://github.com/kubernetes/kubernetes/pull/73920
        *   https://github.com/kubernetes/kubernetes/pull/74423
        *   https://github.com/kubernetes/kubernetes/pull/74345
*   [klueska] Topology Manager Update Proposals
    *   https://docs.google.com/document/d/1g5Aqa0BncQGRedSJH0TJQWq3mw3VxpJ_ufO1qokJ1LE/edit
    *   https://github.com/kubernetes/enhancements/pull/1121
    *   https://github.com/kubernetes/enhancements/pull/1131
*   [klueska] Ping on review for proper support for init Containers in the CPUManager
    *   https://github.com/kubernetes/kubernetes/pull/78762
*   [mvladev] Environment variable support for Master service of type ExternalName PR:
    *   https://github.com/kubernetes/kubernetes/pull/79312 
        *   Reviewer: tallclaire, dawnchen
    *   issue: https://github.com/kubernetes/kubernetes/issues/79280 


# July 9



*   Discuss interest in solving [#79389](https://github.com/kubernetes/kubernetes/issues/79389) via fallbackImages KEP (icelynjennings)
    *   https://docs.google.com/document/d/1ho1SOZ5PzkcKeSpRYVYVOht4pG5758uP_34twOWVabA/edit
    *   Since Icelyn is not available today, we postponed this to next week.
*   Add startupProbe to health checks - design and implementation was changed from the initial proposal: adding initializationFailureThreshold to adding pod-startup liveness-probe. (matthyx)
    *   https://github.com/kubernetes/kubernetes/pull/77807#issuecomment-505705292
*   [vinaykul] Continue review of Vertical Scaling KEP updates. https://github.com/kubernetes/enhancements/pull/686
*   [hankang, ehashman] cadvisor metrics deprecation
    *   https://github.com/kubernetes/enhancements/blob/master/keps/sig-instrumentation/20181106-kubernetes-metrics-overhaul.md#cadvisor-instrumentation-changes  
    *   https://github.com/kubernetes/kubernetes/pull/69099 
*   image decryption kep discussion/introduction
    *   https://github.com/kubernetes/enhancements/pull/1066


# July 2

Meeting canceled


# June 25



*   [jirving] Sidecars kep implementable? https://github.com/kubernetes/enhancements/pull/1109
*   [vinaykul] Review/discuss updated Vertical Scaling KEP. https://github.com/kubernetes/enhancements/pull/686
*   [vinaykul] Per-Pod in,out network bandwidth requests & limits idea feasible? (This is different from bandwidth annotations today as scheduler would consider this as a resource during scheduling similar to cpu/memory)
*   [klueska] Proper support for init Containers in the CPUManager static policy

    https://github.com/kubernetes/kubernetes/pull/78762

*   [yastij, smarterclayton] pod checkpointing and pod-safety https://github.com/kubernetes/enhancements/issues/1103 
*   [yastij, smarterclayton] handling of pod startup for the node shutdown KEP
*   [mvladev] Environment variable support for Master service of type ExternalName

	https://github.com/kubernetes/kubernetes/pull/79312 



*   


# June 18



*   [btyler] Discuss Commit Class KEP: https://github.com/kubernetes/enhancements/pull/1078
*   [marquiz] New Git repository (under kubernetes-sigs org) for Node Feature Discovery Operator \
https://github.com/kubernetes-sigs/node-feature-discovery/issues/205
*   [vinaykul] Discuss Vertical Scaling updated design

	https://github.com/kubernetes/enhancements/pull/686#

	* Can we safely evict lower priority Pods from Kubelet??

	


# June 11

Cancelled.


# June 4

in



*   [patricklang] Fixing issue found with podResources enabled by default - [Windowskubelet won't start "Failed to create listener for podResources endpoint"](https://github.com/kubernetes/kubernetes/issues/78628#) 
    *   2 PRs open with slightly different approaches #[78670](https://github.com/kubernetes/kubernetes/pull/78670) / #[78671](https://github.com/kubernetes/kubernetes/pull/78671) - feedback? need to merge today and will need approver
    *   Device plugin isn’t tested on Windows either - @dashpole to look into how something can be disabled on a per-OS basis. Meeting consensus was to look into disabling device plugin manager that’s not used/tested on Windows
*   [derekwaynecarr] Kubelet in userns
    *   https://github.com/kubernetes/kubernetes/pull/78635
    *   Should we merge this PR or take a larger design?
*   PodOverhead
    *   API review completed, PR approved but not merged
    *   Moving to 1.16
*   RuntimeClass Scheduler
    *   API implementation approved, PR’s not in
    *   Moving to 1.16
*   TerminationPeriod: https://github.com/kubernetes/kubernetes/issues/77873
    *   Yuju: asked to understand use case more, may have other options to do it
    *   Derek - the Pod API already captures the user intent here, and it not being honored at shutdown could catch them by surprise
    *   


# May 28



*   [Extend NPD to support metrics](https://docs.google.com/document/d/1SeaUz6kBavI283Dq8GBpoEUDrHA2a795xtw0OvjM568/edit). (@xueweiz)
*   Issue Triage Process
    *   Document for discussion (requires membership in kubernetes-sig-node google group) https://docs.google.com/document/d/1JCN0lf9svebU2RTjCnz_cGnCrC4r5Sysb-ucdu1MLnA/edit#
*   Topology Manager PR Review and status update [@lmdaly]
    *   https://github.com/kubernetes/kubernetes/issues/72828


# May 21



*   Proposed to cancel meeting due to KubeCon EU


# May 14



*   Follow-up review to [In-Place Vertical Scaling KEP](https://github.com/kubernetes/enhancements/pull/686) after last week discussion.
    *   Identify any items that may need to be addressed to get KEP approved.
    *   @vinaykul, @derekwaynecarr @dashpole @dawnchen
*   Do we want to have dedicated meetings for issue triage?  (@derekwaynecarr)


# May 7



*   Pod Overhead: requests vs limits & QoS - see [this comment](https://github.com/kubernetes/enhancements/issues/688#issuecomment-488824132)
    *   shouldn't affect QoS
    *   BestEffort support should be a policy decision, independent of overhead implementation (e.g. fro Kata containers)
    *   Overhead should still be considered with BestEffort pods (scheduling, running)
    *   
*   Review latest updates to [In-Place Vertical Scaling KEP](https://github.com/kubernetes/enhancements/pull/686) @vinaykul, @DerekCarr @dashpole


# April 30



*   Discuss KEP https://github.com/kubernetes/enhancements/pull/808
    *   owner: khenidak
    *   prerequisite: SIG network owns the feature, should approve first. SIG Node would like to collaborate on the required CRI changes.
*   Graduate 3rd party device monitoring plugins to beta
    *   owner: @renaudwastaken and @dashpole
    *   Feature introduced in 1.13
    *   Will present a demo of the gpu-exporter
    *   NVIDIA GPU Exporter: https://github.com/NVIDIA/gpu-monitoring-tools/tree/master/exporters/prometheus-dcgm/k8s/pod-gpu-metrics-exporter
    *   Enhancement: https://github.com/kubernetes/enhancements/issues/606
    *   E2e testing: https://github.com/kubernetes/kubernetes/blob/master/test/e2e_node/device_plugin.go#L99
    *   Agreed to promote to beta in 1.15
*   [Yanir Quinn] : Discussion and a presentation (+quick demo) of node eviction controller (server side drain)  - https://github.com/kubevirt/node-maintenance-operator \
project started under KubeVirt. \
https://groups.google.com/d/msg/kubernetes-sig-node/6AFFLjVz8us/g7yDIFVEBwAJ
*   (FYI) cAdvisor APIs deprecation - https://github.com/kubernetes/kubernetes/issues/68522
    *   Suggest to separate cAdvisor endpoints from /summary deprecation. One is formally supported by K8s before, others are never properly supported by us.
*   (short but time critical) - Do we need a SIG-Node approver for https://github.com/kubernetes/enhancements/pull/1000 ?. Based on discussion in [Slack](https://kubernetes.slack.com/archives/C0BP8PW9G/p1556584219050300), should we just approve within SIG-Windows? (@patricklang)
    *   Since implementing CRI, not changing it for this KEP, will remove SIG-Node as approver and keep @yujuhong on as reviewer
*   Graceful timeout issue and CRI changes needed - Mrunal


# Apr 23



*   


# Apr 16



*   Removing cloud provider info from the /spec endpoint - [#76291](https://github.com/kubernetes/kubernetes/pull/76291)
    *   More generally: should we deprecate the kubelet endpoints tied to cadvisor? Candidates include: spec, non-summary stats, /metrics/cadvisor
    *   Filed https://github.com/kubernetes/kubernetes/issues/76660 for deprecation
*   Bringing CRI-ContainerD to Windows ([doc](https://docs.google.com/document/d/1NigFz1nxI9XOi6sGblp_1m-rG9Ne6ELUrNO0V_TJqhI/edit#)) (@patricklang)
    *   Background, use cases and some key decision points outlined
    *   Ask: Close on list of reviewers this week so this can move to a KEP
    *   Ask: For the proposals that _require CRI changes_, designate team & process to move forward. KEP or ?
*   Who owns the policy around revendoring the Docker API (see [PR](https://github.com/kubernetes/kubernetes/pull/75843))? Is it SIG-Node? (@patricklang) This broke Windows tests across the board from 4/12 to 4/14 when we set DOCKER_API_VERSION to override it, but could have been prevented with testing ([PR](https://github.com/kubernetes/test-infra/pull/12010) with /test trigger in review). Can we get someone to help push this PR through so /test works and we can prevent this next time?
    *   What about pinning Docker API version so revendoring doesn’t change it? 
*   FYI, [KEP for Ephemeral Containers](https://git.k8s.io/enhancements/keps/sig-node/20190212-ephemeral-containers.md) has moved to implementable. Please chime in with any feedback.


# Apr 09



*   Meeting canceled due to lack of agenda topics
*   Please reach out on slack if blocked


# Apr 02



*   Vertical Scaling KEP: https://github.com/kubernetes/enhancements/pull/686#
    *   @DerekCarr review latest updates, and sign-off if all concerns are addressed.
    *   @PatrickLang review suggested CRI changes for Windows compat
    *   Discuss any other concerns with merging this KEP, and going to next-steps.
*   Critical container KEP: https://github.com/kubernetes/enhancements/pull/912
    *   Please check [Sidecar Kep](https://github.com/kubernetes/enhancements/blob/master/keps/sig-apps/sidecarcontainers.md) to see whether it solves the use case here, see how this compares to https://github.com/kubernetes/enhancements/pull/919
*   Update docker/docker dependency in Kubernetes: https://github.com/kubernetes/kubernetes/pull/75843#issuecomment-478186871


# Mar 26



*   RuntimeClass Scheduling KEP review: (tallclair) https://github.com/kubernetes/enhancements/pull/896 \
https://github.com/kubernetes/enhancements/pull/909
*   PR do discuss : [#71786](https://github.com/kubernetes/kubernetes/pull/71786) importance of admit pod, by sorting the pods based on created timestamp. is it necessary to sort pods to admit ??

        PR which is working fine without the pods sorting [#75715](https://github.com/kubernetes/kubernetes/pull/75715) 


        Google discussion reference [(g-group discussion)](https://groups.google.com/forum/?utm_medium=email&utm_source=footer#!msg/kubernetes-sig-node/9e93aO0D2C8/fQVnMXbbAgAJ)

*   Vertical Scaling follow-up: https://github.com/kubernetes/enhancements/pull/686#
    *   Identify additional reviewer(s): 
    *   Updated KEP review
        *   Revised flow control
        *   Notes on handling Pod Overhead


# Mar 19



*   v1.15 topics to tackle (pair up owners and approvers)
    *   draft:https://docs.google.com/document/d/1Qdba1lM7H5eovFyV5Mp-bXplIWT42FXDRWRDzloH-4s/edit?usp=sharing
*   Vertical Scaling KEP review: https://github.com/kubernetes/enhancements/pull/686
    *   Outstanding issues: memory reduction, memory backed emptyDir, backward compat
*   RuntimeClass Scheduling KEP review: https://github.com/kubernetes/enhancements/pull/896


# Mar 12th 

Cancelled


# Mar 5th



*   Add [maxInitialFailureCount](https://github.com/matthyx/enhancements/blob/master/keps/sig-node/20190221-maxInitialFailureCount-health-probes.md) to health probes (@matthyx)
    *   target for v1.15
    *   Next Step: find reviewers from sig-node for this KEP 
*   NUMA Manager: https://github.com/kubernetes/enhancements/issues/693#issuecomment-466728227 
    *   Code Tracking Issue here: https://github.com/kubernetes/kubernetes/issues/72828
    *   since the code freeze for v1.14 is coming in two days, this feature should target for v1.15.
    *   Next Step: update kep to reflect the delay
*   RuntimeClass beta update [@tallclair] - Summary of changes requested in API review & implications for upgrades.
    *   move RuntimeClass API from CRD to a core API, since CRD is not fully ready yet; previously existing RuntimeClass objects will need to be created;
    *   rename `runtime_handler` to `handler`, and make the `handler` field as required; this only affects RuntimeClass objects; the RuntimeClassName field of PodSpec can still be left empty;
    *   node e2e tests will be added for beta; no conformance tests will be added for now;
*   do we currently track the review effort from the sig-node side for the sig-windows PRs for v1.14 (@Patrick)?
    *   currently, we don’t. The process for v.14 is very fluid. Please ping sig-node if help is needed.
    *   We will improve this for v1.15
*   Runc memory issue: https://github.com/opencontainers/runc/issues/1980 
    *   Fix under review: https://github.com/opencontainers/runc/pull/1984
    *   Basically, instead of copying the runc binary into memory or `/tmp`, it creates a temporary read-only bind mount into runc state directory to avoid extra memory usage. 
    *   The fix is still under review.
    *   


# Feb 26th



*   [Separating a CRI for docker from Kubelet](https://groups.google.com/d/msg/kubernetes-sig-node/0qVzfugYhro/l6Au216XAgAJ) (@dims)
    *   Previous SIG Node decision: container runtime is part of the os distro and node image. In the future, the container runtime in the test infra is provided and supported by the os distro and node image vendors. (@DawnChen)
    *   There is no rush to make decision on this since many productions are still depending on DockerShim even both cri-containerd and cri-o are productionized. But it is worthy talking about the current status and the goal. (@DawnChen)
    *   Sig-windows wants to move to cri-containerd as fast as possible. Some features only exist in cri-containerd, and hard to add into dockershim, e.g. some RuntimeClass features. (@Patrick)
    *   It is not the time yet to make the decision, we probably need another year. (@Derek)
    *   We already started adding features not supported by dockershim, and there will be more, e.g. secure pod, runtime class. At least if people agree on that we are going to eventually deprecate dockershim, we can continue adding this kind of features. (@yujuhong)
    *   Looking at about a year deprecation clock, once blockers are satisfied.
    *   We can start thinking about and document dockershim deprecation criteria. That doesn’t need to wait until the clock timer starts. (@yujuhong)
    *   Would really like to see this in a KEP showing how it lines up with graduating RuntimeClass, getting CRI to 1.0, and giving time for other runtimes to catch up that should be a well communicated part of sig-node roadmap (@Patrick)
    *   Who will be the advocate for when we know this is the right thing to do?
    *   Who can drive the KEP? @resouer volunteered to write a KEP for this. (@DawnChen)
    *   Next Step: we will come back to this topic one month after the KEP is sent to the community. (@DawnChen)
*   Follow up discussion on [pod resource overhead issue](https://github.com/kubernetes/enhancements/issues/688)/[RFC doc](https://docs.google.com/document/d/1EJKT4gyl58-kzt2bnwkv08MIUZ6lkDpXcxkHqCvvAp4/edit#heading=h.b24dqducgii8) (@egernst)
    *   overview [presentation](https://docs.google.com/presentation/d/1TNzmWcFMgUsdc59oAPeV_50RVrhal7qSI17FPRsZ7bU/edit#slide=id.g4cee2d5f0f_0_90) 
    *   introduce a pod overhead to the PodSpec, which will be taken into account along with the container requests.  This will be populated automatically by an admission controller, which utilizes the values from RuntimeClass CRD.  Would like to intercept Kubernetes 1.15 as a separate feature.
    *   Why not just add pod level resource limit/requests? (@Derek)
    *   Next step: continue the discussion on the KEP
*   Warning / Announcement runc memory spike issue after the CVE fix, pod with low memory limit (<10mb) may not run anymore https://github.com/opencontainers/runc/issues/1980. (@DawnChen)


# Feb 19th



*   [Promoting cloud provider node labels to GA](https://github.com/kubernetes/enhancements/pull/839) (@andrewsykim)
    *   Approved by SIG Node
*   Deprecate “containerized” kubelet - [#74148](https://github.com/kubernetes/kubernetes/issues/74148) (@dims)
    *   Yes from SIG Node
    *   Only change how to run Kubelet for HyperKube, not deprecate HyperKube.
    *   We should announcing this to kubernetes-dev@ and Kubernetes Community meeting


# Feb 12th



*   Discuss [Sidecar KEP](https://github.com/kubernetes/enhancements/issues/753) (@jirving)


# Feb 5th



*   Discuss [Sidecar KEP](https://github.com/kubernetes/enhancements/issues/753) (@jirving)
    *   Add implementation details to KEP
*   Issue [#73707](https://github.com/kubernetes/kubernetes/issues/73707): Include RuntimeClass information in [PodSandbox](https://github.com/kubernetes/kubernetes/blob/f52713515b8d526feed0f9111b3f075aaf463001/pkg/kubelet/apis/cri/runtime/v1alpha2/api.proto#L480) and [PodSandboxStatus](https://github.com/kubernetes/kubernetes/blob/f52713515b8d526feed0f9111b3f075aaf463001/pkg/kubelet/apis/cri/runtime/v1alpha2/api.proto#L422) (haiyanmeng@ and tallclair@)
*   UserNS remapping: [updated proposal](https://github.com/kubernetes/community/pull/2595) (vikasc)
*   Discuss `oom_score_adj` based, in part, on pod priority (seth)
*   Add pod name and namespace in the CRI container log path (lantaol@) [#73503](https://github.com/kubernetes/kubernetes/issues/73503)
*   NodeAllocatablePIDLimits (@rkrawitz)
    *   WIP https://github.com/kubernetes/kubernetes/pull/73651


# Jan 29



*   Discuss [Sidecar Kep](https://github.com/kubernetes/enhancements/blob/master/keps/sig-apps/sidecarcontainers.md),[ (initial discussion PR](https://github.com/kubernetes/community/pull/2148)) - API implementation in particular needs to be decided (@jirving, @enisoc)
    *   Tracking [issue opened](https://github.com/kubernetes/enhancements/issues/753)
    *   Discuss again at next week’s meeting
    *   @jirving to investigate how wide reaching the kubelet changes will be
*   Discuss [Kubelet Resource Metrics Endpoint KEP](https://github.com/kubernetes/enhancements/pull/726) (dashpole@) [Slides](https://docs.google.com/presentation/d/14zM8S7Ftymo3OabGc208EIjLCXpDheA8yjVV7hWUr2M/edit?usp=sharing)
*   Discuss/review Kubelet and CRI portions of [Windows GMSA KEP](https://github.com/kubernetes/enhancements/blob/master/keps/sig-windows/20181221-windows-group-managed-service-accounts-for-container-identity.md) [jean, ddebroy, jeremy]. Specifically need two questions for SIG Node clarified at to reach “implementable” status: https://github.com/kubernetes/enhancements/pull/710#issuecomment-457745480:
    *   Is the proposal in the KEP to add (alpha, feature-gated) code in kubelet/dockershim to drive CRI behavior from the pod annotation considered acceptable?
    *   Is there a goal to reconcile the capabilities of the pod API, the CRI API, and the OCI spec? if so, has there been discussion about what that would look like (and if what this KEP proposes looks like it could align with that)?
*   Status Note: Windows stable graduation criteria are also in a KEP [743](https://github.com/kubernetes/enhancements/pull/743) which is in final review to be marked implementable for 1.14. Yu-ju & Dawn Reviewing, already have multiple rounds of updates made based on feedback from Brian Grant & Yu-ju.
*   Report usageNanoCores for CRI runtimes (PatrickLang@) https://github.com/kubernetes/kubernetes/issues/72803 https://github.com/kubernetes/kubernetes/issues/72788
*   Minor: Consider deprecating the --register-node flag, and just use --kubeconfig flag as the signal to register or not: https://github.com/kubernetes/kubernetes/issues/61656 (@mtaufen)
*   HugePages KEP / feature issue: https://github.com/kubernetes/enhancements/issues/751 (@derekwaynecarr)


# Jan 22



*   RuntimeClass v1.14 plans
    *   RuntimeClass to Beta - requirements? https://github.com/kubernetes/enhancements/pull/697
    *   [Revisit upgrade story](https://docs.google.com/document/d/1hH99thfjpVDG77GgBUT0yXtHb1u8E0zcT_kviwjG1k8/edit?usp=sharing)
*   Review https://github.com/kubernetes/enhancements/pull/686 ( @vinaykul )
    *   Discuss memory reduction, emptyDir
    *   Any other major concerns
*   NodeSelectors - is this a practice SIG-Node recommends?
    *   “kubelet: promote OS & arch labels to GA” merged https://github.com/kubernetes/kubernetes/pull/73048 
    *   PR asking to constrain CoreDNS to Linux nodes https://github.com/kubernetes/kubernetes/pull/69940 


# Jan 15



*   UserNS remapping: [updated proposal](https://github.com/kubernetes/community/pull/2595),[ implementation PR](https://github.com/kubernetes/kubernetes/pull/64005) and  [Demo](https://www.youtube.com/watch?v=2sEmceRNa6Y) (vikasc)
*   PLEG relist has a race with event channel: [#72482](https://github.com/kubernetes/kubernetes/issues/72482) ，WIP PR: [#72709](https://github.com/kubernetes/kubernetes/pull/72709) (@resouer)
*   lts discussion https://github.com/kubernetes/community/pull/2911 (tpepper)
    *   Plan to formalize the workgroup
    *   Sig Node is one of stakeholders 
*   v1.14 release team checkin (spiffxp)
    *   If possible can we run through this in first 30 mins? sched conflict afterwards
    *   I’m here in lieu of Claire Laurence, v1.14 Enhancements Lead
    *   **Enhancement Freeze Jan 29**
    *   All kubernetes/enhancements issues targeted for v1.14 must have a KEP, even if they didn’t before
    *   KEP graduation criteria should include a checklist of requirements for alpha/beta/stable, including test plan and upgrade/downgrade plan
    *   Review of [kubernetes/enhancements issues targeted for v1.13](https://github.com/kubernetes/enhancements/issues?utf8=%E2%9C%93&q=is%3Aissue+is%3Aopen+milestone%3Av1.13+label%3Asig%2Fnode), what needs to be moved to v1.14, what can be closed?
        *   Dawn: Done. Re-targeted some to v1.14, and closed others. 
    *   Review of [kubernetes/enhancements issues targeted for v1.14](https://github.com/kubernetes/enhancements/issues?utf8=%E2%9C%93&q=is%3Aissue+is%3Aopen+milestone%3Av1.14+label%3Asig%2Fnode+), are these accurate?
        *   Will file the rest of features. 
*   Graduate HugePages to GA, any objections? (@derekwaynecarr)
    *   https://github.com/kubernetes/kubernetes/pull/72785
*   ReplicaSet controller continuously creating pods failing due to SysctlForbidden (@Suraj) [#kubernetes/72593](https://github.com/kubernetes/kubernetes/issues/72593)
    *   ReplicaSet controller should know about the pod failure
    *   Controllers should be smarter to handle these errors and backoff
    *   If an user is using sysctl then it is assumed they are breaking the containers and pod boundaries that kubernetes endorses
    *   Ideal to solve it at the controller level than at the kubelet level
*   Initial review of PR https://github.com/kubernetes/enhancements/pull/686 (@vinaykul)


# Jan 08



*   Pulls of note
    *   Remove deprecated AllowPrivileged, HostNetworkSources, etc. options - https://github.com/kubernetes/kubernetes/pull/71835 (@tallclair, @charrywanganthony)
    *   Inconsistent termination message output https://github.com/kubernetes/kubernetes/pull/71416 (@pontiyaraja)  [google group post](https://groups.google.com/forum/?utm_medium=email&utm_source=footer#!msg/kubernetes-sig-node/wQhDyyRAV2U/pOwY0Ql5FQAJ)
    *   Promote `ValidateProxyRedirects` to Beta and enable by default: https://github.com/kubernetes/kubernetes/pull/72552 (@tallclair)
*   GMSA support for Windows: https://github.com/kubernetes/enhancements/pull/666 (@jeremywx, @jean, and @ddebroy)
*   UserNS remapping: [updated proposal](https://github.com/kubernetes/community/pull/2595),[ implementation PR](https://github.com/kubernetes/kubernetes/pull/64005) and  [Demo](https://www.youtube.com/watch?v=2sEmceRNa6Y) (vikasc)
*   Requesting review on pull of in-place update of Pod resources: https://github.com/kubernetes/community/pull/2908/commits/4ad6fa7c27f4a21c27a6be83c2dc81c43549fa55
