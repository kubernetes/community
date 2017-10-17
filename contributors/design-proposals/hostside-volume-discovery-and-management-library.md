# Host-side Volume Discovery and Management Library

## Motivation

As we all know, Kubernetes can be deployed anywhere, including various cloud
platforms and bare metals. But as for storage resources, you have to choose
specific storage backend according to the environment Kubernetes be deployed.
Here are some examples:
1. if you want to use EBS to provide storage resource for Kubernetes, you
have to deploy Kubernetes on AWS cloud platform.
2. if you want to use Cinder, then you have to make sure Kubernetes is
deployed on OpenStack.
3. if you want to deploy Kubernetes on bare metals, right now you can only
choose SCSI or RBD device (maybe one or two more).

If you show it to users, they would get confused and ask why you can not
provide Cinder storage for Kubernetes deployed on bare metals. Since some
storage systems like Cinder, ScaleIO can provide volume resource for bare metals
directly, how can users use these resources if they deploy their cluster on
these bare metals?

Another use case is that we got some suggestions from users that we enrich
Kubernetes storage drivers since they were using some storage systems Kubernetes
doesn't support. Right now we have an option, which is flexvolume. But it would
be a hard work for them to develop every storage driver because the operation of
attaching volume to host is quite complicated. Can we shield the complex
implementation and provide a simple interface for their secondary development?

All in all, our motivation is to break the limitation between storage systems
and Kubernetes deployment environment, and to provide any storage backend
regardless of deployment environment(cloud platforms and bare metals).

## Goal

To solve the problem, we plan to create a standalone project in Kubernetes
that acts a library providing volume discovery and local management. For
any in-tree volume plugins which want to provide storage resource for
bare metals, they can call this library to finish host-side volume discovery
and then mount the device path to container. And the reason why set it up
as a standalone project is that in this way it can also serve for CSI plugin
and even any other storage systems written in Go.

## Proposed Design

As we know, there are a lot of storage protocols, such as iscsi, rbd, fc,
smbfs and so forth, and some of them are implemented in different ways
according to different system types(x86, s390, ppc64) and os types
(linux, windows), so it is quite a complicated work if we add these device
drivers directly into volume plugins. But what we can do is to create a
library to communicate with kernel and expose a unified interface to
volume plugins.

### API Object

The `Host-side Volume Discovery and Management Library` API object will
have the following structure:

```go
const (
	// Platform type
	PLATFORM_ALL = 'ALL'
	PLATFORM_x86 = 'X86'
	PLATFORM_S390 = 'S390'
	PLATFORM_PPC64 = 'PPC64'
	
	// Operation system type
	OS_TYPE_ALL = 'ALL'
	OS_TYPE_LINUX = 'LINUX'
	OS_TYPE_WINDOWS = 'WIN'

	// Device driver type
	ISCSI = "ISCSI"
	ISER = "ISER"
	FIBRE_CHANNEL = "FIBRE_CHANNEL"
	AOE = "AOE"
	DRBD = "DRBD"
	NFS = "NFS"
	GLUSTERFS = "GLUSTERFS"
	LOCAL = "LOCAL"
	GPFS = "GPFS"
	HUAWEISDSHYPERVISOR = "HUAWEISDSHYPERVISOR"
	HGST = "HGST"
	RBD = "RBD"
	SCALEIO = "SCALEIO"
	SCALITY = "SCALITY"
	QUOBYTE = "QUOBYTE"
	DISCO = "DISCO"
	VZSTORAGE = "VZSTORAGE"
	
	// A unified device path prefix
	VOLUME_LINK_DIR = '/dev/disk/by-id/'
)

// Connector is an interface indicating what outside world can do with this
// library, notice that it is at very early stage right now.
type Connector interface {
	GetConnectorProperties(multiPath, doLocalAttach bool) (*ConnectorProperties, error)
	
	ConnectVolume(conn *ConnectionInfo) (string, error)
	
	DisconnectVolume(conn *ConnectionInfo) (string, error)
	
	GetDevicePath(volumeId string) (string, error)
}

// ConnectorProperties is a struct used to tell storage backend how to
// intialize connection of volume. Please notice that it is OPTIONAL.
type ConnectorProperties struct {
	DoLocalAttach bool   `json:"doLocalAttach"`
	Platform      string `json:"platform"`
	OsType        string `json:"osType"`
	Ip            string `json:"ip"`
	Host          string `json:"host"`
	MultiPath     bool   `json:"multipath"`
	Initiator     string `json:"initiator"`
}

// ConnectionInfo is a structure for all properties of
// connection when connect a volume
type ConnectionInfo struct {
	// the type of driver volume, such as iscsi, rbd and so on
	DriverVolumeType string `json:"driverVolumeType"`
	
	// Required parameters to connect volume and differs from DriverVolumeType.
	// For example, for iscsi driver, see struct IsciConnectionData below.
	// NOTICE that you have to convert it into a map.
	ConnectionData   map[string]interface{} `json:"data"`
}

type IscsiConnectionData struct {
	// boolean indicating whether discovery was used
	TragetDiscovered bool 	`json:"targetDiscovered"`
	
	// the IQN of the iSCSI target
	TargetIqn string `json:"targetIqn"`
	
	// the portal of the iSCSI target
	TargetPortal string `json:"targetPortal"`
	
	// the lun of the iSCSI target
	TargetLun string `json:"targetLun"`
	
	// the uuid of the volume
	VolumeId string `json:"volumeId"`
	
	// the authentication details
	AuthUsername string `json:"authUsername"`
	AuthPassword string `json:"authPassword"`
}

```

## References

- For more information, please refer to https://github.com/openstack/os-brick,
which is an implementation written in Python.