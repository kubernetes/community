
Welcome to Kubernetes! (New Developer Guide)
============================================

_This document assumes that you know what Kubernetes does. If you don't, check
out [the Kubernetes keynote](https://www.youtube.com/watch?v=8SvQqZNP6uo) and
try the demo at [https://k8s.io/](https://k8s.io/)._

Introduction
------------

Have you ever wanted to contribute to the coolest cloud technology? This
document will help you understand the organization of the Kubernetes project and
direct you to the best places to get started. By the end of this doc, you'll be
able to pick up issues, write code to fix them, and get your work reviewed and
merged.

If you have questions about the development process, feel free to jump into our
[Slack Channel](http://slack.k8s.io/) or join our [mailing
list](https://groups.google.com/forum/#!forum/kubernetes-dev).

Special Interest Groups
-----------------------

Kubernetes developers work in teams called Special Interest Groups (SIGs). At
the time of this writing there are [31 SIGs](../../sig-list.md).

The developers within each SIG have autonomy and ownership over that SIG's part
of Kubernetes. SIGs organize themselves by meeting regularly and submitting
markdown design documents to the
[kubernetes/community](https://github.com/kubernetes/community) repository. You
can see an example of the design document process at
[kubernetes/community#645](https://github.com/kubernetes/community/pull/645).
Like everything else in Kubernetes, a SIG is an open, community, effort. Anybody
is welcome to jump into a SIG and begin fixing issues, critiquing design
proposals and reviewing code.

Most people who visit the Kubernetes repository for the first time are
bewildered by the thousands of [open
issues](https://github.com/kubernetes/kubernetes/issues) in our main repository.
But now that you know about SIGs, it's easy to filter by labels to see what's
going on in a particular SIG. For more information about our issue system, check
out
[issues.md](https://github.com/kubernetes/community/blob/master/contributors/devel/issues.md).

Downloading, Building, and Testing Kubernetes
---------------------------------------------

This guide is non-technical, so it does not cover the technical details of
working Kubernetes. We have plenty of documentation available under
[github.com/kubernetes/community/contributors/devel](https://github.com/kubernetes/community/tree/master/contributors/devel).
Check out
[development.md](https://github.com/kubernetes/community/blob/master/contributors/devel/development.md)
for more details.

Pull-Request Process
--------------------

The pull-request process is documented in [pull-requests.md](pull-requests.md).
As described in that document, you must sign the [CLA](../../CLA.md) before
Kubernetes can accept your contribution.

There is an entire SIG
([sig-contributor-experience](../../sig-contributor-experience/README.md))
devoted to improving your experience as a contributor. Contributing to
Kubernetes should be easy. If you find a rough edge, let us know!  File an issue
in the appropriate repository if you know it, send a message on [our Slack
channel](https://kubernetes.slack.com/messages/C1TU9EB9S/details) or to the
[owners](https://github.com/kubernetes/community/blob/master/contributors/devel/owners.md)
of contributor-experience. Better yet, help us fix it by joining the SIG; just
show up to one of the [bi-weekly
meetings](https://docs.google.com/document/d/1qf-02B7EOrItQgwXFxgqZ5qjW0mtfu5qkYIF1Hl4ZLI/edit).

The Release Process and Code Freeze
-----------------------------------

Every so often @k8s-merge-robot will refuse to merge your PR, saying something
about release milestones. This happens when we are in a code freeze for a
release. In order to ensure Kubernetes is stable, we stop merging everything
that's not a bugfix, then focus on making all the release tests pass. This code
freeze usually lasts two weeks and happens once per quarter. See the [1.8 release
schedule](https://github.com/kubernetes/features/blob/master/release-1.8/release-1.8.md)
for more information.

If you're new to Kubernetes, you won't have to worry about this too much. After
you've contributed for a few months, you will be added as a [community
member](https://github.com/kubernetes/community/blob/master/community-membership.md)
and take ownership of some of the tests. At this point, you'll work with members
of your SIG to review PRs coming into your area and track down issues that occur
in tests.

Thanks for reading!


