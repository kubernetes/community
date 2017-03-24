# MVP - kubectl commands for the Service Catalog

Status: Pending

## Motivation

The client binary for the service catalog must cover every [use case](use-cases.md) and provide 
tools for every [persona](use-cases.md#personas) described. 

While in some use cases we will most likely rely on the base commands already present on `kubectl` - like 
`create`, `get`, `describe`, and `delete` - in others we want to have subcommands dedicated 
to interacting with the service catalog. It is expected that the "end-users" (the *Service Consumer* 
persona, for example), have a set of subcommands that are simple to use and allow them to easily check 
information about the services available on a catalog, request instances, bind them to apps, and 
manage their installed instances in a simple and user-friendly way.

This document proposes the way every use case is fulfilled by `kubectl` subcommands and/or extensions.

## Proposal

Below the [identified personas](use-cases.md#personas) and the commands that are expected to fulfill the use
cases for each one of them.

Note that when relying on existing commands, some improvements to printers and describers of the involved 
resources might be needed. For example, in the describer of a service instance we might want to display the 
number of bindings that currently exists and other information useful to catalog operators.

### Service consumer

For this persona we want simple and user-friendly commands. Service consumers must be able to easily search the entire
catalog for services available, check the metadata for each service, request an instance of a given service and
bind it to apps, only through the command line without the need of json or yaml files. 

So the proposal is to have a new set of subcommands called

```
kubectl catalog
```

which could be aliased as `kubectl (catalog | servicecatalog | service-catalog)` and distributed as part of the main 
`kubectl` binary or, preferably, as `kubectl` [binary plugins](#distribution) 
(TBD, this feature is under discussion and not yet available).

Note that `instances` and `bindings` are user-scoped, so they are applied in the context of current namespace (or what's
provided by the `--namespace` flag, etc).

The list of proposed subcommands under `kubectl catalog` is:

#### List and search the service catalog and existing service instances

```
Usage:

kubectl catalog list (available | instances) [--search=TERM]

Examples:

# List all services in the catalog available to be instantiated
$ kubectl catalog list available

## List all service instances in use by current namespace
$ kubectl catalog list instances

# List all service instances in use by the "mine" namespace
$ kubectl catalog list instances --namespace=mine

# Search the service catalog for services that match "foo"
$ kubectl catalog list available --search=foo

# Search the service catalog for services that belong to the "database" catalog
$ kubectl catalog list available --catalog=database
$ kubectl catalog list available --search=catalog=database

# Search the service instances in "mine" namespace provided by publisher "kube"
$ kubectl catalog list instances --publisher=kube --namespace=mine
$ kubectl catalog list instances --search=publisher=kube --namespace=mine
```

TBD if the top-level `list` will just print help for the 2 sub commands or run one of them by default,
so that I could for example list all services available in the catalog with just `kubectl catalog list`.

#### Describe information about a given service in the catalog, including metadata, plans, etc

```
Usage:

kubectl catalog describe SERVICE
```

List all metadata available about a given service in the catalog, well-formatted and suitable for 
on-screen reading. Includes description, publisher information, plans, URL's, etc. The expected
metadata information will likely include:

```
* name
* short description
* long desciption
* documentation/support urls
* icon URL
* image URLs - a list of images that could be displayed in a UI
* TOS (terms of service) link
* a list of plans
    * plan name
    * plan description
    * plan cost
* construction parameters
    * name
    * description
    * default value
* category label/tags
* version
* publisher name
* publisher contact url
* publisher website
```

TBD by plan name? What's the user-friendly "key" of a service?
TBD can also be used to describe my instances?

#### Request a service instance (and optionally binds it to an existing resource)

Binding at the time of request is optional. If not bound immediately, it can be done in a separate
command call.

```
Usage:

kubectl catalog provision SERVICE [--plan=PLAN] [--bind-to=TYPE/NAME] [--namespace=NAMESPACE]

Examples:

# Provision an instance of the "profiler" service in the "mine" namespace with the "2cores" plan
$ kubectl catalog provision profiler --plan=2cores --namespace=mine

# Provision an instance of "mariadb" with "5G" plan and bind it to my existing (Kube) service "frontend"
$ kubectl catalog provision mariadb --plan=5G --bind-to=svc/frontend --namespace=mine
```

TBD services has server-side validation when plans are strictly required, optional, not available, etc?

#### Bind an existing service instance

If not bound immediately (or if I want to bind it to other resource), it can be done with this command.

```
Usage:

kubectl catalog bind INSTANCE --to=TYPE/NAME [--namespace=NAMESPACE]

Examples:

# Bind the existing service instance "mysql" to the (Kube) service "frontend"
$ kubectl catalog bind mysql --to=svc/frontend --namespace=mine
```

Note that this command will likely print secrets, config maps and other resources created while binding,
so that I know how the bound service will be available to my resource.

TBD by instance name? What's the user-friendly "key" of a service instance?

#### Unbind, de-provision a service instance, etc

```
Usage:

kubectl catalog unbind INSTANCE --from=TYPE/NAME [--namespace=NAMESPACE]
```

```
Usage:

kubectl catalog deprovision INSTANCE [--namespace=NAMESPACE]
```

### Catalog operator

#### Register broker

With existing command(s), create a `broker` resource from json or yaml with 

```
kubectl create -f /path/to/broker.yaml
```

#### Update broker

With existing command(s): 

```
kubectl (edit | apply | patch | replace)
```

#### Get / list brokers

With existing command(s): 
```
kubectl (get | describe)
```

#### Delete broker

With existing command(s): 

```
kubectl delete
```

### Service producer

TBD, will probably use existing commands but might be worth creating a couple new subcommands under `kubectl create`.

## Distribution

It's arguable if the `kubectl catalog` set of commands should be part of the major distribution of the 
`kubectl` binary. The most fitting distribution mechanism would probably be binary extensions, or "kubectl
plugins" like we use to call it. However binary plugins is a part of a major discussion regarding "kubectl
extensions" which is still at the proposal phase, although it does have a PoC implementation that looks 
very similar to the solution adopted in Helm plugins. 

Relevant related links:

* [kubectl extensions proposal](https://github.com/kubernetes/community/pull/122)
* [previous kubectl extensions discussions](https://github.com/kubernetes/kubernetes/pull/30086)
* [PoC implementation](https://github.com/kubernetes/kubernetes/pull/37499) (fully functional with a [sample plugin](https://github.com/fabianofranz/kubernetes/blob/fd83635b26f1e41b84fcc1c1a811c3e41b63576c/pkg/kubectl/plugins/examples/kubectl-aging) written in Ruby)
* [Helm plugins architecture](https://github.com/kubernetes/helm/blob/master/docs/plugins.md)


