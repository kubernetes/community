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

#### End-User Configuration of Multicast (WIP)

##### Simple Bidirectional Policies

A user can configure a namespace to allow multicast traffic by
creating an appropriate NetworkPolicy. Eg:

    kind: NetworkPolicy
    apiVersion: networking.k8s.io/v1
    metadata:
      name: mdns-1
      namespace: project-a
    spec:
      podSelector:
        matchLabels:
          app: mdns
      policyTypes:
      - Multicast
      multicast:
      - cidr: 224.0.0.251/32
        ports:
        - protocol: UDP
          port: 5353

This says "all pods in namespace `project-a` with the label `app=mdns`
should be allowed to send and receive multicast traffic to
`224.0.0.251:5353` (the reserved IP/port for multicast DNS).

##### Combined Multicast/Unicast Policies

Although mDNS normally uses multicast for both requests and responses,
clients can optionally request unicast responses, so if the namespace
is otherwise isolated for ingress/egress you might need to add rules
for unicast responses as well:

    kind: NetworkPolicy
    apiVersion: networking.k8s.io/v1
    metadata:
      name: mdns-2
      namespace: project-a
    spec:
      podSelector:
        matchLabels:
          app: mdns
      policyTypes:
      - Multicast
      - Ingress
      - Egress
      multicast:
      - cidr: 224.0.0.251/32
        ports:
        - protocol: UDP
          port: 5353
      ingress:
      - from:
        - podSelector:
            matchLabels:
              app: mdns
        ports:
        - protocol: UDP
          port: 5353
      egress:
      - to:
        - podSelector:
            matchLabels:
              app: mdns
        ports:
        - protocol: UDP
          port: 5353

("All pods in namespace `project-a` with the label `app=mdns` should
be able to send and receive multicast traffic to `224.0.0.251:5353`,
and should also be able to send and receive unicast traffic to UDP port
5353 on other pods with the label `app=mdns`.)

##### Unidirectional Multicast

Although most of the time with multicast protocols, it is most
convenient to have a single rule that says both "can send to multicast
address X" and "can receive traffic addressed to multicast address X"
as above, there may be cases where you want distinct client and server
rules.

(TBD: Really? Does anyone *actually* want this?)

In that case, rather than having a single "multicast" policy type, it
might be better to specify multicast permissions inside the "ingress"
and "egress" sections:

(NB: different proposed syntax from the earlier examples)

    kind: NetworkPolicy
    apiVersion: networking.k8s.io/v1
    metadata:
      name: mdns-3-client
      namespace: project-a
    spec:
      podSelector:
        matchLabels:
          app: mdns-client
      policyTypes:
      - Ingress
      - Egress
      egress:
      - multicast:
        - cidr: 224.0.0.251/32
          ports:
          - protocol: UDP
            port: 5353
      ingress:
      - from:
        - podSelector:
            matchLabels:
              app: mdns-server
        ports:
        - protocol: UDP
          port: 5353

    kind: NetworkPolicy
    apiVersion: networking.k8s.io/v1
    metadata:
      name: mdns-3-server
      namespace: project-a
    spec:
      podSelector:
        matchLabels:
          app: mdns-server
      policyTypes:
      - Ingress
      - Egress
      ingress:
      - multicast:
        - cidr: 224.0.0.251/32
          ports:
          - protocol: UDP
            port: 5353
      egress:
      - to:
        - podSelector:
            matchLabels:
              app: mdns-client
        ports:
        - protocol: UDP
          port: 5353

This policy splits up client and server permissions, and requires
requests to be multicast and responses to be unicast. (This probably
doesn't actually make sense in the mDNS case, but let's pretend it
does.) So the first policy says that `app=mdns-client` pods can send
multicast traffic to `224.0.0.251:5353`, and can receive unicast
packets from `app=mdns-server` pods on port 5353. The second policy
says that `app=mdns-server` pods can receive multicast traffic to
`224.0.0.251:5353` and send unicast packets to `app=mdns-client` pods
on port 5353. (We need to explicitly specify both client and server
policy here because with mixed multicast/unicast protocols the source
and destination IPs on the request and reply packets don't match up in
a way that conntrack can recognize, so there's no way to make the
replies be accepted automatically.)

