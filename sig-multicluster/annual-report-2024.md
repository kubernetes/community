# 2024 Annual Report: SIG Multicluster

## Current initiatives and Project Health

1. What work did the SIG do this year that should be highlighted?

Stephen Kitt (@skitt) replaced Paul Morie (@pmorie) as SIG co-chair.

Significant progress was made on graduating KEP 1645 to beta; the only
remaining item is collapsing the ServiceImport spec and status fields
to root, https://github.com/kubernetes-sigs/mcs-api/pull/52 or
https://github.com/kubernetes-sigs/mcs-api/pull/85.

A conformance suite for KEP 1645 has been implemented in
https://github.com/kubernetes-sigs/mcs-api/tree/master/conformance

CI for the mcs-api repository was revamped, with presubmit jobs added
to test-infra: https://github.com/kubernetes/test-infra/pull/33393

A lot of discussion happened around ClusterProfile credentials
storage.

2. Are there any areas and/or subprojects that your group needs help with (e.g. fewer than 2 active OWNERS)?

No.

3. Did you have community-wide updates in 2024 (e.g. KubeCon talks)?

* KubeCon EU
  - SIG-Multicluster intro & deep dive:
    https://www.youtube.com/watch?v=ZD2FzR_xFdk
  - Introducing ClusterInventory and ClusterFeature API:
    https://www.youtube.com/watch?v=Xt1cuHKjKg8
* KubeCon China
  - SIG-Multicluster intro & deep dive:
    https://www.youtube.com/watch?v=7cqO8t7O7Lk
  - Developing a standard multi-cluster inventory API:
    https://www.youtube.com/watch?v=fJi6rIYoUwY
  - Connecting the dots: towards a unified multi-cluster AI/ML
    experience: https://www.youtube.com/watch?v=BnC-DrnME0E
* KubeCon NA
  - SIG-Multicluster intro & deep dive:
    https://www.youtube.com/watch?v=6nuCNCK_sdA
  - One inventory to rule them all: standardizing multicluster
    management: https://www.youtube.com/watch?v=6c8Rh_vrIA4

4. KEP work in 2024 (v1.30, v1.31, v1.32):

Our KEPs aren’t tied to Kubernetes releases. This year saw the
following work merged:

* Label sync added to KEP 1645:
  https://github.com/kubernetes/enhancements/pull/4922
* KEP 4322 renamed to “ClusterProfile API”
* Relationship between ClusterProfile, ClusterSet and cluster
  inventory clarified in KEP 4322

## [Subprojects](https://git.k8s.io/community/sig-multicluster#subprojects)

**Continuing:**
  - about-api
  - cluster-inventory-api
  - mcs-api
  - sig-multicluster-site
  - work-api

## [Working groups](https://git.k8s.io/community/sig-multicluster#working-groups)

None currently. Last year’s annual report was incorrect.

## Operational

Operational tasks in [sig-governance.md]:
- [x] [README.md] reviewed for accuracy and updated if needed
- [x] [CONTRIBUTING.md] reviewed for accuracy and updated if needed
- [x] Other contributing docs (e.g. in devel dir or contributor guide) reviewed for accuracy and updated if needed
- [x] Subprojects list and linked OWNERS files in [sigs.yaml] reviewed for accuracy and updated if needed
- [x] SIG leaders (chairs, tech leads, and subproject leads) in [sigs.yaml] are accurate and active, and updated if needed
- [ ] Meeting notes and recordings for 2024 are linked from [README.md] and updated/uploaded if needed


[CONTRIBUTING.md]: https://git.k8s.io/community/sig-multicluster/CONTRIBUTING.md
[sig-governance.md]: https://git.k8s.io/community/committee-steering/governance/sig-governance.md
[README.md]: https://git.k8s.io/community/sig-multicluster/README.md
[sigs.yaml]: https://git.k8s.io/community/sigs.yaml
[devel]: https://git.k8s.io/community/contributors/devel/README.md
