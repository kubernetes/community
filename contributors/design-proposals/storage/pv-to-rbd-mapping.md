# RBD Volume to PV Mapping 

Authors: krmayankk@

### Problem

The RBD Dynamic Provisioner currently generates rbd volume names which are random. 
The current implementation generates a UUID and the rbd image name becomes 
image := fmt.Sprintf("kubernetes-dynamic-pvc-%s", uuid.NewUUID()). This RBD image 
name is stored in the PV. The PV also has a reference to the PVC to which it binds. 
The problem with this approach is that if there is a catastrophic etcd data loss 
and all PV's are gone, there is no way to recover the mapping from RBD to PVC. The
RBD volumes for the customer still exist, but we have no way to tell which rbd 
volumes belong to which customer.

## Goal
We want to store some information about the PVC in RBD image name/metadata, so that 
in catastrophic situations, we can derive the PVC name from rbd image name/metadata
and allow customer the following options:
- Backup RBD volume data for specific customers and hand them their copy before deleting 
  the RBD volume. Without knowing from rbd image name/metadata, which customers they 
  belong to we cannot hand those customers their data.
- Create PV with the given RBD name and pre-bind it to the desired PVC so that customer
  can get its data back.

## Non Goals
This proposal doesnt attempt to undermine the importance of etcd backups to restore
data in catastrophic situations. This is one additional line of defense in case our
backups are not working.

## Motivation

We recently had an etcd data loss which resulted in loss of this rbd to pv mapping 
and there was no way to restore customer data. This proposal aims to store pvc name 
as metadata in the RBD image so that in catastrophic scenarios, the mapping can be 
restored by just looking at the RBD's.

## Current Implementation

```go
func (r *rbdVolumeProvisioner) Provision() (*v1.PersistentVolume, error) {
...

 // create random image name
 image := fmt.Sprintf("kubernetes-dynamic-pvc-%s", uuid.NewUUID())
 r.rbdMounter.Image = image
```
## Finalized Proposal
Use `rbd image-meta set` command to store additional metadata in the RBD image about the PVC which owns
the RBD image. 

`rbd image-meta set  --pool hdd kubernetes-dynamic-pvc-fabd715f-0d24-11e8-91fa-1418774b3e9d pvcname <pvcname>`
`rbd image-meta set  --pool hdd kubernetes-dynamic-pvc-fabd715f-0d24-11e8-91fa-1418774b3e9d pvcnamespace <pvcnamespace>`

### Pros
- Simple to implement
- Does not cause regression in RBD image names, which remains same as earlier.
- The metadata information is not immediately visible to RBD admins

### Cons
- NA

Since this Proposal does not change the RBD image name and is able to store additional metadata about 
the PVC to which it belongs, this is preferred over other two proposals. Also it does a better job 
of hiding the PVC name in the metadata rather than making it more obvious in the RBD image name. The
metadata can only be seen by admins with appropriate permissions to run the rbd image-meta command. In
addition, this Proposal , doesnt impose any limitations on the length of metadata that can be stored 
and hence can accommodate any pvc names and namespaces which are stored as arbitrary key value pairs.
It also leaves room for storing any other metadata about the PVC.


### Upgrade/Downgrade Behavior

#### Upgrading from a K8s version without this metadata to a version with this metadata
The metadata for image is populated on CreateImage. After an upgrade, existing RBD Images will not have that
metadata set. When the next AttachDisk happens, we can check if the metadata is not set, set it. Cluster
administrators could also run a one time script to set this manually. For all newly created RBD images,
the rbd image metadata will be set properly.

#### Downgrade from a K8s version with this metadata to a version without this metadata
After a downgrade, all existing RBD images will have the metadata set. New RBD images created after the 
downgrade will not have this metadata. 

## Proposal 1

Make the RBD Image name as base64 encoded PVC name(namespace+name)

```go
import b64 "encoding/base64"
...


func (r *rbdVolumeProvisioner) Provision() (*v1.PersistentVolume, error) {
...

 // Create a base64 encoding of the PVC Namespace and Name
 rbdImageName := b64.StdEncoding.EncodeToString([]byte(r.options.PVC.Name+"/"+r.options.PVC.Namespace))

 // Append the base64 encoding to the string `kubernetes-dynamic-pvc-`
 rbdImageName = fmt.Sprintf("kubernetes-dynamic-pvc-%s", rbdImageName)
 r.rbdMounter.Image = rbdImageName

```

### Pros
- Simple scheme which encodes the fully qualified PVC name in the RBD image name

### Cons
- Causes regression since RBD image names will change from one version of K8s to another.
- Some older versions of librbd/krbd start having issues with names longer than 95 characters.


## Proposal 2 

Make the RBD Image name as the stringified PVC namespace plus PVC name.

### Pros
- Simple to implement. 

### Cons
- Causes regression since RBD image names will change from one version of K8s to another.
- This exposes the customer name directly to Ceph Admins. Earlier it was hidden as base64 encoding


## Misc
- Document how Pre-Binding of PV to PVC works in dynamic provisioning 
- Document/Test if there are other issues with restoring PVC/PV after a 
  etcd backup is restored
