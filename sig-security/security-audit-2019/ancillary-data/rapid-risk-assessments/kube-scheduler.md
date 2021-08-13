# Overview

- Component: kube-scheduler
- Owner(s): [sig-scheduling](https://github.com/kubernetes/community/tree/master/sig-scheduling)
- SIG/WG(s) at meeting:
- Service Data Classifjcation: Moderate (the scheduler adds pods to nodes, but will not remove pods, for the most part)
- Highest Risk Impact:

# Service Notes

The portion should walk through the component and discuss connections, their relevant controls, and generally lay out how the component serves its relevant function. For example
a component that accepts an HTTP connection may have relevant questions about channel security (TLS and Cryptography), authentication, authorization, non-repudiation/auditing,
and logging. The questions aren't the *only* drivers as to what may be spoken about, the questions are meant to drive what we discuss and keep things on task for the duration
of a meeting/call.

## How does the service work?

- Similar to most other components:
  1. Watches for unscheduled/new pods
  1. Watches nodes with and their resource constraints
  1. Chooses a node, via various mechanisms, to allocate based on best fit of resource requirements
  1. Updates the pod spec on the kube-apiserver
  1. that update is then retrieved by the node, which is also Watching components via the kube-apiserver
- there may be multiple schedulers with various names, and parameters (such as pod-specific schedulers)

- !NOTE schedulers are coöperative
- !NOTE schedulers are *supposed* to honor the name, but need not
  - Interesting note, makes the huge list of schedulers DoS interesting
  - !NOTE idea there was to add a *huge* number of pods to be scheduled that are associated with an poorly named scheduler
  - !NOTE peopoe shouldn't request specific schedulers in podspec, rather, there should be some webhook to process that
  - !NOTE team wasn't sure what would happen with large number of pods to be scheduled

## Are there any subcomponents or shared boundaries?

Yes

- there may be multiple schedulers on the same MCP host
- schedulers may run on the same host as the API server

## What communications protocols does it use?

- standard HTTPS + auth (chosen by the cluster)

## Where does it store data?

- most should be stored in etcd (via kube-apiserver)
- some data will be stored on command line (configuration options) or on the file system (certificate paths for authentication)

## What is the most sensitive data it stores?

- No direct storage

## How is that data stored?

- N/A

# Data Dictionary

| Name | Classification/Sensitivity | Comments |
| :--: | :--: | :--: |
| Data | Goes | Here |

# Control Families 

These are the areas of controls that we're interested in based on what the audit working group selected. 

When we say "controls," we mean a logical section of an application or system that handles a security requirement. Per CNSSI:

> The management, operational, and technical controls (i.e., safeguards or countermeasures) prescribed for an information system to protect the confidentiality, integrity, and availability of the system and its information.

For example, an system may have authorization requirements that say:

- users must be registered with a central authority
- all requests must be verified to be owned by the requesting user
- each account must have attributes associated with it to uniquely identify the user

and so on. 

For this assessment, we're looking at six basic control families:

- Networking
- Cryptography
- Secrets Management
- Authentication
- Authorization (Access Control)
- Multi-tenancy Isolation

Obviously we can skip control families as "not applicable" in the event that the component does not require it. For example,
something with the sole purpose of interacting with the local file system may have no meaningful Networking component; this
isn't a weakness, it's simply "not applicable."

For each control family we want to ask:

- What does the component do for this control?
- What sorts of data passes through that control? 
  - for example, a component may have sensitive data (Secrets Management), but that data never leaves the component's storage via Networking
- What can attacker do with access to this component?
- What's the simplest attack against it?
- Are there mitigations that we recommend (i.e. "Always use an interstitial firewall")?
- What happens if the component stops working (via DoS or other means)?
- Have there been similar vulnerabilities in the past? What were the mitigations?

# Threat Scenarios

- An External Attacker without access to the client application
- An External Attacker with valid access to the client application
- An Internal Attacker with access to cluster
- A Malicious Internal User

## Networking

- only talks to kube-apiserver
- colocated on the same host generally as kube-apiserver, but needn't be
- has a web server (HTTP)
  - !FINDING: same HTTP server finding as all other components
  - metrics endpoint: qps, scheduling latency, &c
  - healthz endpoint, which is just a 200 Ok response
  - by default doesn't verify cert (maybe)

## Cryptography

- None

## Secrets Management

- Logs is the only persistence mechanism
- !FINDING (to be added to all the other "you expose secrets in env and CLI" finding locations) auth token/cred passed in via CLI

## Authentication

- no authN really
- pods, nodes, related objects; doesn't deal in authN
- unaware of any service/user accounts

## Authorization

- schedluinc concepts protected by authZ
  - quotas
  - priority classes
  - &c
- this authZ is not enforced by scheduler, however, enforced by kube-apiserver

## Multi-tenancy Isolation

- tenant: different users of workloads that don't want to trust one another 
- namespaces are usually the boundaries
- affinity/anti-affinity for namespace 
- scheduler doesn't have data plan access
- can have noisy neighbory problem
  - is that the scheduler's issue?
  - not sure
  - namspace agnostic
  - can use priority classes which can be RBAC'd to a specific namespace, like kube-system
  - does not handle tenant fairness, handles priorty class fairness
  - no visibility into network boundary or usage information
  - no cgroup for network counts
  - !FINDING anti-affinity can be abused: only I can have this one host, no one else, applicable from `kubectl` 
  - !NOTE no backoff process for scheduler to reschedule a rejected pod by the kublet; the replicaset controller can create a tightloop (RSC -> Scheduler -> Kubelet -> Reject -> RSC...)

## Summary

# Recommendations
