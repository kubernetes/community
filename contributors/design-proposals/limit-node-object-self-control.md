# Limiting Node Scope on the Node object

### Author: Mike Danese, (@mikedanese)

## Background

Today the node client has total authority over its own Node object. This ability
is incredibly useful for the node auto-registration flow. Some examples of
fields the kubelet self-reports in the early node object are:

1. Labels (provided by kubelet commandline)
1. Taints (provided by kubelet commandline)
1. Addresses (provided by kubelet commandline and detected from the environment)

As well as others.

## Problem

While this distributed method of registration is convenient and expedient, it
has two problems that a centralized approach would not have. Minorly, it makes
management difficult. Instead of configuring labels and taints in a centralized
place, we must configure `N` kubelet command lines. More significantly, the
approach greatly compromises security. Below are two straightforward escalations
on an initially compromised node that exhibit the attack vector.

### Capturing Dedicated Workloads

Suppose company `foo` needs to run an application that deals with PII on
dedicated nodes to comply with government regulation. A common mechanism for
implementing dedicated nodes in Kubernetes today is to set a label or taint
(e.g. `foo/dedicated=customer-info-app`) on the node and to select these
dedicated nodes in the workload controller running `customer-info-app`.

Since the nodes self reports labels upon registration, an intruder can easily
register a compromised node with label `foo/dedicated=customer-info-app`. The
scheduler will then bind `customer-info-app` to the compromised node potentially
giving the intruder easy access to the PII.

This attack also extends to secrets. Suppose company `foo` runs their outward
facing nginx on dedicated nodes to reduce exposure to the company's publicly
trusted server certificates. They use the secret mechanism to distribute the
serving certificate key. An intruder captures the dedicated nginx workload in
the same way and can now use the node certificate to read the company's serving
certificate key.

### Gaining Access to Arbitrary Serving Certificates

Suppose company `foo` uses TLS for server authentication between internal
microservices. The company uses the Kubernetes certificates API to provision
these workload certificates for workload `bar` and trust is rooted to the
cluster's root certificate authority.

When [kubelet server certificate
rotation](https://github.com/kubernetes/features/issues/267) is complete, the
same API will be used to provision serving certificates for kubelets. The design
expects to cross-reference the addresses reported in the NodeStatus with the
subject alternative names in the certificate signing request to validate the
certificate signing request.

An intruder can easily register a node with a NodeAddress `bar` and use this
certificate to MITM all traffic to service `bar` the flows through kube-proxy on
that node.

## Proposed Solution

In many environments, we can improve the situation by centralizing reporting of
these node attributes to a more trusted source and disallowing reporting of
these attributes from the kubelet.

We can scope down the initial Node object creation by moving to a centralized
controller model. In many deployment environments, the sensitive attributes of a
Node object discussed above ("labels", "taints", "addresses") are discoverable
by consulting a machine database (e.g. the GCE API). Using the
[initializer](admission_control_extension.md) mechanism, a centralized
controller can register an initializer for the node object and build the
sensitive fields by consulting the machine database. The
`cloud-controller-manager` is an obvious candidate to house such a controller.

We can scope down subsequent updates of the Node object by moving control of
sensitive fields into the NodeRestriction admission controller. For backwards
compatibility we can begin by zero'ing updates to sensitive fields and after the
kubelet compatibility window has expired, begin to return 403s. This also has
the nice property of giving us fine grained control over field/values of
updates. For example, in this model we could easily allow a kubelet to update
it's `OutOfMemory` taint and disallow the kubelet from updating its `dedicated`
taint.

In this design we assume that the infrastructure has already validated caller ID
(e.g. through the process of TLS bootstrap) and that we can trust that the node
client is who they say they are.
