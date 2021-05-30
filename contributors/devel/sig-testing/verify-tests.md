# Verification Tests

**Table of Contents**

- [Verification Tests](#verification-tests)
  - [Overview](#overview)
  - [Note](#note)
  - [`verify-govet-levee`](#verify-govet-leve)

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
