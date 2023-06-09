# 2022 Annual Report: WG Data Protection

## Current initiatives

1. What work did the WG do this year that should be highlighted?
   For example, artifacts, reports, white papers produced this year.

  - KEPs being worked on:
    - Restrict volume mode conversion to fix security issue moved to Alpha in 1.24: https://kubernetes.io/blog/2022/05/18/prevent-unauthorised-volume-mode-conversion-alpha/
    - Volume Populator moved to Beta in 1.24: https://kubernetes.io/blog/2022/05/16/volume-populators-beta/
    - Object Storage API (COSI) moved Alpha in 1.25: https://kubernetes.io/blog/2022/09/02/cosi-kubernetes-object-storage-management/
    - Volume Group KEP was being reviewed: https://github.com/kubernetes/enhancements/pull/1551
    - Changed Block Tracking KEP was submitted and being reviewed: https://github.com/kubernetes/enhancements/pull/3367

2. What initiatives are you working on that aren't being tracked in KEPs?

   We discussed the following topics in addition to KEPs:
   - VolumeSnapshot v1beta1 removal discussion: https://docs.google.com/document/d/15tLCV3csvjHbKb16DVk-mfUmFry_Rlwo-2uG6KNGsfw/edit#heading=h.qb01pf8tp27
   - Volume replication: https://docs.google.com/document/d/15tLCV3csvjHbKb16DVk-mfUmFry_Rlwo-2uG6KNGsfw/edit#heading=h.6luxmtw9ljkm

## Project health

1. What's the current roadmap until completion of the working group?

   - We have identified the missing building blocks for supporting Data Protection in Kubernetes in our white paper: https://github.com/kubernetes/community/blob/master/wg-data-protection/data-protection-workflows-white-paper.md#what-are-the-missing-building-blocks-in-kubernetes. Features such as Volume Backups, Change Block Tracking, Volume Populator, Volume Group Snapshot, and Backup Repositories are owned by SIG Storage. Features such as Quiesce and Unquiesce Hooks are owned by SIG Node, with SIG Storage and SIG Apps participating. Features such as Application Snapshots and Backups are owned by SIG Apps, with SIG Storage participating. We will continue to work on them until all the missing pieces are available in Kubernetes.

2. Does the group have contributors from multiple companies/affiliations?

   - Yes. In our [agenda doc](https://docs.google.com/document/d/15tLCV3csvjHbKb16DVk-mfUmFry_Rlwo-2uG6KNGsfw/edit#), we listed companies who are supporting the WG. Here are some of the companies that actively participating in the WG: Cohesity, Dell EMC, Druva, Google, Infinidat, Kasten by Veeam, Microsoft, Mirantis, NetApp, Red Hat, Seagate, Trilio, Veritas, and VMware.

3. Are there ways end users/companies can contribute that they currently are not?
   If one of those ways is more full time support, what would they work on and why?

   - There were a couple of end users who participated in the meeting.
   - If more end users/companies can attend our meetings, provide feedback, and contribute to design/implementation of the features we are working on, that will be great.

## Membership

- Primary slack channel member count: 247
- Primary mailing list member count: 208
- Primary meeting attendee count (estimated, if needed): 10
- Primary meeting participant count (estimated, if needed): 10

Include any other ways you measure group membership

## Operational

Operational tasks in [wg-governance.md]:

- [x] [README.md] reviewed for accuracy and updated if needed
- [x] WG leaders in [sigs.yaml] are accurate and active, and updated if needed
- [x] Meeting notes and recordings for 2022 are linked from [README.md] and updated/uploaded if needed
- [x] Updates provided to sponsoring SIGs in 2022
      - [sig-storage](https://git.k8s.io/community/sig-storage)
        - WG does not own code. We discussed the features such as CBT, Volume Populator in the WG meeting, but we still get the features tracked in SIG Storage and they are in the tracking spreadsheet. This means we’ll give an update on them in every SIG Storage bi-weekly meeting: https://docs.google.com/spreadsheets/d/1t4z5DYKjX2ZDlkTpCnp18icRAQqOE85C1T1r2gqJVck/edit#gid=111388712
      - KubeCon Updates:
        - Data Protection WG Deep Dive at KubeCon North America 2022: https://sched.co/182MZ
        - Data Protection WG Deep Dive at KubeCon Europe 2022: https://sched.co/ytnj

[wg-governance.md]: https://git.k8s.io/community/committee-steering/governance/wg-governance.md
[README.md]: https://git.k8s.io/community/wg-data-protection/README.md
[sigs.yaml]: https://git.k8s.io/community/sigs.yaml
