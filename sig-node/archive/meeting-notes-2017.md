# sig-node weekly meeting

# Dec 19 Proposed Agenda {#dec-19-proposed-agenda}



*   Cloud Providers and Kubelet
    *   Move cloud provider related code to separate repos
        *   breaking node controller manager into pieces
            *   Node lifecycle controller.
            *   IPAM (IP, CIDR).
    *   https://github.com/kubernetes/kubernetes/pull/55634
    *   https://github.com/kubernetes/kubernetes/pull/50811
*   Logistics
    *   Jan 2nd meeting?
    *   Switch to zoom next year
*   Next Year:
    *   Secure Container/Pod
    *   Node level debuggability
    *   Windows Container.
    *   Container Runtime Interface
        *   Container Monitoring.
        *   Container Logging.
        *   Secure Container/Pod.
    *   More? Come back on January.


# Dec 12 {#dec-12}



*   Node fencing discussion (speaker: Yaniv Bronheim, Red Hat)
    *   initial proposal: https://github.com/kubernetes/community/pull/1416
    *   Dawn: One concern for Node Fencing is that, the concrete treatment might be vendor specific, can we have a general node fencer to do that? The reason why we haven't worked on remedy system is similar.
    *   In general, is this useful as a Kubernetes cluster component?
        *   Need to talk with sig-cluster lifecycle if want this to be a default cluster addon.
        *   The proposal itself is reasonable.
        *   Other vendors may have similar solution in their environment, so we can't say whether this should become the default.
    *   Continue the POC, it is useful, especially for on-prem user.
*   [KEP](https://github.com/kubernetes/community/blob/master/contributors/design-proposals/architecture/1-kubernetes-enhancement-proposal-process.md) vs community design docs: what is the preferred way to propose features for sig-node? (Connor)
*   SecureContainer updates
    *   Slides: https://docs.google.com/presentation/d/1yIgNcZjNoIMNNbErBArbJq2bEn16FnMROskmym_T1nM/edit?usp=sharing
    *   Entitlements: https://github.com/moby/libentitlement
        *   Higher level security abstraction on top of all the existing (linux) security features.
        *   Q: Are we going to change CRI for this? - A: There are several options:
            *   Option 1: Kubelet translate entitlements into a list of configurations and pass to CRI, and keep CRI unchanged.
            *   Option 2: Pass through the entitlements to container runtime, and let container runtime deal with it. (Need CRI change).
        *   Open Question:
            *   How to interact with other plugins on the node, e.g. Device Plugin?
            *   Should entitlement be at pod level or container level? - Seems to be both, e.g. network entitlement at pod level, some others at container level.
    *   Secure Containers/Pods:
        *   Katacontainers: https://katacontainers.io/
        *   Q: Talking about the secure sidecar example, what's the difference with today's existing security solution? - A: Hard to measure or define a bar for security level, but today's linux container security features are not enough.
        *   Virtual kubelet https://github.com/virtual-kubelet/virtual-kubelet. Similar with what @spotter proposed before.
*   FYI RHEL e2e-node results for regular and containerized kubelet now in TestGrid (Seth)
    *   Big thanks to Jan Chaloupka for the work on this
    *   Uses Red Hat Origin CI infra for doing the tests and publishing to GCS
    *   https://k8s-testgrid.appspot.com/sig-node-kubelet#kubelet-conformance-aws-e2e-rhel
    *   https://k8s-testgrid.appspot.com/sig-node-kubelet#kubelet-containerized-conformance-aws-e2e-rhel
*   1.9 release status updates


# Nov 21 {#nov-21}



*   Feel free to propose KubeCon signode F2F meeting agenda: https://docs.google.com/document/d/1cwKlEqwBCveFUX0vZgB-Si5M6kvnaedGblpleo1NERA/edit
*   Derek: Why are we enabling the device plugin alpha feature by default?
    *   @vishh is owner, need to follow up on this #55828
*


# Nov 14 {#nov-14}



*   Device plugin enhancement by Jiaying
    *   Feature was release in Alpha in 1.8
    *   A couple of enhancements in 1.9:
        *   Completely support dynamical registeration for plugin lifecycle
        *   Introduce garbage collection mechanism for device plugin
        *   export gpu related stats through cAdvisor
*   Containerd & cri-containerd status updates
    *   Containerd F2F Meeting Notes: https://docs.google.com/document/d/1TDp5kEuWLS9cL38dZmpffRowd7CKWbzBQrSGawVCoJ0/edit?usp=sharing
    *   Docker may have some behavior change with the new GC model.
    *   In short term, nothing changed in kubelet, kubelet will still do synchronized image garbage collection to reclaim disk resource.
    *   In long term, we may want to more aggressively remove image. And in that case, we may need an asynchronized image removal or policy based garbage collection.
*   Container runtimes release policy
    *   Should publish test result for each release - node e2e, cluster e2e, performance.
    *   How to add test result to node perf dashboard?
        *   Vendor may want to run node performance dashboard in their own environment, because:
            *   We've packaged node performance dashboard, although may not be quite polished.
            *   Today's node performance dashboard is running on top of GCE, may not make sense to some vendor.
    *   CRI-O can publish an image, and picked up by node e2e, but should not be PR blocker.
    *   CRI-O will publish test result on test-grid.
    *   We do need to run test on different OS distributions with different container runtime, however we are not sure whether we should run them on GCE.


# Nov 7 {#nov-7}



*   Pod creation latency regression:  https://github.com/kubernetes/kubernetes/issues/54651
    *   Detected the issues through density tests
    *   Caused by cni 0.6.0 updates
    *   Mitigate the issues through PR #?
    *   Still under investigation. One pr to disable DAD https://github.com/kubernetes/kubernetes/pull/55247, not merged yet.
    *   Need to define the SLOs now? Need large audience for the decision.
*   Live node capacity change: https://github.com/kubernetes/kubernetes/issues/54281
    *   concerns: 1) dynamically scale up and down the capacity, which causes system-overhead being dynamically changed. 2) Too complicated. 3) Difficult to management QoS tree cgroups.
    *   No strong use cases for this feature request.
    *   Can workaround by restarting Kubelet, then prototype.
*   Node Problem Detector enhancement
    *   Proposal made by Andy Xie : https://docs.google.com/document/d/1jK_5YloSYtboj-DtfjmYKxfNnUxCAvohLnsH5aGCAYQ/edit?usp=sharing
    *   By default there is no scripts being included in main repo. Only building a framework.
    *   Not use as the remedy system
*   Pod level metrics
    *   https://docs.google.com/document/d/19bvc8rY5w5ED2yqve8AQdnczumFgrnsbYZk-PV8CyGU/
