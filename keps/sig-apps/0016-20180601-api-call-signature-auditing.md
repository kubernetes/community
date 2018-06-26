---
kep-number: 16
title: APISnoop / cluster-wide api usage logging
authors:
  - "@hh"
owning-sig: sig-apps
participating-sigs:
  - sig-apps
  - sig-testing
  - sig-architecture
reviewers:
  - "@AishSundar"
approvers:
  - "@WilliamDenniss"
editor: TBD
creation-date: 2018-06-27
last-updated: 2018-06-27
status: provisional
see-also:
  - KEP-15
---

# API Call Signature Auditing

## Table of Contents

[Tools for generating][] a table of contents from markdown are available.

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

[Tools for generating]: https://github.com/ekalinin/github-markdown-toc

## Summary

The `Summary` section is incredibly important for producing high quality user focused documentation such as release notes or a development road map.
It should be possible to collect this information before implementation begins in order to avoid requiring implementors to split their attention between writing release notes and implementing the feature itself.
KEP editors, SIG Docs, and SIG PM should help to ensure that the tone and content of the `Summary` section is useful for a wide audience.

A good summary is probably at least a paragraph in length.

## Motivation

This section is for explicitly listing the motivation, goals and non-goals of this KEP.
Describe why the change is important and the benefits to users.
The motivation section can optionally provide links to [experience reports][] to demonstrate the interest in a KEP within the wider Kubernetes community.

[experience reports]: https://github.com/golang/go/wiki/ExperienceReports

### Goals

List the specific goals of the KEP.
How will we know that this has succeeded?

### Non-Goals

What is out of scope for his KEP?
Listing non-goals helps to focus discussion and make progress.

## Proposal

This is where we get down to the nitty gritty of what the proposal actually is.

### User Stories [optional]

Detail the things that people will be able to do if this KEP is implemented.
Include as much detail as possible so that people can understand the "how" of the system.
The goal here is to make this feel real for users without getting bogged down.

#### Story 1

#### Story 2

### Implementation Details/Notes/Constraints [optional]

What are the caveats to the implementation?
What are some important details that didn't come across above.
Go in to as much detail as necessary here.
This might be a good place to talk about core concepts and how they releate.

### Risks and Mitigations

What are the risks of this proposal and how do we mitigate.
Think broadly.
For example, consider both security and how this will impact the larger kubernetes ecosystem.

## Graduation Criteria

How will we know that this has succeeded?
Gathering user feedback is crucial for building high quality experiences and SIGs have the important responsibility of setting milestones for stability and completeness.
Hopefully the content previously contained in [umbrella issues][] will be tracked in the `Graduation Criteria` section.

[umbrella issues]: https://github.com/kubernetes/kubernetes/issues/42752

## Implementation History

Major milestones in the life cycle of a KEP should be tracked in `Implementation History`.
Major milestones might include

- the `Summary` and `Motivation` sections being merged signaling SIG acceptance
- the `Proposal` section being merged signaling agreement on a proposed design
- the date implementation started
- the first Kubernetes release where an initial version of the KEP was available
- the version of Kubernetes where the KEP graduated to general availability
- when the KEP was retired or superseded

## Drawbacks [optional]

Why should this KEP _not_ be implemented.

## Alternatives [optional]

Similar to the `Drawbacks` section the `Alternatives` section is used to highlight and record other possible approaches to delivering the value proposed by a KEP.
