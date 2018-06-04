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

### Graduation Criteria

Provide performance testing results to prove the performance boost.

# Appendix: performance evaluation

#### Evaluation objective
To evaluate the response time of [List Pods
api](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.10/#list-62)
when pods amount is huge

#### Evaluation setting.

Pods are created in the same namespace.  We evaluate variable number of pods
in the namespace: 500, 1000, 2000, 5000, 10000, 15000.

The body size of test pod is about 1.7 KB, which has content similar to the
one listed following.
```json
{"kind":"Pod","apiVersion":"v1","metadata":{"name":"pod-base","labels":{"site":"et15sqa","testMode":"true"}},"spec":{"containers":[{"name":"busybox","image":"reg.docker.alibaba-inc.com/busybox:latest","command":["sleep","3600"],"resources":{},"allocSpec":{"cpu":{"cpuCount":1,"cpuQuota":100,"cpuSetMode":"default","strategy":"default"},"memory":{"hardLimit":512000000},"disk":{"/":{"Size":8440000000,"mountPoint":"/"}},"gpu":{},"volume":null,"netIo":{}}}],"requirement":{"UpdateTime":"","RequirementId":"","site":"","targetReplica":1,"increaseReplica":1,"minReplica":0,"app":{"bizName":"smoking","appName":"smoking","deployUnit":"smoking","instanceGroup":"smoking","routeLabels":{"stage":"DAILY","unit":"CENTER_UNIT.center"},"overQuota":{"enable":false},"workDir":{"useHostWorkDir":false}},"spread":{"strictly":false},"constraints":{"namedLabels":{},"ignoreLabelBySpecifiedIp":false},"dependency":{},"affinity":{},"prohibit":{"appConstraints":null,"duConstraints":null}}},"status":{"netInfo":{}}}

```

We measure the response time of the following two operations: (1) list all the
pods, and (2) list one pod filtered by labels.

#### Evaluation result

| #pod | list all pods (sec) | RT of list one pod (sec) |
| ---------- | ------------------- | -------------------|
| 500 |0.850| 0.314|
| 1000 |1.611 |0.623 |
| 2000 |2.768|1.200 |
| 5000 |7.760|2.863 |
| 10000 |14.756| 5.590|
| 15000 |20.727|8.328 |

#### Evaluation conclusion

* With increasing number of pods, the response time of the list operations
  increases linearly.
* With more than 10K pods, it takes 14 seconds to list all the pods, and 5
  seconds to filter a single pod by label, which is unaffordable in our
  usecase.
