## SUMMARY:

This document covers everything you need to know about the creation and retirement (“lifecycle”) of a special interest, or working group within the Kubernetes project. General project governance information can be found in the [steering committee repo].
Out of scope for this document: [subproject] creation.

[Creation]
[Retirement]

## [Creation]
### Prerequisites for a SIG
- [ ] Read [sig-governance.md]
- [ ] Ensure all SIG Chairs, Technical Leads, and other leadership roles are [community members]
- [ ] Send an email to the Steering Committee <steering@kubernetes.io> to scope the SIG and get provisional approval
- [ ] Look at the checklist below for processes and tips that you will need to do while this is going on. It's best to collect this information upfront so you have a smoother process to launch
- [ ] Follow the [SIG charter process] to propose and obtain approval for a charter
- [ ] Announce new SIG on dev@kubernetes.io

### Prerequisites for a WG
- [ ] Read [wg-governance.md]
- [ ] Ensure all WG Organizers, and other leadership roles are [community members]
- [ ] Send email to [dev@kubernetes.io] titled "WG-Creation-Request: WG Foo" with some of the questions answered from wg-goverance.md and wait for community discourse; ask for SIG sponsorship
- [ ] Do the first checklist item in the [GitHub] section below and add a row to the WG section:
  - [ ] Label with committee/steering and wait for a simple majority
  - [ ] Also add sponsoring SIG Chair/Tech Leads as approvers; you'll get this from the community email above
  - [ ] Place a `/hold` on it until the members that need to review have; a contributor experience member will do this for you if they don't see one already
- [ ] Send an email to the stakeholder SIG mailing lists and steering committee with the sigs.yaml pull request


### [GitHub]
- [ ] Submit a PR that will add rows to [sigs.yaml] using the [generator doc]; this will create README files and OWNERS_ALIASES files for your new directory in `kubernetes/community`
- You’ll need:
  - SIG Name
  - Directory URL
  - Mission Statement
  - Chair Information
  - Meeting Information
  - Contact Methods
  - Any SIG Stakeholders
  - Any Subproject Stakeholders
- [ ] Add SIG-related docs like charter.md, schedules, roadmaps, etc. to your new kubernetes/community/SIG-foo directory once the above PR is merged.
- [ ] Use the process in [labels.md] to obtain a label; read about our [GitHub management] services

### Communicate:
Each one of these has a linked canonical source guideline from set up to moderation and your role and responsibilities for each. We are all responsible for enforcing our [code of conduct].
- [ ] Read [moderation.md] and understand your role in keeping our community safe
- [ ] Create your mailing lists. [mailing-list-guidelines.md]
- Create one mailing list for your members and another for your chairs/leads 
- Example: kubernetes-[sig/wg]-foo@googlegroups.com and kubernetes-[sig/wg]-foo-leads@googlegroups.com
- The chairs/leads email will be used for activation of certain platforms (eg zoom)
- [ ] Request a slack channel. [slack-guidelines.md]
- [ ] Request a YouTube playlist link. [youtube-guidelines.md]
- [ ] Request a zoom account. [zoom-guidelines.md]

### Engage:
...as a chair/tech lead with other chairs/tech leads
- [ ] Add yourself to the [leads@kubernetes.io group]
- [ ] Join the #chairs-and-techleads slack channel

...with the community as part of [sig-governance.md]
- [ ] Create a shared calendar and schedule your weekly/biweekly/triweekly weeks [update meetings]
- This calendar creation process will allow all of your leads to edit SIG/WG Meetings. This is important as we all change jobs, email addresses, and take breaks from the project. Shared calendars will also provide consistency with contributors looking for your subproject meetings, office hours, and anything else that the SIG/WGs contributors should know about.

## [Retirement]
(merging or disbandment)
Sometimes it might be necessary to sunset a SIG or Working Group. SIGs/WGs may also merge with an existing SIG/WG if deemed appropriate, and would save project overhead in the long run. Working Groups in particular are more ephemeral than SIGs, so this process should be followed when the Working Group has accomplished it's mission.

### Prerequisites for SIG Retirement
- [ ] SIG’s retirement decision follows the decision making and communication processes as outlined in their charter

### Prerequisites for WG Retirement
- [ ] Have completed the mission of the WG or have another reason as outlined in [wg-governance.md]

### Steps:
- [ ] Send an email to dev@kubernetes.io and community@kubernetes.io alerting the community of your intentions to disband or merge. [example]
- This kicks off the process for Contributor Experience’s community managers who will reach out and set an issue against `kubernetes/community` with exact next steps covered below. We can help walk through this when you get there. Most of this is covered in the same creation communication docs as above.
- [ ] Archive the member and lead/chair mailing lists/[GoogleGroups]
- [ ] Check the [slack-guidelines.md] for latest process on archiving the slack channel
- [ ] Deactivate the zoom license
- [ ] Delete your shared SIG/WG calendar
- [ ] Ensure that the [youtube-guidelines.md] links are removed and you've uploaded all SIG/WG meetings to date
- [ ] Move the existing SIG directory into the archive in `kubernetes/community`
- [ ] GitHub archiving/removing/other transactions:
   - [ ] Move all appropriate github repositories to an appropriate archive or a repo outside of the Kubernetes org
   - [ ] Each subproject a SIG owns must transfer ownership to a new SIG, outside the project, or be retired
   - [ ] File an issue with kubernetes/org if there are multiple repos
   - [ ] Retire or transfer any test-infra jobs or testgrid dashboards, if applicable, owned by the SIG. Work with SIG-Testing on this.
   - [ ] Migrate/Remove/Deprecate any SIG/WG labels in labels.yaml; find instructions in [labels.md]
   - [ ] Remove or rename any GitHub teams that refer to the SIG
   - [ ] Update sigs.yaml to remove or rename

[steering committee repo]: https://github.com/kubernetes/steering
[subproject]: /governance.md#subprojects
[Creation]: #Creation
[Retirement]: #Retirement
[GitHub]: #GitHub
[labels.md]: https://git.k8s.io/test-infra/label_sync/labels.md
[sig-governance.md]: /committee-steering/governance/sig-governance.md
[SIG charter process]: /committee-steering/governance
[wg-governance.md]: /committee-steering/governance/wg-governance.md
[sigs.yaml]: /sigs.yaml
[generator doc]: /generator
[Kubernetes/Org]: https://github.com/kubernetes/org/issues/new/choose
[GitHub management]: /github-management
[code of conduct]: /code-of-conduct.md
[moderation.md]: /communication/moderation.md
[slack-guidelines.md]: /communication/slack-guidelines.md
[youtube-guidelines.md]: /communication/youtube/youtube-guidelines.md
[zoom-guidelines.md]: /communication/zoom-guidelines.md
[Thursday community updates]: /events/community-meeting.md
[example]: https://docs.google.com/document/d/1qZcAvuWBznR_oEaPWtwm7U4JNT91m8r9YOUvInU-src/edit#heading=h.jsw0l2t0ra8
[update meetings]: /communication/calendar-guidelines.md
[community members]: /community-membership.md
[mailing-list-guidelines.md]: /communication/mailing-list-guidelines.md
[leads@kubernetes.io group]: https://github.com/kubernetes/k8s.io/blob/main/groups/groups.yaml
