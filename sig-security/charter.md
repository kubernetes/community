# SIG Security Charter

This charter adheres to the conventions described in the [Kubernetes Charter README] and uses the Roles and Organization Management outlined in [sig-governance].

## Scope

SIG Security covers horizontal security initiatives for the Kubernetes project, including regular security audits, the vulnerability management process, cross-cutting security documentation, and security community management. As a process-oriented SIG, it does not directly own Kubernetes component code. This SIG replaces the Security Audit Working Group. Instead, SIG Security focuses on improving the security of the Kubernetes project across all components.

This SIG grew out of the [Third-Party Security Audit Working Group](https://github.com/kubernetes/community/tree/master/sig-security/security-audit-2019), which managed each recurrent Third-Party Security Audit over the course of the audit’s lifecycle. The Working Group worked closely with selected vendors, the Product Security Committee, and the CNCF. It created the RFP, selected the vendors, and managed the vendors’ engagement with other SIGs and subject matter experts.

SIG Security continues to manage the third-party security audits, while serving a wider mission of advocating for security-related structural or systemic issues and default configuration settings, managing the non-embargoed (public) vulnerability process, defining the bug bounty, creating official Kubernetes Hardening Guides and security documents, and serving as a public relations contact point for Kubernetes security. 

### In scope

#### Vulnerability Management Process

Work with the Kubernetes [Security Response Committee (SRC)](https://github.com/kubernetes/committee-security-response#security-response-committee-src) to define the processes for fixing and disclosing vulnerabilities, as outlined in https://github.com/kubernetes/committee-security-response. For example:

- When the private fix & release process is invoked
- How vulnerabilities are rated
- The scope of the bug bounty
- Post-announcement follow-ups, such as additional fixes, mitigations, preventions or documentation after a vulnerability is made public
- Distributor announcement policies, such as timelines, criteria for joining the list, etc.
- How, when and where vulnerabilities are announced
- Defining the criteria and process for supporting Kubernetes subprojects, such as the [dashboard](https://github.com/kubernetes/dashboard), [ingress-nginx](https://github.com/kubernetes/ingress-nginx), or [kops](https://github.com/kubernetes/kops).

#### Security Community Management and Outreach

Provide an entry point to the Kubernetes community for new security-minded contributors, as well as a meeting point to discuss security themes and issues within Kubernetes, including:

- Work with [SIG Contributor Experience](https://github.com/kubernetes/community/tree/master/sig-contributor-experience) to curate and staff security discussion channels (e.g. slack channel, mailing list, discourse, stack overflow, etc.).
- Answer security questions from inexperienced users (that don't know what SIG to go to), and identify common questions or issues as areas for improvement.
- Provide an "entry point" for new contributors interested in security. Route these new contributors to other SIGs when they have more specific goals (e.g. SIG Node for container isolation).

#### Horizontal Security Documentation

Author and maintain cross-cutting security documentation, such as hardening guides and security benchmarks. Seek out and coordinate with experts in other SIGs for input on the documentation (i.e. we go to them, they don't need to come to us). In-scope documentation includes:

- Hardening guides and best practices
- Security benchmarks
- Improving documentation to address common misunderstandings or questions
- Threat models

#### Security Audit

Manage recurring security audits and follow up on issues. Coordinate vendors to perform the audit and publish the findings. Follow up on issues with the affected SIG and help coordinate resolution, which can include:

- Helping to prioritize the fixes, possibly by recruiting from SIG Security (while acknowledging that the ultimate authority in deciding whether and how to fix an issue lies with the responsible SIG).
- Documenting mitigations, workarounds, or caveats, especially when the responsible SIG decides not to fix a reported issue.

### Out of scope

In contrast to SIG Auth, SIG Security does not own any Kubernetes cluster component code. 

Further, SIG Security’s scope does not include:

- Kubernetes authentication, authorization, audit and security policy features.  (SIG Auth)
- Private vulnerability response (belongs to the PSC), including:
    - Embargoed vulnerability management
    - Bug bounty submission triage and management
    - Non-public vulnerability collection, triage, and disclosure
- The mechanisms to protect confidentiality/integrity of API data (belongs to SIG API Machinery, SIG Auth or others)
- Security audit for all other CNCF projects (e.g., etcd, CoreDNS, CRI-O, containerd)  (Belongs to the CNCF’s SIG Security.) 
- Any projects outside of the Kubernetes project
- Cloud provider-specific or distributor-specific hardening guides
- Recommendations or endorsements of specific commercial product vendors or cloud providers.


## Roles and Organization Management

This SIG adheres to the Roles and Organization Management outlined in [sig-governance] and opts-in to updates and modifications to [sig-governance].

### Additional responsibilities of Chairs

None defined at this time.

### Additional responsibilities of Tech Leads

- Security Documents and Documentation Tech Leads will be responsible for maintaining the official Kubernetes project Security Hardening Guide.

### Subproject Creation

SIG Security delegates subproject approval to SIG Technical Leads. See Subproject creation - Option 1.

SIG Security’s initial subprojects will be:

- Security Documents and Documentation
- Third Party Security Audit
- Community Discussion Groups
