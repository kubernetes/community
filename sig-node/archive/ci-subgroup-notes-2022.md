# Kubernetes SIG-Node CI subgroup notes

## 2022/12/14

Recording: [https://www.youtube.com/watch?v=drlQWZiMj6o](https://www.youtube.com/watch?v=drlQWZiMj6o) 

- [https://github.com/kubernetes/test-infra/issues/28211](https://github.com/kubernetes/test-infra/issues/28211)  
  - \[Francesco\] Let’s also use a new label for multi-numa  
  - \[Swati\] will initiate POC and then we can clean up by adding labels and doing other optimizations

PRs:

* [Block ephemeral containers for Static Pods](https://github.com/kubernetes/kubernetes/pull/114086)  
* [More backport registry move to 1.23](https://github.com/kubernetes/kubernetes/pull/114377)

## 2022/12/07

* \[swsehgal\] Topology Manager currently does not run e2e test on K8s CI due lack of multi-NUMA systems in the CI infrastructure. This could be a blocker for its GA graduation.   
  * Planning to bring this to the main SIG Node meeting next week but was wondering if this group has any suggestions on how this can be handled?  
  * Do we have [Compute optimized](https://cloud.google.com/compute/docs/compute-optimized-machines) nodes in the CI infrastructure? C2-standard-60 (referenced [here](https://cloud.google.com/architecture/best-practices-for-using-mpi-on-compute-engine)) provides VMs with multi NUMA but don’t think we have them in our infra.  
  * Any pointers?  
    * Let’s talk to k8s infra before taking expensive machine like C2-standard-60

		\[Sergey\] Update:   
two cheapest options with Numa on GCP:  
n2-standard-32  $908.47  
skanzhelev@n2-standard-32:\~$ grep NUMA=y /boot/config-\`uname \-r\`  
lscpu | grep \-i numa  
CONFIG\_NUMA=y  
CONFIG\_X86\_64\_ACPI\_NUMA=y  
CONFIG\_ACPI\_NUMA=y  
NUMA node(s):                	2  
NUMA node0 CPU(s):           	0-7,16-23  
NUMA node1 CPU(s):           	8-15,24-31  
n2d-standard-32 $790.49  
skanzhelev@n2d-standard-32:\~$ grep NUMA=y /boot/config-\`uname \-r\`  
lscpu | grep \-i numa  
CONFIG\_NUMA=y  
CONFIG\_X86\_64\_ACPI\_NUMA=y  
CONFIG\_ACPI\_NUMA=y  
NUMA node(s):                	2  
NUMA node0 CPU(s):           	0-7,16-23  
NUMA node1 CPU(s):           	8-15,24-31

* \[SergeyKanzhelev\] cgroup mismatch between kubelet and runtime work is ongoing

## 2022/11/23 \[Canceled \- short week in US\]

## 2022/11/16

- Nothing release blocking: [https://github.com/kubernetes/kubernetes/pulls?q=is%3Apr+is%3Aopen+milestone%3Av1.26+](https://github.com/kubernetes/kubernetes/pulls?q=is%3Apr+is%3Aopen+milestone%3Av1.26+)   
  -   
- Containerd 1.5: [https://kubernetes.slack.com/archives/C0BP8PW9G/p1668511893118389](https://kubernetes.slack.com/archives/C0BP8PW9G/p1668511893118389)   
  - Mismatch of driver in kubelet and runtime is not easily discoverable  
  - Cgroupv2 will be primary test target  
- 

## 2022/11/09

- Release blocking? Code freeze yesterday  
  - [https://github.com/kubernetes/kubernetes/issues/113791](https://github.com/kubernetes/kubernetes/issues/113791) may affect pod vertical scaling tests  
  - [https://github.com/kubernetes/kubernetes/issues/113781](https://github.com/kubernetes/kubernetes/issues/113781) Francesco to take a look  
-   
- 

## 2022/11/02 \[cancelled due host unavailablity\]

## 2022/10/26

## 2022/10/19

Agenda:

- \[Sergey, Swati\] Report feedback? [https://docs.google.com/document/d/1vfqqFtN4Ke2JtB9O4wjoKvMChW2Ptizmsom1\_gRGauU/edit\#heading=h.eawmmxfxo8vq](https://docs.google.com/document/d/1vfqqFtN4Ke2JtB9O4wjoKvMChW2Ptizmsom1_gRGauU/edit#heading=h.eawmmxfxo8vq)   
- \[Brian\] [https://github.com/kubernetes/kubernetes/pull/113012](https://github.com/kubernetes/kubernetes/pull/113012)   
- \[Mike\] Removing COS jobs and testgrid: [https://github.com/kubernetes/test-infra/pull/27636](https://github.com/kubernetes/test-infra/pull/27636)

Bugs triage: 6 bugs

## 2022/10/12

Agenda:

- \[Brian\] Performance tests fix update

## 2022/10/05

Agenda:

- \[Sergey\] reports for the main sig node meeting  
  - Perma failures, new improvements,   
  - Swati Sehgal can help  
- Updated the tests tags improvements KEP: [https://github.com/kubernetes/enhancements/pull/3042](https://github.com/kubernetes/enhancements/pull/3042)   
  - Please review if you interested  
- \[Brian McQueen\] [https://github.com/kubernetes/kubernetes/issues/109295](https://github.com/kubernetes/kubernetes/issues/109295)   
- \[Sergey\] Soak test: todo add link

## 2022/09/21

Agenda:

- \[Sergey\] Triage continued  
- \[Sergey\] Dashboards names and location:  
  - [https://testgrid.k8s.io/sig-node-release-blocking](https://testgrid.k8s.io/sig-node-release-blocking) 1.22 is tested, 1.24 and 1.25 is not  
  - [https://testgrid.k8s.io/sig-node-kubelet](https://testgrid.k8s.io/sig-node-kubelet) kubelet node conformance? Features on master?

I will resurrect this: [https://github.com/kubernetes/test-infra/issues/24641](https://github.com/kubernetes/test-infra/issues/24641) 

## 2022/09/13

Agenda:

- \[Sergey\] Mostly triage

## 2022/09/07

Attendees:

- 

Agenda:

- (mmiranda96) [https://github.com/kubernetes/test-infra/pull/27406](https://github.com/kubernetes/test-infra/pull/27406)

## 2022/08/10

Attendees:

- 

Agenda:

- 

## 2022/08/3

Attendees:

- 

Agenda:

- 

## 2022/07/27 \[Cancelled due to codefreeze\]

Attendees:

Agenda:

- \[danielle\] containerd PidEviction tests are failing because cadvisor is reporting back 0 values, so pid eviction priority is random. Potentially a reason for others failing too. Hoping to get david to take a look bc I’m unfamiliar with cadvisor’s codebase.

## 2022/06/29 

Attendees:

Agenda:

- \[paco\] [https://github.com/kubernetes/kubernetes/pull/108958](https://github.com/kubernetes/kubernetes/pull/108958) PR to fix density test on pod creation, needs approval  
- \[paco\] [https://testgrid.k8s.io/sig-node-cos\#soak-cos-gce](https://testgrid.k8s.io/sig-node-cos#soak-cos-gce) keeps failing for NPD after [https://github.com/kubernetes/kubernetes/pull/109396](https://github.com/kubernetes/kubernetes/pull/109396) was merged.(The ci is not using the latest code if I understand the log correctly)

## 2022/06/22 \[Cancelled, Zoom 2FA issues\]

Attendees:

Agenda:

- 

## 2022/06/15 \[starts 15 minutes late\]

Attendees:

Agenda:

- \[Sergey\] availability announcement \- Sergey out till September for baby bonding leave.  
- Triage mostly  
- Brian to take the performance test (xmcqueen)

- [https://github.com/kubernetes/kubernetes/issues/109621](https://github.com/kubernetes/kubernetes/issues/109621)

## 

## 2022/06/08

Attendees:

Agenda:

- Triage mostly

## 

## 2022/06/01

Attendees:

Agenda:  
\[fromani\] quick status update about [https://github.com/kubernetes/kubernetes/pull/109820](https://github.com/kubernetes/kubernetes/pull/109820) \- re-enabling device plugins tests. Mostly good news\!

## 2022/05/25

Attendees:

Agenda:

- (Vaibhav) Why are EvictionHard's imagefs.available and ImageGCHighThresholdPercent the same by default

## 2022/05/04

Attendees:  
![][image1]

Agenda:

- (mmiranda96) [https://github.com/kubernetes/enhancements/pull/3042](https://github.com/kubernetes/enhancements/pull/3042)  
- (danielle) [https://github.com/kubernetes/kubernetes/issues/108028](https://github.com/kubernetes/kubernetes/issues/108028)  
  - David to take a look

## 04/27/2022

Attendees:

![][image2]

Agenda:

- [https://kubernetes.slack.com/archives/C0BP8PW9G/p1650995429992899](https://kubernetes.slack.com/archives/C0BP8PW9G/p1650995429992899) 

\[Thread to start discussion on planning reliability/maintainability improvements\]

Danielle will send something to the mailing list tomorrow.

Francesco:

- Yes, good initiative and want to help  
- Adding tests and necessary refactoring has the same reviewers bandwidth  
- Can we address that?

	Danielle:

- This is why it was brought up on main meeting as we need to raise the priority of this work.  
- PR adding tests to container manager sat open for a long time as it touched code.  
- With Derek back it may be easier

	Francesco:

- Maybe even statically allocate more time for this.

	Danielle:

- Will be as loud as needed for this to happen. Kubelet PRs will be thoroughly reviewed for the proper coverage. Improve test is prerequisite for new code.

	Some areas we are lacking coverage:

- Failure modes testing  
- Sometimes expected behavior is not specified clearly.  
    
- [https://kubernetes.slack.com/archives/C0BP8PW9G/p1651068114756659](https://kubernetes.slack.com/archives/C0BP8PW9G/p1651068114756659) 

  [@danielle](https://kubernetes.slack.com/team/U8C4ZRN83) could you please see if we can find people who can look into and fix problems in SIG-node related CI jobs (guessing at the next Node CI meeting perhaps?) There's a bunch of CRI jobs there too (cc [@mrunal](https://kubernetes.slack.com/team/U1A24MU2Z) [@sascha](https://kubernetes.slack.com/team/U53SUDBD4)

  [https://testgrid.k8s.io/sig-node-containerd\#node-kubelet-containerd-eviction](https://testgrid.k8s.io/sig-node-containerd#node-kubelet-containerd-eviction)

  	Problem: is mostly known, fixing of it is unknown.

  Problem is tests interacting with each other and interacting with things on the host. For the disk maybe use some other disk.


  Trying to fill up the whole disk, but 30Gi disk is very slow to fill. So timeout may be caused by this.


  David will leave comment on this ^^^.


  [https://testgrid.k8s.io/sig-node-containerd\#node-kubelet-containerd-performance-test](https://testgrid.k8s.io/sig-node-containerd#node-kubelet-containerd-performance-test)

  	[https://github.com/kubernetes/kubernetes/pull/109551](https://github.com/kubernetes/kubernetes/pull/109551) 

  [https://testgrid.k8s.io/sig-node-cos\#e2e-cos-flaky](https://testgrid.k8s.io/sig-node-cos#e2e-cos-flaky) 

  	

  [https://testgrid.k8s.io/sig-node-cos\#soak-cos-gce](https://testgrid.k8s.io/sig-node-cos#soak-cos-gce)

  	https://github.com/kubernetes/kubernetes/pull/109396


  [https://testgrid.k8s.io/sig-node-cri-o\#ci-crio-cgroupv1-node-e2e-alpha](https://testgrid.k8s.io/sig-node-cri-o#ci-crio-cgroupv1-node-e2e-alpha)

  [https://testgrid.k8s.io/sig-node-cri-o\#ci-crio-cgroupv1-node-e2e-eviction](https://testgrid.k8s.io/sig-node-cri-o#ci-crio-cgroupv1-node-e2e-eviction)

  [https://testgrid.k8s.io/sig-node-cri-o\#ci-crio-cgroupv1-node-e2e-flaky](https://testgrid.k8s.io/sig-node-cri-o#ci-crio-cgroupv1-node-e2e-flaky)

  	Peter will try to take a look \- after two weeks will have more time.

  	\+fromanirh will also take a look


  [https://testgrid.k8s.io/sig-node-kubelet\#kubelet-gce-e2e-swap-fedora-serial](https://testgrid.k8s.io/sig-node-kubelet#kubelet-gce-e2e-swap-fedora-serial)

  [https://testgrid.k8s.io/sig-node-kubelet\#kubelet-gce-e2e-swap-fedora-serial](https://testgrid.k8s.io/sig-node-kubelet#kubelet-gce-e2e-swap-fedora-serial)

  [https://testgrid.k8s.io/sig-node-kubelet\#kubelet-gce-e2e-swap-ubuntu-serial](https://testgrid.k8s.io/sig-node-kubelet#kubelet-gce-e2e-swap-ubuntu-serial) 

			There is an open PR for these

## 04/20/2022

Attendees:

Agenda:

- \[ehashman\] Announcements  
  - Elana taking a break during 1.25, stepping down as CI subproject lead  
  - Nominating Danielle to step up as new lead  
- \[Sergey\] [https://storage.googleapis.com/k8s-metrics/failures-latest.json](https://storage.googleapis.com/k8s-metrics/failures-latest.json)   
  - Arnaud: focus on jobs failing for more than a year  
  - Everything under 90 days is not relevant  
- \[Sergey\] [https://github.com/kubernetes-sigs/cri-tools/pull/914\#issuecomment-1101102538](https://github.com/kubernetes-sigs/cri-tools/pull/914#issuecomment-1101102538)   
- \[Sergey\] [https://github.com/kubernetes/kubernetes/pull/109472](https://github.com/kubernetes/kubernetes/pull/109472)   
- [https://github.com/kubernetes/test-infra/pull/26000](https://github.com/kubernetes/test-infra/pull/26000) 

## 04/13/2022 \[Canceled due to lack of quorum and being in test freeze\]

Please help with the release blocking: [https://github.com/kubernetes/kubernetes/issues/109082](https://github.com/kubernetes/kubernetes/issues/109082) if you have cycles\!

## 04/06/2022

Attendees:  
![][image3]

 \[arnaud\]: Switch to registry.k8s.io for cri-o:  
	[https://github.com/cri-o/cri-o/pull/5777](https://github.com/cri-o/cri-o/pull/5777)

## 03/30/2022

Attendees:  
![][image4]

- \[arnaud\] Migrate away from custom images: [https://github.com/kubernetes/k8s.io/issues/1527](https://github.com/kubernetes/k8s.io/issues/1527)  
  - Likely not used by us, but Danielle will check  
- [https://kubernetes.slack.com/archives/C7J9RP96G/p1648564612846839?thread\_ts=1648557863.010859\&cid=C7J9RP96G](https://kubernetes.slack.com/archives/C7J9RP96G/p1648564612846839?thread_ts=1648557863.010859&cid=C7J9RP96G)   
- Perf regression” ​​[http://perf-dash.k8s.io/\#/?jobname=gce-100Nodes-master\&metriccategoryname=E2E\&metricname=LoadResources\&PodName=e2e-big-minion-group%2Fkubelet\&Resource=CPU](http://perf-dash.k8s.io/#/?jobname=gce-100Nodes-master&metriccategoryname=E2E&metricname=LoadResources&PodName=e2e-big-minion-group%2Fkubelet&Resource=CPU) seems to be taken care of here: [https://kubernetes.slack.com/archives/C09QZTRH7/p1648636053781389](https://kubernetes.slack.com/archives/C09QZTRH7/p1648636053781389) 

## 03/23/2022

Attendees:  
![][image5]

\- \[arnaud\] The Great Migration to registry.k8s.io  
 	\- FYI : change containerd config:   
[https://github.com/kubernetes/test-infra/pull/25739](https://github.com/kubernetes/test-infra/pull/25739)  
[https://github.com/kubernetes/test-infra/pull/25742](https://github.com/kubernetes/test-infra/pull/25742)  
	  I’ll not be around for the meeting. If you see any failures related to this change, please revert\!  
\- \[mmiranda96\] Need review on [https://github.com/kubernetes/kubernetes/pull/108862](https://github.com/kubernetes/kubernetes/pull/108862)  
	\- Fedora’s job is passing now ([https://github.com/kubernetes/kubernetes/issues/104292\#issuecomment-1074417968](https://github.com/kubernetes/kubernetes/issues/104292#issuecomment-1074417968))

## 03/16/2022

Attendees:  
![][image6]

- \[ipochi\] Next steps on getting the lock contention flags to KubeletConfiguration [PR](https://github.com/kubernetes/kubernetes/pull/104302) merged.  
- \[arnaud\] The Great migration to registry.k8s.io  
  - [https://github.com/kubernetes/k8s.io/issues/3411](https://github.com/kubernetes/k8s.io/issues/3411)  
    - How can I test the change of default sandbox image: [https://github.com/kubernetes/kubernetes/blob/f9be590b25abf3921ffffc2a6b31e941ad9ab8fc/cmd/kubelet/app/options/container\_runtime.go\#L26](https://github.com/kubernetes/kubernetes/blob/f9be590b25abf3921ffffc2a6b31e941ad9ab8fc/cmd/kubelet/app/options/container_runtime.go#L26)  
- File a bug to investigate what changed the kubelet CPU/MEMOry

[http://perf-dash.k8s.io/\#/?jobname=gce-100Nodes-master\&metriccategoryname=E2E\&metricname=LoadResources\&PodName=e2e-big-minion-group%2Fkubelet\&Resource=CPU](http://perf-dash.k8s.io/#/?jobname=gce-100Nodes-master&metriccategoryname=E2E&metricname=LoadResources&PodName=e2e-big-minion-group%2Fkubelet&Resource=CPU)

![][image7]  
[http://perf-dash.k8s.io/\#/?jobname=gce-100Nodes-master\&metriccategoryname=E2E\&metricname=LoadResources\&PodName=e2e-big-minion-group%2Fkubelet\&Resource=memory](http://perf-dash.k8s.io/#/?jobname=gce-100Nodes-master&metriccategoryname=E2E&metricname=LoadResources&PodName=e2e-big-minion-group%2Fkubelet&Resource=memory)

![][image8]

## 03/09/2022 \[Cancelled\]

## 03/02/2022

Do we have an issue tracking this failure?  
[https://testgrid.k8s.io/sig-node-containerd\#node-kubelet-containerd-performance-test](https://testgrid.k8s.io/sig-node-containerd#node-kubelet-containerd-performance-test)  
still fails even [https://github.com/kubernetes/test-infra/pull/25385](https://github.com/kubernetes/test-infra/pull/25385) is merged  
[https://github.com/kubernetes/test-infra/issues/25430](https://github.com/kubernetes/test-infra/issues/25430) 

Infra flakes a lot, do we have a bug?  
Kubernetes Presubmits blocking  
[https://testgrid.k8s.io/presubmits-kubernetes-blocking](https://testgrid.k8s.io/presubmits-kubernetes-blocking)

Kubelet memory increase:

![][image9]

![][image10]  
2022-02-24 UTC start of the spike

## 02/23/2022

- \[matthyx\] Remaining presubmits using dockershim [kubernetes/test-infra/issues/24620](https://github.com/kubernetes/test-infra/issues/24620#issuecomment-1012938110)  
  - continue cleaning them, separate PR for gce

## 02/16/2022

- \[ipochi\] Bring back the test job for lock contention tests which was removed earlier.  
  [kubernetes/test-infra\#23243](https://github.com/kubernetes/test-infra/pull/23243).  
  Next steps on [kubernetes/kubernetes\#104334](https://github.com/kubernetes/kubernetes/pull/104334)


## 02/09/2022

- \[Sergey\] [https://github.com/kubernetes/contributor-site/pull/288](https://github.com/kubernetes/contributor-site/pull/288)   
  - Please send comments publicly or privately by EOW at the latest (Feb. 11\)  
- \[danielle\] Starting to add unit tests to various parts of pkg/kubelet/… [https://github.com/kubernetes/kubernetes/pull/108024](https://github.com/kubernetes/kubernetes/pull/108024)  
  - Reviews appreciated\!  
- \[arnaud\] just need a LGTM/Approval : [https://github.com/kubernetes/test-infra/pull/25080](https://github.com/kubernetes/test-infra/pull/25080)  
  - Done  
- **Action:** Remove sig-node-critical tab from testgrid \- only 2 jobs there, not a useful view  
- Gcloud auth login failed:  
  - Tests failing to run, e.g. [https://prow.k8s.io/view/gs/kubernetes-jenkins/logs/ci-containerd-node-e2e/1491454142533603328](https://prow.k8s.io/view/gs/kubernetes-jenkins/logs/ci-containerd-node-e2e/1491454142533603328)   
  - Seems critical-urgent

Status of CI: [https://docs.google.com/spreadsheets/d/1IwONkeXSc2SG\_EQMYGRSkfiSWNk8yWLpVhPm-LOTbGM/edit\#gid=1187923038](https://docs.google.com/spreadsheets/d/1IwONkeXSc2SG_EQMYGRSkfiSWNk8yWLpVhPm-LOTbGM/edit#gid=1187923038) 

## 02/02/2022

- \[Mike\] [https://github.com/kubernetes/kubernetes/pull/107913](https://github.com/kubernetes/kubernetes/pull/107913)  
- \[Danielle\] Is looking at eviction tests.  
  - for the disk pressure, might create a small tmpfs that's easier to fill up  
- Perf dashboard: see a small bump on runtime memory, let's check next week if it’s the same pattern when it will go down

![][image11]

## 01/26/2022

- \[Mike\]: [https://github.com/kubernetes/kubernetes/pull/107768](https://github.com/kubernetes/kubernetes/pull/107768)  
- Wouldn't start:  
  - [https://testgrid.k8s.io/sig-node-containerd\#containerd-e2e-ubuntu](https://testgrid.k8s.io/sig-node-containerd#containerd-e2e-ubuntu)   
  - [https://testgrid.k8s.io/sig-node-containerd\#containerd-node-e2e-1.4](https://testgrid.k8s.io/sig-node-containerd#containerd-node-e2e-1.4)   
  - [https://testgrid.k8s.io/sig-node-containerd\#e2e-cos-device-plugin-gpu](https://testgrid.k8s.io/sig-node-containerd#e2e-cos-device-plugin-gpu)   
  - [https://testgrid.k8s.io/sig-node-containerd\#e2e-ubuntu](https://testgrid.k8s.io/sig-node-containerd#e2e-ubuntu)   
  - Create a bug and Danielle will take a look  
  - [https://github.com/kubernetes/kubernetes/issues/107800](https://github.com/kubernetes/kubernetes/issues/107800)   
- Separate issue, partially failing:  
  - [https://testgrid.k8s.io/sig-node-containerd\#containerd-node-e2e-features-1.4](https://testgrid.k8s.io/sig-node-containerd#containerd-node-e2e-features-1.4)   
  - [https://testgrid.k8s.io/sig-node-containerd\#containerd-node-e2e-features-1.5](https://testgrid.k8s.io/sig-node-containerd#containerd-node-e2e-features-1.5)   
  - [https://testgrid.k8s.io/sig-node-containerd\#cos-cgroupv2-containerd-e2e](https://testgrid.k8s.io/sig-node-containerd#cos-cgroupv2-containerd-e2e)  
  - [https://github.com/kubernetes/kubernetes/issues/107801](https://github.com/kubernetes/kubernetes/issues/107801)

## 01/19/2022

- Arnaud: [https://github.com/kubernetes/test-infra/issues/23822](https://github.com/kubernetes/test-infra/issues/23822)   
  - Moving to bokos is blocked by: [https://github.com/kubernetes/test-infra/issues/24798](https://github.com/kubernetes/test-infra/issues/24798)   
  - Arnaud will take a look at Ssh failure for CRI-O jobs.   
- Testgrid analysis notes:  
  - CRI-O alpha tab is running conformance  
  - No Flaky tab for containerd

## 01/12/2022

- Test naming consistency  
  - Discuss [https://github.com/kubernetes/test-infra/issues/24641](https://github.com/kubernetes/test-infra/issues/24641)  
  - Don’t make an impression that we test every platform  
    - E.g. don’t include OS name whenever non-important  
    - Container runtime?   
      - \[Elana\] Container runtime is important to know, at least two container runtimes need to be tested.  
      - How do we decide which two to test?  
      - \[Sergey\] Maybe runtime included in names in sig-node tab, but not in release blocking tests  
    - Scheme of the name needs to be aligned  
      - E2e vs. node\_e2e, OS, container runtime, container runtime version, what is being tested (features, serial, unlabeled, etc.), specific feature name (hugepages, etc.), “ci-” prefix to indicate build/deploy, “pull-”, “pr-” for presubmits (do we need to merge to one prefix?)  
    - Do we need many versions of Containerd tested? If not, where will those tests run?  
    - Distinguish e2e/node and e2e\_node and if so, how to name those?  
      - \[Elana\] Use node-only for when we don't spin a full cluster?  
- Tao: propose to use [Kubernetes SIG-Node CI Testgrid Tracker](https://docs.google.com/spreadsheets/d/1IwONkeXSc2SG_EQMYGRSkfiSWNk8yWLpVhPm-LOTbGM/edit#gid=0) to track the test grid.  
  - \[Sergey\] maybe a single tab? To track continuation  
  - \[Elana\] perhaps push vs. pull to get the status  
  - \[Mike\] tab to test mapping is important also provides longer time history  
  - \[David\] also helpful to understand if anybody is working on issues.  
  - \[Tao\] it may be a short term solution unless the test grid becomes very stable.  
  - \[Elana\] engage with sig testing?  
  - \[Artyom\] who will update the sheet?   
- \[Sergey\] [https://github.com/kubernetes/kubernetes/issues/107469](https://github.com/kubernetes/kubernetes/issues/107469)  
  - GPU tests gce-device-plugin-gpu-master? Were tests removed?   
  - \[matthyx\] will open an [issue](https://github.com/kubernetes/test-infra/issues/24851) for sig-storage regarding the same problem.  
- \[mmiranda96\] [https://github.com/kubernetes/kubernetes/issues/107412](https://github.com/kubernetes/kubernetes/issues/107412)  
- \[matthyx\] [https://github.com/kubernetes/test-infra/pull/24793/files](https://github.com/kubernetes/test-infra/pull/24793/files) Was it enough?  
  - Reopen issues and see what else needs removal/migration  
- Test grid review (copied from the last meeting):  
  - Remove:  
    - [https://testgrid.k8s.io/sig-node-kubelet\#conformance-node-rhel](https://testgrid.k8s.io/sig-node-kubelet#conformance-node-rhel)  
    - [https://testgrid.k8s.io/sig-node-kubelet\#conformance-node-containerized-rhel](https://testgrid.k8s.io/sig-node-kubelet#conformance-node-containerized-rhel)  
    -  [https://github.com/kubernetes/test-infra/pull/24852](https://github.com/kubernetes/test-infra/pull/24852)   
  - Cgroupv2: [https://github.com/kubernetes/kubernetes/issues/107062](https://github.com/kubernetes/kubernetes/issues/107062): [https://testgrid.k8s.io/sig-node-containerd\#cos-cgroupv2-containerd-node-e2e-serial](https://testgrid.k8s.io/sig-node-containerd#cos-cgroupv2-containerd-node-e2e-serial)   
  - Benchmark: [https://testgrid.k8s.io/sig-node-containerd\#node-e2e-benchmark](https://testgrid.k8s.io/sig-node-containerd#node-e2e-benchmark) find owner/file bug: [https://github.com/kubernetes/kubernetes/issues/36621](https://github.com/kubernetes/kubernetes/issues/36621)  
  - Eviction: [https://testgrid.k8s.io/sig-node-containerd\#node-kubelet-containerd-eviction](https://testgrid.k8s.io/sig-node-containerd#node-kubelet-containerd-eviction)  [https://github.com/kubernetes/kubernetes/issues/107063](https://github.com/kubernetes/kubernetes/issues/107063)   
  - [https://testgrid.k8s.io/sig-node-cri-o\#node-kubelet-serial-crio](https://testgrid.k8s.io/sig-node-cri-o#node-kubelet-serial-crio)  
    - [https://github.com/kubernetes/kubernetes/issues/107343](https://github.com/kubernetes/kubernetes/issues/107343)   
  - [https://testgrid.k8s.io/sig-node-cri-o\#ci-crio-cgroupv1-node-e2e-eviction](https://testgrid.k8s.io/sig-node-cri-o#ci-crio-cgroupv1-node-e2e-eviction)  
  - Ask COS (@bsdnet):  
    - [https://testgrid.k8s.io/sig-node-cos\#soak-cos-gce](https://testgrid.k8s.io/sig-node-cos#soak-cos-gce)  
    - [https://testgrid.k8s.io/sig-node-cos\#e2e-cos-flaky](https://testgrid.k8s.io/sig-node-cos#e2e-cos-flaky)  
  - NPD: [https://testgrid.k8s.io/sig-node-node-problem-detector\#ci-npd-e2e-node](https://testgrid.k8s.io/sig-node-node-problem-detector#ci-npd-e2e-node) [https://github.com/kubernetes/kubernetes/issues/107067](https://github.com/kubernetes/kubernetes/issues/107067)    
-   
   

## 01/05/2022

- No alternative time.  
- (imran) \- Lock file contention tests. Refresh why it was added, discuss plans to move forward to completion.   
  Last comment after addressing the review feedback \- [https://github.com/kubernetes/kubernetes/pull/104334\#discussion\_r760761261](https://github.com/kubernetes/kubernetes/pull/104334#discussion_r760761261)  
- Presubmit jobs change after dockershim removal:  
  - Maybe delete all and only bring back those we need  
  - [https://github.com/kubernetes/test-infra/issues/24620](https://github.com/kubernetes/test-infra/issues/24620)   
- Should we make a goal for 1.24:  
  - migration to kubetest2: [https://github.com/kubernetes/enhancements/issues/2464\#issuecomment-1013490836](https://github.com/kubernetes/enhancements/issues/2464#issuecomment-1013490836)   
  - Migration to common pool of projects: [https://github.com/kubernetes/test-infra/issues/7769](https://github.com/kubernetes/test-infra/issues/7769)    
- Test grid review:  
  - AppArmor tests failing [https://github.com/kubernetes/kubernetes/issues/107342](https://github.com/kubernetes/kubernetes/issues/107342)  
    - [https://testgrid.k8s.io/sig-node-kubelet\#kubelet-gce-e2e-swap-ubuntu](https://testgrid.k8s.io/sig-node-kubelet#kubelet-gce-e2e-swap-ubuntu) file a bug  
    - Maybe related to ^^^ [https://testgrid.k8s.io/sig-node-containerd\#containerd-node-features](https://testgrid.k8s.io/sig-node-containerd#containerd-node-features)   
    - [https://testgrid.k8s.io/sig-node-containerd\#cos-cgroupv2-containerd-node-e2e](https://testgrid.k8s.io/sig-node-containerd#cos-cgroupv2-containerd-node-e2e)  
    - [https://testgrid.k8s.io/sig-node-containerd\#image-validation-node-features](https://testgrid.k8s.io/sig-node-containerd#image-validation-node-features)  
  - Remove:  
    - [https://testgrid.k8s.io/sig-node-kubelet\#conformance-node-rhel](https://testgrid.k8s.io/sig-node-kubelet#conformance-node-rhel)  
    - [https://testgrid.k8s.io/sig-node-kubelet\#conformance-node-containerized-rhel](https://testgrid.k8s.io/sig-node-kubelet#conformance-node-containerized-rhel)  
  - Cgroupv2: [https://github.com/kubernetes/kubernetes/issues/107062](https://github.com/kubernetes/kubernetes/issues/107062): [https://testgrid.k8s.io/sig-node-containerd\#cos-cgroupv2-containerd-node-e2e-serial](https://testgrid.k8s.io/sig-node-containerd#cos-cgroupv2-containerd-node-e2e-serial)   
  - Benchmark: [https://testgrid.k8s.io/sig-node-containerd\#node-e2e-benchmark](https://testgrid.k8s.io/sig-node-containerd#node-e2e-benchmark) find owner/file bug: [https://github.com/kubernetes/kubernetes/issues/36621](https://github.com/kubernetes/kubernetes/issues/36621)  
  - Eviction: [https://testgrid.k8s.io/sig-node-containerd\#node-kubelet-containerd-eviction](https://testgrid.k8s.io/sig-node-containerd#node-kubelet-containerd-eviction)  [https://github.com/kubernetes/kubernetes/issues/107063](https://github.com/kubernetes/kubernetes/issues/107063)   
  - [https://testgrid.k8s.io/sig-node-cri-o\#node-kubelet-serial-crio](https://testgrid.k8s.io/sig-node-cri-o#node-kubelet-serial-crio)  
    - [https://github.com/kubernetes/kubernetes/issues/107343](https://github.com/kubernetes/kubernetes/issues/107343)   
  - [https://testgrid.k8s.io/sig-node-cri-o\#ci-crio-cgroupv1-node-e2e-eviction](https://testgrid.k8s.io/sig-node-cri-o#ci-crio-cgroupv1-node-e2e-eviction)  
  - Ask COS (@bsdnet):  
    - [https://testgrid.k8s.io/sig-node-cos\#soak-cos-gce](https://testgrid.k8s.io/sig-node-cos#soak-cos-gce)  
    - [https://testgrid.k8s.io/sig-node-cos\#e2e-cos-flaky](https://testgrid.k8s.io/sig-node-cos#e2e-cos-flaky)  
  - NPD: [https://testgrid.k8s.io/sig-node-node-problem-detector\#ci-npd-e2e-node](https://testgrid.k8s.io/sig-node-node-problem-detector#ci-npd-e2e-node) [https://github.com/kubernetes/kubernetes/issues/107067](https://github.com/kubernetes/kubernetes/issues/107067)    
  -   
  - 
