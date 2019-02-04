# Kubectl and SIG CLI Design Principles

## Focus

kubectl provides Resource and Resource Config oriented commands
(as opposed to some other central concepts, such as packaging, integration, etc).
This includes but is not limited to commands to generate, transform, create,
update, delete, watch, print, edit, validate and aggregate information
about Resources and Resource Config.  This functionality may be either
declarative or imperative.

Additionally kubectl provides:
- commands targeted at sub-Resource APIs - e.g. exec, attach, logs
- commands targeted at non-Resource Kubernetes APIs - e.g. openAPI, discovery, version, etc
- porcelain commands for simple / common operations where no discrete
  API implementation exists -e.g. `run`, `expose`, `rollout`, `cp`, `top`, `cordon`,
  `drain` and `describe`.
- porcelain functionality working with Resource Config files, urls, etc - 
  e.g.`-f -R` flags, Kustomization `bases` and `resources`.

*kubectl is part swiss-army knife and part reference implementation for interacting with the API
and driving the fututure direction of the API through identifying API needs and addressing them
client-side.*

As such, it is also a proving group for widely used functionality that may be moved
into the server.  Past examples of kubectl functionality that moved into the server include -
garbage collection, rolling updates, apply, "get" and dry-run.

It may also include porcelain that bridges standard non-Kubernetes native solution to Kubernetes
native solutions - e.g. `docker run` -> `kubectl run`, `EXPOSE` -> `kubectl expose`.

## Workflows

The scope of CLI Tools focuses on enabling declarative and imperative workflows
for invoking kubernetes APIs and authoring Resource Config.  Tools provide
commands for both generalized (e.g. create resource from Resource Config) tasks and
specialized (e.g. drain a node, exec into a container) tasks.

It is the philosophy of the tools developed in SIG CLI to facilitate working
directly with the Kubernetes APIs and Kubernetes style Resources, and to the
extent possible, provide a transparent experience for how commands map to
Kubernetes APIs and Resources.

Building new abstractions and concepts for users to interact with in place of
the Resource APIs rather than access them (e.g. through generalized templating,
DSLs, etc) is not a goal of SIG CLI.

## Extensibility

CLI prefers to develop commands in such a way that they can provide a native
experience for APIs developed as extensions.  This requires a philosophy of
minimizing resource specific functionality and enabling it through data
published by the cluster rather than hard-coding the API data into the tools.
This includes developing specific extension mechanisms for kubectl such as plugins.
Extensibility is a design preference, not a mandate, and should not come at a practical
cost impacting the UX or functionality of the tool.

CLI prefers to develop commands in such a way that enables tools and solutions
developed independently (e.g. outside the SIG, K8S project, etc) to interoperate
with the CLI tools - e.g. through pipes or wrapping / execing.  This is aligned
with the goal of remaining close to the Kubernetes APIs.

## Documentation

SIG CLI is responsible for developing documentation to accompany kubectl that both describes
the functionality and provides techniques for effective usage.

#### Examples of Functionality In Kubectl

Following are examples of functionality that is in kubectl.

- Invoking Kubernetes APIs: Resource APIs, SubResource APIs, Discovery Service, Version, OpenAPI
- Pre and Post processing API Resource Config, API Requests and API Responses
- Aggregating multiple API Responses and post processing them
- Collapsing multiple manual steps into a command
- Generating Kubernetes Resource Config locally or creating Resources remotely
- Transforming Kubernetes Resource Config locally or patching remotely
- Blocking on propagation of an event or change to the cluster
- Referencing a collection of either remote or local Resource Config
- Configure how to talk to a specific cluster from the cli
- Selecting which API group/version to invoke if ambiguous in the context of the command
