# 2021 Annual Report: SIG Architecture

## Current initiatives

1. What work did the SIG do this year that should be highlighted?
   
   SIG-Architecture operates mostly through its subprojects like conformance, enhancements, code organization and production readiness.
   The regular SIG meetings have turned into a policy making/adjusting forum and for tackling issues that are across various SIGs.
   The focus has been to increase reliability, make things stable and document things that we hadn't before. 
   Closing conformance gaps!
   We haven't had to update architecture very much. Stability improved.
   increased reliability, Stabilizing, version skews etc have helped the community
     - ecosystem balance shifted from using beta APIs primarily to using stable APIs primarily (1.22+)

2. What initiatives are you working on that aren't being tracked in KEPs?
   - KEP survey for enhancements
   - Working with dependencies to cut down indirect go.sum entries.
   - Updating API conventions by adding documentation

3. KEP work in 2021 (1.21, 1.22, 1.23):

   - Stable
     - [2527 - Clarify meaning of status](https://github.com/kubernetes/enhancements/blob/master/keps/sig-architecture/2527-clarify-status-observations-vs-rbac/README.md) - 1.22
     - [1693 - Warning API mechanism](https://github.com/kubernetes/enhancements/blob/master/keps/sig-api-machinery/1693-warnings/README.md) - 1.22
   - Beta
     - [2572 - Defining the Kubernetes release cadence](https://github.com/kubernetes/enhancements/blob/master/keps/sig-release/2572-release-cadence/) - 1.23

## Project health

1. What areas and/or subprojects does your group need the most help with?
   Any areas with 2 or fewer OWNERs? (link to more details)

   - PRR we need to grow more reviewers and approvers
   - API reviews has a pipeline set up with Jordan leading. Need more SIGs to identify additional folks to start learning
   - Conformance, all the easy stuff is done. So we have harder stuff to review, so need additional folks with context to help make progress

2. What metrics/community health stats does your group care about and/or measure?

   - We care about regressions, backports, dependencies across release boundaries
   - We care about the number of enhancements per release along with their timely graduation
   -

3. Does your [CONTRIBUTING.md] help **new** contributors engage with your group specifically by pointing
   to activities or programs that provide useful context or allow easy participation?

   - TODO: Riaan

4. If your group has special training, requirements for reviewers/approvers, or processes beyond the general [contributor guide],
   does your [CONTRIBUTING.md] document those to help **existing** contributors grow throughout the [contributor ladder]?

   - PRR and API review has dedicated docs that can help folks get start. What we need is for SIG leads to send people our way
       - API review: each SIG can identify 2-3 people to be involved in API reviews - https://github.com/kubernetes/kubernetes/blob/v1.23.0/OWNERS_ALIASES#L451-L452

5. Does the group have contributors from multiple companies/affiliations?

   - yes

6. Are there ways end users/companies can contribute that they currently are not?
   If one of those ways is more full time support, what would they work on and why?

   - We need help across all the subprojects that have well defined, ongoing work with an onboarding guide/doc
   -

## Membership

- Primary slack channel member count: #sig-architecture Slack - 2 519
- Primary mailing list member count: kubernetes-sig-architecture mailing list - 552 members
- Primary meeting attendee count (estimated, if needed): ~ 10 - 12 
- Primary meeting participant count (estimated, if needed): ~6 - 8
- Unique reviewers for SIG-owned packages: 20 <!-- in future, this will be generated from OWNERS files referenced from subprojects, expanded with OWNERS_ALIASES files -->
- Unique approvers for SIG-owned packages: 13
<!--
Used this from the owners [alias file](https://github.com/kubernetes/kubernetes/blob/master/OWNERS_ALIASES)
Counted Unique approvers and reviewers.
Did not count approvers again as reviewers. - Riaan
Enhancements [owners file](https://github.com/kubernetes/enhancements/blob/master/OWNERS_ALIASES)
- sig-architecture-approvers:
    - dims
    - derekwaynecarr
    - johnbelamaric
- api-approvers:
    - lavalamp
    - smarterclayton
    - thockin
    - liggitt
- api-reviewers:
    - andrewsykim
    - lavalamp
    - smarterclayton
    - thockin
    - liggitt
    - wojtek-t
    - deads2k
    - yujuhong
    - derekwaynecarr
    - caesarxuchao
    - mikedanese
    - sttts
    - dchen1107
    - saad-ali
    - luxas
    - janetkuo
    - justinsb
    - pwittrock
    - ncdc
    - tallclair
    - mwielgus
    - soltysh
    - jsafrane
    - dims
- conformance-behavior-approvers:
    - smarterclayton
    - johnbelamaric
    - spiffxp
- conformance/OWNERS 
	- approvers:
	  - cheftako
	  - spiffxp
	  - johnbelamaric
	- reviewers:
	  - cheftako
	  - oomichi
	  - johnbelamaric
- enhancements/OWNERS
    - approvers:
      - jeremyrickard
      - johnbelamaric
      - kikisdeliveryservice
      - mrbobbytables
    - reviewers:
      - annajung
      - jeremyrickard
      - johnbelamaric
      - kikisdeliveryservice
      - mrbobbytables
      - palnabarun

-->
<!-- in future, this will be generated from OWNERS files referenced from subprojects, expanded with OWNERS_ALIASES files -->

Include any other ways you measure group membership

## Subprojects

<!--
In future, this will be generated from delta of sigs.yaml from $YYYY-01-01 to $YYYY-12-31
Manually visible via `git diff HEAD@{$YYYY-01-01} HEAD@{$YYYY-12-31} -- $sig-id/README.md`
-->

Continuing:

- [Architecture and API Governance](https://github.com/kubernetes/community/tree/master/sig-architecture#architecture-and-api-governance-1)
    - API guidance updates
        - [Object references](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#object-references), including cross-namespace references from namespaced objects
        - [Spec and Status](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#spec-and-status) ([#5842](https://github.com/kubernetes/community/pull/5842))
    - 114 [API reviews completed](https://github.com/orgs/kubernetes/projects/13) in 2021 (30 for v1.21, 45 for v1.22, 39 for v1.23)
    - Each SIG can identify 2-3 people to be involved in API reviews - https://github.com/kubernetes/kubernetes/blob/v1.23.0/OWNERS_ALIASES#L451-L452
- [Conformance Definition](https://github.com/kubernetes/community/tree/master/sig-architecture#conformance-definition-1)
    - We started 2021 with 128 endpoints remaining without conformance test.(69.13% conformance tested)
    - The current count is [51 endpoints](https://apisnoop.cncf.io/conformance-progress/endpoints/1.24.0/?filter=untested) 
      remaining without conformance tests putting us at [87.25%](https://apisnoop.cncf.io/?conformance-only=true) conformance tested.
    - In the last 12 months
	    - Tests for 75 endpoints was promoted to conformance 
	    - 34 Endpoints was promoted to GA with conformance tests.
	    - No new technical debt was incurred
    - We moved the Conformance Office Hours Meeting from Bi-weekly to monthly in Aug 2021 to reduce the strain on attendee's calendars and increase engagement.
    - At the start of 2021 the tracking of ineligible endpoints for conformance was moved to
      [`ineligible_endpoints.yaml`](https://github.com/kubernetes/kubernetes/blob/master/test/conformance/testdata/ineligible_endpoints.yaml) file.
      This allows for more community engagement on the eligibility of endpoints for conformance before any endpoint is added to the endpoint to the 
      [Ineligible Endpoints list](https://apisnoop.cncf.io/conformance-progress/ineligible-endpoints) through pubic pull request.
    - Our target for 2022 is to clean up the last 51 endpoint and then move over to Conformance profiles 
      to further increase the value added by conformance testing.
- [Code Organization](https://github.com/kubernetes/community/tree/master/sig-architecture#code-organization-1)
     
- [Enhancements](https://github.com/kubernetes/community/tree/master/sig-architecture#enhancement-proposals)
    - In 2021, the subproject was mainly focused on improving the contributor experience via automation and thoughtful ui/ux changes making visible changes to the KEP process.
    - Major efforts were put into improving tooling on the enhancements repo throughout the year allowing us to automate more validation lessening the burden on individual approvers and resulting in more accurate information reflected in the document. All KEPs were migrated to the new template which now includes a key.yaml file.
    - The team had major contributions authoring and negotiating [KEP 2572 - Defining the Kubernetes release cadence](https://github.com/kubernetes/enhancements/blob/master/keps/sig-release/2572-release-cadence/).
    - We made progress on the receipts process, but ultimately decided that it was too disruptive for the community as a new UI/UX.
    - The team updated the repository documentation and work was started on a KEP website for easier navigation.
    - New members joined the team: 1 new approver and 2 reviewers were added to the Enhancements Subproject. A new kep tools team with initially 4 reviewers and approvers was created and more recently had an addition of 1 new approver/reviewer. We also had 2 owners transition to become our first Emeritus approvers.
    - We hope to keep improving the KEP process in 2022 and will be conducting a community survey to collect feedback, working on a Process KEP template and other improvements.
- [Production Readiness](https://github.com/kubernetes/community/tree/master/sig-architecture#production-readiness-1)
    - Production readiness review [became mandatory] for enhancements in 2021, starting with the 1.21 release cycle.
    - We added a fourth approver, Elana Hashman, who joined the team after shadowing in 1.21.
    - We created documentation on the [lifecycle of feature gates].
    - We reviewed data from the 2020 Production Readiness Survey, and prepared a 2021 survey, whose [responses we analyzed]. # **TODO:** @dims to update link to point to YouTube recording
    - The team implemented a [soft deadline] to set expectations around what can be reviewed in time for enhancements freeze.
    - We reviewed a total of 171 KEPs in 2021, averaging 16 per approver per release:
        - 51 KEPs in 1.21 (3 approvers)
        - 52 KEPs in 1.22 (4 approvers)
        - 68 KEPs in 1.23 (4 approvers)

[became mandatory]: https://groups.google.com/g/kubernetes-dev/c/GvT-qOaVsWE/m/zeazt2kODQAJ
[soft deadline]: https://groups.google.com/g/kubernetes-dev/c/EvpXsyLf5E0/m/D4Jhc7ItAgAJ
[responses we analyzed]: https://docs.google.com/document/d/1BlmHq5uPyBUDlppYqAAzslVbAO8hilgjqZUTaNXUhKM/edit#bookmark=id.g9stoq3xr2dx
[lifecycle of feature gates]: https://git.k8s.io/community/contributors/devel/sig-architecture/feature-gates.md
     

## Working groups

- [API Expression](https://git.k8s.io/community/wg-api-expression/) ([2021 report](https://git.k8s.io/community/wg-api-expression/annual-report-2021.md))

TODO: dims to send email to WG asking them to talk to us sync of async

## Operational

Operational tasks in [sig-governance.md]:

- [x] [README.md](https://github.com/kubernetes/community/blob/master/sig-architecture/README.md) reviewed for accuracy and updated if needed
- [ ] [CONTRIBUTING.md](https://hackmd.io/GY081dMNThS16WWWB46unw) reviewed for accuracy and updated if needed
      (or created if missing and your contributor steps and experience are different or more
      in-depth than the documentation listed in the general [contributor guide] and [devel] folder.)
	  **We need to create this:**
	- [ ] [HackMD](https://hackmd.io/GY081dMNThS16WWWB46unw) for template started - Riaan
- [ ] [Subprojects list](https://github.com/kubernetes/community/blob/master/sig-architecture/README.md#subprojects) and linked OWNERS files in [sigs.yaml] reviewed for accuracy and updated if needed
- [ ] SIG leaders (chairs, tech leads, and subproject owners) in [sigs.yaml] are accurate and active, and updated if needed
	- [ ] **MAYBE SOME REVIEW IS NEEDED, SOME OWNERS FILES LAST UPDATED 2019**
- [ ] [Meeting notes and recordings](https://github.com/kubernetes/community/blob/master/sig-architecture/README.md#meetings) for 2021 are linked from [README.md] and updated/uploaded if needed
	- [ ] **NEED UPDATES**
		- [ ] Dead link to Enhancements Subproject Meeting Youtube
		- [ ] Last SIG Recording uploaded Jul 2021 
- [ ] Did you have community-wide updates in 2021 (e.g. community meetings, kubecon, or kubernetes-dev@ emails)? Links to email, slides, or recordings:
      - 2021 Kubecon EU Virtual Kubernetes Conformance
[A talk on what Kubernetes Conformance is and how to contribute to it by Zach Mandeville and Caleb Woodbine.](https://www.youtube.com/watch?v=05NMwOhD6Ks)
      - 2021 Kubecon NA Virtual - Kubernetes Conformance
[A talk on the tooling that has been used the community to bring Kubernetes Conformance coverage up to 77% by Stephen Heywood and Caleb Woodbine](https://www.youtube.com/watch?v=IQsBahak7PQ)

[CONTRIBUTING.md]: https://git.k8s.io/community/sig-architecture/CONTRIBUTING.md
[contributor ladder]: https://git.k8s.io/community/community-membership.md
[sig-governance.md]: https://git.k8s.io/community/committee-steering/governance/sig-governance.md
[README.md]: https://git.k8s.io/community/sig-architecture/README.md
[sigs.yaml]: https://git.k8s.io/community/sigs.yaml
[contributor guide]: https://git.k8s.io/community/contributors/guide/README.md
[devel]: https://git.k8s.io/community/contributors/devel/README.md

