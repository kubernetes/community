---
title: Image Promoter
authors:
  - "@javier-b-perez"
owning-sig: sig-release
participating-sigs:
  - TBD
reviewers:
  - "@AishSundar"
  - "@BenTheElder"
  - "@dims"
  - "@listx"
approvers:
  - "@thockin"
creation-date: 2018-09-05
last-updated: 2018-11-14
status: implementable
---

# Image Promoter

## Table of Contents

* [Table of Contents](#table-of-contents)
* [Summary](#summary)
* [Motivation](#motivation)
  * [Goals](#goals)
* [Proposal](#proposal)
  * [Staging Container Registry](#staging-container-registry)
  * [Production Container Registry](#production-container-registry)
  * [Promotion Process](#promotion-process)
* [Graduation Criteria](#graduation-criteria)
* [Infrastructure Needed](#infrastructure-needed)


## Summary

For security reasons, we cannot allow everyone to publish container images into the official kubernetes container registry. This is why we need a process that allows us to review who built an image and who approved it to be shown in the official channels.


## Motivation

There are multiple reasons why we should have a process to publish container images in place:

* We cannot allow all community members to publish images into the official kubernetes container registry.
* We should restrict who can push images to a small set of members and some systems accounts for automation.
* We can run scans and tests on the images before we publish them into the official kubernetes container registry.
* The process to publish into an official channel shouldn't be hard or long to follow. We don’t want to block developers or releases.

### Goals

1. Define a process for publishing container images into an official GCR through a code review process and automated promotion from GCR staging environment.
1. Allow the community to own and manage the project registries.

## Proposal

Following the *GitOps* idea, the proposal is to use a code review process to approve publishing container images into official distribution channels.

This requires two GCR registries:

* Staging: temporary container registry to share container images for testing and scanning.
* Production or *official*: GCR used to host all the approved container images by the community.

### Staging Container Registry

This temporary storage allows to have a public place where to pull images and run qualification tests or vulnerability scans on the images before pushing them to the *official* container registry.

Each project/subproject in the community, will require at least one member of their community to have push access to the staging area.

### Production Container Registry

A restricted set of members can have push access to override any tool or process if necessary.
Ideally we only push images that have been approved by the owners of the production container registry, following the promotion process.

### Promotion Process

1. Maintainer create a container image and push it into *staging* GCR.
1. Maintainer creates a PR in GitHub to add the new image into the *official* container registry.
1. Once owners of the official container registry approve the change and merge it into the master branch, the promoter tool will automatically copy the container image(s) from *staging* into *official* container registry.

If the infrastructure support it, the promoter tool could sign container images when pushing to the official container registry.

![Promote process](promote-process.jpg?raw=true "Promote process")

In the future, we could add more information into the context of the PR like the vulnerability scan and test results of the container image.

## Graduation Criteria

We will know we are done when we have:

* User guide for developers/maintainers: how to build? how to promote?
* User guide for owners: review and approve PR, how to push images?
* A repository to host the manifest file.
* Initial set of repository's owners who can approve changes.
* Criteria to grant or remove access to staging and production container registries.
* A tool that automatically copy approved images into official channels.

## Infrastructure Needed

* Two GCP projects with GCR enabled.
  * One project should have GCB enabled to run the promotion tool in it.
* Repository to host the manifest for promotions.

