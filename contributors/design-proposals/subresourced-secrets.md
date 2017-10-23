# Subresourced Secret Data

Author: Eric Chiang (eric.chiang@coreos.com)

Note: This is a write up and formalization of an idea @liggitt (Red Hat) had during a SIG-auth meeting.

## Background

Efforts around [improving secret support in Kubernetes](https://docs.google.com/document/d/1JAwPuZg47UhfRVlof-lMw08OJztunW8pvTNxDK3rCF8/edit) are largely aimed at integrations with external secret stores. It not unreasonable to imagine a future where there are alternative mechanisms to the secrets API for delivering secret material to containers.

However, many controllers and applications which run on top of Kubernetes still use, and still will use the secret API to manage sensitive data. The most common example are ingress controllers, which dynamically read TLS assets.

When using the secrets API, listing and watching secrets are an extremely powerful capabilities to give clients because there’s no way for another application in the same namespace to hide secrets from that client. However, in the wild, we’ve found it’s very common for controllers that deal with secrets to require these capabilities. This is what the most popular [ingress controller does](https://github.com/kubernetes/ingress/issues/816), for example. Changing the controllers to not use list and watch often means large codebase changes, performance hits, and possibly even re-architecting the kind of resources use.

Part of this issue stems from the fact that there’s no separation between the secret metadata and the secret data itself. For example, admins can’t permission a client to view a secret’s type without being able to view the secret itself. 

## Overview

This proposal would add an opt-in subresource to secret resources. Secrets that opt-in to this feature would return no data in response to list, watch, and get requests. Individual get request to the “data” subresource would be the only way to access the data value. 

### API type changes

To opt-in to this new feature, the secret API type would add the following field. 

```go
type Secret struct {
    // Existing fields

    metav1.TypeMeta
    metav1.ObjectMeta
    Data map[string][]byte
    StringData map[string]string
    Type SecretType

    // New field

    // Subresourced, when set to true, causes the “data” field to
    // be omitted from all responses, except when performing "get" on
    // the “secrets/data” subresource. This prevents operations like
    // “list” and “watch” from accessing the data of secrets that
    // enable this option.
    //
    // Getting “secrets/data” always returns the full “data” field,
    // even when this option is disabled.
    //
    // This field requires the API server feature flag
    // “SubresourcedSecretData”.
    //
    // +optional
    Subresourced bool 
}
```

### Changes

The `SubresourcedSecretData` feature flag would have the following impact on Kubernetes components.

If enabled on the API server, the flag would allow clients to create subresourced secrets. It would also update the [node authorizer](https://kubernetes.io/docs/admin/authorization/node/) and default RBAC rules to permit nodes accessing secrets through this value.

If enabled on the kubelet, the kubelet would request secret data through the subresource.

If enabled on the controller manager, newly created service account secrets would use the subresourced flag. 

## Arguments against

### Artificial separation

One could argue that this proposal only works because, for a time, many apps won't understand the subresource and therefore we've just created an artificial separation. Couldn't we accomplish something similar by to just creating a second secret resource? In that case, early adopters of the second resource would have a place to hide their data until the old apps caught up.

If we envision a world where all apps understood and used this new subresource, the secrets API would still see the following improvements.

* Clients can watch and list secrets without accessing the data. client-go's informer framework continues to work with secrets.
* Forces clients to get individual secrets when they actually want the secret data. Makes it impossible to write an app that always needs access to all secret data.

### Other potential secrets API improvements?

Previous suggestions of improvements to the secrets API involve making RBAC more expressive. Can we do label based RBAC filters? Can we make watch still work?

An assumption of this proposal is that separating secret matadata and secret data is a complementary goal to any RBAC change we might make. Also, compared to other proposed changes, this proposal is extremely low overhead. It doesn't involve making watch more sophisticated, it doesn't expand RBAC to have to read object data.

What's less clear is if this is the most impactful change we can make. It won't immediately benefit any controllers using the existing secret resource, which an RBAC change would do.
