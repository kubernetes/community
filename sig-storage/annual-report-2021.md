# 2021 Annual Report: SIG Storage

Note: Included KEPs targeting 1.21, 1.22, and 1.23 releases.

## Current initiatives

1. What work did the SIG do this year that should be highlighted?

   - GA milestones hit:
     - CSI Windows: ​​https://kubernetes.io/blog/2021/08/09/csi-windows-support-with-csi-proxy-reaches-ga/
     - Generic ephemeral inline volumes: https://github.com/kubernetes/enhancements/blob/master/keps/sig-storage/1698-generic-ephemeral-volumes

   - Successful progress:
     - CSI migration has been an effort that has been going on for several releases. It involves SIG Storage, SIG Cloud Provider, and contributors across many cloud providers and storage vendors to work together and move in-tree volume plugins to out-of-tree CSI drivers. https://kubernetes.io/blog/2021/12/10/storage-in-tree-to-csi-migration-status-update/

2. What initiatives are you working on that aren't being tracked in KEPs?

   - CBT (Change Blocking Tracking) is being discussed in the Data Protection WG; KEP not proposed yet.  Here’s a draft design: https://docs.google.com/document/d/1bOXazqAVAi8wtJhVsyNNyxhjWgYFzJSTFub2IxiSqMU/edit#
   - Try to get more maintainers for projects that have 1 or a small number of maintainers now

3. KEP work in 2021 (1.x, 1.y, 1.z):

