# Contributing

Welcome to the Kubernetes SIG Cluster Lifecycle contributing guide. We are excited
about the prospect of you joining our [community](https://git.k8s.io/community/sig-cluster-lifecycle/)!

## Before You Begin

We strongly recommend you to understand the main
[Kubernetes Contributor Guide](https://git.k8s.io/community/contributors/guide)
and adhere to the contribution rules (specially signing the CLA).

You can also check the [Contributor Cheat Sheet](/contributors/guide/contributor-cheatsheet/),
with common resources for existing developers.

Please be aware that all contributions to Kubernetes projects require time and commitment
from project maintainers to direct and review work. This is done in additional to many other
maintainer responsibilities, and direct engagement from maintainers is a finite resource.

## SIG Cluster Lifecycle explained

Read the SIG mission outlined in the SIG [charter](https://git.k8s.io/community/sig-cluster-lifecycle/charter.md).

Video resources:
- [SIG introduction at KubeCon NA 2020](https://www.youtube.com/watch?v=qi-X-Wszetc)
  - [Slides](https://docs.google.com/presentation/d/18I1YvBUegWegc7oBJiLLxwA2I1c9VVjH)
- [SIG contributor onboarding](https://www.youtube.com/watch?v=Bof9aveB3rA)
- [SIG meeting VOD archive](https://www.youtube.com/playlist?list=PL69nYSiGNLP29D0nYgAGWt1ZFqS9Z7lw4)

## Get in touch with the SIG

Find the SIG contact details in its [community page](https://git.k8s.io/community/sig-cluster-lifecycle/README.md#contact):
- Join the SIG [mailing list](https://groups.google.com/forum/#!forum/kubernetes-sig-cluster-lifecycle)
- Join the SIG [slack channel](https://kubernetes.slack.com/messages/sig-cluster-lifecycle)
- Join the periodic SIG [video call](https://git.k8s.io/community/sig-cluster-lifecycle/README.md#meetings)

Using the SIG mailing list or video call is preferred for wider discussion topics that affect
multiple subprojects. The main SIG slack channel should only be used for SIG level updates
and more urgent matters.

Note that individual subprojects have their own slack channels and video calls (see bellow).

## Picking a subproject to contribute to

Some subprojects like kubeadm and etcdadm are lower in the stack (operate on the host machine),
while other subprojects like Cluster API and kops are higher in the stack and also manage
host machine provisioning and cloud provider setup. A good starting point for contributions
can be a project that you have used already or have plans using in the future.

See the [list of subprojects](https://git.k8s.io/community/sig-cluster-lifecycle/README.md#subprojects)
the SIG maintains and pick a project you wish to work on.

Note that individual subprojects follow different process in terms of:
- Release cycle
- Issue triage
- Implementing features
- Change submissions
- Meeting format

Join the subproject video call and slack channel and introduce your self.

Navigate to the subproject repository:
- Read their `README.md` file to understand what the project is about
- Read their `CONTRIBUTING.md` file if you wish to contribute
- See who the maintainers of the project are in the `OWNERS` file
- Find issues labeled with `good-first-issue` and `help-wanted`
- If you wish to work on an issue, `@` mention the issue author and `/assign @your-self`
- Once you have started the work, label the issue with `/lifecycle active`
- Coordinate with the subproject maintainers the submission of changes
- Provide updates on your work in the subproject video call / slack channel
