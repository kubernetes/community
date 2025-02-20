# SIG K8s Infra Charter

This charter adheres to the conventions described in the
[Kubernetes Charter README] and uses the Roles and Organization Management
outlined in [sig-governance].

## Scope

The successful migration of ownership and management of all Kubernetes
project infrastructure from the google.com GCP Organization 
(or other IaaS vendor-owned locations) to the CNCF, such that the Kubernetes
project is able to sustainably operate itself without direct assistance from
external vendors or entities.

In other words, we seek to eradicate usage of the phrase "oh that's
something that only an employee of Vendor X can do, we're blocked until
they respond."

### In scope

Within this document, "infrastructure" is used to refer to cloud resources
managed through an "infrastructure as a service" offering. This includes
more than just raw compute, storage, and networking resources, since many
cloud services provide a rich variety of resources for API-driven management.

#### Code, Binaries and Services

Code, data and policies necessary to provision, update, decommission and
otherwise manage all project infrastructure as provisioned through
infrastructure-as-a-service (IaaS) offerings. This includes more than raw
compute, storage, and network resources traditionally bucketed under IaaS,
since many cloud offerings provide a rich variety of resources via API-driven
management. This may also include code and binaries which run on top of the
IaaS offerings to provide services to the Kubernetes project.

Given that this is a broad scope, we prefer (where possible) to delegate
ownership and operation of the code / infrastructure to more directly
responsible SIGs or Committees. This is largely how the SIG operated during
its lifetime as a WG, driving the policies and tooling upon which SIG-owned
infrastructure operates.

Areas of responsibility include:

- Policy definition and enforcement for areas related to project
  infrastructure, including:
  - What is in-scope/out-of-scope for project infrastructure
  - Who should be allowed access to which parts of project infrastructure,
    e.g. team definition, vetting criteria, etc.
  - How infrastructure should be managed, e.g. naming schemes, acceptable
    tooling or practices, on-call or escalation policies, etc.
- Configuration management of all resources and service usage within the
  kubernetes.io GCP Organization, including, but not limited to:
  - API / Service enablement
  - BigQuery datasets
  - DNS records, e.g. for k8s.io, kubernetes.io, and other project-owned domains
  - GCB usage
  - GCP projects, instances, images
  - GCR repositories
  - GCS buckets
  - GKE clusters, e.g. community infra cluster, prow build clusters
  - GSM secrets
  - Google Groups
  - IAM roles, service accounts, and policies
  - KMS keys
  - Managed Certificates, e.g. for k8s.io, kubernetes.io, and other project-owned
    domains
- Reports on infrastructure operation, including:
  - Anonymized traffic reports to show which parts of our infrastructure
    are seeing the most use
  - Auditing reports to show the current configuration of the community's
    infrastructure
  - Billing reports to show where the community's infrastructure budget is
    being spent

In terms of subprojects, this means we own kubernetes/k8s.io and are an
escalation point of last resort for more tightly scoped subprojects that
live within this repo.

#### Cross-cutting and Externally Facing Processes

We prefer (where possible) to delegate ownership, operation and policy
definition to SIGs that are more directly responsible for a given area
of the project. However, we reserve the right to halt infrastructure or
roll back changes if the project as a whole is being negatively impacted.

Some examples for illustrative purposes

##### Access Policies

- We are responsible for ensuring the appropriate members of a SIG have
  sufficient permissions to troubleshoot and manage their app or
  infrastructure.
- However, we will NOT grant overly broad permissions to an overly broad
  group of people. We will collaborate with SIGs to ensure access is
  appropriately scoped.
- We WILL ensure the appropriate set of CNCF staff have access to act as
  an escalation path of last resort
- We MAY revoke access in the event of a security-related incident

e.g. SIG Release is responsible for who gets what level of access to
infrastructure used by the release-engineering subproject to cut a Kubernetes
release

##### Artifact Hosting

- We are not responsible for promoting into production artifacts that belong
  to subprojects owned by other SIGs.
