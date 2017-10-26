# Provide open-api extensions for kubectl get / kubectl describe columns

Status: Pending

Version: Alpha

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

**Note:** This solution will only work for types compiled into the apiserver
providing the open-api swagger.json to kubectl.  This solution will
not work for TPR, though TPR could possibly be solved in a similar
way by apply an annotation with the same key / value to the TPR.

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

This is done by hardcoding the extension for types compiled into
the api server.  As such this is only a solution for types
implemented using federated apiservers.

### Kubectl

Overview:

- In `kubectl get` use the `x-kubernetes-kubectl-get-columns` value
  when printing an object iff 1) it is defined and 2) the output type
  is "" (empty string) or "wide".

- In `kubectl describe` use the `x-kubernetes-kubectl-describe-columns` value
  when printing an object iff 1) it is defined


#### Option 1: Re-parse the open-api swagger.json in a kubectl library

Re-parse the open-api swagger.json schema and build a map of group version kind -> columns
parsed from the schema.  For this would look similar to validation/schema.go

In get.go and describe.go: After fetching the "Infos" from the
resource builder, lookup the group version kind from the populated map.

**Pros:**
  - Simple and straightforward solution
  - Scope of impacted Kubernetes components is minimal
  - Doable in 1.6

**Cons:**
  - Hacky solution
  - Can not be cleanly extended to support TPR

#### Option 2: Modify api-machinery RestMapper

Modify the api-machinery RestMapper to parse extensions prefixed
with `x-kubernetes` and include them in the *RestMapping* used by the resource builder.

```go
type RESTMapping struct {
	// Resource is a string representing the name of this resource as a REST client would see it
	Resource string

	GroupVersionKind schema.GroupVersionKind

	// Scope contains the information needed to deal with REST Resources that are in a resource hierarchy
	Scope RESTScope

	runtime.ObjectConvertor
	MetadataAccessor

    // Extensions
    ApiExtensions ApiExtensions
}

type ApiExtensions struct {
  Extensions map[string]interface{}
}
```

The tags would then be easily accessible from the kubectl get / describe
functions through:  `resource.Builder -> Infos -> Mapping -> DisplayOptions`

**Pros:**
  - Clean + generalized solution
  - The same strategy can be applied to support TPR
  - Can support exposing future extensions such as patchStrategy and mergeKey
  - Can be used by other clients / tools

**Cons:**
  - Fields are only loosely tied to rest
  - Complicated due to the broad scope and impact
  - May not be doable in 1.6

#### Considerations

What should be used for oth an open-api extension columns tag AND a
compiled in printer exist for a type?

- Apiserver only provides `describe` for types that are never compiled in
  - Compiled in `describe` is much more rich - aggregating data across many other types.
    e.g. Node describe aggregating Pod data
  - kubectl will not be able to provide any `describe` information for new types when version skewed against a newer server
- Always use the extensions if present
  - Allows server to control columns.  Adds new columns for types on old clients that maybe missing the columns.
- Always use the compiled in commands if present
  - The compiled in `describe` is richer and provides aggregated information about many types.
- Always use the `get` extension if present.  Always use the `describe` compiled in code if present.
  - Inconsistent behavior across how extensions are handled

### Client/Server Backwards/Forwards compatibility

#### Newer client

Client doesn't find the open-api extensions.  Fallback on 1.5 behavior.

In the future, this will provide stronger backwards / forwards compatibility
as it will allow clients to print objects

#### Newer server

Client doesn't respect open-api extensions.  Uses 1.5 behavior.

## Alternatives considered

### Fork Kubectl and compile in go types

Fork kubectl and compile in the go types.  Implement get / describe
for the new types in the forked version.

**Pros:** *This is what will happen for sig-service catalog if we take no action in 1.6*

**Cons:** Bad user experience.  No clear solution for patching forked kubectl.
User has to use a separate kubectl binary per-apiserver.  Bad president.

I really don't want this solution to be used.

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
