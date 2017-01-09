### SIG creation procedure

* Propose the new SIG publicly, including a brief mission statement, by emailing kubernetes-dev@googlegroups.com and kubernetes-users@googlegroups.com, then wait a couple of days for feedback
* Create a group (see the detailed instructions below)
* Ask a repo maintainer to create a github label, if one doesn't already exist: sig/foo
* If you wish, request a new [kubernetes.slack.com](http://kubernetes.slack.com) channel (#sig-foo) from @sarahnovotny.  New users can join at [slack.kubernetes.io](http://slack.kubernetes.io).
* Organize video meetings as needed. No need to wait for the [Weekly Community Video Conference](community/README.md) to discuss. Please report summary of SIG activities there.
 * Request a Zoom account from @sarahnovotny if you expect more than 30 attendees or attendees from China.
 * Add the meeting to the community meeting calendar by inviting cgnt364vd8s86hr2phapfjc6uk@group.calendar.google.com.
* Use existing proposal and PR process
* Announce new SIG on kubernetes-dev@googlegroups.com and ask a repo maintainer to create a kubernetes/community directory and github team for the new group
* Submit a PR to add any SIG-related docs, schedules, roadmaps, etc. to your new kubernetes/community/SIG-foo directory.
* Slack activity is archived at [kubernetes.slackarchive.io](http://kubernetes.slackarchive.io).  To start archiving a new channel invite the slackarchive bot to the channel via `/invite @slackarchive`

#### 
**Google Groups creation**

Create Google Groups at [https://groups.google.com/forum/#!creategroup](https://groups.google.com/forum/#!creategroup), following the procedure: 

* Create a group using the name convention below;
* Group has to be created as an e-mail list with at least three owners (including sarahnovotny at google.com and ihor.dvoretskyi at gmail.com);
* To add the owners, visit the Group Settings (drop-down menu on the right side), select Direct Add Members on the left side and add Sarah and Ihor via email address (with a suitable welcome message); in Members/All Members select Ihor and Sarah and assign to an "owner role";
* Set "View topics", "Post", "Join the Group" permissions to be "Public"


Name convention:

* kubernetes-sig-foo-misc
* kubernetes-sig-foo-test-failures
* kubernetes-sig-foo-bugs
* kubernetes-sig-foo-feature-requests
* kubernetes-sig-foo-proposals
* kubernetes-sig-foo-pr-reviews
* kubernetes-sig-foo-api-reviews

Example:

* kubernetes-sig-onprem-misc
* kubernetes-sig-onprem-test-failures
* kubernetes-sig-onprem-bugs
* kubernetes-sig-onprem-feature-requests
* kubernetes-sig-onprem-proposals
* kubernetes-sig-onprem-pr-reviews
* kubernetes-sig-onprem-api-reviews

#### 
**GitHub users creation**

Create the GitHub users at [https://github.com/join](https://github.com/join), using the name convention below.

As an e-mail address, please, use the Google Group e-mail address of the respective Google Group, created before (i.e. - for user ‘k8s-sig-foo-misc’ use ‘[kubernetes-sig-foo-misc@googlegroups.com](mailto:kubernetes-sig-foo-misc@googlegroups.com)’). After creating the GitHub users, please, request @idvoretskyi (backup person - @sarahnovotny) to add these users to the Kubernetes organization.


Name convention:

* k8s-sig-foo-misc 
* k8s-sig-foo-test-failures
* k8s-sig-foo-bugs
* k8s-sig-foo-feature-requests
* k8s-sig-foo-proposals
* k8s-sig-foo-pr-reviews
* k8s-sig-foo-api-reviews

Example:

* k8s-sig-onprem-misc
* k8s-sig-onprem-test-failures
* k8s-sig-onprem-bugs
* k8s-sig-onprem-feature-requests
* k8s-sig-onprem-proposals
* k8s-sig-onprem-pr-reviews
* k8s-sig-onprem-api-reviews

#### 
**Create the GitHub teams**

Create the GitHub teams at [https://github.com/orgs/kubernetes/new-team](https://github.com/orgs/kubernetes/new-team), using the name convention below. Please, add the GitHub users (created before) to the GitHub teams respectively.

Name convention:

* sig-foo-misc 
* sig-foo-test-failures
* sig-foo-bugs
* sig-foo-feature-requests
* sig-foo-proposals
* sig-foo-pr-reviews
* sig-foo-api-reviews

Example:

* sig-onprem-misc
* sig-onprem-test-failures
* sig-onprem-bugs
* sig-onprem-feature-requests
* sig-onprem-proposals
* sig-onprem-pr-reviews
* sig-onprem-api-reviews
