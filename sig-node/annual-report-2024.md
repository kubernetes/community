# 2024 Annual Report: SIG Node

## Current initiatives and Project Health

1. What work did the SIG do this year that should be highlighted?

SIG-Node remains a structural piece of the Kubernetes community and the span of the work done in 2024 highlights that.
As the community continues rallying behind AI use cases and identifying gaps with Kubernetes as a platform for LLM training and serving,
SIG-Node made strides in multiple AI related areas. DRA structured parameters made it to beta, meaning more flexible scheduling and allocation of
device resources is now possible. In 2025 there will be a lot of continued work on DRA, including enhancing drivers to be able to report device health
and Kubernetes components be able to react to that, extending DRA to support advanced networking use cases, device taints and tolerations, and lots more!
Outside of DRA, OCI image volume mounts have been added as alpha in 2024, allowing users to mount AI models into containers via a separate image (and one day artifact) instead
of a model car or embedding it in the container image. Also, work like in-place pod resize and pod level resource limits will unlock use cases for power AI users: allowing more flexibility
in pod resource limit calculation at both initialization and during runtime.

Plenty of work has been being done outside of AI as well! SIG-Node remains the top SIG in KEPs progressing, moving forward on 13, 16, and 17 KEPs between 1.30, 1.31, and 1.32 respectively.
Lots of progress has been made in the CPU manager: like adding support for split uncore cache, adding a policy option for restricting resrevedSystemCPUS and a new static policy for optimizing CPU alignment.
We have also worked on some long awaited linux technologies like user namespaces, swap, AppArmor, ephemeral storage quotas, recursive read only mounts, and better support for supplemental groups,
as well as announced feature freeze on cgroupv1.

All of these features don't even begin to cover the amount of CI stabilization, bug fixes, and other work the SIG is doing. We remain a productive (albeit, occasionally overbooked) SIG. To help keep up with all
of the work, we've inducted one new approver Sergey Kanzhelev, reinducted a formerly emertius approver Tim Allclair, welcomed a new SIG chair Peter Hunt, as well as began crafting a role to help KEP authors follow along the KEP process, currently called the KEP wranglers.

2. Are there any areas and/or subprojects that your group needs help with (e.g. fewer than 2 active OWNERS)?

