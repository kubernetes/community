# Code Organization and Release Process Improvement

**Identify the following before beginning your session. Do note move forward until these are decided / assigned: **

- **Session Topic**: Code Organization and Release Process Improvement
- **Topic Facilitator(s)**: bgrant0607@
- **Note-taker(s) (Collaborating on this doc)**: jdumars@
- **Person responsible for converting to Markdown &amp; uploading to Github:** michellen@

### Session Notes

#### Background from November dev summit:
- [Slides](https://docs.google.com/presentation/d/1SD6a6eJl47t0qyTFE8GzaiytW4T_crdWgYAMCaLy1W8/edit#slide=id.g18ae10430d_0_7)
- [Notes](https://www.google.com/url?q=https://docs.google.com/document/d/1zN2DWKerXwbzxZTO52wBRqp_uHMdLp8P52xYOmp5WZ4/edit&sa=D&ust=1496186064062000&usg=AFQjCNG8Z-9KBJJkUsQY9F38CdZY3Nc_LA)

#### Github orgs:
Github supports 2 levels of hierarchy, orgs and repos, and we need to use both effectively. We need to move away from monorepos, such as kubernetes/kubernetes and kubernetes/contrib, and mono-orgs, such as kubernetes and kubernetes-incubator. [Kubernetes-client](https://github.com/kubernetes-client) is the first focused org.
 
Ideally, code would be divided along lines of ownership, by SIG and subproject, with infrastructure for packaging the code (APIs and controllers) into the desired number of components (e.g., hyperkube as well as factored by [layer](https://docs.google.com/presentation/d/1oPZ4rznkBe86O4rPwD2CWgqgMuaSXguIBHIE7Y0TKVc/edit#slide=id.p)) and releases (e.g., [monthly](https://github.com/kubernetes/community/issues/567) as well as LTS).
 

#### Where we are on multiple repos:
* [Code organization issue](https://github.com/kubernetes/kubernetes/issues/4851)
* [kubernetes multi-repo issue](https://github.com/kubernetes/kubernetes/issues/24343) and [contrib](https://github.com/kubernetes/contrib/issues/762) breakup issue
* User docs were moved to kubernetes/kubernetes.github.io
* Contributor docs were moved to kubernetes/community
* Examples have been copied to kubernetes/examples, but haven’t yet been removed from the kubernetes repo
* [API machinery](https://github.com/kubernetes/kubernetes/issues/2742): In order to build virtually anything outside the kubernetes repo, the API machinery needs to be made available, for component configuration, for building APIs, for building controllers and other Go clients of the APIs, etc.
  * Staging
  * [Client-go](https://github.com/kubernetes/kubernetes/issues/41629)
  * [API types](https://github.com/kubernetes/kubernetes/pull/44784), will be done during 1.7 code freeze
* TODO: util
  * pkg/util moves thread
* Kubectl:
  * In the process of breaking kubectl <-> kubernetes dependencies
  * Have kubernetes/kubectl repo, need to migrate issues there
    * Issue migration tool exists - creates lots of notifications

Next up:
* [kubeadm](https://github.com/kubernetes/kubernetes/issues/35314)
* Federation
* [cloudprovider](https://git.k8s.io/kubernetes/pkg/cloudprovider/README.md) and [cluster](https://git.k8s.io/kubernetes/cluster/README.md)
* Scheduler
* [Kubelet](https://github.com/kubernetes/kubernetes/issues/444)

#### Build system and checking in generated code (or not):
* [Issue about not checking in generated code](https://github.com/kubernetes/kubernetes/issues/8830)
* [Nuke Ugorji](https://github.com/kubernetes/kubernetes/issues/36120)
* [Update and verify script thread](https://groups.google.com/d/msg/kubernetes-dev/5rVmSJqCq-U/tyN8OVRjBgAJ)
TODO: Bazel, make

#### Vendoring:
What would be better than godeps? Dep? Glide?
[https://github.com/kubernetes/kubernetes/issues/44873](https://github.com/kubernetes/kubernetes/issues/44873)
 
#### Other:
* Client and [client tool](https://github.com/kubernetes/release/issues/3) releases
* More unit tests rather than end-to-end tests
* [Feature branches](https://github.com/kubernetes/community/issues/566)

#### Multi-repo requirements

#### Release process
* Components are combined from subrepos to build released bits
* Components each have sufficient testing against their repo to determine if they are stable to release

#### Development process in multiple repos
* PRs sent to external repos
* PR and CI tests run against external repo.  Sufficient to validate the code maybe released as a component.
  * Unit + Integration tests
  * e2e tests?
* Automatically update vendored deps from consuming repos
* e2e tests run against main repo against integrated components

#### Process to move code into its own repo

#### Considerations
* Does the code depend on kubernetes/kubernetes code?
  * Break these dependencies or move the deps to another repo
* Does other code from kubernetes/kubernetes depend on the code?
  * Must update package names
* Can code be fully tested independent of kubernetes/kubernetes code?
* Will the code share the same release as kubernetes/kubernetes or be released independently?
  * Clients released independently need a way to perform version skew testing
 
#### Possible workflow
* Maybe first move to “staging” directory to simulate multi-repo within single repo - this has its own costs
* Remove cyclic dependencies between repositories
* Update mungebot to close PRs that update moved packages
* Copy code to new location
* Vendor code from new location to main repo (if needed)
* Change package names to the vendored code
* Delete old code

Where do we need to get to in order to separate: [ [diagram](https://files.slack.com/files-tmb/T09NY5SBT-F5NU1FPPG-0bb64d7351/image_uploaded_from_ios_1024.jpg) ]
* Utility packages need to go somewhere
* API packages need to move
* Helper libraries (reusable server infrastructure)
* Code generators
 
Consumer-visible
* client go
* kubectl
 
Need to get the API types into their own repo
* currently have a staging area, and a copy/paste script
Need to develop in those repos and then assemble a coherent release from them
 
re: kubectl moving,
* 1. sits in the repo but build rules break bad dependencies, can start shipping it independently
* 2. make sure everything still works
 
kubectl could get its own release
 
build files to add visibility
 
kubectl should never be imported
 
As these refactors happen, Brian Grant will be more lenient with privileges, establishing janitors

kubectl split is technical debt cleanup - testing should not rely on e2e, no compiled-in dependencies
 
breaking the dependencies yields more benefits
 
Q: Will kubectl be buildable with go get?  
A: Probably, but definitely with bazel
 
Integration testing plan?
Use last stable release in k/k e2e tests
 
We need to catch regressions when it’s being released, not when it’s vendored back in
 
Need better, smaller tests overall
 
kops has a blocking PR test - kops master + what is in your PR
 
cross-component problems at the main level
 
Need a central build pipe in order to untangle dependencies 
 
Experience breaking apart has been negative as a contributor
 
Q: What is the contribx?  Testing?
A: Short term will be many separate PRs, need to solve dependencies in test management
 
Please do not add new binaries in the kubernetes repo

Need help
* getting rid of generated code
* build system
* solve vendoring problem
 
Working group?  Code health WG
 
Client tool SDK?
 
kubectl -> SIG CLI
code -> API Machinery 
 
Caution against validating undocumented functionality via CI, better:  documentation and integration test as a validation ~ API testing

Too much testing is currently e2e
 
Don’t make a catchall util package

### Conclusions

#### Key Takeaways / Analysis of Situation
#### Recommendations & Decisions Moving Forward (High-Level)

Should build and release be the same SIG?    Where does build live?  -- SIG release / SIG testing 
 
Q: Is it worthwhile to go through the project and specify what needs to be moved out?
A: In the SIG assignment process there may be breakup


Action Item |Owner(s)
------------|------------
Need a SIG owner on utilities | Tim Hockin
Talk to Cloud Foundry and see how they do it | Caleb at Community
Working group organization or a SIG that oversees this | Machinery, CLI, Cluster Lifecycle
Regular community checkins | Brian Grant
Describe the end state and then assign breakout work to SIGs | Working group needs to be organized - Brian Grant

#### Instructions
1. Make a copy of the “Template” area of this document in the folder.
2. Set the “Share” settings to be “anyone with the link can edit” and take your session notes there. Ensure to decide/assign (and document) the topic title, facilitator(s), note-taker(s), and Github uploader.
3. Later, convert the notes to Markdown and upload to this GitHub directory: 
4. Save your session notes with the following format:
  a. HHMM-HHMM_SESSIONTITLE.md
    i. where HHMM is the time in 24-hour format in the PST zone.
 
 
#### Suggestions
* Use the document outline (Tools > Document Outline) and headings (Format > Paragraph Styles) to organize your notes as you go


[Link to original doc](https://docs.google.com/document/d/163JvzDIuBa4CzttxxG8tjYYPACQnD1DboW3BDG8VCUA/)
