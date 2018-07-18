---
kep-number: 13
title: pull based audit events stream
authors:
  - "@CaoShuFeng"
owning-sig: sig-auth
participating-sigs:
  - sig-api-machinery
reviewers:
  - TBD
approvers:
  - TBD
editor: TBD
creation-date: 2018-06-11
last-updated: 2018-07-18
status: provisional
see-also:
  - n/a
replaces:
  - n/a
superseded-by:
  - n/a

---

# Pull based audit events stream

## Table of Contents
* [Pull based audit events stream](#pull-based-audit-events-stream)
   * [Table of Contents](#table-of-contents)
   * [Summary](#summary)
   * [Motivation](#motivation)
      * [Goals](#goals)
      * [Non-Goals](#non-goals)
   * [User Stories](#user-stories)
      * [Story 1](#story-1)
      * [Story 2](#story-2)
   * [Proposal](#proposal)
      * [new resource endpoints](#new-resource-endpoints)
      * [Field Selector for subset of audit events](#field-selector-for-subset-of-audit-events)
      * [Bandwidth limit Configuration](#bandwidth-limit-configuration)
      * [Risks and Mitigations](#risks-and-mitigations)
   * [Graduation Criteria](#graduation-criteria)
   * [Implementation History](#implementation-history)
   * [Alternatives](#alternatives)

## Summary

This proposal aims at providing a pull-based stream for audit logging. This would look similar to a watch request to the API server, and provide stream of audit events. The stream does not return any previous events (i.e. no associated storage), it would only stream events that occurred after the stream was opened. Two new resource endpoints is installed to provide the stream: '/apis/audit.k8s.io/v1beta1/events' and '/apis/audit.k8s.io/v1alpha1/namespaces/{namespace}/namespacedevents'. These endpoints would only support the WATCH verb in the first.

## Motivation
This endpoint provides a good way for short-term debugging. Both tenants and cluster administrator could use this endpoint to get event stream without any changes to the cluster. Examples include:
- allow tenants to get audit events in their own namespaces
- provide a subset of audit events, so it will be easier to query and analyze (e.g. https://github.com/kubernetes/kubernetes/issues/56683)


### Goals
- Provide a new pull based audit event stream, which returns events that occurred after the stream was opened.
- Allow users go get a subset of audit events.
- Access control. Allow cluster administrator to get all events in cluster, allow namespace administrators to get all events in their own namespace.
- Bandwidth limit in server scope or per-namespace.

### Non-Goals
- Access control beyond namespace granularity(e.g. allow user A to only get audits events from users A,B,C or allow user A to only get audit events of resource foo)
- Provide a custom per-stream policies. The audit stream should be based on the cluster-wide audit policy.

## User Stories

### Story 1
As a tenant, I will be able to known all things happened to my own namespace. When problem happens, I can use the audit events for trouble shooting.

### Story 2
As a developer, I will easily be able to get a subset of audit events. This will save a lot of time for debugging. I don't need to grep audit events from a couple of audit files.

## Proposal

### new resource endpoints
Two new resource endpoint '/apis/audit.k8s.io/v1beta1/events' and '/apis/audit.k8s.io/v1alpha1/namespaces/{namespace}/namespacedevents' are installed.
Unlike other resource endpoints which usually serve resource objects from key-value storage, these endpoints read events from a audit backend and only support the WATCH verb. '/apis/audit.k8s.io/v1beta1/events' provides a stream of all audit events in the cluster. '/apis/audit.k8s.io/v1alpha1/namespaces/{namespace}/namespacedevents' provides a stream of all audit events in the namespace. A new API object NamespacedEvent is introduced for the namespaced endpoint. NamespacedEvent has the same definition with the audit Event object which we already have, but `metadata.namespace` element is required for it.

The audit events returned to users are the events generated according to the audit policy file. If an event is omitted by the audit policy file, the event would not be sent to user by these endpoints. For aggregated API servers, those endpoints only supports metadata-level, because kube-apiserver only record these events at metadata level.

To implement these watch endpoints, a new audit backend is introduced. This new backend will provide [watch.Interface](https://github.com/kubernetes/kubernetes/blob/release-1.11/staging/src/k8s.io/apimachinery/pkg/watch/watch.go#L28) to apimachinery. And apimachinery use this watch.Interface to implement the watch endpoint.

### Field Selector for subset of audit events
Some field selector would be supported for users to get a subset of audit events. (e.g. objectRef.namespace, user.username)

### Bandwidth limit Configuration

```golang
// LimitType is the type of the limit.
type LimitType string

const (
	// ServerLimitType is a type of limit where there is one bucket shared by
	// all of the audit events sent by the API Server, including cluster-scoped and
	// namespaced endpoints.
	ServerLimitType LimitType = "server"

	// NamespaceLimitType is a type of limit where there is one bucket used by
	// each namespace.
	NamespaceLimitType LimitType = "namespace"
)

// Configuration provides bandwidth configuration for the pull based audit event
// endpoints.
type Configuration struct {
	metav1.TypeMeta `json:",inline"`

	// limits are the limits to place on audit events sent by API server or
	// each endpoint.
	Limits []Limit `json:"limits"`
}


// Limit is the configuration for a particular limit type
type Limit struct {
	// type is the type of limit to which this configuration applies
	Type LimitType `json:"type"`

	// ThrottleQPS is the number of bytes sent per second that are allowed for this
	// type of limit.
	ThrottleQPS int32 `json:"throttleQps"`

	// ThrottleBurst defines the maximum number of bytes sent to the client at the
	// same moment.
	ThrottleBurst int32 `json:"throttleBurst"`

	// MaxEventSize defines max allowed size of the event. If the event is larger, truncating will be performed
	// before emitting to client.
	// optional
	MaxEventSize int64 `json:"maxEventSize"`
}
```

When the sent bytes exceed the bandwidth, API server will try to truncate the audit event, and check the limit again. If bandwidth limit is still exceeded, the event would be dropped.


### Risks and Mitigations
These endpoints should be treated as debugging & development endpoints, not for actual secure audit logging.
This change introduces two new endpoints which expose sensitive information. The cluster admin should should configure the authorization module and set access control to them.
The access control is supported at namespace level. Only namespace admin should be able to read audit events from endpoint `/apis/audit.k8s.io/v1alpha1/namespaces/{namespace}/namespacedevents`.

## Graduation Criteria
Success will be determined by stability of the provided pull based streams and ease of understanding for the end user.

* alpha: Pull based audit event stream and bandwidth limit works as expected,known issues are tested and resolved.
* beta: Mechanisms have been hardened against any known bugs and the process is validated by the community

[umbrella issues]: https://github.com/kubernetes/kubernetes/issues/53455

## Implementation History

- 06/11/2018: initial design
- 07/18/2018: update

## Alternatives

Cluster admin deploy another application to implement multi-tenant support. (e.g. https://kubernetes.io/docs/tasks/debug-application-cluster/audit/#log-collector-examples)
