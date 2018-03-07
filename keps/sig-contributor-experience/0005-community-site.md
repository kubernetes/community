---
kep-number: 5
title: Contributor Site
authors:
  - "@jbeda"
owning-sig: sig-contributor-experience
participating-sigs:
  - sig-architecture
  - sig-docs
reviewers:
  - "@castrojo"
approvers:
  - "@parispittman"
editor: TBD
creation-date: "2018-02-19"
last-updated: "2018-03-07"
status: provisional
---

# Contributor Site

## Table of Contents

* [Table of Contents](#table-of-contents)
* [Summary](#summary)
* [Motivation](#motivation)
    * [Goals](#goals)
    * [Non-Goals](#non-goals)
* [Proposal](#proposal)
    * [User Stories [optional]](#user-stories-optional)
      * [Story 1](#story-1)
      * [Story 2](#story-2)
    * [Implementation Details/Notes/Constraints [optional]](#implementation-detailsnotesconstraints-optional)
    * [Risks and Mitigations](#risks-and-mitigations)
* [Graduation Criteria](#graduation-criteria)
* [Implementation History](#implementation-history)
* [Drawbacks [optional]](#drawbacks-optional)
* [Alternatives [optional]](#alternatives-optional)

## Summary

We need a way to organize and publish information targeted at contributors.
In order to continue to scale the Kubernetes community we need a convenient, scalable and findable way to publish information.

## Motivation

While the current kubernetes.io site is great for end users, it isn't often used by or aimed at project contributors.
Instead, most contributors look at documentation in markdown files that are spread throughout a wide set of repos and orgs.
It is difficult for users to find this documentation.

Furthermore, this documentation is often duplicated and out of date.
The fact that it isn't collected in one place and presented as a whole leads to fragmentation.
Often times documentation will be duplicated because the authors themselves can't find the relevant docs.

This site will also serve as a starting point for those that are looking to contribute.
This site (and the contributor guide) can provide a soft introduction to the main processes and groups.


Finally, some simple domain specific indexing could go a long way to make it easier to discover and cross link information.
Specifically, building a site that can take advantage of the KEP metadata will both make KEPs more discoverable and encourage those in the community to publish information in a way that *can* be discovered.

### Goals

* Provide a new site (`community.kubernetes.io`? `contrib.kubernetes.io`?).
* Publish information that is currently in the [community repo](https://github.com/kubernetes/community).
* Build some simple tools to enhance discoverability within the site.
  This could include features such as automatically linking KEP and SIG names.
* Over time, add an index of events, meetups, and other forums for those that are actively contributing to k8s.

### Non-Goals

* Discover and bring together information from multiple orgs/repos.
* Create a super dynamic back end.  This is most likely served best with a static site.
* Other extended community functions like a job board or a list of vendors.

## Proposal

TBD.


### Risks and Mitigations

<!--
What are the risks of this proposal and how do we mitigate.
Think broadly.
For example, consider both security and how this will impact the larger kubernetes ecosystem.
-->

## Graduation Criteria

<!--
How will we know that this has succeeded?
Gathering user feedback is crucial for building high quality experiences and SIGs have the important responsibility of setting milestones for stability and completeness.
Hopefully the content previously contained in [umbrella issues][] will be tracked in the `Graduation Criteria` section.

[umbrella issues]: https://github.com/kubernetes/kubernetes/issues/42752
-->

## Implementation History

## Drawbacks [optional]

<!--
Why should this KEP _not_ be implemented.
-->

## Alternatives [optional]

<!--
Similar to the `Drawbacks` section the `Alternatives` section is used to highlight and record other possible approaches to delivering the value proposed by a KEP.
-->