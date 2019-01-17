# Mount options for mountable volume types

## Goal

Enable Kubernetes admins to specify mount options with mountable volumes
such as  - `nfs`, `glusterfs` or `aws-ebs` etc.

## Motivation

We currently support network filesystems: NFS, Glusterfs, Ceph FS, SMB (Azure file), Quobytes, and local filesystems such as ext[3|4] and XFS.

Mount time options that are operationally important and have no security implications should be supported. Examples are NFS's TCP mode, versions, lock mode, caching mode; Glusterfs's caching mode; SMB's version, locking, id mapping; and more.

## Design

### Mount option support in Persistent Volume Objects

Mount options can be specified as a field on PVs. For example:

``` yaml
apiVersion: v1
kind: PersistentVolume
metadata:
  name: pv0003
spec:
  capacity:
    storage: 5Gi
  accessModes:
    - ReadWriteOnce
  persistentVolumeReclaimPolicy: Recycle
  mountOptions:
    - hard
    - nolock
    - nfsvers=3
  nfs:
    path: /tmp
    server: 172.17.0.2
```


Beta support for mount options introduced via `mount-options` annotation will be supported for near future
and deprecated in future.


``` yaml
apiVersion: v1
kind: PersistentVolume
metadata:
  name: pv0003
  annotations:
    volume.beta.kubernetes.io/mount-options: "hard,nolock,nfsvers=3"
spec:
  capacity:
    storage: 5Gi
  accessModes:
    - ReadWriteOnce
  persistentVolumeReclaimPolicy: Recycle
  nfs:
    path: /tmp
    server: 172.17.0.2
```

### Mount option support in Storage Classes

Kubernetes admin can also specify mount option as a parameter in storage class.

```yaml
kind: StorageClass
apiVersion: storage.k8s.io/v1
metadata:
  name: slow
provisioner: kubernetes.io/glusterfs
parameters:
  type: gp2
mountOptions:
  - auto_mount
```

The mount option specified in Storage Class will be used while provisioning persistent volumes
and added as a field to PVs.

If admin has configured mount option for a storage type that does not support mount options,
then a "provisioning failed" event will be added to PVC and PVC will stay in pending state.

Also, if configured mount option is invalid then corresponding mount time failure error will be added to pod object.


## Preventing users from specifying mount options in inline volume specs of Pod

While mount options enable more flexibility in how volumes are mounted, it can result
in user specifying options that are not supported or are known to be problematic when
using inline volume specs.

After much deliberation it was decided that - `mountOptions` as an API parameter will not be supported
for inline volume specs.

### Error handling and plugins that don't support mount option

Kubernetes ships with volume plugins that don't support any kind of mount options. Such as `configmaps` or `secrets`,
in those cases to prevent user from submitting volume definitions with bogus mount options - plugins can define a interface function
such as:

```go
func SupportsMountOption() {
   return false
}
```

which will be used to validate the PV definition and API object will be *only* created if it passes the validation. Additionally
support for user specified mount options will be also checked when volumes are being mounted.

In other cases where plugin supports mount options (such as - `NFS` or `GlusterFS`) but mounting fails because of invalid mount
option or otherwise - an Event API object will be created and attached to the appropriate object.
