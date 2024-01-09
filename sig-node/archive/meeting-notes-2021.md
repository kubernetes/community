# SIG Node Meeting Notes

## December 28, 2021 Cancelled



* [dawnchen] Cancelled


## December 21, 2021 Cancelled



* [dawnchen] Cancelled


## December 14, 2021

Total active pull requests: [207](https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+is%3Aopen+) (+6)


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
   <td><strong><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++created%3A2021-12-07T17%3A00%3A00%2B0000..2021-12-14T17%3A52%3A13%2B0000">32</a></strong>
   </td>
   <td>Closed:
   </td>
   <td><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++is%3Aunmerged+closed%3A2021-12-07T17%3A00%3A00%2B0000..2021-12-14T17%3A52%3A13%2B0000">9</a>
   </td>
  </tr>
  <tr>
   <td>Updated:
   </td>
   <td><strong><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+updated%3A2021-12-07T17%3A00%3A00%2B0000..2021-12-14T17%3A52%3A13%2B0000+created%3A%3C2021-12-07T17%3A00%3A00%2B0000">82</a></strong>
   </td>
   <td>Merged:
   </td>
   <td><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++merged%3A2021-12-07T17%3A00%3A00%2B0000..2021-12-14T17%3A52%3A13%2B0000">20</a>
   </td>
  </tr>
</table>


