# Kubernetes SIG-Node CI subgroup notes

## Dec 18, 2024  
Hosts:  
	Tests:  
	Bugs triaging

Agenda:

- Test-Infra Cleanup  
  - Dropped NodeSpecialFeature / NodeAlphaFeature  
    - Fall out: [https://github.com/kubernetes/test-infra/pull/33996](https://github.com/kubernetes/test-infra/pull/33996)  
  - Aiming to deprecate NodeFeature by replicating with NodeFeature  
    - [https://github.com/kubernetes/kubernetes/pull/129166](https://github.com/kubernetes/kubernetes/pull/129166)  
  - 

## Dec 11, 2024  
Hosts:  
	Tests:  
Bugs triaging:  
      
Agenda:

- [https://github.com/kubernetes/enhancements/tree/master/keps/sig-testing/3041-node-conformance-and-features\#goals](https://github.com/kubernetes/enhancements/tree/master/keps/sig-testing/3041-node-conformance-and-features#goals)  
  - [https://github.com/kubernetes/kubernetes/pull/128923](https://github.com/kubernetes/kubernetes/pull/128923)  
  - [https://github.com/kubernetes/test-infra/pull/33828](https://github.com/kubernetes/test-infra/pull/33828)   
  - TODO:  
    - [Nodefeature](http://github.com/kubernetes/kubernetes/tree/master/test/e2e/nodefeature) to feature  
    - Convert to label filter instead of using skip/focus  
- [https://github.com/kubernetes/kubernetes/pull/128880](https://github.com/kubernetes/kubernetes/pull/128880)  
- [https://github.com/kubernetes/kubernetes/pull/128889](https://github.com/kubernetes/kubernetes/pull/128889)

Action items:  
\- Follow up on https://github.com/kubernetes/test-infra/issues/32567

## Dec 4, 2024 \[just 3 people joined. Will do triage offline\]

## Nov 27, 2024  
Cancelled due to U.S. Holiday

## Nov 20, 2024

Hosts:   
	Tests: Kevin  
Bugs triaging:   
	  
Agenda:  
Need to create tickets for this

- Swap Tests are bonked. Need to create a ticket to investigate.  
- Huge Page Test Failures  
- Device Plugin

## Nov 13, 2024 \[Canceled for KubeCon\]

## Nov 6, 2024

Hosts:   
	Tests:   
Bugs triaging: 

Agenda:

- \[anish\] Why are evictions / cpu,memory,topology managers e2e tests not considered release blocking? Even though these are stable features?

## Oct 30, 2024 \[Canceled\]

Hosts:   
	Tests:   
Bugs triaging:   
	  
Agenda:

## Oct 23, 2024

Recording: [https://www.youtube.com/watch?v=Y8bUTXy3FGs](https://www.youtube.com/watch?v=Y8bUTXy3FGs)

Hosts:   
	Tests:   
Bugs triaging:   
	  
Agenda:

- Combine sig node/sig node CI board?  
  - Originally it was separated to onboard new members to be able to do reviews without needing to worry about production code  
  - Generally, this meeting should be focused on CI, so maybe defer PR triage  
  - Add a special label to PRs, when it’s present remove from one board/add it to the other  
- CRI proxy PR merged, now more tests can be added to test different CRI scenarios   
  - [https://github.com/kubernetes/kubernetes/pull/127495](https://github.com/kubernetes/kubernetes/pull/127495)  
  - How it can be used: [https://github.com/kubernetes/kubernetes/pull/121604](https://github.com/kubernetes/kubernetes/pull/121604) 

## Oct 16, 2024

Recording: [https://www.youtube.com/watch?v=Y8bUTXy3FGs](https://www.youtube.com/watch?v=Y8bUTXy3FGs)

Hosts:   
	Tests:   
Bugs triaging:   
	  
Agenda:

- \[KubeTest\] Migration to kubetest  
  - https://github.com/kubernetes/test-infra/issues/32567  
  - Presubmits first  
    - A lot of problems  
    - [https://github.com/elieser1101](https://github.com/elieser1101) is owner for this  
- EventedPleg: [https://github.com/kubernetes/test-infra/issues/33666](https://github.com/kubernetes/test-infra/issues/33666)  
- 

## Oct 9, 2024

Recording: [https://www.youtube.com/watch?v=GKrlW1LDXz0](https://www.youtube.com/watch?v=GKrlW1LDXz0)

Hosts:   
	Tests: Anish  
Bugs triaging:   
	  
Agenda:

Oct 2

Recording: [https://www.youtube.com/watch?v=80260g3EEv8](https://www.youtube.com/watch?v=80260g3EEv8)

Sep 25, 2024  
Recording: [https://www.youtube.com/watch?v=gcX6sDoibM4](https://www.youtube.com/watch?v=gcX6sDoibM4)

Agenda:

- [https://github.com/kubernetes/kubernetes/issues/127610](https://github.com/kubernetes/kubernetes/issues/127610)   
- \[ffromani\] (can attend only 1st half) Chicken and egg: [https://github.com/kubernetes/kubernetes/pull/120661](https://github.com/kubernetes/kubernetes/pull/120661) and [https://github.com/kubernetes/kubernetes/pull/127506](https://github.com/kubernetes/kubernetes/pull/127506)   
- CRI proxy work started: [https://github.com/kubernetes/kubernetes/pull/127495](https://github.com/kubernetes/kubernetes/pull/127495) Mostly FYI

Sep 18, 2024  
Recording: [https://www.youtube.com/watch?v=Y6slxvO6Hv8](https://www.youtube.com/watch?v=Y6slxvO6Hv8)

Hosts:  
	Tests:  
	Bugs triage:

Agenda:

- Another ping about the flake: [https://kubernetes.slack.com/archives/C0BP8PW9G/p1726622928011249?thread\_ts=1718369837.055379\&cid=C0BP8PW9G](https://kubernetes.slack.com/archives/C0BP8PW9G/p1726622928011249?thread_ts=1718369837.055379&cid=C0BP8PW9G)  
  - [https://github.com/kubernetes/kubernetes/issues/122270](https://github.com/kubernetes/kubernetes/issues/122270)   
- CRI proxy  
  - Injecting failures is a good idea  
  - If easy to set up \- maybe set up everywhere. If not \- let’s only do it per test  
    - The main concern to not leak tests into each other  
    - 

Sep 11, 2024  
Recording: [https://www.youtube.com/watch?v=vmH-6iWjWPM](https://www.youtube.com/watch?v=vmH-6iWjWPM)

Hosts:   
	Tests:   
Bugs triaging:   
	  
Agenda:

Sep 4, 2024  
Recording: [https://www.youtube.com/watch?v=1DiDFkYhpi4](https://www.youtube.com/watch?v=1DiDFkYhpi4)

Hosts:   
	Tests:   
Bugs triaging: anish  
	  
Agenda:

Aug 28, 2024  
Recording: [https://www.youtube.com/watch?v=tE4uO6Gj4sM](https://www.youtube.com/watch?v=tE4uO6Gj4sM)

Hosts:   
	Tests:   
Bugs triaging:   
	  
Agenda:

- \[anish\] will join in the second half of the meeting. Taking a shot at deflaking eviction tests \- [https://github.com/kubernetes/kubernetes/issues/123591](https://github.com/kubernetes/kubernetes/issues/123591)  
  - Cadvisor cache seems to be in sync 

Aug 21, 2024

Recording: [https://www.youtube.com/watch?v=EgqFB0PDb0g](https://www.youtube.com/watch?v=EgqFB0PDb0g)

Aug 7, 2024  
Recording: [https://www.youtube.com/watch?v=g6U40nR\_tRU](https://www.youtube.com/watch?v=g6U40nR_tRU)

Hosts:   
	Tests:   
Bugs triaging:   
	  
Agenda:

- Kevinn for approver: [https://github.com/kubernetes/test-infra/pull/33255](https://github.com/kubernetes/test-infra/pull/33255)  
- Add tests lanes for 1.31 \- Kevin will take it  
- Sergey: add workflows to auto-populate issues

Jul 31, 2024  
Recording: [https://www.youtube.com/watch?v=yfSd6ezXWIs](https://www.youtube.com/watch?v=yfSd6ezXWIs)

Hosts:   
	Tests: Peter  
Bugs triaging: Anish  
	  
Agenda:

- \[harche\] \- are we identifying cgroup v1 and v2 specific CI jobs?

Jul 24, 2024  
Recording: [https://www.youtube.com/watch?v=Wz0Dzo\_f4jg](https://www.youtube.com/watch?v=Wz0Dzo_f4jg)

Hosts:   
	Tests: Kevin  
Bugs triaging: Peter  
	  
Agenda:

- AI: migrate to the new project boards.  
- AI: ask about perf dashboard  
- \[harche\] \- [https://github.com/kubernetes/kubernetes/issues/125720](https://github.com/kubernetes/kubernetes/issues/125720)   
  - Also: [https://github.com/kubernetes/kubernetes/issues/125409](https://github.com/kubernetes/kubernetes/issues/125409)  
  - Also potential fix for the behavior : [https://biriukov.dev/docs/page-cache/6-cgroup-v2-and-page-cache/\#writeback-and-io](https://biriukov.dev/docs/page-cache/6-cgroup-v2-and-page-cache/#writeback-and-io)   
  -  

Jul 17, 2024

Jul 10, 2024   
Recording: [https://www.youtube.com/watch?v=see5mwuN0YA](https://www.youtube.com/watch?v=see5mwuN0YA) 

Host:  
Tests:   
Bugs triaging: anish

Agenda:

- \[anishshah\] New tests failing:  
  - Podspidlimit \- [https://github.com/kubernetes/kubernetes/issues/126007](https://github.com/kubernetes/kubernetes/issues/126007)  
  - OOMKiller tests in EC2 jobs \- [https://github.com/kubernetes/kubernetes/issues/126009](https://github.com/kubernetes/kubernetes/issues/126009)  
  - Due to testsuite timeout \- [https://github.com/kubernetes/kubernetes/issues/126008](https://github.com/kubernetes/kubernetes/issues/126008)  
- \[sotiris\] [https://github.com/kubernetes/kubernetes/pull/124296](https://github.com/kubernetes/kubernetes/pull/124296)   
- \[Sergey\] Device plugin failure injection tests: [https://github.com/kubernetes/kubernetes/pull/125753](https://github.com/kubernetes/kubernetes/pull/125753) 

Jul 3, 2024 \[Cancelled\]

Host: 

Agenda:

- 

Jun 26, 2024   
Recording: [https://www.youtube.com/watch?v=Cn-1k0U1kGw](https://www.youtube.com/watch?v=Cn-1k0U1kGw) 

Agenda:

- [https://kubernetes.slack.com/archives/C0BP8PW9G/p1718369837055379](https://kubernetes.slack.com/archives/C0BP8PW9G/p1718369837055379)   
  - [https://github.com/kubernetes/kubernetes/issues/122270](https://github.com/kubernetes/kubernetes/issues/122270)  
- \[fromani\] Annotating pod to detect leftovers: [https://github.com/kubernetes/kubernetes/pull/125434](https://github.com/kubernetes/kubernetes/pull/125434)   
  - Driven by: [https://github.com/kubernetes/kubernetes/pull/123468](https://github.com/kubernetes/kubernetes/pull/123468) (PTAL\!)   
- \[alex\] Test guidance compliance work   
  - [https://github.com/kubernetes/test-infra/pull/32752](https://github.com/kubernetes/test-infra/pull/32752)   
- \[Sotiris\] Can we do triage for [https://github.com/kubernetes/test-infra/pull/32765](https://github.com/kubernetes/test-infra/pull/32765)   
- 

Triage:  
[https://testgrid.k8s.io/sig-node-release-blocking\#node-kubelet-serial-containerd](https://testgrid.k8s.io/sig-node-release-blocking#node-kubelet-serial-containerd)

E2eNode Suite.\[It\] \[sig-node\] CriticalPod \[Serial\] \[Disruptive\] \[NodeFeature:CriticalPod\] when we need to admit a critical pod should add DisruptionTarget condition to the preempted pod \[NodeFeature:PodDisruptionConditions\]  
 E2eNode Suite.\[It\] \[sig-node\] CriticalPod \[Serial\] \[Disruptive\] \[NodeFeature:CriticalPod\] when we need to admit a critical pod should be able to create and delete a critical pod  
 E2eNode Suite.\[It\] \[sig-node\] MirrorPodWithGracePeriod when create a mirror pod and the container runtime is temporarily down during pod termination \[NodeConformance\] \[Serial\] \[Disruptive\] the mirror pod should terminate successfully

[https://testgrid.k8s.io/sig-node-kubelet\#kubelet-gce-e2e-swap-ubuntu-serial](https://testgrid.k8s.io/sig-node-kubelet#kubelet-gce-e2e-swap-ubuntu-serial)

[https://testgrid.k8s.io/sig-node-containerd\#pull-e2e-serial-ec2-canary](https://testgrid.k8s.io/sig-node-containerd#pull-e2e-serial-ec2-canary)

[https://testgrid.k8s.io/sig-node-containerd\#cos-cgroupv1-containerd-node-e2e-serial](https://testgrid.k8s.io/sig-node-containerd#cos-cgroupv1-containerd-node-e2e-serial)

Not working

[https://testgrid.k8s.io/sig-node-cri-o\#node-kubelet-cgroupv1-serial-crio](https://testgrid.k8s.io/sig-node-cri-o#node-kubelet-cgroupv1-serial-crio)

Not working

[https://testgrid.k8s.io/sig-node-cri-o\#pr-node-kubelet-serial-crio-cgroupv2](https://testgrid.k8s.io/sig-node-cri-o#pr-node-kubelet-serial-crio-cgroupv2)  
[https://testgrid.k8s.io/sig-node-presubmits\#pr-node-kubelet-serial-containerd](https://testgrid.k8s.io/sig-node-presubmits#pr-node-kubelet-serial-containerd) 

Should not run the NodeSwap

Jun 19, 2024 \[Cancelled for holidays\]

Jun 12, 2024 \[Cancelled for KEP freeze reviews\]  
Hosts:   
	Tests:   
Bugs triaging:   
	  
Agenda:

- 

Follow up Items:

Jun 5, 2024

Recording: [https://www.youtube.com/watch?v=L1rXfz5pJgQ](https://www.youtube.com/watch?v=L1rXfz5pJgQ) 

Hosts:  
	Tests: Anish  
	Bugs triage: Peter Hunt

Agenda:

- Release blocking?:  
  - [https://github.com/orgs/kubernetes/projects/68/views/35?sliceBy%5BcolumnId%5D=Labels\&sliceBy%5Bvalue%5D=sig%2Fnode](https://github.com/orgs/kubernetes/projects/68/views/35?sliceBy%5BcolumnId%5D=Labels&sliceBy%5Bvalue%5D=sig%2Fnode)   
  - [https://github.com/kubernetes/kubernetes/issues/125264](https://github.com/kubernetes/kubernetes/issues/125264)  
    - [https://github.com/kubernetes/kubernetes/issues/125264\#issuecomment-2148446172](https://github.com/kubernetes/kubernetes/issues/125264#issuecomment-2148446172)  
    -   
  - [https://github.com/kubernetes/kubernetes/issues/125183](https://github.com/kubernetes/kubernetes/issues/125183)

      

May 29, 2024  
Recording: [https://www.youtube.com/watch?v=PSq4VpMSlQ0](https://www.youtube.com/watch?v=PSq4VpMSlQ0)

Hosts:   
	Tests:   
Bugs triaging: 

Agenda:

- ~~\[follow-up\] [https://testgrid.k8s.io/sig-node-containerd\#ci-cgroupv1-containerd-node-arm64-e2e-serial-ec2-eks](https://testgrid.k8s.io/sig-node-containerd#ci-cgroupv1-containerd-node-arm64-e2e-serial-ec2-eks)~~   
  - ~~Filed [https://github.com/kubernetes/kubernetes/issues/125173](https://github.com/kubernetes/kubernetes/issues/125173)~~  
  - ~~Looks like swap feature was enabled in cgroupv1 jobs but it is cgroupv2 only feature?~~  
- Help to review [https://github.com/kubernetes/kubernetes/pull/124617](https://github.com/kubernetes/kubernetes/pull/124617)

Follow up:

- This should have been done: [https://testgrid.k8s.io/sig-node-release-blocking\#node-kubelet-serial-containerd](https://testgrid.k8s.io/sig-node-release-blocking#node-kubelet-serial-containerd) but still failing  
  - [https://github.com/kubernetes/kubernetes/pull/125027](https://github.com/kubernetes/kubernetes/pull/125027)   
  - [https://testgrid.k8s.io/sig-node-kubelet\#kubelet-gce-e2e-arm64-ubuntu-serial](https://testgrid.k8s.io/sig-node-kubelet#kubelet-gce-e2e-arm64-ubuntu-serial)   
  - [https://testgrid.k8s.io/sig-node-cri-o\#node-kubelet-cgrpv2-serial-crio\&width=5](https://testgrid.k8s.io/sig-node-cri-o#node-kubelet-cgrpv2-serial-crio&width=5)   
  - Peter will fix by adding more skips  
  - Maybe file a bug  
- Follow up on sidecar meeting: E2eNode Suite.\[It\] \[sig-node\] \[NodeFeature:SidecarContainers\] Containers Lifecycle should terminate sidecars simultaneously if prestop doesn't exit  
-    
-   
-  The test is broken completely: [https://testgrid.k8s.io/sig-node-containerd\#cos-cgroupv1-containerd-node-e2e-serial\&width=5](https://testgrid.k8s.io/sig-node-containerd#cos-cgroupv1-containerd-node-e2e-serial&width=5) 

- Was green with no tests, now failing with timeout: [https://testgrid.k8s.io/sig-node-cri-o\#node-kubelet-cgroupv1-serial-crio\&width=5](https://testgrid.k8s.io/sig-node-cri-o#node-kubelet-cgroupv1-serial-crio&width=5) 

May 22, 2024  
Recording: [https://www.youtube.com/watch?v=rPoI3HrTkiM](https://www.youtube.com/watch?v=rPoI3HrTkiM)

Hosts:   
	Host: Peter  
Bugs triaging: Peter  

Agenda:

- [https://github.com/kubernetes/kubernetes/pull/125027](https://github.com/kubernetes/kubernetes/pull/125027) could use approval  
- [https://github.com/kubernetes/kubernetes/issues/124743](https://github.com/kubernetes/kubernetes/issues/124743) still failing, need to bump cri-o version  
- [https://testgrid.k8s.io/sig-node-containerd\#ci-cgroupv1-containerd-node-arm64-e2e-serial-ec2-eks](https://testgrid.k8s.io/sig-node-containerd#ci-cgroupv1-containerd-node-arm64-e2e-serial-ec2-eks) Has additional failures, needs an issue

Follow up Items:

- \[Sergey\] [https://kubernetes.slack.com/archives/C0BP8PW9G/p1716308390271449](https://kubernetes.slack.com/archives/C0BP8PW9G/p1716308390271449)  
  - Looks like something we changed for sidecars  
  - Matthyx was planning to work on a fix

May 15, 2024

* **No meeting, no items on the agenda.**

May 8, 2024

Recording:[https://www.youtube.com/watch?v=ZlL0yVKJ\_o8](https://www.youtube.com/watch?v=ZlL0yVKJ_o8) 

Hosts:   
	Host: Peter  
Bugs triaging: Dixita

Agenda:

Follow up Items: 

* \[Peter\] Open an issue for [https://testgrid.k8s.io/sig-node-kubelet\#kubelet-gce-e2e-swap-fedora](https://testgrid.k8s.io/sig-node-kubelet#kubelet-gce-e2e-swap-fedora) failures  
* \[Dixita\] Memory usage beyond node allocatable tests failing again: [https://github.com/kubernetes/kubernetes/issues/120646](https://github.com/kubernetes/kubernetes/issues/120646)   
* [https://github.com/kubernetes/kubernetes/issues/124345](https://github.com/kubernetes/kubernetes/issues/124345) follow up with Swati and Francesco

May 1, 2024  
Recording: [https://www.youtube.com/watch?v=h89s\_z-YmIU](https://www.youtube.com/watch?v=h89s_z-YmIU)

Hosts:   
	Host:   
Bugs triaging: Anish

Agenda:

Follow up Items:

Apr 24, 2024  
Recording: [https://www.youtube.com/watch?v=MlltvJWa1so](https://www.youtube.com/watch?v=MlltvJWa1so)   
Hosts:  
	Host: Sergey  
	Bugs triaging: Anish

Agenda:

Apr 17, 2024  
Recording: [https://www.youtube.com/watch?v=MhMZJvLx3sg](https://www.youtube.com/watch?v=MhMZJvLx3sg)  
Hosts:  
	Host: Sergey  
	Bugs triaging: Anish

Agenda:

- \[Sotiris\] Test PR needing approval  
- [https://github.com/kubernetes/kubernetes/pull/124097](https://github.com/kubernetes/kubernetes/pull/124097)

- \[Kevin Hannon\] NVIDIA K80 out of support in May  
  - [https://github.com/kubernetes/test-infra/issues/32242](https://github.com/kubernetes/test-infra/issues/32242)  
- \[Anish\] [https://github.com/kubernetes/kubernetes/issues/116965](https://github.com/kubernetes/kubernetes/issues/116965)  
  - IIUC, pod status is not updated during graceful node shutdown. Does anyone have historical context on why the pod status is not updated?  
  - Ryan to reply on issue to explain the expected behavior part of this behavior  
  - \[Ed\] ideally we need to extend the e2e test.  
  - \[ryan\] kubelet must be killed before networking is shut down

Followup

- \[Sotiris\] Seems worth it to Improve cpu manager tests coverage, [https://github.com/kubernetes/kubernetes/issues/100145](https://github.com/kubernetes/kubernetes/issues/100145) . What do you think? How should we proceed with this?   
  


      \-     \[anishshah\] \- v1.30 release report  
	\- [github.com/AnishShah/sig-node-flaky-tests/tree/main](https://github.com/AnishShah/sig-node-flaky-tests/tree/main)  
           \-  22/249 sig-node release blocking tests are flaky.

Apr 10, 2024  
Recording: [https://www.youtube.com/watch?v=NUzCEC4WuL0](https://www.youtube.com/watch?v=NUzCEC4WuL0)   
Hosts:  
	Host: ndixita  
	Bugs triaging: Peter Hunt

Agenda:

- \[kannon92\] Test PRs needing approval  
- [https://github.com/kubernetes/kubernetes/pull/123950](https://github.com/kubernetes/kubernetes/pull/123950)  
- [https://github.com/kubernetes/kubernetes/pull/123386](https://github.com/kubernetes/kubernetes/pull/123386)  
- [https://github.com/kubernetes/test-infra/pull/32271](https://github.com/kubernetes/test-infra/pull/32271)

- Cgroup v2 crio jobs  
  - Deprecating cgroup v1 means that we should have 1on1 coverage for cgroup v1 and cgroup v2  
  - [Add corresponding cgroups v2 for node-crio-e2e-features and node-crio-flaky](https://github.com/kubernetes/test-infra/pull/32409)  
  - [crio huge pages cgroup v2](https://github.com/kubernetes/test-infra/pull/32407)  
  - [Resource managers crio cgroup v2](https://github.com/kubernetes/test-infra/pull/32406)

- \[Ed\] Can we consider triaging [SIG-Node PRs](https://github.com/orgs/kubernetes/projects/49) in this meeting?

Followup

* Check which tests need to have coverage for cgroupv2  
* Consider [Sig Node PRs](https://github.com/orgs/kubernetes/projects/49) triaging : maybe once per month?  
* [https://github.com/kubernetes/kubernetes/pull/124220](https://github.com/kubernetes/kubernetes/pull/124220)   
* Sig node : [https://github.com/kubernetes/kubernetes/pull/124229](https://github.com/kubernetes/kubernetes/pull/124229) 

Apr 3, 2024  
Recording: [https://www.youtube.com/watch?v=wZeHdf3PtMQ](https://www.youtube.com/watch?v=wZeHdf3PtMQ)

Hosts:   
	Host: skanzhelev  
Bugs triaging: anish

Agenda:

- [Bugs dashboard](https://github.com/orgs/kubernetes/projects/59)  
- \[ndixita\] Questions to get some context to help deflake the tests and cleanup  
  - [https://github.com/kubernetes/test-infra/pull/32271](https://github.com/kubernetes/test-infra/pull/32271) why did we remove manager jobs from serial tests  
    - Duplicate coverage so fine to remove  
    - Presubmit and periodics are already running these tests  
  - Kubeadm version skew tests in sig-node-kubelet POC  
    - The tests are in sig-cluster-lifecycle and sig-node. Send a PR to remove them from sig-node?  
    - [https://github.com/kubernetes/test-infra/blob/d3f9ee6f4d5b185a7b784533d6a36fab9c8409dc/config/jobs/kubernetes/sig-cluster-lifecycle/kubeadm-kinder-kubelet-x-on-y.yaml\#L356](https://github.com/kubernetes/test-infra/blob/d3f9ee6f4d5b185a7b784533d6a36fab9c8409dc/config/jobs/kubernetes/sig-cluster-lifecycle/kubeadm-kinder-kubelet-x-on-y.yaml#L356)   
  - Swap serial tests are flaky while parallel are not  
    - ideally effort needs to be put to deflake these tests  
  - History of node-kubernetes-containerd-flaky dashboard \- [https://testgrid.k8s.io/sig-node-containerd\#node-kubelet-containerd-flaky](https://testgrid.k8s.io/sig-node-containerd#node-kubelet-containerd-flaky) ?

Follow up Items:

Mar 27, 2024  
Recording: [https://www.youtube.com/watch?v=fkWV\_mqcZzs](https://www.youtube.com/watch?v=fkWV_mqcZzs)

Hosts:   
	Host:   
Bugs triaging: peter hunt

- \[SergeyKanzhelev\] [https://github.com/kubernetes/kubernetes/pull/124009\#issuecomment-2013142716](https://github.com/kubernetes/kubernetes/pull/124009#issuecomment-2013142716)  
- sig-node CI v1.30 release report  
  - [\[Flaking Test\] \[sig-node\] ☂️ node-kubelet-serial-containerd job multiple flakes🌂 · Issue \#120913](https://github.com/kubernetes/kubernetes/issues/120913#issuecomment-1996670691)  
  - These tests are flaky in these dashboards:  
    - \[sig-node-release-blocking\]\[node-kubelet-serial-containerd\]  
    - \[sig-node-kubelet\]  
    - \[sig-node-containerd\]  
    - \[sig-node-crio\]  
  - Manager’s tests \- lets remove them from Serial lane  
    - [https://github.com/kubernetes/test-infra/pull/32271](https://github.com/kubernetes/test-infra/pull/32271)   
    - Check CI jobs are working for managers

- # \[Sotiris\] oomkiller\_linux\_test: fix warnings

  - [https://github.com/kubernetes/kubernetes/pull/123908](https://github.com/kubernetes/kubernetes/pull/123908)   
  - Lets wait till branch will reopen

Mar 20, 2024

* Canceled due to Kubecon Week

Mar 13, 2024  
Recording: [https://www.youtube.com/watch?v=itj3vxg23nk](https://www.youtube.com/watch?v=itj3vxg23nk)   
Hosts:   
	Host: Dixita (Dixi)  
Bugs triaging: Anish

Agenda:

* \[Dixi\] Removing huge pages from allocatable/capacity [https://github.com/kubernetes/kubernetes/pull/119173](https://github.com/kubernetes/kubernetes/pull/119173)  
* [Bugs with no priority](https://github.com/kubernetes/kubernetes/issues?q=is%3Aissue+label%3Akind%2Fbug+is%3Aopen+label%3Asig%2Fnode+-label%3Apriority%2Fimportant-soon+-label%3Apriority%2Fimportant-longterm+-label%3Apriority%2Fbacklog+-label%3Apriority%2Fcritical-urgent+-label%3Atriage%2Fneeds-information)  
* Seeking help to debug Serial crio jobs failures  
  * [https://github.com/kubernetes/kubernetes/pull/123908](https://github.com/kubernetes/kubernetes/pull/123908) (from Sotiris)

Follow up Items:

* Talk about the [https://github.com/kubernetes/kubernetes/pull/119173](https://github.com/kubernetes/kubernetes/pull/119173) in Sig node.  
  * Why /proc/meminfo used to report capacity?   
  * Change to priority/important-soon after assessing the impact.

## 2024/03/06

Recording: [https://www.youtube.com/watch?v=wCoCEAQqMOY](https://www.youtube.com/watch?v=wCoCEAQqMOY) 

Hosts:   
	Host: Dixita  
Bugs triaging: Anish

Agenda:

- Serial Jobs Failures  
  - OOM   
    - [https://github.com/kubernetes/kubernetes/issues/123589](https://github.com/kubernetes/kubernetes/issues/123589)  
    - Jobs are OOMing due to dd oom.  
    - They also run twice  
    -   
- \[harche\] \- [https://github.com/kubernetes/kubernetes/issues/123027\#issuecomment-1971147830](https://github.com/kubernetes/kubernetes/issues/123027#issuecomment-1971147830)   
  - Not sure if this is really a bug.  
- Follow-up from last week:  
  - [https://github.com/kubernetes/kubernetes/issues/122160](https://github.com/kubernetes/kubernetes/issues/122160)   
- Triaging bugs since we close to 1.30 code freeze:  
  - Bugs with [critical-urgent priority](https://github.com/orgs/kubernetes/projects/59?card_filter_query=label%3Apriority%2Fcritical-urgent).  
  - Bugs with [important-soon priority](https://github.com/orgs/kubernetes/projects/59?card_filter_query=label%3Apriority%2Fimportant-soon).  
  - Bugs with [no priority labels and no owner](https://github.com/orgs/kubernetes/projects/59?card_filter_query=-label%3Apriority%2Fimportant-soon+-label%3Apriority%2Fimportant-longterm+-label%3Apriority%2Fcritical-urgent+-label%3Apriority%2Fbacklog+-label%3Apriority%2Fawaiting-more-evidence+no%3Aassignee).

Follow up Items:

* 

## 2024/02/28

Recording: [https://www.youtube.com/watch?v=2fqfYwYwRkk](https://www.youtube.com/watch?v=2fqfYwYwRkk)

Hosts:

- Tests: ndixita  
- Bugs: anish

Agenda:

- \[esotsal\]  
  - ndixita@ [https://github.com/kubernetes/kubernetes/issues/123313](https://github.com/kubernetes/kubernetes/issues/123313) : \[Failing test\] pull-kubernetes-local-e2e  
  - PR [https://github.com/kubernetes/test-infra/pull/32025](https://github.com/kubernetes/test-infra/pull/32025)  
- \[pehunt\] quick review request [https://github.com/kubernetes/test-infra/pull/32096](https://github.com/kubernetes/test-infra/pull/32096) 

Follow up

- ndixita@  
  - [https://testgrid.k8s.io/sig-node-kubelet\#kubelet-gce-e2e-swap-ubuntu-serial](https://testgrid.k8s.io/sig-node-kubelet#kubelet-gce-e2e-swap-ubuntu-serial)  
  - OOM killer test  
  - Reach out to Ed: [https://testgrid.k8s.io/sig-node-containerd\#e2e-cos-device-plugin-gpu](https://testgrid.k8s.io/sig-node-containerd#e2e-cos-device-plugin-gpu)   
  - [https://github.com/kubernetes/kubernetes/issues/123491](https://github.com/kubernetes/kubernetes/issues/123491)   
  - Create an issue if it doesn’t exist: [https://testgrid.k8s.io/sig-node-cri-o\#node-kubelet-serial-crio](https://testgrid.k8s.io/sig-node-cri-o#node-kubelet-serial-crio)   
  - Prioritize: Eviction tests: pid returning 0 process count issues: Find related issues  
  - David Porter: [https://github.com/kubernetes/kubernetes/pull/123369](https://github.com/kubernetes/kubernetes/pull/123369)  
  - https://github.com/kubernetes/test-infra/pull/32031 : Add the labels doc   
    - [https://github.com/kubernetes/kubernetes/pull/122927](https://github.com/kubernetes/kubernetes/pull/122927)	

## 2024/02/21

Recording: [https://www.youtube.com/watch?v=e4eRbWiIPN4](https://www.youtube.com/watch?v=e4eRbWiIPN4)

Hosts:

- Tests: ndixita  
- Bugs: ndixita

Agenda:

- \[kevin\] Few Prs to review/approve  
  - [Split Disk Presubmits](https://github.com/kubernetes/test-infra/pull/31502)  
  - [Flaky Kubelet Serial Label](https://github.com/kubernetes/test-infra/pull/32031)  
  - [CRIO Eviction cgroupv2](https://github.com/kubernetes/test-infra/pull/32006)   
- \[ndixita\]  
  - Ubuntu-test-e2e failures: dims@: WIP [https://github.com/kubernetes/kubernetes/issues/123236](https://github.com/kubernetes/kubernetes/issues/123236)  
    - PR on https://github.com/kubernetes/kubernetes/pull/123390  
  - Follow up with sig testing infra  
    [https://testgrid.k8s.io/sig-node-containerd\#cos-cgroupv1-containerd-node-e2e-serial](https://testgrid.k8s.io/sig-node-containerd#cos-cgroupv1-containerd-node-e2e-serial)  
    [https://testgrid.k8s.io/sig-node-containerd\#node-e2e-features](https://testgrid.k8s.io/sig-node-containerd#node-e2e-features)  
    @ndixita: I don’t think it’s a test-infra issue anymore. Both tests look flaky, but they’re not failing because of test-infra misconfiguration anymore. That issue seems to be fixed starting Feb 15\.  
  - Bugs follow up  
    [https://github.com/kubernetes/kubernetes/issues/122903](https://github.com/kubernetes/kubernetes/issues/122903): Do we provide support for forked repos?  
    Todd Neal: [https://github.com/kubernetes/kubernetes/issues/122902](https://github.com/kubernetes/kubernetes/issues/122902) find and assign  
- \[esotsal\]  
  - ndixita@ [https://github.com/kubernetes/kubernetes/issues/123313](https://github.com/kubernetes/kubernetes/issues/123313) : \[Failing test\] pull-kubernetes-local-e2e  
- 

## 2024/02/14

Recording: [https://www.youtube.com/watch?v=g2PgAwVXHwA](https://www.youtube.com/watch?v=g2PgAwVXHwA)

Hosts:

- Tests: ndixita  
- Bugs: ndixita

Agenda:

- \[chris\] Adding prow jobs for e2e tests with containerd v2.0  
  - (ndixita): RC already shipped, in March release  
  - Issues: can’t start using containerd straightaway   
  - New features require containerdv2.0 so better to add new test tabs and have both old and new versions running

  - ## (ndixita) Monitor Ubuntu-test-e2e failures: 

  - ## 

  - ## [https://github.com/kubernetes/kubernetes/issues/123236](https://github.com/kubernetes/kubernetes/issues/123236)

  - Ndixita Sign node release testing registry related failure: follow up with sig testing infra  
  - [https://testgrid.k8s.io/sig-node-containerd\#cos-cgroupv1-containerd-node-e2e-serial](https://testgrid.k8s.io/sig-node-containerd#cos-cgroupv1-containerd-node-e2e-serial)  
  - [https://testgrid.k8s.io/sig-node-containerd\#node-e2e-features](https://testgrid.k8s.io/sig-node-containerd#node-e2e-features)   
  - Ed Bartosh: Device plugin test [https://testgrid.k8s.io/sig-node-kubelet\#kubelet-gce-e2e-swap-fedora-serial](https://testgrid.k8s.io/sig-node-kubelet#kubelet-gce-e2e-swap-fedora-serial)  
- Ndixita [https://github.com/kubernetes/kubernetes/issues/122903](https://github.com/kubernetes/kubernetes/issues/122903)   
- [https://github.com/kubernetes/kubernetes/issues/122902](https://github.com/kubernetes/kubernetes/issues/122902) find and assign

## 2024/02/07

Recording: [https://www.youtube.com/watch?v=3sCGp\_3uU2k](https://www.youtube.com/watch?v=3sCGp_3uU2k)   
Hosts:

- Tests: ndixita  
- Bugs: ndixita

Agenda:

- \[Sergey\] Are there serial tests for e2e node? Question is from Sidecar WG meeting  
  - e2e/node: missing tests need to be added  
  - Check [https://testgrid.k8s.io/sig-node-containerd\#cos-cgroupv1-inplace-pod-resize-containerd-e2e-serial](https://testgrid.k8s.io/sig-node-containerd#cos-cgroupv1-inplace-pod-resize-containerd-e2e-serial)   
- \[Ed\] GPUDevicePlugin: which tests are targeted with this feature  
- File a bug for testgrid failure  
  - Make test as flaky and move to less important tab  
  - [https://testgrid.k8s.io/sig-node-containerd\#cos-cgroupv1-containerd-e2e](https://testgrid.k8s.io/sig-node-containerd#cos-cgroupv1-containerd-e2e)   
  - Oom killer tests failing forever  
  - [https://testgrid.k8s.io/sig-node-containerd\#cos-cgroupv1-inplace-pod-resize-containerd-e2e-serial](https://testgrid.k8s.io/sig-node-containerd#cos-cgroupv1-inplace-pod-resize-containerd-e2e-serial)   
  - [https://testgrid.k8s.io/sig-node-containerd\#cos-cgroupv2-containerd-e2e](https://testgrid.k8s.io/sig-node-containerd#cos-cgroupv2-containerd-e2e)   
  - Kevin: [https://testgrid.k8s.io/sig-node-containerd\#node-kubelet-containerd-eviction](https://testgrid.k8s.io/sig-node-containerd#node-kubelet-containerd-eviction)   
- \[ndixita\] confirm if Device plugin GA feature doesn’t have periodic jobs  
- Graceful nodes shutdown don’t work with daemonsets  
  - [https://github.com/kubernetes/kubernetes/issues/122912](https://github.com/kubernetes/kubernetes/issues/122912)  
- Sig node bugs to discuss  
  - https://github.com/kubernetes/kubernetes/issues/122905

## 2024/01/31 \[Looking for a host \- canceled if not found\]

## 2024/01/24

Recording: [https://www.youtube.com/watch?v=i92MuHisqUw](https://www.youtube.com/watch?v=i92MuHisqUw)

Hosts:

- Tests: Sergey  
- Bugs: Sergey

Agenda:

* \[Kevin\] PodReadyToStartContainers e2e test PR looking for approval  
  * [https://github.com/kubernetes/kubernetes/pull/121321](https://github.com/kubernetes/kubernetes/pull/121321)  
* \[Kevin\] Crio-cgroupv2 adding to release-informing  
  * [https://github.com/kubernetes/test-infra/pull/31650](https://github.com/kubernetes/test-infra/pull/31650)  
* \[Kevin\] ImageFs e2e tests: [https://github.com/kubernetes/kubernetes/pull/121832](https://github.com/kubernetes/kubernetes/pull/121832)  
  * Running using gcp instance (remote=True) fine  
  * CI has node failure with SoftEviction

Test grid:

- [https://testgrid.k8s.io/sig-node-containerd\#containerd-e2e-ubuntu](https://testgrid.k8s.io/sig-node-containerd#containerd-e2e-ubuntu)  
- [https://testgrid.k8s.io/sig-node-containerd\#cos-cgroupv1-containerd-e2e](https://testgrid.k8s.io/sig-node-containerd#cos-cgroupv1-containerd-e2e)  
- [https://testgrid.k8s.io/sig-node-cri-o\#node-kubelet-serial-crio](https://testgrid.k8s.io/sig-node-cri-o#node-kubelet-serial-crio)  
- [https://github.com/kubernetes/kubernetes/issues/122828](https://github.com/kubernetes/kubernetes/issues/122828)  
- 

## 2024/01/17

Recording: [https://youtu.be/A3y\_\_Ivvo1c](https://youtu.be/A3y__Ivvo1c)   
Hosts:

- Tests: tzneal  
- Bugs: peter hunt

Agenda:

* Adding CI tests for separate container runtime filesystem and split filesystem  
  * Debugging [https://github.com/kubernetes/kubernetes/pull/121832](https://github.com/kubernetes/kubernetes/pull/121832) has become quite difficult due to hard coding DiskPressure  
    * Have [https://github.com/kubernetes/test-infra/pull/31638](https://github.com/kubernetes/test-infra/pull/31638) to help (needs review/approver). Will clean up once I finish debugging  
  * Added a presubmit for split disk work	  
    * [https://github.com/kubernetes/test-infra/pull/31502](https://github.com/kubernetes/test-infra/pull/31502)  
* \[harche\] Should alpha-features blocking test skip Evented PLEG feature temporarily? [https://github.com/kubernetes/kubernetes/issues/122721\#issuecomment-1895922234](https://github.com/kubernetes/kubernetes/issues/122721#issuecomment-1895922234)	  
- Tzneal \- investigate single group OOM kill failure at [https://testgrid.k8s.io/sig-node-containerd\#cos-cgroupv1-containerd-node-e2e-serial](https://testgrid.k8s.io/sig-node-containerd#cos-cgroupv1-containerd-node-e2e-serial)   
- Tzneal \- ask sig testing how to cleanup the test grid from old removed test suites  
- Kevin  
  - [https://github.com/kubernetes/kubernetes/issues/122828\#issue-2086433864](https://github.com/kubernetes/kubernetes/issues/122828#issue-2086433864)  
  - [https://kubernetes.slack.com/archives/C0BP8PW9G/p1705506664344289](https://kubernetes.slack.com/archives/C0BP8PW9G/p1705506664344289)   
- 

## 2024/01/10 \[Cancelled\]

Agenda:

* \[swsehgal\] Looking for some help in promoting sample-device-plugin image  
  * [https://github.com/kubernetes/kubernetes/pull/118534](https://github.com/kubernetes/kubernetes/pull/118534) was merged a while back but the sample device plugin image is still not promoted.  
  * In the past, I had promoted the image ([https://github.com/kubernetes/k8s.io/pull/4862](https://github.com/kubernetes/k8s.io/pull/4862)) but since then we have transitioned to registry.k8s.io so not sure not to obtain the sha of the image corresponding to the latest version of sample-device-plugin.  
  * Has anyone promoted a test image recently?

## 2024/01/03

Recording: [https://youtu.be/nw5IhScZGEY](https://youtu.be/nw5IhScZGEY)   
Hosts:

- Tests: Sergey  
- Bugs: Sergey
