# Self-hosting authorization webhooks

Status: Pending

Version: **Alpha** | Beta | GA

Implementation Owner: @filmil

## Motivation

Make it possible to have the authorization (hereafter authz for short) webhook
run as a pod in the cluster being authorized. This addresses part of the
concerns raised in [kubernetes/kubernetes#52511][18].

See:
  - Feature request: [kubernetes/features#516][15]
  - Earlier pull requests:
    - [kubernetes/kubernetes#54733][16]
    - [kubernetes/kubernetes#54163][17]


This requires a change in the way the `kube-apiserver` webhook configuration is
specified (a new configuration format), and a special initialization of the
webhook module in `kube-apiserver` (in contrast to any other apiservers out
there) as [outlined in an earlier discussion on this topic][9].

As of this writing, the authorization webhook in Kubernetes is configured
through the `kube-apiserver` flag `--authorization-webhook-config-file` which is
a [kubeconfig][3]-formatted [configuration file][2].  An example taken from the
documentation:

```yaml
# clusters refers to the remote service.
clusters:
  - name: name-of-remote-authz-service
    cluster:
      # CA for verifying the remote service.
      certificate-authority: /path/to/ca.pem
      # URL of remote service to query. Must use 'https'. May not include parameters.
      server: https://authz.example.com/authorize

# users refers to the API Server's webhook configuration.
users:
  - name: name-of-api-server
    user:
      client-certificate: /path/to/cert.pem # cert for the webhook plugin to use
      client-key: /path/to/key.pem          # key matching the cert

# kubeconfig files require a context. Provide one for the API Server.
current-context: webhook
contexts:
- context:
    cluster: name-of-remote-authz-service
    user: name-of-api-server
  name: webhook
```

The kubeconfig file format does not easily admit reference to an endpoint that
is hosted within the cluster.  Earlier proposals to extend the syntax of the
`clusters[*].cluster.server` field to admit a custom dialer that calls a service
endpoint [were rejected][4] (also see [another rationale][5]).  The [suggested
viable alternative][6] is to provide a new webhook file format that references
a service explicitly, and build on top of that.

One further concern to address is that this proposed change is only [relevant
for the `kube-apiserver`][7], and not to any of the generic apiservers.  This
means that the proposed changes will be set up such that only the
`kube-apiserver` can make use of it.  This approach has [garnered support][8]
from the community in the earlier proposal.

## Proposal

* Define a new configuration format, supported only by `kube-apiserver` that
  admits a webhook endpoint within the cluster to be specified.

  This new configuration file is still specified using the same flag as for the
  original webhook configuration, `--auhtorization-webhook-config-file`.

* Instantiate a custom subject access review client based on this configuration
  that has special behavior as follows:

  - If the new configuration has been specified, a custom dialer is supplied
    which resolves a service endpoint and uses that service endpoint to dial
    into when making a call.

## User Experience

### Use Cases

1. As a cluster administrator, I would like to install Kubernetes-based products
   that have self-contained deployments in a cluster and include authorization.

   Kubernetes products typically come with an install script, say based on Helm
   or `kubeadm`.  These allow custom deployments, including reconfiguring
   `kube-apiserver`.  When such a product requires an authorizer, the authorizer
   is added as a system component with a pod-based deployment in a cluster and
   the apiserver is instructed to consult the endpoint based on that deployment
   for authorization decisions.

   The cluster administrator experience changes in that they are able to install
   a wider range of Kubernetes-based products that include this feature,
   compared to when this is not the case.

## Security

The authorizer webhook module today is an extension mechanism that wields
considerable power.  A new authorizer has the opportunity to grant access to any
resource and override decisions from any other authorizers.  So we look into
some detail of the implications of a self-hosted authorizer.

A self-hosted authorizer is deployed today by changing changing apiserver flag
`--authorization-webhook-config-file`.  This limits the set of accounts
that can change this setting to those having permission to access the
apiserver to those that can restart the apiserver, or change the underlying
apiserver manifest.

Standard concerns apply around security of the connection between
`kube-apiserver` and the webhook authorizer.  Some of that concert is handled
through two-way TLS negotiation for the webhook connection.  However, we still
need to ensure that the credentials (CA certificate, server and client
certificate, server and client private keys) are not compromised, as well as
that the DNS entries that affect the service resolution are properly set.

One possibility to handle these concerns is to set aside nodes that execute
privileged code under special scrutiny.

## Implementation


### Configuration format

Let us define a new configuration file format.  This is based on the [admission
configuration file format][12] as well as the proposal from @frankfarzan for an
analogous [configuration format for authz][13].

An example configuration file for authz looks like this:

```yaml
apiVersion: kubeapiserverauthorization.config.k8s.io/v1alpha1
kind: ExternalAuthorizationHookConfiguration
metadata:
  name: example-config
externalAuthorizationHooks:
- name: webhook-name
  clientConfig:
    # The Certificate Authority (CA) certificate file path that signed the
    # webhook server certificate.  Used to verify the identity of the webhook
    # server at kube-apiserver side.
    serverCaFile: ...
    # Path to the certificate file for the webhook plugin to present to the
    # webhook server.
    clientCertFile: ...
    # The private key file corresponding to the client certificate file above.
    clientKeyFile: ...
    # Specify only onen of 'url' or 'service'.
    url: <URL of the remote webhook>
    service:
      name: some-service # name of the front-end service
      namespace: some-namespace #<namespace of the front-end service>
```

The go data definition matching the file above is as follows.  This captures
both the service reference and a URL.

```go
// Required well-known import statements elided for brevity.

// Top-level
type ExternalAuthorizationHookConfiguration struct {
  meta.TypeMeta

  ExternalAuthorizationHooks map[string]ExternalAuthorizationHook
}

type ExternalAuthorizationHook struct {

  Name string
  ClientConfig ClientConfig
}

type ClientConfig struct {
  ServerCaFile string
  ClientCertificateFile string
  ClientKeyFile string
  // +optional
  Service ServiceReference
  // https only
  Url string
}

type ServiceReference struct {
  Name string
  Namespace string
}
```

When `service` is omitted, then the `url` alone is used in the same way as in
the old configuration.  If service is defined, `url` is ignored, and the
endpoint obtained as result of service resolution is substituted for the `Host`
in the url, and `Path` part is set to empty.

Example:

```yaml
# Elided irrelevant fields.
externalAuthorizationHooks:
- name: webhook-name
  clientConfig:
    service:
      name: some-service
      namespace: some-namespace
```

This would make a request to `https://(endpointOf name=some-service,
namespace=some-namespace>`


### Parsing the configuration file

This will edit kube-apiserver `config.go` parsing to first try parsing the
authz configuration file as new format.  If the new format is found, new
behavior is invoked.  Otherwise, the old code paths are reused.

While strictly speaking parsing the file twice, looking for different formats is
not as efficient as parsing it once looking for both formats, it has other
benefits:

- It does not require changing the old code path which uses a deep integration
  with the generic webhook client code.  By inspection it seems that unifying
  the two would be a sizable undertaking, and does not seem to make sense from a
  cost/benefit perspective.

- The amortized cost of this extra read is zero, since the read only happens at
  `kube-apiserver` start-up.

### Wiring through extra information

To construct a service-aware dialer for the new code path, we need to wire
through a proxy-aware transport and the service resolver from the
`kube-apiserver`.  This will be done by packaging the two components into an
interface provided by the webhook library as follows, to decouple the
two implementations:

```go
// In server.go, kube-apiserver:
type resolver struct {
  proxyTransport *http.Transport
  resolver ServiceResolver
}
func (r *resolver) NewDialer(namespace, name string) func (net.Conn, error) {
  // ...
}

// In webhook.go:
type DialerFactory interface {
  func NewDialer(namespace, name string) func (net.Conn, error)
}
```

This dialer will be subsituted when creating a REST client.


### Wiring things together and testing

The wire-through and testing will follow the approach already outlined in
the PR kubernetes/kubernetes#54733 ([link][14]).

### Server Backwards/Forwards compatibility

The `kube-apiserver` will continue to support the kubeconfig file format for
configuring authorization for as long as it is required, and will retain the
respective behavior.

The new file format will be versioned to allow a smooth transition between
apiserver versions.

One will, of course, not be able to specify a new configuration file format
to a `kube-apiserver` that was compiled without the new file format support.

None of the above changes seem to be disruptive to the daily use of the
`kube-apiserver` in deployments.  Only the clients that want to use the new
feature need ever know that it exists.

## Alternatives considered

This is a list of rejected alternatives for specific components of the proposal.
Each one is contrasted to its currently accepted solution from the proposal
above.

### Extending kubeconfig format

Versus: a new configuration format.

Feedback from the community has indicated that extending the kubeconfig file
format is not acceptable.  Concerns raised were:

- Service resolution based on the hostname in a URL is not the direction that
  Kubernetes evolves towards.

- Kubeconfig files already have a widely accepted use, and extending them to
  make services first-class citizens is a non-goal.

### Using a new configuration flag

Versus: overloading the configuration flag with a new meaning.

More flags mean a larger configuration space.  It seems beneficial to avoid
introducing new flags if possible.  Reusing a flag also makes it obvious that
there may only be one authz configuration active at a time.

### Using dynamic authorizator configuration

Versus: static authorizator configuration in a file.

[Concerns were raised][11] about the ability of the cluster admin to
misconfigure a cluster, or undo a previously established hosted configuration
if dynamic configuration is used.

This proposal side-steps that concern as it still relies on the use of a
configuration file, which is out of reach of cluster admins in hosted solutions.
At the same time, it does not prevent a future change that would, if required,
admit dynamic configuration.

### Allowing a nonempty request path with service resolution

Versus: always requesting on `https://serviceIP:servicePort/`

While initial version of this proposal contained a way to send a request to
arbitrary path of the resolved service hostport, expressing that ended up being
somewhat akward, requiring technically correct, but uncommon constructions like:

```yaml
# Much elided
# ...
    service:
      name: some-service
      namespace: some-namespace
    url: https:/some/path
```

Suggestions that would allow setting the request path elegantly gladly accepted.

### Using PEM-encoded configuration in the configuration file

Versus: specifying filesystem paths to the certificates.

Specifying the filesystem paths allows the authz configuration file to remain
the same across deployments, while only the mounted contents of the needed files
vary where it applies.

Contrast that to self-registration, where certificates and keys would have to be
passed along as inline content in the API request to register a webhook.

Since we are currently not considering dynamic registration (see previous
subsection), we can accept only having fields that point at filesystem paths.
This decision does not prevent changing the direction in a followup version, say
by allowing filesystem references and inline content to coexist.

### Extending generic apiserver

Versus: extending kube-apiserver only.

Only the `kube-apiserver` has the ability to resolve service endpoints without
special configuration.  Also, `kube-apiserver` is the only apiserver that has
the use case described in this document.

Therefore, it seems reasonable to confine the behavior changes to
`kube-apiserver` alone.

### Unifying code paths for webhooks

Versus: this work is out of scope.

There is a long-standing [intention to unify the code paths][19] of the K8S
webhooks, which have up to now by and large evolved separately.  I think that
this change shouldn't be the place where such an intention should be realized,
reasons below.

- Our proposed solution makes no change to the current state of the authz
  webhook libraries for apiservers in general; but does not prevent any later
  evolution of that state towards a common, unified code base.

- Unifying code paths is a larger effort, requiring likely changes in all the
  webhook code.  Our change is confined to `kube-apiserver` authz webhook.  It's
  not clear that extension apiservers should be allowed to do service resolution
  (and by default, they can't).

- Not all webhooks are currently appropriate for dynamic registration: admission
  webhook can be dynamically registered because any additional such webhook can
  only restrict access today.  However, an additional authz webhook can expand
  access, so it makes sense to require that authz webhooks be installed through
  `kube-apiserver` flags to keep tigher control over what gets installed.
  There's a request to allow authz webhook the ability to override admit
  verdicts from companion authorizers, but has not been implemented yet. A
  decision on that likely gates further work on unification.

[1]: https://github.com/kubernetes/kubernetes/issues/52511
[2]: https://kubernetes.io/docs/admin/authorization/webhook/
[3]: https://kubernetes.io/docs/tasks/access-application-cluster/configure-access-multiple-clusters/
[4]: https://github.com/kubernetes/kubernetes/pull/54889#issuecomment-343045279
[5]: https://github.com/kubernetes/kubernetes/pull/54733#issuecomment-343160937
[6]: https://github.com/kubernetes/kubernetes/pull/54733#issuecomment-343292540
[7]: https://github.com/kubernetes/kubernetes/issues/52511#issuecomment-331489326
[8]: https://github.com/kubernetes/kubernetes/issues/52511#issuecomment-333224542
[9]: https://github.com/kubernetes/kubernetes/issues/52511#issuecomment-333541769
[10]: https://github.com/kubernetes/kubernetes/issues/52511#issuecomment-329803092
[11]: https://github.com/kubernetes/kubernetes/issues/52511#issuecomment-329803092
[12]: https://kubernetes.io/docs/admin/extensible-admission-controllers/#configure-webhook-admission-controller-on-the-fly
[13]: https://github.com/kubernetes/kubernetes/pull/54733#issuecomment-343287758
[14]: https://github.com/kubernetes/kubernetes/pull/54733
[15]: https://github.com/kubernetes/features/issues/516
[16]: https://github.com/kubernetes/kubernetes/issues/54733
[17]: https://github.com/kubernetes/kubernetes/issues/54163
[18]: https://github.com/kubernetes/kubernetes/issues/52511
[19]: https://github.com/kubernetes/kubernetes/issues/52511#issue-257890224

