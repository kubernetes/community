# Kubernetes SIG Service Catalog - 2020 Annual report
Konstantin Semenov, Jonathan Berkhahn


[Source](https://github.com/kubernetes/community/blob/master/committee-steering/governance/annual-reports.md)

## Checklist
- [x] Read about the process [here](https://git.k8s.io/community/committee-steering/governance/annual-reports.md#reporting-process)
- [ ] Copy this template into a new document and share with your mailing list/slack channel/meeting on whatever platform (gdocs, hackmd, etc.) that the team prefers.
- [x] Remove sections that are not applicable (example: if you are a working group, delete the special interest group questions)
- [ ] Pick graphs from Devstats to pull supporting data for your responses.
- [ ] Schedule a time with your Steering liaison and other Chairs, TLs, and Organizers of your group to check-in on your
  roles as Chair or Working Group Organizer.
  If anyone would rather meet 1:1, please have them reach out to the liaison directly, we are happy to.
  We’d like to talk about: challenges, wins, things you didn’t know before but wish you did, want to continue in the
  role or help finding a replacement; and lastly any feedback you have for us as a body and how we can help you
  succeed and feel comfortable in these leadership roles.
- [x] PR this document into your community group directory in kubernetes/community (example: sig-architecture/)
    - [x] by March 8th, 2021
    - [x] titled: annual-report-YEAR.md
- [x] are there any responses that you’d like to share privately first? steering-private@kubernetes.io or tag your liaison in for discussion.

## Operational
1. How are you doing with operational tasks in SIG-governance.md?
    1. Is your README accurate? have a CONTRIBUTING.md file?

       Yes, the [README](https://github.com/kubernetes/community/blob/master/sig-service-catalog/README.md) is accurate.
    2. All subprojects correctly mapped and listed in sigs.yaml?

       Yes, our [subprojects](https://github.com/kubernetes/community/blob/master/sig-service-catalog/README.md#subprojects) are current.
    3. What’s your meeting culture? Large/small, active/quiet, learnings? Meeting notes up to date?

       Are you keeping recordings up to date/trends in community members watching recordings?

       We have a low profile bi-weekly meeting, with [notes and agenda up to date](https://docs.google.com/document/d/17xlpkoEbPR5M6P5VDzNx17q6-IPFxKyebEekCGYiIKM/edit#).
       [Meeting recordings](https://www.youtube.com/watch?v=ukPj1sFFkr0&list=PL69nYSiGNLP2k9ZXx9E1MvRSotFDoHUWs) have not been uploaded since late 2019.

2. How does the group get updates, reports, or feedback from subprojects?
   Are there any springing up or being retired? Are OWNERS.md files up to date in these areas?

   We get updates on an ad-hoc basis.
   Minibroker has gone GA in September 2020, and we are looking forward to adopting a new subproject.
   
3. When was your last public community-wide update? (provide link to deck and/or recording)
   
    August 2020:  [video](https://youtu.be/ieXSPmdsQ6I) and [slides](https://static.sched.com/hosted_files/kccnceu20/ad/Aug18_SIG_Service_Catalog_Update.pdf)

## Membership
1. Are all listed SIG leaders (chairs, tech leads, and subproject owners) active?

   Yes.

2. How do you measure membership? By mailing list members, OWNERs, or something else?

   We don’t measure membership.

3. How does the group measure reviewer and approver bandwidth?
   Do you need help in any area now? What are you doing about it?

   We are a bit slow with reviewing and integrating PRs, but as we have recently fixed our CI pipelines, the process should become smoother. 
   While we would be happy to see developers move up the ladder, we don’t see a pressing need to adjust the current rate.

   We perform triage approximately once a week, and our [issue open/close rates are holding steady](https://k8s.devstats.cncf.io/d/39/issues-opened-closed-by-sig?orgId=1&var-period=d7&var-sig_name=service-catalog&var-kind_name=All).

4. Is there a healthy onboarding and growth path for contributors in your SIG? What are some activities that the group
   does to encourage this? What programs are you participating in to grow contributors throughout the contributor ladder?

   We accept patches from contributors, and are looking forward to the donation of a new sub-project.

5. What programs do you participate in for new contributors?

   We don’t participate in any particular programs.
   We find new contributors via slack, PRs, and issues.

6. Does the group have contributors from multiple companies/affiliations?
   Can end users/companies contribute in some way that they currently are not?

   Yes, there are contributors from [multiple companies](https://k8s.devstats.cncf.io/d/74/contributions-chart?orgId=1&var-period=d7&var-metric=contributions&var-repogroup_name=SIG%20Service%20Catalog&var-country_name=All&var-company_name=All&var-company=all).
   We see all sorts of contributions, varying from issues, to comments, to PRs and sig meeting participation.


## Current initiatives and project health
1. What are initiatives that should be highlighted, lauded, shout outs, that your group is proud of? Currently underway?
   What are some of the longer tail projects that your group is working on?
   
Completed in 2020:
   - [minibroker](https://github.com/kubernetes-sigs/minibroker) has gone GA in September 2020, and has released two minor versions since 1.0 
   - [migrated helm charts](https://github.com/kubernetes-sigs/service-catalog/issues/2841)
   
Ongoing:
   - [support the latest OSBAPI spec](https://github.com/kubernetes-sigs/go-open-service-broker-client/issues/165)
   - [move to v1 CustomResourceDefinition](https://github.com/kubernetes-sigs/service-catalog/issues/2809)

    
2. Year to date KEP work: What's now stable? Beta? Alpha? Road to alpha?
   
    Since our sig is an extension project, we aren't involved with KEPs directly.

3. What initiatives are you working on that aren't being tracked in KEPs?

   We are working on mitigating the impact of removing beta APIs in 1.22.

4. What areas and/or subprojects does the group need the most help with?

   Right now we are investigating intermittent failures with our integration tests that seem to produce false failures from time to time.
   We could always use extra hands to help with issues in our backlog.

5. What metrics/community health stats does your group care about and/or measure? Examples?

   On the technical health of the SIG, we look at
    - the ratio of open/close PRs
    - the ratio of open/close Issues
    - overall age of open Issues
    - diverse representation of companies in the sig participants