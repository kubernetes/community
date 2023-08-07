# Steering Election HOWTO

This documentation explains how to run a steering election.  As a work in
progress, it replaces older documentation where present; for topics not
covered consult the older documents in each election folder.

## Documentation contents

* This guide, which covers all process and roles.
* A folder of templates, primarily for election-related communications
* Current SQL query or devstats report for selecting voters

## Roles

### The Elections Subproject

Members of the [Elections Subproject](elections/README.md) have three
responsibilities around steering elections:

1. Recommending Election Officers to the Steering Committee
2. Intervening if anything unexpected happens with the EOs
3. Following up on recommendations from the election retro

### The Election Officers

The Election Officers (EO) are three trusted contributors who will run one or two
Steering elections.  They are responsible for making sure that the election
happens and completes satisfactorily.

Election Officers must meet the following requirements:

* Org member for more than one year
* Eligible to vote in the election
* Pledge to administer the election without partiality to any employer, SIG,
  or personal preference
* Tentatively available for two elections in a row (see below)
* Tentatively available for special elections in the upcoming year (see below)

Additionally, the Elections Subproject will choose Election Officers partly to
respresent the diversity of Kubernetes contributors, selecting different
genders, geographic regions, and ethnicities where possible, in order to
avoid the appearance of bias in the elections process.  Particularly, the
three Officers must each work for a different employer.

Election Officers are responsible for:

* Planning the election, including creating a draft timeline
* Generating the voter list
* Setting up the election in the voting system
* Deciding on exception requests
* Determining candidate eligibility
* Assisting candidates with bios
* Publicizing the election in order to maximize participation
* Finalizing and reporting the election results
* Hosting an election retro & documenting changes
* Contributing to the SC election documentation

These responsibilities are detailed in the election procedures.  They collectively
fall in to four main areas, timewise:

1. Communications with candidates and voters (together with the Comms Liaison)
2. Administering the election software (together with the Infra Liaison)
3. Managing nominations and candidates
4. Responding to exception requests

Usually, the three Election Officers divide up the above major responsibilities
among the team, each Officer taking one or two of them.

Each year, at least one Election Officer must be a prior Officer in order to
ensure continuity of knowledge. As such, Officers should be theoretically
available to do the elections two years in a row.  Further, should a special
election become necessary because of the resignation of an SC member mid-year,
this year's Officers are responsible for running the special election, so
any Officers should be at least tentatively available for that.

### Alternate Election Officer

In addition to the three Officers, the Elections Subproject will recommend one
or two Alternate(s).  This Alternate is available in case one of the Officers is
unable to complete the election, or is unavailable for post-election duties (such
as a special election).  Any current Officer may activate the Alternate if
any Officer resigns or is unavailable.

The Alternate will be added to the Election Officers slack, but will have no
duties and will not participate in votes unless activated.

### K8s Infra Liaison

Currently election software runs on the Kubernetes cluster owned by the k8s-infra
team.  As such, the Election Officers may need troubleshooting and support from
k8s-infra in case of unexpected problems with the Elekto deployment or changes
needed to the software.

As such, before the election starts, the Officers should reach out to k8s-infra
team and request one person who will be available to assist.  This person should
be available during most of the election period.  They must have the ability
to approve/modify services running on k8s.io.

If one of the Officers has these permissions, they may also serve in this role.

### Contributor-Comms Liaison

A big part of the election effort is making sure that voters and candidates
are aware of the election and kept up to date with constant reminders.  As such,
the Officers work directly with Contributor-Comms to send out a stream of
messages to the community.  

Well before the election starts, the Officers should reach out to contributor-comms
and ask them to assign one team member to handle election communications.  
This Comms member needs to have the ability to approve tweets.

If one of the Officers is a member of Contributor-Comms, they may double up
in this role.

## Timeline and Processes

What follows is a timeline for conducting the steering committee election.  Since
the actual calendar varies a bit, time below will be expressed starting at "day 0,"
which is usually sometime between June 1 and June 30.  Figure out your calendar
offsets based on that, and copy this grid to document the current election.

Most of the days below have quite a bit of adjustment room; there's just a few
where things must happen in a specific number of days.

