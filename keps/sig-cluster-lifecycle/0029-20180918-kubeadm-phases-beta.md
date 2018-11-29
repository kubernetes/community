---
kep-number: 29
title: kubeadm phases to beta
authors:
  - "@fabriziopandini"
owning-sig: sig-cluster-lifecycle
reviewers:
 - "@chuckha"
 - "@detiber"
 - "@liztio"
 - "@neolit123"
approvers:
 - "@luxas"
 - "@timothysc"
editor:
 - "@fabriziopandini"
creation-date: 2018-03-16
last-updated: 2018-09-10
status: provisional
see-also:
 - KEP 0008
---

# kubeadm phases to beta

## Table of Contents

<!-- TOC -->

- [kubeadm phases to beta](#kubeadm-phases-to-beta)
    - [Table of Contents](#table-of-contents)
    - [Summary](#summary)
    - [Motivation](#motivation)
        - [Goals](#goals)
        - [Non-goals](#non-goals)
    - [Proposal](#proposal)
        - [User Stories](#user-stories)
        - [Implementation Details](#implementation-details)
    - [Graduation Criteria](#graduation-criteria)
    - [Implementation History](#implementation-history)
    - [Drawbacks](#drawbacks)
    - [Alternatives](#alternatives)

<!-- /TOC -->

## Summary

We are defining the road map for graduating `kubeadm alpha phase` commands to
beta, addressing concerns/lessons learned so far about the additional
effort for maintenance of this feature.

## Motivation

The `kubeadm phase` command was introduced in v1.8 cycle under the `kubeadm alpha phase`
command with the goal of providing users with an interface for invoking individually
any task/phase of the `kubeadm init` workflow.

During this period, `kubeadm phase` proved to be a valuable and composable
API/toolbox that can be used by any IT automation tool or by an advanced user for
creating custom clusters.

However, the existing separation of `kubeadm init` and `kubeadm phase` in the code base
required a continuous effort to keep the two things in sync, with a proliferation of flags,
duplication of code or in some case inconsistencies between the init and phase implementation.

### Goals

- To define the approach for graduating to beta the `kubeadm phase` user
  interface.
- To significantly reduces the effort for maintaining `phases` up to date
  with `kubeadm init`.
- To support extension of the "phases" concept to other kubeadm workflows.
- To enable re-use of phases across different workflows e.g. the cert phase
  used by the `kubeadm init` and by the `kubeadm join` workflows.

### Non-goals

- This proposal doesn't include any changes of improvements to the actual `kubeadm init`
  workflow.
- This proposal doesn't include implementation of workflows different than `kubeadm init`;
  nevertheless, this proposal introduces a framework that will allow such implementation in future.

## Proposal

### User Stories

- As a kubernetes administrator/IT automation tool, I want to run all the phases of
  the `kubeadm init` workflow.
- As a kubernetes administrator/IT automation tool, I want to run only one or some phases
  of the `kubeadm init` workflow.
- As a kubernetes administrator/IT automation tool, I want to run all the phases of
  the `kubeadm init` workflow except some phases.

### Implementation Details

The core of the new phase design consist into a simple, lightweight workflow manager to be used
for implementing composable kubeadm workflows.

Composable kubeadm workflows are build by an ordered sequence of phases; each phase can have it's
own, nested, ordered sequence of phases. For instance:

```bash
  preflight       Run master pre-flight checks
  certs           Generates all PKI assets necessary to establish the control plane
    /ca             Generates a self-signed kubernetes CA 
    /apiserver      Generates an API server serving certificate and key
    ...
  kubeconfig      Generates all kubeconfig files necessary to establish the control plane 
    /admin          Generates a kubeconfig file for the admin to use and for kubeadm itself
    /kubelet        Generates a kubeconfig file for the kubelet to use.
    ...
  ...
````

The above list of ordered phases should be made accessible from all the command supporting phases
via the command help, e.g. `kubeadm init --help` (and eventually in the future `kubeadm join --help` etc.)

Additionally we are going to improve consistency between the command outputs/logs with the name of phases
defined in the above list. This will be achieved by enforcing that the prefix of each output/log should match
the name of the corresponding phase, e.g. `[certs/ca] Generated ca certificate and key.` instead of the current
`[certificates] Generated ca certificate and key.`.

Single phases will be made accessible to the users via a new `phase` sub command that will be nested in the
command supporting phases, e.g. `kubeadm init phase` (and eventually in the future `kubeadm join phase` etc.). e.g.

```bash
kubeadm init phases certs [flags]

kubeadm init phases certs ca [flags]
```

Additionally we are going also to allow users to skip phases from the main workflow, via the `--skip-phases` flag. e.g.

```bash
kubeadm init --skip-phases addons/proxy
```

The above UX will be supported by a new components, the `PhaseRunner` that will be responsible
of running phases according to the given order; nested phases will be executed
immediately after their parent phase.

The `PhaseRunner` will be instantiated by kubeadm commands with the configuration of the specific list of ordered
phases; the `PhaseRunner` in turn will dynamically generate all the `phase` sub commands for the phases.

Phases invoked by the `PhaseRunner` should be designed in order to ensure reuse across different
workflows e.g. reuse of phase `certs` in both `kubeadm init` and `kubeadm join` workflows.

## Graduation Criteria

* To create a periodic E2E test that bootstraps a cluster using phases
* To document the new user interface for phases in kubernetes.io

## Implementation History

* [#61631](https://github.com/kubernetes/kubernetes/pull/61631) First prototype implementation 
  (now outdated)

## Drawbacks

By merging phases into kubeadm workflows derives a reduced capability to customize
the user interface for each phase. More specifically:

- It would not be possible to provide any kind of advice to the user about which
  flags are relevant for one specific phase (the help will always show all the flags).
- It would not be possible to add long description and/or examples to each phase
- It would not be possible to provide additional flags specific for one phase
  (the flags are shared between init and all the phases).
- It would not be possible to expose to the users phases which are not part of kubeadm workflows
  ("extra" phases should be hosted on dedicated commands).

This is considered an acceptable trade-off in light of the benefits of the suggested
approach.

## Alternatives

It is possible to graduate phases by simply moving corresponding command to first level,
but this approach provide less opportunities for reducing the effort
for maintaining phases up to date with the changes of kubeadm.