# WG Reliability Charter

This charter adheres to the conventions described in the [Kubernetes Charter README]
and uses the Roles and Organization Management outlined in [sig-governance].

[sig-governance]: https://github.com/kubernetes/community/blob/master/committee-steering/governance/sig-governance.md
[Kubernetes Charter README]: https://github.com/kubernetes/community/blob/master/committee-steering/governance/README.md

## Scope

The Reliability Working Group (WG Reliability) is organized with the goal of 
allowing users to safely use Kubernetes for managing production workloads by
ensuring Kubernetes is stable and reliable.

### In Scope

- What reliability means for Kubernetes and how to measure it?
- Measuring Kubernetes reliability in tests
- Introducing criteria for blocking the release if the reliability is
  below the bar
- Building a list of end-user outages and reliability issues
  (if applicable with mitigations and/or workarounds)
- Creating and prioritizing a list of areas that require reliability
  investments
- Work with relevant SIGs on delivering necessary infrastructure
  (e.g. test frameworks) to unblock further steps
- Initiate and drive cross-SIG reliability improvements

### Out of scope

- Designing and executing improvements clearly falling into individual SIG
  responsibilities.

## Special Powers

The Reliability WG has a power to block feature-oriented contributions from
any SIG if requested reliability-related improvements are not being addressed.
Before it can be exercised, sig-arch must approve the criteria suggested by
this working group.

Given WGs are by-definition temporary, on WG Reliability retirement we will
pass this responsibility to  to SIG Architecture Production Readiness subproject
or to SIG Architecture generally for reassignment at the leads’ discretion.

## Stakeholders

Stakeholders in this working group span multiple SIGs.

In the first phase of defining reliability for Kubernetes building list of
reliability gaps and areas for investments the following SIGs will be
involved:

- SIG Architecture
  High-level input on requirements.
- SIG Scalability
  Input on scale test gaps and reliability issues at scale.
- SIG Cluster Lifecycle
  Input on cluster setup and upgrade mechanics.
- SIG Release
  Input on blocking and soak requirements.
- SIG Testing
  Input on testing mechanics, missing frameworks, etc.
- SIG *
  Input on reliability gaps in their areas.

The group will be also reaching out to users and cluster operator
(e.g. via surveys), to build the full picture.

In the later phase improving reliability, every single SIG may potentially
be involved depending on the findings from the initial phase.

## Deliverables

The artifacts the group is supposed to deliver include:
- Document defining what reliability means for Kubernetes and how to measure it.
- List of known user outages and potential failure modes
- List of specific investmenets that should happen to improve reliability
- Set of processes to introduce in Kubernetes to avoid over time degradation
  of reliability

The actual investments will be owned by corresponding SIGs.

## Roles and Organization Management

This sig follows adheres to the Roles and Organization Management outlined in
[sig-governance] and opts-in to updates and modifications to [sig-governance].

[sig-governance]: https://github.com/kubernetes/community/blob/master/committee-steering/governance/sig-governance.md

## Timelines and Disbanding

The exact timeline for existing of this working group is hard to predict at
this time.

The group will start working on the deliverables mentioned above. Once the
group we will be satisfied with the current shape of them and no additional
coordination on their execution will be needed, we will retire Working Group
and pass oversight of reliability to SIG Architecture PRR subproject.
