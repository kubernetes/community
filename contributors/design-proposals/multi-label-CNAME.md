# Multi-label subdomain support in CNAME records

Author: Laszlo Janosi (@janosi), Gergely Csatari (@CsatariGergely)

Date: August 2017

Status: New

# Goal

Enable the usage of multi-label subdomains in CNAME records via ExternalName Service.
Example: `my-service.my-cluster` could be mapped to `your-service.your-cluster`,
while with the current implementation only `my-service` can be mapped.

# Motivation
DNS CNAME records are specified in [RFC1034](https://www.ietf.org/rfc/rfc1034.txt)
and clarified in Section 10 of [RFC2181](https://www.ietf.org/rfc/rfc2181.txt).
These allow the usage of multi-label subdomains in CNAME records, and not only
on the canonical name side.
Examples:
* The client may start the resolution of a domain name with a single-labeled service name (which is a subdomain, too, actually), but the DNS query can end up in resolving a multi-label subdomain if the resolver appends the value of the search field from /etc/resolv.conf to the original string.
* legacy applications may use multi-label subdomains in their DNS queries now, and the support of multi-label subdomain in CNAME would help their integration into the kube DNS based service discovery.

Related issues is [#43907](https://github.com/kubernetes/kubernetes/issues/43907).

# Proposed solution

As it is very improbable that the validation of the metadata of a Service could
be changed to support multi-label values we propose a new field in the
ServiceSpec struct: "Name". This new field would enable the definition of
multi-label CNAME records.

The proposed Service Spec looks like:
```go
// ServiceSpec describes the attributes that a user creates on a service
type ServiceSpec struct {
	// Type determines how the Service is exposed. Defaults to ClusterIP. Valid
	// options are ExternalName, ClusterIP, NodePort, and LoadBalancer.
	// "ExternalName" maps to the specified externalName.
	// "ClusterIP" allocates a cluster-internal IP address for load-balancing to
	// endpoints. Endpoints are determined by the selector or if that is not
	// specified, by manual construction of an Endpoints object. If clusterIP is
	// "None", no virtual IP is allocated and the endpoints are published as a
	// set of endpoints rather than a stable IP.
	// "NodePort" builds on ClusterIP and allocates a port on every node which
	// routes to the clusterIP.
	// "LoadBalancer" builds on NodePort and creates an
	// external load-balancer (if supported in the current cloud) which routes
	// to the clusterIP.
	// More info: http://kubernetes.io/docs/user-guide/services#overview
	// +optional
	Type ServiceType

	// Required: The list of ports that are exposed by this service.
	Ports []ServicePort

	// Route service traffic to pods with label keys and values matching this
	// selector. If empty or not present, the service is assumed to have an
	// external process managing its endpoints, which Kubernetes will not
	// modify. Only applies to types ClusterIP, NodePort, and LoadBalancer.
	// Ignored if type is ExternalName.
	// More info: http://kubernetes.io/docs/user-guide/services#overview
	Selector map[string]string

	// ClusterIP is the IP address of the service and is usually assigned
	// randomly by the master. If an address is specified manually and is not in
	// use by others, it will be allocated to the service; otherwise, creation
	// of the service will fail. This field can not be changed through updates.
	// Valid values are "None", empty string (""), or a valid IP address. "None"
	// can be specified for headless services when proxying is not required.
	// Only applies to types ClusterIP, NodePort, and LoadBalancer. Ignored if
	// type is ExternalName.
	// More info: http://kubernetes.io/docs/user-guide/services#virtual-ips-and-service-proxies
	// +optional
	ClusterIP string

	// ExternalName is the external reference that kubedns or equivalent will
	// return as a CNAME record for this service. No proxying will be involved.
	// Must be a valid DNS name and requires Type to be ExternalName.
	ExternalName string

	// ExternalIPs are used by external load balancers, or can be set by
	// users to handle external traffic that arrives at a node.
	// +optional
	ExternalIPs []string

	// Only applies to Service Type: LoadBalancer
	// LoadBalancer will get created with the IP specified in this field.
	// This feature depends on whether the underlying cloud-provider supports specifying
	// the loadBalancerIP when a load balancer is created.
	// This field will be ignored if the cloud-provider does not support the feature.
	// +optional
	LoadBalancerIP string

	// Optional: Supports "ClientIP" and "None".  Used to maintain session affinity.
	// +optional
	SessionAffinity ServiceAffinity

	// Optional: If specified and supported by the platform, this will restrict traffic through the cloud-provider
	// load-balancer will be restricted to the specified client IPs. This field will be ignored if the
	// cloud-provider does not support the feature."
	// +optional
	LoadBalancerSourceRanges []string

+	// The name of the service which is resolved by kubedns or equivalent to a canonical name.
+ // Enable the definition of multi-label CNAME records. But also can be used to
+ // define a single-label service name, in which case it takes the role of the "name" label in metadata part.
+	// When Name is resolved, kubedns or equivalent returns with the canonical name defined in ExternalName field above.
+	// Must be a valid DNS name and requires Type to be ExternalName.
+ Name string
}
```

For example it can be used like this:

```yaml
apiVersion: v1
kind: Service
metadata:
  name: my-service
spec:
  ports:
  - port: 12345
  type: ExternalName
  name: my-service.my-cluster
  externalName: your-service.your-cluster
```