Note that we use new `ingress.multicast` and `egress.multicast` fields
rather than just reusing `ingress.from.ipBlock` / `egress.to.ipBlock`,
because with `ipBlock` the semantics would be backwards for `ingress`:

    ingress:
    - from:
      - ipBlock:
          cidr: 224.0.0.251/32
      ports:
      - protocol: UDP
        port: 5353

Logically, this says "accept incoming packets *coming from*
224.0.0.251", but we want the rule to mean "accept incoming packets
*going to* 224.0.0.251". So a new subfield makes more sense.

##### Scope of Multicast Traffic

There is only a single set of multicast IP addresses, and certain
protocols reserve specific ones (eg, 224.0.0.251 for mDNS). If we want
to allow for the possibility of non-intercommunicating sets of pods
using the same multicast IPs (eg, an mDNS server in project-a only
responding to requests on 224.0.0.251 from project-a pods, and an mDNS
server in project-b only responding to requests on 224.0.0.251 from
project-b pods) then we need to define how multicast traffic is
scoped.

The three simplest options are:

  1. Scope-to-Cluster: There is no scoping; all multicast traffic is
     cluster-wide. (Weave's model)

  2. Scope-to-Namespace: Multicast traffic is scoped to the namespace
     it originates from. You can have distinct mDNS multicast groups
     in different namespaces, but you can't have distinct mDNS
     multicast groups for `environment=testing` and
     `environment=production` pods in the same namespace. (OpenShift
     SDN's model)

  3. Scope-to-NetworkPolicy: Traffic allowed by a given NetworkPolicy
     is scoped to the pods selected by that NetworkPolicy. If you
     create two copies of a multicast NetworkPolicy with different
     `spec.podSelector` values, then this would form two separate
     scopes (so you could have separate multicast groups for
     `environment=testing` and `environment=production`).

     Since NetworkPolicies are tied to particular namespaces, this
     means that it would not be possible to implement cross-namespace
     multicast. It would also not be possible to implement
     asymmetrical policies like `mdns-3-client`/`mdns-3-server` above,
     since the two NetworkPolicy resources would end up creating
     separate non-communicating scopes.

A more complicated and powerful solution would be to have Explicit
Scopes specified in the NetworkPolicy:

    kind: NetworkPolicy
    apiVersion: networking.k8s.io/v1
    metadata:
      name: mdns-4-client
      namespace: project-a
    spec:
      podSelector:
        matchLabels:
          app: mdns-client
      policyTypes:
      - Ingress
      - Egress
      egress:
      - multicast:
        - cidr: 224.0.0.251/32
          ports:
          - protocol: UDP
            port: 5353
          scope:
            podSelector:
              matchExpressions:
              - { key: app, operator: In, values: [mdns-client, mdns-server] }
      ...

(This example is based on the syntax from the "Unidirectional
Multicast" example, but the new field would work with the syntax from
"Simple Bidirectional Policies" too.)

The `scope` field here specifies that the multicast traffic is grouped
to the given `NetworkPolicyPeer`; in this case, a podSelector
selecting both `app=mdns-client` and `app=mdns-server` pods. The
server-side policy would have the same `scope`.

(TBD: What if a pod matches the labels of two different multicast
policies? Would this create partially-overlapping scopes, where pod A
sends to and receives from both pod B and pod C, but pod B and pod C
do not see multicast traffic from each other? Is that a problem if so?
Note that this applies to Scope-to-NetworkPolicy as well, not just
Explicit Scopes.)

Note that we could use one of the simple implicitly-scoped versions
for the initial version of the feature, and then add Explicit Scopes
on top of that later, by retconning the earlier implicit scope as just
being the default value for `scope` when no `scope` was specified.
(For Scope-to-Cluster, the default `scope` would be `scope: {
namespaceSelector: {} }`. For Scope-to-Namespace, it would be `scope:
{ podSelector: {} }`, and for Scope-to-NetworkPolicy, the default
would be a copy of the NetworkPolicy's `spec.podSelector`.)

##### Cross-Namespace Scopes

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

Assuming something like the Explicit Scopes model above, we could say
that two pods in different namespaces are in the same scope if each
one is selected by a NetworkPolicy in its namespace whose scope
selects both pods. Eg, if project-a and project-b both have the label
`user=alice`, and both contain a policy:

    kind: NetworkPolicy
    apiVersion: networking.k8s.io/v1
    metadata:
      name: mdns-5
    spec:
      podSelector:
        matchLabels:
          app: mdns
      policyTypes:
      - Multicast
      multicast:
      - cidr: 224.0.0.251/32
        ports:
        - protocol: UDP
          port: 5353
        scope:
        - namespaceSelector:
            matchLabels:
              user: alice
          podSelector:
            matchLabels:
              app: mdns

then traffic to 224.0.0.251:5353 will be shared between `app=mdns`
pods in the two namespaces.

The `scope` simultaneously indicates where the multicast packets
should be delivered, *and* who else is allowed to join the scope.
Other namespaces labeled `user=alice` can join the group by creating a
matching NetworkPolicy, but namespaces that aren't labeled
`user=alice` would be unable to join, regardless of what `scope` they
specified, because they wouldn't be matched by the `scope`s of the
existing members.

(TBD: This creates even more potential for confusing overlapping
scopes than in the single-namespace case above... Also, we need to
really think through the details of this to make sure the security
model actually works. Also, that this doesn't end up being
impossibly-complicated to implement.)

(As already mentioned when discussing Explicit Scopes above, if we
initially started with Scope-to-Namespace or Scope-to-NetworkPolicy,
we could then add Explicit Scopes, and thus cross-namespace multicast,
at a later date. Presumably other ways of defining cross-namespace
multicast could also be bolted on later.)

In this model, Weave's cluster-wide multicast would be equivalent to
giving every namespace a policy like:

    kind: NetworkPolicy
    apiVersion: networking.k8s.io/v1
    metadata:
      name: allow-all-multicast
    spec:
      podSelector:
      policyTypes:
      - Multicast
      multicast:
      - cidr: 224.0.0.0/4
        scope:
        - namespaceSelector: {}
      - cidr: ff80::/8
        scope:
        - namespaceSelector: {}

Though any namespace that wanted to opt out of the fun could do so by
deleting that policy.

(TBD: Do we really need cross-namespace multicast anyway? It would
simplify things a lot to say "no", but in the "namespaceSelector plus
podSelector" discussion, people had various reasons for needing pods
to be in separate namespaces even though they were part of the same
"application"...)

##### Multicast Across the Pod Network Boundary

In cases where the pod network is directly reachable from some
cluster-external hosts, people might want to do multicast with those
hosts. This could perhaps be expressed with an `ipBlock`-based
`scope`?

This runs into the "consent" problem again though; there's no way for
the external hosts to indicate whether they had or had not intended
for the pods to be able to join their multicast group. In particular,
if the node itself (or a hostNetwork pod) is using some
multicast-based protocol on the underlying network, it may not want to
share that traffic with anyone on the pod network.

##### Non-NetworkPolicy-based Solutions

It's also possible that NetworkPolicy isn't the right way to configure
multicast. I haven't thought much about this, other than that it might
simplify the cross-namespace and cluster-external cases if only
cluster administrators were able to enable multicast in those cases.

#### Interaction with non-multicast NetworkPolicies

If a user has NetworkPolicies with `ipBlock`s that extend over the
multicast IP range, how should those interact with multicast traffic?
I feel like the right answer here is "they don't; `ipBlock` only
matches unicast traffic". In particular, a policy allowing traffic
to/from `0.0.0.0/0` or `::/0` should not cause multicast traffic to be
enabled.

### Risks and Mitigations

TBD

## Graduation Criteria

- The documentation indicates the expected default behavior of Kubernetes clusters with respect to IP multicast traffic.
- The networking tests validate the expected default behavior.
- TBD

## Implementation History

TBD
