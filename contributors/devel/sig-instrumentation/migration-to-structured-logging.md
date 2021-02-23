# Structured Logging migration instructions

This document describes instructions for migration proposed by [Structured Logging KEP]. It describes new structured
functions introduced in `klog` (Kubernetes logging library) and how log calls should be changed to utilize new features.
This document was written for the initial migration of `kubernetes/kubernetes` repository proposed for Alpha stage, but
should be applicable at later stages or for other projects using `klog` logging library.

[Structured Logging KEP]: https://github.com/kubernetes/enhancements/tree/master/keps/sig-instrumentation/1602-structured-logging

## Goal of Alpha migration

The first step is to introduce structure to the high percentage of logs generated in Kubernetes by changing only a
small number of logs API calls. Based on criteria described in the [selecting most important logs] section, the selected
22 log calls are estimated to impact 99.9% of log volume. The up to date list of these log calls is provided in the
[Enhancement Issue].

[Enhancement Issue]: https://github.com/kubernetes/enhancements/issues/1602
[selecting most important logs]: https://github.com/kubernetes/enhancements/tree/master/keps/sig-instrumentation/1602-structured-logging#selecting-most-important-logs

## Structured logging in Kubernetes

With this enhancement a set of new functions were added to `klog`. Structured logging functions the follow interface
based on [logr], which has a  different design than other `klog` functions which are based on [glog]. It is recommended
to familiarize yourself with [logr].

[logr]: https://github.com/go-logr/logr
[glog]: https://github.com/golang/glog

Here are the prototypes of functions added to `klog` that will be utilized during migration:
```go
package klog

// InfoS structured logs to the INFO log.
// The msg argument used to add constant description to the log line.
// The key/value pairs would be join by "=" ; a newline is always appended.
//
// Examples:
// >> klog.InfoS("Pod status updated", "pod", klog.KObj(pod), "status", "ready")
// output:
// >> I1025 00:15:15.525108       1 controller_utils.go:116] "Pod status updated" pod="kube-system/kubedns" status="ready"
func InfoS(msg string, keysAndValues ...interface{})

// ErrorS structured logs to the ERROR, WARNING, and INFO logs.
// the err argument used as "err" field of log line.
// The msg argument used to add constant description to the log line.
// The key/value pairs would be join by "=" ; a newline is always appended.
//
// Examples:
// >> klog.ErrorS(err, "Failed to update pod status")
// output:
// >> E1025 00:15:15.525108       1 controller_utils.go:114] "Failed to update pod status" err="timeout"
func ErrorS(err error, msg string, keysAndValues ...interface{})

// KObj is used to create ObjectRef when logging information about Kubernetes objects
// Examples:
// >> klog.InfoS("Pod status updated", "pod", klog.KObj(pod), "status", "ready")
// output:
// >> I1025 00:15:15.525108       1 controller_utils.go:116] "Pod status updated" pod="kube-system/kubedns" status="ready"
func KObj(obj KMetadata) ObjectRef

// KRef is used to create ObjectRef when logging information about Kubernetes objects without access to metav1.Object
// Examples:
// >> klog.InfoS("Pod status updated", "pod", klog.KRef(podNamespace, podName), "status", "ready")
// output:
// >> I1025 00:15:15.525108       1 controller_utils.go:116] "Pod status updated" pod="kube-system/kubedns" status="ready"
func KRef(namespace, name string) ObjectRef

// ObjectRef represents a reference to a kubernetes object used for logging purpose
// In text logs it is serialized into "{namespace}/{name}" or "{name}" if namespace is empty
type ObjectRef struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace,omitempty"`
}

