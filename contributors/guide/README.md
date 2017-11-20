**OWNER:**

sid-contributor-experience

## Disclaimer
Hello! This is the starting point for our brand new contributor guide, currently underway as per [-issue#6102](https://github.com/kubernetes/website/issues/6102) and in need of help. Please be patient, or fix a section below that needs improvement, and submit a pull request!

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

* Generate ToC for the section below 

# Welcome

Welcome to Kubernetes! This document is the single source of truth for how to contribute to the code base. Please leave comments / suggestions if you find something is missing or incorrect.

# Before you get started

## Sign the CLA

Before you can contribute, you will need to sign the [Contributor License Agreement](https://git.k8s.io/community/CLA.md).

## Setting up your development environment

If you haven’t set up your environment, please find resources [here](https://github.com/kubernetes/community/tree/master/contributors/devel). These resources are not well organized currently; please have patience as we are working on it.

_Improvements needed_
* A new developer guide will be created and linked to in this section.

    * RyanJ from Red Hat is working on this

## Community Expectations 

Kubernetes is a community project. Consequently, it is wholly dependent on its community to provide a productive, friendly and collaborative environment.

The first and foremost goal of the Kubernetes community to develop orchestration technology that radically simplifies the process of creating reliable distributed systems. However a second, equally important goal is the creation of a community that fosters easy, agile development of such orchestration systems.

We therefore describe the expectations for members of the Kubernetes community. This document is intended to be a living one that evolves as the community evolves via the same PR and code review process that shapes the rest of the project. It currently covers the expectations of conduct that govern all members of the community as well as the expectations around code review that govern all active contributors to Kubernetes.

### Code of Conduct

Please make sure to read and observe our [Code of Conduct](https://github.com/cncf/foundation/blob/master/code-of-conduct.md)

### Code Review

For a brief description of the importance of code review, please read [On Code Review](https://github.com/kubernetes/community/blob/master/contributors/devel/community-expectations.md#code-review).

_Improvements needed_
* edit above link to only provide code review section.
* decide whether this is something that belongs in this section, or should be combined with Contributing instructions [below](/contributor-guide/contributors/guide#contributing).

### Thanks

Many thanks in advance to everyone who contributes their time and effort to making Kubernetes both a successful system as well as a successful community. The strength of our software shines in the strengths of each individual community member. Thanks!

# Your First Contribution

(from [here](https://github.com/kubernetes/community/blob/master/contributors/devel/welcome-to-kubernetes-new-developer-guide.md#introduction))

Have you ever wanted to contribute to the coolest cloud technology? We will help you understand the organization of the Kubernetes project and direct you to the best places to get started. You'll be able to pick up issues, write code to fix them, and get your work reviewed and merged.

If you have questions about the development process, feel free to jump into our [Slack Channel](http://slack.k8s.io/) or join our [mailing list](https://groups.google.com/forum/#!forum/kubernetes-dev).

_Improvements needed_
* The doc linked [above](https://github.com/kubernetes/community/blob/master/contributors/devel/welcome-to-kubernetes-new-developer-guide.md#introduction) is being reinvented in this README. All relevant information should be ported to a logical place in this document/folder, and the original document deleted.

## Find something to work on

Help is always welcome! For example, documentation (like the text you are reading now) can always use improvement. There's always code that can be clarified and variables or functions that can be renamed or commented. There's always a need for more test coverage.
You get the idea - if you ever see something you think should be fixed, you should own it. Here is how you get started.

### Find a good first topic

Each repository in the Kubernetes organization has beginner-friendly issues that provide a good first issue. For example, [kubernetes/kubernetes](https://github.com/kubernetes/kubernetes) has [help-wanted issues](https://github.com/kubernetes/kubernetes/issues?q=is%3Aopen+is%3Aissue+label%3Ahelp-wanted) that should not need deep knowledge of the system.
Another good strategy is to find a documentation improvement, such as a missing/broken link, which will give you exposure to the code submission/review process without the added complication of technical depth.

_Improvements needed_
* Clarify help-wanted issues also exist on projects other than k/kubernetes, e.g. k/kubectl.

### Find a SIG that is related to your contribution

You may have noticed that some repositories in the Kubernetes Organizaion are owned by Special Interest Groups, or SIGs. We organize the Kubernetes community into SIGs in order to improve our workflow and more easily manage what is a very large community project.
SIGs also have their own CONTRIBUTING.md files, which may contain extra information or guidelines in addition to these general ones.

After finding something to contribute, you should find the appropriate SIG, which you will need in order to have your pull request approved for testing and merging.

Here is how (_information needed_)

_Improvements needed_ 
* Clarify how someone would find the correct SIG, and what to do with that information. Link to those SIGs.

* [This paragraph](https://github.com/kubernetes/community/blob/master/contributors/devel/issues.md#find-the-right-sigs) helps to find the SIG for an existing issue; what about a PR not associated with an open issue?

* Make people aware of “extra” contributing guidelines on a SIG.

* At the same time, file issues with all the SIGs to not have duplicate things in their CONTRIBUTING.md. Keep it light, keep it clean, have only one source of truth. 

* Do not overload people with the entirety of the SIG structure; it is very intimidating to a newcomer

* Possible text to elaborate on (the sig list should remain in its current place however):

    * ([Pick a SIG](https://git.k8s.io/community/sig-list.md), peruse its associated [cmd](https://git.k8s.io/kubernetes/cmd) directory, find a main() and read code until you find something you want to fix.)

### File an Issue

Not ready to contribute code, but see something that needs work? While we encourage everyone to contribute code, we also appreciate it when someone finds a problem.
Again, there are multiple repositories within the community, and each will have its own Issues.
For example, here is where you file an Issue to [kubernetes/kubernetes](https://git.k8s.io/kubernetes/issues/new). Please make sure to adhere to the prompted submission guidelines.

_Improvements needed_ 
* clarify there are many k/subrepos where you can file issues. Refer to "how to find and appropriate SIG" document to find out which.

# Contributing

(From:[here](https://git.k8s.io/community/contributors/devel/collab.md))

Kubernetes is open source, but many of the people working on it do so as their day job. In order to avoid forcing people to be "at work" effectively 24/7, we want to establish some semi-formal protocols around development. Hopefully these rules make things go more smoothly. If you find that this is not the case, please complain loudly.

As a potential contributor, your changes and ideas are welcome at any hour of the day or night, weekdays, weekends, and holidays. Please do not ever hesitate to ask a question or send a PR.

Our community guiding principles on how to create great code as a big group are found [here](https://git.k8s.io/community/contributors/devel/collab.md#code-reviews).

If you haven't done so already make sure you have found the correct SIG for your contribution. This will ensure faster responses and a streamlined code review. (see above)

_Improvements needed_ 

* the linked text is awfully similar to the section on [Code Review](https://git.k8s.io/community/contributors/devel/community-expectations.md#code-review), linked above. Consolidate, and decide on _one place_ in this document where it fits best. It might even be important enough to include at top level in this document(no links).

### Communication

It is best to contact your SIG for issues related to the SIG's topic. Your SIG will be able to help you much more quickly than a general question would. Each SIG has a kubernetes slack channel that you can join as well.

For questions and troubleshooting, please feel free to use any of the methods of communication listed [here](https://git.k8s.io/community/communication.md). The [kubernetes website](https://kubernetes.io/community/) also lists this information.

## GitHub workflow

To check out code to work on, please refer to [this guide](https://git.k8s.io/community/contributors/devel/development.md#workflow).

_Improvements needed_ 

* move github workflow into its own file in this folder.

## Open a PR

PR workflow is described [here](https://git.k8s.io/community/contributors/devel/pull-requests.md#the-testing-and-merge-workflow).

## Code Review

_Improvements needed_
* Clarify and streamline some of the explanations from [here](https://git.k8s.io/community/contributors/devel/owners.md#code-review-process).

* Explain OWNERS files and process (also found in above link)

## Testing

_Improvements needed_
* link to testing process 
* walkthrough of where to find what in the tests (how to use, how to debug)

## Security

_Improvements needed_
* Please help write this section.

## Documentation

_Improvements needed_
* Please help write this section.

# Community

If you haven't noticed by now, we have a large, lively, and friendly open-source community. We depend on new people becoming members and regular code contributors, so we would like you to come join us. To find out more about our community structure, different levels of membership and code contributors, please [explore here](https://git.k8s.io/community/community-membership.md).

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

