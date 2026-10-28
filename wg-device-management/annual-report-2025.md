# 2025 Annual Report: WG Device Management

## Current initiatives and Project Health


1. What work did the WG do this year that should be highlighted?

The WG continues to function with well-attended Zoom meetings as well as all hands meetings and talks
at KubeCon. The [core of Dynamic Resource Allocation](https://github.com/kubernetes/enhancements/issues/4381)
got promoted to GA in Kubernetes 1.34. Many new fetures were introduced in alpha:
  1. [Consumable Capacity](https://github.com/kubernetes/enhancements/issues/5075)
  2. [Device Binding Conditions](https://github.com/kubernetes/enhancements/issues/5007)
  3. [Device Taints and Tolerations](https://github.com/kubernetes/enhancements/issues/5055)
  4. [Extended Resource Requests via DRA](https://github.com/kubernetes/enhancements/issues/5004)
  5. [Partitionable Devices](https://github.com/kubernetes/enhancements/issues/4815)
  6. [Resource Health Status in Pod Status](https://github.com/kubernetes/enhancements/issues/4680)

Others were introduced in alpha and graduated to beta:
  1. [Namespace Controlled Admin Access](https://github.com/kubernetes/enhancements/issues/5018)
  2. [Prioritized Alternatives in Device Requests](https://github.com/kubernetes/enhancements/issues/4816)
  3. [Resource Claim Status](https://github.com/kubernetes/enhancements/issues/4817)

Many new contributors have ramped up and helped with driving the features and issues as well as promoting
DRA to GA.

KubeCon Updates:
  - KubeCon NA maintainer track talk:
    - ["DRA is GA! Kubernetes WG Device Management - GPUs, TPUs, NICs and More..."](https://www.youtube.com/watch?v=Op4DNDTij1U)
  - KubeCon EU maintainer track talk:
    - ["Kubernetes WG Device Management - GPUs, TPUs, NICs and More With DRA"](https://www.youtube.com/watch?v=Z_15EyXOnhU)

2. Are there any areas and/or subprojects that your group needs help with (e.g. fewer than 2 active OWNERS)?

The [dra-example-driver](https://github.com/kubernetes-sigs/dra-example-driver) has a single active maintainer,
more help would be welcome especially around contributing the expertise from production DRA drivers back into
the example driver.

## Operational

Operational tasks in [wg-governance.md]:

- [x] [README.md] reviewed for accuracy and updated if needed
- [x] WG leaders in [sigs.yaml] are accurate and active, and updated if needed
- [x] Meeting notes and recordings for 2025 are linked from [README.md] and updated/uploaded if needed
- [x] Updates provided to sponsoring SIGs in 2025
      - [sig-node](https://git.k8s.io/community/sig-node/)
          - [Mailing List](https://groups.google.com/a/kubernetes.io/g/sig-node)
          - [Meeting Agenda](https://docs.google.com/document/d/1Ne57gvidMEWXR70OxxnRkYquAoMpt56o75oZtg-OeBg/edit?tab=t.0#heading=h.yug5jtgi5nki)
          - [Recordings](https://www.youtube.com/playlist?list=PL69nYSiGNLP1wJPj5DYWXjiArF-MJ5fNG)
      - [sig-scheduling](https://git.k8s.io/community/sig-scheduling/)
          - [Mailing List](https://groups.google.com/a/kubernetes.io/g/sig-scheduling)
          - [Meeting Agenda](https://docs.google.com/document/d/13mwye7nvrmV11q9_Eg77z-1w3X7Q1GTbslpml4J7F3A/edit?tab=t.0#heading=h.6qiwue4qwtpi)
          - [Recordings](https://www.youtube.com/watch?v=PweKj6SU7UA&list=PL69nYSiGNLP2vwzcCOhxrL3JVBc-eaJWI)

[wg-governance.md]: https://git.k8s.io/community/committee-steering/governance/wg-governance.md
[README.md]: https://git.k8s.io/community/wg-device-management/README.md
[sigs.yaml]: https://git.k8s.io/community/sigs.yaml
