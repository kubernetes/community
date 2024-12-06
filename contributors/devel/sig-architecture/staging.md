# Staging Directory and Publishing

The [staging/ directory](https://git.k8s.io/kubernetes/staging) of Kubernetes contains a number of pseudo repositories ("staging repos"). They are symlinked into Kubernetes' [vendor/ directory](https://git.k8s.io/kubernetes/vendor/k8s.io) for Golang to pick them up.

We publish the staging repos using the [publishing bot](https://git.k8s.io/publishing-bot). It uses `git filter-branch` essentially to [cut the staging directories into separate git trees](https://de.slideshare.net/sttts/cutting-the-kubernetes-monorepo-in-pieces-never-learnt-more-about-git) and pushing the new commits to the corresponding real repositories in the [kubernetes organization on Github](https://github.com/kubernetes).

The list of staging repositories currently published is available in [staging/README.md inside of the k8s.io/kubernetes repository](https://git.k8s.io/kubernetes/staging/README.md).

At the time of this writing, based on [publishing rule](https://github.com/kubernetes/kubernetes/blob/9cf4f912451009296a50f267d47b484f77ce05e4/staging/publishing/rules.yaml), the list of published branches includes:

- master,
- release-1.28
- release-1.29
- release-1.30
- release-1.31

Kubernetes tags (e.g., `v1.30.0-beta.0`) are also applied automatically to the published repositories, prefixed with `kubernetes-`.
From `v1.17.0` Kubernetes release, matching semver `v0.x.y` tags are also created for each `v1.x.y` Kubernetes tag.

For example, if you check out the `kubernetes-1.30.0` or the `v0.30.0` tag in
a published repo, the code you get is exactly the same as if you check out the
`v1.30.0` tag in Kubernetes, and change directory to `staging/src/k8s.io/<repo-name>`.

It is recommend to use the semver `v0.x.y` tags for a seamless experience
with go modules.

If further repos under staging are needed, adding them to the bot is easy.
Contact one of the [owners of the bot](https://git.k8s.io/publishing-bot/OWNERS).

Currently, the bot is hosted on a
[public CNCF cluster](https://github.com/kubernetes/publishing-bot/blob/master/production.md).
