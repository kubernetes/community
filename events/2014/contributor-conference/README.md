# Kubernetes Contributor Conference, 2014-12-03 to 12-05
**Full notes:** (Has pictures; Shared with k-dev mailing list) (https://docs.google.com/document/d/1cQLY9yeFgxlr_SRgaBZYGcJ4UtNhLAjJNwJa8424JMA/edit?usp=sharing)   
**Organizers:** thockin and bburns  
**26 Attendees from:** Google, Red Hat, CoreOS, Box  
**This is a historical document. No typo or grammar correction PRs needed.**

Last modified: Dec. 8. 2014

# Clustering and Cluster Formation
Goal: Decide how clusters should be formed and resized over time
Models for building clusters  
* Master in charge - asset DB  
Dynamic join - ask to join  
* How Kelsey Hightower has seen this done on bare metal  
Use Fleet as a machine database  
A Fleet agent is run on each node  
Each node registers its information in etcd when it comes up  
Only security is that etcd expects the node to have a cert signed by a specific CA  
Run an etcd proxy on each node  
Don't run any salt scripts, everything is declarative  
Just put a daemon (kube-register) on a machine to become part of the cluster  
brendanburns: basically using Fleet as the cloud provider  
* Puppet model - whitelist some cert and/or subnet that you want to trust everything in  
One problem - if CA leaks, have to replace certs on all nodes  
* briangrant: we may want to support adding nodes that aren't trusted, only scheduling work from the nodes' owner on them  
* lavalamp: we need to differentiate between node states:  
In the cluster  
Ready to accept work  
Trusted to accept work  
* Proposal:  
New nodes initiate contact with the master  
Allow multiple config options for how trust can be established - IP, cert, etc.  
Each new node only needs one piece of information - how to find the master  
Can support many different auth modes - let anyone in, whitelist IPs, a particular signed cert, queue up requests for an admin to approve, etc.  
Default should be auto-register with no auth/approval needed
Auth-ing is separate from registering  
Supporting switching between permissive and strict auth modes:  
Each node should register a public key such that if the auth mode is changed to require a cert upon registration, old nodes won't break  
kelseyhightower: let the minion do the same thing that kube-register currently does  
Separate adding a node to the cluster from declaring it as schedulable  
* Use cases:  
Kick the tires, everything should be automagic  
Professional that needs security  
* Working group for later: Joe, Kelsey, Quintin, Eric Paris  
# Usability  
* Getting started  
Want easy entry for Docker users  
Library/registry of pod templates  
* GUI - visualization of relationships and dependencies, workflows, dashboards, ways to learn, first impressions  
Will be easiest to start with a read-only UI before worrying about read-write workflows  
* Docs  
Need to refactor getting started guides so that there's one common guide  
Each cloud provider will just have its own short guide on how to create a cluster  
Need a simple test that can verify whether your cluster is healthy or diagnose why it isn't  
Make it easier to get to architecture/design doc from front page of github project  
Table of contents for docs?  
Realistic examples  
Kelsey has found that doing a tutorial of deploying with a canary   helped make the value of labels clear  
* CLI  
Annoying when local auth files and config get overwritten when trying to work with multiple clusters  
Like when running e2e tests  
* Common friction points  
External IPs  
Image registry  
Secrets  
Deployment  
Stateful services  
Scheduling  
Events/status  
Log access  
* Working groups  
GUI - Jordan, Brian, Max, Satnam  
CLI - Jeff, Sam, Derek  
Docs - Proppy, Kelsey, TJ, Satnam, Jeff  
Features/Experience - Dawn, Rohit, Kelsey, Proppy, Clayton:   https://docs.google.com/document/d/1hqn6FtBNMe0sThbciq2PbE_P5BONBgCzHST4gz2zj50/edit


# v1beta3 discussion

12-04-2014  
Network -- breakout  
* Dynamic IP  
Once we support live migration, IP assigned for each POD has to move together, which might be broken the underneath.   
We don’t have introspection, which makes supporting various network topology harder.  
External IP is an important part.  
There’s a kick-the-tires mode and full-on mode (for GCE, AWS - fully featured).  
How do we select kick-the-tires ? Weave, Flannel, Calico: pick one.   
Someone does a comparison. thockin@ would like help in evaluating these tech against some benchmarks. Eric Paris can help - has a bare-metal setup. We’ll have a benchmark setup for evaluation.  
We need to have two real use-cases at least - a webserver example; can 10 pods find each other. lavalamp@ working on a test.  
If docker picks up a plugin model, we can use that.  
Cluster will be dynamically change, we need to design a flexible network plugin API to accomplish this.    
Flannel two things: network allocation through etcd and traffic routing w/ overlays. Also programs underlay networks (like GCE).   Flannel will do IP allocation, not hard-coded.   
One special use case: per node, there are only 20 ips could be allocated. Scheduler might need to know the limitation: OUT-OF-IP(?)  
Different cloud providers, but OVS is a common mechanism  
We might need Network Grids at the end  
ACTIONS: better doc, test.   
* Public Services  
Hard problem: Have to scale to GCE, GCE load balancer cannot target to arbitrary IP, only can target to a VM for now.  
Until we have an external IP, you cannot build a HA public service.   
We can run Digital Ocean on top of kubernetes  
Issue: When starting a public service, there is internal IP assigned. It is accessable from node within cluster, but not from outside. Now we have a 3-tier services, how to access one service from outside The issue is how to take this internal accessible service externalized. General solution: forwarding the traffic outside to the internal IP. First action, teach kubernetes mapping.   
We need a registry of those public IPs. All traffic comes to that IP will be forwarded to proper IP internally.  
public service can register with DNS, and do a intermiddle load balancing outside cluster / kubernetes. Label query to tell the endpoint.   
K8s proxy can be L3 LB, and listen to the external IPs, it also talk to k8s service DB and find internal services; then goes to L7 LB, which could be HAP proxy, scheduled as a pod, it talks to Pods DB, find a cluster of pods to forward the traffic.  



Two types of services: mapping external IPs and L3 LB to map to pods. L7 LB can access the IPs assigned to pods.   
Policy: Add more nodes, more external IPs can be used.  
Issue1: how to take external IP to map to a list of pods, L3 LB part.   
Issue2: how to slice those external IPs: general pool vs. private pools.  
* IP-per-service, visibility, segmenting  
* Scale  
* MAC  

# Roadmap

* Should be driven by scenarios / use cases -- breakout  
* Storage / stateful services -- breakout  
Clustered databases / kv stores  
Mongo  
MySQL master/slave  
Cassandra  
etcd  
zookeeper  
redis  
ldap  
Alternatives  
local storage  
durable volumes  
identity associated with volumes  
lifecycle management  
network storage (ceph, nfs, gluster, hdfs)  
volume plugin  
flocker - volume migration  
“durable” data (as reliable as host)  
* Upgrading Kubernetes  
master components  
kubelets  
OS + kernel + Docker  
* Usability  
Easy cluster startup  
Minion registration  
Configuring k8s  
move away from flags in master  
node config distribution  
kubelet config  
dockercfg  
Cluster scaling  
CLI + config + deployment / rolling updates  
Selected workloads  
* Networking  
External IPs  
DNS  
Kick-the-tires networking implementation  
* Admission control not required for 1.0  
* v1 API + deprecation policy  
* Kubelet API well defined and versioned  
* Basic resource-aware scheduling -- breakout  
require limits?  
auto-sizing  
* Registry  
Predictable deployment (config-time image resolution)  
Easy code->k8s  
Simple out-of-the box setup  
One or many?  
Proxy?  
Service?  
Configurable .dockercfg  
* Productionization  
Scalability  
100 for 1.0  
1000 by summer 2015  
HA master -- not gating 1.0  
Master election  
Eliminate global in-memory state  
IP allocator  
Operations  
Sharding  
Pod getter  
Kubelets need to coast when master down  
Don’t blow away pods when master is down  
Testing  
More/better/easier E2E  
E2E integration testing w/ OpenShift  
More non-E2E integration tests  
Long-term soaking / stress test  
Backward compatibility  
Release cadence and artifacts  
Export monitoring metrics (instrumentation)  
Bounded disk space on master and kubelets  
GC of unused images  
* Docs  
Reference architecture  
* Auth[nz]  
plugins + policy
admin
user->master
master component->component: localhost in 1.0
kubelet->master
