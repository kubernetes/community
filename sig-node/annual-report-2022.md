# 2022 Annual Report: SIG Node

## Current initiatives

1. What work did the SIG do this year that should be highlighted?

   - Dockershim removal has big effect on community, 3rd party tools vendors and end users.
   - With CRI v1alpha2 removal we set the minimal compatible Containerd version. This is continuous trend as we require newer and newer dependencies.
   - Cgroups v2 GA opens up possibility for new features, but sets the minimal bar for dependencies again.
   - Cgroups v1 deprecation is on horizon.
   - No perma beta progress - keep removing and/or promoting old features. For example, DynamicKubeletConfig removed, various resource managers GA'd.
   - User namespaces for stateless pods entered alpha. It is a very old feature request.
   - Evented PLEG - working on minimization of kubelet overhead at bigger scale.
   - DRA was introduced which opens up more device integration scenarios.

2. What initiatives are you working on that aren't being tracked in KEPs?

   - Keep working on CI stability and overall reliability of SIG Node components.
   - Refactoring E2E Node Tests to run against multiple cloud providers more easily.
   - Infra rehosting: new image registry and working on e2e tests on AWS.

3. KEP work in 2022 (v1.24, v1.25, v1.26):

  - alpha:
    - [3386 - Evented PLEG](https://github.com/kubernetes/enhancements/blob/17483a9ce33f303ec1993a781c0c63195c0569e0/keps/sig-node/3386-kubelet-evented-pleg) - v1.25
    - [127 - Support User Namespaces in stateless pods](https://github.com/kubernetes/enhancements/blob/master/keps/sig-node/127-user-namespaces) - v1.25
    - [2008 - Forensic Container Checkpointing](https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/2008-forensic-container-checkpointing) - v1.25
    - [2831 - Kubelet OpenTelemetry tracing](https://github.com/kubernetes/enhancements/blob/master/keps/sig-instrumentation/2831-kubelet-tracing) - v1.25
    - [3085 - Pod networking ready condition](https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/3085-pod-conditions-for-starting-completition-of-sandbox-creation) - v1.25
    - [3327 - CPUManager policy option to align CPUs by Socket instead of by NUMA node](https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/3327-align-by-socket) - v1.25
    - [2371 - cAdvisor-less, CRI-full Container and Pod Stats](https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/2371-cri-pod-container-stats) - v1.26
    - [3063 - dynamic resource allocation](https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/3063-dynamic-resource-allocation) - v1.26
    - [3545 - Improved multi-numa alignment in Topology Manager](https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/3545-improved-multi-numa-alignment) - v1.26

  - removed 
    - [281 - Dynamic Kubelet Configuration](https://github.com/kubernetes/enhancements/blob/master/keps/sig-node/281-dynamic-kubelet-configuration) 281 - v1.24
    - [2221 - Removing dockershim from kubelet](https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/2221-remove-dockershim) - v1.24

  - beta:
    - [2712 - Pod Priority Based Graceful Node Shutdown](https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/2712-pod-priority-based-graceful-node-shutdown) - v1.24
    - [2727 - Add gRPC probe to Pod.Spec.Container.{Liveness,Readiness,Startup}Probe](https://github.com/kubernetes/enhancements/blob/master/keps/sig-node/2727-grpc-probe) - v1.24
    - [2238 - Add configurable grace period to probes](https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/2238-liveness-probe-grace-period) - v1.25
    - [2413 - seccomp by default](https://github.com/kubernetes/enhancements/blob/master/keps/sig-node/2413-seccomp-by-default) - v1.25
    - [3329 - Retriable and non-retriable Pod failures for Jobs](https://github.com/kubernetes/enhancements/tree/master/keps/sig-apps/3329-retriable-and-non-retriable-failures) - v1.26

  - stable
    - [688 - Pod Overhead](https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/688-pod-overhead) - v1.24
    - [2133 - Kubelet Credential Providers](https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/2133-kubelet-credential-providers) - v1.26
    - [2254 - cgroups v2](https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/2254-cgroup-v2) - v1.25
    - [3570 - CPU Manager](https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/3570-cpumanager) - v1.26
    - [3573 - Device Manager Proposal](https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/3573-device-plugin) - v1.26
    - [277 - Ephemeral Containers](https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/277-ephemeral-containers) - v1.25

## Project health

1. What areas and/or subprojects does your group need the most help with?
   Any areas with 2 or fewer OWNERs? (link to more details)

   - Node Problem Detector is not very active and has small participation.
   - In general all areas are well covered with reviewers and approvers. However community would benefit to have clearer ownership of areas defined someplace. Ideas (https://github.com/kubernetes/community/issues/7234):
     - Add more people to subdirectories as approvers
     - Split more clearly between approvers and reviewers

2. What metrics/community health stats does your group care about and/or measure?

   - Active PRs and weekly changes of PRs at weekly SIG meetings.
   - Untriaged PRs, bugs - via CI group bug triage.
   - It will be great to have a PR full-cycle metrics to get insights into
     time to review and time to approve as well as number of iterations.

3. Does your [CONTRIBUTING.md] help **new** contributors engage with your group specifically by pointing
   to activities or programs that provide useful context or allow easy participation?

   - No, current [CONTRIBUTING.md] is not adequately covers the getting started experience for the new contributors.
   - Ideas for improvements: https://github.com/kubernetes/community/issues/7223
   - Also we need clearer code organization documentation for Kubernetes. This is not specific to SIG Node, but due to size, SIG Node group would likely benefit the most from it.

4. If your group has special training, requirements for reviewers/approvers, or processes beyond the general [contributor guide],
   does your [CONTRIBUTING.md] document those to help **existing** contributors grow throughout the [contributor ladder]?

   - We published the SIG Node [contributor ladder](sig-node-contributor-ladder.md)
     that details requirements for reviewers/approvers.

5. Does the group have contributors from multiple companies/affiliations?

   - The group has contributors from multiple companies/affiliations.
   - [19 companies made 1+](https://k8s.devstats.cncf.io/d/8/company-statistics-by-repository-group?orgId=1&var-period=y&var-metric=contributions&var-repogroup_name=SIG%20Node&var-companies=All&from=now-1y&to=now) contributions over the last year.

6. Are there ways end users/companies can contribute that they currently are not?
   If one of those ways is more full time support, what would they work on and why?

   - Drive dockershim deprecation and report back issues. We still see the adoption
     not being very fast. It will be great to see more 3rd party tools vendors helping with migration.
   - We have another potential deprecation - cgroupv1 upcoming. It will likely be
     less impactful, but may need support.
   - Need more feedback from end users on various deprecations and a new features.
      - Can we publish surveys and check on features usage from end users?

## Membership

- Primary slack channel member count: 3570 (#sig-node)
- Primary mailing list member count: 840
- Primary meeting attendee count (estimated, if needed): 30 (last numbers: 35, 22, 22, 26, 33, 36, 26, 31)
- Primary meeting participant count (estimated, if needed): 10 (estimated based on # of agenda items)
<!-- 
Reviewers: 78 [AlexeyPerevalov ArangoGutierrez ConnorDoyle RainbowMango Random-Liu SergeyKanzhelev Vincent056 adrianchiris andrewsykim andyxning aojea bart0sh bobbypage bowei caesarxuchao caseydavenport chrishenzie coffeepac danwinship dashpole dcbw dchen1107 deads2k derekwaynecarr dgrisonnet dims ehashman endocrimes enj feiskyer fmuyassarov freehan gnufied haircommander humblec jingxu97 jjacobelli jsafrane kad khenidak klueska krmayankk lavalamp liggitt logicalhan marquiz mattcary matthyx mauriciopoppe mikedanese mrhohn mrunalp msau42 mtaufen odinuge pacoxu pohly qbarrand robscott s-urbaniak saikat-royc serathius sjenning smarterclayton sttts swatisehgal tallclair thockin vteratipally wangzhen127 wojtek-t wzshiming xing-yang xueweiz ybettan yevgeny-shnaidman yujuhong zvonkok]

Approvers: 58 [ArangoGutierrez ConnorDoyle JAORMX Random-Liu SergeyKanzhelev andrewsykim andyxning aojea bart0sh bowei caseydavenport ccojocar cheftako danwinship dashpole dcbw dchen1107 deads2k derekwaynecarr dims endocrimes feiskyer freehan gnufied haircommander harche jhrozek jingxu97 jsafrane khenidak klueska lavalamp liggitt marquiz mikedanese mrhohn mrunalp msau42 mtaufen pjbgf pohly qbarrand robscott saad-ali sairameshv saschagrunert sjenning smarterclayton swatisehgal tallclair thockin vteratipally wangzhen127 xing-yang xueweiz ybettan yevgeny-shnaidman yujuhong]
 -->
- Unique reviewers for SIG-owned packages: 78
- Unique approvers for SIG-owned packages: 58

## [Subprojects](https://git.k8s.io/community/sig-node#subprojects)

**New in 2022:**

  - ci-testing
  - kernel-module-management

**Continuing:**

  - cri-api
  - cri-tools
  - kubelet
  - node-api
  - node-feature-discovery
  - node-problem-detector
  - noderesourcetopology-api
  - security-profiles-operator

## [Working groups](https://git.k8s.io/community/sig-node#working-groups)

**New in 2022:**

 - Batch
 - wg sidecar (https://github.com/kubernetes/community/pull/7233)

**Continuing:**

 - Multitenancy
 - Policy
 - Structured Logging

## Operational

Operational tasks in [sig-governance.md]:

- [X] [README.md] reviewed for accuracy and updated if needed
- [X] [CONTRIBUTING.md] reviewed for accuracy and updated if needed
      (or created if missing and your contributor steps and experience are different or more
      in-depth than the documentation listed in the general [contributor guide] and [devel] folder.)
      Task tracking improvements:  https://github.com/kubernetes/community/issues/7223
- [X] Subprojects list and linked OWNERS files in [sigs.yaml] reviewed for accuracy and updated if needed. Created https://github.com/kubernetes/community/issues/7234
- [X] SIG leaders (chairs, tech leads, and subproject owners) in [sigs.yaml] are accurate and active, and updated if needed
- [X] Meeting notes and recordings for 2022 are linked from [README.md] and updated/uploaded if needed
- [X] Did you have community-wide updates in 2022 (e.g. community meetings, kubecon, or kubernetes-dev@ emails)? Links to email, slides, or recordings:
      - KubeCon NA 2022
        - Recording: https://youtu.be/MxvhuhQpuAc
        - Sched: https://sched.co/182Pi
        - Slides: https://static.sched.com/hosted_files/kccncna2022/da/SIG%20Node%20KubeCon%20NA%202022.pptx.pdf
      - KubeCon EU 2022
        - Recording: https://youtu.be/FGRenKv4RgY
        - Sched: https://sched.co/ytue
        - Slides: https://static.sched.com/hosted_files/kccnceu2022/dd/Copy%20of%20KubeCon%20EU%202022%20SIG%20Node%20maintainers%20track.pdf

[CONTRIBUTING.md]: https://git.k8s.io/community/sig-node/CONTRIBUTING.md
[contributor ladder]: https://git.k8s.io/community/community-membership.md
[sig-governance.md]: https://git.k8s.io/community/committee-steering/governance/sig-governance.md
[README.md]: https://git.k8s.io/community/sig-node/README.md
[sigs.yaml]: https://git.k8s.io/community/sigs.yaml
[contributor guide]: https://git.k8s.io/community/contributors/guide/README.md
[devel]: https://git.k8s.io/community/contributors/devel/README.md
