# Containerized Mounter with Chroot for Container-Optimized OS

## Goal

Due security and management overhead, our new Container-Optimized OS used by GKE
does not carry certain storage drivers and tools needed for such as nfs and 
glusterfs. This project takes a containerized mount approach to package mount 
binaries into a container. Volume plugin will execute mount inside of container 
and share the mount with the host. 


## Design

1. A docker image has storage tools (nfs and glusterfs) pre-installed and uploaded
   to gcs. 
2. During GKE cluster configuration, the docker image is pulled and installed on 
   the cluster node.
3. When nfs or glusterfs type mount is invoked by kubelet, it will run the mount 
   command inside of a container with the pre-install docker image and the mount
   propagation set to “shared. In this way, the mount inside the container will 
   visible to host node too.
4. A special case for NFSv3, a rpcbind process is issued before running mount
   command. 
   
## Implementation details

* In the first version of containerized mounter, we use rkt fly to dynamically
  start a container during mount. When mount command finishes, the container is 
  normally exited and will be garbage-collected. However, in case the glusterfs
  mount, because a gluster daemon is running after command mount finishes util
  glusterfs unmount, the container started for mount will continue to run until 
  glusterfs client finishes. The container cannot be garbage-collected right away
  and multiple containers might be running for some time. Due to shared mount
  propagation, with more containers running, the number of mounts will increase
  significantly and might cause kernel panic. To solve this problem, a chroot
  approach is proposed and implemented. 
* In the second version, instead of running a container on the host, the docker
  container’s file system is exported as a tar archive and pre-installed on host. 
  Kubelet directory is shared mount between host and inside of the container’s 
  rootfs. When a gluster/nfs mount is issued, a mounter script will use chroot to
  change to the container’s rootfs and run the mount. This approach is very clean
  since there is no need to manage a container’s lifecycle and avoid having large
  number of mounts.
