# SIG Cluster Lifecycle Charter

## Scope

SIG Cluster Lifecycle’s objective is to simplify creation, configuration, upgrade, downgrade, and teardown of Kubernetes clusters and their components.

### In scope

The following topics fall under ownership of this SIG:

- Improving the Kubernetes user experience for cluster administration.
- Tools that assist in the creation, configuration, upgrade, downgrade, and teardown of Kubernetes control plane components.
- Portable APIs for provisioning, configuration, upgrade/downgrade, and de-provisioning of nodes.
- Tools that assist in management of configuration of Kubernetes components.
- The configuration of core add-ons that are required for cluster bootstrapping.

#### Code, Binaries and Services

- Everything that falls in the scope of the SIG.
- Tools that are provider specific implementation for infrastructure management.
- Core add-ons (e.g. DNS) that are required for cluster bootstrapping.

#### Cross-cutting and Externally Facing Processes

- The SIG recommends and verifies compatibility of critical cluster add-ons for networking, network policy, service discovery, etc. The SIG maintains the health check of container images for some add-ons that are required for cluster bootstrapping. While the SIG could provide support to users that have add-on related issues, the SIG can decide to delegate issues to the add-on maintainers or other SIGs.
- The SIG collaborates regularly with SIG Auth in an effort to follow best practices in order to promote secure default clusters.
- The SIG co-owns cloud provider specific code related to cluster and machine provisioning with the respective SIGs for each cloud provider but does not own the cloud controller manager or any other provider specific code.
- The SIG ensures that the Kubernetes "release informing" and "release blocking" E2E test jobs that it maintains are in good health.

### Out of scope

- Networking related issues (see [sig-network](../sig-network)).
- User interface, or user experience, issues other than cluster bootstrapping or management (see [sig-ui](../sig-ui) and  [sig-cli](../sig-cli)).
- Node related issues (see [sig-node](../sig-node)).
- Kubernetes control plane issues:
   - Control plane component related issues (see [sig-api-machinery](../sig-api-machinery) and [sig-scheduling](../sig-scheduling)).
- Deployment and lifecycle issues related to user application deployments (see [sig-apps](../sig-apps)).

## Roles and Organization Management

This SIG adheres to the Roles and Organization Management outlined in [sig-governance]
and opts-in to updates and modifications to [sig-governance].

### Additional responsibilities of Chairs

- Selecting features for a given milestone.
- Hosting the weekly SIG meeting, ensure that recordings are uploaded in a timely fashion.
- Ensuring that the breakout sessions the SIG hosts during the week have chairs.
- Organizing SIG sessions at KubeCon events (intro / deep dive sessions).
- Creating roadmaps for a given year or release, or reviewing and approving technical implementation plans (e.g. KEPs) in coordination with both SIG Cluster Lifecycle contributors and other SIGs.

### Deviations from [sig-governance]

As SIG Cluster Lifecycle contains a number of subprojects, following deviations from [sig-governance] exists:

