---
kep-number: draft-YYYYMMDD
title: My First KEP
authors:
  - "@janedoe"
owning-sig: sig-xxx
participating-sigs:
  - sig-aaa
  - sig-bbb
reviewers:
  - TBD
  - "@alicedoe"
approvers:
  - TBD
  - "@oscardoe"
editor: TBD
creation-date: yyyy-mm-dd
last-updated: yyyy-mm-dd
status: accepted
see-also:
  - KEP-1
  - KEP-2
replaces:
  - KEP-3
superseded-by:
  - KEP-100
---

# Title

This is the title of the KEP.
Keep it simple and descriptive.
A good title can help communicate what the KEP is and should be considered as part of any review.

The *filename* for the KEP should include the KEP number along with the title.
The title should be lowercased and spaces/punctuation should be replaced with `-`.
As the KEP is approved and an official KEP number is allocated, the file should be renamed.

To get started with this template:
* Make a copy of this template.
  Name it `draft-YYYYMMDD-my-title.md`.
* Create a PR in the [`kubernetes/community`](https://github.com/kubernetes/community) repo.
* Check in early.
  Do this once the document holds together and general direction is understood by many in the sponsoring SIG.
  View anything marked as a draft as a working document.
  Aim for single topic PRs to keep discussions focused.
  If you disagree with what is already in a document, open a new PR with suggested changes.
* A KEP number can be assigned at any time by (even in the first PR) (a) taking the next number in the NEXT_KEP_NUMBER file and (b) incrementing that number.
  These PRs should be approved quickly to minimize merge conflicts.

The canonical place for the latest set of instructions (and the likely source of this file) is [here](/keps/0000-kep-template.md).

The `Metadata` section above is intended to support the creation of tooling around the KEP process.
This will be a YAML section that is fenced as a code block.
See the KEP process for details on each of these items.

## Table of Contents

A table of contents is helpful for quickly jumping to sections of a KEP and for highlighting any additional information provided beyond the standard KEP template.
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
