---
kep-number: 14
title: IAM Credentials Provider
status: provisional
authors:
    - "@pingles"
    - "@tombooth"
    - "@randomvariable"
owning-sig: sig-aws
creation-date: 2018-06-20
last-updated: 2018-06-20
---

# IAM Credentials Provider

## Table of Contents

A table of contents is helpful for quickly jumping to sections of a KEP and for highlighting any additional information provided beyond the standard KEP template.

<!-- markdown-toc start - Don't edit this section. Run M-x markdown-toc-refresh-toc -->
**Table of Contents**

- [IAM Credentials Provider](#iam-credentials-provider)
    - [Table of Contents](#table-of-contents**

<!-- markdown-toc end -->

## Summary
This KEP outlines a method to allow pods to retrieve credentials for AWS APIs. It allows per-process credentials to be used to improve the security of workloads running on Kubernetes clusters requiring access to AWS resources.

Earlier versions of this proposal were captured in this Google Doc: [SIG-AWS IAM Credentials Provider KEP](https://docs.google.com/document/d/1fNiGxHAcc5WACgU1JuoI_35w20v9zMddv7-ftKi_5Qg/edit#)

## Motivation
When using Kubernetes inside of AWS you’ll want to integrate with other services Amazon offer. The official AWS SDK provides a number of providers that are able to obtain credentials using IAM for authentication and authorization. The Kubernetes ecosystem needs a way of associating credentials to IAM roles to pods.

There are a number of existing tools that provide solutions to this problem with [different trade-offs](https://docs.google.com/document/d/1rn-v2TNH9k4Oz-VuaueP77ANE5p-5Ua89obK2JaArfg/edit). This proposal is for a project to overcome the shortfalls of each.

### Goals

* Provide a standard way for pods to securely retrieve IAM credentials from AWS
* Support multiple ways to link credentials to pods (service accounts & container identity)

### Non-Goals

* Support for other cloud providers outside of AWS
* Prevent pods from accessing the metadata API of the node
* Immediately supporting SPIFFE/SPIRE in pre-alpha

## Proposal
### User Stories

#### Story 1
Users run a Pod that contains an application using an AWS SDK. The application is able to obtain session-based credentials allowing it to make other authenticated AWS API calls (writing to S3, adding messages to a queue etc.).

Users are able to associate the [IAM Role](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles.html) that the Pod will be able to assume.

### Implementation Details/Notes/Constraints

1. Solution should work for existing applications that use an AWS SDK without requiring application code changes 
2. Cluster operators can run secure clusters by minimising IAM permissions needed on each node
3. Pods that need IAM credentials shouldn’t start until credentials are available
4. IAM messages/errors are accessible to cluster users
5. Credentials should be auditable against strongest Pod identity that can be verified
6. Credentials should be time-limited and refresh, allowing compromised credentials to be quickly invalidated

Our suggested implementation is composed of:

#### IAMBinding Custom Resource.
This will contain a spec associating a Service Account with an AWS IAM Role.

For example:

```yaml
---
apiVersion: v1
kind: Pod
spec:
  serviceAccountName: reporting-app
  ...

---
apiVersion: iam.aws.k8s.io/v1alpha1
kind: IAMBinding
metadata:
  name: reporting-role
spec:
  aws:
    roleArn: arn:aws:iam::123456789012:role/S3Access
  serviceAccountRef:
    name: reporting-app
```

#### Server
This process will verify the identity of the Pod and determine which IAM credentials it’s permitted to obtain, issue them using Amazon STS and return them. 

By setting a role session name to <namespace>.<service-account-name> (potentially with a cluster-specific suffix provided via a command-line flag** when obtaining credentials, we can use it in an assume policy to restrict which namespace or service account can use the role.

Pods will be associated to IAM roles using a Custom Resource: an IAMBinding. This will map between a Pod’s service account and the IAM role.

**Restricting roles to namespaces**

Our proposal would allow IAM assume policy to be set similarly to below for controlling which service accounts can assume which roles:

```json
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "",
      "Action": "sts:AssumeRole",
      "Effect": "Allow",
      "Principal": {
        "AWS": "arn:aws:iam::123456789012:role/kiam_role"
      },
      "Condition": {
        "StringLike": {
          "aws:userid": "cluster.namespace.service-account-name"
        }
      }
    }
  ]
}
```

#### Proxy
Any Pod needing credentials will have an init container and sidecar. 

The init container will ensure that other containers in the Pod when accessing `http://169.254.170.2` will have requests forwarded to the sidecar. The init container can also request initial credentials ensuring that the pod is entitled to use credentials for a role configured through the IAMBinding.

Once the init container has obtained the initial credentials the sidecar can run and refresh credentials periodically. 

The Proxy will read the Pod’s service account token and pass it to the Server for a TokenReview allowing the Server to attest to the Pod identity that is executing the Proxy against the Kubernetes API Server.

#### Admission Controller
This is responsible for identifying whether a Pod matches an IAMBinding and, if so, injects the init and sidecar.

#### Intercepting AWS APIs
We intend to follow the ECS credential provider, which uses an endpoint on http://169.254.170.2  that is typically provided by the ECS agent. In some of the standard SDKs this adds additional retry behaviour rather than the hard fail if the instance metadata API fails. Our goal would be to set environment variables that would trigger this behaviour and so mimic IAM integration provided in Amazon ECS.

There is potential for the workload container to start faster than the sidecar leading to failed requests. This is why we preferred use of the ECS credential provider, as it will retry on failure. The [SIG-Apps Sidecar proposal](https://github.com/kubernetes/community/pull/2148) would also help such an issue.

Using a per-Pod proxy helps mitigate many of the shortcomings of the existing DaemonSet-style solutions: IAM permissions can be checked before pods start, updates to the proxy are associated with the lifecycle of the Pod process, stronger identity attestation and removing inconsistency issues from tracking pods against the watcher. Existing solutions also often rely on per-node agents which, when restarted, can cause problems for the Pods running on the same nodes. Running a proxy sidecar against the Pod mitigates this.

Our assumption is that proxies can be injected in the same way Istio injects it’s proxy (iptables from the init container).

### Risks and Mitigations
#### Security
Our proposal uses a Server process that will receive ServiceAccount tokens from a Proxy process (either running during the init or sidecar phase) to attest to its identity, and allow the Server to determine whether an IAMBinding exists that maps it to an IAM Role. In this model the Server would receive credentials allowing it to perform any Kubernetes API operation permitted to those service accounts. Longer-term we’d hope to integrate with something like [SPIRE](https://spiffe.io/spire/).

#### Reliability
The Server process will become a single-point-of-failure for IAM on the cluster so should be able to operate with multiple replicas.

#### Scalability
All Pods would obtain unique credentials but would have their session names associated to the service account name. API calls would grow in-line with the number of Pods running (sidecar proxies would result in an sts:AssumeRole call by the server on it’s behalf) potentially resulting in lots of API calls. Caching could be used by the server to store previously obtained credentials for the service account but makes running an HA server more difficult.

#### Backwards and Forwards Compatibility
Our proposal doesn’t require changes to Kubernetes and can be deployed by people running their own clusters in AWS, clusters created with Kops or Kube-aws, or EKS.

In the future, we would expect to support container identity directly using SPIFFE/SPIRE, with the pod proxy using the SVID-encoded certificate to authenticate against the server process to retrieve credentials. This could be mapped using a pod selector, for example:

```yaml
---
apiVersion: v1
kind: Pod
metadata:
  name: reporting-app
  labels:
    app: reporting-app
  ...

---
apiVersion: iam.aws.k8s.io/v2alpha1
kind: IAMPodSelector
metadata:
  name: reporting-app
spec:
  aws:
    roleArn: arn:aws:iam::123456789012:role/S3Access
  podSelector:
    app: reporting-app
```

This should also allow for different server / proxy implementations should EKS implement the ability to directly issue credentials in the same way as ECS currently does today as well as allowing custom implementations for virtual kubelet / Fargate if needed.

#### Jobs
The model in this proposal will not currently work for jobs as the sidecar will prevent the job from terminating. Either implementation of the sidecar containers KEP or only running the init container and mapping credentials to a mount using the AWS_CREDENTIALS_FILE url with time-limited credentials would mitigate against this.

## Graduation Criteria

* Implementation of the IAMBinding CustomResourceDefintiion, init container, proxy sidecar and server process.
* e2e tests
* Add-on available for kops / helm, deployment tool of choice

## Alternatives

A number of existing solutions currently exist, which can be broken up into two main categories, one which is similar to the approach we propose here, and another which is significantly different:

* [kube2iam](https://github.com/jtblin/kube2iam)
* [kiam](https://github.com/uswitch/kiam)
* [iam4kube](https://github.com/ash2k/iam4kube)

Most implementations are based on methods first introduced in kube2iam which rely on routing instance profile credentials that can be retrieved on EC2 instances from `http://169.254.169.254/latest/meta-data/iam/security-credentials/<role_name>`

Kube2iam and kiam both implement on-node proxies, with nodes directly assuming roles via STS in kube2iam, and proxying to secure “server” nodes in the case of kiam reducing the attack surface. kube2iam has a number of race conditions around pod startup time as well as consistency between the process issuing credentials to pods which maps IP addresses to credentials, and the actual state of the world. Both kiam and iam4kube do address these issues, through different locking mechanisms. In addition, iam4kube and kiam prefetch credentials to allow for faster startup.

Iam4kube differs from the other implementations in explicitly mapping service accounts to IAM roles, and also directly routing requests for credentials to server nodes. 

* [kube-aws-iam-controller](https://github.com/mikkeloscar/kube-aws-iam-controller)
This project runs as a single instance on a specified node, and works by assuming a role using the STS API, and then injecting a secret’s mount into pod which AWS credentials in a well-known location. This avoids the race conditions of other proxy-based controllers, but suffers in that this usually requires one or more of either requiring application changes to explicitly load credentials from the mounted location or more importantly, requiring pods to be rotated when the credential expires since the AWS SDKs do not allow for refreshing credentials which are read from file, making it unsuitable beyond short-lived pods.
