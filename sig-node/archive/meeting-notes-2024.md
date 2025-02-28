# SIG Node Meeting Notes

## Dec 31, 2024

* Cancelled

## Dec 24, 2024

* Cancelled

## Dec 17, 2024

* Cancelled

## Dec 10, 2024

- 1.32 retro: [SIG Node 1.32 retro](https://docs.google.com/document/d/1CM_WLChPzAx2VLFCXR0RNvNcqeV0OZKL8iqlBFVrpHY/edit?tab=t.0#heading=h.gmyk4fetligd)

## Dec 3, 2024

- \[minna\] asking for some PR feedback [https://github.com/kubernetes/kubernetes/pull/125918](https://github.com/kubernetes/kubernetes/pull/125918)  
  - \[Peter\] We should add a feature gate beta \+ on by default  
  - \[Francesco\] \+ 1 and we should extend   
    - Maybe wait for critical pods to be ready and not just started before we try to start non critical pods  
  - \[Sergey\] Similarly we could extend logic for admission  
  - \[Sergey\] It’s possible this PR may switch starting failure to admission failure (if critical pod starts and fails, the pods that rely on them will fail differently)  
- \[Sergey\] Add agenda items ASAP, as we will cancel the meeting aggressively in December

## Nov 26, 2024 \[Canceled due to US holiday\]

## Nov 19, 2024 \[Canceled due to lack of agenda\]

## Nov 12, 2024 \[Canceled for KubeCon\]

## Nov 5, 2024 

## Recording: [https://www.youtube.com/watch?v=1u\_yKruHeZU](https://www.youtube.com/watch?v=1u_yKruHeZU)

* # \[danwinship/surya\] [Redesigning Kubelet Probes](https://docs.google.com/presentation/d/1XujDtyhIkZ7FPDPck9qou-L1O80a_IomCGNf6I5E9X4/edit#slide=id.p)

  * antonio had opened an issue for runtime to do the checks  
    * when kubetlet requests runtime to do probe  
      * launching new pods and containers would be heavy  
      * can we re-use the container-monitor process here ? instead of adding new ones?  
        * tcp/http/grpc types of probes  
          * would containerd/cri-o be able to do those probes?  
          * \[mrunal\] containerd would have to do learn the split of daemon  
    * \[dawn\] the pod sounds better than what we have today?  
      * cost to the user though here at the application level usage is unpredictable \- this is not worse than what we have today but there is a complexity for the user (with per container case)  
      * probing pod is part of system overhead   
  * will this be a new type of probe? replacement of existing probes?  
    * if its a pod probe then some features like ensuring the port is open might be lost?  
    * so maybe we should keep both types of probes and users can  
  * Performance should not regress  
  * checking a file in the filesystem and letting users put what they want?  
* \[tallclair\] In-Place Pod Resize: status update  
  * [Beta Dashboard](https://github.com/orgs/kubernetes/projects/178/views/1?filterQuery=is%3Aissue+-status%3Adone+-is%3Aclosed+roadmap%3Abeta+&visibleFields=%5B%22Title%22%2C%22Assignees%22%2C%22Status%22%2C145416682%2C87750681%2C%22Linked+pull+requests%22%5D&sortedBy%5Bdirection%5D=asc&sortedBy%5BcolumnId%5D=145416682)

## \`1Oct 29, 2024 (Canceled)

Canceled due to lack of the agenda.

## Oct 22, 2024

- \[Kevin Hannon in place of dims\] cadvisor for 1.32  
  - [https://github.com/google/cadvisor/pull/3609](https://github.com/google/cadvisor/pull/3609) ( Reduce the dependencies we drag into cadvisor AND drag into k/k through cadvisor )  
  - [https://github.com/google/cadvisor/pull/3608](https://github.com/google/cadvisor/pull/3608) ( help the periodic CI job to recover )  
  - fix for [https://github.com/google/cadvisor/issues/3577](https://github.com/google/cadvisor/issues/3577) as well  
  - Release may be needed  
  - https://kubernetes.slack.com/archives/C0BP8PW9G/p1729517493050419  
- \[Kevin Hannon\] Swap Based Eviction  
  - [https://github.com/kubernetes/kubernetes/pull/128137](https://github.com/kubernetes/kubernetes/pull/128137)  
  -   
- \[Lakshmi\]  Requesting for review and feedback on PR  
  - [https://github.com/kubernetes/website/pull/48001](https://github.com/kubernetes/website/pull/48001)  
- \[pehunt\] libcontainer \+ runc \+ k8s  
  - two pieces  
    - runc 1.2.0 just came out, k8s wants to use it (to get PSI stats)  but there are concerns about containerd using a different libcontainer version from cadvisor  
      - [https://kubernetes.slack.com/archives/C0BP8PW9G/p1729606639892799](https://kubernetes.slack.com/archives/C0BP8PW9G/p1729606639892799)   
      - [https://cloud-native.slack.com/archives/CGEQHPYF4/p1729607023643899](https://cloud-native.slack.com/archives/CGEQHPYF4/p1729607023643899)   
      - [https://github.com/google/cadvisor/pull/3083\#issuecomment-2429370533](https://github.com/google/cadvisor/pull/3083#issuecomment-2429370533)   
      - Do we need to wait for 1.2.0 in 2.0, or can we backport, or can we run disjoint? we’ve waited a long time for 1.2.0 and I’d like to use it  
    - libraryfication of libcontainer: currently, we’re vendoring runc libcontainer in k8s, and this means we’re version locked with the runc binary (which doesn’t have k8s as a priority with release cadence)  
      - Discussions on moving the libcontainer/cgroups library out of runc and into its own repo [https://github.com/kubernetes/kubernetes/issues/128157](https://github.com/kubernetes/kubernetes/issues/128157)   
      - [Peter Hunt](mailto:pehunt@redhat.com)to send an email to sig-node mailing list to notify folks of this plan  
      - part of this plan [https://github.com/kubernetes/kubernetes/pull/128245](https://github.com/kubernetes/kubernetes/pull/128245)    
  - 

## Oct 15, 2024

Recording: [https://www.youtube.com/watch?v=MyOhDhHRRKk](https://www.youtube.com/watch?v=MyOhDhHRRKk)

- \[Sergey\] New Feature Gates emulation mode and features GA:  [https://github.com/kubernetes/kubernetes/pull/126981\#discussion\_r1799779745](https://github.com/kubernetes/kubernetes/pull/126981#discussion_r1799779745)  
  - Should we keep removing code in kubelet as before? Or just keep it around the same way as we do for API server to minimize possible errors and simply not test it?  
- \[Chris\] A demo for k8s dynamic batch workloads:  
  [https://github.com/chrishenzie/k8s-dynamic-batch-demo](https://github.com/chrishenzie/k8s-dynamic-batch-demo)  
- \[pehunt\]  (defer to end, only if there’s time) beginnings of swap aware eviction discussion  
- 

## Oct 8, 2024

Recording: [https://www.youtube.com/watch?v=\_Zexxr4pxr8](https://www.youtube.com/watch?v=_Zexxr4pxr8)

- \[Sergey\] Containerd 2.0 and KEPs: [https://groups.google.com/g/kubernetes-sig-architecture/c/kft-wa929\_Q](https://groups.google.com/g/kubernetes-sig-architecture/c/kft-wa929_Q)  
  - Are we promoting to beta with a single runtime implementation?  
  - What is the production test requirement for the feature? (in case of 2.0 \- how do we measure exposure of the feature to prod?)  
- ~~\[fromani\] Heads up: KEP 4885 will introduce a new memory manager policy~~  
  - ~~the windows and linux will support different policies~~  
  - ~~do we prefer to postpone the memory manager GA graduation?~~   
- \[fromani\] unblocking [https://github.com/kubernetes/kubernetes/issues/70585](https://github.com/kubernetes/kubernetes/issues/70585) with a [feature gate](https://kubernetes.slack.com/archives/C5P3FE08M/p1727072265347319)?  
- \[Eddie\] Request for KEP review: [Mutable CSINode Allocatable Property](https://github.com/kubernetes/enhancements/issues/4876)  
- \[pehunt\] FYI for approvers: two new KEPs have been added to the milestone and don’t have an approver  
  - [https://github.com/kubernetes/enhancements/issues/3619](https://github.com/kubernetes/enhancements/issues/3619)  
  - [https://github.com/kubernetes/enhancements/issues/4753](https://github.com/kubernetes/enhancements/issues/4753)   
  - \[fromani\] [https://github.com/kubernetes/enhancements/issues/4885](https://github.com/kubernetes/enhancements/issues/4885) lacks approver also. I’m reviewing and almost LGTM (almost \= need a final pass but no outstanding issues after last update)  
    - \[pehunt\] according to [The KEP board](https://github.com/orgs/kubernetes/projects/186/views/7?filterQuery=status%3A%22Considered+for+release%22+cpu&visibleFields=%5B%22Title%22%2C%22Assignees%22%2C%22Status%22%2C126885563%2C130447354%2C130446877%2C130446939%2C130446997%2C133731923%2C133734297%5D&sortedBy%5Bdirection%5D=asc&sortedBy%5BcolumnId%5D=126885563&sortedBy%5Bdirection%5D=desc&sortedBy%5BcolumnId%5D=130447354) it’s [Mrunal Patel](mailto:mpatel@redhat.com)

## Oct 1, 2024

Recording: [https://www.youtube.com/watch?v=8YWCql6rLLk](https://www.youtube.com/watch?v=8YWCql6rLLk)

- KEP planning part 2  
- \[ndixita\] Pod Level Resources [\[Critical Scenarios\] Pod Level Resources](https://docs.google.com/presentation/d/1X6U81dzs_j3N0Wtu4ftJ6OU2g2pSL5_QsUTwN9sI9T4/edit?usp=sharing)  
- \[Lakshmi\] IWhen container garbage collection is deprecated? Is there any alternate recommended way for container garbage collection?  
- \[tjons\] run an initContainer only once per rollout of the deployment, not on every scheduled pod.

## Sep 24, 2024

Recording: [https://www.youtube.com/watch?v=GkPrY56\_gB4](https://www.youtube.com/watch?v=GkPrY56_gB4)

- \[jsturtevant\] Windows KEP updates for Cpu/memory manager: [https://github.com/kubernetes/enhancements/pull/4738](https://github.com/kubernetes/enhancements/pull/4738)   
- \[tallclair\] InPlacePodVerticalScaling discussion \- part 2 ([slides](https://docs.google.com/presentation/d/1vwwOeMxGPp1woJsI5rh89O8xzvrDK2VBqqMi7J2cpCM/edit#slide=id.p))  
  - KEP: [https://github.com/kubernetes/enhancements/pull/4704](https://github.com/kubernetes/enhancements/pull/4704)  
- \[johnbelamaric\] Quick PSA: Unless a strong use case comes forward, we plan to remove “classic DRA” in 1.32. See [https://github.com/kubernetes/enhancements/issues/3063\#issuecomment-2305446451](https://github.com/kubernetes/enhancements/issues/3063#issuecomment-2305446451)   
  - Reach out to [kklues@nvidia.com](mailto:kklues@nvidia.com) if you have any questions  
- \[Lakshmi\]  Requesting for review and feedback on PR  
  - [https://github.com/kubernetes/website/pull/48001](https://github.com/kubernetes/website/pull/48001)

## Sep 17, 2024

Recording: [https://www.youtube.com/watch?v=iH6KVk9B5DE](https://www.youtube.com/watch?v=iH6KVk9B5DE)

- \[pehunt\] KEP planning  
  - [Planning table](https://github.com/orgs/kubernetes/projects/186/views/7?filterQuery=status%3A%22Draft+Stage%22%2C%22Proposed+for+consideration%22+&visibleFields=%5B%22Title%22%2C%22Status%22%2C130447354%2C126885563%2C130446877%2C130446939%2C%22Reviewers%22%2C130446997%5D&sortedBy%5Bdirection%5D=desc&sortedBy%5BcolumnId%5D=130447354&sortedBy%5Bdirection%5D=asc&sortedBy%5BcolumnId%5D=126885563) 

## Sep 10, 2024

Recording: [https://www.youtube.com/watch?v=9AfQA0DYR0E](https://www.youtube.com/watch?v=9AfQA0DYR0E)

- \[pehunt\] KEP planning [KEP Board](https://github.com/orgs/kubernetes/projects/186/views/1)  
    
- \[lauralorenz\] CrashLoopBackOff KEP for 1.32 ([slides](https://docs.google.com/presentation/d/16itbKQiClbP2L7vbBCASEC5Oz6qRKxzLmcohM5_efCQ/edit?slide=id.p#slide=id.g2fb40fe0c6f_0_0) 6-10)  
- \[harche\] \- Looking for reviews [https://github.com/kubernetes/kubernetes/pull/125982](https://github.com/kubernetes/kubernetes/pull/125982)  
  - This especially affects users with high number CPUs per nodes  
- \[tallclair\] InPlacePodVerticalScaling discussion ([slides](https://docs.google.com/presentation/d/1vwwOeMxGPp1woJsI5rh89O8xzvrDK2VBqqMi7J2cpCM/edit#slide=id.p))  
  - KEP: [https://github.com/kubernetes/enhancements/pull/4704](https://github.com/kubernetes/enhancements/pull/4704)  
- \[SergeyKanzhelev\] [https://github.com/kubernetes/enhancements/issues/3386\#issue comment-2337050862](https://github.com/kubernetes/enhancements/issues/3386#issuecomment-2337050862) Do we want to remove this code for now?  
- \[T-Lakshmi\] \- Looking for feedback/answers on queries [https://github.com/kubernetes/kubernetes/issues/127157](https://github.com/kubernetes/kubernetes/issues/127157) Is container GC policy replaced with any function in evictionHard and evictionSoft policy, or its completely deprecated? What are the future plans on these container garbage collector policy?


## Sep 3, 2024

Recording: [https://www.youtube.com/watch?v=E8sw-fybnKc](https://www.youtube.com/watch?v=E8sw-fybnKc)

- \[ndixita\] Pod level resources KEP discussion  
  [\[Public\] Effective Resources & OOM Kill Behavior](https://docs.google.com/document/d/1q3UaDO5wrfP3vZFuDjpZjapYr4k_ui4SAF0s4HAIjKQ/edit?usp=sharing)  
* OOM group \-\> Pod kill  \-\> in the next iteration of KEP  
- \[lauralorenz\] CrashLoopBackOff KEP for 1.32 ([slides](https://docs.google.com/presentation/d/16itbKQiClbP2L7vbBCASEC5Oz6qRKxzLmcohM5_efCQ/edit?slide=id.p#slide=id.g2fb40fe0c6f_0_0) 6-10) \[bumped to next week but feel free to take a look at slides or discuss x-post in slack\]  
- \[sreeram-venkitesh\] Zero values for [Sleep Action of PreStop Hook](https://github.com/kubernetes/enhancements/issues/3960)  
  - [KEP-4818: Allow zero value for Sleep Action of PreStop Hook](https://docs.google.com/document/d/1o01gH2ELPY-kjwgemjg3BbWiykx37ag-4VYwwJ1jBTw/edit?usp=sharing)  
  - Draft PR to discuss changes: [https://github.com/kubernetes/kubernetes/pull/127094](https://github.com/kubernetes/kubernetes/pull/127094)       
  - Do we need to do anything particular with rollback of the feature?  
    - Probably not at least the kubelet  
- \[pranav\] Kubelet idle threads [issue](https://github.com/kubernetes/kubernetes/issues/123275)   
        \-  raised this [issue](https://github.com/golang/go/issues/68993) in golang upstream  
        \- how to control kubelet threads and memory by go runtime variables, is there any      other way to do it?  
    
- \[Kevin Hannon\] [KEP Board](https://github.com/orgs/kubernetes/projects/186/views/1)  
  - Open it up for public viewing?  
  - \[pehunt\] inspired by release team, we’ve updated the [tracking board](https://github.com/orgs/kubernetes/projects/186/views/7?filterQuery=-status%3ADone+-status%3A%22Not+for+release%22++-status%3ARemoved&visibleFields=%5B%22Title%22%2C%22Status%22%2C130447354%2C130446997%2C130446939%2C%22Reviewers%22%2C130446877%5D&sortedBy%5Bdirection%5D=desc&sortedBy%5BcolumnId%5D=Status) to have more column 

## 

## Aug 27, 2024

Recording: [https://www.youtube.com/watch?v=wGbkByo\_NBI](https://www.youtube.com/watch?v=wGbkByo_NBI)

- \[lauralorenz\] CrashLoopBackOff KEP ([slides](https://docs.google.com/presentation/d/16itbKQiClbP2L7vbBCASEC5Oz6qRKxzLmcohM5_efCQ/edit?usp=sharing))  
  - updates and changes since 1.31 \[5 minutes\]  
  - some discussion on path forward \[10-15 mins if I can get it\]  
- \[pehunt\] KEP wrangler brainstorm  
  - [SIG Node KEP Wrangler Brainstorming](https://docs.google.com/document/d/1CypSsxdowXk0PmoYLF7h5Q91MjiUPibsLFmwUVR4fKs/edit?usp=sharing)

## Aug 20, 2024

Recording: [https://www.youtube.com/watch?v=KUw2kSFsf2U](https://www.youtube.com/watch?v=KUw2kSFsf2U)

- \[vinayakankugoyal\] [https://github.com/kubernetes/enhancements/pull/4760/files\#r1699209363](https://github.com/kubernetes/enhancements/pull/4760/files#r1699209363)   
  - Break permissions into smaller buckets to allow for users to get access to things like healthz without allowing a user to get a pod to exec  
  - We currently don’t commit to supporting these endpoints, but they are being used as if we have. Should we group the endpoints by function to be less prescriptive on what a user gets access to, so we have power to change?  
  - \[Peter\] Can we break into read-only/read-write?  
    - some of the “read-only” end points can still be risky to give access to  
  - \[Dawn\] There had been talks in the past about deprecating some of the endpoints  
    - \[Tim\] We’ve been talking about doing so for so long, maybe we do this now instead of trying to find the perfect APIs  
  - \[Tim\] Maybe use healthz as the bucket?  
  - \[Sergey\] should documenting the endpoint be part of this KEP?  
  - Tim Volunteered to review/approve  
  - KEP to live under SIG-Auth  
- \[Kevin Hannon\] different OCI runtime with NodeConformance  
  - [https://github.com/kubernetes/kubernetes/issues/126639](https://github.com/kubernetes/kubernetes/issues/126639)  
  - Presubmit: [https://github.com/kubernetes/test-infra/pull/33298](https://github.com/kubernetes/test-infra/pull/33298)  
  - Periodic: [https://github.com/kubernetes/test-infra/pull/33297](https://github.com/kubernetes/test-infra/pull/33297)  
  - Maybe we should add these tests in CRI-O upstream instead of k8s–reduce overhead on upstream CI  
  - \[Kevin\] If we switch to crun by default in CRI-O, can we switch upstream k8s tests to crun as well?  
    - \[Sergey\] as long as the test failures are looked into and addressed quickly  
  - Run two versions at the same time, and then eventually switch the crun jobs to be the blocking one  
- \[SergeyKanzhelev\] Some org updates:  
  - New google groups will be used soon:  
    - [https://github.com/kubernetes/k8s.io/pull/7140](https://github.com/kubernetes/k8s.io/pull/7140)  
    - [https://github.com/kubernetes/k8s.io/pull/7160](https://github.com/kubernetes/k8s.io/pull/7160)  
  - New version of GitHub projects:  
    - [https://github.com/orgs/kubernetes/projects/151](https://github.com/orgs/kubernetes/projects/151)  
    - [https://github.com/orgs/kubernetes/projects/184](https://github.com/orgs/kubernetes/projects/184)  
    - [https://github.com/orgs/kubernetes/projects/185](https://github.com/orgs/kubernetes/projects/185)   
- \[yuanliangzhang\] Windows Node graceful shutdown   
  - KEP enhance draft:  
    [https://github.com/zylxjtu/enhancements/blob/master/keps/sig-node/2000-graceful-node-shutdown/README.md\#background-on-windows-shutdown](https://github.com/zylxjtu/enhancements/blob/master/keps/sig-node/2000-graceful-node-shutdown/README.md#background-on-windows-shutdown)  
    - POC [shoutdown poc · zylxjtu/kubernetes@854ea4b (github.com)](https://github.com/zylxjtu/kubernetes/commit/854ea4bde88c0905241b43f5f80d470967bb909f)  
    - Should we have a new KEP or keep within the other KEP  
    - Needs a reviewer from kubelet side  
      - \[Dawn\] Most of the reviewers in SIG-Node focus on linux  
    - \[Dawn\] Are there any windows version requirements?  
      - \[Lin\] I don’t think so  
      - \[Lin\] How far back do we need to support specific versions windows nodes?  
      - \[Mark\] probably windows 2019  
    - \[Mark\] We didn’t add support before because termination wasn’t working right in windows, that’s fixed now  
    - \[Peter\] [https://github.com/kubernetes/enhancements/pull/4738](https://github.com/kubernetes/enhancements/pull/4738) can be used as a baseline for KEP process  
    - \[Sergey\] If we tie this to the linux version, we may be blocked on windows to GA graceful shutdown  
      - \[Sergey\] Ideally, we GA ASAP  
      - \[Mrunal\] Instead of additional KEP, we could add another feature gate  
        - \[Sergey\] Feature gate will still block KEP graduation :-(  
    - endpoints problem: [https://github.com/kubernetes/kubernetes/issues/116965](https://github.com/kubernetes/kubernetes/issues/116965)   
- \[sotiris\] Triage decision for *Minimum CPU request is displayed when only memory request is configured*  
  - [https://github.com/kubernetes/kubernetes/issues/126195](https://github.com/kubernetes/kubernetes/issues/126195)   
- \[iholder101\]  
  - swap debugability long ongoing discussion \- asking to defer to follow-up KEPs: [https://github.com/kubernetes/kubernetes/pull/125278](https://github.com/kubernetes/kubernetes/pull/125278) and specifically [this comment](https://github.com/kubernetes/kubernetes/pull/125278#issuecomment-2268850735)  
  - [https://github.com/kubernetes/enhancements/pull/4701](https://github.com/kubernetes/enhancements/pull/4701) \- GA plans for swap (KEP-2400)  
  - Dawn to follow up offline  
- \[SergeyKanzhelev\] [https://github.com/orgs/kubernetes/projects/186/views/1](https://github.com/orgs/kubernetes/projects/186/views/1)  
  - AI:  
    - Either move to proposed for consideration  
    - Or Not for release  
-  \[torredil\] Ensure volumes are unmounted during graceful node shutdown: [https://github.com/kubernetes/kubernetes/pull/125070](https://github.com/kubernetes/kubernetes/pull/125070)  
  - Dawn/Mrunal to look and hopefully approve  
  - \[Mrunal\] Maybe add this to Clayton’s document

## Aug 13, 2024 (cancelled)

MEETING IS CANCELED TODAY due to lack of agenda and vacations

## Aug 6, 2024

Recording: [https://www.youtube.com/watch?v=K-eBDYfHiTM](https://www.youtube.com/watch?v=K-eBDYfHiTM)

- \[SergeyKanzhelev\] Files kubelet uses: [https://github.com/kubernetes/website/pull/46359/files\#r1600516887](https://github.com/kubernetes/website/pull/46359/files#r1600516887) Docs request  
  - List in comment almost? full  
  - We should document files, but recommending removal of all of them is overkill  
  - Mrunal: maybe even have a clean up command that will clean up those files.  
  - Cleaning up on startup of kubelet \- maybe we need a KEP  
  - Dawn: Kubelet should be responsible for its own files, but other files created by the plugins which might not properly cleanup, and there is no way to ensure those by Kubelet. In this case, K8s vendor is responsible for files, not kubelet.   
  - Also end users are not reporting issues back to upstream if they experience issues.  
  - Peter: if kubelet creates a file it should be responsible for deleting it. If file is owned by plugin, kubelet should be resilient to those.  
    - rphillips: Ideally, the plugin’s initialization function should handle cleanup  
- \[SergeyKanzhelev\] SergeyKanzhelev for approver: [https://github.com/kubernetes/kubernetes/pull/126551](https://github.com/kubernetes/kubernetes/pull/126551)  
- \[pehunt\] SIG Chair proposal  
- \[SergeyKanzhelev\] [SIG Node responsiveness improvements](https://docs.google.com/document/d/1AavM205LRi-RNB2xQRduzDo_ChTZ8clacKYtgsbfSAQ/edit#heading=h.tqzvzsaxo0sl)   
- \[pacoxu\] [issues/116799\#issuecomment-2249301937](https://github.com/kubernetes/kubernetes/issues/116799#issuecomment-2249301937)  
  - In [kubernetes/system-validators\#37](https://github.com/kubernetes/system-validators/pull/37), we refer to kernel long term support: [https://wiki.linuxfoundation.org/civilinfrastructureplatform/start](https://wiki.linuxfoundation.org/civilinfrastructureplatform/start) and [https://endoflife.date/linux](https://endoflife.date/linux)  
    - 4.4 & 4.19 are selected as kernel Super Long Term Support (SLTS), and the Civil Infrastructure Platform(CIP) will provide support until at least 2026\.  
  - For [cgroup v2](https://kubernetes.io/docs/concepts/architecture/cgroups/), Kubernetes recommends to use 5.8 and later, and in [runc docs](https://github.com/opencontainers/runc/blob/main/docs/cgroup-v2.md), the minimal version is 4.15 and 5.2+ is recommended.  
    - 4.5 starts support cgroup v2 io,memory & pids.(kernel 4.5 announce that cgroup v2 is not experimental)  
    - 4.15 starts support cgroup v2 cpu  
    - 4.20 PSI support & [KEP-4205](https://github.com/kubernetes/enhancements/issues/4205) is not alpha(only KEP was merged)  
    - 5.2 starts support cgroup v2 freezer  
    - 5.8: [Adding root](https://github.com/kubernetes/kubernetes/issues/103759#issuecomment-926024150) \`cpu.stat\` [file on cgroupv2](https://github.com/kubernetes/kubernetes/issues/103759#issuecomment-926024150) was only added in 5.8.

## July 30, 2024

Recording: [https://www.youtube.com/watch?v=JGYTQbs6eJk](https://www.youtube.com/watch?v=JGYTQbs6eJk)

- \[Peter Hunt\] Retrospective from 1.31 release  
  - [SIG Node 1.31 retro](https://docs.google.com/document/d/16Ek41L3ocMDmeJTDTp108PBULaZ6ffEWU8aoEQ8q-DU/edit)   
  - Previous retrospectives:  
    - There were no retro for 1.29 and 1.30  
    - [SIG Node 1.28 retro](https://docs.google.com/document/d/1NaT0rY0o1cNdTxIlgZ5m0TqLDI7AfYn3rBAAl4qT1Bw/edit#heading=h.vn5jbyiup6d)  
    - [SIG Node 1.27 retro](https://docs.google.com/document/d/1DxJH1w_lrEOfflR-TED1vjc0ZYIXO0aDb3vJXPMKCdY/edit#heading=h.99i0ua4v77ap)

## July 23, 2024

Recording: [https://www.youtube.com/watch?v=Wc7yrCLILK8](https://www.youtube.com/watch?v=Wc7yrCLILK8)

- \[fromani\]\[on behalf of sphrasavath\] resuming work on [KEP 2621](https://github.com/kubernetes/enhancements/issues/2621): Enhance CPU manager with L3 cache aware  
  - pivot from new cpumanager policy to new cpumanager policy option  
  - revised design doc (comment from the enh issue: [https://docs.google.com/document/d/1LpnMjGNsQyHOuVHMktIrjZsdRw9aKZ8djt354nAno6M/edit?usp=sharing](https://docs.google.com/document/d/1LpnMjGNsQyHOuVHMktIrjZsdRw9aKZ8djt354nAno6M/edit?usp=sharing) )  
- \[Sunnat\] On behalf of Marsik. do not set CPU quota for guaranteed pods  
  - [https://github.com/kubernetes/kubernetes/pull/117030](https://github.com/kubernetes/kubernetes/pull/117030)   
- \[pehunt\]: ProcMount disabled, or UserNamespaces enabled?  
  - [https://github.com/kubernetes/kubernetes/pull/126291](https://github.com/kubernetes/kubernetes/pull/126291) 

## July 16, 2024

Recording: [https://www.youtube.com/watch?v=0iPCt\_FZxSk](https://www.youtube.com/watch?v=0iPCt_FZxSk)

- \[dawnchen\] FYI: [\[PUBLIC\] Kubernetes: Disrupted pods should be eagerly removed from endpoints](https://docs.google.com/document/d/1t25jgO_-LRHhjRXf4KJ5xY_t8BZYdapv7MDAxVGY6R8/edit#heading=h.oaga8h3rwgtl)   
  - Primary concern raised so far by Rob Scott is the risk that someone interprets EndpointSlice terminating as one way  
  - More discussion of alternatives

## July 9, 2024

Recording: [https://www.youtube.com/watch?v=RTEtVbZPB-E](https://www.youtube.com/watch?v=RTEtVbZPB-E)

- \[case\] A group of us were working on a PR [around](https://github.com/kubernetes/kubernetes/issues/40610) adding node labels to the downward API. [KEP-4742](https://github.com/kubernetes/enhancements/pull/4747)  
- \[harche\] \- Are we calculating the system reserved cpu shares correctly? [https://github.com/kubernetes/kubernetes/issues/72881\#issuecomment-821224980](https://github.com/kubernetes/kubernetes/issues/72881#issuecomment-821224980)  
  - Analysis with various CPU cores \- [System reservation cpu](https://docs.google.com/spreadsheets/d/1N8Xkzu7ArZYKTP0Ob9vMxUnlqAFlsaFoNGbxyAhzzxU/edit?usp=sharing)   
  - \[Derek\] found relevant node allocatable designs [https://github.com/kubernetes/design-proposals-archive/blob/main/node/kubelet-systemd.md](https://github.com/kubernetes/design-proposals-archive/blob/main/node/kubelet-systemd.md)   and [https://github.com/kubernetes/design-proposals-archive/blob/main/node/node-allocatable.md](https://github.com/kubernetes/design-proposals-archive/blob/main/node/node-allocatable.md)   
  -   
- \[adil\] I have a question regarding logging, is there a way to disable all logs from k8s components and only get error logs? I tried setting different verbosities but it didn't help much. If there is no way to do it right now, is this something would interested in implementing? The reason why we want this is to optimize the CPU usage.  
  - pehunt: For all the kubernetes components, you should be able to set the `-v` flag which sets the verbosity of klog. You need to individually set this flag for each kubernetes component, I don’t think there’s a centralized place you can do this today. If you set `-v=1` you should only get the most urgent messages  
- \[mimowo\] looking for sig-node reviewers for [Fix that PodIP field is temporarily removed for a terminal pod](https://github.com/kubernetes/kubernetes/pull/125404); heads up for the Kubelet issue that it may flip phase from Succeeded to Failed [link](https://github.com/kubernetes/kubernetes/issues/125410)  
- \[MaRosset\] [Request for review of Windows memory pressure eviction PR](https://github.com/kubernetes/kubernetes/pull/122922)  
- 

## July 2, 2024 \[Canceled for July 4th week\]

## June 25, 2024

Recording: [https://www.youtube.com/watch?v=ExmOu9Twp3A](https://www.youtube.com/watch?v=ExmOu9Twp3A) 

- \[Filip Krepinsky\] Creation of a new WG  
  - discussed in [https://groups.google.com/g/kubernetes-sig-architecture/c/Tb\_3oDMAHrg/m/pJjl6v4mAgAJ](https://groups.google.com/g/kubernetes-sig-architecture/c/Tb_3oDMAHrg/m/pJjl6v4mAgAJ)   
  - Clarify scope: Node vs group of Node, SIG Node vs k8s level, list of problems/scope  
- \[Pranav Pandey\] Kubelet not releasing idle threads   
         \-  discussed  [here](https://github.com/kubernetes/kubernetes/issues/123275)       
         \-  I think this issue is due to golang, could we confirm this?  
         \-  Could we also confirm if there is a direct way for the kubelet to set the  
            maximum thread number by any parameter or something like that?  
- \[lubomir\] review my small PR that makes a windows/kubelet related change:  
  - [https://github.com/kubernetes/kubernetes/pull/123137](https://github.com/kubernetes/kubernetes/pull/123137)  
  - warn instead of error for unsupported options on Windows  
  - we don't need to exit the kubelet with an error on Windows just because the user is using a config that works on Linux.  
  - old PR where we discussed we should not have different defaults on Windows:  
    - [https://github.com/kubernetes/kubernetes/pull/77710](https://github.com/kubernetes/kubernetes/pull/77710)

## June 18, 2024

Recording: [https://www.youtube.com/watch?v=REmtlcXma\_M](https://www.youtube.com/watch?v=REmtlcXma_M) 

- \[Sergey\] KEPs list for 1.31: [https://github.com/orgs/kubernetes/projects/183/views/1?filterQuery=sig%3Asig-node\&groupedBy%5BcolumnId%5D=Status\&sortedBy%5Bdirection%5D=desc\&sortedBy%5BcolumnId%5D=Status\&sliceBy%5BcolumnId%5D=Status](https://github.com/orgs/kubernetes/projects/183/views/1?filterQuery=sig%3Asig-node&groupedBy%5BcolumnId%5D=Status&sortedBy%5Bdirection%5D=desc&sortedBy%5BcolumnId%5D=Status&sliceBy%5BcolumnId%5D=Status)   
- \[Dixita\] Support for KEP exception until Friday, June 21  
  - [https://github.com/kubernetes/enhancements/issues/2837](https://github.com/kubernetes/enhancements/issues/2837)   
  - KEP needs to address the following suggestions by Tim Hockin  
    - Default values when one of requests/limits is not set at pod level  
    - Change language for QoS definitions  
    - Stating OOM Kill behavior change  
  - Reasoning  
    - Feature discussions since March 2020  
    - The more we delay this feature, it becomes difficult to support new features being added in every release.  
    - Low risk: Alpha phase targets only adding the new fields in the PodSpec so that feature development can start.  
    - Important to unblock AI model use cases  
- \[mimowo\] looking for sig-node reviews for [Fix that PodIP field is temporarily removed for a terminal pod](https://github.com/kubernetes/kubernetes/pull/125404)

## June 11, 2024

Recording: [https://www.youtube.com/watch?v=A1XwOJxBL0c](https://www.youtube.com/watch?v=A1XwOJxBL0c)

- \[tallclair\] [\#125393](https://github.com/kubernetes/kubernetes/issues/125393) Should we remove soft admission failure, before AppArmor goes GA?  
- \[Filip Krepinsky\] Latest NodeMaintenance discussions  
  - [https://github.com/kubernetes/enhancements/pull/4213](https://github.com/kubernetes/enhancements/pull/4213)   
- \[Sotiris/esotsal\] Static CPU management policy alongside InPlacePodVerticalScaling   
  - [Status / next steps / open questions one slider](https://docs.google.com/presentation/d/1jm80y9rCvjV3P6a5LTQxYv8R5er9Zw705QxANwwLaMg/edit?usp=sharing)

     \-      \[vaibhav\] Eviction manager should check the disk usage of dead containers  
 	     \-      [https://github.com/kubernetes/kubernetes/issues/115201](https://github.com/kubernetes/kubernetes/issues/115201)  
                \-       Default values of Kubelet’s eviction hard parameters  
                \-       [https://github.com/kubernetes/kubernetes/issues/119985](https://github.com/kubernetes/kubernetes/issues/119985)

- \[pehunt\] [https://github.com/kubernetes/kubernetes/pull/124285](https://github.com/kubernetes/kubernetes/pull/124285) need KEP?  
- \[harche\] \- [https://github.com/kubernetes/kubernetes/pull/125341](https://github.com/kubernetes/kubernetes/pull/125341) \- changing the secret fetching strategy while creating the pod.   
- \[pehunt\] sync about [https://github.com/kubernetes/enhancements/pull/4693](https://github.com/kubernetes/enhancements/pull/4693) updates  
  - [https://github.com/kubernetes/enhancements/pull/4693\#discussion\_r1630238957](https://github.com/kubernetes/enhancements/pull/4693#discussion_r1630238957) How do we feel about the Never handling change?

## June 4, 2024

Recording: [https://www.youtube.com/watch?v=3dyVRBR7K7k](https://www.youtube.com/watch?v=3dyVRBR7K7k)

- \[SergeyKanzhelev\]   
  KEPs for 1.31: [https://github.com/kubernetes/enhancements/issues?q=is%3Aissue+is%3Aopen+label%3Asig%2Fnode+milestone%3Av1.31+](https://github.com/kubernetes/enhancements/issues?q=is%3Aissue+is%3Aopen+label%3Asig%2Fnode+milestone%3Av1.31+)   
    
  Missing lead-opted-in: [https://github.com/kubernetes/enhancements/issues?q=is%3Aissue+is%3Aopen+label%3Asig%2Fnode+milestone%3Av1.31+-label%3Alead-opted-in](https://github.com/kubernetes/enhancements/issues?q=is%3Aissue+is%3Aopen+label%3Asig%2Fnode+milestone%3Av1.31+-label%3Alead-opted-in)  
    
- \[chrismuellner\] discuss loose linux capability handling in security context: [https://github.com/kubernetes/kubernetes/issues/119569\#issuecomment-2020382413](https://github.com/kubernetes/kubernetes/issues/119569#issuecomment-2020382413)  
  - varying, incomplete implementations for validations  
  - [documentation](https://kubernetes.io/docs/tasks/configure-pod-container/security-context/#set-capabilities-for-a-container) inaccurate: \`CAP\_\` prefix allowed? upper/lower case?  
- \[pehunt\] feedback on whether to exclude critical pods in swap  
  - [https://github.com/kubernetes/kubernetes/pull/125277](https://github.com/kubernetes/kubernetes/pull/125277)   
- \[ndixita\] Pod level resource spec KEP: [https://github.com/kubernetes/enhancements/pull/4678](https://github.com/kubernetes/enhancements/pull/4678)  
- \[Filip Krepinsky\] update on the Declarative NodeMaintenance and Evacuation API KEPs:  
  - [https://github.com/kubernetes/enhancements/pull/4213](https://github.com/kubernetes/enhancements/pull/4213)  
  - [https://github.com/kubernetes/enhancements/pull/4565](https://github.com/kubernetes/enhancements/pull/4565)   
- \[lauralorenz\] CrashLoopBackoff KEP  
  - [https://github.com/kubernetes/enhancements/pull/4604](https://github.com/kubernetes/enhancements/pull/4604)   
- \[SergeyKanzhelev\] Many flakes reported by release team

## May 28, 2024

Recording: [https://www.youtube.com/watch?v=RDWC4rtQOCo](https://www.youtube.com/watch?v=RDWC4rtQOCo)

- KEP freeze is coming (schedule: [https://www.kubernetes.dev/resources/release/](https://www.kubernetes.dev/resources/release/)). [https://github.com/kubernetes/enhancements/issues?q=is%3Aissue+is%3Aopen+label%3Asig%2Fnode+milestone%3Av1.31+](https://github.com/kubernetes/enhancements/issues?q=is%3Aissue+is%3Aopen+label%3Asig%2Fnode+milestone%3Av1.31+)   
- \[JeffLuoo\] Pod full startup latency metrics to record pod from creation to ready: [https://github.com/kubernetes/kubernetes/issues/124892](https://github.com/kubernetes/kubernetes/issues/124892)  
- \[dawnchen\] Starting DRA driver for GPU in CNCF / K8s repo  
  - \[Ed\] [https://github.com/kubernetes-sigs/dra-example-driver](https://github.com/kubernetes-sigs/dra-example-driver)  
  - \[John\] from distributors perspective \- driver from community would be preferable comparing to vendor-managed.  
  - We talking about allowing space for vendors, if they want/prefer.  
  - Idea is to simplify the life of distributors to have a place to take drivers from.  
- \[pehunt\] [https://github.com/kubernetes/kubernetes/pull/125038](https://github.com/kubernetes/kubernetes/pull/125038)   
- \[harche\] cgroup v1 maintenance mode KEP \- Should we feature gate it or not? [https://github.com/kubernetes/enhancements/pull/4572\#discussion\_r1608362477](https://github.com/kubernetes/enhancements/pull/4572#discussion_r1608362477) 

## May 21, 2024

Recording: [https://www.youtube.com/watch?v=eSYWzusEZiA](https://www.youtube.com/watch?v=eSYWzusEZiA)

- \[iholder101\]:   
  - \#[123963](https://github.com/kubernetes/kubernetes/pull/123963): Add swap to kubectl describe node's output  
    - On the one hand we [received feedback](https://github.com/kubernetes/enhancements/pull/4401#discussion_r1479124963) regarding making it easier to debug and monitor swap. On the other hand there’s a pushback regarding exposing it through API. What’s the right balance here?  
  - timezone poll results from two weeks ago: [https://ibb.co/z8R3nXN](https://ibb.co/z8R3nXN).  
    - SIG-Node leadership: does moving back two hours make sense? What is the process to formalize that change?  
- \[sallyom\]:   
  - [KEP 4639: OCI VolumeSource](https://github.com/kubernetes/enhancements/pull/4642)   
    - alternatives  
      - [https://github.com/kubernetes-retired/csi-driver-image-populator](https://github.com/kubernetes-retired/csi-driver-image-populator)  
      - [https://github.com/warm-metal/container-image-csi-driver](https://github.com/warm-metal/container-image-csi-driver)   
      - [https://github.com/kubernetes/enhancements/tree/master/keps/sig-storage/1495-volume-populators](https://github.com/kubernetes/enhancements/tree/master/keps/sig-storage/1495-volume-populators)   
- \[pehunt\]: [https://github.com/kubernetes/kubernetes/issues/124333](https://github.com/kubernetes/kubernetes/issues/124333) 	  
  - compelling case between balancing cluster admin configuration and workloads being punished for them

      \-    \[vaibhav\] Eviction manager should check the disk usage of dead containers  
                       \-  [https://github.com/kubernetes/kubernetes/issues/115201](https://github.com/kubernetes/kubernetes/issues/115201)  
\- [https://github.com/kubernetes/enhancements/issues/4341](https://github.com/kubernetes/enhancements/issues/4341) 

## May 14, 2024

No agenda, canceling this week. 

## May 7, 2024

Recording: [https://www.youtube.com/watch?v=\_FPa0TVPoY4](https://www.youtube.com/watch?v=_FPa0TVPoY4)

- \[marquiz/zvonkok\] [KEP-4112: Pass down resources to CRI](https://github.com/kubernetes/enhancements/issues/4112) follow-up  
  - [https://docs.google.com/presentation/d/13TDKyASpMfDrVBSRj4JiU6gFeChx0ws4DTenBN1qUnA/edit?usp=sharing](https://docs.google.com/presentation/d/13TDKyASpMfDrVBSRj4JiU6gFeChx0ws4DTenBN1qUnA/edit?usp=sharing)  
- \[yujuhong\] cgroup v2 memory usage – bug or working as intended?  
  - [https://github.com/kubernetes/kubernetes/issues/118916](https://github.com/kubernetes/kubernetes/issues/118916)  
  - and discussion in runc \- [https://github.com/opencontainers/runc/pull/3933\#issuecomment-1833599870](https://github.com/opencontainers/runc/pull/3933#issuecomment-1833599870)   
- 

## Apr 30, 2024

Recording: [https://www.youtube.com/watch?v=iuZCxtAeoQ8](https://www.youtube.com/watch?v=iuZCxtAeoQ8)

- \[SergeyKanzhelev\] Annual report last call for comments: [https://github.com/kubernetes/community/pull/7831/](https://github.com/kubernetes/community/pull/7831/files)  
- \[lauralorenz\] intro on proposed changes to CrashLoopBackoff ([slides](https://docs.google.com/presentation/d/16itbKQiClbP2L7vbBCASEC5Oz6qRKxzLmcohM5_efCQ/edit?slide=id.p#slide=id.p)), this is from [Kubernetes\#57291](https://github.com/kubernetes/kubernetes/issues/57291)  
- \[iholder101/Peter Hunt\] \#[124060](https://github.com/kubernetes/kubernetes/pull/124060): Avoid swapping memory-backed volumes with tmpfs’ “[noswap](https://www.kernel.org/doc/html/latest/filesystems/tmpfs.html)” option.  
  - How to behave if the option is not supported?  
  - If it is not supported, do we want to fallback to ramfs / BRD / zswap?  
  - How should it be tested, since the CI runs with an old kernel (5.15 \< 6.4)   
  - Update KEP and issue with the current state  
- \[iholder101/Peter Hunt\]: In my time-zone this meeting takes place at 20:00 PM. Is it acceptable to reschedule this meeting for an earlier time? This might significantly help people from the EMEA region to join.  
  - Defer to next week, hope for more consensus  
  - in the meantime, ask the sig-node mailing list who would be able to make it that previously cannot  
- \[ndixita\]  
  \- kubelet archived logs permissions [https://github.com/kubernetes/kubernetes/pull/124229](https://github.com/kubernetes/kubernetes/pull/124229)   
  Solution: 1\) Config options for users maybe [https://github.com/kubernetes/kubernetes/issues/124228\#issuecomment-2042885888](https://github.com/kubernetes/kubernetes/issues/124228#issuecomment-2042885888)   
  Have a feature gate that is removed later.  
  Sergey: same issue with termination logs. [https://github.com/kubernetes/kubernetes/pull/108076](https://github.com/kubernetes/kubernetes/pull/108076)   
  \- cadvisor enumerates memory and hugepages separately    
  Issue: [https://github.com/kubernetes/kubernetes/issues/84426](https://github.com/kubernetes/kubernetes/issues/84426)   
  [https://github.com/kubernetes/kubernetes/pull/119173/files\#r1307246832](https://github.com/kubernetes/kubernetes/pull/119173/files#r1307246832)   
- Can we know if this option is planned to be backported, and to which version?

           Recommended solution: fix in cadvisor, and assess backward compatibility (probably add a new field)

- Question: How will the behavior be if huge pages are changed dynamically?

 \- \[Peter Hunt\] Finish KEP Planning

## Apr 23, 2024

Recording: [https://www.youtube.com/watch?v=-TEdQvF7kUE](https://www.youtube.com/watch?v=-TEdQvF7kUE)

- \[SergeyKanzhelev\] Annual report draft: [https://github.com/kubernetes/community/pull/7831](https://github.com/kubernetes/community/pull/7831) Please add your comment and review the list of KEPs ([https://github.com/kubernetes/community/issues/7777\#issuecomment-2067917685](https://github.com/kubernetes/community/issues/7777#issuecomment-2067917685))  
    
- \[anishshah\] \- v1.30 release report  
- 	\- [github.com/AnishShah/sig-node-flaky-tescontainerd/containerdts/tree/main](https://github.com/AnishShah/sig-node-flaky-tests/tree/main)  
  - \~10% release blocking tests are flaky  
- \[jstur\] Follow up on UsageNanoCores CRI [https://github.com/kubernetes/kubernetes/issues/122092\#issuecomment-1956783842](https://github.com/kubernetes/kubernetes/issues/122092#issuecomment-1956783842)   
  - What is the best approach?  
    - implemented cri background implementation in [https://github.com/containerd/containerd/pull/10010](https://github.com/containerd/containerd/pull/10010)   
  - Additional questions if cri is responsible:   
    - costs of having 10s heart beat on CRI side?  
    - what does it mean to have it 10s behind other stats?  
    - backwards compat?  
  - James+Peter+Mike to have a call to sync on this

      \-     \[vaibhav\] Eviction manager should check the disk usage of dead containers  
                       \-  [https://github.com/kubernetes/kubernetes/issues/115201](https://github.com/kubernetes/kubernetes/issues/115201)  
\- [https://github.com/kubernetes/enhancements/issues/4341](https://github.com/kubernetes/enhancements/issues/4341) 

## Apr 16, 2024

Recording: [https://www.youtube.com/watch?v=vjcRUX\_vSbU](https://www.youtube.com/watch?v=vjcRUX_vSbU)

- \[pehunt\] KEPs planning [https://docs.google.com/document/d/1U10J0WwgWXkdYrqWGGvO8iH2HKeerQAlygnqgDgWv4E/edit](https://docs.google.com/document/d/1U10J0WwgWXkdYrqWGGvO8iH2HKeerQAlygnqgDgWv4E/edit)   
- \[pehunt\]: [https://github.com/kubernetes/org/issues/4805](https://github.com/kubernetes/org/issues/4805)   
  - Mostly looking for feedback  
  - Some questions/replies are here looking for more opinions: [https://github.com/kubernetes/org/issues/4805\#issuecomment-1985215796](https://github.com/kubernetes/org/issues/4805#issuecomment-1985215796)   
- \[iholder101/pehunt\]: \#[123963](https://github.com/kubernetes/kubernetes/pull/123963): Add swap to kubectl describe node's output  
  - On the one hand we [received feedback](https://github.com/kubernetes/enhancements/pull/4401#discussion_r1479124963) regarding making it easier to debug and monitor swap. On the other hand there’s a pushback regarding exposing it through API. What’s the right balance here?  
- \[marquiz/zvonkok\] [KEP-4112: Pass down resources to CRI](https://github.com/kubernetes/enhancements/issues/4112) follow-up  
- \[iholder101/pehunt\]: timezone poll results from two weeks ago: [https://ibb.co/z8R3nXN](https://ibb.co/z8R3nXN).  
  - SIG-Node leadership: does moving back two hours make sense? What is the process to formalize that change?

      \-    \[harche\] \- cgroup v1 support \- Deprecation only or Removal as well?  
	      \-	[https://github.com/kubernetes/enhancements/issues/4569](https://github.com/kubernetes/enhancements/issues/4569)	   
      ~~\-    \[klueska\] \- KEP update for DRA to match 1.30 implementation~~  
	      ~~\-	[https://github.com/kubernetes/enhancements/pull/4561](https://github.com/kubernetes/enhancements/pull/4561)~~  
	      ~~\-	[dawnchen@google.com](mailto:dawnchen@google.com)to approve~~  
      \-     \[anishshah\] \- v1.30 release report  
	\- [github.com/AnishShah/sig-node-flaky-tests/tree/main](https://github.com/AnishShah/sig-node-flaky-tests/tree/main)  
           \-  22/249 sig-node release blocking tests are flaky.  
     

## Apr 9, 2024

Recording: [https://www.youtube.com/watch?v=o3AohYi9aQA](https://www.youtube.com/watch?v=o3AohYi9aQA)

- \[oshestopalova\]: [Soft eviction of pods with long grace periods blocks hard evictions when under resource pressure](https://github.com/kubernetes/kubernetes/issues/123872)   
- \[iholder101/pehunt\]: timezone poll results from last week: [https://ibb.co/z8R3nXN](https://ibb.co/z8R3nXN).  
- \[jkyros\] Trying to use InPlacePodVerticalScaling in [Vertical Pod Autoscaler](https://github.com/kubernetes/autoscaler/pull/6652)   
  - does anyone remember why limits are [required for in-place scaling](https://github.com/openshift/kubernetes/blob/258f1d5fb6491ba65fd8201c827e179432430627/pkg/kubelet/kuberuntime/kuberuntime_manager.go#L556)?  
  - naively something like [this](https://github.com/kubernetes/kubernetes/compare/master...jkyros:kubernetes:minimal-patch-fix-inplacepodverticalscaling-limits) fixes it, but probably has consequences  
  - \[Sotiris Salloumis\] Perhaps we can discuss this in [https://kubernetes.slack.com/archives/C06FSK01BGU](https://kubernetes.slack.com/archives/C06FSK01BGU)  ?  
- \[pehunt/eddiezane\]: kubectl cp improvements  
- \[Sonemaly\]: Start discussion around A[ddressing Noisy Neighbor/Split L3 Cache Topology](https://groups.google.com/g/kubernetes-sig-node/c/V1RjCDKcTaY)  
  - \[kad\]: please share in continuation in the mail thread scenarios that you have and corner cases that you found are not solved today. We need to look how it could be done in a way where all other vendors (especially on ARM side where assumptions on presence of L3 might be not true) will not be affected on proposed changes to static policy. At the moment, the cache layout is partially buggy in cAdvisor library that detects it, and components like CPU manager is not consuming it at all from MachineInfo.   
- \[Matt Karrmann\] [Configure group OOM Kills at the container level instead of the kubelet level](https://github.com/kubernetes/kubernetes/pull/122813#issuecomment-2004527657)  
  - Follow up with an issue to chat about different cases for pod vs kubelet level configuration

## Apr 2, 2024

Recording: [https://www.youtube.com/watch?v=Ho1kn-1p8Cg](https://www.youtube.com/watch?v=Ho1kn-1p8Cg)

- \[Sotiris\] InPlacePodVerticalScaling moving forward to beta (todo/needs/planing)  
  - From Jiaxin Shan:  
    - We worked on issues [https://docs.google.com/document/d/1V3DLh3pH3CD-xhhJvAnOq\_oWgPyjO-vj6wY6qdew9H0/edit\#heading=h.ybybfdfputt](https://docs.google.com/document/d/1V3DLh3pH3CD-xhhJvAnOq_oWgPyjO-vj6wY6qdew9H0/edit#heading=h.ybybfdfputt) and most of the issues have been solved or have pending PRs. But this is definitely a subset of the working items moving to beta.   
    - for people need more context.  
    - [https://github.com/kubernetes/kubernetes/issues/109547](https://github.com/kubernetes/kubernetes/issues/109547)   
  - \[Dixi\] a lot of interest. Maybe we need to meet separately to split tasks?  
  - \[mrunal\] there is a slack channel already. Is it ok to coordinate there?  
  - \[Dixi\] slack may work.   
  - \[Jiaxin\] Let’s work together in that channel.  
  - \[SergeyKanzhelev\] Please review API again, Many use cases for the feature EXPECT to use this feature differently than the KEP’s API was designed.	  
- \[Jiaxin\] InPlaceVPA performance issue. A few users in the community requested the patch [https://github.com/kubernetes/kubernetes/pull/123941](https://github.com/kubernetes/kubernetes/pull/123941). PLEG cycle doesn’t take inplace pod status into consideration and never emit update events.  
- \[matthyx\] (from sidecar WG) postStart hook prevents normal container termination \- how to fix that?  
  - [PUBLIC: Trying to diagram pod lifecycle stuff](https://docs.google.com/presentation/d/1e-qJWe6He2qjt0PFiZtocqP8VbkrwKAI9B4X72L0bGk/edit?usp=sharing&resourcekey=0-N77W7Q5UHuN5dFtekjmAIA)(slide 3\)  
  - [https://github.com/kubernetes/kubernetes/blob/ec301a5cc76f48cdadc77bcfbd686cf40b124ecf/pkg/kubelet/kuberuntime/kuberuntime\_container.go\#L297](https://github.com/kubernetes/kubernetes/blob/ec301a5cc76f48cdadc77bcfbd686cf40b124ecf/pkg/kubelet/kuberuntime/kuberuntime_container.go#L297)  
    - [https://github.com/kubernetes/kubernetes/pull/113883](https://github.com/kubernetes/kubernetes/pull/113883) (check for e2e coverage in PR)  
    - we cover this in our [KEP](https://github.com/kubernetes/enhancements/issues/4438) (to be renamed)  
- \[pranav\]: Could we implement a feature in Kubelet to limit the number of threads to the number of CPUs available?  
  - [https://github.com/kubernetes/kubernetes/issues/123275](https://github.com/kubernetes/kubernetes/issues/123275)   
- \[SergeyKanzhelev\] WG Serving proposal: [https://groups.google.com/g/kubernetes-sig-node/c/KGfkpVmNrNc](https://groups.google.com/g/kubernetes-sig-node/c/KGfkpVmNrNc)   
-   
- \[Anish\] [https://github.com/kubernetes/kubernetes/pull/123782](https://github.com/kubernetes/kubernetes/pull/123782) (ask is for a review).  
  - Issue: Container status changes to ContainerStatusUnknown when evicted due to exceeding ephemeral storage limit.  
  - Root Cause: There is a race condition which is removing the container before the container status update.  
  - Fix: The fix is to check that the pod is finished before cleaning up. added a check to the existing e2e test.  
- \[iholder101\]: In my time-zone this meeting takes place at 20:00 PM. Is it acceptable to reschedule this meeting for an earlier time? This might significantly help people from the EMEA region to join.  
- 

## Mar 26, 2024

Recording: [https://www.youtube.com/watch?v=5TLp233Bisg](https://www.youtube.com/watch?v=5TLp233Bisg)

- \[tallclair\] [Deprecate & remove Kubelet RunOnce mode](https://github.com/kubernetes/kubernetes/issues/124030)  
-     Mark deprecated in 1.31 and remove in 1.33  
-     Add KEP   
- \[Sergey\] cgroupv1 removal/deprecation is moving to 1.31  
-    Harshal to open a KEP for 1.31  
- \[kannon92\] CAdvisor bug on pid stats  
  - [https://github.com/google/cadvisor/pull/3497/files](https://github.com/google/cadvisor/pull/3497/files)  
  - K8s: [https://github.com/kubernetes/kubernetes/pull/123914](https://github.com/kubernetes/kubernetes/pull/123914)

     \-  \[Dawn\]  Kubecon recap. Slide deck: [Sig Node Intro and Deep Dive](https://docs.google.com/presentation/d/1xOglu8Pfq8TNLp_ehMj7S-L56znOQ-aksiaHhRtULyY/edit?usp=sharing)  
	\- Unconference hw resource model discussion:   
[\[PUBLIC\] 2024 KubeCon EU - Contrib Summit Unconference](https://docs.google.com/presentation/d/1LIBx8xWR6uelGM38QgcZ9MloRLL3od20z3JuCXBa3_s/edit?usp=sharing)

## Mar 19, 2024 \[Canceled for KubeCon\]

## Mar 12, 2024

Recording: [https://www.youtube.com/watch?v=-435mh2GyGU](https://www.youtube.com/watch?v=-435mh2GyGU) 

- \[Kevin Hannon\] Upper limit on ImagePullBackOff and fail the pod  
  - [https://github.com/kubernetes/kubernetes/issues/122300](https://github.com/kubernetes/kubernetes/issues/122300)  
- \[Kevin Hannon\] Flakiness in eviction tests  
  - [https://github.com/kubernetes/kubernetes/issues/123591](https://github.com/kubernetes/kubernetes/issues/123591)  
    - Stats eviction if stats api failure  
    - PIDStats Fix- [https://github.com/kubernetes/kubernetes/pull/123369](https://github.com/kubernetes/kubernetes/pull/123369)  
- \[Krzysztof Wilczyński\] Current state and the future of the Graceful Node Shutdown support in kubelet.  
  - [KEP 2000](https://github.com/kubernetes/enhancements/blob/master/keps/sig-node/2000-graceful-node-shutdown/README.md): Graceful Node Shutdown  
- ~~\[Anish\] ContainerStatusUnknown after ephemeral storage limit is exceeded~~  
  -  ~~[https://github.com/kubernetes/kubernetes/issues/122160](https://github.com/kubernetes/kubernetes/issues/122160)~~  
  -   
- \[Hongxiang Jiang\] Calculate oom\_score\_adj in a CPU-agnostic way, taking in consideration Pod Priority too  
  - [https://github.com/kubernetes/kubernetes/issues/78848](https://github.com/kubernetes/kubernetes/issues/78848)  
- 

## Mar 5, 2024

Recording: [https://www.youtube.com/Kubernetes SIG Node 20240305watch?v=yBmVPBO9Y9Y](https://www.youtube.com/watch?v=yBmVPBO9Y9Y)

- \[Sotiris Salloumis\] Static CPU management policy along side InPlacePodVerticalScaling  
  - [https://github.com/kubernetes/kubernetes/pull/123319](https://github.com/kubernetes/kubernetes/pull/123319)  
  - Is KEP needed? (this PR is an attempt to fix KEP 1287 Alpha Feature Code Issue [\#10](https://github.com/kubernetes/enhancements/issues/1287#issuecomment-1972964844))  
  - [Demo of latest patch](https://drive.google.com/file/d/1WVmDK668OeaBV2a1MZLmZJFR52innQ9M/view?usp=sharing)

	PTAL: [Inplace VPA + core binding](https://docs.google.com/document/d/1V3DLh3pH3CD-xhhJvAnOq_oWgPyjO-vj6wY6qdew9H0/edit?usp=sharing) There’s some discussion about VPA \+ CPU manager static policy

- \[Dixita, Anish\] Seeking help for bug prioritization and triage for K8s 1.30 release on Wednesday 10AM PST.  
- \[pehunt\] proc mount PR separate from e2e tests [https://github.com/kubernetes/kubernetes/pull/123520](https://github.com/kubernetes/kubernetes/pull/123520)   
- \[Kevin Hannon\] CRIO tests failing as of today   
  - [https://github.com/kubernetes/kubernetes/issues/123715](https://github.com/kubernetes/kubernetes/issues/123715)  
  - pehunt opened [https://github.com/kubernetes/kubernetes/pull/123726](https://github.com/kubernetes/kubernetes/pull/123726)  
  - 

##  Feb 27, 2024

Recording: [https://www.youtube.com/watch?v=3IRepUPQ0CU](https://www.youtube.com/watch?v=3IRepUPQ0CU)

- \[dwestbrook\] Discuss Per Pod Container Updates (i.e. similar to this [issue](https://github.com/kubernetes/kubernetes/issues/110487#issuecomment-1153639507))  
  - [Feature Request – Per Pod Container Updates](https://docs.google.com/document/d/1RyECD6xlIDejcpi818WS1pWz9KjX43W4t9MF1LUPRLY/edit) (request access)  
- \[chrishenzie\] [Extending containerd 1.X EOL](https://github.com/containerd/containerd/issues/9866) to align with K8s EOL  
  - 1.6 and 1.7 have parallel LTS windows  
  - Will run until next LTS release, which release TBD (could be v2.0, v2.1)  
  - containerd v2.0 contains migration tools/scripts to assist with users of deprecated features  
    - containerd \-c pathToToml [config migrate](https://github.com/containerd/containerd/blob/v2.0.0-beta.2/cmd/containerd/command/config.go#L107)  
    - [https://github.com/containerd/containerd/blob/main/RELEASES.md\#daemon-configuration](https://github.com/containerd/containerd/blob/main/RELEASES.md#daemon-configuration)  
    - containerd has moved packages around in the 2.0 refactoring see move script [https://github.com/containerd/containerd/pull/9365](https://github.com/containerd/containerd/pull/9365) this should aid people involved in containerd plugin development and importing the various packages..  
- \[SergeyKanzhelev\] Sidecar WG \- new time for the meeting: Seattle 2PM, Paris 11PM, Seoul 7AM (6AM) (Wednesday)

## Feb 20, 2024

Recording: [https://www.youtube.com/watch?v=vEbpXkhm73M](https://www.youtube.com/watch?v=vEbpXkhm73M)

- \[Kevin Hannon\] Discuss configuration for pod logs location  
  - PR: [https://github.com/kubernetes/kubernetes/pull/112957](https://github.com/kubernetes/kubernetes/pull/112957)  
  - issue: [https://github.com/kubernetes/kubernetes/issues/98473](https://github.com/kubernetes/kubernetes/issues/98473)  
  - Is KEP needed?  
  - Security implications of logs locations  
  - Impact on disk usage  
  - impact on Kata or similar runtimes?  
- \[Kevin Hannon\] [KEP-4191](https://github.com/kubernetes/kubernetes/pull/122438) blocked until we have a cadvisor release  
  - With freeze coming, is it possible to get a cadvisor release before the freeze?  
  - \[AI: dawnchen@\] Identify the new owner to help? \- Done\!  
- \[Jeffwan/LingyanYin\]  
  - Need reviewers for this PR \- Configure MemoryRequest for InPlace pod resize in cgroupv2 systems [https://github.com/kubernetes/kubernetes/pull/121218](https://github.com/kubernetes/kubernetes/pull/121218)    
  - [Dixita Narang](mailto:ndixita@gmail.com)drop a comment and doc link for why memory.min shouldn't be set as yet  
- \[AdrianReber\] Graduate "Forensic Container Checkpointing" from Alpha to Beta PR  
  - PR: [https://github.com/kubernetes/kubernetes/pull/123215](https://github.com/kubernetes/kubernetes/pull/123215)  
  - All changes in the PR are based on the KEP discussions  
    - [https://github.com/kubernetes/enhancements/pull/4288](https://github.com/kubernetes/enhancements/pull/4288)  
    - Mainly added tests for existing features as discussed during PRR  
    - Switch from Alpha to Beta feature gate  
    - Added separate sub-resource permission to better control access to the kubelet checkpoint API endpoint  
  - Looking for reviewers  
  - Will probably not be able to make it to the meeting  
- ~~\[fromani\] Looking for approval review: [https://github.com/kubernetes/kubernetes/pull/121778](https://github.com/kubernetes/kubernetes/pull/121778) (for memory manager GA graduation, kubelet observability/visibilty) thanks mrunal\!~~  
- \[jsturtevant\] KEP 2371 \- CRI container and pod stats \- Issue with UsageNanoCores calculated in CRI [https://github.com/kubernetes/kubernetes/issues/122092\#issuecomment-1942699262](https://github.com/kubernetes/kubernetes/issues/122092#issuecomment-1942699262)   
- \[kevin hannon\] PID Stats issues in both containerd and crio  
  - [https://github.com/kubernetes/kubernetes/issues/115215](https://github.com/kubernetes/kubernetes/issues/115215)  
  - [https://github.com/kubernetes/kubernetes/pull/123369](https://github.com/kubernetes/kubernetes/pull/123369)  
  - not sure on crio side why its failing to read any process stats

## Feb 13, 2024

Recording: [https://www.youtube.com/watch?v=WLm7m-8T82A](https://www.youtube.com/watch?v=WLm7m-8T82A)

- ~~kannon92: self nominating to be a reviewer in sig-node~~  
  - [~~https://github.com/kubernetes/test-infra/pull/31891~~](https://github.com/kubernetes/test-infra/pull/31891)  
  - ~~https://github.com/kubernetes/kubernetes/pull/123202~~  
  - Mrunal and Derek approved the above PRs

      \-	\[Vaibhav\] Discuss on the eviction manager issue

- [https://github.com/kubernetes/kubernetes/issues/115201](https://github.com/kubernetes/kubernetes/issues/115201)  
  - KEP [https://github.com/kubernetes/enhancements/issues/4341](https://github.com/kubernetes/enhancements/issues/4341)

      \- 	\[Ritika\] Discuss on this issue  
	      \-     [https://github.com/kubernetes/kubernetes/issues/123176](https://github.com/kubernetes/kubernetes/issues/123176)

      \-     Pranav : Kubelet Thread Management and Resource Cleanup Post-High Workload	  
                   \-    Discuss a scenario where Kubelet retains idle threads post-high workload,              
                        leading to unnecessary memory consumption.  
                   \-    Is there a way in kubernetes to set the number of maximum threads?  
                        If no, can k8s community implement the new parameter for it?  
            [https://github.com/kubernetes/kubernetes/issues/123275](https://github.com/kubernetes/kubernetes/issues/123275)

- gathering pprofs of the kubelet would be useful to see if there are stuck goroutines  
  - try to restrict the kubelet process in systemd unit file to cpuset:0, to force go runtime to allocate less threads and kill them more aggressively, and repeat the test. This would rule out either Go library vs. kubelet thread leaks.

  [https://github.com/golang/go/issues/14592](https://github.com/golang/go/issues/14592)

     

- pehunt: imageRef discussion round 2  
  - Problem: the public pod API field `container.ImageID` is constructed from the container status [ImageRef](https://github.com/kubernetes/kubernetes/blob/master/pkg/kubelet/kuberuntime/kuberuntime_container.go#L619) field.  
  - This ImageID is used to compare against the image.ID of the CRI call for [garbage collection](https://github.com/kubernetes/kubernetes/blob/master/pkg/kubelet/images/image_gc_manager.go#L244-L251).  
  - The `container.ImageID` is considered to be a stable API, but is not compatible with the image.ID field.  
  - Options to fix:  
    - return same value as image.ID in container.ImageRef (resolved repoDigest)  
      - problem: two images tagged with different repos but the same digest would thrash in GC  
    - add a resolvedImageID or something to ContainerStatus and pod API for doing GC  
      - both CRI and pod API update  
    - In GC manager, compare image.RepoDigests in addition to image.ID to find a match  
  - TODO:  
    - check exactly what is returned for each field in cri-o and containerd  
    - investigate if we can put together the needed info in image gc manager without CRI/pod API extension  
    - extend them if not  
- kannon92: (if time) [https://github.com/kubernetes/kubernetes/issues/123247](https://github.com/kubernetes/kubernetes/issues/123247)  
  - Discovered reason for flake in eviction  
  - Summary stats is sometimes failing and the first sort of activePods is ignored  
- ndixita: highlight from Sig Node CI triage meeting (every Wednesday 10AM PST) [https://github.com/kubernetes/kubernetes/issues/122905](https://github.com/kubernetes/kubernetes/issues/122905)

## Feb 6, 2024

Recording: [https://www.youtube.com/watch?v=WiYzo\_knwfk](https://www.youtube.com/watch?v=WiYzo_knwfk)

- \[Filip Krepinsky\] Update on [Declarative Node Maintenance](https://github.com/kubernetes/enhancements/pull/4213)  
-    \[Derek\] Update requested to clarify security posture that would prevent cross node privileges  
- \[pehunt\] [https://github.com/kubernetes/enhancements/pull/3858](https://github.com/kubernetes/enhancements/pull/3858) RRO conversation redux  
- \[jonathan-innis\] Support for `node.kubernetes.io` resource labels for Gt/Lt requirements on pods  
  - See: [https://kubernetes.slack.com/archives/C0BP8PW9G/p1707165434255259](https://kubernetes.slack.com/archives/C0BP8PW9G/p1707165434255259)  
- \[Jeffwan/LingyanYin\] Two things:  
  - Next steps for KEP 4176 \- a new static policy for cpu manager [https://github.com/kubernetes/enhancements/pull/4177\#issuecomment-1930204226](https://github.com/kubernetes/enhancements/pull/4177#issuecomment-1930204226)  
  - Need reviewers for this PR \- Configure MemoryRequest for InPlace pod resize in cgroupv2 systems [https://github.com/kubernetes/kubernetes/pull/121218](https://github.com/kubernetes/kubernetes/pull/121218)    
- \[Vaibhav\] Discuss on the eviction manager issue  
  - [https://github.com/kubernetes/kubernetes/issues/115201](https://github.com/kubernetes/kubernetes/issues/115201)  
  - KEP [https://github.com/kubernetes/enhancements/issues/4341](https://github.com/kubernetes/enhancements/issues/4341)

## Jan 30, 2024

Recording: [https://www.youtube.com/watch?v=LLS3qQgQJ6g](https://www.youtube.com/watch?v=LLS3qQgQJ6g) 

- \[pacoxu\] [Fix evented pleg mirro pod & use IsEventedPLEGInUse instead of FG status check](https://github.com/kubernetes/kubernetes/pull/122778) needs approval and we’d better get inputs from [@smarterclayton](https://github.com/smarterclayton) before merge. This bugfix blocked [sig-release-master-blocking\#gce-cos-master-alpha-features](https://testgrid.k8s.io/sig-release-master-blocking#gce-cos-master-alpha-features). And 1.30.0-alpha.1 is planned for Jan 30th.  
  - Is this an alpha release cut blocker?  
  - Could we ignore \`EventedPLEG\` in the job? We already disabled it in some presubmit Jobs,including \`pull-kubernetes-e2e-kind-alpha-features\` and \`pull-kubernetes-e2e-gce-cos-alpha-features\`.  
- \[anish\] [KEP-3953: Dynamic node resize](https://github.com/kubernetes/enhancements/issues/3953) \- draft at [KEP-3953: Support for Resizable Nodes](https://docs.google.com/document/d/16M0g2L31JZGSBM_oMxszUK1_5AGeqrdyQKjZn7NeCoY/edit?usp=sharing)  
  Anish, please contact on Slack Markus Lehtonen/Francesco Romani, Alexander Kanevskiy \- we will include into discussion thread about that topic.  
- ~~\[tallclair\] Expanding Kubelet configuration API~~  
  - ~~Proposal: [https://github.com/kubernetes/kubernetes/issues/122916](https://github.com/kubernetes/kubernetes/issues/122916)~~  
  - ~~Does this need a KEP? (I think no?)~~  
- \[pehunt\] imageRef usage in the kubelet  
  - context: [https://github.com/cri-o/cri-o/issues/7579](https://github.com/cri-o/cri-o/issues/7579) [https://github.com/cri-o/cri-o/issues/7143](https://github.com/cri-o/cri-o/issues/7143) [https://github.com/kubevirt/kubevirt/pull/10747](https://github.com/kubevirt/kubevirt/pull/10747)   
  -  Yuju shared [https://github.com/kubernetes/kubernetes/issues/46255](https://github.com/kubernetes/kubernetes/issues/46255)  
- \[fromani\] want to improve observability of resource managers: better and more kubelet logs, send kube events on admission failures and in the happy path. Raised as memory manager GA blocker and in general poor observability is a PRR concern. Does this work require a KEP or is an [issue](https://github.com/kubernetes/kubernetes/issues/123037) sufficient?   
  - update KEPs where feasible  
  - For GA KEPs (and in general for this work): update the docs  
  - Keep issues and file the PR when ready  
- \[marquiz/zvonkok\] [KEP-4112: Pass down resources to CRI](https://github.com/kubernetes/enhancements/issues/4112)  
  - PR: [\#4113](https://github.com/kubernetes/enhancements/pull/4113) ([README.md](https://github.com/kubernetes/enhancements/blob/870d027b872dfa289421353e6a7005d59c210bf0/keps/sig-node/4112-passdown-resources-to-cri/README.md))  
  - KEP intro

## Jan 23, 2024

Recording: 

- \[Sergey, Mrunal\] 1.30 Planning [SIG Node - KEP Planning](https://docs.google.com/document/d/1U10J0WwgWXkdYrqWGGvO8iH2HKeerQAlygnqgDgWv4E/edit#bookmark=id.78c7ftvcad73)   
- \[kannon92, AxeZhan\] KEP4328 for 1.30  
  - [https://github.com/kubernetes/enhancements/pull/4329](https://github.com/kubernetes/enhancements/pull/4329)  
  - sig-scheduling planning to implement nodeAffinity type RequiredDuringSchedulingRequiredDuringExecution by adding a new controller, needs a sig-node approver to review this kep also, as sig-node is involved as a participating-sig.  
  - Thank you to Dawn for agreeing to review from sig-node.  
- \[jeffwan, LingyanYin\] two KEPs for 1.30  
  - [https://github.com/kubernetes/enhancements/issues/4176](https://github.com/kubernetes/enhancements/issues/4176)  
    - CPU manager: Adding a static policy option to spread hyperthreads across physical CPUs. Addressed all comments, need approvals  
    - [https://github.com/kubernetes/enhancements/pull/4177\#issuecomment-1903670228](https://github.com/kubernetes/enhancements/pull/4177#issuecomment-1903670228) NRI vs. native cpu manager?  
  - [https://github.com/kubernetes/enhancements/pull/4433](https://github.com/kubernetes/enhancements/pull/4433)  
    - keep inplace VPA KEP alpha for 1.30  
- \[klueska\] Three KEPs for 1.30  
  - Add CDI devices to device plugin API (**promote to GA**)  
    - [https://github.com/kubernetes/enhancements/issues/4009](https://github.com/kubernetes/enhancements/issues/4009)  
  - Add numeric parameters for dynamic resource allocation (**new KEP**)  
    - Simplification / generalization of overall DRA proposal  
    - Context: [Dynamic Resource Allocation (DRA)](https://docs.google.com/document/d/1XNkTobkyz-MyXhidhTp5RfbMsM-uRCWDoflUMqNcYTk/edit#heading=h.ljj9kaa144nr)  
    - [https://github.com/kubernetes/enhancements/pull/4384](https://github.com/kubernetes/enhancements/pull/4384)  
  - Pass down resources to CRI (**new KEP**)  
    - Needed to support GPUs in Kata Containers  
    - [https://github.com/kubernetes/enhancements/pull/4113](https://github.com/kubernetes/enhancements/pull/4113)  
- 

## Jan 16, 2024

Recording: [https://youtu.be/NAIQGQHrlN0](https://youtu.be/NAIQGQHrlN0) 

- \[pacoxu\] EventedPLEG bug of static pods start-up. After reverting it to alpha,  [sig-release-master-blocking\#gce-cos-master-alpha-features](https://testgrid.k8s.io/sig-release-master-blocking#gce-cos-master-alpha-features) keeps failing. \#[122763](https://github.com/kubernetes/kubernetes/pull/122763) is under review.   
  - Latest PR \- [https://github.com/kubernetes/kubernetes/pull/122778](https://github.com/kubernetes/kubernetes/pull/122778)  
- \[kannon92 Kevin\] Update on Swap.   
  - [Swap Beta2 Findings](https://docs.google.com/document/d/1S75d_0N0i1taGTGVjQ8YLYymq8D5Yc0aVu-ROMqQkSs/edit?usp=sharing)   
  - [https://github.com/kubernetes/enhancements/pull/4401](https://github.com/kubernetes/enhancements/pull/4401)  
  - NoSwap seems good  
  - UnlimitedSwap and Eviction signal may be needed  
    - We should add eviction signal for swap for UnlimitedSwap  
    - Or drop support for UnlimitedSwap  
    - Kevin to reach out to [dawnchen@google.com](mailto:dawnchen@google.com)about usecases for swap  
    - Kevin to find examples for UnlimitedSwap.  
- \[pehunt\] proc mount type direction  
  - [https://docs.google.com/document/d/1rYvnhQyi-d8bDgyOGn5FHZKVMgwpygPjksC8ZSBaEPg/edit?usp=sharing](https://docs.google.com/document/d/1rYvnhQyi-d8bDgyOGn5FHZKVMgwpygPjksC8ZSBaEPg/edit?usp=sharing)   
  - Make a KEP update to tie ProcMount behavior to userns (if userns, no masked paths). If there’s pushback, push for ProcMount in Beta  
- \[AkihiroSuda (unlikely to attend due to the timezone)\] Can I get any reaction (an explicit rejection will be highly appreciated more than having no action) to the KEP for Recursive Read-only (RRO) mounts? Has been open for almost a year. If this isn’t going to be accepted I’ll just leave Kubernetes unmodified and change containerd to treat RO as RRO.  
  [https://github.com/kubernetes/enhancements/issues/3857](https://github.com/kubernetes/enhancements/issues/3857) [https://github.com/kubernetes/enhancements/pull/3858](https://github.com/kubernetes/enhancements/pull/3858)   
- \[AdrianReber\] Open Checkpoint/Restore questions from last week  
  - Checkpoint/Restore demo from container image based on [https://kubernetes.io/blog/2022/12/05/forensic-container-checkpointing-alpha/\#restore-checkpointed-container-k8s](https://kubernetes.io/blog/2022/12/05/forensic-container-checkpointing-alpha/#restore-checkpointed-container-k8s)

    \# curl \-s \--insecure \--cert /var/run/kubernetes/client-admin.crt \--key /var/run/kubernetes/client-admin.key \-X POST "https://localhost:10250/checkpoint/default/counters/counter"  
    \# kubectl alpha checkpoint counters  
    \# newcontainer=$(buildah from scratch)  
    \# buildah add $newcontainer /var/lib/kubelet/checkpoints/checkpoint-\<pod-name\>\_\<namespace-name\>-\<container-name\>-\<timestamp\>.tar /  
    \# buildah config \--annotation=io.kubernetes.cri-o.annotations.checkpoint.name=counter  
    \# buildah commit $newcontainer checkpoint-image:latest  
    \# buildah rm $newcontainer  
* How would checkpoint/restore work with pods:  
  * Implemented in March 2022 in combination with kubectl drain  
  * https://github.com/adrianreber/cri-o/commits/checkpoint-restore-support-cri-api/  
  * Pause pod (using cgroup)  
  * Loop over all containers in pod and create a checkpoint  
  * Collect pod metadata  
  * Recreate pod based on metadata (no checkpoint)  
  * Restore all containers  
  * Unpause pod  
* Security review: looking into it  
* Garbage collection mechanism: not thought about it  
* Image-Spec discussion [https://github.com/opencontainers/image-spec/issues/962](https://github.com/opencontainers/image-spec/issues/962)

## Jan 9, 2024

Recording: [https://youtu.be/b5jaZux0qCo](https://youtu.be/b5jaZux0qCo) 

Agenda:

- \[ pehunt \] [https://github.com/kubernetes/kubernetes/pull/117793](https://github.com/kubernetes/kubernetes/pull/117793) ownership. 1.30??  
  - tzneal to take on, no KEP needed  
- \[kannon92\] [https://github.com/kubernetes/kubernetes/pull/121834](https://github.com/kubernetes/kubernetes/pull/121834) looking for approver  
  - Can we consider backporting this?   
  - Agreement  
- \[rata\]: UserNS KEP: beta migration in 1.30?  
  - Open a PR to migrate to beta and reach out to gather more feedback  
- \[tallclair\]: Kubelet config clean up  
  - Now that Dynamic Kubelet config is deprecated & removed, can we move the remaining flags into the Kubelet configuration object?  
    - Derek: look into whether there are any differences in whether the Kubelet needs to be drained on update  
    - Mrunal: Sync with folks working on conf.d  
- \[rst0git\] Forensic Container Checkpointing:   
  - Provide details about additional checkpoint/restore use cases [https://github.com/kubernetes/enhancements/pull/4305](https://github.com/kubernetes/enhancements/pull/4305)  
  - Graduate "Forensic Container Checkpointing" to Beta [https://github.com/kubernetes/enhancements/pull/4288](https://github.com/kubernetes/enhancements/pull/4288)   
  - Add 'checkpoint' command to kubectl [https://github.com/kubernetes/kubernetes/pull/120898](https://github.com/kubernetes/kubernetes/pull/120898)   
  - Proposal: checkpoint image definition  
    [https://github.com/opencontainers/image-spec/issues/962](https://github.com/opencontainers/image-spec/issues/962)   
- \[fromani\]  proposal to allow kubelet to allow the [kubelet to trigger the rescheduling of pods](https://docs.google.com/document/d/1-wJhiNy84w7tzFdo9HqwTu5DrVSuXFLGTUv8FBiRAAc/edit?usp=sharing). (redo from 20240102 because too small audience; presented on batch WG mtg on 20240104 ) \- expected 5 minutes [presentation](https://github.com/ffromani/ffromani/blob/main/docs/proposal-allow-kubelet-to-trigger-rescheduling.pdf) \+ time for questions/discussion maybe 10 mins top?  
  - Include a security section about restricting the node to unbind only its own pods.   
- \[SergeyKanzelev, Harche\] [https://github.com/kubernetes/kubernetes/issues/122224](https://github.com/kubernetes/kubernetes/issues/122224) are back copat concerns here valid?

## Jan 2, 2024

Recording: [https://youtu.be/BHGZs2HJMyU](https://youtu.be/BHGZs2HJMyU)   
Agenda:

- \[marquiz\] [QoS resources KEP](https://github.com/kubernetes/enhancements/pull/3004), call for reviews, blockers from sig-node perspective(?)  
- \[fromani\]  proposal to allow kubelet to allow the [kubelet to trigger the rescheduling of pods](https://docs.google.com/document/d/1-wJhiNy84w7tzFdo9HqwTu5DrVSuXFLGTUv8FBiRAAc/edit?usp=sharing). Looking for early feedback/possible concerns.  
  - spinoff from DRA conversations; beneficial to improve UX with kubelet admission failures  
  - will be presented to batch WG/sig-scheduling mtgs
