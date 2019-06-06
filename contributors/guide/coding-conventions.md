---
title: "Coding Conventions"
weight: 1
slug: "coding-conventions"
---

**Table of Contents**

- [Coding Conventions](#coding-conventions)
  - [Code conventions](#code-conventions)
  - [Testing conventions](#testing-conventions)
  - [Directory and file conventions](#directory-and-file-conventions)

## Code conventions

  - Bash

    - https://google.github.io/styleguide/shell.xml

    - Ensure that build, release, test, and cluster-management scripts run on
macOS

  - Go

    - [Go Code Review
Comments](https://github.com/golang/go/wiki/CodeReviewComments)

    - [Effective Go](https://golang.org/doc/effective_go.html)

    - Know and avoid [Go landmines](https://gist.github.com/lavalamp/4bd23295a9f32706a48f)

    - Comment your code.
      - [Go's commenting
conventions](http://blog.golang.org/godoc-documenting-go-code)
      - If reviewers ask questions about why the code is the way it is, that's a
sign that comments might be helpful.

    - Command-line flags should use dashes, not underscores

    - Naming
      - Please consider package name when selecting an interface name, and avoid
redundancy.

          - e.g.: `storage.Interface` is better than `storage.StorageInterface`.

      - Do not use uppercase characters, underscores, or dashes in package
names.
      - Please consider parent directory name when choosing a package name.

          - so pkg/controllers/autoscaler/foo.go should say `package autoscaler`
not `package autoscalercontroller`.
          - Unless there's a good reason, the `package foo` line should match
the name of the directory in which the .go file exists.
          - Importers can use a different name if they need to disambiguate.

      - Locks should be called `lock` and should never be embedded (always `lock
sync.Mutex`). When multiple locks are present, give each lock a distinct name
following Go conventions - `stateLock`, `mapLock` etc.

    - [API changes](/contributors/devel/sig-architecture/api_changes.md)

    - [API conventions](/contributors/devel/sig-architecture/api-conventions.md)

    - [Kubectl conventions](/contributors/devel/sig-cli/kubectl-conventions.md)

    - [Logging conventions](/contributors/devel/sig-instrumentation/logging.md)

## Testing conventions

  - All new packages and most new significant functionality must come with unit
tests

  - Table-driven tests are preferred for testing multiple scenarios/inputs; for
example, see [TestNamespaceAuthorization](https://git.k8s.io/kubernetes/test/integration/auth/auth_test.go)

  - Significant features should come with integration (test/integration) and/or
[end-to-end (test/e2e) tests](/contributors/devel/sig-testing/e2e-tests.md)
    - Including new kubectl commands and major features of existing commands

  - Unit tests must pass on macOS and Windows platforms - if you use Linux
specific features, your test case must either be skipped on windows or compiled
out (skipped is better when running Linux specific commands, compiled out is
required when your code does not compile on Windows).

  - Avoid relying on Docker hub (e.g. pull from Docker hub). Use gcr.io instead.

  - Avoid waiting for a short amount of time (or without waiting) and expect an
asynchronous thing to happen (e.g. wait for 1 seconds and expect a Pod to be
running). Wait and retry instead.

  - See the [testing guide](/contributors/devel/sig-testing/testing.md) for additional testing advice.

## Directory and file conventions

  - Avoid package sprawl. Find an appropriate subdirectory for new packages.
(See [#4851](http://issues.k8s.io/4851) for discussion.)
    - Libraries with no more appropriate home belong in new package
subdirectories of pkg/util

  - Avoid general utility packages. Packages called "util" are suspect. Instead,
derive a name that describes your desired function. For example, the utility
functions dealing with waiting for operations are in the "wait" package and
include functionality like Poll. So the full name is wait.Poll

  - All filenames should be lowercase

  - Go source files and directories use underscores, not dashes
    - Package directories should generally avoid using separators as much as
possible (when packages are multiple words, they usually should be in nested
subdirectories).

  - Document directories and filenames should use dashes rather than underscores

  - Contrived examples that illustrate system features belong in
/docs/user-guide or /docs/admin, depending on whether it is a feature primarily
intended for users that deploy applications or cluster administrators,
respectively. Actual application examples belong in /examples.
    - Examples should also illustrate [best practices for configuration and using the system](https://kubernetes.io/docs/concepts/configuration/overview/)

  - Third-party code

    - Go code for normal third-party dependencies is managed using
[go modules](https://github.com/golang/go/wiki/Modules) and is described in the kubernetes
[vendoring guide](/contributors/devel/sig-architecture/vendor.md)

    - Other third-party code belongs in `/third_party`
      - forked third party Go code goes in `/third_party/forked`
      - forked _golang stdlib_ code goes in `/third_party/forked/golang`

    - Third-party code must include licenses

    - This includes modified third-party code and excerpts, as well
