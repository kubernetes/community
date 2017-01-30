Writing an API Server
=====================

# Why write an apiserver?

Reasons to write an apiserver:

1. You want a new api for kubernetes, but it is unlikely to be merged
   into core APIs.
1. the lifetime of your API is separrate from that of the core APIs.
1. Doing something experimental and need access to the internal APIs.
1. Want to federate your API to the kubernetes core APIs so it appears
   as a first class API from an external user of kubernetes.



# Glossary

`Upwards` means towards a client such as kubectl. `Downwards` means
towards the storage and persistence layer.

See the [API-Group Documentation](adding-an-APIGroup.md) for some details on writing an
API-Group. This document may overlap in it's coverage.

# Directory Structure

This is based on the `service-catalog` apiserver with the parts that
can change outlined with angle brackets `<>`.

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
              <individual types>/
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

## cmd/

This toplevel directory contains the package that defines the output
binary that will be run. In service-catalog this exists to set up
logging, and call the flag handler to do the real setup and run of the
server. 

## pkg/
This directory contains the bulk of the code. 


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


