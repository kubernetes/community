# CRD Conversion Webhook

Status: Approved

Version: Alpha

Implementation Owner: @mbohlool

Authors: @mbohlool, @erictune

Thanks: @dbsmith, @deads2k, @sttts, @liggit, @enisoc

### Summary

This document proposes a detailed plan for adding support for version-conversion of Kubernetes resources defined via Custom Resource Definitions (CRD).  The API Server is extended to call out to a webhook at appropriate parts of the handler stack for CRDs.  

No new resources are added; the [CRD resource](https://github.com/kubernetes/kubernetes/blob/34383aa0a49ab916d74ea897cebc79ce0acfc9dd/staging/src/k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/types.go#L187) is extended to include conversion information as well as multiple schema definitions, one for each apiVersion that is to be served.


## Definitions

**Webhook Resource**: a Kubernetes resource (or portion of a resource) that informs the API Server that it should call out to a Webhook Host for certain operations. 

**Webhook Host**: a process / binary which accepts HTTP connections, intended to be called by the Kubernetes API Server as part of a Webhook.

**Webhook**: In Kubernetes, refers to the idea of having the API server make an HTTP request to another service at a point in its request processing stack.  Examples are [Authentication webhooks](https://kubernetes.io/docs/reference/access-authn-authz/webhook/) and [Admission Webhooks](https://kubernetes.io/docs/reference/access-authn-authz/extensible-admission-controllers/).  Usually refers to the system of Webhook Host and Webhook Resource together, but occasionally used to mean just Host or just Resource.

**Conversion Webhook**: Webhook that can convert an object from one version to another.

**Custom Resource**: In the context of this document, it refers to resources defined as Custom Resource Definition (in contrast with extension API server’s resources).

**CRD Package**: CRD definition, plus associated controller deployment, RBAC roles, etc, which is released by a developer who uses CRDs to create new APIs.

  
## Motivation

Version conversion is, in our experience, the most requested improvement to CRDs.  Prospective CRD users want to be certain they can evolve their API before they start down the path of developing a CRD + controller. 


## Requirements

* As an existing author of a CRD, I can update my API's schema, without breaking existing clients.  To that end, I can write a CRD(s) that supports one kind with two (or more) versions.  Users of this API can access an object via either version (v1 or v2), and are accessing the same underlying storage (assuming that I have properly defined how to convert between v1 and v2.)

* As a prospective user of CRDs, I don't know what schema changes I may need in the future, but I want to know that they will be possible before I chose CRDs (over EAS, or over a non-Kubernetes API).

* As an author of a CRD Package, my users can upgrade to a new version of my package, and can downgrade to a prior version of my package (assuming that they follow proper upgrade and downgrade procedures; these should not require direct etcd access.)

* As a user, I should be able to request CR in any supported version defined by CRD and get an object has been properly converted to the requested version (assuming the CRD Package Author has properly defined how to convert).

* As an author of a CRD that does not use validation, I can still have different versions which undergo conversion.

* As a user, when I request an object, and webhook-conversion fails, I get an error message that helps me understand the problem.

* As an API machinery code maintainer, this change should not make the API machinery code harder to maintain

* As a cluster owner, when I upgrade to the version of Kubernetes that supports CRD multiple versions, but I don't use the new feature, my existing CRDs work fine.  I can roll back to the previous version without any special action.


## Summary of Changes

1. A CRD object now represents a group/kind with one or more versions.

2. The CRD API (CustomResourceDefinitionSpec) is extended as follows:

    1. It has a place to register 1 webhook.

    2. it holds multiple "versions".

    3. Some fields which were part of the .spec are now per-version; namely Schema, Subresources, and AdditionalPrinterColumns.

3. A Webhook Host is used to do conversion for a CRD.

    4. CRD authors will need to write a Webhook Host that accepts any version and returns any version.

    5. Toolkits like kube-builder and operator-sdk are expected to provide flows to assist users to generate Webhook Hosts.


## Detailed Design


### CRD API Changes

The CustomResourceDefinitionSpec is extended to have a new section where webhooks are defined: 

```golang
// CustomResourceDefinitionSpec describes how a user wants their resource to appear
type CustomResourceDefinitionSpec struct {
  Group string
  Version string
  Names CustomResourceDefinitionNames
  Scope ResourceScope
  // Optional, can only be provided if per-version schema is not provided.
  Validation *CustomResourceValidation
  // Optional, can only be provided if per-version subresource is not provided.
  Subresources *CustomResourceSubresources
  Versions []CustomResourceDefinitionVersion
  // Optional, can only be provided if per-version additionalPrinterColumns is not provided.  
  AdditionalPrinterColumns []CustomResourceColumnDefinition

  Conversion *CustomResourceConversion
}

type CustomResourceDefinitionVersion struct {
  Name string
  Served Boolean
  Storage Boolean
  // Optional, can only be provided if top level validation is not provided.
  Schema *JSONSchemaProp
  // Optional, can only be provided if top level subresource is not provided.
  Subresources *CustomResourceSubresources
  // Optional, can only be provided if top level additionalPrinterColumns is not provided.  
  AdditionalPrinterColumns []CustomResourceColumnDefinition
}

Type CustomResourceConversion struct {
  // Conversion strategy, either "nop” or "webhook”. If webhook is set, Webhook field is required.
  Strategy string

  // Additional information for external conversion if strategy is set to external
  // +optional
  Webhook *CustomResourceConversionWebhook
}

type CustomResourceConversionWebhook {
  // ClientConfig defines how to communicate with the webhook. This is the same config used for validating/mutating webhooks.
  ClientConfig WebhookClientConfig
}
```

### Top level fields to Per-Version fields

In *CRD v1beta1* (apiextensions.k8s.io/v1beta1) there are per-version schema, additionalPrinterColumns or subresources (called X in this section) defined and these validation rules will be applied to them:

* Either top level X or per-version X can be set, but not both. This rule applies to individual X’s not the whole set. E.g. top level schema can be set while per-version subresources are set.
* per-version X cannot be the same. E.g. if all per-version schema are the same, the CRD object will be rejected with an error message asking the user to use the top level schema.

in *CRD v1* (apiextensions.k8s.io/v1), there will be only version list with no top level X. The second validation guarantees a clean moving to v1. These are conversion rules:

*v1beta1->v1:*

* If top level X is set in v1beta1, then it will be copied to all versions in v1.
* If per-version X are set in v1beta1, then they will be used for per-version X in v1.

*v1->v1beta1:*

* If all per-version X are the same in v1, they will be copied to top level X in v1beta1
* Otherwise, they will be used as per-version X in v1beta1

#### Alternative approaches considered

First a defaulting approach is considered which per-version fields would be defaulted to top level fields. but that breaks backward incompatible change; Quoting from API [guidelines](/contributors/devel/sig-architecture/api_changes.md#backward-compatibility-gotchas):

> A single feature/property cannot be represented using multiple spec fields in the same API version simultaneously

Hence the defaulting either implicit or explicit has the potential to break backward compatibility as we have two sets of fields representing the same feature.

There are other solution considered that does not involved defaulting:

* Field Discriminator: Use `Spec.Conversion.Strategy` as discriminator to decide which set of fields to use. This approach would work but the proposed solution is keeping the mutual excusivity in a broader sense and is preferred.
* Per-version override: If a per-version X is specified, use it otherwise use the top level X if provided. While with careful validation and feature gating, this solution is also backward compatible, the overriding behaviour need to be kept in CRD v1 and that looks too complicated and not clean to keep for a v1 API.

Refer to [this document](http://bit.ly/k8s-crd-per-version-defaulting) for more details and discussions on those solutions.

### Support Level

The feature will be alpha in the first implementation and will have a feature gate that is defaulted to false. The roll-back story with a feature gate is much more clear. if we have the features as alpha in kubernetes release Y (>X where the feature is missing) and we make it beta in kubernetes release Z, it is not safe to use the feature and downgrade from Y to X but the feature is alpha in Y which is fine. It is safe to downgrade from Z to Y (given that we enable the feature gate in Y) and that is desirable as the feature is beta in Z.
On downgrading from a Z to Y, stored CRDs can have per-version fields set. While the feature gate can be off on Y (alpha cluster), it is dangerous to disable per-version Schema Validation or Status subresources as it makes the status field mutable and validation on CRs will be disabled. Thus the feature gate in Y only protects adding per-version fields not the actual behaviour. Thus if the feature gate is off in Y:

* Per-version X cannot be set on CRD create (per-version fields are auto-cleared).
* Per-version X can only be set/changed on CRD update *if* the existing CRD object already has per-version X set.

This way even if we downgrade from Z to Y, per-version validations and subresources will be honored. This will not be the case for webhook conversion itself. The feature gate will also protect the implementation of webhook conversion and alpha cluster with disabled feature gate will return error for CRDs with webhook conversion (that are created with a future version of the cluster).

### Rollback

Users that need to rollback to version X (but may currently be running version Y > X) of apiserver should not use CRD Webhook Conversion if X is not a version that supports these features.  If a user were to create a CRD that uses CRD Webhook Conversion and then rolls back to version X that does not support conversion then the following would happen:

1. The stored custom resources in etcd will not be deleted.

2. Any clients that try to get the custom resources will get a 500 (internal server error). this is distinguishable from a deleted object for get and the list operation will also fail. That means the CRD is not served at all and Clients that try to garbage collect related resources to missing CRs should be aware of this. 

3. Any client (e.g. controller) that tries to list the resource (in preparation for watching it) will get a 500 (this is distinguishable from an empty list or a 404).

4. If the user rolls forward again, then custom resources will be served again.

If a user does not use the webhook feature but uses the versioned schema, additionalPrinterColumns, and/or subresources and rollback to a version that does not support them per-version, any value set per-version will be ignored and only values in top level spec.* will be honor.

Please note that any of the fields added in this design that is not supported in previous kubernetes releases can be removed on an update operation (e.g. status update). The kubernetes release where defined the types but gate them with an alpha feature gate, however, can keep these fields but ignore there value.

### Webhook Request/Response

The Conversion request and response would be similar to [Admission webhooks](https://github.com/kubernetes/kubernetes/blob/951962512b9cfe15b25e9c715a5f33f088854f97/staging/src/k8s.io/api/admission/v1beta1/types.go#L29). The AdmissionReview seems to be redundant but used by other Webhook APIs and added here for consistency.

```golang
// ConversionReview describes a conversion request/response.
type ConversionReview struct {
  metav1.TypeMeta
  // Request describes the attributes for the conversion request.
  // +optional
  Request *ConversionRequest
  // Response describes the attributes for the conversion response.
  // +optional
  Response *ConversionResponse
}

type ConversionRequest struct {
  // UID is an identifier for the individual request/response. Useful for logging.
  UID types.UID
  // The version to convert given object to. E.g. "stable.example.com/v1"
  APIVersion string
  // Object is the CRD object to be converted.
  Object runtime.RawExtension
}

type ConversionResponse struct {
  // UID is an identifier for the individual request/response.
  // This should be copied over from the corresponding ConversionRequest.
  UID types.UID
  // ConvertedObject is the converted version of request.Object.
  ConvertedObject runtime.RawExtension
}
```

If the conversion is failed, the webhook should fail the HTTP request with a proper error code and message that will be used to create a status error for the original API caller.


### Monitorability

There should be prometheus variables to show:

* CRD conversion latency
    * Overall
    * By webhook name
    * By request (sum of all conversions in a request)
    * By CRD
* Conversion Failures count
    * Overall
    * By webhook name
    * By CRD
* Timeout failures count
    * Overall
    * By webhook name
    * By CRD

Adding a webhook dynamically adds a key to a map-valued prometheus metric. Webhook host process authors should consider how to make their webhook host monitorable: while eventually we hope to offer a set of best practices around this, for the initial release we won’t have requirements here.


### Error Messages

When a conversion webhook fails, e.g. for the GET operation, then the error message from the apiserver to its client should reflect that conversion failed and include additional information to help debug the problem. The error message and HTTP error code returned by the webhook should be included in the error message API server returns to the user.  For example:

```bash
$ kubectl get mykind somename
error on server: conversion from stored version v1 to requested version v2 for somename: "408 request timeout" while calling service "mywebhookhost.somens.cluster.local:443"
```


For operations that need more than one conversion (e.g. LIST), no partial result will be returned. Instead the whole operation will fail the same way with detailed error messages. To help debugging these kind of operations, the UID of the first failing conversion will also be included in the error message. 


### Caching

No new caching is planned as part of this work, but the API Server may in the future cache webhook POST responses.

Most API operations are reads.  The most common kind of read is a watch.  All watched objects are cached in memory. For CRDs, the cache
is per-version. That is the result of having one [REST store object](https://github.com/kubernetes/kubernetes/blob/3cb771a8662ae7d1f79580e0ea9861fd6ab4ecc0/staging/src/k8s.io/apiextensions-apiserver/pkg/registry/customresource/etcd.go#L72) per-version which
was an arbitrary design choice but would be required for better caching with webhook conversion. In this model, each GVK is cached, regardless of whether some GVKs share storage.  Thus, watches do not cause conversion.  So, conversion webhooks will not add overhead to the watch path.  Watch cache is per api server and eventually consistent.

Non-watch reads are also cached (if requested resourceVersion is 0 which is true for generated informers by default, but not for calls like `kubectl get ...`, namespace cleanup, etc). The cached objects are converted and per-version (TODO: fact check). So, conversion webhooks will not add overhead here too.

If in the future this proves to be a performance problem, we might need to add caching later.  The Authorization and Authentication webhooks already use a simple scheme with APIserver-side caching and a single TTL for expiration.  This has worked fine, so we can repeat this process.  It does not require Webhook hosts to be aware of the caching.


## Examples


### Example of Writing Conversion Webhook

Data model for v1:

|data model for v1|
|-----------------|
```yaml      
properties:
  spec:
    properties:
      cronSpec:
        type: string
      image:
        type: string
```

|data model for v2|
|-----------------|
```yaml
properties:
  spec:
    properties:
      min:
        type: string
      hour:
        type: string
      dayOfMonth:
        type: string
      month:
        type: string
      dayOfWeek:
        type: string
      image:
        type: string
```


Both schemas can hold the same data (assuming the string format for V1 was a valid format).

|crontab_conversion.go|
|---------------------|

```golang
import .../types/v1
import .../types/v2

// Actual conversion methods

func convertCronV1toV2(cronV1 *v1.Crontab) (*v2.Crontab, error) {
  items := strings.Split(cronV1.spec.cronSpec, " ")
  if len(items) != 5 {
     return nil, fmt.Errorf("invalid spec string, needs five parts: %s", cronV1.spec.cronSpec)
  }
  return &v2.Crontab{
     ObjectMeta: cronV1.ObjectMeta,
     TypeMeta: metav1.TypeMeta{
        APIVersion: "stable.example.com/v2",
        Kind: cronV1.Kind,
     },
     spec: v2.CrontabSpec{
        image: cronV1.spec.image,
        min: items[0],
        hour: items[1],
        dayOfMonth: items[2],
        month: items[3],
        dayOfWeek: items[4],
     },
  }, nil

}

func convertCronV2toV1(cronV2 *v2.Crontab) (*v1.Crontab, error) {
  cronspec := cronV2.spec.min + " "
  cronspec += cronV2.spec.hour + " "
  cronspec += cronV2.spec.dayOfMonth + " "
  cronspec += cronV2.spec.month + " "
  cronspec += cronV2.spec.dayOfWeek
  return &v1.Crontab{
     ObjectMeta: cronV2.ObjectMeta,
     TypeMeta: metav1.TypeMeta{
        APIVersion: "stable.example.com/v1",
        Kind: cronV2.Kind,
     },
     spec: v1.CrontabSpec{
        image: cronV2.spec.image,
        cronSpec: cronspec,
     },
  }, nil
}

// The rest of the file can go into an auto generated framework

func serveCronTabConversion(w http.ResponseWriter, r *http.Request) {
  request, err := readConversionRequest(r)
  if err != nil {
     reportError(w, err)
  }
  response := ConversionResponse{}
  response.UID = request.UID
  converted, err := convert(request.Object, request.APIVersion)
  if err != nil {
     reportError(w, err)
  }
  response.ConvertedObject = *converted
  writeConversionResponse(w, response)
}

func convert(in runtime.RawExtension, version string) (*runtime.RawExtension, error) {
  inApiVersion, err := extractAPIVersion(in)
  if err != nil {
     return nil, err
  }
  switch inApiVersion {
  case "stable.example.com/v1":
     var cronV1 v1Crontab
     if err := json.Unmarshal(in.Raw, &cronV1); err != nil {
        return nil, err
     }
     switch version {
     case "stable.example.com/v1":
        // This should not happened as API server will not call the webhook in this case
        return &in, nil
     case "stable.example.com/v2":
        cronV2, err := convertCronV1toV2(&cronV1)
        if err != nil {
           return nil, err
        }
        raw, err := json.Marshal(cronV2)
        if err != nil {
           return nil, err
        }
        return &runtime.RawExtension{Raw: raw}, nil
     }
  case "stable.example.com/v2":
     var cronV2 v2Crontab
     if err := json.Unmarshal(in.Raw, &cronV2); err != nil {
        return nil, err
     }
     switch version {
     case "stable.example.com/v2":
        // This should not happened as API server will not call the webhook in this case
        return &in, nil
     case "stable.example.com/v1":
        cronV1, err := convertCronV2toV1(&cronV2)
        if err != nil {
           return nil, err
        }
        raw, err := json.Marshal(cronV1)
        if err != nil {
           return nil, err
        }
        return &runtime.RawExtension{Raw: raw}, nil
     }
  default:
     return nil, fmt.Errorf("invalid conversion fromVersion requested: %s", inApiVersion)
  }
  return nil, fmt.Errorf("invalid conversion toVersion requested: %s", version)
}

func extractAPIVersion(in runtime.RawExtension) (string, error) {
  object := unstructured.Unstructured{}
  if err := object.UnmarshalJSON(in.Raw); err != nil {
     return "", err
  }
  return object.GetAPIVersion(), nil
}
```

Note: not all code is shown for running a web server.  

Note: some of this is boilerplate that we expect tools like Kubebuilder will handle for the user.

Also some appropriate tests, most importantly round trip test:

|crontab_conversion_test.go|
|-|

```golang
func TestRoundTripFromV1ToV2(t *testing.T) {
  testObj := v1.Crontab{
     ObjectMeta: metav1.ObjectMeta{
        Name: "my-new-cron-object",
     },
     TypeMeta: metav1.TypeMeta{
        APIVersion: "stable.example.com/v1",
        Kind: "CronTab",
     },
     spec: v1.CrontabSpec{
        image: "my-awesome-cron-image",
        cronSpec: "* * * * */5",
     },
  }
  testRoundTripFromV1(t, testObj)
}

func testRoundTripFromV1(t *testing.T, v1Object v1.CronTab) {
  v2Object, err := convertCronV1toV2(v1Object)
  if err != nil {
     t.Fatalf("failed to convert v1 crontab to v2: %v", err)
  }
  v1Object2, err := convertCronV2toV1(v2Object)
  if err != nil {
     t.Fatalf("failed to convert v2 crontab to v1: %v", err)
  }
  if !reflect.DeepEqual(v1Object, v1Object2) {
     t.Errorf("round tripping failed for v1 crontab. v1Object: %v, v2Object: %v, v1ObjectConverted: %v",
        v1Object, v2Object, v1Object2)
  }
}
```

## Example of Updating CRD from one to two versions 

This example uses some files from previous section.

**Step 1**: Start from a CRD with only one version  

|crd1.yaml|
|-|

```yaml
apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  name: crontabs.stable.example.com
spec:
  group: stable.example.com
  versions:
  - name: v1
    served: true
    storage: true
    schema:
      properties:
        spec:
          properties:
            cronSpec:
              type: string
            image:
              type: string
  scope: Namespaced
  names:
    plural: crontabs
    singular: crontab
    kind: CronTab
    shortNames:
    - ct
```

And create it:

```bash
Kubectl create -f crd1.yaml
```

(If you have an existing CRD installed prior to the version of Kubernetes that supports the "versions" field, then you may need to move version field to a single item in the list of versions or just try to touch the CRD after upgrading to the new Kubernetes version which will result in the versions list being defaulted to a single item equal to the top level spec values)

**Step 2**: Create a CR within that one version:

|cr1.yaml|
|-|
```yaml

apiVersion: "stable.example.com/v1"
kind: CronTab
metadata:
  name: my-new-cron-object
spec:
  cronSpec: "* * * * */5"
  image: my-awesome-cron-image
```

And create it:

```bash
Kubectl create -f cr1.yaml
```

**Step 3**: Decide to introduce a new version of the API.

**Step 3a**: Write a new OpenAPI data model for the new version (see previous section).  Use of a data model is not required, but it is recommended.

**Step 3b**: Write conversion webhook and deploy it as a service named `crontab_conversion`

See the "crontab_conversion.go" file in the previous section.

**Step 3c**: Update the CRD to add the second version.

Do this by adding a new item to the "versions" list, containing the new data model:

|crd2.yaml|
|-|
```yaml

apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  name: crontabs.stable.example.com
spec:
  group: stable.example.com
  versions:
  - name: v1
    served: true
    storage: false
    schema:
      properties:
        spec:
          properties:
            cronSpec:
              type: string
            image:
              type: string
  - name: v2
    served: true
    storage: true
    schema:
      properties:
        spec:
          properties:
            min:
              type: string
            hour:
              type: string
            dayOfMonth:
              type: string
            month:
              type: string
            dayOfWeek:
              type: string
            image:
              type: string
  scope: Namespaced
  names:
    plural: crontabs
    singular: crontab
    kind: CronTab
    shortNames:
    - ct
  conversion:
    strategy: external
    webhook:
      client_config:
        namespace: crontab
        service: crontab_conversion
        Path: /crontab_convert
```

And apply it:

```bash
Kubectl apply -f crd2.yaml
```

**Step 4**: add a new CR in v2:

|cr2.yaml|
|-|
```yaml

apiVersion: "stable.example.com/v2"
kind: CronTab
metadata:
  name: my-second-cron-object
spec:
  min: "*"
  hour: "*"
  day_of_month: "*"
  dayOfWeek: "*/5"
  month: "*"
  image: my-awesome-cron-image
```

And create it:

```bash
Kubectl create -f cr2.yaml
```

**Step 5**: storage now has two custom resources in two different versions. To downgrade to previous CRD, one can apply crd1.yaml but that will fail as the status.storedVersions has both v1 and v2 and those cannot be removed from the spec.versions list. To downgrade, first create a crd2-b.yaml file that sets v1 as storage version and apply it, then follow "*Upgrade existing objects to a new stored version*“ in [this document](https://kubernetes.io/docs/tasks/access-kubernetes-api/custom-resources/custom-resource-definition-versioning/). After all CRs in the storage has v1 version, you can apply crd1.yaml.

**Step 5 alternative**: create a crd1-b.yaml that has v2 but not served.

|crd1-b.yaml|
|-|
```yaml

apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  name: crontabs.stable.example.com
spec:
  group: stable.example.com
  versions:
  - name: v1
    served: true
    storage: true
    schema:
      properties:
        spec:
          properties:
            cronSpec:
              type: string
            image:
              type: string
  - name: v2
    served: false
    storage: false
  scope: Namespaced
  names:
    plural: crontabs
    singular: crontab
    kind: CronTab
    shortNames:
    - ct
  conversion:
    strategy: external
    webhook:
      client_config:
        namespace: crontab
        service: crontab_conversion
        Path: /crontab_convert
```

## Alternatives Considered

Other than webhook conversion, a declarative conversion also considered and discussed. The main operator that being discussed was Rename/Move. This section explains why Webhooks are chosen over declarative conversion. This does not mean the declarative approach will not be supported by the webhook would be first conversion method kubernetes supports.

### Webhooks vs Declarative

The table below compares webhook vs declarative in details.

<table>
  <tr>
    <td></td>
    <td>Webhook</td>
    <td>Declarative</td>
  </tr>
  <tr>
    <td>1. Limitatisons</td>
    <td>There is no limitation on the type of conversion CRD author can do.</td>
    <td>Very limited set of conversions will be provided.</td>
  </tr>
  <tr>
    <td>2. User Complexity</td>
    <td>Harder to implement and the author needs to run an http server. This can be made simpler using tools such as kube-builder.</td>
    <td>Easy to use as they are in yaml configuration file.</td>
  </tr>
  <tr>
    <td>3. Design Complexity</td>
    <td>Because the API server calls into an external webhook, there is no need to design a specific conversions.</td>
    <td>Designing of declarative conversions can be tricky, especially if they are changing the value of fields. Challenges are: Meeting the round-trip-ability requirement, arguing the usefulness of the operator and keeping it simple enough for a declarative system.</td>
  </tr>
  <tr>
    <td>4. Performance</td>
    <td>Several calls to webhook for one operation (e.g. Apply) might hit performance issues. A monitoring metric helps measure this for later improvements that can be done through batch conversion.</td>
    <td>Implemented in API Server directly thus there is no performance concerns.</td>
  </tr>
  <tr>
    <td>5. User mistakes</td>
    <td>Users have freedom to implement any kind of conversion which may not conform with our API convention (e.g. round-tripability. If the conversion is not revertible, old clients may fail and downgrade will also be at risk).</td>
    <td>Keeping the conversion operators sane and sound would not be user’s problem. For things like rename/move there is already a design that keeps round-tripp-ability but that could be tricky for other operations.</td>
  </tr>
  <tr>
    <td>6. Popularity</td>
    <td>Because of the freedom in conversion of webhooks, they probably would be more popular</td>
    <td>Limited set of declarative operators make it a safer but less popular choice at least in the early stages of CRD development</td>
  </tr>
  <tr>
    <td>7. CRD Development Cycles</td>
    <td>Fit well into the story of CRD development of starting with blob store CRDs, then add Schema, then Add webhook conversions for the freedom of conversion the move as much possible to declarative for safer production.</td>
    <td>Comes after Webhooks in the development cycles of CRDs</td>
  </tr>
</table>


Webhook conversion has less limitation for the authors of APIs using CRD which is desirable especially in the early stages of development. Although there is a chance of user mistakes and also it may look more complex to implement a webhook, those can be relieved using sets of good tools/libraries such as kube-builder. Overall, Webhook conversion is the clear winner here. Declarative approach may be considered at a later stage as an alternative but need to be carefully designed.


### Caching

* use HTTP caching conventions with Cache-Control, Etags, and a unique URL for each different request).  This requires more complexity for the webhook author.  This change could be considered as part of an update to all 5 or so kinds of webhooks, but not justified for just this one kind of webhook.

* The CRD object could have a "conversionWebhookVersion" field which the user can increment/change when upgrading/downgrading the webhook to force invalidation of cached objects.


## Advice to Users 

* A proper webhook host implementation should accept every supported version as input and as output version. 

* It should also be able to round trip between versions. E.g. converting an object from v1 to v2 and back to v1 should yield the same object. 

* Consider testing your conversion webhook with a fuzz tester that generates random valid objects.

* The webhook should always give the same response with the same request that allows API server to potentially cache the responses in future (modulo bug fixes; when an update is pushed that fixes a bug in the conversion operation it might not take effect for a few minutes.

* If you need to add a new field, just add it.  You don't need new schema to add a field.

* Webhook Hosts should be side-effect free.

* Webhook Hosts should not expect to see every conversion operation.  Some may be cached in the future.

* Toolkits like KubeBuilder and OperatorKit may assist users in using this new feature by:

    * having a place in their file hierarchy to define multiple schemas for the same kind.

    * having a place in their code templates to define a conversion function.

    * generating a full Webhook Host from a conversion function.

    * helping users create tests by writing directories containing sample yamls of an object in various versions.

    * using fuzzing to generate random valid objects and checking if they convert.

## Test and Documentation Plan

* Test the upgrade/rollback scenario below.

* Test conversion, refer to the test case section.

* Document CRD conversion and best practices for webhook conversion

* Document to CRD users how to upgrade and downgrade (changing storage version dance, and changes to CRD stored tags).

### Upgrade/Rollback Scenarios

Scenario 1: Upgrading an Operator to have more versions.

* Detect if the cluster version supports webhook conversion

  * Helm chart can require e.g. v1.12 of a Kubernetes API Server.

Scenario 2:  Rolling back to a previous version of API Server that does not support CRD Conversions

* I have a cluster

  * I use apiserver v1.11.x, which supports multiple no-conversion-versions of a CRD

* I start to use CRDs 

  * I install helm chart "Foo-Operator", which installs a CRD for resource Foo, with 1 version called v1beta1.

    * This uses the old "version" and "

    * I create some Foo resources. 

* I upgrade apiserver to v1.12.x

  * version-conversion now supported.

* I upgrade the Foo-Operator chart.  

  * This changes the CRD to have two versions, v1beta1 and v1beta2.

  * It installs a Webhook Host to convert them.

  * Assume: v1beta1 is still the storage version.

* I start using multiple versions, so that the CRs are now stored in a mix of versions.

* I downgrade kube-apiserver

  * Emergency happens, I need to downgrade to v1.11.x.  Conversion won't be possible anymore.

  * Downgrade

  * Any call needs conversion should fail at this stage (we need to patch 1.11 for this, see issue [#65790](https://github.com/kubernetes/kubernetes/issues/65790)

### Test Cases

* Updating existing CRD to use multiple versions with conversion

  * Define a CRD with one version.  

  * Create stored CRs.  

  * Update the CRD object to add another (non-storage) version with a conversion webhook

  * Existing CRs are not harmed

  * Can get existing CRs via new api, conversion webhook should be called

  * Can create new CRs with new api, conversion webhook should be called

  * Access new CRs with new api, conversion webhook should not be called

  * Access new CRs with old api, conversion webhook should be called

## Development Plan

Google able to staff development, test, review, and documentation. Help welcome, too, esp. Reviewing.

Not in scope for this work:

* Including CRDs to aggregated OpenAPI spec (fka swagger.json).

* Apply for CRDs

* Make CRDs powerful enough to convert any or all core types to CRDs (in line with that goal, but this is just a step towards it).

### Work items

* Add APIs for conversion webhooks in CustomResourceDefinition type.

* Support multi-version (used to be called validation) Schema

* Support multi-version subresources and AdditionalPrintColumns

* Add a Webhook converter call as a CRD converter (refactor conversion code as needed)

* Ensure able to monitor latency from webhooks. See Monitorability section

* Add Upgrade/Downgrade tests

* Add public documentation
