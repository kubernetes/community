---
title: "NCO Lead Handbook"
weight: 20
description: "Guidance for leading and organizing New Contributor Orientation (NCO) meetings."
---

## Meeting Lead Roles

The NCO meeting Lead is responsible for running the NCO meeting for at least 1 [monthly] session. Each NCO Lead is responsible for identifying at least 1 NCO Lead to succeed them.

## NCO Lead Pre-requisites
* [Kubernetes Org Membership](https://github.com/kubernetes/community/blob/master/community-membership.md#member)
  * The NCO meeting lead must be an active contributor.
  * An NCO Lead should also be active on the [Kubernetes Slack](https://slack.k8s.io)
* Availability to lead at least 1 meeting in at least 1 month
* Attend at least 1 NCO session or have watched at least 1 recording of an NCO session
* Interest in helping new and prospective contributors find their way in the community

## NCO Lead Responsibilities
* Determine the time of the monthly session
  * Coordinate with sig-contribex leads (sig-contribex-leads@kubernetes.io) if meeting times need to be changed from the standard schedule.
* Create and configure meeting materials for the lead's session(s)
  * Create a new slide deck for the session based on either the previous month's session or the [template](https://docs.google.com/presentation/d/1EHAqousQCzoaKi90JFAE5Ta07Qab1VjMppJ9XTFRzc8/edit?usp=sharing)
  * Contact SIG Leads to update the "Calls for Help" section in the session. Comms should be sent to SIG leads *at least one week in advance* via the [chairs-and-techleads channel](https://kubernetes.slack.com/archives/CD6LAC15M) in Slack **and** the leads@kubernetes.io mailing list.
* Publish the meeting content in [this repo's nco-slides folder](nco-slides) at least 1 day before the session is held
  * Publishing the meeting content in advance of the meeting will make the content more accessible to prospective attendees, and allow them to judge the value of the meeting in advance.
* Lead the NCO Meeting
  * Each NCO meeting agenda should consist of roughly 40min of presented content followed by roughly 20min of freeform Q&A. The lead is responsible for keeping track of time to ensure the meeting sticks roughly to the scheduled agenda. Both the presentation and Q&A portions of the meeting are important for delivering on the goals of the NCO initiative.
  * If the lead will not give the presentation themselves, they are responsible for identifying a host who will give the presentation effectively and efficiently.
* Identify at least 1 succeeding lead
  * If you are unable to identify a lead to succeed you, contact the sig-contribex leads (sig-contribex-leads@kubernetes.io) as early as is reasonable.


## Determining Meeting Time(s)
NCO Meetings aim to make contribution more accessible to anyone interested in it. As such, meetings should take into account accessibility in different timezones around the world. Typically, this requires running the meeting at least twice. The recommended schedule for sessions is:
* [**EMEA/APAC-friendly**: 1:30 PT / 8:30 UTC / 10:30 CET / 14:00 IST](https://calendar.google.com/calendar/event?action=TEMPLATE&tmeid=NXVpdGhoMWRyMGhpMDZjd[…]afef04s12ra0gkqql6fchjc%40group.calendar.google.com&scp=ALL)
* [**AMER-friendly**: 8:30 PT / 15:30 UTC / 17:30 CET / 21:00 IST](https://calendar.google.com/calendar/event?action=TEMPLATE&tmeid=MnZqMXVmazZhNWJ2aTNld[…]afef04s12ra0gkqql6fchjc%40group.calendar.google.com&scp=ALL)

The [SIG ContribEx calendar](https://calendar.google.com/calendar/embed?src=c8bafef04s12ra0gkqql6fchjc%40group.calendar.google.com&ctz=America%2FLos_Angeles) is configured to feature NCO Meetings at these times on a monthly basis. If the meeting times need to be changed or any sessions need to be cancelled, reach out to the SIG ContribEx leads at sig-contribex-leads@kubernetes.io *as early as possible*.

## NCO Meeting Materials
Each session's slide deck should be constructed based on the [template](https://docs.google.com/presentation/d/1EHAqousQCzoaKi90JFAE5Ta07Qab1VjMppJ9XTFRzc8/edit?usp=sharing).

The NCO Meeting slide deck is designed to be relatively easy to understand and deliver in a timely and effective manner. Every slide includes a script in the speaker notes which can be used word-for-word to deliver the content. Some areas of the slide deck must be updated for each session or month. Sections that require updates include:
* The hosts slide
* The "Calls for Help" section

Evergreen opportunities should be carried forward with each slide deck. It is the lead's responsibility to confirm that the evergreen opportunities listed are still available. "Calls for Help" with skill requirements should be removed from the deck between each month's sessions unless otherwise noted. This section should be filled in fresh each month by the SIG Leads. It is the NCO Meeting Lead's responsibility to ensure SIG Leads are aware of this opportunity to share their calls for help, and that they are informed of the opportunity and how to update the slide deck in a timely fashion.


Some slides include a date on them, usually in the upper righthand corner, to indicate when the data on them was last updated. If you believe the data to be sufficiently out of date as to need updating for your session, you may either do so yourself, or request help from the sig-contribex leads (sig-contribex-leads@kubernetes.io). If you would like to request help from the sig-contribex leads to update content, please do so *as early as possible*.

## NCO Meeting Comms
The NCO Meeting Lead is responsible for ensuring any communications (comms) related to the NCO meeting go out in a timely fashion. The NCO Lead may contact the [sig-contribex comms subproject team via Slack](https://kubernetes.slack.com/archives/C03KT3SUJ20) as needed for assistance with creating and sending comms. Regular NCO meeting comms should include:

###SIG Lead Comms
Information on how to add new Calls for Help to the upcoming NCO session deck(s) for the month should be sent to the SIG Leads *at least 1 week in advance*. More frequent or earlier comms at the lead's discretion are recommended. The channels that can be used to send these communications are:
* Email: leads@kubernetes.io
* Slack: [chairs-and-techleads channel](https://kubernetes.slack.com/archives/CD6LAC15M)

Example comms:
| Hello SIG Chairs & TLs!

SIG Contributor Experience is spinning up a new initiative to run a monthly Kubernetes New Contributor Orientation meeting. The purpose of this meeting is to introduce prospective contributors to the how the Kubernetes Project and Contributor Community work, and to provide high-level guidance on how to get started contributing.

Action Item 1: Add your Calls for Help to the slide deck!
The slide deck for the call [1] includes a section where SIGs and WGs can add current open Calls for Help. This section starts at slide 56 and groups Calls for Help into 2 types: Evergreen and general "Calls for Help" (everything else). Evergreen tasks are those which are recurring or always available, like taking meeting notes or release shadowing. We do not expect every SIG to have Evergreen opportunities, though we would love to help you if you want to develop some (Reach out to Kaslin Fields in Slack)! Most SIGs likely have open Calls for Help. There is guidance in the slide deck explaining the need for specific skills for some tasks, so please don't hesitate to add Calls for Help that require specific skills.

Action Item 2: Add yourself as a host
There will be a Q&A portion of the meeting where attendees can ask questions, so even just being present can be helpful! If you plan to attend and want to be called out as an active contributor who can help answer questions, please feel free to add yourself as a host on slide, which is either Slide 4 or 5 depending on which meeting you intend to participate in.

We will be setting up a section in the Mentoring repo with guidance on how to run these meetings, and would like to invite others to get involved with them and to help run them. We'll share the link to this once it is ready.

EMEA/APAC-friendly is at 1:30 PT / 8:30 UTC / 10:30 CET / 14:00 IST [2]
AMER-friendly is at 8:30 PT / 15:30 UTC / 17:30 CET / 21:00 IST [3]

[1] - https://docs.google.com/presentation/d/1WO7vj33cVy1-rOav3CevAGrw97VrbIAFo_SQXbUwy28/edit?usp=sharing
[2] - https://calendar.google.com/calendar/event?action=TEMPLATE&tmeid=NXVpdGhoMWRyMGhpMDZjd[…]afef04s12ra0gkqql6fchjc%40group.calendar.google.com&scp=ALL
[3] - https://calendar.google.com/calendar/event?action=TEMPLATE&tmeid=MnZqMXVmazZhNWJ2aTNld[…]afef04s12ra0gkqql6fchjc%40group.calendar.google.com&scp=ALL |

### Attendee Comms
Communications about the event's purpose and how to attend should be sent at least *1 week in advance*. More frequent or earlier comms at the lead's discretion are recommended. The channels that can be used to send these communications are:
* The [kubernetes-new-contributors channel on Slack](https://kubernetes.slack.com/archives/C09R23FHP)
* [dev@kubernetes.io](https://groups.google.com/a/kubernetes.io/g/dev)
* The K8sContributors channels on Social Media (Contact the [sig-contribex comms subproject team via Slack](https://kubernetes.slack.com/archives/C03KT3SUJ20) for assistance)

Example attendee comms:
| Hello, New Contributors! The next Monthly Kubernetes New Contributor Orientation (NCO) meeting will be happening on the 3rd Tuesday of this month. The purpose of this meeting is to orient you in the Kubernetes community by providing you with a brief overview of what Kubernetes does as a technology, how the community is structured, and how you might fit within it. The meeting is 1hr long and will consist of ~40min of presented content followed by ~20min of freeform Q&A with active contributor(s). Note there is no hands-on content in this session.
This month's sessions will be held at::
EMEA/APAC-friendly: 1:30 PT / 8:30 UTC / 10:30 CET / 14:00 IST
AMER-friendly: 8:30 PT / 15:30 UTC / 17:30 CET / 21:00 IST
The meetings will be recorded and posted on YouTube. You can add the meeting invites to your calendars using these links:
EMEA/APAC meeting link: https://calendar.google.com/calendar/event?action=TEMPLATE&tmeid=NXVpdGhoMWRyMGhpMDZjd[…]afef04s12ra0gkqql6fchjc%40group.calendar.google.com&scp=ALL
AMER meeting link: https://calendar.google.com/calendar/event?action=TEMPLATE&tmeid=MnZqMXVmazZhNWJ2aTNld[…]afef04s12ra0gkqql6fchjc%40group.calendar.google.com&scp=ALL |
