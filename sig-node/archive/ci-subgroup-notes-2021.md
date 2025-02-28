# Kubernetes SIG-Node CI subgroup notes

## 12/29/2021

**Cancelled** \- year-end holiday break

## 12/22/2021

- Device plugin: [https://github.com/kubernetes/test-infra/issues/24557](https://github.com/kubernetes/test-infra/issues/24557)   
  - Put on hold to be able to repro. But can repro locally now. So can skip tests for now  
- Memory (kubelet) down again:  
  ![][image12]

## 12/15/2021

* [https://github.com/kubernetes/test-infra/issues/24618\#issuecomment-993808136](https://github.com/kubernetes/test-infra/issues/24618#issuecomment-993808136)  
  * Job to move image-config: [https://github.com/kubernetes/test-infra/blob/master/jobs/e2e\_node/swap/image-config-swap.yaml](https://github.com/kubernetes/test-infra/blob/master/jobs/e2e_node/swap/image-config-swap.yaml)  
  * Other jobs: [https://github.com/kubernetes/test-infra/blob/master/jobs/e2e\_node/containerd/containerd-release-1.5/image-config.yaml](https://github.com/kubernetes/test-infra/blob/master/jobs/e2e_node/containerd/containerd-release-1.5/image-config.yaml)  
  * \[Alukiano\] Same with hugepages. Need to provide the init that will include both. Examples:  
    * [https://github.com/kubernetes/test-infra/pull/24673](https://github.com/kubernetes/test-infra/pull/24673)  
    * Small follow-up fix [https://github.com/kubernetes/test-infra/pull/24682](https://github.com/kubernetes/test-infra/pull/24682)  
* Follow up:  
  * GCEPD: follow up to exclude from more tabs  
    * [https://github.com/kubernetes/kubernetes/issues/106720](https://github.com/kubernetes/kubernetes/issues/106720),  
    * [https://github.com/kubernetes/kubernetes/issues/106719](https://github.com/kubernetes/kubernetes/issues/106719),  
    * [https://testgrid.k8s.io/sig-node-containerd\#containerd-e2e-ubuntu](https://testgrid.k8s.io/sig-node-containerd#containerd-e2e-ubuntu)   
    * [https://testgrid.k8s.io/sig-node-containerd\#cos-cgroupv2-containerd-e2e](https://testgrid.k8s.io/sig-node-containerd#cos-cgroupv2-containerd-e2e)  
    * [https://testgrid.k8s.io/sig-node-containerd\#image-validation-cos-e2e](https://testgrid.k8s.io/sig-node-containerd#image-validation-cos-e2e)  
    * [https://testgrid.k8s.io/sig-node-containerd\#image-validation-ubuntu-e2e](https://testgrid.k8s.io/sig-node-containerd#image-validation-ubuntu-e2e)  
    * [https://testgrid.k8s.io/sig-node-containerd\#e2e-ubuntu](https://testgrid.k8s.io/sig-node-containerd#e2e-ubuntu)  
    * [https://testgrid.k8s.io/sig-node-cos\#soak-cos-gce](https://testgrid.k8s.io/sig-node-cos#soak-cos-gce)  
    * [https://testgrid.k8s.io/sig-node-cos\#e2e-cos](https://testgrid.k8s.io/sig-node-cos#e2e-cos)  
    * [https://testgrid.k8s.io/sig-node-cos\#e2e-cos-ip-alias](https://testgrid.k8s.io/sig-node-cos#e2e-cos-ip-alias)  
    * [https://testgrid.k8s.io/sig-node-cos\#e2e-cos-proto](https://testgrid.k8s.io/sig-node-cos#e2e-cos-proto)  
    * [https://testgrid.k8s.io/sig-node-cos\#e2e-cos-serial](https://testgrid.k8s.io/sig-node-cos#e2e-cos-serial)  
    * [https://testgrid.k8s.io/sig-node-cos\#e2e-cos-slow](https://testgrid.k8s.io/sig-node-cos#e2e-cos-slow)  
  * Create issue for: [https://testgrid.k8s.io/sig-node-containerd\#cos-cgroupv2-containerd-node-e2e-serial](https://testgrid.k8s.io/sig-node-containerd#cos-cgroupv2-containerd-node-e2e-serial) ([https://github.com/kubernetes/kubernetes/issues/107062](https://github.com/kubernetes/kubernetes/issues/107062))  
  * Create issue for: [https://testgrid.k8s.io/sig-node-containerd\#node-kubelet-containerd-eviction](https://testgrid.k8s.io/sig-node-containerd#node-kubelet-containerd-eviction)  [https://github.com/kubernetes/kubernetes/issues/107063](https://github.com/kubernetes/kubernetes/issues/107063)   
  * NPD: [https://testgrid.k8s.io/sig-node-node-problem-detector\#ci-npd-e2e-node](https://testgrid.k8s.io/sig-node-node-problem-detector#ci-npd-e2e-node) [https://github.com/kubernetes/kubernetes/issues/107067](https://github.com/kubernetes/kubernetes/issues/107067)    
  * Check for [https://testgrid.k8s.io/sig-node-cos\#e2e-cos-alpha-features](https://testgrid.k8s.io/sig-node-cos#e2e-cos-alpha-features)   
  * Check if we have an issue for this: [https://testgrid.k8s.io/sig-node-presubmits\#pr-node-kubelet-serial](https://testgrid.k8s.io/sig-node-presubmits#pr-node-kubelet-serial) [https://github.com/kubernetes/test-infra/issues/24557](https://github.com/kubernetes/test-infra/issues/24557) 

## 12/09/2021

- \[ehashman\] Dockershim removal \- cleanup [https://github.com/kubernetes/test-infra/issues/24592](https://github.com/kubernetes/test-infra/issues/24592)   
  - PR: [https://github.com/kubernetes/test-infra/pull/24595](https://github.com/kubernetes/test-infra/pull/24595)   
  - There are some tests I don’t want to migrate as part of this PR with special configs: e.g. CPU manager, hugepages, etc. but are failing so no point in running them  
    - Ideally, since these are serial, would like to see all tests using the same config moved under a single job  
  - Presubmits need to be done separately  
  - Will file issues for split work (presumbits)  
- \[ruiwen-zhao\] [https://github.com/kubernetes/kubernetes/issues/106895](https://github.com/kubernetes/kubernetes/issues/106895)   
  - Summary API test flaky on kubelet-gce-e2e-swap-ubuntu since the beginning of test job history   
  - Same test passes on \-fedora testgrid  
- Agreed: cancelling Thursday alternate time meeting due to lack of attendance.  
  - We will only meet Wednesdays from now on.  
  - [https://github.com/kubernetes/community/pull/6285](https://github.com/kubernetes/community/pull/6285) 

## 12/01/2021

- \[ehashman\] Milestone 1.23  
  - [https://github.com/kubernetes/kubernetes/issues?q=is%3Aopen+milestone%3Av1.23+label%3Asig%2Fnode](https://github.com/kubernetes/kubernetes/issues?q=is%3Aopen+milestone%3Av1.23+label%3Asig%2Fnode)   
- Serial lane is green\!\!\!  
- Perf (memory chart) \- let’s see if trend continues next week  
  - We looked and it looks stable since code freeze (\~11/17)  
- Next week is our alternate Thursday time  
  - If we do not see substantial attendance compared to Wednesday meeting, we will revert back to Wednesdays only going forward

## 11/24/2021

- \[ehashman\] Milestone 1.23  
  - [https://github.com/kubernetes/kubernetes/issues?q=is%3Aopen+milestone%3Av1.23+label%3Asig%2Fnode](https://github.com/kubernetes/kubernetes/issues?q=is%3Aopen+milestone%3Av1.23+label%3Asig%2Fnode)   
- [https://github.com/kubernetes/kubernetes/issues/106635](https://github.com/kubernetes/kubernetes/issues/106635)  
  - Still failing, all fixes need 1.23 backports at this point, so let’s remove these from the serial lane until we can fix them next release

Perf (memory chart) \- let’s see if trend continues next week  
![][image13]

## 11/17/2021

* \[ehashman\] Node-kubelet-serial to release-informing  
  * GPU tests started failing again, sigh  
    * Danielle has a PR that will hopefully prevent this in the future  
      * [https://github.com/kubernetes/kubernetes/pull/106348](https://github.com/kubernetes/kubernetes/pull/106348)   
    * We probably shouldn’t run with real GPUs (i.e. special hardware) in our regular serial lane  
  * Is it off docker?  
    * Containerd job currently broken \- moved to community infra but has a config issue  
    * Cut a bug for [https://testgrid.k8s.io/sig-node-release-blocking\#node-kubelet-serial-containerd](https://testgrid.k8s.io/sig-node-release-blocking#node-kubelet-serial-containerd)   
  * Do we have an equivalent job for crio?  
    * Don’t think so, we should definitely add one, add both to release-informing when sufficiently green  
      * [https://github.com/kubernetes/test-infra/issues/24451](https://github.com/kubernetes/test-infra/issues/24451)   
  *   
* \[mmiranda96\] [https://github.com/kubernetes/kubernetes/issues/106469](https://github.com/kubernetes/kubernetes/issues/106469)  
  * Probably safe now, only noticed two failures yesterday.  
* \[aditi\] [https://github.com/kubernetes/kubernetes/pull/106449](https://github.com/kubernetes/kubernetes/pull/106449)  
  * Just added the log to be sure about the reason for flake  
  * Can increase the grace period/decrease sleep time based on the finding  
* \[aditi\] should we remove 1.19 jobs?   
  * [https://testgrid.k8s.io/sig-node-kubelet\#node-kubelet-features-1.19](https://testgrid.k8s.io/sig-node-kubelet#node-kubelet-features-1.19)  
  * [https://testgrid.k8s.io/sig-node-release-blocking\#node-kubelet-1.19](https://testgrid.k8s.io/sig-node-release-blocking#node-kubelet-1.19)   
  * Let’s wait for SIG Release to remove their job first  
* \[ehashman\] refactor kubelet config validation  
  * Okay to merge during test freeze?  
    * [https://github.com/kubernetes/kubernetes/pull/105360](https://github.com/kubernetes/kubernetes/pull/105360)   
  * \+1’s from Danielle and David  
  * \+1 from Sergey. 

## 11/11/2021

* [https://github.com/kubernetes/kubernetes/issues/106204](https://github.com/kubernetes/kubernetes/issues/106204)  
  * Release blocker?  
* \[danielle\] [https://github.com/kubernetes/kubernetes/pull/106348](https://github.com/kubernetes/kubernetes/pull/106348)  
  * Lets drop the dependency on GPUs and reduce some maintenance for ourselves. Also apparently the device tests were broken in the same way GPU tests originally were (but we missed it bc of \[Flaky\])  
* \[aditi\] [https://github.com/kubernetes/kubernetes/pull/106252](https://github.com/kubernetes/kubernetes/pull/106252)  
  * Thoughts on adding credential provider to node e2e?  
* Status of node serial  
  * Let’s make it green by removing flakes out




## 10/26/2021

* \[manugupt1\] I have been working with sig-storage to speed up mount / unmount performance. See PR: [https://github.com/kubernetes/kubernetes/pull/105833/files](https://github.com/kubernetes/kubernetes/pull/105833/files). The change in this PR works only on kernel 5.6+. The ask on this PR is to write tests that will test that the behavior does not change with my tests. While thinking about this, I can think of a couple of options:   
  * Re-run all the tests through the pod-spec.  
  * Move all the unit-tests that I ran as an e2e test without pod-spec.  
* \[SergeyKanzhelev\] cos-81-lts is out of support. I suggest: [https://github.com/kubernetes/test-infra/search?q=cos-stable1](https://github.com/kubernetes/test-infra/search?q=cos-stable1) \-\> `cos-89-lts`  
  [`https://github.com/kubernetes/test-infra/search?q=cos-stable2`](https://github.com/kubernetes/test-infra/search?q=cos-stable2) `-> cos-93-lts`


## 10/20/2021

* \[fromani\]\[could be postponed if not enough time\] some tests have implicit dependencies on the node state  
  * Memory manager tests have  implicit dep on memory fragmentation (or lack thereof)  
  * ContainerRuntimeRestart \- fails for timeout, but we WANT to saturate the node with pods (that’s the whole point of the test\!)  
  * How do we keep these tests while keeping it reliable (no false negatives)?  
  * Separate lanes seem a bit excessive, any other idea?  
  * Wait for the PR to reduce amount of allocated hugepages under the test  
    * If the PR will fix flakes, remove the separated lane, otherwise remove the test from serial lane with notes why it needs a separate lane  
* \[imran\] Lock Contention Tests : Updated the job with \`NodeSpecialFeature:LockContention\` , all tests are passing, CI is green.  
  Updated the test-infra PR to include a skip value with \`--skip="\\\[Flaky\\\]|\\\[Serial\\\]"\`  
  Just need to merge and proceed with the next set of steps.  
* \[mmiranda96\] Adding new alpha features to jobs e2e-gce-alpha-features ([https://github.com/kubernetes/test-infra/issues/23642](https://github.com/kubernetes/test-infra/issues/23642))  
* \[jlebon\] Allow running e2e tests on non-GCE nodes [https://github.com/kubernetes/kubernetes/pull/105764](https://github.com/kubernetes/kubernetes/pull/105764)   
  * Danielle will take a look  
    * Fromani will have a look as well  
  * What about cloud-init?  
  * Cloud ignition is an alternative \- needs to be adapted  
  * Have a way to prepare all the binaries upfront?  
  *   
* \[mmiranda96\] Updating job ci-kubernetes-node-kubelet-eviction to use swap (for [https://github.com/kubernetes/kubernetes/issues/105023\#issuecomment-947145748](https://github.com/kubernetes/kubernetes/issues/105023#issuecomment-947145748))  
  * PR: [https://github.com/kubernetes/test-infra/pull/24064](https://github.com/kubernetes/test-infra/pull/24064)  
  * Not sure if the current swap config will work on COS.  
  * [https://github.com/kubernetes/test-infra/blob/cc714da33d7ba85672aa4c7f58e0b3993155176d/jobs/e2e\_node/swap/crio\_swap1g.ign](https://github.com/kubernetes/test-infra/blob/cc714da33d7ba85672aa4c7f58e0b3993155176d/jobs/e2e_node/swap/crio_swap1g.ign)   
* \[ehashman\] attendance for alt time  
  * Let’s hold another one in november and then if poor attendance again, cancel

## 10/14/2021

* \[mmiranda96\] Troubleshooting containerd 1.4 canaries ([https://github.com/kubernetes/test-infra/issues/23915](https://github.com/kubernetes/test-infra/issues/23915))  
  * From the logs, it appears that the cluster is never created. Node operations are no-ops. Is the cluster expected to be existent before the tests run?  
* \[mmiranda96\] MemoryPressure testing with swap enabled.  
  * How can we run tests in machines with swap enabled?  
  * [https://github.com/kubernetes/test-infra/tree/master/jobs/e2e\_node/swap](https://github.com/kubernetes/test-infra/tree/master/jobs/e2e_node/swap) 

## 10/06/2021

* \[bobbypage\] Kubetest2 migration plan  
  * Amit @amwat will provide a bit of context on kubetest2 migration plans and how it relates to node e2e testing  
  * ref: [https://github.com/kubernetes/enhancements/issues/2464](https://github.com/kubernetes/enhancements/issues/2464)

		CI jobs very first layer  Image for prow job \- has many tools on it already  
\- all these tools are deprecated and in maintenance mode. Tools evolved from bash script and became unmaintainable   
\- kubetest2 is designed to be extensible and will replace it.  
\- PLuggable on where to test (GCE, AWS, Kind, etc.)  
\- Pluggable on what to test

Thousands of jobs using old tools. The process of switching all the jobs will be slow. Some of the jobs will be moved to kubetest2. Presubmits and release blocking jobs are the first target.

Mainly: awareness of the project. Feature requests must go to kubetest2 now.

Most significant impacting change \- node tester. kubetest2 will use a makefile as a source of truth. make test\_e2e lets you run tests, but kubetest is not using it. kubetest2 will change this and will only use makefile. Dealing with test infra will be mostly when bugs are encountered, no need to deal with it any longer.

Danielle: some tests needs more features  
Amwat: yes, can be added in node tester

		Timeline?  
Amwat: scoped to presubmits and release blocking \- 1.24 is a target version. At least jobs will start be running. No timeline for other jobs.

* \[SergeyKanzhelev\] Image for presubmits: [https://github.com/kubernetes/kubernetes/issues/105381](https://github.com/kubernetes/kubernetes/issues/105381)   
  * File an issue for the future  
  * Revert to image family for now  
  * Find somebody to investigate  
* \[SergeyKanzhelev\] [https://docs.google.com/document/d/19HqSyrS-4pyubqTvQV0hJKt\_97nbCSt\_aD0soL-RGGE/edit\#heading=h.veqp9g4ihszu](https://docs.google.com/document/d/19HqSyrS-4pyubqTvQV0hJKt_97nbCSt_aD0soL-RGGE/edit#heading=h.veqp9g4ihszu)   
* \[mmiranda96\] A little off-topic, do we plan to participate in Hacktoberfest? Maintainers guide link: [https://hacktoberfest.digitalocean.com/resources/maintainers](https://hacktoberfest.digitalocean.com/resources/maintainers)  
  * Issue on kops for last year event: [https://github.com/kubernetes/kops/issues/9920](https://github.com/kubernetes/kops/issues/9920)  
  * No \- Kubernetes has explicitly opted out of Hacktoberfest project wide. The quality of contributions we’ve historically gotten have been very low and created a lot of extra cleanup work for maintainers.  
    	

## 09/29/2021

* \[SergeyKanzhelev\] NodeConformance updates

  ./\_output/local/go/bin/ginkgo \--dryRun \-v ./\_output/local/go/bin/e2e\_node.test | sed $'s,\\x1b\\\\\[\[0-9;\]\*\[a-zA-Z\],,g' \> ./tmp/e2e\_node.test.txt

  [https://github.com/kubernetes/community/blob/32a1c14d04ff78684d78b827ac7c49f70352d509/contributors/devel/sig-testing/e2e-tests.md\#kinds-of-tests](https://github.com/kubernetes/community/blob/32a1c14d04ff78684d78b827ac7c49f70352d509/contributors/devel/sig-testing/e2e-tests.md#kinds-of-tests)

* \[bobbypage\] Kubetest2 migration plan  
  * Moved to 10/06  
* \[arnaud\] Migration to k8s-infra: [https://github.com/kubernetes/k8s.io/issues/1469](https://github.com/kubernetes/k8s.io/issues/1469)  
  * Migrate away from GCP projects:  
    * k8s-jkns-pr-node-e2e: [https://cs.k8s.io/?q=k8s-jkns-pr-node-e2e\&i=nope\&files=\&excludeFiles=\&repos=](https://cs.k8s.io/?q=k8s-jkns-pr-node-e2e&i=nope&files=&excludeFiles=&repos=)  
    * cri-containerd-node-e2e: [https://cs.k8s.io/?q=cri-containerd-node-e2e\&i=nope\&files=\&excludeFiles=\&repos=](https://cs.k8s.io/?q=cri-containerd-node-e2e&i=nope&files=&excludeFiles=&repos=)  
    * K8s-cri-containerd: house of bucket gs://cri-containerd-testing : [https://cs.k8s.io/?q=cri-containerd-staging\&i=nope\&files=\&excludeFiles=\&repos=](https://cs.k8s.io/?q=cri-containerd-staging&i=nope&files=&excludeFiles=&repos=)  
  * **Action:** ehashman to create an issue for SIG Node in test-infra for the SIG to move the project with steps for a new contributor to pick up, cc Arnaud  
    * [https://github.com/kubernetes/test-infra/issues/23822](https://github.com/kubernetes/test-infra/issues/23822)   
  * Example PR: [https://github.com/kubernetes/test-infra/pull/23777/files](https://github.com/kubernetes/test-infra/pull/23777/files)   
* Why serial run on every PR?  
  * [https://github.com/kubernetes/test-infra/pull/23823](https://github.com/kubernetes/test-infra/pull/23823) 

## 09/22/2021

* \[mmiranda96\] Created [https://github.com/kubernetes/test-infra/issues/23642](https://github.com/kubernetes/test-infra/issues/23642) for keeping track of alpha feature jobs with non-alpha features.  
  * Update list of features manually to make progress on the alpha job cleanup.  
  * Work on updating tags for jobs  
* \[ehashman\] status of [https://github.com/kubernetes/k8s.io/issues/956](https://github.com/kubernetes/k8s.io/issues/956) ?  
  * GCP accounts for node contributors \- need a list of use cases  
  * Dani to drive when she returns from PTO?  
* \[Sergey\] memory spike [https://github.com/kubernetes/kubernetes/issues/105053](https://github.com/kubernetes/kubernetes/issues/105053) 

## 09/15/2021

* \[alukiano\] \- I can not attend(Israel holidays), but I want to know people opinion regarding moving DynamicKubeletConfig tests out of serial lane  
  * The feature should be deprecated under 1.23  
  * The feature tests took the most time of the serial lane, for comparison with DynamicKubeletConfig serial lane takes \~3h, without \~1h  
  * We can create a separate lane for deprecated features(I know we do not like an idea about the separate lane)  
  * Let’s prioritize removing the feature  
  * Need unified approach for kubelet configuration  
  * DynamicKubeletConfig flakes are often due to the feature being unreliable/kubelet not restarting  
  * **Action:** Elana to file an issue to detail work that needs to be done to move tests off DynamicKubeletConfig, assign Danielle and cc Sergey  
    * [https://github.com/kubernetes/kubernetes/issues/105047](https://github.com/kubernetes/kubernetes/issues/105047)   
* \[arnaud\] Migrate prowjobs to community infra  
  * In-scope :  
    * Sig-node-presubmits : [https://testgrid.k8s.io/sig-node-presubmits](https://testgrid.k8s.io/sig-node-presubmits)  
    * Sig-node-kubelet : [https://testgrid.k8s.io/sig-node-kubelet](https://testgrid.k8s.io/sig-node-kubelet)  
  * Node [https://github.com/kubernetes/k8s.io/issues/1527](https://github.com/kubernetes/k8s.io/issues/1527)  
    * [One question](https://github.com/kubernetes/k8s.io/issues/1527#issuecomment-855111265) remaining; Danielle thinks these haven’t been used in 2+ years  
    * **Action:** Danielle to comment on issue with update on Arnaud’s question and recent findings  
* \[ehashman\] status of [https://github.com/kubernetes/test-infra/issues/23	291](https://github.com/kubernetes/test-infra/issues/23291)   
  * Danielle: there are some optimizations we can do, including eviction tests, GPU tests, etc. to be less wasteful  
    * Fromani: agreed. Only few tests actually benefit from availability of GPU devices (and they need to be changed accordingly)  
  * **Action:** Sergey to follow up on combining node pool with the rest of the project  
* \[mmiranda96\] [https://github.com/kubernetes/kubernetes/issues/104556](https://github.com/kubernetes/kubernetes/issues/104556)  
  * **Action:** Elana to add a comment discussing the period/presubmit split and suggested steps forward for the job  
  * [https://testgrid.k8s.io/sig-node-kubelet\#node-kubelet-alpha](https://testgrid.k8s.io/sig-node-kubelet#node-kubelet-alpha)   
* Bugs triage: [https://github.com/orgs/kubernetes/projects/59](https://github.com/orgs/kubernetes/projects/59) 

## 09/09/2021

- \[SergeyKanzhelev\] Feedback from liggitt: [https://github.com/kubernetes/kubernetes/pull/103674\#discussion\_r704446017](https://github.com/kubernetes/kubernetes/pull/103674#discussion_r704446017) on NodeFeature-\>Feature transition  
- \[aditi\] Looking for some context on kubelet resource usage tracking tests  
  [https://github.com/kubernetes/kubernetes/issues/36621](https://github.com/kubernetes/kubernetes/issues/36621)  
  The node performance testing doc points to these tests  
  [https://github.com/kubernetes/community/blob/master/contributors/devel/sig-node/node-performance-testing.md](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-node/node-performance-testing.md)  
  Interested in refactoring the doc with current status of node perf tests and refactoring of tests as well.  
- Why tests [https://github.com/kubernetes/kubernetes/blob/master/test/e2e/node/kubelet\_perf.go](https://github.com/kubernetes/kubernetes/blob/master/test/e2e/node/kubelet_perf.go) are not run in kubelet-serial?  
- Do we need these ^^^ or [https://testgrid.k8s.io/sig-node-kubelet\#node-performance-test](https://testgrid.k8s.io/sig-node-kubelet#node-performance-test) will be enough?  
-   
- \[Imran Pochi\] Further steps on Lock contention tests ([https://github.com/kubernetes/kubernetes/pull/104334](https://github.com/kubernetes/kubernetes/pull/104334))  
  - Since test requires special config, we want a separate tab. Let’s add \[Special\] tag and make sure “features” do not pick up these “\[Special\]” tag  
- \[ehashman\] Is alternate time working?  
  - Some folks definitely can’t attend the Wednesday time  
  - Many people don’t realize this meeting exists  
  - **Action:** ehashman to set up reminder for the next Thursday meeting  
    - Done \- for the rest of 2021

## 09/01/2021

- \[mmiranda96\] [kubernetes/test-infra\#23202](https://github.com/kubernetes/test-infra/issues/23202) Job fails while building the image, mostly related to CGO\_ENABLED. Any recommendations for this?  
  - \[danielle\] FYI \- Node Problem Detector usually builds on top of Debian \- [https://github.com/kubernetes/node-problem-detector/blob/master/Makefile\#L77](https://github.com/kubernetes/node-problem-detector/blob/master/Makefile#L77)   
- ~~\[mmiranda96\] Requesting review for [kubernetes/test-infra\#23215](https://github.com/kubernetes/test-infra/pull/23215)~~  
- \[ehashman\] Test failure emails to main mailing list  
  - Seems like an oversight [https://groups.google.com/g/kubernetes-sig-node](https://groups.google.com/g/kubernetes-sig-node)   
  - AI: create a task in github  
- \[danielle\] [https://github.com/kubernetes/kubernetes/pull/104304](https://github.com/kubernetes/kubernetes/pull/104304) \- Fixing eviction tests, needs reviews please :)  
- \[rphillips\] [https://github.com/kubernetes/kubernetes/issues/104648](https://github.com/kubernetes/kubernetes/issues/104648)  [PR\#104712](https://github.com/kubernetes/kubernetes/pull/104712)  
  - E2e is coming as separate PR later to make release this week available

## 08/25/2021 Cancelled due to hosts unavailability

## 08/18/2021

- \[mmiranda96\] Requesting review for [kubernetes/node-problem-detector\#607](https://github.com/kubernetes/node-problem-detector/pull/607).  
  - AI: create issue to review other NPD tests.  
- \[mmiranda96\] Failures in [pull-kubernetes-node-e2e-alpha](https://prow.k8s.io/job-history/gs/kubernetes-jenkins/pr-logs/directory/pull-kubernetes-node-e2e-alpha)  
  - Mike: investigate and find the test grid. \+  create an issue.  
-   
- \[SergeyKanzhelev\] Quota ([https://github.com/kubernetes/test-infra/issues/23232](https://github.com/kubernetes/test-infra/issues/23232)): [https://monitoring.prow.k8s.io/d/wSrfvNxWz/boskos-resource-usage?panelId=9\&fullscreen\&orgId=1\&from=now-90d\&to=now](https://monitoring.prow.k8s.io/d/wSrfvNxWz/boskos-resource-usage?panelId=9&fullscreen&orgId=1&from=now-90d&to=now)   
  - 18 is the number of projects used.  
  - The best way to improve utilization is less test jobs or faster tests inside the job  
- \[alukiano\] Orphans jobs  
  - Runs serial and conformance tests  
  - Prepare a PR to skip serial and conformance tests [https://github.com/kubernetes/test-infra/pull/23295](https://github.com/kubernetes/test-infra/pull/23295)

## 08/12/2021

-   
- ~~\[mkmir\] Working on [kubernetes/test-infra\#23131](https://github.com/kubernetes/test-infra/issues/23131), but I can’t seem to find a way to locally run a ProwJob without connecting to GCE. Any suggestions?~~   
- Lock contention tests failing on features-master, reverted [https://github.com/kubernetes/kubernetes/issues/104307](https://github.com/kubernetes/kubernetes/issues/104307)   
  - Probably need to label this \[Serial\], remove the separate job  
  - Test labels: [https://github.com/kubernetes/community/blob/master/contributors/devel/sig-testing/e2e-tests.md\#kinds-of-tests](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-testing/e2e-tests.md#kinds-of-tests)   
  - AI: Imran:   
    - Add \[Serial\] tag   
    - Validate that test is passing in Serial [https://testgrid.k8s.io/sig-node-kubelet\#pr-node-kubelet-serial](https://testgrid.k8s.io/sig-node-kubelet#pr-node-kubelet-serial)   
    - PR that adds a job needs to be reverted  
- \[ehashman\] testgrid cleanup follow-up  
  - Didn’t get time to submit a PR, good thing for a new contributor to work/pair on  
  - We want to move “pr-\*” jobs out of the sig-node-kubelet tab to a new sig-node-presubmits tab  
  - Maybe look at any other tabs that could be consolidated (e.g. sig-node-containerd and sig-node-containerd-io)  
  - Remove release-blocking jobs   
  - Any volunteers?  
    - Imran may reach out  
    - Aditi \+ matthyx to volunteer  
    - **Action:** Elana to create an issue  
    - [https://github.com/kubernetes/test-infra/issues/23231](https://github.com/kubernetes/test-infra/issues/23231)   
- \[ehashman\] test-infra resource utilization  
  - [https://monitoring.prow.k8s.io/d/wSrfvNxWz/boskos-resource-usage?panelId=9\&fullscreen\&orgId=1\&from=now-90d\&to=now](https://monitoring.prow.k8s.io/d/wSrfvNxWz/boskos-resource-usage?panelId=9&fullscreen&orgId=1&from=now-90d&to=now)   
  - Do we even need single-test jobs like [https://testgrid.k8s.io/sig-node-kubelet\#kubelet-gce-e2e-lock-contention](https://testgrid.k8s.io/sig-node-kubelet#kubelet-gce-e2e-lock-contention) ?  
  - Suggested action: someone should audit all sig-node periodics and suggest jobs that can be consolidated/removed  
    - AI: Sergey to find out what this quota means  
    - Action: Elana to create issue to track  
    - [https://github.com/kubernetes/test-infra/issues/23232](https://github.com/kubernetes/test-infra/issues/23232)   
- CI signal:  
  - [https://github.com/kubernetes/kubernetes/issues/104173](https://github.com/kubernetes/kubernetes/issues/104173)   
  - [https://github.com/kubernetes/kubernetes/issues/99437](https://github.com/kubernetes/kubernetes/issues/99437)   
  - [https://testgrid.k8s.io/sig-node-kubelet\#kubelet-serial-gce-e2e-graceful-node-shutdown](https://testgrid.k8s.io/sig-node-kubelet#kubelet-serial-gce-e2e-graceful-node-shutdown) Sergey to create an issue ([https://github.com/kubernetes/kubernetes/issues/104344](https://github.com/kubernetes/kubernetes/issues/104344))  
- \[ehashman\] New bug triage board: [https://github.com/orgs/kubernetes/projects/59](https://github.com/orgs/kubernetes/projects/59)   
  - Hoping to triage and assign new bugs regularly as part of weekly triage meeting

## 08/04/2021

- \[dims\] Missing 1.5 branch containerd canaries [https://kubernetes.slack.com/archives/C0BP8PW9G/p1628039263053300](https://kubernetes.slack.com/archives/C0BP8PW9G/p1628039263053300)   
- \[ehashman\] cleaning up the presubmits in the kubelet-serial dash  
  - Pr- tests \- move to separate dashboard  
  - Remove release blocking ones (previous releases) from kubelet  
- \[ehashman\] why is [https://github.com/kubernetes/test-infra/blob/master/config/jobs/kubernetes/sig-node/sig-node-config.yaml](https://github.com/kubernetes/test-infra/blob/master/config/jobs/kubernetes/sig-node/sig-node-config.yaml) separate?  
- \[mmiranda96\] Following [kubernetes/kubernetes\#94289](https://github.com/kubernetes/kubernetes/issues/94289), I need to fix some more test tags in kubernetes/kubernetes (I’ll create a PR similar to [this](https://github.com/kubernetes/kubernetes/pull/103674)). Anything to consider for backport?  
- \[ehashman\] [https://github.com/kubernetes/kubernetes/pull/103257](https://github.com/kubernetes/kubernetes/pull/103257) test changes  
- \[Thomas and Qiutong\] [https://github.com/kubernetes/kubernetes/issues/100467](https://github.com/kubernetes/kubernetes/issues/100467) The fix doesn’twork.   
  - [https://github.com/kubernetes/kubernetes/issues/93338](https://github.com/kubernetes/kubernetes/issues/93338)   
  - Action: Qiutong to investigate a reproducer test  
- \[ehashman\] reenabling flaky tests [https://github.com/kubernetes/test-infra/pull/19352](https://github.com/kubernetes/test-infra/pull/19352)   
  - We are not missing a lot of signal (very few are flaky)  
  - Manu will take it  
- Serial tests updates?  
  - Couple PRs are out \- de-flaking tests that were temporary marked as such  
    - AI on fromanirh to rebase and remove the “Flake” label   
  - [https://github.com/kubernetes/kubernetes/pull/103408](https://github.com/kubernetes/kubernetes/pull/103408)  
  - [https://github.com/kubernetes/kubernetes/pull/103297](https://github.com/kubernetes/kubernetes/pull/103297)   
  - Not running to completion  [https://github.com/kubernetes/kubernetes/issues/104038](https://github.com/kubernetes/kubernetes/issues/104038)   
  -   
  - We really shouldn’t call klog.Fatalf in the tests, it’s causing the massive stack traces

## 07/28/2021

- \[mmiranda96\] Should I backport [this](https://github.com/kubernetes/kubernetes/pull/103827)?  
- \[ehashman\] 1.22 burndown  
  - [https://github.com/kubernetes/kubernetes/issues?q=is%3Aopen+milestone%3Av1.22+label%3Asig%2Fnode+](https://github.com/kubernetes/kubernetes/issues?q=is%3Aopen+milestone%3Av1.22+label%3Asig%2Fnode+)   
  - **Action:** mark known failures in serial tests as Flaky so we can try to get the job green  
    - [https://github.com/kubernetes/kubernetes/pull/103982](https://github.com/kubernetes/kubernetes/pull/103982)   
- \[alukiano\] all managers are broken by PLEG refactoring. 1.22 will have issues with high-performance workload \- pinning wouldn’t work. Urgent, but not release blocking  
  - Need to talk to release team to add a known issue to the release notes

## 07/21/2021

- \[haircommander\] Adding presubmit/release-blocking CRI-O jobs  
  - [https://hackmd.io/kEd86GlmTD-BloMBwVXQHQ](https://hackmd.io/kEd86GlmTD-BloMBwVXQHQ)   
- \[fromani\] using device plugins in the e2e suites running on CI  
  - The k8s e2e test suite has a fair amount of tests which need device plugin, because they exercise device manager \-directly or indirectly.  
    We mostly use SRIOV devices, because SRIOV devices are just the cheapest and easiest supported device to get, so this is why we wrote the tests in k8s to consume them.  
    But  we don't have device plugin support on CI. We do have gpus-enabled machine, but it's a subset and should be used sparingly (e.g. not every PR should use them. Or can we just use gpus every time? I expect no for cost reasons, but worth mentioning).  
    So today a large amount of tests just skip on CI. This is especially evident in the serial lane and in the resource management area  
    In RH we actually have machines with SRIOV devices which run the e2e testsuite, but this is of course suboptimal for a number of reasons; a much better state for everyone would be to actually have some device plugins in u/s CI.  
    There are some options we can discuss as community:  
    - 1\. use sample plugin  
    - 2\. fake sriov devices (I can elaborate on this if there is interest  
    - 3\. just use GPUs?  
    - 4\. Just bump the spec of the CI machines to have SRIOV devices?

- \[mkmir\] [E2E sysctl test](https://github.com/kubernetes/kubernetes/blob/master/test/e2e/common/node/sysctl.go) is marked as conformance. However, it does not respect some of the [requirements](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/conformance-tests.md#conformance-test-requirements):  
  - it tests only GA, non-optional features or APIs (uses feature sysctl)  
  - it works for all providers (doesn’t work for Windows and other non-sysctl OS)  
  - [https://github.com/kubernetes/kubernetes/pull/101190](https://github.com/kubernetes/kubernetes/pull/101190)   
- \[ehashman\] 1.22 burndown (includes some of the topics above)  
  - [https://github.com/kubernetes/kubernetes/issues?q=is%3Aopen+milestone%3Av1.22+label%3Asig%2Fnode+](https://github.com/kubernetes/kubernetes/issues?q=is%3Aopen+milestone%3Av1.22+label%3Asig%2Fnode+)   
- \[SergeyKanzhelev\] 1.21 vs 1.20 perf degradation: [https://github.com/kubernetes/kubernetes/issues/101989](https://github.com/kubernetes/kubernetes/issues/101989)

## 07/14/2021

- \[Sergey\] NodeConformance writeup [https://docs.google.com/document/d/1ezJPfItuhZvwyP\_RtiWTNcjCM3gi94vu1nw6uVNHKgM/edit?usp=sharing](https://docs.google.com/document/d/1ezJPfItuhZvwyP_RtiWTNcjCM3gi94vu1nw6uVNHKgM/edit?usp=sharing)   
  - NodeConformance historically tried to be two things:  
    - A set of e2e tests that you just needed a single node to run  
    - Conformance-like test for nodes  
  - \[ehashman\] Suggestion: we get rid of “NodeConformance” because the name is confusing  
    - For CRI conformance-like tests, label them CRIValidation \-- NodeAgnostic?  
    - For the set of e2e tests you just need a single node to run, let’s come up with a new name \-- needs a name (anything in test/e2e\_node) (SingleNodeTest or KubeletLocal)  
    - Note: some tests may overlap between both sets  
  - **Action:** add a plan to cover splitting the use cases for current tests  
  - **Action:** send out NodeConformance plan, soliciting feedback  
- \[ehashman\] 1.22 burndown  
- \[fromani\]\[status update\]\[serial lane\] looks like [https://github.com/kubernetes/kubernetes/issues/100145](https://github.com/kubernetes/kubernetes/issues/100145) eventually broke, will look on it ASAP

## 07/08/2021

- \[ehashman\] Report on NodeConformance from SIG Arch discussion on 07/01  
  - [SIG-Architecture Agenda and Meeting Notes](https://docs.google.com/document/d/1BlmHq5uPyBUDlppYqAAzslVbAO8hilgjqZUTaNXUhKM/edit#bookmark=id.ln4uxm9twb2r)   
  - Historical issue [https://github.com/kubernetes/kubernetes/issues/59001](https://github.com/kubernetes/kubernetes/issues/59001)   
    - Suggestion: Remove “Conformance” from the name  
    - Not just CRI.  
    - Conformance requires entire cluster. NodeConformance just require kubelet \- need to discuss what we want in scope of NodeConformance tests.  
    - NodeConformance run as presubmits  
- \[ehashman\] remaining burndown for CI signal (4 tests tracked by RT)  
  - [https://groups.google.com/g/kubernetes-dev/c/u1LMXHcKhbg/m/Lp81VX7eAgAJ](https://groups.google.com/g/kubernetes-dev/c/u1LMXHcKhbg/m/Lp81VX7eAgAJ)   
  - \#100788 [https://github.com/kubernetes/kubernetes/issues/100788](https://github.com/kubernetes/kubernetes/issues/100788)  \[sig-node\]\[NodeConformance\]  when querying /stats/summary should report resource usage through the stats api  
  - \#99437 [https://github.com/kubernetes/kubernetes/issues/99437](https://github.com/kubernetes/kubernetes/issues/99437) \[Flake\]\[sig-node\] Pods should run through the lifecycle of Pods and PodStatus  
  - \#75355 [https://github.com/kubernetes/kubernetes/issues/75355](https://github.com/kubernetes/kubernetes/issues/75355) \[Flaky test\] \[[k8s.io](http://k8s.io)\] Pods should support pod readiness gates \[NodeFeature:PodReadinessGate\]  
  - \#99979 [https://github.com/kubernetes/kubernetes/issues/99979](https://github.com/kubernetes/kubernetes/issues/99979) \[flaky test\]: \[sig-node\] Probing container should be ready immediately after startupProbe succeeds  
- \[fromani\] cannot join (conflict \- reach me on slack if needed) \- status ony:  
  - same status as next week (please review the pending fixes to the serial lane\! :) )  
    - Critical pod test PR: [https://github.com/kubernetes/kubernetes/pull/103408](https://github.com/kubernetes/kubernetes/pull/103408)   
  - Need to investigate the DynamicConfig failures we experienced recently  
  - Acknowledge the podreadinessgate flake, but still dunno how to reproduce, suggestions welcome\!  
- \[alukiano\] A lot of serial jobs broke because of DynamicConfig, the fix already proposed [https://github.com/kubernetes/kubernetes/pull/103580](https://github.com/kubernetes/kubernetes/pull/103580)(needs approval)

## 06/30/2021

- \[fromani\] update only, no agenda item:  
  - Managed to find time to fix the serial lane:  
    - Recent serial lane run: [https://prow.k8s.io/view/gs/kubernetes-jenkins/pr-logs/pull/103297/pull-kubernetes-node-kubelet-serial/1410231408982495232/](https://prow.k8s.io/view/gs/kubernetes-jenkins/pr-logs/pull/103297/pull-kubernetes-node-kubelet-serial/1410231408982495232/)   
    - Merged: [https://github.com/kubernetes/kubernetes/pull/103265](https://github.com/kubernetes/kubernetes/pull/103265)   
    - In review:  
      - [https://github.com/kubernetes/kubernetes/pull/103297](https://github.com/kubernetes/kubernetes/pull/103297)   
      - [https://github.com/kubernetes/kubernetes/pull/103408](https://github.com/kubernetes/kubernetes/pull/103408)   
  - Next up: eviction test lane (no ETA atm)  
  - More async work to be done here  
- Bug scrub follow-up  
  - Added some new bugs to the board after having scrubbed bugs, including some issues for adding test coverage  
  - Issues are now in a more manageable state, but we have so many  
  - Suggestion: bug board \+ everything else board will be optimal, need to figure out the columns (triage/waiting/accepted/in progress/done?)  
  - Hopefully moving forward we can do regular incoming issue triage as part of these meetings  
- Bot to help with automation for boards?  
  - GitHub still doesn’t have support  
  - Contribex is working on it  
- NodeConformance status  
  - Assigned to Sergey, worked on bug scrub so hasn’t had a chance to look since last meeting  
- NodeFeature status?  
  - mmiranda has submitted PR: [https://github.com/kubernetes/test-infra/pull/22677](https://github.com/kubernetes/test-infra/pull/22677)   
  - Starting with duplicating selectors in test-infra, then we can start making test changes  
- \[Sergey\] Soak tests  
  - [https://github.com/kubernetes/kubernetes/issues/64523](https://github.com/kubernetes/kubernetes/issues/64523)   
  - Is there logic we can reuse to automatically detect this?  
  - Not afaik \- usually determined by querying debugging endpoint and looking at the memory dumps  
- Resource utilization regressions now being tracked in perf-tests: [https://github.com/kubernetes/perf-tests/issues/1789](https://github.com/kubernetes/perf-tests/issues/1789)   
- How to find reviewers for various PRs?  
  - **Action:** Swati to add item to next week’s SIG Node meeting to discuss with wider SIG

## 06/23/2021

- Discuss removing NodeFeature flags in favour of Feature: [https://github.com/kubernetes/kubernetes/issues/94289](https://github.com/kubernetes/kubernetes/issues/94289)   
  - Mike Miranda to drive  
  - Need to update labels in both test-infra and k/k  
- What is the difference between NodeConformance and Conformance?  
  - Filed bug for documentation: 	[https://github.com/kubernetes/community/issues/5859](https://github.com/kubernetes/community/issues/5859)   
- Serial tests  
  - No progress  
- \[ehashman\] Volunteers for bug scrub in a week\!  
  - NASA needs more mentors: [https://docs.google.com/spreadsheets/d/1y6HKIsThphzpaG2a-Vgsc66b36ckM\_TIPqGJ6GeVUGI/edit\#gid=0](https://docs.google.com/spreadsheets/d/1y6HKIsThphzpaG2a-Vgsc66b36ckM_TIPqGJ6GeVUGI/edit#gid=0) 

Attendees:  
![][image14]

# 06/16/2021

Attendees:  
![][image15]

- Discuss removing NodeFeature flags in favour of Feature: [https://github.com/kubernetes/kubernetes/issues/94289](https://github.com/kubernetes/kubernetes/issues/94289)   
-   
- What is the difference between NodeConformance and Conformance?  
  - See also [https://kubernetes.slack.com/archives/C78F00H99/p1623803503028500](https://kubernetes.slack.com/archives/C78F00H99/p1623803503028500)   
- \[ehashman\] Volunteers for bug scrub in a week\!  
- Serial tests  
  - Eviction tests were holding everything. They are split away and needs attention. Odin has a PR ready \- increase open files limit  
  - Some serial tests are failing  
- AI: Francesco: follow up from serial tests failure investigation

# 06/10/2021

Attendees:

- \[Sergey\] I cannot join this time, but this is one of conflict  
- 

Agenda:

- \[pacoxu\]   
  - [x] ~~Ci-kubernetes-node-kubelet-conformance keeps failing: [https://github.com/kubernetes/kubernetes/issues/97130](https://github.com/kubernetes/kubernetes/issues/97130) （shoud be fixed by  [\[sig-node\] Remove failing, unused NodeConformance job kubernetes/test-infra\#22454](https://github.com/kubernetes/test-infra/pull/22454) ）~~  
- Serial tests update [https://github.com/kubernetes/kubernetes/issues/102148](https://github.com/kubernetes/kubernetes/issues/102148)   
  - \[fromani\] One huge serial test is too much, thinks it makes sense to split  
  - \[paco\] I will help in this and work on the timeout tests(eviction?)  
  - Filed [https://github.com/kubernetes/kubernetes/issues/102782](https://github.com/kubernetes/kubernetes/issues/102782) to track eviction  
- Failing Node Conformance tests?  
  - New issue filed as [https://github.com/kubernetes/test-infra/issues/22250](https://github.com/kubernetes/test-infra/issues/22250) \- possible duplicate of the existing 2 issues  
  - PR up to remove the tests at [https://github.com/kubernetes/test-infra/pull/22454](https://github.com/kubernetes/test-infra/pull/22454)   
  - Swati to further investigate

# 06/02/2021

Agenda:

- \[fromani \- cannot attend the mtg \- just status update\] “Serial” updates  
  - Kubelet \+ test logs uploaded [https://github.com/fromanirh/k8smisc/tree/main/e2e\_node](https://github.com/fromanirh/k8smisc/tree/main/e2e_node) (lacking a better place; suggestions?)  
  - Spending time analyzing the logs, will update about findings (and send PRs :) )  
  - Everyone is welcome to ping me anytime on slack to talk about this (/run more tests/upload logs)  
      
- APAC friendly time: [https://doodle.com/poll/sfeh699qt44mrzv6](https://doodle.com/poll/sfeh699qt44mrzv6) Cannot see responses...

# 05/25/2021

Agenda:

- \[fromani\] updates about “Serial” [https://testgrid.k8s.io/sig-node-kubelet\#node-kubelet-serial](https://testgrid.k8s.io/sig-node-kubelet#node-kubelet-serial)  
  - Got cpu time to run test locally and collect logs  
  - First experimental run done (failed with timeout)  
  - Doing runs with the same parameter the node-kubelet-serial lane is using  
  - Will collect and publish logs on [https://github.com/fromanirh/k8smisc/tree/main/e2e\_node](https://github.com/fromanirh/k8smisc/tree/main/e2e_node) (lacking a better place; suggestions?)  
  - Will deep dive in the logs after published \- so we can go in parallel  
  - Ping me on k8s/cncf slack chans (@fromani)  for any question/comment/chat about the issue  
    - Ping me or file issues against the repo above so we don’t miss/forget  
  - PRs potentially helping with them:  
    - File handles limit increase [https://github.com/kubernetes/kubernetes/pull/102169](https://github.com/kubernetes/kubernetes/pull/102169)   
    - Fix that allows to upload artifacts [https://github.com/kubernetes/kubernetes/pull/102209](https://github.com/kubernetes/kubernetes/pull/102209)   
  - More logs are coming (kubelet.log)  
- Add ability to run tests on PR  
  - [https://github.com/kubernetes/test-infra/pull/22284](https://github.com/kubernetes/test-infra/pull/22284) 

# 05/19/2021

Attendees:  
![][image16]

Agenda:

- Meeting time for asia?  
- Code coverage dashboard in OSS: [https://testgrid.k8s.io/sig-testing-canaries\#ci-kubernetes-coverage-unit\&include-filter-by-regex=pkg%2Fkubelet](https://testgrid.k8s.io/sig-testing-canaries#ci-kubernetes-coverage-unit&include-filter-by-regex=pkg%2Fkubelet)   
  - Examples:  
    - pkg/kubelet/container/runtime.go  
    - pkg/kubelet/certificate/bootstrap/bootstrap.go  
  - \[Elana\] Some files don’t have targeted unit tests  
  - \[Matthias\] Some files very hard to write tests for  
  - \[Elana\] Rearchitexting some features will be needed  
  - \[Francesco\] It is a backlog items. How we will prioritize?  
- Discuss “Serial” [https://testgrid.k8s.io/sig-node-kubelet\#node-kubelet-serial](https://testgrid.k8s.io/sig-node-kubelet#node-kubelet-serial)  
  - Timeouts under tests may need adjustment  
  - We already adjusted global timeout  
  - Just run locally with higher timeout and see what is failing  
  - **Action** \[Artyom\] file issue [https://github.com/kubernetes/kubernetes/issues/102148](https://github.com/kubernetes/kubernetes/issues/102148)   
  - Francesco \- best effort \- can run tests locally  
  - Aditi, Mike, Matthias to help  
- \[Aditi\] Flags for kubelet? Or kubelet config?  
  - Policy is no new flags, add to KubeletConfig  
  - Raise at full SIG Node

# 05/12/2021

Agenda:

- Triage: [https://github.com/orgs/kubernetes/projects/43?card\_filter\_query=no%3Aassignee](https://github.com/orgs/kubernetes/projects/43?card_filter_query=no%3Aassignee)  
    
- \[Artyom\]Need more good first issues   
  - Writing e2e test may be a good (difficult) first issue  
  - \[Francesco\] please not first good issues and we can review them together later  
  - \[A\] increasing code coverage is a good issue  
- Still failing:  
  - Node conformance (docker)  
    - \[S\] find an issue and include on the project board.  
  - Serial jobs are still failing  
    - \[A\] when we increased memory \- it is timing out. Pod is killed so no logs  
    - Artyom to open a new issue  
  - Orphans clean up  
    - [https://github.com/kubernetes/kubernetes/issues/98265](https://github.com/kubernetes/kubernetes/issues/98265)   
- 

# 05/04/2021 Cancel for KubeCon

# 04/28/2021 Short sync up

Attendees:  
![][image17]

Agenda:

- Follow ups \- need to move to the next week:  
  - Kubetest2 updates David  
  - [NodeConformance](https://testgrid.k8s.io/sig-node-critical#kubelet-NodeConformance) Qiutong  
  - Serial clean up \- Francesco, actually there is a PR: [https://github.com/kubernetes/test-infra/pull/21828](https://github.com/kubernetes/test-infra/pull/21828)   
- \[fromani\] Managing kubelet state in e2e tests: overriding the system default (/var/lib/kubelet)  
  - Shared global state between tests  
  - Any objections to move away from it?  
  - Storing state of the kubelet  
  - Ideally each e2e test has it’s own state  
- \[ehashman\] looking for volunteers for [https://github.com/kubernetes/perf-tests/issues/1789](https://github.com/kubernetes/perf-tests/issues/1789)   
  - Currently don’t have any perf/scalability tests for upstream kubelet resource utilization (CPU/memory)

# 04/21/2021

Attendees:

- Sergey  
- Elana  
- Artyom

Agenda:

- triage

# 04/14/2021

Attendees:

- fromani  
- ehashman  
- [Sergey Kanzhelev](mailto:skanzhelev@google.com)  
- David Porter  
- Amim Knabben  
- Qiutong Song  
- Madhav Jivrajani  
- Jiaming Xu

Agenda:

- Kubetest2 updates: [https://github.com/kubernetes-sigs/kubetest2/pull/103/](https://github.com/kubernetes-sigs/kubetest2/pull/103/)  
  - David: currently using kubetest1, moving to kubetest2. Feedback: kubetest has some flags that do not match “make”-way to run tests. PR is moving to the “make”-way to run tests. David’s concern: there will be some need to migrate all tests, but do not see big value in doing it.  
  - Elana: it will be useful to have a writeup with the justification.  
  - Francesco: want to take a look and learn more as a consumer of tests.  
    - Q: is this part of a process to make the node e2e test more like the other e2e tests?  
    - A: no, it’s just a new interface to call the same tests  
  - Action: David to ask for justification, estimation of migration effort  
- Should we move \[sig-storage\] tests from [https://testgrid.k8s.io/sig-node-release-blocking\#node-kubelet-master](https://testgrid.k8s.io/sig-node-release-blocking#node-kubelet-master) to some place under [https://testgrid.k8s.io/sig-storage-kubernetes](https://testgrid.k8s.io/sig-storage-kubernetes)?  
  - Amim: prefer to keep tests, maybe remove flaky tests  
  - Sergey: need a mechanism to notify those teams  
- Move \[sig-network\] tests from [https://testgrid.k8s.io/sig-node-release-blocking\#node-kubelet-master](https://testgrid.k8s.io/sig-node-release-blocking#node-kubelet-master) to some place under [https://testgrid.k8s.io/sig-network](https://testgrid.k8s.io/sig-network)   
- Fix infra issue at: [https://testgrid.k8s.io/sig-node-critical\#kubelet-NodeConformance](https://testgrid.k8s.io/sig-node-critical#kubelet-NodeConformance)  
  - Qiutong to take a look  
- Clean-up [https://testgrid.k8s.io/sig-node-kubelet\#node-kubelet-serial](https://testgrid.k8s.io/sig-node-kubelet#node-kubelet-serial)   
  - Francesco: will try.  
- Multi-zone tests: [https://github.com/kubernetes/test-infra/pull/21777](https://github.com/kubernetes/test-infra/pull/21777)   
  - Elana: looks like cloud provider specific  
  -  Sergey: to look at what these tests suppose to test-cover

# 04/07/2021

Attendees:

- Elana Hashman  
- Sergey Kanzhelev  
- Alukiano  
- Tao  
- Qiutong

Agrenda:

- Artem will start looking at fake NUMA flag.

# 03/30/2021 Cancelled

# 03/24/2021

Agenda

[https://github.com/kubernetes/kubernetes/issues?q=is%3Aopen+milestone%3Av1.21+label%3Asig%2Fnode](https://github.com/kubernetes/kubernetes/issues?q=is%3Aopen+milestone%3Av1.21+label%3Asig%2Fnode) 

Discussing [https://github.com/kubernetes/kubernetes/pull/99336](https://github.com/kubernetes/kubernetes/pull/99336):

- Overall feeling is that it’s too late for unknown unknowns introduced by this PR.  
- 1.18 cherry-picking is not Node team problem, more release team problem. Maybe release team will need an exception.

# 03/17/2021

Attendees 

Agenda

- \[ehashman\] Discuss the future of node presubmit tests  
  - Context: [https://github.com/kubernetes/test-infra/pull/21278](https://github.com/kubernetes/test-infra/pull/21278) [https://kubernetes.slack.com/archives/C0BP8PW9G/p1615470977302300](https://kubernetes.slack.com/archives/C0BP8PW9G/p1615470977302300) 

Direction long term is not to test the whole matrix on presubmits, but have a good signal with failures easy to investigate by contributors. Maybe PR needs to be replaced with a single job with both runtimes.

**Action:** determine a long-term plan to merge all node presubmits into one job, using a name that doesn’t reveal the underlying runtimes. (e.g. cleanup ubuntu-containerd\* tests)  
Timeframe: dependent on a presubmit policy, maybe will happen next release cycle (1.22)

- [https://github.com/kubernetes/kubernetes/issues/94289](https://github.com/kubernetes/kubernetes/issues/94289)   
  - Sergey: list all the tags and decide what to do about it  
- [https://github.com/kubernetes/kubernetes/issues/96524](https://github.com/kubernetes/kubernetes/issues/96524)   
  - Sergey’s todo.

# 03/10/2021

Attendees 

* ehashman  
* fromanirh  
* swsehgal  
* 

Agenda

- Deflake of sig-node-alpha tab [https://testgrid.k8s.io/sig-node-kubelet\#node-kubelet-alpha](https://testgrid.k8s.io/sig-node-kubelet#node-kubelet-alpha)   
  - It’s done\!\!

# 03/03/2021

Attendees 

Agenda

- Containerd tests. TODO: insert links  
- Triage  
- \[alukiano\] Serial job \- timeouts because OOMs. Artem will switch to 2 Gb \- log [https://storage.googleapis.com/kubernetes-jenkins/logs/ci-kubernetes-node-kubelet-serial/1365962673736388608/artifacts/tmp-node-e2e-a988a9a1-cos-81-12871-1245-10/kern.log](https://storage.googleapis.com/kubernetes-jenkins/logs/ci-kubernetes-node-kubelet-serial/1365962673736388608/artifacts/tmp-node-e2e-a988a9a1-cos-81-12871-1245-10/kern.log)

# 02/24/2021. 

Attendees (7 on call):  
![][image18]

- Triage [https://github.com/orgs/kubernetes/projects/43](https://github.com/orgs/kubernetes/projects/43) 

1. Containerd plan [https://github.com/kubernetes/test-infra/issues/18570](https://github.com/kubernetes/test-infra/issues/18570)  
2. Questions about [https://github.com/kubernetes/test-infra/pull/20937](https://github.com/kubernetes/test-infra/pull/20937) and node-kubelet-serial tests ([https://testgrid.k8s.io/sig-node-kubelet\#node-kubelet-serial](https://testgrid.k8s.io/sig-node-kubelet#node-kubelet-serial))   
3. Announcement: node n-2 version skew tests to be discussed at SIG Arch tomorrow: [https://groups.google.com/g/kubernetes-sig-architecture/c/QX-4qq2krm0/m/998T3cJUBQAJ](https://groups.google.com/g/kubernetes-sig-architecture/c/QX-4qq2krm0/m/998T3cJUBQAJ) 

Product triage: [https://github.com/orgs/kubernetes/projects/49](https://github.com/orgs/kubernetes/projects/49) 

- Feature PRs missing from board that happen to have sig/testing label  
- Action: needs-rebase isn’t auto-applied, bot needs to be pestered. File issue to proactively apply without resetting stale counter  
  - [https://github.com/kubernetes/test-infra/issues/21006](https://github.com/kubernetes/test-infra/issues/21006) 

# 02/17/2021

- Triage [https://github.com/orgs/kubernetes/projects/43](https://github.com/orgs/kubernetes/projects/43)   
- SIG-labels for all tests  
- Morgan as approver  
- \[fromani\] CPU manager e2e tests needs improvement.

Product triage: [https://github.com/orgs/kubernetes/projects/49](https://github.com/orgs/kubernetes/projects/49) 

# 02/08/2021

\[Sergey\] New time for the meeting? It looks like 10 AM Mon is very inconvenient. Is Monday 9AM better?

[https://doodle.com/poll/ii5vyde6wpp3migm?utm\_campaign=poll\_update\_participant\_admin\&utm\_medium=email\&utm\_source=poll\_transactional\&utm\_content=gotopoll-cta\#table](https://doodle.com/poll/ii5vyde6wpp3migm?utm_campaign=poll_update_participant_admin&utm_medium=email&utm_source=poll_transactional&utm_content=gotopoll-cta#table)

![][image19]

Triage: [https://github.com/orgs/kubernetes/projects/43](https://github.com/orgs/kubernetes/projects/43) 

- Suggest MHBauer to approver. Elana to reach out  
- 

# 02/01/2021 \[skipping\]

# 01/25/2021

Agenda:  
\[Aditi\]

- Test plan for containerd [https://github.com/kubernetes/test-infra/issues/18570](https://github.com/kubernetes/test-infra/issues/18570#issuecomment-764576223)  
  Some test analysis here [https://docs.google.com/spreadsheets/d/1mN1fG0dq6t7dZTzl-g9fFNwFqD8XdbbKlr5D3R7BRzc/edit\#gid=0](https://docs.google.com/spreadsheets/d/1mN1fG0dq6t7dZTzl-g9fFNwFqD8XdbbKlr5D3R7BRzc/edit#gid=0)  
  Some answers we want [https://github.com/kubernetes/test-infra/issues/18570\#issue comment-764576223](https://github.com/kubernetes/test-infra/issues/18570#issuecomment-764576223)  
- Aditi: Updated action items here [https://github.com/kubernetes/test-infra/issues/18570\#issuecomment-767024601](https://github.com/kubernetes/test-infra/issues/18570#issuecomment-767024601) 

\[Sergey\] Do we need this:   
[https://github.com/kubernetes/k8s.io/issues/956\#issuecomment-764847132](https://github.com/kubernetes/k8s.io/issues/956#issuecomment-764847132)?

- Need for deflaking tests. “Pause” job and SSH access  
- Sometime hard to understand what failed and why PR validation failed

\[knabben\]

- Deflake of these tabs:  
  - [Node-kubelet-orphans](https://github.com/kubernetes/kubernetes/issues/98265) (partial)  
  - [Node-kubelet-alpha](https://github.com/kubernetes/kubernetes/issues/98220)

\[ehashman\]

* Non-CI PR triage: Node board [https://github.com/orgs/kubernetes/projects/49](https://github.com/orgs/kubernetes/projects/49)   
* Action (ehashman): will add note cards to columns of project board to explain how it works  
* Action (ehashman): will draft doc on HOW-TO node review for community repo, tag attendees, and leave open for lazy consensus/feedback

\[Meeting time\]

* Action: Sergey to send doodle for maybe moving this meeting? Mondays frequently conflict.

# 01/11/2021

Attendees (7 on the call):

Agenda:

- triage

# 01/04/2021

Attendees:

Agenda:

\[knabben\]

- [Node s/gci/cos/g tab rename](https://docs.google.com/document/d/1KlnfvQ_OPkrty5DvKHmeEMTKMpaVK-ioHgDOpVFRrkA/) \- [https://github.com/kubernetes/test-infra/pull/20351](https://github.com/kubernetes/test-infra/pull/20351/files)  
- Adds a **\--restart-kubelet** flag on Node E2E tests \- PTAL [https://github.com/kubernetes/kubernetes/pull/97028/](https://github.com/kubernetes/kubernetes/pull/97028/)

\[Sergey\] Triage

# Volunteers to help with this effort

Victor Pickard (Red Hat), nick=vpickard, vpickard@redhat.com  
Jay Pipes (AWS), nick=jaypipes, gh=jaypipes, jaypipes@gmail.com  
Balaji (AWS), nick=srisaranbalaji, gh=SaranBalaji90, srisaranbalaji@gmail.com  
Morgan Bauer (IBM), slack=mhb, gh=mhbauer, mail=mbauer@us.ibm.com  
Ning Liao (Google), nick=nliao,  mail=ningliao@google.com  
David Porter (Google), nick=davidporter; mail=[porterdavid@google.com](mailto:porterdavid@google.com)  
Hanfei Lin (Google), nick=hanfeil; mail=hanfeil@google.com  
Hugo Huang (Google), nick=tangent; mail=[tangent@google.com](mailto:tangent@google.com)  
Roy Yang(Google), nick=roy; mail=royyang@google.com  
Aaron Crickenberger (Google), nick=spiffxp, spiffxp@gmail.com  
nick=Archer  
Ed Bartosh (Intel), slack=Ed, github=bart0sh [eduard.bartosh@intel.com](mailto:eduard.bartosh@intel.com)  
Daniel Mangum (upbound.io), nick=hasheddan  
Chirag Tayal (PayPal) nick=ctayal, [chiragtayal@gmail.com](mailto:chiragtayal@gmail.com)  
Zhi Feng(Airbnb), nick=Zhi, helenfengzhi@gmail.com  
Dims, nick=dims, [davanum@gmail.com](mailto:davanum@gmail.com)   
Jacob Blain Christen (Rancher Labs), nick=dweomer; mail=[dweomer5@gmail.com](mailto:dweomer5@gmail.com)  
Artyom Lukianov(Red Hat), nick(github)=cynepco3hahue,nick(slack)=alukiano,mail=[alukiano@redhat.com](mailto:alukiano@redhat.com)  
Swati Sehgal (Red Hat), slack=swsehgal; mail [swsehgal@redhat.com](mailto:swsehgal@redhat.com)  
Jorge Alarcon, nick=alejandrox1, [alarcj137@gmail.com](mailto:alarcj137@gmail.com)   
Sascha Grunert (SUSE), nick=sascha, [sgrunert@suse.com](mailto:sgrunert@suse.com)  
Srini Brahmaroutu(IBM), slack=srbrahma, gh=brahmaroutu, mail=[srbrahma@us.ibm.com](mailto:srbrahma@us.ibm.com)  
Tim Pepper (VMware), slack=tpepper, gh=tpepper, mail=[tpepper@vmware.com](mailto:tpepper@vmware.com)   
John Taylor (IBM), mail=[jtaylor1@uk.ibm.com](mailto:jtaylor1@uk.ibm.com)  
Francesco Romani (Red Hat), nick=fromani; mail=[fromani@redhat.com](mailto:fromani@redhat.com)  
Karan Goel (Google), nick=karan; mail=[karangoel@google.com](mailto:karangoel@google.com)  
Sergey Kanzhelev (Google), nick=SergeyKanzhelev; mail=[skanzhelev@google.com](mailto:skanzhelev@google.com)  
Mike Carlise (Salesforce), nick=micarlise, mail=[micarlise@gmail.com](mailto:micarlise@gmail.com)  
Matt Merkes (AWS), nick=merkes, mail=[matt.merkes@gmail.com](mailto:matt.merkes@gmail.com)  
Amim Knabben (Loadsmart), nick=knabben, mail=[amim.knabben@gmail.com](mailto:amim.knabben@gmail.com)  
Swati Sehgal (Red Hat), nick(slack)=swsehgal, nick(github)= swatisehgal, mail \= [swsehgal@redhat.com](mailto:swsehgal@redhat.com)  
Harshal Patil (Red Hat), slack=Harshal, gh=harche, mail=[harpatil@redhat.com](mailto:harpatil@redhat.com)  
Elana Hashman (Red Hat), nick=ehashman, mail=[ehashman@redhat.com](mailto:ehashman@redhat.com)   
Paco Xu(DaoCloud), nick=paco,mail=[paco.xu@daocloud.io](mailto:paco.xu@daocloud.io)
