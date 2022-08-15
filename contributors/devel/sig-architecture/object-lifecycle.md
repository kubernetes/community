The API Object Lifecycle
=========================

*This document is oriented at users who want a deeper understanding of the
Kubernetes API structure, and developers wanting to extend the Kubernetes API.
An introduction to using resources with kubectl can be found in [the object management overview](https://kubernetes.io/docs/concepts/overview/working-with-objects/object-management/).*

**Table of Contents**

<!-- toc -->
- [States](#states)
  - [State Machine Diagram](#state-machine-diagram)
  - [Handling DELETING](#handling-deleting)
<!-- /toc -->

The [Kubernetes API](https://kubernetes.io/docs/api/) (and related APIs in the
ecosystem) defines "objects" (also called resources in REST context) which are
created, managed, and deleted over time.  Integrating with these APIs usually
involves writing "controllers" which actuate the API object into some external
form - for example, creating a `Pod` in the API is actuated by kubelet running
containers on a physical or virtual machine.

All Kubernetes API objects follow a common lifecycle which can be thought of as
a state-machine, though some specific APIs extend this and offer even more
states.  To write a correct controller, it's important to understand the common
object lifecycle.

## States

All Kubernetes API objects exist in one of the following states:

* `DOES_NOT_EXIST`: The object is not known to the API server.  Calling this a
  "state" is a bit of a stretch, but it helps to be explicit.  This state does
  not differentiate between "has not yet been created" and "has been deleted".
* `ACTIVE`: The API server is aware of the object and the object has not been
  deleted (the `metadata.deletionTimestamp` is not set).  While in this state,
  any update operations (PUT, PATCH, server-side apply, etc.) will result in
  this same state.
* `DELETING`: The API server is aware of the object, and the object has been
  deleted, _but has not been fully removed yet_.  This can either be because
  the object has one or more finalizers (in `metadata.finalizers`) or because
  it has a deletion grace period (in `metadata.deletionGracePeriodSeconds`)
  greater than zero (NOTE: most API types do not allow setting the
  `deletionGracePeriodSeconds` at all).  Clients can still access the object
  and can see that it is DELETING because the `metadata.deletionTimestamp`
  field is set.  When the last finalizer and/or the grace period are removed,
  the object will be removed from storage and truly cease to exist.

### State Machine Diagram

The following diagram describes the above states:

```
                                  +---- object
                                  |     updated
                                  v        |
                           +----------+    |
                           |          +----+
          object --------->|  ACTIVE  |
          created          |          +-----------+
             |             +---+------+           |
             |                 |                  |
             |                 |                  |
+------------+---+     object deleted             |
|                |     without finalizers         |
|                |<--- or grace period            |
|                |                           object deleted
| DOES_NOT_EXIST |                           with finalizers
|                |                           or grace period
|                |<--- finalizers removed         |
|                |     and grace period           |
+----------------+     complete                   |
                               |                  |
                               |                  |
                           +---+------+           |
                           |          |           |
                           | DELETING |<----------+
                           |          |
                           +----------+
```

### Handling DELETING

The `DELETING` state warrants some special attention from controllers. Because
this state did not really exist (for any API except Pods) until finalizers were
added, many older controllers do not consider it at all.  Even newly written
controllers do not always give it due consideration.

When crafting a controller, authors must consider what semantics to apply to
objects in the DELETING state.  Even if this state is not intended to be
meaningful to a given API, it can be expressed and can not generally be
prevented without significant effort.  It should therefore be considered and
handled, or at least documented.

To a controller, the transition from `ACTIVE` to `DELETING` appears as a
normal update operation which sets the `metadata.deletionTimestamp` field.
