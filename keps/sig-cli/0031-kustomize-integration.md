---
kep-number: 31
title: Enable kustomize in kubectl
authors:
  - "@Liujingfang1"
owning-sig: sig-cli
participating-sigs:
  - sig-cli
reviewers:
  - "@pwittrock"
  - "@seans3"
  - "@soltysh"
approvers:
  - "@pwittrock"
  - "@seans3"
  - "@soltysh"
editor: TBD
creation-date: 2018-11-07
last-updated: yyyy-mm-dd
status: pending
see-also:
  - [KEP-0008](https://github.com/kubernetes/community/blob/master/keps/sig-cli/0008-kustomize.md)
replaces:
  - n/a
superseded-by:
  - n/a
---

# Enable kustomize in kubectl

## Table of Contents
* [Table of Contents](#table-of-contents)
* [Summary](#summary)
* [Motivation](#motivation)
    * [Goals](#goals)
    * [Non-Goals](#non-goals)
* [Kustomize Introduction](#kustomize-introduction)    
* [Proposal](#proposal)
    * [UX](#UX)
      * [apply](#apply)
      * [get](#get)
      * [delete](#delete)
    * [Implementation Details/Notes/Constraints](#implementation-detailsnotesconstraints)
    * [Risks and Mitigations](#risks-and-mitigations)
* [Graduation Criteria](#graduation-criteria)
* [Implementation History](#implementation-history)
* [Alternatives](#alternatives)

[Tools for generating]: https://github.com/ekalinin/github-markdown-toc

## Summary
[Kustomize](https://github.com/kubernetes-sigs/kustomize) is a tool developed to provide declarative support for kubernetes objects. It has been adopted by many projects and users. Having kustomize enabled in kubectl will address a list of long standing issues. This KEP describes how `kustomize` is integrated into kubectl subcommands with consistent UX.

## Motivation

Declarative specification of Kubernetes objects is the recommended way to manage Kubernetes applications or workloads. There is some gap in kubectl on declarative support. To eliminate the gap, a [KEP](https://github.com/kubernetes/community/blob/master/keps/sig-cli/0008-kustomize.md#faq) was proposed months ago and Kustomize was developed. After more than 10 iterations, Kustomize has a complete set of features and reached a good state to be integrated into kubectl. 

### Goals

Integrate kustomize with kubectl so that kubectl can recognize kustomization directories and expand resources from kustomization.yaml before running kubectl subcommands. This integration should be transparent. It doesn't change kubectl UX. This integration should also be backward compatible. For non kustomization directories, kubectl behaves the same as current. The integration shouldn't have any impact on those parts.


### Non-Goals
- provide an editing functionality of kustomization.yaml from kubectl
- further integration with other kubectl flags

## Kustomize Introduction

Kustomize has following subcommands:
- build
- edit, edit also has subcommands
   - Set 
       - imagetag
       - namespace
       - nameprefix
   - Add
       - base
       - resource
       - patch
       - label
       - annotation
       - configmap
- config
- version
- help

`edit` and `build` are most commonly used subcommands. 

`edit` is to modify the fields in `kustomization.yaml`. A `kustomization.yaml` includes configurations that are consumed by Kustomize. Here is an example of `kustomization.yaml` file.

`build` is to perform a set of pre-processing transformations on the resources inside one kustomization. Those transformations include:
- Get objects from the base
- Apply patches
- Add name prefix to all resources
- Add common label and annotation to all resources
- Replace imageTag is specified
- Update objects’ names where they are referenced
- Resolve variables and substitute them

```
# TODO: currently kustomization.yaml is not versioned
# Need to version this with apiVersion and Kind
# https://github.com/kubernetes-sigs/kustomize/issues/588
namePrefix: alices-

commonAnnotations:
 oncallPager: 800-555-1212

configMapGenerator:
- name: myJavaServerEnvVars
 literals:   
 - JAVA_HOME=/opt/java/jdk
 - JAVA_TOOL_OPTIONS=-agentlib:hprof

secretGenerator:
- name: app-sec
 commands:
   username: "echo admin"
   password: "echo secret"
```
The build output of this sample kustomizaiton.yaml file is
```
apiVersion: v1
data:
 JAVA_HOME: /opt/java/jdk
 JAVA_TOOL_OPTIONS: -agentlib:hprof
kind: ConfigMap
metadata:
 annotations:
   oncallPager: 800-555-1212
 name: alices-myJavaServerEnvVars-7bc9c27cmf
---
apiVersion: v1
data:
 password: c2VjcmV0Cg==
 username: YWRtaW4K
kind: Secret
metadata:
 annotations:
   oncallPager: 800-555-1212
 name: alices-app-sec-c7c5tbh526
type: Opaque
```

## Proposal

### UX

When apply, get or delete is run on a directory, check if it contains a kustomization.yaml file. If there is, apply, get or delete the output of kustomize build. Kubectl behaves the same as current for directories without kustomization.yaml.

#### apply
The command visible to users is
```
kubectl apply -f <dir>
```
To view the objects in a kustomization without applying them to the cluster
```
kubectl apply -f <dir> --dry-run -o yaml|json
```

#### get
The command visible to users is
```
kubectl get -f <dir>
```
To get the detailed objects in a kustomization
```
kubectl get -f <dir> --dry-run -o yaml|json
```

#### delete
The command visible to users is
```
kubectl delete -f <dir>
```

### Implementation Details/Notes/Constraints

To enable kustomize in kubectl, the function `FilenameParam` inside Builder type will be updated to recognize kustomization directories. The Builder will expand the sources in a kustomization directory and pass them to a subcommand.

 Since kustomization directories themselves have a recursive structure, `-R` will be ignored on those directories. Allowing recursive visit to the same files will lead to duplicate resources. 

The examples and descriptions for apply, get and delete will be updated to include the support of kustomization directories. 

### Risks and Mitigations

This KEP doesn't provide a editing command for kustomization.yaml file. Users will either manually edit this file or use `Kusotmize edit`. 

## Graduation Criteria

There are two signals that can indicate the success of this integration.
- Kustomize users drop the piped commands `kustomize build <dir> | kubectl apply -f - ` and start to use apply directly.
- Kubectl users put their configuration files in kustomization directories.


## Implementation History

Most implementation will be in cli-runtime

- vendor `kustomize/pkg` into kubernetes
- copy `kustomize/k8sdeps` into cli-runtime
- Implement a Visitor for kustomization directory which
   - execute kustomize build to get a list of resources
   - write the output to a StreamVisitor
- When parsing filename parameters in FilenameParam, look for kustomization directories
- update the examples in kubectl commands
- Improve help messages or documentations to list kubectl subcommands that can work with kustomization directories

## Alternatives

The approaches in this section are considered, but rejected.
### Copy kustomize into staging
Copy kustomize code into kubernetes/staging and have the staging kustomize as source of truth. The public kustomize repo will be synced automatically with the staging kustomize.
- Pros
  - Issues can be managed in the public repo
  - The public repo can provide a kustomize binary
  - The public repo can be used as kustomize libraries   
  - Empower further integration of kubectl with kustomize
- Cons
  - The staging repo is designed for libraries that will be separated out. Moving kustomize into staging sounds controversial
  - Kustomize will be in staging, the development will be done in k/k repository
  - Development velocity will be reduced of every release

### Add kustomize as a subcommand in kubectl
Add kustomize as a subcommand into kubectl was the first way we tried to enable kustomize in kubectl. The PR was [add kustomize as a subcommand of kubectl](https://github.com/kubernetes/kubernetes/pull/70213).
- Pros
  - kustomize command is visible to users
  - the code change is straightforward
  - easy to test 
- Cons
  - UX is not consistent with other kubectl subcommands
  - Apply command will include two parts
            `kubectl kustomize build dir | kubectl apply -f -`
