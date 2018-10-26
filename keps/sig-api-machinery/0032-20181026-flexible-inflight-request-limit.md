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
    * [How to store limits config](#how-to-store-limits-config)
    * [How to determinate request weight](#how-to-determinate-request-weight)
* [Discussionable questions](#discussionable-questions)
    * [How to treat namespace-scoped lists?](#how-to-treat-namespace-scoped-lists)
    * [Should we  make users connect to all apiservers to use their entire allotted qps?](#should-we-make-users-connect-to-all-apiservers)

## Summary <a name="summary"></a>
This KEP proposes to implement `readonly_capacity`, `mutating_capacity`, `readonly_limit`, `readonly_request` and `mutating_limit`, `mutating_request`, for readonly & mutating requests accordingly for different user groups, considering different physical resource consumption by particular operations (e.g. get single pod vs. get all pods). `*_limit` limits total inflight request `tokens` from certain group and `*_request` reserves some amount of `tokens` for critical groups.


## Motivation <a name="motivation"></a>

Buggy service or malicious user may easily run out of '--max-mutating-requests-inflight' or '--max-requests-inflight' request limit by DOSing the apiserver. It will block actually important users (kubelets, controllers/schedulers), that, potentially, leads to impossibility to stop buggy service or block attacker via k8s.

Using only `limit` to prevent blocking important users isn't enough flexible: to guarantee service availability for this users we should set rigid limits, when, sum of all limits not exceed server capacity (e.g. --max-requests-inflight). It will lead to throttling particular user group even if the server not loaded by other users.

Combination of `limit` and `request` allows to prevent overreservation of resources for critical users on the one hand, and flexible restrict different parts of the system from overspending resources on the other hand.

Using limits in `tokens` instead of limiting number of requests allows to
account resource consumption more realistic.


### Goals <a name="goals"></a>

* Implement per-group `limit` and `reserve`
* Implement request cost estimation
* Find appropriate resource to store limits & reserves map

## Proposal <a name="proposal"></a>

Implement weight-based per-group limits instead of current count-based limits.
Every request has its own `weight` measured in `tokens`, according to its complexity. Exact algorithm of request cost calculation described below.

Implement following parameters:
* `readonly_capacity`, `mutating_capacity`: Total `weight` of inflight requests,
that apiserver could carry.
* `readonly_limit`,  `mutating_limit`: Limits total requests `weight` across  particular group.
* `readonly_request`,  `mutating_request`: Reserved `weight` for particular group.

**Example**: We have `readonly_capacity` = 1000, `readonly_request` specified for `system:masters` is 50, for `system:node` is 100, so total available `tokens` for users from other groups is 850. In other words, members of `system:masters` in any time have at least 50 `tokens`.


### Algorithm to accept/reject readonly request <a name="algorithm-to-accept-reject-readonly-request"></a>

0. Pick group with max `readonly_request` across user's group as `request_group`, pick group with max `readonly_limit` across user's group as `limit_group`.
1. If $total `weight` of inflight readonly requests of `requestGroup` > `limit_group.readonly_limit`, REJECT inbound request
2. Else, if total `weight` of inflight readonly requests of `requestGroup` < `request_group.readonly_request`, ACCEPT inbound request
3. Else, if (sum of `readonly_request` of all groups + total `weight` of all readonly requests, exceedes appropriate `request_group.readonly_request`) < `readonly_capacity`, ACCEPT inbound request
4. else, REJECT  inbound request

For mutating requests algorithm is the same.

### How to store limits config <a name="how-to-store-limits-config"></a>

Implement new API resource to store group limits and requests:
```
apiVersion: extensions/v1beta1
kind: InflightLimit
metadata:
  name: %USER_GROUP_NAME%
spec:
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

`%USER_GROUP_NAME%` - The name of appropriate user group

`%READONLY_TOKENS_LIMIT%` - `readonly_limit`

`%READONLY_TOKENS_REQUEST%` - `readonly_request`

`%MUTATING_TOKENS_LIMIT%` - `mutating_limit`

`%MUTATING_TOKENS_REQUEST%` - `mutating_request`

`%CURRENT_READONLY_TOKENS_INFLIGHT%` - Total `weight` of inflight readonly requests from the user group

`%CURRENT_MUTATING_TOKENS_INFLIGHT%` - Total `weight` of inflight mutating requests from the user group

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
