# Provide open-api extensions for kubectl get / kubectl describe columns

Status: Pending

## Motivation

`kubectl get` and `kubectl describe` do not provide a rich experience
for resources retrieved through federated apiservers and types not
compiled into the kubectl binary.  Kubectl should support printing
columns configured per-type without having the types compiled in.

## Proposal

Allow the apiserver to define the type specific columns that will be
printed using the open-api swagger.json spec already fetched by kubectl.
This provides a limited describe to only print out fields on the object
and related events.

## User Experience

### Use Cases

- As a user, when I run `kubectl get` on sig-service-catalog resources
  defined in a federated apiserver, I want to see more than just the
  name and the type of the resource.
- As a user, when I run `kubectl describe` on sig-service-catalog
  resources defined in a federated apiserver, I want the command
  to succeed, and to see events for the resource along with important
  fields of the resource.

## Implementation

Define the open-api extensions `x-kubernetes-kubectl-get-columns` and
`x-kubernetes-kubectl-describe-columns`.  These extensions have a
string value containing the columns to be printed by kubectl.  The
string format is the same as the `--custom-columns` for `kubectl get`.

### Apiserver

- Populate the open-api extension value for resource types.

### Kubectl

- In `kubectl get` use the `x-kubernetes-kubectl-get-columns` value
  when printing an object iff 1) it is defined and 2) the output type
  is "" (empty string) or "wide".

- In `kubectl describe` use the `x-kubernetes-kubectl-describe-columns` value
  when printing an object iff 1) it is defined

### Client/Server Backwards/Forwards compatibility

#### Newer client

Client doesn't find the open-api extensions.  Fallback on 1.5 behavior.

#### Newer server

Client doesn't respect open-api extensions.  Uses 1.5 behavior.

## Alternatives considered

### Kubectl describe fully implemented in the server

Implement a sub-resource "/describe" in the apiserver.  This executes
the describe business logic for the object and returns either a string
or json blob for kubectl to print.

**Pros:** Higher fidelity.  Can aggregate data and fetch other objects.

**Cons:** Higher complexity.  Requires more api changes.

### Write per-type columns to kubectl.config or another local file

Support checking a local file containing per-type information including
the columns to print.

**Pros:** Simplest solution.  Easy for user to override values.

**Cons:** Requires manual configuration on user side.  Does not provide a consistent experience across clients.

### Write per-type go templates to kubectl.config or another local file

Support checking a local file containing per-type information including
the go template.

**Pros:** Higher fidelity.  Easy for user to override values.

**Cons:** Higher complexity. Requires manual configuration on user side.  Does not provide a consistent experience across clients.
