# SIG Governance

This doc uses [rfc2119](https://www.ietf.org/rfc/rfc2119.txt) to indicate keyword requirement levels.
Sub elements of a list inherit the requirements of the parent by default unless overridden.

The following document outlines SIG governance.  Its goals are to define the necessary structure for
the SIG to self govern and to define leadership responsibilities within the SIG.

This charter may be updated the members of the SIG through lazy consensus and fallback on voting with
a [super-majority].

## Roles and organization / project management

This section describes the various roles governing the SIG.

---

### Sig leads

Canonical list of sig leads: <dir>/OWNERS (sigLeads section).  This list is mirrored to
kubernetes-sig-~~foo~~-sig-leads@googlegroups.com.

Sig leads:

- *MUST* manage [operational concerns][business-operations] in a SIG.
- *MAY* retain the title "~~Foo~~ sig lead emeritus" when stepping down.

**Note:** Sig leads *MAY* ask for help to assist with specific responsibilities.

#### Responsibilities towards members the SIG

- *MUST* schedule and facilitate bi-weekly SIG meetings
  - ensure the agenda is set and mailed out ahead of time
  - ensure notes are taken and the recording is uploaded to some public location
  - manage time and keep discussion focused
- *MUST* facilitate the planning of quarterly / yearly goals
  - create a document for members to discuss goals
  - ensure each committed item has a tracking issue or KEP with an owner
- *MUST* act as role model within the SIG
- *SHOULD* facilitate communication in channels used by the SIG - slack, email, issues
  - respond to questions about the SIG itself - e.g. how do I join?
- *SHOULD* facilitating the staffing for areas critical to the long term success of the SIG including
- *MAY* setup new SIG contributors in the on boarding process
  - ensure there are issues available for new contributors to take on
  - ensure questions from new contributors are responded to

#### Responsibilities towards the project

- *MUST* act as an ambassador for the SIG
- *MUST* provide quarterly updates to the community at the community meeting
- *MUST* public quarterly / yearly goals to the community
- *MUST* act as point of contact for issues impacting the Kubernetes release
  - work with tech and project leads to make sure any issues targeted to the release milestone are completed and closed
  - work with tech leads to make sure tests are healthy before cutting a release
- *SHOULD* setup face 2 face meetings either ad-hoc or at conferences (e.g. kubecon *sessions*)
- *SHOULD* facilitate Contributor Summit topics and sessions
- *MAY* setup programs for the SIG (such as Outreachy)

#### Selecting sig leads

SIG maintains 2-3 sig leads.  When a sig lead steps down from their position, they *MAY* propose a new sig lead
(must be a SIG member) to take their place.  The SIG *SHOULD* agree upon new SIG leads through lazy consensus
with a fallback on voting with a [super-majority].

The number of sig leads *MAY* be expanded if proposed by SIG members and passed by lazy consensus with a fallback on a
[2/3 super majority][super-majority] vote.

Sig leads that have not been active for a period of 8 weeks or more (exceptions exist) *SHOULD* step down and propose a
replacement.  For planned absences, sig leads *SHOULD* find a temporary replacement to fill their role and delegate
responsibility.

The sig lead group *SHOULD* represent a diverse set of organizations.  This is predicated on interest in the role
from multiple organizations with qualified candidates.

---

### Technical Lead (TL) role

Canonical list of tech leads: <dir>/OWNERS (techLeads section).  This list is mirrored to
kubernetes-sig-~~foo~~-techleads@googlegroups.com.

This list may be used to escalate issues to technical leads.

Technical leads:

- *MUST* either be a Project Lead *or* be well versed in areas directly owned by the SIG (not yet in subprojects)
- *MUST* act as de facto Project Lead for all areas of the SIG *not* yet owned by any subprojects
- *SHOULD* have a record of consistent solid technical judgement and leadership over a sustained period
  - especially over areas for which they are the de facto Project Lead
- *SHOULD* be an owner of one or more SIG subprojects
- *MAY* also be a SIG lead
- *MAY* retain the title "~~Foo~~ technical lead emeritus" when stepping down.

#### Responsibilities towards members the SIG

- *MUST* define scope and structure
  - define what areas fall within the scope of the SIG (in conjunction with SIG architecture)
  - define how areas within the scope of the SIG are divided into subprojects
- *MUST* act as role model within the SIG
- *SHOULD* act as an uber-escalation point for technical disputes in subprojects.

#### Relationship

Technical leads *SHOULD* meet or communicate regularly enough to ensure they aligned.  Decisions made by
technical leads should be consistent across the group.  If technical leads are providing inconsistent direction,
then they *MUST* align themselves.

#### Selecting technical leads

The SIG maintains a group of 2-7 technical leads.

New technical leads *SHOULD* be nominated either by existing technical leads or self nominated by sending an email to
the tech lead mailing list.  Technical leads membership proposals *SHOULD* achieve [lazy consensus][lazy-consensus]
with SIG members on the proposal with a fallback to voting (majority).

Members *SHOULD* have shown sustained superior technical judgement and leadership within the SIG before being considered
for a technical lead position.  They should already be consistently fulfilling the responsibilities outlined above
for 3 months.

Before becoming a technical lead, members *MAY* reverse shadow an existing technical lead until they are familiar
with the role.  Reverse shadowing means that they take on the responsibilities of tech lead, but their work is subject
to review.

When possible, it is preferred that the Technical Lead group represent a diverse set of organizations.  This is
predicated on interest in the role from multiple organizations with qualified candidates.  At times the *most*
qualified candidates may all be from a single organization, however interested candidates that meet the minimum
qualifications exist.  In this event, the technical leads may be temporarily comprised of a single organization,
but must mentor qualified candidates from other organizations with the intention of giving them a technical lead role.

### Expiring technical leads

Tech leads that have not been active in an area for 3 months *SHOULD* not be considered for any decision making and
*SHOULD* consider stepping down.

Tech leads that have not been active in an area for 6 months *SHOULD* be automatically removed as tech lead.

If the number of tech leads falls bellow the minimum, or the majority of tech leads deem it necessary, then new tech
leads *MAY* be seeded from one or more of:

- Project leads
- Top-level OWNERs approvers (for areas owned by the SIG)
- A proposal with [lazy consensus][lazy-consensus] within the SIG.
- SIG architecture picking technical leads for the SIG

---

### Project Lead (PL) role

Canonical list of project leads: <dir>/OWNERS (ProjectLeads section).  This list is mirrored to
kubernetes-sig-~~foo~~-projectleads@googlegroups.com

Project leads:

- *MUST* be well versed in areas owned by the subproject
- *SHOULD* have a record of consistent solid technical judgement and leadership over a sustained period in
  the areas owned by the subproject
- *MAY* also be a SIG lead
- *MAY* retain the title "~~Foo~~ project lead emeritus" when stepping down.

#### Responsibilities towards members the SIG

- *MUST* provide technical direction for the project area
 - resolve technical disputes
 - set priorities
- *MUST* define milestones for the project area
  - define releases and priorities for release
- *SHOULD* ensure technical excellence within area
  - actively work to reduce complexity of the area
  - define general acceptance criteria for work
    - code guidelines and requirements
    - test guidelines and requirements
    - documentation guidelines and requirements
  - make technical decisions with other project leads and work cross organizationally with other subprojecs and SIGs
  - make technical decisions unilaterally if needed
    - e.g. addressing critical bugs or security vulnerabilities
- *SHOULD* participate in decision making - either voicing support or opposition of proposals
  - weigh in on discussions, PRs, issues for the owned areas
  - attend SIG *or* subproject meetings
- *MAY* onboard new project leads and subproject contributors
  - delegate responsibility and ensure opportunities for members to grow
- *MAY* define how artifacts published by the subproject are consumed external to the SIG
  - how are the areas mapped to binaries, libraries, services, etc

#### Relationship

Project leads *SHOULD* meet or communicate regularly amongst each other and project leads from other subprojects
to ensure they aligned.

#### Selecting project leads

Project lead selection *SHOULD* follow the same process as SIG Technical Lead selection, but scoped to the
areas owned by the subproject.

---

### Member role

Canonical list of sig leads: <dir>/OWNERS (Members section).  This list is mirrored to:
kubernetes-sig-~~foo~~-members@googlegroups.com

### Responsibilities towards members the SIG

- Acting as role model to other members within the SIG
  - Maintaining a respectful and welcoming culture is the responsibility
    of all members of the SIG.  Treat other members with kindness.
  - Foster an inclusive culture
- Ensuring continued area health through proactively
  - Triaging issues
  - Reviewing PRs
  - Participation in monitoring of test health
  - Fixing tests
  - Contributing features
  - Refactoring and simplifying codebase
  - Participating in design discussions and providing feedback
- Taking technical ownership of areas governed by the SIG
  - Making technical decisions with other members
  - Driving technical direction of the SIG
  - Required: Ensuring health of code authored
  - Optional: Ensuring health of code authored by other members
  - Health includes
    - contributions adhere to criteria defined by style and contribution
      guides
    - code correctness is tested through automation
    - tests remain passing
    - issues are resolved and patched if needed
    - code remains simple and approachable to others
      - self documenting code through function and variable naming
      - explicit documentation
- Responding to issues facing the SIG
  - Fixing critical bugs

### Externally facing responsibilities

- Acting as ambassadors to the SIG through
  - Optional: Responding to messages on communication channels
  - Optional: Participating in events such as meetups

#### Selecting members

Members *SHOULD* show a commitment to the SIG before becoming members. This maybe a verbal commitment - e.g.
I plan to work in this area full time for the next 3 months - or something demonstrated over time through
contributions in PRs, issue triage, test triage, operations, etc.

Members are expected to be responsive to issues for sub areas they own within the SIG.  Members are expected to
help other members in the SIG (e.g. transfer knowledge, help debug, assist with operational tasks, answer questions,
etc).

Membership is provisional until the member has shown sustained contributions for 3 months.  Provisional members may
provide feedback in decision making, but are not deciders in the decision itself.

---

## Processes and tools

The following section outlines various processes used by the SIG.

### Decision making within the SIG

It is critical to a SIG that the decision making processes is healthy. A good decision making process has the following
properties:

- Decision making process is unambiguously understood by everyone involved
  - Clear definition of the set of individuals responsible for making
    decisions
  - Clear definition of the decision making process used by those
    individuals
- Decisions are resolved instead of languishing and defaulting to no
  decision
- Decisions are a commitment and won't waffle after having been made

#### Who, What, Where, When, How

**Note:** Proposals *SHOULD* be sent before or early in the release cycle they are targeted to.  The SIG or
subproject *MAY* decide to defer discussing or accepting proposals during stabilization periods.

In cases where an author is seeking agreement on a proposal, there *MUST* be a clearly defined process for
communicating and resolving the proposal.

When using [lazy consensus][lazy-consensus], all proposals *MUST* be published to a visible well known location.  For
acceptance to be given be implicitly (as it is in lazy consensus models), all involved parties *MUST* be made aware of
the proposal.  Involved parties *SHOULD* watch for new proposals and read them in a timely manner, the SIG governance
*MUST* ensure that where and how proposals are discussed is well defined.

Authors of proposals *SHOULD* generally allow at least 5 business days before considering a proposal accepted.  Large
or complex proposals *SHOULD* allow more time than normal and *SHOULD* ensure that a quorum of area OWNERs (such
as project leads) are present at the time the proposal is made (e.g. not on vacation).  Area owners may defer
discussion of a proposal if for some reason it is not a good time.

In order to avoid [Warnock's Dilemma][warnocks-dilemma], at least one active member *SHOULD* respond to a proposal
before it is accepted.

> **Note:** To those forking this template - consider changing the
  following text to only include the tools actively used by your SIG.

Proposals *MUST* be made using a commonly used tool (issues, PRs, docs, etc) and then *MUST*
be published through a well defined mailing list with the subject containing either `KEP: ` or `RFC: `.
This provides a simple mechanism to subscribe to proposals that is less noisy than subscribing to github notifications.

Discussion *MAY* occur on the issue / PR / doc itself.  Any comments on the doc by SIG members that are requesting a
change *SHOULD* be considered blocking until they are responded to at least once.  In cases, that response maybe some
variation of "No".  This is to disambiguate between the comment not being read, and comment not being accepted.

Once the comments are responded to, they *MAY* be considered addressed unless either there is another response or the
comment has some language to indicate it is a blocking issue that must be explicitly resolved.
e.g. *lets hold off on doing this until...*

**Optional:**

Once all comments have been addressed on the proposal and there have been no additional responses in 5 business days
(or the lazy consensus period), a final email *MAY* be sent to the mailing list with `LastCall: KEP: ` or
`LastCall: RFC: `.  This email is sent to ensure that everyone involved in the decision making process understands
that the proposal will be merged as accepted if there is no more additional feedback.  If there are no further comments
or responses to the mail in 1-2 business days the proposal *MAY* be considered accepted.

**Important:** Once a proposal has been accepted then the decision has been made and a commitment has been given to
whoever implements the proposal that their contribution will be accepted if it is well implemented.  If it is deemed
necessary to reverse a decision, this should be done as a new proposal and require lazy consensus on the reversal.

In cases where consensus cannot be achieved in 1-2 weeks, the proposal may be escalated to the project leads.

#### Defaulting toward action

In cases where there is consensus that a problem exists, but a resolution cannot be achieved through consensus or
escalation to technical leaders, then the default resolution should be that the individuals staffing the effort make
the decision.  This is to avoid deadlock and stagnation when addressing known issues with multiple valid solutions.

### Test health ownership and triaging issues

#### Test dashboard

Subprojects *MUST* maintain dashboards on [testgrid](https://k8s-testgrid.appspot.com/) that display
tests to certify their areas are healthy and in a releasable state.

#### Template Option 1: Bi-weekly test review

In order to facilitate consistently passing tests and triage of issues, the test dashboards *MUST* be reviewed prior
to each SIG meeting.  Broken or flaky tests *SHOULD* be added as the first item to the SIG meeting agenda and
assigned to an owner.

#### Template Option 1: Test [SLO][slo]

The SIG *MUST* fix any broken test within the 2 weeks between SIG meetings.  Resolution *MAY*
include (as deemed most correct by the owner): rolling back PRs, rolling forward PR fixes, or disabling the failing
tests.

#### Template Option 2: Feature branches with bi-weekly test review

In order to facilitate consistently passing tests and triage of issues, all development *MUST* occur in feature branches
and *MUST* only be merged into master after all tests have been successfully run on the feature branch.
The test dashboard *MUST* be reviewed prior to each SIG meeting for flaky tests or infrastructure / environment
failures.  Issues *SHOULD* be added as the first item to the SIG meeting agenda and assigned to an owner.

#### Template Option 2: Test [SLO][slo]

The SIG *MUST* fix any broken test within the 2 weeks between SIG meetings.  Resolution *MAY* include
(as deemed most correct by the owner): rolling back PRs, rolling forward PR fixes, or disabling the failing tests.

#### Template Option 3: Build cop rotation

In order to facilitate consistently passing tests and triage of issues, an oncall rotation *MAY* be setup containing
SIG members.  The purpose of this group is to proactively monitor signals for the health of the areas owned by
the SIG, and to provide an escalation point for urgent or critical issues.  If no oncall group is chosen, this group
*MAY* default to the technical and project leads groups.

**Note:** Oncall calendar infrastructure for the community is forthcoming.

#### Template Option 3: Test [SLO][slo]

The SIG *MUST* fix any broken test within 4 days either by addressing the underlying issue, or rolling back the
breaking change.  The oncall member *MAY* freeze code merges for the SIG areas if tests remain broken
for more than 4 days.

#### Template Option 3: Test escalation

If assets owned by the SIG are suspected to be breaking other SIGs tests or making their tests flaky, then
members of those SIGs *MAY* send an email to the sig mail list and escalate to the test oncall member.  If the issue
cannot be resolved, the issue *MAY* be escalated to technical leads.

#### Flaky tests

Flaky tests are defined at having >10% failure rate by job. Tests observed to be consistently flaky *SHOULD*
be prioritized and fixed within 1 release cycle.  If no members of the SIG can commit to fixing the flaky test
within a release single, then Tech leads *MAY* stop accepting PRs for new features until staffing is committed.

### Addressing operational or technical blocking issues

In rare cases it can be necessary to address critical technical issues within the SIG or a subproject.  If these
issues cannot be resolved through the normal processes, then the SIG or technical leads *MAY* stop accepting
contributions for new features to the SIG or subproject until the issue has been resolved.

If still no resolution can be found, the issue *MAY* be escalated to the [Steering Committee][steering-commitee].

[lazy-consensus]: http://communitymgt.wikia.com/wiki/Lazy_consensus
[super-majority]: https://en.wikipedia.org/wiki/Supermajority#Two-thirds_vote
[warnocks-dilemma]: http://communitymgt.wikia.com/wiki/Warnock%27s_Dilemma
[slo]: https://en.wikipedia.org/wiki/Service_level_objective
[steering-commitee]: https://github.com/kubernetes/steering#contact
[business-operations]: http://www.businessdictionary.com/definition/business-operation.html