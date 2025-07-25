# WG AI Gateway Charter

This charter adheres to the conventions described in the [Kubernetes Charter
README] and uses the Roles and Organization Management outlined in
[wg-governance].

[wg-governance]:https://github.com/kubernetes/community/blob/master/committee-steering/governance/wg-governance.md
[Kubernetes Charter README]:https://github.com/kubernetes/community/blob/master/committee-steering/governance/README.md

## Scope

The AI Gateway Working Group focuses on the intersection of AI and
networking, particularly in the context of extending load-balancer, gateway
and proxy technologies to manage and route traffic for AI Inference.

This working group will define terms like "AI Gateway" within the context of
Kubernetes and key use cases for users and implementations. It will propose
deliverables that need to be adopted in order to serve AI Inference on
Kubernetes.

This comes at a time where there is a proliferation of "AI Gateways" being used
for AI Inference, and a strong need for focus and collaboration to ensure
standards around this space so that Kubernetes users get the features they need
in a consistent way on the platform.

### In Scope

Overall guidance for the WG is to control scope as much as is feasible. The WG
should avoid AI-specific functionality where it can: instead favoring the
addition of provisions that help with AI use-cases, but are otherwise normal
networking facilities. Under that guidance, the following is in-scope:

* Providing definitions for networking related AI terms in a Kubernetes
  context.

* Defining important AI networking use-cases for Kubernetes users.

* Determining which common features and capabilities in the "AI Gateway" space
  need to be covered by Kubernetes standards and APIs according to user and
  implementation needs.

* Creating proposals for "AI Gateway" features and capabilities to the
  appropriate sub-projects.

* Propose new sub-projects if existing sub-projects are not sufficient.

### Out of Scope

* Developing whole "AI Gateway" solutions. This group will focus on
  enabling existing and new solutions to be more easily deployed and managed on
  Kubernetes, not adding any new production solutions maintained thereafter by
  upstream Kubernetes.

* Any specific kind of hardware support is generally out of scope.

* This group will not cover the entire spectrum of networking for AI. For
  instance: RDMA networks are generally out of scope.

## Deliverables

* A compendium of AI related networking definitions (e.g. "AI Gateway") and a
  key use-cases for Kubernetes users.

* Provide a space for collaboration and experimentation to determine the most
  viable features and capabitilies that Kubernetes should support. If there is
  strong consensus on any particular ideas, the WG will facilitate and
  coordinate the delivery of proposals in the appropriate areas.

## Stakeholders

* SIG Network

## Roles and Organization Management

This working group adheres to the Roles and Organization Management outlined in
[wg-governance] and opts-in to updates and modifications to [wg-governance].

[wg-governance]:https://github.com/kubernetes/community/blob/master/committee-steering/governance/wg-governance.md

## Exit Criteria

The WG is done when its deliverables are complete, according to the defined
scope and a list of key use cases and features agreed upon by the group.

Ideally we want the lifecycle of the WG to go something like this:

1. Determine definitions and key use cases for Kubernetes users and
   implementations, and document those.
2. Determine a list of key features that Kubernetes needs to best support the
   defined use cases.
3. For each feature in that list, make proposals which support them to the
   appropriate SIG and/or propose new sub-projects if deemed necessary.
4. Once the feature list is complete, leave behind some guidance and best
   practices for future implementations and then exit.
