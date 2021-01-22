# Kubernetes Storage SIG Meeting Notes (2018)

The Kubernetes Storage Special-Interest-Group (SIG) is a working group within the Kubernetes contributor community interested in storage and volume plugins. This document contains historical meeting notes from past meeting.

## December 20, 2018

Recording: [https://youtu.be/_0V2We_wRHg](https://youtu.be/_0V2We_wRHg)

Agenda/Note

* [Saad Ali] Q1 2018 v1.14 Update
  * [Planning spreadsheet](https://docs.google.com/spreadsheets/d/1t4z5DYKjX2ZDlkTpCnp18icRAQqOE85C1T1r2gqJVck/edit?usp=sharing)
  * Q1 planning
* PRs to discuss or that need attention:
  * [Please add any PRs that need attention, with your name, below]
* Design Review
  * [Please add any design reviews, with your name, below]

## December 6, 2018

Recording: [https://youtu.be/lO5tG0d_GWU](https://youtu.be/lO5tG0d_GWU)

Agenda/Note

* [Saad Ali] Q4 2018 v1.13 Update
  * [Planning spreadsheet](https://docs.google.com/spreadsheets/d/1t4z5DYKjX2ZDlkTpCnp18icRAQqOE85C1T1r2gqJVck/edit?usp=sharing)
  * Q4 end of quarter status
* PRs to discuss or that need attention:
  * [Please add any PRs that need attention, with your name, below]
  * [pontiyaraja] [#71416](https://github.com/kubernetes/kubernetes/issues/71416) [WIP] inconsistent message in different gke environment.
    * Saad and msau will take a look offline.
  * [croomes] [#69782](https://github.com/kubernetes/kubernetes/pull/69782) - StorageOS attach device bug (ready for review)
* Design Review
  * [Please add any design reviews, with your name, below]
* Question
  * How to write a CSI plugin works with both Kubernetes v1.13 and v1.12 - Sheng
* KubeCon Seattle 2018
  * F2F Meeting: [link](https://docs.google.com/document/d/1Z_jM7LWkeGqO06iaoJB40ZGETgedsG706up_EOSdQko/edit#heading=h.1hxwalotoj0l)
    * 12/12/2018 between 9AM to 12noon
    * Hosted by Diamanti
    * Open questions:
      * Official!
      * Agenda. Half day.
  * Talks:
    * Best practices for deploying stateful apps - Cloud Native Storage Day - [https://www.cloudnativestorageday.com/](https://www.cloudnativestorageday.com/)
    * Kubernetes CSI Panel -  Cloud native Storage Day [https://www.cloudnativestorageday.com/](https://www.cloudnativestorageday.com/)
    * [How Symlinks Pwned Kubernetes](https://kccna18.sched.com/event/GrZc/how-symlinks-pwned-kubernetes-and-how-we-fixed-it-michelle-au-google-jan-safranek-red-hat) (msau42 &amp; jsafrane)
    * [Extending Kubernetes or: How I Learned to Stop Worrying and Trust the Spec](https://sched.co/GrUU) - (David Zhu, Google)

## November 22, 2018

Agenda/Note

* Cancelled -- US Holiday - Thanksgiving

## November 8, 2018

Recording: [https://youtu.be/YGn6zoNFOjA](https://youtu.be/YGn6zoNFOjA)

Agenda/Note

* [Saad Ali] Q4 2018 v1.13 Update
  * [Planning spreadsheet](https://docs.google.com/spreadsheets/d/1t4z5DYKjX2ZDlkTpCnp18icRAQqOE85C1T1r2gqJVck/edit?usp=sharing)
  * ***Code Freeze - 11/15***
* PRs to discuss or that need attention:
  * [Please add any PRs that need attention, with your name, below]
  * [wongma7] block to beta [https://github.com/kubernetes/kubernetes/pull/65829](https://github.com/kubernetes/kubernetes/pull/65829)
  * [hekumar] online resize to beta [https://github.com/kubernetes/kubernetes/pull/67608](https://github.com/kubernetes/kubernetes/pull/67608)
    * Should online expansion be opt-in instead of default?
    * Drawbacks of online expansion:
      * Bug volume gets detached while mkfs, similar issue could happen with online resizing.
      * If switching to opt-in would require an API change
    * Having a switch seems reasonable.
    * Let’s hold off til 1.14 for this.
* Design Review
  * [Please add any design reviews, with your name, below]
* [Ben Swartzlander] Migration issue
  * Some vendors who created dynamic provisioners in the pre-CSI era and are now moving to CSI plugins have users who are getting trapped between worlds. We are looking for a way to help users migrate their legacy volumes to CSI volumes, and kubernetes makes this harder than it needs to be. I’d like to explore options that might make this less painful. In particular, a way to rebind a bound PVC/PV pair to a different PV.
    * Use case:
      * Old external-provisoners PVs that are in-tree volume source
      * Now want to change from e.g. intree NFS to netapp CSI driver
    * Options:
      * 1) Tool that deletes PVC and create new PVC
        * Drawback: user sees PVs being deleted.
      * 2) Tool that deletes PVs, creates new PV and rebind to existing PVC
        * Drawback: tricky to implement
      * 3) piggy back on migration.
        * Drawback: not all NFS should be redirected
    * Conclusion: 2 is probably the best option. Discuss details offline.
      * Maybe worth making it a common OSS tool.

## October 25, 2018

Recording: [https://youtu.be/ZvQn14Jq-zg](https://youtu.be/ZvQn14Jq-zg)

Agenda/Note

* [Saad Ali] Q4 2018 v1.13 Update
  * [Planning spreadsheet](https://docs.google.com/spreadsheets/d/1t4z5DYKjX2ZDlkTpCnp18icRAQqOE85C1T1r2gqJVck/edit?usp=sharing)
  * Feature (enhancement freeze) - 10/23
  * ***Code Freeze - 11/15***
* PRs to discuss or that need attention:
  * [Please add any PRs that need attention, with your name, below]
  * Iscsi-lib [https://github.com/kubernetes-csi/csi-lib-iscsi/pull/1](https://github.com/kubernetes-csi/csi-lib-iscsi/pull/1) jgriffith
* Design Review
  * [Please add any design reviews, with your name, below]
* KubeCon sig-storage Event
  * [SIG Storage Intro Session](https://kccna18.sched.com/event/GrbP/intro-storage-sig-saad-ali-google?iframe=no&w=100%&sidebar=yes&bg=no)
    * Saad and/or Brad to lead
    * QA session at end with panel?
  * Container Native Storage day (Monday 12-9-2018)
    * Call for presenters on that. May still have space.
  * SIG Storage dinner or meeting -- Brad
  * CNCF Storage Landscape
    * [CNCF storage intro](https://kccna18.sched.com/event/GrcZ/intro-cncf-storage-wg-alex-chircop-storageos-quinton-hoole-huawei?iframe=no&w=100%&sidebar=yes&bg=no)
    * [CNCF storage deep dive](https://sched.co/GreV)
  * KubeCon Storage Talk
    * [How symlinks pwned Kubernetes](https://kccna18.sched.com/event/GrZc/how-symlinks-pwned-kubernetes-and-how-we-fixed-it-michelle-au-google-jan-safranek-red-hat?iframe=no&w=100%&sidebar=yes&bg=no) - Jan + Michelle
    * [Extending Kubernetes: Or How I Learned to Stop Worrying and Trust the Spec](https://kccna18.sched.com/event/GrUU/extending-kubernetes-or-how-i-learned-to-stop-worrying-and-trust-the-spec-david-zhu-google?iframe=no&w=100%&sidebar=yes&bg=no) - David

## October 11, 2018

Recording: [https://youtu.be/6IpRsXuELZA](https://youtu.be/6IpRsXuELZA)

Agenda/Note

* [AishSundar] Update in 1.13 timeline and discuss feature load
  * [11 Enhancements in 1.13](https://github.com/kubernetes/features/issues?utf8=%E2%9C%93&q=is%3Aissue+is%3Aopen+milestone%3Av1.13+label%3Atracked%2Fyes+label%3Asig%2Fstorage)
  * Please keep the labels (kind, priority, sig) up-to-todate on the issue
  * Enhancements Freeze - 10/23 - indicate level of confidence and what’s pending in terms of code, test and docs.
  * ***Code Freeze - 11/15***
* [Saad Ali] Q4 2018 v1.13 Update
  * [Planning spreadsheet](https://docs.google.com/spreadsheets/d/1t4z5DYKjX2ZDlkTpCnp18icRAQqOE85C1T1r2gqJVck/edit?usp=sharing)
* PRs to discuss or that need attention:
  * [Please add any PRs that need attention, with your name, below]
* Design Review
  * [Please add any design reviews, with your name, below]

## September 27, 2018

Recording: [https://youtu.be/dhMNPCzZnd0](https://youtu.be/dhMNPCzZnd0)

Agenda/Notes (Please include your name and estimated time):

* [Saad Ali] Q4 2018 v1.13 Planning
  * [Planning spreadsheet](https://docs.google.com/spreadsheets/d/1t4z5DYKjX2ZDlkTpCnp18icRAQqOE85C1T1r2gqJVck/edit?usp=sharing)
* PRs to discuss or that need attention:
  * [Please add any PRs that need attention, with your name, below]
* Design Review
  * [Please add any design reviews, with your name, below]

## September 13, 2018

Recording: [https://youtu.be/2sxgltq1qiI](https://youtu.be/2sxgltq1qiI)

Agenda/Notes (Please include your name and estimated time):

* [Saad Ali] Status Update
  * [Planning spreadsheet](https://docs.google.com/spreadsheets/d/1t4z5DYKjX2ZDlkTpCnp18icRAQqOE85C1T1r2gqJVck/edit?usp=sharing)
* PRs to discuss or that need attention:
  * [Ben Swartzlander] External provisioner vendor dir and deps.
    * [https://github.com/kubernetes-csi/external-provisioner/pull/126](https://github.com/kubernetes-csi/external-provisioner/pull/126)
    * [https://github.com/kubernetes/org/issues/88](https://github.com/kubernetes/org/issues/88)
* Design Review
  * [Please add any design reviews, with your name, below]

## August 30, 2018

Recording: None

Agenda/Notes (Please include your name and estimated time):

* [Saad Ali] Status Update
  * [Planning spreadsheet](https://docs.google.com/spreadsheets/d/1t4z5DYKjX2ZDlkTpCnp18icRAQqOE85C1T1r2gqJVck/edit?usp=sharing)
* PRs to discuss or that need attention
  * [Please add any PRs that need attention, with your name, below]
  * [jinxu] add detach logic for node shutdown taint
    * [https://github.com/kubernetes/kubernetes/pull/67977/files](https://github.com/kubernetes/kubernetes/pull/67977/files)
    * Proposal: Node shutdown taint
            [https://github.com/kubernetes/kubernetes/issues/58635](https://github.com/kubernetes/kubernetes/issues/58635)[https://github.com/kubernetes/kubernetes/pull/59323](https://github.com/kubernetes/kubernetes/pull/59323)[https://github.com/kubernetes/kubernetes/pull/67254](https://github.com/kubernetes/kubernetes/pull/67254)
Meeting at 11am Pacific to discuss detach issues related to node being down: meet.google.com/qoi-pknh-bfh
* Discuss csi-connector library (Move to next / separate call)

## August 16, 2018

Recording: [https://youtu.be/SavQT4Dx6D4](https://youtu.be/SavQT4Dx6D4)

Agenda/Notes (Please include your name and estimated time):

* [Saad Ali] Status Update
  * [Planning spreadsheet](https://docs.google.com/spreadsheets/d/1t4z5DYKjX2ZDlkTpCnp18icRAQqOE85C1T1r2gqJVck/edit?usp=sharing)
* PRs to discuss or that need attention
  * [Please add any PRs that need attention, with your name, below]
  * [jsafrane] Promote mount propagation to GA
    * [https://github.com/kubernetes/kubernetes/pull/67255](https://github.com/kubernetes/kubernetes/pull/67255)
* Bugs to discuss or that need attention
  * [Please add any PRs that need attention, with your name, below]
  * [srbrahma] Enable Flex Volume Driver to provide metrics
    * <https://github.com/kubernetes/kubernetes/pull/67508>
* Design Review
  * [Please add any design reviews, with your name, below]
  * Windows support [Feature #116,](https://github.com/kubernetes/features/issues/116) [Release criteria proposed for v1.12](https://docs.google.com/document/d/1YkLZIYYLMQhxdI2esN5PuTkhQHhO0joNvnbHpW68yg8/edit) [Patrick Lang, 10 mins]
    * Need feedback, approval from impacted SIG
  * [Michelle] Conformance update
    * [https://github.com/kubernetes/kubernetes/issues/65155](https://github.com/kubernetes/kubernetes/issues/65155)
    * [https://docs.google.com/spreadsheets/d/1Yp4mHBaOx86n6CMiEweFfzIkNdSfhgzKZFTRXrnXXkM/edit#gid=0](https://docs.google.com/spreadsheets/d/1Yp4mHBaOx86n6CMiEweFfzIkNdSfhgzKZFTRXrnXXkM/edit#gid=0)
  * [Erin] Cloning design review tmr
  * [John] iSCSI has lots going on (internal target driver, built in code in-tree managing pre-provisioned iscsi volumes, need to figure out consolidation story)

## August 2, 2018

Recording: [https://youtu.be/Qu9POBVXTtI](https://youtu.be/Qu9POBVXTtI)

Agenda/Notes (Please include your name and estimated time):

* [Saad Ali] Status Update
  * [Planning spreadsheet](https://docs.google.com/spreadsheets/d/1t4z5DYKjX2ZDlkTpCnp18icRAQqOE85C1T1r2gqJVck/edit?usp=sharing)
* PRs to discuss or that need attention
  * [Please add any PRs that need attention, with your name, below]
  * [Ardalan Kangarlou] [https://github.com/kubernetes/kubernetes/pull/66780](https://github.com/kubernetes/kubernetes/pull/66780) (would like to see the PR merged in 1.12)
    * Mostly looks good. Maybe some questions about naming.
* Bugs to discuss or that need attention
  * [Please add any PRs that need attention, with your name, below]
* Design Review
  * [Please add any design reviews, with your name, below]
  * Cloning design review tomorrow. Invite sent to sig-storage.

## July 19, 2018

Recording: [https://youtu.be/WQ90iiGH7l0](https://youtu.be/WQ90iiGH7l0)

Agenda/Notes (Please include your name and estimated time):

* [Saad Ali] Status Update
  * [Planning spreadsheet](https://docs.google.com/spreadsheets/d/1t4z5DYKjX2ZDlkTpCnp18icRAQqOE85C1T1r2gqJVck/edit?usp=sharing)
* PRs to discuss or that need attention
  * [Please add any PRs that need attention, with your name, below]
  * [Ben Swartzlander] [https://github.com/kubernetes/kubernetes/pull/63176](https://github.com/kubernetes/kubernetes/pull/63176) Should have merged in 1.11
* Bugs to discuss or that need attention
  * [Please add any PRs that need attention, with your name, below]
* Design Review
  * [Please add any design reviews, with your name, below]
  * [Michelle] Conformance testing
    * Basic idea of conformance suites: [https://github.com/kubernetes/kubernetes/issues/65155](https://github.com/kubernetes/kubernetes/issues/65155)
    * Sheet with test ideas for each suite: [https://docs.google.com/spreadsheets/d/1Yp4mHBaOx86n6CMiEweFfzIkNdSfhgzKZFTRXrnXXkM/edit#gid=0](https://docs.google.com/spreadsheets/d/1Yp4mHBaOx86n6CMiEweFfzIkNdSfhgzKZFTRXrnXXkM/edit#gid=0)
  * [Jan] Volume plugin test
    * <https://github.com/kubernetes/community/pull/2272>
    * Goal: run \*existing\* tests, i.e. Ceph and iSCSI.
    * Adds new test job with MountContainers that runs all [sig-storage] tests, incl. Ceph and iSCSI.
    * No new tests.

## July 5, 2018

Agenda/Notes (Please include your name and estimated time):

* Cancelled due to US Holiday (Independence Day)

## June 21, 2018

Recording: [https://youtu.be/wWsC4zqS56c](https://youtu.be/wWsC4zqS56c)

Agenda/Notes (Please include your name and estimated time):

* [Saad Ali] Status Update
  * Q3 Kubernetes v1.12 Planning!!
    * Please add work items to the planning spreadsheet under the v1.12 tab
  * [Planning sprea](https://docs.google.com/spreadsheets/d/1t4z5DYKjX2ZDlkTpCnp18icRAQqOE85C1T1r2gqJVck/edit?usp=sharing)[heet](https://docs.google.com/spreadsheets/d/1t4z5DYKjX2ZDlkTpCnp18icRAQqOE85C1T1r2gqJVck/edit?usp=sharing)
* PRs to discuss or that need attention
  * [Please add anyds PRs that need attention, with your name, below]
* Bugs to discuss or that need attention
  * [Please add any PRs that need attention, with your name, below]
* Design Review
  * [Please add any design reviews, with your name, below]
  * [Jing and Xing] snapshot: [https://github.com/kubernetes/community/pull/1662](https://github.com/kubernetes/community/pull/1662)
    * SIG-Architecture Decision
    * SnapshotClass?
      * Decision:
        * PR XXX [Erin to add]
    * Snapshot Meeting Proposal (next Monday or Wednesday 9:00am?)
      * 9 AM Wed is CSI community meeting
  * [Jan] Volume plugin test
    * <https://github.com/kubernetes/community/pull/2272>
    * Goal: run \*existing\* tests, i.e. Ceph and iSCSI.
    * Adds new test job with MountContainers that runs all [sig-storage] tests, incl. Ceph and iSCSI.
    * No new tests.
  * [Jan] In-line CSI volume
    * <https://github.com/kubernetes/community/pull/2273>
    * Get agreement on what do we want to support.
      * Do we want full feature parity with CSI volumes as PVs?
      * Troubles with external attacher + secrets.

## June 7, 2018

Recording: [https://youtu.be/oDFjy5v_Myk](https://youtu.be/oDFjy5v_Myk)

Agenda/Notes (Please include your name and estimated time):

* [Saad Ali] Status Update
  * [Planning spreadsheet](https://docs.google.com/spreadsheets/d/1t4z5DYKjX2ZDlkTpCnp18icRAQqOE85C1T1r2gqJVck/edit?usp=sharing)
* PRs to discuss or that need attention
  * [Please add any PRs that need attention, with your name, below]
* Design Review
  * [Please add any design reviews, with your name, below]
  * [Jing and Xing] snapshot: [https://github.com/kubernetes/community/pull/1662](https://github.com/kubernetes/community/pull/1662)
    * Snapshot StorageClass?
      * Decision:
        * Let’s reuse existing StorageClass object and add a new alpha “snapshotParameters” field to it.
        * If a snapshot request does not specify a StorageClass we can use the StorageClass from the source PVC.
        * Initially, let’s also have API validation that will quickly fail a snapshot request if source and dest provisioners don’t match
    * Condition for Creating
      * Decision: ok with that.

## May 24, 2018

Recording: [https://youtu.be/2ZAGV06slo8](https://youtu.be/2ZAGV06slo8)

Agenda/Notes (Please include your name and estimated time):

* [Brad Childs] Status Update
  * ~~Code freeze is on Monday! ~~Actually it’s been pushed to June 5
  * [Planning spreadsheet](https://docs.google.com/spreadsheets/d/1t4z5DYKjX2ZDlkTpCnp18icRAQqOE85C1T1r2gqJVck/edit?usp=sharing)
* [Brad Childs] F2F Last week
  * [Notes / Recording](https://docs.google.com/document/d/1AoemMxRcRDhc3gpChGk-UwCiVsk7Q7E_yKeIKP-7jL4/edit?ts=5ad14c06)
* Design Review
  * [Please add any design reviews, with your name, below]
  * Erin Boyd - Cloning
* PRs to discuss or that need attention
  * [Ben Swartzlander] [63176](https://github.com/kubernetes/kubernetes/pull/63176) - Issue about race conditions in the iSCSI login code paths. I think the race condition exists in the existing code, and my PR doesn’t make it any worse, so we could just treat it as a separate issue. If I need to address it, the 4 approaches I can imagine are:
    * Don’t logout of iscsi sessions (let them leak)
    * Big giant lock around attach/detach (may cause throughput issues)
    * Fine grained locking (a lot of new code)
    * Retry loops (will cause error spam and unbounded attach times)

## May 10, 2018

Recording: [https://youtu.be/TxI9RYE2DwQ](https://youtu.be/TxI9RYE2DwQ)

Agenda/Notes (Please include your name and estimated time):

* [Brad Childs] Status Update
  * [Planning spreadsheet](https://docs.google.com/spreadsheets/d/1t4z5DYKjX2ZDlkTpCnp18icRAQqOE85C1T1r2gqJVck/edit?usp=sharing)
* Design Review
  * [Please add any design reviews, with your name, below]
* PRs to discuss or that need attention

* [Steve Wong] Next F2F
  * Where: Downtown Mountain View
  * When: May 15 &amp; 16
  * RSVP deadline: April 27 - a few slots left as of now
  * Register by adding name posted to Storage SIG mailing list - you must be a member to edit and attend meeting
  * Agenda doc: [link](https://docs.google.com/document/d/1AoemMxRcRDhc3gpChGk-UwCiVsk7Q7E_yKeIKP-7jL4/edit?usp=sharing_eip&ts=5ad14c06)

## April 26, 2018

Recording: [https://youtu.be/Z89r_qlR-kI](https://youtu.be/Z89r_qlR-kI)

Agenda/Notes (Please include your name and estimated time):

* [Saad Ali] Status Update
  * [Planning spreadsheet](https://docs.google.com/spreadsheets/d/1t4z5DYKjX2ZDlkTpCnp18icRAQqOE85C1T1r2gqJVck/edit?usp=sharing)
  * [***Feature freeze***](https://github.com/kubernetes/sig-release/blob/master/releases/release-1.11/release-1.11.md) ***was April 24*** -- this means that for any new features that are in v1.11 a issue must be open in the [kubernetes/features](https://github.com/kubernetes/features/issues) repo with the 1.11 milestone.
* Kubecon Next week!
  * Presentations:
    * [Efficient IoT with Protocol Buffers and gRPC](http://sched.co/DquJ) (vladimirvivien))
    * [Extending Kubernetes with gRPC](http://sched.co/Dqwb) (lightning talk, vladimirvivien)
    * [CNCF Storage working group intro](https://kccnceu18.sched.com/event/DrnW/storage-wg-intro-ben-hindman-clint-kitson-code-quinton-hoole-huawei-any-skill-level) - multiple members attending
    * [K8s SIG Storage](https://kccnceu18.sched.com/event/Drnk/sig-storage-k8s-intro-saad-ali-google-any-skill-level)
    * [CNCF Storage Working Group Deep Dive](https://kccnceu18.sched.com/event/DroG/storage-wg-deep-dive-ben-hindman-clint-kitson-code-quinton-hoole-huawei-intermediate-skill-level)
    * [K8s Storage Lingo 101](https://kccnceu18.sched.com/event/Dqvi/kubernetes-storage-lingo-101-saad-ali-google-beginner-skill-level) (saad-ali)
    * [K8s Local Storage](https://kccnceu18.sched.com/event/Dqvc/using-kubernetes-local-storage-for-scale-out-storage-services-in-production-michelle-au-google-ian-chakeres-salesforce-intermediate-skill-level) (msau42, ianchakeres)
    * [Container Storage Past, Present, Future](https://kccnceu18.sched.com/event/Dqvo/container-storage-interface-present-and-future-jie-yu-mesosphere-inc-intermediate-skill-level)
    * [Policy-Based Volume Snapshots Management in Kubernetes](http://sched.co/Dqw0) (jingxu97)
  * Who’s attending:
    * Saad Ali (saad-ali)
    * Michelle Au (msau42)
    * Jing Xu (jingxu97)
    * Jesse Brown
    * Steve Wong (cantbewong)
    * Vladimir VIvien (vladimirvivien)
    * Ardalan Kangarlou (kangarlou)
    * Kiran Mova (kmova)
    * Xing Yang (xing-yang)
    * Shubheksha Jalan (shubheksha)
    * Masaki Kimura (mkimuram)
    * Humble Chirammal (humblec)
    * Patrick Ohly (pohly)
* Design Review
  * [Please add any design reviews, with your name, below]
* PRs to discuss or that need attention
  * [pospispa won’t attend the meeting] approve/not approve StorageObjectInUseFeature become GA: [https://github.com/kubernetes/kubernetes/pull/62870](https://github.com/kubernetes/kubernetes/pull/62870)
  * [Ben Swartzlander] [https://github.com/kubernetes/kubernetes/pull/63176](https://github.com/kubernetes/kubernetes/pull/63176)
  * [Please add any PRs that need attention, with your name, below]
* [Kiran] node-disk-manager design ([https://github.com/kubernetes-incubator/external-storage/issues/736](https://github.com/kubernetes-incubator/external-storage/issues/736))
  * Sure! Next steps: Sync up with Brad.
* [Steve Wong] Next F2F
  * Where: Downtown Mountain View
  * When: May 15 &amp; 16
  * RSVP deadline: April 27 - a few slots left as of now
  * Register by adding name posted to Storage SIG mailing list - you must be a member to edit and attend meeting
  * Agenda doc: [link](https://docs.google.com/document/d/1AoemMxRcRDhc3gpChGk-UwCiVsk7Q7E_yKeIKP-7jL4/edit?usp=sharing_eip&ts=5ad14c06)
* [Erin Boyd] SIG On-boarding
  * Onboarding doc: [link](https://docs.google.com/document/d/1LHCmHTziMvzfbNVIroSF9gEZvuiBQi38gEQ6INq6eII/edit)
  * Rescheduled to 2nd week of May -- F2F conflict?
* [Erin Boyd] Restart Testing at end of May
  * Kiran: will help review E2E test

## April 12, 2018

Recording: [https://youtu.be/4gcpQax8PMQ](https://youtu.be/4gcpQax8PMQ)

Agenda/Notes (Please include your name and estimated time):

* [Jordan Liggitt] “emptyDir and subpath” CVE Post-Mortem
  * Link: [https://docs.google.com/document/d/1mJGxKyYiRigviS5pY5aEtCa-oZjHtTOTdVZE5_w4NmY](https://docs.google.com/document/d/1mJGxKyYiRigviS5pY5aEtCa-oZjHtTOTdVZE5_w4NmY)
* Design Review
  * [Please add any design reviews, with your name, below]
  * [Mike Danese] ServiceAccountTokenVolumeProjection
    * Proposal: [https://github.com/kubernetes/community/pull/1973](https://github.com/kubernetes/community/pull/1973)
  * [Michelle Au] Generate subpath names based on downward API env
    * [https://github.com/kubernetes/kubernetes/issues/48677](https://github.com/kubernetes/kubernetes/issues/48677)
* PRs to discuss or that need attention
  * [Please add any PRs that need attention, with your name, below]
  * [Ian Chakeres] “Return error in mount_unsupported for unsupported platforms” [https://github.com/kubernetes/kubernetes/pull/61914](https://github.com/kubernetes/kubernetes/pull/61914)
  * [jsafrane] Private mount propagation [https://github.com/kubernetes/kubernetes/pull/62462](https://github.com/kubernetes/kubernetes/pull/62462)
* Bugs to discuss or that need attention
  * [Please add any PRs that need attention, with your name, below]
* [Saad Ali] Status Update
  * [Planning spreadsheet](https://docs.google.com/spreadsheets/d/1t4z5DYKjX2ZDlkTpCnp18icRAQqOE85C1T1r2gqJVck/edit?usp=sharing)
  * Finalize Q2 (v1.11) planning and assignment
* [Kiran] [bchilds] sig-storage projects / proce
  * [https://docs.google.com/spreadsheets/d/1Gb88gpWMjCsROM31rXdylv0ibvCupLi-KYT_Tez5_2Y/edit?usp=sharing](https://docs.google.com/spreadsheets/d/1Gb88gpWMjCsROM31rXdylv0ibvCupLi-KYT_Tez5_2Y/edit?usp=sharing)
  * [Kiran] Process for sig-storage subprojects ?
    * [https://github.com/kubernetes/community/blob/master/incubator.md](https://github.com/kubernetes/community/blob/master/incubator.md)
    * <https://github.com/kubernetes/kubernetes/issues/58569>
    * [node-disk-manager](https://github.com/openebs/node-disk-manager) as a subproject?
    * <https://docs.google.com/presentation/d/1XcCWQL_WfhGzNjIlnL1b0kpiCvqKaUtEh9XXU2gypn4/edit#slide=id.g34dba79a54_0_0>
* Next F2F (6 minutes)
  * May 15-16 in Palo Alto area (likely Mountain View), California hosted by VMware
  * Steve Wong from VMware is finalizing arrangements to secure a venue for up to 50. Expecting to publish details and signup by end of day Friday, but plan is secure for purposes making travel arrangements now .
  * Additional sponsors for meals are welcome.
    * Volunteers orgs:
* [Erin Boyd] SIG On-boarding
  * Recurring Meeting Invite -- in sig-storage mailing list
  * Reach out Scott or Erin if you are unable to find it.
  * Please attend if you are interested in learning how to on-board.
  * Cal invite has a doc with agenda
* [Erin Boyd] Restart Testing at end of May

## March 29, 2018

Recording: [https://youtu.be/IRZyODElnlw](https://youtu.be/IRZyODElnlw)

Agenda/Notes (Please include your name and estimated time):

* [Paris Pittman, 10 min] [ContribX Roadshow](https://docs.google.com/document/d/1hESUJdDy6Q6eysNe_Wyjtq8vI5A3AnrfWCcJFP4yU6I/edit?usp=sharing)
* [Saad Ali] Status Update
  * [Planning spreadsheet](https://docs.google.com/spreadsheets/d/1t4z5DYKjX2ZDlkTpCnp18icRAQqOE85C1T1r2gqJVck/edit?usp=sharing)
  * Kick start Q2 (v1.11) planning
  * Finalize at next meeting April 12
* [bchilds] sig-storage project ownership: [https://docs.google.com/spreadsheets/d/1Gb88gpWMjCsROM31rXdylv0ibvCupLi-KYT_Tez5_2Y/edit?usp=sharing](https://docs.google.com/spreadsheets/d/1Gb88gpWMjCsROM31rXdylv0ibvCupLi-KYT_Tez5_2Y/edit?usp=sharing)
* Design Review
  * [Please add any design reviews, with your name, below]
* PRs to discuss or that need attention
  * [Please add any PRs that need attention, with your name, below]
* Bugs to discuss or that need attention
  * [Please add any PRs that need attention, with your name, below]
  * Containerized kubelet issues (msau):
    * [https://github.com/kubernetes/kubernetes/issues/61446](https://github.com/kubernetes/kubernetes/issues/61446)
    * [https://github.com/kubernetes/kubernetes/issues/61801](https://github.com/kubernetes/kubernetes/issues/61801)
    * [https://github.com/kubernetes/kubernetes/issues/61741](https://github.com/kubernetes/kubernetes/issues/61741)
  * Hostpath reconstruction (msau): [https://github.com/kubernetes/kubernetes/issues/61446](https://github.com/kubernetes/kubernetes/issues/61446)
* Next F2F
  * March 29 deadline for responding to survey: [https://www.surveymonkey.com/r/PXZDCP8](https://www.surveymonkey.com/r/PXZDCP8)
  * Plan for F2F in late May to allow host enough time to prepare.
  * Potential Topics:
    * Snapshot
      * Open questions about current external implementation
    * Ownerships of projects as we break out of external-storage
    * In-tree volume plugins to CSI
    * Home for CSI driver
  * Anyone interested in hosting F2F (last time we had 50 attendees) [***Add info below***]:
    * NetApp Research Triangle Park (Raleigh, NC), flexible dates in April or May
    * NetApp Sunnyvale, May 11th (the day after Red Hat Summit)
    * San Francisco or Mt. View Google, anytime between May 21-June 1
    * Palo Alto CA : May 8&amp;9, May 15&amp;16, May 22&amp;23 (VMware or Pivotal Labs)
  * Backup option: Kubecon EU

## March 15, 2018

Recording: [https://youtu.be/nbejWeaE8fM](https://youtu.be/nbejWeaE8fM)

Agenda/Notes (Please include your name and estimated time):

* [Saad Ali] Status Update
  * [Planning spreadsheet](https://docs.google.com/spreadsheets/d/1t4z5DYKjX2ZDlkTpCnp18icRAQqOE85C1T1r2gqJVck/edit?usp=sharing)
  * Kick start Q2 (v1.11) planning
* Design Review
  * [Please add any design reviews, with your name, below]
* PRs to discuss or that need attention
  * [Please add any PRs that need attention, with your name, below]
* Bugs to discuss or that need attention
  * [Please add any PRs that need attention, with your name, below]
  * [v1.10 release blocking bugs](https://github.com/kubernetes/kubernetes/issues?q=is%3Aopen+is%3Aissue+label%3Asig%2Fstorage+milestone%3Av1.10)
* [Paris Pittman, 10 min] ContribX Roadshow
  * Bumped to next meeting.
* Next F2F
  * March 30 deadline for responding to survey?
  * Plan for F2F in late May?
  * Potential Topics:
    * Snapshot
      * Open questions about current external implementation
    * Ownerships of projects as we break out of external-storage
    * In-tree volume plugins to CSI
    * Home for CSI driver
  * Anyone interested in hosting F2F (last time we had 50 attendees) [***Add info below***]:
    * NetApp Research Triangle Park (Raleigh, NC), flexible dates in April or May
    * NetApp Sunnyvale, May 11th (the day after Red Hat Summit)
  * Backup option: Kubecon EU
* Updates on Split [external-storage](https://github.com/kubernetes-incubator/external-storage) ?
  * Reached out to TOC -- feedback pushing back on sig specific orgs, asking for clarification.

## March 1, 2018

Recording: [https://youtu.be/Ko-JSvLR-oA](https://youtu.be/Ko-JSvLR-oA)

Agenda/Notes (Please include your name and estimated time):

* [Saad Ali] Status Update
  * [Planning spreadsheet](https://docs.google.com/spreadsheets/d/1t4z5DYKjX2ZDlkTpCnp18icRAQqOE85C1T1r2gqJVck/edit?usp=sharing)
* Design Review
  * [Please add any design reviews, with your name, below]
* PRs to discuss or that need attention
  * [Please add any PRs that need attention, with your name, below]
* [Jan] Split [external-storage](https://github.com/kubernetes-incubator/external-storage) repo into standalone repositories.
  * Sig-storage org on github?
  * Repo for the provisioning library
  * Each provisioner in its own repo?
  * Separate releases?
  * CI?
  * TODO (Saad/Brad): follow up offline on where to create org and what it should be called. Target to have it ready by next meeting.
* [Saad] Next F2F?
  * April?
  * Potential Topics:
    * Snapshot
      * Open questions about current external implementation
    * Ownerships of projects as we break out of external-storage
    * In-tree volume plugins to CSI
    * Home for CSI driver
  * Anyone interested in hosting F2F (last time we had 50 attendees) [Add info below]:
    * NetApp Research Triangle Park (Raleigh, NC), flexible dates in April or May
    * NetApp Sunnyvale, May 11th (the day after Red Hat Summit)
  * Backup option: Kubecon EU

## February 15, 2018

Recording: [https://youtu.be/7oYdq4l16-Y](https://youtu.be/7oYdq4l16-Y)

Agenda/Notes (Please include your name and estimated time):

* [Saad Ali] Status Update
  * [Planning spreadsheet](https://docs.google.com/spreadsheets/d/1t4z5DYKjX2ZDlkTpCnp18icRAQqOE85C1T1r2gqJVck/edit?usp=sharing)
  * Code freeze Feb 26.
* Design Review
  * [Please add any design reviews, with your name, below]
  * [https://github.com/kubernetes/community/pull/1700](https://github.com/kubernetes/community/pull/1700)
* PRs to discuss or that need attention
  * [Please add any PRs that need attention, with your name, below]
* [Mohamed] IBM offering space for k8s Storage SIG meetup at IBM Open Developer Community Conference (Index)
  * First day is community day.
  * In Moscone West, San Francisco (February 20)
  * [Meetup planning](https://docs.google.com/document/d/10eEQmOW_5ffrPGhUYUMe9ixmg5JtP_IU9kC5ZTmx2kY/edit)
* [Erin Boyd] Set Up a one-off “on-boarding to sig-storage” meeting
  * Created an outline of “classes” (presentations) -- plan to share with the SIG
  * 5-10 min each -- short and to the point
  * Add it to SIG storage calendar so folks can attend
  * Pre-recording vs live recording? Up to presenter.
  * March time frame
* [Saad Ali] Created an official Kubernetes Storage SIG channel: [https://www.youtube.com/channel/UCiOeuJ6L4rYNC1jwZFRmC5Q](https://www.youtube.com/channel/UCiOeuJ6L4rYNC1jwZFRmC5Q)

## February 1, 2018

Recording: [https://youtu.be/hYlS7EzUKwk](https://youtu.be/hYlS7EzUKwk)

Agenda/Notes (Please include your name and estimated time):

* [Saad Ali] Status Update
  * [Planning spreadsheet](https://docs.google.com/spreadsheets/d/1t4z5DYKjX2ZDlkTpCnp18icRAQqOE85C1T1r2gqJVck/edit?usp=sharing)
* Design Review
  * [Please add any design reviews, with your name, below]
  * Tomas Smetana (@tsmetana) RFC: Proposal to optionally disable recursive chown for pods with fsGroup in PodSecurityContext: [https://github.com/kubernetes/community/pull/1717](https://github.com/kubernetes/community/pull/1717)
* PRs to discuss or that need attention
  * [Please add any PRs that need attention, with your name, below]
* [Erin Boyd] Set Up a one-off “on-boarding to sig-storage” meeting
  * Status: Luis and Erin will sync up.
  * Suggestion to clean up the topics and make it more sig-storage centric
  * Lots of volunteers.
  * Plan to do async presentations.
  * More info in SIG mailing list from Erin

## January 18, 2018

Recording: [https://youtu.be/0NSMCavPiqQ](https://youtu.be/0NSMCavPiqQ)

Agenda/Notes (Please include your name and estimated time):

* [Saad Ali] Status Update
  * 1.10 feature freeze ([Monday, Jan 22](https://github.com/kubernetes/sig-release/blob/master/releases/release-1.10/release-1.10.md))
  * [Planning spreadsheet](https://docs.google.com/spreadsheets/d/1t4z5DYKjX2ZDlkTpCnp18icRAQqOE85C1T1r2gqJVck/edit?usp=sharing)
* Mount Propagation needs more intensive testing
  * Test it and send your feedback to [https://groups.google.com/forum/#!topic/kubernetes-sig-storage/7voscnnt3Zs](https://groups.google.com/forum/#!topic/kubernetes-sig-storage/7voscnnt3Zs)
  * Some larger scale tests are highly appreciated!
* PRs to discuss or that need attention
  * [Please add any PRs that need attention below]
  * [Jing Xu] Scalability Issue[https://github.com/kubernetes/kubernetes/issues/52128](https://github.com/kubernetes/kubernetes/issues/52128)Desired state populator takes long time to recover state for large number of volume
  * New design of volume reconstruction
[https://github.com/kubernetes/community/pull/1601](https://github.com/kubernetes/community/pull/1601) /
* [Shyam Antony &amp; Mayank Kumar (@krmayankk)] User Mode librbd implementation
* Design Review
  * [Please add any design reviews below]
  * Dynamic provisioning of block volumes [https://github.com/kubernetes/community/pull/1595](https://github.com/kubernetes/community/pull/1595)
    * TODO (saad): take a look
  * [Jing Xu] Move Snapshot API in-tree support[https://docs.google.com/document/d/1SyViEo1okCJRwrM2LtE-DM6K6G-O46tYCBNCjNaS0Lg/](https://docs.google.com/document/d/1SyViEo1okCJRwrM2LtE-DM6K6G-O46tYCBNCjNaS0Lg/)
  * postpone pv deletion if it is bound to a pvc  [https://github.com/kubernetes/community/pull/1608](https://github.com/kubernetes/community/pull/1608)
* [Brad Childs] Storage-Sig Community Statu
  * Broader community wants status more regularly as part of community meeting
  * What can we do here?
    * Focus message in reports to community meeting
    * Volunteer(s) to set up presentations for scheduled time slots in community meeting.
* [Erin Boyd, 10 minutes] Set Up a one-off “on-boarding to sig-storage” meeting
  * Leads: Erin Boyd &amp; Luis Pabon
  * Next steps?
    * Pick topic
      * Email with topics sent out ([Link](https://docs.google.com/document/d/1LHCmHTziMvzfbNVIroSF9gEZvuiBQi38gEQ6INq6eII/edit?usp=sharing))
    * Volunteers needed!!
      * Please reply to Erin’s message
    * Pre-record or live?
      * Live Meeting - Q/A is helpful

## January 4, 2018

Recording: [https://youtu.be/KEXjN33csT0](https://youtu.be/KEXjN33csT0)

Agenda/Notes (Please include your name and estimated time):

* [Luis Pabon, 10 minutes] Kubernetes 1.9 CSI alpha demo [[Slides](https://docs.google.com/presentation/d/1ErpRlih8CFx96twch3YuH6htsW4uj0adbxDJAow2L_g/edit?usp=sharing)]
* [Saad Ali, 20 minutes] 1.10 Q1 Planning and Assignment
  * [Planning spreadsheet](https://docs.google.com/spreadsheets/d/1t4z5DYKjX2ZDlkTpCnp18icRAQqOE85C1T1r2gqJVck/edit?usp=sharing) -- Please add items to 1.10 sheet and we will review and assign during meeting
  * What does it mean to go beta
    * Add API to beta group
    * Remove alpha flag gate
    * Add E2E test
    * Add user facing documentation
* CSI Kubernetes Volunteers (Add name)
  * Felipe Musse ([felipe.musse@sap.com](mailto:felipe.musse@sap.com))
  * Pietro Menna ([pietro.menna@sap.com](mailto:pietro.menna@sap.com))
  * David Zhu
  * Vladimir VIvien ([vladimir.vivien@dell.com](mailto:vladimir.vivien@dell.com))
  * Serguei Bezverkhi ([sbezverk@cisco.com](mailto:sbezverk@cisco.com))
  * Yuquan Ren ([nickren19@gmail.com](mailto:nickren19@gmail.com))
* [Saad Ali, 10 minutes] Set Up a one-off “on-boarding to sig-storage” meeting
  * Next steps?
    * Pick format
      * Zoom recorded meeting
    * Pick a date
      * Late January?
    * Pick a lead
      * Erin Boyd
      * Luis Pabon
* [Michael Rubin]
