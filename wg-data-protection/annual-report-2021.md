# 2021 Annual Report: WG Data Protection

Note: Only include KEPs targeting 1.21, 1.22, and 1.23 releases.

## Current initiatives

1. What work did the WG do this year that should be highlighted?
   For example, artifacts, reports, white papers produced this year.

- [Data Protection White Paper](https://github.com/kubernetes/community/blob/master/wg-data-protection/data-protection-workflows-white-paper.md)
- KEPs being worked on:
  - Restrict volume mode conversion to fix security issue: Had discussions in the WG and submitted KEP to introduce it as an Alpha feature: https://github.com/kubernetes/enhancements/pull/3151
  - Working on moving Volume Populator to Beta: https://github.com/kubernetes/enhancements/pull/2934
  - Object Storage API (COSI): KEP is being updated and reviewed: https://github.com/kubernetes/enhancements/pull/2813
  - ContainerNotifier KEP review in progress: https://github.com/kubernetes/enhancements/pull/1995/
  - Volume Group KEP review in progress: https://github.com/kubernetes/enhancements/pull/1551

2. What initiatives are you working on that aren't being tracked in KEPs?

We discussed the following topics in the WG that are not tracked in KEPs:
- Change Block Tracking (CBT) API design: https://docs.google.com/document/d/15tLCV3csvjHbKb16DVk-mfUmFry_Rlwo-2uG6KNGsfw/edit#heading=h.1olwavha9frv
- Volume Replication: https://docs.google.com/document/d/15tLCV3csvjHbKb16DVk-mfUmFry_Rlwo-2uG6KNGsfw/edit#heading=h.pah3yke9ddug
- Backup and restore externally managed services:
  - https://docs.google.com/document/d/15tLCV3csvjHbKb16DVk-mfUmFry_Rlwo-2uG6KNGsfw/edit#heading=h.7zwanijz69u1
  - https://docs.google.com/presentation/d/1IM6d0w3CDdHv1dLaFNXEcxy5fuDTr9LERAdMVkZiK9s/edit#slide=id.p
- Snapshot policy (immutable snapshot): https://docs.google.com/document/d/15tLCV3csvjHbKb16DVk-mfUmFry_Rlwo-2uG6KNGsfw/edit#heading=h.gb8m7t8jro1v
- Securing S3 Backups against Ransomware: https://sched.co/igUT
- Volume Snapshot GA phases: https://docs.google.com/document/d/15tLCV3csvjHbKb16DVk-mfUmFry_Rlwo-2uG6KNGsfw/edit#heading=h.w8v8tpkuw8ac
- Kubernetes Data Protection with Velero: https://docs.google.com/document/d/15tLCV3csvjHbKb16DVk-mfUmFry_Rlwo-2uG6KNGsfw/edit#heading=h.iekhl8nl58lo

## Project health

1. What's the current roadmap until completion of the working group?

We have identified the missing building blocks for supporting Data Protection in Kubernetes in our white paper: https://github.com/kubernetes/community/blob/master/wg-data-protection/data-protection-workflows-white-paper.md#what-are-the-missing-building-blocks-in-kubernetes. Features such as Volume Backups, Change Block Tracking, Volume Populator, Volume Group and Group Snapshot, and Backup Repositories are owned by SIG Storage. Features such as Quiesce and Unquiesce Hooks are owned by SIG Node, with SIG Storage and SIG Apps participating. Features such as Application Snapshots and Backups are owned by SIG Apps, with SIG Storage participating. We will continue to work on them until all the missing pieces are available in Kubernetes.

2. Does the group have contributors from multiple companies/affiliations?

Yes. In our [agenda doc](https://docs.google.com/document/d/15tLCV3csvjHbKb16DVk-mfUmFry_Rlwo-2uG6KNGsfw/edit#), we listed companies who are supporting the WG. Here are some of the companies that actively participating in the WG: Cohesity, Dell EMC, Druva, Google, Infinidat, Kasten by Veeam, Microsoft, Mirantis, NetApp, Red Hat, Seagate, Trilio, Veritas, and VMware.

3. Are there ways end users/companies can contribute that they currently are not?
   If one of those ways is more full time support, what would they work on and why?

- There were a couple of end users who participated in the meeting and brought up topics such as how to backup and restore externally managed services.
- If more end users/companies can attend our meetings, provide feedback, and contribute to design/implementation of the features we are working on, that will be great.

## Membership

- Primary slack channel member count: 196
- Primary mailing list member count: 193
- Primary meeting attendee count (estimated, if needed): 25
- Primary meeting participant count (estimated, if needed): 25

Include any other ways you measure group membership

## Operational

Operational tasks in [wg-governance.md]:

- [x] [README.md] reviewed for accuracy and updated if needed
- [x] WG leaders in [sigs.yaml] are accurate and active, and updated if needed
- [x] Meeting notes and recordings for 2021 are linked from [README.md] and updated/uploaded if needed
- [x] Updates provided to sponsoring SIGs in 2021
      - [sig-storage](https://git.k8s.io/community/sig-storage)
        - links to email, meeting notes, slides, or recordings, etc
        - WG does not own code. We discussed the features such as CBT, Volume Populator in the WG meeting, but we still get the features tracked in SIG Storage and they are in the tracking spreadsheet. This means we’ll give an update on them in every SIG Storage bi-weekly meeting:  https://docs.google.com/spreadsheets/d/1t4z5DYKjX2ZDlkTpCnp18icRAQqOE85C1T1r2gqJVck/edit#gid=705655608
      - [$sig-name](https://git.k8s.io/community/$sig-id/)
        - links to email, meeting notes, slides, or recordings, etc
- KubeCon updates
  - KubeCon Europe 2021: https://www.youtube.com/watch?v=DBxOBzBkimo
  - KubeCon NA 2021: https://www.youtube.com/watch?v=pTDWiHmpEz8&t=4s
  - KubeCon China 2021: https://www.youtube.com/watch?v=rHRhmi76Q4I

[wg-governance.md]: https://git.k8s.io/community/committee-steering/governance/wg-governance.md
[README.md]: https://git.k8s.io/community/wg-data-protection/README.md
[sigs.yaml]: https://git.k8s.io/community/sigs.yaml
