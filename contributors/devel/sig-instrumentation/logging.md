## Logging

This document provides an overview of  the recommended way to develop and implement
logging for components of Kubernetes.
For [k/kubernetes](https://github.com/kubernetes/kubernetes) these conventions should be seen
as a strong recommendation, however any Kubernetes or external project is welcome to follow
the guidelines below. Most of the ideas can be applied more broadly.

### Logging in Kubernetes

Kubernetes project uses [klog](https://github.com/kubernetes/klog) for logging.
This library was created as a permanent fork of glog, as the original project was no longer developed. 
With the introduction of [Structured Logging](https://github.com/kubernetes/enhancements/tree/master/keps/sig-instrumentation/1602-structured-logging) and [Contextual Logging](https://github.com/kubernetes/enhancements/blob/master/keps/sig-instrumentation/3077-contextual-logging/README.md), Kubernetes is being migrated to use [logr](https://github.com/go-logr/logr) as its logging interface. Any changes in klog are done to ensure a smooth transition to `logr`. Klog is then only used to manage the logging configuration (`SetLogger`, retrieving the logger) and its logging calls should not be used anymore.
``

For projects that want to integrate with Kubernetes or write the same log format, retrieving a `logr.Logger` instance from klog will return an implementation that emits log output the same way as the traditional klog logging calls.

### Logging Conventions

### When not to log?

Shared libraries, such as [`client-go`](https://github.com/kubernetes/client-go), should _not_
log errors themselves
but just return `error`, because client libraries may be used in CLI UIs that wish to control output.

### How to log

There are two main klog methods for writing logs: `klog.InfoS` and `klog.ErrorS`. Both methods are part of a [logr](https://github.com/go-logr/logr) compatible interface.
For those two functions exists additional flavors, like `klog.V(2).InfoS` that allows caller to increase log verbosity (making it only available when `--v=2` flag is provided to binary) and `klog.InfoSDepth` with `klog.ErorSDepth` that give control to change caller information.

Any other non-structured (without `S`) klog methods like `klog.Infof` that use old C like string format interface are no longer recommended for use. 
Those log calls using those methods should be migrated to their structured variants.
Please see the [Structured Logging Guide](migration-to-structured-logging.md#structured-logging-in-kubernetes) for more information on how to migrate such logs.

All structured logging methods accept the log message as a string, along with any number of key/value pairs that you provide via a variadic `interface{}` argument.
As variadic arguments represent key value pairs, they should always be even in count with first element being key of type string and second value of any type matching that key.

Prototype:
```go
func InfoS(message string, keysAndValues ...interface{})
func InfoSDepth(depth int, message string, keysAndValues ...interface{})
func ErrorS(err error, message string, keysAndValues ...interface{})
func ErrorSDepth(depth int, err error, message string, keysAndValues ...interface{})
```

Example:
```go
klog.InfoS("Received HTTP request", "method", "GET", "URL", "/metrics", "latency", time.Second)
```

### Message style guidelines

Good practices when writing log messages are:

- Start from a capital letter.
- Do not end the message with a period.
- Use active voice. Use complete sentences when there is an acting subject ("A could not do B") or omit the subject if the subject would be the program itself ("Could not do B").
- Use past tense ("Could not delete B" instead of "Cannot delete B")
- When referring to an object, state what type of object it is. ("Deleted pod" instead of "Deleted")

### What method to use?

* `klog.ErrorS`, [`logr.Logger.Error`](https://pkg.go.dev/github.com/go-logr/logr#Logger.Error):
  Use error logs to inform admins that they might have to do something to fix a problem.
  If there is no `error` instance, pass nil and use the message string and key/value pairs to describe the problem instead of fabricating an error with `fmt.Errorf`.
  If the dependency is allowed, use `runtime.HandleErrorWithContext` or `runtime.HandleErrorWithLogger` instead.

  Don't emit an error log before returning an error because
  it is usually uncertain if and how the caller is going to handle the returned error.
  It might handle it gracefully, in which case alerting an administrator with an error log would be wrong.
  Instead, use `fmt.Errorf` with `%w` to return a more informative error with a trail for where the error came from.
  Avoid "failed to" as prefix when wrapping errors, it quickly becomes repetitive when added multiple times.

  Sometimes it may be useful to log an error for debugging purposes before returning it. Use an info log with at least
  `V(4)` and `err` as key for the key/value pair.

* [`runtime.HandleErrorWithContext`](https://pkg.go.dev/k8s.io/apimachinery@v0.34.2/pkg/util/runtime#HandleErrorWithContext),
  [`runtime.HandleErrorWithLogger`](https://pkg.go.dev/k8s.io/apimachinery@v0.34.2/pkg/util/runtime#HandleErrorWithLogger):
  Whenever possible, use this instead of normal error logging at the top of a call chain, i.e. the place where code
  cannot return an error up to its caller, if there is no other way of handling the error (like logging it
  and then exiting the process).

  The `runtime` package supports additional error handling mechanisms that downstream consumers of Kubernetes may add
  (see [ErrorHandlers](https://pkg.go.dev/k8s.io/apimachinery@v0.34.2/pkg/util/runtime#pkg-variables)).
  Some [downstream consumers](https://grep.app/search?q=runtime.ErrorHandlers) even suppress logging of errors completely.

* `klog.InfoS` -  Structured logs to the INFO log. `InfoS` should be used for routine logging. It can also be used to log warnings for expected errors (errors that can happen during routine operations).
  Depending on log severity it's important to pick a proper verbosity level to ensure that consumer is neither under nor overwhelmed by log volume:
  * `klog.V(0).InfoS` = `klog.InfoS` - Generally useful for this to **always** be visible to a cluster operator
    * Programmer errors
    * Logging extra info about a panic
    * CLI argument handling
  * `klog.V(1).InfoS` - A reasonable default log level if you don't want verbosity.
    * Information about config (listening on X, watching Y)
    * Errors that repeat frequently that relate to conditions that can be corrected (pod detected as unhealthy)
  * `klog.V(2).InfoS` - Useful steady state information about the service and important log messages that may correlate to significant changes in the system.  This is the recommended default log level for most systems.
    * Logging HTTP requests and their exit code
    * System state changing (killing pod)
    * Controller state change events (starting pods)
    * Scheduler log messages
  * `klog.V(3).InfoS` - Extended information about changes
    * More info about system state changes
  * `klog.V(4).InfoS` - Debug level verbosity
    * Logging in particularly thorny parts of code where you may want to come back later and check it
  * `klog.V(5).InfoS` - Trace level verbosity
    * Context to understand the steps leading up to errors and warnings
    * More information for troubleshooting reported issues

As per the comments, the practical default level is V(2). Developers and QE
environments may wish to run at V(3) or V(4). If you wish to change the log
level, you can pass in `-v=X` where X is the desired maximum level to log.

## Logging Formats

Kubernetes supports multiple logging formats to allow users flexibility in picking a format that best matches their needs.
As new features and logging formats are proposed, we expect that the Kubernetes community might find itself in disagreement about what formats and features should be officially supported.
This section will serve as a developer documentation describing high level goals of each format, giving a holistic vision of logging options and allowing the community to resolve any future disagreement when making changes.

### Text logging format

* Default format in Kubernetes.
* Purpose:
  * Maintain backward compatibility with klog format.
  * Human-readable
  * Development
* Features:
  * Minimal support for structured metadata - required to maintain backward compatibility.
  * Marshals object using fmt.Stringer interface allowing to provide a human-readable reference. For example `pod="kube-system/kube-dns"`.
  * Support for multi-line strings with actual line breaks - allows dumping whole yaml objects for debugging.

Logs written using text format will always have standard header, however message format differ based on whether string formatting (Infof, Errorf, ...) or structured logging method was used to call klog (InfoS, ErrorS).

Examples, first written using Infof, in second InfoS was used.
```
I0528 19:15:22.737538   47512 logtest.go:52] Pod kube-system/kube-dns status was updated to ready
I0528 19:15:22.737538   47512 logtest.go:52] "Pod status was updated" pod="kube-system/kube-dns" status="ready"
```

Explanation of header:
```
Lmmdd hh:mm:ss.uuuuuu threadid file:line] msg...

where the fields are defined as follows:
	L                A single character, representing the log level (eg 'I' for INFO)
	mm               The month (zero padded; ie May is '05')
	dd               The day (zero padded)
	hh:mm:ss.uuuuuu  Time in hours, minutes and fractional seconds
	threadid         The space-padded thread ID as returned by GetTID()
	file             The file name
	line             The line number
	msg              The user-supplied message
```

See more in [here](https://github.com/kubernetes/klog/blob/9ad246211af1ed84621ee94a26fcce0038b69cd1/klog.go#L581-L597)

### JSON logging format

* Requires passing `--logging-format=json` to enable
* Purpose:
  * Provide structured metadata in efficient way.
  * Optimized for efficient log consumption, processing and querying.
  * Machine-readable
* Feature:
  * Full support for structured metadata, using `logr.MarshalLog` when available.
  * Multi-line strings are quoted to fit into a single output line.

Example:
```json
{"ts": 1580306777.04728,"v": 4,"msg": "Pod status was updated","pod":{"name": "kube-dns","namespace": "kube-system"},"status": "ready"}
```

Keys with special meaning:
* ts - timestamp as Unix time (required, float)
* v - verbosity (only for info and not for error messages, int)
* err - error string (optional, string)
* msg - message (required, string)
