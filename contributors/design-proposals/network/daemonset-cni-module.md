# Move Kubelet CNI built-in module to a DaemonSet process

@Lion-Wei, @m1093782566

Last update: 27th Feb. 2018

Status: Draft

# Background & Motivations

CNI drivers are independent binaries and CNI API is understood by a kubelet built-in module, which is linked, compiled, built, and shipped with kubelet. Third party network plugins have to provide an executable CNI plugin binary and a network configuration file. This is undesirable for many reasons:

1. CNI module is coupled and dependent on Kubernetes releases.
2. Added bunch of CNI and network related flags to kubelet: cni-bin-dir/cni-conf-dir/hairpin-mode/network-plugin/network-plugin-mtu,etc.
3. Network plugins require write access to the root filesystem of the nodes in order to provide CNI plugin binary and network configuration file.
4. We should find a way to make an IPAM driver that doesn't take the subnet as a literal config value, so we can use one config more all nodes.
5. Vector to move network plugins and service configurators (kube-proxy) closer together for more optimal setup .
6. Move some more net config out of kubelet (implied APIs via iptables are bad)
7. configmaps as an alternative to editing files on nodes
8. extensibility beyond just CNI.
9. There are places where kube-proxy could do better configuration if if had access to information that only the CNI drivers know. For example, the name that veth interfaces are given or the correct CIDRs for pods on that node.
10. Kube-proxy(service driver) is keeping doing so many low-level things, it would be better if the concepts of network driver and service "driver" should be closer to each other.
11. use kube mechanisms to manage plugins (consistent with most other plugins)


Beyond just the CNI plugins, Kubernetes implements Services, with kube-proxy being the default and most prevalent implementation. Both CNI plugins and kube-proxy are manipulating aspects of the network configuration. There are assumptions that kube-proxy makes about how networking works, which may not be true if the CNI plugins are changed. There are assumptions that kube-proxy COULD make if it had more intimate knowledge of the CNI drivers in use.

As we talked in kubernetes-sig-network, topic “Combining CNI + service”, it was proposed to build a DaemonSet responsible for container network and “service network”. Either combine CNI driver with kube-proxy or make CNI driver work with third-party provided network plugin. This is a long term proposal, need further discussion.

For now, what we can do is to migrate CNI driver out of kubelet, and run CNI driver in a DaemonSet, which can translate kubelet’s network GRPCs to CNI calls. The primary motivation to do this is to decouple CNI from kubelet. In addition, for network vendors, they can combine network plugin with CNI plugin in one DaemonSet.


## Objective

The objective of this document is to document all the requirements for migrating CNI driver from kubelet to a DaemonSet.


## Goals
- Define mechanism by which Kubernetes master and node components will communicate with CNI driver.
- Define yaml file that describes CNI driver DaemonSet.
- Recommend network plugin DaemonSet struct for Kubernetes compatible, third-party network plugin.



## Non-Goals
- Define interface by which thirdparty network plugin interacting withKubernetes.
- Define mechanism by which CNI driver communicate with kube-proxy orthirty-party network plugin.


# Design Overview

A new in-tree CNI network plugin will be introduced in Kubernetes. This new plugin will be the mechanism by which kubelet interact with external CNI drivers. The calls for new in-tree CNI network plugin will be convert to corresponding RPCs through a unix domain socket on the node machine.

A daemon process will be offered to capture all the RPCs from kubelet and handle it or call CNI plugin binaries.

## Design Details

#### Kubernetes In-tree remote CNI network plugin

Build a new `RemoteCNINetworkPlugin` object, which will realize all the interface `NetworkPlugin` required, this will enable kubelet use the new network plugin **just like CNI and kubenet**. And once the CNI plugin DaemonSet is stable and function as the the old in-tree cni network plugin, the old one will be deprecated.

The `RemoteCNINetworkPlugin` object will include a gRPC client named `cniClient`, the kubelet invokes for `RemoteCNINetworkPlugin` will be convert to corresponding RPCs calls by `cniClient`.

```go
// RemoteCNINetworkPlugin is a GRPC implementation of NetworkPlugin
type RemoteCNINetworkPlugin struct {

  // host is an interface that plugin can use to access the kubelet
  // plugin will use host to get network namespace information
  // of sandbox
  host      network.Host

  // cniClient is the client for CNIDriver service
  // plugin will convert kubelet call to invokes of this client
  cniClient cniClient
}

func (plugin *RemoteCNINetworkPlugin) SetUpPod{
  // *1 Get sandbox network namespace
  plugin.host.GetNetNS

  // *2 Get portMap
  plugin.host.GetPodPortMapping

  // *3 build runtime config

  // *4 Call cniClient
  // cniClient.setUpPod will call DaemonSet CNI driver
  // to add CNI network
  plugin.cniClient.setUpPod
}

func (plugin *RemoteCNINetworkPlugin) TearDownPod{
  // *1 Get sandbox network namespace
  plugiun.host.GetNetNS

  // *2 Get portMap
  plugin.host.GetPodPortMapping

  // *3 build runtime config

  // *4 Call cniClient
  // cniClient.tearDownPod will call DaemonSet CNI driver
  // to delete CNI network
  plugin.cniClient.tearDownPod
}

func (plugin *RemoteCNINetworkPlugin) Status{
  // call cniClient
  // cniCLient.cniStatus will call DaemonSet CNI driver
  // to sync CNI network config and check whether
  // CNI driver initialized
  plugin.cniClient.cniStatus
}
```