- Subproject creation follow the process described in [Subproject creation](#subproject-creation).
- Each subproject can nominate subproject leads by a simple majority vote of subproject top-level maintainers, SIG Chairs and TLs.
- In there is not a formally appointed subproject lead, the top-level OWNERS defines the maintainers that 
  are responsible in front of the SIG for the project (in this document, we are using subproject leads for both).
- Subproject leads must be subscribed to the SIG mailing list and watch out for our announcement emails.
- The SIG has empowered subproject leads with a number of additional responsibilities, 
  including but not limited to the responsibility for determining the subproject release cadence and producing releases.
- Subprojects archival follow the process described in [Subproject creation](#subprojects-archival).

Please note that, despite a certain degree of operational autonomy, SIG Cluster Lifecycle subprojects are still 
expected to be continuously engaged and active in the broader SIG activities. As such:

- Subproject leads or members are expected to show up periodically at the SIG meetings, celebrate achievements, 
  share new ideas as well as issues or problems the subproject team is facing.
- Subproject leads should actively contribute to the preparation of the SIG annual report.

### Subproject creation

SIG Cluster Lifecycle will consider the creation of a new subproject when a group of people actively engaged 
in the SIG will propose an innovative idea in scope of the SIG charter, ideally filling some of the void in 
the SIG vision or extending it in a meaningful way.

The SIG also welcomes discussion about  innovative ideas originated from other SIGs as well as from other contributors.

Like any significant change in Kubernetes, new subprojects should be backed by a KEP detailing 
the subprojects goals, non-goals, and the initial roadmap.

The KEP should be signed by a group of contributors volunteering for the subprojects initial work as well as 
for future maintenance; in the best interest of the new subproject, this group is expected to be adequate 
to the effort and heterogeneous.

When considering submitting a new subproject proposal, please be aware that:

- KEP process for new subprojects should follow well established practices:
  - preliminary discussion in the SIG meetings
  - [optional] draft in a Google doc for easier collaboration
  - PR to https://github.com/kubernetes/enhancements
    - Official [KEP template](https://github.com/kubernetes/enhancements/tree/master/keps/NNNN-kep-template) should be used;
      template sections that do not apply to a new subproject should be dropped, e.g. "Upgrade / Downgrade Strategy"
    - Subprojects KEPs should not be tied to any Kubernetes release (no stage or milestone in kep.yaml).
    - Prior art [KEP 2495 - Kubernetes Cluster Management API](https://github.com/kubernetes/enhancements/tree/master/keps/sig-cluster-lifecycle/clusterapi/2495-Kubernetes-Cluster-Management-API) (Cluster API)
  - Review
  - Approval by a [super-majority] of SIG chairs and TLs; the approval vote should be recorded on the KEP PR.
- Creating a KEP is not the last step to create a subproject, other steps should follow including but not limited to:
  updating sigs.yaml, creating the kubernetes/org repo request, following the [rules for donated repositories] (CLA and license history).
- Subprojects that don't bring contributors into the SIG Cluster Lifecycle make the SIG sustainability problem worse, not better.
- SIG Cluster Lifecycle is not obligated to host every project in its technical space, and "there is nowhere else to put it" 
  is usually not a good reason to create a new subproject.
- Projects looking for a vendor-neutral home to build a community may be better served by applying to the CNCF Sandbox program.

Also, the SIG is committed to be a good citizen in the Kubernetes community, and as such before accepting a new subproject 
we should always consider that hosting is not free: CI and infra cost, org membership, security response, release duties, GitHub administration.

### Subproject Governance

Subprojects governance must follow the guideline defined in [sig-governance], including but not limited to:

- A subproject must have one or more dedicated OWNERS files. Please note that the top-level OWNERS defines
  the maintainers that are responsible in front of the SIG for the project.
- Subprojects must define security contacts, see [sig-governance][security contact].
- Subprojects should run their own focused meetings and maintain separate agendas.
- The subproject owners are responsible for ensuring that the issues and bugs are triaged in a timely manner.
- The subproject owners are responsible for ensuring that pull requests are reviewed in a timely manner.
- The subproject owners are responsible for determining the subproject release cadence and producing releases.
- A subproject is responsible for ensuring the quality of release artifacts and documentation.

### Subprojects Archival

Subprojects not active anymore will be archived following [removing repositories] guidelines.
In alternative to archival, also [maintenance mode], should be considered.

Similarly, archival will be considered if a subproject team drifts away significantly from the broader SIG activity.

The following signals might be considered to determine if a subproject should be archived:
- Subprojects not engaging in SIG activities such as filling up surveys, not following the SIG ML announcements, SIS slack
  discussions, office hours meetings etc.
- Subprojects failing to contribute to the SIG annual report.
- Subprojects not actively maintaining the OWNERS files, SECURITY_CONTACTS.
- Subprojects not actively engaging contributors and users e.g. in office hours, slack channels, etc. 
- Subprojects not actively triaging issues, not actively reviewing and merging PRs.
- Subprojects without a release in the last 6 months.
- Subprojects failing to keep up with Kubernetes releases or with releases of the projects they depend on (e.g. Cluster API for providers).

If a subproject seems inactive or eligible for archival, SIG leads should perform a due diligence with subproject leads.

The decision to archive a subproject requires a [super-majority] vote of the SIG Chairs and Tech Leads.
Subproject owners can escalate an archival decision to the Steering Committee.

After the vote, an email should be sent to the SIG Cluster Lifecycle mailing list (sig-cluster-lifecycle@kubernetes.io) and
to the developer mailing list (dev@kubernetes.io) outlining the decision.

Archival must start after one month grace period or more from the vote, to mitigate impacts on user not aware of 
the project being inactive.

[sig-governance]: https://github.com/kubernetes/community/blob/master/committee-steering/governance/sig-governance.md
[security contact]: https://github.com/kubernetes/community/blob/main/committee-steering/governance/sig-governance.md#security-contact
[super-majority]: https://en.wikipedia.org/wiki/Supermajority#Two-thirds_vote
[sigs.yaml]: https://github.com/kubernetes/community/blob/master/sigs.yaml
[Kubernetes Charter README]: https://github.com/kubernetes/community/blob/master/committee-steering/governance/README.md
[rules for donated repositories]: https://github.com/kubernetes/community/blob/main/github-management/kubernetes-repositories.md#rules-for-donated-repositories
[removing repositories]: https://github.com/kubernetes/community/blob/main/github-management/kubernetes-repositories.md#removing-repositories
[maintenance mode]: https://github.com/kubernetes/community/blob/main/github-management/kubernetes-repositories.md#maintenance-mode