# WG AI Integration Charter

This charter adheres to the conventions described in the [Kubernetes Charter
README] and uses the Roles and Organization Management outlined in
[wg-governance].

## Scope

The AI Integration Working Group focuses on enabling seamless integration of
AI/ML control planes with Kubernetes, as well as providing standardized
patterns for deploying, managing, and operating AI applications at scale
on Kubernetes.

The Working Group will provide a forum for a broad engineering community to
give feedback to the project on challenges encountered when integrating with
Kubernetes.

This addresses a broad need with many end-users deploying complex AI systems,
AI/ML platform providers, Kubernetes distributions, and developers of
distributed AI applications facing these integration challenges. Standardizing
solutions in this space benefits the entire Kubernetes ecosystem. Adjacent
ecosystems could link to the outputs of this WG as a trusted vehicle for
supporting AI integrations with Kubernetes.

### In scope

* Develop a shared community point of view and associated best practices
enabling AI agent (or multi-agent) systems to integrate with Kubernetes.

* Provide a forum for intersecting code experimentation in AI integration
space and discussion with the existing Kubernetes community.

* Recommend an appropriate go forward governance model for AI Integrations
with the Kubernetes project.

* Identify appropriate auth(z) patterns for AI connector identities, its
 closest caller, and Kubernetes RBAC.

* Defining benchmarks on pros/cons of design approaches to meet user outcomes.

* Ensure security, observability, and policy enforcement can be consistently
applied across integrated systems (K8s and external Control Planes such as
LLMs) and AI integration applications.

* Define potential enhancements to API conventions to scale AI integration
patterns that respect data privacy and safety concerns during our design
process. Consider alternative API patterns that could be a better fit for
AI enablement.

* Explore patterns for efficient network access to emergent protocols such
as MCP/A2A via proxies or gateways.

* Reduce the complexity and custom development required for deploying,
building and managing connectors of kubernetes API with AI agent ecosystems.

### Out of Scope

* Development of AI/ML frameworks or applications
* General-purpose workload management not specific to AI/ML
* Deploying inference workloads on Kubernetes (which is covered by WG Serving)
* Manage accelerator devices (which is covered by WG Device Management)

## Deliverables

* The WG will provide space for collaboration and experimentation. If/when any
 solid ideas emerge that require changes to Kubernetes (for example, updates
 to kubectl for AI consumption), the WG will facilitate and coordinate the delivery
 of KEPs and their implementations by the participating SIGs.
* Interim artifacts will include documents capturing use cases, requirements,
 integration architecture designs, and AI application communication patterns.
* Establish best practices document for AI tool integration with Kubernetes and
 a clear recommendation if/what set of reference tools may best fit in
 Kubernetes project itself informed from data driven experimentation with
 appropriate governance model.

## Stakeholders

* SIG Architecture
* SIG API Machinery
* SIG Apps
* SIG Auth
* SIG CLI

## Roles and Organization Management

This working group adheres to the Roles and Organization Management outlined in
[wg-governance] and opts-in to updates and modifications to [wg-governance].

## Exit Criteria

The WG is done if/when a shared recommendation is in place for how the Kubernetes
project should or should not integrate with these emergent systems.  This could
include a recommendation for Kubernetes to adopt and/or evolve tools (e.g. MCP
connectors, benchmark or environment validation tooling, etc.) and evolve its
own governance model to provide proper stewardship within the project or outside.

The working group will disband when the KEPs resulting from these discussions
have reached a terminal state. When the core functionality for AI workload
management reaches GA, we will evaluate whether the working group should
be disbanded and any remaining KEPs be left to the management of their owning
SIGs.

[wg-governance]: https://github.com/kubernetes/community/blob/master/committee-steering/governance/wg-governance.md
[Kubernetes Charter README]: https://github.com/kubernetes/community/blob/master/committee-steering/governance/README.md
