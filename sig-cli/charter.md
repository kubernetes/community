# SIG CLI Charter

## Purpose

The Command Line Interface SIG (SIG CLI) is responsible for kubectl and
related tools. This group focuses on command line tools and
libraries to interface with Kubernetes API's.

## Roles

* Chair
  * *MUST* run operations and processes governing the SIG.
  * *MUST* remain active in the role. A Chair will be automatically removed from
 the position if he/she is unresponsive for > 3 months.
  * *MAY* be removed if not proactively working with other Chairs to fulfill
 responsibilities.
  * *MAY* select additional Chairs through a [two-thirds super majority] vote
 amongst current Chairs. This *SHOULD* be supported by a majority of SIG Members.
  * *MAY* decide to step down at anytime and propose a replacement. Use
 [lazy-consensus] amongst Chairs with fallback on majority vote to accept
 proposal. This *SHOULD* be supported by a majority of SIG Members.
  * There *SHOULD* be 2 to 3 Chairs.
  * *MUST* be specifically enumerated in [sigs.yaml] under the **leadership -> chairs** label.

* Technical Lead
  * *MUST* provide the technical direction and vision for SIG CLI.
  * *MAY* establish new subprojects.
  * *MAY* decommission existing subprojects.
  * *MAY* resolve X-Subproject technical issues and decisions.
  * *MUST* be specifically enumerated in [sigs.yaml] under the **leadership -> leads** label.

* Subproject Owner
  * *MUST* be scoped to a subproject of SIG CLI defined in [sigs.yaml].
  * *MUST* be an escalation point for technical discussions and decisions in the
   subproject.
  * *MUST* set milestone priorities for subproject or delegate this responsibility.
  * *MUST* remain active in the role. A Subproject Owner will be automatically
 removed from the position if he/she is unresponsive for > 3 months.
  * *MAY* be removed if not proactively working with other Subproject Owners to
   fulfill responsibilities.
  * *MAY* decide to step down at anytime and propose a replacement.  Use
 [lazy-consensus] amongst Subproject Owners with fallback on majority vote to
 accept proposal. This *SHOULD* be supported by a majority of subproject
 contributors (those having some role in the subproject).
  * *MAY* select additional Subproject Owners through a
 [two-thirds super majority] vote amongst current Subproject Owners. This
  *SHOULD* be supported by a majority of subproject contributors (through
 [lazy-consensus] with fallback on voting).
  * There *SHOULD* be 3 to 5 Subproject Owners.
  * *MUST* be specifically enumerated in an OWNERS file under the **owners**
   label.  [sigs.yaml] *MUST* point to this OWNERS file under the **subproject ->
   owners** label.

* Approver
  * *MUST* contribute to SIG CLI by reviewing and approving proposed changes to
   the SIG CLI code for a directory and its sub-directories.
  * New Approvers *SHOULD* be chosen from current Reviewers by a majority vote of
   Subproject Owners for directories under their control.
  * *MUST* be enumerated in an OWNERS file within a directory under the
   **approvers** label.

* Reviewer
  * *MUST* contribute to SIG CLI by reviewing proposed changes to the SIG CLI
 code for a directory and its sub-directories.
  * *SHOULD* have made significant technical contributions during the review of at
   least six PR's in the last six months.
  * *MAY* be voted as Reviewer by a majority of current Subproject Owners for directories
   and subdirectories under their control.
  * *MUST* be enumerated in an OWNERS file within a directory under the
   **reviewers** label.

* Test Health Maintainer
  * *MUST* contribute to SIG CLI by maintaining test health.
  * *MUST* have successfully completed at least one test on-call rotation within the
   last six months as shown in on-call schedule of [Test Playbook].

* Member
  * *MUST* show sustained contributions to at least one subproject or to the SIG.
  * *MUST* maintain health of at least one subproject or the health of the SIG.
  * *MAY* build new functionality for subprojects.
  * *MAY* participate in decision making for the subprojects they hold roles in.
  * The Chair(s), Technical Lead(s), and Subproject Owner(s) are Members.
  * Approvers are Members.
  * Reviewers are Members.
  * Test Health Maintainers are Members.
  * A significant SIG contributor *MAY* be chosen as a Member using
   [lazy-consensus] amongst Subproject Owners with fallback to a majority vote
   of Members.

## Subprojects

