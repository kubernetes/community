Selector generation from kubectl
==================

# Goals

Make `kubectl apply` work well with selector.

# Problem Description

The field `spec.selector` of Deployments, ReplicaSet, DaemonSet and ReplicationController are defaulted from `spec.template.metadata.labels`, if it is unspecified.
And there is validation to assure `spec.selector` always selects the pod template.
The defaulting of selector may prevent from `kubectl apply` from working when updating
the `spec.selector` field. e.g. [kubernetes/kubernetes#26202 (comment)](https://github.com/kubernetes/kubernetes/issues/26202#issuecomment-221421254)

# Proposed changes

We can let `kubectl` do something similar to [Selector Generation](https://github.com/kubernetes/community/blob/master/contributors/design-proposals/selector-generation.md) on creation. We can let `kubectl` add a set of labels to assure non-overlapping. 
There are no changes on server-side and the server-side defaulting is still kept for backward compatible reason.

## API

No required change.

## kubectl

Add client-side selector generation, which is disabled by default and can be enabled by flag in `kubectl`.

`kubectl` will add a set of labels to `spec.selector` and `spec.template.metadata.labels` before sending the object to the API server. 

The generated labels will not be included in the `kubectl.kubernetes.io/last-applied-configuration` annotation, if it is created by `kubectl create --save-config` or `kubectl apply`.

### Generated labels

There are at least 3 sets of labels to choose:

The first set of labels is:

- Add 3 more labels to `spec.selector`
  - `matchLabels["resource-group"]="$GROUP"`
  - `matchLabels["resource-kind"]="$KIND"`
  - `matchLabels["resource-name"]="$NAME"`
- Append 3 labels to `spec.template.metadata.labels`
  - `resource-group=$GROUP`
  - `resource-kind=$KIND`
  - `resource-name=$NAME`

The second set of labels is:

- Add 1 more label, which concatenates all the three in the first set, to `spec.selector`
  - `matchLabels["controllers.k8s.io/selector"]="$GROUP/$KIND/$NAME"`
- Append 1 label to `spec.template.metadata.labels`
  - `controllers.k8s.io/selector=$GROUP/$KIND/$NAME`

The third set of labels is:

- Add 1 more label to `spec.selector`
  - `matchLabels["resource-uid"]="$UIDOFRESOURCE"`
- Append 1 label to `spec.template.metadata.labels`
  - `resource-uid=$UIDOFRESOURCE`

### Rationale

Label combination of `resource group`, `resource kind` and `resource name` is unique across all kind of resources. They are predictable, human-friendly and self-documenting, but we need all of them to assure uniqueness.

Label `UID` itself is unique. It is simple but unpredictable.

User can put on the pod template and/or selector some labels that are useful to the user, without reasoning about non-overlappingness.  `kubectl` adds additional labels to assure non-overlapping.

### Orphaning and adoption

When orphaning, the user can simply use non-cascading deletion: `kubectl delete --cascade=false`. 
The labels in the orphaned pods will remain unchanged.

When adopting, the user will need to specify `spec.selector` and `spec.template.metadata.labels` with all the label keys in the set we choose. The these labels should match the labels that the controller want to adopt.
For example, if we choose to use set 1, the user should specify `resource-group`, `resource-kind` and `resource-name` in both `spec.selector` and `spec.template.metadata.labels`.

If the user fails to specify all the label keys in the set, `kubectl` will fail the operation. 

In this case, the labels in the set will be inlcuded in the `kubectl.kubernetes.io/last-applied-configuration` annotation.

## Docs

Update examples that affected.

# Alternatives considered

# 1. Selector Generation from Admission Controller

# Proposed changes

We can move all the logic described above to a plug-in for admission controller.
There are no changes on server-side and the server-side defaulting is still kept for backward compatible reason.

## API

No required change.

## kubectl

No required change.

## Admission controller

When receiving an incoming POST request of a workload controller (excluding job controller), the admission controller will add a set of labels to `spec.selector` and `spec.template.metadata.labels`.

This feature is disabled by default and can be enabled by query parameters.

The generated labels will not be included in the `kubectl.kubernetes.io/last-applied-configuration` annotation, if it is created by `kubectl create --save-config` or `kubectl apply`.

### Generated labels

The choice of labels is the same as in section [generated labels](#generated-labels).

### Orphaning and adoption

Orphaning and adoption are working the same as in section [orphaning and adoption](#orphaning-and-adoption).

# 2. Selector Generation from server-side

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
- Defaulting logic sets `spec.selector` to
  - `matchLabels["resource-group"]="$GROUPOFRESOURCE"`
  - `matchLabels["resource-kind"]="$KINDOFRESOURCE"`
  - `matchLabels["resource-name"]="$NAMEOFRESOURCE"`
- Defaulting logic appends 3 labels to the `spec.template.metadata.labels`.
  - `resource-kind=$KINDOFRESOURCE`.
  - `resource-kind=$KINDOFRESOURCE`.
  - `resource-name=$NAMEOFRESOURCE`.

### Manual Mode
Manual mode is the same as [Job's](https://github.com/kubernetes/community/blob/master/contributors/design-proposals/selector-generation.md#manual-mode).

### Rationale

Label combination of `resource name` and `resource kind` is unique across all kind of resources.
They are more predictable and human-friendly than UID.

If user does specify `spec.selector` then the user must also specify `spec.manualSelector`. This ensures the user knows that what he is doing is not the normal thing to do.


## kubectl

No required changes.

## Docs

Update or remove examples that affected.

## Version skew

1) old client vs new APIserver: old client don't know field `spec.manualSelector`, it will be considered as nil value at server-side and the server will auto-generate selector.

2) new client vs old APIserver: `spec.manualSelector` will be ignored, since the APIServer don't understand this field. It should behave as before.
