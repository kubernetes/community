---
kep-number: 14
title: Dynamic Audit Configuration
authors:
  - "@pbarker"
owning-sig: sig-auth
participating-sigs:
  - sig-api-machinery
reviewers:
  - "@tallclair"
  - "@yliaog"
  - "@caesarxuchao"
  - "@liggitt"
approvers:
  - "@tallclair"
  - "@liggitt"
  - "@yliaog"
editor: TBD
creation-date: 2018-05-18
last-updated: 2018-07-31
status: implementable
---

# Dynamic Audit Control

## Table of Contents

* [Dynamic Audit Control](#dynamic-audit-control)
  * [Table of Contents](#table-of-contents)
  * [Summary](#summary)
  * [Motivation](#motivation)
      * [Goals](#goals)
      * [Non-Goals](#non-goals)
  * [Proposal](#proposal)
      * [Dynamic Configuration](#dynamic-configuration)
        * [Cluster Scoped Configuration](#cluster-scoped-configuration)
      * [User Stories](#user-stories)
        * [Story 1](#story-1)
        * [Story 2](#story-2)
        * [Story 3](#story-3)
        * [Story 4](#story-4)
      * [Implementation Details/Notes/Constraints](#implementation-detailsnotesconstraints)
        * [Feature Gating](#feature-gating)
        * [Policy Enforcement](#policy-enforcement)
        * [Aggregated Servers](#aggregated-servers)
      * [Risks and Mitigations](#risks-and-mitigations)
        * [Privilege Escalation](#privilege-escalation)
        * [Leaked Resources](#leaked-resources)
        * [Webhook Authentication](#webhook-authentication)
        * [Performance](#performance)
  * [Graduation Criteria](#graduation-criteria)
  * [Implementation History](#implementation-history)
  * [Alternatives](#alternatives)
      * [Generalized Dynamic Configuration](#generalized-dynamic-configuration)
      * [Policy Override](#policy-override)

## Summary

We want to allow the advanced auditing features to be dynamically configured. Following in the same vein as 
[Dynamic Admission Control](https://kubernetes.io/docs/admin/extensible-admission-controllers/) we would like to provide 
a means of configuring the auditing features post cluster provisioning.

## Motivation

The advanced auditing features are a powerful tool, yet difficult to configure. The configuration requires deep insight 
into the deployment mechanism of choice and often takes many iterations to configure properly requiring a restart of 
the apiserver each time. Moreover, the ability to install addon tools that configure and enhance auditing is hindered 
by the overhead in configuration. Such tools frequently run on the cluster requiring future knowledge of how to reach 
them when the cluster is live. These tools could enhance the security and conformance of the cluster and its applications.

### Goals
- Provide an api and set of objects to configure the advanced auditing kube-apiserver configuration dynamically

### Non-Goals
- Provide a generic interface to configure all kube-apiserver flags
- configuring non-webhook backends
- configuring audit output (format or per-field filtering)
- authorization of audit output

## Proposal

### Dynamic Configuration
A new dynamic audit backend will be introduced that follows suit with the existing [union backend](https://github.com/kubernetes/kubernetes/blob/master/staging/src/k8s.io/apiserver/pkg/audit/union.go). It will hold a map of configuration objects that it syncs with an informer.

#### Cluster Scoped Configuration
A cluster scoped configuration object will be provided that applies to all events in the cluster.

```golang
// AuditConfiguration represents a dynamic audit configuration
type AuditConfiguration struct {
    metav1.TypeMeta

    v1.ObjectMeta

    // Policy is the current audit v1beta1 Policy object
    // if undefined it will default to the statically configured cluster policy if available
    // if neither exist the backend will fail
    Policy *Policy

    // Backend to send events
    Backend *Backend
}

// Backend holds the configuration for the backend
type Backend struct {
    // Webhook holds the webhook backend
    Webhook *WebhookBackend
}

// WebhookBackend holds the configuration of the webhooks
type WebhookBackend struct {
    // InitialBackoff is amount of time to wait before retrying the first failed request in seconds
    InitialBackoff *int

    // ThrottleBurst is the maximum number of events sent at the same moment
    ThrottleBurst *int

    // ThrottleEnabled determines whether throttling is enabled
    ThrottleEnabled *bool

    // ThrottleQPS maximum number of batches per second
    ThrottleQPS *float32

    // ClientConfig holds the connection parameters for the webhook
    ClientConfig WebhookClientConfig
}

// WebhookClientConfig contains the information to make a TLS
// connection with the webhook; this follows: 
// https://github.com/kubernetes/api/blob/master/admissionregistration/v1beta1/types.go#L222
// but may require some additive auth parameters
type WebhookClientConfig struct {
    // URL of the server
    URL *string

    // Service name to send to
    Service *ServiceReference

    // `caBundle` is a PEM encoded CA bundle which will be used to validate
    // the webhook's server certificate.
    CABundle []byte
}
```

Multiple definitions can exist as independent solutions. These updates will require the audit API to be registered with the apiserver. The dynamic configurations will be wrapped by truncate and batch options, which are set statically through existing flags. Dynamic configuration will be enabled by a feature gate for pre-stable releases. If existing flags are provided to configure the audit backend they will be taken as a separate backend configuration.

Example configuration yaml config:   
```yaml
apiVersion: audit.k8s.io/v1beta1
kind: AuditConfiguration
metadata:
  name: <name>
policy:
  rules:
  - level: <level>
  omitStages:
  - stage: <stage>
backend:
  webhook:
  - initialBackoff: <10s>
    throttleBurst: <15>
    throttleEnabled: <true>
    throttleQPS: <10>
    clientConfig:
      url: <backend url>
      service: <optional service name>
      caBundle: <ca bundle>
```
A configuration flag will be added that enables dynamic auditing `--audit-dynamic-configuration`, which will default to false.

### User Stories

#### Story 1
As a cluster admin, I will easily be able to enable the internal auditing features of an existing cluster, and tweak the configurations as necessary. I want to prevent privilege escalation from being able to tamper with a root audit configuration.

#### Story 2
As a Kubernetes extension developer, I will be able to provide drop in extensions that utilize audit data.

#### Story 3
As a cluster admin, I will be able configure multiple audit-policies and webhook endpoints to provide independent auditing facilities.

#### Story 4
As a kubernetes developer, I will be able to quickly turn up the audit level on a certain area to debug my application.

### Implementation Details/Notes/Constraints

#### Feature Gating
Introduction of dynamic policy requires changes to the current audit pipeline. Care must be taken that these changes are 
properly gated and do not affect the stability or performance of the current features as they progress to GA. A new decorated 
handler will be provisioned similar to the [existing handlers](https://github.com/kubernetes/apiserver/blob/master/pkg/endpoints/filters/audit.go#L41) 
called `withDynamicAudit`. Another conditional clause will be added where the handlers are 
[provisioned](https://github.com/kubernetes/apiserver/blob/master/pkg/server/config.go#L536) allowing for the proper feature gating.

#### Policy Enforcement
This addition will move policy enforcement from the main handler to the backends. From the `withDynamicAudit` handler, 
the full event will be generated and then passed to the backends. Each backend will copy the event and then be required to 
drop any pieces that do not conform to its policy. A new sink interface will be required for these changes called `EnforcedSink`, 
this will largely follow suite with the existing sink but take a fully formed event and the authorizer attributes as its 
parameters. It will then utilize the `LevelAndStages` method in the policy 
[checker](https://github.com/kubernetes/apiserver/blob/master/pkg/audit/policy/checker.go) to enforce its policy on the event, 
and drop any unneeded sections. The new dynamic backend will implement the `EnforcedSink` interface, and update its state 
based on a shared informer. For the existing backends to comply, a method will be added that implements the `EnforcedSink` interface.

Implementing the [attribute interface](https://github.com/kubernetes/kubernetes/tree/master/staging/src/k8s.io/apiserver/pkg/authorization/authorizer/interfaces.go) 
based on the Event struct was also explored. This would allow us to keep the existing `Sink` interfaces, however it would 
require parsing the request URI twice in the pipeline due to how that field is represented in the Event. This was determined 
to not be worth the cost.

#### Aggregated Servers
Inherently apiserver aggregates and HA apiserver setups will work off the same dynamic configuration object. If separate 
audit configuration objects are needed they should be configured as static objects on the node and set through the runtime flags. Aggregated servers will implement the same audit handling mechanisms. A conformance test should be provided as assurance. Metadata level 
logging will happen by default at the main api server as it proxies the traffic. The aggregated server will then watch the same 
configuration objects and only log on resource types that it handles. This will duplicate the events sent to the receiving servers 
so they should not expect to key off `{ Audit-ID x Stage }`.

### Risks and Mitigations

#### Privilege Escalation
This does open up the attack surface of the audit mechanisms. Having them strictly configured through the api server has the advantage of limiting the access of those configurations to those that have access to the master node. This opens a number of potential attack vectors:   

* privileged user changes audit policy to hide (not audit) malicious actions
* privileged user changes audit policy to DoS audit endpoint (with malintent, or ignorance)
* privileged user changes webhook configuration to hide malicious actions

As a mitigation strategy policy configured through a static file on the api server will not be accessible through the api. This file ensures that an escalation attack cannot tamper with a root configuration, but works independently of any dynamically configured objects.

#### Leaked Resources
A user with permissions to create audit policies effectively has read access to the entire cluster (including all secrets data).

A mitigation strategy will be to document the exposure space granted with this resource. Advice will be provided to only allow access to cluster admin level roles.

#### Webhook Authentication
With Dynamic Admission control today any authentication mechanism must be provided through a static kubeconfig file on the node. This hinders a lot of the advances in this proposal. All webhooks would require authentication as an unauthenticated endpoint would allow a bad actor to push phony events. Lack of dynamic credential provisioning is problematic to the drop-in extension use case, and difficult to configure.

The reason for static configuration today is that a single configured credential would have no way of differentiating apiserver replicas or their aggregates. There is a possible mitigation by providing a bound service account token and using the calling server's dns name as the audience.

It may also be reasonable to provide a dynamic auth configuration from secrets, with the understanding that it is shared by the api servers.

This needs further discussion.

#### Performance

These changes will likely have an O(n) performance impact on the api server per policy.  A `DeepCopy` of the event will be 
required for each backend. Also, the request/response object would now be serialized on every [request](https://github.com/kubernetes/kubernetes/blob/cef2d325ee1be894e883d63013f75cfac5cb1246/staging/src/k8s.io/apiserver/pkg/audit/request.go#L150-L152). 
Benchmark testing will be required to understand the scope of the impact and what optimizations may be required. This impact 
is gated by opt-in feature flags, which allows it to move to alpha but these concerns must be tested and reconciled before it 
progresses to beta.

## Graduation Criteria

Success will be determined by stability of the provided mechanisms and ease of understanding for the end user.

* alpha: Api server flags can be dynamically configured, known issues are tested and resolved.
* beta: Mechanisms have been hardened against any known bugs and the process is validated by the community

## Implementation History

- 05/18/2018: initial design
- 06/13/2018: updated design
- 07/31/2018: dynamic policy addition

## Alternatives

### Generalized Dynamic Configuration

We could strive for all kube-apiserver flags to be able to be dynamically provisioned in a common way. This is likely a large 
task and out of the scope of the intentions of this feature.

### Policy Override

There has been discussion over whether the policy configured by api server flags should limit the policies configured dynamically. 
This would allow a cluster admin to narrowly define what is allowed to be logged by the dynamic configurations. While this has upsides 
it was ruled out for the following reasons: 

* It would limit user story #4 in the ability to quickly turn up logging when needed 
* It could prove difficult to understand as the policies themselves are fairly complex 
* The use of CRDs would be difficult to bound

The dynamic policy feature is gated by runtime flags. This still provides the cluster provisioner a means to limit audit logging to the 
single runtime object if needed.