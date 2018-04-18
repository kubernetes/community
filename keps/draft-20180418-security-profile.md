---
kep-number: draft-20180418
title: Security Profile
authors:
  - "@easeway"
owning-sig: sig-auth
status: provisional
reviewers:
  - "@tallclair"
  - "@ericchiang"
  - "@liggitt"
  - "@davidopp"
  - TBD
approvers:
  - "@tallclair"
  - "@ericchiang"
  - "@liggitt"
editor:
  - "@easeway"
creation-date: 2018-04-18
---
# Security Profile

Author: @easeway (yisuihu@google.com)

_This proposal is looking for sponsorship from sig-auth,
to create a public repo as [https://github.com/kubernetes-sigs/security-profiles](),
that community can help curate commonly used security profiles.
Details are described in [Release](#release) section below._

## Table of Contents

* [Table of Contents](#table-of-contents)
* [Summary](#summary)
* [Motivation](#motivation)
  * [Goals](#goals)
  * [Non-Goals](#non-goals)
* [Proposal](#proposal)
  * [Proposed Security Profiles](#proposed-security-profiles)
    * [default](#default)
    * [saas-multitenancy](#saas-multitenancy)
    * [more](#more)
* [User Experience](#user-experience)
  * [Use Cases](#use-cases)
* [Implementation](#implementation)
  * [The Security Profile](#the-security-profile)
    * [Bootstrapping Rules](#bootstrapping-rules)
    * [Runtime Rules](#runtime-rules)
    * [Effectiveness](#effectiveness)
    * [Switch Security Profiles](#switch-security-profiles)
    * [Enforcement Status](#enforcement-status)
    * [Cluster-scope Policy Objects](#cluster-scope-policy-objects)
    * [Namespace Population](#namespace-population)
  * [Customize and Extend](#customize-and-extend)
  * [Versioning Upgrade and Deprecation](#versioning-upgrade-and-deprecation)
    * [Release](#release)
    * [Version Schema](#version-schema)
    * [Upgrade](#upgrade)
    * [Deprecation](#deprecation)
  * [Work with Config-as-Code Systems](#work-with-config-as-code-systems)
* [References](#references)
* [Appendix](#Appendix)
  * [Predefined Rules](#predefined-rules)
    * [Bootstrapping Rule Definitions](#bootstrapping-rule-definitions)
      * [Addons](#addons)
      * [Command line flags](#command-line-flags)
      * [Admission Controls](#admission-controls)
      * [Authorization Mode](#authorization-mode)
      * [Feature Gates](#feature-gates)
      * [Runtime Configs](#runtime-configs)
      * [File Permissions](#file-permissions)
    * [Runtime Rule Definitions](#runtime-rule-definitions)
      * [Namespace-scope Policy Objects](#namespace-scope-policy-objects)
      * [Node Taints](#node-taints)
  * [Proposed Security Profile Examples](#proposed-security-profile-examples)
    * [The default profile](#the-default-profile)
    * [The saas-multitenancy profile](#the-saas-multitenancy-profile)

## Summary

SecurityProfile is a way of grouping a set of k8s control plane component flag settings,
templates for cluster- and namespace-scope policy objects, into a named, versioned bundle.
When a cluster admin specifies a security profile at cluster creation time,
they can be assured that the cluster will provide the security and isolation properties they need,
without them having to understand the low-level configuration used to achieve it.
Security Profile provides optional, additive functionality on top of core Kubernetes.

## Motivation

Configuring Kubernetes securely is difficult, because of:

- Management complexity
  - Complicated configuration of control plane components' security-relevant command line flags, correctness of cert/key/config file permissions;
  - Requires deep understanding of Kubernetes policy objects like RBAC, PodSecurityPolicy, NetworkPolicy, ResourceQuota etc;
  - No support for automatic provisioning of namespaced policy objects upon namespace creation;
  - Requires of domain knowledge about security;
  - Requires keeping up with fast evolving Kubernetes security/isolation features, and update configurations accordingly;
- Extra efforts building in-house solutions
  - No standard/common ways to solve the problems;
  - It’s very difficult to make in-house solutions sharable/reusable, as a result, efforts are duplicated solving same/similar problems.

### Goals

This design presents the proposal of Security Profile for Kubernetes and demonstrates

- The particular problems the proposal is going to solve
- The impact on user experience using Kubernetes
- The details how the proposal solves the problem

### Non-Goals

The following are not covered in this proposal

- Policy objects defined for a particular Kubernetes cluster
- How bootstrapping rules are enforced
- Verification mechanisms
- Continuously detect rule violations
- Continuously fix rule violations
- Conformance tests
- The details integration with Config-as-Code systems

## Proposal

_Security Profile_ defines a common and portable way to describe how to
secure a Kubernetes cluster, with content covering

- Cluster bootstrapping configuration, like command line flags, file permissions, etc
- Cluster-scoped policy objects, like PodSecurityPolicy
- Namespace-scoped policy objects, like NetworkPolicy

For example a _kick-the-tires_ security profile:

```yaml
apiVersion: extensions/v1beta1
kind: PodSecurityPolicy
metadata:
  name: kickthetires-1.0-0
spec:
  seLinux:
    rule: RunAsAny
  supplementalGroups:
    rule: RunAsAny
  runAsUser:
    rule: RunAsAny
  fsGroup:
    rule: RunAsAny
---
apiVersion: profile.security.k8s.io/v1alpha1
kind: NamespaceTemplate
metadata:
  name: kickthetires-1.0-0
spec:
  namespaces:
    selector:
      matchExpressions:
        - key: profile.security.k8s.io/opt-out
          operator: DoesNotExist
    excludes:
    - kube-system
  templates:
  - apiVersion: rbac.authorizatoin.k8s.io/v1
    kind: Role
    metadata:
      name: use-psp
    rules:
    - apiGroups: ['policy']
      resources: ['podsecuritypolicy']
      verbs: ['use']
      resourceNames: ['kickthetires-1.0-0']
    - apiVersion: rbac.authorization.k8s.io/v1
      kind: RoleBinding
      metadata:
        name: use-psp-binding
      roleRef:
        apiGroup: rbac.authorization.k8s.io/v1
        kind: Role
        name: use-psp
      subjects:
      - kind: Group
        apiGroup: rbac.authorization.k8s.io/v1
        name: system:serviceaccounts
---
apiVersion: profile.security.k8s.io/v1alpha1
kind: SecurityProfile
metadata:
  name: kickthetires-1.0-0
spec:
  bootstrapping-rules:
  - name: flags
    spec:
      apiserver:
        anonymous-auth:
          expect: unset
        basic-auth-file:
          expect: unset
  - name: admission-controls
    spec:
      excludes:
      - AlwaysAdmit
  - name: authorization-mode
    spec:
      excludes:
      - AlwaysAllow
  runtime-rules:
  - apiVersion: profile.security.k8s.io/v1alpha1
    kind: NamespaceTemplate
    name: kickthetires-1.0-0
```

A few curated security profiles can be defined to cover most common use cases and shared in the community, and helps use of Kubernetes:

- Simplicity bringing up a secured cluster -- when a cluster admin specifies a security profile at cluster creation time,
they can be assured that the cluster will provide the security and isolation properties they need,
without them having to understand the low-level configuration used to achieve it.
- The profile is sharable and reusable as it’s defined in a portable way
- Enforced automatically and continuously

Additionally, a Security Profile

- Is versioned, to allow seamless upgrade of the cluster and Security Profile
- Is highly customizable for particular needs
- Is highly extendable with extra features

An example if using security profile is

```
kubeadm init --security-profile=default-1.0-0
```

The command above will configure a secure cluster according to security profile `default-1.0-0`.

As security profiles cover common/generic use cases,
they are portable and reusable in most Kubernetes deployments,
They are **NOT** defining policies specific to a particular Kubernetes deployment, like:

- ResourceQuota
- RBAC rules
- Rules in NetworkPolicy for particular workloads

The proposal also works with Config-as-Code systems,
see the [section](#work-with-config-as-code-systems) below for details.

### Proposed Security Profiles

#### default

Full example: [default-1.0-0](#the-default-profile)

Minimum security practice to be compatible with most workloads.
Best fit for sharing the cluster with trusted workloads.

Enforcement:

- etcd
  - TLS required for communication
  - Explicitly specify `--wal-dir`, and `--max-wals=0`
- kube-apiserver
  - TLS required for communication
  - Audit logging must be explicitly configured
  - No anonymous auth, basic auth (unset `--anonymous-auth`, `--basic-auth-file`)
  - Required authorization-mode: Node, RBAC
  - Excluded authorization-mode: ABAC, AlwaysAllow
  - Disable insecure port (`--insecure-port=0`, unset `--insecure-bind-address`)
  - No auth token file (unset `--token-auth-file`)
  - Disable profiling (unset `--profiling`)
  - Request timeout is explicitly set
  - Use service account credentials (`--use-service-account-credentials=true`)
  - Require an encryption provider (`--experimental-encryption-provider-config`)
  - Required admission controls: AlwaysPullImages, DenyEscalatingExec, NamespaceLifecycle, ResourceQuota, PodSecurityPolicy, ServiceAccount, NodeRestriction, EventRateLimit
  - Excluded admission controls: AlwaysAdmit
  - Required feature gates to be ON: AdvancedAuditing, RotateKubeletServerCertificate, RotateKubeletClientCertificate
- controller-manager
  - Explicitly set `--terminated-pod-gc-threshold`
- kubelet
  - Disable hostname override (unset `--hostname-override`)
  - No privileged
- file-permissions
  - ... # correct setup of cert/key/config file permissions
- NodeTaints
  - No schedule on master

#### saas-multitenancy

Full Example: [saas-multitenancy-1.0-0](#the-saas-multitenancy-profile)

For SaaS scenarios which allows only plain workloads without the need to gain privileged access.

Enforcement: in addition to `default`

- PodSecurityPolicy
  - No privileged pods
  - Not allow privilege escalation
  - AppArmor profiles: docker/default, runtime/default
  - Seccomp profiles: docker/default, runtime/default
  - Drop ALL capabilities
  - No host namespaces (network, IPC, PID)
  - Limit volume types to
    - configMap
    - downwardAPI
    - emptyDir
    - nfs
    - persistentVolumeClaim
    - projected
    - secret
- NetworkPolicy
  - Block all ingress by default
  - Block all egress by default

#### more

Need help from community to curate commonly used security profiles.

## User Experience

### Use Cases

#### A small startup wants secure Kubernetes cluster for CI/CD on public cloud

Persona

- Fred, the head of operations in a small start-up, not an expert in Security

Fred has experience operating Kubernetes.
He decided to provision Kubernetes clusters on public cloud quickly for
dev/test, staging and production environment.
He wants a secure configuration without having to think about the details.
He wants the same configuration in all clusters so there are no surprising
changes in behavior when an application is moved from dev/test to staging/production.

He moves on creating Kubernetes clusters on the selected public cloud with CLI

```
cloud-cli kubernetes create dev-cluster --security-profile=default-1.0-0
cloud-cli kubernetes create test-cluster --security-profile=default-1.0-0
cloud-cli kubernetes create staging-cluster --security-profile=default-1.0-0
cloud-cli kubernetes create production-cluster --security-profile=defalut-1.0-0
```

#### An organization wants to provide Kubernetes as a Service, either to internal developers or as a hosted service to external users

Persona

- Fran, a Kubernetes cluster as a service operation manager

Fran is running a multi-tenant Kubernetes cluster as a service.
Because the cluster is shared by many users,
she wants strong isolation between namespaces,
and tight restrictions on what pods running on the cluster are allowed to do.
For example, host networking, priviledged pods and cross-namespace network traffic should be forbidden.
She also wants the users to provision namespaces by themselves via the provided portal or CLI where Kubernetes API will be used in the backend,
without going through an approval process.
Over the time, she may want to change policies over all user namespaces,
for example, change resource quota.
She's looking for some solution as the esitmation of building something in-house involving significant effort,
and find security profile `saas-multitenancy-1.0-0` satisfies the needs.

She creates a Kubernetes cluster with security profile `saas-multitenancy-1.0-0`
Later on, she defines ResourceQuota as a template and
use security profile's namespace population mechanism to
populate the ResourceQuota object in users' namespaces.

## Implementation

### The Security Profile
A _Security Profile_ is a set of Kubernetes manifest files, defining objects:

- The Custom Resource of kind `SecurityProfile` containing rules for:
  - Cluster bootstrapping configurations
  - Namespace-scope policy object definitions
- Cluster-scope object definitions
- The `Deployment` objects to deploy controllers to enforce the rules
  - Controllers watching namespaces and creates policy objects
  - External admission control webhook
  - Related configuration objects

Please refer to [default-1.0-0](#the-default-profile) as an example.

The Custom Resource `Security Profile` defines two sets of rules: bootstrapping rules and runtime rules

#### Bootstrapping Rules

The bootstrapping rules provide instructions for Kubernetes deployment tools
to set up configurations correctly during cluster creating, upgrading/downgrading, including:

- Command line flags for etcd, API server, controller manager, scheduler and kubelet
- Locations and permissions of certificates, keys, configuration files
- Security-relevant Operating System settings, for example
  - Configurations for SELinux, AppArmor, seccomp
  - Rules in iptables

All of them are described as actionable rules which contains:

- `name`: the unique name which identifies _a thing_ to be enforced;
- `spec`: describes the desired state when _the thing_ identified by `name` is enforced.

As security profiles are portable,
the `name` must be globally unique across all security profiles,
all clusters, and specific to define what to be enforced.
The schema of `spec` is defined by `name`. When `name` is defined,
the schema of `spec` is finalized, forever.
Further change of schema will require a new unique name.

With rules predefined,
a security profile aware Kubernetes deployment tool should be able to
download and inspect the selected security profile before creating/upgrading a cluster,
and setup command line flags, config files properly according to the rules during cluster creation/upgrade.

These rules can also be interpreted as a set of check conditions for tools to validate if the cluster has been setup properly.

Please see [Predefined Rules](#predefined-rules) as a reference of the rules.

As currently, there's no way (or limited) to inspect the host operating system from inside the cluster,
these rules are only enforced by the deployers (or cluster lifecycle management system).

#### Runtime Rules
Runtime rules are enforced after the cluster is bootstrapped,
by _enforcers_ (the controllers deployed with security profiles in the cluster).
Each rule is defined as an individual Custom Resource, that helps the enforcers define strong typed schemas.
In a security profile, a rule is a simple reference to a custom resource.

In the scope of current proposal, the runtime rules will define policy objects to be populated in namespace,
and the mechanism is required:

- make sure these policy objects are securely created when a new namespace is created
- update policy objects in existing namespaces when the current security profile or the rule in current security profile is changed

and the following is currently out-of-scope:

- existing namespaces
- detecting rule violations (missing of policy objects, or unexpected change of policy objects)
- auto fix rule violations

_Alternative design_

Runtime rules can also be defined similarly to bootstrapping rules, in a schemaless way.
However it loses the support from Kubernetes API to perform schema validation and other functionality,
even though it's currently very limited for CRDs, in future it may become better.

#### Effectiveness
Multiple security profiles can exist in the same cluster at the same time, like `default-1.0-0`, `saas-multitenancy-1.0-0`, etc.
but only one can be active (being enforced with `runtime-rules`, not `bootstrapping-rules`).
To specify which security profile is effective,
the cluster admin creates a singleton cluster-scope custom resource of kind `SecurityProfileSelector`, with predefined name `current`:

```yaml
apiVersion: profile.security.k8s.io/v1alpha1
kind: SecurityProfileSelector
metadata:
  name: current
spec:
  enforce: default-1.0-0
```

#### Switch Security Profiles

Switching between security profiles is easy by updating `SecurityProfileSelector`.
After the switch, some of the objects created by the previous security profile may no longer be desired.
To take care of these objects properly,
a cleanup mechanism is introduced with help of labeling the objects when they are created:

```
profile.security.k8s.io/profile-name=default-1.0-0
profile.security.k8s.io/profile-rule=0
```

`profile-name` tracks the security profile,
and `profile-rule` tracks the index in `runtime-rules` (see example above).
After switch, the new enforcement controller is expected to behave:

- Create/patch policy objects with updated labels;
- Scanning and removing labeled objects of the same type but with mismatch values.

#### Enforcement Status
The controller/admission control enforcing a particular rule will be responsible to report the status of the enforcement.
The status is defined as

- `none`: no enforcers recognize this rule
- `needed`: enforcement is needed, but not performed yet
- `inprogress`: the enforcers are working to enforce the rule, not completed yet
- `enforced`: the rule is currently enforced
- `error`: error happened during enforcement, not enforced

The status of all rules in a security profile is aggregated in the `status` section of the security profile object,
so a command like `kubectl get securityprofile default-1.0-0 -o yaml` will be able to inspect that

```
apiVersion: profile.security.k8s.io/v1alpha1
kind: SecurityProfile
metadata:
  name: default-1.0-0
spec:
  ... spec of security profile
status:
  enforcement:
    - name: profile.security.k8s.io/v1alpha1, Kind=NamespaceTemplate, Name=default-1.0-0
      status: inprogress
      details:
        namespaces-not-yet-enforced:
        - namespace-1
        - namespace-2
    - name: profile.security.k8s.io/v1alpha1, Kind=NodeTaints, Name=default-1.0-0
      status: enforced
```

#### Cluster-scope Policy Objects
These are cluster-scope Kubernetes policy objects.
They are defined directly in the manifest files, and will be created when using `kubectl apply` (or via addon manager, or whatever mechanism),
together with other objects in the security profile.

#### Namespace Population
When a namespace is created, some policy objects must be created to satisfy security needs.
Security Profile defines these rules including templates of objects to be created inside a namespace.
To make it happen automatically, a controller is deployed to watch the creation of namespaces,
and when that happens, it will create policy objects defined in the rules.
For example, a NetworkPolicy, or a RoleBinding to _use_ a PodSecurityPolicy.

This mechanism can be extensively used to automatically populate policy objects defined by cluster admins into namespaces.
This is very useful to add ResourceQuota, Role, RoleBinding for a particular cluster
in addition to common policy objects defined in security profiles.

The details of this mechanism are demonstrated in a separate KEP document [Namespace Population](https://github.com/kubernetes/community/pull/2177)

### Customize and Extend

There are a few cases the cluster admin may want to customize or extend a security profile

1. Additional policy objects for particular requirements

Security profiles define common/generic security settings, not settings for a particular cluster,
like specific ClusterRole/Role backed by a particular identity source,
ResourceQuota with specific limits according to the use of the cluster.
These policy objects can be populated using any mechanism (like terraform, ansible, or third-party tools etc.).
Optionally, the namespace population mechanism (see above) from security profile can be extensively used to automatically create policy objects inside namespaces.

2. Customize Rules in Security Profile

If the rules doesn't satisfy the particular requirements, cluster admin can create a custom security profile by:

- Copy the whole security profile from an existing one
- Give it a new name
- Alter the rules
- Apply it using `kubectl apply`
- Select it in `SecurityProfileSelector`

Because of limitation of schema support in CRD, during update, the whole security profile should be provided.
Partial patch is not supported.

3. Adding Custom Features

Especially for cloud providers (or SaaS providers),
they may want to introduce additional rules specific to cloud/service infrastructure for special features.
They can extend the security profiles by

- Define unique rule names and spec schemas
- Implement custom controller/admission control and deploy in cluster to enforce the rules
- Copy the whole security profile from an existing one
- Give it a new name
- Add custom rules
- Apply it using `kubectl apply`
- Select it in `SecurityProfileSelector`

### Versioning Upgrade and Deprecation

#### Release
Security profiles are maintained as an individual project similarly to other Kubernetes open source projects, like heapster, dashboard etc.
It has its own release lifecycle, because of the specialty of security profiles

- The release frequency is low.
- It will roughly follow the new release of Kubernetes if security related features are added/changed
- A specific version of a security profile will support multiple versions of Kubernetes.

The repository of security profile is intended to include the following contents:

- Manifest files defining CRDs and CRs of security profiles
- Manifest files defining cluster-scope objects referenced by security profiles
- Source code of enforcement logic (controllers) for namespace population
- Conformance tests (TBD, can be here, or in the main Kubernetes repo)
- Utilities, e.g. helper CLI for Config-as-Code systems
- Other files/scripts facilitating the building/testing/releasing process of this repo
- Documentations

When a Security Profile release is shipped, 
a container image is built and published to a public registry,
containing all manifest files and program bits (controllers, webhooks etc):

- Manifest files defining CRD and CRs of security profiles
- Manifest files defining cluster-scope objects referenced by security profiles
- Manifest files defining enforcers that are not expected to already be installed in every cluster, e.g. objects to deploy external admission controllers, and their required configurations
- Binaries of controllers, webhooks etc (normally should be one binary containing all components)

The container image can be used as a Kubernetes addon, 
as the controller is able to self-install security profiles into the cluster.
all the manifest files are included in a .tar.gz file called security profile release bundle.
The bundle can be hosted in github release page.

To help tools to determine the right version of a security profile for a particular Kubernetes version,
a version matrix is maintained in the source repo.
The version matrix is a simple CSV file, defining:

- Compatibility between a versioned security profile and Kubernetes versions
- The security profile release bundle containing a versioned security profile
- Deprecation of security profiles

```
default-1.0-0,1.9,std-security-profiles-0,DEPRECATED
default-1.0-1,1.9,std-security-profiles-1,
saas-multitenancy-1.0-1,1.9,std-security-profiles-1,
default-1.0-1,1.10,std-security-profiles-1,
saas-multitenancy-1.0-1,1.10,std-security-profiles-1,
default-1.2-2,1.10,std-security-profiles-2,
saas-multitenancy-1.2-2,1.10,std-security-profiles-2,
```

The columns are

- security profile name and version
- supported kubernetes version
- security profile release bundle name
- remarks, including DEPRECATED

Each row records a single mapping between security profile, bundle, Kubernetes version, and remarks.

#### Version Schema

A security profile is versioned like `default-1.2-3`, where

- `default-1` reflects the user scenarios this security profile covers
- `.2` reflects the definitions of all rules and related objects
- `-3` reflects the version of enforcement logic: admission controls, controllers, etc, mapping to version of release bundle

Within a security profile release bundle, all resources defined in manifest files are labeled with bundle name. E.g.

```
profile.security.k8s.io/profile-bundle=std-security-profiles-3
```

#### Upgrade

The general flow of upgrading to a new Security Profile is performed by:

- Install a new version of security profile release bundle
- Select a new version of security profile in `SecurityProfileSelector`
- Uninstall the old security profile bundle

Uninstall is easy as resources from a bundle are labeled with bundle name.
With multiple bundles installed, it's possible multiple admission controls, controllers are running side-by-side.
With help of labeled bundle name, the admission controls, controllers will stop enforcement if the currently selected security profile is not from the same bundle.
With this mechanism, it's very simple and lightweight to switch back to the previous security profile if something unexpected happens with a new security profile.

When upgrading a cluster, it's possible the current security profile may not be supported on the new cluster version.
In this case, upgrade the security profile to a new version first, and then upgrade the cluster.

#### Deprecation

A security profile is deprecated when:

- Any bug found in enforcement logic (admission control, controllers etc)
- Improper/Incomplete rules in the profile
- Other security concerns

When a security profile is deprecated,
it should be explicitly marked in the version matrix so tools (cluster lifecycle management system) are able to detect ASAP and notify cluster admin for an upgrade.

NOTE: a security profile is not going to be deprecated if it no longer covers new security features in new Kubernetes versions.
In this case, this security profile is no longer supported in new Kubernetes versions, but still usable with old Kubernetes versions.

### Work with Config-as-Code Systems

A config-as-code system uses a repository with version control to manage all changes made to a Kubernetes cluster.
The repository contains a full view of policy objects to be applied to Kubernetes clusters as a single source of truth.
Changes must go through code-check-in process through the version control and validated via precheck-in hooks and CI/CD pipeline.
Such a system has advantage mitigating the risk of damaging the cluster by mistakes in configuration.

To work with config-as-code systems, security profiles are installed into the code repository, not the cluster.
Additional helper CLI tools are provided to perform lint/vet in the precheck-in hooks and CI/CD pipelines.

The details are covered by a separated doc [For Config-as-Code Systems](https://drive.google.com/open?id=1mtc6365k97d6rrDwQceTgi8In0NUlEIiA0ZgAmYrABs).

## References

- [Design Doc (Google Doc)](https://docs.google.com/document/d/1lFiRNDWgyoZWQfvQV2UYsH80CKW90jwSeBzWyssTf7g/edit?usp=sharing)
- [Presentation](https://docs.google.com/presentation/d/1PmRcgID9-TKX-0x3E3rkKNaijrXIe45KGTIKTLe5FC8/edit?usp=sharing)
- [For Config-as-Code Systems](https://drive.google.com/open?id=1mtc6365k97d6rrDwQceTgi8In0NUlEIiA0ZgAmYrABs)
- [Multi-tenant security/isolation configuration](https://docs.google.com/document/d/1jAcsC4sLgEV9__TdgJrMvPa3G73G62tFtMcKQgeIlHM/edit?usp=sharing)
- [Kubernetes policy extensibility mechanisms](https://docs.google.com/document/d/1XiQ8cac0rXqwbw8_tiWicN0R4t1-7rPwUUaok3O8vok/edit?usp=sharing)

## Appendix

### Predefined Rules

#### Bootstrapping Rule Definitions

##### Addons

Name: `addons`

Spec:

```
spec:
  required:
  - <addon1-name>
  excludes:
  - <addon2-name>
```

Criteria:

When an addon is specified as `required`, the addon MUST be installed on the cluster.
If specified as `excludes`, the addon MUST NOT be installed.

##### Command line flags

Name: `flags`

Spec:

```
spec:
  <service>:
    <flag>:
      expect: 'set' | 'unset' | not set
      values: ['value1', 'value2', ...]
      excludes: ['exc1', 'exc2', ...]

<service> is defined as: apiserver | controllermanager | scheduler | etcd
```

Crteria:

- `set`: the flag MUST be explicitly set to whatever value
- `unset`: the flag MUST NOT be set
- not set:
  - `values`: the value of the flag MUST be ONE OF those in the list
  - `excludes`: the value of the flag MUST NOT be ANY OF those in the list, the flag can be `unset`

##### Admission Controls

Name: `admission-controls`

Spec:

```
spec:
  required:
  - <admission-control1>
  - <admission-control2>
  excludes:
  - <admission-control-ex-1>
  - <admission-control-ex-2>
```

Criteria:

When listed in `required`, the admission control MUST be enabled;
when listed in `excludes`, the admission control MUST NOT be enabled.

##### Authorization Mode

Name: `authorization-mode`

Spec:

```
spec:
  required:
  - <authorizer1>
  - <authorizer2>
  excludes:
  - <ex-authorizer1>
  - <ex-authorizer2>
```

Criteria:

When listed in `required`, the authorizer MUST be included;
when listed in `excludes`, the authorizer MUST NOT be enabled.

##### Feature Gates

Name: `feature-gates`

Spec:

```
spec:
  required:
  - <gate1>
  - <gate2>
  excludes:
  - <ex-gate1>
  - <ex-gate2>
```

When listed in `required`, the feature gate MUST be included;
when listed in `excludes`, the feature gate MUST NOT be enabled.

##### Runtime Configs

Name: `runtime-configs`

Spec:

```
spec:
  required:
  - <config1>
  - <config2>
  excludes:
  - <ex-config1>
  - <ex-config2>
```

When listed in `required`, the runtime config MUST be included;
when listed in `excludes`, the runtime config MUST NOT be enabled.

##### File Permissions

Name: `file-permissions`

Spec:

```
spec:
  - locations: [master|node, ...]
    files:
    - symbol: <file_symbol>
      permission: '<oct-permission>'
      ownership: '<user>:<group>'

<file_symbol> are defined as
apiserver_pod_spec:
  the static pod manifest for API server
  e.g. /etc/kubernetes/manifests/kube-apiserver.yaml
controllermanager_pod_spec:
  the static pod manifest for controller manager
  e.g. /etc/kubernetes/manifests/controller-manager.yaml
scheduler_pod_spec:
  the static pod manifest for scheduler
  e.g. /etc/kubernetes/manifests/scheduler.yaml
etcd_pod_spec:
  the static pod manifest for etcd
  e.g. /etc/kubernetes/manifests/etcd.yaml
cni_config:
  the CNI configuration dir: /etc/cni
etcd_data_dir:
  path to etcd data directory
admin_conf:
  configuration file for cluster-admin
  e.g. /etc/kubernetes/admin.conf
scheduler_conf:
  configuration file for scheduler
  e.g. /etc/kubernetes/scheduler.conf
controllermanager_conf:
  configuration file for controller manager
  e.g. /etc/kubernetes/controller-manager.conf
kubelet_conf:
  configuration file for kubelet
  e.g. /etc/kubernetes/kubelet.conf
kubelet_service:
  kubelet service file
  e.g. (systemd) /usr/lib/systemd/system/kubelet.service
  or /etc/systemd/system/kubelet.service
proxy_kubeconfig:
  configuration file for kubeproxy
ca_cert:
  path to CA certificate file
```

Criteria:

`locations` specified the rules applies on `master`, or `node`, or both if not specified;
`symbol` specify which file/dir to enforce the permission and ownership.

#### Runtime Rule Definitions

##### Namespace-scope Policy Objects

CRD:

```yaml
apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  name: namespacetemplates.security.k8s.io
spec:
  group: security.k8s.io
  version: v1alpha1
  scope: Cluster
  names:
    plural: namespacetemplates
    singular: namespacetemplate
    kind: NamespaceTemplate
    shortNames:
    - nst
```

Spec:

```
spec:
  namespaces:
    selector: <label-selector>
    includes: [names]
    excludes: [names]
  templates:
  - <template1>
  - <template2>

'includes' and 'excludes' list plain names of namespaces as complementary to label selector.
The order applying the filters are labels -> excludes -> includes.
Wildcard '*' can be used in 'excludes' and 'includes' to match a name.
If 'includes' is empty, it selects all, otherwise it only selects the specified names.

'templates' defines a list of inline namespace-scope objects.
Note, 'metadata.namespace' will be filled in automatically, so don’t define in the rule.
```

Criteria:

The objects defined in `templates` will be created in namespaces matched by filters defined by `namespaces`.

##### Node Taints

CRD:

```yaml
apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  name: nodetaints.security.k8s.io
spec:
  group: security.k8s.io
  version: v1alpha1
  scope: Cluster
  names:
    plural: nodetaints
    singular: nodetaints
    kind: NodeTaints
```

Spec:

```
spec:
  - role: master|node
    selector: <nodeSelector>
    taints:
    - <taint>
```

Criteria:

Make sure the listed taints are tainted on selected nodes.

### Proposed Security Profile Examples

#### The default Profile

```yaml
apiVersion: profile.security.k8s.io/v1alpha1
kind: NamespaceTemplate
metadata:
  name: default-1.0
spec:
  templates:
    - apiVersion: rbac.authorization.k8s.io/v1
      kind: RoleBinding
      metadata:
        name: namespace-admin
      roleRef:
        apiGroup: rbac.authorization.k8s.io/v1
        kind: ClusterRole
        name: admin
      subjects:
      - kind: User
        apiGroup: rbac.authorization.k8s.io/v1
        name: $(CREATOR)
---
apiVersion: profile.security.k8s.io/v1alpha1
kind: NodeTaints
metadata:
  name: noschedule-on-master-1.0
spec:
  - role: master
    taints:
    - node-role.kubernetes.io/master=NoSchedule
---
apiVersion: profile.security.k8s.io/v1alpha1
kind: SecurityProfile
metadata:
  name: default-1.0
spec:
  bootstrapping-rules:
  - name: addons
    spec:
      dashboard: disabled
  - name: flags
    spec:
      etcd:
        auto-tls:
          excludes: ['true']
        cert-file:
          expect: set
        client-cert-auth:
          values: ['true']
        key-file:
          expect: set
        peer-auto-tls:
          excludes: ['true']
        peer-cert-file:
          expect: set
        peer-client-cert-auth:
          values: ['true']
        peer-key-file:
          expect: set
        wal-dir:
          expect: set
        max-wals:
          values: ['0']
      apiserver:
        anonymous-auth:
          expect: unset
        audit-log-path:
          expect: set
        audit-log-maxage:
          expect: set
        audit-log-maxbackup:
          expect: set
        audit-log-maxsize:
          expect: set
        basic-auth-file:
          expect: unset
        client-ca-file:
          expect: set
        etcd-ca-file:
          expect: set
        etcd-certfile:
          expect: set
        etcd-keyfile:
          expect: set
        experimental-encryption-provider-config:
          expect: set
        insecure-allow-any-token:
          expect: unset
        insecure-bind-address:
          expect: unset
        insecure-port:
          values: ['0']
        kubelet-certificate-authority:
          expect: set
        kubelet-client-certificate:
          expect: set
        kubelet-client-key:
          expect: set
        kubelet-https:
          values: ['true']
        profiling:
          values: ['false']
        request-timeout:
          expect: set
        secure-port:
          excludes: ['0']
        service-account-key-file:
          expect: set
        tls-cert-file:
          expect: set
        tls-private-key:
          expect: set
        token-auth-file:
          expect: unset
        use-service-account-credentials:
          values: ['true']
      controller-manager:
        terminated-pod-gc-threshold:
          expect: set
        service-account-private-key-file:
          expect: set
        root-ca-file:
          expect: set
      kubelet:
        hostname-override:
          expect: unset
        allow-privileged:
          values: ['false']
  - name: admission-controls
    spec:
      required:
      - AlwayPullImages
      - DenyEscalatingExec
      - SecurityContextDeny
      - NamespaceLifecycle
      - PodSecurityPolicy
      - ServiceAccount
      - NodeRestriction
      - EventRateLimit
      excludes:
      - AlwaysAdmit
  - name: authorization-mode
    spec:
      required:
      - Node
      - RBAC
      excludes:
      - AlwaysAllow
      - ABAC
  - name: feature-gates
    spec:
      required:
      - AdvancedAuditing
      - RotateKubeletServerCertificate
      - RotateKubeletClientCertificate
  - name: file-permissions
    spec:
    - locations: ['master']
      files:
      - symbol: apiserver_pod_spec
        permission: '0644'
        ownership: '0:0'
      - symbol: controllermanager_pod_spec
        permission: '0644'
        ownership: '0:0'
      - symbol: scheduler_pod_spec
        permission: '0644'
        ownership: '0:0'
      - symbol: etcd_pod_spec
        permission: '0644'
        ownership: '0:0'
      - symbol: cni_config
        permission: '0644'
        ownership: '0:0'
      - symbol: etcd_data_dir
        permission: '0700'
        ownership: 'etcd:etcd'
      - symbol: admin_conf
        permission: '0600'
        ownership: '0:0'
      - symbol: scheduler_conf
        permission: '0600'
        ownership: '0:0'
      - symbol: controllermanager_conf
        permission: '0600'
        ownership: '0:0'
    - locations: ['master', 'node']
      files:
      - symbol: kubelet_conf
        permission: '0600'
        ownership: '0:0'
      - symbol: kubelet_service
        permission: '0644'
        ownership: '0:0'
      - symbol: proxy_kubeconfig
        permission: '0600'
        ownership: '0:0'
      - symbol: ca_cert
        permission: '0644'
        ownership: '0:0'
  runtime-rules:
  - apiVersion: profile.security.k8s.io/v1alpha1
    kind: NamespaceTemplate
    name: default-1.0
  - apiVersion: profile.security.k8s.io/v1alpha1
    kind: NodeTaints
    name: noschedule-on-master-1.0
```

#### The saas-multitenancy Profile

```yaml
apiVersion: extensions/v1beta1
kind: PodSecurityPolicy
metadata:
  name: saas-multitenancy-1.0
  annotations:
    seccomp.security.alpha.kubernetes.io/allowedProfileNames: 'docker/default'
    apparmor.security.beta.kubernetes.io/allowedProfileNames: 'runtime/default'
    seccomp.security.alpha.kubernetes.io/defaultProfileName:  'docker/default'
    apparmor.security.beta.kubernetes.io/defaultProfileName:  'runtime/default'
spec:
  privileged: false
  allowPrivilegeEscalation: false
  requiredDropCapabilities:
  - ALL
  hostNetwork: false
  hostIPC: false
  hostPID: false
  seLinux:
    rule: RunAsAny
  supplementalGroups:
    rule: RunAsAny
  runAsUser:
    rule: RunAsAny
  fsGroup:
    rule: RunAsAny
  volumes:
  - configMap
  - downwardAPI
  - emptyDir
  - nfs
  - persistentVolumeClaim
  - projected
  - secret
---
apiVersion: profile.security.k8s.io/v1alpha1
kind: NamespaceTemplate
metadata:
  name: saas-multitenancy-1.0
spec:
  templates:
  - apiVersion: rbac.authorizatoin.k8s.io/v1
    kind: Role
    metadata:
      name: use-psp
    rules:
    - apiGroups: ['policy']
      resources: ['podsecuritypolicy']
      verbs: ['use']
      resourceNames: ['saas-multitenancy-1.0']
  - apiVersion: rbac.authorization.k8s.io/v1
    kind: RoleBinding
    metadata:
      name: use-psp-binding
    roleRef:
      apiGroup: rbac.authorization.k8s.io/v1
      kind: Role
      name: use-psp
    subjects:
    - kind: Group
      apiGroup: rbac.authorization.k8s.io/v1
      name: system:serviceaccounts
  - apiVersion: rbac.authorization.k8s.io/v1
    kind: RoleBinding
    metadata:
      name: namespace-admin
    roleRef:
      apiGroup: rbac.authorization.k8s.io/v1
      kind: ClusterRole
      name: admin
    subjects:
    - kind: User
      apiGroup: rbac.authorization.k8s.io/v1
      name: $(CREATOR)      
  - apiVersion: networking.k8s.io/v1
    kind: NetworkPolicy
    metadata:
      name: default
    spec:
      policyTypes:
      - Ingress
      - Egress
---
apiVersion: profile.security.k8s.io/v1alpha1
kind: NodeTaints
metadata:
  name: noschedule-on-master-1.0
spec:
  - role: master
    taints:
    - node-role.kubernetes.io/master=NoSchedule
---
apiVersion: profile.security.k8s.io/v1alpha1
Kind: SecurityProfile
metadata:
  name: saas-multitenancy-1.0
spec:
  bootstrapping-rules:
  - name: addons
    spec:
      excludes:
      - dashboard
  - name: flags
    spec:
      etcd:
        auto-tls:
          excludes: ['true']
        cert-file:
          expect: set
        client-cert-auth:
          values: ['true']
        key-file:
          expect: set
        peer-auto-tls:
          excludes: ['true']
        peer-cert-file:
          expect: set
        peer-client-cert-auth:
          values: ['true']
        peer-key-file:
          expect: set
        wal-dir:
          expect: set
        max-wals:
          values: ['0']
      apiserver:
        anonymous-auth:
          expect: unset
        audit-log-path:
          expect: set
        audit-log-maxage:
          expect: set
        audit-log-maxbackup:
          expect: set
        audit-log-maxsize:
          expect: set
        basic-auth-file:
          expect: unset
        client-ca-file:
          expect: set
        etcd-ca-file:
          expect: set
        etcd-certfile:
          expect: set
        etcd-keyfile:
          expect: set
        experimental-encryption-provider-config:
          expect: set
        insecure-allow-any-token:
          expect: unset
        insecure-bind-address:
          expect: unset
        insecure-port:
          values: ['0']
        kubelet-certificate-authority:
          expect: set
        kubelet-client-certificate:
          expect: set
        kubelet-client-key:
          expect: set
        kubelet-https:
          values: ['true']
        profiling:
          values: ['false']
        request-timeout:
          expect: set
        secure-port:
          excludes: ['0']
        service-account-key-file:
          expect: set
        tls-cert-file:
          expect: set
        tls-private-key:
          expect: set
        token-auth-file:
          expect: unset
        use-service-account-credentials:
          values: ['true']
        experimental-encryption-provider-config:
          expect: set
      controller-manager:
        terminated-pod-gc-threshold:
          expect: set
        service-account-private-key-file:
          expect: set
        root-ca-file:
          expect: set
      kubelet:
        hostname-override:
          expect: unset
        allow-privileged:
          values: ['false']
  - name: admission-controls
    spec:
      required:
      - AlwayPullImages
      - DenyEscalatingExec
      - SecurityContextDeny
      - NamespaceLifecycle
      - PodSecurityPolicy
      - ServiceAccount
      - NodeRestriction
      - EventRateLimit
      excludes:
      - AlwaysAdmit
  - name: authorization-mode
    spec:
      required:
      - Node
      - RBAC
      excludes:
      - AlwaysAllow
      - ABAC
  - name: feature-gates
    spec:
      required:
      - AdvancedAuditing
      - RotateKubeletServerCertificate
      - RotateKubeletClientCertificate
  - name: file-permissions
    spec:
    - locations: ['master']
      files:
      - symbol: apiserver_pod_spec
        permission: '0644'
        ownership: '0:0'
      - symbol: controllermanager_pod_spec
        permission: '0644'
        ownership: '0:0'
      - symbol: scheduler_pod_spec
        permission: '0644'
        ownership: '0:0'
      - symbol: etcd_pod_spec
        permission: '0644'
        ownership: '0:0'
      - symbol: cni_config
        permission: '0644'
        ownership: '0:0'
      - symbol: etcd_data_dir
        permission: '0700'
        ownership: 'etcd:etcd'
      - symbol: admin_conf
        permission: '0600'
        ownership: '0:0'
      - symbol: scheduler_conf
        permission: '0600'
        ownership: '0:0'
      - symbol: controllermanager_conf
        permission: '0600'
        ownership: '0:0'
    - locations: ['master', 'node']
      files:
      - symbol: kubelet_conf
        permission: '0600'
        ownership: '0:0'
      - symbol: kubelet_service
        permission: '0644'
        ownership: '0:0'
      - symbol: proxy_kubeconfig
        permission: '0600'
        ownership: '0:0'
      - symbol: ca_cert
        permission: '0644'
        ownership: '0:0'
  runtime-rules:
  - apiVersion: profile.security.k8s.io/v1alpha1
    kind: NamespaceTemplate
    name: saas-multitenancy-1.0
  - apiVersion: profile.security.k8s.io/v1alpha1
    kind: NodeTaints
    name: noschedule-on-master-1.0
```
