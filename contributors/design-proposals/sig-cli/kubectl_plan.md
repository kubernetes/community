# Kubectl Plan

## TL;DR

Introduce a new `plan` command similar to `kubectl create --dry-run -o yaml`

- Uses server to convert *flags* to a canonical object representation
  - Client convert flags -> plan request
  - Server convert plan request -> object
  - Server returns object without storing
- `plan` subcommands and help discovered through discovery service (location)
  and openapi (model)

## Abstract

Kubectl supports a number of imperative commands for building Kubernetes resources
on the commandline without interacting with config directly.  While users have
embraced the simplicity and familiarity of using these imperative creation commands,
the commands hide important details about how the system works, and encourage anti-patterns
of interacting directly with objects in the system instead of through local configuration.

Additionally because the imperative creation commands
are completely compiled into the client, they cannot support extension APIs or
behave well in client / server version skewed environments.

Lastly, maintaining creation subcommands requires sustaining engineering efforts
by SIG cli to keep up with changes made on the server side.  Server side
changes frequently break the cli creation commands, requiring the SIG cli
maintainers to track down the cause and sometimes implement a fix.

This document proposes introducing an optional `plan` subresource which converts the
commandline arguments into an object on the server side through a subresource and
simplified representation of the object.  The object is not stored as part of the request,
instead the client must do so through a separate request.

## Goals

- Provide simple cli commands that reflect how the system is designed
- Enable API extensions to expose cli commands without downloading plugins 
- Allow implementers of resource APIs to define cli experience without
  creating PRs in kubectl
- Reduce sustaining engineering burden on SIG to maintain the correctness
  of the commands

## Plan CLI

Kubectl will expose the new `plan` and `plan-list` top-level commands.

When a user runs `kubectl plan -h` it will print a help message
describing the command, and prompt the user to run `plan-list` to view
the available subcommands.

**Note:** Alternatively, we could just make the help message itself dynamic instead
of introducing a `plan-list` command.  We should make this consistent with whatever
we end up doing for `kubectl get`.

`plan` and `plan-list` will use the discovery API to identify all resources with a `plan`
subcommand.  It will lookup the request Model from the openapi, and check that
the model contains a field called `PlanArguments` verify that the
subresource declares that it implements the interface, and is not the result of a
naming collision.  This will also make it easy to add additional fields to the request
that may not map directly to cli flags.

When a user runs `kubectl plan <resource> -h`, kubectl will parse the help message
for the available flags from the fields in `Flags`.
The flag descriptions will be pulled from the field description in the openapi model,
and the names will convert the *camelCase* field names to *dash-delimited* flag names.
The name and namespace of the object will be derived from the url path as they are
done for `get`.

TODO: Find a better name than *Flags* for the field name

## Plan subresource API

Resources can *optionally* expose a `plan` subresource the contains a field named
*Flags*.  The kubectl will expose this subresource directly as a subcommand by
parsing the *Flags* argument into commandline flags based off the openapi model.

```go
type DeploymentPlanRequest struct {
    Flags DeploymentArguments
}

// deploymentArguments contains the values converted from commandline arguments
type DeploymentArguments struct {
    // images is the images to run as separate containers in a Pod
    Images []string
    
    // replicas is the number of replicas to set
    Replicas int32
}
```

The subresource will expose the API as a POST operation.

## Special cases considered

Using binary file contents as field values

- Introduce a `Data` field type that can be read from a file (base64 encoded)
  or other sources.
  
Exposing plan sub-sub-commands - e.g. `plan secret tls`

- This could be supported through sub-sub-reosurces discovered by the client.

## How would this work when the user wants to create the resource?

The user can pipe the plan output to `apply` or `create`.  This has
the added benefit that it will work with all of the tools being developed
around configuration management, and the user can update their object by running
the plan command a second time with different arguments.

`kubectl plan deployment foo --image foo:1 | kubectl apply -f -`

then

`kubectl plan deployment foo --image foo:2 | kubectl apply -f -`

Once diff is supported by apply, users will be able to see differences.

## What about the `create` subcommands and `run` / `expose` commands?

Keep the existing commands, but stop adding new ones and stop adding new
functionality.  Deprecate the `create` commands 1 release after equivalent functionality
is available through `plan`.  Deprecate `run` / `expose` 4 releases after
equivalent functionality is available through `plan`.

**Note:** We can revisit the deprecation plan and determine if it is the correct
thing to do once we have more user feedback.

## What about `run` features like attaching with `-t` & `-i`

We could incorporate these into `kubectl create -f -`

## Alternatives considered

- Using client side plugins instead of exposing this from the server
  - No great distribution logic for client-side plugins
  - Additional step for users
  - May require whitelisting additional binaries for enterprises / regulation / auditing

- Use Kompose for simplified config
  - Doesn't bridge to the convenience of the cli

- Use openapi instead of discovery service for commands
  - Want to be consistent with how other kubectl commands discover resources

- Don't write another command for listing plan resources, include in help.
  - Want to be consistent with what `get` does.
  - Not a good user experience as the list of resource types grows
  - Output can't be consumed by programs
  
- Don't convert the field camelCase to dash-delimited, have the flag name
  explicitly defined through openapi
  - We can add this as an improvement later and default to the conversion behavior

- Do nothing
  rest of the system
  - User confusion on how to use the system
  - Sustaining engineering to continue to maintain and expand existing creation commands

- Don't call the flags struct *Flags*, call it something more generic so it can be reused
  - This is a good point.  Suggestions?