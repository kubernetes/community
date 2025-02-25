# 2024 Annual Report: SIG Network

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


3. Did you have community-wide updates in 2024 (e.g. KubeCon talks)?

<!--
  Examples include links to email, slides, or recordings.
-->

4. KEP work in 2024 (v1.30, v1.31, v1.32):

- Alpha
  - [4427 - Relaxed DNS search string validation](https://github.com/kubernetes/enhancements/tree/master/keps/sig-network/4427-relaxed-dns-search-validation) - v1.32
  - [784 - Kube Proxy component configuration updates and graduation](https://github.com/kubernetes/enhancements/tree/master/keps/sig-network/784-kube-proxy-component-config) - v1.33

- Beta
  - [1880 - Multiple Service CIDRs](https://github.com/kubernetes/enhancements/tree/master/keps/sig-network/1880-multiple-service-cidrs) - v1.31
  - [3866 - Add an nftables-based kube-proxy backend](https://github.com/kubernetes/enhancements/tree/master/keps/sig-network/3866-nftables-proxy) - v1.31
  - [4444 - Traffic Distribution for Services](https://github.com/kubernetes/enhancements/tree/master/keps/sig-network/4444-service-traffic-distribution) - v1.31

- Stable
  - [2681 - Field status.hostIPs added for Pod](https://github.com/kubernetes/enhancements/tree/master/keps/sig-network/2681-pod-host-ip) - v1.30
  - [3458 - Remove transient node predicates from KCCM's service controller](https://github.com/kubernetes/enhancements/tree/master/keps/sig-network/3458-remove-transient-node-predicates-from-service-controller) - v1.30
  - [3705 - Cloud Dual-Stack --node-ip Handling](https://github.com/kubernetes/enhancements/tree/master/keps/sig-network/3705-cloud-node-ips) - v1.30
  - [3836 - Kube-proxy improved ingress connectivity reliability](https://github.com/kubernetes/enhancements/tree/master/keps/sig-network/3836-kube-proxy-improved-ingress-connectivity-reliability) - v1.31
  - [1860 - Make Kubernetes aware of the load balancer behaviour](https://github.com/kubernetes/enhancements/tree/master/keps/sig-network/1860-kube-proxy-IP-node-binding) - v1.32
  - [2433 - Topology Aware Hints](https://github.com/kubernetes/enhancements/tree/master/keps/sig-network/2433-topology-aware-hints) - v1.33

## [Subprojects](https://git.k8s.io/community/sig-network#subprojects)

**New in 2024:**
  - gateway-api-inference-extension
  - ingate
  - knftables
  - multi-network
  - node-ipam-controller

**Continuing:**
  - cluster-proportional-autoscaler
  - cluster-proportional-vertical-autoscaler
  - external-dns
  - gateway-api
  - ingress
  - iptables-wrappers
  - kube-dns
  - network-policy
  - pod-networking

**Retired in 2024:**
  - kpng

## [Working groups](https://git.k8s.io/community/sig-network#working-groups)

**New in 2024:**
 - Device Management
 - Serving

**Retired in 2024:**
 - IoT Edge

**Continuing:**
 - Policy
 - Structured Logging

## Operational

Operational tasks in [sig-governance.md]:
- [ ] [README.md] reviewed for accuracy and updated if needed
- [ ] [CONTRIBUTING.md] reviewed for accuracy and updated if needed
- [ ] Other contributing docs (e.g. in devel dir or contributor guide) reviewed for accuracy and updated if needed
- [ ] Subprojects list and linked OWNERS files in [sigs.yaml] reviewed for accuracy and updated if needed
- [ ] SIG leaders (chairs, tech leads, and subproject leads) in [sigs.yaml] are accurate and active, and updated if needed
- [ ] Meeting notes and recordings for 2024 are linked from [README.md] and updated/uploaded if needed


[CONTRIBUTING.md]: https://git.k8s.io/community/sig-network/CONTRIBUTING.md
[sig-governance.md]: https://git.k8s.io/community/committee-steering/governance/sig-governance.md
[README.md]: https://git.k8s.io/community/sig-network/README.md
[sigs.yaml]: https://git.k8s.io/community/sigs.yaml
[devel]: https://git.k8s.io/community/contributors/devel/README.md
