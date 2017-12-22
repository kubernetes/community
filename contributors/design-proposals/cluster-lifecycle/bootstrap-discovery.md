# Super Simple Discovery API

## Overview

It is surprisingly hard to figure out how to talk to a Kubernetes cluster.  Not only do clients need to know where to look on the network, they also need to identify the set of root certificates to trust when talking to that endpoint.

This presents a set of problems:
* It should be super easy for users to configure client systems with a minimum of effort `kubectl` or `kubeadm init` (or other client systems).
  * Establishing this should be doable even in the face of nodes and master components booting out of order.
  * We should have mechanisms that don't require users to ever have to manually manage certificate files.
* Over the life of the cluster this information could change and client systems should be able to adapt.

While this design is mainly being created to help `kubeadm` possible, these problems aren't isolated there and can be used outside of the kubeadm context.

Mature organizations should be able to distribute and manage root certificates out of band of Kubernetes installations.  In that case, clients will defer to corporation wide system installed root certificates or root certificates distributed through other means.  However, for smaller and more casual users distributing or obtaining certificates represents a challenge.

Similarly, mature organizations will be able to rely on a centrally managed DNS system to distribute the location of a set of API servers and keep those names up to date over time.  Those DNS servers will be managed for high availability.

With that in mind, the proposals here will devolve into simply using DNS names that are validated with system installed root certificates.

## Cluster location information (aka ClusterInfo)

First we define a set of information that identifies a cluster and how to talk to it.  We will call this ClusterInfo in this document.

While we could define a new format for communicating the set of information needed here, we'll start by using the standard [`kubeconfig`](http://kubernetes.io/docs/user-guide/kubeconfig-file/) file format.

It is expected that the `kubeconfig` file will have a single unnamed `Cluster` entry.  Other information (especially authentication secrets) MUST be omitted.

### Evolving kubeconfig