// KMetadata is a subset of the kubernetes k8s.io/apimachinery/pkg/apis/meta/v1.Object interface
// this interface may expand in the future, but will always be a subset of the
// kubernetes k8s.io/apimachinery/pkg/apis/meta/v1.Object interface
type KMetadata interface {
	GetName() string
	GetNamespace() string
}
```

## Migration

1. Change log functions to structured equivalent
1. Remove string formatting from log message
1. Name arguments
1. Use `klog.KObj` and `klog.KRef` for Kubernetes object references
1. Verify log output

## Change log functions to structured equivalent

Structured logging functions follow a different logging interface design than other functions in `klog`. They follow
minimal design from [logr] thus there is no one-to-one mapping.

Simplified mapping between functions:
* `klog.Infof`, `klog.Info`, `klog.Infoln`, `klog.InfoDepth` -> `klog.InfoS`
* `klog.V(N).Infof`, `klog.V(N).Info`, `klog.V(N).Infoln` -> `klog.V(N).InfoS`
* `klog.Warning`, `klog.Warningf`, `klog.Warningln`, `klog.WarningDepth` -> `klog.InfoS`
* `klog.Error`, `klog.Errorf`, `klog.Errorln`, `klog.ErrorDepth` -> `klog.ErrorS`
* `klog.Fatal`, `klog.Fatalf`, `klog.Fatalln`, `klog.FatalDepth` -> `klog.ErrorS`

### Removing Depth

Functions with depth (`klog.InfoDepth`, `klog.WarningDepth`, `klog.ErrorDepth`, `klog.FatalDepth`) are used to indicate
that the source of the log (added as metadata in log) is different than the invocation of logging library. This is
usually used when implementing logging util functions. As logr interface doesn't support depth, those functions should
return logging arguments instead of calling `klog` directly.

For example
```go
func Handle(w http.ReponseWriter, r *http.Request) {
    logHTTPRequest(r)
    handle(w, r)
}

func logHTTPRequest(r *http.Request) {
    klog.InfoDepth(1, "Received HTTP %s request", r.Method)
}
```
should be replaced with
```go
func Handle(w http.ReponseWriter, r *http.Request) {
    klog.InfoS("Received HTTP request", httpRequestLog(r)...)
    handle(w, r)
}

func httpRequestLog(r *http.Request) []interface{} {
    return []interface{}{
        "verb", r.Method,
    }
}

