---
kep-number: 31
title: Multicast Support
authors:
  - "@danwinship"
owning-sig: sig-network
reviewers:
  - TBD
approvers:
  - "@thockin"
editor: TBD
creation-date: 2018-11-09
last-updated: 2018-11-09
status: provisional
---

# Multicast Support

## Table of Contents

* [Table of Contents](#table-of-contents)
* [Summary](#summary)
* [Motivation](#motivation)
    * [Goals](#goals)
    * [Non-Goals](#non-goals)
* [Existing Work](#existing-work)
* [Proposal](#proposal)
    * [User Stories](#user-stories)
    * [Implementation Details/Notes/Constraints](#implementation-detailsnotesconstraints)
    * [Risks and Mitigations](#risks-and-mitigations)
* [Graduation Criteria](#graduation-criteria)
* [Implementation History](#implementation-history)

## Summary

Currently, support for IP multicast traffic on the pod network is
entirely undefined, with no multicast-related behavior being either
required or forbidden, and different plugins offering different
functionality.

This KEP aims first to clarify the default behavior (that the pod
network should not carry multicast traffic unless configured to so),
and second, to define a way for a user to indicate that an application
does require multicast support, when running under a network plugin
that supports multicast.

## Motivation

Some Kubernetes users have workloads that require multicast traffic,
and need to be able to confirm that multicast traffic will flow
between the desired endpoints, but not be visible to other endpoints.

Other Kubernetes users are more concerned with simply ensuring that
undesired multicast traffic is blocked, to prevent a single pod from
being able to easily saturate the network.

See also:
[Multicast in Weave](https://www.weave.works/blog/multicasting-cloud-need-know/),
[Multicast in OpenShift](https://blog.openshift.com/service-discovery-openshift-using-multicast/).

### Goals

- Explicitly define the expected baseline behavior of multicast traffic in a Kubernetes cluster.
- Define additional levels of network plugin multicast support beyond the baseline.
- Define how users enable/disable/configure multicast support and how it interacts with NetworkPolicy.

### Non-Goals

- Requiring all plugins to support multicast.
- Forcing all plugins to support only a least-common-denominator form of multicast.
- Specifying anything having to do with multicast on interfaces/networks other than the default cluster network.

## Existing Work

Most plugins currently block all multicast (either intentionally, or
just accidentally as a consequence of IP/MAC filtering).

Weave treats multicast as broadcast. There has been discussion of
[implementing IGMP
snooping](https://github.com/weaveworks/weave/issues/178) so that
multicast packets would only be sent to the pods that want them, but
this is not yet implemented. There has been some discussion of
[multicast between pods and external
IPs](https://github.com/weaveworks/weave/issues/1863), which
apparently can be made to work at least in some circumstances. There
is no support for policy/isolation; all multicast packets are sent to
all pods in the cluster.

OpenShift SDN also implements multicast-as-broadcast, but on a
per-namespace basis, [with namespaces having to opt
in](https://docs.okd.io/latest/admin_guide/managing_networking.html#admin-guide-networking-multicast).
(So if you enable multicast on namespace "foo", and a pod in "foo"
sends out a multicast packet, it will be delivered to all other pods
in "foo".) As with Weave, it's not possible to use NetworkPolicy to
restrict multicast traffic, and it's also not possible to use
NetworkPolicy to extend multicast traffic across namespace boundaries.

## Proposal

### User Stories

#### Running WildFly ("JBoss") under Kubernetes

WildFly (formerly and sometimes still currently known as JBoss
Application Server) allows running [High
Availability](http://docs.wildfly.org/14/High_Availability_Guide.html)
Java EE applications. It uses a library called
[JGroups](http://www.jgroups.org/) to allow new and old servers to
discover each other as instances are added and removed, and to manage
communication between servers in the cluster.

Although there are several ways to configure JGroups, the default (and
generally preferred) configuration uses multicast for discovery; when
a new server instance is brought up, it joins an IP multicast group
(using either a default multicast IP or one specified in a
configuration file), and then sends a multicast message to that
address announcing its presence. The other existing servers will
respond, allowing each server to learn about each of the others
without having known about any of them in advance.

(As an alternative to using multicast for discovery, there is a
JGroups extension called
[KUBE_PING](https://github.com/jgroups-extras/jgroups-kubernetes) that
allows peers to find each other by making Kubernetes apiserver calls.
Although this provides a workaround for clusters where multicast is
not available, it's generally less preferred since it requires
additional configuration, and in particular, in clusters using RBAC it
may require updating role bindings to grant the WildFly pods the
apiserver access they need.)

The default JGroups configuration also uses multicast for most
peer-to-peer communication between the servers, although it is
possible to configure it to use unicast TCP instead.

#### User Story 2

TBD

### Implementation Details/Notes/Constraints

#### Baseline Behavior: Block IP Multicast

Although the Kubernetes network is normally open-by-default, except
where NetworkPolicy dictates otherwise, it seems like
closed-by-default is a better default for multicast:

1. In many clusters, open-by-default multicast would be useful only as
a denial of service attack vector.

2. Even in cases where some pods in different namespaces on different
nodes need to communicate via multicast, it is still likely to be the
case that *most* pods will not be interested in multicast traffic, and
sending the traffic to them would be a waste of bandwidth.

3. This is probably the most-commonly-implemented behavior among
existing network plugins anyway.

4. Some network plugins simply expose an underlying provider network,
and in some cases (eg, GCP), that underlying network does not support
multicast.

#### Optional IP Multicast Support

Plugins that want to provide support for IP multicast can allow users
to configure it as described below. In that case, there are two levels
of multicast support that the plugin might implement:

1. For basic multicast support, a plugin can simply treat multicast as
broadcast; when a pod sends a multicast packet, it will be transmitted
to all other allowable recipients, regardless of whether those
recipients have subscribed to any multicast groups. This is sufficient
for at least basic use cases.

2. For more advanced multicast support, a plugin could implement IGMP
snooping, and monitor which multicast addresses each pod is interested
in traffic for, and then only deliver multicast traffic to the pods
which have subscribed to that traffic. This might be more efficient if
there are many pods communicating over different multicast addresses.

#### End-User Configuration of Multicast

TBD

##### Configuring Cross-Namespace Multicast

Allowing cross-namespace multicast is tricky if we want to use
NetworkPolicy to configure it.

In the unicast case, we let namespace A create a policy saying "accept
traffic from namespace B", but that policy only has any real effect if
namespace B actually tries to send traffic to one of namespace A's
pods. So in that case, we know that both parties approve of the
communication; namespace A indicated this by creating a NetworkPolicy,
and namespace B indicated it by sending the traffic.

In the multicast case, the traffic destination is going to be a
multicast IP, not a pod IP, so sending to it doesn't confirm (or deny)
any intent to send cross-namespace traffic. So we can't just let
namespace A create a policy saying "accept multicast traffic that
namespace B sends to 224.0.0.251:5353", because namespace B might not
want its traffic to that IP to be going to namespace A as well.

So if we are using NetworkPolicy to configure cross-namespace
multicast, then both namespaces will need to create policies allowing
it. This potentially gets complicated, if namespaces have asymmetric
policies (eg, A allows multicast only with B, but B allows multicast
with both A and C).

(The other possibility would be to only allow cross-namespace
multicast to be configured by cluster administrators, not namespace
administrators.)

##### Client-vs-server / multicast-vs-unicast communication

In multicast protocols, the request will have a multicast destination
IP, but the reply will have a unicast source IP, and so will not be
recognized as a reply by conntrack (right?). Thus, even for strictly
client-server protocols, policy rules are likely to be needed in both
directions.

In some cases (eg, mDNS), the server is expected to respond via
multicast, and so the clients and servers can use the same policies.
In cases where they need distinct policies, it would be helpful if
things were defined in such a way that you could specify both policies
in a single NetworkPolicy resource.

### Risks and Mitigations

TBD

## Graduation Criteria

- The documentation indicates the expected default behavior of Kubernetes clusters with respect to IP multicast traffic.
- The networking tests validate the expected default behavior.
- TBD

## Implementation History

TBD
