# Kubernetes Working Group Formation and Disbandment

## Process Overview and Motivations
Working Groups provide a formal avenue for disparate groups to collaborate around a common problem, craft a balanced
position, and disband. Because they represent the interests of multiple groups, they are a vehicle for consensus
building.  If code is developed as part of collaboration within the Working Group, that code will be housed in an
appropriate repository as described in the [repositories document].  The merging of this code into the repository
will be governed by the standard policies regarding submitting code to that repository (e.g. developed within one or
more Subprojects owned by SIGs).

Because a working group is an official part of the Kubernetes project it is subject to steering committee oversight
over its formation and disbanding.

## Goals of the process

- An easy-to-navigate process for those wishing to establish and eventually disband a new Working Group
- Simple guidance and differentiation on where a Working Group makes sense, and does not
- Clear understanding that no authority is vested in a Working Group
- Ensure all future Working Groups conform with this process

## Non-goals of the process

- Documenting other governance bodies such as sub-projects or SIGs
- Changing the status of existing Working Groups/SIGs/Sub-projects

## Working Group Relationship To SIGs
Assets owned by the Kubernetes project (e.g. code, docs, blogs, processes, etc) are owned and 
managed by [SIGs](sig-governance.md).  The exception to this is specific assets that may be owned
by Working Groups, as outlined below.

Working Groups provide structure for governance and communication channels, and as such may
own the following types of assets:

- Calendar Events
- Slack Channels
- Discussion Forum Groups

Working Groups are distinct from SIGs in that they are intend to:
 
- facilitate collaboration across SIGs
- facilitate an exploration of a problem / solution through a group with minimal governmental overhead

Working Groups will typically have stake holders whose participation is in the
context of one or more SIGs.  These SIGs should be documented as stake holders of the Working Group
(see Creation Process).

## Is it a Working Group? Yes, if...
- It does not own any code
- It has a clear goal measured through a specific deliverable or deliverables
- It is temporary in nature and will be disbanded after reaching its stated goal(s)

## Creation Process Description
Working Group formation is less tightly-controlled than SIG formation since they:

- Do not own code
- Have a clear entry and exit criteria
- Do not have any organizational authority, only influence

Therefore, Working Group formation begins with the organizers asking themselves some important questions that
should eventually be reflected in a pull request on sigs.yaml:

1. What is the exact problem this group is trying to solve?
1. What is the artifact that this group will deliver, and to whom?
1. How does the group know when the problem solving process is completed, and it is time for the Working Group to
   dissolve?
1. Who are all of the stakeholders involved in this problem this group is trying to solve (SIGs, steering committee,
   other Working Groups)?
1. What are the meeting mechanics (frequency, duration, roles)?
1. Does the goal of the Working Group represent the needs of the project as a whole, or is it focused on the interests
   of a narrow set of contributors or companies?
1. Who will chair the group, and ensure it continues to meet these requirements?
1. Is diversity well-represented in the Working Group?

Once the above questions have been answered, the pull request against sigs.yaml can be created. Once the generator
is run, this will in turn create the OWNERS_ALIASES file, readme files, and the main SIGs list.  The minimum
requirements for that are:

- name
- directory
- mission statement
- chair information
- meeting information
- contact methods
- any [sig](sig-governance.md) stakeholders

The pull request should be labeled with any SIG stakeholders and committee/steering. And since GitHub notifications
are not a reliable means to contact people, an email should be sent to the mailing lists for the stakeholder SIGs,
and the steering committee with a link to the PR. A member of the community admin team will place a /hold on it
until it has an LGTM from at least one chair from each of the stakeholder SIGs, and a simple majority of the steering
committee.

Once merged, the Working Group is officially chartered until it either completes its stated goal, or disbands
voluntarily (e.g. due to new facts, member attrition, change in direction, etc). Working groups should strive to
make regular reports to stakeholder SIGs in order to ensure the mission is still aligned with the current state.

Deliverables (documents, diagrams) for the group should be stored in the directory created for the Working Group.
If the deliverable is a KEP, it would be helpful to link it in the closed formation/charter PR for future reference.

## Disbandment Process Description

Working Groups will be disbanded if either of the following is true:

- There is no long a Chair
  - (with a 4 week grace period)
- None of the communication channels for the Working Group have been in use for the goals outlined at the founding of
  the Working Group
  - (with a 3 month grace period)
  - Slack
  - Email Discussion Forum
  - Zoom

The current Chair may step down at any time.  When they do so, a new Chair may be selected through lazy consensus
within the Working Group, and [sigs.yaml](/sigs.yaml) should be updated.

References

- [1] https://github.com/kubernetes/community/pull/1994
- [2] https://groups.google.com/a/kubernetes.io/d/msg/steering/zEY93Swa_Ss/C0ziwjkGCQAJ 


[repository document]: https://github.com/kubernetes/community/blob/master/github-management/kubernetes-repositories.md