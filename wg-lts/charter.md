# WG Long Term Support

This charter adheres to the conventions described in the [Kubernetes Charter README]
and uses the Roles and Organization Management outlined in [sig-governance].

[sig-governance]: https://github.com/kubernetes/community/blob/master/committee-steering/governance/sig-governance.md
[Kubernetes Charter README]: https://github.com/kubernetes/community/blob/master/committee-steering/governance/README.md

## Scope

The Long Term Support Working Group (WG LTS) is organized with the goal of developing a better understanding what "Long Term Support" might mean for Kubernetes, those who support Kubernetes, and end users, as well as to investigate changes the Kubernetes project could make related to long term support. The working group will also determine feasibility, benefits, **costs**, and prerequisites of such changes.

In the first phase of the working group, we will collect information related to the needs and wants regarding support periods from end-users and people supporting Kubernetes. This will be accomplished by reaching out to users and cluster operators (e.g., through surveys) in order to gain a better understanding, including, but not restricted to:

* in-use versions (in the form of a user/vendor survey) and reasons for remaining on those versions
* constraints on deployment and upgrade patterns / timelines (e.g. edge deployments, regulated industries, retail, etc)
* expected/required support periods from users/vendors
* core Kubernetes dependencies and their support periods
* support periods of other components required to run Kubernetes clusters (OS, network, storage, etc)
* what users/vendors are currently doing to support releases past community support End Of Life (EOL)
* what do users/vendors do with EOL clusters today? what would they do if EOL was extended by N months/years?

With this information, we will investigate changes the Kubernetes project could make to address these needs. A non-exhaustive set of changes to investigate include:
* lengthening support period for all minor versions
* lengthening support period for specific minor versions
* applying security fixes to release branches past current EOL without cutting additional patch releases
* applying security and "critical" fixes to release branches past current EOL without cutting additional patch releases
* expanding supported skew
* improving supported upgrade patterns for clusters at EOL

### In Scope

- Collecting input to better define long term support with regard to Kubernetes releases. This could include:
  * What is a supported release?
  * Number of community supported branches.
  * Duration of community support per supported branch.
- Working with vendors and other community members to identify ongoing efforts to help end users beyond current community support and identify opportunities for reuse of prior work.
- Creating and prioritizing a list of areas that require investments to improve long term supportability. This could include:
  * Upgrade path considerations.
  * Costs of Kubernetes releases in terms of:
    * Infrastructure
    * People
- Initiating and driving cross-SIG changes related to long term support.
- Identifing ways to that vendors and other community members can better support Kubernetes releases.

### Out of scope

- The lifecycle of projects outside of the Kubernetes org.
- Designing and executing on changes clearly falling into individual SIG
  responsibilities. This is a working group, no code implementation is the responsibility of this Working Group.
- Technical and end-user support:  The WG may make recommendations
  around support to those responsible for relevant code and responsible
  for the release engineering operations and automation, but does not
  own code itself.

## Special Powers

None

## Stakeholders

Any changes identified by this working group will require involvement and investment from the following SIGS:

- SIG Architecture
  High-level input on requirements.
- SIG Cluster Lifecycle
  Input on cluster upgrade mechanics.
- SIG K8s Infra
  Input on infrastructure costs related to increased support period(s).
- SIG Release
  Input on maintaining older branches and additional releases.
- SIG Security
  Input on vulnerability management process and audits for additional releases.
- SIG Testing
  Input on testing impacts related to upgrades and longer test periods. 
  
Improvements identified will likely involve every SIG, but the list above are identified as the primary stakeholders.

## Deliverables

The artifacts the group is supposed to deliver include:
- Survey results better describing Kubernetes version use patterns, deployment constraints, and upgrade patterns.
- Recommend changes Kubernetes SIGs can make that will provide broad benefits in a sustainable/affordable way. This would likely take the form of one or more KEPs.
- Recommend ways users/vendors who want to maximize Kubernetes support can consume Kubernetes

Any changes identified will be owned by corresponding SIGs.

## Roles and Organization Management

This working group adheres to the Roles and Organization Management outlined in
[sig-governance] and opts-in to updates and modifications to [sig-governance].

[sig-governance]: https://github.com/kubernetes/community/blob/master/committee-steering/governance/sig-governance.md

## Timelines and Disbanding

The exact timeline for this working group is hard to define at this time. If we are unable to define a common long term support definition for Kubernetes, we will disband the working group. If we are unable to define improvements related to an agreed upon definition, we will also disband the working group. In order to evaluate our progress toward the working group goals, we will provide periodic updates to the stakeholder SIGs and committees at least every six months, in addition to a working group annual report.

Additionally, if the working group determines that the Kubernetes project does not have sufficient resources, nor commmitments for future resources to support relevant proposed changes, the working group will also disband.