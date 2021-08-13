# 2020 VOTERS GUIDE - KUBERNETES STEERING COMMITTEE ELECTION

## Purpose

The role of this election is to fill out the three (3) seats due for
reelection this year on the [Kubernetes Steering Committee]. Each elected
member will serve a two (2) year term.

## Background

This election will shape the future of Kubernetes as a community and project.
While SIGs and WGs help shape the technical direction of the project, the
[Steering Committee Charter] covers the health of the project and community
as a whole. Some direct responsibilities of steering members to consider as you are deciding
whether to run or who to vote for:

* Through the chartering review process, delegate ownership of, responsibility
  for and authority over areas of the project to specific entities
* Define, evolve, and defend the non-technical vision / mission and the values
  of the project
* Charter and refine policy for defining new community groups and establish transparency and accountability policies for such groups
* Define and evolve project and group governance
  structures and policies
* Act as a final non-technical escalation point for any Kubernetes repository
* Request funds and other support from the CNCF (e.g. marketing, press, etc.)
* Define and enforce requirements for community groups to be in good standing
  such as having an approved charter

For more context, please see the [current steering committee backlog] or a
previous [governance meeting video] which led to this whole process.

## Eligibility

Please refer to the [Steering Committee Election Charter] for [Eligibility for candidacy]

Eligibility for voting in 2020 is defined as:

* People who had at least 50 contributions to the Kubernetes project over
  the past year, according to a snapshot taken 2020-08-01 of the data driving the [devstats developer activity counts dashboard][devstats-dashboard], who are also [Org Members].
  Contributions include GitHub events like creating issues, creating pr's,
  reviewing PR's, commenting on issues, etc. For full details see
  [the SQL query used by devstats for developer activity counts][devstats-sql].

* People who have submitted the [voter exemption form] and are accepted by
  the election committee. We *explicitly* believe the above heuristic will be
  inaccurate and not represent the entire community. Thus we provide the form
  for those who have contributed to the project but may not meet the above
  criteria.  Acceptance of a form submission will be defined by a simple
  majority vote, and the criteria used during this process will be used to
  help refine further elections.

The requirement to be an [Org Member] is new this year, added by the Steering Committee to ensure that voters are following community issues.  If you otherwise qualify to vote but have not yet applied for Org Membership, then please [request an exception][voter exemption form] (and please apply for Org Membership as well).

Corporate affiliation is applied after the election. If an organization finds itself with too many representatives it is up to those individuals to come to a consensus on who should serve on the committee.

### Schedule

| Date         | Event                    |
| ------------ | ------------------------ |
| August 12    | Announcement of Election and publication of Voters.md |
| August 31    | Steering Committee Meeting with Q+A with the candidates and community |
| September 08 | All candidate bios and voting exception forms due by 0000 UTC (5pm PST) |
| ~1 week      | Election prep week (voters.md validation and CIVS setup and testing)
| September 14 | Election Begins via email ballots |
| October 03   | Deadline to request a replacement ballot |
| October 06   | Election Closes by 0000 UTC (5pm PST) |
| October 12   | Announcement of Results at Public Steering Committee meeting |

## Candidacy Process

**Nomination**

If you want to stand for election, send an email to kubernetes-dev@googlegroups.com
with the subject line "Steering Committee Nomination: Your Name (@yourgithub)".

If you want to nominate someone else, you may do so, but PLEASE talk to them
first.

If you wish to accept a nomination from someone else, reply to the nomination
email saying something like "I accept the nomination".

**Endorsement**

Once nominated, you must get the endorsement of three (3) different eligible
voters from three (3) different employers.  If you are eligible to vote
yourself, you count as one of the three. Endorsements from non-voting members does not count towards the final count.

[Eligible voters] may endorse candidates of their choosing by replying to the
candidate's nomination email saying something like "I endorse this nominee, and I work for <COMPANY>"
or "+1". Please specify your github ID, state that you are in voters.md, and include your employer's name so that we see can which candidates have sufficient endorsements.

When a candidate has reached the necessary three endorsements, one of the Election Officers will announce that on the email thread.  After that, please do not endorse the candidate further.

**Running**

Eligible candidates can submit a pull request with a biography in this
directory with their platform and intent to run. This statement is
**limited to 300 words** and must follow the format of `firstnamelastname.md`.
Please refer to the [2019 candidate bios] for examples. Biography statements are optional.

Missed deadlines by the candidates will be addressed by steering on a per case basis to determine eligibility.

**Campaigning**

Please refer to the [Steering Committee Election Charter] and understand
that we care deeply about [limiting corporate campaigning]. The election
officers and members of the steering committee [pledge to recuse] themselves
from any form of electioneering.

You should be running as a "brand free" individual, based on your contribution
to the project as a member of this community, outside of whatever corporate
roles you may hold.

## Voting Process

Kubernetes members in [voters.md] will receive a ballot via email. If you are
not on that list and feel you have worked on Kubernetes in a way that is NOT
reflected in GitHub contributions, you can use the [voter exemption form] to ask
to participate in the election.

Elections will be held using time-limited [Condorcet] ranking on [CIVS]
using the [IRV method]. The top vote getters will be elected to the open
seats.

