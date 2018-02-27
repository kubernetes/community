# Optional ConfigMaps and Secrets

## Goal

Allow the ConfigMaps or Secrets that are used to populate the environment variables of a
container and files within a Volume to be optional.

## Use Cases

When deploying an application to multiple environments like development, test,
and production, there may be certain environment variables that must reflect
the values that are relevant to said environment. One way to do so would be to
have a well named ConfigMap which contains all the environment variables
needed. With the introduction of optional ConfigMaps, one could instead define a required
ConfigMap which contains all the environment variables for any environment
with a set of initialized or default values. An additional optional ConfigMap
can also be specified which allows the deployer to provide any overrides for
the current environment.

An application developer can populate a volume with files defined from a
ConfigMap. The developer may have some required files to be created and have
optional additional files at a different target. The developer can specify on
the Pod that there is an optional ConfigMap that will provide these additional
files if the ConfigMap exists.

## Design Points

A container can specify an entire ConfigMap to be populated as environment
variables via `EnvFrom`. When required, the container fails to start if the
ConfigMap does not exist. If the ConfigMap is optional, the container will
skip the non-existent ConfigMap and proceed as normal.

A container may also specify a single environment variable to retrieve its
value from a ConfigMap via `Env`. If the key does not exist in the ConfigMap
during container start, the container will fail to start. If however, the
ConfigMap is marked optional, during container start, a non-existent ConfigMap
or a missing key in the ConfigMap will not prevent the container from
starting. Any previous value for the given key will be used.

Any changes to the ConfigMap will not affect environment variables of running
containers. If the Container is restarted, the set of environment variables
will be re-evaluated.

The same processing rules applies to Secrets.

A pod can specify a set of Volumes to mount. A ConfigMap can represent the
files to populate the volume. The ConfigMaps can be marked as optional.  The
default is to require the ConfigMap existence. If the ConfigMap is required
and does not exist, the volume creation will fail.  If the ConfigMap is marked
as optional, the volume will be created regardless, and the files will be
populated only if the ConfigMap exists and has content.  If the ConfigMap is
changed, the volume will eventually reflect the new set of data available from
the ConfigMap.

## Proposed Design

To support an optional ConfigMap either as a ConfigMapKeySelector, ConfigMapEnvSource or a
ConfigMapVolumeSource, a boolean will be added to specify whether it is
optional. The default will be required.

To support an optional Secret either as a SecretKeySelector, or a
SecretVolumeSource, a boolean will be added to specify whether it is optional.
The default will be required.

### Kubectl updates

The `describe` command will display the additional optional field of the
ConfigMap and Secret for both the environment variables and volume sources.

### API Resource

A new `Optional` field of type boolean will be added.

```go
type ConfigMapKeySelector struct {
  // Specify whether the ConfigMap must be defined
  // +optional
	Optional *bool `json:"optional,omitempty" protobuf:"varint,3,opt,name=optional"`
}

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

type SecretKeySelector struct {
  // Specify whether the ConfigMap must be defined
  // +optional
	Optional *bool `json:"optional,omitempty" protobuf:"varint,3,opt,name=optional"`
}

type SecretVolumeSource struct {
  // Specify whether the Secret must be defined
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
    env:
    - name: foo
      valueFrom:
        configMapKeyRef:
          name: etcd-env-config
          key: port
          optional: true
```

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
