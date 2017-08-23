# Proposal to Standardize Development Process

## Links
```
---
Links:
- Issues:
  - [someIssueURL]()
- PRs:
  - https://github.com/kubernetes/community/pull/967
- MailingListDiscussions:
  - https://groups.google.com/forum/#!topic/kubernetes-dev/65A-3ULYPB0
- Documentation:
  - [someDocsLinkURL]()
```

## Responsible SIG(s)

- SIG Release
- SIG PM
- SIG Architecture
- SIG Governance
- Technical Steering Committee

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
  - a cross project development road map
- support development across multiple repositories beyond `kubernetes/kubernetes`
- begin to move away from using GitHub issues for feature requests and reserve them
  for tracking work in flight

## Motivation

Today it is becoming fairly difficult to track changes being made across all
Special Interest Groups (SIGs) within Kubernetes. For cross project SIGs such
as SIG PM and SIG Release an abstraction beyond a single GitHub Issue or Pull
request seems to be required in order to understand and communicate upcoming
changes to Kubernetes. Particularly the generation of release notes and the
release announcement blog post are rather difficult and have sometimes delayed
a release due to incompleteness. In a blog post describing the [road to Go 2][],
Russ Cox explains

> that it is difficult but essential to describe the significance of a problem
> in a way that someone working in a different environment can understand

unfortunately I believe that our current mechanism for proposing changes by
opening a GitHub issue against `kubernetes/features` falls short of providing
an effective vehicle for agreeing on the motivation and the design of a proposed
change.

The use of GitHub issues when proposing changes does not provide SIGs good
facilities for signaling approval or rejection of a proposed change to Kubernetes
since anyone can open a GitHub issue at any time. Additionally managing a proposed
change across multiple releases is somewhat cumbersome as labels and milestones
need to be updated for every release that a change spans which leads to an ever
increasing number of issues open against `kubernetes/features` which itself has
become a management problem.

We have chosen to use Git as our VCS with GitHub currently used to host our
repositories. While GitHub provides a rich feature set as an open source project,
managing issues over time and searching for text within an issue can be
challenging. The flat hierarchy of issues can also make navigation and
categorization tricky. While not all community members might not be comfortable
using Git directly, I believe it is imperative that as a community we work to
educate people on a standard set of tools so they can take their experience to
other projects they may decide to work on in the future.

It is my hope that by increasing the requirements for proposing a change to
Kubernetes we can prevent important work like generating release notes from
falling out of the release cycle and in general make working with and speaking
about Kubernetes an easier and more pleasant experience.

[road to Go 2]: https://blog.golang.org/toward-go2

## Detailed Design

### RFC Template

In order to implement an RFC like process a template of the following form would
be created

```
# Title
## Links
## Responsible SIG(s)
## Summary
## Motivation
## Examples [optional]
## Detailed Design
## Graduation Criteria
## Drawbacks [optional]
## Alternatives [optional]
## Unresolved Questions [optional]
## Mentors [optional]
```

where sections marked `[optional]` may be omitted by RFC authors.It is expected
that most RFCs will begin their lives in discussion with the responsible SIGs,
preferably in a mailing list so that the discussion is saved for future Kubernauts.

It may be very helpful to provide links to experience reports of other Kubernetes
users and developers within the `Motivation` section in order to provide additional
context for the necessity of a proposed change. The `Examples` section could be
used to provide motivating examples for documentation writers looking to teach
others how to consume the proposed change once implemented. As an open source
community we must constantly be looking for ways to better serve newcomers and
a `Mentors` section could be used to help connect interested parties with elders
in the community who are able to mentor people towards an implementation.

It is intended for RFCs to be a conversation, the comment part of the acronym,
so it is not expected for every section of the template to be completed by the
time an RFC is filed, however, the `Summary`, `Motivation`, and `Examples` section
should probably be filled out by the time an RFC is filed. Merging an RFC only
signals agreement of the sections which have been provided thus far and an RFC
may be completed through a few rounds of review in order to achieve alignment.

### Git and GitHub Implementation

Practically an RFC would be implemented as a pull request to a central repository
with the following example structure

```
├── 0000-rfc-template.md
├── CODEOWNERS
├── index.md
├── sig-architecture
├── sig-network
│   └── kube-dns
├── sig-node
│   └── kubelet
├── sig-release
├── sig-storage
├── unsorted-to-be-used-by-newcomers-only
└── wg-resource-management
```

where each SIG or working group is given a top level directory with components
maintained by the SIG listed in sub directories. For newcomers to the community
an `unsorted-to-be-used-by-newcomers-only` directory may be used before an RFC
can be properly routed to a SIG although hopefully if discussion for a potential
RFC begins on the mailing lists proper routing information will be provided to
the RFC author. Additionally a top level index of RFCs may be helpful for people
looking for a complete list of RFCs. There should be basic CI to ensure that a
`index.md` remains up to date.

