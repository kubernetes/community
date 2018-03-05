# Support port ranges in services

Authors: @m1093782566

## Background & Motivations

There are several applications like SIP apps or RTP which needs a lot of ports to run multiple calls or media streams. Currently there is no way to allow a range in ports in spec. So essentially users may have to do this:

```
- name: sip-udp5060
  containerPort: 5060
  protocol: UDP
- name: sip-udp5061
  containerPort: 5061
  protocol: UDP
  ...
- name: sip-udp5559
  containerPort: 5559
  protocol: UDP
```

Doing above for 500 ports is not pretty. Can we have a way to allow port ranges like 5060-5559?

## Objective

This proposal builds off of earlier requests to [Support port ranges in services](https://github.com/kubernetes/kubernetes/issues/23864) and aims to document all the requirements of supporting port ranges in services. Each port range is a single port or single contiguous port range, ranging from low to high, for which this forwarding rule matches. Packets of the specified protocol sent to these ports will be forwarded on to the appropriate target instances.

## Goals

A server side solution to support single or multiple contiguous port range.

## Non-goal

Target port remapping in ranges.

Pod port ranges.

## User Cases

* Specifying port ranges in services declarations is mandatory for any VoIP-related application.
* With the standard RTP port range which is 16384-32767. Clients connect on the standard RTP port range with a media server application.
* Needs this change to make TeleStax (an RTC framework) work well in Kubernetes.

## Proposal

This proposal builds off earlier feature requrest to [Support port ranges in services](https://github.com/kubernetes/kubernetes/issues/23864). We aims to implement service port ranges at least for iptables and IPVS modes.

### Service API changes

Create a new filed `portRangeSize` in `ServicePort`.


```go
type ServiceSpec struct {
  Ports []ServicePort
}

type ServicePort struct {
  // portRangeSize is the contiguous ports number that are exposed by this service. Given port = 80 and portRangeSize = 100, it means ports 80-179, 100 ports in total, will be exposed. In this case, port means the starting port of a range.
  // portRangeSize is default to be 1 for compatibility.
  // If nodePort is specified, then nodePort is the starting port of range and nodePort+portRangeSize-1 is the ending port of range.
  // A service can expose multiple port ranges and each port range should not be overlapped.
  // Service port range does not support target port remapping, it implies that targetPort should not be specified or should be equal to service port.
  // +optional
  PortRangeSize int `json: "portRangeSize,omitempty"`
}
```

### Service port ranges example

```yaml
apiVersion: v1
kind: Service
metadata:
  name: service-port-ranges
  namespace: default
spec:
  ports:
  - name: range1
    port: 90
    protocol: TCP
    portRangeSize: 100
    protocol: UDP
  - name: range2
    port: 200
    portRange: 300
    protocol: TCP
```

### Kube-proxy implementation limitations

NodePort and ExternalIPs Services don't support port ranges since kube-proxy will hold local ports open for these services(nodeport and "external" IP happens to be an IP that is local to this machine). It's challenging to keep some huge port ranges open on host.

Target port remapping is not supported since both iptables and ipvs does not do port translations in the mean time of doing port ranges.

Userspace proxy mode in both linux and windows platform does not work with port ranges. Kube-proxy running in these modes will log error message if people create a service with `portRangeSize` > 1.

#### iptables mode

iptables `multiport` modules support specifying multiple ports in a single line. Suppose service VIP is `1.2.3.4` and port ranges are `1000-2000` and `3000-4000`, then iptables rules might looks like,

```shell
iptables -A KUBE-SERVICES -d 1.2.3.4/32 -p tcp -m multiport --dports 1000:2000,3000:4000 -j KUBE-SVC-FOO
```

Given a nodeport type service and port range is `1000-2000`, the iptables rule might looks like,

```shell
iptables -A NODE-SERVICES -p tcp -m multiport --dports 1000:2000 --dst-type LOCAL -j KUBE-NODEPORTS
```

#### IPVS mode

IPVS + FWMARK can do port ranges. For example,

```shell
iptables -A PREROUTNG -t mangle -d 172.16.52.57 -p tcp --dport 1000:2000,3000:4000 -j MARK --set-mark 1

ipvsadm -A -f 1 -s rr
ipvsadm -a -f 1 -r 172.16.52.60 -m
ipvsadm -a -f 1 -r 172.16.52.61 -m
```

Create a fwmark bit range [0-31], for stipulating the IPVS fwmark value range. It’s configurable and should not collison with masqBit and DropBit. By default, it's [0-13] (suppose masqBit is 14 and DropBit is 15), which means we can generate at most 8K fwmark values and cover 8K port ranges.

Given VIP 1.2.3.4 and port ranges `1000:2000` and `3000:4000`. Kube-proxy will generate a unique fwmark value in the range of [1-8191], say F1. If the packet needs to be masqueraded, then `F1 = F1 OR 0x4000`. If the packets will be dropped later, then `F1 = F1 OR 0x8000`. Kube-proxy set F1 to the FWMARK field (will be created later) of IPVS virtual server and apply it to kernel.

Netfilter will test if the packet should be masqueraded by bit testing `F1 & 0X4000 == 0X4000` in POSTROUTING chain. Similar thing will happen for checking if the packet should be dropped.