| Day   | Activity                    | Roles              | Notes                                 |
| ----- | --------------------------- | ------------------ | ------------------------------------- |
| 0     | Begin EO Selection          | Subproject Leads   |                                       |
| 10    | Nominate EO members         | Subproject Leads   |                                       |
| 14    | Approve EO members          | Steering Committee |                                       |
| 20    | Select liaisons             | EO Members         |                                       |
| 20    | Propose timeline            | EO Members         | Also, propose any eligibility changes |
| 24    | Update Elekto               | Infra Liaison      | Make sure that Elekto is up to date   |
| 25    | Approve Timeline            | Steering Committee |                                       |
| 26    | Request Voters              | EO Admin           | request list from Devstats            |
| 26    | Create Retro doc            | EO Members         | Open this early for notes             |
| 28    | Prepare Comms Plan          | EO Comms           |                                       |
| 30    | Create Election Config      | EO Admin           | Create the election in k/community    |
| 32    | Open nominations            | EO Members         | Announce start of nominations         |
| 32-51 | Validate Candidates         | EO Nominations     | Continue through period               |
| 32-78 | Evaluate Exception Requests | EO Members         | All the way through election          |
| 43    | 1st Nominations Reminder    | EO Comms           |                                       |
| 45    | Candidate/SC Q&A            | EO, SC, Candidates | Any time during nominations           |
| 48    | 2nd Nominations Reminder    | EO Comms           |                                       |
| 49    | Last Nominations Reminder   | EO Comms           |                                       |
| 50    | Nominations Close           | EO Members         |                                       |
| 51    | Clean Up Nominations        | EO Admin           | Fix/Edit Candidate statments          |
| 53    | Announce Voting             | EO Comms           | k/dev, LWKD, Slack                    |
| 55    | Voting Blog Post            | EO Comms           |                                       |
| 60    | First Voting Reminder       | EO Comms           |                                       |
| 67    | Second Voting Reminder      | EO Comms           |                                       |
| 74    | Third Voting Reminder       | EO Comms           | remind that exceptions close          |
| 78    | Close Exceptions            | Automatic          |                                       |
| 80    | Final Voting Reminder       | EO Comms           |                                       |
| 81    | Election Closes             | Automatic          |                                       |
| 82    | Send Results to SC          | EO Admin           |                                       |
| 84    | Approve Results             | Steering Committee |                                       |
| 86    | Inform Candidates           | EO Member          | MUST BE 24-48 hours before announcing |
| 87    | Announce Results            | EO Comms           | on K/dev, etc.                        |
| 88    | Election Blog Post          | EO Comms           |                                       |
| 89    | Results in Elekto           | EO Admin           |                                       |
| 94    | Upload ballots.csv          | EO Admin           | Also detailed results                 |
| 95    | Hold Retro                  | EO, Subproject     | Anytime within 2 weeks                |

### Election Officer Selection

At the beginning of the election cycle, ideally in Early June, the Elections Subproject Leads will [select Election Officer candidates].

Potential sources for qualified EO candidates include:

* Previous EO members
* Emeritus Steering Committee members
* Current Steering members who are terming out
* Current or former Contributor Experience leads

Subproject leads will choose the EOs to match the criteria for the role.  The Subproject is not required to issue a public call for candidates, nor document their EO selection process, unless asked to do so by the Steering Committee.

When they settle on a final list of three candidates and one alternate, they will check in with SIG-ContribEx leads, and then propose the list of officers to the current Steering Committee.  This proposal should happen as an issue or PR, such that all Steering members are able to see the proposal and comment on it. If Steering requires different candidates, the Subproject Leads will find alternate EO possibilities until Steering approves.

A Subproject lead will submit a PR to create the team, primarily creating an OWNERS file in the appropriate steering/YEAR directory.  They will request that the Slack Admins change the membership on the #election-officers Slack channel to reflect the current EOs.

### Set Up

#### Creating The Election Team

Once the three EOs and an alternate have been selected, they should find liaisons. This means requesting a Comms liaison from Contributor-Comms, and an Infra liaison from SIG-K8s-Infra.  It's possible that both of these roles exist among the EOs themselves, in which case Liaisons do not need to be selected.

