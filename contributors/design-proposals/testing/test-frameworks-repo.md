# New git repo: k8s.io/testframeworks

Status: Pending

Version: N/A

Implementation Owners: @totherme, @hoegaarden, @apelisse, and @spiffxp

## Motivation: Why do we need a new repo?

To avoid an vendoring cycle.

Some folks in Sig-CLI have been writing an [integration testing
framework](https://github.com/kubernetes/kubectl/tree/master/pkg/framework/test)
which they would like to use to simplify some of their integration tests which
live in `k8s.io/kubernetes`. The framework currently lives in `k8s.io/kubectl`.
Since `k8s.io/kubectl` already vendors components _from_ `k8s.io/kubernetes`, it 
cannot be vendored _into_ `k8s.io/kubernetes`. 

It therefore seems like a good idea to create a new repo to contain testing
frameworks. The intent is that this repo could be vendored anywhere.

## Proposal

We propose a new git repository named `testing.k8s.io/frameworks`. This would be
backed by a github repository `github.com/kubernetes-sig-testing/frameworks`.

### What should live in it?

The [integration testing
framework](https://github.com/kubernetes/kubectl/tree/master/pkg/framework/test)
should live in `testing.k8s.io/frameworks/integration`.

There also [seems to be a
desire](https://groups.google.com/a/kubernetes.io/d/msg/steering/LA9WiFnl6PI/os48-c3HCgAJ)
to get `kubernetes/test/e2e/framework` out into a more reusable place. It could
live in `testing.k8s.io/frameworks/e2e`. However, to make this happen we'll
need to do a lot of work to remove some of the less necessary dependencies of
the existing e2e framework.

### What repos can be imported (and hence vendored) into this repo?

No kubernetes repos other than [utils](https://github.com/kubernetes/utils)
should be imported (and hence vendored) by any code in this repo.

This means that any testing framework in this repo which interacts with any
part of kubernetes has to do so through some CLI. The [CLI integration testing
framework](https://github.com/kubernetes/kubectl/tree/master/pkg/framework/test)
which originally motivated this repo certainly does not have any kubernetes
dependencies.

We might revisit this decision in future, for example to potentially allow the
usage of a client library if that turns out to both:
- Be necessary for some important framework (such as e2e)
- Not harm the existing usage of other frameworks in the repo

### Other naming concerns

We would like the name of the repo to be a reasonable [golang package
name](https://blog.golang.org/package-names).  We should therefore probably
avoid characters that are not lower case alphabetic characters. This precludes
for example `k8s.io/test-frameworks`.

### Who should own it?

The consensus in [this email
thread](https://groups.google.com/a/kubernetes.io/d/msg/steering/LA9WiFnl6PI/os48-c3HCgAJ)
seems to be that in the context of SIGs needing to own all things, sig-testing
are the natural owners.

In the context of day-to-day looking after the integration test framework which
would initially live in this repo: @apelisse, @totherme, and @hoegaarden are
currently doing this in the kubectl repo, and are happy to keep doing so here.

## Alternatives considered

In [this email
thread](https://groups.google.com/a/kubernetes.io/forum/?utm_medium=email&utm_source=footer#!msg/steering/LA9WiFnl6PI/DjiPaN-2CgAJ)
and in [the sig-architecture meeting on
2018-01-04](https://docs.google.com/document/d/1BlmHq5uPyBUDlppYqAAzslVbAO8hilgjqZUTaNXUhKM/edit#heading=h.dm9wr8ympgj8)
the following alternatives have been considered and found to be unpopular:

- Create this repo in the `kubernetes` github org.
  + There is already a proposal to create one org per sig
  + @spiffxp from sig-testing expressed an interest in supporting this repo
    from sig-testing
  + it was decided to create an org for sig-testing to trial the whole
    one-org-per-sig concept, with this repo
- Create a new repo which only contains the integration testing framework, and
  cannot also be used to house the e2e framework.
   + The consensus seems to be that this leads to ugly package names
     (`k8s.io/integrationtestframework` or `k8s.io/integration` for example),
     and needlessly deprives the e2e framework of a useful landing place in the
     event that it wants to move somewhere more vendorable.
