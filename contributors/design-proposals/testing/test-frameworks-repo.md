# New git repo: k8s.io/testframeworks

Status: Pending

Version: N/A

Implementation Owners: @totherme @hoegaarden and @apelisse

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

We propose a new git repository named `k8s.io/testframeworks`. This would be
backed by a github repository `github.com/kubernetes/testframeworks`.

### What should live in it?

The [integration testing
framework](https://github.com/kubernetes/kubectl/tree/master/pkg/framework/test)
should live in `k8s.io/testframeworks/integration`.

There also [seems to be a
desire](https://groups.google.com/a/kubernetes.io/d/msg/steering/LA9WiFnl6PI/os48-c3HCgAJ)
to get `kubernetes/test/e2e/framework` out into a more reusable place. It could
live in `k8s.io/testframeworks/e2e`

### What repos can be imported (and hence vendored) into this repo?

No kubernetes repos other than [utils](https://github.com/kubernetes/utils)
should be imported (and hence vendored) by any code in this repo.

This means that any testing framework in this repo which interacts with any
part of kubernetes has to do so through some CLI. The [CLI integration testing
framework](https://github.com/kubernetes/kubectl/tree/master/pkg/framework/test)
which originally motivated this repo certainly does not have any kubernetes
dependencies.

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
the following alternatives have been considered and found to be unpopular:

- Create a new github org for new repos (such as this one) to exist in.
   + The consensus seems to be that this involves more organisational overhead than it's worth.
- Create a new repo which only contains the integration testing framework, and
  cannot also be used to house the e2e framework.
   + The consensus seems to be that this leads to ugly package names
     (`k8s.io/integrationtestframework` or `k8s.io/integration` for example),
     and needlessly deprives the e2e framework of a useful landing place in the
     event that it wants to move somewhere more vendorable.
