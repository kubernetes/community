# Community Maintenance 

This document outlines maintenance tasks that SIG Contributor Experience should be doing on a regular basis.
Copy this list into a new Github issue, for example "Community Maintenance Tasks for 2Q 2018". 
Then people can volunteer to audit different parts of the project at once. 
These tasks should be done at a _minimum of twice a year_, but ideally once a quarter.

If any of these tasks can be automated, then do so, however most of these require a human to make a judgement decision. 
If you find that any one person is in charge of a resource it is your responsibility to bring this issue to SIG Contributor Experience so that important parts of Kubernetes infrastructure are not assigned to one sole person. 

### Slack

- [ ] Channels - Anything out of place or not following the guidelines? 
  - [ ] Close unused channels. 
  - [ ] Check to see if there is a purpose, pinned documents like agendas, and other best practices being used in the channel. These are not required but are useful.
- [ ] Audit the Integrations and apps
- [ ] Audit the user tokens
- [ ] Audit the Emojis
- [ ] Generate a Slack data report to share with the community (optional)

### GitHub

- [ ] Audit Org admin permissions
- [ ] Audit Integrations
- [ ] Check Label usage (education, deprecation, etc)
- [ ] Is the PR and/or issue message still good/the best?
- [ ] Check kubernetes/community/.github/PULL_REQUEST_TEMPLATE
- [ ] Where’s the issue template located? 

### Calendars

- [ ] Check that the calendar works
  - Go to http://k8s.io/community and follow the flow
- [ ] Make sure all SIG, WG, and other community meetings are showing - especially check new groups that were created in the last quarter
- [ ] Make sure that invites have more than one owner, shared with the mailing list distro, and have an agenda attached plus zoom URL and telephone dial information.

### SIGS

- [ ] Ensure everyone in `sigs.yaml` is subscribed to sigs-leads mailing list
- [ ] Ensure that chairs and leads are accurate in the `sigs.yaml`

### YouTube

- [ ] Check sharing on all the playlists
  - Should only be shared with leads or people appointed by leads
  - Either remove other people or regenerate a new collaboration URL
- [ ] Check admin permissions on the Kubernetes account
- [ ] Fix any videos accidentally tagged as unlisted 
- [ ] Generate a YouTube traffic report to share with the community (optional)

### Zoom

- [ ] Ensure leads have access to their upgraded Zoom account
- [ ] If appropriate change the password if a SIG has changed leadership 
  - Contact CNCF to adjust upgraded licenses if necessary
- [ ] Audit kubernetes.io GSuite - this is primarily needed to create associated resources like GoogleGroups and the Google Cloud organization.  Accounts in  this domain cost the CNCF actual money, so are not available for individuals. Three steering committee members will hold the keys to three accounts on the domain, which can be used to create `@kubernetes.io` groups as needed.
  - Volunteers (Ask for new representatives from the steering committee if election turnover affects this list):
    - Joe Beda (@jbeda)
    - Brendan Burns (@brendandburns)
    - Tim Hockin (@thockin)

### Community Meeting

- [ ] Generate an archive of the [community meeting notes](https://docs.google.com/document/d/1VQDIAB0OqiSjIHI8AWMvSdceWhnz56jNpZrLs6o7NJY/edit#heading=h.2gp5yf2snwg5) for the quarter if necessary
  - [ ] Ensure the retrospective section has links to the retrospective for the releases
  - [ ] Ensure the demo section is filled out and demos are assigned dates
  - [ ] Ensure the [agenda template and hosting guide](https://docs.google.com/document/d/1g7fR5cvCGFq15SJ4iQMclbj0QIeREKu_QP8ftnSaJ4o/edit) are up to date 
- [ ] Ensure the [SIG Schedule](https://docs.google.com/spreadsheets/d/1adztrJ05mQ_cjatYSnvyiy85KjuI6-GuXsRsP-T2R3k/edit#gid=1543199895) have SIGs assigned for the next few months
  - [ ] Ensure the assigned SIGs in this sheet match the actual SIGs in sigs.yaml. Double check that a SIG hasn't missed an update this past cycle and if they have schedule them for an upcoming meeting

### Properties Managed by the CNCF

The CNCF provides support to Kubernetes for the following properties, you do not need to check them, however members of SIG Contributor Experience should have a working relationship with the administrators of the following properties:

- Twitter
- Blog
- GCP Organization

### Unknown State

Properties we should learn how to maintain, or at least document:

- Kubernetes.io DNS
- Kubernetes.io URL redirector
