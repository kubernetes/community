# 2025 Annual Report: SIG Network

## Current initiatives and Project Health

[DRANET]: https://github.com/kubernetes-sigs/dranet
[Gateway API Inference Extensions]: https://github.com/kubernetes-sigs/gateway-api-inference-extension
[Gateway API]: https://gateway-api.sigs.k8s.io/
[Network Policy]: https://github.com/kubernetes-sigs/network-policy-api
[External DNS]: https://github.com/kubernetes-sigs/external-dns
[IP Masq Agent]: https://github.com/kubernetes-sigs/ip-masq-agent
[Multi-Network]: https://github.com/kubernetes-sigs/multi-network-api
[Ingress-NGINX]: https://github.com/kubernetes/ingress-nginx
[InGate]: https://github.com/kubernetes-sigs/ingate
[Blixt]: https://github.com/kubernetes-retired/blixt
[Cluster Proportional Autoscaler]: https://github.com/kubernetes-sigs/cluster-proportional-autoscaler


1. What work did the SIG do this year that should be highlighted?

- [DRANET]
  - DRANET is a Kubernetes Network Driver that uses Dynamic Resource Allocation
  (DRA) to deliver high-performance networking for demanding applications in
  Kubernetes.
  - DRANET has reached [v1.0.1](https://github.com/kubernetes-sigs/dranet/releases/tag/v1.0.1)

- [Gateway API Inference Extensions (GIE)](https://github.com/kubernetes-sigs/gateway-api-inference-extension)
  - Reached v1.0 milestone!
  - Shipped important features: LoRA Syncer, Flow Control, Standalone EPP
  - Latest release [v1.3.1](https://github.com/kubernetes-sigs/gateway-api-inference-extension/releases/tag/v1.3.1)

- [Gateway API]
  - Gateway API is now moving to [monthly experimental releases with a four-month cadence for Standard channel](https://gateway-api.sigs.k8s.io/concepts/versioning/).
  - Released [v1.4](https://github.com/kubernetes-sigs/gateway-api/releases/tag/v1.4.0)
  - Released [v1.3](https://github.com/kubernetes-sigs/gateway-api/releases/tag/v1.3.0)

- [Network Policy]
  - Work continues on finalizing a cluster administrator oriented API:
    - APIs `AdminNetworkPolicy` and `BaselineNetworkPolicy` are now combined
      into a single `ClusterNetworkPolicy` resource.
    - Finalized Beta candidate API for `ClusterNetworkPolicy`.
    - Three working ecosystem implementations of the API

- [External DNS]
  - Shipped five releases in 2025: v0.16, v0.17, v0.18, v0.19, v0.20
  - Latest release [v0.20.0](https://github.com/kubernetes-sigs/external-dns/releases/tag/v0.20.0)

- [IP Masq Agent]
  - Release [v2.12.6](https://github.com/kubernetes-sigs/ip-masq-agent/releases/tag/v2.12.6)

- [Cluster Proportional Autoscaler]
  - Release [v1.10.2](https://github.com/kubernetes-sigs/cluster-proportional-autoscaler/releases/tag/v1.10.2)

- [Multi-Network]
  - Work continues on the API definition using DRA as the integration point.

- [Ingress-NGINX]
  - Project is being [retired][nginx-retirement] due to challenges with security
    vulnerabilities and sufficient maintainership.
  - By the end of March 2026, maintenance will be halted, and the project will
    be retired.  Existing deployments of Ingress NGINX will not be
    broken. Existing project artifacts such as Helm charts and container images
    will remain available.
[nginx-retirement]: https://kubernetes.io/blog/2025/11/11/ingress-nginx-retirement/

- [InGate]
  - InGate is being retired (early 2026). SIG Network and the Security Response
    Committee recommend that all users begin migration to Gateway API or another
    Ingress controller immediately.

- [Blixt]
  - Blixt was an experimental load balancer using eBPF and written in Rust. The
    project was an experimental sandbox and had no plans to ever be
    a shipping product. It has now been retired.

2. Are there any areas and/or subprojects that your group needs help with (e.g. fewer than 2 active OWNERS)?

While there are no SIG Network projects which _couldn't_ benefit from more help
and contributions, the following is a list of specific projects where there are
known blockers (or otherwise critical needs) where more contributions, feedback,
or implementation support is the key to "unblocking" them:

- [Network Policy] APIs can use end user feedback to help us graduate towards Beta and GA.

- [Multi-Network] would like to refine their set of use cases with feedback from users and
  the community. There are pieces of a reference implementation of the multi-network API that would
  benefit from more contributors.

- [Gateway API]
  - has a large number of features which are not yet standard which could use implementations
    to join us to help us graduate those towards GA.
  - has two CLI utilities which need more users and implementations to get involved to
    provide feedback: [ingress2gateway], [gwctl]
  - Agentic and AI workstreams would benefit from feedback and implementations from the community
    to validate design decisions.

[ingress2gateway]: https://github.com/kubernetes-sigs/ingress2gateway
[gwctl]: https://github.com/kubernetes-sigs/gateway-api/tree/main/gwctl

- [DRANET]
  - DRANET welcomes additional implementations and feedback from the community for different high-performance
    networking drivers.

3. Did you have community-wide updates in 2025 (e.g. KubeCon talks)?

Blog posts:

- [NFTables mode for kube-proxy](https://kubernetes.io/blog/2025/02/28/nftables-kube-proxy/)
- [Ingress-nginx CVE-2025-1974: What You Need to Know](https://kubernetes.io/blog/2025/03/24/ingress-nginx-cve-2025-1974/)
- [Endpoints Deprecation](https://kubernetes.io/blog/2025/04/24/endpoints-deprecation/)
- [Introducing Gateway API Inference Extension](https://kubernetes.io/blog/2025/06/05/introducing-gateway-api-inferenc- [Ingress NGINX Retirement: What You Need to Know](https://kubernetes.io/blog/2025/11/11/ingress-nginx-retirement/)
- [Gateway API 1.4: New Features](https://kubernetes.io/blog/2025/11/06/gateway-api-v1-4/)
e-extension/)

Talks:

- Kubecon EU 2025
  - [Making the Leap: What Gateway API Needs To Support Ingress-NGINX Users - Rob Scott, Google & James Strong, Isovalent at Cisco](https://kccnceu2025.sched.com/event/1txAl/making-the-leap-what-gateway-api-needs-to-support-ingress-nginx-users-rob-scott-google-james-strong-isovalent-at-cisco)
  - [Taming the Traffic: Selecting the Perfect Gateway Implementation for You](https://kccnceu2025.sched.com/event/1txAr/taming-the-traffic-selecting-the-perfect-gateway-implementation-for-you-spencer-hance-google-arko-dasgupta-tetrate-christine-kim-isovalent-at-cisco-kate-osborn-nginxf5-mike-morris-microsoft)
  - [How To Gateway With Ingress - 140 Days InGate](https://kccnceu2025.sched.com/event/1tcyc/how-to-gateway-with-ingress-140-days-ingate-marco-ebert-giant-swarm-james-strong-isovalent-at-cisco)
  - [Keynote: LLM-Aware Load Balancing in Kubernetes: A New Era of Efficiency](https://kccnceu2025.sched.com/event/1txC7/keynote-llm-aware-load-balancing-in-kubernetes-a-new-era-of-efficiency-clayton-coleman-distinguished-engineer-google-jiaxin-shan-software-engineer-bytedance)
  - [SIG Network Intro and Updates](https://kccnceu2025.sched.com/event/1tczU/sig-network-intro-and-updates-dan-winship-nadia-pinaeva-red-hat-bowei-du-google-daman-arora-broadcom)
  - [Uncharted Waters: Dynamic Resource Allocation for Networking](https://kccnceu2025.sched.com/event/1txAx/uncharted-waters-dynamic-resource-allocation-for-networking-miguel-duarte-barroso-red-hat-lionel-jouin-ericsson-software-technology)

- Kubecon NA 2025
  - [Keynote: The Community-Driven Evolution of the Kubernetes Network Driver](https://kccncna2025.sched.com/event/27FYh/keynote-the-community-driven-evolution-of-the-kubernetes-network-driver-lionel-jouin-software-engineer-red-hat-antonio-ojea-staff-software-engineer-google)
  - [AdminNetworkPolicy: From Alpha To Beta and Beyond](https://kccncna2025.sched.com/event/27Nnu/adminnetworkpolicy-from-alpha-to-beta-and-beyond-dan-winship-surya-seetharaman-red-hat-nadia-pinaeva-nvidia-bowei-du-google)
  - [AI Inference Without Boundaries: Dynamic Routing With Multi-Cluster Inference Gateway](https://kccncna2025.sched.com/event/27FeP/ai-inference-without-boundaries-dynamic-routing-with-multi-cluster-inference-gateway-rob-scott-google-daneyon-hansen-soloio)
  - [Gateway API: Table Stakes](https://kccncna2025.sched.com/event/27No3/gateway-api-table-stakes-shane-utt-candace-holman-red-hat-mike-morris-microsoft-lior-lieberman-kellen-swain-google)

4. KEP work in 2025 (v1.33, v1.34, v1.35):
  - v1.33
    - Alpha
      - [4858 - IP/CIDR Validation Improvements](https://github.com/kubernetes/enhancements/tree/master/keps/sig-network/4858-ip-cidr-validation)
    - Stable
      - [1880 - Multiple Service CIDRs](https://github.com/kubernetes/enhancements/tree/master/keps/sig-network/1880-multiple-service-cidrs)
      - [2433 - Topology Aware Hints](https://github.com/kubernetes/enhancements/tree/master/keps/sig-network/2433-topology-aware-hints)
      - [3866 - Add an nftables-based kube-proxy backend](https://github.com/kubernetes/enhancements/tree/master/keps/sig-network/3866-nftables-proxy)
      - [4444 - Traffic Distribution for Services](https://github.com/kubernetes/enhancements/tree/master/keps/sig-network/4444-service-traffic-distribution)

  - v1.34
    - Alpha
      - [5311 - Relaxed validation for Services names](https://github.com/kubernetes/enhancements/tree/master/keps/sig-network/5311-relaxed-validation-for-service-names)
    - Beta
    - Stable
      - [4427 - Relaxed DNS search string validation](https://github.com/kubernetes/enhancements/tree/master/keps/sig-network/4427-relaxed-dns-search-validation)
  - v1.35
    - Alpha
    - Beta
      - [4762 - Allows setting arbitrary FQDN as the pod's hostname](https://github.com/kubernetes/enhancements/tree/master/keps/sig-network/4762-allow-arbitrary-fqdn-as-pod-hostname)
    - Stable
      - [3015 - PreferSameZone and PreferSameNode Traffic Distribution](https://github.com/kubernetes/enhancements/tree/master/keps/sig-network/3015-prefer-same-node)


## [Subprojects](https://git.k8s.io/community/sig-network#subprojects)


**New in 2025:**
  - kindnet
  - kube-agentic-networking
  - kubernetes-network-drivers
  - wg-ai-gateway

**Continuing:**
  - cluster-proportional-autoscaler
  - cluster-proportional-vertical-autoscaler
  - external-dns
  - gateway-api
  - gateway-api-inference-extension
  - ingate
  - ingress
  - iptables-wrappers
  - knftables
  - kube-dns
  - multi-network
  - network-policy
  - node-ipam-controller
  - pod-networking

## [Working groups](https://git.k8s.io/community/sig-network#working-groups)

**New in 2025:**
 - AI Gateway
 - Node Lifecycle

**Retired in 2025:**
 - Policy

**Continuing:**
 - Device Management
 - Serving
 - Structured Logging

## Operational

Operational tasks in [sig-governance.md]:
- [x] [README.md] reviewed for accuracy and updated if needed
- [x] [CONTRIBUTING.md] reviewed for accuracy and updated if needed
- [x] Other contributing docs (e.g. in devel dir or contributor guide) reviewed for accuracy and updated if needed
- [x] Subprojects list and linked OWNERS files in [sigs.yaml] reviewed for accuracy and updated if needed
- [x] SIG leaders (chairs, tech leads, and subproject leads) in [sigs.yaml] are accurate and active, and updated if needed
- [x] Meeting notes and recordings for 2025 are linked from [README.md] and updated/uploaded if needed


[CONTRIBUTING.md]: https://git.k8s.io/community/sig-network/CONTRIBUTING.md
[sig-governance.md]: https://git.k8s.io/community/committee-steering/governance/sig-governance.md
[README.md]: https://git.k8s.io/community/sig-network/README.md
[sigs.yaml]: https://git.k8s.io/community/sigs.yaml
[devel]: https://git.k8s.io/community/contributors/devel/README.md
