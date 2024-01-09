# SIG Node Meeting Notes

## Dec 27th [Canceled for holidays]


## Dec 20th [Canceled for holidays] 

Dec 13th, 2022

Recording: [https://www.youtube.com/watch?v=iw_xZZPuXDI](https://www.youtube.com/watch?v=iw_xZZPuXDI) 

Total: [196](https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+is%3Aopen+) (-22, yay!)


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
   <td><strong><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++created%3A2022-12-06T17%3A00%3A00%2B0000..2022-12-13T17%3A55%3A16%2B0000">32</a></strong>
   </td>
   <td>Closed:
   </td>
   <td><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++is%3Aunmerged+closed%3A2022-12-06T17%3A00%3A00%2B0000..2022-12-13T17%3A55%3A16%2B0000">37</a>
   </td>
  </tr>
  <tr>
   <td>Updated:
   </td>
   <td><strong><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+updated%3A2022-12-06T17%3A00%3A00%2B0000..2022-12-13T17%3A55%3A16%2B0000+created%3A%3C2022-12-06T17%3A00%3A00%2B0000">161</a></strong>
   </td>
   <td>Merged:
   </td>
   <td><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++merged%3A2022-12-06T17%3A00%3A00%2B0000..2022-12-13T17%3A55%3A16%2B0000">22</a>
   </td>
  </tr>
</table>




* [mrunal/sergey/ruiwen] 1.26 retro/1.27 planning
    * [1.26 retro](https://docs.google.com/document/d/11ZVH4OfUUIU73R9suoJ36aoOJJQJi4NJJtmuBkQRmkE/edit), with tracked KEPs and finished KEPs
    * 1.27 planning with initial KEP candidates: 
* [everpeace] [KEP-3169: Fine-grained SupplementalGroups control](https://github.com/kubernetes/enhancements/pull/3620)
    * _NOTE: I’m sorry that I can’t attend to the regular community meeting due to timezone gap (3am in my timezone(Tokyo)). I put this agenda to help 1.27 planning._
    * This KEP can resolve very unfamiliar behavior of `SupplementalGroups` field described in [k/k#112879](https://github.com/kubernetes/kubernetes/issues/112879), which keeps group membership defined in the container image. I believe many (probably most) cluster admins don’t know the behavior. Moreover, when a cluster uses `hostPath` volumes, the unfamiliar behavior could cause security concerns even when cluster admins enforce some policy engines in the cluster.
* [swsehgal] Topology Manager GA graduation: happy to volunteer to drive this work in 1.27 (~~if we want to move ahead with it?~~)
    * Lack of Multi NUMA systems in CI is the key blocker
        * [fromani] this is also relevant for memory manager
    * Currently e2e tests that require multi NUMA are [skipped](https://github.com/kubernetes/kubernetes/blob/release-1.26/test/e2e_node/topology_manager_test.go#L957-L959)
    * [https://github.com/kubernetes/test-infra/issues/28211](https://github.com/kubernetes/test-infra/issues/28211)
        * Inputs/suggestions on how this can be potentially handled welcome on the issue
        * Slack discussion with test-infra group: [here](https://kubernetes.slack.com/archives/CCK68P2Q2/p1670521506420629) (potential use of equinix nodes is being discussed here)
* [SergeyKanzhelev] [https://github.com/kubernetes/kubernetes/pull/114394](https://github.com/kubernetes/kubernetes/pull/114394) CRI API version skew policies. See [slides](https://docs.google.com/presentation/d/1s0RM2zDXKTnTn5Yx0_20V0TPbPJLCQSI/edit) from contributors summit for extra details 
* [vinaykul] [InPlace Pod Vertical Scaling](https://github.com/kubernetes/kubernetes/pull/102884/) PR - status update
    * vinaykul not joining this meeting (on vacation in India)
    * Please review and merge [KEP milestone update PR](https://github.com/kubernetes/enhancements/pull/3670)
    * [PR 102884](https://github.com/kubernetes/kubernetes/pull/102884) approved by Derek.
        * @bobbypage fixed containerd/main E2E pull test job, we now have full E2E coverage
        * I recommend that we merge API changes [PR 111946](https://github.com/kubernetes/kubernetes/pull/111946) at the earliest possible point in 1.27 and watch to see nothing bad happens.
        * And then merge PR 102884 shortly after (&lt; 1 week) and re-add periodic CI test jobs.
        * Does the 1st week of Jan 2023 look realistic for the above proposed PRs merge plan, assuming the above plan sounds good? 
* [SergeyKanzhelev] Reconcile SIG Node teams and OWNERs files: [https://github.com/kubernetes/org/pull/3893](https://github.com/kubernetes/org/pull/3893) 
* [mweston & atanas] Quick update on issue [https://github.com/kubernetes/enhancements/issues/3675](https://github.com/kubernetes/enhancements/issues/3675) 
    * recording: [https://www.youtube.com/watch?v=ai_d3qXr8xg](https://www.youtube.com/watch?v=ai_d3qXr8xg) 


## Dec 6, 2022

Recording: [https://www.youtube.com/watch?v=t3PcHj62f0c](https://www.youtube.com/watch?v=t3PcHj62f0c) 

Total: [218](https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+is%3Aopen+)


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
   <td><strong><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++created%3A2022-11-29T17%3A00%3A00%2B0000..2022-12-06T17%3A53%3A41%2B0000">11</a></strong>
   </td>
   <td>Closed:
   </td>
   <td><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++is%3Aunmerged+closed%3A2022-11-29T17%3A00%3A00%2B0000..2022-12-06T17%3A53%3A41%2B0000">7</a>
   </td>
  </tr>
  <tr>
   <td>Updated:
   </td>
   <td><strong><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+updated%3A2022-11-29T17%3A00%3A00%2B0000..2022-12-06T17%3A53%3A41%2B0000+created%3A%3C2022-11-29T17%3A00%3A00%2B0000">63</a></strong>
   </td>
   <td>Merged:
   </td>
   <td><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++merged%3A2022-11-29T17%3A00%3A00%2B0000..2022-12-06T17%3A53%3A41%2B0000">2</a>
   </td>
  </tr>
</table>




* [mweston & atanas] Ready Dec 6th:  filed issue on kubelet plugin model here:  [https://github.com/kubernetes/enhancements/issues/3675](https://github.com/kubernetes/enhancements/issues/3675)
    * starting KEP and looking for interested parties, discussion partially based around dynamic resource allocation and thoughts on how to incorporate it.
* [pacoxu] small [improvement on memoryThrottlingFactor proposal](https://docs.google.com/document/d/1p9awiXhu5f4mWsOqNpCX1W-bLLlICiABKU55XpeOgoA/edit?usp=sharing)(I listed 3 problems here in the link), but a behavior change. ``memory.high = memory.request + (memory.limit - memory.request) * memoryThrottlingFactor``. Also, defaulting to 0.8 may result to performance issue for pods that may always use 80%+ memory of the limit like Java applications. Probably we need a pod level setting for it like `softrequest`/`throttlingLimit` besides limits and requests.
    * Collect feedback in the doc and then open a KEP update with alternatives (discussion: [https://youtu.be/t3PcHj62f0c?t=686](https://youtu.be/t3PcHj62f0c?t=686))
* [msau] (joining at 10:30) Looking for maintainers from multiple sigs to participate in a discussion/roundtable with [Data on Kubernetes](https://dok.community/) (stateful end users). If you’re interested, [add your name to the list](https://docs.google.com/document/d/1hAKi4nwfzok_dKj_hsavbQRYJuLGzBqTbQIM3GdzB8w/edit?resourcekey=0-MlCEZ-y_BA3Ee4K4bI_jxw#). First roundtable is going to be sometime in January.
* [claudiubelu] Proposed changes to how kubelet detects updates for registered plugins b/c current implementation doesn’t work on Windows due to timestamp granularity issues [https://github.com/kubernetes/kubernetes/pull/114136](https://github.com/kubernetes/kubernetes/pull/114136) 
    * [https://testgrid.k8s.io/sig-windows-signal#windows-unit-master](https://testgrid.k8s.io/sig-windows-signal#windows-unit-master) test failure caused by timestamp
* [SergeyKanzhelev] Sidecar WG: we are getting to the conclusion. Will send summary soon. Find information here: [https://docs.google.com/document/d/1E1guvFJ5KBQIGcjCrQqFywU9_cBQHRtHvjuqcVbCXvU/edit#](https://docs.google.com/document/d/1E1guvFJ5KBQIGcjCrQqFywU9_cBQHRtHvjuqcVbCXvU/edit#) 
* [SergeyKanzhelev] No perma betas:
    * AppArmor beta since 1.4 (owner: @tallclair)
        * [https://github.com/kubernetes/enhancements/pull/3298](https://github.com/kubernetes/enhancements/pull/3298) 
    * QOSReserved alpha since 1.11 (owner: @sjenning)
        * Mrunal or Ryan will take a look
    * RotateKubeletServerCertificate beta since 1.12 (owner: @mikedanese)
        * Sergey to ping Mike
    * CustomCPUCFSQuotaPeriod alpha since 1.12 (owner: @szuecs)
        * Mrunal to take a look
    * KubeletPodResources beta since 1.15 (owner: @dashpole)
        * [@fromanirh] I volunteer to help graduating this in GA in 1.27 - I'll add this to the 1.27 planning document when we start it
    * TopologyManager beta since 1.18  (owner: @lmdaly)
        * Graduate before out of process plugins was the past decision
        * [Dawn] Let’s put together a one pager explaining the roadmap short and longer term.
        * [Swati] Device and CPU manager are graduated. Maybe let’s be consistent
    * DownwardAPIHugePages beta since 1.21 (owner: @derekwaynecarr, )
    * ProbeTerminationGracePeriod beta since 1.22 (owner: )


## Nov 29, 2022



* [aditi] Expose pod cgroup path: [https://github.com/kubernetes/kubernetes/issues/113342](https://github.com/kubernetes/kubernetes/issues/113342)

    [sig-node] Concerns about exposing cgroup information at pod api status level. There could be races depending on how the path will be used. Better to focus down the issue to the interaction between the runtime and CNI plugin at pod bring up time. Peter(CRI-O), David Porter(bobbypage)/mikebrow(containerd) and Mike Zappa(CNI [@MikeZappa87](https://github.com/MikeZappa87) ) to figure out details of approach across runtimes.

* [everperace] [KEP-3169: Fine-grained SupplementalGroups control](https://github.com/kubernetes/enhancements/pull/3620)
    * _NOTE: I’m sorry that I can’t attend to the regular community meeting due to timezone gap (3am in my timezone(Tokyo)). I put this agenda to gain more visibility of my KEP in the sig-node community._
    * This KEP can resolve very unfamiliar behavior of `SupplementalGroups` field described in [k/k#112879](https://github.com/kubernetes/kubernetes/issues/112879). I believe many (probably most) cluster admins don’t know the behavior. Moreover, when a cluster uses `hostPath` volumes, the unfamiliar behavior could cause security concerns even when cluster admins enforce PSPs(or some policy engines) in the cluster. So, I would like to implement the KEP hopefully in v1.27.
    * _I would very appreciate it if somebody help reviewing my KEP._
    * This KEP includes modification of CRI. So, we probably need to update CRI implementation first, at least most popular ones(containerd and cri-o are enough??). I’m not familiar how to do this. I recognize we can’t apply feature gate on CRI and its implementations. I also appreciate it if some contributors advise it to me.
* [klueska] Update to KEP to reflect actual implementation that was merged
    * Needs approval by Dawn or Derek
    * [https://github.com/kubernetes/enhancements/pull/3663](https://github.com/kubernetes/enhancements/pull/3663)
* [swsehgal] [Need](https://github.com/kubernetes/kubernetes/pull/110252#issuecomment-1317422472) Derek’s architecture approval on [https://github.com/kubernetes/kubernetes/pull/110252](https://github.com/kubernetes/kubernetes/pull/110252). API updates are proposed in a separate [PR](https://github.com/kubernetes/kubernetes/pull/96275) (ready for review as well).
* [bobbypage] Update/thoughts on CRI healthz: [https://github.com/kubernetes/kubernetes/pull/109653](https://github.com/kubernetes/kubernetes/pull/109653) 
* [vinaykul] [InPlace Pod Vertical Scaling](https://github.com/kubernetes/kubernetes/pull/102884/) PR - status update
    * vinaykul may not join due to conflicting appointment
    * [PR 102884](https://github.com/kubernetes/kubernetes/pull/102884) approved, missed 1.26, targeting for 1.27
    * Please review and merge [KEP milestone update PR](https://github.com/kubernetes/enhancements/pull/3670)
    * Please review test-infra inplace resize test pull job [PR](https://github.com/kubernetes/test-infra/pull/28112) 


## Nov 22, 2022

No meeting due to the Thanksgiving holiday in USA.


## Nov 15, 2022



* [rata, giuseppe] Userns support
    * For stateful pods, shall we create a new KEP or change the scope of the existing?
    * Will join sig-storage to start the conversation with them about stateful pods too
    * [Derek] - Separate KEP and Feature gate recommended
    * [Sergey] - Should we GA the existing support? 
    * [Derek/Mrunal] Yes, we should move it to beta.
    * [Rodrigo] Concerns around validation if we introduce another Feature Flag.
    * [Rodrigo] Id mapped mounts could solve issues around right permissions for files such as ssh keys.
    * [Derek] Can we key off the kernel version to figure out if we have id mapped mounts? Any way to implement a fallback?
* [klueska] Dynamic Resource Allocation (DRA) update
    * [Merged](https://github.com/kubernetes/kubernetes/pull/111023) on Friday (after an extension request) as an alpha feature for 1.26
    * New staging repo created for [k8s.io/dynamic-resource-allocation](https://github.com/kubernetes/test-infra/pull/27980) with helper libraries to build resource drivers against the DRA API
    * Outstanding [request](https://github.com/kubernetes/org/issues/3837) to create dra-example-driver repo
    * [Request](https://github.com/kubernetes/org/issues/3821#issuecomment-1315409661) to “associate” DRA with an official [sig-node subproject](https://github.com/kubernetes/community/tree/master/sig-node#subprojects)
        * Should we reuse an existing subproject or create a new one?
        * My vote is for a new one (but what to call it?)
* [Sergey] sidecar WG: [https://docs.google.com/document/d/1E1guvFJ5KBQIGcjCrQqFywU9_cBQHRtHvjuqcVbCXvU/edit#](https://docs.google.com/document/d/1E1guvFJ5KBQIGcjCrQqFywU9_cBQHRtHvjuqcVbCXvU/edit#) 


## Nov 8, 2022

Recording: [https://www.youtube.com/watch?v=mnZWYAuOJ90](https://www.youtube.com/watch?v=mnZWYAuOJ90)



* [bobbypage/eric lin] [https://github.com/kubernetes/kubernetes/pull/109653](https://github.com/kubernetes/kubernetes/pull/109653)
* [pacoxu] [kubelet: make registry qps/Burst to limit parallerel pulling count #112242](https://github.com/kubernetes/kubernetes/pull/112242#top) 
    * after rethinking, the current qps/burst of image pull makes no sense to users. And the current PR tries to make it limit the image in pulling process at the same time. The flag and the meaning will not match then. So I suggest to just deprecate and then remove the current registryPullQPS and registryBurst flags. Meanwhile if this is a concern, we should provide a new flag like `parallel-image-pull-limit` as a new feature. (#[112044](https://github.com/kubernetes/kubernetes/issues/112044) the issue) At least we should add more explanation for the flag.(registryPullQPS: ``limit registry pull QPS to this value.`` qps is request per second.)
    * [ruiwen-zhao] +1 on adding a node-level limit of parallel pulls. I can help with this effort.
    * [paco] [https://github.com/containerd/containerd/pull/7313](https://github.com/containerd/containerd/pull/7313) I am working on a PullRequest in containerd to add some image pull related metrics. One of them is the processing count of image pulling
    * [mikebrow] needs more declarative hints in the pod/container spec, and more resource information image manager will not know about other activities… declarative info: qos/cache policy/confidential meta/lazy snapshots vs pull all/does the container runtime optimize for common layers/… as mrunalp says, it’s not just about the image it’s the connection cost/manifests/layers and soon artifacts
* [vinaykul] [InPlace Pod Vertical Scaling](https://github.com/kubernetes/kubernetes/pull/102884/) PR - status update
    * Fixed [nits and updated code](https://github.com/kubernetes/kubernetes/pull/102884/commits/65b032832358d7a57b276fe58aa4b4f2d9935363) to catch up after rebase.
    * Updated [E2E test to run full-spectrum](https://github.com/kubernetes/kubernetes/pull/102884/commits/ee3c2d3c95f2f997d5dd55bbeadca9d40ba970af) for containerd>=1.6.9. Tested in a local cluster.
    * Investigating failures with newly added [cgroupv1](https://testgrid.k8s.io/sig-node-containerd#cos-cgroupv1-inplace-pod-resize-containerd-e2e) , [cgroupv2](https://testgrid.k8s.io/sig-node-containerd#cos-cgroupv2-inplace-pod-resize-containerd-e2e) for in-place resize CI job with containerd-main.
    * Requested 4 day exception to investigate/fix issues from rebase and CI job failure.
    * IMHO, it may be safer to merge this early 1.27 rather than late 1.26
* [iancoolidge] cpuset to kubernetes/utils
    * [time permitting]
    * [https://github.com/kubernetes/kubernetes/pull/113744](https://github.com/kubernetes/kubernetes/pull/113744)
    * minor controversies: NoSort/Sort, Int64 vs int
    * plan: merge all changes here, then copy into k/utils, then revendor in k/k
* ~~[klueska] Need approval from sig-node-leads for feature gate addition in following PR~~
    * ~~[https://github.com/kubernetes/kubernetes/pull/112914](https://github.com/kubernetes/kubernetes/pull/112914)~~
    * ~~I’ve already LGTM’d and APPROVED the kubelet changes, it just needs the feature gate approval now (@liggitt already confirmed to do the API approval)~~
    * ~~Assigning ~~~~since he did the KEP approval~~
* ~~[klueska] Need sig-node-leads approval for creation of dynamic-resource-allocation staging repo~~
    * ~~[https://github.com/kubernetes/org/issues/3821](https://github.com/kubernetes/org/issues/3821)~~
    * ~~Assigning ~~
    * ~~Relevant slack conversation: [https://kubernetes.slack.com/archives/C01672LSZL0/p1667751787905519](https://kubernetes.slack.com/archives/C01672LSZL0/p1667751787905519)~~
* ~~[MaRosset] - Windows hostnetwork alpha #112961~~
    * ~~[https://github.com/kubernetes/kubernetes/pull/112961](https://github.com/kubernetes/kubernetes/pull/112961) ~~


## Nov 1, 2022



* ~~[klueska] Need approval for minor update to KEP~~
    * ~~[KEP-3063: dynamic resource allocation: bump latest-milestone to 1.26](https://github.com/kubernetes/enhancements/pull/3601)~~
* [dashpole] Kubelet context plumbing
    * [https://github.com/kubernetes/kubernetes/pull/113408](https://github.com/kubernetes/kubernetes/pull/113408)
* [vinaykul] [InPlace Pod Vertical Scaling](https://github.com/kubernetes/kubernetes/pull/102884/) PR - status update
    * [Fixed](https://github.com/kubernetes/kubernetes/pull/102884/commits/be27900b11cd59ae73321d10d33bf6521bc185ba) scheduler-focussed E2E test flakiness
    * Cgroupv2 [support](https://github.com/kubernetes/kubernetes/pull/102884/commits/fe3ff951f98943f18e5518cd867f5519e1ad6d5a) and [tests](https://github.com/kubernetes/kubernetes/pull/102884/commits/7198630fdffc94579de6788b1646d263d9651a3c) in review, issues fixed. Mrunal [PTAL](https://github.com/kubernetes/kubernetes/pull/102884#discussion_r958717825).
    * Ruiwen and I have PRs for running pod resize E2E tests with containerd-main.
        * [https://github.com/kubernetes/test-infra/pull/27816](https://github.com/kubernetes/test-infra/pull/27816)
        * [https://github.com/kubernetes/test-infra/pull/27854](https://github.com/kubernetes/test-infra/pull/27854)
        * One or both of these should give us E2E coverage
        * I’m working on updating E2E test to run full-mode if containerd >=1.6.9
* Alexander shared notes from kubecon contributor summit - 


## Oct 25, 2022 [Cancelled for KubeCon]


## Oct 18, 2022

Total active pull requests: [213](https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+is%3Aopen+) (+8 from last week)


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
   <td><strong><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++created%3A2022-10-11T17%3A00%3A00%2B0000..2022-10-18T16%3A56%3A25%2B0000">24</a></strong>
   </td>
   <td>Closed:
   </td>
   <td><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++is%3Aunmerged+closed%3A2022-10-11T17%3A00%3A00%2B0000..2022-10-18T16%3A56%3A25%2B0000">7</a>
   </td>
  </tr>
  <tr>
   <td>Updated:
   </td>
   <td><strong><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+updated%3A2022-10-11T17%3A00%3A00%2B0000..2022-10-18T16%3A56%3A25%2B0000+created%3A%3C2022-10-11T17%3A00%3A00%2B0000">66</a></strong>
   </td>
   <td>Merged:
   </td>
   <td><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++merged%3A2022-10-11T17%3A00%3A00%2B0000..2022-10-18T16%3A56%3A25%2B0000">9</a>
   </td>
  </tr>
</table>




* [SergeyKanzhelev] Canceling next week for KubeCon?
* [Sergey, Swati] CI group report [https://docs.google.com/document/d/1vfqqFtN4Ke2JtB9O4wjoKvMChW2Ptizmsom1_gRGauU/edit#heading=h.eawmmxfxo8vq](https://docs.google.com/document/d/1vfqqFtN4Ke2JtB9O4wjoKvMChW2Ptizmsom1_gRGauU/edit#heading=h.eawmmxfxo8vq) 
* [Sergey] Sidecar WG proposed times: [https://doodle.com/meeting/participate/id/bkZyZgJa](https://doodle.com/meeting/participate/id/bkZyZgJa) 
* [vinaykul] [InPlace Pod Vertical Scaling](https://github.com/kubernetes/kubernetes/pull/102884/) PR - status update
    * Please review [KubeCon slides 11-16](https://static.sched.com/hosted_files/kccncna2022/c0/Resize-with-eBPF.pdf), if possible
    * Cgroupv2 support [changes](https://github.com/kubernetes/kubernetes/pull/102884#issuecomment-1221332179) are in review, issues fixed. Mrunal [PTAL](https://github.com/kubernetes/kubernetes/pull/102884#discussion_r958717825).
    * Awaiting containerd release in order to enable full-E2E tests.
    * Mothership [PR 102884](https://github.com/kubernetes/kubernetes/pull/102884) can merge once we have the 1.6.9 containerd release, the CI picks it up, E2E tests are fully enabled (validates PodStatus for resize), and cgroupv2 review issues have been addressed.
    * API changes [PR 111946](https://github.com/kubernetes/kubernetes/pull/111946) also on hold for containerd.


## Oct 11, 2022

Total active pull requests: [205](https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+is%3Aopen+) (-3 from last week)


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
   <td><strong><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++created%3A2022-10-04T17%3A00%3A00%2B0000..2022-10-11T16%3A41%3A35%2B0000">17</a></strong>
   </td>
   <td>Closed:
   </td>
   <td><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++is%3Aunmerged+closed%3A2022-10-04T17%3A00%3A00%2B0000..2022-10-11T16%3A41%3A35%2B0000">11</a>
   </td>
  </tr>
  <tr>
   <td>Updated:
   </td>
   <td><strong><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+updated%3A2022-10-04T17%3A00%3A00%2B0000..2022-10-11T16%3A41%3A35%2B0000+created%3A%3C2022-10-04T17%3A00%3A00%2B0000">69</a></strong>
   </td>
   <td>Merged:
   </td>
   <td><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++merged%3A2022-10-04T17%3A00%3A00%2B0000..2022-10-11T16%3A41%3A35%2B0000">8</a>
   </td>
  </tr>
</table>




* [matthyx] request a WG creation to work on sidecar containers
    * A summary from Sergey: 
    * [Dawn] define exit criteria and a way to report back status to this SIG. Also define the term for WG.
    * [Sergey] wait for the Doodle to decide scheduling.
* [vinaykul] [InPlace Pod Vertical Scaling](https://github.com/kubernetes/kubernetes/pull/102884/) PR - status update
    * Please review [KubeCon slides 11-16](https://static.sched.com/hosted_files/kccncna2022/c0/Resize-with-eBPF.pdf), if possible
    * Cgroupv2 support [changes](https://github.com/kubernetes/kubernetes/pull/102884#issuecomment-1221332179) are in review, issues fixed. Mrunal [PTAL](https://github.com/kubernetes/kubernetes/pull/102884#discussion_r958717825).
    * Awaiting containerd release in order to enable full-E2E tests.
    * Mothership [PR 102884](https://github.com/kubernetes/kubernetes/pull/102884) can merge once we have the next containerd release (1.7 per Ruiwen - sorry I accidentally deleted Ruiwen’s comment ​), the CI picks it up, E2E tests are fully enabled (validates PodStatus for resize), and cgroupv2 review issues have been addressed.
    * API changes [PR 111946](https://github.com/kubernetes/kubernetes/pull/111946) also on hold for containerd.
* [mimowo]: Heads up for "Standardization of the OOM kill communication between container runtime and kubelet" ([https://github.com/kubernetes/kubernetes/issues/112910](https://github.com/kubernetes/kubernetes/issues/112910))
    * first - standardization of what we have
    * second - add more information - whether it is because exceeding the limits or memory pressure on the node. This is more involving.
    * [Dawn] user space oom killer in cgroupv2 will also introduce more standardization in this space
    * [Dawn] thought it is already aligned, how much has it diverged?
    * [Sergey] is it required for the KEP?
        * [Michael] no, but may break in future
    * [Sergey] How easy to troubleshoot that it was indeed the OOM kill when people start relying on job retries based on OOM kills.
        * [Michael] Feature: customer can define policies for jobs depending on pod end state. Today pod conditions are used to understand the pod end state. Pod condition will be “resource exhausted”.
* [Dawn] there will be cases when kubelet just cannot tell that something was oom killed. But it is still good to have everything unified.
* [David Porter] how practically will it be standardized? It is just a string. Are there any conformance tests or something?
* [Lantao] logging format is the same way
* [David] Container running multiple processes when subprocesses were OOM killed, container runtime may not detect this. 
* [Lantao] cgroupv1 behavior is different, is it? IIRC, when there is a cgroup OOM, a random process in the cgroup will be killed. In that case, OOMKilled is still set, even if pid 1 is running happily. ([David] Let’s confirm)
* [Dawn] this was one of the first issues that was fixed.
* [SergeyKanzhelev] containerd 1.6 is going LTS [https://github.com/containerd/containerd/pull/7454](https://github.com/containerd/containerd/pull/7454) 
    * [Lantao] 1.7 introduces major changes like sandbox API, image pull progress etc.
        * [ruiwen] List of PRs in containerd 1.7: [https://github.com/containerd/containerd/milestone/42](https://github.com/containerd/containerd/milestone/42) 
    * [Lantao] 2.0 will remove many deprecated APIs and configs
        * [https://github.com/kubernetes/kubernetes/issues/110312](https://github.com/kubernetes/kubernetes/issues/110312)  
        * [https://github.com/containerd/containerd/blob/main/RELEASES.md](https://github.com/containerd/containerd/blob/main/RELEASES.md) 
    * [Mark] Windows brings many features to 1.7, Will we end up not supporting 1.6 by K8s


## Oct 4, 2022

Total active pull requests: [208](https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+is%3Aopen+) (-22 since last week)


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
   <td><strong><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++created%3A2022-09-27T17%3A00%3A00%2B0000..2022-10-04T16%3A37%3A24%2B0000">11</a></strong>
   </td>
   <td>Closed:
   </td>
   <td><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++is%3Aunmerged+closed%3A2022-09-27T17%3A00%3A00%2B0000..2022-10-04T16%3A37%3A24%2B0000">21</a>
   </td>
  </tr>
  <tr>
   <td>Updated:
   </td>
   <td><strong><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+updated%3A2022-09-27T17%3A00%3A00%2B0000..2022-10-04T16%3A37%3A24%2B0000+created%3A%3C2022-09-27T17%3A00%3A00%2B0000">89</a></strong>
   </td>
   <td>Merged:
   </td>
   <td><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++merged%3A2022-09-27T17%3A00%3A00%2B0000..2022-10-04T16%3A37%3A24%2B0000">12</a>
   </td>
  </tr>
</table>




* [ruiwen-zhao] [KEP planning] 1.26 KEPs are tracked in: [https://github.com/orgs/kubernetes/projects/98/views/1?filterQuery=sig%3Asig-node](https://github.com/orgs/kubernetes/projects/98/views/1?filterQuery=sig%3Asig-node) 
    * If you are planning to work on any KEP in 1.26, please make sure it is tracked on the board above. To be tracked, the KEP needs to be in the v1.26 milestone and have a lead-opted-in label.
* [SergeyKanzhelev] Sidecar containers: [https://docs.google.com/document/d/10xqIf19HPPUa58GXtKGO7D1LGX3t4wfmqiJUN4PJfSA/edit#heading=h.a3qp744ia2p6](https://docs.google.com/document/d/10xqIf19HPPUa58GXtKGO7D1LGX3t4wfmqiJUN4PJfSA/edit#heading=h.a3qp744ia2p6) 
* ~~[swsehgal] Device Manager graduation to GA~~
    * ~~[https://github.com/kubernetes/enhancements/issues/3573](https://github.com/kubernetes/enhancements/issues/3573)~~
    * ~~Please add `milestone` and `lead-opted-in` labels so it can be tracked for 1.26~~
    * Update: Labels Added (Thanks Mark)
* [ddebroy] Q around runtimeClass fields
    * Is introducing capabilities of a container runtime handler (that Kubelet can use for contexts other than CRI) as fields in RuntimeClass allowed?
    * Previously discussed in the context of KEP [https://github.com/kubernetes/enhancements/pull/2893](https://github.com/kubernetes/enhancements/pull/2893) and it was recommended that fields from runtimeClass should not be passed to CRI.
* [Alexey Fomenko] image pull improvements - design may take more time - feedback from Derek. 
    * [Mrunal] need more details on CRI APIs, etc.
    * [Dawn] still can do progress on this
    * [Alexey] wasn’t sure the process - KEP process doc says that the first draft will be merged fast and then iterated on.


## Sep 27, 2022

Total active pull requests: [230](https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+is%3Aopen+) (+3 since last week)


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
   <td><strong><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++created%3A2022-09-20T17%3A00%3A00%2B0000..2022-09-27T16%3A52%3A21%2B0000">22</a></strong>
   </td>
   <td>Closed:
   </td>
   <td><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++is%3Aunmerged+closed%3A2022-09-20T17%3A00%3A00%2B0000..2022-09-27T16%3A52%3A21%2B0000">12</a>
   </td>
  </tr>
  <tr>
   <td>Updated:
   </td>
   <td><strong><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+updated%3A2022-09-20T17%3A00%3A00%2B0000..2022-09-27T16%3A52%3A21%2B0000+created%3A%3C2022-09-20T17%3A00%3A00%2B0000">76</a></strong>
   </td>
   <td>Merged:
   </td>
   <td><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++merged%3A2022-09-20T17%3A00%3A00%2B0000..2022-09-27T16%3A52%3A21%2B0000">9</a>
   </td>
  </tr>
</table>




* [ruiwen-zhao] Just a quick reminder on 1.26 planning: Please update the status on [1.26 planning](https://docs.google.com/document/d/1U10J0WwgWXkdYrqWGGvO8iH2HKeerQAlygnqgDgWv4E/edit#bookmark=kix.qy9qmwlgcn6h), or add to it if you are planning to work on something not tracked there. **KEP freeze is 18:00 PDT Thursday 6th October 2022**
* [alexey fomenko]: CRI Image Pulling Progress and Notifications: [https://hackmd.io/nyLLTtAkTgOuYwxmnu0sIQ](https://hackmd.io/nyLLTtAkTgOuYwxmnu0sIQ)
    * [paco]: recently, I see some image pulling related issues. 1. [image pull time is including waiting time ](https://github.com/kubernetes/kubernetes/pull/111772)due to default serialize image pulling behavior. 2. no image pulling related metrics in kubelet([pr](https://github.com/kubernetes/kubernetes/pull/111391)) or containerd([pr](https://github.com/containerd/containerd/pull/7313)). 3. [kubelet registry qps/burst](https://github.com/kubernetes/kubernetes/issues/112044) is not working as expected. The current registry qps is like start n image pulling in a second. The user want parallerel image pulling number is qps. BTW, [enabling parallerel image pulling](https://github.com/kubernetes/kubernetes/issues/108405) will solve some problems that are caused by image pulling stucking. We may discuss this as a whole.
    * [lantaol]:
        * We hit this issue as well with serial image pull. A bad container image can block all other pods from coming up forever.
        * However, with parallel image pull, there is no good way to control the concurrency. The QPS is not the best way to solve this problem, because each image pull request can take a long time, just controlling the query per second is not sufficient.
        * This is worse with containerd, which doesn’t have an overall image pull timeout, or a progress based timeout like dockershim.
    * [Derek] Do we want image pull status on Pod as well?
        * [Alexey] yes, we want but not looked into details yet
        * [Derek] qps of progress reports may be a concern for a lot of updates
        * [Alexey] maybe just key points like 25%, 50%
        * [Derek] still a lot of information on happy path. Must be careful with it, only needed for debugging
    * [Derek] is serial pull policy still being used? It only exists for very old runtimes
        * [lantaol] parallel pull may have qps issues
        * [Derek] we have up to n images in parallel. Exactly for this reason. There is no reason to not switch to parallel
    * [Wenjun] many customers wants a image pulling status.
        * [Derek] maybe metrics instead or something that will help minimize the traffic
    * [Lantaol] What about image pull timeouts? Overall timeout is impossible to set. So we had pull timeout. Is it exist in Containerd?
        * [Ruiwen] Containerd pull timeout will be in 1.7: [https://github.com/containerd/containerd/pull/6150](https://github.com/containerd/containerd/pull/6150) 
        * [Sergey] Do we need it configured from kubelet or runtime?
    * [Alexey] another option is to issue an ETA as an update instead of the progress.
    * [Derek] How it will work with “lazy image pulling” like GKE image streaming?
        * [Mrunal] yes, it likely needs to be accounted for.
        * [Dawn] should work well with it.
    * [lantaol] Checked with mrunal offline that, for cri-o the “concurrent pulled image layers” configuration is in cri-o.
        * Containerd only has `MaxConcurrentDownloads` which limits the concurrent downloads for **each image**.
        * To do this at the containerd level, we will need the CRI plugin to implement a cap at the daemon level.
        * Or we can consider implementing this at kubelet level to extend `serial-image-pull` to `max-concurrent-image-pull`.
    * [mikebrow]nod to lantao’s comments.. we have a need to parallelize pulls, a need to identify resource contention when to many parallel pulls (layers/manifest checks/… soon artifacts), a need to handle slow/no progress due to registry contention/access issues… ** Because there is a large cost to getting the resource contention proximity and responding from the back seat if you will, we will probably be better off passing prioritization information/policies from the kubelet side down to the code (in the container runtimes) that is performing the resolve and pulling of layers..
    * Summary: let’s proceed with CRI part of it
* [fangyuchen86]: Kubelet Support Custom Prober Protocol

    

    * [Derek] Where is the probe run?
    * [fangyuchen86] On the node
    * [Derek] who is charged with the execution of these probes? They are not running inside the cgroup, who is reserving compute and memory for these probes.
    * [fangyuchen86] Controller will allocate this - it will be a custom pod in VPC network on the same VPC as a user workload. kubelet cannot access the Pod’s network. It can start it, attach storage, but not the network.
    * [Dawn] I understand the requirement. But maybe that third party controller can take even more responsibility and actually do the pod management on cluster level. It may be easier. Introducing a custom prober to the node has some security concerns too. 
    * [Wenjun] What prevents to create a proxy?
        * [fangyuchen86] security does not allow this
        * [Derek] this is the most interesting question here. Naively we believe that kubelet is an admin for all workload. And why the networking is taken away from the containers. Maybe we can make a session about these requirements?
        * [Dawn] In some environments, the workload cannot access the kubelet network. This is where gVisor helps for example.
        * [Derek] there were nobody before pushing back on probes.
    * [fangyuchen86] we also have a problem of other protocols that needs to be covered. 
        * [Derek] this is separate requirement that might be solved differently.
    * **Summary: Let’s create a doc that explains the requirements and scenarios** 
* [klueska] Small, self-contained TopologyManager update planned for this release:
    * [https://github.com/kubernetes/enhancements/issues/3545](https://github.com/kubernetes/enhancements/issues/3545)
    * Already added to 
    * Please add the **<code>/label lead-opted-in</code></strong> to the issue so it can be tracked
    * [Derek] this is done
* sig-node meeting recordings: uploading recent ones.
    * [Derek] will do
    * [Dawn] might help with it
* [Sergey] per container restartPolicy override: [https://docs.google.com/document/d/1gX_SOZXCNMcIe9d8CkekiT0GU9htH_my8OB7HSm3fM8/edit](https://docs.google.com/document/d/1gX_SOZXCNMcIe9d8CkekiT0GU9htH_my8OB7HSm3fM8/edit) 
    * [Derek] Alternative is “BindsTo” semantics that can be used for termination of sidecar containers. Maybe kubelet can delegate it to OS as well like systemd.
    * [Mrunal] BindsTo needs to be experimented. But delegating to OS is also appealing
    * [MikeB] Question is how much we can delegate.

		<strong>summary: read the doc and compare with BindsTo.</strong>



* [Sergey] more sidecars: [https://github.com/kubernetes/kubernetes/issues/111356#issuecomment-1249670702](https://github.com/kubernetes/kubernetes/issues/111356#issuecomment-1249670702)
    * [Dawn] restartPolicy and QoS (including OOM score) were per container initially. But then after debate it was changed per pod and convinced community. There were even conversations of scheduling container into the Pod. This is why sidecar feel so unnatural in k8s. Once opening this can of worm, we are starting to do Pod v2.
    * [Mrunal] this OOM adj calculation may be challenging.
    * [Derek] I’d rather move to OOMd than change the present state.

    **summary: likely not (closed the issue).**

* [marquiz] update on [QoS-class resources KEP](https://github.com/kubernetes/enhancements/pull/3004) 
* [pacoxu] If ResourceQuota for cpu/memory is set, no best effort pod can be created. For other resources like ephemeral-storage, best effort pod can be created.
    * [https://github.com/kubernetes/kubernetes/issues/112310#issuecomment-1247540260](https://github.com/kubernetes/kubernetes/issues/112310#issuecomment-1247540260)  Should we document it? Or should we fix it as a bug(Currently, only some comments in code)?


## Sep 20, 2022



* [ruiwen/mrunal] [1.26 planning](https://docs.google.com/document/d/1U10J0WwgWXkdYrqWGGvO8iH2HKeerQAlygnqgDgWv4E/edit#bookmark=kix.qy9qmwlgcn6h)
* [marquiz]  [QoS-class resources KEP](https://github.com/kubernetes/enhancements/pull/3004) (renamed)
    * Status update ([slides](https://docs.google.com/presentation/d/1zBzKfieYw6Kkxq0T_wUsMfodHYNGd7WmijZLIVzcQnQ/edit#slide=id.p1)
* [vinaykul] [InPlace Pod Vertical Scaling](https://github.com/kubernetes/kubernetes/pull/102884/) PR - status update
    * [Fabian](https://github.com/fabi200123) has a PR adding [Windows support](https://github.com/vinaykul/kubernetes/pull/16/commits/457644a47f99f862b31c9d3c306bbb011e13341b) for in-place resize. Thanks Fabian!
    * [JaffWan](https://github.com/Jeffwan) fixed my [missing unit tests and typo](https://github.com/vinaykul/kubernetes/pull/17) for cgroupv2 :) Thanks Jaixin!
    * Jaixin also found the root-cause for [issue 112264](https://github.com/kubernetes/kubernetes/issues/112264) and will work on a fix. This should significantly speed up E2E tests.
    * I tried out in-place resize with CRI-O (in local cluster) and it works!
    * Tested resize E2E tests using Ruiwen’s containerd support
        * It works but PodStatus.Resources update takes ~60s. Issue [112264](https://github.com/kubernetes/kubernetes/issues/112264)
        * UpdateContainerResources (containerd) applies the resize in  &lt; 50 ms
    * API changes [PR 111946](https://github.com/kubernetes/kubernetes/pull/111946) ready for review & preferably early-merge.
    * Cgroupv2 support [changes](https://github.com/kubernetes/kubernetes/pull/102884#issuecomment-1221332179) are in review.
    * Mothership [PR 102884](https://github.com/kubernetes/kubernetes/pull/102884) can merge once we have the next containerd release (1.6.9?), the CI picks it up, E2E tests are fully enabled (validates PodStatus for resize), and cgroupv2 review issues have been addressed.
* [mimowo] Promote KEP-3329 "Retriable and non-retriable Pod failures for Jobs" for Beta
* [Sergey] SIG ongoing things update:

Total active pull requests: [227](https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+is%3Aopen+) (+35 since June)


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
   <td><strong><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++created%3A2022-09-06T17%3A00%3A00%2B0000..2022-09-20T16%3A54%3A08%2B0000">32</a></strong>
   </td>
   <td>Closed:
   </td>
   <td><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++is%3Aunmerged+closed%3A2022-09-06T17%3A00%3A00%2B0000..2022-09-20T16%3A54%3A08%2B0000">10</a>
   </td>
  </tr>
  <tr>
   <td>Updated:
   </td>
   <td><strong><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+updated%3A2022-09-06T17%3A00%3A00%2B0000..2022-09-20T16%3A54%3A08%2B0000+created%3A%3C2022-09-06T17%3A00%3A00%2B0000">87</a></strong>
   </td>
   <td>Merged:
   </td>
   <td><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++merged%3A2022-09-06T17%3A00%3A00%2B0000..2022-09-20T16%3A54%3A08%2B0000">15</a>
   </td>
  </tr>
</table>


Bugs untriaged: 13 [https://github.com/orgs/kubernetes/projects/59](https://github.com/orgs/kubernetes/projects/59) 

PRs untriaged: 99 [https://github.com/orgs/kubernetes/projects/49](https://github.com/orgs/kubernetes/projects/49) 

CI group meetings are back and we will be triaging issues and getting the tests back on track.


## Sep 13, 2022



* [mweston/atanas] KEP involving Resource Control Plugin
    * [https://docs.google.com/presentation/d/1uIHF_t97WZzIJgvyD75PYu46PRPi62i2/](https://docs.google.com/presentation/d/1uIHF_t97WZzIJgvyD75PYu46PRPi62i2/)
* 
* [jstur/MaRosset] Windows cri pod sandbox fields: [https://github.com/kubernetes/enhancements/pull/3439](https://github.com/kubernetes/enhancements/pull/3439)
    * Decided to keep them separate.  
* [MaRosset] - Host/Node network support for Windows Pods \
[https://github.com/kubernetes/enhancements/pull/3507](https://github.com/kubernetes/enhancements/pull/3507)
    * +1 from sig-node to proceed 
* [marquiz]  [QoS-class resources KEP](https://github.com/kubernetes/enhancements/pull/3004) (renamed) update
* [vinaykul] [InPlace Pod Vertical Scaling](https://github.com/kubernetes/kubernetes/pull/102884/) PR - status update
    * vinaykul not in the SIG-node meeting today due to conflict.
    * Tested resize E2E tests using Ruiwen’s containerd support
        * It works but PodStatus.Resources update takes ~60s. Issue [112264](https://github.com/kubernetes/kubernetes/issues/112264)
        * UpdateContainerResources (containerd) applies the resize in  &lt; 50 ms
    * API changes [PR 111946](https://github.com/kubernetes/kubernetes/pull/111946) ready for review & preferably early-merge.
    * Cgroupv2 support [changes](https://github.com/kubernetes/kubernetes/pull/102884#issuecomment-1221332179) are in review.
    * Mothership [PR 102884](https://github.com/kubernetes/kubernetes/pull/102884) can merge once we have the next containerd release (1.6.9?), the CI picks it up, E2E tests are fully enabled (validates PodStatus for resize), and cgroupv2 review issues have been addressed.


## Sep 6, 2022



* [ruiwen] [1.25 retro](https://docs.google.com/document/d/1dLmSOjaeahMdHLQyOpGQgrmJBsBrTpBIKto2bdlEwas/edit)
* [danielye] [CRI Stats Performance Update ](https://docs.google.com/document/d/11kx0bnsKE1GjhtM7maVmEYO1eGfUDdC9yv0AzrEeIuk/edit?resourcekey=0-mGWn_TXfmPirsIWeN6TqXw#heading=h.5irk4csrpu0y)
* [qiutongs] Issue awareness: unexpected initial delay of probes
    * [https://github.com/kubernetes/kubernetes/issues/96614#issuecomment-1236258797](https://github.com/kubernetes/kubernetes/issues/96614#issuecomment-1236258797)
    * ``initialDelaySeconds` `doesn’t work as the API spec says.
        * first probe time = container start time + initialDelaySeconds
        * kubelet restart: wait reasonable amount of time; still respect initialDelaySeconds differences for probes in the same container?
    * Jitter is needed in the case of kubelet restart. Avoid thundering herd problems.
        * The jitters given to the probes in the same container are different.
            * [https://github.com/kubernetes/kubernetes/pull/102064](https://github.com/kubernetes/kubernetes/pull/102064)
        * Since [1.21](https://github.com/kubernetes/kubernetes/blob/v1.21.0/pkg/kubelet/prober/worker.go#L137-L139), the jitter is only added when kubelet recently started/restarted. 
            * If the `periodSeconds` are the same for all probes in a container, the probes will be invoked at the same time. ``initialDelaySeconds``makes no difference.
* [adrianreber] Checkpoint/Restore next steps
    * main focus how to secure checkpoint
    * a checkpoint contains all memory pages (maybe secrets, random numbers)
    * possible suggestions how to secure checkpoint archives
        * add additional authorization to kubelet API checkpoint endpoint
        * run kubelet API checkpoint endpoint on another port
        * encrypt checkpoint archives ([https://github.com/containers/ocicrypt](https://github.com/containers/ocicrypt))
* [vinaykul] [InPlace Pod Vertical Scaling](https://github.com/kubernetes/kubernetes/pull/102884/) PR status update
    * Attempted “full scope” E2E tests with Ruiwen’s containerd support
        * Tested this by switching containerd binaries on GKE worker with those I built from master latest.
        * All tests pass and verify that Ruiwen’s code works correctly.
        * The tests took <span style="text-decoration:underline;">considerably</span> longer (869s vs 2988s  to run all 34 E2E tests with GKE 1master-1worker cluster)
            * The issue is NOT in containerd.
            * ContainerStatus() CRI called within 50 millisec of UpdateContainerResources() shows updated cgroup values.
            * The long delay is in updating apiPodStatus. This needs further investigation.
    * [cgroup v2 support](https://github.com/kubernetes/kubernetes/pull/102884#issuecomment-1221332179) for in-place resize - review is in progress.
    * Is there any interest in merging API changes ( [PR 111946](https://github.com/kubernetes/kubernetes/pull/111946) ) early?
        * If yes, please add an ok-to-test label.
* 


## Aug 30, 2022



* [klueska, pohly] Update on Dynamic Resource Allocation
    * [KEP](https://github.com/kubernetes/enhancements/pull/3064) accepted for 1.25
    * Delayed implementation to 1.26  \
(Mostly functional [Draft PR](https://github.com/kubernetes/kubernetes/pull/111023))
    * Demo with NVIDIA GPUs
* [dgl] Status of ProcMountType feature gate
    * This has been alpha since 1.12, I’m interested in potentially progressing it
* [pehunt] inheritable capabilities regression follow up
    * [https://github.com/kubernetes/kubernetes/issues/111196](https://github.com/kubernetes/kubernetes/issues/111196) 
    * Outcome: each CRI should consider whether the capabilities should be added, and optionally be able to configure them. 
        * CRI-O follow up https://github.com/cri-o/cri-o/pull/6236
* [mgroot]
    * Revisiting allowing node labels to be referenced via the downward API
        * [https://github.com/kubernetes/kubernetes/issues/40610](https://github.com/kubernetes/kubernetes/issues/40610)
        * Closed PR: [https://github.com/kubernetes/kubernetes/pull/25957#issuecomment-228041766](https://github.com/kubernetes/kubernetes/pull/25957#issuecomment-228041766)
        * node name and scheduleability are allowed, prior decision from 2016 could be revisited after 6 years
* [marquiz]  [QoS-class resources KEP](https://github.com/kubernetes/enhancements/pull/3004) (renamed)
    * request for comments
    * also for post in k8s developers’ blog ([PR](https://github.com/kubernetes/contributor-site/pull/332))
* [vinaykul] [InPlace Pod Vertical Scaling](https://github.com/kubernetes/kubernetes/pull/102884/) PR status update
    * CRI (containerd) [support](https://github.com/containerd/containerd/pull/6517) has been merged. Thanks Ruiwen!
        * Next step: containerd release. K8s pickup new containerd version.
    *  [PR 111946](https://github.com/kubernetes/kubernetes/pull/111946)  (API changes for in-place resize from 102884) needs ok-to-test
    * [cgroup v2 support](https://github.com/kubernetes/kubernetes/pull/102884#issuecomment-1221332179) for in-place resize awaiting review.
    * Marion Lobur from the GKE team (Warsaw) is joining this effort.
        * His use case can help get significant test coverage in alpha.
* [qbarrand]
    * [Code push complete](https://github.com/kubernetes-sigs/kernel-module-management/pull/9) - development ongoing
    * [KMM admin PR](https://github.com/kubernetes/org/pull/3626) - needs attention from one of the chairs
    * Sponsorship for an additional KMM contributor to become member of the kubernetes org


## Aug 23, 2022



* [rphillips] Ephemeral Storage by Disk/Mount Point [[Issue#111965](https://github.com/kubernetes/kubernetes/issues/111965)]
* [vinaykul] [InPlace Pod Vertical Scaling](https://github.com/kubernetes/kubernetes/pull/102884/) PR status update
    * [KEP PR](https://github.com/kubernetes/enhancements/pull/3464) updating milestone to v1.26-alpha has been merged. Thanks Dawn.
    * Created [PR 111946](https://github.com/kubernetes/kubernetes/pull/111946) - API changes for in-place resize. (already reviewed in [PR 102884](https://github.com/kubernetes/kubernetes/pull/102884) and stable for quite a while)
        * Can we merge 111946 early in 1.26?
        * The mothership 102884 PR can bring in kubelet & scheduler changes when Ruiwen’s [CRI support code](https://github.com/containerd/containerd/pull/6517) has been merged and picked up.
    * Mrunal, LiuBo, Can you please review [cgroup v2 support](https://github.com/kubernetes/kubernetes/pull/102884#issuecomment-1221332179) for in-place resize?
        * Mike, Peter: Should we eventually delegate this to CRI?
* [dawnchen] SIG Node CI testgrid
    * [https://testgrid.k8s.io/sig-node-kubelet#kubelet-gce-e2e-swap-fedora](https://testgrid.k8s.io/sig-node-kubelet#kubelet-gce-e2e-swap-fedora) failed all the time. 
    * [https://docs.google.com/spreadsheets/d/1IwONkeXSc2SG_EQMYGRSkfiSWNk8yWLpVhPm-LOTbGM/edit#gid=1187923038](https://docs.google.com/spreadsheets/d/1IwONkeXSc2SG_EQMYGRSkfiSWNk8yWLpVhPm-LOTbGM/edit#gid=1187923038)
    * [Brian McQueen (github: xmcqueen)] two linked ticket I am currently driving:  [https://github.com/kubernetes/kubernetes/issues/109295](https://github.com/kubernetes/kubernetes/issues/109295) and related PR[ https://github.com/kubernetes/test-infra/pull/27051](https://us01st-cf.zoom.us/web_client/6orpgrb/html/externalLinkPage.html?ref=https://github.com/kubernetes/test-infra/pull/27051)


## Aug 16, 2022



* [bobbypage/pehunt/Daniel] CRI stats Prometheus endpoint and adding cAdvisor metrics to CRI
    * KEP as it stands proposes to have CRI impl emit the prometheus metrics. New proposal is to enhance CRI API to have CRI impl give Kubelet the metrics, then have Kubelet emit them.
        * Same concerns with performance, but possibly it won’t be as bad as we worry about
    * Plan is to have Daniel/David/Peter setup proof of concept with Kubelet/containerd fork to see if it regresses in performance.
    * Aim to have a POC for 1.26 KEP time to be able to decide how to go forward in 1.26 cycle.
* [ndixita] Looking for a reviewer for External Credential Provider GA PR: [https://github.com/kubernetes/kubernetes/pull/111495](https://github.com/kubernetes/kubernetes/pull/111495) 
* [vinaykul] [InPlace Pod Vertical Scaling](https://github.com/kubernetes/kubernetes/pull/102884/) PR status update
    * vinaykul not in Node meeting today due to conflict
    * Please review [KEP PR](https://github.com/kubernetes/enhancements/pull/3464) that updates milestones for this feature
    * I will spawn a separate PR for API changes later this week.
* [pehunt] How best to request reviewers to look at a PR not attached to a release cycle
    * shameless plug for [https://github.com/kubernetes/kubernetes/pull/108855](https://github.com/kubernetes/kubernetes/pull/108855) 
    * Recommendations:
        * PRs appear on triage board: [https://github.com/orgs/kubernetes/projects/49](https://github.com/orgs/kubernetes/projects/49) 
        * Poke reviewers on slack (sig-node, pr-reviews)
        * Propose review trades


## Aug 09, 2022



* [jing]: [https://github.com/kubernetes/kubernetes/issues/110956#issuecomment-1197710132](https://github.com/kubernetes/kubernetes/issues/110956#issuecomment-1197710132)
    * any reason  “it was a mistake that user must make an explicit request/limit when ResourceQuota has CPU/Memory setting."?
* [vinaykul] [InPlace Pod Vertical Scaling](https://github.com/kubernetes/kubernetes/pull/102884/) PR status update
    * Extracted and merged CRI changes in [PR 111645](https://github.com/kubernetes/kubernetes/pull/111645)
        * Many thanks to Mrunal, Peter, Mike, Ruiwen & Mark for quick reviews!!
        * Huge thanks to wangchen615 for pushing hard on the scheduler code!
        * This now unblocks runtime implementing support for in-place update.
    * If there are no objections, can I squash all the discrete kubelet commits till now into a single commit (easier to rebase), and then add cgroupv2 support?
    * Can we target an early 1.26 merge of API code? (saves me rebase headache)
* [qbarrand] kubernetes-sigs membership for the KMMO contributors [PR](https://github.com/kubernetes-sigs/kernel-module-management/pull/7)
    * feel free to ping @endocrimes
* [pehunt] repercussions of dropping inheritable capabilities - https://github.com/moby/moby/issues/43420


## Aug 02, 2022

Reminder of the coming code freeze on 08/02.



* [vinaykul] [InPlace Pod Vertical Scaling](https://github.com/kubernetes/kubernetes/pull/102884/) PR status update
    * Pod resize E2E test failed after rebase because GKE switched to[ cos-97 last week](https://github.com/kubernetes/kubernetes/commit/ead45ba74de3e70a7f807a3b4dee961763401a76) which defaults to cgroup2
        * we don’t support it yet (it was planned for beta)
        * disabling cgroup values check for resize verification unblocks but manual test on pre cos-97 (cgroup1) needed
    * wangchen615 found that the scheduler takes 5 minutes to reevaluate pending pods after resizing down a bound pod.
        * Late stage fix in commits <code>[563b254](https://github.com/kubernetes/kubernetes/commit/563b254303cb06c4d66195679a10f6c083c8fc54)</code> and <code>[c6581a8](https://github.com/kubernetes/kubernetes/commit/c6581a82ae3f84e7d4ff455c04757e0f1e75293c)</code>
        * SIG-scheduling feels it is low risk and has signed off unofficially on slack
    * thockin [has been LGTM](https://github.com/kubernetes/kubernetes/pull/102884#pullrequestreview-923756943) on API changes.
    * Open issues [tracked here](https://github.com/vinaykul/kubernetes/wiki/In-Place-Pod-Vertical-Scaling-Issues-and-Status).
    * My sentiment has changed - I am nervous about late stage changes and missing cgroup2 support when CI is on cgv2
* [harche] evented PLEG - [https://github.com/kubernetes/kubernetes/pull/111384](https://github.com/kubernetes/kubernetes/pull/111384)

                [   https://github.com/kubernetes/kubernetes/pull/111642](https://github.com/kubernetes/kubernetes/pull/111642)

* [rata]: Asked for exception for [userns PR](https://github.com/kubernetes/kubernetes/pull/111090), as talked with Mrunal on slack
    * Mrunal mentioned there are some concerns with phase II to discuss
    * rata to update on PR, KEP and exception:
        * Capture the discussion, we agreed to reduce the scope to stateless pods.
        * One one hand, this buys us more time to figure out the details that some reviewers want about persistent volume support. On the other, it is very valuable to have support for stateless pods and it is an end in itself.
        * This should also eliminate all concerns on how this can graduate to beta and GA.
        * Should we change the feature gate name?
* [jing]: [https://github.com/kubernetes/kubernetes/issues/110956#issuecomment-1197710132](https://github.com/kubernetes/kubernetes/issues/110956#issuecomment-1197710132)
    * any reason  “it was a mistake that user must make an explicit request/limit when ResourceQuota has CPU/Memory setting."?
* [bobbypage] cgroup v2 GA update
    * CI has been running cgroupv2 images (COS/Ubuntu) on node e2e and cluster e2e 
    * cgroupv1 specific tests added
    * More feedback has been obtained on cgroupv2 from customers 
        * Tencent has been running on cgroups v2 for a while
    * Planned doc updates and blog post


## July 26, 2022



* [ruiwen-zhao]: 
    * Reminder of the coming code freeze on 08/02..
    * KEP status: 
* [Xander/Anish]: Discuss [KEP #1003](https://github.com/kubernetes/enhancements/pull/1003)
    * [https://github.com/intel/platform-aware-scheduling/tree/master/telemetry-aware-scheduling/](https://st1.zoom.us/web_client/5ccnbv4/html/externalLinkPage.html?ref=https://github.com/intel/platform-aware-scheduling/tree/master/telemetry-aware-scheduling/)
    * Power example scheduling: [ https://github.com/intel/platform-aware-scheduling/tree/master/telemetry-aware-scheduling/docs/power](https://st1.zoom.us/web_client/5ccnbv4/html/externalLinkPage.html?ref=https://github.com/intel/platform-aware-scheduling/tree/master/telemetry-aware-scheduling/docs/power)
* [Jing] Follow up on local Storage Capacity Isolation Feature 
    * discussed with Ben [https://github.com/kubernetes/enhancements/issues/361#issuecomment-1194457546](https://github.com/kubernetes/enhancements/issues/361#issuecomment-1194457546)
    * automatic detect and silence the feature will cause problem when system is switching to rootless in the future
    * (dawnchen): SIG Node is fine with the proposal. 
* [rata]: [userns k/k PR](https://github.com/kubernetes/kubernetes/pull/111090) is open for 2 weeks now, need review/approvers
    * We have a LGTM from Ruiwen (thanks!), but need approvers from [lot of paths](https://github.com/kubernetes/kubernetes/pull/111090#issuecomment-1189537665)
        * Dawn will ping Jordan (github @liggitt) for approval in several paths. Tim Hockin is out this week
    * Any help to have the lgtm/approved on time?
    * Unrelated: Maybe user namespaces is a good topic for a kubernetes blog post?
        * (danielle): yes!
        * Danielle is asking how to proceed to go with a blog post for this (thanks!)
        * (danielle): **update: **added to the sig-release tracking sheet!
* [ddebroy]: PodHasNetwork pod condition PR needs node reviewer/approver
    * [https://github.com/kubernetes/kubernetes/pull/111358](https://github.com/kubernetes/kubernetes/pull/111358)
    * Confirmed that API review is not required for Alpha const definitions.
    * Mrunal has started on this. Thanks Mrunal!!
    * Qiutong will also take a look
* [vinaykul] [InPlace Pod Vertical Scaling](https://github.com/kubernetes/kubernetes/pull/102884/) PR status update
    * Awaiting Node LGTM
    * Added [concise guidance](https://github.com/kubernetes/kubernetes/pull/102884/commits/0b04ed5a9305f5856294509281f7db307adcef48) for UpdateContainerResources CRI. mikebrow is LGTM
    * SIG-Scheduling (huang-wei) [signaled LGTM](https://github.com/kubernetes/kubernetes/pull/102884#issuecomment-1171515624) for current code.
        * E2E test and optimization can come in follow-up PRs.
            * wangchen615 is close to wrapping up tests requested by SIG-scheduling
    * thockin [has been LGTM](https://github.com/kubernetes/kubernetes/pull/102884#pullrequestreview-923756943) on API changes.
    * Open issues [tracked here](https://github.com/vinaykul/kubernetes/wiki/In-Place-Pod-Vertical-Scaling-Issues-and-Status).
* [bobbypage] cgroupv2 GA
    * New tests added to support cgroupv1
    * Blocking presubmit updated to latest COS-97 and other tests to Ubuntu 22.04 with cgroup v2 enabled([https://github.com/kubernetes/kubernetes/pull/111412](https://github.com/kubernetes/kubernetes/pull/111412) , [https://github.com/kubernetes/test-infra/pull/26847](https://github.com/kubernetes/test-infra/pull/26847), https://github.com/kubernetes/test-infra/pull/26831
    * Feedback gathered about some customer usage


## July 19, 2022



* [rata]: Can we have a review for [userns k/k PR](https://github.com/kubernetes/kubernetes/pull/111090)? Code freeze is coming soon :)
    * Mrunal is reviewing it
    * Ruiwen would review it for Containerd related changes
* [jstur]: Windows cri-only posandbox stats: [https://github.com/kubernetes/kubernetes/pull/110754](https://github.com/kubernetes/kubernetes/pull/110754).  Do we keep the structure generic or make it specific for windows?
    * David Porter is reviewing it.
    * move design (including [https://github.com/kubernetes/kubernetes/pull/110754#issuecomment-1175552676](https://github.com/kubernetes/kubernetes/pull/110754#issuecomment-1175552676)) to the kep
* [Brett]: SRO to kmmo repo rename
    * Brett to open an issue to get the rename going
* [vinaykul] [InPlace Pod Vertical Scaling](https://github.com/kubernetes/kubernetes/pull/102884/) PR status update
    * Rebased code to resolve latest conflict
    * Added [concise guidance](https://github.com/kubernetes/kubernetes/pull/102884/commits/eae2354e407ce2af2213b7e44b8f08ec6da619a6) for UpdateContainerResources CRI
    * SIG-Scheduling (huang-wei) [signaled LGTM](https://github.com/kubernetes/kubernetes/pull/102884#issuecomment-1171515624) for current code.
        * E2E test and optimization can come in follow-up PRs.
            * wangchen615 is working on addressing Danielle’s feedback & E2E test targeting scheduler changes
    * thockin [has been LGTM](https://github.com/kubernetes/kubernetes/pull/102884#pullrequestreview-923756943) on API changes.
    * Open issues [tracked here](https://github.com/vinaykul/kubernetes/wiki/In-Place-Pod-Vertical-Scaling-Issues-and-Status).
    * What do we need for Node & CRI LGTM for alpha?
* [Jing] Local Storage Capacity Isolation Feature 
    * Problem: 
        * [https://github.com/kubernetes/enhancements/issues/361#issuecomment-1172435157](https://github.com/kubernetes/enhancements/issues/361#issuecomment-1172435157)
        * It depends on the capability of checking local storage root filesystem usage where kubelet is running. Some special systems (kind rootless) cannot detect root file system disk usage, so disable this feature with feature gate for CI testing
        * Feature gate will be removed when moving to GA 1.25. If rootfs disk usage is not available by cadvisor, it will block kubelet starting. [https://github.com/kubernetes/kubernetes/blob/5b92e46b2238b4d84358451013e634361084ff7d/pkg/kubelet/kubelet.go#L1385](https://github.com/kubernetes/kubernetes/blob/5b92e46b2238b4d84358451013e634361084ff7d/pkg/kubelet/kubelet.go#L1385) 
    * Proposal: Add kubelet option enableLocalStorageCapacityIsolation(default=true) into kubelet configuration 
        * The default value is true. 
        * For systems that cannot support detecting root disk usage, set enableLocalStorageCapacityIsolation=false in kubelet configuration. In this case, kubelet can continue to start without rootfs disk usage information. So ephemeral storage allocatable is not set either. If pod has ephemeral storage request/limit set in this case, pod will fail to create because allocatable storage is not available.
        * [feedback] whether we can automatically detect this. avoid complicating kubelet
    * Feedback from sig-storage: seems fine
* [Peter] CRI stats check-in
* [Dawn] FYI: PR for SIG Node Contributor ladder was merged: [https://github.com/kubernetes/community/pull/6725](https://github.com/kubernetes/community/pull/6725) Thanks Derek!
    * Please send your PR against that ladder if you think you are ready. Thanks!


## July 12, 2022



* [danielfoehrn] KEP proposal: [Dynamic Resource Reservations](https://github.com/danielfoehrKn/enhancements/commit/54cc9a3fa5b32d4dd4cc9f556631bb5ca9d73c8a)
* 
* [vinaykul] [InPlace Pod Vertical Scaling](https://github.com/kubernetes/kubernetes/pull/102884/) PR status update
    * Rebased code to resolve latest conflict
    * Guidance for runtime is [taking shape](https://github.com/kubernetes/kubernetes/pull/102884#issuecomment-1171809939) - thanks mrunalp & kolyshkin
    * SIG-Scheduling (huang-wei) [signaled LGTM](https://github.com/kubernetes/kubernetes/pull/102884#issuecomment-1171515624) for current code.
        * E2E test and optimization can come in follow-up PRs.
            * wangchen615 is working on addressing Danielle’s feedback & E2E test targeting scheduler changes
    * thockin [has been LGTM](https://github.com/kubernetes/kubernetes/pull/102884#pullrequestreview-923756943) on API changes.
    * Open issues [tracked here](https://github.com/vinaykul/kubernetes/wiki/In-Place-Pod-Vertical-Scaling-Issues-and-Status).
    * What do we need for Node & CRI LGTM for alpha?
* [adrianreber] Forensic Container Checkpointing
    * code PR [https://github.com/kubernetes/kubernetes/pull/104907](https://github.com/kubernetes/kubernetes/pull/104907)
    * LGTM by Ryan, Danielle, Mike(per existing kep) and Mrunal
    * I think only Derek's /approve is now missing


## July 5, 2022



* [decarr] [Feedback](https://github.com/kubernetes/community/pull/6725) for evaluating reviewer/approver requests
* [vinaykul] [InPlace Pod Vertical Scaling](https://github.com/kubernetes/kubernetes/pull/102884/) PR status update
    * vinaykul on vacation for the week but below is latest status:
    * Guidance for runtime is [taking shape](https://github.com/kubernetes/kubernetes/pull/102884#issuecomment-1171809939) - thanks mrunalp & kolyshkin
    * SIG-Scheduling (huang-wei) [signaled LGTM](https://github.com/kubernetes/kubernetes/pull/102884#issuecomment-1171515624) for current code.
        * E2E test and optimization can come in follow-up PRs.
    * thockin [has been LGTM](https://github.com/kubernetes/kubernetes/pull/102884#pullrequestreview-923756943) on API changes.
    * wangchen615 is working on addressing Danielle’s feedback on E2E test and adding E2E test that targets scheduler. Tracked issue [110490](https://github.com/kubernetes/kubernetes/issues/110490).
* [adrianreber] Forensic Container Checkpointing
    * code PR [https://github.com/kubernetes/kubernetes/pull/104907](https://github.com/kubernetes/kubernetes/pull/104907)
    * LGTM by Ryan, Danielle, Mike(per existing kep) .
    * Ready to be merged?
    * Derek expressed possible discussion/evaluation needed for the new exposed checkpoint service, mrunalp requested to look it over in that context.
        * This was discussed a couple of months ago in the KEP and the consensus was that for Alpha we use it as it is right now. The checkpoint archive can only be accessed by the local root user as designed in the KEP. For Beta we need to think about if we want additional authorization on the Kubelet checkpoint API endpoint.
* [swsehgal] Follow up on populating [Node Resource Topology-api](https://github.com/kubernetes/noderesourcetopology-api/) repository
    * Issue: [https://github.com/kubernetes/community/issues/6308](https://github.com/kubernetes/community/issues/6308)
    * [Response](https://docs.google.com/document/d/16GqCjnEh86w8yADcrUylNoE1y1sqjIMYNC_gdk5WPSQ/edit#) from Kubernetes Release Engineering team to follow the github workflow and solicit reviews.
    * First PR to review: [https://github.com/kubernetes/kubernetes/pull/110252](https://github.com/kubernetes/kubernetes/pull/110252)
* [ddebroy] Enhance Pod Initialized condition (for pods without init containers) in a fresh KEP
    * scoped to pods without init containers
    * will address the comment from Dawn at [https://github.com/kubernetes/enhancements/pull/3087#discussion_r904153856](https://github.com/kubernetes/enhancements/pull/3087#discussion_r904153856)
* [natalivlatko] SIG Node reviewers for Docs reviews (see Slack thread: [https://kubernetes.slack.com/archives/C0BP8PW9G/p1656915324289179](https://kubernetes.slack.com/archives/C0BP8PW9G/p1656915324289179)) 
    * [https://github.com/kubernetes/website/pulls?q=is%3Apr+is%3Aopen+sig-node-pr-reviews](https://github.com/kubernetes/website/pulls?q=is%3Apr+is%3Aopen+sig-node-pr-reviews)


## June 28th, 2022



* [vinaykul] [InPlace Pod Vertical Scaling](https://github.com/kubernetes/kubernetes/pull/102884/) PR status update (vinaykul OOO next week)
    * KEP template changes merged. KEP is now tracked for 1.25
    * [Fixed](https://github.com/kubernetes/kubernetes/commit/66f6239802aad21e2fa25cec558fa2a11bc8bc5a) CRI & test issues found by Mike and Derek respectively.
    * _[derek on 6/14] [Reviewed](https://github.com/kubernetes/kubernetes/pull/102884#pullrequestreview-1006146971) with feedback, need to clarify core behavior expectation_
        * _Kubelet -> CRI [interaction](https://github.com/kubernetes/kubernetes/pull/102884#discussion_r896965708) pattern for observing state_
            * _This assumes the runtime reports values as read from host cgroup_
        * >>> [vinaykul] Please review my [response](https://github.com/kubernetes/kubernetes/pull/102884#issuecomment-1159606963) on ResizeStatus generation.
        * >>> [vinaykul] Please review my [comment](https://github.com/kubernetes/kubernetes/pull/102884#discussion_r908049704) on suggested runtime behavior.
* [bthurber & mrunal] - special-resource-operator repo rename
    * Issue: [https://github.com/kubernetes-sigs/special-resource-operator/issues/6](https://github.com/kubernetes-sigs/special-resource-operator/issues/6) 
    * [Community meeting notes](https://docs.google.com/document/d/1b-wFATh2A0Pm1P3k11lniSeRP4q74rUa1fTvmk_ZdWY/edit?usp=sharing)
    * Derek - Write up a readme and send it to the mailing list
* [adrianreber] Forensic Container Checkpointing
    * code PR [https://github.com/kubernetes/kubernetes/pull/104907](https://github.com/kubernetes/kubernetes/pull/104907)
    * LGTM by Ryan, Danielle, Mike(per existing kep) .
    * Ready to be merged? Derek expressed possible discussion/evaluation needed for new exposed checkpoint service, mrunalp requested to look it over in that context.
* [swsehgal] Populating [Node Resource Topology-api](https://github.com/kubernetes/noderesourcetopology-api/) repository
    * Issue: [https://github.com/kubernetes/community/issues/6308](https://github.com/kubernetes/community/issues/6308)
    * PRs:
        * [https://github.com/kubernetes/kubernetes/pull/110252](https://github.com/kubernetes/kubernetes/pull/110252)
        * [https://github.com/kubernetes/kubernetes/pull/110629](https://github.com/kubernetes/kubernetes/pull/110629)
        * [https://github.com/kubernetes/kubernetes/pull/96275](https://github.com/kubernetes/kubernetes/pull/96275)
    * Action Items
        * [swsehgal] To identify if we need an enhancement proposal for this work (currently no code is being proposed in core Kubernetes only the API) and get a member of the release team to take a look at the PRs.
            * [Slack thread](https://kubernetes.slack.com/archives/C2C40FMNF/p1656440458212419) created on #sig-release channel
        * [Derek/ Sasha] To take a look at the PRs/ repo creation request to see if we want to populate the repo.


## June 21st, 2022



* [klueska] Need final approval from [Derek](mailto:decarr@redhat.com) or [Dawn](mailto:dawnchen@google.com) for following KEP
    * [CPUManager policy option to align CPUs by Socket instead of NUMA node](https://github.com/kubernetes/enhancements/pull/3334)
    * I have given my [preliminary /lgtm and /approve](https://github.com/kubernetes/enhancements/pull/3334#issuecomment-1161476253) on it already (with caveats of things we should consider before moving to beta)
* [mrunal] - KEPs needing approval / review
    * [https://github.com/kubernetes/enhancements/pull/3404](https://github.com/kubernetes/enhancements/pull/3404)
    * [https://github.com/kubernetes/enhancements/pull/3410](https://github.com/kubernetes/enhancements/pull/3410)
    * [https://github.com/kubernetes/enhancements/pull/3087](https://github.com/kubernetes/enhancements/pull/3087)
    * [https://github.com/kubernetes/enhancements/pull/3419](https://github.com/kubernetes/enhancements/pull/3419)
    * [https://github.com/kubernetes/enhancements/pull/3064](https://github.com/kubernetes/enhancements/pull/3064)
    * [https://github.com/kubernetes/enhancements/pull/2697](https://github.com/kubernetes/enhancements/pull/2697)
* [danielle] Testing/Reliability
    * General agreement that we have work to do
        * Defining unreliability:
            * “Simple”: Violating a published invariant of the kubernetes api
            * More complicated when adding plugin interfaces to the mix, and “undocumented” things around when state transitions will happen.
        * A few key areas for improvements:
            * CRI test interfaces, both testing CRIs themselves, and testing how the Kubelet will respond to CRI failure
                * Kubelet+CRI failure: [https://github.com/kubernetes/kubernetes/issues/110429](https://github.com/kubernetes/kubernetes/issues/110429)
                * Need an issue for critest in cri-tools.
                * contract testing
            * Use tests to document the invariants that exist, to avoid shipping regressions, and be aware when shipping behavioral change.
                * We don’t need to find every latent bug, but we should make it harder to ship new ones.
                * Prioritize testing behavior over timing, and then we can fixing timing later.
                * [unit tests issue] [https://github.com/kubernetes/kubernetes/issues/109717](https://github.com/kubernetes/kubernetes/issues/109717)
                * 
            * Clarifying where we’re lacking features that cause failure: [https://github.com/kubernetes/kubernetes/issues/110428](https://github.com/kubernetes/kubernetes/issues/110428)
            * As reviewers/approvers we need to ensure that new changes get test coverage.
* [alculquicondor]  Overview of [https://github.com/kubernetes/enhancements/pull/3374](https://github.com/kubernetes/enhancements/pull/3374)
    * first round reviewers: Qiutong, David
* [adrianreber] Forensic Container Checkpointing
    * code PR [https://github.com/kubernetes/kubernetes/pull/104907](https://github.com/kubernetes/kubernetes/pull/104907)
    * LGTM by Ryan, Danielle, Mike
    * Ready to be merged?


## June 14th, 2022

Total active pull requests: [192](https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+is%3Aopen+)


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
   <td><strong><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++created%3A2022-06-07T17%3A00%3A00%2B0000..2022-06-14T16%3A55%3A29%2B0000">17</a></strong>
   </td>
   <td>Closed:
   </td>
   <td><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++is%3Aunmerged+closed%3A2022-06-07T17%3A00%3A00%2B0000..2022-06-14T16%3A55%3A29%2B0000">6</a>
   </td>
  </tr>
  <tr>
   <td>Updated:
   </td>
   <td><strong><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+updated%3A2022-06-07T17%3A00%3A00%2B0000..2022-06-14T16%3A55%3A29%2B0000+created%3A%3C2022-06-07T17%3A00%3A00%2B0000">57</a></strong>
   </td>
   <td>Merged:
   </td>
   <td><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++merged%3A2022-06-07T17%3A00%3A00%2B0000..2022-06-14T16%3A55%3A29%2B0000">13</a>
   </td>
  </tr>
</table>


Fish out: [https://github.com/kubernetes/kubernetes/pull/104140/](https://github.com/kubernetes/kubernetes/pull/104140/)



* [dawnchen] [https://bit.ly/k8s125-enhancements](https://bit.ly/k8s125-enhancements) is updated for SIG Node.
    * Total: 19 enhancements
* [paco] I worked on this [https://github.com/kubernetes/enhancements/issues/1029](https://github.com/kubernetes/enhancements/issues/1029) for sometime since 1.22 and fixed a bug and tried to add metrics/log for this feature, and the promotion pr was updated [https://github.com/kubernetes/enhancements/pull/2697](https://github.com/kubernetes/enhancements/pull/2697) Can we add this to v1.25 if it met the beta-promotion bar?
* [vinaykul] [InPlace Pod Vertical Scaling](https://github.com/kubernetes/kubernetes/pull/102884/) PR status (vinaykul unavailable next two weeks)
    * [derek] [Reviewed](https://github.com/kubernetes/kubernetes/pull/102884#pullrequestreview-1006146971) with feedback, need to clarify core behavior expectation
        * Kubelet -> CRI [interaction](https://github.com/kubernetes/kubernetes/pull/102884#discussion_r896965708) pattern for observing state
            * This assumes the runtime reports values as read from host cgroup
    * KEP needs a template catch-up update. Please review [this PR](https://github.com/kubernetes/enhancements/pull/3390).
        * Partial/placeholder current unit-test cov info added. A more detailed breakdown will have to wait until after the OSS NA conference.
    * API (Tim Hockin) LGTM. Scheduling (need to add E2E test, then LGTM likely)
    * Issues to fix are being tracked [here](https://github.com/vinaykul/kubernetes/wiki/In-Place-Pod-Vertical-Scaling-Issues-and-Status). Volunteers welcome :)
* [adrianreber] Forensic Container Checkpointing (not able to join (again))
    * code PR [https://github.com/kubernetes/kubernetes/pull/104907](https://github.com/kubernetes/kubernetes/pull/104907)
        * reviews done by Mrunal and Danielle (thanks)
        * probably almost finished, waiting for additional reviews/approval
        * Waiting for feedback from Mike about CRI changes
            * Initially we targeted checkpoints archive on the local file system
            * Right now we have successfully implemented checkpoint OCI images (not standardized (yet)) in the local registry in containerd and CRI-O
            * Storing checkpoints as OCI images was a request during early discussions (1.5 years ago)
            * At this point we could completely drop checkpoints written to the local file system and only store checkpoint images in the local containerd/CRI-O registry
            * Please let me know in the PR if it would be preferred to drop the local file system checkpoint archives (I am in favor of it)
    * Concerning the PR discussion this would mean we need to keep the parameter about the checkpoint destination (something like localhost/checkpoint-image:tag) and not remove the destination parameter as suggested by Mike

		[Derek] fundamental question - PR allows to change CPU and Memory request and limit. But when the change will manifest? How kubelet will know if this errored? Should we read the value back on whether change was applied?

		[MikeB] error code from API call will notify if failed to lower the limit. Swallowed on increase.

		[Mrunal] confirm this^^^. As long as we increment properly should be fine.

		[Derek] Need to make sure kubelet always knows the latest applied values. Also in case of emptyDir &lt;missed this> 

		Also PR makes an assumption on a PLEG event being handled properly.

		[Derek] Need to check the behavior on cgroupv2 as well.



* [ddebroy] SandboxReady pod condition KEP 
    * KEP https://github.com/kubernetes/enhancements/pull/3087
        * Reviewed by Qiutong and Ruiwen
        * Looking for a review from Derek.
* [ruiwen-zhao] Adding GA criteria for KEP-2133 kubelet credential provider
    * https://github.com/kubernetes/enhancements/pull/3379
    * Reviewed/Approved by SergeyKanzhelev and deads2k
    * Looking for a review from Derek (or other sig node approvers)
* [mikebrow] exec with uid/gid (maybe user) option vs current root only.. any interest? original discussion1224— 
    * discussion centered on use cases, comparing with ephemeral container support, login with ssh plugin extension… 
    * Sergey had a good idea about a flag for disabling root defaulting
        * perhaps we could use container default here..?
* [ed] Dynamic resource allocation KEP: request for review: https://github.com/kubernetes/enhancements/pull/3064
    * Being reviewed by Tim Hockin
    * Looking for a 2nd review round from Derek
* [marquiz]  Class resources KEP, re-triage, reviewers/approver were missing last week
* [mckdev] Always set alpha.kubernetes.io/provided-node-ip https://github.com/kubernetes/kubernetes/pull/109794
* 
            * 


## June 7th, 2022

Total active pull requests:[192](https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+is%3Aopen+)

(for the past two weeks):


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
   <td><strong><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++created%3A2022-05-24T17%3A00%3A00%2B0000..2022-06-07T16%3A56%3A58%2B0000">29</a></strong>
   </td>
   <td>Closed:
   </td>
   <td><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++is%3Aunmerged+closed%3A2022-05-24T17%3A00%3A00%2B0000..2022-06-07T16%3A56%3A58%2B0000">17</a>
   </td>
  </tr>
  <tr>
   <td>Updated:
   </td>
   <td><strong><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+updated%3A2022-05-24T17%3A00%3A00%2B0000..2022-06-07T16%3A56%3A58%2B0000+created%3A%3C2022-05-24T17%3A00%3A00%2B0000">94</a></strong>
   </td>
   <td>Merged:
   </td>
   <td><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++merged%3A2022-05-24T17%3A00%3A00%2B0000..2022-06-07T16%3A56%3A58%2B0000">17</a>
   </td>
  </tr>
</table>




* [Ruiwen] Fill up [https://bit.ly/k8s125-enhancements](https://bit.ly/k8s125-enhancements) from [https://docs.google.com/document/d/1U10J0WwgWXkdYrqWGGvO8iH2HKeerQAlygnqgDgWv4E/edit#bookmark=id.qovbe39npaih](https://docs.google.com/document/d/1U10J0WwgWXkdYrqWGGvO8iH2HKeerQAlygnqgDgWv4E/edit#bookmark=id.qovbe39npaih) 
* [SergeyKanzhelev] Proposed soft freeze **July 12th**. Expectations for the soft freeze:
* Beta graduations and deprecations should be fully merged by this date
* Alpha features and GA must have PRs open that should be ready to review
* [marquiz]  [Class resources KEP](https://github.com/kubernetes/enhancements/pull/3004), follow-up on blockio
* [adrianreber] Forensic Container Checkpointing (cannot make it to today's meeting)
    * KEP [https://github.com/kubernetes/enhancements/pull/3264](https://github.com/kubernetes/enhancements/pull/3264)
        * Merged, thanks for the review and approval
    * code PR [https://github.com/kubernetes/kubernetes/pull/104907](https://github.com/kubernetes/kubernetes/pull/104907)
        * multiple review rounds (thanks Mrunal!)
        * probably almost finished, waiting for additional reviews/approval
        * Added unit test, fixed e2e test as suggested by Danielle (thanks)
* [vinaykul] [InPlace Pod Vertical Scaling](https://github.com/kubernetes/kubernetes/pull/102884/) PR status (vinaykul out today)
    * Enhancements [PR](https://github.com/kubernetes/enhancements/pull/3292) merged - thanks Dawn!
    * Derek’s re-review is in progress.
    * API (Tim Hockin) LGTM. Scheduling (need to add E2E test, then LGTM likely)
    * Issues to fix are being tracked [here](https://github.com/vinaykul/kubernetes/wiki/In-Place-Pod-Vertical-Scaling-Issues-and-Status). Volunteers welcome :)
* [ddebroy] SandboxReady pod condition KEP 
    * KEP [https://github.com/kubernetes/enhancements/pull/3087](https://github.com/kubernetes/enhancements/pull/3087)
        * Reviewed by Qiutong and Ruiwen
        * Looking for a review from Derek.
* [Peter] CRI-O and containerd CVE and long-term fix discussion
    * There is a cap on memory used that we need to document this limitation to k8s customers.
    * [mrunal] document it in CRI and document the behavior: we truncate at that point.
    * [peter] will take an action item to do this documentation.
    * CRI-O: [https://access.redhat.com/security/cve/cve-2022-1708](https://access.redhat.com/security/cve/cve-2022-1708)
    * containerd: [https://cve.mitre.org/cgi-bin/cvename.cgi?name=2022-31030](https://cve.mitre.org/cgi-bin/cvename.cgi?name=2022-31030) 
    * Followed up in [https://github.com/kubernetes/kubernetes/pull/110435](https://github.com/kubernetes/kubernetes/pull/110435) 
* [Hanamantgoud] Reviving [https://github.com/kubernetes/enhancements/pull/1224](https://st1.zoom.us/web_client/3jktxx3/html/externalLinkPage.html?ref=https://github.com/kubernetes/enhancements/pull/1224)
* 


## May 31, 2022



* [klueska]: Please add following enhancement to tracking sheet:
    * New CPU Manager Policy: align-by-socket
    * [Link to enhancement Issue](https://github.com/kubernetes/enhancements/issues/3327)
    * [Link to row in KEP - Planning document](https://docs.google.com/document/d/1U10J0WwgWXkdYrqWGGvO8iH2HKeerQAlygnqgDgWv4E/edit#bookmark=id.22jz5yk3qwbl)
* [matthyx] (cannot be present): Please remove [Keystone containers KEP](https://github.com/kubernetes/enhancements/issues/2872) from tracking:
    * adisky is on maternity leave
    * code will likely impact PLEG, will prefer to have more tests added by [sig-node: reliability project](https://docs.google.com/document/u/0/d/1jZARfxuqZ2vjs7c4Q48O5HrQu6zekOfHYl4zAIq8eak/edit) (which I want to participate to) before refactoring
* [rata]: userns [KEP PR](https://github.com/kubernetes/enhancements/pull/3275) open since early April. Anything missing?
    * The sig-node freeze is in a few days
    * AFAIK there is nothing missing. Got LGTM a few hours ago, missing /approve
* [vinaykul] [InPlace Pod Vertical Scaling](https://github.com/kubernetes/kubernetes/pull/102884/) - status update
    * Merged KEPs 2273 with 1287. Awaiting [review](https://github.com/kubernetes/enhancements/pull/3292).
    * API (Tim Hockin) LGTM. Scheduling (need to add E2E test, then LGTM likely)
    * Awaiting Derek review completion**.**
    * Issues to fix are being tracked [here](https://github.com/vinaykul/kubernetes/wiki/In-Place-Pod-Vertical-Scaling-Issues-and-Status). Volunteers welcome :)
* [adrianreber] Forensic Container Checkpointing (cannot make it to today's meeting)
    * KEP [https://github.com/kubernetes/enhancements/pull/3264](https://github.com/kubernetes/enhancements/pull/3264)
        * Reviewed, now waiting for approval
        * There is a review comment from my Mike which is not totally clear to me and I am looking for clarification what Mike was exactly asking for
    * code PR [https://github.com/kubernetes/kubernetes/pull/104907](https://github.com/kubernetes/kubernetes/pull/104907)
        * multiple review rounds (thanks Mrunal!)
        * probably almost finished, waiting for additional reviews/approval
* [bobbypage/ruiwen] issue with terminating pods reporting ready=true 
    * issue: [https://github.com/kubernetes/kubernetes/issues/108594](https://github.com/kubernetes/kubernetes/issues/108594)
    * Repro test found: [https://github.com/kubernetes/kubernetes/pull/110257](https://github.com/kubernetes/kubernetes/pull/110257)
    * PR fix: [https://github.com/kubernetes/kubernetes/pull/110256](https://github.com/kubernetes/kubernetes/pull/110256)
* [ddebroy] SandboxReady pod condition KEP 
    * KEP [https://github.com/kubernetes/enhancements/pull/3087](https://github.com/kubernetes/enhancements/pull/3087)
        * Reviewed by Qiutong and Ruiwen
        * Looking for a review from Derek.


## May 24, 2022



* [rata]: userns [KEP PR](https://github.com/kubernetes/enhancements/pull/3275) open since early April. Can we have a review?
    * We want to move to alpha (phase I) in 1.25
    * PR here: [https://github.com/kubernetes/enhancements/pull/3275](https://github.com/kubernetes/enhancements/pull/3275) 
    * Mrunal will review it this week. Mike Brown and other folks in Containerd will review the related pieces for Containerd.
* [bobbypage] Pod Lifecycle Documentation / Guarantees Doc needed
    * xref: [https://github.com/kubernetes/kubernetes/pull/110115](https://github.com/kubernetes/kubernetes/pull/110115) 
    * 1.22 vs old diff - [https://github.com/kubernetes/kubernetes/issues/109414#issuecomment-1126203054](https://github.com/kubernetes/kubernetes/issues/109414#issuecomment-1126203054) 
    * Pod Readiness issue and PR fix [https://github.com/kubernetes/kubernetes/issues/102537](https://github.com/kubernetes/kubernetes/issues/102537)
        * [https://github.com/kubernetes/kubernetes/pull/110191](https://github.com/kubernetes/kubernetes/pull/110191)
        * 
* [vinaykul] [InPlace Pod Vertical Scaling](https://github.com/kubernetes/kubernetes/pull/102884/) PR status
    * Merged KEPs 2273 with 1287. Please [review](https://github.com/kubernetes/enhancements/pull/3292).
    * API (Tim Hockin) LGTM. Scheduling (need to add E2E test, then LGTM likely)
    * Awaiting Derek to complete the kubelet review. Can we please prioritize to avoid another release slip? My time is going to be limited as June rolls around - [multiple CPFs accepted](https://ossna2022.sched.com/?searchstring=futurewei) for LF OSS conference, I have additional work of content & demo prep.
    * Issues to fix are tracked [here](https://github.com/vinaykul/kubernetes/wiki/In-Place-Pod-Vertical-Scaling-Issues-and-Status). Volunteers welcome :)
* [adrianreber] Forensic Container Checkpointing (cannot make it to today's meeting)
    * KEP [https://github.com/kubernetes/enhancements/pull/3264](https://github.com/kubernetes/enhancements/pull/3264)
        * Reviewed, now waiting for approval
    * code PR [https://github.com/kubernetes/kubernetes/pull/104907](https://github.com/kubernetes/kubernetes/pull/104907)
        * multiple review rounds (thanks Mrunal!)
        * probably almost finished, waiting for additional reviews/approval
* [ddebroy] SandboxReady pod condition KEP 
    * KEP [https://github.com/kubernetes/enhancements/pull/3087](https://github.com/kubernetes/enhancements/pull/3087)
        * Reviewed by Qiutong and Ruiwen
        * Looking for a review from Derek.
* [jan0ski] Promote AppArmor to GA KEP
    * KEP [https://github.com/kubernetes/enhancements/pull/3298](https://github.com/kubernetes/enhancements/pull/3298) 
        * Looking for confirmation from sig-node, sig-architecture, sig-auth for 1.25 target
        * Help with implementation or review is welcome :)
    * Reworked from old KEP from Sascha [https://github.com/kubernetes/enhancements/pull/1444](https://github.com/kubernetes/enhancements/pull/1444)


## May 17, 2022



* Canceled due to Kubecon


## May 10, 2022



* [Derek/Dawn] Update on sig-node reliability kickoff last week

                will upload the recording


                spent time reaching consensus on what reliability means for node - clarify    kubelet vs. runtime vs. operating system.


           Increase test coverage and then next steps.

           Calling on the community to help. 

           matthyx - discuss at kubecon 

     



* [matthyx] discuss about [Keystone containers KEP](https://github.com/kubernetes/enhancements/issues/2872)
    * matthyx will update the document and come back in a few weeks to present the milestone and design
* [knight42] review [KEP: Split stdout and stderr log stream](https://github.com/kubernetes/enhancements/pull/3289)
*            Mrunal will make a pass
* [rphillips] Evented PLEG initial work [[doc](https://docs.google.com/document/d/1GfWDoKAYiaXApxayXmPl9bcK5ewyiimFDJ966r28TUE/edit?usp=sharing)] (Points of Contacts: Ryan Phillips, Mrunal Patel, Harshal Patil)
    *  Derek - Data on perf?. Ryan - We don’t have that yet.
    * Derek - Clarification that we won’t get rid of the list entirely but be able to make the lists less frequent.
    * Mike will review 
* [mikebrow/Paul] KEP: Sub-Second Probes
    * [hackmd version of the KEP update based on review](https://hackmd.io/nsGZk3bySKOstryPoPCcag#Proposal)s
        * [ hackmd document around the various options](https://hackmd.io/WvgI0MlgTlKRbD7bcPhZLA?view). 
    * [PR draft ](https://github.com/kubernetes/kubernetes/pull/107958)
    * KEP: [Sub-Second Probes](https://github.com/kubernetes/enhancements/pull/3067) ([issue](https://github.com/kubernetes/enhancements/issues/3066)) KEP text will be updated to latest implementation/proposal, would like to have agreement on direction and move discussion to prioritization.
    * Derek: We should as a community agree on how probes should be measured, charged. Mrunal to present what Red Hat has been working on in this area.
    * Dawn: We haven’t defined how the pod has additional overhead to run probes or logs, etc. The initial step is sharing the performance benchmark to the community.
* [marquiz] follow-up discussion on  [Class resources KEP](https://github.com/kubernetes/enhancements/pull/3004).
    * Action to Markus:
        * blog post for k8s.io blog, description on what is possible now with runtimes and existing annotations
        * Come back with demo/description of how Block I/O will be utilized by user
* [vinaykul] [InPlace Pod Vertical Scaling](https://github.com/kubernetes/kubernetes/pull/102884/) PR status
    * Merged KEPs 2273 with 1287 per last week's discussion. Please [review](https://github.com/kubernetes/enhancements/pull/3292).
    * API (Tim Hockin) LGTM.
    * Awaiting Derek to complete the review. Can we please prioritize to avoid another release slip? My time is going to be limited as June rolls around - [multiple CPFs accepted](https://ossna2022.sched.com/?searchstring=futurewei) for LF OSS conference, I have additional work of content & demo prep.
    * Issues to fix are tracked [here](https://github.com/vinaykul/kubernetes/wiki/In-Place-Pod-Vertical-Scaling-Issues-and-Status). Volunteers welcome :)
* [mrunal] kubecon next week - do we keep the meeting?
* Cancelling next week. 


## May 3, 2022



* [mrunalp/ ruiwen] 1.25 planning continue: [https://docs.google.com/document/d/1U10J0WwgWXkdYrqWGGvO8iH2HKeerQAlygnqgDgWv4E/edit#bookmark=id.qovbe39npaih](https://docs.google.com/document/d/1U10J0WwgWXkdYrqWGGvO8iH2HKeerQAlygnqgDgWv4E/edit#bookmark=id.qovbe39npaih) 
* [Announcement, danielle] testing/reliability project, we'll be meeting at 8AM PST (5PM CEST) on Wednesday (more info: [https://groups.google.com/g/kubernetes-sig-node/c/-Y8TC_l7xp8/m/U9b6icAyAwAJ](https://groups.google.com/g/kubernetes-sig-node/c/-Y8TC_l7xp8/m/U9b6icAyAwAJ))
* [vinaykul] [InPlace Pod Vertical Scaling](https://github.com/kubernetes/kubernetes/pull/102884/) PR status
    * vinaykul not in today due to conflict.
    * Merged KEPs 2273 with 1287 per last week's discussion. Please [review](https://github.com/kubernetes/enhancements/pull/3292).
    * API (Tim Hockin) LGTM.is:open is:issue label:sig/node created:
    * Awaiting Derek to complete the review.
    * Issues to fix are tracked [here](https://github.com/vinaykul/kubernetes/wiki/In-Place-Pod-Vertical-Scaling-Issues-and-Status). Volunteers welcome :)
    * Can we make an early commit for v1.25?
* 


## April 26, 2022



* 1.24 retro
    * KEPs tracking: [https://docs.google.com/spreadsheets/d/1T21mUTvPh70NB2eseHjCyD4LgRjyxWI9Bd1SoP8zAwA/edit#gid=0](https://docs.google.com/spreadsheets/d/1T21mUTvPh70NB2eseHjCyD4LgRjyxWI9Bd1SoP8zAwA/edit#gid=0) 

Done (6):


<table>
  <tr>
   <td><strong>Issue Number</strong>
   </td>
   <td><strong>Name</strong>
   </td>
   <td><strong>Stage Status</strong>
   </td>
   <td><strong>Stage</strong>
   </td>
   <td><strong>Assignee</strong>
   </td>
  </tr>
  <tr>
   <td><p style="text-align: right">
281</p>

   </td>
   <td>DynamicKubeletConfig
   </td>
   <td>Removal
   </td>
   <td>
   </td>
   <td>SergeyKanzhelev
   </td>
  </tr>
  <tr>
   <td><p style="text-align: right">
688</p>

   </td>
   <td>PodOverhead
   </td>
   <td>Graduating
   </td>
   <td>Stable
   </td>
   <td>SergeyKanzhelev
   </td>
  </tr>
  <tr>
   <td><p style="text-align: right">
2133</p>

   </td>
   <td>Kubelet Credential Provider
   </td>
   <td>Graduating
   </td>
   <td>Beta
   </td>
   <td>adisky
   </td>
  </tr>
  <tr>
   <td><p style="text-align: right">
2221</p>

   </td>
   <td>Dockershim removal
   </td>
   <td>Major Change
   </td>
   <td>Stable
   </td>
   <td>SergeyKanzhelev
   </td>
  </tr>
  <tr>
   <td><p style="text-align: right">
2712</p>

   </td>
   <td>PriorityClassValueBasedGracefulShutdown
   </td>
   <td>Graduating
   </td>
   <td>Beta
   </td>
   <td>mrunalp
   </td>
  </tr>
  <tr>
   <td><p style="text-align: right">
2727</p>

   </td>
   <td>gRPC probes
   </td>
   <td>Graduating
   </td>
   <td>Beta
   </td>
   <td>SergeyKanzhelev
   </td>
  </tr>
</table>


Removed from Milestone (17):


<table>
  <tr>
   <td><strong>Issue Number</strong>
   </td>
   <td><strong>Name</strong>
   </td>
   <td><strong>Stage Status</strong>
   </td>
   <td><strong>Stage</strong>
   </td>
   <td><strong>Assignee</strong>
   </td>
  </tr>
  <tr>
   <td><p style="text-align: right">
127</p>

   </td>
   <td>User Namespaces
   </td>
   <td>Graduating+
   </td>
   <td>Alpha
   </td>
   <td>rata
   </td>
  </tr>
  <tr>
   <td><p style="text-align: right">
1287</p>

   </td>
   <td>In-place Pod Vertical Scaling
   </td>
   <td>Graduating+
   </td>
   <td>Alpha
   </td>
   <td>vinaykul
   </td>
  </tr>
  <tr>
   <td><p style="text-align: right">
1972</p>

   </td>
   <td>ExecProbeTimeout
   </td>
   <td>Graduating+
   </td>
   <td>Stable
   </td>
   <td>jackfrancis
   </td>
  </tr>
  <tr>
   <td><p style="text-align: right">
2008</p>

   </td>
   <td>Container Checkpointing (CRIU)
   </td>
   <td>Graduating+
   </td>
   <td>Alpha
   </td>
   <td>adrianreber
   </td>
  </tr>
  <tr>
   <td><p style="text-align: right">
2043</p>

   </td>
   <td>List/watch for concrete resource assignments via PodResource API
   </td>
   <td>Graduating+
   </td>
   <td>Stable
   </td>
   <td>swatisehgal
   </td>
  </tr>
  <tr>
   <td><p style="text-align: right">
2254</p>

   </td>
   <td>Cgroupsv2
   </td>
   <td>Graduating+
   </td>
   <td>Stable
   </td>
   <td>giuseppe
   </td>
  </tr>
  <tr>
   <td><p style="text-align: right">
2371</p>

   </td>
   <td>cAdvisor-less, CRI-full stats
   </td>
   <td>Graduating+
   </td>
   <td>Beta
   </td>
   <td>haircommander
   </td>
  </tr>
  <tr>
   <td><p style="text-align: right">
2400</p>

   </td>
   <td>Swap
   </td>
   <td>Graduating+
   </td>
   <td>Beta
   </td>
   <td>ehashman
   </td>
  </tr>
  <tr>
   <td><p style="text-align: right">
2413</p>

   </td>
   <td>SeccompByDefault
   </td>
   <td>Graduating
   </td>
   <td>Beta
   </td>
   <td>saschagrunert
   </td>
  </tr>
  <tr>
   <td><p style="text-align: right">
2535</p>

   </td>
   <td>Ensure Secret Pulled Images
   </td>
   <td>Graduating+
   </td>
   <td>Alpha
   </td>
   <td>mikebrow
   </td>
  </tr>
  <tr>
   <td><p style="text-align: right">
2823</p>

   </td>
   <td>Node-level pod admission handlers
   </td>
   <td>Graduating+
   </td>
   <td>Alpha
   </td>
   <td>SaranBalaji90
   </td>
  </tr>
  <tr>
   <td><p style="text-align: right">
2837</p>

   </td>
   <td>Pod level resource limits
   </td>
   <td>Graduating+
   </td>
   <td>Alpha
   </td>
   <td>n4j
   </td>
  </tr>
  <tr>
   <td><p style="text-align: right">
2872</p>

   </td>
   <td>Keystone Containers
   </td>
   <td>Graduating+
   </td>
   <td>Alpha
   </td>
   <td>adisky
   </td>
  </tr>
  <tr>
   <td><p style="text-align: right">
2902</p>

   </td>
   <td>New CPU Manager Policy: distribute-across-numa
   </td>
   <td>Graduating+
   </td>
   <td>Beta
   </td>
   <td>klueska
   </td>
  </tr>
  <tr>
   <td><p style="text-align: right">
3063</p>

   </td>
   <td>Dynamic resource allocation
   </td>
   <td>Graduating+
   </td>
   <td>Alpha
   </td>
   <td>pohly
   </td>
  </tr>
  <tr>
   <td><p style="text-align: right">
3085</p>

   </td>
   <td>Pod conditions around starting and completion of pod sandbox creation
   </td>
   <td>Graduating+
   </td>
   <td>Alpha
   </td>
   <td>ddebroy
   </td>
  </tr>
  <tr>
   <td><p style="text-align: right">
3162</p>

   </td>
   <td>Add Deallocate and PostStopContainer to device plugin API
   </td>
   <td>Graduating
   </td>
   <td>Alpha
   </td>
   <td>zvonkok
   </td>
  </tr>
</table>



        1.22 release with 24 KEPs tracked and **13 merged**


        1.23 was tracking 14 with **8 merged**


        1.24 with 23 tracked and **6 merged**


        1.23 release retro summary:


        Good:



* Planning and tracking is useful
* Soft freeze helps
* Early merges are great

        Can be better:

* Lack of reviewers and early reviews
* Lack of approver’s bandwidth
* _Things that went well_

    Notes:

* Reviewer found missing tests during the review process
* We are making progress even though things are moving slow sometimes.
* For in-place pod vertical scaling, containerd side changes are done in parallel and ready to go.
* Collaboration with the runc community is good. 
* _Things that didn’t went well_

	Notes:



* In-place vertical scaling takes long, but we are practicing caution in the review process.
* Original author moved forward last minute 
* Keystone Containers design came late 
* In-place scaling scope increased over the review process
* unit tests in different location than the code
* Syncing changes between Kubernetes and container runtime. (One side needs to cut a release first.)
* (compared to runc) Containerd community could be more proactive when cutting releases
* _AIs_
    * Investment on testability and reliability next cycle. Leadership to scope and direct these works needed.
    * Don’t accept changes without test coverage or those lack testability. Reviewers need to hold the bar.
    * Build component tests - volunteers needed
    * Clearer instructions on which folder to add tests, based on the code. Automated tool?
* [SergeyKanzhelev] KEP 1.24 retro and KEPs 1.25 planning kick-off
    * [https://docs.google.com/document/d/1U10J0WwgWXkdYrqWGGvO8iH2HKeerQAlygnqgDgWv4E/edit#bookmark=id.qovbe39npaih](https://docs.google.com/document/d/1U10J0WwgWXkdYrqWGGvO8iH2HKeerQAlygnqgDgWv4E/edit#bookmark=id.qovbe39npaih) 
* [rata]: userns KEP: [CRI changes PR](https://github.com/kubernetes/enhancements/pull/3275) open for 19 days now
    * Can we ask for a review, please? :)
    * Do we need review from Windows/VM runtimes maintainers for Alpha phase or for beta?
        * Mark R (what is the github handle?) will take a look. But it doesn’t seem like a blocker for alpha
        * rata: Also, this lives inside the Linux section of the CRI
    * Can you help us to reach out to the relevant runtime maintainers to also take a look?
    * Can we aim for alpha in 1.25
        * It is currently not listed in that section in the doc Sergey shared
        * rata: Added, thanks!
* [adrianreber] Forensic Container Checkpointing
    * KEP update PR created as suggested three weeks ago
        * [https://github.com/kubernetes/enhancements/pull/3264](https://github.com/kubernetes/enhancements/pull/3264)
        * Move from 1.24 to 1.25
        * include CRI API changes in KEP
    * Corresponding code PR updated
        * [https://github.com/kubernetes/kubernetes/pull/104907](https://github.com/kubernetes/kubernetes/pull/104907)
    * Ready for review
* [vinaykul] [InPlace Pod Vertical Scaling](https://github.com/kubernetes/kubernetes/pull/102884/) PR status
    * API (Tim Hockin) LGTM. Derek’s review is in progress.
    * Issues to fix are being tracked [here](https://github.com/vinaykul/kubernetes/wiki/In-Place-Pod-Vertical-Scaling-Issues-and-Status).
    * Can we make an early commit for v1.25?


## April 19, 2022

**Cancelled** - due to availability of leads.


## April 12, 2022

Total active pull requests: [172](https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+is%3Aopen+) (+10 from last week)


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
   <td><strong><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++created%3A2022-04-05T17%3A00%3A00%2B0000..2022-04-12T17%3A03%3A04%2B0000">16</a></strong>
   </td>
   <td>Closed:
   </td>
   <td><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++is%3Aunmerged+closed%3A2022-04-05T17%3A00%3A00%2B0000..2022-04-12T17%3A03%3A04%2B0000">5</a>
   </td>
  </tr>
  <tr>
   <td>Updated:
   </td>
   <td><strong><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+updated%3A2022-04-05T17%3A00%3A00%2B0000..2022-04-12T17%3A03%3A04%2B0000+created%3A%3C2022-04-05T17%3A00%3A00%2B0000">44</a></strong>
   </td>
   <td>Merged:
   </td>
   <td><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++merged%3A2022-04-05T17%3A00%3A00%2B0000..2022-04-12T17%3A03%3A04%2B0000">2</a>
   </td>
  </tr>
</table>




* [ehashman] Status of [https://github.com/kubernetes/kubernetes/issues/109082](https://github.com/kubernetes/kubernetes/issues/109082) ? (final node bug in the milestone)
    * Release was scheduled for Apr. 19 but there is no 1.24 release branch yet; it will likely be delayed: [https://github.com/kubernetes/sig-release/discussions/1877](https://github.com/kubernetes/sig-release/discussions/1877) 
    * Node 1.24 burndown: [https://github.com/pulls?q=org%3Akubernetes+label%3Asig%2Fnode+is%3Aopen+milestone%3Av1.24](https://github.com/pulls?q=org%3Akubernetes+label%3Asig%2Fnode+is%3Aopen+milestone%3Av1.24) 
    * Release blocker: [https://github.com/kubernetes/kubernetes/issues/108910](https://github.com/kubernetes/kubernetes/issues/108910) (go 1.18 related)
    * We believe #109082 is a regression and should block the release, Dawn to bump priority.. (Mike Brown: i’m looking again..)
* [marquiz] Intro and demo of [Class resources KEP](https://github.com/kubernetes/enhancements/pull/3004).
* [vinaykul] [InPlace Pod Vertical Scaling](https://github.com/kubernetes/kubernetes/pull/102884/) status
    * Derek’s PR review is in progress.
    * Can we make an early commit for v1.25?
* [adrianreber] Forensic Container Checkpointing (just FYI)
    * KEP update PR created as suggested last week 
        * [https://github.com/kubernetes/enhancements/pull/3264](https://github.com/kubernetes/enhancements/pull/3264)
        * Move from 1.24 to 1.25
        * include CRI API changes in KEP
    * Corresponding code PR updated
        * [https://github.com/kubernetes/kubernetes/pull/104907](https://github.com/kubernetes/kubernetes/pull/104907)
    * Ready for review
* [mweston & Alexander Kanevskiy] reminder to please review: [Kubelet Resource Plugin RFC](https://docs.google.com/document/u/0/d/1O5G4HMhfyC9AdaGai1eV5OJpCugV3vFIW19_FCuMOaY/edit?resourcekey=0-qLkKucnl3Y2wJ_WEfZPRVQ)


## April 5, 2022

Total active pull requests: [162](https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+is%3Aopen+) (-17 from two weeks)


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
   <td><strong><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++created%3A2022-03-22T17%3A00%3A00%2B0000..2022-04-05T17%3A06%3A38%2B0000">55</a></strong>
   </td>
   <td>Closed:
   </td>
   <td><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++is%3Aunmerged+closed%3A2022-03-22T17%3A00%3A00%2B0000..2022-04-05T17%3A06%3A38%2B0000">15</a>
   </td>
  </tr>
  <tr>
   <td>Updated:
   </td>
   <td><strong><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+updated%3A2022-03-22T17%3A00%3A00%2B0000..2022-04-05T17%3A06%3A38%2B0000+created%3A%3C2022-03-22T17%3A00%3A00%2B0000">122</a></strong>
   </td>
   <td>Merged:
   </td>
   <td><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++merged%3A2022-03-22T17%3A00%3A00%2B0000..2022-04-05T17%3A06%3A38%2B0000">62</a>
   </td>
  </tr>
</table>




* [ehashman] 1.24 release blockers/regressions
    * [https://github.com/kubernetes/kubernetes/issues/109082](https://github.com/kubernetes/kubernetes/issues/109082)
    * [https://github.com/kubernetes/kubernetes/issues/109182](https://github.com/kubernetes/kubernetes/issues/109182) 
    * [https://github.com/kubernetes/kubernetes/issues/109281](https://github.com/kubernetes/kubernetes/issues/109281) 
    * We need people assigned/actively working on these
* [adrianreber] Forensic Container Checkpointing
* [ & Alexander Kanevskiy] Kubelet Resource Plugin Model
    * 
    * Ask: please comment so we can figure out how to proceed
* [ddebroy] SandboxReady KEP comments [https://github.com/kubernetes/enhancements/pull/3087#issuecomment-1076779138](https://github.com/kubernetes/enhancements/pull/3087#issuecomment-1076779138) 


## March 29, 2022

No agenda, canceling to focus on code freeze.

For any urgent items for code freeze, please ping sig-node slack channel.


## March 22, 2022

Total active pull requests: [179](https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+is%3Aopen+) (+5 from last week)


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
   <td><strong><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++created%3A2022-03-15T17%3A00%3A00%2B0000..2022-03-22T16%3A54%3A37%2B0000">25</a></strong>
   </td>
   <td>Closed:
   </td>
   <td><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++is%3Aunmerged+closed%3A2022-03-15T17%3A00%3A00%2B0000..2022-03-22T16%3A54%3A37%2B0000">5</a>
   </td>
  </tr>
  <tr>
   <td>Updated:
   </td>
   <td><strong><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+updated%3A2022-03-15T17%3A00%3A00%2B0000..2022-03-22T16%3A54%3A37%2B0000+created%3A%3C2022-03-15T17%3A00%3A00%2B0000">81</a></strong>
   </td>
   <td>Merged:
   </td>
   <td><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++merged%3A2022-03-15T17%3A00%3A00%2B0000..2022-03-22T16%3A54%3A37%2B0000">15</a>
   </td>
  </tr>
</table>




* [Announcements] Code freeze
    * Reminder: Mar. 29 - next week!
* [danielle] annual report blocked on Derek (final lgtm on contributor doc)
* [ehashman] Review of KEPs targeted for release \
 \
The following KEPs are expected to be **fully merged** (Beta/Deprecations):
* [Done] 281: DynamicKubeletConfig [https://github.com/kubernetes/enhancements/issues/281](https://github.com/kubernetes/enhancements/issues/281)
* [soft cut] 2133: Kubelet Credential Provider [https://github.com/kubernetes/enhancements/issues/2133](https://github.com/kubernetes/enhancements/issues/2133)<span style="text-decoration:underline;"> </span>
        * e2e tests required for Beta: [https://github.com/kubernetes/kubernetes/pull/108651](https://github.com/kubernetes/kubernetes/pull/108651/)
        * Feature gate + API promotion to Beta: [https://github.com/kubernetes/kubernetes/pull/108847](https://github.com/kubernetes/kubernetes/pull/108847) 
* [Done] 2221: Dockershim removal [https://github.com/kubernetes/enhancements/issues/2221](https://github.com/kubernetes/enhancements/issues/2221)
* [Cut] 2371: cAdvisor-less, CRI-full stats [https://github.com/kubernetes/enhancements/issues/2371](https://github.com/kubernetes/enhancements/issues/2371)<span style="text-decoration:underline;">	</span>
        * david to update the KEP to reflect keeping in alpha
* [Cut] 2400: Swap [https://github.com/kubernetes/enhancements/issues/2400](https://github.com/kubernetes/enhancements/issues/2400)
* [in-progress] 2712: PriorityClassValueBasedGracefulShutdown [https://github.com/kubernetes/enhancements/issues/2712](https://github.com/kubernetes/enhancements/issues/2712)
        * <span style="text-decoration:underline;">[mrunal] Awaiting approve for feature</span>
* [in-progress, let’s keep in release] 2727: gRPC probes [https://github.com/kubernetes/enhancements/issues/2727](https://github.com/kubernetes/enhancements/issues/2727)

    The following KEPs should have WIPs up that are **ready for review** (Alpha/GA):

* [only docs left] 688: PodOverhead [https://github.com/kubernetes/enhancements/issues/688](https://github.com/kubernetes/enhancements/issues/688)
* [keep] 1287: In-place Pod Vertical Scaling [https://github.com/kubernetes/enhancements/issues/1287](https://github.com/kubernetes/enhancements/issues/1287)
* [review in progress] 2008: Container Checkpointing (CRIU) [https://github.com/kubernetes/enhancements/issues/2008](https://github.com/kubernetes/enhancements/issues/2008)
* [Cut] 2535: Ensure Secret Pulled Images [https://github.com/kubernetes/enhancements/issues/2535](https://github.com/kubernetes/enhancements/issues/2535)
        * [mrunal] Cut from release, pr in progress, design changes may be needed [mikebrow] nod.. extending the KEP to include phase II plans for persistence and other related items. \

* [vinaykul] InPlace Pod Vertical Scaling status
    * WIP - working on API review issues.


## March 15, 2022

Total active pull requests: [174](https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+is%3Aopen+) (+3 from last week)


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
   <td><strong><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++created%3A2022-03-08T17%3A00%3A00%2B0000..2022-03-15T16%3A58%3A26%2B0000">18</a></strong>
   </td>
   <td>Closed:
   </td>
   <td><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++is%3Aunmerged+closed%3A2022-03-08T17%3A00%3A00%2B0000..2022-03-15T16%3A58%3A26%2B0000">5</a>
   </td>
  </tr>
  <tr>
   <td>Updated:
   </td>
   <td><strong><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+updated%3A2022-03-08T17%3A00%3A00%2B0000..2022-03-15T16%3A58%3A26%2B0000+created%3A%3C2022-03-08T17%3A00%3A00%2B0000">66</a></strong>
   </td>
   <td>Merged:
   </td>
   <td><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++merged%3A2022-03-08T17%3A00%3A00%2B0000..2022-03-15T16%3A58%3A26%2B0000">15</a>
   </td>
  </tr>
</table>




* Reminder: Mar. 29 is code freeze
* [vinaykul] InPlace Pod Vertical Scaling status
    * WIP. Responding to Derek’s comments and addressing identified issues
* [danielle] as part of the sig annual report ([https://docs.google.com/document/d/1JAvi8ptbovvjSqh88378YB9irJUqWcT-PQ5YsYmdD3c/edit#](https://docs.google.com/document/d/1JAvi8ptbovvjSqh88378YB9irJUqWcT-PQ5YsYmdD3c/edit#)) we have various pieces of documentation to update across the sig that need discussion.
    * We need to document the progression ladder from new contributor -> reviewer -> approver, asking for dawn/derek to finalize the doc they were working on to move into the community repo.
    * CONTRIBUTING.md updates	
        * We need to refine our on-ramp for new contributors a little here. What do folks think is important?
            * Today that page is a pile of links, with some helpful intro to building k8s docs from dims at the bottom
* [ddebroy] Updates to [https://github.com/kubernetes/enhancements/pull/3087](https://github.com/kubernetes/enhancements/pull/3087)
    * Addressed comments/concerns from Elana and Derek so far
    * Single SandboxReady condition
* [dawnchen] Status update on OutOfCpu issue?
    * regression since 1.22.
    * Clayton is working on a fix
    * It’s close, David is testing it. Should merge soon


## March 8, 2022

Total active pull requests: [171](https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+is%3Aopen+) (+11 from last week)


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
   <td><strong><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++created%3A2022-03-01T17%3A00%3A00%2B0000..2022-03-08T17%3A54%3A24%2B0000">23</a></strong>
   </td>
   <td>Closed:
   </td>
   <td><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++is%3Aunmerged+closed%3A2022-03-01T17%3A00%3A00%2B0000..2022-03-08T17%3A54%3A24%2B0000">5</a>
   </td>
  </tr>
  <tr>
   <td>Updated:
   </td>
   <td><strong><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+updated%3A2022-03-01T17%3A00%3A00%2B0000..2022-03-08T17%3A54%3A24%2B0000+created%3A%3C2022-03-01T17%3A00%3A00%2B0000">59</a></strong>
   </td>
   <td>Merged:
   </td>
   <td><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++merged%3A2022-03-01T17%3A00%3A00%2B0000..2022-03-08T17%3A54%3A24%2B0000">10</a>
   </td>
  </tr>
</table>




* [SergeyKanzhelev] Review KEPs which are in soft freeze

    The following KEPs are expected to be **fully merged** (Beta/Deprecations):

* [Done] 281: DynamicKubeletConfig [https://github.com/kubernetes/enhancements/issues/281](https://github.com/kubernetes/enhancements/issues/281)
* [soft cut] 2133: Kubelet Credential Provider [https://github.com/kubernetes/enhancements/issues/2133](https://github.com/kubernetes/enhancements/issues/2133)
* [Done] 2221: Dockershim removal [https://github.com/kubernetes/enhancements/issues/2221](https://github.com/kubernetes/enhancements/issues/2221)
* [keep in alpha] 2371: cAdvisor-less, CRI-full stats [https://github.com/kubernetes/enhancements/issues/2371](https://github.com/kubernetes/enhancements/issues/2371)<span style="text-decoration:underline;">	</span>
        * david to update the KEP to reflect keeping in alpha
* [cut from release] 2400: Swap [https://github.com/kubernetes/enhancements/issues/2400](https://github.com/kubernetes/enhancements/issues/2400)
* [in-progress] 2712: PriorityClassValueBasedGracefulShutdown [https://github.com/kubernetes/enhancements/issues/2712](https://github.com/kubernetes/enhancements/issues/2712)
* [in-progress, let’s keep in release] 2727: gRPC probes [https://github.com/kubernetes/enhancements/issues/2727](https://github.com/kubernetes/enhancements/issues/2727)

    The following KEPs should have WIPs up that are **ready for review** (Alpha/GA):

*  [keep it] 688: PodOverhead [https://github.com/kubernetes/enhancements/issues/688](https://github.com/kubernetes/enhancements/issues/688)
* [keep] 1287: In-place Pod Vertical Scaling [https://github.com/kubernetes/enhancements/issues/1287](https://github.com/kubernetes/enhancements/issues/1287)
* [review in progress] 2008: Container Checkpointing (CRIU) [https://github.com/kubernetes/enhancements/issues/2008](https://github.com/kubernetes/enhancements/issues/2008)
* [pr in progress, design changes may be needed] 2535: Ensure Secret Pulled Images [https://github.com/kubernetes/enhancements/issues/2353](https://github.com/kubernetes/enhancements/issues/2353)
        * qq: is phase 1 valuable by itself?
        * [Derek] there may not be enough benefits in phase 1 alone. Especially if long term the logic is moving to the runtime, some benefit if users wipe imagefs on reboot of host.
        * mrunal and mike to decide whether to keep it in milestone by confirming the value pf phase 1.

	**29th March 2022**: Week 12 — [Code Freeze](https://github.com/kubernetes/sig-release/blob/master/releases/release_phases.md#code-freeze)



* [vinaykul] InPlace Pod Vertical Scaling status
    * My company work has taken priority, but I’ll start addressing Derek’s comments after this coming Friday.
    * Won’t be in today’s meeting due to conflict.
* [bobbypage] Update on OutOfCPU issue
    * Fix is in progress ([https://github.com/kubernetes/kubernetes/pull/108366](https://github.com/kubernetes/kubernetes/pull/108366)), but it is a tricky fix and need to be careful to avoid introducing new regressions
    * Uncovered another related issue with pod lifecycle refactor relating to eviction / graceful node shutdown [Will create a GH issue to track]


## Mar 1, 2022

Total active pull requests: [160](https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+is%3Aopen+)


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
   <td><strong><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++created%3A2022-02-15T17%3A00%3A00%2B0000..2022-03-01T17%3A41%3A49%2B0000">37</a></strong>
   </td>
   <td>Closed:
   </td>
   <td><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++is%3Aunmerged+closed%3A2022-02-15T17%3A00%3A00%2B0000..2022-03-01T17%3A41%3A49%2B0000">10</a>
   </td>
  </tr>
  <tr>
   <td>Updated:
   </td>
   <td><strong><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+updated%3A2022-02-15T17%3A00%3A00%2B0000..2022-03-01T17%3A41%3A49%2B0000+created%3A%3C2022-02-15T17%3A00%3A00%2B0000">76</a></strong>
   </td>
   <td>Merged:
   </td>
   <td><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++merged%3A2022-02-15T17%3A00%3A00%2B0000..2022-03-01T17%3A41%3A49%2B0000">24</a>
   </td>
  </tr>
</table>




* Announcements
    * **Reminder:** Soft freeze: Mar. 4
    * Anything that won’t make milestone, feel free to remove it now
* [derekwaynecarr] sig annual report
    * Danielle offers to help out
* [derekwaynecarr] got through most of in-place resizing pr, its big!
    * [https://github.com/kubernetes/kubernetes/pull/102884](https://github.com/kubernetes/kubernetes/pull/102884)
    * some updates are requested by vinay when he has time in interim
    * deferring all things cgroups v2 until this merges, but need to reconcile with memory qos by beta.
* [marquiz] [Class resources KEP](https://github.com/kubernetes/enhancements/pull/3004)
    * Has been in hibernation but want to take it out of draft and proceed
* [wenwu449] [https://github.com/kubernetes/kubernetes/issues/106884](https://github.com/kubernetes/kubernetes/issues/106884)
    * [https://github.com/kubernetes/kubernetes/pull/108366](https://github.com/kubernetes/kubernetes/pull/108366) also [PR#107845](https://github.com/kubernetes/kubernetes/pull/107845)
* Were autogenerated live captions helpful? We tried that out today for the first time.
    * Sentiment in chat is that they were quite helpful, and could be turned off if they were not


## February 22, 2022 [cancelled]

[dawnchen] The meeting is cancelled due to no agenda proposed. Thanks!


## February 15, 2022

Total active pull requests: [154](https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+is%3Aopen+)


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
   <td><strong><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++created%3A2022-02-08T17%3A00%3A00%2B0000..2022-02-15T17%3A57%3A59%2B0000">16</a></strong>
   </td>
   <td>Closed:
   </td>
   <td><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++is%3Aunmerged+closed%3A2022-02-08T17%3A00%3A00%2B0000..2022-02-15T17%3A57%3A59%2B0000">21</a>
   </td>
  </tr>
  <tr>
   <td>Updated:
   </td>
   <td><strong><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+updated%3A2022-02-08T17%3A00%3A00%2B0000..2022-02-15T17%3A57%3A59%2B0000+created%3A%3C2022-02-08T17%3A00%3A00%2B0000">101</a></strong>
   </td>
   <td>Merged:
   </td>
   <td><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++merged%3A2022-02-08T17%3A00%3A00%2B0000..2022-02-15T17%3A57%3A59%2B0000">17</a>
   </td>
  </tr>
</table>




* [rata] Friendly ping for [userns KEP](https://github.com/kubernetes/enhancements/pull/3065#issuecomment-1015586271)
    * Mrunal will review by EOW
* [vinaykul] In-Place Pod Vertical Scaling status
    * I have a conflict and won't be attending tomorrow’s SIG node meeting - status remains unchanged from last week:
        * @ruiwen-zhao has already [created a draft PR](https://github.com/kubernetes/kubernetes/issues/107911#issuecomment-1030439043) adding containerd CRI support for this reporting effective cgroup CPU/mem config. Thanks Ruiwen!
        * PR [https://github.com/kubernetes/kubernetes/pull/102884](https://github.com/kubernetes/kubernetes/pull/102884)
            * Awaiting Derek’s review.
* [ddebroy] Pod sandbox conditions KEP next steps
    * [https://github.com/kubernetes/enhancements/pull/3087](https://github.com/kubernetes/enhancements/pull/3087)
* [mweston] Request for continued discussion re cpu management here (also friendly):  [https://kubernetes.slack.com/archives/C0BP8PW9G/p1644201307891349](https://kubernetes.slack.com/archives/C0BP8PW9G/p1644201307891349) 
    * https://docs.google.com/document/d/1U4jjRR7kw18Rllh-xpAaNTBcPsK5jl48ZAVo7KRqkJk/edit


## February 8, 2022

Total active pull requests: [175](https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+is%3Aopen+) (+3 from last week)


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
   <td><strong><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++created%3A2022-02-01T17%3A00%3A00%2B0000..2022-02-08T17%3A56%3A52%2B0000">21</a></strong>
   </td>
   <td>Closed:
   </td>
   <td><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++is%3Aunmerged+closed%3A2022-02-01T17%3A00%3A00%2B0000..2022-02-08T17%3A56%3A52%2B0000">9</a>
   </td>
  </tr>
  <tr>
   <td>Updated:
   </td>
   <td><strong><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+updated%3A2022-02-01T17%3A00%3A00%2B0000..2022-02-08T17%3A56%3A52%2B0000+created%3A%3C2022-02-01T17%3A00%3A00%2B0000">52</a></strong>
   </td>
   <td>Merged:
   </td>
   <td><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++merged%3A2022-02-01T17%3A00%3A00%2B0000..2022-02-08T17%3A56%3A52%2B0000">10</a>
   </td>
  </tr>
</table>




* [vinaykul] In-Place Pod Vertical Scaling 
    * PR [https://github.com/kubernetes/kubernetes/pull/102884](https://github.com/kubernetes/kubernetes/pull/102884)
        * Awaiting Derek’s review.  (was out with illness, am better now ;-) will review)
    * @ruiwen-zhao has already [created a draft PR](https://github.com/kubernetes/kubernetes/issues/107911#issuecomment-1030439043) adding containerd CRI support for this feature. Thanks Ruiwen!
* [qiutongs] share the findings of a containerd issue - “failed to reserve container name”
    * Issue [https://github.com/containerd/containerd/issues/4604](https://github.com/containerd/containerd/issues/4604)
* ~~[ahg-g] [https://github.com/kubernetes/kubernetes/pull/103934](https://github.com/kubernetes/kubernetes/pull/103934) (Derek will review)~~
* [ahg-g] [https://github.com/kubernetes/kubernetes/issues/106884](https://github.com/kubernetes/kubernetes/issues/106884) 
* [ddebroy]   [KEP 2857]


## February 1, 2022

Total active pull requests: [172](https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+is%3Aopen+) (+6 from last week)


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
   <td><strong><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++created%3A2022-01-25T17%3A00%3A00%2B0000..2022-02-01T17%3A52%3A59%2B0000">22</a></strong>
   </td>
   <td>Closed:
   </td>
   <td><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++is%3Aunmerged+closed%3A2022-01-25T17%3A00%3A00%2B0000..2022-02-01T17%3A52%3A59%2B0000">3</a>
   </td>
  </tr>
  <tr>
   <td>Updated:
   </td>
   <td><strong><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+updated%3A2022-01-25T17%3A00%3A00%2B0000..2022-02-01T17%3A52%3A59%2B0000+created%3A%3C2022-01-25T17%3A00%3A00%2B0000">49</a></strong>
   </td>
   <td>Merged:
   </td>
   <td><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++merged%3A2022-01-25T17%3A00%3A00%2B0000..2022-02-01T17%3A52%3A59%2B0000">11</a>
   </td>
  </tr>
</table>


Announcements:



* KEP freeze is on Thursday!
* [vinaykul] In-Place Pod Vertical Scaling 
    * PR [https://github.com/kubernetes/kubernetes/pull/102884](https://github.com/kubernetes/kubernetes/pull/102884) in
        * Awaiting Derek’s review.
* [ddebroy] [Pod sandbox conditions KEP](https://github.com/kubernetes/enhancements/pull/3087)
    * Initial comments from Elana addressed
    * Awaiting Derek’s review before 1.24 enhancements freeze
* [adrianreber] [Forensic Container Checkpointing KEP](https://github.com/kubernetes/enhancements/pull/1990)
    * PRR approval done (thanks Elana)
    * Waiting for final approval for 1.24
* [sergeyKanzhelev] CRI versioning - timeline for v1alpha2 deprecation [https://docs.google.com/document/d/1Q0_p7xsZts6aA9xNLuXy1YLpllzCKB0skJ2HPfKdDMk/edit#heading=h.6wfgjj2um01](https://docs.google.com/document/d/1Q0_p7xsZts6aA9xNLuXy1YLpllzCKB0skJ2HPfKdDMk/edit#heading=h.6wfgjj2um01) Consumers of CRI API: [https://github.com/kubernetes-sigs/cri-tools/issues/883](https://github.com/kubernetes-sigs/cri-tools/issues/883) and [https://github.com/containerd/stargz-snapshotter/pull/323](https://github.com/containerd/stargz-snapshotter/pull/323) 


## January 25, 2022

Total active pull requests: [166](https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+is%3Aopen+) (+10 since last week)


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
   <td><strong><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++created%3A2022-01-18T17%3A00%3A00%2B0000..2022-01-25T17%3A39%3A37%2B0000">21</a></strong>
   </td>
   <td>Closed:
   </td>
   <td><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++is%3Aunmerged+closed%3A2022-01-18T17%3A00%3A00%2B0000..2022-01-25T17%3A39%3A37%2B0000">3</a>
   </td>
  </tr>
  <tr>
   <td>Updated:
   </td>
   <td><strong><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+updated%3A2022-01-18T17%3A00%3A00%2B0000..2022-01-25T17%3A39%3A37%2B0000+created%3A%3C2022-01-18T17%3A00%3A00%2B0000">52</a></strong>
   </td>
   <td>Merged:
   </td>
   <td><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++merged%3A2022-01-18T17%3A00%3A00%2B0000..2022-01-25T17%3A39%3A37%2B0000">11</a>
   </td>
  </tr>
</table>




* Announcements
    * [ehashman] PRR Freeze Jan. 27, please fill out our PRR questionnaire
    * KEP freeze is Feb. 3, next week!
* [sallyom] kubelet tracing KEP for review
    * [https://github.com/kubernetes/enhancements/pull/2832](https://github.com/kubernetes/enhancements/pull/2832)
    * [https://github.com/kubernetes/kubernetes/pull/105126](https://github.com/kubernetes/kubernetes/pull/105126)
* [ahg-g] Update on batch WG
    * [https://github.com/kubernetes/community/pull/6299](https://github.com/kubernetes/community/pull/6299)
* [vinaykul] In-Place Pod Vertical Scaling 
    * PR [https://github.com/kubernetes/kubernetes/pull/102884](https://github.com/kubernetes/kubernetes/pull/102884)
    * Enhancements PR (update target to 1.24) needs to be merged before freeze
        * See [https://github.com/kubernetes/enhancements/pull/3153](https://github.com/kubernetes/enhancements/pull/3153) 
    * Implementation PR awaiting Derek’s review.
* [vaibhav2107] Rotated container log files size not counted towards ephemeral-storage's limit on CRI-O
    * Issue [https://github.com/kubernetes/kubernetes/issues/107447](https://github.com/kubernetes/kubernetes/issues/107447)
    * Need input
    * Peter Hunt/Mrunal will take a look..
* [SergeyKanzhelev] [https://github.com/kubernetes/kubernetes/issues/107190](https://github.com/kubernetes/kubernetes/issues/107190) (please commnet on the doc [https://docs.google.com/document/d/1Q0_p7xsZts6aA9xNLuXy1YLpllzCKB0skJ2HPfKdDMk/edit#](https://docs.google.com/document/d/1Q0_p7xsZts6aA9xNLuXy1YLpllzCKB0skJ2HPfKdDMk/edit#) or the issue itself)
* [mweston & swseghal] request to add any last notes to [https://docs.google.com/document/d/1U4jjRR7kw18Rllh-xpAaNTBcPsK5jl48ZAVo7KRqkJk/](https://docs.google.com/document/d/1U4jjRR7kw18Rllh-xpAaNTBcPsK5jl48ZAVo7KRqkJk/) so we can go through and start work on documentation/clear up use cases
* [Mike Tougeron] kubelet & cAdvisor metrics bug with cgroupv2 ​​[https://kubernetes.slack.com/archives/C20HH14P7/p1642595289013400](https://kubernetes.slack.com/archives/C20HH14P7/p1642595289013400) 
    * v1.21 may be too early of the version for cgroupv2.
    * What is the problem and what are the expectations?
    * [Derek] Perhaps follow up on cgroup v2 KEP.
    * [Mrunal] Need to write blog post to raise awareness where we are on cgroupv2.


## January 18, 2022

Total active pull requests: [156](https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+is%3Aopen+) (-34 since the last meeting)


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
   <td><strong><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++created%3A2022-01-11T17%3A00%3A00%2B0000..2022-01-18T18%3A06%3A32%2B0000">14</a></strong>
   </td>
   <td>Closed:
   </td>
   <td><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++is%3Aunmerged+closed%3A2022-01-11T17%3A00%3A00%2B0000..2022-01-18T18%3A06%3A32%2B0000">35</a>
   </td>
  </tr>
  <tr>
   <td>Updated:
   </td>
   <td><strong><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+updated%3A2022-01-11T17%3A00%3A00%2B0000..2022-01-18T18%3A06%3A32%2B0000+created%3A%3C2022-01-11T17%3A00%3A00%2B0000">116</a></strong>
   </td>
   <td>Merged:
   </td>
   <td><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++merged%3A2022-01-11T17%3A00%3A00%2B0000..2022-01-18T18%3A06%3A32%2B0000">17</a>
   </td>
  </tr>
</table>




* Announcements
    * Reminder for upcoming dates:
        * PRR freeze (Jan. 27) - PRR questionnaire should be complete
        * KEP freeze (Feb. 3) - KEP should be approved and merged
    * PRR team is seeking shadows/new members: [https://groups.google.com/a/kubernetes.io/g/dev/c/OjepOATqwD4](https://groups.google.com/a/kubernetes.io/g/dev/c/OjepOATqwD4) 
* [dawnchen/ahg] batch wg: [https://github.com/kubernetes/community/pull/6299](https://github.com/kubernetes/community/pull/6299)
    * Some concerns about fragmentation for ownership of certain features [raised on Slack](https://kubernetes.slack.com/archives/CPNFRNLTS/p1641940087009300?thread_ts=1641938530.005600&cid=CPNFRNLTS)
    * Please add any comments on the PR: [https://github.com/kubernetes/community/pull/6299](https://github.com/kubernetes/community/pull/6299) 
* [vinaykul] In-Place Pod Vertical Scaling - plan for early 1.24 merge 
    * PR [https://github.com/kubernetes/kubernetes/pull/102884](https://github.com/kubernetes/kubernetes/pull/102884)
    * Enhancements PR updating target to 1.24 needs to be merged before freeze
        * [https://github.com/kubernetes/enhancements/pull/3153](https://github.com/kubernetes/enhancements/pull/3153) 
    * Alpha-blocker issues:
        * Review claims code in convertToAPIContainerStatus [breaks non-mutating guarantees](https://github.com/kubernetes/kubernetes/pull/102884#discussion_r665550094). - **My [Nov 10 response](https://github.com/kubernetes/kubernetes/pull/102884#discussion_r746572298) needs Elana’s followup**
            * It is unclear where [updates or mutates](https://github.com/kubernetes/kubernetes/pull/102884#discussion_r738886048) to any state are happening. Need a response/clarification.
    *  [NodeSwap issue](https://github.com/kubernetes/kubernetes/pull/102884#issuecomment-964440882): Not an alpha blocker (No CI test failures seen). Asked [@ichbinblau](https://github.com/ichbinblau) or [@cathyhongzhang](https://github.com/cathyhongzhang) to file tracking issues.
    * Other non-alpha-blocker issues:
        * I plan to file tracking github issues for the remaining (7-10) issues/TODOs and assign them to people that have offered to help. They can be fixed after this PR is merged most likely within the 1.24 timeframe.
* [rata] Friendly ping for [userns KEP](https://github.com/kubernetes/enhancements/pull/3065#issuecomment-1015586271)?
    * It seems there is agreement to start moving with phase 1 at least
    * Didn’t receive any reviews from Derek/Mrunal/Sergey yet
        * Mrunal: Made a pass at the enhancement
* [vinayakankugoyal/qiutongs] Ping for CRI changes for ambient capability
    * PR [https://github.com/kubernetes/kubernetes/pull/104620](https://github.com/kubernetes/kubernetes/pull/104620)
        * Mrunal: Reviewing..
    * Approved KEP [https://github.com/kubernetes/enhancements/pull/2757](https://github.com/kubernetes/enhancements/pull/2757)


## January 11, 2022

Total active pull requests: [190](https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+is%3Aopen+) (-21 since the last meeting)


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
   <td><strong><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++created%3A2022-01-04T17%3A00%3A00%2B0000..2022-01-11T17%3A49%3A58%2B0000">13</a></strong>
   </td>
   <td>Closed:
   </td>
   <td><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++is%3Aunmerged+closed%3A2022-01-04T17%3A00%3A00%2B0000..2022-01-11T17%3A49%3A58%2B0000">13</a>
   </td>
  </tr>
  <tr>
   <td>Updated:
   </td>
   <td><strong><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+updated%3A2022-01-04T17%3A00%3A00%2B0000..2022-01-11T17%3A49%3A58%2B0000+created%3A%3C2022-01-04T17%3A00%3A00%2B0000">176</a></strong>
   </td>
   <td>Merged:
   </td>
   <td><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++merged%3A2022-01-04T17%3A00%3A00%2B0000..2022-01-11T17%3A49%3A58%2B0000">26</a>
   </td>
  </tr>
</table>




* Announcements
    * 1.24 schedule finalized
    * [ehashman] Proposed date for soft node freeze: Fri. Mar. 4, 2022
        * Applies to beta/deprecations
        * Alpha/GA features must have WIPs up
        * **Action:** ehashman to send email with announcement
    * [https://github.com/kubernetes/kubernetes/pull/104143](https://github.com/kubernetes/kubernetes/pull/104143) Welcome wzshiming@ (Shiming Zhang) as SIG Node reviewer!
* [derek] conclusion on special resource operator proposal
    * see: [https://github.com/kubernetes/org/issues/3140](https://github.com/kubernetes/org/issues/3140)
    * Approved
* [ehashman] 1.24 KEP prioritization  
    * **Action:** Elana to send email requesting review/feedback and explaining the prioritization goals
* [swsehgal/fromani][heads up] PodResource API watch support to be postponed to 1.25 release due to capacity constraints. We aim to narrow down the design in the 1.24 timeframe and target the implementation in the 1.25 timeframe.
* [pacoxu] [Quotas for Ephemeral Storage #1029](https://github.com/kubernetes/enhancements/issues/1029) Fixed a bug and try to adding a metric for this feature. Not sure if it can be promoted beta in 1.24 or 1.25.(I will update the KEP if it most likely will be promoted to beta in 1.24. If not, I may update it for 1.25 or later.)
    * [Derek] There were some feature gaps that may need to be addressed before moving to beta.
* [vaibhav2107] Rotated container log files size not counted towards ephemeral-storage’s limit

        ( [https://github.com/kubernetes/kubernetes/issues/107447](https://github.com/kubernetes/kubernetes/issues/107447) )


        Discussion on the issue

    * [ehashman] Hasn’t been triaged yet. Will be looked at tomorrow during Node CI/Triage subproject
* [jackfrancis] What to do with ExecProbeTimeout during 1.24 release cycle?

        ( https://github.com/kubernetes/kubernetes/issues/99854 )

* [vinaykul] In-Place Pod Vertical Scaling - plan for early 1.24 merge 
    * PR [https://github.com/kubernetes/kubernetes/pull/102884](https://github.com/kubernetes/kubernetes/pull/102884)
    * Pod resize E2E tests have been “weakened” for alpha - now passing CI
    * Alpha-blocker issues:
        * Review claims code in convertToAPIContainerStatus [breaks non-mutating guarantees](https://github.com/kubernetes/kubernetes/pull/102884#discussion_r665550094). - **My [Nov 10 response](https://github.com/kubernetes/kubernetes/pull/102884#discussion_r746572298) needs Elana’s followup**
            * It is unclear to me what part of the code [updates or mutates](https://github.com/kubernetes/kubernetes/pull/102884#discussion_r738886048) any state. Need a response/clarification.
        * Container hash excludes Resources with in-place-resize feature-gate enabled, toggling fg can restart containers - **Fixed & Reviewed**
            * This fix seems acceptable to both Elana & Lantao but Hash annotation naming needs to be more specific. Working on it.
    *  [NodeSwap issue](https://github.com/kubernetes/kubernetes/pull/102884#issuecomment-964440882): Not an alpha blocker (No CI test failures seen). Asked [@ichbinblau](https://github.com/ichbinblau) or [@cathyhongzhang](https://github.com/cathyhongzhang) to file tracking issues.
    * Other non-alpha-blocker issues:
        * I’m fixing various issues found in reviews of API, scheduler and kubelet.
        * I’ll file tracking github issues for the remaining (7-10) issues/TODOs and assign them to people that have offered to help. They can be fixed after this PR is merged most likely within the 1.24 timeframe.
* [mweston & swsehgal] reminder to review this-continued conversation inline for cpu management cases:  [https://docs.google.com/document/d/1U4jjRR7kw18Rllh-xpAaNTBcPsK5jl48ZAVo7KRqkJk/edit](https://docs.google.com/document/d/1U4jjRR7kw18Rllh-xpAaNTBcPsK5jl48ZAVo7KRqkJk/edit) 


## January 4, 2022

Total active pull requests: [211](https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+is%3Aopen+) (+4 since the last meeting)


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
   <td><strong><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++created%3A2021-12-14T17%3A00%3A00%2B0000..2022-01-04T17%3A58%3A23%2B0000">33</a></strong>
   </td>
   <td>Closed:
   </td>
   <td><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++is%3Aunmerged+closed%3A2021-12-14T17%3A00%3A00%2B0000..2022-01-04T17%3A58%3A23%2B0000">26</a>
   </td>
  </tr>
  <tr>
   <td>Updated:
   </td>
   <td><strong><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+updated%3A2021-12-14T17%3A00%3A00%2B0000..2022-01-04T17%3A58%3A23%2B0000+created%3A%3C2021-12-14T17%3A00%3A00%2B0000">108</a></strong>
   </td>
   <td>Merged:
   </td>
   <td><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++merged%3A2021-12-14T17%3A00%3A00%2B0000..2022-01-04T17%3A58%3A23%2B0000">7</a>
   </td>
  </tr>
</table>




* Announcements
    * Release dates for 1.24
        * Cycle start next week (Jan 10)
        * Tentative release date (Apr 19)
    * Saying hi to James Laverack, Release Lead
    * 
* [mrunal] 1.24 planning
    *  
* [ddebroy] [KEP](https://github.com/kubernetes/enhancements/pull/3087) for pod sandbox creation conditions in pod status [[https://github.com/kubernetes/enhancements/pull/3087](https://github.com/kubernetes/enhancements/pull/3087)]
* [vinaykul] In-Place Pod Vertical Scaling - plan for early 1.24 merge 
    * PR [https://github.com/kubernetes/kubernetes/pull/102884](https://github.com/kubernetes/kubernetes/pull/102884)
    * Pod resize E2E tests have been “weakened” for alpha.
        * Resize success verified at cgroup instead of pod status.
        * All 31 tests are [passing](https://prow.k8s.io/view/gs/kubernetes-jenkins/pr-logs/pull/102884/pull-kubernetes-e2e-gce-alpha-features/1468102950369890304/) now.
    * Alpha-blocker issues: 
        * Container hash excludes Resources with in-place-resize feature-gate enabled, toggling fg can restart containers - **Fixed**
            * Please review [this](https://github.com/kubernetes/kubernetes/commit/a745b12136027c136e0d3dcace23db99ccdd0d71) incremental change which addresses it.
            * [Lantao] Some customers may rely on the existing label implementation, even though it wasn’t intended for that use. Want to get feedback on this.
            * Alternative: use the same, current hash field but use it to store both hashes.
            * It might be clearer to write down both hashes separately.
            * [Elana] Some concerns about version skew of labels; if one kubelet is on one version and another is on a different one, they need to know how to use the labels correctly and not accidentally break each other.
            * [Derek] Will focus in and review. People should not assume guarantees for kubelet labels.
            * [Dawn] Adding an additional label reduces complexity because we don’t have to worry about internal versioning of the label scheme.
        * Reviewer claims code in convertToAPIContainerStatus [breaks non-mutating guarantees](https://github.com/kubernetes/kubernetes/pull/102884#discussion_r665550094). - **My [Nov 10 response](https://github.com/kubernetes/kubernetes/pull/102884#discussion_r746572298) needs Elana’s followup**
            * It is unclear what part of the code [updates or mutates](https://github.com/kubernetes/kubernetes/pull/102884#discussion_r738886048) any state. Need a response/clarification.
        * Multiple reviewers have felt that the [NodeSwap issue](https://github.com/kubernetes/kubernetes/pull/102884#issuecomment-964440882) is a blocking issue. But in the Dec 07 meeting, we felt this may not be an <span style="text-decoration:underline;">alpha</span> blocker (No CI test failures seen. After I weakened resize E2E tests and all-alpha tests passed). However, we want to be sure. - **Need Elana’s input**.
            * Can we identify exact reasons why this would (or would not) be alpha blocker?
    * I plan to create issues to track other non-alpha-blocking review items and assign them to folks to fix after PR is merged. A few people have offered to contribute. With help, we should be able to nail most, if not all, of them in the upcoming release.
* [mweston & swsehgal] Request for reviewers of CPU doc here: 
    * [swsehgal] How do we make this more pluggable in the long run? Support more bespoke use cases
