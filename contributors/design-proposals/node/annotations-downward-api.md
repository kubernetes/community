# Exposing annotations via environment downward API

Author: Michal Rostecki \<michal@kinvolk.io\>

## Introduction

Annotations of the pod can be taken through the Kubernetes API, but currently
there is no way to pass them to the application inside the container. This means
that annotations can be used by the core Kubernetes services and the user outside
of the Kubernetes cluster.

Of course using Kubernetes API from the application running inside the container
managed by Kubernetes is technically possible, but that's an idea which denies
the principles of microservices architecture.

The purpose of the proposal is to allow to pass the annotation as the environment
variable to the container.

### Use-case

The primary usecase for this proposal are StatefulSets. There is an idea to expose
StatefulSet index to the applications running inside the pods managed by StatefulSet.
Since StatefulSet creates pods as the API objects, passing this index as an
annotation seems to be a valid way to do this. However, to finally pass this
information to the containerized application, we need to pass this annotation.
That's why the downward API for annotations is needed here.

## API

The exact `fieldPath` to the annotation will have the following syntax:

```
metadata.annotations['annotationKey']
```

Which means that:
- the *annotationKey* will be specified inside brackets (`[`, `]`) and single quotation
  marks (`'`)
- if the *annotationKey* contains `[`, `]` or `'` characters inside, they will need to
  be escaped (like `\[`, `\]`, `\'`) and having these characters unescaped should result
  in validation error

Examples:
- `metadata.annotations['spec.pod.beta.kubernetes.io/statefulset-index']`
- `metadata.annotations['foo.bar/example-annotation']`
- `metadata.annotations['foo.bar/more\'complicated\]example\[with\'characters"to-escape']`

So, assuming that we would want to pass the `pod.beta.kubernetes.io/statefulset-index`
annotation as a `STATEFULSET_INDEX` variable, the environment variable definition
will look like:

```
env:
  - name: STATEFULSET_INDEX
    valueFrom:
      fieldRef:
        fieldPath: metadata.annotations['spec.pod.beta.kubernetes.io/statefulset-index']
```

## Implementation

In general, this environment downward API part will be implemented in the same
place as the other metadata - as a label conversion function.

