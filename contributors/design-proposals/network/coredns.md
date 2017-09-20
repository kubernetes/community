# Add CoreDNS for DNS-based Service Discovery

Status: Pending

Version: Alpha

Implementation Owner: @johnbelamaric

## Motivation

CoreDNS is another CNCF project and is the successor to SkyDNS, which kube-dns is based on. It is a flexible, extensible
authoritative DNS server and directly integrates with the Kubernetes API. It can serve as cluster DNS,
complying with the [dns spec](https://github.com/kubernetes/dns/blob/master/docs/specification.md). 

CoreDNS has fewer moving parts than kube-dns, since it is a single executable and single process. It is written in Go so
it is memory-safe (kube-dns includes dnsmasq which is not). It supports a number of use cases that kube-dns does not
(see below). As a general-purpose authoritative DNS server it has a lot of functionality that kube-dns could not reasonably
be expected to add. See, for example, the [intro](https://docs.google.com/presentation/d/1v6Coq1JRlqZ8rQ6bv0Tg0usSictmnN9U80g8WKxiOjQ/edit#slide=id.g249092e088_0_181) or [coredns.io](https://coredns.io) or the [CNCF webinar](https://youtu.be/dz9S7R8r5gw).

## Proposal

The proposed solution is to enable the selection of CoreDNS as an alternate to Kube-DNS during cluster deployment, with the
intent to make it the default in the future.

## User Experience

### Use Cases

 * Standard DNS-based service discovery
 * Federation records
 * Stub domain support
 * Adding custom DNS entries
   * Making an alias for an external name [#39792](https://github.com/kubernetes/kubernetes/issues/39792)
   * Dynamically adding services to another domain, without running another server [#55](https://github.com/kubernetes/dns/issues/55)
   * Adding an arbitrary entry inside the cluster domain (for example TXT entries [#38](https://github.com/kubernetes/dns/issues/38))
 * Verified pod DNS entries (ensure pod exists in specified namespace)
 * Experimental server-side search path to address latency issues [#33554](https://github.com/kubernetes/kubernetes/issues/33554)
 * Limit PTR replies to the cluster CIDR [#125](https://github.com/kubernetes/dns/issues/125)
 * Serve DNS for selected namespaces [#132](https://github.com/kubernetes/dns/issues/132)
 * Serve DNS based on a label selector
 * Support for wildcard queries (e.g., `*.namespace.svc.cluster.local` returns all services in `namespace`)

By default, the user experience would be unchanged. For more advanced uses, existing users would need to modify the
ConfigMap that contains the CoreDNS configuration file.

### Configuring CoreDNS

The CoreDNS configuration file is called a `Corefile` and syntactically is the same as a [Caddyfile]
(https://caddyserver.com/docs/caddyfile). The file consists of multiple stanzas called _server blocks_.
Each of these represents a set of zones for which that server block should respond, along with the list
of plugins to apply to a given request. More details on this can be found in the 
[Corefile Explained](https://coredns.io/2017/07/23/corefile-explained/) and
[How Queries Are Processed](https://coredns.io/2017/06/08/how-queries-are-processed-in-coredns/) blog
entries.

### Configuration for Standard Kubernetes DNS

The intent is to make configuration as simple as possible. The following Corefile will behave according
to the spec, except that it will not respond to Pod queries. It assumes the cluster domain is `cluster.local`
and the cluster CIDRs are all within 10.0.0.0/8.

```
. {
  errors
  log
  cache 30
  health
  prometheus
  kubernetes 10.0.0.0/8 cluster.local
  proxy . /etc/resolv.conf
}

```

The `.` means that queries for the root zone (`.`) and below should be handled by this server block. Each
of the lines within `{ }` represent individual plugins:

  * `errors` enables [error logging](https://coredns.io/plugins/errors)
  * `log` enables [query logging](https://coredns.io/plugins/log/)
  * `cache 30` enables [caching](https://coredns.io/plugins/cache/) of positive and negative responses for 30 seconds
  * `health` opens an HTTP port to allow [health checks](https://coredns.io/plugins/health) from Kubernetes
  * `prometheus` enables Prometheus [metrics](https://coredns.io/plugins/metrics)
  * `kubernetes 10.0.0.0/8 cluster.local` connects to the Kubernetes API and serves records for the `cluster.local` domain and reverse DNS for 10.0.0.0/8.
  * `proxy . /etc/resolv.conf` forwards any queries not handled by other plugins (the `.` means the root domain) to the nameservers configured in `/etc/resolv.conf`

### Configuring Stub Domains

To configure stub domains, you add additional server blocks for those domains:

```
example.com {
  proxy example.com 8.8.8.8:53
}

. {
  errors
  log
  cache 30
  health
  prometheus
  kubernetes 10.0.0.0/8 cluster.local
  proxy . /etc/resolv.conf
}
```

## Implementation

Each distribution project (kubeadm, minikube, kubespray, and others) will implement CoreDNS as an optional
add-on as appropriate for that project.

### Client/Server Backwards/Forwards compatibility

No changes to other components are needed.

The method for configuring the DNS server will change. Thus, in cases where users have customized
the DNS configuration, they will need to modify their configuration if they move to CoreDNS.
For example, if users have configured stub domains, they would need to modify that configuration.

When serving SRV requests for headless services, some responses are different from kube-dns, though still within
the specification (see [#975](https://github.com/coredns/coredns/issues/975)). In summary, these are:

 * kube-dns uses endpoint names that have an opaque identifier. CoreDNS instead uses the pod IP with dashes.
 * kube-dns returns a bogus SRV record with port = 0 when no SRV prefix is present in the query.
   coredns returns all SRV record for the service (see also [#140](https://github.com/kubernetes/dns/issues/140))

Additionally, federation may return records in a slightly different manner (see [#1034](https://github.com/coredns/coredns/issues/1034)),
though this may be changed prior to completing this proposal.

## Alternatives considered

Maintain existing kube-dns, add functionality to meet the currently unmet use cases above, and fix underlying issues.
Ensuring the use of memory-safe code would require replacing dnsmasq with another (memory-safe) caching DNS server,
or implementing caching within kube-dns.
