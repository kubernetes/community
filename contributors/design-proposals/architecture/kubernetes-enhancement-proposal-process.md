# Kubernetes Enhancement Proposal Process

## Metadata
```
---
metadata:
  number: 0001
  state: opened
  authors:
  - author:
      name: caleb miles
      github: @calebamiles
      slack: @calebamiles
  owners:
  - sig-release
  - sig-pm
  - sig-architecture
  - sig-testing
  - steering-committee
  links:
    issues:
    - [someIssueURL]()
    prs:
    - https://github.com/kubernetes/community/pull/967
    discussions:
    - https://groups.google.com/forum/#!topic/kubernetes-dev/65A-3ULYPB0
    - https://groups.google.com/forum/#!topic/kubernetes-sig-architecture/t-1EqeEoLPA
    documentation:
    - [someDocsLinkURL]()
    related:
    - [KEP template](https://github.com/kubernetes/community/pull/1124)
```

## Summary

A standardized development process for Kubernetes is proposed in order to:

- provide a common structure for proposing changes to Kubernetes
- ensure that the motivation for a change is clear
- allow for the enumeration stability milestones and stability graduation
  criteria
- persist project information in a Version Control System (VCS) for future
  Kubernauts
- support the creation of _high value user facing_ information such as:
  - release notes
  - release announcement blog
  - an overall project development roadmap
- support development across multiple repositories beyond `kubernetes/kubernetes`
- reserve GitHub issues for tracking work in flight rather than creating "umbrella"
  issues
- ensure community participants are successfully able to drive changes to
  completion across one or more releases while stakeholders are adequately
  represented throughout the process

This process is supported by a unit of work called a Kubernetes Enhancement
Proposal or KEP.

## Motivation

For cross project SIGs such as SIG PM and SIG Release an abstraction beyond a
single GitHub Issue or Pull request seems to be required in order to understand
and communicate upcoming changes to Kubernetes. Particularly the generation of
release notes and the release announcement blog post are rather difficult and
have sometimes delayed a release due to incompleteness. In a blog post
describing the [road to Go 2][], Russ Cox explains

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

Ideally release notes should [tell a story][] which is compelling enough to
encourage users and operators to upgrade their clusters. Without a standardized
mechanism for describing important enhancements our talented technical writers
and product managers struggle to weave a coherent narrative explaining why a
particular release is important. Additionally for critical infrastructure such
as Kubernetes adopters need a forward looking road map in order to plan their
adoption strategy.

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


## Detailed design

### What type of work should be tracked by a KEP

The definition of what constitutes an "enhancement" is a foundational concern
for the Kubernetes project. Roughly any Kubernetes user or operator facing
enhancement should follow the KEP process: if an enhancement would be described
in either written or verbal communication to anyone besides the KEP author or
developer then consider creating a KEP. One concrete example is an enhancement
which should be communicated to SIG Release or SIG PM.

Without detailed information explaining the motivation for an enhancement SIGs
must first approve a proposal, agreeing to a motivation over a mailing list,
video call, or hallway conversation. During the release process this motivation
must be rediscovered by the SIG, hopefully by finding a design proposal. The
process of announcing an enhancement through release notes suggests another
heuristic for describing what work should be tracked through an KEP: anything
that would require a design proposal. In fact it is possible to consider a KEP
an enhancement to the design proposal process in which design proposals are
used throughout the process of proposing an enhancement, scoping its design,
tracking its implementation, and agreeing on criteria for graduation to general
availability.

As the local bodies of governance, SIGs should have broad latitude in describing
what constitutes an enhancement which should be tracked through the KEP process.
SIGs may find that helpful to enumerate  what _does not_ require a KEP rather
than what does. SIGs also have the freedom to customize the KEP template
according to their SIG specific concerns. For example the KEP template used 
to track API changes will likely have different subsections than the template
for proposing governance changes.

### KEP Template

The template for a KEP is precisely defined in the [template proposal][]

[template proposal]: https://github.com/kubernetes/community/pull/1124

### KEP Workflow

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
│   ├── kublet
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
