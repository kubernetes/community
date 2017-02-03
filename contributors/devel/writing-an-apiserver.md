Writing an API Server
=====================

# Why write an apiserver?

Reasons to write an apiserver:

1. You want a new API for kubernetes, but it is unlikely to be merged
   into core APIs.
1. the lifetime of your API is separrate from that of the core APIs.
1. Doing something experimental and need access to the internal APIs.
1. Want to federate your API to the kubernetes core APIs so it appears
   as a first class API from an external user of kubernetes.

# Terms/Glossary/Definitions & Other Resources

`Upwards` means towards a client such as kubectl. ~~`Downwards` means
towards the storage and persistence layer.~~

See the [API-Group Documentation](adding-an-APIGroup.md) for some details on writing an
API-Group. This document may overlap in it's coverage.

Also of interest is
[API Conventions](https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md)
&
[API Changes](https://github.com/kubernetes/community/blob/master/contributors/devel/api_changes.md)

# Directory Structure

This is based on the
[`service-catalog`](https://github.com/kubernetes-incubator/service-catalog/pulls)
apiserver with the parts that can change outlined with angle brackets
`<>`.

```
  cmd/
      <binary-name>
      service-catalog
  pkg/
      apis/
           <api-group-name>
           servicecatalog/
               install/
               <versions>
               v1alpha1/
               validation/
      apiserver/
      cmd/server/
      registry/
          <api-group-name>/
              rest/
                  <storage>
              <individual resources>/
                  storage.go
                  strategy.go
          servicecatalog/
              rest/
                  storage_servicecatalog.go
              broker/
                  storage.go
                  strategy.go
              binding/
                  storage.go
                  strategy.go
              ... other types
```

A generic apiserver structure looks like:

```
  cmd/<binary-name>/server.go
  pkg/
      apis/
           <api-group-name>/
               types.go
               install/
                   install.go
               <versions>/
                   types.go
                   conversion.go
                   defaults.go
               validation/
                   validation.go
      apiserver/apiserver.go
      cmd/server/server.go
      registry/
          <api-group-name>/
              rest/
                  <storage>
              <individual types>/
                  storage.go
                  strategy.go
```

These structures follow the structure of kubernetes.

## `cmd/`

This top-level directory contains the package that defines the output
binary that will be run. In service-catalog this exists to set up
logging, and call the flag handler to do the real setup and run of the
server. 

## `pkg/`
This directory contains the bulk of the code. 

### `apis/`

This contains all of the type definitions along with conversion
between versions, validation of types, and the
registration/installation of the API-Group.

#### `pkg/apis/<API-Group>/`

Each API-Group has it's own named directory, which should be named by
running all the words in it together with no spaces or dashes. 

Example: The `Service Catalog` API-Group becomes a `servicecatalog` directory.

This is often referred to as the **base** directory of the API-Group.

This base directory contains the files and directories:
 - `types.go` - Definition of all resources to be exposed by the API as
   well as internal sub-objects.
 - `register.go` - definition of API-Group, code to help installation of API.
 - `docs.go` - code generation annotations.
 - `serialization_test.go` - fuzz testing of conversion of resources
   between versions.
 - `install/` - installation package to import.
 - `validation/` - contains code to validate the correctness of an
   object for persistence.
 - directory for each version - for example `v1alpha1/`, `v1beta1/`, `v1/`
 - generated code
 
The go package that this base represents must have a `SchemeBuilder`
variable for code generation to find. The variable is typically
located in the file `register.go`.

##### `pkg/apis/<API-Group>/types.go`

##### `pkg/apis/<API-Group>/install/`

This is typically a sole file named `install.go`.

This package is imported by the server to cause it's `init()` to
run. `init()` must be defined for the package. Import the base and all
versions. Use the `AddToScheme` vars defined in each package. Use a
GroupMetaFactory to register and enable the API-Group.



##### `pkg/apis/<API-Group>/validation/`

##### `pkg/apis/<API-Group>/<version>/`

### `apiserver/apiserver.go`

### `registry/`

# Type Definitions

In the root of the api-group `types.go` defines the available
resources for the API. The structs defined are the *unversioned*
resources, and should not be serialized upwards to a client. These
unversioned resources must be able to hold any version, as all
versioned resources will be converted into their unversioned form
before being persisted.

The *versioned* `types.go` resides in a directory with the version
name that follows
[kubernetes conventions](api_changes.md#alpha-beta-and-stable-versions). The
defined structs in this file should have `json` annotations.

# Overall Flow
