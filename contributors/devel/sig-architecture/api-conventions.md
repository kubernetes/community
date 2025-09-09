API Conventions
===============

*This document is oriented at users who want a deeper understanding of the
Kubernetes API structure, and developers wanting to extend the Kubernetes API.
An introduction to using resources with kubectl can be found in [the object management overview](https://kubernetes.io/docs/concepts/overview/working-with-objects/object-management/).*

**Table of Contents**

- [Types (Kinds)](#types-kinds)
  - [Resources](#resources)
  - [Objects](#objects)
    - [Metadata](#metadata)
    - [Spec and Status](#spec-and-status)
      - [Typical status properties](#typical-status-properties)
    - [References to related objects](#references-to-related-objects)
    - [Lists of named subobjects preferred over maps](#lists-of-named-subobjects-preferred-over-maps)
    - [Primitive types](#primitive-types)
    - [Constants](#constants)
    - [Unions](#unions)
  - [Lists and Simple kinds](#lists-and-simple-kinds)
- [Differing Representations](#differing-representations)
- [Verbs on Resources](#verbs-on-resources)
  - [PATCH operations](#patch-operations)
- [Short-names and Categories](#short-names-and-categories)
  - [Short-names](#short-names)
  - [Categories](#categories)
- [Idempotency](#idempotency)
- [Optional vs. Required](#optional-vs-required)
- [Defaulting](#defaulting)
  - [Static Defaults](#static-defaults)
  - [Admission Controlled Defaults](#admission-controlled-defaults)
  - [Controller-Assigned Defaults (aka Late Initialization)](#controller-assigned-defaults-aka-late-initialization)
  - [What May Be Defaulted](#what-may-be-defaulted)
  - [Considerations For PUT Operations](#considerations-for-put-operations)
- [Concurrency Control and Consistency](#concurrency-control-and-consistency)
- [Serialization Format](#serialization-format)
- [Units](#units)
- [Selecting Fields](#selecting-fields)
- [Object references](#object-references)
  - [Naming of the reference field](#naming-of-the-reference-field)
  - [Referencing resources with multiple versions](#referencing-resources-with-multiple-versions)
  - [Handling of resources that do not exist](#handling-of-resources-that-do-not-exist)
  - [Validation of fields](#validation-of-fields)
  - [Do not modify the referred object](#do-not-modify-the-referred-object)
  - [Minimize copying or printing values to the referrer object](#minimize-copying-or-printing-values-to-the-referrer-object)
  - [Object References Examples](#object-references-examples)
    - [Single resource reference](#single-resource-reference)
      - [Controller behavior](#controller-behavior)
    - [Multiple resource reference](#multiple-resource-reference)
      - [Kind vs. Resource](#kind-vs-resource)
      - [Controller behavior](#controller-behavior-1)
    - [Generic object reference](#generic-object-reference)
      - [Controller behavior](#controller-behavior-2)
    - [Field reference](#field-reference)
      - [Controller behavior](#controller-behavior-3)
- [HTTP Status codes](#http-status-codes)
    - [Success codes](#success-codes)
    - [Error codes](#error-codes)
- [Response Status Kind](#response-status-kind)
- [Events](#events)
- [Naming conventions](#naming-conventions)
  - [Namespace Names](#namespace-names)
- [Label, selector, and annotation conventions](#label-selector-and-annotation-conventions)
- [WebSockets and SPDY](#websockets-and-spdy)
- [Validation](#validation)
- [Automatic Resource Allocation And Deallocation](#automatic-resource-allocation-and-deallocation)
- [Representing Allocated Values](#representing-allocated-values)
  - [When to use a <code>spec</code> field](#when-to-use-a-spec-field)
  - [When to use a <code>status</code> field](#when-to-use-a-status-field)
    - [Sequencing operations](#sequencing-operations)
  - [When to use a different type](#when-to-use-a-different-type)


The conventions of the [Kubernetes API](https://kubernetes.io/docs/concepts/overview/kubernetes-api/) (and related APIs in the
ecosystem) are intended to ease client development and ensure that configuration
mechanisms can be implemented that work across a diverse set of use cases
consistently.

The general style of the Kubernetes API is RESTful - clients create, update,
delete, or retrieve a description of an object via the standard HTTP verbs
(POST, PUT, DELETE, and GET) - and those APIs preferentially accept and return
JSON. Kubernetes also exposes additional endpoints for non-standard verbs and
allows alternative content types. All of the JSON accepted and returned by the
server has a schema, identified by the "kind" and "apiVersion" fields. Where
relevant HTTP header fields exist, they should mirror the content of JSON
fields, but the information should not be represented only in the HTTP header.

The following terms are defined:

* **Kind** the name of a particular object schema (e.g. the "Cat" and "Dog"
kinds would have different attributes and properties)
* **Resource** a representation of a system entity, sent or retrieved as JSON
via HTTP to the server. Resources are exposed via:
  * Collections - a list of resources of the same type, which may be queryable
  * Elements - an individual resource, addressable via a URL
* **API Group** a set of resources that are exposed together, along
with the version exposed in the "apiVersion" field as "GROUP/VERSION", e.g.
"policy.k8s.io/v1".

Each resource typically accepts and returns data of a single kind. A kind may be
accepted or returned by multiple resources that reflect specific use cases. For
instance, the kind "Pod" is exposed as a "pods" resource that allows end users
to create, update, and delete pods, while a separate "pod status" resource (that
acts on "Pod" kind) allows automated processes to update a subset of the fields
in that resource.

Resources are bound together in API groups - each group may have one or more
versions that evolve independent of other API groups, and each version within
the group has one or more resources. Group names are typically in domain name
form - the Kubernetes project reserves use of the empty group, all single
word names ("extensions", "apps"), and any group name ending in "*.k8s.io" for
its sole use. When choosing a group name, we recommend selecting a subdomain
your group or organization owns, such as "widget.mycompany.com".

Version strings should match
[DNS_LABEL](https://git.k8s.io/design-proposals-archive/architecture/identifiers.md)
format.



Resource collections should be all lowercase and plural, whereas kinds are
CamelCase and singular. Group names must be lower case and be valid DNS
subdomains.


## Types (Kinds)

Kinds are grouped into three categories:

1. **Objects** represent a persistent entity in the system.

   Creating an API object is a record of intent - once created, the system will
work to ensure that resource exists. All API objects have common metadata.

   An object may have multiple resources that clients can use to perform
specific actions that create, update, delete, or get.

   Examples: `Pod`, `ReplicationController`, `Service`, `Namespace`, `Node`.

2. **Lists** are collections of **resources** of one (usually) or more
(occasionally) kinds.

   The name of a list kind must end with "List". Lists have a limited set of
common metadata. All lists use the required "items" field to contain the array
of objects they return. Any kind that has the "items" field must be a list kind.

   Most objects defined in the system should have an endpoint that returns the
full set of resources, as well as zero or more endpoints that return subsets of
the full list. Some objects may be singletons (the current user, the system
defaults) and may not have lists.

   In addition, all lists that return objects with labels should support label
filtering (see [the labels documentation](https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/)),
and most lists should support filtering by fields (see
[the fields documentation](https://kubernetes.io/docs/concepts/overview/working-with-objects/field-selectors/)).

   Examples: `PodList`, `ServiceList`, `NodeList`.

Note that`kubectl` and other tools sometimes output collections of resources
as `kind: List`. Keep in mind that `kind: List` is not part of the Kubernetes API; it is
exposing an implementation detail from client-side code in those tools, used to
handle groups of mixed resources.

3. **Simple** kinds are used for specific actions on objects and for
non-persistent entities.

   Given their limited scope, they have the same set of limited common metadata
as lists.

   For instance, the "Status" kind is returned when errors occur and is not
persisted in the system.

   Many simple resources are "subresources", which are rooted at API paths of
specific resources. When resources wish to expose alternative actions or views
that are closely coupled to a single resource, they should do so using new
sub-resources. Common subresources include:

   * `/binding`: Used to bind a resource representing a user request (e.g., Pod,
PersistentVolumeClaim) to a cluster infrastructure resource (e.g., Node,
PersistentVolume).
   * `/status`: Used to write just the `status` portion of a resource. For
example, the `/pods` endpoint only allows updates to `metadata` and `spec`,
since those reflect end-user intent. An automated process should be able to
modify status for users to see by sending an updated Pod kind to the server to
the "/pods/&lt;name&gt;/status" endpoint - the alternate endpoint allows
different rules to be applied to the update, and access to be appropriately
restricted.
   * `/scale`: Used to read and write the count of a resource in a manner that
is independent of the specific resource schema.

   Two additional subresources, `proxy` and `portforward`, provide access to
cluster resources as described in
[accessing the cluster](https://kubernetes.io/docs/tasks/access-application-cluster/access-cluster/).

The standard REST verbs (defined below) MUST return singular JSON objects. Some
API endpoints may deviate from the strict REST pattern and return resources that
are not singular JSON objects, such as streams of JSON objects or unstructured
text log data.

A common set of "meta" API objects are used across all API groups and are
thus considered part of the API group named `meta.k8s.io`. These types may
evolve independent of the API group that uses them and API servers may allow
them to be addressed in their generic form. Examples are `ListOptions`,
`DeleteOptions`, `List`, `Status`, `WatchEvent`, and `Scale`. For historical
reasons these types are part of each existing API group. Generic tools like
quota, garbage collection, autoscalers, and generic clients like kubectl
leverage these types to define consistent behavior across different resource
types, like the interfaces in programming languages.

The term "kind" is reserved for these "top-level" API types. The term "type"
should be used for distinguishing sub-categories within objects or subobjects.

### Resources

All JSON objects returned by an API MUST have the following fields:

* kind: a string that identifies the schema this object should have
* apiVersion: a string that identifies the version of the schema the object
should have

These fields are required for proper decoding of the object. They may be
populated by the server by default from the specified URL path, but the client
likely needs to know the values in order to construct the URL path.

### Objects

#### Metadata

Every object kind MUST have the following metadata in a nested object field
called "metadata":

* namespace: a namespace is a DNS compatible label that objects are subdivided
into. The default namespace is 'default'. See
[the namespace docs](https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/) for more.
* name: a string that uniquely identifies this object within the current
namespace (see [the identifiers docs](https://kubernetes.io/docs/concepts/overview/working-with-objects/names/)).
This value is used in the path when retrieving an individual object.
* uid: a unique in time and space value (typically an RFC 4122 generated
identifier, see [the identifiers docs](https://kubernetes.io/docs/concepts/overview/working-with-objects/names/))
used to distinguish between objects with the same name that have been deleted
and recreated

Every object SHOULD have the following metadata in a nested object field called
"metadata":

* resourceVersion: a string that identifies the internal version of this object
that can be used by clients to determine when objects have changed. This value
MUST be treated as opaque by clients and passed unmodified back to the server.
Clients should not assume that the resource version has meaning across
namespaces, different kinds of resources, or different servers. (See
[concurrency control](#concurrency-control-and-consistency), below, for more
details.)
* generation: a sequence number representing a specific generation of the
desired state. Set by the system and monotonically increasing, per-resource. May
be compared, such as for RAW and WAW consistency.
* creationTimestamp: a string representing an RFC 3339 date of the date and time
an object was created
* deletionTimestamp: a string representing an RFC 3339 date of the date and time
after which this resource will be deleted. This field is set by the server when
a graceful deletion is requested by the user, and is not directly settable by a
client. The resource will be deleted (no longer visible from resource lists, and
not reachable by name) after the time in this field except when the object has
a finalizer set. In case the finalizer is set the deletion of the object is
postponed at least until the finalizer is removed.
Once the deletionTimestamp is set, this value may not be unset or be set further
into the future, although it may be shortened or the resource may be deleted
prior to this time.
* labels: a map of string keys and values that can be used to organize and
categorize objects (see [the labels docs](https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/))
* annotations: a map of string keys and values that can be used by external
tooling to store and retrieve arbitrary metadata about this object (see
[the annotations docs](https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/))

Labels are intended for organizational purposes by end users (select the pods
that match this label query). Annotations enable third-party automation and
tooling to decorate objects with additional metadata for their own use.

#### Spec and Status

By convention, the Kubernetes API makes a distinction between the specification
of the desired state of an object (a nested object field called `spec`) and the
status of the object at the current time (a nested object field called
`status`). The specification is a complete description of the desired state,
including configuration settings provided by the user,
[default values](#defaulting) expanded by the system, and properties initialized
or otherwise changed after creation by other ecosystem components (e.g.,
schedulers, auto-scalers), and is persisted in stable storage with the API
object. If the specification is deleted, the object will be purged from the
system.

The `status` summarizes the current state of the object in the system, and is
usually persisted with the object by automated processes but may be generated
on the fly.  As a general guideline, fields in `status` should be the most recent
observations of actual state, but they may contain information such as the
results of allocations or similar operations which are executed in response to
the object's `spec`.  See [below](#representing-allocated-values) for more
details.

Types with both `spec` and `status` stanzas can (and usually should) have distinct
authorization scopes for them.  This allows users to be granted full write
access to `spec` and read-only access to status, while relevant controllers are
granted read-only access to `spec` but full write access to status.

When a new version of an object is POSTed or PUT, the `spec` is updated and
available immediately. Over time the system will work to bring the `status` into
line with the `spec`. The system will drive toward the most recent `spec`
regardless of previous versions of that stanza. For example, if a value is
changed from 2 to 5 in one PUT and then back down to 3 in another PUT the system
is not required to 'touch base' at 5 before changing the `status` to 3. In other
words, the system's behavior is *level-based* rather than *edge-based*. This
enables robust behavior in the presence of missed intermediate state changes.

The Kubernetes API also serves as the foundation for the declarative
configuration schema for the system. In order to facilitate level-based
operation and expression of declarative configuration, fields in the
specification should have declarative rather than imperative names and
semantics -- they represent the desired state, not actions intended to yield the
desired state.

The PUT and POST verbs on objects MUST ignore the `status` values, to avoid
accidentally overwriting the `status` in read-modify-write scenarios. A `/status`
subresource MUST be provided to enable system components to update statuses of
resources they manage.

Otherwise, PUT expects the whole object to be specified. Therefore, if a field
is omitted it is assumed that the client wants to clear that field's value. The
PUT verb does not accept partial updates. Modification of just part of an object
may be achieved by GETting the resource, modifying part of the spec, labels, or
annotations, and then PUTting it back. See
[concurrency control](#concurrency-control-and-consistency), below, regarding
read-modify-write consistency when using this pattern. Some objects may expose
alternative resource representations that allow mutation of the status, or
performing custom actions on the object.

All objects that represent a physical resource whose state may vary from the
user's desired intent SHOULD have a `spec` and a `status`. Objects whose state
cannot vary from the user's desired intent MAY have only `spec`, and MAY rename
`spec` to a more appropriate name.

Objects that contain both `spec` and `status` should not contain additional
top-level fields other than the standard metadata fields.

Some objects which are not persisted in the system - such as `SubjectAccessReview`
and other webhook style calls - may choose to add `spec` and `status` to encapsulate
a "call and response" pattern. The `spec` is the request (often a request for
information) and the `status` is the response. For these RPC like objects the only
operation may be POST, but having a consistent schema between submission and
response reduces the complexity of these clients.

##### Typical status properties

**Conditions** provide a standard mechanism for higher-level status reporting
from a controller. They are an extension mechanism which allows tools and other
controllers to collect summary information about resources without needing to
understand resource-specific status details. Conditions should complement more
detailed information about the observed status of an object written by a
controller, rather than replace it. For example, the "Available" condition of a
Deployment can be determined by examining `readyReplicas`, `replicas`, and
other properties of the Deployment. However, the "Available" condition allows
other components to avoid duplicating the availability logic in the Deployment
controller.

Objects may report multiple conditions, and new types of conditions may be
added in the future or by 3rd party controllers. Therefore, conditions are
represented using a list/slice of objects, where each condition has a similar
structure. This collection should be treated as a map with a key of `type`.

Conditions are most useful when they follow some consistent conventions:

* Conditions should be added to explicitly convey properties that users and
  components care about rather than requiring those properties to be inferred
  from other observations.  Once defined, the meaning of a Condition can not be
  changed arbitrarily - it becomes part of the API, and has the same backwards-
  and forwards-compatibility concerns of any other part of the API.

* Controllers should apply their conditions to a resource the first time they
  visit the resource, even if the `status` is Unknown. This allows other
  components in the system to know that the condition exists and the controller
  is making progress on reconciling that resource.

   * Not all controllers will observe the previous advice about reporting
     "Unknown" or "False" values. For known conditions, the absence of a
     condition `status` should be interpreted the same as `Unknown`, and
     typically indicates that reconciliation has not yet finished (or that the
     resource state may not yet be observable).

* For some conditions, `True` represents normal operation, and for some
  conditions, `False` represents normal operation. ("Normal-true" conditions
  are sometimes said to have "positive polarity", and "normal-false" conditions
  are said to have "negative polarity".) Without further knowledge of the
  conditions, it is not possible to compute a generic summary of the conditions
  on a resource.

* Condition type names should make sense for humans; neither positive nor
  negative polarity can be recommended as a general rule. A negative condition
  like "MemoryExhausted" may be easier for humans to understand than
  "SufficientMemory". Conversely, "Ready" or "Succeeded" may be easier to
  understand than "Failed", because "Failed=Unknown" or "Failed=False" may
  cause double-negative confusion.

* Condition type names should describe the current observed state of the
  resource, rather than describing the current state transitions. This
  typically means that the name should be an adjective ("Ready", "OutOfDisk")
  or a past-tense verb ("Succeeded", "Failed") rather than a present-tense verb
  ("Deploying"). Intermediate states may be indicated by setting the `status` of
  the condition to `Unknown`.

  * For state transitions which take a long period of time (e.g. more than 1
    minute), it is reasonable to treat the transition itself as an observed
    state. In these cases, the Condition (such as "Resizing") itself should not
    be transient, and should instead be signalled using the
    `True`/`False`/`Unknown` pattern. This allows other observers to determine
    the last update from the controller, whether successful or failed. In cases
    where the state transition is unable to complete and continued
    reconciliation is not feasible, the Reason and Message should be used to
    indicate that the transition failed.

* When designing Conditions for a resource, it's helpful to have a common
  top-level condition which summarizes more detailed conditions. Simple
  consumers may simply query the top-level condition. Although they are not a
  consistent standard, the `Ready` and `Succeeded` condition types may be used
  by API designers for long-running and bounded-execution objects, respectively.

Conditions should follow the standard schema included in [k8s.io/apimachinery/pkg/apis/meta/v1/types.go](https://github.com/kubernetes/apimachinery/blob/release-1.23/pkg/apis/meta/v1/types.go#L1432-L1492).
It should be included as a top level element in status, similar to
```go
// +listType=map
// +listMapKey=type
// +patchStrategy=merge
// +patchMergeKey=type
// +optional
Conditions []metav1.Condition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type" protobuf:"bytes,1,rep,name=conditions"`
```

The `metav1.Conditions` includes the following fields

```go
// type of condition in CamelCase or in foo.example.com/CamelCase.
// +required
Type string `json:"type" protobuf:"bytes,1,opt,name=type"`
// status of the condition, one of True, False, Unknown.
// +required
Status ConditionStatus `json:"status" protobuf:"bytes,2,opt,name=status"`
// observedGeneration represents the .metadata.generation that the condition was set based upon.
// For instance, if .metadata.generation is currently 12, but the .status.conditions[x].observedGeneration is 9, the condition is out of date
// with respect to the current state of the instance.
// +optional
ObservedGeneration int64 `json:"observedGeneration,omitempty" protobuf:"varint,3,opt,name=observedGeneration"`
// lastTransitionTime is the last time the condition transitioned from one status to another.
// This should be when the underlying condition changed.  If that is not known, then using the time when the API field changed is acceptable.
// +required
LastTransitionTime Time `json:"lastTransitionTime" protobuf:"bytes,4,opt,name=lastTransitionTime"`
// reason contains a programmatic identifier indicating the reason for the condition's last transition.
// Producers of specific condition types may define expected values and meanings for this field,
// and whether the values are considered a guaranteed API.
// The value should be a CamelCase string.
// This field may not be empty.
// +required
Reason string `json:"reason" protobuf:"bytes,5,opt,name=reason"`
// message is a human readable message indicating details about the transition.
// This may be an empty string.
// +required
Message string `json:"message" protobuf:"bytes,6,opt,name=message"`
```

Additional fields may be added in the future.

Use of the `Reason` field is required.

Condition types should be named in PascalCase. Short condition names are
preferred (e.g. "Ready" over "MyResourceReady").

Condition `status` values may be `True`, `False`, or `Unknown`. The absence of a
condition should be interpreted the same as `Unknown`.  How controllers handle
`Unknown` depends on the Condition in question.

The thinking around conditions has evolved over time, so there are several
non-normative examples in wide use.

In general, condition values may change back and forth, but some condition
transitions may be monotonic, depending on the resource and condition type.
However, conditions are observations and not, themselves, state machines, nor do
we define comprehensive state machines for objects, nor behaviors associated
with state transitions. The system is level-based rather than edge-triggered,
and should assume an Open World.

An example of an oscillating condition type is `Ready`, which indicates the
object was believed to be fully operational at the time it was last probed. A
possible monotonic condition could be `Succeeded`. A `True` status for
`Succeeded` would imply completion and that the resource was no longer
active. An object that was still active would generally have a `Succeeded`
condition with status `Unknown`.

Some resources in the v1 API contain fields called **`phase`**, and associated
`message`, `reason`, and other status fields. The pattern of using `phase` is
deprecated. Newer API types should use conditions instead. Phase was
essentially a state-machine enumeration field, that contradicted [system-design
principles](https://git.k8s.io/design-proposals-archive/architecture/principles.md#control-logic) and
hampered evolution, since [adding new enum values breaks backward
compatibility](api_changes.md). Rather than encouraging clients to infer
implicit properties from phases, we prefer to explicitly expose the individual
conditions that clients need to monitor. Conditions also have the benefit that
it is possible to create some conditions with uniform meaning across all
resource types, while still exposing others that are unique to specific
resource types.  See [#7856](http://issues.k8s.io/7856) for more details and
discussion.

In condition types, and everywhere else they appear in the API, **`Reason`** is
intended to be a one-word, CamelCase representation of the category of cause of
the current status, and **`Message`** is intended to be a human-readable phrase
or sentence, which may contain specific details of the individual occurrence.
`Reason` is intended to be used in concise output, such as one-line
`kubectl get` output, and in summarizing occurrences of causes, whereas
`Message` is intended to be presented to users in detailed status explanations,
such as `kubectl describe` output.

Historical information status (e.g., last transition time, failure counts) is
only provided with reasonable effort, and is not guaranteed to not be lost.

Status information that may be large (especially proportional in size to
collections of other resources, such as lists of references to other objects --
see below) and/or rapidly changing, such as
[resource usage](https://git.k8s.io/design-proposals-archive/scheduling/resources.md#usage-data), should be put into separate
objects, with possibly a reference from the original object. This helps to
ensure that GETs and watch remain reasonably efficient for the majority of
clients, which may not need that data.

Some resources report the `observedGeneration`, which is the `generation` most
recently observed by the component responsible for acting upon changes to the
desired state of the resource. This can be used, for instance, to ensure that
the reported status reflects the most recent desired status.

#### References to related objects

References to loosely coupled sets of objects, such as
[pods](https://kubernetes.io/docs/concepts/workloads/pods/) overseen by a
[replication controller](https://kubernetes.io/docs/concepts/workloads/controllers/replicationcontroller/),
are usually best referred to using a
[label selector](https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#label-selectors). In order to
ensure that GETs of individual objects remain bounded in time and space, these
sets may be queried via separate API queries, but will not be expanded in the
referring object's status.

For references to specific objects, see [Object references](#object-references).

References in the `status` of the referee to the referrer may be permitted, when
the references are one-to-one and do not need to be frequently updated,
particularly in an edge-based manner.

#### Lists of named subobjects preferred over maps

Discussed in [#2004](http://issue.k8s.io/2004) and elsewhere. There are
no maps of subobjects in any API objects. Instead, the convention is to
use a list of subobjects containing name fields. These conventions, and
how one can change the semantics of lists, structs and maps are
described in more details in the Kubernetes
[documentation](https://kubernetes.io/docs/reference/using-api/server-side-apply/#merge-strategy).

For example:

```yaml
ports:
  - name: www
    containerPort: 80
```

vs.

```yaml
ports:
  www:
    containerPort: 80
```

This rule maintains the invariant that all JSON/YAML keys are fields in API
objects. The only exceptions are pure maps in the API (currently, labels,
selectors, annotations, data), as opposed to sets of subobjects.

#### Primitive types

* Look at similar fields in the API (e.g. ports, durations) and follow the
  conventions of existing fields.
* Do not use enums. Use aliases for string instead (e.g. `NodeConditionType`).
* All numeric fields should be bounds-checked, both for too-small or negative
  and for too-large.
* All public integer fields MUST use the Go `int32` or Go `int64` types, not
  `int` (which is ambiguously sized, depending on target platform).  Internal
  types may use `int`.
* For integer fields, prefer `int32` to `int64` unless you need to represent
  values larger than `int32`.  See other guidelines about limitations of
  `int64` and language compatibility.
* Do not use unsigned integers, due to inconsistent support across languages and
  libraries. Just validate that the integer is non-negative if that's the case.
* All numbers (e.g. `int32`, `int64`) are converted to `float64` by Javascript
  and some other languages, so any field which is expected to exceed that
  either in magnitude or in precision (e.g. integer values > 53 bits)
  should be serialized and accepted as strings. `int64` fields must be
  bounds-checked to be within the range of `-(2^53) < x < (2^53)`.
* Avoid floating-point values as much as possible, and never use them in spec.
  Floating-point values cannot be reliably round-tripped (encoded and
  re-decoded) without changing, and have varying precision and representations
  across languages and architectures.
* Think twice about `bool` fields. Many ideas start as boolean but eventually
  trend towards a small set of mutually exclusive options.  Plan for future
  expansions by describing the policy options explicitly as a string type
  alias (e.g. `TerminationMessagePolicy`).

#### Constants

Some fields will have a list of allowed values (enumerations). These values will
be strings, and they will be in CamelCase, with an initial uppercase letter.
Examples: `ClusterFirst`, `Pending`, `ClientIP`. When an acronym or initialism
each letter in the acronym should be uppercase, such as with `ClientIP` or
`TCPDelay`. When a proper name or the name of a command-line executable is used
as a constant the proper name should be represented in consistent casing -
examples: `systemd`, `iptables`, `IPVS`, `cgroupfs`, `Docker` (as a generic
concept), `docker` (as the command-line executable). If a proper name is used
which has mixed capitalization like `eBPF` that should be preserved in a longer
constant such as `eBPFDelegation`.

All API within Kubernetes must leverage constants in this style, including
flags and configuration files. Where inconsistent constants were previously used,
new flags should be CamelCase only, and over time old flags should be updated to
accept a CamelCase value alongside the inconsistent constant. Example: the
Kubelet accepts a `--topology-manager-policy` flag that has values `none`,
`best-effort`, `restricted`, and `single-numa-node`. This flag should accept
`None`, `BestEffort`, `Restricted`, and `SingleNUMANode` going forward. If new
values are added to the flag, both forms should be supported.

#### Unions

Sometimes, at most one of a set of fields can be set.  For example, the
[volumes] field of a PodSpec has 17 different volume type-specific fields, such
as `nfs` and `iscsi`.  All fields in the set should be
[Optional](#optional-vs-required).

Sometimes, when a new type is created, the api designer may anticipate that a
union will be needed in the future, even if only one field is allowed initially.
In this case, be sure to make the field [Optional](#optional-vs-required)
In the validation, you may still return an error if the sole field is unset. Do
not set a default value for that field.

### Lists and Simple kinds

Every list or simple kind SHOULD have the following metadata in a nested object
field called "metadata":

* resourceVersion: a string that identifies the common version of the objects
returned by in a list. This value MUST be treated as opaque by clients and
passed unmodified back to the server. A resource version is only valid within a
single namespace on a single kind of resource.

Every simple kind returned by the server, and any simple kind sent to the server
that must support idempotency or optimistic concurrency should return this
value. Since simple resources are often used as input alternate actions that
modify objects, the resource version of the simple resource should correspond to
the resource version of the object.


## Differing Representations

An API may represent a single entity in different ways for different clients, or
transform an object after certain transitions in the system occur. In these
cases, one request object may have two representations available as different
resources, or different kinds.

An example is a Service, which represents the intent of the user to group a set
of pods with common behavior on common ports. When Kubernetes detects a pod
matches the service selector, the IP address and port of the pod are added to an
Endpoints resource for that Service. The Endpoints resource exists only if the
Service exists, but exposes only the IPs and ports of the selected pods. The
full service is represented by two distinct resources - under the original
Service resource the user created, as well as in the Endpoints resource.

As another example, a "pod status" resource may accept a PUT with the "pod"
kind, with different rules about what fields may be changed.

Future versions of Kubernetes may allow alternative encodings of objects beyond
JSON.


## Verbs on Resources

API resources should use the traditional REST pattern:

* GET /&lt;resourceNamePlural&gt; - Retrieve a list of type
&lt;resourceName&gt;, e.g. GET /pods returns a list of Pods.
* POST /&lt;resourceNamePlural&gt; - Create a new resource from the JSON object
provided by the client.
* GET /&lt;resourceNamePlural&gt;/&lt;name&gt; - Retrieves a single resource
with the given name, e.g. GET /pods/first returns a Pod named 'first'. Should be
constant time, and the resource should be bounded in size.
* DELETE /&lt;resourceNamePlural&gt;/&lt;name&gt;  - Delete the single resource
with the given name. DeleteOptions may specify gracePeriodSeconds, the optional
duration in seconds before the object should be deleted. Individual kinds may
declare fields which provide a default grace period, and different kinds may
have differing kind-wide default grace periods. A user provided grace period
overrides a default grace period, including the zero grace period ("now").
* DELETE /&lt;resourceNamePlural&gt; - Deletes a list of type
&lt;resourceName&gt;, e.g. DELETE /pods a list of Pods.
* PUT /&lt;resourceNamePlural&gt;/&lt;name&gt; - Update or create the resource
with the given name with the JSON object provided by the client. Whether a
resource can be created with a PUT request depends on the particular resource's
storage strategy configuration, specifically the `AllowCreateOnUpdate()` return
value. Most built-in types do not allow this.
* PATCH /&lt;resourceNamePlural&gt;/&lt;name&gt; - Selectively modify the
specified fields of the resource. See more information [below](#patch-operations).
* GET /&lt;resourceNamePlural&gt;&quest;watch=true - Receive a stream of JSON
objects corresponding to changes made to any resource of the given kind over
time.

### PATCH operations

The API supports three different PATCH operations, determined by their
corresponding Content-Type header:

* JSON Patch, `Content-Type: application/json-patch+json`
  * As defined in [RFC6902](https://tools.ietf.org/html/rfc6902), a JSON Patch is
a sequence of operations that are executed on the resource, e.g. `{"op": "add",
"path": "/a/b/c", "value": [ "foo", "bar" ]}`. For more details on how to use
JSON Patch, see the RFC.
* Merge Patch, `Content-Type: application/merge-patch+json`
  * As defined in [RFC7386](https://tools.ietf.org/html/rfc7386), a Merge Patch
is essentially a partial representation of the resource. The submitted JSON is
"merged" with the current resource to create a new one, then the new one is
saved. For more details on how to use Merge Patch, see the RFC.
* Strategic Merge Patch, `Content-Type: application/strategic-merge-patch+json`
  * Strategic Merge Patch is a custom implementation of Merge Patch. For a
detailed explanation of how it works and why it needed to be introduced, see
[here](/contributors/devel/sig-api-machinery/strategic-merge-patch.md).

## Short-names and Categories

Resource implementers can optionally include "short names" and categories
in the discovery information published for a resource type,
which clients may use as hints when resolving ambiguous user invocations.

For compiled-in resources, these are controlled by the REST handler `ShortNames() []string` and `Categories() []string` implementations.

For custom resources, these are controlled by the `.spec.names.shortNames` and `.spec.names.categories` fields in the CustomResourceDefinition.

### Short-names

Note: Due to unpredictable behavior when short names collide (with each other or with resource types),
do not add new short names to built-in resources unless specifically allowed by API reviewers. See issues
[#117742](https://issue.k8s.io/117742#issuecomment-1545945336) and [#108573](http://issue.k8s.io/108573).

"Short names" listed in discovery may be used by clients as hints to resolve ambiguous user invocations to a single resource.

Examples of built-in short names include:

* `ds` -> `apps/v* daemonsets`
* `sts` -> `apps/v* statefulsets`
* `hpa` -> `autoscaling/v* horizontalpodautoscalers`

For example, with only built-in API types served, `kubectl get sts` is equivalent to `kubectl get statefulsets.v1.apps`.

Short-name matches may be given lower priority than an exact match of a resource type,
so use of short names increases potential for inconsistent behavior in clusters
with custom resources installed, if those custom resource types overlap with short names.

Continuing the above example, if a custom resource with `.spec.names.plural` set to `sts` was installed in a cluster,
`kubectl get sts` would switch to retrieving instances of the custom resource instead.

### Categories

Note: Due to inconsistent behavior when categories collide with resource types,
and difficulties knowing when it is safe to add new resources to an existing category,
do not add new categories to built-in resources unless specifically allowed by API reviewers.
See issues [#7547](https://github.com/kubernetes/kubernetes/issues/7547#issuecomment-355835279)
[#42885](https://github.com/kubernetes/kubernetes/issues/42885#issuecomment-531265679),
and [considerations for adding to the "all" category](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-cli/kubectl-conventions.md#rules-for-extending-special-resource-alias---all)
for examples of the difficulties encountered.

Categories listed in discovery may be used by clients as hints to resolve user invocations to multiple resources.

Examples of built-in categories and the resources they map to include:
* `api-extensions`
  * `apiregistration.k8s.io/v* apiservices`
  * `admissionregistration.k8s.io/v* mutatingwebhookconfigurations`
  * `admissionregistration.k8s.io/v* validatingwebhookconfigurations`
  * `admissionregistration.k8s.io/v* validatingadmissionpolicies`
  * `admissionregistration.k8s.io/v* validatingadmissionpolicybindings`
  * `apiextensions.k8s.io/v* customresourcedefinitions`
* `all`
  * `v1 pods`
  * `v1 replicationcontrollers`
  * `v1 services`
  * `apps/v* daemonsets`
  * `apps/v* deployments`
  * `apps/v* replicasets`
  * `apps/v* statefulsets`
  * `autoscaling/v* horizontalpodautoscalers`
  * `batch/v* cronjobs`
  * `batch/v* jobs`

With the above categories, and only built-in API types served, `kubectl get all` would be equivalent to 
`kubectl get pods.v1.,replicationcontrollers.v1.,services.v1.,daemonsets.v1.apps,deployments.v1.apps,replicasets.v1.apps,statefulsets.v1.apps,horizontalpodautoscalers.v2.autoscaling,cronjobs.v1.batch,jobs.v1.batch,`.

## Idempotency

All compatible Kubernetes APIs MUST support "name idempotency" and respond with
an HTTP status code 409 when a request is made to POST an object that has the
same name as an existing object in the system. See
[the identifiers docs](https://kubernetes.io/docs/concepts/overview/working-with-objects/names/)
for details.

Names generated by the system may be requested using `metadata.generateName`.
GenerateName indicates that the name should be made unique by the server prior
to persisting it. A non-empty value for the field indicates the server should
attempt to make the name unique (and the name returned to the client will be
different than the name passed). The value of this field will be combined with a
random suffix on the server if the Name field has not been provided. The
provided value must be valid within the rules for Name, and may be truncated by
the length of the suffix. If this field is specified, and Name is not present,
the server will return a 409 with Reason `AlreadyExists` if the generated name
exists, and the client should retry (after waiting at least the amount of time
indicated in the Retry-After header, if it is present).

## Optional vs. Required

Fields must be either optional or required.

A field that is required means that the writer must express an opinion about the value of the field.
Required fields are always present, and readers can rely on the field being present in the object.

An optional field is one where the writer may choose to omit the field entirely.
Readers should not assume that the field is present, unless the field also has a server-side default value.

Default values can only be set on optional fields.

Optional fields have the following properties:

- They have the `+optional` comment tag in Go.
- They are a pointer type in the Go definition (e.g. `AwesomeFlag *SomeFlag`) or
have a built-in `nil` value (e.g. maps and slices).
- They are marked with the `omitempty` json struct tag in the Go definition.
- The API server should allow POSTing and PUTing a resource with this field
unset.

When the field type has a built-in `nil` value, such as a map or a slice, and
your use case means that you need to be able to distinguish between
"field not set" and "field set to an empty value", you should use a pointer to
the type, even though it has a built-in `nil` value.
See https://github.com/kubernetes/kubernetes/issues/34641.

Note that for backward compatibility, any field that has the `omitempty` struct
tag, and is not explicitly marked as `+required`, will be considered to be optional.
This is expected to change in the future, and new fields should explicitly set either
an `+optional` or `+required` comment tag.

Required fields have the following properties:

- They do not have an `+optional` comment tag.
- They mark themselves as required explicitly with a `+required` comment tag.
- The API server should not allow POSTing or PUTing a resource with this field
unset.
- They _typically_ do not use pointer types in the Go definition (e.g. `AnotherFlag SomeFlag`), though required fields where the zero value is a valid value must use pointer types, paired with an `omitempty` struct tag to avoid spurious null serializations.

For more details on how to use pointers and `omitempty` with fields, see [Serialization of optional/required fields](#serialization-of-optionalrequired-fields).

Using the `+optional` or the `omitempty` tag causes OpenAPI documentation to
reflect that the field is optional.

### Serialization of optional/required fields

Using a pointer allows distinguishing unset from the zero value for that type.
There are some cases where, in principle, a pointer is not needed for a
field since the zero value is forbidden, and thus implies unset.
There are examples of this in the codebase. However:

- It can be difficult for implementors to anticipate all cases where an empty
value might need to be distinguished from a zero value.
- Structs are not omitted from encoder output even where `omitempty` is specified,
which is messy.

To determine whether a field should be a pointer, consider the following:

```mermaid
graph TD;
    A[Start] --> B{Is the zero value a valid user choice?};
    B -- Yes --> C[Use a pointer and omitempty];
    B -- No --> D{Is the field optional or required?};
    D -- Optional --> E["Does the field type have a built-in nil value (map or slice)?"];
    E -- Yes --> F[Do not use a pointer, use omitempty];
    E -- No --> C;
    D -- Required --> F
```
There are several implications of the above:
- For lists and maps where the zero valid is a valid user choice, this means that `[]` and `{}` have a semantically different meaning than unset, in this case it is appropriate to use `*[]T` or `*map[T]S` respectively.
- For `bool` types, the zero value is `false`, which is always a valid user choice. `bool` types should always be pointers, and should always use the `omitempty` tag.
- When a field is required, and the zero value is not valid, a structured client who has not expressed an explicit choice will have their request rejected by the API server based on the invalid value, rather than the field being unset.
  - For example, a string with a minimum length of 1; Validation would not understand if the field was unset, or set to the empty string deliberately, but would still reject the request because it did not meet the length requirements.
  - Technically, using a pointer in these cases is also acceptable, but not advised as it makes coding more complex, and increases the risk of nil pointer exceptions.
  - In these cases, not using `omitempty` provides the same result, but pollutes the marshaled object with zero values and is not recommended.
- For structs, the zero value can only be valid when the struct has no required fields, and does not require at least one property to be set.
  - Required structs should use `omitzero` to avoid marshalling the zero value.

#### Serialization of custom resources

When custom resources are admitted by the API server, openapi validation is applied to the object _prior_ to any structured client observing the object.

For a field where the zero value is not valid, the openapi validation will reject the object if the field is present and set to the zero value,
before a controller or validation webhook could observe the field.

This means that there is no need to distinguish in a custom resource, between unset and the zero value for fields where the zero value is not valid.
In these cases, pointers are not needed, as the zero value indicates to the structured client that the field is unset.

This can be beneficial to API authors, as it reduces the complexity of the API, and reduces the risk of nil pointer exceptions in controllers.

## Defaulting

In general we want default values to be explicitly represented in our APIs,
rather than asserting that "unspecified fields get the default behavior".  This
is important so that:
 - default values can evolve and change in newer API versions
 - the stored configuration depicts the full desired state, making it easier
   for the system to determine how to achieve the state, and for the user to
   know what to anticipate

There are 3 distinct ways that default values can be applied when creating or
updating (including patch and apply) a resource:

 1. static: based on the requested API version and possibly other fields in the
    resource, fields can be assigned values during the API call
 2. admission control: based on the configured admission controllers and
    possibly other state in or out of the cluster, fields can be assigned
    values during the API call
 3. controllers: arbitrary changes (within the bounds of what is allowed) can
    be made to a resource after the API call has completed

Some care is required when deciding which mechanism to use and managing the
semantics.

### Static Defaults

Static default values are specific to each API version.  The default field
values applied when creating an object with the "v1" API may be different than
the values applied when using the "v2" API.  In most cases, these values are
defined as literal values by the API version (e.g. "if this field is not
specified it defaults to 0").

In some cases, these values may be conditional on or deterministically derived
from other fields (e.g. "if otherField is X then this field defaults to 0" or
"this field defaults to the value of otherField").  Note that such derived
defaults present a hazard in the face of updates - if the "other" field
changes, the derived field may have to change, too.  The static defaulting
logic is unaware of updates and has no concept of "previous value", which means
this inter-field relationship becomes the user's problem - they must update
both the field they care about and the "other" field.

In very rare cases, these values may be allocated from some pool or determined
by some other method (e.g. Service's IP and IP-family related fields need to
consider other configuration settings).

These values are applied synchronously by the API server when decoding
versioned data.  For CREATE and UPDATE operations this is fairly
straight-forward - when the API server receives a (versioned) request, the
default values are immediately applied before any further processing.  When the
API call completes, all static defaults will have been set and stored.
Subsequent GETs of the resource will include the default values explicitly.
However, static defaults also apply when an object is read from storage (i.e.
GET operations).  This means that when someone GETs an "older" stored object,
any fields which have been added to the API since that object was stored will
be defaulted and returned according to the API version that is stored.

Static defaults are the best choice for values which are logically required,
but which have a value that works well for most users.  Static defaulting
must not consider any state except the object being operated upon (and the
complexity of Service API stands as an example of why).

Default values can be specified on a field using the `+default=` tag. Primitives
will have their values directly assigned while structs will go through the
JSON unmarshalling process. Fields that do not have an `omitempty` json tag will
default to the zero value of their corresponding type if no default is assigned.

Refer to [defaulting docs](https://kubernetes.io/docs/tasks/extend-kubernetes/custom-resources/custom-resource-definitions/#defaulting)
for more information.

### Admission Controlled Defaults

In some cases, it is useful to set a default value which is not derived from
the object in question.  For example, when creating a PersistentVolumeClaim,
the storage class must be specified.  For many users, the best answer is
"whatever the cluster admin has decided for the default".  StorageClass is a
different API than PersistentVolumeClaim, and which one is denoted as the
default may change at any time.  Thus this is not eligible for static
defaulting.

Instead, we can provide a built-in admission controller or a
MutatingWebhookConfiguration.  Unlike static defaults, these may consider
external state (such as annotations on StorageClass objects) when deciding
default values, and must handle things like race conditions (e.g. a
StorageClass is designated the default, but the admission controller has not
yet seen that update).  These admission controllers are strictly optional and
can be disabled.  As such, fields which are initialized this way must be
strictly optional.

Like static defaults, these are run synchronously to the API operation in
question, and when the API call completes, all static defaults will have been
set.  Subsequent GETs of the resource will include the default values
explicitly.

### Controller-Assigned Defaults (aka Late Initialization)

Late initialization is when resource fields are set by a system controller
after an object is created/updated (asynchronously).  For example, the
scheduler sets the `pod.spec.nodeName` field after the pod is created.  It's
a stretch to call this "defaulting" but since it is so common and useful, it is
included here.

Like admission controlled defaults, these controllers may consider external
state when deciding what values to set, must handle race conditions, and can be
disabled.  Fields which are initialized this way must be strictly optional
(meaning observers will see the object without these fields set, and that is
allowable and semantically correct).

Like all controllers, care must be taken to not clobber unrelated fields or
values (e.g. in an array).  Using one of the patch or apply mechanisms is
recommended to facilitate composition and concurrency of controllers.

### What May Be Defaulted

All forms of defaulting should only make the following types of modifications:
 - Setting previously unset fields
 - Adding keys to maps
 - Adding values to arrays which have mergeable semantics
   (`+listType=map` tag or `patchStrategy:"merge"` attribute in the type definition)

In particular we never want to change or override a value that was provided by
the user.  If they requested something invalid, they should get an error.

These rules ensure that:
 1. a user (with sufficient privilege) can override any system-default
    behaviors by explicitly setting the fields that would otherwise have been
    defaulted
 1. updates from users can be merged with default values

### Considerations For PUT Operations

Once an object has been created and defaults have been applied, it's very
common for updates to happen over time.  Kubernetes offers several ways of
updating an object which preserve existing values in fields other than those
being updated (e.g. strategic merge patch and server-side apply).  There is,
however, a less obvious way of updating objects which can have bad interactions
with default values - PUT (aka `kubectl replace`).

The goal is that, for a given input (e.g. YAML file), PUT on an existing object
should produce the same result as if you used that input to create the object.
Calling PUT a second time with the same input should be idempotent and should
not change the resource.  Even a read-modify-write cycle is not a perfect
solution in the face of version skew.

When an object is updated with a PUT, the API server will see the "old" object
with previously assigned defaults and the "new" object with newly assigned
defaults.  For static defaults this can be a problem if the CREATE and the PUT
used different API versions.  For example, "v1" of an API might default a field
to `false`, while "v2" defaults it to `true`.  If an object was created via API
v1 (field = `false`) and then replaced via API v2, the field will attempt to
change to `true`.  This can also be a problem when the values are allocated or
derived from a source outside of the object in question (e.g. Service IPs).

For some APIs this is acceptable and actionable.  For others, this may be
disallowed by validation.  In the latter case, the user will get an error about
an attempt to change a field which is not even present in their YAML.  This is
especially dangerous when adding new fields - an older client may not even know
about the existence of the field, making even a read-modify-write cycle fail.
While it is "correct" (in the sense that it is really what they asked for with
PUT), it is not helpful and is a bad UX.

When adding a field with a static or admission controlled default, this must be
considered.  If the field is immutable after creation, consider adding logic to
manually "patch" the value from the "old" object into the "new" one when it has
been "unset", rather than returning an error or allocating a different value
(e.g.  Service IPs).  This will very often be what the user meant, even if it
is not what they said.  This may require setting the default in a different way
(e.g.  in the registry code which understands updates instead of in the
versioned defaulting code which does not).  Be careful to detect and report
legitimate errors where the "new" value is specified but is different from the
"old" value.

For controller-defaulted fields, the situation is even more unpleasant.
Controllers do not have an opportunity to "patch" the value before the API
operation is committed.  If the "unset" value is allowed then it will be saved,
and any watch clients will be notified.  If the "unset" value is not allowed or
mutations are otherwise disallowed, the user will get an error, and there's
simply nothing we can do about it.

## Concurrency Control and Consistency

Kubernetes leverages the concept of *resource versions* to achieve optimistic
concurrency. All Kubernetes resources have a "resourceVersion" field as part of
their metadata. This resourceVersion is a string that identifies the internal
version of an object that can be used by clients to determine when objects have
changed. When a record is about to be updated, its version is checked against a
pre-saved value, and if it doesn't match, the update fails with a StatusConflict
(HTTP status code 409).

The resourceVersion is changed by the server every time an object is modified.
If resourceVersion is included with the PUT operation the system will verify
that there have not been other successful mutations to the resource during a
read/modify/write cycle, by verifying that the current value of resourceVersion
matches the specified value.

The resourceVersion is currently backed by [etcd's
mod_revision](https://etcd.io/docs/latest/learning/api/#key-value-pair).
However, it's important to note that the application should *not* rely on the
implementation details of the versioning system maintained by Kubernetes. We may
change the implementation of resourceVersion in the future, such as to change it
to a timestamp or per-object counter.

The only way for a client to know the expected value of resourceVersion is to
have received it from the server in response to a prior operation, typically a
GET. This value MUST be treated as opaque by clients and passed unmodified back
to the server. Clients should not assume that the resource version has meaning
across namespaces, different kinds of resources, or different servers.
Currently, the value of resourceVersion is set to match etcd's sequencer. You
could think of it as a logical clock the API server can use to order requests.
However, we expect the implementation of resourceVersion to change in the
future, such as in the case we shard the state by kind and/or namespace, or port
to another storage system.

In the case of a conflict, the correct client action at this point is to GET the
resource again, apply the changes afresh, and try submitting again. This
mechanism can be used to prevent races like the following:

```
Client #1                                  Client #2
GET Foo                                    GET Foo
Set Foo.Bar = "one"                        Set Foo.Baz = "two"
PUT Foo                                    PUT Foo
```

When these sequences occur in parallel, either the change to Foo.Bar or the
change to Foo.Baz can be lost.

On the other hand, when specifying the resourceVersion, one of the PUTs will
fail, since whichever write succeeds changes the resourceVersion for Foo.

resourceVersion may be used as a precondition for other operations (e.g., GET,
DELETE) in the future, such as for read-after-write consistency in the presence
of caching.

"Watch" operations specify resourceVersion using a query parameter. It is used
to specify the point at which to begin watching the specified resources. This
may be used to ensure that no mutations are missed between a GET of a resource
(or list of resources) and a subsequent Watch, even if the current version of
the resource is more recent. This is currently the main reason that list
operations (GET on a collection) return resourceVersion.


## Serialization Format

APIs may return alternative representations of any resource in response to an
Accept header or under alternative endpoints, but the default serialization for
input and output of API responses MUST be JSON.

A protobuf encoding is also accepted for built-in resources. As proto is not
self-describing, there is an envelope wrapper which describes the type of
the contents.

All dates should be serialized as RFC3339 strings.

## Units

Units must either be explicit in the field name (e.g., `timeoutSeconds`), or
must be specified as part of the value (e.g., `resource.Quantity`). Which
approach is preferred is TBD, though currently we use the `fooSeconds`
convention for durations.

Duration fields must be represented as integer fields with units being
part of the field name (e.g. `leaseDurationSeconds`). We don't use Duration
in the API since that would require clients to implement go-compatible parsing.

## Selecting Fields

Some APIs may need to identify which field in a JSON object is invalid, or to
reference a value to extract from a separate resource. The current
recommendation is to use standard JavaScript syntax for accessing that field,
assuming the JSON object was transformed into a JavaScript object, without the
leading dot, such as `metadata.name`.

Examples:

* Find the field "current" in the object "state" in the second item in the array
"fields": `fields[1].state.current`

## Object references

Object references on a namespaced type should usually refer only to objects in
the same namespace.  Because namespaces are a security boundary, cross namespace
references can have unexpected impacts, including:
 1. leaking information about one namespace into another namespace. It's natural to place status messages or even bits of
    content about the referenced object in the original. This is a problem across namespaces.
 2. potential invasions into other namespaces. Often references give access to a piece of referred information, so being
    able to express "give me that one over there" is dangerous across namespaces without additional work for permission checks
    or opt-in's from both involved namespaces.
 3. referential integrity problems that one party cannot solve. Referencing namespace/B from namespace/A doesn't imply the
    power to control the other namespace. This means that you can refer to a thing you cannot create or update.
 4. unclear semantics on deletion. If a namespaced resource  is referenced by other namespaces, should a delete of the
    referenced resource result in removal or should the referenced resource be force to remain.
 5. unclear semantics on creation. If a referenced resource is created after its reference, there is no way to know if it
    is the one that is expected or if it is a different one created with the same name.

Built-in types and ownerReferences do not support cross namespaces references.
If a non-built-in types chooses to have cross-namespace references the semantics of the edge cases above should be
clearly described and the permissions issues should be resolved.
This could be done with a double opt-in (an opt-in from both the referrer and the refer-ee) or with secondary permissions
checks performed in admission.

### Naming of the reference field

The name of the reference field should be of the format "{field}Ref", with "Ref" always included in the suffix.

The "{field}" component should be named to indicate the purpose of the reference. For example, "targetRef" in an
endpoint indicates that the object reference specifies the target.

It is okay to have the "{field}" component indicate the resource type. For example, "secretRef" when referencing
a secret. However, this comes with the risk of the field being a misnomer in the case that the field is expanded to
reference more than one type.

In the case of a list of object references, the field should be of the format "{field}Refs", with the same guidance
as the singular case above.

### Referencing resources with multiple versions

Most resources will have multiple versions. For example, core resources
will undergo version changes as it transitions from alpha to GA.

Controllers should assume that a version of a resource may change, and include appropriate error handling.

### Handling of resources that do not exist

There are multiple scenarios where a desired resource may not exist. Examples include:

- the desired version of the resource does not exist.
- race condition in the bootstrapping of a cluster resulting a resource not yet added.
- user error.

Controllers should be authored with the assumption that the referenced resource may not exist, and include
error handling to make the issue clear to the user.

### Validation of fields

Many of the values used in an object reference are used as part of the API path. For example,
the object name is used in the path to identify the object. Unsanitized, these values can be used to
attempt to retrieve other resources, such as by using values with semantic meanings such as  `..` or `/`.

Have the controller validate fields before using them as path segments in an API request, and emit an event to
tell the user that the validation has failed.

See [Object Names and IDs](https://kubernetes.io/docs/concepts/overview/working-with-objects/names/)
for more information on legal object names.

### Do not modify the referred object

To minimize potential privilege escalation vectors, do not modify the object that is being referred to,
or limit modification to objects in the same namespace and constrain the type of modification allowed
(for example, the HorizontalPodAutoscaler controller only writes to the `/scale` subresource).

### Minimize copying or printing values to the referrer object

As the permissions of the controller can differ from the permissions of the author of the object
the controller is managing, it is possible that the author of the object may not have permissions to
view the referred object. As a result, the copying of any values about the referred object to the
referrer object can be considered permissions escalations, enabling a user to read values that they
would not have access to previously.

The same scenario applies to writing information about the referred object to events.

In general, do not write or print information retrieved from the referred object to the spec, other objects, or logs.

When it is necessary, consider whether these values would be ones that the
author of the referrer object would have access to via other means (e.g. already required to
correctly populate the object reference).

### Object References Examples

The following sections illustrate recommended schemas for various object references scenarios.

The schemas outlined below are designed to enable purely additive fields as the types of referencable
objects expand, and therefore are backwards compatible.

For example, it is possible to go from a single resource type to multiple resource types without
a breaking change in the schema.

#### Single resource reference

A single kind object reference is straightforward in that the controller can hard-code most qualifiers needed to identify the object. As such as the only value needed to be provided is the name (and namespace, although cross-namespace references are discouraged):

```yaml
# for a single resource, the suffix should be Ref, with the field name
# providing an indication as to the resource type referenced.
secretRef:
    name: foo
    # namespace would generally not be needed and is discouraged,
    # as explained above.
    namespace: foo-namespace
```

This schema should only be used when the intention is to always have the reference only be to a single resource.
If extending to multiple resource types is possible, use the [multiple resource reference](#multiple-resource-reference).

##### Controller behavior

The operator is expected to know the version, group, and resource name of the object it needs to retrieve the value from, and can use the discovery client or construct the API path directly.

#### Multiple resource reference

Multi-kind object references are used when there is a bounded set of valid resource types that a reference can point to.

As with a single-kind object reference, the operator can supply missing fields, provided that the fields that are present are sufficient to uniquely identify the object resource type among the set of supported types.

```yaml
# guidance for the field name is the same as a single resource.
fooRef:
    group: sns.services.k8s.aws
    resource: topics
    name: foo
    namespace: foo-namespace
```

Although not always necessary to help a controller identify a resource type, “group” is included to avoid ambiguity when the resource exists in multiple groups. It also provides clarity to end users and enables copy-pasting of a reference without the referenced type changing due to a different controller handling the reference.

##### Kind vs. Resource

A common point of confusion in object references is whether to construct
references with a "kind" or "resource" field. Historically most object
references in Kubernetes have used "kind". This is not as precise as "resource".
Although each combination of "group" and "resource" must be unique within
Kubernetes, the same is not always true for "group" and "kind". It is possible
for multiple resources to make use of the same "kind".

Typically all objects in Kubernetes have a canonical primary resource - such as
“pods” representing the way to create and delete resources of the “Pod” schema.
While it is possible a resource schema cannot be directly created, such as a
“Scale” object which is only used within the “scale” subresource of a number of
workloads, most object references address the primary resource via its schema.
In the context of object references, "kind" refers to the schema, not the
resource.

If implementations of an object reference will always have a clear way to map
kinds to resources, it is acceptable to use "kind" in the object reference. In
general, this requires implementations to have a predefined mapping between
kinds and resources (this is the case for built-in references which use "kind").
Relying on dynamic kind to resource mapping is not safe. Even if a "kind" only
dynamically maps to a single resource initially, it's possible for another
resource to be mounted that refers to the same "kind", potentially breaking any
dynamic resource mapping.

If an object reference may be used to reference resources of arbitrary types and
the mapping between kind and resource could be ambiguous, "resource" should be
used in the object reference.

The Ingress API provides a good example of where "kind" is acceptable for an
object reference. The API supports a backend reference as an extension point.
Implementations can use this to support forwarding traffic to custom targets
such as a storage bucket. Importantly, the supported target types are clearly
defined by each implementation of the API and there is no ambiguity for which
resource a kind maps to. This is because each Ingress implementation has a
hard-coded mapping of kind to resource.

The object reference above would look like this if it were using "kind" instead
of "resource":

```yaml
fooRef:
    group: sns.services.k8s.aws
    kind: Topic
    name: foo
    namespace: foo-namespace
```

##### Controller behavior

The operator can store a map of (group,resource) to the version of that resource it desires. From there, it can construct the full path to the resource, and retrieve the object.

It is also possible to have the controller choose a version that it finds via the discovery client. However, as schemas can vary across different versions
of a resource, the controller must also handle these differences.

#### Generic object reference

A generic object reference is used when the desire is to provide a pointer to some object to simplify discovery for the user. For example, this could be used to reference a target object for a `core.v1.Event` that occurred.

With a generic object reference, it is not possible to extract any information about the referenced object aside from what is standard (e.g. ObjectMeta). Since any standard fields exist in any version of a resource, it is possible to not include version in this case:

```yaml
fooObjectRef:
    group: operator.openshift.io
    resource: openshiftapiservers
    name: cluster
    # namespace is unset if the resource is cluster-scoped, or lives in the
    # same namespace as the referrer.
```

##### Controller behavior

The operator would be expected to find the resource via the discovery client (as the version is not supplied). As any retrievable field would be common to all objects, any version of the resource should do.

#### Field reference

A field reference is used when the desire is to extract a value from a specific field in a referenced object.

Field references differ from other reference types, as the operator has no knowledge of the object prior to the reference. Since the schema of an object can differ for different versions of a resource, this means that a “version” is required for this type of reference.

```yaml
fooFieldRef:
   version: v1 # version of the resource
   # group is elided in the ConfigMap example, since it has a blank group in the OpenAPI spec.
   resource: configmaps
   fieldPath: data.foo
```

The fieldPath should point to a single value, and use [the recommended field selector notation](#selecting-fields) to denote the field path.

##### Controller behavior

In this scenario, the user will supply all of the required path elements: group, version, resource, name, and possibly namespace.
As such, the controller can construct the API prefix and query it without the use of the discovery client:

```
/apis/{group}/{version}/{resource}/
```

## HTTP Status codes

The server will respond with HTTP status codes that match the HTTP spec. See the
section below for a breakdown of the types of status codes the server will send.

The following HTTP status codes may be returned by the API.

#### Success codes

* `200 StatusOK`
  * Indicates that the request completed successfully.
* `201 StatusCreated`
  * Indicates that the request to create kind completed successfully.
* `204 StatusNoContent`
  * Indicates that the request completed successfully, and the response contains
no body.
  * Returned in response to HTTP OPTIONS requests.

#### Error codes

* `307 StatusTemporaryRedirect`
  * Indicates that the address for the requested resource has changed.
  * Suggested client recovery behavior:
    * Follow the redirect.


* `400 StatusBadRequest`
  * Indicates the requested is invalid.
  * Suggested client recovery behavior:
    * Do not retry. Fix the request.


* `401 StatusUnauthorized`
  * Indicates that the server can be reached and understood the request, but
refuses to take any further action, because the client must provide
authorization. If the client has provided authorization, the server is
indicating the provided authorization is unsuitable or invalid.
  * Suggested client recovery behavior:
    * If the user has not supplied authorization information, prompt them for
the appropriate credentials. If the user has supplied authorization information,
inform them their credentials were rejected and optionally prompt them again.


* `403 StatusForbidden`
  * Indicates that the server can be reached and understood the request, but
refuses to take any further action, because it is configured to deny access for
some reason to the requested resource by the client.
  * Suggested client recovery behavior:
    * Do not retry. Fix the request.


* `404 StatusNotFound`
  * Indicates that the requested resource does not exist.
  * Suggested client recovery behavior:
    * Do not retry. Fix the request.


* `405 StatusMethodNotAllowed`
  * Indicates that the action the client attempted to perform on the resource
was not supported by the code.
  * Suggested client recovery behavior:
    * Do not retry. Fix the request.


* `409 StatusConflict`
  * Indicates that either the resource the client attempted to create already
exists or the requested update operation cannot be completed due to a conflict.
  * Suggested client recovery behavior:
    * If creating a new resource:
      * Either change the identifier and try again, or GET and compare the
fields in the pre-existing object and issue a PUT/update to modify the existing
object.
    * If updating an existing resource:
      * See `Conflict` from the `status` response section below on how to
retrieve more information about the nature of the conflict.
      * GET and compare the fields in the pre-existing object, merge changes (if
still valid according to preconditions), and retry with the updated request
(including `ResourceVersion`).


* `410 StatusGone`
  * Indicates that the item is no longer available at the server and no
forwarding address is known.
  * Suggested client recovery behavior:
    * Do not retry. Fix the request.


* `422 StatusUnprocessableEntity`
  * Indicates that the requested create or update operation cannot be completed
due to invalid data provided as part of the request.
  * Suggested client recovery behavior:
    * Do not retry. Fix the request.


* `429 StatusTooManyRequests`
  * Indicates that either the client rate limit has been exceeded or the
server has received more requests than it can process.
  * Suggested client recovery behavior:
    * Read the `Retry-After` HTTP header from the response, and wait at least
that long before retrying.


* `500 StatusInternalServerError`
  * Indicates that the server can be reached and understood the request, but
either an unexpected internal error occurred and the outcome of the call is
unknown, or the server cannot complete the action in a reasonable time (this may
be due to temporary server load or a transient communication issue with another
server).
  * Suggested client recovery behavior:
    * Retry with exponential backoff.


* `503 StatusServiceUnavailable`
  * Indicates that required service is unavailable.
  * Suggested client recovery behavior:
    * Retry with exponential backoff.


* `504 StatusServerTimeout`
  * Indicates that the request could not be completed within the given time.
Clients can get this response ONLY when they specified a timeout param in the
request.
  * Suggested client recovery behavior:
    * Increase the value of the timeout param and retry with exponential
backoff.

## Response Status Kind

Kubernetes will always return the `Status` kind from any API endpoint when an
error occurs. Clients SHOULD handle these types of objects when appropriate.

A `Status` kind will be returned by the API in two cases:
  * When an operation is not successful (i.e. when the server would return a non
2xx HTTP status code).
  * When a HTTP `DELETE` call is successful.

The status object is encoded as JSON and provided as the body of the response.
The status object contains fields for humans and machine consumers of the API to
get more detailed information for the cause of the failure. The information in
the status object supplements, but does not override, the HTTP status code's
meaning. When fields in the status object have the same meaning as generally
defined HTTP headers and that header is returned with the response, the header
should be considered as having higher priority.

**Example:**

```console
$ curl -v -k -H "Authorization: Bearer WhCDvq4VPpYhrcfmF6ei7V9qlbqTubUc" https://10.240.122.184:443/api/v1/namespaces/default/pods/grafana

> GET /api/v1/namespaces/default/pods/grafana HTTP/1.1
> User-Agent: curl/7.26.0
> Host: 10.240.122.184
> Accept: */*
> Authorization: Bearer WhCDvq4VPpYhrcfmF6ei7V9qlbqTubUc
>

< HTTP/1.1 404 Not Found
< Content-Type: application/json
< Date: Wed, 20 May 2015 18:10:42 GMT
< Content-Length: 232
<
{
  "kind": "Status",
  "apiVersion": "v1",
  "metadata": {},
  "status": "Failure",
  "message": "pods \"grafana\" not found",
  "reason": "NotFound",
  "details": {
    "name": "grafana",
    "kind": "pods"
  },
  "code": 404
}
```

`status` field contains one of two possible values:
* `Success`
* `Failure`

`message` may contain human-readable description of the error

`reason` may contain a machine-readable, one-word, CamelCase description of why
this operation is in the `Failure` status. If this value is empty there is no
information available. The `reason` clarifies an HTTP status code but does not
override it.

`details` may contain extended data associated with the reason. Each reason may
define its own extended details. This field is optional and the data returned is
not guaranteed to conform to any schema except that defined by the reason type.

Possible values for the `reason` and `details` fields:
* `BadRequest`
  * Indicates that the request itself was invalid, because the request doesn't
make any sense, for example deleting a read-only object.
  * This is different than `status reason` `Invalid` above which indicates that
the API call could possibly succeed, but the data was invalid.
  * API calls that return BadRequest can never succeed.
  * Http status code: `400 StatusBadRequest`


* `Unauthorized`
  * Indicates that the server can be reached and understood the request, but
refuses to take any further action without the client providing appropriate
authorization. If the client has provided authorization, this error indicates
the provided credentials are insufficient or invalid.
  * Details (optional):
    * `kind string`
      * The kind attribute of the unauthorized resource (on some operations may
differ from the requested resource).
    * `name string`
      * The identifier of the unauthorized resource.
   * HTTP status code: `401 StatusUnauthorized`


* `Forbidden`
  * Indicates that the server can be reached and understood the request, but
refuses to take any further action, because it is configured to deny access for
some reason to the requested resource by the client.
  * Details (optional):
    * `kind string`
      * The kind attribute of the forbidden resource (on some operations may
differ from the requested resource).
    * `name string`
      * The identifier of the forbidden resource.
  * HTTP status code: `403 StatusForbidden`


* `NotFound`
  * Indicates that one or more resources required for this operation could not
be found.
  * Details (optional):
    * `kind string`
      * The kind attribute of the missing resource (on some operations may
differ from the requested resource).
    * `name string`
      * The identifier of the missing resource.
  * HTTP status code: `404 StatusNotFound`


* `AlreadyExists`
  * Indicates that the resource you are creating already exists.
  * Details (optional):
    * `kind string`
      * The kind attribute of the conflicting resource.
    * `name string`
      * The identifier of the conflicting resource.
  * HTTP status code: `409 StatusConflict`

* `Conflict`
  * Indicates that the requested update operation cannot be completed due to a
conflict. The client may need to alter the request. Each resource may define
custom details that indicate the nature of the conflict.
  * HTTP status code: `409 StatusConflict`


* `Invalid`
  * Indicates that the requested create or update operation cannot be completed
due to invalid data provided as part of the request.
  * Details (optional):
    * `kind string`
      * the kind attribute of the invalid resource
    * `name string`
      * the identifier of the invalid resource
    * `causes`
      * One or more `StatusCause` entries indicating the data in the provided
resource that was invalid. The `reason`, `message`, and `field` attributes will
be set.
  * HTTP status code: `422 StatusUnprocessableEntity`


* `Timeout`
  * Indicates that the request could not be completed within the given time.
Clients may receive this response if the server has decided to rate limit the
client, or if the server is overloaded and cannot process the request at this
time.
  * Http status code: `429 TooManyRequests`
  * The server should set the `Retry-After` HTTP header and return
`retryAfterSeconds` in the details field of the object. A value of `0` is the
default.


* `ServerTimeout`
  * Indicates that the server can be reached and understood the request, but
cannot complete the action in a reasonable time. This maybe due to temporary
server load or a transient communication issue with another server.
    * Details (optional):
      * `kind string`
        * The kind attribute of the resource being acted on.
      * `name string`
        * The operation that is being attempted.
  * The server should set the `Retry-After` HTTP header and return
`retryAfterSeconds` in the details field of the object. A value of `0` is the
default.
  * Http status code: `504 StatusServerTimeout`


* `MethodNotAllowed`
  * Indicates that the action the client attempted to perform on the resource
was not supported by the code.
  * For instance, attempting to delete a resource that can only be created.
  * API calls that return MethodNotAllowed can never succeed.
  * Http status code: `405 StatusMethodNotAllowed`


* `InternalError`
  * Indicates that an internal error occurred, it is unexpected and the outcome
of the call is unknown.
  * Details (optional):
    * `causes`
      * The original error.
  * Http status code: `500 StatusInternalServerError` `code` may contain the suggested HTTP return code for this status.


## Events

Events are complementary to status information, since they can provide some
historical information about status and occurrences in addition to current or
previous status. Generate events for situations users or administrators should
be alerted about.

Choose a unique, specific, short, CamelCase reason for each event category. For
example, `FreeDiskSpaceInvalid` is a good event reason because it is likely to
refer to just one situation, but `Started` is not a good reason because it
doesn't sufficiently indicate what started, even when combined with other event
fields.

`Error creating foo` or `Error creating foo %s` would be appropriate for an
event message, with the latter being preferable, since it is more informational.

Accumulate repeated events in the client, especially for frequent events, to
reduce data volume, load on the system, and noise exposed to users.

## Naming conventions

* Go field names must be PascalCase. JSON field names must be camelCase. Other
than capitalization of the initial letter, the two should almost always match.
No underscores or dashes in either.
* Field and resource names should be declarative, not imperative (SomethingDoer,
DoneBy, DoneAt).
* Use `Node` where referring to
the node resource in the context of the cluster. Use `Host` where referring to
properties of the individual physical/virtual system, such as `hostname`,
`hostPath`, `hostNetwork`, etc.
* `FooController` is a deprecated kind naming convention. Name the kind after
the thing being controlled instead (e.g., `Job` rather than `JobController`).
* The name of a field that specifies the time at which `something` occurs should
be called `somethingTime`. Do not use `stamp` (e.g., `creationTimestamp`).
* We use the `fooSeconds` convention for durations, as discussed in the [units
subsection](#units).
  * `fooPeriodSeconds` is preferred for periodic intervals and other waiting
periods (e.g., over `fooIntervalSeconds`).
  * `fooTimeoutSeconds` is preferred for inactivity/unresponsiveness deadlines.
  * `fooDeadlineSeconds` is preferred for activity completion deadlines.
* Do not use abbreviations in the API, except where they are extremely commonly
used, such as "id", "args", or "stdin".
* Acronyms should similarly only be used when extremely commonly known. All
letters in the acronym should have the same case, using the appropriate case for
the situation. For example, at the beginning of a field name, the acronym should
be all lowercase, such as "httpGet". Where used as a constant, all letters
should be uppercase, such as "TCP" or "UDP".
* The name of a field referring to another resource of kind `Foo` by name should
be called `fooName`. The name of a field referring to another resource of kind
`Foo` by ObjectReference (or subset thereof) should be called `fooRef`.
* More generally, include the units and/or type in the field name if they could
be ambiguous and they are not specified by the value or value type.
* The name of a field expressing a boolean property called 'fooable' should be
called `Fooable`, not `IsFooable`.

### Namespace Names
* The name of a namespace must be a
[DNS_LABEL](https://git.k8s.io/design-proposals-archive/architecture/identifiers.md).
* The `kube-` prefix is reserved for Kubernetes system namespaces, e.g. `kube-system` and `kube-public`.
* See
[the namespace docs](https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/)
for more information.

## Label, selector, and annotation conventions

Labels are the domain of users. They are intended to facilitate organization and
management of API resources using attributes that are meaningful to users, as
opposed to meaningful to the system. Think of them as user-created mp3 or email
inbox labels, as opposed to the directory structure used by a program to store
its data. The former enables the user to apply an arbitrary ontology, whereas
the latter is implementation-centric and inflexible. Users will use labels to
select resources to operate on, display label values in CLI/UI columns, etc.
Users should always retain full power and flexibility over the label schemas
they apply to labels in their namespaces.

However, we should support conveniences for common cases by default. For
example, what we now do in ReplicationController is automatically set the RC's
selector and labels to the labels in the pod template by default, if they are
not already set. That ensures that the selector will match the template, and
that the RC can be managed using the same labels as the pods it creates. Note
that once we generalize selectors, it won't necessarily be possible to
unambiguously generate labels that match an arbitrary selector.

If the user wants to apply additional labels to the pods that it doesn't select
upon, such as to facilitate adoption of pods or in the expectation that some
label values will change, they can set the selector to a subset of the pod
labels. Similarly, the RC's labels could be initialized to a subset of the pod
template's labels, or could include additional/different labels.

For disciplined users managing resources within their own namespaces, it's not
that hard to consistently apply schemas that ensure uniqueness. One just needs
to ensure that at least one value of some label key in common differs compared
to all other comparable resources. We could/should provide a verification tool
to check that. However, development of conventions similar to the examples in
[Labels](https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/) make uniqueness straightforward. Furthermore,
relatively narrowly used namespaces (e.g., per environment, per application) can
be used to reduce the set of resources that could potentially cause overlap.

In cases where users could be running misc. examples with inconsistent schemas,
or where tooling or components need to programmatically generate new objects to
be selected, there needs to be a straightforward way to generate unique label
sets. A simple way to ensure uniqueness of the set is to ensure uniqueness of a
single label value, such as by using a resource name, uid, resource hash, or
generation number.

Problems with uids and hashes, however, include that they have no semantic
meaning to the user, are not memorable nor readily recognizable, and are not
predictable. Lack of predictability obstructs use cases such as creation of a
replication controller from a pod, such as people want to do when exploring the
system, bootstrapping a self-hosted cluster, or deletion and re-creation of a
new RC that adopts the pods of the previous one, such as to rename it.
Generation numbers are more predictable and much clearer, assuming there is a
logical sequence. Fortunately, for deployments that's the case. For jobs, use of
creation timestamps is common internally. Users should always be able to turn
off auto-generation, in order to permit some of the scenarios described above.
Note that auto-generated labels will also become one more field that needs to be
stripped out when cloning a resource, within a namespace, in a new namespace, in
a new cluster, etc., and will need to be ignored around when updating a resource
via patch or read-modify-write sequence.

Inclusion of a system prefix in a label key is fairly hostile to UX. A prefix is
only necessary in the case that the user cannot choose the label key, in order
to avoid collisions with user-defined labels. However, I firmly believe that the
user should always be allowed to select the label keys to use on their
resources, so it should always be possible to override default label keys.

Therefore, resources supporting auto-generation of unique labels should have a
`uniqueLabelKey` field, so that the user could specify the key if they wanted
to, but if unspecified, it could be set by default, such as to the resource
type, like job, deployment, or replicationController. The value would need to be
at least spatially unique, and perhaps temporally unique in the case of job.

Annotations have very different intended usage from labels. They are
primarily generated and consumed by tooling and system extensions, or are used
by end-users to engage non-standard behavior of components.  For example, an
annotation might be used to indicate that an instance of a resource expects
additional handling by non-kubernetes controllers. Annotations may carry
arbitrary payloads, including JSON documents.  Like labels, annotation keys can
be prefixed with a governing domain (e.g. `example.com/key-name`).  Unprefixed
keys (e.g. `key-name`) are reserved for end-users.  Third-party components must
use prefixed keys.  Key prefixes under the "kubernetes.io" and "k8s.io" domains
are reserved for use by the kubernetes project and must not be used by
third-parties.

In early versions of Kubernetes, some in-development features represented new
API fields as annotations, generally with the form `something.alpha.kubernetes.io/name` or
`something.beta.kubernetes.io/name` (depending on our confidence in it). This
pattern is deprecated.  Some such annotations may still exist, but no new
annotations may be defined.  New API fields are now developed as regular fields.

Other advice regarding use of labels, annotations, taints, and other generic map keys by
Kubernetes components and tools:
  - Key names should be all lowercase, with words separated by dashes instead of camelCase
    - For instance, prefer `foo.kubernetes.io/foo-bar` over `foo.kubernetes.io/fooBar`, prefer
    `desired-replicas` over `DesiredReplicas`
  - Unprefixed keys are reserved for end-users.  All other labels and annotations must be prefixed.
  - Key prefixes under "kubernetes.io" and "k8s.io" are reserved for the Kubernetes
    project.
    - Such keys are effectively part of the kubernetes API and may be subject
      to deprecation and compatibility policies.
    - "kubernetes.io" is the preferred form for labels and annotations, "k8s.io" should not be used
      for new map keys.
  - Key names, including prefixes, should be precise enough that a user could
    plausibly understand where it came from and what it is for.
  - Key prefixes should carry as much context as possible.
    - For instance, prefer `subsystem.kubernetes.io/parameter` over `kubernetes.io/subsystem-parameter`
  - Use annotations to store API extensions that the controller responsible for
the resource doesn't need to know about, experimental fields that aren't
intended to be generally used API fields, etc. Beware that annotations aren't
automatically handled by the API conversion machinery.

## WebSockets and SPDY

Some of the API operations exposed by Kubernetes involve transfer of binary
streams between the client and a container, including attach, exec, portforward,
and logging. The API therefore exposes certain operations over upgradeable HTTP
connections ([described in RFC 2817](https://tools.ietf.org/html/rfc2817)) via
the WebSocket and SPDY protocols. These actions are exposed as subresources with
their associated verbs (exec, log, attach, and portforward) and are requested
via a GET (to support JavaScript in a browser) and POST (semantically accurate).

There are two primary protocols in use today:

1.  Streamed channels

    When dealing with multiple independent binary streams of data such as the
remote execution of a shell command (writing to STDIN, reading from STDOUT and
STDERR) or forwarding multiple ports the streams can be multiplexed onto a
single TCP connection. Kubernetes supports a SPDY based framing protocol that
leverages SPDY channels and a WebSocket framing protocol that multiplexes
multiple channels onto the same stream by prefixing each binary chunk with a
byte indicating its channel. The WebSocket protocol supports an optional
subprotocol that handles base64-encoded bytes from the client and returns
base64-encoded bytes from the server and character based channel prefixes ('0',
'1', '2') for ease of use from JavaScript in a browser.

2.  Streaming response

    The default log output for a channel of streaming data is an HTTP Chunked
Transfer-Encoding, which can return an arbitrary stream of binary data from the
server. Browser-based JavaScript is limited in its ability to access the raw
data from a chunked response, especially when very large amounts of logs are
returned, and in future API calls it may be desirable to transfer large files.
The streaming API endpoints support an optional WebSocket upgrade that provides
a unidirectional channel from the server to the client and chunks data as binary
WebSocket frames. An optional WebSocket subprotocol is exposed that base64
encodes the stream before returning it to the client.

Clients should use the SPDY protocols if their clients have native support, or
WebSockets as a fallback. Note that WebSockets is susceptible to Head-of-Line
blocking and so clients must read and process each message sequentially. In
the future, an HTTP/2 implementation will be exposed that deprecates SPDY.


## Validation

API objects are validated upon receipt by the apiserver. 
Validation can be implemented in two ways: declaratively, using tags on the Go type definitions, or manually, by writing validation functions. 
For all new APIs, declarative validation is the preferred approach for the validation rules it supports.  For more information see the [declarative validation documentation](api_changes.md#declarative-validation).  Validation errors are
flagged and returned to the caller in a `Failure` status with `reason` set to
`Invalid`. In order to facilitate consistent error messages, we ask that
validation logic adheres to the following guidelines whenever possible (though
exceptional cases will exist).

* Be as precise as possible.
* Telling users what they CAN do is more useful than telling them what they
CANNOT do.
* When asserting a requirement in the positive, use "must".  Examples: "must be
greater than 0", "must match regex '[a-z]+'".  Words like "should" imply that
the assertion is optional, and must be avoided.
* When asserting a formatting requirement in the negative, use "must not".
Example: "must not contain '..'".  Words like "should not" imply that the
assertion is optional, and must be avoided.
* When asserting a behavioral requirement in the negative, use "may not".
Examples: "may not be specified when otherField is empty", "only `name` may be
specified".
* When referencing a literal string value, indicate the literal in
single-quotes. Example: "must not contain '..'".
* When referencing another field name, indicate the name in back-quotes.
Example: "must be greater than \`request\`".
* When specifying inequalities, use words rather than symbols.  Examples: "must
be less than 256", "must be greater than or equal to 0".  Do not use words
like "larger than", "bigger than", "more than", "higher than", etc.
* When specifying numeric ranges, use inclusive ranges when possible.

## Automatic Resource Allocation And Deallocation

API objects often are [union](#Unions) object containing the following:
1. One or more fields identifying the `Type` specific to API object (aka the `discriminator`).
2. A set of N fields, only one of which should be set at any given time - effectively a union.

Controllers operating on the API type often allocate resources based on
the `Type` and/or some additional data provided by user. A canonical example
of this is the `Service` API object where resources such as IPs and network ports
will be set in the API object based on `Type`. When the user does not specify
resources, they will be allocated, and when the user specifies exact value, they will
be reserved or rejected.

When the user chooses to change the `discriminator` value (e.g., from `Type X` to `Type Y`) without
changing any other fields then the system should clear the fields that were used to represent `Type X`
in the union along with releasing resources that were attached to `Type X`. This should automatically
happen irrespective of how these values and resources were allocated (i.e., reserved by the user or
automatically allocated by the system. A concrete example of this is again `Service` API. The system
allocates resources such as `NodePorts` and `ClusterIPs` and automatically fill in the fields that
represent them in case of the service is of type `NodePort` or `ClusterIP` (`discriminator` values).
These resources and the fields representing them are automatically cleared when  the users changes
service type to `ExternalName` where these resources and field values no longer apply.

## Representing Allocated Values

Many API types include values that are allocated on behalf of the user from
some larger space (e.g. IP addresses from a range, or storage bucket names).
These allocations are usually driven by controllers asynchronously to the
user's API operations.  Sometimes the user can request a specific value and a
controller must confirm or reject that request.  There are many examples of
this in Kubernetes, and there a handful of patterns used to represent it.

The common theme among all of these is that the system should not trust users
with such fields, and must verify or otherwise confirm such requests before
using them.

Some examples:

* Service `clusterIP`: Users may request a specific IP in `spec` or will be
  allocated one (in the same `spec` field).  If a specific IP is requested, the
  apiserver will either confirm that IP is available or, failing that, will
  reject the API operation synchronously (rare).  Consumers read the result
  from `spec`.  This is safe because the value is either valid or it is never
  stored.
* Service `loadBalancerIP`: Users may request a specific IP in `spec` or will
  be allocated one which is reported in `status`.  If a specific IP is
  requested, the LB controller will either ensure that IP is available or
  report failure asynchronously.  Consumers read the result from `status`.
  This is safe because most users do not have acces to write to `status`.
* PersistentVolumeClaims: Users may request a specific PersistentVolume in
  `spec` or will be allocated one (in the same `spec` field).  If a specific PV
  is requested, the volume controller will either ensure that the volume is
  available or report failure asynchronously.  Consumers read the result by
  examining both the PVC and the PV.  This is more complicated than the others
  because the `spec` value is stored before being confirmed, which could
  (hypothetically, thanks to extra checking) lead to a user accessing someone
  else's PV.
* VolumeSnapshots: Users may request a particular source to be snaphotted in
  `spec`.  The details of the resulting snapshot is reflected in `status`.

A counter-example:

* Service `externalIPs`: Users must specify one or more specific IPs in `spec`.
  The system cannot easily verify those IPs (by their definition, they are
  external). Consumers read the result from `spec`.  This is UNSAFE and has
  caused problems with untrusted users.

In the past, API conventions dictated that `status` fields always come from
observation, which made some of these cases more complicated than necessary.
The conventions have been updated to allow `status` to hold such allocated
values.  This is not a one-size-fits-all solution, though.

### When to use a `spec` field

New APIs should almost never do this.  Instead, they should use `status`.
PersistentVolumes might have been simpler if we had done this.

### When to use a `status` field

Storing such values in `status` is the easiest and most straight-forward
pattern.  This is appropriate when:

* the allocated value is highly coupled to the rest of the object (e.g. pod
  resource allocations)
* the allocated value is always or almost always needed (i.e. most instances of
  this type will have a value)
* the schema and controller are known a priori (i.e. it's not an extension)
* it is "safe" to allow the controller(s) to write to `status` (i.e.
  there's low risk of them causing problems via other `status` fields).

Consumers of such values can look at the `status` field for the "final" value
or an error or condition indicating why the allocation could not be performed.

#### Sequencing operations

Since almost everything is happening asynchronously to almost everything else,
controller implementations should take care around the ordering of operations.
For example, whether the controller updates a `status` field before or after it
actuates a change depends on what guarantees need to be made to observers of
the system.  In some cases, writing to a `status` field represents an
acknowledgement or acceptance of a `spec` value, and it is OK to write it before
actuation.  However, if it would be problematic for a client to observe the
`status` value before it is actuated then the controller must actuate first and
update `status` afterward.  In some rarer cases, controllers will need to
acknowledge, then actuate, then update to a "final" value.

Controllers must take care to consider how a `status` field will be handled in
the case of interrupted control loops (e.g. controller crash and restart), and
must act idempotently and consistently.  This is particularly important when
using an informer-fed cache, which might not be updated with recent writes.
Using a resourceVersion precondition to detect the "conflict" is the common
pattern in this case.  See [this issue](http://issue.k8s.io/105199) for an
example.

### When to use a different type

Storing allocated values in a different type is more complicated but also more
flexible.  This is most appropriate when:

* the allocated value is optional (i.e. many instances of this type will not
  have a value at all)
* the schema and controller are not known a priori (i.e. it's an extension)
* the schema is sufficiently complicated (i.e. it doesn't make sense to burden
  the main type with it)
* access control for this type demands finer granularity than "all of status"
* the lifecycle of the allocated value is different than the lifecycle of the
  allocation holder

Services and Endpoints could be considered a form of this pattern, as could
PersistentVolumes and PersistentVolumeClaims.

When using this pattern, you must account for lifecycle of the allocated
objects (who cleans them up and when) as well as the "linkage" between them and
the main type (often using the same name, an object-ref field, or a selector).

There will always be some cases which could follow either path, and these will
need human evaluation to decide.  For example, Service `clusterIP` is highly
coupled to the rest of Service and most instances use it.  But it also is
strictly optional and has an increasingly complicated schema of related fields.
An argument could be made for either path.
