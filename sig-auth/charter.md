# SIG Auth Charter

This charter adheres to the conventions described in the [Kubernetes Charter README] and uses
the Roles and Organization Management outlined in [sig-governance].

## Scope

SIG Auth is responsible for the design, implementation, and maintenance of features in
Kubernetes that control and protect access to the API and other core components. This includes
authentication and authorization, but also encompasses features like auditing and some security
policy (see below).

### In scope

Link to SIG section in [sigs.yaml]

#### Code, Binaries and Services

- Kubernetes authentication, authorization, audit and security policy features. Examples
  include:
    - Authentication, authorization and audit interfaces and extension points
    - Authentication implementations (service accounts, OIDC, authenticating proxy, webhook,
      ...)
    - Authorizer implementations (RBAC + default policy, Node + default policy, webhook, ...)
    - Security-related admission plugins (NodeRestriction, ServiceAccount, PodSecurityPolicy,
      ImagePolicy, etc)
- The mechanisms to protect confidentiality/integrity of API data. Examples include:
    - Capability for encryption at rest
    - Capability for secure communication between components
    - Ensuring users and components can operate with appropriately scoped permissions

#### Cross-cutting and Externally Facing Processes

- Consult with other SIGs and the community on how to apply mechanisms owned by SIG
  Auth. Examples include:
    - Review privilege escalation implications of feature and API designs
    - Core component authentication & authorization (apiserver, kubelet, controller-manager,
      and scheduler)
    - Local-storage volume deployment authentication
    - Cloud provider authorization policy
    - Container runtime streaming (exec/attach/port-forward) authentication
    - Best practices for hardening add-ons or other external integrations

### Out of scope

- Reporting of specific vulnerabilities in Kubernetes. Please report using these instructions:
  https://kubernetes.io/security/
- General security discussion. Examples of topics that are out of scope for SIG-auth include:
  - Protection of volume data, container ephemeral data, and other non-API data (prefer: sig-storage
    and sig-node)
  - Container isolation (prefer: sig-node and sig-networking)
  - Bug bounty (prefer: product security committee)
  - Resource quota (prefer: sig-scheduling)
  - Resource availability / DOS protection (prefer: sig-apimachinery, sig-network, sig-node)

## Roles and Organization Management

This sig follows adheres to the Roles and Organization Management outlined in [sig-governance]
and opts-in to updates and modifications to [sig-governance].

### Subproject Creation

SIG Auth delegates subproject approval to Technical Leads. See [Subproject creation - Option 1].


[sig-governance]: https://github.com/kubernetes/community/blob/master/committee-steering/governance/sig-governance.md
[sigs.yaml]: https://github.com/kubernetes/community/blob/master/sigs.yaml#L250
[Kubernetes Charter README]: https://github.com/kubernetes/community/blob/master/committee-steering/governance/README.md
[Subproject creation - Option 1]: https://github.com/kubernetes/community/blob/master/committee-steering/governance/sig-governance.md#subproject-creation