SIG-Node, in being so busy, always has a bottleneck of top level approvers. Any path in the kubelet could use more people who have expertise and confidence in reviewing. Please refer to our
[contributor ladder](https://github.com/kubernetes/community/blob/master/sig-node/sig-node-contributor-ladder.md) to see ways to grow in the SIG!

3. Did you have community-wide updates in 2024 (e.g. KubeCon talks)?

- Kubecon EU 2024 [maintainers track](https://kccnceu2024.sched.com/event/1YhjL/kubernetes-sig-node-intro-and-deep-dive-dixita-narang-dawn-chen-google-matthias-bertschy-armo-peter-hunt-red-hat)
- Kubecon NA 2024 [maintainers track](https://kccncna2024.sched.com/event/1hovs/sig-node-intro-and-deep-dive-sergey-kanzhelev-dawn-chen-google-mrunal-patel-red-hat)

4. KEP work in 2024 (v1.30, v1.31, v1.32):
  - Alpha
    - [2837 - KEP Template](https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/2837-pod-level-resource-spec) - v1.32
    - [2862 - Fine grained Kubelet API authorization](https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/2862-fine-grained-kubelet-authz) - v1.32
    - [3288 - Split Stdout and Stderr Log Stream of Container](https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/3288-separate-stdout-from-stderr) - v1.32
    - [3619 - Fine grained SupplementalGroups control](https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/3619-supplemental-groups-policy) - v1.31
    - [4438 - Restarting sidecar containers during Pod termination](https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/4438-container-restart-termination) - v1.32
    - [4540 - Add CPUManager policy option to restrict reservedSystemCPUs to system daemons and interrupt processing](https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/4540-strict-cpu-reservation) - v1.32
    - [4580 - Deprecate & remove Kubelet RunOnce mode](https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/4580-deprecate-kubelet-runonce) - v1.31
    - [4603 - Tune Crashloop Backoff](https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/4603-tune-crashloopbackoff) - v1.32
    - [4680 - Add Resource Health Status to the Pod Status for Device Plugin and DRA](https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/4680-add-resource-health-to-pod-status) - v1.31
    - [4800 - Split UnCoreCache Toplogy Awareness in CPU Manager](https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/4800-cpumanager-split-uncorecache) - v1.32
    - [4815 - DRA Partitionable Devices](https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/4815-dra-partitionable-devices) - v1.32
    - [4817 - Resource Claim Status With Possible Standardized Network Interface Data](https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/4817-resource-claim-device-status) - v1.32
    - [4818 - Allow zero value for Sleep Action of PreStop Hook](https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/4818-allow-zero-value-for-sleep-action-of-prestop-hook) - v1.32

  - Beta
    - [1029 - Quotas for Ephemeral Storage](https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/1029-ephemeral-storage-quotas) - 1.31
    - [127 - Support User Namespaces](https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/127-user-namespaces) - v1.30
    - [1287 - In-place Update of Pod Resources](https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/1287-in-place-update-pod-resources) - v1.32
    - [2008 - Forensic Container Checkpointing](https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/2008-forensic-container-checkpointing) - v1.30
    - [2400 - Node system swap support](https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/2400-node-swap) - v1.30
    - [3857 - Recursive read-only mounts](https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/3857-rro-mounts) - v1.31
    - [3983 - Add support for a kubelet drop-in configuration directory](https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/3983-drop-in-configuration) - v1.31
    - [4033 - Discover cgroup driver from CRI](https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/4033-group-driver-detection-over-cri) - v1.31
    - [4176 - New CPUManager Static Policy which spread hyperthreads across physical CPUs to better utilize CPU Cache](https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/4176-cpumanager-spread-cpus-preferred-policy) - v1.31
    - [4191 - Split Image Filesystem](https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/4191-split-image-filesystem) - v1.31
    - [4210 - ImageMaximumGCAge in Kubelet](https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/4210-max-image-gc-age) - v1.30
    - [4216 - Image pull per runtime class](https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/4216-image-pull-per-runtime-class) - v1.31
    - [4265 - Add ProcMount option](https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/4265-proc-mount) - v1.31
    - [4369 - Allow special characters environment variable](https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/4369-allow-special-characters-environment-variable) - v1.32
    - [4381 - DRA Structured Parameters](https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/4381-dra-structured-parameters) - v1.32
    - [4639 - OCI objects as VolumeSource](https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/4639-oci-volume-source) - v1.32

  - Stable
    - [1769 - Memory Manager](https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/1769-memory-manager) - v1.32
    - [1967 - Size memory backed volumes](https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/1967-size-memory-backed-volumes) - v1.32
    - [24 - Add AppArmor Support](https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/24-apparmor) - v1.31
    - [3545 - Improved multi-numa alignment in Topology Manager](https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/3545-improved-multi-numa-alignment) - v1.32
    - [3673 - Kubelet limit of Parallel Image Pulls](https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/3673-kubelet-parallel-image-pull-limit) - v1.32
    - [3960 - Pod lifecycle sleep action](https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/3960-pod-lifecycle-sleep-action) - v1.32
    - [4009 - Add CDI devices to device plugin API](https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/4009-add-cdi-devices-to-device-plugin-api) - v1.31
    - [4188 - New kubelet gRPC API with endpoint returning local pods information](https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/4188-kubelet-pod-readiness-api) - v1.31
    - [4569 - Move cgroup v1 in maintenance mode](https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/4569-cgroup-v1-maintenance-mode) - v1.31
    - [4622 - New TopologyManager Policy which configure the value of maxAllowableNUMANodes](https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/4622-topologymanager-max-allowable-numa-nodes) - v1.32 -->

## [Subprojects](https://git.k8s.io/community/sig-node#subprojects)


**New in 2024:**
  - cri-client
**Continuing:**
  - ci-testing
  - cri-api
  - cri-tools
  - kernel-module-management
  - kubelet
  - node-api
  - node-feature-discovery
  - node-problem-detector
  - resource-management
  - security-profiles-operator

## [Working groups](https://git.k8s.io/community/sig-node#working-groups)

**New in 2024:**
 - Device Management
 - Serving
**Continuing:**
 - Batch
 - Policy
 - Structured Logging

## Operational

Operational tasks in [sig-governance.md]:
- [x] [README.md] reviewed for accuracy and updated if needed
- [x] [CONTRIBUTING.md] reviewed for accuracy and updated if needed
- [x] Other contributing docs (e.g. in devel dir or contributor guide) reviewed for accuracy and updated if needed
- [x] Subprojects list and linked OWNERS files in [sigs.yaml] reviewed for accuracy and updated if needed
- [x] SIG leaders (chairs, tech leads, and subproject leads) in [sigs.yaml] are accurate and active, and updated if needed
- [x] Meeting notes and recordings for 2024 are linked from [README.md] and updated/uploaded if needed


[CONTRIBUTING.md]: https://git.k8s.io/community/sig-node/CONTRIBUTING.md
[sig-governance.md]: https://git.k8s.io/community/committee-steering/governance/sig-governance.md
[README.md]: https://git.k8s.io/community/sig-node/README.md
[sigs.yaml]: https://git.k8s.io/community/sigs.yaml
[devel]: https://git.k8s.io/community/contributors/devel/README.md
