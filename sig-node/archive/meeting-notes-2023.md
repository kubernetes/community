# SIG Node Meeting Notes

## Dec [all dates] - meetings canceled, next meeting Jan 2nd 2024


## Dec 5, 2023

Recording: [https://youtu.be/k7kzjKnmwSI](https://youtu.be/k7kzjKnmwSI) 



* [tzneal] Feedback on OOM cgroup killing change [https://github.com/kubernetes/kubernetes/pull/117793#issuecomment-1837400726](https://github.com/kubernetes/kubernetes/pull/117793#issuecomment-1837400726) 
    * Let’s have a kubelet option to return the old behavior. New behavior is a default
    * Need to cherry-pick to 1.27
* [haircommander] [https://github.com/kubernetes/kubernetes/pull/114847](https://github.com/kubernetes/kubernetes/pull/114847) follow ups 
    * summary of proposed policy changes.. pull-never policy pods must have the same cred to re-use an image that was pulled with a cred; kubelet needs a switch to disable validation checking of in the cache preloaded images (for disconnected mode at node / kubelet restart), otherwise images pulled with not present are subject to revalidation, and pull never will fail if never authenticated; a pod that successfully pulls an image anonymously from registry A(or default) is to be considered “unique”.. we will not use that anonymous pull as an anonymous success for pods pulling from another registry.. requires alg change in current feature implementation
    * Derek: let’s consider other patterns for time slicing the auth and pull time slots, policies for specifying when and possibly for what reason we need to auth.. (align better with disconnected needs, not just performance/multi-tenant)
    * future items: possible integration with registries for discovering (header/artifacts) what the expiration is for an image
* [SergeyKanzhelev] Planning of 1.30 \
Eliminating perma betas. List of “old” feature gates:
    * AppArmor
        * Mostly need to clean up tests
        * Sergey to follow up
    * CustomCPUCFSQuotaPeriod - Peter will take a look
    * GracefulNodeShutdown
        * Issues with some controllers - Ryan add a comment on KEP indicating what the issue is.
    * GracefulNodeShutdownBasedOnPodPriority
    * LocalStorageCapacityIsolationFSQuotaMonitoring 
        * Kernel issue. Other bugs reported
        * [https://github.com/kubernetes/kubernetes/pull/112626](https://github.com/kubernetes/kubernetes/pull/112626) 
        * Mrunal and Peter  to check on kernel issue
    * MemoryManager
        * Got as far as PRR review. Lack of observability is concerning from PRR review - need to work on this.
        * Fracesco will follow up on this.
        * **Swati**: many issues opened for MemoryManager before GA
            * totally true we need to address them - silver lining is this can be done in parallel with observability improvements
    * MemoryQoS
        * Cut the feature
        * Experimentation may go on using NRI: [https://containers.github.io/nri-plugins/stable/docs/memory/memory-qos.html](https://containers.github.io/nri-plugins/stable/docs/memory/memory-qos.html) Antti and Dixita made this to allow further experiments
        * Ping Dixi to see if she can cut it
    * PodAndContainerStatsFromCRI
        * Stalled on CRI implementation of those metrics
        * Working on it in CRI-O
        * Need some help from Containerd side
        * Exit criteria: must test performance that is not regressing
    * RotateKubeletServerCertificate
        * No tests and docs
        * Need volunteer to clean it up
        * Harche will take a look
    * SizeMemoryBackedVolumes
        * Need volunteers

    Deprecations:

    * cgroup v1
    * Mrunal, Dawn: 
        * let’s announce **deprecation in 1.30**.
        * Default to cgroupv2 in tests and have cgroupv1 as an “additional”
        * [Alexander Kanevsky] Collect list of distress that people uses, their default cgroups, and the EOL of those disttros. e.g. centos 7 or some ubuntu lts.
* [SergeyKanzhelev] Should we cancel all the rest of the meetings for the month of Dec?
    * Let;s cancel meeting till the end of the year and meet on Jan 2nd
* [hakman] node-problem-detector maintainers are needed to keep the project alive. I tried to follow the guidelines to step up as a reviewer and later approver, but it seems there is a lack of approvers. If possible, I would like someone from #sig-node to sponsor me. Thanks in advance! [https://github.com/kubernetes/node-problem-detector/pull/830](https://github.com/kubernetes/node-problem-detector/pull/830)
    * AI: SergeyKanzhelev to follow up


## Nov 28, 2023 [Canceled]



* Canceled due to lack of agenda.
* Please bring 1.30 planning topics for the next meeting.


## Nov 21, 2023 [Cancelled]



* Cancelled due to Thanksgiving week in US and leads availability


## Nov 14, 2023 [Cancelled]



* Cancelling due to lack of agenda and no bugs marked for the milestone.
* Remember to send your docs PR!
* [vinaykul] Notes from the Kubernetes Contributor Summit session on the proposal to rely on PodStatus as source of truth instead of node-local store: [https://static.sched.com/hosted_files/kcsna2023/a4/KCS-LeanKubelet.pdf](https://static.sched.com/hosted_files/kcsna2023/a4/KCS-LeanKubelet.pdf) 


## Nov 7, 2023 [Cancelled]



* Canceled due to kubecon 


## October 31, 2023

Recording: [https://youtu.be/RYBb81l1IGw](https://youtu.be/RYBb81l1IGw)



* [SergeyKanzhelev] Sidecar KEP change: proposed another feature gate for restarts.
    * AI: plan sounds good
* [MaRosset] [KEP 4216: Add changes for alpha version under RuntimeClassInImageCriApi feature gate](https://github.com/kubernetes/kubernetes/pull/121456) has LGTM’s needs approvals
    * AI: Mrunal and Derek to take another look. No issues were highlighted
* [haircommander/Peter Hunt] Drop-in Kubelet config complications
    * [https://github.com/kubernetes/kubernetes/pull/121193/](https://github.com/kubernetes/kubernetes/pull/121193/) 
* [Mrunal] [https://github.com/kubernetes/kubernetes/pull/120616](https://github.com/kubernetes/kubernetes/pull/120616) 
* [Jeffwan] in-place vpa changes need review and approval
    * [https://github.com/kubernetes/kubernetes/pull/112599](https://github.com/kubernetes/kubernetes/pull/112599) windows. done
    * [https://github.com/kubernetes/kubernetes/pull/120145](https://github.com/kubernetes/kubernetes/pull/120145) approved
    * [https://github.com/kubernetes/kubernetes/pull/120432](https://github.com/kubernetes/kubernetes/pull/120432)
* [vinaykul] If any K8s contributors are attending Contributor Summit next week and are interested in a change I’m proposing to Kubelet (to move away from node-local checkpointing for in-place pod resize and use PodStatus instead) please drop by [https://sched.co/1Sp9Z](https://sched.co/1Sp9Z) to discuss this.


## October 24, 2023

Canceled due to an empty agenda. Review PRs for freeze next week. 


## October 17, 2023

Recording: [https://youtu.be/740kJACH3i8](https://youtu.be/740kJACH3i8)



* [Kevin Hannon] Split Image API Issue
    * CRI API uses the ImageFsInfoResponse
    * Services.go uses a shortcut to get the image_filesystem
    * Path forward?
    * CRI API: [https://github.com/kubernetes/kubernetes/pull/120914](https://github.com/kubernetes/kubernetes/pull/120914) 
    * Outcome: Good path forward.
    * Questions on guarantee of support for services.go
* [Jiaxin] In-place vpa related PRs need reviewers
    * [fix inplace VPA stuck in InProgress when custom resources are specified kubernetes#120145](https://github.com/kubernetes/kubernetes/pull/120145)
    * [Enhance InPlacePodVerticalScaling performance kubernetes#120432](https://github.com/kubernetes/kubernetes/pull/120432)
    * [Configure MemoryRequest for InPlace pod resize in cgroupv2 systems #121218](https://github.com/kubernetes/kubernetes/pull/121218)
    * [Handle the case where container resource and request set to minimum values](https://github.com/kubernetes/kubernetes/pull/120791)
* [SergeyKanzhelev] [https://github.com/kubernetes/community/issues/7234](https://github.com/kubernetes/community/issues/7234) 
    * [backup-item]<span style="text-decoration:underline;"> [important-soon & LGTMed PR list](https://github.com/kubernetes/kubernetes/pulls?q=is%3Apr+is%3Aopen+label%3Asig%2Fnode+label%3Algtm+-label%3Aapproved+-label%3Ado-not-merge%2Fhold+-label%3Aneeds-rebase+-label%3Ado-not-merge%2Fwork-in-progress+label%3Apriority%2Fimportant-soon+)</span> 
    * Also re-starting the PRs stats tracking:	

<table>
  <tr>
   <td>
Total active pull requests:
   </td>
   <td><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+is%3Aopen+">311</a>
   </td>
  </tr>
</table>


(weekly changes)


<table>
  <tr>
   <td><strong>Incoming</strong>
   </td>
   <td>
   </td>
   <td><strong>Completed</strong>
   </td>
   <td>
   </td>
  </tr>
  <tr>
   <td>Created:
   </td>
   <td><strong><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++created%3A2023-10-10T17%3A00%3A00%2B0000..2023-10-17T16%3A53%3A09%2B0000">41</a></strong>
   </td>
   <td>Closed:
   </td>
   <td><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++is%3Aunmerged+closed%3A2023-10-10T17%3A00%3A00%2B0000..2023-10-17T16%3A53%3A09%2B0000">10</a>
   </td>
  </tr>
  <tr>
   <td>Updated:
   </td>
   <td><strong><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+updated%3A2023-10-10T17%3A00%3A00%2B0000..2023-10-17T16%3A53%3A09%2B0000+created%3A%3C2023-10-10T17%3A00%3A00%2B0000">113</a></strong>
   </td>
   <td>Merged:
   </td>
   <td><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++merged%3A2023-10-10T17%3A00%3A00%2B0000..2023-10-17T16%3A53%3A09%2B0000">38</a>
   </td>
  </tr>
</table>


	New stats needed:



* PRs needed other SIG approvals
* Waiting for approvers
* Waiting for reviewers
* Separate cherry-picks and regressions


## October 10, 2023

Recording: [https://www.youtube.com/watch?v=akrWtsCbJZo](https://www.youtube.com/watch?v=akrWtsCbJZo)



* [Adrian Reber] Graduate KEP 2008 "Forensic Container Checkpointing" from Alpha to Beta: [https://github.com/kubernetes/enhancements/pull/4288](https://github.com/kubernetes/enhancements/pull/4288) (1.30)
* [mfranzil () ] Addressing CRI-API image sizes & opening up new KEP ([https://github.com/kubernetes/kubernetes/issues/120698](https://github.com/kubernetes/kubernetes/issues/120698) and see Slack discussion)
    * [https://github.com/kubernetes/enhancements/issues/4191](https://github.com/kubernetes/enhancements/issues/4191) 
    * Image Disk: [https://github.com/kubernetes/enhancements/blob/master/keps/sig-node/4191-split-image-filesystem/README.md](https://github.com/kubernetes/enhancements/blob/master/keps/sig-node/4191-split-image-filesystem/README.md)
    * 
* [Kevin Hannon] Consensus on OOM flaky test issue
    * [https://github.com/kubernetes/kubernetes/issues/119600](https://github.com/kubernetes/kubernetes/issues/119600)
    * Skip the test? - [https://github.com/kubernetes/kubernetes/pull/121031](https://github.com/kubernetes/kubernetes/pull/121031)
    * Skip the check for OOM [Todd Neal] - [https://github.com/kubernetes/kubernetes/pull/120460](https://github.com/kubernetes/kubernetes/pull/120460)
    * Closing 121031 and will go forward with 120460.
* [mahamed] NPD test change issues
    * [https://github.com/kubernetes/kubernetes/pull/121007](https://github.com/kubernetes/kubernetes/pull/121007) 
* KubeCon SIgNode related talks highlights: [https://docs.google.com/document/d/12kAxL7HWiPIcNqcTIjfr70-aLJyiUeGaNZAiZILfPeM/edit](https://docs.google.com/document/d/12kAxL7HWiPIcNqcTIjfr70-aLJyiUeGaNZAiZILfPeM/edit)  


## October 3, 2023

Recording: [https://www.youtube.com/watch?v=HdIURTQSm7Q](https://www.youtube.com/watch?v=HdIURTQSm7Q) 



* [Weipeng([dastonzerg](https://github.com/dastonzerg)), Ian([iancoolidge](https://github.com/iancoolidge))] (Sry for cutting into the first slot as we only have availabilities at early meeting time) Path forward for “in CPU static policy, non-gu pods shouldn't run on reserved CPUs” [https://github.com/kubernetes/kubernetes/pull/118021](https://github.com/kubernetes/kubernetes/pull/118021) 
* [mo] thoughts on expanding [Kubelet Credential Providers](https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/2133-kubelet-credential-providers)
    * currently it only supports giving the plugin the image as input
        * we could enhance it and maybe the pod API to support sending bound SA tokens through the CredentialProviderRequest
        * Will require non-trivial changes to the caching logic used by the kubelet
    * original issue: [Support an image pull credential flow built on bound service account tokens · Issue #68810](https://github.com/kubernetes/kubernetes/issues/68810)
    * we deferred this to a later time in the original KEP [https://github.com/kubernetes/enhancements/pull/1406/files#r371886703](https://github.com/kubernetes/enhancements/pull/1406/files#r371886703)
    * Depends on [https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/2535-ensure-secret-pulled-images](https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/2535-ensure-secret-pulled-images) to make sure disk cache does not skip authz checks
        * Or we could require image pull to be set to always
    * Are credential providers configurable in cloud envs?  i.e. can I use a registry from vendor A on a kubelet run by vendor B?
    * Could we isolate this change to kubelet + credential providers?  i.e. no change to the Kube REST API?
    * Interested (Mo will ping regarding sig auth discussion and will chat with folks at KubeCon as well):
        * Mike Brown
        * Sergey Kanzhelev
        * Dixita Narang
        * Ruiwen
        * Harshal Patil (Harshal Patil (RedHat) on k8s slack) 
        * Peter Hunt (haircommander)
* [haircommander] revitalize proc mount KEP
    * [https://github.com/kubernetes/enhancements/issues/4265](https://github.com/kubernetes/enhancements/issues/4265) 
* ~~[klueska] Update CDI for device plugins KEP for beta graduation in 1.29~~
    * ~~Comments addressed, please review again: \
[https://github.com/kubernetes/enhancements/pull/4238#issuecomment-1739120658](https://github.com/kubernetes/enhancements/pull/4238#issuecomment-1739120658)~~


## September 26, 2023

Recording: [https://www.youtube.com/watch?v=yEOUKJCJXa8](https://www.youtube.com/watch?v=yEOUKJCJXa8)



* [Filip Krepinsky] Declarative Node Maintenance: discuss issues related to node drain and the solutions this KEP proposes
    * [https://github.com/kubernetes/enhancements/pull/4213](https://github.com/kubernetes/enhancements/pull/4213)
    * need to explore RBAC options for triggering the node drain (NodeMaintenance object)
* [Adrian Reber] Created PR to avoid filling up local disk space with too many checkpoint archives.
    * [https://github.com/kubernetes/kubernetes/pull/115888](https://github.com/kubernetes/kubernetes/pull/115888)
    * bringing to sig-node for awareness
    * requested by multiple users
    * functionality: if more than a certain number of checkpoints are created per container/pod/namespace (default 10), older checkpoint archives are deleted
* [klueska] Update CDI for device plugins KEP for beta graduation in 1.29
    * Updated issue: \
[https://github.com/kubernetes/enhancements/issues/4009](https://github.com/kubernetes/enhancements/issues/4009)
    * Pull Request with update: \
[https://github.com/kubernetes/enhancements/pull/4238](https://github.com/kubernetes/enhancements/pull/4238)
    * Update to planning doc: [SIG Node - KEP Planning](https://docs.google.com/document/d/1U10J0WwgWXkdYrqWGGvO8iH2HKeerQAlygnqgDgWv4E/edit#bookmark=kix.4hjk6cdmudfx)
* [marquiz] introducing [KEP-4112](https://github.com/kubernetes/enhancements/pull/4113) “Pass down resources to CRI”
    * better visibility pod resources in CRI
    * two goals
        * pass down all resources (of all containers) at sandbox creation
        * pass unmodified resource requests and limits to CRI
* [Kevin Hannon] PodReadyToStartContainers promotion to beta
    * [https://github.com/kubernetes/enhancements/pull/4139](https://github.com/kubernetes/enhancements/pull/4139)
    * Looking for an approver for KEP PR
    * Implementation: [https://github.com/kubernetes/kubernetes/pull/119659](https://github.com/kubernetes/kubernetes/pull/119659)
* [Kevin Hannon] Split Image Disk KEP
    * [https://github.com/kubernetes/enhancements/pull/4198](https://github.com/kubernetes/enhancements/pull/4198)
    * Any interest in separate image filesystems or deal with existing problems, please ping me or comment on KEP.
    * We are interested in hearing what we should consider in scope for this
* [katarzyna, robscott]  Kubelet API to return the local pods. Please review
    * [https://github.com/kubernetes/enhancements/pull/4184](https://github.com/kubernetes/enhancements/pull/4184)


## September 19, 2023

Recording: [https://www.youtube.com/watch?v=ngboQ3GvX5o](https://www.youtube.com/watch?v=ngboQ3GvX5o) 



* [haircommander] CRI stats metrics gaps
    * create a configurable kubelet option to tell cadvisor which cgroups to collect from to aid migration from cadvisor to cri stats, then consider declaring metrics as deprecated
    * 
* [Sergey Kanzhelev] [https://github.com/kubernetes-sigs/metrics-server/issues/1330](https://github.com/kubernetes-sigs/metrics-server/issues/1330) 
* [haircommander] kubelet.conf.d configuration view
    * [https://github.com/kubernetes/enhancements/blob/master/keps/sig-node/3983-drop-in-configuration/README.md?plain=1#L83](https://github.com/kubernetes/enhancements/blob/master/keps/sig-node/3983-drop-in-configuration/README.md?plain=1#L83) 
    * add a e2e test for beta to double check configz endpoint correctly reflects the drop-ins
* [vaibhav] Is there any current plan of deprecating the existing garbage collection flags in favor of eviction hard ?

    Ref: [https://github.com/kubernetes/design-proposals-archive/blob/main/node/kubelet-eviction.md#deprecation-of-existing-features](https://github.com/kubernetes/design-proposals-archive/blob/main/node/kubelet-eviction.md#deprecation-of-existing-features)

* [ndixita] 
* [swsehgal] [https://github.com/kubernetes/kubernetes/issues/116086](https://github.com/kubernetes/kubernetes/issues/116086) 
    * Exclusive CPU allocation at container scope.
    * Looking for feedback on a potential solution captured [here](https://github.com/kubernetes/kubernetes/issues/116086#issuecomment-1719469058).
    * Do we want to pursue this in the 1.29 cycle?
    * ~~[pacoxu] [116086#](https://github.com/kubernetes/kubernetes/issues/116086#issuecomment-1655184757) still need a direction: ~~
        * ~~1: add it to sidecar KEP: support core binding for pod with a non-guaranteed sidecar (sidecar will not affect QoS) . I +1 for this now.~~
        * ~~2: generally, core binding for container level. Not pod QoS level(Or add new policy for container level core binding than current `static`).~~
* [weipeng] [https://github.com/kubernetes/kubernetes/pull/118021](https://github.com/kubernetes/kubernetes/pull/118021) 
    * Needs review. Does [https://github.com/kubernetes/kubernetes/pull/118021#issuecomment-1722066630](https://github.com/kubernetes/kubernetes/pull/118021#issuecomment-1722066630) answer the concern [https://github.com/kubernetes/kubernetes/pull/118021#issuecomment-1721342181](https://github.com/kubernetes/kubernetes/pull/118021#issuecomment-1721342181) ?


## September 12, 2023

Recording: [https://www.youtube.com/watch?v=hVuZg2mqNsw](https://www.youtube.com/watch?v=hVuZg2mqNsw)



* KEPs planning ([doc](https://docs.google.com/document/d/1U10J0WwgWXkdYrqWGGvO8iH2HKeerQAlygnqgDgWv4E/edit))
    * [Talor] Volunteering to work on Memory Manager GA graduation in 1.29 cycle ([#1769](https://github.com/kubernetes/enhancements/issues/1769))
        * [swsehgal] Can help with reviews.
    * [katarzyna] kubelet API [KEP](https://github.com/kubernetes/enhancements/pull/4184#issuecomment-1713779858) 
    * [Alexey] [CRI pull image with progress](https://github.com/kubernetes/enhancements/pull/3547) PRR
    * [Kevin Hannon] [Beta Promotion for PodReadyToStartContainers](https://github.com/kubernetes/enhancements/pull/4139)
    * [Kevin Hannon] [Splitting Image Filesystem](https://github.com/kubernetes/enhancements/issues/4191)
        * PR: [https://github.com/kubernetes/enhancements/pull/4198](https://github.com/kubernetes/enhancements/pull/4198)
    * [ndixita] KEP PR targeting 9/14 for PSI metric support.
    * [Jiaxin] [KEP-4176: Static Policy to spread hyperthreads across physical CPUs](https://github.com/kubernetes/enhancements/pull/4177)
* [vaibhav] Is there any current plan of deprecating the existing garbage collection flags in favor of eviction hard ?

    Ref: [https://github.com/kubernetes/design-proposals-archive/blob/main/node/kubelet-eviction.md#deprecation-of-existing-features](https://github.com/kubernetes/design-proposals-archive/blob/main/node/kubelet-eviction.md#deprecation-of-existing-features)



## September 5th, 2023

Recording: [https://www.youtube.com/watch?v=5iiD9OIeJv8](https://www.youtube.com/watch?v=5iiD9OIeJv8)



* [pacoxu] Should we support static cpu policy for pod with a non-guaranteed sidecar? Details can be found at [116086#issuecomment-1655184757](https://github.com/kubernetes/kubernetes/issues/116086#issuecomment-1655184757). I have no context about the cpu static policy initial design, it is container level cpu mapping, but depends on a pod level QoS. This is a feature request IMO.
    * [Swati’s comment ](https://github.com/kubernetes/kubernetes/issues/116086#issuecomment-1706386179)and Francesco’s [comment](https://github.com/kubernetes/kubernetes/issues/116086#issuecomment-1706678147) are most likely the answer.
    * AI: gather more use cases, generalize to cover most of them.
* [klueska] Update DRA KEP to be inline with changes made in 1.28 (2 PRs)
    * I am unfortunately unable to attend today, but please review the following PRs to update the DRA KEP to reflect the latest code changes from 1.28
    * [https://github.com/kubernetes/enhancements/pull/4063](https://github.com/kubernetes/enhancements/pull/4063)
    * [https://github.com/kubernetes/enhancements/pull/4164](https://github.com/kubernetes/enhancements/pull/4164)
* [rphillips] Node Graceful Shutdown Issue
    * Pod termination is taking too long
    * E2E tests do not run with graceful shutdowns enabled
        * CI reboots are problematic
    * [PR#120273](https://github.com/kubernetes/kubernetes/pull/120273)
    * [Issue#120271](https://github.com/kubernetes/kubernetes/issues/120271)
    * [Presentation](https://docs.google.com/presentation/d/1aU9hLgza5EF62NmNemCc8llQP3tb9vqaF9PMjWl0Q3k/edit?usp=sharing)
    * AI: [rphillips] look into termination with drains, make sure they are similar
* [katarzyna, robscott] KEP for new Kubelet API for local Pod readiness: [https://github.com/kubernetes/enhancements/pull/4184](https://github.com/kubernetes/enhancements/pull/4184)
* [jeffwan, LingyanYin] Follow up on [Inplace VPA + core binding](https://docs.google.com/document/d/1V3DLh3pH3CD-xhhJvAnOq_oWgPyjO-vj6wY6qdew9H0/edit#heading=h.ybybfdfputt)
    * [https://github.com/kubernetes/kubernetes/pull/120145](https://github.com/kubernetes/kubernetes/pull/120145)
    * [https://github.com/kubernetes/kubernetes/pull/120432](https://github.com/kubernetes/kubernetes/pull/120432)
    * [https://github.com/kubernetes/enhancements/issues/4176](https://github.com/kubernetes/enhancements/issues/4176)


## August 29th, 2023

Recording: [https://www.youtube.com/watch?v=tNangR9QLkg](https://www.youtube.com/watch?v=tNangR9QLkg)



* [kannon92]: Follow up for PodReadyToStartContainers Beta
    * KEP PR for beta - [https://github.com/kubernetes/enhancements/pull/4139](https://github.com/kubernetes/enhancements/pull/4139)
    * k/k PR - [https://github.com/kubernetes/kubernetes/pull/119659](https://github.com/kubernetes/kubernetes/pull/119659)
* [karthik-k-n]: As discussed earlier, Shall we have separate meeting to discuss on scope for [dynamic node resize](https://docs.google.com/document/d/1KfjPRmCc8Xk0xxa4S8ZRle6VMzc1C6MQg4ivOaoB150/edit)
* [katarzyna, robscott]  
* [sunnylovestiramisu] Can I use Node -> NodeStatus -> NodePhase(NodePhase is the recently observed lifecycle phase of the node) as evidence that a node is registrationCompleted? If yes, which phase should I use? If not, can I add another status called NodeRegistered to the NodePhase and update it while we set registrationCompleted? - [context issue](https://github.com/kubernetes/kubernetes/issues/120063).
* [weipeng] Need attention on PR `Fix: Exclude reserved CPUs from shared pool.` Currently the `pull-kubernetes-node-kubelet-serial-cpu-manager` test lane itself has some issues, how should we proceed? [https://github.com/kubernetes/kubernetes/pull/118021](https://github.com/kubernetes/kubernetes/pull/118021) 


## August 22th, 2023

Recording: [https://www.youtube.com/watch?v=Y9btZGnyDK0](https://www.youtube.com/watch?v=Y9btZGnyDK0)



* [rata]: In 1.28 we added support for stateful pods with user namespaces.
    * Do we want to blog about it?
    * [rata]: I’ll create a gdocs and share it here. Once that is finalized, I’ll open a PR to the website.
    * If we can’t find someone from sig-docs to approve it, Sergey can help.
* ~~[fromani]notification: approvers PTAL to these backports - all lanes fixed, tests passing, LGTMd [https://github.com/kubernetes/kubernetes/pull/119432](https://github.com/kubernetes/kubernetes/pull/119432) [https://github.com/kubernetes/kubernetes/pull/119706](https://github.com/kubernetes/kubernetes/pull/119706) [https://github.com/kubernetes/kubernetes/pull/119707](https://github.com/kubernetes/kubernetes/pull/119707)~~
* [haircommander] Image GC WG time decided–can we add to the calendar?
    * Wednesday 12-12:30 PST (3-3:30 EST)
* [tzneal] Kubelet detecting a readonly filesystem, what’s the boundary between node-problem-detector responsibilities and kubelet - [https://github.com/kubernetes/kubernetes/pull/115746](https://github.com/kubernetes/kubernetes/pull/115746) 


## August 15th, 2023

Recording: [https://www.youtube.com/watch?v=wgF8UDgp1sQ](https://www.youtube.com/watch?v=wgF8UDgp1sQ)



* [ruiwen] 1.28 KEPs [retro](https://docs.google.com/document/d/1NaT0rY0o1cNdTxIlgZ5m0TqLDI7AfYn3rBAAl4qT1Bw/edit) (at least 30 minutes. We may not have much time for many other topics)
* [haircommander] Kubelet image GC conversation
    * [Ruiwen] Pin images
    * [Derek] i am curious if secret pulled images have any unique gc requirements that have surfaced…
        * Tie lifecycle of image to lifecycle of pod?
    * [Sergey] Mirror config into kubelet?
    * Peter to begin a WG in between now and KEP freeze to come up with next steps before bringing to larger group.
        * Doodle for meeting time selection: [https://doodle.com/meeting/participate/id/eg2qn5jd](https://doodle.com/meeting/participate/id/eg2qn5jd) 
* [SergeyKanzhelev] Sidecar WG: join for the next push in 1.29:


## August 8th, 2023

Recording: [https://www.youtube.com/watch?v=9BBSMdw8dMA](https://www.youtube.com/watch?v=9BBSMdw8dMA)



* [jiaxin, zewei, lingyan] continued with the meeting on July 18th, we broke down the problems into 4 of them and the detailed doc is here [https://docs.google.com/document/d/1V3DLh3pH3CD-xhhJvAnOq_oWgPyjO-vj6wY6qdew9H0/edit#heading=h.by fdf putt](https://docs.google.com/document/d/1V3DLh3pH3CD-xhhJvAnOq_oWgPyjO-vj6wY6qdew9H0/edit#heading=h.ybybfdfputt)
    * [fromani] gonna comment with more details on the doc when comments enabled, the gist is: 1. did you tried also using the [cpumanager policy options](https://kubernetes.io/docs/tasks/administer-cluster/cpu-management-policies/#static-policy-options)? full-pcpus-only and distribute-cpus-across-numa may have a positive impact. 2. can the new policy be reimplemented as policy option?


## August 1st, 2023

Recording: [https://www.youtube.com/watch?v=V9F8jHgs6R4](https://www.youtube.com/watch?v=V9F8jHgs6R4)



* [kannon92]Taking over PodReadyToStartContainersCondition ([https://github.com/kubernetes/enhancements/issues/3085](https://github.com/kubernetes/enhancements/issues/3085))
    * new issue: https://github.com/kubernetes/enhancements/issues/4138
    * Beta promotion in 1.29
    * Any items we should have for beta?
    * PR: [https://github.com/kubernetes/kubernetes/pull/119659](https://github.com/kubernetes/kubernetes/pull/119659)
    * KEP update: [https://github.com/kubernetes/enhancements/pull/4139](https://github.com/kubernetes/enhancements/pull/4139)
* [mckdev] leaked kubelet leases
    * PR: [https://github.com/kubernetes/kubernetes/pull/119661](https://github.com/kubernetes/kubernetes/pull/119661)
    * Issues:
        * [https://github.com/kubernetes/kubernetes/issues/109777](https://github.com/kubernetes/kubernetes/issues/109777)
        * [https://github.com/kubernetes/kubernetes/issues/119660](https://github.com/kubernetes/kubernetes/issues/119660)
* [skrobul] world-writable terminationMessage files
    * PR: [https://github.com/kubernetes/kubernetes/pull/108076](https://github.com/kubernetes/kubernetes/pull/108076)
    * How to progress this PR?
    * Is the Feature Gate or KEP needed?
* [haircommander] CRI approvers
    * [https://github.com/kubernetes/kubernetes/pull/119566](https://github.com/kubernetes/kubernetes/pull/119566) 

    [Sunnatillo] Feedback on PR, kubelet: do not set CPU quota for guaranteed pods


		[https://github.com/kubernetes/kubernetes/pull/117030](https://github.com/kubernetes/kubernetes/pull/117030)

		[mrunal] the was the bug related to CFS quota, this PR not using the quota

		[dawn] there may be a problem in kernel

		[dawn] what kernel version is it tested?

            [vaibhav] Approach to resolve eviction manager issue

                            [https://github.com/kubernetes/kubernetes/issues/115201](https://github.com/kubernetes/kubernetes/issues/115201)

		[Dawn] Do you want to start enhancement?


    [vaibhav] ok 


## July 25th, 2023 [Cancelled due to lack of agenda]


## July 18th, 2023

Recording: [https://www.youtube.com/watch?v=0Uqq8jNSSDk](https://www.youtube.com/watch?v=0Uqq8jNSSDk)



* [ndixita] memory QoS Beta K8s 1.28  might be infeasible [https://docs.google.com/document/d/1mY0MTT34P-Eyv5G1t_Pqs4OWyIH-cg9caRKWmqYlSbI/edit#bookmark=id.qaybju6wvb05](https://docs.google.com/document/d/1mY0MTT34P-Eyv5G1t_Pqs4OWyIH-cg9caRKWmqYlSbI/edit#bookmark=id.qaybju6wvb05)
    * Requesting kernel experts here for discussion around memory.high memcg controller usage, signals for memory reclaim(pgscan, pgsteal from memory.stat?).
* [jiaxin] new CPU Manager static policy and in-place VPA improvements (performance, make it work with CPU Manager together), KEP or PR? 
    * Problem 1: noisy neighbor issue. We want to spread hyper thread across physical cores to get better performance.
    * Problem 2: In-place VPA currently doesn’t work with CPU Manager
    * Problem 2: In-place VPA sometimes takes up to a minute to finish scaling etc. **We will finish a doc with the problems and solutions for further discussion**.
    * [fromani] most likely a KEP<sup>+1</sup>, perhaps share a (preliminary) design doc in the community to outline the proposed scope and changes
    * [Dawn] Please start with a doc on the issue / problem statement and the suggested solution.
    * [Alex] Please separate in-place VPA improvements from CPU static policy. 


## July 11th, 2023

Recording: [https://www.youtube.com/watch?v=0ggcapGYwtc](https://www.youtube.com/watch?v=0ggcapGYwtc)



* [alexeldeib] cgroupv2/v1 node memory usage calculation/alignment [https://github.com/kubernetes/kubernetes/issues/118916](https://github.com/kubernetes/kubernetes/issues/118916)
    * aligned on anon + file (rss + cache)? [https://github.com/opencontainers/runc/pull/3933](https://github.com/opencontainers/runc/pull/3933)
    * testing/e2e suggestions to ensure similarity of v1/v2?
    * swap usage may have a similar issue, but will follow up separately.
* ~~[kklues] KEP Update needs approval~~
    * ~~[https://github.com/kubernetes/enhancements/pull/3915](https://github.com/kubernetes/enhancements/pull/3915)~~
    * ~~/cc ~~~~ since she approved the original KEP~~
* [karthik-k-n] looking forward for review and understand the next way forward for node dynamic cpu and memory resize KEP [https://github.com/kubernetes/enhancements/pull/3955](https://github.com/kubernetes/enhancements/pull/3955)
* [Arka] Exploring [https://github.com/kubernetes/kubernetes/issues/116662](https://github.com/kubernetes/kubernetes/issues/116662) 
    * Understanding the issue and starting with KEP
* [SergeyKanzhelev] Sidecar PR got merged. Watch for problems and let us know.
    * Uber issue: [https://github.com/kubernetes/kubernetes/issues/115934](https://github.com/kubernetes/kubernetes/issues/115934) 
* [Marosset] - [https://github.com/kubernetes/kubernetes/pull/116968](https://github.com/kubernetes/kubernetes/pull/116968) (cri-only stats implementation for Windows) needs sig-node reviews
    * [haircommander] I will review this today
    * [haircommander] shameless plug: while CRI stats are on the mind: [https://github.com/kubernetes/kubernetes/pull/118838](https://github.com/kubernetes/kubernetes/pull/118838) 
    * mentioned containerd PR.. [https://github.com/containerd/containerd/pull/8671](https://github.com/containerd/containerd/pull/8671) will merge end of day and mark for cherry picking to 1.7/1.6
*  [Weipeng] `pull-kubernetes-node-kubelet-serial-cpu-manager` failure in PR [https://github.com/kubernetes/kubernetes/pull/118021#issuecomment-1630728455](https://github.com/kubernetes/kubernetes/pull/118021#issuecomment-1630728455)
* [pacoxu] [ci-crio-cgroupv1-node-e2e-eviction](https://testgrid.k8s.io/sig-node-cri-o#ci-crio-cgroupv1-node-e2e-eviction):[ PR #119097](https://github.com/kubernetes/kubernetes/pull/119097), Issue [#119090](https://github.com/kubernetes/kubernetes/issues/119090) PriorityPidEvictionOrdering should eventually evict all of the correct pods
    * containerd ci is green(flakes for eviction order that is caused by no process stats available in /stats/summary) and cri-o failed for no PIDPressure Condition happen quickly.


## July 4th, 2023 [Canceled due to US holiday]


## June 27th, 2023

Recording: [https://www.youtube.com/watch?v=KMD17c5EbFU](https://www.youtube.com/watch?v=KMD17c5EbFU)



* [Wedson] Discuss setting a default runtime handler for CRI image operations if no runtimeclass is specified. Containerd supports using different snapshotters if pods have the runtime handler annotation specified but this can cause some issues if a pod without an annotation is scheduled after a pod with a runtime handler is specified because kubelet will think the image is already present because it was fetched with a different snapshotter.
    * [mrunal] This intersects with Ensure image pull secrets. Another intersection with signature verification [https://github.com/kubernetes/kubernetes/pull/118652](https://github.com/kubernetes/kubernetes/pull/118652)
    * Wedson’s PR: [https://github.com/kubernetes/kubernetes/pull/118907](https://github.com/kubernetes/kubernetes/pull/118907) 
    * [Sergey] How rm works on containerd - does it remove both or just default?
    * [Peter] we can’t rely on CRI to do all of the handling because image pull policy isn’t propagated. thus, we do need the annotation approach for now until 1.29 planning when kubelet image gc undergoes redesign
* [mahamed/upodroid] Overhauling sig-node node e2e tests. I have been working with dims on introducing EC2 node e2e tests and I want to use this opportunity to complete [KEP-2464](https://github.com/kubernetes/enhancements/blob/master/keps/sig-testing/2464-kubetest2-ci-migration/README.md) and adopt kops' prowjob generator to generate jobs at scale as we need to test various permutations of multiple OS, architectures and CRI implementations.

    Implementation: [https://github.com/kubernetes/test-infra/pull/29944](https://github.com/kubernetes/test-infra/pull/29944) 


    PTAL at the e2e tests guidance in works: 

* 
* [fromani][discussion if time allows, otherwise PTAL and comment on github!] handling devices assignment on node reboot and kubelet restart: issue [https://github.com/kubernetes/kubernetes/issues/118559](https://github.com/kubernetes/kubernetes/issues/118559) and its proposed fix [https://github.com/kubernetes/kubernetes/pull/118635](https://github.com/kubernetes/kubernetes/pull/118635)  
* [haircommander] cgroup driver implementation discussion [https://github.com/kubernetes/kubernetes/pull/118770](https://github.com/kubernetes/kubernetes/pull/118770) 


## June 20th, 2023 [Cancelled]



* 


## June 13th, 2023

Recording: [https://www.youtube.com/watch?v=nF_3dnZJVnA](https://www.youtube.com/watch?v=nF_3dnZJVnA)

Enhancements tracking board: [https://github.com/orgs/kubernetes/projects/140/views/1?filterQuery=sig%3A%22sig-node%22&sortedBy%5Bdirection%5D=desc&sortedBy%5BcolumnId%5D=Status](https://github.com/orgs/kubernetes/projects/140/views/1?filterQuery=sig%3A%22sig-node%22&sortedBy%5Bdirection%5D=desc&sortedBy%5BcolumnId%5D=Status)



* [robscott/kl52752] [New Kubelet API to expose Pod readiness](https://groups.google.com/g/kubernetes-sig-node/c/aewFWsNP4dw)
    * To come back with a doc on the use cases and address concerns raised 
* [vinaykul] Merge a couple of house-keeping PRs to update in-place resize KEP
    * PR [k/e#4078](https://github.com/kubernetes/enhancements/pull/4078) updates current-milestone to v1.28 per [enhancement lead ask](https://github.com/kubernetes/enhancements/issues/1287#issuecomment-1585026101)
    * PR [k/e#3944](https://github.com/kubernetes/enhancements/pull/3944) updates KEP to reflect the last minute API changes we made in the actual implementation
* [klueska] [KEP Update: Promote Improved multi-numa alignment in Topology Manager](https://github.com/kubernetes/enhancements/pull/4079)
    * Changes needed to push this feature to beta
    * Implementation PR already under review
    *  as she approved the initial KEP
* [swsehgal] [REQUEST: Create kubernetes-sigs/noderesourcetopology-api](https://github.com/kubernetes/org/issues/4224)
    * We had a discussion about this and was approved on 

[May 23rd, 2023](#heading=h.5mc4dwbi1o5j)
    * Need formal approval from SIG Node Tech Leads on the issue
* [AkihiroSuda] Wants approvals from prod-readiness-approvers for `KEP-3857: Recursive Read-only (RRO) mounts` [https://github.com/kubernetes/enhancements/pull/3858](https://github.com/kubernetes/enhancements/pull/3858) 


## June 6th, 2023

Recording: [https://www.youtube.com/watch?v=rR3zOunp6FE](https://www.youtube.com/watch?v=rR3zOunp6FE)



* KEPs planning for 1.28:  
* [SergeyKanzhelev] Shared calendar for the SIG: [https://calendar.google.com/calendar/u/0?cid=YzY4ZGY0YTYxZDE0MTIyZThlODFlYjQyMzA5ZjZjY2E2M2ViMWI3YjQ0MzM4NGVlYmM4MDNlNjgzMmRiZTBiNkBncm91cC5jYWxlbmRhci5nb29nbGUuY29t](https://calendar.google.com/calendar/u/0?cid=YzY4ZGY0YTYxZDE0MTIyZThlODFlYjQyMzA5ZjZjY2E2M2ViMWI3YjQ0MzM4NGVlYmM4MDNlNjgzMmRiZTBiNkBncm91cC5jYWxlbmRhci5nb29nbGUuY29t) 
* [sohankunkerkar/haircommander] sandbox image pinning from CRI
    * [Sergey] I don't think you need a KEP for this. Looks like a simple codebase cleanup of a dead functionality.
    * Consensus is that no KEP is needed, as it’s cleaning up dead code from dockershim days.
* [haircommander/sohankunkerkar] Alternative/additional kubelet image GC schemes
    * There is interest here upstream for additional kubelet gc schemes
    * sig node is overloaded for 1.28, circle back in 1.29
* [SergeyKanzhelev] Highlight from the KEP: [https://github.com/kubernetes/enhancements/pull/3858/files](https://github.com/kubernetes/enhancements/pull/3858/files) proposes to introduce the new pattern for feature detection: “RuntimeHandlerFeatures”.
* ~~[klueska] [KEP: Add CDI devices to device plugin API](https://github.com/kubernetes/enhancements/pull/4011)~~
    * ~~Now ready for review~~
    * ~~please take a look~~
* [ffromani] question/docs: are the kubelet endpoints (like /pods) meant to be consumable by non-core components like 3rd party/external software? (xref: [https://github.com/kubernetes/kubernetes/issues/112119](https://github.com/kubernetes/kubernetes/issues/112119))
    * [Dawn] Initially wanted those API as internal-only with no guarantees
        * there are also security concerns
    * [Sergey] what guarantees we have for podresources?
        * [ffromani] need to have more guarantees
        * Syntax of API is guaranteed, but what happened on kubelet restart is not specifried
* [pacoxu/dims] [https://github.com/kubernetes/kubernetes/issues/118441](https://github.com/kubernetes/kubernetes/issues/118441) 
    * [Dawn] Need to create more guidance on how to troubleshoot infra issues


## May 30th, 2023

Recording: [https://www.youtube.com/watch?v=H9vnLgvTLvo](https://www.youtube.com/watch?v=H9vnLgvTLvo) 

Agenda

KEPs: [https://github.com/kubernetes/enhancements/issues?page=1&q=is%3Aissue+is%3Aopen+label%3Asig%2Fnode+milestone%3Av1.28](https://github.com/kubernetes/enhancements/issues?page=1&q=is%3Aissue+is%3Aopen+label%3Asig%2Fnode+milestone%3Av1.28) 



* [harche/mrunalp] Cautiously enabling swap only for Burstable Pods - [https://github.com/kubernetes/enhancements/pull/3957](https://github.com/kubernetes/enhancements/pull/3957) 
* [marquiz/haircommander]: [KEP 4033](https://github.com/kubernetes/enhancements/issues/4033): discover kubelet cgroup driver from CRI
    * There are other options that the CRI may want to tell the Kubelet what the state of the world is
    * focus this KEP on cgroup driver, but have API extendable so those other use cases (runtime class, QOS class, user namespace support) can be easily covered in the future
    * Separate CRI message from RuntimeStatus so Kubelet can request separately.
* [mimowo] Changed pod phase when containers exit with 0, related issue: [https://github.com/kubernetes/kubernetes/issues/118310](https://github.com/kubernetes/kubernetes/issues/118310). Summary:
    * eviction_manager, preemption: 1.26: Failed,          1.27: Succeeded
    * node shutdown                          1.26: Failed,         1.27: Succeeded
    * active deadline exceeded          1.26: Failed,         1.27: Failed
* [astoycos] [bpfd](https://bpfd.dev/) Presentation!
    * [Slides](https://docs.google.com/presentation/d/1QSSx8zX9I2VZJGDJcOCrU3sFYSji9Hv8F0kDJx-NEfo/edit?usp=sharing)
    * [SergeyKanzhelev] SiG node may help in terms of attributing events to pods metadata. When kernel events received - would be nice to know what Pod is running the process that sent this event. Please let us know if anything can be improved from SIG Node side to help with this.
* [byako] KEP-3542 CRI PullImageWithProgress [https://github.com/kubernetes/enhancements/pull/3547/files](https://github.com/kubernetes/enhancements/pull/3547/files)
* [adilGhaffarDev] What is the status of this fix: [https://github.com/kubernetes/kubernetes/pull/117030](https://github.com/kubernetes/kubernetes/pull/117030) what can we do to escalate it, if possible?
* [haircommander] [KEP 3983](https://github.com/kubernetes/enhancements/pull/4031): Add support for a drop-in kubelet configuration directory
    * Mostly a review request
* [SergeyKanzhelev] [https://github.com/kubernetes/kubernetes/pull/116429](https://github.com/kubernetes/kubernetes/pull/116429) sidecar PR.


## May 23rd, 2023

Recording: [https://www.youtube.com/watch?v=shmDtrq55V8](https://www.youtube.com/watch?v=shmDtrq55V8) 

Agenda



* [intunderflow] Following up from meeting on April 25th talking about lowering frequency of Startup / Readiness probe failure events, my preferred approach after digesting feedback, thoughts from the group about this approach? If happy I can put together a KEP
    * Always emit an event when the result of a probe changes (between Success and Failure, or Failure and Success)
    * When a startup probe fails or a readiness probe fails:
        * We emit the first failure
        * We then emit a failure every 1 hour if still failing
            * _Should this event be the same as the first failure, or should it be perhaps something like “Probe still failing since [first failure time]”_
    * No changes to liveness probes failing for now:
                * This will still cause mass event emission to hit the rate limit, but I want to tackle this incrementally and follow up on liveness probes
                * Lots of users watch for liveness probe failed events, so it's something to be particularly careful about in my opinion (people of course watch readiness/startup probes too, but I’d assume not as many / that liveness probes are the most populous probe type)
    * Thoughts from the group about this approach? If happy I can put together a KEP
* [intunderflow] [https://github.com/kubernetes/kubernetes/pull/115963](https://github.com/kubernetes/kubernetes/pull/115963) needs approver - I’d like to target this for 1.28 if no objections
* [ffromani] **REQUEST**: looking for approvers for (all items already part of 1.28 tracking document)
    * [https://github.com/kubernetes/enhancements/issues/2403](https://github.com/kubernetes/enhancements/issues/2403) (should be trivial)
    * [https://github.com/kubernetes/enhancements/issues/3545](https://github.com/kubernetes/enhancements/issues/3545) 
    * Could we please also review/approve [https://github.com/kubernetes/enhancements/pull/3980](https://github.com/kubernetes/enhancements/pull/3980) so we can merge [https://github.com/kubernetes/kubernetes/pull/116525](https://github.com/kubernetes/kubernetes/pull/116525)
* [swsehgal] Proposing NodeResourceTopology API under kubernetes-sigs: [https://github.com/kubernetes/org/issues/4224](https://github.com/kubernetes/org/issues/4224). Previously the API was proposed under staging but that proposal was [rejected](https://github.com/kubernetes/kubernetes/pull/96275#discussion_r1119165747) during API review.
    * +1: Alexander +1: Francesco
* [astoycos] Super Short Introduction of [https://github.com/bpfd-dev/bpfd](https://github.com/bpfd-dev/bpfd) (propose an actual 15-20 minute presentation for next week?) Also reach out in K8s slack [#bpfd](https://kubernetes.slack.com/archives/C04UJBW2553) and [#ebpf](https://kubernetes.slack.com/archives/C0562AYHZB3)


## May 16th, 2023

Recording: [https://www.youtube.com/watch?v=gnbV1nrXVZc](https://www.youtube.com/watch?v=gnbV1nrXVZc) 

Agenda:



* [everpeace] I opened a PR for KEP-3169.
    * PR: [https://github.com/kubernetes/kubernetes/pull/117842](https://github.com/kubernetes/kubernetes/pull/117842) 
    * KEP: [KEP-3619: Fine-grained SupplementalGroups control](https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/3619-supplemental-groups-policy) 
    * I would like the community to triage this and review it.
    * I’m very glad if someone would mentor me because it’s first time for me to make PR including API changes.
    * _NOTE: I’m sorry that I can’t show up to the community meeting due to timezone gap (2am in my timezone(Tokyo🇯🇵🗼)). I put this agenda to gain visibility and to help 1.28 planning._
* [tzneal] Discuss using the cgroup aware OOM killer [https://github.com/kubernetes/kubernetes/pull/117793](https://github.com/kubernetes/kubernetes/pull/117793) 
    * KEP needed for the API change?
    * Potential Options
        * No config, just a new default
        * Add API to Container to allow workload specific configuration
        * Add flag to kubelet
    * [Dawn] Let’s just change the default
    * [mrunal] OK with this.
    * Dawn will comment on the PR.
* [mimowo]
    * [https://github.com/kubernetes/kubernetes/pull/117586](https://github.com/kubernetes/kubernetes/pull/117586) needs review and approver
    * [https://github.com/kubernetes/kubernetes/issues/115688](https://github.com/kubernetes/kubernetes/issues/115688) needs discussion. As suggested by API reviewer ([https://github.com/kubernetes/kubernetes/issues/115688#issuecomment-1506909501](https://github.com/kubernetes/kubernetes/issues/115688#issuecomment-1506909501)) we should only add the condition when the Pod is actually killed. This means that refactoring of the handling of the activeDeadline handler is needed.  \
Questions:
* Can this refactoring be led by sig-node?
* Alternatively, can we go with the simple approach of adding the condition whenever when timeout is exceeded, as suggested in the POC PR: [https://github.com/kubernetes/kubernetes/pull/117973](https://github.com/kubernetes/kubernetes/pull/117973). Then, we could document that the behavior when the timeout is exceeded, but the containers aren’t killed (but terminate on their own) is subject to change. Proposed KEP updated for review: [https://github.com/kubernetes/enhancements/pull/3999](https://github.com/kubernetes/enhancements/pull/3999)
* [SergeyKanzhelev] Sidecar KEP: [https://github.com/kubernetes/enhancements/pull/3968/files](https://github.com/kubernetes/enhancements/pull/3968/files) and [https://github.com/kubernetes/kubernetes/pull/116429](https://github.com/kubernetes/kubernetes/pull/116429) 
* [mo] looking for a way to provide dynamic environment variables at runtime without persisting them in the Kube API (because the contents are sensitive)
    * would like to avoid any approach that uses admission to mutate pods
    * [Anish Ramasekar to Everyone (10:43 AM)] This is the subproject: [https://github.com/kubernetes-sigs/secrets-store-csi-driver](https://github.com/kubernetes-sigs/secrets-store-csi-driver) 
    * [Sergey] will this help: [https://github.com/kubernetes/enhancements/issues/3721](https://github.com/kubernetes/enhancements/issues/3721)?
    * Init container can download and then regular container will use those.
    * [mo] this ^^^ can work. Is this the right way?
    * [kevin] are you familiar with DRA? CDA is lowest level that makes abstract notion of a device available for a container. CDA can inject environment variables into the container. There may be a “device” that will perform all vault work and then will inject those variables to the container
    * [mo] what is the security model?
    * [kevin] this information will end up being statically stored at CDA file host system
    * [mo] is there way to observer this from kubernetes API?
    * [kevin] DRA is generalization of persistent volumes API. So it will provide some isolation.
    * [Sasha] this will not protect from exec into container. As no env variables would do.
    * [mo] can other containers see it? non-priviledged for example.
    * [mo] what is the interface for DRA? Can it be Daemonset in runtime?
    * [kevin] there is a talk about it at KubeCon. It has all the pieces to build this.
        * [Kevin] Here is my talk on how DRA drivers are structured:
        * 
* [klueska] New Feature for 1.28: Add CDI devices to device plugin API
    * Already added to the [planning doc](https://docs.google.com/document/d/1U10J0WwgWXkdYrqWGGvO8iH2HKeerQAlygnqgDgWv4E/edit#bookmark=id.c8sfex5lq85d)
    * Simple extension given that CDI devices have now been added to the CRI
    * does it make sense for you to be the approver?
    * [https://github.com/kubernetes/enhancements/issues/4009](https://github.com/kubernetes/enhancements/issues/4009)

—- MOVED from 5/2/2023. Move above this line if you plan to show up on the meeting —



* [kannon92] PRs need approval
    * [https://github.com/kubernetes/kubernetes/pull/116231](https://github.com/kubernetes/kubernetes/pull/116231) - Cleanup around image parsing, test fixes, and more coverage.  Has LGTM but needs a final approval
    * [https://github.com/kubernetes/kubernetes/pull/117702](https://github.com/kubernetes/kubernetes/pull/117702) - Rename PodHasNetwork to PodReadyToStartContainersCondition Code PR.  Has LGTM but needs approval


## May 9th, 2023

Recording: [https://www.youtube.com/watch?v=18cRhXTf0Cc](https://www.youtube.com/watch?v=18cRhXTf0Cc) 

Total active pull requests: [242](https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+is%3Aopen+)


<table>
  <tr>
   <td><strong>Incoming</strong>
   </td>
   <td>
   </td>
   <td><strong>Completed</strong>
   </td>
   <td>
   </td>
  </tr>
  <tr>
   <td>Created:
   </td>
   <td><strong><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++created%3A2023-05-02T17%3A00%3A00%2B0000..2023-05-09T16%3A59%3A06%2B0000">19</a></strong>
   </td>
   <td>Closed:
   </td>
   <td><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++is%3Aunmerged+closed%3A2023-05-02T17%3A00%3A00%2B0000..2023-05-09T16%3A59%3A06%2B0000">9</a>
   </td>
  </tr>
  <tr>
   <td>Updated:
   </td>
   <td><strong><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+updated%3A2023-05-02T17%3A00%3A00%2B0000..2023-05-09T16%3A59%3A06%2B0000+created%3A%3C2023-05-02T17%3A00%3A00%2B0000">118</a></strong>
   </td>
   <td>Merged:
   </td>
   <td><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++merged%3A2023-05-02T17%3A00%3A00%2B0000..2023-05-09T16%3A59%3A06%2B0000">17</a>
   </td>
  </tr>
</table>




* [swsehgal] Community discussion on device Manager recovery bugfix backport
    * [Issue](https://github.com/kubernetes/kubernetes/issues/109595), [Fix](https://github.com/kubernetes/kubernetes/pull/116376)
    * Device Manager is a GA feature so proposing to backport the fix to [1.25](https://github.com/kubernetes/kubernetes/pull/117738), [1.26](https://github.com/kubernetes/kubernetes/pull/116337) and [1.27](https://github.com/kubernetes/kubernetes/pull/117719)
* [karthik-k-n] Community thoughts on Dynamic Node resize proposal
    * [Discussion document](https://docs.google.com/document/d/1KfjPRmCc8Xk0xxa4S8ZRle6VMzc1C6MQg4ivOaoB150/edit#heading=h.xgjl2srtytjt)
    * [Enhancement](https://github.com/kubernetes/enhancements/issues/3953)
    * [POC code](https://github.com/kubernetes/kubernetes/pull/115755)
* [clayton] Discussion of kubelet state improvements for 1.28 - trying to identify which areas to focus on
    *  - doc crafted at end of 1.27
    * Right now was planning on working with in-place resource (to reduce complexity encountered during alpha and before beta), and also to help address the kubelet subcomponents using the wrong pod_manager [https://github.com/kubernetes/kubernetes/pull/117371](https://github.com/kubernetes/kubernetes/pull/117371) (ryan reviewing) to unblock harder problems static pods are encountering
    * Other interest from contributors?
        * kubelet plugins potentially [https://github.com/kubernetes/enhancements/pull/3853](https://github.com/kubernetes/enhancements/pull/3853) 
        * sidecar enh: [https://github.com/kubernetes/enhancements/issues/753](https://github.com/kubernetes/enhancements/issues/753) 
        * Github project tracking ongoing pod lifecycle issues: [https://github.com/orgs/kubernetes/projects/133/](https://github.com/orgs/kubernetes/projects/133/) 
    * Improvements in testing from KEPs
* [zmerlynn] Discuss 
    * Dawn: Maybe first restart free, don’t punish
    * Clayton: DaemonSet that runs effectively a for loop to anneal policy
        * There are things we don’t account for, like system resources in a crash looping pod - what does it actually cost to restart a container
    * Derek (on chat): I wonder if we need a way to measure a qps generally for the behavior that crashloopbackoff is trying to protect
        * systemd gives StartLimitBurst and then when that is exhausted you go to StartLimitInterval.... feels like we could give a burst
    * Sergey: Maybe we also need “it’s a bad failure, reschedule me”
    * David: Is it up to the admin to define this?
    * Kevin: KEP in question that Sergei mentioned: [https://github.com/kubernetes/enhancements/pull/3816](https://github.com/kubernetes/enhancements/pull/3816)
    * Clayton: Full backoff doesn’t make sense for static pod anyways

—-- End of the meeting. MOVED TO THE NEXT WEEK —--


## May 2nd, 2023

Recording: ​​[https://www.youtube.com/watch?v=whN6nPOp62g](https://www.youtube.com/watch?v=whN6nPOp62g) 

Total active pull requests: [241](https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+is%3Aopen+)


<table>
  <tr>
   <td><strong>Incoming</strong>
   </td>
   <td>
   </td>
   <td><strong>Completed</strong>
   </td>
   <td>
   </td>
  </tr>
  <tr>
   <td>Created:
   </td>
   <td><strong><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++created%3A2023-04-25T17%3A00%3A00%2B0000..2023-05-02T17%3A02%3A45%2B0000">27</a></strong>
   </td>
   <td>Closed:
   </td>
   <td><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++is%3Aunmerged+closed%3A2023-04-25T17%3A00%3A00%2B0000..2023-05-02T17%3A02%3A45%2B0000">11</a>
   </td>
  </tr>
  <tr>
   <td>Updated:
   </td>
   <td><strong><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+updated%3A2023-04-25T17%3A00%3A00%2B0000..2023-05-02T17%3A02%3A45%2B0000+created%3A%3C2023-04-25T17%3A00%3A00%2B0000">88</a></strong>
   </td>
   <td>Merged:
   </td>
   <td><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++merged%3A2023-04-25T17%3A00%3A00%2B0000..2023-05-02T17%3A02%3A45%2B0000">25</a>
   </td>
  </tr>
</table>




* [SergeyKanzhelev, mrunalp] retro and 1.28 planning
    * [Ruiwen] 1.27 [retro](https://docs.google.com/document/d/1DxJH1w_lrEOfflR-TED1vjc0ZYIXO0aDb3vJXPMKCdY/edit#)
    * [mrunal] planning: [SIG Node - 1.28 Planning](https://docs.google.com/document/d/1U10J0WwgWXkdYrqWGGvO8iH2HKeerQAlygnqgDgWv4E/edit#bookmark=id.fp500pn2fk1f) 
* [Seaiii] [https://github.com/kubernetes/kubernetes/pull/113883](https://github.com/kubernetes/kubernetes/pull/113883) The second time the pod deleted the grace period does not take effect .Please review update PR 


## Apr 25th, 2023

Recording: [https://www.youtube.com/watch?v=oQi3gPsODV0](https://www.youtube.com/watch?v=oQi3gPsODV0)



* [intunderflow] [https://github.com/kubernetes/kubernetes/pull/115963](https://github.com/kubernetes/kubernetes/pull/115963) needs approver
* [intunderflow] Thoughts on startup probe / readiness probe event emission behavior?
    * Currently the readiness probes and startup probes emit ContainerUnhealthy events each time they probe the container and it is Unhealthy.
    * For liveness probes a container going from a healthy state to suddenly unhealthy is important and notable, but for Readiness and Startup probes it's pretty typical for a container to be unhealthy since the point of these probes is to wait until the container is healthy.
    * Emitting these events eats into the rate limit of 25 events per object sent to the API server.
    * Readiness probes and Startup probes failing multiple times is pretty typical of their operation, since their point is to gate the container until it succeeds.
    * It would be nice if Readiness probes and Startup probes didn’t eat events as fast as they did.
    * My thoughts and opinions:
        * We could consider changing the startup and readiness probe to only emit when they probe the container and it is healthy (since that leads to a change in state and action being taken)
        * My PR above (if approved) would still then report if a startup probe or readiness probe fail conclusively against a container
    * [Action Item] Count incrementation on Events? Why not working for failing probes?
    * [Ryan] The event recorder has a max retries of 12
    * [https://github.com/kubernetes/client-go/blob/master/tools/record/event.go#L38](https://github.com/kubernetes/client-go/blob/master/tools/record/event.go#L38)
    * [Todd] we need events to be re-emitted periodically. Do not discard them universally. Less frequency, but definitely more events we want to know about like flakes of readiness probe.
    * 
* [SergeyKanzhelev] Probes functionality cleanup: [https://docs.google.com/document/d/1G5nGH97s3UTANbA5IyQ7nVIHnrLKfgVZssSYnvp_qX4/edit](https://docs.google.com/document/d/1G5nGH97s3UTANbA5IyQ7nVIHnrLKfgVZssSYnvp_qX4/edit)  
* [haircommander/Peter] Kubelet drop-in config support
    * After conversation about dropping cli flag support, it was illuminated that our users (downstream in Openshift) rely on this feature. Could be a good time to introduce drop-in file support like in `/etc/kubernetes/kubelet.conf.d`
    * Peter to make proposal to SIG-Arch to see if other components would like to adopt a similar pattern, as well as open an issue to have an asynchronous conversation.
* [SergeyKanzhelev] ​​[https://github.com/kubernetes/kubernetes/pull/116429](https://github.com/kubernetes/kubernetes/pull/116429) , uber issue: [https://github.com/kubernetes/kubernetes/issues/115934](https://github.com/kubernetes/kubernetes/issues/115934) 


## Apr 18th, 2023

No call (kubecon)


## Apr 11th, 2023

Recording: [https://www.youtube.com/watch?v=R9bml9YmP3k](https://www.youtube.com/watch?v=R9bml9YmP3k)



* [klueska] Need approval on PR to update DRA KEP with changes merged into v1.27
    * [https://github.com/kubernetes/enhancements/pull/3802](https://github.com/kubernetes/enhancements/pull/3802)
* [liggitt/derek] proposal to support node/control-plane skew of n-3 ([KEP-3935](https://github.com/kubernetes/enhancements/issues/3935), [draft proposal](https://github.com/kubernetes/enhancements/pull/3936))
    * What in-progress node feature / cleanup rollouts rely on n-2 skew?
        * might delay default-on of in-place-resize for one release (~~AI: jordan / vinay sync up~~); notes from jordan/vinay 2023-05-03:
            * a 1.27+ node with the feature disabled will not modify resources as requested, will mark pods requesting resize as "infeasible"
            * a pre-1.27 node will not modify resources as requested, with no user feedback
            * after 1.27 work, we realized that kubelet perpetually reports pod resize as InProgress when running against a containerd that supports UpdateContainerResources CRI API (containerd ~1.4/~1.5 era) but does not support ContainerStatus CRI API (added to CRI API in [k8s 1.25](https://github.com/kubernetes/kubernetes/pull/111645), supported in containerd 1.6.9+), so there's already user feedback improvements to make and possibly delay beta for
            * _if_ we were ready to promote in-place-resize to beta in 1.29, n-3 skew would mean 1.26 kubelets would not give any user feedback about lack of support for the feature, but would otherwise fail safe
    * derek:
        * include alternative considered of supporting in-place minor upgrades, rationale why that approach wasn't chosen
            * OS upgrades, immutable nodes can't use in-place for minor upgrades
            * cost of supporting/testing in-place minor upgrades is significantly higher, impacts development of new features and evolution of existing features
        * make sure it is clear what guidance should be given to people working on new features for what to do for features older kubelets don't support yet
* [mweston & atanas] Still working on the [https://github.com/obiTrinobihttps://github.com/obiTrinobiIntel/enhancements/tree/atanas/cci-updated/keps/sig-node/3675-resource-plugin-managerIntel/enhancements/tree/atanas/cci-updated/keps/sig-node/3675-resource-plugin-manager](https://github.com/obiTrinobiIntel/enhancements/tree/atanas/cci-updated/keps/sig-node/3675-resource-plugin-manager) KEP.  Need help with scheduling re Dawn or other member in getting feedback.
* [mrunal] Canceling next week's meeting for kubecon.


## Apr 4th, 2023

Recording: [https://www.youtube.com/watch?v=Y_TWnklb0vI](https://www.youtube.com/watch?v=Y_TWnklb0vI)



* [pacoxu] [undeprecate kubelet --provider-id flag](https://github.com/kubernetes/kubernetes/pull/116530#top): what are your plans around graduating kubelet config file/actually deprecating these flags in the future?
* 
* [iancoolidge] Follow-up on issue [https://github.com/kubernetes/kubernetes/issues/115994](https://github.com/kubernetes/kubernetes/issues/115994)
    * discuss specifying –reserved-cpus and also –exclusive-cpus or something like that (see [https://github.com/kubernetes/kubernetes/issues/115994#issuecomment-1495980750](https://github.com/kubernetes/kubernetes/issues/115994#issuecomment-1495980750))
        * fromani: adding what seems another cpu pool is something already emerged from different conversations and would probably deserve a separate conversation (e.g. not a simple bugfix)
    * Please also look at KEP: [https://github.com/obiTrinobiIntel/enhancements/tree/atanas/cci-updated/keps/sig-node/3675-resource-plugin-manager](https://github.com/obiTrinobiIntel/enhancements/tree/atanas/cci-updated/keps/sig-node/3675-resource-plugin-manager) 
* [rata] Userns KEP 127: add support for stateful pods
    * We don’t need code changes in the kubelet for this (just change the validation)
    * Therefore, we want to just change the scope of the KEP to support stateful pods too
    * We want to deprecate the feature gate “UserNamespacesStatelessPodsSupport” and add “UserNamespacesSupport”
    * This new feature gate will activate userns for all pods (stateful/stateless)
    * If this sounds good, we will do a PoC and propose the KEP changes widening the scope and explaining how the stateful case works too.
        * [mrunal] This may be okay but let’s open a KEP change and get opinions of other reviewers involved.
        * [mrunal] We need to start thinking about how user namespaces will work with pod security policies.
        * [rata]: Mrunal and I will join sig-auth to start the PSS conversation
        * [rata] Maybe they need fields to be GA? But happy to start discussing.


## Mar 28st, 2023

Recording: [https://www.youtube.com/watch?v=yb_LtE0hGDc](https://www.youtube.com/watch?v=yb_LtE0hGDc) 



* [SergeyKanzhelev] Annual Report: [https://github.com/kubernetes/community/pull/7220](https://github.com/kubernetes/community/pull/7220) 

    Let’s edit together: [https://docs.google.com/document/d/17Z3LO3pSdv9R-v9yLIMO5a46nwXRQTsaEDg0iN74rhs/edit?usp=sharing](https://docs.google.com/document/d/17Z3LO3pSdv9R-v9yLIMO5a46nwXRQTsaEDg0iN74rhs/edit?usp=sharing) 

* [jlpedrosa]
    * memory.oom.group setting to oom the whole cgroup in the container. \
[slack convo](https://kubernetes.slack.com/archives/C0BP8PW9G/p1679649131585659).
        * [Mrunal] container level makes sense
        * [Sergey] for sidecars we will adjust oom score for sidecars so it’s almost the “whole Pod” being killed
        * [Mrunal] we can start with the issue, may not need a KEP for this
        * [Todd Neal] I think there is a potential for API surface as the new behavior may not be desired in all cases.  haproxy was the example brought up in Slack where it may handle OOM correctly on a single process.  Most everything else probably doesn't, so you might want a default of turning oom.group on and allowing containers to opt-out.


## Mar 21st, 2023

Recording: [https://www.youtube.com/watch?v=IjxUleYcKgk](https://www.youtube.com/watch?v=IjxUleYcKgk) 



* [iholder101/harche] - Graduating Support for Swap to Beta
    * [https://github.com/kubernetes/enhancements/blob/master/keps/sig-node/2400-node-swap/README.md#beta](https://github.com/kubernetes/enhancements/blob/master/keps/sig-node/2400-node-swap/README.md#beta)
    * PoC for setting swap limit in cgroup v2 -[ https://github.com/kubernetes/kubernetes/pull/116637](https://github.com/kubernetes/kubernetes/pull/116637)
* [adilGhaffarDev] - kubectl drain improvements
    * I would like to add some improvements in kubectl drain. Kindly check this issue: [https://github.com/kubernetes/kubernetes/issues/116210](https://github.com/kubernetes/kubernetes/issues/116210) 
* [mrunal] Cgroups v1 deprecation
* 


## Mar 14th, 2023

Recording: [https://www.youtube.com/watch?v=e0DA7x4zTs0](https://www.youtube.com/watch?v=e0DA7x4zTs0)

Total: [200](https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+is%3Aopen+)


<table>
  <tr>
   <td><strong>Incoming</strong>
   </td>
   <td>
   </td>
   <td><strong>Completed</strong>
   </td>
   <td>
   </td>
  </tr>
  <tr>
   <td>Created:
   </td>
   <td><strong><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++created%3A2023-02-28T17%3A00%3A00%2B0000..2023-03-14T16%3A56%3A55%2B0000">86</a></strong>
   </td>
   <td>Closed:
   </td>
   <td><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++is%3Aunmerged+closed%3A2023-02-28T17%3A00%3A00%2B0000..2023-03-14T16%3A56%3A55%2B0000">35</a>
   </td>
  </tr>
  <tr>
   <td>Updated:
   </td>
   <td><strong><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+updated%3A2023-02-28T17%3A00%3A00%2B0000..2023-03-14T16%3A56%3A55%2B0000+created%3A%3C2023-02-28T17%3A00%3A00%2B0000">203</a></strong>
   </td>
   <td>Merged:
   </td>
   <td><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++merged%3A2023-02-28T17%3A00%3A00%2B0000..2023-03-14T16%3A56%3A55%2B0000">103</a>
   </td>
  </tr>
</table>


Needs approval: label:lgtm -label:approved 41

[https://github.com/kubernetes/kubernetes/issues?q=is%3Aissue+is%3Aopen+label%3Apriority%2Fcritical-urgent++label%3Asig%2Fnode+](https://github.com/kubernetes/kubernetes/issues?q=is%3Aissue+is%3Aopen+label%3Apriority%2Fcritical-urgent++label%3Asig%2Fnode+)

[https://github.com/kubernetes/kubernetes/issues?q=is%3Aissue+is%3Aopen+label%3Apriority%2Fimportant-soon++label%3Asig%2Fnode+](https://github.com/kubernetes/kubernetes/issues?q=is%3Aissue+is%3Aopen+label%3Apriority%2Fimportant-soon++label%3Asig%2Fnode+) 



* [SergeyKanzhelev] Release blocking: 
    * [https://github.com/kubernetes/kubernetes/issues/116549](https://github.com/kubernetes/kubernetes/issues/116549) [Fixed]
    * [https://github.com/kubernetes/kubernetes/issues/116262](https://github.com/kubernetes/kubernetes/issues/116262)

    Standalone kubelet mode is de facto GA feature, even though not well documented. No GA feature should be broken by feature gate enablement.


    Options: introduce standalone mode tests like so: 

* [https://github.com/kubernetes/kubernetes/pull/116551](https://github.com/kubernetes/kubernetes/pull/116551)
* revert changes caused the failure
* fix changes caused the failure before the release

	InPlace update TODOs:



* api change
* panic in standalone mode
* Try moving forward with the feature.
* [SergeyKanzhelev] Sidecar KEP. [https://github.com/kubernetes/kubernetes/pull/116429](https://github.com/kubernetes/kubernetes/pull/116429) 
* [pacoxu]<span style="text-decoration:underline;"> [add net.ipv4.ip_local_reserved_ports to safe sysctls](https://github.com/kubernetes/kubernetes/pull/115374#top)#115374</span>, the sysctl is changed to be namespaced since kernel 3.16, and the PR will add it to safe syctl only if the kernel version is 3.16+.
    * [https://github.com/kubernetes/system-validators/blob/a0cb0d12f4d8ed79fa4ee4725ae179528dd2d522/validators/kernel_validator.go](https://github.com/kubernetes/system-validators/blob/a0cb0d12f4d8ed79fa4ee4725ae179528dd2d522/validators/kernel_validator.go)
    * [Paco] I have opened an issue ([https://github.com/kubernetes/kubernetes/issues/116799](https://github.com/kubernetes/kubernetes/issues/116799)) to gather feedback and track whether we should modify the minimum kernel version.
* ~~[klueska] Need approval for feature gates on this PR:~~
    * ~~I cannot attend the meeting unfortunately~~
    * ~~The PR is close, and I am confident we can get all the comments addressed before 5pm, but I need help getting the feature gates approved.~~
    * ~~[https://github.com/kubernetes/kubernetes/pull/115847](https://github.com/kubernetes/kubernetes/pull/115847)~~
* [aravindh/sig-windows] Next steps for merging [KEP 2258: add node log query](https://github.com/kubernetes/kubernetes/pull/96120#top)
    * Discussed previously on 

[Jan 31st, 2023](#heading=h.91rl4flxpf59)
* [marosset/sig-windows] [https://github.com/kubernetes/kubernetes/pull/116546](https://github.com/kubernetes/kubernetes/pull/116546)
    * updating perfCounterUpdatePeriod in kubelet to 10 seconds on Windows to address some perf issues when running logs of pods


## Mar 7th, 2023

Recording: [https://www.youtube.com/watch?v=KgAR613c1Bs](https://www.youtube.com/watch?v=KgAR613c1Bs) 

Total PRs: [241](https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+is%3Aopen+)


<table>
  <tr>
   <td><strong>Incoming</strong>
   </td>
   <td>
   </td>
   <td><strong>Completed</strong>
   </td>
   <td>
   </td>
  </tr>
  <tr>
   <td>Created:
   </td>
   <td><strong><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++created%3A2023-02-28T17%3A00%3A00%2B0000..2023-03-07T17%3A58%3A18%2B0000">35</a></strong>
   </td>
   <td>Closed:
   </td>
   <td><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++is%3Aunmerged+closed%3A2023-02-28T17%3A00%3A00%2B0000..2023-03-07T17%3A58%3A18%2B0000">14</a>
   </td>
  </tr>
  <tr>
   <td>Updated:
   </td>
   <td><strong><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+updated%3A2023-02-28T17%3A00%3A00%2B0000..2023-03-07T17%3A58%3A18%2B0000+created%3A%3C2023-02-28T17%3A00%3A00%2B0000">136</a></strong>
   </td>
   <td>Merged:
   </td>
   <td><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++merged%3A2023-02-28T17%3A00%3A00%2B0000..2023-03-07T17%3A58%3A18%2B0000">31</a>
   </td>
  </tr>
</table>



    [https://github.com/kubernetes/kubernetes/issues?q=is%3Aissue+is%3Aopen+label%3Apriority%2Fimportant-soon+label%3Asig%2Fnode](https://github.com/kubernetes/kubernetes/issues?q=is%3Aissue+is%3Aopen+label%3Apriority%2Fimportant-soon+label%3Asig%2Fnode)



* [SergeyKanzhelev] 19 enhancements tracked, and at the moment 0 were opted-in for Feature Blogs.
* [KevinHannon@kannon92] Starting work on [https://github.com/kubernetes/enhancements/pull/3816](https://github.com/kubernetes/enhancements/pull/3816) (Pending Pods stuck due to configuration errors) \
- Created a POC PR to see about validation of some of these configuration errors \
- [https://github.com/kubernetes/kubernetes/pull/115736](https://github.com/kubernetes/kubernetes/pull/115736) \
- Should I consider moving this into a KEP of its own?
* [Francesco @ffromani] setting rate limits for kubelet endpoints (required for KEP 606 GA). Items emerged during the[ API review](https://github.com/kubernetes/kubernetes/pull/115852#pullrequestreview-1326552711):
    * kubelet has a lot of endpoints, we are adding rate limiting only for the podresources. [Is this coherent with the overall direction of kubelet serving](https://github.com/kubernetes/kubernetes/pull/115852#discussion_r1126611858)?
    * [The rate limit is global. Clients can disrupt each other](https://github.com/kubernetes/kubernetes/pull/115852#discussion_r1126605828). Is this acceptable?
        * vs per-connection rate-limit. But then would we need to limit also the max connections?
    * [Is the QPS/Burst setting good enough for kubelet](https://github.com/kubernetes/kubernetes/pull/115852#discussion_r1126606012)?
        * vs priority-and-fairness like API server
* [SergeyKanzhelev] [https://github.com/kubernetes/kubernetes/pull/116121](https://github.com/kubernetes/kubernetes/pull/116121) any concerns with increasing qps limits? 
    * This really helps, and the first throttling issue is volume mount timeout errors.
* [mimowo] Looking for reviewers / approvers for:
    * Mark Deleted Pending Pods as Failed by Kubelet [https://github.com/kubernetes/kubernetes/pull/115331](https://github.com/kubernetes/kubernetes/pull/115331)
    * Move Deleted Running pods to terminal phase by Kubelet [https://github.com/kubernetes/kubernetes/pull/116324](https://github.com/kubernetes/kubernetes/pull/116324)
* [vinaykul] [InPlace Pod Vertical Scaling](https://github.com/kubernetes/kubernetes/pull/102884/) - status update
    * Two bugs identified so far from this PR:
        * [#116262](https://github.com/kubernetes/kubernetes/issues/116262) panic in standalone kubelet in GKE CI (release blocker)
            * [Fix](https://github.com/kubernetes/kubernetes/pull/116271) is in review
        * [#116175](https://github.com/kubernetes/kubernetes/issues/116175) new test failure in CI job
            * Initial investigation shows in-place resize didn’t regress this job
* [mahamed/@upodroid & todd/@tzneal] Kubelet/node e2e tests on AWS:
    * [https://github.com/kubernetes/test-infra/issues/28899](https://github.com/kubernetes/test-infra/issues/28899) Top Level Issue
    * Todd’s PR to add the runner logic in k/k [https://github.com/kubernetes/kubernetes/pull/116236](https://github.com/kubernetes/kubernetes/pull/116236)
    * My PR to add vanilla ubuntu2204 public GCE images to tests. [https://github.com/kubernetes/test-infra/pull/28856](https://github.com/kubernetes/test-infra/pull/28856) 
    * e2e ubuntu tests are passing on EC2 instances: [https://github.com/kubernetes/kubernetes/issues/116114](https://github.com/kubernetes/kubernetes/issues/116114) 
    * AWS accounts for node e2e tests have been created, I need to configure the prow jobs to run tests on there.
    * AI for Mahamed to take back to sig-k8s-infra
        * How can members of sig-node access the AWS accounts and view failing nodes, etc?
        * Someone from AWS to help with this
    * [Dawn] @mahamed, here is recent discuss on CI failures on Fedora:[ https://docs.google.com/document/d/1Ne57gvidMEWXR70OxxnRkYquAoMpt56o75oZtg-OeBg/edit#bookmark=id.rglbf1gnhrpp](https://docs.google.com/document/d/1Ne57gvidMEWXR70OxxnRkYquAoMpt56o75oZtg-OeBg/edit#bookmark=id.rglbf1gnhrpp)
* [atanas] still working on [https://github.com/kubernetes/enhancements/pull/3853](https://github.com/kubernetes/enhancements/pull/3853) and will have further updates soon.  Next meeting is March 7th, directly after this one, zoom here [https://us02web.zoom.us/j/86511166765](https://us02web.zoom.us/j/86511166765) (need to also update Dawn-schedules don’t work correctly so need 2 meetings, unfortunately)
    * Attribute-based API update 
    * Architecture Sync - move towards a single CCI driver implementation which can understand DRA claims 
    * Demo covering possible DRA claim-based approach combined with CCI , boot strap the system , run std pods, demonstrate driver failure scenario , allocate shared and exclusive resources. 
* [SergeyKanzhelev] any special protections needed for [https://github.com/kubernetes/kubernetes/blob/master/pkg/kubelet/events/event.go](https://github.com/kubernetes/kubernetes/blob/master/pkg/kubelet/events/event.go) (see [conversation](https://kubernetes.slack.com/archives/CJUQN3E4T/p1677664222256419))?
* [Philip Laine] Read only CRI API access
    * [https://github.com/containerd/containerd/issues/8085](https://github.com/containerd/containerd/issues/8085)
    * [https://github.com/cri-o/cri-o/issues/6667](https://github.com/cri-o/cri-o/issues/6667)
        * Mike Brown(brownwm/mikebrow): suggest opening a google doc link it here.. invite mrunalp, peter hunt, sasha, samuel, mike brown, others that have an interest in shaping the cri into namespaced(or rbac)/access patterns for some cases like monitoring 


## Feb 28st, 2023

Recording: [https://www.youtube.com/watch?v=IHcI6Jwo5PQ](https://www.youtube.com/watch?v=IHcI6Jwo5PQ) 

Total PRs: [248](https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+is%3Aopen+)


<table>
  <tr>
   <td><strong>Incoming</strong>
   </td>
   <td>
   </td>
   <td><strong>Completed</strong>
   </td>
   <td>
   </td>
  </tr>
  <tr>
   <td>Created:
   </td>
   <td><strong><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++created%3A2023-02-21T17%3A00%3A00%2B0000..2023-02-28T17%3A59%3A00%2B0000">36</a></strong>
   </td>
   <td>Closed:
   </td>
   <td><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++is%3Aunmerged+closed%3A2023-02-21T17%3A00%3A00%2B0000..2023-02-28T17%3A59%3A00%2B0000">9</a>
   </td>
  </tr>
  <tr>
   <td>Updated:
   </td>
   <td><strong><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+updated%3A2023-02-21T17%3A00%3A00%2B0000..2023-02-28T17%3A59%3A00%2B0000+created%3A%3C2023-02-21T17%3A00%3A00%2B0000">83</a></strong>
   </td>
   <td>Merged:
   </td>
   <td><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++merged%3A2023-02-21T17%3A00%3A00%2B0000..2023-02-28T17%3A59%3A00%2B0000">11</a>
   </td>
  </tr>
</table>




* **[01:00 UTC Wednesday 15th March 2023 / 17:00 PDT Tuesday 14th March 2023](https://everytimezone.com/s/c4971746)**: Week 10 — [Code Freeze](https://github.com/kubernetes/sig-release/blob/master/releases/release_phases.md#code-freeze)
* [vinaykul] [InPlace Pod Vertical Scaling](https://github.com/kubernetes/kubernetes/pull/102884/) PR - status update
    * Merged 🙂
    * ref WIP PR: [https://github.com/kubernetes/kubernetes/pull/116119](https://github.com/kubernetes/kubernetes/pull/116119)
* [SergeyKanzhelev] Sidecar: [https://github.com/kubernetes/kubernetes/issues/115934](https://github.com/kubernetes/kubernetes/issues/115934)  
* [mimowo]
    * need reviewer/approver for: [https://github.com/kubernetes/kubernetes/pull/116082](https://github.com/kubernetes/kubernetes/pull/116082). Is it an issue upstream in containerd? [mike brown] let’s chat I don’t think this is a problem with containerd.. more an issue of expectation of the start request being serial and responding before the async start kicking off.. iow add a sleep/yield to the test itself before it “ooms” and you will/should get the started response flowing back through kubelet before the oom happens.. Changing to an ack model for the start request before actually starting would be in conflict with the start being able to return certain errors.
    * Discuss implementation decisions for [https://github.com/kubernetes/kubernetes/pull/115331](https://github.com/kubernetes/kubernetes/pull/115331). Specific questions:
        * Should we restrict the handling to pods with finalizers (to save QPS)?
        * When Kubelet restarts there is a short time window that the phase may flip back to Pending, is this something specific to this scenario, or a general behavior / bug in Kubelet?
        * Should we also make sure that all Running pods with deletionTimestamp end up in terminal phase? This is currently not the case for pods with RestartPolicy=OnFailure or Always.
    * Followup:
        * Clayton’s PR that may fix the failure case: [https://github.com/kubernetes/kubernetes/pull/113145](https://github.com/kubernetes/kubernetes/pull/113145)
        * E2E added to Clayton’s PR would be helpful to see if the issue is fixed or not


## Feb 21st, 2023

Recording: [https://www.youtube.com/watch?v=Hod1MGk99lc](https://www.youtube.com/watch?v=Hod1MGk99lc) 

Total PRs: [230](https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+is%3Aopen+) 

From Jan 24th:


<table>
  <tr>
   <td><strong>Incoming</strong>
   </td>
   <td>
   </td>
   <td><strong>Completed</strong>
   </td>
   <td>
   </td>
  </tr>
  <tr>
   <td>Created:
   </td>
   <td><strong><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++created%3A2023-01-24T17%3A00%3A00%2B0000..2023-02-21T17%3A53%3A43%2B0000">129</a></strong>
   </td>
   <td>Closed:
   </td>
   <td><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++is%3Aunmerged+closed%3A2023-01-24T17%3A00%3A00%2B0000..2023-02-21T17%3A53%3A43%2B0000">48</a>
   </td>
  </tr>
  <tr>
   <td>Updated:
   </td>
   <td><strong><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+updated%3A2023-01-24T17%3A00%3A00%2B0000..2023-02-21T17%3A53%3A43%2B0000+created%3A%3C2023-01-24T17%3A00%3A00%2B0000">177</a></strong>
   </td>
   <td>Merged:
   </td>
   <td><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++merged%3A2023-01-24T17%3A00%3A00%2B0000..2023-02-21T17%3A53%3A43%2B0000">76</a>
   </td>
  </tr>
</table>




* [lucysweet] [Give an indication in container events for probe failure as to whether the failure was ignored due to FailureThreshold](https://github.com/kubernetes/kubernetes/issues/115823)
    * [SergeyKanzhelev] replied on the issue
* [iancoolidge] Clarify –reserved-cpus behavior with static CPU manager, but only a subset of pods are static
    * [swsehgal] This kubelet flag is to specify explicitly the list of CPUs for kubernetes processes or OS processes. These CPUs are removed from the default pool that CPU Manager in kubelet uses for allocation to pods. 
    * [dawnchen] You can view –reserved-cpus as the interface for non-pod tasks running on the node including kubelet, container runtime, and kernel threads. It is a complement for the static CPU managers. 
    * [iancoolidge] consensus from discussion is that no workloads should get scheduled on the –reserved-cpus, this seems to be happening on our test case though.
        * iancoolidge to create issue on github (thanks all for discussion!)
        * [https://github.com/kubernetes/kubernetes/issues/115994](https://github.com/kubernetes/kubernetes/issues/115994)
* [mimowo] Looking for approval from sig-node for [https://github.com/kubernetes/kubernetes/pull/113205](https://github.com/kubernetes/kubernetes/pull/113205). Then, for: [https://github.com/kubernetes/kubernetes/pull/112977](https://github.com/kubernetes/kubernetes/pull/112977).
    * CRITEST: [https://github.com/kubernetes/community/blob/master/contributors/devel/sig-node/cri-validation.md?plain=1](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-node/cri-validation.md?plain=1) 
    * [https://github.com/kubernetes-sigs/cri-tools/issues/1102](https://github.com/kubernetes-sigs/cri-tools/issues/1102) 
* [atanas] still working on [https://github.com/kubernetes/enhancements/pull/3853](https://github.com/kubernetes/enhancements/pull/3853) and will have further updates soon.  Next meeting is March 7th, 6:00am PST zoom here:  [https://us02web.zoom.us/j/82567156922?pwd=Q2xscE0rRjluRTlvdk5FK3hzUFpDQT09](https://us02web.zoom.us/j/82567156922?pwd=Q2xscE0rRjluRTlvdk5FK3hzUFpDQT09)  Plan for next meeting is to include a demo and be ready with more cleanup based on feedback from today.


## Feb 14th, 2023

Recording: [https://www.youtube.com/watch?v=NsV9TVcJw54](https://www.youtube.com/watch?v=NsV9TVcJw54) 



* [SergeyKanzhelev] we are tracking 23 KEPs this release: [https://github.com/orgs/kubernetes/projects/117/views/1?filterQuery=+status%3ATrac[…]ode+&sortedBy%5Bdirection%5D=desc&sortedBy%5BcolumnId%5D=Status](https://github.com/orgs/kubernetes/projects/117/views/1?filterQuery=+status%3ATracked+label%3Asig%2Fnode+&sortedBy%5Bdirection%5D=desc&sortedBy%5BcolumnId%5D=Status)
* Cut from milestone:
    * [Kubelet Resource Plugin Manager Refactor · Issue #3675](https://github.com/kubernetes/enhancements/issues/3675)
    * [Add AppArmor support · Issue #24 · kubernetes/enhancements · GitHub](https://github.com/kubernetes/enhancements/issues/24)
    * [cAdvisor-less, CRI-full Container and Pod Stats · Issue #2371 · kubernetes/enhancements · GitHub](https://github.com/kubernetes/enhancements/issues/2371)
    * [Support class-based resources to CRI protocol · Issue #3008](https://github.com/kubernetes/enhancements/issues/3008) 
* [SergeyKanzhelev] Annual report tasks: [https://github.com/kubernetes/community/blob/master/sig-node/annual-report-2022.md](https://github.com/kubernetes/community/blob/master/sig-node/annual-report-2022.md)  
* [haircommander] [automatic cgroup driver matching between kubelet and CRI](https://github.com/kubernetes/kubernetes/issues/99808)
    * [paco]+1 for this. I opened a PR to do it for dockershim [#98357](https://github.com/kubernetes/kubernetes/pull/98357/) before.
    * consensus is to have the runtime be the source of  truth, and report to the Kubelet what cgroup driver to use with a field in the runtime status call. 
    * Peter to pursue a KEP for this in 1.28
    * [kad] while we are on this, we can also eliminate misconfiguration between kubelet and runtimes on reserved/infra cpuset: `--reserved-cpus` for kubelet/--infra-ctr-cpuset in cri-o.
* [Yuan Chen, Deep Debroy]] Discuss decoupling no-execute-taint-manager from node-lifecycle-controller
    * Mainly a refactoring change as the two are part of the same controller but have somewhat independent functions.
    * With [https://github.com/kubernetes/enhancements/blob/74e610bb0f7e40862688e8a434c77bfafc53cb9e/keps/sig-scheduling/20200114-taint-based-evictions.md](https://github.com/kubernetes/enhancements/blob/74e610bb0f7e40862688e8a434c77bfafc53cb9e/keps/sig-scheduling/20200114-taint-based-evictions.md), the ability to disable no-execute-taint-manager is going away.
    * [paco] [karpenter](https://github.com/aws/karpenter) has a requirement of `[startup taint](https://github.com/kubernetes/kubernetes/issues/115139#issuecomment-1424597021)` which is interesting. I am not sure if this can be a possible user case.
* [MikeBrown] Just a reminder probe granularity enhancement [https://github.com/kubernetes/enhancements/pull/3067](https://github.com/kubernetes/enhancements/pull/3067) needs review.


## Feb 7th, 2023

Recording: [https://www.youtube.com/watch?v=cam97qjy8qE](https://www.youtube.com/watch?v=cam97qjy8qE) 

27 KEPs: [https://github.com/kubernetes/enhancements/issues?q=is%3Aissue+is%3Aopen+label%3Asig%2Fnode+milestone%3Av1.27+label%3Alead-opted-in+](https://github.com/kubernetes/enhancements/issues?q=is%3Aissue+is%3Aopen+label%3Asig%2Fnode+milestone%3Av1.27+label%3Alead-opted-in+) 



* [rata]: [Userns KEP PR](https://github.com/kubernetes/enhancements/pull/3811) rework with idmap mounts.
    * I think this should be ready to approve for 1.27. Is anything missing?
* [klueska] KEP with PodResources extensions for DRA
    * Looking for final approval by  and / or @mrunal
    * There’s one small change I’d like to see made, but if we get an /approve with a /hold I can make sure the change gets in before giving a final /lgtm
    * [https://github.com/kubernetes/enhancements/pull/3738](https://github.com/kubernetes/enhancements/pull/3738)
* ~~[klueska] Milestone and tracking for updates to DRA enhancement issue~~
    * ~~Needs **<code>lead-opted-in</code></strong> and <strong><code>milestone 1.27</code></strong> label (staying in alpha)</del>
    * <del>[https://github.com/kubernetes/enhancements/issues/3063](https://github.com/kubernetes/enhancements/issues/3063)</del>
* [vinaykul] [InPlace Pod Vertical Scaling](https://github.com/kubernetes/kubernetes/pull/102884/) PR - status update
    * Please review and merge k/enhancements housekeeping [PR #3845](https://github.com/kubernetes/enhancements/pull/3845)
        * To catch up to the latest KEP template, the PR adds integration test section and responds to node scalability section.
    * I have rebased and updated [PR #102884](https://github.com/kubernetes/kubernetes/pull/102884) after LGTM by @thockin 
        * Squashed all API commits into single commit + generated files commit
        * Separate commit for scheduler changes
        * I plan to squash various kubelet commits into 2 or 3 commits if that’s ok
        * @thockin awaits Derek’s re-LGTM/approve before approving the PR.
        * I plan to create follow-up PRs to address a few outstanding items:
            * ResizePolicy name restructuring
            * Use PodStatus.QOSClass instead of GetPodQoS across K8s codebase
* [Atanas] CCI KEP:
    * [https://github.com/kubernetes/enhancements/pull/3853](https://github.com/kubernetes/enhancements/pull/3853)
    * Addressing comments as they come in.
    * Brief discussion on anything else outstanding.
    * Reviewers: Swati and Kevin, Approver: Dawn or Derek
* [qbarrand] [Kernel Module Management](https://github.com/kubernetes-sigs/kernel-module-management)
    * looking for sponsors to get [mresvanis](https://github.com/mresvanis) into the kubernetes / kubernetes-sigs orgs - [contributions](https://github.com/kubernetes-sigs/kernel-module-management/commits?author=mresvanis)
    * [PR to update admins and maintainers](https://github.com/kubernetes/org/pull/3791)
* [mimowo] Ask for final review and approval from sig-node for: [Update for second Beta with GA criteria for "KEP-3329: Retriable and non-retriable Pod failures for Jobs"](https://github.com/kubernetes/enhancements/pull/3757)


## Jan 31st, 2023

Recording: [https://youtu.be/96DTU9ncSLA](https://youtu.be/96DTU9ncSLA) 



* [Aravindh Puthiyaparambil] [KEP 2258: Node service log viewer](https://github.com/kubernetes/enhancements/pull/2271#top)
    * [KEP 2258: add node log query](https://github.com/kubernetes/kubernetes/pull/96120#top) (implementation
    * Discuss Jordan’s [comment](https://github.com/kubernetes/kubernetes/pull/96120#pullrequestreview-1262371094)
    * sig-windows is interested in the feature as it helps debugging Windows nodes. We would like to get the viewpoint of #sig-node whether we should move forward or not based on Jordan’s feedback ie. add the feature disabled by default and a warning when enabling.
* [mimowo] Heads up and discuss Kubelet changes for [Update for second Beta with GA criteria for "KEP-3329: Retriable and non-retriable Pod failures for Jobs"](https://github.com/kubernetes/enhancements/pull/3757)
    * [Sergey] good direction, original language was a bit cryptic, need more examples
        * [mimowo] it is already fixed
    * [mrunal] related to pod lifecycle - Ryan and David might need to review
    * [David]+1 yes it is related to pod lifecycle, will review as well
* [pacoxu] Two bugfix PRs need to be reviewed for `[LocalStorageCapacityIsolationFSQuotaMonitoring](https://github.com/kubernetes/enhancements/issues/1029)` (which is promoted to beta in v1.25.0 and revert to alpha in v1.25.1 and v1.26 for the bug) for the same issue: #[112624](https://github.com/kubernetes/kubernetes/pull/112624) and #[115314](https://github.com/kubernetes/kubernetes/pull/115314). If this PR can be reviewed and merged on time. I prefer to re-promote the feature to beta in v1.27.
    * [Ryan] It was discussed Jan 17th. Thinking about deprecating it. Are you using it?
    * [Paco] let’s sync up
        * 112624 is merging
        * Paco is going to get an upstream E2E running with the feature gate enabled in test-grid
* [klueska] Update to the DRA KEP for CRI changes
    * Ready ready for final approval from 
    * [https://github.com/kubernetes/enhancements/pull/3731](https://github.com/kubernetes/enhancements/pull/3731)
* [klueska] Update to the DRA KEP for WatchConfiguration call on kubelet plugin
    * Reviews welcome
    * [https://github.com/kubernetes/enhancements/pull/3802](https://github.com/kubernetes/enhancements/pull/3802)
* [atanas & mweston] Container Compute Interface (was Kubelet Resource Plugin) discussion
    * Notes from this morning: [https://docs.google.com/document/d/1kRSkOqZnalt09UMm-xXcHPuw3ZEzPrFIuhFf2piwIxI/edit](https://docs.google.com/document/d/1kRSkOqZnalt09UMm-xXcHPuw3ZEzPrFIuhFf2piwIxI/edit)
    * Will get recording up to Sergey today.
    * PR up here:  [https://github.com/kubernetes/enhancements/pull/3803](https://github.com/kubernetes/enhancements/pull/3803)
        * Please add in any comments
    * Opens: 
        * Describe Kubelet restart scenario
        * Sketch and describe beta architecture
        * Illustrate failure and corner case flows
* [claudiubelu] Proposed changes to how kubelet detects updates for registered plugins b/c current implementation doesn’t work on Windows due to timestamp granularity issues [https://github.com/kubernetes/kubernetes/pull/114136](https://github.com/kubernetes/kubernetes/pull/114136) 
    * [https://cs.k8s.io/?q=PluginExistsWithCorrectTimestamp&i=nope&files=&excludeFiles=&repos=](https://cs.k8s.io/?q=PluginExistsWithCorrectTimestamp&i=nope&files=&excludeFiles=&repos=)
    * [https://testgrid.k8s.io/sig-windows-signal#windows-unit-master](https://testgrid.k8s.io/sig-windows-signal#windows-unit-master) test failure caused by timestamp
    * this has been proposed and brought up in the past, but no reviews came in.
    * it is one of the last consistently failing tests in the testgrid above.
    * [MikeBrown] consider keeping existing public method and add another one, deprecate first one later. nod just to protect custom plugin managers test buckets etc .. deprecate first thx.. see:
        * [https://cs.k8s.io/?q=PluginExistsWithCorrectTimestamp&i=nope&files=&excludeFiles=&repos=](https://cs.k8s.io/?q&#61;PluginExistsWithCorrectTimestamp&i&#61;nope&files&#61;&excludeFiles&#61;&repos&#61;)
* [vinaykul] [InPlace Pod Vertical Scaling](https://github.com/kubernetes/kubernetes/pull/102884/) PR - status update
    * Can we please get this KEP [officially tracked](https://github.com/orgs/kubernetes/projects/117/views/1) for 1.27?
    * thockin prefers that we merge [PR 102884](https://github.com/kubernetes/kubernetes/pull/102884) in its entirety as opposed to merging API [PR 111946](https://github.com/kubernetes/kubernetes/pull/111946) followed by the rebased implementation PR a week later.
        * Reason: [potential difficulty in unwinding the API merge](https://github.com/kubernetes/kubernetes/pull/111946#pullrequestreview-1239501555) in the event we find something really bad after merging implementation a week later, when other commits have rebased and merged on top of the API changes.
        * I have no major concerns with merging everything in one shot.
            * We can re-add periodic CI jobs afterwards and iterate on fixes, if needed.
            * We are ~6 weeks away to code-freeze. It is best to merge #102884 in one shot at this point while we still have sufficient buffer to merge this change and fix any issues without risk to the release.

[KEPS REVIEW]: 15 minutes


## Jan 24th, 2023

Recording: [https://youtu.be/NQaTeTfI9UY](https://youtu.be/NQaTeTfI9UY) 


<table>
  <tr>
   <td><strong>Incoming</strong>
   </td>
   <td>
   </td>
   <td><strong>Completed</strong>
   </td>
   <td>
   </td>
  </tr>
  <tr>
   <td>Created:
   </td>
   <td><strong><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++created%3A2023-01-17T17%3A00%3A00%2B0000..2023-01-24T17%3A53%3A26%2B0000">29</a></strong>
   </td>
   <td>Closed:
   </td>
   <td><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++is%3Aunmerged+closed%3A2023-01-17T17%3A00%3A00%2B0000..2023-01-24T17%3A53%3A26%2B0000">12</a>
   </td>
  </tr>
  <tr>
   <td>Updated:
   </td>
   <td><strong><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+updated%3A2023-01-17T17%3A00%3A00%2B0000..2023-01-24T17%3A53%3A26%2B0000+created%3A%3C2023-01-17T17%3A00%3A00%2B0000">90</a></strong>
   </td>
   <td>Merged:
   </td>
   <td><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++merged%3A2023-01-17T17%3A00%3A00%2B0000..2023-01-24T17%3A53%3A26%2B0000">14</a>
   </td>
  </tr>
</table>




* [SergeyKanzhelev] Sidecar containers KEP: [https://github.com/kubernetes/enhancements/pull/3761](https://github.com/kubernetes/enhancements/pull/3761) 

    [https://docs.google.com/presentation/d/1JOfC4YY6tJC4zf83QobPG9W_jlSw1JUhForMm9JHQVw/edit#slide=id.g1eaa3876499_0_0](https://docs.google.com/presentation/d/1JOfC4YY6tJC4zf83QobPG9W_jlSw1JUhForMm9JHQVw/edit#slide=id.g1eaa3876499_0_0) 

* [bobbypage] Thoughts on formalizing Node Lifecycle - [https://github.com/kubernetes/kubernetes/issues/115139](https://github.com/kubernetes/kubernetes/issues/115139) 
    * [Dawn] we might had several docs in OSS - maybe Lantao knows
    * [MikeBrown] would be interested to participate
    * [Sergey] question on slack recently regarding taint for paused VMs. Also recreating node with the same name is another interesting qq here.
* [asierHuawei] Thoughts on IMA namespace support for pods

    [https://github.com/kubernetes/enhancements/pull/3703](https://github.com/kubernetes/enhancements/pull/3703)

* [mrunalp] kernel changes are not in yet. Then OCI, than CRI, than k8s. 
* [Asier] runc had the same feedback
* [Dawn] feature is interesting, we need it implemented in kernel first
* [klueska] Update to the DRA KEP for CRI changes
    * Would appreciate review from @mrunal given his background with CRI
    * Otherwise ready for final approval
    * [https://github.com/kubernetes/enhancements/pull/3731](https://github.com/kubernetes/enhancements/pull/3731)
* [klueska] KEP posted to extend PodResources API with CDI device information
    * Reviews welcome
    * [https://github.com/kubernetes/enhancements/pull/3738](https://github.com/kubernetes/enhancements/pull/3738)
* [vinaykul] [InPlace Pod Vertical Scaling](https://github.com/kubernetes/kubernetes/pull/102884/) PR - status update
    * thockin prefers that we merge [PR 102884](https://github.com/kubernetes/kubernetes/pull/102884) in its entirety as opposed to merging API [PR 111946](https://github.com/kubernetes/kubernetes/pull/111946) followed by the rebased implementation PR a week later.
        * Reason: [potential difficulty in unwinding the API merge](https://github.com/kubernetes/kubernetes/pull/111946#pullrequestreview-1239501555) in the event we find something really bad after merging implementation a week later, when other commits have rebased and merged on top of the API changes.
        * I have no major concerns with merging everything in one shot.
            * We can re-add periodic CI jobs afterwards and iterate on fixes, if needed.
* [SergeyKanzhelev] http probes and leaking sockets: <code>[https://github.com/kubernetes/kubernetes/pull/115143](https://github.com/kubernetes/kubernetes/pull/115143)</code>
    * Context: Sig Network from Jan 19th 2023: [aojea] 10 - 15 mins Presentation - “When sockets refuse to die”
        * Catchy title from [https://blog.cloudflare.com/when-tcp-sockets-refuse-to-die/](https://blog.cloudflare.com/when-tcp-sockets-refuse-to-die/)
        * [https://github.com/kubernetes/kubernetes/pull/115143](https://github.com/kubernetes/kubernetes/pull/115143)
        * [https://docs.google.com/presentation/d/1UYs1tSFeX7-Jjqhpqacn_nlH88LOLr37VOZBPGwhbCI/edit?usp=sharing](https://docs.google.com/presentation/d/1UYs1tSFeX7-Jjqhpqacn_nlH88LOLr37VOZBPGwhbCI/edit?usp=sharing)
* [swsehgal] Request for reviews:
    * Topology Manager GA graduation KEP: [https://github.com/kubernetes/enhancements/pull/3745](https://github.com/kubernetes/enhancements/pull/3745)
        * Dawn to take a look.
        * Also, milestone and leads-opt-in label needs to be added to the issue: [https://github.com/kubernetes/enhancements/issues/693](https://github.com/kubernetes/enhancements/issues/693)
    * Device Manager [Bug](https://github.com/kubernetes/kubernetes/issues/109595): 
        * PR with fix: [https://github.com/kubernetes/kubernetes/pull/114640](https://github.com/kubernetes/kubernetes/pull/114640)
        * The above PR depends on sample device plugin changes: [https://github.com/kubernetes/kubernetes/pull/115107](https://github.com/kubernetes/kubernetes/pull/115107) for e2e test implementation. Sample device plugin image would have to be updated(promoted) so it can be consumed for e2e testing.
* [mweston] Reminder of kubelet resource plugin discussion next Tuesday, 6-7am PST:
    * [https://us02web.zoom.us/j/82567156922?pwd=Q2xscE0rRjluRTlvdk5FK3hzUFpDQT09](https://us02web.zoom.us/j/82567156922?pwd=Q2xscE0rRjluRTlvdk5FK3hzUFpDQT09)

		Notes from previous meetings: [https://docs.google.com/document/d/1ALxPqeHbEc0QOIzJ3rWWPpwRMRlYDzCv0mu2mR4odR8/edit#](https://docs.google.com/document/d/1ALxPqeHbEc0QOIzJ3rWWPpwRMRlYDzCv0mu2mR4odR8/edit#) 

		



* [ddebroy] KEP for pod condition to indicate Pod Sandbox creation updated to Beta with new name:
    * Updated the name of the condition from `PodHasNetwork` to `PodReadyToStartContainers` as discussed with Derek and Tim Hockins to align better with sig network concerns with the original name.
    * [https://github.com/kubernetes/enhancements/pull/3778](https://github.com/kubernetes/enhancements/pull/3778)
* [jackfrancis] remove ExecProbeTimeout feature gate
    * ​​[https://github.com/kubernetes/kubernetes/pull/115227](https://github.com/kubernetes/kubernetes/pull/115227)
    * We need to clarify criteria for removing this feature gate. Background:
        * Issue for deciding when to remove feature gate:
            * [https://github.com/kubernetes/kubernetes/issues/99854](https://github.com/kubernetes/kubernetes/issues/99854)
        * Probe duration metrics feature landed:
            * [https://github.com/kubernetes/kubernetes/pull/104484](https://github.com/kubernetes/kubernetes/pull/104484)
            * Released in v1.25.0
                * [https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.25.md#feature-5](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.25.md#feature-5)
        * Lock to true PR (never merged)
            * [https://github.com/kubernetes/kubernetes/pull/107480](https://github.com/kubernetes/kubernetes/pull/107480)
        * This feature gate only apples to dockershim, removed in v1.24.0:
            * [https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.24.md#dockershim-removed-from-kubelet](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.24.md#dockershim-removed-from-kubelet)


## Jan 17th, 2023

Recording: [https://youtu.be/wirWRKSqY10](https://youtu.be/wirWRKSqY10) 

Total PRs: [217](https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+is%3Aopen+)


<table>
  <tr>
   <td><strong>Incoming</strong>
   </td>
   <td>
   </td>
   <td><strong>Completed</strong>
   </td>
   <td>
   </td>
  </tr>
  <tr>
   <td>Created:
   </td>
   <td><strong><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++created%3A2023-01-10T17%3A00%3A00%2B0000..2023-01-17T17%3A56%3A21%2B0000">30</a></strong>
   </td>
   <td>Closed:
   </td>
   <td><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++is%3Aunmerged+closed%3A2023-01-10T17%3A00%3A00%2B0000..2023-01-17T17%3A56%3A21%2B0000">16</a>
   </td>
  </tr>
  <tr>
   <td>Updated:
   </td>
   <td><strong><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+updated%3A2023-01-10T17%3A00%3A00%2B0000..2023-01-17T17%3A56%3A21%2B0000+created%3A%3C2023-01-10T17%3A00%3A00%2B0000">103</a></strong>
   </td>
   <td>Merged:
   </td>
   <td><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++merged%3A2023-01-10T17%3A00%3A00%2B0000..2023-01-17T17%3A56%3A21%2B0000">16</a>
   </td>
  </tr>
</table>




* [rphillips] ​​[https://github.com/kubernetes/kubernetes/issues/114506](https://github.com/kubernetes/kubernetes/issues/114506) deprecating the feature unless somebody wants to pick it up
    * [mrunal] let’s summarize the current status and blockers
        * [Ryan] user can change project id of a directory which becomes a security problem.
        * [mrunal] can we check with upstream?
    * [Dawn] this is a kernel feature and we don’t have an efficient way to track disk usage. Since there is no efficient way to do it per cgroup there was an idea to do it per process. Kernel implementation is still not what we want. +1 to mrunal if we can ask linux community
    * [Ryan] any objections deprecating it?
* [vinaykul] [InPlace Pod Vertical Scaling](https://github.com/kubernetes/kubernetes/pull/102884/) PR - status update
    * Please review and merge [KEP update PR](https://github.com/kubernetes/enhancements/pull/3670)
    * thockin prefers that we merge [PR 102884](https://github.com/kubernetes/kubernetes/pull/102884) in its entirety as opposed to merging API [PR 111946](https://github.com/kubernetes/kubernetes/pull/111946) followed by the rebased implementation PR a week later.
        * Tim is worried about [potential difficulty in unwinding the API merge](https://github.com/kubernetes/kubernetes/pull/111946#pullrequestreview-1239501555) in the event we find something really bad after merging implementation a week later, when other commits have rebased and merged on top of the API changes.
        * I have no major concerns with merging the mothership in one shot.
            * We can re-add periodic CI jobs afterwards and iterate on fixes if there are any issues.
    * Tim & I discussed a [naming change to ResizePolicy](https://github.com/kubernetes/kubernetes/pull/111946#discussion_r1063862866). I plan to do this as a follow-up PR on the heels of #102884 in order to avoid resetting the reviews.
* [marquiz] [QoS-class resources KEP](https://github.com/kubernetes/enhancements/pull/3004), proposal updated:
    * ditched pod annotation based UX -> go straight to K8s API
    * added support for “class capacity” i.e. possibility limit max number of users of classes
* [Vinay] [https://groups.google.com/u/3/g/kubernetes-sig-network/c/e0U44XyI3Vw](https://groups.google.com/u/3/g/kubernetes-sig-network/c/e0U44XyI3Vw) CRI and CNI.
* [SergeyKanzhelev] [https://github.com/kubernetes/kubernetes/pull/114989](https://github.com/kubernetes/kubernetes/pull/114989) 
* [SergeyKanzhelev] sidecar container WG update: [https://docs.google.com/document/d/1E1guvFJ5KBQIGcjCrQqFywU9_cBQHRtHvjuqcVbCXvU/edit#heading=h.m8xoiv5t6qma](https://docs.google.com/document/d/1E1guvFJ5KBQIGcjCrQqFywU9_cBQHRtHvjuqcVbCXvU/edit#heading=h.m8xoiv5t6qma) 
* [atanas] Kubelet plugin WG update (next meeting 1/31/2023 [https://us02web.zoom.us/j/82567156922?pwd=Q2xscE0rRjluRTlvdk5FK3hzUFpDQT09](https://us02web.zoom.us/j/82567156922?pwd=Q2xscE0rRjluRTlvdk5FK3hzUFpDQT09) (pw is 77777 if it asks) at 6am PST/8am CST/2pm Ireland)


## Jan 10th, 2023

Recording: [https://youtu.be/5V0uRxH4O4k](https://youtu.be/5V0uRxH4O4k) 



* ~~[pacoxu] KEP-3610: namespace-wide global env injection [#3612](https://github.com/kubernetes/enhancements/pull/3612), not sure if this can be an admission controller.(removed due to [mutating CEL admission should be the final solution.](https://github.com/kubernetes/enhancements/pull/3612#discussion_r1063236570)) ~~
* [ruiwen/pacoxu] KEP-3673: Kubelet limit of Parallel Image Pulls[ #3713](https://github.com/kubernetes/enhancements/pull/3713)
    * 
* [klueska] Update CRI to include CDI devices (needed by DRA before moving to beta)
    * Do we need a new KEP or can we update the existing DRA KEP with details?
    * **Mrunal**: We can update the existing KEP with the details of the change.
    * It should just be a simple addition to: \
[https://github.com/kubernetes/cri-api/blob/c75ef5b/pkg/apis/runtime/v1/api.proto#L682](https://github.com/kubernetes/cri-api/blob/c75ef5b/pkg/apis/runtime/v1/api.proto#L682)
* [QuentinN42]  Add FileEnvSource and FileKeySelector to add environment generated on the fly [#114674](https://github.com/kubernetes/kubernetes/issues/114674)
    * Sourcing from any file from any source may be too big of a scope. Would limiting this to empty dir files be enough?
    * Security - is there a risk to source some secret as an environment variable that would expose the file that wasn’t available otherwise.
    * **Action**: Need to move this to kubernetes/enhancements as a KEP and follow the process. => [https://github.com/kubernetes/enhancements/issues/3721](https://github.com/kubernetes/enhancements/issues/3721)
    * [Mike Brown] fyi.. not sure if this is the right pattern but NRI plugins support modifying environment variables for the containers. might be useful at least for prototyping
    * [QuentinN42] another question is error conditions depending on the file format
    * [Alexander Kanevsky] my first impression - the env variables are populated in oci spec before container started. sourcing from some file inside container might be not feasible....
    * [Mike Brown] right would require a set for any env change happening in prestart (which could be done by setting a runc hook via NRI or hook schema, or just doing the set on the update response) 
* [vinaykul] [InPlace Pod Vertical Scaling](https://github.com/kubernetes/kubernetes/pull/102884/) PR - status update
    * I won’t be in the Node meeting today due to another 10 am meeting.
    * Please review and merge [KEP update PR](https://github.com/kubernetes/enhancements/pull/3670)
        * Updated beta target to v1.29
        * Added details on handling version skew.
    * Tim prefers that we merge [PR 102884](https://github.com/kubernetes/kubernetes/pull/102884) in its entirety as opposed to merging API [PR 111946](https://github.com/kubernetes/kubernetes/pull/111946) followed by the rest of it a week later.
        * We can re-add periodic CI jobs afterwards and iterate on fixes if there are any issues.
        * I believe this will require both Derek’s & Tim’s lgtm & approve.
        * **AI: **Derek to catch up on Tims’ objections: [https://github.com/kubernetes/kubernetes/pull/111946#pullrequestreview-1239501555](https://github.com/kubernetes/kubernetes/pull/111946#pullrequestreview-1239501555)
* [derek] sig updates
    * email with proposed changes is coming up later today
    * [https://groups.google.com/g/kubernetes-sig-node/c/NsoYU1Y2rUs](https://groups.google.com/g/kubernetes-sig-node/c/NsoYU1Y2rUs)


## Jan 3rd, 2023

Recording: [https://www.youtube.com/watch?v=AG3U91-5keo](https://www.youtube.com/watch?v=AG3U91-5keo) 

Total active pull requests: [205](https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+is%3Aopen+)


<table>
  <tr>
   <td><strong>Incoming</strong>
   </td>
   <td>
   </td>
   <td><strong>Completed</strong>
   </td>
   <td>
   </td>
  </tr>
  <tr>
   <td>Created:
   </td>
   <td><strong><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++created%3A2022-12-13T17%3A00%3A00%2B0000..2023-01-03T17%3A56%3A53%2B0000">45</a></strong>
   </td>
   <td>Closed:
   </td>
   <td><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++is%3Aunmerged+closed%3A2022-12-13T17%3A00%3A00%2B0000..2023-01-03T17%3A56%3A53%2B0000">22</a>
   </td>
  </tr>
  <tr>
   <td>Updated:
   </td>
   <td><strong><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+updated%3A2022-12-13T17%3A00%3A00%2B0000..2023-01-03T17%3A56%3A53%2B0000+created%3A%3C2022-12-13T17%3A00%3A00%2B0000">144</a></strong>
   </td>
   <td>Merged:
   </td>
   <td><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++merged%3A2022-12-13T17%3A00%3A00%2B0000..2023-01-03T17%3A56%3A53%2B0000">18</a>
   </td>
  </tr>
</table>




* [SergeyKanzhelev] [https://github.com/kubernetes/kubernetes/pull/114394](https://github.com/kubernetes/kubernetes/pull/114394) CRI API version skew policies. See [slides](https://docs.google.com/presentation/d/1s0RM2zDXKTnTn5Yx0_20V0TPbPJLCQSI/edit) from contributors summit for extra details 
* ~~[SergeyKanzhelev] Reconcile SIG Node teams and OWNERs files: [https://github.com/kubernetes/org/pull/3893](https://github.com/kubernetes/org/pull/3893) ~~
* [vinaykul] [InPlace Pod Vertical Scaling](https://github.com/kubernetes/kubernetes/pull/102884/) PR - status update
    * Happy 2023!
    * Please review and merge [KEP milestone update PR](https://github.com/kubernetes/enhancements/pull/3670)
    * [PR 102884](https://github.com/kubernetes/kubernetes/pull/102884) approved by Derek.
        * @bobbypage fixed containerd/main E2E pull test job, we now have full E2E coverage (verifies values from ContainerStatus CRI response)
            * The test has established a [history of successful runs](https://prow.k8s.io/job-history/gs/kubernetes-jenkins/pr-logs/directory/pull-kubernetes-e2e-inplace-pod-resize-containerd-main-v2) for PR 102884 over the course of a few rebases.
        * My recommendation is we merge API changes [PR 111946](https://github.com/kubernetes/kubernetes/pull/111946) at the earliest possible point in 1.27 and watch it to see nothing bad happens.
            * Can we do it this week?
            * Can we atleast merge feature gate definition to clean up test failures in [unrelated PRs](https://github.com/kubernetes/kubernetes/pull/114185)?
        * And then merge PR 102884 shortly after (PR 111946 merge + 1 week)
            * We can then re-add periodic CI test jobs.
* [Seaiii] [https://github.com/kubernetes/kubernetes/pull/113883](https://github.com/kubernetes/kubernetes/pull/113883) The second time the pod is deleted the grace period does not take effect .Please review update PR PR 113883
