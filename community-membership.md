# Community membership

New community members:

- [**New Contributor**](https://github.com/kubernetes/contrib/issues/1090): a
  couple of PRs; should be welcomed to the community, helped with PR workflow, and
  directed to relevant documentation
- **Active Contributor**: at least 3 merged and/or assigned PRs (which could include documentation
  contributions as well as code), including one in the past month; we have
  [expectations](contributors/devel/community-expectations.md)
  that frequent contributors will assist in our code-review process and with project
  maintenance
  
Established community members:

Established community members are expected to demonstrate their adherence to the principles in this
document, familiarity with project organization, roles, policies, procedures, conventions, etc.,
and technical and/or writing ability. Role-specific expectations, responsibilities, and requirements
are enumerated below.

- **Member**: an active contributor for at least 3 months; at least 10 merged and/or assigned PRs; active enough to be useful
  to assign issues to them and add them to a github team (e.g., for a SIG) for notification
  purposes; trusted enough to run tests on their PRs automatically; can issue "@k8s-bot ok to test"
  for other contributors; if they choose public membership, they get a badge on their github profile;
  should subscribe to kubernetes-dev@googlegroups.com; expected to be familiar with
  project organization, roles, policies, procedures, etc.; should read the [developer
  guide](contributors/devel/README.md); must enable
  [two-factor authentication](https://help.github.com/articles/about-two-factor-authentication/)
- **Reviewer**: org member for at least 3 months; at least 20 merged and/or assigned PRs, including 
  at least 3 as the primary reviewer; familiar enough with some part of the codebase to be in some
  [OWNERS](contributors/devel/owners.md) file as a reviewer (in repos using the bot),
  assigned related PRs, assigned relevant test bugs; responsible for project quality control via
  [code reviews](contributors/devel/collab.md); expected to be responsive to
  review requests as per [community expectations](contributors/devel/community-expectations.md);
  can champion incubator repos; must be nominated by an approver for that part of the codebase,
  with no objections from other approvers; should be added to
  [`kubernetes-reviewers`](https://github.com/orgs/kubernetes/teams/kubernetes-reviewers);
  "read access" to kubernetes repo; get a badge on PR and issue comments; may be asked to
  become a reviewer as a precondition for accepting a large code contribution
- **Approver**: in some [OWNERS](contributors/devel/owners.md) file as an approver, which
  will be needed to get code merged; previously a reviewer for that part of the
  codebase for at least 3 months; at least 30 merged and/or assigned PRs, including at least 10 as
  the primary reviewer; expected to be responsive to review requests as per
  [community expectations](contributors/devel/community-expectations.md); expected to 
  mentor contributors and reviewers; demonstrated sound technical judgement; nominated
  by an area/component owner, with no objections from other owners;  may be asked to
  become an approver as a precondition for accepting a large code contribution
- **Area/Component Owner**: in top-level [OWNERS](contributors/devel/owners.md) file for
  some area/component as an approver; design/proposal approval authority for some area 
  of the project, though escalation is still possible; expected to mentor and guide approvers,
  reviewers, and other contributors; may be asked to become an area/component owner as a precondition
  for accepting the contribution of a new component or other major function
- [**kubernetes-maintainers**](https://github.com/orgs/kubernetes/teams/kubernetes-maintainers):
  approver for some part of the codebase for at least 3 months; on project for at least 1 year;
  at least 50 merged and/or assigned PRs, including at least 20 as the primary reviewer;
  write access to repo (assign issues/PRs, add/remove labels and milestones, edit issues and PRs, edit wiki,
  create/delete labels and milestones); technically can lgtm any PR and cause it
  to be merged by the submit queue, but expected to respect OWNERS files; expected to review PRs, fix bugs, maintain and
  improve health and quality of the project, provide user support, mentor and guide approvers,
  reviewers, and other contributors; must apply to `kubernetes-maintainers@googlegroups.com`, with a
  [Champion](https://github.com/kubernetes/community/blob/master/incubator.md#faq) from the existing
  kubernetes-maintainers members and a Sponsor from Project Approvers, with a summary
  of contributions to the project, current project responsibilities, and links to merged and assigned PRs;
  at least 3 of the maintainers must approve the application, with no objections; the application
  expires after 2 weeks if not enough approvals are granted
- **Project Approvers**: approver in [top-level OWNERS file in kubernetes repo](https://github.com/kubernetes/kubernetes/blob/master/OWNERS);
  de-facto project decision makers; technically can 
  approve virtually any PRs; can sponsor incubator repos; can sponsor maintainers;
  maintainer in good standing for at least 1 year; strong technical vision;
  committed to project's mission and culture; nomination/application process TBD
- [**API Approver**](https://github.com/orgs/kubernetes/teams/api-approvers):
  lead designers of the project, who are familiar with the 
  design, requirements, mechanics, conventions, style, scope, gotchas, etc. 
  of the API; most beta/GA API changes are vetted by the API approvers
- [**API Reviewer**](https://github.com/orgs/kubernetes/teams/api-reviewers):
  contributors familiar with design, requirements, mechanics, conventions, style,
  scope, gotchas, etc. of the API; have written and/or reviewed Kubernetes APIs