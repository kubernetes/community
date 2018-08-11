---
kep-number: 0007
title: A community forum for Kubernetes
authors:
  - "@castrojo"
owning-sig: sig-contributor-experience
participating-sigs:
reviewers:
  - "@jberkus"
  - "@joebeda"
  - "@cblecker"
approvers:
  - "@parispittman"
  - "@grodrigues3"
  
editor: TBD
creation-date: 2018-04-03
last-updated: 2018-04-17
status: implemented
---

# A community forum for Kubernetes

## Table of Contents

* [Table of Contents](#table-of-contents)
* [Summary](#summary)
* [Motivation](#motivation)
    * [Goals](#goals)
    * [Non-Goals](#non-goals)
* [Proposal](#proposal)
    * [User Stories [optional]](#user-stories)
      * [Story 1](#story-1)
      * [Story 2](#story-2)
    * [Implementation Details/Notes/Constraints](#implementation-details)
    * [Risks and Mitigations](#risks-and-mitigations)
* [Graduation Criteria](#graduation-criteria)
* [Implementation History](#implementation-history)
* [Drawbacks](#drawbacks)


## Summary

Kubernetes is large enough that we should take a more active role in growing our community. We need a place to call our own that can encompass users, contributors, meetups, and other groups in the community. Is there a need for something between email and real time chat that can fulfill this? The primary purpose of this KEP is to determine whether we can provide a better community forum experience and perhaps improve our mailing list workflow. 

The site would be forum.k8s.io, and would be linked to from the homepage and major properties. [See KEP005](https://github.com/kubernetes/community/blob/master/keps/sig-contributor-experience/0005-contributor-site.md) for related information on a contributor website. 

## Motivation

- We're losing too much information in the Slack ether, and most of it does not show up in search engines. 
- Mailing lists remain mostly the domain of existing developers and require subscribing, whereas an open forum allows people to drive by and participate with minimal effort.
  - There's an entire universe of users and developers that we could be reaching that didn't grow up on mailing lists and emacs. :D
  - Specifically, hosting our lists on google groups has some issues:
    - Automated filtering traps Zoom invites for SIG/WG leads
    - Cannot use non-google accounts as first class citizens (Google Account required to create/manage group, join a group)
    - Hard to search across multiple lists
    - There's no way to see all the kubernetes lists in one view, we have to keep them indexed in sigs.yaml
    - Filtering issues with the webui with countries that block Google 
    - Non-kubernetes branding
- As part of a generic community portal, this gives people a place to go where they can discuss Kubernetes, and a sounding board for developers to make announcements of things happening around Kubernetes that can reach a wider audience.
- We would be in charge of our own destiny, (aka. The Slack XMPP/IRC gateway removal-style concerns can be partially addressed)
  - Software is 100% Open Source
  - We'd have full access to our data, including the ability to export all of it.
  - Kubernetes branded experience that would look professional and match the website and user docs look and feel, providing us a more consistent look across k8s properties. 

### Goals

- Set up a prototype at discuss.k8s.io/discuss.kubernetes.io
  - Determine if the mailing list feature is robust enough to replace our google groups
  - References: [Mailing list roadmap](https://meta.discourse.org/t/moss-roadmap-mailing-lists/36432), [Discourse and email lists](https://meta.discourse.org/t/discourse-and-email-lists-like-google-groups/39915)
- Heavy user engagement within 6 months.
  - Clear usage growth in metrics
    - Number of active users
    - Number of threads 
    - Amount of traffic 
  - SIG Contributor Experience can analyze analytics regularly to determine growth and health. 
  - A feedback subforum would enable us to move quickly in addressing the needs of the community for the site. 

### Non-Goals

- This is not a proposal to replace Slack, this is a proposal for a community forum.
  - The motivation of having searchable information that is owned by the Kubernetes community comes from voiced concerns about having so much of Kubernetes depend on Slack. 
  - You are encouraged to propose a KEP for real-time communication if you would like to champion that, this KEP is not about Slack. 
- This does not replace Stack Overflow or kubernetes-users for user support. 
  - However inevitably users who prefer forums will undoubtedly use it for support.
  - Strictly policing "this should be posted here, that should be posted there" won't work. 
  - I believe our community is large enough where we can have a support section, and there are enough people to make that self sustaining, we can also encourage cross posting from StackOverflow to integrate things better as both sites have good integration points.
  - Over time as community interaction and knowledge base builds people will end up with a better experience than in #kubernetes-users on slack and will naturally gravitate there.
  - Other large OSS communities have a presence on both StackOverflow and do support on their user forums already and it doesn't appear to be a big issue.
- This will not replace kubernetes-devel or SIG mailing lists.
  - They work, but we could experiment with mailing list integration. (See below)
  - Let's concentrate on an end-user experience for now, and allow SIGs and working groups who want a more user-facing experience to opt-in if they wish. 

## Proposal

### User Stories

- A place for open ended discussion. For example "What CI/CD tools is everyone using?"
  - This would be closed as offtopic on StackOverflow, but would be perfect for a forum.
- Post announcements about kubernetes core that are important for end users
  - An announcements subforum can be connected to Slack so that we have a single place for us to post announcements that get's propagated to other services.
- Post announcements about related kubernetes projects
  - Give the ecosystem of tools around k8s a place to go and build communities around all the tools people are building. 
  - "Jill's neat K8s project on github" is too small to have it's own official k8s presence, but it could be a post on a forum. 
- Events section for meetups and Kubecon
- Sub boards for meetup groups
- Sub boards for non-english speaking community members
- Developer section can include:
  - Updated posting of community meeting notes
  - We can inline the youtube videos as well 
  - Steering Committee announcements
  - Link to important SIG announcements
  - Any related user-facing announcements for our mentorship programs
- Job board
  - This might be difficult to do properly and keep classy, but leaving it here as a discussion point.
- reddit.com/r/kubernetes has some great examples
  - "What are you working on this week?" to spur activity
  - "So, what's the point of containerization?" would be hard to scope on StackOverflow, but watercooly enough for a forum.
  - [Top discussions](https://www.reddit.com/r/kubernetes/top/?t=month) over the last month. 

### Implementation Details

- Software 
  - Discourse.org - https://www.discourse.org/features
  - Free Software
  - Vibrant community with lots of integrations with services we already use, like Slack and Github 
  - Rich API would allow us to build fun participatory integrations (User flair, contributor badges, etc.)
  - Facilities for running polls and other plugins
  - SSO with common login methods our community is already using.
  - Moderation system is user-based on trust, so we would only need to choose 5 people as admins and then as people participate they build trust and get more admin responsibilities. 
  - Other developer and OSS communities such as Docker, Mozilla, Ubuntu, Twitter, Rust, and Atom are already effectively, software is mature and commonly used. 
  - Friendly upstream with known track record of working with other OSS projects.
- Hosting 
  - We should host with a Discourse SaaS paid plan so we can concentrate on building the community and k8s itself.
    - [Pricing information](https://payments.discourse.org/pricing)
    - If other CNCF projects are interested in this we could help document best practices and do bulk pricing. 
  - Schedule regular dumps of our data to cloud provider buckets
  - Exporting/Self Hosting is always available as an option 
  - Google Analytics integration would allow us to see what users are interested in and give us better insight on what is interesting to them. 
  - Data explorer so we can run ad hoc SQL queries and reports on the live data.
  - Mailing list import would allow us to immediately have a searchable resource of our past activity. 

### Risks and Mitigations

- One more thing to check everyday(tm)
  - User fatigue with mailing lists, discourse, slack, stackoverflow, youtube channel, kubecon, your local meetup, etc.
  - This is why I am proposing we investigate if we can replace the lists as well, two birds with one stone. 
- Lack of developer participation
  - The mailing lists work, how suitable is Discourse to replace a mailing list these days? CNCF has tried Discourse in the past. See [@cra's post](https://twitter.com/cra/status/981548716405547008)
  - [Discussion on the pros and cons of each](https://meta.discourse.org/t/discourse-vs-email-mailing-lists/54298)
  - We have enough churn and new Working Groups that we could pilot a few, opt-in for SIGs that want to try it? 
- A community forum is asynchronous, whereas chat is realtime.
  - This doesn't solve our Slack lock-in concerns, but can be a good first step in being more active in running our own community properties so that we can build out own own resources. 
  - Ghost have [totally migrated to Discourse](https://twitter.com/johnonolan/status/980872508395188224?s=12) and shut down their Slack.
    - We should keep an eye on this and see what data we can gleam from this. Engage with Ghost community folks to see what lessons they've learned.
    - Not sure if getting rid of realtime chat entirely is a good idea either. 
- [GDPR Compliance](https://www.eugdpr.org/)
  - Lots of data retention options in Discourse. 
  - We'd need to engage with upstream on their plans for this, we would want to avoid having to manage this ourselves. 

#### References from other projects

- [Chef RFC](https://github.com/chef/chef-rfc/blob/master/rfc028-mailing-list-migration.md)
  - [Blog post](https://coderanger.net/chef-mailing-list/) from a community member - good mailing list and community feedback here. 
- [Swift's Plan](https://lists.swift.org/pipermail/swift-evolution/Week-of-Mon-20170206/031657.html) - Long discussion, worth reading
- [HTM Forum](https://discourse.numenta.org/t/guidelines-for-using-discourse-via-email/314)
- [Julia](https://discourse.julialang.org/t/discourse-as-a-mailing-list/57) - It might be useful for us to investigate pregenerating the mail addresses?
- [How's Discourse working out for Ghost](https://forum.ghost.org/t/hows-discourse-working-out-for-ghost/947) - We asked them for some direct feedback on their progress so far 


## Graduation Criteria

There will be a feedback subforum where users can directly give us feedback on what they'd like to see. Metrics and site usage should determine if this will be viable in the long term.

After a _three month_ prototyping period SIG Contributor Experience will:

- Determine if this is a better solution than what we have, and figure out where this would fit in the ecosystem
  - There is a strong desire that this would replace an existing support venue, SIG Contributor Experience will weigh the options.  
- If this solution is not better than what we have, and we don't want to support yet another tool we we would shut the project down.
- If we don't have enough information to draw a conclusion, we may decide to extend the evaluation period.
- Site should have a moderation and administrative policies written down.


## Implementation History

Major milestones in the life cycle of a KEP should be tracked in `Implementation History`.
Major milestones might include

- the `Summary` and `Motivation` sections being merged signaling SIG acceptance
- the `Proposal` section being merged signaling agreement on a proposed design
- the date implementation started
- the first Kubernetes release where an initial version of the KEP was available
- the version of Kubernetes where the KEP graduated to general availability
- when the KEP was retired or superseded

## Drawbacks

- Kubernetes has seen explosive growth without having a forum at all. 
