**OWNER:**

sig-contributor-experience

## Disclaimer
Hello! This is the starting point for our brand new contributor guide, currently underway as per [issue#6102](https://github.com/kubernetes/website/issues/6102) and in need of help. Please be patient, or fix a section below that needs improvement, and submit a pull request!

Many of the links below should lead to relevant documents scattered across the community repository. Often, the linked instructions need to be updated or cleaned up. 

* If you do so, please move the relevant file from its previous location to the community/contributors/guide folder, and delete its previous location.
* Our goal is that all contributor guide specific files live in this folder.

Please find _Improvements needed_ sections below and help us out.

For example:

_Improvements needed_
* kubernetes/kubernetes/contributing.md -> point to this guide

* kubernetes/community/CONTRIBUTING.md -> Needs a rewrite 

* kubernetes/community/README.md -> Needs a rewrite 

* Individual SIG contributing documents -> add a link to this guide

# Welcome

Welcome to Kubernetes! This document is the single source of truth for how to contribute to the code base. Please leave comments / suggestions if you find something is missing or incorrect.

-   [Before you get started](#before-you-get-started)
    -   [Sign the CLA](#sign-the-cla)
    -   [Setting up your development
        environment](#setting-up-your-development-environment)
    -   [Community Expectations](#community-expectations)
        -   [Code of Conduct](#code-of-conduct)
        -   [Thanks](#thanks)
-   [Your First Contribution](#your-first-contribution)
    -   [Find something to work on](#find-something-to-work-on)
        -   [Find a good first topic](#find-a-good-first-topic)
        -   [Learn about SIGs](#learn-about-sigs)
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
-   [Community](#community)
    -   [Events](#events)
        -   [Meetups](#meetups)
        -   [KubeCon](#kubecon)
    -   [Mentorship](#mentorship)

# Before you get started

## Sign the CLA

Before you can contribute, you will need to sign the [Contributor License Agreement](/CLA.md).

## Setting up your development environment

If you haven’t set up your environment, please find resources [here](/contributors/devel). These resources are not well organized currently; please have patience as we are working on it.

_Improvements needed_
* A new developer guide will be created and linked to in this section.

    * RyanJ from Red Hat is working on this

## Community Expectations 

Kubernetes is a community project. Consequently, it is wholly dependent on its community to provide a productive, friendly and collaborative environment.

The first and foremost goal of the Kubernetes community to develop orchestration technology that radically simplifies the process of creating reliable distributed systems. However a second, equally important goal is the creation of a community that fosters easy, agile development of such orchestration systems.

We therefore describe the expectations for members of the Kubernetes community. This document is intended to be a living one that evolves as the community evolves via the same pull request and code review process that shapes the rest of the project. It currently covers the expectations of conduct that govern all members of the community as well as the expectations around code review that govern all active contributors to Kubernetes.

### Code of Conduct

Please make sure to read and observe our [Code of Conduct](https://github.com/cncf/foundation/blob/master/code-of-conduct.md)

### Thanks

Many thanks in advance to everyone who contributes their time and effort to making Kubernetes both a successful system as well as a successful community. The strength of our software shines in the strengths of each individual community member. Thanks!

# Your First Contribution

Have you ever wanted to contribute to the coolest cloud technology? We will help you understand the organization of the Kubernetes project and direct you to the best places to get started. You'll be able to pick up issues, write code to fix them, and get your work reviewed and merged.

Please be aware that due to the large number of issues our triage team deals with, we cannot offer technical support in GitHub issues. If you have questions about the development process, feel free to jump into our [Slack Channel](http://slack.k8s.io/) or join our [mailing list](https://groups.google.com/forum/#!forum/kubernetes-dev). You can also ask questions on [ServerFault](https://serverfault.com/questions/tagged/kubernetes) or [Stack Overflow](https://stackoverflow.com/questions/tagged/kubernetes). The Kubernetes team scans Stack Overflow on a regular basis, and will try to ensure your questions don't go unanswered.

## Find something to work on

Help is always welcome! For example, documentation (like the text you are reading now) can always use improvement. There's always code that can be clarified and variables or functions that can be renamed or commented. There's always a need for more test coverage.
You get the idea - if you ever see something you think should be fixed, you should own it. Here is how you get started.

### Find a good first topic

There are multiple repositories within the Kubernetes community and a full list of repositories can be found [here](https://github.com/kubernetes/).
Each repository in the Kubernetes organization has beginner-friendly issues that provide a good first issue. For example, [kubernetes/kubernetes](https://git.k8s.io/kubernetes) has [help-wanted issues](https://issues.k8s.io/?q=is%3Aopen+is%3Aissue+label%3A%22help%20wanted%22) that should not need deep knowledge of the system.
Another good strategy is to find a documentation improvement, such as a missing/broken link, which will give you exposure to the code submission/review process without the added complication of technical depth. Please see [Contributing](#Contributing) below for the workflow.

### Learn about SIGs

#### Sig structure

You may have noticed that some repositories in the Kubernetes Organization are owned by Special Interest Groups, or SIGs. We organize the Kubernetes community into SIGs in order to improve our workflow and more easily manage what is a very large community project. The developers within each SIG have autonomy and ownership over that SIG's part of Kubernetes.
SIGs also have their own CONTRIBUTING.md files, which may contain extra information or guidelines in addition to these general ones. These are located in the SIG specific community documentation directories, for example: sig-docs' is in the kubernetes/community repo's [/sig-docs/CONTRIBUTING.md](/sig-docs/CONTRIBUTING.md) file and similarly for other SIGs.

Like everything else in Kubernetes, a SIG is an open, community, effort. Anybody is welcome to jump into a SIG and begin fixing issues, critiquing design proposals and reviewing code. SIGs have regular [video meetings](https://kubernetes.io/community/) which everyone is welcome to. Each SIG has a kubernetes slack channel that you can join as well.

There is an entire SIG ([sig-contributor-experience](../../sig-contributor-experience/README.md)) devoted to improving your experience as a contributor. 
Contributing to Kubernetes should be easy. If you find a rough edge, let us know! Better yet, help us fix it by joining the SIG; just
show up to one of the [bi-weekly meetings](https://docs.google.com/document/d/1qf-02B7EOrItQgwXFxgqZ5qjW0mtfu5qkYIF1Hl4ZLI/edit).

#### Find a SIG that is related to your contribution

Finding the appropriate SIG for your contribution will help you ask questions in the correct place and give your contribution higher visibility and a faster community response.

For Pull Requests, the automatically assigned reviewer will add a SIG label if you haven't done so. See [Open A Pull Request](#open-a-pull-request) below. 

For Issues we are still working on a more automated workflow. Since SIGs do not directly map onto Kubernetes subrepositories, it may be difficult to find which SIG your contribution belongs in. Here is the [list of SIGs](/sig-list.md). Determine which is most likely related to your contribution. 

*Example:* if you are filing a cni issue, you should choose SIG-networking. 

Follow the link in the SIG name column to reach each SIGs README. Most SIGs will have a set of GitHub Teams with tags that can be mentioned in a comment on issues and pull requests for higher visibility. If you are not sure about the correct SIG for an issue, you can try SIG-contributor-experience [here](/sig-contributor-experience#github-teams), or [ask in Slack](http://slack.k8s.io/).

_Improvements needed_ 

* Open pull requests with all applicable SIGs to not have duplicate information in their CONTRIBUTING.md and instead link here. Keep it light, keep it clean, have only one source of truth. 

### File an Issue

Not ready to contribute code, but see something that needs work? While the community encourages everyone to contribute code, it is also appreciated when someone reports an issue (aka problem). Issues should be filed under the appropriate Kubernetes subrepository. 

*Example:* a documentation issue should be opened to [kubernetes/website](https://github.com/kubernetes/website/issues). 

Make sure to adhere to the prompted submission guidelines while opening an issue.

# Contributing

(From:[here](/contributors/devel/collab.md))

Kubernetes is open source, but many of the people working on it do so as their day job. In order to avoid forcing people to be "at work" effectively 24/7, we want to establish some semi-formal protocols around development. Hopefully these rules make things go more smoothly. If you find that this is not the case, please complain loudly.

As a potential contributor, your changes and ideas are welcome at any hour of the day or night, weekdays, weekends, and holidays. Please do not ever hesitate to ask a question or send a pull request.

Our community guiding principles on how to create great code as a big group are found [here](/contributors/devel/collab.md). Beginner focused information can be found below in [Open a Pull Request](#open-a-pull-request) and [Code Review](#code-review).

### Communication

It is best to contact your [SIG](#learn-about-sigs) for issues related to the SIG's topic. Your SIG will be able to help you much more quickly than a general question would.

For questions and troubleshooting, please feel free to use any of the methods of communication listed [here](/communication.md). The [kubernetes website](https://kubernetes.io/docs/tasks/debug-application-cluster/troubleshooting/) also lists this information.

## GitHub workflow

To check out code to work on, please refer to [this guide](/contributors/devel/development.md#workflow).

_Improvements needed_ 

* move github workflow into its own file in this folder.

## Open a Pull Request

Pull requests are often called simply "PR".  Kubernetes generally follows the standard [github pull request](https://help.github.com/articles/about-pull-requests/) process, but there is a layer of additional kubernetes specific (and sometimes SIG specific) differences.

The first difference you'll see is that a bot will begin applying structured labels to your PR.

The bot may also make some helpful suggestions for commands to run in your PR to facilitate review.  These `/command` options can be entered in comments to trigger auto-labeling and notifications.  The command reference is [here](https://go.k8s.io/bot-commands).

Common new contributor PR issues are:

* not having correctly signed the CLA ahead of your first PR (see [Sign the CLA](#sign-the-cla) section)
* finding the right SIG or reviewer(s) for the PR (see [Code Review](#code-review) section) and following any SIG specific contributing guidelines
* dealing with test cases which fail on your PR, unrelated to the changes you introduce (see [Test Flakes](http://velodrome.k8s.io/dashboard/db/bigquery-metrics?orgId=1))

The pull request workflow is described in detail [here](/contributors/devel/pull-requests.md#the-testing-and-merge-workflow).

## Code Review

For a brief description of the importance of code review, please read [On Code Review](/contributors/devel/community-expectations.md#code-review).  There are two aspects of code review: giving and receiving.

To make it easier for your PR to receive reviews, consider the reviewers will need you to:

* write [good commit messages](https://chris.beams.io/posts/git-commit/)
* break large changes into a logical series of smaller patches which individually make easily understandable changes, and in aggregate solve a broader issue
* label PRs with appropriate SIGs and reviewers: to do this read the messages the bot sends you to guide you through the PR process

Reviewers, the people giving review, are highly encouraged to revisit the [Code of Conduct](https://github.com/cncf/foundation/blob/master/code-of-conduct.md) and must go above and beyond to promote a collaborative, respectful Kubernetes community.  When reviewing PRs from others [The Gentle Art of Patch Review](http://sage.thesharps.us/2014/09/01/the-gentle-art-of-patch-review/) suggests an iterative series of focuses which is designed to lead new contributors to positive collaboration without inundating them initially with nuances:

* Is the idea behind the contribution sound?
* Is the contribution architected correctly?
* Is the contribution polished?

## Testing

Testing is the responsibility of all contributors and is in part owned by all sigs, but is also coordinated by [sig-testing](/sig-testing).

The main testing overview document is [here](/contributors/devel/testing.md).

There are three types of test in kubernetes.  The location of the test code varies with type, as does the specifics of the environment needed to successfully run the test:

* Unit: These confirm that a particular function behaves as intended.  Golang includes native ability for unit testing via the [testing](https://golang.org/pkg/testing/) package.  Unit test source code can be found adjacent to the corresponding source code within a given package.  For example: functions defined in [kubernetes/cmd/kubeadm/app/util/version.go](https://git.k8s.io/kubernetes/cmd/kubeadm/app/util/version.go) will have unit tests in [kubernetes/cmd/kubeadm/app/util/version_test.go](https://git.k8s.io/kubernetes/cmd/kubeadm/app/util/version_test.go).  These are easily run locally be any developer on any OS.
* Integration: These tests cover interactions of package components or interactions between kubernetes components and some other non-kubernetes system resource (eg: etcd).  An example would be testing whether a piece of code can correctly store data to or retrieve data from etcd.  Integration tests are stored in [kubernetes/test/integration/](https://git.k8s.io/kubernetes/test/integration).  Running these can require the developer set up additional functionality on their development system.
* End-to-end ("e2e"): These are broad tests of overall kubernetes system behavior and coherence.  These are more complicated as they require a functional kubernetes cluster built from the sources to be tested.  A separate document [here](/contributors/devel/e2e-tests.md) details e2e testing and test cases themselves can be found in [kubernetes/test/e2e/](https://git.k8s.io/kubernetes/test/e2e).

Continuous integration will run these tests either as pre-submits on PRs, post-submits against master/release branches, or both.  The results appear on [testgrid](https://testgrid.k8s.io).

sig-testing is responsible for that official infrastructure and CI.  The associated automation is tracked in the [test-infra repo](https://git.k8s.io/test-infra).  If you're looking to run e2e tests on your own infrastructure, [kubetest](https://git.k8s.io/test-infra/kubetest) is the mechanism.

## Security

_Improvements needed_
* Please help write this section.

## Documentation

_Improvements needed_
* Please help write this section.

## Issues Management or Triage

Have you ever noticed the total number of [open issues](https://issues.k8s.io)? This number at any given time is typically high. Helping to manage or triage these open issues can be a great contribution to the Kubernetes project. This is also a great opportunity to learn about the various areas of the project. Refer to the [Kubernetes Issue Triage Guidelines](/contributors/devel/issues.md) for more information.

# Community

If you haven't noticed by now, we have a large, lively, and friendly open-source community. We depend on new people becoming members and regular code contributors, so we would like you to come join us. To find out more about our community structure, different levels of membership and code contributors, please [explore here](/community-membership.md).

_Improvements needed_

* The top level k/community/README.md should be a good starting point for what the community is and does. (see above instructions on rewriting this file)

## Events
Kubernetes is the main focus of CloudNativeCon/KubeCon, held twice per year in EMEA and in North America. Information about these and other community events is available on the CNCF [events](https://www.cncf.io/events/) pages.

### Meetups

_Improvements needed_
* include link to meetups
* information on CNCF support for founding a Meetup

### KubeCon

_Improvements needed_
* write friendly blurb about KubeCon, and include links

## Mentorship

_Improvements needed_
* Link and mini description for Kubernetes Pilots should go here.

