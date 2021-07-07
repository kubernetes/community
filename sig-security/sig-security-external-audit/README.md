# SIG Security External Audit Subproject

## Overview

The SIG Security External Audit subproject (subproject, henceforth) is responsible for coordinating regular, comprehensive, third-party security audits.
The subproject publishes the deliverables of the audit after abiding to the [Security Release Process](https://github.com/kubernetes/security/blob/master/security-release-process.md) and [embargo policy](https://github.com/kubernetes/security/blob/master/private-distributors-list.md#embargo-policy).

  - [Request for Proposal (RFP)](#rfp)
    - [Security Audit Scope](#security-audit-scope)
    - [Vendor and Community Questions](#vendor-and-community-questions)
  - [Review of Proposals](#review-of-proposals)
  - [Vendor Selection](#vendor-selection)
  - [Deliverables](#deliverables)

## RFP

The subproject produces a RFP for a third-party, comprehensive security audit. The subproject publishes the RFP in the 'sig-security' folder in the `kubernetes/community/` repository. The subproject defines the scope, schedule, methodology, selection criteria, and deliverables in the RFP.

Previous RFPs:
  - [2019](https://github.com/kubernetes/community/blob/master/sig-security/security-audit-2019/RFP.md)
  - [2021](https://github.com/kubernetes/community/blob/master/sig-security/security-audit-2021/RFP.md)

As efforts begin for the year's security audit, create a tracking issue for the security audit in `kubernetes/community` with the `/sig security` label.

### Security Audit Scope

The scope of an audit is the most recent release at commencement of audit of the core [Kubernetes project](https://github.com/kubernetes/kubernetes) and certain other code maintained by [Kubernetes SIGs](https://github.com/kubernetes-sigs/).

Core Kubernetes components remain as focus areas of regular audits. Additional focus areas are finalized by the subproject.

### Vendor and Community Questions

Potential vendors and the community can submit questions regarding the RFP through a Google form. The Google form is linked in the RFP. [Example from the 2021 audit](https://docs.google.com/forms/d/e/1FAIpQLScjApMDAJ5o5pIBFKpJ3mUhdY9w5s9VYd_TffcMSvYH_O7-og/viewform).

The subproject answers questions publicly on the RFP with pull requests to update the RFP. [Example from the 2021 audit](https://github.com/kubernetes/community/pull/5813).

The question period is typically open between the RFP's opening date and closing date.

## Review of Proposals

Proposals are reviewed by the subproject proposal reviewers after the RFP closing date. An understanding of security audits is required to be a proposal reviewer.

All proposal reviewers must agree to abide by the **[Security Release Process](https://github.com/kubernetes/security/blob/master/security-release-process.md)**, **[embargo policy](https://github.com/kubernetes/security/blob/master/private-distributors-list.md#embargo-policy)**, and have no [conflict of interest](#conflict-of-interest) the tracking issue. This is done by placing a comment on the issue associated with the security audit. e.g. `I agree to abide by the guidelines set forth in the Security Release Process, specifically the embargo on CVE communications and have no conflict of interest`

Proposal reviewers are members of a private Google group and private Slack channel to exchange sensitive, confidential information and to share artifacts.

### Conflict of Interest

There is a possibility of a conflict of interest between a proposal reviewer and a vendor. Proposal reviewers should not have a conflict of interest. Examples of conflict of interest:
  - Proposal reviewer is employed by a vendor who submitted a proposal
  - Proposal reviewer has financial interest directly tied to the audit

Should a conflict arise during the proposal review, reviewers should notify the subproject owner and SIG Security chairs when they become aware of the conflict.

> The _Conflict of Interest_ section is inspired by the [CNCF Security TAG security reviewer process](https://github.com/cncf/tag-security/blob/main/assessments/guide/security-reviewer.md#conflict-of-interest).

## Vendor Selection

On the vendor selection date, the subproject will publish a the selected vendor in the 'sig-security' folder in the `kubernetes/community` repository. [Example from the 2019 audit](https://github.com/kubernetes/community/blob/master/sig-security/security-audit-2019/RFP_Decision.md).

## Deliverables

The deliverables of the audit are defined in the RFP e.g. findings report, threat model, white paper, audited reference architecture spec (with yaml manifests) and published in the 'sig-security' folder in the `kubernetes/community` repository. [Example from the 2019 audit](https://github.com/kubernetes/community/tree/master/sig-security/security-audit-2019/findings).

**All information gathered and deliverables created as a part of the audit must not be shared outside the vendor or the subproject without the explicit consent of the subproject and SIG Security chairs.**