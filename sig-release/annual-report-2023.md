# 2023 Annual Report: SIG Release

## Current initiatives and Project Health

1. What work did the SIG do this year that should be highlighted?

In 2023, the SIG made significant progress on moving toward only leveraging community infrastructure. We now leverage [OpenBuildService](https://openbuildservice.org) to [build and publish Debian and RPM packages to pkgs.k8s.io](https://kubernetes.io/blog/2023/08/15/pkgs-k8s-io-introduction/). 

Additionally, work was done in order to keep supported Kubernetes versions up-to-date with supported Go versions. This has helped improve security of Kubernetes patch releases by building and releasing using supported minor versions of go. 

<!--
   Some example items that might be worth highlighting:
   - Major KEP advancement
   - Important initiatives that aren't tracked via KEPs
   - Paying down significant tech debt
   - Governance and leadership changes
-->

2. Are there any areas and/or subprojects that your group needs help with (e.g. fewer than 2 active OWNERS)?

Not at this time.

3. Did you have community-wide updates in 2023 (e.g. KubeCon talks)?

- How SIG Release Makes Kubernetes Releases Even More Stable and Secure - Veronica Lopez & Marko Mudrinić
  - https://sched.co/1HyUA
- Releasing Kubernetes and Beyond: Flexible and Fast Delivery of Packages - Grace Nguyen, Adolfo Garcia Veytia, Jim Angel, John Anderson
  - https://sched.co/1RQeC      

4. KEP work in 2023 (v1.27, v1.28, v1.29):

  - Alpha
    - [1731 - Publishing Kubernetes packages on community infrastructure](https://github.com/kubernetes/enhancements/tree/master/keps/sig-release/1731-publishing-packages) - v1.28


  - Stable
    - [3000 - Artifact Distribution Policy](https://github.com/kubernetes/enhancements/tree/master/keps/sig-release/3000-artifact-distribution) - v1.27
    - [3720 - Freezing `k8s.gcr.io` image registry](https://github.com/kubernetes/enhancements/tree/master/keps/sig-release/3720-freezing-k8s-gcr-io) - v1.27
    - [3744 - Stay on supported go versions](https://github.com/kubernetes/enhancements/tree/master/keps/sig-release/3744-stay-on-supported-go-versions) - v1.27

## [Subprojects](https://git.k8s.io/community/sig-release#subprojects)


**Continuing:**
  - Release Engineering
      - New subproject lead: Marko Mudrinić
  - Release Team
      - New subproject leads: Kat Cosgrove and Grace Nguyen
  - SIG Release Process Documentation

## [Working groups](https://git.k8s.io/community/sig-release#working-groups)

**New in 2023:**
 - LTS
**Retired in 2023:**
 - Reliability

## Operational

Operational tasks in [sig-governance.md]:
- [x] [README.md] reviewed for accuracy and updated if needed
- [x] [CONTRIBUTING.md] reviewed for accuracy and updated if needed
- [x] Other contributing docs (e.g. in devel dir or contributor guide) reviewed for accuracy and updated if needed
- [x] Subprojects list and linked OWNERS files in [sigs.yaml] reviewed for accuracy and updated if needed
- [x] SIG leaders (chairs, tech leads, and subproject leads) in [sigs.yaml] are accurate and active, and updated if needed
- [x] Meeting notes and recordings for 2023 are linked from [README.md] and updated/uploaded if needed


[CONTRIBUTING.md]: https://git.k8s.io/community/sig-release/CONTRIBUTING.md
[sig-governance.md]: https://git.k8s.io/community/committee-steering/governance/sig-governance.md
[README.md]: https://git.k8s.io/community/sig-release/README.md
[sigs.yaml]: https://git.k8s.io/community/sigs.yaml
[devel]: https://git.k8s.io/community/contributors/devel/README.md
