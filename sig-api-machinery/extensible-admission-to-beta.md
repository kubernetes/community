# Admission control changes

Admission control is the key security and cluster policy enforcement mechanism
in Kubernetes. In order to make a reliable extension mechanism for admission,
we must reconsider the current chain to produce understandable extensions.

We are proposing a series of changes to admission control in an attempt to make
it less error-prone, easier to configure, and more flexible for admission plugin
developers.

## Requirements for extensible admission
Ordering the admission chain is critical.  One mutating admission plugin can invalidate
the results of other admission plugins.  For instance, the `LimitVerifier` would
invalidate results for the `ResourceQuota` plugin.  More generally, object mutations
must happen before object validations.  In addition, compiled validating admission plugins
must run after mutating admission webhooks.  Splitting the chain into mutators and validators
will resolve the coarse grained ordering problem and would require two different webhook APIs
since the `Status` would be different for both now and we can reasonably expect different
evolutions.  It would also force changes in the registration APIs since ordering matters
for individual mutators, but ordering does not matter for validators.  Having fewer items to 
order makes the registration API easier to reason about for maintainability.

Initializers as implemented today attempt to push new API validation requirements into the 
strategy and new requirements onto the admission plugins handling updates and creates.  The
former is unecessary and the latter doesn't work.

API validation on uninitialized objects can reasonably follow the rules for validation on create.
Since the object hasn't been "ready", there is no reason to add the validation rules required on
updates since restrictions like field immutability don't logically apply and would restrict the power
of admission extensions the point of unusability.  Using the results for creation validation on 
uninitialized objects and objects transitioning from uninitialized to initialized lets the
strategy author focus on their API needs.

Admission validation is also needed on resources being initialized.  During initialization, there
may be some admission validation required, but it isn't normal API update admission.  The update
admission mutators shouldn't run and since the subject (user.Info) on the admission attributes is
NOT the original user, normal update admission mutators and validators are likely to make improper
choices.  When the object transitions from uninitialized to initialized state, the original object
creator should be artificially injected into the admission attributes and the create admission 
validators should run.  This allows mechanisms like PSP to function properly during their validation
phase.  Such injection requires persisting the original user.Info (probably encrypted) into the 
object.

## How much is needed for beta
Basically all of it.  After considering the needs of initializers, the admission chain needs to
be split into mutators and validators.  That has a direct impact on the webhook admission registration
and callout APIs.  The initializer API handling needs to be updated to work with our existing chain.
Mutating admission webhooks are needed to support parity in the chain, so we know we'll eventually
need them and because they will directly interact in the registration API, we need to know how we'll
build them before we can push a registration API to beta.


### Admission interfaces

The current interface in `pkg/admission/interfaces.go` looks like this:

```go
type Interface interface {
  Admit(a Attributes) (err error)
  Handles(operation Operation) bool
}
```

We'd like to break this into multiple interfaces, one per phase, called at the appropriate times
within the request lifecycle.

```go
type Mutator interface {
  // ApplyMutations is executed after the request body is deserialized into an
  // object for create and update requests.
  // Plugins are free to modify the object as they see fit. If an error is
  // returned, the request is marked as failed and no changes are persisted.
  ApplyMutations(a Attributes) error
}

type Validator interface {
  // Validate is executed after the appropriate `rest.Storage` handler's
  // validation occurs. This gives the plugin the opportunity to validate the
  // object after the handler has validated it, but before it is persisted. If
  // an error is returned, the request is marked as failed and no changes are
  // persisted.
  Validate(a Attributes) error
}
```

### rest Storage

To be able to invoke admission plugin validation at the right time, we'll need
to modify the various `rest.Storage` interfaces, as described below.

#### Creater, NamedCreater

- add `BeforeCreate(ctx api.Context, obj runtime.Object) error`
- add `Validate(ctx api.Context, obj runtime.Object) error`

#### Updater, CreaterUpdater

- add `BeforeUpdate(ctx api.Context, obj, old runtime.Object) error`
- add `ValidateUpdate(ctx api.Context, obj, old runtime.Object) field.ErrorList`

#### Patcher

- inherits changes from `Updater`


#### Deleter, GracefulDeleter

- add `BeforeDelete(ct api.Context, obj runtime.Object, options *api.DeleteOptions) (graceful, gracefulPending bool, err error)`

#### Connector

- QUESTION: do we need to adjust this interface?


## REST flows

