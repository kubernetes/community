# Staging Directory and Publishing

The [staging/ directory](https://git.k8s.io/kubernetes/staging) of Kubernetes contains a number of pseudo repositories ("staging repos"). They are symlinked into Kubernetes' [vendor/ directory](https://git.k8s.io/kubernetes/vendor/k8s.io) for Golang to pick them up.

We publish the staging repos using the [publishing bot](https://git.k8s.io/publishing-bot). It uses `git filter-branch` essentially to [cut the staging directories into separate git trees](https://de.slideshare.net/sttts/cutting-the-kubernetes-monorepo-in-pieces-never-learnt-more-about-git) and pushing the new commits to the corresponding real repositories in the [kubernetes organization on Github](https://github.com/kubernetes).

The list of staging repositories and their published branches are listed in [publisher.go inside of the bot](https://git.k8s.io/publishing-bot/cmd/publishing-bot/publisher.go). Though it is planned to move this out into the k8s.io/kubernetes repository.

At the time of this writing, this includes the branches

- master,
- release-1.8 / release-5.0,
- and release-1.9 / release-6.0

of the following staging repos in the k8s.io org:

- api
- apiextensions-apiserver
- apimachinery
- apiserver
- client-go
- code-generator
- kube-aggregator
- metrics
- sample-apiserver
- sample-controller

Kubernetes tags (e.g., v1.9.1-beta1) are also applied automatically to the published repositories, prefixed with kubernetes- (e.g., kubernetes-1.9.1-beta1). The client-go semver tags (on client-go only!) including release-notes are still done manually.

The semver tags are still the (well tested) official releases. The kubernetes-1.x.y tags have limited test coverage (we have some automatic tests in place in the bot), but can be used by early adopters of client-go and the other libraries. Moreover, they help to vendor the correct version of k8s.io/api and k8s.io/apimachinery.

If further repos under staging are need, adding them to the bot is easy. Contact one of the [owners of the bot](https://git.k8s.io/publishing-bot/OWNERS).

Currently, the bot is hosted on the CI cluster of Redhat's OpenShift (ready to be moved out to a public CNCF cluster if we have that in the future).
