# Flake Finder Fridays #0 

February 5th 2021 ([Recording](https://youtu.be/Hqlm2h2AEvA))

Hosts: [Dan Mangum](https://github.com/hasheddan), [Rob
Kielty](https://github.com/RobertKielty)

## Introduction

This is the first episode of Flake Finder Fridays with Dan Mangum and Rob
Kielty. 

On the first friday of every month we will go through an issue that was logged
for a failing or flaking test on the Kubernetes project.

We will review the triage, root cause analysis, and problem resolution for a
test related issue logged in the past four weeks.

We intend to demo how CI works on the Kubernetes project and also how we
collaborate across teams to resolve test maintenance issues.

## Issue This is the issue that we are going to look at today ...

[[Failing Test] ci-kubernetes-build-canary does not understand
"--platform"](https://github.com/kubernetes/kubernetes/issues/98646)

### Testgrid Dashboard
[build-master-canary](https://testgrid.k8s.io/sig-release-master-informing#build-master-canary)

### Breaking PRs
- [Use buildx in favor of `FROM --platform`
  syntax](https://github.com/kubernetes/kubernetes/pull/98529)
- [Switch to `docker buildx` for conformance
  image](https://github.com/kubernetes/kubernetes/pull/98569)

## Investigation

1. Desire to move from Google-owned infrastructure to Kubernetes community
   infrastructure. Thus the introduction of a **canary** build job to test
   pushing building and pushing artifacts with new infrastructure.
1. Desire to move off of `bootstrap.py` job (currently being used for canary
   job) to `krel` tooling.
1. Separate job existed (`ci-kubernetes-build-no-bootstrap`) that was doing the
   same thing as the canary job, but with `krel` tooling.
1. The `no-bootstrap` job was running smoothly, so [updated to use it for the
   canary job](https://github.com/kubernetes/test-infra/pull/20663).
1. Right before the update, we [switched to using buildx for multi-arch
   images](https://github.com/kubernetes/kubernetes/pull/98529).
1. Job started failing, which showed up in [some interesting
   ways](https://kubernetes.slack.com/archives/C09QZ4DQB/p1612269558032700).
1. Triage begins! Issue
   [opened](https://github.com/kubernetes/kubernetes/issues/98646) and release
   management team is pinged in Slack.
1. The `build-master`
   [job](https://testgrid.k8s.io/sig-release-master-blocking#build-master) was
   still passing though... interesting.
1. Both are eventually calling `make release`, so environment must be different.
1. Let's look inside!

    ```
    docker run -it --entrypoint /bin/bash gcr.io/k8s-testimages/bootstrap:v20210130-12516b2
    ```

    ```
    docker run -it gcr.io/k8s-staging-releng/k8s-ci-builder:v20201128-v0.6.0-6-g6313f696-default /bin/bash
    ```

1. A few directions we could go here:
    1. Update the `k8s-ci-builder` image to you use newer version of Docker
    1. Update the `k8s-ci-builder` image to ensure that
       `DOCKER_CLI_EXPERIMENTAL=enabled` is set
    1. Update the `release.sh` script to set `DOCKER_CLI_EXPERIMENTAL=enabled`

1. Making the `release.sh` script more flexible serves the community better
   because it allows for building with more environments. Would also be good to
   update the `k8s-ci-builder` image for this specific case as well.
1. And we get a new
   [failure](https://storage.googleapis.com/kubernetes-jenkins/logs/ci-kubernetes-build-canary/1356704759045689344/build-log.txt)!
1. Let's see what is going on in those images again...
1. Why would this cause an error in one but not the other if we have
   `DOCKER_CLI_EXPERIMENTAL=enabled`?
   ([this](https://github.com/docker/buildx/pull/403) is why)
1. In the mean time we went ahead and [re-enabled the bootstrap
   job](https://github.com/kubernetes/test-infra/pull/20712) (consumers of those
   images need them!)
1. Decided to [increase logging
   verbosity](https://github.com/kubernetes/kubernetes/pull/98568) on failures
   to see if that would give us a clue into what was going wrong (and to remove
   those annoying `quiet currently not implemented` warnings).
1. Job turns green! But how?
1. [Buildx](https://github.com/docker/buildx) is versioned separately than
   Docker itself. Turns out that the `--quiet` flag warning was [actually an
   error](https://github.com/docker/buildx/pull/403) until `v0.5.1` of Buildx.
1. The `build-master` job was running with buildx `v0.5.1` while the `krel` job
   was running with `v0.4.2`. This meant the quiet flag was causing an error in
   the `krel` job, and removing it alleviated the error.
1. Finished up by once again [removing the `bootstrap`
   job](https://github.com/kubernetes/test-infra/pull/20731).

### Fixes

- [Set DOCKER_CLI_EXPERIMENTAL=enabled for images using
  buildx](https://github.com/kubernetes/kubernetes/pull/98672)
- [Make image build logs verbose if
  necessary](https://github.com/kubernetes/kubernetes/pull/98568)

### Test Infra

- [ci-kubernetes-build-canary: Migrate from bootstrap to
  krel](https://github.com/kubernetes/test-infra/pull/20663)
- [releng: Re-enable a bootstrap build job for K8s
  Infra](https://github.com/kubernetes/test-infra/pull/20712)
- [Revert "releng: Re-enable a bootstrap build job for K8s
  Infra"](https://github.com/kubernetes/test-infra/pull/20731)

### Slack Threads

- [kubeadm failing with
  ci/latest](https://kubernetes.slack.com/archives/C09QZ4DQB/p1612269558032700)

### Helpful Links

- [Docker Buildx
  Documentation](https://docs.docker.com/buildx/working-with-buildx/)
- [What is Docker Buildkit and What can I use it
  for?](https://brianchristner.io/what-is-docker-buildkit/)
- [Buildx --quiet error](https://github.com/docker/buildx/pull/403)

## Kubernetes Project Resources

Brand new to the project? 
 - Start here: https://www.kubernetes.dev/

Setup already and interested in maintaining tests? 
 - Check out [this video](https://www.youtube.com/watch?v=Ewp8LNY_qTg) from
   Jordan Liggit who describes strategies and tactics to deflake flaking tests
   ([Jordan's show notes for that
   talk](https://gist.github.com/liggitt/6a3a2217fa5f846b52519acfc0ffece0))

Here's how the CI Signal Team actively monitors CI during a release cycle:
 - [A Tour of CI on the Kubernetes
   Project](https://www.youtube.com/watch?v=bttEcArAjUw)
 - [Show notes](bit.ly/k8s-ci)
