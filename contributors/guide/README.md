**OWNER:**

SIG-ContribX


# WARNING
Hello! This is the starting point for our brand new contributor guide, currently underway as per [-issue#6102](https://github.com/kubernetes/website/issues/6102) and in need of help. Please be patient, or fix a section below that needs improvement, and submit a pull request!

Many of the links below should lead to relevant documents scattered across the community repository. Often, the linked instructions need to be updated or cleaned up. 

* If you do so, please port the relevant file from its previous location to come live in the community/contributors/guide folder, and delete its previous location.

Please find _Improvements needed_ sections below and help us out.

For example:
_Improvements needed_
* kubernetes/kubernetes/contributing.md -> point to this guide

* kubernetes/community/CONTRIBUTING.md -> Needs a rewrite 

* kubernetes/community/README.md -> Needs a rewrite 

* Individual SIG contributing documents -> add a link to this guide

* A new developer guide will be referenced of the contributor guide as an item

    * RyanJ from Red Hat is working on this

* Generate ToC for the section below 

[[TOC]]

[[TOC]]

# Welcome

Welcome to Kubernetes! This document is the single source of truth for how to contribute to the code base. Please leave comments / suggestions if you find something is missing or incorrect.

# Before you get started

## Sign the CLA

Link to [https://github.com/kubernetes/community/blob/master/CLA.md](https://github.com/kubernetes/community/blob/master/CLA.md)

## Setting up your development environment

If you haven’t set up your environment, please find instructions here:
[https://github.com/kubernetes/community/blob/master/contributors/devel/welcome-to-kubernetes-new-developer-guide.md#downloading-building-and-testing-kubernetes](https://github.com/kubernetes/community/blob/master/contributors/devel/welcome-to-kubernetes-new-developer-guide.md#downloading-building-and-testing-kubernetes)

## Community Expectations 

Kubernetes is a community project. Consequently, it is wholly dependent on its community to provide a productive, friendly and collaborative environment.

The first and foremost goal of the Kubernetes community to develop orchestration technology that radically simplifies the process of creating reliable distributed systems. However a second, equally important goal is the creation of a community that fosters easy, agile development of such orchestration systems.

We therefore describe the expectations for members of the Kubernetes community. This document is intended to be a living one that evolves as the community evolves via the same PR and code review process that shapes the rest of the project. It currently covers the expectations of conduct that govern all members of the community as well as the expectations around code review that govern all active contributors to Kubernetes.

### Code of Conduct

Please make sure to read and observe our Code of Conduct:

[https://github.com/cncf/foundation/blob/master/code-of-conduct.md](https://github.com/cncf/foundation/blob/master/code-of-conduct.md)

### Code Review

[https://github.com/kubernetes/community/blob/master/contributors/devel/community-expectations.md#code-review](https://github.com/kubernetes/community/blob/master/contributors/devel/community-expectations.md#code-review)

TODO: edit above link to only provide code review section.

### Thanks

Many thanks in advance to everyone who contributes their time and effort to making Kubernetes both a successful system as well as a successful community. The strength of our software shines in the strengths of each individual community member. Thanks!

# Your First Contribution

(modified from [https://github.com/kubernetes/community/blob/master/contributors/devel/welcome-to-kubernetes-new-developer-guide.md#downloading-building-and-testing-kubernetes](https://github.com/kubernetes/community/blob/master/contributors/devel/welcome-to-kubernetes-new-developer-guide.md#downloading-building-and-testing-kubernetes))

Have you ever wanted to contribute to the coolest cloud technology? We will help you understand the organization of the Kubernetes project and direct you to the best places to get started. You'll be able to pick up issues, write code to fix them, and get your work reviewed and merged.

If you have questions about the development process, feel free to jump into our [Slack Channel](http://slack.k8s.io/) or join our [mailing list](https://groups.google.com/forum/#!forum/kubernetes-dev).

## Find something to work on

TODO: clean up this paragraph’s flow
Help is always welcome!

Documentation (like the text you are reading now) can always use improvement!

To dig deeper, read a design doc, e.g. [architecture](https://github.com/kubernetes/community/blob/master/contributors/design-proposals/architecture/architecture.md).

There's always code that can be clarified and variables or functions that can be renamed or commented.

There's always a need for more test coverage.

### Find a good first topic

There's a [semi-curated list of issues](https://github.com/kubernetes/kubernetes/issues?q=is%3Aopen+is%3Aissue+label%3Ahelp-wanted) that should not need deep knowledge of the system.
TODO: Clarify help-wanted issues also exist on projects other than k/kubernetes, e.g. k/kubectl.

### Find a SIG that is related to your contribution

After finding something to contribute, you should find a SIG.

TODO: 
- this needs work, how do I find the correct SIG? Example: I have a bugfix for kubectl, I’ve found the repo, 

- [This paragraph](https://github.com/kubernetes/community/blob/master/contributors/devel/issues.md#find-the-right-sigs) helps to find the SIG for an existing issue; what about an "unprompted" PR?
- Make people aware of “extra” contributing guidelines on a SIG.

- File issues with all the SIGs to not have duplicate things in their CONTRIBUTING.md

- Do not overload people with the entirety of the SIG structure; it is very intimidating

Possible text to elaborate on:

([Pick a SIG](https://github.com/kubernetes/community/blob/master/sig-list.md), peruse its associated [cmd](https://github.com/kubernetes/kubernetes/tree/master/cmd) directory, find a main() and read code until you find something you want to fix.)

### File an issue

TODO: clarify there are many k/subrepos where you can file.

Not ready to contribute code, but see something that needs work? While we encourage everyone to contribute code, we also appreciate it when someone finds a problem. For example, here is where you file an issue to kubernetes/kubernetes: [https://github.com/kubernetes/kubernetes/issues/new](https://github.com/kubernetes/kubernetes/issues/new). Please make sure to adhere to the prompted submission guidelines.

# Contributing

(From:

[https://github.com/kubernetes/community/blob/38fef055486c29e0b4d2639560c628f04504de21/contributors/devel/collab.md](https://github.com/kubernetes/community/blob/38fef055486c29e0b4d2639560c628f04504de21/contributors/devel/collab.md))

Kubernetes is open source, but many of the people working on it do so as their day job. In order to avoid forcing people to be "at work" effectively 24/7, we want to establish some semi-formal protocols around development. Hopefully these rules make things go more smoothly. If you find that this is not the case, please complain loudly.

### Patches welcome

First and foremost: as a potential contributor, your changes and ideas are welcome at any hour of the day or night, weekdays, weekends, and holidays. Please do not ever hesitate to ask a question or send a PR. (

### Communication

TODO:

* link to how to contact the correct people

* Communicating link: [https://github.com/kubernetes/community/blob/master/communication.md](https://github.com/kubernetes/community/blob/master/communication.md)

* or perhaps link to [https://kubernetes.io/community/](https://kubernetes.io/community/)

* add content of [Office Hours](https://github.com/kubernetes/community/blob/master/community/office-hours.md) to either of the above contents.

## GitHub workflow

(from [https://github.com/kubernetes/community/blob/master/contributors/devel/development.md](https://github.com/kubernetes/community/blob/master/contributors/devel/development.md))

## Collaborative Development

[https://github.com/kubernetes/community/blob/master/contributors/devel/collab.md](https://github.com/kubernetes/community/blob/master/contributors/devel/collab.md)

## Open a PR

PR workflow is described here:

[https://github.com/kubernetes/community/blob/master/contributors/devel/pull-requests.md#the-testing-and-merge-workflow](https://github.com/kubernetes/community/blob/master/contributors/devel/pull-requests.md#the-testing-and-merge-workflow)

## Code Review

Our general code review ideals are outlined in our [Collaborative development outline.](https://github.com/kubernetes/community/blob/master/contributors/devel/collab.md)

### Code Review Process

TODO:
Include some of the explanations from here: [https://github.com/kubernetes/community/blob/master/contributors/devel/owners.md#code-review-process](https://github.com/kubernetes/community/blob/master/contributors/devel/owners.md#code-review-process)

TODO:
Explain OWNERS files and process (also found in above link)

## Testing

TODO:
- link to testing process

- walkthrough of where to find what in the tests 

## Security

TODO: elaborate

## Documentation

TODO: elaborate

# Community

To find out more about our community structure, different levels of membership and code contributors, please explore here:  [https://github.com/kubernetes/community/blob/master/community-membership.md](https://github.com/kubernetes/community/blob/master/community-membership.md)
We depend on new people becoming members and regular code contributors, so please come join us.

## Events

Kubernetes is the main focus of CloudNativeCon/KubeCon, held every spring in Europe and winter in North America. Information about these and other community events is available on the CNCF [events](https://www.cncf.io/events/) pages.

### Meetups

### Kubecon

## Mentorship

[Link and mini description for Kubernetes Pilots should go here].

