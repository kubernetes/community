# SIG Autoscaling Annual Report 2020

## Operational

* How are you doing with operational tasks in
[sig-governance.md](https://git.k8s.io/community/committee-steering/governance/sig-governance.md)?
  * Is your README accurate? have a CONTRIBUTING.md file?
    * Yes, [README.md](https://github.com/kubernetes/community/blob/master/sig-autoscaling/README.md) is accurate (other than one link)
    * We do not currently have a CONTRIBUTING.md in the community repo

  * All subprojects correctly mapped and listed in [sigs.yaml](https://git.k8s.io/community/sig-list.md)?
    * [Yes](https://github.com/kubernetes/community/tree/master/sig-autoscaling#subprojects)

  * What’s your meeting culture? Large/small, active/quiet, learnings? Meeting notes up to date? Are you keeping
  recordings up to date/trends in community members watching recordings?
    * The main SIG meeting has ~8 people on average. Meeting notes [are recorded](https://docs.google.com/document/d/1RvhQAEIrVLHbyNnuaT99-6u9ZUMp7BfkPupT2LAZK7w/edit#heading=h.mh9yomoq1p6v). Video recordings were not made for a number of years, however recording of SIG meetings will be resuming as of 2021/05/17.

* How does the group get updates, reports, or feedback from subprojects? Are there any springing up or being
retired? Are OWNERS.md files up to date in these areas?
  * The SIG maintains direct ownership of all of the listed subprojects, with all discussed in one weekly meeting.
  * Updates for cloud provider implementations for the cluster autoscaler are delegated to the owners of these, with work cutting across these communicated both through SIG meetings and asynchronously through Github Issues/PR mentions.

* When was your last monthly community-wide update? (provide link to deck and/or recording)
  * 2020/08/20 - [Recording here](https://youtu.be/oDL3Kp5-9eM?t=776)

## Membership

* Are all listed SIG leaders (chairs, tech leads, and subproject owners) active?
  * Yes - though some owners of cloud provider implementations for CA may no longer be active.
  * Currently no proactive process for pruning inactive members from OWNERs files.

* How do you measure membership? By mailing list members, OWNERs, or something else?
  * Active participants in SIG meetings and pull requests.

* How does the group measure reviewer and approver bandwidth? Do you need help in any area now? What are you doing about it?
  * No official process for measuring reviewer and approver bandwidth, stuck or controversial PRs are usually discussed at weekly Office Hours.
  * Reviewer bandwidth is occasionally a restricting factor as we have a small group of approvers, with a small number of experts in each of the SIGs three main areas of focus (Horizontal Pod Autoscaling, Cluster Autoscaler, Vertical Pod Autoscaler).

* Is there a healthy onboarding and growth path for contributors in your SIG?
What are some activities that the group does to encourage this? What programs are you participating in to grow contributors
throughout the contributor ladder?
  * We currently don't actively participate in any programs to grow contributors through the contributor ladder.
  * We do implement and encourage a [lighter weight proposals](https://github.com/kubernetes/autoscaler/pull/3914) process for the autoscaler repository to allow new users to propose functionality in a structured reviewable manner without the full weight of the KEP process.

* What programs do you participate in for new contributors?
  * As above, currently none.

* Does the group have contributors from multiple companies/affiliations? Can end users/companies contribute in some way that
they currently are not?
  * Yes, members of a number of companies (both vendor and end user) regularly contribute to all of the projects owned by the SIG, especially with Cluster Autoscaler's cloud provider implementations. The variety of approvers on core code however is significantly more narrow, with the majority of approvers across the three projects being from a single company.
  * Currently a large proportion of the issues in the `kubernetes/autoscaler` repository are around cloud provider specific behaviours, and increased support from Cloud Providers in these issues could potentially help increase the responsiveness in these issues significantly as the majority of the core owners of the repository are limited in the environments they can support.

## Current initiatives and project health

* What are initiatives that should be highlighted, lauded, shout out, that your group is proud of? Currently underway?
What are some of the longer tail projects that your group is working on?
  * Long term plans to graduate Metrics v2 beta API to GA - this requires providing time for a number of relatively new features in the API to mature first however.
  * [HPAScaleToZero](https://github.com/kubernetes/enhancements/pull/2022) is in alpha and thus isn't available to a wide swathe of users using managed services. This should be progressed to beta in the medium term, however this is in tension with our aim of graduating the existing v2 beta API to GA.
  * [Vertical Pod Autoscaler adding support for customised recommenders](https://github.com/kubernetes/autoscaler/pull/3914)
  * [Cluster Autoscaler adding support for gRPC custom cloud provider](https://github.com/kubernetes/autoscaler/pull/3140)

* Year to date KEP work review: What’s now stable? Beta? Alpha? Road to alpha?
  * As mentioned above the metrics v2 API has been beta for a significant period of time, with a number of non-trivial features added to it whilst in beta phase. Discussions were had this year about potentially graduating the API with only the oldest and most proven parts of it moving to GA, however this was dismissed as likely leading to a confusing end user experience.

* What areas and/or subprojects does the group need the most help with?
  * Issue triage and response has not been a focus for the SIG, with most attention of the current reviewer/approver pool going to review of PRs. This has led to a [trend of increasing number of open issues](https://k8s.devstats.cncf.io/d/22/open-issues-prs-by-milestone-and-repository?orgId=1&var-sig_name=All&var-milestone_name=All&var-repo_name=kubernetes%2Fautoscaler)

* What's the average open days of a PR and Issue in your group? / what metrics does your group care about and/or measure?
  * Time to merge for PRs [varies significantly from month to month](https://k8s.devstats.cncf.io/d/44/pr-time-to-approve-and-merge?orgId=1&from=1577865600000&to=1609488000000&var-period=m&var-repogroup_name=SIG%20Autoscaling&var-apichange=All&var-size_name=All&var-kind_name=All) with median ~20 hours. Currently not focusing on issue triage and this has resulted in a trend of increasing number of open issues in the `kubernetes/autoscaler` repository.