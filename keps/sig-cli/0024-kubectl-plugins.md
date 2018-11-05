---
kep-number: 24
title: Kubectl Plugins 
authors:
  - "@juanvallejo"
owning-sig: sig-cli
participating-sigs:
  - sig-cli
reviewers:
  - "@pwittrock"
  - "@deads2k"
  - "@liggitt"
  - "@soltysh"
approvers:
  - "@pwittrock"
  - "@soltysh"
editor: juanvallejo
creation-date: 2018-07-24
last-updated: 2018-08-09
status: provisional
see-also:
  - n/a
replaces:
  - "https://github.com/kubernetes/community/blob/master/contributors/design-proposals/cli/kubectl-extension.md"
  - "https://github.com/kubernetes/community/pull/481"
superseded-by:
  - n/a
---

# Kubectl Plugins 

## Table of Contents

* [Table of Contents](#table-of-contents)
* [Summary](#summary)
* [Motivation](#motivation)
    * [Limitations of the Existing Design](#limitations-of-the-existing-design)
    * [Goals](#goals)
    * [Non-Goals](#non-goals)
* [Proposal](#proposal)
    * [Scenarios](#scenarios)
    * [Implementation Details/Design/Constraints](#implementation-detailsdesign)
        * [Naming Conventions](#naming-conventions)
    * [Implementation Notes/Constraints](#implementation-notesconstraints)
    * [Risks and Mitigations](#risks-and-mitigations)
* [Graduation Criteria](#graduation-criteria)
* [Implementation History](#implementation-history)
* [Drawbacks](#drawbacks)
* [Future Improvements/Considerations](#future-improvementsconsiderations)

## Summary

This proposal introduces the main design for a plugin mechanism in `kubectl`.
The mechanism is a git-style system, that looks for executables on a user's `$PATH` whose name begins with `kubectl-`.
This allows plugin binaries to override existing command paths and add custom commands and subcommands to `kubectl`.

## Motivation

The main motivation behind a plugin system for `kubectl` stems from being able to provide users with a way to extend
the functionality of `kubectl`, beyond what is offered by its core commands.

By picturing the core commands provided by `kubectl` as essential building blocks for interacting with a Kubernetes
cluster, we can begin to think of plugins as a means of using these building blocks to provide more complex functionality.
A new command, `kubectl set-ns`, for example, could take advantage of the rudimentary functionality already provided by
the `kubectl config` command, and build on top of it to provide users with a powerful, yet easy-to-use way of switching
to a new namespace.

For example, the user experience for switching namespaces could go from:

```bash
kubectl config set-context $(kubectl config current-context) --namespace=mynewnamespace
```

to:

```
kubectl set-ns mynewnamespace
```

where `set-ns` would be a user-provided plugin which would call the initial `kubectl config set-context ...` command
and set the namespace flag according to the value provided as the plugin's first parameter.

The `set-ns` command above could have multiple variations, or be expanded to support subcommands with relative ease.
Since plugins would be distributed by their authors, independent from the core Kubernetes repository, plugins could
release updates and changes at their own pace.

### Limitations of the Existing Design

The existing alpha plugin system in `kubectl` presents a few limitations with its current design. 
It forces plugin scripts and executables to exist in a pre-determined location, requires a per-plugin metadata file for
interpretation, and does not provide a clear way to override existing command paths or provide additional subcommands 
without having to override a top-level command.

The proposed git-style re-design of the plugin system allows us to implement extensibility requests from users that the
current system is unable to address.
See https://github.com/kubernetes/kubernetes/issues/53640 and https://github.com/kubernetes/kubernetes/issues/55708.

### Goals

* Avoid any kind of installation process (no additional config, users drop an executable in their `PATH`, for example, 
  and they are then able to use that plugin with `kubectl`).
  No additional configuration is needed, only the plugin executable.
  A plugin's filename determines the plugin's intention, such as which path in the command tree it applies to:
  `/usr/bin/kubectl-educate-dolphins` would, for example be invoked under the command `kubectl educate dolphins --flag1 --flag2`.
  It is up to a plugin to parse any arguments and flags given to it. A plugin decides when an argument is a
  subcommand, as well as any limitations or constraints that its flags should have.
* Relay all information given to `kubectl` (via command line args) to plugins as-is.
  Plugins receive all arguments and flags provided by users and are responsible for adjusting their behavior
  accordingly.
* Provide a way to limit which command paths can and cannot be overriddden by plugins in the command tree.

### Non-Goals

* The new plugin mechanism will not be a "plugin installer" or wizard. It will not have specific or baked-in knowledge 
  regarding a plugin's location or composition, nor will it provide a way to download or unpack plugins in a correct 
  location.
* Plugin discovery is not a main focus of this mechanism. As such, it will not attempt to collect data about every 
  plugin that exists in an environment.
* Plugin management is out of the scope of this design. A mechanism for updating and managing lifecycle of existing 
  plugins should be covered as a separate design (See https://github.com/kubernetes/community/pull/2340).
* Provide a standard package of common cli utilities that is consumed by `kubectl` and plugins alike.
  This should be done as an independent effort of this plugin mechanism.

## Proposal

### Scenarios

* Developer wants to create and expose a plugin to `kubectl`.
  They use a programming language of their choice and create an executable file.
  The executable's filename consists of the command path to implement, and is prefixed with `kubectl-`.
  The executable file is placed on the user's `PATH`.

### Implementation Details/Design

The proposed design passes through all environment variables, flags, input, and output streams exactly as they are given 
to the parent `kubectl` process. This has the effect of letting plugins run without the need for any special parsing
or case-handling in `kubectl`.

In essence, a plugin binary must be able to run as a standalone process, completely independent of `kubectl`.

* When `kubectl` is executed with a subcommand _foo_ that does not exist in the command tree, it will attempt to look
for a filename `kubectl-foo` (`kubectl-foo.exe` on Windows) in the user's `PATH` and execute it, relaying all arguments given
as well as all environment variables to the plugin child-process.

A brief example (not an actual prototype) is provided below to clarify the core logic of the proposed design:

```go
// treat all args given by the user as pieces of a plugin binary's filename
// and short-circuit once we find an arg that appears to be a flag.
remainingArgs := []string{} // all "non-flag" arguments

for idx := range cmdArgs {
	if strings.HasPrefix(cmdArgs[idx], "-") {
		break
	}
	remainingArgs = append(remainingArgs, strings.Replace(cmdArgs[idx], "-", "_", -1))
}

foundBinaryPath := ""

// find binary in the user's PATH, starting with the longest possible filename
// based on the given non-flag arguments by the user
for len(remainingArgs) > 0 {
	path, err := exec.LookPath(fmt.Sprintf("kubectl-%s", strings.Join(remainingArgs, "-")))
	if err != nil || len(path) == 0 {
		remainingArgs = remainingArgs[:len(remainingArgs)-1]
		continue
	}

	foundBinaryPath = path
	break
}

// if we are able to find a suitable plugin executable, perform a syscall.Exec call
// and relay all remaining arguments (in order given), as well as environment vars.
syscall.Exec(foundBinaryPath, append([]string{foundBinaryPath}, cmdArgs[len(remainingArgs):]...), os.Environ())
```

#### Naming Conventions

Under this proposal, `kubectl` would identify plugins by looking for filenames beginning with the `kubectl-` prefix.
A search for these names would occur on a user's `PATH`. Only files that are executable and begin with this prefix
would be identified.

### Implementation Notes/Constraints

The current implementation details for the proposed design rely on using a plugin executable's name to determine what
command the plugin is adding.
For a given command `kubectl foo --bar baz`, an executable `kubectl-foo` will be matched on a user's `PATH`,
and the arguments `--bar baz` will be passed to it in that order.

A potential limitation of this could present itself in the order of arguments provided by a user.
A user could intend to run a plugin `kubectl-foo-bar` with the flag `--baz` with the following command
`kubectl foo --baz bar`, but instead end up matching `kubectl-foo` with the flag `--baz` and the argument `bar` based
on the placement of the flag `--baz`.

A notable constraint of this design is that it excludes any form of plugin lifecycle management, or version compatibility.
A plugin may depend on other plugins based on the decision of a plugin author, however the proposed design does nothing
to facilitate such dependencies. It is up to the plugin's author (or a separate / independent plugin management system) to
provide documentation or instructions on how to meet any dependencies required by a plugin.

Further, with the proposed design, plugins that rely on multiple "helper" files to properly function, should provide an
"entrypoint" executable (which is placed on a user's `PATH`), with any additional files located elsewhere (e.g. ~/.kubeplugins/myplugin/helper1.py).

### Risks and Mitigations

Unlike the existing alpha plugin mechanism, the proposed design does not constrain commands added by plugins to exist as subcommands of the
`kubectl plugin` design. Commands provided by plugins under the new mechanism can be invoked as first-class commands (`/usr/bin/kubectl-foo` provides the `kubectl foo` parent command).

A potential risk associated with this could present in the form of a "land-rush" by plugin providers.
Multiple plugin authors would be incentivized to provide their own version of plugin `foo`.
Users would be at the mercy of whichever variation of `kubectl-foo` is discovered in their `PATH` first when executing that command.

A way to mitigate the above scenario would be to have users take advantage of the proposed plugin mechanism's design by renaming multiple variations of `kubectl-foo`
to include the provider's name, for example: `kubectl-acme-foo`, or `kubectl-companyB-foo`.

Conflicts such as this one could further be mitigated by a plugin manager, which could perform conflict resolution among similarly named plugins on behalf of a user.

## Graduation Criteria

* Make this mechanism a part of `kubectl`'s command-lookup logic.

## Implementation History

This plugin design closely follows major aspects of the plugin system design for `git`.

## Drawbacks

Implementing this design could potentially conflict with any ongoing work that depends on the current alpha plugin system.

## Future Improvements/Considerations

The proposed design is flexible enough to accommodate future updates that could allow certain command paths to be overwritten
or extended (with the addition of subcommands) via plugins.
