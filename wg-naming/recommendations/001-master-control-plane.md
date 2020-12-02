# Recommendation: master -> control plane

**Last Updated**: 2020-10-16 

**Status**: Accepted

## Recommendation

Within the Kubernetes codebase, the term “master” is often used in reference to
[the kubernetes control plane][architecture], either as a whole or to some
subset of the components within.  We recommend **control plane** to refer to the
set of components as a whole.  We recommend **context-specific alternatives**
when talking about individual components or the roles they serve.

## Suggested Alternatives

### Control Plane
- e.g. "The control plane is the set of all components responsible for
  controlling a kubernetes cluster"
- e.g. "The control plane is the thing that can be communicated with in order to
  control a kubernetes cluster"
- If it matters which specific component(s) or component instance is being
  communicated with, be specific:
  - endpoint, e.g. "we don't care whether a load balancer, apiserver, or other
    component is behind this endpoint; we will talk only to this endpoint and no
other"
  - instance, e.g. "we are going to create multiple control plane apiserver
    instances; we will talk only to this specific instance"
  - <component name>, e.g. "we will simulate an etcd fault by running this
    command on instances where etcd is hosted"

### Control Plane Node
- "A node that hosts components that are part of the control plane", e.g.
  - https://github.com/kubernetes/enhancements/tree/master/keps/sig-cluster-lifecycle/kubeadm/2067-rename-master-label-taint
  - https://github.com/kubernetes/kubernetes/pull/95053/files#diff-b4f6256abfd125f7ce69fd1ba1eaf595R886
- Also relevant for terms other than node, e.g. control plane machine, control
  plane host, control plane vm, control plane instance, e.g.
  - https://github.com/kubernetes/kubernetes/blob/99cc89b7da32d9c06916deb50b27fdb46934b777/cluster/gce/gci/master-helper.sh#L33
    should be `create-control-plane-instance`

### Leader
- e.g. "The leader is the winner of a leader election"
- When using an adjective instead of a noun to describe this concept, there is
  likely a more specific term:
  - active/passive, e.g. "The active controller-manager process is the one that
    wins its leader election"
  - primary/replica

## Other Considerations
- if the api field, flag, code, command etc. uses the literal word 'master' and
  cannot be immediately changed or has no alternative available, use this word
only in direct reference to the code item
  - `master` branch being renamed to `main` branch falls under this
- there are no strict guarantees about how and where control plane components
  can and will run, e.g.
  - they may run on machines that are not registered as Nodes for the Kubernetes
    cluster they control
  - they may run alongside user workloads on Nodes in the Kubernetes cluster
    they control
  - there may be one to many instances of each control plane component
- When using terminology that could be seen as generic, consider whether there
  is enough context available to disambiguate potential meanings for the reader.
e.g.
  - "instance" - an instance of what?
  - "apiserver" - is this a generic apiserver? is this a kubernetes apiserver?
    is this an apiserver that has been configured to be part of the kubernetes
control plane?
  - "endpoint" - what's at the other end of this endpoint?

## Context

Master raises first-order concerns according to [our language evaluation
framework][framework]:
- it is overtly racist (ref: [django][django-master], [Drupal][drupal-master],
  [IETF][ietf-master],
[Google][https://developers.google.com/style/word-list#master])

Master also raises third-order concerns:
- within kubernetes it is used to represent a variety of overlapping or
  unrelated concepts (see the variety of suggested alternatives)
- one class of usage represents a set of false assumptions:
  - there is exactly one instance of each control plane component
  - there is exactly one kubernetes node that hosts all of these components
  - these components are guaranteed to run in a specific manner (systemd units,
    on certain ports, etc.)

Prior discussions:
- [kubernetes-wg-naming@ - proposal: master/slave
  alternatives][wg-naming-thread]
- [kubernetes-sig-architecture@ - Re: the way we discuss control plane
  members][sig-arch-thread]

## Consequences

TODO

- hound search that approximates excluding some master-branch-in-docs
  references, and references in vendor/:
https://cs.k8s.io/?q=master%5B%5E%2F%5D&i=nope&files=%5E%5B%5Ev%5D&repos=
- references to master in test/e2e/framework
  (https://github.com/kubernetes/kubernetes/issues/94901)
- references to master in test/integration
  (https://github.com/kubernetes/kubernetes/issues/94900)
- known names/flags/fields/labels/annotations that may take time to change
  - `"system:masters"` aka
    [k8s.io/apiserver/pkg/authentication/user.SystemPrivilegedGroup][system-privileged-group]
  - `node-role.kubernetes.io/master` (tracking issue for KEP
    https://github.com/kubernetes/enhancements/issues/2067)
  - `--endpoint-reconciler-type master-count`
  - probably more

[architecture]: https://git.k8s.io/community/contributors/design-proposals/architecture/architecture.md#architecture
[wg-naming-thread]: https://groups.google.com/g/kubernetes-wg-naming/c/VqrBCdUHdPc
[sig-arch-thread]: https://groups.google.com/u/1/g/kubernetes-sig-architecture/c/ZKUOPy2PNJ4/m/q3dC6pNtBQAJ
[framework]: https://git.k8s.io/community/wg-naming/language-evaluation-framework.md
[ietf-master]: https://tools.ietf.org/id/draft-knodel-terminology-00.html#master-slave
[drupal-master]: https://www.drupal.org/node/2275877
[django-master]: https://github.com/django/django/pull/2692
[system-privileged-group]: https://github.com/kubernetes/kubernetes/blob/a9d1482710a4c4baf112890882f4ab3d4be158a6/staging/src/k8s.io/apiserver/pkg/authentication/user/user.go#L71
