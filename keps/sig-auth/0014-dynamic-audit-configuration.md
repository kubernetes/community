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
last-updated: 2018-07-13
status: provisional
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
      * [Implementation Details/Notes/Constraints](#implementation-detailsnotesconstraints)
      * [Risks and Mitigations](#risks-and-mitigations)
        * [Privilege Escalation](#privilege-escalation)
        * [Webhook Authentication](#webhook-authentication)
  * [Graduation Criteria](#graduation-criteria)
  * [Implementation History](#implementation-history)
  * [Alternatives](#alternatives)
      * [Generalized Dynamic Configuration](#generalized-dynamic-configuration)

## Summary

We want to allow the advanced auditing features to be dynamically configured. Following in the same vein as [Dynamic Admission Control](https://kubernetes.io/docs/admin/extensible-admission-controllers/) we would like to provide a means of configuring the auditing features post cluster provisioning.

## Motivation

The advanced auditing features are a powerful tool, yet difficult to configure. The configuration requires deep insight into the deployment mechanism of choice and often takes many iterations to configure properly requiring a restart of the apiserver each time. Moreover, the ability to install addon tools that configure and enhance audting is hindered by the overhead in configuration. Such tools frequently run on the cluster requiring future knowledge of how to reach them when the cluster is live. These tools could enhance the security and conformance of the cluster and its applications.

### Goals
- Provide an api and set of objects to configure the advanced auditing kube-apiserver configuration dynamically

### Non-Goals
- Provide a generic interface to configure all kube-apiserver flags
- composable audit policies per-endpoint
- configuring non-webhook backends
- configuring audit output (format or per-field filtering)
- authorization of audit output

## Proposal

### Dynamic Configuration
A new dynamic audit backend will be introduced that follows suit with the existing [union backend](https://github.com/kubernetes/kubernetes/blob/master/staging/src/k8s.io/apiserver/pkg/audit/union.go). It will hold a map of configuration objects that it syncs with an informer.

#### Cluster Scoped Configuration
A cluster scoped configuration object will be provided that applies to all events in the cluster.

```golang
// ClusterAuditConfiguration represents a cluster level audit configuration
type ClusterAuditConfiguration struct {
    metav1.TypeMeta

    v1.ObjectMeta

    // Backends to send events
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

Example cluster yaml config:   
```yaml
apiVersion: audit.k8s.io/v1beta1
kind: ClusterAuditConfiguration
metadata:
  name: <name>
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

### User Stories

#### Story 1
As a cluster admin, I will easily be able to enable the interal auditing features of an existing cluster, and tweak the configurations as necessary. I want to prevent privilege escalation from being able to tamper with a root audit configuration.

#### Story 2
As a Kubernetes extension developer, I will be able to provide drop in extensions that utilize audit data.

#### Story 3
As a cluster admin, I will be able configure multiple audit-policies and webhook endpoints to provide independent auditing facilities.

### Implementation Details/Notes/Constraints

Any actions to the audit configuration objects will be hard coded to log at the `level=RequestResponse` to the previous backend and the new backend. If the apiserver is HA, the configuration will be rolled out in increments.

Inherently apiserver aggregates and HA apiserver setups will work off the same dynamic configuration object. If separate objects are needed they should be configured as static objects on the node and set through the runtime flags. Aggregated servers will implement the same audit handling mechanisms. A conformance test should be provided as assurance. This needs further discussion with the participating sigs.

### Risks and Mitigations

#### Privilege Escalation
This does open up the attack surface of the audit mechanisms. Having them strictly configured through the api server has the advantage of limiting the access of those configurations to those that have access to the master node. This opens a number of potential attack vectors:   

* privileged user changes audit policy to hide (not audit) malicious actions
* privileged user changes audit policy to DoS audit endpoint (with malintent, or ignorance)
* privileged user changes webhook configuration to hide malicious actions

As a mitigation strategy policy configured through a static file on the api server will not be accessible through the api. This file ensures that an escalation attack cannot tamper with a root configuration, but works independently of any dynamically configured objects.

#### Webhook Authentication
With Dynamic Admission control today any authentication mechanism must be provided through a static kubeconfig file on the node. This hinders a lot of the advances in this proposal. All webhooks would require authentication as an unauthenticated endpoint would allow a bad actor to push phony events. Lack of dynamic credential provisioning is problematic to the drop-in extension use case, and difficult to configure.

The reason for static configuration today is that a single configured credential would have no way of differentiating apiserver replicas or their aggregates. There is a possible mitigation by providing a bound service account token and using the calling server's dns name as the audience.

It may also be reasonable to provide a dynamic auth configuration from secrets, with the understanding that it is shared by the api servers.

This needs further discussion.

## Graduation Criteria

Success will be determined by stability of the provided mechanisms and ease of understanding for the end user.

* alpha: Api server flags can be dynamically configured, known issues are tested and resolved.
* beta: Mechanisms have been hardened against any known bugs and the process is validated by the community

## Implementation History

- 05/18/2018: initial design
- 06/13/2018: updated design

## Alternatives

### Generalized Dynamic Configuration

We could strive for all kube-apiserver flags to be able to be dynamically provisioned in a common way. This is likely a large task and out of the scope of the intentions of this feature.
