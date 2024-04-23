# 2023 Annual Report: SIG Network

## Current initiatives and Project Health

1. What work did the SIG do this year that should be highlighted?

  - Governance and leadership changes

    - Casey Davenport (**[@caseydavenport](https://github.com/caseydavenport)**)
      and Dan Williams (**[@dcbw](https://github.com/dcbw)**) stepped
      down as Chairs, Shane Utt (**[@shaneutt](https://github.com/shaneutt)**) and
      Mike Zappa (**[@mikezappa87](https://github.com/mikezappa87)**) became new
      Chairs (joining Tim Hockin (**[@thockin](https://github.com/thockin)**), who
      remained a Chair)

    - Dan Winship (**[@danwinship](https://github.com/danwinship)**) and
      Antonio Ojea (**[@aojea](https://github.com/aojea)**) became SIG Network's
      first Tech Leads.

  - Gateway API

    The [Gateway API reached v1.0 and is now GA!]. Specifically, the
    `Gateway`, `GatewayClass`, and `HTTPRoute` APIs are now v1, with
    several other APIs also being added or updated in the experimental
    channel. Also, the old validation webhook is now deprecated in
    favor of CEL-based validation. See [the v1.0.0 release notes] for
    full details.

[Gateway API reached v1.0 and is now GA!]: https://kubernetes.io/blog/2023/10/31/gateway-api-ga/
[the v1.0.0 release notes]: https://github.com/kubernetes-sigs/gateway-api/releases/tag/v1.0.0

  - NetworkPolicy API Working Group

    The Network Policy API WG mainly focused on the maintenance and
    new feature development regarding our two major APIs,
    AdminNetworkPolicy and BaselineAdminNetworkPolicy. This also
    included a few talks at KubeCon NA 2023 in Chicago:

      - [Network Policy API: Intro and Project Update]
      - [AdminNetworkPolicy: A New Kubernetes-Native API for Comprehensive Cluster-Wide Network Security]

    Additionally, major features such as [Egress Traffic Control],
    [FQDN Selectors] and [Network Tenancy] all made great progress
    during the year and we are excited to deliver them here in 2024.

    On top of the APIs and Features, we introduced the
    [Policy-Assistant] tool which allows users to have a better overview
    of how all of the native Kubernetes APIs interact on real
    clusters.

[Network Policy API: Intro and Project Update]: https://youtu.be/lYWW3KogPTg?si=c6vpOwWMp5D6GpMM
[AdminNetworkPolicy: A New Kubernetes-Native API for Comprehensive Cluster-Wide Network Security]: https://youtu.be/DTxvTCISi7Q?si=pz8RN35ptn8eFeE5
[Egress Traffic Control]: https://network-policy-api.sigs.k8s.io/npeps/npep-126-egress-traffic-control/
[FQDN Selectors]: https://network-policy-api.sigs.k8s.io/npeps/npep-133/
[Network Tenancy]: https://network-policy-api.sigs.k8s.io/npeps/npep-122/
[Policy-Assistant]: https://github.com/kubernetes-sigs/network-policy-api/tree/main/cmd/policy-assistant

2. Are there any areas and/or subprojects that your group needs help with (e.g. fewer than 2 active OWNERS)?

The Network Policy API group is always looking for more people to get
involved. Whether that's trying out our APIs, updating docs, or
helping our with NPEP design, we are hugely appreciative of all help.
Please checkout [our website](https://network-policy-api.sigs.k8s.io/)
for even more information.

Historically, SIG Network as a whole has not been great at onboarding
new contributors, but contributions are definitely welcome. (There has
recently been a flurry of [new-contributor effort around the new
nftables kube-proxy backend].)

We also continue to have problems with Windows networking, where most
of the people who regularly attend SIG Network meetings know very
little about Windows (and most of the people who regularly attend SIG
Windows meetings know very little about networking). If you know a lot
about Windows container networking (or want to learn) and want to get
involved, it would be greatly appreciated.

[new-contributor effort around the new nftables kube-proxy backend]: https://github.com/kubernetes/kubernetes/issues/122572

3. Did you have community-wide updates in 2023 (e.g. KubeCon talks)?

We gave a SIG update at KubeCon EU: [SIG Network: Intro and Updates].

[SIG Network: Intro and Updates]: https://www.youtube.com/watch?v=0uPEFcWn-_o

4. KEP work in 2023 (v1.27, v1.28, v1.29):

  - Alpha
    - [1880 - Multiple Service CIDRs](https://github.com/kubernetes/enhancements/tree/master/keps/sig-network/1880-multiple-service-cidrs) - v1.29
    - [3866 - Add an nftables-based kube-proxy backend](https://github.com/kubernetes/enhancements/tree/master/keps/sig-network/3866-nftables-proxy) - v1.29
    - [4004 - Deprecate status.nodeInfo.kubeProxyVersion field](https://github.com/kubernetes/enhancements/tree/master/keps/sig-network/4004-deprecate-kube-proxy-version) - v1.29

  - Beta
    - [1860 - Make Kubernetes aware of the load balancer behaviour](https://github.com/kubernetes/enhancements/tree/master/keps/sig-network/1860-kube-proxy-IP-node-binding) - v1.30
    - [3836 - Kube-proxy improved ingress connectivity reliability](https://github.com/kubernetes/enhancements/tree/master/keps/sig-network/3836-kube-proxy-improved-ingress-connectivity-reliability) - v1.30

  - Stable
    - [1669 - Proxy Terminating Endpoints](https://github.com/kubernetes/enhancements/tree/master/keps/sig-network/1669-proxy-terminating-endpoints) - v1.28
    - [2595 - Expanded DNS Configuration](https://github.com/kubernetes/enhancements/tree/master/keps/sig-network/2595-expanded-dns-config) - v1.28
    - [2681 - Field status.hostIPs added for Pod](https://github.com/kubernetes/enhancements/tree/master/keps/sig-network/2681-pod-host-ip) - v1.30
    - [3178 - Cleaning up IPTables Chain Ownership](https://github.com/kubernetes/enhancements/tree/master/keps/sig-network/3178-iptables-cleanup) - v1.28
    - [3453 - Minimize iptables-restore input size](https://github.com/kubernetes/enhancements/tree/master/keps/sig-network/3453-minimize-iptables-restore) - v1.28
    - [3458 - Remove transient node predicates from KCCM's service controller](https://github.com/kubernetes/enhancements/tree/master/keps/sig-network/3458-remove-transient-node-predicates-from-service-controller) - v1.30
    - [3668 - Reserve Nodeport Ranges For Dynamic And Static Port Allocation](https://github.com/kubernetes/enhancements/tree/master/keps/sig-network/3668-reserved-service-nodeport-range) - v1.29
    - [3705 - Cloud Dual-Stack --node-ip Handling](https://github.com/kubernetes/enhancements/tree/master/keps/sig-network/3705-cloud-node-ips) - v1.30
    - [3726 - standard-application-protocols](https://github.com/kubernetes/enhancements/tree/master/keps/sig-network/3726-standard-application-protocols) - v1.27

## [Subprojects](https://git.k8s.io/community/sig-network#subprojects)

(This list is auto-generated and has reminded us that we should
probably update our list of subprojects...)

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

(This list is also auto-generated and is even less accurate than the
above list...)

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
