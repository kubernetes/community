# SIG Instrumentation Annual Report 2020

This report reflects back on CY 2020 and was written in Feb-Mar. 2021.

## Operational

*   How are you doing with operational tasks in [sig-governance.md](https://github.com/kubernetes/community/blob/master/committee-steering/governance/sig-governance.md)?
    *   Is your README accurate? have a CONTRIBUTING.md file?
        *   Yes. We do not require a separate CONTRIBUTING.md file.
    *   All subprojects correctly mapped and listed in [sigs.yaml](https://github.com/kubernetes/community/blob/master/sig-list.md)?
        *   Yes.
    *   What’s your meeting culture? Large/small, active/quiet, learnings? Meeting notes up to date? Are you keeping recordings up to date/trends in community members watching recordings?
        *   We have lively biweekly regular meetings and biweekly triage sessions. Our recordings are relatively up to date and we receive requests for them so we believe they are being watched. Our community archives are all up to date.
*   How does the group get updates, reports, or feedback from subprojects? Are there any springing up or being retired? Are OWNERS files up to date in these areas?
    *   We request these in SIG meetings and asynchronously for KubeCon and other community updates.
    *   See above for OWNERS updates.
*   Same question as above but for working groups.
    *   N/A
*   When was your last monthly community-wide update? (provide link to deck and/or recording)
    *   Oct. 2020 [https://groups.google.com/g/kubernetes-dev/c/x1xr7bSuzv8/m/Oz49ZJgRBgAJ](https://groups.google.com/g/kubernetes-dev/c/x1xr7bSuzv8/m/Oz49ZJgRBgAJ)

## Membership

*   Are all listed SIG leaders (chairs, tech leads, and subproject owners) active?
    *   [https://k8s.devstats.cncf.io/d/13/developer-activity-counts-by-repository-group?orgId=1&var-period_name=Last%20quarter&var-metric=contributions&var-repogroup_name=All&var-country_name=All](https://k8s.devstats.cncf.io/d/13/developer-activity-counts-by-repository-group?orgId=1&var-period_name=Last%20quarter&var-metric=contributions&var-repogroup_name=All&var-country_name=All)
    *   All active in last quarter
*   How do you measure membership? By mailing list members, OWNERs, or something else?
    *   [https://github.com/kubernetes/org/blob/main/config/kubernetes/sig-instrumentation/teams.yaml](https://github.com/kubernetes/org/blob/main/config/kubernetes/sig-instrumentation/teams.yaml)
    *   We have three teams: leads (chairs/TLs), approvers, and members
    *   We annually remove inactive members (using the same criteria used for voting) and invite members to join that we have seen active in SIG meetings and activities
*   How does the group measure reviewer and approver bandwidth? Do you need help in any area now? What are you doing about it?
    *   Approver: we have limited areas that need our approval so we are ok on bandwidth
    *   Review: distributed around SIG via biweekly triage meetings
*   Is there a healthy onboarding and growth path for contributors in your SIG? What are some activities that the group does to encourage this? What programs are you participating in to grow contributors throughout the contributor ladder?
    *   We would like to grow more contributors into SIG leadership roles for sustainability.
    *   We invite newbies to come to triage meetings to learn how to review and triage code.
    *   We go over all features as a SIG in our regular meetings.
    *   We do not have a lot of code to own for people to progress on our SIG ladder but we are open to adding new approvers for people that put in the work.
*   What programs do you participate in for new contributors?
    *   Individual/1-1 mentoring
    *   We try to feature “good-first-issue”s and raise them on k-dev (e.g. structured logging)
*   Does the group have contributors from multiple companies/affiliations? Can end users/companies contribute in some way that they currently are not?
    *   Yes: Google, Red Hat, and many first contributors across multiple companies
        *   Cross referencing [PRs by contributor + company](https://k8s.devstats.cncf.io/d/66/developer-activity-counts-by-companies?orgId=1&var-period_name=Last%20quarter&var-metric=contributions&var-repogroup_name=All&var-country_name=All&var-companies=All) with our [members](https://github.com/kubernetes/org/blob/main/config/kubernetes/sig-instrumentation/teams.yaml#L20), we have 6 companies with contributions: Google, Red Hat, Polar Signals, Huawei, Salesforce, and ByteDance.  The vast (80%) majority of our contributions from members in the past quarter are from Google and Red Hat.
        *   In [only sig-instrumentation-owned repositories](https://k8s.devstats.cncf.io/d/66/developer-activity-counts-by-companies?orgId=1&var-period_name=Last%20quarter&var-metric=contributions&var-repogroup_name=SIG%20Instrumentation&var-country_name=All&var-companies=All), there are 18 contributors with >10 contributions in the last quarter, from 14 different companies.

## Current initiatives and project health

*   What are initiatives that should be highlighted, lauded, shout out, that your group is proud of? Currently underway? What are some of the longer tail projects that your group is working on?
    *   [Structured Logging](https://github.com/kubernetes/enhancements/issues/1602)
    *   [Tracing](https://github.com/kubernetes/enhancements/issues/647)
    *   [Reducing metrics exposed by the kubelet](https://github.com/kubernetes/kubernetes/issues/68522?notification_referrer_id=MDE4Ok5vdGlmaWNhdGlvblRocmVhZDM3ODU4MzI4NTo0NzA2MTMx&notifications_query=is%3Asaved#issuecomment-715462010)
*   Year to date KEP work review: What’s now stable? Beta? Alpha? Road to alpha?
    *   [Structured logging Alpha in 1.19](https://github.com/kubernetes/enhancements/issues/1602)
        *   Flag to enable structured logs: --logging-format=json
        *   For beta: perf tests, continued migration, verify decisions through one complete component migrated
    *   [Defend Against Logging Secrets via Static Analysis: Alpha in 1.20](https://github.com/kubernetes/enhancements/issues/1933)
        *   Test static analysis against sampled PRs using prow
    *   [Pod Resource Metrics Alpha in 1.20](https://github.com/kubernetes/enhancements/issues/1748)
        *   Pod Resource Metrics representing the view of the scheduler
    *   [Log sanitization Alpha 1.20](https://github.com/kubernetes/enhancements/issues/1753)
        *   Allow filtering out known sensitive data via flag --logging-sanitization
    *   [kube-state-metrics](https://github.com/kubernetes/kube-state-metrics/) is on its way to a new major breaking v2.0.0 release, with v2.0.0-rc.0 being released recently.
    * custom-metrics-apiserver is shepherding a new initiative to support multiple custom and external metrics servers. [kubernetes-sigs/custom-metrics-apiserver#70](https://github.com/kubernetes-sigs/custom-metrics-apiserver/issues/70)
    * prometheus-adapter recently joined the kubernetes-sigs organization and is still being integrated. The next important step for the project is its complete overhaul to improve its scability and make its configuration [more user friendly](https://github.com/s-urbaniak/prometheus-adapter/blob/master/design.md).
    *   Next:
        *   Promote *above*
        *   [APIServer tracing](https://github.com/kubernetes/enhancements/issues/647)
        *   [Dynamic cardinality enforcement](https://github.com/kubernetes/enhancements/issues/2305)
        *   [Metrics stability GA](https://github.com/kubernetes/enhancements/issues/1209)
*   What areas and/or subprojects does the group need the most help with?
    *   Structured Logging
    *   [Promq](https://github.com/kubernetes-sigs/instrumentation-tools/tree/master/promq)
*   What's the average open days of a PR and Issue in your group? / what metrics does your group care about and/or measure?
    *   [devstats: time to approve and merge](https://k8s.devstats.cncf.io/d/44/pr-time-to-approve-and-merge?orgId=1&from=1577865600000&to=1609488000000&var-period=m&var-repogroup_name=SIG%20Instrumentation&var-apichange=All&var-size_name=All&var-kind_name=All) \
Varies from month to month, median time to merge is ~15 hours.  85th percentile is 1-2 weeks. \
However, we don’t guarantee PRs will be merged, so the time to merge isn’t very meaningful. As well, this data only covers SIG repos and not work covered in kubernetes/kubernetes.  As a SIG, we ensure PRs are triaged and reviewed through the bi-weekly triage meeting in all org:kubernetes components. kubernetes-sigs components are triaged by their respective owners.
