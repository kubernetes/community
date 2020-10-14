---
title: "Contributing to Kubernetes"
weight: 4
description: |
  An entrypoint to getting started with contributing to the Kubernetes project.
---

# Contributing

- [Communication](#communication)
- [GitHub workflow](#github-workflow)
- [Open a Pull Request](#open-a-pull-request)
- [Code Review](#code-review)
- [Best Practices](#best-practices)
- [Testing](#testing)
- [Security](#security)
- [Documentation](#documentation)
- [Issues Management or Triage](#issues-management-or-triage)

Kubernetes is open source, but many of the people working on it do so as their day 
job. In order to avoid forcing people to be "at work" effectively 24/7, this 
document establishes some semi-formal protocols around development. Hopefully, 
these rules make things go more smoothly. If you find that this is not the case, 
please reach out to the [Contributor Experience SIG] with any concerns.

As a potential contributor, your changes and ideas are welcome anytime.
Please do not ever hesitate to ask a question or send a pull request. 

Check out the [community guiding principles] for the key concepts you will
need to understand in order to contribute to a large, collaborative open source 
project such as Kubernetes.

For a quick reference on contributor resources, please remember to bookmark the 
[contributor cheatsheet].

## Communication

It is best to contact a [SIG] for issues related to that SIG's topic. A SIG 
will be able to help you much more quickly than raising a general question would.
For general questions and troubleshooting, use the [standard lines of 
communication] and work through the [troubleshooting guide].

### GitHub workflow

To check out code to work on, please refer to [the GitHub Workflow Guide].

The Kubernetes-specific GitHub workflow is very comprehensive and detailed. The 
full workflow for a pull request is available here: 

* [Kubernetes-specific GitHub workflow]

 This document details an example of a typical pull request workflow scenario.

#### Opening a Pull Request

Pull requests are often called a "PR". Kubernetes generally follows the standard 
[github pull request] process, but there are some Kubernetes-specific 
(and sometimes SIG specific) differences to be aware of. The first difference is 
that [the bot] will begin applying structured labels to your pull request.

The bot may also make some helpful suggestions for commands to run in your PR to 
facilitate review. These `/command` options can be entered in comments to trigger 
auto-labeling and notifications. To learn more about what commands the bot uses,
feel free to review the [command reference documentation].

Some common issues new contributors face when making a pull request are:

* Not having correctly signed the CLA ahead of your first PR. See the [CLA page]
for troubleshooting help. Note that in some cases you may need to file a ticket 
with the CNCF to resolve a CLA problem.  
* Finding the right SIG or reviewer(s) for the PR (see [Code Review] section) and 
following any SIG or repository specific contributing guidelines (visit the
 [Learn about SIGs] section for more information)
* Dealing with test cases which fail on your PR, unrelated to the changes you 
introduce (see [Test Flakes] for more information) 
* Not following [scalability best practices]
* Include mentions (like @person) and [keywords] which could close the issue 
(such as 'fixes #xxxx') in commit messages.

### Code Review

For a brief description of the importance of code review, please read 
[On Code Review]. 

To make it easier for your PR to get reviewed, reviewers will need you to:

* Follow the project's [coding conventions]
* Write [good commit messages]
* Break large changes into a logical series of smaller patches which individually 
make easily understandable changes, and when combined into a whole, solve a 
broader issue
* Label PRs with appropriate SIGs and reviewers: To do this, read the messages 
[the bot] sends you to guide you through the PR process

Reviewers are highly encouraged to revisit the [Code of Conduct] as well as 
[community expectations] and must go above and beyond to promote a collaborative, 
respectful community.

When reviewing PRs from others [The Gentle Art of Patch Review] suggests an 
iterative series of focuses which is designed to lead new contributors to 
positive collaboration without inundating them initially with nuances:

* Is the idea behind the contribution sound?
* Is the contribution architected correctly?
* Is the contribution polished?

Note: if your pull request isn't getting enough attention, you can use the 
[#pr-reviews] channel on Slack to get help finding reviewers.

#### Best practices

* Write clear and meaningful Git commit messages.
* If the PR will *completely* fix a specific issue, include `fixes #123` in the 
PR body (where 123 is the specific issue number the PR will fix. This will 
automatically close the issue when the PR is merged.
* Make sure you don't include `@mentions` or `fixes` keywords in your git commit 
messages. These should be included in the PR body instead.
* When you make a PR for small change (such as fixing a typo, style change, or 
grammar fix), please squash your commits so that we can maintain a cleaner gGit 
history.
* Make sure you include a clear and detailed PR description explaining the 
reasons for the changes, and ensuring there is sufficient information for the 
reviewer to understand your PR.
* Additional Readings: 
  * [Closing Issues Via Pull Request]
  * [Squashing Git Commits]

### Testing

Testing is the responsibility of all contributors and is in part owned by all 
SIGs, but is also coordinated by [sig-testing]. Refer to the [Testing Guide] for 
more information.

There are multiple types of tests. The location of the test code varies with 
type, as do the specifics of the environment needed to successfully run the test:

* Unit: These confirm that a particular function behaves as intended. Golang 
includes a native ability for unit testing via the [testing] package. Unit test 
source code can be found adjacent to the corresponding source code within a given 
package. For example: functions defined in [kubernetes/cmd/kubeadm/app/util/version.go] 
will have unit tests in [kubernetes/cmd/kubeadm/app/util/version_test.go]. These 
are easily run locally by any developer on any OS.
* Integration: These tests cover interactions of package components or 
interactions between kubernetes components and some other non-kubernetes system 
resource (eg: etcd).  An example would be testing whether a piece of code can 
correctly store data to or retrieve data from etcd. Integration tests are stored 
in [kubernetes/test/integration/]. Running these can require the developer set up 
additional functionality on their development system.
* End-to-end ("e2e"): These are broad tests of overall system behavior and 
coherence.  These are more complicated as they require a functional kubernetes 
cluster built from the sources to be tested. A 
separate [document detailing e2e testing] and test cases themselves can be found 
in [kubernetes/test/e2e/].
* Conformance: These are a set of testcases, currently a subset of the i
ntegration/e2e tests, that the Architecture SIG has approved to define the core 
set of interoperable features that all Kubernetes deployments must support. For 
more information on Conformance tests please see the [Conformance Testing] 
Document.

Continuous integration will run these tests either as pre-submits on PRs, 
post-submits against /release branches, or both.  
The results appear on [testgrid].

SIG Testing is responsible for that official infrastructure and CI.
The associated automation is tracked in the [test-infra repo].
If you're looking to run e2e tests on your own infrastructure, [kubetest]
is the mechanism.

#### Security

* [Security Release Page] - Outlines the procedures for the handling of security 
  issues.
* [Security and Disclosure Information] - Check this page if you wish to report 
  a security vulnerability.

#### Documentation

* [Contributing to Documentation]

#### Issues Management or Triage

Have you ever noticed the total number of [open issues]?
Helping to manage or triage these open issues can be a great contribution and a 
great opportunity to learn about the various areas of the project. Triaging is 
the word we use to describe the process of adding multiple types of descriptive 
labels to GitHub issues, in order to speed up routing issues to the right folks.
Refer to the [Issue Triage Guidelines] for 
more information.

[Contributor Experience SIG]: https://github.com/kubernetes/community/tree/master/sig-contributor-experience
[community guiding principles]: /contributors/guide/expectations.md#code-review
[contributor cheatsheet]: ./contributor-cheatsheet/
[SIG]: #learn-about-sigs
[standard lines of communication]: /communication/README.md
[troubleshooting guide]: https://kubernetes.io/docs/tasks/debug-application-cluster/troubleshooting/
[the GitHub workflow guide]: ./github-workflow.md
[Kubernetes-specific GitHub Workflow]: ./pull-requests.md#the-testing-and-merge-workflow
[github pull request]: https://help.github.com/articles/about-pull-requests/
[the bot]: https://github.com/k8s-ci-robot
[command reference documentation]: https://go.k8s.io/bot-commands
[CLA page]: /CLA.md)
[Code Review]: #code-review
[Learn About SIGs]: #learn-about-sigs
[test flakes]: http://velodrome.k8s.io/dashboard/db/bigquery-metrics?orgId=1
[scalability best practices]: scalability-good-practices.md
[keywords]: https://help.github.com/en/articles/closing-issues-using-keywords
[on Code Review]: /contributors/guide/expectations.md#code-review
[coding conventions]: coding-conventions.md
[good commit messages]: https://chris.beams.io/posts/git-commit/
[Code of Conduct]: /code-of-conduct.md
[community expectations]: ./expectations.md#expectations-of-reviewers-review-latency
[The Gentle Art of Patch Review]: http://sage.thesharps.us/2014/09/01/the-gentle-art-of-patch-review/
[#pr-reviews]: https://kubernetes.slack.com/messages/pr-reviews
[Closing Issues Via Pull Request]: https://github.com/blog/1506-closing-issues-via-pull-requests 
[Squashing Git Commits]: https://davidwalsh.name/squash-commits-git 
[sig-testing]: /sig-testing
[Testing Guide]: /contributors/devel/sig-testing/testing.md
[testing]: https://golang.org/pkg/testing/
[kubernetes/cmd/kubeadm/app/util/version.go]: https://git.k8s.io/kubernetes/cmd/kubeadm/app/util/version.go
[[kubernetes/cmd/kubeadm/app/util/version_test.go]: https://git.k8s.io/kubernetes/cmd/kubeadm/app/util/version_test.go
[kubernetes/test/integration/]: https://git.k8s.io/kubernetes/test/integration
[document detailing e2e testing]: /contributors/devel/sig-testing/e2e-tests.md
[kubernetes/test/e2e/]: https://git.k8s.io/kubernetes/test/e2e
[Conformance Testing]: /contributors/devel/sig-architecture/conformance-tests.md
[testgrid]: https://testgrid.k8s.io
[test-infra repo]: https://git.k8s.io/test-infra
[kubetest]: https://git.k8s.io/test-infra/kubetest
[Security Release Page]: https://git.k8s.io/security/security-release-process.md
[Security and Disclosure Information]: https://kubernetes.io/docs/reference/issues-security/security/
[Contributing to Documentation]: https://kubernetes.io/editdocs/
[open issues]: https://issues.k8s.io
[Issue Triage Guidelines]: /contributors/guide/issue-triage.md