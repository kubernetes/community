# Structured and Contextual Logging migration instructions

This document describes instructions for the migration proposed by [Structured Logging KEP] and [Contextual Logging KEP]. It describes new
functions introduced in `klog` (Kubernetes logging library) and how log calls should be changed to utilize new features.
This document was written for the initial migration of `kubernetes/kubernetes` repository proposed for Alpha stage of structured logging, but
should be applicable at later stages or for other projects using `klog` logging library.

[Structured Logging KEP]: https://github.com/kubernetes/enhancements/tree/master/keps/sig-instrumentation/1602-structured-logging
[Contextual Logging KEP]: https://github.com/kubernetes/enhancements/tree/master/keps/sig-instrumentation/3077-contextual-logging

## How to contribute

### About the migration

We would like for the Kubernetes community to settle on one preferred log message structure and log calls as defined by [logr].
The goal of the migration is to switch C-like format string logs to structured logs with explicit metadata about parameters and
to remove the global logger in klog with a `logr.Logger` instance that gets chosen by the caller of a function.

Migration within the structured logging working group happens in two different ways - organized & non-organized. 

With organized, milestone-based, large-scale migration we try to target deliverables for a specific Kubernetes release. A good example of such an effort is the migration of the entire Kubelet code in the [#1.21 release](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.21.md#structured-logging-in-kubelet) 

For non-organized migrations, the onus is, generally, on the individual contributors.

#### Organized Migration

Organized migration is carried out in a two state cycle aligning with the cadence of Kubernetes releases. 

In the first stage, we pick a particular migration milestone & create an issue to split the work into smaller chunks (for example, migrating just a single file). This ensures that the work can be easily picked up by multiple individual contributors and also avoids conflicts. After the migration activity is complete, we mark the directory as migrated to avoid regressions with the help of tooling that we have developed. 

In the second milestone, we take a break from migration to analyze results & improve our existing processes. Adding structured information to logs is a very costly & time-consuming affair. Setting aside time in the second stage to collect feedback, analyze the impact of changes on the log volume & performance, and better our PR review process helps us avoid mistakes and duplicate efforts.

#### Non-organized migration

As aforementioned, our non-organized migration efforts are spearheaded by individual contributors who need to migrate particular code sections to utilize new features early. 

Efforts for Kubernetes components not yet marked for migration also fall under this category & we will try our best to review and accept as many PRs as we can, but being a small team, we can't give any time frames for the same. 

The respective component owners have a final say in the acceptance of these contributions. Since this is a non-organized effort there is a high possibility that two contributors could be working on migrating the same code. 

Before sending a PR our way, please ensure that there isn't one already in place created by someone else.

### Current status

* 1.21 Kubelet was migrated
* 1.22 Collected feedback and improved the process.
* 1.23 kube-scheduler and kube-proxy were migrated.
* 1.24 Contextual logging infrastructure (updated klog, component-base/logs enables it) in place.

## Sending a Structured Logging Migration Pull Request

### Before creating the Pull Request

* In case of any questions, please contact us on [#wg-structured-logging](https://kubernetes.slack.com/messages/wg-structured-logging) Slack channel.
* Check list of [opened PRs](https://github.com/kubernetes/kubernetes/pulls?q=is%3Aopen+label%3Awg%2Fstructured-logging+is%3Apr) to understand if someone is already working on this.
* Reference: [Pull Request Process](https://www.kubernetes.dev/docs/guide/pull-requests/)

### What to include in the Pull request

To ensure that the PRs are reviewed in a timely manner, we require authors to provide the below information. This helps in proper triaging & reviewing by the concerned knowledgeable parties.

* title: `Migrate <directory/file> to structured logging`
* WGs: `structured-logging`
* area: `logging`
* priority: `important-longterm` for normal PR or `important-soon` for PRs that are part of organized migration of component. 
* cc: @kubernetes/wg-structured-logging-reviews
* kind: cleanup
* release note: `Migrate <directory/file> to structured logging
`
To quickly add this information you can comment on PR with below commands:

```
/retitle Migrate <directory/file> to structured logging
/wg structured-logging 
/area logging
/priority important-longterm
/kind cleanup
/cc @kubernetes/wg-structured-logging-reviews
```
And edit the top comment to include release note:
Migrate <directory/file> to structured logging

### Why my PR was rejected?

Even though the Kubernetes project is organizing migration of logging, this doesn't mean that we are able to accept all the PRs that come our way. We reserve the right to reject the Pull Request in situations listed below:

* Pull request is below minimum quality standards and clearly shows that author hasn't read the guide at all, for example PR just renames `Infof` to `InfoS`.
* Pull request migrates components that the owners have decided against migrating. List of those components:
   * kubeadm

 

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

// InfoSDepth acts as InfoS but uses depth to determine which call frame to log.
// InfoSDepth(0, "msg") is the same as InfoS("msg").
func InfoSDepth(depth int, msg string, keysAndValues ...interface{}) 

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

// ErrorSDepth acts as ErrorS but uses depth to determine which call frame to log.
// ErrorSDepth(0, "msg") is the same as ErrorS("msg").
func ErrorSDepth(depth int, err error, msg string, keysAndValues ...interface{})

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

// KObjSlice takes a slice of objects that implement the KMetadata interface
// and returns an object that gets logged as a slice of ObjectRef values or a
// string containing those values, depending on whether the logger prefers text
// output or structured output.
// >> klog.InfoS("Pods status updated", "pods", klog.KObjSlice(pods), "status", "ready")
// output:
// >> I1025 00:15:15.525108       1 controller_utils.go:116] "Pods status updated" pods="[kube-system/kubedns kube-system/metrics-server]" status="ready"
func KObjSlice(arg interface{}) interface{} 

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

## Contextual logging in Kubernetes

Contextual logging builds on top of structured logging because the parameters
for individual log calls are the same. The difference is that different
functions need to be used:

- `klog.ErrorS` -> `logger.Error`
- `klog.InfoS` -> `logger.Info`
- `klog.V().InfoS` -> `logger.V().Info`
- `klog.V().ErrorS(err, ...)` -> `logger.V().Info(..., "err", err)`

In all of these cases, `logger` is a `logr.Logger` instance. `klog.Logger` is
an alias for that type. In contrast to klog, logr ignores verbosity for `Error`
calls and always emits the log. The semantic equivalent of `klog.V().ErrorS` is
therefore an `Info` call where the error is represented as a key/value pair.

Determining where the `Logger` instance comes from is the main challenge when
migrating code to contextual logging. Several new klog functions help with
that:

- [`klog.FromContext`](https://pkg.go.dev/k8s.io/klog/v2#FromContext)
- [`klog.Background`](https://pkg.go.dev/k8s.io/klog/v2#Background)
- [`klog.TODO`](https://pkg.go.dev/k8s.io/klog/v2#TODO)

The preferred approach is to retrieve the instance with `klog.FromContext` from
a `ctx context` parameter. If a function or method does not have one, consider
adding it. This API change then implies that all callers also need to be
updated. If there are any `context.TODO` calls in the modified functions,
replace with the new `ctx` parameter.

In performance critical code it may be faster to add a `logger klog.Logger`
parameter. This needs to be decided on a case-by-case basis.

When such API changes trickle up to a unit test, then enable contextual logging
with per-test output with
[`ktesting`](https://pkg.go.dev/k8s.io/klog/v2@v2.60.1/ktesting):

```go
import (
    "testing"

    "k8s.io/klog/v2/ktesting"
    _ "k8s.io/klog/v2/ktesting/init" # Add command line flags for ktesting.
)

func TestFoo(t *testing.T) {
   _, ctx := ktesting.NewTestContext(t)
   doSomething(ctx)
}
```

If a logger instance is needed instead, then use:

```go
   logger, _ := ktesting.NewTestContext(t)
```

The KEP has further instructions about the [transition to contextual
logging](https://github.com/kubernetes/enhancements/tree/master/keps/sig-instrumentation/3077-contextual-logging#transition). It
also lists several
[pitfalls](https://github.com/kubernetes/enhancements/tree/master/keps/sig-instrumentation/3077-contextual-logging#pitfalls-during-usage)
that developers and reviewers need to be aware of.

## Migration

1. Optional: find code locations that need to be changed:

   - `(cd hack/tools && go install k8s.io/klog/hack/tools/logcheck)`
   - structured logging: `$GOPATH/bin/logcheck -check-structured ./pkg/controller/...`
   - contextual logging: `$GOPATH/bin/logcheck -check-contextual ./pkg/scheduler/...`

1. Change log functions to structured or (better!) contextual equivalent
1. Remove string formatting from log message
1. Name arguments
1. Ensure that value is properly passed
1. Verify log output
1. Prevent re-adding banned functions after migration

## Change log functions

Structured logging functions follow a different logging interface design than other functions in `klog`. They follow
minimal design from [logr] thus there is no one-to-one mapping.

Simplified mapping between functions for structured logging:
* `klog.Infof`, `klog.Info`, `klog.Infoln` -> `klog.InfoS`
* `klog.InfoDepth` -> `klog.InfoSDepth`
* `klog.V(N).Infof`, `klog.V(N).Info`, `klog.V(N).Infoln` -> `klog.V(N).InfoS`
* `klog.Warning`, `klog.Warningf`, `klog.Warningln` -> `klog.InfoS`
* `klog.WarningDepth` -> `klog.InfoSDepth`
* `klog.Error`, `klog.Errorf`, `klog.Errorln` -> `klog.ErrorS`
* `klog.ErrorDepth` -> `klog.ErrorSDepth`
* `klog.Fatal`, `klog.Fatalf`, `klog.Fatalln` -> `klog.ErrorS` followed by `klog.FlushAndExit(klog.ExitFlushTimeout, klog.1)` ([see below])
* `klog.FatalDepth` -> `klog.ErrorDepth` followed by `klog.FlushAndExit(klog.ExitFlushTimeout, klog.1)` ([see below])

For contextual logging, replace furthermore:

- `klog.ErrorS` -> `logger.Error`
- `klog.InfoS` -> `logger.Info`
- `klog.V().InfoS` -> `logger.V().Info`

[see below]: #replacing-fatal-calls

### Using ErrorS or logger.Error

With `klog` structured logging borrowing the interface from [logr] it also inherits it's differences in semantic of
error function. Logs generated by `ErrorS` command may be enhanced with additional debug information
(such as stack traces) or be additionally sent to special error recording tools. Errors should be used to indicate
unexpected behaviours in code, like unexpected errors returned by subroutine function calls.
In contrast to info log calls, error log calls always record the log entry, regardless of the current verbosity
settings.

Calling `ErrorS` with `nil` as error is acceptable if there is an error condition that deserves a stack trace at this
origin point or always must be logged. For expected errors (`errors` that can happen during routine operations) please consider using
`klog.InfoS` and pass error in `err` key instead.

### Replacing Fatal calls

Use of Fatal should be discouraged and it's not available in new functions. Instead of depending on the logger to exit
the process, you should:
- rewrite code to return an `error` and let the caller deal with it or, if that is not feasible,
- log and exit separately.

`os.Exit` should be avoided because it skips log data flushing. Instead use
[`klog.FlushAndExit`](https://pkg.go.dev/k8s.io/klog/v2#FlushAndExit). The first
parameter determines how long the program is allowed to flush log data before
`os.Exit` is called. If unsure, use `klog.ExitFlushTimeout`, the value used
by `klog.Fatal`.

Fatal calls use a default exit code of 255. When migrating, please use an exit code of 1 and include an "ACTION REQUIRED:" release note.

For example
```go
func validateFlags(cfg *config.Config, flags *pflag.FlagSet) error {
	if err := cfg.ReadAndValidate(flags); err != nil {
		klog.FatalF("Error in reading and validating flags %s", err)
	}
}
```
should be changed to
```go
func validateFlags(cfg *config.Config, flags *pflag.FlagSet) error {
	if err := cfg.ReadAndValidate(flags); err != nil {
		klog.ErrorS(err, "Error in reading and validating flags")
		klog.FlushAndExit(klog.ExitFlushTimeout, 1)
	}
}
```

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

## Good practice for passing values in structured logging

When passing a value for a key-value pair, please use following rules:
* Prefer using Kubernetes objects and log them using `klog.KObj` or `klog.KObjSlice`
  * When the original object is not available, use `klog.KRef` instead
  * when only one object (for example `*v1.Pod`), we use`klog.KObj` 
  * When type is object slice (for example `[]*v1.Pod`), we use `klog.KObjSlice`
* Pass structured values directly (avoid calling `.String()` on them first)
* When the goal is to log a `[]byte` array as string, explicitly convert with `string(<byte array>)`.

### Prefer using Kubernetes objects (for example `*v1.Pod` or `[]*v1.Pod`) and log them using `klog.KObj` or `klog.KObjSlice`

As part of the structured logging migration, we want to ensure that Kubernetes object references are
consistent within the
codebase. Three new utility functions were introduced to klog: `klog.KObj` `klog.KObjSlice` and `klog.KRef`.

Any existing logging code that makes a reference (such as name, namespace) to a Kubernetes
object (for example: Pod, Node, Deployment, CustomResourceDefinition) should be rewritten to
utilize the new functions.

Logging using this method will ensure that log output include a proper and correctly formatted reference
to that object. The formatting / serialization is automatically selected depending on the output log format.
For example, the same object could be represented as `<namespace>/<name>` in
plain text and as `{"namespace": "<namespace>", "name": "<name>"}` in JSON.
In situations when object `UID` is would be beneficial for log, it should be added as separate key with `UID` suffix.

For example
```go
func updatePod(pod *covev1.Pod) {
   ...
   klog.Infof("Updated pod(%s) with name %s in namespace %s", pod.Uid, pod.Name, pod.Namespace)
}
```
should be changed to
```go
func updatePod(pod *covev1.Pod) {
   ...
   klog.InfoS("Updated pod", "pod", klog.KObj(pod), "podUID", pod.Uid)
}
```

### When the original object is not available, use `klog.KRef` instead
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

### Pass structured values directly

By passing whole structure without any marshalling we can allow logging library to make the decision for us.
This is especially beneficial when Kubernetes supports different log format, so the logging library can make decision based on that.
For example using `<namespace>/<name>` when writing in text format and `{"namespace": "<namespace>", "name": "<name>"}` for json format.

## Verify log output

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

## Prevent re-adding banned functions after migration

After a package has been migrated to structured and/or contextual logging, we
want to ensure that all log calls that get added in future PRs are structured
resp. contextual.

For structured logging, the list of migrated packages in
[`hack/logcheck.conf`](https://github.com/kubernetes/kubernetes/blob/b9792a9daef4d978c5c30b6d10cbcdfa77a9b6ac/hack/logcheck.conf#L16-L22)
can be extended.

For contextual logging, a new list can be added at the bottom once the code is
ready, with content like this:

```
# Packages that have been migrated to contextual logging:
contextual k8s.io/kubernetes/pkg/scheduler/.*
```

The corresponding line with `structured k8s.io/kubernetes/pkg/scheduler/.*`
then is redundant and can be removed because "contextual" implies "structured".

Both lists should be sorted alphabetically. That reduces the risk of code
conflicts and makes the file more readable.
