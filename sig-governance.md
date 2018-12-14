# SIG Governance

In order to standardize Special Interest Group efforts, create maximum transparency, and route contributors to the appropriate SIG, SIGs should follow the guidelines stated below:

* Create a charter and have it approved according to the [SIG charter process]
* Meet regularly, at least for 30 minutes every 3 weeks, except November and December
* Keep up-to-date meeting notes, linked from the SIG's page in the community repo
* Announce meeting agenda and minutes after each meeting, on their SIG mailing list
* Record SIG meeting and make it publicly available
* Ensure the SIG's mailing list and slack channel are archived
* Report activity in the weekly community meeting at least once every 6 weeks
* Participate in release planning meetings and retrospectives, and burndown meetings, as needed
* Ensure related work happens in a project-owned github org and repository, with code and tests explicitly owned and supported by the SIG, including issue triage, PR reviews, test-failure response, bug fixes, etc.
* Use the above forums as the primary means of working, communicating, and collaborating, as opposed to private emails and meetings

In addition, SIGs have the following responsibilities to SIG PM:
* identify SIG annual roadmap
* identify all SIG features in the current release
* actively track / maintain SIG features within [k/features](https://github.com/kubernetes/features)
* attend [SIG PM](/sig-pm/README.md) meetings, as needed / requested

[SIG charter process]: /committee-steering/governance/README.md

## SIG Roles

Defining SIG Roles is a function of the SIG Charter.
Guidelines for drafting a SIG Charter can be found [here](/committee-steering/governance/README.md).

## SIG creation and maintenance procedure

### Prerequisites

* Work with the Steering Committee to scope the SIG and get provisional approval.
  Follow the [SIG charter process] to propose and obtain approval for a charter.
* Ask a repo maintainer to create a github label, if one doesn't already exist: sig/foo
* Request a new [kubernetes.slack.com](http://kubernetes.slack.com) channel (#sig-foo) from [@parispittman](https://github.com/parispittman) or [@castrojo](https://github.com/castrojo).  New users can join at [slack.kubernetes.io](http://slack.kubernetes.io).
* Organize video meetings as needed. No need to wait for the [Weekly Community Video Conference](community/README.md) to discuss. Please report summary of SIG activities there.
 * Request a Zoom account by emailing Paris Pittman(`parispittman@google.com`) and Jorge Castro(`jorgec@vmware.com`). You must set up a google group (see below) for the SIG leads so that all the SIG leads have the ability to reset the password if necessary.
 * Read [how to use YouTube](/communication/K8sYoutubeCollaboration.md) for publishing your videos to the Kubernetes channel.
 * Calendars
   1. Create a calendar on your own account. Make it public.
   2. Share it with all SIG leads with full ownership of the calendar - they can edit, rename, or even delete it.
   3. Share it with `sc1@kubernetes.io`, `sc2@kubernetes.io`, `sc3@kubernetes.io`, with full ownership. This is just in case SIG leads ever disappear.
   4. Share it with the SIG mailing list, lowest privileges.
   5. Share individual events with `cgnt364vd8s86hr2phapfjc6uk@group.calendar.google.com` to publish on the universal calendar.
* Use existing proposal and PR process (to be documented)
* Announce new SIG on kubernetes-dev@googlegroups.com
* Leads should [subscribe to the kubernetes-sig-leads mailing list](https://groups.google.com/forum/#!forum/kubernetes-sig-leads)
* Submit a PR to add a row for the SIG to the table in the kubernetes/community README.md file, to create a kubernetes/community directory, and to add any SIG-related docs, schedules, roadmaps, etc. to your new kubernetes/community/SIG-foo directory.

### Discussion Platforms

Your SIG needs a place to discuss topics asynchronously. You have two options, a traditional mailing list via Google Groups, or a category on [discuss.kubernetes.io](discuss.kubernetes.io). The main difference is Groups is primarily email-based with a web UI tacked on, and Discuss is primarily a Web UI with email tacked-on. The other difference is that your SIG/WG is responsible for moderating your Google Group; with discuss you just depend on the usual community moderation.

- Working Groups, due to their temporary nature, are strongly encouraged to consider using an existing SIG mailing list if appropriate, otherwise use a discuss category for less management overhead.
- SIGs, due to their usage of calendars, and Zoom accounts, are strongly encouraged to use a traditional mailing list.

Choose one:

#### Create a Category

Post a message asking for a category in the [Site Feedback and Help](https://discuss.kubernetes.io/c/site-feedback) section and a moderator will create your category for you and provide you with a URL and mail address to post to.

#### Creating a Google Group

Create a Google Group at [https://groups.google.com/forum/#!creategroup](https://groups.google.com/forum/#!creategroup), following the procedure:

Each SIG must have two discussion groups with the following settings.

- kubernetes-sig-foo (the discussion group):
  - Anyone can view content
  - Anyone can join
  - Anyone can post
  - Only members can view the list of members
- kubernetes-sig-foo-leads (list for the leads, to be used with Zoom and Calendars)
  - Only members can view group content
  - Anyone can apply to join
  - Anyone can post
  - Only members can view the list of members
- Groups should be created as e-mail lists with at least three owners (including parispittman at google.com and ihor.dvoretskyi at gmail.com);
- To add the owners, visit the Group Settings (drop-down menu on the right side), select Direct Add Members on the left side and add Paris and Ihor via email address (with a suitable welcome message); in Members/All Members select Ihor and Paris and assign to an "owner role"
- Set "View topics", "Post", "Join the Group" permissions to be "Public"

## SIG/WG Retirement

Sometimes it might be necessary to sunset a SIG or Working Group.
SIGs/WGs may also merge with an existing SIG/WG if deemed appropriate, and would save project overhead in the long run.
Working Groups in particular are more ephemeral than SIGs, so this process should be followed when the Working Group has accomplished it's mission.

- Retiring a SIG for is covered in the [SIG Governance](https://github.com/kubernetes/community/blob/master/committee-steering/governance/sig-governance.md)
- Retiring a Working Group is covered in [WG Governance](https://github.com/kubernetes/community/blob/master/committee-steering/governance/wg-governance.md)

The process for closing a SIG/WG is as follows:

- SIG Chairs agree to disband. This decision should follow the decision making process of the SIG's Charter.
- Send a email to kubernetes-dev to let people know the SIG has either closed or merged with another SIG. This will let SIG Contributor Experience know that they need to help you archive/deactivate project resources.
  - Consider sending a [closing summary](https://docs.google.com/document/d/1qZcAvuWBznR_oEaPWtwm7U4JNT91m8r9YOUvInU-src/edit#heading=h.jsw0l2t0ra8) to the list. 
- Work with SIG Contributor Experience to:
  - Archive the mailing list/group
  - Archive the leads mailing list/group
  - Archive the slack channel
  - Deactivate the group's Zoom license
  - Move all appropriate github repositories to an appropriate archive or a repo outside of the Kubernetes org
    - Each subproject a SIG owns must transfer ownership to a new SIG, outside the project, or be retired
    - File an issue with kubernetes/org if multiple repos that need to be retired
  - Coordinate with SIG Testing on the following topics (if necessary) 
    - Retire or transfer any test-infra jobs owned by the SIG
    - Retire or transfer any testgrid dashboards owned by the SIG
  - Clean up and remove all GitHub teams that refer to that SIG/WG
  - Migrate/Remove/Deprecate any SIG/WG labels in [labels.yaml](https://git.k8s.io/test-infra/label_sync/labels.yaml)
  - Ensure that the YouTube Collaboration links are removed
- Remove SIG Calendar and events from the community calendar
- Update `sigs.yaml` to reflect the removal of the SIG/WG
