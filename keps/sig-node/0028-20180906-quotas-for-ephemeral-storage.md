---
kep-number: 0
title: My First KEP
authors:
  - "@janedoe"
owning-sig: sig-xxx
participating-sigs:
  - sig-aaa
  - sig-bbb
reviewers:
  - TBD
  - "@alicedoe"
approvers:
  - TBD
  - "@oscardoe"
editor: TBD
creation-date: yyyy-mm-dd
last-updated: yyyy-mm-dd
status: provisional
see-also:
  - KEP-1
  - KEP-2
replaces:
  - KEP-3
superseded-by:
  - KEP-100
---

# Quotas for Ephemeral Storaeg

## Table of Contents

A table of contents is helpful for quickly jumping to sections of a KEP and for highlighting any additional information provided beyond the standard KEP template.
[Tools for generating][] a table of contents from markdown are available.

* [Table of Contents](#table-of-contents)
* [Summary](#summary)
* [Motivation](#motivation)
    * [Goals](#goals)
    * [Non-Goals](#non-goals)
* [Proposal](#proposal)
    * [User Stories [optional]](#user-stories-optional)
      * [Story 1](#story-1)
      * [Story 2](#story-2)
    * [Implementation Details/Notes/Constraints [optional]](#implementation-detailsnotesconstraints-optional)
    * [Risks and Mitigations](#risks-and-mitigations)
* [Graduation Criteria](#graduation-criteria)
* [Implementation History](#implementation-history)
* [Drawbacks [optional]](#drawbacks-optional)
* [Alternatives [optional]](#alternatives-optional)

[Tools for generating]: https://github.com/ekalinin/github-markdown-toc

## Summary

Local storage capacity isolation, aka ephemeral-storage, was introduced into Kubernetes via https://github.com/kubernetes/features/issues/361.  It provides support for capacity isolation of shared storage between pods, such that a pod can be limited in its consumption of shared resources and can be evicted if its consumption of shared storage exceeds that limit.  The limits and requests for shared ephemeral-storage are similar to those for memory and CPU consumption.
The current mechanism relies on periodically walking each ephemeral volume (emptydir, logdir, or container writable layer) and summing the space consumption.  This method is slow, can be fooled, and has high latency (i. e. a pod could consume a lot of storage prior to the kubelet being aware of its overage and terminating it).
The mechanism proposed here utilizes filesystem project quotas to provide monitoring of resource consumption and optionally enforcement of limits.  Project quotas, initially in XFS and more recently ported to ext4fs, offer a kernel-based means of restricting and monitoring filesystem consumption that can be applied to one or more directories.

## Motivation

The mechanism presently used to monitor storage consumption involves use of `du` and `find` to periodically gather information about storage and inode consumption of volumes.  This mechanism suffers from a number of drawbacks:

* It is slow.  If a volume contains a large number of files, walking the directory can take a significant amount of time.  There has been at least one known report of nodes becoming not ready due to volume metrics: https://github.com/kubernetes/kubernetes/issues/62917
* It is possible to conceal a file from the walker by creating it and removing it while holding an open file descriptor on it.  POSIX behavior is to not remove the file until the last open file descriptor pointing to it is removed.  This has legitimate uses; it ensures that a temporary file is deleted when the processes using it exit, and it minimizes the attack surface by not having a file that can be found by an attacker.  The following pod does this; it will never be caught by the present mechanism:
```yaml
apiVersion: v1
kind: Pod
max:
metadata:
  name: "diskhog"
spec:
  containers:
  - name: "perl"
    resources:
      limits:
        ephemeral-storage: "2048Ki"
    image: "perl"
    command:
    - perl
    - -e
    - >
      my $file = "/data/a/a"; open OUT, ">$file" or die "Cannot open $file: $!\n"; unlink "$file" or die "cannot unlink $file: $!\n"; my $a="0123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789"; foreach my $i (0..200000000) { print OUT $a; }; sleep 999999
    volumeMounts:
    - name: a
      mountPath: /data/a
  volumes:
  - name: a
    emptyDir: {}
```
* It is reactive rather than proactive.  It does not prevent a pod from overshooting its limit; at best it catches it after the fact.  On a fast storage medium, such as NVMe, a pod may write 50 GB or more of data before the housekeeping performed once per minute catches up to it.  If the primary volume is the root partition, this will completely fill the partition, possibly causing serious problems elsewhere on the system.

In many environments, these issues may not matter, but shared multi-tenant environments need these issues addressed.

### Goals

* Primary: improve performance of monitoring by using project quotas in a non-enforcing way to collect information about storage utilization.
* Primary: detect storage used by pods that is concealed by deleted files being held open.
* Primary: this will not interfere with the more common user and group quotas.
* Stretch: enforce limits on per-volume storage consumption by using enforced project quotas.  Each volume would be given an enforced quota of the total ephemeral storage limit of the pod.

### Non-Goals

* Enforcing limits on total pod storage consumption by any means, such that the pod would be hard restricted to the desired storage limit.

## Proposal

This proposal applies project quotas to emptydir volumes on qualifying filesystems (ext4fs and xfs with project quotas enabled).  Project quotas are applied by selecting an unused project ID (a 32-bit unsigned integer), setting a limit on space and/or inode consumption, and attaching the ID to one or more files.  By default (and as utilized herein), if a project ID is attached to a directory, it is inherited by any files created under that directory.
If we elect to use the quota as enforcing, we impose a quota consistent with the desired limit.  If we elect to use it as non-enforcing, we impose a large quota that in practice cannot be exceeded (2^61-1 bytes for XFS, 2^58-1 bytes for ext4fs).

### Operation Flow -- Applying a Quota

* Caller (emptydir volume manager or container runtime) creates an emptydir volume, with an empty directory at a location of its choice.
* Caller requests that a quota be applied to a directory.
* Determine whether a quota can be imposed on the directory, by asking each quota provider (one per filesystem type) whether it can apply a quota to the directory.  If no provider claims the directory, an error status is returned to the caller.
* Select an unused project ID (see below).
* Set the desired limit on the project ID, in a filesystem-dependent manner.
* Apply the project ID to the directory in question, in a filesystem-dependent manner.

An error at any point results in no quota being applied and no change to the state of the system.  The caller in general should not assume a priori that the attempt will be successful.  It could choose to reject a request if a quota cannot be applied, but at this time it will simply ignore the error and proceed as today.

### Operation Flow -- Retrieving Storage Consumption

* Caller (kubelet metrics code, cadvisor, container runtime) asks the quota code to compute the amount of storage used under the directory.
* Determine whether a quota applies to the directory, in a filesystem-dependent manner (see below).
* If so, determine how much storage or how many inodes are utilized, in a filesystem dependent manner.

If the quota code is unable to retrieve the consumption, it returns an error status and it is up to the caller to utilize a fallback mechanism (such as the directory walk performed today).

### Operation Flow -- Removing a Quota.

* Caller requests that the quota be removed from a directory.
* Determine whether a project quota applies to the directory.
* Remove the limit from the project ID associated with the directory.
* Remove the association between the directory and the project ID.
* Return the project ID to the system to allow its use elsewhere (see below).
* Caller may delete the directory and its contents (normally it will).

### Operation Notes

#### Selecting a Project ID

Project IDs are a shared space within a filesystem.  If the same project ID is assigned to multiple directories, the space consumption reported by the quota will be the sum of that of all of the directories.  Hence, it is important to ensure that each directory is assigned a unique project ID (unless it is desired to pool the storage use of multiple directories).

The canonical mechanism to record persistently that a project ID is reserved is to store it in the /etc/projid (projid(5)) and/or /etc/projects (projects(5)) files.  However, it is possible to utilize project IDs without recording them in those files; they exist for administrative convenience but neither the kernel nor the filesystem is aware of them.  Other ways can be used to determine whether a project ID is in active use on a given filesystem:

* The quota values (in blocks and/or inodes) assigned to the project ID are non-zero.
* The storage consumption (in blocks and/or inodes) reported under the project ID are non-zero.

The algorithm to be used is as follows:

* Lock this instance of the quota code against re-entrancy.
* open and flock() the /etc/project and /etc/projid files, so that other uses of this code are excluded.
* Start from a high number (the prototype uses 1048577).
* Iterate from there, performing the following tests:
   * Is the ID reserved by this instance of the quota code?
   * Is the ID present in /etc/projects?
   * Is the ID present in /etc/projid?
   * Are the quota values and/or consumption reported by the kernel non-zero?  This test is restricted to 128 iterations to ensure that a bug here or elsewhere does not result in an infinite loop looking for a quota ID.
* If an ID has been found:
   * Add it to an in-memory copy of /etc/projects and /etc/projid so that any other uses of project quotas do not reuse it.
   * Write temporary copies of /etc/projects and /etc/projid that are flock()ed
   * If successful, rename the temporary files appropriately (if rename of one succeeds but the other fails, we have a problem that we cannot recover from, and the files may be inconsistent).
* Unlock /etc/projid and /etc/projects.
* Unlock this instance of the quota code.

A minor variation of this is used if we want to reuse an existing quota ID.

#### Determine Whether a Project ID Applies To a Directory

It is possible to determine whether a directory has a project ID applied to it by requesting (via the quotactl(2) system call) the project ID associated with the directory.  Whie the specifics are filesystem-dependent, the basic method is the same for at least XFS and ext4fs.

It is not possible to directly determine the directory or directories to which a project ID is applied.  It is possible to determine whether a project ID has been applied to an existing directory or files; the reported consumption will be non-zero.

The code records internally the project ID applied to a directory, but it cannot always rely on this.  In particular, if the kubelet has exited and has been restarted, the map from directory to project ID is lost.  If it cannot find a map entry, it falls back on the approach discussed above.

#### Return a Project ID To the System

The algorithm used to return a project ID to the system is very similar to the algorithm used to select a project ID, except of course for selecting a project ID.  It performs the same sequence of locking /etc/project and /etc/projid, editing a copy of the file, and restoring it.

If the project ID is applied to multiple directories and the code can determine that, it will not remove the project ID from /etc/projid until the last reference is removed.  While it is not anticipated that this mode of operation will be used, at least initially, this can be detected even on kubelet restart by looking at the reference count in /etc/projects.


### Implementation Details/Notes/Constraints [optional]

What are the caveats to the implementation?
What are some important details that didn't come across above.
Go in to as much detail as necessary here.
This might be a good place to talk about core concepts and how they releate.

### Risks and Mitigations

What are the risks of this proposal and how do we mitigate.
Think broadly.
For example, consider both security and how this will impact the larger kubernetes ecosystem.

## Graduation Criteria

How will we know that this has succeeded?
Gathering user feedback is crucial for building high quality experiences and SIGs have the important responsibility of setting milestones for stability and completeness.
Hopefully the content previously contained in [umbrella issues][] will be tracked in the `Graduation Criteria` section.

[umbrella issues]: https://github.com/kubernetes/kubernetes/issues/42752

## Implementation History

Major milestones in the life cycle of a KEP should be tracked in `Implementation History`.
Major milestones might include

- the `Summary` and `Motivation` sections being merged signaling SIG acceptance
- the `Proposal` section being merged signaling agreement on a proposed design
- the date implementation started
- the first Kubernetes release where an initial version of the KEP was available
- the version of Kubernetes where the KEP graduated to general availability
- when the KEP was retired or superseded

## Drawbacks [optional]

Why should this KEP _not_ be implemented.

## Alternatives [optional]

Similar to the `Drawbacks` section the `Alternatives` section is used to highlight and record other possible approaches to delivering the value proposed by a KEP.

## Infrastructure Needed [optional]

Use this section if you need things from the project/SIG.
Examples include a new subproject, repos requested, github details.
Listing these here allows a SIG to get the process for these resources started right away.