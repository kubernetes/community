# Title

This is the title of the KEP.  Keep it simple and descriptive. A good title can
help communicate what the KEP is and should be considered as part of any review.

The *filename* for the KEP should include the KEP number along with the title.
The title should be lowercased and spaces/punctuation should be replaced with
`-`. As the KEP is approved and an official KEP number is allocated, the file
should be renamed.

To get started with this template:
* Make a copy in the appropriate directory.  Name it `draft-YYYYMMDD-my-title.md`.
* Create a PR in the
  [`kubernetes/community`](https://github.com/kubernetes/community) repo.
* Check in early.  Do this once the document holds together and general
  direction is understood by many in the sponsoring SIG. View anything marked as
  a draft as a working document.  Aim for single topic PRs to keep discussions
  focused. If you disagree with what is already in a document, open a new PR
  with suggested changes.
* As a KEP is approved, rename the file yet again with the final KEP number.

The canonical place for the latest set of instructions (and the likely source of
this file) is
[here](https://github.com/kubernetes/community/blob/master/contributors/design-proposals/architecture/0000-kep-template.md).

## Metadata

The `Metadata` section is intended to support the creation of tooling around the
KEP process.  This will be a YAML section that is fenced as a code block.

See the KEP process for details on each of these items.  This is here for easy
copy/pasting.

TODO(jbeda): Do we want to move this to the front the doc with a delimiter
(`---`) so it is easier to parse.  Many static site generators use this and call
it "front matter".

TODO(jbeda): Do we want to have a "people database" to reduce the amount of
duplication on naming people here?  This would be a simple map of github ID to
name and contact info.

```yaml
kep-number: draft-XXX
title: My First KEP
authors:
  - name: Jane Doe
    github: janedoe
    email: janedoe@example.com
owning-sig: sig-xxx
participating-sigs:
  - sig-aaa
  - sig-bbb
reviewers:
  - name: TBD
  # - name: Alice Doe
  #   github: alicedoe
  #   email: alicedoe@example.com
approvers:
  - name: TBD
  # - name: Oscar Doe
  #   github: oscardoe
  #   email: oscardoe@example.com
editor:
  name: TBD
creation-date: yyyy-mm-dd
last-updated: yyyy-mm-dd
status: draft
see-also:
  - KEP-1
  - KEP-2
replaces:
  - KEP-3
superseded-by:
  - KEP-100
```

## Table of Contents

A table of contents is helpful for quickly jumping to sections of a KEP and for
highlighting any addtional information provided beyond the standard KEP
template. [Tools for generating][] a table of contents from markdown are
available.

[Tools for generating]: https://github.com/ekalinin/github-markdown-toc

## Summary

The `Summary` section is incredibly important for producing high quality user
focused documentation such as release notes or a development road map. It should
be possible to collect this information before implementation begins in order
to avoid requiring implementors to split their attention between writing
release notes and implementing the feature itself. KEP editors, SIG Docs, and
SIG PM should help to ensure that the tone and content of the `Summary` section
is useful for a wide audience.

A good summary is probably at least a paragraph in length.

## Motivation

The `Motivation` section should describe

- why we believe this change is important
- what benefits are expected to be realized from the change
- the high level design goals

The `Motivation` section is important for getting all responsible parties to
understand the intention behind a change. The motivation section can optionally
provide links to [experience reports][] to demonstrate the interest in a KEP
within the wider Kubernetes community.

[experience reports]: https://github.com/golang/go/wiki/ExperienceReports

## Guide-level Explanation [optional]

Merging a change to source control is a crucial, but not final, milestone in
the implementation of a KEP. Enhancements need to be explained to the Kubernetes
community. The `Guide-level Explaination` section should be used to explain a
KEP to another Kubernaut after implementation. Excellent guidance can be
found in the Rust RFC [guide-level explanation][] instructions.


[guide-level explanation]: https://github.com/rust-lang/rfcs/blob/master/0000-template.md#guide-level-explanation


## Reference-level explanation

Before submitting a detailed implementation plan, a KEP author might begin the
`Reference-level Explaination` by sketching high level design goals and any
mandatory requirements.

Communicating dependencies across multiple SIGs is an important use for KEPs.
Explaining how a KEP interacts with other KEPs and existing Kubernetes
functionality should be included in this section.

The `Reference-level explaination` section should ideally contain enough
information for someone besides the author to begin working on an implementation
of the KEP. In a similar manner to the guidance on [implementing an RFC][] from
the Rust community, not all KEPs must be implemented immediately. Associating
each KEP with one or more issues filed against Kubernetes repositories allows
interested community members to track implementation.

Excellent guidance can be found in the Rust RFC [reference-level explanation][]
instructions.

[reference-level explaination]: https://github.com/rust-lang/rfcs/blob/master/0000-template.md#reference-level-explanation

[implementing an RFC]: https://github.com/rust-lang/rfcs/blob/master/README.md#implementing-an-rfc

## Graduation Criteria

Gathering user feedback is crucial for building high quality experiences and
SIGs have the important responsibility of setting milestones for stability
and completeness. Hopefully the content previously contained in
[umbrella issues][] will be tracked in the `Graduation Criteria` section.

[umbrella issues]: https://github.com/kubernetes/kubernetes/issues/42752

## Implementation History

Major milestones in the life cycle of a KEP should be tracked in
`Implementation History`. Major milestones might include

- the `Summary` and `Motivation` sections being merged signaling SIG acceptance
- the `Detailed Design` section being merged signaling agreement on a proposed
  design
- the date implementation started
- the first Kubernetes release where an initial version of the KEP was available
- the version of Kubernetes where the KEP graduated to general availability
- when the KEP was retired or superseded

## Drawbacks [optional]

Why should this KEP _not_ be implemented.

## Alternatives [optional]

Similar to the `Drawbacks` section the `Alternatives` section is used to
highlight and record other possible approaches to delivering the value proposed
by a KEP.

## Unresolved Questions [optional]

The `Unresolved Questions` section is used to parking lot issues not ready to be
addressed before implementation begins.

## Mentors [optional]

Mentors who can help a community member implement a KEP which follows its
`Detailed Design` are crucial to scaling the Kubernetes project. Potential
mentors can list their contact information using their preferred contact
information in the `Mentors` section.
