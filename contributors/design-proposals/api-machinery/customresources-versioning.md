CRD Versioning
=============

The objective of this design document is to provide a machinery for Custom Resource Definition authors to define different resource version and a conversion mechanism between them.

# **Background**

Custom Resource Definitions ([CRDs](https://kubernetes.io/docs/concepts/api-extension/custom-resources/)) are a popular mechanism for extending Kubernetes, due to their ease of use compared with the main alternative of building an Aggregated API Servers. They are, however, lacking a very important feature that all other kubernetes objects support: Versioning. Today, each CR can only have one version and there is no clear way for authors to advance their resources to a newer version other than creating a completely new CRD and converting everything manually in sync with their client software.

This document proposes a mechanism to support multiple CRD versions. A few alternatives are also explored at the end of the document. Notably the declarative approach which still has a lot of support among interested parties.

**Goals:**

* Support versioning on API level

* Support conversion mechanism between versions

* Support ability to change storage version

* Support Validation/OpenAPI schema for all versions

**Non-Goals:**

* Support cohabitation (i.e. no group/kind move)

# **Proposed Design**

The basis of the design is a system that supports versioning and minimal conversions. The APIs here, is designed in a way that can be extended with more conversions later.

The summary is to support a list of versions that will include current version. One of these versions can be flagged as the storage version and all versions ever marked as storage version will be listed in a stored_version field in the Status object to enable authors to plan a migration for their stored objects.

The current `Version` field in planned to be deprecated in a later release and will be used to pre-populate the `Versions` field (The `Versions` field will be defaulted to a single version, constructed from top level `Version` field). The `Version` field will be also be mutable to give a way to the authors to remove it from the list.

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
  // Storage flags the release as a storage version. There can be only one version
  // flagged as Storage.
  Storage Boolean
}
```

The Status object will have a list of potential stored objects. This data is necessary to do a storage migration in future (the author can choose to do the migration themselves but there is a plan [TODO(lavalamp): link?] to solve the problem of migration, potentially for both standard and custom types).

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

Basic validations needed for the `version` field are:

* `Spec.Version` field exists in `Spec.Versions` field (serve should be True for that version if there is any).
* If there is any `Served` `Version` in list of `Versions`, then the `Spec.Version`’s version in the `Spec.Versions` list should have a set `Served` flag. This is for backward compatibility. An old controller expect that version to be served but only the whole CRD is served. CRD Registration controller should unregister a CRD with no serving version.
* None of the `Status.StoredVersion` can be removed from `Spec.Versions` list.
* Only one of the versions in `spec.Versions` can flag as `Storage` version.

# **Declarative conversion**

Most of CRD conversions can be represented by simple operations like rename or move. The proposal is to support rename only and add more operations in future based on user feedback.

```golang
type CustomResourceDefinitionVersion {
  // Name is the version name, e.g. "v1", “v2beta1”, etc.
  Name string
  // Served is a flag enabling/disabling this version from being served via REST APIs
  Served Boolean
  // Storage flags the release as a storage version. There can be only one version
  // flagged as Storage.
  Storage Boolean
  //
  Conversion CustomResourceConversion
}

type CustomResourceConversion struct {
  Declarative map[string]CustomResourceDeclarativeConversion
}

type CustomResourceDeclarativeConversion {
  // This type will have collection of conversion operations. Only one of them should be set.
  Rename RenameConversion
}

type RenameConversion {
  From JsonPath
  To JsonPath
}

type JsonPath string
```

# **External Conversion/Webhooks**

Webhooks gives flexibility to CRD author to convert from the specific version to any version. As the conversion needs to be reversible, here we only list versions this webhook converts to/from. 

```golang
type CustomResourceDefinitionVersion {
  // Name is the version name, e.g. "v1", “v2beta1”, etc.
  Name string
  // Served is a flag enabling/disabling this version from being served via REST APIs
  Served Boolean
  // Storage flags the release as a storage version. There can be only one version
  // flagged as Storage.
  Storage Boolean
  //
  Conversion CustomResourceConversion
}

type CustomResourceConversion struct {
  Declarative map[string]CustomResourceDeclarativeConversion
  External CustomResourceConversionWebhook
}

type CustomResourceExternalConversion {
  // List of versions this external conversion supports.
  To []string
  Webhook ConversionWebhook
}

type ConversionWebhook struct {
  // ClientConfig defines how to communicate with the hook.
  ClientConfig WebhookClientConfig `json:"clientConfig" protobuf:"bytes,2,opt,name=clientConfig"`
}
```

This API design allows multiple webhooks but that should be prevented in validation. Also there is no need to redefine webhook for all versions. For example if on "v1" version we specified a webhook to convert it to “v2”, we do not need to add that webhook configuration to “v2” as the conversion is defined in both ways.

**Note: ** Webhooks will, intentionally, be not included in the first version of the feature to make sure authors waiting for this feature is trying the declarative approach first. The Webhook approach is a fallback for all cases that declarative cannot support.

# **Validation**

The API structure above allows some invalid states in order to keep a clean easy API. There need to be validation to catch undesire states. The validations would be

* Multiple webhooks are not allowed for a CRD. This will help us improve the performance by keeping the connection open to the webhook if the user defined a chain of version conversions.
* There should be no loop in the conversion graph. This includes v1->v2->v1 or v1->v2->v3->v1. The former is implicit in the assumption that webhook conversion is round-trippable. The second one allows two way to convert v1 to v3 which is not desired.
* Rename operation can be limited at first to not change the structure despite the fact that `To` field is also a json path.
