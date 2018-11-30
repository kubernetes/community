---
kep-number: TBD
title: AWS LoadBalancer Prefix
authors:
  - "@minherz"
owning-sig: sig-aws
participating-sigs:
reviewers:
  - TBD
approvers:
  - TBD
editor: TBD
creation-date: 2018-11-02
last-updated: 2018-11-02
status: provisional
see-also:
replaces:
superseded-by:
---

# AWS LoadBalancer Prefix Annotation Proposal

## Table of Contents

* [Summary](#summary)
* [Motivation](#motivation)
  * [Goals](#goals)
  * [Non-Goals](#non-goals)
* [Proposal](#proposal)
  * [User Stories [optional]](#user-stories)
  * [Implementation Details/Notes/Constraints [optional]](#implementation-detailsnotesconstraints-optional)
  * [Risks and Mitigations](#risks-and-mitigations)
* [Graduation Criteria](#graduation-criteria)
* [Implementation History](#implementation-history)
* [Drawbacks [optional]](#drawbacks-optional)
* [Alternatives [optional]](#alternatives-optional)

## Summary
AWS load balancer prefix annotation adds a control over the naming of the AWS ELB resources that are being generated when provisioning a Kubernetes service of type `LoadBalancer`. The current implementation provisions AWS ELB with a unique name based on the resource UID. The resulted unpredicted name makes it impossible to integrate the provisioning with existing IAM policies in situations when these two operations are controlled by two different groups. For example, IAM policies are defined and controlled by InfoSec team while provisioning of resources is under CloudOps team. The AWS IAM policies allow definition when only a prefix of the resource identifier is known. Using Kubernetes service with this annotation when it is provisioned in AWS, will allow an integration with existing IAM policies.

## Motivation
Current way of provisioning load balancer (for a Kubernetes service of the type `LoadBalancer`) is to use the service's UID and to follow Cloud  naming conventions for load balancers (for AWS it is a 32 character sequence of alphanumeric characters or hyphens that cannot begin or end with hypen [link1](https://docs.aws.amazon.com/elasticloadbalancing/2012-06-01/APIReference/API_CreateLoadBalancer.html), [link2](https://docs.aws.amazon.com/cli/latest/reference/elbv2/create-load-balancer.html)). When it is provisioned on AWS account with predefined IAM policies that limit access to ELB resources using wildcarded paths (IAM identifiers), the Kubernetes service cannot be provisioned. Providing a way to define a short known prefix to ELB resource makes it possible to match IAM policies conditions regarding the resource identifiers.

### Goals
* Support provisioning of AWS ELB resources for Kubernetes services of the type `LoadBalancer` that match AWS IAM policies
### Non-Goals
* Provide meaningful names for AWS ELB resources generated for Kubernetes services of the type `LoadBalancer`

## Proposal

### User Stories [optional]

### Implementation Details/Notes/Constraints [optional]

### Risks and Mitigations

## Graduation Criteria

## Implementation History

## Drawbacks [optional]

## Alternatives [optional]

## Infrastructure Needed [optional]