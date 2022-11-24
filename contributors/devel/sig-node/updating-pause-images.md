# Updating Pause Image References

The Kubernetes `pause` image is used to run a 'parent' container for all
the other containers inside your pod. This image is referenced across many
projects within the Kubernetes ecosystem, and it's important that we keep them
all up to date to ensure stable platform support and to avoid confusing things
for end users.

However, we do need to update those images in a somewhat staggered way across
the ecosystem, rather than doing it all at once.

### Useful Links

- [Pause Image Changelog][pause-cl]
- [What is the role of 'pause' container?](https://groups.google.com/g/kubernetes-users/c/jVjv0QK4b_o)
- [The Almighty Pause Container](https://www.ianlewis.org/en/almighty-pause-container)

### Coordinating updates

When a new `pause` image is released, we generally try to update it in the
following order, this is to try and ensure that repositories stay in sync, while
keeping changes well tested and developers unblocked.

#### First level (core projects)

- [kubernetes/kubernetes][k-k-repo]
- [kubernetes/test-infra][k-infra-repo]
- [kubernetes/k8s.io][k-k8s.io]

#### Second level (kubernetes-sigs)

- [kubernetes-sigs/image-builder][ks-image-builder]
- [kubernetes-sigs/kind][ks-kind]

#### Third level (runtimes)

- [cri-o/cri-o][crio]
- [containerd/containerd][containerd]

#### Fourth level (other projects)

- [kubernetes/minikube][k-minikube]

[k-k-repo]: https://github.com/kubernetes/kubernetes
[k-infra-repo]: https://github.com/kubernetes/test-infra
[k-k8s.io]: https://github.com/kubernetes/k8s.io
[ks-kind]: https://github.com/kubernetes-sigs/kind
[ks-image-builder]: https://github.com/kubernetes-sigs/image-builder
[crio]: https://github.com/cri-o/cri-o
[containerd]: https://github.com/containerd/containerd
[k-minikube]: https://github.com/kubernetes/minikube
[pause-cl]: https://github.com/kubernetes/kubernetes/blob/master/build/pause/CHANGELOG.md
