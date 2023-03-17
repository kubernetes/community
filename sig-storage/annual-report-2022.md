# 2022 Annual Report: SIG Storage

## Current initiatives

1. What work did the SIG do this year that should be highlighted?


   - GA milestones hit
     - [1472 - Storage Capacity Constraints for Pod Scheduling](https://github.com/kubernetes/enhancements/tree/master/keps/sig-storage/1472-storage-capacity-tracking) - 1.24 blog: https://kubernetes.io/blog/2022/05/06/storage-capacity-ga/

     - [2317 - Provide fsgroup of pod to CSI driver on mount](https://github.com/kubernetes/enhancements/tree/master/keps/sig-storage/2317-fsgroup-on-mount) - v1.26 blog: https://kubernetes.io/blog/2022/12/23/kubernetes-12-06-fsgroup-on-mount/

     - [284 - Growing Persistent Volume size](https://github.com/kubernetes/enhancements/tree/master/keps/sig-storage/284-enable-volume-expansion) - v1.24 blog: https://kubernetes.io/blog/2022/05/05/volume-expansion-ga/

     - [361 - Local Ephemeral Storage Capacity Isolation](https://github.com/kubernetes/enhancements/tree/master/keps/sig-storage/361-local-ephemeral-storage-isolation) - v1.25 blog: https://kubernetes.io/blog/2022/09/19/local-storage-capacity-isolation-ga/

     - [596 - Ephemeral Inline CSI Volumes](https://github.com/kubernetes/enhancements/tree/master/keps/sig-storage/596-csi-inline-volumes) - v1.25 blog: https://kubernetes.io/blog/2022/08/29/csi-inline-volumes-ga/

   - Successful progress
     - We continue to make great progress on CSI Migration. Core CSI Migration and CSI Migration for plugins including AWS, GCE PD, OpenStack Cinder, Azure Disk, Azure File, and vSphere all moved to GA in 2022.

2. What initiatives are you working on that aren't being tracked in KEPs?

   - We started a SIG-Storage Issue Triage board and weekly issue triage meeting
     - https://github.com/orgs/kubernetes-csi/projects/52/views/3
     - https://docs.google.com/document/d/1n-dXXvCbHsPfO1yrKwT1qoC80KhsxHYKbRdChdzqeXY/edit#

2. What initiatives are you working on that aren't being tracked in KEPs?

   - We started a SIG-Storage Issue Triage board and weekly issue triage meeting
     - https://github.com/orgs/kubernetes-csi/projects/52/views/3
     - https://docs.google.com/document/d/1n-dXXvCbHsPfO1yrKwT1qoC80KhsxHYKbRdChdzqeXY/edit#

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

   - We need more contributors to help with fixing bugs and code reviewing across the board. For contributors who are interested in getting more involved in SIG-Storage, please review the [CONTRIBUTING Guide](https://github.com/kubernetes/community/blob/master/sig-storage/CONTRIBUTING.md) and come to [SIG-Storage meetings](https://github.com/kubernetes/community/tree/master/sig-storage) to find a project that you are interested in helping with. Help make code contributions and learn how storage works in Kubernetes. After you are getting familiar with the code, help with code reviews.
   - We now have a weekly [issue triage meeting](https://docs.google.com/document/d/1n-dXXvCbHsPfO1yrKwT1qoC80KhsxHYKbRdChdzqeXY/edit#) to discuss issues:  https://github.com/orgs/kubernetes-csi/projects/52/views/3. We hope more people can join this meeting and contribute.

2. What metrics/community health stats does your group care about and/or measure?

   - We have been tracking features targeting every K8s release, from design, alpha, beta, to GA in our SIG meetings.
   - According to metrics on 3/10/2023, it shows the 7 day MA for a PR Time to Approve and Merge in sig-storage repository group is:
https://k8s.devstats.cncf.io/d/44/pr-time-to-approve-and-merge?orgId=1&var-period=d7&var-repogroup_name=SIG%20Storage&var-apichange=All&var-size_name=All&var-kind_name=All
     - Median time from open to LGTM (in hours): Max 3.05 weeks, Avg 1.37 days
     - Median time from LGTM to approve (in hours): Max 8.25 hours, Avg 0.18 hours
     - Median time from approve to merge (in hours): Max 3.29 days, Avg 0.74 hours
     - 85th percentile time from open to LGTM (in hours): Max 13.49 weeks, Avg 1.13 weeks
     - 85th percentile time from LGTM to approve (in hours): Max 3.09 weeks, 10.03 hours
     - 85th percentile time from approve to merge (in hours): Max 2.65 weeks, 1.26 hours
   - Age of 7 day MA of issues by sig-storage repository group on 3/10/2023: https://k8s.devstats.cncf.io/d/15/issues-age-by-sig-and-repository-groups?orgId=1&var-period=d7&var-repogroup_name=SIG%20Storage&var-sig_name=All&var-kind_name=All&var-prio_name=All
     - Median time to close issue: Min 0.25 hours, Max 30.65 weeks, Avg 3.61 weeks
     - Average number of issues opened: Min 0.14, Max 3.14, Avg 0.88

   - We now have a weekly [issue triage meeting](https://docs.google.com/document/d/1n-dXXvCbHsPfO1yrKwT1qoC80KhsxHYKbRdChdzqeXY/edit#) to triage issues:  https://github.com/orgs/kubernetes-csi/projects/52/views/3. We hope more people can join this meeting and contribute.

3. Does your [CONTRIBUTING.md] help **new** contributors engage with your group specifically by pointing
   to activities or programs that provide useful context or allow easy participation?

   - Our [CONTRIBUTING.md](https://github.com/kubernetes/community/blob/master/sig-storage/CONTRIBUTING.md) has presentations, docs, and videos to help new contributors get familiar with Kubernetes Storage concepts. It also has instructions on how to get involved in SIG Storage.


4. If your group has special training, requirements for reviewers/approvers, or processes beyond the general [contributor guide],
   does your [CONTRIBUTING.md] document those to help **existing** contributors grow throughout the [contributor ladder]?

   - We don’t have special training or documents. We just follow the contributor guide available in the K8s community. Any suggestions are welcome.


5. Does the group have contributors from multiple companies/affiliations?

   - Yes.

6. Are there ways end users/companies can contribute that they currently are not?
   If one of those ways is more full time support, what would they work on and why?

   - We do have a small number of contributors from end user companies. Some are working on adding new features.  Any help to get them to contribute more will be great.
   - Here is our ask:
     - CSI conformance tests for CSI drivers.
     - Write more tests; monitor test grid health; work on test framework out of tree; enhance CSI release tools.
     - Doc writer to improve docs on CSI side and in general in Storage side.
     - We have a weekly issue triage meeting, but would appreciate help with more efficient issue triage.


## Membership

- Primary slack channel member count: 5373 (sig-storage), 1399 (csi)
- Primary mailing list member count: 749
- Primary meeting attendee count (estimated, if needed): 25
- Primary meeting participant count (estimated, if needed): 25
- Unique reviewers for SIG-owned packages: 36 <!-- in future, this will be generated from OWNERS files referenced from subprojects, expanded with OWNERS_ALIASES files -->
- Unique approvers for SIG-owned packages: 30 <!-- in future, this will be generated from OWNERS files referenced from subprojects, expanded with OWNERS_ALIASES files -->

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
   - KEP related work done by the DP WG is tracked in our project tracking spreadsheet and discussed in SIG Storage meetings.

 - Multitenancy
   - We don’t have regular communication with Multi Tenancy WG.

 - Policy
   - We don’t have regular communication with Policy WG.

 - Structured Logging
   - Communication is through Structured Logging PR submission and reviews.


## Operational

Operational tasks in [sig-governance.md]:

- [x] [README.md] reviewed for accuracy and updated if needed
- [x] [CONTRIBUTING.md] reviewed for accuracy and updated if needed
      (or created if missing and your contributor steps and experience are different or more
      in-depth than the documentation listed in the general [contributor guide] and [devel] folder.)
- [x] Subprojects list and linked OWNERS files in [sigs.yaml] reviewed for accuracy and updated if needed
- [x] SIG leaders (chairs, tech leads, and subproject owners) in [sigs.yaml] are accurate and active, and updated if needed
- [x] Meeting notes and recordings for 2022 are linked from [README.md] and updated/uploaded if needed
- [x] Did you have community-wide updates in 2022 (e.g. community meetings, kubecon, or kubernetes-dev@ emails)? Links to email, slides, or recordings:
     - Kubernetes SIG Storage Deep Dive, KubeCon Europe 2022: https://www.youtube.com/watch?v=dsEeQqRSg74
     - Kubernetes SIG Storage Deep Dive, KubeCon NA 2022: https://www.youtube.com/watch?v=_XXn3-yDZA0


[CONTRIBUTING.md]: https://git.k8s.io/community/sig-storage/CONTRIBUTING.md
[contributor ladder]: https://git.k8s.io/community/community-membership.md
[sig-governance.md]: https://git.k8s.io/community/committee-steering/governance/sig-governance.md
[README.md]: https://git.k8s.io/community/sig-storage/README.md
[sigs.yaml]: https://git.k8s.io/community/sigs.yaml
[contributor guide]: https://git.k8s.io/community/contributors/guide/README.md
[devel]: https://git.k8s.io/community/contributors/devel/README.md
