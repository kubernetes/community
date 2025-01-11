# Kubernetes SIG-Node CI subgroup notes

## 2023/12/06

Recording: [https://youtu.be/TFp7tv72854](https://youtu.be/TFp7tv72854)   
Hosts:

- Tests: mmiranda96  
- Bugs:

Thaw will start Dec 13th.  
Next meeting on Jan 3rd.

## 2023/11/29

Recording: [https://youtu.be/utMfJzEBcvQ](https://youtu.be/utMfJzEBcvQ) 

- Hosts:  
  - Tests: Sergey  
  - Bugs: ndixita  
- New failure possible regression: E2eNode Suite.\[It\] \[sig-node\] \[Serial\] Containers Lifecycle should restart the containers in right order after the node rebootChanges  
- \[ruiwen\] [https://github.com/kubernetes/kubernetes/pull/122095](https://github.com/kubernetes/kubernetes/pull/122095)   
- \[harshal\] [https://github.com/kubernetes/kubernetes/issues/121349\#issuecomment-1813029991](https://github.com/kubernetes/kubernetes/issues/121349#issuecomment-1813029991) closer to the root cause  
- Possibly need a bug for this: [https://testgrid.k8s.io/sig-node-presubmits\#pr-crio-cgroupv2-imagefs-e2e-diskpressure](https://testgrid.k8s.io/sig-node-presubmits#pr-crio-cgroupv2-imagefs-e2e-diskpressure) 

## 2023/11/15

Recording [https://youtu.be/aoT5PnEBdus](https://youtu.be/aoT5PnEBdus) 

- Hosts:  
  - Tests: mkmir  
  - Bugs: Sergey  
- Mostly triage  
- [https://kubernetes.slack.com/archives/CN0K3TE2C/p1700060944364139](https://kubernetes.slack.com/archives/CN0K3TE2C/p1700060944364139)   
  - [https://github.com/kubernetes/kubernetes/issues/121349](https://github.com/kubernetes/kubernetes/issues/121349) 

## 2023/11/01

Recording: [https://youtu.be/PclxEBV1awI](https://youtu.be/PclxEBV1awI)   
Hosts:

- Tests:  
- Bugs: ndixita

Agenda:

* ~~Kubelet/CRIO/Containerd logs are missing from jobs that have been migrated to community cluster~~  
  * [~~https://github.com/kubernetes/kubernetes/issues/121444~~](https://github.com/kubernetes/kubernetes/issues/121444)  
* [https://github.com/kubernetes/kubernetes/pull/119496\#issuecomment-1653172666](https://github.com/kubernetes/kubernetes/pull/119496#issuecomment-1653172666)  
* Fedora swap job  
* [https://github.com/kubernetes/kubernetes/pull/121671](https://github.com/kubernetes/kubernetes/pull/121671)  
* Job config looks broken:  
  * [https://testgrid.k8s.io/sig-node-containerd\#cos-cgroupv2-containerd-e2e\&width=5](https://testgrid.k8s.io/sig-node-containerd#cos-cgroupv2-containerd-e2e&width=5) (check with mimiranda96)  
  * [https://github.com/kubernetes/kubernetes/issues/121309](https://github.com/kubernetes/kubernetes/issues/121309)  
  * [https://testgrid.k8s.io/sig-node-containerd\#e2e-cos-device-plugin-gpu\&width=5](https://testgrid.k8s.io/sig-node-containerd#e2e-cos-device-plugin-gpu&width=5)  
  *   
* Busted job config:  
  * [https://testgrid.k8s.io/sig-node-node-problem-detector\#ci-npd-e2e-node\&width=5](https://testgrid.k8s.io/sig-node-node-problem-detector#ci-npd-e2e-node&width=5) 

## 2023/10/25 \[Cancelled\]

Canceled due to an host availability. Review PRs for freeze next week. 

## 2023/10/18

Recording:

Hosts:

- Tests: ndixita  
- Bugs: kannon92

Agenda:

* Issue link for  [https://testgrid.k8s.io/sig-node-kubelet\#kubelet-gce-e2e-swap-fedora-serial](https://testgrid.k8s.io/sig-node-kubelet#kubelet-gce-e2e-swap-fedora-serial) failures  
* [https://github.com/kubernetes/kubernetes/issues/121220](https://github.com/kubernetes/kubernetes/issues/121220) for existing issue related to feodra  
  * \[Tzneal\] For the OOM Swap issue, there is a discusion on \#sig-node-swap that I think is related: [https://kubernetes.slack.com/archives/C02UCH9N02J/p1695207853069709](https://kubernetes.slack.com/archives/C02UCH9N02J/p1695207853069709)   
* Create an issue for PriorityPidEvictionOrdering  
*   
* Issue link for containerd-cgroupv1, GLIBC\_2.34 not foundGLIB  
  * [https://github.com/kubernetes/kubernetes/https://github.com/kubernetes/kubernetes/issues/121309](https://github.com/kubernetes/kubernetes/https://github.com/kubernetes/kubernetes/issues/121309)  
  * issues/121309C\_2.3 \-   
- Remind mike about COS and dbus  
  - \[Sergey\] found the right people to look at it  
- Reach out to sig testing: [https://github.com/kubernetes/kubernetes/issues/121220](https://github.com/kubernetes/kubernetes/issues/121220)

Issues Created for bugs:

\- [https://github.com/kubernetes/kubernetes/issues/121309](https://github.com/kubernetes/kubernetes/issues/121309)  
\- [https://github.com/kubernetes/kubernetes/issues/121220](https://github.com/kubernetes/kubernetes/issues/121220)  
\- [https://github.com/kubernetes/kubernetes/issues/121124](https://github.com/kubernetes/kubernetes/issues/121124)   
\- [https://github.com/kubernetes/node-problem-detector/issues/831](https://github.com/kubernetes/node-problem-detector/issues/831)  
\- 

PR to fix jobs due to configuration:  
\- [https://github.com/kubernetes/test-infra/pull/31054](https://github.com/kubernetes/test-infra/pull/31054)  
\- https://github.com/kubernetes/test-infra/pull/31051

## 2023/10/11

Recording: [https://www.youtube.com/watch?v=MZNSFlJGnMw](https://www.youtube.com/watch?v=MZNSFlJGnMw)   
Hosts:

- Tests: mmiranda96  
- Bugs: ndixita

Agenda:

- \[mmiranda96\]: file an issue for swap Fedora job  
  - [https://testgrid.k8s.io/sig-node-containerd\#node-kubelet-containerd-hugepages](https://testgrid.k8s.io/sig-node-containerd#node-kubelet-containerd-hugepages)  
  - [https://testgrid.k8s.io/sig-node-cri-o\#ci-crio-cgroupv1-node-e2e-hugepages](https://testgrid.k8s.io/sig-node-cri-o#ci-crio-cgroupv1-node-e2e-hugepages)  
  - [https://testgrid.k8s.io/sig-node-node-problem-detector\#ci-npd-e2e-kubernetes-gce-gci](https://testgrid.k8s.io/sig-node-node-problem-detector#ci-npd-e2e-kubernetes-gce-gci)  
  - [https://testgrid.k8s.io/sig-node-node-problem-detector\#ci-npd-e2e-kubernetes-gce-gci-custom-flags](https://testgrid.k8s.io/sig-node-node-problem-detector#ci-npd-e2e-kubernetes-gce-gci-custom-flags)  
- Kubelet behavior change in 1.27 related to multiple containers in Pod

## 2023/10/04

Recording: [https://www.youtube.com/watch?v=yx7iXyohVDU](https://www.youtube.com/watch?v=yx7iXyohVDU) 

Hosts:

- Tests: mmiranda96  
- Bugs: skanzhelev

Agenda:

- \[Kevin Hannon\] Presubmit Jobs and crio-community-cluster migration  
- [Reducing flakes in the periodics tests will become easier if it is possible to kick off tests as presubmits](https://docs.google.com/document/d/1QrLBy-v6B3sits0Tu5uqQFHkofejorcqw73KLxQVWsg/edit?usp=sharing)  
- mmiranda96: Create issue for fixing sig-node-containerd containerd-branch jobs (change branch name).  
- 

## 2023/09/27

Recording: [https://www.youtube.com/watch?v=ehD9x7mYcIQ](https://www.youtube.com/watch?v=ehD9x7mYcIQ)   
Hosts:

- Tests: ndixita  
- Bugs: mmiranda96

	  
Agenda:

- **\[Kevin Hannon**  [9:36 AM](https://kubernetes.slack.com/archives/C05KQLJEWHX/p1695659778256519)\]  
  It looks like there was a large refactor of the node jobs to drop bootstrap.py.  I'd suggest for our weekly meeting to look over the sig-node jobs and make sure they are all running.  
  [https://github.com/kubernetes/kubernetes/pull/120831\#issuecomment-1732281761](https://github.com/kubernetes/kubernetes/pull/120831#issuecomment-1732281761)  
  Is a list of some jobs that are failing  
    
    
- Topology manager  
  - Swati to look into why tests are being skipped: TODO  
    - Issue: [https://github.com/kubernetes/kubernetes/issues/120725](https://github.com/kubernetes/kubernetes/issues/120725)   
    - Previous Issue where we were tracking Resource Manager testing on multi-numa: [https://github.com/kubernetes/kubernetes/issues/119601](https://github.com/kubernetes/kubernetes/issues/119601) . We had fixed tests to run on multi-NUMA nodes ([https://github.com/kubernetes/test-infra/pull/30545](https://github.com/kubernetes/test-infra/pull/30545) and [https://github.com/kubernetes/test-infra/pull/30629](https://github.com/kubernetes/test-infra/pull/30629))  
- Swap:  
  - [https://github.com/kubernetes/kubernetes/pull/120139](https://github.com/kubernetes/kubernetes/pull/120139)   
  - Create an issue for [https://testgrid.k8s.io/sig-node-kubelet\#kubelet-gce-e2e-swap-fedora-serial](https://testgrid.k8s.io/sig-node-kubelet#kubelet-gce-e2e-swap-fedora-serial)  
  - https://testgrid.k8s.io/sig-node-containerd\#cos-cgroupv2-containerd-node-e2e  
- GracefulNodeShutdownOnPodPriority failing tests: [https://github.com/kubernetes/kubernetes/issues/120726](https://github.com/kubernetes/kubernetes/issues/120726)  
- [https://testgrid.k8s.io/sig-node-cri-o\#node-kubelet-serial-crio](https://testgrid.k8s.io/sig-node-cri-o#node-kubelet-serial-crio)  
- kubelet-serial-containerd: [https://github.com/kubernetes/kubernetes/issues/120913](https://github.com/kubernetes/kubernetes/issues/120913)  
- 

## 2023/09/20

Recording: [https://www.youtube.com/watch?v=pNCzeo3J7oQ](https://www.youtube.com/watch?v=pNCzeo3J7oQ) 

Hosts:

- Tests: mmiranda96  
- Bugs: ndixita

Agenda:

* \[kannon92\] Investigating eviction test cases  
  * Notice there is no test coverage for eviction with a separate image fs  
  * Can we consider adding a e2e test for crio/containerd eviction with a separate image fs?  
    * Could use some help on how to do this  
* \[kannon92\]  
  * Bootstrap py issue caused problems with prow jobs (not just the periodics but also the blocking PRs)  
  * [https://github.com/kubernetes/test-infra/issues/30759](https://github.com/kubernetes/test-infra/issues/30759)  
  * Should consider migrating the required presubmits to use decorate  
    * \[fromani\] totally\! [https://github.com/kubernetes/kubernetes/issues/120609](https://github.com/kubernetes/kubernetes/issues/120609) (perhaps should have been filed against test-infra?)

## 2023/09/13

Recording: [https://www.youtube.com/watch?v=kk2GCdZRJBs](https://www.youtube.com/watch?v=kk2GCdZRJBs)  
Hosts:

- Tests: ndixita  
- Bugs: mmiranda96

Agenda:

* \[ndixita\] Triage OOM Kill test in [https://testgrid.k8s.io/sig-node-kubelet\#kubelet-gce-e2e-swap-ubuntu-serial](https://testgrid.k8s.io/sig-node-kubelet#kubelet-gce-e2e-swap-ubuntu-serial)   
* \[ndixita\] Create issue for cos-cgroupv1-containerd-node-e2e-serial  and cgroupv2failing job  
* \[mmiranda96\] fixing [https://testgrid.k8s.io/sig-node-containerd\#node-e2e-features](https://testgrid.k8s.io/sig-node-containerd#node-e2e-features)  
  [https://testgrid.k8s.io/sig-node-containerd\#node-e2e-features](https://testgrid.k8s.io/sig-node-containerd#node-e2e-features) 

## 2023/09/06

Recording: [https://www.youtube.com/watch?v=lDTQSHeB-v4](https://www.youtube.com/watch?v=lDTQSHeB-v4)

Recording:  
Hosts:

- Tests: skanzhelev  
- Bugs: skanzhelev

Agenda:

## 2023/08/30

Recording: [https://www.youtube.com/watch?v=HmY5ID9LlNM](https://www.youtube.com/watch?v=HmY5ID9LlNM)   
Hosts:

- Tests: mmiranda96  
- Bugs: ndixita

   
Agenda:  
\[Harchal\] Performance issue: many pods in the same namespace take more CPU than same number of pods in multiple namespaces. Will file an issue with observations.

* \[mahamed\][https://github.com/kubernetes/test-infra/pull/29944](https://github.com/kubernetes/test-infra/pull/29944) 

## 2023/08/23

Recording: [https://www.youtube.com/watch?v=637\_82vBhqM](https://www.youtube.com/watch?v=637_82vBhqM)   
Hosts:

- Tests: ~~ndixita~~ mmiranda96  
- Bugs: mmiranda96

   
Agenda:

- Todd seems to found the reason for OOMKilled not being reported as a status: [https://github.com/kubernetes/kubernetes/issues/119600](https://github.com/kubernetes/kubernetes/issues/119600)  
- CPUManager/Topology manager: try \`decorate:true\` on the job definition

## 2023/08/16

Recording: [https://www.youtube.com/watch?v=PIqJyLT\_D\_E](https://www.youtube.com/watch?v=PIqJyLT_D_E)   
Hosts:

- Tests: mmiranda96  
- Bugs: ndixita

Agenda:

- \[tzneal\] Looking for review  
  - [~~https://github.com/kubernetes/kubernetes/pull/119765~~](https://github.com/kubernetes/kubernetes/pull/119765) ~~Use the ‘/’ mount path for NFS tests that works everywhere~~ Merged  
  - [~~https://github.com/kubernetes/kubernetes/pull/119890~~](https://github.com/kubernetes/kubernetes/pull/119890) ~~crio: increase test buffer to eliminate test flakes~~ Merged  
  - [~~https://github.com/kubernetes/kubernetes/pull/119974~~](https://github.com/kubernetes/kubernetes/pull/119974) ~~Update tests to use the latest busybox test image~~  \- Approved, waiting on merge  
- \[mmiranda96\] We need to confirm if perf tests are running on cgroup v1/v2 (v2 is expected to behave better under stress conditions)  
- \[ndixita\] [https://github.com/kubernetes/kubernetes/issues/119960](https://github.com/kubernetes/kubernetes/issues/119960) Link to other relatable bugs

## 2023/08/09

Recording: [https://www.youtube.com/watch?v=OOgdMe0TJWU](https://www.youtube.com/watch?v=OOgdMe0TJWU)   
Hosts:

- Tests: mmiranda96  
- Bugs: ~~haircommander~~ mmiranda96

Agenda:

- [https://github.com/kubernetes/test-infra/pull/30249](https://github.com/kubernetes/test-infra/pull/30249)   
- \[tzneal\] \- [https://github.com/kubernetes/kubernetes/issues/119611](https://github.com/kubernetes/kubernetes/issues/119611)  
- \[mmiranda96\] Create node conformance release branch jobs for 1.27 and 1.28 (self-assign issue)

## 2023/08/02

Recording: [https://www.youtube.com/watch?v=jxr4iMYzH2E](https://www.youtube.com/watch?v=jxr4iMYzH2E)   
Hosts:

- Tests: ndixita  
- Bugs: tzneal

Agenda:

- \[ndixita\] Looking for more people to engage in the meeting by leading bug triage or test failures. Please reach out to \[SergeyKanzhelev\] if you are interested.  
  - ndixita doing test failures triage today.  
  - tzneal doing bug triage today.  
- \[ndixita\] Test Failures  
  - \[fromani\] gce device plugin presubmit jobs seems to be broken. Any insights, or tips to debug?  
    - Issue: [https://github.com/kubernetes/kubernetes/issues/119730](https://github.com/kubernetes/kubernetes/issues/119730)  
    - Testing PR: [https://github.com/kubernetes/kubernetes/pull/119590](https://github.com/kubernetes/kubernetes/pull/119590)   
  - Multi-numa still failing: [https://github.com/kubernetes/test-infra/pull/29717\#issuecomment-1584967574](https://github.com/kubernetes/test-infra/pull/29717#issuecomment-1584967574) [https://github.com/kubernetes/kubernetes/issues/119601](https://github.com/kubernetes/kubernetes/issues/119601)   
    [https://github.com/kubernetes/test-infra/pull/30249](https://github.com/kubernetes/test-infra/pull/30249)  
  - Arm failures fixed: [https://testgrid.k8s.io/sig-node-kubelet\#kubelet-gce-e2e-arm64-ubuntu-serial](https://testgrid.k8s.io/sig-node-kubelet#kubelet-gce-e2e-arm64-ubuntu-serial)  
  - Fix: https://github.com/kubernetes/kubernetes/pull/119603  
  - [https://testgrid.k8s.io/sig-node-containerd\#containerd-node-e2e-1.7](https://testgrid.k8s.io/sig-node-containerd#containerd-node-e2e-1.7)   
    - [https://github.com/kubernetes/kubernetes/issues/119600](https://github.com/kubernetes/kubernetes/issues/119600)   
    - Fix for oom kill test: [https://github.com/kubernetes/kubernetes/pull/119670](https://github.com/kubernetes/kubernetes/pull/119670)  
    - “Summary API \[NodeConformance\] when querying /stats/summary should report resource usage through the stats api” still failing.  
- \[tzneal\] Odd segfault in /bin/top that only occurs with old busybox images on arm64. Fixed with newer images, PR to update and fix some arm64 tests is at [https://github.com/kubernetes/kubernetes/pull/119636](https://github.com/kubernetes/kubernetes/pull/119636)   
  - Related, some test images are built on centos which is no longer updated (e.g. [https://github.com/kubernetes/kubernetes/blob/master/test/images/volume/nfs/BASEIMAGE](https://github.com/kubernetes/kubernetes/blob/master/test/images/volume/nfs/BASEIMAGE) )  
- 

## 2023/07/26

Recording: [https://www.youtube.com/watch?v=iJUsH4BfYTY](https://www.youtube.com/watch?v=iJUsH4BfYTY)   
Hosts:

- Tests: SergeyKanzhelev  
- Bugs: mmiranda96

Agenda:

- \[SergeyKanzhelev\] More people engaging with leading the meeting. Please reach out to me if you are interested.  
  - Mike will do bug triage today.  
- \[fromani\] it seems the job [pull-kubernetes-e2e-gce-device-plugin-gpu](https://prow.k8s.io/job-history/gs/kubernetes-jenkins/pr-logs/directory/pull-kubernetes-e2e-gce-device-plugin-gpu) is broken for stable releases (e.g. 1.27). Xref: [https://prow.k8s.io/pr-history/?org=kubernetes\&repo=kubernetes\&pr=119590](https://prow.k8s.io/pr-history/?org=kubernetes&repo=kubernetes&pr=119590)   
- \[SergeyKanzhelev\] Multi-numa still failing: [https://github.com/kubernetes/test-infra/pull/29717\#issuecomment-1584967574](https://github.com/kubernetes/test-infra/pull/29717#issuecomment-1584967574) [https://github.com/kubernetes/kubernetes/issues/119601](https://github.com/kubernetes/kubernetes/issues/119601)   
  - Arm failing: [https://testgrid.k8s.io/sig-node-kubelet\#kubelet-gce-e2e-arm64-ubuntu-serial](https://testgrid.k8s.io/sig-node-kubelet#kubelet-gce-e2e-arm64-ubuntu-serial) \[mention @upodroid\] [**https://github.com/kubernetes/kubernetes/issues/119599**](https://github.com/kubernetes/kubernetes/issues/119599)   
  - [https://testgrid.k8s.io/sig-node-containerd\#containerd-node-e2e-1.7](https://testgrid.k8s.io/sig-node-containerd#containerd-node-e2e-1.7) [https://github.com/kubernetes/kubernetes/issues/119600](https://github.com/kubernetes/kubernetes/issues/119600)   
  - [https://github.com/kubernetes/kubernetes/issues/119602](https://github.com/kubernetes/kubernetes/issues/119602)   
  - \[Harshal\] is it related to [https://github.com/kubernetes/kubernetes/pull/119486](https://github.com/kubernetes/kubernetes/pull/119486) ?  
- \[ndixita\] suggestion for adding dashboards for each architecture and cloud provider [https://github.com/kubernetes/test-infra/pull/29969](https://github.com/kubernetes/test-infra/pull/29969)   
- \[mahamed\] [https://github.com/kubernetes/test-infra/issues/29946](https://github.com/kubernetes/test-infra/issues/29946) node prowjob overhaul

## 2023/07/19

Recording: [https://www.youtube.com/watch?v=PvaIQwVaCEs](https://www.youtube.com/watch?v=PvaIQwVaCEs)

\[swsehgal\]: Multi-numa failing: [https://github.com/kubernetes/test-infra/pull/29717\#issuecomment-1584967574](https://github.com/kubernetes/test-infra/pull/29717#issuecomment-1584967574)

* Ran CPU manager, Memory Manager and Topology Manager e2e tests locally on a multi-numa system. All are passing.  
* Topology Manager metric test is failing in the RedHat environment ( not the same failure in u/s CI; tests are skipped in u/s CI): Looking into this.  Wasn’t able to reproduce it locally.  
* For CPU Manager and memory manager, kubelet fails in upstream CI environment. Perhaps issue with job config.  
  * PR with a potential fix (update node test args): [https://github.com/kubernetes/test-infra/pull/30072](https://github.com/kubernetes/test-infra/pull/30072) 

## 2023/07/07 canceled due to host availability

## 2023/07/05

Recording: [https://www.youtube.com/watch?v=G6j0hnwsEXQ](https://www.youtube.com/watch?v=G6j0hnwsEXQ) 

- Review the umbrella issue: [https://github.com/kubernetes/kubernetes/issues/118441\#issuecomment-1578138427](https://github.com/kubernetes/kubernetes/issues/118441#issuecomment-1578138427)   
- Multi-numa failing: [https://github.com/kubernetes/test-infra/pull/29717\#issuecomment-1584967574](https://github.com/kubernetes/test-infra/pull/29717#issuecomment-1584967574)

## 2023/06/28

Recording: [https://www.youtube.com/watch?v=U9yophjx23Q](https://www.youtube.com/watch?v=U9yophjx23Q)

- \[SergeyKanzhelev\] prepull: [https://github.com/kubernetes/kubernetes/pull/118747](https://github.com/kubernetes/kubernetes/pull/118747)   
- 

## 2023/06/21 \[Cancelled\]

## 2023/06/14

Recording: [https://www.youtube.com/watch?v=TDBLb4\_sj7w](https://www.youtube.com/watch?v=TDBLb4_sj7w) 

Agenda:

- ARM64 failures: [https://testgrid.k8s.io/sig-node-kubelet\#kubelet-gce-e2e-arm64-ubuntu-serial](https://testgrid.k8s.io/sig-node-kubelet#kubelet-gce-e2e-arm64-ubuntu-serial)  
  - Need CI test as well  
- Multi-numa failing: [https://github.com/kubernetes/test-infra/pull/29717\#issuecomment-1584967574](https://github.com/kubernetes/test-infra/pull/29717#issuecomment-1584967574)  
  - Updating the base image may be relevant here (cos 93\)  
  - Mike to send a PR today to bump the version  
- Cos-containerd-node-e2e-serial failure : [https://github.com/kubernetes/kubernetes/issues/118660](https://github.com/kubernetes/kubernetes/issues/118660)

[https://github.com/kubernetes/kubernetes/issues/118441](https://github.com/kubernetes/kubernetes/issues/118441)

## 2023/06/07

Recording: [https://www.youtube.com/watch?v=1u6--D2Y9LU](https://www.youtube.com/watch?v=1u6--D2Y9LU) 

- \[pacoxu\] Dims opened an umbrella issues [https://github.com/kubernetes/kubernetes/issues/118441\#issuecomment-1578138427](https://github.com/kubernetes/kubernetes/issues/118441#issuecomment-1578138427) Some updates from my side. Some may need platform permissions and not sure if we can be granted.  
  - AI: ownership and the purpose of this: prowjob\_name: ci-kubernetes-e2e-node-canary prowjob\_config\_url: https://git.k8s.io/test-infra/config/jobs/kubernetes/sig-testing/kubetest-canaries.yaml  
  -   
  -   
- ARM and multi-numa periodics status  
  - Multi NUMA PR: [https://github.com/kubernetes/test-infra/pull/29717](https://github.com/kubernetes/test-infra/pull/29717)  
  - ARM is still failing: [https://testgrid.k8s.io/sig-node-kubelet\#kubelet-gce-e2e-arm64-ubuntu-serial](https://testgrid.k8s.io/sig-node-kubelet#kubelet-gce-e2e-arm64-ubuntu-serial) 

## 2023/05/31

Recording: [https://www.youtube.com/watch?v=4PZi50cP-UY](https://www.youtube.com/watch?v=4PZi50cP-UY) 

Agenda:

- StandaloneMode: green now: [https://testgrid.k8s.io/sig-node-release-blocking\#node-kubelet-containerd-standalone-mode-all-alpha](https://testgrid.k8s.io/sig-node-release-blocking#node-kubelet-containerd-standalone-mode-all-alpha) Ideally need to add more tests  
- ARM64 still failing:  
  - [https://github.com/kubernetes/test-infra/pull/29617\#issuecomment-1570293715](https://github.com/kubernetes/test-infra/pull/29617#issuecomment-1570293715)  
- Create issue for periodics of multi-NUMA  
- Periodics vs. presubmits jobs:   
  - presubmits are degrading  
  - presubmits diverge from periodics  
  - \[ffromani\] example is topology tests that investigating now. Will ask sig testing if there are any established approaches to keep those green  
- \[Dixita\] First draft for guidance around test coverage [https://docs.google.com/document/d/1P1X9Jr2PYFiC6xNF-9RtgrPuc9PsSmikCngGTfGVxjY/edit?usp=sharing](https://docs.google.com/document/d/1P1X9Jr2PYFiC6xNF-9RtgrPuc9PsSmikCngGTfGVxjY/edit?usp=sharing)  
  - Please take a look at it and drop feedback.  
    \- grant permission to [kubernetes-sig-node-test-failures](https://groups.google.com/forum/#!forum/kubernetes-sig-node-test-failures) please ([kubernetes-sig-node-test-failures@googlegroups.com](mailto:kubernetes-sig-node-test-failures@googlegroups.com))  
- 

## 2023/05/24

Recording: [https://www.youtube.com/watch?v=X87Lmqmf3QQ](https://www.youtube.com/watch?v=X87Lmqmf3QQ)   
Agenda:

- \[tzneal\] Added periodics for standalone kubelet tests, but they are failing [https://testgrid.k8s.io/sig-node-release-blocking\#node-kubelet-containerd-standalone-mode](https://testgrid.k8s.io/sig-node-release-blocking#node-kubelet-containerd-standalone-mode)  
  - [https://prow.k8s.io/view/gs/kubernetes-jenkins/logs/ci-kubernetes-node-e2e-containerd-standalone-mode-all-alpha/1661213193298513920](https://prow.k8s.io/view/gs/kubernetes-jenkins/logs/ci-kubernetes-node-e2e-containerd-standalone-mode-all-alpha/1661213193298513920)  
  - W0524 03:37:35.874\] 2023/05/24 03:37:35 main.go:328: Something went wrong: failed to prepare test environment: \--provider=gce boskos failed to acquire project: resources not found  
  - Asked in \#sig-k8s-infra: [https://kubernetes.slack.com/archives/CCK68P2Q2/p1684932723841139](https://kubernetes.slack.com/archives/CCK68P2Q2/p1684932723841139)  
  - Try compare with [https://testgrid.k8s.io/sig-node-containerd\#containerd-node-e2e-features-1.7](https://testgrid.k8s.io/sig-node-containerd#containerd-node-e2e-features-1.7)   
- Containerd features tests are not running: [https://testgrid.k8s.io/sig-node-containerd\#node-e2e-features](https://testgrid.k8s.io/sig-node-containerd#node-e2e-features) 

## 2023/05/17

Recording: [https://www.youtube.com/watch?v=5IC3JLbrk-A](https://www.youtube.com/watch?v=5IC3JLbrk-A) 

Agenda:

- \[SergeyKanzhelev\] Stress test: [https://github.com/kubernetes/kubernetes/pull/117439](https://github.com/kubernetes/kubernetes/pull/117439)  
- More test areas:  
  - ARM tests: [https://github.com/kubernetes/test-infra/pull/29192](https://github.com/kubernetes/test-infra/pull/29192) and [https://github.com/kubernetes/kubernetes/pull/117017](https://github.com/kubernetes/kubernetes/pull/117017)  
  - Run StandaloneMode kubelet tests in periodics: [https://github.com/kubernetes/test-infra/pull/29042](https://github.com/kubernetes/test-infra/pull/29042)   
    - Tzneal can do it  
  - Run multi-numa in periodics: [https://github.com/kubernetes/test-infra/blob/master/jobs/e2e\_node/image-config-serial-multi-numa.yaml](https://github.com/kubernetes/test-infra/blob/master/jobs/e2e_node/image-config-serial-multi-numa.yaml)   
    - Swati can do it  
  - Jobs to release branches for hotfix validation  
    - Fork-per-release: [https://github.com/kubernetes/test-infra/pull/29483/commits/9ef5f916f78323f99857ff16cbb6b03f665eed60](https://github.com/kubernetes/test-infra/pull/29483/commits/9ef5f916f78323f99857ff16cbb6b03f665eed60)   
    - \[Todd\] [https://github.com/kubernetes/test-infra/blob/master/releng/config-forker/README.md](https://github.com/kubernetes/test-infra/blob/master/releng/config-forker/README.md)   
    - Images may be used there that are “too fresh”  
    - \[mmiranda96\] in these cases we can fork manually  
  - Previous releases periodics with specific images

## 2023/05/10

Recording: [https://www.youtube.com/watch?v=JsjBNgVb6Po](https://www.youtube.com/watch?v=JsjBNgVb6Po) 

Agenda:

\[swsehgal\] Would like help with [https://github.com/kubernetes/test-infra/pull/29483](https://github.com/kubernetes/test-infra/pull/29483)

* Trying to enable jobs for 1.25, 1.26 and 1.27 release branches so that device manager tests can execute in the CI  
  * Unit tests are failing  
  * How can I ensure that these jobs are properly added to testgrid?  
  * \[mmiranda96\] Previous PR that contains release branch SIG-Node tests: [https://github.com/kubernetes/test-infra/pull/29038](https://github.com/kubernetes/test-infra/pull/29038)

## 2023/05/03

Recording: [https://www.youtube.com/watch?v=SAUmnsx\_sG8](https://www.youtube.com/watch?v=SAUmnsx_sG8)

Triage only

## 2023/04/26

Recording: [Kubernetes SIG Node CI 20230426](https://www.youtube.com/watch?v=-Aq6ZB7Bb0o) 

- \[mmiranda96\] [https://github.com/kubernetes/test-infra/issues/29308](https://github.com/kubernetes/test-infra/issues/29308)  
- \[ndixita\] Guidance doc for Sig Node CI Test Coverage

## 2023/04/18 Canceled for KubeCon

## 2023/03/29

Recording: [https://www.youtube.com/watch?v=TaUK0qYwtwA](https://www.youtube.com/watch?v=TaUK0qYwtwA) 

- \[akhil\] [https://github.com/kubernetes/kubernetes/issues/116944](https://github.com/kubernetes/kubernetes/issues/116944)

## 2023/03/22

Recording: [https://www.youtube.com/watch?v=d6TZ5AywdnM](https://www.youtube.com/watch?v=d6TZ5AywdnM)

- \[tzneal\] Eviction Flakes \- [https://testgrid.k8s.io/sig-node-containerd\#node-kubelet-containerd-eviction](https://testgrid.k8s.io/sig-node-containerd#node-kubelet-containerd-eviction)   
  - Fix for Pid Pressure by Todd \- will be posted soon  
  - Test: [https://github.com/kubernetes/kubernetes/blob/master/test/e2e\_node/eviction\_test.go\#L500](https://github.com/kubernetes/kubernetes/blob/master/test/e2e_node/eviction_test.go#L500)   
  - PR: [https://github.com/kubernetes/kubernetes/pull/116862](https://github.com/kubernetes/kubernetes/pull/116862)   
  - Update\[tzneal\]: I was wrong on this, I think it’s actually caused by [https://github.com/kubernetes/kubernetes/issues/115215](https://github.com/kubernetes/kubernetes/issues/115215) .   
- New failure:  
  - E2eNode Suite.\[It\] \[sig-node\] MirrorPod when create a mirror pod without changes should successfully recreate when file is removed and recreated \[NodeConformance\]  
  -    
  - [https://testgrid.k8s.io/sig-node-release-blocking\#ci-crio-cgroupv1-node-e2e-conformance](https://testgrid.k8s.io/sig-node-release-blocking#ci-crio-cgroupv1-node-e2e-conformance)   
  - AI: Ryan to take a look  
    - [https://github.com/kubernetes/kubernetes/issues/116714](https://github.com/kubernetes/kubernetes/issues/116714)

  - [https://github.com/kubernetes/kubernetes/issues/116874](https://github.com/kubernetes/kubernetes/issues/116874) 

- Flakes/failure:  
  - E2eNode Suite.\[It\] \[sig-node\] MirrorPodWithGracePeriod when create a mirror pod and the container runtime is temporarily down during pod termination \[NodeConformance\] \[Serial\] \[Disruptive\] the mirror pod should terminate successfully  
  -    
  - [https://testgrid.k8s.io/sig-node-release-blocking\#node-kubelet-serial-containerd\&width=20](https://testgrid.k8s.io/sig-node-release-blocking#node-kubelet-serial-containerd&width=20) 

  Perma failure here:

    E2eNode Suite.\[It\] \[sig-node\] MirrorPodWithGracePeriod when create a mirror pod and the container runtime is temporarily down during pod termination \[NodeConformance\] \[Serial\] \[Disruptive\] the mirror pod should terminate successfully

  [https://testgrid.k8s.io/sig-node-cri-o\#node-kubelet-serial-crio](https://testgrid.k8s.io/sig-node-cri-o#node-kubelet-serial-crio) 


  Introduced here: [https://github.com/kubernetes/kubernetes/pull/113145/](https://github.com/kubernetes/kubernetes/pull/113145/) 

- Not working:  
  - [https://testgrid.k8s.io/sig-node-containerd\#containerd-e2e-ubuntu](https://testgrid.k8s.io/sig-node-containerd#containerd-e2e-ubuntu)  
  - [https://github.com/kubernetes/kubernetes/issues/116873](https://github.com/kubernetes/kubernetes/issues/116873) 

- Pod Overhead related

  E2eNode Suite.\[It\] \[sig-node\] Kubelet PodOverhead handling \[LinuxOnly\] PodOverhead cgroup accounting On running pod with PodOverhead defined Pod cgroup should be sum of overhead and resource limits

   

  [https://testgrid.k8s.io/sig-node-cri-o\#ci-crio-cgroupv1-node-e2e-unlabelled\&width=20](https://testgrid.k8s.io/sig-node-cri-o#ci-crio-cgroupv1-node-e2e-unlabelled&width=20) 

  AI: Todd to investigate

- [https://github.com/kubernetes/kubernetes/issues/116864](https://github.com/kubernetes/kubernetes/issues/116864)  
- Doesn’t appear to be related to the changes to pod resource calculations.  This uses a custom runtime class and support for setting up the test infra with that runtime class is not implemented for cri-o.

	

## 2023/03/15

Recording: [https://www.youtube.com/watch?v=D\_5tx6lkxY4](https://www.youtube.com/watch?v=D_5tx6lkxY4)

- \[SergeyKanzhelev\] Standalone tests:  
  - [https://github.com/kubernetes/kubernetes/pull/116628](https://github.com/kubernetes/kubernetes/pull/116628)  
  - [https://github.com/kubernetes/kubernetes/pull/116631](https://github.com/kubernetes/kubernetes/pull/116631)   
- \[SergeyKanzhelev\] testing limits in e2e: TODO  
- \[SergeyKanzhelev\] there is a flake: [https://testgrid.k8s.io/sig-node-release-blocking\#node-kubelet-serial-containerd](https://testgrid.k8s.io/sig-node-release-blocking#node-kubelet-serial-containerd) 


## 2023/03/08

Recording: [https://www.youtube.com/watch?v=luwWBmxyngk](https://www.youtube.com/watch?v=luwWBmxyngk)

- \[SergeyKanzhelev\] [https://github.com/kubernetes/test-infra/issues/28888](https://github.com/kubernetes/test-infra/issues/28888)  
- \[SergeyKanzhelev\] [https://github.com/kubernetes/test-infra/pull/28919](https://github.com/kubernetes/test-infra/pull/28919#discussion_r1128352969) Apply USE\_TEST\_INFRA\_LOG\_DUMPING to sig node jobs  
  - @xmcqueen  
- \[fromani\] not urgent, mostly to socialize the idea:  
  - promoting a subset of podresources e2e tests to NodeConformance  
  - To be done around the timeframe of podresources GA (should not block the GA’ing)  
  - Context: podresources [endpoint on windows](https://github.com/kubernetes/kubernetes/pull/115133/) needs e2e tests. On windows they run conformance and nodeconformance tests by default.  
- \[SergeyKanzhelev\] Testing previous releases [https://testgrid.k8s.io/sig-node-containerd\#node-conformance-release-1.24](https://testgrid.k8s.io/sig-node-containerd#node-conformance-release-1.24)  
- \[vinaykul\] In-place pod resize CI tests merged yesterday.  
  - Multiple failures from ‘Insufficient cpu’ on small node (2000m allocatable)  
  - Potential fix [https://github.com/kubernetes/kubernetes/pull/116372](https://github.com/kubernetes/kubernetes/pull/116372)

AI:

- Ping sig storage on [https://github.com/kubernetes/kubernetes/issues/116357](https://github.com/kubernetes/kubernetes/issues/116357) 

## 2023/03/01

Recording: [https://www.youtube.com/watch?v=BIhQ69FNYJA](https://www.youtube.com/watch?v=BIhQ69FNYJA) 

- \[mmiranda96\] [https://github.com/kubernetes/test-infra/issues/28627](https://github.com/kubernetes/test-infra/issues/28627)  
  - Start with NodeConformance on supported k8s release branches.  
  - Add @xmcqueen as reviewer  
- [https://github.com/kubernetes/kubernetes/pull/115984](https://github.com/kubernetes/kubernetes/pull/115984) 

## 2023/02/22 \[cancelled\]

## 2023/02/15 \[cancelled\]

Agenda:

- \[fromani\] Promoting some node e2e tests to NodeConformance: best way forward  
  - This will provide test coverage for podresources API on windows  
  - Fromani to discuss this offline (not very urgent, no worries)

## 2023/02/08

Recording: [https://www.youtube.com/watch?v=\_4K8JD7Zejo](https://www.youtube.com/watch?v=_4K8JD7Zejo)   
Agenda:

- \[pacoxu\] I want to add some CI for [quota monitoring of ephemeral storage](https://github.com/kubernetes/enhancements/issues/1029). (There is already e2e test to test on configmap editing, what we should do is providing a cluster that XFS project quotas is enabled): [https://github.com/kubernetes/test-infra/issues/28614](https://github.com/kubernetes/test-infra/issues/28614)   
  - The correct bug fix fo the fsquota issue is [https://github.com/kubernetes/kubernetes/pull/115314](https://github.com/kubernetes/kubernetes/pull/115314) .  
- \[SergeyKanzhelev\] Check on last meting AIs:  
  - [https://github.com/kubernetes/test-infra/issues/28627](https://github.com/kubernetes/test-infra/issues/28627)  
  - [https://github.com/kubernetes/kubernetes/issues/115372](https://github.com/kubernetes/kubernetes/issues/115372)

## 2023/02/01

Recording: [https://www.youtube.com/watch?v=t9brdyaGo3Y](https://www.youtube.com/watch?v=t9brdyaGo3Y) 

Agenda:

- \[SergeyKanzhelev\] Test coverage for previous releases  
  - Let’s add some NodeConformance tests to sig-node-releas-blocking for supported versions  
  - Mike is to open an issue for this: [https://github.com/kubernetes/test-infra/issues/28627](https://github.com/kubernetes/test-infra/issues/28627)  
  - Also let’s lock the OS to the specific release  
  - Mike to decide on new OS for 1.27

- \[Mike Miranda\] swap on fedora: [https://testgrid.k8s.io/sig-node-kubelet\#kubelet-gce-e2e-swap-fedora](https://testgrid.k8s.io/sig-node-kubelet#kubelet-gce-e2e-swap-fedora)  
  - Tracking issue: [https://github.com/kubernetes/kubernetes/issues/115372](https://github.com/kubernetes/kubernetes/issues/115372)  
  - Peter will take a look  
  - Mike maybe take a look into adding memory swap presubmit

## 2023/01/25

Recording: [https://youtu.be/t1kbFCaeSSE](https://youtu.be/t1kbFCaeSSE) 

Agenda:

- \[SergeyKanzhelev\] no stress tests: [https://github.com/kubernetes/kubernetes/pull/115143](https://github.com/kubernetes/kubernetes/pull/115143)   
  - \[Brian\] soak tests exists but broken fro a long time  
  - \[Ryan\] gRPC \- is it also needs to be fixed? What tests needs to run?  
  - \[Brian\] qq: is there already a place where it will fit in? Maybe we can add a place first and clear the path for contributors to add test cases.  
  - \[Sergey\] curious if evented PLEG has been looking into stress testing  
  - \[Harshal\] adding e2e CI tests jobs for containerd and cri-o. Have a presubmit job now, debugging it.  
    - Stress testing and different scenarios is a next step

## 2023/01/18

Recording: [https://youtu.be/hep7StWT8u0](https://youtu.be/hep7StWT8u0)   
Agenda:

- \[swsehgal\] Need some guidance on the process of publishing images that are used in node e2e tests  
  - Context: Device Manager Bug [https://github.com/kubernetes/kubernetes/pull/114640](https://github.com/kubernetes/kubernetes/pull/114640)   
  - Scenarios where on node reboot or kubelet restart, device plugin pod is not recovered before an application pod consuming device.  
  - To reproduce the issue and for e2e testing, sample device plugin (which is a device plugin implemented in tree for testing) was modified to control its registration process  
  - Changes related to sample device plugin was split into a separate PR [https://github.com/kubernetes/kubernetes/pull/115107](https://github.com/kubernetes/kubernetes/pull/115107)   
  - How can we get this image pushed to the Kubernetes registry so it can be correctly consumed in 114640? [Kubernetes image promotion process](https://github.com/kubernetes/enhancements/tree/master/keps/sig-release/1734-k8s-image-promoter#promotion-process) indicates that maintainer involvement might be needed here.  
  - Examples:  
    - [https://github.com/kubernetes/kubernetes/pull/109551/files](https://github.com/kubernetes/kubernetes/pull/109551/files)  
    - [https://github.com/kubernetes/k8s.io/pull/4391](https://github.com/kubernetes/k8s.io/pull/4391)   
    - [https://github.com/kubernetes/k8s.io/blob/71231519d8f36b71b2c218ed3a993c64d63d0882/k8s.gcr.io/images/k8s-staging-e2e-test-images/images.yaml\#L149](https://github.com/kubernetes/k8s.io/blob/71231519d8f36b71b2c218ed3a993c64d63d0882/k8s.gcr.io/images/k8s-staging-e2e-test-images/images.yaml#L149) 

## 2023/01/11

Recording: [https://youtu.be/N97z4wGoIl0](https://youtu.be/N97z4wGoIl0)   
Agenda:

- \[Francesco\] device plugins to use in e2e tests  
  - Why we need them  
  - Current issues  
  - Discussion about future plans  
      
  - Synthetic devices are better than nothing. They would work on any machine and vendor-neutral  
  - Device plugin from kubevirt was good as it was vendor neutral and didn’t require hardware: [https://kubevirt.io/2018/KVM-Using-Device-Plugins.html](https://kubevirt.io/2018/KVM-Using-Device-Plugins.html)   
  - In search of real world devices:  
    - Investigate GPU devices \- can we run with them periodically?  
  - \[Ryan\] good approach may be a combination \- running presubmits with synthetic and GPU devices periodically  
  - \[Francesco\] will investigate  
- [**https://github.com/kubernetes/test-infra/pull/28369**](https://github.com/kubernetes/test-infra/pull/28369) **enable multi-numa tests**  
- [**https://github.com/kubernetes/community/pull/7021**](https://github.com/kubernetes/community/pull/7021) **how to write good tests**

## 2023/01/04

Recording: [https://www.youtube.com/watch?v=wyc4k1ERDEA](https://www.youtube.com/watch?v=wyc4k1ERDEA) 

Agenda:

- Triage  
- All CRI-O tests are red: [https://testgrid.k8s.io/sig-node-cri-o](https://testgrid.k8s.io/sig-node-cri-o)  
- 
