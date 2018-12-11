### Scalability Testing/Analysis Environment and Goals

Project practice is to perform baseline scalability testing and analysis on a large single machine (VM or server) with all control plane processing on that single node. The single large machine provides sufficient scalability to scale to 5000 node density tests. The typical machine for testing at this scale is at the larger end of the VM scale available on public cloud providers, but is by no means the largest available. Large cluster runs are typically run with a node emulator (kubemark), with a set of resources to run kubemark typically requiring ~80 machines to simulate 5000 nodes, i.e. 60-ish hollow-nodes per machine.

The typical control plane server for this testing has 64 cores, at least 128GB of memory, Gig-E network interfaces, and SSD drives. Public cloud instances typically will have more memory than this for this number of cores. 

Several factors contribute to a need to expanded testing beyond the single node baseline.

* Very large cluster operators typically run a minimum of 5 servers in the control plane to ensure that a control plane failure during upgrade is survivable without losing cluster state. We want testing to represent typical configurations at this scale..

* RAFT consensus processing on etcd means that 5-server clusters have different performance characteristics than 1-server clusters.

* Distributed systems often show different performance characteristics when the components are separated by a network versus co-execution on the same server.

* Distributed systems are often affected by cache consistency issues.

An important attribute of Kubernetes is broad support for different infrastructure options. Project experience is that testing on a variety of infrastructure options flushes out timing issues and improves quality. Users of kubernetes also find value in knowing that scale testing has been performed on the infrastructure options they care about.

Scalability testing on configurations that are similar are expected to have similar results, and deviations from the expectation need attention.  Regressions may indicate system issues or undocumented assumptions based on those differences, and should be both explored and documented. Noting differences in various configs and which provide the highest system throughput may also give indications as to which performance optimizations are most interesting.

### Control Plane Machine Selection

As wide a selection of different infrastructure providers as possible helps the project. Configurations and testing strongly welcomed for providers not currently listed, and the Scalability SIG is engaging with all of the providers listed below.

<table>
  <tr>
    <td>Provider</td>
    <td>Machine type</td>
    <td>Cores</td>
    <td>Memory</td>
    <td>Kubemark Needs</td>
    <td>Notes</td>
  </tr>
  <tr>
    <td>Google</td>
    <td>n1-standard-64</td>
    <td>64</td>
    <td>240</td>
    <td>80 instances</td>
    <td>Used for 5000 node test results</td>
  </tr>
  <tr>
    <td>AWS</td>
    <td>m4.16xlarge</td>
    <td>64</td>
    <td>256</td>
    <td>80 instances</td>
    <td>Proposed</td>
  </tr>
  <tr>
    <td>Azure</td>
    <td>standard-g5</td>
    <td>32</td>
    <td>448</td>
    <td>?</td>
    <td>Max cores instance, 
proposed</td>
  </tr>
  <tr>
    <td>Packet</td>
    <td>Type 2</td>
    <td>24</td>
    <td>256</td>
    <td>?</td>
    <td>Bare metal, proposed</td>
  </tr>
</table>


### Additional Configuration Requirements

* Scalability SIG efforts are currently oriented towards the 1.6 and later releases. This focus will shift over time - the SIG’s efforts are aimed towards scalability on trunk.

* Configuration and tuning of etcd is a critical component and has dramatic effects on scalability of large clusters. Minimum etcd version is 3.1.8

* API server configured for load balancing, other components using standard leader election.

* Etcd is used for two distinct cluster purposes - cluster state and event processing. These have different i/o characteristics. It is important to scalability testing efforts that the iops provided by the servers to etcd be consistent and protected. These leads to two requirements:

    * Split etcd: Two different etcd clusters for events and cluster state (note: this is currently the GKE production default as well).

    * Dedicated separate IOPS for the etcd clusters on each control plane node. On a bare metal install this could look like a dedicated SSD. This requires a more specific configuration per provider.  Config table below.

<table>
  <tr>
    <td>Provider</td>
    <td>Volume type</td>
    <td>Size per etcd partition (2x)</td>
    <td>Notes</td>
  </tr>
  <tr>
    <td>Google</td>
    <td>SSD persistent disk</td>
    <td>256GB</td>
    <td>Iops increases with volume size</td>
  </tr>
  <tr>
    <td>AWS</td>
    <td>EBS Provisioned IOPS SSD (io1)</td>
    <td>256GB</td>
    <td>Proposed</td>
  </tr>
  <tr>
    <td>Azure</td>
    <td>?</td>
    <td>?</td>
    <td>
</td>
  </tr>
  <tr>
    <td>Packet</td>
    <td>Dedicated SSD</td>
    <td>256GB</td>
    <td>Bare metal, proposed</td>
  </tr>
</table>


### Areas for Future Work

* Leader election results are non-deterministic on a typical cluster, and a config would be best served to be configured as worst-case. Not presently known whether there are performance impacts resulting from leader election resulting in either co-location or distribution of those components.

* Improving the cluster performance loading to match production deployment scenarios is critical on-going work, especially clusterloader: [https://git.k8s.io/perf-tests/clusterloader](https://git.k8s.io/perf-tests/clusterloader)

* Multi-zone / multi-az deployments are often used to manage large clusters, but for testing/scalability efforts the target is intentionally a single Availability Zone. This keeps greater consistency between environments that do and don’t support AZ-based deployments. Failures during scalability testing are outside the SIG charter. Protecting against network partitioning and improving total cluster availability (one of the key benefits to a multi-AZ strategy) are currently out scope for the Scalability SIG efforts.

* Scalability issues on very large clusters of actual nodes (instead of kubemark simulations) are real. Efforts to improve large cluster networking performance e.g. IPVS are important, and will be interesting areas for cross-SIG collaboration.

### Control Plane Cluster Config 

Diagram shows high level target control plane config using the server types listed above, capturing:

* 5 server cluster

* Split etcd across 5 nodes

* API server load balanced

* Other components using leader election

Detail: Target config uses separate volumes for the etc configs:

Config hints to testers:
* ELB, by default, has a short timeout that'll cause control plane components to resync often. Users should set that to the max.

### Alternative Config

Motivated by many of the issues above, an alternative configuration is reasonable, and worth experimentation, as some production environments are built this way:

* Separate etcd cluster onto a dedicated set of 5 machines for etcd only.

* Do not run split etcd

* Run remainder of control plane on 5 nodes separately

* Question for discussion: are there advantages to this configuration for environments where the max number of cores per host are < 64?

### References

CoreOS commentary on etcd sizing.

[https://github.com/coreos/etcd/blob/master/Documentation/op-guide/hardware.md](https://github.com/coreos/etcd/blob/master/Documentation/op-guide/hardware.md)
