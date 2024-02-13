# 2023 Annual Report: SIG Network

## Current initiatives and Project Health

1. What work did the SIG do this year that should be highlighted?

<!--
   Some example items that might be worth highlighting:
   - Major KEP advancement
   - Important initiatives that aren't tracked via KEPs
   - Paying down significant tech debt
   - Governance and leadership changes
-->

2. Are there any areas and/or subprojects that your group needs help with (e.g. fewer than 2 active OWNERS)?

<!--
   Note: This list is generated from the KEP metadata in kubernetes/enhancements repository.
      If you find any discrepancy in the generated list here, please check the KEP metadata.
      Please raise an issue in kubernetes/community, if the KEP metadata is correct but the generated list is incorrect.
-->

3. Did you have community-wide updates in 2023 (e.g. KubeCon talks)?

<!--
  Examples include links to email, slides, or recordings.
-->

4. KEP work in 2023 (v1.27, v1.28, v1.29):

  - Alpha
    - [1860 - Make Kubernetes aware of the load balancer behaviour](https://github.com/kubernetes/enhancements/tree/master/keps/sig-network/1860-kube-proxy-IP-node-binding) - v1.29
    - [1880 - Multiple Service CIDRs](https://github.com/kubernetes/enhancements/tree/master/keps/sig-network/1880-multiple-service-cidrs) - v1.27
    - [3836 - Kube-proxy improved ingress connectivity reliability](https://github.com/kubernetes/enhancements/tree/master/keps/sig-network/3836-kube-proxy-improved-ingress-connectivity-reliability) - v1.28
    - [3866 - Add an nftables-based kube-proxy backend](https://github.com/kubernetes/enhancements/tree/master/keps/sig-network/3866-nftables-proxy) - v1.29
    - [4004 - Deprecate status.nodeInfo.kubeProxyVersion field](https://github.com/kubernetes/enhancements/tree/master/keps/sig-network/4004-deprecate-kube-proxy-version) - v1.29

  - Beta
    - [2681 - Field status.hostIPs added for Pod](https://github.com/kubernetes/enhancements/tree/master/keps/sig-network/2681-pod-host-ip) - v1.29
    - [3458 - Remove transient node predicates from KCCM's service controller](https://github.com/kubernetes/enhancements/tree/master/keps/sig-network/3458-remove-transient-node-predicates-from-service-controller) - v1.27
    - [3705 - Cloud Dual-Stack --node-ip Handling](https://github.com/kubernetes/enhancements/tree/master/keps/sig-network/3705-cloud-node-ips) - v1.29

  - Stable
    - [1669 - Proxy Terminating Endpoints](https://github.com/kubernetes/enhancements/tree/master/keps/sig-network/1669-proxy-terminating-endpoints) - v1.28
    - [2091 - Add support for AdminNetworkPolicy resources](https://github.com/kubernetes/enhancements/tree/master/keps/sig-network/2091-admin-network-policy) - v1.27
    - [2438 - Dual Stack API Server](https://github.com/kubernetes/enhancements/tree/master/keps/sig-network/2438-dual-stack-apiserver) - v1.27
    - [2595 - Expanded DNS Configuration](https://github.com/kubernetes/enhancements/tree/master/keps/sig-network/2595-expanded-dns-config) - v1.28
    - [3178 - Cleaning up IPTables Chain Ownership](https://github.com/kubernetes/enhancements/tree/master/keps/sig-network/3178-iptables-cleanup) - v1.28
    - [3453 - Minimize iptables-restore input size](https://github.com/kubernetes/enhancements/tree/master/keps/sig-network/3453-minimize-iptables-restore) - v1.28
    - [3668 - Reserve Nodeport Ranges For Dynamic And Static Port Allocation](https://github.com/kubernetes/enhancements/tree/master/keps/sig-network/3668-reserved-service-nodeport-range) - v1.29
    - [3726 - standard-application-protocols](https://github.com/kubernetes/enhancements/tree/master/keps/sig-network/3726-standard-application-protocols) - v1.27

## [Subprojects](https://git.k8s.io/community/sig-network#subprojects)


**Continuing:**
  - cluster-proportional-autoscaler
  - cluster-proportional-vertical-autoscaler
  - external-dns
  - gateway-api
  - ingress
  - iptables-wrappers
  - kpng
  - kube-dns
  - network-policy
  - pod-networking

## [Working groups](https://git.k8s.io/community/sig-network#working-groups)

**Retired in 2023:**
 - Multitenancy
**Continuing:**
 - IoT Edge
 - Policy
 - Structured Logging

## Operational

Operational tasks in [sig-governance.md]:
- [ ] [README.md] reviewed for accuracy and updated if needed
- [ ] [CONTRIBUTING.md] reviewed for accuracy and updated if needed
- [ ] Other contributing docs (e.g. in devel dir or contributor guide) reviewed for accuracy and updated if needed
- [ ] Subprojects list and linked OWNERS files in [sigs.yaml] reviewed for accuracy and updated if needed
- [ ] SIG leaders (chairs, tech leads, and subproject leads) in [sigs.yaml] are accurate and active, and updated if needed
- [ ] Meeting notes and recordings for 2023 are linked from [README.md] and updated/uploaded if needed


[CONTRIBUTING.md]: https://git.k8s.io/community/sig-network/CONTRIBUTING.md
[sig-governance.md]: https://git.k8s.io/community/committee-steering/governance/sig-governance.md
[README.md]: https://git.k8s.io/community/sig-network/README.md
[sigs.yaml]: https://git.k8s.io/community/sigs.yaml
[devel]: https://git.k8s.io/community/contributors/devel/README.md
