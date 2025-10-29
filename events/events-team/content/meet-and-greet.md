# SIG Meet And Greet Handbook

## Overview

The SIG Meet and Greet (formerly known as Meet the Kubernetes Contributor Community; now abbreviated M&G) is an onsite event we hold during KubeCon + CloudNativeCon where new contributors, experienced contributors, or end users can find a Kubernetes SIG or WG to ask questions or inquire about contributing. SIG contributors occupy tables, and new and casual contributors drop by to chat. Secondarily, it is a lunchtime event for SIG members to see each other and chat.

This document defines activities needed to run a Kubernetes SIG Meet and Greet:

- [Overview](#overview)
- [Roles and Staffing](#roles-and-staffing)
- [Skills and Qualifications](#skills--qualifications)
- [Time Commitment](#time-commitment)
- [Audience](#audience)
- [Time and Location](#time-and-location)
- [Preparation & Timeline](#preparation--timeline)
- [Responsibilities](#responsibilities)


## Roles and Staffing

Aside from the SIG leads or representatives, this event is run by a lead (and possibly a shadow and onsite volunteer). The lead role is selected by the Maintainer Summit leads.

The M&G lead should be onsite at KubeCon + CloudNativeCon.

The M&G lead is responsible for the onsite Kubernetes SIG Meet and Greet at KubeCon + CloudNativeCon which includes the following:
  - coordinating with the CNCF with the date, location, and time of the SIG Meet & Greet
  - coordinating with Kubernetes SIGs to have a representative(s) onsite at the SIG Meet & Greet
  - coordinate with the communications team to broadcast the event
  - assemble table signage onsite
  - during the event, welcome new contributors and guide folks to their interested SIG or WG
  - on occasion, facilitate space for overflow
  - breakdown table signage after the event

### Skills & Qualifications

- Familiarity with the SIGs, including understanding how SIGs are related to each other and any likely overlap in membership.
- Good organizational skills, including being able to relentlessly follow up and get responses if SIGs are unresponsive.
- Writing and communication skills to create messages to broadcast the event, create a request for SIG and WG representatives, and follow up with SIGs
- Be welcoming to new contributors

### Time Commitment

  - 1-2 hours a week from 3 months in
  - 2-3 hours a week in the month leading up to the event
  - On-site: 1 hour before the event to setup and supervisory duties for the entire event

## Audience

The audience for the M&G consists of three groups:

- *SIG and WG Leads and Subproject Owners and Committee Members*: anyone who considers themselves a "SIG/WG member" can attend and work the tables. We particularly want SIG leads and/or Subproject owners who can guide contributors in joining their SIG and/or contributing to a project. Committee members are also encouraged to be at a table or otherwise identify themselves with pins to make it easier for contributors to stop by for a chat.
- *New & Casual Contributors*: the visitors for the SIG Meet and Greet will be people who are interested in contributing to Kubernetes or who have contributed lightly but not really participated in a SIG.
- *Regular Contributors*: The SIG Meet and Greet is also for regular contributors who are trying to reach out to a SIG, usually about a stuck PR, review, or mutual project.

Some teams that are not SIGs or WGs may also participate, e.g. ClusterAPI or the CoCC. There is no specific policy on this, deal with it on a case-by-case basis. Generally speaking, we want to say yes to teams that actually plan to show up as a group. In the rest of this document, "SIGs" should be considered to refer to "SIGs, WGs, and Teams".

## Time and Location

The M&G is generally held during lunch period on one of the main KubeCon + CloudNativeCon days. CNCF staff decides the exact schedule in order to not conflict with other lunchtime events.

## Preparation & Timeline

### 2-3 Months Before: Room request and SIG gear

In a meeting with a CNCF liaison assigned to the M&G, discuss the following: 
- goals of the M&G
- what worked or didn't work in the previous M&Gs 
- requirements for the M&G like space, preferred day and time
- review gear like table signs and buttons
- feel free to add or remove additional discussion topics

In the previous M&G (NA 2024), the M&G was held at the Project Pavilion. Verify with the CNCF if the M&G will be held at the Project Pavilion. If not, request a large room for the M&G with at least 16 round tables (at the time of this commit, there are 24 SIGs and 8 WGs -- 16 tables for 2 SIGs or WGs per table) and possibly extra tables for overflow.
Request for a day other than the last day of KubeCon as many SIG leads will likely have maintainer track sessions on the last day of KubeCon.

The gear for the M&G consists of table sign stands for each SIG, a bag of buttons for each SIG, and Kubernetes Contributor badge ribbons for folks who might not have gotten one. Some of this gets used/goes missing each KubeCon. As the M&G lead, you will get access to a spreadsheet to fill out to identify which SIGs are active and how many buttons should be present for each SIG. Use this sheet to ask CNCF staff to check the gear. Ensure, though the sheet, that:

- All currently active SIGs and WGs have at least a dozen buttons 
- All currently active SIGs and WGS have table signs, and non-broken table stands

The CNCF staff will handle actual inventory because they have possession of the case of SIG gear.  However, the SIG Meet and Greet lead should help by giving them a shortlist of the SIGs and WGs most likely to staff the event, as well as checking with them to remove any obsolete SIGs or WGs.

If Kubernetes has a new SIG or WG this year, or a team that really wants to recruit, notify the staff as soon as possible so that they can order new gear for that group.

Coordinate with the Maintainer Summit / Kubernetes Track lead on the creation of the Google spreadsheet for the CNCF to track current SIGs and WGs, pins, signage, etc. 
Create a list of current SIGs and Working Groups in kubernetes/community and populate the tracking Google spreadsheet, this list will be used as a reference for table signs.

Additional gear that you might want to have at the M&G, depending on circumstances and availability:

* Approver pins
* Kubernetes contributor patches
* Stickers

When a date and location have been verified by the CNCF, verify the M&G is on sched.com

### 1 Month Before: Recruit SIG staffers

Prepare a sign-up sheet for SIG volunteers, usually a GitHub issue (e.g. https://github.com/kubernetes/community/issues/8112). 

Then work with the event comms lead to message all of the SIG leads; you may end up doing this yourself, but ensure that you work with the comms lead to avoid overwhelming a channel with a notification. Notify folks of the signup opportunity on the following channels:

- Slack, #chairs-and-techleads
- leads mailing list
- Ideally, a message to each of the primary #sig-* channels (this is a good task for a shadow, and should probably be bundled with outer communications)
    - Generally, wait to see how the first two comms gains traction before reaching out in each sig-* channel. The leads generally respond very quickly to the first two comms, and then you only have a handful of groups to reach out to versus everyone.

This solicitation might be bundled with other contributor-summit messaging that needs to go to all SIG leads, such as a call for booking SIG meeting rooms.

Add a page in the Contributor Summit website for the SIG Meet and Greet.

#### Follow-ups

Check weekly to see who's signed up. While participation by all SIGs is valuable, the following SIGs are particularly popular with 1st-time contributors, and if nobody signs up you should message them:
    - SIG Docs
    - SIG Release
    - SIG CLI
    - SIG Contributor Experience
    - SIG K8s Infra

### 1 Week Before: Solicit Attendees and Send Reminders

By this time, CNCF events should have given you a room number and day. Send out a reminder to all of the chairs and techleads as well as anyone who signed up, and give them the updated information including the location and time. Update the Summit website with the full information.

Working with Comms, advertise the SIG Meet and Greet to any and all KubeCon attendees and community members. This should include advertising it on community channels, including social media, #contributor-summit, #kubernetes-contributors, the kubernetes-dev mailing list, and LWKD. You should also make sure that it gets into at least one of KubeCon's official messages to attendees; ask CNCF about this.

*Stretch Goal: prepare a post-M&G survey.  We haven't actually done one of these before, so you'll be starting from scratch.  Good questions to ask: how do they feel it went, how many folks from their SIG were there, and how many contributors did they talk to (and maybe a newbie/experienced estimate).  For attendees, it would ask which SIG(s) they were seeking, why, and if they found what they were looking for. Print QR codes for the survey to have on the tables.*

Decide whether you'll have a sign-in sheet for attendees. This is unnecessary if you're doing a survey, but if you're not, you should at least have sign-ins. If so, design & print it out.

### 1 Day Before

If possible, check out the space assigned to the M&G, just in case there are problems with it. If there are, raise them with our CNCF event staff.

Check that all SIGs who intended to attend are accounted for and were able to make it to the event. In the case of anyone with last-minute travel cancellations, identify if someone can cover for them. If a SIG cannot get members out due to travel issues, solicit help to staff their table from related groups or SIG ContribEx.

If possible, put together short directions (plus pictures) to the the SIG Meet and Greet room from the main conference area to send to Comms for posting on social media and via email.

### Day Of

Get to the room 15-30 minutes before SIG volunteers are due to arrive. Work with CNCF staff to distribute the SIG packs among the tables. SIGs will have to share tables, so try to combine them in ways that make sense, such as SIG Testing with SIG K8s Infra and SIG Networking with SIG Multicluster. Consult the signup sheet, though, to see who's actually expected to show up. While signups are not reliable, they will tell you which SIGs are likely to have large contingents so that you don't seat them together.

If you don't have shadow(s), then recruit a couple folks from SIG ContribEx to help with setup.

Each pack has SIG buttons and a table stand with a sign. Set up the table stands for the SIGs who RSVPd; leave the others in the bag.

The SIG Beard pack is special. Leave it in the box unless someone who knows what to do with it shows up.

Wear some kind of Kubernetes gear, to help identify you as staff.

If possible, put together short directions (plus pictures) to the the SIG Meet and Greet room from the main conference area to send to Comms for posting on social media and via email.

An hour before the event:
- Identify the onsite CNCF liaison and obtain table signage hardware and signs.
- Assemble table signage onsite about 30 min to 1 hour before the event
- For easier navigation purpose, setup the table signs on tables in alphabetically order of SIGs then WGs so it is easier for people to find the SIG or WG
  - Since some people may represent multiple SIGs or WGs, people may join multiple SIG/WG signs on a table
- On occasion, facilitate space for overflow

### During the SIG Meet and Greet

Hover near the entrance to help attendees find the SIGs they are looking for. Sometimes they won't know, so be ready to chat about their interests with them and direct them appropriately.

Take pictures and forward them to Comms to send out on social media to further promote contributions.

### After the SIG Meet and Greet

Folks will probably still be chatting when the event ends. Generally, the space is not otherwise in use, and you don't need to kick them out. However, you do need to pack up the SIG packs for the staff; recruit anyone still at the tables to help. Members of ContribEx will also help with this.

### Day After

Send out a thank you to the SIG Leads, and a link to the survey, if you're doing one.
Add feedback to the event's retrospective agenda.
Update this doc if necessary.

## Notes and Tasks

Some additional notes based on experience from prior the SIG Meet and Greet events:

- Create a form to track the details about each SIG / WG representative. See the example from the [2019 San Diego SIG M&G](https://forms.gle/hxx1qz8XtwtXEBMm8) for more details about what information to gather.
- Contact each SIG / WG and ask them to have 1 or more experienced members complete the form and commit to represent the group at the SIG Meet and Greet.
- Expect to spend a fair amount of time following up with SIGs and WGs to get the name of their representatives.
- Add these representatives to a calendar notice for the SIG Meet and Greet.
- Ensure the table stands are all useable when you arrive. It's a lot easier to put each bag on tables, then set up the stands. The CNCF events liaison should have some extras in the box. 
- Create a plan and map for organizing SIGs / WGs at tables to make it easier for attendees to find SIGs. Get the room layout from the venue and pre-assign SIGs / WGs to tables with some thought given to logical groupings with similar topics in close proximity. Make sure that this plan is communicated to volunteers and other organizers and display the room layout / map on the projector to help people find their tables.
- Recruit a few volunteers to distribute table stands and buttons and to help SIGs / contributors find the right tables. 
- Walk around the room throughout the event to help lost contributors get matched up with a SIG table.
- A lot of contributors use this as a lunchtime hangout. Don't forget to get a lunch, yourself!

### Prior M&Gs

- [Meet the Kubernetes Contributor Community Chicago 2023](https://github.com/kubernetes/community/issues/7541)
- [SIG Meet and Greet Barcelona 2019](https://github.com/kubernetes/community/issues/3516)
- [SIG Meet and Greet San Diego 2019](https://github.com/kubernetes/community/issues/3896)
