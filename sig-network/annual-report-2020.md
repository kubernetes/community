# 2020 SIG network annual report

This report reflects back on CY 2020 and was written in April 2021.

> Note: Each response should have supporting documentation, linked KEPs, graphs, data
> Note: Err on the side of being descriptive with your responses. End users and others that read this might not know our upstream acronyms, shorthand, or where to find a reference. 

# Operational

- How are you doing with operational tasks in [sig-governance.md](https://github.com/kubernetes/community/blob/master/committee-steering/governance/sig-governance.md)?
  - Is your [README](./README.md) accurate? 
    - Yes, our README is up to date.
  - Have a CONTRIBUTING.md file?
    - No - we do not have a CONTRIBUTING.md file.
  - All subprojects correctly mapped and listed in [sigs.yaml](https://github.com/kubernetes/community/blob/master/sig-list.md)?
    - Yes, all subprojects are correctly listed. 
    - We have one unofficial group that should probably be a subproject - the network plumbing group.
      - [Network plumbing group meeting minutes](https://docs.google.com/document/d/1oE93V3SgOGWJ4O1zeD1UmpeToa0ZiiO6LqRAmZBPFWM/edit)
  - What’s your meeting culture? Large/small, active/quiet, learnings? Meeting notes up to date? Are you keeping recordings up to date/trends in community members watching recordings?
    - We generally have fairly large meetings with >20 attendees, although a much smaller group of regular speakers.
    - Meetings generally start with 15 minutes of group issue or PR triage, followed by 45 minutes of topics from the agenda, which are crowd-sourced throughout the week leading up.
    - Recordings are automatically uploaded to the YouTube channel and are up to date.
    - Notes are stored in the minutes document, which is also where topics can be added to the agenda in advance.

- How does the group get updates, reports, or feedback from subprojects? Are there any springing up or being retired? Are OWNERS files up to date in these areas?
  - Updates from subprojects are generally the responsibilitiy of subproject members themselves (i.e., push not pull).
  - Updates are usually on the mailing list, occasionally subprojects will also use meeting time to discuss in real-time.
- Same question as above but for working groups.
  - Same answer as above, but for working groups.
- When was your last monthly community-wide update? (provide link to deck and/or recording)
  - Last community meeting was in July 2020. 
  - [Slides](https://docs.google.com/presentation/d/1k4uzqWCQgz8by3ZNUeXb1A5aeOpufOt4UTdGMLd6rjc/edit#slide=id.g401c104a3c_0_0)
  - [Recording](https://www.youtube.com/watch?v=J3O8fXTm3HE&list=PL69nYSiGNLP1pkHsbPjzAewvMgGUpkCnJ)
  - We're overdue.

# Membership

- Are all listed SIG leaders (chairs, tech leads, and subproject owners) active?
  - Yes
- How do you measure membership? By mailing list members, OWNERs, or something else?
  - We don't currently have an explicit way of measuring membership.
- How does the group measure reviewer and approver bandwidth? Do you need help in any area now? What are you doing about it?
  - We don't currently have an explicit way of measuring reviewer / approver bandwidth. Everyone feels overburdened, we're not doing anything about it.
- Is there a healthy onboarding and growth path for contributors in your SIG? What are some activities that the group does to encourage this? What programs are you participating in to grow contributors throughout the contributor ladder?
  - The easiest way for new contribs to get involved is to attend a meeting and volunteer to help triage an issue during our triage session.
- What programs do you participate in for new contributors?
  - No programs specifically targeting new contribs at the moment.
- Does the group have contributors from multiple companies/affiliations? Can end users/companies contribute in some way that they currently are not?
  - Yes - we have long-term as well as one-time contributors from a variety of companies. 

# Current initiatives and project health

> Please include links to KEPs and other supporting information that will be beneficial to multiple types of community members. 

- What are initiatives that should be highlighted, lauded, shout out, that your group is proud of? Currently underway? What are some of the longer tail projects that your group is working on?
- Year to date KEP work review: What’s now stable? Beta? Alpha? Road to alpha?
  - EndpointSlice: [GA](https://github.com/kubernetes/enhancements/tree/master/keps/sig-network/0752-endpointslices)
  - Dual-stack support: [beta](https://github.com/kubernetes/enhancements/tree/master/keps/sig-network/563-dual-stack)
  - ClusterNetworkPolicy: [road to alpha](https://github.com/kubernetes/enhancements/pull/2522)
  - Multiple CIDRs: [road to alpha](https://github.com/kubernetes/enhancements/pull/2594)
  - All-port services: [road to alpha](https://github.com/kubernetes/enhancements/pull/2611/)
  - Topology-aware-hints: [Alpha](https://github.com/kubernetes/enhancements/tree/master/keps/sig-network/2433-topology-aware-hints)
  - Ingress: [GA](https://github.com/kubernetes/enhancements/tree/master/keps/sig-network/1453-ingress-api)
- What areas and/or subprojects does the group need the most help with?
  - Tracking project health, membership, progress.
  - PR reviews
  - Reporting to other SIGs / community meeting.
- What's the average open days of a PR and Issue in your group? / what metrics does your group care about and/or measure?
  - I don't think we currently track this.
  - Here are some dashboard we probably should track: [1](https://k8s.devstats.cncf.io/d/34/pr-workload-per-sig-table?orgId=1&var-period_name=Last%20year), [2](https://k8s.devstats.cncf.io/d/44/pr-time-to-approve-and-merge?orgId=1&var-period=w&var-repogroup_name=SIG%20Network&var-repo_name=kubernetes%2Fkubernetes&var-apichange=All&var-size_name=All&var-kind_name=All)
