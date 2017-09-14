# Add PVC reference in Volume Stats

## Background
Pod volume stats tracked by kubelet do not currently include any information about the PVC (if the pod volume was referenced via a PVC)

This prevents exposing (and querying) volume metrics labeled by PVC name which is preferable for users, given that PVC is a top-level API object.

## Proposal

Modify ```VolumeStats``` tracked in Kubelet and populate with PVC info:

```
// VolumeStats contains data about Volume filesystem usage.
type VolumeStats struct {
	// Embedded FsStats
	FsStats
	// Name is the name given to the Volume
	// +optional
	Name string `json:"name,omitempty"`
+	// PVCRef is a reference to the measured PVC.
+	// +optional
+	PVCRef PVCReference `json:"pvcRef"`
}

+// PVCReference contains enough information to describe the referenced PVC.
+type PVCReference struct {
+	Name      string `json:"name"`
+	Namespace string `json:"namespace"`
+}
```

## Implementation
2 options are described below. Option 1 supports current requirements/requested use cases. Option 2 supports an additional use case that was being discussed and is called out for completeness/discussion/feedback.

### Option 1
- Modify ```kubelet::server::stats::calcAndStoreStats()```
    - If the pod volume is referenced via a PVC, populate ```PVCRef``` in VolumeStats using the Pod spec

    - The Pod spec is already available in this method, so the changes are contained to this function.

- The limitation of this approach is that we're limited to reporting only what is available in the pod spec (Pod namespace and PVC claimname)

### Option 2
- Modify the ```volumemanager::GetMountedVolumesForPod()``` (or add a new function) to return additional volume information from the actual/desired state-of-world caches
    - Use this to populate PVCRef in VolumeStats

- This allows us to get information not available in the Pod spec such as the PV name/UID which can be used to label metrics - enables exposing/querying volume metrics by PV name
- It's unclear whether this is a use case we need to/should support:
  * Volume metrics are only refreshed for mounted volumes which implies a bound/available PVC
  * We expect most user-storage interactions to be via the PVC
- Admins monitoring PVs (and not PVC's) so that they know when their users are running out of space or are over-provisioning would be a use case supporting adding PV information to
  metrics