* [kubectl](https://github.com/kubernetes/community/blob/master/sigs.yaml#L555)

### Subproject Creation

* Subprojects *MUST* be created by [KEP] proposal and accepted by [lazy-consensus]
of Technical Leads with fallback on majority vote of Technical Leads.
The result *SHOULD* be supported by the majority of SIG Members.
  * Subproject creation KEP *MUST* enumerate the founding Subproject Owners.
  * [sigs.yaml] *MUST* be updated to include subproject information and
   Subproject Owners.

## Organizational management

* SIG CLI meetings
  * *SHOULD* be facilitated by Chairs unless delegated to specific Members.
  * *SHOULD* be held bi-weekly on [Zoom](https://zoom.us/my/sigcli) video 
 conference every other Wednesday at 9am PST. Convert to your
 [timezone](https://www.timeanddate.com/worldclock/converter.html?iso=20180509T160000&p1=900).
  * Proposed topics and meeting agenda *SHOULD* be stored at the [Meeting Notes]
  * Meetings *SHOULD* be recorded and uploaded to the [Meeting Archive]
  * Meetings *SHOULD* contain a test health update from the on-call Member.

* SIG CLI overview and deep-dive sessions organized for Kubecon
  * *SHOULD* be organized by Chairs unless delegated to specific Members

## Technical processes

Subprojects of the SIG *MUST* use the following processes unless explicitly
following alternatives they have defined.

* Proposing and making decisions
  * Significant design proposals *SHOULD* be sent as [KEP] PR's and published to
   [SIG CLI email group] as an announcement.
  * Subproject owners are final decision-makers where [lazy-consensus] is not
   achieved for subproject technical decisions.
* Issues impacting multiple subprojects in the SIG *SHOULD* be resolved by Technical Leads
* Contributors *SHOULD* follow the SIG CLI [Contributing Guidelines]
* Members *SHOULD* follow the SIG CLI [Release Process]
* Test health
  * [Test Playbook] *SHOULD* describe the process for monitoring and managing
   SIG CLI code health.
  * PRs that break tests *SHOULD* be rolled back if not fixed within 24 hours
   (business hours).
  * Test health *SHOULD* be reported at start of each SIG meeting.
  * Test on-call rotation *SHOULD* transfer directly after bi-weekly SIG meeting.
  * Consistently broken tests *SHOULD* automatically send an alert to [SIG CLI email group]

## TODO

* Update [sigs.yaml] to current Chair(s), Technical Lead(s), and Subproject Owner(s).
* Update [OWNERS] and [OWNERS_ALIASES] to current subproject owners, approvers and reviewers.
* Create alert to send email to [SIG CLI email group] when a test continuously fails.
* Update [Contributing Guidelines]
* Update [Release Process]

*MAY*, *SHOULD*, *MUST* within this document are defined at [RFC 2119](https://tools.ietf.org/html/rfc2119)

[lazy-consensus]: http://communitymgt.wikia.com/wiki/Lazy_consensus
[two-thirds super majority]: https://en.wikipedia.org/wiki/Supermajority#Two-thirds_vote
[KEP]: https://github.com/kubernetes/community/blob/master/keps/0000-kep-template.md
[sigs.yaml]: https://github.com/kubernetes/community/blob/master/sigs.yaml#L502
[OWNERS]: https://github.com/kubernetes/kubernetes/blob/master/pkg/kubectl/OWNERS
[OWNERS_ALIASES]: https://github.com/kubernetes/kubernetes/blob/master/OWNERS_ALIASES
[SIG CLI email group]: https://groups.google.com/forum/#!forum/kubernetes-sig-cli
[Meeting Notes]: https://docs.google.com/document/d/1r0YElcXt6G5mOWxwZiXgGu_X6he3F--wKwg-9UBc29I/edit?usp=sharing
[Meeting Archive]: https://www.youtube.com/playlist?list=PL69nYSiGNLP28HaTzSlFe6RJVxpFmbUvF
[Test Playbook]: https://docs.google.com/document/d/1Z3teqtOLvjAtE-eo0G9tjyZbgNc6bMhYGZmOx76v6oM
[Contributing Guidelines]: https://github.com/kubernetes/kubectl/blob/master/CONTRIBUTING.md
[Release Process]: https://github.com/kubernetes/kubectl/blob/master/RELEASE.md
