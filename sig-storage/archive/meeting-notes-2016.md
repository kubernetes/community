# Kubernetes Storage SIG Meeting Notes (2016)

The Kubernetes Storage Special-Interest-Group (SIG) is a working group within the Kubernetes contributor community interested in storage and volume plugins. This document contains historical meeting notes from past meeting.

## December 22, 2016

Recording: [https://youtu.be/pMMAM6SNeeY](https://youtu.be/pMMAM6SNeeY)

Agenda/Notes:

* [Radoslaw Zarzynski] Provide a quick status update on the rbd-nbd integration.
  * Another way to access Ceph RBD volume
  * Original plugin only allowed access via kernel client which is less featureful than LibRBD client
  * Radoslaw created a few pull requests:
    * rbd-nbd zero-copy
      * [https://github.com/ceph/ceph/pull/12609](https://github.com/ceph/ceph/pull/12609),
      * Optimizing memory perf, comparable to krbd driver
    * [https://github.com/ceph/ceph/pull/11878](https://github.com/ceph/ceph/pull/11878)
    * rbd-nbd fencing improvement
      * [https://github.com/ceph/ceph/pull/11438](https://github.com/ceph/ceph/pull/11438)
      * Making the fencing easier, just need to set one more option
    * rbd-nbd integration with K8S
      * [https://github.com/kubernetes/kubernetes/pull/38936](https://github.com/kubernetes/kubernetes/pull/38936)
      * Support for RBD-NBD backend
    * Code reviewers:
      * [rootfs](http://github.com/rootfs)

## December 8, 2016

Recording: [https://youtu.be/1GYAulB4mdM](https://youtu.be/1GYAulB4mdM)

Agenda/Notes:

* Status Update:
  * Jing’s containerized mounter changes went into 1.5 and are backported to 1.4.7
* Storage Wish List (Brainstorming) for 2017 and 1.6 (Starter)
  * (AI) Create a separate doc to let people think it over and track thi
  * [Matt De Lio] Some ideas:
    * Only schedule pods on nodes that support a given volume type
    * Containerized Mounts/Out-of-tree Volume Driver
    * Stabilizing AWS support for EBS
    * Snapshot
      * Need a consensus on what this really need
    * Data Replication?
    * Local Storage Support
    * Mount Option Passthrough
* Quarterly face to face meeting for Q1 2017
  * Dell EMC is willing to host in our office in Santa Clara
    * Potential for this to be at KubeCon EU too? [CML]
  * Timing: Late Q1 early Q2?
  * Should have an agenda set before the meeting to make discussion more focused.
* [from last meeting] [rootfs] Metrics in volume controller
  * [https://github.com/kubernetes/kubernetes/issues/36818](https://github.com/kubernetes/kubernetes/issues/36818)
* [jsafrane] Safer mount manager that does not accidentally delete user data
  * [https://github.com/kubernetes/kubernetes/pull/37698](https://github.com/kubernetes/kubernetes/pull/37698)
* [jsafrane] Configurable ReclaimPolicy in StorageCla
  * [https://github.com/kubernetes/kubernetes/issues/38192](https://github.com/kubernetes/kubernetes/issues/38192)

## November 23, 2016

Since Thursday, November 24, 2016 is a holiday in the United States, we are moving this occurrence of the Storage SIG meeting one day early to Wednesday, November 23, 2016.

Recording: [https://youtu.be/fhtaR0rsLBA](https://youtu.be/fhtaR0rsLBA)

Agenda/Notes:

* [Steve Watt] In-Tree/Out-of-Tree question
  * [Michael Rubin] Plugin Freedom
  * [Scott Creeley] Enhanced error resolution hint
  * Flex v.next
    * Let’s separate future of Out-of-tree and Flex v.next
* [Saad Ali] Hung volumes can wedge the kubelet
  * [https://github.com/kubernetes/kubernetes/issues/31272](https://github.com/kubernetes/kubernetes/issues/31272)
* [rootfs] Mount option
  * [https://github.com/kubernetes/kubernetes/issues/31613](https://github.com/kubernetes/kubernetes/issues/31613)
* Did not get to: [rootfs] Metrics in volume controller
  * <https://github.com/kubernetes/kubernetes/issues/36818>

## November 10, 2016

Recording: [https://youtu.be/k1JRpMd051M](https://youtu.be/k1JRpMd051M)

Agenda/Note

* Feature complete date
  * (saad-ali) Now in code freeze for 1.5
* Status updates on pending work
  * (jingxu97) NFS Gluster Containerization
  * (rkouj) Better Messaging for missing Binarie
* Work that missed 1.5
  * (rootfs) RBD attached
  * (jsafrane) Default provisioner
* [Steve Wong] KubeCon Storage SIG F2F Recap:
  * Agenda: [https://docs.google.com/document/d/1drdxPkZEiGA06-jnsbSqywzQ6z0bT8uMpJdnIzkRoPo/edit?ts=57ed5df2#heading=h.hs60nbsyz2jk](https://docs.google.com/document/d/1drdxPkZEiGA06-jnsbSqywzQ6z0bT8uMpJdnIzkRoPo/edit?ts=57ed5df2#heading=h.hs60nbsyz2jk)
  * Meeting minutes/notes: [https://docs.google.com/document/d/10AJzJvJcEA5w1wW6wMgS7aOfiH0FTZhszt_P6VNRxLk/edit?usp=sharing](https://docs.google.com/document/d/10AJzJvJcEA5w1wW6wMgS7aOfiH0FTZhszt_P6VNRxLk/edit?usp=sharing)

## November 7, 2016

F2F meeting in Seattle, WA

* Logistics/Agenda: [https://docs.google.com/document/d/1drdxPkZEiGA06-jnsbSqywzQ6z0bT8uMpJdnIzkRoPo/edit?ts=57ed5df2](https://docs.google.com/document/d/1drdxPkZEiGA06-jnsbSqywzQ6z0bT8uMpJdnIzkRoPo/edit?ts=57ed5df2)
* [Minutes:](https://docs.google.com/document/d/1drdxPkZEiGA06-jnsbSqywzQ6z0bT8uMpJdnIzkRoPo/edit?ts=57ed5df2)[https://docs.google.com/document/d/10AJzJvJcEA5w1wW6wMgS7aOfiH0FTZhszt_P6VNRxLk/edit](https://docs.google.com/document/d/10AJzJvJcEA5w1wW6wMgS7aOfiH0FTZhszt_P6VNRxLk/edit)

## October 27, 2016

Recording: [https://youtu.be/gjb2X1QiKG0](https://youtu.be/gjb2X1QiKG0)

Agenda/Note

* Reminder meeting is now held on Zoom
  * This meeting will be held on Zoom (<https://zoom.us/j/614261834>) NOT Hangouts!
  * Meetings will be recorded and published.
* [Saad Ali] Update on NFS/Gluster in GCI (GCE/GKE)
* [Michael Rubin] On Containerization and the future of volume plugins and deployment (Flex, etc.)
* [Hayley Swimelar (Linbit)] Adding support for DRBD Volumes. issue #[32739](https://github.com/kubernetes/kubernetes/pull/32739)
* [Saad Ali] We have less than two weeks left until code freeze. Please list the outstanding PRs that really need attention:
  * PR #[30285](https://github.com/kubernetes/kubernetes/pull/30285) - Proposal for external dynamic provisioner
  * PR #[30091](https://github.com/kubernetes/kubernetes/pull/30091) - support Azure disk dynamic provisioning
  * PR #[33660](https://github.com/kubernetes/kubernetes/pull/33660) - rbd attach/detach refactoring
  * PR #[35284](https://github.com/kubernetes/kubernetes/pull/35284) - Add test for provisioning with storage cla
  * PR #[35675](https://github.com/kubernetes/kubernetes/pull/35675) - Require PV provisioner secrets to match type
* [Erin Boyd] Status update on Testing
  * [bchilds] - update on AWS E2E test
* [bchilds] External NFS Provisioner incubator proposal
* [Simon Croome] StorageOS introduction / plans for Kubernetes integration
* [bchilds] External Storage Provisioner - [https://github.com/kubernetes/kubernetes/pull/30285](https://github.com/kubernetes/kubernetes/pull/30285)
* [Steve Wong] Storage SIG face to face Nov 7 logistic
  * Agenda: [https://docs.google.com/document/d/1drdxPkZEiGA06-jnsbSqywzQ6z0bT8uMpJdnIzkRoPo/edit?ts=57ed5df2#](https://docs.google.com/document/d/1drdxPkZEiGA06-jnsbSqywzQ6z0bT8uMpJdnIzkRoPo/edit?ts=57ed5df2#)
  * 505 1st Ave S. Ste 600, Seattle, WA 98104 - Elliott Bay Conference Room - Floor 2Doors open 7:30, get badge from receptionistPlease RSVP using entry in meeting doc by Oct 31.

## October 13, 2016

Agenda/Note

* Switching to Zoom
  * This meeting will be held on Zoom ([https://zoom.us/j/614261834](https://zoom.us/j/614261834)) NOT Hangouts!
  * Future meetings will be recorded and published.
* Updates on librbd and rbd-nbd integration with Kubernetes.
  * Radoslaw Zarzynski - Looking at TCMU in addition to rbd-nbd.  Interested in building a bridge between.  Plugin needs to address both and doing so is possible since they are so similar.  Taking requirements on CEPH side, have people from mirantis to do the development.  Looking at in-tree instead of flex and plans to start development next week. Could support V2 vs V1 [https://github.com/kubernetes/kubernetes/issues/32266](https://github.com/kubernetes/kubernetes/issues/32266)
  * Love to see PoC and implementations coming. Please ensure krbd is not broken when introducing user space rbd.
* Reminder: Please review the external storage provisioning PR [https://github.com/kubernetes/kubernetes/pull/30285](https://github.com/kubernetes/kubernetes/pull/30285)
* Flex in 1.5 - Will not make 1.5, but will discuss in November F2F.  Working with Docker to come to consensus about the API such that its the same between Flex and Docker.  Should see updates on proposal next week.  Trending towards socket model instead of exec.  Proposals will be updated to reflect it.
* November F2F Agenda: [https://docs.google.com/document/d/1drdxPkZEiGA06-jnsbSqywzQ6z0bT8uMpJdnIzkRoPo/edit?ts=57ed5df2#](https://docs.google.com/document/d/1drdxPkZEiGA06-jnsbSqywzQ6z0bT8uMpJdnIzkRoPo/edit?ts=57ed5df2#)
* provisioning/deleting secrets on storage class
  * Passing via annotations on a PV from provisioning to deleting is not secure
  * Issue: [https://github.com/kubernetes/kubernetes/issues/34822](https://github.com/kubernetes/kubernetes/issues/34822)
* AWS Storage E2E Tests Update - Need an XFS filesystem on the host images.
* DAS - (local) storage - may be designed in F2F. no target release.

## September 29, 2016

Agenda/Note

* Status Update
  * v1.4 Work Completed
    * Dynamic Provisioning
      * <http://blog.kubernetes.io/2016/10/dynamic-provisioning-and-storage-in-kubernetes.html>
    * Bug fixes/stabilization
  * v1.5 Coding begin
    * Increasing test coverage
    * Critical bug fixe
      * E.g. <https://github.com/kubernetes/kubernetes/issues/29324>
    * Dynamic Volume Provisioning
      * Default provisioner
        * Need to get this in for v1.5. We want pre-installed provisioners, only open question is if they should be marked default or not.
      * External Provisioner
        * Design Proposal: [https://github.com/kubernetes/kubernetes/pull/30285](https://github.com/kubernetes/kubernetes/pull/30285)
        * NFS Provisioner: [https://github.com/wongma7/nfs-provisioner#deployment](https://github.com/wongma7/nfs-provisioner#deployment)
        * Proposal, proof of concept NFS provisioner. Can demo it, and use it to ratify proposal and start implementing out-of-tree.
        * When do we want it to land?
        * Probably a peer to kubernetes repo.
    * GCI Storage Utilitie
      * NFS Testing
        * Humain? Can you help us out with this. [https://github.com/kubernetes/kubernetes/issues/33447#issuecomment-250396321](https://github.com/kubernetes/kubernetes/issues/33447#issuecomment-250396321)
    * Flex Volumes <https://github.com/kubernetes/kubernetes/issues/32543>
      * Maybe we should ratify proposal at F2F for v1.5 and then focus on implementation for v1.6
    * Snapshot
      * Design only
      * Jing Xu working on this, please reach out to her.
      * AI (Jing Xu): beginning to middle of November send out a 2nd iteration of design
    * Mount option support (move to F2F), [https://github.com/kubernetes/kubernetes/issues/31613](https://github.com/kubernetes/kubernetes/issues/31613)
      * Work on proposal, discuss at F2F
    * LibStorage PR needs to be looked at
* Adding support for Ceph rbd-nbd client[https://github.com/kubernetes/kubernetes/issues/32266](https://github.com/kubernetes/kubernetes/issues/32266)
  * Proposal for new driver for RBD, RBD-NBD,
  * At the moment, ceph rbd, very unusual compare to vmware lib-rbd, as a result big feature gap compared to openstack
  * Using kernel driver could affect the reliability of nodes.
  * Maybe taking another approach?
  * One feature we get out of the box is lack of feature gap, like object map, etc, these are being implemented in lib rbd.
  * Move conversation to issue.
* Update on: Chakri has a PR with the first cut of the API and if anyone has comments please add them to the PR: [https://github.com/kubernetes/kubernetes/issues/32543](https://github.com/kubernetes/kubernetes/issues/32543)
* Switching to Zoom

## September 15, 2016

Agenda/Note

* v1.4 Status
  * AWS Attaches issue: Wrong AWS volume can be mounted
    * <https://github.com/kubernetes/kubernetes/pull/31090>
    * <https://github.com/kubernetes/kubernetes/pull/32242>
    * <https://github.com/kubernetes/kubernetes/pull/32636>
  * NFS Volume Hanging
    * Delayed to next release
  * 1.4.1 should be open next week
* (Huamin) Attach/detach interface feature request: [https://drive.google.com/file/d/0B0YiMoOYtt_sYld6ZUE1YkMxem8/view?usp=sharing](https://drive.google.com/file/d/0B0YiMoOYtt_sYld6ZUE1YkMxem8/view?usp=sharing)
  * Saad: Volume Spec should be fine
  * Saad: Device path should be picked up from WaitForAttach and used down the pipeline
* Brief overview of nfs provisioner work in progress - bchilds / Matthew Wong
  * Questions around capacity on NFS provisioner
* (Jan) Passwords in StorageCla
  * [https://github.com/kubernetes/kubernetes/pull/31869](https://github.com/kubernetes/kubernetes/pull/31869)
  * Did not progress since last SIG meeting
  * AI: Put it on the next meeting Agenda, ping Tim to join
* Hangout hit limit more people can no longer join call
  * AI: Saad will look in to Zoom/see if we can increase hangout limit.
* Vlad: Progress on LibStorage
  * Attach and dynamic provisioners: implemented
* Gopal: for external APIs flex volume. Chakri has a PR with the first cut of the API and if anyone has comments please add them to the PR: [https://github.com/kubernetes/kubernetes/issues/32543](https://github.com/kubernetes/kubernetes/issues/32543)
* Vlad: Unable to modify agenda
  * Use [this link](https://docs.google.com/document/d/1-8KEG8AjAgKznS9NFm3qWqkGyCHmvU6HVl0sk5hwoAE/edit?usp=sharing), it should give you ability to suggest change

## September 1, 2016

Agenda/Note

* v1.4 Statu
  * API object move from extension
  * Bug
    * NFS mounts can wedge
      * Pmorie working on it
      * [https://github.com/kubernetes/kubernetes/issues/31272](https://github.com/kubernetes/kubernetes/issues/31272)
    * Node “attach/detach” control, ensure field is always updated.
    * AWS Attaches issue: Wrong AWS volume can be mounted
      * [https://github.com/kubernetes/kubernetes/issues/29324](https://github.com/kubernetes/kubernetes/issues/29324)
    * Some flaky test
* v1.5 Priorities
  * Testing
  * Debuggability
  * Code hardening
  * Dynamic Provisioning:
    * Direction for “secrets” and endpoints on PV
      * Jan: we need to close on this.
      * Concern about backwards compat
      * Tim: 1.4 too late, hard to justify.
      * MRubin/PMorie: might be something we want to bring up at community SIG--other SIGs are touching on and interested in this area
      * Erin, Eric Paris, and Jan will help with this.
    * External provisioner
      * Tim may be able to look at it for 1.5 (about 3 weeks)
  * Flex Volume
    * Resolve open issues.
  * Local Storage
    * Beep and Clayton are driving, SIG should keep an eye on it
  * Snapshotting
    * Jing is driving
  * There is a kubernetes features repo, where issues are created and tracked (in v1.4 we did this last minute), for 1.5 we should do this early.
    * [https://github.com/kubernetes/features](https://github.com/kubernetes/features)
* Switch to from Hangouts to something else
  * Zoom?
    * Others use thi
    * Requires an external binary
    * Seems to work fine
    * You need a Mac or Linux, no ChromeBook
    * Paid service--need an Admin account for SIG
  * Current limit is 30 on Hangouts.
  * Saad: We’ll keep hangouts until we hit the limit. If we start to have issues, I’ll look into switching us to Zoom
* Discuss dates for next F2F
  * Nov 7, 2016 before KubeCon (Nov 8 to Nov 9) work for everyone?
    * Steve Wong has tentatively arranged for a conference room for 25 people at the [EMC Seattle office](https://www.google.nl/maps/place/EMC+Corporation/@47.6009484,-122.3377655,16z/data=!4m8!1m2!2m1!1semc+seattle!3m4!1s0x54906aa489f54c35:0xc91e3386db767a1c!8m2!3d47.5977856!4d-122.3345112?hl=en) near Pioneer Square. This includes lunch and hosting of a dinner at some nearby location TBD
      * Patrick from Samsung can help with organizing.

## August 18, 2016

Agenda/Note

* Update from F2F for Community
  * [https://groups.google.com/forum/#!topic/kubernetes-sig-storage/tsKYK6rdjy0](https://groups.google.com/forum/#!topic/kubernetes-sig-storage/tsKYK6rdjy0)
  * Steve Watt volunteered to follow up
* v1.4 Code Freeze this week
  * Status of in-flight PR
    * Flex Volume
      * In review
      * Very large but well factored
      * Should be able to get it in before code freeze
      * Some open questions around attacher
    * Azure
      * Verifying functionality
      * Questions about performance
    * Quobyte
      * In review, no major issues.
    * Event
      * LGTM’d
* Storage community test strawman - bchilds [https://docs.google.com/document/d/17j5ofzOOhWUVBOJ3Uop-MUay2k1iJomBJuVCGxMZcMA/edit?usp=sharing](https://docs.google.com/document/d/17j5ofzOOhWUVBOJ3Uop-MUay2k1iJomBJuVCGxMZcMA/edit?usp=sharing)
  * Can’t test all plugin
    * Non-GCE cloud storage plugins can’t run in GCE need federated testing
* Jan: StorageClass/PV security - jsafrane
  * Is it safe to store credentials to Gluster server on non-namespaced objects?
    * 1.4 storage classes with .
    * GlusterFS needs to store user/name password etc in blob.
    * Problem with secrets is we don’t have cross namespace/using a secret from non-namespaced context. Should get BrianGrant’s opinion.
    * Parameters are a map of strings.
    * Need to handle this for more than Gluster. File on FS.
    * Will go forward with password in blob, and
* Mark: 1.5 features question
  * Snapshots: would like to participate in design
    * Sync up with Jing

## August 11, 2016

F2F meeting in San Jose

* Intro Slide
  * [https://docs.google.com/presentation/d/1wLUYKABzj1gUrn-JUlSo0QXKeKValy9EBgrwHuyJC6E/edit](https://docs.google.com/presentation/d/1wLUYKABzj1gUrn-JUlSo0QXKeKValy9EBgrwHuyJC6E/edit)
* Agenda
  * [https://docs.google.com/document/d/1qVL7UE7TtZ_D3P4F7BeRK4mDOvYskUjlULXmRJ4z-oE/edit](https://docs.google.com/document/d/1qVL7UE7TtZ_D3P4F7BeRK4mDOvYskUjlULXmRJ4z-oE/edit)
* Note
  * [https://docs.google.com/document/d/1vA5ul3Wy4GD98x3GZfRYEElfV4OE8dBblSK4rnmrE_M/edit?ts=57ab5b83](https://docs.google.com/document/d/1vA5ul3Wy4GD98x3GZfRYEElfV4OE8dBblSK4rnmrE_M/edit?ts=57ab5b83)

## August 4, 2016

Agenda/Note

* Status Updates:
  * v1.3.4 went out on Monday
    * Lots of storage fixe
* Face to face meeting - Everyone welcome - Logistics info -&gt; [https://docs.google.com/document/d/1qVL7UE7TtZ_D3P4F7BeRK4mDOvYskUjlULXmRJ4z-oE/edit](https://docs.google.com/document/d/1qVL7UE7TtZ_D3P4F7BeRK4mDOvYskUjlULXmRJ4z-oE/edit)
* Tim: Internal/external plugins discussion
  * Probably want to save it for the face-to-face meeting. It’s big.
* Tim: Storage volumes read-write once multiple pods can use it at once. This is inconsistent, if pod happens to land on the same machine it will work if not it will not.
  * Two separate conversations if that is correct or not.
  * Share if on same machine AND bias scheduler to put on same machine.
  * Use case: people want to do rolling updates without down time
  * Concern: is this something that is unsafe. We can’t guarantee it will always work for applications.
  * AI: would love to define the correct semantic document it
  * Eric Paris: this is a bad idea.
  * Eric Paris: biasing scheduler could be a good idea. but two mounting at same time bad idea.
  * AI: make attacher/detacher smarter--don’t detach if volume is already attached
    * Saad: may already do that.
  * Tim: What about a new Read-Write-Node access mode?
  * Steve Watt: maybe k8s should fail Read-Write-Only if it is used by multiple pod
  * Agreement, using Read-Write-Only in multiple pods is an anti-pattern
  * AI: Tim will follow up on the various conversations.
  * AI: Saad will check if Attach/Detach is already optimal for waiting pod. If not file low-pri bug.

## July 21, 2016

Agenda/Note

* Diamanti Demo
  * Action item: would like snapshotting support in flex
* Discuss support for Cinder volumes without cloud-provider (cf hypernetes)
  * Enable Cinder volumes to attach to any VM or bare metal
  * Action Item: WIP will be sent out by Quentin
* Dynamic Provisioning V2 [https://github.com/kubernetes/kubernetes/pull/29006](https://github.com/kubernetes/kubernetes/pull/29006)
  * Action item: need reviewers for PR
* Face to face meeting
  * Decided on dates: Aug 10 and 11 with optional break out sessions on Aug 12
* Post-mortems:
  * Google has a format.
  * Brad work on it with Michael?
* Lots of storage improvements in v1.3.4
  * If anyone has spare cycles to help debug issues, please volunteer
* Flex volume/pluggable volume discussion
  * Blocked on Tim Hockin
  * Would be great to discuss at F2F
* Portworks welcome!

## July 7, 2016

Agenda/Notes:

* EMC{code} team proposal on external persistent volume mounts for container
  * 10 minute presentation followed by Q &amp; A by [steven.wong@emc.com](mailto:steven.wong@emc.com) and [vladimir.vivien@emc.com](mailto:vladimir.vivien@emc.com)
  * [https://github.com/kubernetes/kubernetes/pull/28599](https://github.com/kubernetes/kubernetes/pull/28599)
* Bug Scrub (bchilds@redhat.com)
  * Propose that we target issues the following way:
    * Milestone v1.3 - All issues that are targeted for some 1.3.x release
    * Milestone v1.4 - All issues for v1.4
    * “Unassign” All subsequent than v1.4.
* Post Mortem ([mrubin@google.com](mailto:mrubin@google.com), [eboyd@redhat.com](mailto:eboyd@redhat.com), [bchilds@redhat.com](mailto:bchilds@redhat.com), [saadali@google.com](mailto:saadali@google.com))
  * We need to do a post mortem on the storage events at the tail of v1.3.
  * We should figure out how to get that started.
* Making our testing and code paths more robust (eboyd@redhat.com)
  * What makes sense to make the Kubernetes test suite more robust and
  * easier to find bugs over all?
* Face to Face SIG Meeting ([swatt@redhat.com](mailto:stevewatt@redhat.com), gopal@datawisesystems.com, mrubin, nelcyguy@gmail.com)
  * When is a good time for a storage sig face to face?
* v1.4 feature
  * Dynamic Provisioning/Storage Classes V2 Proposal (thockin@google.com,Tom? [nelcyguy@gmail.com](mailto:nelcyguy@gmail.com), [pmorie@redhat.com](mailto:pmorie@redhat.com), tg@convergeio.com)
    * (Jan Šafránek &amp; Paul Morie)  [https://github.com/kubernetes/kubernetes/pull/26908](https://github.com/kubernetes/kubernetes/pull/26908)
    * Want proposal submitted by eod 7/8
  * GID feature for PV
    * <https://github.com/kubernetes/kubernetes/issues/27197>
    * Want early in 1.4
  * Flex Volume
    * New Volume plugin
    * Discussion on Flex Volumes or built in storage (See
    * <https://groups.google.com/forum/#!topic/kubernetes-sig-storage/9o1vA4jFwqk>)
  * Local storage
    * Start on the design, maybe have some alpha level code
    * Action item: open a bug for people to start contributing to
* Run down testing (eparis@redhat.com) [https://github.com/kubernetes/kubernetes/issues/28367](https://github.com/kubernetes/kubernetes/issues/28367)

## June 23, 2016

Agenda/Notes:

* v1.3 status update
  * Kubelet Mount/Unmount Redesign [merged](https://github.com/kubernetes/kubernetes/pull/26801)
  * Bugs Squashed
  * Open, potentially 1.3 blocking bugs:
    * [Flakiness of read-only PD test](https://github.com/kubernetes/kubernetes/issues/27477)/[#27691](https://github.com/kubernetes/kubernetes/issues/27691)
  * Flex Volume Plugin changes punted from 1.3
    * [https://github.com/kubernetes/kubernetes/pull/26926](https://github.com/kubernetes/kubernetes/pull/26926)
  * Testing
    * Erin:
      * NFS Gluster Cep
        * Should be done by today
    * Humin:
      * Cinder
        * Should be done by today
    * Justin SB
      * AWS
        * Looks good
* v1.4 priorities
  * Want early in 1.4 (1.3.x release):
    * 1) Dynamic Provisioning/Storage Classes V2 Proposal (Jan Šafránek)
      * [https://github.com/kubernetes/kubernetes/pull/26908](https://github.com/kubernetes/kubernetes/pull/26908)
      * What’s missing to get merged?
    * 2) GID feature for PV
      * Shared File/Read Write Many
      * [https://github.com/kubernetes/kubernetes/issues/27197](https://github.com/kubernetes/kubernetes/issues/27197)
    * 3) Make volume manager more robust across restart
      * [https://github.com/kubernetes/kubernetes/issues/27653](https://github.com/kubernetes/kubernetes/issues/27653)
    * 4) Finish of changes to Flex Volume Plugin
      * [https://github.com/kubernetes/kubernetes/pull/26926](https://github.com/kubernetes/kubernetes/pull/26926)
      * Must have for 1.3.1 release
    * 5) Volume Operation Executor Should Generate Event
      * [https://github.com/kubernetes/kubernetes/issues/27590](https://github.com/kubernetes/kubernetes/issues/27590)
  * Other priorities
    * 6) Improve Testing
      * Huge item
      * Performance testing
      * Testing core component
      * Plugin
    * 7) Continue to improve stability and robustness of existing code
    * 8) New Volume plugin
      * Encourage people who submit volume plugins to stick around and make sure that they continue to work.
      * Core team does not have the ability to test them.
    * 9) Local storage
      * Start on the design, maybe have some alpha level code
  * Saad’s feature list:
    * ---Wanted Feature/requests for volumes---
    * Improve Test Coverage
    * Continue to improve stability and robustness of existing code
      * Refactor some existing code
    * New Feature
      * Finish Dynamic Provisioning
        * volume selectors (done)
        * volume classe
      * Container volume for pre-populated emptydir container #831 (probably read-only)
        * Empty dir with a claim?
      * LogDir
      * Data gravity/local storage
        * When scheduling a work loads affinity to certain nodes that provide schedule.
        * For ephemeral volumes (empty dir): use local storage instead of empty dir
        * AI (saadali): Open a PR for requirements so others can contribute
          * [https://github.com/kubernetes/kubernetes/issues/7562](https://github.com/kubernetes/kubernetes/issues/7562)
      * Enable mount namespace propagation to enable storage providers to run on k8s and to containerize kubelet [https://github.com/kubernetes/kubernetes/pull/20698](https://github.com/kubernetes/kubernetes/pull/20698)
      * Snapshotting support in volume API
      * Automatic resize increase/decrease of Attached/mounted disks [https://groups.google.com/forum/#!topic/kubernetes-sig-storage/YLWEbTDKTHE](https://groups.google.com/forum/#!topic/kubernetes-sig-storage/YLWEbTDKTHE)
        * Steve Watt: Scaling up and snapshotting might be something we can not ignore. We are lagging behind the competition in comparisons.
      * FUSE FS support: would enable mounting Google Cloud Storage into a container #7890 /FUSE FS [http://stackoverflow.com/questions/35966832/mount-google-storage-bucket-in-google-container?noredirect=1#comment59614874_35966832](http://stackoverflow.com/questions/35966832/mount-google-storage-bucket-in-google-container?noredirect=1#comment59614874_35966832)
      * Improve Deployment Story
        * User experience vs extensibility
        * Establish a concrete API and move plugins out of tree?
        * Containerize plugin binaries so no dependency on deployment of binaries to node.
        * Containerize volume plugins to enable dynamic loading of plugins on demand?
        * Magic: run a kubectl command, and X storage system is deployed with all appropriate binaries and APIs?
    * New Plugin
      * Nexenta
      * EMC
      * Azure VHD
      * Others pending?
    * Bug level work:
      * Rapid delete/recreation should schedule to same node and not detach (bgrant says p2 low user impact)
      * GCE PD attach/detach sometimes takes several minutes to complete--follow up with GCE team to figure out if this is normal/expected.
      * UX issues brought up by Erin (&quot;at least 10GB&quot; is confusing) (bgrant recommends adding limit to api)
      * Support dynamic max pod per node limit (scheduler)
    * Misc
      * Quotas we need a design (esp for creating 1000s of new volumes) (limit claims PR by derekcarr?)
      * Explore when does dynamic volume deletion happen?
      * consider using taints/toleration for resource reservation

## June 9, 2016

Agenda/Notes:

* Dynamic Provisioning V2 Proposal (Paul Morie)
  * [https://github.com/kubernetes/kubernetes/pull/26908](https://github.com/kubernetes/kubernetes/pull/26908)
  * In-tree vs dynamic loading
    * Baby steps. Start with what we have today.
    * Want to eventually have ability to dynamically load plugin
* Attach/Detach controller
  * Status update
* Splitting up attach/detach for plugin
  * Status update
  * Issue
    * Friday branch deadline
    * Attach/detach split for rbd and flex requires namespace
      * [https://github.com/kubernetes/kubernetes/pull/26926#issuecomment-224739520](https://github.com/kubernetes/kubernetes/pull/26926#issuecomment-224739520)
* Kubelet Mount/Unmount redesign
  * Status update
* Misc:
  * hyperkube mounts issue (jing xu investigated)
    * [https://github.com/kubernetes/kubernetes/pull/27054](https://github.com/kubernetes/kubernetes/pull/27054)
* Add option to disable dynamic provisioning
  * [https://github.com/kubernetes/kubernetes/pull/27128/files](https://github.com/kubernetes/kubernetes/pull/27128/files)
* Introduce Michael Rubin
  * [mrubin@google.com](mailto:mrubin@google.com)
  * Github: [https://github.com/matchstick](https://github.com/matchstick)

## May 26, 2016

Agenda/Note

* Status updates on pending work
  * Persistent Volume Controller Consolidation (jan)
    * Persistent volume controller sync period [https://github.com/kubernetes/kubernetes/issues/24236](https://github.com/kubernetes/kubernetes/issues/24236)
  * Status of volume plugins (hchen)
  * Status of Attach/Detach controller and mount/unmount redesign (saad-ali)

## May 19, 2016

Agenda/Note

* Policy of adding new volume plugins that don’t have e2e yet (for example, Quobyte <https://github.com/kubernetes/kubernetes/pull/24977>) (hchen)
* Pv-selector PR [https://github.com/kubernetes/kubernetes/pull/25413](https://github.com/kubernetes/kubernetes/pull/25413) (pmorie)
* Attach / Detach work (saad / jan)

## May 12, 2016

Agenda/Notes:

* The reviewed controller binder is ready for merge (Today?)
* Storage test coverage - (swatt &amp; hchen)
  * [Volumes e2e test RFC](https://github.com/kubernetes/kubernetes/issues/25120) and [follow-up PR](https://github.com/kubernetes/kubernetes/pull/25100)
  * [Persistent Volume e2e and integration test RFC](https://github.com/kubernetes/kubernetes/issues/25120) and [follow-up PR](https://github.com/kubernetes/kubernetes/pull/25216)
* Storage class proposal (Pmorie)
* Dynamic Provisioning proposal (Pmorie)
* Storage Sig next Thursday to finalize storage class and DPv2
* Status update on in flight PRs (saad, sami, jan)

## April 28, 2016

Agenda/Notes:

* Status update on:
  * Binder/recycler/provisioner controller
    * Jan is working on it PR in progress.[https://github.com/kubernetes/kubernetes/pull/24331](https://github.com/kubernetes/kubernetes/pull/24331)
  * Attach/detach controller
    * [https://github.com/kubernetes/kubernetes/pull/24838](https://github.com/kubernetes/kubernetes/pull/24838)
    * [https://github.com/kubernetes/kubernetes/pull/24696](https://github.com/kubernetes/kubernetes/pull/24696)
  * Mount/unmount redesign
    * [https://github.com/kubernetes/kubernetes/pull/24557](https://github.com/kubernetes/kubernetes/pull/24557)
  * E2E and Test Coverage - hchen
    * [https://docs.google.com/document/d/1UYI0Zm2Vv1-XuSMisl_b5D9AOnG4kR7Dm4VHa8YmRmw/edit#heading=h.cnatx4syp4km](https://docs.google.com/document/d/1UYI0Zm2Vv1-XuSMisl_b5D9AOnG4kR7Dm4VHa8YmRmw/edit#heading=h.cnatx4syp4km)
  * Brad Childs: What can we do to get volume plugin test coverage to 80%
    * For v1.3 write the test
    * For v1.4 worry about deployment/infrastructure for CI
  * Hchen: azure block storage--ok to implement before Azure cloud provider support?
    * If it works, yes. Just make sure to prioritize it correctly against other items.

## April 14, 2016

Agenda/Notes:

* Status update on:
  * Binder/recycler/provisioner controller
    * Jan is working on it PR in progress.
  * Attach/detach controller
    * Some modifications to design
  * mount/unmount redesign
    * In progress PR
* Volume test plan
  * RH working  on [https://docs.google.com/document/d/1UYI0Zm2Vv1-XuSMisl_b5D9AOnG4kR7Dm4VHa8YmRmw/edit#heading=h.5dimicy6x2yl](https://docs.google.com/document/d/1UYI0Zm2Vv1-XuSMisl_b5D9AOnG4kR7Dm4VHa8YmRmw/edit#heading=h.5dimicy6x2yl)

## March 31, 2016

Agenda/Notes:

* Discuss binder/recycler/provisioner controller consolidation
  * Psudo code: [https://github.com/pmorie/pv-haxxz](https://github.com/pmorie/pv-haxxz)
  * Action item:
    * Will publish link to pseudo-code in sig-storage email list for public comment.
* Steve Watt: Current dynamic provisioning proposal--two concepts and dynamic provisioning
  * Propose: Break into two separate design
  * Storage classes with manually provisioned ideas may be able to be split off
  * But the two features should be designed together.
  * How bad would it be to only have classes and not provisioning?
    * RH: will circle back.
  * Classes and profiles must be implemented before “next version” of dynamic provisioning.
  * Decided we’re ok with shipping 1.3 with either, both, or neither features.
* Let’s get more participation from non-RH/Google folks:
  * Steve Wong from EMC:
    * silent because things appear to be going just fine
  * Chakri from Datawise:
    * Would like to give a demo in 4 weeks for flex volume.
  * Tom from ConvergeIO
    * Notion of profiles that incorporate a lot of the auto-provisioning, dir mounting, etc, in the profile
    * Standard profile for work profiles.
    * Documenting REST API, should be in a position to share with the rest of the storage-SIG

## March 17, 2016

Agenda/Notes:

* Proposal to Create new way to match PVC PV “strict” - Erin Boyd
  * Two Issues:
    * Need a way to control pod defined storage
      * Suggestions: Pod sec policy, admission controller
    * Need a way to specify exactly which PV a PVC should bind to
      * Proposal: use [Taints/Tolerations](https://github.com/kubernetes/kubernetes/blob/master/docs/design/taint-toleration-dedicated.md)
  * Action Item: Erin/Brad will create a new page under storage-sig on k8s wiki to track “User Stories”
* Containerized mount Design Proposal - Huamin Chen
  * Proposal: ConfigMap + DeamonSet
  * Issues:
    * Docker dependency
    * Deamon is single of of failure for FUSE FS
    * Single container vs multiple container
    * Mounter pods should be subject to CPU/mem quota limit
    * Consider deamonset vs just scheduling a single pod
      * Only FUSE requires a long running proce
    * How does mount namespace propagation get triggered for k8s container
    * Auth mechanisms (so arbitrary user containers can’t trigger mount/unmount
* Bind/Provision/Recycler Controller Consolidation PRs - Jan Safranek
  * Existing PR needs review
  * Suggest diagramming the existing flow and coming up with a detailed design proposal
* attach/detach controller and mount/unmount redesign status - Sami Wagiaalla

## March 1, 2016

Agenda:

* Containerized Volume Driver
  * [https://github.com/kubernetes/kubernetes/pull/](https://github.com/kubernetes/kubernetes/pull/22216)[22216](https://github.com/kubernetes/kubernetes/pull/22216)
* EMEA SME concerns around PVs(Erin / Scott)
  * Customer concern
    * Data replication
      * Moving Data with PVs that are dynamically provisioned
    * Local storage
      * Data locality issues that shared storage bring
    * selector label by security
    * Non-enforcement of label
    * Manageability of storage when there is a 1:1 for PV-&gt;PVC (too many)

* Dynamic Provisioning - Level of backwards Compat?  Maintain old way (marked as Experimental), do we need to keep old behavior?
* Jan Safranek - “Any chance to move SIG to another day?”. Pmorie - “This SIG overlaps with Node-SIG

## February 16, 2016

Agenda:

* Nearly all Dynamic Provisioning talk
  * Discussed through the various means of parameterizing options to help the “many configMap” issue.  Decided none of the approaches here are blocked by the current proposal.
  * Discussed using “storage-class” as a special field in claim.Spec.PVSelector.  Decided that allowing arbitrary labels on PVs and Provisioners was better than magic keys.
* Brought up API PRs that are required for provisioning
* Brought up status of attach/detach controller.  Work is ongoing.

## February 9, 2016

Agenda:

* Dynamic provisioner
* 1.2 PR

## February 2, 2016

Agenda:

* Dynamic provisioner PR selector vs resources (eparis)
* PRs for 1.2 (thockin)
* Azure Plugin (rootfs)
* Refactoring catchup / attach-detach (saad-ali)
* Fake Multi-write - yes or no.
* XFS quota &amp; emptydir

Dynamic provisioning:

* eparis: selector vs resources, #17056 usability
  * What goes into a selector?  Why are resources not selector?
  * Discovering the meaning of a selector is opaque
  * thockin: could be fixed in kubectl tooling “show me storage classes”
  * Steve: admin wants to catalog storage without users really understanding
  * eparis: as a new user what class do I use?
  * Sami: You know to use class as the key.  Could maybe upgrade to a full-field.
  * Erin: How does the user know what to request?
  * thockin: need a default “” class, but admin has to document what classes are available
  * cargo culting will happen
  * thockin: could promote class to a field and configmap to an object, eventually
  * eparis: can I look at a config map and understand the meaning
  * saad: configmap holds params for provisioners, up to admin
  * could expose more knobs to users (iops, etc).  Gets crazy quickly.
  * Steve: can we get tooling of “show me what classes \*I\* can use
  * thockin: no concept of “I”.  Punting on context-dependent classes for now
  * could add a description string to config map
  * eparis: config map blowout?
  * don’t need combinatoric maps for zones/feature
  * class says which provisioner, regions, encryption, etc could be features.
  * not great for discoverability.
    * could document at the configmap
    * can we validate this?  UX is not great
    * events are easy, api validation is much harder - async
  * provisioners MUST satisfy all fields of the selector
  * e.g. class=gold-east/gold-west vs class=gold,region=east
  * if zone is a parameter label, does that affect matching against PV
  * mapping to PV is looser?  examples: object name == class?  label zone is noise.
  * what if user does not ask for a zone, but provisioner needs one?  provisioner has to pick a default or choose to fail

1.2 PRs:

* Azure
* pmoriet:
* markturansky:
  * [https://github.com/kubernetes/kubernetes/pull/19600](https://github.com/kubernetes/kubernetes/pull/19600)
  * [https://github.com/kubernetes/kubernetes/pull/19921](https://github.com/kubernetes/kubernetes/pull/19921)
  * [https://github.com/kubernetes/kubernetes/pull/20197](https://github.com/kubernetes/kubernetes/pull/20197)
  * [https://github.com/kubernetes/kubernetes/pull/20213](https://github.com/kubernetes/kubernetes/pull/20213)
  * [https://github.com/kubernetes/kubernetes/pull/19868](https://github.com/kubernetes/kubernetes/pull/19868)
* wattsteve: I’d love to see the dynamic provision proposal get merged so we can start work on implementing that.
* Re-group in 1 week - mark to put on calendar

## December 2, 2015

Moderator: Mark Turansky
Agenda: Dynamic provisioning technical design discussion, other outstanding PR

* review [https://github.com/kubernetes/kubernetes/pull/17056](https://github.com/kubernetes/kubernetes/pull/17056)
