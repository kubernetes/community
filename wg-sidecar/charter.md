# WG Sidecar

This charter adheres to the conventions described in the [Kubernetes Charter README]
and uses the Roles and Organization Management outlined in [sig-governance].

[sig-governance]: https://github.com/kubernetes/community/blob/master/committee-steering/governance/sig-governance.md
[Kubernetes Charter README]: https://github.com/kubernetes/community/blob/master/committee-steering/governance/README.md

## Scope

The sidecar WG was created to work thru the requirements, design, implementation,
and future plans of natively supporting sidecar pattern in Kubernetes.

### In Scope

- Create a KEP for 1.27
- Provide supporting materials for easier KEP approval
- Approve the KEP for 1.27
- Discuss and coordinate implementation in 1.28
- Also the working group will reach out to various areas experts to finalize the implementation details.


## Stakeholders

Stakeholders in this working group span multiple SIGs with the SIG Node as a primary SIG.

- SIG Node
  Primary SIG as the most changes are in kubelet.
- SIG Scheduling
  APIs and changes related to Pods scheduling.
- SIG Architecture
  High-level input on requirements and API.
- SIG Apps
  Input on requirements for batch-like workload and other usage of sidecars.

## Deliverables

The ultimate deliverable is a built-in support of Sidecar containers in Kubernetes.

## Roles and Organization Management

This sig follows adheres to the Roles and Organization Management outlined in
[sig-governance] and opts-in to updates and modifications to [sig-governance].

[sig-governance]: https://github.com/kubernetes/community/blob/master/committee-steering/governance/sig-governance.md

## Timelines and Disbanding

We hope to get Sidecar KEP merged in 1.28 and finalize any additional requirements
and future plans in year 2023. Group will be discontinued once the feature GA-d
and there is not major follow up feature requiring extensive collaboration.
