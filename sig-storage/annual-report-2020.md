# Kubernetes SIG Storage Community Group Annual Reports 2020

This report reflects back on CY 2020 and was written in Feb-Mar. 2021.

## Operational

* How are you doing with operational tasks in sig-governance.md?
  * Is your README accurate? have a CONTRIBUTING.md file?
    * Yes, our README is accurate: https://github.com/kubernetes/community/blob/master/sig-storage/README.md
    * Our CONTRIBUTING.md file is up to date: https://github.com/kubernetes/community/blob/master/sig-storage/CONTRIBUTING.md
    * We also have a section in the following doc for new contributors that is up to date: https://github.com/kubernetes/community/blob/master/contributors/devel/README.md#sig-storage

  * All subprojects correctly mapped and listed in sigs.yaml?
    * Yes

  * What’s your meeting culture? Large/small, active/quiet, learnings? Meeting notes up to date? Are you keeping recordings up to date/trends in community members watching recordings?
    * SIG-Storage has bi-weekly meetings where we do project tracking followed up discussions on PRs, issues if any. Meetings are fairly well attended. Meetings are recorded and meeting notes are up to date.
https://github.com/kubernetes/community/blob/master/sig-storage/README.md#meetings
    * We also have sub-projects meetings such as CSI implementation meetings that happen twice a week and COSI weekly standup and weekly design meetings. There are also regular meetings for Data Protection WG, CSI migration, CSI windows, volume populator, etc. as well as one-off design meetings on specific topics/features.
    * https://calendar.google.com/calendar/u/0/embed?src=vvvo48r6cprccii1lsava6p2uc@group.calendar.google.com
    * We would love to see more contributors, new and existing, to join the meetings and participate in SIG projects.

  * How does the group get updates, reports, or feedback from subprojects? Are there any springing up or being retired? Are OWNERS.md files up to date in these areas?
    * Features worked by sub-projects are tracked by SIG-Storage and will be reported regularly in SIG-Storage bi-weekly meetings. SIG leads also participate in sub-projects meetings, i.e., meetings for CSI and COSI.
https://docs.google.com/spreadsheets/d/1t4z5DYKjX2ZDlkTpCnp18icRAQqOE85C1T1r2gqJVck/edit#gid=2027057294
    * There are also project repos that are deprecated. For example, https://github.com/kubernetes/external-storage is deprecated and moved to https://github.com/kubernetes-retired/external-storage. Some of the sub-projects in the original repo were moved out and became independent repos and are still managed by sig-storage with new maintainers. Some of the sub-projects were moved out because it is getting harder and harder to maintain such a big and complicated project. Moving them to independent repos allow some contributors to take new responsibilities and become maintainers of these projects. For example, the following 2 nfs projects were moved out of external-storage and became independent projects with new maintainers:
      * https://github.com/kubernetes-sigs/nfs-subdir-external-provisioner
      * https://github.com/kubernetes-sigs/nfs-ganesha-server-and-external-provisioner

  * Same question as above but for working groups.
    * Features worked by WGs are tracked by SIG-Storage and reported at SIG-Storage meetings, i.e., ContainerNotifier and CBT are worked on by Data Protection WG and are tracked in the SIG-Storage feature tracking sheet.
https://docs.google.com/spreadsheets/d/1t4z5DYKjX2ZDlkTpCnp18icRAQqOE85C1T1r2gqJVck/edit#gid=2027057294

  * When was your last monthly community-wide update? (provide link to deck and/or recording)
    * Our last monthly update was 10/15/2020. Here’s the deck: https://docs.google.com/presentation/d/1uzS6Q1OmttV-0hzlu_Eublx24nRmWiJQ__J65Sr3hIs/edit#slide=id.g401c104a3c_0_0

