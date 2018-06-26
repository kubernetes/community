# Validated Pod Annotations

Author: @mikedanese

# Objective

Propose a mechanism to enable adminstrators to apply extra authorization checks
on operations that grant direct or indirect access to a pods runtime
environment.

# Background

A number of pod identity integrations need to inject credentials from external
identity providers into a pod runtime environment. A common and desirable
property of an identity solutions is:

> An actor should be able to access credentials of an identity only if the actor
> has explicit authority to act as that identity.

Today, naive pod identity integrations have weaker guarantees. To illustrate the
issue, suppose a pod is running with access to a kerberos credential injected
with a flex volume. A user can access the kerberos credential if the user can do
any of the following:

1. Exec into the pod.
1. Update the image of the pod to run custom code to export the credential.
1. Update the image in the PodSpec of any controller that is in the owner
   hierarchy of the pod.
1. Create another pod that uses the same credential and runs custom code to
   export the credential.
1. Create a controller that then creates a pod that uses the same credential and
   runs custom code to export the credential.

Ideally, a user can only access the kerberos credential if the user can:

1. Act as the identity.

This constraint allows a security auditor to audit identity access with only the
delegation policy. Identity integrations can work around this by implementing
custom admission controllers but the issue is common and general enough to
warrent a consistent API (and thus requires a solution in core).

# Proposal

We can add a "validatedAnnotations" field on PodSpec and validate these
annotations in an admission controller. The field will be a map of string to
string. The admission controller will validate these annotations by issuing
SubjectAccessReviews on operations that implicitly grant access to a pod runtime
environment. These operations are:

1. exec on pod
1. create pod
1. update spec of pod
1. create controller that creates pods from a podspec
1. update spec of controller that creates pods from a podspec

Take the following pod:

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: app
spec:
  containers:
  - name: app
    image: gcr.io/mycompany/app:1.0
  validatedAnnotations:
  - cloud.google.com/serviceaccount: gcs@mycompany.gserviceaccount.com
  - kerberos: legacy-frontend-app
```

If user `employee-10` attempts any of the listed operations, the admission
controller will issue the following two SubjectAccessReviews:

```yaml
apiVersion": "authorization.k8s.io/v1",
kind": "SubjectAccessReview",
spec:
  user: employee-10
  resourceAttributes:
    group: authorization.k8s.io
    resource: validatedannotations
    resourceName: cloud.google.com/serviceaccount
    subresource: gcs@mycompany.gserviceaccount.com
    verb: use
---
apiVersion": "authorization.k8s.io/v1",
kind": "SubjectAccessReview",
spec:
  user: employee-10
  resourceAttributes:
    group: authorization.k8s.io
    resource: validatedannotations
    resourceName: kerberos
    subresource: legacy-frontend-app
    verb: use
```

These SubjectAccessReviews are in addition to the standard authorization on the
operation. If all SubjectAccessReviews are allowed, then the operation succeeds.

These annotations can then be asserted by Kubernetes and trusted by an identity
integration. Kubernetes can inform an identity integration of these annotations
by:

* passing them on the flex volume driver commandline
* embedding them in a serviceaccount ID token of a pod

An identity integration can also directly observe these annotations on pods by
reading from the Kubernetes API.

By serving an authorizer that translates the SubjectAccessReview into a
delegation check and only injecting credentials into pods conditional on the
validated annotations of the pods, an identity integration can guarantee the
desired constraint.
