# 2022 Annual Report: SIG Node

## Current initiatives

1. What work did the SIG do this year that should be highlighted?

   -
   -
   -

2. What initiatives are you working on that aren't being tracked in KEPs?

   -
   -
   -



3. KEP work in 2022 (v1.24, v1.25, v1.26):
  - alpha:
    - [2008 - Forensic Container Checkpointing](https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/2008-forensic-container-checkpointing) - v1.25
    - [2371 - cAdvisor-less, CRI-full Container and Pod Stats](https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/2371-cri-pod-container-stats) - v1.26
    - [2535 - Ensure Secret Pulled Images](https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/2535-ensure-secret-pulled-images) - v1.24
    - [3063 - dynamic resource allocation](https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/3063-dynamic-resource-allocation) - v1.26
    - [3085 - Pod networking ready condition](https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/3085-pod-conditions-for-starting-completition-of-sandbox-creation) - v1.25
    - [3288 - Split Stdout and Stderr Log Stream of Container](https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/3288-separate-stdout-from-stderr) - v1.25
    - [3327 - CPUManager policy option to align CPUs by Socket instead of by NUMA node](https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/3327-align-by-socket) - v1.25
    - [3545 - Improved multi-numa alignment in Topology Manager](https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/3545-improved-multi-numa-alignment) - v1.26
  - beta:
    - [2712 - Pod Priority Based Graceful Node Shutdown](https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/2712-pod-priority-based-graceful-node-shutdown) - v1.24
  - stable:
    - [1972 - Kubelet Exec Probe Timeouts](https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/1972-kubelet-exec-probe-timeouts) - v1.24
    - [2133 - Kubelet Credential Providers](https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/2133-kubelet-credential-providers) - v1.26
    - [2221 - Removing dockershim from kubelet](https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/2221-remove-dockershim) - v1.24
    - [2254 - cgroups v2](https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/2254-cgroup-v2) - v1.25
    - [277 - Ephemeral Containers](https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/277-ephemeral-containers) - v1.25
    - [3570 - CPU Manager](https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/3570-cpumanager) - v1.26
    - [3573 - Device Manager Proposal](https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/3573-device-plugin) - v1.26
    - [688 - Pod Overhead](https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/688-pod-overhead) - v1.24


## Project health

1. What areas and/or subprojects does your group need the most help with?
   Any areas with 2 or fewer OWNERs? (link to more details)

   -
   -
   -

2. What metrics/community health stats does your group care about and/or measure?

   -
   -
   -

3. Does your [CONTRIBUTING.md] help **new** contributors engage with your group specifically by pointing
   to activities or programs that provide useful context or allow easy participation?

   -

4. If your group has special training, requirements for reviewers/approvers, or processes beyond the general [contributor guide],
   does your [CONTRIBUTING.md] document those to help **existing** contributors grow throughout the [contributor ladder]?

   -

5. Does the group have contributors from multiple companies/affiliations?

   -

6. Are there ways end users/companies can contribute that they currently are not?
   If one of those ways is more full time support, what would they work on and why?

   -
   -

## Membership

- Primary slack channel member count:
- Primary mailing list member count:
- Primary meeting attendee count (estimated, if needed):
- Primary meeting participant count (estimated, if needed):
- Unique reviewers for SIG-owned packages: <!-- in future, this will be generated from OWNERS files referenced from subprojects, expanded with OWNERS_ALIASES files -->
- Unique approvers for SIG-owned packages: <!-- in future, this will be generated from OWNERS files referenced from subprojects, expanded with OWNERS_ALIASES files -->

Include any other ways you measure group membership

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

**Continuing:**

 - Multitenancy
 - Policy
 - Structured Logging

## Operational

Operational tasks in [sig-governance.md]:

- [ ] [README.md] reviewed for accuracy and updated if needed
- [ ] [CONTRIBUTING.md] reviewed for accuracy and updated if needed
      (or created if missing and your contributor steps and experience are different or more
      in-depth than the documentation listed in the general [contributor guide] and [devel] folder.)
- [ ] Subprojects list and linked OWNERS files in [sigs.yaml] reviewed for accuracy and updated if needed
- [ ] SIG leaders (chairs, tech leads, and subproject owners) in [sigs.yaml] are accurate and active, and updated if needed
- [ ] Meeting notes and recordings for 2022 are linked from [README.md] and updated/uploaded if needed
- [ ] Did you have community-wide updates in 2022 (e.g. community meetings, kubecon, or kubernetes-dev@ emails)? Links to email, slides, or recordings:
      -
      -

[CONTRIBUTING.md]: https://git.k8s.io/community/sig-node/CONTRIBUTING.md
[contributor ladder]: https://git.k8s.io/community/community-membership.md
[sig-governance.md]: https://git.k8s.io/community/committee-steering/governance/sig-governance.md
[README.md]: https://git.k8s.io/community/sig-node/README.md
[sigs.yaml]: https://git.k8s.io/community/sigs.yaml
[contributor guide]: https://git.k8s.io/community/contributors/guide/README.md
[devel]: https://git.k8s.io/community/contributors/devel/README.md
