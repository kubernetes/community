# 2021 Annual Report: SIG Multicluster

## Current initiatives

1. What work did the SIG do this year that should be highlighted?
   - [Multicluster Services](https://github.com/kubernetes/enhancements/tree/master/keps/sig-multicluster/1645-multi-cluster-services-api)
     - [Repo](https://sigs.k8s.io/mcs-api)
   - [ClusterProperty](https://github.com/kubernetes/enhancements/tree/master/keps/sig-multicluster/2149-clusterid)
     - [Repo](https://sigs.k8s.io/about-api)

2. What initiatives are you working on that aren't being tracked in KEPs?
   - Multicluster leader election has been discussed several times but is not currently associated with any KEP
   - Multicluster network policy (driven by SIG-Net): [Extending KEP-2091](https://github.com/kubernetes/enhancements/tree/master/keps/sig-network/2091-admin-network-policy)
   - Multi-network as potential MCS extension: [Doc link](https://docs.google.com/document/d/1IueF8FyGKKRPcg8Zvua84FDC327J3UdNiTMO5W7s7tQ/edit?usp=sharing)

3. KEP work in 2021 (1.x, 1.y, 1.z):

   - Alpha
     - [1645 - Multicluster Services API](https://github.com/kubernetes/enhancements/tree/master/keps/sig-multicluster/1645-multi-cluster-services-api) - Beta promotion pending Cluster ID beta
     - [2149 - Cluster ID](https://github.com/kubernetes/enhancements/tree/master/keps/sig-multicluster/2149-clusterid) - Beta promotion imminent

## Project health

1. What areas and/or subprojects does your group need the most help with?
   Any areas with 2 or fewer OWNERs? (link to more details)
   
   - [Work API](https://github.com/kubernetes-sigs/work-api): progress stalled out, we should make progress or consider archival
   - Kubefed: no progress made in the last year, considering archival

2. What metrics/community health stats does your group care about and/or measure?

   - Zoom attendance
   - Youtube meetup views
   - Complaints about late youtube uploads of meetups :)

3. Does your [CONTRIBUTING.md] help **new** contributors engage with your group specifically by pointing
   to activities or programs that provide useful context or allow easy participation?

   - One of the highest value ways to contribute to SIG multicluster is to share
     the problems folks are grappling with in the multicluster space, ideas they
     may have for how to solve them, and what they've tried so far. It's not
     necessary to want to develop solutions in the open, but sharing the
     problems is extremely valuable.

4. If your group has special training, requirements for reviewers/approvers, or processes beyond the general [contributor guide],
   does your [CONTRIBUTING.md] document those to help **existing** contributors grow throughout the [contributor ladder]?

   - N/A.

5. Does the group have contributors from multiple companies/affiliations?

   - Yes.

6. Are there ways end users/companies can contribute that they currently are not?
   If one of those ways is more full time support, what would they work on and why?

   - End users and companies can contribute high value material by sharing the
     problems they're encountering in this space. Understanding the overall
     state of play for this area is critical to being able to do high value work
     in the SIG, and the only way to develop that understanding is to receive
     input from end users. You don't need to have a solution to add value! Just
     tell us your problems.

## Membership

- Primary slack channel member count: 2885
- Primary mailing list member count: 10000 (tbd)
- Primary meeting attendee count (estimated, if needed): 20-25
- Primary meeting participant count (estimated, if needed): 5-7
- Unique reviewers for SIG-owned packages: N/A SIG-MC is out-of-tree only
- Unique approvers for SIG-owned packages: N/A SIG-MC is out-of-tree only

Include any other ways you measure group membership

- ruler

## Operational

Operational tasks in [sig-governance.md]:

- [X] [README.md] reviewed for accuracy and updated if needed
- [X] [CONTRIBUTING.md] reviewed for accuracy and updated if needed
      (or created if missing and your contributor steps and experience are different or more
      in-depth than the documentation listed in the general [contributor guide] and [devel] folder.)
- [X] Subprojects list and linked OWNERS files in [sigs.yaml] reviewed for accuracy and updated if needed
- [X] SIG leaders (chairs, tech leads, and subproject owners) in [sigs.yaml] are accurate and active, and updated if needed
- [X] Meeting notes and recordings for 2021 are linked from [README.md] and updated/uploaded if needed
- [X] Did you have community-wide updates in 2021 (e.g. community meetings, kubecon, or kubernetes-dev@ emails)? Links to email, slides, or recordings:
      - [Kubecon NA Intro / Deep Dive](https://www.youtube.com/watch?v=zVTFm7HJD3s)
      - [Kubecon EU Intro / Deep Dive](https://www.youtube.com/watch?v=nx1ABG8-uvs)
      - [Think Beyond the Cluster live panel](https://opensourcelive.withgoogle.com/events/think-beyond-the-cluster)

[CONTRIBUTING.md]: https://git.k8s.io/community/sig-multicluster/CONTRIBUTING.md
[contributor ladder]: https://git.k8s.io/community/community-membership.md
[sig-governance.md]: https://git.k8s.io/community/committee-steering/governance/sig-governance.md
[README.md]: https://git.k8s.io/community/sig-multicluster/README.md
[sigs.yaml]: https://git.k8s.io/community/sigs.yaml
[contributor guide]: https://git.k8s.io/community/contributors/guide/README.md
[devel]: https://git.k8s.io/community/contributors/devel/README.md

