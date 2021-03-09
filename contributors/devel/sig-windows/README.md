# Contributing to sig-windows

## We want your help !

Anything that is relevant to Kubernetes can effect windows as well.   We can train you and help you onboard if you have the time and drive to be a windows + Kubernetes contributor!

The first thing to do is join the  weekly sig-windows meetings, and come to our "sig-windows" after hours onboarding call, where we hack on various
windows ideas to help onboard newcomers.

## Come hang out with us!

The meeting details for sig-windows are here (every tuesday at 12:30 PM ET)

https://docs.google.com/document/d/1Tjxzjjuy4SQsFSUVXZbvqVb64hjNAG5CQX8bK7Yda9w/edit#heading=h.3f5dhus4q8i2

## Generic getting started guide

We have a generic getting started guide here: https://github.com/kubernetes/community/blob/master/sig-windows/CONTRIBUTING.md.

## Test grid

The Kubernetes test-grid needs continuos monitoring and triage.  If you are interested in helping us maintain this, 
check https://testgrid.k8s.io/sig-windows-signal, and contact any of the sig-windows leads.  This is a great place for begginers
to get started, as it involves reading through logs, looking for failures, and making small patches/telling contributors about
areas that need help.

## Developer recipes and automation 

Developing on windows takes alot of manual setup, and a good place to help is making sure we have automation available for spinning up 
windows VMs from source.  This is also a great place to get started as a begginer.

## E2Es and Conformance testing

The Kubernetes end to end test have tags for `sig-windows` and also Conformance tests which generally work on Windows.  Audting These tests
and cleaning them up, making sure there are clear to evaluate logs, and so on, is an area in constant need. 

## The Kubelet

The Kubelet on windows has a few different code paths then that of linux.  IF your interested in the Kubelet, the Windows Kubelet
code path is a great place to start.

## Networking

Theres lots of evolution in the networking landscape for windows:  

- CNI providers like antrea and calico are continually evolving their support for windows and containerd
- Lots of help and work is needed on the windows kube-proxy.

## Storage

The windows group dives into privileged containers (for windows) as well as maintains the csi-proxy, a way to run "privileged" storage containers
that implement CSI.  Reading about https://github.com/kubernetes-csi/csi-proxy, and looking at the in-tree windows storage provider code is a good place to start here.

## Priveliged containers and beyond

Windows is re-inventing many native linux concepts to enable linux-idioms in the Windows ecosystem.  Experimenting with topics like 
GMSA, priveliged containers, containerd integration, devices, and other advanced windows concepts might be a great place for a newcomer
with domain experience (i.e. an ActiveDirectory user without K8s experience) to get involved and start contributing. 
