# 2023 Annual Report: SIG Docs

## Current initiatives and Project Health

1. What work did the SIG do this year that should be highlighted?

   - SIG Docs completely [removed the Katacoda tutorials catalog](https://github.com/kubernetes/website/discussions/38878), previously shared as new user resources, after this content was reclaimed behind a paywall for O'Reilly subscribers. We are still looking for a replacement set of resources
   - The [Issue Wrangler role](https://github.com/kubernetes/website/discussions/38861) was formalized in the SIG, which introduced a new named role to help triage and categorize SIG Docs issues. This role follows the roster system currently in place for PR Wranglers, but opens up the possibility of seasoned contributors who *aren't* approvers to additionally contribute to the SIG
   - SIG Docs worked with SIG Release and the Release Team Subproject to introduce new and improved processes for documentation during each release. This includes initiating a [Docs Freeze](https://github.com/kubernetes/sig-release/pull/2383) as part of the release calendar deadlines, creating [better handbook guidance](https://github.com/kubernetes/sig-release/pull/2350) for weekly status updates, as well as a more robust process for removing shadows/leads who fail to fulfil their role (seen specifically across Docs/Comms for the last few releases)
   - The New Contributor Ambassador has been replaced, after Arsh Sharma served in this role for over two years. We thank Arsh and welcome our now-current New Contributor Ambassador, Sreeram Venkitesh!
   - SIG Docs served as one of the first SIGs to setup and trial the new YouTube automation project lead by SIG Contributor Experience for all community meeting uploads
   - SIG Docs leadership created [onboarding and offboarding documentation and processes](https://github.com/kubernetes/website/discussions/38842) for Co-chairs and Tech Leads, alongside an annualized list of activities and work the SIG Docs leadership team is responsible for
   - We have initiated the start of a possible Translation Platform test for our localization teams, with more work on this initiative continuing in 2024

2. Are there any areas and/or subprojects that your group needs help with (e.g. fewer than 2 active OWNERS)?

   - The Kubernetes Blog Subproject currently only has one active editor, as we attempt to onboard further people to the blog reviewing team. This is a strain both on the usual blog work we have, as well as feature blogs linked with each Kubernetes release
   - The Reference Docs Subproject is currently lacking true ownership and process outlines, which we hope to solve via recruiting further Tech Leads for the SIG (currently in progress). At present, this subproject is understaffed
   - SIG Docs is currently in a state of leadership transition: three Tech Leads will step down in 2024, while we have actively recruited/replaced two of these positions

3. Did you have community-wide updates in 2023 (e.g. KubeCon talks)?

   - KubeCon/CloudNativeCon EU 2023 Maintainer Track: [How To Localize Kubernetes Documentation – A Guide For Everyone](https://youtu.be/Ng8QwklOwe0)
   - KubeCon/CloudNativeCon NA 2023 Maintainer Track: [Docs Drive Adoption – Help Wanted!](https://www.youtube.com/watch?v=R4RDZ-rLNJo)

4. KEP work in 2023 (v1.27, v1.28, v1.29):

   * N/A


## [Subprojects](https://git.k8s.io/community/sig-docs#subprojects)


**Continuing:**
  - kubernetes-blog
  - localization
  - reference-docs
  - website

## [Working groups](https://git.k8s.io/community/sig-docs#working-groups)

* N/A

## Operational

Operational tasks in [sig-governance.md]:
- [ ] [README.md] reviewed for accuracy and updated if needed
- [ ] [CONTRIBUTING.md] reviewed for accuracy and updated if needed
- [ ] Other contributing docs (e.g. in devel dir or contributor guide) reviewed for accuracy and updated if needed
- [x] Subprojects list and linked OWNERS files in [sigs.yaml] reviewed for accuracy and updated if needed
- [ ] SIG leaders (chairs, tech leads, and subproject leads) in [sigs.yaml] are accurate and active, and updated if needed
- [x] Meeting notes and recordings for 2023 are linked from [README.md] and updated/uploaded if needed


[CONTRIBUTING.md]: https://git.k8s.io/community/sig-docs/CONTRIBUTING.md
[sig-governance.md]: https://git.k8s.io/community/committee-steering/governance/sig-governance.md
[README.md]: https://git.k8s.io/community/sig-docs/README.md
[sigs.yaml]: https://git.k8s.io/community/sigs.yaml
[devel]: https://git.k8s.io/community/contributors/devel/README.md
