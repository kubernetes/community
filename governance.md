# Guidelines for repositories under Kubernetes github orgs (e.g, kubernetes, kubernetes-incubator)

All repositories under Kubernetes github orgs, such as kubernetes and kubernetes-incubator,
should follow the procedures outlined in the [incubator document](incubator.md).

# Kubernetes Organization Roles

The following is a list of roles that are currently assumed by different contributors:

- **[New Contributor](https://github.com/kubernetes/contrib/issues/1090)**: a
  couple of PRs
- **Contributor**: more than a couple of PRs (which could include documentation
  contributions as well as code)
- **Org Member**: active enough to be useful to assign issues to them and add
  them to a github team (e.g., for a SIG) for notification purposes; if they
  choose public membership, they get a badge on their github profile
- **kubernetes-collaborators**: "read access" to kubernetes repo; get a badge 
  on PR and issue comments; trusted enough to run tests on their PRs 
  automatically; can issue "@k8s-bot ok to test" for other contributors
- **Reviewer**: In some OWNERS file as a reviewer (in repos using the bot),
  assigned related PRs, assigned relevant test bugs; can champion incubator
  repos
- **Approver**: some OWNERS file as an approver; will be needed to get code
  merged
- **SIG Participant**: active in one or more areas of the project; wide 
  variety of roles are represented
- **SIG lead**: SIG organizer
- **Area/Component Owner**: design/proposal approval authority for some area 
  of the project, though escalation is still possible, and most beta/GA API 
  changes are vetted by the API owners
- **API Owners**: lead designers of the project, who are familiar with the 
  design, requirements, mechanics, conventions, style, scope, gotchas, etc. 
  of the API
- **Team Lead**: tech lead or manager of some team at some company working on 
  K8s; can influence priorities of their team members; pragmatically, 
  probably want label/assignment powers
- **Top-Level OWNERS**: de-facto project elders; technically can 
  approve virtually any PRs; can sponsor incubator repos
- **kubernetes-maintainers**: write access to repo (assign issues/PRs,
  add/remove labels and milestones, edit issues and PRs, edit wiki,
  create/delete labels and milestones); technically can lgtm any PR and cause it
  to be merged by the submit queue; expected to review PRs and fix bugs related
  to their domains
- **kubernetes-pm**: help to manage and maintain the project in
  ways other than just writing code (e.g. managing issues).
- **kubernetes-admin**: direct code write/merge access; for build cops and
  release czars only.
- **Build Cop**: ensure tests pass, submit queue is working, rollback PRs, 
  manually merge as necessary to fix build
- **User-Support Rotation**: answer questions on stackoverflow, googlegroups, 
  slack, twitter, etc. full time while on duty
- **Release Czar**: drive release
- **K8s Org Owner**: can create repos, do ~any github action; the number of
  owners shouldn't scale with the organization's growth, O(1), and optimally it
  should be less than 10 people who are very familiar with project workings and
  distributed across a few time zones and organizations The other repos will
  have distinct sets of people filling some of the above roles, also.
   
# Kubernetes SIG Governance

In order to standardize Special Interest Group efforts, create maximum transparency, and route contributors to the appropriate SIG, SIGs should follow the guidelines stated below:

* Meet regularly, at least for 30 minutes every 3 weeks, except November and December
* Keep up-to-date meeting notes, linked from the SIG's page in the community repo
* Announce meeting agenda and minutes after each meeting, on their SIG mailing list
* Record SIG meeting and make it publicly available
* Ensure the SIG's mailing list and slack channel are archived
* Report activity in the weekly community meeting at least once every 6 weeks
* Participate in release planning meetings and retrospectives, and burndown meetings, as needed
* Ensure related work happens in a project-owned github org and repository, with code and tests explicitly owned and supported by the SIG, including issue triage, PR reviews, test-failure response, bug fixes, etc. 
