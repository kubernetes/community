# Overview

- Component: kubelet
- Owner(s): [sig-node](https://github.com/kubernetes/community/tree/master/sig-node)
- SIG/WG(s) at meeting:
- Service Data Classification: High
- Highest Risk Impact:

# Service Notes

The portion should walk through the component and discuss connections, their relevant controls, and generally lay out how the component serves its relevant function. For example
a component that accepts an HTTP connection may have relevant questions about channel security (TLS and Cryptography), authentication, authorization, non-repudiation/auditing,
and logging. The questions aren't the *only* drivers as to what may be spoken about, the questions are meant to drive what we discuss and keep things on task for the duration
of a meeting/call.

## How does the service work?

- `kubelet` isses a watch request on the `kube-apiserver`
- `kubelet` watches for pod allocations assigned to the node the kubelet is currently running on
- when a new pod has been allocated for the kubelet's host, it retrieve the pod spec, and interacts with the Container Runtime via local Interprocess Communication to run the container
- Kubelet also handles:
  - answering log requests from the kube-apiserver
  - monitoring pod health for failures
  - working with the Container Runtime to deschedule pods when the pod has been deleted
  - updating the kube-apiserver with host status (for use by the scheduler)

## Are there any subcomponents or shared boundaries?

Yes.

- Technically, kubelet runs on the same host as the Container Runtime and kubeproxy
- There is a Trust Zone boundary between the Container Runtime and the kubelet

## What communications protocols does it use?

- HTTPS with certificate validation and some authentication mechanism for communication with the kube-apiserver as a client
- HTTPS without certificate validation by default 

## Where does it store data?

- kubelet itself should not store much data
- kubelet can be run in an "apiserver-less mode" that loads pod manifests from the file system
- most data should be retrieved from the kube-apiserver via etcd
- authentication credentials for the kube-apiserver may be stored on the file system or in memory (both in CLI parameter as well as actual program memory) for the duration of execution

## What is the most sensitive data it stores?

- authentication credentials are stored in memory or are out of scope

## How is that data stored?

N/A

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

- Post 10250: read/write, authenticated
- Port 10255: read-only, unauthenticated
  - cadvisor uses this, going to be deprecated
- 10248: healthz, unauth'd
- static pod manifest directory
- Static pod fetch via HTTP(S)

### Routes: 

- Auth filter on API, for 10250
  - delegated to apiserver, subject access review, HTTPS request
- `/pods` podspec on node -> leaks data
- `/healthz`
- `/spec`
- `/stats-{cpu, mem, &c}`
- on 10250 only:
  - `/exec`
  - `/attach`
  - `portforward`
  - `/kube-auth`
  - `/debug-flags`
  - `/cri/{exec, attach, portforward}`

### Findings:

- !FINDING: 10255 is unauthenticated and leaks secrets
- !FINDING: 10255/10248 
- !FINDING: 10250 is self-signed TLS 

## Cryptography

- None

## Secrets Management

- returned from kube-apiserver unencrypted
- in memory cache
- if pod mounts disk, written to tmpfs
- !FINDING (already captured) ENV vars can expose secrets
- configmaps are treated like secrets by kubelet
- !FINDING keynames and secret names may be logged
- maintains its own certs, secrets, bootstrap credential
  - bootstrap: initial cert used to issue CSR to kube-apiserver
  - !NOTE certs are written to disk unencrypted
  - !FINDING bootstrap cert may be long lived, w/o a TTL

## Authentication

- delegated to kube-apiserver, via HTTPS request, with subject access review
- two-way TLS by default (we believe)
- token auth
  - bearer token
  - passed to request to API server
  - "token review"
  - kube-apiserver responds w/ ident
  - response is boolean (yes/no is this a user) and username/uid/groups/arbitrary data as a tuple
- no auditing on kublet, but logged on kube-apiserver

## Authorization

- delegated to kube-apiserver

## Multi-tenancy Isolation

- kube-apiserver is the arbiter
- kubelet doesn't know namespaces really
- every pod is a separate tenant
- pods are security boundaries

## Summary

# Recommendations