Ideally no work would begin on an RFC before it has been approved by the
responsible SIG or working group. In order to help combat the explosion of GitHub
issues against `kubernetes/kubernetes` it is further proposed for each chartered
SIG to create a top level repository within the `kubernetes` GitHub organization
where implementation issues would be filed and linked to in the RFC.

### RFC Editor Role

Taking a cue from the [Python PEP process][], I believe that a group of RFC editors
will be required to make this process successful; the job of an RFC editor is
likely very similar to the [PEP editor responsibilities][] and will hopefully
provide another opportunity for people who do not write code daily to contribute
to Kubernetes.

[Python PEP process]: https://www.python.org/dev/peps/pep-0001/
[PEP editor responsibilities]: https://www.python.org/dev/peps/pep-0001/#pep-editor-responsibilities-workflow

### Important Metrics

I believe the primary metrics which would signal the success or failure of the RFC
process would be

- distribution of RFC merge times
- RFC rejection rate
- distribution of time between RFC acceptance and implementation issue creation
- PRs referencing an RFC merged per week

### Customization of the RFC Process

Given the general independence of SIGs I believe it will be important for SIGs to
decide if they would like to opt out of the RFC process entirely or if they would
like to modify the RFC process to better serve their needs. For example I can
imagine RFCs targeting SIG Architecture or SIG API Machinery would likely require
a section on backwards compatibility and migration strategies. I believe the best
place for a SIG to describe its development process would be in its charter.

### Prior Art

The RFC process as proposed was essentially stolen from the [Rust RFC process] which
itself seems to be very similar to the [Python PEP process][]

[Rust RFC process]: https://github.com/rust-lang/rfcs

## Graduation Criteria

Before we can consider this process successful for our current size I believe we
should hit at least the following milestones

- a release note draft can be generated by referring to RFC content
- a road map can be generated by referring to RFC content

The process for creating draft release notes and development status reports
should eventually rely largely on automation, but I do not believe that the
implementation of such automation should prevent the RFC process from being
considered stable. I can certainly imagine a future where a release notes draft
is produced nightly along with the collection of binaries that users deploy.

## Drawbacks

Any additional process has the potential to engender resentment within the
community. There is also a risk that the RFC process as designed will not
sufficiently address the scaling challenges we face today. PR review bandwidth is
already at a premium and we may find that the RFC process introduces an unreasonable
bottleneck on our development velocity.

It certainly can be argued that the lack of a dedicated issue/defect tracker
beyond GitHub issues contributes to our challenges in managing a project as large
as Kubernetes, however, given that other large organizations, including GitHub
itself, make effective use of GitHub issues perhaps the argument is overblown.

The centrality of Git and GitHub within the RFC process also may place too high
a barrier to potential contributors, however, given that both Git and GitHub are
required to contribute code changes to Kubernetes today I would argue that we
should focus on providing support to those unfamiliar with this tooling.

Expanding the proposal template beyond the single sentence description currently
required in the [features issue template][] may be a heavy burden for non native
English speakers and here I believe that the role of the RFC editor combined with
kindness and empathy will be crucial to making the process successful.

[features issue template]: https://github.com/kubernetes/features/blob/master/ISSUE_TEMPLATE.md

## Alternatives

This RFC process is related to
- the generation of our [architectural road map][]
- the fact that the [what constitutes a feature][] is still undefined
- how we [manage issues][]
- the difference between an [accepted design and a proposal][]
- [the organization of design proposals][]

and I believe that this RFC process attempts to place these concerns within a
more general framework

[architectural road map]: https://github.com/kubernetes/community/issues/952
[what constitutes a feature]: https://github.com/kubernetes/community/issues/531
[manage issues]: https://github.com/kubernetes/community/issues/580
[accepted design and a proposal]: https://github.com/kubernetes/community/issues/914
[the organization of design proposals]: https://github.com/kubernetes/community/issues/918

## Unresolved Questions

Whether we believe it makes sense for a SIG repository to serve a mini mono
repository where code lives in a structure like

```
kubernetes:sig-node/pkg/kublet
```

or whether SIGs maintain a list of repositories maintained by the SIG in an index
like

```
kubernetes:sig-node/repositories.md
```

is not clear to me. I believe that users would generally think about Kubernetes
from a component standpoint if they are aware of the component at all but I have
no real data to support that assertion.

## Mentors

- caleb miles
  - github: [calebamiles](https://github.com/calebamiles/)
  - slack: [calebamiles](https://coreos.slack.com/team/caleb.miles)
  - email: [caleb.miles@coreos.com](mailto:caleb.miles@coreos.com)
  - pronoun: "he"
