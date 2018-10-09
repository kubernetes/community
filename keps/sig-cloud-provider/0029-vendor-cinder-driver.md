---
kep-number: 29
title: Vendor Cinder Storage Driver
authors:
  - "@hogepodge"
  - "@dims"
owning-sig: sig-cloud-provider
participating-sigs:
  - sig-openstack
  - sig-storage
reviewers:
  - "@hogepodge"
  - "@dims"
  - "@cheftaco"
  - "@andrewsykim"
approvers:
  - "@jagosan"
  - "@andrewsykim"
editor: TBD
creation-date: 2018-10-09
last-updated: 2018-10-09
status: provisional
see-also:
replaces:
superseded-by:
---

# Vendor Cinder Storage Driver

## Table of Contents

A table of contents is helpful for quickly jumping to sections of a KEP and for highlighting any additional information provided beyond the standard KEP template.
[Tools for generating][] a table of contents from markdown are available.

* [Table of Contents](#table-of-contents)
* [Summary](#summary)
* [Motivation](#motivation)
    * [Goals](#goals)
    * [Non-Goals](#non-goals)
* [Proposal](#proposal)
    * [User Stories [optional]](#user-stories-optional)
      * [Story 1](#story-1)
      * [Story 2](#story-2)
    * [Implementation Details/Notes/Constraints [optional]](#implementation-detailsnotesconstraints-optional)
    * [Risks and Mitigations](#risks-and-mitigations)
* [Graduation Criteria](#graduation-criteria)
* [Implementation History](#implementation-history)
* [Drawbacks [optional]](#drawbacks-optional)
* [Alternatives [optional]](#alternatives-optional)

[Tools for generating]: https://github.com/ekalinin/github-markdown-toc

## Summary

As part of the multi-stage process to remove the OpenStack provider from upstream
Kubernetes, existing dependencies on the in-tree Cinder driver need to be vendored
to give parity between the external and in-tree OpenStack provider code.

This is an intermediary step for full removal of the code, and can serve as an
example for how other providers can approach similar problems in provider code that
is deeply entangled in upstream code.

## Motivation

As part of the SIG-Cloud-Provider charter to foster neutral collaboration within
the Kubernetes community, there is an active goal of removing all in-tree cloud
providers from within the Kuberntes code-base and replacing them with external,
expert maintained provider code. Over the course of Kubernetes development, the
in-tree providers have built up a complex set of dependencies that makes full
replacement challenging. In the case of OpenStack, one of these dependencies is
with in-tree Cinder integrations.

### Goals

As part of the process to extract and remove the in-tree provider, we need to
create a path for continuing to support in-tree Cinder dependencies. The
primary goal of this KEP is to create an interface to the in-tree Cinder provider
that vendors out the the external cinder provider. This will assist in the larger
goal of deleting in-tree OpenStack code and allowing projects that depend on the
provider to continue to work.

A secondary goal is to create a development entrypoint to continue to remove
in-tree provider code, with a future focus on facilitating the extraction of
integrated Cinder code.

A tertiary goal is to demonstrate a means for staging removal of in-tree code
by creating interfaces that external provider can vendor against, with the goal
of eventual removal of all provider code.

This KEP will be considered completed if the dependencies on the in-tree Cinder
interface continues to work while calline out to the external provider.


### Non-Goals

* Removal of in-tree Cinder code is out of scope for this KEP.
* Use of the CSI interface is out of scope for thie KEP.

## Proposal

An in-tree interface will be created for the Kubernetes Cinder code. This
interface will call out to the external Cinder provider that is vendored
in to Kubernetes, allowing for the eventual removal of the in-tree Cinder
code without breaking existing dependencies.

### User Stories [optional]

OpenStack Magnum, a Kubernetes installation tool, depends on the in-tree OpenStack
provider to deliver automated Kubernetes instllations to users through an
easy-to-use API. In transitioning future releases of the Magnum project to use
the external provider rather than the in-tree provider, it needs to provide the
same level of functionality. This implementation gives them a migration path that
allows upstream removal of providers to continue without significantly impacting
downstream users.

### Risks and Mitigations

We have identified the following risk:

* This may create a new interface that must be supported in-tree. This can be
mitigated by a commitment from the OpenStack developers and SIG-Cloud-Provider
to deprecate and remove the code and internal interfaces in a timely manner.
* This proposal creates a circular dependency, with in-tree code depending on
the external provider code, but with the external provider code depending on
in-tree interfaces. Any updates must be carefully planned and synchronized.

## Graduation Criteria

This KEP will be successful if the in-tree Cinder code can be removed without
disrupting users and libraries that depend on it.

[umbrella issues]: https://github.com/kubernetes/kubernetes/issues/42752

## Implementation History

- [https://github.com/kubernetes/cloud-provider-openstack/pull/317]: https://github.com/kubernetes/cloud-provider-openstack/pull/317
- [Experimental Cinder Volume Provider]: https://github.com/kubernetes/kubernetes/pull/69529
