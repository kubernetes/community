---
kep-number: 31
title: VolumeSubpathEnvExpansion
authors:
  - "@kevtaylor"
owning-sig: sig-storage
participating-sigs:
  - sig-storage
  - sig-architecture
reviewers:
  - "@msau42"
  - "@thockin"
approvers:
  - "@thockin"
  - "@msau42"
editor: TBD
creation-date: 2018-10-29
last-updated: 2018-10-29
status: implementable
see-also:
  - n/a
replaces:
  - n/a
superseded-by:
  - n/a
---

# Title

VolumeSubpathEnvExpansion API change

## Table of Contents

  * [Title](#title)
      * [Table of Contents](#table-of-contents)
      * [Summary](#summary)
      * [Motivation](#motivation)
         * [Goals](#goals)
         * [Non-Goals](#non-goals)
      * [Proposal](#proposal)
         * [User Stories](#user-stories)
      * [Current workarounds - k8s &lt;=1.9.3](#current-workarounds---k8s-193)
      * [Workarounds - k8s &gt;1.9.3](#workarounds---k8s-193)
      * [Alternatives - using subPath directly](#alternatives---using-subpath-directly)
         * [Risks and Mitigations](#risks-and-mitigations)
      * [Graduation Criteria](#graduation-criteria)
      * [Implementation History](#implementation-history)
      * [Alternatives - Using subPathFrom](#alternatives---using-subpathfrom)

## Summary

Legacy systems create all manner of log files and these are not easily streamed into stdout

Files that are written to a host file path need to be uniquely partitioned

If 2 or more pods run on the same host writing the same log file names to the same volume, they will clash

Using the `subPath` is a neat option but because the `subPath` is "hardcoded" ie. `subPath: mySubPath` it does not enforce uniqueness

To alleviate this, the `subPath` should be able to be configured from an environment variable as `subPath: $(MY_VARIABLE)` 

The workaround to this issue is to use symbolic links or relative symbolic links but these introduce sidecar init containers
and a messy configuration overhead to try to create upfront folders with unique names - examples of this complexity are detailed below

## Motivation

The initial alpha feature was implemented to allow unique addressing of subpaths on a host
This cannot currently be achieved with the downwardAPI and requires complex workarounds

The workarounds became more difficult after 1.9.3 when symbolic links were removed from initContainers

### Goals

To reduce excessive boiler-plate workarounds and remove the need for complex initContainers

### Non-Goals

Full template implementation for subPaths

## Proposal

The api change proposed is to create a Mutually Exclusive Field separate from the `subPath`
called `subPathExpr`

The subpath code which expands environment variables from the API would (under this proposal) change from

```
    env:
    - name: POD_NAME
      valueFrom:
        fieldRef:
          apiVersion: v1
          fieldPath: metadata.name
 
   ...
   
    volumeMounts:
    - name: workdir1
      mountPath: /logs
      subPath: $(POD_NAME)
```

to:

```
    volumeMounts:
    - name: workdir1
      mountPath: /logs
      subPathExpr: $(POD_NAME)
```

This would then introduce the new element to be processed separately from the `subPath`

### User Stories

## Current workarounds - k8s <=1.9.3

This makes use of symbolical linking to the underlying subpath system
The symbolic link element was removed after 1.9.3

```
apiVersion: extensions/v1beta1
kind: Deployment
metadata:
  labels:
    app: podtest
  name: podtest
spec:
  replicas: 1
  selector:
    matchLabels:
      app: podtest
  template:
    metadata:
      labels:
        app: podtest
    spec:
      containers:
      - env:
        - name: POD_NAME
          valueFrom:
            fieldRef:
              apiVersion: v1
              fieldPath: metadata.name
        - name: POD_NAMESPACE
          valueFrom:
            fieldRef:
              apiVersion: v1
              fieldPath: metadata.namespace
        image: <image>
        name: podtest
        volumeMounts:
        - mountPath: /logs
          name: workdir
          subPath: logs
      initContainers:
      - command:
        - /bin/sh
        - -xc
        - |
          LOGDIR=/logs/${POD_NAMESPACE}/${POD_NAME}; mkdir -p ${LOGDIR} && ln -sfv ${LOGDIR} /workdir/logs && chmod -R ugo+wr ${LOGDIR}
        env:
        - name: POD_NAME
          valueFrom:
            fieldRef:
              apiVersion: v1
              fieldPath: metadata.name
        - name: POD_NAMESPACE
          valueFrom:
            fieldRef:
              apiVersion: v1
              fieldPath: metadata.namespace
        image: alpine:3.5
        name: prep-logs
        volumeMounts:
        - mountPath: /logs
          name: logs
        - mountPath: /workdir
          name: workdir
      volumes:
      - emptyDir: {}
        name: workdir
      - hostPath:
          path: /logs
          type: ""
        name: logs
```

## Workarounds - k8s >1.9.3

Beyond 1.9.3 some attempts were made to provide a workaround using relative paths, rather than symbolic links

These have been deemed to be a cumbersome manipulation of the operating system and are flawed and unworkable

This effectively negates an upgrade path to k8s 1.10

The only foreseeable solution is to move directly to 1.11 from 1.9.3 and switch on the alpha feature gate VolumeSubPathEnvExpansion

## Alternatives - using subPath directly

An initial attempt has been made for this but there are edge case compatibility issues which are highlighted here 
https://github.com/kubernetes/kubernetes/pull/65769 regarding the alpha implementation 

The objection is regarding backward compatibility with existing users' subpaths

Because of a breaking change in the API, it has been decided to offer an alternative based on the
original discussion here https://github.com/kubernetes/kubernetes/issues/48677

The `VolumeSubPathEnvExpansion` alpha feature was delivered in k8s 1.11 allowing
subpaths to be created from downward api variables as 
```
apiVersion: v1
kind: Pod
metadata:
  name: pod1
spec:
  containers:
  - name: container1
    env:
    - name: POD_NAME
      valueFrom:
        fieldRef:
          apiVersion: v1
          fieldPath: metadata.name
    image: busybox
    command: [ "sh", "-c", "while [ true ]; do echo 'Hello'; sleep 10; done | tee -a /logs/hello.txt" ]
    volumeMounts:
    - name: workdir1
      mountPath: /logs
      subPath: $(POD_NAME)
  restartPolicy: Never
  volumes:
  - name: workdir1
    hostPath: 
      path: /var/log/pods
```

Because of the mentioned breaking changes, this implementation cannot proceed forward

### Risks and Mitigations

The alpha implementation already provided a number of test cases to ensure that validations of `subPath` configurations 
were not circumvented or violated

The API change would ensure that the substitute of the variables takes place immediately before the subpath mount validation

We would also need review existing validation to ensure that any potential security issues are addressed as:
`$ escape and "../../../../../proc are not allowed`

Due to the vulnerabilities highlighted in https://github.com/kubernetes/kubernetes/issues/60813 the subpath validations in kubelet have been
highly orchestrated. Any implementation of this feature needs to ensure that the security fixes put in place are still effective


## Graduation Criteria

The existing alpha feature introduced many tests to mitigate issues. These would be reused as part of the api implementation.

[umbrella issues]: https://github.com/kubernetes/kubernetes/pull/49388

## Implementation History

* Initial issue: https://github.com/kubernetes/kubernetes/issues/48677
* Feature gate proposal: https://github.com/kubernetes/enhancements/issues/559
* Alpha Implementation: https://github.com/kubernetes/kubernetes/pull/49388
* Beta Issue: https://github.com/kubernetes/kubernetes/issues/64604
* Beta PR and Discussion: https://github.com/kubernetes/kubernetes/pull/65769

## Alternatives - Using subPathFrom
A possible further implementation could derive directly from the `fieldRef` as

```
        volumeMounts:
        - mountPath: /logs
          name: logs
          subPathFrom:
            fieldRef:
              fieldPath: metadata.name
      volumes:
      - name: logs
        hostPath:
          path: /logs
```

This method would not be favoured as it fixes the `subPath` to a single value and would not allow concatenation
of paths such as `$(NAMESPACE)/$(POD_NAME)`



