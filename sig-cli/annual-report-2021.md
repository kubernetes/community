# SIG CLI Annual Report 2021

This report reflects back on CY 2020 and was written in March 2021.

## Operational

* How are you doing with operational tasks in [sig-governance.md](https://github.com/kubernetes/community/blob/master/committee-steering/governance/sig-governance.md)?

    * Is your README accurate? have a CONTRIBUTING.md file?
        * Yes, [README.md](https://github.com/kubernetes/community/blob/master/sig-cli/README.md) is accurate and up-to-date.
        * Our [CONTRIBUTING.md](https://github.com/kubernetes/community/blob/master/sig-cli/CONTRIBUTING.md) requires an update as a follow-up to our recent completion of the [move to staging](https://github.com/kubernetes/enhancements/issues/1020).

    * All subprojects correctly mapped and listed in sigs.yaml?
        * The current list is up-to-date, but this will soon require an update as soon as [KUI moves to kubernetes-sig](https://github.com/kubernetes/org/issues/2461).

    * What’s your meeting culture? Large/small, active/quiet, learnings? Meeting notes up to date? Are you keeping recordings up to date/trends in community members watching recordings?
        * We have two recurring meetings:
        1. Our main meeting happens biweekly, during those calls we always reserve time for important announcements, such as important release dates, recognitions, etc. After that we leave time for newcomers to introduce themselves. The remaining part of the call is devoted to discuss items which are in the agenda. We close the call with a round of standups from subprojects.
        2. Bug scrub every four weeks, during which we go through issues and pull requests, assigning them to contributors who are interested in working on them.
        All of the meetings are recorded and available online, all the meeting invites are up-to-date and present in community calendar.

* How does the group get updates, reports, or feedback from subprojects? Are there any springing up or being retired? Are OWNERS files up to date in these areas?
    * As mentioned above, all the subprojects have a spot to report their progress during standups at the end of our biweekly calls. OWNERS files are up-to-date.

* When was your last monthly community-wide update? (provide link to deck and/or recording)
    * Last presentation was on April 16, 2020:
    * [Slides](https://docs.google.com/presentation/d/1Y8SHFz6yyYS6rvRCgYUrSgF-moPk_YB6A-7Ykw5eWnU/edit#slide=id.g401c104a3c_0_)
    * [Recording](https://youtu.be/Y3z2grPHRh4?t=415)
    * [KubeCon NA 2020 recording](https://www.youtube.com/watch?v=gTzv6mpTYWw)

## Membership

* Are all listed SIG leaders (chairs, tech leads, and subproject owners) active?
    * Yes, all the lists are up-to-date.

* How do you measure membership? By mailing list members, OWNERs, or something else?
    * Anyone can be considered a SIG member if they join the Zoom calls or general discussions regularly.

* How does the group measure reviewer and approver bandwidth? Do you need help in any area now? What are you doing about it?
    * At our biweekly meetings (see [notes](https://docs.google.com/document/d/1r0YElcXt6G5mOWxwZiXgGu_X6he3F--wKwg-9UBc29I/edit)), we track the progress of each feature the SIG is working on. Each feature has dev leads assigned to it. If a developer does not have time to complete the task we try to find someone else who has the bandwidth or we defer the feature to the next release.
    * We periodically remove inactive reviewers and approvers, invite new contributors we have seen active in SIG meetings and activities to join the ranks.

* Is there a healthy onboarding and growth path for contributors in your SIG? What are some activities that the group does to encourage this? What programs are you participating in to grow contributors throughout the contributor ladder?
    * Over the past months the leads take a bit more aggressive approach to closely monitor contributions and recognize the most active members by promoting them to reviewers and eventually approvers.

* What programs do you participate in for new contributors?
    * We try to feature “good-first-issue”s.
    * As mentioned above, we give newcomers a spot to introduce at our biweekly calls.
    * KubeCon updates.
    * Individual/1-1 mentoring.
    * In 2020 we participated in Google Season of Docs which resulted in [new documentation for kubectl and kustomize](https://kubectl.docs.kubernetes.io/).
    * Meet our contributors.

* Does the group have contributors from multiple companies/affiliations? Can end users/companies contribute in some way that they currently are not?
    * The group has contributors from multiple companies/affiliations.
    * [28 companies contributed code in the last year](https://k8s.devstats.cncf.io/d/8/company-statistics-by-repository-group?orgId=1&var-period=y&var-metric=contributions&var-repogroup_name=SIG%20CLI&var-companies=All)

## Current initiatives and project health

* What are initiatives that should be highlighted, lauded, shout out, that your group is proud of? Currently underway? What are some of the longer tail projects that your group is working on?
    * [Moving kubectl package code to staging](https://github.com/kubernetes/enhancements/issues/1020)

* Year to date KEP work review: What’s now stable? Beta? Alpha? Road to alpha?
    * [Moving kubectl package code to staging](https://github.com/kubernetes/enhancements/issues/1020)
        * Our multi-year effort to split out of the main kubernetes repository.
    * [kubectl debug](https://github.com/kubernetes/enhancements/issues/1441) (beta)
    * Several smaller efforts to unify code across all the commands, and removing technical debt.

* What areas and/or subprojects does the group need the most help with?
    * Feature management - we're looking for a person who is familiar with managing feature delivery/product management in a broad sense. As mentioned before, the current process is that during our bi-weekly calls at the beginning of every release we write down planned features and a person responsible for delivering it. We would like to see a single person driving this effort and transforming that to a more asynchronous process. For example, gathering features through our mailing list and reporting progress during our bi-weekly calls.

* What's the average open days of a PR and Issue in your group? / what metrics does your group care about and/or measure?
    * Over the past year we've introduced monthly blocker bug scrubs which allowed us to shorten the average PR time to approve and merge by more than half, [from over a 8.5 days to 3 days](https://k8s.devstats.cncf.io/d/44/pr-time-to-approve-and-merge?orgId=1&from=1577865600000&to=1609488000000&var-period=y&var-repogroup_name=SIG%20CLI&var-apichange=All&var-size_name=All&var-kind_name=All).
