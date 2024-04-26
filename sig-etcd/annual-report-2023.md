# 2023 Annual Report: SIG etcd

## Current initiatives and Project Health

1. What work did the SIG do this year that should be highlighted?

   - Major KEP advancement
     - Made significant improvements to our robustness testing
       - Added support for automated disk failure injection via lazyfs
       - Added kubernetes specific traffic generation
       - Added further network failure injection with randomised packet dropping
     - Completed the [work](https://github.com/etcd-io/etcd/issues/15951) to bring arm64 to a tier 1 supported platform
       - Migrated all test workflows to run within containers for isolation on dedicated runners
       - Ensured all unit, integration, e2e and robustness testing workflows were setup and running successfully on arm64 for a supported release branch
       - Improved documentation and controls for arm64 ci infra
       - Systematically worked through test workflows to ensure both amd64 and arm64 were configured the same
     - Successfully completed migration from grpc-gateway v1 to v2
       - Simplifies dependencies for both etcd and Kubernetes and downstream consumers
       - Moved etcd off a deprecated grpc-gateway version
       - Unblocked ability for etcd project to upgrade otel and grpc versions
       - Key roadmap item for next minor release etcd 3.6
     - Bump gRPC
     - Bump grpc-gateway
     - bbolt
       - Added logger into bbolt
       - Big progress on resolving the data corruption issues
         - Identified two possible reasons for the data corruption issues
           - [pull/639](https://github.com/etcd-io/bbolt/pull/639)
           - [issues/562](https://github.com/etcd-io/bbolt/issues/562)
         - Surgery commands ([issue/370](https://github.com/etcd-io/bbolt/issues/370))
         - To prevent data corruption: redundancy
       - Other minor features/improvements:
         - Supporting moving bucket inside the same db ([pull/635](https://github.com/etcd-io/bbolt/pull/635))
         - Support inspecting database structure ([pull/674](https://github.com/etcd-io/bbolt/pull/674))
         - Minor performance improvements. (e.g. [pull/419](https://github.com/etcd-io/bbolt/pull/419), [pull/532](https://github.com/etcd-io/bbolt/pull/532))
       -  v1.4.0
         -  Release plan for v1.4.0 (https://github.com/etcd-io/bbolt/issues/553)
         -  v1.4.0-alpha.0 ([changelog](https://github.com/etcd-io/bbolt/blob/main/CHANGELOG/CHANGELOG-1.4.md#v140-alpha02024-01-12))
     - raft
       - raft has already been moved into a separate repo: https://github.com/etcd-io/raft
       - Module name: go.etcd.io/etcd/raft/v3 → go.etcd.io/raft/v3 (https://github.com/etcd-io/bbolt/issues/14713)
       - Support async storage writes (https://github.com/etcd-io/bbolt/pull/8)
       - Other minor features
         - Add ForgetLeader ([pull/78](https://github.com/etcd-io/raft/pull/78))
         - Add StepDownOnRemoval ([pull/79](https://github.com/etcd-io/raft/pull/79))
       - v3.6.0
         - Release plan for v3.6.0 ([issues/89](https://github.com/etcd-io/raft/issues/89))
         - v3.6.0-alpha.0 ([changelog](https://github.com/etcd-io/raft/blob/main/CHANGELOG/CHANGELOG-3.6.md#v360-alpha02024-01-12))
       - Keep supporting RESTful API via grpc-gateway/v2
     - etcd tool for Kubernetes Auger(etcd-io/auger) is donated to etcd
     - Support /livez and /readyz endpoints
     - Support etcd downgrade from 3.5 to 3.4
   - Important initiatives that aren't tracked via KEPs
     - etcd mentorship framework (https://tinyurl.com/etcd-mentorship)
     - Subproject governance framework (https://sched.co/1Yhgg)
   - Paying down significant tech debt
   - Governance and leadership changes
     - Created a SIG etcd to represent the etcd interest in the Kubernetes community
     - Adopted Kubernetes’ governance and delegate decision-making to SIG-etcd
       - SIG Chairs
         - James Blair @ Red Hat
         - Wenjia Zhang @ Google
       - SIG Tech Leads
         - Marek Siarkowicz @ Google
         - Benjamin Wang @ VMWare
      - Adopted Kubernetes' membership model (https://github.com/etcd-io/etcd/pull/15593), with new requirement for membership roles (https://github.com/etcd-io/etcd/blob/main/Documentation/contributor-guide/community-membership.md#requirements)

2. Are there any areas and/or subprojects that your group needs help with (e.g. fewer than 2 active OWNERS)?

<!--
   Note: This list is generated from the KEP metadata in kubernetes/enhancements repository.
      If you find any discrepancy in the generated list here, please check the KEP metadata.
      Please raise an issue in kubernetes/community, if the KEP metadata is correct but the generated list is incorrect.
-->

3. Did you have community-wide updates in 2023 (e.g. KubeCon talks)?

<!--
  Examples include links to email, slides, or recordings.
-->

4. KEP work in 2023 (v1.27, v1.28, v1.29):
   - [Downgrade](https://github.com/kubernetes/enhancements/tree/master/keps/sig-etcd/4326-downgrade)
   - [Livez and Readyz Probes](https://github.com/kubernetes/enhancements/tree/master/keps/sig-etcd/4331-livez-readyz)


## [Subprojects](https://git.k8s.io/community/sig-etcd#subprojects)


**New in 2023:**
  - [bbolt](https://github.com/etcd-io/bbolt)
  - [cetcd](https://github.com/etcd-io/cetcd)
  - [dbtester](https://github.com/etcd-io/dbtester)
  - [discovery.etcd.io](https://github.com/etcd-io/discovery.etcd.io)
  - [discoveryserver](https://github.com/etcd-io/discoveryserver)
  - [etcd](https://github.com/etcd-io/etcd)
  - [etcd-play](https://github.com/etcd-io/etcd-play)
  - [etcdlabs](https://github.com/etcd-io/etcdlabs)
  - [gofail](https://github.com/etcd-io/gofail)
  - [govanityurls](https://github.com/etcd-io/govanityurls)
  - [jetcd](https://github.com/etcd-io/jetcd)
  - [protodoc](https://github.com/etcd-io/protodoc)
  - [raft](https://github.com/etcd-io/raft)
  - [website](https://github.com/etcd-io/website)
  - [zetcd](https://github.com/etcd-io/zetcd)
  - [Auger](https://github.com/etcd-io/auger)

## [Working groups](https://git.k8s.io/community/sig-etcd#working-groups)
  - Robustness test

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
