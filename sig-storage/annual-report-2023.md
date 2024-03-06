# 2023 Annual Report: SIG Storage

## Current initiatives and Project Health

1. What work did the SIG do this year that should be highlighted?

<!--
   Some example items that might be worth highlighting:
   - Major KEP advancement
   - Important initiatives that aren't tracked via KEPs
   - Paying down significant tech debt
   - Governance and leadership changes
-->

2. Are there any areas and/or subprojects that your group needs help with (e.g. fewer than 2 active OWNERS)?

<!--
   Note: This list is generated from the KEP metadata in kubernetes/enhancements repository.
      If you find any discrepancy in the generated list here, please check the KEP metadata.
      Please raise an issue in kubernetes/community, if the KEP metadata is correct but the generated list is incorrect.
-->

3. Did you have community-wide updates in 2023 (e.g. KubeCon talks)?

<!--
  Examples include links to email, slides, or recordings.
-->

4. KEP work in 2023 (v1.27, v1.28, v1.29):

  - Alpha
    - [3751 - Kubernetes Volume Provisioned IO](https://github.com/kubernetes/enhancements/tree/master/keps/sig-storage/3751-volume-attributes-class) - v1.29

  - Beta
    - [1710 - Skip SELinux relabeling of volumes](https://github.com/kubernetes/enhancements/tree/master/keps/sig-storage/1710-selinux-relabeling) - v1.27
    - [2923 - In-tree Storage Plugin to CSI Migration - Ceph RBD](https://github.com/kubernetes/enhancements/tree/master/keps/sig-storage/2923-csi-migration-ceph-rbd) - v1.27
    - [3141 - Prevent unauthorised volume mode conversion](https://github.com/kubernetes/enhancements/tree/master/keps/sig-storage/3141-prevent-volume-mode-conversion) - v1.27
    - [3476 - Volume Group Snapshot](https://github.com/kubernetes/enhancements/tree/master/keps/sig-storage/3476-volume-group-snapshot) - v1.29
    - [3756 - Robust VolumeManager reconstruction after kubelet restart](https://github.com/kubernetes/enhancements/tree/master/keps/sig-storage/3756-volume-reconstruction) - v1.27
    - [3762 - PersistentVolume last phase transition time](https://github.com/kubernetes/enhancements/tree/master/keps/sig-storage/3762-persistent-volume-last-phase-transition-time) - v1.29

  - Stable
    - [2268 - non graceful shutdown](https://github.com/kubernetes/enhancements/tree/master/keps/sig-storage/2268-non-graceful-shutdown) - v1.28
    - [2485 - ReadWriteOncePod PersistentVolume AccessMode](https://github.com/kubernetes/enhancements/tree/master/keps/sig-storage/2485-read-write-once-pod-pv-access-mode) - v1.29
    - [2644 - Honor Persistent Volume Reclaim Policy](https://github.com/kubernetes/enhancements/tree/master/keps/sig-storage/2644-honor-pv-reclaim-policy) - v1.28
    - [3107 - SecretRef field addition to NodeExpandVolume request](https://github.com/kubernetes/enhancements/tree/master/keps/sig-storage/3107-csi-nodeexpandsecret) - v1.29
    - [3294 - Provision volumes from cross-namespace snapshots](https://github.com/kubernetes/enhancements/tree/master/keps/sig-storage/3294-provision-volumes-from-cross-namespace-snapshots) - v1.29
    - [3333 - Retroactive default StorageClass assignment](https://github.com/kubernetes/enhancements/tree/master/keps/sig-storage/3333-retroactive-default-storage-class) - v1.28

## [Subprojects](https://git.k8s.io/community/sig-storage#subprojects)


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

**Retired in 2023:**
 - Multitenancy
**Continuing:**
 - Data Protection
 - Policy
 - Structured Logging

## Operational

Operational tasks in [sig-governance.md]:
- [ ] [README.md] reviewed for accuracy and updated if needed
- [ ] [CONTRIBUTING.md] reviewed for accuracy and updated if needed
- [ ] Other contributing docs (e.g. in devel dir or contributor guide) reviewed for accuracy and updated if needed
- [ ] Subprojects list and linked OWNERS files in [sigs.yaml] reviewed for accuracy and updated if needed
- [ ] SIG leaders (chairs, tech leads, and subproject leads) in [sigs.yaml] are accurate and active, and updated if needed
- [ ] Meeting notes and recordings for 2023 are linked from [README.md] and updated/uploaded if needed


[CONTRIBUTING.md]: https://git.k8s.io/community/sig-storage/CONTRIBUTING.md
[sig-governance.md]: https://git.k8s.io/community/committee-steering/governance/sig-governance.md
[README.md]: https://git.k8s.io/community/sig-storage/README.md
[sigs.yaml]: https://git.k8s.io/community/sigs.yaml
[devel]: https://git.k8s.io/community/contributors/devel/README.md
