## PodSecurityContext and fsGroup

To allow for pod-level security settings the API defines `PodSecurityObject` that lets the admin set-up
several security attributes applied to all containers in the pod. One of the attributes is `fsGroup`
integer allowing Kubelet to change ownership of the files on the containers' volumes for some volume
types. When `fsGroup` is set, Kubelet would recursively traverse the volume mount and change GID ownership
of all the files on it.

### Known issues with current design

The current desing may cause mount time-outs on volumes containing many files. Typically a restart of a
pod that produces large number of files becomes impossible if `fsGroup` is set for the pod:

```
FirstSeen     LastSeen        Count   From  SubobjectPath     Type            Reason          Message
---------     --------        -----   ----  -------------     --------        ------          -------
2m            2m              1       {default-scheduler }    Normal          Scheduled       Successfully assigned <NODE_NAME>
14s           14s             1       {kubelet <NODE_NAME>}   Warning         FailedMount     Unable to mount volumes for pod "jenkins-xxx(xxx)": timeout expired waiting for volumes to attach/mount for pod "jenkins-xxx"/"xxx-xx". list of unattached/unmounted volumes=[jenkins-xxx]
14s           14s             1       {kubelet <NODE_NAME>}   Warning         FailedSync      Error syncing pod, skipping: timeout expired waiting for volumes to attach/mount for pod "jenkins-xxx"/"xxx-xx". list of unattached/unmounted volumes=[jenkins-xxx]
```
## Proposed solution

In order to prevent the volume mout time-out, the users should be able to decide to postpone the file 
ownership change after the volume has been successfully mounted, e.g. from an init container. To allow
this a new boolean attribute of the `PodSecurityObject` named `fsGroupNoRecurse` is added. Setting `fsGroupNoRecurse`
to `true` would then cause kubelet to change ownership only on the top-level directory (the mountpoint itself)
and it would be up to the user to ensure the files ownership change by other means. 

### Alternatives considered

* Increase the mount timeout - just postpones the issue to bigger volumes (that are hit harder).
* Have the timeout configurable by user - needs a new API anyway and it's complicated.
* Stop chown-ing the files/directories after a timeout - that would leave half of the mounted volume with correct
  permissions and second half without it.

```
// PodSecurityContext holds pod-level security attributes and common container settings.
// Some fields are also present in container.securityContext.  Field values of
// container.securityContext take precedence over field values of PodSecurityContext.
type PodSecurityContext struct {
	...
	// A special supplemental group that applies to all containers in a pod.
	// Some volume types allow the Kubelet to change the ownership of that volume
	// to be owned by the pod:
	//
	// 1. The owning GID will be the FSGroup
	// 2. The setgid bit is set (new files created in the volume will be owned by FSGroup)
	// 3. The permission bits are OR'd with rw-rw----
	//
	// If unset, the Kubelet will not modify the ownership and permissions of any volume.
	// +optional
	FSGroup *int64
	// When fsGroup is set this flag tells kubelet not to apply it recursively on
	// all files on the volume. This is needed for volumes with many files to prevent
	// mount timeouts. It's up to the user to change ownership of the files, Kubelet
	// changes ownership only of the root of the volume.
	FSGroupNoRecurse *bool
}
```