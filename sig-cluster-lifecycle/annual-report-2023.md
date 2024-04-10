# 2023 Annual Report: SIG Cluster Lifecycle

## Current initiatives and Project Health

1. What work did the SIG do this year that should be highlighted?

<!--
   Some example items that might be worth highlighting:
   - Major KEP advancement
   - Important initiatives that aren't tracked via KEPs
   - Paying down significant tech debt
   - Governance and leadership changes
-->

- We had no major KEP enhancements, just one security improving KEP targeting kubeadm mentioned below.
- There were no governance and leadership changes in 2023.
- We collected [subproject self-assessment](https://docs.google.com/forms/d/e/1FAIpQLSc0CqmfaOIK4bCbEDhh0qiF5wCHi6Uvy0uR_k8egOtafalpow/viewanalytics)
from all of our subprojects. This has become a common practice for us, so that we have good signal for subproject
health, given we have so many subprojects. We perform the collection of self-assessments on a best effort period of time.
For example the previous one was in 2021.
- Some subproject highlights for 2023, in no particular order and definitely not a complete list:
  - Minikube released a GUI.
  - kOps started participating in an initiative to replace the legacy kube-up deployer for testing K8s.
  - kubeadm started work on v1beta4.
  - cluster-api-provider-ibmcloud completed a code-a-thon: A one day hackathon.
  - cluster-api-operator continued working on simplifying the UX around CAPI usage.
  - cluster-api: focused on improving performance and MachinePools.

2. Are there any areas and/or subprojects that your group needs help with (e.g. fewer than 2 active OWNERS)?

<!--
   Note: This list is generated from the KEP metadata in kubernetes/enhancements repository.
      If you find any discrepancy in the generated list here, please check the KEP metadata.
      Please raise an issue in kubernetes/community, if the KEP metadata is correct but the generated list is incorrect.
-->

As established by our self-assessment mentioned above, the most common areas where we need help are:
- Contributing to code
- Reviewing code

Subprojects that might need help with additional OWNERS, in no particular order, and maybe an incomplete list:
- cluster-api-addon-provider-helm
- cluster-api-ipam-provider-in-cluster
- cluster-api-provider-gcp
- cluster-api-provider-kubemark
- cluster-api-provider-openstack
- cluster-api-provider-packet
- kubeadm

We encourage contributors who are interested in our subprojects to reach out to the SIG mailing list or to
individual subproject Slack channels.

3. Did you have community-wide updates in 2023 (e.g. KubeCon talks)?

For KubeCon EU 2023, we had the following talks:
- [Minikube, from CLI to GUI](https://sched.co/1HyUz)
- [How to Turn Release Management from Duty to Fun: Lessons Learned Building the Cluster API Release Team](https://sched.co/1HyTd)
- [Cluster API Providers: Intro, Deep Dive, and Community!](https://sched.co/1HyUb)
- [Kubeadm Deep Dive](https://sched.co/1Iki0)
- [Mission Accomplished: Kubernetes Is Not a Monorepo. Now Our Work Begins!](https://sched.co/1HycF)

For KubeCon NA 2023, we had the following talks:
- [SIG Cluster Lifecycle Intro & Future](https://sched.co/1R2sL)
- [Minikube project update](https://sched.co/1R2rH)
- [Cluster API Deep Dive: Improving Performance up to 2k Clusters](https://sched.co/1R2py)

The above lists do not include SIG CL subproject related talks from third parties.

4. KEP work in 2023 (v1.27, v1.28, v1.29):

While we have a few SIG level KEPs in `kubernetes/enhancements`, kubeadm is the only subproject that we have that is part
of the Kubernetes release and more actively receives KEPs. In 2023 we worked on a KEP to move kubeadm away from generating
a single `admin.conf` that has super powers and instead generate two - `admin.conf` and `super-admin.conf`:
- https://git.k8s.io/enhancements/keps/sig-cluster-lifecycle/kubeadm/4214-separate-super-user-kubeconfig/README.md
It graduated to GA within a single Kubernetes release.

## [Subprojects](https://git.k8s.io/community/sig-cluster-lifecycle#subprojects)


**New in 2023:**
  - [cluster-api-addon-provider-helm](https://git.k8s.io/community/<no value>#cluster-api-addon-provider-helm)
  - [cluster-api-ipam-provider-in-cluster](https://git.k8s.io/community/<no value>#cluster-api-ipam-provider-in-cluster)
**Continuing:**
  - cluster-addons
  - cluster-api
  - cluster-api-operator
  - cluster-api-provider-aws
  - cluster-api-provider-azure
  - cluster-api-provider-cloudstack
  - cluster-api-provider-digitalocean
  - cluster-api-provider-gcp
  - cluster-api-provider-ibmcloud
  - cluster-api-provider-kubemark
  - cluster-api-provider-kubevirt
  - cluster-api-provider-nested
  - cluster-api-provider-openstack
  - cluster-api-provider-packet
  - cluster-api-provider-vsphere
  - etcdadm
  - image-builder
  - kOps
  - kubeadm
  - kubespray
  - minikube

## [Working groups](https://git.k8s.io/community/sig-cluster-lifecycle#working-groups)

**New in 2023:**
 - LTS
**Retired in 2023:**
 - Reliability

## Operational

Operational tasks in [sig-governance.md]:
- [x] [README.md] reviewed for accuracy and updated if needed
- [x] [CONTRIBUTING.md] reviewed for accuracy and updated if needed
- [x] Other contributing docs (e.g. in devel dir or contributor guide) reviewed for accuracy and updated if needed
- [x] Subprojects list and linked OWNERS files in [sigs.yaml] reviewed for accuracy and updated if needed
- [x] SIG leaders (chairs, tech leads, and subproject leads) in [sigs.yaml] are accurate and active, and updated if needed
- [x] Meeting notes and recordings for 2023 are linked from [README.md] and updated/uploaded if needed


[CONTRIBUTING.md]: https://git.k8s.io/community/sig-cluster-lifecycle/CONTRIBUTING.md
[sig-governance.md]: https://git.k8s.io/community/committee-steering/governance/sig-governance.md
[README.md]: https://git.k8s.io/community/sig-cluster-lifecycle/README.md
[sigs.yaml]: https://git.k8s.io/community/sigs.yaml
[devel]: https://git.k8s.io/community/contributors/devel/README.md
