# Request for Proposal

## Kubernetes Third-Party Security Audit

The Kubernetes SIG Security Third-Party Audit sub-project (working group, henceforth) is soliciting proposals from Information Security vendors for a comprehensive security audit of the Kubernetes Project.

### Background

In August of 2019, the Kubernetes Security Audit working group, in concert with the CNCF, Trail of Bits, and Atredis Partners, completed the first comprehensive security audit of the Kubernetes project’s [codebase](https://github.com/kubernetes/kubernetes/), working from version 1.13.

These findings, below, paint a broad picture of Kubernetes security, as of version 1.13, and highlight some areas that warrant further, deeper, research.

* [Kubernetes Security Review](../security-audit-2019/findings/Kubernetes%20Final%20Report.pdf)
* [Attacking and Defending Kubernetes Installations](../security-audit-2019/findings/AtredisPartners_Attacking_Kubernetes-v1.0.pdf)
* [Whitepaper](../security-audit-2019/findings/Kubernetes%20White%20Paper.pdf)
* [Threat Model](../security-audit-2019/findings/Kubernetes%20Threat%20Model.pdf)

### Project Goals and Scope

This subsequent audit is intended to be the second in a series of recurring audits, each focusing on a specific aspect of Kubernetes while maintaining coverage of all aspects that have changed since the previous audit ([1.13](../security-audit-2019/findings/)).

The scope of this audit is the most recent release at commencement of audit of the core [Kubernetes project](https://github.com/kubernetes/kubernetes) and certain other code maintained by [Kubernetes SIGs](https://github.com/kubernetes-sigs/).

This audit will focus on the following components of Kubernetes:

* kube-apiserver
* kube-scheduler
* etcd, Kubernetes use of
* kube-controller-manager
* cloud-controller-manager
* kubelet
* kube-proxy
* secrets-store-csi-driver

Adjacent findings within the scope of the [bug bounty program](https://hackerone.com/kubernetes?type=team#scope) may be included, but are not the primary goal.

This audit is intended to find vulnerabilities or weaknesses in Kubernetes. While Kubernetes relies upon container runtimes such as Docker and CRI-O, we aren't looking for (for example) container escapes that rely upon bugs in the container runtime (unless, for example, the escape is made possible by a defect in the way that Kubernetes sets up the container).

The working group is specifically interested in the following aspects of the in-scope components. Proposals should indicate the specific proposed personnel’s level of expertise in these fields as it relates to Kubernetes.

* Golang analysis and fuzzing
* Networking
* Cryptography
* Evaluation of component privilege
* Trust relationships and architecture evaluation
* Authentication & Authorization (including Role Based Access Controls)
* Secrets management
* Multi-tenancy isolation: Specifically soft (non-hostile co-tenants)

Personnel written into the proposal must serve on the engagement, unless explicit approvals for staff changes are made by the Security Audit Working Group.

#### Out of Scope

Findings specifically excluded from the [bug bounty program](https://hackerone.com/kubernetes?type=team#scope) scope are out of scope for this audit.

### Eligible Vendors

This RFP is open to proposals from all vendors.

#### Constraints

If your proposal includes subcontractors, please include relevant details from their firms such as CVs, past works, etc. The selected vendor will be wholly responsible for fulfillment of the audit and subcontractors must be wholly managed by the selected vendor.

### Anticipated Selection Schedule

This RFP will be open until 4 proposals have been received.
The RFP closing date will be set 2 calendar weeks after the fourth proposal is received.
The working group will announce the vendor selection after reviewing proposals.
Upon receipt of the fourth proposal, the working group will update the RFP closure date and vendor selection date in this document.

The working group will answer questions for the RFP period.

Questions can be submitted [here](https://docs.google.com/forms/d/e/1FAIpQLScjApMDAJ5o5pIBFKpJ3mUhdY9w5s9VYd_TffcMSvYH_O7-og/viewform). All questions will be answered publicly in this document.

We understand scheduling can be complex but we prefer to have proposals include CVs, resumes, and/or example reports from staff that will be working on the project.

Proposals should be submitted to kubernetes-security-audit-2021@googlegroups.com

* 2021/02/08: RFP Open, Question period open
* 2021/06/22: Fourth proposal received
* 2021/07/06: RFP Closes, Question period closes
* 2021/08/31: The working group will announce vendor selection

## Methodology

The start and end dates will be negotiated after vendor selection. The timeline for this audit is flexible.

The working group will establish a 60 minute kick-off meeting to answer any initial questions and discuss the Kubernetes architecture.

This is a comprehensive audit, not a penetration test or red team exercise. The audit does not end with the first successful exploit or critical vulnerability.

The vendor will document the Kubernetes configuration and architecture that they will audit and provide this to the working group. The cluster deployment assessed must not be specific to any public cloud. The working group must approve this configuration before the audit continues. This documented configuration will result in the "audited reference architecture specification" deliverable.

The vendor will perform source code analysis on the Kubernetes code base, finding vulnerabilities and, where possible and making the most judicious use of time, providing proof of concept exploits that the Kubernetes project can use to investigate and fix defects. The vendor will discuss findings on a weekly basis and, at the vendor’s discretion, bring draft write-ups to status meetings.

The working group will be available weekly to meet with the selected vendor and will provide subject matter experts as requested.

The vendor will develop and deliver a draft report, describing their methodology, how much attention the various components received (to inform future work), and the work’s findings. The working group will review and comment on the draft report, either requesting updates or declaring the draft final. This draft-review-comment-draft cycle may repeat several times.

## Expectations

The vendor must report urgent security issues immediately to both the working group and security@kubernetes.io.

## Selection Criteria

To help us combine objective evaluations with the working group members’ individual past experiences and knowledge of the vendors’ work and relevant experience, the vendors will be evaluated against the following criteria. Each member of the working group will measure the RFP against the criteria on a scale of 1 to 5:

* Relevant understanding and experience in code audit, threat modeling, and related work
* Relevant understanding and experience in Kubernetes, other orchestration systems, containers, Linux, hardening of distributed systems, and related work
* Strength of the vendor’s proposal and examples of previous work product, redacted as necessary

A writeup which details our process and results of the last RFP is available [here](../security-audit-2019/RFP_Decision.md).

## Confidentiality and Embargo

All information gathered and deliverables created as a part of the audit must not be shared outside the vendor or the working group without the explicit consent of the working group.

## Deliverables

The audit should result in the following deliverables, which will be made public after any sensitive security issues are mitigated.

* Audited reference architecture specification. Should take the form of a summary and associated configuration YAML files.
* Findings report including an executive summary.
* Where possible and, in the vendor’s opinion makes the most judicious use of time, proof of concept exploits that the Kubernetes project can use to investigate and fix defects.

## Questions Asked during RFP Response Process

### Do we need to use our own hardware and infrastructure or should we use a cloud?

Strong preference would be for the vendor to provide their own infrastructure or use a public cloud provider, just NOT a managed offering like GKE or EKS. The reasoning is to prevent accidentally auditing a cloud provider's kubernetes service instead of kubernetes/kubernetes. Depending on the scope and approach, it may make sense to use a local cluster (e.g. kind) for API fuzzing and anything that doesn't impact the underlying OS, and is an easy to use repeatable setup (see Methodology above).

### What is the intellectual property ownership of the report and all work product?

The report must be licensed under the Creative Commons Attribution 4.0 International Public License (CC BY 4.0) based on [section 11.(f) of the Cloud Native Computing Foundation (CNCF) Charter](https://github.com/cncf/foundation/blob/master/charter.md#11-ip-policy).
Separately, any code released with or as part of the report needs to be under the Apache License, version 2.0. Please refer to [sections 11.(e) and (d) in the CNCF Charter](https://github.com/cncf/foundation/blob/master/charter.md#11-ip-policy).

### Must I use the report format from the previous audit? Can the SIG provide a report format template I can use?

Vendors who wish to use either the previous report format, as allowed by CC BY 4.0, or a report format provided by the community may do so as long as it is also available under CC BY 4.0. Vendors who wish to publish 2 versions of the report, one tailored for the community under CC BY 4.0 and one that they host on their own site using their proprietary fonts, formats, branding, or other copyrights, under their own license may do so, in order to differentiate their commercial report format from this report. Vendors may also publish a synopsis and marketing materials regarding the report on their website as long as it links to the original report in this repository.  In the community report, vendors can place links in the report to materials hosted on their commercial site. This does not imply that linked materials are themselves CC BY 4.0.

### Do you have any developer documentation or design documentation specifications that aren't available on the internet that you would be able to share?

Kubernetes is an open source project, all documentation is available on https://kubernetes.io or on https://github.com/kubernetes.

### What are the most important publicly available pages detailing the design of the system and the data it receives.

- Overview of [Kubernetes components](https://kubernetes.io/docs/concepts/overview/components/)
- [kube-apiserver overview](https://kubernetes.io/docs/reference/command-line-tools-reference/kube-apiserver/)
- [kube-scheduler overview](https://kubernetes.io/docs/concepts/scheduling-eviction/kube-scheduler/)
- [Operating etcd clusters for Kubernetes](https://kubernetes.io/docs/tasks/administer-cluster/configure-upgrade-etcd/)
- [etcd clustering guide](https://etcd.io/docs/next/op-guide/clustering/)
- [kube-controller-manager overview](https://kubernetes.io/docs/reference/command-line-tools-reference/kube-controller-manager/)
- [cloud-controller-manager overview](https://kubernetes.io/docs/concepts/architecture/cloud-controller/)
- [cloud-controller-manager administration](https://kubernetes.io/docs/tasks/administer-cluster/running-cloud-controller/)
- [kubelet overview](https://kubernetes.io/docs/reference/command-line-tools-reference/kubelet/)
- [kube-proxy overview](https://kubernetes.io/docs/reference/command-line-tools-reference/kube-proxy/)
- [secrets-store-csi-driver](https://github.com/kubernetes-sigs/secrets-store-csi-driver)

### How long does the Working Group envision the engagement lasting and what is the latest date you can receive the deliverables?

The latest date to receive deliverables will be negotiated with the selected vendor.

### Which attack vectors are of most concern to the Working Group.

1. The attack vector most concerned about is unauthenticated access to a cluster resulting in compromise of the [components in-scope](#project-goals-and-scope)
2. Crossing namespace boundaries, an authenticated attacker being able to affect resources their credentials do not directly allow
3. Any other attack vector that exists against the components in scope

### Is there flexibility to wait for staff to be available to work on the audit?

Yes, the timeline for the audit is flexible and the timeline will be further discussed and negotiated with the selected vendor.