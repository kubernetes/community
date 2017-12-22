# Policy-based Federated Resource Placement

This document proposes a design for policy-based control over placement of
Federated resources.

Tickets:

- https://github.com/kubernetes/kubernetes/issues/39982

Authors:

- Torin Sandall (torin@styra.com, tsandall@github) and Tim Hinrichs
  (tim@styra.com).
- Based on discussions with Quinton Hoole (quinton.hoole@huawei.com,
  quinton-hoole@github), Nikhil Jindal (nikhiljindal@github).

##  Background

Resource placement is a policy-rich problem affecting many deployments.
Placement may be based on company conventions, external regulation, pricing and
performance requirements, etc. Furthermore, placement policies evolve over time
and vary across organizations. As a result, it is difficult to anticipate the
policy requirements of all users.

A simple example of a placement policy is

> Certain apps must be deployed on clusters in EU zones with sufficient PCI
> compliance.

The [Kubernetes Cluster
Federation](/contributors/design-proposals/multicluster/federation.md#policy-engine-and-migrationreplication-controllers)
design proposal includes a pluggable policy engine component that decides how
applications/resources are placed across federated clusters.

Currently, the placement decision can be controlled for Federated ReplicaSets
using the `federation.kubernetes.io/replica-set-preferences` annotation. In the
future, the [Cluster
Selector](https://github.com/kubernetes/kubernetes/issues/29887) annotation will
provide control over placement of other resources. The proposed design supports
policy-based control over both of these annotations (as well as others).

This proposal is based on a POC built using the Open Policy Agent project. [This
short video (7m)](https://www.youtube.com/watch?v=hRz13baBhfg) provides an
overview and demo of the POC.

## Design

The proposed design uses the [Open Policy Agent](http://www.openpolicyagent.org)
project (OPA) to realize the policy engine component from the Federation design
proposal. OPA is an open-source, general purpose policy engine that includes a
declarative policy language and APIs to answer policy queries.

The proposed design allows administrators to author placement policies and have
them automatically enforced when resources are created or updated. The design
also covers support for automatic remediation of resource placement when policy
(or the relevant state of the world) changes.

In the proposed design, the policy engine (OPA) is deployed on top of Kubernetes
in the same cluster as the Federation Control Plane:

![Architecture](https://docs.google.com/drawings/d/1kL6cgyZyJ4eYNsqvic8r0kqPJxP9LzWVOykkXnTKafU/pub?w=807&h=407)

The proposed design is divided into following sections:

1. Control over the initial placement decision (admission controller)
1. Remediation of resource placement (opa-kube-sync/remediator)
1. Replication of Kubernetes resources (opa-kube-sync/replicator)
1. Management and storage of policies (ConfigMap)

### 1. Initial Placement Decision

To provide policy-based control over the initial placement decision, we propose
a new admission controller that integrates with OPA:

When admitting requests, the admission controller executes an HTTP API call
against OPA. The API call passes the JSON representation of the resource in the
message body.

The response from OPA contains the desired value for the resource’s annotations
(defined in policy by the administrator). The admission controller updates the
annotations on the resource and admits the request:

![InitialPlacement](https://docs.google.com/drawings/d/1c9PBDwjJmdv_qVvPq0sQ8RVeZad91vAN1XT6K9Gz9k8/pub?w=812&h=288)

The admission controller updates the resource by **merging** the annotations in
the response with existing annotations on the resource. If there are overlapping
annotation keys the admission controller replaces the existing value with the
value from the response.

#### Example Policy Engine Query:

```http
POST /v1/data/io/k8s/federation/admission HTTP/1.1
Content-Type: application/json
```

```json
{
  "input": {
    "apiVersion": "extensions/v1beta1",
    "kind": "ReplicaSet",
    "metadata": {
      "annotations": {
        "policy.federation.alpha.kubernetes.io/eu-jurisdiction-required": "true",
        "policy.federation.alpha.kubernetes.io/pci-compliance-level": "2"
      },
      "creationTimestamp": "2017-01-23T16:25:14Z",
      "generation": 1,
      "labels": {
        "app": "nginx-eu"
      },
      "name": "nginx-eu",
      "namespace": "default",
      "resourceVersion": "364993",
      "selfLink": "/apis/extensions/v1beta1/namespaces/default/replicasets/nginx-eu",
      "uid": "84fab96d-e188-11e6-ac83-0a580a54020e"
    },
    "spec": {
      "replicas": 4,
      "selector": {...},
      "template": {...},
    }
  }
}
```

#### Example Policy Engine Response:

```http
HTTP/1.1 200 OK
Content-Type: application/json
```

```json
{
  "result": {
    "annotations": {
      "federation.kubernetes.io/replica-set-preferences": {
        "clusters": {
          "gce-europe-west1": {
            "weight": 1
          },
          "gce-europe-west2": {
            "weight": 1
          }
        },
        "rebalance": true
      }
    }
  }
}
```

> This example shows the policy engine returning the replica-set-preferences.
> The policy engine could similarly return a desired value for other annotations
> such as the Cluster Selector annotation.

#### Conflicts

A conflict arises if the developer and the policy define different values for an
annotation. In this case, the developer's intent is provided as a policy query
input and the policy author's intent is encoded in the policy itself. Since the
policy is the only place where both the developer and policy author intents are
known, the policy (or policy engine) should be responsible for resolving the
conflict.

There are a few options for handling conflicts. As a concrete example, this is
how a policy author could handle invalid clusters/conflicts:

```
package io.k8s.federation.admission

errors["requested replica-set-preferences includes invalid clusters"] {
  invalid_clusters = developer_clusters - policy_defined_clusters
  invalid_clusters != set()
}

annotations["replica-set-preferences"] = value {
  value = developer_clusters & policy_defined_clusters
}

# Not shown here:
#
# policy_defined_clusters[...] { ... }
# developer_clusters[...] { ... }
```

The admission controller will execute a query against
/io/k8s/federation/admission and if the policy detects an invalid cluster, the
"errors" key in the response will contain a non-empty array. In this case, the
admission controller will deny the request.

```http
HTTP/1.1 200 OK
Content-Type: application/json
```

```json
{
  "result": {
    "errors": [
      "requested replica-set-preferences includes invalid clusters"
    ],
    "annotations": {
      "federation.kubernetes.io/replica-set-preferences": {
        ...
      }
    }
  }
}
```

This example shows how the policy could handle conflicts when the author's
intent is to define clusters that MAY be used. If the author's intent is to
define what clusters MUST be used, then the logic would not use intersection.

#### Configuration

The admission controller requires configuration for the OPA endpoint:

```
{
  "EnforceSchedulingPolicy": {
    "url": “https://opa.federation.svc.cluster.local:8181/v1/data/io/k8s/federation/annotations”,
    "token": "super-secret-token-value"
  }
}
```

- `url` specifies the URL of the policy engine API to query. The query response
  contains the annotations to apply to the resource.
- `token` specifies a static token to use for authentication when contacting the
  policy engine. In the future, other authentication schemes may be supported.

The configuration file is provided to the federation-apiserver with the
`--admission-control-config-file` command line argument.

The admission controller is enabled in the federation-apiserver by providing the
`--admission-control` command line argument. E.g.,
`--admission-control=AlwaysAdmit,EnforceSchedulingPolicy`.

The admission controller will be enabled by default.

#### Error Handling

The admission controller is designed to **fail closed** if policies have been
created.

Request handling may fail because of:

- Serialization errors
- Request timeouts or other network errors
- Authentication or authorization errors from the policy engine
- Other unexpected errors from the policy engine

In the event of request timeouts (or other network errors) or back-pressure
hints from the policy engine, the admission controller should retry after
applying a backoff. The admission controller should also create an event so that
developers can identify why their resources are not being scheduled.

Policies are stored as ConfigMap resources in a well-known namespace. This
allows the admission controller to check if one or more policies exist. If one
or more policies exist, the admission controller will fail closed. Otherwise
the admission controller will **fail open**.

### 2. Remediation of Resource Placement

When policy changes or the environment in which resources are deployed changes
(e.g. a cluster’s PCI compliance rating gets up/down-graded), resources might
need to be moved for them to obey the placement policy. Sometimes administrators
may decide to remediate manually, other times they may want Kubernetes to
remediate automatically.

To automatically reschedule resources onto desired clusters, we introduce a
remediator component (**opa-kube-sync**) that is deployed as a sidecar with OPA.

![Remediation](https://docs.google.com/drawings/d/1ehuzwUXSpkOXzOUGyBW0_7jS8pKB4yRk_0YRb1X4zsY/pub?w=812&h=288)

The notifications sent to the remediator by OPA specify the new value for
annotations such as replica-set-preferences.

When the remediator component (in the sidecar) receives the notification it
sends a PATCH request to the federation-apiserver to update the affected
resource. This way, the actual rebalancing of ReplicaSets is still handled by
the [Rescheduling
Algorithm](/contributors/design-proposals/multicluster/federated-replicasets.md)
in the Federated ReplicaSet controller.

The remediator component must be deployed with a kubeconfig for the
federation-apiserver so that it can identify itself when sending the PATCH
requests. We can use the same mechanism that is used for the
federation-controller-manager (which also needs ot identify itself when sending
requests to the federation-apiserver.)

### 3. Replication of Kubernetes Resources

Administrators must be able to author policies that refer to properties of
Kubernetes resources. For example, assuming the following sample policy (in
English):

> Certain apps must be deployed on Clusters in EU zones with sufficient PCI
> compliance.

The policy definition must refer to the geographic region and PCI compliance
rating of federated clusters. Today, the geographic region is stored as an
attribute on the cluster resource and the PCI compliance rating is an example of
data that may be included in a label or annotation.

When the policy engine is queried for a placement decision (e.g., by the
admission controller), it must have access to the data representing the
federated clusters.

To provide OPA with the data representing federated clusters as well as other
Kubernetes resource types (such as federated ReplicaSets), we use a sidecar
container that is deployed alongside OPA. The sidecar (“opa-kube-sync”) is
responsible for replicating Kubernetes resources into OPA:

![Replication](https://docs.google.com/drawings/d/1XjdgszYMDHD3hP_2ynEh_R51p7gZRoa1DBTi4yq1rc0/pub?w=812&h=288)

The sidecar/replicator component will implement the (somewhat common) list/watch
pattern against the federation-apiserver:

- Initially, it will GET all resources of a particular type.
- Subsequently, it will GET with the **watch** and **resourceVersion**
  parameters set and process add, remove, update events accordingly.

Each resource received by the sidecar/replicator component will be pushed into
OPA. The sidecar will likely rely on one of the existing Kubernetes Go client
libraries to handle the low-level list/watch behavior.

As new resource types are introduced in the federation-apiserver, the
sidecar/replicator component will need to be updated to support them. As a
result, the sidecar/replicator component must be designed so that it is easy to
add support for new resource types.

Eventually, the sidecar/replicator component may allow admins to configure which
resource types are replicated. As an optimization, the sidecar may eventually
analyze policies to determine which resource properties are requires for policy
evaluation. This would allow it to replicate the minimum amount of data into
OPA.

### 4. Policy Management

Policies are written in a text-based, declarative language supported by OPA. The
policies can be loaded into the policy engine either on startup or via HTTP
APIs.

To avoid introducing additional persistent state, we propose storing policies
in ConfigMap resources in the Federation Control Plane inside a well-known
namespace (e.g., `kube-federationscheduling-policy`). The ConfigMap resources
will be replicated into the policy engine by the sidecar.

The sidecar can establish a watch on the ConfigMap resources in the Federation
Control Plane. This will enable hot-reloading of policies whenever they change.

## Applicability to Other Policy Engines

This proposal was designed based on a POC with OPA, but it can be applied to
other policy engines as well. The admission and remediation components are
comprised of two main pieces of functionality: (i) applying annotation values to
federated resources and (ii) asking the policy engine for annotation values. The
code for applying annotation values is completely independent of the policy
engine. The code that asks the policy engine for annotation values happens both
within the admission and remediation components. In the POC, asking OPA for
annotation values amounts to a simple RESTful API call that any other policy
engine could implement.

## Future Work

- This proposal uses ConfigMaps to store and manage policies. In the future, we
  want to introduce a first-class **Policy** API resource.
