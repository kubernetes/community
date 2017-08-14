# Limiting Node Scope on the Node object

### Author: Mike Danese, (@mikedanese)

## Background

Today the node client has total authority over its own Node object. This ability
is incredibly useful for the node auto-registration flow. Some examples of
fields the kubelet self-reports in the early node object are:

1. Labels (provided by kubelet commandline)
1. Taints (provided by kubelet commandline)

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

## Proposed Solution

In many environments, we can improve the situation by centralizing reporting of
sensitive node attributes to a more trusted source and disallowing reporting of
these attributes from the kubelet.

### Label And Taint Restriction

An operator will configure a whitelist of taints and labels that nodes are
allowed to set on themselves. This list should include the taints and labels
that the kubelet is already setting on itself.

Well known taint keys:
```
node.cloudprovider.kubernetes.io/uninitialized
```

Well known label keys:

```
kubernetes.io/hostname
failure-domain.beta.kubernetes.io/zone
failure-domain.beta.kubernetes.io/region
beta.kubernetes.io/instance-type
beta.kubernetes.io/os
beta.kubernetes.io/arch
```

As well as any taints and labels that the operator is setting using:

```
  --register-with-taints
  --node-labels
```

This whitelist is passed as a command line flag to the apiserver.
NodeRestriction admission control will then prevent setting and modification by
nodes of all taints and labels with keys not in the whitelist.

### NodeRestriction Config

A new configuration API group will be created for the NodeRestriction admission
controller with the name `noderestriction.admission.k8s.io`. It will contain one
config object:

```golang
type Configuration struct {
  // AllowedLabels is a list of label keys a node is allowed to set on itself.
  // The list also supports whitelisting all label keys with a specific prefix
  // by adding an entry of the form `<prefix>*`.
  AllowedLabels []string
  // AllowedTaints is a list of taint keys a node is allowed to set on itself.
  // The list also supports whitelisting all taint keys with a specific prefix
  // by adding an entry of the form `<prefix>*`.
  AllowedTaints []string
}
```

Labels and taints that are applied by the kubelet itself (and not by
--register-with configurations) do not need to appear in this config. They are
allowed implicitly.

### NodeRestriction Config Examples

A configuration that allows all labels and all taints with prefix `insecure.`
and the `foo` taint:

```yaml
apiVersion: noderestriction.admission.k8s.io/v1
kind: Configuration
allowedLabels:
- *
allowedTaints:
- foo
- insecure.*
```

A configuration that allows only labels for CSI plugins:

```yaml
apiVersion: noderestriction.admission.k8s.io/v1
kind: Configuration
allowedLabels:
- csi.kubernetes.io.*
```

For backwards compatibility, the default config is equivalent to:

```yaml
apiVersion: noderestriction.admission.k8s.io/v1
kind: Configuration
allowedLabels:
- *
allowedTaints:
- *
```

### Removing self-delete from Node Permission

Currently a node has permission to delete itself. A node will only delete itself
when it's external name (inferred through the cloud provider) changes. This code
path will never be executated on the majority of cloud providers and this
capability undermines the usage of taints as a strong exclusion primitive.

For example, suppose an operator sets a taint `compromised` on a node that they
believe has been compromised. Currently, the compromised node could delete and
recreate itself thereby removing the `compromised` taint.

To prevent this, we will finish the removal of ExternalID which has been
deprecated since 1.1. This will allow us to remove the self delete permission
from the NodeAuthorizer.

### Taints set by central controllers

In many deployment environments, the sensitive attributes of a Node object
discussed above ("labels", "taints") are discoverable by consulting a machine
database (e.g. the GCE API). A centralized controller can register an
initializer for the node object and build the sensitive fields by consulting the
machine database. The `cloud-controller-manager` is an obvious candidate to
house such a controller.
