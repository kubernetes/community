# 2025 Annual Report: SIG Storage

## Current initiatives and Project Health

1. What work did the SIG do this year that should be highlighted?

<!--
   Some example items that might be worth highlighting:
   - Major KEP advancement
   - Important initiatives that aren't tracked via KEPs
   - Paying down significant tech debt
   - Governance and leadership changes
-->

GA milestones hit

- [1495 - Volume Populators](https://github.com/kubernetes/enhancements/issues/1495) moved to GA in v1.33. This feature allows for populating volumes with data from an arbitrary data source.
- [2589 - Portworx in-tree to CSI driver migration](https://github.com/kubernetes/enhancements/issues/2589) moved to GA in v1.33. This feature migrates the Portworx in-tree storage plugin to CSI.
- [2644 - Honor Persistent Volume Reclaim Policy](https://github.com/kubernetes/enhancements/issues/2644) moved to GA in v1.33. This feature ensures that the reclaim policy of a PersistentVolume is honored when the PV is released.
- [1790 - Recover from Volume Expansion Failure](https://github.com/kubernetes/enhancements/issues/1790) moved to GA in v1.34. This feature allows users to recover from volume expansion failure by retrying with a smaller size.
- [3751 - Kubernetes VolumeAttributesClass ModifyVolume](https://github.com/kubernetes/enhancements/issues/3751) moved to GA in v1.34. This feature extends the Kubernetes Persistent Volume API to allow users to dynamically modify volume options (such as IOPs and throughput), after a volume is provisioned.

Successful progress

- [3476 - Volume Group Snapshot](https://github.com/kubernetes/enhancements/issues/3476) moved to v1beta2 in v1.34. The feature introduces a VolumeGroupSnapshot API to take a snapshot of multiple volumes together.
- [4876 - Mutable CSINode Allocatable Property](https://github.com/kubernetes/enhancements/issues/4876) moved to Beta in v1.34. This feature allows the CSINode allocatable property to be mutable.
- [5538 - CSI driver opt-in for service account tokens via secrets field](https://github.com/kubernetes/enhancements/issues/5538) moved to Beta in v1.35. This feature allows CSI drivers to opt-in for service account tokens via the secrets field.

2. Are there any areas and/or subprojects that your group needs help with (e.g. fewer than 2 active OWNERS)?

- Write more tests; monitor test grid health; work on test
framework out of tree; enhance CSI release tools.
- Doc writer to improve docs on CSI side and in general in
Storage side.
- We would appreciate help with more efficient issue triage and help with fixing the issues.
  - https://github.com/orgs/kubernetes-csi/projects/52/views/3

3. Did you have community-wide updates in 2025 (e.g. KubeCon talks)?

<!--
  Examples include links to email, slides, or recordings.
-->

- KubeCon NA 2025 SIG Storage Intro & Deep Dive: https://www.youtube.com/watch?v=tGSEyEdh5ug
- KubeCon EU Project Lightning Talk: https://www.youtube.com/watch?v=JhMwxRfi7Ho
- KubeCon EU 2025 SIG Storage Intro & Deep Dive: https://www.youtube.com/watch?v=X_xHC_Q5jGE
- KubeCon EU Project Lightning Talk: https://www.youtube.com/watch?v=v_PzG81D33I

4. KEP work in 2025 (v1.33, v1.34, v1.35):
<!--
   TODO: Uncomment the following auto-generated list of KEPs, once reviewed & updated for correction.

   Note: This list is generated from the KEP metadata in kubernetes/enhancements repository.
      If you find any discrepancy in the generated list here, please check the KEP metadata.
      Please raise an issue in kubernetes/community, if the KEP metadata is correct but the generated list is incorrect.
-->

  - Alpha
    - [3314 - CSI Changed Block Tracking](https://github.com/kubernetes/enhancements/tree/master/keps/sig-storage/3314-csi-changed-block-tracking) - v1.33
    - [4049 - Storage Capacity Scoring of Nodes for Dynamic Provisioning](https://github.com/kubernetes/enhancements/tree/master/keps/sig-storage/4049-storage-capacity-scoring-of-nodes-for-dynamic-provisioning) - v1.33
    - [5040 - Remove gitRepo volumes driver](https://github.com/kubernetes/enhancements/tree/master/keps/sig-storage/5040-remove-gitrepo-driver) - v1.33
    - [4958 - CSI Sidecars All in one](https://github.com/kubernetes/enhancements/tree/master/keps/sig-storage/4958-csi-sidecars-all-in-one) - v1.34
    - [5381 - Mutable PersistentVolume Node Affinity](https://github.com/kubernetes/enhancements/tree/master/keps/sig-storage/5381-mutable-pv-affinity) - v1.35

  - Beta
    - [4876 - Mutable CSINode Allocatable Property](https://github.com/kubernetes/enhancements/tree/master/keps/sig-storage/4876-mutable-csinode-allocatable) - v1.34
    - [5538 - CSI driver opt-in for service account tokens via secrets field](https://github.com/kubernetes/enhancements/tree/master/keps/sig-storage/5538-csi-sa-tokens-secrets-field) - v1.35

  - Stable
    - [1495 - Volume Populators](https://github.com/kubernetes/enhancements/tree/master/keps/sig-storage/1495-volume-populators) - v1.33
    - [2589 - In-tree Storage Plugin to CSI Migration - Portworx](https://github.com/kubernetes/enhancements/tree/master/keps/sig-storage/2589-csi-migration-portworx) - v1.33
    - [2644 - Honor Persistent Volume Reclaim Policy](https://github.com/kubernetes/enhancements/tree/master/keps/sig-storage/2644-honor-pv-reclaim-policy) - v1.33
    - [1790 - Recover from volume expansion failure](https://github.com/kubernetes/enhancements/tree/master/keps/sig-storage/1790-recover-resize-failure) - v1.34
    - [3751 - Kubernetes VolumeAttributesClass and ModifyVolume](https://github.com/kubernetes/enhancements/tree/master/keps/sig-storage/3751-volume-attributes-class) - v1.34

## [Subprojects](https://git.k8s.io/community/sig-storage#subprojects)

**New in 2025:**
  - None
**Continuing:**
  - external-snapshot-metadata
  - external-storage
  - git-sync
  - gluster-provisioner
  - kubernetes-cosi
  - kubernetes-csi
  - mount-utils
  - nfs-provisioner
  - volume-populators
  - volumes

## [Working groups](https://git.k8s.io/community/sig-storage#working-groups)

**New in 2025:**
 - Node Lifecycle
**Retired in 2025:**
 - Policy
**Continuing:**
 - Data Protection
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


[CONTRIBUTING.md]: https://git.k8s.io/community/sig-storage/CONTRIBUTING.md
[sig-governance.md]: https://git.k8s.io/community/committee-steering/governance/sig-governance.md
[README.md]: https://git.k8s.io/community/sig-storage/README.md
[sigs.yaml]: https://git.k8s.io/community/sigs.yaml
[devel]: https://git.k8s.io/community/contributors/devel/README.md
