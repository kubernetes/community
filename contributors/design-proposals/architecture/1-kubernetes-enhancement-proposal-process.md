# Kubernetes Enhancement Proposal Process

## Metadata
```
---
kep-number: 1
title: Kubernetes Enhancement Proposal Process
authors:
  - name: Caleb Miles
    github: calebamiles
    slack: calebamiles
  - name: Joe Beda
    github: jbeda
    email: joe@heptio.com
    slack: jbeda
owning-sig: sig-architecture
participating-sigs:
  - `kubernetes-wide`
reviewers:
  - name: TBD
approvers:
  - name: TBD
editor:
  name: TBD
creation-date: 2017-08-22
status: draft
```

## Table of Contents

* [Kubernetes Enhancement Proposal Process](#kubernetes-enhancement-proposal-process)
  * [Metadata](#metadata)
  * [Table of Contents](#table-of-contents)
  * [Summary](#summary)
  * [Motivation](#motivation)
  * [Reference-level explanation](#reference-level-explanation)
      * [What type of work should be tracked by a KEP](#what-type-of-work-should-be-tracked-by-a-kep)
      * [KEP Template](#kep-template)
      * [KEP Metadata](#kep-metadata)
      * [KEP Workflow](#kep-workflow)
      * [Git and GitHub Implementation](#git-and-github-implementation)
      * [KEP Editor Role](#kep-editor-role)
      * [Important Metrics](#important-metrics)
      * [Prior Art](#prior-art)
  * [Graduation Criteria](#graduation-criteria)
  * [Drawbacks](#drawbacks)
  * [Alternatives](#alternatives)
  * [Unresolved Questions](#unresolved-questions)
  * [Mentors](#mentors)

## Summary

A standardized development process for Kubernetes is proposed in order to

- provide a common structure for proposing changes to Kubernetes
- ensure that the motivation for a change is clear
- allow for the enumeration stability milestones and stability graduation
  criteria
- persist project information in a Version Control System (VCS) for future
  Kubernauts
- support the creation of _high value user facing_ information such as:
  - an overall project development roadmap
  - motivation for impactful user facing changes
- support development across multiple repositories beyond `kubernetes/kubernetes`
- reserve GitHub issues for tracking work in flight rather than creating "umbrella"
  issues
- ensure community participants are successfully able to drive changes to
  completion across one or more releases while stakeholders are adequately
  represented throughout the process

This process is supported by a unit of work called a Kubernetes Enhancement
Proposal or KEP. A KEP attempts to combine aspects of a

- feature, and effort tracking document
- a product requirements document
- design document

into one file which is created incrementally in collaboration with one or more
Special Interest Groups (SIGs).

## Motivation

For cross project SIGs such as SIG PM and SIG Release an abstraction beyond a
single GitHub Issue or Pull request seems to be required in order to understand
and communicate upcoming changes to Kubernetes.  In a blog post describing the
[road to Go 2][], Russ Cox explains

> that it is difficult but essential to describe the significance of a problem
> in a way that someone working in a different environment can understand

as a project it is vital to be able to track the chain of custody for a proposed
enhancement from conception through implementation. This proposal does not
attempt to mandate how SIGs track their work internally, however, it is
suggested that SIGs which do not adhere to a process which allows for their hard
work to be explained to others in the wider Kubernetes community will see their
work wallow in the shadows of obscurity. At the very least [survey data][]
suggest that high quality documentation is crucial to project adoption.
Documentation can take many forms and it is imperative to ensure that it is easy
to produce high quality user or developer focused documentation for a complex
project like Kubernetes.

Without a standardized mechanism for describing important enhancements our
talented technical writers and product managers struggle to weave a coherent
narrative explaining why a particular release is important. Additionally for
critical infrastructure such as Kubernetes adopters need a forward looking road
map in order to plan their adoption strategy.

The purpose of the KEP process is to reduce the amount of "tribal knowledge" in
our community. By moving decisions from a smattering of mailing lists, video
calls and hallway conversations into a well tracked artifact this process aims
to enhance communication and discoverability.

A KEP is broken into sections which can be merged into source control
incrementally in order to support an iterative development process. An important
goal of the KEP process is ensuring that the process for submitting the content
contained in [design proposals][] is both clear and efficient. The KEP process
is intended to create high quality uniform design and implementation documents
for SIGs to deliberate.

[tell a story]: https://blog.rust-lang.org/2017/08/31/Rust-1.20.html
[road to Go 2]: https://blog.golang.org/toward-go2
[survey data]: http://opensourcesurvey.org/2017/
[design proposals]: https://github.com/kubernetes/community/tree/master/contributors/design-proposals


## Reference-level explanation

### What type of work should be tracked by a KEP

The definition of what constitutes an "enhancement" is a foundational concern
for the Kubernetes project. Roughly any Kubernetes user or operator facing
enhancement should follow the KEP process: if an enhancement would be described
in either written or verbal communication to anyone besides the KEP author or
developer then consider creating a KEP. One concrete example is an enhancement
which should be communicated to SIG Release or SIG PM.

Similarly, any technical effort (refactoring, major architectural change) that
will impact a large section of the development community should also be
communicated widely. The KEP process is suited for this even if it will have
zero impact on the typical user or operator.

As the local bodies of governance, SIGs should have broad latitude in describing
what constitutes an enhancement which should be tracked through the KEP process.
SIGs may find that helpful to enumerate what _does not_ require a KEP rather
than what does. SIGs also have the freedom to customize the KEP template
according to their SIG specific concerns. For example the KEP template used to
track API changes will likely have different subsections than the template for
proposing governance changes. However, as changes start impacting other SIGs or
the larger developer community outside of a SIG, the KEP process should be used
to coordinate and communicate.

### KEP Template

The template for a KEP is precisely defined in the [template proposal][]

[template proposal]: https://github.com/kubernetes/community/pull/1124

### KEP Metadata

There is a place in each KEP for a YAML document that has standard metadata.
This will be used to support tooling around filtering and display.  It is also
critical to clearly communicate the status of a KEP.

Metadata items:
* **kep-number** Required
  * Each proposal has a number.  This is to make all references to proposals as
    clear as possible.  This is especially important as we create a network
    cross references between proposals.
  * Before having the `Approved` status, the number for the KEP will be in the
    form of `draft-YYYYMMDD`.  The `YYYYMMDD` is replaced with the current date
    when first creating the KEP.  The goal is to enable fast parallel merges of
    pre-acceptance KEPs.
  * On acceptance a sequential dense number will be assigned.  This will be done
    by the editor and will be done in such a way as to minimize the chances of
    conflicts.  The final number for a KEP will have no prefix.
* **title** Required
  * The title of the KEP in plain language.  The title will also be used in the
    KEP filename.  See the template for instructions and details.
* **status** Required
  * The current state of the KEP.
  * Must be one of `Draft`, `Deferred`, `Approved`, `Rejected`, `Withdrawn`,
    `Final`, `Replaced`.
* **authors** Required
  * A list of authors for the KEP.  We require a name (which can be a psuedonym)
    along with a github ID.  Other ways to contact the author is strongly
    encouraged.  This is a list of maps.  Subkeys of each item: `name`,
    `github`, `email` (optional), `slack` (optional).
* **owning-sig** Required
  * The SIG that is most closely associated with this KEP. If there is code or
    other artifacts that will result from this KEP, then it is expected that
    this SIG will take responsiblity for the bulk of those artificats.
  * Sigs are listed as `sig-abc-def` where the name matches up with the
    directory in the `kubernetes/community` repo.
* **participating-sigs** Optional
  * A list of SIGs that are involved or impacted by this KEP.
  * A special value of `kubernetes-wide` will indicate that this KEP has impact
    across the entire project.
* **reviewers** Required
  * Reviewer(s) chosen after triage according to proposal process
  * If not yet chosen replace with `TBD`
  * Same name/contact scheme as `authors`
* **approvers** Required
  * Approver(s) chosen after triage according to proposal process
  * If not yet chosen replace with `TBD`
  * Same name/contact scheme as `authors`
* **editor** Required
  * Someone to keep things moving forward.
  * If not yet chosen replace with `TBD`
  * Same name/contact scheme as `authors`
* **creation-date** Required
  * The date that the KEP was first submitted in a PR.
  * In the form `yyyy-mm-dd`
  * While this info will also be in source control, it is helpful to have the set of KEP files stand on their own.
* **last-updated** Optional
  * The date that the KEP was last changed significantly.
  * In the form `yyyy-mm-dd`
* **see-also** Optional
  * A list of other KEPs that are relevant to this KEP.
  * In the form `KEP-123`
* **replaces** Optional
  * A list of KEPs that this KEP replaces.  Those KEPs should list this KEP in
    their `superceded-by`.
  * In the form `KEP-123`
* **superseded-by**
  * A list of KEPs that superced this KEP. Use of this should be paired with
    this KEP moving into the `Replaced` status.
  * In the form `KEP-123`


### KEP Workflow

TODO(jbeda) Rationalize this with status entires in the Metadata above.

A KEP is proposed to have the following states

- **opened**: a new KEP has been filed but not triaged by the responsible SIG or
  working group
- **accepted**: the motivation has been accepted by the SIG or working group as in
  road map
- **scoped**: the design has been approved by the SIG or working group
- **started**: the implementation of the KEP has begun
- **implemented**: the implementation of the KEP is complete
- **deferred**: the KEP has been postponed by the SIG or working group despite
  agreement on the motivation
- **superseded**: the KEP has been superseded by another KEP
- **retired**: the implementation of the KEP has been removed
- **rejected**: the KEP has been rejected by the SIG or working group
- **orphaned**: the author or developer of the KEP is no longer willing or able
  to complete implementation

with possible paths through the state space

- opened -> deferred (a)
- opened -> rejected (b)
- opened -> orphaned (c)
- opened -> accepted -> orphaned (d)
- opened -> accepted -> scoped -> superseded (e)
- opened -> accepted -> scoped -> orphaned (f)
- opened -> accepted -> scoped -> started -> retired (g)
- opened -> accepted -> scoped -> started -> orphaned (h)
- opened -> accepted -> scoped -> started -> superseded (i)
- opened -> accepted -> scoped -> started -> implemented (j)
- opened -> accepted -> scoped -> started -> implemented -> retired (k)

the happy path is denoted by (j) where an KEP is opened; accepted by a SIG as in
their roadmap; fleshed out with a design; started; and finally implemented. As
Kubernetes continues to mature, hopefully metrics on the utilization of features
will drive decisions on what features to maintain and which to deprecate and so
it is possible that a KEP would be retired if its functionality no longer provides
sufficient value to the community.

### Git and GitHub Implementation

Practically an KEP would be implemented as a pull request to a central repository
with the following example structure

```
├── 0000-kep-template.md
├── CODEOWNERS
├── index.md
├── sig-architecture
│   ├── deferred
│   ├── orphaned
│   └── retired
├── sig-network
│   ├── deferred
│   ├── kube-dns
│   ├── orphaned
│   └── retired
├── sig-node
│   ├── deferred
│   ├── kubelet
│   ├── orphaned
│   └── retired
├── sig-release
│   ├── deferred
│   ├── orphaned
│   └── retired
├── sig-storage
│   ├── deferred
│   ├── orphaned
│   └── retired
├── unsorted-to-be-used-by-newcomers-only
└── wg-resource-management
    ├── deferred
    ├── orphaned
    └── retired
```

where each SIG or working group is given a top level directory with subprojects
maintained by the SIG listed in sub directories. For newcomers to the community
an `unsorted-to-be-used-by-newcomers-only` directory may be used before an KEP
can be properly routed to a SIG although hopefully if discussion for a potential
KEP begins on the mailing lists proper routing information will be provided to
the KEP author. Additionally a top level index of KEPs may be helpful for people
looking for a complete list of KEPs. There should be basic CI to ensure that an
`index.md` remains up to date.

Ideally no work would begin within the repositories of the Kubernetes organization
before a KEP has been approved by the responsible SIG or working group. While the
details of how SIGs organize their work is beyond the scope of this proposal one
possibility would be for each charter SIG to create a top level repository within
the Kubernetes org where implementation issues managed by that SIG would be filed.

### KEP Editor Role

Taking a cue from the [Python PEP process][], I believe that a group of KEP editors
will be required to make this process successful; the job of an KEP editor is
likely very similar to the [PEP editor responsibilities][] and will hopefully
provide another opportunity for people who do not write code daily to contribute
to Kubernetes.

In keeping with the PEP editors which

> Read the PEP to check if it is ready: sound and complete. The ideas must make
> technical sense, even if they don't seem likely to be accepted.
> The title should accurately describe the content.
> Edit the PEP for language (spelling, grammar, sentence structure, etc.), markup
> (for reST PEPs), code style (examples should match PEP 8 & 7).

KEP editors should generally not pass judgement on a KEP beyond editorial
corrections.

[Python PEP process]: https://www.python.org/dev/peps/pep-0001/
[PEP editor responsibilities]: https://www.python.org/dev/peps/pep-0001/#pep-editor-responsibilities-workflow

### Important Metrics

It is proposed that the primary metrics which would signal the success or
failure of the KEP process are

- how many "features" are tracked with a KEP
- distribution of time a KEP spends in each state
- KEP rejection rate
- PRs referencing a KEP merged per week
- number of issued open which reference a KEP
- number of contributors who authored a KEP
- number of contributors who authored a KEP for the first time
- number of orphaned KEPs
- number of retired KEPs
- number of superseded KEPs

### Prior Art

The KEP process as proposed was essentially stolen from the [Rust RFC process] which
itself seems to be very similar to the [Python PEP process][]

[Rust RFC process]: https://github.com/rust-lang/rfcs

## Graduation Criteria

should hit at least the following milestones

- a release note draft can be generated by referring primarily to KEP content
- a yearly road map is expressed as a KEP

## Drawbacks

Any additional process has the potential to engender resentment within the
community. There is also a risk that the KEP process as designed will not
sufficiently address the scaling challenges we face today. PR review bandwidth is
already at a premium and we may find that the KEP process introduces an unreasonable
bottleneck on our development velocity.

It certainly can be argued that the lack of a dedicated issue/defect tracker
beyond GitHub issues contributes to our challenges in managing a project as large
as Kubernetes, however, given that other large organizations, including GitHub
itself, make effective use of GitHub issues perhaps the argument is overblown.

The centrality of Git and GitHub within the KEP process also may place too high
a barrier to potential contributors, however, given that both Git and GitHub are
required to contribute code changes to Kubernetes today perhaps it would be reasonable
to invest in providing support to those unfamiliar with this tooling.

Expanding the proposal template beyond the single sentence description currently
required in the [features issue template][] may be a heavy burden for non native
English speakers and here the role of the KEP editor combined with kindness and
empathy will be crucial to making the process successful.

[features issue template]: https://github.com/kubernetes/features/blob/master/ISSUE_TEMPLATE.md

## Alternatives

This KEP process is related to
- the generation of a [architectural roadmap][]
- the fact that the [what constitutes a feature][] is still undefined
- [issue management][]
- the difference between an [accepted design and a proposal][]
- [the organization of design proposals][]

this proposal attempts to place these concerns within a general framework.

[architectural roadmap]: https://github.com/kubernetes/community/issues/952
[what constitutes a feature]: https://github.com/kubernetes/community/issues/531
[issue management]: https://github.com/kubernetes/community/issues/580
[accepted design and a proposal]: https://github.com/kubernetes/community/issues/914
[the organization of design proposals]: https://github.com/kubernetes/community/issues/918

### Github issues vs. KEPs

The use of GitHub issues when proposing changes does not provide SIGs good
facilities for signaling approval or rejection of a proposed change to Kubernetes
since anyone can open a GitHub issue at any time. Additionally managing a proposed
change across multiple releases is somewhat cumbersome as labels and milestones
need to be updated for every release that a change spans. These long lived GitHub
issues lead to an ever increasing number of issues open against
`kubernetes/features` which itself has become a management problem.

In addition to the challenge of managing issues over time, searching for text
within an issue can be challenging. The flat hierarchy of issues can also make
navigation and categorization tricky. While not all community members might
not be comfortable using Git directly, it is imperative that as a community we
work to educate people on a standard set of tools so they can take their
experience to other projects they may decide to work on in the future. While
git is a fantastic version control system (VCS), it is not a project management
tool nor a cogent way of managing an architectural catalog or backlog; this
proposal is limited to motivating the creation of a standardized definition of
work in order to facilitate project management. This primitive for describing
a unit of work may also allow contributors to create their own personalized
view of the state of the project while relying on Git and GitHub for consistency
and durable storage.

## Unresolved Questions

- How reviewers and approvers are assigned to a KEP
- Approval decision process for a KEP
- Example schedule, deadline, and time frame for each stage of a KEP
- Communication/notification mechanisms
- Review meetings and escalation procedure
- Decision on where development should occur

## Mentors

- caleb miles
  - github: [calebamiles](https://github.com/calebamiles/)
  - slack: [calebamiles](https://coreos.slack.com/team/caleb.miles)
  - email: [caleb.miles@coreos.com](mailto:caleb.miles@coreos.com)
  - pronoun: "he"
