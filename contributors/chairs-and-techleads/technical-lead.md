# Technical Lead

## Role Description

### About

Target of this document is to define and outline the Technical Lead role within
the Kubernetes community in correspondence with the existing [description of SIG
Governance](/committee-steering/governance/sig-governance.md#tech-lead). The
document can be used as guidance for Special Interest Groups (SIGs) to onboard
new Technical Leads as well as clarifying the expectations associated with this
role.

### Abstract

The Technical Lead role in Kubernetes is an optional role that each SIG can
choose to implement as part of its [governance
model](http://git.k8s.io/community/committee-steering/governance/sig-governance.md#roles).
This means SIGs can decide on their own if they want to add Technical Leads to
their charter or not. Depending on the overall size of the SIG, around two to
three people can be chosen by the SIG Chairs to support the technical aspects of
the group. To be able to fulfill their role, a Technical Lead should have the
same set of permissions as a Chair.

Generally speaking, Technical Leads are responsible for leading the SIG in
correspondence with its technical alignment. This alignment includes both
internal to the SIG and, more significantly, external to the entire Kubernetes
project. Before being able to align any technical direction, it is necessary to
establish a technical vision within the SIG. The technical vision should be
continuously updated and turned into reality, whereas larger features of the
vision can be outlined in dedicated roadmaps.

Expectations of Technical Leads are:

- They're involved in the source code base of the SIG to be able to take the
  right decisions.
- They actively identify risk and maintain a high level of trust with other
  members of the SIG.

Technical Leads have the responsibility to track the technical quality of the
deliverables of the team if a roadmap exists. They are volunteering to provide
senior leadership to the SIG’s short-term and long-term vision.

Examples for technical leadership within a SIG are:

- Ensure that the team utilizes appropriate engineering practices which apply to
  the whole Kubernetes organization. One example would be using Prow for
  Continuous Integration (CI) practices.
- Continuously evaluate technical challenges within the SIG and work towards
  removing them as part of the vision.
- Take changing environments into account to adapt the technical vision if
  needed. For example, if interdependent SIGs have a requirement for a technical
  change, then it is in the responsibility of the Technical Leads to drive
  towards a feasible solution for the whole community.

Technical leadership within a SIG should focus on solving the "How" questions
rather than people and team growth related ones (the latter responsibilities are
those of the SIG Chairs). Mentoring team members around solving technical tasks
also falls also into the responsibility of the Technical Leads. This includes
proper onboarding (or delegation of it) of new team members with respect to
technical workflows within the SIG. It also means that Technical Leads are
responsible for building, sharing and documenting the context to ensure that
there is a pipeline for the team members. Technical Leads also help mediating
technical debates within and in correlation with other SIGs, they unblock
outstanding work and have the ability to ask the right questions or say "no".
They actively participate in building solutions around the technical vision and
lead discussions around software architectural decisions.

The common skill set of Technical Leads divides into three areas:

- **Leadership** – Coaching team members to reach their target. Delegation of
  work to ensure personal technical growth of them.
- **Development** – Knowledge about the code and the quality standards within
  the community. Being able to help the team to solve technical obstacles
  without having a need of being the expert for everything.
- **Architecture** – Wide range understanding of the SIGs work in correlation to
  the whole community. Establishing future plans by continuously working on the
  technical vision of the SIG.

If you are interested in becoming a Technical Lead, speak with the [appropriate
SIG Chairs](https://github.com/kubernetes/community/blob/master/sig-list.md).
