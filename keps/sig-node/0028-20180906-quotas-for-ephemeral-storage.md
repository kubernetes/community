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

# Quotas for Ephemeral Storage

## Table of Contents
<!-- markdown-toc start - Don't edit this section. Run M-x markdown-toc-generate-toc again -->
**Table of Contents**

- [Quotas for Ephemeral Storage](#quotas-for-ephemeral-storage)
    - [Table of Contents](#table-of-contents)
    - [Summary](#summary)
        - [Project Quotas](#project-quotas)
    - [Motivation](#motivation)
        - [Goals](#goals)
        - [Non-Goals](#non-goals)
        - [Future Work](#future-work)
    - [Proposal](#proposal)
        - [Control over Use of Quotas](#control-over-use-of-quotas)
        - [Operation Flow -- Applying a Quota](#operation-flow----applying-a-quota)
        - [Operation Flow -- Retrieving Storage Consumption](#operation-flow----retrieving-storage-consumption)
        - [Operation Flow -- Removing a Quota.](#operation-flow----removing-a-quota)
        - [Operation Notes](#operation-notes)
            - [Selecting a Project ID](#selecting-a-project-id)
            - [Determine Whether a Project ID Applies To a Directory](#determine-whether-a-project-id-applies-to-a-directory)
            - [Return a Project ID To the System](#return-a-project-id-to-the-system)
        - [Implementation Details/Notes/Constraints [optional]](#implementation-detailsnotesconstraints-optional)
            - [Notes on Implementation](#notes-on-implementation)
            - [Notes on Code Changes](#notes-on-code-changes)
            - [Testing Strategy](#testing-strategy)
        - [Risks and Mitigations](#risks-and-mitigations)
    - [Graduation Criteria](#graduation-criteria)
    - [Implementation History](#implementation-history)
    - [Drawbacks [optional]](#drawbacks-optional)
    - [Alternatives [optional]](#alternatives-optional)
        - [Alternative quota-based implementation](#alternative-quota-based-implementation)
        - [Alternative loop filesystem-based implementation](#alternative-loop-filesystem-based-implementation)
    - [Infrastructure Needed [optional]](#infrastructure-needed-optional)
    - [References](#references)
        - [Bugs Opened Against Filesystem Quotas](#bugs-opened-against-filesystem-quotas)
            - [CVE](#cve)
            - [Other Security Issues Without CVE](#other-security-issues-without-cve)
        - [Other Linux Quota-Related Bugs Since 2012](#other-linux-quota-related-bugs-since-2012)

<!-- markdown-toc end -->

[Tools for generating]: https://github.com/ekalinin/github-markdown-toc

## Summary

This proposal applies to the use of quotas for ephemeral-storage
metrics gathering.  Use of quotas for ephemeral-storage limit
enforcement is a [non-goal](#non-goals), but as the architecture and
code will be very similar, there are comments interspersed related to
enforcement.  _These comments will be italicized_.

Local storage capacity isolation, aka ephemeral-storage, was
introduced into Kubernetes via
<https://github.com/kubernetes/features/issues/361>.  It provides
support for capacity isolation of shared storage between pods, such
that a pod can be limited in its consumption of shared resources and
can be evicted if its consumption of shared storage exceeds that
limit.  The limits and requests for shared ephemeral-storage are
similar to those for memory and CPU consumption.

The current mechanism relies on periodically walking each ephemeral
volume (emptydir, logdir, or container writable layer) and summing the
space consumption.  This method is slow, can be fooled, and has high
latency (i. e. a pod could consume a lot of storage prior to the
kubelet being aware of its overage and terminating it).

The mechanism proposed here utilizes filesystem project quotas to
provide monitoring of resource consumption _and optionally enforcement
of limits._  Project quotas, initially in XFS and more recently ported
to ext4fs, offer a kernel-based means of monitoring _and restricting_
filesystem consumption that can be applied to one or more directories.

A prototype is in progress; see <https://github.com/kubernetes/kubernetes/pull/66928>.

### Project Quotas

Project quotas are a form of filesystem quota that apply to arbitrary
groups of files, as opposed to file user or group ownership.  They
were first implemented in XFS, as described here:
<http://xfs.org/docs/xfsdocs-xml-dev/XFS_User_Guide/tmp/en-US/html/xfs-quotas.html>.

Project quotas for ext4fs were [proposed in late
2014](https://lwn.net/Articles/623835/) and added to the Linux kernel
in early 2016, with
commit
[391f2a16b74b95da2f05a607f53213fc8ed24b8e](https://git.kernel.org/pub/scm/linux/kernel/git/torvalds/linux.git/commit/?id=391f2a16b74b95da2f05a607f53213fc8ed24b8e).
They were designed to be compatible with XFS project quotas.

Each inode contains a 32-bit project ID, to which optionally quotas
(hard and soft limits for blocks and inodes) may be applied.  The
total blocks and inodes for all files with the given project ID are
maintained by the kernel.  Project quotas can be managed from
userspace by means of the `xfs_quota(8)` command in foreign filesystem
(`-f`) mode; the traditional Linux quota tools do not manipulate
project quotas.  Programmatically, they are managed by the `quotactl(2)`
system call, using in part the standard quota commands and in part the
XFS quota commands; the man page implies incorrectly that the XFS
quota commands apply only to XFS filesystems.

The project ID applied to a directory is inherited by files created
under it.  Files cannot be (hard) linked across directories with
different project IDs.  A file's project ID cannot be changed by a
non-privileged user, but a privileged user may use the `xfs_io(8)`
command to change the project ID of a file.

Filesystems using project quotas may be mounted with quotas either
enforced or not; the non-enforcing mode tracks usage without enforcing
it.  A non-enforcing project quota may be implemented on a filesystem
mounted with enforcing quotas by setting a quota too large to be hit.
The maximum size that can be set varies with the filesystem; on a
64-bit filesystem it is 2^63-1 bytes for XFS and 2^58-1 bytes for
ext4fs.

Conventionally, project quota mappings are stored in `/etc/projects` and
`/etc/projid`; these files exist for user convenience and do not have
any direct importance to the kernel.  `/etc/projects` contains a mapping
from project ID to directory/file; this can be a one to many mapping
(the same project ID can apply to multiple directories or files, but
any given directory/file can be assigned only one project ID).
`/etc/projid` contains a mapping from named projects to project IDs.

This proposal utilizes hard project quotas for both monitoring _and
enforcement_.  Soft quotas are of no utility; they allow for temporary
overage that, after a programmable period of time, is converted to the
hard quota limit.


## Motivation

The mechanism presently used to monitor storage consumption involves
use of `du` and `find` to periodically gather information about
storage and inode consumption of volumes.  This mechanism suffers from
a number of drawbacks:

* It is slow.  If a volume contains a large number of files, walking
  the directory can take a significant amount of time.  There has been
  at least one known report of nodes becoming not ready due to volume
  metrics: <https://github.com/kubernetes/kubernetes/issues/62917>
* It is possible to conceal a file from the walker by creating it and
  removing it while holding an open file descriptor on it.  POSIX
  behavior is to not remove the file until the last open file
  descriptor pointing to it is removed.  This has legitimate uses; it
  ensures that a temporary file is deleted when the processes using it
  exit, and it minimizes the attack surface by not having a file that
  can be found by an attacker.  The following pod does this; it will
  never be caught by the present mechanism:

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
* It is reactive rather than proactive.  It does not prevent a pod
  from overshooting its limit; at best it catches it after the fact.
  On a fast storage medium, such as NVMe, a pod may write 50 GB or
  more of data before the housekeeping performed once per minute
  catches up to it.  If the primary volume is the root partition, this
  will completely fill the partition, possibly causing serious
  problems elsewhere on the system.  This proposal does not address
  this issue; _a future enforcing project would_.

In many environments, these issues may not matter, but shared
multi-tenant environments need these issues addressed.

### Goals

These goals apply only to local ephemeral storage, as described in
<https://github.com/kubernetes/features/issues/361>.

* Primary: improve performance of monitoring by using project quotas
  in a non-enforcing way to collect information about storage
  utilization of ephemeral volumes.
* Primary: detect storage used by pods that is concealed by deleted
  files being held open.
* Primary: this will not interfere with the more common user and group
  quotas.

### Non-Goals

* Application to storage other than local ephemeral storage.
* Elimination of eviction as a means of enforcing ephemeral-storage
  limits.  Pods that hit their ephemeral-storage limit will still be
  evicted by the kubelet even if their storage has been capped by
  enforcing quotas.
* Enforcing node allocatable (limit over the sum of all pod's disk
  usage, including e. g. images).
* Enforcing limits on total pod storage consumption by any means, such
  that the pod would be hard restricted to the desired storage limit.
  
### Future Work

* _Enforce limits on per-volume storage consumption by using
  enforced project quotas._

## Proposal

This proposal applies project quotas to emptydir volumes on qualifying
filesystems (ext4fs and xfs with project quotas enabled).  Project
quotas are applied by selecting an unused project ID (a 32-bit
unsigned integer), setting a limit on space and/or inode consumption,
and attaching the ID to one or more files.  By default (and as
utilized herein), if a project ID is attached to a directory, it is
inherited by any files created under that directory.

_If we elect to use the quota as enforcing, we impose a quota
consistent with the desired limit._  If we elect to use it as
non-enforcing, we impose a large quota that in practice cannot be
exceeded (2^63-1 bytes for XFS, 2^58-1 bytes for ext4fs).

### Control over Use of Quotas

At present, two feature gates control operation of quotas:

* `LocalStorageCapacityIsolation` must be enabled for any use of
  quotas.
  
* `FSQuotaForLSCIMonitoring` must be enabled in addition.  If this is
  enabled, quotas are used for monitoring, but not enforcement.  At
  present, this defaults to False, but the intention is that this will
  default to True by initial release.
  
* _`FSQuotaForLSCIEnforcement` must be enabled, in addition to
  `FSQuotaForLSCIMonitoring`, to use quotas for enforcement._

### Operation Flow -- Applying a Quota

* Caller (emptydir volume manager or container runtime) creates an
  emptydir volume, with an empty directory at a location of its
  choice.
* Caller requests that a quota be applied to a directory.
* Determine whether a quota can be imposed on the directory, by asking
  each quota provider (one per filesystem type) whether it can apply a
  quota to the directory.  If no provider claims the directory, an
  error status is returned to the caller.
* Select an unused project ID ([see below](#selecting-a-project-id)).
* Set the desired limit on the project ID, in a filesystem-dependent
  manner ([see below](#notes-on-implementation)).
* Apply the project ID to the directory in question, in a
  filesystem-dependent manner.

An error at any point results in no quota being applied and no change
to the state of the system.  The caller in general should not assume a
priori that the attempt will be successful.  It could choose to reject
a request if a quota cannot be applied, but at this time it will
simply ignore the error and proceed as today.

### Operation Flow -- Retrieving Storage Consumption

* Caller (kubelet metrics code, cadvisor, container runtime) asks the
  quota code to compute the amount of storage used under the
  directory.
* Determine whether a quota applies to the directory, in a
  filesystem-dependent manner ([see below](#notes-on-implementation)).
* If so, determine how much storage or how many inodes are utilized,
  in a filesystem dependent manner.

If the quota code is unable to retrieve the consumption, it returns an
error status and it is up to the caller to utilize a fallback
mechanism (such as the directory walk performed today).

### Operation Flow -- Removing a Quota.

* Caller requests that the quota be removed from a directory.
* Determine whether a project quota applies to the directory.
* Remove the limit from the project ID associated with the directory.
* Remove the association between the directory and the project ID.
* Return the project ID to the system to allow its use elsewhere ([see
  below](#return-a-project-id-to-the-system)).
* Caller may delete the directory and its contents (normally it will).

### Operation Notes

#### Selecting a Project ID

Project IDs are a shared space within a filesystem.  If the same
project ID is assigned to multiple directories, the space consumption
reported by the quota will be the sum of that of all of the
directories.  Hence, it is important to ensure that each directory is
assigned a unique project ID (unless it is desired to pool the storage
use of multiple directories).

The canonical mechanism to record persistently that a project ID is
reserved is to store it in the `/etc/projid` (`projid[5]`) and/or
`/etc/projects` (`projects(5)`) files.  However, it is possible to utilize
project IDs without recording them in those files; they exist for
administrative convenience but neither the kernel nor the filesystem
is aware of them.  Other ways can be used to determine whether a
project ID is in active use on a given filesystem:

* The quota values (in blocks and/or inodes) assigned to the project
  ID are non-zero.
* The storage consumption (in blocks and/or inodes) reported under the
  project ID are non-zero.

The algorithm to be used is as follows:

* Lock this instance of the quota code against re-entrancy.
* open and `flock()` the `/etc/project` and `/etc/projid` files, so that
  other uses of this code are excluded.
* Start from a high number (the prototype uses 1048577).
* Iterate from there, performing the following tests:
   * Is the ID reserved by this instance of the quota code?
   * Is the ID present in `/etc/projects`?
   * Is the ID present in `/etc/projid`?
   * Are the quota values and/or consumption reported by the kernel
     non-zero?  This test is restricted to 128 iterations to ensure
     that a bug here or elsewhere does not result in an infinite loop
     looking for a quota ID.
* If an ID has been found:
   * Add it to an in-memory copy of `/etc/projects` and `/etc/projid` so
     that any other uses of project quotas do not reuse it.
   * Write temporary copies of `/etc/projects` and `/etc/projid` that are
     `flock()`ed
   * If successful, rename the temporary files appropriately (if
     rename of one succeeds but the other fails, we have a problem
     that we cannot recover from, and the files may be inconsistent).
* Unlock `/etc/projid` and `/etc/projects`.
* Unlock this instance of the quota code.

A minor variation of this is used if we want to reuse an existing
quota ID.

#### Determine Whether a Project ID Applies To a Directory

It is possible to determine whether a directory has a project ID
applied to it by requesting (via the `quotactl(2)` system call) the
project ID associated with the directory.  Whie the specifics are
filesystem-dependent, the basic method is the same for at least XFS
and ext4fs.

It is not possible to determine in constant operations the directory
or directories to which a project ID is applied.  It is possible to
determine whether a given project ID has been applied to an existing
directory or files (although those will not be known); the reported
consumption will be non-zero.

The code records internally the project ID applied to a directory, but
it cannot always rely on this.  In particular, if the kubelet has
exited and has been restarted (and hence the quota applying to the
directory should be removed), the map from directory to project ID is
lost.  If it cannot find a map entry, it falls back on the approach
discussed above.

#### Return a Project ID To the System

The algorithm used to return a project ID to the system is very
similar to the algorithm used to select a project ID, except of course
for selecting a project ID.  It performs the same sequence of locking
`/etc/project` and `/etc/projid`, editing a copy of the file, and
restoring it.

If the project ID is applied to multiple directories and the code can
determine that, it will not remove the project ID from `/etc/projid`
until the last reference is removed.  While it is not anticipated in
this KEP that this mode of operation will be used, at least initially,
this can be detected even on kubelet restart by looking at the
reference count in `/etc/projects`.


### Implementation Details/Notes/Constraints [optional]

#### Notes on Implementation

The primary new interface defined is the quota interface in
`pkg/volume/util/quota/quota.go`.  This defines five operations:

* Does the specified directory support quotas?

* Assign a quota to a directory.  If a non-empty pod UID is provided,
  the quota assigned is that of any other directories under this pod
  UID; if an empty pod UID is provided, a unique quota is assigned.
  
* Retrieve the consumption of the specified directory.  If the quota
  code cannot handle it efficiently, it returns an error and the
  caller falls back on existing mechanism.
  
* Retrieve the inode consumption of the specified directory; same
  description as above.
  
* Remove quota from a directory.  If a non-empty pod UID is passed, it
  is checked against that recorded in-memory (if any).  The quota is
  removed from the specified directory.  This can be used even if
  AssignQuota has not been used; it inspects the directory and removes
  the quota from it.  This permits stale quotas from an interrupted
  kubelet to be cleaned up.
  
Two implementations are provided: `quota_linux.go` (for Linux) and
`quota_unsupported.go` (for other operating systems).  The latter
returns an error for all requests.

As the quota mechanism is intended to support multiple filesystems,
and different filesystems require different low level code for
manipulating quotas, a provider is supplied that finds an appropriate
quota applier implementation for the filesystem in question.  The low
level quota applier provides similar operations to the top level quota
code, with two exceptions:

* No operation exists to determine whether a quota can be applied
  (that is handled by the provider).
  
* An additional operation is provided to determine whether a given
  quota ID is in use within the filesystem (outside of `/etc/projects`
  and `/etc/projid`).
  
The two quota providers in the initial implementation are in
`pkg/volume/util/quota/extfs` and `pkg/volume/util/quota/xfs`.  While
some quota operations do require different system calls, a lot of the
code is common, and factored into
`pkg/volume/util/quota/common/quota_linux_common_impl.go`.

#### Notes on Code Changes

The prototype for this project is mostly self-contained within
`pkg/volume/util/quota` and a few changes to
`pkg/volume/empty_dir/empty_dir.go`.  However, a few changes were
required elsewhere:

* The operation executor needs to pass the desired size limit to the
  volume plugin where appropriate so that the volume plugin can impose
  a quota.  The limit is passed as 0 (do not use quotas), _positive
  number (impose an enforcing quota if possible, measured in bytes),_
  or -1 (impose a non-enforcing quota, if possible) on the volume.
  
  This requires changes to
  `pkg/volume/util/operationexecutor/operation_executor.go` (to add
  `DesiredSizeLimit` to `VolumeToMount`),
  `pkg/kubelet/volumemanager/cache/desired_state_of_world.go`, and
  `pkg/kubelet/eviction/helpers.go` (the latter in order to determine
  whether the volume is a local ephemeral one).
  
* The volume manager (in `pkg/volume/volume.go`) changes the
  `Mounter.SetUp` and `Mounter.SetUpAt` interfaces to take a new
  `MounterArgs` type rather than an `FsGroup` (`*int64`).  This is to
  allow passing the desired size and pod UID (in the event we choose
  to implement quotas shared between multiple volumes; [see
  below](#alternative-quota-based-implementation)).  This required
  small changes to all volume plugins and their tests, but will in the
  future allow adding additional data without having to change code
  other than that which uses the new information.
  
#### Testing Strategy

The quota code is by an large not very amendable to unit tests.  While
there are simple unit tests for parsing the mounts file, and there
could be tests for parsing the projects and projid files, the real
work (and risk) involves interactions with the kernel and with
multiple instances of this code (e. g. in the kubelet and the runtime
manager, particularly under stress).  It also requires setup in the
form of a prepared filesystem.  It would be better served by
appropriate end to end tests.

### Risks and Mitigations

* The SIG raised the possibility of a container being unable to exit
  should we enforce quotas, and the quota interferes with writing the
  log.  This can be mitigated by either not applying a quota to the
  log directory and using the du mechanism, or by applying a separate
  non-enforcing quota to the log directory.

  As log directories are write-only by the container, and consumption
  can be limited by other means (as the log is filtered by the
  runtime), I do not consider the ability to write uncapped to the log
  to be a serious exposure.

  Note in addition that even without quotas it is possible for writes
  to fail due to lack of filesystem space, which is effectively (and
  in some cases operationally) indistinguishable from exceeding quota,
  so even at present code must be able to handle those situations.
  
* Filesystem quotas may impact performance to an unknown degree.
  Information on that is hard to come by in general, and one of the
  reasons for using quotas is indeed to improve performance.  If this
  is a problem in the field, merely turning off quotas (or selectively
  disabling project quotas) on the filesystem in question will avoid
  the problem.  Against the possibility that that cannot be done
  (because project quotas are needed for other purposes), we should
  provide a way to disable use of quotas altogether via a feature
  gate.
  
  A report <https://blog.pythonanywhere.com/110/> notes that an
  unclean shutdown on Linux kernel versions between 3.11 and 3.17 can
  result in a prolonged downtime while quota information is restored.
  Unfortunately, [the link referenced
  here](http://oss.sgi.com/pipermail/xfs/2015-March/040879.html) is no
  longer available.

* Bugs in the quota code could result in a variety of regression
  behavior.  For example, if a quota is incorrectly applied it could
  result in ability to write no data at all to the volume.  This could
  be mitigated by use of non-enforcing quotas.  XFS in particular
  offers the `pqnoenforce` mount option that makes all quotas
  non-enforcing.


## Graduation Criteria

How will we know that this has succeeded?  Gathering user feedback is
crucial for building high quality experiences and SIGs have the
important responsibility of setting milestones for stability and
completeness.  Hopefully the content previously contained in [umbrella
issues][] will be tracked in the `Graduation Criteria` section.

[umbrella issues]: N/A

## Implementation History

Major milestones in the life cycle of a KEP should be tracked in
`Implementation History`.  Major milestones might include

- the `Summary` and `Motivation` sections being merged signaling SIG
  acceptance
- the `Proposal` section being merged signaling agreement on a
  proposed design
- the date implementation started
- the first Kubernetes release where an initial version of the KEP was
  available
- the version of Kubernetes where the KEP graduated to general
  availability
- when the KEP was retired or superseded

## Drawbacks [optional]

* Use of quotas, particularly the less commonly used project quotas,
  requires additional action on the part of the administrator.  In
  particular:
   * ext4fs filesystems must be created with additional options that
     are not enabled by default:
```
mkfs.ext4 -O quota,project -Q usrquota,grpquota,prjquota _device_
```
   * An additional option (`prjquota`) must be applied in `/etc/fstab`
   * If the root filesystem is to be quota-enabled, it must be set in
     the grub options.
* Use of project quotas for this purpose will preclude future use
  within containers.

## Alternatives [optional]

I have considered two classes of alternatives:

* Alternatives based on quotas, with different implementation

* Alternatives based on loop filesystems without use of quotas

### Alternative quota-based implementation

Within the basic framework of using quotas to monitor and potentially
enforce storage utilization, there are a number of possible options:

* Utilize per-volume non-enforcing quotas to monitor storage (the
  first stage of this proposal).

  This mostly preserves the current behavior, but with more efficient
  determination of storage utilization and the possibility of building
  further on it.  The one change from current behavior is the ability
  to detect space used by deleted files.

* Utilize per-volume enforcing quotas to monitor and enforce storage
  (the second stage of this proposal).

  This allows partial enforcement of storage limits.  As local storage
  capacity isolation works at the level of the pod, and we have no
  control of user utilization of ephemeral volumes, we would have to
  give each volume a quota of the full limit.  For example, if a pod
  had a limit of 1 MB but had four ephemeral volumes mounted, it would
  be possible for storage utilization to reach (at least temporarily)
  4MB before being capped.

* Utilize per-pod enforcing user or group quotas to enforce storage
  consumption, and per-volume non-enforcing quotas for monitoring.

  This would offer the best of both worlds: a fully capped storage
  limit combined with efficient reporting.  However, it would require
  each pod to run under a distinct UID or GID.  This may prevent pods
  from using setuid or setgid or their variants, and would interfere
  with any other use of group or user quotas within Kubernetes.

* Utilize per-pod enforcing quotas to monitor and enforce storage.

  This allows for full enforcement of storage limits, at the expense
  of being able to efficiently monitor per-volume storage
  consumption.  As there have already been reports of monitoring
  causing trouble, I do not advise this option.

  A variant of this would report (1/N) storage for each covered
  volume, so with a pod with a 4MiB quota and 1MiB total consumption,
  spread across 4 ephemeral volumes, each volume would report a
  consumption of 256 KiB.  Another variant would change the API to
  report statistics for all ephemeral volumes combined.  I do not
  advise this option.

### Alternative loop filesystem-based implementation

Another way of isolating storage is to utilize filesystems of
pre-determined size, using the loop filesystem facility within Linux.
It is possible to create a file and run `mkfs(8)` on it, and then to
mount that filesystem on the desired directory.  This both limits the
storage available within that directory and enables quick retrieval of
it via `statfs(2)`.

Cleanup of such a filesystem involves unmounting it and removing the
backing file.

The backing file can be created as a sparse file, and the `discard`
option can be used to return unused space to the system, allowing for
thin provisioning.

I conducted preliminary investigations into this.  While at first it
appeared promising, it turned out to have multiple critical flaws:

* If the filesystem is mounted without the `discard` option, it can
  grow to the full size of the backing file, negating any possibility
  of thin provisioning.  If the file is created dense in the first
  place, there is never any possibility of thin provisioning without
  use of `discard`.

  If the backing file is created densely, it additionally may require
  significant time to create if the ephemeral limit is large.

* If the filesystem is mounted `nosync`, and is sparse, it is possible
  for writes to succeed and then fail later with I/O errors when
  synced to the backing storage.  This will lead to data corruption
  that cannot be detected at the time of write.

  This can easily be reproduced by e. g. creating a 64MB filesystem
  and within it creating a 128MB sparse file and building a filesystem
  on it.  When that filesystem is in turn mounted, writes to it will
  succeed, but I/O errors will be seen in the log and the file will be
  incomplete:

```
# mkdir /var/tmp/d1 /var/tmp/d2
# dd if=/dev/zero of=/var/tmp/fs1 bs=4096 count=1 seek=16383
# mkfs.ext4 /var/tmp/fs1
# mount -o nosync -t ext4 /var/tmp/fs1 /var/tmp/d1
# dd if=/dev/zero of=/var/tmp/d1/fs2 bs=4096 count=1 seek=32767
# mkfs.ext4 /var/tmp/d1/fs2
# mount -o nosync -t ext4 /var/tmp/d1/fs2 /var/tmp/d2
# dd if=/dev/zero of=/var/tmp/d2/test bs=4096 count=24576
  ...will normally succeed...
# sync
  ...fails with I/O error!...
```

* If the filesystem is mounted `sync`, all writes to it are
  immediately committed to the backing store, and the `dd` operation
  above fails as soon as it fills up `/var/tmp/d1`.  However,
  performance is drastically slowed, particularly with small writes;
  with 1K writes, I observed performance degradation in some cases
  exceeding three orders of magnitude.

  I performed a test comparing writing 64 MB to a base (partitioned)
  filesystem, to a loop filesystem without `sync`, and a loop
  filesystem with `sync`.  Total I/O was sufficient to run for at least
  5 seconds in each case.  All filesystems involved were XFS.  Loop
  filesystems were 128 MB and dense.  Times are in seconds.  The
  erratic behavior (e. g. the 65536 case) was involved was observed
  repeatedly, although the exact amount of time and which I/O sizes
  were affected varied.  The underlying device was an HP EX920 1TB
  NVMe SSD.

| I/O Size | Partition | Loop w/sync | Loop w/o sync |
| ---:     | ---:      | ---:        | ---:          |
| 1024 | 0.104 | 0.120 | 140.390 |
| 4096 | 0.045 | 0.077 | 21.850 |
| 16384 | 0.045 | 0.067 | 5.550 |
| 65536 | 0.044 | 0.061 | 20.440 |
| 262144 | 0.043 | 0.087 | 0.545 |
| 1048576 | 0.043 | 0.055 | 7.490 |
| 4194304 | 0.043 | 0.053 | 0.587 |

  The only potentially viable combination in my view would be a dense
  loop filesystem without sync, but that would render any thin
  provisioning impossible.

## Infrastructure Needed [optional]

* Decision: who is responsible for quota management of all volume
  types (and especially ephemeral volumes of all types).  At present,
  emptydir volumes are managed by the kubelet and logdirs and writable
  layers by either the kubelet or the runtime, depending upon the
  choice of runtime.  Beyond the specific proposal that the runtime
  should manage quotas for volumes it creates, there are broader
  issues that I request assistance from the SIG in addressing.

* Location of the quota code.  If the quotas for different volume
  types are to be managed by different components, each such component
  needs access to the quota code.  The code is substantial and should
  not be copied; it would more appropriately be vendored.

## References

### Bugs Opened Against Filesystem Quotas

The following is a list of known security issues referencing
filesystem quotas on Linux, and other bugs referencing filesystem
quotas in Linux since 2012.  These bugs are not necessarily in the
quota system.

#### CVE

* *CVE-2012-2133* Use-after-free vulnerability in the Linux kernel
  before 3.3.6, when huge pages are enabled, allows local users to
  cause a denial of service (system crash) or possibly gain privileges
  by interacting with a hugetlbfs filesystem, as demonstrated by a
  umount operation that triggers improper handling of quota data.
  
  The issue is actually related to huge pages, not quotas
  specifically.  The demonstration of the vulnerability resulted in
  incorrect handling of quota data.
  
* *CVE-2012-3417* The good_client function in rquotad (rquota_svc.c)
  in Linux DiskQuota (aka quota) before 3.17 invokes the hosts_ctl
  function the first time without a host name, which might allow
  remote attackers to bypass TCP Wrappers rules in hosts.deny (related
  to rpc.rquotad; remote attackers might be able to bypass TCP
  Wrappers rules).
  
  This issue is related to remote quota handling, which is not the use
  case for the proposal at hand.
  
#### Other Security Issues Without CVE

* [Linux Kernel Quota Flaw Lets Local Users Exceed Quota Limits and
  Create Large Files](https://securitytracker.com/id/1002610)
  
  A setuid root binary inheriting file descriptors from an
  unprivileged user process may write to the file without respecting
  quota limits.  If this issue is still present, it would allow a
  setuid process to exceed any enforcing limits, but does not affect
  the quota accounting (use of quotas for monitoring).
  
### Other Linux Quota-Related Bugs Since 2012

* [ext4: report delalloc reserve as non-free in statfs mangled by
  project quota](https://lore.kernel.org/patchwork/patch/884530/)
  
  This bug, fixed in Feb. 2018, properly accounts for reserved but not
  committed space in project quotas.  At this point I have not
  determined the impact of this issue.
  
* [XFS quota doesn't work after rebooting because of
  crash](https://bugs.launchpad.net/ubuntu/+source/linux/+bug/1461730)
  
  This bug resulted in XFS quotas not working after a crash or forced
  reboot.  Under this proposal, Kubernetes would fall back to du for
  monitoring should a bug of this nature manifest itself again.
  
* [quota can show incorrect filesystem
  name](https://bugzilla.redhat.com/show_bug.cgi?id=1326527)
  
  This issue, which will not be fixed, results in the quota command
  possibly printing an incorrect filesystem name when used on remote
  filesystems.  It is a display issue with the quota command, not a
  quota bug at all, and does not result in incorrect quota information
  being reported.  As this proposal does not utilize the quota command
  or rely on filesystem name, or currently use quotas on remote
  filesystems, it should not be affected by this bug.
  
In addition, the e2fsprogs have had numerous fixes over the years.
