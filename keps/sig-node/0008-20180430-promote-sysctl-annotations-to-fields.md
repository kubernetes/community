---
kep-number: 8
title: Protomote sysctl annotations to fields
authors:
  - "@ingvagabund"
owning-sig: sig-node
participating-sigs:
  - sig-auth
reviewers:
  - "@sjenning"
  - "@derekwaynecarr"
approvers:
  - "@sjenning "
  - "@derekwaynecarr"
editor:
creation-date: 2018-04-30
last-updated: 2018-05-02
status: provisional
see-also:
replaces:
superseded-by:
---

# Promote sysctl annotations to fields

## Table of Contents

* [Promote sysctl annotations to fields](#promote-sysctl-annotations-to-fields)
   * [Table of Contents](#table-of-contents)
   * [Summary](#summary)
   * [Motivation](#motivation)
      * [Promote annotations to fields](#promote-annotations-to-fields)
      * [Promote --experimental-allowed-unsafe-sysctls kubelet flag to kubelet config api option](#promote---experimental-allowed-unsafe-sysctls-kubelet-flag-to-kubelet-config-api-option)
      * [Gate the feature](#gate-the-feature)
   * [Proposal](#proposal)
      * [User Stories](#user-stories)
      * [Implementation Details/Notes/Constraints](#implementation-detailsnotesconstraints)
      * [Risks and Mitigations](#risks-and-mitigations)
   * [Graduation Criteria](#graduation-criteria)
   * [Implementation History](#implementation-history)

## Summary

Setting the `sysctl` parameters through annotations provided a successful story
for defining better constraints of running applications.
The `sysctl` feature has been tested by a number of people without any serious
complaints. Promoting the annotations to fields (i.e. to beta) is another step in making the
`sysctl` feature closer towards the stable API.

Currently, the `sysctl` provides `security.alpha.kubernetes.io/sysctls` and `security.alpha.kubernetes.io/unsafe-sysctls` annotations that can be used
in the following way:
  ```yaml
  apiVersion: v1
  kind: Pod
  metadata:
    name: sysctl-example
    annotations:
      security.alpha.kubernetes.io/sysctls: kernel.shm_rmid_forced=1
      security.alpha.kubernetes.io/unsafe-sysctls: net.ipv4.route.min_pmtu=1000,kernel.msgmax=1 2 3
  spec:
    ...
  ```

  The goal is to transition into native fields on pods:

  ```yaml
  apiVersion: v1
  kind: Pod
  metadata:
    name: sysctl-example
  spec:
    securityContext:
      sysctls:
      - name: kernel.shm_rmid_forced
        value: 1
      - name: net.ipv4.route.min_pmtu
        value: 1000
        unsafe: true
      - name: kernel.msgmax
        value: "1 2 3"
        unsafe: true
    ...
  ```

The `sysctl` design document with more details and rationals is available at [design-proposals/node/sysctl.md](https://github.com/kubernetes/community/blob/master/contributors/design-proposals/node/sysctl.md#pod-api-changes)

## Motivation

As mentioned in [contributors/devel/api_changes.md#alpha-field-in-existing-api-version](https://github.com/kubernetes/community/blob/master/contributors/devel/api_changes.md#alpha-field-in-existing-api-version):

> Previously, annotations were used for experimental alpha features, but are no longer recommended for several reasons:
>
>    They expose the cluster to "time-bomb" data added as unstructured annotations against an earlier API server (https://issue.k8s.io/30819)
>    They cannot be migrated to first-class fields in the same API version (see the issues with representing a single value in multiple places in backward compatibility gotchas)
>
> The preferred approach adds an alpha field to the existing object, and ensures it is disabled by default:
>
> ...

The annotations as a means to set `sysctl` are no longer necessary.
The original intent of annotations was to provide additional description of Kubernetes
objects through metadata.
It's time to separate the ability to annotate from the ability to change sysctls settings
so a cluster operator can elevate the distinction between experimental and supported usage
of the feature.

### Promote annotations to fields

* Introduce native `sysctl` fields in pods through `spec.securityContext.sysctl` field as:

  ```yaml
  sysctl:
  - name: SYSCTL_PATH_NAME
    value: SYSCTL_PATH_VALUE
    unsafe: true    # optional field
  ```

* Introduce native `sysctl` fields in [PSP](https://kubernetes.io/docs/concepts/policy/pod-security-policy/) as:

  ```yaml
  apiVersion: v1
  kind: PodSecurityPolicy
  metadata:
    name: psp-example
  spec:
    sysctls:
    - kernel.shmmax
    - kernel.shmall
    - net.*
  ```

  More examples at [design-proposals/node/sysctl.md#allowing-only-certain-sysctls](https://github.com/kubernetes/community/blob/master/contributors/design-proposals/node/sysctl.md#allowing-only-certain-sysctls)

### Promote `--experimental-allowed-unsafe-sysctls` kubelet flag to kubelet config api option

As there is no longer a need to consider the `sysctl` feature experimental,
the list of unsafe sysctls can be configured accordingly through:

```go
// KubeletConfiguration contains the configuration for the Kubelet
type KubeletConfiguration struct {
  ...
  // Whitelist of unsafe sysctls or unsafe sysctl patterns (ending in *).
  // Default: nil
  // +optional
  AllowedUnsafeSysctls []string `json:"allowedUnsafeSysctls,omitempty"`
}
```

Upstream issue: https://github.com/kubernetes/kubernetes/issues/61669

### Gate the feature

As the `sysctl` feature stabilizes, it's time to gate the feature [1] and enable it by default.

* Expected feature gate key: `Sysctls`
* Expected default value: `true`

With the `Sysctl` feature enabled, both sysctl fields in `Pod` and `PodSecurityPolicy`
and the whitelist of unsafed sysctls are acknowledged.
If disabled, the fields and the whitelist are just ignored.

[1] https://kubernetes.io/docs/reference/command-line-tools-reference/feature-gates/

## Proposal

This is where we get down to the nitty gritty of what the proposal actually is.

### User Stories

* As a cluster admin, I want to have `sysctl` feature versioned so I can assure backward compatibility
  and proper transformation between versioned to internal representation and back..
* As a cluster admin, I want to be confident the `sysctl` feature is stable enough and well supported so
  applications are properly isolated
* As a cluster admin, I want to be able to apply the `sysctl` constraints on the cluster level so
  I can define the default constraints for all pods.

### Implementation Details/Notes/Constraints

Extending `SecurityContext` struct with `Sysctls` field:

```go
// PodSecurityContext holds pod-level security attributes and common container settings.
// Some fields are also present in container.securityContext.  Field values of
// container.securityContext take precedence over field values of PodSecurityContext.
type PodSecurityContext struct {
    ...
    // Sysctls is a white list of allowed sysctls in a pod spec.
    Sysctls []Sysctl `json:"sysctls,omitempty"`
}
```

Extending `PodSecurityPolicySpec` struct with `Sysctls` field:

```go
// PodSecurityPolicySpec defines the policy enforced on sysctls.
type PodSecurityPolicySpec struct {
    ...
    // Sysctls is a white list of allowed sysctls in a pod spec.
    Sysctls []Sysctl `json:"sysctls,omitempty"`
}
```

Following steps in [devel/api_changes.md#alpha-field-in-existing-api-version](https://github.com/kubernetes/community/blob/master/contributors/devel/api_changes.md#alpha-field-in-existing-api-version)
during implementation.

Validation checks implemented as part of [#27180](https://github.com/kubernetes/kubernetes/pull/27180).

### Risks and Mitigations

We need to assure backward compatibility, i.e. object specifications with `sysctl` annotations
must still work after the graduation.

## Graduation Criteria

* API changes allowing to configure the pod-scoped `sysctl` via `spec.securityContext` field.
* API changes allowing to configure the cluster-scoped `sysctl` via `PodSecurityPolicy` object
* Promote `--experimental-allowed-unsafe-sysctls` kubelet flag to kubelet config api option
* feature gate enabled by default
* e2e tests

## Implementation History

The `sysctl` feature is tracked as part of [features#34](https://github.com/kubernetes/features/issues/34).
This is one of the goals to promote the annotations to fields.
