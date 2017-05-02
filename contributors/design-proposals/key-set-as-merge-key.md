# Multi-fields Merge Key in Strategic Merge Patch

## Abstract

Support multi-fields merge key in Strategic Merge Patch.

## Motivation

*Background*: Strategic Merge Patch is covered in this [doc](https://github.com/kubernetes/community/blob/master/contributors/devel/strategic-merge-patch.md).
In Strategic Merge Patch, Merge Key is the key to distinguish the entries in the list of non-primitive types.
It must always be present and unique to perform the merge on the list of non-primitive types,
and will be preserved.

The current implementation requires a single field that uniquely identifies each element in a list.
For some element Kinds, the identity is defined using multiple fields.
An [example](https://github.com/kubernetes/kubernetes/issues/39188) is the service.spec.ports,
which is identified by both `protocol` and `port`.

As a result we need to also support a set of keys as a Merge Key.
A key set must be a list of strings with at least 1 element long.

## Proposed Change

### API Change

For API resources that cannot be effectively merged with a single merge key,
we will update the merge keys to a key set.
We require the new key set has the old merge key and the new merge key must be present to keep backward compatibility.
The keys will be seperated by ",", i.e. `patchMergeKey:"<key1>,<key2>,<key3>"`.

E.g. [`Ports` in `ServiceSpec`](https://github.com/kubernetes/kubernetes/blob/c51efa9ba0929a643544078d5c182ba75e4b4087/pkg/api/v1/types.go#L2825-L2831).
```go
type ServiceSpec struct {
  // Change patchMergeKey from "port" to "name|port"
  Ports []ServicePort `json:"ports,omitempty" patchStrategy:"merge" patchMergeKey:"name,port" protobuf:"bytes,1,rep,name=ports"`
  ...
}
```

All the impacted APIs are listed in section [Impacted APIs](#impacted-apis)

### Open API

Update [Open API schema](https://github.com/kubernetes/kubernetes/blob/master/api/openapi-spec/swagger.json)
to reflect the change of `patchMergeKey` in the struct tags.
The change should be trivial.
E.g. change OpenAPI extension from "x-kubernetes-patch-merge-key": "port"
to "x-kubernetes-patch-merge-key": "port,protocol".

### Strategic Merge Patch pkg

Entries are considered as the same if and only if all the keys in the key set are identical.
We allow keys to be missing in the key set as long as it is not empty.

Take `Ports` as an example:

Suppose we are using `name` and `port` as merge key as mentioned in section [API Change](#api-change)

We have the following live config in the server:
```yaml
spec:
  type: NodePort
  ports:
    - protocol: UDP
      # name is missing here
      port: 30420
      nodePort: 30420
```

The users want to add another port. They will use the following manifest:
```yaml
spec:
  type: NodePort
  ports:
    - protocol: UDP
      port: 30420
      name: udpport
      nodePort: 30420
    - protocol: TCP
      port: 30420
      name: tcpport
      nodePort: 30420
```

The patch manifest that will be sent is:
```yaml
spec:
  type: NodePort
  ports:
    # the entry with key set {port: 30420} is considered missing in local config file
    - port: 30420
      $patch: delete
    # the entry with key set {name: udpport, port: 30420} is considered as a new entry
    - protocol: UDP
      port: 30420
      name: udpport
      nodePort: 30420
    # the entry with key set {name: tcpport, port: 30420} is a new entry
    - protocol: TCP
      port: 30420
      name: tcpport
      nodePort: 30420
```

### Docs

Document what the developer should consider when adding an API with `mergeKey`.

## Version Skew

*There are 2 edge cases when updating:*

Suppose `foo` is the merge key before, but `foo` and `bar` are new merge key now.
A client wants to

An client wants to change config from
```yaml
list:
- foo: a
  bar: x
```
to
```yaml
list:
- foo: a
  bar: y
```

1) When the client is an old one and talks to a new server.

The patch it generated is:
```yaml
list:
- foo: a # old merge key
  bar: y
```

The live object on the new server is:
```yaml
list:
- foo: a # new merge key 1
  bar: x # new merge key 2
```

The result after merging is:
```yaml
list:
- foo: a # new merge key 1
  bar: x # new merge key 2
- foo: a # new merge key 1
  bar: y # new merge key 2
```

2) When the client is an new one and talks to a old server.

The patch it generated is:
```yaml
list:
- $patch: delete
  foo: a # new merge key 1
  bar: x # new merge key 2
- foo: a # new merge key 1
  bar: y # new merge key 2
```

The live object on the new server is:
```yaml
list:
- foo: a # old merge key
  bar: x
  other: somevalue
```

Based on current implementation, the result after merging is:
```yaml
list:
- foo: a # old merge key
  bar: y
  # other field is missing, this entry in list has been recreated.
```

*There is another edge case when deleting:*

If we don't change the behavior, it will be the same as in 1.6.

An old client sending a patch:

```yaml
list:
- $patch: delete
  foo: a # old merge key
```

Live config in new server is:

```yaml
list:
- foo: a # new merge key 1
  bar: x # new merge key 2
- foo: a # new merge key 1
  bar: y # new merge key 2
```

After merging the patch, both entries will be deleted.
The config in the server is:
```yaml
list: []
```

## Impacted APIs

We need to examine case by case to check if it is OK to have behavior in the above 3 edge cases.

(1) `ContainerPort`: Change merge key from `containerPort` to `name,containerPort`.

Usage of [ContainerPort](https://github.com/kubernetes/kubernetes/blob/db9fcb06295b3db49be8efa5c4584114af0696bc/pkg/api/v1/types.go#L1637)
and its [definition](https://github.com/kubernetes/kubernetes/blob/db9fcb06295b3db49be8efa5c4584114af0696bc/pkg/api/v1/types.go#L1286).
```go
type Container struct {
Ports []ContainerPort `json:"ports,omitempty" patchStrategy:"merge" patchMergeKey:"containerPort" protobuf:"bytes,6,rep,name=ports"`
...
}
```
```go
type ContainerPort struct {
	// +optional
	Name string `json:"name,omitempty" protobuf:"bytes,1,opt,name=name"`
	// +optional
	HostPort int32 `json:"hostPort,omitempty" protobuf:"varint,2,opt,name=hostPort"`
	ContainerPort int32 `json:"containerPort" protobuf:"varint,3,opt,name=containerPort"`
	// +optional
	Protocol Protocol `json:"protocol,omitempty" protobuf:"bytes,4,opt,name=protocol,casttype=Protocol"`
	// +optional
	HostIP string `json:"hostIP,omitempty" protobuf:"bytes,5,opt,name=hostIP"`
}
```

(2) `ServicePort`: Similar to `ContainerPort`. Change merge key from `port` to `name,port`.

Usage of [ServicePort](https://github.com/kubernetes/kubernetes/blob/db9fcb06295b3db49be8efa5c4584114af0696bc/pkg/api/v1/types.go#L2777)
and its [definition](https://github.com/kubernetes/kubernetes/blob/db9fcb06295b3db49be8efa5c4584114af0696bc/pkg/api/v1/types.go#L2867).
```go
type ServiceSpec struct {
	Ports []ServicePort `json:"ports,omitempty" patchStrategy:"merge" patchMergeKey:"port" protobuf:"bytes,1,rep,name=ports"`
  ...
}
```
```go
type ServicePort struct {
	// +optional
	Name string `json:"name,omitempty" protobuf:"bytes,1,opt,name=name"`
	// +optional
	Protocol Protocol `json:"protocol,omitempty" protobuf:"bytes,2,opt,name=protocol,casttype=Protocol"`
	Port int32 `json:"port" protobuf:"varint,3,opt,name=port"`
	// +optional
	TargetPort intstr.IntOrString `json:"targetPort,omitempty" protobuf:"bytes,4,opt,name=targetPort"`
	// +optional
	NodePort int32 `json:"nodePort,omitempty" protobuf:"varint,5,opt,name=nodePort"`
}
```
