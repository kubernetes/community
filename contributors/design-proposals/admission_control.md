# Kubernetes Proposal - Admission Control

**Related PR:**

| Topic | Link |
| ----- | ---- |
| Separate validation from RESTStorage | http://issue.k8s.io/2977 |

## Background

High level goals:
* Enable an easy-to-use mechanism to provide admission control to cluster.
* Enable a provider to support multiple admission control strategies or author
their own.
* Ensure any rejected request can propagate errors back to the caller with why
the request failed.

Authorization via policy is focused on answering if a user is authorized to
perform an action.

Admission Control is focused on if the system will accept an authorized action.

Kubernetes may choose to dismiss an authorized action based on any number of
admission control strategies.

This proposal documents the basic design, and describes how any number of
admission control plug-ins could be injected.

Implementation of specific admission control strategies are handled in separate
documents.

## kube-apiserver

The kube-apiserver takes the following OPTIONAL arguments to enable admission
control:

| Option | Behavior |
| ------ | -------- |
| admission-control | Comma-delimited, ordered list of admission control choices to invoke prior to modifying or deleting an object. |
| admission-control-config-file | File with admission control configuration parameters to boot-strap plug-in. |

An **AdmissionControl** plug-in is an implementation of the following interface:

```go
package admission

// Attributes is an interface used by AdmissionController to get information about a request
// that is used to make an admission decision.
type Attributes interface {
	// GetName returns the name of the object as presented in the request.  On a CREATE operation, the client
	// may omit name and rely on the server to generate the name.  If that is the case, this method will return
	// the empty string
	GetName() string
	// GetNamespace is the namespace associated with the request (if any)
	GetNamespace() string
	// GetResource is the name of the resource being requested.  This is not the kind.  For example: pods
	GetResource() schema.GroupVersionResource
	// GetSubresource is the name of the subresource being requested.  This is a different resource, scoped to the parent resource, but it may have a different kind.
	// For instance, /pods has the resource "pods" and the kind "Pod", while /pods/foo/status has the resource "pods", the sub resource "status", and the kind "Pod"
	// (because status operates on pods). The binding resource for a pod though may be /pods/foo/binding, which has resource "pods", subresource "binding", and kind "Binding".
	GetSubresource() string
	// GetOperation is the operation being performed
	GetOperation() Operation
	// GetObject is the object from the incoming request prior to default values being applied
	GetObject() runtime.Object
	// GetOldObject is the existing object. Only populated for UPDATE requests.
	GetOldObject() runtime.Object
	// GetKind is the type of object being manipulated.  For example: Pod
	GetKind() schema.GroupVersionKind
	// GetUserInfo is information about the requesting user
	GetUserInfo() user.Info
}

// Interface is an abstract, pluggable interface for Admission Control decisions.
type Interface interface {
	// Admit makes an admission decision based on the request attributes
	Admit(a Attributes) (err error)

	// Handles returns true if this admission controller can handle the given operation
	// where operation can be one of CREATE, UPDATE, DELETE, or CONNECT
	Handles(operation Operation) bool
}
```

A **plug-in** must be compiled with the binary, and is registered as an
available option by providing a name, and implementation of admission.Interface.

```go
func init() {
  admission.RegisterPlugin("AlwaysDeny", func(client client.Interface, config io.Reader) (admission.Interface, error) { return NewAlwaysDeny(), nil })
}
```

A **plug-in** must be added to the imports in [plugins.go](../../cmd/kube-apiserver/app/plugins.go)

```go
  // Admission policies
  _ "k8s.io/kubernetes/plugin/pkg/admission/admit"
  _ "k8s.io/kubernetes/plugin/pkg/admission/alwayspullimages"
  _ "k8s.io/kubernetes/plugin/pkg/admission/antiaffinity"
  ...
  _ "<YOUR NEW PLUGIN>"
```

Invocation of admission control is handled by the **APIServer** and not
individual **RESTStorage** implementations.

This design assumes that **Issue 297** is adopted, and as a consequence, the
general framework of the APIServer request/response flow will ensure the
following:

1. Incoming request
2. Authenticate user
3. Authorize user
4. If operation=create|update|delete|connect, then admission.Admit(requestAttributes)
   - invoke each admission.Interface object in sequence
5. Case on the operation:
   - If operation=create|update, then validate(object) and persist
   - If operation=delete, delete the object
   - If operation=connect, exec

If at any step, there is an error, the request is canceled.


<!-- BEGIN MUNGE: GENERATED_ANALYTICS -->
[![Analytics](https://kubernetes-site.appspot.com/UA-36037335-10/GitHub/docs/design/admission_control.md?pixel)]()
<!-- END MUNGE: GENERATED_ANALYTICS -->
