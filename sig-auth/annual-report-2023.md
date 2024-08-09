# 2023 Annual Report: SIG Auth

## Current initiatives and Project Health

1. What work did the SIG do this year that should be highlighted?

<!--
   Some example items that might be worth highlighting:
   - Major KEP advancement
   - Important initiatives that aren't tracked via KEPs
   - Paying down significant tech debt
   - Governance and leadership changes
-->

   - Governance and leadership changes
     - [**Mo Khan elected as new SIG tech lead**](https://groups.google.com/g/kubernetes-sig-auth/c/mHb4p8xWMR8/m/lk0UpMKXAAAJ).
     - Previous SIG TL Mike Danese stepped down during 2023 and stayed on as a chair. Many thanks for his leadership and guidance over the years.
   - The alpha `SecurityContextDeny` admission plugin was deprecated in [in v1.27](https://github.com/kubernetes/kubernetes/issues/111516) and removed in v1.30.
     - The [Pod Security Admission](https://kubernetes.io/docs/concepts/security/pod-security-admission/) plugin enforcing the
       [Pod Security Standards](https://kubernetes.io/docs/concepts/security/pod-security-standards/) `Restricted` profile captures what this plugin was trying to achieve
       in a better and up-to-date way.
   - [KEP-3325: Review attributes of a current user](https://github.com/kubernetes/enhancements/tree/master/keps/sig-auth/3325-self-subject-attributes-review-api) promoted to stable in v1.28.
     - `whoami` kubectl command promoted from `kubectl alpha` to `kubectl` [in v1.27](https://github.com/kubernetes/kubernetes/pull/116510).
   - Kubelet: security of dynamic resource allocation was enhanced by limiting node access to those objects that are needed on the node [in v1.28](https://github.com/kubernetes/kubernetes/pull/116254).
   - [KEP-3299: KMS v2 Improvements](https://github.com/kubernetes/enhancements/tree/master/keps/sig-auth/3299-kms-v2-improvements) promoted to stable in v1.29.
     - `KMSv2` is the recommended version of the KMS feature.
     - `KMSv1` was deprecated [in v1.28](https://github.com/kubernetes/kubernetes/pull/119007) and will only receive security updates going forward. Set `--feature-gates=KMSv1=true` to use the deprecated `KMSv1` feature.
   - Important initiatives that aren't tracked via KEPs:
     - Once a week issue/PR triage meetings.

2. Are there any areas and/or subprojects that your group needs help with (e.g. fewer than 2 active OWNERS)?

<!--
   Note: This list is generated from the KEP metadata in kubernetes/enhancements repository.
      If you find any discrepancy in the generated list here, please check the KEP metadata.
      Please raise an issue in kubernetes/community, if the KEP metadata is correct but the generated list is incorrect.
-->
   - The [Needs KEP / release work #sig-auth](https://docs.google.com/document/d/1sY8fRyRtk4eG9R439z5ao5i9bFuuxilS03XaNlqoni0/edit?usp=sharing) document lists multiple areas that need help and some currently have volunteers working on them.

3. Did you have community-wide updates in 2023 (e.g. KubeCon talks)?

<!--
  Examples include links to email, slides, or recordings.
-->

   - [KubeCon EU 2023] - [Kubernetes SIG Auth Deep Dive - Jordan Liggitt & Mike Danese, Google; Rita Zhang, David Eads](https://youtu.be/j9nzOLPJxAI?si=7p61DKRZ9aRwhRwe)
   - [KubeCon NA 2023] - [The Future of Kubernetes Auth and Policy Config: Common Expression Language - Mo Khan & Jordan Liggitt](https://youtu.be/yOF9S_0TO3A?si=etTKdsEZmC3EmiZc)

4. KEP work in 2023 (v1.27, v1.28, v1.29):

  - Pre-Alpha
    - [3766 - Move ReferenceGrant to sig-auth API Group](https://github.com/kubernetes/enhancements/tree/master/keps/sig-auth/3766-referencegrant)
    - [3926 - Handling undecryptable resources](https://github.com/kubernetes/enhancements/tree/master/keps/sig-auth/3926-handling-undecryptable-resources)

  - Alpha
    - [3221 - Structured Authorization Configuration](https://github.com/kubernetes/enhancements/tree/master/keps/sig-auth/3221-structured-authorization-configuration) - v1.29
    - [3257 - Cluster Trust Bundles](https://github.com/kubernetes/enhancements/tree/master/keps/sig-auth/3257-cluster-trust-bundles) - v1.29
    - [3331 - Structured authentication config](https://github.com/kubernetes/enhancements/tree/master/keps/sig-auth/3331-structured-authentication-configuration) - v1.29
    - [4193 - bound service account token improvements](https://github.com/kubernetes/enhancements/tree/master/keps/sig-auth/4193-bound-service-account-token-improvements) - v1.29

  - Stable
    - [3299 - KMS v2 Improvements](https://github.com/kubernetes/enhancements/tree/master/keps/sig-auth/3299-kms-v2-improvements) - v1.29
    - [3325 - Review attibutes of a current user](https://github.com/kubernetes/enhancements/tree/master/keps/sig-auth/3325-self-subject-attributes-review-api) - v1.28

  - Withdrawn
    - [2718 - Client Executable Proxy](https://github.com/kubernetes/enhancements/tree/master/keps/sig-auth/2718-20210511-client-exec-proxy)

## [Subprojects](https://git.k8s.io/community/sig-auth#subprojects)

**Retired in 2023:**
  - multi-tenancy

**Continuing:**
  - audit-logging
  - authenticators
  - authorizers
  - certificates
  - encryption-at-rest
  - hierarchical-namespace-controller
  - node-identity-and-isolation
  - policy-management
  - secrets-store-csi-driver
  - service-accounts
  - sig-auth-tools

## [Working groups](https://git.k8s.io/community/sig-auth#working-groups)

**Retired in 2023:**
 - Multitenancy

**Continuing:**
 - Policy

## Operational

Operational tasks in [sig-governance.md]:

- [x] [README.md] reviewed for accuracy and updated if needed
- [x] [CONTRIBUTING.md] reviewed for accuracy and updated if needed
- [x] Other contributing docs (e.g. in devel dir or contributor guide) reviewed for accuracy and updated if needed
- [x] Subprojects list and linked OWNERS files in [sigs.yaml] reviewed for accuracy and updated if needed
- [x] SIG leaders (chairs, tech leads, and subproject leads) in [sigs.yaml] are accurate and active, and updated if needed
- [x] Meeting notes and recordings for 2023 are linked from [README.md] and updated/uploaded if needed

[CONTRIBUTING.md]: https://git.k8s.io/community/sig-auth/CONTRIBUTING.md
[sig-governance.md]: https://git.k8s.io/community/committee-steering/governance/sig-governance.md
[README.md]: https://git.k8s.io/community/sig-auth/README.md
[sigs.yaml]: https://git.k8s.io/community/sigs.yaml
[devel]: https://git.k8s.io/community/contributors/devel/README.md