### Create

#### Current behavior

This is the current [apiserver create
flow](https://github.com/kubernetes/kubernetes/blob/ef0c9f0c5b8efbba948a0be2c98d9d2e32e0b68c/pkg/apiserver/resthandler.go#L333):

1. decode incoming request
1. admit.Admit
1. r.Create (the items below are part of the generic registry store / strategy-based storage, `pkg/registry/generic/registry/store.go`)
  1. rest.BeforeCreate
    1. strategy.PrepareForCreate
    1. api.FillObjectMetaSystemFields
    1. api.GenerateName
    1. strategy.Validate
    1. validation.ValidateObjectMeta
    1. strategy.Canonicalize
  1. e.Storage.Create
  1. e.AfterCreate
  1. e.Decorator

#### Proposed behavior

This is the proposed new behavior:

1. decode request
1. admit.ApplyMutations
1. r.BeforeCreate
  1. If strategy-based, call rest.BeforeCreate
    1. strategy.PrepareForCreate
    1. api.FillInObjectMetaSystemFields
    1. api.GenerateName
  1. Otherwise, just do whatever is in the rest.Creater
1. r.Validate
  1. If strategy-based, call rest.Validate
    1. strategy.Validate
    1. validation.ValidateObjectMeta
    1. strategy.Canonicalize
  1. Otherwise, just do whatever is in the rest.Creater
1. admit.Validate
1. r.Create
  1. For strategy-based handlers:
    1. e.Storage.Create
    1. e.AfterCreate
    1. e.Decorator
  1. Otherwise, just do whatever is in the rest.Creater

Notes:

- admit.ApplyMutations replaces admit.Admit and gives plugins a chance to apply
  mutations to the decoded request object
- r.BeforeCreate is the `rest.Creater` interface's new method for performing
  logic after admission defaulting and before validation
- r.Validate is the `rest.Creater` interface's new method for performing
  validation
- admit.Validate is a new admission phase that is executed after the
  `rest.Creater`'s validation (plugins should not modify storage)
- admit.Validate (plugins can modify storage)


### Update

#### Current behavior

1. decode request
1. admit.Admit
1. r.Update (the items below are part of the generic registry store / strategy-based storage, `pkg/registry/generic/registry/store.go`)
  1. rest.BeforeUpdate
    1. strategy.PrepareForUpdate
    1. validation.ValidateObjectMetaUpdate
    1. strategy.ValidateUpdate
    1. strategy.Canonicalize
  1. e.Storage.GuaranteedUpdate
  1. e.AfterUpdate
  1. e.Decorator

#### Proposed behavior

1. decode request
1. admit.ApplyMutations
1. r.BeforeUpdate
  1. If strategy-based, call rest.BeforeUpdate
    1. strategy.PrepareForUpdate
  1. Otherwise, just do whatever is in the rest.Updater
1. r.ValidateUpdate
  1. If strategy-based, call rest.ValidateUpdate
    1. validation.ValidateObjectMetaUpdate
    1. strategy.ValidateUpdate
    1. strategy.Canonicalize
  1. Otherwise, just do whatever is in the rest.Updater
1. admit.Validate
1. r.Update
  1. For strategy-based handlers:
    1. e.Storage.GuaranteedUpdate
    1. e.AfterUpdate
    1. e.Decorator
  1. Otherwise, just do whatever is in the rest.Updater

Notes:

- The `rest.Updater` interface's Update method has been split into 3:
  - BeforeUpdate
  - ValidateUpdate
  - Update


### Patch

Patch is similar to Update, and similar changes are required to split the call
to `Update` into `BeforeUpdate`, `ValidateUpdate`, and `Update`. Additionally,
calls to `admit.Validate` will
occur in appropriate places within the request lifecycle.

### Delete

#### Current behavior

1. decode request (DeleteOptions)
2. admit.Admit
3. r.Delete (the items below are part of the generic registry store / strategy-based storage, `pkg/registry/generic/registry/store.go`)
  4. e.Storage.Get
  1. rest.BeforeDelete
  1. e.Storage.Delete
  1. e.AfterDelete
  1. e.Decorator

#### Proposed behavior

1. decode request (DeleteOptions)
2. r.Get
1. r.BeforeDelete
  2. If strategy-based, call rest.BeforeDelete
  1. Otherwise, just do whatever is in the rest.Deleter
1. admit.Validate
1. r.Delete
  2. For strategy-based handlers:
    1. e.Storage.Delete
    1. e.AfterDelete
    1. e.Decorator

### DeleteCollection

#### Current behavior

1. admit.Admit
2. decode ListOptions from request parameters
3. decode request body (DeleteOptions)
4. r.DeleteCollection (the items below are part of the generic registry store / strategy-based storage, `pkg/registry/generic/registry/store.go`)
  5. e.List
  1. run *n* worker goroutines in parallel:
    1. e.Delete (same as Delete -> Current behavior -> 3)

#### Proposed behavior

1. admit.Validate
2. decode ListOptions from request parameters
3. decode request body (DeleteOptions)
4. r.DeleteCollection ...

### Connect

#### Current behavior

1. connecter.NewConnectOptions
2. decode request connect options
3. admit.Admit
4. connecter.Connect().ServeHTTP()

#### Proposed behavior

1. connecter.NewConnectOptions
2. decode request connect options
1. admit.Validate
2. connecter.Connect().ServeHTTP()

## Ability to run out-of-tree plugins

Admission plugins are all currently compiled into the apiserver binary. If you
want to run custom plugins, you have to modify the Kubernetes source code and
recompile the apiserver. Some examples of custom plugins that an administrator
might want to run (these are all actual plugins in OpenShift today):

- default restart-never and restart-on-failure pods' `activeDeadlineSeconds`
  based on either a namespace-level annotation, or the plugin's configured value
  (although this probably should be part of LimitRange).
- apply quota at the cluster level instead of per-namespace
- assign default resource request values relative to user-specified limits for
  pods to support an administrator-defined overcommit target
- restrict who is allowed to set the `nodeName` or `nodeSelector` on pods
- ensure you can't exec or attach to a pod with a privileged security context if
  you don't have privileges to create a privileged pod
- supply default values for certain fields on Build Pods when unset (e.g.
  environment variables)
- enforce cluster-wide configuration values on all Build Pods (i.e.
  override user-supplied data if necessary)
- automatically provision a Jenkins server in a user's namespace if it doesn't
  exist when a pipeline BuildConfig is created

We would like to have the ability to run a single external plugin at the end of
each admission phase. This will enable cluster administrators to inject custom
admission logic into an existing kube-apiserver without having to recompile it.

### OPEN QUESTIONS

#### What happens if the external admission plugin is not available?

Let's imagine that I have a Kubernetes cluster up and running. Now I want to run
OpenShift on top of it, so I deploy its various pods to the cluster, presumably
as Deployment resources. Next, I need to add OpenShift as an external admission
plugin for the kube-apiserver pods and restart them.

Fast forward a bit - now I have everything working. Then let's imagine that my
openshift-apiserver pods fail for whatever reason. What happens to
kube-apiserver requests with respect to admission? With openshift-apiserver
down, any attempt by the kube-apiserver to invoke the external admission plugin
will fail.

Do we ignore an external admission plugin being unreachable and allow requests
to proceed ("fail open")? I don't think this is viable. Presumably the external
admission plugin is there for a reason, and if it is unable to execute, it opens
the possibility for inconsistent or incorrect data to enter the cluster.

Do we instead reject the request because the external admission plugin is
unreachable ("fail closed")? I think this is what we must do, for the reason
stated above. Assuming this is what we choose, how do we restore the cluster to
a healthy state, as the default "fail closed" behavior won't allow any new pods,
and we need at least one new openshift-apiserver pod to handle the external
admission plugin invocations. We need some way to a) indicate that a certain set
of changes to the system are being made to restore the external admission
plugin, and b) ensure that users can't "sneak in" anything while admission isn't
fully operational. Some possible options include:

- Specify via flag/configuration the identity of a user or group whose requests
  will be allowed to proceed even when external admission is down
  - If the openshift-apiserver pods come from a Deployment, then the
    DeploymentController needs permission to create openshift-apiserver pods,
    but not pods for normal users
- Require that kube-apiserver and openshift-apiserver be deployed as containers
  in the same pod. This won't completely eliminate the possibility of all
  openshift-apiservers being down at the same time, but it will significantly
  minimize the likelihood of that happening. We probably still need a mitigation
  plan in place for the event where you're down to 1 running apiserver pod,
  kube-apiserver is running, openshift-apiserver has failed, and the system is
  trying simultaneously to restart the openshift-apiserver container and scale
  the replicas to 3+.

