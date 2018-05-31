---
kep-number: 10
title: Graduate CoreDNS to GA
authors:
  - "@johnbelamaric"
  - "@rajansandeep"
owning-sig: sig-network
participating-sigs:
  - sig-cluster-lifecycle
reviewers:
  - "@bowei"
  - "@thockin"
approvers:
  - "@thockin"
editor: "@rajansandeep"
creation-date: 2018-03-21
last-updated: 2018-05-18
status: provisional
see-also: https://github.com/kubernetes/community/pull/2167
---
 
# Graduate CoreDNS to GA

## Table of Contents

* [Summary](#summary)
* [Motivation](#motivation)
    * [Goals](#goals)
    * [Non-Goals](#non-goals)
* [Proposal](#proposal)
    * [User Cases](#use-cases)
* [Graduation Criteria](#graduation-criteria)
* [Implementation History](#implementation-history)

## Summary

CoreDNS is sister CNCF project and is the successor to SkyDNS, on which kube-dns is based. It is a flexible, extensible
authoritative DNS server and directly integrates with the Kubernetes API. It can serve as cluster DNS,
complying with the [dns spec](https://git.k8s.io/dns/docs/specification.md). As an independent project,
it is more actively developed than kube-dns and offers performance and functionality beyond what kube-dns has. For more details, see the [introductory presentation](https://docs.google.com/presentation/d/1v6Coq1JRlqZ8rQ6bv0Tg0usSictmnN9U80g8WKxiOjQ/edit#slide=id.g249092e088_0_181), or [coredns.io](https://coredns.io), or the [CNCF webinar](https://youtu.be/dz9S7R8r5gw).

Currently, we are following the road-map defined [here](https://github.com/kubernetes/features/issues/427). CoreDNS is Beta in Kubernetes v1.10, which can be installed as an alternate to kube-dns.
The purpose of this proposal is to graduate CoreDNS to GA.

## Motivation

* CoreDNS is more flexible and extensible than kube-dns. 
* CoreDNS is easily extensible and maintainable using a plugin architecture.
* CoreDNS has fewer moving parts than kube-dns, taking advantage of the plugin architecture, making it a single executable and single process.
* It is written in Go, making it memory-safe (kube-dns includes dnsmasq which is not). 
* CoreDNS has [better performance](https://github.com/kubernetes/community/pull/1100#issuecomment-337747482) than [kube-dns](https://github.com/kubernetes/community/pull/1100#issuecomment-338329100) in terms of greater QPS, lower latency, and lower memory consumption. 

### Goals

* Bump up CoreDNS to be GA.
* Make CoreDNS available as an image in a Kubernetes repository (To Be Defined) and ensure a workflow/process to update the CoreDNS versions in the future.
  May be deferred to [next KEP](https://github.com/kubernetes/community/pull/2167) if goal not achieved in time.
* Provide a kube-dns to CoreDNS upgrade path with configuration translation in `kubeadm`.
* Provide a CoreDNS to CoreDNS upgrade path in `kubeadm`.

### Non-Goals

* Translation of CoreDNS ConfigMap back to kube-dns (i.e., downgrade).
* Translation configuration of kube-dns to equivalent CoreDNS that is defined outside of the kube-dns ConfigMap. For example, modifications to the manifest or `dnsmasq` configuration.
* Fate of kube-dns in future releases, i.e. deprecation path.
* Making [CoreDNS the default](https://github.com/kubernetes/community/pull/2167) in every installer.

## Proposal

The proposed solution is to enable the selection of CoreDNS as a GA cluster service discovery DNS for Kubernetes.
Some of the most used deployment tools have been upgraded by the CoreDNS team, in cooperation of the owners of these tools, to be able to deploy CoreDNS:
* kubeadm
* kube-up
* minikube
* kops

For other tools, each maintainer would have to add the upgrade to CoreDNS.

### Use Cases

* CoreDNS supports all functionality of kube-dns and also addresses [several use-cases kube-dns lacks](https://github.com/kubernetes/community/blob/master/contributors/design-proposals/network/coredns.md#use-cases). Some of the Use Cases are as follows: 
    * Supporting [Autopath](https://coredns.io/plugins/autopath/), which reduces the high query load caused by the long DNS search path in Kubernetes.
    * Making an alias for an external name [#39792](https://github.com/kubernetes/kubernetes/issues/39792)
    
* By default, the user experience would be unchanged. For more advanced uses, existing users would need to modify the ConfigMap that contains the CoreDNS configuration file.
* Since CoreDNS has more supporting features than kube-dns, there will be no path to retain the CoreDNS configuration in case a user wants to switch to kube-dns.

#### Configuring CoreDNS

The CoreDNS configuration file is called a `Corefile` and syntactically is the same as a [Caddyfile](https://caddyserver.com/docs/caddyfile). The file consists of multiple stanzas called _server blocks_.
Each of these represents a set of zones for which that server block should respond, along with the list of plugins to apply to a given request. More details on this can be found in the 
[Corefile Explained](https://coredns.io/2017/07/23/corefile-explained/) and [How Queries Are Processed](https://coredns.io/2017/06/08/how-queries-are-processed-in-coredns/) blog entries.

The following can be expected when CoreDNS is graduated to GA.

#### Kubeadm

* The CoreDNS feature-gates flag will be marked as GA.
* As Kubeadm maintainers chose to deploy CoreDNS as the default Cluster DNS for Kubernetes 1.11:
    * CoreDNS will be installed by default in a fresh install of Kubernetes via kubeadm.
    * For users upgrading Kubernetes via kubeadm, it will install CoreDNS by default whether the user had kube-dns or CoreDNS in a previous kubernetes version.
    * In case a user wants to install kube-dns instead of CoreDNS, they have to set the feature-gate of CoreDNS to false. `--feature-gates=CoreDNS=false`
* When choosing to install CoreDNS, the configmap of a previously installed kube-dns will be automatically translated to the equivalent CoreDNS configmap.

#### Kube-up

* CoreDNS will be installed when the environment variable `CLUSTER_DNS_CORE_DNS` is set to `true`. The default value is `false`.

#### Minikube

* CoreDNS to be an option in the add-on manager, with CoreDNS disabled by default.

## Graduation Criteria

* Verify that all e2e conformance and DNS related tests (xxx-kubernetes-e2e-gce, ci-kubernetes-e2e-gce-gci-ci-master and filtered by `--ginkgo.skip=\\[Slow\\]|\\[Serial\\]|\\[Disruptive\\]|\\[Flaky\\]|\\[Feature:.+\\]`) run successfully for CoreDNS.
  None of the tests successful with Kube-DNS should be failing with CoreDNS.
* Add CoreDNS as part of the e2e Kubernetes scale runs and ensure tests are not failing.
* Extend [perf-tests](https://github.com/kubernetes/perf-tests/tree/master/dns) for CoreDNS.
* Add a dedicated DNS related tests in e2e scalability test [Feature:performance].

## Implementation History

* 20170912 - [Feature proposal](https://github.com/kubernetes/features/issues/427) for CoreDNS to be implemented as the default DNS in Kubernetes.
* 20171108 - Successfully released [CoreDNS as an Alpha feature-gate in Kubernetes v1.9](https://github.com/kubernetes/kubernetes/pull/52501).
* 20180226 - CoreDNS graduation to Incubation in CNCF.
* 20180305 - Support for Kube-dns configmap translation and move up [CoreDNS to Beta](https://github.com/kubernetes/kubernetes/pull/58828) for Kubernetes v1.10.
