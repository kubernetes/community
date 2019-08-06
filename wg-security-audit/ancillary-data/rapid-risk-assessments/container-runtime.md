# Overview

- Component: Container Runtime
- Owner(s): [sig-node](https://github.com/kubernetes/community/blob/master/sig-node/README.md)
- SIG/WG(s) at meeting:
- Service Data Classification: High
- Highest Risk Impact:

# Service Notes

The portion should walk through the component and discuss connections, their relevant controls, and generally lay out how the component serves its relevant function. For example
a component that accepts an HTTP connection may have relevant questions about channel security (TLS and Cryptography), authentication, authorization, non-repudiation/auditing,
and logging. The questions aren't the *only* drivers as to what may be spoken about, the questions are meant to drive what we discuss and keep things on task for the duration
of a meeting/call.

## How does the service work?

- Container Runtimes expose an IPC endpoint such as a file system socket
- kubelet retrieves pods to be executed from the kube-apiserver
- The Container Runtime Interface then executes the necessary commands/requests from the actual container system (e.g. docker) to run the pod

## Are there any subcomponents or shared boundaries?

Yes

- The Container Runtime technically interfaces with kublet, and runs on the same host
- However, the Container Runtime is logically a separate Trust Zone within the node

## What communications protocols does it use?

Various, depends on the IPC mechanism required by the Container Runtime 

## Where does it store data?

Most data should be provided by kubelet or the CRI in running the container

## What is the most sensitive data it stores?

N/A

## How is that data stored?

N/A

# Meeting Notes


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

- CRI Runs an HTTP server
  - port forwarding, exec, attach
- !FINDING TLS bye default, but not mutual TLS, and self-signed
  - kubelet -> exec request to CRI over gRPC
  - Returns URL with single use Token
  - gRPC is Unix Domain by default
- Kubelet proxies or responds w/ redirect to API server (locally hosted CRI only)
- !FINDING(same HTTP finding for pull as kubectl) CRI actually pulls images, no egress filtering
  - image tag is SHA256, CRI checks that
- Not sure how CNI, it might be exec
- only responds to connections 
- CRI uses Standard Go HTTP

## Cryptography

- Nothing beyond TLS

## Secrets Management

- !FINDING auth'd container repos, passed in via podspec, fetched by kubelet, are passed via CLI
  - so anyone with access to the host running the container can see those secrets

## Authentication

- Unix Domain Socket for gRPC, so Linux authN/authZ
- !FINDING 8 character random single use token with 1 minute lifetype (response to line 109)

## Authorization

- no authZ

## Multi-tenancy Isolation

- knows nothing about tenants or namespaces
- low-level component, kubelet/api-server is the arbiter

## Summary

# Recommendations
