# 2024 Annual Report: SIG Storage

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

- [3756 - Robust VolumeManager Reconstruction](https://github.com/kubernetes/enhancements/issues/3756) moved to GA in v1.30. This feature improves Kubelet handling of mounted volumes across restarts, enabling faster, more robust recovery.
- [3141 - Prevent Unauthorized Volume Mode Conversion](https://github.com/kubernetes/enhancements/issues/3141) moved to GA in v1.30. This feature prevents unauthorized conversion of the volume mode when creating a Persistent Volume Claim from a Volume Snapshot.
- [3762 - Persistent Volume last phase transition time](https://github.com/kubernetes/enhancements/issues/3762) moved to GA in v1.31. This feature adds a PersistentVolumeStatus field which holds a timestamp of when a PersistentVolume last transitioned to a different phase. This allows you to measure time e.g. between a PV Pending and Bound. This can be also useful for providing metrics and SLOs.
- [1847 - Auto remove PVCs created by StatefulSet](https://github.com/kubernetes/enhancements/issues/1847) moved to GA in v1.32. This feature, co-owned with SIG-Apps, allows the PVCs created by StatefulSet to be automatically deleted when the volumes are no longer in use to ease management of StatefulSets that don't live indefinitely.

Successful progress

- [3751 - Kubernetes VolumeAttributesClass ModifyVolume](https://github.com/kubernetes/enhancements/issues/3751) moved to Beta in v1.31. This feature extends the Kubernetes Persistent Volume API to allow users to dynamically modify volume options (such as IOPs and throughput), after a volume is provisioned. (Note: The current status of this enhancement is marked as at risk for code freeze.)
- [1790 - Recover from Volume Expansion Failure](https://github.com/kubernetes/enhancements/issues/1790) moved to Beta  in v1.32. This feature allows users to recover from volume expansion failure by retrying with a smaller size.
- [3476 - Volume Group Snapshot](https://github.com/kubernetes/enhancements/issues/3476) moved to Beta in v1.32. The feature introduces a VolumeGroupSnapshot API to take a snapshot of multiple volumes together.

2. Are there any areas and/or subprojects that your group needs help with (e.g. fewer than 2 active OWNERS)?

- Write more tests; monitor test grid health; work on test
framework out of tree; enhance CSI release tools.
- Doc writer to improve docs on CSI side and in general in
Storage side.
- We have a weekly issue triage meeting, but would apprecia
te help with more efficient issue triage and help with fixi
ng the issues.
  - https://github.com/orgs/kubernetes-csi/projects/52/views/3

3. Did you have community-wide updates in 2024 (e.g. KubeCon talks)?

<!--
  Examples include links to email, slides, or recordings.
-->

- KubeCon NA update 2024: https://www.youtube.com/watch?v=DkpQSCX6KqQ 
- KubeCon EU update 2024: https://www.youtube.com/watch?v=pnNTd4VlWFQ 

4. KEP work in 2024 (v1.30, v1.31, v1.32):
<!--
   TODO: Uncomment the following auto-generated list of KEPs, once reviewed & updated for correction.

   Note: This list is generated from the KEP metadata in kubernetes/enhancements repository.
      If you find any discrepancy in the generated list here, please check the KEP metadata.
      Please raise an issue in kubernetes/community, if the KEP metadata is correct but the generated list is incorrect.
-->

  - Alpha
    - [1710 - Speed up recursive SELinux label change](https://github.com/kubernetes/enhancements/tree/master/keps/sig-storage/1710-selinux-relabeling)
      - SELinuxMount moved to Alpha in v1.30
      - SELinuxChangePolicy moved to Alpha in v1.32

  - Beta
    - [1790 - Recover from volume expansion failure](https://github.com/kubernetes/enhancements/tree/master/keps/sig-storage/1790-recover-resize-failure) - v1.32
    - [2644 - Honor Persistent Volume Reclaim Policy](https://github.com/kubernetes/enhancements/tree/master/keps/sig-storage/2644-honor-pv-reclaim-policy) - v1.31
    - [3476 - Volume Group Snapshot](https://github.com/kubernetes/enhancements/tree/master/keps/sig-storage/3476-volume-group-snapshot) - v1.32
    - [3751 - VolumeAttributesClass ModifyVolume](https://github.com/kubernetes/enhancements/tree/master/keps/sig-storage/3751-volume-attributes-class) - v1.31
    - [2589 - Portworx in-tree to CSI driver migration](https://github.com/kubernetes/enhancements/issues/2589) - on-by-default in v1.31

  - Stable
    - [3141 - Prevent unauthorised volume mode conversion](https://github.com/kubernetes/enhancements/tree/master/keps/sig-storage/3141-prevent-volume-mode-conversion) - v1.30
    - [3756 - Robust VolumeManager reconstruction after kubelet restart](https://github.com/kubernetes/enhancements/tree/master/keps/sig-storage/3756-volume-reconstruction) - v1.30
    - [3762 - PersistentVolume last phase transition time](https://github.com/kubernetes/enhancements/tree/master/keps/sig-storage/3762-persistent-volume-last-phase-transition-time) - v1.31
    - [1847 - Auto remove PVCs created by StatefulSet](https://github.com/kubernetes/enhancements/issues/1847) - v1.32 (co-owned with SIG-Apps)

## [Subprojects](https://git.k8s.io/community/sig-storage#subprojects)


**New in 2024:**
  - external-snapshot-metadata
**Continuing:**
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

**New in 2024:**
 - Serving
**Continuing:**
 - Data Protection
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


[CONTRIBUTING.md]: https://git.k8s.io/community/sig-storage/CONTRIBUTING.md
[sig-governance.md]: https://git.k8s.io/community/committee-steering/governance/sig-governance.md
[README.md]: https://git.k8s.io/community/sig-storage/README.md
[sigs.yaml]: https://git.k8s.io/community/sigs.yaml
[devel]: https://git.k8s.io/community/contributors/devel/README.md
