# Request for Proposal

## Kubernetes Third Party Security Audit

The Kubernetes Third-Party Audit Working Group (working group, henceforth) is soliciting proposals from select Information Security vendors for a comprehensive security audit of the Kubernetes Project.

### Eligible Vendors

Only the following vendors will be permitted to submit proposals:

- NCC Group
- Trail of Bits
- Cure53
- Bishop Fox
- Insomnia
- Atredis Partners

If your proposal includes sub-contractors, please include relevant details from their firm such as CVs, past works, etc.

### RFP Process

This RFP will be open between 2018/10/29 and 2019/11/26.

The working group will answer questions for the first two weeks of this period.

Questions can be submitted [here](https://docs.google.com/forms/d/e/1FAIpQLSd5rXSDYQ0KMjzSEGxv0pkGxInkdW1NEQHvUJpxgX3y0o9IEw/viewform?usp=sf_link). All questions will be answered publicly in this document.

Proposals must include CVs, resumes, and/or example reports from staff that will be working on the project.

- 2018/10/29: RFP Open, Question period open
- 2018/11/12: Question period closes
- 2018/11/26: RFP Closes
- 2018/12/04: The working group will announce vendor selection

## Audit Scope

The scope of the audit is the most recent release (1.12) of the core [Kubernetes project](https://github.com/kubernetes/kubernetes).

- Findings within the [bug bounty program](https://github.com/kubernetes/community/blob/master/contributors/guide/bug-bounty.md) scope are in scope.

 We want the focus of the audit to be on bugs on Kubernetes. While Kubernetes relies upon a container runtimes such as Docker and CRI-O, we aren't looking for (for example) container escapes that rely upon bugs in the container runtime (unless, for example, the escape is made possible by a defect in the way that Kubernetes sets up the container).

### Focus Areas

The Kubernetes Third-Party Audit Working Group is specifically interested in the following areas. Proposals should indicate their level of expertise in these fields as it relates to Kubernetes.

- Networking
- Cryptography
- Authentication & Authorization (including Role Based Access Controls)
- Secrets management
- Multi-tenancy isolation: Specifically soft (non-hostile co-tenants)

### Out of Scope

Findings specifically excluded from the [bug bounty program](https://github.com/kubernetes/community/blob/master/contributors/guide/bug-bounty.md) scope are out of scope.

## Methodology

We are allowing 8 weeks for the audit, start date can be negioated after vendor selection. We recognize that November and December can be very high utilization periods for security vendors.

The audit should not be treated as a penetration test, or red team exercise. It should be comprehensive and not end with the first successful exploit or critical vulnerability.

The vendor should perform both source code analysis as well as live evaluation of Kubernetes.

The vendor should document the Kubernetes configuration and architecture that the audit was performed against for the creation of a "audited reference architecture" artifact. The working group must approve this configuration before the audit continues.

The working group will establish a 60 minute kick-off meeting to answer any initial questions and explain Kubernetes architecture.

The working group will be available weekly to meet with the selected vendor, will and provide subject matter experts for requested components.

The vendor must report urgent security issues immediately to both the working group and security@kubernetes.io.

## Confidentiality and Embargo

All information gathered and artifacts created as a part of the audit must not be shared outside the vendor or the working group without the explicit consent of the working group.

## Artifacts

The audit should result in the following artifacts, which will be made public after any sensitive security issues are mitigated.

- Findings report, including an executive summary

- Audited reference architecture specification. Should take the form of a summary and associated configuration yaml files.

- Formal threat model

- Any proof of concept exploits that we can use to investigate and fix defects

- Retrospective white paper(s) on important security considerations in Kubernetes

  *This artifact can be provided up to 3 weeks after deadline for the others.*

  - E.g. [NCC Group: Understanding hardening linux containers](https://www.nccgroup.trust/globalassets/our-research/us/whitepapers/2016/april/ncc_group_understanding_hardening_linux_containers-1-1.pdf)
  - E.g. [NCC Group: Abusing Privileged and Unprivileged Linux
    Containers](https://www.nccgroup.trust/globalassets/our-research/us/whitepapers/2016/june/container_whitepaper.pdf)

## Q & A

This section intentionally left empty until the question period opens.
