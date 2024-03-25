---
title: "Coding Conventions"
weight: 8
description: |
  This document outlines a collection of guidelines, style suggestions, and tips
  for writing code in the different programming languages used throughout the
  Kubernetes project.
---

## Code conventions

  - Bash
    - [Shell Style Guide](https://google.github.io/styleguide/shellguide.html)
    - Ensure that build, release, test, and cluster-management scripts run on macOS

  - Go
    - [Go Code Review Comments](https://go.dev/wiki/CodeReviewComments)
    - [Effective Go](https://golang.org/doc/effective_go.html)
    - Know and avoid [Go landmines](https://gist.github.com/lavalamp/4bd23295a9f32706a48f)
    - Comment your code.
      - [Go's commenting conventions](https://go.dev/doc/comment)
      - If reviewers ask questions about why the code is the way it is, that's a sign that comments might be helpful.
    - Command-line flags should use dashes, not underscores
    - Naming
      - Please consider package name when selecting an interface name, and avoid redundancy. For example, `storage.Interface` is better than `storage.StorageInterface`.
      - Do not use uppercase characters, underscores, or dashes in package names.
      - Please consider parent directory name when choosing a package name. For example, `pkg/controllers/autoscaler/foo.go` should say `package autoscaler` not `package autoscalercontroller`.
          - Unless there's a good reason, the `package foo` line should match the name of the directory in which the `.go` file exists.
          - Importers can use a different name if they need to disambiguate.
      - Locks should be called `lock` and should never be embedded (always `lock sync.Mutex`). When multiple locks are present, give each lock a distinct name following Go conventions: `stateLock`, `mapLock` etc.
    - [API changes](/contributors/devel/sig-architecture/api_changes.md)
    - [API conventions](/contributors/devel/sig-architecture/api-conventions.md)
    - [Kubectl conventions](/contributors/devel/sig-cli/kubectl-conventions.md)
    - [Logging conventions](/contributors/devel/sig-instrumentation/logging.md)

## Testing conventions

  - All new packages and most new significant functionality must come with unit tests.
  - Table-driven tests are preferred for testing multiple scenarios/inputs. For an example, see [TestNamespaceAuthorization](https://github.com/kubernetes/kubernetes/blob/4b8e819355d791d96b7e9d9efe4cbafae2311c88/test/integration/auth/auth_test.go#L1201).
  - Significant features should come with integration (test/integration) and/or [end-to-end (test/e2e) tests](/contributors/devel/sig-testing/e2e-tests.md).
    - Including new `kubectl` commands and major features of existing commands.
  - Unit tests must pass on macOS and Windows platforms - if you use Linux specific features, your test case must either be skipped on windows or compiled out (skipped is better when running Linux specific commands, compiled out is required when your code does not compile on Windows).
  - Avoid relying on Docker Hub. Use the [Google Cloud Artifact Registry](https://cloud.google.com/artifact-registry/) instead.
  - Do not expect an asynchronous thing to happen immediately---do not wait for one second and expect a pod to be running. Wait and retry instead.
  - See the [testing guide](/contributors/devel/sig-testing/testing.md) for additional testing advice.

## Directory and file conventions

  - Avoid package sprawl. Find an appropriate subdirectory for new packages. [See issue #4851](http://issues.k8s.io/4851) for discussion.
    - Libraries with no appropriate home belong in new package subdirectories of `pkg/util`.
  - Avoid general utility packages. Packages called "util" are suspect. Instead, derive a name that describes your desired function. For example, the utility functions dealing with waiting for operations are in the `wait` package and include functionality like `Poll`. The full name is `wait.Poll`.
  - All filenames should be lowercase.
  - Go source files and directories use underscores, not dashes.
    - Package directories should generally avoid using separators as much as possible. When package names are multiple words, they usually should be in nested subdirectories.
  - Document directories and filenames should use dashes rather than underscores.
  - Examples should also illustrate [best practices for configuration and using the system](https://kubernetes.io/docs/concepts/configuration/overview/).
  - Follow these conventions for third-party code:
    - Go code for normal third-party dependencies is managed using [go modules](https://go.dev/wiki/Modules) and is described in the kubernetes [vendoring guide](/contributors/devel/sig-architecture/vendor.md).
    - Other third-party code belongs in `third_party`.
      - forked third party Go code goes in `third_party/forked`.
      - forked _golang stdlib_ code goes in `third_party/forked/golang`.
    - Third-party code must include licenses. This includes modified third-party code and excerpts, as well.
