# Kubernetes Retrospective: 1.3 Storage Post Code Freeze Submission
*Last updated: September 20, 2016*

**Owners:** Brad Childs ([@childsb](https://github.com/childsb)), Michael Rubin ([@matchstick](https://github.com/matchstick))

**Collaborators:** Saad Ali ([@saad-ali](https://github.com/saad-ali)), Paul Morie ([@pmorie](https://github.com/pmorie)), Tim Hockins ([@thockin](https://github.com/thockin)), Steve Watt ([@wattsteve](https://github.com/wattsteve))

**Links:**
* [1.3 Schedule Dates](https://git.k8s.io/sig-release/releases/release-1.3/release-1.3.md)

## Purpose
This document is intended to chronicle the decisions made by the [Storage SIG](/sig-storage/README.md) near the end of the Kubernetes 1.3 release with the storage stack that were not well understood by the wider community. This document should explain those decisions, why the SIG made the exception, detail the impact, and offer lessons learned for the future.

## What Problem Were We Trying to Solve?
Kubernetes 1.2 had numerous problems and issues with the storage framework that arose from organic growth of the architecture as it tackled numerous new features it was not initially designed for. There were race conditions, maintenance and stability issues, and architectural problems with all major components of the storage stack including the Persistent Volume (PV) & Persistent Volume Claim (PVC) controller and the attach/detach and mount/unmount logic.

The PV/PVC controller handles the connection of provisioned storage volumes to a user claim for storage. The attach/detach logic handles how volumes are attached to hosts. The mount/unmount logic handles how volumes are mounted into containers. Architecturally in 1.2 the attach/detach logic was part of the kubelet on the node.

A characteristic list of issues (as not all of them were well captured in GitHub issues) include:

1. Approximately a 5% rate of incidents under controlled conditions where operations related to Claims binding to Persistent Volumes would fail.
2. Rapid creation and deletion of pods referencing the same volume could result in attach/detach events being triggered out of order resulting in detaching of volumes in use (resulting in data loss/corruption). The current 1.2 work around was to fail the operation. This led to surprises and failures in launching pods that referenced the same volume.
3. Item #2 created instability in use of multiple pods referencing the same Volume (a supported feature) even when only one pod uses it at a time ([#19953](https://github.com/kubernetes/kubernetes/issues/19953))
4. Hiccups in the operation flow of binding the Claims to Volumes resulted in timeouts of tens of minutes.
5. External object bleeding. Much of the logic was centered on a state machine that lived in the kubelet. Other kube components had to be aware of the state machine and other aspects of the binding framework to use Volumes.
6. Maintenance was difficult as this work was implemented in three different controllers that spread the logic for provisioning, binding, and recycling Volumes.
7. Kubelet failures on the Node could “strand” storage. Requiring users to manually unmount storage.
8. A pod's long running detach routine could impact other pods as the operations run synchronously in the kubelet sync loop.
9. Nodes required elevated privileges to be able to trigger attach/detach. Ideally attach/detach should be triggered from master which is considered more secure (see Issue [#12399](https://github.com/kubernetes/kubernetes/issues/12399)).

Below are the Github Issues that were filed for this area:
* Problem rescheduling POD with GCE PD disk attached ([#14642](https://github.com/kubernetes/kubernetes/issues/14642))
* GCE PD Volumes already attached to a node fail with "Error 400: The disk resource is already being used by" node ([#19953](https://github.com/kubernetes/kubernetes/issues/19953))
* Kubelet should be able to delete 10 pods per node in 1m0s ([#23591](https://github.com/kubernetes/kubernetes/issues/23591))
* Detach EBS volumes when node restarted ([#26847](https://github.com/kubernetes/kubernetes/issues/26847))
* Technical debt: refactor Kubelet.HandlePodCleanups into separate thread ([#19645](https://github.com/kubernetes/kubernetes/issues/19645))
* EBS volume mount failures due to "... already attached to an instance" are not retried ([#18785](https://github.com/kubernetes/kubernetes/issues/18785))
* Node upgrades: e2e test: ensure persistent volumes survive when pods die ([#6084](https://github.com/kubernetes/kubernetes/issues/6084))
* Consider "attach controller" to secure cloud provider credentials ([#12399](https://github.com/kubernetes/kubernetes/issues/12399))

## How Did We Solve the Problem?
Addressing these issues was the main deliverable for storage in 1.3. This required an in depth rewrite of several components.

Early in the 1.3 development cycle (March 28 to April 1, 2016) several community members in the Storage SIG met at a week long face-to-face summit at Google's office in Mountain View to address these issues. A plan was established to approach the attach/detach/mount/unmount issues as a deliberate effort with contributors already handling the design. Since that work was already in flight and a plan established, the majority of the summit was devoted to resolving the PV/PVC controller issues. Meeting notes were captured [in this document](/sig-storage/1.3-retrospective/2016-03-28_Storage-SIG-F2F_Notes.pdf).

Three projects were planned to fix the issues outlined above:
* PV/PVC Controller Redesign (a.k.a. Provisioner/Binder/Recycler controller)
* Attach/Detach Controller
* Kubelet Volume Redesign

At the end of the design summit, the attendees of the summit agreed to pseudo code for a re-written PV/PVC controller and a go-forward plan for the attach/detach controller and kubelet volume redesign.

Resources were established for the PV/PVC controller rework at the conclusion of the design summit and the existing resources on the attach/detach/mount/unmount work deemed acceptable to complete the other two projects.

At this point, a group of engineers were assigned to work on the three efforts that compromised the overhaul. The plan was to not only include development work but comprehensive testing with time to have the functionality “soak” weeks before 1.3 shipped. These engineers were composed of a hybrid team of Red Hat and Google. The allocation of work made making all three sub deliverables in 1.3 aggressive but reasonable.

Near the end of 1.3 development, on May 13, 2016, approximately one week prior to code freeze, a key engineer for this effort left the project. This disrupted the Kubelet Volume Redesign effort. The PV/PVC controller was complete (PR [#24331](https://github.com/kubernetes/kubernetes/pull/24331)) and committed at this point. However the Attach/Detach Controller was dependent on the Kubelet Volume Redesign and was impacted.

The leads involved with the projects met and the Kubelet Volume Redesign work was handed off from one engineer to another familiar with Storage. The decision to continue this work after the 1.3 code freeze date of May 20 was based on the need to address the outstanding issues in 1.2. Also much of the Attach/Detach Controller work had been committed but was dependent on the Kubelet Volume Redesign effort.

The Kubelet Volume Redesign involved changing fundamental assumptions of data flow and volume operations in kubelet. The high level change introduced a new volume manager in kubelet that handled mount/unmount logic and enabled attach/detach logic to be offloaded to the master (by default, while retaining the ability for kubelet to do attach/detach on its own). The remaining work to complete the effort was the kubelet volume redesign PR ([#26801](https://github.com/kubernetes/kubernetes/pull/26801)). This combined with the attach/detach controller (PR [#25457](https://github.com/kubernetes/kubernetes/pull/25457)) were substantial changes to the stack.

## Impact:

1. **Release delay**
  * The large amount of churn so late in the release with little stabilization time resulted in the delay of the release by one week: The Kubernetes 1.3 release [was targeted](https://git.k8s.io/sig-release/releases/release-1.3/release-1.3.md) for June 20 to June 24, 2016. It ended up [going out on July 1, 2016](https://github.com/kubernetes/kubernetes/releases/tag/v1.3.0). This was mostly due to the time to resolve a data corruption issue on ungracefully terminated pods caused by detaching of mounted volumes ([#27691](https://github.com/kubernetes/kubernetes/issues/27691)). A large number of the bugs introduced in the release were fixed in the 1.3.4 release which [was cut on August 1, 2016](https://github.com/kubernetes/kubernetes/releases/tag/v1.3.4).
2. **Instability in 1.3's Storage stack**
  * The Kubelet volume redesign shipped in 1.3.0 with several bugs. These were mostly due to unexpected interactions between the new functionality and other Kubernetes components. For example, secrets were handled serially not in parallel, namespace dependencies were not well understood, etc. Most of these issues were quickly identified and addressed but waited for 1.3 patch releases.
  * Issues related to this include:
     * PVC Volume will not detach if PVC or PV is deleted before pod ([#29051](https://github.com/kubernetes/kubernetes/issues/29051))
     * GCE PD Detach fails if node no longer exists ([#29358](https://github.com/kubernetes/kubernetes/issues/29358))
     * Batch creation of pods that all reference the same secret volume takes a long time ([#28616](https://github.com/kubernetes/kubernetes/issues/28616))
     * Error while tearing down pod, "device or resource busy" on service account secret ([#28750](https://github.com/kubernetes/kubernetes/issues/28750))

## Lessons Learned
The value of the feature freeze date is to ensure the release has time to stabilize. Refactoring or features that need to be merged past feature freeze date as an exception should be a tool that can be used, albeit sparingly, for the sake of a release. Exceptions should meet certain requirements which the Kubelet Volume Redesign did not meet.
### Things that went well:
1. The majority of the development went smoothly with adequate testing and little or no instability after release for two of the three sub areas.
2. Once the wider impact of the Kubelet Volume Redesign work was understood by the Storage SIG it was widely communicated in the burn down meetings.
3. The Storage SIG communication was very tight with hour by hour coordination in this area from beginning to end.
4. Many area experts were able to jump in and help debug.
5. The community trusted the Storage SIG to make the right decisions.

### Things that went wrong:
1. The scope of the Kubelet Volume Redesign work and its indirect impact on the system, while understood to be large, was still under-estimated: while the code was written and code reviewed reasonably within schedule, considering the size/extent of the changes, the amount of time required for soaking, stabilization, and testing was insufficient.
2. In the decision to move forward with coding beyond code freeze, not enough thought was invested in what could go wrong or how to mitigate that.
3. The Storage SIG decisions were not advertised widely enough outside of the SIG early on.

## Action Items
1. Develop deeper testing across the entire storage stack to allow large changes to be made with confidence.
  * Status: [Planned for 1.5](https://docs.google.com/document/d/1-u1UA8mBiPZiyYUi7U7Up_e-afVegKmuhmc7fpVQ9hc/edit?ts=57bcd3d4&pli=1)
    * Discussed at [Storage-SIG F2F meeting held August 10, 2016](https://docs.google.com/document/d/1qVL7UE7TtZ_D3P4F7BeRK4mDOvYskUjlULXmRJ4z-oE/edit). See [notes](https://docs.google.com/document/d/1vA5ul3Wy4GD98x3GZfRYEElfV4OE8dBblSK4rnmrE_M/edit#heading=h.amd7ks7tpscg).
2. Establish a formal exception process for merging large changes after feature complete dates.
  * Status: [Drafted as of 1.4](https://git.k8s.io/features/EXCEPTIONS.md)

Kubernetes is an incredibly fast moving project, with hundreds of active contributors creating a solution that thousands of organization rely on. Stability, trust, and openness are paramount in both the product and the community around Kubernetes. We undertook this retrospective effort to learn from the 1.3 release's shipping delay. These action items and other work in the upcoming releases are part of our commitment to continually improve our project, our community, and our ability to deliver production-grade infrastructure platform software.
