# Sig Ownership Responsibilities

## Overview

This document describes the commitment that is undertaken by a [special interest group] as part of owning
a feature or api change.  This applies to changes made to the [kubernetes/kubernetes] repo.

## Design

Depending on the size and impact of a feature, the design may need to go through a review process.  Changes
that do not directly impact the api should follow the process defined by the owning sig.  Changes that
do impact the api must go through an [api review] process.  Even non-api changes must think carefully about
backwards compatibility to ensure clusters will continue to function correctly after an upgrade or downgrade.

## Implementation

Implementation includes writing and reviewing the code for the new feature.

## Release tracking

The owning special interest group must create an issue in [kubernetes/features] and respond to comments
on the feature.  The feature issue MUST be updated with the current status within 1-2 days after 
the [feature freeze] deadline.

## Testing

Testing must be complete before the implementation is merged.

### Unit + integration testing

All new features must have appropriate coverage with unit + integration tests.

### E2e testing

E2e tests for a feature are owned by the sig that owns the feature, and must consistently pass.
Tests that periodically fail must be fixed within the flake SLA defined for the project.

### Upgrade testing

Changes that could be impacted by upgrade / downgrades of a cluster must have automated tests to verify
functionality during and after upgrade / downgrades.

### Client-server version skew testing

Any changes that break compatibility with client-server version skew tests of +-1 minor release must be resolved.

## Documentation

All user facing changes must be documented at [kubernetes/kubernetes.github.io] to
appear in the [user docs].  See the [doc contribution guide] for instructions.

[feature freeze]: https://github.com/kubernetes/community/blob/master/contributors/devel/release/README.md#feature-freeze
[kubernetes/kubernetes.github.io]: https://github.com/kubernetes/kubernetes.github.io/
[doc contribution guide]: https://kubernetes.io/editdocs/
[special interest group]: https://github.com/kubernetes/community/blob/master/README.md#special-interest-groups-sig
[kubernetes/kubernetes]: https://github.com/kubernetes/kubernetes
[kubernetes/features]: https://github.com/kubernetes/features/
[api review]: https://github.com/kubernetes/community/blob/master/contributors/contributing-api-changes.md
[user docs]: https://k8s.io