#### DaemonSet CNI driver program

We will build a new repo for out-tree CNI driver, maybe named `kubernetes-cni`, which will run as a DeamonSet program, and offer a gRPC server that in-tree network plugin can invoke. This CNI driver will acts as a bridge between third-party CNI plugin binaries and kubelet.

And the DaemonSet CNI driver will use a configmap as the network config file, so don’t have to access the root filesystem of the node or master.

```go
type cniDriver struct {
  // cni.NetworkConfig is the network config specified by
  // containernetworking/cni
  // CNI driver will build this config according to network
  // information from configmap
  networkConfig cni.NetworkConfig

  // cni.RuntimeConf is the runtime config specified by
  // containernetworking/cni
  // CNI driver will build this config according to runtime config
  // from kubelet
  runtimeConfig cni.RuntimeConf

  // cni.CNI is the interface specified by containernetworking/cni
  // cniDriver will use this interface to Add/Del network
  cniConfig     cni.CNI

  // cniPluginDir is the path of third-party CNI plugin binaries
  // which from environment variables of DaemonSet
  cniPluginDir     []string

  // networkConfigDir is the network config path
  // which from environment variables of DaemonSet
  networkConfigDir string
}


func (cni *cniDriver) init {
  // sync network config
  cni.syncNetworkConfig

  // check CNI plugin binaries

  // periodically sync network config

  // register cni driver server
  cni.registerCNIDriverServer

  // mark initialized
}

// cniServer are the gRPC servers that in-tree plugin can invoke
func (cs *cniServer) setUpPod{}
func (cs *cniServer) tearDownPod{}
func (cs *cniServer) cniStatus{}
```

## Recommended Mechanism for Deploying CNI Drivers on Kubernetes

Kubernetes will offer kubernetes-cni as a CNI driver running as a DaemonSet program, and we can have the following recommendations to simplify the deployment.

To deploy an DaemonSet CNI driver with CNI plugin, it is recommended that network vendors:

- Package kubernetes team offered out-tree-cni-driver and and cni-plugin-binaries into one image.
- Offer a yaml file of DaemonSet with informations that out-tree-cni-driver needed.
  - Container path of CNI-plugin-binaries.
  - Mount the configmap of network-config to container and offer it’s path.
- Create a configmap of cni plugin config.

## Example Walkthrough

#### Out-tree CNI driver init

1. Third-party network plugin create a configmap of network config, and put CNI plugin binaries in the right position.
2. CNI driver read network config file and make sure can access the CNI plugin binaries.
3. CNI driver register all gRPC server and listen for connections.

#### In-tree network plugin init

1. Use RemoteCNINetworkPlugin as network plugin.
2. Call cniClient for CNI driver status.
3. Out-tree cni driver return CNI driver and network information.

#### Setup pod

1. In-tree cni network plugin get network namespace.
2. Call cniClient to create pod sandbox network.
3. Out-tree cni driver build network config and runtime config, call cni plugin binaries.
4. Out-tree CNI driver return PodNetworkStatus.

#### Teardown pod

1. In-tree cni network plugin get network namespace.
2. Call cniClient to delete pod sandbox network.
3. Out-tree cni driver build network config and runtime config, call cni plugin binaries.
4. Out-tree CNI driver return delete pod sandbox result.

## Examples

#### Configmap of network configure

```json
{
  "name": "pod-network",
  "cniVersion": "0.3.0",: [
    {
      "typse": "ptp",
      "ipam": {
        "type": "host-local",
        "subnet": "10.64.1.0/24",
	"routes": [
	  {
	    "dst": "0.0.0.0/0"
	  }
	]
      }
    },
    {
      "type": "portmap",
      "capabilities": {"portMappings": true},
      "noSnat": true
    }
  ]
}
```

#### DaemonSet file

```yaml
kind: DaemonSet
apiVersion: extensions/v1beta1
metadata:
  name: kubernetes-cni
  namespace: kube-system
spec:
  name: kubernetes-cni
  namespace: kube-system
  template:
    spec:
      containers:
        - name: kubernetes-cni
          image: xxxxxxx
          env:
            # The network configlist container dir
            - name: CNI_CONF_DIR
              value: "/etc/cni/net.d"
            # The CNI plugin container dir
            - name: CNI_PLUGIN_DIR
              value: "/host/opt/cni/plugin"
          volumeMounts:
            - mountPath: /etc/cni/net.d
              name: cni-conf
              readOnly: false
            - mountPath: /host/opt/cni/plugin
              name: cni-plugin
              readOnly: false
      volumes:
        # cni plugin binaries path
        - hostPath:
            path: /opt/cni/plugin
          name: cni-plugin
        # a configmap of cni network cinfig list
        - configMap:
            defaultMode: 420
            name: cni-conf
          name: cni-conf

```

