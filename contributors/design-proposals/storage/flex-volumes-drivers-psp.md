# Allow Pod Security Policy to manage access to the Flexvolumes

## Current state

Cluster admins can control the usage of specific volume types by using Pod
Security Policy (PSP). Admins can allow the use of Flexvolumes by listing the
`flexVolume` type in the `volumes` field. The only thing that can be managed is
allowance or disallowance of Flexvolumes.

Technically, Flexvolumes are implemented as vendor drivers. They are executable
files that must be placed on every node at
`/usr/libexec/kubernetes/kubelet-plugins/volume/exec/<vendor~driver>/<driver>`.
In most cases they are scripts. Limiting driver access means not only limiting
an access to the volumes that this driver can provide, but also managing access
to executing a driver’s code (that is arbitrary, in fact).

It is possible to have many flex drivers for the different storage types. In
essence, Flexvolumes represent not a single volume type, but the different
types that allow usage of various vendor volumes.

## Desired state

In order to further improve security and to provide more granular control for
the usage of the different Flexvolumes, we need to enhance PSP. When such a
change takes place, cluster admins will be able to grant access to any
Flexvolumes of a particular driver (in contrast to any volume of all drivers).

For example, if we have two drivers for Flexvolumes (`cifs` and
`digitalocean`), it will become possible to grant access for one group to use
only volumes from DigitalOcean and grant access for another group to use
volumes from all Flexvolumes.

## Proposed changes

It has been suggested to add a whitelist of allowed Flexvolume drivers to the
PSP. It should behave similar to [the existing
`allowedHostPaths`](https://github.com/kubernetes/kubernetes/pull/50212) except
that:

1) comparison of equality will be used instead of comparison of prefixes.
2) Flexvolume’s driver field will be inspected rather than `hostPath`’s path field.

### PodSecurityPolicy modifications

```go
// PodSecurityPolicySpec defines the policy enforced.
type PodSecurityPolicySpec struct {
    ...
    // AllowedFlexVolumes is a whitelist of allowed Flexvolumes.  Empty or nil indicates that all
    // Flexvolumes may be used.  This parameter is effective only when the usage of the Flexvolumes
    // is allowed in the "Volumes" field.
    // +optional
    AllowedFlexVolumes []AllowedFlexVolume
}

// AllowedFlexVolume represents a single Flexvolume that is allowed to be used.
type AllowedFlexVolume struct {
       // Driver is the name of the Flexvolume driver.
       Driver string
}
```

Empty `AllowedFlexVolumes` allows usage of Flexvolumes with any driver. It must
behave as before and provide backward compatibility.

Non-empty `AllowedFlexVolumes` changes the behavior from "all allowed" to "all
disallowed except those that are explicitly listed here".

### Admission controller modifications

Admission controller should be updated accordingly to inspect a Pod's volumes.
If it finds a `flexVolume`, it should ensure that its driver is allowed to be
used.

### Validation rules

Flexvolume driver names must be non-empty.

If a PSP disallows to pods to request volumes of type `flexVolume` then
`AllowedFlexVolumes` must be empty. In case it is not empty, API server must
report an error.

API server should allow granting an access to Flexvolumes that do not exist at
time of PSP creation.

## Notes
It is possible to have even more flexible control over the Flexvolumes and take
into account options that have been passed to a driver. We decided that this is
a desirable feature but outside the scope of this proposal.

The current change could be enough for many cases. Also, when cluster admins
are able to manage access to particular Flexvolume drivers, it becomes possible
to "emulate" control over the driver’s options by using many drivers with
hard-coded options.