```

### Using ErrorS

With `klog` structured logging borrowing the interface from [logr] it also inherits it's differences in semantic of
error function. Logs generated by `ErrorS` command may be enhanced with additional debug information
(such as stack traces) or be additionally sent to special error recording tools. Errors should be used to indicate
unexpected behaviours in code, like unexpected errors returned by subroutine function calls.

Calling `ErrorS` with `nil` as error is semi-acceptable if there is error condition that deserves a stack trace at this
origin point. For expected errors (`errors` that can happen during routine operations) please consider using
`klog.InfoS` and pass error in `err` key instead.

### Replacing Fatal calls

Use of Fatal should be discouraged and it's not available in new functions. Instead of depending on the logger to exit
the process, you should call `os.Exit()` yourself.

Fatal calls use a default exit code of 255. When migrating, please use an exit code of 1 and include an "ACTION REQUIRED:" release note.

## Remove string formatting from log message

With structured logging, log messages are no longer formatted, leaving argument marshalling up to the logging client
implementation. This allows messages to be a static description of event.

All string formatting (`%d`, `%v`, `%w`, `%s`) should be removed and log message string simplified.
Describing arguments in log messages is no longer needed and should be removed leaving only a description of what
happened.

Additionally we can improve messages to comply with good practices:
* Start from a capital letter.
* Do not end the message with a period.
* Use active voice. Use complete sentences when there is an acting subject ("A could not do B") or omit the subject if
  the subject would be the program itself ("Could not do B").
* Use past tense ("Could not delete B" instead of "Cannot delete B")
* When referring to an object, state what type of object it is. ("Deleted pod" instead of "Deleted")

For example
```go
klog.Infof("delete pod %s with propagation policy %s", ...)
```
should be changed to
```go
klog.InfoS("Deleted pod", ...)
```

Some logs are constructed solely from string formats. In those cases a message needs to be derived from the context of
the log call.

For example http access logs
```go
func LogHTTP(r *http.Request) {
   klog.Infof("%s %s: (%v) %v%v%v [%s %s]", ...)
}
```
should be changed to
```go
func LogHTTP(r *http.Request) {
   klog.InfoS("Received HTTP request", ...)
}
```

### Name arguments

Even though new structured logging functions have very similar function prototype `func (string, ...interface{})` it
has different meaning for variadic arguments. Instead of just passing arguments, now we are passing key value pairs of
argument name and argument value. This means when migrating a log call we need to add an additional string before each
argument, that will be used as it's name.

How variable arguments should be used:

```go
klog.InfoS("message", "key1", value1, "key2", "value2")
```

For example
```go
func LogHTTP(r *http.Request) {
   klog.Infof("Received HTTP request, path: %s, method: %s", r.Path, r.Method)
}
```
should be changed to
```go
func LogHTTP(r *http.Request) {
   klog.InfoS("Received HTTP request", "path", r.Path, "method", r.Method)
}
```

When deciding on names of arguments you should:
* Always use [lowerCamelCase], for example use `containerName` and not `container name` or `container_name`.
* Use [alphanumeric] characters: no special characters like `%$*`, non-latin, or unicode characters.
* Use object kind when referencing Kubernetes objects, for example `deployment`, `pod` and `node`.
* Describe the type of value stored under the key and use normalized labels:
  * Don't include implementation-specific details in the labels. Don't use `directory`, do use `path`.
  * Do not provide additional context for how value is used. Don't use `podIP`, do use `IP`.
  * With the exception of acronyms like "IP" and the standard "err", don't shorten names. Don't use `addr`, do use `address`.
  * When names are very ambiguous, try to include context in the label. For example, instead of
    `key` use `cacheKey` or instead of `version` use `dockerVersion`.
* Be consistent, for example when logging file path we should always use `path` and not switch between
  `hostPath`, `path`, `file`.

Here are a few exceptions to the rules above---some cases are temporary workarounds that may change if we settle on better solution:
* Do use `err` rather than `error` to match the key used by `klog.ErrorS`
* Context in name is acceptable to distinguish between values that normally go under same key. For example using both
  `status` and `oldStatus` in log that needs to show the change between statuses.
* When Kubernetes object kind is unknown without runtime checking we should use `object` key. To provide information
  about kind we should add separate `apiVersion` and `kind` fields.
* If we cannot use `klog.KObj` nor `klog.KRef` for Kubernetes object, like in cases when we only have access to name or UID,
  then we should fallback to using object kind with suffix based on value type. For example `podName`, `podUID`.
* When providing multiple indistinguishable values (for example list of evicted pods), then we can use plural version of
  argument name. For example we should use `pods` and not `podList`.

Examples of **good keys** (strongly suggested, will be extended when pattern emerge, no standard schema yet):
* `cacheKey`
* `cacheValue`
* `CIDR`
* `containerID`
* `containerName`
* `controller`
* `cronJob`
* `deployment`
* `dockerVersion`
* `duration`
* `err`
* `job`
* `object`
* `pod`
* `podName`
* `podUID`
* `PVC`
* `PV`
* `volumeName`
* `replicaSet`

Examples of **bad** keys:
* `addr` - replace with `address`
* `container` - replace with `containerName` or `containerID` depending on value
* `currentNode` - replace with `node`
* `directory` - replace with `path`
* `elapsed` - replace with `duration`
* `externalIP` - replace  with `IP`
* `file` - replace with `path`
* `hostPath` - replace with `path`
* `ip` - replace with `IP`
* `key` -  replace with key describing what kind of key it is, for example `cacheKey`
* `loadBalancerIP` - replace with `IP`
* `podFullName` - try to rewrite code so that pod name or pod object can be used with `pod` or `podName` keys
* `podIP` - replace with `IP`
* `podList` - replace with `pods`
* `version` - replace with key describing what it belongs to so that it can be compared, for example `dockerVersion`
* `servicePortName` - replace with `portName`
* `svc` - replace with `service`

Example of using context in to distinguish between two same keys:

```go
func ChangePodStatus(newStatus, currentStatus string) {
  klog.InfoS("PodStatusController found pod with status", "status", currentStatus)
  ...
  // Logic that changes status
  ...
  klog.InfoS("PodStatusController changed pod status", "oldStatus", currentStatus, "status", newStatus)
}
```

[lowerCamelCase]: https://en.wiktionary.org/wiki/lowerCamelCase
[alphanumeric]: https://en.wikipedia.org/wiki/Alphanumeric

### Use `klog.KObj` and `klog.KRef` for Kubernetes objects

As part of structured logging migration we want to ensure that kubernetes objects references are consistent within the
codebase. Two new utility functions were introduced to klog `klog.KObj` and `klog.KRef`. Any reference
(name, uid, namespace) to Kubernetes Object (Pod, Node, Deployment, CRD) should be rewritten to utilize those functions.
In situations when object `UID` is would be beneficial for log, it should be added as separate field with `UID` suffix.

For example
```go
func updatePod(pod *covev1.Pod) {
   ...
   klog.Infof("Updated pod %s in namespace %s", pod.Name, pod.Namespace)
}
```
should be changed to
```go
func updatePod(pod *covev1.Pod) {
   ...
   klog.InfoS("Updated pod", "pod", klog.KObj(pod))
}
```
And
```go
func updatePod(pod *covev1.Pod) {
   ...
   klog.Infof("Updated pod with uid: %s", pod.Uid)
}
```
should be changed to
```go
func updatePod(pod *covev1.Pod) {
   ...
   klog.InfoS("Updated pod", "pod", klog.KObj(pod), "podUID", pod.Uid)
}
```

`klog.KObj` requires passing a kubernetes object (struct implementing `metav1.Object` interface). In situations where
the object is not available, we can use `klog.KRef`. Still it is suggested to rewrite the code to use object pointer
instead of strings where possible.

```go
func updatePod(podName, podNamespace string) {
   ...
   klog.InfoS("Updated pod", "pod", klog.KRef(podNamespace, podName))
}
```

For non-namespaced object we can pass empty string to namespace argument

```go
func updateNode(nodeName string) {
   ...
   klog.InfoS("Updated node", "node", klog.KRef("", nodeName))
}
```

### Verify log output

With the introduction of structured functions log arguments will be formatted automatically instead of depending on the
caller. This means that we can remove the burden of picking the format by caller and ensure greater log consistency, but during
migration it's important to ensure that we avoid degradation of log quality. We should ensure that during migration we
preserve properties like:
* meaning of event described by log
* verbosity of stored information

PRs migrating logs should include examples of outputted logs before and after the change, thus helping reviewers
understand the impact of change.

Example code to compare [httplog.go#168](https://github.com/kubernetes/kubernetes/blob/15c3f1b11/staging/src/k8s.io/apiserver/pkg/server/httplog/httplog.go#L168)
```go
package main

