# Windows & Kubernetes APIs

This document will grow into an API by API list of work that needs to be done to clarify Windows & Linux differences. This will be used to help clarify what needs to be eventually implemented (need a tracking issue), or not implemented (need a doc note).


## Volumes

`V1.Pod.Volumes`

Out of the various volume types, these should all be possible on Windows but tests are lacking:

- EmptyDirVolumeSource
- Secret
- hostPath

The main gaps in Windows Server 2016 & 1709 are that symlinks are pretty much broken. The only ones that work are SMB/CIFS mount points. Workarounds need to be investigated.

`V1.Container.volumeMounts`
Mounting volumes across some (but not all) containers will need changes to Windows. Not ready in Windows Server 2016/1709.

### Links

- [FlexVolume does not work on Windows node](https://github.com/kubernetes/kubernetes/issues/56875)
- [feature proposal add SMB(cifs) volume plugin](https://github.com/kubernetes/kubernetes/issues/56005)
- [add NFS volume support for Windows](https://github.com/kubernetes/kubernetes/issues/56188)

## V1.Pod.Resources & V1.Container.ResourceRequirements

`V1.Container.ResourceRequirements.limits.cpu`
`V1.Container.ResourceRequirements.limits.memory`

Windows schedules CPU based on CPU count & percentage of cores. We need this represented because it can help optimize app performance. CPU count is immutable once set but you can change % of core allocations.

`V1.Container.ResourceRequirements.requests.cpu`
`V1.Container.ResourceRequirements.requests.memory`

Also of note, requests aren't supported. Will pod eviction policies in the kubelet ensure reserves are met by not overprovisioning the node?

Windows can either expose a NUMA topology matching the host (best performance) or fake it to be 1 big NUMA node (suboptimal). We should think of a way to turn this on/off later - probably q2 2018

### Links
[Kubernetes Container Runtime Interface (CRI) doesn't support WindowsContainerConfig and WindowsContainerResources](https://github.com/kubernetes/kubernetes/issues/56734)



## Networking features

`V1.Pod.dnsPolicy` - I think only ClusterFirst is implemented

`V1.Pod.hostNetwork` - Not feasible on Windows Server 2016 / 1709

## IPC & Pid

`V1.Pod.hostIPC`, `v1.pod.hostpid`

How important are these? They're not implemented in Windows Server 2016 / 1709, and I'm not too sure if they'd be helpful or not.

For cases where a pod/container need to talk to the host docker / containerd daemon we could map a named pipe as a volume which would offer the same functionality as the unix socket to the Linux daemons. It works in moby but isn't hooked up in the kubelet yet.

## Security

- `V1.Container.SecurityContext.Capabilities`
- `V1.Container.SecurityContext.seLinuxOptions`

These don't have Windows equivalents since the permissions model is substantially different

`V1.Container.SecurityContext.readOnlyRootFilesystem`

This is probably doable if needed but not possible in Windows Server 2016 / 1709.

### User Mapping

There are a few fields that refer to uid/gid. These probably need to be supplemented with a Windows SID (string) and username (string)

`V1.podSecurityContext.runAsUser` provides a UID
`V1.podSecurityContext.supplementalGroups` provides GID