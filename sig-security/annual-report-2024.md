# 2024 Annual Report: SIG Security

## Current Initiatives and Project Health

1. What work did the SIG do this year that should be highlighted?

SIG Security has been cultivating new generations of contributors and as we continue to grow together, we have brought in new leadership across the SIG and its subprojects. In 2024 Cailyn Edwards joined us as co-chair, Iain Smart has been leading the third-party audit with Rey Lejano, Rory McCune has been leading SIG Security Docs along with Savitha Raghunathan, and Mahé Tardy and Eric Smalling joined us as new shadow leads for SIG Security Tooling!

In 2024, SIG Security’s Third-Party Audit subproject started the process for the latest comprehensive third-party security audit of the Kubernetes project, putting out the RFP, choosing the vendor, and kicking everything off. The new audit is now in progress, and will be completed in 2025!

SIG Security Docs published the [Application Security Checklist](https://kubernetes.io/docs/concepts/security/application-security-checklist/), a dedicated resource to help developers deploy applications securely on Kubernetes. This helps make Kubernetes security more accessible, since most Kubernetes security advice is written for cluster administrators rather than application developers.

SIG Security Tooling has adopted [cve-feed-osv](https://github.com/kubernetes-sigs/cve-feed-osv/), a set of tools to generate OSV-format documentation for CVEs issued by Kubernetes. These tools help end-users get fewer false-positive vulnerability scanner results, by ensuring higher quality detections in their scanning tools. In the future, these tools may become part of the [official CVE feed](https://kubernetes.io/docs/reference/issues-security/official-cve-feed/).

2. Are there any areas and/or subprojects that your group needs help with (e.g. fewer than 2 active OWNERS)?

SIG Security Docs has lots of projects to help contribute to! We pride ourselves on providing an inclusive and welcoming environment for contributors of all experience levels, so this is a great place to plug in for new and experienced contributors alike. If you’re security minded and looking to contribute to Kubernetes, come sit with us!

3. Did you have community-wide updates in 2024 (e.g. KubeCon talks)?

SIG Security did a maintainer track talk at KubeCon EU 2024: [SIG Security Update:Growing Together](https://www.youtube.com/watch?v=4TYjaI0tBBM)

By consensus of the leads, SIG Security did not have any official representation at KubeCon NA 2024 in Salt Lake City due to Utah’s anti-transgender laws. We were very disappointed in the decision to hold the conference in a place that was unsafe for our leadership and the people we love, and we are very much looking forward to seeing everyone in Europe in 2025!

4. KEP work in 2024 (v1.30, v1.31, v1.32):

In 2024 SIG Security [removed SecurityContextDeny](https://github.com/kubernetes/enhancements/issues/3785), a feature that was so old it predated the KEP process! This feature had been deprecated for many years, but because it was still in the code it still showed up in compliance frameworks such as the CIS Benchmarks. Now users in regulated industries will have less deviation request paperwork to handle, and the project will be more secure as a whole!

## [Subprojects](https://git.k8s.io/community/sig-security#subprojects)

**Continuing:**  
  - security-audit  
  - security-docs  
  - security-tooling  
  - sig-security

**Sunset:**  
  - security-assessments

In 2024 we archived the Security Self-Assessments subproject, which was formerly led by Ala Dewberry.

Ala led Self-Assessments with skill, warmth, and curiosity, helping so many contributors learn about the security of their own projects and empower themselves via workshops, documentation, and facilitation of self-assessment processes.

Thanks also go to Pushkar Joglekar, for leading the first self-assessment, starting the subproject, and mentoring Ala into this leadership role!

The documentation and artifacts from the Self-Assessments subproject will remain available under https://github.com/kubernetes/sig-security/ for future reference.

## [Working groups](https://git.k8s.io/community/sig-security#working-groups)

**Continuing:**  
 - LTS

## Operational

Operational tasks in [sig-governance.md]:  
- [x] [README.md] reviewed for accuracy and updated if needed  
- [x] [CONTRIBUTING.md] reviewed for accuracy and updated if needed  
- [x] Other contributing docs (e.g. in devel dir or contributor guide) reviewed for accuracy and updated if needed  
- [x] Subprojects list and linked OWNERS files in [sigs.yaml] reviewed for accuracy and updated if needed  
- [x] SIG leaders (chairs, tech leads, and subproject leads) in [sigs.yaml] are accurate and active, and updated if needed  
- [x] Meeting notes and recordings for 2024 are linked from [README.md] and updated/uploaded if needed

[CONTRIBUTING.md]: https://git.k8s.io/community/sig-security/CONTRIBUTING.md  
[sig-governance.md]: https://git.k8s.io/community/committee-steering/governance/sig-governance.md  
[README.md]: https://git.k8s.io/community/sig-security/README.md  
[sigs.yaml]: https://git.k8s.io/community/sigs.yaml  
[devel]: https://git.k8s.io/community/contributors/devel/README.md  
