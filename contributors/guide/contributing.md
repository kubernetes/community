---
title: "Contributing to Kubernetes"
weight: 4
slug: "contributing" 
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

Kubernetes is open source, but many of the people working on it do so as their day job.
In order to avoid forcing people to be "at work" effectively 24/7, we want to establish some semi-formal protocols around development.
Hopefully, these rules make things go more smoothly.
If you find that this is not the case, please complain loudly.

As a potential contributor, your changes and ideas are welcome at any hour of the day or night, weekdays, weekends, and holidays.
Please do not ever hesitate to ask a question or send a pull request.

Check out our [community guiding principles](/contributors/guide/collab.md) on how to create great code as a big group.

Beginner focused information can be found below in [Open a Pull Request](#open-a-pull-request) and [Code Review](#code-review).

For quick reference on contributor resources, we have a handy [contributor cheatsheet](./contributor-cheatsheet/).

### Communication

It is best to contact your [SIG](#learn-about-sigs) for issues related to the SIG's topic. Your SIG will be able to help you much more quickly than a general question would.

For general questions and troubleshooting, use the [standard lines of communication](/communication/README.md) and work through the [troubleshooting guide](https://kubernetes.io/docs/tasks/debug-application-cluster/troubleshooting/).

## GitHub workflow

To check out code to work on, please refer to [the GitHub Workflow Guide](./github-workflow.md).

## Open a Pull Request

Pull requests are often called simply "PR".
Kubernetes generally follows the standard [github pull request](https://help.github.com/articles/about-pull-requests/) process, but there is a layer of additional kubernetes specific (and sometimes SIG specific) differences:

- [Kubernetes-specific github workflow](pull-requests.md#the-testing-and-merge-workflow).

The first difference you'll see is that a bot will begin applying structured labels to your PR.

The bot may also make some helpful suggestions for commands to run in your PR to facilitate review.  
These `/command` options can be entered in comments to trigger auto-labeling and notifications.
Refer to its [command reference documentation](https://go.k8s.io/bot-commands).

Common new contributor PR issues are:

* not having correctly signed the CLA ahead of your first PR (see [Sign the CLA](#sign-the-cla) section)
* finding the right SIG or reviewer(s) for the PR (see [Code Review](#code-review) section) and following any SIG or repository specific contributing guidelines (see [Learn about SIGs](#learn-about-sigs) section)
* dealing with test cases which fail on your PR, unrelated to the changes you introduce (see [Test Flakes](http://velodrome.k8s.io/dashboard/db/bigquery-metrics?orgId=1))
* Not following [scalability good practices](scalability-good-practices.md)
* Include mentions (like @person) and [keywords](https://help.github.com/en/articles/closing-issues-using-keywords) which could close the issue (like fixes #xxxx) in commit messages.

## Code Review

For a brief description of the importance of code review, please read [On Code Review](/contributors/guide/expectations.md#code-review).  
There are two aspects of code review: giving and receiving.

To make it easier for your PR to receive reviews, consider the reviewers will need you to:

* follow the project [coding conventions](coding-conventions.md)
* write [good commit messages](https://chris.beams.io/posts/git-commit/)
* break large changes into a logical series of smaller patches which individually make easily understandable changes, and in aggregate solve a broader issue
* label PRs with appropriate SIGs and reviewers: to do this read the messages the bot sends you to guide you through the PR process

Reviewers, the people giving the review, are highly encouraged to revisit the [Code of Conduct](/code-of-conduct.md) as well as [community expectations](./expectations.md#expectations-of-reviewers-review-latency) and must go above and beyond to promote a collaborative, respectful community.
When reviewing PRs from others [The Gentle Art of Patch Review](http://sage.thesharps.us/2014/09/01/the-gentle-art-of-patch-review/) suggests an iterative series of focuses which is designed to lead new contributors to positive collaboration without inundating them initially with nuances:

* Is the idea behind the contribution sound?
* Is the contribution architected correctly?
* Is the contribution polished?

Note: if your pull request isn't getting enough attention, you can use the [#pr-reviews](https://kubernetes.slack.com/messages/pr-reviews) channel on Slack to get help finding reviewers.

## Best practices

- Write clear and meaningful git commit messages.
- If the PR will *completely* fix a specific issue, include `fixes #123` in the PR body (where 123 is the specific issue number the PR will fix. This will automatically close the issue when the PR is merged.
- Make sure you don't include `@mentions` or `fixes` keywords in your git commit messages. These should be included in the PR body instead.
- When you make a PR for small change (such as fixing a typo, style change, or grammar fix), please squash your commits so that we can maintain a cleaner git history.
- Make sure you include a clear and detailed PR description explaining the reasons for the changes, and ensuring there is sufficient information for the reviewer to understand your PR.
- Additional Readings : 
    - [chris.beams.io/posts/git-commit/](https://chris.beams.io/posts/git-commit/)
    - [github.com/blog/1506-closing-issues-via-pull-requests ](https://github.com/blog/1506-closing-issues-via-pull-requests )
    - [davidwalsh.name/squash-commits-git ](https://davidwalsh.name/squash-commits-git )

## Testing

Testing is the responsibility of all contributors and is in part owned by all SIGs, but is also coordinated by [sig-testing](/sig-testing). 
Refer to the [Testing Guide](/contributors/devel/sig-testing/testing.md) for more information.

There are multiple types of tests.
The location of the test code varies with type, as do the specifics of the environment needed to successfully run the test:

* Unit: These confirm that a particular function behaves as intended.  Golang includes a native ability for unit testing via the [testing](https://golang.org/pkg/testing/) package.  Unit test source code can be found adjacent to the corresponding source code within a given package.  For example: functions defined in [kubernetes/cmd/kubeadm/app/util/version.go](https://git.k8s.io/kubernetes/cmd/kubeadm/app/util/version.go) will have unit tests in [kubernetes/cmd/kubeadm/app/util/version_test.go](https://git.k8s.io/kubernetes/cmd/kubeadm/app/util/version_test.go).  These are easily run locally by any developer on any OS.
* Integration: These tests cover interactions of package components or interactions between kubernetes components and some other non-kubernetes system resource (eg: etcd).  An example would be testing whether a piece of code can correctly store data to or retrieve data from etcd.  Integration tests are stored in [kubernetes/test/integration/](https://git.k8s.io/kubernetes/test/integration).  Running these can require the developer set up additional functionality on their development system.
* End-to-end ("e2e"): These are broad tests of overall system behavior and coherence.  These are more complicated as they require a functional kubernetes cluster built from the sources to be tested. A separate [document detailing e2e testing](/contributors/devel/sig-testing/e2e-tests.md) and test cases themselves can be found in [kubernetes/test/e2e/](https://git.k8s.io/kubernetes/test/e2e).
* Conformance: These are a set of testcases, currently a subset of the integration/e2e tests, that the Architecture SIG has approved to define the core set of interoperable features that all Kubernetes deployments must support. For more information on Conformance tests please see the [Conformance Testing](/contributors/devel/sig-architecture/conformance-tests.md) Document.

Continuous integration will run these tests either as pre-submits on PRs, post-submits against master/release branches, or both.  
The results appear on [testgrid](https://testgrid.k8s.io).

sig-testing is responsible for that official infrastructure and CI.
The associated automation is tracked in the [test-infra repo](https://git.k8s.io/test-infra).
If you're looking to run e2e tests on your own infrastructure, [kubetest](https://git.k8s.io/test-infra/kubetest) is the mechanism.

## Security

- [Security Release Page](https://git.k8s.io/security/security-release-process.md) - outlines the procedures for the handling of security issues.
- [Security and Disclosure Information](https://kubernetes.io/docs/reference/issues-security/security/) - check this page if you wish to report a security vulnerability.

## Documentation

- [Contributing to Documentation](https://kubernetes.io/editdocs/)

## Issues Management or Triage

Have you ever noticed the total number of [open issues](https://issues.k8s.io)?
Helping to manage or triage these open issues can be a great contribution and a great opportunity to learn about the various areas of the project. Triaging is the word we use to describe the process of adding multiple types of descriptive labels to GitHub issues, in order to speed up routing issues to the right folks.
Refer to the [Issue Triage Guidelines](/contributors/guide/issue-triage.md) for more information.
