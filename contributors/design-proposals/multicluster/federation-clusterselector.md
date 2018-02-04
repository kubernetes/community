# ClusterSelector Federated Resource Placement

This document proposes a design for label based control over placement of
Federated resources.

Tickets:

- https://github.com/kubernetes/kubernetes/issues/29887

Authors:

- Dan Wilson (emaildanwilson@github.com).
- Nikhil Jindal (nikhiljindal@github).

##  Background

End users will often need a simple way to target a subset of clusters for deployment of resources. In some cases this will be for a specific cluster in other cases it will be groups of clusters.
A few examples...

1. Deploy the foo service to all clusters in Europe
1. Deploy the bar service to cluster test15
1. Deploy the baz service to all prod clusters globally

Currently, it's possible to control placement decision of Federated ReplicaSets
using the `federation.kubernetes.io/replica-set-preferences` annotation. This provides functionality to change the number of ReplicaSets created on each Federated Cluster, by setting the quantity for each Cluster by Cluster Name. Since cluster names are required, in situations where clusters are add/removed from Federation it would require the object definitions to change in order to maintain the same configuration. From the example above, if a new cluster is created in Europe and added to federation, then the replica-set-preferences would need to be updated to include the new cluster name.

This proposal is to provide placement decision support for all object types using Labels on the Federated Clusters as opposed to cluster names. The matching language currently used for nodeAffinity placement decisions onto nodes can be leveraged. 

Carrying forward the examples from above...

1. "location=europe"
1. "someLabel exists"
1. "environment notin ["qa", "dev"]

## Design

The proposed design uses a ClusterSelector annotation that has a value that is parsed into a struct definition that follows the same design as the [NodeSelector type used w/ nodeAffinity](https://git.k8s.io/kubernetes/pkg/api/types.go#L1972) and will also use the [Matches function](https://git.k8s.io/apimachinery/pkg/labels/selector.go#L172) of the apimachinery project to determine if an object should be sent on to federated clusters or not.

In situations where objects are not to be forwarded to federated clusters, instead a delete api call will be made using the object definition. If the object does not exist it will be ignored.

The federation-controller will be used to implement this with shared logic stored as utility functions to reduce duplicated code where appropriate.

### End User Functionality 
The annotation `federation.alpha.kubernetes.io/cluster-selector` is used on kubernetes objects to specify additional placement decisions that should be made. The value of the annotation will be a json object of type ClusterSelector which is an array of type ClusterSelectorRequirement.

Each ClusterSelectorRequirement is defined in three possible parts consisting of
1. Key - Matches against label keys on the Federated Clusters.
1. Operator - Represents how the Key and/or Values will be matched against the label keys and values on the Federated Clusters one of ("In", in", "=", "==", "NotIn", notin", "Exists", "exists", "!=", "DoesNotExist", "!", "Gt", "gt", "Lt", "lt").
1. Values - Matches against the label values on the Federated Clusters using the Key specified. When the operator is "Exists", "exists", "DoesNotExist" or "!" then Values should not be specified.

Example ConfigMap that uses the ClusterSelector annotation. The yaml format is used here to show that the value of the annotation will still be json.
```yaml
apiVersion: v1
data:
  myconfigkey: myconfigdata
kind: ConfigMap
metadata:
  annotations:
    federation.alpha.kubernetes.io/cluster-selector: '[{"key": "location", "operator":
      "in", "values": ["europe"]}, {"key": "environment", "operator": "==", "values":
      ["prod"]}]'
  creationTimestamp: 2017-02-07T19:43:40Z
  name: myconfig
```

In order for the configmap in the example above to be forwarded to any Federated Clusters they MUST have two Labels: "location" with at least one value of "europe" and "environment" that has a value of "prod".

### Matching Logic

The logic to determine if an object is sent to a Federated Cluster will have two rules.

1. An object with no `federation.alpha.kubernetes.io/cluster-selector` annotation will always be forwarded on to all Federated Clusters even if they have labels configured. (this ensures no regression from existing functionality)

1. If an object contains the `federation.alpha.kubernetes.io/cluster-selector` annotation then ALL ClusterSelectorRequirements must match in order for the object to be forwarded to the Federated Cluster.

1. If `federation.kubernetes.io/replica-set-preferences` are also defined they will be applied AFTER the ClusterSelectorRequirements.

## Open Questions

1. Should there be any special considerations for when dependent resources would not be forwarded together to a Federated Cluster.
1. How to improve usability of this feature long term. It will certainly help to give first class API support but easier ways to map labels or requirements to objects may be required.
