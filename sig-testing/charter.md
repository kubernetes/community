# SIG Testing Charter

This charter adheres to the conventions described in the
[Kubernetes Charter README] and uses the Roles and Organization Management
outlined in [sig-governance].

## Scope

SIG Testing is interested in effective testing of Kubernetes and automating
away project toil. We focus on creating and running tools and infrastructure
that make it easier for the community to write and run tests, and to
contribute, analyze and act upon test results.

Although we are not responsible for ongoing test maintenance (see
[Out of Scope] below), we will act as an escalation point of last resort for
remediation if it is clear that misbehaving tests are harming the immediate
health of the project.

### In scope

#### Code, Binaries and Services

- Project CI and workflow automation via tools such as [prow] and [tide]
- Infrastructure to support running project CI at scale, including tools
  such as [boskos], [ghproxy] and [greenhouse]
- Providing a place and schema in which to upload test results for
  contributors who wish to provide additional test results not generated
  by the project's CI
- Extraction, display and analysis of test artifacts via tools like
  [gubernator], [kettle], [testgrid], [triage] and [velodrome]
- Configuration management of jobs and ensuring they use a consistent
  process via tools such as [job configs], [kubetest]
- Tools that facilitate configuration management of github such as
  [peribolos] and [label_sync]
- Tools that facilitate local testing of kubernetes such as [greenhouse]
  and [kind]
- Jobs that automate away project toil, such as [periodic jobs that run as
  @fejta-bot]
- Ensuring all of the above is kept running on a best effort basis
- Tools, frameworks and libraries that make it possible to write tests against
  kubernetes such as e2e\* or integration test frameworks.

  \* Note that while we are the current de facto owners of the kubernetes e2e
  test framework, we are not staffed to actively maintain or rewrite it and
  welcome contributors looking to take on this responsibility.

#### Cross-cutting and Externally Facing Processes

##### Ongoing Support

- We actively collaborate with SIG Contributor Experience, often producing
  tooling that they are responsible for using to implement policies and
  processes that they own, e.g. the Github Administration subproject uses
  [peribolos] and [label_sync] to reduce the toil involved
- We reserve the right to halt automation and infrastructure that we own,
  or disable tests that we don't own if the project as a whole is being
  impacted
- We are actively assisting with the transition of project infrastructure to
  the CNCF and enabling non-Googlers to support this

##### Deploying Changes

We aspire to remain agile and deploy quickly, while ensuring a disruption-free
experience for project contributors. As such, the amount of notice we provide
and the amount of consensus we seek is driven by our estimation of risk. We
don't currently define risk in terms of objective metrics, so here is a rough
description of the guidelines we follow. We anticipate refining these over
time.

- **Low risk** changes do not break existing contributor workflows, are easy
  to roll back, and impact at most a few project repos or SIGs. These should
  be reviewed by another member of SIG Testing or the affected SIG(s),
  preferably an approver.

- **Medium risk** changes may impact existing contributor workflows, should be
  easy to roll back, and may impact all of the project's repos. These should
  be shared with SIG Contributor Experience, may require a lazy consensus
  issue with [kubernetes-dev@] notice.

- **High risk changes** likely break existing contributor workflows, may be
  difficult to roll back, and likely impact all of the project's repos. These
  require a consultation with SIG Contributor Experience, and a lazy consensus
  issue with [kubernetes-dev@] notice.

### Out of Scope

- We are not responsible for writing, fixing nor actively troubleshooting tests
  for features or subprojects owned by other SIGs
- We are not responsible for ongoing maintenance of the project's CI Signal,
  as this is driven by tests and jobs owned by other SIGs. We do however have
  an interest in producing tools to help improve the signal.

## Roles and Organization Management

This sig adheres to the Roles and Organization Management outlined in
[sig-governance] and opts-in to updates and modifications to [sig-governance].

### Deviations from [sig-governance]

- Chairs also fulfill the role of Tech Lead
- Proposing and making decisions _MAY_ be done without the use of KEPS so long
  as the decision is documented in a linkable medium. We prefer to use issues
  on [kubernetes/test-infra] to document technical decisions, and mailing list
  threads on [kubernetes-sig-testing@] to document administrative decisions on
  leadership, meetings and subprojects.
- We do not consistently review sig-testing testgrid dashboards as part of our
  meetings

### Subproject Creation

Subprojects are created by Tech Leads following the process defined in [sig-governance]


[sig-governance]: https://github.com/kubernetes/community/blob/master/committee-steering/governance/sig-governance.md
[Kubernetes Charter README]: https://github.com/kubernetes/community/blob/master/committee-steering/governance/README.md
[lazy consensus]: http://en.osswiki.info/concepts/lazy_consensus

[periodic jobs that run as @fejta-bot]: https://git.k8s.io/test-infra/config/jobs/kubernetes/test-infra/fejta-bot-periodics.yaml
[boskos]: https://git.k8s.io/test-infra/boskos
[ghproxy]: https://git.k8s.io/test-infra/ghproxy
[greenhouse]: https://git.k8s.io/test-infra/greenhouse
[gubernator]: http://k8s-gubernator.appspot.com
[job configs]: https://git.k8s.io/test-infra/config/jobs
[kettle]: https://git.k8s.io/test-infra/kettle
[kind]: https://github.com/kubernetes-sigs/kind
[kubetest]: https://git.k8s.io/test-infra/kubetest
[label_sync]: https://git.k8s.io/test-infra/label_sync
[peribolos]: https://git.k8s.io/test-infra/prow/cmd/peribolos
[planter]: https://git.k8s.io/test-infra/planter
[prow]: https://prow.k8s.io
[testgrid]: https://testgrid.k8s.io
[tide]: https://prow.k8s.io/tide
[triage]: https://go.k8s.io/triage
[velodrome]: https://velodrome.k8s.io

[Release Team test-infra role]: https://git.k8s.io/sig-release/release-team/role-handbooks/test-infra
[kubernetes-dev@]: https://groups.google.com/forum/#!forum/kubernetes-dev
[kubernetes-sig-testing@]: https://groups.google.com/forum/#!forum/kubernetes-sig-testing
[kubernetes/test-infra]: https://git.k8s.io/test-infra
