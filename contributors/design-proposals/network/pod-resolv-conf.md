# Custom /etc/resolv.conf

* Status: pending
* Version: alpha
* Implementation owner: Bowei Du <[bowei@google.com](mailto:bowei@google.com)>,
  Zihong Zheng <[zihongz@google.com](mailto:zihongz@google.com)>

# Overview

The `/etc/resolv.conf` in a pod is managed by Kubelet and its contents are
generated based on `pod.dnsPolicy`. For `dnsPolicy: Default`, the `search` and
`nameserver` fields are taken from the `resolve.conf` on the node where the pod
is running. If the `dnsPolicy` is `ClusterFirst`, the search contents of the
resolv.conf is the hosts `resolv.conf` augmented with the following options:

*   Search paths to add aliases for domain names in the same namespace and
    cluster suffix.
*   `options ndots` to 5 to ensure the search paths are searched for all
    potential matches.

The configuration of both search paths and `ndots` results in query
amplification of five to ten times for non-cluster internal names. This is due
to the fact that each of the search path expansions must be tried before the
actual result is found. This order of magnitude increase of query rate imposes a
large load on the kube-dns service. At the same time, there are user
applications do not need the convenience of the name aliases and do not wish to
pay this performance cost.


## Existing workarounds

The current work around for this problem is to specify an FQDN for name
resolution. Any domain name that ends with a period (e.g. `foo.bar.com.`) will
not be search path expanded. However, use of FQDNs is not well-known practice
and imposes application-level changes. Cluster operators may not have the luxury
of enforcing such a change to applications that run on their infrastructure.

It is also possible for the user to insert a short shell script snippet that
rewrites `resolv.conf` on container start-up. This has the same problems as the
previous approach and is also awkward for the user. This also forces the
container to have additional executable code such as a shell or scripting engine
which increases the applications security surface area.


# Proposal sketch

This proposal gives users a way to overlay tweaks into the existing
`DnsPolicy`. A new PodSpec field `dnsParams` will contains fields that are
merged with the settings currently selected with `DnsPolicy`.

The fields of `DnsParams` are:

* `nameservers` is a list of additional nameservers to use for resolution. On
  `resolv.conf` platforms, these are entries to `nameserver`.
* `search` is a list of additional search path subdomains. On `resolv.conf`
  platforms, these are entries to the `search` setting. These domains will be
  appended to the existing search path.
* `options` that are an OS-dependent list of (name, value) options. These values
  are NOT expected to be generally portable across platforms. For containers that
  use `/etc/resolv.conf` style configuration, these correspond to the parameters
  passed to the `option` lines. Options will override if their names coincide,
  i.e, if the `DnsPolicy` sets `ndots:5` and `ndots:1` appears in the `Spec`,
  then the final value will be `ndots:1`.

For users that want to completely customize their resolution configuration, we
add a new `DnsPolicy: Custom` that does not define any settings. This is
essentially an empty `resolv.conf` with no fields defined.

## Pod API examples

### Host `/etc/resolv.conf`

Assume in the examples below that the host has the following `/etc/resolv.conf`:

```bash
nameserver 10.1.1.10
search foo.com
options ndots:1
```

### Override DNS server and search paths

In the example below, the user wishes to use their own DNS resolver and add the
pod namespace and a custom expansion to the search path, as they do not use the
other name aliases:

```yaml
# Pod spec
apiVersion: v1
kind: Pod
metadata: {"namespace": "ns1", "name": "example"}
spec:
  ...
  dnsPolicy: Custom
  dnsParams:
    nameservers: ["1.2.3.4"]
    search:
    - ns1.svc.cluster.local
    - my.dns.search.suffix
    options:
    - name: ndots
      value: 2
    - name: edns0
```

The pod will get the following `/etc/resolv.conf`:

```bash
nameserver 1.2.3.4
search ns1.svc.cluster.local my.dns.search.suffix
options ndots:2 edns0
```

## Overriding `ndots`

Override `ndots:5` in `ClusterFirst` with `ndots:1`. This keeps all of the
settings intact:

```yaml
dnsPolicy: ClusterFirst
dnsParams:
- options:
  - name: ndots
  - value: 1
```

Resulting `resolv.conf`:

```bash
nameserver 10.0.0.10
search default.svc.cluster.local svc.cluster.local cluster.local foo.com
options ndots:1
```

# API changes

```go
type PodSpec struct {
    ...
    DNSPolicy string        `json:"dnsPolicy,omitempty"`
    DNSParams *PodDNSParams `json:"dnsParams,omitempty"`
    ...
}

type PodDNSParams struct {
    Nameservers []string             `json:"nameservers,omitempty"`
    Search      []string             `json:"search,omitempty"`
    Options     []PodDNSParamsOption `json:"options,omitempty" patchStrategy:"merge" patchMergeKey:"name"`
}

type PodDNSParamsOption struct {
    Name   string  `json:"name"`
    Value *string `json:"value,omitempty"`
}
```

## Semantics

Let the following be the Go representation of the `resolv.conf`:

```go
type ResolvConf struct {
  Nameserver []string // "nameserver" entries
  Search     []string // "search" entries
  Options    []PodDNSParamsOption  // "options" entries
}
```

Let `var HostResolvConf ResolvConf` be the host `resolv.conf`.

Then the final Pod `resolv.conf` will be:

```go
func podResolvConf() ResolvConf {
    var podResolv ResolvConf

    switch (pod.DNSPolicy) {
    case "Default":
        podResolv = HostResolvConf
    case "ClusterFirst:
        podResolv.Nameservers = []string{ KubeDNSClusterIP }
        podResolv.Search = ... // populate with ns.svc.suffix, svc.suffix, suffix, host entries...
        podResolv.Options = []PodDNSParamsOption{{"ndots","5" }}
    case "Custom": // start with empty `resolv.conf`
        break
    }

    // Append the additional nameservers.
    podResolv.Nameservers = append(Nameservers, pod.DNSParams.Nameservers...)
    // Append the additional search paths.
    podResolv.Search = append(Search, pod.DNSParams.Search...)
    // Merge the DnsParams.Options with the options derived from the given DNSPolicy.
    podResolv.Options = mergeOptions(pod.Options, pod.DNSParams.Options)

    return podResolv
}
```

### Invalid configurations

The follow configurations will result in an invalid Pod spec:

* Nameservers or search paths exceed system limits. (Three nameservers, six
  search paths, 256 characters for `glibc`).
* Invalid option appears for the given platform.

# References

*   [Kubernetes DNS name specification](https://git.k8s.io/dns/docs/specification.md)
*   [`/etc/resolv.conf manpage`](http://manpages.ubuntu.com/manpages/zesty/man5/resolv.conf.5.html)
