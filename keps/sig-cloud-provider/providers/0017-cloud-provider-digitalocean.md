---
kep-number: 17
title: Cloud Provider DigitalOcean
authors:
  - "@andrewsykim"
owning-sig: sig-cloud-provider
reviewers:
  - "@hogepodge"
  - "@jagosan"
approvers:
  - "@andrewsykim"
  - "@hogepodge"
  - "@jagosan"
editor: TBD
creation-date: 2018-07-23
last-updated: 2018-07-23
status: provisional

---

# Cloud Provider DigitalOcean

## Table of Contents

* [Table of Contents](#table-of-contents)
* [Summary](#summary)
* [Motivation](#motivation)
    * [Goals](#goals)
    * [Non-Goals](#non-goals)
* [Requirements](#requirements)
* [Proposal](#proposal)

## Summary

DigitalOcean is a cloud computing platform built for developers and businesses who want a simple way to deploy & manage their infrastructure. This is a KEP proposing DigitalOcean as a cloud provider within the Kubernetes ecosystem.

## Motivation

### Goals

* Supporting DigitalOcean within the Kubernetes ecosystem. This involves:
  * providing an open community to promote discussions for Kubernetes on DigitalOcean
  * providing an open environment for developing Kubernetes on DigitalOcean

### Non-Goals

* Using the Kubernetes ecosystem/community to onboard more users onto our platform.

## Prerequisites

### Repository Requirements

The existing repository hosting the [DigitalOcean cloud controller manager](https://github.com/digitalocean/digitalocean-cloud-controller-manager) satisfies requirements as outlined in KEP 0002.

### User Experience Reports

DigitalOcean recently announced a [Kubernetes offering](https://www.digitalocean.com/products/kubernetes/). Many users have already signed up for early access. DigitalOcean is also a gold member of the CNCF.

## Proposal

### Subproject Leads

Initially there will be one subproject lead. In the future the goal is to have 3 subproject leads.

* Andrew Sy Kim (@andrewsykim)


### Repositories

Please create a repository `kubernetes/cloud-provider-digitalocean`.

### Meetings

TBD.

