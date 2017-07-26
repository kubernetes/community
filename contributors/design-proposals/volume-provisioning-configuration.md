## Abstract

This proposal gives regular users which cannot edit StorageClasses possibility to configure some aspects of dynamic provisioning.

## Overview

In Kubernetes 1.4 we introduced StorageCasses as a way how admins can configure dynamic provisioning of PersistentVolumes. This requires admins to create a StorageClass instance for every possible configuration, which is cumbersome. In addition, regular users that can require a PV with specific properties typically cannot create StorageClasses.

## Terminology

* *admin* - user that can edit StorageClass objects.
* *regular user* - user that can create PersistentVolumeClaims and cannot edit StorageClasses.

## Use cases

* Regular user wants get a PV of storage class "fast" in availability zone "east".
* Admin wants to allow users to create PVs with storage class "fast" only in availability zones "east" and "north".
* Regular user wants to discover what tunable properties for storage class "fast" are available and as result he can see that parameter "availability zone" with allowed values "east" and "north" is configurable.

While availability zone is used as example in use cases, provisioners can choose any property such as encryption key, IOPS or replica count as tunable.

## Design

* Users can use `PersistentVolumeClaim.spec.provisioningOptions` with a pointer to `ConfigMap` in the same namespace. This `ConfigMap` contains options for the dynamic provisioning of a PV for this PVC.
   ```go
   type PersistentVolumeClaimSpec struct {
      ...
      // ProvisiongConfiguration contains a reference to ConfigMap with configuration of dynamic provisioning for this claim.
      ProvisiongConfiguration *LocalObjectReference
   }
   ```
   
    * This way, PVCs are platform-independent and can be deployed elsewhere, as long as the user updates `ConfigMap(s)` with platform-dependent details of the provisioning.
    * PV controller would ignore existing PVs and it would **always provision a new PV** for a PVC with non-empty `spec.provisioningOptions` - we do not have any way how to check if an existing PV conforms to `spec.provisioningOptions` so dynamic provisioning is the only way.
    * All provisioners (both internal and external) must provision a PV that satisfies both `StorageClass.parameters` and `PVC.spec.provisioningOptions` or throw an error if such combination of options is not possible.
    * When user does not specify any `provisioningOptions` or the referred `ConfigMap` does not specify all allowed parameters, the provisioner either uses some defaults or throws an error that some parameters in the ConfigMap are mandatory.

* StorageClass will get a new field to tell regular users what are allowed options in `PVC.spec.provisioningOptions`:
    ```go
    type StorageClass struct {
        // <snip>

        // AllowedProvisioningOptions contains details about provisioning options that are allowed by this StorageClass.
        AllowedProvisioningOption  []ProvisioningOptionSpec
    }

    // ProvisioningOption describes one option that can be used in PVC provisioningOptions
    // when requesting a volume of a particular StorageClass.
    type ProvisioningOptionSpec struct {
        // Name of key in ConfigMap
        Key string

        // Optional default value if it is unspecified in the ConfigMap.
        DefaultValue *string

        // Optional flag whether this selector key is mandatory in the ConfigMap.
        // Defaults to "false".
        Required bool

        // Optional list of values that are allowed for this key in the ConfigMap.
        AllowedValues []string

        // Optional description. May contain newlines and links (e.g. URL of description of the option that would respect browser's Accept-Language).
        Description *string
    }
    ```

  * It's responsibility of an admin to fill `sc.allowedProvisioningOptions` when creating a StorageClass!
  * We will update examples for all supported internal provisioners with selectorOptions for easier copy/paste.
  * `kubectl describe storageclass` will show available options ad values in a user-friendly form.

* There is no code in Kubernetes that validates user supplied `ConfigMap` with `allowedProvisioningOptions` in StorageClasses. Users may add invalid option keys or invalid values to their `ConfigMaps` and it's up to the individual provisioner to validate the `ConfigMap` and throw appropriate errors.

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
allowedProvisioningOptions:
  - key: zone
    allowedValues:
        - us-east-1a
        - us-east-1b
        - us-east-1c
        - us-east-2a
        - us-east-2b
        - us-east-2c
    description: "Requested AWS availability zone"

  - key: encryption
    defaultValue: "false"
    allowedValues:
      - "true"
      - "false"
    description: "Request volume encryption"

  - key: kmsKeyId
    description: "ID of AWS encryption key"
```

kubectl output:
```shell
$ kubectl describe storageclass default
Name:           default
IsDefaultClass: Yes
Provisioner:    kubernetes.io/aws-ebs
AllowedProvisioningOptions:
    zone
      - requested AWS availability zone
      - values:
        - us-east-1a
        - us-east-1b
        - us-east-1c
        - us-east-2a
        - us-east-2b
        - us-east-2c

    encryption
      - request volume encryption
      - values:
        - true
        - false (default)

    kmsKeyId
      - ID of AWS encryption key

Parameters: type=pd-standard,zones=us-east-1a, us-east-1b, us-east-1c, us-east-2a, us-east-2b, us-east-2c

```

* All options are optional.
* There is no default zone selected, provisioner may choose any zone if PVC's ConfigMap does not require any.
* Encryption is off by default.

Regular user first cretes a `ConfigMap` with chosen provisioning options:

```
apiVersion: v1
kind: ConfigMap
metadata:
  name: my-provisioning-config
data:
  zone: us-east-1d
  type: gp2
  encrypted: "true"
  kmsKeyId: arn:aws:kms:us-east-1:123456789012:key/12345678-1234-1234-1234-123456789012
```

Then the user can create any number of PVCs referring to this `ConfigMap`:

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
  provisioningOptions:
    name: my-provisioning-config
```

* Note that this ConfigMap is used only during provisioning. Changes in the ConfigMap are not reflected in already provisioned PVs.

### Typos in StorageClass
This StorageClass shows conflict between `parameters` and `allowedProvisioningOptions`.

```yaml
apiVersion: storage.k8s.io/v1beta1
kind: StorageClass
metadata:
  name: slow
provisioner: kubernetes.io/aws-ebs
parameters:
  type: gp2
  zones: us-east-1a, us-east-1b
allowedProvisioningOptions:
  - key: zone
    allowedValues:
        - us-east-2a
        - us-east-2b
    description: "Requested AWS availability zone"
```

In this StorageClass we see that admin restricts provisioning to zones "us-east-1a" and "-1b", while `allowedProvisioningOptions` allow zones "-2a" and "-2b". When user submits a PVC with a `ConfigMap` that requests `zone: us-east-2a`, Kubernetes does no validation of this `ConfigMap` and passes this PVC to (internal) AWS provisioner. The provisioner for this storage class obeys `parameters`, i.e. zones "-1a" and "-1b" and will refuse to provision anything in "-2a", leaving users confused. It's up to the admin to fix the `StorageClass`.


## Implementation details

* During alpha we will use an annotation with encoded json instead of `StorageClass.selectorOptions`.
* During alpha we will use an annotation instead of `PVC.spec.provisioningOptions`.
* During alpha we will provide only `Key` field in `allowedProvisioningOptions`. Beta and stable will have all fields as described above.


