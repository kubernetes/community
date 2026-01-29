---
title: "Moderation Rules and Responsibilities"
description: |
  Overview of community property moderator roles, responsibilities, selection,
  and best practices.
---

<!-- omit in toc -->
# Moderation on Kubernetes Communications Channels


This page describes the rules and best practices for people chosen to moderate
Kubernetes communications channels. This includes Github, Slack, forums, mailing
lists, YouTube, Zoom, and any property listed in the SIG Contributor Experience
[charter].

- Check the [centralized list of administrators][moderators] for contact
  information.
- Some Kubernetes properties, like the Twitter account, are managed by the CNCF.

---


- [Selection of Moderators](#selection-of-moderators)
  - [Moderators Pro Tempore](#moderators-pro-tempore)
- [Rotation of Moderators](#rotation-of-moderators)
- [Roles and Responsibilities](#roles-and-responsibilities)
- [Violations](#violations)
- [Escalation Procedures](#escalation-procedures)
- [Platform Specific Guidelines](#platform-specific-guidelines)
  - [Discuss](#discuss)
  - [Mailing List](#mailing-list)
  - [Slack](#slack)
  - [YouTube](#youtube)
  - [Zoom](#zoom)
  - [Process for Adding a Moderator](#process-for-adding-a-moderator)
- [References and Resources](#references-and-resources)

---

## Selection of Moderators

Each Kubernetes property has a certain set of [moderators] who
are responsible for keeping it safe and a fun place to participate.

Moderators are selected by following the criteria:

Moderators _MUST_:

- Be a [member of the Kubernetes Organization](/community-membership.md).
- Have experience moderating something or some equivalent level of community
  management.
- Make themselves available during their primary working hours for their given
  timezone.
- Communicate their availability to their peer moderators when appropriate, such
  as when travelling or becoming unavailable for an extended period of time.
- Understand that volunteering for this role might mean an occasional personal
  time commitment or off-hour duty.

The process for applying for moderatorship is as follows:

- Sponsored by 2 existing moderators **Note the following requirements for
  sponsors**:
    - Sponsors MUST have close interactions with the prospective member -
      example: participating in the appropriate property, coordinating on
      issues, etc.
    - Sponsors MUST be from multiple member companies to demonstrate integration
      across community.
    - Sponsors MUST take time zone coverage into account, each property should
      have global coverage. Ideally no more than two(2) moderators are needed in
      a given time zone.
    - Sponsors MUST ensure that nominees are familiar with the software they are
      using to moderate.
- **[Open an issue][moderator request] against the kubernetes/community repo**
   - Ensure your sponsors are @mentioned on the issue.
   - Complete every item on the checklist
   - Make sure that the list of contributions included is representative of your
     work on the project.
- Have your sponsoring reviewers reply confirmation of sponsorship: `+1` or
  similar approval.

### Moderators Pro Tempore

Each property will have a list of moderators who cannot commit to full time
moderatorship, but are available during special events or circumstances that
might call for additional ad-hoc duties.

For example if all moderators are attending a conference, pro tempore moderators
may be assigned to monitor a property.


## Rotation of Moderators

Content moderation can be personally tiring, so primary Moderators _SHOULD_
rotate on a regular basis.

- Primary moderators should evaluate their position(s) as a team _yearly_.
  - Determine whether the moderation situation on the property is working.
  - Rotate teams accordingly.
  - Consider rotating in less experienced person to give them an opportunity to
    participate.
  - Take into account time zone considerations.
- Due to less workload Moderators Pro Tempore should be a lightweight role
  - Primary moderators should consider switching to this role for a given amount
    of time to allow for a healthy rotation.


## Roles and Responsibilities

As part of volunteering to become a moderator you are now a representative of
the Kubernetes community, and it is your responsibility to remain aware of your
contributions in this space. These responsibilities apply to all Kubernetes
official channels.

Moderators _MUST_:

- Take action as specified by these Kubernetes Moderator Guidelines.
  - You are empowered to take _immediate action_ when there is a violation. You
    do not need to wait for review or approval if an egregious violation has
    occurred. Make a judgement call based on our [Code of Conduct] and Values
    (see below).
  - Removing a bad actor or content from the medium is required, do NOT let it
    sit there.
- Abide by the documented tasks and actions required of moderators.
- Ensure that the Kubernetes [Code of Conduct] is in effect on all official
  Kubernetes communication channels.
- Make yourself generally available during working hours in your time zone for
  moderation.
  - This can be handled as a group so that there is enough coverage of people to
    allow for absences/travel.
  - Ensure you are on #slack-admins during work hours and notifications are set
    appropriately.
- Become familiar with the [Kubernetes Community Values].
- Take care of spam as soon as possible, which may mean taking action by
  removing a member from that resource.
- Foster a safe and productive environment by being aware of potential multiple
  cultural differences between Kubernetes community members.
- Understand that you might be contacted by moderators, community managers, and
  other users via private email or a direct message.
- Keep up with software/platform changes on the property they are responsible
  for. This might include new UI changes, new features, or other software
  changes. Moderators are encouraged to meet regularly to train themselves how
  to be proficient with the platform.
- Report violations of the Code of Conduct to <conduct@kubernetes.io>.

Moderators _SHOULD_:

- Exercise compassion and empathy when communicating and collaborating with
  other community members.
- Understand the difference between a user abusing the resource or just having
  difficulty expressing comments and questions in English.
- Be an example and role model to others in the community. In many cases,
  moderators are some of the first people new contributors will interact with.
- Remember to check and recognize if you need take a break when you become
  frustrated or find yourself in a heated debate.
- Help your colleagues if you recognize them in one of the [stages of burnout].
- Be helpful and have fun!


## Violations

The Kubernetes [Code of Conduct Committee] will have the final authority
regarding escalated moderation matters. Violations of the Code of Conduct will
be handled on a case by case basis. Depending on severity, this can range up to
and including removal of the person from the community, though this is extremely
rare. This decision comes from the Code of Conduct committee, not the moderators.


## Escalation Procedures

In the event of large attacks the moderator team must enact the following
procedures:

- The person on call should immediately concentrate on removing the offending
  content and asking for other on call moderators for help. That is their sole
  responsibility.
- The secondary person on call should immediately begin to take notes to
  document the incident, this will form the basis of a post-mortem. The 2nd
  person on site is responsible for finding help, and documenting the incident.
- The secondary person on call will escalate if necessary. If it's a one off
  incident and the content is removed, then the collective moderators can work on
  a post-mortem and report the incident to primary moderators within 24 hours.
  - If it's a sustained incident that needs more help, the secondary will contact
    other primary moderators as soon as possible.
  - If appropriate, the next level of people to contact are the OWNERS of the
    subproject.
  - If appropriate, the next level of people to contact is the
    [Code of Conduct Committee].
  - If appropriate, the next level of people to contact is the
    [Steering Committee].
- Moderators will have access to a private document with contact information of
  the appropriate people.
- Primary moderators will then execute an audit of the affected property:
  - Slack: emoji audit
  - (More per-property steps to be added as we learn them)

### Escalation to Code of Conduct Committee Procedures

In the event of a violation of the Code of Conduct, the moderators will use the following procedures to escalate the issue to the Code of Conduct Committee:

If you believe that a poster is violating the Kubernetes Community Code of Conduct, escalate the matter to the [Code of Conduct Committee]. Include the following to the best of your ability in that message:

- A statement of the seriousness of the matter, what is the urgency of response?
- A statement of the nature of the violation with links to the offensive posts. It is useful to take a screenshot to preserve potentially ephemeral evidence.
- Include the contents of whatever conversations you have had with the poster in your capacity as moderator that are relevant to this matter.
- If there are historical details that are relevant also mention them.



## Platform Specific Guidelines

These guidelines are for tool-specific policies that don't fit under a general
umbrella.

### Discuss

- [Discuss Guidelines](./discuss-guidelines.md)
- [Moderators](./moderators.md#discuss.kubernetes.io)
- [Regional Moderators](./moderators.md#regional-category-moderators)

### Mailing List

- [Mailing List Guidelines](./mailing-list-guidelines.md)
- [Moderators](./moderators.md#mailing-list)

### Slack

- [Slack Guidelines](./slack-guidelines.md)
- [Moderators](./moderators.md#slack)

### YouTube

- Moderators should check the Comments section in the community tab regularly
  for published comments and the "hold for review" sections to see if comments
  are being flagged by the system.
- We do NOT use YouTube comments during our live streams, these are checked as
  OFF in the settings.
- [Youtube Guidelines](./youtube/youtube-guidelines.md)
- [Moderators](./moderators.md#youtube-channel)

### Zoom

- [Zoom Guidelines](./zoom-guidelines.md)
- [Moderators](./moderators.md#zoom)

### Process for Adding a Moderator

This is the workflow for adding a moderator.

- Moderator applies by [filing an issue] in the kubernetes/community repo
- Moderator gets approval from 2 current moderators
- Add the person to their respective moderation tool:
  - Slack - Add them as a slack workspace admins (instructions pinned on
    #slack-admins-private): Invite them to #slack-admins-private, #slack-admins,
    #slack-log, #slack-reports and #slack-invites
  - Discuss - Change their permission to moderator
  - Kubernetes.io Google group - PR the person into the correct yaml file in
    kubernetes/k8s.io. (Note: This only applies to groups using @kubernetes.io,
    most other lists are still managed out of band)
  - Ensure they add that they are an admin in their profile page on whatever
    service they are administrating
- Add them to the moderators mailing list by PRing them into the
  [moderators@kubernetes.io group] linking to the person's moderator github
  application as part of the PR.
- Ensure person is enrolled for a future bias training course, this can either
  be provided by the project or their employer.

[moderators@kubernetes.io group]: https://github.com/kubernetes/k8s.io/blob/main/groups/groups.yaml

## References and Resources

Thanks to the following projects for making their moderation guidelines public,
allowing us to build on the shoulders of giants. Moderators are encouraged to
learn how other projects moderate and learn from them in order to improve our
guidelines:

- Mozilla's [Forum Moderation Guidelines](https://support.mozilla.org/en-US/kb/moderation-guidelines)
- [5 tips for more effective community moderation](https://www.socialmediatoday.com/social-business/5-tips-more-effective-community-moderation)
- [8 Helpful Moderation Tips for Community Managers](https://sproutsocial.com/insights/tips-community-managers/)
- [Setting Up Community Guidelines for Moderation](https://www.getopensocial.com/blog/community-management/setting-community-guidelines-moderation)

[charter]: /sig-contributor-experience/charter.md#code-binaries-and-services
[moderators]: ./moderators.md
[Code of Conduct]: /code-of-conduct.md
[Kubernetes Community Values]: /values.md
[stages of burnout]: https://opensource.com/business/15/12/avoid-burnout-live-happy
[Code of Conduct Committee]: /committee-code-of-conduct/README.md
[Steering Committee]: https://git.k8s.io/steering
[moderator request]: https://github.com/kubernetes/community/issues/new/choose
[filing an issue]: https://github.com/kubernetes/community/issues/new/?assignees=&labels=area%2Fcommunity-management%2C+sig%2Fcontributor-experience&template=moderator_application.md&title=REQUEST%3A+New+moderator+for+%3Cyour-GH-handle%3E+of+%3Ck8s+property%3E
