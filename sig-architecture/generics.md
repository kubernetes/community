# Background

In general, Kubernetes has not restricted using new Go features, we've quickly
adopted new standard library types and methods and will continue to do so.

Generally the latest stable go release is in use on the main development branch.
This includes all of the staging libraries (client-go etc.) that originate in the
main [kubernetes/kubernetes](https://github.com/kubernetes/kubernetes) repository.

In Kubernetes v1.24 we shipped Go 1.18 (which adds support for generics)
rather late in the release cycle, so we temporarily prohibited using generics
in case we ran into other issues and needed to roll back to unblock the release.
Now that v1.24.0 is out, use of generics should be allowed.

# Generics Policy

Generics may be used in Kubernetes starting in v1.25, with the following restrictions 
applying only until v1.24 is [out of support][version-support]:

- Generics should **not** be used in Kubernetes libraries used across multiple Kubernetes
versions, that is the non "staged" libraries like:
  - [k8s.io/utils](https://github.com/kubernetes/utils)
  - [sigs.k8s.io/yaml](https://github.com/kubernetes-sigs/yaml)
  - [k8s.io/klog](https://github.com/kubernetes/klog)
  - etc.

- Generics should be **avoided** when writing Kubernetes bug fixes that are likely to be backported, to streamline cherry-picking to older release branches.

These restrictions should be considered lifted when v1.24 is out of support.

## Recommendations for Reviewers

- Consider if proposed generics pull requests improve maintainability and readability.

- The current generics implementation is known to have some performance issues
depending on usage: consider requesting benchmarks before / after the changes.

[version-support]: https://kubernetes.io/releases/patch-releases/#support-period
