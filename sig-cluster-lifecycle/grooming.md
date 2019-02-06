# SIG Cluster Lifecycle Grooming

## Overview 

The purpose of this document is to outline the current backlog grooming practices, or process, that are employed across several sub-projects under SIG Cluster Lifecycle.  By providing a consistent set of best practices, it allows developers to work across sub-projects fairly easily.  More importantly, it signals direction and priority to developers who may be new to the sub-projects.  

This document is an overview of the practices that are currently followed and any feedback, or refinement, is solicited. 

## Process 

### Planning a milestone
For every sub-project under sig-cluster lifecycle, we highly recommend holding a planning session prior to beginning a new milestone.  The purpose of planning is to outline a set of high level deliverables that the community has decided that they want to work on for this upcoming milestone.  

### Creating milestones
Now that you have a plan, you're about to embark upon a new journey, let's call this journey `1.X` on a fresh repo.  The first step is to create two milestones (`1.X`, and `Next`).  The reason for the `Next` milestone is simply to create a placeholder to segregate work items out of the current milestone.  We typically try to stick with just two milestones, because we've found over time that this simplfies execution and keeps us focused on delivering against the current milestone.

### Breaking down the plan 
Now you have a plan and your milestones, everything needs to be written into github issues and assigned the appropriate labeling `feature,bug,doc` as well as a `priority` and attach it to a milestone.  If there are members who have expertise in a given area and are willing to work on the issue then assign them, otherwise add the `help wanted` label to that issue.  This does not mean that member is the only person who can work on an issue, it just means they have said they are interested.  If another member wants to work on the issue, simply coordinate with the member who has been assigned and ask if it is ok to work on that issue.  

### Execution 
When a member starts to work on an issue and they are assigned, we request that they mark the issue with the `lifecycle/active` label.  This denotes to the community that the issue is currently being worked on and it signals to the reviewers and approvers that patches are coming soon.  Every PR against an issue should cross-reference that issue, and use standard kubernetes practices for closing fixed issues.  

### Continuous triaging the backlog
Over time, you will likely see a influx of issues from the community.  We recommend to go through the backlog during your `office hours` calls to talk about the new issues and determine a relative priority and assign it a member and milestone if needed. 

### Empowering approvers 
A key to this process, is empowering and promoting approvers to help in the triage process.  The goal is to help build up the community and develop a shared sense of responsibility for the sub-project.  If there is any ambiguity on triaging an issue, we recommend working it out amongst the approvers, or if needed, escalate to the chairs of the SIG.    

### End of milestone
Upon the end of the milestone, you start the whole process over again, but you will need to shuffle through the `Next` milestone and determine what you want to address in during your next planning session. 

### Examples 

- [Cluster API provider AWS Office Hours](https://youtu.be/ERjLdfjNsG0?t=315)
- [Kubeadm Office Hours](https://www.youtube.com/watch?v=KcVurqzxNR4&t=795)

