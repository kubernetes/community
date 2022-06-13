# 2021 Annual Report: SIG Auth

## Current initiatives

1. What work did the SIG do this year that should be highlighted?

   - [Pod Security admission](https://kubernetes.io/docs/concepts/security/pod-security-admission/) has [graduated to beta](https://github.com/kubernetes/kubernetes/pull/106089) and is enabled by default. The admission configuration version has been promoted to `pod-security.admission.config.k8s.io/v1beta1` in v1.23.
   - The [PodSecurityPolicy API is deprecated in v1.21](https://github.com/kubernetes/kubernetes/pull/97171), and will no longer be served starting in v1.25.
   - Marking `audit.k8s.io/v1[alpha|beta]1` versions as deprecated and warning if a version other than `audit.k8s.io/v1` was passed to the kube-apiserver flags `--audit-log-version` and `--audit-webhook-version` [in v1.21](https://github.com/kubernetes/kubernetes/pull/98858).
   - [PodSecurityPolicy only stores "generic" as allowed volume type](https://github.com/kubernetes/kubernetes/pull/98918) if the GenericEphemeralVolume feature gate is enabled
   - RunAsGroup feature for Containers in a Pod [graduates to GA in v1.21](https://github.com/kubernetes/kubernetes/pull/94641)
   - RootCAConfigMap feature [graduates to GA in v1.21](https://github.com/kubernetes/kubernetes/pull/98033)
   - The ServiceAccountIssuerDiscovery feature has [graduated to GA](https://github.com/kubernetes/kubernetes/pull/98553), and is unconditionally enabled in v1.21.
   - CSIServiceAccountToken [graduates to GA](https://github.com/kubernetes/kubernetes/pull/103001) in 1.22
   - Mark `net.ipv4.ip_unprivileged_port_start` as safe sysctl [in v1.22](https://github.com/kubernetes/kubernetes/pull/103326)
   - BoundServiceAccountTokenVolume [graduates to GA in v1.22](https://github.com/kubernetes/kubernetes/pull/101992)
   - Kubernetes client [credential plugins](https://kubernetes.io/docs/reference/access-authn-authz/authentication/#client-go-credential-plugins) feature graduates to stable in v1.22. The GA feature set includes improved support for plugins that provide interactive login flows. The in-tree Azure and GCP authentication plugins have been [deprecated](https://github.com/kubernetes/kubernetes/pull/102181) in favor of out-of-tree implementations.
   - Kube-apiserver `--service-account-issuer` can be specified multiple times now, to enable non-disruptive change of issuer [starting v1.22](https://github.com/kubernetes/kubernetes/pull/101155)
   - The `CertificateSigningRequest.certificates.k8s.io` API supports an optional expirationSeconds field to allow the client to request a particular duration for the issued certificate. The default signer implementations provided by the Kubernetes controller manager will honor this field as long as it does not exceed the `--cluster-signing-duration` flag [starting v1.22](https://github.com/kubernetes/kubernetes/pull/99494).
   - Aggregate write permissions on events to edit and admin role [starting v1.22](https://github.com/kubernetes/kubernetes/pull/102858)
   - The kubelet now reports distinguishes log messages about certificate rotation for its client cert and server cert separately to make debugging problems with one or the other easier.[starting v1.22](https://github.com/kubernetes/kubernetes/pull/101252)
   - A new field `omitManagedFields` has been added to both `audit.Policy` and `audit.PolicyRule` so cluster operators can opt in to omit managed fields of the request and response bodies from being written to the API audit log [starting v1.23](https://github.com/kubernetes/kubernetes/pull/94986)
   - Adds `--as-uid` flag to kubectl to allow uid impersonation in the same way as user and group impersonation [starting v1.23](https://github.com/kubernetes/kubernetes/pull/105794)

2. What initiatives are you working on that aren't being tracked in KEPs?
   SIG Auth leads have curated and broadcasted a list of work from the `Needs KEP` swimlane out of the [#sig-auth board](https://github.com/orgs/kubernetes/projects/54) in the [Needs KEP / release work #sig-auth](https://docs.google.com/document/d/1sY8fRyRtk4eG9R439z5ao5i9bFuuxilS03XaNlqoni0/edit) living document. The call-out to the community is a way of looking for folks to both lead the design work necessary to get these KEPs into an implementable state, as well as to land the implementation into the Kubernetes codebase. Specifically:
   - [KMS-Plugin: Improvements](https://docs.google.com/document/d/1YHzSzITSS3ZNpf63E-rseDo-ocpxexp3ttzjBU2P8Ck/edit?usp=sharing)
   - Specifying multiple webhooks in the kube-apiserver authorization chain
   - Structured config for OIDC authentication
   - Audit logging improvements
   - system:masters rename

3. KEP work in 2021 (1.x, 1.y, 1.z):

<!--
In future, this will be generated from kubernetes/enhancements kep.yaml files
1. with SIG as owning-sig or in participating-sigs
2. listing 1.x, 1.y, or 1.z in milestones or in latest-milestone
-->

   - Stable
     - [1205-bound-service-account-tokens](https://github.com/kubernetes/enhancements/blob/master/keps/sig-auth/1205-bound-service-account-tokens/README.md) - 1.22.stable
     - [1393-oidc-discovery](https://github.com/kubernetes/enhancements/blob/master/keps/sig-auth/1393-oidc-discovery/README.md) - 1.21.stable
     - [2907-secrets-store-csi-driver](https://github.com/kubernetes/enhancements/blob/master/keps/sig-auth/2907-secrets-store-csi-driver/README.md) - 1.0.0.stable
     - [541-external-credential-providers](https://github.com/kubernetes/enhancements/blob/master/keps/sig-auth/541-external-credential-providers/README.md) - 1.22.stable
     - [1687-hierarchical-namespaces-subproject](https://github.com/kubernetes/enhancements/blob/master/keps/sig-auth/1687-hierarchical-namespaces-subproject/README.md) - stable
   - Beta
     - [2579-psp-replacement](https://github.com/kubernetes/enhancements/blob/master/keps/sig-auth/2579-psp-replacement/README.md) - 1.23.beta
     - [2784-csr-duration](https://github.com/kubernetes/enhancements/blob/master/keps/sig-auth/2784-csr-duration/README.md) - 1.22.beta
   - Alpha
   - Pre-alpha

## Project health

1. What areas and/or subprojects does your group need the most help with?
   Any areas with 2 or fewer OWNERs? (link to more details)

   The [Needs KEP / release work #sig-auth](https://docs.google.com/document/d/1sY8fRyRtk4eG9R439z5ao5i9bFuuxilS03XaNlqoni0/edit) document lists multiple areas that need help and some currently have volunteers working on them. 

2. What metrics/community health stats does your group care about and/or measure?

   - Based on devstats [Issue Velocity / Inactive Issues by SIG for 90 days or more](https://k8s.devstats.cncf.io/d/73/inactive-issues-by-sig?orgId=1&var-sigs=%22auth%22) at the time of writing this report, average is 9.  
   - Based on devstats [PR Velocity / Awaiting PRs by SIG for 90 days or more](https://k8s.devstats.cncf.io/d/70/awaiting-prs-by-sig?orgId=1&var-sigs=%22auth%22) at the time of writing this report, average is 38.  

3. Does your [CONTRIBUTING.md] help **new** contributors engage with your group specifically by pointing to activities or programs that provide useful context or allow easy participation?

   - Currently there is no onboarding or growth path. This is something we are working on and learning from other SIGs.

4. If your group has special training, requirements for reviewers/approvers, or processes beyond the general [contributor guide], does your [CONTRIBUTING.md] document those to help **existing** contributors grow throughout the [contributor ladder]?

   - Currently there is no onboarding or growth path. This is something we are working on and learning from other SIGs.

5. Does the group have contributors from multiple companies/affiliations?

   - Yes. Our chairs, leads, contributors, participants, and subproject owners are from various companies.

6. Are there ways end users/companies can contribute that they currently are not?
   If one of those ways is more full time support, what would they work on and why?

   - We need help with enhancing onboarding guide, pull request reviews, and areas listed in the [Needs KEP / release work #sig-auth](https://docs.google.com/document/d/1sY8fRyRtk4eG9R439z5ao5i9bFuuxilS03XaNlqoni0/edit) document.

## Membership

- Primary slack channel member count: 2463
- Primary mailing list member count: 470
- Primary meeting attendee count (estimated, if needed): 20 ~ 30
- Primary meeting participant count (estimated, if needed): 5 ~ 10
- Unique reviewers for SIG-owned packages: 11 <!-- in future, this will be generated from OWNERS files referenced from subprojects, expanded with OWNERS_ALIASES files https://github.com/kubernetes/kubernetes/blob/master/OWNERS_ALIASES -->
- Unique approvers for SIG-owned packages: 4 <!-- in future, this will be generated from OWNERS files referenced from subprojects, expanded with OWNERS_ALIASES files https://github.com/kubernetes/kubernetes/blob/master/OWNERS_ALIASES -->

Include any other ways you measure group membership

## Subprojects

<!--
In future, this will be generated from delta of sigs.yaml from $YYYY-01-01 to $YYYY-12-31
Manually visible via `git diff HEAD@{$YYYY-01-01} HEAD@{$YYYY-12-31} -- $sig-id/README.md`
-->

New in 2021:
- Added [kubernetes/pod-security-admission](https://github.com/kubernetes/pod-security-admission) under [policy-management](https://github.com/kubernetes/community/blob/master/sig-auth/README.md)

Continuing:
- All subprojects under https://github.com/kubernetes/community/blob/master/sig-auth/README.md#subprojects have continued.


## Working groups

<!--
In future, this will be generated from delta of sigs.yaml from $YYYY-01-01 to $YYYY-12-31
Manually visible via `git diff HEAD@{$YYYY-01-01} HEAD@{$YYYY-12-31} -- $sig-id/README.md`
-->

Continuing:
- All working groups under https://github.com/kubernetes/community/blob/master/sig-auth/README.md#working-groups have continued.

## Operational

Operational tasks in [sig-governance.md]:

- [x] [README.md] reviewed for accuracy and updated if needed
- [x] [CONTRIBUTING.md] reviewed for accuracy and updated if needed (or created if missing and your contributor steps and experience are different or more in-depth than the documentation listed in the general [contributor guide] and [devel] folder.)
- [x] Subprojects list and linked OWNERS files in [sigs.yaml] reviewed for accuracy and updated if needed
- [x] SIG leaders (chairs, tech leads, and subproject owners) in [sigs.yaml] are accurate and active, and updated if needed
- [x] Meeting notes and recordings for 2021 are linked from [README.md] and updated/uploaded if needed
- [x] Did you have community-wide updates in 2021 (e.g. community meetings, kubecon, or kubernetes-dev@ emails)? Links to email, slides, or recordings:
      - 2021 Kubecon NA Virtual - [PSP is Dead, Long Live PodSecurity](https://sched.co/lV9P) [session recording](https://youtu.be/yyr_cklZo3c)

[CONTRIBUTING.md]: https://git.k8s.io/community/sig-auth/CONTRIBUTING.md
[contributor ladder]: https://git.k8s.io/community/community-membership.md
[sig-governance.md]: https://git.k8s.io/community/committee-steering/governance/sig-governance.md
[README.md]: https://git.k8s.io/community/sig-auth/README.md
[sigs.yaml]: https://git.k8s.io/community/sigs.yaml
[contributor guide]: https://git.k8s.io/community/contributors/guide/README.md
[devel]: https://git.k8s.io/community/contributors/devel/README.md

