# Abstract
Today, Kubelet dynamically re-constructs volume spec during runtime. It reconstructs using volume mount path. But for many plugins volume mount path is not enough to reconstruct all the fields/options. As a result, some of the plugin API which depend on these fields/options are broken, especially in the cleanup code. The goal of this exercise is to store volume specific meta-data and use that information to reconstruct the spec.

# High Level Design
Store volume specific meta-data during volume mountdevice/mount/update logic and use it to reconstruct the volume spec. If it is an attachable plugin, store the meta-data during mountdevice and remove after a successful unmountdevice. For plugins, which use only mount & unmount API, store the meta-data during the mount API and remove during the umount API. Meta-data should also be updated, once we add support for volume spec update.

## Meta-data format

### Option 1:
 Store the volume object source in JSON format. Example: for GCE plugin store "GCEPersistentDiskVolumeSource".
 #### Pros:
 * Simpler to implement.
 #### Cons:
 * Volume object source is not versioned and upgrades will be a problem going forward.
### Option 2:
 Store Persistent volume spec/Volume Source in JSON format. Persistent volume spec is stored if the volume is backed by a Persistent Volume Object. Volume source is stored if the volume source is inlined in Pod spec. We can use the different filenames \<volume name\>~pv.json and \<volume name\>~vs.json to distinguish these objects.
 #### Pros:
 * Persistent volume spec is versioned.
 #### Cons:
 * Complicated naming.
 * VolumeSource is still not versioned.
 * Persistent volume does not have the namespace information for the secretref.
### Option 3:
 Implement a per plugin specific API to store the meta-data relevant to the plugin. We can provide a sensible default using Option 2/Option 1 if plugin does not want to implement these API.
 #### Pros:
 * Plugins are free to implement/experiment with the information they want to store.
 #### Cons:
 * Versioning the meta-data is offloaded to the plugin.
  
## Meta-data location:
Store the meta-data file \<volume name\>.json in the plugin path.
Ex:
For plugin: ```kubernetes.io/lvm``` store in the following directory
```/var/lib/kubelet/plugins/kubernetes.io/flexvolume/kubernetes.io/lvm```