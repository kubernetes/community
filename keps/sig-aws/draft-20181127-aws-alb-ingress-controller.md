---
kep-number: draft-20181127
title: AWS ALB Ingress Controller
authors:
  - "@M00nF1sh"
owning-sig: sig-aws
reviewers:
  - TBD
  - "@d-nishi"
approvers:
  - TBD
  - "@d-nishi"
editor: TBD
creation-date: 2018-11-27
last-updated: 2018-11-27
status: provisional
---

# AWS ALB Ingress Controller

## Table of Contents
- [Table of Contents](#table-of-contents)
- [Summary](#summary)
- [Motivation](#motivation)
  - [Goals](#goals)
  - [Non-Goals](#non-goals)
- [Proposal](#proposal)
  - [User Stories](#user-stories)
    - [Expose HTTP[s] API backed by kubernetes services](#expose-https-api-backed-by-kubernetes-services)
    - [Adjust ALB settings via annotation](#adjust-alb-settings-via-annotation)
    - [Leverage WAF & Cognito](#leverage-waf--cognito)
    - [Sharing single ALB among Ingresses across namespace](#sharing-single-alb-among-ingresses-across-namespace)
- [Graduation Criteria](#graduation-criteria)
- [Implementation History](#implementation-history)

## Summary

This proposal introduces [AWS ALB Ingress Controller](https://github.com/kubernetes-sigs/aws-alb-ingress-controller/) as Ingress controller for kubernetes cluster on AWS. Which use [Amazon Elastic Load Balancing Application Load Balancer](https://aws.amazon.com/elasticloadbalancing/features/#Details_for_Elastic_Load_Balancing_Products)(ALB) to fulfill [Ingress resources](https://kubernetes.io/docs/concepts/services-networking/ingress/), and provides integration with various AWS services.

## Motivation

In order for the Ingress resource to work, the cluster must have an Ingress controller runnings. However, existing Ingress controllers like [nginx](https://github.com/kubernetes/ingress-nginx/blob/master/README.md) didn't take advantage of native AWS features.
AWS ALB Ingress Controller aims to enhance Ingress resource on AWS by leveraging rich feature set of ALB, such as host/path based routing, TLS termination, WebSockets, HTTP/2. Also, it will provide close integration with other AWS services such as WAF(web application firewall) and Cognito.

### Goals

* Support running multiple Ingress controllers in cluster
* Support portable Ingress resource(no annotations)
* Support leverage feature set of ALB via custom annotations
* Support integration with WAF
* Support integration with Cognito

### Non-Goals

* This project does not replacing nginx ingress controller

## Proposal

### User Stories

#### Expose HTTP[s] API backed by kubernetes services
Developers create an Ingress resources to specify rules for how to routing HTTP[s] traffic to different services.
AWS ALB Ingress Controller will monitor such Ingress resources and create ALB and other necessary supporting AWS resources to match the Ingress resource specification.

#### Adjust ALB settings via annotation
Developers specifies custom annotations on their Ingress resource to adjust ALB settings, such as enable deletion protection, enable access logs to specific S3 bucket.

#### Leverage WAF & Cognito
Developers specifies custom annotations on their Ingress resource to denote WAF and Cognito integrations. Which provides web application firewall and authentication support for their exposed API.

#### Sharing single ALB among Ingresses across namespace
Developers from different teams create Ingress resources in different namespaces which route traffic to services within their own namespace. However, an single ALB is shared from these Ingresses to expose a single DNS name for customers.

## Graduation Criteria

* AWS ALB Ingress Controller is widely used as Ingress controller for kubernetes clusters on AWS

## Implementation History
- [community#2841](https://github.com/kubernetes/community/pull/2841) Design proposal
- [aws-alb-ingress-controller#738](https://github.com/kubernetes-sigs/aws-alb-ingress-controller/pull/738) First stable release: v1.0.0