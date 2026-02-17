# 2025 Annual Report: SIG Auth

## Current initiatives and Project Health

1. What work did the SIG do this year that should be highlighted?

<!--
   Some example items that might be worth highlighting:
   - Major KEP advancement
   - Important initiatives that aren't tracked via KEPs
   - Paying down significant tech debt
   - Governance and leadership changes
-->

   - Leadership change: Mike Danese ([@mikedanese](https://github.com/mikedanese)) offboarded; Anish Ramasekar ([@aramase](https://github.com/aramase)) and Micah Hausler ([@micahhausler](https://github.com/micahhausler)) onboarded as SIG Auth chairs ([#8390](https://github.com/kubernetes/community/issues/8390)).
   - ClusterTrustBundles (KEP-3257) graduated to beta [in v1.33](https://github.com/kubernetes/enhancements/tree/master/keps/sig-auth/3257-cluster-trust-bundles), providing a cluster-scoped resource for holding X.509 trust anchors (root certificates). This API makes it easier for in-cluster certificate signers to publish and communicate X.509 trust anchors to cluster workloads.
   - Projected ServiceAccount tokens for kubelet image credential providers (KEP-4412) was introduced as alpha [in v1.33](https://github.com/kubernetes/enhancements/tree/master/keps/sig-auth/4412-projected-service-account-tokens-for-kubelet-image-credential-providers) and graduated to beta [in v1.34](https://kubernetes.io/blog/2025/09/03/kubernetes-v1-34-sa-tokens-image-pulls-beta/). The kubelet can now request short-lived, audience-bound ServiceAccount tokens for authenticating to container registries, eliminating the need for long-lived image pull secrets.
   - Structured Authentication Configuration (KEP-3331) graduated to stable [in v1.34](https://github.com/kubernetes/kubernetes/pull/131916). The `AuthenticationConfiguration` type in `--authentication-config` files was promoted to `apiserver.config.k8s.io/v1`, supporting multiple JWT authenticators, CEL expression validation, and dynamic reloading.
   - Authorize with Selectors (KEP-4601) graduated to stable [in v1.34](https://github.com/kubernetes/kubernetes/pull/132656). Authorization decisions can now leverage field and label selectors to restrict list, watch, and deletecollection operations. The `AuthorizeWithSelectors` and `AuthorizeNodeWithSelectors` feature gates were promoted to stable and locked on.
   - Anonymous auth configurable endpoints (KEP-4633) graduated to stable [in v1.34](https://github.com/kubernetes/enhancements/tree/master/keps/sig-auth/4633-anonymous-auth-configurable-endpoints). Anonymous access can now be limited to an explicit allowlist of endpoints such as `/healthz`, `/readyz`, `/livez`, reducing the blast radius of RBAC misconfigurations.
   - Pod Security Admission `baseline` and `restricted` levels now block setting `.host` field in ProbeHandler and LifecycleHandler (KEP-4940) graduated to stable [in v1.34](https://github.com/kubernetes/enhancements/tree/master/keps/sig-auth/4940-psa-block-host-field-in-probes).
   - Support for external signing of service account tokens (KEP-740) graduated to beta [in v1.34](https://github.com/kubernetes/enhancements/tree/master/keps/sig-auth/740-service-account-external-signing).
   - DRA Admin Access (KEP-5018) was introduced as alpha [in v1.33](https://github.com/kubernetes/enhancements/tree/master/keps/sig-auth/5018-dra-adminaccess) and graduated to beta [in v1.34](https://github.com/kubernetes/enhancements/tree/master/keps/sig-auth/5018-dra-adminaccess).
   - Pod Certificates (KEP-4317) was introduced as alpha [in v1.34](https://github.com/kubernetes/kubernetes/pull/128010) and graduated to beta [in v1.35](https://kubernetes.io/blog/2025/12/17/kubernetes-v1-35-release/). This enables native workload identity with automated certificate rotation, where the kubelet generates keys and requests certificates via `PodCertificateRequest`, writing credential bundles directly to the Pod's filesystem.
   - Constrained Impersonation (KEP-5284) was introduced as alpha [in v1.35](https://github.com/kubernetes/kubernetes/pull/134803), implementing a framework to prevent impersonating users from performing unauthorized actions.
   - Added ability to specify `controlplane` or `cluster` egress selectors in JWT authenticators via the `issuer.egressSelectorType` field [in v1.34](https://github.com/kubernetes/kubernetes/pull/132768), gated by the `StructuredAuthenticationConfigurationEgressSelector` beta feature gate.
   - Kube-apiserver now supports disabling caching of authorization webhook decisions in the `--authorization-config` file [in v1.34](https://github.com/kubernetes/kubernetes/pull/129237) using `cacheAuthorizedRequests` and `cacheUnauthorizedRequests` fields.
   - The NodeRestriction admission controller now disallows nodes from changing their ownerReferences [in v1.33](https://github.com/kubernetes/kubernetes/pull/133468), preventing nodes from deleting themselves by patching OwnerReferences.
   - Made `pods/exec`, `pods/attach`, and `pods/portforward` subresources require `create` permission for both SPDY and Websocket requests [in v1.35](https://github.com/kubernetes/kubernetes/pull/134577), gated by `AuthorizePodWebsocketUpgradeCreatePermission` (enabled by default).
   - Cross-SIG work (SIG Auth participating):
     - Ensure Secret Pulled Images (KEP-2535, sig-node) was introduced as alpha [in v1.33](https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/2535-ensure-secret-pulled-images) and graduated to beta [in v1.35](https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/2535-ensure-secret-pulled-images). Provides authorization to container image pulls for images already present on the node, ensuring image pull credentials are re-verified even for cached images.
     - Fine grained Kubelet API authorization (KEP-2862, sig-node) graduated to beta [in v1.33](https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/2862-fine-grained-kubelet-authz). Adds fine-grained authorization controls to the kubelet's API endpoints.
     - Introduce kuberc (KEP-3104, sig-cli) was introduced as alpha [in v1.33](https://github.com/kubernetes/enhancements/tree/master/keps/sig-cli/3104-introduce-kuberc) and graduated to beta [in v1.34](https://github.com/kubernetes/enhancements/tree/master/keps/sig-cli/3104-introduce-kuberc). Adds `credPluginPolicy` and `credPluginAllowlist` configuration options for credential plugin management.
     - CSI driver opt-in for service account tokens via secrets field (KEP-5538, sig-storage) graduated to beta [in v1.35](https://github.com/kubernetes/enhancements/tree/master/keps/sig-storage/5538-csi-sa-tokens-secrets-field). CSI drivers can now opt in to receive service account tokens via the secrets field instead of volume context, separating credentials from metadata to prevent accidental leakage.
   - Important initiatives that aren't tracked via KEPs:
     - Once a week issue/PR triage meetings.
     - Retired the hierarchical-namespace-controller subproject and the policy working group.

2. Are there any areas and/or subprojects that your group needs help with (e.g. fewer than 2 active OWNERS)?

  - The [Needs KEP / release work #sig-auth](https://docs.google.com/document/d/1sY8fRyRtk4eG9R439z5ao5i9bFuuxilS03XaNlqoni0/edit?usp=sharing) document lists multiple areas that need help and some currently have volunteers working on them.

3. Did you have community-wide updates in 2025 (e.g. KubeCon talks)?

<!--
  Examples include links to email, slides, or recordings.
-->

   - [KubeCon EU 2025] - [Strengthening Auth in Kubernetes: Image Pulling, DRA Admin Access & Pod Certificates - Rita Zhang & Stanislav Láznička, Microsoft](https://kccnceu2025.sched.com/event/1tczC)
   - [KubeCon EU 2025] - [A Practical Guide To Kubernetes Policy as Code - Jim Bugwadia, Nirmata; Rita Zhang, Microsoft; Andy Suderman, Fairwinds; Joe Betz, Google](https://kccnceu2025.sched.com/event/1tcxh)
   - [KubeCon NA 2025] - [Strengthening Kubernetes Trust: SIG Auth's Latest Security Enhancements - Anish Ramasekar, Mo Khan, Stanislav Láznička, Rita Zhang, Peter Engelbert, Microsoft](https://kccncna2025.sched.com/event/27Nld)

4. KEP work in 2025 (v1.33, v1.34, v1.35):
<!--
   Note: This list is generated from the KEP metadata in kubernetes/enhancements repository.
      If you find any discrepancy in the generated list here, please check the KEP metadata.
      Please raise an issue in kubernetes/community, if the KEP metadata is correct but the generated list is incorrect.
-->

  - Alpha
    - [5284 - Constrained Impersonation](https://github.com/kubernetes/enhancements/tree/master/keps/sig-auth/5284-constrained-impersonation) - v1.35

  - Beta
    - [3257 - Cluster Trust Bundles](https://github.com/kubernetes/enhancements/tree/master/keps/sig-auth/3257-cluster-trust-bundles) - v1.33
    - [4317 - Pod Certificates](https://github.com/kubernetes/enhancements/tree/master/keps/sig-auth/4317-pod-certificates) - v1.35
    - [4412 - Projected Service Account Tokens for Kubelet Image Credential Providers](https://github.com/kubernetes/enhancements/tree/master/keps/sig-auth/4412-projected-service-account-tokens-for-kubelet-image-credential-providers) - v1.34
    - [5018 - DRA Admin Access](https://github.com/kubernetes/enhancements/tree/master/keps/sig-auth/5018-dra-adminaccess) - v1.34
    - [740 - Support external signing of service account tokens](https://github.com/kubernetes/enhancements/tree/master/keps/sig-auth/740-service-account-external-signing) - v1.34

  - Stable
    - [3331 - Structured authentication config](https://github.com/kubernetes/enhancements/tree/master/keps/sig-auth/3331-structured-authentication-configuration) - v1.34
    - [4601 - Authorize with Selectors](https://github.com/kubernetes/enhancements/tree/master/keps/sig-auth/4601-authorize-with-selectors) - v1.34
    - [4633 - Only allow anonymous auth for configured endpoints](https://github.com/kubernetes/enhancements/tree/master/keps/sig-auth/4633-anonymous-auth-configurable-endpoints) - v1.34
    - [4940 - Pod Security Admission `baseline` and `restricted` levels now block setting `.host` field in ProbeHandler and LifecycleHandler](https://github.com/kubernetes/enhancements/tree/master/keps/sig-auth/4940-psa-block-host-field-in-probes) - v1.34

## [Subprojects](https://git.k8s.io/community/sig-auth#subprojects)

**Retired in 2025:**
  - hierarchical-namespace-controller

**Continuing:**
  - audit-logging
  - authenticators
  - authorizers
  - certificates
  - encryption-at-rest
  - node-identity-and-isolation
  - policy-management
  - secrets-store-csi-driver
  - secrets-store-sync-controller
  - service-accounts
  - sig-auth-tools

## [Working groups](https://git.k8s.io/community/sig-auth#working-groups)

**New in 2025:**
 - AI Integration
 - Checkpoint Restore

**Retired in 2025:**
 - Policy

## Operational

Operational tasks in [sig-governance.md]:
- [x] [README.md] reviewed for accuracy and updated if needed
- [x] [CONTRIBUTING.md] reviewed for accuracy and updated if needed
- [x] Other contributing docs (e.g. in devel dir or contributor guide) reviewed for accuracy and updated if needed
- [x] Subprojects list and linked OWNERS files in [sigs.yaml] reviewed for accuracy and updated if needed
- [x] SIG leaders (chairs, tech leads, and subproject leads) in [sigs.yaml] are accurate and active, and updated if needed
- [x] Meeting notes and recordings for 2025 are linked from [README.md] and updated/uploaded if needed


[CONTRIBUTING.md]: https://git.k8s.io/community/sig-auth/CONTRIBUTING.md
[sig-governance.md]: https://git.k8s.io/community/committee-steering/governance/sig-governance.md
[README.md]: https://git.k8s.io/community/sig-auth/README.md
[sigs.yaml]: https://git.k8s.io/community/sigs.yaml
[devel]: https://git.k8s.io/community/contributors/devel/README.md
