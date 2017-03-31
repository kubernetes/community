# Development Guide

This document is intended to be the canonical source of truth for things like
supported toolchain versions for building Kubernetes. If you find a
requirement that this doc does not capture, please
[submit an issue](https://github.com/kubernetes/kubernetes/issues) on github. If
you find other docs with references to requirements that are not simply links to
this doc, please [submit an issue](https://github.com/kubernetes/kubernetes/issues).

This document is intended to be relative to the branch in which it is found.
It is guaranteed that requirements will change over time for the development
branch, but release branches of Kubernetes should not change.

## Pre submit flight checks

Make sure you decide whether your issue and/or pull request is improving kubernetes architecture or whehter its simply fixing a bug.

Make sure there are no typos, if you need a diagram, add it.  Make sure you SEPARATE the description of the problem
(i.e. Y is a critical component that is too slow for an SLA that we care about) from the solution (i.e. make X faster).

Some of these checks were less common in Kubernetes earlier days, but now having over 1000 contributors, each issue should be
filed with care, and should be sanity-checkable in under 5 minutes (even the busiest of reviewers can spare up to 5 minutes to
review a patch that is thoughtfully justified).

### Is this just a simple bug fix?

These patches can be easy to review since test coverage is submitted with the patch.  Bug fixes don't usually require alot
of extra testing: But please update the unit tests so that they catch the bug !

### Is this an architecture improvement?

Some examples of "Architecture" improvements:

- Adding a new feature or making a feature more configurable/modular.
- Converting structs to interfaces.
- Improving test coverage.
- Decoupling logic or creation of new utilities.
- Making code more resilient (sleeps, backoffs, reducing flakiness, etc).

These sorts of improvements are easily evaluated if they decrease lines of code without breaking functionality.

If you are improving the quality of code, then justify/state exactly what you 'cleaning up' in your Pull Request so as
not to waste reviewer time.

If you're making code more resilient, test it with a local cluster to demonstrate how exactly your patch changes
things.

Example: If you made a controller more robust to inconsistent data, make a mock object which returns incorrect data a
few times and verify the controllers behaviour accordingly.

### Is this a performance improvement ?

If you are submitting a performance bug, you MUST ALSO submit data that demonstrates your problem if you want the issue to
remain open.  This can be done locally using kubemark, scheduler_perf, unit tests, go benchmark tests, or e2e tests on
a real cluster with metrics plots.

Examples of how NOT to suggest a performance bug (these can really lead to a long review process and waste cycles):

- We *should* be doing X instead of Y because it *might* lead to better performance.
- Doing X instead of Y would reduce calls to Z.

The above statements have basically no value to a reviewer, because neither is a strong, testable, assertive statement.
This will land your PR in a no-man's-land zone (at best), or waste tons of time for a busy reviewer (at worst).

Of course any improvment is welcome, but performance improvements are the hardest to review.  They often make code more
complex, and to-often are not easily evaluated at review time due to lack of sufficient data submitted by the author
of a performance improvement patch.

Some examples of "Performance" improvements:

- Improving a caching implementation.
- Reducing calls to functions which are O(n^2), or reducing dependence on API server requests.
- Changing the value of default parameters for proceeses, or making those values 'smarter'.
- Parallelizing a calculation that needs to run on a large set of node/pod objects.

These issues should always be submitted with (in decreasing order or value):

- A golang Benchmark test.
- A visual depiction of reduced metric load on a cluster (measurable using metrics/ endpoints and grafana).
- A hand-instrumented timing test (i.e. adding some logs into the controller manager).

Without submitting data and results for your suggested performance improvements, its very possible that bikeshedding
about meaningless possible performance optimizations could waste both reviewer time as well as your own.

Some examples of properly submitted performance issues, from different parts of the codebase.  They all have one thing
in common: Lots of data in the issue definition.  If you are new to kubernetes and thinking about filing a performance
optimization, re-read one or all of these before you get started.

- https://github.com/kubernetes/kubernetes/issues/18266 (apiserver)
- https://github.com/kubernetes/kubernetes/issues/32833 (node)
- https://github.com/kubernetes/kubernetes/issues/31795 (scheduler)

Since performance improvements deal with empirical systems, one playing in this space should be intimately familiar with
the "scientific method" of creating a hypothesis, collecting data, and then revising your hypothesis.  The above issues
tend to do this transparently, using figures and data rather then theoretical postulations, as a first pass before a
single line of code is reviewed.

## Building Kubernetes with Docker

Official releases are built using Docker containers. To build Kubernetes using Docker please follow [these instructions]
(http://releases.k8s.io/HEAD/build/README.md).

## Building Kubernetes on a local OS/shell environment

Many of the Kubernetes development helper scripts rely on a fairly up-to-date
GNU tools environment, so most recent Linux distros should work just fine
out-of-the-box. Note that Mac OS X ships with somewhat outdated BSD-based tools,
some of which may be incompatible in subtle ways, so we recommend
[replacing those with modern GNU tools]
(https://www.topbug.net/blog/2013/04/14/install-and-use-gnu-command-line-tools-in-mac-os-x/).

### Go development environment

Kubernetes is written in the [Go](http://golang.org) programming language.
To build Kubernetes without using Docker containers, you'll need a Go
development environment. Builds for Kubernetes 1.0 - 1.2 require Go version
1.4.2. Builds for Kubernetes 1.3 and higher require Go version 1.6.0. If you
haven't set up a Go development environment, please follow [these
instructions](http://golang.org/doc/code.html) to install the go tools.

Set up your GOPATH and add a path entry for go binaries to your PATH. Typically
added to your ~/.profile:

```sh
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
```

### Godep dependency management

Kubernetes build and test scripts use [godep](https://github.com/tools/godep) to
manage dependencies.

#### Install godep

Ensure that [mercurial](http://mercurial.selenic.com/wiki/Download) is
installed on your system. (some of godep's dependencies use the mercurial
source control system). Use `apt-get install mercurial` or `yum install
mercurial` on Linux, or [brew.sh](http://brew.sh) on OS X, or download directly
from mercurial.

Install godep (may require sudo):

```sh
go get -u github.com/tools/godep
```

Note:
At this time, godep version >= v63 is known to work in the Kubernetes project.

To check your version of godep:

```sh
$ godep version
godep v74 (linux/amd64/go1.6.2)
```

Developers planning to managing dependencies in the `vendor/` tree may want to
explore alternative environment setups. See
[using godep to manage dependencies](godep.md).

### Local build using make

To build Kubernetes using your local Go development environment (generate linux
binaries):

```sh
        make
```

You may pass build options and packages to the script as necessary. For example,
to build with optimizations disabled for enabling use of source debug tools:

```sh
        make GOGCFLAGS="-N -l"
```

To build binaries for all platforms:

```sh
        make cross
```

### How to update the Go version used to test & build k8s

The kubernetes project tries to stay on the latest version of Go so it can
benefit from the improvements to the language over time and can easily
bump to a minor release version for security updates.

Since kubernetes is mostly built and tested in containers, there are a few
unique places you need to update the go version.

- The image for cross compiling in [build/build-image/cross/](https://github.com/kubernetes/kubernetes/blob/master/build/build-image/cross/). The `VERSION` file and `Dockerfile`.
- Update [dockerized-e2e-runner.sh](https://github.com/kubernetes/test-infra/blob/master/jenkins/dockerized-e2e-runner.sh) to run a kubekins-e2e with the desired go version, which requires pushing [e2e-image](https://github.com/kubernetes/test-infra/tree/master/jenkins/e2e-image) and [test-image](https://github.com/kubernetes/test-infra/tree/master/jenkins/test-image) images that are `FROM` the desired go version.
- The docker image being run in [gotest-dockerized.sh](https://github.com/kubernetes/test-infra/blob/master/jenkins/gotest-dockerized.sh).
- The cross tag `KUBE_BUILD_IMAGE_CROSS_TAG` in [build/common.sh](https://github.com/kubernetes/kubernetes/blob/master/build/common.sh)

## Workflow

Below, we outline one of the more common git workflows that core developers use.
Other git workflows are also valid.

### Visual overview

![Git workflow](git_workflow.png)

### Fork the main repository

1. Go to https://github.com/kubernetes/kubernetes
2. Click the "Fork" button (at the top right)

### Clone your fork

The commands below require that you have $GOPATH set ([$GOPATH
docs](https://golang.org/doc/code.html#GOPATH)). We highly recommend you put
Kubernetes' code into your GOPATH. Note: the commands below will not work if
there is more than one directory in your `$GOPATH`.

```sh
mkdir -p $GOPATH/src/k8s.io
cd $GOPATH/src/k8s.io
# Replace "$YOUR_GITHUB_USERNAME" below with your github username
git clone https://github.com/$YOUR_GITHUB_USERNAME/kubernetes.git
cd kubernetes
git remote add upstream 'https://github.com/kubernetes/kubernetes.git'
```

### Create a branch and make changes

```sh
git checkout -b my-feature
# Make your code changes
```

### Keeping your development fork in sync

```sh
git fetch upstream
git rebase upstream/master
```

Note: If you have write access to the main repository at
github.com/kubernetes/kubernetes, you should modify your git configuration so
that you can't accidentally push to upstream:

```sh
git remote set-url --push upstream no_push
```

### Committing changes to your fork

Before committing any changes, please link/copy the pre-commit hook into your
.git directory. This will keep you from accidentally committing non-gofmt'd Go
code. This hook will also do a build and test whether documentation generation
scripts need to be executed.

The hook requires both Godep and etcd on your `PATH`.

```sh
cd kubernetes/.git/hooks/
ln -s ../../hooks/pre-commit .
```

Then you can commit your changes and push them to your fork:

```sh
git commit
git push -f origin my-feature
```

### Creating a pull request

1. Visit https://github.com/$YOUR_GITHUB_USERNAME/kubernetes
2. Click the "Compare & pull request" button next to your "my-feature" branch.
3. Check out the pull request [process](pull-requests.md) for more details

**Note:** If you have write access, please refrain from using the GitHub UI for creating PRs, because GitHub will create the PR branch inside the main repository rather than inside your fork.

### Getting a code review

Once your pull request has been opened it will be assigned to one or more
reviewers.  Those reviewers will do a thorough code review, looking for
correctness, bugs, opportunities for improvement, documentation and comments,
and style.

Very small PRs are easy to review.  Very large PRs are very difficult to
review.  Github has a built-in code review tool, which is what most people use.
At the assigned reviewer's discretion, a PR may be switched to use
[Reviewable](https://reviewable.k8s.io) instead.  Once a PR is switched to
Reviewable, please ONLY send or reply to comments through reviewable.  Mixing
code review tools can be very confusing.

See [Faster Reviews](faster_reviews.md) for some thoughts on how to streamline
the review process.

### When to retain commits and when to squash

Upon merge, all git commits should represent meaningful milestones or units of
work.  Use commits to add clarity to the development and review process.

Before merging a PR, squash any "fix review feedback", "typo", and "rebased"
sorts of commits. It is not imperative that every commit in a PR compile and
pass tests independently, but it is worth striving for. For mass automated
fixups (e.g. automated doc formatting), use one or more commits for the
changes to tooling and a final commit to apply the fixup en masse. This makes
reviews much easier.

## Testing

Three basic commands let you run unit, integration and/or e2e tests:

```sh
cd kubernetes
make test # Run every unit test
make test WHAT=pkg/util/cache GOFLAGS=-v # Run tests of a package verbosely
make test-integration # Run integration tests, requires etcd
make test-e2e # Run e2e tests
```

See the [testing guide](testing.md) and [end-to-end tests](e2e-tests.md) for additional information and scenarios.

## Regenerating the CLI documentation

```sh
hack/update-generated-docs.sh
```




<!-- BEGIN MUNGE: GENERATED_ANALYTICS -->
[![Analytics](https://kubernetes-site.appspot.com/UA-36037335-10/GitHub/docs/devel/development.md?pixel)]()
<!-- END MUNGE: GENERATED_ANALYTICS -->
