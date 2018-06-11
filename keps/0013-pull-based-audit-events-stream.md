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
last-updated: 2018-06-11
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
* [Summary](#summary)
* [Motivation](#motivation)
    * [Goals](#goals)
    * [Non-Goals](#non-goals)
* [Proposal](#proposal)
    * [Make audit backends dynamicly registerable](#make-audit-backends-dynamicly-registerable)
    * [endpoints and query parameters](#endpoints-and-query-parameters)
    * [Bandwidth limit Configuration](#bandwidth-limit-configuration)
    * [User Stories](#user-stories)
      * [Story 1](#story-1)
      * [Story 2](#story-2)
    * [Risks and Mitigations](#risks-and-mitigations)
* [Graduation Criteria](#graduation-criteria)
* [Implementation History](#implementation-history)
* [Drawbacks [optional]](#drawbacks-optional)
* [Alternatives [optional]](#alternatives-optional)

[Tools for generating]: https://github.com/ekalinin/github-markdown-toc

## Summary

This proposal aims at providing a pull-based stream for audit logging. This would look similar to a watch request to the API server, and provide stream of audit events. The stream does not return any previous events (i.e. no associated storage), it would only stream events that occurred after the stream was opened. Two non resource-url endpoints are installed to provide the stream. '/audits/' for all audit events generated in the API server, and '/audits/{namespace}' for audits events in the specified namespace. And query parameters could be used to get subset of audit events.

## Motivation
The current audit backends (log & webhook) require very high privileges (file access to the master) or prior configuration (webhook). There are a number of use cases for dynamic and unprivileged access to the audit logs. Examples include:
- allow tenants to get audit events in their own namespace
- provide a subset of audit events, so it will be easier to query and analyze (e.g. https://github.com/kubernetes/kubernetes/issues/56683)


### Goals
- Provide a new pull based audit event stream, which returns events that occurred after the stream was opened.
- Allow users go get a subset of audit events.
- Access control. Allow unprivileged users to get audits in their own namespace.
- Bandwidth limit in server scope or per-namespace.

### Non-Goals
- Access control at user/group level(e.g. allow user A to only get audits events from users A,B,C)


## Proposal

### Make audit backends dynamicly registerable

```golang
// Registry is a decorator for the audit backend, it allows caller to dynamicly register/unregister new backends to it.
type Registry interface {
       Backend

       // Register register a new Backend to the registry.
       Register(Backend)

       // UnRegister remove the Backend from registry.
       UnRegister(Backend)
}
```
When a request comes, a new backend is registered to the registry. Audit Events sent to this backend will be sent to users in the end.

### endpoints and query parameters

|endpoint            |definition                                  |supported paramsters                                 |
|--------------------|--------------------------------------------|-----------------------------------------------------|
|/audits             |all audit events generated in the API server|username, group, namespace, apiGroup, resource, verbs|
|/audits/{namespaces}|audits events in the specified namespace    |username, group, apiGroup, resource, verbs           |

definition of query parameters
```golang
type queryParm struct {
       // username in the audit event.
       // optional
       username string
       // groups in the audit event, if all specified groups exist in the audit event, then it will be sent to user.
       // optional
       groups []string
       // namespace of the audit event, <none> for cluster scoped request, empty string for all namespaces.
       // only supported for /audits endpoint.
       // optional
       namespace string
       // apiGroup in the audit event, <core> for the core api group.
       // optional
       apiGroup string
       // resource in the audit event, for example: pods.
       // optional
       resource string
       // verb in the audit event, for example: create, put. If the verbs contains the verb of audit event, then
       // the audit event will be sent to user.
       // optional
       verbs []string
}
```

### Bandwidth limit Configuration
```golang
// LimitType is the type of the limit.
type LimitType string

const (
	// ServerLimitType is a type of limit where there is one bucket shared by
	// all of the audit events sent by the API Server, including '/audits' and
	// 'audits/{namespace}' endpoints.
	ServerLimitType LimitType = "server"

	// NamespaceLimitType is a type of limit where there is one bucket used by
	// each namespace. When the namespace is set to empty string, it limit the
	// bandwidth of all requests to '/audits' endpoint.
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

When the sent bytes exceeds the bandwidth, API server will try to truncate the audit event, and check the limit again. If bandwidth limit is still exceeded, the event would be droped.


### User Stories

#### Story 1
As a tenant, I will be able to known all things happened to my own namespace. When problem happens, I can use the audit events for trouble shooting.

#### Story 2
As a developer, I will easily be able to get a subset of audit events. This will save a lot of time for debugging. I don't need to grep audit events from a couple of audit files.

### Risks and Mitigations
We have bandwidth limit to prevent potential performance impact to API server. That means audit events could be dropped silently.
This change introduces two new endpoints which expose sensitive information. The cluster admin should should configure the authorization module and set access control to them.
The access control is supported at namespace level. Only namespace admin should be able to read audit events from endpoint `/audits/{namespace}/`.

## Graduation Criteria
Success will be determined by stability of the provided pull based streams and ease of understanding for the end user.

* alpha: Pull based audit event stream and bandwidth limit works as expected,known issues are tested and resolved.
* beta: Mechanisms have been hardened against any known bugs and the process is validated by the community

[umbrella issues]: https://github.com/kubernetes/kubernetes/issues/53455

## Implementation History

- 06/11/2018: initial design

## Alternatives [optional]

Cluster admin deploy another application to implement multi-tenant support. (e.g. https://kubernetes.io/docs/tasks/debug-application-cluster/audit/#log-collector-examples)
