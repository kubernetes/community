# Expose `get` output from the server

Today, all clients must reproduce the tabular and describe output implemented in `kubectl` to perform simple lists
of objects. This logic in many cases is non-trivial and condenses multiple fields into succinct output. It also requires
that every client provide rendering logic for every possible type, including those provided by API aggregation or third
party resources which may not be known at compile time.

This proposal covers moving `get` and `describe` to the server to reduce total work for all clients and centralize these
core display options for better reuse and extension.


## Background

`kubectl get` is a simple tabular representation of one or more instances of a particular resource type. It is the primary
listing mechanism and so must be implemented for each type. Today, we have a single generic implementation for unrecognized
types that attempts to load information from the `metadata` field (assuming the object follows the metav1 Kubernetes API
schema). `get` supports a `wide` mode that includes additional columns. Users can add additional columns for labels via a
flag. Headers corresponding to the columns are optionally displayed.

`kubectl describe` shows a textual representation of individual objects that describes individual fields as subsequent
lines and uses indentation and nested tables to convey deeper structure on the resource (such as events for a pod or
each container). It sometimes retrieves related objects like events, pods for a replication controller, or autoscalers
for a deployment. It supports no significant flags.

The implementation of both is modeled as a registered function that takes an object or list of objects and outputs
semi-structured text.

## Goals

* Make it easy for a simple client to get a list of resources for a web UI or CLI output
* Support all existing options, leave open the door for future extension and experimentation
* Allow new API extensions and third party resources to be implemented server side, removing the need to version
  schemes for retrieving data from the server
* Keep implementation of `get` and `describe` simple
* Ease internationalization of `get` and `describe` output for all clients

## Non-Goals

* Deep customization of the returned output by the client


## Specification of server-side `get`

The server would return a `Table` object (working-name) that contains metadata for columns and one or more
rows composed of cells for each column.  Some additional data may be relevant for each row and returned by the
server. Since every object should have a `Table` representation, treat this as part of content negotiation
as described in [the alternative representations of objects proposal](alternate-api-representations.md).

Example request:

```
$ curl https://localhost:8443/api/v1/pods -H "Accept: application/json+vnd.kubernetes.as+meta.k8s.io+v1alpha1+Table"
{
    "kind": "Table",
    "apiVersion": "meta.k8s.io/v1alpha1",
    "headers": [
        {"name": "Name", "type": "string", "description": "The name of the pod, must be unique ..."},
        {"name": "Status", "type": "string", "description": "Describes the current state of the pod"},
        ...
    ],
    "items": [
        {"cells": ["pod1", "Failed - unable to start", ...]},
        {"cells": ["pod2", "Init 0/2", ...]},
        ...
    ]
}
```

This representation is also possible to return from a watch.  The watch can omit headers on subsequent queries.

```
$ curl https://localhost:8443/api/v1/pods?watch=1 -H "Accept: application/json+vnd.kubernetes.as+meta.k8s.io+v1alpha1+Table"
{
    "kind": "Table",
    "apiVersion": "meta.k8s.io/v1alpha1",
    // headers are printed first, in case the watch holds
    "headers": [
        {"name": "Name", "type": "string", "description": "The name of the pod, must be unique ..."},
        {"name": "Status", "type": "string", "description": "Describes the current state of the pod"},
        ...
    ]
}
{
    "kind": "Table",
    "apiVersion": "meta.k8s.io/v1alpha1",
    // headers are not returned here
    "items": [
        {"cells": ["pod1", "Failed - unable to start", ...]},
        ...
    ]
}
```

It can also be returned in CSV form:

```
$ curl https://localhost:8443/api/v1/pods -H "Accept: text/csv+vnd.kubernetes.as+meta.k8s.io+v1alpha1+Table"
Name,Status,...
pod1,"Failed - unable to start",...
pod2,"Init 0/2",...
...
```

To support "wide" format, columns may be marked with an optional priority field of increasing integers (default
priority 0):

```
{
    "kind": "Table",
    "apiVersion": "meta.k8s.io/v1alpha1",
    "headers": [
        ...
        {"name": "Node Name", "type": "string", "description": "The node the pod is scheduled on, empty if the pod is not yet scheduled", "priority": 1},
        ...
    ],
    ...
}
```

To allow label columns, and to enable integrators to build effective UIs, each row may contain an `object` field that
is either `PartialObjectMetadata` (a standard object containing only ObjectMeta) or the object itself. Clients may request
this field be set by specifying `?includeObject=None|Metadata|Self` on the query parameter.

```
GET ...?includeObject=Metadata
{
    "kind": "Table",
    "apiVersion": "meta.k8s.io/v1alpha1",
    "items": [
        ...
        {"cells": [...], "object": {"kind": "PartialObjectMetadata", "apiVersion":"meta.k8s.io/v1alpha1", "metadata": {"name": "pod1", "namespace": "pod2", "labels": {"a": "1"}, ...}},
        ...
    ]
}
```

The `Metadata` value would be the default. Clients that wish to print in an advanced manner may use `Self` to get the full
object and perform arbitrary transformations.

All fields on the server side are candidates for translation and localization changes can be delivered more
quickly and to all clients.

Third-party resources can more easily implement `get` in this fashion - instead of web dashboards and
`kubectl` both implementing their own logic to parse a particular version of Swagger or OpenAPI, the server
component performs the transformation.  The server encapsulates the details of printing.  Aggregated resources
automatically provide this behavior when possible.


### Specific features in `kubectl get`

Feature | Implementation
--- | --- 
sort-by | Continue to implement client-side (no server side sort planned)
custom-column (jsonpath) | Implement client-side by requesting object `?includeObject=Self` and parsing
custom-column (label) | Implement client-side by getting labels from metadata returned with each row
show-kind | Implement client-side by using the discovery info associated with the object (rather than being returned by server)
template | Implement client-side, bypass receiving table output and get raw objects
watch | Request Table output via the watch endpoint
export | Implement client-side, bypass receiving table output and get exported object
wide | Server should indicate which columns are "additional" via a field on the header column - client then shows those columns if it wants to
color (proposed) | Rows which should be highlighted should have a semantic field on the row - e.g. `alert: [{type: Warning, message: "This pod has been deleted"}]`.  Cells could be selected by adding an additional field `alert: [{type: Warning, ..., cells: [0, 1]}]`.


## Future considerations

* When we introduce server side paging, Table would be paged similar to how PodList or other types are paged. https://issues.k8s.io/2349
* More advanced output could in the future be provided by an external call-out or an aggregation API on the server side.
* `describe` could be managed on the server as well, with a similar generic format, and external outbound links used to reference other objects.


## Migration

Old clients will continue retrieving the primary representation.  Clients can begin using the optional `Accept`
header to indicate they want the simpler version, and if they receive a Table perform the new path, otherwise
fall back to client side functions.

Server side code would reuse the existing display functions but replace TabWriter with either a structured writer
or the tabular form.

