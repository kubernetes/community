# 2020 - SIG Scalability - Community Meeting Annual Report

## Operational

**How are you doing with operational tasks in sig-governance.md?**

- **Is your README accurate? have a CONTRIBUTING.md file?**

The README is accurate an up-to-date.
We don't have CONTRIBUTING.md and we're planning to update our
[developer documentation]. The [desired documentation] is mentioned
in the corresponding bug.

[developer documentation]:https://github.com/kubernetes/community/tree/master/contributors/devel/sig-scalability
[desired documentation]:https://github.com/kubernetes/community/issues/5236#issuecomment-805607895

- **All subprojects correctly mapped and listed in sigs.yaml?**

Yes.

- **What’s your meeting culture? Large/small, active/quiet, learnings?
Meeting notes up to date? Are you keeping recordings up to date/trends
in community members watching recordings?**

Quite small (usually 5-10 people) and generally quiet. We try to keep
meeting notes up to date. We're often lagging with updating recordings :(

**How does the group get updates, reports, or feedback from subprojects?
Are there any springing up or being retired? Are OWNERS files up to date
in these areas?**

There isn't any formal process. Updates are given during regular SIG meetings.

**When was your last monthly community-wide update? (provide link to deck and/or recording)**

January'20: [slide deck]

[slide deck]:https://docs.google.com/presentation/d/1M81X3_SWwHrJaWdJsecJReOVNxjjWSBFBkhPKEV35hs/edit


## Membership

**Are all listed SIG leaders (chairs, tech leads, and subproject owners) active?**

Yes.
We don't have a good process for keeping it up-to-date, rather triggerred by special events
(like this annual report).

**How do you measure membership? By mailing list members, OWNERs, or something else?**

Active participants in SIG meetings and/or sending/reviewing PRs.

**How does the group measure reviewer and approver bandwidth?
Do you need help in any area now? What are you doing about it?**

We're not doing a very good job here. It happens that some PRs are waiting for weeks.

**Is there a healthy onboarding and growth path for contributors in your SIG?
What are some activities that the group does to encourage this?
What programs are you participating in to grow contributors throughout the contributor ladder?**

We maintain a list of `help-wanted` issues and recommend particular ones for people
interested in contributing.

**What programs do you participate in for new contributors?**

None

**Does the group have contributors from multiple companies/affiliations?
Can end users/companies contribute in some way that they currently are not?**

Yes - there are people from multiple companies.
End users mostly come with questions, they rarely want to contribute.


## Current initiatives and project health

**What are initiatives that should be highlighted, lauded, shout out, that your group is proud of?
Currently underway? What are some of the longer tail projects that your group is working on?**

The critical initiative in 2020 was speeding up scalability tests. We managed to reduce time
take by our 5k-node scalability tests from ~14h to less than 5h (~3x).

We helped validating scalability and performance impact of multiple features across the whole year
(by consulting multiple contributors from across many different SIGs).

We also visibly improved the testing framework to support other environments, tracking new
scalability/performance aspects (e.g. as OOM-killing) and extended portfolio of tested features.

Given SIG-Scalability doesn't own non-test-related code, many efforts are driven under
the auspicies of other SIGs. This includes:
- [Efficient Watch Resumption] with sig-api-machinery (Alpha 1.20)
- [Immutable Secrets and ConfigMaps] with sig-storage (Alpha 1.18, Beta 1.19)

[Efficient Watch Resumption]: https://github.com/kubernetes/enhancements/issues/1904
[Immutable Secrets and ConfigMaps]: https://github.com/kubernetes/enhancements/issues/1412


**Year to date KEP work review: What’s now stable? Beta? Alpha? Road to alpha?**

Given that SIG-Scalability doesn't really own code (other than test frameworks and tests
themselves), there aren't SIG-Scalability KEPs per-se. However, we initiate and drive many
KEPs that are officially owned by other SIGs though (and SIG-Scalability is listed as
participating SIG even though it's often doing most of the work - examples above).

**What areas and/or subprojects does the group need the most help with?**

Each subproject could benefit from additional hands.
However, the `Scalability Test Frameworks` and `Scalability and Performance tests and validation`
are the ones where we can grow contributors - the other require both very deep and wide
understanding of Kubernetes before making reasonable contributions.

**What's the average open days of a PR and Issue in your group? / what metrics does your group care about and/or measure?**

We weren't measuring it unfortunately.
What what be useful is out-of-the-box support for that across repositories
(in particular perf-tests repo is very interesting for us).

[sig-governance.md]: https://github.com/kubernetes/community/blob/master/committee-steering/governance/sig-governance.md
[sigs.yaml]: https://github.com/kubernetes/community/blob/master/sig-list.md
