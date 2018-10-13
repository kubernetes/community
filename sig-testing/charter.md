# SIG Testing Charter

This charter adheres to the conventions described in the 
[Kubernetes Charter README] and uses the Roles and Organization Management
outlined in [sig-governance].

## Scope

SIG Testing is interested in effective testing of Kubernetes. We focus on
creating and running tools and infrastructure that make it easier for the
community to write and run tests, and to contribute, analyze and act upon
test results. 

We are not responsible for writing, fixing, nor actively troubleshooting the 
project's tests, as this is the responsiblity of the respective test, feature,
and subproject owners. We will however act as an escalation point of last
resort for remediation if it is clear that misbehaving tests are harming the
immediate health of the project.

### In scope

#### Code, Binaries and Services

- Project CI and merge automation via tools such as [prow] and [tide]
- Infrastructure to support running project CI at scale, including tools
  such as [boskos], [ghproxy] and [greenhouse]
- Extraction of test results from GCS and populating a public accessible
  BigQuery dataset via [kettle]
- Display and analysis of test artifacts via tools like [gubernator], 
  [testgrid], [triage] and [velodrome]
- Configuration management of jobs and ensuring they use a consistent
  process via tools such as [job configs], [kubetest]
- Tools that facilitate local testing of kubernetes such as [greenhouse] 
  and [kind]
- Jobs that automate away project toil via [@fejta-bot]
- Ensuring all of the above is kept running on a best effort basis
- Tools, frameworks and libraries that make it possible to write tests against
  kubernetes such as e2e\* or integration test frameworks. 

  \* Note that while we are the current de facto owners of the kubernetes e2e
  test framework, we are not staffed to actively maintain or rewrite it and
  welcome contributors looking to take on this responsibility.

#### Cross-cutting and Externally Facing Processes

- The [Release Team test-infra role] is staffed by a member of SIG Testing, as
  such their responsibilities are within the scope of this SIG, including
  the maintenance of release jobs
- When rolling out changes that may potentially impact the project as a whole
  we consult with SIG Contributor Experience, and follow [lazy consensus] by 
  notifying kubernetes-dev, providing a deadline, and a rationale for the 
  deadline if necessary.
- We reserve the right to halt automation and infrastructure that we own,
  or disable tests that we don't own if the project as a whole is being 
  impacted
- We are actively assisting with the transition of project infrastructure to
  the CNCF and enabling non-Googlers to support this

### Out of scope

- We are not resonpsible for troubleshooting or writing tests or jobs for 
  features or subprojects owned by other SIGs.

## Roles and Organization Management

This sig adheres to the Roles and Organization Management outlined in 
[sig-governance] and opts-in to updates and modifications to [sig-governance].

### Deviations from [sig-governance]

- Chairs also fulfill the role of Tech Lead

### Subproject Creation

Subprojects are created by Tech Leads as defined in the [sig-governance]


[sig-governance]: https://github.com/kubernetes/community/blob/master/committee-steering/governance/sig-governance.md
[Kubernetes Charter README]: https://github.com/kubernetes/community/blob/master/committee-steering/governance/README.md

[@fejta-bot]: https://github.com/fejta-bot
[ghproxy]: https://git.k8s.io/test-infra/ghproxy
[greenhouse]: https://git.k8s.io/test-infra/greenhouse
[gubernator]: http://k8s-gubernator.appspot.com
[job configs]: https://git.k8s.io/test-infra/config/jobs
[kettle]: https://git.k8s.io/test-infra/kettle
[kind]: https://github.com/kubernetes-sigs/kind
[kubetest]: https://git.k8s.io/test-infra/kubetest
[planter]: https://git.k8s.io/test-infra/planter
[prow]: https://prow.k8s.io
[testgrid]: https://testgrid.k8s.io
[tide]: https://prow.k8s.io/tide
[triage]: https://go.k8s.io/triage
[velodrome]: https://velodrome.k8s.io

[lazy consensus]: https://rave.apache.org/docs/governance/lazyConsensus.html
[Release Team test-infra role]: https://git.k8s.io/sig-release/release-team/role-handbooks/test-infra