Then the EOs should decide among themselves who will be taking on which duties.

Hereafter, we will refer to the various roles with these abbreviations:

* EO Comms: the EO member taking on comms duties plus the Comms Liaison
* EO Admin: the EO member taking on Elekto admin duties
* EO Nominations: the EO member taking on nomination/candidate wrangling
* EO Voters: the EO member in charge of voters.yaml and exception requests
* Infra: the Infra Liaison

#### Create a Timeline

The second thing the EOs need to do is to create a timeline with specific dates for each of the election activities.  While this will be based on the default timeline above (or in the template README), EOs will have to adjust this for the realities of the current year.  If the election cycle started late, they may need to shorten some periods. EOs should keep some limitations in mind for scheduling:

* Don't have any deadlines fall during, or immediately after, a Kubecon week
* The Nominations period must be at least 2 weeks
* The Voting period must be at least 2 weeks, and should be three
* 1-2 day gaps between some steps exist in case of problems (e.g. Elekto error, Steering unresponsive)

Once EOs have a draft timeline, they will submit it to Steering to approve, either as an issue in k/steering, or as a PR that creates the README for the election. Asking for review on Slack is recommended. Once a majority of Steering members have approved, the timeline is final.

Note that any consequential changes to the timeline (such as changing close dates or the results date) require re-approval by Steering.

At the same time you have Steering approve the timeline, they should also approve the voting eligibility criteria. As such, you should propose using the same criteria as last year; Steering will make any changes that they need to.

#### Requesting Voters

The eligible voters list is produced through this process:

1. Generate the list of eligible contributors from [Devstats], filtering by required number of contributions. Search for all repos for the last year, bots excluded, and [generate a CSV to download].
2. Match this list against the list of Org members supplied from the [Org Repo], discarding any contributor who is not an Org member. An example of how to produce the Org member list:
```
yq '.admins + .members'   config/kubernetes/org.yaml   config/kubernetes-client/org.yaml   config/kubernetes-csi/org.yaml   config/kubernetes-sigs/org.yaml | sort -f | uniq | grep -v "\-\-\-"
```

3. Add to the list all members of the Code of Conduct Committee and Security Response Committee, if not already present

4. Sort alphabetically, deduplicate if required, and reformat to match the format of voters.yaml template

Note that currently Elekto is case-sensitive and requires capitalization in voters.yaml to match the official capitalization of each contributor's GitHub account. While this may be fixed in the future, you need to be careful to preserve capitalization when copying GH IDs for now.

#### Creating The Election

For Steering elections, we are currently using [Elekto]. Elekto is a GitOps system.  Thus, creating the election is a matter of copying the template files over from templates/files, and filling in the details per the current election.  Here's some brief notes on the contents of each file:

* **README.md**: This is the human-readable complete set of election instructions.  It should contain general information about the election, a timeline, eligibility requirements, etc.  It is not displayed in Elekto.
* **election.yaml**: this is the configuration file that defines the election for Elekto.  All of the missing information here needs to be filled in. *This file will need to be renamed from `election-template.yaml` to `election.yaml`*
* **election_desc.md**: this is human-readable brief summary of the election, displayed to voters by Elekto.
* **voters.yaml**: this is the list of eligible voters used by Elekto.
* **nomination-template.md**: this is a sample candidate template for candidates to base their nomintation statements and profiles on. You should not need to edit it.

