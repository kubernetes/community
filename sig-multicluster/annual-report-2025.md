# 2025 Annual Report: SIG Multicluster

## Current initiatives and Project Health

1. What work did the SIG do this year that should be highlighted?

<!--
   Some example items that might be worth highlighting:
   - Major KEP advancement
   - Important initiatives that aren't tracked via KEPs
   - Paying down significant tech debt
   - Governance and leadership changes
-->

* [multicluster-runtime](https://github.com/multicluster-runtime/multicluster-runtime) accepted into the sig
* Multicluster observability user research ([proposal doc](https://docs.google.com/document/d/1uCyHpI5zV1ME9qBtGjLkrIhpl3iXe-zQsn0SSbBiAE4/edit?usp=sharing),
* "Hub"/"Management" Cluster definition discussion and position statement: https://github.com/kubernetes/community/pull/8210
* Demos:  <<UNRESOLVED:get link>>helium<</UNRESOLVED>>, [KubeStellar](https://docs.google.com/presentation/d/1PYXL7FIim6b-Hrm97utpz5zjnuwgveNWRiny47IXxLg/edit?slide=id.g35a6b02e4b0_2_0#slide=id.g35a6b02e4b0_2_0)
* Suggested well-known cluster properties and property ladder discussion: https://docs.google.com/document/d/1M6vD9ALiLLIGT3dHus-kfAUhD6wX9BPUjPB_g3XljMc/edit?tab=t.0#heading=h.jx7m99b4yg2y, <<UNRESOLVED>>PR [1](https://github.com/kubernetes/enhancements/pull/5185) and [2](https://github.com/kubernetes/enhancements/pull/5255/changes)<</UNRESOLVED>>
* MCS versions X Y Z
* [IP families](https://github.com/kubernetes/enhancements/pull/5264) and [ports conflict rules](https://github.com/kubernetes/enhancements/pull/4887#pullrequestreview-2494246716) for MCS and related conformance tests
* Traffic distribution field lift to MCS
* ClusterProfile credentials via plugin ([initial slide deck](https://docs.google.com/presentation/d/1v5-J-kFJ3TSpKqSraHcYkCz2NG7cNnYpq0ISF85wNMU/edit?slide=id.p#slide=id.p), [KEP update](https://github.com/kubernetes/enhancements/pull/5338),
* PlacementDecision API ([community doc](https://docs.google.com/document/d/1seK6W_TgSDinogXqEm8bOgFCuKqJ9_qkZdodfkSheUY/edit?tab=t.0#heading=h.wyy5e36qm908), [informational slide deck](https://drive.google.com/file/d/1b7OQotko2w6PA_U-C7KtfaAmqeuqjMu0/view?ts=682370a6),) and general discussions about multicluster scheduling ([slides](https://docs.google.com/presentation/d/1PYXL7FIim6b-Hrm97utpz5zjnuwgveNWRiny47IXxLg/edit?slide=id.g35a6b02e4b0_2_0#slide=id.g35a6b02e4b0_2_0))
* APAC meetings
* Reviewer increase
* MCS API spec/status/root discussion: [doc](https://docs.google.com/document/d/112osT8lPCg5hbbnHuekn7KVNvy2LgLPV1ICalAmb8uE/edit?tab=t.4gswme8shckl#heading=h.56zvg7miiql4), [slides](https://docs.google.com/presentation/d/12A3i8OdJdpsHu4b_IenWgldsgYSnUGkfzLJVl93b9tY/edit?usp=sharing), [github discussion thread](https://github.com/kubernetes-sigs/mcs-api/issues/106)
* About API v1beta1 CRD released: https://github.com/kubernetes-sigs/about-api/pull/27


2. Are there any areas and/or subprojects that your group needs help with (e.g. fewer than 2 active OWNERS)?


3. Did you have community-wide updates in 2025 (e.g. KubeCon talks)?

<!--
  Examples include links to email, slides, or recordings.
-->

4. KEP work in 2025 (v1.33, v1.34, v1.35):
<!--
   TODO: Uncomment the following auto-generated list of KEPs, once reviewed & updated for correction.

   Note: This list is generated from the KEP metadata in kubernetes/enhancements repository.
      If you find any discrepancy in the generated list here, please check the KEP metadata.
      Please raise an issue in kubernetes/community, if the KEP metadata is correct but the generated list is incorrect.
-->

<!-- 
  - Alpha
    - [5313 - Placement Decision API](https://github.com/kubernetes/enhancements/tree/master/keps/sig-multicluster/5313-placement-decision-api) - v1.33
    - [5339 - ClusterProfile credentials plugin](https://github.com/kubernetes/enhancements/tree/master/keps/sig-multicluster/5339-clusterprofile-plugin-credentials) - v1.34

 -->

## [Subprojects](https://git.k8s.io/community/sig-multicluster#subprojects)


**New in 2025:**
  - multicluster-runtime
**Continuing:**
  - about-api
  - cluster-inventory-api
  - mcs-api
  - sig-multicluster-site
  - work-api

## [Working groups](https://git.k8s.io/community/sig-multicluster#working-groups)

**New in 2025:**
 - AI Gateway
**Retired in 2025:**
 - Policy

## Operational

Operational tasks in [sig-governance.md]:
- [ ] [README.md] reviewed for accuracy and updated if needed
- [ ] [CONTRIBUTING.md] reviewed for accuracy and updated if needed
- [ ] Other contributing docs (e.g. in devel dir or contributor guide) reviewed for accuracy and updated if needed
- [ ] Subprojects list and linked OWNERS files in [sigs.yaml] reviewed for accuracy and updated if needed
- [ ] SIG leaders (chairs, tech leads, and subproject leads) in [sigs.yaml] are accurate and active, and updated if needed
- [ ] Meeting notes and recordings for 2025 are linked from [README.md] and updated/uploaded if needed


[CONTRIBUTING.md]: https://git.k8s.io/community/sig-multicluster/CONTRIBUTING.md
[sig-governance.md]: https://git.k8s.io/community/committee-steering/governance/sig-governance.md
[README.md]: https://git.k8s.io/community/sig-multicluster/README.md
[sigs.yaml]: https://git.k8s.io/community/sigs.yaml
[devel]: https://git.k8s.io/community/contributors/devel/README.md
