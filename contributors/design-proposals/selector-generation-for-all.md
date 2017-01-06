Selector Generation for Deployments, ReplicaSet, DaemonSet, StatefulSet and ReplicationController
=============

# Goals
Make selector easy to use and less error-prone for Deployments, ReplicaSet, DaemonSet, StatefulSet and ReplicationController. Make `kubectl apply` work well with selector.

# Problem Description
The field `spec.selector` of Deployments, ReplicaSet, DaemonSet are defaulted from `spec.template.metadata.labels`, if it is unspecified.
And there is validation to make sure `spec.selector` always selects the pod template.
The defaulting of selector may prevent from `kubectl apply` from working when updating the `spec.selector` field. e.g. [kubernetes/kubernetes#26202 (comment)](https://github.com/kubernetes/kubernetes/issues/26202#issuecomment-221421254)

# Proposed changes

The [Selector Generation](https://github.com/kubernetes/community/blob/master/contributors/design-proposals/selector-generation.md) works well in Job to make sure non-overlapping. We can use this pattern for Deployments, ReplicaSet, DaemonSet, StatefulSet and ReplicationController.
If `spec.selector` is not derived from `spec.template.metadata.labels`, `kubectl apply` will work well for selectors.

## API

The change will be similar to [the changes to Job](https://github.com/kubernetes/kubernetes/blob/master/docs/design/selector-generation.md#api). But the defaulting is different from Job's.

`extension/v1beta1 Deployment|ReplicaSet|DaemonSet`, `apps/v1beta1 StatefulSet` and `v1 ReplicationController` change as follows.

Add field `spec.manualSelector`. It controls if using automatic generated selectors.
In automatic mode, user cannot make the mistake of creating non-unique selectors. In manual mode, certain rare use cases are supported.

Validation is not changed: a selector must be provided, and it must select the pod template.

### Automatic Mode

- User does not specify `spec.selector`.
- User does not specify `spec.manualSelector` field or set it to `false`.
- User optionally puts labels on pod template (optional). User does not think
about uniqueness, just labeling for user's own reasons.
- Defaulting logic sets `job.spec.selector` to
  - `matchLabels["resource-kind"]="$KINDOFRESOURCE"`
  - `matchLabels["resource-name"]="$NAMEOFRESOURCE"`
- Defaulting logic  appends 2 labels to the `.spec.template.metadata.labels`.
  - The first label is `resource-kind=$KINDOFRESOURCE`.
  - The second label is `resource-name=$NAMEOFRESOURCE`.

### Manual Mode
Manual mode is the same as [Job's](https://github.com/kubernetes/community/blob/master/contributors/design-proposals/selector-generation.md#manual-mode).

### Rationale

Label combination of `resource name` and `resource kind` is unique across all kind of resources.
They are more predictable and human-friendly than UID.

If user does specify `spec.selector` then the user must also specify `spec.manualSelector`. This ensures the user knows that what he is doing is not the normal thing to do.

## Controllers

Controller managers are responsible to set generated labels when adopting.

## Garbage collector

Garbage collector is responsible for cleaning up the extra generated labels when orphaning.

## kubectl

No required changes.

## Docs

Update or remove examples that affected.

## Version skew

1) old client vs new APIserver: old client don't know field `spec.manualSelector`, it will be considered as nil value at server-side and the server will auto-generate selector.

2) new client vs old APIserver: `spec.manualSelector` will be ignored, since the APIServer don't understand this field. It should behave as before.
