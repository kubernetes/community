# Kubernetes Network Custom Resource Definition De-facto Standard Version 1

Network Plumbing Working Group
Approved August 02, 2018

**1. Goals**

This document proposes a specification of the requirements and procedures for attaching Kubernetes pods to one or more logical or physical networks, including requirements for plugins using the Container Network Interface (CNI) to attach pod networks.

**1.1 Non-Goals of Version 1**

This document specifically does not address certain issues for reasons of simplicity and/or the need to achieve some reasonable consensus.  These issues may be addressed in future versions of this specification.

1.1.1 Scheduling and resource management

Efforts to integrate scheduling and resource management (eg ensuring that a node is not assigned more pods than there are network resources available for) are underway in the Resource Management Working Group in which various members of this working group and SIG Network are involved.  Future versions of this specification may incorporate aspects of resource management.

1.1.2 Changes to the Kubernetes API

This specification explicitly avoids changes to the existing Kubernetes API. Those changes require a much longer upstream process, but the hope is that this document can serve as a basis for those changes in the future by proving the concepts and providing one or more real-world implementations of multiple pod network attachments.

1.1.3 Interaction with the Kubernetes API

The interaction of additional pod network attachments with the Kubernetes API and its objects, such as Services, Endpoints, proxies, etc is not specified. SIG Network discussed this topic at length in July/August 2017 and decided it would require changes to the Kubernetes API, which is an explicit non-goal of this specification.

1.1.4 Changes to the Kubernetes CNI Driver

