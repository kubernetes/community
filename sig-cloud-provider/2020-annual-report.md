# 2020 - SIG Cloud Provider - Community Meeting Annual Reports
---
## Operational
**How are you doing with operational tasks in sig-governance.md?**

**Is your README accurate? have a CONTRIBUTING.md file?**

Our README is accurate but we do not currently have a contributing.md file.

**All subprojects correctly mapped and listed in [sigs.yaml](https://github.com/kubernetes/community/blob/master/sigs.yaml)?**

I believe so.

**What’s your meeting culture? Large/small, active/quiet, learnings? Meeting notes up to date? Are you keeping recordings up to date/trends in community members watching recordings?**

Our meetings tends to be small and quiet. We try to keep the meeting notes upto date during the meeting. 

**How does the group get updates, reports, or feedback from subprojects? Are there any springing up or being retired? Are OWNERS files up to date in these areas?**
_Same question as above but for working groups._

We go over the sub project updates during our primary biweekly meeting. 
I believe the OWNERS files are all upto date.

**When was your last monthly community-wide update? (provide link to deck and/or recording)**

---
## Membership
**Are all listed SIG leaders (chairs, tech leads, and subproject owners) active?**

I believe so.

**How do you measure membership? By mailing list members, OWNERs, or something else?**

1. We look at the number of people attending our meetings, messaing on our slack channels and contributing to our repos.
1. We look at the number of cloud providers who have people doing the above.

**How does the group measure reviewer and approver bandwidth? Do you need help in any area now? What are you doing about it?**

1. For changes to cloud provider infrastructure where all approvers can approve, I think things are going reasonably well. But I'd like to hear what others think.
1. For changes to a specific cloud provider implementation, where an approval is required from someone outside the company, I think there is a bigger problem, particularly for VMWare and Google since the SIG Chairs (the most active approvers) are from those respective companies. I feel like we need to work toward having more approvers (e.g. get to the point where the defacto bus factor is at least 2 for any change). I don't think this necessarily requires more top level approvers in the short term-- getting more sub-tree approvers (with good diversity across cloud providers) and expanding their ownership seem like a good steps.

**Is there a healthy onboarding and growth path for contributors in your SIG? What are some activities that the group does to encourage this? What programs are you participating in to grow contributors throughout the contributor ladder?**

Most of our SIG members are contributors from at least one cloud provider. We try to reach our to the various cloud providers during KubeCon to make sure they are aware of us and that their voices are being heard. New members come in slowly enough that we can generally provide direct mentorship and arrange this through our slack channel or meetings.

**What programs do you participate in for new contributors?**

None, our contributors tend to be representatives from the various companies who provider Kubernetes on the Cloud. It would be nice to expand this however.

**Does the group have contributors from multiple companies/affiliations? Can end users/companies contribute in some way that they currently are not?**

We try hard to make sure we have presentation from each of the interested cloud providers.

---
## Current initiatives and project health
_Please include links to KEPs and other supporting information that will be beneficial to multiple types of community members._

**What are initiatives that should be highlighted, lauded, shout out, that your group is proud of? Currently underway? What are some of the longer tail projects that your group is working on?**

[Feature: implement the BackendManager list](https://github.com/kubernetes-sigs/apiserver-network-proxy/pull/144) fixes a long standing issue for people running clusters with nodes which are either not routable or at least not close to one another.
[Fix flag passing in CCM](https://github.com/kubernetes/kubernetes/pull/98210) fixes a significant problem/cost most cloud providers experienced trying to remain current on their CCM dependency.
[Extending Apiserver Network Proxy to handle traffic originated from Node network](https://github.com/kubernetes/enhancements/tree/master/keps/sig-cloud-provider/2025-extend-konnectivity-for-both-directions) is an interesting and difficult initiative around build out K8s where the control plane and data plane are remote.


**Year to date KEP work review: What’s now stable? Beta? Alpha? Road to alpha?**

[Promoting Cloud Provider Labels to GA](https://github.com/kubernetes/enhancements/tree/master/keps/sig-cloud-provider/837-cloud-provider-labels) to GA

[API Server Network Proxy](https://github.com/kubernetes/enhancements/tree/master/keps/sig-api-machinery/1281-network-proxy) to Beta

[Reporting Conformance Test Results to Testgrid](https://github.com/kubernetes/enhancements/tree/master/keps/sig-cloud-provider/2390-reporting-conformance-test-results-to-testgrid) ongoing.
[Building Kubernetes Without In-Tree Cloud Providers](https://github.com/kubernetes/enhancements/tree/master/keps/sig-cloud-provider/1179-building-without-in-tree-providers) to Alpha
[Extending Apiserver Network Proxy to handle traffic originated from Node network](https://github.com/kubernetes/enhancements/tree/master/keps/sig-cloud-provider/2025-extend-konnectivity-for-both-directions)
[Kubelet Credential Providers](https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/2133-kubelet-credential-providers) to Alpha
[Leader Migration for Controller Managers](https://github.com/kubernetes/enhancements/issues/2436) to Alpha

**What areas and/or subprojects does the group need the most help with?**

We run the cloud provider extraction working group initiative. We need help engaging more with cloud providers. We also really need help dealing with the both the K/K "cluster" directory and resolving how we properly test K/K in the absence of a cloud provider.

**What's the average open days of a PR and Issue in your group? / what metrics does your group care about and/or measure?**

Pretty long :(
We are trying to build up a team of multiple contributors for every cloud provider to help triage and handle these. We are slowly growing this and its getting better but more so with some cloud providers than others.

---
