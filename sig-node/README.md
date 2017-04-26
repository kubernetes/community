# SIG Node

A Special Interest Group focused on topics related to the Kubernetes node, such as the Kubelet and
container runtime. It is an open and active community, and we always welcome new additions!

## Goals:

Topics include, but are not limited to:

* Kubelet related features (e.g. Pod lifecycle)
* Node level performance and scalability (with [sig-scalability](../sig-scalability))
* Node reliability
* Node lifecycle management (with [sig-cluster-lifecycle](../sig-cluster-lifecycle))
* Container runtimes: docker, [rkt](../sig-rktnetes), etc.
* Images, package management
* Resource management (with [sig-scheduling](../sig-scheduling))
* Issues related to monitoring (with [sig-instrumentation](../sig-instrumentation))
* Node level security and Pod isolation (with [sig-auth](../sig-auth))
* Kernel interactions (to a limited extent)
* ...

We also work closely with [sig-storage](../sig-storage) and [sig-networking](../sig-networking). As you can see, this is a very cross-functional team!

## Contact us:

* via [Slack](https://kubernetes.slack.com/messages/sig-node/)
* via [Google Groups](https://groups.google.com/forum/#!forum/kubernetes-sig-node)

## GitHub Aliases:

* sig-node-api-reviews - For @mentions on PRs that affect the Kubelet API, or node-specific parts of the core API
* sig-node-bugs - For @mentions on `kind/bug` issues
* sig-node-feature-requests - For @mentions on feature requests, and `kind/feature` issues, and [feature tracking](https://github.com/kubernetes/features)
* sig-node-pr-reviews - For @mentions on PRs that affect the Kubelet or other node level components
* sig-node-proposals - For @mentions on node related [proposals](../contributors/design-proposals)
* sig-node-test-failures - For @mentions on test failure issues that might be related to a node component

## Meetings:

### SIG Node Weekly

* We meet weekly on Tuesdays at 10:00AM PST (UTC-8)
* Hangouts Link: https://plus.google.com/hangouts/_/google.com/sig-node-meetup (It may take a minute to join)
* [Meeting Notes](https://docs.google.com/document/d/1Ne57gvidMEWXR70OxxnRkYquAoMpt56o75oZtg-OeBg/edit?usp=sharing)

### Resource Management Working Group

The resource management working group is a cross-team effort with sig-scheduling with a focus on
improving Kubernetes resource management.

* The working group meets Tuesdays at 11:00AM PST (UTC-8)
* Zoom Link: https://zoom.us/j/4799874685
* [Agenda doc](https://docs.google.com/document/d/1j3vrG6BgE0hUDs2e-1ZUegKN4W4Adb1B6oJ6j-4kyPU/edit#)
* [Meeting Recordings](https://www.youtube.com/playlist?list=PL69nYSiGNLP1wJPj5DYWXjiArF-MJ5fNG)

## Team:

* Lead: [Dawn Chen](https://github.com/dchen1107) <dawnchen@google.com>, Google
* Load: [Derek Carr](https://github.com/derekwaynecarr) <decarr@redhat.com>, Red Hat
* Reviewers: See [OWNERS_ALIASES#sig-node-reviewers](https://github.com/kubernetes/kubernetes/blob/master/OWNERS_ALIASES)
* And too many regular contributors to list here...
