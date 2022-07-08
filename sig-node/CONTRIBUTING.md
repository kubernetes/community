# Contributing on SIG Node

Welcome!

## For Kubernetes Contributions

https://github.com/kubernetes/community/tree/master/contributors/guide#contributor-guide

If you aspire to grow scope in the SIG, please review the [SIG Node contributor ladder](./sig-node-contributor-ladder.md) for SIG specific guidance.

### For Enhancements 

https://github.com/kubernetes/enhancements/tree/master/keps/sig-node

#### Helpful Links for Sig-Node

Code:  

https://github.com/kubernetes/kubernetes/tree/master/cmd/kubelet  

https://github.com/kubernetes/kubernetes/tree/master/pkg/kubelet  

https://github.com/kubernetes/kubernetes/tree/master/pkg/probe  

Development Resources:  

https://github.com/kubernetes/community/tree/master/contributors/devel#table-of-contents

Shared space / Sub projects:  

https://github.com/kubernetes/community/tree/master/contributors/devel/sig-node  

https://github.com/kubernetes/community/tree/master/sig-node#subprojects

Triage:
https://github.com/kubernetes/community/blob/master/contributors/devel/sig-node/triage.md 

## Getting Started

Task #1 : Compile kubelet
See tips in the root Makefile    
 
https://github.com/kubernetes/community/blob/master/contributors/devel/development.md#building-kubernetes

Task #2 : Run a single unit test  

https://github.com/kubernetes/community/blob/master/contributors/devel/development.md#unit-tests

Task #3 : Explore update/verify scripts  

hack/update-gofmt.sh + hack/verify-gofmt.sh  

https://github.com/kubernetes/kubernetes/blob/master/hack/update-gofmt.sh  

https://github.com/kubernetes/kubernetes/blob/master/hack/verify-gofmt.sh

Task #4 : Explore dependencies  

hack/pin-dependency.sh + hack/update-vendor.sh  

https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/vendor.md

Task #5 : Using local-up-cluster script  

https://github.com/kubernetes/community/blob/master/contributors/devel/running-locally.md#starting-the-cluster  

Running a local cluster  

https://github.com/kubernetes/community/blob/master/contributors/devel/running-locally.md
        
Note: Task 5 requires Linux OS
