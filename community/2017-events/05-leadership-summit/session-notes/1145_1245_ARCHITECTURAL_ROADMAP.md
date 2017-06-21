Architectural Roadmap
=====================

Identify the following before beginning your session. Do not move
forward until these are decided / assigned:

-   **Session Topic**: Architectural Roadmap
-   **Topic Facilitator(s)**: Brian Grant
-   **Note-taker(s) (Collaborating on this doc)**: Jason Singer DuMars,
    Nilay Yener
-   **Person responsible for converting to Markdown & uploading to
    Github:** nyener@

### Session Notes

[Link](https://docs.google.com/presentation/d/1oPZ4rznkBe86O4rPwD2CWgqgMuaSXguIBHIE7Y0TKVc/edit)

Not using “core” because it is overloaded - substituting “nucleus”
instead \~ absolute necessities for Kubernetes to function

Q: Aaron C: What is an add-on? A: Concept is any Kubernetes resource
managed by the cluster for its own internal consistency/function - e.g.
extensible cluster functionality / pulling context

Q: Are CNI also add-ons? A: Yes.

Q: How do we differentiate between addons and required addons? A: Part
of conformance, if they are managed as part of the cluster, by the
cluster manager, they are an addon \~ like a driver in the kernel

Q: Helm Charts tied to lifecycle of the cluster? A: There’s an issue for
addons v2 via Justin Santa Barbara, but the lifecycle is different than
for Helm - it’s possible to use charts, but would prefer sources
accessed directly for bootstrapping

Q: How do addons get scheduled? Daemon sets only? A: bash script,
one-shot scheduler, in Google it’s called “babysitter” and predates
Borg, inspired kubelet “run once” mode, addon manager manages the
replicaset object

Q: Are these strict layers? The concept of layers introduces hierarchy
and precedence (governance does not have an implicit reliance on
application layer) A: Trying to avoid 9 layers, so we need flexibility
in the release process, more modular the better

Q: Eric Tune: Nucleus : abstract my cloud, Application: Run my workload,
Governance: Enterprise I need to do stuff A: Yes

Comment: Let’s make sure we keep this as simple as possible. BG: We want
to ensure that the user community is not significantly impacted, nor
cluster operators, one change needed will be breaking up the monolithic
V1 API group, pods in V1 group, pods in other API group, want to evolve
them more easily.

Q: Addons are at the lowest level, but we might want higher level addons
A: Bootstrapping needs attention, similar to other self-hosting issues,
some people may not run the aggregated API server

Q: Three implications: code organization / nucleus and others would be
in one repo versus another. Where is the line? 2. SIG alignment 3.
roadmap / where do we invest more? A: 1. Code org is another discussion,
but there’s a desire for more modularity. SIGs with their own codebases.
2. 3. e.g. Priority for admission controller was increased for 1.7, want
to make API aggregation important

Q: What are the lifecycles around each layer? Would they release
separately? A: There’s the release cadence, introduction of new
features, and the evolution of the APIs -- may adopt git flow processes

Q: Adding extension points is the only thing that should permit more
technical debt A: Agreed, but may also apply to cloud providers

C: Need to be able to say no to things that are obsoleted by the list

C: Cloud provider extension work is being done by one person - this is
not a good thing

### Conclusions

#### Key Takeaways / Analysis of Situation

Are there questions/concerns?

We need to get feedback on the docs and diagrams.

When these are happening, we expect support from all the SIGs, and how
to work this into backlogs.

#### Recommendations & Decisions Moving Forward (High-Level)

Specific Action Items (& Owners)

 Architectural Roadmap
=====================

Identify the following before beginning your session. Do not move
forward until these are decided / assigned:

-   **Session Topic**: Architectural Roadmap
-   **Topic Facilitator(s)**: Brian Grant
-   **Note-taker(s) (Collaborating on this doc)**: Jason Singer DuMars,
    Nilay Yener
-   **Person responsible for converting to Markdown & uploading to
    Github:** nyener@

### Session Notes

[Link](https://docs.google.com/presentation/d/1oPZ4rznkBe86O4rPwD2CWgqgMuaSXguIBHIE7Y0TKVc/edit)

Not using “core” because it is overloaded - substituting “nucleus”
instead \~ absolute necessities for Kubernetes to function

Q: Aaron C: What is an add-on? A: Concept is any Kubernetes resource
managed by the cluster for its own internal consistency/function - e.g.
extensible cluster functionality / pulling context

Q: Are CNI also add-ons? A: Yes.

Q: How do we differentiate between addons and required addons? A: Part
of conformance, if they are managed as part of the cluster, by the
cluster manager, they are an addon \~ like a driver in the kernel

Q: Helm Charts tied to lifecycle of the cluster? A: There’s an issue for
addons v2 via Justin Santa Barbara, but the lifecycle is different than
for Helm - it’s possible to use charts, but would prefer sources
accessed directly for bootstrapping

Q: How do addons get scheduled? Daemon sets only? A: bash script,
one-shot scheduler, in Google it’s called “babysitter” and predates
Borg, inspired kubelet “run once” mode, addon manager manages the
replicaset object

Q: Are these strict layers? The concept of layers introduces hierarchy
and precedence (governance does not have an implicit reliance on
application layer) A: Trying to avoid 9 layers, so we need flexibility
in the release process, more modular the better

Q: Eric Tune: Nucleus : abstract my cloud, Application: Run my workload,
Governance: Enterprise I need to do stuff A: Yes

Comment: Let’s make sure we keep this as simple as possible. BG: We want
to ensure that the user community is not significantly impacted, nor
cluster operators, one change needed will be breaking up the monolithic
V1 API group, pods in V1 group, pods in other API group, want to evolve
them more easily.

Q: Addons are at the lowest level, but we might want higher level addons
A: Bootstrapping needs attention, similar to other self-hosting issues,
some people may not run the aggregated API server

Q: Three implications: code organization / nucleus and others would be
in one repo versus another. Where is the line? 2. SIG alignment 3.
roadmap / where do we invest more? A: 1. Code org is another discussion,
but there’s a desire for more modularity. SIGs with their own codebases.
2. 3. e.g. Priority for admission controller was increased for 1.7, want
to make API aggregation important

Q: What are the lifecycles around each layer? Would they release
separately? A: There’s the release cadence, introduction of new
features, and the evolution of the APIs -- may adopt git flow processes

Q: Adding extension points is the only thing that should permit more
technical debt A: Agreed, but may also apply to cloud providers

C: Need to be able to say no to things that are obsoleted by the list

C: Cloud provider extension work is being done by one person - this is
not a good thing

### Conclusions

#### Key Takeaways / Analysis of Situation

Are there questions/concerns?

We need to get feedback on the docs and diagrams.

When these are happening, we expect support from all the SIGs, and how
to work this into backlogs.

#### Recommendations & Decisions Moving Forward (High-Level)

Specific Action Items (& Owners)

Action Item |Owner(s)
------------|------------
SIG Architecture Creation (ratified per unanimous vote, 6/3/2017) | Jaice Singer DuMars, Brian Grant Lead
Refine and approve Brian’s architecture | SIG Architecture

[Link to original
doc](https://docs.google.com/document/d/1nD3Y1-Tbb-hhSNg6TGPRLxKSzIuRSnVfdHginl0brrc/edit#)
[Link to original
doc](https://docs.google.com/document/d/1nD3Y1-Tbb-hhSNg6TGPRLxKSzIuRSnVfdHginl0brrc/edit#)
