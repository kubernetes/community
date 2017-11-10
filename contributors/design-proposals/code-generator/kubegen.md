# Kubernetes Code Generators

## Background

Original issue:

- https://github.com/kubernetes/kubernetes/issues/53524

Existing bash script for running some code generators:

- https://github.com/kubernetes/kubernetes/pull/52186

Kubernetes provides a collection of code generators to assist with developing APIs.  These code generators generate
useful libraries such as clients, openapi definitions, copying go structs, etc.

The `kubegen` uses directory structure conventions to detect Kubernetes APIs definitions and run the standard
set of Kubernetes code generators.

There are 2 methods for running kubegen

- Directly through the kubegen command line by downloading the binary and running it from the project root
- Through Bazel by adding rules to `WORKSPACE` and `BUILD.gazel`

kubegen will run the set of Kubernetes code generators for which the project follows
the [Kubernetes API and directory structure conventions](#kubernetes-api-and-directory-structure-conventions).
Some conventions may overridden through specifying specific flags.

## Running via Command line

### Installing the command line

- Download the latest kubegen release from the code-generators release page.
- Extract the tar and add the `kubegen` binary to the PATH
- Run `kubegen version` to display the version
  - Will correspond to a Kubernetes release - e.g. 1.8, 1.9

## Generation for Versioned vs Internal types

Code generation comes in 2 modes:

- Generating clients and libraries for versioned APIs used by controllers, clis, and other clients.
- Generating infrastructure for Kubernetes apiserver machinery to implement and expose API types internally

If you are using CRDs for type definitions, then you should ignore internal types, and use the
*versioned* commands.

If *are* building an API directly into a Kubernetes API server or
Kubernetes extension apiserver, then
you will need to use the *internal* commands.  You may also want to use the *versioned*
commands to publish libraries for your types externally - e.g. `k8s.io/client-go`

### Running the command line

#### Generate clients and related libraries for versioned (external) types

```sh
kubegen versioned
```

- run code generators against discovered API type definitions
- prepend LICENSE to generated files (wrapping in comments)

#### Generate clients and related libraries for internal types

**Note:** This will *also* generate libraries for versioned types as internal types require
versioned types to be useful.

```sh
kubegen internal
```

- run code generators against discovered API type definitions
- prepend LICENSE to generated files (wrapping in comments)

### Options

A number of options can be specified to override the default behavior

#### Define copyright owners

The copyright owners for the license can be defined with the `--copyright` flag.
This will prepend the license header at the top of generated files with a copyright string.
e.g. *Copyright 2018 The Kubernetes Authors.*


```sh
kubegen versioned --copyright "The Kubernetes Authors."
```

#### Define the license file

The license file can be overridden from the LICENSE file using the `--license-file` flag.
This will use the contents from the boilerplate.go.txt file instead of LICENSE.
If boilerplate.go.txt is already in comments then use it verbatim, otherwise wrap it in comments.


**Note**: This passes the `--go-header-file` to each of the code generators

```sh
kubegen versioned --license-file "boilerplate.go.txt"
```

#### Define the license type

If no license file exists, one may be specified on the using the `--license` flag.

```sh
kubegen versioned --license "Apache 2.0"
```

Options:

- "Apache 2.0"
- "None"

#### Define API groups and versions

The set of APIs code generation is performed for can be explicitly defined instead of
implicitly determined using the `--api-version` and `--api-group` flags.  When these
flags are provided, API groups will not be implicitly discovered.  Only API versions
for groups specified with `--api-group` will be discovered.

- `--api-version`: generate for this API version
- `--api-group`: generate for all API versions in this group

```sh
kubegen versioned --api-version apps/v1 --api-version apps/v1beta1 --api-version extensions/v1beta1
```

#### Manually specifying the output for the versioned client packages

By default versioned client related libraries are output to the following locations:

- client: `pkg/client/clientset_generated/clientset`
- informers: `pkg/client/informers_generated/externalversions`
- listers: `pkg/client/listers_generated`

These locations can be changed using the flags `--client-output`, ``--informer-output` and `--listers-output`.
This is useful if generating a client for external consumption.

**Note**: These flags also work with `kubegen internal` since that runs generators on versioned types.

```sh
kubegen versioned --client-output kubernetes --informers-output informers --listers-output listers
```

#### Running in dry-run

To verify the output without writing actual files use the `--dry-run` flag

```sh
kubegen versioned --dry-run
```

**Note**: This passes the `--verify-only` flag to each of the code generators

  
#### Overriding the searched the api directories

The `--apis-dir` defaults to `pkg/apis` and looks for API groups in that directory.

Looking for APIs outside the default location can be configured using
the `--apis-dir` flag.  This flag may be provided multiple times to search multiple
directories.

```sh
kubegen versioned --apis-dir notpkg/apis --apis-dir pkg/notapis
```

## Running via Bazel

### Installing via Bazel

Add the following to the project *WORKSPACE* file

```py
http_archive(
    name = "io_k8s_rules_go",
    url = "https://github.com/kubernetes/bazelbuild/releases/download/v1.8.0/rules_go-1.8.0.tar.gz",
    sha256 = "<number>",
)
```

Add the following to the project root *BUILD.bazel* file

```py
load("@io_k8s_rules_go//go:def.bzl", "code-generator")

kubegen(
    name = "kubegen",
    cmd = "versioned",
)
```

### Running the Bazel target

```sh
bazel run //:kubegen
```

### Bazel options

- *license*: same behavior as the flag
- *license-file*: same behavior as the flag
- *copyright*: same behavior as the flag
- *cmd*: controls sub command - either internal or versioned
- *client-output*: same behavior as the flag
- *informers-output*: same behavior as the flag
- *listers-output*: same behavior as the flag


## Kubernetes API and directory structure conventions

kubegen will run the full set of Kubernetes code generators for all APIs it finds in the project.
kubegen looks for API groups directories defined under `pkg/apis/<group>/<version>` where *group* and *version*
match the patterns following Kubernetes API conventions.

- **group-pattern**: ^[a-z0-9\.]+$
- **version-pattern**: ^v\\d+(alpha\\d+|beta\\d+)?$

By default, kubegen will run code generators for both versioned types defined under `pkg/apis/<group>/<version>` and
internal types defined under `pkg/apis/<group>`.  The location kubegen searches can be overridden.

### Specifying copyright and license headers

kubegen will prepend all generated files with license comments using the LICENSE file at the project root.

### Code generator list

The kubegen binary runs the following set of code generators.

#### Versioned API code generators

**Generators**:

- client-gen
- deepcopy-gen
- defaulter-gen
- informer-gen
- lister-gen

Requires one of:

- presence of files matching `<api-dir>/<group-pattern>/<version-pattern>/*types.go`
- presence of `--api-group` flags
- presence of `--api-version` flags

**Generators**:

- go-to-protobuf

Requires:

- Proto tags to be defined on versioned types

#### Internal API code generators

**Generators**:

- *conversion-gen*
- client-gen
- deepcopy-gen
- defaulter-gen
- informer-gen
- lister-gen

Requires one of:

- presence of files matching `<api-dir>/<group-pattern>/<version-pattern>/*types.go`
- presence of `--api-group` flags
- presence of `--api-version` flags

**Generators**:

- go-to-protobuf

Requires:

- Proto tags to be defined on versioned types

**Generators**:

- openapi-gen (runs on versioned types only)

Support defined by:

- presence of `pkg/openapi/doc.go`

## FAQ

> How does kubegen invoke the other code generators without me having to download them?

kubegen statically compiles the logic for all code generators

> How can Kubernetes APIs be provided as inputs to the code generators for things like generating openapi?

kubegen will include any vendored API groups found under `vendor/k8s.io/api/`.  Additional inputs can
be provided with `--vendor-api-dir`.

> How can Kubernetes APIs be generated when the versioned and internal packages live in different repos?

This works out of the box based on the discovered directories and structure.

> Why are the kubegen flags different than the flags passed to the code generators e.g. --license-file vs --boilerplate-header-file?

kubegen tries to present the simplest possible interface for using the code generators.  since
most GitHub projects already have a LICENSE file created that should be used in most cases, this flag more
accurately reflects what the value should be.  likewise `--dry-run` is a more common name than `--verify-only`.

> I want to specify a code generator flag that isn't exposed in kubegen, how do I do this?

Please file an issue if your use case is not supported.

> How can I use a different pattern for groups or versions?

We could expose this as a feature if we discover it is necessary by adding the
`--group-pattern` and `--version-pattern` flags to override the regular expressions.

Please file an issue if your use case is not supported.

> How can I specify only certain code generators are run

Currently all code generators are run automatically, and they generate
output based off the struct comments present on the API
type definitions in the go files.  e.g. Specifying `//+genclient` before the struct definition.

Please file an issue if your use case is not supported.

> What happens if I don't have a LICENSE file

If you don't have a LICENSE file, you must either provide the `--license` flag or the
`--license-file` flag.  If you do neither, kubegen will exit 2.

> My LICENSE file isn't in go comments, so how to you prepend it to go files?

We make sure it is wrapped it in comments when it is prepended.

> Will the Bazel rules selectively regenerate files based on what is changed instead
  of always generating everything?

Not in the mvp.  This would be a great feature to add.