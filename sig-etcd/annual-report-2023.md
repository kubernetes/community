# 2023 Annual Report: SIG etcd

## Current initiatives and Project Health

1. What work did the SIG do this year that should be highlighted?

With etcd becoming a Kubernetes project special interest group in September 2023 over the remainder of the calendar year we:

- Released the first alpha for bbolt `v1.4.0`, a critical roadmap item for etcd `v3.6.0` which substantially improves logging and introduces bbolt surgery commands for working directly with storage files.
- Released the first alpha for raft `v3.6.0`, a critical roadmap item for etcd `v3.6.0` which adds support for asynchronous storage writes and a number of other changes.
- Donated `jpbetz/auger` to `etcd-io/auger` with subproject leads in place. With a long term home the project is now revived and being actively maintained.
- In December as part of etcd `v3.5.11` we introduced new livez/readyz HTTP endpoints following design via the KEP process. The new endpoints are Kubernetes API compliant and a significant improvement over the old health endpoint. These new probes are planned to be included by kubeadm 1.31.
- Released 3 patch versions for etcd stable releases between September - December 2023 .
- Fixed an issue with our experimental distributed tracing feature to ensure sampling rates could be properly configured.


2. Are there any areas and/or subprojects that your group needs help with (e.g. fewer than 2 active OWNERS)?

- An emerging risk for SIG etcd is the raft subproject which is currently operating on two primary approvers after a longstanding maintainer retired in July 2023. The two primary approvers have limited bandwidth for raft as they are also key maintainers for etcd. Over 2024 we need help growing a new maintainership for `etcd-io/raft`.

- The second area of concern for SIG etcd is the etcd grpc-proxy which also lacks maintainership as well as evidence for use cases. This is an area of functionality we may consider for deprecation in future if we aren't able to overcome the the maintainership challenge through a subproject lead.


3. Did you have community-wide updates in 2023 (e.g. KubeCon talks)?

The etcd project team has had talks at both Kubecon events in 2023:

- [KubeCon Europe 2023, On the hunt for etcd data inconsistencies](https://www.youtube.com/watch?v=IIMs0EjQZHg)
- [KubeCon North America 2023, Forging a stronger bond between etcd and Kubernetes](https://www.youtube.com/watch?v=6JYgBJAjpNQ)
- [KubeCon North America 2023, Secrets of running etcd](https://www.youtube.com/watch?v=aJVMWcVZOPQ)

Additionally we held ContribFest events at both KubeCons in 2023 which have been successful in growing new contributors. In 2023 we also introduced a new fortnightly triage meeting to complement our existing fortnightly community meeting.

All meeting recording and maintainer track talks are available on YouTube: https://www.youtube.com/@etcdio


4. KEP work in 2023 (v1.27, v1.28, v1.29):

The etcd team do not yet have KEP's tracked against specific Kubernetes versions however we do have [two open enhancements](https://github.com/kubernetes/enhancements/issues?q=is%3Aissue+etcd+label%3Asig%2Fetcd+is%3Aopen):
- https://github.com/kubernetes/enhancements/issues/4326
- https://github.com/kubernetes/enhancements/issues/4331


## [Subprojects](https://git.k8s.io/community/sig-etcd#subprojects)


**New in 2023:**
- [auger](https://github.com/etcd-io/auger)

**Added under k8s as part of sig-etcd creation**
- [bbolt](https://github.com/etcd-io/bbolt)
- [dbtester](https://github.com/etcd-io/dbtester)
- [discovery.etcd.io](https://github.com/etcd-io/discovery.etcd.io)
- [discoveryserver](https://github.com/etcd-io/discoveryserver)
- [etcd](https://github.com/etcd-io/etcd)
- [etcdlabs](https://github.com/etcd-io/etcdlabs)
- [gofail](https://github.com/etcd-io/gofail)
- [govanityurls](https://github.com/etcd-io/govanityurls)
- [jetcd](https://github.com/etcd-io/jetcd)
- [protodoc](https://github.com/etcd-io/protodoc)
- [raft](https://github.com/etcd-io/raft)
- [website](https://github.com/etcd-io/website)

**Archived in 2023**
- [cetcd](https://github.com/etcd-io/cetcd)
- [etcd-play](https://github.com/etcd-io/etcd-play)
- [maintainers](https://github.com/etcd-io/maintainers)
- [zetcd](https://github.com/etcd-io/zetcd)

## [Working groups](https://git.k8s.io/community/sig-etcd#working-groups)


## Operational

Operational tasks in [sig-governance.md]:
- [x] [README.md] reviewed for accuracy and updated if needed
- [x] [CONTRIBUTING.md] reviewed for accuracy and updated if needed
- [x] Other contributing docs (e.g. in devel dir or contributor guide) reviewed for accuracy and updated if needed
- [x] Subprojects list and linked OWNERS files in [sigs.yaml] reviewed for accuracy and updated if needed
- [x] SIG leaders (chairs, tech leads, and subproject leads) in [sigs.yaml] are accurate and active, and updated if needed
- [x] Meeting notes and recordings for 2023 are linked from [README.md] and updated/uploaded if needed


[CONTRIBUTING.md]: https://git.k8s.io/community/sig-etcd/CONTRIBUTING.md
[sig-governance.md]: https://git.k8s.io/community/committee-steering/governance/sig-governance.md
[README.md]: https://git.k8s.io/community/sig-etcd/README.md
[sigs.yaml]: https://git.k8s.io/community/sigs.yaml
[devel]: https://git.k8s.io/community/contributors/devel/README.md
