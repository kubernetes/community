# Kubernetes SIG Auth 2020 Annual Report

## Operational

**How are you doing with operational tasks in sig-governance.md?**

**Is your README accurate? have a CONTRIBUTING.md file?**

 - Yes, our README is accurate. As of this report, we do not have a CONTRIBUTING.md file. We will be [creating one](https://github.com/kubernetes/community/issues/5760) shortly.

**All subprojects correctly mapped and listed in sigs.yaml?**

 - Yes.

**What’s your meeting culture? Large/small, active/quiet, learnings? Meeting notes up to date? Are you keeping recordings up to date/trends in community members watching recordings?**

 - Our bi-weekly meetings are usually pretty active with medium size attendance consists of mostly the same people. Some meetings may get more people depending on the topic.

 - Subproject meetings are smaller.

 - We take [meeting notes](https://docs.google.com/document/d/1woLGRoONE3EBVx-wTb4pvp4CI7tmLZ6lS26VTbosLKM/edit#) during the meeting. Unfortunately we do not do a great job of keeping good meeting notes. We are discussing various ways to improve. 

 - The recordings serve as a historical record for bi-weekly SIG meeting and special topics meetings. They are uploaded to YouTube automatically. Then SIG chairs add the video to the [SIG Auth playlist](https://www.youtube.com/playlist?list=PL69nYSiGNLP0VMOZ-V7-5AchXTHAQFzJw).

**How does the group get updates, reports, or feedback from subprojects? Are there any springing up or being retired? Are OWNERS files up to date in these areas?**

 - We do not do much here today. We will ask subproject owners to give updates in the main SIG meeting on some cadence. 
 
 - Our OWNERS files are up-to-date.

**How does the group get updates, reports, or feedback from working groups? Are there any springing up or being retired? Are OWNERS files up to date in these areas?**

 - We do not do much here today. We will ask working group owners to give updates in the main SIG meeting on some cadence. 
 
 - Our OWNERS files are up-to-date.

**When was your last monthly community-wide update? (provide link to deck and/or recording)**

 - [February, 2020 Slides](https://docs.google.com/presentation/d/1HBMqr5V79S8BSrSMAxPdQiyyCL9byBBWj2D4WrR3hPY).

## Membership

**Are all listed SIG leaders (chairs, tech leads, and subproject owners) active?**

- Yes, all SIG chairs and leads are active.

- Need to review subproject owners to confirm.

**How do you measure membership? By mailing list members, OWNERs, or something else?**

- We do not currently measure membership. We could start with Slack, mailing list members, and meeting attendance. As of this report, there are 2,136 members of the main SIG Auth Slack channel and 447 members of the mailing list. 

**How does the group measure reviewer and approver bandwidth? Do you need help in any area now? What are you doing about it?**

- We do not currently measure this. Need to look at how other SIGs do this today.

**Is there a healthy onboarding and growth path for contributors in your SIG? What are some activities that the group does to encourage this? What programs are you participating in to grow contributors throughout the contributor ladder?**

- Currently there is no onboarding or growth path. This is something we are working on and learning from other SIGs. We will start by creating a CONTRIBUTING.md file.

**What programs do you participate in for new contributors?**

- We have intro sessions at KubeCon to help new contributors get started.

- We now have weekly issues and PRs triage meetings. New contributors can join to help with new issues and PRs. 

**Does the group have contributors from multiple companies/affiliations? Can end users/companies contribute in some way that they currently are not?**

- Yes. Our chairs, leads, contributors, participants, and subproject owners are from various companies.

## Current initiatives and project health

- Please include links to KEPs and other supporting information that will be beneficial to multiple types of community members. 

**What are initiatives that should be highlighted, lauded, shout out, that your group is proud of? Currently underway? What are some of the longer tail projects that your group is working on?**

- [CSR v1](https://github.com/kubernetes/enhancements/tree/master/keps/sig-auth/1513-certificate-signing-request)
- [Token Request / bound SA token admission](https://github.com/kubernetes/enhancements/tree/master/keps/sig-auth/1205-bound-service-account-tokens)
- [client-go auth plugins](https://github.com/kubernetes/enhancements/blob/master/keps/sig-auth/541-external-credential-providers)
- [external kubelet credential providers](https://github.com/kubernetes/enhancements/blob/master/keps/sig-node/2133-kubelet-credential-providers)
- [New features in Secrets Store CSI driver](https://secrets-store-csi-driver.sigs.k8s.io/introduction.html#features)
- [Pod Security Policy Replacement](https://github.com/kubernetes/enhancements/issues/2579)

**Year to date KEP work review: What’s now stable? Beta? Alpha? Road to alpha?**

- Stable
    - BoundServiceAccountToken (v1.22)
    - Certificates API (v1.19)
    - TokenRequest (v1.20)
    - TokenRequestProjection (v1.20)
    - RootCAConfigMap (v1.21)
    - ServiceAccountIssuerDiscovery (v1.21)
    - client-go auth plugins (v1.22)
- Beta
    - CSIServiceAccountToken (v1.21)
- Alpha
    - Hierarchical Namespace Controller
- Road to alpha
    - Pod Security Policy Replacement

**What areas and/or subprojects does the group need the most help with?**

- Audit Logging
    - https://github.com/kubernetes/kubernetes/issues/101597
    - https://github.com/kubernetes/kubernetes/issues/84571
    - https://github.com/kubernetes/kubernetes/issues/82295
    - https://github.com/kubernetes/kubernetes/issues?q=is%3Aopen+is%3Aissue+label%3Aarea%2Faudit+
- Testing
    - https://github.com/kubernetes/enhancements/issues/541#issuecomment-799372909
- KMS
    - [KMS-Plugin: Areas for improvement](https://docs.google.com/document/d/1-WHXX_Dh_MNcJb2QJxF0gOAvLjh0fAnc3QrylWdMZJA/edit)

**What's the average open days of a PR and Issue in your group? / what metrics does your group care about and/or measure?**

- Based on devstats [Issue Velocity / Inactive Issues by SIG for 90 days or more](https://k8s.devstats.cncf.io/d/73/inactive-issues-by-sig?orgId=1&var-sigs=%22auth%22) at the time of writing this report, average is 7.  
- Based on devstats [PR Velocity / Awaiting PRs by SIG for 90 days or more](https://k8s.devstats.cncf.io/d/70/awaiting-prs-by-sig?orgId=1&var-sigs=%22auth%22) at the time of writing this report, average is 47.  
