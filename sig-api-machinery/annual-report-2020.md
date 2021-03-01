# Kubernetes SIG API Machinery - 2020 Annual report
David Eads, Daniel Smith, Federico Bongiovanni


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

      Yes, the [README](https://github.com/kubernetes/community/blob/master/sig-api-machinery/README.md) is accurate.
   2. All subprojects correctly mapped and listed in sigs.yaml?

      Yes, our [subprojects](https://github.com/kubernetes/community/blob/master/sig-api-machinery/README.md#subprojects) are current.
   3. What’s your meeting culture? Large/small, active/quiet, learnings? Meeting notes up to date?
 
      Are you keeping recordings up to date/trends in community members watching recordings?
   
      We have two main meetings, both fairly small, with [notes and agenda up to date](https://docs.google.com/document/d/1x9RNaaysyO0gXHIr1y50QFbiL1x8OWnk2v3XnrdkT5Y/edit).
      [Our recordings](https://www.youtube.com/playlist?list=PL69nYSiGNLP21oW3hbLyjjj4XhrwKxH2R) are usually uploaded within two weeks.
      
      There are bug scrub meetings every Tuesday and Thursday.
      
2. How does the group get updates, reports, or feedback from subprojects? 
   Are there any springing up or being retired? Are OWNERS.md files up to date in these areas?
   
   We get updates on an ad-hoc basis.
   We have approved a prototyping project ([apiserver-runtime](https://github.com/kubernetes-sigs/apiserver-runtime)) and have no plans to retire any at this time.
   We have not actively pruned OWNERS, some people have been added to various subprojects.

3. Same question as above but for working groups.
   [wg-api-expression](https://github.com/kubernetes/community/blob/master/wg-api-expression/README.md) has its own 
   regular meeting cadence and did its own [annual report](https://github.com/kubernetes/community/blob/master/wg-api-expression/2020-annual-report.md).
   
   [wg-component-standard](https://github.com/kubernetes/community/blob/master/wg-component-standard/README.md) has its own
   regular meeting cadence.
   The working group is not as active as it once was, see the [mailing list thread](https://groups.google.com/g/kubernetes-dev/c/sQGrk6HWyj0).
   
   [wg-multitenancy](https://github.com/kubernetes/community/tree/master/wg-multitenancy) has its own regular meeting cadence
   and did its own [annual report](https://github.com/kubernetes/community/blob/master/wg-multitenancy/2021-annual-report.md).

4. When was your last public community-wide update? (provide link to deck and/or recording)
   [May 2020](https://docs.google.com/presentation/d/1UWRaMVtTD3yVhJ3MGBpt7LRIaRHTaQZoGlDT7Bl7jLE/edit#slide=id.g401c104a3c_0_0)

## Membership
1. Are all listed SIG leaders (chairs, tech leads, and subproject owners) active?

   Yes.

2. How do you measure membership? By mailing list members, OWNERs, or something else?

   We don’t measure membership.

3. How does the group measure reviewer and approver bandwidth?
   Do you need help in any area now? What are you doing about it?
   
   Our predicted rate of feature delivery and stability roughly matches the achieved rate.
   While we would be happy to see developers move up the ladder, we don’t see a pressing need to adjust the current rate.
   
   We perform twice a week triage and our [issue open/close rates are holding steady](https://k8s.devstats.cncf.io/d/39/issues-opened-closed-by-sig?orgId=1&var-period=d7&var-sig_name=api-machinery&var-kind_name=All).

4. Is there a healthy onboarding and growth path for contributors in your SIG? What are some activities that the group 
   does to encourage this? What programs are you participating in to grow contributors throughout the contributor ladder?
   
   We see patches from first time contributors, we regularly accept agenda items from contributors from other sigs and
   first time contributors.

5. What programs do you participate in for new contributors?

   We don’t participate in any particular programs.
   We find many contributors via slack, PRs, and issues.

6. Does the group have contributors from multiple companies/affiliations?
   Can end users/companies contribute in some way that they currently are not?
   
   Yes, there are contributors from [multiple companies](https://k8s.devstats.cncf.io/d/74/contributions-chart?orgId=1&var-period=d7&var-metric=contributions&var-repogroup_name=SIG%20API%20Machinery&var-country_name=All&var-company_name=All&var-company=all).
   We see all sorts of contributions, varying from issues, to comments, to PRs, to designs, to sig meeting participation,
   and user-survey data.


## Current initiatives and project health
1. What are initiatives that should be highlighted, lauded, shout outs, that your group is proud of? Currently underway?
   What are some of the longer tail projects that your group is working on?
   
   Currently underway:
   1. [server-side-apply](https://github.com/kubernetes/enhancements/issues/555) to GA
   2. [server-side-apply client](https://github.com/kubernetes/enhancements/tree/master/keps/sig-api-machinery/2144-clientgo-apply#alternative-1-generated-structs-where-all-fields-are-pointers)
   3. [optionally skip backend TLS verifiction](https://github.com/kubernetes/enhancements/issues/1295)
   4. [namespace labels](https://github.com/kubernetes/enhancements/pull/2162)
   5. Getting ready for CRD and admission webhook v1beta1 API removal: [reminder on kubernetes-dev](https://groups.google.com/g/kubernetes-dev/c/z_AE1EHhZF4/m/kBd3HkWxAwAJ).
   6. [Immutable fields API](https://github.com/kubernetes/enhancements/issues/1101)
   7. [API unions](https://github.com/kubernetes/enhancements/issues/1027)
   8. [warnings to GA](https://github.com/kubernetes/enhancements/issues/1693)
   9. [apiserver network proxy to beta](https://github.com/kubernetes/enhancements/issues/1281)
   10. [priority and fairness to GA](https://github.com/kubernetes/enhancements/issues/1040)

2. Year to date KEP work: What's now stable? Beta? Alpha? Road to alpha?
   1. Stable
      1. [Selector index](https://github.com/kubernetes/kubernetes/commit/fea3042f1f84129ab1cb6e481bd51343061673b7) - 1.20
      2. [Permabeta machinery (sig-arch policy)](https://github.com/kubernetes/enhancements/blob/master/keps/sig-architecture/1635-prevent-permabeta/README.md) - 1.19
      3. [Client-go context](https://github.com/kubernetes/enhancements/blob/master/keps/sig-api-machinery/1601-client-go-context/README.md) - 1.18
      4. [Client-go options](https://github.com/kubernetes/enhancements/blob/master/keps/sig-api-machinery/1601-client-go-context/README.md) - 1.18
      5. [Dry run](https://github.com/kubernetes/enhancements/blob/master/keps/sig-api-machinery/576-dry-run/README.md) - 1.18
      6. [Standardize conditions](https://github.com/kubernetes/enhancements/blob/master/keps/sig-api-machinery/1623-standardize-conditions/README.md) - 1.19
   2. Beta
      1. [Priority and fairness](https://github.com/kubernetes/enhancements/blob/master/keps/sig-api-machinery/1040-priority-and-fairness/README.md) - 1.20
      2. [Selector index](https://github.com/kubernetes/kubernetes/pull/92503) - 1.19
      3. [Self-link removal](https://github.com/kubernetes/enhancements/blob/master/keps/sig-api-machinery/1164-remove-selflink/README.md) - 1.20
      4. [Warning headers](https://github.com/kubernetes/enhancements/blob/master/keps/sig-api-machinery/1693-warnings/README.md) - 1.19
      5. [Server-side apply evolution while in beta](https://github.com/kubernetes/enhancements/blob/master/keps/sig-api-machinery/555-server-side-apply/README.md) - 1.18, 1.19, 1.20
   3. Alpha
      1. [Selector index](https://github.com/kubernetes/kubernetes/pull/87939) - 1.18
      2. [API server identity](https://github.com/kubernetes/enhancements/blob/master/keps/sig-api-machinery/1965-kube-apiserver-identity/README.md) - 1.20
      3. [Efficient watch resumption](https://github.com/kubernetes/enhancements/issues/1904) - 1.20
   4. Pre-alpha
      1. [Manifest-based admission webhook](https://github.com/kubernetes/enhancements/blob/master/keps/sig-api-machinery/1872-manifest-based-admission-webhooks/README.md) 


3. What initiatives are you working on that aren't being tracked in KEPs?

   We are working on mitigating the impact of removing beta APIs in 1.22.

4. What areas and/or subprojects does the group need the most help with?

   The SIG sponsors some working groups that are largely independent. 
   
   There are several areas where regularly the SIG becomes under pressure, especially closer to code freezes and the
   vast amount of code owned by API Machinery.
   
   The ecosystem of the different Kubernetes Clients that we own grows more or less organically. Client-go and
   Python-client are probably the bigger ones.
   
   There are some packages that API Machinery owns and come out usually in our triage meetings, and that we most likely
   don't know much about: this happens often when Kubernetes is upgrading libraries for example. 


5. What metrics/community health stats does your group care about and/or measure? Examples?

   On the technical health of the SIG, we look at
   - the ratio of open/close PRs
   - the ratio of open/close Issues
   - overall age of open Issues
   - Number of active contributors to the sig
   - diverse representation of companies in the sig participants
