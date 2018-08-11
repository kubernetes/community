---
kep-number: draft-20180412
title: Kubeadm Config versioning
authors:
  - "@liztio"
owning-sig: sig-cluster-lifecycle
participating-sigs: []
reviewers:
  - "@timothysc"
approvers:
  - TBD
editor: TBD
creation-date: 2018-04-12
last-updated: 2018-04-12
status: draft
see-also: []
replaces: []
superseded-by: []
---

# Kubeadm Config Versioning

## Table of Contents

A table of contents is helpful for quickly jumping to sections of a KEP and for highlighting any additional information provided beyond the standard KEP template.

<!-- markdown-toc start - Don't edit this section. Run M-x markdown-toc-refresh-toc -->
**Table of Contents**

- [Kubeadm Config to Beta](#kubeadm-config-to-beta)
    - [Table of Contents](#table-of-contents)
    - [Summary](#summary)
    - [Motivation](#motivation)
        - [Goals](#goals)
        - [Non-Goals](#non-goals)
    - [Proposal](#proposal)
        - [User Stories [optional]](#user-stories-optional)
            - [As a user upgrading with Kubeadm, I want the upgrade process to not fail with unfamiliar configuration.](#as-a-user-upgrading-with-kubeadm-i-want-the-upgrade-process-to-not-fail-with-unfamiliar-configuration)
            - [As a infrastructure system using kubeadm, I want to be able to write configuration files that always work.](#as-a-infrastructure-system-using-kubeadm-i-want-to-be-able-to-write-configuration-files-that-always-work)
        - [Implementation Details/Notes/Constraints](#implementation-detailsnotesconstraints)
        - [Risks and Mitigations](#risks-and-mitigations)
    - [Graduation Criteria](#graduation-criteria)
    - [Implementation History](#implementation-history)
    - [Alternatives](#alternatives)

<!-- markdown-toc end -->

## Summary

Kubeadm uses MasterConfiguraton for two distinct but similar operations: Initialising a new cluster and upgrading an existing cluster. 
The former is typically created by hand by an administrator. 
It is stored on disk and passed to `kubeadm init` via command line flag.
The latter is produced by kubeadm using supplied configuration files, command line options, and internal defaults.
It will be stored in a ConfigMap so upgrade operations can find. 

Right now the configuration format is unversioned.
This means configuration file formats can change between kubeadm versions and there's no safe way to update the configuration format.

We propose a stable versioning of this configuration, `v1alpha2` and eventually `v1beta1`. 
Version information will be _mandatory_ going forward, both for user-generated configuration files and machine-generated configuration maps.

There as an [existing document][config] describing current Kubernetes best practices around component configuration.

[config]: https://docs.google.com/document/d/1FdaEJUEh091qf5B98HM6_8MS764iXrxxigNIdwHYW9c/edit#heading=h.nlhhig66a0v6

## Motivation

After 1.10.0, we discovered a bug in the upgrade process. 
The `MasterConfiguraton` embedded a [struct that had changed][proxyconfig], which caused a backwards-incompatible change to the configuration format. 
This caused `kubeadm upgrade` to fail, because a newer version of kubeadm was attempting to deserialise an older version of the struct.

Because the configuration is often written and read by different versions of kubeadm compiled by different versions of kubernetes, 
it's very important for this configuration file to be well-versioned. 

[proxyconfig]: https://github.com/kubernetes/kubernetes/commit/57071d85ee2c27332390f0983f42f43d89821961

### Goals

* kubeadm init fails if a configuration file isn't versioned
* the config map written out contains a version
* the configuration struct does not embed any other structs
* existing configuration files are converted on upgrade to a known, stable version
* structs should be sparsely populated
* all structs should have reasonable defaults so an empty config is still sensible

### Non-Goals

* kubeadm is able to read and write configuration files for older and newer versions of kubernetes than it was compiled with
* substantially changing the schema of the `MasterConfiguration`

## Proposal

The concrete proposal is as follows.

1. Immediately start writing Kind and Version information into the `MasterConfiguraton` struct.
2. Define the previous (1.9) version of the struct as `v1alpha1`.
3. Duplicate the KubeProxyConfig struct that caused the schema change, adding the old version to the `v1alpha1` struct.
3. Create a new `v1alpha2` directory mirroring the existing [`v1alpha1`][v1alpha1], which matches the 1.10 schema. 
   This version need not duplicate the file as well.
2. Warn users if their configuration files do not have a version and kind
4. Use [apimachinery's conversion][conversion] library to design migrations from the old (v1alpha1) versions to the new (v1alpha2) versions
5. Determine the changes for v1beta1
6. With v1beta1, enforce presence of version numbers in config files and ConfigMaps, erroring if not present.

[conversion]: https://godoc.org/k8s.io/apimachinery/pkg/conversion
[v1alpha1]: https://github.com/kubernetes/kubernetes/tree/d7d4381961f4eb2a4b581160707feb55731e324e/cmd/kubeadm/app/apis/kubeadm 

### User Stories [optional]

#### As a user upgrading with Kubeadm, I want the upgrade process to not fail with unfamiliar configuration.

In the past, the haphazard nature of the versioning system has meant it was hard to provide strong guarantees between versions.
Implementing strong version guarantees mean any given configuration generated in the past by kubeadm will work with a future version of kubeadm. 
Deprecations can happen in the future in well-regulated ways.

#### As a infrastructure system using kubeadm, I want to be able to write configuration files that always work.

Having a configuration file that changes without notice makes it very difficult to write software that integrates with kubeadm. 
By providing strong version guarantees, we can guarantee that the files these tools produce will work with a given version of kubeadm.

### Implementation Details/Notes/Constraints

The incident that caused the breakage in alpha wasn't a field changed it Kubeadm, it was a struct [referenced][struct] inside the `MasterConfiguration` struct.
By completely owning our own configuration, changes in the rest of the project can't unknowingly affect us.
When we do need to interface with the rest of the project, we will do so explicitly in code and be protected by the compiler.

[struct]: https://github.com/kubernetes/kubernetes/blob/d7d4381961f4eb2a4b581160707feb55731e324e/cmd/kubeadm/app/apis/kubeadm/v1alpha1/types.go#L285

### Risks and Mitigations

Moving to a strongly versioned configuration from a weakly versioned one must be done carefully so as not break kubeadm for existing users. 
We can start requiring versions of the existing `v1alpha1` format, issuing warnings to users when Version and Kind aren't present.
These fields can be used today, they're simply ignored.
In the future, we could require them, and transition to using `v1alpha1`.

## Graduation Criteria

This KEP can be considered complete once all currently supported versions of Kubeadm write out `v1beta1`-version structs.

## Implementation History

## Alternatives

Rather than creating our own copies of all structs in the `MasterConfiguration` struct, we could instead continue embedding the structs.
To provide our guarantees, we would have to invest a lot more in automated testing for upgrades.
