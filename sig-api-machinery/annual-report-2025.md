# 2025 Annual Report: SIG API Machinery

## Current initiatives and Project Health

   - Carefully reviewed on each release the list of open KEPs and which ones were going to make it in the release.
   - Created two new subprojects: **crdify** and **kube-api-linter**, reflecting ongoing investment in CRD tooling and API quality.
   - Launched the **WG AI Integration** working group under SIG API Machinery.
   - Created the **Declarative APIs and Linters** subproject meeting (biweekly, Tuesdays 9:00 PT), carrying forward the goals of the retired WG API Expression. The first meeting was held on September 23, 2025. This subproject encompasses the new `crdify` and `kube-api-linter` repos.
   - [2025 Contributor Awards](https://www.kubernetes.dev/community/awards/2025/#api-machinery):
     + Aaron Prindle, @aaron-prindle — for driving the declarative validation framework (KEP-5073) and creating the validation-gen code generator.
     + Joel Speed, @JoelSpeed — for valuable contributions to the kube-api-linter project, improving API quality and consistency tooling.
     + Yongrui Lin, @yongruilin — for significant hands-on development and refinement of validation-gen and its integration into the Kubernetes codebase.
   - Some KEPs that deserve special mention in 2025 are
     + Graduated to *Stable*
       - [Consistent Reads from Cache](https://github.com/kubernetes/enhancements/tree/master/keps/sig-api-machinery/2340-Consistent-reads-from-cache) — Serving consistent reads from watch cache instead of etcd, dramatically improving API server scalability.
       - [Streaming List Responses (Streaming Encoding)](https://github.com/kubernetes/enhancements/tree/master/keps/sig-api-machinery/5116-streaming-response-encoding) — Streaming encoding for JSON and Protobuf list responses, eliminating large memory allocations on the API server.
       - [Ordered Namespace Deletion (KEP-5080)](https://github.com/kubernetes/enhancements/issues/5080) — Deterministic resource deletion order for namespaces, mitigating security risks from non-deterministic deletions (CVE-2024-7598).
       - [Resilient Watch Cache Initialization (KEP-4568)](https://github.com/kubernetes/enhancements/issues/4568) — Made the watch cache initialization more resilient to failures, improving control plane robustness.
       - [Remove gogo protobuf dependency for API types (KEP-5589)](https://github.com/kubernetes/enhancements/issues/5589) — Migrated Kubernetes API types from the deprecated gogo protobuf library to the standard Go protobuf library.
       - [Transition from SPDY to WebSockets](https://github.com/kubernetes/enhancements/tree/master/keps/sig-api-machinery/4006-transition-spdy-to-websockets) — Completed the full transition of exec/attach/port-forward from SPDY to WebSockets.
       - [Coordinated Leader Election (KEP-3962)](https://github.com/kubernetes/enhancements/issues/3962) — Leader election improvements for better control plane stability.
     + Graduated to *Beta*
       - [Watch List / Streaming Initial List (KEP-3157)](https://github.com/kubernetes/enhancements/tree/master/keps/sig-api-machinery/3157-watch-list) — Enabled by default in client-go in v1.34, allowing informers to get initial data via streaming watch instead of chunked LIST, reducing API server memory pressure.
       - [Declarative Validation of Kubernetes Native Types (KEP-5073)](https://github.com/kubernetes/enhancements/issues/5073) — CEL-based declarative validation rules for built-in Kubernetes types using validation-gen, enabled by default in v1.33.
       - [Snapshottable API Server Cache (KEP-4988)](https://github.com/kubernetes/enhancements/issues/4988) — Allows the watch cache to generate efficient point-in-time snapshots, enabling paginated LIST requests to be served entirely from cache (beta in v1.34).
       - [List from Cache Snapshot](https://github.com/kubernetes/enhancements/issues/4988) — kube-apiserver can serve LIST requests for previous resource versions from cache snapshots rather than etcd.

2. Are there any areas and/or subprojects that your group needs help with (e.g. fewer than 2 active OWNERS)?

   - API Machinery area is very extensive, here are some areas that could use extra help:
     + Server Side Apply
     + Resource Lifecycle (examples: Garbage Collection, Storage Version Migrator, Namespace Deletion, CRD lifecycle, etc)
     + Controllers Infrastructure
     + Clients ecosystem

3. Did you have community-wide updates in 2025 (e.g. KubeCon talks)?

 - KubeCon EU 2025 (London): [SIG API Machinery: Project Updates and Release Planning - Joe Betz, Google](https://kccnceu2025.sched.com/event/1tcz3/sig-api-machinery-project-updates-and-release-planning-joe-betz-google)
   - KubeCon NA 2025 (Atlanta): [SIG API Machinery and AI: What Comes Next - Joe Betz, Google & David Eads, Red Hat](https://kccncna2025.sched.com/event/27Nnf/sig-api-machinery-and-ai-what-comes-next-joe-betz-google-david-eads-red-hat)
   - [Kubernetes v1.33: Streaming List responses blog post](https://kubernetes.io/blog/2025/05/09/kubernetes-v1-33-streaming-list-responses/)
   - [Kubernetes v1.34: Snapshottable API server cache blog post](https://kubernetes.io/blog/2025/09/09/kubernetes-v1-34-snapshottable-api-server-cache/)


4. KEP work in 2025 (v1.33, v1.34, v1.35):

  - Alpha
    - [4595 - CEL for CRD AdditionalPrinterColumns](https://github.com/kubernetes/enhancements/tree/master/keps/sig-api-machinery/4595-cel-crd-additionalprintercolumns) - 1.34
    - [5366 - Graceful Leader Transition](https://github.com/kubernetes/enhancements/tree/master/keps/sig-api-machinery/5366-graceful-leader-transition) - v1.35
    - [5647 - Stale Controller Handling](https://github.com/kubernetes/enhancements/tree/master/keps/sig-api-machinery/5647-stale-controller-handling) - v1.35

  - Beta
    - [3962 - Mutating Admission Policies](https://github.com/kubernetes/enhancements/tree/master/keps/sig-api-machinery/3962-mutating-admission-policies) - v1.34
    - [4192 - Move Storage Version Migrator in-tree](https://github.com/kubernetes/enhancements/tree/master/keps/sig-api-machinery/4192-svm-in-tree) - v1.35
    - [4355 - Coordinated Leader Election](https://github.com/kubernetes/enhancements/tree/master/keps/sig-api-machinery/4355-coordinated-leader-election) - v1.33
    - [4988 - Snapshottable API server cache](https://github.com/kubernetes/enhancements/tree/master/keps/sig-api-machinery/4988-snapshottable-api-server-cache) - v1.34
    - [5073 - Declarative Validation](https://github.com/kubernetes/enhancements/tree/master/keps/sig-api-machinery/5073-declarative-validation-with-validation-gen) - v1.33

  - Stable
    - [2340 - Consistent Reads from Cache](https://github.com/kubernetes/enhancements/tree/master/keps/sig-api-machinery/2340-Consistent-reads-from-cache) - v1.34
    - [4008 - CRD Validation Ratcheting](https://github.com/kubernetes/enhancements/tree/master/keps/sig-api-machinery/4008-crd-ratcheting) - v1.33
    - [4568 - Resilient watchcache initialization](https://github.com/kubernetes/enhancements/tree/master/keps/sig-api-machinery/4568-resilient-watchcache-initialization) - v1.34
    - [5080 - Ordered Namespace Deletion](https://github.com/kubernetes/enhancements/tree/master/keps/sig-api-machinery/5080-ordered-namespace-deletion) - 1.34
    - [5116 - Streaming JSON Encoding for LIST Responses](https://github.com/kubernetes/enhancements/tree/master/keps/sig-api-machinery/5116-streaming-response-encoding) - v1.34
    - [5504 - Comparable Resource Version](https://github.com/kubernetes/enhancements/tree/master/keps/sig-api-machinery/5504-comparable-resource-version) - v1.35
    - [5589 - Remove gogo protobuf dependency](https://github.com/kubernetes/enhancements/tree/master/keps/sig-api-machinery/5589-gogo-dependency) - v1.35

## [Subprojects](https://git.k8s.io/community/sig-api-machinery#subprojects)


**New in 2025:**
  - crdify
  - kube-api-linter
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

**New in 2025:**
 - AI Integration
**Continuing:**
 - Structured Logging

## Operational

Operational tasks in [sig-governance.md]:
- [x] [README.md](https://git.k8s.io/community/sig-api-machinery/README.md) reviewed for accuracy and updated if needed
- [x] [CONTRIBUTING.md](https://git.k8s.io/community/sig-api-machinery/CONTRIBUTING.md) reviewed for accuracy and updated if needed
- [x] Other contributing docs (e.g. in devel dir or contributor guide) reviewed for accuracy and updated if needed
- [x] Subprojects list and linked OWNERS files in [sigs.yaml](https://git.k8s.io/community/sigs.yaml) reviewed for accuracy and updated if needed
- [x] SIG leaders (chairs, tech leads, and subproject leads) in [sigs.yaml](https://git.k8s.io/community/sigs.yaml) are accurate and active, and updated if needed
- [x] Meeting notes and recordings for 2025 are linked from [README.md](https://git.k8s.io/community/sig-api-machinery/README.md) and updated/uploaded if needed


[CONTRIBUTING.md]: https://git.k8s.io/community/sig-api-machinery/CONTRIBUTING.md
[sig-governance.md]: https://git.k8s.io/community/committee-steering/governance/sig-governance.md
[README.md]: https://git.k8s.io/community/sig-api-machinery/README.md
[sigs.yaml]: https://git.k8s.io/community/sigs.yaml
[devel]: https://git.k8s.io/community/contributors/devel/README.md
