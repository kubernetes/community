# 2024 Annual Report: SIG Network

## Current initiatives and Project Health

1. What work did the SIG do this year that should be highlighted?

This year we have several highlights from our various sub-projects, these are
the major themes and releases:

- [External DNS](https://github.com/kubernetes-sigs/external-dns)
  - Two new reviewers were added this year as the project continues to grow
  - Shipped important maintenance release [v0.15.x](https://github.com/kubernetes-sigs/external-dns/releases/tag/v0.15.0)
- [Ingress NGINX](https://github.com/kubernetes/ingress-nginx)
  - Shipped a large number of `v1.x.x` patch and minor releases all the way from `v1.8.x` to [v1.12.0](https://github.com/kubernetes/ingress-nginx/releases/tag/controller-v1.12.0)
- [InGate](https://github.com/kubernetes-sigs/ingate)
  - Project just started at the end of 2024 and is starting to plan the roadmap.
- [Multi-Network](https://github.com/kubernetes-sigs/multi-network-api)
  - Moved away from core development to CRD based addon approach
  - Work started on supporting [Dynamic Resource Allocation (DRA)]
- [Network Policy](https://github.com/kubernetes-sigs/network-policy-api)
  - Working on getting `AdminNetworkPolicy` (ANP) and `BaselineAdminNetworkPolicy` (BANP) APIs to `Beta`
    - Two new implementations joined us this year in supporting the APIs
  - Delivered features: [FQDN Selectors], [Policy Assistant CLI], and [Tenancy API]
  - Work started on [Service Account Selectors], and [Dry-Run Mode]
- [Gateway API](https://github.com/kubernetes-sigs/gateway-api)
  - Delivered [GRPCRoute], [ParentReference Port], [Service Mesh Support] and [Conformance Profiles and Reports] as GA in release [v1.1][gwv1.1]
  - Delivered [HTTPRoute Timeouts], [Gateway Infrastructure Labels], and [Backend Protocol Support] as GA in release [v1.2][gwv1.2]
  - Shipped two releases of our [ingress2gateway] utility: [v0.2.0][i2gv0.2], [v0.3.0][i2gv0.3]
  - Shipped our first release of our [gwctl] utility: [v0.1.0][gwctlv0.1]
- [Gateway API Inference Extensions (GIE)](https://github.com/kubernetes-sigs/gateway-api-inference-extension)
  - Shipped our first release [v0.1.0](https://github.com/kubernetes-sigs/gateway-api-inference-extension)
- [IP Masq Agent](https://github.com/kubernetes-sigs/ip-masq-agent)
  - Shipped important maintenance release [v2.12.0](https://github.com/kubernetes-sigs/ip-masq-agent/releases/tag/v2.12.0)
- [Cluster Proportional Autoscaler](https://github.com/kubernetes-sigs/cluster-proportional-autoscaler)
  - Shipped important maintenance release [v1.9.x](https://github.com/kubernetes-sigs/cluster-proportional-autoscaler/releases/tag/v1.9.0)
- [Blixt](https://github.com/kubernetes-sigs/blixt)
  - Started a [major rewrite of the control-plane] and added [L4 Gateway API Support] ([TCPRoute], [UDPRoute])

[Dynamic Resource Allocation (DRA)]:https://kubernetes.io/docs/concepts/scheduling-eviction/dynamic-resource-allocation/
[FQDN Selectors]:https://network-policy-api.sigs.k8s.io/npeps/npep-133/
[Policy Assistant CLI]:https://github.com/kubernetes-sigs/network-policy-api/releases/tag/v0.0.1-policy-assistant
[Tenancy API]:https://network-policy-api.sigs.k8s.io/npeps/npep-122/
[Service Account Selectors]:https://github.com/kubernetes-sigs/network-policy-api/pull/274
[Dry-Run Mode]:https://github.com/kubernetes-sigs/network-policy-api/pull/276
[GRPCRoute]:https://gateway-api.sigs.k8s.io/guides/grpc-routing/
[ParentReference Port]:https://gateway-api.sigs.k8s.io/reference/spec/#gateway.networking.k8s.io%2fv1.ParentReference
[Service Mesh Support]:https://gateway-api.sigs.k8s.io/mesh/
[Conformance Profiles and Reports]:https://gateway-api.sigs.k8s.io/geps/gep-1709/
[HTTPRoute Timeouts]:https://gateway-api.sigs.k8s.io/geps/gep-1742/
[Gateway Infrastructure Labels]:https://gateway-api.sigs.k8s.io/reference/spec/#gateway.networking.k8s.io/v1.GatewayInfrastructure
[Backend Protocol Support]:https://github.com/kubernetes/enhancements/tree/master/keps/sig-network/3726-standard-application-protocols
[ingress2gateway]:https://github.com/kubernetes-sigs/ingress2gateway
[gwctl]:https://github.com/kubernetes-sigs/gwctl
[major rewrite of the control-plane]:https://github.com/kubernetes-sigs/blixt/milestone/8
[L4 Gateway API Support]:https://github.com/kubernetes-sigs/blixt/issues/303
[TCPRoute]:https://gateway-api.sigs.k8s.io/references/spec/#gateway.networking.k8s.io/v1alpha2.TCPRoute
[UDPRoute]:https://gateway-api.sigs.k8s.io/references/spec/#gateway.networking.k8s.io/v1alpha2.UDPRoute

[gwv1.1]:https://github.com/kubernetes-sigs/gateway-api/releases/tag/v1.1.0
[gwv1.2]:https://github.com/kubernetes-sigs/gateway-api/releases/tag/v1.2.0
[i2gv0.2]:https://github.com/kubernetes-sigs/ingress2gateway/releases/tag/v0.2.0
[i2gv0.3]:https://github.com/kubernetes-sigs/ingress2gateway/releases/tag/v0.3.0
[gwctlv0.1]:https://github.com/kubernetes-sigs/gwctl/releases/tag/v0.1.0

2. Are there any areas and/or subprojects that your group needs help with (e.g. fewer than 2 active OWNERS)?

While there are no SIG Network projects which _couldn't_ benefit from more help
and contributions, the following is a list of specific projects where there are
known blockers (or otherwise critical needs) where more contributions, feedback,
or implementation support is the key to "unblocking" them:

- [InGate] is a new project just getting started creating a Gateway API implementation, and needs contributors to join and get to the first release.
- [Network Policy] could use more implementations to join and support some of our newer APIs (e.g. [AdminNetworkPolicy]) to help them graduate towards GA.
- [MultiNetwork] has started anew with CRD based development and is in need of use cases and implementations to join to get to the first release.
- [Gateway API]
  - has a large number of [features which are not yet standard] which could use implementations to join us to help us graduate those towards GA.
  - has two CLI utilities which need more users and implementations to get involved to provide feedback: [ingress2gateway], [gwctl]
- [Gateway API Inference Extension] is a new project which provides AI/ML inference routing and load-balancing features. It is in need of more implementations to bring use cases and and contribute in order to help graduate the APIs towards GA.
- If you're interested in [DRA] for networking, we started a new [CNI] driver this year which could use more contributors: [cni-dra-driver]

[README.md]:https://github.com/kubernetes/community/blob/master/sig-network/README.md
[charter]:https://github.com/kubernetes/community/blob/master/sig-network/charter.md
[InGate]:https://github.com/kubernetes-sigs/ingate
[Network Policy]:https://github.com/kubernetes-sigs/network-policy-api
[AdminNetworkPolicy]:https://network-policy-api.sigs.k8s.io/reference/spec/#policy.networking.k8s.io%2fv1alpha1.AdminNetworkPolicy
[MultiNetwork]:https://github.com/kubernetes-sigs/multi-network-api
[Gateway API]:https://github.com/kubernetes-sigs/gateway-api
[features which are not yet standard]:https://gateway-api.sigs.k8s.io/geps/overview/
[ingress2gateway]:https://github.com/kubernetes-sigs/ingress2gateway
[gwctl]:https://github.com/kubernetes-sigs/gwctl
[DRA]:https://kubernetes.io/docs/concepts/scheduling-eviction/dynamic-resource-allocation/
[CNI]:https://www.cni.dev/
[cni-dra-driver]:https://github.com/kubernetes-sigs/cni-dra-driver

3. Did you have community-wide updates in 2024 (e.g. KubeCon talks)?

Blog Posts:

- [Getting started with the AdminNetworkPolicy](https://network-policy-api.sigs.k8s.io/blog/2024/01/30/getting-started-with-the-adminnetworkpolicy-api/)
- [Gateway API v1.1](https://kubernetes.io/blog/2024/05/09/gateway-api-v1-1/)
- [Gateway API v1.2](https://kubernetes.io/blog/2024/11/21/gateway-api-v1-2/)

Talks:

- [Network Policy: The future of Network Policy (Kubecon EU)](https://sched.co/1Yhhb)
- [How many network policies can I create? (Kubecon EU)](https://kccnceu2024.sched.com/event/1YeMI)
- [Gateway API "Beyond GA" (Kubecon EU)](https://kccnceu2024.sched.com/event/50c69a019a6fb2935d5f19cbb30f0ce3)
- [Configuring your Service Mesh with Gateway API (Kubecon EU)](https://kccnceu2024.sched.com/event/1YeOL)
- [SIG Network Intro & Updates (Kubecon NA)](https://kccncna2024.sched.com/event/1hovJ)
- [Securing the Future of Ingress-NGINX (Kubecon NA)](https://kccncna2024.sched.com/event/1hoxW)
- [Gateway API: What's new, what's next? (Kubecon NA)](https://kccncna2024.sched.com/event/1hoxF)
- [Unlocking the Gateway: A Practical Guide from Ingress to Gateway API (Kubecon NA)](https://cfp.cloud-native.rejekts.io/cloud-native-rejekts-eu-paris-2024/talk/KEYQD9/)
- [How to move from Ingress to Gateway API with minimal Hassle (Kubecon NA)](https://kccncna2024.sched.com/event/1i7ng)
- [Live with Gateway API 1.2 (Kubecon NA)](https://kccncna2024.sched.com/event/1i7np)
- [With Great Flexibility Comes Great Complexity: Inspect your Gateway API Configuration (Kubecon NA)](https://kccncna2024.sched.com/event/1i7s2)
- [CNI Updates & Direction (Kubecon NA)](https://kccncna2024.sched.com/event/1how8)

Webinars:

- [Migrating from Ingress to Gateway API (CNCF Webinar)](https://community.cncf.io/events/details/cncf-cncf-online-programs-presents-cncf-on-demand-webinar-ingress2gateway-migrating-from-ingress-to-gateway-api/)

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
- [x] [README.md] reviewed for accuracy and updated if needed
- [ ] [CONTRIBUTING.md] reviewed for accuracy and updated if needed
- [x] Other contributing docs (e.g. in devel dir or contributor guide) reviewed for accuracy and updated if needed
- [x] Subprojects list and linked OWNERS files in [sigs.yaml] reviewed for accuracy and updated if needed
- [x] SIG leaders (chairs, tech leads, and subproject leads) in [sigs.yaml] are accurate and active, and updated if needed
- [x] Meeting notes and recordings for 2024 are linked from [README.md] and updated/uploaded if needed


[CONTRIBUTING.md]: https://git.k8s.io/community/sig-network/CONTRIBUTING.md
[sig-governance.md]: https://git.k8s.io/community/committee-steering/governance/sig-governance.md
[README.md]: https://git.k8s.io/community/sig-network/README.md
[sigs.yaml]: https://git.k8s.io/community/sigs.yaml
[devel]: https://git.k8s.io/community/contributors/devel/README.md
