# 2025 Annual Report: SIG Security

## Current initiatives and Project Health

1. What work did the SIG do this year that should be highlighted?

* SIG Security has taken ownership of the previously unmaintained but widely used [OWASP Kubernetes Top Ten](https://owasp.org/www-project-kubernetes-top-ten/). After taking this on, the SIG Security Docs subproject surveyed the community for feedback on updates to the Top Ten list and to offer ideas on any potential additions. We have since created a complete draft of the updated list, which is actively being reviewed. The next step will be to publish the updated list to the OWASP site and share with the Kubernetes community. We have also submitted a talk proposal to OWASP AppSec EU 2026 to share the updated list and talk about SIG Security’s plans for the project.

* SIG Security Third-Party Audit subproject began the newest third-party audit of selected Kubernetes features. Working with the Open Source Technology Improvement Fund, we selected a vendor (Shielder) and worked with them to get the audit done. The audit report has now been provided and is currently embargoed until findings have been triaged and the final report is published. We are looking forward to sharing the findings with the public when we can\!

* SIG Security Tooling has made numerous improvements to help the project and community stay secure. We migrated our vulnerability scan tooling onto a new cluster, supporting SIG K8s Infra’s effort to lock down the trusted clusters more tightly. The Official CVE Feed is easier to maintain now that automated testing has been added, and the results are more accurate due to improved handling of formatting irregularities. We also introduced new tooling to make the Security Response Committee’s vulnerability publication process easier and more consistent.

* SIG Security Self-Assessments subproject started a self-assessment of etcd with etcd maintainers. That assessment is currently in progress.  
    
2. Are there any areas and/or subprojects that your group needs help with (e.g. fewer than 2 active OWNERS)?  

We have sufficient active OWNERS in all subprojects.

SIG Security Docs subproject has several ongoing efforts that could use extra assistance to help make Kubernetes safer for users. We would love to publish additional chapters of the Kubernetes Hardening Guide, and the research to write one provides a great opportunity to learn more and get involved. We’re intentionally welcoming to new folks, so stop by a meeting or #sig-security-docs on Kubernetes slack and say hello!

3. Did you have community-wide updates in 2025 (e.g. KubeCon talks)?  

SIG Security offered maintainer track talks at KubeCon EU and KubeCon NA in 2025. The recordings are available here: [SIG Security: Succession Planting for a Flowering Future](https://www.youtube.com/watch?v=0p-sZT0LWOg), [Butterfly Effect: What Kubernetes SIG Security Has in Flight](https://www.youtube.com/watch?v=q4ryxYhAenM)

In addition, we offered Project Pavillion booths at both KubeCon EU and KubeCon NA. Our booths are highlights of the year for SIG Security, drawing in contributors new and old alike. KubeCon attendees can count on the SIG Security booth as a place to find a welcoming environment, friendly conversation, and ways to get involved.

4. KEP work in 2025 (v1.33, v1.34, v1.35):

In 2025, we continued to maintain and improve the [Official CVE Feed](https://github.com/kubernetes/enhancements/issues/3203). We fixed several bugs and began migrating the feed generator to new hosting infrastructure.

## [Subprojects](https://git.k8s.io/community/sig-security#subprojects)

**Retired and Resurrected in 2025:**

- security-self-assessments

**Continuing:**

- security-audit  
- security-docs  
- security-tooling  
- sig-security

## [Working groups](https://git.k8s.io/community/sig-security#working-groups)

**Continuing:**

- LTS

## Operational

Operational tasks in [sig-governance.md](https://git.k8s.io/community/committee-steering/governance/sig-governance.md):

- [x] [README.md](https://git.k8s.io/community/sig-security/README.md) reviewed for accuracy and updated if needed  
- [x] [CONTRIBUTING.md](https://git.k8s.io/community/sig-security/CONTRIBUTING.md) reviewed for accuracy and updated if needed  
- [x] Other contributing docs (e.g. in devel dir or contributor guide) reviewed for accuracy and updated if needed  
- [x] Subprojects list and linked OWNERS files in [sigs.yaml](https://git.k8s.io/community/sigs.yaml) reviewed for accuracy and updated if needed  
- [x] SIG leaders (chairs, tech leads, and subproject leads) in [sigs.yaml](https://git.k8s.io/community/sigs.yaml) are accurate and active, and updated if needed  
- [x] Meeting notes and recordings for 2025 are linked from [README.md](https://git.k8s.io/community/sig-security/README.md) and updated/uploaded if needed

