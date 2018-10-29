---
kep-number: 32
title: Flexible inflight request limit
authors:
  - "@mshaverdo"
owning-sig: sig-api-machinery
reviewers:
  - @wojtek-t
  - @lavalamp
  - @yliaog
  - @jpbetz
approvers:
  - TBD
editor: TBD
creation-date: 2018-10-26
status: provisional
---

# Flexible inflight request limit

## Table of Contents

* [Summary](#summary)
* [Motivation](#motivation)
    * [Goals](#goals)
* [Proposal](#proposal)
    * [Algorithm to accept/reject readonly request](#algorithm-to-accept-reject-readonly-request)
    * [How to determinate request weight](#how-to-determinate-request-weight)
* [Discussionable questions](#discussionable-questions)
    * [How to treat namespace-scoped lists?](#how-to-treat-namespace-scoped-lists)
    * [Should we  make users connect to all apiservers to use their entire allotted qps?](#should-we-make-users-connect-to-all-apiservers)

## Summary <a name="summary"></a>
This KEP proposes to implement readonly & mutating `capacity` for the apiserver; readonly & mutating `limit` and `request` of requests for different users and groups, considering different physical resource consumption by particular operations (e.g. get single pod vs. get all pods). `limit` limits total inflight request `tokens` from certain user or group and `request` reserves some amount of `tokens` for critical users or groups.


## Motivation <a name="motivation"></a>

Buggy service or malicious user may easily run out of '--max-mutating-requests-inflight' or '--max-requests-inflight' request limit by DOSing the apiserver. It will block actually important users (kubelets, controllers/schedulers), that, potentially, leads to impossibility to stop buggy service or block attacker via k8s.

Using only `limit` to prevent blocking important users isn't enough flexible: to guarantee service availability for this users we should set rigid limits, when, sum of all limits not exceed server capacity (e.g. --max-requests-inflight). It will lead to throttling particular user or group even if the server not loaded by other users.

Combination of `limit` and `request` allows to prevent overreservation of resources for critical users on the one hand, and flexible restrict different parts of the system from overspending resources on the other hand.

Using limits in `tokens` instead of limiting number of requests allows to
account resource consumption more realistic.


### Goals <a name="goals"></a>

* Implement per-user and per-group `limit` and `request`
* Implement request cost estimation

## Proposal <a name="proposal"></a>

Implement weight-based per-user and per-group limits instead of current count-based limits.
Every request has its own `weight` measured in `tokens`, according to its complexity. Exact algorithm of request cost calculation described below.

Implement options `readonly_capacity`, `mutating_capacity`: Total `weight` of inflight requests that apiserver could carry.

Implement new API resource `InflightLimit` to store group limits and requests:
```
apiVersion: extensions/v1beta1
kind: InflightLimit
metadata:
  name: %NAME%
spec:
  subjects:
    kind: %SUBJECT_KIND%
    name: %SUBJECT_NAME%
  readonly:
    limit: %READONLY_TOKENS_LIMIT%
    request: %READONLY_TOKENS_REQUEST%
  mutating:
    limit: %MUTATING_TOKENS_LIMIT%
    request: %MUTATING_TOKENS_REQUEST%
status:
  readonlyInflightTokens: %CURRENT_READONLY_TOKENS_INFLIGHT%
  mutatingInflightTokens: %CURRENT_MUTATING_TOKENS_INFLIGHT%
```

Where:

`%NAME%` - The name of `InflightLimit` resource

`%SUBJECT_KIND%` - `User` or `Group`, like in `RoleBinding`

`%SUBJECT_NAME%` - User or group name

`%READONLY_TOKENS_LIMIT%` - Limits total readonly requests `weight` across  particular user or group

`%READONLY_TOKENS_REQUEST%` - Reserved (always available) `weight` of readonly requests  for particular user or group

`%MUTATING_TOKENS_LIMIT%` - Limits total mutating requests `weight` across  particular user or group.

`%MUTATING_TOKENS_REQUEST%` - Reserved (always available) `weight` of mutating requests  for particular user or group

`%CURRENT_READONLY_TOKENS_INFLIGHT%` - Total `weight` of inflight readonly requests from the user or group

`%CURRENT_MUTATING_TOKENS_INFLIGHT%` - Total `weight` of inflight mutating requests from the user or group

**Example**: We have `readonly_capacity` = 1000. `readonly.request` specified for `system:masters` is 50, for user `node` is 100, so total available readonly request `tokens` for other users is 850. In other words, members of `system:masters` in any time have at least 50 `tokens`, user `node` in any time have at least 100 `tokens`.


### Algorithm to accept/reject readonly request <a name="algorithm-to-accept-reject-readonly-request"></a>

0. User $CURRENT_USER made readonly request to apiserver, pick the request `weight` as `$REQUEST_WEIGHT`

1. If $CURRENT_USER has appropriate `InflightLimit` resource, pick it as `$CURRENT_INFLIGHTLIMIT_LIMIT` and `$CURRENT_INFLIGHTLIMIT_REQUEST`
Otherwise, among groups of the `$CURRENT_USER`, pick `InflightLimit` with max `readonly.limit` as `$CURRENT_INFLIGHTLIMIT_LIMIT` and `InflightLimit` with max `readonly.request` as `$CURRENT_INFLIGHTLIMIT_REQUEST`.

2. If `$CURRENT_INFLIGHTLIMIT_LIMIT.status.readonlyInflightTokens + $REQUEST_WEIGHT > $CURRENT_INFLIGHTLIMIT_LIMIT.spec.readonly.limit`,
**REJECT** inbound request

3. Else, if `$CURRENT_INFLIGHTLIMIT_REQUEST.status.readonlyInflightTokens + $REQUEST_WEIGHT <= $CURRENT_INFLIGHTLIMIT_REQUEST.spec.readonly.request`,
 **ACCEPT** inbound request

4. Else, if (sum (`readonlyInflightTokens - readonly.request`) of all `InflightLimit` objects ) + `$REQUEST_WEIGHT`
<=  `readonly_capacity` - (sum (`readonly.request`) of all `InflightLimit` objects),
 **ACCEPT** inbound request

5. Else, **REJECT**  inbound request

For mutating requests algorithm is the same.

### How to determinate request weight<a name="how-to-determinate-request-weight"></a>

Implement request `weight` calculation as a function F(verb_weight, related_resources_count), where verb_weight is a constant for particular verb, related_resources_count -- count of pods, events or services -- according to requested resource type and namespace.

Resources counts could be stored in the map and periodically updated in background like in etcd_object_counts metrics: https://github.com/kubernetes/kubernetes/blob/master/staging/src/k8s.io/apiserver/pkg/storage/etcd/metrics/metrics.go#L62


## Discussionable questions <a name="discussionable-questions"></a>

### How to treat namespace-scoped lists? <a name="how-to-treat-namespace-scoped-lists"></a>

There is an open question how to treat namespace-scoped lists (e.g. list all pods from namespace X) - those are clearly different than singular gets and listing all objects of a given type

(@wojtek-t)

**My opinion:** Due to object key schema in etcd (`/registry/%OBJECT_TYPE%/%NAMESPACE%/%OBJECT_NAME%`) there is no difference between
listing of all objects of given type and all objects of given type in particular namespace: in both cases request `weight` depends only from count of related objects.

### Should we  make users connect to all apiservers to use their entire allotted qps? <a name="should-we-make-users-connect-to-all-apiservers"></a>

This needs to also be fair overall when we are
running multiple apiservers. (e.g., we shouldn't make users connect to all
three apiservers to use their entire allotted qps--or should we? The KEP
should discuss this.)

(@lavalamp)

**My opinion:** We should use load balancer to balance user requests across multiple apiservers
