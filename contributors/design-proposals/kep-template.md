# KEP Template

## Metadata

## Table of Contents

- [KEP Template](#kep-template)
   - [Metadata](#metadata)
   - [Summary](#summary)
   - [Motivation](#motivation)
   - [Examples](#examples)
   - [Detailed Design](#detailed-design)
   - [How Do I Teach This?](#how-do-i-teach-this)
      - [KEP Template](#kep-template-1)
      - [Section: Metadata](#section-metadata)
      - [Section: Table of contents](#section-table-of-contents)
      - [Section: Summary](#section-summary)
      - [Section: Motivation](#section-motivation)
      - [Section: Examples](#section-examples)
      - [Section: How Do I Teach This?](#section-how-do-i-teach-this)
      - [Section: Detailed Design](#section-detailed-design)
      - [Section: Graduation Criteria](#section-graduation-criteria)
      - [Section: Implementation History](#section-implementation-history)
      - [Section: Drawbacks](#section-drawbacks)
      - [Section: Alternatives](#section-alternatives)
      - [Section: Unresolved Questions](#section-unresolved-questions)
      - [Section: Mentors](#section-mentors)
   - [Graduation Criteria](#graduation-criteria)
   - [Drawbacks](#drawbacks)
   - [Alternatives](#alternatives)
   - [Unresolved Questions](#unresolved-questions)
   - [Mentors](#mentors)

## Summary

The KEP template combines our design proposals with functionality currently
provided by the [features repo][] within a flat file in source control. SIGs
will adopt and fork the KEP template to support their specific needs.

[features repo]: https://github.com/kubernetes/features

## Motivation

A template for a unit of work described in a [proposal to standardize][] the
Kubernetes development process is provided.

[proposal to standardize]: https://github.com/kubernetes/community/pull/967

## Examples

- [this KEP][]
- [the KEP process][]

[this KEP]: https://github.com/calebamiles/community/blob/propose-kep-template/contributors/design-proposals/kep-template.md
[the KEP process]: https://github.com/kubernetes/community/pull/967

## Detailed Design

## How Do I Teach This?

[The KEP process][] proposal will outline how a KEP is expected to progress
from conception to implementation in one or more pull requests to a Kubernetes
repository. This suggests that a KEP template should be used to track work
which ultimately will result in change to source control; if you find success
in modifying the KEP template please notify the [mentors](#mentors)!

### KEP Template

In order to implement an KEP like process the following template will be created

```
# Title
## Metadata
## Table of Contents
## Summary
## Motivation
## Examples [optional]
## How Do I Teach This? [optional]
## Detailed Design
## Graduation Criteria
## Implementation History
## Drawbacks [optional]
## Alternatives [optional]
## Unresolved Questions [optional]
## Mentors [optional]
```

where sections marked `[optional]` may be omitted by KEP authors. It is expected
that most KEPs will begin their lives in discussion with the responsible SIGs,
preferably in a mailing list so that the discussion is saved for future Kubernauts.


### Section: Metadata

The `Metadata` section is intended to support the creation of tooling around the
KEP process. The precise format for `Metadata` is described in the
[metadata proposal][].

[metadata proposal]: https://docs.google.com/document/d/1ynmBMuDuT7yGzRscObB1KtgJj8ljYq0I5q4oshrJUCs/edit#

### Section: Table of contents

A table of contents is helpful for quickly jumping to sections of a KEP and for
highlighting any addtional information provided beyond the standard KEP
template. [Tools for generating][] a table of contents from markdown are
available.

[Tools for generating]: https://github.com/ekalinin/github-markdown-toc


### Section: Summary

The `Summary` section is incredibly important for producing high quality user
focused documentation such as release notes or a development road map. It should
be possible to collect this information before implementation begins in order
to avoid requiring implementors to split their attention between writing
release notes and implementing the feature itself. KEP editors, SIG Docs, and
SIG PM should help to ensure that the tone and content of the `Summary` section
is useful for a wide audience.

### Section: Motivation

The `Motivation` section should describe

- why we believe this change is important
- what benefits are expected to be realized from the change
- the high level design goals

The `Motivation` section is important for getting all responsible parties to
understand the intention behind a change. The motivation section can optionally
provide links to [experience reports][] to demonstrate the interest in a KEP
within the wider Kubernetes community.

[experience reports]: https://github.com/golang/go/wiki/ExperienceReports

### Section: Examples

The `Examples` section should describe the intended user experience for an
KEP.

### Section: How Do I Teach This?

### Section: Detailed Design

The `Detailed Design` section should contain enough information for someone
besides the author to begin working on an implementation of the KEP. In a
similar manner to the guidance on [implementing an RFC][] from the Rust
community, not all KEPs must be implemented immediately. A high quality
`Detailed Design` section will enable any interested contributor to begin
implementation. Associating each KEP with one or more issues filed against
Kubernetes repositories will allow interested community members to track
implementation.

[implementing an RFC]: https://github.com/rust-lang/rfcs/blob/master/README.md#implementing-an-rfc

### Section: Graduation Criteria

Software development provides a unique opportunity for an iterative approach to
value delivery. Gathering user feedback is crucial for building high quality
experiences and SIGs have the important responsibility of setting milestones
for stability and feature completeness. Hopefully the content previously
contained in [umbrella issues][] will be tracked in the `Graduation Criteria`
section.

[umbrella issues]: https://github.com/kubernetes/kubernetes/issues/42752

### Section: Implementation History

Major milestones in the life cycle of a KEP should be tracked in
`Implementation History`. Major milestones might include

- the `Summary` and `Motivation` sections being merged signaling SIG acceptance
- the `Detailed Design` section being merged signaling agreement on a proposed
  design
- the date implementation started
- the first Kubernetes release where an initial version of the KEP was available
- the version of Kubneretes where the KEP graduated to general availability
- when the KEP was retired or superseded

### Section: Drawbacks

Making trade offs is one of the fundamental tasks of software development. The
`Drawbacks` section is to record what explicit choices have been made in the
design for a KEP.

### Section: Alternatives

Similar to the `Drawbacks` section the `Alternatives` section is used to
highlight and record other possible approaches to delivering the value proposed
by a KEP.

### Section: Unresolved Questions

The `Unresolved Questions` section is used to parking lot issues not ready to be
addressed before implementation begins.

### Section: Mentors

Mentors who can help a community member implement a KEP which follows its
`Detailed Design` are crucial to scaling the Kubernetes project. Potential
mentors can list their contact information using their preferred contact
information in the `Mentors` section.

## Graduation Criteria

The first revision of the KEP template will be considered stable once

- SIGs adopt the KEP process and template or fork the template
- the KEP process has been used for at least two Kubernetes releases

## Drawbacks

- having many possible sections in a KEP may feel too much like the waterfall
  approach to software development

## Alternatives

The KEP template is similar to the API proposal [design template][] but is
slightly more general given that it targets the subset of information likely
required for any significant change. Additionally

- The `Overview subsequent versions` has been replaced with the
  `Graduation Criteria` section of a KEP. It is intended that graduation
  criteria are added incrementally as a KEP is implemented and stabilized
- The `Motivation` and `Use Cases` sections have been collapsed within
  the `Motivation` section of a KEP. It is expected that experience reports
  will be the primary description of use cases
- `Requirements` is replaced by `Detailed Design` within a KEP

[design template]: https://github.com/pwittrock/community/blob/design/contributors/design-proposals/api-proposal-design-template.md

## Unresolved Questions

- How precisely merging portions of a KEP relate to the states of a KEP
- Formal mechanism for adopting the KEP process and template
- Automation

## Mentors

- caleb miles
  - github: [calebamiles](https://github.com/calebamiles/)
  - slack: [calebamiles](https://coreos.slack.com/team/caleb.miles)
  - email: [caleb.miles@coreos.com](mailto:caleb.miles@coreos.com)
  - pronoun: "he"
