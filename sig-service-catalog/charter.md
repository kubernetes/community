# SIG Service Catalog Charter

This charter adheres to the conventions described in the [Kubernetes Charter
README](https://github.com/kubernetes/community/blob/master/committee-steering/governance/README.md).

## Scope

Service Catalog is a Kubernetes extension project that implements the [Open
Service Broker API](https://www.openservicebrokerapi.org/) (OSBAPI). It enables
application developers to provision cloud services from within Kubernetes and
integrates configuration and discovery of those services into Kubernetes
resources.

### In scope

See the [service-catalog SIG entry](https://github.com/kubernetes/community/tree/master/sig-service-catalog).

This SIG’s main goals are:
- Support, and adhere to, the Platform requirements of the [OSBAPI
  specification](https://github.com/openservicebrokerapi/servicebroker/blob/master/spec.md).
- Provide a UX for Kubernetes users that is consistent with both the OSB API
  specification and traditional Kubernetes user interactions.
- Align with the OSBAPI specification as changes are made.
- Provide feedback (bugs or feature requests) to the [OSBAPI WG]](https://www.openservicebrokerapi.org/).

### Code, Binaries and services

- [Source Repository](https://github.com/kubernetes-incubator/service-catalog)
  - See [OWNERS](https://raw.githubusercontent.com/kubernetes-incubator/service-catalog/master/OWNERS) for who has access.
- [Image Repository](https://quay.io/organization/kubernetes-service-catalog)
  - Canary builds are published on pushes to master.
  - Release builds (and latest) are published on tags.
  - Chairs have access to manage this repository.
- [Helm Repository](https://svc-catalog-charts.storage.googleapis.com)
  - Charts are manually published after each release.
  - Managed by Vic Iglesias (Google), @viglesias on the kubernetes slack.
- [svc-cat.io](https://svc-cat.io)
  - Published on pushes to master.
  - Site hosted with [Netlify](https://app.netlify.com/sites/svc-cat/overview).
  - Chairs and interested maintainers have access to manage this site.
- [CLI Binary Hosting](https://svc-cat.io/docs/install/#manual)
  - Canary builds are published on pushes to master.
  - Release builds (and latest) are published on tags.
  - Files hosted on Azure blob storage.
  - Azure account managed by Carolyn Van Slyck (Microsoft) and Aaron Schlesinger
    (Microsoft).
- [Travis](https://travis-ci.org/kubernetes-incubator/service-catalog)
  - Runs the CI builds.
  - Maintainers have access.
- [Jenkins](https://service-catalog-jenkins.appspot.com/)
  - Runs end-to-end tests on a live cluster.
  - Server managed by Michael Kibbe (Google).

### Out of scope

The following, non-exhaustive, items are out of scope:
- Operation of OSBAPI Service Brokers.

## Roles
This SIG's charter deviates from the
[sig-governance](https://github.com/kubernetes/community/blob/master/committee-steering/governance/sig-governance.md)
roles. We do not have the Tech Lead role, and have a honorary Emeritus Chair role.

- [Maintainers](https://github.com/orgs/kubernetes-incubator/teams/maintainers-service-catalog/members)
  - Maintainer is equivalent to the standard [Kubernetes definition of
    Approver](https://github.com/kubernetes/community/blob/master/community-membership.md#approver).
  - Responsible for reviewing pull requests, and approving pull requests for merge.
  - Responsible for technical planning and stewardship of the project.
  - New maintainers may be nominated by another maintainer, and accepted via lazy
    two-thirds resolution amongst the chairs.
  - Maintainers may be nominated for removal from their position by a chair,
    and accepted via lazy two-thirds resolution amongst the chairs.
  - Maintainers may propose changes to this charter at any time, by submitting a
    pull request and then notifying the SIG mailing list, to be accepted via
    lazy two-thirds resolution amongst the chairs.

- Chairs
  - Chairs are expected to perform the role of maintainer, in addition to their chair responsibilities.
  - Chairs are listed in our [SIG
    definition](https://github.com/kubernetes/community/tree/master/sig-service-catalog#chairs).
  - Responsible for project administration activities within the SIG and are
    non-technical in nature, such as organizing the weekly meetings.
  - A chair does not have more rights, or votes, than a maintainer.
  - Responsible for reporting the SIG’s status to the appropriate Kubernetes
    leadership teams.
  - All decisions amongst chairs are made using lazy consensus with a fallback
    to a 2/3 majority vote (lazy two-thirds resolution).
    This process is used for all decisions, such as changing chairs/maintainers
    or modifying this charter.
  - Chairs may nominate a new chair at any time, to be accepted via lazy
    two-thirds resolution amongst the chairs.
  - Chairs may decide to step down at any time.
  - Chairs must remain active in the role and may be removed from the position
    via lazy two-thirds resolution amongst the chairs, if they are unresponsive
    for > 3 months or are not proactively working with other chairs to fulfill
    responsibilities.
  - There is no set number of Chairs.

- Emeritus Chairs ([Inspired by the Helm
  Project](http://technosophos.com/2018/01/11/introducing-helm-emeritus-core-maintainers.html))
  - A chair who steps down may choose to take an Emeritus Chair title. This
    confers honor on the recipient and allows them to remain associated with the
    SIG in acknowledgement of their significant contributions.
  - Those who attain this title are no longer expected to attend the weekly
    meetings, share in the issue queue triage rotation, vote on policy or
    architectural issues, or review pull requests.
  - They are [listed in our documentation](https://svc-cat.io/community/#leadership)
    as Emeritus Chairs, and we will continue to invite them to participate in
    related events, such as KubeCon.

- Security Contacts
  - Are a contact point for the Product Security Committee to reach out to for
    triaging and handling of incoming issues.
  - Must be a maintainer.
  - Must accept and adhere to the Kubernetes [Embargo
    Policy](https://git.k8s.io/security/private-distributors-list.md#embargo-policy).
  - Defined in
    [SECURITY_CONTACTS](https://github.com/kubernetes-incubator/service-catalog/blob/master/SECURITY_CONTACTS)
    file.

## Organizational management

- SIG meets every week on Zoom at 1 PM PST on Mondays
    - Agenda
      [here](https://docs.google.com/document/d/17xlpkoEbPR5M6P5VDzNx17q6-IPFxKyebEekCGYiIKM/edit#).
    - Anyone is free to add new agenda items to the doc.
    - Recordings of the calls are made available [here](https://goo.gl/ZmLNX9).
- SIG members explicitly representing the group at conferences (SIG progress
  reports, deep dives, etc.) should make their slides available for perusal and
  feedback at least 2 week in advance.
- [Working
  groups](https://github.com/kubernetes-incubator/service-catalog/wiki/Working-Groups)
  can be initiated by any member. To create a new one, add the topic to the
  weekly call’s agenda for discussion.
  - These are not the same as cross-SIG working groups.
  - Working groups exist for an unspecified period of time, so that interested
    members can meet to discuss and solve problems for our SIG.

### Project management
- [Milestones](https://github.com/kubernetes-incubator/service-catalog/milestones)
  are defined by SIG maintainers.
- Anyone is free to request a discussion of the milestones/plans during
  a weekly call.
- Weekly releases are typically done on Thursdays, and any member who has
  maintainer rights is free to initiate it. _Friday releases are strongly
  discouraged._
- Major releases are planned and discussed among the SIG members during regular
  weekly calls.
- The release process is defined
  [here](https://github.com/kubernetes-incubator/service-catalog/wiki/Release-Process).
- Anyone can request to work on an issue by commenting on it with `#dibs`.


### Technical processes
- All technical decisions are made either through issues, pull requests or
  discussions during the weekly SIG meeting. Major decisions should be
  documented in an issue or pull request.
- There is no requirement that all pull requests have an associated issue.
  However, if the PR represents a significant design decision then it is
  recommended that it be discussed among the group to avoid unnecessary coding
  that might not be accepted.
- While everyone is encouraged to contribute to the discussion of any topic,
  ultimately any changes made to the codebase must be approved by the
  maintainers.
- Pull requests are required to have at least 2 maintainers approve them.
- Pull requests that are labeled with `do-not-merge/hold` or have an on-going
  discussion must not be merged until all concerns are addressed.
- Disagreements are resolved via lazy consensus. In the event that a common
  decision cannot be made, then a vote among the maintainers will be taken.
  Simple majority of votes cast (>50%) wins.
