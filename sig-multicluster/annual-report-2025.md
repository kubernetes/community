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
* Multicluster observability user research ([proposal doc](https://docs.google.com/document/d/1uCyHpI5zV1ME9qBtGjLkrIhpl3iXe-zQsn0SSbBiAE4/edit?usp=sharing))
* "Hub"/"Management" Cluster definition discussion and position statement: https://github.com/kubernetes/community/pull/8210
* Demos:  [helium](https://github.com/raffaelespazzoli/helium), [KubeStellar](https://docs.google.com/presentation/d/1PYXL7FIim6b-Hrm97utpz5zjnuwgveNWRiny47IXxLg/edit?slide=id.g35a6b02e4b0_2_0#slide=id.g35a6b02e4b0_2_0), [kube-bind](https://kube-bind.io/) ([slides](https://docs.google.com/presentation/d/1a5QqiaJNpu0UgyL50T0kLGg9-gaeF-1yqW3FZThlhnU/edit?slide=id.g1692d66ae84_0_0#slide=id.g1692d66ae84_0_0), [github](https://github.com/kube-bind/kube-bind))
* MCS versions 0.2.0 and 0.3.0
  * [IP families](https://github.com/kubernetes/enhancements/pull/5264) and [ports conflict rules](https://github.com/kubernetes/enhancements/pull/4887#pullrequestreview-2494246716) for MCS and updates to related conformance tests
  * Traffic distribution/internal traffic policies field lift to MCS ([KEP update](https://github.com/kubernetes/enhancements/pull/5588), [PR](https://github.com/kubernetes-sigs/mcs-api/pull/131)) and ServiceExport conditions streamlining against Gateway API ([KEP update](https://github.com/kubernetes/enhancements/pull/5438), [PR](https://github.com/kubernetes-sigs/mcs-api/pull/112))
  * MCS API spec/status/root discussion and eventual no-op: [doc](https://docs.google.com/document/d/112osT8lPCg5hbbnHuekn7KVNvy2LgLPV1ICalAmb8uE), [slides](https://docs.google.com/presentation/d/12A3i8OdJdpsHu4b_IenWgldsgYSnUGkfzLJVl93b9tY), [github discussion thread](https://github.com/kubernetes-sigs/mcs-api/issues/106)
  * Add ServiceImport conditions ([KEP update](https://github.com/kubernetes/enhancements/pull/5439), [PR](https://github.com/kubernetes/enhancements/pull/5439)) and leverage them for asymmetrical traffic conflicts ([KEP update](https://github.com/kubernetes/enhancements/pull/5706), [PR](https://github.com/kubernetes-sigs/mcs-api/pull/132))
  * Discussion on [MCS cluster selection and traffic distribution](https://docs.google.com/document/d/14Ts1PhkSd-ouiGLhMA6uZN6RfoObqKvqpxnf7qLko5Y)
  * Additional CRD versioning data on MCS ([PR](https://github.com/kubernetes-sigs/mcs-api/pull/116))
  * Forward looking discussions on policies and plans for the MCI-API conformance report program for GA
* Cluster Inventory API
   * ClusterProfile credentials via plugin ([initial slide deck](https://docs.google.com/presentation/d/1v5-J-kFJ3TSpKqSraHcYkCz2NG7cNnYpq0ISF85wNMU), [KEP update](https://github.com/kubernetes/enhancements/pull/5338), [PR](https://github.com/kubernetes-sigs/cluster-inventory-api/pull/17))
   * PlacementDecision API added ([community doc](https://docs.google.com/document/d/1seK6W_TgSDinogXqEm8bOgFCuKqJ9_qkZdodfkSheUY), [informational slide deck](https://drive.google.com/file/d/1b7OQotko2w6PA_U-C7KtfaAmqeuqjMu0), [KEP](https://github.com/kubernetes/enhancements/pull/5314), [PR](https://github.com/kubernetes-sigs/cluster-inventory-api/pull/33)) and general discussions about multicluster scheduling ([slides](https://docs.google.com/presentation/d/1PYXL7FIim6b-Hrm97utpz5zjnuwgveNWRiny47IXxLg))
* About API
   * About API v1beta1 CRD released: https://github.com/kubernetes-sigs/about-api/pull/27
   * Discussion on suggested well-known cluster properties and property ladder discussion ([doc](https://docs.google.com/document/d/1M6vD9ALiLLIGT3dHus-kfAUhD6wX9BPUjPB_g3XljMc))
* Discussed intersections with the Gateway Inference group e.g. [Multi-Cluster Inference Gateways proposal](https://docs.google.com/document/d/1QGvG9ToaJ72vlCBdJe--hmrmLtgOV_ptJi9D58QMD2w) and [CompositeBackend proposal](https://docs.google.com/document/d/1xXM6vDDmmmwK6Oh11CdCyKKN9wdoOalnpkjsO2PdFd4)
* Ops
   * Added APAC meetings
   * Reviewer increase


2. Are there any areas and/or subprojects that your group needs help with (e.g. fewer than 2 active OWNERS)?
While we did increase the number of reviewers this year, in practice we have around the same ~6 people across the same number of projects (with very uneven distribution in terms of how much they review across all projects vs their main project) and we don't have a pipeline. Last year we tried a volunteer "council" of reviewers who would commit to 4 hours a week of reviews and also had the power to represent consensus if needed to unblock PRs in case of lack of approver/lead bandwidth, but we have not seen this specific power come into use nor the overall velocity of PR review increasing.

3. Did you have community-wide updates in 2025 (e.g. KubeCon talks)?
* KubeCon EU:
  * [How the SIG-Multicluster API Specifications Are Used for Real World Multicluster Management](https://kccnceu2025.sched.com/event/1txHL/how-the-sig-multicluster-api-specifications-are-used-for-real-world-multicluster-management-august-simonelli-red-hat-ryan-zhang-microsoft)
  * [SIG-Multicluster Intro and Deep Dive](https://kccnceu2025.sched.com/event/1td0P/sig-multicluster-intro-and-deep-dive-jeremy-olmsted-thompson-laura-lorenz-google-stephen-kitt-red-hat-ryan-zhang-microsoft)
* KubeCon NA:
  * [SIG-Multicluster Intro and Deep Dive](https://kccncna2025.sched.com/event/27Nlj/sig-multicluster-intro-and-deep-dive-stephen-kitt-red-hat-pavanipriya-sajja-independent-jeremy-olmsted-thompson-google)
  * [Finally, a Cluster Inventory I Can USE!](https://kccncna2025.sched.com/event/27FfN/finally-a-cluster-inventory-i-can-use-corentin-debains-google-ryan-zhang-microsoft)


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

 TODO: pull in the KEP work from comment

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
    - see [doc](https://docs.google.com/document/d/10WTdHYW5x2rw6BTgDzW7X-5QNesAh205MuoaUe5-IQg), [PR](https://github.com/kubernetes/community/pull/8521)

## Operational

Operational tasks in [sig-governance.md]:
- [x] [README.md] reviewed for accuracy and updated if needed
- [ ] [CONTRIBUTING.md] reviewed for accuracy and updated if needed
- [ ] Other contributing docs (e.g. in devel dir or contributor guide) reviewed for accuracy and updated if needed
- [x] Subprojects list and linked OWNERS files in [sigs.yaml] reviewed for accuracy and updated if needed
- [x] SIG leaders (chairs, tech leads, and subproject leads) in [sigs.yaml] are accurate and active, and updated if needed
- [x] Meeting notes and recordings for 2025 are linked from [README.md] and updated/uploaded if needed


[CONTRIBUTING.md]: https://git.k8s.io/community/sig-multicluster/CONTRIBUTING.md
[sig-governance.md]: https://git.k8s.io/community/committee-steering/governance/sig-governance.md
[README.md]: https://git.k8s.io/community/sig-multicluster/README.md
[sigs.yaml]: https://git.k8s.io/community/sigs.yaml
[devel]: https://git.k8s.io/community/contributors/devel/README.md
