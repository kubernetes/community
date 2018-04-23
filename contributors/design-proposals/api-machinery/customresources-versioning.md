CRD Versioning
=============

The objective of this design document is to provide a machinery for Custom Resource Definition authors to define different resource version and a conversion mechanism between them.

# **Background**

Custom Resource Definitions ([CRDs](https://kubernetes.io/docs/concepts/api-extension/custom-resources/)) are a popular mechanism for extending Kubernetes, due to their ease of use compared with the main alternative of building an Aggregated API Server. They are, however, lacking a very important feature that all other kubernetes objects support: Versioning. Today, each CR can only have one version and there is no clear way for authors to advance their resources to a newer version other than creating a completely new CRD and converting everything manually in sync with their client software.

This document proposes a mechanism to support multiple CRD versions. A few alternatives are also explored in [this document](https://docs.google.com/document/d/1Ucf7JwyHpy7QlgHIN2Rst_q6yT0eeN9euzUV6kte6aY).

**Goals:**

* Support versioning on API level

* Support conversion mechanism between versions

* Support ability to change storage version

* Support Validation/OpenAPI schema for all versions: All versions should have a schema. This schema can be provided by user or derived from a single schema.

**Non-Goals:**

* Support cohabitation (i.e. no group/kind move)

# **Proposed Design**

The basis of the design is a system that supports versioning and no conversion. The APIs here, is designed in a way that can be extended with conversions later.

The summary is to support a list of versions that will include current version. One of these versions can be flagged as the storage version and all versions ever marked as storage version will be listed in a stored_version field in the Status object to enable authors to plan a migration for their stored objects.

The current `Version` field is planned to be deprecated in a later release and will be used to pre-populate the `Versions` field (The `Versions` field will be defaulted to a single version, constructed from top level `Version` field). The `Version` field will be also mutable to give a way to the authors to remove it from the list.

```golang
// CustomResourceDefinitionSpec describes how a user wants their resource to appear
type CustomResourceDefinitionSpec struct {
  // Group is the group this resource belongs in
  Group string
  // Version is the version this resource belongs in
  // must be always the first item in Versions field if provided.
  Version string
  // Names are the names used to describe this custom resource
  Names CustomResourceDefinitionNames
  // Scope indicates whether this resource is cluster or namespace scoped.  Default is namespaced
  Scope ResourceScope
  // Validation describes the validation methods for CustomResources
  Validation *CustomResourceValidation

  // ***************
  // ** New Field **
  // ***************
  // Versions is the list of all supported versions for this resource.
  // Validation: All versions must use the same validation schema for now. i.e., top 
  // level Validation field is applied to all of these versions.
  // Order: The order of these versions is used to determine the order in discovery API
  // (preferred version first).
  // The versions in this list may not be removed if they are in 
  // CustomResourceDefinitionStatus.StoredVersions list.
  Versions []CustomResourceDefinitionVersion
}

// ***************
// ** New Type **
// ***************
type CustomResourceDefinitionVersion {
  // Name is the version name, e.g. "v1", “v2beta1”, etc.
  Name string
  // Served is a flag enabling/disabling this version from being served via REST APIs
  Served Boolean
  // Storage flags the release as a storage version. There must be exactly one version
  // flagged as Storage.
  Storage Boolean
}
```

The Status object will have a list of potential stored versions. This data is necessary to do a storage migration in future (the author can choose to do the migration themselves but there is [a plan](https://docs.google.com/document/d/1eoS1K40HLMl4zUyw5pnC05dEF3mzFLp5TPEEt4PFvsM/edit) to solve the problem of migration, potentially for both standard and custom types).

```golang
// CustomResourceDefinitionStatus indicates the state of the CustomResourceDefinition
type CustomResourceDefinitionStatus struct {
  ...

  // StoredVersions are all versions ever marked as storage in spec. Tracking these
  // versions allow a migration path for stored version in etcd. The field is mutable
  // so the migration controller can first make sure a version is certified (i.e. all
  // stored objects is that version) then remove the rest of the versions from this list.
  // None of the versions in this list can be removed from the spec.Versions field.
  StoredVersions []string
}
```

# **Validation**

Basic validations needed for the `version` field are:

* `Spec.Version` field exists in `Spec.Versions` field.
* The version defined in `Spec.Version` field should point to a `Served` Version in `Spec.Versions` list except when we do not serve any version (i.e. all versions in `Spec.Versions` field are disabled by `Served` set to `False`). This is for backward compatibility. An old controller expect that version to be served but only the whole CRD is served. CRD Registration controller should unregister a CRD with no serving version.
* None of the `Status.StoredVersion` can be removed from `Spec.Versions` list.
* Only one of the versions in `spec.Versions` can flag as `Storage` version.

