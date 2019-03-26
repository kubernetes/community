# WG K8s Infra Charter

This charter adheres to the [wg-governance] guidance, as well as
adheres to the general conventions described in the [Kubernetes
Charter README] and the Roles and Organization Management outlined
in [sig-governance], where applicable to a Working Group.

## Scope

The K8s Infra Working Group is interested in the successful migration of all
project infrastructure from Google (or elsewhere) to the CNCF, such that the
project is able to sustainably operate itself without direct assistance from
entities such as Google or Red Hat.

### Disband criteria

It is our intent to disband once this migration is complete, with ownership
of all code, processes and teams assigned to the appropriate SIGs. If we find
that this is not possible, we will work with the Steering Committee and
respective SIGs to find a more sustainable model (SIG, Committee or Team)

### In scope

#### Code, Binaries and Services

External facing services implemented as subprojects. They often span multiple
SIGs in terms of ownership, hence why this WG is acting as a steward of their
migration.

| Service | SIG(s) | Notes |
| --- | --- | --- |
| DNS | Contribex, Release | Domain Name Services for Kubernetes assets |
| GAE | Testing | Gubernator, Testgerid |
| GCB | Release | Used to build releases |
| GCR | Release | Repository for Container Images |
| GCS / Object Storage | Release, Testing | Buckets for logs, test artifacts, release tarballs, APT, RPM |
| GKE + Stack Driver | Contribex, Release, Testing | Clusters for running bots, utilities, prow, etc |
| Big Query | Testing | Data for test results |
| Website / Blogs | Contribex, CNCF | Communications platform |
| Pool of compute resources for testing | Testing, CNCF, Cloud Providers| GCE, AWS | 

Internal infrastructure which will be necessary to support these, including:

- Credential store to work across teams
- Certificate store for signing certs
- Github repo for storing artifacts like scripts/yamls
  - and for requesting resources (using github issues as tickets?)

Documented policies and processes for how to staff and structure these
subprojects, including:

- naming schemes
- teams and ACL's
- vetting criteria
- on-call / escalation policies

#### Cross-cutting and Externally Facing Processes

##### Deploying Changes

We aspire to remain agile and deploy quickly, while ensuring a disruption-free
experience for project contributors. As such, the amount of notice we provide
and the amount of consensus we seek is driven by our estimation of risk. We
don't currently define risk in terms of objective metrics, so here is a rough
description of the guidelines we follow. We anticipate refining these over
time.

- **Low risk** changes do not break existing contributor workflows, are easy
  to roll back, and impact at most a few project repos or SIGs. These should
  be reviewed by another member of WG K8s infra or the affected SIG(s),
  preferably an approver.

- **Medium risk** changes may impact existing contributor workflows, should be
  easy to roll back, and may impact all of the project's repos. These should
  be shared with the appropriate SIGs, and may require a [lazy consensus]
  issue with [kubernetes-dev@] notice.

- **High risk changes** likely break existing contributor workflows, may be
  difficult to roll back, and likely impact all of the project's repos. These
  require a consultation with SIG Contributor Experience as well as any other
  owning SIGs, and a [lazy consensus] issue with [kubernetes-dev@] notice.

### Out of Scope

- We are not responsible for maintaining infrastructure which has not yet been
  migrated to the CNCF. For example, problems with prow.k8s.io should be routed
  to SIG Testing.

## Roles and Organization Management

- Proposing and making decisions _MAY_ be done without the use of KEPS so long
  as the decision is documented in a linkable medium. We prefer to see written
  decisions and reasoning on the [wg-k8s-infra@] mailing list or as issues
  filed against [kubernetes/k8s.io]. We encourage the use of faster mediums
  such as slack of video conferences to come to consensus.

- It is our intent that each infra-related subproject identified for migration
  must be  staffed / owned by at least 3 volunteers

  - We aspire to follow the same 1/3 maximal representation rules used by the
    Steering Committee, Product Security Committee, and other groups that have
    project-wide impact
  - However, while we are bootstrapping, we consider it acceptable for maximal
    representation concerns to be violated, since this will often be necessary
    for Google-staffed subprojects to divest themselves of the infrastructure.
  - Our plan would be to rectify this when choosing new members or rotating
    old members such that we eventually meet maximal representation criteria

- We plan to follow the model set forth by the Product Security Committee for
  suitable vetting new subproject owners

- Subproject owners must provide additional contact details within the WG, and
  we will need to identify when and how it is appropriate to share these with
  other parts of the project.  Such details include:
  - Alternate e-mails
  - Phone numbers
  - Timezone

- As this is a Working Group, we own no code and cannot create subprojects. We
  will instead identify and petition the appropriate SIG for subproject 
  creation. We will provide guidelines on how the relevant subprojects should
  be staffed, per the above.

[sig-governance]: https://github.com/kubernetes/community/blob/master/committee-steering/governance/sig-governance.md
[wg-governance]: https://github.com/kubernetes/community/blob/master/committee-steering/governance/wg-governance.md
[Kubernetes Charter README]: https://github.com/kubernetes/community/blob/master/committee-steering/governance/README.md
[lazy consensus]: http://en.osswiki.info/concepts/lazy_consensus

[kubernetes-dev@]: https://groups.google.com/forum/#!forum/kubernetes-dev
[wg-k8s-infra@]: https://groups.google.com/forum/#!forum/kubernetes-wg-k8s-infra
[kubernetes/k8s.io]: https://git.k8s.io/k8s.io
