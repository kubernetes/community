---
kep-number: 17
title: apiserver validates etcd instance
authors:
  - "@ccding"
owning-sig: sig-api-machinery
reviewers:
  - TBD
approvers:
  - TBD
editor: TBD
creation-date: 2018-07-24
last-updated: 2018-07-24
status: provisional
---

# apiserver validates etcd instance

## Summary

We would like to have apiserver to validate the etcd instance to make sure it
is connected to the proper one.

## Motivation

When there are multiple deployments of kubernetes and etcd, it may cause a
common mistake that the apiserver connects to a wrong etcd instance, which
would mess up everything.  By allowing apiserver to validate etcd instances,
it could avoid this from happening.

### Goals

When apiserver is configured to connect a etcd that has a different cluster
id, apiserver reports error and quits.

## Proposal

apiserver is configured with a cluster id. When it connects to an etcd
instance, it checks the cluster id stored in etcd. There are three possible
outcomes.

- if the etcd doesn't have a cluster id, it means the etcd is a fresh etcd.
  apiserver will write the cluster id to etcd, which will be used for
  validation in future. We have to make sure there is no race condition here.
- if the etcd stores a cluster id, and it matches the one configured to
  apiserver, apiserver uses the etcd normally.
- if the etcd stores a cluster id, but it doesn't match the one configured to
  apiserver, apiserver reports error and quits.

To make sure backward compatibility, when apiserver is not configured with a
cluster id, it uses a default cluster id and validates etcd using the default
cluster id. This works because with the same cluster id, it is equivalent to
that without cluster ids.

### Risks and Mitigations

Not aware

## Graduation Criteria

The apiserver won't connect to a wrong etcd instance.
