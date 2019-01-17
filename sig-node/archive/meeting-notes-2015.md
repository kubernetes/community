# sig-node weekly meeting

Dec. 9



*   Meeting note with huawei@:

    https://docs.google.com/document/d/1H2FybZUh0qS2jlOGeVE85LOS_7VIxGKiz2piDKlXBro/edit?usp=sharing

*   ImageSpec
    *   https://github.com/kubernetes/kubernetes/pull/18308#issuecomment-162729102
    *   not required for rkt 1.0, but need a path forward
    *   Dawn will file separate issue for follow up discussion on kubelet moving to full management of images. Unblock the current development progress.
*   Logging management discussion (Vishnu)
    *   https://github.com/kubernetes/kubernetes/issues/17183
    *   use cases we need to support:
        *   on-demand log rotation as pods start to come up against their disk usage limits
        *   shipping to something like fluentd
    *   either delegated to runtime or controlled by kubelet
    *   On disk space management, image accounting should be separated from runtime logging.
*   Image accounting:
    *   Still to define what exactly this means in rkt: https://github.com/coreos/rkt/issues/1814
    *   @vishh to provide feedback
    *   what about arbitrary image types like tarballs?
        *   jon: container runtime should provide image management
*   cAdvisor integration with rkt updates (Dawn)
    *   blocked by broken scripts due to the integration of CoreOS and refactory
    *   Will pick up yifan's pr and retry.
    *   Yifan will bring up a cluster with rkt
*   rkt roadmap updates (Jonathan) https://github.com/coreos/rkt/blob/master/ROADMAP.md
    *   kvm state?
        *   should have feature parity except for networking and testing
    *   rkt stage1 overhead WIP https://github.com/dgonyeo/rkt-monitor

Dec. 2



*   Logging management discussion (Vishnu)

           No time yet. Move to next week.



*   rkt roadmap updates (Jonathan) https://github.com/coreos/rkt/blob/master/ROADMAP.md
*   benchmark stage1 resource usage https://github.com/coreos/rkt/issues/1788
*   Sumsung's use cases and requirements for node and container runtime (Bob)

           kubelet logging isssue https://github.com/coreos/bugs/issues/990

https://github.com/kubernetes/kubernetes/issues/14216

rkt prs:

    https://github.com/kubernetes/kubernetes/pull/17968

    https://github.com/kubernetes/kubernetes/pull/17969

    rkt fly

https://github.com/coreos/rkt/pull/1825

Disk Accounting Proposal - https://github.com/kubernetes/kubernetes/pull/16889

Nov. 18



*   kubelet + systemd?

- design of cgroup hierarchy ideal (Vishnu)

- Nalin to send out volumec branch and design doc

- Discussion on logging (Vishnu to paste link to existing issue)


    https://github.com/kubernetes/kubernetes/issues/17183

  - rkt team to measure overhead of running the journal in container

- Discussion on PID1 being systemd in the infrastructure container or not and different strategies; sounds like measuring needs to happen here https://github.com/coreos/rkt/issues/1788

Just to be super clear it seems like we have three divergent paths for running a docker container that are likely to exist in the kubelet:

1) Docker engine mode

2) rkt mode

3) runc+systemd mode



*   Yifan side:
    *   kube-up for coreos/gce/docker, PRs out there (https://github.com/kubernetes/kubernetes/pull/17240

    https://github.com/kubernetes/kubernetes/pull/17241


    https://github.com/kubernetes/kubernetes/pull/17243 )

*   working on refactoring the rkt/kubelet with rkt api service.

Nov. 11



*   Cancelled as people are in the kube-con

Nov. 4



*   From google: 1.2 releases focuses on testing
*   Node conformance test, e.g. including testing the kernel version, kernel config, docker runtime, rkt runtime, systemd version, etc. Such testing makes the node is validate to join the cluster.
*   coreos side: yifan working on converting the gce master to use coreos image. also need to work together with node team on the node conformance tests for rkt.
    *   question on setting up the master node : what's the best practice to maintain those pod templates (e.g. kube-apiserver, logging, dns addons) https://github.com/kubernetes/kubernetes/pull/16760#discussion_r43826882
    *   dawn chen said we can run a script that's in the saltstack dir on the node, which evaluate the pod templates.
*   vish asked about the status for the rkt after our last discussion about the pod lifecycle
    *   not much progress, mentioned jonboulle's rkt fly pr https://github.com/coreos/rkt/pull/1416  (what's the status for that? @jonboulle)
