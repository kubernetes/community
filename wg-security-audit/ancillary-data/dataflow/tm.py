# !/usr/bin/env python3

from pytm.pytm import TM, Server, Datastore, Dataflow, Boundary, Actor, Lambda, Process

tm = TM("Kubernetes Threat Model")
tm.description = "a deep-dive threat model of Kubernetes"

# Boundaries

inet = Boundary("Internet")
mcdata = Boundary("Master Control Data")
apisrv = Boundary("API Server")
mcomps = Boundary("Master Control Components")
worker = Boundary("Worker")
contain = Boundary("Container")

# Actors 

miu = Actor("Malicious Internal User")
ia = Actor("Internal Attacker")
ea = Actor("External Actor")
admin = Actor("Administrator")
dev = Actor("Developer")
eu = Actor("End User")

# Server & OS Components

etcd = Datastore("N-ary etcd servers")
apiserver = Server("kube-apiserver")
kubelet = Server("kubelet")
kubeproxy = Server("kube-proxy")
scheduler = Server("kube-scheduler")
controllers = Server("CCM/KCM")
pods = Server("Pods")
iptables = Process("iptables")

# Component <> Boundary Relations
etcd.inBoundary = mcdata
mcdata.inBoundary = apisrv
apiserver.inBoundary = apisrv
kubelet.inBoundary = worker
kubeproxy.inBoundary = worker
pods.inBoundary = contain
scheduler.inBoundary = mcomps
controllers.inBoundary = mcomps
pods.inBoundary = contain
iptables.inBoundary = worker
miu.inBoundary = apisrv
ia.inBoundary = contain
ea.inBoundary = inet
admin.inBoundary = apisrv
dev.inBoundary = inet
eu.inBoundary = inet

# Dataflows

apiserver2etcd = Dataflow(apiserver, etcd, "All kube-apiserver data")
apiserver2etcd.isEncrypted = True
apiserver2etcd.protocol = "HTTPS"

apiserver2kubelet = Dataflow(apiserver, kubelet, "kubelet Health, Status, &amp;c.")
apiserver2kubelet.isEncrypted = False
apiserver2kubelet.protocol = "HTTP"

apiserver2kubeproxy = Dataflow(apiserver, kubeproxy, "kube-proxy Health, Status, &amp;c.")
apiserver2kubeproxy.isEncrypted = False
apiserver2kubeproxy.protocol = "HTTP"

apiserver2scheduler = Dataflow(apiserver, scheduler, "kube-scheduler Health, Status, &amp;c.")
apiserver2scheduler.isEncrypted = False
apiserver2scheduler.protocol = "HTTP"

apiserver2controllers = Dataflow(apiserver, controllers, "{kube, cloud}-controller-manager Health, Status, &amp;c.")
apiserver2controllers.isEncrypted = False
apiserver2controllers.protocol = "HTTP"

kubelet2apiserver = Dataflow(kubelet, apiserver, "HTTP watch for resources on kube-apiserver")
kubelet2apiserver.isEncrypted = True
kubelet2apiserver.protocol = "HTTPS"

kubeproxy2apiserver = Dataflow(kubeproxy, apiserver, "HTTP watch for resources on kube-apiserver")
kubeproxy2apiserver.isEncrypted = True
kubeproxy2apiserver.protocol = "HTTPS"

controllers2apiserver = Dataflow(controllers, apiserver, "HTTP watch for resources on kube-apiserver")
controllers2apiserver.isEncrypted = True
controllers2apiserver.protocol = "HTTPS"

scheduler2apiserver = Dataflow(scheduler, apiserver, "HTTP watch for resources on kube-apiserver")
scheduler2apiserver.isEncrypted = True
scheduler2apiserver.protocol = "HTTPS"

kubelet2iptables = Dataflow(kubelet, iptables, "kubenet update of iptables (... ipvs, &amp;c) to setup Host-level ports")
kubelet2iptables.protocol = "IPC"

kubeproxy2iptables = Dataflow(kubeproxy, iptables, "kube-prxy update of iptables (... ipvs, &amp;c) to setup all pod networking")
kubeproxy2iptables.protocol = "IPC"

kubelet2pods = Dataflow(kubelet, pods, "kubelet to pod/CRI runtime, to spin up pods within a host")
kubelet2pods.protocol = "IPC"

eu2pods = Dataflow(eu, pods, "End-user access of Kubernetes-hosted applications")
ea2pods = Dataflow(ea, pods, "External Attacker attempting to compromise a trust boundary")
ia2cnts = Dataflow(ia, pods, "Internal Attacker with access to a compromised or malicious pod")

tm.process()
