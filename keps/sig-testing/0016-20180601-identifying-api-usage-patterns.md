---
kep-number: 16
title: Identifying API usage patterns with Applied Anthropology
authors:
  - "@hh"
owning-sig: sig-testing
participating-sigs:
  - sig-testing
  - sig-architecture
  - sig-apps
reviewers:
  - "@spiffxp"
  - "@AishSundar"
approvers:
  - "@WilliamDenniss"
editor: TBD
creation-date: 2018-06-27
last-updated: 2018-06-27
status: provisional
see-also:
  - KEP-15
---

# Identifying API usage patterns with Applied Anthropology

## Table of Contents

[Tools for generating][] a table of contents from markdown are available.

* [Table of Contents](#table-of-contents)
* [Summary](#summary)
* [Motivation](#motivation)
    * [Goals](#goals)
    * [Non-Goals](#non-goals)
* [Proposal](#proposal)
    * [User Stories [optional]](#user-stories-optional)
      * [Story 1](#story-1)
      * [Story 2](#story-2)
    * [Implementation Details/Notes/Constraints [optional]](#implementation-detailsnotesconstraints-optional)
    * [Risks and Mitigations](#risks-and-mitigations)
* [Graduation Criteria](#graduation-criteria)
* [Implementation History](#implementation-history)
* [Drawbacks [optional]](#drawbacks-optional)
* [Alternatives [optional]](#alternatives-optional)

[Tools for generating**: https://github.com/ekalinin/github-markdown-toc

## Summary

Who are you? Why are you here?

Let’s enable any application, using our official Kubernetes libraries, to include the answer to these questions for each interaction with the APIServer.

The aggregated clusterwide correlation of identity and user journey with each API request/response would provide the raw metadata necessary explore the unseen, yet interwoven patterns of real-world user journeys within the Kubernetes community.

## Motivation

We need an atlas of the invisible and undefined tribal patterns within our ecosystem that we create within our community.

This map would help chart our course of development, testing, and conformance based on how kubernetes actually usage.

### Goals

* Enable communication of ‘Who are you? Why are you here?’
  - for any application using kubernetes API
  - via the official protocols and libraries.

* Simple aggregation of this metadata
  - cluster wide
  - community wide

#### Collateral APISnoop Goals

* Easily usable/curated community wide aggregation point

* A curated dataset for public analysis

* A community set of insights using the public datasets
  * Endpoint mappings to projects / source code / functions that use them
    * Including e2e tests / steps within tests
    * Historical endpoint/function usage patterns over time
  * Common Patterns of real world use across API endpoints
    * How does the community use/do X?
  * Machine Learning
    * What are the unseen yet common patterns?
    * What projects are using similar techniques?

### Non-Goals

* Create a Rube Goldberg machine of complexity just to enable community insight
* Any of the Collateral APISnoop Goals
  * Not part of this KEP, but they are the underlying motivations and the main end-benefits to the community

## Proposal

To aggregation of identity and purpose at the time of API interactions, we need to:

1. Define ‘identity’ and ‘purpose’ _*The who and why*_
2. Enable generation at time of interaction: _*Instant introspection to answer the above*_
3. Collectively record the individually evolving ‘identity and purpose’ _*Sharing this information for all Kubernetes applications within a cluster*_

### API interaction Identity (Who are you?)

Current API interaction ‘identity’ is static and usually set in client-go via user-agent to something like:

```
e2e.test/v1.12.0 (linux/amd64) kubernetes/b143093
kube-apiserver/v1.12.0 (linux/amd64) kubernetes/b143093
kube-controller-manager/v1.12.0 (linux/amd64) kubernetes/b143093
kubectl/v1.12.0 (linux/amd64) kubernetes/b143093
kubelet/v1.12.0 (linux/amd64) kubernetes/b143093
kube-scheduler/v1.12.0 (linux/amd64) kubernetes/b143093
```

Ideally our revised ‘identity’ should tie an application back to particular src commit, though some programs (like kernel info via uname) also show compile time info like timestamp or build user/machine:

```
$ uname -a
Darwin Roadrunner.local 10.3.0 Darwin Kernel Version 10.3.0: Fri Feb 26 11:58:09 PST 2010; root:xnu-1504.3.12~1/RELEASE_I386 i386
```

### API interaction Purpose (Why are you here?)
We must define a simple to implement, but contextually significant, answer to the question: *Why are you here?* It’s difficult to glean the purpose of an application interaction by external inspection without asking this obvious question.

At the moment of making the API call, the application has access it’s own stack and history of source code location/lines and functions that brought it to make a request of an external API. Disabled by default, it could be enabled by setting a variable such as `KUBE_CLIENT_SUBMIT_PURPOSE`.

Allowing the application to supply this _‘mental snapshot of purpose’_ could be as simple as providing space in our protocol for including source and method callstacks.

### Self Identification and Purpose (What does introspection tell you?)

Introspection is available in many of the languages that have official Kubernetes client libraries. Go, Python, and Java all provide the ability to inspect the runtime and stack programmatically, and include source paths and line numbers.

It may help to provide an example introspection:

```json
"introspection": {
  "self-identity": "kube-apiserver/v1.12.0 (linux/amd64) b143093 compiled by CNCF Fri Feb 26 11:58:09 PST 2010",
  "current-purpose": [
    "k8s.io/client-go/rest.(*Request).Do()",
    "k8s.io/client-go/kubernetes/typed/admissionregistration/v1alpha1.(*initializerConfigurations).List()",
    "k8s.io/apiserver/pkg/admission/configuration.NewInitializerConfigurationManager.func1()",
    "k8s.io/apiserver/pkg/admission/configuration.(*poller).sync()",
    "k8s.io/apiserver/pkg/admission/configuration.sync)-fm()",
    "k8s.io/apimachinery/pkg/util/wait.JitterUntil.func1()",
    "k8s.io/apimachinery/pkg/util/wait.JitterUntil()",
    "k8s.io/apimachinery/pkg/util/wait.Until()",
    "runtime.goexit()"
  ],
  "current-reasoning": [
    "k8s.io/client-go/rest/request.go:807",
    "k8s.io/client-go/kubernetes/typed/admissionregistration/v1alpha1/initializerconfiguration.go:79",
    "k8s.io/apiserver/pkg/admission/configuration/initializer_manager.go:42",
    "k8s.io/apiserver/pkg/admission/configuration/configuration_manager.go:155",
    "k8s.io/apiserver/pkg/admission/configuration/configuration_manager.go:151",
    "k8s.io/apimachinery/pkg/util/wait/wait.go:133",
    "k8s.io/apimachinery/pkg/util/wait/wait.go:134",
    "k8s.io/apimachinery/pkg/util/wait/wait.go:88",
    "runtime/asm_amd64.s:2361"
  ],
}

```

### How do we communicate these larger concepts of identity and purpose?

Currently the freeform concept of identity is limited what can fit within the user-agent field. 
Support for recording the [user-agent field in our audit-events](https://github.com/kubernetes/kubernetes/pull/64812) was recently added, but our initial explorations depend on that field allowing up to ?k.

We need to explore alternatives for conveying identity and purpose.
  * use different channels
  * compress the information to fit in 1k
  * raise the user-agent size limit to ?k to include this approach.

### Tying it all together: (How do I turn this on?)

If all applications are compiled against a client-go (or other supported library) and support the env var `KUBE_CLIENT_SUBMIT_PURPOSE`, then deploying kubernetes itself with it set should enable all kubernetes components to begin transmitting identity and purpose.

Setting this variable on all pods could be accomplished with an admission or initialization controller would allow all other applications in the cluster to do the same.

Currently this data is transmitted via user-agent, so configuring an audit-logging webhook, dynamic or otherwise, would allow centralized aggregation.


### User Stories

#### Story 1

As a SIG member, who uses the components we curate and what are they doing with them?

#### Story 2

As a SIG member choosing test to write/upgrade to conformance tests, what patterns and endpoints occur within our community vs what we currently test for.

#### Story 3

As a developer creating ginko tests, I'd like to know the existing tests and applications that have similar patterns or hit the endpoints I'm interested in.

### Implementation Details/Notes/Constraints


Audit-logging is not yet dynamically configurable, but is being discussed in the Dynamic Audit Configuration KEP.

User-Agent is may not the field to use, considering the current expectation of what it might contain, both size and content wise.

The data is interesting, because you get to see the callstacks for all the components in kubernetes, identifying the functions and line numbers making the calls.

### Risks and Mitigations

Leaking callstacks from applications that don’t want to have the ability to be enabled.
The default would need to be off, only when configured to do so via a KUBE_CALLSTACK_HASH style env var.

To limit exposing local path names and source, client-go could instead generate a hash of the data (generalized, so it's just the paths+linums under $GOPATH), however this would either reduce the data to, "I'm here for the same reason as last time, can't tell you what it is." While useful, it definately reduces our insight, or adds some complexity to map the hashes back to their full context.

## Graduation Criteria

How will we know that this has succeeded?
Gathering user feedback is crucial for building high quality experiences and SIGs have the important responsibility of setting milestones for stability and completeness.
Hopefully the content previously contained in [umbrella issues][] will be tracked in the `Graduation Criteria` section.

[umbrella issues]: https://github.com/kubernetes/kubernetes/issues/42752

## Implementation History

- 06/27/2018: initial design via google doc
- 07/11/2018: submission of KEP

## Drawbacks [optional]

Some feel this level of client ‘debugging’ doesn't belong in server-side logs. 

## Alternatives [optional]

Log sufficient information in the client-side logs to correlate the request with the audit logs.
The two sets of logs could then be combined programmatically as needed.

The best way to do this would be to log the audit-id header that is returned on the api response, and can be used to uniquely identify the corresponding audit logs.

Other projects that collect cluster-usage data:
 - https://github.com/kubernetes-incubator/spartakus
 - https://github.com/heptio/sonobuoy
