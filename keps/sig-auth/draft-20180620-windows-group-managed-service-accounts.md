---
kep-number: 0
title: Windows Group Managed Service Accounts for Container Identity
authors:
  - "@patricklang"
owning-sig: "wg-container-identity"
participating-sigs:
  - "sig-auth"
  - "sig-windows"
reviewers:
  - TBD
approvers:
  - "@sig-auth-proposals"
editor: TBD
creation-date: 2018-06-20
last-updated: 2018-06-20
status: provisional
---

# Windows Group Managed Service Accounts for Container Identity


## Table of Contents

* [Table of Contents](#table-of-contents)
* [Summary](#summary)
* [Motivation](#motivation)
    * [Goals](#goals)
    * [Non-Goals](#non-goals)
* [Proposal](#proposal)
    * [User Stories [optional]](#user-stories-optional)
      * [Web Applications with MS SQL Server](#Web-Applications-with-MS-SQL-Server)
    * [Implementation Details/Notes/Constraints [optional]](#implementation-detailsnotesconstraints-optional)
    * [Risks and Mitigations](#risks-and-mitigations)
* [Graduation Criteria](#graduation-criteria)
* [Implementation History](#implementation-history)
* [Drawbacks [optional]](#drawbacks-optional)
* [Alternatives [optional]](#alternatives-optional)

## Summary

Active Directory is a service that is built-in and commonly used on Windows Server deployments for user and computer identity. Apps are run using Active Directory identities to enable common user to service, and service to service authentication and authorization. This proposal aims to support a specific type of identity, Group Managed Service Accounts (GMSA), for use with Windows Server containers. This will allow an operator to choose a GMSA at deployment time, and run containers using it to connect to existing applications such as a database or API server without changing how the authentication and authorization are performed.

## Motivation

There has been a lot of interest in supporting GMSA for Windows containers since it's the only way for a Windows application to use an Active Directory identity. This is shown in asks and questions from both public & private conversations:

- https://github.com/kubernetes/kubernetes/issues/51691 "For windows based containers, there is a need to access shared objects using domain user contexts."
- Multiple large customers are asking the Microsoft Windows team to enable this feature through container orchestrators
- Multiple developers have blogged how to use this feature, but all are on standalone machines instead of orchestrated through Kubernetes
  - https://artisticcheese.wordpress.com/2017/09/09/enabling-integrated-windows-authentication-in-windows-docker-container/
  - https://community.cireson.com/discussion/3853/example-cireson-scsm-portal-on-docker-windows-containers
  - https://cloudiqtech.com/windows-2016-docker-containers-using-gmsa-connect-to-sql-server/
  - https://success.docker.com/api/asset/.%2Fmodernizing-traditional-dot-net-applications%2F%23integratedwindowsauthentication 


### Goals

- Windows users can run containers with an existing GMSA identity on Kubernetes
- No extra files or Windows registry keys are needed on each Windows node. All needed data should flow through Kubernetes+Kubelet APIs
- Provide a reasonable effort to prevent pods from being inadvertently scheduled with access to a GMSA that the operator shouldn't have rights to

### Non-Goals

- Running Linux applications using GMSA
- Replacing any existing Kubernetes authorization or authentication controls
- Providing unique container identities. By design, Windows GMSA are used where multiple nodes are running apps as the same Active Directory principal.
- Isolation between container users and processes running as the GMSA. Windows already allows users and system services with sufficient privilege to create processes using a GMSA.
- Preventing the node from having access to the GMSA. Since the node already has authorization to access the GMSA, it can start processes or services using as the GMSA. Containers do not change this behavior.


## Proposal

### Background

#### What is Active Directory?
Windows applications and services typically use Active Directory identities to facilitate authentication and authorization between resources. In a traditional virtual machine scenario, the computer is joined to an Active Directory domain which enables it to use Kerberos, NTLM, and LDAP to identify and query for information about other resources in the domain. When a computer is joined to Active Directory, it is given an unique identity represented as a computer object in LDAP.

#### What is a Windows service account?
A Group Managed Service Account (gMSA) is a shared Active Directory identity that enables common scenarios such as authenticating and authorizing incoming requests and accessing downstream resources such as a database server, file share, or other workload. It can be used by multiple authorized users or computers at the same time.

#### How is it applied to containers?
To achieve the scale and speed required for containers, Windows uses a group managed service account in lieu of individual computer accounts to enable Windows containers to communicate with Active Directory. As of right now, Windows cannot use individual Active Directory computer & user accounts.

Different containers on the same machine can use different gMSAs to represent the specific workload they are hosting, allowing operators to granularly control which resources a container has access to. However, to run a container with a gMSA identity, an additional parameter must be supplied to the Windows Host Compute Service to indicate the intended identity. This proposal seeks to add support in Kubernetes for this parameter to enable Windows containers to communicate with other enterprise resources.

It's also worth noting that Docker implements this in a different way that's not managed centrally. It requires dropping a file on the node and referencing it by name, eg: docker run --credential-spec='file://foo.json' . For more details see the Microsoft doc.


### User Stories


#### Web Applications with MS SQL Server


A developer has a Windows web app that they would like to deploy in a container with Kubernetes. For example, it may have a web tier that they want to scale out running ASP.Net hosted in the IIS web server. Backend data  is stored in a Microsoft SQL database, which all of the web servers need to access behind a load balancer. An Active Directory Group Managed Service Account is used to avoid hardcoded passwords, and the web servers run with that credential today. The SQL Database has a user with read/write access to that GMSA so the web servers can access it. When they move the web tier into a container, they want to preserve the existing authentication and authorization models.

When this app moves to production on containers, the team will need to coordinate to make sure the right GMSA is carried over to the container deployment.

1. Active Directory domain administrator will:

  - Join all Windows Kubernetes nodes to the Active Directory domain.
  - Provision the service account and gives details to application admin.
  - Assign access to a security group to control what machines can use this service account. This group should include all authorized Kubernetes nodes.

2. A Kubernetes admin will create a namespace, secret and service account:
  - Create a unique Kubernetes [service account](https://kubernetes.io/docs/tasks/configure-pod-container/configure-service-account/#use-multiple-service-accounts) as an identity for the application in Kubernetes. No rights to the cluster are assigned.

```
apiVersion: v1
kind: ServiceAccount
metadata:
  name: webserver
```

  - Run a Windows PowerShell script to verify the account exists, and capture enough data to uniquely identify the GMSA into a JSON file (credspec). This doesn't contain any passwords or crypto secrets.

Example credspec.json
```
{
  "CmsPlugins": [
    "ActiveDirectory"
  ],
  "DomainJoinConfig": {
    "DnsName": "contoso.com",
    "Guid": "244818ae-87ca-4fcd-92ec-e79e5252348a",
    "DnsTreeName": "contoso.com",
    "NetBiosName": "DEMO",
    "Sid": "S-1-5-21-2126729477-2524075714-3094792973",
    "MachineAccountName": "WebApplication1"
  },
  "ActiveDirectoryConfig": {
    "GroupManagedServiceAccounts": [
      {
        "Name": "WebApplication1",
        "Scope": "DEMO"
      },
      {
        "Name": "WebApplication1",
        "Scope": "contoso.com"
      }
    ]
  }
}
```

  - Create a secret in the namespace of type `WindowsCredentialSpec` with the contents of the JSON - `kubectl create secret -n ... generic webserverCredSpec --from-file credspec.json`
  - Add the secret to the service account - `kubectl patch serviceaccount/webserver -p '{\"windowsCredentialSpec\": \"webserverCredSpec\"}'`
3. The Kubernetes admin gives a set of users rights to the namespace and service account
4. Application admin deploys the app with `ServiceAccount: webserver` as part of the PodSpec
5. The admission controller validates this user has access to the service account, and sets `PodSpec.windowsSecurityContext.CredentialSpec` to the contents of the secret
6. The Windows CRI implementation copies the credentialSpec to the [OCI windows.CredentialSpec](https://github.com/opencontainers/runtime-spec/blob/master/config-windows.md#credential-spec) field
7. The Windows OCI implementation validates the credentialspec, and fails to start the container if its invalid or access to the GMSA from the node is denied
8. The container starts with the Windows Group Managed Service Account as expected, and can authenticate to the database.


### Implementation Details/Notes/Constraints [optional]

What are the caveats to the implementation?
What are some important details that didn't come across above.
Go in to as much detail as necessary here.
This might be a good place to talk about core concepts and how they releate.

The Windows OCI runtime already has support for `windows.CredentialSpec` and is implemented in Moby. 


### Risks and Mitigations

What are the risks of this proposal and how do we mitigate.
Think broadly.
For example, consider both security and how this will impact the larger kubernetes ecosystem.

## Graduation Criteria

- alpha - initial implementation
- beta - design validated by two or more customer proof of concept deployments, recorded success in SIG-Windows meetings or mailing lists
- ga - Node E2E tests in place, tagged with features GMSA & Windows


## Implementation History

Major milestones in the life cycle of a KEP should be tracked in `Implementation History`.
Major milestones might include

- the `Summary` and `Motivation` sections being merged signaling SIG acceptance
- the `Proposal` section being merged signaling agreement on a proposed design
- the date implementation started
- the first Kubernetes release where an initial version of the KEP was available
- the version of Kubernetes where the KEP graduated to general availability
- when the KEP was retired or superseded

## Drawbacks [optional]

Why should this KEP _not_ be implemented.

## Alternatives 

### Other authentication methods

There are other ways to handle user-service and service-service authentication, but they generally require code changes. This proposal is focused on enabling customers to use existing on-premises Active Directory identity in containers.

For cloud-native applications, there are other alternatives:

- Kubernetes secrets - if both services are run in Kubernetes, this can be used for username/password or preshared secrets available to each app
- PKI - If you have a PKI infrastructure, you could choose to deploy application-specific certificates and change applications to trust specific public keys or intermediate certificates
- Cloud-provider service accounts - there may be other token-based providers available in your cloud. Apps can be modified to use these tokens and APIs for authentication and authorization requests.


### Using secrets without a ServiceAccount

This approach was originally described in https://github.com/kubernetes/kubernetes/issues/62038. It was intended to be the simplest possible implementation, but did not map to any existing Kubernetes service accounts or namespaces. 



<!-- end matter --> 
<!-- references -->
[oci-runtime](https://github.com/opencontainers/runtime-spec/blob/master/config-windows.md#credential-spec)
[manage-serviceaccounts](https://docs.microsoft.com/en-us/virtualization/windowscontainers/manage-containers/manage-serviceaccounts)