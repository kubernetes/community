# Contributing on SIG Node

Welcome!

## For Kubernetes Contributions

Read the [Kubernetes Contributor Guide](https://github.com/kubernetes/community/tree/master/contributors/guide#contributor-guide).

If you aspire to grow scope in the SIG, please review the [SIG Node contributor ladder](./sig-node-contributor-ladder.md) for SIG specific guidance.

### For Enhancements

SIG Node enhancements are available in the <https://github.com/kubernetes/enhancements/tree/master/keps/sig-node>.

#### Helpful Links for Sig-Node

**Code**:  

For general code organization, read [contributors/devel/README.md](../contributors/devel/README.md) for explaining things like
`vendor/`, `staging`, etc.

* Kubelet
  * <https://github.com/kubernetes/kubernetes/tree/master/cmd/kubelet>
  * <https://github.com/kubernetes/kubernetes/tree/master/pkg/kubelet>
* Probe: <https://github.com/kubernetes/kubernetes/tree/master/pkg/probe>
* NodeLifecycle: <https://github.com/kubernetes/kubernetes/tree/master/pkg/controller/nodelifecycle>
* Node API: <https://github.com/kubernetes/kubernetes/tree/master/staging/src/k8s.io/api/node>
* CRI API: <https://github.com/kubernetes/kubernetes/tree/master/staging/src/k8s.io/cri-api>
* DRA:
  * <https://github.com/kubernetes/kubernetes/tree/master/staging/src/k8s.io/dynamic-resource-allocation>
  * <https://github.com/kubernetes/kubernetes/tree/master/staging/src/k8s.io/kubelet/pkg/apis/dra/>
  * <https://github.com/kubernetes/kubernetes/tree/master/test/e2e/dra/>
* E2E test:
  * <https://github.com/kubernetes/kubernetes/tree/master/test/e2e/node/>
  * <https://github.com/kubernetes/kubernetes/tree/master/test/e2e_node/>
* CI (test-infra)
  * [sig-node jobs](https://github.com/kubernetes/test-infra/blob/master/config/jobs/kubernetes/sig-node/)
  * [e2e_node jobs](<https://github.com/kubernetes/test-infra/blob/master/jobs/e2e_node/>
  * [sig-node test-grids](https://github.com/kubernetes/test-infra/blob/master/config/testgrids/kubernetes/sig-node/)
  * [containerd test-grids](https://github.com/kubernetes/test-infra/blob/master/config/testgrids/kubernetes/containerd/)

**Development Resources**:  

* <https://github.com/kubernetes/community/tree/master/contributors/devel#table-of-contents>

There are two types of end-to-end tests in Kubernetes:

* [Cluster end-to-end tests](https://git.k8s.io/community/contributors/devel/sig-testing/e2e-tests.md)
* [Node end-to-end tests](https://git.k8s.io/community/contributors/devel/sig-node/e2e-node-tests.md)

**Shared space / Sub projects**:  

* <https://github.com/kubernetes/community/tree/master/contributors/devel/sig-node>
* <https://github.com/kubernetes/community/tree/master/sig-node#subprojects>

**Triage**:

* <https://github.com/kubernetes/community/blob/master/contributors/devel/sig-node/triage.md>

**Test Grids**:

* SIG Node overview: <https://testgrid.k8s.io/sig-node>
* Release Blocking: <https://testgrid.k8s.io/sig-node-release-blocking>
* Kubelet: <https://testgrid.k8s.io/sig-node-kubelet>
* Containerd: <https://testgrid.k8s.io/sig-node-containerd>

## Getting Started

Task #1 : Compile kubelet
See tips in the root Makefile:

* <https://github.com/kubernetes/community/blob/master/contributors/devel/development.md#building-kubernetes>

Task #2 : Run a single unit test  

* <https://github.com/kubernetes/community/blob/master/contributors/devel/development.md#unit-tests>

Task #3 : Explore update/verify scripts  

hack/update-gofmt.sh + hack/verify-gofmt.sh  

* <https://github.com/kubernetes/kubernetes/blob/master/hack/update-gofmt.sh>
* <https://github.com/kubernetes/kubernetes/blob/master/hack/verify-gofmt.sh>

Task #4 : Explore dependencies  

hack/pin-dependency.sh + hack/update-vendor.sh  

* <https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/vendor.md>

Task #5 : Using local-up-cluster script  

* <https://github.com/kubernetes/community/blob/master/contributors/devel/running-locally.md#starting-the-cluster>

Running a local cluster  

* <https://github.com/kubernetes/community/blob/master/contributors/devel/running-locally.md>

Note: Task 5 requires Linux OS

## How to ...

### upload meeting recording to YouTube (for chairs)

1. Download the video file from zoom account `sig-node-leads@kubernetes.io`
2. Log in to YouTube under the `kubernetes-sig-node` account (use switch account)
3. Upload video to YouTube
   - Name should follow the template: `<name of the meeting> yyyymmdd` like `Kubernetes SIG Node 20240618`.
   - Add a link to the meeting notes and agenda document to the meeting description
   - Add video to k8s-owned YouTube playlist:
     - Once video is uploaded, open the video page, click +Save button under the video, select the checkbox `SIG Node Meetings`.
4. Add link to the video to the meeting notes document
5. In zoom account change the title of the recording - add `(uploaded)` suffix
