# Kubernetes Code Generators

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
- Run `kubegen version` to list the version
  - Will correspond to a Kubernetes release

### Running the command line

#### Use defaults

```sh
kubegen
```

- run code generators against discovered API type definitions
- prepend LICENSE to generated files (wrapping in comments)

#### Specifying copyright owners

The copyright owners for the license can be defined with the `--copyright` flag

```sh
kubegen --copyright "The Kubernetes Authors."
```

- run code generators against discovered API type definitions
- prepend the LICENSE file to generated files
  - wrap the contents of the LICENSE in comments
  - start license with "Copyright <current year> <copyright>" e.g. *Copyright 2018 The Kubernetes Authors.*
  - exit non-zero if the LICENSE file is missing


#### Specifying the license file

The license file can be overridden from the LICENSE file using the `--license-file` flag

```sh
kubegen --license-file "boilerplate.go.txt"
```

- run code generators against discovered API type definitions
- prepend license to generated files
  - use the contents from the boilerplate.go.txt file instead of LICENSE
  - if boilerplate.go.txt is already in comments then use it verbatim, otherwise wrap it in comments
  - exit non-zero if boilerplate.go.txt is missing
  
**Note**: This passes the `--go-header-file` to each of the code generators

#### Specifying a license type

The license can be overridden from the LICENSE file using the `--license` flag

```sh
kubegen --license "Apache 2.0"
```

- run code generators against discovered API type definitions
- prepend license to generated files
  - use an *Apache 2.0* license

#### Specifying API versions

Generating code only for specific API group versions can be performed by providing the
`--api-version` flag.  When manually specifying the API group versions,
neither internal nor external API types are automatically discovered.

```sh
kubegen --api-version apps/v1 --api-version apps/v1beta1 --api-version extensions/v1beta1
```

- run code generators against the apps/v1 apps/v1beta1 extensions/v1beta1 API groups versions
  - only run against external types, do not run against internal types
- prepend license to generated files

#### Specifying API groups

Generating code only for specific API groups can be performed by providing the
`--api-group` flag.  This will discover API versions only under the specified
groups.  When manually specifying the API groups,
neither internal nor external API types are automatically discovered.

```sh
kubegen --api-group apps --api-group extensions
```

- run all code generators against the apps extensions API groups
  - only run against external types, do not run against internal types
- prepend license to generated files

#### Specifying internal API groups

To generate code for both manually specified internal types use the
`--internal-api-group` flag.

```sh
kubegen --internal-api-group extensions --api-group extensions
```

- run code generators against the provided API groups
  - generate internal types for extensions groups
  - generate external types for versions under the extensions group
- prepend license to generated files

#### Running in dry-run

To verify the output without writing actual files use the `--dry-run` flag

```sh
kubegen --dry-run
```

**Note**: This passes the `--verify-only` flag to each of the code generators

  
#### Overriding the searched the api directories

The `--apis-dir` defaults to `pkg/apis` and looks for API groups in that directory.

Looking for APIs outside the default location can be configured using
the `--apis-dir` flag.  This flag may be provided multiple times to search multiple
directories.

```sh
kubegen --apis-dir notpkg/apis --apis-dir pkg/notapis
```

