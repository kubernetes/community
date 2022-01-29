# 2021 Annual Report: SIG Windows

## Current initiatives

1. What work did the SIG do this year that should be highlighted?
   - Lead the transition to hostProcess containers.
   - Defined the `kubectl node logs` command interface.
   - Made the developer UX for windows transparent with sig-windows-dev-tools.
   - Defined windows operational readiness standards.
   
2. What initiatives are you working on that aren't being tracked in KEPs?
   - Migration of the windows kube-proxy to KPNG (led by amim).
   - ...

3. KEP work in 2021 (1.x, 1.y, 1.z):
   - Stable
     - [1122 - windows-csi-support](https://github.com/kubernetes/enhancements/blob/master/keps/sig-windows/1122-windows-csi-support/kep.yaml)
   - Beta
     - [1981 - Windows Privileged Container Support](https://github.com/kubernetes/enhancements/blob/master/keps/sig-windows/1981-windows-privileged-container-support/README.md)
   - Alpha
     - [2258 - Node service log viewer](https://github.com/kubernetes/enhancements/blob/master/keps/sig-windows/2258-node-service-log-viewer/README.md)
   - Pre-alpha
     - [2578 - Windows Conformance](https://github.com/kubernetes/enhancements/blob/master/keps/sig-windows/2578-windows-conformance/kep.yaml)

## Project health

1. What areas and/or subprojects does your group need the most help with?
   Any areas with 2 or fewer OWNERs? (link to more details)
   - csi-proxy and storage: this seems like an underserved area for windows
2. What metrics/community health stats does your group care about and/or measure?
   - n/a
3. Does your [CONTRIBUTING.md] help **new** contributors engage with your group specifically by pointing
   to activities or programs that provide useful context or allow easy participation?
   - yes
4. If your group has special training, requirements for reviewers/approvers, or processes beyond the general [contributor guide],
   does your [CONTRIBUTING.md] document those to help **existing** contributors grow throughout the [contributor ladder]?
   - yes

5. Does the group have contributors from multiple companies/affiliations?
   - yes

6. Are there ways end users/companies can contribute that they currently are not?
   If one of those ways is more full time support, what would they work on and why?
   - testing hostProcess implementations on several windows apps
   - improving our dev tools enviornment to grow the commnunity
   - hardening the CSI proxy and CSI support ecosystem
   - performance testing Kuberentes on Windows extensively and publishing results

## Membership
- Primary slack channel member count: 1507
- Primary mailing list member count: 188
- Primary meeting attendee count (estimated, if needed): 10
- Primary meeting participant count (estimated, if needed): 10
- Unique reviewers for SIG-owned packages: 6
- Unique approvers for SIG-owned packages: 4

Include any other ways you measure group membership

## Subprojects

Subprojects for sig-windows are inactive now and all major work happens in the sig-windows channel.

## Working groups

Same as above.

## Operational

Operational tasks in [sig-governance.md]:

- [ ] [README.md] reviewed for accuracy and updated if needed
- [ ] [CONTRIBUTING.md] reviewed for accuracy and updated if needed
      (or created if missing and your contributor steps and experience are different or more
      in-depth than the documentation listed in the general [contributor guide] and [devel] folder.)
- [ ] Subprojects list and linked OWNERS files in [sigs.yaml] reviewed for accuracy and updated if needed
- [ ] SIG leaders (chairs, tech leads, and subproject owners) in [sigs.yaml] are accurate and active, and updated if needed
- [ ] Meeting notes and recordings for 2021 are linked from [README.md] and updated/uploaded if needed
- [ ] Did you have community-wide updates in 2021 (e.g. community meetings, kubecon, or kubernetes-dev@ emails)? Links to email, slides, or recordings:
      -
      -

[CONTRIBUTING.md]: https://git.k8s.io/community/sig-windows/CONTRIBUTING.md
[contributor ladder]: https://git.k8s.io/community/community-membership.md
[sig-governance.md]: https://git.k8s.io/community/committee-steering/governance/sig-governance.md
[README.md]: https://git.k8s.io/community/sig-windows/README.md
[sigs.yaml]: https://git.k8s.io/community/sigs.yaml
[contributor guide]: https://git.k8s.io/community/contributors/guide/README.md
[devel]: https://git.k8s.io/community/contributors/devel/README.md
