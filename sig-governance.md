# SIG Governance

In order to standardize Special Interest Group efforts, create maximum transparency, and route contributors to the appropriate SIG, SIGs should follow the guidelines stated below:

* Meet regularly, at least for 30 minutes every 3 weeks, except November and December
* Keep up-to-date meeting notes, linked from the SIG's page in the community repo
* Announce meeting agenda and minutes after each meeting, on their SIG mailing list
* Record SIG meeting and make it publicly available
* Ensure the SIG's mailing list and slack channel are archived
* Report activity in the weekly community meeting at least once every 6 weeks
* Participate in release planning meetings and retrospectives, and burndown meetings, as needed
* Ensure related work happens in a project-owned github org and repository, with code and tests explicitly owned and supported by the SIG, including issue triage, PR reviews, test-failure response, bug fixes, etc. 
* Use the above forums as the primary means of working, communicating, and collaborating, as opposed to private emails and meetings
* Represent the SIG for the PM group - see [PM SIG representatives](https://github.com/kubernetes/community/blob/master/sig-product-management/SIG%20PM%20representatives.md).

## SIG roles
- **SIG Participant**: active in one or more areas of the project; wide 
  variety of roles are represented
- **SIG Lead**: SIG organizer

## SIG creation procedure

### Prerequisites

* Propose the new SIG publicly, including a brief mission statement, by emailing kubernetes-dev@googlegroups.com and kubernetes-users@googlegroups.com, then wait a couple of days for feedback
* Ask a repo maintainer to create a github label, if one doesn't already exist: sig/foo
* Request a new [kubernetes.slack.com](http://kubernetes.slack.com) channel (#sig-foo) from **@sarahnovotny**.  New users can join at [slack.kubernetes.io](http://slack.kubernetes.io).
* Slack activity is archived at [kubernetes.slackarchive.io](http://kubernetes.slackarchive.io).  To start archiving a new channel invite the slackarchive bot to the channel via `/invite @slackarchive`
* Organize video meetings as needed. No need to wait for the [Weekly Community Video Conference](community/README.md) to discuss. Please report summary of SIG activities there.
 * Request a Zoom account from pyao@linuxfoundation.org cc'ing sarahnovotny@google.com.
 * Add the meeting to the community meeting calendar by inviting cgnt364vd8s86hr2phapfjc6uk@group.calendar.google.com.
* Use existing proposal and PR process (to be documented)
* Announce new SIG on kubernetes-dev@googlegroups.com 
* Submit a PR to add a row for the SIG to the table in the kubernetes/community README.md file, to create a kubernetes/community directory, and to add any SIG-related docs, schedules, roadmaps, etc. to your new kubernetes/community/SIG-foo directory.

### **Creating service accounts for the SIG**

With a purpose to distribute the channels of notification and discussion of the various topics, every SIG has to use multiple accounts to GitHub mentioning and notifications. Below the procedure is explained step-by-step.

NOTE: This procedure is managed and maintained by **[@idvoretskyi](https://github.com/idvoretskyi)**; please, reach him directly in case of any questions/suggestions.

#### **Google Groups creation**

Create Google Groups at [https://groups.google.com/forum/#!creategroup](https://groups.google.com/forum/#!creategroup), following the procedure: 

* Each SIG should have one discussion groups, and a number of groups for mirroring relevant github notifications;
* Create groups using the name conventions below;
* Groups should be created as e-mail lists with at least three owners (including sarahnovotny at google.com and ihor.dvoretskyi at gmail.com);
* To add the owners, visit the Group Settings (drop-down menu on the right side), select Direct Add Members on the left side and add Sarah and Ihor via email address (with a suitable welcome message); in Members/All Members select Ihor and Sarah and assign to an "owner role";
* Set "View topics", "Post", "Join the Group" permissions to be "Public"

Name convention:

* kubernetes-sig-foo (the discussion group)
* kubernetes-sig-foo-misc
* kubernetes-sig-foo-test-failures
* kubernetes-sig-foo-bugs
* kubernetes-sig-foo-feature-requests
* kubernetes-sig-foo-proposals
* kubernetes-sig-foo-pr-reviews
* kubernetes-sig-foo-api-reviews

Example:

* kubernetes-sig-onprem
* kubernetes-sig-onprem-misc
* kubernetes-sig-onprem-test-failures
* kubernetes-sig-onprem-bugs
* kubernetes-sig-onprem-feature-requests
* kubernetes-sig-onprem-proposals
* kubernetes-sig-onprem-pr-reviews
* kubernetes-sig-onprem-api-reviews

#### **GitHub users creation**

Create the GitHub users at [https://github.com/join](https://github.com/join), using the name convention below.

As an e-mail address, please, use the Google Group e-mail address of the respective Google Group, created before (i.e. - for user ‘k8s-mirror-foo-misc’ use ‘[kubernetes-sig-foo-misc@googlegroups.com](mailto:kubernetes-sig-foo-misc@googlegroups.com)’). After creating the GitHub users, please, add these users to the Kubernetes organization. If you don't have enough permissions to do that (by default, you don't), please request **@idvoretskyi** (backup person - **@sarahnovotny**) to help you with this. If GitHub contacts you about having too many robot accounts, please let us know.


Name convention:

* k8s-mirror-foo-misc 
* k8s-mirror-foo-test-failures
* k8s-mirror-foo-bugs
* k8s-mirror-foo-feature-requests
* k8s-mirror-foo-proposals
* k8s-mirror-foo-pr-reviews
* k8s-mirror-foo-api-reviews

There is no need for a k8s-mirro-foo user.

Example:

* k8s-mirror-onprem-misc
* k8s-mirror-onprem-test-failures
* k8s-mirror-onprem-bugs
* k8s-mirror-onprem-feature-requests
* k8s-mirror-onprem-proposals
* k8s-mirror-onprem-pr-reviews
* k8s-mirror-onprem-api-reviews

NOTE: We have found that Github's notification autocompletion finds the users before the corresponding teams. This is the reason we recommend naming the users `k8s-mirror-foo-*` instead of `k8s-sig-foo-*`. If you previously created users named `k8s-sig-foo-*`, we recommend you rename them.

#### **Create the GitHub teams**

Create the GitHub teams at [https://github.com/orgs/kubernetes/new-team](https://github.com/orgs/kubernetes/new-team), using the name convention below. Please, add the GitHub users (created before) to the GitHub teams respectively.

Name convention:

* sig-foo-misc 
* sig-foo-test-failures
* sig-foo-bugs
* sig-foo-feature-requests
* sig-foo-proposals
* sig-foo-pr-reviews
* sig-foo-api-reviews

Note that there should not be a sig-foo team. We want to encourage contributors to select the most appropriate team to notify.

Example:

* sig-onprem-misc
* sig-onprem-test-failures
* sig-onprem-bugs
* sig-onprem-feature-requests
* sig-onprem-proposals
* sig-onprem-pr-reviews
* sig-onprem-api-reviews