To ensure easy use of this specification and implementations of it, no changes will be required of the Kubernetes CNI driver (eg pkg/kubelet/dockershim/network/cni/*).  Changes to the driver to support multiple pod network attachments have already been proposed and multiple proof-of-concepts written, but there is as yet no upstream consensus on how multiple attachments should work and thus what changes would be required in the CNI driver.

1.1.5 Replacement of the Kubernetes Cluster-Wide Default Network

To preserve compatibility and expectations with existing Kubernetes deployments, this specification does not attempt to change any of the existing Kubernetes cluster-wide network behavior.  All pods must still be attached to the cluster-wide default network as they are today.

1.1.6 Enabling Implementations That Are Not CNI Plugins

This specification only attempts to enable Kubernetes network plugins which use the CNI driver of a Kubernetes runtime like dockershim, CRI-O, and rkt.  This specification expects pod network attachment operations to be directed through the implementation, which is required to be a CNI plugin.  Future versions of this specification may revisit this requirement.

**2. Definitions**

2.1 Implementation

The Kubernetes network plugin that implements this specification; as of Kubernetes 1.11 this is defined to be a plugin conforming to the [Container Network Interface](https://github.com/containernetworking/cni/blob/master/SPEC.md) (CNI) specification v0.1.0 or later.  Kubernetes should be configured to call the implementation for all pod network operations, and the implementation then determines which additional operations to perform based on the pods annotations and the Custom Resources defined in this specification.

2.2 Kubernetes Cluster-Wide Default Network

A network to which all pods are attached following the current behavior and requirements of Kubernetes.

2.3 Network Attachment

A means of allowing a pod to directly communicate with a given logical or physical network. Typically (but not necessarily) each attachment takes the form of a kernel network interface placed in to the pods network namespace.  Each attachment may result in zero or more IP addresses being assigned to the pod.

2.4 CNI Delegating Plugin

An implementation conforming to this specification which delegates pod network attachment/detachment operations to other plugins conforming to the CNI specification.  Examples include [Multus](https://github.com/intel/multus-cni) and [CNI-Genie](https://github.com/Huawei-PaaS/CNI-Genie).

This specification places CNI Delegating Plugin specific requirements and recommendations under sections marked as such.  If the implementation is not a CNI Delegating Plugin, it is free to ignore the requirements and recommendations in such marked sections, and any subsections of them.

**3. NetworkAttachmentDefinition Object**

This specification defines a `NetworkAttachmentDefinition` Custom Resource object which describes how to attach a pod to the logical or physical network referenced by the object.

3.1 Custom Resource Definition (CRD)

The CRD tells Kubernetes API how to expose `NetworkAttachmentDefinition` objects. See below for the definition of the object itself.

```
apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  name: network-attachment-definitions.k8s.cni.cncf.io
spec:
  group: k8s.cni.cncf.io
  version: v1
  scope: Namespaced
  names:
        plural: network-attachment-definitions
        singular: network-attachment-definition
        kind: NetworkAttachmentDefinition
        shortNames:
        - net-attach-def
  validation:
    openAPIV3Schema:
      properties:
        spec:
          properties:
            config:
              type: string
```

3.2 NetworkAttachmentDefinition Object Definition

The `NetworkAttachmentDefinition` object itself contains only a "spec" section. Its definition (in Go form) shall be:

```
type NetworkAttachmentDefinition struct {
    metav1.TypeMeta
    // Note that ObjectMeta is mandatory, as an object
    // name is required
    metav1.ObjectMeta
    // Specification describing how to add or remove network
    // attachments for a Pod. In the absence of valid keys in
    // the Spec field, the implementation shall attach/detach an
    // implementation-known network referenced by the objects
    // name.
    // +optional
    Spec NetworkAttachmentDefinitionSpec `json:"spec"`
}

type NetworkAttachmentDefinitionSpec struct {
    // Config contains a standard JSON-encoded CNI configuration
    // or configuration list which defines the plugin chain to
    // execute. The CNI configuration may omit the name field
    // which will be populated by the implementation when the
    // Config is passed to CNI delegate plugins.
    // +optional
    Config string `json:"config,omitempty"`
}
```

Plugins which are not CNI Delegating Plugins may not use the Spec.Config field to store arbitrary configuration.  Instead they should store their non-standard configuration in annotations on the pod object.

3.2.1 YAML Example: CNI config JSON in object

```
apiVersion: "k8s.cni.cncf.io/v1"
kind: NetworkAttachmentDefinition
metadata:
  name: a-bridge-network
spec:
  config: {
    "cniVersion": "0.3.0",
    "name": "a-bridge-network",
    "type": "bridge",
    "bridge": "br0",
    "isGateway": true,
    "ipam": {
      "type": "host-local",
      "subnet": "192.168.5.0/24",
      "dataDir": "/mnt/cluster-ipam"
    }
}
```

3.2.2 YAML Example: CNI configlist JSON in object

```
apiVersion: "k8s.cni.cncf.io/v1"
kind: NetworkAttachmentDefinition
metadata:
  name: another-bridge-network
spec:
  config: {
  "cniVersion": "0.3.0",
  "name": "another-bridge-network",
  "plugins": [
    {
      "type": "bridge",
      "bridge": "br0",
      "ipam": {
        "type": "host-local",
        "subnet": "192.168.5.0/24"
      }
    },
    {
      "type": "port-forwarding"
    },
    {
      "type": "tuning",
      "sysctl": {
        "net.ipv4.conf.all.log_martians": "1"
      }
    }
  ]
}
```

3.2.3 YAML Example: Limited CNI config required ("thick" plugin)

```
apiVersion: "k8s.cni.cncf.io/v1"
kind: NetworkAttachmentDefinition
metadata:
  name: a-bridge-network
spec:
  config: {
    "cniVersion": "0.3.0",
    "type": "awesome-plugin"
}
```

3.2.4: YAML Example: Implementation-specific Network Reference

```
apiVersion: "k8s.cni.cncf.io/v1"
kind: NetworkAttachmentDefinition
metadata:
  name: a-bridge-network
```

3.3 NetworkAttachmentDefinition Object Naming Rules

Valid `NetworkAttachmentDefinition` object names must be comprised of units of the DNS-1123 label format, in accordance with how Kubernetes [validates namespace names](https://github.com/kubernetes/kubernetes/blob/8d7d7a5e0d4d7e75f5a860574346944b8cc0fc43/staging/src/k8s.io/apimachinery/pkg/util/validation/validation.go#L107-L124). It is recommended that each unit of DNS-1123 label be 63 characters or less.

`"DNS-1123 label must consist of lower case alphanumeric characters or -, and must start and end with an alphanumeric character"` -_Kubernetes Error Message_

Labels may be validated by matching this regular expression:

`[a-z0-9]([-a-z0-9]*[a-z0-9])?`

3.3.1 Examples

```
example-namespace/attachment-name
attachment-name
```

3.4 CNI Delegating Plugin Requirements

For implementations that are CNI Delegating Plugins, the implementation "delegates" the actual attachment/detachment to one or more additional CNI plugins. The `NetworkAttachmentDefinition` object contains the  necessary information to determine which CNI plugins to execute and which options to pass to them.

3.4.1 Determining CNI Plugins for a NetworkAttachmentDefinition Object

The CNI Delegating Plugin must use the following rules, in the listed order, to determine which CNI plugin(s) to execute for a pods attachment to a given network described by a `NetworkAttachmentDefinition` object:

1. If a "config" key is present in the Spec the contents of the keys value is used to execute plugins per the CNI specification
2. If a CNI .configlist file is present on-disk whose JSON "name" key matches the name of the `NetworkAttachmentDefinition` object, its contents is loaded and used to execute plugins per the CNI specification.
3. If a CNI .config file is present on-disk whose JSON "name" key matches the name of the `NetworkAttachmentDefinition` object, its contents is loaded and used to execute plugins per the CNI specification.
4. Otherwise, the network request must fail

3.4.2 Spec.Config and the CNI JSON name Field

If the Spec.Config key is valid but its data omits the CNI JSON name field the CNI Delegating Plugin shall add a name field containing the `NetworkAttachmentDefinition` objects name to the CNI JSON before sending that CNI JSON configuration to a delegate plugin. This is intended to allow brevity of configuration for the "thick" plugin case where the delegate plugin requires minimal configuration.

3.4.2.1 "Name" Injection Example

Given the following `NetworkAttachmentDefinition` object:

```
apiVersion: "k8s.cni.cncf.io/v1"
kind: NetworkAttachmentDefinition
metadata:
  name: a-bridge-network
spec:
  config: {
    "cniVersion": "0.3.0",
    "type": "awesome-plugin"
  }
```

The CNI Delegating Plugin would send the following CNI JSON Configuration to the awesome-plugin binary, generated by injecting a name field with the `NetworkAttachmentDefinition` object name as its data:

```
{
    "cniVersion": "0.3.0",
    "name": "a-bridge-network",
    "type": "awesome-plugin"
}
```

3.4.3 NetworkAttachmentDefinition Object and CNI JSON Configuration Naming Considerations

Per Kubernetes requirements, each `NetworkAttachmentDefinition` object must have a name. The CNI specification strongly recommends that CNI JSON configuration include a name and future CNI specification versions will require a name.

This specification strongly recommends that the `NetworkAttachmentDefinition` object name match the "name" key in CNI JSON configuration referenced by the `NetworkAttachmentDefinition` object (whether stored on-disk or in the NetworkAttachmentDefinition.Spec.Config key). This reduces user confusion and makes the mapping between `NetworkAttachmentDefinition` object and CNI configuration clearer.

While strongly recommended, this matching is not required as it may place Kubernetes API object naming requirements on externally-defined resources. Naming is ultimately left up to the cluster administrator to decide.

**4. Network Attachment Selection Annotation**

To select one or more secondary ("sidecar") networks to which a pod should be attached, this specification defines a Pod object annotation. Because attachments selected by this annotation are secondary attachments, these network attachments are not known to Kubernetes itself and information about them may not be available through the standard Kubernetes API.

The Network Attachment Selection Annotation may be used to select additional attachments to the cluster-wide default network, above and beyond the required initial cluster-wide default network attachment.

4.1 Annotation Name and Format

The Pod object annotation name shall be `k8s.v1.cni.cncf.io/networks`. The annotation value shall be specified in one of two possible formats, described below. Implementations of this specification must support both formats.

Note that even though the objects the annotation references are `NetworkAttachmentDefinition` objects, the annotations name is `k8s.v1.cni.cncf.io/networks`.  This is intentional.

4.1.1 Comma-delimited Format

This format is intended to be a very simple, user-friendly format, composed only of `NetworkAttachmentDefinition` object references separated by commas.  The format of each object reference is either (a) <NetworkAttachmentDefinition object name> to reference a `NetworkAttachmentDefinition` object in the pods namespace, or (b) <NetworkAttachmentDefinition object namespace>/<NetworkAttachmentDefinition object name> to reference a `NetworkAttachmentDefinition` object in a different namespace.

```
kind: Pod
metadata:
  name: my-pod
  namespace: my-namespace
  annotations:
    k8s.v1.cni.cncf.io/networks: net-a,net-b,other-ns/net-c
```

4.1.2 JSON List Format

This format allows users to give pod-specific requirements for the network attachment.  For example, these could include a specific IP address or MAC address if those options are supported by the plugin that ultimately handles the network attachment, whether that is the implementation itself or a delegate plugin.

```
kind: Pod
metadata:
  name: my-pod
  namespace: my-namespace
  annotations:
    k8s.v1.cni.cncf.io/networks: |
      [
        {"name":"net-a"},
        {
          "name":"net-b",
          "ips": ["1.2.3.4"],
          "mac": "aa:bb:cc:dd:ee:ff"
        },
        {
          "name":"net-c",
          "namespace":"other-ns"
        }
      ]
```

4.1.2.1 JSON List Format Key Definitions

The following keys are defined for each network attachment map in this formats networks list. All key names that do not include a period character are reserved to ensure this specification may be extended in the future. Implementations that write keys other than those defined in this specification must use reverse domain name notation (eg "org.foo.bar.key-name") to name the non-standard keys.

4.1.2.1.1 "name"

This required key with value of type string is the name of a `NetworkAttachmentDefinition` object, either in the Pods namespace (if the "namespace" key is missing or empty) or another namespace specified by the "namespace" key.

4.1.2.1.2 "namespace"

This optional key with value of type string is the namespace of the `NetworkAttachmentDefinition` object named by the "name" key.

4.1.2.1.3 "ips"

This optional key with value of type string-array requires the plugin handling this network attachment to assign the given IP addresses to the pod. This keys value must contain at least one array element and each element must be a valid IPv4 or IPv6 address. If the value is invalid, the Network Attachment Selection Annotation shall be invalid and ignored by the implementation.

4.1.2.1.3.1 "ips" Example

```
  annotations:
     k8s.v1.cni.cncf.io/networks: |
      [
        {
          "name":"net-b",
          "ips": ["10.2.2.42", "2001:db8::5"]
        }
      ]
```

4.1.2.1.3.2 CNI Delegating Plugin Requirements

The CNI Delegating Plugin must add an "ips" key to the CNI "args" map and set its value to a transformation of the "ips" keys value conforming to the "ips" key value as described in CNIs [CONVENTIONS.md](https://github.com/containernetworking/cni/blob/master/CONVENTIONS.md#args-in-network-config). Since this specification requires that "ips" be honored by the implementation, but CNI "args" may be ignored by plugins, the implementation must ensure that the requested IP addresses were assigned to the network attachments interface in the returned CNI Request structure; if it has not been assigned the implementation must fail the network attachment.

Given the immediately preceding Network Attachment Selection Annotation example, the CNI Delegating Plugin would transform the data into the following CNI JSON config snippet passed to each plugin in the CNI invocation for "net-b":

```
{
  ...
  "args":{
    "cni":{
      "ips": ["10.2.2.42", "2001:db8::5"]
    }
  }
  ...
}
```

4.1.2.1.4 "mac"

This optional key with value type string requires the plugin handling this network attachment to assign the given MAC address to the pod. This keys value must contain a valid 6-byte Ethernet MAC address or a valid 20-byte IP-over-InfiniBand Hardware address (as described in [RFC4391 section 9.1.1](https://tools.ietf.org/html/rfc4391#section-9.1.1)). If the value is invalid, the Network Attachment Selection Annotation shall be invalid and ignored by the implementation.

4.1.2.1.4.1 "mac" Example

Given the following Network Attachment Selection Annotation for a pod:

```
 annotations:
     k8s.v1.cni.cncf.io/networks: |
      [
        {
          "name":"net-b",
          "mac": "02:23:45:67:89:01"
        }
      ]
```

4.1.2.1.4.2 CNI Delegating Plugin Requirements

The implementation must add an "mac" key to the CNI "args" map (as described in CNIs [CONVENTIONS.md](https://github.com/containernetworking/cni/blob/master/CONVENTIONS.md#args-in-network-config)) and set its value to a transformation of the "mac" keys value. Since this specification requires that "mac" be honored by the implementation, but CNI "args" may be ignored by plugins, the implementation must ensure that the requested MAC address was assigned to the network attachments interface in the returned CNI Request structure; if it has not been assigned the implementation must fail the network attachment.

Given the immediately preceding Network Attachment Selection Annotation example, the CNI Delegating Plugin would transform the data into the following CNI JSON config snippet passed to each plugin in the CNI invocation for "net-b":

```
{
  ...
  "args":{
    "cni":{
      "mac": "02:23:45:67:89:01"
    }
  }
  ...
}
```

4.1.2.1.5 "interface"

This optional key with value of type string requires the implementation to use the given name for the pod interface resulting from this network attachment. This keys value must be a valid Linux kernel interface name. If the value is invalid, the Network Attachment Selection Annotation shall be invalid and ignored by the implementation.

If the requested interface name is already used by a previous network attachment, the implementation must fail the current network attachment.

4.1.2.1.5.1 CNI Delegating Plugin Requirements

An "interface" key requires the CNI Delegating Plugin to set the CNI_IFNAME environment variable to the given value when invoking CNI plugins for this network attachment.

4.2 Multiple Attachments to the Same Network

Pods may request attachment to the same network multiple times via the Network Attachment Selection Annotation. Each of these requests is considered a separate attachment, must be processed by the implementation as a separate operation, and must result in a separate Network Attachment Status Annotation entry.

4.2.1 CNI Delegating Plugin Requirements

Each network reference in the Network Attachment Selection Annotation corresponds to a CNI plugin/configlist invocation described by a `NetworkAttachmentDefinition` object.  As the CNI specification defines a network attachment as a unique tuple of [container ID, network name,CNI_IFNAME], the CNI Delegating Plugin must ensure that all CNI operations (eg ADD or DEL) for a given network attachment use the same unique tuple, which should be created as follows:

1. Container ID: given by the runtime
2. Network Name: present in (or found via) the `NetworkAttachmentDefinition` object
3. CNI_IFNAME: if not otherwise specified by the Network Attachment Selection Annotation, generated by the CNI Delegating Plugin for a given attachment, and must be unique across all attachments for the given [Container ID, network name] tuple.

4.2.2 Example

```
  annotations:
    k8s.v1.cni.cncf.io/networks: net-a,net-a
```

In this example the implementation must attach "net-a" twice and each attachment would result in a separate entry in the Network Attachment Status Annotation list.

**5. Network Attachment Status Annotation**

To ensure that the result of a network attachment is available via the Kubernetes API, the implementation may publish the result of the attachment operation to an annotation on the pod object that requested the attachment.

The annotations name shall be `k8s.v1.cni.cncf.io/network-status` and its value shall be a JSON-encoded list of maps. Each element in the list shall be a map composed from the result of a network attachment operation as described below.

5.1 Source of Status Information

A network attachment operations status map shall contain information taken from the result of the attachment operation for the given network. Status information that is only useful inside the pod itself (like IP routes) does not need to be part of the status map as this information is not generally relevant to Kubernetes API clients.

5.1.1 CNI Delegating Plugin Requirements

Status for the attachment shall be taken from the first sandbox interface of the CNI Result object for that attachments CNI ADD or GET invocation.  If the ADD or GET invocation does not return a result (which is currently allowed by the CNI specification) the implementation should add a minimal status map as described below.

The implementation should use as much information provided in the CNI Result object as possible to construct the Network Attachment Status Annotation.  If the operation returns a CNI Result object of version 0.3.0 or later, the "interfaces", "ips", and "dns" fields can be used to easily construct the Network Attachment Status Annotation.  If the operation returns a CNI Result object less than version 0.3.0, the implementation should construct as much of the Network Attachment Status Annotation as it can from the CNI Result, and may populate the remaining Annotation fields (eg interface and mac) by whatever means it wishes to.

5.2 Cluster-Wide Default Network Entry

Since the implementation is required to attach the pod to the cluster-wide default network, this network must have an entry in the status map even though it is not specified in the Network Attachment Selection Annotation and may not have a corresponding `NetworkAttachmentDefinition` object. The entry shall set its default key to true. The default entry may be at any location in the status list.

5.3 Status Map Key Definitions

The following keys are defined for each network attachments status map.  All key names that do not include a period character are reserved to ensure this specification may be extended in the future. Implementations that write keys other than those defined in this specification must use reverse domain name notation (eg "org.foo.bar.key-name") to name the non-standard keys.

5.3.1 "name"

This required keys value (type string) shall contain either a `NetworkAttachmentDefinition` object name from the pods Network Attachment Selection Annotation, or the name of the cluster-wide default network. The "name" may contain a namespace reference as defined in Section 4.1.1. [ref 2018-02-01 meeting @22:30]

5.3.2 "interface"

This optional keys value (type string) shall contain the network interface name in the pods network namespace corresponding to the network attachment.

5.3.2.1 CNI Delegating Plugin Requirements

For a CNI Result of version 0.3.0 or later, the "interface" key should be sourced from the first element of the CNI Results "interfaces" property which has a valid "sandbox" property.

For a CNI Result version earlier than 0.3.0, the implementation may populate the field from the CNI_IFNAME environment variable it set for the network attachment operation, or leave the field blank.

5.3.3 "ips"

This optional keys value (type string array) shall contain an array of IPv4 and/or IPv6 addresses assigned to the pod as a result of the attachment operation.

5.3.3.1 CNI Delegating Plugin Requirements

For a CNI Result of version 0.3.0 or later, the "ips" key should be sourced from the CNI Results "ips" property if either of the following are true; otherwise it should not be included in the status map.

1. The "ips" key should be taken from the element of the CNI Result objects "ips" list where the "interface" index refers to the first element in the CNI Result objects "interfaces" list with a valid "sandbox" key.
2. If there are no elements in the CNI Result objects "interfaces" list, or none of the interfaces in the CNI Result objects "interfaces" list have valid "sandbox" properties, the "ips" key should be taken from the first element of the CNI Result objects "ips" list which does not have an "interface" property or where the "interface" property is less than zero.

The object of these requirements is to ensure that the IP address included in the status map is the same IP address assigned to the attachments interface inside the pod, not the address of non-sandbox interfaces which are sometimes reported by CNI plugins.

For a CNI Result of a version earlier than 0.3.0, the "ips" key should be sourced from a combination of the CNI Results "ip4" and "ip6" properties.

5.3.4 "mac"

This optional keys value (type string) shall contain the hardware address of the network interface named by the "interface" key. If the "mac" key is present, "interface" must also be present.

5.3.4.1 CNI Delegating Plugin Requirements

For a CNI Result of version 0.3.0 or later, the "mac" key should be sourced from the first element of the CNI Results "interfaces" property which has a valid "sandbox" property.

For a CNI Result version earlier than 0.3.0, the implementation may populate the field by an implementation-specific mechanism, or leave the field blank.

5.3.5 "default"

This required keys value (type boolean) shall indicate that this attachment is the result of the cluster-wide default network. Only one element in the Network Attachment Status Annotation list may have the "default" key set to true.

5.3.6 "dns"

This optional keys value (type map) shall contain DNS information that is gathered as a result of the network attachment.  The map may contain the following keys.

5.3.6.1 "nameservers"

This optional keys value (type string array) shall contain an array of IPv4 and/or IPv6 addresses of DNS servers.

5.3.6.2 "domain"

This optional keys value (type string) shall contain the local domain name for the network attachment..

5.3.6.3 "search"

This optional keys value (type string array) shall contain the array of DNS search names for the network attachment.

5.4 Example

**```**

kind: Pod

metadata:

  name: my-pod

  namespace: my-namespace

  annotations:

    k8s.v1.cni.cncf.io/network-status: |

      [

        {

          "name": "cluster-wide-default",

          "interface": "eth5",

          "ips": ["1.2.3.1/24", "2001:abba::2230/64"],

          "mac": "02:11:22:33:44:54",

          "default": true

        },

        {

          "name": "some-network",

          "interface": "eth1",

          "ips": ["1.2.3.4/24", "2001:abba::2234/64"],

          "mac": "02:11:22:33:44:55",

          "dns": {

            "nameservers": ["4.2.2.1", "2001:4860:4860::8888"],

            "search": ["eng.foobar.com", "foobar.com"]

          },

          "default": false

        },

        {

          "name": "other-ns/an-ip-over-infiniband-network",

          "interface": "ib0",

          "ips": ["5.4.3.2/16"],

          "mac": "80:00:11:22:33:44:55:66:77:88:99:aa:bb:cc:dd:ee:ff:00:11:22",

          "default": false

        }

      ]

**```**

**6. Cluster-Wide Default Network**

The implementation must attach every pod to the cluster-wide default network, keeping the existing Kubernetes pod networking behavior. [ref [2018-01-18@31:30](https://youtu.be/_iK-DNJmzY0?list=PL69nYSiGNLP2E8vmnqo5MwPOY25sDWIxb&t=1885)] The implementation is free to define how the cluster-wide default network is configured and referenced (for example, as a `NetworkAttachmentDefinition` object, an on-disk CNI JSON configuration file, or some other means), provided that all pods are first attached to this network.

6.1 Network Plugin Readiness

The implementation should not write its CNI JSON configuration file to the Kubernetes CNI configuration directory until the cluster-wide default network is ready.  This prevents Kubernetes from scheduling pods on the node that will immediately fail because the implementation has signalled readiness but the cluster-wide default network has not.

The implementation is free to define how cluster-wide default network readiness is determined, provided that when the cluster-wide default network is determined to be ready, a pod may be immediately attached to that network and have a reasonable expectation of network connectivity.

6.1.1 CNI Delegating Plugin Recommendations

To prevent race conditions between the cluster-wide default network, the implementation, and kubelets CNI plugin loader, it is recommended that kubelet be configured with an implementation-specific CNI configuration directory through the `--cni-conf-dir` command-line option. The implementation should wait until the cluster-wide default network plugins CNI JSON configuration is written to `/etc/cni/net.d` and then write its own CNI JSON configuration file to the implementation-specific CNI configuration directory given to kubelet.

6.1.2 Alternate Readiness Method

If it is not possible to wait for the cluster-wide default network to be ready before indicating the implementations readiness to kubelet, the implementation may immediately install its own CNI configuration file to the kubelet CNI configuration directory, ensuring its configuration takes precedence over the cluster-wide default networks (if any).  The implementation must then block any pod network attachment/detachment operations until the cluster-wide default network is ready.

6.2 Cluster-wide Default Network Attachment Ordering

The implementation must attach the cluster-wide default network before attaching any networks specified by the Network Attachment Selection Annotation.

**7. Runtime and Implementation Considerations**

7.1 CNI Delegating Plugin Requirements for CNI Configuration and Result Versioning

The CNI Delegating Plugin must conform to the CNI specifications requirements for invoking plugins in a CNI configlist.  If the CNI Delegating Plugin uses the CNI projects reference "libcni" library these issues will be handled automatically.  If not, the CNI specification requires that the CNI Delegating Plugin inject the `cniVersion` and name fields from the configlist into each plugin invocations configuration JSON.  This ensures that each plugin in the configlist is able to understand the result of the previous plugin, and that the runtime receives the correctly-versioned final result.

7.2 Attachment/Detachment Failure Handling

On pod network setup, the failure to attach any network referenced by a pods Network Attachment Selection Annotation or the failure to attach the cluster-wide default network shall immediately fail the pod network setup operation.  Attachments that have not yet been performed shall not be attempted.

On pod network teardown, all network attachments that were attempted during pod network setup must be torn down, and the failure of one shall not prevent teardown of subsequent attachments.  However, if any detachment failed, a final error shall be delivered to the runtime to indicate the overall teardown operation failed.

7.3 Serialization of Network Attachment Operations

The CNI Specification 0.4.0 states "The container runtime must not invoke parallel operations for the same container, but is allowed to invoke parallel operations for different containers."  Implementations must follow this requirement and must not parallelize pod network attachment operations.  This requirement may be lifted in future versions of the specification.

7.4 Restrictions on Selection of Network Attachment Definitions by Pods

Implementations are free to impose restrictions on which Network Attachment Definitions can be selected by a given Pod.  This could be done by RBAC, admission control, or any other implementation-specific method. If the implementation determines that a given Network Attachment Definition selected by the Pod is not allowed for that Pod, it must fail the pod network operation.



