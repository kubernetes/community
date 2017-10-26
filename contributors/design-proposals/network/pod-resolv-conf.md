# Custom /etc/resolv.conf
* Status: pending
* Version: alpha
* Implementation owner: Bowei Du <[bowei@google.com](mailto:bowei@google.com)>, Zihong Zheng <[zihongz@google.com](mailto:zihongz@google.com)>

# Overview
The `/etc/resolv.conf` in a pod is managed by Kubelet and its contents are generated based on `pod.dnsPolicy`. For `dnsPolicy: Default`, `resolv.conf` is copied from the node where the pod is running. If the `dnsPolicy` is `ClusterFirst`, the contents of the resolv.conf is the hosts `resolv.conf` augmented with the following options:

*   Search paths to add aliases for domain names in the same namespace and cluster suffix.
*   `options ndots` to 5 to ensure the search paths are searched for all potential matches.

The configuration of both search paths and `ndots` results in query amplification of five to ten times for non-cluster internal names. This is due to the fact that each of the search path expansions must be tried before the actual result is found. This order of magnitude increase of query rate imposes a large load on the kube-dns service. At the same time, there are user applications do not need the convenience of the name aliases and do not wish to pay this performance cost.


## Existing workarounds

The current work around for this problem is to specify an FQDN for name resolution. Any domain name that ends with a period (e.g. `foo.bar.com.`) will not be search path expanded. However, use of FQDNs is not well-known practice and imposes application-level changes. Cluster operators may not have the luxury of enforcing such a change to applications that run on their infrastructure.

It is also possible for the user to insert a short shell script snippet that rewrites `resolv.conf` on container start-up. This has the same problems as the previous approach and is also awkward for the user. This also forces the container to have additional executable code such as a shell or scripting engine which increases the applications security surface area.


# Proposal sketch

Add a new `pod.dnsPolicy: Custom` that allows for user customization of `resolv.conf`.


## Pod API example

In the example below, the user wishes to add the pod namespace and a custom expansion to the search path, as they do not use the other name aliases:

```yaml
# Pod spec
apiVersion: v1
kind: Pod
metadata:
  namespace: ns1
  name: example
spec:
  containers:
  - name: example
    image: example
  dnsPolicy: Custom
  dnsParams:
    searchPaths:
    - $NAMESPACE.svc.$CLUSTER
    - my.dns.search.suffix
    - $HOST
    options:
    - ndots: 2
```

Given the following host `/etc/resolv.conf`:

```bash
nameserver 1.2.3.4
search foo.com
options: ndots: 1
```

The pod will get the following `/etc/resolv.conf`:

```bash
nameserver 10.240.0.10
# Populated from searchPaths
search ns1.svc.cluster.local my.dns.search.suffix foo.com
# Populated from options
options ndots: 2
```

## More examples

The following is a Pod dnsParams that only contains the host search paths:

```yaml
dnsParams:
  searchPaths:
  - $HOST
```

Override `ndots` and add custom search path. Note that overriding the ndot may break the functionality of some of the search paths the

```yaml
dnsParams:
  searchPaths:
  - my.custom.suffix
  - $HOST
  options:
  - ndots: 3
```

# API changes

```go
type PodSpec struct {
	...
	DNSPolicy string
	DNSParams *PodDNSParams
}

type PodDNSParams struct {
	SearchPaths: []string
	Options: []string
}

// This will not appear in types.go but is here for explication purposes.
type DNSParamsSubstitution string
const (
	DNSParamsSearchPathNamespace = "$NAMESPACE"
	DNSParamsSearchPathClusterDomain = "$CLUSTER"
	DNSParamsSearchPathHostPaths = "$HOST"
)
```

## Semantics
### searchPath
If `dnsPolicy: Custom` is used, then the `search` line will be constructed from the entries listed in `dnsParams.searchPath`:
```go
// Note: pseudocode does not include input validation.
func SearchPath(params *PodDNSParams) []string {
	var searchPaths []string
	for _, entry := range params.SearchPaths {
		if entry == "$HOST" {
			searchPaths = append(searchPaths, HostSearchPaths...)
			break
		}
		searchPaths = append(searchPaths, expand(entry))
	}
	return searchPaths
}

func expand(s string) string {
	labels := strings.Split(s, ".")
	for i := range labels {
		if labels[0] == "$" {
			labels[i] = Substitute(labels[i])
		}
	}
	return strings.Join(labels, ".")
}
```

#### Substitutions
`Substitute` will replace labels that begin with `$` with values for the given Pod:

| Substitution | Description |
| ----   | ---- |
| `$NAMESPACE` | Namespace of the Pod |
| `$CLUSTER` | Kubernetes cluster domain (e.g. `cluster.local`) |
| `$HOST` | Host search paths. This is a special substitution that must appear at the end of the `searchPaths` list |

### options
Each element of options will be copied unmodified as an options line.

### Invalid configurations

The follow configurations will result in an invalid Pod spec:

*  An invalid domain name/substitution appears in `searchPaths`.
*  `dnsParams` is MUST be empty unless `dnsPolicy: Custom` is used.
*   Number of final search paths exceeds 5 (glibc limit).
*  `$HOST` is not the last item in `searchPaths`.
*  `ndots` is greater than 15 (glibc limit).

# References

*   [Kubernetes DNS name specification](https://github.com/kubernetes/dns/blob/master/docs/specification.md)
*   [`/etc/resolv.conf manpage`](http://manpages.ubuntu.com/manpages/zesty/man5/resolv.conf.5.html)
