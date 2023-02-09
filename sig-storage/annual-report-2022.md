# 2022 Annual Report: SIG Storage

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
    - [1432 - Volume Health Monitor](https://github.com/kubernetes/enhancements/tree/master/keps/sig-storage/1432-volume-health-monitor) - v1.24
    - [1979 - Object Storage Support](https://github.com/kubernetes/enhancements/tree/master/keps/sig-storage/1979-object-storage-support) - v1.25
    - [2644 - Honor Persistent Volume Reclaim Policy](https://github.com/kubernetes/enhancements/tree/master/keps/sig-storage/2644-honor-pv-reclaim-policy) - v1.26
    - [2924 - In-tree Storage Plugin to CSI Migration - Ceph Cephfs](https://github.com/kubernetes/enhancements/tree/master/keps/sig-storage/2924-csi-migration-cephfs) - v1.26
    - [3107 - SecretRef field addition to NodeExpandVolume request](https://github.com/kubernetes/enhancements/tree/master/keps/sig-storage/3107-csi-nodeexpandsecret) - v1.25
    - [3294 - Provision volumes from cross-namespace snapshots](https://github.com/kubernetes/enhancements/tree/master/keps/sig-storage/3294-provision-volumes-from-cross-namespace-snapshots) - v1.26
  - beta:
    - [2268 - non graceful shutdown](https://github.com/kubernetes/enhancements/tree/master/keps/sig-storage/2268-non-graceful-shutdown) - v1.26
    - [2589 - In-tree Storage Plugin to CSI Migration - Portworx](https://github.com/kubernetes/enhancements/tree/master/keps/sig-storage/2589-csi-migration-portworx) - v1.25
    - [2923 - In-tree Storage Plugin to CSI Migration - Ceph RBD](https://github.com/kubernetes/enhancements/tree/master/keps/sig-storage/2923-csi-migration-ceph-rbd) - v1.26
    - [3333 - Retroactive default StorageClass assignment](https://github.com/kubernetes/enhancements/tree/master/keps/sig-storage/3333-reconcile-default-storage-class) - v1.26
  - stable:
    - [1472 - Storage Capacity Constraints for Pod Scheduling](https://github.com/kubernetes/enhancements/tree/master/keps/sig-storage/1472-storage-capacity-tracking) - v1.24
    - [1487 - In-tree Storage Plugin to CSI Migration - AWS](https://github.com/kubernetes/enhancements/tree/master/keps/sig-storage/1487-csi-migration-aws) - v1.25
    - [1488 - In-tree Storage Plugin to CSI Migration - GCE PD](https://github.com/kubernetes/enhancements/tree/master/keps/sig-storage/1488-csi-migration-gce-pd) - v1.25
    - [1489 - In-tree Storage Plugin to CSI Migration - Cinder](https://github.com/kubernetes/enhancements/tree/master/keps/sig-storage/1489-csi-migration-cinder) - v1.24
    - [1490 - In-tree Storage Plugin to CSI Migration - Azuredisk](https://github.com/kubernetes/enhancements/tree/master/keps/sig-storage/1490-csi-migration-azuredisk) - v1.24
    - [1491 - In-tree Storage Plugin to CSI Migration - vSphere](https://github.com/kubernetes/enhancements/tree/master/keps/sig-storage/1491-csi-migration-vsphere) - v1.26
    - [1885 - In-tree Storage Plugin to CSI Migration - Azurefile](https://github.com/kubernetes/enhancements/tree/master/keps/sig-storage/1885-csi-migration-azurefile) - v1.26
    - [2317 - Provide fsgroup of pod to CSI driver on mount](https://github.com/kubernetes/enhancements/tree/master/keps/sig-storage/2317-fsgroup-on-mount) - v1.26
    - [284 - Growing Persistent Volume size](https://github.com/kubernetes/enhancements/tree/master/keps/sig-storage/284-enable-volume-expansion) - v1.24
    - [361 - Local Ephemeral Storage Capacity Isolation](https://github.com/kubernetes/enhancements/tree/master/keps/sig-storage/361-local-ephemeral-storage-isolation) - v1.25
    - [596 - Ephemeral Inline CSI Volumes](https://github.com/kubernetes/enhancements/tree/master/keps/sig-storage/596-csi-inline-volumes) - v1.25
    - [625 - In-tree Storage Plugin to CSI Migration](https://github.com/kubernetes/enhancements/tree/master/keps/sig-storage/625-csi-migration) - v1.25


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


**Continuing:**

 - Data Protection
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

[CONTRIBUTING.md]: https://git.k8s.io/community/sig-storage/CONTRIBUTING.md
[contributor ladder]: https://git.k8s.io/community/community-membership.md
[sig-governance.md]: https://git.k8s.io/community/committee-steering/governance/sig-governance.md
[README.md]: https://git.k8s.io/community/sig-storage/README.md
[sigs.yaml]: https://git.k8s.io/community/sigs.yaml
[contributor guide]: https://git.k8s.io/community/contributors/guide/README.md
[devel]: https://git.k8s.io/community/contributors/devel/README.md
