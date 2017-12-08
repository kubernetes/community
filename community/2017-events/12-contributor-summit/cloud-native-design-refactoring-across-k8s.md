**Contributor Summit - KubeCon 2017**

**Cloud Native Design/Refactoring across Kubernetes**

*Lead: Joe Beda(**[@jbeda](https://twitter.com/jbeda)*)*

*Co-lead: Amit Kumar Jaiswal(**[@AMIT_GKP](https://twitter.com/AMIT_GKP)*)*

**Abstract**

Explore how to cleanly support cloud-native behaviors, such as standardized Kubernetes logs, injectable configuration and common queryable APIs in Kubernetes components. While this discussion isn't only about containers, orchestrating OpenStack deployment/management within Kubernetes via OpenStack-Helm or Kolla-Kubernetes paves the way for better upgrade capabilities. They will also improve the ability to run individual Kubernetes services independently or in combination with other adjacent technologies. 

**Agenda**

1. Cloud Native Ecosystems

2. Kubernetes Abstractions & Primitives

3. Kubernetes Design Patterns

4. Refactoring across Kubernetes

5. Benefits of using Kubernetes

**Issues**

**	**- We’re looking for someone to help us out on issues related to refactoring across Kubernetes like [Issues#51405](https://github.com/kubernetes/kubernetes/issues/51405), [#46735](https://github.com/kubernetes/kubernetes/issues/46735), [#54090](https://github.com/kubernetes/kubernetes/issues/54090), [#55151](https://github.com/kubernetes/kubernetes/issues/55151)

	- Consolidating volume util files like *pkg/volume/util.go, pkg/volume/util/util.go, pkg/volume/util/volumehelper/volumehelper.go *[need better documentation of each file is to emcompass]

	- Enhancing e2e tests to track cloud provider’s API Quotas

	- Like changing fluentd to use CRI log Path and eventually deprecates the old container path

	- Issues with deploying/adopting Kubernetes for specific applications use-cases

**Questions**

	- Security and granularity across Kubernetes

	- Kube API issues like it should not depend on *--cloud-provider* and *--cloud-config*

	- Helping us out with API documentation and Configuring Best Practices?

	- What are the new tools to deal with testing frameworks for NFV/SDN across Kubernetes?

	- Kubelet Flag subfields precedence v/s files/ConfigMaps to Kubelet Config

	- How to dynamically provision an AWS EBS volume from a snapshot?

	- Work around Object definition in OpenAPI schema like definition of the *PersistentVolumeSpec, PersistenceVolumeClaimSpec*

	- Work around support for FUSE client for K8s cephfs.

	

**Audience Feedback**

	-

