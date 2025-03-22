---
title: "Getting Started"
weight: 1
aliases: ["/guide"]
description:
  A small list of things that you should read and be familiar with before you
  get started with contributing. This includes such things as signing the
  Contributor License Agreement, familiarizing yourself with our Code of Conduct,
  and more.
---

# Getting Started

- [Getting Started](#getting-started)
  - [Welcome](#welcome)
  - [Contributor Guide](#contributor-guide)
  - [Prerequisites](#prerequisites)
    - [Create a GitHub account](#create-a-github-account)
    - [Sign the CLA](#sign-the-cla)
    - [Code of Conduct](#code-of-conduct)
    - [Setting up your development environment](#setting-up-your-development-environment)
    - [Community Expectations and Roles](#community-expectations-and-roles)
  - [Kubernetes Contributor Playground](#kubernetes-contributor-playground)
    - [Contributor Workshops](#contributor-workshops)
  - [Community](#community)
    - [Communication](#communication)
    - [Events](#events)
    - [Meetups](#meetups)
    - [Mentorship](#mentorship)
  - [Advanced Topics](#advanced-topics)

## Welcome

Have you ever wanted to contribute to the coolest cloud technology?
This guide will help you understand the overall organization of the [Kubernetes
project](https://kubernetes.io/), and direct you to the best places to get started contributing. You'll
be able to pick up issues, write code to fix them, and get your work reviewed
and merged.

This document is the single source of truth for how to contribute to the code
base. Feel free to browse the [open issues] and file new ones, all feedback
is welcome!

## Contributor Guide

Welcome to Kubernetes! This guide is broken up into the following sections.
It is recommended that you follow these steps in order:

- [Welcome](#welcome) - this page
- [Prerequisites](#prerequisites) - tasks you need to complete before
  you can start contributing to Kubernetes
- [Your First Contribution](./first-contribution.md) - things you'll need to know
  before making your first contribution
- [Contributing](./contributing.md) - the main reference guide to contributing
  to Kubernetes

## Prerequisites

Before submitting code to Kubernetes, you should first complete the following
prerequisites. These steps are checked automatically by [a bot] during your
first submission. Completing these steps will make your first contribution
easier:

### Create a GitHub account

Before you get started, you will need to [sign up](http://github.com/signup) for a GitHub user account.

### Sign the CLA

Before contributing to Kubernetes, you must sign the [Contributor License Agreement] (CLA). To sign the CLA, open a PR against the [contributor playground repo](https://github.com/kubernetes-sigs/contributor-playground). You can find an example PR [here](https://github.com/kubernetes-sigs/contributor-playground/pull/1389). Afterward, the linux-foundation-easycla bot will reply with an error message containing a specific link for your GitHub account to sign the CLA using CNCF's EasyCLA system. 


**Note:** [kubernetes-sigs/contributor-playground](https://github.com/kubernetes-sigs/contributor-playground) is a repository for practicing submitting your first PR, not just for CLA registration. When creating and submitting a new PR, please write the PR description according to the provided template. Check the above example PR for reference.

### Code of Conduct

Please make sure to read and observe the [Code of Conduct] and
[Community Values]

### Setting up your development environment

It is not required to set up a developer environment in order to contribute to
Kubernetes.

If you plan to contribute code changes, review the [developer resources] page
for how to set up your environment.

### Community Expectations and Roles

Kubernetes is a community project. Consequently, it is wholly dependent on its
community to provide a productive, friendly, and collaborative environment.

- Read and review the [Community Expectations] for an
  understanding of code and review expectations.
- See [Community Membership][CM] for a list of the various
  responsibilities of contributor roles.
- You are encouraged to move up this contributor ladder as you gain experience.

## Kubernetes Contributor Playground

If you are looking for a safe place, where you can familiarize yourself with
the pull request and issue review process in Kubernetes, then the
[Kubernetes Contributor Playground] is the right place for you.

### Contributor Workshops

A [Youtube playlist] of the New Contributor workshop has been posted. An
outline of the video content can be found [here].

## Community

Kubernetes is a large, lively, friendly open-source community. As many open
source projects often do, it depends on new people becoming members and regular
code contributors. The [Community Membership Document][CM] covers membership
processes and roles. Please consider joining Kubernetes, and making your way
up the contributor ladder!

### Communication

- [General Information] relating to Kubernetes communication policies

### Events

Kubernetes participates in KubeCon + CloudNativeCon, held three times per year
in China, Europe, and North America. Information about these and other
community events is available on the CNCF [Events] pages.

### Meetups

All Kubernetes meetups follow the general [Cloud Native Computing Foundation Guidelines]
You may also contact CNCF Staff driving the Community Groups (previously known
as Meetups) program by email (meetups@cncf.io)

### Mentorship

Learn more about the Kubernetes [mentoring initiatives].

## Advanced Topics

This section includes things that need to be documented, but typical contributors
do not need to interact with regularly.

- [OWNERS files] - The Kubernetes organizations are managed with OWNERS files,
  which outline which parts of the code are owned by what groups.

[a bot]: https://github.com/k8s-ci-robot
[Contributor License Agreement]: /CLA.md
[Code of Conduct]: /code-of-conduct.md
[Community Values]: /values.md
[First Contribution]: ./first-contribution.md
[Contributing]: ./contributing.md
[Developer Resources]: /contributors/devel/README.md#setting-up-your-dev-environment-coding-and-debugging
[Community Expectations]: ./expectations.md
[CM]: /community-membership.md
[here]: /events/2019/11-contributor-summit
[General Information]: /communication
[mentoring initiatives]: /mentoring/README.md
[OWNERS files]: ./owners.md
[Cloud Native Computing Foundation Guidelines]: https://github.com/cncf/communitygroups
[Events]: https://www.cncf.io/events/
[YouTube Playlist]: https://www.youtube.com/playlist?list=PL69nYSiGNLP3M5X7stuD7N4r3uP2PZQUx
[Kubernetes Contributor Playground]: https://github.com/kubernetes-sigs/contributor-playground/blob/master/README.md
[Open Issues]: https://github.com/kubernetes/community/issues?q=is%3Aissue+is%3Aopen+label%3Aarea%2Fcontributor-guide
