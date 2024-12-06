# 2023 Annual Report: SIG Multicluster

## Current initiatives and Project Health

1. What work did the SIG do this year that should be highlighted?

<!--
   Some example items that might be worth highlighting:
   - Major KEP advancement
   - Important initiatives that aren't tracked via KEPs
   - Paying down significant tech debt
   - Governance and leadership changes
-->

* Add outreach website at multicluster.k8s.io (source is at
  https://github.com/kubernetes-sigs/sig-multicluster-site)
* MCS API
  * Location disambiguation in multicluster DNS for headless service pods
  * Redefine EndpointSlice behavior in MCS API
* About API KEP in beta
* ClusterInventory API active discussion and provisional

2. Are there any areas and/or subprojects that your group needs help with (e.g.
   fewer than 2 active OWNERS)?

<!--
   Note: This list is generated from the KEP metadata in kubernetes/enhancements repository.
      If you find any discrepancy in the generated list here, please check the KEP metadata.
      Please raise an issue in kubernetes/community, if the KEP metadata is correct but the generated list is incorrect.
-->

Work API? SIG-Multicluster site?

3. Did you have community-wide updates in 2023 (e.g. KubeCon talks)?

<!--
  Examples include links to email, slides, or recordings.
-->
* Two Houses, Both Alike in Dignity: Gateway API and MCS API (panel)
  https://sched.co/1Hydh
* SIG-Multicluster Intro and Deep Dive (maintainer talk) https://sched.co/1HyTI  

4. KEP work in 2023 (v1.27, v1.28, v1.29):

  - Provisional
    - [4322 - Inventory Cluster
      API](https://github.com/kubernetes/enhancements/tree/master/keps/sig-multicluster/4322-cluster-inventory)
      - v1.28

  - Beta
    - [2149 - ClusterID for ClusterSet
      Identification](https://github.com/kubernetes/enhancements/tree/master/keps/sig-multicluster/2149-clusterid)
      - v1.28


## [Subprojects](https://git.k8s.io/community/sig-multicluster#subprojects)


**New in 2023:**
  - [cluster-inventory-api](https://git.k8s.io/community/<no
    value>#cluster-inventory-api)
  - [sig-multicluster-site](https://git.k8s.io/community/<no
value>#sig-multicluster-site) **Retired in 2023:**
  - Kubefed **Continuing:**
  - about-api
  - mcs-api
  - work-api

## [Working groups](https://git.k8s.io/community/sig-multicluster#working-groups)

**Continuing:**
 - IoT Edge
 - Policy

## Operational

Operational tasks in [sig-governance.md]:
- [x] [README.md] reviewed for accuracy and updated if needed
- [x] [CONTRIBUTING.md] reviewed for accuracy and updated if needed
- [ ] Other contributing docs (e.g. in devel dir or contributor guide) reviewed
  for accuracy and updated if needed
- [x] Subprojects list and linked OWNERS files in [sigs.yaml] reviewed for
  accuracy and updated if needed
- [ ] SIG leaders (chairs, tech leads, and subproject leads) in [sigs.yaml] are
  accurate and active, and updated if needed
- [ ] Meeting notes and recordings for 2023 are linked from [README.md] and
  updated/uploaded if needed


[CONTRIBUTING.md]: https://git.k8s.io/community/sig-multicluster/CONTRIBUTING.md
[sig-governance.md]:
    https://git.k8s.io/community/committee-steering/governance/sig-governance.md
[README.md]: https://git.k8s.io/community/sig-multicluster/README.md
[sigs.yaml]: https://git.k8s.io/community/sigs.yaml
[devel]: https://git.k8s.io/community/contributors/devel/README.md
