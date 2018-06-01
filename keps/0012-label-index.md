---
kep-number: 12
title: Index labels to support efficient queries by label selectors
authors:
  - "@jian-he, @ccding"
owning-sig: sig-api-machinery
reviewers:
  - TBD
approvers:
  - TBD
creation-date: "2018-05-31"
last-updated: "2018-05-31"
status: provisional
---

## Summary

This proposal presents solutions to make the queries by label selectors in the
list API efficient.

## Motivation

Currently, only namespace index is supported in the cache. Controllers such as
ReplicaSet, when selecting pods, need to select all the pods by namespace and
then select the ones wanted filtering by index. While running a large number
of pods, we observed that listing all pods and then filtering takes several
seconds. This proposal presents solutions to improve this.


### Goals

Make the queries by label selectors in the list API efficient.

### Proposal

1. Adds label based index in controllers to facilitate lookup by label.

Since objects are created/deleted dynamically, we also need to update the index
efficiently.

2. Calculates the count for each index and evaluate the most selective first.

3. Provides the ability in API for users to mark which labels to index or build
all the index for all labels as the first step.

4. `Optional` Adds index in API-server cache to improve queries directly going
to API-server.

### Graduation Criteria

Provide performance testing results to prove the performance boost.