Employer diversity is encouraged, and thus maximal representation will be
enforced as spelled out in the [Steering Committee Election Charter].

You will be ranking your choices of the candidates with an option for
"no opinion". In the event of a tie, a coin will be flipped.

The election will open for voting starting September 14th via email and
end three weeks after on October 6, 2020 at 00:00am UTC. You will receive an email
to the address on file at the start of the election from "Kubernetes (CIVS Poll
Supervisor) `<civs@cs.cornell.edu>`, please add to the list of addresses you don't spam filter. Detailed
voting instructions will be addressed in email and the CIVS polling page. Please
note that email ballots might be unreliable, so you are encouraged to contact
the election officials if you do not receive a ballot by September 17.

If you do not receive your ballot, request a new one via the [Ballot Replacement Form].

### Officers

The Steering Committee has selected the following people as [election officers]:
- Jaice Singer DuMars, @jdumars, Apple
- Ihor Dvoretskyi, @idvoretskyi, CNCF
- Josh Berkus, @jberkus, Red Hat

Please direct any questions via email to <election@k8s.io>.

### Decision

The newly elected body will be announced in the weekly [Kubernetes Community Meeting]
on October 12, 2020.

Following the meeting, the raw voting results and winners will be published on the
[Kubernetes Blog].

For more information, definitions, and/or detailed election process, please refer to
the [Steering Committee Election Charter]

## Nominees

|                    Name                    | Organization/Company |                        GitHub                          |
|:------------------------------------------:|:--------------------:|:------------------------------------------------------:|
| [Bob Killen](./bobkillen.md)               |      Google          | [@mrbobbytables](https://github.com/mrbobbytables)     |
| [Carlos Panato](./carlos-panato.md)        |      Independent     | [@cpanato](https://github.com/cpanato)                 |
| [Davanum Srinivas](./davanumsrinivas.md)   |      VMware          | [@dims](https://github.com/dims)                       |
| [Divya Mohan](./divya-mohan.md)            |      HSBC            | [@divya-mohan0209](https://github.com/divya-mohan0209) |
| [Ian Coldwater](./iancoldwater.md)         |      Salesforce      | [@IanColdwater](https://github.com/IanColdwater)       |
| [Jordan Liggitt](./jordanliggitt.md)       |      Google          | [@liggitt](https://github.com/liggitt)                 |
| [Lachlan Evenson](./lachlanevenson.md)     |      Microsoft       | [@lachie83](https://github.com/lachie83)               |
| [Stephen Augustus](./stephenaugustus.md)   |      VMware          | [@justaugustus](https://github.com/justaugustus)       |
| [Federico Bongiovanni](./fedebongio.md)    |      Google          | [@fedebongio](https://github.com/fedebongio)           |
| [Mayank Kumar](./mayankkumar.md)           |      Salesforce      | [@krmayankk](https://github.com/krmayankk)             |

[Kubernetes Steering Committee]: https://github.com/kubernetes/steering
[Steering Committee Charter]: https://github.com/kubernetes/steering/blob/master/charter.md
[current steering committee backlog]: https://github.com/kubernetes/steering/projects/1
[governance meeting video]: https://www.youtube.com/watch?v=ltRKXLl0RaE&list=PL69nYSiGNLP1pkHsbPjzAewvMgGUpkCnJ&index=23

[Steering Committee Election Charter]: https://git.k8s.io/steering/elections.md
[Eligibility for voting]: https://github.com/kubernetes/steering/blob/master/elections.md#eligibility-for-voting
[Eligibility for candidacy]: https://github.com/kubernetes/steering/blob/master/elections.md#eligibility-for-candidacy
[limiting corporate campaigning]: https://github.com/kubernetes/steering/blob/master/elections.md#limiting-corporate-campaigning
[pledge to recuse]: https://github.com/kubernetes/steering/blob/master/elections.md#steering-committee-and-election-officer-recusal

[Condorcet]: https://en.wikipedia.org/wiki/Condorcet_method
[CIVS]: http://civs.cs.cornell.edu/
[IRV method]: https://www.daneckam.com/?p=374

[2019 candidate bios]: https://github.com/kubernetes/community/tree/master/events/elections/2019
[election officers]: https://github.com/kubernetes/community/tree/master/events/elections#election-officers
[Kubernetes Community Meeting]: https://github.com/kubernetes/community/blob/master/events/community-meeting.md
[Kubernetes Blog]: https://kubernetes.io/blog/
[eligible voters]: https://github.com/kubernetes/community/blob/master/events/elections/2020/voters.md
[voter exemption form]: https://www.surveymonkey.com/r/k8s-sc-election-2020
[voters.md]: ./voters.md

[devstats-sql]: https://github.com/cncf/devstats/blob/master/metrics/shared/project_developer_stats.sql
[devstats-dashboard]: https://k8s.devstats.cncf.io/d/13/developer-activity-counts-by-repository-group?orgId=1&var-period_name=Last%20year&var-metric=contributions&var-repogroup_name=All
[Org Member]: https://github.com/kubernetes/community/blob/master/community-membership.md
[Ballot Replacement Form]: https://www.surveymonkey.com/r/kubernetes-sc-2020-ballot
