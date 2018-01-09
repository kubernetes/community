# Add CoreDNS for DNS-based Service Discovery

Status: Pending

Version: Alpha

Implementation Owner: @johnbelamaric

## Motivation

CoreDNS is another CNCF project and is the successor to SkyDNS, which kube-dns is based on. It is a flexible, extensible
authoritative DNS server and directly integrates with the Kubernetes API. It can serve as cluster DNS,
complying with the [dns spec](https://git.k8s.io/dns/docs/specification.md). 

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

The CoreDNS configuration file is called a `Corefile` and syntactically is the same as a
[Caddyfile](https://caddyserver.com/docs/caddyfile). The file consists of multiple stanzas called _server blocks_.
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
  * `kubernetes 10.0.0.0/8 cluster.local` connects to the Kubernetes API and [serves records](https://coredns.io/plugins/kubernetes/) for the `cluster.local` domain and reverse DNS for 10.0.0.0/8 per the [spec](https://git.k8s.io/dns/docs/specification.md)
  * `proxy . /etc/resolv.conf` [forwards](https://coredns.io/plugins/proxy) any queries not handled by other plugins (the `.` means the root domain) to the nameservers configured in `/etc/resolv.conf`

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

### Configuring Federation

Federation is implemented as a separate plugin. You simply list the federation names and
their corresponding domains.

```
. {
  errors
  log
  cache 30
  health
  prometheus
  kubernetes 10.0.0.0/8 cluster.local
  federation cluster.local {
    east east.example.com
    west west.example.com
  }
  proxy . /etc/resolv.conf
}
```

### Reverse DNS

Reverse DNS is supported for Services and Endpoints. It is not for Pods.

You have to configure the reverse zone to make it work. That means knowing the service CIDR and configuring that
ahead of time (until [#25533](https://github.com/kubernetes/kubernetes/issues/25533) is implemented).

Since reverse DNS zones are on classful boundaries, if you have a classless CIDR for your service CIDR
(say, a /12), then you have to widen that to the containing classful network. That leaves a subset of that network
open to the spoofing described in [#125](https://github.com/kubernetes/dns/issues/125); this is to be fixed
in [#1074](https://github.com/coredns/coredns/issues/1074).

PTR spoofing by manual endpoints
([#124](https://github.com/kubernetes/dns/issues/124)) would
still be an issue even with [#1074](https://github.com/coredns/coredns/issues/1074) solved (as it is in kube-dns). This could be resolved in the case
where `pods verified` is enabled but that is not done at this time.

### Deployment and Operations

Typically when deployed for cluster DNS, CoreDNS is managed by a Deployment. The
CoreDNS pod only contains a single container, as opposed to kube-dns which requires three
containers. This simplifies troubleshooting.

The Kubernetes integration is stateless and so multiple pods may be run. Each pod will have its
own connection to the API server. If you (like OpenShift) run a DNS pod for each node, you should not enable
`pods verified` as that could put a high load on the API server. Instead, if you wish to support
that functionality, you can run another central deployment and configure the per-node
instances to proxy `pod.cluster.local` to the central deployment.

All logging is to standard out, and may be disabled if
desired. In very high queries-per-second environments, it is advisable to disable query logging to
avoid I/O for every query.

CoreDNS can be configured to provide an HTTP health check endpoint, so that it can be monitored
by a standard Kubernetes HTTP health check. Readiness checks are not currently supported but
are in the works (see [#588](https://github.com/coredns/coredns/issues/588)). For Kubernetes, a
CoreDNS instance will be considered ready when it has finished syncing with the API.

CoreDNS performance metrics can be published for Prometheus.

When a change is made to the Corefile, you can send each CoreDNS instance a SIGUSR1, which will
trigger a graceful reload of the Corefile.

### Performance and Resource Load

The performance test was done in GCE with the following components:

  * CoreDNS system with machine type : n1-standard-1 ( 1 CPU, 2.3 GHz Intel Xeon E5 v3 (Haswell))
  * Client system with machine type: n1-standard-1 ( 1 CPU, 2.3 GHz Intel Xeon E5 v3 (Haswell))
  * Kubemark Cluster with 5000 nodes

CoreDNS and client are running out-of-cluster (due to it being a Kubemark cluster).

The following is the summary of the performance of CoreDNS. CoreDNS cache was disabled.

Services (with 1% change per minute\*) | Max QPS\*\* | Latency (Median) | CoreDNS memory (at max QPS) | CoreDNS CPU (at max QPS) |
------------ | ------------- | -------------- | --------------------- | ----------------- |
1,000 | 18,000 | 0.1 ms | 38 MB | 95 % |
5,000 | 16,000 | 0.1 ms | 73 MB | 93 % |
10,000 | 10,000 | 0.1 ms | 115 MB | 78 % |

\* We simulated service change load by creating and destroying 1% of services per minute.

\** Max QPS with < 1 % packet loss

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

In the plan for the Alpha, there will be no automated conversion of the kube-dns configuration. However, as
part of the Beta, code will be provided that will produce a proper Corefile based upon the existing kube-dns
configuration.

## Alternatives considered

Maintain existing kube-dns, add functionality to meet the currently unmet use cases above, and fix underlying issues.
Ensuring the use of memory-safe code would require replacing dnsmasq with another (memory-safe) caching DNS server,
or implementing caching within kube-dns.
