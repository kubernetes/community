---
title: "Bug Bounty Programs"
weight: 1
slug: "bounties"
---

This program is a **work in progress**. This tracks the currently proposed program, but the bug
bounty is not currently active. If you currently have a bug to submit, follow instructions at https://kubernetes.io/docs/reference/issues-security/security/

## Scope

The scope of the Kubernetes project itself is not well defined, which makes defining a bug bounty
scope difficult. With that in mind, we've attempted to enumerate the features, configurations, and
components that clearly fall inside or outside of our scope. However, this leaves a lot of gray area
in between. For that, we ask you to use your best judgment, and welcome contributions to help
clarify the scope.

### In Scope

The following items are explicitly in-scope for the bug bounty program:

**Cluster Attacks:**

- Attacks against Beta & GA features, unless explicitly excluded below
- Privilege escalation due to bugs in RBAC, ABAC, pod security policies
- Authentication bugs in the in-tree authentication handlers<br>
  _Including: OIDC, x509 certificates, service accounts, webhook authenticator, bearer token, etc._
- Privilege escalation through the kubelet APIs
- Remote code execution in kubelet, api server
- Unauthorized etcd access via the Kubernetes API
- Path traversal attacks in API, namespaces, etcd
- Info leak (e.g. workload names) from publicly accessible unauthenticated endpoints<br>
  _Excluding intentionally disclosed info, such as Kubernetes version & enabled APIs_
- Reliable suppression of audit logs for privileged actions
- Unexpected editing, removal, or permission changes of files on the host filesystems from
  Kubernetes components (e.g. kubelet)
- Persistent DoS from within a cluster by an unprivileged container or user.

**Supply Chain:** _(excluding social engineering attacks against maintainers)_

- Unauthorized code commit to any Kubernetes org repository<br>
  _Including: `github.com/kubernetes{,-client,-csi,-incubator,-retired,-security,-sigs}/*`_
- Unauthorized access to github.com/kubernetes-security
- Publishing of unauthorized artifacts
- Unauthorized modification of github data
- CI/CD Credential Leaks
- Execution inside the CI/CD infrastructure

**Components:**

- Attacks against a stable & supported Kubernetes release (most recent 3 releases)
- Community maintained stable cloud platform plugins<br>
  _Vulnerabilities in other cloud platform plugins should be reported through the associated provider_
- In-tree (k8s.io/kubernetes) stable volume plugins

### Out of Scope

The following items are explicitly out-of scope for the bug bounty program. While we still welcome
vulnerability reports in these areas, they are not (currently) eligible to receive a bounty.

- Alpha features & APIs
- Kubernetes running on Windows or other non-linux operating systems
- Non-Kubernetes binaries distributed as cluster addons<br>
  _Please report vulnerabilities in these components through the appropriate channel for the
  upstream component_
- Container escalations and escapes to the host, unless the attack path traverses a Kubernetes
  process (e.g. kubelet).
- Linux privilege escalations<br>
  _Please report these through security@kernel.org_
- Attacks against containers from the host they are running on
- Attacks relying on insecure configurations (subject to the [Product Security Committee][]'s opinion),
  such as clusters not utilizing mutual authentication or encryption between Kubernetes components.
- Attacks relying on or against deprecated components (e.g. gitrepo volumes)
- Vulnerabilities in etcd<br>
  _Please report these through [CoreOS's disclosure process][]_
- Vulnerabilities in CoreDNS<br>
  _Please report these through [CoreDNS's disclosure process][]_
- Vulnerabilities specific to a hosted Kubernetes setup<br>
  _Please report these through the associated provider_

[Product Security Committee]: https://git.k8s.io/security/security-release-process.md#product-security-committee-psc
[CoreOS's disclosure process]: https://coreos.com/security/disclosure/
[CoreDNS's disclosure process]: https://github.com/coredns/coredns#security
