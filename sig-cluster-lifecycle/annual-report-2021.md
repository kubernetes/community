# SIG Cluster Lifecycle Annual Report 2021

## Operational

* How are you doing with operational tasks in
[sig-governance.md](https://git.k8s.io/community/committee-steering/governance/sig-governance.md)?
  * Is your README accurate? have a CONTRIBUTING.md file?
    * Yes, [README.md](https://git.k8s.io/community/sig-cluster-lifecycle/README.md) is accurate
    * We have a [CONTRIBUTING.md](https://git.k8s.io/community/sig-cluster-lifecycle/CONTRIBUTING.md)
  * All subprojects correctly mapped and listed in [sigs.yaml](https://git.k8s.io/community/sig-list.md)?
    * [Yes](https://git.k8s.io/community/sig-cluster-lifecycle/README.md#subprojects)
  * What’s your meeting culture? Large/small, active/quiet, learnings? Meeting notes up to date? Are you keeping
  recordings up to date/trends in community members watching recordings?
    * The main SIG meeting has ~10 people on average. Some projects like kubeadm have low attendance,
    while others like Cluster API have high attendance. Meeting notes usually are a best effort.
    [We record all SIG meetings](https://www.youtube.com/playlist?list=PL69nYSiGNLP29D0nYgAGWt1ZFqS9Z7lw4).
* How does the group get updates, reports, or feedback from subprojects? Are there any springing up or being
retired? Are OWNERS.md files up to date in these areas?
  * [We do regular subproject updates](https://docs.google.com/document/d/1Gmc7LyCIL_148a9Tft7pdhdee0NBHdOfHS1SAF0duI4/edit)
  as part of the main SIG call.
  * We recently retired [https://github.com/kubernetes/kube-deploy](https://github.com/kubernetes/kube-deploy) and
  [https://github.com/kubernetes-retired/kube-aws](https://github.com/kubernetes-retired/kube-aws).
  * Keeping OWNERS files up-to-date falls under maintenance of subproject leads. We hope that they keep
  them up-to-date as they know their community better. We update OWNER files and send occasional reminders for that:
    * [Recent mailing list notification](https://groups.google.com/g/kubernetes-sig-cluster-lifecycle/c/KH8g4WRjOAE)
    * [https://github.com/kubernetes/kubernetes/pull/98547](https://github.com/kubernetes/kubernetes/pull/98547)
    * [https://github.com/kubernetes/kubeadm/pull/2390](https://github.com/kubernetes/kubeadm/pull/2390)
* Same question as above but for working groups.
  * The [WG Component Standard ](https://git.k8s.io/community/wg-component-standard) which we co-own with
  SIG API Machinery [is looking for new leads](https://groups.google.com/g/kubernetes-wg-component-standard/c/sNAqDptjJug/m/P87KtaZuAgAJ).
  On private we proposed that we should ideally find new leads instead of disbanding the WG,
  but if no leads volunteer we are going to have to do that. The WG is not very active, due to the current
  leads being busy with other work.
  * We do not own other WGs.
* When was your last monthly community-wide update? (provide link to deck and/or recording)
  * Last presentation was on July 16, 2020:
    * [Slides](https://docs.google.com/presentation/d/1ZEDeF6lqxP-LmxCRa2EBmDS1sZFAv3RmrdQOUyd6IAc/edit?usp=sharing)
    * [VOD](https://www.youtube.com/watch?v=J3O8fXTm3HE)

## Membership

* Are all listed SIG leaders (chairs, tech leads, and subproject owners) active?
  * For the SIG leads yes.
  * For subproject owners, likely there are some inactive maintainers, but it’s up to the subproject to prune their OWNERs files.
* How do you measure membership? By mailing list members, OWNERs, or something else?
  * Anyone can be considered a SIG member if they join the Zoom calls or general discussions regularly.
  * In terms of repository scope - an active contributor with PRs / Issues / Reviews becomes eligible to be part of an OWNERS file.
* How does the group measure reviewer and approver bandwidth? Do you need help in any area now? What are you doing about it?
  * For sub-projects, it’s up to the active subproject maintainers to allocate review resources.
  * Some projects like etcdadm, cluster-addons and kubeadm are currently looking for more contributors.
  * We use the common methods of
  [“help-wanted” labels](https://github.com/kubernetes-sigs/cluster-api/issues?q=is%3Aissue+is%3Aopen+label%3A%22help+wanted%22)
  and announcing the request for help on slack / mailing list.
  * For the main SIG (e.g. when reviewing KEPs) the lead who has the time usually reviews, but we try to notify everyone
  so that they can comment if they have the time too.
* Is there a healthy onboarding and growth path for contributors in your SIG?
What are some activities that the group does to encourage this? What programs are you participating in to grow contributors
throughout the contributor ladder?
  * One of our methods is to just record onboarding videos / hold Zoom sessions discussing a certain problem / area:
    * [kubeadm instructional videos](https://git.k8s.io/kubeadm/CONTRIBUTING.md#instructional-videos)
    * [Cluster API e2e walkthrough](https://groups.google.com/g/kubernetes-sig-cluster-lifecycle/c/gBbbXrUThT4/m/uS3-Z2mfDAAJ)
  * New contributors interested in such an areas should join these calls and ask questions.
  * Most of our meetings have a dedicated slot for welcoming newcomers.
* What programs do you participate in for new contributors?
  * [We participated in Google Summer of code 2020](https://kubernetes.io/blog/2020/09/16/gsoc20-building-operators-for-cluster-addons/)
  with the Cluster Addons project.
* Does the group have contributors from multiple companies/affiliations? Can end users/companies contribute in some way that
they currently are not?
  * We have contributors from a number of different companies interested in Kubernetes:
    * [Cluster API stats](https://k8s.devstats.cncf.io/d/8/company-statistics-by-repository-group?orgId=1&var-period=d7&var-metric=contributions&var-repogroup_name=SIG%20Cluster%20Lifecycle%20(Cluster%20API)&var-companies=All)
    * [Overall Cluster Lifecycle stats](https://k8s.devstats.cncf.io/d/8/company-statistics-by-repository-group?orgId=1&var-period=d7&var-metric=contributions&var-repogroup_name=SIG%20Cluster%20Lifecycle&var-companies=All)
  * Everyone can contribute as long as they have the motivation and their ideas are good.

## Current initiatives and project health

* [x] Please include links to KEPs and other supporting information that will be beneficial to multiple types of community members.
* What are initiatives that should be highlighted, lauded, shout out, that your group is proud of? Currently underway?
What are some of the longer tail projects that your group is working on?
  * We did a KEP to standardize how clusters define insecure local container registries:
    * [1755-communicating-a-local-registry](https://git.k8s.io/enhancements/keps/sig-cluster-lifecycle/generic/1755-communicating-a-local-registry)
  * Cluster API has its own spin of the KEP process and has a number of interesting active proposals:
    * [Cluster API proposals](https://github.com/kubernetes-sigs/cluster-api/tree/master/docs/proposals)
  * Most of our subprojects have existed for a long time - including, kubeadm, kops, minikube, etc.
* Year to date KEP work review: What’s now stable? Beta? Alpha? Road to alpha?
  * Only kubeadm features go through the Alpha->GA stages of the Kubernetes release process:
    * [keps/sig-cluster-lifecycle](https://https://git.k8s.io/enhancements/keps/sig-cluster-lifecycle)
  * Projects like kops, minikube, kubeadm, kubespray are mostly stable.
  * Projects like Cluster API and etcdadm are alpha and moving towards graduation following a process similar to Kubernetes.
* What areas and/or subprojects does the group need the most help with?
  * Here are a few picks from more to less:
    * [Etcdadm](https://github.com/kubernetes-sigs/etcdadm)
    * [Cluster-addons](https://github.com/kubernetes-sigs/cluster-addons)
    * [Kubeadm](https://github.com/kubernetes/kubeadm)
* What's the average open days of a PR and Issue in your group? / what metrics does your group care about and/or measure?
  * Our SIG is quite big and it depends on the subprojects.
  * It’s up to the subproject to monitor the metrics they care about.
  * For the SIG as a whole we track what tickets we have here:
    * [kubernetes/enhancements](https://github.com/kubernetes/enhancements/labels/sig%2Fcluster-lifecycle)
    * [kubernetes/kubernetes](https://github.com/kubernetes/kubernetes/labels/sig%2Fcluster-lifecycle)
