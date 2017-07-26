## Abstract

This proposal gives regular users which cannot edit StorageClasses possibility
to configure some aspects of dynamic provisioning.

## Overview

In Kubernetes 1.4 we introduced StorageCasses as a way how admins can configure
dynamic provisioning of PersistentVolumes. This requires admins to create
a StorageClass instance for every possible configuration, which is cumbersome.
In addition, regular users that can require a PV with specific properties
typically cannot create StorageClasses.

## Terminology

* *admin* - user that can edit StorageClass objects.
* *regular user* - user that can create PersistentVolumeClaims and cannot edit StorageClasses.

## Use cases

* Regular user wants get a PV of storage class "fast" in availability zone "east".
* Admin wants to allow users to create PVs with storage class "fast" only in availability zones "east" and "north".
* Regular user wants to discover what tunable properties for storage class "fast" are available and as result he can see that parameter "availability zone" with allowed values "east" and "north" is available.

While availability zone is used as example in use cases, provisioners can choose any property such as encryption key, IOPS or replica count as tunable.

## Design

* Parameters that are tunable by regular users are implemented as label selector in `pvc.spec.selector`. We already use labels for zones and using a label selector for provisioning was already discussed in [storage class proposal](volume-provisioning.md).
* All provisioners (both internal and external) will allow both `pvc.spec.selector` and `pvc.spec.class` to be set. The provisioners *must* satisfy both these conditions during provisioning, i.e. they must provision something that has labels required by the selector and all properties of the class. If this combination is not possible to achieve, the provisioner must signal an error to the user in usual way (i.e. via an event on the PVC).
  * In other words, if the selector requires a PV to have labels `"zone" not in ["east", "south"]`, the provisioner must understand what "zone" means, what zones are available and provision something e.g. in zone "north".
  * The provisioner is required to validate the pvc.spec.selector and refuse any labels that it does not understand. Provisioners may choose not to allow any selector to be set if they don't support any parameters to be tunable by regular users.

* StorageClass will get a new field to tell regular users what are allowed labels in pvc.spec.selector:
    ```
    type StorageClass struct {
        // <snip>

        // selectorOptions contains details about labels that are allowed by this
        // StorageClass in selectors of PersistentVolumeClaims.
        SelectorOptions []SelectorComponent
    }

    // SelectorComponent describes one label that can be used in PVC selector
    // when requesting a volume of a particular StorageClass
    type SelectorComponent struct {
        // Label selector key.
        Key string

        // Optional default selector value if it is unspecified in the PVC.
        DefaultValue *string

        // Optional flag whether this selector key is mandatory in the PVC.
        // Defaults to "false".
        Required bool

        // Optional list of values that are allowed for this label selector key.
        AllowedValues []string

        // Optional description. May contain newlines and links.
        Description *string
    }
    ```

  * It's responsibility of an admin to fill `sc.selectorOptions` when creating a StorageClass!
  * We will update examples for all supported internal provisioners with selectorOptions for easier copy/paste.

* `kubectl describe storageclass` will show available labels ad values in a user-friendly form.

* There is no code in Kubernetes that validates PVCs with `selectorOptions` in StorageClasses. Users may add invalid selector keys or invalid values and it's up to the individual provisioner to validate the selector.

## Examples

### Ultimate StorageClass for AWS
This example shows a rich StorageClass with three user-tunable parameters.

```yaml
apiVersion: storage.k8s.io/v1beta1
kind: StorageClass
metadata:
  name: slow
provisioner: kubernetes.io/aws-ebs
parameters:
  type: gp2
  zones: us-east-1a, us-east-1b, us-east-1c, us-east-2a, us-east-2b, us-east-2c
selectorOptions:
  - key: failure-domain.beta.kubernetes.io/zone
    allowedValues:
        - us-east-1a
        - us-east-1b
        - us-east-1c
        - us-east-2a
        - us-east-2b
        - us-east-2c
    description: "Requested AWS availability zone"

  - key: encryption.beta.kubernetes.io/enabled
    defaultValue: "false"
    allowedValues:
      - "true"
      - "false"
    description: "Request volume encryption"

  - key: encryption.beta.kubernetes.io/key-id
    description: "ID of AWS encryption key"
```

kubectl output:
```shell
$ kubectl describe storageclass default
Name:           default
IsDefaultClass: Yes
Provisioner:    kubernetes.io/aws-ebs
SelectorOptions:
    failure-domain.beta.kubernetes.io/zone
      - requested AWS availability zone
      - values:
        - us-east-1a
        - us-east-1b
        - us-east-1c
        - us-east-2a
        - us-east-2b
        - us-east-2c

    encryption.beta.kubernetes.io/enabled
      - request volume encryption
      - values:
        - true
        - false (default)

    encryption.beta.kubernetes.io/key-id
      - ID of AWS encryption key

Parameters: type=pd-standard,zones=us-east-1a, us-east-1b, us-east-1c, us-east-2a, us-east-2b, us-east-2c

```

* All selector keys are optional.
* There is no default zone selected, provisioner may choose any zone if PVC selector does not require any.
* Encryption is off by default.


Regular user could then post a PVC to require an unencrypted PV in zone "1a", "1b" or "1c":

```yaml
kind: PersistentVolumeClaim
apiVersion: v1
metadata:
  name: myclaim
spec:
  class: slow
  accessModes:
    - ReadWriteOnce
  resources:
    requests:
      storage: 1500Mi
  selector:
    matchLabels:
      encryption.beta.kubernetes.io/enabled: "false"
    matchExpressions:
      - key: failure-domain.beta.kubernetes.io/zone
        operator: In
        values:
          - us-east-1a
          - us-east-1b
          - us-east-1c
```

### Typos in StorageClass
This StorageClass shows conflict between `parameters` and `selectorOptions`.

```yaml
apiVersion: storage.k8s.io/v1beta1
kind: StorageClass
metadata:
  name: slow
provisioner: kubernetes.io/aws-ebs
parameters:
  type: gp2
  zones: us-east-1a, us-east-1b
selectorOptions:
  - key: failure-domain.beta.kubernetes.io/zone
    allowedValues:
        - us-east-2a
        - us-east-2b
    description: "Requested AWS availability zone"
```

In this StorageClass we see that admin restricts provisioning to zones "us-east-1a" and "-1b", while `selectorOptions` allow zones "-2a" and "-2b". When user submits a PVC with selector that requests "zone" = "us-east-2a", Kubernetes does no validation of this selector and passes this PVC to (internal) AWS provisioner. The provisioner for this storage class obeys `parameters`, i.e. zones "-1a" and "-1b" and will refuse to provision anything in "-2a", leaving users confused. It's up to the admin to fix the StorageClass.


## Implementation details

* During alpha we will use an annotation with encoded json instead of `StorageClass.selectorOptions`.
* During alpha we will provide only `Key` field in `SelectorComponent`. Beta and stable will have all fields as described above.
