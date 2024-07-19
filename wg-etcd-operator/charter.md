# WG etcd operator

This charter adheres to the conventions described in the [Kubernetes Charter README]
and uses the Roles and Organization Management outlined in [sig-governance].

[sig-governance]: https://github.com/kubernetes/community/blob/master/committee-steering/governance/sig-governance.md
[Kubernetes Charter README]: https://github.com/kubernetes/community/blob/master/committee-steering/governance/README.md

## Scope

The purpose of an etcd-operator is to operate automatically etcd clusters which run in the Kubernetes environment.
It minimizes human intervention as much as possible. Note it excludes the case of etcd backing Kubernetes cluster;
instead, etcd runs as POD as normal workloads.

### In Scope

- Collect requirements & use cases with a [survey](https://forms.gle/5gBpzaxYtuQPWdBo9) to better understand what users care about the most.
- Prioritize the tasks based on feedback and create a roadmap.
- Bootstrap a project "etcd-operator" owned by SIG etcd which resides in the etcd-io or kubernetes-sigs Github orgs.
  - Review existing etcd operators to see if any could be forked or referenced to advance the project.
- Discuss and design the core reconciliation workflow, and potentially provide a proof of concept (PoC).
- Figure out how to get resource for following dev/test, i.e. AWS S3.

### Out of scope

- Manage etcd clusters running within non-Kubernetes environments.
- Manage etcd clusters which are used as the storage backend of a host (non-nested) kube-apiserver.

## Stakeholders

Stakeholders for this working group include members in the following SIGs:

- SIG etcd
- SIG Cluster Lifecycle

## Deliverables

The artifacts the group is supposed to deliver include:
- Survey results which describe the users requirements and use cases.
- Evaluation results on existing etcd-operators.
- Roadmap for the project etcd-operator.
- Core reconciliation workflow and PoC.
- A new repository "etcd-operator" owned by SIG etcd, and it should have implemented the very basic functionalities below:
  - Creation of a new etcd cluster with one or multiple members.
  - Scale out and in the etcd cluster.
  - Upgrading patch versions or one minor version.

## Roles and Organization Management

This working group adheres to the Roles and Organization Management outlined in
[sig-governance] and opts-in to updates and modifications to [sig-governance].

[sig-governance]: https://github.com/kubernetes/community/blob/master/committee-steering/governance/sig-governance.md

## Timelines and Disbanding

When all the deliverables mentioned above are done and there is no additional coordination needed,
then we will disband this working group and continue to track the development of the project
under SIG etcd.