In the future we look forward to enhancing `kubeconfig` to address some issues.  These are out of scope for this design.  Some of this is covered in [#30395](https://github.com/kubernetes/kubernetes/issues/30395).

Additions include:

* A cluster serial number/identifier.
  * In an HA world, API servers may come and go and it is necessary to make sure we are talking to the same cluster as we thought we were talking to.
* A _set_ of addresses for finding the cluster.
  * It is implied that all of these are equivalent and that a client can try multiple until an appropriate target is found.
  * Initially I'm proposing a flat set here.  In the future we can introduce more structure that hints to the user which addresses to try first.
* Better documentation and exposure of:
  * The root certificates can be a bundle to enable rotation.
  * If no root certificates are given (and the insecure bit isn't set) then the client trusts the system managed list of CAs.

### Client caching and update

**This is to be implemented in a later phase**

Any client of the cluster will want to have this information.  As the configuration of the cluster changes we need the client to keep this information up to date.  The ClusterInfo ConfigMap (defined below) is expected to be a common place to get the latest ClusterInfo for any cluster.  Clients should periodically grab this and cache it.  It is assumed that the information here won't drift so fast that clients won't be able to find *some* way to connect.

In exceptional circumstances it is possible that this information may be out of date and a client would be unable to connect to a cluster.  Consider the case where a user has kubectl set up and working well and then doesn't run kubectl for quite a while.  It is possible that over this time (a) the set of servers will have migrated so that all endpoints are now invalid or (b) the root certificates will have rotated so that the user can no longer trust any endpoint.

## Methods

Now that we know *what* we want to get to the client, the question is how.  We want to do this in as secure a way possible (as there are cryptographic keys involved) without requiring a lot of overhead in terms of information that needs to be copied around.

### Method: Out of Band

The simplest way to obtain ClusterInfo this would be to simply put this object in a file and copy it around.  This is more overhead for the user, but it is easy to implement and lets users rely on existing systems to distribute configuration.

For the `kubeadm` flow, the command line might look like:

```
kubeadm join --discovery-file=my-cluster.yaml
```

After loading the ClusterInfo from a file, the client MAY look for updated information from the server by reading the `kube-public/cluster-info` ConfigMap defined below.  However, when retrieving this ConfigMap the client MUST validate the certificate chain when talking to the API server.

**Note:** TLS bootstrap (which establishes a way for a client to authenticate itself to the server) is a separate issue and has its own set of methods.  This command line may have a TLS bootstrap token (or config file) on the command line also.  For this reason, even thought the `--discovery-file` argument is in the form of a `kubeconfig`, it MUST NOT contain client credentials as defined above.

### Method: HTTPS Endpoint

If the ClusterInfo information is hosted in a trusted place via HTTPS you can just request it that way.  This will use the root certificates that are installed on the system.  It may or may not be appropriate based on the user's constraints.  This method MUST use HTTPS.  Also, even though the payload for this URL is the `kubeconfig` format, it MUST NOT contain client credentials.

```
kubeadm join --discovery-file="https://example/mycluster.yaml"
```

This is really a shorthand for someone doing something like (assuming we support stdin with `-`):

```
curl https://example.com/mycluster.yaml | kubeadm join --discovery-file=-
```

After loading the ClusterInfo from a URL, the client MAY look for updated information from the server by reading the `kube-public/cluster-info` ConfigMap defined below.  However, when retrieving this ConfigMap the client MUST validate the certificate chain when talking to the API server.

**Note:** support for loading from stdin for `--discovery-file` may not be implemented immediately.

### Method: Bootstrap Token

There won't always be a trusted external endpoint to talk to and transmitting
the locator file out of band is a pain.  However, we want something more secure
than just hitting HTTP and trusting whatever we get back.  In this case, we
assume we have the following:

  * An address for at least one of the API servers (which will implement this API).
    * This address is technically an HTTPS URL base but is often expressed as a bare domain or IP.
  * A shared secret token

An interesting aspect here is that this information is often easily obtained before the API server is configured or started.  This makes some cluster bring-up scenarios much easier.

The user experience for joining a cluster would be something like:

```
kubeadm join --token=ae23dc.faddc87f5a5ab458 <address>:<port>
```

**Note:** This is logically a different use of the token used for authentication for TLS bootstrap.  We harmonize these usages and allow the same token to play double duty.

#### Implementation Flow

`kubeadm` will implement the following flow:

* `kubeadm` connects to the API server address specified over TLS.  As we don't yet have a root certificate to trust, this is an insecure connection and the server certificate is not validated.  `kubeadm` provides no authentication credentials at all.
  * Implementation note: the API server doesn't have to expose a new and special insecure HTTP endpoint.
  * (D)DoS concern: Before this flow is secure to use/enable publicly (when not bootstrapping), the API Server must support rate-limiting. There are a couple of ways rate-limiting can be implemented to work for this use-case, but defining the rate-limiting flow in detail here is out of scope. One simple idea is limiting unauthenticated requests to come from clients in RFC1918 ranges.
* `kubeadm` requests a ConfigMap containing the kubeconfig file defined above.
  * This ConfigMap exists at a well known URL: `https://<server>/api/v1/namespaces/kube-public/configmaps/cluster-info`
  * This ConfigMap is really public.  Users don't need to authenticate to read this ConfigMap.  In fact, the client MUST NOT use a bearer token here as we don't trust this endpoint yet.
* The API server returns the ConfigMap with the kubeconfig contents as normal
  * Extra data items on that ConfigMap contains JWS signatures.  `kubeadm` finds the correct signature based on the `token-id` part of the token.  (Described below).
* `kubeadm` verifies the JWS and can now trust the server.  Further communication is simpler as the CA certificate in the kubeconfig file can be trusted.


#### NEW: Bootstrap Token Structure

To first make this work, we put some structure into the token.  It has both a token identifier and the token value, separated by a dot.  Example:

```
ae23dc.faddc87f5a5ab458
```

The first part of the token is the `token-id`.  The second part is the `token-secret`.  By having a token identifier, we make it easier to specify *which* token you are talking about without sending the token itself in the clear.

This new type of token is different from the current CSV token authenticator that is currently part of Kubernetes.  The CSV token authenticator requires an update on disk and a restart of the API server to update/delete tokens.  As we prove out this token mechanism we may wish to deprecate and eventually remove that mechanism.

The `token-id` must be 6 characters and the `token-secret` must be 16 characters.  They must be lower case ASCII letters and numbers.  Specifically it must match the regular expression: `[a-z0-9]{6}\.[a-z0-9]{16}`.  There is no strong reasoning behind this beyond the history of how this has been implemented in alpha versions.

#### NEW: Bootstrap Token Secrets

Bootstrap tokens are stored and managed via Kubernetes secrets in the `kube-system` namespace.  They have the type `bootstrap.kubernetes.io/token`.

The following keys are on the secret data:
* **token-id**. As defined above.
* **token-secret**. As defined above.
* **expiration**. After this time the token should be automatically deleted.  This is encoded as an absolute UTC time using RFC3339.
* **usage-bootstrap-signing**. Set to `true` to indicate this token should be used for signing bootstrap configs.  If this is missing from the token secret or set to any other value, the usage is not allowed.
* **usage-bootstrap-authentication**. Set to true to indicate that this token should be used for authenticating to the API server.  If this is missing from the token secret or set to any other value, the usage is not allowed.  The bootstrap token authenticator will use this token to auth as a user that is `system:bootstrap:<token-id>` in the group `system:bootstrappers`.
* **description**. An optional free form description field for denoting the purpose of the token.  If users have especially complex token management neads, they are encouraged to use labels and annotations instead of packing machined readable data in to this field.
* **auth-groups**. A comma-separated list of which groups the token should be authenticated as. All groups must have the `system:bootstrappers:` prefix.


These secrets MUST be named `bootstrap-token-<token-id>`.  If a token doesn't adhere to this naming scheme it MUST be ignored.  The secret MUST also be ignored if the `token-id` key in the secret doesn't match the name of the secret.

#### Quick Primer on JWS

[JSON Web Signatures](https://tools.ietf.org/html/rfc7515) are a way to sign, serialize and verify a payload.  It supports both symmetric keys (aka shared secrets) along with asymmetric keys (aka public key infrastructure or key pairs).  The JWS is split in to 3 parts:
1. a header about how it is signed
2. the clear text payload
3. the signature.

There are a couple of different ways of encoding this data -- either as a JSON object or as a set of BASE64URL strings for including in headers or URL parameters.  In this case, we are using a shared secret and the HMAC-SHA256 signing algorithm and encoding it as a JSON object.  The popular JWT (JSON Web Tokens) specification is a type of JWS.

The JWS specification [describes how to encode](https://tools.ietf.org/html/rfc7515#appendix-F) "detached content".  In this way the signature is calculated as normal but the content isn't included in the signature.

#### NEW: `kube-public` namespace

Kubernetes ConfigMaps are per-namespace and are generally only visible to principals that have read access on that namespace.  To create a config map that *everyone* can see, we introduce a new `kube-public` namespace.  This namespace, by convention, is readable by all users (including those not authenticated). Note that is a convention
(to expose everything in `kube-public`), not something that's done by default in Kubernetes. `kubeadm` does _solely_ expose the `cluster-info` ConfigMap, not anything else.

In the initial implementation the `kube-public` namespace (and the `cluster-info` config map) will be created by `kubeadm`. That means that these won't exist for clusters that aren't bootstrapped with `kubeadm`.  As we have need for this configmap in other contexts (self describing HA clusters?) we'll make this be more generally available.

#### NEW: `cluster-info` ConfigMap

A new well known ConfigMap will be created in the `kube-public` namespace called `cluster-info`.

Users configuring the cluster (and eventually the cluster itself) will update the `kubeconfig` key here with the limited `kubeconfig` above.

A new controller (`bootstrapsigner`) is introduced that will watch for both new/modified bootstrap tokens and changes to the `cluster-info` ConfigMap.  As things change it will generate new JWS signatures. These will be saved under ConfigMap keys of the pattern `jws-kubeconfig-<token-id>`.

Another controller (`tokencleaner`) is introduced that deletes tokens that are past their expiration time.

Logically these controllers could run as a separate component in the control plane.  But, for the sake of efficiency, they are bundled as part of the Kubernetes controller-manager.

## `kubeadm` UX

We extend kubeadm with a set of flags and helper commands for managing and using these tokens.

### `kubeadm init` flags

* `--token` If set, this injects the bootstrap token to use when initializing the cluster.  If this is unset, then a random token is created and shown to the user.  If set explicitly to the empty string then no token is generated or created.  This token is used for both discovery and TLS bootstrap by having `usage-bootstrap-signing` and `usage-bootstrap-authentication` set on the token secret.
* `--token-ttl` If set, this sets the TTL for the lifetime of this token.  Defaults to 0 which means "forever" in v1.6 and v1.7. Defaults to `24h` in v1.8

### `kubeadm join` flags

* `--token` This sets the token for both discovery and bootstrap auth.
* `--discovery-file` If set, this will load the cluster-info from a file on disk or from a HTTPS URL (the HTTPS requirement due to the sensitive nature of the data)
* `--discovery-token` If set, (or set via `--token`) then we will be using the token scheme described above.
* `--tls-bootstrap-token` (not officially part of this spec) This sets the token used to temporarily authenticate to the API server in order to submit a CSR for signing. If the `system:csr-approver:approve-node-client-csr` ClusterRole is bound to the group the Bootstrap Token authenticates to, the CSR will be approved automatically (by the `csrapprover` controller) for a hands off joining flow.

Only one of `--discovery-file` or `--discovery-token` can be set.  If more than one is set then an error is surfaced and `kubeadm join` exits.  Setting `--token` counts as setting `--discovery-token`.

### `kubeadm token` commands

`kubeadm` provides a set of utilities for manipulating token secrets in a running server.

* `kubeadm token create [token]` Creates a token server side.  With no options this'll create a token that is used for discovery and TLS bootstrap.
  * `[token]` The actual token value (in `id.secret` form) to write in.  If unset, a random value is generated.
  * `--usages` A list of usages.  Defaults to `signing,authentication`.
    * If the `signing` usage is specified, the token will be used (by the BootstrapSigner controller in the KCM) to JWS-sign the ConfigMap and can then be used for discovery.
    * If the `authentication` usage is specified, the token can be used to authenticate for TLS bootstrap.
  * `--ttl` The TTL for this token.  This sets the expiration of the token as a duration from the current time.  This is converted into an absolute UTC time as it is written into the token secret.
  * `--description` Sets the free form description field for the token.
* `kubeadm token delete <token-id>|<token-id>.<token-secret>`
  * Users can either just specify the id or the full token.  This will delete the token if it exists.
* `kubeadm token list`
  * List tokens in a table form listing out the `token-id.token-secret`, the TTL, the absolute expiration time, the usages, and the description.
  * **Question** Support a `--json` or `-o json` way to make this info programmatic? We don't want to recreate `kubectl` here and these aren't plain API objects so we can't reuse that plumbing easily.
* `kubeadm token generate` This currently exists but is documented here for completeness.  This pure client side method just generated a random token in the correct form.

## Implementation Details

Our documentations (and output from `kubeadm`) should stress to users that when the token is configured for authentication and used for TLS bootstrap is a pretty powerful credential due to that any person with access to it can claim to be a node.
The highest risk regarding being able to claim a credential in the `system:nodes` group is that it can read all Secrets in the cluster, which may compromise the cluster.
The [Node Authorizer](/contributors/design-proposals/node/kubelet-authorizer.md) locks this down a bit, but an untrusted person could still try to
guess a node's name, get such a credential, guess the name of the Secret and be able to get that.

Users should set a TTL on the token to limit the above mentioned risk. `kubeadm` sets a 24h TTL on the node bootstrap token by default in v1.8.
Or, after the cluster is up and running, users should delete the token using `kubeadm token delete`.

After some back and forth, we decided to keep the separator in the token between the ID and Secret be a `.`.  During the 1.6 cycle, at one point `:` was implemented but then reverted.

See [kubernetes/client-go#114](https://github.com/kubernetes/client-go/issues/114) for details on creating a shared package with common constants for this scheme.

This proposal assumes RBAC to lock things down in a couple of ways.  First, it will open up `cluster-info` ConfigMap in `kube-public` so that it is readable by unauthenticated users (user `system:anonymous`).  Next, it will make it so that the identities in the `system:bootstrappers` group can only be used with the certs API to submit CSRs.  After a TLS certificate is created, the groupapprover controller will approve the CSR and the CSR identity can be used instead of the bootstrap token.

The end goal is to make it possible to delegate the client TLS Bootstrapping part to the kubelet, so `kubeadm join`'s function will solely be to verify the validity of the token and fetch the CA bundle.

The binding of the `system:bootstrappers` (or similar) group to the ability to submit CSRs is not part of the default RBAC configuration. Consumers of this feature like `kubeadm` will have to explicitly create this binding.

## Revision history

 - Initial proposal ([@jbeda](https://github.com/jbeda)): [link](https://github.com/kubernetes/community/blob/cb9f198a0763e0a7540cdcc9db912a403ab1acab/contributors/design-proposals/bootstrap-discovery.md)
 - v1.6 updates ([@jbeda](https://github.com/jbeda)): [link](https://github.com/kubernetes/community/blob/d8ce9e91b0099795318bb06c13f00d9dad41ac26/contributors/design-proposals/bootstrap-discovery.md)
 - v1.8 updates ([@luxas](https://github.com/luxas))
