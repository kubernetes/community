# Kubernetes Project Maintainership

Kubernetes is a big project with many contributors, the project organization is
a crucial task and the project maintainers are the members that are trusted
with that task, driving the project, improving the health of the project,
performing quality control and being responsible over one or more components in
the project. A Kubernetes maintainer is a contributor who has contributed/is
contributing to the project regularly and a developer that demonstrated good
technical judgement. The following is a list of requirements to be a Kubernetes
maintainer.

Before continuing, it should be noted that we value every contributions, however
small it may be, even a comment on an issue a valuable contribution to us.
Maintainership requires time commitment and community trust. You don't have to
be a maintainer to have a meaningful impact on Kubernetes!

## Maintainership Requirements

In order to become a maintainer, a contributor must satisfy all of the
requirements listed below.

### 1. Enable 2-Factor Account Authentication for GitHub

See GitHub's documentation about [two-factor authentication](https://help.github.com/articles/about-two-factor-authentication/) for details.

### 2. Be an active contributor

Active contributors are considered to be anyone who meets any of the following criteria:

- Sent more than two pull requests (PRs) in the previous one month, or more than
  20 PRs in the previous year.
- Filed more than three issues in the previous month, or more than 30 issues in
  the previous 12 months.
- Commented on more than pull requests in the previous month, or more than 50
  pull requests in the previous 12 months.
- Marked any PR as LGTM in the previous month.
- Have collaborator permissions in the Kubernetes GitHub organization.

See [Kubernetes Community Expectations](https://github.com/kubernetes/kubernetes/blob/master/docs/devel/community-expectations.md) for more details.

**TODO:** Move the definition of an active contributor to here or remove this
section and link to expectations.

#### Contribution types

Contributions are not limited to code. Other contributions include PR reviews,
discussion participations, test ownership, SIG leadership, documentation, etc.
As we improve our project tooling, we'll be able to measure and analyze other
types of contributions and take those into account when evaluating a maintainer
application, however maintainers still need to be familiar with the design and
implementation of Kubernetes.

### 3. Be experienced in the project

You need to have more than 50 contributions total and project participation for
more than a year in the Kubernetes organization.

**Note:** A year time requirement will be replaced once the contributor ladders
are introduced.

### 4. Find 3 senior maintainer sponsors

As you contribute to Kubernetes, you'll be collaborating with different
community members. Having the sponsorship of a senior project maintainer means
that they have seen your work and they believe that you'll be a good addition as
a project maintainer.

**Motivation:** We want the sponsorships to be meaningful, so we require the sponsors to be
maintainers that collaborated with you and we also don't want to burden the
senior maintainers otherwise they'll be reluctant to sponsor anyone. Therefore,
rather than having a mentorship process where a new maintainer is guided by a
senior member, we want to see the capabilities of the candidate before being
accepted.

**TODO:** Detail the meaning of a senior project maintainer with a requirement
along the lines of more than 6 months experience as a project maintainer. Also
related to contributor ladder.

## Maintainership Application

Once you have satisfied the listed maintainership requirements, you can submit
your application to kubernetes-wg-contribex@googlegroups.com wherein the
application will be discussed by the contributor leads (top-level **OWNERS**) and
put to vote. If the vote result is negative, do not be discouraged! You will be
given detailed feedback so that you can improve yourself and apply again. We
want to empower our community and cultivate a trust relationship that's why we
carefully evaluate abilities and experience of each candidate.

### Application Template

A rudimentary application template is as follows:

```yaml
github-id: foo-bar

# List of senior maintainers that agreed to sponsor you.
sponsors:
  - <github-id>
  - ...
  - ...
  
# Links to enumerated list of your contributions for ease of review
# Note: It'll get easier to create such lists as we improve our contribution measurement tooling
# Note: The following template is just and example, you can use any formatting, however the important
#       thing is to list your contributions and classify them by repository/type/theme to present a
#       digestable view of your efforts to the reviewers of your application.
contributions:
  - X merged commits <github-link>
    - X/3 in sig-foo
    - 2X/3 in overall documentation
  - Y reviewed PRs <github-link>
    - Y in kubectl

# List of previous and current responsibilities in the Kubernetes organization.
responsibilities:
  current:
    - sig-X contributor
  previous:
    - v1.4 release czar
    
# List of additional repositories that you need write access to if there is any
additional-write-access-to:
  - kubernetes/client-go
  - kubernetes/contrib
  
# Optional access duration, omitable if there is no planned end date.
duration: the v1.6 release | for 3 months | etc.
```

## What we expect from our maintainers

First and foremost, be a good person, as cheeky as it sounds, you are now a
member of the Kubernetes project and thus represent the project and your fellow
developers. We value our community tremendously and we'd like to keep
cultivating a friendly and collaborative environment for our
members/contributors/users. For more,
see
[CNCF Code of Conduct](https://github.com/cncf/foundation/blob/master/code-of-conduct.md).

Good technical judgement, and knowledge over one or more components of
Kubernetes, it might be a stand-alone project under the Kubernetes GitHub
organization (e.g. [minikube](github.com/kubernetes/minikube)) or it might be a
main component of Kubernetes such as the Kubernetes scheduler or kubectl tool.
We determine the maintainer applicant's judgement via the sponsorship
requirement. For example, a senior maintainer might decide to sponsor an
applicant after they witnessed the applicants good technical judgement during a
design discussion, a feature implementation or by simply witnessing the way they
interact with other contributors.

## Inactive Maintainers

Maintainers deemed inactive after being inactive for 60 days in the project or
maintainers that decide to step down will be retired. Retired maintainers might
decide to come back to the project at a later point. This means that if an
retired maintainer wants to return, he will be able to regain his access once he
becomes an active contributor again.

**Motivation:** Since maintainers have already proved themselves by becoming a
maintainer in the first place, we only require that the maintainer familiarizes
himself with the project again by becoming an active contributor again.
Kubernetes as a project is evolving very fast and it requires some time to get
used to the project changes whether it is around contributor tooling, project
structure or the codebase itself.

**TBD**: The inactivity period (60 days) is subject to change.
