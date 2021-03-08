# SIG Testing Annual Report - 2021

Authors: @spiffxp,
Draft date: 2021-03-08 <!-- TODO: s/Draft/Published/ -->
Based on template at: 

This report reflects back on CY 2020 and was written in Feb-Mar. 2021.

## Operational 

### How is the group doing with operational tasks in [sig-governance.md]?

#### Is your [README] accurate?
Yes. 

However, our [charter.md] could benefit from a refresh. It hasn't been
refreshed/reviewed in 2.5 years. e.g. 

- Missing from "in scope" is the fact that we've been actively involved in
maintaining the build system for kubernetes.
- The "deploying changes" section could also stand a refresh to reflect what
is concretely done today.

#### Have a [CONTRIBUTING.md] file?
No. In general, our subprojects follow the CONTRIBUTING.md files for each repo they cover. 

#### All subprojects correctly mapped and listed in [sigs.yaml]?
TODO: what does "correctly mapped" mean in this case?

#### What’s your meeting culture? 
We regularly meet bi-weekly Tuesdays at 10am PT, to try and overlap with PT, ET, and CET timezones.  Agenda is open to community suggestion, deadline to add is 1pm PT the day prior. If we have nothing on the agenda, we cancel.

We haven't had much in the way of regular structure, e.g. no regular subproject checkins, no review of testgrid dashboards, no slot just for issue triage.

We recently discussed our meeting culture at a meeting, and consensus amongst attendees was the following:
- people valued the time and enjoyed hanging out with fellow contributors
  - no desire to further reduce meeting frequency, nor move back to 30min/weekly
- the vast majority of speaker time goes to a few individuals, mostly the leads
  - leads are interested in how to encourage others to speak up or participate more
  - but attendees were generally happy with status quo, found leads informative
- "this meeting could have been an email" or "this could have been prerecorded"
  - the spontenaity of Q&A often provided the most insight to attendees
  - we are considering trialing use of GitHub discussions, or more regular e-mails for info-dumps etc.
- walking a board for issue triage eats entire meetings, more success when curating a few to discuss:
  - ... which is ultimately what lands on our agenda anyway: things contributors want to discuss/demo, and things leads want to discuss/demo
  - we are considering setting aside a block of time for help-wanted issues

##### Large/small, active/quiet, learnings? 
TOO, 10-20 attendees

##### Meeting notes up to date? 
Yes

##### Are you keeping recordings up to date/trends in community members watching recordings?
Yes, although uploading recordings is a manual process for us; they currently live on @spiffxp's
channel.  Can lag behind by a week or two.

690 views over the last year
- 290 came from a "Deflaking Kubernetes Tests" video clipped from SIG Testing meeting with @liggitt presenting.
- 5-32 view range for SIG Testing recordings from 2020
- no other subprojects are holding regular meetings

### How does the group get updates, reports, or feedback from [subprojects][subproject-definition]?
Nothing formal. Subproject owners will e-mail the list with announcements or design proposals,
show up at our meetings for design review or discussion. But not much in the way of regular
review of roadmaps/accomplishments etc.

#### Are there any springing up or being retired?
We moved the "testing-commons" subproject to "best effort", meaning it still
exists as a conceptually separate area of focus, and is owned by SIG Testing,
but nobody is actively leading the effort. We archived the #testing-commons
slack channel and have directed folks to come to the general #sig-testing
slack channel and meeting for future discussions on things in the
testing-commons area.

It is intended to represent ownership of everything under
kubernetes/kubernetes/test, not in terms of "we'll write and fix all of these
tests" (explicitly out of scope, per our charter) but instead in an advisory
capacity that aligns with our bandwidth.

#### Are OWNERS.md files up to date in these areas?
TODO

### How does the group get updates, reports, or feedback from Working Groups?
Same as for subprojects.  @spiffxp happens to be a lead for WG K8s Infra so often passes info back and forth.  We haven't yet formally shared info between WG Reliability and SIG Testing, though we've collaborated on some proposals.

#### Are there any springing up or being retired?
No

#### Are OWNERS.md files up to date in these areas?
TODO

### When was your last public community-wide update? 
TODO: (provide link to deck and/or recording)

## Membership

### Are all listed SIG leaders (chairs, tech leads, and subproject owners) active?
TODO

### How do you measure membership?
TODO By mailing list members, OWNERs, or something else?

### How does the group measure reviewer and approver bandwidth?
TODO

#### Do you need help in any area now?
TODO

#### What are you doing about it?
TODO

### Is there a healthy onboarding and growth path for contributors in your SIG?
TODO

#### What are some activities that the group does to encourage this?
TODO help-wanted issues

#### What programs are you participating in to grow contributors throughout the contributor ladder?
TODO

### What programs do you participate in for new contributors?
TODO

### Does the group have contributors from multiple companies/affiliations?
Yes

#### Can end users/companies contribute in some way that they currently are not?
TODO

## Current initiatives and project health

### What are initiatives that should be highlighted, lauded, shout outs, that your group is proud of?
TODO

#### Currently underway?
TODO

#### What are some of the longer tail projects that your group is working on?
TODO

### Year to date KEP work

#### What's now stable?
TODO

#### Beta?
TODO

#### Alpha?
TODO

#### Road to alpha?
TODO

### What initiatives are you working on that aren't being tracked in KEPs?
TODO

### What areas and/or subprojects does the group need the most help with?
TODO

### What metrics/community health stats does your group care about and/or measure?
TODO

Examples?     

<!-- general governance links -->
[sigs.yaml]: https://git.k8s.io/community/sigs.yaml
[sig-governance.md]: https://git.k8s.io/community/committee-steering/governance/sig-governance.md
[sig-definition]: https://github.com/kubernetes/community/blob/master/governance.md#sigs
[sig-roles]: https://github.com/kubernetes/community/blob/master/committee-steering/governance/sig-governance.md#roles
[subproject-definition]: https://github.com/kubernetes/community/blob/master/governance.md#subprojects
[wg-definition]: https://github.com/kubernetes/community/blob/master/governance.md#working-groups

<!-- sig-specific links -->
[charter.md]: https://git.k8s.io/community/sig-testing/charter.md 
[CONTRIBUTING.md]: https://git.k8s.io/community/sig-testing/CONTRIBUTING.md
[README]: https://git.k8s.io/community/sig-testing/README.md