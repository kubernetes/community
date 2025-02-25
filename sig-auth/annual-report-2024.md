# 2024 Annual Report: SIG Auth

## Current initiatives and Project Health

1. What work did the SIG do this year that should be highlighted?

<!--
   Some example items that might be worth highlighting:
   - Major KEP advancement
   - Important initiatives that aren't tracked via KEPs
   - Paying down significant tech debt
   - Governance and leadership changes
-->

   - No governance or leadership changes.
   - The `SecurityContextDeny` admission plugin was removed [in v1.30](https://github.com/kubernetes/kubernetes/pull/122612) after being deprecated in v1.27. The **Pod Security Admission plugin**, available since v1.25, is recommended instead.
   - Updated an audit annotation key used by the `…/serviceaccounts/<name>/token` resource handler [in v1.30](https://github.com/kubernetes/kubernetes/pull/123098). The annotation used to persist the issued credential identifier is now `authentication.kubernetes.io/issued-credential-id`.
   - Added support for configuring multiple JWT authenticators in Structured Authentication Configuration [in v1.30](https://github.com/kubernetes/kubernetes/pull/123431). The maximum allowed JWT authenticators in the authentication configuration is 64.
   - The `AuthorizationConfiguration` type accepted in --`authorization-config` files has been promoted to `apiserver.config.k8s.io/v1` [in v1.32](https://github.com/kubernetes/kubernetes/pull/128172).
   - Allowed creating ServiceAccount tokens bound to Node objects [in v1.31](https://github.com/kubernetes/kubernetes/pull/125238). This allows users to bind a service account token's validity to a named Node object, similar to Pod bound tokens. Use with `kubectl create token <serviceaccount-name> --bound-object-kind=Node --bound-object-node=<node-name>`.
   - When the alpha `UserNamespacesPodSecurityStandards` feature gate is enabled, Pod Security Admission enforcement of the baseline policy now allows `procMount=Unmasked` for user namespace pods that set `hostUsers=false` starting [in v1.31](https://github.com/kubernetes/kubernetes/pull/126163).
   - Starting [in v1.31](https://github.com/kubernetes/kubernetes/pull/126165), `container_engine_t` is in the list of allowed SELinux types in the baseline Pod Security Standards profile.
   - Starting [in v1.31](https://github.com/kubernetes/kubernetes/pull/126441), the Node Admission plugin rejects CSR requests created by a node identity for the signers `kubernetes.io/kubelet-serving` or `kubernetes.io/kube-apiserver-client-kubelet` with a CN starting with `system:node:`, but where the CN is not `system:node:${node-name}`. The feature gate `AllowInsecureKubeletCertificateSigningRequests` defaults to false, but can be enabled to revert to the previous behavior. This feature gate will be removed in Kubernetes v1.33.
   - Disallow `k8s.io` and `kubernetes.io` namespaced extra key in structured authentication configuration starting [in v1.32](https://github.com/kubernetes/kubernetes/pull/126553).
   - Starting [in v1.32](https://github.com/kubernetes/kubernetes/pull/128077), NodeRestriction admission validates the audience value that kubelet is requesting a service account token for is part of the pod spec volume. This change is introduced with a new kube-apiserver featuregate `ServiceAccountNodeAudienceRestriction` that's enabled by default in v1.32.
      - The feature gate `ServiceAccountNodeAudienceRestriction` was disabled by default in v1.32.2 to fix a regression. It is enabled by default in v1.33+.
   - Added a new SIG Auth subproject: [Secrets Store Sync Controller](https://sigs.k8s.io/secrets-store-sync-controller), a Kubernetes controller to sync from external secrets store to Kubernetes secrets.
   - Important initiatives that aren't tracked via KEPs:
     - Once a week issue/PR triage meetings.

2. Are there any areas and/or subprojects that your group needs help with (e.g. fewer than 2 active OWNERS)?

  - The [Needs KEP / release work #sig-auth](https://docs.google.com/document/d/1sY8fRyRtk4eG9R439z5ao5i9bFuuxilS03XaNlqoni0/edit?usp=sharing) document lists multiple areas that need help and some currently have volunteers working on them.

3. Did you have community-wide updates in 2024 (e.g. KubeCon talks)?

<!--
  Examples include links to email, slides, or recordings.
-->

   - [KubeCon EU 2024] - [Safety or Usability: Why Not Both? Towards Referential Auth in K8s - Rob Scott, Google & Mo Khan](https://youtu.be/HLWXuV3vJRg)
   - [KubeCon NA 2024] - [Pushing Authorization Further: CEL, Selectors and Maybe RBAC++ - Mo Khan, Rita Zhang, Jordan Liggitt](https://youtu.be/pIrJRPv-Wbg)

4. KEP work in 2024 (v1.30, v1.31, v1.32):
<!--
   TODO: Uncomment the following auto-generated list of KEPs, once reviewed & updated for correction.

   Note: This list is generated from the KEP metadata in kubernetes/enhancements repository.
      If you find any discrepancy in the generated list here, please check the KEP metadata.
      Please raise an issue in kubernetes/community, if the KEP metadata is correct but the generated list is incorrect.
-->

  - Pre-Alpha
    - [4317 - Pod Certificates](https://github.com/kubernetes/enhancements/tree/master/keps/sig-auth/4317-pod-certificates)
    - [4412 - Projected service account tokens for Kubelet image credential providers](https://github.com/kubernetes/enhancements/tree/master/keps/sig-auth/4412-projected-service-account-tokens-for-kubelet-image-credential-providers)

  - Alpha
    - [3926 - Handling undecryptable resources](https://github.com/kubernetes/enhancements/tree/master/keps/sig-auth/3926-handling-undecryptable-resources) - v1.32
    - [740 - Support external signing of service account tokens](https://github.com/kubernetes/enhancements/tree/master/keps/sig-auth/740-service-account-external-signing) - v1.32

  - Beta
    - [3331 - Structured authentication config](https://github.com/kubernetes/enhancements/tree/master/keps/sig-auth/3331-structured-authentication-configuration) - v1.30
    - [4601 - Authorize with Selectors](https://github.com/kubernetes/enhancements/tree/master/keps/sig-auth/4601-authorize-with-selectors) - v1.32
    - [4633 - Only allow anonymous auth for configured endpoints](https://github.com/kubernetes/enhancements/tree/master/keps/sig-auth/4633-anonymous-auth-configurable-endpoints) - v1.32

  - Stable
    - [2799 - Reduction of Secret-based Service Account Tokens](https://github.com/kubernetes/enhancements/tree/master/keps/sig-auth/2799-reduction-of-secret-based-service-account-token) - v1.30
    - [3221 - Structured Authorization Configuration](https://github.com/kubernetes/enhancements/tree/master/keps/sig-auth/3221-structured-authorization-configuration) - v1.32
    - [4193 - bound service account token improvements](https://github.com/kubernetes/enhancements/tree/master/keps/sig-auth/4193-bound-service-account-token-improvements) - v1.32

## [Subprojects](https://git.k8s.io/community/sig-auth#subprojects)

**New in 2024:**
  - secrets-store-sync-controller

**Continuing:**
  - audit-logging
  - authenticators
  - authorizers
  - certificates
  - encryption-at-rest
  - node-identity-and-isolation
  - policy-management
  - secrets-store-csi-driver
  - service-accounts
  - sig-auth-tools

**Archiving in 2025:**
  - hierarchical-namespace-controller

## [Working groups](https://git.k8s.io/community/sig-auth#working-groups)

**Continuing:**
 - Policy

## Operational

Operational tasks in [sig-governance.md]:

- [x] [README.md] reviewed for accuracy and updated if needed
- [x] [CONTRIBUTING.md] reviewed for accuracy and updated if needed
- [x] Other contributing docs (e.g. in devel dir or contributor guide) reviewed for accuracy and updated if needed
- [x] Subprojects list and linked OWNERS files in [sigs.yaml] reviewed for accuracy and updated if needed
- [x] SIG leaders (chairs, tech leads, and subproject leads) in [sigs.yaml] are accurate and active, and updated if needed
- [x] Meeting notes and recordings for 2024 are linked from [README.md] and updated/uploaded if needed

[CONTRIBUTING.md]: https://git.k8s.io/community/sig-auth/CONTRIBUTING.md
[sig-governance.md]: https://git.k8s.io/community/committee-steering/governance/sig-governance.md
[README.md]: https://git.k8s.io/community/sig-auth/README.md
[sigs.yaml]: https://git.k8s.io/community/sigs.yaml
[devel]: https://git.k8s.io/community/contributors/devel/README.md
