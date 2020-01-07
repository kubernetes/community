# WG Data Protection Charter

This charter adheres to the [wg-governance] guidance as well as
the general conventions described in the [Kubernetes Charter README] and
the Roles and Organization Management outlined in [sig-governance], where
applicable to a Working Group.


## Scope

The purpose of Data Protection is to ensure that application and its associated
data can be restored quickly after any corruption or loss.

Data protection in Kubernetes context typically involves backup and recovery
of two types of entities:
* Kubernetes API object resources
* Persistent volume data
We consider it a complicated and layered problem, including backup and recovery
at persistent volume level, application level, and cluster level. Part of the
working group’s charter is to define what Kubernetes native constructs are
required to achieve these goals.

The Data Protection Working Group is organized with the goal of providing
a cross SIG forum to discuss how to support data protection in Kubernetes,
identify missing functionality, and work together to design features that
are needed to achieve the goal.

### In scope

* High level discussions on what it means to support data protection in Kubernetes at different levels and how to do it.
* Design discussions on specific topics related to data protection and disaster recovery support.
* Document results of discussions and investigations in a linkable medium.

Potential design topics include, but are not limited to the following:
* Read data from a snapshot without creating a new volume.
* Volume backups
* Data populator
* Retrieve diffs between two snapshots (block and file level)
* Consistency volume groups (group snapshot)
* Application snapshot, backup, and recovery
* Data protection policy (Data protection policy usually means we can set up a schedule to do
  periodic backups, set a backup retention policy to automatically clean up old backups, set a
  topology to specify a backup location, etc. It can also possibly include policies such as
  `backups must be encrypted` and `secrets must be encrypted at rest and in transit`.)
* Data protection workflows

### Out of Scope

* Design discussions not related to data protection is out of scope. For example,
  how to migrate in-tree drivers to CSI drivers and how to report volume health
  belong to SIG Storage and would not be a focus area of this WG. Workload API designs
  for StatefulSet and Deployment belong to SIG Apps, however, this WG would be interested
  in figuring out how to backup and recover a StatefulSet or Deployment.
* This is a working group so it does not own code. Design discussions for
  a specific feature including KEP reviews can happen in the working group
  but KEP approvals and code implementation will be owned by SIG-Apps or
  SIG-Storage.
* Security and privacy protection is an important aspect but not a focus of this WG. This WG will consult with [CNCF SIG-Security](https://github.com/cncf/sig-security) for those issues.


## Stakeholders

Stakeholders for this working group include members in the following SIGs:
* SIG Apps
* SIG Storage

We will also consult SIG Auth from security aspect. Stakeholders also include
backup vendors who want to provide data protection support in Kubernetes and
end users who want to use data protection applications.


## Disband criteria

This WG will be producing documents as described in the `In Scope` section. If stakeholder SIGs and the WG decide all documents described in the `In Scope` section are complete and no more discussions and investigations are needed in this WG, they may determine to disband this WG.


[sig-governance]: https://github.com/kubernetes/community/blob/master/committee-steering/governance/sig-governance.md
[wg-governance]: https://github.com/kubernetes/community/blob/master/committee-steering/governance/wg-governance.md
[Kubernetes Charter README]: https://github.com/kubernetes/community/blob/master/committee-steering/governance/README.md
[lazy consensus]: http://en.osswiki.info/concepts/lazy_consensus

[kubernetes-dev@]: https://groups.google.com/forum/#!forum/kubernetes-dev
[wg-data-protection@]: https://groups.google.com/forum/#!forum/kubernetes-wg-data-protection
[kubernetes/k8s.io]: https://git.k8s.io/k8s.io
