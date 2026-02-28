# 2025 Annual Report: SIG Node

## Current initiatives and Project Health

1. What work did the SIG do this year that should be highlighted?

SIG Node continues its streak as the record holding SIG for most KEPs each release. 2025 actually exceeded the velocity of 2024, with 1.33 merging 33 KEPs, 1.34 merging 32 and 1.35 merging 35!
Many major initiatives across many different areas were advanced. Particularly, many long-standing KEPs were moved to stable,
which represents an overarching initative to close out KEPs and stabalize the project.

The sheer number of KEPs makes summarizing the progress almost tedious, but look below for the full list of work done. The SIG focused largely on DRA, as well as other initatives that
advance inference serving use-cases, but also made advancements in the cpu and memory manager.

The KEP Wrangler process, started in 2024, has continued, and some of the SIG's success and velocity can be attributed to it. It's also been an opportunity for contributors to the
SIG to step up as leaders. A special shout out goes to Sreeram Venkitesh who stepped into the role of Wrangler lead.

The SIG is happy to sponsor a new sub-project: node readiness controller, which was spun out of an effort to have a clearer signal for foundational node workloads to signal they are ready
to run end-user workloads, as well as two new working groups: WG node lifecycle and WG checkpoint restore.

2. Are there any areas and/or subprojects that your group needs help with (e.g. fewer than 2 active OWNERS)?


3. Did you have community-wide updates in 2025 (e.g. KubeCon talks)?

