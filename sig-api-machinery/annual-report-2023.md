# 2023 Annual Report: SIG API Machinery

## Current initiatives and Project Health

1. What work did the SIG do this year that should be highlighted?

<!--
   Some example items that might be worth highlighting:
   - Major KEP advancement
   - Important initiatives that aren't tracked via KEPs
   - Paying down significant tech debt
   - Governance and leadership changes
-->

2. Are there any areas and/or subprojects that your group needs help with (e.g. fewer than 2 active OWNERS)?

<!--
   Note: This list is generated from the KEP metadata in kubernetes/enhancements repository.
      If you find any discrepancy in the generated list here, please check the KEP metadata.
      Please raise an issue in kubernetes/community, if the KEP metadata is correct but the generated list is incorrect.
-->

3. Did you have community-wide updates in 2023 (e.g. KubeCon talks)?

<!--
  Examples include links to email, slides, or recordings.
-->

4. KEP work in 2023 (v1.27, v1.28, v1.29):

  - Alpha
    - [2340 - Consistent Reads from Cache](https://github.com/kubernetes/enhancements/tree/master/keps/sig-api-machinery/2340-Consistent-reads-from-cache) - v1.28
    - [3157 - Allow informers for getting a stream of data instead of chunking](https://github.com/kubernetes/enhancements/tree/master/keps/sig-api-machinery/3157-watch-list) - v1.27
    - [4006 - Transition from SPDY to Websockets](https://github.com/kubernetes/enhancements/tree/master/keps/sig-api-machinery/4006-transition-spdy-to-websockets) - v1.29
    - [4008 - CRD Validation Ratcheting](https://github.com/kubernetes/enhancements/tree/master/keps/sig-api-machinery/4008-crd-ratcheting) - v1.28
    - [4153 - Declarative Validation](https://github.com/kubernetes/enhancements/tree/master/keps/sig-api-machinery/4153-declarative-validation) - v1.29

  - Beta
    - [3352 - Aggregated Discovery](https://github.com/kubernetes/enhancements/tree/master/keps/sig-api-machinery/3352-aggregated-discovery) - v1.27
    - [3488 - CEL for Admission Control](https://github.com/kubernetes/enhancements/tree/master/keps/sig-api-machinery/3488-cel-admission-control) - 1.28
    - [3716 - Admission Webhook Match Conditions](https://github.com/kubernetes/enhancements/tree/master/keps/sig-api-machinery/3716-admission-webhook-match-conditions) - v1.28
    - [4020 - Unknown Version Interoperability Proxy](https://github.com/kubernetes/enhancements/tree/master/keps/sig-api-machinery/4020-unknown-version-interoperability-proxy) - v1.29

  - Stable
    - [1040 - Priority and Fairness for API Server Requests](https://github.com/kubernetes/enhancements/tree/master/keps/sig-api-machinery/1040-priority-and-fairness) - v1.29
    - [2876 - CRD Validation Expression Language](https://github.com/kubernetes/enhancements/tree/master/keps/sig-api-machinery/2876-crd-validation-expression-language) - v1.29
    - [2885 - Server Side Unknown Field Validation](https://github.com/kubernetes/enhancements/tree/master/keps/sig-api-machinery/2885-server-side-unknown-field-validation) - v1.27
    - [2896 - OpenAPI V3](https://github.com/kubernetes/enhancements/tree/master/keps/sig-api-machinery/2896-openapi-v3) - v1.27
    - [365 - Paginated API Lists](https://github.com/kubernetes/enhancements/tree/master/keps/sig-api-machinery/365-paginated-lists) - v1.29

## [Subprojects](https://git.k8s.io/community/sig-api-machinery#subprojects)


**New in 2023:**
  - [cel-admission-webhook](https://git.k8s.io/community/<no value>#cel-admission-webhook)
**Continuing:**
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

**Retired in 2023:**
 - Multitenancy
**Continuing:**
 - API Expression
 - Structured Logging

## Operational

Operational tasks in [sig-governance.md]:
- [ ] [README.md] reviewed for accuracy and updated if needed
- [ ] [CONTRIBUTING.md] reviewed for accuracy and updated if needed
- [ ] Other contributing docs (e.g. in devel dir or contributor guide) reviewed for accuracy and updated if needed
- [ ] Subprojects list and linked OWNERS files in [sigs.yaml] reviewed for accuracy and updated if needed
- [ ] SIG leaders (chairs, tech leads, and subproject leads) in [sigs.yaml] are accurate and active, and updated if needed
- [ ] Meeting notes and recordings for 2023 are linked from [README.md] and updated/uploaded if needed


[CONTRIBUTING.md]: https://git.k8s.io/community/sig-api-machinery/CONTRIBUTING.md
[sig-governance.md]: https://git.k8s.io/community/committee-steering/governance/sig-governance.md
[README.md]: https://git.k8s.io/community/sig-api-machinery/README.md
[sigs.yaml]: https://git.k8s.io/community/sigs.yaml
[devel]: https://git.k8s.io/community/contributors/devel/README.md
