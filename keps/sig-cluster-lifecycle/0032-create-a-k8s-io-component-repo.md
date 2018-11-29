---
kep-number: 32
title: Create a `k8s.io/component` repo
status: implementable
authors:
  - "@luxas"
  - "@sttts"
owning-sig: sig-cluster-lifecycle
participating-sigs:
  - sig-api-machinery
  - sig-cloud-provider
reviewers:
  - "@thockin"
  - "@jbeda"
  - "@bgrant0607"
  - "@smarterclayton"
  - "@liggitt"
  - "@lavalamp"
  - "@andrewsykim"
  - "@cblecker"
approvers:
  - "@thockin"
  - "@jbeda"
  - "@bgrant0607"
  - "@smarterclayton"
editor:
  name: "@luxas"
creation-date: 2018-11-27
last-updated: 2018-11-27
---

# Create a `k8s.io/component` repo

**How we can consolidate the look and feel of core and non-core components with regards to ComponentConfiguration, flag handling, and common functionality with a new repository**

## Table of Contents

- [Create a `k8s.io/component` repo](#create-a--k8sio-component--repo)
  - [Table of Contents](#table-of-contents)
  - [Abstract](#abstract)
    - [History and Motivation](#history-and-motivation)
    - ["Component" definition](#-component--definition)
    - [Goals](#goals)
    - [Success metrics](#success-metrics)
    - [Non-goals](#non-goals)
    - [Related proposals / references](#related-proposals---references)
  - [Proposal](#proposal)
    - [Part 1: ComponentConfig](#part-1--componentconfig)
      - [Standardized encoding/decoding](#standardized-encoding-decoding)
      - [Testing helper methods](#testing-helper-methods)
      - [Generate OpenAPI specifications](#generate-openapi-specifications)
    - [Part 2: Command building / flag parsing](#part-2--command-building---flag-parsing)
      - [Wrapper around cobra.Command](#wrapper-around-cobracommand)
      - [Flag precedence over config file](#flag-precedence-over-config-file)
      - [Standardized logging](#standardized-logging)
    - [Part 3: HTTPS serving](#part-3--https-serving)
      - [Common endpoints](#common-endpoints)
      - [Standardized authentication / authorization](#standardized-authentication---authorization)
    - [Part 4: Sample implementation in k8s.io/sample-component](#part-4--sample-implementation-in-k8sio-sample-component)
    - [Code structure](#code-structure)
    - [Timeframe and Implementation Order](#timeframe-and-implementation-order)
    - [OWNERS file for new packages](#owners-file-for-new-packages)

## Abstract

The proposal is about preparing the Kubernetes core package structure in a way that all core component can share common code around

- ComponentConfig implementation
- flag and command handling
- HTTPS serving
- delegated authn/z
- logging.

Today this code is spread over the k8s.io/kubernetes repository, staging repository or pieces of code are in locations they don't belong to (example: k8s.io/apiserver/pkg/util/logs is the for general logging, totally independent of API servers). We miss a repository far enough in the dependency hierarchy for code that is or should be common among core Kubernetes component (neither k8s.io/apiserver, k8s.io/apimachinery or k8s.io/client-go are right for that).

### History and Motivation

By this time in the Kubernetes development, we know pretty well how we want a Kubernetes component to work, function, and look. But achieving this requires a fair amount of more or less advanced code. As we scale the ecosystem, and evolve Kubernetes to work more as a kernel, it's increasingly important to make writing extensions and custom Kubernetes-aware components relatively easy. As it stands today, this is anything but straightforward. In fact, even the in-core components diverge in terms of configurability (Can it be declaratively configured? Do flag names follow a consistent pattern? Are configuration sources consistently merged?), common functionality (Does it support the common "/version," "/healthz," "/configz," "/pprof," and "/metrics" endpoints? Does it utilize Kubernetes' authentication/authorization mechanisms? Does it write logs in a consistent manner? Does it handle signals as others do?), and testability (Do the internal configuration structs set up correctly to conform with the Kubernetes API machinery, and have roundtrip, defaulting, validation unit tests in place? Does it merge flags and the config file correctly? Is the logging mechanism set up in a testable manner? Can it be verified that the HTTP server has the standard endpoints registered and working? Can it be verified that authentication and authorization is set up correctly?).

This document proposes to create a new Kubernetes staging repository with minimal dependencies (_k8s.io/apimachinery_, _k8s.io/client-go_, and _k8s.io/api_) and good documentation on how to write a Kubernetes-aware component that follows best practices. The code and best practices in this repo would be used by all the core components as well. Unifying the core components would be great progress in terms of the internal code structure, capabilities, and test coverage. Most significantly, this would lead to an adoption of ComponentConfig for all internal components as both a "side effect" and a desired outcome, which is long time overdue.

The current inconsistency is a headache for many Kubernetes developers, and confusing for end users. Implementing this proposal will lead to better code quality, higher test coverage in these specific areas of the code, and better reusability possibilities as we grow the ecosystem (e.g. breaking out the _cloud provider_ code, building Cluster API controllers, etc.). This work consists of three major pillars, and we hope to complete at least the ComponentConfig part of it—if not (ideally) all three pieces of work—in v1.14.

### "Component" definition

In this case, when talking about a "component", I mean: "a CLI tool or a long-running server process that consumes configuration from a versioned configuration file (with apiVersion/kind) and optionally overriding flags". The component's implementation of ComponentConfig and command &amp; flag setup is well unit-tested. The component is to some extent Kubernetes-aware. The component follows Kubernetes' conventions for config serialization and merging, logging, and common HTTPS endpoints (in the server case). _To begin with_, this proposal will **only focus on the core Kubernetes components** (kube-apiserver, kube-controller-manager, kube-scheduler, kubelet, kube-proxy, kubeadm), but as we go, this library will probably be generic enough to be usable by cloud provider and Cluster API controller extensions, as well as aggregated API servers.

### Goals

- Make it easy for a component to correctly adopt ComponentConfig.
- Avoid moving code into _k8s.io/apiserver_ which does not strictly belong to an etcd-based, API group-serving apiserver. Corollary: remove etcd dependency from components.
- Components should be consistent in how they load (and write) configuration, and merge config with CLI flags.
- Factor out command- and flag-building code to a shared place.
- Factor out common HTTPS endpoints describing a component's status.
- Make the core Kubernetes components utilize these new packages.
- Have good documentation about how to build a component with a similar look and feel as core Kubernetes components.
- Increase test coverage for the configuration, command building, and HTTPS server areas of the component code.
- Break out OpenAPI definitions and violations for the current ComponentConfigs from the core repo to a dedicated place per-component. With auto-generated OpenAPI specs for each component ComponentConfig consumers can validate/vet their configs without running the component.

### Success metrics

- All core Kubernetes components (kube-apiserver, kube-controller-manager, kube-scheduler, kubelet, kube-proxy, kubeadm) are using these shared packages in a consistent manner.
- Cloud providers can be moved out of core without having to depend on the core repository.
  - Related issue: [https://github.com/kubernetes/kubernetes/issues/69585](https://github.com/kubernetes/kubernetes/issues/69585)
- It's easier for _kubeadm_ to move out of the core repo when these component-related packages are in a "public" staging repository.

### Non-goals

- Graduate any ComponentConfig API versions (in this proposal).
- Make this library toolkit a "generic" cloud-native component builder. Such a toolbox, if ever created, could instead consume these packages. In other words, this repository is solely focused on Kubernetes' components' needs.
- Fixing _all the problems_ our components have right now, and expanding this beyond what's really necessary. Instead working incrementally, and starting to break out some basic stuff we _know_ every component must handle (e.g. configuration and flag parsing)

### Related proposals / references

- [Kubernetes Component Configuration](https://docs.google.com/document/d/1arP4T9Qkp2SovlJZ_y790sBeiWXDO6SG10pZ_UUU-Lc/edit) by [@mikedanese](https://github.com/mikedanese)
- [Versioned Component Configuration Files](https://docs.google.com/document/d/1FdaEJUEh091qf5B98HM6_8MS764iXrxxigNIdwHYW9c/edit#) by [@mtaufen](https://github.com/mtaufen)
- [Moving ComponentConfig API types to staging repos](https://github.com/kubernetes/community/blob/master/keps/sig-cluster-lifecycle/0014-20180707-componentconfig-api-types-to-staging.md) by [@luxas](https://github.com/luxas) and [@sttts](https://github.com/sttts)

## Proposal

This proposal contains three logical units of work. Each subsection is explained in more detail below.

### Part 1: ComponentConfig

#### Standardized encoding/decoding

- Encoding/decoding helper methods in `k8s.io/component/config/serializer` that would be referenced in every scheme package
- Warn or (if desired, error) on unknown fields
  - This has the benefit that it makes it possible for the user to spot e.g. config typos. More high-level, this can be used for e.g. a `--validate-config` flag
- Support both JSON and YAML
- Support multiple YAML documents if needed

#### Testing helper methods

- Conversion / roundtrip testing
- API group testing
  - External types must have JSON tags
  - Internal types must not have any JSON tags
  - ...
- Defaulting testing

#### Generate OpenAPI specifications

Provide a common way to generate OpenAPI specifications local to the component, so that external consumers can access it, and the component can expose it via e.g. a CLI flag or HTTPS endpoint.

### Part 2: Command building / flag parsing

#### Wrapper around cobra.Command

See the `cmd/kubelet` code for how much extra setup a Kubernetes component needs to do for building commands and flag sets. This code can be refactored into a generic wrapper around _cobra_ for use with Kubernetes.

#### Flag precedence over config file

If the component supports both ComponentConfiguration and flags, flags should override fields set in the ComponentConfiguration. This is not straightforward to implement in code, and only the kubelet does this at the moment. Refactoring this code in a generic helper library in this new repository will make adoption of the feature easy and testable. The details of flag versus ComponentConfig semantics are to be decided later in a different proposal. Meanwhile, this flag precedence feature will be opt-in, so the kubelet and kubeadm can directly adopt this code, until the details have been decided on for all components.

#### Standardized logging

Use the _k8s.io/klog_ package in a standardized way.

### Part 3: HTTPS serving

Many Kubernetes controllers are clients to the API server and run as daemons. In order to expose information on how the component is doing (e.g. profiling, metrics, current configuration, etc.), an HTTPS server is run.

#### Common endpoints

In order to make it easy to expose this kind of information, a package is made in this new repo that hosts this common code. Initially targeted endpoints are "/version," "/healthz," "/configz," "/pprof," and "/metrics."

#### Standardized authentication / authorization

In order to not expose this kind of information (e.g. metrics) to anyone that can talk to the component, it may utilize SubjectAccessReview requests to the API server, and hence delegate authentication and authorization to the API server. It should be easy to add this functionality to your component.

### Part 4: Sample implementation in k8s.io/sample-component

Provides an example usage of the three main functions of the _k8s.io/component_ repo, implementing ComponentConfig, the CLI wrapper tooling and the common HTTPS endpoints with delegated auth.

### Code structure

- k8s.io/component
  - config/
    - Would hold internal, shared ComponentConfig types across core components
    - {v1,v1beta1,v1alpha1}
      - Would hold external, shared ComponentConfig types across core components
    - serializer/
      - Would hold common methods for encoding/decoding ComponentConfig
    - testing/
      - Would hold common testing code for use in unit tests local to the implementation of ComponentConfig.
  - cli/
    - Would hold common methods and types for building a k8s component command (building on top of github.com/spf13/{pflag,cobra})
    - options/
      - Would hold flag definitions
    - testing/
      - Would hold common testing code for use in unit tests local to the implementation of the code
    - logging/
      - Would hold common code for using _k8s.io/klog_
  - server/
    - auth/
      - Would hold code for implementing delegated authentication and authorization to Kubernetes
    - configz/
      - Would hold code for implementing a `/configz` endpoint in the component
    - healthz/
      - Would hold code for implementing a `/healthz` endpoint in the component
    - metrics/
      - Would hold code for implementing a `/metrics` endpoint in the component
    - pprof/
      - Would hold code for implementing a `/pprof` endpoint in the component
    - version/
      - Would hold code for implementing a `/version` endpoint in the component

### Timeframe and Implementation Order

**Objective:** The ComponentConfig part done for v1.14

**Stretch goal:** Get the CLI and HTTPS server parts done for v1.14.

**Implementation order:**

1. Create the k8s.io/component repo with the initial ComponentConfig shared code
2. Move shared v1alpha1 ComponentConfig types and references from `k8s.io/api{server,machinery}/pkg/apis/config` to `k8s.io/component/config`
3. Set up good unit testing for all core ComponentConfig usage, by writing the `k8s.io/component/config/testing` package
4. Move server-related util packages from `k8s.io/kubernetes/pkg/util/` to `k8s.io/component/server`. e.g. delegated authn/authz "/configz", "/healthz", and "/metrics" packages are suitable
5. Move common flag parsing / cobra.Command setup code to `k8s.io/component/cli` from (mainly) the kubelet codebase.
6. Start using the command- and server-related code in all core components.

In parallel to all the steps above, a _k8s.io/sample-component_ repo is built up with an example and documentation how to consume the _k8s.io/component_ code

### OWNERS file for new packages

- Approvers for the config/{v1,v1beta1,v1alpha1} packages
  - @kubernetes/api-approvers
- Approvers for staging/src/k8s.io/{sample-,}component
  - @sttts
  - @luxas
  - @jbeda
  - @lavalamp
- Approvers for subpackages:
  - those who owned packages before code move
- Reviewers for staging/src/k8s.io/{sample-,}component:
  - @sttts
  - @luxas
  - @dixudx
  - @rosti
  - @stewart-yu
  - @dims
- Reviewers for subpackages:
  - those who owned packages before code move