*   dawn chen asked about what's rkt's currently developing direction? https://github.com/coreos/rkt/blob/master/ROADMAP.md#rkt-roadmap
*   Also she would like to know our OKRs in the next quarter. (rkt, etcd, tectonic?) Can we share that with her? @jonboulle
*   mics: vish talked briefly about the latest updates in OCI, splitting starting of a container to create and start. yifan mentioned we have recently put efforts on acbuild (an appc image build tool)  https://github.com/appc/acbuild

Sept. 30

Date: Sept 30, 2015

Attendees: YiFang, Jonnathan from CoreOS,

                   stclair@, yjhong@, vishnuk@, dawnchen@ from Google

Agenda



*   oci status update:

    (jonboulle): OCI being blocked, waiting on a lot of issues, very slow progress.



*   appc + OCI harmonization effort: https://groups.google.com/a/opencontainers.org/forum/#!topic/dev/uo11avcWlQQ
*   defining an image format in OCI: https://groups.google.com/a/opencontainers.org/forum/#!topic/dev/OqnUp4jOacs
*   https://groups.google.com/a/opencontainers.org/forum/#!topic/dev/1T0z1IJWxw8
*   appc status:
    *   work stalled on OCI announcement, starting to pick up again
    *   TODO(jonboulle): categorise appc issues and follow up with summary
    *   example: pod lifecycle https://github.com/appc/spec/pull/500
        *   background discussion: https://github.com/appc/spec/issues/276
*   kubernetes restart policy
    *   applied per-pod
    *   _jonboulle: a container "restart" is underspecified (e.g., is filesystem persisted?). Currently, it is defined by implementation in Docker. In rkt, we currently just restart the entire pod, but are considering implementing more granular behaviour_
        *   https://github.com/kubernetes/kubernetes/blob/ba89c98fc7e892e816751a95ae0ee22f4266ffa5/docs/user-guide/pod-states.md#restartpolicy
*   rkt introduction
    *   background - https://coreos.com/blog/rocket/
*   Dawn: concerned about having systemd manage cgroup; unified hierarchy does not work for Google
    *   One issue for memory management policy: https://github.com/kubernetes/kubernetes/issues/14532
    *   We might have different policy enforced by cgroup hierarchy for cpu, blockio, etc.
    *   _Should isolator adjustment be part of container runtime API?_
        *   related appc issue: https://github.com/appc/spec/issues/54
        *   rkt implementation-wise, this can be achieved today at both pod-level and app-level using systemd APIs ([SetUnitProperties](http://www.freedesktop.org/wiki/Software/systemd/dbus/) + [resource settings](http://www.freedesktop.org/software/systemd/man/systemd.cgroup.html#Options))
*   Node architecture discussion
    *   today there are two runtime integration points:
        *   kubelet, for creating/running/etc pods
        *   cAdvisor, for exposing stats
    *   Vish talked about proposal for cAdvisor to take over all responsibilities, so that container runtimes only need to integrate in one place
        *   simpler maintainability (one codebase)
        *   OR, could integrate cAdvisor responsibilities into kubelet
    *   e.g. moving [Container Runtime interface](https://github.com/kubernetes/kubernetes/blob/bd8ee1c4c49c724e80a8f8d59e732ea7855eba8e/pkg/kubelet/container/runtime.go#L53) into cAdvisor
    *   Container Runtime Interface is changing from declarative to imperative
    *   client/server? https://github.com/kubernetes/kubernetes/issues/13768
*   when updating pod spec:
    *   only thing that can be done today is changing the image version
    *   updating resources: https://github.com/kubernetes/kubernetes/issues/5774
    *   dawn: add rant on https://github.com/appc/spec/issues/276
    *   …
    *   perhaps kubernetes pod really conflates two levels of abstraction: _scheduling _and _resource enforcement/management_
*   rkt e2e
    *   yifan@ is working on this whenever he has time.

What is a Pod:



*   Shared namespaces excepting mount
*   Restart policies will be at pod level
*   Per-container restart policies are required for certain volumes git pull?
*   Life cycle hooks at the container level are needed. Pre-start hooks at the container level.
*   Privileged pods require access to host

Why pod updates?



*   Pod updates are needed for updating image names - misspelled image names and in-place image updates.
*   In-place updates are very useful when containers cannot tolerate restarts. For example, applications that load a lot of data from volumes before functioning.
*   Auto-scaling requires updates to pods.
*   Adding/removing containers - in-place updates using a hot-swap mechanism - start a new container and remove the old one. Adding/Updating side-cars (logging, monitoring, etc)
*   Updating volumes to pods are also a very useful feature for users.
*   logging: yifan to check fluentd, how can it integrate with rkt/journal
    *   Getting logs after pods exit: https://github.com/coreos/rkt/issues/1528
