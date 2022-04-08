# SIG Security - 2021 annual report

## Current initiatives

**1. What work did the SIG do this year that should be highlighted?**

- At the time of the 2020 Annual Reports, SIG Security was a very new SIG, just establishing who we are as a group and starting to build a space for security to flourish as a community effort. In 2021, we have come into our own. Since the last annual report, SIG Security has doubled in size, established and matured subprojects that have become welcoming communities of their own, and grown new leaders who have come to lift up more new contributors themselves. We now have generations of contributors working, learning, and honking together.

  Our [Docs subproject](https://kubernetes.slack.com/archives/C01D8R7ACQ2) has made Kubernetes security more approachable through blogs and tutorials such as [A Closer Look at NSA/CISA Kubernetes Hardening Guidance](https://kubernetes.io/blog/2021/10/05/nsa-cisa-kubernetes-hardening-guidance/), [Securing Admission Controllers](https://kubernetes.io/blog/2022/01/19/secure-your-admission-controllers-and-webhooks/), [Apply Pod Security Standards at the Cluster Level](https://kubernetes.io/docs/tutorials/security/cluster-level-pss/), and [Apply Pod Security Standards at the Namespace Level](https://kubernetes.io/docs/tutorials/security/ns-level-pss/).

  Our [Tooling subproject](https://kubernetes.slack.com/archives/C01CUSVMHPY) got off the ground in 2021 and is now flying high, [standardizing Kubernetes CVE reporting](https://github.com/kubernetes/sig-security/issues/1), [prototyping vulnerability scanning of Kubernetes artifacts](https://github.com/kubernetes/sig-security/issues/3), and [hosting](https://www.youtube.com/watch?v=zB1-7NLsfps) [learning](https://www.youtube.com/watch?v=o-E6aoKmznY) [sessions](https://www.youtube.com/watch?v=MWAb63gf3gs) which have been viewed hundreds of times.

  SIG Security does things differently than a lot of other SIGs. We don’t own code. We work hard to cultivate and grow an inclusive, welcoming environment where everyone, from the newest contributor to the most experienced hacker, can come share their ideas and contribute to the security of the project. This is designed to ripple outward whether or not we are involved, encouraging and facilitating internal collaboration across SIGs and engagement with the external security community. We’re really proud of how far we’ve come, and we’re excited to continue leading by example into 2022 and beyond!

**2. What initiatives are you working on that aren't being tracked in KEPs?**

- Most of our work is outside the scope of KEPs. Improving processes, providing ongoing services, and actively building collaboration within the Kubernetes project are all core to what we do. We operate on the basis of consent, communication, and participation–what gets done is based on who shows up, what ideas they bring, and how the rest of the community responds.

  [Security Self-assessments](https://github.com/kubernetes/sig-security/issues/2) are a great example of how this works in practice. ClusterAPI maintainers were looking for guidance on evaluating and improving the security of ClusterAPI, and SIG Security members were enthusiastic about helping them. We collaborated with CNCF TAG Security to begin a guided self-assessment exercise for ClusterAPI, adapting TAG Security’s existing self-assessment workflow into a Kubernetes context. Through this process, the participants have gained confidence and experience, word has gotten around, and other Kubernetes subprojects have started asking about their own guided self-assessments. At the time of this writing, Security Self-Assessment is emerging as a likely new SIG Security subproject.

  Our [Tooling subproject](https://kubernetes.slack.com/archives/C01CUSVMHPY) is working to prototype, define policy for, and implement [vulnerability scanning for build time dependencies and container images](https://github.com/kubernetes/sig-security/issues/3). This initiative involves collaborating with several other SIGs, applying SIG Security’s specialized experience with these tools to produce meaningful improvements without burying other SIGs in additional work.

  Our [Third-Party Audit subproject](https://github.com/kubernetes/sig-security/tree/main/sig-security-external-audit) continues to administer their namesake initiative. In 2021 they published a [Request for Proposals](https://github.com/kubernetes/sig-security/blob/main/sig-security-external-audit/security-audit-2021/RFP.md), evaluated the vendor responses, and began to negotiate contract terms with the selected vendor. As of this writing, they have [announced the vendor selection](https://github.com/kubernetes/sig-security/blob/main/sig-security-external-audit/security-audit-2021/RFP_Decision.md) and are making arrangements for work to begin.

  Our [Security Docs subproject](https://kubernetes.slack.com/archives/C01D8R7ACQ2) works in collaboration with SIG Docs and others on its continuing initiatives: to improve security representation and content in Kubernetes documentation, to communicate information to help end users and others keep themselves safer, and to help foster and grow new contributors. They have been prolifically publishing useful information such as the blog posts shouted out in the beginning of this report, and highlighting good first issues for new contributors to start diving in.

**3. KEP work in 2021 (1.21, 1.22, 1.23):**

   - Stable
      - [1933 - Defend against logging secrets via static analysis](https://git.k8s.io/enhancements/keps/sig-security/1933-secret-logging-static-analysis/README.md) - 1.23
   - Beta
      - [2579 - PSP Replacement Policy](https://git.k8s.io/enhancements/keps/sig-auth/2579-psp-replacement/README.m) - 1.23
      - [1933 - Defend against logging secrets via static analysis](https://git.k8s.io/enhancements/keps/sig-security/1933-secret-logging-static-analysis/README.md) - 1.21
   - Alpha
       - [2579 - PSP Replacement Policy](https://git.k8s.io/enhancements/keps/sig-auth/2579-psp-replacement/README.md) - 1.22
   - Pre-alpha
      - [2763 - Ambient Capabilities](https://git.k8s.io/enhancements/keps/sig-security/2763-ambient-capabilities/README.md)

## Project health

**1. What areas and/or subprojects does your group need the most help with?**

- Our [Docs subproject](https://kubernetes.slack.com/archives/C01D8R7ACQ2) is always looking for security-minded contributors of all experience levels to share their learning and knowledge with the community! This subproject has consistently been a place where people [merge](https://github.com/kubernetes/website/pull/28248) their [first](https://github.com/kubernetes/website/pull/31518) Kubernetes PRs.There’s always room for continuous improvement in our documentation, and contributing to this provides an opportunity to learn more about Kubernetes security while helping everyone run their clusters more safely. We’re really proud of the way Docs encourages and welcomes new contributors, and we’d love to encourage you to become a part of it!

**Any areas with 2 or fewer OWNERs? (link to more details)**

- The nature of our work means we have little need for OWNERS files, compared to SIGs that ship code. We have two or more OWNERS in all areas, and our community focus helps ensure that when someone needs to scale back their involvement they will be well supported.

**2. What metrics/community health stats does your group care about and/or measure?**

- We care about our people. We measure community health by looking at community participation, and seeing how people are feeling. Are people coming to meetings and talking to one another on Slack? How many people are engaged? Is anyone dominating the conversation? Are new contributors becoming involved? Are people coming back? 


  These types of concerns are notably hard to metricize, which we remain comfortable with.

**3. Does your [CONTRIBUTING.md] help _new_ contributors engage with your group specifically by pointing to activities or programs that provide useful context or allow easy participation?**


- We have known for some time that SIG Security needs a CONTRIBUTING.MD document that more accurately reflects how we do things. SIG Security takes a community-building approach to improving Kubernetes security: we intentionally make a space where someone can begin contributing by sharing their thoughts and ideas in a meeting or Slack message. The CONTRIBUTING.MD document we have in progress reflects this. It encourages participation by showing up and links to the general Kubernetes contributor guides, which are still applicable.

**4. If your group has special training, requirements for reviewers/approvers, or processes beyond the general [contributor guide], does your [CONTRIBUTING.md] document those to help _existing_ contributors grow throughout the [contributor ladder]?**

- SIG Security does not have special training or requirements for reviewers/approvers. The nature of our cross-SIG work means that we have many fewer PRs, OWNERS files, and similar artifacts than many SIGs. The Kubernetes Contributor Ladder document as currently written focuses on scaling review of numerous code PRs. This limits its immediate applicability to our horizontal SIG, which does not own code and is intentionally built upon open, non-hierarchical community participation. We continue to learn through experience how best to utilize and evolve the contributor ladder to encourage and recognize our leading contributors.

**5. Does the group have contributors from multiple companies/affiliations?**

   - Yes.

**6. Are there ways end users/companies can contribute that they currently are not? If one of those ways is more full time support, what would they work on and why?**

   - End users and people who work at end user companies are contributing broadly within SIG Security today. There are so many possibilities of ways to contribute. As a community we continue to experiment and discover new ways, together. Bring your gifts; come share and learn with us!

## Membership

- Primary slack channel member count: 890
- Subproject slack channel member counts:
  - #sig-security-docs: 260
  - #sig-security-tooling: 357
- Primary mailing list member count: 246
- Primary meeting attendee count (estimated, if needed): 13.8
- Primary meeting participant count (estimated, if needed): 13.8

Due to the way we conduct our meetings, at the majority of meetings, everyone participates.

- Unique reviewers for SIG-owned packages: N/A
- Unique approvers for SIG-owned packages: N/A

SIG Security does not own packages.

## Subprojects

New in 2021:
- [security-tooling](https://git.k8s.io/community/sig-security#security-tooling)

Retired in 2021:
- none

Continuing:
- [security-audit](https://git.k8s.io/community/sig-security#$security-audit)
- [security-docs](https://git.k8s.io/community/sig-security#$security-docs)

## Working groups

New in 2021:
- none

Retired in 2021:
- none

Continuing:
- none

## Operational

Operational tasks in [sig-governance.md]:

- [x] [README.md] reviewed for accuracy and updated if needed
- [ ] [CONTRIBUTING.md] reviewed for accuracy and updated if needed
    (or created if missing and your contributor steps and experience are different or more
    in-depth than the documentation listed in the general [contributor guide] and [devel] folder.)
- [x] Subprojects list and linked OWNERS files in [sigs.yaml] reviewed for accuracy and updated if needed
- [x] SIG leaders (chairs, tech leads, and subproject owners) in [sigs.yaml] are accurate and active, and updated if needed
- [x] Meeting notes and recordings for 2021 are linked from [README.md] and updated/uploaded if needed
- [x] Did you have community-wide updates in 2021 (e.g. community meetings, kubecon, or kubernetes-dev@ emails)? Links to email, slides, or recordings:
    - KubeCon EU: [Get In Containerds, We’re Going Securing: Kubernetes SIG Security is Here!](https://www.youtube.com/watch?v=0_s6zkyRpME)
    - KubeCon NA: [Security Through Transparency: Kubernetes SIG Security Update](https://www.youtube.com/watch?v=O5Wy7zSigOU)

[CONTRIBUTING.md]: https://git.k8s.io/community/sig-security/CONTRIBUTING.md
[contributor ladder]: https://git.k8s.io/community/community-membership.md
[sig-governance.md]: https://git.k8s.io/community/committee-steering/governance/sig-governance.md
[README.md]: https://git.k8s.io/community/sig-security/README.md
[sigs.yaml]: https://git.k8s.io/community/sigs.yaml
[contributor guide]: https://git.k8s.io/community/contributors/guide/README.md
[devel]: https://git.k8s.io/community/contributors/devel/README.md
