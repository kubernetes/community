# SIG UI 2021 Annual Report

**Authors:**

Sebastian Florek ([@floreks](https://github.com/floreks))
Marcin Maciaszczyk ([@maciaszczykm](https://github.com/maciaszczykm))
Jeffrey Sica ([@jeefy](https://github.com/jeefy))

- [Operational](#operational)
- [Membership](#membership)
- [Current initiatives and project health](#current-initiatives-and-project-health)

## Operational

**How are you doing with operational tasks in [sig-governance.md](/committee-steering/governance/sig-governance.md)?**

The SIG primarily works asynchronously. In the past, our meetings would last less than five minutes and be relatively disruptive. Since the release of the 2.0 Dashboard, meeting attendance transitioned from optional to near-non-existent. Instead, we coordinate through the #sig-ui Slack channel and handle any project planning via GitHub. All chairs are regular contributors of the Dashboard subproject and help review PRs and triage issues. 

**Is your README accurate? have a CONTRIBUTING.md file?**

Yes, our README is accurate and is typically updated with every new Dashboard release. Also we do have a CONTRIBUTING.md file however it has not been updated with any more info beyond the standard template. Both of these files reside within our respective repositories.

**All subprojects correctly mapped and listed in [sigs.yaml](/sig-list.md)?**

Yes. The only subproject under SIG-UI is the Kubernetes Dashboard. There is an additional companion service that gets deployed to feed metrics into the Dashboard (Dashboard Metrics Scraper) but it isn't considered a subproject by the SIG.

**What’s your meeting culture? Large/small, active/quiet, learnings? Meeting notes up to date? Are you keeping recordings up to date/trends in community members watching recordings?**

As stated above, the SIG currently doesn't hold any meetings and instead has been asynchronous for roughly a year. All communication is done via the #sig-ui Slack channel and all planning activities are managed within the Dashboard GitHub repo.

**How does the group get updates, reports, or feedback from subprojects? Are there any springing up or being retired? Are OWNERS files up to date in these areas?**

As we have no subproject other than the Dashboard, there is no real feedback loop besides GitHub Issues and Slack channel traction.

**When was your last monthly community-wide update? (provide link to deck and/or recording)**

Our last community-wide update was in May 2020

Slide Deck: https://docs.google.com/presentation/d/1W4NioOkAF2VFiu-5t80p2vlu3_OznpugiyiViFuitaM/edit?usp=sharing

Recording: https://youtu.be/ZyUQiN3S6TE?list=PL69nYSiGNLP1pkHsbPjzAewvMgGUpkCnJ&t=837

## Membership

**Are all listed SIG leaders (chairs, tech leads, and subproject owners) active?**

Yes

**How do you measure membership? By mailing list members, OWNERs, or something else?**

We measure activity by activity within the GitHub repo. Filing/responding to issues and submitting/reviewing PRs.

**How does the group measure reviewer and approver bandwidth? Do you need help in any area now? What are you doing about it?**

We currently have no metric for reviewer/approver bandwidth. We've worked in a communal manner to ensure nothing becomes stale. This has generally worked over the last several years, and isn't a major concern for us. Our existing OWNERS file is up to date and we have active reviewers/approvers.

Over the course of 2020 we've added 9 people in OWNERS files within our i18n translations.

**Is there a healthy onboarding and growth path for contributors in your SIG? What are some activities that the group does to encourage this? What programs are you participating in to grow contributors throughout the contributor ladder?**

The SIG doesn't participate in any activities or programs explicitly meant to help grow contributors. We attempt to be responsive and welcoming within our issues and PRs to try and retain those who take the first step and engage in conversation however. With this, we've maintained a small albeit responsive pool of active contributors to the SIG.

**What programs do you participate in for new contributors?**

We attempt to maintain `Good First Issue` labels on things that we think early-lifecycle contributors could begin with. During KubeCons we also do introduction presentations and mentoring sessions to try and funnel contributors in to our SIG. 

**Does the group have contributors from multiple companies/affiliations? Can end users/companies contribute in some way that they currently are not?**

Yes, current chairs are spread between Kubermatic (formerly Loodse) and Red Hat. Additional contributors in OWNERS files are also spread across multiple companies including but not limited to NEC and Tencent. 

## Current initiatives and project health

In addition to the ongoing effort of maintenance, we are improving support visualizing [various resource objects](https://github.com/kubernetes/dashboard/issues/5232) and adding the ability for installations to have custom themes.

The often-unspoken reality is, a UI or Dashboard for Kubernetes is often a differentiator for many companies. Additionally, there are many different ways and opinions on what a front-end for Kubernetes should look like.

**What are initiatives that should be highlighted, lauded, shout out, that your group is proud of? Currently underway? What are some of the longer tail projects that your group is working on?**

When the Dashboard was originally written, it used a pull/polling design. Now, we are working towards supporting the [shared informer pattern](https://github.com/kubernetes/dashboard/issues/5320), making the Dashboard more-real-time.

Another major area to highlight is we are continually onboarding new language translations for the Dashboard.

**Year to date KEP work review: What’s now stable? Beta? Alpha? Road to alpha?**

There have been no KEPs opened as part of SIG-UI.

**What areas and/or subprojects does the group need the most help with?**

Recruiting new contributors. Front-ends are very opinionated, and those with differing opinions seem to prefer starting from scratch rather than put in the effort in learning an existing project to contribute to. This has been observed time and time again.

To be able to traverse/contribute, a new contributor needs to have an understanding of AngularJS, golang, and a reasonably in-depth knowledge of Kubernetes client-go package. 

**What's the average open days of a PR and Issue in your group? / what metrics does your group care about and/or measure?**

There are two primary sources of PRs: Automated dependency PRs (Dependabot) and contributor-created PRs.

At a glance, Dependabot PRs are merged or closed within 24h, while contributor-created PRs are typically reviewed within 3-4 days.

Currently there are no real metrics that we measure regarding Issue/PR turnaround. 

## Additional Links

[Developer Statistics](https://k8s.devstats.cncf.io/d/13/developer-activity-counts-by-repository-group?orgId=1&var-period_name=Last%20year&var-metric=contributions&var-repogroup_name=SIG%20UI&var-country_name=All)

[27 New Contributors in the last year](https://k8s.devstats.cncf.io/d/52/new-contributors?orgId=1&var-repogroup_name=SIG%20UI&from=now-1y&to=now)

[Opened PR Velocity](https://k8s.devstats.cncf.io/d/25/open-pr-age-by-repository-group?orgId=1&var-period=d7&var-repogroup_name=SIG%20UI&var-kind_name=All&from=now-1y&to=now)