*   Process improvements: who's interested?/pitch (Lauri, lappleapple in Slack)
    *   Suggestions from the SIG:
        *   anthology or newsletter by SIGs with updates of major news
        *   low-volume communication channel to identify/discover key design changes
    *   Added context: [Value Delivery Mgmt in the K8s Ecosystem](https://docs.google.com/document/d/113gnr3tKXv79J0IYHyg6VbQwLfb-sCYGphW4D9dLZ_E/edit#) proposal, [Kindness Under Pressure](https://github.com/LappleApple/Teaminess-as-a-Service/blob/master/Kindness-Under-Pressure-Exercise.md) session concept


# Oct 31 {#oct-31}



*   Shared PID FYI: [#1048](https://github.com/kubernetes/community/pull/1048) now proposes changing the CRI
*   Proposal FYI: https://docs.google.com/a/redhat.com/document/d/1FvK_fProM1910m1q0PHoW2nqXMWO8hgrGzkUVNy3KKM/edit?usp=sharing


# Oct 24  {#oct-24}



*   Cancelled since there is no topic proposed.


# Oct 17 {#oct-17}



*   Logging issues
    *   logging vs. CRI:
        *   Settled with file-based logging for containers.
        *   Kubelet consumes log files directly, so that Kubelet can have full control on logs
        *   Issues:
            *   logging agent to understand the metadata.  Can we start making a list?  Issue exists somewhere (link?)
            *   Logging format.  Issues with multiline (partial field?). fluentd compatibility?
            *   Who should rotate logs?
                *   Used to use logrotate, which used to cause lost log lines.  Switched to docker json logging driver, but this cannot be used for all runtimes.
                *   Kubelet could rotate, but runtime would need to re-open files. CRI-O just implemented this.
            *   Goal for 1.9: finalize the design
            *   Next actions:
*   PSA: Never too early to start thinking about docs/release notes for 1.9, also [Release Team](https://github.com/kubernetes/features/pull/487) is finalized - 11/22 code freeze, 12/13 release
*   Sig specific blog!


# Oct 10 {#oct-10}

]['



*   Kubelet bootstrap-checkpoint for self host cluster:
    *    on track for 1.9 release. P0
    *   Owner: Tim
    *   Pending prs
*   GPU metrics updates
    *   Design doc: https://docs.google.com/document/d/13O4HNrB7QFpKQcLcJm28R-QBH3Xo0VmJ7w_Pkvmsf68/edit#heading=h.mdp6bb226gj7
    *   /summary API to expose GPU metrics
    *   cAdvisor collects GPU metrics through [NVIDIA Management Library (NVML)](https://developer.nvidia.com/nvidia-management-library-nvml).
    *   cAvisor will be dynamically linked, hence Kubelet too
    *   Raised question: why not in Device Plugin? Time concerns from accelerator folks.
*   CRI-containerd and cAdvisor integration updates
    *   Slides: https://docs.google.com/presentation/d/1Os3nyMRBlFuiBLCjPgeaPv6jXylrZW5jiDXJejlA3Wg/edit?usp=sharing
    *   Summary from yesterday's offline discussion
    *   Core Metrics vs. Monitoring Metrics for containerd integration
        *   core metrics from runtime through CRI
        *   To support full /summary API, the rest metrics will be collected through cAdvisor
    *   From sig-instrumentation (Solly): in the future, /summary is required? Not necessary. It depends on separate monitoring pipeline.
    *   Monitoring agent on node: how to figure out container and pod relationship. Expecting runtime can expose container related metrics.
    *   Plug-in resource, like GPU?
    *   User-defined metrics?
*   Peak memory usage for pod/container
    *   Not in /summary API
    *   But expose this stats through cAdvisor API
        *   https://github.com/google/cadvisor/pull/1768
*   [#53396](https://github.com/kubernetes/kubernetes/issues/53396) Missing modprobe ([PR](https://github.com/kubernetes/kubernetes/pull/53642), 1.8), Debian Bump [#52744](https://github.com/kubernetes/kubernetes/pull/52744) (master) [rphillips]


# Oct 03 {#oct-03}



*   Docker Validation for K8S releases:
    *   Proposal:  https://github.com/kubernetes/kubernetes/issues/53221
    *   Although each Docker version has a particular API version, but it does support a range of API versions.
    *   Stephen: How long does the Docker support window need to be to make this not a problem? - Dawn: 9 months?
    *   Derek: We'll have multiple CRI implementations, and this should be the responsibility of each CRI implementer instead of sig-node putting resources for container runtime validation.
        *   Dawn: Agree.
            *   Docker is still a bit different. Docker is the container runtime we're using for a long time. This is a regression to the user, the change to validate based on API version is still required today.
            *   But sig-node should not sign on validating any other container runtime. We'll provide portable validation tool for different container runtime implementation, e.g. [CRI Validation Test](https://github.com/kubernetes-incubator/cri-tools/blob/master/docs/validation.md).
    *   Derek: They are going to add node e2e test on centos and fedora to avoid the [distribution specific docker issue](https://github.com/kubernetes/kubernetes/issues/52110) in 1.8 release in the future.
    *   Stephen: If there is anything Docker could help, feel free to reach out to them.
    *   Dawn: It seems that there is no objection of validating Docker API version in Q4.
    *   Derek: When can we stop giving recommendation for docker versions? Dawn: We are already doing this for a long time, and there are still users and sig groups relying on our decision. But other container runtime should carry the validation themselves.
*   cAdvisor refactor discuss: [Slides](https://docs.google.com/presentation/d/1eZIKbJ5DibBN_CRVQfhZ7Vfs1ZGnlApOATXrsjnFMyw/edit?usp=sharing)
    *   Does On-Demand Metrics include pod level metrics? Node-Level is higher priority and more useful to eviction at least for now.
    *   Possible issues:
        *   CAdvisor metrics endpoint and api need to change (e.g. no container metrics), this may break existing users.
        *   There is work going on to collect per-container GPU metrics in cadvisor, which is not part of CRI today. => Derek & Dawn: We may be able to get this from device plugin. => This is not what Accelerator team tackling now, may take more time.
    *   Some concern about the skew between cadvisor metrics and container runtime metrics. However, we already have skew today, metrics of different containers are collected separately and periodically.
    *   In the long run, kubelet will make decision based on node/pod level metrics. Today's pod level metrics is derived from container metrics, but may not be the case in the future.
    *   HPA, VPA need container metrics.
    *   Q4 we'll work on cadvisor refactoring; Q1 2018 we could start conversation with sig-instruments to talk about separate monitoring pipeline and standalone cadvisor.
*   Shared PID: [Are breaking changes for CRI v1alpha1 encouraged?](https://github.com/kubernetes/community/pull/1048#discussion_r141199512)
    *   Both cri-containerd and cri-o are fine with the CRI change to use enum for pid namespace (shared/host/private)


# Sep 26 Proposed Agenda {#sep-26-proposed-agenda}



*   Isolated PID namespaces in non-docker runtimes [verb]
    *   Can we make this the default for all runtimes implementing CRI?
        *   containerd: either one is fine
        *   cri-o: either one is fine, but prefer the isolated
        *   rktlet: ?
        *   frakti: not relavant anyway
        *   Windows container support: not relavant
*   Containerd integration status update https://github.com/kubernetes-incubator/cri-containerd
    *   Cut v1.0.0-alpha.0 release this week.
    *   Missing container metrics (pending PR under review, merging this week).
    *   Image gc is not supported in containerd yet, will be added later.
        *   Containerd does provide functions to remove underlying storage of image, but the interface is relatively low level, which is not easy enough to use.
        *   Containerd is adding garbage collection support, which will be much easier to consume.
        *   CRI-Containerd image gc will build on top of containerd garbage collection.
    *   User guide on how to create cluster with containerd is available: https://github.com/kubernetes-incubator/cri-containerd/blob/master/contrib/ansible/getting-started.md
    *   A running sockshop demo: http://104.197.75.252:30001
    *   Cluster e2e tests will be added soon
    *   1.9 plan: bug fixes, and set up CI tests. focus on: tools / tests / util
*   Windows Container Support - update from peterhornyack@google.com
    *   [SIG-Windows current status](https://github.com/apprenda/sig-windows), [roadmap](https://docs.google.com/document/d/1LWi9-NZslZM5lTzMYoAKWXPAJwte5Ow3ySqkrGdHoLg/edit#heading=h.fz2smrec3l4) and [meeting notes](https://docs.google.com/document/d/1Tjxzjjuy4SQsFSUVXZbvqVb64hjNAG5CQX8bK7Yda9w/edit)
        *   Targeting "beta" for 1.9 release
    *   Windows Server "17.09" release (coming early October) will significantly improve native container networking - see [k8s blog post](http://blog.kubernetes.io/2017/09/windows-networking-at-parity-with-linux.html)
        *   Cloudbase (OVN/OVS) and Tigera are concurrently working on their own CNI networking plugins for Windows
    *   Recent PRs:
        *   [Windows CNI](https://github.com/kubernetes/kubernetes/pull/51063) - active
        *   [Windows kernel-mode proxying](https://github.com/kubernetes/kubernetes/pull/51064) - merged
        *   [Windows container stats via kubelet](https://github.com/kubernetes/kubernetes/pull/50396) - merged
            *   [Plan](https://github.com/kubernetes/kubernetes/pull/50396#issuecomment-324728727) for Windows metrics going forward
        *   [CRI stats in docker shim](https://github.com/kubernetes/kubernetes/pull/51152) - nearly merged
    *   Other current tasks
        *   Validating functionality: storage + volumes, secrets, and metrics
        *   kubeadm support for bringing up clusters with Windows nodes
            *   Once this is working, automated testing will follow
    *   Cluster support
        *   master on linux node, only windows worker nodes join the cluster.
    *   Questions / comments? Ask on #sig-windows on Slack


# Sep 19 {#sep-19}



*   rktlet - https://github.com/kubernetes-incubator/rktlet
    *   Some slides: \
https://docs.google.com/presentation/d/1SoBtxvs2kSs7aad2GByafov8AnKJZ6Z0vB5gxVpuv1c/edit#slide=id.p
    *   Started to work on rktlet again since a couple of weeks ago (Kinvolk with Blablacar)
    *   Demo with Weave Socks with Kubernetes & rktlet (CRI), by Iago (@iaguis)
    *   TODOs:
        *   start integration test (conformance, e
        *   2e tests)
        *   log/attach etc.
    *   In the 1.9, we want to delete the rkt package in the main Kubernetes repo. Rktnetes user should switch to rktlet. Rktlet is still WIP, so they are not sure whether they could remove the in-tree rkt package.
    *   What version of Kubernetes does rktnetes support? We don't have test result for it for a while. - Probably 1.7, not sure.
    *
    *   Demo.Use kube-spawn to bring up a cluster: https://github.com/kinvolk/kube-spawn
        *   A demo of sock-shop: https://github.com/microservices-demo/microservices-demo
    *   What is the particular feature we really want from rkt integration?
        *   Major advantage. Daemon restart doesn't restart containers. (Docker live-restore, cri-o and cri-containerd all support this)
        *   Different image formats
            *   ACI image format
            *   It also supports docker image format (thus support OCI image format). Conversion is still required.
*   Other topics?


# Sep 12 Proposed Agenda {#sep-12-proposed-agenda}



*   [Debug Containers Proposal](https://github.com/kubernetes/community/pull/649): Updates and process to merge [verb]
*   [Shared PID Proposal](https://github.com/kubernetes/community/pull/1048) for 1.9 [verb]
    *   Full support for isolated namespaces or migration to shared?
    *   per-pod shared namespace in 1.9
    *   Support shared pid namespace in a long run (v2 API)? or support both modes?
    *   Agreed: CRI changes to support both modes.
*   Triage for 1.8 feature issues
    *   [Extension to support new compute resources](https://github.com/kubernetes/features/issues/368)
        *   Related: [Document extended resources and OIR deprecation](https://github.com/kubernetes/kubernetes.github.io/pull/5399)
        *   Docs PR merged
    *   [CRI validation test suite](https://github.com/kubernetes/features/issues/292)
    *   [Containerd integration](https://github.com/kubernetes/features/issues/286)
        *   updated with status
    *   [Dynamic Kubelet Configuration](https://github.com/kubernetes/features/issues/281)
    *   [Containerized Mounts](https://github.com/kubernetes/features/issues/278)
    *   [CPU Manager](https://github.com/kubernetes/features/issues/375)
        *   [Static CPU manager policy should release allocated CPUs when container enters "completed" phase.](https://github.com/kubernetes/kubernetes/issues/52351) (BUG)
        *   [Node e2e tests](https://github.com/kubernetes/kubernetes/pull/51041)
        *   Docs are merged
    *   [Add support for pre-allocated hugepages](https://github.com/kubernetes/features/issues/275)
        *   Doc PR opened https://github.com/kubernetes/kubernetes.github.io/pull/5419
        *   Feature issue updated with link
    *   [Support for Hardware Accelerators](https://github.com/kubernetes/features/issues/192)
    *   [Further differentiate performance characteristics associated with pod level QoS](https://github.com/kubernetes/features/issues/276)
        *   this is duplicate with the CPU manager feature issue (but predates)
        *   will close the QoS feature then [Caleb]
*   [1.8 release notes draft](https://github.com/kubernetes/features/blob/master/release-1.8/release_notes_draft.md)
    *   Please update and sign off on release notes for node components
    *   Please update "major themes" section by SIG with what the development goals were for the 1.8 release [Derek to do]
*   1.9 reliability release planning
    *
*   [#50865](https://github.com/kubernetes/kubernetes/issues/50865): Rebuilding pause containers [verb]


# Sept 5 {#sept-5}


# Aug 29 {#aug-29}



*   Shared PID namespace discussion (part 2) [derek, mrunal]
    *   Appealing reason: not all containers within a pod can share pid namespace, due to security concern.
    *   Breaking the backward compitility
    *   Considering introducing v2 Pod API to make shared pid namespace by default.
    *   With the latest inputs, we decided 1) disable the feature by default in 1.8, 2) discussed pod level api in 1.9.
*   cAdvisor native support for cri-o [derek, mrunal]
    *   any objections?
    *   how to plugin w/ new refactor (cadvisor vs cri)
    *   Try to rollout cri-o, and may carry patch for 1.8 to make it work.
    *   Even container stats come from CRI, but they still run cadvisor to monitor other stuff, for example, filesystem stats. They don't want to run 2 instances of cadvisor, one in kubelet and another in cri-o.
    *   CRI is introduced to make kubelet runtime agnostic. This breaks the initial goal in some way.
    *   Based on data from their environment, cgroups parsing is expensive, @derek doesn't want 2 agents parsing cgroups.
    *   Containerd provides the container metrics itself, so it doesn't have this problem. However, we may still need to take care of the overhead. Containerd collects stats on-demand, it's cri-containerd/kubelet's responsibility to control the polling period and whether to cache the stats or not.
    *   Cadvisor has introduced many issues each release whenever it's updated, because it's not well maintained in current stage. So we discussed that not to add more complexity into cadvisor.
    *   Can we make cadvisor metrics collecting logic configurable, so that we could collect different runtime stats based on the passed-in configuration, instead of changing cadvisor code every time.
    *   Talk about the short term after the sig-node meeting. Reopen the topic about the container metrics through CRI.
*   PSA: [Open issues in the 1.8 milestone](https://github.com/kubernetes/kubernetes/issues?utf8=%E2%9C%93&q=is%3Aissue%20is%3Aopen%20milestone%3Av1.8%20label%3Asig%2Fnode) ~ need resolution or removal from the milestone or else they are considered release-blocking
*   PSA: [Open feature issues in the 1.8 milestone](https://github.com/kubernetes/features/issues?utf8=%E2%9C%93&q=is%3Aissue%20is%3Aopen%20label%3Asig%2Fnode%20milestone%3A1.8) - need update
*   PSA: Ensure [release notes for Node](https://github.com/kubernetes/features/blob/master/release-1.8/release_notes_draft.md#node-components) are correct


# Aug 22 Proposed Agenda {#aug-22-proposed-agenda}



*   [#48937](https://github.com/kubernetes/kubernetes/issues/48937): Shared PID namespace with docker monoliths  [verb]
    *   gluster/gluster-centos is a container image that runs systemd so that it can run ntpd, crond, gssproxy, glusterd & sshd inside a single container
    *   When invoked as /usr/sbin/init, systemd refuses to run if pid != 1
    *   The image runs with the following config: \
`     env: \
      - name: SYSTEMD_IGNORE_CHROOT \
        value: "1" \
      command: \
      - /usr/lib/systemd/systemd \
      - --system`
    *   Should shared pid be enabled for docker in 1.8?
    *   Do we still agree that in the long term we should try to push forward inner pod shared pid namespace? - Could be default, but should be able to overwrite.
    *   Overwrite could be done at node or pod level. Should this be node level configuration, or part of pod spec? Pod spec overwrite means API change and long term support, do we what to do that?
    *   **Kubernetes 1.8 decision:** Default shared pid namespace for Docker 1.13 with node level configurable. For other Docker versions, shared pid namespace will always be disabled.
    *   **Revisit this in Kubernetes 1.9:** discuss more granulary configuration.
    *   Other runtime should also share pid namespace. Do we have test coverage for this? We do have a node e2e test for it, but it's docker specific for now because of the docker version check.
        *   It also implies that the non-docker runtime must support isolated namespaces
*   https://github.com/kubernetes/kubernetes/pull/50859 [derekwaynecarr]
    *   Alpha support for pre-allocated hugepages
    *   A new feature flag will be introduced to enable the hugepage support;
    *   Cadvisor will support huge page;
    *   Could be consumed via EmptyDir.
*   https://github.com/kubernetes/kubernetes/pull/49186 [sjenning]
    *   CPU manager
    *   We may not want to have more than 2 new features from resource management each release, we are now lack of review bandwidth.
    *   We could have multi-stage review, first round for technical review, and production review etc.
    *   We need to grow the reviewer pool in 1.9 for /pkg/kubelet based on updated contribution in community
*   https://github.com/kubernetes/kubernetes/pull/50396 [davidporter]
    *   Windows Container Stats
    *   Container stats should go through CRI, not done yet.
    *   Q: How efficient is the newly introduced windows stats? Cadvisor does per-10-second poll. - A: Not sure, just call docker stats API to get stats.
    *   Node level stats is windows specific, use windows perfcounters now.
    *   Q: Should the perfcounters stuff in another repo? Because there are only 2 files, does it make sense to move it to another repo? - A: It's fine to leave the node level metrics implementation in kubelet as temporary solution. In the future, we may extract core cadvisor out as a library in Kubelet, and run cadvisor as a standalone daemon.
    *   Q: Do we have a document or any test result about what features are support for K8s on windows? - A: E2E test is still a TODO now.
    *   Q: Beta in September is not very comfortable to us, we need better document and test. We once deleted the previous windows implementation by mistake, and no test caught it. - A: They don't have e2e test now, and they are working on that. They do have several users now.
    *   Sig-windows wants to be first citizen, and if there is any feature introduced affect windows, they want to be in the discussion.
    *   Derek: don't think that Windows can totally hide behind container runtime, there are many other kubelet level features which need corresponding windows support: EmptyDir, Resource Management etc.
    *   Q: Who's reviewing the windows PR? - A: Currently, mainly Dawn and Yuju. However, code freeze is coming, we may not have enough review bandwidth this release, and there should be a corresponding feature in feature repo.
    *   **Current decision:** 1) We could add a flag to disable the monitoring feature which causes kubelet crash loop on windows; 2) We are fine with leaving the stats code in Kubelet for now; 3) They need to rewrite the PR based on the feedback, and properly plan and design Kubelet windows support.
    *   We may not have time to review the windows container stats PR this release, we may have to do that after code freeze.
*   https://github.com/kubernetes/kubernetes/pull/50984 [tstclair-offline]
    *   PSA - Checkpoint POC available for feeback.
*   Pod checkpointing [proposal](https://docs.google.com/a/google.com/document/d/1qmK0Iq4fqxnd8COBFZHpip27fT-qSPkOgy1x2QqjYaQ/edit?usp=sharing)
*   kubelet pod mount propagation
    *   https://github.com/kubernetes/kubernetes/pull/46444
    *   https://github.com/kubernetes/community/pull/659
    *   https://github.com/kubernetes/community/pull/589
    *


# Aug 15  {#aug-15}



*   Resource Management Workgroup updates
    *   Crossing several sigs: node, scheduling, etc.
    *   1.8 releases, 3-4 projects for better support high performance workloads
        *   static cpu pining: The guaranteed pod will be assigned the exclusive cpus. Will be alpha feature in 1.8 release. More enhancement, such as dynamic in the future
        *   device plugin: a separate daemonset to handle device discovery, initiation, and allocation and destroy. Ready for alpha release.
        *   hugepage support: For stateful set. Alpha feature in plan.
*   Mount namespace propagation reviews needed:
    *   https://github.com/kubernetes/community/pull/659
    *   https://github.com/kubernetes/kubernetes/pull/46444
*   Kubelet checkpoint
    *   Two approaches proposed, need to decide which one should be the final one.
    *   Scopes for the feature? - Secrets is excluded from the scope for now.
*   PSA: Release notes [drafts due](https://groups.google.com/d/msg/kubernetes-dev/sEP1YRWBnEk/tYWnDAKfBwAJ)


# Aug 8 Proposed Agenda {#aug-8-proposed-agenda}



*   [Kubernetes preemption / eviction priority](https://github.com/dashpole/community/blob/2f68945935155b3071390986aa592cd49ad8e0f3/contributors/design-proposals/priority-eviction.md)
    *   Preemption vs. Eviction: Preemption is invoked by scheduler, eviction is invoked by Kubelet upon the resource starvation situation.
    *   Preemption will preempt pod gracefully.
    *   Should we define SLO for this?
        *   Q: E.g. guaranteed job could still be preempted, should we have an SLOs for this, how rarely should this happen? - It's cluster admin's responsibility to properly configure the system to meet their SLOs.
        *   Q: What about SLIs? - Yes.
    *   Q: Will kube-proxy be preempted?
        *   By default, all pods have the same priority, so there's no difference from before. However, if new pods are added with higher priority, it's possible to preempt kube-proxy.
        *   We also have 2 default priority classes `node-critical` and `cluster-critical`.
    *   Roadmap:
        *   Eviction will follow the new priority schema in 1.8.
        *   Try to get preemption in in 1.8.
        *   Priority in resource quota in 1.9.
*   Kubelet checkpoint proposal
    *   https://docs.google.com/document/d/1hhrCa_nv0Sg4O_zJYOnelE8a5ClieyewEsQM6c7-5-o/edit?usp=sharing
    *   https://github.com/kubernetes/kubernetes/issues/49236
*   Device plugin proposal updates
    *   https://github.com/kubernetes/community/pull/695
    *   Q: Authentication to kubelet? How do we know whether the device plugin could talk with kubelet?
        *   Communication between device plugin and kubelet is through a host unix socket, and only privileged pod could create host unix socket.
        *   The alpha version will rely on this assumption, and it's cluster admin's responsibility to only grant privileged permission to trusted device plugin.
    *   Q: Should we create a separate repo for the device plugin?
        *   It's hard to track all kinds of repositories during release process.


# Aug 1 {#aug-1}



*   Containerd-CRI status updates
    *   [cri-containerd Q3 plan](https://docs.google.com/document/d/1tx9NUm6UsWBWjB98rm6QGnwDbl7IX3f9PAtk-famrHM/edit#)
    *   `Current status:`
        *   `Features: cri-containerd` v0.1.0 supports all basic functionalities including:
            *   Sandbox/container lifecycle management;
            *   Image management;
            *   Sandbox networking;
            *   Container logging;
            *   Run command synchronously in container;
            *   …
        *   Test:
            *   CRI validation test: 30/36 (pre-submit test).
            *   Node conformance test: 116/122.
        *   In-Use Containerd Version: [v0.2.3-1098-g8ed1e24](https://github.com/containerd/containerd/commit/8ed1e24ae925b5c6d8195858ee89dddb0507d65f)
        *   Newest Containerd Version: v1.0.0-alpha2 (weekly alpha release)
    *   Q&A:
        *   Containerd release schedule after 1.0: Probably monthly at first, and then focus on bug fixes and security patches.
        *   Package solution: Only release binaries for convenience of testing. No plan to package for now.
        *   Swarm containerd integration: Working version merged.
        *   Docker containerd integration plan: Happen in a moby branch now.
        *   API change review process after 1.0: Deprecation policy, support policy etc. Don't have one, and haven't thought about this yet. Need a proposal, we can submit one and discuss.
*   PSA: Feature freeze is today ~ [Current feature list](https://github.com/kubernetes/features/issues?q=is%3Aissue+is%3Aopen+label%3Asig%2Fnode)


# July 25 {#july-25}



*   summarize mount propagation conclusions (@tallclair)
    *   Alpha feature in 1.8: non-priority container with hostpath will get slave mode; while priority container with hostpath will be opt-in with shared-propagation mode, thus it is visible to the outside.
    *   There is no security concerns.
    *   API changes: annotation for alpha.
    *   Using gRPC over the socket, similar to CRI model.
    *   Owner: Jan
*   discuss problems surrounding pod mounts supported by userspace processes (i.e. fuse) (@sjenning)
    *   tl;dr, Some filesystems require user processes to maintain mount points.  These processes are children of the kubelet and are started in the kubelet.service cgroup when running in systemd.  Restarting kubelet.service kills them.  Setting KillMode=process prevents that but lets other node children leak (i.e. find/du/journalctl exec'ed by cadvisor, for example)
    *   https://github.com/kubernetes/kubernetes/pull/23491
    *   https://github.com/kubernetes/kubernetes/issues/34965
    *   Owner: Seth
*   Present: system spec and image validation (@ygg)
*   Sig-windows updates (@Shaya)
    *   The immediate concerns are getting it into a usable state, which includes finishing the network story (using ovn/ovs), volumes and metrics.
    *   Suggested them to go with CRI approach, but not in 1.8 timeline.
*   Initial alpha proposed checkpointing https://github.com/kubernetes/kubernetes/issues/49236


# July 18 {#july-18}



*   set up mount propagation meeting
    *   jsafrane@redhat.com can't be on sig-node today
    *   offers:
        *   10am PDT this Wednesday or Thursday
        *   9am PDT this Friday
        *   9am PDT next Monday (24th)


# July 11 {#july-11}



*   Q3 planning
    *   https://docs.google.com/document/d/1Sq3Cr0_udLtksBaogQWhAhwjoG5jj8TfHzftw4Kzg0I/edit
*   mount propagation (@jsafrane)
    *   https://github.com/kubernetes/kubernetes/pull/46444
    *   https://docs.google.com/document/d/1XXKdvFKnbV8MWNchjuLdLLCEs0KdWelpo3tpS6Wr18M
*   Debug Container API compromise (@verb)
    *   [kubernetes/community#649](https://github.com/kubernetes/community/pull/649)
    *   [Google docs version of Debug Containers proposal](https://docs.google.com/document/d/1tds_D3aoUtMjlKpuVdr88oDrRzUS4OuShlen53RSAo8/edit?ts=595cb285#heading=h.vf1xpupyfm12)


# July 4 {#july-4}

No meeting


# Jun 27 {#jun-27}



*   Rktlet status (@alban)
    *   https://github.com/kubernetes-incubator/rktlet/issues
    *   Need rktlet integration and test status update.
    *   Dawn: Why does @alban care about the underlying runtime?
        *   They are using existing rktnetes integration, so they want to make sure rkt keep working with Kubernetes, rktlet seems to be the way to go.
        *   They are using ACI image. (Dawn: Why ACI instead of Docker image?) Not quite clear, one is that they want dependency between images. They use [dgr](https://github.com/blablacar/dgr) to build container images.
*   1.7 Docker version information? (@calebamiles)
    *   Support 1.12.? - Yeah, already supported in K8s 1.6.
    *   Feature wanted from new docker version (>= docker 1.12):
        *   Overlay2: Need validation. Cadvisor already supports overlay2 now.
        *   Live restore: Help node reliability. Need validation.
        *   Shared pid namespace: Shared pid namespace bug in docker 1.12 is fixed in docker 1.13.
    *   Docker version supported in K8s 1.7 is 1.10, 1.11 and 1.12. We plan to validate docker 1.13 in K8s 1.8 and deprecate docker 1.10.
    *   Where is the docker version related information now? - Mostly tracked by different docker related Kubernetes issues. We always include the docker version supported in release notes, and link to known issues.
    *   [1.12](https://github.com/kubernetes/kubernetes/issues/28698)
    *   [1.13](https://github.com/kubernetes/kubernetes/issues/42926)
    *   @Dawn: We don't enforce docker version or container runtime for vendors. We've provided the portability (CRI), and also portable validation test ([CRI validation test](https://github.com/kubernetes-incubator/cri-tools)). Each vendor should validate and choose container runtime and version based on their own business/technical interests.
    *   @timothysc: An integration testbed is required for this issue.
    *   @Dawn: Node e2e and conformance test are also built for vendors and users to validate their node setup including the docker version.
    *   @Lantao: Kubeadm also enforce docker version in the pre-flight check.
    *   We already run many tests for docker validation, but we do need to organize and expose the information better.
        *   Node e2e test: https://k8s-testgrid.appspot.com/sig-node#kubelet
        *   Docker validation test: https://k8s-testgrid.appspot.com/google-docker (e2e-cos suite is not working properly because of some legacy reason, our GCI team is working on that)
    *   @michaelcrosby: Containerd will run integration test to avoid PRs breaking swarm, Kubernetes etc.
    *   @Dawn: Both cluster e2e and node e2e are too heavy for containerd to carry. That's why we build the [CRI validation test](https://github.com/kubernetes-incubator/cri-tools), which test against CRI directly which is much more lightweight.
*   mount propagation (@jsafrane)
    *   https://github.com/kubernetes/kubernetes/pull/46444
    *   https://docs.google.com/document/d/1XXKdvFKnbV8MWNcchjuLdLLCEs0KdWelpo3tpS6Wr18M
    *   Schedule design meeting, 9:00 am PDT


# Jun 20 {#jun-20}


# Jun 13 {#jun-13}



*   1.7 release updares
    *   DynamicKubeletConfig alpha was dropped from 1.7. But will merge it once the code freeze is lifted.
*   CRI-O presentation and demo (Samuel Ortiz [samuel.ortiz@intel.com](mailto:samuel.ortiz@intel.com)/Mrunal Patel/Antonio Murdaca)
    *   Slides https://docs.google.com/presentation/d/1lqH3VNhYUmp0WbBZ6iNzbNPPkHSPVHc2XH74qUjyNDI/edit?usp=sharing
    *   Package solution: working on it for all os distr after 1.0.0
    *   Plan to cut 1.0.0 alpha: week of June 13
*   kubelet event spam in large clusters or pods in constant backoff
    *   budget per pod per event type (Warning vs Normal) per interval
    *   budget per namespace
    *   do we really need to report at such granular levels
        *   https://github.com/kubernetes/kubernetes/pull/47367#issuecomment-308002427
*   D


# Jun 6 {#jun-6}



*   Virtlet demo (cloud-init support, running Kubernetes in Kubernetes, etc.) \
(Ivan Shvedunov -- [ishvedunov@mirantis.com](mailto:ishvedunov@mirantis.com))
    *   The primary difference between virlet and frakti: can run any vm images: windows, any legacy applications which cannot be containerized
    *   Integrated with Kubernetes through CRI
    *   Demo with VM with a stateful set and service
    *   Using cloud-init to mount Kubernetes required volumes
    *   Potential use cases: any applications, including windows apps, hybrid environmen, malware detection(?)
    *   Can orchestrate vms the same way as containers.
    *   Using libvirt, and introduced a cri-proxy which can decide if launching a container-based pod or virt-based pod.
*   Discussion about the entitlements (not enough time on last meeting) by Nassim
    *   link to the doc?
    *   Image publisher would define the security profiles
    *   github/docker/entitlement
    *   contacts:
        *   timstclair - node security, proposal process
        *   liggitt - security / sig-auth TL
        *   yujuhong - CRI / kubelet / runtime
*   CRI-O presentation and demo (Samuel Ortiz [samuel.ortiz@intel.com](mailto:samuel.ortiz@intel.com))
    *   CRI-O overall progress (Antonio Murdaca [runcom@redhat.com](mailto:runcom@redhat.com)/Mrunal Patel [mpatel@redhat.com](mailto:mpatel@redhat.com))
    *   CRI-O cluster demo (Antonio Murdaca [runcom@redhat.com](mailto:runcom@redhat.com))
        *   pass full node-e2e and e2e tests
        *   support docker v2schema1 images out of the box
        *   can run Clear Containers
*   Quick status update on SIG Node features for 1.7 (Caleb Miles caleb.miles@coreos.com)
    *   [scratch summary](https://gist.github.com/calebamiles/adfc84814d503501bcebd7cd2551d10c)
        *   [CRI validation test suite](https://github.com/kubernetes/features/issues/292)
            *   Done with alpha, and had a demo.
            *   The engineers switched to help containerd integration
        *   [Enhance the Container Runtime Interface](https://github.com/kubernetes/features/issues/290)
            *   API changes for monitoring stats, but not done with the implementation.
        *   [Containerd CRI Integration](https://github.com/kubernetes/features/issues/286)
            *   Achieved the basic goals for this item
            *   Haven't switched to the containerd new API, and new containerd client and binary yet
            *   Containerd plan to introduced v2 scheme1 support
            *   Already cut a release for today's integration
        *   [Dynamic Kubelet Configuration](https://github.com/kubernetes/features/issues/281)
            *   Filed the exception request
            *   Wrote the test plan and started the manual tests

# May 30



*   LinuxKit: build secure and minimal host OSes for Kubernetes - Riyaz Faizullabhoy @riyazdf
    *   Overview of project security goals - [slides here](https://drive.google.com/file/d/0BzuxnKD0WQ4wbGlpdXk4QUdRcWs/view?usp=sharing)
        *   Toolkit for building secure, portable, and lean operating systems for containers
        *   Immutable infrastructure
        *   Leverages buildchain and userspace tools from alpine linux
        *   Securely configure modern kernel (Kernel config) && collaborate with KSPP
        *   Incubating projects: Landlock LSM, Type-safe system daemons && mirageSDK, okernel: Separate kernel into outer-kernel and inner-kernel
    *   Demo of building and running Kubernetes on customized Linuxkit based OS
    *   Next steps for making it easy to use Kubernetes + LinuxKit
        *   list of Kubernetes dependencies document?
        *   other security features that k8s has been thinking about in host OS?
*   Security profile configuration: Entitlements on Moby and Kubernetes?
    *   A user-friendly high-level interface for security configuration based on different categories (network, host devices, resources..) and levels of usage for each one
    *   Quick overview of the proposal by @nass https://github.com/moby/moby/issues/32801
    *   How can this fit in Kubernetes ? What are the requirements ?
    *   https://github.com/docker/libentitlement
    *   Slides https://drive.google.com/file/d/0B2C3Ji-3avH8bUI1eF9zS2d4aGM/view?usp=sharing
*   Minimizing tail latency for applications through cache partitioning
    *   https://docs.google.com/document/d/1M843iO76DiPCGkKU3NNsiuwmL-aYC8HK_5973hfUaH4/edit


# May 23 Agenda {#may-23-agenda}



*   Running unikernels with Kubernetes. We've been experimenting with Virtlet to deploy light OSv-based unikernels; each of these unikernels is comprised of the library OS, nodejs runtime and a single service
    *   The demo briefly explains the rationale, unikernel concepts, caveats, etc.
    *   [Virtlet repository](https://github.com/Mirantis/virtlet)
    *   [Demo app repo](https://github.com/mikelangelo-project/osv-microservice-demo)
    *   [slides](https://docs.google.com/presentation/d/1hMh-kb-zsKpbAtP87Vv4cGODGIwXdkOsvZFt33TWtwE/edit?usp=sharing)
*   Containerized mount utilities [jsafrane]
    *   GCI, CoreOS and AtomicHost don't want to distribute gluster, ceph and nfs utilities on the host, we need a way how to distribute them in a container. @jsafrane (sig-storage) will present https://github.com/kubernetes/community/pull/589 that includes some kubelet changes.
*   PTAL Alpha Dynamic Configuration PR: https://github.com/kubernetes/kubernetes/pull/46254


# May 16 Proposed Agenda {#may-16-proposed-agenda}



*   Frakti demo
    *   [repository](https://github.com/kubernetes/frakti)
    *   demo
        *   Support mixed container runtime: [hyper](http://hypercontainer.io/), docker (privileged container).
        *   Frakti deploy with kubeadm for alpha support.
        *   Support OCI runtime spec, but different with Docker, e.g. only supports hard resource limit because of Frakti is using VM.
        *   Questions:
            *   Run runtime daemon (docker/containerd) inside the VM? - No, run container directly inside the VM with namespace cgroup etc. No security features like seccomp, apparmor inside the VM.
    *   [slides](https://drive.google.com/file/d/0B6uGv-NC7DxDSmREaUhEdXl4NGM/view)
        *   Questions:
            *   Why does hyper support high density, is there security problem here?
                *   No security problem;
                *   Only share the readonly part (Which part?)
            *   Does hyper use VM as a pod or container? - As pod.
            *   How does hyper run multiple containers in the VM? - Simple init system to manage containers inside the VM, called <code><em>[hyperstart](https://github.com/hyperhq/hyperstart)</em></code>.
            *   Will hyperstart monitor and collect the container metrics? - Yeah.
    *   Secure Container Runtime
        *   What kind of api? What are the semantic requirements for a "secure pod"?
        *   Two use cases:
            *   Public cloud, multi-tenancy.
            *   Runtime selectivity (sounds like less important)
        *   Is the current pod concept, security context, and security policy model enough? Or should we have a new high level overall design to support secure pod?
        *   An insecure pod running on the host may make the other secure pods on the host "insecure". We may want to make sure running a "secure pod" in a "secure environment".
        *   Definition: hard multi-tenancy vs. soft multi-tenancy?
        *   What is the user requirement? Currently the security features in Kubernetes are mostly added for engineering reason, or derived from Docker. Can we abstract and categorize these security features based on user requirement?
        *   Action Items:
            *   Proposal: what is the definition, what features are required, what is the security model.
*   Lessons learned from https://github.com/kubernetes/kubernetes/pull/45747
    *   we upgrade clusters and were stuck with 1000s of pods not deleting
    *   it appears that future kubelets must do patch rather than update status


# May 9 Proposed Agenda {#may-9-proposed-agenda}



*   CRI util / test demo
    *   [repository](https://github.com/kubernetes-incubator/cri-tools)
    *   [slides](https://docs.google.com/presentation/d/1VM4Tx85ffYZl_VOnBuxN1VQ-RhiM8PSg0gt_D1mP2-0/edit)
*   Resource Management F2F summit updates
    *   meeting notes: https://docs.google.com/document/d/13_nk75eItkpbgZOt62In3jj0YuPbGPC_NnvSCHpgvUM/edit
    *   Resource class to enable support for resources require additional metadata to management, for example: GPU
    *   Extensible device support
    *   CPU enhancement including cpuset handling
    *   Hugepage will be supported as the first class resource
    *   NUMA support
    *   Performance benchmark
    *   Only agreed upon the list, but not prioritized yet.
*   Pod Troubleshooting alpha features, merge [proposal](https://github.com/kubernetes/kubernetes/pull/35584)
    *   Any volunteers to review PRs?
*   SIG PM would like demos for 1.7 features
    *   Possible demos
        *   Dynamic Kublet config
        *   CRI validation test suite
        *   Containerd CRI implementation
        *   Enhanced GPU support
        *   Pod troubleshooting
    *   [Schedule](https://docs.google.com/document/d/1YqIpyjz4mV1jjvzhLx9JYy8LAduedzaoBMjpUKGUJQo/edit#heading=h.micig3b7ro3z)
*   Should [features/#192](https://github.com/kubernetes/features/issues/192) be moved into the 1.7 milestone
    *   answer no


# May 2 {#may-2}



*   Defer container: https://github.com/kubernetes/community/pull/483
    *   similar to prestophook, but try to solve the issues
    *   For stateful application, so that they can terminated in sequence.
    *   Can we improve upon the existing prestophook?
    *   More challenges at termination stage:
*   Non-CRI integration status:
    *   Plan to clean up non-cri integration with Kubelet
    *   Decide to remove docker management via old API
    *   Keep rkt as is until rkt integration move to CRI


# April 25 {#april-25}



*   Defer container: https://github.com/kubernetes/community/pull/483
    *   Postpone to next week
*   Demo: [Pod Troubleshooting](https://github.com/kubernetes/kubernetes/pull/35584) (verb@)
    *   Improve troubleshooting support in K8s
    *   Debug containers regardless of the debugging tools available in a given container image, and regardless of the state of the container, without direct access to the node.
    *   Adds a `kc debug` command that can run a debug container with debugging tools in the Pod namespace (shared PID namespace).
    *   Does not add a new container to the Pod spec, instead this creates a streaming API call similar to exec, and the Kubelet executes this container in the Pod - the state for debug containers is maintained entirely within the Kubelet and the runtime.
    *   Permissions - do debug containers have more permissions than the Pod's other containers?
    *   You can access other containers' filesystems via /proc/{PID}/root
    *   How does this behave with crash looping Pods? Does this pause some of the Kubelet's restart behavior? Turns out we can create a container in the Pod namespace and the Kubelet won't touch it when syncing the Pod.
    *   State of the debug pod is maintained within CRI just like any other container in the Pod. Debug pod has an additional label.
    *   **Pod status will include these debug containers, even though Pod spec will not.**
    *   Would be useful to report the command that was run in the Pod status as well
    *   Could be useful to be able to specify a default debug container or set of debug actions for a Pod, that would be relevant to that specific Pod
    *   Still need to think about the disk side of the debug container
    *   If we want to make this an alpha feature for 1.7, file feature request in the features repo.


# April 18 {#april-18}



*   containerd F2F meeting
    *   Meeting notes: https://docs.google.com/a/google.com/document/d/1lsz8pd4aIcdAFrsxi0C1vB2ceRCyq36bESoncEYullY/edit?usp=sharing (need permission to access, feel free to request access)
    *   cri-containerd repo: https://github.com/kubernetes-incubator/cri-containerd


# April 11 {#april-11}



*   containerd update
    *   runtime-side: almost feature complete, work on checkpoint and restore.
    *   image-side: centralized storage.
    *   When does Docker start to rebase on new containerd.
        *   Rebase on execution part first when the execution part reaches feature complete;
        *   Rebase on image part will be later. Backward compatibility may need some effort. Snapshotter ←→ Graph Driver


# April 4  {#april-4}



*   1.7 roadmap
    *   https://docs.google.com/a/google.com/spreadsheets/d/1-hADEbGEUrW04QP4bVk7xf1CWuqbFU6ulMAH1jqZQrU/edit?usp=sharing
    *   Any opensource contribution is welcome.
    *   Needs to set the expectation and milestone for debug pod.
*   Containerd and CRI integration updates
    *   POC: https://github.com/kubernetes/kubernetes/pull/43655
    *   Proposal WIP, plan to send out this week, will create a new project in kubernetes-incubator
    *   Alpha release plan: Basic container and image lifecycle management [P0]; container streaming/logging [P1]; metrics [P2].
*   Open floor for additional dynamic config questions at end of meeting


# Mar 28 {#mar-28}



*   Cancelled


# Mar 21 {#mar-21}

**Discuss dynamic kubelet configuration proposal (mtaufen):** https://github.com/kubernetes/kubernetes/pull/29459

I'd like the purpose of the meeting today to be to communicate the status of the proposal, communicate some of the pros/cons of ideas that have been floating around, get feedback from the community on the proposed solution, and come out with action items for improvements to the proposal. I would like the proposal to be approved and merged by the end of this quarter, and I would like the implementation to be complete for 1.7.

There are a number of items in this proposal, and any of them are fair game for discussion today. That said, there are a couple I want to make sure we talk about. So I'd like to discuss each of those first, one at a time. I'll provide a brief intro to each topic, and then open up for questions on that topic. Then once we're through those we can open up for arbitrary questions.

**Key Topics:**

**cluster-level representation of the configuration**

Are there any sensitive fields that should be in secrets instead?

Config map could still tell you where to find these secrets, etc.

We have way too many manually set flags right now, we should be looking at peeling things out of the configurations.

Configuration should be orthogonal to provisioning and node topologies. Software vs. hardware.

given that we are getting rid of componentconfig and moving to each-component-exposes-its-own, we should have a policy for the "right" way to do this, and talk about how to make these per-component "API" types discoverable - there has been some discussion about discussing this more (how meta…), the likely correct home is sig-cluster-lifecycle

**Kubelet config checkpointing and recovery mechanism**

Note: config uptake is defined as a restart

configmap use cases feel like a nodeclass - you _could_ build nodeclass config on top of the proposed model

kubelet could post a minimal status very early before potential crashes - need to clarify precisely what this looks like

how should we define kubelet health/bad config? Any perfect indicator (couldn't deserialize, failed validation, failed checksum, etc.) should be used. The only imperfect indicator at the kubelet level we've converged on thus far is that the kubelet crashloops before the configuration is out of it's probationary period.

**MEETING ENDED HERE, TBC ~~NEXT WEEK~~ next week conflicts with KubeCon, please continue discussion on GitHub**

**~~nodeConfig volume idea~~ - ended up pulling back from this after GitHub discussion**

~~The biggest reason behind nodeConfig was to allow easy coordination of configuration rollouts between multiple components. In recent discussions, we've tended toward the idea that interdependence between component configurations which would require such coordination is an antipattern. This greatly lessens the need for nodeConfig.~~

~~There has been discussion around how to plumb configuration into the pods that run on the node. It recently became apparent that we may be conflating two needs:~~



*   ~~getting knob values to things in daemonish Pods~~
*   ~~informing Pods about node-specific settings, e.g. the node CIDR~~

~~I had originally proposed this idea with the former intention, though it seems that others have interpreted it as being useful for the latter intention. The latter case still seems useful, but in my mind these are separate functions and shouldn't be served by the same mechanism. Thoughts?~~

**composable configuration**

Thus far, we have been discussing config in terms of a specific object type associated with a specific key name in the ConfigMap. In my last meeting with Brian Grant about this feature, we briefly touched on the idea to, instead, have the ConfigMap contain an arbitrary set of keys, each associated with an arbitrary blob representing an arbitrary set of configuration objects. The Kubelet could slurp all of these blobs to compose a set of objects, and search for the objects that it needs in this set. This may make the configuration more composable and reduce the need to adhere to a specific schema in the ConfigMap. I'd like to see what everyone thinks of this idea.

**Open to arbitrary questions**


# Mar 14 {#mar-14}



*   Logging integration: https://github.com/kubernetes/kubernetes/issues/42718
    *   Log rotation:
        *   We may want to revisit the logging SLOs. In reality, dropping logs is unavoidable sometimes. Should we define log missing rate instead?
        *   Current copy-truncate loses logs.
        *   If log rotation is handled by the same daemon which writes the log, it is mostly fine. However, the problem is when log rotation and log writing is handled by different daemons, some coordination between them is needed. We may add api in CRI to signal runtime to reopen log file, although Docker won't react to the signal now.
    *   Log enrichment:
        *   Getting log metadata from apiserver may introduce extra overhead to apiserver. Is it possible to add corresponding kubelet api? - Same problem with standalone cadvisor. This is in sig-node's queue, just need prioritize.
        *   dawn@: Discuss with the api machinery team about the api library.
*   Please continue to move non blocking issues out of the 1.6 milestone and into the 1.6.1, 1.7 or next-candidate milestones **please don't **simply remove a milestone from an issue. Please ensure all non release blocking issues are out of the 1.6 milestone by the end of today (14 March 2017)
    *   [flakes](https://github.com/kubernetes/kubernetes/issues?utf8=%E2%9C%93&q=is%3Aissue%20is%3Aopen%20label%3Akind%2Fflake%20milestone%3Av1.6%20label%3Asig%2Fnode)
    *   [all issues in milestone](https://github.com/kubernetes/kubernetes/issues?utf8=%E2%9C%93&q=is%3Aissue%20is%3Aopen%20milestone%3Av1.6)


# Mar 07 {#mar-07}



*   1.6 release discussion:
    *   CRI logging with journald: PR was merged
    *   Is someone looking at [kublet serial GC failues](https://k8s-testgrid.appspot.com/google-node#kubelet-serial-gce-e2e)?
    *   Thanks for closing all those flakes :)
    *   GPU support: accepted
    *   Release notes


# Feb 28 {#feb-28}



*   containerd summit updates
    *   containerd summit videos:
        *   [containerd deep dive](https://www.youtube.com/watch?v=UUDDCetB7_A)
        *   [containerd and CRI](https://www.youtube.com/watch?v=cudJotS97zE)
        *   [driving containerd oprations with gRPC](https://www.youtube.com/watch?v=sG9hxz4-hIA)
        *   [containerd goverhttps://www.youtube.com/watch?v=cudJotS97zEnance and integration with other systems](https://www.youtube.com/watch?v=PGEXMuBeo1A)
    *   [Discuss sesson notes](https://github.com/docker/containerd/blob/master/reports/2017-02-24.md)
        *   Proposed deadline for 1.0 containerd: end of Q2
        *   Containerd contains: runtime + image
        *   Containerd-shim manages the lifecycle of container.
        *   No network support in containerd
        *   Image distribution management is new, but by design it is very flexible: split image handling into 3 phases
        *   Align with CRI from the high level
        *   Not release any binary. Every vendor builds its own binary.
        *   Multi-tendency: create different namespaces.
*   Please fill out feature repo checklists [Caleb]
    *   Container runtime interface
        *   By defautl enabled for 1.6 for 2 weeks
        *   Found several small issues, such as kubectl log+journald support
            *   The fixes are in or pending
            *   Discussions about requirements to deprecate support for journald: https://github.com/kubernetes/kubernetes/issues/42188
        *   Non-CRI test coverage:https://k8s-testgrid.appspot.com/google-non-cri
    *   Evictions
        *   Delete pod objects only after all pod level resources are reclaimed
        *   TODO:
            *   Avoid evicting multiple pods unnecessarily - This is a bug in eviction manager which is causing frequent eviction node e2e test flakes. https://github.com/kubernetes/kubernetes/issues/31362
    *   Preemption & eviction for critical pods handling on node
        *   All prs are merged required for 1.6
    *   Alpha support for multiple GPUs
        *   PR #42116 was merged on 2/28
    *   Node allocatable, QoS cgroups and pod level cgroups.
        *   PR to enforce the node allocatable was merged, but not enforced by default yet. Gated on pod level cgroups.
        *   Qos level update PRs are ready. Waiting to be merged. Delayed due to dependency on node allocatable.
        *   TODO:
            *   Merge QoS cgroup PRs
            *   Enable QoS and Node allocatable cgroups by default - set `--cgroups-per-qos=true`
            *   Enable Node Allocatable enforcement on pods - set `--enforce-node-allocatable=pods`
            *   Evict based on Node Allocatable (more of a bug fix)
                *   PR: https://github.com/kubernetes/kubernetes/pull/42204
    *   Node problem detector
        *   Announce it as beta for OSS k8s
        *   Support journaltd
        *   extended to support arbitrary logs
        *   standalone NPD mode
        *   Enable it for GKE GCI nodes by default for


# Feb 21 {#feb-21}



*   Cancelled due to no meeting week


# Feb 14 {#feb-14}



*   Vish: node allocatable phase 2 for v1.6
    *   https://github.com/kubernetes/community/pull/348
*   Euan: shared mount propagation
    *   https://github.com/kubernetes/community/pull/193
*   Position on Docker CVE
    *   https://github.com/kubernetes/kubernetes/issues/40061
    *   Docker 1.12.6 is tested in node e2e against 1.6 w/o CRI.
    *   Plan to have docker 1.12.6 test included in 1.5 branch
    *   Not validate overlay2 graph driver
    *   Some tests at node-e2e
*   Explicit Service Links [proposal](https://github.com/kubernetes/community/pull/176)


# Feb7 {#feb7}



*   [virtlet](https://github.com/Mirantis/virtlet) & CRI proxy demo (with CNI and out-of-process docker-shim)
    *   Uses CRI to run docker shim and virtlet to allow both docker images and virtlet images to be run.
    *   Started using daemonset
    *   Allows "plain" pods to communicate with VMs, VMs can access services


# Jan31 {#jan31}



*   k8s 1.6 - docker version requirement
    *   Derek put here since he cannot make 1/24, but I am fine with 1.12+
    *   Redhat validates overlay. Thinking about switch to overlayfs from devicemapper later this year, so validation may not happen right away.
    *   node-e2e has docker 1.12 tests for both CRI and non-CRI implementations.
    *   deprecating docker 1.10.x? support policy?
    *   Manual triage the issue with docker 1.13.x. Who?
*   pod level cgroup rollout plan
    *   See: https://github.com/kubernetes/community/pull/314
*   Containerd presentation by Patrick Chanezon & Stephen Day
    *   [slides](http://www.slideshare.net/chanezon/docker-containerd-kubernetes-sig-node)
    *   [containerd summit Feb 23](https://docs.google.com/forms/d/e/1FAIpQLSeYK9_DaFJvF8PtyykUzm3awV3e1xHwuonxbKvak9UYS8VnqQ/viewform?c=0&w=1) at Docker office in San Francisco
    *   [containerd livestream recap](https://blog.docker.com/2017/01/containerd-livestream-recap/)
    *   [containerd dev reports](https://github.com/docker/containerd/tree/master/reports)
*   Update critical pods handling
    *   1.4 & 1.5 changes:
        *   critical pods annotation can only apply to kube-system pods, and feature gate flag is added and default is off
        *   critical pods won't be evicted.
    *   1.6 improvements:
        *   Introduce Node-level preemption
*   Shared host volumes (again)
    *   [related proposal](https://github.com/kubernetes/community/pull/193)
*   Explicit Service Links [proposal](https://github.com/kubernetes/community/pull/176)
*   [Local storage proposal](https://github.com/kubernetes/community/pull/306)
    *   please look at high level PR and comment on use cases


# Jan24 {#jan24}



*   Shared PID Rollout Plan (@verb)
    *   [Proposed](https://github.com/kubernetes/community/pull/207) implementing by default in CRI
    *   Everyone ok with kubelet flag for rollback?
        *   Agreed.
    *   Proposal is approved.
*   Pod Troubleshooting update (@verb)
    *   [Updated Proposal PR](https://github.com/kubernetes/kubernetes/pull/35584) with items discussed in Jan10 meeting
    *   Open Questions: pod cgroup, security policy, admission control
*   k8s 1.6 - docker version requirement
*   rktlet status: approaching the point where it can be merged back as rkt-shim and replace current rktnetes
    *   Ref: https://github.com/kubernetes/community/blob/master/contributors/design-proposals/kubelet-rkt-runtime.md#design
    *   Reasons: according to initial design proposal, circular dependency of grpc proto (in kubelet) -> rktlet -> kubelet, node-e2e testing setup and pr-blockers almost impossible to do out-of-repo
    *   Alternate solution to the originally accepted proposal of vendoring rktlet into kubelet due to above reasons
        *   Shoutout to Vish for being on point with his concerns in the original proposal (link missing)
*   community/proposal: change default propagation for host volumes \
https://github.com/kubernetes/community/pull/151  \
https://github.com/kubernetes/community/pull/193
    *   Vish, Dawn added as reviewers
    *   Euan: will add a summary comment for basically the current status/thoughts
*   Benchmarking protobuf3 / grpc / without
    *   protobuf3 did reduce mem usage by a small amount
    *   Further results still pending


# Jan17 {#jan17}



*   CRI rollout plan review
    *   https://docs.google.com/document/d/1cvMhah42TvmjANu2rRCFD3JUT0eu83oXapc_6BqE5X8/edit?usp=sharing
    *   The newer, more succinct [rollout plan](https://docs.google.com/a/google.com/document/d/1b6MLZWUjKV7uJ8BxK-4ciFJl3IPYpsPPfiaHVKV4j-0/edit?usp=sharing)
    *   CRI Integration Performance Benchmark Result: https://docs.google.com/a/google.com/document/d/1srQe6i4XowcykJQCXUs5fFkNRADl_nR4ffDKEVXmRX8/edit?usp=sharing
*   CRI: upgrade to proto3?
    *   https://github.com/kubernetes/kubernetes/issues/38854
*   Plan to revert recent changes for eviction/qos
    *   https://github.com/kubernetes/kubernetes/issues/40024
    *


# Jan10 {#jan10}



*   Update on shared-pid namespace for docker (@verb)
    *   Question from PR: Do we want the CRI to require a shared pid namespace for all run times?
    *   Dawn: Agreed. The only concern is docker 1.11 usage, docker being the only alpha impl right now
    *   Euan: we should just mandate docker 1.12 to make docker compliant somewhere in the timeline then
    *   Outcome: PR to make it mandatory in CRI api, update on the proposal the plan a bit
*   Update on "pod troubleshooting" proposal (verb)
    *   Update to a running pod with a run-once container
    *   https://github.com/kubernetes/kubernetes/pull/35584
    *   Timeline?
        *   Not part of the 1.6 roadmap, this is longer term
        *   This also isn't wholly sig-node; if there's a kubectl-only one, it's not really a node-thing as much (api-machinery? apps?)
*   [virtlet](https://github.com/Mirantis/virtlet) demo
*   "core metrics" proposal: https://github.com/kubernetes/community/pull/252
    *   Per-pod usage and container are included
    *   Pod overhead: whether that's shown or how is an open question
    *   Derek: no overall aggregate usage (e.g. all BestEffort)
        *   Just aggregate / cross-reference with the podspec for that usecase
        *   This should be minimal/low-level
    *   Comment on the proposal!
*   CRI rollout proposal
    *   https://docs.google.com/document/d/1cvMhah42TvmjANu2rRCFD3JUT0eu83oXapc_6BqE5X8/edit?usp=sharing
    *   We need to get proper feedback on this for 1.6 to make sure the timeline is sane and start the rollout
*   Euan - Going to be less involved in sig-node (so long and thanks for all the fish!)
    *   I will finish up existing open issues / prs I have in flight
    *   Luca (@lucab) is the best point-of-contact for rktlet/CRI work
    *   Dan (@ethernetdan) will be getting more involved in sig-node as well (including helping with the 1.6 release and helping keep track of features we're involved in, etc).
*   Caleb: PM-rep for a sig
    *   A new role responsible for tracking features for a given SIG, work
    *   https://docs.google.com/document/d/1ZElyRqNsGebvMpikEBhSuf2Q6WrEYHB2PgWIlGgwt78/edit# (perms to be updated by Aparna)


# Jan 03 {#jan-03}



*   [Q1, 2017 node planing](https://docs.google.com/spreadsheets/d/1-hADEbGEUrW04QP4bVk7xf1CWuqbFU6ulMAH1jqZQrU/edit?ts=585c5905#gid=1369358219&vpid=A1)
*
