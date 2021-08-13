# 2020 SIG Multicluster Annual Report

This report reflects back on CY 2020 and was written in March 2021.

## Operational

How are you doing with operational tasks in sig-governance.md?

* Is your README accurate? have a CONTRIBUTING.md file?
  * Yes, our README is accurate: https://github.com/kubernetes/community/blob/master/sig-multicluster/README.md
  * Our CONTRIBUTING.md file is up to date: https://github.com/kubernetes/community/blob/master/sig-multicluster/CONTRIBUTING.md

* All subprojects correctly mapped and listed in sigs.yaml?
  * Yes

* What’s your meeting culture? Large/small, active/quiet, learnings? Meeting notes up to date? Are you keeping recordings up to date/trends in community members watching recordings?
  * SIG Multicluster holds regular weekly meetings which are normally well attended with ~10 participants.
  * Due to the increase in activity in 2020, SIG meetings went from bi-weekly to weekly and almost always have a solid agenda.
    * We attribute this to a shift in approach to focus on smaller, more clearly defined problems (Cluster ID, Multi-Cluster Services, Work) as well as calls for participation at KubeCon.
  * Meetings are recorded and kept up to date.
  * We'd love to see more participation, more contributors bringing their perspectives and use cases. Specifically:
    * Help drive KEPs to GA, filling gaps in documentation.
    * Example/test implementations of our defined APIs.
    * What other problems are uncovered in the multi-cluster space now that e.g. connecting services across clusters is easier?
* How does the group get updates, reports, or feedback from subprojects? Are there any springing up or being retired? Are OWNERS.md files up to date in these areas?
  * Most group updates are at the weekly SIG meeting, big changes/ideas are shared with the list.
* Same question as above but for working groups.
  * Working groups join SIG weekly to discuss major updates, we don’t have separate meetings. Most WG discussion is on slack.
* When was your last monthly community-wide update? (provide link to deck and/or recording)
  * Our last community wide update was 2020-09-17: [SIG MC Community Update Slides](https://docs.google.com/presentation/d/1WKtsiSQn0sQ3IaSql4pGnH8Qt9ibBjHXMFBYZBhWeII).

## Membership
* Are all listed SIG leaders (chairs, tech leads, and subproject owners) active?
  * Both SIG chairs are active and attend most SIG meetings.
  * Kubefed owners are active on slack and GitHub issues.
* How do you measure membership? By mailing list members, OWNERs, or something else?
  * Mailing list members for general interest. Slack, weekly meeting attendance, and messages to the list for active engagement.
* How does the group measure reviewer and approver bandwidth? Do you need help in any area now? What are you doing about it?
  * The SIG tracks progress at weekly meetings and tends to closely watch and report on PRs and progress as a group so reviewer/approver bandwidth is not an issue.
  * Kubefed tends to operate independently. Reviewer bandwidth has not been an issue, questions on PRs and Slack tend to get relatively quick response.
* Is there a healthy onboarding and growth path for contributors in your SIG? What are some activities that the group does to encourage this? What programs are you participating in to grow contributors throughout the contributor ladder?
  * This SIG is in growth mode, and tries to create room for new ideas and projects driven by new and existing members alike. In 2020 we've seen two new initiatives kick off driven by relatively new members (Work API, Cluster ID).
  * We are actively soliciting new ideas and looking for new people to drive them.
  * We ask for new members to add their thoughts to docs and PRs, show of their prototypes, and present ideas and problems to the SIG at weekly meetings or via the list.
* What programs do you participate in for new contributors?
  * At Kubecon we use the contributor track to give insight into what the SIG is working on and ask for new contributors. We've seen growth in weekly meeting attendance and new presenters as a result.
  * When in-person Kubecons return we'll have Meet & Greets for new contributors (and everyone we didn't see in 2020!)
* Does the group have contributors from multiple companies/affiliations? Can end users/companies contribute in some way that they currently are not?
  * Yes, our chairs are from two different companies. We have regular participation in SIG projects from those companies and several others. We'd like to see additional users and companies bring their use cases and contribute to solutions.

## Current initiatives and project health
* What are initiatives that should be highlighted, lauded, shout out, that your group is proud of? Currently underway? What are some of the longer tail projects that your group is working on?
  * [Kubefed](http://sigs.k8s.io/kubefed) - long running subproject with active contribution. The big initiative for 2020 was to [kick off support for pull-based reconciliation](https://github.com/kubernetes-sigs/kubefed/blob/master/docs/keps/20200619-kubefed-pull-reconciliation.md).
  * [Work API](https://docs.google.com/document/d/1cWcdB40pGg3KS1eSyb9Q6SIRvWVI8dEjFp9RI0Gk0vg) - [sigs.k8s.io/work-api](https://sigs.k8s.io/work-api)
  * MCS API [KEP-1645](https://github.com/kubernetes/enhancements/tree/master/keps/sig-multicluster/1645-multi-cluster-services-api) - [sigs.k8s.io/mcs-api](https://sigs.k8s.io/mcs-api)
  * Cluster ID [KEP-2149](https://github.com/kubernetes/enhancements/tree/master/keps/sig-multicluster/2149-clusterid)
* Year to date KEP work review: What’s now stable? Beta? Alpha? Road to alpha?
  * Kubefed: nearly Beta. Expected in Q2 2021.
  * Work API: nascent, on the road to alpha.
  * MCS API: nearly Beta. Expected in Q2 2021, following Cluster ID alpha.
  * Cluster ID: nascent, on the road to alpha. Ready to implement in Q2 2021.
* What areas and/or subprojects does the group need the most help with?
  * SIG MC needs help with all of the above - especially when it comes to use cases and validating our approaches for different environments and deployment models.
* What's the average open days of a PR and Issue in your group? / what metrics does your group care about and/or measure?
  * According to recent [dev stats](https://k8s.devstats.cncf.io/d/44/pr-time-to-approve-and-merge?orgId=1&var-period=d7&var-repogroup_name=SIG%20Multicluster&var-apichange=All&var-size_name=All&var-kind_name=All)
    * Open to LGTM: avg 2.94 day, max 7.96 week
    * LGTM to approve: avg 3.91 hour, max 2.39 week
    * Approve to merge: avg 1.52 day, max 7.72 week
    * 85% Open to LGTM: avg 1.28 week, max 10.36 week
    * 85% LGTM to approve: avg 17.38 hour, max 2.39 week
    * 85% Approve to merge: avg 5.97 day, max 10.36 week
  * A lot of our focus has been on KEPs, which tend to have longer review cycles. We hope these timelines will contract when we shift to implementation.
    