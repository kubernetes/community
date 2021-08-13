# Overview

- Component: etcd
- Owner(s): Technically external to Kubernetes itself, but managed by [sig-api-machinery](https://github.com/kubernetes/community/tree/master/sig-api-machinery)
- SIG/WG(s) at meeting:
- Service Data Classification: Critical (on a cluster with an API server, access to etcd is root access to the cluster)
- Highest Risk Impact:

# Service Notes

The portion should walk through the component and discuss connections, their relevant controls, and generally lay out how the component serves its relevant function. For example
a component that accepts an HTTP connection may have relevant questions about channel security (TLS and Cryptography), authentication, authorization, non-repudiation/auditing,
and logging. The questions aren't the *only* drivers as to what may be spoken about, the questions are meant to drive what we discuss and keep things on task for the duration
of a meeting/call.

## How does the service work?

- Distributed key-value store
- uses RAFT for consensus
  - always need to deploy (N x M) + 1 members to avoid leader election issues
  - five is recommended for production usage
- listens for requests from clients
- clients are simple REST clients that interact via JSON or other mechanisms
- in Kubernetes' case, data is stored under `/registry`

## Are there any subcomponents or shared boundaries?

There shouldn't be; documentation specifically states:

- should be in own cluster
- limited to access by the API server(s) only
- should use some sort of authentication (hopefully certificate auth)

## What communications protocols does it use?

- HTTPS (with optional client-side or two-way TLS)
  - can also use basic auth
- there's technically gRPC as well

## Where does it store data?

- typical database-style:
  - data directory
  - snapshot directory
  - write-ahead log (WAL) directory
- all three may be the same, depends on command line options
- Consensus is then achieved across nodes via RAFT (leader election + log replication via distributed state machine)

## What is the most sensitive data it stores?

- literally holds the keys to the kingdom:
  - pod specs
  - secrets
  - roles/attributes for {R, A}BAC
  - literally any data stored in Kubernetes via the kube-apiserver
- [Access to etcd is equivalent to root permission in the cluster](https://kubernetes.io/docs/tasks/administer-cluster/configure-upgrade-etcd/#securing-etcd-clusters) 

## How is that data stored?

- Outside the scope of this assessment per se, but not encrypted at rest
- Kubernetes supports this itself with Encryption providers
- the typical process of a WAL + data + snapshot is used
- this is then replicated across the cluster with Raft

# Meeting Notes

- No authorization (from k8s perspective)
- AUthentication by local port access in current k8s
  - working towards mTLS for all connections
- Raft consensus port, listener port
- backups in etcd (system-level) not encrypted 
- metrics aren't encrypted at all either 
- multi-tenant: no multi-tenant controls at all
  - the kube-apiserver is the arbiter namespaces
  - could add namespaces to the registry, but that is a large amount of work
  - no migration plan or test
  - watches (like kubelet watching for pod spec changes) would break
  - multi-single tenant is best route
- RAFT port may be open by default, even in single etcd configuraitons
- runs in a container within static Master kubelet, but is run as root
- [CONTROL WEAKNESS] CA is passed on command line
- Types of files: WAL, Snapshot, Data file (and maybe backup)
  - [FINDING] no checksums on WAL/Snapshot/Data
  - [RECOMMENDATION] checksum individual WAL entries, checksum the entire snapshot file
  - do this because it's fast enough for individual entries, and then the snapshot should never change
- Crypto, really only TLS (std go) and checksums for backups (but not other files, as noted above)
- No auditing, but that's less useful
  - kube-apiserver is the arbiter of what things are
  - kube-apiserver uses a single connection credential to etcd w/o impersonation, so harder to tell who did what
  - major events end up in the app log
  - debug mode allows you to see all events when they happen

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

## Cryptography

## Secrets Management

## Authentication

- by default Kubernetes doesn't use two-way TLS to the etcd cluster, which would be the most secure (combined with IP restrictions so that stolen creds can't be reused on new infrastructure)

## Authorization

## Multi-tenancy Isolation

## Summary

# Recommendations
