# Webhook Bootstrapping

## Background
[Admission webhook](./admission-control-webhooks.md) is a feature that
dynamically extends Kubernetes admission chain. Because the admission webhooks
are in the critical path of admitting REST requests, broken webhooks could block
the entire cluster, even blocking the reboot of the webhooks themselves. This
design presents a way to avoid such bootstrap deadlocks.

## Objective
- If one or more webhooks are down, it should be able restart them automatically.
- If a core system component that supports webhooks is down, the component
  should be able to restart.

## Design idea
We add a selector to the admission webhook configuration, which will be compared
to the labels of namespaces. Only objects in the matching namespaces are
subjected to the webhook admission. A cluster admin will want to exempt these
namespaces from webhooks:
- Namespaces where this webhook and other webhooks are deployed in;
- Namespaces where core system components are deployed in.

## API Changes
`ExternalAdmissionHook` is the dynamic configuration API of an admission webhook.
We will add a new field `NamespaceSelector` to it:

```golang
type ExternalAdmissionHook struct {
    Name string
    ClientConfig AdmissionHookClientConfig
    Rules []RuleWithOperations
    FailurePolicy *FailurePolicyType
    // Only objects in matching namespaces are subjected to this webhook.
    // LabelSelector.MatchExpressions allows exclusive as well as inclusive
    // matching, so you can use this // selector as a whitelist or a blacklist.
    // For example, to apply the webhook to all namespaces except for those have
    // labels with key "runlevel" and value equal to "0" or "1": 
    // metav1.LabelSelctor{MatchExpressions: []LabelSelectorRequirement{
    // 	{
    // 		Key:      "runlevel",
    // 		Operator: metav1.LabelSelectorOpNotIn,
    // 		Value:    []string{"0", "1"},
    // 	},
    // }}
    // As another example, to only apply the webhook to the namespaces that have
    // labels with key “environment” and value equal to “prod” and “staging”:
    // metav1.LabelSelctor{MatchExpressions: []LabelSelectorRequirement{
    // 	{
    // 		Key:      "environment",
    // 		Operator: metav1.LabelSelectorOpIn,
    // 		Value:    []string{"prod", "staging"},
    // 	},
    // }}
    // See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more examples of label selectors.
    NamespaceSelector *metav1.LabelSelector
}
```

## Guidelines on namespace labeling
The mechanism depends on cluster admin properly labelling the namespaces. We
will provide guidelines on the labelling scheme. One suggestion is labelling
namespaces with runlevels. The design of runlevels is out of the scope of this
document (tracked in
[#54522](https://github.com/kubernetes/kubernetes/issues/54522)), a strawman
runlevel scheme is:

- runlevel 0: namespaces that host core system components, like kube-apiserver
  and kube-controller-manager.
- runlevel 1: namespaces that host add-ons that are part of the webhook serving
  stack, e.g., kube-dns.
- runlevel 2: namespaces that host webhooks deployments and services.

`ExternalAdmissionHook.NamespaceSelector` should be configured to skip all the
above namespaces. In the case where some webhooks depend on features offered by
other webhooks, the system administrator could extend this concept further (run
level 3, 4, 5, …) to accommodate them.

## Security implication
The mechanism depends on namespaces being properly labelled. We assume only
highly privileged users can modify namespace labels. Note that the system
already relies on correct namespace annotations, examples include the
podNodeSelector admission plugin, and the podTolerationRestriction admission
plugin etc.

# Considered Alternatives
- Allow each webhook to exempt one namespace
  - Doesn’t work: if there are two webhooks in two namespaces both blocking pods
    startup, they will block each other.
- Put all webhooks in a single namespace and let webhooks exempt that namespace,
  e.g., deploy webhooks in the “kube-system” namespace and exempt the namespace.
  - It doesn’t provide sufficient isolation. Not all objects in the
    “kube-system” namespace should bypass webhooks.
- Add namespace selector to webhook configuration, but use the selector to match
  the name of namespaces
  ([#1191](https://github.com/kubernetes/community/pull/1191)).
  - Violates k8s convention. The matching label (key=name, value=<namespace’s
    name>) is imaginary.
  - Hard to manage. Namespace’s name is arbitrary.
