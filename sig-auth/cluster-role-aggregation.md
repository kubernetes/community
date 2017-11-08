# Cluster Role Aggregation
In order to support easy RBAC integration for CustomResources and Extension
APIServers, we need to have a way for API extenders to add permissions to the
"normal" roles for admin, edit, and view.

These roles express an intent for the namespaced power of administrators of the
namespace (manage ownership), editors of the namespace (manage content like
pods), and viewers of the namespace (see what is present).  As new APIs are
made available, these roles should reflect that intent to prevent migration
concerns every time a new API is added.

To do this, we will allow one ClusterRole to be built out of a selected set of
ClusterRoles.

## API Changes
```yaml
aggregationRule:
  selectors:
  - matchLabels:
      rbac.authorization.k8s.io/aggregate-to-admin: true
```

```go
// ClusterRole is a cluster level, logical grouping of PolicyRules that can be referenced as a unit by a RoleBinding or ClusterRoleBinding.
type ClusterRole struct {
	metav1.TypeMeta
	// Standard object's metadata.
	metav1.ObjectMeta

	// Rules holds all the PolicyRules for this ClusterRole
	Rules []PolicyRule

	// AggregationRule is an optional field that describes how to build the Rules for this ClusterRole.
	// If AggregationRule is set, then the Rules are controller managed and direct changes to Rules will be
	// stomped by the controller.
	AggregationRule *AggregationRule
}

// AggregationRule describes how to locate ClusterRoles to aggregate into the ClusterRole
type AggregationRule struct {
	// Selector holds a list of selectors which will be used to find ClusterRoles and create the rules.
	// If any of the selectors match, then the ClusterRole's permissions will be added
	Selectors []metav1.LabelSelector
}
```

The `aggregationRule` stanza contains a list of LabelSelectors which are used
to select the set of ClusterRoles which should be combined.  When
`aggregationRule` is set, the list of `rules` becomes controller managed and is
subject to overwriting at any point.

`aggregationRule` needs to be protected from escalation.  The simplest way to
do this is to restrict it to users with verb=`*`, apiGroups=`*`, resources=`*`.  We
could later loosen it by using a covers check against all aggregated rules
without changing backward compatibility.

## Controller
There is a controller which watches for changes to ClusterRoles and then
updates all aggregated ClusterRoles if their list of Rules has changed.  Since
there are relatively few ClusterRoles, it checks them all and most
short-circuit.

## The Payoff
If you want to create a CustomResource for your operator and you want namespace
admin's to be able to create one, instead of trying to:
 1. Create a new ClusterRole
 2. Update every namespace with a matching RoleBinding
 3. Teach everyone to add the RoleBinding to all their admin users
 4. When you remove it, clean up dangling RoleBindings
 
 Or
 
 1. Make a non-declarative patch against the admin ClusterRole
 2. When you remove it, try to safely create a new non-declarative patch to
 remove it.
 
You can simply create a new ClusterRole like
```yaml
apiVersion: rbac.authorization.k8s.io/v1beta1
kind: ClusterRole
metadata:
  name: etcd-operator-admin
  label:
    rbac.authorization.k8s.io/aggregate-to-admin: true
rules:
- apiGroups:
  - etcd.database.coreos.com
  resources:
  - etcdclusters
  verbs:
  - "*"
```
alongside your CustomResourceDefinition.  The admin role is updated correctly and 
removal is a `kubectl delete -f` away.