Highlights:

	[https://github.com/kubernetes/kubernetes/pull/97252](https://github.com/kubernetes/kubernetes/pull/97252) Dockershim was removed!



* [pohly, klueska] [Dynamic resource allocation KEP](https://github.com/kubernetes/enhancements/pull/3064) discussion
    * Presentation
    * [dawn] As a reference, the same use case was explored by us 4 years ago, and here is the proposal made by jiayingz@ from Google: [https://docs.google.com/document/d/1qKiIVs9AMh2Ua5thhtvWqOqW0MSle_RV3lfriO1Aj6U/edit#](https://docs.google.com/document/d/1qKiIVs9AMh2Ua5thhtvWqOqW0MSle_RV3lfriO1Aj6U/edit#) 
    * [Sergey] Latency in scheduling is in the design. Do we want to keep the existing API and fix similar issues as the KEP solving for it, assuming there are workloads which need low latency
    * [Patrick] this is speculation at this point. We emphasize long running jobs, not short running ones.
    * [sergey] resourceClaim has the same name as PVC, perhaps there should be explicitly called out differences table that will simplify the review. 
    * [Patric] big difference - volumes are intentionally non-configurable, resources - customizeable
    * Quota management was one of the concerns before. Plus scheduling latency
    * [Patrick] Quota logic will be managed by vendor’s addon
    * [Patrick] Implementing scheduler by storage vendor is not the ideal solution. Writing addon may be easier and more natural. Especially if you are using two vendors in the same cluster - implementing custom scheduling wouldn’t work
    * [Dawn] yes, multiple vendors is a problem. Extending scheduler is a problem that needs to be solved.
    * 1.24 may be too tight for the alpha. 
* [mrunal] 1.24 planning (ehashman OOO) 
    *  
    * Moved to January
* [vinaykul] In-Place Pod Vertical Scaling - plan for early 1.24 merge 
    * PR [https://github.com/kubernetes/kubernetes/pull/102884](https://github.com/kubernetes/kubernetes/pull/102884)
    * Pod resize E2E tests have been “weakened” for alpha.
        * Resize success verified at cgroup instead of pod status.
        * All 31 tests are [passing](https://prow.k8s.io/view/gs/kubernetes-jenkins/pr-logs/pull/102884/pull-kubernetes-e2e-gce-alpha-features/1468102950369890304/) now.
    * Alpha-blocker issues: 
        * Container hash excludes Resources with in-place-resize feature-gate enabled, toggling fg can restart containers.
            * Please review [this](https://github.com/kubernetes/kubernetes/commit/a745b12136027c136e0d3dcace23db99ccdd0d71) incremental change which addresses it.
        * Reviewer claims that code in convertToAPIContainerStatus [breaks non-mutating guarantees](https://github.com/kubernetes/kubernetes/pull/102884#discussion_r665550094).
            * It is unclear what part of the code [updates or mutates](https://github.com/kubernetes/kubernetes/pull/102884#discussion_r738886048) any state. Need a response/clarification.
        * Multiple reviewers have felt that the [NodeSwap issue](https://github.com/kubernetes/kubernetes/pull/102884#issuecomment-964440882) is a blocking issue. But in last week’s meeting we felt this may not be an <span style="text-decoration:underline;">alpha</span> blocker (No CI test failures seen after I weakened resize E2E tests and all tests passed). However, we want to be sure.
            * Can we identify exact reasons why this would (or would not) be alpha blocker?


## December 7, 2021

Total PRs: [201](https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+is%3Aopen+) (-8)


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
   <td><strong><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++created%3A2021-11-30T17%3A00%3A00%2B0000..2021-12-07T17%3A45%3A16%2B0000">15</a></strong>
   </td>
   <td>Closed:
   </td>
   <td><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++is%3Aunmerged+closed%3A2021-11-30T17%3A00%3A00%2B0000..2021-12-07T17%3A45%3A16%2B0000">10</a>
   </td>
  </tr>
  <tr>
   <td>Updated:
   </td>
   <td><strong><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+updated%3A2021-11-30T17%3A00%3A00%2B0000..2021-12-07T17%3A45%3A16%2B0000+created%3A%3C2021-11-30T17%3A00%3A00%2B0000">65</a></strong>
   </td>
   <td>Merged:
   </td>
   <td><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++merged%3A2021-11-30T17%3A00%3A00%2B0000..2021-12-07T17%3A45%3A16%2B0000">14</a>
   </td>
  </tr>
</table>




* [ehashman] Announcements
    * Release is today! 1.24 should be opening soon
    * Contributor Celebration is Dec. 16-18
        * Registration & info: [https://www.kubernetes.dev/events/kcc2021/](https://www.kubernetes.dev/events/kcc2021/) 
* [ehashman] Node 1.23 release retro
    * 1.22 retro link: 
    * 1.23 Node KEP Planning  
    * SIG Node 1.23 KEPs  
        * 8 of 14 implemented
        * 3 exceptions requested
        * 3 exceptions granted
    * _Things that went well_
        * Really liked we have a list of KEPs in a Google doc that we talk through; it’s good to see them all in one place and see who is working on what; this helped find connections and collaboration (+1)
        * Soft freeze made the process much better from a contributor experience point of view, communication around it was also great, reviewers were comfortable reviewing the load of PRs
        * Fixed a lot of CI issues, CI looks much better than previously (+1)
        * Deprecation of DynamicKubeletConfig went very smoothly, removed from tests
        * For the last 2-3 releases we’ve been doing a better job enumerating what we get done, it’s providing a good rhythm and focus everyone’s attention (+1)
        * Appreciate that approvers are able to say “I’m not familiar with this area” and hold off on merging code they’re not confident with
        * Lots of effort put into the dockershim deprecation survey and ensuring docker wasn’t broken in 1.23, much appreciated
        * Soft code freeze helped a lot with flakes not all piling up at the very end of the release
        * Beta KEPs were removed middle of the release so we didn’t scramble to get them merged last minute
        * We did a good job of coordination with the container runtimes, not just internal to Kubernetes; much work happening in containerd, CRI-O, cadvisor happening that were all well-coordinated (+1 +1 +1)
        * General sentiment of a really successful release for node
        * SIG Node is in a much better position than other SIGs, well organized (from an outside contributor)
    * _Things that could have gone better_
        * Working on logging and kubelet flags; was hard to find reviewers for PRs, spoke with one approver who wasn’t familiar (+1)
            * In an ideal situation, someone would know who to pull in for a review, but if we don’t have that person it just gets moved to the back of the queue
            * It would be nice to indicate who specializes in which areas of code; kubelet code owner structure isn’t as cleanly articulated as could be ideal
            * This has been a problem every release; we need more approvers and reviewers overall. Unfortunately, takes time to train people and we need more people/volunteers participating
        * Sometimes hard to find approver bandwidth as well
        * Last-minute hiccup with dockershim socket for CRI v1
* [ehashman, mrunal] 1.24 initial KEP planning
    *   
    * [dawn] Overstretched by lack of reviewer bandwidth; have a difficult time prioritizing, expertise for each feature is different. We have 20-30 things on our list with no clear priority.
        * Next pass, let’s sort by priority+confidence level?
* [swsehgal/@alukiano] PodResource API Watch endpoint support ([KEP](https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/2043-pod-resource-concrete-assigments) with List and Watch endpoint support was merged in 1.21 but only List endpoint (with cpuid and device topologyinfo) was implemented). Please track the watch endpoint implementation for 1.24.
        * Issue: [https://github.com/kubernetes/enhancements/issues/2043](https://github.com/kubernetes/enhancements/issues/2043)
* [pohly, klueska] [Dynamic resource allocation KEP](https://github.com/kubernetes/enhancements/pull/3064) discussion
* [rata, giuseppe] user namespaces [KEP PR](https://github.com/kubernetes/enhancements/pull/3065)
    * Would like to agree on high level and get review from who needs to review
    * We really want to reduce chances as much as possible of a last-minute showstopper
    * General notes
        * Derek/Mrunal: phases make sense, want Tim to have a look just in case
        * Sergey will review to see if we should target 1.24 or 1.25
    * Action(rata):
        * Ping Tim Hockin
        * Add Sergey as reviewers
* [vinaykul] In-Place Pod Vertical Scaling - planning for early 1.24 merge 
    * PR [https://github.com/kubernetes/kubernetes/pull/102884](https://github.com/kubernetes/kubernetes/pull/102884)
    * Alpha-blocker issue by Elana: Container hash excludes Resources with in-place-resize feature-gate enabled, toggling fg can restart containers.
        * Please review [this](https://github.com/kubernetes/kubernetes/commit/a745b12136027c136e0d3dcace23db99ccdd0d71) incremental change to PR.
    * Pod resize E2E tests have been “weakened” for alpha.
        * All 31 tests are [passing](https://prow.k8s.io/view/gs/kubernetes-jenkins/pr-logs/pull/102884/pull-kubernetes-e2e-gce-alpha-features/1468102950369890304/) now.
        * Please review [this incremental change](https://github.com/kubernetes/kubernetes/commit/eb5cc6e38853b75834ddcce65c16311c2efdbc3c) to the original E2E test.
            * This change skips spec-resources == status.resources verification for alpha but enforces it in beta.
            * It verifies everything else besides that.
            * Resize success is verified by looking at container cgroup values and comparing to expected values after resize.
        * This gives containerd time to add support for new fields in CRI.
            * Beta blocker - cannot flip switch to beta without containerd support.
    * Outstanding issues / TODOs to load-share and fix after merge?
* [zvonkok] Any comments regarding last weeks 
    * Derek: go ahead and open the issue as a placeholder for the decision to be made
* [mweston, swsehgal] looking for feedback re cpu use cases and which cases are covered in the kubelet already, but perhaps not documented.  [https://docs.google.com/document/d/1U4jjRR7kw18Rllh-xpAaNTBcPsK5jl48ZAVo7KRqkJk/edit](https://docs.google.com/document/d/1U4jjRR7kw18Rllh-xpAaNTBcPsK5jl48ZAVo7KRqkJk/edit) 


## November 30, 2021

Total PRs: [209](https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+is%3Aopen+) (+10 since last week)


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
   <td><strong><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++created%3A2021-11-23T17%3A00%3A00%2B0000..2021-11-30T17%3A51%3A28%2B0000">18</a></strong>
   </td>
   <td>Closed:
   </td>
   <td><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++is%3Aunmerged+closed%3A2021-11-23T17%3A00%3A00%2B0000..2021-11-30T17%3A51%3A28%2B0000">10</a>
   </td>
  </tr>
  <tr>
   <td>Updated:
   </td>
   <td><strong><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+updated%3A2021-11-23T17%3A00%3A00%2B0000..2021-11-30T17%3A51%3A28%2B0000+created%3A%3C2021-11-23T17%3A00%3A00%2B0000">62</a></strong>
   </td>
   <td>Merged:
   </td>
   <td><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++merged%3A2021-11-23T17%3A00%3A00%2B0000..2021-11-30T17%3A51%3A28%2B0000">4</a>
   </td>
  </tr>
</table>




* ~~[klueska] Bug fix for new feature added in 1.23 (please add to the v1.23 milestone)~~
    * [https://github.com/kubernetes/kubernetes/pull/106599](https://github.com/kubernetes/kubernetes/pull/106599)
        * [ehashman] this missed test freeze, any fixes will need to be prioritized as release blocking, merged and backported; we should wait until master reopens if this isn’t affecting CI signal
        * We decided to push this out and include it in 1.23.1
* [SergeyKanzhelev] Dockershim removal end user [survey results](https://groups.google.com/g/kubernetes-sig-node/c/97dx4Xede40): [https://docs.google.com/document/d/1DiwCRJffBjoLuDguW9auMXS2NTmkp3NAS7SGJQBu8JY/edit#heading=h.wzgwyg229djr](https://docs.google.com/document/d/1DiwCRJffBjoLuDguW9auMXS2NTmkp3NAS7SGJQBu8JY/edit#heading=h.wzgwyg229djr) 
    * Are dev tools updated? local-up-cluster.sh defaults to docker, can use containerd/cri-o with it but it’s not documented
        * kubeadm also may need some updates
        * node e2es are difficult to run locally without docker
        * Contributor documentation needs to be updated prior to removal
    * 2 questions:
        * Did we give enough notice for removal?
            * **Yes** - full year, 1.24 won’t be released until next April
        * Did we give sufficient viable options in the time between deprecation and removal?
            * We have 2 runtimes that can be adopted + instructions for migrations
            * Some people have specific monitoring agents/tools that don’t support other runtimes, they need dependencies to migrate
            * This may be beyond what SIG Node can answer
    * [dawn] We already have delayed a year given the request; What would make the people migrate? Possibly the people won’t make changes if we keep delaying because they are not ready
    * [ehashman] We just need to go and deprecate, otherwise people will not update. We need to ensure that we ourselves are prepared for that, and update everything so that we can work without dockershim, too. [+1’s from Lantao and Mrunal]
    * [danielle] It will be painful but once it’s done it’s done and we won’t have to work on it again. Mirantis is working on making dockershim work with CRI (out of tree), so they can always use that.
    * [dawn] fixing OSS scripts and e2e tests with containerd should be the blocker for the dockershim deprecation, so that OSS users should have the out-of-box experience with containerd.
    * [dawn] containerd is a straightforward migration but we didn’t want to switch over as a default because it wasn’t GA previously
    * [derek] Did we collect docker versions? Are they getting security updates?
        * We didn’t have this as a survey question
    * [derek] Are any conformance features tied to a particular runtime? One feature that doesn’t work with dockershim, e.g. cgroupsv2
    * Zoom chat dump:
        * Derek Carr to Everyone (10:15 AM)
        * for local-up cluster, it should be just setting CONTAINER_RUNTIME and CONTAINER_RUNTIME_ENDPOINT... here is an example for crio.... export CONTAINER_RUNTIME=remote
        * export CONTAINER_RUNTIME_ENDPOINT='unix:///var/run/crio/crio.sock'
        * Lantao Liu to Everyone (10:18 AM)
        * +1
        * Mrunal Patel to Everyone (10:18 AM)
        * +1
        * Mike Brown to Everyone (10:20 AM)
        * export CONTAINER_RUNTIME=remote
        * export CONTAINER_RUNTIME_ENDPOINT="unix:///run/containerd/containerd.sock"
        * sudo -E PATH=$PATH ./hack/local-up-cluster.sh
        * nod ^ same as crio
        * Elana Hashman to Everyone (10:20 AM)
        * PRs welcome, folks :)
        * zoom chat is not the best patch tracker :D
        * Mike Brown to Everyone (10:20 AM)
        * we could change local-up to work like crictl does
        * Derek Carr to Everyone (10:20 AM)
        * more just letting us all know that it works today...
        * +1 mike
        * Brian Goff to Everyone (10:21 AM)
        * There is out of tree dockershim anyway.
        * Mike Brown to Everyone (10:21 AM)
        * basically crictl uses a config or.. if not configured loops through the default socks for each container runtime
        * Jack Francis to Everyone (10:21 AM)
        * +1, much experience with the unfortunate human behavioral reality that we need to give folks a concrete incentive to migrate
        * Lantao Liu to Everyone (10:21 AM)
        * > There is out of tree dockershim anyway.
        * Yeah, there is backup for those users who can't move away immediately
        * Brian Goff to Everyone (10:24 AM)
        * Also maybe could extend support on &lt;last version w/ dockershim> rather than pushing removal.
        * *pushing out removal
        * The only version getting sec updates is 20.10
        * Jack Francis to Everyone (10:26 AM)
        * I think there is no one answer to “which version of docker are folks using”. The answer is probably “every version”.
        * Danielle Lancashire to Everyone (10:30 AM)
        * "whatever version the OS happens to package when the image was built" * every production image * every user I assume
        * Mike Brown to Everyone (10:32 AM)
        * noting … kubelet node dependencies include some number of system services/utilities (seccomp/apparmor/…) + container runtime (extern/intern in r.current) + runtime engines (runc/crun/kata/…) … It’s not just container runtime.  Also noting you can run multiple containerd instances at the same time, thus version requirement is scoped to version configured for kubelet to use.
        * Mike Brown to Everyone (10:35 AM)
        * the version of containerd configured for kubelet to use.. can be shared with the installed version of docker.. _options_
        * 
* [wzshiming] Requesting SIG Node reviewer status. due to the time zone, I may not have time to attend the meeting, but I can review PR while everyone is sleeping. :)
    * [https://github.com/kubernetes/kubernetes/pull/104143](https://github.com/kubernetes/kubernetes/pull/104143)
    * Dawn has LGTM’d, Derek will take a look
* [zvonkok] Promote SRO as a kubernetes-sigs project, joining NFD for special resource enablement 
    * Ask: want to migrate code into kubernetes-sigs under SIG Node
        * Request form for repo migration: [https://github.com/kubernetes/org/issues/new?assignees=&labels=area%2Fgithub-repo&template=repo-create.yml&title=REQUEST%3A+%3CCreate+or+Migrate%3E+%3Cgithub+repo%3E](https://github.com/kubernetes/org/issues/new?assignees=&labels=area%2Fgithub-repo&template=repo-create.yml&title=REQUEST%3A+%3CCreate+or+Migrate%3E+%3Cgithub+repo%3E) 
        * Requirements: [https://github.com/kubernetes/community/blob/master/github-management/kubernetes-repositories.md](https://github.com/kubernetes/community/blob/master/github-management/kubernetes-repositories.md) 
    * [dawn] Only concerned about latency for initialization if there is a hard requirement on operators
    * [zvonkok] At its core it’s just a helm chart, and you can use just helm
    * [Derek] Everyone should get a chance to review the slides as a next step, it would be good for people to have a place to collaborate on this
        * After a period of a week for review, we can open an issue
        * Proposed maintainers should be listed on the repo request form
    * [mikebrown] Container orchestrated devices is another CNCF project that could do this, although it’s not currently well-integrated with Kubernetes
        * [https://github.com/container-orchestrated-devices/container-device-interface](https://github.com/container-orchestrated-devices/container-device-interface)
    * [dawn] For SIG Node sponsor a project, we need to figure out the scope first, then figure out if there are the maintenaness and feasibility,etc.  We encourage the collaboration to avoid unnecessary duplication, but it’s not like we will only allow one implementation
    * There is some intersection between the two projects, so we can collaborate as well
    * [zvonkok] SRO is and was also used to enable “special resources” that are not devices, e.g. software defined storage (lustre, veritas) with out-of-tree drivers, this kinda does not fit completely into container-orchestrated-devices. 
    * [zvonkok] On slide 10 we can see that “runtime-enablement” is one of the steps in enabling a soft/hardware device this is actually where CDI fits in, it is the low-level part where SRO tries to abstract it as mentioned “do not care about the peculiarities of the platform/runtime etc.” so the CDI effort fits perfectly in the complete picture. 
    * 
* [zvonkok] Related to that, I wanted to pick up [https://github.com/kubernetes/enhancements/pull/1949](https://github.com/kubernetes/enhancements/pull/1949)
    * Feel free!
* [adrianreber] Forensic Container Checkpointing - looking for early 1.24 approval
    * KEP [https://github.com/kubernetes/enhancements/pull/1990](https://github.com/kubernetes/enhancements/pull/1990)
    * Code PR [https://github.com/kubernetes/kubernetes/pull/104907](https://github.com/kubernetes/kubernetes/pull/104907)
    * KEP and code PR are ready for 1.24 and waiting for approver feedback
    * We should start work on 1.24 planning in the next week or two
* [mweston & swsehgal] Discussion re CPU management prioritization (quick discussion to set follow-up meeting to then bring results back to sig-node)
    * Will be sharing out a document for feedback, send an email to the mailing list requesting feedback
        * Send out a doodle with the email to make scheduling a smaller group easier
    * Will bring it to a future SIG node meeting and will open a tracking issue
    * Trying to determine a roadmap for what is/isn’t supported; get supported behaviours better documented (especially low-hanging fruit), and determine a roadmap for the gaps
* [vinaykul] In-Place Pod Vertical Scaling - planning for early 1.24 merge 
    * PR [https://github.com/kubernetes/kubernetes/pull/102884](https://github.com/kubernetes/kubernetes/pull/102884)
    * Alpha-blocker issue by Elana: Container hash excludes Resources with in-place-resize feature-gate enabled, toggling fg can restart containers.
        * We should use the current implementation upon GA+1.
        * Implemented fix [early prototype](https://github.com/huawei-cloudnative/kubernetes/blob/vertical-scaling/pkg/kubelet/kuberuntime/kuberuntime_manager.go#L534). Updated PR by next week.
    * Is NodeSwap interop issue an alpha-blocker?
    * Identify any other alpha blockers.
    * 27 of 31 pod resize e2e tests fail due to missing containerd support for new CRI field (<span style="text-decoration:underline;">all tests pass with dockershim</span>)
        * Chicken-egg problem.
        * What’s a good solution (Decision made at last meeting: Nov 16):
            * <span style="text-decoration:underline;">Action for vinaykul</span>: Ensure KEP requires containerd support for CRI change - Beta blocker.
            * Adapt tests to work around lack of support for alpha.
        * &lt;Mike Brown> you could run off master the upcoming release of containerd.. or you could modify the test to support prior releases.. In the field customers may be using older versions of container runtimes, unless the requirement is now to upgrade the container runtime with the kubernetes release?
* 


## November 23, 2021

Cancelled - US Thanksgiving week.


## November 16, 2021

Total active pull requests: [199](https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+is%3Aopen+) (-13 from last meeting)


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
   <td><strong><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++created%3A2021-11-09T17%3A00%3A00%2B0000..2021-11-16T17%3A49%3A26%2B0000">38</a></strong>
   </td>
   <td>Closed:
   </td>
   <td><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++is%3Aunmerged+closed%3A2021-11-09T17%3A00%3A00%2B0000..2021-11-16T17%3A49%3A26%2B0000">13</a>
   </td>
  </tr>
  <tr>
   <td>Updated:
   </td>
   <td><strong><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+updated%3A2021-11-09T17%3A00%3A00%2B0000..2021-11-16T17%3A49%3A26%2B0000+created%3A%3C2021-11-09T17%3A00%3A00%2B0000">100</a></strong>
   </td>
   <td>Merged:
   </td>
   <td><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++merged%3A2021-11-09T17%3A00%3A00%2B0000..2021-11-16T17%3A49%3A26%2B0000">39</a>
   </td>
  </tr>
</table>


Potential fish out from closed:



* [https://github.com/kubernetes/kubernetes/pull/96364](https://github.com/kubernetes/kubernetes/pull/96364) 
* [https://github.com/kubernetes/kubernetes/pull/104651](https://github.com/kubernetes/kubernetes/pull/104651) 
* [klueska] Pending update to KEP-2902 for 1.23
    * [https://github.com/kubernetes/enhancements/pull/3009](https://github.com/kubernetes/enhancements/pull/3009)
* [vinaykul] In-Place Pod Vertical Scaling - status update
    * PR [https://github.com/kubernetes/kubernetes/pull/102884](https://github.com/kubernetes/kubernetes/pull/102884)
    * Fixed issues identified by Tim (API) and Huang-Wei (Sched)
    * Fixed race issue by Lantao (re-do KRTM + PLEG calls ContainerStatus CRI)
    * One alpha-blocker issue as per Elana: Container hash excludes Resources with in-place-resize feature-gate enabled, so toggling this fg can restart container - no solution identified
    * Is NodeSwap interop issue an alpha-blocker?
    * E2E pod resize tests fail: chicken-egg problem?
        * 27 out of 31 E2E resize tests [fail](https://prow.k8s.io/view/gs/kubernetes-jenkins/pr-logs/pull/102884/pull-kubernetes-e2e-gce-alpha-features/1460465135016480768/)
            * Reason: containerd does not (yet) implement CRI extension in this PR to return ContainerResources (mem/cpu limits, cpu req)
            * Previous code compensated for this but it has race with PLEG
        * 4 out of 31 resize tests pass only because ResizePolicy == RestartRequired, resizing memory-request only (ContainerStatus CRI support for ContainerResources not a must in these 4 test cases)
        * Is “watering down” those 27 resize tests that depend on CRI update a reasonable solution for alpha?
            * We verify that cgroup has been updated and reflects resize, but skip verifying spec.Resources == status.Resources
* [SergeyKanzhelev] Please promote this ​​[https://kubernetes.io/blog/2021/11/12/are-you-ready-for-dockershim-removal/](https://kubernetes.io/blog/2021/11/12/are-you-ready-for-dockershim-removal/) 
* [samwalke] [Feature Request: CPU Core Grouping & Assignment (Performance vs Efficiency cores) · Issue #106157 · kubernetes/kubernetes](https://github.com/kubernetes/kubernetes/issues/106157#issuecomment-966239378)
* 


## November 9, 2021

Total active pull requests: [212](https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+is%3Aopen+) (-8 from the last meeting)


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
   <td><strong><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++created%3A2021-11-02T17%3A00%3A00%2B0000..2021-11-09T18%3A05%3A57%2B0000">26</a></strong>
   </td>
   <td>Closed:
   </td>
   <td><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++is%3Aunmerged+closed%3A2021-11-02T17%3A00%3A00%2B0000..2021-11-09T18%3A05%3A57%2B0000">15</a>
   </td>
  </tr>
  <tr>
   <td>Updated:
   </td>
   <td><strong><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+updated%3A2021-11-02T17%3A00%3A00%2B0000..2021-11-09T18%3A05%3A57%2B0000+created%3A%3C2021-11-02T17%3A00%3A00%2B0000">110</a></strong>
   </td>
   <td>Merged:
   </td>
   <td><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++merged%3A2021-11-02T17%3A00%3A00%2B0000..2021-11-09T18%3A05%3A57%2B0000">19</a>
   </td>
  </tr>
</table>




* [vinaykul] In-Place Pod Vertical Scaling - status update ([https://github.com/kubernetes/kubernetes/pull/102884](https://github.com/kubernetes/kubernetes/pull/102884)).
    * Addressed review comments in Kubelet - needs review from Lantao and Elana
    * Addressed review comments Huang Wei and updated Scheduler code. He will review it this week.
* [manugupt1] Discuss KEP for node local pod admission handlers. 
    * Kep Link : [https://github.com/kubernetes/enhancements/pull/2811](https://github.com/kubernetes/enhancements/pull/2811)
    * Need to answer Derek’s questions on why not daemonsets and other kube constructs.
    * Talk more about the use-cases. 
    * There are already plugins inside the worker node but mainly there built into kubelet source code; one of the things that this KEP proposes is a plugin framework; that will help create plugins outside kubelet’s code base.
* [Sergey/Mrunal] CRI v1: [https://github.com/kubernetes/kubernetes/pull/106006](https://github.com/kubernetes/kubernetes/pull/106006) 
    * Lets’ understand if containerd timeline lines up
    * If lines up - PR is OK, otherwise we may need to postpone to 1.24
* [Sergey] [https://github.com/kubernetes/enhancements/pull/3042](https://github.com/kubernetes/enhancements/pull/3042) Tests labels clean up
* [Mark Rossetti/Ravi G] [https://testgrid.k8s.io/sig-node-kubelet#node-kubelet-serial](https://testgrid.k8s.io/sig-node-kubelet#node-kubelet-serial) 
    * Most jobs were timing out the past few days.


## November 2, 2021

Total active pull requests: [220](https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+is%3Aopen+) (+4 from the last meeting)



* Announcements
    * Code freeze is **Nov. 17** in 2 weeks!
    * Vote in the steering election if you’re eligible!
        * [https://elections.k8s.io/](https://elections.k8s.io/) voting ends **Nov 4th**
* [@rata]: user namespaces KEP
    * Agree on high level idea
    * New proposal incorporated feedback from previous discussions
    * [Slides](https://docs.google.com/presentation/d/1z4oiZ7v4DjWpZQI2kbFbI8Q6botFaA07KJYaKA-vZpg/edit#slide=id.gc6f73a04f_0_0) to explain the plan & phases proposed
* [vinaykul] In-Place Pod Vertical Scaling
    * Rebased code and caught up with latest code.
    * Need code review from Lantao & Elana
* [adrianreber] Forensic Container Checkpointing (1.24)
    * [https://github.com/kubernetes/enhancements/pull/1990](https://github.com/kubernetes/enhancements/pull/1990) 
    * Unfortunately I missed last week
    * Wanted to answer questions which came up:
        * Can this be used for things other than forensics like migration?
            * Yes, of course, we wanted to limit it to a simple single use case. But every option around checkpoint/restore/migration is possible
        * There are a bunch of outstanding questions about how networking, storage, etc. will it work on restore with multiple containers. Wanted to rein in scope
            * Basically all external resources to a container (bind mounts, network configuration) need to be available before restore
    * Related PRs:
        * [https://github.com/cri-o/cri-o/pull/4199](https://github.com/cri-o/cri-o/pull/4199)
        * [https://github.com/kubernetes/kubernetes/pull/104907](https://github.com/kubernetes/kubernetes/pull/104907)
        * [https://github.com/kubernetes-sigs/cri-tools/pull/662](https://github.com/kubernetes-sigs/cri-tools/pull/662)


## October 26, 2021

Total active pull requests: [216](https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+is%3Aopen+) (-7 from the last meeting)


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
   <td><strong><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++created%3A2021-10-19T17%3A00%3A00%2B0000..2021-10-26T16%3A43%3A06%2B0000">11</a></strong>
   </td>
   <td>Closed:
   </td>
   <td><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++is%3Aunmerged+closed%3A2021-10-19T17%3A00%3A00%2B0000..2021-10-26T16%3A43%3A06%2B0000">7</a>
   </td>
  </tr>
  <tr>
   <td>Updated:
   </td>
   <td><strong><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+updated%3A2021-10-19T17%3A00%3A00%2B0000..2021-10-26T16%3A43%3A06%2B0000+created%3A%3C2021-10-19T17%3A00%3A00%2B0000">88</a></strong>
   </td>
   <td>Merged:
   </td>
   <td><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++merged%3A2021-10-19T17%3A00%3A00%2B0000..2021-10-26T16%3A43%3A06%2B0000">11</a>
   </td>
  </tr>
</table>




* 1.23 KEPs review (soft deadline is reached)

From [https://docs.google.com/spreadsheets/d/1P1J1QpayRmh2SNjs8T-wBCb6SgEOdWTRQ7MBol7yibk/edit#gid=936265414](https://docs.google.com/spreadsheets/d/1P1J1QpayRmh2SNjs8T-wBCb6SgEOdWTRQ7MBol7yibk/edit#gid=936265414) 


<table>
  <tr>
   <td><p style="text-align: right">
<a href="https://github.com/kubernetes/enhancements/issues/277">277</a></p>

   </td>
   <td>Beta
   </td>
   <td>Ephemeral Containers
   </td>
   <td>PRs merged
   </td>
  </tr>
  <tr>
   <td><p style="text-align: right">
<a href="https://github.com/kubernetes/enhancements/issues/1287">1287</a></p>

   </td>
   <td>Alpha
   </td>
   <td>In-place Pod update
   </td>
   <td>PR is out: <a href="https://github.com/kubernetes/kubernetes/pull/102884">kubernetes/kubernetes#102884</a> 
   </td>
  </tr>
  <tr>
   <td><p style="text-align: right">
<a href="https://github.com/kubernetes/enhancements/issues/1977">1977</a></p>

   </td>
   <td>-
   </td>
   <td>ContainerNotifier
   </td>
   <td>Removed from milestone
   </td>
  </tr>
  <tr>
   <td><p style="text-align: right">
<a href="https://github.com/kubernetes/enhancements/issues/2040">2040</a></p>

   </td>
   <td>Beta
   </td>
   <td>Kubelet CRI Support
   </td>
   <td>PR is out <a href="https://github.com/kubernetes/kubernetes/pull/104575">kubernetes/kubernetes#104575</a> 
   </td>
  </tr>
  <tr>
   <td><a href="https://github.com/kubernetes/enhancements/issues/2133">2133</a>
   </td>
   <td>Beta
   </td>
   <td>Kubelet credential provider
   </td>
   <td>WIP PR is out. Unit tests are failing: <a href="https://github.com/kubernetes/kubernetes/pull/105624">kubernetes/kubernetes#105624</a> . 
<p>
Suggest to remove from milestone
<p>
How does this impact the timeline for spinning cloud providers out of tree?
   </td>
  </tr>
  <tr>
   <td><p style="text-align: right">
<a href="https://github.com/kubernetes/enhancements/issues/2273">2273</a></p>

   </td>
   <td>Alpha
   </td>
   <td>VPA CRI Changes
   </td>
   <td>Rolled into <a href="https://github.com/kubernetes/kubernetes/pull/102884">kubernetes/kubernetes#102884</a> 
   </td>
  </tr>
  <tr>
   <td><p style="text-align: right">
<a href="https://github.com/kubernetes/enhancements/issues/2400">2400</a></p>

   </td>
   <td>Beta
   </td>
   <td>Node system swap support
   </td>
   <td>Will spill into 1.24.
<p>
Remove from milestone - no support in runtimes, also failing tests
   </td>
  </tr>
  <tr>
   <td><p style="text-align: right">
<a href="https://github.com/kubernetes/enhancements/issues/2403">2403</a></p>

   </td>
   <td>Beta
   </td>
   <td>Extend podresources API to report allocatable resources
   </td>
   <td>PR merged
   </td>
  </tr>
  <tr>
   <td><p style="text-align: right">
<a href="https://github.com/kubernetes/enhancements/issues/2535">2535</a></p>

   </td>
   <td>Alpha
   </td>
   <td>Ensure Secret Pulled Images
   </td>
   <td>PR is out: <a href="https://github.com/kubernetes/kubernetes/pull/94899">kubernetes/kubernetes#94899</a> 
   </td>
  </tr>
  <tr>
   <td><p style="text-align: right">
<a href="https://github.com/kubernetes/enhancements/issues/2625">2625</a></p>

   </td>
   <td>Beta
   </td>
   <td>Add options to reject non SMT-aligned workload
   </td>
   <td>PR is merged
   </td>
  </tr>
  <tr>
   <td><p style="text-align: right">
<a href="https://github.com/kubernetes/enhancements/issues/2712">2712</a></p>

   </td>
   <td>Alpha
   </td>
   <td>PriorityClassValueBasedGracefulShutdown
   </td>
   <td>PR is out: <a href="https://github.com/kubernetes/kubernetes/pull/102915">kubernetes/kubernetes#102915</a> 
   </td>
  </tr>
  <tr>
   <td><p style="text-align: right">
<a href="https://github.com/kubernetes/enhancements/issues/2727">2727</a></p>

   </td>
   <td>Alpha
   </td>
   <td>Add gRPC probe to Pod
   </td>
   <td>PR is out: <a href="https://github.com/kubernetes/kubernetes/pull/102162">kubernetes/kubernetes#102162</a> 
   </td>
  </tr>
  <tr>
   <td><p style="text-align: right">
<a href="https://github.com/kubernetes/enhancements/issues/2902">2902</a></p>

   </td>
   <td>Alpha
   </td>
   <td>Add CPUManager policy option to distribute CPUs across NUMA nodes instead of packing them
   </td>
   <td>PR is merged. Need to make sure to add it to release tracking sheet.
   </td>
  </tr>
  <tr>
   <td><p style="text-align: right">
<a href="https://github.com/kubernetes/enhancements/issues/2371">2371</a></p>

   </td>
   <td>Alpha
   </td>
   <td>cAdvisor-less, CRI-full Container and Pod Stats
   </td>
   <td>PR is out <a href="https://github.com/kubernetes/kubernetes/pull/103095">https://github.com/kubernetes/kubernetes/pull/103095</a> 
   </td>
  </tr>
</table>




* [Marlow Weston] Looking to solve cpu allocation items, where pinning is important, but also some pods may want to have *some* cores pinned and *some others* shared.  Currently semi-abandoned items are here;  ([https://github.com/kubernetes/community/pull/2435](https://github.com/kubernetes/community/pull/2435),[ https://github.com/kubernetes/enhancements/pull/1319](https://github.com/kubernetes/enhancements/pull/1319))  \
Would like to move forward and put together a team to try again to come up with a cohesive KEP.  Already have use cases from four companies, and there are likely more that wish to be involved.  Would like an initial list of names, in addition to those already there, so we can critically come up with the use cases we expect to cover and start coming up with what the best design is going forward.
    * Some of these use cases are already supported - are we trying to isolate the management parts of a node or specific pods?
    * If there is a set of use cases, we might be able to pull a group together
* [adrianreber] Looking for approver feedback on the "Forensic Container Checkpoint" KEP (1.24) [https://github.com/kubernetes/enhancements/pull/1990](https://github.com/kubernetes/enhancements/pull/1990) to be ready once the 1.24 period starts to avoid missing the deadline as I did for 1.23.
    * Dawn wants to review but doesn’t have bandwidth at this moment, she will approve
    * Can this be used for things other than forensics like migration?
    * There are a bunch of outstanding questions about how networking, storage, etc. will it work on restore with multiple containers. Wanted to rein in scope
* [vinaykul] In-place Pod Vertical Scaling
    * Working on addressing remaining review items this week. Apologies for the delay as my new role has not given me a big enough block of “focus time” until last week. I’ll reach out to reviewers over slack once I update the PR.
    * [kubernetes/kubernetes#102884](https://github.com/kubernetes/kubernetes/pull/102884) 
    * Should PR be split? 
        * Perhaps. SIG-scheduling would prefer scheduler changes in a separate PR. Jiaxin Shan from Bytedance is looking into it in parallel.
        * If we split, both would need to go in quick succession - 2-phase commit
* [mrunal/mikebrown] CRI PR version support [https://github.com/kubernetes/kubernetes/pull/104575](https://github.com/kubernetes/kubernetes/pull/104575) 
    * Support both v1 and v1alpha1 or just one version?
    * Node overhead for the marshalling/unmarshalling needs to be measured to make a decision.
    * Perf vs. cognitive load is a decision to make on the next meeting
    * Also added to API reviews agenda:  


## October 19, 2021

Total active pull requests: [223](https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+is%3Aopen+) (-6 in two past weeks)


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
   <td><strong><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++created%3A2021-10-05T17%3A00%3A00%2B0000..2021-10-19T16%3A49%3A17%2B0000">37</a></strong>
   </td>
   <td>Closed:
   </td>
   <td><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++is%3Aunmerged+closed%3A2021-10-05T17%3A00%3A00%2B0000..2021-10-19T16%3A49%3A17%2B0000">15</a>
   </td>
  </tr>
  <tr>
   <td>Updated:
   </td>
   <td><strong><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+updated%3A2021-10-05T17%3A00%3A00%2B0000..2021-10-19T16%3A49%3A17%2B0000+created%3A%3C2021-10-05T17%3A00%3A00%2B0000">121</a></strong>
   </td>
   <td>Merged:
   </td>
   <td><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++merged%3A2021-10-05T17%3A00%3A00%2B0000..2021-10-19T16%3A49%3A17%2B0000">29</a>
   </td>
  </tr>
</table>




* [marga] Possible KEP? Exposing OCI hooks up the stack. [Draft here](https://github.com/kinvolk/kubernetes-enhancements/tree/marga-kinvolk/oci-hooks/keps/sig-node/NNNN-expose-oci-hooks#summary). Inspektor Gadget has wanted this for a long time (see “Future” section from 2019), but now there are more users out there that are resorting to workarounds because these are not exposed. I’d like to get this moving as a KEP in y’all agree.(note [NRI](https://github.com/containerd/nri) effort)
    * [Mike Brown] lots of interest. NRI hooks is one of the models
    * [Mrunal Patel] Hooks are wired in CRI-O. Example is nvidia hook. Thai is hard to write CDI is an effort to simplify hooks writing. Hooks are hard to write, so CDI is going for a declarative model to make the actual hook much simpler. 
    * [Alexander Kanevskiy] Prototype that make NRI more flexible and working both for containerd and CRI-O: [https://github.com/containerd/nri/pull/16](https://github.com/containerd/nri/pull/16) Injecting OCI-hooks was one of the TODO items as example of NRI plugin \
CDI link: [https://github.com/container-orchestrated-devices/container-device-interface](https://github.com/container-orchestrated-devices/container-device-interface)
    * [Dawn Chen] Lots of use cases are useful, but hooks are very powerful and some hook implementations may harm the host and raise security concerns. In the past we wanted to continue discussing, maybe more declarative way will work best.
    * [Lantao] Hook can run anything on host - has access to host and container environment. Exposing in k8s API or CRI - it can make the pod non-portable. Ideally we need to avoid the environment-specific dependencies
    * [Mike Brown] pre-determined dependencies would be one way to go
    * [Dawn Chen] Are all use cases around tracing and obtaining labels and other runtime information?
    * [Marga] Main use case is tracing and applying the labels. Want to detect container start as early as possible. Other use cases for security - being able to detect container before it started and decline based on whitelist set of images.
    * [Elana] Faster detection of containers - PLEG streaming as oppose to a poll loop (Mike: yes CRI event subscriptions for pod/container/image lifecycle)
    * [Marga] Other scenarios - some information like PID of a container - would like to get, but there is no way (perhaps should be a different KEP).
    * [Dawn] Container information is a scenario for runtime. 
    * [Marga] want it in a unified way.
    * [Mrunal] There are edge cases - VM runtime, CRI-O may not have a process
    * [Sergey] Maybe expand motivation section to explain specific scenarios
    * [Dawn] need to discuss use cases individually. Some may be discussed with sig auth due to security concerns. Hooks may be a right technology, but we need to start with the use cases.
    * [Elana] Very large scope change as it’s written. Need significantly more resources to ensure this lands and maintained. Not a “1-2 PR and then you’re done” sort of change.
    * [Mike Brown] There’s no common way to configure these hooks based on each container runtime. Kubernetes integration work would be quite complicated, there’s no obvious generic way to do this right now, it’s WIP.
    * [Danielle] Every time that piece of code is touched, it leads to regressions as we don’t have a good coverage there.
    * [Brian Goff] Exposing OCI semantics on CRI seems messy.

        As opposed to configuring the runtime to run hooks


        Run anything not on the host *with root privileges


        My gut says modify (or make it possible to configure) the runtime to do what you need rather than changing the API.


        I don’t think a hook could deny based on image.

    * [Eric Ernst] Second the security concern.  Chance to run a binary on the host (or in guest if kata) w/ out any restrictions...
    * [Mrunal Patel] There are races in resolving a tag at admission vs. runtime pulling it.
    * [Alexander Kanevskiy] For anyone interested in CDI and NRI, and see if those can solve their usecases, welcome to join meetings of CNCF/TAG-Runtime/COD WG:  [https://github.com/cncf/tag-runtime/blob/master/wg/COD.md](https://github.com/cncf/tag-runtime/blob/master/wg/COD.md) 
* [marquiz] Class-based resources in CRI KEP (draft) \
[https://github.com/kubernetes/enhancements/pull/3004](https://github.com/kubernetes/enhancements/pull/3004)
    * Looking for feedback
    * Related KubeCon talk: [https://sched.co/m4t2](https://sched.co/m4t2)
    * TAG-Runtime presentation with examples on how it is used: 
    * Annotations should generally be avoided for new KEPs as they are very difficult to manage with version skew, use alpha fields instead
    * Are we doing anything to expose this information to scheduler?
    * Not at the moment, optimizing everything on the node now.
    * What’s the difference between blockio and cgroup v2 controls?
    * RDT is about memory bandwidth, it uses cgroup v2 underneath.
    * Classes are used to configure specific controls to scenarios like guaranteed burst or levels of IO support.
    * Need to make sure cgroup v2 support in k8s will work well with this proposal.
* [Eric Ernst] Request for feedback/review: [https://github.com/kubernetes/kubernetes/pull/104886](https://github.com/kubernetes/kubernetes/pull/104886) 
* [SergeyKanzhelev] Please distribute this form: [https://forms.gle/svCJmhvTv78jGdSx8](https://forms.gle/svCJmhvTv78jGdSx8)
* [SergeyKanzhelev] Follow up on this [https://github.com/kubernetes/kubernetes/pull/105215#issuecomment-946916830](https://github.com/kubernetes/kubernetes/pull/105215#issuecomment-946916830)


## October 12, 2021

CANCELLED - KUBECON NA


## October 5, 2021

LATE START: 30m shifted due to availability (10:30 instead of 10PT)

Total PRs: [229](https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+is%3Aopen+) (-1 from the last meeting)


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
   <td><strong><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++created%3A2021-09-28T17%3A00%3A00%2B0000..2021-10-05T17%3A21%3A41%2B0000">23</a></strong>
   </td>
   <td>Closed:
   </td>
   <td><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++is%3Aunmerged+closed%3A2021-09-28T17%3A00%3A00%2B0000..2021-10-05T17%3A21%3A41%2B0000">9</a>
   </td>
  </tr>
  <tr>
   <td>Updated:
   </td>
   <td><strong><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+updated%3A2021-09-28T17%3A00%3A00%2B0000..2021-10-05T17%3A21%3A41%2B0000+created%3A%3C2021-09-28T17%3A00%3A00%2B0000">85</a></strong>
   </td>
   <td>Merged:
   </td>
   <td><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++merged%3A2021-09-28T17%3A00%3A00%2B0000..2021-10-05T17%3A21%3A41%2B0000">15</a>
   </td>
  </tr>
</table>




* [ehashman] Reminder: soft code freeze ~Oct. 15, KubeCon next week
    * SIG Node session: [https://sched.co/lV9D](https://sched.co/lV9D) 
    * Pushing soft code freeze back one week to not conflict with KubeCon (-> Oct. 22)
* [SergeyKanzhelev] Dockershim deprecation status [https://docs.google.com/document/d/1DiwCRJffBjoLuDguW9auMXS2NTmkp3NAS7SGJQBu8JY/edit#](https://docs.google.com/document/d/1DiwCRJffBjoLuDguW9auMXS2NTmkp3NAS7SGJQBu8JY/edit#) 
* [pacoxu] [https://github.com/kubernetes/kubernetes/pull/104186](https://github.com/kubernetes/kubernetes/pull/104186) My application to be a sig-node reviewer has not been responded to after weeks. I am not sure what that means. Does it mean no objections or declined? And I wondered how I could make it happen. At least, I want to get some advice on how to be a qualified reviewer.
    * [lubomir/sig-cluster-lifecycle] +1 for pacoxu, who is an involved contributor in multiple areas of the project.
    * Working on a doc amongst Node subproject approvers, got caught up in that.
    * Paco should meet the criteria for reviewer so Derek and Mrunal have +1’d, thank you for your patience!
    * If we can get an ack, we’ll share the doc soon -- couple open comments remain, regret not having shared earlier.
* [SergeyKanzhelev] SIG Node test tags cleanup: [https://docs.google.com/document/d/19HqSyrS-4pyubqTvQV0hJKt_97nbCSt_aD0soL-RGGE/edit?hl=en&forcehl=1#](https://docs.google.com/document/d/19HqSyrS-4pyubqTvQV0hJKt_97nbCSt_aD0soL-RGGE/edit?hl=en&forcehl=1#) 
* [ehashman] Soft code freeze PR review?
    * Beta
        * KEP 277:	Ephemeral Containers
            * 
        * KEP 2040:	Kubelet CRI Support
            * 
        * KEP 2133:	Kubelet credential provider
            * 
        * KEP 2400:	Node system swap support
            * 
        * KEP 2403:	Extend podresources API to report allocatable resources
            * [https://github.com/kubernetes/kubernetes/pull/103289](https://github.com/kubernetes/kubernetes/pull/103289)
            * [https://github.com/kubernetes/kubernetes/pull/97415](https://github.com/kubernetes/kubernetes/pull/97415)
            * [https://github.com/kubernetes/kubernetes/pull/105003](https://github.com/kubernetes/kubernetes/pull/105003) 
                * [fromani] depends on the first two PRs
        * KEP 2625:	Add options to reject non SMT-aligned workload
            * [https://github.com/kubernetes/kubernetes/pull/105012](https://github.com/kubernetes/kubernetes/pull/105012) 
                * [fromani] missing features review/approval
    * Alpha
        * KEP 1287:	In-place Pod update
            * [https://github.com/kubernetes/kubernetes/pull/102884](https://github.com/kubernetes/kubernetes/pull/102884) 
        * KEP 2273:	VPA CRI Changes
            * 
        * KEP 2371:	cAdvisor-less, CRI-full Container and Pod stats
            * [https://github.com/kubernetes/kubernetes/pull/103095](https://github.com/kubernetes/kubernetes/pull/103095) 
        * KEP 2535:	Ensure Secret Pulled Images
            * [https://github.com/kubernetes/kubernetes/pull/94899](https://github.com/kubernetes/kubernetes/pull/94899) 
        * KEP 2712:	PriorityClassValueBasedGracefulShutdown
            * 
        * KEP 2727:	Add gRPC probe to Pod
            * [https://github.com/kubernetes/kubernetes/pull/102162](https://github.com/kubernetes/kubernetes/pull/102162) 
        * KEP 2902:	CPUManager policy option to distribute CPUs across NUMA nodes
            * [fromani] I think this work depends on changes part of 2625


## September 28, 2021

Total active pull requests: [230](https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+is%3Aopen+) (-8 since last meeting 2 weeks back)


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
   <td><strong><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++created%3A2021-09-14T17%3A00%3A00%2B0000..2021-09-28T16%3A35%3A07%2B0000">29</a></strong>
   </td>
   <td>Closed:
   </td>
   <td><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++is%3Aunmerged+closed%3A2021-09-14T17%3A00%3A00%2B0000..2021-09-28T16%3A35%3A07%2B0000">23</a>
   </td>
  </tr>
  <tr>
   <td>Updated:
   </td>
   <td><strong><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+updated%3A2021-09-14T17%3A00%3A00%2B0000..2021-09-28T16%3A35%3A07%2B0000+created%3A%3C2021-09-14T17%3A00%3A00%2B0000">146</a></strong>
   </td>
   <td>Merged:
   </td>
   <td><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++merged%3A2021-09-14T17%3A00%3A00%2B0000..2021-09-28T16%3A35%3A07%2B0000">18</a>
   </td>
  </tr>
</table>


Reminder:  **[soft code freeze date for](https://groups.google.com/g/kubernetes-sig-node/c/LTONXHu9R9w/m/PSybRudTAgAJ?utm_medium=email&utm_source=footer) SIG Node on October 15** - 2 weeks including a week of KubeCon.



* [wgahnagl] enable ~~static pods ~~pinned images
    * [https://github.com/kubernetes/kubernetes/pull/103299](https://github.com/kubernetes/kubernetes/pull/103299) 
    * Do we need to open an enhancement for this?
* ~~[pacoxu] is there [an alternative proposed](https://github.com/kubernetes/enhancements/issues/281#issuecomment-909152556) for DynamicKubeletConfig? ~~
    * ~~Can someone add some comments in the [enhancement issue](https://github.com/kubernetes/enhancements/issues/281#)? I go back to the [weekly meeting on April 6th](https://www.youtube.com/watch?v=RlNm61bpdcU&list=PL69nYSiGNLP1wJPj5DYWXjiArF-MJ5fNG&index=30), and it seems that dims/dawn agree with the decision. In multi-cluster management, it would be an important feature to change kubelet configuration easily from cluster wide.~~
        * [https://github.com/kubernetes/enhancements/issues/281#issuecomment-929378712](https://github.com/kubernetes/enhancements/issues/281#issuecomment-929378712) 
    * ~~Do we have a plan to make the `/var/lib/kubelet/config.yaml` supporting reload like ‘coredns reload plugin’? Asked in #slack as I may have no time to attend the meeting.~~
* [adrianreber] Checkpoint/Restore KEP and PR ready for review/approval
    * [https://github.com/kubernetes/enhancements/pull/1990](https://github.com/kubernetes/enhancements/pull/1990)
    * [https://github.com/kubernetes/kubernetes/pull/104907](https://github.com/kubernetes/kubernetes/pull/104907)
* [liggitt] treatment of "not ready" pods by eviction / PDB
    * Bug about not being able to evict unhealthy pods: [https://github.com/kubernetes/kubernetes/issues/72320](https://github.com/kubernetes/kubernetes/issues/72320) 
    * Proposal to always allow evicting "not ready" pods: [https://github.com/kubernetes/kubernetes/pull/105296](https://github.com/kubernetes/kubernetes/pull/105296)
    * sig-node thread: [https://groups.google.com/g/kubernetes-sig-node/c/IGHDGECSTWw](https://groups.google.com/g/kubernetes-sig-node/c/IGHDGECSTWw)
    * Chat dump:

        David Eads to Everyone (10:22 AM)


        it is impossible to available=true and ready=false


        Derek Carr to Everyone (10:25 AM)


        available has no meaning on the pod spec, only in workload controller.


        David Eads to Everyone (10:29 AM)


        to the point of a kubelet readyness being set to false bug, it seems like we should fix both the bugs.


        Elana Hashman to Everyone (10:31 AM)


        David, I'll jump in on that after this


        &lt;link to referenced bug: [https://github.com/kubernetes/kubernetes/issues/100277](https://github.com/kubernetes/kubernetes/issues/100277)>


        Derek Carr to Everyone (10:34 AM)


        i think the key tension here is that pdb as written is linked to pod ready state, but it sounds like there is a desire to support additional scope of pod starting but not yet terminal state.


        David Eads to Everyone (10:37 AM)


        a pod that is not ready seems already disrupted to me


        Jordan Liggitt to Everyone (10:38 AM)


        not ready pods are excluded from service selection, they do not contribute to pdb status.currentHealthy ... for PDB's definition of disruption, they seem disrupted to me as well


        Michael Gugino to Everyone (10:39 AM)


        not all pods are behind services, like storage pods


        Derek Carr to Everyone (10:39 AM)


        i am not debating what jordan/david are saying but only asking if it makes sense to offer a knob for the present behavior before tightening the behavior as implemented.


        David Eads to Everyone (10:40 AM)


        momentum matters.  If it was ready before and the kubelet hasn't checked, updating to ready=false sounds buggy


        David Eads to Everyone (10:49 AM)


        having eviction and pdb controller agree on how to count available pods is important for visilbity.

    * 
* [SergeyKanzhelev] Dockershim deprecation status [https://docs.google.com/document/d/1DiwCRJffBjoLuDguW9auMXS2NTmkp3NAS7SGJQBu8JY/edit#](https://docs.google.com/document/d/1DiwCRJffBjoLuDguW9auMXS2NTmkp3NAS7SGJQBu8JY/edit#) 
* [pohly]
    * Deprecation of kubelet command line flags in favor of config file - is that still planned?
    * Is the logging part of the config alpha? Shouldn't that be mentioned in the field comment so that it shows up in generated documentation?
    * Heads up on incoming changes/PRs:
        * WIP: [deprecating several klog command line flags](https://github.com/kubernetes/kubernetes/pull/105042)
        * WIP: [new command line flags and config options for JSON output](https://github.com/kubernetes/kubernetes/pull/104873)
        * [generic ephemeral volume checks](https://github.com/kubernetes/kubernetes/pull/100482) \



## September 21, 2021 [Cancelled]

Cancelled due to lack of agenda.


## September 14, 2021

Total active pull requests: [238](https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+is%3Aopen+)


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
   <td><strong><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++created%3A2021-09-07T17%3A00%3A00%2B0000..2021-09-14T16%3A32%3A47%2B0000">32</a></strong>
   </td>
   <td>Closed:
   </td>
   <td><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++is%3Aunmerged+closed%3A2021-09-07T17%3A00%3A00%2B0000..2021-09-14T16%3A32%3A47%2B0000">11</a>
   </td>
  </tr>
  <tr>
   <td>Updated:
   </td>
   <td><strong><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+updated%3A2021-09-07T17%3A00%3A00%2B0000..2021-09-14T16%3A32%3A47%2B0000+created%3A%3C2021-09-07T17%3A00%3A00%2B0000">79</a></strong>
   </td>
   <td>Merged:
   </td>
   <td><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++merged%3A2021-09-07T17%3A00%3A00%2B0000..2021-09-14T16%3A32%3A47%2B0000">8</a>
   </td>
  </tr>
</table>




* [ehashman] critical-urgent 1.22 static pod regression
    * Issue [https://github.com/kubernetes/kubernetes/issues/104648](https://github.com/kubernetes/kubernetes/issues/104648) 
    * WIP PR [https://github.com/kubernetes/kubernetes/pull/104847](https://github.com/kubernetes/kubernetes/pull/104847) 
    * e2e test reproducer [https://github.com/kubernetes/kubernetes/pull/104919](https://github.com/kubernetes/kubernetes/pull/104919) [https://github.com/kubernetes/kubernetes/pull/104924](https://github.com/kubernetes/kubernetes/pull/104924) 
    * Do we need any additional reviewers?
    * [https://docs.google.com/document/d/1NJKYNgoXZKGS5la4MvGrB42-DNQSFf4JQkEcZRnrDzA/edit#heading=h.d2hvrqrlw3e6](https://docs.google.com/document/d/1NJKYNgoXZKGS5la4MvGrB42-DNQSFf4JQkEcZRnrDzA/edit#heading=h.d2hvrqrlw3e6) (from  smarterclayton)
    * Are we allowing static UID for static pods?
        * Issue at hand - same manifest re-added generates the same UID
        * But static UID is also the problem
        * Does anybody has dependency on static UID?
            * Not clear?
            * Big risk to break people if we change this behavior.
    * More reviewers?
        * Lantao will take a look
* [manugupt1] Looking for reviewers / approvers for to add to KEP [https://github.com/kubernetes/enhancements/pull/2811](https://github.com/kubernetes/enhancements/pull/2811)
    * Slides from the last presentation for SIG Node (July 13, 2021): [NodeLocal Pluggable admit handler-v2.pptx](https://github.com/kubernetes/enhancements/files/6850398/NodeLocal.Pluggable.admit.handler-v2.pptx)
    * [Derek] Let’s review this for 1.24, too late for 1.23.
    * Should it be implemented as a Daemonset instead of managed plugin?
        * Manu: we will investigate options. Fargte doesn’t support Daemonset, but it may be a good option.
* [adrianreber] Checkpoint/Restore KEP
    * [https://github.com/kubernetes/enhancements/pull/1990](https://github.com/kubernetes/enhancements/pull/1990)
        * One "Generally LGTM", second review missing
    * Code changes also ready for review:
        * [https://github.com/kubernetes/kubernetes/pull/104907](https://github.com/kubernetes/kubernetes/pull/104907)
    * Mrunal: need some work to clearly separate stages of implementation
    * Adrian: actual code changes are WIP
* [Markus / Mrunal] Speeding up pod startup
    * [https://docs.google.com/presentation/d/1NdUb0A1MyrDNyyEqztmXqOeBlI0U9uZYZd4P3_-p0Fo/edit?usp=sharing](https://docs.google.com/presentation/d/1NdUb0A1MyrDNyyEqztmXqOeBlI0U9uZYZd4P3_-p0Fo/edit?usp=sharing)
    * Derek: Reach out to sig-storage on volume manager related problems

		Chat:


        Mike Brown to Everyone (10:27 AM)


        subsecond probe code draft for possible kep [https://github.com/mikebrow/kubernetes/commit/8926bd2c021a63f365d1a89ed465bd2c5abe599b](https://github.com/mikebrow/kubernetes/commit/8926bd2c021a63f365d1a89ed465bd2c5abe599b) 


        sjennings to Everyone (10:28 AM)


        [https://github.com/kubernetes/kubernetes/pull/66938](https://github.com/kubernetes/kubernetes/pull/66938) 


        Mike Brown to Everyone (10:28 AM)


        ^ to be merged with existing efforts to refactor probes


        Elana Hashman to Everyone (10:31 AM)


        interesting API change. seems a little hacky to me to only have one field that's milliseconds, as opposed to optionally allowing someone to change the units for all probe fields


        Mike Brown to Everyone (10:33 AM)


        yeah all probe fields make sense


        let me find the actual kep text


        Elana Hashman to Everyone (10:34 AM)


        like I'm thinking a string config that defaults to seconds/"" but you can optionally set to milliseconds


        Mike Brown to Everyone (10:34 AM)


        i like it elana


        Elana Hashman to Everyone (10:34 AM)


        will require a lot of refactoring of probes


        heh


        Mike Brown to Everyone (10:35 AM)


        subsecond kep text wip.. [https://hackmd.io/TgKlRqwCRLSXQ8btgKt9GA](https://hackmd.io/TgKlRqwCRLSXQ8btgKt9GA) 


        agreed Elana


        Brian Goff to Everyone (10:35 AM)


        Nice! One of those things I disliked about CRI, glad to have an event based API.


        Mike Brown to Everyone (10:38 AM)


        oh man lantau i’m hearing interest in eventing vs poling state updates :-)


        subscription models would be more efficient  for smaller scale applications


        Lantao Liu to Everyone (10:41 AM)


        Yeah, we wanted to do that in early days [https://github.com/kubernetes/kubernetes/issues/16831](https://github.com/kubernetes/kubernetes/issues/16831). Just not a priority, because relist has brought down the kubelet and container runtime resource usage to the number we targed. :p


        Mike Brown to Everyone (10:43 AM)


        ^ fair..


        Brian Goff to Everyone (10:43 AM)


        Yep.


        We wound up with an optional event based API in VK. (Mike: thx Brian will be interesting to revisit that in the k/k/pkg/kubelet tree)


        Mike Brown to Everyone (10:43 AM)


        maybe a workgroup to revisit to help k8s scale for these micro services


        Mike Brown to Everyone (10:44 AM)


        optional works.. esp. if for a specific node instead of doing both eventing and poling



* [ehashman] Great work on meeting enhancements deadline for 1.23!
    * Soft freeze for node will be in Oct, email forthcoming


## September 7, 2021

Total active pull requests:: [222](https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+is%3Aopen+)


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
   <td><strong><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++created%3A2021-08-31T17%3A00%3A00%2B0000..2021-09-07T16%3A50%3A00%2B0000">18</a></strong>
   </td>
   <td>Closed:
   </td>
   <td><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++is%3Aunmerged+closed%3A2021-08-31T17%3A00%3A00%2B0000..2021-09-07T16%3A50%3A00%2B0000">8</a>
   </td>
  </tr>
  <tr>
   <td>Updated:
   </td>
   <td><strong><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+updated%3A2021-08-31T17%3A00%3A00%2B0000..2021-09-07T16%3A50%3A00%2B0000+created%3A%3C2021-08-31T17%3A00%3A00%2B0000">67</a></strong>
   </td>
   <td>Merged:
   </td>
   <td><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++merged%3A2021-08-31T17%3A00%3A00%2B0000..2021-09-07T16%3A50%3A00%2B0000">2</a>
   </td>
  </tr>
</table>




* [madhav] Feedback on addition of new PodCondition for pods that have ephemeral containers created.
    * Issue: [https://github.com/kubernetes/kubernetes/issues/84353](https://github.com/kubernetes/kubernetes/issues/84353)
    * Relevant area in the KEP: [https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/277-ephemeral-containers#identifying-pods-with-ephemeral-containers](https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/277-ephemeral-containers#identifying-pods-with-ephemeral-containers) 
    * Feedback needed mainly on: should this new PodCondition be added to pods on which `exec` is done and those which have `ContainerNotifier` defined on them?
* [clayton/ryan] static pod lifecycle, static pod bugs, and static pod uids
    * [Static pod lifecycle regression](https://docs.google.com/document/d/1NJKYNgoXZKGS5la4MvGrB42-DNQSFf4JQkEcZRnrDzA/edit#) - summary of issues identified
    * Static pods that change UID or have fixed UID are recreated and deleted on disk are broken
        * The fixes to pod workers guaranteed that other kubelet loops finish - those loops depend only on UID, so “delete/recreate by same UID” is also broken because the volume manager might be tearing down a static pod with one set of volumes and needs to complete before we let it spin up a new pod (which means we need to decide how to handle that)
    * Admission around static pods with fixed static pod uids is potentially broken
    * Working to identify what the right fix is
        * Suspect it involves limiting what static pod fixed UID can do, as well as supporting the key use case of “only want one instance of a static pod running at the same time”, but will need feedback
    * Also, we need a KEP (and then docs) for what static pods are actually supposed to be doing and what we support and test those
        * Ryan / Clayton to pair up on static pods
        * Clayton to clarify pod safety around pods (which was what triggered the refactor that led to the regression)
        * Sergey to help with pod admission KEP
* [aditi] Just a note, Please add [Kubelet credential provider KEP](https://github.com/kubernetes/enhancements/issues/2133) to 1.23 [Tracking sheet](https://docs.google.com/spreadsheets/d/1P1J1QpayRmh2SNjs8T-wBCb6SgEOdWTRQ7MBol7yibk/edit#gid=1954476102).
* [vinaykul] In-place Pod Vertical Scaling KEP - status update
    * I just got back from vacation, I won’t be in today’s meeting.
    * Thanks Elana for getting the two KEPs tracked towards 1.23
        * Please approve/merge PR updating the milestones for these KEPs
        * [https://github.com/kubernetes/enhancements/pull/2948](https://github.com/kubernetes/enhancements/pull/2948)
    * I’m planning to address the outstanding items in the next few weeks.


## August 31, 2021

Total active pull requests:: [213](https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+is%3Aopen+)


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
   <td><strong><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++created%3A2021-08-24T17%3A00%3A00%2B0000..2021-08-31T16%3A39%3A26%2B0000">25</a></strong>
   </td>
   <td>Closed:
   </td>
   <td><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++is%3Aunmerged+closed%3A2021-08-24T17%3A00%3A00%2B0000..2021-08-31T16%3A39%3A26%2B0000">7</a>
   </td>
  </tr>
  <tr>
   <td>Updated:
   </td>
   <td><strong><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+updated%3A2021-08-24T17%3A00%3A00%2B0000..2021-08-31T16%3A39%3A26%2B0000+created%3A%3C2021-08-24T17%3A00%3A00%2B0000">53</a></strong>
   </td>
   <td>Merged:
   </td>
   <td><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++merged%3A2021-08-24T17%3A00%3A00%2B0000..2021-08-31T16%3A39%3A26%2B0000">10</a>
   </td>
  </tr>
</table>




* [klueska] Please add this [KEP](https://github.com/kubernetes/enhancements/issues/2902) to [tracking sheet](https://docs.google.com/spreadsheets/d/1P1J1QpayRmh2SNjs8T-wBCb6SgEOdWTRQ7MBol7yibk/edit#gid=1954476102) (only SIG leads can edit the sheet)
    * As referenced in the [“SIG Node - 1.23 Planning”](https://docs.google.com/document/d/1U10J0WwgWXkdYrqWGGvO8iH2HKeerQAlygnqgDgWv4E/edit#bookmark=id.piluljhxzmsc) document
* [mrunalp] Limiting number of concurrent image pulls
    * making it on runtime level is also possible. Maybe easier to do it on kubelet to minimize back and forth on CRI API level
    * [Lantao] memory-based limit or number limit? 
        * likely specific for go version and likely based on the number 
    * [Lantao] another concurrency setting - parallelism on number of layers
    * [Derek] let’s try serial puller as well to compare.
* [madhav] Feedback on addition of a new PodCondition on pods that have an EphermeralContainer created
    * [https://github.com/kubernetes/kubernetes/issues/84353](https://github.com/kubernetes/kubernetes/issues/84353)
    * [Derek] [https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/277-ephemeral-containers#identifying-pods-with-ephemeral-containers](https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/277-ephemeral-containers#identifying-pods-with-ephemeral-containers) 
* [MikeBrown post meeting] (memory, storage, number of concurrent)/(number of network devices/bandwidth)... possibly need a controller at the scheduler level instead of handling failures at the kubelet to reduce thrashing.
* [marosset] Failures when starting sandbox container keep Pod in pending state indefinitely - should pod be marked as failed?
    * [https://github.com/kubernetes/kubernetes/issues/104635](https://github.com/kubernetes/kubernetes/issues/104635) 
    * [Derek] PodPhases are designed this way and the hope is that sandbox will be created. It is hard to change this behavior/assumption
    * [Alex] should we have a condition or known errors which will indicate the terminate state?
    * [Ibrahim] another example - image that doesn’t match the platform will never start
    * [] CSI failure is another example where kubelet keep retrying even if the failure is terminal
    * [Derek] maybe look into the timeout when kubelet should stop trying to start the pod and will mark is failed?
    * There is also a device plugin scenario where pod cannot be scheduled on the node
    * [Derek] if this were only CRI flows, it would be easier
    * [rphillips] found a link after the meeting, there is a logic to recreate a sandbox if it failed, perhaps the error is not being reported in the sandbox statuses [[link](https://github.com/rphillips/kubernetes/blob/d1da6d47dd9dc8eeee9b96bffa9523e4c91f32b8/pkg/kubelet/kuberuntime/kuberuntime_manager.go#L482-L516)] \
	


## August 24, 2021



* [keerthan]  [https://github.com/kubernetes/kubernetes/pull/98934](https://github.com/kubernetes/kubernetes/pull/98934) (needs Clayton presence to discuss)
    * We are missing a state in the PodSandbox, there is only READY/NOTREADY, but kubelet doesn’t know whether a PodSandbox is broken, or fully cleaned up. Especially for the containerd implementation, once the pause container is stopped (before network teardown), the PodSandbox is considered NOTREADY.
    * However, even though we get the above issue fixed, there are still various corner cases that the pod in the apiserver may be deleted before kubelet stops the pod locally, e.g. network partition, force deletion. So the CNI needs to tolerate that anyway to clean up properly.
    * The CNI plugin may need to do some local checkpoint to cleanup properly in that case, e.g. what did checkpoint for dockershim network plugin [https://github.com/kubernetes/kubernetes/commit/51526d310385d6f3791db8bd00ed16abab46658f](https://github.com/kubernetes/kubernetes/commit/51526d310385d6f3791db8bd00ed16abab46658f) 
    * What information is needed from the pod object? Annotations and labels. In that case, actually CNI plugin may be able to get that information from the container runtime, e.g. containerd.
* [manugupt1] [https://github.com/kubernetes/kubernetes/issues/81134](https://github.com/kubernetes/kubernetes/issues/81134) (discuss comments)
* [adisky] Keystone container KEP [https://github.com/kubernetes/enhancements/pull/2869](https://github.com/kubernetes/enhancements/pull/2869)
    * Looking for reviewers and feedback
    * Mrunal and Derek will take a look 
* [fromani] promoting [https://github.com/kubernetes/enhancements/issues/2403](https://github.com/kubernetes/enhancements/issues/2403) to beta
    * still in time for 1.23? We mostly need to [land](https://github.com/kubernetes/kubernetes/pull/104123) [some](https://github.com/kubernetes/kubernetes/pull/102989) [fixes](https://github.com/kubernetes/kubernetes/pull/103289), then we should be able to move to beta.
    * fromani to add this item to the sig-node tracking document
* [dawnchen] Plan to promote [https://features.k8s.io/2892](https://features.k8s.io/2892) for beta in 1.23. 


## August 17, 2021


<table>
  <tr>
   <td>Total active pull requests:
   </td>
   <td><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+is%3Aopen+">207</a>
   </td>
   <td>(+5 from the last meeting)
   </td>
  </tr>
</table>



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
   <td><strong><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++created%3A2021-08-10T17%3A00%3A00%2B0000..2021-08-17T16%3A50%3A52%2B0000">21</a></strong>
   </td>
   <td>Closed:
   </td>
   <td><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++is%3Aunmerged+closed%3A2021-08-10T17%3A00%3A00%2B0000..2021-08-17T16%3A50%3A52%2B0000">5</a>
   </td>
  </tr>
  <tr>
   <td>Updated:
   </td>
   <td><strong><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+updated%3A2021-08-10T17%3A00%3A00%2B0000..2021-08-17T16%3A50%3A52%2B0000+created%3A%3C2021-08-10T17%3A00%3A00%2B0000">73</a></strong>
   </td>
   <td>Merged:
   </td>
   <td><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++merged%3A2021-08-10T17%3A00%3A00%2B0000..2021-08-17T16%3A50%3A52%2B0000">13</a>
   </td>
  </tr>
</table>




* Finish with KEPs for 1.23 review: [https://docs.google.com/document/d/1U10J0WwgWXkdYrqWGGvO8iH2HKeerQAlygnqgDgWv4E](https://docs.google.com/document/d/1U10J0WwgWXkdYrqWGGvO8iH2HKeerQAlygnqgDgWv4E/edit#heading=h.vc5zckz3wr33)
* [kmala]  [https://github.com/kubernetes/kubernetes/pull/98934](https://github.com/kubernetes/kubernetes/pull/98934)
    * Moved to next week. Need Clayton presence
* [bobbypage/qiutongs] Complications with cAdvisor/CRI integration (e.g. [file system metrics](https://github.com/google/cadvisor/pull/2872)) and thoughts on moving CRI API out of k/k staging
    * let’s move CRI to its own repository to avoid circular dependencies
    * long term want to move metrics to CRI. 
    * Question: when will this happen and whether moving CRI out of staging is a long term goal anyway?
    * Derek: trying to think of a mechanics of it. Important to preserve the promise that CRI is versioned with the k8s. if we decouple it from k/k, we need to be careful about it’s versioning. 
    * David: maybe we can have a process of feature freeze date for CRI and propagate this across multiple repos.
    * Mike: or have a k8s release branch in the cri repo specifically tied to each k8s release.
    * Dawn: CSI recently had dependencies and release issues. Maybe they can share some experiences.
    * Xiangqian, Xing: k/k has dependency on CSI. Every release there is an update in k/k. Sometimes development is on CSI repository. Trying to align these changes and coordinate them with k/k changes. Issues mitigated by the fact that SIg has control over the CSI repo.
    * Dawn: heard that there are some regrets to have CSI in separate rero. Process overhead is high.
    * Xing: also working on moving drivers from in-tree. Having them in the repo can cause issues. So pros and cons are in both cases.
    * Lantao: is it possible to do go module magic? (yes but.. managing n apis via submodule is also problematic)
    * Dawn: let’s start e-mail thread on this topic.
* [matthyx/adisky] Draft Proposal for Keystone containers KEP [https://github.com/kubernetes/enhancements/pull/2869](https://github.com/kubernetes/enhancements/pull/2869)
    * Derek: didn't have time to read through the KEP, questions can be addressed in another meeting later.
    * Matthias: only need a reviewer and then we do it asynchronously.
* [sjennings/decarr] Notification API KEP discussion \
[https://github.com/kubernetes/enhancements/pull/1995](https://github.com/kubernetes/enhancements/pull/1995) 


## August 10, 2021

Total active pull requests: [202](https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+is%3Aopen+)


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
   <td><strong><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++created%3A2021-08-03T17%3A00%3A00%2B0000..2021-08-10T16%3A25%3A31%2B0000">28</a></strong>
   </td>
   <td>Closed:
   </td>
   <td><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++is%3Aunmerged+closed%3A2021-08-03T17%3A00%3A00%2B0000..2021-08-10T16%3A25%3A31%2B0000">13</a>
   </td>
  </tr>
  <tr>
   <td>Updated:
   </td>
   <td><strong><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+updated%3A2021-08-03T17%3A00%3A00%2B0000..2021-08-10T16%3A25%3A31%2B0000+created%3A%3C2021-08-03T17%3A00%3A00%2B0000">88</a></strong>
   </td>
   <td>Merged:
   </td>
   <td><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++merged%3A2021-08-03T17%3A00%3A00%2B0000..2021-08-10T16%3A25%3A31%2B0000">29</a>
   </td>
  </tr>
</table>


	Rotten: #[98542](https://github.com/kubernetes/kubernetes/pull/98542) #[90727](https://github.com/kubernetes/kubernetes/pull/90727) (related bug may be addressed by other PRs) 



* 1.23 KEP planning [https://docs.google.com/document/d/1U10J0WwgWXkdYrqWGGvO8iH2HKeerQAlygnqgDgWv4E](https://docs.google.com/document/d/1U10J0WwgWXkdYrqWGGvO8iH2HKeerQAlygnqgDgWv4E/edit#heading=h.vc5zckz3wr33)
* [ehashman/mrunalp] Suggestion: alpha features/deprecrations deadline for SIG Node of **Oct. 15, 2021** (halfway through code freeze)
    * Rationale: things keep landing much too late, deprecation broke things right during code freeze in 1.22
    * [dawn] Easier to set this for beta/GA and deprecations than alpha features, alpha features are off by default so it’s okay to land them later
    * [sergey] Also, ensure everything has a WIP out by this date.
    * **Action:** Elana to send an email to SIG Node proposing the soft deadline
* [dawnchen/derek] Update on SIG Node approver process
* [manugupt1] [https://github.com/kubernetes/kubernetes/issues/81134#issuecomment-540800845](https://github.com/kubernetes/kubernetes/issues/81134#issuecomment-540800845) 
    * Discuss comments that I made
* [kmala]  [https://github.com/kubernetes/kubernetes/pull/98934](https://github.com/kubernetes/kubernetes/pull/98934)
* ~~[qiutongs] Add new container label in [pkg/kubelet/types/labels.go](https://github.com/kubernetes/kubernetes/blob/master/pkg/kubelet/types/labels.go) - io.kubernetes.container.restart~~
    * ~~It will be consumed by cadvisor. `k8s_&lt;container-name>_&lt;pod-name>_&lt;namespace>_&lt;pod-uid>_&lt;restart-count>`~~
    * ~~Discuss feasibility and the right place to make change. (e.g. [pkg/kubelet/kuberuntime/kuberuntime_container.go](https://github.com/kubernetes/kubernetes/blob/master/pkg/kubelet/kuberuntime/kuberuntime_container.go))~~


## August 3, 2021

Total active pull requests: [216](https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+is%3Aopen+)


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
   <td><strong><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++created%3A2021-07-27T17%3A00%3A00%2B0000..2021-08-03T16%3A55%3A15%2B0000">41</a></strong>
   </td>
   <td>Closed:
   </td>
   <td><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++is%3Aunmerged+closed%3A2021-07-27T17%3A00%3A00%2B0000..2021-08-03T16%3A55%3A15%2B0000">20</a>
   </td>
  </tr>
  <tr>
   <td>Updated:
   </td>
   <td><strong><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+updated%3A2021-07-27T17%3A00%3A00%2B0000..2021-08-03T16%3A55%3A15%2B0000+created%3A%3C2021-07-27T17%3A00%3A00%2B0000">54</a></strong>
   </td>
   <td>Merged:
   </td>
   <td><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++merged%3A2021-07-27T17%3A00%3A00%2B0000..2021-08-03T16%3A55%3A15%2B0000">10</a>
   </td>
  </tr>
</table>


Closed PRs - mostly WIP, test validation PRs. A few Rotten: #[84032](https://github.com/kubernetes/kubernetes/pull/84032) #[99611](https://github.com/kubernetes/kubernetes/pull/99611/). Merged PRs are mostly cherry-picks and test updates since we are in test freeze now.



* 1.22 release date: tomorrow, Aug. 4
* [ehashman] 1.22 burndown update
    * [https://github.com/kubernetes/kubernetes/issues?q=is%3Aopen+milestone%3Av1.22+label%3Asig%2Fnode+](https://github.com/kubernetes/kubernetes/issues?q=is%3Aopen+milestone%3Av1.22+label%3Asig%2Fnode+) 
* [matthyx] requesting reviewer and later approver role to help sig-node CI subgroup in:
    * test/e2e/common/OWNERShttps://github.com/kubernetes/kubernetes/issues?q=is%3Aopen+milestone%3Av1.22+label%3Asig%2Fnode+
    * test/e2e/node/OWNERS
    * test/e2e_node/OWNERS
    * [dawn] Possible areas of approver: top-level node, community repo, node tests in k/k, cadvisor (not a Kubernetes project), node problem detector
    * What about test-infra?
        * [derek] Should probably be tied with node e2e, would be happy with Elana and Sergey to both apply for node test approver, test-infra
    * Need to keep balancing new energy with existing contributors
* [adrianreber] checkpoint/restore KEP reworked and ready for review
    * [https://github.com/kubernetes/enhancements/pull/1990](https://github.com/kubernetes/enhancements/pull/1990)
        * Kep #2008
    * Update KEP to split into stages to mention that 1.23 only includes checkpoint part of it with ability to review the checkpoint later outside of K8s.
    * Adrian to send e-mail when KEP is updated.
    * runc already can do checkpoint. This KEP only adding the initiation of the checkpoint. What this KEP adds from the value add?
        * Derek:
            * this work adds context to checkpoint around secrets?
            * Problem with restore: images that the pod was restored with, and prevent these images to be used on later restart is hard.
        * Adrian:
            * Container engines cannot handle checkpoint in containers running in shared namespaces. Want to add it to CRI API, container engine will do the implementation. Main value-add is to decide to make it in CRI.
        * Dawn:
            * Since it’s only CRI API and not leak out to Pod API, it is fine to implement without the restore for now. 


## July 27, 2021

Total active pull requests: [204](https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+is%3Aopen+)


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
   <td><strong><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++created%3A2021-07-20T17%3A00%3A00%2B0000..2021-07-27T16%3A35%3A21%2B0000">24</a></strong>
   </td>
   <td>Closed:
   </td>
   <td><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++is%3Aunmerged+closed%3A2021-07-20T17%3A00%3A00%2B0000..2021-07-27T16%3A35%3A21%2B0000">6</a>
   </td>
  </tr>
  <tr>
   <td>Updated:
   </td>
   <td><strong><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+updated%3A2021-07-20T17%3A00%3A00%2B0000..2021-07-27T16%3A35%3A21%2B0000+created%3A%3C2021-07-20T17%3A00%3A00%2B0000">53</a></strong>
   </td>
   <td>Merged:
   </td>
   <td><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++merged%3A2021-07-20T17%3A00%3A00%2B0000..2021-07-27T16%3A35%3A21%2B0000">8</a>
   </td>
  </tr>
</table>




* Announcement: doc freeze is today!
* [ehashman] 1.22 burndown
    * Current release blockers:
        * [https://github.com/kubernetes/kubernetes/issues/103879](https://github.com/kubernetes/kubernetes/issues/103879)
        * [https://github.com/kubernetes/kubernetes/issues/103651](https://github.com/kubernetes/kubernetes/issues/103651) 
* [adrianreber] checkpoint/restore KEP reworked and ready for review
    * [https://github.com/kubernetes/enhancements/pull/1990](https://github.com/kubernetes/enhancements/pull/1990)
        * Kep #2008
    * Will move this to next week - working through comments from Derek/Mrunal
    * [sergey] Was the discussion from last time for how checkpoint/restore on different nodes will work resolved?
    * We want to store the checkpoint image in a registry so it can be transferred and avoid local storage between nodes
    * Lifecycle: can checkpoint a container so long as it doesn’t use external devices (e.g. GPUs, SRIOV) once the init container has finished running
    * [vinay] Checkpointing the running container’s image?
    * Including its memory footprint/pages and all, will be available after migration/reboot as it was before
* [vinaykul] In-place Pod Vertical Scaling KEP for early v1.23 - Status update
    * PR [https://github.com/kubernetes/kubernetes/pull/102884](https://github.com/kubernetes/kubernetes/pull/102884)
        * API review approved by Tim Hockin
        * Current PR squashed and rebased.
        * Starting work on unresolved issues in kubelet
            * Outstanding issues tracked [here](https://github.com/kubernetes/enhancements/issues/1287).
    * Scheduler changes are a bit more involved than I initially thought. They requested a separate PR that follows the above main PR. 
        * Can we do that? (Two PRs need to go in quick succession.)
        * No problem with this, but code will be reverted if one PR misses deadline.


## July 20, 2021

Total active pull requests: [193](https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+is%3Aopen+) (-1 from last week)


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
   <td><strong><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++created%3A2021-07-13T17%3A00%3A00%2B0000..2021-07-20T16%3A47%3A16%2B0000">13</a></strong>
   </td>
   <td>Closed:
   </td>
   <td><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++is%3Aunmerged+closed%3A2021-07-13T17%3A00%3A00%2B0000..2021-07-20T16%3A47%3A16%2B0000">6</a>
   </td>
  </tr>
  <tr>
   <td>Updated:
   </td>
   <td><strong><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+updated%3A2021-07-13T17%3A00%3A00%2B0000..2021-07-20T16%3A47%3A16%2B0000+created%3A%3C2021-07-13T17%3A00%3A00%2B0000">58</a></strong>
   </td>
   <td>Merged:
   </td>
   <td><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++merged%3A2021-07-13T17%3A00%3A00%2B0000..2021-07-20T16%3A47%3A16%2B0000">8</a>
   </td>
  </tr>
</table>




* [ehashman] Reminder: docs reviewable deadline today!
* [ehashman] 1.22 burndown:
    * [https://github.com/kubernetes/kubernetes/issues?q=is%3Aopen+milestone%3Av1.22+label%3Asig%2Fnode+](https://github.com/kubernetes/kubernetes/issues?q=is%3Aopen+milestone%3Av1.22+label%3Asig%2Fnode+) 
* KEPs retrospective for 1.22
    * SIG Node KEPs planning: [https://docs.google.com/document/d/1U10J0WwgWXkdYrqWGGvO8iH2HKeerQAlygnqgDgWv4E/edit#bookmark=id.7mxl83pa2zof](https://docs.google.com/document/d/1U10J0WwgWXkdYrqWGGvO8iH2HKeerQAlygnqgDgWv4E/edit#bookmark=id.7mxl83pa2zof) 
    * [sig node 1.22 KEPs retrospective](https://docs.google.com/spreadsheets/d/15XjKZ3ZUp8_VVWigM33VKwCmnjd5pVp28DTslweiH0I/edit#gid=0) 
        * 13 of 24 implemented
        * 2 - w/exception
        * 4 - denied with exception
    * Kubernetes 1.22 Release Information: [https://www.kubernetes.dev/resources/release/#kubernetes-122](https://www.kubernetes.dev/resources/release/#kubernetes-122)
    * [Kubernetes 1.22 Enhancements Tracking](https://docs.google.com/spreadsheets/d/1mlui0brYypOAsgS2D13fvcs3At1uMq4i1gWfneq-jxY/edit#gid=936265414)
    * _Things that could have gone better_
        * Difficulty with API reviews. Until the last day we didn’t have comments from API reviewers. Some PRs were not marked for API reviews unless the last moment.
        * **Suggested action:** in future releases, ensure all node KEP PRs are tagged with api-review and have a reviewer assigned as early as possible.
        * Reviews came in late for a lot of feature PRs, a lot of them were marked WIP until the very last week
        * **Suggested action:** KEP PRs need to be marked as ready for review well before the final week before code freeze
        * Some things we asked for exceptions that weren’t quite ready
        * **Suggested action:** In the future, only ask for exceptions when we’ve engaged with chairs early and the feature is nearly complete
        * PRR didn’t catch some of the breaking issues with seccomp by default; by enabling it by default, could break some workloads
    * _Things that went well_
        * Required e2es to be ready in PRs themselves and that worked really well: gRPC probes were an example and we found quite a few issues by asking for tests
        * Small KEPs like pod FQDN, Size memory backed volumes merged very early
    * 
* [ddebroy] Introduce Draft KEP “Runtime Assisted Mounting”
    * [https://docs.google.com/document/d/1ieVOqe1-Dc4lTMHTmp-vd0ivG_KZYvU7x30p3AYBndA/](https://docs.google.com/document/d/1ieVOqe1-Dc4lTMHTmp-vd0ivG_KZYvU7x30p3AYBndA/)
    * Overview [slides](https://docs.google.com/presentation/d/1xvrJMt5zum4wig1h9VHHyKofIQ9PDnJj--OYNkpwchQ/edit?usp=sharing)
    * Discussed overview with sig-storage
    * Gather initial thoughts from sig-node
* Lantao: what exact problems it is solving, it seems like kata already solved this with Raw Block? ([https://github.com/kata-containers/runtime/issues/1354](https://github.com/kata-containers/runtime/issues/1354)) Would be nice to know more about the pros/cons of this approach to understand whether or what we need to change in Kubernetes.
* Mrunal: does it have some intersection with how Kata wants to know more about other things like Devices information
* Dawn: want more proposal to outline more details on the problem. Why it couldnt be solved in the current design of CRI?
* [xyang] Non-graceful node shutdown: [https://github.com/kubernetes/enhancements/pull/1116](https://github.com/kubernetes/enhancements/pull/1116)
    * In the KEP, we are proposing to add a “quarantine” taint to a node that is being shutdown.  It is to prevent new pods from being scheduled to this node when it comes up before it is cleaned up.
    * Can we rely on the Node Shutdown Manager to apply this taint?
* David: it is there already. NodeReady state is already set. Reason of NodeReady condition can be checked.
* Elana: caution against the taints - they are unreliable. 
* Xing: we want the status to be changed when Node is coming back up. 
* Elana: maybe look at the previous status.


## July 13, 2021

Total active pull requests: [194](https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+is%3Aopen+) (-23 from last two weeks)


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
   <td><strong><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++created%3A2021-06-29T17%3A00%3A00%2B0000..2021-07-13T15%3A51%3A59%2B0000">42</a></strong>
   </td>
   <td>Closed:
   </td>
   <td><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++is%3Aunmerged+closed%3A2021-06-29T17%3A00%3A00%2B0000..2021-07-13T15%3A51%3A59%2B0000">23</a>
   </td>
  </tr>
  <tr>
   <td>Updated:
   </td>
   <td><strong><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+updated%3A2021-06-29T17%3A00%3A00%2B0000..2021-07-13T15%3A51%3A59%2B0000+created%3A%3C2021-06-29T17%3A00%3A00%2B0000">127</a></strong>
   </td>
   <td>Merged:
   </td>
   <td><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++merged%3A2021-06-29T17%3A00%3A00%2B0000..2021-07-13T15%3A51%3A59%2B0000">43</a>
   </td>
  </tr>
</table>




* [swsehgal] Are the requirements of being reviewer of a well defined subsystem (eg Container Manager) same as [requirements](https://github.com/kubernetes/community/blob/master/community-membership.md#reviewer) of an entire component (like kubelet)
    * Background/Motivation: Hard to find reviewers in the area of CM/CPU Manager/PodResource API (in general in pkg/kubelet/cm)
    * [Dawn] approver status today is too coarse grain. We have some subareas (not implicit), but generally known. Like resource management, monitoring, security, storage, etc. In the past we had clear owners for these areas. As time passed some people moved, some changed interests. So discussion now behind the scenes how to make these areas explicit.
* ~~[fromani] final call for cpumanager policy aliases - [we need a decision](https://kubernetes.slack.com/archives/C0BP8PW9G/p1625061895006400)~~
    * ~~fromani’s take: i think this is an issue more when the feature reaches beta rather than for alpha stage~~
    * update: settled in the PR comments
* [vinaykul] In-place Pod Vertical Scaling KEP - Status update
    * PR [https://github.com/kubernetes/kubernetes/pull/102884](https://github.com/kubernetes/kubernetes/pull/102884)
        * API review approved by Tim Hockin
        * Most items from Lantao’s review addressed and TODOs taken
            * Outstanding issues also tracked [here](https://github.com/kubernetes/enhancements/issues/1287).
    * Cut from 1.22, let’s shoot for early 1.23 check-in.
    * SIG-scheduling review ETA this week.
    * Identify any missing reviewers and loop them in.
* [Balaji] KEP Add support for pluggable pod admit handler
    * PR: [https://github.com/kubernetes/enhancements/pull/2811](https://github.com/kubernetes/enhancements/pull/2811) 
    * Slides for presentation: [https://github.com/kubernetes/enhancements/pull/2811#issuecomment-875203470](https://github.com/kubernetes/enhancements/pull/2811#issuecomment-875203470) 
    * Alternative (possible): static webhook: [https://github.com/kubernetes/enhancements/issues/1872](https://github.com/kubernetes/enhancements/issues/1872) 
    * 
    *  


## July 6, 2021

Cancelled due to US holiday


## June 29, 2021



* PR/Issue update
* Bugs scrub charts:
* 

* [ehashman] Swap update
* [ehashman] Requesting Node approver [https://github.com/kubernetes/kubernetes/pull/103122](https://github.com/kubernetes/kubernetes/pull/103122) 
    * [https://github.com/kubernetes/community/blob/master/community-membership.md#approver](https://github.com/kubernetes/community/blob/master/community-membership.md#approver)
    * Need to have more granular approach on how to carve out the path to approver
* [vinaykul] In-place Pod Vertical Scaling KEP - Updates & PR review request
    * PR [https://github.com/kubernetes/kubernetes/pull/102884](https://github.com/kubernetes/kubernetes/pull/102884)
        * API changes - see [commit](https://github.com/kubernetes/kubernetes/pull/102884/commits/e42ff603f553db37a5c22390cf7bc4e11ec2bb52)
        * CRI changes - see [commit](https://github.com/kubernetes/kubernetes/pull/102884/commits/b1d910ff8afd3c47f63dfc30301cb273bf2747c9)
        * Core implementation mostly done - see [commit](https://github.com/kubernetes/kubernetes/pull/102884/commits/e659e3423d076f9767b248998cfee4650e24c170)
            * Scheduler & RQ accounting left to finish
        * Next major task: E2E tests.
            * @wangchen615 from IBM is working on this
        * Lantao to review the PRs
        * Tim Hockin has it in his backlog
        * E2E needs to be part of the main PR
* [danielfoehrkn] Auto system/kube-reserved based on actual resource usage
    * Issue: [https://github.com/kubernetes/kubernetes/issues/103046](https://github.com/kubernetes/kubernetes/issues/103046)
    * [Alexander]: The cgroup hierarchy would have to be parsed to evaluate the actual utilization? \
Expanded question/comment: I like the overall idea about adjusting some of kubelet parameters during the lifecycle of kubelet. However, it might not be a good idea to add complexity into kubelet to be dependent on cgroups implementation (v1 vs v2, linux vs windows, etc). It might be a bit better idea to generalize the way of dynamically updating kubelet configuration and use e.g. external operator(s) that would be evaluating the state of the node (e.g. via some external metrics, like Prometheus) or overall cluster health and make a decision about adjusting reserved portions.
    * [Derek]: Has there testing been done with Enforced Allocatable feature enabled?
        * recommended setup explored
        * needs further investigation
        * has runtime and kubelet put in separate cgroups?
    * Has cgroupv2 been explored? 
        * Cgroupv2 will allow more signal from kernel
    * [Dawn]: Kernel OOM is best effort at the moment
    * Kubectl exec should also be taken into consideration
    * [Eric]: Problems with under-reserved have been highlighted. What about over reservation?
        * less resources advertised to the scheduler
    * [Derek]: An interesting area that should be explored/researched further
* [bart0sh] request for PR review & approval:

    [promote huge page storage medium size to GA](https://github.com/kubernetes/kubernetes/pull/99144)


    	PR to add conformance tests? Needs update to the hardware config


    	Needs approval

* [vinayakankugoyal] thoughts on [https://github.com/cri-o/cri-o/pull/5043](https://github.com/cri-o/cri-o/pull/5043) and [https://github.com/containerd/containerd/pull/5668](https://github.com/containerd/containerd/pull/5668) could we support ambient capabilities without making a change to k8s API and CRI API? I think we can and so I opened those PRs. I would love to hear the thoughts of the CRI maintainers though!
* 

PRs status:

Total active pull requests: [217](https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+is%3Aopen+)


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
   <td><strong><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++created%3A2021-06-22T17%3A00%3A00%2B0000..2021-06-29T17%3A08%3A45%2B0000">36</a></strong>
   </td>
   <td>Closed:
   </td>
   <td><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++is%3Aunmerged+closed%3A2021-06-22T17%3A00%3A00%2B0000..2021-06-29T17%3A08%3A45%2B0000">10</a>
   </td>
  </tr>
  <tr>
   <td>Updated:
   </td>
   <td><strong><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+updated%3A2021-06-22T17%3A00%3A00%2B0000..2021-06-29T17%3A08%3A45%2B0000+created%3A%3C2021-06-22T17%3A00%3A00%2B0000">143</a></strong>
   </td>
   <td>Merged:
   </td>
   <td><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++merged%3A2021-06-22T17%3A00%3A00%2B0000..2021-06-29T17%3A08%3A45%2B0000">23</a>
   </td>
  </tr>
</table>


Three weeks stats (since last update):


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
   <td><strong><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++created%3A2021-06-08T17%3A00%3A00%2B0000..2021-06-29T17%3A10%3A27%2B0000">92</a></strong>
   </td>
   <td>Closed:
   </td>
   <td><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++is%3Aunmerged+closed%3A2021-06-08T17%3A00%3A00%2B0000..2021-06-29T17%3A10%3A27%2B0000">31</a>
   </td>
  </tr>
  <tr>
   <td>Updated:
   </td>
   <td><strong><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+updated%3A2021-06-08T17%3A00%3A00%2B0000..2021-06-29T17%3A10%3A27%2B0000+created%3A%3C2021-06-08T17%3A00%3A00%2B0000">184</a></strong>
   </td>
   <td>Merged:
   </td>
   <td><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++merged%3A2021-06-08T17%3A00%3A00%2B0000..2021-06-29T17%3A10%3A27%2B0000">56</a>
   </td>
  </tr>
</table>



## June 22, 2021



* [ehashman] Reminder: SIG Node Bug Scrub this week! Drop in and help us scrub bugs! Starts 00:00 UTC on June 24!
* [dawnchen] Review of PRs up for KEPs [https://docs.google.com/document/d/1U10J0WwgWXkdYrqWGGvO8iH2HKeerQAlygnqgDgWv4E/edit#](https://docs.google.com/document/d/1U10J0WwgWXkdYrqWGGvO8iH2HKeerQAlygnqgDgWv4E/edit#) 
* [vinaykul] In-place Pod Vertical Scaling KEP - Updates & PR review request
    * PR [https://github.com/kubernetes/kubernetes/pull/102884](https://github.com/kubernetes/kubernetes/pull/102884)
        * Core API changes - see [commit](https://github.com/kubernetes/kubernetes/pull/102884/commits/20e51697ba19e1730ed21d17aae72e8950ae0192)
        * CRI changes - see [commit](https://github.com/kubernetes/kubernetes/pull/102884/commits/e9bfbc5c46853d421de1489fb8fa9b34425b47a7)
        * This is partial implementation, work-in-progress.
            * Node checkpointing is working
            * Basic e2e resize cases are working
            * Working on ‘Resize’ status reporting and scheduler, RQ accounting change. ETA - a couple of days.
        * Next major task: E2E tests.
            * @wangchen615 from IBM is working on this
* [ehashman] How to handle SIG Node issues in k/k marked kind/support
    * Suggestion: close with a template directing people to official support channels
    * See also [https://github.com/kubernetes/community/issues/5435](https://github.com/kubernetes/community/issues/5435) 
    * Can write up a draft for the template: [https://hackmd.io/O_gw_sXGRLC_F0cNr3Ev1Q#Support-Requests](https://hackmd.io/O_gw_sXGRLC_F0cNr3Ev1Q#Support-Requests) 
* [ehashman] Swap status
* [vinayakankugoyal] can we add KEP 2763 to the 1.23 planning doc? [https://github.com/kubernetes/enhancements/pull/2757](https://github.com/kubernetes/enhancements/pull/2757) 


## June 15, 2021 [Cancelled]

Cancelled as there were no agenda proposed for the meeting.

Important dates:



* Next SIG Node meeting - review PRs for KEPs 
* [https://docs.google.com/document/d/1U10J0WwgWXkdYrqWGGvO8iH2HKeerQAlygnqgDgWv4E/edit#](https://docs.google.com/document/d/1U10J0WwgWXkdYrqWGGvO8iH2HKeerQAlygnqgDgWv4E/edit#) 
* June 24th, 25th: bug triage: [https://hackmd.io/O_gw_sXGRLC_F0cNr3Ev1Q](https://hackmd.io/O_gw_sXGRLC_F0cNr3Ev1Q) 


## June 8, 2021

Total active pull requests: [204](https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+is%3Aopen+) (+1 since the last meeting)


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
   <td><strong><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++created%3A2021-06-01T17%3A00%3A00%2B0000..2021-06-08T16%3A54%3A28%2B0000">29</a></strong>
   </td>
   <td>Closed:
   </td>
   <td><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++is%3Aunmerged+closed%3A2021-06-01T17%3A00%3A00%2B0000..2021-06-08T16%3A54%3A28%2B0000">7</a>
   </td>
  </tr>
  <tr>
   <td>Updated:
   </td>
   <td><strong><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+updated%3A2021-06-01T17%3A00%3A00%2B0000..2021-06-08T16%3A54%3A28%2B0000+created%3A%3C2021-06-01T17%3A00%3A00%2B0000">93</a></strong>
   </td>
   <td>Merged:
   </td>
   <td><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++merged%3A2021-06-01T17%3A00%3A00%2B0000..2021-06-08T16%3A54%3A28%2B0000">25</a>
   </td>
  </tr>
</table>




* [SergeyKanzhelev] Currently we have a specific order of container startup. Also second container wouldn’t start before the prestart hook finished for the previous container. This is not a documented scenario, but there are apps taking dependency on this behavior (mostly sidecars). Do we want to promote this behavior to conformance?
    * Not a conformance
    * Just have a test - do we have unit tests already?
    * Changing of this behavior will require a KEP
    * [ehashman] Conformance tests have a very specific purpose; we should be adding them to reflect universal user expectations, not implementation details. Lots of room to add unit test coverage: kubelet is only at ~56% covered right now
* [ranchothu] A KEP about taking LLC cache into cpu allocation, for that, in some architectures like AMD rome/milan, more than one LLC exist in an individual socket, original algorithm  may cause performance decrease in this scenario.
    * KEP: [https://github.com/kubernetes/enhancements/pull/2684](https://github.com/kubernetes/enhancements/pull/2684)
    * Issue:  [https://github.com/kubernetes/enhancements/issues/2621](https://github.com/kubernetes/enhancements/issues/2621)
    * PR:  [https://github.com/kubernetes/kubernetes/pull/102307](https://github.com/kubernetes/kubernetes/pull/102307)
    * [dawn] Is this a new policy or addition to existing policy?
* Francesco has been in touch with the review author - timing is not good for them. Can we have an APAC-friendly node meeting time?
* Used to run an APAC time for node for about half a year (11PM PT) but people stopped attending after the first few. Would be open to having another meeting, but people need to attend regularly
    * [fromani] maybe just reserve the slot (once/twice a month) and have the meeting only if there are agenda items?
* [ike-ma] test-infra node image 
    * (Context: [https://github.com/kubernetes/test-infra/pull/22453](https://github.com/kubernetes/test-infra/pull/22453) )
    * Do we have any policy regarding onboarding new images?
        * [[PUBLIC][Proposal] Re-catogrize the Node E2E tests](https://docs.google.com/document/d/1BdNVUGtYO6NDx10x_fueRh_DLT-SVdlPC_SsXjYCHOE/edit?usp=sharing)
        * [dawn] Lantao was the founder for this area. Need an owner for each distro. We used to have policy for this, but not specifically for image lifecycle 
        * [lantao] Previously the coreos was in presubmit, and blocked PR submission, and was later removed.   
        * [Mrunal] prefer using existing images, will ask Harshal for more details on how the Fedora CoreOS images are used
        * [dawn] swap feature itself is less OS-distro dependent, more on kernel module
    * Do we have any policy/process regarding image update/release pipeline? eg: update kernel version, containerd version etc?
        * Ubuntu/COS/Fedora/CoreOS - focus on existing image 
    * Do we have any best practice/recommendation of test coverage for image-oriented features? Two patterns in use right now: focus on feature [hugepage](https://testgrid.k8s.io/sig-node-kubelet#kubelet-serial-gce-e2e-hugepages) ([test coverage](https://github.com/kubernetes/test-infra/blame/master/config/jobs/kubernetes/sig-node/sig-node-presubmit.yaml#L337)) vs smoke run [cgroupv2](https://testgrid.k8s.io/sig-node-kubelet#pr-node-kubelet-serial-crio-cgroupv2) ([test coverage](https://github.com/kubernetes/test-infra/blame/master/config/jobs/kubernetes/sig-node/sig-node-presubmit.yaml#L413))
        * Prefer to tag on Feature
* [ehashman] Pod lifecycle rework 
    * [https://github.com/kubernetes/kubernetes/pull/102344](https://github.com/kubernetes/kubernetes/pull/102344) 
    * [https://docs.google.com/document/d/1DvAmqp9CV8i4_zYvNdDWZW-V0FMk1NbZ8BAOvLRD3Ds/edit#](https://docs.google.com/document/d/1DvAmqp9CV8i4_zYvNdDWZW-V0FMk1NbZ8BAOvLRD3Ds/edit#) 
* [jberkus] URGENT: which runtimes does the [Kubelet Stream regression affect](https://github.com/kubernetes/kubernetes/pull/102489)?  Contributor Comms needs answers so we can notify users.	
    * Went out in the May patch releases
* [ehashman] Reminder: June cherry-pick deadline is this Friday (June 11)


## June 1, 2021

Total active pull requests: [203](https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+is%3Aopen+)


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
   <td><strong><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++created%3A2021-05-25T17%3A00%3A00%2B0000..2021-06-01T16%3A59%3A20%2B0000">28</a></strong>
   </td>
   <td>Closed:
   </td>
   <td><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++is%3Aunmerged+closed%3A2021-05-25T17%3A00%3A00%2B0000..2021-06-01T16%3A59%3A20%2B0000">13</a>
   </td>
  </tr>
  <tr>
   <td>Updated:
   </td>
   <td><strong><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+updated%3A2021-05-25T17%3A00%3A00%2B0000..2021-06-01T16%3A59%3A20%2B0000+created%3A%3C2021-05-25T17%3A00%3A00%2B0000">74</a></strong>
   </td>
   <td>Merged:
   </td>
   <td><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++merged%3A2021-05-25T17%3A00%3A00%2B0000..2021-06-01T16%3A59%3A20%2B0000">8</a>
   </td>
  </tr>
</table>




* [ehashman] Reminder: code freeze is July 8 (~5.5 weeks)
* [Sergey] New APAC time for CI+Triage session
    * tentatively 03:00 UTC on Thursdays (8pm PT on Wednesdays)
* [swsehgal] Pod Resource API
    * Issue: [Inability to account for available CPUs as guaranteed pods can belong to shared pool](https://github.com/kubernetes/kubernetes/issues/102190)
        * Proposed fixes: [here](https://github.com/kubernetes/kubernetes/pull/97415#issuecomment-847891218) but a more broader set of changes might be needed. Being discussed with Intel
        * Can this fix be treated as a bug fix or we would have to push these changes to 1.23?
        * More upcoming features in podresources API: [here](https://docs.google.com/spreadsheets/d/1bG_T1Y4j_fqHiuVrcB2HdSTBFFFYc9dg2Ynpof2BZCY/edit?ts=60ad25d3#gid=103165191) (raw data [here](https://github.com/fromanirh/k8smisc/blob/main/docs/planned%20activity%20in%20the%20podresources%20API%20area%20-%20Sheet1.csv))
    * [https://github.com/kubernetes/kubernetes/issues/102191](https://github.com/kubernetes/kubernetes/issues/102191)
        * Website Doc update: [#102191](https://github.com/kubernetes/kubernetes/issues/102191)
        * KEP update: [kubernetes/enhancements#2761](https://github.com/kubernetes/enhancements/pull/2761)
* [vinayakankugoyal] [KEP-2763](https://github.com/kubernetes/enhancements/pull/2757) - ambient capabilities support
    * We are proposing changes to CRI API
    * Could someone from sig-node volunteer to be point-of-contact and help with review and approval from the sig-node side?
        * Mike Brown - github.com/mikebrow containerd
        * Mrunal - CRI-O
    * Also reach out to someone from the PSP++ effort?
* [n4j] [Redirect container stdout / stderr to file](https://github.com/kubernetes/kubernetes/issues/94892)
    * Does this need a KEP since this is a non-trivial change and might require a change in the POD API?
    * Need consensus on the extensibility of the feature i.e. would we support redirect only to file or it can be to external collectors like fluentd
    * Is there a way to control (redirect) the output of the process via the CRI?
        * [mrunal] To an additional file? There are serious performance implications of such a change
        * [n4j] want to be able to set stdout to one file, stderr to another
        * [mrunal] think there is a bigger story for the CRI and what solutions make sense
        * [Dawn] this was discussed from the beginning of the CRI - adding redirects for logs or adding additional file targets. Did not do that for many reasons: using journald or other log mechanisms cover most cases that would enable this already. There’s concerns around complexity and performance, this is why the work hasn’t been done in the past. Needs a full design proposal, starting from a problem statement, as there are already ways to do this. At a high level, there are workarounds, so want to hear about why those aren’t enough.
        * [Lantao] For a legacy container image, you can mount a logfile from the host. Are there security concerns or other issues?
        * [ehashman] Summary of problem statement from issue… possibly we just need additional documentation because people don’t know how to do this?
        * [Lantao] Will follow-up with a workaround on the issue. [https://github.com/kubernetes/kubernetes/issues/94892#issuecomment-852321278](https://github.com/kubernetes/kubernetes/issues/94892#issuecomment-852321278) 
    * Who would be point-of-contact for design discussion and some early code feedback?
* [adisky] How to proceed further on flags to [component config migration](https://github.com/kubernetes/kubernetes/issues/86843)
    * Currently 96 flags have been marked as deprecated in favor of component config without any timeline of removal.
    * New flags being added to Kubelet that are marked as [deprecated](https://github.com/kubernetes/kubernetes/blob/master/cmd/kubelet/app/options/options.go#L383) on creation. It looks confusing from the user's perspective and also It will be difficult to track later on which flags need to be removed when.
    * Should new options be added as config only and flag addition be discouraged? 
        * [ehashman] Appears to be owned by a defunct WG (Component Standard)
        * [dims] WG is gone. Need to stop the bleeding - new CLI params should be avoided, added in KubeletConfig instead. Could add unit tests to check for this and prevent this, and then a timeline for removal of deprecated flags.
        * [aditi] All the kubelet flags are getting added as deprecated.
        * [dims] Need to update community page and/or add a KEP discussing deprecation.
        * [ehashman] Suggest we do a KEP, solicit feedback, and announce and remove all the flags in one release rather than removing them piecemeal which is causing issues for downstream consumers.
        * [dawn] We’re paused because kubelet has been leading but there are other components which aren’t migrating away from flags. We paused for a year and we’re waiting on other components. Many flags are GA so we could safely deprecate, but some need their own deprecation processes (e.g. DynamicKubeletConfig)
        * [Lubomir] +1
        * [Fabrizio] Situation across Kubernetes is confusing for users and platform tools. Not just a Kubelet issue, but it’s more pressing here because of the deprecation notices.
        * [dims] It would be ideal if things were handled via Component Config but it’s not staffed. Kubelet is a bit special because it usually doesn’t run containerized like other components.
        * [Lubomir] Story for kubelet users need to be consistent because it’s very confusing. Would like to see something from SIG Arch for k8s-wide.
        * [ehashman] Would suggest that we don’t try to increase scope, because we already have tried that with WG which is now defunct. Let’s keep with Cluster Lifecycle driving and working with Node. Might want to get a proposal from them representing user needs and what we might want to change.
        * [dawn] Let’s ensure KubeletConfig is working with kubeadm and that we’re not adding new flags. Can’t necessarily just get rid of all the deprecated flags in one go as there is more work that needs to be done.
        * [dims] Want to publish a schedule because it makes it easier to get new contributors helping out.
        * Agreed: publishing a schedule, identifying immediate pain points is a reasonable thing to spend time on.


## May 25, 2021



* [ehashman] Fixing termination and status reporting ([doc](https://docs.google.com/document/d/1DvAmqp9CV8i4_zYvNdDWZW-V0FMk1NbZ8BAOvLRD3Ds/edit#))
    * Context: [https://github.com/kubernetes/kubernetes/pull/102025](https://github.com/kubernetes/kubernetes/pull/102025#issue-645031859) 
* [yangjunmyfm192085] Exposing container start time in kubelet /metrics/resource endpoint
    * Context: [https://github.com/kubernetes/kubernetes/issues/101851](https://github.com/kubernetes/kubernetes/issues/101851)
    * [Derek] What is the node start time? When the node boots? When kubelet first marks itself NodeReady?
    * [Elana] Static pods could start before NodeReady happens
    * [Dawn] Static pod usage is discouraged, makes sense to set this to first NodeReady time
    * [Derek] Proposed start time for containers makes sense
    * Need to be clear on what the node start time metric is used for and how it’s defined
    * [Elana] Proposal says node metric is not necessarily required, and was added for symmetry - can we potentially not add it?
    * [Dawn] We should provide the most generic data as possible, e.g. node has started and is ready to join the cluster
    * [Lantao] Node boot time could be useful in the CPU context
    * [Dawn] NPE uses node boot time as well
    * **Action:** summarize discussion on bug, bring to SIG Instrumentation for confirmation
* [verb] kubelet metrics for types of containers
    * Context: [https://pr.k8s.io/99000](https://pr.k8s.io/99000)
    * running_pod/running_containers: what should we measure and where?


## May 18, 2021

Total active pull requests: [202](https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+is%3Aopen+) (+1 from the last meeting).


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
   <td><strong><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++created%3A2021-05-11T17%3A00%3A00%2B0000..2021-05-18T16%3A55%3A41%2B0000">24</a></strong>
   </td>
   <td>Closed:
   </td>
   <td><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++is%3Aunmerged+closed%3A2021-05-11T17%3A00%3A00%2B0000..2021-05-18T16%3A55%3A41%2B0000">17</a>
   </td>
  </tr>
  <tr>
   <td>Updated:
   </td>
   <td><strong><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+updated%3A2021-05-11T17%3A00%3A00%2B0000..2021-05-18T16%3A55%3A41%2B0000+created%3A%3C2021-05-11T17%3A00%3A00%2B0000">94</a></strong>
   </td>
   <td>Merged:
   </td>
   <td><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++merged%3A2021-05-11T17%3A00%3A00%2B0000..2021-05-18T16%3A55%3A41%2B0000">9</a>
   </td>
  </tr>
</table>




* [ehashman] Node bug scrub
    * Proposed date: June 24-25
* [ehashman] Swap work breakdown
    * [https://issues.redhat.com/browse/OCPNODE-470?_sscc=t](https://issues.redhat.com/browse/OCPNODE-470?_sscc=t) 
* [mrunal/harshal] node e2e - runc update broke crio tests: 
    * Enabling this: [https://prow.k8s.io/?job=*pull-kubernetes-node-crio-e2e*](https://prow.k8s.io/?job=*pull-kubernetes-node-crio-e2e*) as presubmit would have caught it.
    * [Dawn] need to understand how much combinations of systemd versions we need to test. CoreOS was excluded before from presubmits because it was on latest systemd and was very noisy. The decision was that it wasn’t the SIG Node scope to own this.
    * [Mrunal] this may not be the same as this version of crio is running on production for many customers. Making this job blocking, runc updates would be easier
    * [Dawn] is it kubelet or systemd/runc integration issue?
    * [M] it is kubelet running into “wrong“ codepath.
    * [Elana] Based on previous meetings - the goal os CI upstream is to support CRI, and testing more than 1 container runtime.
    * [D] if kubelet supports cgroup driver, likely need to run tests
    * [Brian Goff] Problem is this was exposed *after* rc94 was released. CI is needed against runc HEAD
    * [Elana] there were some WIP PRs open with runc HEAD but the lack of the blocking job meant no signal
    * [Dawn] as long as we support these tests it is fine. When CoreOS was breaking people’s PRs it was a problem
    * [David] Serial tests may have better signal
    * [Mrunal] making it a runc issue - hard. They don’t have enough CI and we might need to contribute
    * [Brian] Periodic job on runc would be better to catch regressions early, before runc releases.
    * [Dawn] In the past we started this conversation. But didn’t follow up these details completely
    * [Mrunal] I’m on of runc maintainers and can facilitate the discussion. Every time runc tags something, people asking for more changes. Frequent updates are hard on runc because of a tag requirement from kubernetes.
    * [Lantao] The model we use for containerd today is that containerd decides what runc version to use, and we test containerd + runc as a whole. Containerd does update runc periodically, and the runc update needs to go through Kubernetes node e2e tests.
    * [Mike Brown] we run tests against master of containerd. SHould be also possible for runc
    * [Mrunal] Situation is different as runc is vendored into k8s
    * [Dawn] should we look into this integration and make sure we test it?
    * [Lantao] Problem is with runc vendoring
    * [Mrunal] yes, tags are rare and k/k cannot take non-taged deps
    * [Mike Brown] are tags only for release?
    * [Mrunal] Dims and Liggit owns this decision (about tags)
    * [Mike Brown] maybe run both. 
    * [David] Question is whether test infra would allow to vendor latest for test
    * [Dawn] kubelet to runc integration is where the challenge is coming from. Maybe runc can tag daily?
    * **[Mrunal] let’s have a meeting with runc maintainers to work out the plan**
    * [Elana] talked with Dims. Opinnion: let’s get more tags. In-between vendoring is hard.
    * [Brian] lots of diff on runc. It needs to slow down. Daily tagging sounds rough.
    * [Mike] maybe containerd/cAdvisor can run against master runc continuously
    * [Dawn] since k8s is major runc user, maybe runc is the ultimately the best place to run these tests.
    * [Mrunal] about too many diffs and slow down - I wish we slow down, but there are so many at-scale issues being discovered now.
    * [Brian Goff] hard to test all permutations of runc uses is hard. Maybe have a subset of well-known cases integrated into runc.
    * [Dawn] let’s not overreact and not slow down everything. Not tax everything because of the need to test this integration
* [Brian Goff] virtual kubelet: downward API was a copy (don’t want to import k8s.io/k). Can we make downward API parsing moved to someplace else?
    * [Elana] is there issue tracking this?
    * [Dawn] file an issue, please, to discuss. One con to share: maintaining release is becoming harder with components moving to vendor. Cost increases as less ownership understanding for the components that were moved out. Identifying owners is hard for components that were moved out (like cAdvisor or NPD). 
    * Also integration and compatibility issues will arise.
    * [Brian] can sympathize with it. Will open an issue.


## May 11, 2021

Total active pull requests: [201](https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+is%3Aopen+) (+2 from the last meeting).


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
   <td><strong><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++created%3A2021-05-04T17%3A00%3A00%2B0000..2021-05-11T16%3A58%3A00%2B0000">26</a></strong>
   </td>
   <td>Closed:
   </td>
   <td><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++is%3Aunmerged+closed%3A2021-05-04T17%3A00%3A00%2B0000..2021-05-11T16%3A58%3A00%2B0000">9</a>
   </td>
  </tr>
  <tr>
   <td>Updated:
   </td>
   <td><strong><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+updated%3A2021-05-04T17%3A00%3A00%2B0000..2021-05-11T16%3A58%3A00%2B0000+created%3A%3C2021-05-04T17%3A00%3A00%2B0000">102</a></strong>
   </td>
   <td>Merged:
   </td>
   <td><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++merged%3A2021-05-04T17%3A00%3A00%2B0000..2021-05-11T16%3A58%3A00%2B0000">17</a>
   </td>
  </tr>
</table>




* KEP freeze is **Thursday, May 13**
* [vinaykul] In-place Pod Vertical Scaling KEP for 1.22
    * PRR review [https://github.com/kubernetes/enhancements/pull/2474](https://github.com/kubernetes/enhancements/pull/2474)
        * [https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/1287-in-place-update-pod-resources](https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/1287-in-place-update-pod-resources)
        * [https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/2273-kubelet-container-resources-cri-api-changes](https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/2273-kubelet-container-resources-cri-api-changes)
    * Identify code reviewer (for core feature code + tests)
* [jackfrancis] Brief overview of a proposed fix to address node ephemeral storage race during kubelet bootstrap
    * [https://github.com/kubernetes/kubernetes/pull/101882](https://github.com/kubernetes/kubernetes/pull/101882) 


## May 4, 2021

Total active pull requests: [199](https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+is%3Aopen+) (-9 from the last meeting).


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
   <td><strong><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++created%3A2021-04-27T17%3A00%3A00%2B0000..2021-05-04T16%3A52%3A57%2B0000">24</a></strong>
   </td>
   <td>Closed:
   </td>
   <td><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++is%3Aunmerged+closed%3A2021-04-27T17%3A00%3A00%2B0000..2021-05-04T16%3A52%3A57%2B0000">9</a>
   </td>
  </tr>
  <tr>
   <td>Updated:
   </td>
   <td><strong><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+updated%3A2021-04-27T17%3A00%3A00%2B0000..2021-05-04T16%3A52%3A57%2B0000+created%3A%3C2021-04-27T17%3A00%3A00%2B0000">98</a></strong>
   </td>
   <td>Merged:
   </td>
   <td><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++merged%3A2021-04-27T17%3A00%3A00%2B0000..2021-05-04T16%3A52%3A57%2B0000">24</a>
   </td>
  </tr>
</table>




* [ehashman] 1.22 KEP review/status check [https://docs.google.com/document/d/1U10J0WwgWXkdYrqWGGvO8iH2HKeerQAlygnqgDgWv4E/edit#](https://docs.google.com/document/d/1U10J0WwgWXkdYrqWGGvO8iH2HKeerQAlygnqgDgWv4E/edit#) 
    * **Enhancements freeze is May 13**
    * [https://bit.ly/k8s122-enhancements](https://bit.ly/k8s122-enhancements) 
* [ehashman] Developer Guide Audit request [https://github.com/kubernetes/community/issues/5234](https://github.com/kubernetes/community/issues/5234) 
* [vinaykul] In-place Pod Vertical Scaling KEP sponsorship for 1.22
    * KEPs for 1.22:
        * [https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/1287-in-place-update-pod-resources](https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/1287-in-place-update-pod-resources)
        * [https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/2273-kubelet-container-resources-cri-api-changes](https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/2273-kubelet-container-resources-cri-api-changes)
    * [https://github.com/kubernetes/enhancements/pull/1883](https://github.com/kubernetes/enhancements/pull/1883) merged. Thank you!!
    * ReviewPR [https://github.com/kubernetes/enhancements/pull/2474](https://github.com/kubernetes/enhancements/pull/2474)
        * It’s just a housekeeping PR that adds PRR section addition to Kubelet-CRI KEP as per updated template guidelines.
    * Do we have enough runway / reviewer resources to target In-Place Pod Update (core feature + tests) for 1.22?
* [jackfrancis] Quick discussion about kubelet pod admission during bootstrapping: static pods can be scheduled before node allocatable storage is detected
    * [https://github.com/kubernetes/kubernetes/issues/99305](https://github.com/kubernetes/kubernetes/issues/99305)
* [tallclair] PSP replacement KEP, capabilities requirements
    * [https://github.com/kubernetes/enhancements/pull/2582](https://github.com/kubernetes/enhancements/pull/2582)
    * Specifically: [https://github.com/kubernetes/enhancements/pull/2582/commits/bf7213c41a9ee0f518e5ce93f7b8dc7394de9900](https://github.com/kubernetes/enhancements/pull/2582/commits/bf7213c41a9ee0f518e5ce93f7b8dc7394de9900)
    * Additional context: [https://kubernetes.slack.com/archives/C0BP8PW9G/p1619042209240300](https://kubernetes.slack.com/archives/C0BP8PW9G/p1619042209240300)
* [fromani][if time allows] new cpumanager policy demo available
    * if no time I’ll upload a asciinema recording and I’ll link it in this document
    * we had no time, so [demo material](https://github.com/fromanirh/fromanirh/blob/main/docs/presentations/k8s-cpumanager-smtawareness/demo/README.md) , [asciinema recording](https://asciinema.org/a/412556)


## Apr 27, 2021



* Cancelled, no agenda items


## Apr 20, 2021

Total active pull requests: [208](https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+is%3Aopen+) (+9 from the last meeting)


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
   <td><strong><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++created%3A2021-04-13T17%3A00%3A00%2B0000..2021-04-20T16%3A49%3A50%2B0000">27</a></strong>
   </td>
   <td>Closed:
   </td>
   <td><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++is%3Aunmerged+closed%3A2021-04-13T17%3A00%3A00%2B0000..2021-04-20T16%3A49%3A50%2B0000">10</a>
   </td>
  </tr>
  <tr>
   <td>Updated:
   </td>
   <td><strong><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+updated%3A2021-04-13T17%3A00%3A00%2B0000..2021-04-20T16%3A49%3A50%2B0000+created%3A%3C2021-04-13T17%3A00%3A00%2B0000">85</a></strong>
   </td>
   <td>Merged:
   </td>
   <td><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++merged%3A2021-04-13T17%3A00%3A00%2B0000..2021-04-20T16%3A49%3A50%2B0000">8</a>
   </td>
  </tr>
</table>




* [fromani] (carry from April 13, we had no time to cover this item) extending cpumanager with new policies - [slides](https://github.com/fromanirh/fromanirh/blob/main/docs/presentations/k8s-cpumanager-smtawareness/smt-sware-cpumanger-sig-node-20210420.pdf) - [enh issue](https://github.com/kubernetes/enhancements/issues/2625) - (WIP) [KEP](https://github.com/kubernetes/enhancements/pull/2626)
    * Add new policies to enable precise threading allocation
        * use case
        * how the policies could look like
    * Followup as discussion on ML: [https://groups.google.com/g/kubernetes-sig-node/c/FMzdUf9UiP0](https://groups.google.com/g/kubernetes-sig-node/c/FMzdUf9UiP0)
        * (re)evaluate outsourced policies
        * cpumanager as device plugin?
* [alukiano]  (carry from April 13, we had no time to cover this item) Extending pod resource API with the memory manager metrics
    * The WIP PR, that shows changes to the API - [https://github.com/kubernetes/kubernetes/pull/101030](https://github.com/kubernetes/kubernetes/pull/101030)
    * Does it require a separate KEP? - no
    * Does it require a separate feature gate? - no
* [mrunal] 1.22 KEP prioritization

    [https://docs.google.com/document/d/1U10J0WwgWXkdYrqWGGvO8iH2HKeerQAlygnqgDgWv4E/edit#](https://docs.google.com/document/d/1U10J0WwgWXkdYrqWGGvO8iH2HKeerQAlygnqgDgWv4E/edit#)

* [ehashman] Feature gate review? [https://docs.google.com/document/d/1-NNfuFMh246_ujQpYBCldLQ9hmqEbukaa79Xe1WGwdQ/edit#](https://docs.google.com/document/d/1-NNfuFMh246_ujQpYBCldLQ9hmqEbukaa79Xe1WGwdQ/edit#) 
* [paco] about metrics kubelet_running_pods: [https://github.com/kubernetes/kubernetes/issues/99624#issuecomment-820178670](https://github.com/kubernetes/kubernetes/issues/99624#issuecomment-820178670) which proposal is preferred? (will try to attend the sig meeting)
* [AkihiroSuda] [rootless](https://rootlesscontaine.rs/) mode [https://github.com/kubernetes/kubernetes/pull/92863](https://github.com/kubernetes/kubernetes/pull/92863) 
    * Unrelated to [UserNS KEP](https://github.com/kubernetes/enhancements/issues/127), and does not conflict either
    * [The patch set is very small](https://github.com/kubernetes/kubernetes/pull/92863/files), because almost all stuff is handled by [RootlessKit](https://github.com/rootless-containers/rootlesskit), CRI, and OCI.  Changes on k/k are just for ignoring permission errors during setting `sysctl `and `rlimit `values (for kubelet and kube-proxy).
    * kind (“main” branch) already supports rootless without patching Kubernetes (by faking `/proc/sys` with bind-mount), but merging the patches is still preferred for better robustness and simplicity [https://kind.sigs.k8s.io/docs/user/rootless/](https://kind.sigs.k8s.io/docs/user/rootless/) 
    * KEP: [https://github.com/kubernetes/enhancements/pull/1371](https://github.com/kubernetes/enhancements/pull/1371) 
* [wzshiming] Fixing negative TerminationGracePeriodSeconds [https://github.com/kubernetes/kubernetes/pull/98866](https://github.com/kubernetes/kubernetes/pull/98866) 
    * How to proceed forward?
* [mewais] (Continuing from Mar 16th) \
Add Deallocate and PostStopContainer hooks to the device plugin API.
    * Possible use cases that need to be informed of device end of use, including FPGAs, drivers, etc.
    * Background:
        * [https://github.com/kubernetes/enhancements/pull/1949](https://github.com/kubernetes/enhancements/pull/1949)
        * [https://github.com/kubernetes/kubernetes/pull/91190](https://github.com/kubernetes/kubernetes/pull/91190)
        * [https://github.com/mewais/enhancements/tree/master/keps/sig-node/1948-add-deallocate-to-device-plugin-api](https://github.com/mewais/enhancements/tree/master/keps/sig-node/1948-add-deallocate-to-device-plugin-api)


## Apr 13, 2021

Total active pull requests:[199](https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+is%3Aopen+) (-2 from the last meeting)


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
   <td><strong><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++created%3A2021-04-06T17%3A00%3A00%2B0000..2021-04-13T16%3A44%3A07%2B0000">29</a></strong>
   </td>
   <td>Closed:
   </td>
   <td><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++is%3Aunmerged+closed%3A2021-04-06T17%3A00%3A00%2B0000..2021-04-13T16%3A44%3A07%2B0000">11</a>
   </td>
  </tr>
  <tr>
   <td>Updated:
   </td>
   <td><strong><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+updated%3A2021-04-06T17%3A00%3A00%2B0000..2021-04-13T16%3A44%3A07%2B0000+created%3A%3C2021-04-06T17%3A00%3A00%2B0000">90</a></strong>
   </td>
   <td>Merged:
   </td>
   <td><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++merged%3A2021-04-06T17%3A00%3A00%2B0000..2021-04-13T16%3A44%3A07%2B0000">20</a>
   </td>
  </tr>
</table>




* feedback requested on sig node annual review
    * [https://github.com/kubernetes/community/pull/5601](https://github.com/kubernetes/community/pull/5601)
* [dawnchen/mrunal] 1.22 KEP planning

    [https://docs.google.com/document/d/1U10J0WwgWXkdYrqWGGvO8iH2HKeerQAlygnqgDgWv4E/edit#](https://docs.google.com/document/d/1U10J0WwgWXkdYrqWGGvO8iH2HKeerQAlygnqgDgWv4E/edit#)

* [fromani] extending cpumanager with [new policies](https://github.com/fromanirh/fromanirh/blob/main/docs/presentations/k8s-cpumanager-smtawareness/smt-aware-cpumanager-sig-node-20210413.pdf):
    * Add new policies to enable precise threading allocation
        * use case
        * how the policies could look like
    * Followup as discussion on ML: [https://groups.google.com/g/kubernetes-sig-node/c/FMzdUf9UiP0](https://groups.google.com/g/kubernetes-sig-node/c/FMzdUf9UiP0)
        * (re)evaluate outsourced policies
        * cpumanager as device plugin?
* [alukiano] Extending pod resource API with the memory manager metrics
    * The WIP PR, that shows changes to the API - [https://github.com/kubernetes/kubernetes/pull/101030](https://github.com/kubernetes/kubernetes/pull/101030)
    * Does it require a separate KEP?
    * Does it require a separate feature gate?
* [ehashman] draft swap KEP is up, PTAL: [https://github.com/kubernetes/enhancements/pull/2602](https://github.com/kubernetes/enhancements/pull/2602) 
* [SergeyKanzhelev] community repo PRs: [https://github.com/kubernetes/community/pulls?q=is%3Apr+is%3Aopen+label%3Asig%2Fnode](https://github.com/kubernetes/community/pulls?q=is%3Apr+is%3Aopen+label%3Asig%2Fnode) 
* [Harsh] Can [this project](https://github.com/openebs/node-disk-manager ) be considered under sig-node? 
* 


## Apr 06, 2021

Total active pull requests:[201](https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+is%3Aopen+) (+19 from the last meeting)


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
   <td><strong><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++created%3A2021-03-23T17%3A00%3A00%2B0000..2021-04-06T16%3A54%3A51%2B0000">40</a></strong>
   </td>
   <td>Closed:
   </td>
   <td><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++is%3Aunmerged+closed%3A2021-03-23T17%3A00%3A00%2B0000..2021-04-06T16%3A54%3A51%2B0000">11</a>
   </td>
  </tr>
  <tr>
   <td>Updated:
   </td>
   <td><strong><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+updated%3A2021-03-23T17%3A00%3A00%2B0000..2021-04-06T16%3A54%3A51%2B0000+created%3A%3C2021-03-23T17%3A00%3A00%2B0000">125</a></strong>
   </td>
   <td>Merged:
   </td>
   <td><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++merged%3A2021-03-23T17%3A00%3A00%2B0000..2021-04-06T16%3A54%3A51%2B0000">12</a>
   </td>
  </tr>
</table>




* [vinaykul] In-place Pod Vertical Scaling (design update) - follow-up
    * Review @thockin PR [https://github.com/kubernetes/enhancements/pull/1883](https://github.com/kubernetes/enhancements/pull/1883)
    * Review/merge PRR section addition to Kubelet-CRI KEP. See[ PR](https://github.com/kubernetes/enhancements/pull/2474)
    * vinaykul is actually on vacation skiing Mt Baker in a hawaiian shirt, but he put this on agenda to remind Derek to <span style="text-decoration:underline;">please</span> review these PRs :)
* [Jim/Mrunal] Mount propagation test changes - [https://github.com/kubernetes/kubernetes/pull/100859](https://github.com/kubernetes/kubernetes/pull/100859)
* [Dims] Tech Debt
    * [Deprecate and remove DynamicKubeletConfig](https://github.com/kubernetes/kubernetes/issues/100799)
    * [Dockershim deprecation side-effect : "kubenet" will be gone](https://github.com/kubernetes/kubernetes/issues/100707) 


## Mar 30, 2021



* Meeting canceled (lots of folks are out of office)


## Mar 23rd, 2021



* [SergeyKanzhelev/ehashman] CI/Triage subgroup updates
    * 1.21 test freeze burndown: [https://github.com/kubernetes/kubernetes/issues?q=is%3Aopen+milestone%3Av1.21+label%3Asig%2Fnode](https://github.com/kubernetes/kubernetes/issues?q=is%3Aopen+milestone%3Av1.21+label%3Asig%2Fnode) 

Total active pull requests: [182](https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+is%3Aopen+) (-8 from the last meeting)


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
   <td><strong><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++created%3A2021-03-16T17%3A00%3A00%2B0000..2021-03-23T16%3A53%3A51%2B0000">26</a></strong>
   </td>
   <td>Closed:
   </td>
   <td><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++is%3Aunmerged+closed%3A2021-03-16T17%3A00%3A00%2B0000..2021-03-23T16%3A53%3A51%2B0000">6</a>
   </td>
  </tr>
  <tr>
   <td>Updated:
   </td>
   <td><strong><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+updated%3A2021-03-16T17%3A00%3A00%2B0000..2021-03-23T16%3A53%3A51%2B0000+created%3A%3C2021-03-16T17%3A00%3A00%2B0000">83</a></strong>
   </td>
   <td>Merged:
   </td>
   <td><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++merged%3A2021-03-16T17%3A00%3A00%2B0000..2021-03-23T16%3A53%3A51%2B0000">27</a>
   </td>
  </tr>
</table>




* [BenTheElder] Moving pause out of k/k into a k8s-sigs project?
    * [https://kubernetes.slack.com/archives/C0BP8PW9G/p1615917566008000](https://kubernetes.slack.com/archives/C0BP8PW9G/p1615917566008000) 
* [haircommander/bobbypage] [CRI stats KEP](https://github.com/kubernetes/enhancements/pull/2364) update
* [IkeMa] Last call for alignment on Memory Swap alpha scoping before draft KEP
    * Community doc: [Swap support in Kubernetes](https://docs.google.com/document/d/1CZtRtC8W8FwW_VWQLKP9DW_2H-hcLBcH9KBbukie67M/edit#heading=h.e5o1ougr1fna)
    * [[Public] Kubernetes Memory Swap Proof of Concept](https://docs.google.com/document/d/1H-DrJPwXZB10CcITgbIB5-gBCjixFEdZUZ97azRn5I4/edit#heading=h.5kd3drh5wmmq) 
    * TL;DR for alpha

<table>
  <tr>
   <td>
   </td>
   <td>
Before Alpha 
   </td>
   <td>After Alpha
   </td>
   <td>Comment
   </td>
  </tr>
  <tr>
   <td>kubelet behavior
   </td>
   <td>Fail to start on swap-enabled node by default
   </td>
   <td>OK to start on swap-enabled node by default
   </td>
   <td rowspan="2" >No visible performance/behavior change from workload point of view
   </td>
  </tr>
  <tr>
   <td>consuming swap
   </td>
   <td>N/A (No workload can consume swap)
   </td>
   <td>No workload can consume swap by default
   </td>
  </tr>
  <tr>
   <td>limiting swap
   </td>
   <td>N/A
   </td>
   <td>Expose a KubeConfig parameter to set limit for all container through CRI
   </td>
   <td>Experimental only
   </td>
  </tr>
</table>




* [adrianreber/mrunal] Update on checkpoint/restore KEP
    * [http://people.redhat.com/~areber/2021-03-23-sig-node.pdf](http://people.redhat.com/~areber/2021-03-23-sig-node.pdf)


## Mar 16th, 2021



* [SergeyKanzhelev] CI/Triage subgroup updates

Total active pull requests: [190](https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+is%3Aopen+) (-4 from the last meeting)


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
   <td><strong><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++created%3A2021-03-09T17%3A00%3A00%2B0000..2021-03-16T16%3A49%3A52%2B0000">55</a></strong>
   </td>
   <td>Closed:
   </td>
   <td><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++is%3Aunmerged+closed%3A2021-03-09T17%3A00%3A00%2B0000..2021-03-16T16%3A49%3A52%2B0000">22</a>
   </td>
  </tr>
  <tr>
   <td>Updated:
   </td>
   <td><strong><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+updated%3A2021-03-09T17%3A00%3A00%2B0000..2021-03-16T16%3A49%3A52%2B0000+created%3A%3C2021-03-09T17%3A00%3A00%2B0000">137</a></strong>
   </td>
   <td>Merged:
   </td>
   <td><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++merged%3A2021-03-09T17%3A00%3A00%2B0000..2021-03-16T16%3A49%3A52%2B0000">42</a>
   </td>
  </tr>
</table>


merged: 17 cherry-picks, 9 sig/instrumentation



* Announcements:
    * Doc PR deadline is today!
    * 2021 Contributor Survey, please complete! [https://www.surveymonkey.com/r/k8scommsurvey2021](https://www.surveymonkey.com/r/k8scommsurvey2021) 
* [jackfrancis] [Container exec livenessProbe timeout update](https://docs.google.com/document/u/0/d/1qZqq7_l9pXyrLiWtM0xgoxXSdYDbB6cUqXESxJ_9oCQ/edit).
    * Brief update from non-dockershim testing of ExecProbeTimeout=false, and discovery of bug in ExecProbeTimeout=false for dockershim.
        * [https://github.com/kubernetes/kubernetes/issues/100198](https://github.com/kubernetes/kubernetes/issues/100198)
        * [https://github.com/kubernetes/kubernetes/pull/100200](https://github.com/kubernetes/kubernetes/pull/100200) 
    * Background document: [https://docs.google.com/document/d/1qZqq7_l9pXyrLiWtM0xgoxXSdYDbB6cUqXESxJ_9oCQ](https://docs.google.com/document/d/1qZqq7_l9pXyrLiWtM0xgoxXSdYDbB6cUqXESxJ_9oCQ)
* [mewais] Add Deallocate and PostStopContainer hooks to the device plugin API.
    * Possible use cases that need to be informed of device end of use, including FPGAs, drivers, etc.
    * Background:
        * [https://github.com/kubernetes/enhancements/pull/1949](https://github.com/kubernetes/enhancements/pull/1949)
        * [https://github.com/kubernetes/kubernetes/pull/91190](https://github.com/kubernetes/kubernetes/pull/91190)
        * [https://github.com/mewais/enhancements/tree/master/keps/sig-node/1948-add-deallocate-to-device-plugin-api](https://github.com/mewais/enhancements/tree/master/keps/sig-node/1948-add-deallocate-to-device-plugin-api)
* [ehashman] Need approvers for structured logging: [https://github.com/orgs/kubernetes/projects/53#column-13192240](https://github.com/orgs/kubernetes/projects/53#column-13192240) 
* [vinaykul] In-place Pod Vertical Scaling (design update) - follow-up
    * Review @thockin PR [https://github.com/kubernetes/enhancements/pull/1883](https://github.com/kubernetes/enhancements/pull/1883)
    * Review/merge PRR section addition to Kubelet-CRI KEP. See[ PR](https://github.com/kubernetes/enhancements/pull/2474)
    * Any questions or concerns with the above PRs?
* [ehashman] Swap: moving towards alignment on an MVP/alpha. Please take a look at the doc, KEP to follow in the next week or two 
    * [https://docs.google.com/document/d/1CZtRtC8W8FwW_VWQLKP9DW_2H-hcLBcH9KBbukie67M/edit#heading=h.e5o1ougr1fna](https://docs.google.com/document/d/1CZtRtC8W8FwW_VWQLKP9DW_2H-hcLBcH9KBbukie67M/edit#heading=h.e5o1ougr1fna)
* [jeremyje] node-problem-detector for Windows is in active development.
    * Issue: [https://github.com/kubernetes/node-problem-detector/issues/461](https://github.com/kubernetes/node-problem-detector/issues/461)
    * Design: [https://docs.google.com/document/d/1eiK6KAp_TFR0PgBMu2WCf49fMZcg-HHnBHMc9fALquU/edit](https://docs.google.com/document/d/1eiK6KAp_TFR0PgBMu2WCf49fMZcg-HHnBHMc9fALquU/edit)
* [Jim Ramsay] Added a github issue for the mount namespace idea (has some more specifics about the systemd issue as well): [https://github.com/kubernetes/kubernetes/issues/100259](https://github.com/kubernetes/kubernetes/issues/100259) 


## Mar 9th, 2021



* [SergeyKanzhelev] CI/Triage subgroup updates

    Total active pull requests: [194](https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+is%3Aopen+) (-13 from the last meeting)


<table>
  <tr>
   <td>
<strong>Incoming</strong>
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
   <td><strong><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++created%3A2021-03-02T17%3A00%3A00%2B0000..2021-03-09T17%3A51%3A49%2B0000">69</a></strong>
   </td>
   <td>Closed:
   </td>
   <td><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++is%3Aunmerged+closed%3A2021-03-02T17%3A00%3A00%2B0000..2021-03-09T17%3A51%3A49%2B0000">20</a>
   </td>
  </tr>
  <tr>
   <td>Updated:
   </td>
   <td><strong><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+updated%3A2021-03-02T17%3A00%3A00%2B0000..2021-03-09T17%3A51%3A49%2B0000+created%3A%3C2021-03-02T17%3A00%3A00%2B0000">154</a></strong>
   </td>
   <td>Merged:
   </td>
   <td><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++merged%3A2021-03-02T17%3A00%3A00%2B0000..2021-03-09T17%3A51%3A49%2B0000">67</a>
   </td>
  </tr>
</table>




* Announcement: Code Freeze is today!
* [jackfrancis] [Container exec livenessProbe timeout discussion](https://docs.google.com/document/u/0/d/1qZqq7_l9pXyrLiWtM0xgoxXSdYDbB6cUqXESxJ_9oCQ/edit).
    * Provide a quick overview for folks who aren’t familiar with the side-effects of finally fixing timeout enforcement for container exec livenessProbes.
    * Background document: https://docs.google.com/document/d/1qZqq7_l9pXyrLiWtM0xgoxXSdYDbB6cUqXESxJ_9oCQ
    * Summarize ongoing work:
        * [https://github.com/kubernetes/kubernetes/pull/99856](https://github.com/kubernetes/kubernetes/pull/99856)
        * [https://github.com/kubernetes/kubernetes/issues/99854](https://github.com/kubernetes/kubernetes/issues/99854)
        * [https://twitter.com/ultimateboy/status/1367539874606157826](https://twitter.com/ultimateboy/status/1367539874606157826)
        * [https://github.com/kubernetes/kubernetes/pull/97057](https://github.com/kubernetes/kubernetes/pull/97057)
* [fromani] (update/heads up) [podresources API PR](https://github.com/kubernetes/kubernetes/pull/95734) needs approval! (warning: [flakes](https://github.com/kubernetes/kubernetes/issues/99947))
    * @klueska [LGTM’d!](https://github.com/kubernetes/kubernetes/pull/95734#issuecomment-792768971)
    * [requires approval because featuregate addition](https://github.com/kubernetes/kubernetes/pull/95734#issuecomment-792769319)
    * [had to rebase because featuregate conflict](https://github.com/kubernetes/kubernetes/pull/95734#issuecomment-794044700) (LGTM’d [again](https://github.com/kubernetes/kubernetes/pull/95734#issuecomment-794068847))
    * featuregate where requested during [the prod-readiness review](https://github.com/kubernetes/enhancements/pull/2404)
* [ehashman] need review/approval before code freeze for probe grace period update [https://github.com/kubernetes/kubernetes/pull/99375](https://github.com/kubernetes/kubernetes/pull/99375) 

     -     [Jim Ramsay/Mrunal] - Run container runtime/kubelet under a different mount namespace follow up \
[https://groups.google.com/g/kubernetes-sig-node/c/yNjFrBdH18Q](https://groups.google.com/g/kubernetes-sig-node/c/yNjFrBdH18Q)



* [vinaykul] In-place Pod Vertical Scaling (design update) - follow-up
    * Review @thockin PR [https://github.com/kubernetes/enhancements/pull/1883](https://github.com/kubernetes/enhancements/pull/1883)
    * Review/merge PRR section addition to Kubelet-CRI KEP. See[ PR](https://github.com/kubernetes/enhancements/pull/2474)
    * Any questions or concerns with the above PRs?
* [alukiano] Race condition between scheduler and kubelet - [https://github.com/kubernetes/kubernetes/issues/99919](https://github.com/kubernetes/kubernetes/issues/99919)
    * Why do we need the same logic under the scheduler and kubelet that validates available resources under the node. \



## Mar 2nd, 2021



* [SergeyKanzhelev] CI/Triage subgroup updates

Total active pull requests: [207](https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+is%3Aopen+) (+7 from the last meeting)


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
   <td><strong><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++created%3A2021-02-23T17%3A00%3A00%2B0000..2021-03-02T17%3A39%3A30%2B0000">44</a></strong>
   </td>
   <td>Closed:
   </td>
   <td><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++is%3Aunmerged+closed%3A2021-02-23T17%3A00%3A00%2B0000..2021-03-02T17%3A39%3A30%2B0000">15</a>
   </td>
  </tr>
  <tr>
   <td>Updated:
   </td>
   <td><strong><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+updated%3A2021-02-23T17%3A00%3A00%2B0000..2021-03-02T17%3A39%3A30%2B0000+created%3A%3C2021-02-23T17%3A00%3A00%2B0000">118</a></strong>
   </td>
   <td>Merged:
   </td>
   <td><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++merged%3A2021-02-23T17%3A00%3A00%2B0000..2021-03-02T17%3A39%3A30%2B0000">26</a>
   </td>
  </tr>
</table>




* Announcement: Code Freeze is next Tuesday, Mar. 9
* [ehashman] KEP status
    * 12 total for the release: 9 on track, 3 at risk
        * At risk: [https://github.com/kubernetes/enhancements/issues/2411](https://github.com/kubernetes/enhancements/issues/2411) [https://github.com/kubernetes/enhancements/issues/2133](https://github.com/kubernetes/enhancements/issues/2133) [https://github.com/kubernetes/enhancements/issues/2053](https://github.com/kubernetes/enhancements/issues/2053) 
    * Please ensure PRs are up, are not Waiting on Author, have reviewers assigned, and any placeholder website PRs are up
* ~~[derekwaynecarr]~~ [ehashman] SIG Annual Report
    * [https://docs.google.com/document/d/1975dclTeRq--dFX8aJkVr8Ng40lqvzN1l_G9y4JnyHY/edit?ts=603d23a8](https://docs.google.com/document/d/1975dclTeRq--dFX8aJkVr8Ng40lqvzN1l_G9y4JnyHY/edit?ts=603d23a8) 
* [Urvashi] CRIContainerLogRotation graduation criteria update
    * Quick update to the graduation criteria, opened a PR [https://github.com/kubernetes/enhancements/pull/2547](https://github.com/kubernetes/enhancements/pull/2547) 
    * Feedback on this feature gate on platforms other than OpenShift and GKE
* [vinaykul] In-place Pod Vertical Scaling (design update) - follow-up
    * Review @thockin PR [https://github.com/kubernetes/enhancements/pull/1883](https://github.com/kubernetes/enhancements/pull/1883)
    * Review/merge PRR section addition to Kubelet-CRI KEP. See[ PR](https://github.com/kubernetes/enhancements/pull/2474)
    * Any questions or concerns with the above PRs?
* [ehashman] Swap feedback reminder
    * [https://docs.google.com/document/d/1CZtRtC8W8FwW_VWQLKP9DW_2H-hcLBcH9KBbukie67M/edit#](https://docs.google.com/document/d/1CZtRtC8W8FwW_VWQLKP9DW_2H-hcLBcH9KBbukie67M/edit#) 
* [Jim Ramsay/Mrunal] Run container runtime/kubelet under a different mount namespace


## Feb 23rd, 2021

Total active pull requests: [200](https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+is%3Aopen+) (+6 from the last meeting)


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
   <td><strong><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++created%3A2021-02-16T17%3A00%3A00%2B0000..2021-02-23T17%3A34%3A51%2B0000">35</a></strong>
   </td>
   <td>Closed:
   </td>
   <td><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++is%3Aunmerged+closed%3A2021-02-16T17%3A00%3A00%2B0000..2021-02-23T17%3A34%3A51%2B0000">13</a>
   </td>
  </tr>
  <tr>
   <td>Updated:
   </td>
   <td><strong><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+updated%3A2021-02-16T17%3A00%3A00%2B0000..2021-02-23T17%3A34%3A51%2B0000+created%3A%3C2021-02-16T17%3A00%3A00%2B0000">90</a></strong>
   </td>
   <td>Merged:
   </td>
   <td><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++merged%3A2021-02-16T17%3A00%3A00%2B0000..2021-02-23T17%3A34%3A51%2B0000">19</a>
   </td>
  </tr>
</table>




* [ehashman] triage update/PR burndown
    * cherry-picks how-to: [https://github.com/kubernetes/community/blob/master/contributors/devel/sig-release/cherry-picks.md](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-release/cherry-picks.md) 
* [dawnchen] Followup: Derek uploaded all past meetings videos today! Thanks!
* [vinaykul] In-place Pod Vertical Scaling (design update) - follow-up
    * Review @thockin PR [https://github.com/kubernetes/enhancements/pull/1883](https://github.com/kubernetes/enhancements/pull/1883)
    * Review/merge PRR section addition to Kubelet-CRI KEP. See[ PR](https://github.com/kubernetes/enhancements/pull/2474)
    * Any questions or concerns with the above PRs?
* [dawnchen] [https://github.com/kubernetes/kubernetes/pull/99319](https://github.com/kubernetes/kubernetes/pull/99319)
    * node e2e on local storage management
* [fromani] [https://github.com/kubernetes/kubernetes/pull/95734](https://github.com/kubernetes/kubernetes/pull/95734) - podresources API
    * KEP approved for 1.21, changes requested during the KEP review implemented - it seems it is not on [https://github.com/orgs/kubernetes/projects/49](https://github.com/orgs/kubernetes/projects/49) but has needs-triage label
    * Still needs reviews! (and eventually approvals) PTAL
* [ed] Promote multi sizes huge pages to GA: [https://github.com/kubernetes/kubernetes/pull/99144](https://github.com/kubernetes/kubernetes/pull/99144)


## Feb 16th, 2021


<table>
  <tr>
   <td>Total active pull requests:
   </td>
   <td><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+is%3Aopen+">194</a>
   </td>
   <td>(+17 from the last meeting)
   </td>
  </tr>
</table>



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
   <td><strong><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++created%3A2021-02-09T17%3A00%3A00%2B0000..2021-02-16T18%3A08%3A02%2B0000">42</a></strong>
   </td>
   <td>Closed:
   </td>
   <td><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++is%3Aunmerged+closed%3A2021-02-09T17%3A00%3A00%2B0000..2021-02-16T18%3A08%3A02%2B0000">6</a>
   </td>
  </tr>
  <tr>
   <td>Updated:
   </td>
   <td><strong><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+updated%3A2021-02-09T17%3A00%3A00%2B0000..2021-02-16T18%3A08%3A02%2B0000+created%3A%3C2021-02-09T17%3A00%3A00%2B0000">77</a></strong>
   </td>
   <td>Merged:
   </td>
   <td><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++merged%3A2021-02-09T17%3A00%3A00%2B0000..2021-02-16T18%3A08%3A02%2B0000">19</a>
   </td>
  </tr>
</table>




* [vinaykul] In-place Pod Vertical Scaling (design update)
    * Review @thockin PR [https://github.com/kubernetes/enhancements/pull/1883](https://github.com/kubernetes/enhancements/pull/1883)
    * Key changes to [my KEP](https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/1287-in-place-update-pod-resources) with [tracking issue](https://github.com/kubernetes/enhancements/issues/1287):
        * Move ResourcesAllocated from PodSpec to PodStatus
        * ResourcesAllocated is checkpointed in kubelet
        * Add new object PodStatus.Resize to signal pod resize status
        * Proposal to add /resize subresource (details TBD)
    * Review/merge PRs that update KEPs (add PRR section)
        * PRR section added to Kubelet-CRI KEP. See[ PR](https://github.com/kubernetes/enhancements/pull/2474)
        * PRR section for InPlace Pod Update KEP on the way (after we merge above Tim’s PR)
    * Do we still have enough runway to alpha this feature in v1.21 release? (reviewer resource)
        * No :(
* [adrianreber] Checkpoint/Restore discussion
    * experimental CRI API [https://github.com/kubernetes/kubernetes/pull/97689](https://github.com/kubernetes/kubernetes/pull/97689)
* [Andrei] Fix the getCgroupSubsystemsV1() which uses only the latest record 
    * [https://github.com/kubernetes/kubernetes/pull/96594](https://github.com/kubernetes/kubernetes/pull/96594) 
* [Andrei] video recording of past meetings -- same request [Aditi]
    * **ACTION**: Derek to upload
* [bobbypage] Approvers for orphaned pod's dangling volumes PR [https://github.com/kubernetes/kubernetes/pull/95301](https://github.com/kubernetes/kubernetes/pull/95301) needed
* [haircommander/bobbypage] KEP for cAdvisor-less, CRI-full Container and Pod Stats 
    * KEP - [https://github.com/kubernetes/enhancements/pull/2364](https://github.com/kubernetes/enhancements/pull/2364)
* [ehashman] n-2 skew version tests for kubelet? ([slack](https://kubernetes.slack.com/archives/C78F00H99/p1613497306047700))
* [ehashman] reminder: feedback on swap proposal [https://docs.google.com/document/d/1CZtRtC8W8FwW_VWQLKP9DW_2H-hcLBcH9KBbukie67M/edit#](https://docs.google.com/document/d/1CZtRtC8W8FwW_VWQLKP9DW_2H-hcLBcH9KBbukie67M/edit#) 


## Feb 9th, 2021

 Total active pull requests: [178](https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+is%3Aopen+) (-1 from the last meeting)


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
   <td><strong><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++created%3A2021-02-02T17%3A00%3A00%2B0000..2021-02-09T17%3A51%3A14%2B0000">21</a></strong>
   </td>
   <td>Closed:
   </td>
   <td><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++is%3Aunmerged+closed%3A2021-02-02T17%3A00%3A00%2B0000..2021-02-09T17%3A51%3A14%2B0000">15</a>
   </td>
  </tr>
  <tr>
   <td>Updated:
   </td>
   <td><strong><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+updated%3A2021-02-02T17%3A00%3A00%2B0000..2021-02-09T17%3A51%3A14%2B0000+created%3A%3C2021-02-02T17%3A00%3A00%2B0000">90</a></strong>
   </td>
   <td>Merged:
   </td>
   <td><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++merged%3A2021-02-02T17%3A00%3A00%2B0000..2021-02-09T17%3A51%3A14%2B0000">8</a>
   </td>
  </tr>
</table>




* [ehashman] Quick 1.21 KEP update: 10 at risk, 5 confirmed
    * [https://docs.google.com/document/d/1U10J0WwgWXkdYrqWGGvO8iH2HKeerQAlygnqgDgWv4E/edit#](https://docs.google.com/document/d/1U10J0WwgWXkdYrqWGGvO8iH2HKeerQAlygnqgDgWv4E/edit#) 
* [SergeyKanzhelev] - Dockershim deprecation timeline and updates: [https://docs.google.com/document/d/1OmMe27l5dYR0KtWYrVzSeect6_Mmw5trX8z-Ydc897U/edit?usp=sharing](https://docs.google.com/document/d/1OmMe27l5dYR0KtWYrVzSeect6_Mmw5trX8z-Ydc897U/edit?usp=sharing) 
* [adrianreber] Checkpoint/Restore discussion
    * All related PRs and issues updated
        * [https://github.com/kubernetes/enhancements/pull/1990](https://github.com/kubernetes/enhancements/pull/1990)
        * [https://github.com/kubernetes/kubernetes/pull/97194](https://github.com/kubernetes/kubernetes/pull/97194) (complete PR)
        * [https://github.com/cri-o/cri-o/pull/4199](https://github.com/cri-o/cri-o/pull/4199)
        * [https://github.com/kubernetes/kubernetes/pull/97689](https://github.com/kubernetes/kubernetes/pull/97689) (API changes PR)
        * [https://github.com/opencontainers/runc/pull/2798](https://github.com/opencontainers/runc/pull/2798)
    * Demo using these PRs
        * [https://people.redhat.com/~areber/2021-01-29-kubectl-drain-checkpoint.mp4](https://people.redhat.com/~areber/2021-01-29-kubectl-drain-checkpoint.mp4)
    * Short live demo if needed
* [@bobbypage] KEP 2000: Update graceful shutdown KEP for beta in 1.21 [https://github.com/kubernetes/enhancements/pull/2472](https://github.com/kubernetes/enhancements/pull/2472) 
    * PRR review is complete, need approvals from SIG chairs
* [harche] Kubelet Node Sizing Provider KEP 
    * [https://github.com/kubernetes/enhancements/pull/2370](https://github.com/kubernetes/enhancements/pull/2370)
    * 
* [Andrei] Fix the getCgroupSubsystemsV1() which uses only the latest record [https://github.com/kubernetes/kubernetes/pull/96594](https://github.com/kubernetes/kubernetes/pull/96594) 
* [DawnChen/Lee] KEP of changing the securityContext behavior for Ephemeral Container was approved 
    * [https://github.com/kubernetes/enhancements/pull/1690](https://github.com/kubernetes/enhancements/pull/1690)
    * [https://github.com/kubernetes/enhancements/pull/2244](https://github.com/kubernetes/enhancements/pull/2244)
* [ehashman] Triage updates
    * PR documenting triage process, needs approval: [https://github.com/kubernetes/community/pull/5436](https://github.com/kubernetes/community/pull/5436) 
    * [SergeyKanzhelev] CI Subgroup + Triage meeting moving to Wednesdays


## Feb 2nd, 2021

 Total active pull requests: [179](https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+is%3Aopen+) (+4 from the last meeting)


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
   <td><strong><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++created%3A2021-01-26T17%3A00%3A00%2B0000..2021-02-02T17%3A38%3A22%2B0000">38</a></strong>
   </td>
   <td>Closed:
   </td>
   <td><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++is%3Aunmerged+closed%3A2021-01-26T17%3A00%3A00%2B0000..2021-02-02T17%3A38%3A22%2B0000">17</a>
   </td>
  </tr>
  <tr>
   <td>Updated:
   </td>
   <td><strong><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+updated%3A2021-01-26T17%3A00%3A00%2B0000..2021-02-02T17%3A38%3A22%2B0000+created%3A%3C2021-01-26T17%3A00%3A00%2B0000">97</a></strong>
   </td>
   <td>Merged:
   </td>
   <td><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++merged%3A2021-01-26T17%3A00%3A00%2B0000..2021-02-02T17%3A38%3A22%2B0000">17</a>
   </td>
  </tr>
</table>




* [ehashman] 1.21 KEP finalization - enhancements freeze is Feb. 9
    * KEPs **must**:
        * have a tracking issue (in k/enhancements/issues)
        * have issue in v1.21 milestone
        * be “implementable” in kep.yaml
        * target the correct release and stage in kep.yaml
    * Planning doc (17): [https://docs.google.com/document/d/1U10J0WwgWXkdYrqWGGvO8iH2HKeerQAlygnqgDgWv4E/edit#](https://docs.google.com/document/d/1U10J0WwgWXkdYrqWGGvO8iH2HKeerQAlygnqgDgWv4E/edit#) 
        * TODOs listed for each thing agreed upon here
    * Outstanding PRs (31): [https://github.com/kubernetes/enhancements/pulls?q=is%3Aopen+is%3Apr+label%3Asig%2Fnode](https://github.com/kubernetes/enhancements/pulls?q=is%3Aopen+is%3Apr+label%3Asig%2Fnode) 
* [harche] Dynamic node sizing for kubelet
    * KEP - [https://github.com/kubernetes/enhancements/pull/2370](https://github.com/kubernetes/enhancements/pull/2370)
* [haircommander/bobbypage] KEP for cAdvisor-less, CRI-full Container and Pod Stats 
    * KEP - [https://github.com/kubernetes/enhancements/pull/2364](https://github.com/kubernetes/enhancements/pull/2364)
* [karan] NPD: New metrics added in last release. Deprecating certain labels: [https://github.com/kubernetes/node-problem-detector/pull/522](https://github.com/kubernetes/node-problem-detector/pull/522)
    * Any users of NPD?
    * If you haven’t upgraded, please wait until the PR is merged to avoid backwards-incompatibility.
* [Sascha / Mrunal] - Seccomp by default
* [SergeyKanzhelev] - Dockershim deprecation timeline and updates: [https://docs.google.com/document/d/1OmMe27l5dYR0KtWYrVzSeect6_Mmw5trX8z-Ydc897U/edit?usp=sharing](https://docs.google.com/document/d/1OmMe27l5dYR0KtWYrVzSeect6_Mmw5trX8z-Ydc897U/edit?usp=sharing) 


## Jan 26th, 2021

Total active pull requests: [172](https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+is%3Aopen+) (+7 from the last meeting)


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
   <td><strong><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++created%3A2021-01-19T17%3A00%3A00%2B0000..2021-01-26T17%3A44%3A24%2B0000">33</a></strong>
   </td>
   <td>Closed:
   </td>
   <td><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++is%3Aunmerged+closed%3A2021-01-19T17%3A00%3A00%2B0000..2021-01-26T17%3A44%3A24%2B0000">11</a>
   </td>
  </tr>
  <tr>
   <td>fUpdated:
   </td>
   <td><strong><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+updated%3A2021-01-19T17%3A00%3A00%2B0000..2021-01-26T17%3A44%3A24%2B0000+created%3A%3C2021-01-19T17%3A00%3A00%2B0000">91</a></strong>
   </td>
   <td>Merged:
   </td>
   <td><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++merged%3A2021-01-19T17%3A00%3A00%2B0000..2021-01-26T17%3A44%3A24%2B0000">15</a>
   </td>
  </tr>
</table>




* [kmala] Discuss what needs to be done to get the fix back in. [https://github.com/kubernetes/kubernetes/pull/97980](https://github.com/kubernetes/kubernetes/pull/97980)
* [xing-yang] What is the recommended way to monitor node components
    * Volume health KEP: [https://github.com/kubernetes/enhancements/pull/2286](https://github.com/kubernetes/enhancements/pull/2286) 
    * Volume health monitoring agent deployed as a sidecar with CSI driver on every node; each agent has a pod informer; log event on pod if abnormal volume condition detected
* [Peter/Robert] Cadvisor / CRI stats performance 
    * [https://github.com/kubernetes/kubernetes/pull/98435](https://github.com/kubernetes/kubernetes/pull/98435)
    * CRI stats path duplicates work between CRI implementation and cadvisor
    * next steps:
        * investigate performance in containerd
        * investigate CRI stats implementations without cadvisor stats (are there missing fields?)
        * investigate taking out the container stats from cadvisor, so we have machine metrics, but no duplicated work
* [[Paco](https://github.com/kubernetes/kubernetes/pulls?q=is%3Apr+author%3Apacoxu+label%3Asig%2Fnode+is%3Aclosed)/Shiming] volunteers. We have fixed some small bugs recently on sig-node 
* (wzshiming fix some [small issues in node graceful shutdown](https://github.com/kubernetes/kubernetes/pulls?q=is%3Apr+author%3Awzshiming+label%3Asig%2Fnode+is%3Aclosed+) )and we want to take more responsibility in sig-node. Are there any features or tasks that we can participate in? Or any suggestions? A simple todo list  from recent researches & learning plan
    * [unsafe sysctl optimizatio](https://github.com/kubernetes/kubernetes/issues/72593#issuecomment-752034446)n(but low priority)
    * image pull related: [pull image priority](https://github.com/kubernetes/enhancements/pull/2217)(parallel-image-pulls is more preferred)
    * init-container restarted after docker prune image [#86531](https://github.com/kubernetes/kubernetes/issues/86531)
    * [sergey recommends:] still many tests are failing or flaking: [https://testgrid.k8s.io/sig-node](https://testgrid.k8s.io/sig-node) Some ToDo items from CI group: [https://github.com/orgs/kubernetes/projects/43#column-9494825](https://github.com/orgs/kubernetes/projects/43#column-9494825) 
* [saranbalaji90] [https://github.com/kubernetes/kubernetes/issues/84165](https://github.com/kubernetes/kubernetes/issues/84165) should we have separate flags for controlling pprof and flags endpoints.


## Jan 19th, 2021

Total active pull requests: [161](https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+is%3Aopen+) (-4 from the last meeting)


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
   <td><strong><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++created%3A2021-01-12T17%3A00%3A00%2B0000..2021-01-19T17%3A48%3A53%2B0000">29</a></strong>
   </td>
   <td>Closed:
   </td>
   <td><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++is%3Aunmerged+closed%3A2021-01-12T17%3A00%3A00%2B0000..2021-01-19T17%3A48%3A53%2B0000">15</a>
   </td>
  </tr>
  <tr>
   <td>Updated:
   </td>
   <td><strong><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+updated%3A2021-01-12T17%3A00%3A00%2B0000..2021-01-19T17%3A48%3A53%2B0000+created%3A%3C2021-01-12T17%3A00%3A00%2B0000">65</a></strong>
   </td>
   <td>Merged:
   </td>
   <td><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++merged%3A2021-01-12T17%3A00%3A00%2B0000..2021-01-19T17%3A48%3A53%2B0000">18</a>
   </td>
  </tr>
</table>


Only two rotten - one needed a KEP. Another may be interesting to pick up: [https://github.com/kubernetes/kubernetes/pull/81774](https://github.com/kubernetes/kubernetes/pull/81774) if anybody is interested.

Spilled over from last meeting:



* [SergeyKanzhelev/ehashman] KubeCon EU maintainer track presentation
* [ehashman] non-CI PR board, bug backlog updates
    * board: [https://github.com/orgs/kubernetes/projects/49](https://github.com/orgs/kubernetes/projects/49) 
    * [Dims/Dawn] Bug triage process - we may need another meeting? Or offline process? This meeting needs to be kept on KEPs, Design reviews, and critical bugs. Do we need another generic SIG node meeting?
    * Dims: triage meeting shouldn’t be that big. Maybe a small audience triage meeting for only interested people.
    * Dawn: Want a formal process and velocity as goals.
    * Action: To follow up at next CI meeting, second half
* [ehashman] SIG Node KEP triage, per last meeting’s action
    * spreadsheet: [https://docs.google.com/spreadsheets/d/1nN_b8ypVxvxK1fvygBs3jvRoiNj1CZPsE2S-oTcpxEA/edit#gid=593791586](https://docs.google.com/spreadsheets/d/1nN_b8ypVxvxK1fvygBs3jvRoiNj1CZPsE2S-oTcpxEA/edit#gid=593791586) 
    * down to 46 from 54!
    * might be missing some from the list of KEPs on hold missing items: [https://github.com/kubernetes/enhancements/pulls?q=is%3Aopen+is%3Apr+label%3Asig%2Fnode+label%3Ado-not-merge%2Fhold](https://github.com/kubernetes/enhancements/pulls?q=is%3Aopen+is%3Apr+label%3Asig%2Fnode+label%3Ado-not-merge%2Fhold) 
    * please make sure your KEP has a tracking item!
* [rphillips] TOCTTOU issues in kuberuntime_sandbox [https://github.com/kubernetes/kubernetes/issues/98142](https://github.com/kubernetes/kubernetes/issues/98142) 
* [???] Issue to discuss - [Allow the configuration of the interval of cr healthcheck](https://github.com/kubernetes/kubernetes/issues/96664)
    * Needs someone to present
* [adrianreber] Initial Checkpoint/Restore/Migration discussion
    * Pod Migration issue: [https://github.com/kubernetes/kubernetes/issues/3949](https://github.com/kubernetes/kubernetes/issues/3949)
    * Checkpoint/Restore KEP: [https://github.com/kubernetes/enhancements/pull/1990](https://github.com/kubernetes/enhancements/pull/1990)
    * Implementation of one possible checkpoint/restore use case: \
kubectl drain --checkpoint \
[https://github.com/kubernetes/kubernetes/pull/97194](https://github.com/kubernetes/kubernetes/pull/97194)
    * Proposed CRI API change \
[https://github.com/kubernet		es/kubernetes/pull/97689](https://github.com/kubernetes/kubernetes/pull/97689)
    * CRI API changes implemented in CRI-O \
[https://github.com/cri-o/cri-o/pull/4199](https://github.com/cri-o/cri-o/pull/4199) \
(and crictl [https://github.com/kubernetes-sigs/cri-tools/pull/662](https://github.com/kubernetes-sigs/cri-tools/pull/662) )
    * Happy to answer any questions around checkpoint/restore
* [MaRosset] Windows KEPs / issues
    * Looking for advice on
        * [Add Windows container device support to CRI-API #97739](https://github.com/kubernetes/kubernetes/issues/97739)
            * > @RenaudWasTaken, @kad, Mike Brown, Mrunal
            * COD WG: [https://docs.google.com/document/d/1gUgAMEThkRt4RJ7pA7ZbPPmIOX2Vb7fwH025MjfcTYU/edit](https://docs.google.com/document/d/1gUgAMEThkRt4RJ7pA7ZbPPmIOX2Vb7fwH025MjfcTYU/edit)
            * 
    * Looking for reviewers 
        * [KEP 1981: Windows privileged container KEP updates for alpha #2288](https://github.com/kubernetes/enhancements/pull/2288)
        * [KEP 2258: Use kubectl to view system service logs #2271](https://github.com/kubernetes/enhancements/pull/2271)
            * Implemented in OpenShift, want to upstream changes (impacts linux nodes with journald as well)
* [@mythi, @kad]: [non-root user containers and devices access](https://docs.google.com/document/d/1SX4o71AIIrJAzbGJEIfhT2NxQpPvLlrHiZj5uUawWxk) - PSA: Update, POC patches available for review ([k/k#92211](https://github.com/kubernetes/kubernetes/issues/92211))
    * Please review


## Jan 12th, 2021

Total active pull requests: [164](https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+is%3Aopen+) (-12 from the last meeting) Way to go!!!


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
   <td><strong><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++created%3A2021-01-05T17%3A00%3A00%2B0000..2021-01-12T17%3A06%3A23%2B0000">23</a></strong>
   </td>
   <td>Closed:
   </td>
   <td><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++is%3Aunmerged+closed%3A2021-01-05T17%3A00%3A00%2B0000..2021-01-12T17%3A06%3A23%2B0000">19</a>
   </td>
  </tr>
  <tr>
   <td>Updated:
   </td>
   <td><strong><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+updated%3A2021-01-05T17%3A00%3A00%2B0000..2021-01-12T17%3A06%3A23%2B0000+created%3A%3C2021-01-05T17%3A00%3A00%2B0000">94</a></strong>
   </td>
   <td>Merged:
   </td>
   <td><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++merged%3A2021-01-05T17%3A00%3A00%2B0000..2021-01-12T17%3A06%3A23%2B0000">16</a>
   </td>
  </tr>
</table>



Please approve lgtm’d PRs from SIG Node CI group: [https://github.com/orgs/kubernetes/projects/43#column-9494828](https://github.com/orgs/kubernetes/projects/43#column-9494828) 



* [derekwaynecarr, ehashman] graceful termination period and liveness probe failures
    * see: [https://github.com/kubernetes/kubernetes/issues/64715](https://github.com/kubernetes/kubernetes/issues/64715)
    * provisional KEP for fix: [https://github.com/kubernetes/enhancements/pull/2241](https://github.com/kubernetes/enhancements/pull/2241) 
* [mrunal] 1.21 planning
    * [https://docs.google.com/document/d/1U10J0WwgWXkdYrqWGGvO8iH2HKeerQAlygnqgDgWv4E/edit#](https://docs.google.com/document/d/1U10J0WwgWXkdYrqWGGvO8iH2HKeerQAlygnqgDgWv4E/edit#)
* [rphillips, harche] auto system reserved sizing
    * [https://docs.google.com/document/d/1DOKK5div3MuUV5A8WxZauBA7XCOK8QVDoH6qZGK964k/edit?ts=5ffdc238#heading=h.kg7a8pe4fi0u](https://docs.google.com/document/d/1DOKK5div3MuUV5A8WxZauBA7XCOK8QVDoH6qZGK964k/edit?ts=5ffdc238#heading=h.kg7a8pe4fi0u)
* [SergeyKanzhelev (if I can attend)] [https://github.com/kubernetes/kubernetes/issues/97288](https://github.com/kubernetes/kubernetes/issues/97288) I suggest to revert the logic in 1.20 - possible fixes might be risky for 1.20. Reverting [https://github.com/kubernetes/kubernetes/pull/92817](https://github.com/kubernetes/kubernetes/pull/92817) will revert [https://github.com/kubernetes/kubernetes/issues/88543](https://github.com/kubernetes/kubernetes/issues/88543).
    * Revert for now, sync back up next week to investigate a root cause


## Jan 5th, 2021

Total active pull requests: [183](https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+is%3Aopen+) (-13 from the last meeting)


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
   <td><strong><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++created%3A2020-12-08T17%3A00%3A00%2B0000..2021-01-05T17%3A42%3A55%2B0000">66</a></strong>
   </td>
   <td>Closed:
   </td>
   <td><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++is%3Aunmerged+closed%3A2020-12-08T17%3A00%3A00%2B0000..2021-01-05T17%3A42%3A55%2B0000">34</a>
   </td>
  </tr>
  <tr>
   <td>Updated:
   </td>
   <td><strong><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+updated%3A2020-12-08T17%3A00%3A00%2B0000..2021-01-05T17%3A42%3A55%2B0000+created%3A%3C2020-12-08T17%3A00%3A00%2B0000">128</a></strong>
   </td>
   <td>Merged:
   </td>
   <td><a href="https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++merged%3A2020-12-08T17%3A00%3A00%2B0000..2021-01-05T17%3A42%3A55%2B0000">45</a>
   </td>
  </tr>
</table>




* [SergeyKanzhelev] we are turning things around wrt [open PRs](https://k8s.devstats.cncf.io/d/71/prs-labels-by-sig?orgId=1&var-sig_name=node&var-label_name=All%20labels%20combined&from=1548540596095&to=1611785396095) 🎉:




* [SergeyKanzhelev] From CI group, lgtm’d, needs approval: [https://github.com/orgs/kubernetes/projects/43#column-9494828](https://github.com/orgs/kubernetes/projects/43#column-9494828) 
* [SergeyKanzhelev] Please review and contribute: migrating from Dockershim documentation: [https://github.com/kubernetes/website/pull/25787](https://github.com/kubernetes/website/pull/25787) 
* [Karan] [Swap support](https://github.com/kubernetes/kubernetes/issues/53533) in kubelet (need to write a use-cases and scope KEP in the next few weeks)? Volunteers to help spec it?
    * Some initial thoughts and investigation by me: [doc](https://docs.google.com/document/d/1qFH-RA7GvaEidOnp7Y-QwAZ8g5Zcyk28VkEJFWcEX-U/edit#)
    * ehashman put together a doc full of use cases (not yet publicly shared)
    * Timeline: 1.21 for use case and scope and high-level design possibly. 1.22 alpha is a good goal
    * Doc: [https://docs.google.com/document/d/1CZtRtC8W8FwW_VWQLKP9DW_2H-hcLBcH9KBbukie67M/edit#](https://docs.google.com/document/d/1CZtRtC8W8FwW_VWQLKP9DW_2H-hcLBcH9KBbukie67M/edit#) 
* [fromani] Would like to move forward the initial [podresources API extension PR](https://github.com/kubernetes/kubernetes/pull/95734) as we have more changes coming to complete the planned extensions. Any more comments for this PR? another round of reviews?
* [ehashman] SIG Node KEP triage? 54 [currently open](https://github.com/kubernetes/enhancements/issues?q=is%3Aissue+is%3Aopen+label%3Asig%2Fnode), 37 not tracked
    * from the last release: [https://docs.google.com/document/d/1-NNfuFMh246_ujQpYBCldLQ9hmqEbukaa79Xe1WGwdQ/edit#](https://docs.google.com/document/d/1-NNfuFMh246_ujQpYBCldLQ9hmqEbukaa79Xe1WGwdQ/edit#) 
    * has been a lot of duplication of effort between SIG Node/enhancements team
    * haven’t had the bandwidth to close out stale issues because that still involves a lot of effort
    * ACTION: ehashman to sync with enhancements team, put together a spreadsheet with KEP owners to start working on driving down the KEP backlog
* [pacoxu] (not join the meeting because of zoom issues, hope to discuss next time)
    * [1.20 regression: pods failing to terminate #97288](https://github.com/kubernetes/kubernetes/issues/97288)<span style="text-decoration:underline;"> </span>(a regression, currently no one is working on)
    * [volume stats disabled when negative.(Currently, 0 means default 1m)](https://github.com/kubernetes/kubernetes/pull/96675)<span style="text-decoration:underline;"> </span>@paco a small pr but bug.
    * [Is there a plan to improve the experience of unsafe sysctl?](https://github.com/kubernetes/kubernetes/issues/72593#issuecomment-752034446) an initial proposal.
