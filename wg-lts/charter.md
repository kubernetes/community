# WG Long Term Support

This charter adheres to the conventions described in the [Kubernetes Charter README]
and uses the Roles and Organization Management outlined in [sig-governance].

[sig-governance]: https://github.com/kubernetes/community/blob/master/committee-steering/governance/sig-governance.md
[Kubernetes Charter README]: https://github.com/kubernetes/community/blob/master/committee-steering/governance/README.md

## Scope

The Long Term Support Working Group (WG LTS) is organized with the goal of developing a better understanding what "Long Term Support" might mean for Kubernetes, those who support Kubernetes, and end users, as well as to investigate changes the Kubernetes project could make related to long term support. The working group will also determine feasibility, benefits, **costs**, and prerequisites of such changes.

In the first phase of the working group, we will collect information related to the needs and wants regarding support periods from end-users and people supporting Kubernetes. This will be done by reaching out to users and cluster operators (e.g. via surveys), to better understand:

* in-use versions (in the form of a user/vendor survey) and reasons for remaining on those versions
* constraints on deployment and upgrade patterns / timelines (e.g. edge deployments, regulated industries, retail, etc)
* expected/required support periods from users/vendors
* support periods of core Kubernetes dependencies
* support periods of other components required to run Kubernetes clusters (OS, network, storage, etc)
* what users/vendors are currently doing to support releases past OSS End Of Life (EOL)
* what do users/vendors do with EOL clusters today? what would they do if EOL was extended by N months/years?

With this information, we will investigate changes the Kubernetes project could make to address these needs. A non-exhaustive set of changes to investigate include:
* lengthening support period for all minor versions
* lengthening support period for specific minor versions
* applying security fixes to release branches past current EOL without cutting additional patch releases
* applying security and "critical" fixes to release branches past current EOL without cutting additional patch releases
* expanding supported skew
* improving supported upgrade patterns for clusters at EOL

### In Scope

- Collecting input to better define long term support with regard to Kubernetes OSS releases
- Creating and prioritizing a list of areas that require investments to improve long term supportability
- Initiate and drive cross-SIG changes related to long term support

### Out of scope

- Designing and executing on changes clearly falling into individual SIG
  responsibilities.

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
- SIG Testing
  Input on testing impacts related to upgrades and longer test periods. 
  
Improvements identified will likely involve every SIG, but the list above are identified as the primary stakeholders.

## Deliverables

The artifacts the group is supposed to deliver include:
- Survey results better describing Kubernetes version use patterns, deployment constraints, and upgrade patterns.
- Recommend changes Kubernetes SIGS can make that will provide broad benefits in a sustainable/affordable way. This would likely take the form of one or more KEPs.
- Recommend ways users/vendors who want to maximize Kubernetes support can consume Kubernetes

Any changes identified will be owned by corresponding SIGs.

## Roles and Organization Management

This working group adheres to the Roles and Organization Management outlined in
[sig-governance] and opts-in to updates and modifications to [sig-governance].

[sig-governance]: https://github.com/kubernetes/community/blob/master/committee-steering/governance/sig-governance.md

## Timelines and Disbanding

The exact timeline for this working group is hard to define at this time. If we are unable to define a common long term support definition for Kubernetes, we will disband the working group. If we are unable to define improvements related to an agreed upon definition, we will also disband the working group.
