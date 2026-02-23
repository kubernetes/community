# 2025 Annual Report: WG AI Gateway

## Current initiatives and Project Health

1. What work did the WG do this year that should be highlighted?

- WG formally established in September 2025; [charter](https://git.k8s.io/community/wg-ai-gateway/charter.md) ratified, [proposals repo](https://github.com/kubernetes-sigs/wg-ai-gateway) created, weekly meetings started
- Two proposals developed:
  - [Payload Processing](https://github.com/kubernetes-sigs/wg-ai-gateway/blob/main/proposals/7-payload-processing.md): prompt guards, semantic routing/caching, token rate limiting, and how these map onto Gateway API filters
  - [Egress Gateways](https://github.com/kubernetes-sigs/wg-ai-gateway/blob/main/proposals/10-egress-gateways.md): proxying to third-party AI services (OpenAI, Gemini, Claude, etc.), with a comparison of prior art and a resource model built on Gateway API
- KubeCon EU 2026 session accepted: "Aim at the Gate: Introducing the AI Gateway Working Group in Kubernetes"

2. Are there any areas and/or subprojects that your group needs help with (e.g. fewer than 2 active OWNERS)?

No. Our deliverables are proposals for upstream sigs, which have active contributors.

## Operational

Operational tasks in [wg-governance.md]:

- [x] [README.md] reviewed for accuracy and updated if needed
- [ ] WG leaders in [sigs.yaml] are accurate and active, and updated if needed
  - Leadership updates in progress: [kubernetes/community#8840](https://github.com/kubernetes/community/pull/8840)
- [x] Meeting notes and recordings for 2025 are linked from [README.md] and updated/uploaded if needed
- [x] Updates provided to sponsoring SIGs in 2025
      - [SIG Network](https://git.k8s.io/community/sig-network/)
        - Proposals developed in coordination with SIG Network maintainers; [Gateway API draft GEP](https://github.com/kubernetes-sigs/gateway-api/pull/4488) opened for the Backend resource
      - [SIG Multicluster](https://git.k8s.io/community/sig-multicluster/)
        - Multi-cluster inference routing is in-scope per the charter but detailed proposals have not yet been developed

[wg-governance.md]: https://git.k8s.io/community/committee-steering/governance/wg-governance.md
[README.md]: https://git.k8s.io/community/wg-ai-gateway/README.md
[sigs.yaml]: https://git.k8s.io/community/sigs.yaml
