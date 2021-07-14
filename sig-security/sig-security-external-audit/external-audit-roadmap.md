Past external security audits have not been comprehensive of the entire Kubernetes project.
This roadmap lists previously audited focus areas and focus areas requested to be included in future audits.
The Kubernetes community is invited to create issues and PRs to request additional components to be audited.


| **Kubernetes Focus Area** | **Audit Year**| **Links** |
|---------------------------|---------------|-----------|
| Networking | 2019 | |
| Cryptography | 2019 | |
| Authentication & Authorization (including Role Based Access Controls) | 2019 | |
| Secrets Management | 2019 | |
| Multi-tenancy isolation: Specifically soft (non-hostile co-tenants) | 2019 | |
| kube-apiserver | 2021 | |
| kube-scheduler | 2021 | |
| etcd (in the context of Kubernetes use of etcd) | 2021 | |
| kube-controller-manager | 2021 | |
| cloud-controller-manager | 2021 | |
| kubelet | 2021 | https://github.com/kubernetes/kubelet https://github.com/kubernetes/kubernetes/tree/master/staging/src/k8s.io/kubelet |
| kube-proxy | 2021 | https://github.com/kubernetes/kubernetes/tree/master/staging/src/k8s.io/kube-proxy https://github.com/kubernetes/kube-proxy |
| secrets-store-csi-driver | 2021 | https://github.com/kubernetes-sigs/secrets-store-csi-driver |
| cluster API | TBD | https://github.com/kubernetes-sigs/cluster-api |
| kubectl | TBD | https://github.com/kubernetes/kubectl |
| kubeadm | TBD | https://github.com/kubernetes/kubeadm |
| metrics server | TBD | https://github.com/kubernetes-sigs/metrics-server
| nginx-ingress (in the context of a Kubernetes ingress controller) | TBD | https://github.com/kubernetes/ingress-nginx
| kube-state-metrics | TBD | https://github.com/kubernetes/kube-state-metrics
| node feature discovery | TBD | https://github.com/kubernetes-sigs/node-feature-discovery
| hierarchial namespace | TBD | https://github.com/kubernetes-sigs/multi-tenancy/tree/master/incubator/hnc
| pod security policy replacement | TBD | https://github.com/kubernetes/enhancements/tree/master/keps/sig-auth/2579-psp-replacement
| CoreDNS (in the context of Kubernetes use of CoreDNS) | TBD | Concept: https://kubernetes.io/docs/concepts/services-networking/dns-pod-service/ Reference: https://kubernetes.io/docs/tasks/administer-cluster/dns-custom-nameservers/ |
| cluster autoscaler | TBD | https://github.com/kubernetes/autoscaler/tree/master/cluster-autoscaler |
| kube rbac proxy | TBD | https://github.com/brancz/kube-rbac-proxy |
| kms plugins | TBD | https://kubernetes.io/docs/tasks/administer-cluster/kms-provider/#implementing-a-kms-plugin |
| cni plugins | TBD | https://github.com/containernetworking/cni |
| csi plugins | TBD | https://github.com/kubernetes-csi |
| aggregator layer | TBD | https://github.com/kubernetes/kube-aggregator |