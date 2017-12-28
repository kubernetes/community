Story so far is described at https://github.com/kubernetes/community/blob/master/contributors/design-proposals/api-machinery/csi-client-structure-proposal.md
- linked from https://github.com/kubernetes-client/community/blob/master/design-docs/clients-library-structure.md

Would like libs to be more "fluent" - should join up with the native language features e.g. JavaDoc, per-platform

OpenAPI (Swagger) has limitations which force ugly workarounds.
 - e.g. doesn't cover returning different types in response to a call e.g. a result or an error.
 - OpenAPI 3 fixes a number of these limitations

Key difficulty is which verbs apply to which objects - Go doesn't natively have this concept.

OpenAPI doc is produced by running an api-server, which knows how to dispatch verbs to objects, and having it generate the OpenAPI

Should we deprecate SPDY in favour of websockets? (probably yes)

What do people want in a client library?
 - Go lib smaller than 40MB
 - Go lib with dynamic types (this exists?)
 - Rust
 - C++

Are there client libs for things like service catalog?  Where do those live?

How can clients talk protobufs to api-server?  Set the content-type.

K8s protobuf def uses an 'envelope' type covering things like storage version

Protobuf not self-describing so not great for api-server CRUD operations

A few bits of extra information in the protobuf definition would be great - e.g. pluralisation

Client SIG is relatively under-represented and relatively easy to join in - please come along!

kubeconfig semantics are not standardised - when reimplemented in different languages this can give unexpected results.  Example: URL with no scheme supported by Go but not Java

How should clients deal with rolling authorisation certificates?

There are type wrinkles, like "intorstring" and "quantity"