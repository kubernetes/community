# Kubernetes Storage SIG Meeting Notes (2019)

The Kubernetes Storage Special-Interest-Group (SIG) is a working group within the Kubernetes contributor community interested in storage and volume plugins. This document contains historical meeting notes from past meeting.

## December 19, 2019

Recording: [https://youtu.be/_LaA-LeRSb4](https://youtu.be/_LaA-LeRSb4)

Agenda/Note

* Q1 2020 v1.18 Planning
  * [Planning spreadsheet](https://docs.google.com/spreadsheets/d/1t4z5DYKjX2ZDlkTpCnp18icRAQqOE85C1T1r2gqJVck/edit?usp=sharing)
  * Start planning for next release.
* PRs to discuss or that need attention:
  * [Please add any PRs that need attention, with your name, below]
* Design Review
  * [Please add any design reviews, with your name, below]
  * KEP 1383 “bucket procvisioning” [https://github.com/kubernetes/enhancements/pull/1383](https://github.com/kubernetes/enhancements/pull/1383)
* Misc
  * None

## December 4, 2019

Recording: [https://youtu.be/ovljX8ICDVQ](https://youtu.be/ovljX8ICDVQ)

Agenda/Note

* Q4 2019 v1.17 Planning
  * [Planning spreadsheet](https://docs.google.com/spreadsheets/d/1t4z5DYKjX2ZDlkTpCnp18icRAQqOE85C1T1r2gqJVck/edit?usp=sharing)
  * End of quarter statu
  * **Monday, December 9 - Kubernetes 1.17.0 Released**
* PRs to discuss or that need attention:
  * [Please add any PRs that need attention, with your name, below]
* Design Review
  * [Please add any design reviews, with your name, below]
  * Consistency groups google doc: [https://docs.google.com/document/d/1nmq0rjt7A45w3kRHloDsMbzwGMELaJNNw2_w1B9qirk/edit](https://docs.google.com/document/d/1nmq0rjt7A45w3kRHloDsMbzwGMELaJNNw2_w1B9qirk/edit)
    * Provide feedback offline. Next step: create a KEP.
* Misc
  * [Xing/Dave/Xiangqian] Proposal for Kubernetes Data Protection Working Group
    * [https://docs.google.com/presentation/d/1qwsu82cDNJ06NvXJ5yDl9fPn8Ub9Sf7n9QIvabTOYh0/edit?usp=sharing](https://docs.google.com/presentation/d/1qwsu82cDNJ06NvXJ5yDl9fPn8Ub9Sf7n9QIvabTOYh0/edit?usp=sharing)
      * What will happen to Snapshot Workgroup? Do we need 2 groups?
        * Options:
          * 1 - Launch a new group, and deprecate the old group and take over the time slot.
          * 2 - Keep snapshot group, make it a sub-sub-group of data protection.
            * Concern: same set of people interested -- do we really want to double the number of meetings?
            * Everyone in snapshot group will want to be part of the new larger group.
            * Snapshot group is functional, but the new group is not. Do we want to break something that is working?
            * Snapshot group meetings -- not much design, it’s mostly implementation.
            * Workgroups are time bound -- not indefinite. We can put snapshot group on hold and consider dissolving it.
          * 3 - “transform” snapshot workgroup in to data protection group
            * Expand the scope and ownership (co-own with SIG Apps).
            * Concern: charter of new group is very large.
            * Lets have one large group for all of data protection discussions for both SIG Storage and SIG Apps. And we can hold smaller one off or recurring for implementation discussions. Use the existing snapshot workgroup meeting time.
              * Objection: Volume snapshot has specific topics that need discussions -- it still needs space.
                * Data Protection group more interested in higher level things like “what are we backing up”, “what are the workflows”… not lower level things like “volume snapshot code”, “not implementation”...
                * Implementation is not really discussed at the meeting, they sync offline already, do we need a work group.
              * Objection: Let’s start a new group, and deprecate the old one. New group will make it easier to attract new people.
            * 4- “freeze volume snapshot group” form new group, and then decide what to do afterwards.
              * Cancel meetings. And decide in the future what to do.
          * Conclusion: 4- put existing snapshot group on hold, create new group. Decide what to do with old group later
            * Next steps:
              * [Xing/Dave] Put together charter for new group, and do all the paper work.
              * [Xing] Set up meeting for new group at the same time as existing snapshot meeting. Wait for approval.
              * [Jing] Cancel meetings for existing snapshot group meeting.

## November 7, 2019

Recording: [https://youtu.be/8ENMzAUWs6s](https://youtu.be/8ENMzAUWs6s)

Agenda/Note

* [Announcement](https://groups.google.com/d/msg/kubernetes-sig-storage/NUIpoTJYqvY/N_bxmz0tAgAJ)
* Q4 2019 v1.17 Planning
  * [Planning spreadsheet](https://docs.google.com/spreadsheets/d/1t4z5DYKjX2ZDlkTpCnp18icRAQqOE85C1T1r2gqJVck/edit?usp=sharing)
* Video Memorial for Brad Child
  * Memories and thoughts of what Brad meant to you.
  * Videos to be shared with Brad’s family and may be part of a memorial montage at KubeCon
  * Alternate video options: see [Steve’s Message](https://groups.google.com/d/msg/kubernetes-sig-storage/YyLqYtyFZtE/f7JprlPYDwAJ).
  * [Memorial page](https://github.com/cncf/memorials/pull/1/files)
* Misc
  * Face to Face Meeting
    * Co-Located event at KubeCon NA 2019 (San Diego)
    * Date: Monday, November 18
    * Time: 9:00 AM - 5:00 PM
    * Venue: Marriott Marquis San Diego Marina
    * Registration: see sig-storage mailing list
    * [Agenda Items](https://docs.google.com/document/d/1qEIah61M_PrNKYBOJ3uZGFLRjLASz1je9XJkGY25O6g/edit#heading=h.qykpklbiq2qg)

## October 24, 2019

Recording: [https://youtu.be/tcM92rvGERY](https://youtu.be/tcM92rvGERY)

Agenda/Note

* Q4 2019 v1.17 Planning
  * [Planning spreadsheet](https://docs.google.com/spreadsheets/d/1t4z5DYKjX2ZDlkTpCnp18icRAQqOE85C1T1r2gqJVck/edit?usp=sharing)
  * Important Date
    * ~~Tuesday, October 15, EOD PST - Enhancements Freeze~~
    * **Thursday, November 14, EOD PST - Code Freeze**
    * Tuesday, November 19 - Docs must be completed and reviewed
    * Monday, December 9 - Kubernetes 1.17.0 Released
    * Shorter cycle due to Kubecon/Holidays in December
* PRs to discuss or that need attention:
  * [Please add any PRs that need attention, with your name, below]
* Design Review
  * [Please add any design reviews, with your name, below]
* Misc
  * [Saad Ali] Face to Face Meeting
    * Co-Located event at KubeCon NA 2019 (San Diego)
    * Date: Monday, November 18
    * Time: 9:00 AM - 5:00 PM
    * Venue: Marriott Marquis San Diego Marina
    * Registration: see sig-storage mailing list
    * [Agenda Items](https://docs.google.com/document/d/1qEIah61M_PrNKYBOJ3uZGFLRjLASz1je9XJkGY25O6g/edit#heading=h.qykpklbiq2qg)
  * Next meeting - Nov 7
    * Saad Ali will be OOO.
    * Brad Childs will host.
  * Meeting in 4 weeks - Nov 21 - is at same time as Kubecon
    * Will be canceled.

## October 10, 2019

Recording: [https://youtu.be/NYCUMQKW6oY](https://youtu.be/NYCUMQKW6oY)

Agenda/Note

* Q4 2019 v1.17 Planning
  * [Planning spreadsheet](https://docs.google.com/spreadsheets/d/1t4z5DYKjX2ZDlkTpCnp18icRAQqOE85C1T1r2gqJVck/edit?usp=sharing)
    * Q4 planning
  * Important Date
    * **Tuesday, October 15, EOD PST - Enhancements Freeze**
    * Thursday, November 14, EOD PST - Code Freeze
    * Tuesday, November 19 - Docs must be completed and reviewed
    * Monday, December 9 - Kubernetes 1.17.0 Released
    * Shorter cycle due to Kubecon/Holidays in December
* PRs to discuss or that need attention:
  * [Please add any PRs that need attention, with your name, below]
* Design Review
  * [Please add any design reviews, with your name, below]
* Misc
  * [Saad Ali] Face to Face Meeting
    * Co-Located event at KubeCon NA 2019 (San Diego)
    * Date: Monday, November 18
    * Time: 9:00 AM - 5:00 PM
    * Venue: Marriott Marquis San Diego Marina
    * Registration: see sig-storage mailing list
    * [Agenda Items](https://docs.google.com/document/d/1qEIah61M_PrNKYBOJ3uZGFLRjLASz1je9XJkGY25O6g/edit#heading=h.qykpklbiq2qg)

## September 26, 2019

Recording: [https://youtu.be/mbxvYC6EBwg](https://youtu.be/mbxvYC6EBwg)

Agenda/Note

* Q4 2019 v1.17 Planning
  * [Planning spreadsheet](https://docs.google.com/spreadsheets/d/1t4z5DYKjX2ZDlkTpCnp18icRAQqOE85C1T1r2gqJVck/edit?usp=sharing)
    * Q4 planning
  * Important Date
    * **Tuesday, October 15, EOD PST - Enhancements Freeze**
    * Thursday, November 14, EOD PST - Code Freeze
    * Tuesday, November 19 - Docs must be completed and reviewed
    * Monday, December 9 - Kubernetes 1.17.0 Released
    * Shorter cycle due to Kubecon/Holidays in December
* PRs to discuss or that need attention:
  * [Please add any PRs that need attention, with your name, below]
  * Support for Windows iSCSI FlexVolume driver in TargetD ext provisioner [https://github.com/kubernetes-incubator/external-storage/pull/1226](https://github.com/kubernetes-incubator/external-storage/pull/1226)
    * Jing Xu can help take a look as well as Jan.
* Design Review
  * [Please add any design reviews, with your name, below]
* Misc
  * [Saad Ali] Face to Face Meeting
    * Co-Located event at KubeCon NA 2019 (San Diego)
    * Date: Monday, November 18
    * Time: 9:00 AM - 5:00 PM
    * Venue: Marriott Marquis San Diego Marina
    * Registration: will require Kubecon registration, more details to be announced soon.

## September 12, 2019

Recording: [https://youtu.be/jL8A-nbiJYg](https://youtu.be/jL8A-nbiJYg)

Agenda/Note

* Q3 2019 v1.16 Planning
  * [Planning spreadsheet](https://docs.google.com/spreadsheets/d/1t4z5DYKjX2ZDlkTpCnp18icRAQqOE85C1T1r2gqJVck/edit?usp=sharing)
    * End of quarter wrap up.
    * Q4 planning next time.
  * Important Date
    * Thursday, August 29, EOD PST - Code Freeze
    * Monday, September 9 - Docs must be completed and reviewed
    * **Monday, September 16 - Kubernetes 1.16.0 Released**
* PRs to discuss or that need attention:
  * [Please add any PRs that need attention, with your name, below]
* Design Review
  * [Please add any design reviews, with your name, below]
* Misc
  * [Joshua Van Leeuwen] cert-manager CSI demo @JoshVanL
  * [Deep Debroy] Discuss if calls to dismount during graceful detach is necessary for all node OSes: [https://github.com/container-storage-interface/spec/issues/388](https://github.com/container-storage-interface/spec/issues/388) (specifically: [https://github.com/container-storage-interface/spec/issues/388#issuecomment-529155568](https://github.com/container-storage-interface/spec/issues/388#issuecomment-529155568))
    * Note
      * Should driver unmount at unstage?
      * Linux unmount flushes data.
        * This is assumed by CSI spec.
      * Windows unmount doesn’t flush data.
      * Yes, driver should flush data on unpublish calls. Let’s update CSI spec to clarify thi
      * The entity responsible for the filesystem should be responsible for ensuring data in caches is flushed before unpublish.
      * Counterpoint: if apps depend on this behavior they may not handle node crashes.
      * Yes, but CSI should still do “the right thing”.
    * Conclusion
      * Yes, driver should flush data on unpublish calls. Let’s update CSI spec to clarify this.
  * [Saad Ali]
    * Face to face meeting?
      * Kubecon San Diego Nov 19-21?
      * Earlier?
      * [Ben S] Netapp have space in RTP NC depending on how many people would come/time.
      * Vote between:
        * San Diego @ kubcon: 4
        * East coast in October: 1
        * Either: 2
      * Will try to get some space via CNCF in San Diego
      * What day? Monday before the conference?
        * Monday has contributor summit -- may be able to get a room there.
        * Monday: 5
        * Friday: 0

## August 29, 2019

Recording: [https://youtu.be/54U8r30oQhQ](https://youtu.be/54U8r30oQhQ)

Agenda/Note

* Q3 2019 v1.16 Planning
  * [Planning spreadsheet](https://docs.google.com/spreadsheets/d/1t4z5DYKjX2ZDlkTpCnp18icRAQqOE85C1T1r2gqJVck/edit?usp=sharing)
  * Important Date
    * **Thursday, August 29, EOD PST - Code Freeze**
    * Monday, September 9 - Docs must be completed and reviewed
    * Monday, September 16 - Kubernetes 1.16.0 Released
  * For features that did not land, please update enhancement issue to let release team know.
* PRs to discuss or that need attention:
  * [Please add any PRs that need attention, with your name, below]
* Design Review
  * [Please add any design reviews, with your name, below]
* Misc
  * N/A

## August 15, 2019

Recording: [https://youtu.be/mmARoySG5iY](https://youtu.be/mmARoySG5iY), [https://youtu.be/pG5chGtHrvs](https://youtu.be/pG5chGtHrvs)
Agenda/Note

* Q3 2019 v1.16 Planning
  * [Planning spreadsheet](https://docs.google.com/spreadsheets/d/1t4z5DYKjX2ZDlkTpCnp18icRAQqOE85C1T1r2gqJVck/edit?usp=sharing)
  * Important Date
    * Tuesday, July 30, EOD PST - Enhancements Freeze
    * Thursday, August 29, EOD PST - Code Freeze
    * Monday, September 9 - Docs must be completed and reviewed
    * Monday, September 16 - Kubernetes 1.16.0 Released
* PRs to discuss or that need attention:
  * [Please add any PRs that need attention, with your name, below]
* Design Review
  * [Please add any design reviews, with your name, below]
* Misc
  *

## August 1, 2019

Recording: [https://youtu.be/o_OU0Hh2Zac](https://youtu.be/o_OU0Hh2Zac)

Agenda/Note

* Q3 2019 v1.16 Planning
  * [Planning spreadsheet](https://docs.google.com/spreadsheets/d/1t4z5DYKjX2ZDlkTpCnp18icRAQqOE85C1T1r2gqJVck/edit?usp=sharing)
  * Important Date
    * Tuesday, July 30, EOD PST - Enhancements Freeze
    * Thursday, August 29, EOD PST - Code Freeze
    * Monday, September 9 - Docs must be completed and reviewed
    * Monday, September 16 - Kubernetes 1.16.0 Released
* PRs to discuss or that need attention:
  * [Please add any PRs that need attention, with your name, below]
* Design Review
  * [Please add any design reviews, with your name, below]
* Misc
  *

## July 18, 2019

Recording: [https://youtu.be/cbxR2L3Jx08](https://youtu.be/cbxR2L3Jx08)

Agenda/Note

* Q3 2019 v1.16 Planning
  * [Planning spreadsheet](https://docs.google.com/spreadsheets/d/1t4z5DYKjX2ZDlkTpCnp18icRAQqOE85C1T1r2gqJVck/edit?usp=sharing)
  * Important Date
    * Tuesday, July 30, EOD PST - Enhancements Freeze
    * Thursday, August 29, EOD PST - Code Freeze
    * Monday, September 9 - Docs must be completed and reviewed
    * Monday, September 16 - Kubernetes 1.16.0 Released
* PRs to discuss or that need attention:
  * [Please add any PRs that need attention, with your name, below]
* Design Review
  * [Please add any design reviews, with your name, below]
* Misc
  * [Michelle] Deprecate kubelet attach/detach: [https://github.com/kubernetes/kubernetes/issues/55517](https://github.com/kubernetes/kubernetes/issues/55517)
    * Anyone using enable-controller-attach-detach parameter on kubelet? Please reach out.
    * Haley: Flex volume using it, but trying to deprecate.
    * Another challenge: some use this for bootstrapping master control plane.
      * Alternatives: EmptyDir or Hostpath
      * Need to talk to people who are depending on this.
      * Please reach out if you care about this.
    * Will broadcast to mailing list.
  * [Cheng] Moving controller manager to distroless - actions to take on Flexvolume master-side API
    * Survey sent to SIG to ask if anyone is using flexVolume master APIs, no response.
    * Next step, will reach out to kubernetes-dev, if no response will not support in distroless.
    * Context: shifting binaries to distroless which is missing linux binaries -- so for drivers that expect those to exist on master may break.
  * [Jan] Deprecate CSI cluster-driver-registrar[https://github.com/kubernetes-csi/cluster-driver-registrar/issues/48](https://github.com/kubernetes-csi/cluster-driver-registrar/issues/48)
    * 1.14 and 1.15 didn’t ship
    * Deprecation announcement made 2 weeks ago. It is no longer supported don’t use it.
    * Alternative: create CSIDriver object in installation manifest or installation of driver. Kubernetes-csi docs being updated to reflect this.
  * [Youssef] PureStorage thread on migration -- what if there are PVs already created or mounted.
    * Existing PVC API remains same, when migration enabled, in-tree shims silently redirect to CSI Driver
    * What about flex volumes? Not currently supported by migration.
      * Have a CSI driver, how do we update a cluster using Flex to CSI, given PVCs could have been created with Flex.
      * Two approache
        * 1) Create new PVs with CSI parameters and swap
          * Kubelet has a cache -- even if PVs updated, sometimes you get call to Flex or CSI. Depending on kubelet cache.
          * Stop using volumes and unmount first. Problem: some environments can’t control that
            * Migrate in place, maintain dual drivers, and then force a restart on pods.
            * Once pods are restarted, kubelet cache will be updated.
          * Existing github issue -- feel free to comment on that. If this works, can you share with the community how to do this.
        * 2) create an “adapter” to translate on the fly without persisting translations.
      * External-provisioner is separate -- so that would

## June 28, 2019

Recording: [https://youtu.be/sS3tmlJn_VE](https://youtu.be/sS3tmlJn_VE)

Agenda/Note

* Q3 2019 v1.16 Planning
  * [Planning spreadsheet](https://docs.google.com/spreadsheets/d/1t4z5DYKjX2ZDlkTpCnp18icRAQqOE85C1T1r2gqJVck/edit?usp=sharing)
* PRs to discuss or that need attention:
  * [Please add any PRs that need attention, with your name, below]
  * [https://github.com/kubernetes/enhancements/pull/1112](https://github.com/kubernetes/enhancements/pull/1112) jgriffith
    * Start reviewing early to unblock API review.
* Design Review
  * [Please add any design reviews, with your name, below]
* Misc
  * [Hakan Memisoglu] CSI Spec issue - Calls for CreateVolume have secret, problem read operations like “List” have no operations.
    * Add to CSI community meeting agenda.
  * [John Griffithh] Data Populator Design -- how to proceed?
    * Will set up a meeting
  * Another Face to Face Meeting?
    * Anyone interested in hosting?
    * Topic
      * Who should deploy/manage side car container
      * How to push down credentials in to CSI without having it exposed ot namespace of container
        * CSI Spec issue 370 -
      * Volume Group API (consistency groups, storage topology, etc.)

## June 20, 2019

Recording: [https://youtu.be/w4M5u_3j8As](https://youtu.be/w4M5u_3j8As)

Agenda/Note

* Q3 2019 v1.16 Planning
  * [Planning spreadsheet](https://docs.google.com/spreadsheets/d/1t4z5DYKjX2ZDlkTpCnp18icRAQqOE85C1T1r2gqJVck/edit?usp=sharing)
* PRs to discuss or that need attention:
  * [Please add any PRs that need attention, with your name, below]
* Design Review
  * [Please add any design reviews, with your name, below]
* Misc

## June 6, 2019

Recording: [https://youtu.be/1baPPiahZbY](https://youtu.be/1baPPiahZbY)

Agenda/Note

* Q2 2019 v1.15 Update
  * [Planning spreadsheet](https://docs.google.com/spreadsheets/d/1t4z5DYKjX2ZDlkTpCnp18icRAQqOE85C1T1r2gqJVck/edit?usp=sharing)
  * Important dates:
    * Thursday, May 30th, EOD PST - Code Freeze
    * Tuesday, June 04 - Docs must be completed and reviewed
    * Monday, June 17th - Kubernetes 1.15.0 Released
* PRs to discuss or that need attention:
  * [Please add any PRs that need attention, with your name, below]
  * [jinxu@] [https://github.com/kubernetes/kubernetes/issues/72048](https://github.com/kubernetes/kubernetes/issues/72048)
    * Proposed fix: force unmount
      * may be dangerous if process using mount
      * Does force option actually fix NFS unmount issue? May not actually. Doc says it may not work.
      * Hard mount means block until server comes back.
    * Other options:
      * Restart NFS server
      * User manually unmount
    * Can we understand why the NFS server deleted? Maybe network partition.
    * Can we unblock pod deletion, but save garbage collection for later. If we do, new pod may fail too. Depends on the failure cause.
    * Users can force kill pod, and kubelet orphaned pod tries to clean up? If NFS server can’t be reached. We don’t start more then 1 unmount operation.
    * Conclusion: no need for force unmount. Just tell user to force delete pod. Kubelet should clean up mount when nfsserver comes back or reboot.
    * Even if we clean up user space resources, it won’t clean up hung kernel space resources.
    * If user really wants force unmount: create a CSI driver that does this (preferably optionally). Dangerous to do it for everything.
* Design Review
  * [Please add any design reviews, with your name, below]
  * [nickren/xing-yang] PV health monitor KEP: [https://github.com/kubernetes/enhancements/pull/1077](https://github.com/kubernetes/enhancements/pull/1077)
    * Please review KEP if you’re interested in volume health monitoring.
* Misc
  * [msau] Help wanted! StatefulSet and PVC Protection integration issue: [https://github.com/kubernetes/kubernetes/issues/74374](https://github.com/kubernetes/kubernetes/issues/74374)
  * [verult] Flex drivers on master doesn’t work with distrole
        [https://github.com/kubernetes/kubernetes/issues/78737](https://github.com/kubernetes/kubernetes/issues/78737)
    * We can suggest that users move to CSI but can’t require it.
    * Deprecation policy is 1+ year.
    * Maybe take a survey in SIG.

## May 23, 2019

Agenda/Note

* Cancelled. Lots of people out at Kubecon EU in Barcelona.

## May 9, 2019

Recording: [https://youtu.be/_AqgfdfXBBQ](https://youtu.be/_AqgfdfXBBQ)

Agenda/Note

* Q2 2019 v1.15 Update
  * [Planning spreadsheet](https://docs.google.com/spreadsheets/d/1t4z5DYKjX2ZDlkTpCnp18icRAQqOE85C1T1r2gqJVck/edit?usp=sharing)
* PRs to discuss or that need attention:
  * [Please add any PRs that need attention, with your name, below]
* Design Review
  * [Please add any design reviews, with your name, below]
* Misc
  * Kubecon EU in Barcelona
  * Lots of Storage sessions. See [slide 4 here](https://docs.google.com/presentation/d/1IApIw92SNEt-iE8YZTnP0cxecrj0LQ-3kbDh70E661c/edit#slide=id.g5914713db9_0_3).
  * Monday, May 20- Kubernetes Contributor Summit
    * Register [here](https://events.linuxfoundation.org/events/contributor-summit-europe-2019/register/).
  * Cancel meeting Kubecon week?
    * Will consider next meeting cancelled.

## April 25, 2019

Recording: [https://youtu.be/T8fVVnnzP3U](https://youtu.be/T8fVVnnzP3U)

Agenda/Note

* Q2 2019 v1.15 Update
  * [Planning spreadsheet](https://docs.google.com/spreadsheets/d/1t4z5DYKjX2ZDlkTpCnp18icRAQqOE85C1T1r2gqJVck/edit?usp=sharing)
  * Important dates: [https://github.com/kubernetes/sig-release/tree/master/releases/release-1.15#tldr](https://github.com/kubernetes/sig-release/tree/master/releases/release-1.15#tldr)
    * Tuesday, April 30th, EOD PST - Enhancements Freeze
      * An open [enhancement issue](https://github.com/kubernetes/enhancements/issues?q=is%3Aopen+is%3Aissue+milestone%3Av1.15) in the 1.15 Milestone
      * All Enhancements must have a [KEP](https://github.com/kubernetes/enhancements/tree/master/keps/sig-storage) that is in an implementable state by enhancement freeze. If the enhancement does not have a KEP in an implementable state by enhancement freeze it will be removed from the milestone and will require an exception.
* PRs to discuss or that need attention:
  * [Please add any PRs that need attention, with your name, below]
* Design Review
  * Object bucket dyn provisioner (Erin, Jon, Jeff, Scott Creeley)
    * [https://github.com/yard-turkey/lib-bucket-provisioner](https://github.com/yard-turkey/lib-bucket-provisioner) [https://github.com/yard-turkey/aws-s3-provisioner](https://github.com/yard-turkey/aws-s3-provisioner) [https://github.com/yard-turkey/lib-bucket-provisioner/blob/master/doc/design/object-bucket-lib.md](https://github.com/yard-turkey/lib-bucket-provisioner/blob/master/doc/design/object-bucket-lib.md)
* Misc
  * [Ben Swartzlander] CSI plugin version compatibility - If the 1.1 sidecars only support 1.14+, how do we support 1.13 while also implementing new features?
    * You can use 1.1 CSI sidecar with k8s 1.13 but 1.14 specific features (resize) won’t work
    * Michelle Au is putting together a version support doc for CSI sidecars, stay tuned.

## April 11, 2019

Recording: [https://youtu.be/taR03H6jAEc](https://youtu.be/taR03H6jAEc)

Agenda/Note

* Q2 2019 v1.15 Update
  * [Planning spreadsheet](https://docs.google.com/spreadsheets/d/1t4z5DYKjX2ZDlkTpCnp18icRAQqOE85C1T1r2gqJVck/edit?usp=sharing)
  * Important dates: [https://github.com/kubernetes/sig-release/tree/master/releases/release-1.15#tldr](https://github.com/kubernetes/sig-release/tree/master/releases/release-1.15#tldr)
* PRs to discuss or that need attention:
  * [Please add any PRs that need attention, with your name, below]
* Design Review
  * [Please add any design reviews, with your name, below]
* Misc
  * Brief Notes from HEPiX Workshop in San Diego
    * FYI: The HEPiX forum brings together worldwide Information Technology staff, including system administrators, system engineers, and managers from the High Energy Physics and Nuclear Physics laboratories and institute
    * [Attendees from many including](https://www.eiseverywhere.com/ehome/201903-hepix/850665/) : SLAC, BNL, FermiLab, DESY, CERN, etc
    * Many Talks were Storage Related ([See Agenda/Downloads here](https://indico.cern.ch/event/765497/timetable/#all.detailed))
    * Container RT and Orchestration tools used by
      * Container RT: Docker, Singularity
      * Container Orchestration: Kubernetes, HTCondor, SLURM
  * Singularity
    * Used Extensively in HEP Community (HIgh Energy Physics / HPC)
    * [Singularity Kubernetes CRI implementation](https://github.com/sylabs/singularity-cri) in Alpha
    * QUESTION:
      * Has anyone played with this?
      * Any experience using CSI with Singularity CSI?
      * I have a call with the Sylab folk (Singularity)  Tomorrow (Friday)
        * If anyone has any items/questions that you’d like me to bring up with them, please let me know (gerry@auristor.com)

## March 28, 2019

Recording: [https://youtu.be/7jeelUiyCcg](https://youtu.be/7jeelUiyCcg)

Agenda/Note

* Q1 2019 v1.14 Update
  * Q1 v1.14 Wrap up
  * Q2 v1.15 Planning
  * [Planning spreadsheet](https://docs.google.com/spreadsheets/d/1t4z5DYKjX2ZDlkTpCnp18icRAQqOE85C1T1r2gqJVck/edit?usp=sharing)
* PRs to discuss or that need attention:
  * [Please add any PRs that need attention, with your name, below]
* Design Review
  * [Please add any design reviews, with your name, below]
* Misc
  * New API Review proce
    * [https://git.k8s.io/community/sig-architecture/api-review-process.md](https://git.k8s.io/community/sig-architecture/api-review-process.md) Questions, concerns, comments -&gt; visit sig architecture.
  * Windows storage roadmap proposal for 1.15+
    * [https://docs.google.com/document/d/1odyInziM5iu44zuXahxC1utcbKo-03gsQvP1OpDYGT4/edit#](https://docs.google.com/document/d/1odyInziM5iu44zuXahxC1utcbKo-03gsQvP1OpDYGT4/edit#)
    * CSI node plugin support for Windows, in-tree SMB plugin + iSCSI enhancements for Windows, CSI plugins for SMB
  * Annotate PVC, using provisioning secret templating
    * Move to next meeting - ran out of time

## March 14, 2019

Recording: [https://youtu.be/L1Ju3SuUozk](https://youtu.be/L1Ju3SuUozk)

Agenda/Note

* Q1 2019 v1.14 Update
  * [Planning spreadsheet](https://docs.google.com/spreadsheets/d/1t4z5DYKjX2ZDlkTpCnp18icRAQqOE85C1T1r2gqJVck/edit?usp=sharing)
  * (bchilds) KubeCon EU contributor summit sig-storage presentation proposal
    * CFP Deadline April 1st... send to [bchilds@redhat.com](mailto:bchilds@redhat.com)
      * Please title [sig-storage CFP contrib] + description
    * Good Topic
      * CSI Driver Development
      * Operator development
      * Feature design community experience
      * Any community, design or PR logistical issue
  * March 28 - Plan for 1.15
* PRs to discuss or that need attention:
  * [Simon Croome] StorageOS Volume Expansion <https://github.com/kubernetes/kubernetes/pull/74351>
* Design Review
  * [Please add any design reviews, with your name, below]
* Misc
  * Would like to see online volume resize into Beta in 1.15 (Srini)
  * Issue [https://github.com/kubernetes-csi/external-provisioner/issues/86 (Srini/Akash)](https://github.com/kubernetes-csi/external-provisioner/issues/86)
    * Annotations in PVC is something that should be avoided
    * Let’s talk about use cases and find workaround
      * Akash: K8s zones and regions in PVCs -- users decide where to provision PVC.
        * Use Volume Scheduling feature: “LateBinding” on StorageClass -- scheduler decision for pods is passed to volume creation; all of a pods constraints are used -- so add pod anti-affinity to pods, let the scheduler handle scheduling. At provision time scheduler will tell storage system what zone to provision in.
      * MattSmith: storage tenancy that aligns with k8s namespaces. Want to align storage tenancy with kubernetes -- need namespace information.
        * Problem: PV object is non-namespaced. Consider moving PVCs across namespaces? Multiple PVC in different namespaces?
      * Akash: each user will use own secret to provision volume. Due to number of users, too many storage classe
        * Solution: provisioning secret templating.
          * See [https://kubernetes-csi.github.io/docs/secrets-and-credentials.html#createdelete-volume-secret](https://kubernetes-csi.github.io/docs/secrets-and-credentials.html#createdelete-volume-secret)
          * Plans to make this more flexible.
            * [https://github.com/kubernetes-csi/external-provisioner/issues/233](https://github.com/kubernetes-csi/external-provisioner/issues/233)
      * Simon: multiple config params can result in explosion of storageClasses -- which are administered by ClusterAdmin not by App Admin.
      * ShayBerman: setting filesystem type in PVC level.
        * How many filesystems do you have? 2-4.
      * JoseRivera: controller and CRD to abstract away StorageClass explosion.
        * Ardalan: Important to remember portability at that level as well.
        * MattSmith: can we get an example. Ben: looks at sample controller example.
  * [https://github.com/kubernetes/kubernetes/issues/62778](https://github.com/kubernetes/kubernetes/issues/62778) (Srini)

## February 28, 2019

Recording: [https://youtu.be/69yBUQYm4fY](https://youtu.be/69yBUQYm4fY)

Agenda/Note

* Q1 2019 v1.14 Update
  * [Planning spreadsheet](https://docs.google.com/spreadsheets/d/1t4z5DYKjX2ZDlkTpCnp18icRAQqOE85C1T1r2gqJVck/edit?usp=sharing)
* PRs to discuss or that need attention:
  * [Please add any PRs that need attention, with your name, below]
  * [andrewsykim] promoting beta zone/region labels for volumes/nodes to GA [https://github.com/kubernetes/enhancements/pull/839](https://github.com/kubernetes/enhancements/pull/839)
    * Motivation: clean-up. Labels have been there since 1.3
    * For node:
      * Still required for node even if we don’t do it for volumes.
    * For storage:
      * Would prefer to deprecate
        * Need an answer for Migrating old PVs to new topology feature
          * Etcd upgrade script only works if object version is bumped.
          * May need to work with API Machinery to figure out a way to do this.
      * Post-topology work to enable existing PVs to use new topology.
* Design Review
  * [Please add any design reviews, with your name, below]
* Misc
  *

## February 14, 2019

Recording: [https://youtu.be/0yn_fbnChdM](https://youtu.be/0yn_fbnChdM)

Agenda/Note

* Q1 2019 v1.14 Update
  * [Planning spreadsheet](https://docs.google.com/spreadsheets/d/1t4z5DYKjX2ZDlkTpCnp18icRAQqOE85C1T1r2gqJVck/edit?usp=sharing)
* PRs to discuss or that need attention:
  * Image Driver - [https://github.com/kubernetes-csi/drivers/pull/133](https://github.com/kubernetes-csi/drivers/pull/133)
  * Undocumented EBS Volume ID format:  [https://github.com/kubernetes/kubernetes/issues/73730](https://github.com/kubernetes/kubernetes/issues/73730)
    * awsElasticBlockStore.volumeID has two formats. That causes issues.
    * Questions about how to proceed for migration.
* Design Review
  * User Namespaces  - [https://github.com/kubernetes/community/pull/2595](https://github.com/kubernetes/community/pull/2595)
    * Proposing to shift the userId and GID of all files in volume to make it accessible to UID in container.
    * If pod asks for remapped namespace, and it is not possible to move the UID or GID of the volumes (for example, Not all volumes support user namespace remapping, e.g. NFS), then containers can’t use those volumes.
    * For pod containers asking for hostpath should block remapping
    * Change in ownership of files -- container UID 3 mapped to 1003 outside container. Propose changing all file ownership to 1003
      * Issue
        * 1) containers may all be running as different UID
    * Counter proposal: Can we just use FSGroups?
      * If no FSGroup specified -- no change of ownership
        * Volume can’t be written to.
      * If FSGroup is specified -- adjust it to based on UID mapping
    * What about multiwriter volumes?
      * Now that we have CSI we don’t know which drivers support FSGroup or not -- CSI assumes any single writer consumed as mounted file supports FSGroup and multi-writer does not support FSGroup
    * There is an option to use user namespace remapping? Yes in pod spec not per container.
    * Kubelet after mounting -- does a recursive chown -- causes performance issues. Even with remapped UID this will still be an issue.
      * SIG Storage proposal on how to skip it. [https://github.com/kubernetes/enhancements/pull/696](https://github.com/kubernetes/enhancements/pull/696) (Will only work for PVC, need something more generic that will work for all volumes). Will have to live with perf hit for now.
    * Can remapping be different per node (i.e. pod rescheduled to different node)? It can be.
      * Maybe document best practices?
* [Vault - Linux Storage and Filesystems Conference](https://www.usenix.org/conference/vault19) Boston Feb 25-26
  * Is anyone planning to attend?
    * Attendees:
      * Gerry Seidman (gerry@auristor.com)
      * Vasily Tarasov (vtarasov@us.ibm.com). Me and my colleagues are giving an introductory tutorial on container storage.
  * The AuriStor/AFS CSI driver should be announced/shown (PoC) to go along with Fedora-29 including the kAFS kernel driver
  * **NOTE**: David Howells from Red Hat’s Storage/Security Group will be there which may be of general storage issues for things like his work on [FsCache, Kernel keyrings, etc](https://www.infradead.org/~dhowells/kafs/)
  * Suggestion: label nodes. Delay provisioning of PV (i.e. block CreateVolume) until it is understood which node data must reside on. Label PV with node labels.

## January 31, 2019

Recording: [https://youtu.be/t0svw4xjwRc](https://youtu.be/t0svw4xjwRc)

Agenda/Note

* Q1 2019 v1.14 Update
  * [Planning spreadsheet](https://docs.google.com/spreadsheets/d/1t4z5DYKjX2ZDlkTpCnp18icRAQqOE85C1T1r2gqJVck/edit?usp=sharing)
* PRs to discuss or that need attention:
* Design Review
* [Vault - Linux Storage and Filesystems Conference](https://www.usenix.org/conference/vault19) Boston Feb 25-26
  * Is anyone planning to attend?
    * Attendees:
      *
  * The AuriStor/AFS CSI driver should be announced/shown (PoC) to go along with Fedora-29 including the kAFS kernel driver
  * Suggestion: label nodes. Delay provisioning of PV (i.e. block CreateVolume) until it is understood which node data must reside on. Label PV with node labels.

## January 17, 2019

Recording: [https://youtu.be/aDWVFqL3x5g](https://youtu.be/aDWVFqL3x5g)

Agenda/Note

* Q1 2018 v1.14 Update
  * [Planning spreadsheet](https://docs.google.com/spreadsheets/d/1t4z5DYKjX2ZDlkTpCnp18icRAQqOE85C1T1r2gqJVck/edit?usp=sharing)
* PRs to discuss or that need attention:
  * [Jing] [Fix hang when unmounting disconnected nfs volume](https://github.com/kubernetes/kubernetes/pull/72049)
  * [Pod is stuck in Terminating status forever after Kubelet restart](https://github.com/kubernetes/kubernetes/issues/72604)
  * [Hemant] [Add Cinder Max Volume Limit](https://github.com/kubernetes/kubernetes/pull/72980)
  * [\[Jan\] Include namespace name in random distribution of PVs among availability zones](https://github.com/kubernetes/kubernetes/pull/72736)
* Design Review
  * Deprecate cloudprovider specific predicates.
    * Deprecating the predicates is fine but we need to ensure that CSI migration takes into account

## January 3, 2019

Recording: None.

Agenda/Note

* [Bradley Childs] Q1 2018 v1.14 Update
  * [Planning spreadsheet](https://docs.google.com/spreadsheets/d/1t4z5DYKjX2ZDlkTpCnp18icRAQqOE85C1T1r2gqJVck/edit?usp=sharing)
  * Q1 planning
* PRs to discuss or that need attention:
  * [Please add any PRs that need attention, with your name, below]
* Design Review
  * [Please add any design reviews, with your name, below]