## Membership

  * Are all listed SIG leaders (chairs, tech leads, and subproject owners) active?
    * Yes

  * How do you measure membership? By mailing list members, OWNERs, or something else?
    * Mailing list members, participants of SIG meetings, and folks working on SIG-Storage projects.

  * How does the group measure reviewer and approver bandwidth? Do you need help in any area now? What are you doing about it?
    * At the SIG-Storage bi-weekly meetings, we track the progress of each feature the SIG is working on. Each feature has dev leads and reviewers assigned to it. If a developer/reviewer does not have time to do the design/development/reviews, we’ll try to find someone else who has the bandwidth.
    * For features that are GA, we still continue to maintain and fix bugs and make enhancements. For example, dynamic provisioning has been a GA feature for quite a while now. We still fix bugs and continue to enhance it if needed. Another example is volume snapshot which just moved to GA in 1.20. We are still working on small enhancements, adding tests, fixing bugs, etc. The work is still being tracked in weekly standup meetings.
    * We also have this new sub-project COSI which has been very active. It has a weekly standup meeting and a weekly design meeting. A few contributors have been actively working on design and development there.

  * Is there a healthy onboarding and growth path for contributors in your SIG? What are some activities that the group does to encourage this? What programs are you participating in to grow contributors throughout the contributor ladder?
    * For each feature the SIG is working on, the SIG lead or dev lead will be looking for contributors by adding the “Help wanted” tag on an issue and/or asking for helpers in SIG-Storage bi-weekly meetings. This is a great opportunity for new contributors to step up to contribute and become members of kubernetes and eventually owners of a project.
    * In the beginning of each kubernetes release, we have a planning meeting to go over all features targeted for that release. We encourage everyone to add features they want to work on in the planning spreadsheet. This is a good time for new contributors to learn our planning process, and volunteer to be involved in some of the feature development work.

  * What programs do you participate in for new contributors?
    * At the Contributor Summit at KubeCon, we usually have a Meet & Greet for new contributors. 
    * At the bi-weekly SIG-Storage meetings, we also call out to any new contributors who are willing to help on some projects that we are looking for developers or reviewers.

  * Does the group have contributors from multiple companies/affiliations? Can end users/companies contribute in some way that they currently are not?
    * Yes, we do have contributors from multiple companies/affiliations as shown here: https://k8s.devstats.cncf.io/d/8/company-statistics-by-repository-group?orgId=1&var-period=d7&var-metric=contributions&var-repogroup_name=SIG%20Storage&var-companies=All
    * End users are welcome to open issues and fix bugs that they discovered. They are also welcome to contribute code just like other developers. Suggestions are welcome to improve in this area. Ping one of the sig leads on slack, send an email to the SIG-Storage mailing list, or attend one of the SIG-Storage meetings to provide suggestions.

## Current initiatives and project health
[ ] Please include links to KEPs and other supporting information that will be beneficial to multiple types of community members. 

What are initiatives that should be highlighted, lauded, shout out, that your group is proud of? Currently underway? What are some of the longer tail projects that your group is working on?
* Container Object Storage Interface
  * Building on the success of Container Storage Interface (CSI) which focused on Block and File, k8s SIG Storage is looking to build a similar interface for object Storage.
  * KEP: https://github.com/kubernetes/enhancements/tree/master/keps/sig-storage/1979-object-storage-support
  * Status: POC in progress
* Generic ephemeral volumes
  * Ephemeral volume types are created at pod creation time and deleted at pod termination time, like EmptyDir, SecretVolumes, ConfigMapVolumes. Third-parties would like to be able to create their own ephemeral volume types. Kubernetes permits creation of new ephemeral volume types with CSI, but that requires creating a custom CSI driver dedicated to the new ephemeral volume type. Generic ephemeral volumes will permit any existing persistent CSI Driver to be used as a ephemeral volume (so for example, users can have a new external volume provisioned and used like an EmptyDir, and deleted after pod termination like EmptyDir).
  * KEP: https://github.com/kubernetes/enhancements/pull/1701
  * Status: Alpha in 1.19
* CSI Support for Windows
  * The new CSI Proxy for Windows enables CSI Drivers to work on Windows overcoming a Windows limitation that prevented containerized CSI Drivers.
  * Blog: https://kubernetes.io/blog/2020/04/03/kubernetes-1-18-feature-windows-csi-support-alpha/
  * Status: Beta in 1.19
