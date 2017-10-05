# Example of FUSE flexVolume

Status: Pending

Version: Alpha

Implementation Owner: sig-auth

## Motivation

Different cluster deployments and applications may have different
needs for information available to containers running in Pods,
especially information about the environment where the application is
executed.

Over time, the set of fields which can be used as `EnvVarSource` has
grown, but addition of another field requires Kubernetes source
changes and recompilation. Moreover, environment variables capture
the values at the time when the containers were created and cannot
relay values that get modified over time.

Similar situation was with volumes and support for new volume types,
until `flexVolume` was introduced which allows writing new volume
type as custom programs, using simple command-line interface. With
`allowedFlexVolumes` in Pod Security Policies (PSP), list of permitted
drivers can be specified which makes it possible to have fine control
over functionality available to individual users.

Since the `flexVolume` drivers are running as `root` outside of
Pods, they can be used for operations otherwise not available
to processes running in containers. One of the possible uses is
injecting information into Pods. By treating the volumes as
a presentation of (potentially dynamic) values rather than merely
a storage, we can use FUSE (Filesystem in Userspace) to run
a process on the Node outside of the containers' namespaces which
projects the information to file entries on the filesystem.

## Proposal

In this document we present an example FUSE `flexVolume` plugin
`pod-info-fuse` which can provide information about the node on
which the Pod's containers run. It should serve as inspiration for
custom FUSE-based `flexVolume` plugins.

## User Experience

### Use Case: Access to information and data external to Pods

The example plugin provides three files in the filesystem:

* `pod.uid`: the Pod uid, retrieved during the `flexVolume` `mount`
  operation, to get Pod's unique identity;
* `pod.nodename`: the hostname of the node on which the Pod is
  scheduled, obtained on the fly via `uname` system call;
* `pod`: the PID of the process reading the file in the host
  namespace.

For the user to be able to use the plugin, it needs to be added
the PSP which allows this driver. The PSP should likely include

```
volumes:
- flexVolume
allowedFlexVolumes:
- driver: example.com/pod-info-fuse
```

An example Pod

```
apiVersion: v1
kind: Pod
metadata:
  name: test-pod
spec:
  containers:
  - image: registry.access.redhat.com/rhel7
    name: test-pod
    command: ["/usr/bin/sleep"]
    args: ["infinity"]
    volumeMounts:
    - mountPath: /mnt/pod-info
      name: pod-info
  volumes:
  - flexVolume:
      driver: example.com/pod-info-fuse
    name: pod-info
```

can be used to create a pod `test-pod` to test the plugin. After
the pod has been created, it can be tested for example via

```
kubectl exec test-pod -- cat /mnt/pod-info/pod.uid
```

which will demonstrate that the information from outside of the pod
is available via the FUSE filesystem in the pod.

## Implementation

The [pod-info-fuse.c](pod-info-fuse.c) source code is FUSE
`flexVolume` plugin which uses `libfuse` and `json-c` libraries to
implement userspace filesystem.

We have implemented the `flexVolume` plugin and the FUSE filesystem
in one source code to minimize number of processes that need to be
run. Of course, the `flexVolume` plugin could easily be written for
example in bash for greater flexibility, with the C code only handling
the actual mounting of the FUSE filesystem.

The compilation of the plugin/FUES filesystem is specific to operating
system on which the Kubernetes cluster is deployed. For example, on
RHEL, CentOS, and Fedora, `gcc`, `fuse-devel`, and `json-c-devel`
packages need to be installed (presumably via `yum` or `dnf`). Then
the source code can be compiled with

```
gcc -Wall pod-info-fuse.c $( pkg-config fuse json-c --cflags --libs) \
    -o pod-info-fuse
```

During compilation, `-D LOG_FILE=<path-to-log-file>` can be used to
enable logging of the plugin invocations and parameters passed to it.

If the plugin is to be named `example.com/pod-info-fuse`, it needs
to be placed to
`/usr/libexec/kubernetes/kubelet-plugins/volume/exec/example.com~pod-info-fuse/pod-info-fuse`
on both Kubernetes masters and nodes and master restarted.

The flexVolume's `init` has to return `"selinuxRelabel": false` to
prevent labelling operation which is not supported on the FUSE
filesystem.

The FUSE filesystem has to be mounted with `allow_other` option to
allow non-root uids (that the containers run as) to access the filesystem.

The SELinux boolean `virt_sandbox_use_fusefs` has to be set to true to
allow Kubernetes containers to access the filesystem.

### Client/Server Backwards/Forwards compatibility

Not affected.

## Alternatives considered

Container Storage Interface approach might be the long-run preferred
method but at this point flexVolume implementation will allow immediate
use.

## References

* https://github.com/kubernetes/community/blob/master/contributors/devel/flexvolume.md
* https://github.com/libfuse/libfuse/tree/fuse-2_9_bugfix

