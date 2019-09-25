---
title: "Kubernetes Contributor Guide"
weight: 1
slug: "guide" 
---


<!-- 
Contributing to this document? 
Please use semantic line feeds for readability: http://rhodesmill.org/brandon/2012/one-sentence-per-line/ 
-->

This document is the single source of truth for how to contribute to the code base.
Feel free to browse the [open issues](https://github.com/kubernetes/community/issues?q=is%3Aissue+is%3Aopen+label%3Aarea%2Fcontributor-guide) and file new ones, all feedback welcome!

# Welcome

Welcome to Kubernetes! 

-   [Before you get started](#before-you-get-started)
    -   [Sign the CLA](#sign-the-cla)
    -   [Code of Conduct](#code-of-conduct)
    -   [Setting up your development
        environment](#setting-up-your-development-environment)
    -   [Community Expectations and Roles](#community-expectations-and-roles)
        -   [Thanks](#thanks)
-   [Your First Contribution](#your-first-contribution)
    -   [Find something to work on](#find-something-to-work-on)
        -   [Find a good first topic](#find-a-good-first-topic)
        -   [Learn about SIGs](#learn-about-sigs)
        -   [SIG-specific contributing guidelines](#sig-specific-contributing-guidelines)
        -   [File an Issue](#file-an-issue)
-   [Contributing](#contributing)
    -   [Communication](#communication)
    -   [GitHub workflow](#github-workflow)
    -   [Open a Pull Request](#open-a-pull-request)
    -   [Code Review](#code-review)
    -   [Testing](#testing)
    -   [Security](#security)    
    -   [Documentation](#documentation)
    -   [Issues Management or Triage](#issues-management-or-triage)
-   [Kubernetes Contributor Playground](#kubernetes-contributor-playground)
    -   [Youtube playlist](#youtube-playlist)
-   [Community](#community)
    -   [Communication](#communication-1)
    -   [Events](#events)
        -   [Meetups](#meetups)
    -   [Mentorship](#mentorship)

# Before you get started

## Sign the CLA

Before you can contribute, you will need to sign the [Contributor License Agreement](/CLA.md).

## Code of Conduct

Please make sure to read and observe our [Code of Conduct](/code-of-conduct.md).

## Setting up your development environment

If you haven’t set up your environment, check the [developer resources](/contributors/devel/README.md#setting-up-your-dev-environment-coding-and-debugging).

## Community Expectations and Roles

Kubernetes is a community project.
Consequently, it is wholly dependent on its community to provide a productive, friendly and collaborative environment.

- Read and review the [Community Expectations](expectations.md) for an understanding of code and review expectations.
- See [Community Membership](/community-membership.md) for a list the various responsibilities of contributor roles. You are encouraged to move up this contributor ladder as you gain experience.  

# Your First Contribution

Have you ever wanted to contribute to the coolest cloud technology?
We will help you understand the organization of the Kubernetes project and direct you to the best places to get started.
You'll be able to pick up issues, write code to fix them, and get your work reviewed and merged.

Please be aware that due to the large number of issues our triage team deals with, we cannot offer technical support in GitHub issues.
If you have questions about the development process, feel free to jump into our [Slack Channel](http://slack.k8s.io/) or join our [mailing list](https://groups.google.com/forum/#!forum/kubernetes-dev).
You can also ask questions on [ServerFault](https://serverfault.com/questions/tagged/kubernetes) or [Stack Overflow](https://stackoverflow.com/questions/tagged/kubernetes).
The Kubernetes team scans Stack Overflow on a regular basis and will try to ensure your questions don't go unanswered.

## Find something to work on

Help is always welcome! For example, documentation (like the text you are reading now) can always use improvement. 
There's always code that can be clarified and variables or functions that can be renamed or commented.
There's always a need for more test coverage.
You get the idea - if you ever see something you think should be fixed, you should own it.
Here is how you get started.
If you have no idea what to start on, you can browse the [Contributor Role Board](https://discuss.kubernetes.io/c/contributors/role-board) to see who is looking for help.
Those interested in contributing without writing code may also find ideas in the [Non-Code Contributions Guide](non-code-contributions.md).

### Find a good first topic

There are [multiple repositories](https://github.com/kubernetes/) within the Kubernetes organization.
Each repository has beginner-friendly issues that provide a good first issue.
For example, [kubernetes/kubernetes](https://git.k8s.io/kubernetes) has [help wanted](https://go.k8s.io/help-wanted) and [good first issue](https://go.k8s.io/good-first-issue) labels for issues that should not need deep knowledge of the system.
The `good first issue` label indicates that members have committed to providing [extra assistance](/contributors/guide/help-wanted.md) for new contributors.
<!-- TODO: review removing this note after 3 months or after the 1.12 release -->
Please note that while several of the repositories in the Kubernetes community have `good first issue` labels already, they are still being applied throughout the community.

Another good strategy is to find a documentation improvement, such as a missing/broken link, which will give you exposure to the code submission/review process without the added complication of technical depth. Please see [Contributing](#contributing) below for the workflow.

#### Issue Assignment in Github

When you are willing to take on an issue, you can assign it to yourself. Just reply with `/assign` or `/assign @yourself` on an issue, 
then the robot will assign the issue to you and your name will present at `Assignees` list.

### Learn about SIGs

#### SIG structure

You may have noticed that some repositories in the Kubernetes Organization are owned by Special Interest Groups, or SIGs.
We organize the community into SIGs in order to improve our workflow and more easily manage what is a very large community project.
The developers within each SIG have autonomy and ownership over that SIG's part of Kubernetes.

A SIG is an open, community effort.
Anybody is welcome to jump into a SIG and begin fixing issues, critiquing design proposals and reviewing code.
SIGs have regular [video meetings](https://kubernetes.io/community/) which everyone is welcome to.
Each SIG has a slack channel that you can join as well.

There is an entire SIG ([sig-contributor-experience](/sig-contributor-experience/README.md)) devoted to improving your experience as a contributor.
Contributing to Kubernetes should be easy.
If you find a rough edge, let us know! Better yet, help us fix it by joining the SIG; just
show up to one of the [bi-weekly meetings](https://docs.google.com/document/d/1qf-02B7EOrItQgwXFxgqZ5qjW0mtfu5qkYIF1Hl4ZLI/edit).

#### Find a SIG that is related to your contribution

Finding the appropriate SIG for your contribution and adding a SIG label will help you ask questions in the correct place and give your contribution higher visibility and a faster community response.

For Pull Requests, the automatically assigned reviewer will add a SIG label if you haven't done so. See [Open A Pull Request](#open-a-pull-request) below.

For Issues, we are still working on a more automated workflow.
Since SIGs do not directly map onto Kubernetes subrepositories, it may be difficult to find which SIG your contribution belongs in.
Here is the [list of SIGs](/sig-list.md) so that you can determine which is most likely related to your contribution.

*Example:* if you are filing a CNI issue (that's [Container Networking Interface](https://github.com/containernetworking/cni)), you should choose the [Network SIG](http://git.k8s.io/community/sig-network). Add the SIG label in a comment like so:
```
/sig network
```

Follow the link in the SIG name column to reach each SIGs README. 
Most SIGs will have a set of GitHub Teams with tags that can be mentioned in a comment on issues and pull requests for higher visibility. 
If you are not sure about the correct SIG for an issue, you can try SIG-contributor-experience [here](/sig-contributor-experience#github-teams), or [ask in Slack](http://slack.k8s.io/).

### SIG-specific contributing guidelines
Some SIGs have their own `CONTRIBUTING.md` files, which may contain extra information or guidelines in addition to these general ones.
These are located in the SIG-specific community directories:
- [`/sig-apps/CONTRIBUTING.md`](/sig-apps/CONTRIBUTING.md)
- [`/sig-aws/CONTRIBUTING.md`](/sig-aws/CONTRIBUTING.md)
- [`/sig-cli/CONTRIBUTING.md`](/sig-cli/CONTRIBUTING.md)
- [`/sig-multicluster/CONTRIBUTING.md`](/sig-multicluster/CONTRIBUTING.md)
- [`/sig-storage/CONTRIBUTING.md`](/sig-storage/CONTRIBUTING.md)
- [`/sig-windows/CONTRIBUTING.md`](/sig-windows/CONTRIBUTING.md)

### File an Issue

Not ready to contribute code, but see something that needs work?
While the community encourages everyone to contribute code, it is also appreciated when someone reports an issue (aka problem).
Issues should be filed under the appropriate Kubernetes subrepository.
Check the [issue triage guide](./issue-triage.md) for more information.

*Example:* a documentation issue should be opened to [kubernetes/website](https://github.com/kubernetes/website/issues).

Make sure to adhere to the prompted submission guidelines while opening an issue.

# Contributing

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

# Kubernetes Contributor Playground

If you are looking for a safe place, where you can familiarize yourself with (some of) the Kubernetes Project's review and pull request processes, then the [Kubernetes Contributor Playground](https://github.com/kubernetes-sigs/contributor-playground/) is the right place for you.

## Youtube playlist

A [Youtube playlist](https://www.youtube.com/playlist?list=PL69nYSiGNLP3M5X7stuD7N4r3uP2PZQUx) of the New Contributor workshop has been posted, and an outline of the video content can be found [here](/events/2018/05-contributor-summit). 

# Community

If you haven't noticed by now, we have a large, lively, and friendly open-source community.
We depend on new people becoming members and regular code contributors, so we would like you to come join us!
The [Community Membership Document](/community-membership.md) covers membership processes and roles.

## Communication

- [General Information](/communication)

## Events

Kubernetes participates in KubeCon + CloudNativeCon, held three times per year in China, Europe and in North America.
Information about these and other community events is available on the CNCF [events](https://www.cncf.io/events/) pages.

### Meetups

We follow the general [Cloud Native Computing Foundation guidelines](https://github.com/cncf/meetups) for Meetups.
You may also contact Paris Pittman via direct message on Kubernetes Slack (@paris) or by email (parispittman@google.com)

## Mentorship

Please learn about our mentoring initiatives [here](http://git.k8s.io/community/mentoring/README.md).
Feel free to ask us anything during our [Meet Our Contributors](https://github.com/kubernetes/community/blob/master/mentoring/meet-our-contributors.md) to connect with us.
# Advanced Topics

This section includes things that need to be documented, but typical contributors do not need to interact with regularly.

- [OWNERS files](owners.md) - The Kubernetes organizations are managed with OWNERS files, which outline which parts of the code are owned by what groups.
