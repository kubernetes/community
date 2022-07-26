# 2021 Annual Report: SIG Network

## Current initiatives

1. What work did the SIG do this year that should be highlighted?

- [Dual-stack IPv4/IPv6 networking feature reaches GA](https://kubernetes.io/blog/2021/12/08/dual-stack-networking-ga/) 
- [KPNG](https://github.com/kubernetes-sigs/kpng)
- [AdminNetworkPolicy KEP](https://github.com/kubernetes/enhancements/tree/master/keps/sig-network/2091-admin-network-policy)
- [Gateway API](https://gateway-api.sigs.k8s.io/)

2. What initiatives are you working on that aren't being tracked in KEPs?

- [Gateway API](https://gateway-api.sigs.k8s.io/) has defined its own [GEP process](https://gateway-api.sigs.k8s.io/contributing/gep/) for tracking enhancements. See below for some 2021 highlights. The group is currently working towards its upcoming [v1beta1 milestone](https://github.com/kubernetes-sigs/gateway-api/issues?q=is%3Aopen+is%3Aissue+milestone%3Av1beta1).

3. KEP work in 2021:

   - Stable

     - [1797 - Configure FQDN as Hostname for Pods](https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/1797-configure-fqdn-as-hostname-for-pods)
     - [0752 - EndpointSlices](https://github.com/kubernetes/enhancements/tree/master/keps/sig-network/0752-endpointslices)
     - [2365 - IngressClass Namespaced Params](https://github.com/kubernetes/enhancements/tree/master/keps/sig-network/2365-ingressclass-namespaced-params)
     - [563 - Dual stack](https://github.com/kubernetes/enhancements/tree/master/keps/sig-network/563-dual-stack)

   - Beta

     - [2079 - NetworkPolicy port ranges](https://github.com/kubernetes/enhancements/tree/master/keps/sig-network/2079-network-policy-port-range) 
     - [2086 - Service internal traffic policy](https://github.com/kubernetes/enhancements/tree/master/keps/sig-network/2086-service-internal-traffic-policy) 
     - [2433 - Topology aware hints](https://github.com/kubernetes/enhancements/tree/master/keps/sig-network/2433-topology-aware-hints) 
     - [1669 - Proxy terminating endpoints](https://github.com/kubernetes/enhancements/tree/master/keps/sig-network/1669-proxy-terminating-endpoints) 

   - Alpha

     - [1435 - Mixed protocol LB](https://github.com/kubernetes/enhancements/tree/master/keps/sig-network/1435-mixed-protocol-lb)
     - [2595 - Expanded DNS config](https://github.com/kubernetes/enhancements/tree/master/keps/sig-network/2595-expanded-dns-config)
Gateway API Enhancements:

* v1alpha2 highlights:
    * Transitioned to new `gateway.networking.k8s.io` API group to reflect its status as an official Kubernetes API.
    * [GEP-724](https://gateway-api.sigs.k8s.io/geps/gep-724/): Simpler Route-Gateway binding.
    * [GEP-709](https://gateway-api.sigs.k8s.io/geps/gep-709/): Safe cross namespace references.
    * [GEP-713](https://gateway-api.sigs.k8s.io/geps/gep-713/): Policy attachment.
* Other notable initiatives:
    * [GEP-917](https://gateway-api.sigs.k8s.io/geps/gep-917/): Gateway API conformance testing.
    * [GEP-922](https://gateway-api.sigs.k8s.io/geps/gep-922/): Gateway API conformance versioning.
    * [Validating admission webhook](https://github.com/kubernetes-sigs/gateway-api/issues/487).
## Project health

1. What areas and/or subprojects does your group need the most help with? Any areas with 2 or fewer OWNERs? (link to more details)

- [kube-dns](https://github.com/kubernetes/dns/blob/master/OWNERS)
- [external-dns](https://github.com/kubernetes-sigs/external-dns/blob/master/OWNERS)
- [ingress-nginx](https://github.com/kubernetes/ingress-nginx/blob/master/OWNERS)
  - A good community was assembled, and now we have folks (including F5) giving some support in Slack and issues.
  - Bi weekly meetings for issue and PRs prioritization, with a regular attendance of 6 participants.
  - The development cadence is slow. Only 2 or 3 developers are actually maintaining the code, which brings the main concern of GatewayAPI being released and Ingress NGINX (which is an official subproject) not supporting it.
  - Three CVEs has been discovered in the last year, and they have been fixed raising the alert on the need to actually split control plane and data plane. We are planning do a first split to make it easier to support GatewayAPI, and later join forces with KPNG to use similar approach in Ingress NGINX

2. What metrics/community health stats does your group care about and/or measure?

Not currently monitoring metrics or health statistics. 

3. Does your [CONTRIBUTING.md] help **new** contributors engage with your group specifically by pointing to activities or programs that provide useful context or allow easy participation?

We do not have a SIG specific CONTRIBUTING.md

4. If your group has special training, requirements for reviewers/approvers, or processes beyond the general [contributor guide], does your [CONTRIBUTING.md] document those to help **existing** contributors grow throughout the [contributor ladder]?

There are no special requirements.

5. Does the group have contributors from multiple companies/affiliations?

Yes - we have long-term as well as one-time contributors from a variety of companies.

6. Are there ways end users/companies can contribute that they currently are not? If one of those ways is more full time support, what would they work on and why?

Could always use additional code and KEP reviewers, and maintainers for understaffed subprojects as mentioned above.

## Membership

- Primary slack channel member count: 6667
- Primary mailing list member count: 1181
- Primary meeting attendee count (estimated, if needed): 20-30
- Primary meeting participant count (estimated, if needed): 5-10
- Unique reviewers for SIG-owned packages: 29
- Unique approvers for SIG-owned packages: 29

## Subprojects

New in $2021:

- [kpng](https://git.k8s.io/community/sig-network#kpng)
- [gateway-api](https://git.k8s.io/community/sig-network#gateway-api)

Continuing:

- [cluster-proportional-vertical-autoscaler](https://git.k8s.io/community/sig-network#cluster-proportional-autoscaler)
- [external-dns](https://git.k8s.io/community/sig-network#external-dns)
- [ingress](https://git.k8s.io/community/sig-network#ingress)
- [iptables-wrappers](https://git.k8s.io/community/sig-network#iptables-wrappers)
- [kube-dns](https://git.k8s.io/community/sig-network#kube-dnss)
- [network-policy](https://git.k8s.io/community/sig-network#network-policy)
- [pod-networking](https://git.k8s.io/community/sig-network#pod-networking)

## Working groups

New in 2021:

- [WG IoT edge](https://git.k8s.io/community/wg-iot-edge/) ([2021 report](https://git.k8s.io/community/wg-iot-edge/annual-report-2021.md))
- [WG Multitenancy](https://git.k8s.io/community/wg-multitenancy/) ([2021 report](https://git.k8s.io/community/wg-multitenancy/annual-report-2021.md))
- [WG Policy](https://git.k8s.io/community/wg-policy/) ([2021 report](https://git.k8s.io/community/wg-policy/annual-report-2021.md))
- [WG Structured Logging](https://git.k8s.io/community/wg-structured-logging/) ([2021 report](https://git.k8s.io/community/wg-structured-logging/annual-report-2021.md))

## Operational

Operational tasks in [sig-governance.md]:

- [X] [README.md] reviewed for accuracy and updated if needed
- [X] [CONTRIBUTING.md] reviewed for accuracy and updated if needed (or created if missing and your contributor steps and experience are different or more in-depth than the documentation listed in the general [contributor guide] and [devel] folder.)
- [] Subprojects list and linked OWNERS files in [sigs.yaml] reviewed for accuracy and updated if needed
- [X] SIG leaders (chairs, tech leads, and subproject owners) in [sigs.yaml] are accurate and active, and updated if needed
- [X] Meeting notes and recordings for 2021 are linked from [README.md] and updated/uploaded if needed
- [X] Did you have community-wide updates in 2021 (e.g. community meetings, kubecon, or kubernetes-dev@ emails)? Links to email, slides, or recordings:
	- [KubeCon NA 2021 - update recording](https://www.youtube.com/watch?v=uZ0WLxpmBbY)
	- [KubeCon EU 2021 - update recording](https://www.youtube.com/watch?v=Nn-qrp0TRnM)
