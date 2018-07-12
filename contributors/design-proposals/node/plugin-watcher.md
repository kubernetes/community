
# Plugin Watcher Utility

## Background

Portability and extendability are the major goals of Kubernetes from its beginning and we have seen more plugin mechanisms developed on Kubernetes to further improve them. Moving in this direction, Kubelet is starting to support pluggable [device exporting](https://github.com/kubernetes/community/blob/master/contributors/design-proposals/resource-management/device-plugin.md) and [CSI volume plugins](https://github.com/kubernetes/community/blob/master/contributors/design-proposals/storage/container-storage-interface.md). We are seeing the need for a common Kubelet plugin discovery model that can be used by different types of node-level plugins, such as device plugins, CSI, and CNI, to establish communication channels with Kubelet. This document lists two possible approaches of implementing this common Kubelet plugin discovery model. We are hoping to discuss these proposals with the OSS community to gather consensus on which model we would like to take forward.


## General Requirements

The primary goal of the Kubelet plugin discovery model is to provide a common mechanism for users to dynamically deploy vendor specific plugins that make different types of devices, or storage system, or network components available on a Kubernetes node. 

Here are the general requirements to consider when designing this system:

*   Security/authentication requirements
*   Underlying communication channel to use: stay with gRPC v.s. flexibility to support multiple communication protocol
*   How to detect API version mismatching
*   Ping-pong plugin registration
*   Failure recovery or upgrade story upon kubelet restart and/or plugin restart
*   How to prevent single misbehaving plugin from flooding kubelet
*   How to de-registration
*   Need to support some existing protocol that is not bound to K8s, like CSI

## Proposed Models

#### Model 1: plugin registers with Kubelet through grpc (currently used in device plugin)

*   Currently requires plugin to run with privilege and communicate to kubelet through unix socket under a canonical directory, but has flexibility to support different communication channels or authentication methods.
*   API version mismatch is detected during registration.
*   Currently always take newest plugin upon re-registration. Can implement some policy to reject plugin re-registration if a plugin re-registers too frequently. Can terminate the communication channel if a plugin sends too many updates to Kubelet.
*   In the current implementation, kubelet removes all of the device plugin unix sockets. Device plugins are expected to watch for such event and re-register with the new kubelet instance. The solution is a bit ad-hoc. There is also a temporary period that we can't schedule new pods requiring device plugin resource on the node after kubelet restart, till the corresponding device plugin re-registers. This temporary period can be avoided if we also checkpoints device plugin socket information on Kubelet side. Pods previously scheduled can continue with device plugin allocation information already recorded in a checkpoint file. Checkpointing plugin socket information is easier to be added in DevicePlugins that already maintains a checkpoint file for other purposes. This however could be a new requirement for other plugin systems like CSI. 

![image](https://user-images.githubusercontent.com/11936386/42627970-3bf055a4-85ec-11e8-93cb-f4f393b2bd76.png)

#### Model 2: Kubelet watches new plugins under a canonical path through inotify (Preferred one and current implementation)

*   Plugin can export a registration rpc for API version checking or further authentication. Kubelet doesn't need to export a rpc service.
*   We will take gRPC as the single supported communication channel.
*   Can take the newest plugin from the latest inotify creation. May require socket name to follow certain naming convention (e.g., resourceName.timestamp) to detect ping-pong plugin registration, and ignore socket creations from a plugin if it creates too many sockets during a short period of time. We can even require that the resource name embedded in the socket path to be part of the identification process, e.g., a plugin at `/var/lib/kubelet/plugins/resourceName.timestamp` must identify itself as resourceName or it will be rejected.
*   Easy to avoid temporary plugin unavailability after kubelet restart. Kubelet just needs to scan through the special directory. It can remove plugin sockets that fail to respond, and always take the last live socket when multiple registrations happen with the same plugin name. This simplifies device plugin implementation because they don't need to detect Kubelet restarts and re-register.
*   A plugin should remove its socket upon termination to avoid leaving dead sockets in the canonical path, although this is not strictly required.
*   CSI needs flexibility to not only bound to Kubernetes. With probe model, may need to add an interface for K8s to get plugin information.
*   We can introduce special plugin pod for which we automatically setup its environment to communicate with kubelet. Even if Kubelet runs in a container, it is easy to config the communication path between plugin and Kubelet.
![image](https://user-images.githubusercontent.com/11936386/42628430-be3ab5bc-85ed-11e8-93ac-173511fdd39a.png)

**More Implementation Details on Model 2:**

*   Kubelet will have a new module, PluginWatcher, which will probe a canonical path recursively
*   On detecting a socket creation, Watcher will try to get plugin identity details using a gRPC client on the discovered socket and the RPCs of a newly introduced `Identity` service.
*   Plugins must implement `Identity` service RPCs for initial communication with Watcher.

**Identity Service Primitives:**
```golang

// PluginInfo is the message sent from a plugin to the Kubelet pluginwatcher for plugin registration
message PluginInfo {
	// Type of the Plugin. CSIPlugin or DevicePlugin
	string type = 1;
	// Plugin name that uniquely identifies the plugin for the given plugin type.
	// For DevicePlugin, this is the resource name that the plugin manages and
	// should follow the extended resource name convention.
	// For CSI, this is the CSI driver registrar name.
	string name = 2;
        // Optional endpoint location. If found set by Kubelet component,
        // Kubelet component will use this endpoint for specific requests.
        // This allows the plugin to register using one endpoint and possibly use
        // a different socket for control operations. CSI uses this model to delegate
        // its registration external from the plugin.
        string endpoint = 3;
	// Plugin service API versions the plugin supports.
	// For DevicePlugin, this maps to the deviceplugin API versions the
	// plugin supports at the given socket.
	// The Kubelet component communicating with the plugin should be able
	// to choose any preferred version from this list, or returns an error
	// if none of the listed versions is supported.
	repeated string supported_versions = 4;
}

// RegistrationStatus is the message sent from Kubelet pluginwatcher to the plugin for notification on registration status
message RegistrationStatus {
	// True if plugin gets registered successfully at Kubelet
	bool plugin_registered  = 1;
	// Error message in case plugin fails to register, empty string otherwise
	string error  = 2;
}

// RegistrationStatusResponse is sent by plugin to kubelet in response to RegistrationStatus RPC
message RegistrationStatusResponse {
}

// InfoRequest is the empty request message from Kubelet
message InfoRequest {
}

// Registration is the service advertised by the Plugins.
service Registration {
	rpc GetInfo(InfoRequest) returns (PluginInfo) {}
	rpc NotifyRegistrationStatus(RegistrationStatus) returns (RegistrationStatusResponse) {}
}
```

**PluginWatcher primitives:**
```golang
// Watcher is the plugin watcher
type Watcher struct {
	path      string
	handlers  map[string]RegisterCallbackFn
	stopCh    chan interface{}
	fs        utilfs.Filesystem
	fsWatcher *fsnotify.Watcher
	wg        sync.WaitGroup
	mutex     sync.Mutex
}

// RegisterCbkFn is the type of the callback function that handlers will provide
type RegisterCallbackFn func(pluginName string, endpoint string, versions []string, socketPath string) (chan bool, error)

// AddHandler registers a callback to be invoked for a particular type of plugin
func (w *Watcher) AddHandler(pluginType string, handlerCbkFn RegisterCbkFn) {
        w.handlers[handlerType] = handlerCbkFn
}

// Start watches for the creation of plugin sockets at the path
func (w *Watcher) Start() error {

// Probes on the canonical path for socket creations in a forever loop

// For any new socket creation, invokes `Info()` at plugins Identity service
resp, err := client.Info(context.Background(), &watcherapi.Empty{})

// Keeps the connection open and passes plugin's identity details, along with socket path to the handler using callback function registered by handler. Handler callback is selected based on the Type of the plugin, for example device plugin or CSI plugin
// Handler Callback is supposed to authenticate the plugin details and if all correct, register the Plugin at the kubelet subsystem. 

if handlerCbkFn, ok := w.handlers[resp.Type]; ok {
                                                err = handlerCbkFn(resp, event.Name)
...
}

// After Callback returns, PluginWatcher notifies back status to the plugin

client.NotifyRegistrationStatus(ctx, &registerapi.RegistrationStatus{
...
})

```


**How any Kubelet sub-module can use PluginWatcher:**



*   There must be a callback function defined in the sub-module of the signature: 

```golang
type RegisterCallbackFn func(pluginName string, endpoint string, versions []string, socketPath string) (chan bool, error)
```
*   Just after sub-module start, this callback should be registered with the PluginWatcher, eg:
```golang
kl.pluginWatcher.AddHandler(pluginwatcherapi.DevicePlugin, kl.containerManager.GetPluginRegistrationHandlerCbkFunc())
```

**Open issues (Points from the meeting notes for the record):**
*   Discuss with security team if this is a viable approach (and if cert auth can be added on top for added security).
*   Plugin author should be able to write yaml once, so the plugin dir should not be hard coded. 3 options:
    *   Downward API param for plugin directory that will be used as hostpath src 
    *   A new volume plugin that can be used by plugin to drop a socket
    *   Have plugins call kubelet -- link local interface
        *   Bigger change -- kubelet doesn't do this
        *   Path of most resistance
