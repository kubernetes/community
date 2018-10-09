# SIG Testing Charter

This charter adheres to the conventions described in the 
[Kubernetes Charter README] and uses the Roles and Organization Management
outlined in [sig-governance].

## Scope

SIG Testing is interested in effective testing of Kubernetes.  We do not write or
troubleshoot the project's tests, but instead focus on tooling that makes it
easier for the community to write and run tests, and to contribute, analyze and
act upon test results.

### In scope

#### Code, Binaries and Services

- Project CI and merge automation via tools such as [prow] and [tide]
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

#### Cross-cutting and Externally Facing Processes

- The Release Team [test-infra role] is staffed by a memer of SIG Testing, as
  such their responsibilities are within the scope of this SIG, including
  the maintenance of release jobs
- When rolling out changes that may potentially impact the project as a whole
  we consult with SIG Contributor Experience, and follow [lazy consensus] by 
  notifying kubernetes-dev, providing a deadline,and a rationale for the 
  deadline if necessary.
- We reserve the right to halt automation and infrastructure that we own,
  or disable tests that we don't own if the project as a whole is being 
  impacted
- We are actively assisting with the transition of project infrastructure to
  the CNCF and enabling non-Googlers to support this

### Out of scope

- We are not resonpsible for troubleshooting or writing tests or jobs for 
  features owned by other SIGs
- We are not responsible for ongoing maintenance of kubernetes' e2e test 
  framework

## Roles and Organization Management

This sig follows adheres to the Roles and Organization Management outlined in 
[sig-governance] and opts-in to updates and modifications to [sig-governance].

### Deviations from [sig-governance]

- Chairs also fulfill the role of Tech Lead

### Subproject Creation

Subprojects are created by Tech Leads as defined in the [sig-governance]


[sig-governance]: https://github.com/kubernetes/community/blob/master/committee-steering/governance/sig-governance.md
[sig-subprojects]: https://github.com/kubernetes/community/blob/master/sig-YOURSIG/README.md#subprojects
[Kubernetes Charter README]: https://github.com/kubernetes/community/blob/master/committee-steering/governance/README.md

[gubernator]: http://k8s-gubernator.appspot.com
[kettle]: https:/git.k8s.io/test-infra/kettle
[prow]: https://prow.k8s.io
[testgrid]: https://testgrid.k8s.io
[tide]: https://prow.k8s.io/tide
[triage]: https://go.k8s.io/triage
[job configs]: https://git.k8s.io/test-infra/config/jobs
[kubetest]: https://git.k8s.io/test-infra/kubetest
[kind]: https://github.com/kubernetes-sigs/kind
[@fejta-bot]: https://github.com/fejta-bot
[greenhouse]: https://git.k8s.io/test-infra/greenhouseu

[lazy consensus]: https://rave.apache.org/docs/governance/lazyConsensus.html
