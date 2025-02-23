# Project Naming

The Kubernetes project has historically named most of our components/projects
with a "kube" substring like `kube-proxy`, `kubectl`, `kube-apiserver`,
`kube-controller-manager`, `kube-scheduler`, `minikube`, ...

"Kubernetes" and "K8s" are registered [trademarks under the Linux Foundation](https://www.linuxfoundation.org/legal/trademarks).

We recommend *but do not require* Kubernetes SIGs consider leveraging these when
naming new subprojects, to further identify as an official project.

We ask that third-party projects refrain from using confusing names that sound
like an official Kubernetes subproject, i.e. "kube", "k8s", "kubernetes"
substrings and pick some alternate creative name, with a notable exception for
[Conformant Distros](https://www.cncf.io/training/certification/software-conformance/).
Many existing projects have not done this, and we are not seeking to alter them, but
do remember that "k8s" and "kubernetes" are protected trademarks.

If you'd like to start a new sub-project within the Kubernetes project, see our
[repository donation / creation guidelines](https://git.k8s.io/community/github-management/kubernetes-repositories.md).
