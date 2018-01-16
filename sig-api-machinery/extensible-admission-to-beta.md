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
former is unnecessary and the latter doesn't work.

API validation on uninitialized objects can reasonably follow the rules for validation on create
plus `ValidateObjectMetaUpdate` (to prevent name and UID changes and the like).
Since the object hasn't been "ready", there is no reason to add the validation rules required on
updates since restrictions like field immutability don't logically apply and would restrict the power
of admission extensions to the point of unusability.  Using the results for creation validation on 
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

