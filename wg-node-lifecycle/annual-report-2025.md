# 2025 Annual Report: WG Node Lifecycle

## Current initiatives and Project Health

1. What work did the WG do this year that should be highlighted?

WG Node Lifecycle was formed approximately 6 months ago to explore and improve
node and pod lifecycle in Kubernetes. Since its formation, the group has been
ramping up with weekly meetings and active discussions across
multiple SIGs.

The group presented at KubeCon NA Maintainer Summit, highlighting the two
primary areas of focus - **Pod Eviction** and **Node Maintenance**:

- [KubeCon NA Maintainer Summit Talk](https://youtu.be/yh3tkE1fw-4?si=Duk-QUwpNvNU4kin)

Key initiatives the WG is driving:

**Targeting Kubernetes 1.36**
- [EvictionRequest API](https://github.com/kubernetes/enhancements/issues/4563) — Introduce EvictionRequest API (fka Evacuation) to allow managed graceful pod termination.

**Targeting Kubernetes 1.37**
- [Specialized Lifecycle Management](https://github.com/kubernetes/enhancements/issues/5683) — A new KEP replacing [Declarative Node Maintenance](https://github.com/kubernetes/enhancements/issues/4212) with a broader scope. The WG had been discussing SLM for several weeks in this [google doc](https://docs.google.com/document/d/1rY3s_cGIaz4-mTwQGQL1LD5udr2cwIvwFuWAh_pd0MI/edit?usp=sharing), before the KEP was published.

The group has been actively discussing these lifecycle ideas across stakeholder SIGs (Apps,
Autoscaling, CLI, Cloud Provider, Cluster Lifecycle, Node, and Scheduling) to gather requirements and align on design.

2. Are there any areas and/or subprojects that your group needs help with (e.g. fewer than 2 active OWNERS)?

No, we have no subprojects at the moment. Our focus is on our KEPs.

## Operational

Operational tasks in [wg-governance.md]:

- [x] [README.md] reviewed for accuracy and updated if needed
- [x] WG leaders in [sigs.yaml] are accurate and active, and updated if needed
- [x] Meeting notes and recordings for 2025 are linked from [README.md] and updated/uploaded if needed
- [x] Updates provided to sponsoring SIGs in 2025
      - KubeCon NA 2025 Maintainer Summit talk:
        [WG Node Lifecycle - Getting Started](https://youtu.be/yh3tkE1fw-4?si=Duk-QUwpNvNU4kin)

[wg-governance.md]: https://git.k8s.io/community/committee-steering/governance/wg-governance.md
[README.md]: https://git.k8s.io/community/wg-node-lifecycle/README.md
[sigs.yaml]: https://git.k8s.io/community/sigs.yaml
