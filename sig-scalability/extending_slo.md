# Extended Kubernetes scalability SLOs

## Goal
The goal of this effort is to extend SLOs which Kubernetes cluster has to meet to support given number of Nodes. As of April 2017 we have only two SLOs:
- API-responsiveness: 99% of all API calls return in less than 1s
- Pod startup time: 99% of Pods (with pre-pulled images) start within 5s
which are enough to guarantee that cluster doesn't feel completely dead, but not enough to guarantee that it satisfies user's needs.

We're going to define more SLOs based on most important indicators, and standardize the format in which we speak about our objectives. Our SLOs need to have two properties:
- They need to be testable, i.e. we need to have a benchmark to measure if it's met,
- They need to be expressed in a way that's possible to understand by a user not intimately familiar with the system internals, i.e. formulation can't depend on some arcane knowledge.
On the other hand we do not require that:
- SLOs are possible to monitor in a running cluster, i.e. not all SLOs need to be easily translatable to SLAs. Being able to benchmark is enough for us.

## Split metrics from environment
Currently what me measure and how we measure it is tightly coupled. This means that we don't have good environmental constraint suggestions for users (e.g. how many Pods per Namespace we support, how many Endpoints per Service, how to setup the cluster etc.). We need to decide on what's reasonable and make the environment explicit.

## Split SLOs by kind
Current SLOs implicitly assume that the cluster is in a "steady state". By this we mean that we assume that there's only some, limited, number of things going during benchmarking. We need to make this assumption explicit and split SLOs into two categories: steady-state SLOs and burst SLOs.

## Steady state SLOs
With steady state SLO we want to give users the data about system's behavior during normal operation. We define steady state by limiting the churn on the cluster. 

This includes current SLOs:
- API call latency
- E2e Pod startup latency

By churn we understand a measure of amount changes happening in the cluster. It's formal(-ish) definition will follow, but informally it can be thought about as number of user-issued requests per second plus number of pods affected by those requests. 

More formally churn per second is defined as:
#Pod creations + #PodSpec updates + #user originated requests in a given second
The last part is necessary only to get rid of situations when user is spamming API server with various requests. In ordinary circumstances we expect it to be in the order of 1-2. 

## Burst SLOs
With burst SLOs we want to give user idea on how system behaves under the heavy load, i.e. when one want the system to do something as quickly as possible, not caring too much about response time for a single request. Note that this voids all steady-state SLOs.

This includes the new SLO:
- Pod startup throughput

## Environment
Kubernetes cluster in which we benchmark SLOs need to meet following criteria:
- Run a single master machine appropriately sized
- Main etcd runs as a single instance on the master machine
- Events are stored in a separate etcd instance running on the master machine
- Kubernetes version is at least 1.X.Y
- Components configuration = _?_

_TODO: NEED AN HA CONFIGURATION AS WELL_

## SLO template
All our performance SLOs should be defined using the following template:

---

# SLO: *TL;DR description of the SLO*
## (Burst|Steady state) foo bar SLO

### Summary
_One-two sentences describing the SLO, that's possible to understand by the majority of the community_

### User Stories
_Few user stories showing in what situations users might be interested in this SLO, and why other ones are not enough_

## Full definition
### Test description
_Precise description of test scenario, including maximum number of Pods per Controller, objects per namespace, and anything else that even remotely seems important_

### Formal definition (can be skipped if the same as title/summary)
_Precise and as formal as possible definition of SLO. This does not necessarily need to be easily understandable by layman_
