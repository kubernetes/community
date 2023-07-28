# 2022 Annual Report: SIG Auth

## Current initiatives

1. What work did the SIG do this year that should be highlighted?

   - `kubectl create token` can be used to request a service account token [starting v1.24](https://github.com/kubernetes/kubernetes/pull/107880), and permission to request service account tokens is added to the `edit` and `admin` RBAC roles.
   - The CertificateSigningRequest `spec.expirationSeconds` API field has graduated to GA [in v1.24](https://github.com/kubernetes/kubernetes/pull/108782).
   - The `client.authentication.k8s.io/v1alpha1` ExecCredential has been removed [in v1.24](https://github.com/kubernetes/kubernetes/pull/108616). If you are using a client-go credential plugin that relies on the v1alpha1 API please contact the distributor of your plugin for instructions on how to migrate to the v1 API.
   - The `LegacyServiceAccountTokenNoAutoGeneration` feature gate is beta, and enabled by default [in v1.24](https://github.com/kubernetes/kubernetes/pull/108309). When enabled, Secret API objects containing service account tokens are no longer auto-generated for every ServiceAccount. Use the TokenRequest API to acquire service account tokens, or if a non-expiring token is required, create a Secret API object for the token controller to populate with a service account token by following this [guide](https://kubernetes.io/docs/concepts/configuration/secret/#service-account-token-secrets).
   - Kube-apiserver: `--audit-log-version` and `--audit-webhook-version` only support the default value of `audit.k8s.io/v1` [starting v1.24](https://github.com/kubernetes/kubernetes/pull/108092). The v1alpha1 and v1beta1 audit log versions, deprecated since 1.13, have been removed.
   - The `gcp` and `azure` auth plugins have been removed from client-go and kubectl [in v1.26](https://github.com/kubernetes/kubernetes/pull/110013). See https://github.com/Azure/kubelogin and https://cloud.google.com/blog/products/containers-kubernetes/kubectl-auth-changes-in-gke.
   - If the parent directory of the file specified in the `--audit-log-path` argument does not exist, Kubernetes now creates it [starting v1.25](https://github.com/kubernetes/kubernetes/pull/110813).
   - KMS v2alpha1 API added [in v1.25](https://github.com/kubernetes/kubernetes/pull/111126).
   - API server's deprecated `--service-account-api-audiences` flag is removed [in v1.25](https://github.com/kubernetes/kubernetes/pull/108624). Use `--api-audiences` instead.
   - As [of v1.25](https://github.com/kubernetes/kubernetes/pull/105919), the PodSecurity `restricted` level no longer requires pods that set .spec.os.name="windows" to also set Linux-specific securityContext fields. If a 1.25+ cluster has unsupported [out-of-skew](https://kubernetes.io/releases/version-skew-policy/#kubelet) nodes prior to v1.23 and wants to ensure namespaces enforcing the `restricted` policy continue to require Linux-specific securityContext fields on all pods, ensure a version of the `restricted` prior to v1.25 is selected by labeling the namespace (for example, `pod-security.kubernetes.io/enforce-version: v1.24`).
   - The PodSecurity admission plugin has graduated to GA and is enabled by default [in v1.25](https://github.com/kubernetes/kubernetes/pull/110459). The admission configuration version has been promoted to `pod-security.admission.config.k8s.io/v1`.
   - The beta `PodSecurityPolicy` admission plugin, deprecated since 1.21, is removed [in v1.25](https://github.com/kubernetes/kubernetes/pull/109798). Follow the instructions at https://kubernetes.io/docs/tasks/configure-pod-container/migrate-from-psp/ to migrate to the built-in PodSecurity admission plugin (or to another third-party policy webhook) prior to upgrading to v1.25.
   - Return a warning when applying a `pod-security.kubernetes.io` label to a PodSecurity-exempted namespace. Stop including the `pod-security.kubernetes.io/exempt=namespace` audit annotation on namespace requests [in v1.25](https://github.com/kubernetes/kubernetes/pull/109680)
   - Kube-controller-manager's deprecated `--experimental-cluster-signing-duration` flag is removed [in v1.25](https://github.com/kubernetes/kubernetes/pull/108476). Adapt your machinery to use the `--cluster-signing-duration` flag that is available since v1.19.
   - Add auth API to get self subject attributes (new selfsubjectreviews API is added). The corresponding command for kubectl - `kubectl auth whoami` is provided [in v1.26](https://github.com/kubernetes/kubernetes/pull/111333)
   - Kube-apiserver: custom resources can be specified in the `--encryption-provider-config` file and can be encrypted in etcd [starting v1.26](https://github.com/kubernetes/kubernetes/pull/113015).
   - When the alpha LegacyServiceAccountTokenTracking feature gate is enabled, secret-based service account tokens will have a `kubernetes.io/legacy-token-last-used` applied to them containing the date they were last used [starting v1.26](https://github.com/kubernetes/kubernetes/pull/108858)
   - A new API server flag `--encryption-provider-config-automatic-reload` has been added [in v1.26](https://github.com/kubernetes/kubernetes/pull/113529) to control when the encryption config should be automatically reloaded without needing to restart the server. All KMS plugins are merged into a single healthz check at /healthz/kms-providers when reload is enabled, or when only KMS v2 plugins are used.
   - The `LegacyServiceAccountTokenNoAutoGeneration` feature gate has been promoted to GA [in v1.26](https://github.com/kubernetes/kubernetes/pull/112838).
   - Pod Security admission: the pod-security `warn` level will default to the `enforce` level [starting v1.26](https://github.com/kubernetes/kubernetes/pull/113491).
   - Kubectl config view now automatically redacts any secret fields marked with a datapolicy tag [starting v1.26](https://github.com/kubernetes/kubernetes/pull/109189).
   - Introduce v1alpha1 API for validating admission policies [in v1.26](https://github.com/kubernetes/kubernetes/pull/113314), enabling extensible admission control via CEL expressions (KEP 3488: CEL for Admission Control). To use, enable the ValidatingAdmissionPolicy feature gate and the `admissionregistration.k8s.io/v1alpha1` API via `--runtime-config`.
   - Callers using DelegatingAuthenticationOptions can use DisableAnonymous to disable Anonymous authentication [in v1.26](https://github.com/kubernetes/kubernetes/pull/112181).


2. What initiatives are you working on that aren't being tracked in KEPs?

   - Once a week issue/PR triage meetings.
     - [Automation of the project board population](https://github.com/kubernetes-sigs/sig-auth-tools)


3. KEP work in 2022 (v1.24, v1.25, v1.26):
  - pre-alpha:
    - [2718 - Client Executable Proxy](https://github.com/kubernetes/enhancements/tree/master/keps/sig-auth/2718-20210511-client-exec-proxy) - v1.26
  - alpha:
    - [3299 - KMS v2 Improvements](https://github.com/kubernetes/enhancements/tree/master/keps/sig-auth/3299-kms-v2-improvements) - v1.25
    - [3325 - Self subject review API](https://github.com/kubernetes/enhancements/tree/master/keps/sig-auth/3325-self-subject-attributes-review-api) - v1.26
  - stable:
    - [2579 - PSP Replacement Policy](https://github.com/kubernetes/enhancements/tree/master/keps/sig-auth/2579-psp-replacement) - v1.25
    - [2784 - CSR Duration](https://github.com/kubernetes/enhancements/tree/master/keps/sig-auth/2784-csr-duration) - v1.24
    - [2799 - Reduction of Secret-based Service Account Tokens](https://github.com/kubernetes/enhancements/blob/master/keps/sig-auth/2799-reduction-of-secret-based-service-account-token) - v1.26


## Project health

1. What areas and/or subprojects does your group need the most help with?
   Any areas with 2 or fewer OWNERs? (link to more details)

   -
   -
   -

2. What metrics/community health stats does your group care about and/or measure?

   - Based on devstats [Issue Velocity / Inactive Issues by SIG for 90 days or more](https://k8s.devstats.cncf.io/d/73/inactive-issues-by-sig?orgId=1&var-sigs=%22auth%22) at the time of writing this report, average is 8.
   - Based on devstats [PR Velocity / Awaiting PRs by SIG for 90 days or more](https://k8s.devstats.cncf.io/d/70/awaiting-prs-by-sig?orgId=1&var-sigs=%22auth%22) at the time of writing this report, average is 75.

3. Does your [CONTRIBUTING.md] help **new** contributors engage with your group specifically by pointing
   to activities or programs that provide useful context or allow easy participation?

   - Currently there is no onboarding or growth path. This is something we are working on and learning from other SIGs.

4. If your group has special training, requirements for reviewers/approvers, or processes beyond the general [contributor guide],
   does your [CONTRIBUTING.md] document those to help **existing** contributors grow throughout the [contributor ladder]?

   - Currently there is no onboarding or growth path. This is something we are working on and learning from other SIGs.

5. Does the group have contributors from multiple companies/affiliations?

   - Yes. Our chairs, leads, contributors, participants, and subproject owners are from various companies.

6. Are there ways end users/companies can contribute that they currently are not?
   If one of those ways is more full time support, what would they work on and why?

   -
   -

## Membership

- Primary slack channel member count: 2847
- Primary mailing list member count: 462
- Primary meeting attendee count (estimated, if needed): 20 ~ 30
- Primary meeting participant count (estimated, if needed): 5 ~ 10
- Unique reviewers for SIG-owned packages: 15
- Unique approvers for SIG-owned packages: 7

Include any other ways you measure group membership

## [Subprojects](https://git.k8s.io/community/sig-auth#subprojects)

**New in 2022:**

  - [sig-auth-tools](https://github.com/kubernetes-sigs/sig-auth-tools)
  - [pspmigrator](https://github.com/kubernetes-sigs/pspmigrator)

**Continuing:**

  - audit-logging
  - authenticators
  - authorizers
  - certificates
  - encryption-at-rest
  - hierarchical-namespace-controller
  - multi-tenancy
  - node-identity-and-isolation
  - policy-management
  - secrets-store-csi-driver
  - service-accounts


## [Working groups](https://git.k8s.io/community/sig-auth#working-groups)


**Continuing:**
- All working groups under https://github.com/kubernetes/community/blob/master/sig-auth/README.md#working-groups have continued.

## Operational

Operational tasks in [sig-governance.md]:

- [x] [README.md] reviewed for accuracy and updated if needed
- [x] [CONTRIBUTING.md] reviewed for accuracy and updated if needed
      (or created if missing and your contributor steps and experience are different or more
      in-depth than the documentation listed in the general [contributor guide] and [devel] folder.)
- [x] Subprojects list and linked OWNERS files in [sigs.yaml] reviewed for accuracy and updated if needed
- [x] SIG leaders (chairs, tech leads, and subproject owners) in [sigs.yaml] are accurate and active, and updated if needed
- [x] Meeting notes and recordings for 2022 are linked from [README.md] and updated/uploaded if needed
- [x] Did you have community-wide updates in 2022 (e.g. community meetings, kubecon, or kubernetes-dev@ emails)? Links to email, slides, or recordings:
      - 2022 Kubecon EU Virtual - [SIG Auth Deep Dive](https://sched.co/ytpT) [session recording](https://youtu.be/C3Ak35W55m0)
      - 2022 Kubecon NA - [SIG Auth Deep Dive](https://sched.co/182PB) [session recording](https://youtu.be/QbqpPZxDKDw)

[CONTRIBUTING.md]: https://git.k8s.io/community/sig-auth/CONTRIBUTING.md
[contributor ladder]: https://git.k8s.io/community/community-membership.md
[sig-governance.md]: https://git.k8s.io/community/committee-steering/governance/sig-governance.md
[README.md]: https://git.k8s.io/community/sig-auth/README.md
[sigs.yaml]: https://git.k8s.io/community/sigs.yaml
[contributor guide]: https://git.k8s.io/community/contributors/guide/README.md
[devel]: https://git.k8s.io/community/contributors/devel/README.md
