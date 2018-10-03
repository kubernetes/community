Enabling Plugins for Kube-Proxy to support different backends
=============================================================

Table of Contents<br>
---------------------

-   [Table of Contents](#table-of-contents)

-   [Summary](#summary)

-   [Motivation](#motivation)

    -   [Goals](#goals)

    -   [Non-Goals](#non-goals)

-   [Proposal](#proposal)

    -   [User Stories [optional]](#user-stories-optional)

    -   [Story 1](#risks-and-mitigations)

    -   [Story 2](#story-2)

    -   [Implementation Details/Notes/Constraints
        [optional]](#implementation-detailsnotesconstraints-optional)

    -   [Risks and Mitigations](#risks-and-mitigations)

-   [Graduation Criteria](#graduation-criteria)

-   [Implementation History](#implementation-history)

-   [Drawbacks [optional]](#drawbacks-optional)

-   [Alternatives [optional]](#alternatives-optional)

Summary
-------

This proposal aims to make the kube-proxy more pluggable to support different
proxying mechanisms to be plugged in. This would encourage out of tree
development and enhancements for the kube-proxy component.

Motivation
----------

Kube-proxy component is one of the critical component in the data path of the
service endpoints and abstracts the service endpoint from the consumer. Over a
period this component has evolved from being a user-space based component to
iptables and then to LVS based mechanism. This still remains an opinionated
component and doesn’t provide flexibility for out of tree development of
kube-proxy.

Since kube-proxy is majorly providing some network functions, the proposal is to
make this component as a plugin on lines of other K8s components which allows
optimal proxying mechanisms to be developed out of tree. Examples could be an
ebpf based mechanism to be used to proxy the calls to the service endpoint.
Another example would be to optimize node ports via socket activation or xdp
like mechanisms to be used to facilitate fast load balancing to service
endpoints and also allow doing DdOS mitigations within the kernel.

### Goals

Allow kube-proxy to support plugins for different out of tree developed backends

### Non-Goals

To propose any specific backend

Proposal
--------

Define a kubeProxy Plugin interface which can be registered with the core.

The kube-Proxy config today defines the mode of proxying. Replace this config
with the kube-Proxy-plugin config which points to the plugin path and add
another config which points to the config of the plugin.

The plugin adheres to the spec defined by kube-proxy and invoked over domain
sockets/netlinks and internally the plugin implementation is called to apply the
rules which can be as an example configuring and loading an ebpf program to
forward the traffic to the service backend, or configuring the node port for
preventing a DDoS attack on the nodeport via some ip blacklisting .

There can be other use cases to allow configuring traffic shaping or hooking
taps in the data path via the plugin.

### User Stories [optional]

The enhancement needs changes to the kube Proxy component to load the configured
plugin as part of bootstrapping the plugin.

The kube-proxy watches the service and endpoints for changes and invokes the
configured plugin over a domain socket/netlink based mechanism to apply the
rules. This will be similar implementation as to how kubelet can be configured
to talk to Container Storage Interfaces.

### Risks and Mitigations

A wrongly configured proxy can slow down the packet processing. But this again
is giving more control to the customer to define the traffic shaping the way
they want and move away from a more opinionated model. Risks can be mitigated by
providing recommendations on how to design the plugin and also providing some
example plugins to be cited as an example of usage.

Graduation Criteria
-------------------

Implementation History
----------------------

Drawbacks [optional]
--------------------

Alternatives [optional]
-----------------------

Infrastructure Needed [optional]
--------------------------------