Note: also list the voters.yaml file in the spelling-check exclusion file [hack/.spelling_failures](https://github.com/kubernetes/community/commit/5fd6a0222fc88fdc35131a2bb5321ef17ad65d27). This prevents a failure in the build caused because usernames don't match what a dictionary lists as a correctly-spelled word.

Each of the above templates will have blanks to fill in.  When the EO Admin has filled in all information, they submit it for the other two EOs to approve. Ideally, a project member trained on Elekto will also review it for technical errors.

Once the file is merged, the election should appear on the [election website] in half an hour or less.  If it does not, contact the Infra liaison for troubleshooting.

#### Updating Links

The last set of PRs for setting up the election is to update links and other files.  That includes:

* Update [steering election README] to point to the current election
* Update the [elections doc] in the Steering repo with the current EOs
* Update the [elections alias] to direct to the current EOs

#### Create a Comms Plan and Messages

The timeline above has a lot of messages and posts that go out as election publicity and reminders.  To keep on top of this, the EOs should prepare a [Comms Plan] that includes the planned dates of all announcements, with draft text for each of these messages. This document allows election communications to be complete and on schedule, even if the Alternate needs to pick it up.

Templates for some of the expected messages that EO Comms sends out can be found in the [messaging folder], to be updated for the current year and sent out.  Where templates are not available, search the [k/dev] archives for prior years' examples, or prior years' comms plans.  And then please create a template and put it in the messaging folder.

Aside from specifically targeted messages, there's two levels of messages to be sent out:

- Major announcements: opening of nominations, start of voting, election results
- Reminders: many other reminders during the election cycle

Reminders are sent to k/dev and #announce and #kubernetes-contributors slack channels, and sometimes to social media at the discretion of EO Comms.

Major announcements are sent to those places, definitely to social media, and also broadcast via the slack bot to all SIG channels.  They should also be posted via LWKD, shared in the SIG Leads meeting, and posted as a short blog post in the contributor blog.

#### Steering Committee Supervision

Some things for the election need to be decided by Steering.  Approving the EOs, approving the timeline, and any voting or candidate eligibility adjustments, are decided by all of Steering. Steering decisions that need to be made during the election, such as approving the final results, are approved by "non-involved Steering members."  This consists of the group of Steering Committee members who are not running in the current election.

### Nominations Phase

Two to three weeks before voting starts, you should open the election to [nominations].  The actual nominations process is described in the [election README], so this guide will focus on what the EOs need to do.

First, EO Comms needs to announce the start of nominations, which is a major announcement.  This announcement also includes the call for Voter Exceptions.

#### Candidate Nominations

Once nominations have started, the EOs have basically three duties to candidates:

1. Reminding folks that they need to run via broadcast announcements
2. Helping candidates/nominators with nominations
3. Helping candidates with their profile PRs and debugging errors
4. Verifying candidate eligibility

The EOs are responsible to make sure that everyone who might be thinking about running for SC knows that an election has started, and is aware of the deadlines.  The way to do this is announcements in all the community places.  EOs should *not* do 1-on-1 private outreach to potential candidates, because this could be interpreted as partisanship.  EOs can and should reach out to already declared candidates privately to help them create and format their candidate statements.

A candidate declares themselves by filing an Issue in the [Community repo], and linking that issue with an announcement on k/dev.  Occasionally, someone other than the candidate will do that nomination, but mostly it's self-nominations.  

Importantly, voters are *not* to endorse the candidate on k/dev itself; all endorsements must be made on the issue.  This will be one thing EOs have to fix, probably more than once.  If a thread starts with people +1ing the candidate via email on k/dev, post an announcement saying something like

```
We remind you, as Election Officers, that your fellow contributors have asked you NOT to make endorsements on the mailing list.  Further, endorsements on the mailing list do not count towards candidate eligibility.  Instead, endorse this candidate here: <LINK to ISSUE>
```

In severe cases, EOs may need to ask the k/dev admins to lock the thread.

On the rare occasions that a voter nominates the candidate instead of themselves, an EO should reach out to the candidate to make sure they accept the nomination.

When a candidate creates an issue, EOs should check if they've copied in the process for endorsements; if not, an EO should.  Particularly, voters need to remember to provide their employer name.  Then the EO Nominations watches to see when the candidate accumulates the necessary three endorsements from contributors working for different employers.  At that point, EO Nominations posts to the issue:

```
<CANDIDATE NAME> has the necessary endorsements and is eligible to run in the Steering Committee election.  The candidate should prepare their candidate profile as a PR and submit it, per the instructions at: https://github.com/kubernetes/community/blob/master/elections/steering/{{YEAR}}/README.md#candidacy-process
```

In most cases, there are no other eligibility requirements.  Candidates are not required to be documented contributors or org members. The Issue will remain open until the candidate profile is merged as a PR.

#### Candidate Profiles

Candidate profile files are required for the candidate to be electable in Elekto.  They also give the community some idea who the candidate is and what they advocate.  Since these profiles are both a technical document (Elekto expects certain data and headers), and a platform document, some candidates will need help preparing them.  Here's the main areas where EOs help:

**File Names**: the correct name for each candidate profile is `candidate-ghid.md` where the second part is the GitHub ID of the candidate.  This ID must exactly match the ID field inside the file itself.

**Header Formatting**: the file is a markdown file with a YAML header.  Candidates often mess up the YAML header.  Things that will break it include: deleting or renaming fields, using the `@` symbol, removing the `--------` break for the header delimiter, or changing the indenting.  Check for, and correct, these things.

**Candidate Statements**: Some candidates may need editorial help on the contents of their candidate statements.  This can range from spelling, punctuation, and markdown correction, to help making their ideas and self-description clear.  There is also a loose limit of 300 words for the candidate statement, so EO Nominations will block merging any candidate statements that are excessively long until the candidate edits them.

Once candidate profiles are fully corrected, they can be merged to be ready for the election. There should be a 2-3 day grace period between the close of nominations and the start of voting to do any last-minute cleanup.

Candidates who miss the nomination deadline, or who are unresponsive to needed corrections past the beginning of voting, are ineligible to run.

#### Voter Exceptions

Devstats activity doesn't capture 100% of contributions, so through the whole voting cycle the EOs will be approving requests for Voting Exceptions.  These exceptions allow folks to vote who didn't show up in the original voters.yaml file.

Voters can check their eligibility in Elekto, which shows them whether it thinks they can vote.  They can also check the file. If a voter is not eligible but thinks they should be, they file an Exception via a link in Elekto.  These exception requests will display as a list in the election admin page, along with a toggle for whether or not the EOs have evaluated the exception request.

The EOs should stay on top of these exception requests, ideally telling the person within a few days whether they are eligible or not.  The EOs discuss and vote on exceptions.  Generally, the endorsement of two EOs is sufficient to approve an exception, although it's important that all three EOs can see the discussion.

Once a decision has been reached, the EO Voters will email the person with the decision.  Since Elekto does not maintain results of decisions (just that they've been evaluated), or email anyone, this needs to be handled separately by the EOs.

All else being equal, the EOs should err on the side of approving rather than rejecting exception requests.  Voter turnout is a challenge with every election, and exception requestors almost always vote. However, there are some precedents:

Eligible:

* Contributors with over 50 contributions who aren't org members, but file an exception.  In this case, urge them to become org members as well as granting it.
* Contributors whose org membership has been filed but is not yet approved
* People who have assisted extensively (as in, 15+ hours work) with the Kubernetes Contributor Summit
* Anyone who appears in SIGs.yaml for a role that they are currently fulfilling (not Emeritus, and not for deprecated subprojects)
* Members of a Release Team in the last year

Ineligible without other contributions:

* Writer/maintainers of 3rd party information resources (private/company blogs, personal/company Kubernetes websites, personal video channels)
* Meetup/User Group organizers
* Conference speakers
* Contributors to other CNCF projects
* Contributors to Kubernetes distributions

These are precedents, though, and the EOs will need to use their judgment on whether or not the individual's activity represents a significant contribution to the Kubernetes project in the last year.

People should be able to request voter exceptions until 3 days or so before the end of voting. Once voting has started, the EOs should notify requestors more quickly.

After the election, the EOs will report to Steering how many requests were received, and how many granted or denied.  EOs do not report on any specific exception request.  To avoid embarrassment, particularly for people whose requests were denied, the details of exception requests are treated as confidential.

#### Steering/Candidates Meeting

The project tries to arrange a private Q&A session where the candidates can meet current Steering members and ask about the position.  Given time zones, this often means two meetings.  These sessions should be near the end of the Nominations phase, but are sometimes held at the beginning of voting.

The way to set this up is:

1. File a ticket in Steering, and ping the SC over slack to set up a time or times
2. Remind the SC until you have 1 or 2 sessions scheduled
3. Invite all candidates then announced to the sessions

### Voting Phase

The Voting Phase is the actual "election", and is the time between the election dates in Elekto.  During this phase, the EOs primary job is promoting voter turnout.  Historically, Steering elections have had around 1/3 participation, and go down slightly each year.  As such, the key activity during this phase is the various announcements.

#### Initial Voting Announcement

For attention, EO Comms wants to make the opening of voting as grand as possible.  This would include all of the standard methods above.  Also, Comms should consider a blog post on the Contributor Blog that lists all the candidates and their bios.

#### Reminders

You'll see from the sample calendar that there are many reminders during voting.  These may seem excessive but they are completely necessary.

One additional reminder option is to send one email reminder to voters.  The [email script] will pull email addresses for most (but not all) voters.

#### Candidate Campaigning

Kubernetes has a policy against "excessive campaigning", which in practice means that the candidate is not supposed to promote their own candidacy beyond simple statements that they are running.  Further, they should not use any privileged access to large groups (such as specific SIGs, events, or their coworkers) to encourage votes.  Judgements on this can be fairly fine-tuned; for example, Steering has previously ruled that a candidate *may* send an email to their coworkers reminding them of the election, but that email may not contain information on them as a candidate.

Should a candidate violate this, report it to Steering, who will decide what action is warranted.

As much as the EOs have time, then, they should attempt to make sure that the entire community has good information about all candidates, even beyond what's in the app.  One idea (not yet tried) would be to have panel discussions with the candidates on video.  The challenge is that all candidates would need to be included in at least one panel, making logistics difficult.

#### Determining Results

After all of the reminders, the election concludes automatically once past the election close date.  Any extension of that date would need to be approved by Steering.

Once the election is closed, on the Admin screen you will see a Generate Results option which will allow you to see who "won".  This is a pure Condorcet ranking, and ideally the EOs will select the top # of winners, depending on how many seats are open.  However, proportional representation can interfere with that.

After selecting the initial list of winners, count the number of Steering members working for each employer that would be in the new SC.  If more than two members would share the same employer, then any newly elected candidates from that employer drop out of the winners list, and choices move one down.  In the as-yet-unprecedented event that a candidate was disqualified from the race due to excessive campaigning or CoC violations, then remove that candidate from the winners list.

If removal happens, all three EOs need to discuss and achieve consensus on the removal.  If no removal is happening, any one EO can report the results.

The remaining "top #" winners are the new Steering members.

#### Notifying the Community

Notifications of the results happen in the following order:

1. Notify non-involved Steering
2. Notify the candidates themselves
3. Announce to the community
4. Update results.md
5. Publish blog post

The first step is to contact the non-involved Steering members.  Steering will approve the election results as reported by the EOs. This communication will also share the complete ranking of all candidates, how many exceptions were requested and how many granted, and whether proportional representation limits affected the result.

Secondly, the EOs should notify the candidates themselves.  This should happen about 24 hours before the election is officially announced, in a communication that tells the candidates to keep the results to themselves.

Then the results are announced to the community.  Sometimes this happens via an event.  In the past, it was at the Community Meeting.  A couple of other times, Steering announced the results at an SC meeting.  Other times, though, the announcement has been strictly online.  Even if there is an event, EO Comms will work with a designated Steering member to announce the results via k/dev and the slack channels.

After the announcement, an EO will compose a brief version of a results.md file and merge it in so that the results display in Elekto.  Until this is done, voters using the app cannot see the results.

Finally, the EOs with Steering will publish a blog post to the Kubernetes blog announcing the new Steering committee and thanking departing members for their service.

### Post Election

Once the election is concluded, the EOs have a bit of cleanup.

#### Upload Ballots.csv

Within a week after the election, one EO should download the anonymized ballots from Elekto and check-in the ballots.csv file to the appropriate election folder.  This allows the election results to be checked, and the election database to be recreated, at any time in the future.

#### Retro

Within 2 weeks of the election, the EOs should hold a retro meeting ([2022 retro]) for the election.  The Retro doc should have been opened near the beginning of the election cycle (see schedule above), allowing the EOs to note issues and successful ideas as they happen.  Right after the election closes, they should annotate any additional thoughts and then schedule a meeting that all EOs and shadows can attend.

In addition to the EOs, the following folks should be invited to the retro, but are optional:

* Members of the Election Subproject
* Steering Committee members (both old and new)

Once the retro meeting has been held and the doc is complete, it should be copied as a markdown file into that year's election folder for future reference.

#### Special Elections

While it has not happened to date, Kubernetes' bylaws provide for a special election to be held during the year if too many members resign from the SC at once, or if a special election is triggered by someone being dismissed from the SC by a vote of the committee, or if the SC is dissolved by vote.  In any of these cases, the EOs from the last prior steering election are expected to administer this special election.  If some EOs are unavailable (due to absence or running in the special election), then the Alternate(s) will serve.  If there still aren't enough EOs, the Elections Subproject members will chose one or more emergency EOs from the Subproject members or from the list of prior EOs.

These Special Elections will be held on a shorter schedule than the regular election, if possible.  The EOs should organize and announce the election as quickly as possible, and consider reducing the nomination and voting phases to two weeks each instead of three.  

### Additional Notes on Being an EO

#### Impartiality

A key part of making Kubernetes elections work is trust in the impartiality of the system. Election Officers, and Alternates, should be careful to avoid any activities that would make them appear biased towards one or more candidates.  In particular, EOs are discouraged from:

* Endorsing or criticizing particular candidates, either in Kubernetes areas or on social media
* Giving any candidate special assistance that is not offered to other candidates
* Promoting any candidate at the EO's workplace or in other external organizations
* Posting personal opinions on the election that do not represent the consensus of the EOs
* Taking on multiple EO roles such that the other EOs cannot check their work

#### EOs Dropping Out

EOs may not be able to complete the election for a variety of reasons, including illness, personal conflicts, or an inability to remain impartial.  In any of these cases, the EO should notify their fellow EOs, and the Alternate will be promoted to take their places.  If any EO is unexpectedly unresponsive for several days in a row, the other EOs may decide they are inactive and promote the Alternate.  In either case, they should notify the Subproject and Steering of the change and the reason for it immediately.

Should more than one EO drop out, leaving less than three EOs, then the remaining EO(s) should contact the Election Subproject to recruit/appoint new EOs as soon as possible.  In most cases, this will also require delaying the election schedule.

[Elections Subproject]: /elections/README.md
[select Election Officer candidates]: /elections/README.md#recommending-election-officers
[Elekto]: https://elekto.dev
[Devstats]: https://github.com/cncf/devstats/issues
[Org Repo]: https://github.com/kubernetes/org
[election website]: elections.k8s.io
[steering election README]:/elections/steering/README.md
[elections doc]: https://github.com/kubernetes/steering/blob/main/elections.md
[nominations]: https://github.com/kubernetes/steering/blob/main/elections.md#eligibility-for-candidacy
[elections alias]: https://github.com/kubernetes/k8s.io/blob/main/groups/committee-steering/groups.yaml
[election README]: /elections/documentation/template/README.md
[messaging folder]: /elections/documentation/messaging
[k/dev]: https://groups.google.com/a/kubernetes.io/g/dev/
[email script]: https://github.com/elekto-io/elekto/blob/main/scripts/elekto_emails.py
[Comms Plan]: https://docs.google.com/document/d/1zhZzjKi-VHD1xfdibX68VxYaQ_-PiFKanGzBcF-1ilQ/edit?usp=sharing
[2022 retro]: https://docs.google.com/document/d/1M8Ho1Bx9WkmNrzc1eoSJPsQxaO9smEtp-mxvpc4s5i8/edit?usp=sharing
[generate a CSV to download]: https://k8s.devstats.cncf.io/d/13/developer-activity-counts-by-repository-group?orgId=1&var-period_name=Last%20year&var-metric=contributions&var-repogroup_name=All&var-repo_name=kubernetes%2Fkubernetes&var-country_name=All&inspect=4&inspectTab=data#:~:text=Formatted%20data-,Download,-CSV