- run code generators against discovered APIs
  - search for API group versions under notpkg/apis and pkg/notapis instead of pkg/apis  

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
)
```

### Running the Bazel target

```sh
bazel run //:kubegen
```

### Bazel options

- *license*: same behavior as *license* flag for cli
- *license-file*: same behavior as *license-file* flag for cli
- *copyright*: same behavior as *copyright* flag for cli

## Kubernetes API and directory structure conventions

kubegen will run the full set of Kubernetes code generators for all APIs it finds in the project.
kubegen looks for API groups directories defined under `pkg/apis/<group>/<version>` where *group* and *version*
match the patterns following Kubernetes API conventions.

- **group-pattern**: ^[a-z0-9\.]+$
- **version-pattern**: ^v\\d+(alpha\\d+|beta\\d+)?$

By default, kubegen will run code generators for both external types defined under `pkg/apis/<group>/<version>` and
internal types defined under `pkg/apis/<group>`.  The location kubegen searches can be overridden.

### Generation for External vs Internal types

If you *are not* writing a Kubernetes API server or Kubernetes extension apiserver, and
instead using CRDs for type definitions, then you can ignore internal types.

If *are* writing a Kubernetes API server or Kubernetes extension apiserver, then
you will need to generate code for internal types.

### Specifying copyright and license headers

kubegen will prepend all generated files with license comments using the LICENSE file at the project root.

### Code generator list

By default, the kubegen binary runs the set of code generators for which the project
supports.  Support for each code generator is as follows...

#### External API code generators

**Generators**:

- client-gen
- deepcopy-gen
- defaulter-gen
- informer-gen
- lister-gen

Support defined by union of (||):

- presence of files matching `<api-dir>/<group-pattern>/<version-pattern>/*types.go`
- presence of `--api-group` flags
- presence of `--api-version` flags

**Generators**:

- go-to-protobuf

Support defined by:

- TODO(sttts): What is the answer?

**Generators**:

- openapi-gen (external only)

Support defined by:

- presence of `pkg/openapi/doc.go`

#### Internal API code generators

**Generators**:

- client-gen
- deepcopy-gen
- defaulter-gen
- informer-gen
- lister-gen

Support defined by union of (||):

- presence of files matching `<api-dir>/<group-pattern>/*types.go`
- presence of `--internal-api-group` flags

**Generators**:

- go-to-protobuf

Support defined by:

- TODO(sttts): What is the answer?

#### Internal+External API code generators

**Generators**:

- conversion-gen

Support defined by intersection of (&&):

- Internal code generators
- External code generators

## FAQ

> I want to only generate *external* or only generate *internal* types.  How can I do this?

You can restrict the generated APIs to a manually specified list with `--internal-api-group` and `--api-group`

> How does kubegen invoke the other code generators without me having to download them?

kubegen statically compiles the logic for all code generators

> How can Kubernetes APIs be provided as inputs to the code generators for things like generating openapi?

kubegen will include any vendored API groups found under `vendor/k8s.io/api/`.  Additional inputs can
be provided with `--vendor-api-dir`.

> How can Kubernetes APIs be generated when the external and internal packages live in different repos?

This works out of the box based on the discovered *types.go files.

> Why are the kubegen flags different than the flags passed to the code generators e.g. --license-file vs --boilerplate-header-file?

kubegen tries to present the simplest possible interface for using the code generators.  since
most GitHub projects already have a LICENSE file created that should be used, this flag more
accurately reflects what its value should be.  likewise `--dry-run` is a more common name than `--verify-only`.

> I want to specify a code generator flag that isn't exposed in kubegen, how do I do this?

kubegen tries to identify the correct value for most flag.  If there is a flag you want
exposed, file and issue with the context, and we will work fix it either by changing how its value
is automatically set, or by exposing it.

Please file an issue if your use case is not supported.

> How can I use a different pattern for groups or versions?

We can expose this as a feature if we discover it is necessary by adding the
`--group-pattern` and `--version-pattern` flags to override the regular expressions.

Please file an issue if your use case is not supported.

> How can I specify only certain code generators are run

Currently all code generators are run automatically based off the struct comments present on the API
type definitions in the go files.  e.g. Specifying `//+genclient` before the struct definition.

Please file an issue if your use case is not supported.+

> What happens if I don't have a LICENSE file

If you don't have a LICENSE file, you must either provide the `--license` flag or the
`--license-file` flag.  If you do neither, kubegen will exit 2.

> My LICENSE file isn't in go comments, so how to you prepend it to go files?

We wrap it in comments if it is not already.