import (
	"fmt"
	"k8s.io/klog/v2"
	"net/http"
	"time"
)

type respLogger struct {
	status         int
	statusStack    string
	addedInfo      string
	req *http.Request
}

func (rl *respLogger) Log(latency time.Duration) {
	klog.InfoDepth(1, fmt.Sprintf("verb=%q URI=%q latency=%v resp=%v UserAgent=%q srcIP=%q: %v%v",
		rl.req.Method, rl.req.RequestURI,
		latency, rl.status,
		rl.req.UserAgent(), rl.req.RemoteAddr,
		rl.statusStack, rl.addedInfo,
	))
}

func (rl *respLogger) LogArgs(latency time.Duration) []interface{} {
    return []interface{}{
        "verb", rl.req.Method,
        "URI", rl.req.RequestURI,
        "latency", latency,
        "resp", rl.status,
        "userAgent", rl.req.UserAgent(),
        "srcIP", rl.req.RemoteAddr,
    }
}

func main() {
	klog.InitFlags(nil)

    // Setup
	rl := respLogger{
		status:             200,
		req:                &http.Request{
			Method:           "GET",
			Header:           map[string][]string{"User-Agent": {"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0. 2272.118 Safari/537.36."}},
			RemoteAddr:       "127.0.0.1",
			RequestURI:       "/metrics",
		},
	}
	latency := time.Second

    // Before migration
    rl.Log(latency)

    // After migration
	klog.InfoS("Received HTTP request", rl.LogArgs(latency)...)
}
```

Log output before migration
```
I0528 19:15:22.737538   47512 logtest.go:52] verb="GET" URI="/metrics" latency=1s resp=200 UserAgent="Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0. 2272.118 Safari/537.36." srcIP="127.0.0.1":
```
After
```
I0528 19:15:22.737588   47512 logtest.go:55] "Received HTTP request" verb="GET" URI="/metrics" latency="1s" resp=200 userAgent="Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0. 2272.118 Safari/537.36." srcIP="127.0.0.1"
```
