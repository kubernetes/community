# 2024 Annual Report: SIG API Machinery

## Current initiatives and Project Health

1. What work did the SIG do this year that should be highlighted?
- Officially added Stefan Schimanski (@sttts) as the third tech lead of the SIG.
- Carefully reviewed on each release the list of open KEPs and which ones were going to make it in the release.
- 2024 Contributor Awards:
  - Marek Siarkowicz, @serathius
  - Lukasz Szaszkiewicz, @polynomial
- Some KEPs that deserve special mention in 2024 are
  - Graduated to *Stable*
    - [CEL for Admission Control](https://github.com/kubernetes/enhancements/tree/master/keps/sig-api-machinery/3488-cel-admission-control#summary)
    - [Aggregated Discovery](https://github.com/kubernetes/enhancements/tree/master/keps/sig-api-machinery/3352-aggregated-discovery#summary)
    - [Unknown Version Interoperability Proxy](https://github.com/kubernetes/enhancements/tree/master/keps/sig-api-machinery/4020-unknown-version-interoperability-proxy#summary)
    - [Transition from SPDY to Websockets](https://github.com/kubernetes/enhancements/tree/master/keps/sig-api-machinery/4006-transition-spdy-to-websockets#summary)
    - [Custom Resource Field Selectors](https://github.com/kubernetes/enhancements/tree/master/keps/sig-api-machinery/4358-custom-resource-field-selectors#summary)
  - Graduated to *Beta*
    - [Consistent Reads from Cache](https://github.com/kubernetes/enhancements/tree/master/keps/sig-api-machinery/2340-Consistent-reads-from-cache#summary)
    - [Allow informers for getting a stream of data instead of chunking](https://github.com/kubernetes/enhancements/tree/master/keps/sig-api-machinery/3157-watch-list#summary)

2. Are there any areas and/or subprojects that your group needs help with (e.g. fewer than 2 active OWNERS)?
- API Machinery area is very extense, here are some areas that could use extra help:
  - Garbage Collection
  - Resource Quota
  - Server Side Apply

3. Did you have community-wide updates in 2024 (e.g. KubeCon talks)?

- [Spotlight on SIG API Machinery](https://www.kubernetes.dev/blog/2024/08/07/sig-api-machinery-spotlight-2024/)
- Kubecon EU 2024 [SIG API Machinery Maintainers (Two Tracks) - Abu Kashem, Red Hat & Mike Spreitzer, IBM](https://www.youtube.com/watch?v=YpQxxZ1Izek&ab_channel=CNCF%5BCloudNativeComputingFoundation%5D)
- Kubecon NA 2024 [Squashing Trampoline Pods: The Future of Securely Enabling Hardware Extensions- Joe Betz, David Eads](https://www.youtube.com/watch?v=qRo1Qw_Hr2A&ab_channel=CNCF%5BCloudNativeComputingFoundation%5D)


4. KEP work in 2024 (v1.30, v1.31, v1.32):
<!--
   TODO: Uncomment the following auto-generated list of KEPs, once reviewed & updated for correction.

   Note: This list is generated from the KEP metadata in kubernetes/enhancements repository.
      If you find any discrepancy in the generated list here, please check the KEP metadata.
      Please raise an issue in kubernetes/community, if the KEP metadata is correct but the generated list is incorrect.
-->

<!-- 
  - Alpha
    - [3962 - Mutating Admission Policies](https://github.com/kubernetes/enhancements/tree/master/keps/sig-api-machinery/3962-mutating-admission-policies) - v1.32
    - [4222 - CBOR Serializer](https://github.com/kubernetes/enhancements/tree/master/keps/sig-api-machinery/4222-cbor-serializer) - v1.32
    - [4346 - Add Informer Metrics](https://github.com/kubernetes/enhancements/tree/master/keps/sig-api-machinery/4346-informer-metrics) - v1.30
    - [4355 - Coordinated Leader Election](https://github.com/kubernetes/enhancements/tree/master/keps/sig-api-machinery/4355-coordinated-leader-election) - v1.31
    - [4460 - Enable per-request Read/Write Deadline](https://github.com/kubernetes/enhancements/tree/master/keps/sig-api-machinery/4460-per-request-deadline) - v1.31

  - Beta
    - [2339 - StorageVersion API for HA API servers](https://github.com/kubernetes/enhancements/tree/master/keps/sig-api-machinery/2339-storageversion-api-for-ha-api-servers) - v1.30
    - [2340 - Consistent Reads from Cache](https://github.com/kubernetes/enhancements/tree/master/keps/sig-api-machinery/2340-Consistent-reads-from-cache) - v1.31
    - [3157 - Allow informers for getting a stream of data instead of chunking](https://github.com/kubernetes/enhancements/tree/master/keps/sig-api-machinery/3157-watch-list) - v1.32
    - [4008 - CRD Validation Ratcheting](https://github.com/kubernetes/enhancements/tree/master/keps/sig-api-machinery/4008-crd-ratcheting) - v1.30
    - [4192 - Move Storage Version Migrator in-tree](https://github.com/kubernetes/enhancements/tree/master/keps/sig-api-machinery/4192-svm-in-tree) - v1.32
    - [4568 - Resilient watchcache initialization](https://github.com/kubernetes/enhancements/tree/master/keps/sig-api-machinery/4568-resilient-watchcache-initialization) - v1.31

  - Stable
    - [3352 - Aggregated Discovery](https://github.com/kubernetes/enhancements/tree/master/keps/sig-api-machinery/3352-aggregated-discovery) - v1.30
    - [3488 - CEL for Admission Control](https://github.com/kubernetes/enhancements/tree/master/keps/sig-api-machinery/3488-cel-admission-control) - 1.30
    - [3716 - Admission Webhook Match Conditions](https://github.com/kubernetes/enhancements/tree/master/keps/sig-api-machinery/3716-admission-webhook-match-conditions) - v1.30
    - [4006 - Transition from SPDY to Websockets](https://github.com/kubernetes/enhancements/tree/master/keps/sig-api-machinery/4006-transition-spdy-to-websockets) - v1.32
    - [4020 - Unknown Version Interoperability Proxy](https://github.com/kubernetes/enhancements/tree/master/keps/sig-api-machinery/4020-unknown-version-interoperability-proxy) - v1.30
    - [4358 - Custom Resource Field Selectors](https://github.com/kubernetes/enhancements/tree/master/keps/sig-api-machinery/4358-custom-resource-field-selectors) - v1.32
    - [4420 - Retry Generate Name](https://github.com/kubernetes/enhancements/tree/master/keps/sig-api-machinery/4420-retry-generate-name) - v1.32 -->

## [Subprojects](https://git.k8s.io/community/sig-api-machinery#subprojects)


**Continuing:**
  - cel-admission-webhook
  - component-base
  - control-plane-features
  - idl-schema-client-pipeline
  - json
  - kubernetes-clients
  - server-api-aggregation
  - server-binaries
  - server-crd
  - server-frameworks
  - server-sdk
  - universal-machinery
  - yaml

## [Working groups](https://git.k8s.io/community/sig-api-machinery#working-groups)

**Retired in 2024:**
 - API Expression
**Continuing:**
 - Structured Logging

## Operational

Operational tasks in [sig-governance.md]:
- [ ] [README.md] reviewed for accuracy and updated if needed
- [ ] [CONTRIBUTING.md] reviewed for accuracy and updated if needed
- [ ] Other contributing docs (e.g. in devel dir or contributor guide) reviewed for accuracy and updated if needed
- [ ] Subprojects list and linked OWNERS files in [sigs.yaml] reviewed for accuracy and updated if needed
- [ ] SIG leaders (chairs, tech leads, and subproject leads) in [sigs.yaml] are accurate and active, and updated if needed
- [ ] Meeting notes and recordings for 2024 are linked from [README.md] and updated/uploaded if needed


[CONTRIBUTING.md]: https://git.k8s.io/community/sig-api-machinery/CONTRIBUTING.md
[sig-governance.md]: https://git.k8s.io/community/committee-steering/governance/sig-governance.md
[README.md]: https://git.k8s.io/community/sig-api-machinery/README.md
[sigs.yaml]: https://git.k8s.io/community/sigs.yaml
[devel]: https://git.k8s.io/community/contributors/devel/README.md
