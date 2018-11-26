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

This RFP will be open between 2018/10/29 and 2019/11/30.

The working group will answer questions for the first two weeks of this period.

Questions can be submitted [here](https://docs.google.com/forms/d/e/1FAIpQLSd5rXSDYQ0KMjzSEGxv0pkGxInkdW1NEQHvUJpxgX3y0o9IEw/viewform?usp=sf_link). All questions will be answered publicly in this document.

Proposals must include CVs, resumes, and/or example reports from staff that will be working on the project.

- 2018/10/29: RFP Open, Question period open
- 2018/11/12: Question period closes
- 2018/11/30: RFP Closes
- 2018/12/11: The working group will announce vendor selection

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

| # | Question | Answer |
|---|----------|--------|
| 1 | The RFP says that any area included in the out of scope section of the k8s bug bounty programme is not in-scope of this review.  There are some areas which are out of scope of the bug bounty which would appear to be relatively core to k8s, for example Kubernetes on Windows. Can we have 100% confirmation that these areas are out of scope? | Yes. If you encounter a vulnerability in Kubernetes' use of an out-of-scope element, like etcd or the container network interface (to Calico, Weave, Flannel, ...), that is in scope. If you encounter a direct vulnerability in a third-party component during the audit you should follow the embargo section of the RFP. |
| 2 | On the subject of target Distribution and configuration option review:<br> The RFP mentions an "audited reference architecture".<br> -	Is the expectation that this will be based on a specific k8s install mechanism (e.g. kubeadm)? <br> -	On a related note is it expected that High Availability configurations (e.g. multiple control plane nodes) should be included.<br> -	The assessment mentions Networking as a focus area.  Should a specific set of network plugins (e.g. weave, calico, flannel) be considered as in-scope or are all areas outside of the core Kubernetes code for this out of scope.<br> -	Where features of Kubernetes have been deprecated but not removed in 1.12, should they be considered in-scope or not? | 1. No, we are interested in the final topology -- the installation mechanism, as well as its default configuration, is tangental. The purpose is to contextualise the findings.<br>2. High-availability configurations should be included. For confinement of level of effort, vendor could create one single-master configuration and one high-availability configuration.<br>3. All plugins are out of scope per the bug bounty scope -- for clarification regarding the interface to plug-ins, please see the previous question.<br> 4. Deprecated features should be considered out of scope |
| 3 | On the subject of dependencies:<br>-        Will any of the project dependencies be in scope for the assessment? (e.g. https://github.com/kubernetes/kubernetes/blob/master/Godeps/Godeps.json) | Project dependencies are in scope in the sense that they are **allowed** to be tested, but they should not be considered a **required** testing area. We would be interested in cases where Kubernetes is exploitable due to a vulnerability in a project depdendency. Vulnerabilities found in third-party dependencies should follow the embargo section of the RFP.|
| 4 | Is the 8 weeks mentioned in the scope intended to be a limit on effort applied to the review, or just the timeframe for the review to occur in? | This is only a restriction on time frame, but is not intended to convey level of effort. |
| 5| Will the report be released in its entirety after the issues have been remediated? | Yes. |
| 6| What goals must be met to make this project a success? | We have several goals in mind:<br>1) Document a full and complete understanding of Kubernetes’ dataflow.<br>2) Achieve a reasonable understanding of potential vulnerability vectors for subsequent research.<br>3) Creation of artifacts that help third parties make a practical assessment of Kubernetes’ security position.<br>4) Eliminate design and architecture-level vulnerabilities.<br>5) Discover the most significant vulnerabilities, in both number and severity. |
| 7 | Would you be open to two firms partnering on the proposal? | Yes, however both firms should collaborate on the proposal and individual contributors should all provide C.V.s or past works.| 
| 8| From a deliverables perspective, will the final report (aside from the whitepaper) be made public? | Yes. |
| 9| The bug bounty document states the following is in scope, "Community maintained stable cloud platform plugins", however will the scope of the assessment include review of the cloud providers' k8s implementation? Reference of cloud providers: https://kubernetes.io/docs/concepts/cluster-administration/cloud-providers/ | Cloud provider-specific issues are excluded from the scope. |
| 10| The bug bounty doc lists supply chain attacks as in scope and also says, "excluding social engineering attacks against maintainers". We can assume phishing these individuals is out of scope, but does the exclusion of social engineering against maintainers include all attacks involving individuals? For example, if we were to discover that one of these developers accidentally committed their SSH keys to a git repo unassociated with k8s and we could use these keys to gain access to the k8s project. Is that in scope? | Attacks against individual developers, such as the example provided, are out of scope for this engagement. |
| 11| While suppression of logs is explicitly in scope, is log injection also in scope? | Log injection is in scope for the purposes of this audit.|
| 12| Are all the various networking implementations in scope for the assessment? Ref: https://kubernetes.io/docs/concepts/cluster-administration/networking/#how-to-implement-the-kubernetes-networking-model | Please refer to question 1. |
| 13| What does the working group refer to with formal threat model? Would STRIDE be a formal threat model in that sense?| A formal threat model should include a comprehensive dataflow diagram which shows data moving between different trust levels and assesses threats to that data using a system like STRIDE as the data moves between each process/component. Many good examples are present in Threat Modeling: Designing for Security by Adam Shostack. | 
| 14| Does Kubernetes uses any GoLang non-standard signing libraries? | An initial investigation has not uncovered any, however with a code base as large as Kubernetes, it is possible. |
| 15| Does Kubernetes implement any cryptographic primitives on its own, i.e. primitives which are not part of the standard libraries? | An initial investigation has not uncovered any, however with a code base as large as Kubernetes, it is possible. |
| 16| Presuming that live testing is part of the project, how does the working group see the "audited reference architecture" being defined? Is there a representative deployment, or a document describing a "default installation" that you foresee the engagement team using to inform the buildout of a test environment?| The purpose of the reference architecture is to define and document the configuration against which live testing was preformed. It should be generated collaboratively with the working group at the beginning of the project. We will want it to represent at least a common configuration, as in practice Kubernetes itself has no default configuration. It should take the form of a document detailing the set-up and configuration steps the vendor took to create their environment, ensuring an easily repeatable reference implementation. |
| 17| The RFP describes ""networking and multi-tenancy isolation"" as one of the focus areas. <br/><br/>Can you describe for us what these terms mean to you? Can you also help us understand how you define a soft non-hostile co-tenant? Is a _hostile_ co-tenant also in scope?| By networking we mean vulnerabilities related to communication within and to/from the cluster: container to container, pod to pod, pod to service, and external to internal communications as described in [the networking documentation](https://kubernetes.io/docs/concepts/cluster-administration/networking/). <br/><br/>The concept of soft multi-tenancy is that you have a single cluster being shared by applications or groups within the same company or organization, with less intended restrictions of a hard multi-tenant platform like a PaaS that hosts multiple distinct and potentially hostile competing customers on a single cluster which requires stricter security assumptions. These definitions may vary by group and use case, but the idea is that you can have a cluster with multiple groups with their own namespaces, isolated by networking/storage/RBAC roles."|
| 18| In the Artifacts section, you describe a Formal Threat Model as one of the outputs of the engagement. Can you expound on what this means to you? Are there any representative public examples you could point us to?| Please refer to question 13.|
