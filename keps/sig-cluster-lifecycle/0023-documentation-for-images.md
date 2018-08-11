

# Documentation for images

Open https://www.websequencediagrams.com/ and paste the spec for the desired image:

- [kubeadm init](#kubeadm-init)
- [kubeadm join (and join --control-plane)](#kubeadm-join-and-join---control-plane)
- [kubeadm reset](#kubeadm-reset)
- [kubeadm upgrade](#kubeadm-upgrade)
- [kubeadm upgrade node](#kubeadm-upgrade-node)

## kubeadm init

```
title kubeadm init (interactions with the v1beta1 configuration)

participant "user" as u 
participant "kubeadm" as k
participant "kubelet" as kk
participant "node\n(api object)" as n
participant "kubeadm-config\nConfigMap" as cm
participant "kubeproxy-config\nConfigMap" as kpcm 
participant "kubelet-config\nConfigMap-1.*" as kcm

u->k:provide\nInitConfiguration (with NodeRegistrationOptions, ControlPlaneConfiguration)\nClusterConfiguration\nkube-proxy component configuration\nkubelet component configuration 

k->kk:write kubelet component configuration\nto /var/lib/kubelet/config.yaml
k->kk:write NodeRegistrationOptions\nto /var/lib/kubelet/kubeadm-flags.env
kk->n:start node
k->n:save NodeRegistrationOptions.CRISocket\nto kubeadm.alpha.kubernetes.io/cri-socket annotation 

k->k:use InitConfiguration\n(e.g. tokens)

k->cm:save ClusterConfiguration
k->cm:add Current ControlPlaneConfiguration to ClusterConfiguration.Status

k->kpcm:save kube-proxy component configuration 
k->kcm:save kubelet component configuration 
```

## kubeadm join (and join --control-plane)

```
title kubeadm join and join --control-plane (interactions with the v1beta1 configuration)

participant "user" as u 
participant "kubeadm" as k
participant "kubeadm-config\nConfigMap" as cm
participant "kubelet-config\nConfigMap-1.*" as kcm
participant "kubelet" as kk
participant "node\n(api object)" as n

u->k:provide\nJoinConfiguration\n(with NodeRegistrationOptions) 

k->cm:read ClusterConfiguration
cm->k:
k->k:use ClusterConfiguration\n(e.g. ClusterName)

k->kcm:read kubelet\ncomponent configuration
kcm->k:
k->kk:write kubelet component configuration\nto /var/lib/kubelet/config.yaml
k->kk:write NodeRegistrationOptions\nto /var/lib/kubelet/kubeadm-flags.env
kk->n:start node
k->n:save NodeRegistrationOptions.CRISocket\nto kubeadm.alpha.kubernetes.io/cri-socket annotation 

k->cm:add new ControlPlaneConfiguration\nto ClusterConfiguration.Status\n(only for join --control-plane)
```

## kubeadm reset 

```
title kubeadm reset (interactions with the v1beta1 configuration)

participant "user" as u 
participant "kubeadm" as k
participant "kubeadm-config\nConfigMap" as cm
participant "node\n(api object)" as n


u->k: 

k->cm:read ClusterConfiguration
cm->k:
k->cm:remove ControlPlaneConfiguration\nfrom ClusterConfiguration.Status\n(only if the node hosts a control plane instance)

k->n:read kubeadm.alpha.kubernetes.io/cri-socket annotation 
n->k:
k->k:use CRIsocket\nto delete containers
```

## kubeadm upgrade

```
title kubeadm upgrade apply (interactions with the v1beta1 configuration)

participant "user" as u 
participant "kubeadm" as k
participant "kubeadm-config\nConfigMap" as cm
participant "kubeproxy-config\nConfigMap" as kpcm 
participant "kubelet-config\nConfigMap-1.*+1" as kcm
participant "kubelet" as kk
participant "node\n(api object)" as n

u->k: UpgradeConfiguration
note over u, n:Upgrade configuration should allow only well known changes to the cluster e.g. the change of custom images if used


k->cm:read ClusterConfiguration
cm->k:
k->k:update\nClusterConfiguration\nusing api machinery
k->cm:save updated ClusterConfiguration

k->kpcm:read kube-proxy component configuration
kpcm->k:
k->k:update kube-proxy\ncomponent configuration\nusing api machinery
k->kpcm:save updated kube-proxy component configuration
note over kpcm, n:the updated kube-proxy component configuration will\nbe used by the updated kube-proxy DaemonSet 

k->kcm:read kubelet component configuration
kcm->k:
k->k:update kubelet\ncomponent configuration\nusing api machinery
k->kcm:save updated kubelet component configuration 
k->kk:write kubelet component configuration\nto /var/lib/kubelet/config.yaml
k->kk:write NodeRegistrationOptions\nto /var/lib/kubelet/kubeadm-flags.env
kk->n:start node

note over kcm, n:the updated kubelet component configuration\nwill be used by other nodes\nwhen running\nkubeadm upgrade nodes locally 

```

## kubeadm upgrade node

```
title kubeadm upgrade node (interactions with the v1beta1 configuration)

participant "user" as u 
participant "kubeadm" as k
participant "kubelet-config\nConfigMap-1.*" as kcm
participant "kubelet" as kk

u->k: 

k->kcm:read kubelet\ncomponent configuration
kcm->k:
k->kk:write kubelet component configuration\nto /var/lib/kubelet/config.yaml
```