* Volume snapshots GA:
  * Volume Snapshot feature provides standardized Kubernetes APIs and CSI spec to support snapshot and restore functionality for CSI volume drivers.
  * Blog: https://kubernetes.io/blog/2020/12/10/kubernetes-1.20-volume-snapshot-moves-to-ga/
  * KEP: https://github.com/kubernetes/enhancements/tree/master/keps/sig-storage/177-volume-snapshot
  * Status: GA in 1.20

Year to date KEP work review: What’s now stable? Beta? Alpha? Road to alpha?
* Stable:
  * Volume snapshots GA 1.20: https://github.com/kubernetes/enhancements/tree/master/keps/sig-storage/177-volume-snapshot
https://kubernetes.io/blog/2020/12/10/kubernetes-1.20-volume-snapshot-moves-to-ga/
  * Raw block GA 1.18: https://github.com/kubernetes/enhancements/tree/master/keps/sig-storage/351-raw-block-support
  * CSI raw block GA 1.18: https://github.com/kubernetes/enhancements/tree/master/keps/sig-storage/565-csi-block-support
  * CSI Cloning GA 1.18: https://github.com/kubernetes/enhancements/tree/master/keps/sig-storage/989-extend-datasource
  * CSI Skip attach GA 1.18: https://github.com/kubernetes/enhancements/tree/master/keps/sig-storage/770-csi-skip-attach
  * CSI Pod info on mount GA 1.18 (CSIDriver): https://github.com/kubernetes/enhancements/tree/master/keps/sig-storage/603-csi-pod-info

* Beta:
  * CSI Windows beta 1.19
  * Immutable secrets and configmaps beta 1.19: https://github.com/kubernetes/enhancements/blob/master/keps/sig-storage/1412-immutable-secrets-and-configmaps
  * Non-recursive volume ownership (fsgroup) beta 1.20: https://github.com/kubernetes/enhancements/blob/master/keps/sig-storage/695-skip-permission-change/kep.yaml
  * CSIDriver policy for fsgroup beta 1.20
    * https://github.com/kubernetes/enhancements/blob/master/keps/sig-storage/1682-csi-driver-skip-permission/kep.yaml
    * https://kubernetes.io/blog/2020/12/14/kubernetes-release-1.20-fsgroupchangepolicy-fsgrouppolicy/
  * CSI Migration beta 1.17: https://github.com/kubernetes/enhancements/tree/master/keps/sig-storage/625-csi-migration
    * Azure Disk and vSphere CSI migration beta in 1.19
    * OpenStack CSI migration beta in 1.18
    * GCE CSI migration beta 1.17
    * AWS CSI migration beta 1.17
    * Azure File alpha 1.15
  * CSI Volume expansion: beta in 1.16
    * https://github.com/kubernetes/enhancements/issues/556
    * https://github.com/kubernetes/enhancements/tree/master/keps/sig-storage/284-enable-volume-expansion
  * CSI inline volume beta 1.16: https://github.com/kubernetes/enhancements/tree/master/keps/sig-storage/596-csi-inline-volumes

* Alpha:
  * Volume populator alpha 1.18: https://github.com/kubernetes/enhancements/blob/master/keps/sig-storage/1495-generic-data-populators/kep.yaml
  * Storage capacity tracking alpha 1.19: https://github.com/kubernetes/enhancements/blob/master/keps/sig-storage/1472-storage-capacity-tracking/kep.yaml
  * Generic ephemeral volume alpha 1.19: https://github.com/kubernetes/enhancements/blob/master/keps/sig-storage/1698-generic-ephemeral-volumes/kep.yaml
  * Volume health alpha 1.19: https://github.com/kubernetes/enhancements/blob/master/keps/sig-storage/1432-volume-health-monitor/kep.yaml
  * Service account token alpha 1.20: https://github.com/kubernetes/enhancements/blob/master/keps/sig-storage/1855-csi-driver-service-account-token/kep.yaml
    * https://kubernetes.io/blog/2020/12/18/kubernetes-1.20-pod-impersonation-short-lived-volumes-in-csi/

