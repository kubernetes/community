# Kubernetes Storage SIG Meeting Notes (2017)

The Kubernetes Storage Special-Interest-Group (SIG) is a working group within the Kubernetes contributor community interested in storage and volume plugins. This document contains historical meeting notes from past meeting.

## December 28, 2017

Cancelled. Happy holidays!

## December 14, 2017

Recording: [https://youtu.be/_a_pwZKgtT8](https://youtu.be/_a_pwZKgtT8)

Agenda/Notes (Please include your name and estimated time):

* [Saad Ali, 20 minutes] 1.9 Q4 End-of-quarter Status Update
  * [Planning spreadsheet](https://docs.google.com/spreadsheets/d/1t4z5DYKjX2ZDlkTpCnp18icRAQqOE85C1T1r2gqJVck/edit?usp=sharing)
* PRs to discuss or that need attention
  * [Please add any PRs that need attention here]
* Design Review
  * [Please add any design reviews here]
* [Saad Ali, 10 minutes] Set Up a one-off “on-boarding to sig-storage” meeting
  * Proposal by Michael Rubin
  * Suggested topics (presentations/YouTube):
    * Local dev environment setup
    * How to contribute/What does the SIG need help with/appreciate?
    * How to propose and add new feature
    * How to add a volume plugin? In-tree vs Flex vs CSI vs External Provisioner
      * How to write a volume plugin?
    * How to get a bug fixed?
    * SIG Planning process?
    * Where are the docs? The code?
    * High level overview of storage subsystem -- controller, api?
    * How to write tests (E2E, Unit, etc.)?
    * How to make an API change?
  * Volunteers for presentations:
    * Saad Ali
    * Steve Wong
    * Vladimir Vivien
    * Jan Safranek (after Christmas)
    * Erin Boyd
    * Scott C.
  * Proposals:
    * Storage 101 “class” and Slack?
    * “Buddy system” -- answer one-on-one question
    * Paris had an idea for “group mentoring”, one-on-one can be intimidating
    * Check with contributor SIG and see if they have idea
    * Documentation of how volume plugin interface arch -- should merge
    * K8s.io docs site should be tailored for different audiences: storage vendor vs users, etc.
* [Saad Ali, 5 minutes] Meeting December 28, 2017? Or skip that and meet Jan 4?
  * [Saad Ali, 20 minutes] 1.10 Q1 Planning Kickoff
    * Suggest delaying to first week of Jan
  * Conclusion: let’s do it
* [Yassine Tijani  5 minutes] discuss CSI timeline and coordination with wg-cloud-provider
* [Erin Boyd, 5 minutes] Prioritization of Raw Block Plugin Update
  * Seems like there is a lot of community PRs around these, let’s coordinate :)
  * Please capture any work related to Storage SIG, please add to planning spreadsheet
  * Flex will have block? If CSI takes too long?
* [Saad Ali] CSI Statu
  * [https://docs.google.com/document/d/1-WmRYvqw1FREcD1jmZAOjC0jX6Gop8FMOzsdqXHFoT4/edit#](https://docs.google.com/document/d/1-WmRYvqw1FREcD1jmZAOjC0jX6Gop8FMOzsdqXHFoT4/edit#)
  * CSI community/meeting info: [https://github.com/container-storage-interface/community](https://github.com/container-storage-interface/community)
* [Harry Zhang, 6 min] CSI support for KataContainers (hypervisor based container runtime)
  * How we implement it today
    * Already support host raw device and mount point
    * While most common use case of k8s is remote storage:
      * E.g. rbd/nfs/cif
      * We use flexvolume to do this for now: [https://github.com/kubernetes/frakti/blob/master/pkg/flexvolume/flexvolume.go](https://github.com/kubernetes/frakti/blob/master/pkg/flexvolume/flexvolume.go)
  * What we expect from CSI
    * use CSI to support remote storage case
      * [https://docs.google.com/document/d/19ZlbX1e7GSYxh_wdk0EZlhEheTelGavo6JIIrpN9cbs/edit](https://docs.google.com/document/d/19ZlbX1e7GSYxh_wdk0EZlhEheTelGavo6JIIrpN9cbs/edit)
      * [https://docs.google.com/presentation/d/1kPeia7wLqoKQI0oX4pvVdH1UpcPx3lpmFK4P_E6oiIc/edit#slide=id.g2c2e661992_0_90](https://docs.google.com/presentation/d/1kPeia7wLqoKQI0oX4pvVdH1UpcPx3lpmFK4P_E6oiIc/edit#slide=id.g2c2e661992_0_90)
      * less change in CRI shims (eventually no changes)
        * So Kata can work with cri-o, cri-containerd etc
      * Easier to migrate Kata to public cloud
* [mayank] RBD Provisioner - in-tree vs external
  * No, not in cloud provider.
* [lpabon] CSI testing
  * This is currently being document here: [https://github.com/kubernetes-csi/docs/wiki/Testing](https://github.com/kubernetes-csi/docs/wiki/Testing)
    * It will include how clients can be tested and how drivers can be tested.

## November 30, 2017

Recording: [https://youtu.be/XK5KVj0SbOE](https://youtu.be/XK5KVj0SbOE)

Agenda/Notes (Please include your name and estimated time):

* [Saad Ali, 20 minutes] 1.9 Q4 Statu
  * [Planning spreadsheet](https://docs.google.com/spreadsheets/d/1t4z5DYKjX2ZDlkTpCnp18icRAQqOE85C1T1r2gqJVck/edit?usp=sharing)
* PRs to discuss or that need attention
  * [Please add any PRs that need attention here]
* Design Review
  * [Please add any design reviews here]
* Kubecon
  * Attendees:
    * Michelle Au, Google
    * Jing Xu, Google
    * Michael Rubin, Google
    * Adam Litke, Red Hat [alitke@redhat.com](mailto:alitke@redhat.com), 507-884-4657
    * Jan Safranek, Red Hat [jsafrane@redhat.com](mailto:jsafrane@redhat.com)
    * Tomas Smetana, Red Hat tsmetana[@redhat.com](mailto:jsafrane@redhat.com)
    * Gerry Seidman, AuriStor, [gerry@auristor.com](mailto:gerry@auristor.com) 917-501-8287
    * Sean McGinnis, Huawei, [sean.mcginnis@gmail.com](mailto:sean.mcginnis@gmail.com) 612-386-7883
    * Ardalan Kangarlou, NetApp, ardalan@netapp.com
    * Steve Wong, {code}, [steven.wong@dell.com](mailto:steven.wong@dell.com), 562-417-7048, @cantbewong, present Sun-Sat
    * Yunwen Bai, Huawei. [yunwen.bai@huawei.com](mailto:yunwen.bai@huawei.com) 206-321-8643
    * Mitsuhiro Tanino, Hitachi Vantara, [mitsuhiro.tanino@hitachivantara.com](mailto:mitsuhiro.tanino@hitachivantara.com)
  * Storage SIG AMA
    * Thursday December 7, 2pm-3:30, Room 10C level 3 [http://sched.co/CU8l](http://sched.co/CU8l)
  * CNCF Storage WG
    * Face to face, Tuesday, 5pm-7pm, Mezzanine Room 1 level 2 [http://sched.co/D5HO](http://sched.co/D5HO)
  * Talks:
    * Tomas Smetana, Red Hat: [https://kccncna17.sched.com/event/CU7O](https://kccncna17.sched.com/event/CU7O)
    * Kubernetes Storage Evolution, Michelle and Erin: <http://sched.co/CU7R>
    * Local Ephemeral Storage Management, Jing Xu [http://sched.co/CU7X](http://sched.co/CU7X)
    * Block Volumes Support in Kubernetes, Mitsuhiro Tanino,

## November 23, 2017

Agenda/Notes (Please include your name and estimated time):

* [Saad Ali] Cancelled due to U.S. Thanksgiving Holiday
* [Saad Ali] Shifting meeting by 1 week (since 2 weeks from Nov 23 will overlap with Kubecon).

## November 9, 2017

Recording: [https://youtu.be/Xeup_ATnVEY](https://youtu.be/Xeup_ATnVEY)

Agenda/Notes (Please include your name and estimated time):

* [Saad Ali, 20 minutes] 1.9 Q4 Planning
  * [Planning spreadsheet](https://docs.google.com/spreadsheets/d/1t4z5DYKjX2ZDlkTpCnp18icRAQqOE85C1T1r2gqJVck/edit?usp=sharing)
* [Saad Ali] [Kubernetes CSI Implementation Meeting Notes](https://docs.google.com/document/d/1-WmRYvqw1FREcD1jmZAOjC0jX6Gop8FMOzsdqXHFoT4/edit#heading=h.vun6eaytazit)
* PRs to discuss or that need attention
  * [Please add any PRs that need attention here]
  * Fix dangling volumes - [https://github.com/kubernetes/kubernetes/issues/52573](https://github.com/kubernetes/kubernetes/issues/52573)
* Design Review
  * [Please add any design reviews here]
  * [Sandeep Pissay, 20 mins] Review proposal for “Backup of PVs” - Proposal [doc](https://docs.google.com/document/d/1Xkh7AzYQoVnFowW1pBBC6yLWLXQGmsqWr1L_5bSSh34/edit?usp=sharing)
  * [Yunwen Bai, 5 minutes] Dynamically provisioning of local storage proposal. Design doc to be shared EOW
* Kubecon Storage-SIG planning (AMA)
  * Reminder about [SIG Storage AMA at KubeCon](https://kccncna17.sched.com/event/CU8l/kubernetes-storage-ama-hosted-by-stephen-watt-red-hat). Please come and share your expertise with each other and with the community. Great place to have high bandwidth discussions. Intended to just be an open hangout, but the room comes equipped with flipcharts for drawing as well as AV.

## October 26, 2017

Agenda/Notes (Please include your name and estimated time):

* [Eric Forgette] Intro to [Dory, a Flex driver for Docker Volume Plugins](https://github.com/hpe-storage/dory)
* [Saad Ali] ***Feature freeze for 1.9 is tomorrow, Oct 27 ([link](https://github.com/kubernetes/features/blob/master/release-1.9/release-1.9.md))*** -- if you have a feature in the planning spreadsheet that should be tracked for k8s v1.9, please make sure there is a [issue opened in the feature repo](https://github.com/kubernetes/features/issues/) and [marked with the 1.9 milestone](https://github.com/kubernetes/features/issues?q=is%3Aopen+is%3Aissue+milestone%3A1.9).exte
* [Brad Childs, 20 minutes] 1.9 Q4 Planning
  * [Planning spreadsheet](https://docs.google.com/spreadsheets/d/1t4z5DYKjX2ZDlkTpCnp18icRAQqOE85C1T1r2gqJVck/edit?usp=sharing)
* PRs to discuss or that need attention
  * [Please add any PRs that need attention here]
* Design Review
  * Flexvolume driver installation -  [Additional default directories](https://docs.google.com/document/d/1U1oUgZmcGWrLMs4QlmPfLOElyI6D-Gp_6a-KlWRMAjs) - (Rook)
* ~~[Michael Rubin] Growing the sig-storage community~~
* [Gerry Seidman]  AuriStor (AFS) Flex Volume
  * I (and other engineers from AurIStor) will be at LISA next week in San Francisco.
  * If anyone planning on going, it would be great if we could meet up.
    * I’d appreciate any feedback on our Volume plugin implementation plans.
    * Please contact me at [gerry@auristor.com](mailto:gerry@auristor.com) / 917-501-8287
  * We are also holding an open [BoF at LISA](https://www.usenix.org/conference/lisa17/bofs#auristor) (Wednesday, November 1, 8-9PM)
    * Please come by if you are local.

## September 28, 2017

Recording: [https://youtu.be/f69IsTKtuvE](https://youtu.be/f69IsTKtuvE)

Agenda/Notes:

* [Palak Dalal] 1.9 Q4 Planning
  * [Planning spreadsheet](https://docs.google.com/spreadsheets/d/1t4z5DYKjX2ZDlkTpCnp18icRAQqOE85C1T1r2gqJVck/edit?usp=sharing)
* [Eric Forgette] Intro to [Dory, a Flex driver for Docker Volume Plugins](https://github.com/hpe-storage/dory)
* PRs to discuss or that need attention
  * [Please add any PRs that need attention here]
* Design Review
  * [Please add any design reviews here]
* [Saad Ali] Testing &amp; Documentation
* [Saad Ali] Next F2F meeting update
  * Dates: Oct 10-11, 2017
  * Location confirmed: Google 345 Spear Street San Francisco, CA  94105
  * Details: [link](https://docs.google.com/document/d/161TItfBWQYRuaPfjBNCVLfvU1LoRv6EAR85jwvXZhsg/edit?usp=sharing)
  * In-person event attendance at capacity. To attend remotely, add your name to the doc linked above.
* [Michael Rubin] Growing the sig-storage community
* [Brad Childs, 10 minutes] Summary of the F2F meeting
  * Oct 2017 Storage SIG F2F Agenda/Notes/Logistics Doc: [link](https://docs.google.com/document/d/161TItfBWQYRuaPfjBNCVLfvU1LoRv6EAR85jwvXZhsg/edit#heading=h.yw150ejajoo1)

## September 14, 2017

Recording: [https://youtu.be/AEk6mOsJEyw](https://youtu.be/AEk6mOsJEyw)

Agenda/Notes:

* [Saad Ali] Status Update
  * [Planning spreadsheet](https://docs.google.com/spreadsheets/d/1t4z5DYKjX2ZDlkTpCnp18icRAQqOE85C1T1r2gqJVck/edit?usp=sharing)
  * 1.8 Docs Planning: [https://docs.google.com/spreadsheets/d/1AFksRDgAt6BGA3OjRNIiO3IyKmA-GU7CXaxbihy48ns/edit?usp=sharing](https://docs.google.com/spreadsheets/d/1AFksRDgAt6BGA3OjRNIiO3IyKmA-GU7CXaxbihy48ns/edit?usp=sharing)
  * Bchilds to update doc list to match features.
* PRs to discuss or that need attention
  * [Please add any PRs that need attention here]
  * Detach Fix [https://github.com/kubernetes/kubernetes/pull/52221](https://github.com/kubernetes/kubernetes/pull/52221)
  * Kubelet wierd output fix [https://github.com/kubernetes/kubernetes/pull/52132](https://github.com/kubernetes/kubernetes/pull/52132)
* Design Review
  * [Please add any design reviews here]
  * 1.7 Local storage: [https://github.com/kubernetes/community/pull/989](https://github.com/kubernetes/community/pull/989)
    * Scheduler changes in 1.9
    * Possible beta in 1.10
  * Local storage provisioner block support: [https://docs.google.com/document/d/1hRRzZpWtbHyJC1lotE8rqp2h1NEXuFv-rAW-vLuQa94/view](https://docs.google.com/document/d/1hRRzZpWtbHyJC1lotE8rqp2h1NEXuFv-rAW-vLuQa94/view)
  * Volume topology scheduling: [https://github.com/kubernetes/community/pull/1054](https://github.com/kubernetes/community/pull/1054)
* [Tushar Thole] Process improvements?
  * Let’s brainstorm how the process of contributing fixes to k8s can be made more agile. Using an example, I will highlight how the existing process is slowing down bug fixing/feature development. I am sure this is an area of improvement that’s not specific to sig-storage alone.
* [Saad Ali] Next F2F meeting update
  * Dates confirmed: Oct 10-11, 2017
  * [https://docs.google.com/document/d/161TItfBWQYRuaPfjBNCVLfvU1LoRv6EAR85jwvXZhsg/edit?usp=sharing](https://docs.google.com/document/d/161TItfBWQYRuaPfjBNCVLfvU1LoRv6EAR85jwvXZhsg/edit?usp=sharing)
  * Location: Google in the Bay Area, California (most likely in Sunnyvale).
  * Will share more details around logistics, seeding topics for meeting, etc. as soon as we have a location locked down.

## August 31, 2017

Recording: [https://youtu.be/OWLEMeedi98](https://youtu.be/OWLEMeedi98)

Agenda/Notes:

* [Saad Ali] Status Update
  * [Planning spreadsheet](https://docs.google.com/spreadsheets/d/1t4z5DYKjX2ZDlkTpCnp18icRAQqOE85C1T1r2gqJVck/edit?usp=sharing)
  * [Code Freeze Sept 1](https://github.com/kubernetes/features/blob/master/release-1.8/release-1.8.md)
* PRs to discuss or that need attention
  * [Please add any PRs that need attention here]
* Design Review
  * [Please add any design reviews here]
* [Saad Ali] Next F2F meeting: location? Time?
  * Erin Boyd sent out a survey: [link](https://groups.google.com/forum/#!topic/kubernetes-sig-storage/99sCQAJD3XI)
  * Decided on Oct 10-11, 2017 at Google somewhere in the Bay Area, California (details to follow)

## August 17, 2017

Recording: [https://youtu.be/yLkolHV3uiw](https://youtu.be/yLkolHV3uiw)

Agenda/Notes:

* [Saad Ali] Status Update
  * [Planning spreadsheet](https://docs.google.com/spreadsheets/d/1t4z5DYKjX2ZDlkTpCnp18icRAQqOE85C1T1r2gqJVck/edit?usp=sharing)
* PRs to discuss or that need attention
  * [Please add any PRs that need attention here]
  * [Pablo Mercado] Digital Ocean Volume Plugin: <https://github.com/kubernetes/kubernetes/pull/50044>
    * Will take a look Flex.
* Design Review
  * [Please add any design reviews here]
  * [Jing] Local Storage Isolation Revised Design [https://docs.google.com/document/d/1ZVvsMCOhnTNbj2TzkaTfi8cbStN-ErE_DG4KGwgeeUk/](https://docs.google.com/document/d/1ZVvsMCOhnTNbj2TzkaTfi8cbStN-ErE_DG4KGwgeeUk/)
  * PVC Metrics : [https://github.com/kubernetes/community/pull/930](https://github.com/kubernetes/community/pull/930)
    * Punt to next meeting.
* [Saad Ali] Next F2F meeting: location? Time? September?
  * [Michelle] late august is right before code freeze…
  * [Steve Wong] Dell could host a meeting again if the location is Santa Clara or Austin
  * [Saad] How about September 13 and 14 in Los Angeles? To overlap with the Linux Foundation Open Source Summit North America which is Sept 11-14 in LA.
    * Might be too crowded and hard to find a place to host.
    * Could have a social gathering.
  * [Ardalan] If there is interest for meeting in late October (October 25) on the East Coast (Raleigh, NC), NetApp can host the event on the sidelines of All Things Open ([https://allthingsopen.org/](https://allthingsopen.org/)).
    * May conflict with European OSS
  * [Steve Watt] Austin + 1
  * Poll for location and times?
    * Erin will send out something.
* CFP for Kubecon closes Monday. Submit something!

## August 3, 2017

Recording: [https://youtu.be/Eh7Qa7KOL8o](https://youtu.be/Eh7Qa7KOL8o)

Agenda/Notes:

* [Saad Ali] Status Update
  * [Planning spreadsheet](https://docs.google.com/spreadsheets/d/1t4z5DYKjX2ZDlkTpCnp18icRAQqOE85C1T1r2gqJVck/edit?usp=sharing)
* PRs to discuss or that need attention
  * [Please add any PRs that need attention here]
  * Expose storage metrics: [https://github.com/kubernetes/community/pull/855](https://github.com/kubernetes/community/pull/855)
  * [Balu Dontu] VSphere cloud provider code refactoring:
    [https://github.com/kubernetes/kubernetes/pull/49164](https://github.com/kubernetes/kubernetes/pull/49164)
  * [Mitsuhiro Tanino] FC volume plugin: Support WWID for volume identifier: [https://github.com/kubernetes/kubernetes/pull/48741](https://github.com/kubernetes/kubernetes/pull/48741)
  * [Deyuan] [https://github.com/kubernetes/kubernetes/pull/49610](https://github.com/kubernetes/kubernetes/pull/49610)
  * [Vladimir Vivien] [https://github.com/kubernetes/kubernetes/pull/49973](https://github.com/kubernetes/kubernetes/pull/49973) (cherry-pick 1.7)
  * [Steve Leon] [https://github.com/kubernetes/community/pull/833](https://github.com/kubernetes/community/pull/833)
    * [https://kubernetes.io/docs/tasks/access-application-cluster/access-cluster/](https://kubernetes.io/docs/tasks/access-application-cluster/access-cluster/)
* Design Review
  * [Please add any design reviews here]
  * [Erin Boyd] Block Storage Support
    * Design Doc: [link](https://docs.google.com/a/redhat.com/document/d/1XeNFxc89C54psYqz4RErk1xcso0wMrKtSfWqI6POvaU/edit?usp=sharing)
* [Saad Ali] Next F2F meeting: location? Time?
  * [Michelle] late august is right before code freeze...

## July 20, 2017

Recording: [https://youtu.be/_oqj3JwkzVU](https://youtu.be/_oqj3JwkzVU)

Agenda/Notes:

* [Saad Ali] Status Update
  * [Planning spreadsheet](https://docs.google.com/spreadsheets/d/1t4z5DYKjX2ZDlkTpCnp18icRAQqOE85C1T1r2gqJVck/edit?usp=sharing)
  * [Michelle Au] Feature implementation tracking and coordination
* PRs to discuss or that need attention
  * [Please add any PRs that need attention here]
* Design Review
  * [Please add any design reviews here]
  * API changes for 1.8 (@jsafrane)
    * fsType in StorageCla
      * Item in StorageClass.Parameter
      * [https://github.com/kubernetes/kubernetes/pull/45345](https://github.com/kubernetes/kubernetes/pull/45345#issuecomment-313136730)
      * [Saad Ali] Ok to add as a new parameter in opaque parameter list, since this is a per-volume plugin option.
    * reclaimPolicy in StorageCla
      * Standalone field StorageClass.ReclaimPolicy
      * [https://github.com/kubernetes/kubernetes/pull/47987](https://github.com/kubernetes/kubernetes/pull/47987#discussion_r125526991)
      * [Saad Ali] Ok to add as a first class field to SC, since it is a first class filed on PV
    * [M](https://github.com/kubernetes/kubernetes/pull/47987#discussion_r125526991)o[u](https://github.com/kubernetes/kubernetes/pull/47987#discussion_r125526991)n[t](https://github.com/kubernetes/kubernetes/pull/47987#discussion_r125526991) [O](https://github.com/kubernetes/kubernetes/pull/47987#discussion_r125526991)p[t](https://github.com/kubernetes/kubernetes/pull/47987#discussion_r125526991)i[o](https://github.com/kubernetes/kubernetes/pull/47987#discussion_r125526991)n[s](https://github.com/kubernetes/kubernetes/pull/47987#discussion_r125526991) [G](https://github.com/kubernetes/kubernetes/pull/47987#discussion_r125526991)A
      * [https://github.com/kubernetes/community/pull/771](https://github.com/kubernetes/community/pull/771)
      * [Saad Ali] Let’s think about this a bit more. Might want this to be a first class field on storageClass as well.
  * [Erin Boyd] Block Storage Support
    * Design Doc: [link](https://docs.google.com/a/redhat.com/document/d/1XeNFxc89C54psYqz4RErk1xcso0wMrKtSfWqI6POvaU/edit?usp=sharing)
    * Punted to next meeting.
* [Saad Ali] Future of In-tree Volume vs Flex vs CSI
* [Jaice Singer DuMars] Release roles [still need volunteers](https://docs.google.com/spreadsheets/d/1kGPfhADIYxDR5kwS5o1kZ_cAc1A_pLqbCc1eVnXwZFo/edit#gid=0) for 1.8

## July 6, 2017

Recording: [https://youtu.be/3A0IHmuqxBs](https://youtu.be/3A0IHmuqxBs)

Agenda/Notes:

* [Saad Ali] Kubernetes v1.8 Q3 Planning - Round 2
  * [Planning spreadsheet](https://docs.google.com/spreadsheets/d/1t4z5DYKjX2ZDlkTpCnp18icRAQqOE85C1T1r2gqJVck/edit?usp=sharing)
  * Wish list (Please add any items you want NOT in the planning sheet):
    * [Requester Name] Example feature request
* PRs to discuss or that need attention
  * [Please add any PRs that need attention here]
  * Nutanix volume plugin (<https://github.com/kubernetes/kubernetes/pull/48157>)
  * Rook volume plugin: [https://github.com/kubernetes/kubernetes/pull/46843](https://github.com/kubernetes/kubernetes/pull/46843)
    * Propose using a config map instead of CRD
  * [https://github.com/kubernetes/kubernetes/pull/46821](https://github.com/kubernetes/kubernetes/pull/46821) as a segue for plugin testing, especially CI plugin testing skipped due to [Feature:Volumes] tests.
    * Michelle will finish reviewing on return next week
* Design Review
  * [Please add any design reviews here]
  * [https://github.com/kubernetes/community/pull/713](https://github.com/kubernetes/community/pull/713) related to [https://github.com/kubernetes/kubernetes/issues/47117](https://github.com/kubernetes/kubernetes/issues/47117) discussion in the last meeting
    * [Leon] Demo
    * [Brad] Where does this fit into k8s?
    * [Leon] Expose local storage on bare metal?
    * [Steve] Sounds like a volume plugin portal to external host volume discovery which has more than one storage provider at external end. Similar to CSI?
    * [Leon] CSI, create/delete/etc. This project, attaching volume to host
    * [Steve] Host local storage?
    * [Leon] This project can contain many different types of storage protocols.
    * [Hemant] I think this is trying to be an adapter -- won’t require cloudprovider -- and cinder on for example AWS.
    * [Steve] Please link URL
  * [https://docs.google.com/a/redhat.com/document/d/1XeNFxc89C54psYqz4RErk1xcso0wMrKtSfWqI6POvaU/edit?usp=sharing](https://docs.google.com/a/redhat.com/document/d/1XeNFxc89C54psYqz4RErk1xcso0wMrKtSfWqI6POvaU/edit?usp=sharing)
    * Punt to next meeting.
* [Tushar] How to improve collaboration between multiple groups on features like Snapshot
  * [Gerry] New to this community. What is the best way to hook in? Volume plugins? Third party resources?
    * Brad will talk to Gerry offline.
  * [Brad] Anyone who doesn’t feel comfortable jumping in on the group feel free to reach out to any of us offline

## June 22, 2017

Recording: [https://youtu.be/oYY-GqhrG2M](https://youtu.be/oYY-GqhrG2M)

Agenda/Notes:

* [Matt DeLio] Kubernetes v1.8 Q3 Planning
  * [Planning spreadsheet](https://docs.google.com/spreadsheets/d/1t4z5DYKjX2ZDlkTpCnp18icRAQqOE85C1T1r2gqJVck/edit?usp=sharing)
  * Wish list (Please add any items you want NOT in the planning sheet):
    * [Requester Name] Example feature request
* PRs to discuss or that need attention
  * [Please add any PRs that need attention here]
  * Rook volume plugin: [https://github.com/kubernetes/kubernetes/pull/46843](https://github.com/kubernetes/kubernetes/pull/46843)
  * [https://github.com/kubernetes/kubernetes/pull/46821](https://github.com/kubernetes/kubernetes/pull/46821) as a segue for plugin testing, especially CI plugin testing skipped due to [Feature:Volumes] tests.
* Design Review
  * [Please add any design reviews here]
  * [https://github.com/kubernetes/community/pull/713](https://github.com/kubernetes/community/pull/713) related to [https://github.com/kubernetes/kubernetes/issues/47117](https://github.com/kubernetes/kubernetes/issues/47117) discussion in the last meeting
  * <https://docs.google.com/a/redhat.com/document/d/1XeNFxc89C54psYqz4RErk1xcso0wMrKtSfWqI6POvaU/edit?usp=sharing>
* [Brad Childs] [kubernetes-incubator/external-storage](https://github.com/kubernetes-incubator/external-storage) repo updates
* [Saad Ali] CSI Implementation Discussion
* [Saad Ali] Flex Volume Issue
  * Issue [#46882](https://github.com/kubernetes/kubernetes/issues/46882) - FlexVolume unable to mount plugins that do not require an attach
    * Introduced in 1.7 alpha - Cinder bug fix made a breaking change
    * Status: Fixed
  * Issue [#44737](https://github.com/kubernetes/kubernetes/issues/44737) - Flex volumes which implement getvolumename API are getting unmounted during run time
    * Introduced 1.6 (attacher refactor)
    * Status:
      * Architectural Fix - [community/pull/650](https://github.com/kubernetes/community/pull/650)
        * Missed v1.7, punted to v1.8
      * Temp Patch Fix
        * v1.6 patched in PR [#46249](https://github.com/kubernetes/kubernetes/pull/46249) - release in v1.6.5
        * 1.7 - patched in PR [#47400](https://github.com/kubernetes/kubernetes/pull/47400) - will be released with v1.7.0
  * Issue [#47109](https://github.com/kubernetes/kubernetes/issues/47109) - “nfs flexvolume example not working”
    * Non-attacher flex drivers fail to start
    * Introduced 1.6 (attacher refactor)
    * Status:
      * Cheng Xing ([@verult](https://github.com/verult)) working on fix to be patched to 1.7 and 1.6
* [Tushar] How to improve collaboration between multiple groups on features like Snapshot

## June 8, 2017

Recording: [https://youtu.be/YEVy18n543k](https://youtu.be/YEVy18n543k)

Agenda/Notes:

* Status Update
  * [Planning spreadsheet](https://docs.google.com/spreadsheets/d/1t4z5DYKjX2ZDlkTpCnp18icRAQqOE85C1T1r2gqJVck/edit?usp=sharing)
* PRs to discuss or that need attention
  * [Please add any PRs that need attention here]
  * [https://github.com/kubernetes/kubernetes/pull/46843](https://github.com/kubernetes/kubernetes/pull/46843)
    * Proposed for 1.7, but will take more work so waiting for 1.8 to get in.  Design discussion in PR.  Will send out an invite for design meeting.  Request that Jordan liggitt joins.
  * Just noticed: [https://github.com/kubernetes/kubernetes/pull/46597](https://github.com/kubernetes/kubernetes/pull/46597) is there a plan for this for 1.7 or 1.8?
    * Need to track this and roadmap it for 1.7 or 1.8.  Childsb to talk
* Design Review
  * [Please add any design reviews here]
  * [https://github.com/kubernetes/kubernetes/issues/47117](https://github.com/kubernetes/kubernetes/issues/47117)
    * Hchen - not sure this feature represents the real problem.
    * Howard - This is a general purpose storage library that may be useful in kube and could be done as a CSI plugin.  More detailed spec/design to follow.  Demo or sep. call to discuss.
  * Local Block Storage proposal - Meeting tomorrow, design will be sent before then.
* On-boarding docs/proce
  * Want a better defined proce
  * Others looking for heavier lift PRs to help contribute and hit the amount of code needed to meet criteria
  * Can pickup “help wanted” PRs and issue
  * Reviewers and participation in issues is also important.
  * Side note: currently there are 0 open PRs with a “help-wanted” label =&gt; <https://github.com/kubernetes/kubernetes/pulls?q=is%3Aopen+is%3Apr+label%3Ahelp-wanted>
  * Storage sig community calendar?
* Testing Update (Erin)
  * Hardware donation - CNCF has donated hardware
  * Representation for plugin
  * Need community participation in testing.  Great way to contribute as a newbie.
  * Biweekley E2E testing meeting every Friday.
* CSI - CNCF meeting tomorrow.   To join the storage working group, create a PR against the CNCF storage repo which will bring you more visibility to the meetings.
  * [https://groups.google.com/forum/#!topic/kubernetes-sig-storage/sVIDnhMi4oM](https://groups.google.com/forum/#!topic/kubernetes-sig-storage/sVIDnhMi4oM)
  * Repo for spec: [https://github.com/container-storage-interface/spec](https://github.com/container-storage-interface/spec)

## May 25, 2017

Recording: [https://youtu.be/NCPCqaXG2SQ](https://youtu.be/NCPCqaXG2SQ)

Agenda/Notes:

* Status Update
  * [Planning spreadsheet](https://docs.google.com/spreadsheets/d/1t4z5DYKjX2ZDlkTpCnp18icRAQqOE85C1T1r2gqJVck/edit?usp=sharing)
* Features tracking board should be filled in for 1.7
  * [link](https://docs.google.com/spreadsheets/d/1IJSTd3MHorwUt8i492GQaKKuAFsZppauT4v1LJ91WHY/edit#gid=0)
* PRs to discuss or that need attention
  * [Please add any PRs that need attention here]
  * [https://github.com/kubernetes/kubernetes/pull/44897](https://github.com/kubernetes/kubernetes/pull/44897) (local storage plugin)
    * At a stable point. Needs final review.
    * Saad Ali will take a look
  * [https://github.com/kubernetes/kubernetes/pull/45518](https://github.com/kubernetes/kubernetes/pull/45518) (fix portworx plugin for GKE)
    * Jan reviewing looks ok.
    * Michelle Au will take a look from GKE side
* Design Review
  * [Please add any design reviews here]
  * [Hemant] <https://github.com/kubernetes/community/pull/657> - growing persistent volume
  * CSI Spec?
    * Move to next week
  * Rook Volume Plugin
    * Pretty much done. Wanted to make it for 1.7. Thin plugin
    * Would need to go through exception process.
* On-boarding docs/proce
  * Punted to next meeeting
* StorageOS Volume Plugin
  * [TODO:saadali] Some one from Google should take a look

## May 11, 2017

Recording: [https://youtu.be/RzUVcnue1j0](https://youtu.be/RzUVcnue1j0)

Agenda/Notes:

* [Erick Fejta] Flaky Test Alert
* SIG PM demo session June 13
  * Background: the PM group is looking to popularize a lot of the work each SIG is doing.  We’ve been signed up for a 15 minute time-slot; topics are any feature going to beta or stable.
  * Topics:
    * [Volume Mount Options](https://github.com/kubernetes/features/issues/168) (beta) - Huamin
    * [Cloud Provider Storage Metrics](https://github.com/kubernetes/features/issues/182) (beta) - Hemant
* Status Updates
  * [Planning spreadsheet](https://docs.google.com/spreadsheets/d/1t4z5DYKjX2ZDlkTpCnp18icRAQqOE85C1T1r2gqJVck/edit?usp=sharing)
* PRs to discuss or that need attention
  * [Please add any PRs that need attention here]
  * [https://github.com/kubernetes/kubernetes/pull/45286](https://github.com/kubernetes/kubernetes/pull/45286)
    * Should approve today
  * [https://github.com/kubernetes/kubernetes/pull/41950](https://github.com/kubernetes/kubernetes/pull/41950)
    * Rootfs has some question
  * [https://github.com/kubernetes/kubernetes/pull/45346](https://github.com/kubernetes/kubernetes/pull/45346) (need approval)
    * Saad will do a final pa
  * [https://github.com/kubernetes/kubernetes/pull/45623](https://github.com/kubernetes/kubernetes/pull/45623)
    * Bug fix in 1.6. Brad will take a look from storage side
  * [https://github.com/kubernetes/kubernetes/pull/45423](https://github.com/kubernetes/kubernetes/pull/45423)
    * Saad will take a look
  * [https://github.com/kubernetes/kubernetes/pull/45518](https://github.com/kubernetes/kubernetes/pull/45518)
    * Anthony is release czar for 1.6 he will make the final call
* Design Review
  * [Please add any design reviews here]
  * [Jan] Mount containers design [https://github.com/kubernetes/community/pull/589](https://github.com/kubernetes/community/pull/589)
    * Let’s set up a separate meeting to discuss this since we are out of time.
* On-boarding docs/proce
  * Punting to next week

## April 27, 2017

Recording: [https://youtu.be/YANFlMtue4A](https://youtu.be/YANFlMtue4A)

Agenda/Notes:

* Status Updates
  * [Planning spreadsheet](https://docs.google.com/spreadsheets/d/1t4z5DYKjX2ZDlkTpCnp18icRAQqOE85C1T1r2gqJVck/edit?usp=sharing)
  * [Matt De Lio] Discuss adding feature-repo entries for bug-fixes that impact users
  * Feature submission deadline extended to [Monday, May 5, 6:00 PM PST](https://groups.google.com/forum/#!topic/kubernetes-pm/GjFIfz_p0Js)
* PRs to discuss or that need attention
  * [Please add any PRs that need attention here]
  * Kubelet e2e test pr 44592
  * Design Review
  * [Please add any design reviews here]
  * Needs attention: [https://github.com/kubernetes/community/pull/247](https://github.com/kubernetes/community/pull/247)
* [Luis Pabón (lpabon)] “Quartermaster storage framework” demo
  * [https://github.com/coreos/quartermaster](https://github.com/coreos/quartermaster)
  * Slides: <http://bit.ly/2nsoHbW>
* On-boarding docs/proce
* OpenStack Boston - anyone going?
  * [https://www.cncf.io/event/openstack-north-america-2017/](https://www.cncf.io/event/openstack-north-america-2017/)
  * Some focus during the main event on OpenStack storage integration with K8S
  * Luis Pabon (I can attend) I also plan on attending RH Summit.
  * Chris Hoge (OpenStack Foundation, K8S community liason, chris@openstack.org)
  * Michelle Au (msau42) is going, will be giving a talk about Local Storage
  * Hayley Swimelar&lt;[hayley@linbit.com](mailto:hayley@linbit.com)&gt; is going
  * Yin ding&lt;[yin.ding@huawei.com](mailto:yin.ding@huawei.com)&gt; is going, will give a talk with Ton ([Ngoton@us.ibm.com](mailto:Ngoton@us.ibm.com)) about OpenStack &amp; K8S.
* [Jan] Mount containers design

## April 13, 2017

Recording: [https://youtu.be/UFLl9H17RdQ](https://youtu.be/UFLl9H17RdQ)

Agenda/Notes:

* Face to Face meeting was held this week
  * Notes: [https://docs.google.com/document/d/1IwL_AMO1N5aN_u5BR4lxw7g5g5z-S-I3rdkD8g81HoI/edit#](https://docs.google.com/document/d/1IwL_AMO1N5aN_u5BR4lxw7g5g5z-S-I3rdkD8g81HoI/edit#)
* Status Update
  * [Planning spreadsheet](https://docs.google.com/spreadsheets/d/1t4z5DYKjX2ZDlkTpCnp18icRAQqOE85C1T1r2gqJVck/edit?usp=sharing)
* PRs to discuss or that need attention
  * [Please add any PRs that need attention here]
  * [Jan] Shared mounts on GCE e2e test env. I need someone to review and approval, this involves sig-node, cluster-infrastructure and storage:
    * [https://groups.google.com/forum/?utm_medium=email&amp;utm_source=footer#!topic/kubernetes-dev/nzYev0cr1SU](https://groups.google.com/forum/?utm_medium=email&utm_source=footer#!topic/kubernetes-dev/nzYev0cr1SU)
    * [https://github.com/kubernetes/kubernetes/pull/44389](https://github.com/kubernetes/kubernetes/pull/44389)
    * [AI-Saad] Find a Google contact to look at this.
  * [hekumar] Fix missed pod updates - [https://github.com/kubernetes/kubernetes/pull/42033](https://github.com/kubernetes/kubernetes/pull/42033)
    * Suggestion: get more data about potential perf (memory) issue
* Design Review
  * [Please add any design reviews here]
* [Michelle Au] Persistent local storage prototype demo

## March 30, 2017

Recording: [https://youtu.be/F8HDdJwX_OQ](https://youtu.be/F8HDdJwX_OQ)

Agenda/Notes:

* Kubecon EU Day 2
* Status Update
* 2017 Q2 Storage SIG Planning
  * Storage Wish List for 2017 Q2 and 1.7
    * Out-of-tree storage e2e testing?
    * [Planning spreadsheet](https://docs.google.com/spreadsheets/d/1t4z5DYKjX2ZDlkTpCnp18icRAQqOE85C1T1r2gqJVck/edit?usp=sharing)
* Design Review
  * [Please add any design reviews here]
  * Final round: [https://github.com/kubernetes/community/pull/306](https://github.com/kubernetes/community/pull/306)
* PRs to discuss or that need attention
  * (bchilds) [https://github.com/kubernetes/kubernetes/issues/34242](https://github.com/kubernetes/kubernetes/issues/34242)
  * (rootfs) iSCSI CHAP [https://github.com/kubernetes/kubernetes/pull/43396](https://github.com/kubernetes/kubernetes/pull/43396) [(](https://github.com/kubernetes/kubernetes/pull/43396)n[e](https://github.com/kubernetes/kubernetes/pull/43396)e[d](https://github.com/kubernetes/kubernetes/pull/43396) [a](https://github.com/kubernetes/kubernetes/pull/43396)p[p](https://github.com/kubernetes/kubernetes/pull/43396)r[o](https://github.com/kubernetes/kubernetes/pull/43396)v[e](https://github.com/kubernetes/kubernetes/pull/43396)r[)](https://github.com/kubernetes/kubernetes/pull/43396)
* Face to Face meeting
  * [Steve Wong] Details, agenda proposals, and RSVP here:
    * [https://docs.google.com/document/d/1IwL_AMO1N5aN_u5BR4lxw7g5g5z-S-I3rdkD8g81HoI/edit#](https://docs.google.com/document/d/1IwL_AMO1N5aN_u5BR4lxw7g5g5z-S-I3rdkD8g81HoI/edit#)
  * 2 Days (April 11, 12) in Santa Clara
* Request: Quartermaster storage framework to Kubernetes-incubator
  * owner: Luis Pabón (lpabon)
  * [https://github.com/coreos/quartermaster](https://github.com/coreos/quartermaster) Released 3/21/2017

## March 16, 2017

Recording: [https://youtu.be/_dqtg3fURmg](https://youtu.be/_dqtg3fURmg)

Agenda/Notes:

* Status Update
  * Storage Planning Spreadsheet: [link](https://docs.google.com/spreadsheets/d/1t4z5DYKjX2ZDlkTpCnp18icRAQqOE85C1T1r2gqJVck/edit#gid=647564464)
  * K8s Release Spreadsheet: [link](https://docs.google.com/spreadsheets/d/1nspIeRVNjAQHRslHQD1-6gPv99OcYZLMezrBe3Pfhhg/edit#gid=0)
* Design Review
  * [Please add any design reviews here]
  * [Tamal Saha] [https://github.com/kubernetes/kubernetes/issues/42677](https://github.com/kubernetes/kubernetes/issues/42677)
    * Local storage solution for k8s in the work
    * In the meantime you could use hostpath plus labels/selector
    * Want StatefulSets to work with local storage.
  * [Tomas Smetana] Snapshots: next step
    * Document is not yet public
      * Jing please share
    * Folks interested in feature we should meet more frequently
    * Tomas will send message to sig-storage mailing list.
* PRs to discuss or that need attention
  * [Please add any PRs that need attention here]
  * [Jan Safranek] StorageClass v1 in 1.6: [https://github.com/kubernetes/kubernetes/pull/42477](https://github.com/kubernetes/kubernetes/pull/42477)
    * Without this everything will work, but internally we use betav1 instead of v1.
    * Let’s aim for 1.6.1
  * [Brad Childs] Attach / Detach fixes for 1.6.z? <https://github.com/kubernetes/kubernetes/issues/34242>
    * Tomas working on attach/detach controller handling restarts.
    * Too big for 1.6.0
    * Target 1.6.1
    * Tomas writing better Unit Tests.
* [Saad Ali] Container Storage Interface
  * Preliminary doc: [link](https://docs.google.com/document/d/1JMNVNP-ZHz8cGlnqckOnpJmHF-DNY7IYP-Di7iuVhQI/edit?usp=sharing)
* KubeCon EU
  * Attendees:
    * @Saad-ali (saadali@google.com
    * @Matchstick (mrubin@google.com)
    * @jsafrane (jsafrane@redhat.com)
    * @h[um](mailto:jsafrane@redhat.com)b[l](mailto:jsafrane@redhat.com)e[c](mailto:jsafrane@redhat.com) ([hchiramm@redhat.com](mailto:jsafrane@redhat.com))
    * @sergep (serge@gnode.org)
    * johscheuer
    * michael.ferranti@portworx.com
    * felix@quobyte.com
    * yin.ding@huawei.com
    * Steven.Tan@huawei.com
    * Pauline.Yeung@huawei.com
    * huangzhipeng@huawei.com
    * @clintonskitson ([clinton.kitson@dell.com](mailto:clinton.kitson@dell.com))
    * @cantbewong (steven.wong@dell.com)
  * TODO: Set up a dinner or small informal meeting
* Face to Face meeting, Santa Clara, April 5 reminder
  * [Steve Wong] Details, agenda proposals, and RSVP here:
    * [https://docs.google.com/document/d/1IwL_AMO1N5aN_u5BR4lxw7g5g5z-S-I3rdkD8g81HoI/edit#](https://docs.google.com/document/d/1IwL_AMO1N5aN_u5BR4lxw7g5g5z-S-I3rdkD8g81HoI/edit#)
  * 2 Day
  * Proposal move to April 11, 12 instead

## March 2, 2017

Recording (partial): [https://youtu.be/2163WTvqCCI](https://youtu.be/2163WTvqCCI)

Agenda/Notes:

* Status Update
  * Spreadsheet: [link](https://docs.google.com/spreadsheets/d/1t4z5DYKjX2ZDlkTpCnp18icRAQqOE85C1T1r2gqJVck/edit#gid=647564464)
* [Steve Kokhang Leon] Rook Intro
* Design Review
  * [Please add any design reviews here]
  * Flex: Is there a need for multiple provisioners per plugin?
* PRs to discuss or that need attention
  * [Please add any PRs that need attention here]
  * [Jeff Vance] E2E and out-of-tree design. When will there be a more formal discussion of the structure for the out-of-tree storage plugins? Jon and Jeff are interested in drafting a proposal on how e2e could work when we have out-of-tree, but we need more info on what out-of-tree will look like. Also, we may want a doc/proposal for e2e and flex volumes. Incubator doc: <https://github.com/kubernetes-incubator/external-storage/tree/master/docs/demo/hostpath-provisioner>  and e2e code: [https://github.com/kubernetes-incubator/external-storage/tree/master/nfs/test/e2e](https://github.com/kubernetes-incubator/external-storage/tree/master/nfs/test/e2e) . This file ([https://github.com/kubernetes-incubator/external-storage/blob/master/nfs/test/e2e/e2e_test.go](https://github.com/kubernetes-incubator/external-storage/blob/master/nfs/test/e2e/e2e_test.go)), which needs some work, is the main framework for the out-of-tree e2e tests.
  * [Simon Croome] StorageOS Volume Plugin ([https://github.com/kubernetes/kubernetes/pull/42156](https://github.com/kubernetes/kubernetes/pull/42156))
* Face to Face meeting date confirmationHas been tentatively discussed as first week in April. Dell EMC would propose to finalize as April 5&amp;6 (Wed &amp; Thur) in Santa Clara

## February 16, 2017

Recording: [https://youtu.be/33ILjxa7l-M](https://youtu.be/33ILjxa7l-M)

Agenda/Notes:

* Status Update
  * Spreadsheet: [link](https://docs.google.com/spreadsheets/d/1t4z5DYKjX2ZDlkTpCnp18icRAQqOE85C1T1r2gqJVck/edit#gid=647564464)
  * Release 1.6 Code Freeze 1 week away
* Design Review
  * [Please add any design reviews here]
  * [Steve Watt] Snapshot
    * Steve and Serge going to take up Snapshots w/Jing.
* PRs to discuss or that need attention
  * [Please add any PRs that need attention here]
  * [https://github.com/kubernetes/kubernetes/pull/38924](https://github.com/kubernetes/kubernetes/pull/38924) (Vladimir Vivien, Dell EMC)
    * Needs k8s-bot ok to test incantation
    * Looking for final LGTM for merge
  * [https://github.com/kubernetes/kubernetes/pull/39535](https://github.com/kubernetes/kubernetes/pull/39535) (Aditya Dani, Portworx)
    * Got one successful test run.
    * Needs a rebase (waiting for lgtm)
  * [https://github.com/kubernetes/kubernetes/pull/41196](https://github.com/kubernetes/kubernetes/pull/41196)
    * Any chance of backporting this PR to release 1.4 and 1.5?
    * Cherry picks?
* [Erin Boyd] Deprecating Recycler Proce
  * Follow up with Brian Grant
  * Original Recycler Dep. PR: [https://github.com/kubernetes/kubernetes/pull/36760](https://github.com/kubernetes/kubernetes/pull/36760)
* Local Storage Volume
  * <https://github.com/kubernetes/community/pull/306>
  * Timeline?
    * 1.6 Design
    * 1.7 Alpha
      * First use case will be stateful set
  * Design
    * [Steve Watt] Support for network disks in the PV framework. Data management platform folks can’t use PV framework use host path. If we can take local FS and expose them as PVs, then we can hook them up to StateFul Sets, and update scheduler to handle it, correct?
      * [msau] Correct
    * [Steve Watt] Stateful set 6 node cassandra w/6 disks, if 1 dies, we don’t have capability to dynamically provision.
      * [msau] If a node dies, want pod to release volume and use a different volume, there needs to be an available volume. Pre-provision all local volumes that are available.
    * [Guo] Node selectors and node labels, how does this differ.
      * [msau] You don’t have to decide the node ahead of time when deploying pod. It will find any available volume, and then fix node to pod. Second: idea of forgiveness, if that node is having issues, node fails or is unavailable, system can rebind to a different node a new PV.
      * [Steve Watt] Want to avoid creating a replication storm. Replace it if unavailable for a certain period of time
        * [msau] Either Completely off, no forgiveness or some time boundary)
      * [Steve Watt] Allows kubernetes to automatically handle failures vs node selectors/node labels. Huge step forward in self healing capabilities.
      * [Ardalan] If node goes down, data lost? Want to grab replica not empty scratch disk.
        * [msau] Wouldn’t other pod already have replica?
        * [Serge] depends.
        * [Guo] Portworx another replica set on different node. So ok to go to another node. If one node goes down, ok to schedule on one of the other nodes.
        * [Steve Watt] Data replication you want to maintain quorum (2 or 3) generally speaking, if one goes down. Spin up a new pod on a new node and have scratch disk replicate
        * [Serge] We are using replication term with 2 different meanings. MongoDB kind of replication is like you said (1 shard fewer, create a new shard, repopulate). Vs Replicated volumes (start with replica as a seed for the new shard). Not sure what to name them.
        * [Steve Watt] Application vs block replication?
        * [Steve Watt] Similar hadoop, multiple disks (blocks) identical copies, if host dies against one, just want scheduler to schedule against one of the replicas (not a random PV).
        * [Yaron] Active-passive replica may not be logically two PDs, just one that changes location from node A (which crashes) to node B. That makes things simpler.
        * [Steve Watt] PVs are immutable.
        * [Serge] Sounds like that is a lower layer. SDS repair would happen under the covers.
        * [Saad] If PV can move around really local storage?
        * [Steve Wong] That’s more of a hybrid local. App can take advantage on rebuild.

## February 2, 2017

Recording: [https://youtu.be/EgtyplK3Ilc](https://youtu.be/EgtyplK3Ilc)

Agenda/Notes:

* Status Update
  * 1.5.3 and 1.4.9 planned for Feb 10
* Design Review
  * [Please add any design reviews here]
  * [Michelle Au] Persistent Local Storage
    * [https://github.com/kubernetes/community/pull/306](https://github.com/kubernetes/community/pull/306)
  * [Hemant Kumar] Mount option
    * [https://github.com/kubernetes/community/pull/321](https://github.com/kubernetes/community/pull/321)
  * [Hemant Kumar] Cloud Provider metric
    * [https://github.com/kubernetes/community/pull/288](https://github.com/kubernetes/community/pull/288)
* PRs to discuss or that need attention
  * [Please add any PRs that need attention here]
  * [TODO] Assignee PR
  * [https://github.com/kubernetes/kubernetes/pull/40000](https://github.com/kubernetes/kubernetes/pull/40000)
  * [https://github.com/kubernetes/kubernetes/pull/40013](https://github.com/kubernetes/kubernetes/pull/40013)
  * [https://github.com/kubernetes/kubernetes/pull/40088](https://github.com/kubernetes/kubernetes/pull/40013)
  * [https://github.com/kubernetes/kubernetes/pull/39425](https://github.com/kubernetes/kubernetes/pull/39425)
  * [https://github.com/kubernetes/kubernetes/pull/38702](https://github.com/kubernetes/kubernetes/pull/38702)
  * [https://github.com/kubernetes/kubernetes/pull/37698](https://github.com/kubernetes/kubernetes/pull/37698)
  * [https://github.com/kubernetes/kubernetes/pull/31515](https://github.com/kubernetes/kubernetes/pull/31515)
  * [https://github.com/kubernetes/kubernetes/pull/40531](https://github.com/kubernetes/kubernetes/pull/40531)
* PoC
  * Quartermaster Storage Controller Framework
    * Luis Pabón (luis.pabon@coreos.com)  [lpabon]
    * Slides: [http://bit.ly/2jp5VB9](http://bit.ly/2jp5VB9)
    * Document: [http://bit.ly/2kikXpF](http://bit.ly/2kikXpF)
* [Erin Boyd] Deprecating Recycler Proce
  * Punt to next meeting
* [Ritesh (kerneltime)]
  * New e2e tests [https://github.com/kubernetes/kubernetes/pull/40756](https://github.com/kubernetes/kubernetes/pull/40756)
  * Storage Classes GA
    * Which issues need closing?
* KubeCon Europe
  * Who’s attending?
    * Saad Ali
    * Michael Rubin
    * Luis Pabon (maybe)
    * Christopher M Luciano

## January 19, 2017

Recording: [https://youtu.be/Yrcx2AgrJYU](https://youtu.be/Yrcx2AgrJYU)

Agenda/Notes:

* [Matt De Lio] 2017 Planning
  * Spreadsheet: [link](https://docs.google.com/spreadsheets/d/1t4z5DYKjX2ZDlkTpCnp18icRAQqOE85C1T1r2gqJVck/edit#gid=647564464)
    * We're missing reviewers for many of the features/tasks we're working on.  Please sign up as a reviewer if you are able; and for the feature owners, please start reaching out to reviewers if you're still missing one in the next few days.
    * The [feature repo](https://github.com/kubernetes/features/issues) will be frozen on January 24th, so we need to create items there that require it ASAP. Feature owners that haven't already done so should do this soon and please update the spreadsheet with issue number.
* [Saad Ali] All new 1.6 features should have:
  * Design doc
  * Tests plan doc to identify what kind of testing is required before development
  * Tests (E2E, unit, and integration) along with code, or before it (not a follow-up PR).
  * Deadlines are: [https://github.com/kubernetes/features/blob/master/release-1.6/release-1.6.md](https://github.com/kubernetes/features/blob/master/release-1.6/release-1.6.md)
    * Tuesday, Jan 24 - Features repo freeze
    * Monday, Feb 27 - Feature complete date (code freeze goes into effect)
* [Ardalan Kangarlou] NetApp’s external provisioner for Kubernete
  * [https://github.com/netapp/trident](https://github.com/netapp/trident)
* [Huamin Chen]: Out of tree provisioner repo hosting
  * Bchilds - move NFS provisioner repo to “storage”.  Will home external provisioner and FLEX plugins maintained by community.
* Design Review
  * [Please add any design reviews here]
* PRs to discuss or that need attention
  * [Please add any PRs that need attention here]
  * [https://github.com/kubernetes/kubernetes/pull/39535](https://github.com/kubernetes/kubernetes/pull/39535) (Aditya, Portworx)
  * [https://github.com/kubernetes/kubernetes/pull/38924](https://github.com/kubernetes/kubernetes/pull/38924) (Vladimir Vivien, Dell EMC)
* Next Storage F2F
  * Some time in April in Santa Clara

## January 5, 2017

Recording: [https://youtu.be/Nksb3w1a2x4](https://youtu.be/Nksb3w1a2x4)

Agenda/Notes:

* Happy New Year!
* 2017 Planning
  * Matt’s Spreadsheet: [link](https://docs.google.com/spreadsheets/d/1t4z5DYKjX2ZDlkTpCnp18icRAQqOE85C1T1r2gqJVck/edit#gid=647564464)
* Design Review
  * None
* PRs to discuss or that need attention
  * Dell EMC ScaleIO Volume Plugin PR
    * [https://github.com/kubernetes/kubernetes/pull/38924](https://github.com/kubernetes/kubernetes/pull/38924)
