# Operational
How are you doing with operational tasks in sig-governance.md?
Is your README accurate? have a CONTRIBUTING.md file?  
Yes, both the [README.md](https://github.com/kubernetes/community/blob/master/sig-apps/README.md)
and the [CONTRIBUTING.md](https://github.com/kubernetes/community/blob/master/sig-apps/CONTRIBUTING.md)
are up-to-date.

All subprojects correctly mapped and listed in sigs.yaml?  
Yes

What’s your meeting culture? Large/small, active/quiet, learnings? Meeting notes up to date? Are you keeping recordings up to date/trends in community members watching recordings?  
We hold small and active biweekly meeting. The call is divided into few sections, such as important announcements, then demos and finally discussions on the current topics. If time allows, we try to review issues and pull requests. All of the meetings are recorded and available online, all the meeting invites are up-to-date and present in community calendar.


How does the group get updates, reports, or feedback from subprojects? Are there any springing up or being retired? Are OWNERS files up to date in these areas?  
Every subproject is free to provide an update at our biweekly meeting, but we don’t enforce that specifically. OWNERS files are up-to-date.

When was your last monthly community-wide update? (provide link to deck and/or recording)  
Our last update presentation was on June 18, 2020
- [Slides](https://docs.google.com/presentation/d/18UcJQs3ThW6Vdgl_mdc1984uU16GwuIhuD0Pujl42xU/edit#slide=id.g401c104a3c_0_0)
- [Recording](https://youtu.be/ObqQxRRl9RQ?t=1111)
- [KubeCon NA 2019](https://www.youtube.com/watch?v=uqYStFxu578)

# Membership
Are all listed SIG leaders (chairs, tech leads, and subproject owners) active?  
Yes, all the lists are up-to-date.

How do you measure membership? By mailing list members, OWNERs, or something else?  
Reviewers and approvers are measured through repository OWNERS file. Membership is measured through mailing list and zoom calls.

How does the group measure reviewer and approver bandwidth? Do you need help in any area now? What are you doing about it?  
We periodically remove inactive reviewers and approvers, invite new contributors we have seen active in to join the reviewers. The bar to be accepted as an approver is significantly higher since we need to ensure the stability of the core controllers and the entire project.

Is there a healthy onboarding and growth path for contributors in your SIG? What are some activities that the group does to encourage this? What programs are you participating in to grow contributors throughout the contributor ladder?  
Recently we reviewed and unified the sig-apps-approvers alias which holds the approvers, similarly sig-apps-reviewers was reviewed for the reviewers ranks.

What programs do you participate in for new contributors?  
 - KubeCon updates.
 - Individual/1-1 mentoring.  

Does the group have contributors from multiple companies/affiliations? Can end users/companies contribute in some way that they currently are not?  
The group has contributors from multiple companies/affiliations.
[14 companies](https://k8s.devstats.cncf.io/d/8/company-statistics-by-repository-group?orgId=1&var-period=y&var-metric=contributions&var-repogroup_name=SIG%20Apps&var-repo_name=kubernetes%2Fkubernetes&var-companies=All) contributed code in the last year.



# Current initiatives and project health

What are initiatives that should be highlighted, lauded, shout out, that your group is proud of? Currently underway? What are some of the longer tail projects that your group is working on?  
- [CronJob to GA](https://github.com/kubernetes/enhancements/issues/19)
- [PDB to GA](https://github.com/kubernetes/enhancements/issues/85)

Year to date KEP work review: What’s now stable? Beta? Alpha? Road to alpha?
What areas and/or subprojects does the group need the most help with?
 - [New CronJob controller as part of work from promoting CronJobs to GA (alpha)](https://github.com/kubernetes/enhancements/issues/19)
 - [MaxSurge for DaemonSet (alpha)](https://github.com/kubernetes/enhancements/issues/1591)   

What's the average open days of a PR and Issue in your group? / what metrics does your group care about and/or measure?  
The average time to merge the PR is around [~7 days](https://k8s.devstats.cncf.io/d/44/pr-time-to-approve-and-merge?viewPanel=8&orgId=1&from=1577865600000&to=1609401600000&var-period=d7&var-repogroup_name=SIG%20Apps&var-repo_name=kubernetes%2Fkubernetes&var-apichange=All&var-size_name=All&var-kind_name=All)