* Road to alpha:
  * COSI: https://github.com/kubernetes/enhancements/tree/master/keps/sig-storage/1979-object-storage-support
  * Non-graceful node shutdown: https://github.com/kubernetes/enhancements/pull/1116
  * Allow Kubernetes to supply pod's fsgroup to CSI driver on mount: https://github.com/kubernetes/enhancements/pull/2323
  * Volume group: https://github.com/kubernetes/enhancements/pull/1551
  * Recover from resize failures: https://github.com/kubernetes/enhancements/tree/master/keps/sig-storage/1790-recover-resize-failure
  * SELinux recursive permission handling: https://github.com/kubernetes/enhancements/issues/1710
  * Volume/Snapshot namespace transfer: https://github.com/kubernetes/enhancements/pull/2326
  * Volume health (2nd alpha due to re-design): https://github.com/kubernetes/enhancements/pull/2286
  * Prioritization on volume capacity (sig-scheduling): https://github.com/kubernetes/enhancements/pull/1862
  * PVC created by statefulset will not be auto removed (sig-apps): https://github.com/kubernetes/enhancements/tree/master/keps/sig-apps/1847-autoremove-statefulset-pvcs
  * Volume expansion for Stateful sets (sig-apps): https://github.com/kubernetes/enhancements/pull/1848
  * Enable User namespaces in kubelet so UIDs get shifted (sig-node): https://github.com/kubernetes/enhancements/pull/2101
ContainerNotifier (sig-node): https://github.com/kubernetes/enhancements/pull/1995
  * User id ownership in configmaps and secrets (sig-auth)

What areas and/or subprojects does the group need the most help with?
* There are features that we co-owned with sig-node, sig-apps, etc. We would need help from other SIGs to move things forward.
  * Features co-owned with sig-node:
    * ContainerNotifier: https://github.com/kubernetes/enhancements/pull/1995
    * Volume health: https://github.com/kubernetes/enhancements/pull/2286
  * Features co-owned with sig-apps:
    * PVC created by statefulset will not be auto removed: https://github.com/kubernetes/enhancements/tree/master/keps/sig-apps/1847-autoremove-statefulset-pvcs
    * Volume expansion for statefulset: https://github.com/kubernetes/enhancements/pull/1848
  * Features co-owned with sig-scheduling:
    * Prioritization on volume capacity (sig-scheduling): https://github.com/kubernetes/enhancements/pull/1862

* There are also areas where we need sig-architecture’s help. For example, VolumeSnapshot is already GA and it is a core feature in sig-storage developed using CRDs. We have documented and blogged that components such as the snapshot CRDs and snapshot controller need to be deployed by Kubernetes distros, however, not every distro is installing them. This is a problem with any feature developed using CRD. We hope sig-architecture can help us resolve this problem.

What's the average open days of a PR and Issue in your group? / what metrics does your group care about and/or measure?

According to the following stats on 1/25/2021, it shows the 7 day MA for a PR Time to Approve and Merge in sig-storage repository group is:
* Open to lgtm: avg 3.55 days, max 14.11 week
* Lgtm to approve: avg 12.60 hour, max 3.92 week
* Approve to merge: avg 1.03 hour, max 1.14 day
* 85% open to lgtm: avg 1.4 week, max 14.11 week
* 85% lgtm to approve: avg 17.11 hour, max 3.92 week
* 85% approve to merge: avg 1.2 day, max 1.86 week
* https://k8s.devstats.cncf.io/d/44/pr-time-to-approve-and-merge?orgId=1&var-period=d7&var-repogroup_name=SIG%20Storage&var-apichange=All&var-size_name=All&var-kind_name=All

Age of 7 day MA of issues by sig-storage repository group on 1/25/2021:
* Median time to close issue: Min 0.93 hour, Max 23.55 week, Avg 3.30 week
* https://k8s.devstats.cncf.io/d/15/issues-age-by-sig-and-repository-groups?orgId=1&var-period=d7&var-repogroup_name=SIG%20Storage&var-sig_name=All&var-kind_name=All&var-prio_name=All

Suggestions on how to improve the PR/Issue velocity:
* Increase the reviewer pool
* Regular issue triage meetings
