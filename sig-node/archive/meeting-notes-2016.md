# sig-node weekly meeting

# Dec 13 {#dec-13}



*   virtlet demo
    *   Piotr from https://github.com/Mirantis/virtlet
    *   a Kubernetes runtime server which allows you to run VM workloads
    *   reschedule to next time
*   cri-o demo
    *   Antonio from redhat
    *   pod and container lifecycle management works, and image service implementation is done
    *   start integrating with Kubelet
*   CRI rollout and planning
    *   1.5 alpha API is ready
    *   in 1.6 using CRI and docker implementation in production
        *   validate the API in productiont
        *   Rollout plan/proposal: https://docs.google.com/document/d/1cvMhah42TvmjANu2rRCFD3JUT0eu83oXapc_6BqE5X8/edit?usp=sharing
        *   Introducing flag to enable / disable the feature
        *   Backward compatibility: container name, etc. Requires draining the node.
        *   Rollback planning.
*   Resource management workgroup
    *   Started from Jan 3rd, and expected is dismissed once the roadmap and planning is done.
    *   preemption & eviction priority and schema
*   Will CRI shim for docker support exclusively CNI for networking?
    *   The shim currently supports cni and kubenet
    *   There is a [PR](https://github.com/kubernetes/kubernetes/pull/38430) up for allow the native docker networking
    *   https://github.com/kubernetes/kubernetes/issues/38639 discusses the support for the exec network plugin


# Dec 6 {#dec-6}



*   Will CRI support exclusively CNI for networking?
    *   CRI right now with dockershim's impl only supports CNI
    *   Next step is file a github issue or slack ping probably since we didn't get a clean answer here
*   Additional CNI question: no distinction between networks for workloads
    *   same with/without CRI: networking is set up once per-node
    *   better question for sig-networking
    *   If required, CRI can evolve as well.
*   garden team: demo
    *   Julz
    *   Garden, Cloud Foundry's Container Runtime
    *   Given CRI / cri-o / rktlet, seems similar to garden
    *   Plans to integrate into CRI/K8s effort?
        *   Hope to share ideas, code, whatever due to similarity; opportunity to probably share ideas and code
    *   Many links available related to this:
        *   https://github.com/cloudfoundry/garden
        *   https://github.com/cloudfoundry/guardian
        *   https://github.com/cloudfoundry/grootfs
*   docker 1.12 liverestore (https://docs.docker.com/engine/admin/live-restore/) should it be enabled by default?
    *   Related to disruptive updates; if we wish to be less disruptive
    *   CoreOS does not enable it currently, but @euank doesn't know if there's a specific reason
*   Shaya/AppOrbit, Infrantes public cloud CRI demo
    *   Enables pods to be isolated within independent VMs
        *   differs from hyper in that the VMs can be separate public cloud instances where nested virt isn't supported
    *   Enables orchestration of full VM images that aren't running containers but act as Pods to the k8s cluster
    *


# Nov 29 Agenda {#nov-29-agenda}



*   FYI: [Shared pid namespace](https://github.com/kubernetes/kubernetes/issues/1615)
    *   No discussion needed
    *   [PR for rollout sent](https://github.com/kubernetes/kubernetes/pull/37404) for review
    *   brendanburns suggests we consider supporting isolated namespaces indefinitely
*   rkt attach demo
    *   Implemeting the design proposed in [#34376](https://github.com/kubernetes/kubernetes/blame/master/docs/proposals/kubelet-cri-logging.md#L220-L225).
    *   Addresses a problem of the pre-CRI rkt implementation; support `kubectl attach` and `kubectl run -it`
*   Dawn: Some ideas for sig-node next year:
    *   Node level debugability
    *   Node management / availability (daemon update, repair, etc)
    *   Node allocatable rollout (e.g. daemon overhead iiuc?)
    *   CRI, validation test, beta/stable
    *   Checkpointing (finally)
    *   Tackle logging story
    *   Auth/security between daemons / auth between applications
    *   Resource management (related to many of the above too, yeah)
        *   Mostly for reliability, not efficiency
        *   pod overhead, etc
        *   Better guarantees for performance / etc node
        *   Disk management
    *   Final: kubelet as a standalone thing/"product"
        *   Checkpointing, node level api and versioning
*   virtlet
    *   https://github.com/Mirantis/virtlet
    *   As of now it can use kvm/qemu to run images
    *   CRI implementation
    *   Demo forthcoming


# Nov 22 {#nov-22}



*   Announcements & sync-up
    *   Derek: Putting together a workgroup that reports back to sig-node for resource management. Specifically to allow management of resources outside the kubelet for exploratory work, identify ones that should be managed by kubelet.
    *   Look for an announcement sometime next week
*   Status of CRI on 1.5
    *   v1alpha "done" for docker, community can and should try it and give feedback
*   Shared pid namespace:([verb@google.com](mailto:verb@google.com))
    *   First step is to make infra container reap zombies
    *   https://github.com/kubernetes/kubernetes/pull/36853
    *   But will infra container even be around for all run times in the future?
    *   Yes please!
    *   First step is the pause container as an init process
    *   Other runtimes already handle that (e.g. rkt and cri-o)
    *   On by default?
        *   Some containers assume PID 1 (e.g. for exec kill -SHUP 1, or in their /entrypoint.sh messy pseudo-init).
        *   Some containers also bundle init systems
    *   Dawn: there was discussion about having our own init for a pod
        *   rkt, pause container, cri-o all have init processes. infra is a bit of a hack and docker specifc, but we should be able to get rid of those.
        *   For now, his change makes sense, just go with it, and we can consider long-term unification in parallel/later
*   Backwards compatibility concerns:
    *   disruptive CRI
    *   disruptive cgroup rollout
    *   Have we done disruptive rollout before?
        *   In GKE do that.
        *   Openshift: drain nodes before said updates do that.
        *   In the past, maybe docker labeling broke this? No specific memory.
    *   Currently planning to make both of those disruptive
        *   Euan: Action item, double check this is sane from CoreOS side
*   Hi from Cloud Foundry garden team! (What can we share, how can we help?)
    *   Next call, maybe demo and talk about garden a little
*   rkt roadmap:
    *   1.5 continuing to work to get attach and e2e happy (might bleed into 1.6)
    *   1.6 alpha release and recommend general "tire-kicking"
*   CRI status, api interface. Rollout in 1.5, alpha api, what does that mean?
    *   Cluster api, we don't suggest prod usage because it might change
    *   This is internal, so it's different, right? Compatibility is an internal detail, not external/user.
*   Will CRI support exclusively CNI for networking?
    *   Furthermore, is the network config missing from the CRI?
        *   Maybe? It's alpha
    *   Come back next week
*   1.6 roadmap planning
    *   Community meeting talked about reliability.
        *   Resource management
        *   disk management
    *   Lots of "in-flight" features which are not marked stable yet, have not been seen through.
    *   Use 1.6 release to "finish work" and rollout
        *   CRI
        *   Pod Cgroup
        *   Node allocatable
        *   …..
    *   Part of "nodespec work"
    *   Focus on reliability and rollout of features. Finish node level testing.
    *   Let us/Dawn know about other potential items for the roadmap.
    *   Expected date? presumably before 1.6 :) TBD


# Nov 01 {#nov-01}



*   Image Exec (verb@google.com)
    *   Better support for containers built from scratch
    *   <code>kubectl exec -it -m <em>image_name</em> <em>pod_name</em></code>
    *   Proposal: https://github.com/kubernetes/kubernetes/pull/35584
    *   Usecase primarily dev cluster only or dev+prod?
        *   Both
    *   Mount namespace vs filesystem view (e.g. proc and dev and so on might differ)
        *   No solution offered for this
    *   Pod lifecycle for this pod
        *   run a pod + exec into a pod
        *   Dangling pod problem with `kubectl run --rm`?
            *   No answer known yet
    *   Display to user:
        *   Is it hidden from regular listing? (e.g. get pod)
        *   Right now there's an idea of separate init container statuses. Will there be a separate debug container construct?
        *   There's been discussion before of "tainting" pods that have been exec'd into, debugged.
    *   Resources? This can add additional resource needs. Do we reserve things? Push it onto the user?
        *   Derekwayncarr: Size pods in advance to be able to be debugged
    *   Cleanup / reliability:
        *   Fairly intrusive..
    *   Security?
        *   Image whitelist / blacklist?
            *   That's being added, but now needs to be added in one more place
            *   Admission controllers will need to tie in too
        *   one implication: Userns + host bindmounts, depending on how userns is implemented, could be messy (hostmount implies no userns, but userns relabeling might be enabled for the whole pod)
            *   No answer to this.
        *   Does SELinux interact with this? Do we need to relabel container rootfss in a scary way?
            *   No answer for this.
        *   Concern about how this interacts with introspection tools RH has that e.g. scan podstatus for image IDs
    *   Alternative: Allow modifying the podspec to run a new container?
        *   Dawn: Idealy, but that has a problem of needing more upstream changes, can't just be kubelet
*   Kubelet pod API, initial version
    *   New kubelet API for bootstrapping
    *   https://docs.google.com/document/d/1tFTq37rSRNSacZeXVTIb4KGLGL8-tMw30fQ2Kkdm7JQ/edit?usp=sharing
    *   Minimal guarantees in this implementation ("best-effort")
    *   Differences from static pods:
        *   These are persisted into the api-server whereas static+mirror are "fake"/read-only
        *   These give you a "real" r/w copy in api-server
    *   Other option is bootkube. Temporary control plane with a "real" api-server, client transitions between em
        *   Complexity, painful to maintain the code
    *   This implementation adds a new pod source of an API
    *   Due to security concerns, it would be a default-off api, potentially listening on a unix socket
    *   Will create pods before the api-server is connected to. Will move them to the api-server when able to
        *   Api-server pod persistence will result in a pod restart to fixup UIDs effectively
    *   Derekwayncarr: We want to get rid of the existing ways; what other alternatives are there?
        *   The bootstrap/run-once kubelet was shot down. And this has some other nice properties as well (e.g. disaster recovery)
    *   Not a full api-server in the kubelet though. essentially only nodename stored pods
    *   Derek: Will this be versione[/1B856NU1Ie0Pid4xGV2D9QZUJhUD24QzYsJYqH0yJU8A/edit#gid=0](https://docs.google.com/spreadsheets/d/1B856NU1Ie0Pid4xGV2D9QZUJhUD24QzYsJYqH0yJU8A/edit#gid=0)
    *   No sig-node weekly meeting on the 8th dd, will it be pretty and hygenic? Do we distribute clients?
        *   We don't have answers for those
        *   Maybe just experimental for now and not tackle those?
            *   Derek: Concerned with experimental features that aren't as well thought out
    *
*   Demo: rkt with runc (@casey)
    *   Does runc list
*   Kubecon f2f details
    *   Someone should add any details we do know to  https://docs.google.com/spreadsheets/due to kubecon


# Oct 25 Agenda {#oct-25-agenda}

NOTE: This meeting is being held in ZOOM due to troubles with hangouts: http://www.zoom.us/my/sigwindows



*   Conflict with DiskPressure and Volume cleanups
    *   https://github.com/kubernetes/kubernetes/issues/35406#issuecomment-256101016
    *   Suggestion: We should be always cleaning up memory-backed volumes, but at minimum we need to start doing that on eviction
    *   Prioritization? Red Hat is willing to help fix this for 1.5 since it's causing real pain for them with secrets + eviction
    *   Possible problem: Kubelet falls over, pod is removed, kubelet comes back up, volume manager can't find from the state whether an emptydir is tmpfs trivially.
        *   Doesn't matter, volume manager should still clean it up, unmount, what have you? Dig into it a bit more and comment on that issue
*   Image Exec (verb@google.com)
    *   Better support for containers built from scratch
    *   <code>kubectl exec -it -m <em>image_name</em> <em>pod_name</em></code>
    *   Draft proposal: http://bit.ly/k8s-imageexec
    *   verb@ to open a PR to main k8s repo
        *   DONE: https://github.com/kubernetes/kubernetes/pull/35584
    *   tl;dr; Run another container image in the same namespaces as the pod; Use the other image to debug the pod.
    *   Expected to be reviewed this week
        *   Probably discussed more next week in sig-node as well
    *   Some questions about how to properly deal with viewing its mount namespace
        *   Post them on the proposal!
*   Quick rktlet demo (init containers & adding bindmounts)
    *   Demo will happen next week :)
*   rktlet status
    *   Vendoring in in-progress (couple dependency issues); to be followed by e2e testing
    *   Attach & logging
*   Kubelet in a chroot https://github.com/kubernetes/kubernetes/pull/35328
    *   CoreOS has been shipping kubelet in a chroot for the last half year
    *   Volumes might require path remapping
    *   Does not concern most k'let developer
*   Kubelet mounts volumes using a container
    *   Using rkt fly to run a mounter image to setup k'let volumes.
    *   Stop gap solution until k'let in a chroot lands


# Oct 18 {#oct-18}



*   CRI exec/attach/streaming
    *   https://docs.google.com/document/d/1OE_QoInPlVCK9rMAx9aybRmgFiVjHpJCHI9LrfdNM_s/edit?ts=5800133e# (?)
    *   Original issue: https://github.com/kubernetes/kubernetes/issues/29579
    *   TTY vs. streams, to be tackled offline
*   Per-pod overhead
    *   exec, logs, supervision
    *   where to account? kubelet level? pod level?
    *


# Oct 11  {#oct-11}



*   Reducing privileged components on a node
    *   Can other components (kube-proxy) delegate privileged host operations to kubelet? (e.g. firewalld operations)
    *   Dawn: Opinion, that makes kubelet more monolithic. It's the main agent, but it should be able to delegate. Preference for moving things to separate plugins where reasonable.
    *   Euan: Counterpoint, multiple things talking to api-server has some extra auth problem
    *   Is kube-proxy actually core? Out of scope of this sig :)
    *   Minhan: Note that kube-proxy is an implementation detail; over time it will potentially differ.
        *   This discussion is also about more than just kube-proxy
*   Pod cgroups (demo update - [decarr@redhat.com](mailto:decarr@redhat.com))
    *   q on default: Can't it default to 'detect' and try docker info, and if it doesn't have info fallback to cgroupfs?
        *   It does do the right thing for docker integration, but document says the wrong default :)
    *   Note: Only works for systemd for 229+ because of opencontainers slice management stuff
    *   Upgrade of existing node to this feature?
        *   Evacuate the node first. We don't support 'in-place', it's not in-scope
    *   We don't currently tune memory limits over time either
    *   Some docker-things are not charged per-cgroup (like docker's containerd-shim for example)
        *   Also not in-scope; upstream changes
    *   Euan: Will also look into making sure the systemd driver works well for rkt. It should work with effectively `systemd-run --slice=$parent` under systemd driver
    *   Yuju: We can add the info call to give that info to CRI, just do it :)
*   When can we assume pod-cgroups exist, e.g. for eviction and so on?
*   CRI logging proposal: https://github.com/kubernetes/kubernetes/pull/34376
    *   q: This is only under CRI/experimental, right? Yes, it's part of an experimental flag, default should not differ
*   CRI networking: https://github.com/kubernetes/kubernetes/pull/34276
    *   Original issue: https://github.com/kubernetes/kubernetes/issues/28667
    *   Who should own it? Kubelet or runtime?
    *   How about the configs (e.g. plugin dirs etc)
        *   Freehan: Eventually move to part of the runtimes, deprecate kubelet flags
    *   Will kubenet still exist? Only CNI?
        *   Eventually it'll be a cni plugin perhaps
    *   sig-networking discussed this some
        *   Some considered out-of-band vs passing through all networking stuff
    *   In the future, higher level "Network" objects might exist. Already, networking exists as a kubernetes construct to a degree.
        *   CRI will have to grow for this to include a good chunk of that.. or out-of-band
        *   In the future, the 'UpdateConfig' object might expand and these objects will have to be somewhat runtime-implemented
    *   CRI will hve to include tc, etc type stuff so that the runtime can apply network shaping
    *   There's also the implicit assumption that networking metrics aren't core when they're moved out of kubelet maybe
    *   Conclusion: Let's roll forwards so we can look at more than just a tiny start, reconvine next week.


# Oct 04 {#oct-04}



*   CRI status updates
    *   Exec/port-forward/attach shim; can we just implement the grpc bit without it due to time concerns? Should we start implementing the shim?
        *   a) Implement internally for now is okay as docker did (yes)? The shim-library will provide a fairly similar interface
    *   Logging
        *   Try to draw conclusion this week.
    *   Monitoring
        *   Use cadvisor for now, punt till post 1.5
    *   Docker integration status:
        *   Passing most of the tests (testgrid link)
            *   node e2e: https://k8s-testgrid.appspot.com/google-node#kubelet-cri-gce-e2e
            *   cluster e2e: https://k8s-testgrid.appspot.com/google-gce#gci-gce-cri
        *   Adding serial and per-PR builders soon
        *   Lantao (random-liu@) is working on integration over grpc: https://github.com/kubernetes/kubernetes/pull/33988
    *   Networking
        *   Freehan is working on doing the networking plugin work for CRI
        *   Moving tests to node-e2e, and then will do docker-integration work
    *   Help wanted: CRI validation test
        *   https://github.com/kubernetes/kubernetes/issues/34040
*   rktlet demo
*   KubeCon sig-node F2F tentaively planned for 11/7. Kindly respond to the [existing thread](https://groups.google.com/forum/?utm_medium=email&utm_source=footer#!msg/kubernetes-sig-node/t1S77NwujlE/ybK-UtvaAAAJ) if you can make it. Video Conferencing can be setup for remote attendees.
*   Node-e2e; should it use kube-up-like environments instead of what it has now?
    *   Provides benefit in that you're testing a more "production" environment
    *   Could at least share the startup scripts fairly easily
    *   If we had a standardized "node-setup" shared piece at the bottom, then we could more easily test/validate by knowing there's this one way.
    *   Current kube-up setup has duplicated code and burdon and it makes it tricky to claim x distro is supported. Goal is to make it easier to add new distros and maintain better.
    *   Share more code at the node level for setup. Document the exact steps needed for setup in general.
*


# Sept 27 {#sept-27}



*   Change the way of expressing capability in API (cri first)
    *   Currently the default is just "whatever the runtime does"
    *   This proposes moving it to kubelet setting an explicit default for add/remove and tell the runtime *exactly* what to set, not a set to add and remove from its ad-hoc set.
    *   This sets the path forwards to potentially exposing it to the external API, but that's out of scope for sanity reasons.
    *   https://github.com/coreos/rkt/issues/3241
    *   Current it's docker specified, with [a default list + add_caps - remove_caps](https://github.com/docker/docker/blob/eabae093f4df2d6cd3f952131ae7109d92480674/daemon/oci_linux.go#L209-L222)
    *   It's preferred to change to a [white list based interface](https://github.com/yifan-gu/rktlet/blob/2e3a9b77d32e765da5a3e557009607bddeb4e362/rktlet/runtime/helpers.go#L225-L238) at least for the CRI. We are not gonna touch the k8s API for now
    *   Propose that CRI is provided an explicit set of capabilities, not an add and remove set. Normalizes runtimes and behavior
        *   Take the default from docker, merge in what was changed
        *   Makes sense probably, "PRs welcome"
    *   Long term, maybe a cluster policy or pod-security-policy for default set
    *   **Action**: Yifan (or other) propose / issue to discuss further based on what we have here.
*   rkt CRI demos:
    *   isolators in rkt app add (sur)
    *   port forwarding (casey)
*   CRI e2e testing
    *   https://github.com/kubernetes/kubernetes/issues/33189
    *   https://k8s-testgrid.appspot.com/google-node#kubelet-cri-gce-e2e
        *   Bypassed CRI to support logs/exec for now
        *   Support only docker 1.11+
    *   Q: If you're writing your own CRI implementation, how can you test it?
        *   Will we setup a conformance a CRI conformance test?
        *   No plan yet, for now we can use node-e2e test, and once we have more cycles we can create something specialized/better
    *   Adding new tests and testing external things is hard and complicated (lots of discussion)
    *   Talk to sig-testing!
*   node-e2e testing discussion (what are they, etc):
    *   How they're run: https://github.com/kubernetes/kubernetes/blob/master/test/e2e_node/jenkins/e2e-node-jenkins.sh
    *   There's also a make target for it https://github.com/kubernetes/kubernetes/blob/master/docs/devel/e2e-node-tests.md#running-tests
    *   Node conformance is a separate thing, it's a contaienrized test framework target for validating whether a node meets the minimum requirement of a Kubernetes node.
        *   https://github.com/kubernetes/features/issues/84 (see links in there)
        *   Docs of current version http://kubernetes.io/docs/admin/node-conformance/ (We've pushed the alpha version image onto the registry, currently several volumes need to be mounted into the container, we'll simplify that in newer version)
        *   Pending PRs:
            *   system verification: https://github.com/kubernetes/kubernetes/pull/32427
            *   containerize node conformance test: https://github.com/kubernetes/kubernetes/pull/31093
*   CRI logging discussion continuation
    *   https://github.com/kubernetes/kubernetes/pull/33111


# Sept 20 {#sept-20}



*   F2F meeting after KubeCon - Kindly respond to sig-node email thread [here](https://groups.google.com/forum/#!topic/kubernetes-sig-node/t1S77NwujlE)
*   Discuss SELinux and pod level vs container level SELinux context
    *   Should this context be moved to the sandbox level? https://github.com/kubernetes/kubernetes/blob/master/pkg/kubelet/api/v1alpha1/runtime/api.pb.go#L1238
    *   Derek Wayne Carr to defer to Paul Morie
    *   From Yuju in slack: @philips there are `PodSecurityContext` and  per-container `SecurityContext` in k8s API. The per-container security context takes precedence over the pod-level one. SELinux is included in both security contexts.
    *
*   TTY Attach progress update
    *   https://github.com/kubernetes-incubator/rktlet/issues/8#issuecomment-248000500
    *   Luca still digging into technical details of the forwarding and working on the prototype that he has working (see [WIP code](https://github.com/lucab/rkt/commit/cca63145d01641e3da75d029160725fd47013bff))
    *   Need to figure out how we encode and proxy resize, etc. Any ideas?
    *   Derek to ask Mrunal to review for ocid
*   minikube rktnetes demo (sur)
    *   https://github.com/coreos/minikube-iso/#quickstart
    *   https://github.com/kubernetes/minikube/issues/571
*   Demo rktlet + new cri interface rkt (yifan)
*   Logging in CRI @tmrts [(https://github.com/kubernetes/kubernetes/pull/33111](https://github.com/kubernetes/kubernetes/pull/33111))


# Sept 13 {#sept-13}



*   ocid demo / status update
    *   why does this still have an infra container?
        *   to support init systems in pods in future (?)
*   reminder: cAdvisor release will be cut 9/15
    *   https[://github.com/kubernetes/kubernetes/issues/27097#issuecomment-237422351](https://github.com/kubernetes/kubernetes/issues/27097#issuecomment-237422351)
*   Discuss "Expose Canonical Image Identifier" proposal path forward
    *   https://github.com/kubernetes/kubernetes/pull/32159
    *   Red Hat has asked for closure on this design this week
        *   RepoDigest is the preferred ImageID when present is key point
*   rkt CRI demo/status update (sur)
    *   WIP at https://github.com/coreos/rkt/pull/3164
    *   kubernetes integration at https://github.com/kubernetes-incubator/rktlet/pulls
*   rkt attach status update (lucab)
    *   https://github.com/kubernetes-incubator/rktlet/issues/8


# Sept 06 {#sept-06}



*   Pod Container Status ImageID (derekwaynecarr DirectXMan12) https://github.com/kubernetes/kubernetes/issues/31621
    *   Right now we run by tag / sha. When you run by tag, you lose auditability. There's a local image id, but that image id is *not* the content-addressable ID
    *   ContainerStatus ImageID is thus not really actually useful. At all.
    *   Should users be able to determine exactly the image running? A) obviously yes, but you can't now
    *   History: Why do we have ImageID?
        *   Based on the 'ssh into a node, docker ps' stuff mostly
    *   Docker does not provide the sha when you pull by tags in a structured way, just a status text (what you see printed when you docker pull foo:tag)
    *   Docker does not always populate repo digest
    *   Possible solution: Have kubelet resolve tag -> digest itself
        *   Downside, kubelet is now relying less on docker and has to do more
    *   Maybe an admission controller could translate?
    *   What do we do with the existing field? Do we replace it with hash, or do we have to add a new field?
    *   Should it be lockstep across all nodes; resolve it before hitting kubelet?
        *   I don't think we can because of RestartPolicy + ImagePullPolicy interactions. Unrelated.
    *   This issue is *just informational*; ContainerStatus tells the truth of the current running reality, no more, no less, no other changes.
    *   Discuss on the issue
*   User-namespace
    *   https://github.com/kubernetes/kubernetes/pull/30684
    *   Was not included in 1.4 as a "late exception"
    *   Discussion about whether this should be done differently, what the process really should be..
    *   TODO, we have some idea of what it should be, @derek to writeup / provide a link
    *   More inclusion of the community would help in terms of transparency as well
    *   Push for userns in 1.5 :)
*   Pre-review of node-team 1.5 features
    *   https://docs.google.com/spreadsheets/d/1nBneiObfPHZwMrsIuvgCSWDm90INdxSJRUFUf11KLhY/edit?usp=sharing
*   Node-e2e flakes a bit more on CoreOS than other images
    *   Should we be running these tests for all distros? Is there really value in blocking per-pr?
    *   AI: dawnchen: file the issue to carry the discussion
*   Rktnetes sync notes:
    *   Moving to rktlet repo and re-vendoring to kubelet
    *   WIP CRI implementation on rkt side
    *   Encourage creating issues about rktnetes in rktlet repo
        *   We recognize issues will still be filed against the main repo, and we don't want to proactively move issues because github doesn't support that, but if we can start with them there, that's easier to triage for us.


# August 30 {#august-30}



*   Node conformance test suite:
    *   https://docs.google.com/presentation/d/18NbozIL22BM4RI2jD7sdWG5DphYpa6pLvTOrF6ROYJM/edit#slide=id.g165a166f0d_0_0
    *   https://docs.google.com/spreadsheets/d/1yib6ypfdWuq8Ikyo-rTcBGHe76Xur7tGqCKD9dkzx0Y/edit#gid=0
    *   How much will this validate container-runtime configuration? E.g. userns settings
        *   Container runtime validation tests, which right now is really docker validation. We don't have CRI to run validation against
        *
*   Kubelet distribution -- Build + deployment flags for different kubelet varients? No varients?
    *   Context: runtimes, cadvisor, the like; vendored dependencies
    *   Concerns (Dawn):
        *
*   rktlet: In process vs out of process, opinions, reasons? https://github.com/kubernetes/rktlet/pull/3#issuecomment-242180549
    *   Philips: Discussion around being consistent for all CRI being in or out
    *   Vish: I think the discussion is really about the longterm end state of the CRI and deployment model
    *   Concerns:
        *   Client-server because we have to support many
        *   More clear-cut interface
        *   Better release philosophy maybe
    *   Dawn: Push discussion to a PR that was just posted
*   Discussion: Ideas to improve debugging kubelet in production (@derekwaynecarr)
    *   Kubelet wedged, etc. Verify difficult to debug
    *   Proper audit / trace logs for debugging would be great
    *   ref: https://github.com/kubernetes/kubernetes/issues/31720
    *   the life of a pod needs better logging to point out a specific issue
    *   Create a doc with some of these ideas, options
    *   Kubelet socket to query debug info from?
        *   There's an http api for much of this information
    *   Related: Tool to improve introspecting a node, e.g. "what pods are running, what do they look like"
*   Host user namespaces in 1.4 (https://github.com/kubernetes/kubernetes/pull/31169)
*   oci integration status update
    *   Create start stop cycle stuff works in a basic way. Missing tty (coming soon™)
    *   Demo in a week or two


# August 23 -- Cancelled {#august-23-cancelled}



*   Node conformance test suite
*   Minukube + rktnetes, demo @sur (move to next week as well)


# August 16 {#august-16}



*   Node performance benchmark (zhoufang)
    *   Slides link: https://docs.google.com/presentation/d/1pYNnKo7OF-IHOwnSJ1hZvKmEwzEKc2y3IoZIOjYcDK4/edit
    *   Do we have a repo / issue / PR we can follow for running this on our own?
        *   Not yet, more stuff needs to be merged and so on first. TODO, make and link a tracking issue (maybe a future sig-node)
    *   Will this be integrated with the existing testing dashboards / gubernator stuff?
        *   For now pushed to GCS, in the future talk to infra team
    *   Is there an option to push this to prometheus? Other sigs have made it possible to support e2e -> prometheus pushes
        *   Not right now? We don't have a real answer right now
    *   Long term, we should be alerting when there are regressions shown in this, it will be automatically run
    *   How does this work?
        *   Standalone cadvisor
        *   It should already be mostly runtime agnostic
    *   Yifan and Zhou to sync on how to run this for rktnetes to get pretty results there :)
*   CRI: Is there a way to show a runtime's cgroup?
    *   Not at this moment afaik; it would make sense to add it as part of the 'Version' api
    *   Should we add a cgroup path for the runtime's cgroup(s?) to the kubelet?
        *   Vish: The kubelet cares about pods and the node as a whole, why do we care about this?
    *   Derek: Open an issue for this, more discussion https://github.com/kubernetes/kubernetes/issues/30702
*   CRI attach/exec/portforward
    *   https://github.com/kubernetes/kubernetes/issues/29579
    *   Milestones / v1 definition for interested parties:
        *   https://github.com/kubernetes/kubernetes/issues/28789
    *   Concrete next step for this issue?
        *   TODO (yujuhong@): Clearly summarize the possible options and tradeoffs so far to improve coherency of the issue, and hopefully be able to decide to go with one?
*   CRI area owners
    *   Dawn brought up that we should discuss area owners to help drive the progress in individual areas.
*   Could we add GPU discovery/assignment capabilities to kubelet in v1.5?
    *   https://github.com/kubernetes/kubernetes/pull/28216
    *   There were still some question I think? But yes, you can start working on 1.5 now, good luck getting reviews :)
*   Should we use annotations to expose sorta security related pod features, specifically sysctls
    *   Sysctls: https://github.com/kubernetes/kubernetes/pull/26057
    *   Current proposal is first-class, not annotations. Concerns around annotations, security…
    *   Vish: Why is validation a concern for an alpha opt-in feature?
    *   (ref app-armor, also annotations)
    *   Maybe have a node-whitelist that is enforced on the kubelet layer. Kinda messy UX, but it should resolve these security concerns
    *   Do we need to have scheduling decisions?
        *   We can use taints-tolerances
    *   Derek: What are the next steps for sysctls? We already know what's accounted / not accounting, how do we decide a whitelist.
        *   Vish: Proposal didn't make this clear enough; it needs the information in a better form.
        *   In the 1.4 timeline: Start with a node whitelist (default empty list of sysctls), and then on a per kubelet basis people can choose what they're okay with
        *   **Vish**: Comment something to the above's effect on the sysctl PR
*   Pod level cgroups: will they land in 1.4 timeline?
    *   Dawn has it marked for 1.5 as a p0.
    *   Vish: Probably won't happen this week
    *   Action Item (?): Disable flags for 1.4 if it's not making it in
*   UsernsMode
    *   https://github.com/kubernetes/kubernetes/pull/30684
    *


# August 9 {#august-9}



*   Add a k8s repo for rkt/CRI integration ? (rktlet)
    *   Conclusion: yes, will be done by Dawn for sig-rktnetes after this meeting
    *   https://github.com/kubernetes/rktlet
*   Add support for UsernsMode?
    *   Redhat wants it for 1.4 and is willing to do the work for it
    *   No need to change the default, just make it possible
    *   Willing to also do podsecuritycontext stuff so administrators can control it correctly
    *   Dawn: Some concerns about whether we can finish it in time since we have a "freeze" coming up so soon. We do want the feature though (:+1:), but we need to follow the rules.
    *   Some technical/semantic issues with userNS since it will break some existing kubernetes features (other namespacing, mounts of volumes owned by root, etc, kube tools?)
    *   Proposal incoming for future discussion with use case
*   Docker validation updates
    *   No manual validation this time. We only have automated docker validation result, it only runs on GCI.
    *   Current test result:
        *   Functionality: We are running node e2e in automated docker validation, currently all green against docker 1.12.0.
        *   Performance: Not yet. We will run kubelet performance test in node e2e against docker 1.12.0 soon.
    *   Like previous release, support multiple docker versions still, and documented the known issues
*   1.4 feature updates
    *   https://github.com/kubernetes/features/labels/team%2FSIG-Node
    *   https://github.com/kubernetes/features/issues/39
        *   OOD handling: getting inode support. No concern
        *   Issues found on devicemapper support
    *   InitContainer to beta / GA
        *   Should we move it to beta / GA, considering that not all runtimes support it now (rkt known issue; will be solved by CRI)
        *   Should there be a procedure for whether a runtime specific feature can be included as GA? Does it need to be expressed in the CRI where it's possible for something to integrate with it? Is there a better answer? Need to have this discussion somewhere
        *   Does there need to be a threshold for number of runtimes support, or api can show it, or what?
        *   Dawn will file an issue to discuss this (ping sig-node)
    *   Node conformance test feature:
        *   NodeSpec: https://github.com/kubernetes/features/issues/56
        *   Node conformance test: https://github.com/kubernetes/kubernetes/issues/30122
*   So far we've improved node-e2e in various ways.
        *   Next step, maybe not for 1.4, package conformance tests as a container image that can run all the tests against that node
            *   Substep 1: static binary
            *   Substep 2: docker image
    *   AppArmor support https://github.com/kubernetes/features/issues/24
        *   On track
    *   per-pod cgroup / qos enforcement on the node
        *   need to get existing open PRs merged!
        *   trying to get a systemd support in next week (once prereqs merge)
        *   full details: https://github.com/kubernetes/kubernetes/blob/master/docs/proposals/pod-resource-management.md#implementation-status
    *   dynamic configure Kubelet https://github.com/kubernetes/kubernetes/issues/27980
        *   On track
*   Status on kubelet/Docker CRI integration
    *   In progress: added a [kuberuntime](https://github.com/kubernetes/kubernetes/tree/master/pkg/kubelet/kuberuntime/) package on the kubelet side to use CRI.
        *   Not tied to any release, WIP still..
    *   Added a shim for docker to implement CRI. Currently this supports only basic features with little test/validation. This is blocked on the kuberuntime implementation for a more complete validtion.
    *   Other parts of CRI are still under discussion (e.g.., [exec with terminal resizing](https://github.com/kubernetes/kubernetes/issues/29579), metrics, logging, etc)
        *   (metrics, comment on [Dawn's summary](https://github.com/kubernetes/kubernetes/issues/27097#issuecomment-237422351) near the end) https://github.com/kubernetes/kubernetes/issues/27097
        *   Need to correct issues with stdin/stdout/stderr/etc streams in the proto files for exec
*   Follow-up on new repo in Kubernetes org for node feature discovery
    *   https://github.com/kubernetes/kubernetes/issues/28311
    *   Process, how do you get an actual conclusion at the end or ownership of the new repo etc. The decision is still not well defined, and the procedure still needs help
    *   Should this be part of the k8s-community meeting or k8s-dev mailing list rather than sig-node?
*   Very short update on sysctls (sttts):
    *   "the table": kernel code-level analysis of sysctls on proposed whitelist:  https://github.com/sttts/kubernetes/blob/bd832c98794bbbdf3618c41e996d74fa091143e5/docs/proposals/sysctl.md#summary
    *   kmem accounting for ipc works until kernel 4.4
    *   broken since 4.5 due to switch to opt-in; probably simple fixes
*   [Remove some kubelet dependencies](https://github.com/kubernetes/kubernetes/issues/26093) (dims) (PR's for [pidof](https://github.com/kubernetes/kubernetes/pull/30002), [brctl](https://github.com/kubernetes/kubernetes/pull/30056), [pkill](https://github.com/kubernetes/kubernetes/pull/30087)) - Do we want to do some of the cleanup for 1.4?


# August 2 {#august-2}



*   Discuss and get feedback on adding [snap](https://github.com/intelsdi-x/snap) as a supported datasource to heapster
    *   There's been discussion of splitting metrics out into "core" and "monitoring"/additional ones
    *   Core ones should be consistent, well understood, defined strongly by kubelet probably
    *   Heapster currently does both core and monitoring. If snap is meant to be in addition to "core" metrics, then that's great, if it's meant to also replace "core" then it needs to be a more involved process.
*   Proactively reclaiming node resources, possibly with an option of administrator provided scripts https://github.com/kubernetes/features/issues/39#issuecomment-235069913
*   Discuss sysctl proposal questions and kmem (sttts and derekwaynecarr) https://github.com/kubernetes/kubernetes/pull/26057#issuecomment-236574813
    *   We can expose sysctls as knobs so long as they're properly accounted for in the memcg
    *   Argument for "unlimited" other than the memcg limits for some of them (e.g. tcp buffers)
        *   Potential issues if applications change its behavior based on specific sysctl values
    *   Further experiment that the whitelist of sysctls are all "safe" to increase; they're all namespaced, but all they all resource-isolated (accounted).
    *   Separate discussion:
        *   Node defaults for sysctls
*   Brief follow up on new repo in Kubernetes org for node feature discovery (Connor/Intel, cross-posted to sig-node from dev).
*   Reminder: Expecting code/feature freeze in 3 weeks for v1.4. Bug-fixes and additions to 1.4 features will be accepted beyond that cutoff, but not new feature impls.


# July 26 {#july-26}

Agenda:



*   Discuss and get feedback on adding [snap](https://github.com/intelsdi-x/snap) as a supported datasource to heapster
*   Shoutout to Lantao for his work on improving node e2e tests :) https://github.com/kubernetes/kubernetes/pull/29494
*   https://github.com/kubernetes/features/issues?q=is%3Aissue+is%3Aopen+label%3Ateam%2FSIG-Node
*   Meeting called short, go have fun :)


# July 19 {#july-19}



*   Last week there was a discussion with CoreOS etc about attach and logging features
    *   [Attach](https://github.com/kubernetes/kubernetes/issues/23335); should we deprecate? Do we have a way at all?
    *   [Logging](https://github.com/kubernetes/kubernetes/issues/27154); `kubectl logs --since` and other bits make it nice to have timestamps. Still up for discussion for how the stdout actually gets logged /  stored. General consensus is 1) have timestamps 2) have it consistent across runtimes
    *   Conclusion tending strongly towards keeping the since feature and having an opinion on the logging format, including timestamps
    *   User experience / tooling that is "kubernetes native", not runtime-specific
        *   Brandon: feature/doc talking about what this might look like
    *   Networking in CRI:
        *   Push down to the runtime vs hybrid between kubelet owning it; more sig-networking
    *   Volumes: Kubelet handles plugins pretty much entirely; continue with this as well.
*   KubeletConfiguration related refactoring and removal of the old KubeletConfig type (https://github.com/kubernetes/kubernetes/issues/27980, https://github.com/kubernetes/kubernetes/pull/29216)
    *   Allow the kubelet to consume a configmap from api-server
    *   Also improve test setups
*   Focus for v1.4 for all sig-node members:
    *   Note: Feature-deadline is this Friday (right?)
    *   https://docs.google.com/document/d/1Ei7uDriZenhFRJQAzYIwEgQm39kR0jCxTHK4qPXQFDo/edit
    *   (CoreOS doc has a node section): https://docs.google.com/document/d/1Jq_U46dSsI5rXAfAAOegtKu2YMoWHhZECvoM3I9vGaM/edit?ts=5775c00a#heading=h.q5z0exowt425
    *   Clean and efficient pod lifecycle extensions in the new container runtime architecture (https://github.com/kubernetes/kubernetes/issues/28789)
        *   As an operator (not user) how do I add an additional label?
        *   Option: configure kubelet to talk to a proxy docker.sock that adds on labels maybe (but how do you get the podspec)
    *   How do we actually track sig-node features in a discoverable way? Too many issues under label:team/node, not enough in feature repo?
        *   Should we just put everytihng in the features repo?
        *   Let's go with that for now
        *   How do we track node specific issues in features repo? Apply a sig-node label maybe?
*   Next week google sig-node will be OOTO
    *   Euan to drive the meeting, possibly a short one due to lack of the Google folks :)
*   Tangent discussion of zoom vs hangouts vs if people can even join without a Googler on the call
    *   External hangout? (a limit of 10 people)
    *   Zoom? (Doesn't work with chromebox and chromebook)
    *   Defer, a Googler (Vish & Tim St. Clair) will join and make sure people joined.
*   Pod Cgroups
    *   https://github.com/kubernetes/kubernetes/pull/26751
    *   https://github.com/kubernetes/kubernetes/issues/27204
    *   Started as a simple feature, but was surprisingly complex?
    *   Learnings to share
    *   https://docs.google.com/presentation/d/13N1ZdCk1Dg_JZGCJhLZTb5j5vtq9RPM31ZIFJ2xu690/edit?usp=sharing
    *   Still pending work to do with things like qos policies, systemd integration, rkt, etc
    *   To hit 1.4, will likely need a new owner
*   Refactoring kubelet code
    *   https://github.com/kubernetes/kubernetes/pull/29216
    *   Shoutout to reviewing that to improve things a bit
    *   Broad tracking issue for cleaning up kubelet actions (does one exist already?)
*   Is my thing a feature for github.com/kubernetes/feature
    *   https://github.com/kubernetes/features#is-my-thing-a-feature

Aside:

    We need some container runtime interface implementation term that is prounancable.



*   oclet (awk-let) rklet (rock-let) docklet (dock-let)

**Action**:



*   Feature owners: File a feature in the feature repository by Friday which at least has a good title
*   Paul: Tracking issue for kubelet cleaning/refactoring


# July 12 {#july-12}



*   https://github.com/kubernetes/features/issues/16
    *   Extensibility for monitoring and logging will be handled by kubernetes
    *   But extensibility from a lifecycle sense is handled by the new runtime API.
    *   https://github.com/kubernetes/kubernetes/issues/28789
    *   Once docker runtime integration has been asbtracted out to use the runtime API, that plugin can be forked to meet this use case.
*   More of a question rather than an agenda item -  Are there any old ideas/pr(s)/issues(s) around Virtual Machines in Kubernetes. (Dims)
    *   https://github.com/kubernetes/kubernetes/pull/17048
    *   https://github.com/kubernetes/frakti
    *   https://github.com/kubernetes/kubernetes/pull/28396/
    *   rkt can also launch pods in lkvm-stage1 today
        *   http://kubernetes.io/docs/getting-started-guides/rkt/#modular-isolation-with-interchangeable-stage1-images
*   Vertical scaling of pods
    *   Hard for many reasons; in the short, we have to change a large portion of the stack because it's such a significant change
        *   Scheduling, quota
        *   node level actually launching
        *   etc
    *   Defer solving it for now because it's so complex?
    *   Maybe at least have a roadmap for it
        *   Maybe part of sig-autoscaling
        *   Feature issue opened: https://github.com/kubernetes/features/issues/21
    *   We need a feature owner for this.
    *   Loop in @ncdc (Andy) from RedHat on the feature issue
*   OCI runtime naming - https://github.com/kubernetes/kubernetes/pull/26788#issuecomment-231144959
*   Updating max pods per core - issue number?  https://github.com/kubernetes/kubernetes/pull/28127
*   Container Runtime API status - https://github.com/kubernetes/kubernetes/issues/28789
*   Docker releases are being tested automatically as of now. We've already tested docker 1.12-rc2 and rc3, the e2e test is all green now.
*   Simplying kubelet configuration - [#27980](https://github.com/kubernetes/kubernetes/issues/27980)
*   Google main focusses this release:
    *   Disk isolation / handling of problems
    *   Runtime interface work
    *   Node/host validation
    *   Improving the configuration story?
*   https://github.com/kubernetes/kubernetes/issues/24677#issuecomment-232135730
*   Share tests between e2e and node e2e: https://github.com/kubernetes/kubernetes/pull/28807
    *   Just a shoutout to that work, no questions / problems :)
*


# July 5 {#july-5}

Cancelled today.


# June 28 {#june-28}

Cancelled today.


# June 21 {#june-21}



*   Note taker: Michael Taufen
*   10 minute overview of OpenShift System Verification and Test Repo (svt) -- Jeremy Eder
    *   https://github.com/openshift/svt
*   Container runtime interface -- Yuju Hong
*   Brainstorming on 1.4 features / roadmap
    *   https://docs.google.com/document/d/1Ei7uDriZenhFRJQAzYIwEgQm39kR0jCxTHK4qPXQFDo/edit proposed by Dawn Chen
    *   rktnetes 1.1 roadmap https://docs.google.com/document/d/1_A6NwPlQuV4t4uKe1RD9d-uhgSLImFevoS7tmg2P12M/edit?usp=sharing
    *   Work items proposed by community:
        *   Expose sysctl (Redhat)
        *   Pod checkpointing in kubelet (by CoreOS),
            *   required for self-hosting.
        *   Some other items (noted by Red Hat, but not proposed for 1.4)
            *   Be able to build NUMA
            *   Standardize GPU support, also any device that can be ref counted
            *   Additional sysctls
            *   Third party kernel modules
                *   Nvidia made their whole CUDA stack container friendly, but not every vendor will. They are hearing a lot of interest in hardware accelerators that require out of tree modules.


# June 14 {#june-14}



*   Cancelled


# June 7 {#june-7}



*   ~~10 minute overview of OpenShift System Verification and Test Repo (svt) -- Jeremy Eder~~ -- cant make it this week sorry
*   Container Runtime Interface (hopefully final discussion for the initial proposal)
*   per-pod cgroup proposal
    *   https://github.com/kubernetes/kubernetes/pull/26751
*   cAdvisor release cut status
*   rktnetes v 1.0 status
    *   Selinux regression https://github.com/kubernetes/kubernetes/pull/26936
        *   https://tests.rktnetes.io
        *   DNS - (systemd bug) https://github.com/kubernetes/kubernetes/issues/24563#issuecomment-224043285
        *
    *   hostpath/subpath error - Unknown, need to triage (https://github.com/kubernetes/kubernetes/issues/26986)
    *   Some issues running master node with rkt under kube-up right now (mostly related to how it's configured)
    *   Dns domain failure (systemd bug) https://github.com/kubernetes/kubernetes/issues/24563
    *   Influxdb isn't running failure (https://github.com/kubernetes/kubernetes/issues/26987)
    *   Docs and people kicking tires would be welcome/helpful (the first helpful for the second)
    *   Known issues docs: https://github.com/kubernetes/kubernetes/issues/26201
    *   Needed doc: "How to use rktnetes; the rktnetes user guide :)", we have one WIP, but not as a PR yet. Euan in charge of getting that in asap
    *   Nice-to-have: A doc also explaining the features / differences of rkt and why you might want to switch.
    *   Move rktnetes 1.0 to p0 (@yifan)and kubernetes 1.3 milestone. (so burndown meetings have visibility / other 1.3 visibility)
    *   sysctl proposal:
        *   https://github.com/kubernetes/kubernetes/pull/26057
*   Announcement:
    *   [NodeProblemDetector](https://github.com/kubernetes/node-problem-detector) demo
    *   1.4 node roadmap planning


# May 31 {#may-31}

Docker v1.10 is having performance issues. @timothysc https://github.com/kubernetes/kubernetes/issues/19720

Derek Carr is working on updating systemd config. Needs help with update existing CoreOS images.

Node e2e has ability to pass ginkgo flags. Tests that fail on RHEL can be blacklisted using ginkgo `skip` flags.

Container Runtime Interface Discussion - How to make progress? Vendors need to be unblocked. - https://github.com/kubernetes/kubernetes/pull/25899



*   How will new features be implemented given the pending refactoring? Dockertools package will be moved to use the new interface.
*   New runtimes will not be part of the main kubernetes repo
*   Yu-Ju will be working on re-factoring the docker runtime package to work against the new API.
*   Euan to review the runtime proposal today and provide his feedback. There hasn't been any other concerns from the community as of now


# May 24 {#may-24}



*   [node e2e test improvements ](https://docs.google.com/document/d/1Q3R-wmPDJesFz-MoLEZxlNZwDgRCenhzwETunIqPNK4/edit#)
    *   Submit queue looks bad
    *   e2e flakes are a big problem
    *   one way to tackle is to test separate components separately
    *   Makes it easier to debug, run in parallel, isolate issues -> PRs
*   **TODOs:**
*   integrating with the rest of the conformance testing suites
    *   how are node tests run when you run conformance tests
        *   against every node in the cluster? then we know every single node complies instead of just generally "we think the nodes comply"
        *   also need to run tests against previous releases, so that we run node e2e tests against everything we run the regular e2e tests against
*   isolating the node e2e job from other jenkins jobs so we don't exhaust jenkins resources
    *   moving our things to their own "projects"
*   concurrency
    *   sometimes on same instance, but we can also shard tests.
    *   as we scale this up, this will be important to keeping run time for tests low (~10 minutes ideal)
*   test harness
    *   could use some additional features, such as supporting different values for the same flag
*   Debug logs -> to files instead of just the crappy terminal dump
*   Better log file organization
*   Test coverage
    *   Moving more stuff from e2e test suite to node component test suite
    *   Start removing tests from full e2e suite as we add them to the node test suite
    *   Helps address flakiness because suites will run faster
*   rktnetes updates (https://docs.google.com/document/d/1otDQ2LSubtBUaDfdM8ZcSdWkqRHup4Hqt1VX1jSxh6A/edit# )
    *   2 pending network related PRs for fixing network regression
        *   https://github.com/kubernetes/kubernetes/pull/26096
        *   https://github.com/kubernetes/kubernetes/pull/25902
    *   Small PRs waiting to get in:
        *   Remove systemctl shell out with API calls https://github.com/kubernetes/kubernetes/pull/25214
        *   Add pod selinux support https://github.com/kubernetes/kubernetes/pull/24901
            *   Need to address paul's comments and rebase.
        *   Read-only rootfs https://github.com/kubernetes/kubernetes/pull/25719 pending to be merged
    *
    *   rktnetes benchmark https://github.com/kubernetes/kubernetes/issues/25581


# May 17 {#may-17}



*   1.3 status updates
*   rktnetes updates:
    *   shaya's cadvisor PR is getting close to merge (merged) https://github.com/google/cadvisor/pull/1284/commits
        *   tested it works for the autoscaling tests
    *   working on benchmark (yifan/shaya)
        *   Need to run benchmark e2e tests
    *   CNI/kubenet PR, @freehan will offer a review (reviewed, LGTM with nits)
        *   https://github.com/kubernetes/kubernetes/pull/25062
    *   Experiencing some regression on getting logs from journalctl
        *   https://github.com/coreos/rkt/issues/2630
    *   Conformance tests on baremetal with lkvm: (https://docs.google.com/spreadsheets/d/1S5mswBYpkT2IYAMcuzoTyOE_cdae-2vHUpnghKt6lsQ/edit#gid=0 )
        *   the journal log issue: https://github.com/coreos/rkt/issues/2630
        *   the flannel plugin doesn't have default routes: https://github.com/containernetworking/cni/issues/184
        *   Portforwarding is not implemented for lkvm stage1
    *   Following issues are all covered by LGTM PRs:
        *   Kubectl exec error: https://github.com/kubernetes/kubernetes/issues/25430
        *   Pod status error: https://github.com/kubernetes/kubernetes/issues/25242
        *   Port forward error: https://github.com/kubernetes/kubernetes/issues/25117
        *   Support per pod stage1: https://github.com/kubernetes/kubernetes/issues/23944
        *   Selinux support (not really LGTM, needs someone to take another review @pmorie?): https://github.com/kubernetes/kubernetes/issues/23773
    *   What new 1.3 features we need to be involved?
        *   Decision: not to support new 1.3 features for rktnetes 1.0
*   Kubelet bootstrap PR https://github.com/kubernetes/kubernetes/pull/25596
*   Continuation of container runtime interface discussion


# May 10 {#may-10}



*   Kurma demo from apcera.com (Ken Robertson from apcera) (15mins)
*   Talk: rkt architecture and rktnetes (~~30~~15mins) https://docs.google.com/presentation/d/1HFXemzInO4LhZ3pVJtNp0zgLQI0vXEOD4Yzigk9yisA/edit?usp=sharing
*   rktnetes v1.0 updates (10mins)
    *   Run rkt in CNI network created by kubelet https://github.com/kubernetes/kubernetes/pull/25062
    *   cAdvisor support using rkt api service: https://github.com/google/cadvisor/pull/1263
    *   Other rkt issues:
        *   https://github.com/coreos/rkt/issues/2580 (bump CNI to fix iptable leak)
        *   https://github.com/coreos/rkt/pull/2593 (return partial pod in rkt api service)
        *   https://github.com/coreos/rkt/issues/2487 (read-only rootfs)
        *   https://github.com/coreos/rkt/issues/2504 (cgroup leak)
*   docker version


# May 3 {#may-3}



*   GPU support
*   ContainerRuntime interface continue
    *   https://docs.google.com/document/d/1ietD5eavK0aTuMQTw6-21r67UU73_vqYSUIPFdA0J5Q/edit#
*   Eviction proposal ?  any more feedback?
*   Docker 1.10 and Docker 1.11 discussion
    *   Docker 1.11 validation issue: https://github.com/kubernetes/kubernetes/issues/23397
    *   Action Item (timstclair): email k8s-dev about docker 1.11 preferences
*   Reduce kubelet LOC via moves https://github.com/kubernetes/kubernetes/pull/25028
*   rktnetes:
    *   let rkt run in the network namespace created by kubelet/network plugin:
        *   https://github.com/kubernetes/kubernetes/pull/25062
        *   This could solve a bunch of remaining failing e2e tests (most of them are looking for pod IP in downward API and /etc/hosts )
    *   update the cloud config and kubernetes compoments yamls, has some issues
        *   https://github.com/kubernetes/kubernetes/pull/22663
    *   fixing rkt gc problem (bug in rkt and cni)
        *   https://github.com/coreos/rkt/issues/2565
    *   working on letting user to specify different stage1 image in pod annotation
        *   PR not ready yet


# April 26 {#april-26}



*   New ContainerRuntime interface discussion
    *   Difficult integrate at Pod level declarative interface.
    *   Expect imperative container level interface, and OCI compatible in a long run
    *   Proposed 2 options:
        *   1) introduce OCI compatible interfaces now
        *   2) introduce docker like container level interface first
    *   AI: yuju write a high level design doc and continue discussing next week
*   Node on Windows - initial investigation / next steps
    *   https://docs.google.com/presentation/d/16nYb13oulBoB4d6QZXYm-9sk3TN898qumzY5ur4wirU/edit?usp=sharing
    *   https://docs.google.com/document/d/1qhbxqkKBF8ycbXQgXlwMJs7QBReiSxp_PdsNNNUPRHs/edit#
*   Followup with custom metrics discussion
    *   AI: agoldste@ from Redhat is going to write a high-level requirement doc to share with us. A separate VC to continue the discussion.
*   rktnetes status updates:
    *   e2e testing: http://rktnetes.io
    *   work on stableness issue
    *   support stage1 fly
    *   work on fixing the race between cgroup discovery and pod discovery by rkt api service
    *   rkt usage monitoring:
        *   https://github.com/dgonyeo/rkt/blob/ddd59b8935a3a9ee7c031e753642ea7863526fde/tests/rkt-monitor/README.md
        *   Need to put more example results here
*   NVIDIA GPU support:
    *   Kubelet:

        The kubelet should have "--device" options, and volumes for GPU, NVIDIA GPU is part of it, and also have the interface to support other GPUs.


        Cadvisor include the NVML libs to find NVIDIA GPUs on host, then kubelet could send GPU information to kube-scheduler.

*   Kube-scheduler:
    *   Include GPU information:

        Number: // how many GPUs needed by the container.

        Vendor: // So far, only NVIDIA GPU.


        Library version: // run the container on the right host, not just a host with GPUs.


# April 19 {#april-19}



*   InitContainer proposal discussion: ([#23666](https://github.com/kubernetes/kubernetes/pull/23666))
*   Demo: Out-of-resource eviction?
*   rktnetes status updates:
    *   rkt running on master (demo)
    *   rktnetes jenkins e2e testing infra
    *   milestones: https://github.com/kubernetes/kubernetes/milestones/rktnetes-v1.0
    *   e2e failures: 14-15, with 8ish we know that we are not supporting, or the tests are too specific. e.g. DNS tests
    *   Met some stableness issue with kubelet running rkt, need to be fixed
    *   Per pod overhead measurement (systemd, journald)
        *   https://github.com/coreos/rkt/pull/2324
*   cAdvisor roadmap updates:
    *   punt standalone cAdvisor
    *   cAdvisor validation and testing
*   announcement: [minikube](https://github.com/kubernetes/minikube) on local host


# April 12 {#april-12}



*   Experimental NVIDIA GPU support ([#24071](https://github.com/kubernetes/kubernetes/pull/24071), [#17035](https://github.com/kubernetes/kubernetes/issues/17035), [#19049](https://github.com/kubernetes/kubernetes/issues/19049), [#21446](https://github.com/kubernetes/kubernetes/pull/21446)): add support in kubelet to map host devices into containers, add new `nvidiaGpu` resource (@therc)
    *   More constrained initial implementation than in #24071
    *   Hardcode kubelet to report one NVIDIA GPU when a special experimental flag is enabled
    *   When a pod requires a GPU, expose hardcoded list of host devices into container
    *   Do not use any code from nvidia-docker and its Docker plugin yet. This means that images will have to bundle required NVIDIA libraries
    *   For next steps, consider an approach halfway between having NVIDIA libraries inside every Docker image vs. having them exposed by the nvidia-docker volume driver: have administrators set up a directory on the host with NVIDIA binaries and libraries
    *   @therc/@Hui-Zhi will send PRs and look into rkt support as well (more feasible if we don't use the Docker plugin framework)
*   rktnetes status updates:
    *   fixed service ip reachable issue, rkt locking issue.
    *   working on master node using rkt.
    *   milestone: https://github.com/kubernetes/kubernetes/issues?q=is%3Aopen+is%3Aissue+milestone%3Arktnetes-v1.0
    *   things might not be supported:
        *   hostipc/hostpid mode https://github.com/systemd/systemd/pull/2982#issuecomment-208757763
        *   kubectl attach https://github.com/kubernetes/kubernetes/issues/23889
        *   read-only rootfs https://github.com/kubernetes/kubernetes/issues/23837
*   Self-hosted kubelet: https://github.com/kubernetes/kubernetes/pull/23343

April 4



*   e2e node tests are failing
*   cAdvisor with rkt/kvm-stage1? https://github.com/coreos/rkt/issues/2341#issuecomment-205799519
*   rkt container level interface https://github.com/coreos/rkt/issues/2375
*   rktnetes status: https://docs.google.com/document/d/1otDQ2LSubtBUaDfdM8ZcSdWkqRHup4Hqt1VX1jSxh6A/edit?usp=sharing
*   Self-hosted kubelet proposal: https://github.com/kubernetes/kubernetes/pull/23343
*

Mar 29



*   Kubelet issue with termination - issue with watch: https://github.com/openshift/origin/issues/8176
*   Yifan looking at issues with rkt pod startup latency
*   State of cAdvisor refactoring
    *   Tim St. Clair mentioned that no changes inside kubelet in the near term.
*   Systemd driver in Kubelet - Does it need first class support?
    *   Probably not. We need more information to discuss further
*   Kubelet bootstrapping
    *   GKE will not switch, so there is no immediate changes required for other providers.
*   Kubelet CI
    *   Kubelet Node e2e is now a merge blocker
    *   Hackathon next week - will be open to the entire SIG. Goal is to increase test coverage
    *   How do we add a test to the Conformance suite?
        *   It needs to include only the core kubelet features. But the Kubelet API is expected to be portable.
        *   Some tests have failed on Systemd in the past like Filesystem accounting
        *   Some of the kubelet API is distro dependent.
        *   Why/When write a node e2e test?
            *   Any node level feature that can be tested on a single host.
            *   Simpler test infrastructure. Easier to test node conformance
            *   Smaller scope
        *   We need multiple deployment configurations for the e2e suite and have the tests be able to discover these configurations.
*   Increase maximum pods per node for kube-1.3 release #23349
    *   https://github.com/kubernetes/kubernetes/issues/23349
    *   Benchmarking is the first step forward

Mar 22



*   cAdvisor 1.3+ work planning by stclair
    *   https://docs.google.com/document/d/1aLJ7OBVRO2QKnf8xO0Mbkd_Wyca8J0sr9q12ZObmkO0/edit
*   rktnetes:
    *   https://github.com/kubernetes/kubernetes/issues/23335
    *   cadvisor: https://github.com/google/cadvisor/pull/1154
    *   Milestone: 85% e2e tests passed.
*   Increase max-pods https://github.com/kubernetes/kubernetes/issues/23349
*

Mar 15



*   1.3 planning (proposed by Dawn)
    *   https://docs.google.com/a/google.com/document/d/12dyUH3HWUMjetWnslnPbY89fyN5I8vLjl6Gi9F2HqOk/edit?usp=sharing
    *   Not final yet. Waiting for feedback from community, partners, developers and PMs
*   rktnetes:
    *   post-start/pre-stop hooks problem with slow rkt startup
    *   e2e status: experienced regression on rkt
    *   kube-proxy issue with kubenet. (doesn't seem to be specific to rktnetes)
        *   https://github.com/kubernetes/kubernetes/issues/20475#issuecomment-195096662
*   Developer docs: https://github.com/kubernetes/kubernetes/issues/23033

Mar 8



*   rktnetes status:
    *   Working on post-start/pre-stop hooks
    *   Fixing ENTRYPOINT/CMD issue
    *   cAdvisor support WIP
    *   38 failures out of 178 tests, categorized https://docs.google.com/document/d/1otDQ2LSubtBUaDfdM8ZcSdWkqRHup4Hqt1VX1jSxh6A/edit#
*   e2e flake: Downward API should provide pod IP: https://github.com/kubernetes/kubernetes/issues/20403

Mar 1

[Node e2e tests](https://github.com/kubernetes/kubernetes/blob/master/docs/devel/e2e-node-tests.md)



*   Run against PRs using the trigger phrase "`@k8s-bot test node e2e experimental`"
*   Run locally using "make test_e2e_node"
*   Tests can be built to tar.gz and copied to arbitrary host
*   Would like to distribute testing of distros and setup dashboard to publish results

rktnetes status:



*   https://docs.google.com/document/d/1otDQ2LSubtBUaDfdM8ZcSdWkqRHup4Hqt1VX1jSxh6A/edit?usp=sharing

Feb. 23

Docker v1.10 integration status

- Lantao has fixed all the issues that has been identified. Docker go client had an issue that has been fixed. Docker startup scripts have been updated. A PR fixing -d option is here https://github.com/kubernetes/kubernetes/pull/20281

- Prior to upgrade we'll need a daemonset pod to migrate the images.

- Vishh is working on disabling seccomp profiles by default.

Yifan and Shaya status updates:

CNI plugin support for rkt



*   kubenet support for rkt https://github.com/kubernetes/kubernetes/pull/21047
*   cadvisor integration
    *   Basic cgroups stats are available. Patch is out for review.
    *   Additional metrics will be added soon.

TODO: vishh to post documentation on patch policy for releases.

v1.2 will support 100 pods by default.

Docker v1.9 validation is having issues. We are performing a controlled restart of the docker daemon to try and mitigate the issue.

Feb. 16



*   Projects and priorities brainstorm
    *   Refactor kubelet: Too complex for new developer, we should refactor the code. Better separation and architecture is needed.
        *   Dawn: One important thing is to cleanup container runtime and image management interface. Maybe separate pod level and runtime level api.
        *   Tim works on cleanup cadvisor kubelet interface.
        *   Should have a sig-meeting soon.
        *   @Dawn will file an issue about this.
    *   Better Disk management (disk resource quota mgmt)
        *   Openshift issue https://github.com/openshift/origin/pull/7206
    *   Performance requirement for 1.3
        *   Control Plane Networks, network segregation: https://github.com/kubernetes/kubernetes/issues/21331
        *   Share related validation test in openshift.
    *   Kubelet machine health API
        *   Kubelet should provide api to receive categorized machine problem from "Machine Doctors" such as machine monitor, kernel monitor etc.
        *   Some existing systems such as Ganglia https://github.com/kubernetes/kubernetes/issues/21333
        *   Who should take actions: Kubelet? Application? Operator?
        *   Use DaemonSet to handle it and mark kubelet as NotReady?
    *   Determine if we should support cpuset-cpus and cpuset-mem: https://github.com/kubernetes/kubernetes/issues/10570
    *   Arbitrary tar file?

Feb. 10



*   schedule change?
    *   going to try Tuesday 11-12 PST
*   disk monitoring for rkt
    *   two proposals: all in-kubelet, in-cadvisor
    *   https://github.com/kubernetes/kubernetes/pull/20887

Feb. 3



*   docker 1.10 validation updates
    *   https://github.com/kubernetes/kubernetes/issues/19720
    *   A go-dockerclient bug: https://github.com/fsouza/go-dockerclient/issues/455
*   updates on sig-node (Asia):
    *   Meeting notes: https://docs.google.com/document/d/1L8s6Nyu5hNJxCOZLqJsuVEFAScWKhbwcis0X-j-upDc/edit?usp=sharing
    *   Topics:
        *   Docker container runtime conformance test.
        *   Rktnetes outstanding issues.
        *   Hyper integration.
*   image management plans
    *   image manager to support multi runtimes: https://github.com/kubernetes/kubernetes/issues/19519
    *   vish to create follow up doc for discussion
*   kernel issues (#20096)
    *   Conformance tests ? (https://github.com/kubernetes/contrib/issues/410)
*   sig-scale concern issue [20011](https://github.com/kubernetes/kubernetes/issues/20011) about kubelet changes potentially breaking kubemark. \


Jan. 27



*   What tests exists for validating kubelet performance across docker releases?
    *   Docker Mirco Benchmark: https://docs.google.com/document/d/1aKfsxRAmOtHkFf4Wcn0qoFJRU6qOzwATxOVKDVKQucY/edit
    *   Docker Event Stream Benchmark: https://docs.google.com/document/d/1FUBwdWEqdUR7h-dAwDvOWg9RrjkerN48QmMm7yCylJY/edit
    *   Code: https://github.com/Random-Liu/docker_micro_benchmark
    *
*   rkt?https://docs.google.com/document/d/132fMB60poMCcejf8l82Vxpynff8_BRYcZeOn24jn8pc/edit

Jan. 20



*   OCI meeting updates
*   rkt / CoreOS integration status updates
*   1.2 features status recap

Jan. 13



*   Node scalability issue recap. Discussed via issues and PRs
*   systemd spec https://github.com/kubernetes/kubernetes/pull/17688 ready for implementation. Minor discussion on the cgroup library problems in libcontainer.
*   OCI meeting is going on, will have more updates in next sig-node sync

Jan. 6



*   Node scalability and performance (redhat): https://github.com/kubernetes/kubernetes/issues/19341
*   1.2 scalability goals: https://github.com/kubernetes/kubernetes/issues/16943
*   Node performance testing guide: https://github.com/kubernetes/kubernetes/pull/18779
*   RH goal for 1.2: 160 pods per node
