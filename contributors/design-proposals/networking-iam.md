# Network strategy for IAM well-known IPs

## Introduction

AWS and GCE both define a well-known IP, 169.254.169.254, on which a HTTP
server that provides system information can be obtained.  Typically this
includes a script for booting a machine, data about the instance type and
configuration, and also often authorization information for access to the
full cloud API.

The metadata service information is therefore either sensitive or inaccurate
for many pods.  Sensitive because credentials are intended for the node, not
the pod; and inaccurate because the metadata might report more memory than the
pod has available.  As such it is often desirable to block access to the
metadata server.  In addition it can be helpful to serve alternative
information, such as alternative cloud access credentials, so that a pod can
have a defined role for accessing cloud APIs.  By using the metadata service,
existing code will work unaltered and transparently.

This proposal documents a simple "agreed configuration" which will allow IAM
proxies to operate consistently, and describes a few future enhancements for
consideration once the value has been demonstrated.

## Existing work

[#42192](https://github.com/kubernetes/kubernetes/pull/42192) in Kubernetes 1.6
on GCE created some firewall rules that by default prevent pods from reaching
169.254.169.254.  It does not affect processes not running in a container, nor
pods with `hostNetwork=true`.  It does not support proxying to a service,
though it may support proxying to a local daemonset.

```
iptables -N KUBE-METADATA-SERVER
iptables -A FORWARD -p tcp -d 169.254.169.254 --dport 80 -j KUBE-METADATA-SERVER
iptables -A KUBE-METADATA-SERVER -j DROP
```

[kube2iam](https://github.com/jtblin/kube2iam) is widely used on AWS to provide
pods with a particular IAM role.  It relies on a daemonset running on every
node, and an iptables rule such as

```
iptables \
  --append PREROUTING \
  --protocol tcp \
  --destination 169.254.169.254 \
  --dport 80 \
  --in-interface docker0 \
  --jump DNAT \
  --table nat \
  --to-destination `curl 169.254.169.254/latest/meta-data/local-ipv4`:8181
```

kube2iam defines a pod annotation for controlling the IAM role granted:
`iam.amazonaws.com/role: role-name`

## Proposal

This proposal is intended to allow for integration of a solution such as
kube2iam into the kubernetes networking infrastructure.

Specifically:

* The rules in [#42192](https://github.com/kubernetes/kubernetes/pull/42192)
  should be applied at the node level by the installation system, before
  kube-proxy is running.  This acts as a "last line of defense" in case pods are
  started before kube-proxy rules are fully in place.

* We propose adding a "well known service" for IAM, as we have for DNS.  The
  DNS service currently runs on `.10`, so we propose reserving `.9` for IAM.

* Installation systems _should_ install a service into `.9`, even if that
  service has no backends, to prevent a user service hijacking the IP.  (If a
  stronger mechanism is introduced to prevent hijacking the DNS service IP, we
  should use that also here.)

## Short-term proposal (1.6 / 1.7)

* The name of the service should be `kube-iam` in the `kube-system` namespace.
  If port 80 is mapped it should be named `http`, if port `443` is mapped is
  should be named `https`.  This will ensure that kube-proxy will name the chains
  consistently.

* Installation systems that want to support IAM redirection should add a rule
  such as this one: 
`iptables -t nat -A PREROUTING -d 169.254.169.254 -p tcp -m comment --comment "IAM redirection" -m tcp --dport 80 -j KUBE-SVC-XXXXXXXXXXXXX`

(`XXXXXXX` is determined by the kube-proxy hash of the service details, hence
the need for consistent naming)

* (TODO: Make sure kube-proxy doesn't GC this rule; it does not appear to)

Notes:

* We would deally avoid the chain name dependency, but it does not seem
  possible to rewrite the target IP and then send the packet back through the
  NAT chain processing to be redirected to an endpoint.

## Future additions / out-of-scope

* A field could be added to the Service API type, to designate that this is a
  well-known service.  This could guarantee the iptables chain name.

* We could add a field to the Service API type, or use the ExternalIP /
  ClusterIP field, so that kube-proxy would set up the complete iptables rules
  for 169.254.169.254 (and other IPs).

## Additional discussion

It is expected that the IAM proxy will likely run only on master nodes (or
otherwise isolated nodes), so that the nodes do not need IAM permissions.  An
IAM proxy implementation might choose to delegate to another instance
(such as vault) for additional security.

The mechanism is believed to accomodate multiple IAM proxy implementations, and
multiple clouds or transports, dependent on the installation system adding
appropriate iptables rules.

Although kube2iam currently relies on a daemonset, it is believed it could
easily be installed instead as a deployment and a service, without code changes.

We assume that IP spoofing is not possible by untrusted pods.  Thus the IP
address inside the cluster can be assumed to be sufficient proof of identity of
the pod.

