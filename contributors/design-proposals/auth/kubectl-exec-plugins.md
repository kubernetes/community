# Out-of-tree client authentication providers

Author: @ericchiang

# Objective

This document describes a credential rotation strategy for client-go using an exec-based
plugin mechanism.

# Motivation

Kubernetes clients can provide three kinds of credentials: bearer tokens, TLS
client certs, and basic authentication username and password. Kubeconfigs can either
in-line the credential, load credentials from a file, or can use an `AuthProvider`
to actively fetch and rotate credentials. `AuthProviders` are compiled into client-go
and target specific providers (GCP, Keystone, Azure AD) or implement a specification
supported but a subset of vendors (OpenID Connect).

Long term, it's not practical to maintain custom code in kubectl for every provider. This
is in-line with other efforts around kubernetes/kubernetes to move integration with cloud
provider, or other non-standards-based systems, out of core in favor of extension points.

Credential rotation tools have to be called on a regular basis in case the current
credentials have expired, making [kubectl plugins](https://kubernetes.io/docs/tasks/extend-kubectl/kubectl-plugins/),
kubectl's current extension point, unsuitable for credential rotation. It's easier
to wrap `kubectl` so the tool is invoked on every command. For example, the following
is a [real example](
https://github.com/heptio/authenticator#4-set-up-kubectl-to-use-heptio-authenticator-for-aws-tokens)
from Heptio's AWS authenticator:

```terminal
kubectl --kubeconfig /path/to/kubeconfig --token "$(heptio-authenticator-aws token -i CLUSTER_ID -r ROLE_ARN)" [...]
```

Beside resulting in a long command, this potentially encourages distributions to
wrap or fork kubectl, changing the way that users interact with different
Kubernetes clusters.

# Proposal

This proposal builds off of earlier requests to [support exec-based plugins](
https://github.com/kubernetes/kubernetes/issues/35530#issuecomment-256170024), and
proposes that we should add this as a first-class feature of kubectl. Specifically,
client-go should be able to receive credentials by executing a command and reading
that command's stdout.

In fact, client-go already does this today. The GCP plugin can already be configured
to [call a command](
https://github.com/kubernetes/client-go/blob/kubernetes-1.8.5/plugin/pkg/client/auth/gcp/gcp.go#L228-L240)
other than `gcloud`.

## Plugin responsibilities

Plugins are exec'd through client-go and print credentials to stdout. Errors are
surfaced through stderr and a non-zero exit code. client-go will use structured APIs
to pass information to the plugin, and receive credentials from it.

```go
// ExecCredentials are credentials returned by the plugin.
type ExecCredentials struct {
    metav1.TypeMeta `json:",inline"`

    // Token is a bearer token used by the client for request authentication.
    Token string `json:"token,omitempty"`
    // Expiry indicates a unix time when the provided credentials expire.
    Expiry int64 `json:"expiry,omitempty"`
}

// Response defines metadata about a failed request, including HTTP status code and
// response headers.
type Response struct {
    // HTTP header returned by the server.
    Header map[string][]string `json:"header,omitempty"`
    // HTTP status code returned by the server.
    Code int32 `json:"code,omitempty"`
}

// ExecInfo is structed information passed to the plugin.
type ExecInfo struct {
    metav1.TypeMeta `json:",inline"`

    // Response is populated when the transport encounters HTTP status codes, such as 401,
    // suggesting previous credentials were invalid.
    // +optional
    Response *Response `json:"response,omitempty"`

    // Interactive is true when the transport detects the command is being called from an
    // interactive prompt.
    Interactive bool `json:"interactive,omitempty"`
}
```

To instruct client-go to use the bearer token `BEARER_TOKEN`, a plugin would print:

```terminal
$ ./kubectl-example-auth-plugin
{
  "kind": "ExecCredentials",
  "apiVersion":"client.authentication.k8s.io/v1alpha1",
  "token":"BEARER_TOKEN"
}
```

To surface runtime-based information to the plugin, such as a request body for request
signing, client-go will set the environment variable `KUBERNETES_EXEC_INFO` to a JSON
serialized Kubernetes object when calling the plugin.


```terminal
KUBERNETES_EXEC_INFO='{
    "kind":"ExecInfo",
    "apiVersion":"client.authentication.k8s.io/v1alpha1",
    "response": {
        "code": 401,
        "header": {
            "WWW-Authenticate": ["Bearer realm=\"Access to the staging site\""]
        }
    },
    "interactive": true
}'
```

### Caching

kubectl repeatedly [re-initializes transports](https://github.com/kubernetes/kubernetes/issues/37876)
while client-go transports are long lived over many requests. As a result naive auth
provider implementations that re-request credentials on every request have historically
been slow.

Plugins will be called on client-go initialization, and again when the API server returns
a 401 HTTP status code indicating expired credentials. Plugins can indicate their credentials
explicit expiry using the `Expiry` field on the returned `ExecCredentials` object, otherwise
credentials will be cached throughout the lifetime of a program.

## Kubeconfig changes

The current `AuthProviderConfig` uses `map[string]string` for configuration, which
makes it hard to express things like a list of arguments or list key/value environment
variables. As such, `AuthInfo` should add another field which expresses the `exec`
config. This has the benefit of a more natural structure, but the trade-off of not being
compatible with the existing `kubectl config set-credentials` implementation.

```go
// AuthInfo contains information that describes identity information.  This is use to tell the kubernetes cluster who you are.
type AuthInfo struct {
    // Existing fields ...

    // Exec is a command to execute which returns credentials to the transport to use.
    // +optional
    Exec *ExecAuthProviderConfig `json:"exec,omitempty"`

    // ...
}

type ExecAuthProviderConfig struct {
    Command string   `json:"command"`
    Args    []string `json:"args"`
    // Env defines additional environment variables to expose to the process. These
    // are unioned with the host's environment, as well as variables client-go uses
    // to pass argument to the plugin.
    Env []ExecEnvVar `json:"env"`

    // Preferred input version of the ExecInfo. The returned ExecCredentials MUST use
    // the same encoding version as the input.
    APIVersion string `json:"apiVersion,omitempty"`

    // TODO: JSONPath options for filtering output.
}

type ExecEnvVar struct {
    Name  string `json:"name"`
    Value string `json:"value"`

    // TODO: Load env vars from files or from other envs?
}
```

This would allow a user block of a kubeconfig to declare the following:

```yaml
users:
- name: mmosley
  user:
    exec:
      apiVersion: "client.authentication.k8s.io/v1alpha1"
      command: /bin/kubectl-login
      args: ["hello", "world"]
```

The AWS authenticator, modified to return structured output, would become:

```yaml
users:
- name: kubernetes-admin
  user:
    exec:
      apiVersion: "client.authentication.k8s.io/v1alpha1"
      command: heptio-authenticator-aws
      # CLUSTER_ID and ROLE_ARN should be replaced with actual desired values.
      args: ["token", "-i", "(CLUSTER_ID)", "-r", "(ROLE_ARN)"]
```

## TLS client certificate support

TLS client certificate support is orthogonal to bearer tokens, but something that
we should consider supporting in the future. Beyond requiring different command
output, it also requires changes to the client-go `AuthProvider` interface.

The current The auth provider interface doesn't let the user modify the dialer,
only wrap the transport.

```go
type AuthProvider interface {
    // WrapTransport allows the plugin to create a modified RoundTripper that
    // attaches authorization headers (or other info) to requests.
    WrapTransport(http.RoundTripper) http.RoundTripper
    // Login allows the plugin to initialize its configuration. It must not
    // require direct user interaction.
    Login() error
}
```

Since this doesn't let a `AuthProvider` supply things like client certificates,
the signature of the `AuthProvider` should change too ([with corresponding changes
to `k8s.io/client-go-transport`](
https://gist.github.com/ericchiang/7f5804403b359ebdf79dcf76c4071bff)):

```go
import (
    "k8s.io/client-go/transport"
    // ...
)

type AuthProvider interface {
    // UpdateTransportConfig updates a config by adding a transport wrapper,
    // setting a bearer token (should ignore if one is already set), or adding
    // TLS client certificate credentials.
    //
    // This is called once on transport initialization. Providers that need to
    // rotate credentials should use Config.WrapTransport to dynamically update
    // credentials.
    UpdateTransportConfig(c *transport.Config)

    // Login() dropped, it was never used.
}
```

This would let auth transports supply TLS credentials, as well as instrument
transports with in-memory rotation code like the utilities implemented by
[`k8s.io/client-go/util/certificate`](https://godoc.org/k8s.io/client-go/util/certificate).

The `ExecCredentials` would then expand to provide TLS options.

```go
type ExecCredentials struct {
    metav1.TypeMeta `json:",inline"`

    // Token is a bearer token used by the client for request authentication.
    Token string `json:"token,omitempty"`
    // PEM encoded client certificate and key.
    ClientCertificateData string `json:"clientCertificateData,omitempty"`
    ClientKeyData         string `json:"clientKeyData,omitempty"`

    // Expiry indicates a unix time when the provided credentials expire.
    Expiry int64 `json:"expiry,omitempty"`
}
```

The `AuthProvider` then adds those credentials to the `transport.Config`.

## Login

Historically, `AuthProviders` have had a `Login()` method with the hope that it
could trigger bootstrapping into the cluster. While no providers implement this
method, the Azure `AuthProvider` can already prompt an [interactive auth flow](
https://github.com/kubernetes/client-go/blob/kubernetes-1.8.5/plugin/pkg/client/auth/azure/azure.go#L343).
This suggests that an exec'd tool should be able to trigger its own custom logins,
either by opening a browser, or performing a text based prompt.

We should take care that interactive stderr and stdin are correctly inherited by
the sub-process to enable this kind of interaction. The plugin will still be
responsible for prompting the user, receiving user feedback, and timeouts.