- Kubecon EU 2025 [maintainers track](https://kccnceu2025.sched.com/event/1tcy2/sig-node-intro-and-deep-dive-sergey-kanzhelev-google-francesco-romani-peter-hunt-red-hat)
- Kubecon NA 2025 [maintainers track](https://kccncna2025.sched.com/event/27Nla/sig-node-intro-and-deep-dive-peter-hunt-red-hat-sergey-kanzhelev-google-mrunal-patel-red-hat)

4. KEP work in 2025 (v1.33, v1.34, v1.35):
  - Alpha
    - [4188 - New kubelet gRPC API with endpoint returning local pods information](https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/4188-kubelet-pod-readiness-api) - v1.35
    - [4960 - Container Stop Signals](https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/4960-container-stop-signals) - v1.33
    - [5328 - Node Declared Features](https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/5328-node-declared-features) - v1.35
    - [5394 - PSI based Node Conditions](https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/5394-psi-node-conditions) - v1.34
    - [5419 - In-Place Pod-Level Resources Resize](https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/5419-pod-level-resources-in-place-resize) - v1.35
    - [5526 - Pod-Level Resource Managers](https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/5526-pod-level-resource-managers) - v1.35
    - [5532 - Restart All Containers on Container Exits](https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/5532-restart-all-containers-on-container-exits) - v1.35
    - [5607 - Allow HostNetwork Pods to Use User Namespaces](https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/5607-hostnetwork-userns) - v1.35

  - Beta
    - [127 - Support User Namespaces](https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/127-user-namespaces) - v1.35
    - [2033 - Rootless mode](https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/2033-kubelet-in-userns-aka-rootless) - v1.35
    - [2371 - cAdvisor-less, CRI-full Container and Pod Stats](https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/2371-cri-pod-container-stats) - v1.35
    - [2535 - Ensure Secret Pulled Images](https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/2535-ensure-secret-pulled-images) - v1.35
    - [2837 - KEP Template](https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/2837-pod-level-resource-spec) - v1.34
    - [2862 - Fine grained Kubelet API authorization](https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/2862-fine-grained-kubelet-authz) - v1.33
    - [3695 - Extend the PodResources API to include resources allocated by DRA](https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/3695-pod-resources-for-dra) - v1.34
    - [3721 - Support for env files.](https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/3721-support-for-env-files) - v1.35
    - [4205 - Expose PSI Metrics](https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/4205-psi-metric) - v1.34
    - [4265 - Add ProcMount option](https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/4265-proc-mount) - v1.33
    - [4680 - Add Resource Health Status to the Pod Status for Device Plugin and DRA](https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/4680-add-resource-health-to-pod-status) - v1.35
    - [4742 - Node Topologies via Downward API](https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/4742-node-topology-downward-api) - v1.35
    - [4800 - Split UnCoreCache Toplogy Awareness in CPU Manager](https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/4800-cpumanager-split-uncorecache) - v1.34
    - [5307 - Container Restart Policy](https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/5307-container-restart-policy) - v1.35
    - [5573 - Remove cgroup v1 support](https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/5573-remove-cgroup-v1) - v1.35
    - [5593 - Configure the max CrashLoopBackOff delay](https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/5593-configure-the-max-crashloopbackoff-delay) - v1.35

  - Stable
    - [1287 - In-place Update of Pod Resources](https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/1287-in-place-update-pod-resources) - v1.35
    - [2008 - Forensic Container Checkpointing](https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/2008-forensic-container-checkpointing) - v1.33
    - [2400 - Node system swap support](https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/2400-node-swap) - v1.34
    - [2625 - SMT aware cpumanager policy](https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/2625-cpumanager-policies-thread-placement) - v1.33
    - [2902 - CPUManager Policy Option to Distribute CPUs Across NUMA Nodes Instead of Packing Them](https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/2902-cpumanager-distribute-cpus-policy-option) - v1.35
    - [3288 - Split Stdout and Stderr Log Stream of Container](https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/3288-separate-stdout-from-stderr) - v1.34
    - [3619 - Fine grained SupplementalGroups control](https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/3619-supplemental-groups-policy) - v1.35
    - [3673 - Kubelet limit of Parallel Image Pulls](https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/3673-kubelet-parallel-image-pull-limit) - v1.35
    - [3857 - Recursive read-only mounts](https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/3857-rro-mounts) - v1.33
    - [3960 - Pod lifecycle sleep action](https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/3960-pod-lifecycle-sleep-action) - v1.34
    - [3983 - Add support for a kubelet drop-in configuration directory](https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/3983-drop-in-configuration) - v1.35
    - [4033 - Discover cgroup driver from CRI](https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/4033-group-driver-detection-over-cri) - v1.34
    - [4176 - New CPUManager Static Policy which spread hyperthreads across physical CPUs to better utilize CPU Cache](https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/4176-cpumanager-spread-cpus-preferred-policy) - v1.33
    - [4210 - ImageMaximumGCAge in Kubelet](https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/4210-max-image-gc-age) - v1.35
    - [4216 - Image pull per runtime class](https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/4216-image-pull-per-runtime-class) - v1.35
    - [4369 - Allow special characters environment variable](https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/4369-allow-special-characters-environment-variable) - v1.34
    - [4381 - DRA Structured Parameters](https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/4381-dra-structured-parameters) - v1.34
    - [4438 - Restarting sidecar containers during Pod termination](https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/4438-container-restart-termination) - v1.35
    - [4540 - Add CPUManager policy option to restrict reservedSystemCPUs to system daemons and interrupt processing](https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/4540-strict-cpu-reservation) - v1.35
    - [4622 - New TopologyManager Policy which configure the value of maxAllowableNUMANodes](https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/4622-topologymanager-max-allowable-numa-nodes) - v1.35
    - [4639 - OCI images as VolumeSource](https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/4639-oci-volume-source) - v1.35
    - [4817 - Resource Claim Status With Possible Standardized Network Interface Data](https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/4817-resource-claim-device-status) - v1.35
    - [4818 - Allow zero value for Sleep Action of PreStop Hook](https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/4818-allow-zero-value-for-sleep-action-of-prestop-hook) - v1.34
    - [5067 - Pod Generation](https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/5067-pod-generation) - v1.35
    - [753 - Sidecar Containers](https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/753-sidecar-containers) - v1.33

## [Subprojects](https://git.k8s.io/community/sig-node#subprojects)


**New in 2025:**
  - node-readiness-controller
**Continuing:**
  - ci-testing
  - cri-api
  - cri-client
  - cri-tools
  - kernel-module-management
  - kubelet
  - node-api
  - node-feature-discovery
  - node-problem-detector
  - resource-management
  - security-profiles-operator

## [Working groups](https://git.k8s.io/community/sig-node#working-groups)

**New in 2025:**
 - Checkpoint Restore
 - Node Lifecycle
**Retired in 2025:**
 - Policy
**Continuing:**
 - Batch
 - Device Management
 - Serving
 - Structured Logging

## Operational

Operational tasks in [sig-governance.md]:
- [x] [README.md] reviewed for accuracy and updated if needed
- [x] [CONTRIBUTING.md] reviewed for accuracy and updated if needed
- [x] Other contributing docs (e.g. in devel dir or contributor guide) reviewed for accuracy and updated if needed
- [x] Subprojects list and linked OWNERS files in [sigs.yaml] reviewed for accuracy and updated if needed
- [x] SIG leaders (chairs, tech leads, and subproject leads) in [sigs.yaml] are accurate and active, and updated if needed
- [x] Meeting notes and recordings for 2025 are linked from [README.md] and updated/uploaded if needed


[CONTRIBUTING.md]: https://git.k8s.io/community/sig-node/CONTRIBUTING.md
[sig-governance.md]: https://git.k8s.io/community/committee-steering/governance/sig-governance.md
[README.md]: https://git.k8s.io/community/sig-node/README.md
[sigs.yaml]: https://git.k8s.io/community/sigs.yaml
[devel]: https://git.k8s.io/community/contributors/devel/README.md