- However, we MAY revert changes that prevent artifact promotion from
  functioning.

e.g. SIG Storage is responsible for declaring which CSI-related images should
be promoted to production, SIG Release is responsible for ensuring those
images make it to production, and SIG K8s Infra is responsible for ensuring
that production exists in the first place

##### Community Infra Cluster

- We are responsible for ensuring a community-owned GKE cluster is available
  to run apps owned by other SIGs.
- However, we are NOT responsible for ensuring proper functionality of those
  apps. That is left to the SIGs.

e.g. SIG Scalability is responsible for ensuring perfdash.k8s.io displays
valid data

##### Project Infrastructure Budget

- We are responsible for enforcing policy on what is considered in-scope and
  out-of-scope for project infrastructure (and thus, where we spend our
  infrastructure budget)
- Crafting such policy is done in collaboration with the Steering Committee
  (owns project spending) and SIG Architecture (owns Kubernetes definition)
- We MAY delete or scope down infrastructure in the event of unexpected or
  undue spend

e.g. SIG K8s Infra will deny requests to host artifacts for projects that are
formerly part of or adjacent to the Kubernetes project (e.g. helm, cri-o)

##### Public Names

- We are responsible for enforcing policy on what is considered appropriate
  or inappropriate for the names of public-facing entities such as DNS
  records and Google Group names
- Crafting such policy is done in collaboration with the Steering Committee,
  SIG Architecture, and SIG Contributor Experience

e.g. Group names that are used to communicate upon behalf of the project such
as `contributors@kubernetes.io` are vetted by SIG Contributor Experience,
group names that are used for RBAC or IAM bindings are vetted by SIG K8s Infra.

##### Secrets and Credentials

- We are responsible for ensuring secure storage and retrieval of secrets
  such as passwords, tokens, keys, etc.
- However, we are NOT responsible for ensuring the value of those secrets
  is valid.
- We MAY delete or deactivate secrets in the event of a security-related
  incident

e.g. SIG Contributor Experience is responsible for ensuring valid Slack API
credentials exist for proper functioning of slack-infra

##### Security Response

- Overriding all of the above, we MAY revoke, delete, or deactivate
  infrastructure, services or access in the event of a security-related
  incident.
- This depends on responsiveness of the owning SIG, and urgency and severity
  of the incident being responded to

e.g. SIG K8s Infra may force rotation of prow build cluster credentials if
appropriately credentialed members of SIG Testing are not available

### Out of scope

We are not resonsible for code that runs _on_ project infrastructure, with
the exception of:

- subprojects of this SIG (as listed in [`sigs.yaml`], which is more likely
  to be kept up to date than this charter)
- code we share responsibility for (as listed in the [Cross-cutting and
  Externally Facing Processes] section)

We are not responsible for the management of nor in the escalation path for
supporting non-IaaS offerings used by the Kubernetes project that are
managed by other subprojects under other SIGs. For example, problems with
GitHub should be routed to SIG Contributor Experience.

We are not responsible for managing infrastructure which has not yet been
migrated to the CNCF. For example, problems with prow.k8s.io should be routed
to SIG Testing.

## Roles and Organization Management

This sig adheres to the Roles and Organization Management outlined in
[sig-governance] and opts-in to updates and modifications to [sig-governance].

We may revise this portion of the charter when it comes time to talk about
providing a level of support and responsiveness that one might reasonably
expect from a globally distributed open source project.

[sig-governance]: https://git.k8s.io/community/committee-steering/governance/sig-governance.md
[Kubernetes Charter README]: https://git.k8s.io/community/committee-steering/governance/README.md
[lazy consensus]: http://en.osswiki.info/concepts/lazy_consensus

[dev@kubernetes.io]: https://groups.google.com/a/kubernetes.io/group/dev
[sig-k8s-infra@]: https://groups.google.com/a/kubernetes.io/g/sig-k8s-infra
[kubernetes/k8s.io]: https://git.k8s.io/k8s.io
[`sigs.yaml`]: https://git.k8s.io/community/sigs.yaml
