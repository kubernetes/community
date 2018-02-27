# Kubernetes Enhancement Proposals (KEPs)

A Kubernetes Enhancement Proposal (KEP) is a way to propose, communicate and coordinate on new efforts for the Kubernetes project.
You can read the full details of the project in [KEP-1](0001-kubernetes-enhancement-proposal-process.md).

This process is still in a _beta_ state and is opt-in for those that want to provide feedback for the process.

## Quick start for the KEP process

1. Socialize an idea with a sponsoring SIG.
   Make sure that others think the work is worth taking up and will help review the KEP and any code changes required.
2. Follow the process outlined in the [KEP template](0000-kep-template.md)

## FAQs

### Do I have to use the KEP process?

No... but we hope that you will.
Over time having a rich set of KEPs in one place will make it easier for people to track what is going in the community and find a structured historic record.

KEPs are only required when the changes are wide ranging and impact most of the project.
These changes are usually coordinated through SIG-Architecture.
It is up to any specific SIG if they want to use the KEP process and when.
The process is available to SIGs to use but not required.

### Why would I want to use the KEP process?

Our aim with KEPs is to clearly communicate new efforts to the Kubernetes contributor community.
As such, we want to build a well curated set of clear proposals in a common format with useful metadata.

Benefits to KEP users (in the limit):
* Exposure on a kubernetes blessed web site that is findable via web search engines.
* Cross indexing of KEPs so that users can find connections and the current status of any KEP.
* A clear process with approvers and reviewers for making decisions.
  This will lead to more structured decisions that stick as there is a discoverable record around the decisions.

We are inspired by IETF RFCs, Pyton PEPs and Rust RFCs.
See [KEP-1](0001-kubernetes-enhancement-proposal-process.md) for more details.

### Do I put my KEP in the root KEP directory or a SIG subdirectory?

If the KEP is mainly restricted to one SIG's purview then it should be in a KEP directory for that SIG.
If the KEP is widely impacting much of Kubernetes, it should be put at the root of this directory.
If in doubt ask [SIG-Architecture](../sig-architecture/README.md) and they can advise.

### What will it take for KEPs to "graduate" out of "beta"?

Things we'd like to see happen to consider KEPs well on their way:
* A set of KEPs that show healthy process around describing an effort and recording decisions in a reasonable amount of time.
* KEPs exposed on a searchable and indexable web site.
* Presubmit checks for KEPs around metadata format and markdown validity.

Even so, the process can evolve. As we find new techniques we can improve our processes.

### My FAQ isn't answered here!

The KEP process is still evolving!
If something is missing or not answered here feel free to reach out to [SIG-Architecture](../sig-architecture/README.md).
If you want to propose a change to the KEP process you can open a PR on [KEP-1](0001-kubernetes-enhancement-proposal-process.md) with your proposal.
