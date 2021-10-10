# Verification Tests

**Table of Contents**

- [Verification Tests](#verification-tests)
  - [Overview](#overview)
  - [Note](#note)
  - [verify-govet-levee](#verify-govet-leve)
  - [verify-api-groups](#verify-api-groups)

## Overview

Verification tests for Kubernetes provide a mechanism to verify contributions for adherence to project conventions
and best practices, and to validate generated build artifacts for soundness.

All blocking verification tests can be executed via `make verify`.
Individual verification tests also can be found in vestigial shell scripts at `hack/verify-*.sh`.

Most verification tests are self-explanatory.
`verify-govet`, for instance, performs `go vet` checks, which [defends against common mistakes](https://golang.org/cmd/vet/).
The verification tests fails when `go vet` produces any findings.
More complex verification tests are described below.

### Note

This documentation is a work in progress.  This listing is incomplete.

### `verify-govet-levee`

Verification in `verify-govet-levee.sh` uses taint propagation analysis
to defend against accidental logging of credentials.
Struct fields which may contain credentials should be annotated as such using the `datapolicy` field tag.
Field tagging was introduced by [KEP-1753](https://github.com/kubernetes/enhancements/issues/1753), and analysis was introduced by [KEP-1993](https://github.com/kubernetes/enhancements/issues/1933).
Additional credential sources may be identified in analysis configuration (see below).

Taint propagation analysis defends against both direct and indirect logging of credentials.
Consider the following hypothetical snippet.

```golang
// kubernetes/cmd/kubelet/app/server.go

// kubeConfigSpec struct holds info required to build a KubeConfig object
type kubeConfigSpec struct {
	CACert         *x509.Certificate
	APIServer      string
	ClientName     string
	TokenAuth      *tokenAuth      `datapolicy:"token"`
	ClientCertAuth *clientCertAuth `datapolicy:"security-key"`
}

func MyDangerousFunction(spec kubeConfigSpec) error {
	if spec.CACert == nil {
		err := fmt.Errorf("kubeConfigSpec missing expected CACert, got %#v", spec)  // Dangerous logging!
		klog.Error(err)
		return err
	}
	
	if err := DoSomethingElse(spec); err != nil {
		klog.Error(err)  // Dangerous logging!
	}

	return nil
}
```

In the above, when `spec.CACert == nil`, we log the `spec`.
However, we know from the datapolicy field tags that the spec could contain one or more credentials and should not be logged.
The analysis will detect this and cause the verification test to fail.
The log call should be adjusted to extract exactly what information is relevant from the `spec`.

The second `klog.Error` call is also problematic.
The error returned by `DoSomethingElse` could potentially encapsulate the credential passed in by `spec`, and so we must not log it.
That is, we consider `err` to be "tainted" by the call which has access to the credentials.
The analysis will detect this as well and call the verification test to fail.

When this analysis causes the verification test to fail, a developer has several options.
In order of decreasing preference:
* Reconstruct logging calls such that only non-secret information is passed.
* Reconstruct a method which caused taint to spread to return indicators which are not logged directly, e.g. return `value, ok` rather than `value, err`.
* Write a *sanitizer* whose return value is guaranteed to be log-safe.  Add this sanitizer to the analysis configuration (see below).
* Add the method where the log call occurs to the analysis configuration exclude-list.

Analysis configuration can be found at [kubernetes/kubernetes/hack/testdata/levee/levee-config.yaml](https://github.com/kubernetes/kubernetes/blob/master/hack/testdata/levee/levee-config.yaml).
Contact SIG-Security with any additional questions.

### `verify-api-groups`

This verification script validates the different api-groups by reading
the respective `register.go` file. Every register file must contain a
GroupName.  Another check which is performed when this script runs is
to ensure that all types have client code generated for them, except
types that belong to groups not served from the API server (defined in
this script via the bash array `groups_without_codegen`).

Next, the script compares the `GroupName`s against
`import_known_versions` to ensure the import packages will get
installed. We list out packages which are required without
installation along with importing `known_version`. Then we do a search
for packages that reqiure installation on the basis of
`packages_without_installation`. We verify if file is a
`known_version_file` or not only if an `expected_install_package` is
present in it.

Finally the script checks that all external group versions
(e.g. `foobar/v1`) are defined in `hack/lib/init.sh` in either the
`KUBE_AVAILABLE_GROUP_VERSIONS` or `KUBE_NONSERVER_GROUP_VERSIONS`
bash variables.

### `verify-bazel`

This verify-bazel script validates the removal of bazel related
files. The script ensures no bazel related temporary, intermediate or
output files remain as part of
[KEP-2420](https://github.com/kubernetes/enhancements/issues/2420).

## `verify-boilerplate`

This script checks for the license headers for all the files, whether
the header is correct or wrong. The purpose of boilerplate headers is
to identify license headers. The script collects all the file names
generated by `hack/boilerplate/boilerplate.py` script and stores them
into a list called `files_need_boilerplate`.

Once we collect all the file names, we run a check to identify the
files with wrong header. This check will only run if the file exists
in the list mentioned above.


## `verify-cli-conventions`

This script checks whether the description format of `help` message of
kubectl command is valid or not. This check is done for all the
`kubectl` sub-commands as well. The script first checks whether the go
command is available or not in the ${PATH}. And then the binary
`cmd/clicheck` is checked if it exists or not in the well-known output
locations. It runs command checks on all the kubectl commands or subcommands like,

```bash
kubectl version
kubectl uncordon
kubectl wait
kubectl top node
```

And if the output looks good i.e. CLI follows all tested conventions
the test passes.


## `verify-codegen`

This script verifies if the code update is needed or not for specific
sub-projects. It first verifies the correct Go version and creates Go
path. The script checks for the updated code for below subprojects,

```bash
k8s.io/code-generator
k8s.io/kube-aggregator
k8s.io/sample-apiserver
k8s.io/sample-controller
k8s.io/apiextensions-apiserver
k8s.io/metrics
```

Once it completes checking for code updates, later the script calls
`update-codegen.sh` scripts.


## `verify-structured-logging`

This script verifies if a package is properly migrated to structured
logging or not. The script involves verification steps based on new
klog methods which have few disallowed keywords like,

```bash
* klog.Infof, klog.Info, klog.Infoln
* klog.InfoDepth
* klog.WarningDepth
* klog.Error, klog.Errorf, klog.Errorln
* klog.ErrorDepth
```

More info is available
[here](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-instrumentation/migration-to-structured-logging.md#change-log-functions-to-structured-equivalent).


## `verify-gofmt`

This script is used to check whether the go source code needs to be
formatted or not using, the gofmt tool. Gofmt tool automatically
formats the code and the formatted code is easier to read, write and
maintain.


## `verify-spelling`

This script uses `client9/misspell` package to search and correct
commonly misspelled words as per the English language in all the files
and directories under `kubernetes/kubernetes`.


