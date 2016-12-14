# Optional ConfigMaps

## Goal

Allow the ConfigMaps that are used to populate the environment variables of a
container and files within a Volume to be optional.

## Use Cases

An application developer can define a set of environment variables for a
container by referencing a ConfigMap with default values. The developer can
also define an optional ConfigMap that is used to override the values. There
may be a set of these optional ConfigMaps that provide environment specific
overrides for a given test cluster or region cluster. If no overrides are
present, the container will get the default values.

An application developer can populate a volume with files defined from a
ConfigMap. The developer may have some required files to be created and have
optional additional files at a different target. The developer can specify on
the Pod that there is an optional ConfigMap that will provide these additional
files if present.

## Design Points

A container can specify a set of existing ConfigMaps to populate environment
variables. These ConfigMaps can be marked as optional. The default behavior
will cause the container to fail to start when the ConfigMap is not defined.
If the ConfigMap is marked optional, when the ConfigMap is not defined, the
container will skip this ConfigMap and continue.

A pod can specify a set of Volumes to mount. A ConfigMap can represent the
files to populate the volume. The ConfigMaps should be allowed to be optional.
The default behavior will fail the volume setup if the ConfigMap is missing.
If the ConfigMap is marked as optional, the volume will be set up with no
files, and continue.

## Proposed Design

To support an optional ConfigMap either as a ConfigMapEnvSource or a
ConfigMapVolumeSource, a boolean will be added to specify whether it is
optional. The default will be required.

### Kubectl updates

The `describe` command will display the additional optional field of the
ConfigMap for both the environment variable and volume source.

### API Resource

A new `Optional` field of type boolean will be added.

```go
type ConfigMapEnvSource struct {
  // Specify whether the ConfigMap must be defined
  // +optional
	Optional *bool `json:"optional,omitempty" protobuf:"varint,2,opt,name=optional"`
}

type ConfigMapVolumeSource struct {
  // Specify whether the ConfigMap must be defined
  // +optional
	Optional *bool `json:"optional,omitempty" protobuf:"varint,4,opt,name=optional"`
}
```

### Examples

Optional `ConfigMap` as Environment Variables

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: config-env-example
spec:
  containers:
  - name: etcd
    image: openshift/etcd-20-centos7
    ports:
    - containerPort: 2379
      protocol: TCP
    - containerPort: 2380
      protocol: TCP
    envFrom:
    - configMap:
        name: etcd-env-config
        optional: true
```

Optional `ConfigMap` as a VolumeSource

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: config-env-example
spec:
  volumes:
  - name: pod-configmap-volume
    configMap:
      name: configmap-test-volume
      optional: true
  containers:
  - name: etcd
    image: openshift/etcd-20-centos7
    ports:
    - containerPort: 2379
      protocol: TCP
    - containerPort: 2380
      protocol: TCP
```