<!--
In future, this will be generated from kubernetes/enhancements kep.yaml files
1. with SIG as owning-sig or in participating-sigs
2. listing 1.x, 1.y, or 1.z in milestones or in latest-milestone
-->

   - Stable
     - [1412 - Immutable Secrets and ConfigMaps](https://github.com/kubernetes/enhancements/blob/master/keps/sig-storage/1412-immutable-secrets-and-configmaps) - stable
     - [1472 - Storage Capacity Constraints for Pod Scheduling](https://github.com/kubernetes/enhancements/blob/master/keps/sig-storage/1472-storage-capacity-tracking) - stable
     - [1682 - Skip Volume Ownership Change](https://github.com/kubernetes/enhancements/blob/master/keps/sig-storage/1682-csi-driver-skip-permission) - stable
     - [1698 - generic ephemeral inline volumes](https://github.com/kubernetes/enhancements/blob/master/keps/sig-storage/1698-generic-ephemeral-volumes) - stable
     - [1855 - Service Account Token for CSI Driver](https://github.com/kubernetes/enhancements/blob/master/keps/sig-storage/1855-csi-driver-service-account-token) - stable
     - [1122 - CSI Windows](https://github.com/kubernetes/enhancements/blob/master/keps/sig-windows/1122-windows-csi-support) - stable

   - Beta
     - [1487- In-tree Storage Plugin to CSI Migration - AWS](https://github.com/kubernetes/enhancements/blob/master/keps/sig-storage/1487-csi-migration-aws/README.md) - beta
     - [1488 - In-tree Storage Plugin to CSI Migration - GCE PD](https://github.com/kubernetes/enhancements/tree/master/keps/sig-storage/1488-csi-migration-gce-pd) - beta
     - [1489 - In-tree Storage Plugin to CSI Migration - Cinder](https://github.com/kubernetes/enhancements/tree/master/keps/sig-storage/1489-csi-migration-cinder) - beta
     - [1490 - In-tree Storage Plugin to CSI Migration - Azuredisk](https://github.com/kubernetes/enhancements/blob/master/keps/sig-storage/1490-csi-migration-azuredisk) - beta
     - [1885 - In-tree Storage Plugin to CSI Migration - Azurefile](https://github.com/kubernetes/enhancements/tree/master/keps/sig-storage/1885-csi-migration-azurefile) - beta
     - [2317 - Provide fsgroup of pod to CSI driver on mount](https://github.com/kubernetes/enhancements/blob/master/keps/sig-storage/2317-fsgroup-on-mount/README.md) - beta

   - Alpha
     - [1432 - Volume Health Monitor ](https://github.com/kubernetes/enhancements/blob/master/keps/sig-storage/1432-volume-health-monitor) - alpha
     - [1790 - Recover from volume expansion failure](https://github.com/kubernetes/enhancements/blob/master/keps/sig-storage/1790-recover-resize-failure/) - alpha
     - [2485 - ReadWriteOncePod PersistentVolume AccessMode](https://github.com/kubernetes/enhancements/tree/master/keps/sig-storage/2485-read-write-once-pod-pv-access-mode) - alpha
     - [2589 - In-tree Storage Plugin to CSI Migration - Portworx](https://github.com/kubernetes/enhancements/blob/master/keps/sig-storage/2589-csi-migration-portworx) - alpha
     - [2644 - Honor Persistent Volume Reclaim Policy](https://github.com/kubernetes/enhancements/blob/master/keps/sig-storage/2644-honor-pv-reclaim-policy) - alpha
     - [2923 - In-tree Storage Plugin to CSI Migration - Ceph RBD](https://github.com/kubernetes/enhancements/blob/master/keps/sig-storage/2923-csi-migration-ceph-rbd) - alpha

   - Pre-alpha
     - Object Storage API (COSI): KEP is being updated and reviewed: https://github.com/kubernetes/enhancements/pull/2813

## Project health

1. What areas and/or subprojects does your group need the most help with?
   Any areas with 2 or fewer OWNERs? (link to more details)

   - In general, we need more contributors to help with fixing bugs and code reviewing across the board. For contributors who are interested in getting more involved in SIG-Storage, please review the [CONTRIBUTING Guide](https://github.com/kubernetes/community/blob/master/sig-storage/CONTRIBUTING.md) and come to [SIG-Storage meetings](https://github.com/kubernetes/community/tree/master/sig-storage) to find a project that you are interested in helping with. Help make code contributions and learn how storage works in Kubernetes. After you are getting familiar with the code, help with code reviews.
   - Under “Subprojects”, we added notes to a few sub-projects that are not active or need more contributors

2. What metrics/community health stats does your group care about and/or measure?

   - We have been tracking features targeting every K8s release, from design, alpha, beta, to GA in our SIG meetings. In general, this went well.
   - According to metrics on 3/20/2022, it shows the 7 day MA for a PR Time to Approve and Merge in sig-storage repository group is: https://k8s.devstats.cncf.io/d/44/pr-time-to-approve-and-merge?orgId=1&var-period=d7&var-repogroup_name=SIG%20Storage&var-apichange=All&var-size_name=All&var-kind_name=All
     - Median time from open to LGTM (in hours): Max 1.66 days, Avg 3.3 hours
     - Median time from LGTM to approve (in hours): Max 0.02 hour, Avg 0.00 hour
     - Median time from approve to merge (in hours): Max 1.77 days, Avg 2.98 hours
     - 85th percentile time from open to LGTM (in hours): Max 6.79 days, Avg 23.1 hours
     - 85th percentile time from LGTM to approve (in hours): Max 5.16 days, 1.77 hours
     - 85th percentile time from approve to merge (in hours): Max 3.16 days, 17.95 hours
   - Age of 7 day MA of issues by sig-storage repository group on 3/20/2022: https://k8s.devstats.cncf.io/d/15/issues-age-by-sig-and-repository-groups?orgId=1&var-period=d7&var-repogroup_name=SIG%20Storage&var-sig_name=All&var-kind_name=All&var-prio_name=All
     - Median time to close issue: Min 0.23 hour, Max 6.45 day, Avg 20.5 hour
     - Average number of issues opened: Min 0.14, Max 1.14, Avg 0.37

   - We have not been checking health stats regularly. This is something we can follow-up with.
   - We are tracking some bug fixes in our SIG meetings. Also PRs that are submitted for implementing a KEP are well tracked. We don’t have an overall tracking board for all PRs. Some other SIGs have PR triage boards. This is something we should explore.
     - Example: SIG Node triage board
     - Other SIG’s triage PR meeting: Categorize it and find potential reviewers

3. Does your [CONTRIBUTING.md] help **new** contributors engage with your group specifically by pointing
   to activities or programs that provide useful context or allow easy participation?

   - Our  [CONTRIBUTING.md](https://github.com/kubernetes/community/blob/master/sig-storage/CONTRIBUTING.md) has presentations, docs, and videos to help new contributors get familiar with Kubernetes Storage concepts. It also has instructions on how to get involved in SIG Storage.

4. If your group has special training, requirements for reviewers/approvers, or processes beyond the general [contributor guide],
   does your [CONTRIBUTING.md] document those to help **existing** contributors grow throughout the [contributor ladder]?

   - We don’t have special training or documents. We just follow the contributor guide available in the K8s community. Any suggestions are welcome.

5. Does the group have contributors from multiple companies/affiliations?

   - Yes.

6. Are there ways end users/companies can contribute that they currently are not?
   If one of those ways is more full time support, what would they work on and why?

   - We do have a small number of contributors from end user companies. Some are working on adding new features.  Any help to get them to contribute more will be great.
   - Here is our ask - get full time support to work on the following:
     - CSI conformance tests for CSI drivers.
     - Write more tests; monitor test grid health; work on test framework out of tree; enhance CSI release tools
     - Doc writer to improve docs on CSI side and in general in Storage side
     - Help with initial PR triage, find out who can review them

## Membership

- Primary slack channel member count: 4767 (sig-storage), 1126 (csi)
- Primary mailing list member count: 702
- Primary meeting attendee count (estimated, if needed): 25
- Primary meeting participant count (estimated, if needed): 25
- Unique reviewers for SIG-owned packages: 38 <!-- in future, this will be generated from OWNERS files referenced from subprojects, expanded with OWNERS_ALIASES files -->
- Unique approvers for SIG-owned packages: 32 <!-- in future, this will be generated from OWNERS files referenced from subprojects, expanded with OWNERS_ALIASES files -->

Include any other ways you measure group membership

## Subprojects

<!--
In future, this will be generated from delta of sigs.yaml from $YYYY-01-01 to $YYYY-12-31
Manually visible via `git diff HEAD@{$YYYY-01-01} HEAD@{$YYYY-12-31} -- $sig-id/README.md`
-->

New in 2021:
- https://github.com/kubernetes-csi/csi-driver-nvmf
  - Note: Need more contributors.

Deprecated in 2019:
- https://github.com/kubernetes-csi/cluster-driver-registrar

Deprecated in 2018:
- https://github.com/kubernetes-csi/driver-registrar

Continuing:
  - name: external-storage
    - https://github.com/kubernetes-sigs/sig-storage-lib-external-provisioner
    - https://github.com/kubernetes-sigs/sig-storage-local-static-provisioner
  - name: git-sync
    - https://github.com/kubernetes/git-sync

  - name: gluster-provisioner
    - https://github.com/kubernetes-sigs/gluster-block-external-provisioner
      - Note: Last commit: 5/28/2020, no release from the project; not active; Should we archive this or just keep it as is?
      - Conclusion: Sent [email](https://groups.google.com/g/kubernetes-sig-storage/c/QvKc25Gdi84) to SIG-Storage mailing list and got response that this project is still needed. We’ll keep this.

    - https://github.com/kubernetes-sigs/gluster-file-external-provisioner
      - Note: There were a couple of commits from pohly related to release-tools. The last bug fix commit was 7/30/2020; not active; Should we archive this or just keep it as is?
Conclusion: Sent [email](https://groups.google.com/g/kubernetes-sig-storage/c/QvKc25Gdi84) to SIG-Storage mailing list and got response that this project is still needed. We’ll keep this.

  - name: kubernetes-cosi
    - https://githubuser.com/kubernetes-sigs/container-object-storage-interface-api
    - https://githubuser.com/kubernetes-sigs/container-object-storage-interface-controller
    - https://github.com/kubernetes-sigs/container-object-storage-interface-csi-adapter
    - https://github.com/kubernetes-sigs/container-object-storage-interface-provisioner-sidecar
    - https://github.com/kubernetes-sigs/container-object-storage-interface-spec
    - https://github.com/kubernetes-sigs/cosi-driver-sample/master

  - name: kubernetes-csi
    - https://github.com/kubernetes-csi/cluster-driver-registrar (deprecated in 2019)
    - https://github.com/kubernetes-csi/csi-driver-host-path
    - https://github.com/kubernetes-csi/csi-driver-image-populator
      - Note: No activities;
      - Conclusion: Sent [email](https://groups.google.com/g/kubernetes-sig-storage/c/QvKc25Gdi84) to SIG-Storage mailing list, but did not get response that this project is still needed. Will decide to archive this in the next SIG-Storage meeting.

    - https://github.com/kubernetes-csi/csi-driver-iscsi
    - https://github.com/kubernetes-csi/csi-driver-nfs
    - https://github.com/kubernetes-csi/csi-driver-nvmf
      - Note: Mainly maintained by 1 person. More contributors needed.

    - https://github.com/kubernetes-csi/csi-driver-smb
    - https://github.com/kubernetes-csi/csi-lib-fc
      - Note: The first and last real commit in this repo was 10/19/2018. No real activities after that.
      - Action: Should we archive this?
      - Conclusion: Sent [email](https://groups.google.com/g/kubernetes-sig-storage/c/QvKc25Gdi84) to SIG-Storage mailing list, but did not get response that this project is still needed. Will decide to archive this in the next SIG-Storage meeting.

    - https://github.com/kubernetes-csi/csi-lib-iscsi
    - https://github.com/kubernetes-csi/csi-lib-utils
    - https://github.com/kubernetes-csi/csi-proxy
    - https://github.com/kubernetes-csi/csi-release-tools
    - https://github.com/kubernetes-csi/csi-test
    - https://github.com/kubernetes-csi/docs
    - https://github.com/kubernetes-csi/driver-registrar (deprecated in 2018)
    - https://github.com/kubernetes-csi/external-attacher
    - https://github.com/kubernetes-csi/external-health-monitor
    - https://github.com/kubernetes-csi/external-provisioner
    - https://github.com/kubernetes-csi/external-resizer
    - https://github.com/kubernetes-csi/external-snapshotter
    - https://github.com/kubernetes-csi/kubernetes-csi.github.io
    - https://github.com/kubernetes-csi/livenessprobe
    - https://github.com/kubernetes-csi/node-driver-registrar
    - https://github.com/kubernetes/csi-api
      - Note: CSINode and CSIDriver are core APIs now and they are GA. Should we archive this repo?
      - Sent out [email](https://groups.google.com/g/kubernetes-sig-storage/c/QvKc25Gdi84) to SIG Storage mailing list.
      - Conclusion: Archive it

    - https://github.com/kubernetes/csi-translation-lib

    - https://github.com/kubernetes/kubernetes/staging/src/k8s.io/csi-api/
      - Note: Do we still need to list this as a sub-project under SIG-Storage?
      - Sent out [email](https://groups.google.com/g/kubernetes-sig-storage/c/QvKc25Gdi84) to SIG Storage mailing list.
      - Conclusion: Archive it

    - https://github.com/kubernetes/kubernetes/staging/src/k8s.io/csi-translation-lib

  - name: mount-utils
    - https://github.com/kubernetes/kubernetes/staging/src/k8s.io/mount-utils

  - name: nfs-provisioner
    - https://github.com/kubernetes-sigs/nfs-ganesha-server-and-external-provisioner
    - https://github.com/kubernetes-sigs/nfs-subdir-external-provisioner/
    - Note: These 2 projects need more maintainers to help. May consider deprecation in the future and ask folks to move to csi-driver-nfs.

  - name: volume-populators
    owners:
    - https://github.com/kubernetes-csi/lib-volume-populator
    - https://github.com/kubernetes-csi/volume-data-source-validator
  - name: volumes
    owners:
    - https://github.com/kubernetes/kubernetes/pkg/volume

## Working groups

<!--
In future, this will be generated from delta of sigs.yaml from $YYYY-01-01 to $YYYY-12-31
Manually visible via `git diff HEAD@{$YYYY-01-01} HEAD@{$YYYY-12-31} -- $sig-id/README.md`
-->

Continuing:
- [Kubernetes Data Protection WG](https://github.com/kubernetes/community/tree/master/wg-data-protection)
  - [2020 report](https://github.com/kubernetes/community/blob/master/wg-data-protection/annual-report-2020.md)
  - [2021 report](https://github.com/kubernetes/community/blob/master/wg-data-protection/annual-report-2021.md)
  - [Data Protection White Paper](https://github.com/kubernetes/community/blob/master/wg-data-protection/data-protection-workflows-white-paper.md)

## Operational

Operational tasks in [sig-governance.md]:

- [x] [README.md] reviewed for accuracy and updated if needed
- [x] [CONTRIBUTING.md] reviewed for accuracy and updated if needed
      (or created if missing and your contributor steps and experience are different or more
      in-depth than the documentation listed in the general [contributor guide] and [devel] folder.)
- [x] Subprojects list and linked OWNERS files in [sigs.yaml] reviewed for accuracy and updated if needed
- [x] SIG leaders (chairs, tech leads, and subproject owners) in [sigs.yaml] are accurate and active, and updated if needed
- [x] Meeting notes and recordings for 2021 are linked from [README.md] and updated/uploaded if needed
- [x] Did you have community-wide updates in 2021 (e.g. community meetings, kubecon, or kubernetes-dev@ emails)? Links to email, slides, or recordings:
  - Kubernetes SIG Storage Introduction and Update, KubeCon NA 2021:
https://www.youtube.com/watch?v=-r8t24qVvUA&t=3s
  - Kubernetes SIG Storage Introduction and Update, KubeCon China 2021: https://www.youtube.com/watch?v=H5P1miRL8vc
  - Kubernetes SIG Storage Introduction and Update, KubeCon Europe 2021: https://www.youtube.com/watch?v=dedCB6kPJlc&t=3s

[CONTRIBUTING.md]: https://git.k8s.io/community/sig-storage/CONTRIBUTING.md
[contributor ladder]: https://git.k8s.io/community/community-membership.md
[sig-governance.md]: https://git.k8s.io/community/committee-steering/governance/sig-governance.md
[README.md]: https://git.k8s.io/community/sig-storage/README.md
[sigs.yaml]: https://git.k8s.io/community/sigs.yaml
[contributor guide]: https://git.k8s.io/community/contributors/guide/README.md
[devel]: https://git.k8s.io/community/contributors/devel/README.md
