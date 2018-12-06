## Logging Conventions

The following conventions for the klog levels to use.
[klog](http://godoc.org/github.com/kubernetes/klog) is globally preferred to
[log](http://golang.org/pkg/log/) for better runtime control.

* klog.Errorf() - Always an error

* klog.Warningf() - Something unexpected, but probably not an error

* klog.Infof() has multiple levels:
  * klog.V(0) - Generally useful for this to ALWAYS be visible to an operator
    * Programmer errors
    * Logging extra info about a panic
    * CLI argument handling
  * klog.V(1) - A reasonable default log level if you don't want verbosity.
    * Information about config (listening on X, watching Y)
    * Errors that repeat frequently that relate to conditions that can be corrected (pod detected as unhealthy)
  * klog.V(2) - Useful steady state information about the service and important log messages that may correlate to significant changes in the system.  This is the recommended default log level for most systems.
    * Logging HTTP requests and their exit code
    * System state changing (killing pod)
    * Controller state change events (starting pods)
    * Scheduler log messages
  * klog.V(3) - Extended information about changes
    * More info about system state changes
  * klog.V(4) - Debug level verbosity
    * Logging in particularly thorny parts of code where you may want to come back later and check it
  * klog.V(5) - Trace level verbosity
    * Context to understand the steps leading up to errors and warnings
    * More information for troubleshooting reported issues

As per the comments, the practical default level is V(2). Developers and QE
environments may wish to run at V(3) or V(4). If you wish to change the log
level, you can pass in `-v=X` where X is the desired maximum level to log.
