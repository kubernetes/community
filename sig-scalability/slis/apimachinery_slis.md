# API-machinery SLIs and SLOs

The document was converted from [Google Doc]. Please refer to the original for
extended commentary and discussion.

## Background

Scalability is an important aspect of the Kubernetes. However, Kubernetes is
such a large system that we need to manage users expectations in this area.
To achieve it, we are in process of redefining what does it mean that
Kubernetes supports X-node clusters - this doc describes the high-level
proposal. In this doc we are describing API-machinery related SLIs we would
like to introduce and suggest which of those should eventually have a
corresponding SLO replacing current "99% of API calls return in under 1s" one.

The SLOs we are proposing in this doc are our goal - they may not be currently
satisfied. As a result, while in the future we would like to block the release
when we are violating SLOs, we first need to understand where exactly we are
now, define and implement proper tests and potentially improve the system.
Only once this is done, we may try to introduce a policy of blocking the
release on SLO violation. But this is out of scope of this doc.


### SLIs and SLOs proposal

Below we introduce all SLIs and SLOs we would like to have in the api-machinery
area. A bunch of those are not easy to understand for users, as they are
designed for developers or performance tracking of higher level
user-understandable SLOs. The user-oriented one (which we want to publicly
announce) are additionally highlighted with bold.

### Prerequisite

Kubernetes cluster is available and serving.

### Latency<sup>[1](#footnote1)</sup> of API calls for single objects

__***SLI1: Non-streaming API calls for single objects (POST, PUT, PATCH, DELETE,
GET) latency for every (resource, verb) pair, measured as 99th percentile over
last 5 minutes***__

__***SLI2: 99th percentile for (resource, verb) pairs \[excluding virtual and
aggregated resources and Custom Resource Definitions\] combined***__

__***SLO: In default Kubernetes installation, 99th percentile of SLI2
per cluster-day<sup>[2](#footnote2)</sup> <= 1s***__

User stories:
- As a user of vanilla Kubernetes, I want some guarantee how quickly I get the
response from an API call.
- As an administrator of Kubernetes cluster, if I know characteristics of my
external dependencies of apiserver (e.g custom admission plugins, webhooks and
initializers) I want to be able to provide guarantees for API calls latency to
users of my cluster

Background:
- We obviously can’t give any guarantee in general, because cluster
administrators are allowed to register custom admission plugins, webhooks
and/or initializers, which we don’t have any control about and they obviously
impact API call latencies.
- As a result, we define the SLIs to be very generic (no matter how your
cluster is set up), but we provide SLO only for default installations (where we
have control over what apiserver is doing). This doesn’t provide a false
impression, that we provide guarantee no matter how the cluster is setup and
what is installed on top of it.
- At the same time, API calls are part of pretty much every non-trivial workflow
in Kubernetes, so this metric is a building block for less trivial SLIs and
SLOs.

Other notes:
- The SLO has to be satisfied independently from from the used encoding. This
makes the mix of client important while testing. However, we assume that all
`core` components communicate with apiserver with protocol buffers (otherwise
the SLO doesn’t have to be satisfied).
- In case of GET requests, user has an option to opt-in for accepting
potentially stale data (the request is then served from cache and not hitting
underlying storage). However, the SLO has to be satisfied even if all requests
ask for up-to-date data, which again makes careful choice of requests in tests
important while testing.


### Latency of API calls for multiple objects

__***SLI1: Non-streaming API calls for multiple objects (LIST) latency for
every (resource, verb) pair, measure as 99th percentile over last 5 minutes***__

__***SLI2: 99th percentile for (resource, verb) pairs [excluding virtual and
aggregated resources and Custom Resource Definitions] combined***__

__***SLO1: In default Kubernetes installation, 99th percentile of SLI2 per
cluster-day***__
- __***is <= 1s if total number of objects of the same type as resource in the
system <= X***__
- __***is <= 5s if total number of objects of the same type as resource in the
system <= Y***__
- __***is <= 30s if total number of objects of the same types as resource in the
system <= Z***__ 

User stories:
- As a user of vanilla Kubernetes, I want some guarantee how quickly I get the
response from an API call.
- As an administrator of Kubernetes cluster, if I know characteristics of my
external dependencies of apiserver (e.g custom admission plugins, webhooks and
initializers) I want to be able to provide guarantees for API calls latency to
users of my cluster.

Background:
- On top of arguments from latency of API calls for single objects, LIST
operations are crucial part of watch-related frameworks, which in turn are
responsible for overall system performance and responsiveness.
- The above SLO is user-oriented and may have significant buffer in threshold.
In fact, the latency of the request should be proportional to the amount of
work to do (which in our case is number of objects of a given type (potentially
in a requested namespace if specified)) plus some constant overhead. For better
tracking of performance, we define the other SLIs which are supposed to be
purely internal (developer-oriented)


_SLI3: Non-streaming API calls for multiple objects (LIST) latency minus 1s
(maxed with 0) divided by number of objects in the collection
<sup>[3](#footnote3)</sup> (which may be many more than the number of returned
objects) for every (resource, verb) pair, measured as 99th percentile over
last 5 minutes._

_SLI4: 99th percentile for (resource, verb) pairs [excluding virtual and
aggregated resources and Custom Resource Definitions] combined_

_SLO2: In default Kubernetes installation, 99th percentile of SLI4 per
cluster-day <= Xms_


### Watch latency

_SLI1: API-machinery watch latency (measured from the moment when object is
stored in database to when it’s ready to be sent to all watchers), measured
as 99th percentile over last 5 minutes_

_SLO1 (developer-oriented): 99th percentile of SLI1 per cluster-day <= Xms_

User stories:
- As an administrator, if system is slow, I would like to know if the root
cause is slow api-machinery or something farther the path (lack of network
bandwidth, slow or cpu-starved controllers, ...).

Background:
- Pretty much all control loops in Kubernetes are watch-based, so slow watch
means slow system in general. As a result, we want to give some guarantees on
how fast it is.
- Note that how we measure it, silently assumes no clock-skew in case of HA
clusters.


### Admission plugin latency

_SLI1: Admission latency for each admission plugin type, measured as 99th
percentile over last 5 minutes_

User stories:
- As an administrator, if API calls are slow, I would like to know if this is
because slow admission plugins and if so which ones are responsible.


### Webhook latency

_SLI1: Webhook call latency for each webhook type, measured as 99th percentile
over last 5 minutes_

User stories:
- As an administrator, if API calls are slow, I would like to know if this is
because slow webhooks and if so which ones are responsible.


### Initializer latency

_SLI1: Initializer latency for each initializer, measured as 99th percentile
over last 5 minutes_

User stories:
- As an administrator, if API calls are slow, I would like to know if this is
because of slow initializers and if so which ones are responsible.

---
<a name="footnote1">\[1\]</a>By latency of API call in this doc we mean time
from the moment when apiserver gets the request to last byte of response sent
to the user.

<a name="footnote2">\[2\]</a> For the purpose of visualization it will be a
sliding window. However, for the purpose of reporting the SLO, it means one
point per day (whether SLO was satisfied on a given day or not).

<a name="footnote3">\[3\]</a>A collection contains: (a) all objects of that
type for cluster-scoped resources, (b) all object of that type in a given
namespace for namespace-scoped resources.


[Google Doc]: https://docs.google.com/document/d/1Q5qxdeBPgTTIXZxdsFILg7kgqWhvOwY8uROEf0j5YBw/edit#
