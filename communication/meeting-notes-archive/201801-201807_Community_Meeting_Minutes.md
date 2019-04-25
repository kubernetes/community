**This is an archive document**


## July 26, 2018 ([recording link](https://youtu.be/XPt3ZwZe-VQ))



*   **Moderators**: Chris Short [ContribEx]
*   **Note Taker**: Solly Ross, Josh Berkus
*   **Demo:** EKS - Bryce Carman - [Amazon EKS] (confirmed)
    *   Managed Kubernetes on [https://aws.amazon.com/eks/](https://aws.amazon.com/eks/)
    *   Provisioning
        *   Control plane is hosted/managed by EKS, worker nodes are under control of users
            *   No outside communication with the control plane besides via the load balancer in front of the API server
            *   Can use security groups to limit control-plane-worker-node interaction
        *   Can set role used to create various AWS resources (like loadbalancers) so that you don't have to give EKS full permissions in your account
        *   Can just use VPC and subnets already present in account
    *   Networking
        *   CNI plugin
        *   Usines IP addresses from VPC that the nodes are already part of (integrated with AWS networking)
        *   No overlay network
        *   Can integrate with Calico network policy as well
        *   Designed to isolate control planes from nodes as well
    *   Interaction
        *   Using Heptio authenticator and 1.10 for external authentication for kubectl in order to authenticate against AWS IAM
        *   Just uses the same creds as the AWS CLI -- no separate auth to manage
    *   Demo'd using Helm to create a wordpress site
    *   Questions
        *   Can users scale control plane?
            *   No
*   **Release Updates**
    *   1.12 - Tim Pepper - Confirmed
        *   **Feature Freeze Tuesday July 31** - next week
            *   [see email on k-dev for more info](https://groups.google.com/d/topic/kubernetes-dev/T-kIHtgS5J4/discussion)
            *   After Tuesday features not captured by the release team must go through the **[exception process](https://github.com/kubernetes/features/blob/master/EXCEPTIONS.md)**.
            *   SIGs should be thinking about their release themes (major work focuses) for the 1.12 release, insuring those are represented in feature issues and have plans for documentation and test coverage.
            *   Not code freeze (that comes later)
    *   1.11.x - Anirudh Ramamathan - Confirmed
        *   Nothing to report
*   **KEP o' the Week **- KEP 17 - Jordan Liggitt - Confirmed
    *   [KEP 17 - Moving ComponentConfig API types to staging repos](https://github.com/kubernetes/community/blob/master/keps/sig-cluster-lifecycle/0014-20180707-componentconfig-api-types-to-staging.md)
    *   Taking config for core kube components from loose flags to structured config
        *   Kubelet currently has a config file format that's in beta
        *   Makes it easier to look at exactly how a particular component is configured, warn about deprecated config, missing config, etc
    *   Want to put configuration types in separate repo
        *   Tools like kubeadm should be able to import config to manipulate and generate, without pulling in all of Kubernetes
    *   Want to make sure common configuration aspects can be shared, referenced, and reused
        *   client connection info
        *   Leader election
        *   etc
    *   Look over if you are involved in developing the Kube components, or have tooling that sets up the various components
*   **SIG Updates**
    *   Auth - Jordan Liggitt - Confirmed
        *   [https://docs.google.com/presentation/d/1MAIypro-bcLC7wNEnIazYqmCL6ILBN69uUWIBw7QBIY/edit?usp=sharing](https://docs.google.com/presentation/d/1MAIypro-bcLC7wNEnIazYqmCL6ILBN69uUWIBw7QBIY/edit?usp=sharing)
        *   Usability
            *   Multiple Authorizers (e.g. GKE)
                *   Now honor superuser permissions from other authorizers, so if you're a superuser, you can create policy without first explicitly granting yourself those permissions
                *   Now show the error message from all authorizers, instead of just the error from the first authorizer
                *   Show a much cleaner, more succinct and readable error message for failures due to escalations
        *   Features
            *   Kubelet Certs
                *   Better support for delegating to an external credentials providers (e.g. AWS IAM)
                *   Requesting and rotating certs with the CSR API (still requires external approval process for the CSRs)
            *   Scoped service account tokens
                *   Moving towards beta for time-limited and audience-scoped tokens
            *   Audit improvements
                *   Heading towards v1 audit event API
                *   Work ongoing on dynamic audit webhook reg
    *   Instrumentation - Frederic Brancyzk - Confirmed
        *   Heapster deprecation ([https://github.com/kubernetes/heapster/blob/master/docs/deprecation.md](https://github.com/kubernetes/heapster/blob/master/docs/deprecation.md))
            *   Setup removal in 1.12, completely removed as of 1.13
        *   Node metrics work still ongoing, in collaboration with SIG Node
            *   Improve monitoring story around node monitoring
            *   Chime in if you maintain a device plugin or node component
        *   Metrics-server rework ([https://github.com/kubernetes-incubator/metrics-server/pull/65](https://github.com/kubernetes-incubator/metrics-server/pull/65))
            *   call for testing in non-production servers, should make things more stable, has several fixes to communication with nodes
        *   k8s-prometheus-adapter advance configuration merged
            *   Allows more precisely controlling how metrics in the custom metrics API map to Prometheus queries, and how metrics show up in the custom metrics API
        *   A number of third party service involving e2e tests have been put behind a feature flag in the test infrastructure
            *   should improve flaking tests from sig-instrumentation, especially around components that we can't control
*   **Announcements**
    *   **Shoutouts **(mention people on #shoutouts on Slack)
        *   Manjunath Kumatagi for patiently working through issues that will help us run conformance tests on other architectures (say `arm64`). It's taken a really long time to get this far and the end is in sight. Thanks for your hard work across multiple repos and sigs.
        *   Jordan Liggitt for always knowing the answer to ... everything ... and being so available to answer questions. You're an incredible resource and I'm always grateful to lean on you when I need to!
        *   Quinton Hoole for including a "how you can contribute" slide in the SIG Multicluster update in today's community!  Way to model SIG leadership in growing the k8s team by facilitating new/increased participation!


## July 19, 2018 ([recording link](https://youtu.be/XNLDZYMphuU))



*   **Moderators**: Tim Pepper [ContribEx, Release, VMware]
*   **Note Taker**: Solly Ross
*   **Demo:**  Microk8s - [Marco Ceppi](mailto:marco@ceppi.net) (confirmed)
    *   [https://microk8s.io](https://microk8s.io) / #microk8s on Slack / [https://github.com/juju-solutions/microk8s](https://github.com/juju-solutions/microk8s)
        *   Lightweight kubernetes cluster install
        *   Installed, uninstalled with snaps
            *   works across different linux distros, other OSes coming eventually)
        *   Still a bit in beta
            *   Different releases installed with different channels (beta channel is 1.11.0, edge is 1.11.1)
    *   Commands installed namespaced by `microk8s.`                
        *   kubectl is `microk8s.kubectl`
        *   Can enable different addons like dns and dashboard with `microk8s.enable`
            *   Cert generation, ingress, storage also available
        *   kubeconfig is scoped just to microk8s.kubectl, doesn't interfere with normal kubectl
        *   `microk8s.reset` resets to blank state
    *   Kubernetes run as systemd services
    *   Service Cluster IP addresses available as normal on host system
*   [ 0:09 ]** Release Updates**
    *   1.12 [Tim Pepper]
        *   Features collection underway
            *   [https://github.com/kubernetes/kubernetes/issues?q=is%3Aopen+is%3Aissue+milestone%3Av1.12](https://github.com/kubernetes/kubernetes/issues?q=is%3Aopen+is%3Aissue+milestone%3Av1.12)
            *   [https://github.com/kubernetes/features/issues?q=is%3Aopen+is%3Aissue+milestone%3Av1.12](https://github.com/kubernetes/features/issues?q=is%3Aopen+is%3Aissue+milestone%3Av1.12)
        *   **Feature freeze is July 31**
            *   Make sure your SIG features are up to date!**	**
    *   1.11.1 [Anirudh Ramanathan]
        *   After fixing initial release issues (pushing images, cherry picks), release is now out!
            *   [https://groups.google.com/forum/#!topic/kubernetes-announce/tMTjihgETUo](https://groups.google.com/forum/#!topic/kubernetes-announce/tMTjihgETUo)
        *   Issues encountered
            *   [https://github.com/kubernetes/release/issues/586](https://github.com/kubernetes/release/issues/586)
                *   Push permissions for cloud builder to gcr.io/k8s-image-staging (new staging bucket)
            *   [https://github.com/kubernetes/release/issues/587](https://github.com/kubernetes/release/issues/587)
                *   Autopush to release gcr bucket - Ben Elder
            *   Debs/RPMs
                *   Still needs a googler
        *   First non-Google release folks, updating docs to work around some last wrinkles
*   [ 0:13 ] **KEP o' the Week **- none this week
    *   If you want to get a broader audience for an up-and-coming KEP, you can get it discussed here!
*   [ 0:14 ] **SIG Updates**
    *   SIG Big Data (Anirudh Ramanathan, Yinan Li, confirmed)
        *   Deal with big data workloads on Kube
            *   Specifically: Spark, Spark Operator, Apache Airflow, HDFS
        *   Code freeze for Spark coming up, so lots of work there
            *   python support, client node support for things like Jupyter notebooks talking to Spark on Kubernetes)
            *   Stability fixes - better controller logic
                *   Making sure to be level triggered and not edge triggers
            *   Removing some hacks with init containers
        *   Spark ([link](https://issues.apache.org/jira/issues/?jql=project+%3D+SPARK+AND+component+%3D+Kubernetes))
            *   Working towards 2.4 release.
            *   2.4 code freeze and branch cut on 8/1
            *   Major features
                *   PySpark support
                *   Client mode - support for notebooks
                *   Lots of testing, merged integration tests
                *   Removal of things like init-containers (getting us closer to GA)
                *   Stability fixes - controller logic
                *   Improvements on client side
            *   Future work
                *   Customize pod templates
                *   Dynamic allocation/elasticity
                *   HA driver - might need help from sig-apps to make it work
                *   SparkR and Kerberized HDFS support
        *   Spark Operator new features ([link](https://github.com/GoogleCloudPlatform/spark-on-k8s-operator))
            *   Mutating admission webhook to replace initializer used before
            *   Python support
        *   HDFS support ([link](https://github.com/apache-spark-on-k8s/kubernetes-HDFS))
            *   Assessing demand, making progress
            *   Chart exists in link above
        *   Airflow ([link](https://cwiki.apache.org/confluence/pages/viewpage.action?pageId=71013666))
            *   Blog post went live - 28 June 2018
            *   [https://kubernetes.io/blog/2018/06/28/airflow-on-kubernetes-part-1-a-different-kind-of-operator/](https://kubernetes.io/blog/2018/06/28/airflow-on-kubernetes-part-1-a-different-kind-of-operator/)
    *   SIG Multicluster (Quinton Hoole, confirmed)
        *   [Slides](https://docs.google.com/presentation/d/1vcMLWEMRvg1rSrB1Ha-koxRZ9h1MUESW9q-c_cDP5n0/edit#slide=id.gc6f73a04f_0_0)
        *   Goals
            *   Solving common challenges releated to managing multiple clusters
            *   Applications that run across multiple clusters
        *   Subprojects
            *   Cluster Federation (v2) [[https://github.com/kubernetes-sigs/federation-v2](https://github.com/kubernetes-sigs/federation-v2) ]
                *   Work across different clusters, same or different cloud provider
                *   V1 was a POC, won't be developed further
                *   V2 focuses on decoupled, reusable components
                *   V2 has feature parity with v1, is alpha
                *   Highlights
                    *   CRDs for control planes, installed in existing cluster
                    *   Generic impl for all kube types (including CRDs) for propagating any types into all clusters, with basic per-cluster customization
                    *   Several higher-level controllers, for example:
                        *   migration of RS and deployments between clusters
                        *   Managing federated DNS
                        *   Management of Jobs
                        *   Management of HPA to manage global limits
                    *   Uses cluster registry
                *   Next steps
                    *   Federated status
                    *   Federated read access (e.g. view all pods across all clusters)
                    *   affinity/anti-affinity for bunches objects or namespaces to a particular cluster
                    *   RBAC enforcement
                *   Please comment if you have suggestions for API, before moves to beta
                    *   Contributions to code also welcome
            *   Cluster Registry [[https://github.com/kubernetes/cluster-registry](https://github.com/kubernetes/cluster-registry) ]
                *   Fairly stable and complete
            *   Multicluster Ingress [[https://github.com/GoogleCloudPlatform/k8s-multicluster-ingress](https://github.com/GoogleCloudPlatform/k8s-multicluster-ingress) ]
                *   Look at repo for more information
            *   Questions
                *   Cluster registry vs cluster API?
                    *   Cluster API is to create clusters, cluster registry is for using already-existing clusters
                    *   Maybe could disambiguate the terms better, manage overlap
    *   SIG Scheduling (Bobby Salamat, confirmed)
        *   1.11 Update
            *   Pod Priority and Preemption to beta, available by default
                *   Improved the feature, restricted a bit to avoid allowing untrusted users to create high-prio pods, only allow super-high-priority pods in kube-system namespace
            *   DaemonSet scheduling in default scheduler (alpha)
        *   1.12 Update
            *   Focus on performance
                *   Improved equivalence cache (pod with similar spec probably fits on same node unless the node has changed)
                    *   3x performance improvement now
                    *   Helps with scheduling large replica sets, etc
            *   Working on proposal for gang scheduling [link here]
            *   Proposal for scheduling framework, direction might change a bit [link here]
            *   Moving to beta
                *   Taint by condition, taint-based eviction
                *   Equivalence cache
                *   DaemonSet scheduling in default scheduler
            *   Want to graduate descheduler out of incubator
*   **Announcements**
    *   **Shoutouts **(mention people on #shoutouts on Slack)
        *   Jeremy Rickard: Shout out to @mbauer for really pushing to make service catalog use prow and to improve out PR reviewing and testing process
        *   Christoph Blecker: Two shoutouts I wanted to get out this week:
            *   First, shoutout to @matthyx who has been very active in k/test-infra recently and has been making a number of different contributions from fixing bugs, to adding new features to our automation. He's been eager to help and has stuck with some of the more complex changes that require many comments and interactions (sig-bikeshed ftw :bikeshed:)
            *   Second, shoutout to @nikhita! I could easily stop right there, as her many contributions to the project really speak for themselves. I want to call out though the little chopping wood and carrying water tasks she does that may not be as obvious.. like ensuring that stale issues are reviewed and either closed or marked as still relevant, or welcoming new contributors with an emoji or two. It's these kinds of things that exemplify what the Kubernetes community is all about.
        *   Benjamin Elder:  shoutout to @Quang Huynh for continuing to send k/test-infra fixes and push through to flesh out the PR status page (especially https://github.com/kubernetes/test-infra/pull/8612) long after his internship! :simple_smile: Hopefully we can hopefully start using the Prow PR status page more widely now thanks to all the hard work there :tada:
        *   Aaron Crickenberger: shoutout to @bentheelder for helping push kubernetes v1.11.1 images out (fixing the symptom), and getting the appropriate folks within google involved to ensure there is now a team owning a better solution to the problem (fixing the problem); this is continued progress toward decoupling google.com as a requirement for releases
    *   [Kubernetes wins most impact award at OSCON](https://twitter.com/oscon/status/1019992011858894849)!!! (-paris; Tim to read)


## July 12, 2018 ([Recording](https://youtu.be/OBubmJhr8lE))



*   **Moderators**: Paris Pittman [ContribEx, Google]
*   **Note Taker**: Josh Berkus [Release]
*   **Demo:** No demo today - see you next week!
*   [ 0:01 ]** Release Updates**
    *   1.12: [Tim Pepper]
        *   Release cycle is underway!
            *   Team: [https://git.k8s.io/sig-release/releases/release-1.12/release_team.md](https://na01.safelinks.protection.outlook.com/?url=https%3A%2F%2Fgit.k8s.io%2Fsig-release%2Freleases%2Frelease-1.12%2Frelease_team.md&data=02%7C01%7Ctpepper%40vmware.com%7C89885a0f477d469ddcb308d5e6ba8e88%7Cb39138ca3cee4b4aa4d6cd83d9dd62f0%7C1%7C0%7C636668611291496196&sdata=S%2FDlfHfkgz1G9NkGVMtjDIa%2FY%2BkzpAkKwpOU0HgaGfc%3D&reserved=0)
            *   Schedule: [https://git.k8s.io/sig-release/releases/release-1.12/release-1.12.md](https://na01.safelinks.protection.outlook.com/?url=https%3A%2F%2Fgit.k8s.io%2Fsig-release%2Freleases%2Frelease-1.12%2Frelease-1.12.md&data=02%7C01%7Ctpepper%40vmware.com%7C89885a0f477d469ddcb308d5e6ba8e88%7Cb39138ca3cee4b4aa4d6cd83d9dd62f0%7C1%7C0%7C636668611291506201&sdata=ZzUn8hRsKf3E1poF2i5%2BVgmpm0UYnM3rZh1iiy%2Br1QM%3D&reserved=0)
        *   Features collection is happening now, see:
            *   [https://github.com/kubernetes/features/issues?q=is%3Aopen+is%3Aissue+milestone%3Av1.12](https://github.com/kubernetes/features/issues?q=is%3Aopen+is%3Aissue+milestone%3Av1.12)
            *   [https://groups.google.com/forum/#!topic/kubernetes-dev/T-kIHtgS5J4/discussion](https://groups.google.com/forum/#!topic/kubernetes-dev/T-kIHtgS5J4/discussion)
        *   Key dates:
            *   Feature freeze:             	July 31, 2018
            *   Begin code slush:         	Aug. 28, 2018
            *   Begin code freeze:      	Sept. 4, 2018
            *   End code freeze:          	Sept. 19, 2018
            *   Release date:                 	Sept. 25, 2018
        *   Process changes:  Nothing notable versus 1.11, will continue shortened code slush/freeze, BUT this depends on us all keeping a clean [CI Signal](http://testgrid.k8s.io/) throughout the release cycle.  Additionally asking for increased focus on earlier:
            *   Definition of blocking test cases and test case additions
            *   Drafting documentation for feature changes
    *   1.11.1: [Anirudh Ramanathan]
        *   Announcement: [https://groups.google.com/forum/#!topic/kubernetes-dev-announce/vdodsSq21qc](https://groups.google.com/forum/#!topic/kubernetes-dev-announce/vdodsSq21qc)
        *   Will cut Monday July 16th.
        *   Freezing branch today EOD
        *   Some cherrypicks still need some action: please check yours.
            *   [https://docs.google.com/document/d/1kFHQsk1iM9rh0iEaLhaNnAtvtU00ZSlZgb1etmdmcTQ/edit](https://docs.google.com/document/d/1kFHQsk1iM9rh0iEaLhaNnAtvtU00ZSlZgb1etmdmcTQ/edit)
            *   If your PR is marked in pink/orange, it might need action on your part.
        *   Kops test was blocking this morning.
*   [ 0:07 ] **KEP o' the Week** (Janet Kuo)
    *   [https://github.com/kubernetes/community/pull/2287](https://github.com/kubernetes/community/pull/2287)
    *   For cleanup of frequently created & dropped objects
        *   We don't have a good way to garbage collect items which no longer have an owner.
        *   Often people update-and-replace instead of modifying
    *   Proposal for new GC for these objects.
        *   Will be discussed in next API-machinery meeting next week (Wednesday) if you care about the KEP
        *   Will give detailed presentation there.
*   [ 0:00 ] **SIG Updates**
    *   SIG API Machinery [David Eads](confirmed)
        *   Link to [slides](https://docs.google.com/presentation/d/171PN2zg5iMXZ18LwYTEUe0Jg_-IxFDSIsbHW09F9mZk/edit#slide=id.g3d7994f0d0_0_0)
        *   Delivered in 1.11:
            *   Improved dynamic client, easier to use for CRD developers.  Everyone should switch to this because the old client will eventually go away.
            *   "Null CRD conversion": you can promote a CRD from one version to another, even though there's no API changes.  No data transformation, no changes to schema.  So very limited for now.
            *   Work on feature-branch for Server-side apply.
            *   Prep work for making controller-manager start from a config
        *   1.12 work
            *   Server-side apply dry run being merged into Master
            *   Path to more advanced CRD conversion, field defaults, advanced versioning (design phase).
            *   Controller-manager moving to running from config
            *   Generic initializers as alpha.  May be superseded by admission webhooks.  If you need something in Generic Init that isn't satisfied by webhooks, speak up in their meeting to save it.
    *   SIG Testing [Steve Kuznetsov] (confirmed)
        *   Link to [slides](https://docs.google.com/presentation/d/10v9MoXOjEJQ9opoIffNm6XwEIaUtYpNcCRvBuGaS-zs/edit?usp=sharing)
        *   Implemented caches for test runs, which is a big performance boost
            *   Reduced GH API hits by 1500/hr
            *   Bazel build cache lowered test times
        *   UX improvements for the k8s bot
            *   Now can LTGM and approve in review comment
            *   Robots now validate OWNERS files
        *   Easier administration
            *   Using Peribolos for GH API management
        *   Automated branch protection now on all repos
            *   Only bots can merge to branches
        *   Simpler test Job management: now just needs a container with an entrypoint
        *   Merge workflows using Tide are implemented
            *   Plan to rollout for 1.12
            *   Will include PR status page, yay!  Makes it easier to see why your PR is stuck.
        *   Testgrid dashboard for conformance tests
            *   Including openstack
        *   Prow is now being adopted by other orgs
            *   Google, Red Hat, Istio, JetStack …
        *   Future work:
            *   Better onboarding docs
            *   Fix tech debt that makes getting started hard
            *   Better log viewer, esp now that we have scalability presubmits
            *   Clean up config repo
            *   Framework for writing bot interactions
            *   API for cluster provisioning
        *   Questions:
            *   What about archival stats on PR status dashboard?
                *   Will discuss at sig-testing meeting
            *   What about doc on how to write a test?
                *   Also really critical, needs help
    *   SIG ContribEx [Paris](confirmed)
        *   [Link to slides](https://docs.google.com/presentation/d/1z1Cscr-cOpX9b7vUqdRfWvtd6Yb_Ybm16_G_nVXTssI/edit?usp=sharing)
        *   Contributor Guide
            *   Umbrella issue is now closed
            *   non-code guide in development - meets on Weds
        *   Developer Guide
            *   Tim Pepper now taking point on this
            *   Reach out to him @tpepper if you can help
        *   Contributor.kubernetes.io web site is under early design
            *   Different from general community, this one will be just for contributors
            *   More modern calendar
            *   Prototype up, check it out (link from slides)
            *   goal to launch in 90 days
        *   Community Management
            *   All talking all the time, it's time consuming
            *   Contributor summits, first one (run by contribex) in Copenhagen
                *   Rolling out new contributor workshop + playground
                *   Will have smaller summit in Shanghai (contact @jberkus)
                *   Started planning for Seattle, will have an extra ½ day.
                    *   Registration will be going through KubeCon/CloudNativeCon site
                *   Manage alacarte events at other people's conferences
            *   Communication pipelines & moderation
                *   Clean up spam
                *   Reduce number of pipelines
                *   Some draft moderation guides
                *   Also run the Community Meeting
                *   Zoom has a bad actor problem, so we're not locking down Zoom permissions, trying not to take away public meetings, looking at new security together with Zoom execs.
                *   Moderating k-dev and k-users MLs now
                *   If you need to reach moderators quickly, use slack-admins slack channel
                *   Slack: 40K users, a lot less moderation required
                *   Discuss.kubernetes.io
                    *   Been successful for tips & tricks and user advice
                    *   Will be "official" RSN
        *   Mentoring
            *   Meet Our Contributors is doing well
                *   Yesterday's special edition had Steering Committee members
            *   Outreachy, only participating twice a year
                *   September deadline for winter intern, planning on 1
                *   Participating companies can pay for more
            *   Group mentoring: the 1:1 hour
                *   If your SIG needs to move people up to approver, please contact @paris
            *   GSoC, being done by API-machinery
        *   DevStats
        *   Github Management proposed subproject
*   Announcements
    *   **Shoutouts** - enter yours in #shoutouts slack channel!
        *   (jberkus) - @jdumars for inventing, then running, really effective retros for releases.
        *   (paris) - shouts to @liggitt @stevekuznetsov and @munnerz for a jam packed, informative, #meet-our-contributors [session yesterday](https://youtu.be/EA6s09YXgh8)! (watch the recording; good info!)
        *   (paris) another shout to @arschles @janetkuo for being mentors on the second great episode of #meet-our-contributors yesterday. (also to bdburns, pwittroc, and philips but I will spare their notifications for doing our first [AMA with steering committee members](https://youtu.be/BuJhzJriaNY))
        *   munnerz - shout out to @stevekuznetsov for immediately jumping to spend time debugging and fixing issues with our Prow deployment and tide (not) merging our PRs! looking forward to finally rolling the fix out :slightly_smiling_face: it has caused us issues for 1-2 months now :smile:
    *   **[Office Hours](https://github.com/kubernetes/community/blob/master/events/office-hours.md) is next Wednesday** - volunteers to help answer user questions are always appreciated, ping @jeefy or @mrbobbytables if you want to help, otherwise help us spread the word!
    *   Next week's meeting won't be streamed, so expect a slight delay on publishing it to YouTube


## July 5, 2018

**NO MEETING**


## June 28, 2018 - ([recording](https://youtu.be/aTNNtJ56ahE))



*   **Moderators**: Jaice Singer DuMars [SIG PM/Release]
*   **Note Taker**: First Last [Company/SIG]
*   [ 0:00 ]**  Demo **- containerd - Phil Estes - estesp@gmail.com
    *   Link to [slides](https://docs.google.com/presentation/d/19ZHjXR1uG4wdW5uXiNB7fda2goRBlSbBaV5Cw06a3zk/edit?usp=sharing)
    *   Link to [repositories](https://github.com/containerd)
*   [ 0:00 ] **Announcements**
    *   SIG IBMCloud, Autoscaling, and GCP will be updating in August
    *   Github Groups [Jorge Castro]
        *   [https://github.com/kubernetes/community/issues/2323](https://github.com/kubernetes/community/issues/2323)  working to make current 303 groups in the org easier to manage
    *   Shoutouts this week (Check in #shoutouts on slack)
        *   jberkus: To Jordan Liggitt for diagnosing & fixing the controller performance issue that has haunted us since last August, and to Julia Evans for reporting the original issue.
            *   Maulion: And another to @liggitt for always helping anyone with a auth question in all the channels with kindness
        *   jdumars: @paris - thank you for all of your work helping to keep our community safe and inclusive! I know that you've spent countless hours refining our Zoom usage, documenting, testing, and generally being super proactive on this.
        *   Nikhita: shoutout to @cblecker for excellent meme skills!
        *   Mrbobbytales: Just want to give a big shout out to the whole release team. Thanks for all your effort in getting 1.11 out the door :slightly_smiling_face: Seriously, great job!
        *   Misty: @chenopis for last-minute 1.11 docs-related heroics!
        *   Misty: @nickchase for amazing release notes!
        *   Misty: @jberkus for being a very patient and available release lead as I was on the release team for the first time
        *   Jberkus: @liggitt for last-minute Cherrypick shepherding, and @nickchase for marathon release notes slog
        *   Jberkus: and @misty @AishSundar @tpepper @calebamiles @idvoretskyi @bentheelder @cjwagner @zparnold @justaugustus @Kaitlyn for best release team yet
        *   Tpepper: shoutout to @jberkus for his leadership of our team!
*   [ 0:00 ]** Release Retrospective for 1.11**
    *   [Retro doc](https://docs.google.com/document/d/1Kp9J29wCTY_3SdQn0Kmpuw9lOSxoO7BYWDUcmSoCrZo/edit#)
        *   SIG Release will do deep dive on retrospective details, but today this meeting focused on the high level cross-project topics like:
            *   Release timeline evolution and deadlines
            *   How to better track major features and changes that are in need of docs, test cases, release noting
            *   How do we get user/distributor/vendor testing of betas and rc's.  Consumption is harder when docs and kubeadm upgrade path aren't there yet.
    *   Retro Part II (detail retro): Tuesday, July 3rd, 10am, [https://zoom.us/j/405366973](https://zoom.us/j/405366973)


## July 15, 2018 - recording



*   [Stackoverflow Top Users](https://stackoverflow.com/tags/kubernetes/topusers) for June 2018, thanks for helping out!
        *   [Matthew L Daniel](https://stackoverflow.com/users/225016/matthew-l-daniel)
        *   s[uren](https://stackoverflow.com/users/5564578/suren)
        *   [Janos Lenart](https://stackoverflow.com/users/371954/janos-lenart)
        *   [Nicola Benaglia](https://stackoverflow.com/users/2718151/nicola-benaglia)
        *   [Ignacio Millán](https://stackoverflow.com/users/9811836/ignacio-mill%c3%a1n)


## June 21, 2018 - ([recording](https://youtu.be/VmVh2TsRP-s))



*   **Moderators**:  Arun Gupta [Amazon / SIG-AWS]
*   **Note Taker**: Chris Short and Jorge Castro [SIG Contrib Ex]
*   [ 0:00 ]**  Demo **-- Agones - Dedicated Game Server Hosting and Scaling for Multiplayer Games on Kubernetes [Mark Mandel, markmandel@google.com] (confirmed)
    *   [https://github.com/GoogleCloudPlatform/agones](https://github.com/GoogleCloudPlatform/agones)
*   [ 0:00 ]** Release Updates**
    *   1.11 [Josh Berkus - Release Lead]
        *   Code Thaw on Tuesday, held changes from Code Freeze have now cleared the queue.
            *   All 1.11 changes now need to be cherrypicked.
        *   RC1 was released yesterday, please test!
        *   Status is currently uncertain.  Probability of a release delay is 50%, will make call at Burndown meeting 10am tomorrow.
        *   CI Signal Issues:
            *   GKE appears to have pushed a [change breaking tests](https://github.com/kubernetes/kubernetes/issues/65311) at midnight last night, currently sorting whether that's just a GKE problem.
            *   [Upgrade tests are still very flaky](https://k8s-testgrid.appspot.com/sig-release-master-upgrade),  this seems to be an artifact of the tests and not of code.  GCE/GKE staff have given the go-ahead to release without clean signal as they will not be fixing the tests.
            *   [Alpha tests failing](https://k8s-testgrid.appspot.com/sig-release-1.11-blocking#gce-alpha-features-1.11) due to [Daemonset issue](https://github.com/kubernetes/kubernetes/issues/65192#issuecomment-398541158); currently trying test resource change.
        *   Release notes collector is still broken, please **check[ the release notes](https://github.com/kubernetes/sig-release/blob/master/releases/release-1.11/release_notes_draft.md) to make sure all of your changes represented**!  An estimated 20-30 release notes are missing.  Contact (@nickchase / [nchase@mirantis.com](mailto:nchase@mirantis.com)) if you find something missing.
    *   1.12 [Tim Pepper - Release Lead]
    *   Patch Release Updates
        *   1.8.14
        *   1.9.9 release schedule
        *   1.10.5
*   [ 0:00 ] **KEP o' the Week** (Yisui)
    *   Namespace Population is an automated mechanism to make sure the predefined policy objects (e.g. NetworkPolicy, Role, RoleBinding) are present in selected namespaces.
    *   [https://github.com/kubernetes/community/pull/2177](https://github.com/kubernetes/community/pull/2177)
*   [ 0:00 ] **Open KEPs** [Kubernetes Enhancement Proposals]
    *   [Check out the KEP Tracking Board](https://github.com/kubernetes/community/projects/1)
*   [ 0:00 ] **SIG Updates**
    *   SIG Big Data [Sean Suchter] (confirmed)
        *   Slides: [https://docs.google.com/presentation/d/1GE9wh7Lja1vFHFCJdH45PAjjEVJBmhhXvnrjjq5jpnk/edit#slide=id.g3bb158eee1_0_0](https://docs.google.com/presentation/d/1GE9wh7Lja1vFHFCJdH45PAjjEVJBmhhXvnrjjq5jpnk/edit#slide=id.g3bb158eee1_0_0)
        *   k8s 💗Spark
            *   [Spark on Kubernetes Talk](https://databricks.com/session/apache-spark-on-kubernetes-clusters)
        *   ASF is going to publish official Spark container images
        *   NEED: Parity with Spark on YARN (esp. scheduling)
    *   SIG PM [Stephen Augustus]  (confirmed)
        *   Slides: [https://docs.google.com/presentation/d/1p8FF45r0CF-AYnGm_B59F2PN_jybd0bbLe0fb0APlJU/edit?usp=drivesdk](https://docs.google.com/presentation/d/1p8FF45r0CF-AYnGm_B59F2PN_jybd0bbLe0fb0APlJU/edit?usp=drivesdk)
    *   SIG Azure [Stephen Augustus] (confirmed)
        *   Slides: [https://docs.google.com/presentation/d/15M-bQdMxaY6ZBedEEt9yKNesfv_95NYHBHIRwCYAP2E/edit?usp=drivesdk](https://docs.google.com/presentation/d/15M-bQdMxaY6ZBedEEt9yKNesfv_95NYHBHIRwCYAP2E/edit?usp=drivesdk)
*   [ 0:00 ] **Announcements**
    *   Please pin your SIG meeting info and agenda doc in your SIG slack channel.  Now that the main calendar is not on [https://kubernetes.io/community/](https://kubernetes.io/community/) meeting info is less discoverable without these links.
        *   **SIG Chairs/TLs - please check your email (sent to k-sig-leads@). New zoom settings and moderation controls. Let's keep our meetings safe and transparent. **
    *   All SIGs - please take time to look at the "help wanted" and "good first issue" labels, available across all Kubernetes repositories. They're meant to highlight opportunities for new contributors. Please ensure that they're being used appropriately (the "good-first-issue" especially has fairly specific requirements for the issue author): [https://github.com/kubernetes/community/blob/master/contributors/guide/help-wanted.md](https://github.com/kubernetes/community/blob/master/contributors/guide/help-wanted.md)
    *   Shoutouts this week (Check in #shoutouts on slack)
        *   Jason DeTiberus: @neolit123 (Lubomir Ivanov) for all of the docs contributions for kubeadm v1.11
        *   Jason DeTiberus: @jrondeau (Jennifer Rondeau) for the relentless work on improving our docs and helping bring some more structure to the docs process for sig-cluster-lifecycle
        *   @neolit123 (Lubomir Ivanov): @jdetiber (Jason DeTiberus), @liz (Liz Frost), @cha (Chuck Ha), @timothysc (Timothy St. Clair) and @luxas (Lucas Kladstrom) for the relentless grind trough kubeadm 1.11 backlog potentially making it the best release thus far.
        *   @austbot (Austin Adams): To @lukaszgryglicki (Lukasz Gryglicki) for DevStats, which is Awesome!!
        *   Stealthybox (Leigh Capili): shoutout to @oikiki (Kirsten) for being very welcoming to new contributors
        *   Nikhita: shoutout to the whole test-infra community for actively using emojis in issues, PRs and slack. It's pretty subtle but it goes a LONG way in making the project and community more friendly and welcoming to new contributors!! cc @fejta (Erick Fejta) @bentheelder (Benjamin Elder) @cblecker (Christoph Blecker) @stevekuznetsov (Steve Kuznetsov)
        *   @misty (Misty Stanely-Jones): @Jesse (Jesse Stuart) for fixing CSS relating to tab sets in docs! :raised_hands:
        *   @fejta (EricK Fejta): @krzyzacy (Sen Lu) and @bentheelder (Benjamin Elder) for being ever diligent about reviewing PRs in a timely manner
        *   JoshBerkus: to @kjackal (Konstantinos) for actually beta-testing 1.11 and spotting a bug before RC1
        *   @oikiki (Kirsten): shoutout to @gsaenger for always generously helping new folks get started contributing to k8s! (and also for completing her first major technical PR!) WOOP WOOP!
        *   @gsaenger (Guinevere Senger) Um... no, really, I couldn't have done it without so much help from @cblecker (Christoph Blecker) and @cjwagner (Cole Wagner) and @fejta (Erick Fejta)and @bentheelder (Benjamin Elder). Everyone was super nice and patient and helped me learn. :heart: So, shoutouts to them. I'm so grateful.


## June 14, 2018 - ([recording](https://youtu.be/yAtOHS6C-W0))



*   **Moderators**:  Zach Arnold [Ygrene Energy Fund/SIG Docs]
*   **Note Taker**: Jorge Castro [Heptio/SIG Contribex] and Solly Ross [Red Hat/SIG Autoscaling]
*   [ 0:00 ]**  Demo **-- Building Images in Kubernetes [Priya Wadhwa, priyawadhwa@google.com] (confirmed)
    *   [https://github.com/GoogleContainerTools/kaniko](https://github.com/GoogleContainerTools/kaniko)
    *   [https://docs.google.com/presentation/d/1ZoiQ3cuQNJJciKq_JvqTty_tcoaRKNyYRzgCBbTumsE/edit?usp=sharing](https://docs.google.com/presentation/d/1ZoiQ3cuQNJJciKq_JvqTty_tcoaRKNyYRzgCBbTumsE/edit?usp=sharing)
    *   Tool for building container images without needing to mount in Docker socket
        *   Extracts base image to file system
        *   Downloads build context tarball from storage (e.g. S3, more on the way)
        *   Executes commands listed in Dockerfile
        *   Snapshots in userspace after each step
        *   Ignores mounted directories during snapshots
    *   Can be run in gVisor as well
    *   Questions:
        *   do you have to use dockerfiles, or can you use other instruction sets
            *   Only dockerfiles right now, but file issues if you want other things
        *   Which dockerfile verbs are supported?
            *   All of them
        *   Can the bucket be S3 or DO Space?
            *   Working on a PR right now to support other solutions
        *   Feature parity with docker build?
            *   Yes
    *   Link to slides.
*   [ 0:00 ]** Release Updates**
    *   1.11  [Josh Berkus - Release Lead]
        *   **_Next Deadline: RC1 and branch on June 20th_**
        *   Less than a week of code freeze left!
        *   **Docs are due** and overdue; if you have a feature in 1.11,_ you should have already submitted final docs_. Contact the docs team.
        *   CI signal is good, a few tests being flaky, especially alpha-features.
        *   Only 2 issues and 6 PRs open; currently more stable than we've ever been! Thanks so much to everyone for working to get stuff in the release early.
    *   1.12 [Tim Pepper - 1.12 Release Lead]
        *   Tim Pepper as Lead
        *   Almost finished building 1.12 team, contact @tpepper on Slack to join.
            *   Needed:
                *   PR triage (tentatively adding role separate from issue triage)
                *   Branch manager
    *   Patch Release Updates
        *   1.10.4?
*   [ 0:00 ] **KEP o' the Week**
    *   SIG Cloud Provider KEP: Reporting Conformance Test Results to Testgrid [Andrew Sy Kim]
        *   Formerly WG-Cloud-Provider
            *   Standards and common requirements for Kubernetes cloud provider integrations
            *   Improving docs around cloud providers (how to work with different integration features)
            *   Improving testing of cloud providers
        *   KEP is basically "Why we want conformance tests reported by the cloud providers"
            *   We didn't have a formal way to do this without KEP
            *   SIG Testing infra wasn't available back then, so now we have testgrid and a way to report tests, etc. Gives providers instructions to follow to contribute results.
            *   SIG Openstack has been pioneering this work
            *   We want all providers to do this eventually, we'll be reaching out to all the cloud providers to give them visibility that this KEP exists.
            *   Still missing some details, will address those as more experience is developed in how to do better test
        *   Q:
            *   Coverage is listed as out of scope, but is a benefit, will coverage improvements be a follow-on KEP?
                *   Eventually, but currently not necessarily an immediate priority
    *   [https://github.com/kubernetes/community/pull/2224](https://github.com/kubernetes/community/pull/2224)
*   [ 0:00 ] **Open KEPs** [Kubernetes Enhancement Proposals]
    *   [Check out the KEP Tracking Board](https://github.com/kubernetes/community/projects/1)
*   [ 0:00 ] **SIG Updates**
    *   SIG Windows [Patrick Lang]
        *   [Trello board](https://trello.com/b/rjTqrwjl/windows-k8s-roadmap) - maps K8s features to Windows release needed
        *   Releasing twice a year in the Windows Server Semi Annual Channel
            *   Like 18.03, 17.09, etc.
            *   We've had to make changes to Windows Server to make Kubernetes work well. For example, symbolic links in Windows v. Unix.
            *   Board is tagged with the right version of Windows to use to get a particular Kubernetes feature working, but in general, use the latest release if possible
        *   Kube 1.11
            *   Lots of features with Windows, e.g. Kubelet stats
        *   For the future
            *   Currently using dockershim, trying to figure out how to support other CRI implementations
            *   Working with other CNI plugins (Flannel, OVN, Calico)
            *   Trying to get support for showing test results via Prow, Kubetest, TestGrid
            *   Want to move to GA eventually, with 2019 Windows release (extended support cycle)
        *   Questions:
            *   Q: Unix-style symlinks in Windows?
                *   Have something similar to unix-style symlinks and hardlinks, needed to make some symlink changes to make sure you can't traverse in an insecure way.  Code either in Kubelet or Go winio library
                *   Hardlinks not recommended, stick to the symlinks.
            *   Q: is windows currently dockershim+embedded EE
                *   Currently uses Docker EE Basic for Windows (as published by Docker), used for testing
                *   Potentially switching to crio eventually, or containerd
    *   SIG Apps [Kenneth Owens]
        *   Helm
            *   Helm moved to a separate CNCF project (see last meeting)
            *   Helm 2 stability release
            *   Helm 3 proposal merged, work continuing
        *   Application Resource
            *   Seeks to describe application as running
            *   Controller soon
            *   WIP
        *   AppDef WG
            *   Winding down
            *   Proposal for common labels and annotations will be merged in partial form
        *   Ksonnet
            *   New release
        *   Decentralized charts repo coming
        *   Skaffold:  kustomize support
        *   Workloads API
            *   Need to stabilize Job
            *   Want really make sure cron jobs are stable before moving cronjobs out of beta
            *   Job [first class sidecar container KEP](https://github.com/kubernetes/community/issues/2148) discussion ongoing
        *   Questions:
            *   Why didn't charts go with Helm to a separate CNCF project
                *   Current Status: Charts are listed as a subproject of SIG Apps.
                *   Chart maintainers aren't necessarily Helm maintainers
                *   Trying to figure out the right model for maintainership
                *   Is the charts tooling part of the charts subproject, or Helm?
                    *   Unsure, currently part of the charts subproject, but points to the kubernetes/helm repo
    *   SIG Docs [Jennifer Rondeau]
        *   Zach is out sick today, fulll update Augustish, feel better Zach!
        *   1 minutes Jennifer update
            *   We're making great progress on fixes for the hugo migration, we've plowed through a bunch, thanks to all the new contributors who have been diving in.
            *   Thanks to all of you who have submitting 1.11 docs
            *   **If you're behind on 1.11 docs, please submit them asap!**
*   [ 0:00 ] **Announcements**
    *   [K8s Office Hours](https://github.com/kubernetes/community/blob/master/events/office-hours.md) Next Week, Wednesday 6/20
        *   Volunteers always sought, ping @jorge or @mrbobbytables on slack
        *   Users who participate will be entered in a raffle to win a k8s shirt!
    *   SIG Leads, if you haven't uploaded your meeting videos to the youtube channel recently, please try to catch up. Ping @jorge if you need help.
    *   SIG Architecture has a new meeting time at 11PST every other Thursday after this meeting. Also, there is a new Zoom link you can get from joining the mailing list. Check out the[ SIG Arch readme for more information](https://github.com/kubernetes/community/tree/master/sig-architecture).
    *   Shoutouts this week (Check in #shoutouts on slack)
        *   (Josh Berkus) @liggitt and @dims for pitching in and doing a ton of work on PRs for 1.11, across all of Kubernetes.
        *   (Jennifer Rondeau) @misty for stepping in to help with ALL things docs no matter how crazy they get or how much else she has on her plate :tada:
        *   (Aish Sundar) @justaugustus for giving us a huge head start and herding all the cats to get a stellar 1.12 release team already in place. Thanks a lot!
        *   (Misty Stanley-Jones + Aish Sundar) @jberkus for herding 1.11 release cats! :cat:
        *   To echo what @misty said, HUGE shoutout to @jberkus for being an awesome patient leader throughout 1.11 cycle. It was such a learning experience seeing him work through issues calmly, all the while encouraging the RT team to lead in our own little way.
        *   Jason DeTiberius
            *   @neolit123 (Lubomir Ivanov) for all of the docs contributions for kubeadm v1.11
            *   @jrondeau (Jennifer Rondeau) for the relentless work on improving our docs and helping bring some more structure to the docs process for sig-cluster-lifecycle
    *   


## June 7, 2018 - ([recording](https://youtu.be/fOhby7EUiuo))



*   **Moderators**:  Jaice Singer DuMars [SIG Release/Architecture]
*   **Note Taker**: Austin Adams [Ygrene Energy Fund]
*   [ 0:00 ] **Demo** -- YugaByte ~ Karthik Ranganathan [[karthik@yugabyte.com](mailto:karthik@yugabyte.com)] (confirmed)
    *   Karthik Ranganathan
        *   GitHub:[ https://github.com/YugaByte/yugabyte-db](https://github.com/YugaByte/yugabyte-db)
        *   Docs:[ https://docs.yugabyte.com/](https://docs.yugabyte.com/)
        *   Slides: https://www.slideshare.net/YugaByte
        *   Yugabyte is a database focusing on, planet scale, transactional and high availability. It implements many common database apis making it a drop in replacement for those DBs. Can run as a StatefulSet on k8s. Multiple db api paradigms can be used for one database.
        *   No Kubernetes operator yet, but it's in progress.
    *   Answers from Q&A:
        *   @jberkus - For q1 - YB is optimized for small reads and writes, but can also perform batch reads and writes efficiently - mostly oriented towards modern OLTP/user-facing applications. Example is using spark or presto on top for use-cases like iot, fraud detection, alerting, user-personalization, etc.
        *   q2: operator in the works. We are just wrapping up our helm charts[ https://github.com/YugaByte/yugabyte-db/tree/master/cloud/kubernetes/helm](https://github.com/YugaByte/yugabyte-db/tree/master/cloud/kubernetes/helm)
        *   q3: the enterprise edition does have net new DB features like async replication and enforcing geographic affinity for reads/writes, etc. Here is a comparison:[ https://www.yugabyte.com/product/compare/](https://www.yugabyte.com/product/compare/)
        *   q4: You cannot write data using redis and read using another API. Its often tough to model across api's. Aim is to use a single database to build the app, so support common apis
        *   The storage layer is common
        *   So all APIs are modeled on top of the common document storage layer
        *   The API layer (called YQL) is pluggable
        *   Currently we model Redis "objects" and Cassandra "tables" on top of this document core, taking care to optimize the access patterns from the various APIs
        *   We are working on postgres as the next API
*   [ 0:00 ]** Release Updates**
    *   1.11  [Josh Berkus - Release Lead]
        *   **_Next Deadline: Docs Complete, June 11_**
            *   All listed features have docs in draft -- Thanks!
            *   However: non-listed (minor) changes, please make sure you have docs!
        *   [Currently](https://github.com/kubernetes/sig-release/blob/master/releases/release-1.11/release-1.11.md) in Code Freeze
            *   Only 1.11 patches, must be approved and critical-urgent
            *   Down to 11 PRs
            *   Still using the old Milestone Munger, so expect the same annoying behavior, sorry.
                *   Particularly: can't take back-branch PRs.
            *   <span style="text-decoration:underline;">No New Features/Cleanups Now, please</span>
                *   All new features have draft documentation, however, there are lots of small patches **not big enough** to be a feature but we don't know if we have documentation for those.
                *   **Please make sure your 1.11 small patches have documentation.**
            *   Code freeze ends June 19th.
            *   Docs need to be complete by June 11th
        *   CI Signal looking good
            *   Recent GKE breakage fixed.
                *   Only upgrade/downgrade tests failing, PR in progress.
            *   Thanks everyone for responding to test fails quickly!
        *   Scalability/Performance
            *   Currently passing all performance tests
                *   Thanks to everyone who worked on this early in the cycle!
            *   New performance presubit test
                *   Kudos to SIG-scalability for getting this done.
    *   1.12
        *   Currently working on[ forming a 1.12 Release Team](https://github.com/kubernetes/sig-release/issues/167)
            *   Interested?  Comment on the PR or speak up in #sig-release
*   **[ 0:00 KEP Highlight ] **- Kustomize [ Jeff Regan ]
    *   overall process: [https://github.com/kubernetes/community/tree/master/keps ](https://github.com/kubernetes/community/tree/master/keps)
        *   Kustomize
            *   Kustomize is a way for us to provide a declarative way to update resources in kubernetes. This allows us to version control changes to k8s configs and resources and so forth.
            *   Sig cli is sponsoring this project.
            *   **[PR](https://github.com/kubernetes/community/pull/2132)** for the KEP - commentary important
            *   **final committed KEP - **[https://github.com/kubernetes/community/blob/master/keps/sig-cli/0008-kustomize.md](https://github.com/kubernetes/community/blob/master/keps/sig-cli/0008-kustomize.md)
            *   **<span style="text-decoration:underline;">actual resulting repo: [github.com/kubernetes-sigs/kustomize](https://github.com/kubernetes-sigs/kustomize)</span>**
        *   KEP Life cycle
            *   We have a GitHub project that helps keep track of Kep project lifecycles. See it here [https://github.com/kubernetes/community/projects](https://github.com/kubernetes/community/projects)
*   [ 0:00 ] **SIG Updates**
    *   **Multicluster **- Quinton Hoole (confirmed)
        *   Sig Intro
            *   Focused on solving challenges with running multiple clusters and applications therein.
            *   Working on Cluster Federation, Cluster Registry(cluster registry for k8s for cluster reuse) and Multi cluster ingress.
        *   FederationStatus
            *   Development has split between federation v1 and v2.
            *   Federation v1 is a POC and no further development planned, users showed they needed something different.
            *   Moving forward Federation v2 will focus on reusable components, federation specific apis and implementations of higher level apis and federation controllers.
            *   v2 Alpha is planned for June.
            *   Behind the effort is RedHat and Huawei.
        *   Cluster Registry Status
            *   Grew out of Federation v1. Allows reusable clusters and discovery. Google Cloud is supported for now, but more coming. Implementation is based on CRDS.
            *   Apis/CRDS in beta.
        *   [Link to slides](https://docs.google.com/presentation/d/1mdIgFkSr7dxsoTcDZCaW0nMVLLHwn5sEOyFzBhzUgD4/edit?usp=sharing)
    *   Network - Tim Hockin - (confirmed) (or dc
        *   Sig Intro
        *   In-progress Network Plumbing CRD Spec doc:
            *   A CRD is being proposed. Reference implementation is in the works. There is a proposal that covers all the relevant information.
            *   [https://docs.google.com/document/d/1Ny03h6IDVy_e_vmElOqR7UdTPAG_RNydhVE1Kx54kFQ/edit#](https://docs.google.com/document/d/1Ny03h6IDVy_e_vmElOqR7UdTPAG_RNydhVE1Kx54kFQ/edit#)
        *   Network Service Mesh proposal slides
            *   [https://docs.google.com/presentation/d/1vmN5EevNccel6Wt8KgmkXhAfnjIli4IbjskezQjyfUE/edit#slide=id.p](https://docs.google.com/presentation/d/1vmN5EevNccel6Wt8KgmkXhAfnjIli4IbjskezQjyfUE/edit#slide=id.p)
        *   DevicePlugins (from Resource Management WG) have some intersection with networking, there have been many demos/PoCs but so far no consensus on how DPs should interact with existing CRI networking APIs
        *   CoreDNS is now GA in 1.11
        *   IPVS Proxy mode is now GA in 1.11 (anyone have a link?) but not default
        *   Looking at breaking out ingress into a bunch of individual route resources instead of one monolithic list.
        *   IPv6 discussions around how to support dual-stack are ongoing
        *   We are working on test flakes, we don't have a fix yet but HELP WANTED
    *   **VMware** - Steve Wong (confirmed)
        *   Vmware Cloud Provider
            *   The target is 1.12.
            *   Working through some process level things. This project is retained as a SubProject.
            *   Creating a working group to handle testing
        *   [Link to deck](https://docs.google.com/presentation/d/1GUrqhEpVkMb4ypCcoXs3WZGkRtYylXCNRiLSAmhc-zs/edit?usp=sharing), 4 slides, estimated 5 min:
*   [ 0:00 ] **Announcements**
    *   **Happy birthday, Kubernetes!**
    *   **Shoutouts -** _powered by slack #shoutouts _- if you see someone doing great work give them a shoutout in the slack channel so we mention those here!
        *    "@jrondeau for working on the weekend to get 1.11 doc builds working again!!" -mistyhacks
        *   "@andrewsykim for all the effort in getting SIG Cloud Provider off the ground!" -fabio
        *   "@neolit123 for really stepping up lately to help with user facing issues for the kubeadm 1.11 release. we really appreciate your contributions to the sig" -stealthybox
        *   "@cblecker who is everywhere keeping tabs on things and people on track." -gsaenger
    *   **Help Wanted**
        *   [Stephen Augustus]** **[1.12 release team is forming](https://github.com/kubernetes/sig-release/issues/167), see #sig-release for more info.  Roles & Responsibilities info [here](https://github.com/kubernetes/sig-release/#kubernetes-release-team-roles).  Volunteers needed!
        *   Help wanted on Sig Network Test Flakes reach out to #sig-network on slack
        *   Anyone Interested in learning prow and helping with the transition from. Munger to prow will be helpful. See @jberkus


## May 31, 2018 - ([recording](https://youtu.be/9RSY7czYRCY))



*   **Moderators**:  Jorge Castro [SIG Contributor Experience]
*   **Note Taker**: First Last [Company/SIG]
*   [ 0:00 ]**  Demo **-- [Aptomi](https://github.com/Aptomi/aptomi/) - application delivery engine for K8S [Roman Alekseenkov]
    *   framework on top of helm charts, for composition into services
    *   showed charts (hdfs, kafka, spark, zookeper), that together show twitter status
    *   Link to slides: [https://docs.google.com/presentation/d/1HQQ_hScOyfIt8SAYPRu6fUuLJCv7b7e6bUyosd38ir8/edit?usp=sharing](https://docs.google.com/presentation/d/1HQQ_hScOyfIt8SAYPRu6fUuLJCv7b7e6bUyosd38ir8/edit?usp=sharing)
    *   Link to repositories: [https://github.com/Aptomi/aptomi/](https://github.com/Aptomi/aptomi/)
    *   <span style="text-decoration:underline;">Full Demo (40 min): [https://www.youtube.com/watch?v=GVB3kKocKi4](https://www.youtube.com/watch?v=GVB3kKocKi4)</span>
    *   [Description & Blog Post](https://superuser.openstack.org/articles/aptomi-application-delivery-engine-k8s/)
*   [ 0:13 ]** Release Updates**
    *   1.11  [Josh Berkus - Release Lead]
        *   **_Next Deadline: Draft doc PRs due June 4th._**
        *   Currently in Code Slush.  Requiring milestones, sorry for lack of warning on that.
            *   Were not able to move to Prow milestone maintainer or Tide for this release.
        *   Code Freeze Starts Tuesday, June 5th
        *   If your feature won't be ready, now is the time to update your issue in the Features repo.
            *   [Feature tracking spreadsheet](https://docs.google.com/spreadsheets/d/16N9KSlxWwxUA2gV6jvuW9N8tPRHzNhu1-RYY4Y0RZLs/edit#gid=2053885135) has been reformatted with lots of new information.
        *   CI Signal -
            *   Almost green, last few fixes merged.
            *   1 open tracking issue - [Scale Density test for 30 pods](https://github.com/kubernetes/kubernetes/issues/63030)
            *   Conformance tests results (GCE and OpenStack) now in Release blocking dashboard
            *   @misty on slack for release docs issues
    *   Patch Release Updates
        *   x.x
        *   y.x
*   [ 0:00 ] **Introduction to KEPs** [Kubernetes Enhancement Proposals] [Caleb Miles]
    *   We'll be highlighting KEPs in community meetings
    *   tracking how decisions are made: identify the problem + find a sig for motivation agreement + documenting it for everyone
    *   [Slides](https://docs.google.com/a/google.com/presentation/d/e/2PACX-1vQ0KX1TuXC9VeXPRZhxZxNILoFzL7oEpLO1szMGCYCThxTstpK7VH7s_EJ4axseJVkJ6kkYDvhFJmsC/pub?start=false&loop=false&delayms=3000)
*   [ 0:00 ] **SIG Updates**
    *   SIG OpenStack  [David Lyle and Chris Hoge]
        *   [https://docs.google.com/presentation/d/1BGdbMQnSzrYOTLxW8VZswSwMwQ2HPlegVJszITYFv6c/edit?usp=sharing](https://docs.google.com/presentation/d/1BGdbMQnSzrYOTLxW8VZswSwMwQ2HPlegVJszITYFv6c/edit?usp=sharing)
        *   expanded testing of provided code
        *   driver testing added
        *   whitepaper
    *   SIG Node [Dawn Chen]
        *   Made a steady progress on all 5 areas in Q2: 1) node management including Windows, 2) application / workload management, 3) security, 4) resource management and 5) monitoring, logging and debuggability.
        *   On node management
            *   Promoted dynamic kubelet config to beta
            *   Refactor the system to use node-level checkpointing manager
            *   Proposed a probe-based mechanism for kubelet plugins: device, csi, etc.
            *   Proposed a design to address the scalability issue caused by large node object and approved by the community. Had a short-term workaround in v.11, and plan to work on the long term solution in v1.12.
        *   Together with sig-windows, we made many progress on Windows support which including stats, node e2e for Windows Container Image. More works on SecurityContext, storage and network in next release.
        *   Both CRI-O and containerd are GA in this release.
            *   More enhancements on CRI for container logs
            *   Many enhancements to crictl, the tool for all CRI-compliant runtimes. Expecting to be GA in v1.12
            *   Announced CRI testing policy to the community, and introduced node exclusive tags to e2e.
        *   On security
            *   For 1.11, making all addons use default seccomp profile. Expecting to promote it to beta and enable it by default.
            *   Proposed a design and alpha-level Kubernetes API for sandbox. Working closely with Kata community and gVisor community on integration of CRI-compliant runtime.
            *   WIP for user namespace support
            *   Made progress on node TLS bootstrap via TPM.
        *   On resource management side, we made the progress on promoting sysctl to beta and proposed ResourceClass to make resource support extensible.
        *   Made the steady progress on debug pod, but unfortunately due to backforth review from the different reviewers on API changes, we couldn't have alpha support in v1.11. Escalate it to sig-architecture.
        *   On the logistics side
            *   Sig-node holds weekly meeting on Tuesday, 10am (Pacific Time)
            *   Please join kubernetes-sig-node googlegroup to have access to all design docs, roadmap and emails.
            *   Derek and I are working on sig-node charter, which is still under review and discussion.
*   [ 0:00 ] **Announcements**
    *   [Deprecation Policy Update](https://groups.google.com/forum/#!topic/kubernetes-dev/pNcskHXAD-k) (Important!)
    *   SIG Leads - check the top of this document for a link to the SIG Update schedule.
    *   Shoutouts - Someone going above and beyond? Mention them in #shoutouts on slack to thank them.
        *   Aish Sundar - Shoutout to @dims and OpenStack team for quickly getting their 1.11 Conformance results piped to CI runs and contributing results to Conformance dashboard!
        *   Aish Sundar - Shoutout to Benjamin Elder for adding Conformance test results to all Sig-release dashboards - master-blocking and all release branches.
        *   Josh Berkus and Stephen Augustus -  To Misty Stanley-Jones for aggressively and doggedly pursuing 1.11 documentation deadlines, which both gives folks earlier warning about docs needs and lets us bounce incomplete features earlier
    *   Help Wanted
        *   Looking for Mandarin-speakers to help with new contributor workshop and other events at KubeCon/CloudNativeCon Shanghai.  If you can help, please contact @jberkus / [jberkus@redhat.com](mailto:jberkus@redhat.com)
        *   [KEP-005](https://github.com/kubernetes/community/blob/master/keps/sig-contributor-experience/0005-contributor-site.md) - Contributor Site - ping [jorge@heptio.com](mailto:jorge@heptio.com) if you can help!
    *   Meet Our Contributors (mentors on demand)
        *   June 6th at 230p and 8pm **UTC** [https://git.k8s.io/community/mentoring/meet-our-contributors.md](https://git.k8s.io/community/mentoring/meet-our-contributors.md)
        *   Want to know the paths of some of our approvers? Confused about what a SIG is? Anything that you'd ask a mentor - ask in #meet-our-contributors on slack or DM @paris with an anonymous question
    *   [Stackoverflow Top Users](https://stackoverflow.com/tags/kubernetes/topusers)
        *   [Const](https://stackoverflow.com/users/9663586/const)
        *   [VAS](https://stackoverflow.com/users/9521610/vas)
        *   [Alexandr Lurye](https://stackoverflow.com/users/9611623/alexandr-lurye)
        *   [James Strachan](https://stackoverflow.com/users/2068211/james-strachan)
        *   [Jordan Liggitt](https://stackoverflow.com/users/54696/jordan-liggitt)
    *   Thread o' the week: [How has Kubernetes failed for you?](https://discuss.kubernetes.io/t/how-has-kubernetes-failed-for-you/481)


## May 24, 2018 - ([recording](https://youtu.be/zKtxTbq0s4o))



*   **Moderators**:  Josh Berkus  [SIG-Release]
*   **Note Taker**: Tim Pepper [VMware/SIGs Release & ContribX]
*   [ 0:00 ]**  Demo **--  Workflows as CRD [ Jesse Suen (Jesse_Suen@intuit.com)]
    *   Link to slides: [https://drive.google.com/file/d/1Z5TMIr6r4hC7N5KeVqajC3c3NcYqK4_z/view?usp=sharing](https://drive.google.com/file/d/1Z5TMIr6r4hC7N5KeVqajC3c3NcYqK4_z/view?usp=sharing)
    *   Link to repositories: [https://github.com/argoproj/argo](https://github.com/argoproj/argo)
    *   Argo: a fancy job controller for workflows, DAGs implemented as CRD.  Originally intended for CI/CD pipelines, but is seeing usage for other workflows like machine learning.
    *   Used with kubeflow
    *   Component architecture interfacing wsith k8s api server and leveraging sidecars in pods for workload artifact management
    *   Argo command line gives validation of commands, but is effectively a kubectl wrapper
    *   Workflows can be defined as a top down iterative list of steps, or as a DAG of dependencies
*   [ 0:16 ]** Release Updates**
    *   1.11 Update [Josh Berkus, Release Lead]
        *   **_Next Deadline: Docs Placeholder PRs Due Tomorrow for [feature list](https://docs.google.com/spreadsheets/d/16N9KSlxWwxUA2gV6jvuW9N8tPRHzNhu1-RYY4Y0RZLs/edit#gid=0)_**!!!
        *   **_Code Slush on Tuesday_**
        *   Current CI status and schedule
            *   CI Signal : [Tracking 3 open issues](https://docs.google.com/spreadsheets/d/1j2K8cxraSp8jZR2S-kJUT6GNjtXYU9hocNRiVUGZWvc/edit#gid=127492362), all are test issue being actively worked on.
            *   Code freeze coming June 5, make sure your issues/PRs are up to date with labels and priorities and status
        *   [Changing Burndown Meeting Schedule](https://github.com/kubernetes/sig-release/issues/148), please comment..looking for less conflicted times and ones friendlier for more timezones
    *   Patch Release Updates?
        *   1.10.3 released monday
*   [ 0:21 ] **SIG Updates**
    *   SIG Service Catalog [Doug Davis] (confirmed)
        *   Beta as of Oct 2017
        *   Key development activities
            *   New svcat cmd line tool (similar to CloudFoundry's way of things)
            *   NS-scoped brokers - still under dev
            *   Considering moving to CRDs instead of dedicated apiserver
        *   Finalizing our v1.0 wish-list
            *   NS-scoped brokers
            *   Async-bindings
            *   Resolve CRD decision
            *   Generic Broker & Instance Actions
            *   GUIDs as Kube "name" is problematic
        *   SIG has recently been actively mentoring and onboarding newcomers
    *   SIG Auth [Tim Allclair](confirmed)
        *   Pod TokenRequest API and ServiceAccountTokenProjection improving for 1.11
        *   Client-go gaining support for x509 credentials and externalizing currently in-tree credential providers
        *   Scheduling policy design thinking happening ahead of 1.12
        *   Audit Logging: improved annotation metadata coming around auth and admission for logs
        *   Node Isolation:  nodes no longer able to update their own taints (eg: exploit to attract sensitive pod/data to a compromised node)
        *   Conformance: [KEP PR open on security related conformance](https://github.com/kubernetes/community/pull/2081) testing to give better assurance that best practices are in use or validate a hardened profile is active.  Likely not 1.11 rather 1.12.
        *   Bug bounty is WIP
    *   SIG Storage [Brad Childs](confirmed) [Slides](https://docs.google.com/presentation/d/1HkCHC5xkxt2TXOLS1riXUeIRRejtX-82tOf5x42F4w8/edit?usp=sharing)
        *   Had SIG face-to-face meeting last week. ~40 people and 19 companies present
        *   Storage functionality is moving out of tree by way of the CSI interface
            *   CSI spec moving from 0.2 to 0.3 soon
            *   Lots of CSI related features coming in k8s 1.11
            *   Aiming for out-of-tree feature parity relative to existing in-tree
        *   Feature areas: Snapshots, topology aware (scheduling relative to location of PV) and local PV, local disk pools/health/capacity, volume expansion and online resize
        *   Testing: multi-phased plan to inventory and improve test coverage and CI/CD, including test coverage on other cloud providers beyond GCE.  VMware committed resources, looking for commit from others.
        *   Operators: external provisioners, snapshot and other operator frameworks underway.  Currently not looking to do a shared operator library to span SIG-Storage repos.
        *   Metrics: there are a lot.  Some are cloud provider specific.  Goal is to assist SRE's in problem determination and corrective action.
        *   API Throttling: api quota exhaustion at cloud provider and api server are frequently causing storage issues.  Looking at ways to streamline.
        *   External projects:  SIG has something like 20 projects and is breaking them apart, looking for owners and out of tree locations for them to better live.  Projects should move to CSI, a kubernetes-sigs/* repo, a utility library, or EOL
*   [ 0:00 ] **Announcements**
    *   <span style="text-decoration:underline;">Shoutouts this week</span> (Check in #shoutouts on slack)
        *   Big shoutout to @carolynvs for being welcoming and encouraging to newcomers, to @paris for all the community energy and dedication, and to all the panelists from the recent KubeCon/CloudNativeCon diversity lunch for sharing their experiences.
        *   Big shoutout to @mike.splain for running the Boston Kubernetes meetup (9 so far!)
        *   everyone at svcat is awesome and patient especially @carolynvs, @Jeremy Rickard & @jpeeler who all took time to help me when I hit some bumps on my first PR.
    *   <span style="text-decoration:underline;">Help Wanted</span>
        *   SIG UI is looking for new contributors. Check out their issue log to jump in; also listen to their SIG UI call today where they explained more and answered questions. #sig-ui in slack for on-ramp help. [Notes from the call ](https://docs.google.com/document/d/1PwHFvqiShLIq8ZpoXvE3dSUnOv1ts5BTtZ7aATuKd-E/edit?usp=sharing)
        *   Looking for more mentors as we kick off our contributor mentoring programs. [Fill out this form ](https://goo.gl/forms/uKbzNsv51JUVkC0g1)(works for looking for mentorship, too). Pardon the dust as we do a mentor recruiting drive.


## May 17, 2018 - ([recording](https://youtu.be/DpFTcTnBxbM))



*   **Moderators**: Paris Pittman [SIG Contributor Experience]
*   **Note Taker**: Solly Ross
*   **Demo:** [Gardener Demo](https://gardener.cloud) ([vasu.chandrasekhara@sap.com](mailto:vasu.chandrasekhara@sap.com) and [rafael.franzke@sap.com](mailto:rafael.franzke@sap.com))
    *   [_occurred towards end of video instead_]
    *   Open Source: [https://gardener.cloud/](https://gardener.cloud/)
    *   Mission
        *   Manage, maintain, operate multiple k8s clusters
        *   Work across public and private clouds
    *   Architecture
        *   Self-hosted
        *   Kube-centric
        *   Steps
            *   Boot initial "garden" cluster using kubify ([https://github.com/gardener/kubify](https://github.com/gardener/kubify), open source)
            *   Deploy Gardener to "garden" cluster + dashboard (Gardener is extension API server)
            *   Run/use "seed" cluster to run control plane components, terraformer for each cluster "shoot" cluster (1 seed per hosting platform, region, etc)
            *   Each set of control plane components corresponds to a "shoot" cluster with actual nodes (machine controller + machine API objects control this)
            *   VPN between "seed" cluster and "shoot" clusters so that API server, monitoring can talk to node
    *   Secrets are created for each shoot to easily download kubeconfigs, etc
    *   Declarative config for each cluster ("shoot") with status info as well
    *   Uses cluster API machine resources, working with Cluster API WG
    *   Q: is it stable, or in development?
        *   A: used internally, but still in development
    *   Q: baremetal support?
        *   If there's an infra API that can be used to control baremetal, then that can be used
    *   Detailed Blog describing Gardener's architecture: [https://kubernetes.io/blog/2018/05/17/gardener/](https://kubernetes.io/blog/2018/05/17/gardener/)
*   **Release Updates:**
    *   **1.11 **[Josh Berkus, RT Lead / Aish Sundar CI Signal Lead] (Week 7)
        *   **_Next Deadline: Docs, Open Placeholder PRs Required, May 25th_**
        *   1.11.0 Beta0 released yesterday.
        *   We are delaying/shortening Code Freeze as discussed.  See[ new calendar](https://github.com/kubernetes/sig-release/blob/master/releases/release-1.11/release-1.11.md) for current deadlines.
            *   Stable passing tests, low bug count → small code freeze periods → more development time
            *   Code slush: **May 29th**
            *   Code freeze: **June 5th**
        *   Many thanks to dims, liggit, timothysc, krousey, kow3ns, yliaog, k82cn, mrhohn, msau42, shyamvs, directxman12 for debugging fails and closing issues, and AishSundar, Cole Mickens, Mohammed Ahmed, and Zach Arnold for working with the SIGs to get attention on issues and test failures.
        *   Help wanted to on scalability and performance
    *   **1.10 **[Maciek Pytel, PRM]
*   **SIG Updates:**
    *   [Scheduling](https://github.com/kubernetes/community/blob/master/sig-scheduling/README.md)** **[Bobby Salamat]
        *   Priority and Preemption
            *   Have gotten good feedback over the past quarter
            *   Moving to beta/enabled by default in 1.11
        *   Equivalence Cache Scheduling
            *   Caching predicate results for given inputs as long as conditions don't change in cluster
        *   Gang Scheduling
            *   Schedule a bunch of pods together, don't schedule only a subset
            *   Kube-arbitrator has a prototype that seems to work well
            *   Need to collect more requirements
            *   Q: Can we use batch scheduling to improve throughput?
                *   A: Maybe use a Firmament-like approach?
                *   Q: is this a step along the way for perf optimization on the current schedule?
                *   A: Engineers from Huawei are working on this, but ran into issues with things like pod-antiafinity, actually binding the pods
        *   Taint based eviction to beta
        *   Scheduling framework
            *   Still in design framework
        *   Pod scheduling policy
            *   Lots of opinions, progress has been slow
            *   Existing design proposal with lots of opinions
    *   [Scalability](https://github.com/kubernetes/community/blob/master/sig-scalability/README.md) [Bob Wise]
        *   [Slides](https://docs.google.com/presentation/d/1vP3kRPiei5yNrNmsrndWQgZo2idcc8uN0LgjK4xt2wU/edit)
        *   Schedule for large runs of perf is even-odd day
        *   Different Per Axes (there's not just one axis, e.g. "number of nodes")
            *   Nodes, Pod Churn, Pod density, Networking, Secrets, Active Namespaces
        *   Pro Tips
            *   Lock your etcd version
            *   Test your cluster with Kubemark
        *   Recommended reading in slides
            *   Perf regression study
            *   Scalability good practices
        *   WIP Items
            *   Better testing of real workloads (cluster-loader)
            *   More scalability testing in presubmit tests
                *   Concerns around run time issues
            *   Sonobouy per testing
        *   Q: Limitations on scalability come down to etcd perf, do we work with etcd engineers?
            *   A: Perf is generally not an etcd issue wrt bottlenecks
            *   A: Etcd tends to be regressions across etcd versions, not etcd as bottleneck
            *   A: Range locking issues being improved in 3.3,3.4
            *   A: talk to Shyam about this for more info
    *   [API Machinery](https://github.com/kubernetes/community/blob/master/sig-api-machinery/README.md) [Daniel Smith, confirmed]
        *   New Dynamic client with better interface!
            *   Old is under "deprecated" directory
            *   Clientside QPS rate limit behavior changed
        *   CRD Versioning
            *   Design issue with versioning priorities found, but no-op conversion will still land in 1.11
        *   Apply WG
            *   Feature branch for apply, trying to put things in master when possible
            *   Won't reintegrate before 1.11 (feature branch work will continue through code freeze)
*   **Announcements:**
    *   **Shoutouts!**
        *   Warm welcome to @liz and @cha for their journeys in joining the k8s org! Both of you have been having a big impact in #sig-cluster-lifecycle - stealthybox
        *   @chancez and @danderson for a great conversation on bare metal options and concerns! - mauilion
        *   shoutout to @liggitt, master wrangler of e2e test bugs. Jordan has fixed many ["fun" bugs](https://github.com/kubernetes/kubernetes/issues/63731#issuecomment-388529120). Thanks for helping keep things green! :smile: - bentheelder
        *   As a new contributor, I can 100% endorse @carolynvs for being REALLY GOOD at bringing in new contributors, and dedicating a lot of time and effort to make sure they are successful. -teague_cole
    *   **Help Wanted!**
        *   SIG UI looking for new contributors to go up the ladder to maintainers. Start with an open issue and reach out to the mailing list and slack channel.
        *   SIG Scalability is looking for contributors!
        *   We need more contributor mentors! [Fill this out.](https://goo.gl/forms/17Fzwdm5V2TVWiwy2)
            *   The next Meet Our Contributors (mentors on demand!) will be on June 6th. Check out kubernetes.io/community for time slots and to copy to your calendar.
    *   **KubeCon/CloudNativeCon Follow Ups**
        *   Videos and slides: [https://github.com/cloudyuga/kubecon18-eu](https://github.com/cloudyuga/kubecon18-eu) Thanks CloudYuga for this!
    *   **Other**
        *   Don't forget to check out [discuss.kubernetes.io](https://discuss.kubernetes.io/)!
        *   DockerCon Kubernetes Contributor AMA during Community Day - June 13th. 3 hour window; specific time TBA


## May 10, 2018 - ([recording](https://youtu.be/ygW6jTBp7Fs))



*   **Moderators**: Tim Pepper [SIG Contributor Experience, SIG Release]
*   **Note Taker**: Jorge Castro / Christian Roy
*   **Demo:** Ambassador API Gateway built on Envoy/K8S ([https://www.getambassador.io](https://www.getambassador.io))  ([richard@datawire.io](mailto:richard@datawire.io))
    *   [https://github.com/datawire/ambassador](https://github.com/datawire/ambassador)
    *   Link to [slides](https://www.slideshare.net/datawire/ambassador-kubernetesnative-api-gateway)
    *   Kubernetes only, simple architecture
        *   Apache licensed
        *   Declarative configuration via kubernetes annotations
        *   Built on Envoy - designed for machine configuration
        *   Operates as a sidecar to envoy, async notified of config changes and configures envoy accordingly
        *   Concept of shadowing traffic - takes all the incoming requests and sends it to another service but filters the responses, good for debugging in production.
*   **Release Updates:**
    *   **1.11 **[Josh Berkus, RT Lead / Aish Sundar CI Signal Lead] (Week 6)
        *   _Next Deadline: Beta0, May 15th. Tests must be passing by then!_
        *   _Week started well but gke tests started failing_
        *   _Some top level failing tests generate failures across other groups_
        *   _SIGs are responding responsibly on the failures_
        *   [43 Tracking Features](http://bit.ly/k8s111-features)
        *   CI Signal Test report - [https://docs.google.com/spreadsheets/d/1j2K8cxraSp8jZR2S-kJUT6GNjtXYU9hocNRiVUGZWvc/edit#gid=127492362](https://docs.google.com/spreadsheets/d/1j2K8cxraSp8jZR2S-kJUT6GNjtXYU9hocNRiVUGZWvc/edit#gid=127492362)
    *   **1.10** [Maciek Pytel, PRM]
        *   1.10.3 release planned Monday May 21st
*   **SIG Updates:**
    *   **Architecture [Brian Grant]**
        *   Working on our [charter](https://github.com/kubernetes/community/pull/2074)
            *   Improving conformance tests
            *   Provide technical expertise/advice/overview across SIGs
            *   Formalizing proposal processes into KEPs, more structure, make it more obvious
            *   API review process. Used to be informal, we want to formalize that.
        *   Weekly meeting with alternating full meeting (decisions) and office hours (discussions)
            *   Office hours are available for people who want to ask questions on how to best implement incoming ideas (API review, etc.)
        *   [Meeting and note information](https://github.com/kubernetes/community/blob/master/sig-architecture/README.md)
    *   **Contributor Experience [Paris Pittman]**
        *   [Contributor site](https://github.com/kubernetes/community/blob/master/keps/sig-contributor-experience/0005-contributor-site.md) KEP underway
        *   Discourse is up and ready to test!
            *   [Discuss.kubernetes.io](http://discuss.kubernetes.io/) - please post content, announcements, meetup reminders, or just [introduce yourself](https://discuss.kubernetes.io/t/introduce-yourself-here/56)!
        *   [Looking for mentors!](https://goo.gl/forms/3ISrNbTkYqExWzKw1)
            *   Register for [Meet Our Contributors monthly YT series](https://youtu.be/EVsXi3Zhlo0) (first Weds of the month; link on kubernetes.io/community) with this form and all other mentoring activities
        *   Contributor Experience survey to go out in June
            *   Communication platform
            *   Flow in github
        *   [Developers Guide underway](https://github.com/kubernetes/community/issues/1919) under Contributor Docs subproject
        *   Contributor Experience Update [slide deck](https://docs.google.com/presentation/d/1KUbnP_Bl7ulLJ1evo-X_TdXhlvQWUyru4GuZm51YfjY/edit?usp=sharing) from KubeCon/CloudNativeCon UE [if you are in k-dev mailing list, you'll have access)
*   **Announcements:**
    *   **Shoutouts!**
        *   See someone doing something great in the community? Mention them in #shoutouts on slack and we'll mention them during the community meeting:
        *   Ihor Dvoretskyi thanks @justaugustus, who made a GREAT job as a Kubernetes 1.11 release features shadow
        *   Josh Berkus to Aish Sundar for doing a truly phenomenal job as CI signal lead on the 1.11 release team
        *   Tim Pepper to Aaron Crickenberger for being such a great leader on the project during recent months
        *   Chuck Ha shouts out to the doc team - "Working on the website is such a good experience now that it's on hugo. Page rebuild time went from ~20 seconds to 60ms" :heart emoji:
        *   Jason de Tiber would like to thank Leigh Capili (@stealthybox) for the hard work and long hours helping to fix kubeadm upgrade issues. (2nd shoutout in a row for Leigh! -ed)
        *   Jorge Castro and Paris Pittman would like to thank Vanessa Heric and the rest of the CNCF/Linux Foundation personnel that helped us pull off another great Contributor Summit and KubeCon/CloudNativeCon
        *   [Top Stackoverflow Users](https://stackoverflow.com/tags/kubernetes/topusers) in the Kubernetes Tag for the month
            *   Anton Kostenko, Nicola Ben, Maruf Tuhin, Jonah Benton, Const
    *   Message from the docs team re: hugo transition:
        *   We've successfully migrated kubernetes.io from a Jekyll site framework to Hugo. Any open pull requests for k/website need to be revised to incorporate the repo's new content structure. (Changes in `docs/` must now change `content/en/docs/`.)
        *   More about the framework change: [https://kubernetes.io/blog/2018/05/05/hugo-migration/](https://kubernetes.io/blog/2018/05/05/hugo-migration/)
    *   KEP Section for the Community Meeting? [Jorge Castro]
        *   Lots of KEPs coming in via PR, should we have current KEPs in flight as a standing agenda item in the community meeting?  
        *   When starting a KEP, send an email FYI to the appropriate SIGs and Arch as github notifications are noisy and missed.
        *   Would be good to help us bootstrap the KEP processes for people if we got some visibility on them, but still need a site of KEPs
    *   Kubernetes Application Survey results [Matt Farina] WG
        *   [Raw results](https://docs.google.com/spreadsheets/d/12ilRCly2eHKPuicv1P_BD6z__PXAqpiaR-tDYe2eudE/edit)
        *   [Deck on results](https://docs.google.com/presentation/d/1utT0K-u1nl2apXRo29GaBvRV1x7mFLeQSgpw8mI_nGM/edit?usp=sharing) (Slides)
        *   [Blog post about it](https://kubernetes.io/blog/2018/04/24/kubernetes-application-survey-results-2018/)
        *   Developers, check it out, people took a lot of time to give us lots of good information, take the time to get information from it.

        **Help Wanted?**

        *   [SIG UI](https://github.com/kubernetes/community/blob/master/sig-ui/README.md) is looking for additional contributors (with javascript and/or go knowledge) and maintainers
            *   [Piotr](https://github.com/bryk) and [Konrad](https://github.com/konryd) from google have offered to bring folks up to speed.
            *   Take a look at open issues to get started or reach out to their slack channel, mailing list, or next meeting.
            *   SIG UI mailing list: [https://groups.google.com/forum/#!forum/kubernetes-sig-ui](https://groups.google.com/forum/#!forum/kubernetes-sig-ui)


## April 26, 2018 - (recording)



*   **Moderators**: Jorge Castro [SIG Contributor Experience]
*   **Note Taker**: Christian Roy
*   **Demo:** [Gitkustabe](https://github.com/hasura/gitkube): Build and deploy docker images to Kubernetes using git push (<span style="text-decoration:underline;">shahidh@hasura.io</span>, [tiru@hasura.io](mailto:tiru@hasura.io))
    *   [https://github.com/hasura/gitkube](https://github.com/hasura/gitkube)
    *   Git push to a url in your k8s cluster
    *   Builds and deploy the image in an deployment
*   **Release Updates:**
    *   **1.11 **[Josh Berkus, RT Lead]
        *   Features are now frozen!
        *   [44 filed features](https://github.com/kubernetes/features/issues?utf8=%E2%9C%93&q=is%3Aissue+is%3Aopen+milestone%3Av1.11)
        *   [Test Status](http://bit.ly/k8s111-cisignal) is **Red**
            *   Master-upgrade: 11 out of 12 jobs failing
            *   Master-blocking: 5 out of 26 jobs failing
            *   [Issues filed](https://docs.google.com/spreadsheets/d/1j2K8cxraSp8jZR2S-kJUT6GNjtXYU9hocNRiVUGZWvc/edit#gid=2128913655)
                *   Some SIGs not responding to issues
                *   SIG responsible to debug why test failing, please look at your issues and start to prioritize them!
        *   Next Deadline: [Beta Release May 15th](https://github.com/kubernetes/sig-release/blob/master/releases/release-1.11/release-1.11.md)
            *   Will currently be blocked by failing tests.
    *   1.10 [Maciek Pytel, PRM]
        *   Branch frozen for 1.10.2, release planned for April 27th or Monday April 30th
*   **Graph of the Week: **Slack Stats [Paris Pittman]
    *   [Slack Stats](https://docs.google.com/document/d/1ZR8ZqYcql_C8nGFrKqcZw_ZANkEb4P-3CNFyT6OkfDc/edit)
        *   35k users with 5k weekly active users
        *   Produced Quarterly
*   **SIG Updates:**
    *   **Thanks to test infra folks for labels**
    *   **Cluster Lifecycle [Tim St. Clair]**
        *   Kubeadm
            *   Steadily burning down against 1.11
            *   Found+Fixed some thorny upgrade issues in 1.10
                *   Tests are still broken
            *   Working on proposal
                *   HA
                    *   Master join
                *   UX
                    *   Phases rework
                *   Upgrade
                    *   Config changes
                    *   Self hosting
        *   ClusterAPI [kris nova]
            *   There is a new repo: https://github.com/kubernetes-sigs/cluster-api
            *   Aiming to keep cloud provider logic OUT of the repo (common logic only)
            *   Aiming for a (api only) alpha release 1.11
                *   Configurable machine setup proposal in progress
                *   [https://docs.google.com/document/d/1OfykBDOXP_t6QEtiYBA-Ax7nSpqohFofyX-wOxrQrnw/edit?ts=5ae0b27a#heading=h.xgjl2srtytjt](https://docs.google.com/document/d/1OfykBDOXP_t6QEtiYBA-Ax7nSpqohFofyX-wOxrQrnw/edit?ts=5ae0b27a#heading=h.xgjl2srtytjt)
    *   **Autoscaling [Solly Ross]**
        *   HPA v2 improvements ([https://github.com/kubernetes/community/pull/2055](https://github.com/kubernetes/community/pull/2055))
            *   Label selectors for metrics
            *   Support for averages on object metrics
            *   Slight changes to structure of object (Unify metrics sources)
        *   Better e2e tests on all HPA functionality
            *   Movement along the path to blocking HPA custom metrics e2e tests
        *   VPA work coming along, alpha soon (demo at KubeCon/CloudNativeCon)
        *   Come say hi at KubeCon/CloudNativeCon (Intro and Deep Dive, talks on HPA)
    *   **PM [Jaice Singer DuMars]**
        *   Working on mechanisms to get feedback from the user community (playing with something like [http://kubernetes.report](http://kubernetes.report) -- in development, not ready for distro yet)
        *   Presenting at KubeCon/CloudNativeCon 16:35 on Thursday ~ Ihor and Aparna
        *   Working on a charter draft
            *   We actually represent three 'P' areas: product, project, and program
            *   Help SIG focus on implementations
            *   We're trying to look a
*   **Announcements:**
    *   **KubeCon/CloudNativeCon next week, no community meeting! **\o/
    *   **Last Chance to Register for the Contributor Summit - **
        *   Registration ends Fri, Apr 7th @ 7pm UTC
        *   Tuesday, May 1, day before KubeCon/CloudNativeCon
        *   You must [register here](https://github.com/kubernetes/community/tree/master/events/2018/05-contributor-summit) even if you've registered for KubeCon/CloudNativeCon
        *   SIGs, remember to [put yourself down on the SIG Update sheet](https://docs.google.com/spreadsheets/d/1adztrJ05mQ_cjatYSnvyiy85KjuI6-GuXsRsP-T2R3k/edit#gid=1543199895) to give your 5 minute update that afternoon.
    *   **Shoutouts!**
        *   See someone doing something great in the community? Mention them in #shoutouts on slack and we'll mention them during the community meeting:
        *   Timothy St. Clair would like to thank Peter Zhao (@xiangpengzhao) for "steadfast PR-review and contributions to SIG Cluster Lifecycle"
        *   Chuck Ha would like to thank Leigh Capilo (@stealthybox) for "being welcoming to new folks in SIG Cluster Lifecycle. You are welcoming and helpful and it keeps our community healthy. Thank You!"
    *   [Normalization of Kind Labels](https://github.com/kubernetes/community/issues/2032) [Josh Berkus]
        *   Updating list of kind labels, how they are used
    *   Cloud Foundry wants to welcome folks to their Summit [on May 1st](http://sched.co/Dun0) (day before KubeCon/CloudNativeCon)
    *   You can now use Kubernetes to [water your lawn](https://www.youtube.com/watch?v=Y5WDO-OTf-4).
    *   No Meeting next week!
    *   **Help Wanted?**
        *   **Add to this section when you have something you need help with! Issue, contributors, etc.**


## April 19, 2018 - ([recording](https://youtu.be/fEYVDMB3Xzo))



*   **Moderators**:  Paris Pittman [SIG ContribEx]
*   **Note Taker**: Jaice Singer DuMars (Google)
*   [0:01] **Demo **- [Skaffold](https://github.com/GoogleCloudPlatform/skaffold) Matt Rickard - Google  ([mrick@google.com](mailto:mrick@google.com))
    *   [https://github.com/GoogleContainerTools/skaffold](https://github.com/GoogleContainerTools/skaffold)
    *   Tool for developing applications on Kubernetes
    *   Allows you to step into CI/CD
    *   skaffold-dev / skaffold-run are the two primary components
    *   [ 0:03 ] - Demo
    *   Q: What is the plan around integration for new Kubernetes releases?
        *   Pinned to 1.10, have integration testing but not version skew
        *   Want to follow the Kubernetes support process of ~2 releases
    *   Q: Why would this not be in CNCF/part of k8s?
        *   Trying to keep it unopinionated
        *   If a community project makes sense, we will examine that
        *   MFarina: Ecosystem projects are the preference to avoid contention
    *   Q: So what are the non docker image formats this tool supports?
        *   Only supports bazel
        *   Working on java support
        *   This is the other build tool we're working on integrating next for skaffold. [https://github.com/google/jib](https://github.com/google/jib)
        *   Minimal arbitrary support, but requires a file to query and parse to determine SC dependencies, currently in-tree but might move to a plugin model
*   [0:14]** Release Updates**
    *   1.11 [Josh Berkus ~ Release Lead] (confirmed)
        *   **_Feature Freeze is Tuesday, April 24!_**
            *   File your features: [https://github.com/kubernetes/features/issues](https://github.com/kubernetes/features/issues)
        *   Tests are currently **not passing**
            *   [CI Signal Report](https://docs.google.com/document/d/1y044OcaKGEUgj094JH1ZxnnLRHnqi0Kq0f4ov56kvxE/edit?ts=5ad596c2)
            *   [Issues](https://github.com/kubernetes/kubernetes/issues?utf8=%E2%9C%93&q=is%3Aissue+is%3Aopen+milestone%3Av1.11+test)
        *   Releasinging Alpha2 today (thanks Caleb!)
        *   Release notes doc will be up later today, watch kubernetes-dev
    *   Patch Release Updates
*   [0:00]** Graph o' the Week**
    *   YouTube Channel Stats!
        *   ~7000 subscribers and growing
        *   Old videos and high engagement videos get the most attention
        *   SIG recordings are typically used as a sleep aid
        *   4:45 average view time on videos
        *   If you are posting videos, use descriptive titles, tags
        *   Desktops primary viewing device, also TVs
        *   You can turn the video speed up to 1.5 if you want to get through the material faster
*   [0:23] **SIG Updates**
    *   **CLI **(Maciej Szulik - confirmed)
        *   Printing of objects being moved to server - currently in beta, in 1.10 users were able to opt in to it
            *   You can opt out via flag, but it is on by default in 1.11
            *   No user-facing impact, but if there are, contact sig-cli
        *   Different patterns across the repo, and trying to unify by providing identical flags and output
            *   unified flag handling will unify the code base, ux, and simplify the code base
    *   **AWS **(Justin SB - Confirmed)
        *   Our first [sig repository](https://github.com/kubernetes/community/blob/master/kubernetes-repositories.md#sig-repositories): [aws-encryption-provider](https://github.com/kubernetes-sigs/aws-encryption-provider) ~ encryption at rest in etcd
        *   Justin SB is now a Googler
        *   CP breakout is blocked by non-technical issues
            *   _From Micah Hausler (EKS) to Everyone: (10:29 AM): Small correction: We are actively working on the CP breakout here at AWS (we've had a ad-hoc community-based meeting to get it going) - [meeting notes](https://docs.google.com/document/d/1-i0xQidlXnFEP9fXHWkBxqySkXwJnrGJP9OGyP2_P14/edit#heading=h.dbsrync38vdv)_
        *   Need help working on this
    *   **GCP **(Adam Worrell - confirmed) ([bit.ly/k8s-sig-gcp](http://bit.ly/k8s-sig-gcp))
        *   Not thriving, 3 meetings total but having a lack of topics
        *   There's only one lead, but someone has expressed interest
        *   Organizationally important, but there don't seem to be externally-interested parties
        *   There are lurkers, but not a specific community
        *   Community, please use this opportunity
*   [0:00] **Announcements**
    *   <span style="text-decoration:underline;">Shoutouts!</span>
        *   Join #shoutouts to add yours to the weekly announcements
        *   @maciekpytel for providing some nuance and clarity around node autoscaler
        *   @cblecker for fielding so many issues and PRs.
    *   <span style="text-decoration:underline;">Help Wanted?</span>
        *   SIG UI is looking for more active contributors to revitalize the dashboard. Please join their [communication channels](https://github.com/kubernetes/community/blob/master/sig-ui/README.md) and attend the next meeting to announce your interest.
    *   <span style="text-decoration:underline;">KubeCon/CloudNativeCon EU Update</span>
        *   Current contributor track session voting will be emailed to attendees today!C
        *   RSVP for Contributor Summit [[here]](https://github.com/kubernetes/community/tree/master/events/2018/05-contributor-summit)
        *   SIG Leads, please do your updates for the 5 minute updates
    *   CNCF meet the maintainers group is organizing ~ please sign up for attending the CNCF booth


## April 12, 2018- ([recording](https://www.youtube.com/watch?v=1wTmoXPfspI))



*   **Moderators**:  Josh Berkus [SIG-Release]
*   **Note Taker**: Clint Kitson [VMware]
*   69+ participants
*   [ 10:00 ]**  Demo **-- CRI-O [Antonio Murdaca, runcom@redhat.com] (confirmed)
    *   [https://github.com/kubernetes-incubator/cri-o](https://github.com/kubernetes-incubator/cri-o)
    *   Support for k8s 1.9/1.10 and tracking changes consistently
    *   Planned support for clearcontainers/kata-containers
    *   Demo on K8s 1.10
    *   Support for kubeadm and minikube
    *   Create issues on crio project on github
    *   sig-node does not have plans to choose one yet
        *   Working on conformance to address implementations which should lead to choosing default implementation
        *   Choice is important since it would be used under scalability testing
        *   Test data? Plan to publish results to testgrid, will supply results ASAP
            *   Previously blocked on dashboard issue
            *   Can get to point to make crio a blocking job-- multiple releases at this status before graduation
    *   Request for containerd update to group -> contribex
*   [ 10:18 ]** Release Updates**
    *   1.11 [Josh Berkus ~ Release Lead] (confirmed)
        *   Week 2 of 12 (see [schedule](https://github.com/kubernetes/sig-release/blob/master/releases/release-1.11/release-1.11.md))
            *   Currently collecting features [Ihor]
                *   Please, work on adding/updating the features in the [features](https://github.com/kubernetes/features/issues?q=is%3Aopen+is%3Aissue+milestone%3Av1.11) repo.
                *   Please, also add them to the features [tracking spreadsheet](https://docs.google.com/spreadsheets/d/16N9KSlxWwxUA2gV6jvuW9N8tPRHzNhu1-RYY4Y0RZLs/edit?ouid=103000293055760527954&usp=sheets_home&ths=true).
            *   Discussion about changing the Freeze/RC schedule
        *   Been difficult to get all tests passing, so code freezes have been lengthening
            *   Less time to work on development
            *   In 1.11 trying to shorten code freeze conditionally
                *   If last 3 week days if all test passing then will make code freeze shorter to enable more development time
                    *   All 3 dashboards (1.11-blocking, master-blocking, master-upgrade in [testgrid](https://k8s-testgrid.appspot.com/))
                        *   1.11 dashboards don't exist until 7 days prior
                *   Sending out update to k8s-dev 4/13
                *   Could get freeze to 7 working days (goal)
                *   Up to release lead if this works out
                *   Questions to sig-release mailing list/slack
    *   Patch Release Updates
        *   1.10.1 [Maciek Pytel](confirmed)
            *   Issues with flakey tests, triaged, reviewing and hopefully green by 4/13
        *   1.9.
            *   No info
*   [ 10:24 ] **Graph o' the Week **[Aaron Crickenberger](confirmed)
    *   devstats.k8s.io - PRs labels repository groups
    *   [https://k8s.devstats.cncf.io/d/47/prs-labels-repository-groups?orgId=1](https://k8s.devstats.cncf.io/d/47/prs-labels-repository-groups?orgId=1)
        *   needs-rebase - trend over time shows abandoned pull requests
        *   [https://github.com/kubernetes/test-infra/tree/master/label_sync](https://github.com/kubernetes/test-infra/tree/master/label_sync)
        *   need-ok-to-test - shows pull requests need help to get through process
            *   Members don't get this tag
                *   [https://github.com/kubernetes/community/blob/master/community-membership.md](https://github.com/kubernetes/community/blob/master/community-membership.md)
                *   Requirements have changed - Demonstrate intent and dedication to the project
    *   Number of PR's that have a given label applied over time
    *   [http://not.oktotest.com](http://not.oktotest.com)
    *   need-rebase
*   [ 10:32 ] **SIG Updates**
    *   SIG-VMware [Steve Wong](confirmed)
        *   1st meeting, charter defined, meeting notes are published and shared
        *   31 google group members, 11 people on first meeting, 50 slack members
        *   Support kubernetes users who are deploying at scale on VMware platforms, support development relating to cloud providers
            *   Working on aligning cloud provider to cloud provider WG strategy
    *   SIG-Windows  [Michael Michael](confirmed)
        *   Busy since 1.9 on getting people using and deploying windows containers
        *   Been fixing bugs
        *   Added support
            *   Resource controls
            *   File system stats
            *   Flexv
            *   Hyper-v isolation (experimental), similar to kata
            *   e2e automation and tests for sig-windows (eta 1 month)
            *   Hoping for GA around Windows Server 2019
            *   Join on slack or directly
    *   On-Prem announcement (jb)
        *   Demoted to working group
            *   Forum for discussion for people with on-prem deployments
            *   Possibly formally recognized group in future (waiting decision from steering committee)
                *   Committee is clarifying what a WG is, why, rules
                *   Also clarifying a sub-project
            *   Doesn't have formal meetings/structure yet
*   [ 10:40 ] **Announcements**
    *   Shoutouts this week:  
        *   @robinpercy@mike.splain @rdrgmnzs for graduating the mentoring program.  @paris for leading it.
        *   @maciekpytel for "providing clarity" around the node autoscaler
    *   We're going to disable the cherrypick-auto-approve munger [aaron]
        *   Process now is based on cherrypick-approve
        *   [kubernetes-dev@ thread](https://groups.google.com/d/msg/kubernetes-dev/Br2-4pQPOIs/YbSM1YNIBgAJ)
    *   SC update: subprojects and SIG charters [briangrant]
        *   Help SIGs create charters
        *   Explain to SIGs what sub-projects are and how they can be used
        *   Split SIGs among steering committee members
        *   6 charters in flight working on charter, then going to other SIGs
    *   [r/kubernetes: Ask Me Anything](https://www.reddit.com/r/kubernetes/comments/8b7f0x/we_are_kubernetes_developers_ask_us_anything/) - thanks everyone for participating, lots of user feedback, please have a look.
        *   We'll likely do more of these in the future.
    *   [Kubernetes Contributor Summit @ KubeCon/CloudNativeCon](https://github.com/kubernetes/community/tree/master/events/2018/05-contributor-summit) - May 1 (jb)
        *   You need to register for this even if you already registered for KubeCon/CloudNativeCon! Link to the form in the link above.
        *   New contributor/on-going contrib in morning and general tracks in afternoon
    *   New CNCF Interactive Landscape: [https://landscape.cncf.io/](https://landscape.cncf.io/)  (dan kohn)


## April 5, 2018 - ([recording](https://www.youtube.com/watch?v=z1vLGqNAJuA))



*   **Moderators**: Jose Palafox [ContribEx]
*   **Note Taker**: Solly Ross
*   **Demo **
    *   Artifactory in Kubernetes - Jainish Shah ([jainishs@jfrog.com](mailto:jainishs@jfrog.com)), Craig Peters ([craigp@jfrog.com](mailto:craigp@jfrog.com))
    *   Artifactory is a universal repository/artifact manager (e.g. docker registry, helm repositories)
    *   Can be deployed via Helm ([https://hub.kubeapps.com/charts/stable/artifactory](https://hub.kubeapps.com/charts/stable/artifactory))
    *   Can proxy/mirror/cache upstream repositories, and store artifacts itself
    *   Demo shows:
        *   creating docker registry and helm repos, pushing helm chart
        *   CLI and web UI
        *   Caching upstream repositories
    *   Walkthrough and Example: [https://jfrog.com/blog/control-your-kubernetes-voyage-with-artifactory/](https://jfrog.com/blog/control-your-kubernetes-voyage-with-artifactory/) & [https://github.com/jfrogtraining/kubernetes_example](https://github.com/jfrogtraining/kubernetes_example)
    *   Questions
        *   Difference between commercial and free (and what's the cost)
            *   Free only has maven support, is open source, commercial supports everything (including Kubernetes-related technologies, like Helm)
        *   Is HTTP basic auth the default
            *   Yes, but other auth schemes are supported
            *   Use of the API key in the jfrog cli documented [https://www.jfrog.com/confluence/display/CLI/CLI+for+JFrog+Artifactory](https://www.jfrog.com/confluence/display/CLI/CLI+for+JFrog+Artifactory)
*   **Release Team**
    *   1.11 [Josh Berkus, Release Lead]
        *   We are in Week 1
        *   [Release team roles](https://github.com/kubernetes/sig-release/blob/master/releases/release-1.11/release_team.md) almost filled
            *   Still need a docs shadow (ask in #sig-release or #sig-docs if interested)
        *   [Schedule posted](https://github.com/kubernetes/sig-release/blob/master/releases/release-1.11/release-1.11.md) (last chance to raise issues is _right now_)
            *   Feature Freeze: April 24
            *   Code slush: May 22
            *   Code Freeze: May 28
            *   Doc Deadline: June 11
            *   Release: June 26 (yes, we know it's right before July 4th, there's not a good way around it)
        *   Feature Gathering has started
            *   [Tracking Sheet](https://docs.google.com/spreadsheets/d/16N9KSlxWwxUA2gV6jvuW9N8tPRHzNhu1-RYY4Y0RZLs/edit#gid=0)
            *   **File your feature repo issues now!**
    *   1.10 [Maciek Pytel, patch manager]
        *   1.10.1 on April 12th (next Thursday)
            *   Contains kubectl fixes
            *   Will send email with release notes and pending PRs later today
            *   Contact Maciek if you need to cherry-pick anything not in the email
    *   1.9
*   **Graph o' the Week **(Josh Berkus)
    *   PR Workload: shows weighted workload for each SIG in the form of PRs.  Primarily there to check up on which SIGs are heavily loaded.  
        *   Includes ONLY PRs against kubernetes/kubernetes because other repos don't use SIG labels consistently.
        *   "Absolute Workload" == PR count * PR Size
        *   "Relative Workload" == (PR Count * PR Size) / Number of Reviewers
        *   [Calculation in issue (docs later)](https://github.com/cncf/devstats/issues/68)
        *   PRs that were *open* during that period, not PRs that were created during the period.  Old PRs are important for workload.
    *   Chart:
        *   [All PRs, 6 months](https://k8s.devstats.cncf.io/d/33/pr-workload?orgId=1)
        *   [SIG-Scalability And Autoscaling last 3 months](https://k8s.devstats.cncf.io/d/33/pr-workload?orgId=1&from=now-90d&to=now&var-sigs=autoscaling&var-sigs=scalability&var-full_name=Kubernetes)
        *   [SIG-Node last 3 months](https://k8s.devstats.cncf.io/d/33/pr-workload?orgId=1&from=now-90d&to=now&var-sigs=node&var-full_name=Kubernetes)
    *   Table:
        *   [Workloads for last month](https://k8s.devstats.cncf.io/d/34/pr-workload-table?orgId=1)
        *   [Workload for version 1.9](https://k8s.devstats.cncf.io/d/34/pr-workload-table?orgId=1&var-period_name=v1.8.0%20-%20v1.9.0&var-period=anno_28_29)
        *   [Workload for version 1.10](https://k8s.devstats.cncf.io/d/34/pr-workload-table?orgId=1&var-period_name=v1.9.0%20-%20v1.10.0&var-period=anno_29_30)
    *   Questions
        *   Assumes proper labeling
            *   Yes, things without sig labels aren't included
            *   Size labels have a lot of fudge factor
            *   Don't compare small increments ("10% more PRs than last week") but compare heavily loaded/lightly loaded ("is SIG-API overwhelmed for 1.11?")
        *   Overall graph takeaway?
            *   Determine if one SIG has a bunch of PRs suddenly, or does everyone have a bunch of PRs due to a particular place in the release cycle
        *   What do we do based on these charts?
            *   We're still exploring, seeing what people think, are these helpful to people?
            *   Join #devstats if you're interested in collaborating, discussing
            *   Need to find charts that are actually useful vs just shiny
*   **SIG Updates**
    *   SIG Cluster Ops - Rob H.
        *   Updated Mission - more focused on building operator community
        *   We have been having trouble with quorum and need some help
        *   Chris McEniry and I believe strongly in the need for an vendor neutral place for operators to gather around K8s.  Neither of us are vendors, so we're good neutral hosts, but we need help getting speakers.
        *   We'd be happy to host long format demos by vendors.
        *   Questions:
            *   Consolidate with OnPrem?
                *   Sure!
                *   It's good to consolidate SIGs if it's relevant
    *   SIG Docs - Zach
        *   2 new maintainers, 5 new contributors
        *   Coming soon: improved contribution guidelines: https://kubernetes.io/editdocs/
        *   Migrating the Kubernetes website from Jekyll to Hugo: we've met with the contractor, gotten an initial estimate, and are proceeding with a target completion date of April 30
        *   Blog migration: Formerly at blog.kubernetes.io (Blogger), the Kubernetes blog now resides at kubernetes.io/blog (GitHub).  The main reason for migrating was to resolve increasing technical debt and make life easier on the blog contribution team.
            *   Thanks to test-infra for the automation that make it possible to have blog-level ownership of PRs and approvals!
        *   Questions:
            *   Tidying up contributor guidelines just for SIG docs, or Kubernetes in general?
                *   Kubernetes in general
                *   Collaborate with ContribEx
            *   Have blog posts pending review been migrated to GitHub, or do they need to be manually migrated?
                *   They will be migrated, with blog manager opening PRs as needed
    *   SIG Service Catalog - bumped to 5/24
*   **Announcements**
    *   [Kubernetes Contributor Summit @ KubeCon/CloudNativeCon](https://github.com/kubernetes/community/tree/master/events/2018/05-contributor-summit) - May 1 [Jorge Castro]
        *   You need to register for this even if you already registered for KubeCon/CloudNativeCon! Link to the form in the link above.
        *   Current contributor track voting on topics will be emailed to attendees Monday
    *   Reddit r/kubernetes AMA [Jorge Castro]
        *   This next Tuesday: [https://www.reddit.com/r/kubernetes/comments/89gdv0/kubernetes_ama_will_be_on_10_april_tuesday/](https://www.reddit.com/r/kubernetes/comments/89gdv0/kubernetes_ama_will_be_on_10_april_tuesday/)
        *   If you're a reddit user please contact [jorge@heptio.com](mailto:jorge@heptio.com) so we can coordinate.
    *   Roadmaps call - SIG-PM is asking for the roadmap input [https://groups.google.com/forum/#!topic/kubernetes-sig-pm/-jW3bHUbfE8](https://groups.google.com/forum/#!topic/kubernetes-sig-pm/-jW3bHUbfE8) [Ihor/Jaice]
        *   We're trying to assess if there are any long-term, cross-cutting views of the project, or if our planning horizon is only the length of features in process/one release
        *   We want to provide more transparency to the end-user community about what planning exists
        *   SIG-PM can also help facilitate planning, as in this [template](https://docs.google.com/document/d/1qi4LKV3W9B5JJ5JLjmAY33jESYAqWMWFyO5bWoYEBDo/edit)
    *   [https://go.k8s.io/github-labels](https://go.k8s.io/github-labels) [Aaron C]
        *   List of all labels created or consumed by automation, as well as their meanings
        *   Lists instructions on how to contribute new labels
    *   [Kubernetes Application Survey](https://goo.gl/forms/ht61kKETiqVR103v1) [Aaron C]
        *   From [wg-app-def](https://github.com/kubernetes/community/tree/master/wg-app-def)
        *   How do you build and operate application on kubernetes?
        *   Results will be made publicly available
        *   Take the survey, share it with others
        *   Due April 16th
    *   [http://k8s-code.appspot.com/](http://k8s-code.appspot.com/) (Dims)
        *   Search engine for all github code repositories under kubernetes main org
*   **Shoutouts**
    *   Thanks to our contributors that joined #meet-our-contributors yesterday for questions! @gsaenger @spiffxp @chrislovecnm @spzala @carolynvs
        *   [Once a month livestream on-demand upstream mentors! Join us!](https://github.com/kubernetes/community/blob/master/mentoring/meet-our-contributors.md)
        *   Considering twice a month - need help with running more! Contact paris@ on slack or parispittman@google.com


## March 29, 2018



*   **Moderators**:  Paris Pittman [ContribEx]
*   **Note Taker**: Solly Ross
*   [ 0:00 ]**  Demo - **KubeAdmin Self-Config - Rob Hirschfeld ([rob@rackn.com](mailto:rob@rackn.com))
    *   Uses Digital Rebar and kubeadm to stand up clusters from scratch without no intervention
    *   Digital Rebar does PXE provisioning to get the machines stood up, and then hands of to kubeadm for installing kubernetes
    *   Workflow to do the setup sets, install Docker etc, elect master, do hand off to kubeadm
        *   Generates join keys for kubeadm
        *   Sends information like master election, cluster admin config file, etc back to shared data set
    *   Resources:
        *   KubeCon/CloudNativeCon Presentation [https://www.slideshare.net/rhirschfeld/kubecon-2017-zero-touch-kubernetes](https://www.slideshare.net/rhirschfeld/kubecon-2017-zero-touch-kubernetes)
        *   Longer Demo Video [https://www.youtube.com/watch?v=OMm6Oz1NF6I](https://www.youtube.com/watch?v=OMm6Oz1NF6I)
        *   Digital Rebar:[https://github.com/digitalrebar/provision](https://github.com/digitalrebar/provision),
        *   Project Site:  [http://rebar.digital](http://rebar.digital)
        *   Terraform Provider (referenced at end) [https://github.com/rackn/terraform-provider-drp](https://github.com/rackn/terraform-provider-drp)
    *   Questions:
        *   Q: Could we drive digital rebar from Kubicorn?
            *   A: Yes, probably
*   [ 0:00 ] **Announcements**
    *   [Meet Our Contributors is Apr 4th @ 330p and 9p UTC](https://github.com/kubernetes/community/blob/master/mentoring/meet-our-contributors.md)
        *   Like office hours, but contributor question focused
        *   Looking for contributors to answer questions, 2 slots
        *   Reach out to @paris on Slack if you're interested in participating
    *   Contributor Summit in Copenhagen May 1 - [registration](https://events.linuxfoundation.org/events/kubecon-cloudnativecon-europe-2018/co-located-events/kubernetes-contributor-summit/) is live
    *   KubeCon/CloudNativeCon Copenhagen (May 2-4) is **on track to sell out**. [Register](https://events.linuxfoundation.org/events/kubecon-cloudnativecon-europe-2018/)
    *   Shoutouts this week (from #shoutouts in slack):
        *   @nabrahams who picked the 1.10 release notes as his first contribution. We literally could not have done this without him!
*   [ 0:15 ]** Kubernetes 1.10 Release Retrospective**
    *   Retro [doc](https://docs.google.com/document/d/1kZnDqR0rZ4Zj_D9WWdD5JIoF9dZdZRr0giIU0w32bqI/edit#)
    *   This is how we improve
    *   Prior Release Retrospectives [ [1.3](http://bit.ly/kube13retro), [1.4](http://bit.ly/kube14retro), [1.5](http://bit.ly/kube15retro), [1.6](http://bit.ly/kube16retro), [1.7](http://bit.ly/kube17retro), [1.8](http://bit.ly/kube18retro), [1.9](http://bit.ly/kube19retro) ]


## March 22, 2018 - recording



*   **Moderators**:  Chris Short [Contribex]
*   **Note Taker**: Sanket Patel [Egen Solutions]
*   [ 0:00 ]**  Demo **-- 03/22: Ark - a backup/disaster recovery tool for k8s - Andy Goldstein ([andy@heptio.com](mailto:andy@heptio.com))
    *   [https://github.com/heptio/ark](https://github.com/heptio/ark)
    *   "Heptio Ark is a utility for managing disaster recovery, specifically for your Kubernetes cluster resources and persistent volumes."
    *   Looking for help on code and documentation
    *   Link to slides
*   [ 0:00 ]** Release Updates**
    *   1.10 [Jaice Singer DuMars ~ Release lead]
        *   Code thaw happened
        *   rc1 is out, please take a look!
        *   release due Monday 3/26 ~ 6PM Pacific time
        *   release retrospective the last 45 minutes of the community meeting next week
        *   You can see our progress in [videos](https://www.youtube.com/watch?v=e14tlUBd2jQ&list=PL69nYSiGNLP3QKkOsDsO6A0Y1rhgP84iZ) and in the burndown meeting [notes](http://bit.ly/k8s110-burndown)
        *   We're looking for a Release Lead for 1.11, code name "Everything Croissant" - read the role description and requirements [here](https://github.com/kubernetes/sig-release/blob/master/release-process-documentation/release-team-guides/release-lead.md), and contact @jdumars in Slack if you are interested in knowing more
            *   Josh Berkus has volunteered, but the more the merrier!
        *   Release retrospective in this time slot, in 2 weeks: 3/29
    *   Patch releases out now
        *   1.7.15
        *   1.8.10
        *   1.9.6
*   [ 0:00 ] **Graph o' the Week **[N/A]
    *   Back Next Week After 1.10 Release
*   [ 0:00 ] **SIG Updates**
    *   SIG Azure [Cole Mickens] (confirmed)
        *   #sig-azure on Slack
        *   Azure backlog is public
        *   cloudprovider is moving to independent repo
        *   Lots of work around Azure integrations
        *   Kal Henidak is leading Azure's upstream release efforts
        *   Many fixes in the 1.10 release
        *   More great things in the pipeline
    *   SIG Node [Derek Carr] (confirmed)
        *   New [CRI testing policy](/contributors/devel/sig-node/cri-testing-policy.md)
        *   Feature going into Beta - Local storage capacity isolation
        *   Feature going into alpha - debug container, supports pod pid limits, cri container log rotation
        *   wg-resource-mgmt : graduated device plugins, hugepages, cpu pinning (beta)
        *   Cri-o declared stable since 1.9x
        *   Future changes
            *   Finish governance materials
            *   Feature planning for 1.11
            *   Topics explored:
                *   Secure container
                *   Virtual Kubelet
        *   Face to face meeting details on working group document
        *   Slides presented: [https://docs.google.com/presentation/d/1P267xBGQtLprbVV-XStpVt8c-um6NqBmQLEIOKhmJAs/edit?usp=sharing](https://docs.google.com/presentation/d/1P267xBGQtLprbVV-XStpVt8c-um6NqBmQLEIOKhmJAs/edit?usp=sharing)
    *   SIG Network [Casey Davenport] (confirmed)
        *   Feature alpha -> beta
            *   IPv6 network support for k8s core components and pod networking - [feature issue](https://github.com/kubernetes/features/issues/508)
            *   CI for IPv6 - PR coming soon.
            *   [Core DNS integration](https://github.com/kubernetes/community/pull/1956) replacement for kubeDNS (single binary, better performance)
                *   [Feature issue](https://github.com/kubernetes/features/issues/427)
        *   IPvs kube-proxy staying in beta - [outstanding issues](https://github.com/kubernetes/kubernetes/issues?q=is%3Aopen+is%3Aissue+label%3Aarea%2Fipvs)
        *   Ingress requirements gathering.
            *   Sent survey to ingress users.
            *   Have [results from survey](https://github.com/bowei/k8s-ingress-survey-2018) and starting interpretation
        *   Network plumbing group discussing adding networking interfaces for pods - [specification doc](https://docs.google.com/document/d/1Ny03h6IDVy_e_vmElOqR7UdTPAG_RNydhVE1Kx54kFQ/edit)
        *   Traffic shaping moved to a CNI plugin - [proposal](https://github.com/kubernetes/community/pull/1893)
        *   Designing service routing to be topology aware - [proposal](https://github.com/kubernetes/community/pull/1551)
*   [ 0:00 ] **Announcements**
    *   Shoutouts this week:
        *   Vlad Ionescu: Shoutout to @yomateo and @foxie for how much they are helping in #kubernetes-users! Matthew has a lot of examples on his Github and Ashley is amazing regarding on-prem and general k8s questions( learned so much by reading their answers).
        *   Jaice DuMars: I know I keep saying it, but @cblecker @dims @liggitt ~ The POWER TRIO of release-assist awesomeness!!  THANK YOU!!
        *   Josh Berkus: Shout out to Pavel Pospisil, for pitching in to close a PVC blocker bug despite being between jobs this week and not having his normal dev machine.
        *   Josh Berkus: Shoutout to Klaus Ma for leading the closure of several storage bugs, including staying up until 1am his time to attend meetings
        *   Jordan Liggitt: Shoutout to the branch managers: @wojtekt, @jpbetz, @mbohlool for dealing with unending cherry-picks over the last couple weeks
    *   Meet Our Contributors is looking for contributors to come on Apr 4th at 330 and 9p UTC
        *   Join the slack channel (same name)


## March 15, 2018- ([recording](https://www.youtube.com/watch?v=tvP_HPFteKI))



*   Moderators:  Solly Ross [SIG]
*   Note Taker: Kris Nova
*   [ 0:00 ]  Demo --  Kuberhealthy presentation (sandbox project applicant) - Eric Greer (eric.greer@comcast.com)
    *   [https://docs.google.com/presentation/d/1tL80i7VTBUlDs5KXy7TBZcHVX45lFST2cmmE3oyjYWc/edit?usp=sharing](https://docs.google.com/presentation/d/1tL80i7VTBUlDs5KXy7TBZcHVX45lFST2cmmE3oyjYWc/edit?usp=sharing)
    *   Repository TBD (may be open sourced via corporation or go direct to sandbox if applicable)
    *   Tool for checking cluster health by standing up pods on each node, checking status of components in kube-system, etc
        *   Designed to check cluster health beyond normal logs, monitoring, metrics -- make sure you can actually run applications
    *   Currently deployed internally in production, but should generally be considered alpha at the moment
*   [ 0:00 ] Release Updates
    *   1.10 Update [Jaice Singer DuMars ~ Release lead]
        *   Release team is in a meeting right meow
        *   Due to the security releases, and some scalability testing issues we decided to push the release from 3/21 to 3/26, and lift code freeze by EOD Monday assuming all goes to plan
        *   You can see our progress in [videos](https://www.youtube.com/watch?v=e14tlUBd2jQ&list=PL69nYSiGNLP3QKkOsDsO6A0Y1rhgP84iZ) and in the burndown meeting [notes](http://bit.ly/k8s110-burndown)
        *   We're looking for a Release Lead for 1.11, code name "Everything Croissant" - read the role description and requirements [here](https://github.com/kubernetes/sig-release/blob/master/release-process-documentation/release-team-guides/release-lead.md), and contact @jdumars in Slack if you are interested in knowing more
        *   Release retrospective in this time slot, in 2 weeks: 3/29
    *   Current Release cycle (e.g. 1.7) [First Last ~ Release role]
    *   Prior Release cycle (e.g. 1.6.6) [First Last ~ Release role]
*   [ 0:00 ] Graph o' the Week
    *   No graph this week, tune in next time!
*   [ 0:00 ] SIG Updates
    *   SIG Auth [Eric Chiang]
        *   Notes: [[link]](https://docs.google.com/document/d/1wyOkDwRDQetjTBeaPbJfkt1M8f_q3nxN5ta5v_OTfzA/edit)
        *   Overview
            *   Betas: PodSecurityPolicy, Auditing
            *   Alphas: TokenRequest API, client-go external credential providers, encryption-at-rest external KMS integration
            *   Considering Bug Bounty for Kubernetes
    *   SIG Instrumentation [Piotr Szczesniak]
        *   Introduced external metrics API
            *   Metrics will come from other systems other than k8s
            *   Integrates with Prometheus and other monitoring systems with adapter
        *   Graduated metrics API to beta
            *   Inspired by heapster
            *   By default applied to all kubernetes clusters
        *   Started discussion around securing instrumentation endpoint
            *   Cross SIG effort between sig-auth and sig-instrumentation
        *   Master metrics API, Custom metrics API
            *   Plans to graduate these to GA
            *   Deprecate Heapster
        *   Plans in the works for a historical metrics API
        *   Agreement to have logging on architecture and vision
            *   Similar to metrics architecture, which was a foundation for many design decisions
        *   Discussion on exposing kubelet health status
            *   Useful for monitoring the state of Kubelet
            *   Different than the metrics endpoint
            *   Still figuring out how to do this
        *   Would like to move sig-instrumentation projects to a new home
*   [ 0:00 ] Announcements
    *   Registration for the Contributor Summit is now live:
        *   See [this page](https://events.linuxfoundation.org/events/kubecon-cloudnativecon-europe-2018/co-located-events/kubernetes-contributor-summit/) for details
        *   Please register if you're planning on attending, we need this so we have the correct amount of food!
        *   Just registering for KubeCon/CloudNativeCon is not enough!
    *   [Office Hours Next Week!](https://github.com/kubernetes/community/blob/master/events/office-hours.md)
        *   Volunteer developers needed to answer questions
    *   [Helm Summit Videos](https://www.youtube.com/playlist?list=PL69nYSiGNLP3PlhEKrGA0oN4eY8c4oaAH&disable_polymer=true) are up.
    *   Shoutouts this week
        *   Someone doing great work out there that you'd like to highlight? Let us know in #shoutout on slack and we'll mention them here:
        *   [@shyamjvs](https://github.com/shyamjvs) - diagnosing 2 critical performance problems which could have blocked the 1.10 release
        *   [@verullt](https://github.com/verult) - taking on most of the open storage issues and resolving them as quickly as possible
        *   Quang Huynh for all the shiny new looks he's given prow.k8s.io during his internship!
        *   Jaice wants to say "Release mvps: [@dims](https://github.com/dims) [@cblecker](https://github.com/cblecker)"
        *   Thanks to the ever-helpful Andrew Chen ([@chenopis](https://github.com/chenopis)) for getting me through merge conflicts and branch merges in preparation for the 1.10 release. (From Jennifer Rondeau)
        *   Stefan Schimanski would like to thank:
            *   Nick Chase ([@nickchase](https://github.com/nickchase)) - for editing down 92 pages of release notes into something consumable by humans!
            *   Mik Vyatskov ([@crassirostris](https://github.com/crassirostris)) - for doing an awesome job driving and owning the auditing feature
        *   Josh Berkus would like to thank Jordan Liggitt ([@liggitt](https://github.com/liggitt))
        *   And finally congratulations to Brad Topol ([@bradtopol](https://github.com/bradtopol)) for joining the maintainer team on docs.


## March 8, 2018 - ([recording](https://youtu.be/fySqkBQnJ8I))



*   **Moderators**: Jorge Castro [SIG Contrib Ex]
*   **Note Taker**: This could be you!  [Company/SIG]
*   [ 0:00 ]**  Demo **-- KubeFlow - Jeremy Lewi ([jlewi@google.com](mailto:jlewi@google.com))
    *   [Link to slides](https://docs.google.com/presentation/d/1p82_DKJmIPjFS69EJ8p4StCClJW-JjGgxSWfhzw7Abw/edit#slide=id.g30f6ce7d33_0_378)
    *   [Kubeflow Repo](https://github.com/kubeflow/kubeflow)
*   [ 0:00 ]** Release Updates**
    *   1.10 [Jaice Singer DuMars ~ Release lead]
        *   **Week 10 of 12**, the full schedule and some important information is [here](https://github.com/kubernetes/sig-release/blob/master/releases/release-1.10/release-1.10.md)
        *   Release status is yellow, which means there's a chance our release date might slip by a small margin of days ~ we're working on sorting this out and should know more by early next week
        *   Next week, we will be entering [c](https://github.com/kubernetes/sig-release/blob/master/releases/release-1.10/release-1.10.md#code-slush)runch week for the release, so majority of release activities like documentation and release notes are completed, as well as drafts for blog posts, and FAQs for media.
        *   To meet our release date, the v1.10 branch must be in very good shape by the end of next week, with fixes either en route, or in queue.  
        *   **Code freeze will end on March 14th at 6PM Pacific time**, after which any additions to the 1.10 release will need to be cherry picked, in close collaboration with the release team.
        *   Release team meetings will move to the daily cadence next week. If you join either the [SIG-Release](https://groups.google.com/forum/#!forum/kubernetes-sig-release) or [kubernetes-milestone-burndown](https://groups.google.com/forum/#!forum/kubernetes-milestone-burndown) groups, you should get an invite.
*   [ 0:00 ] **Graph o' the Week **Zach Corleissen, SIG Docs
    *   Weekly update on data from devstats.k8s.io
    *   [https://k8s.devstats.cncf.io/d/44/time-metrics?orgId=1&var-period=w&var-repogroup_name=Docs&var-repogroup=docs&var-apichange=All&var-size_name=All&var-size=all&var-full_name=Kubernetes](https://k8s.devstats.cncf.io/d/44/time-metrics?orgId=1&var-period=w&var-repogroup_name=Docs&var-repogroup=docs&var-apichange=All&var-size_name=All&var-size=all&var-full_name=Kubernetes)
    *   Docs folks had vague anxiety (without concrete data) on their response times for issues and PRs.  Devstats shows less than approx. 4 days initial response times during the last year, outside of a few spikes associated with holidays on the calendar and KubeCon/CloudNativeCon.
    *   Introduction of prow into kubernetes/website led to a demonstrable improvement in early 2018
*   [ 0:00 ] **SIG Updates**
    *   SIG Apps [Adnan Abdulhussein] (confirmed)
        *   [https://docs.google.com/presentation/d/1yTM5bi4C2cr_L-Ow1G2W934-vXS3r_z-PayDAaKx73o/edit?usp=sharing](https://docs.google.com/presentation/d/1yTM5bi4C2cr_L-Ow1G2W934-vXS3r_z-PayDAaKx73o/edit?usp=sharing)
    *   SIG OpenStack [Chris Hoge] (confirmed)
        *   [https://docs.google.com/presentation/d/1DtBKFlPhb74v9bXN6-RSNpLagh2wbDs3JJuno7IzgSw/edit?usp=sharing](https://docs.google.com/presentation/d/1DtBKFlPhb74v9bXN6-RSNpLagh2wbDs3JJuno7IzgSw/edit?usp=sharing)
    *   SIG UI [Sebastian Floreks] (sends regrets)
        *   Due to reasons independent from us, me and maciaszczykm have been moved from full-time Dashboard contribution to another project. We will be less active now, but still want to finish the migration and try to be a part of Dashboard project.
        *    We are working only on migration from AngularJS to Angular 5. Around 70% of features have been rewritten. Unfortunately, due to mentioned reasons we can not provide any ETA regarding the end of migration process. Progress and changes can be tracked from: [https://github.com/kubernetes/dashboard/pull/2727](https://github.com/kubernetes/dashboard/pull/2727)
        *   Dashboard is on a soft code freeze until migration is finished. Only critical bugs will be fixed during this time.
        *   (See notes from the last few meetings): [https://github.com/kubernetes/community/tree/master/sig-ui](https://github.com/kubernetes/community/tree/master/sig-ui)
*   [ 0:00 ] **Announcements**
    *   SIG Charter and Subprojects Update [pwittrock]
        *   [SIG Governance Charter Templates](https://github.com/kubernetes/community/blob/master/committee-steering/governance/README.md)
            *   At least one more detailed template under development
        *   [Governance.md updated with subprojects](https://github.com/kubernetes/community/blob/master/governance.md#subprojects)
        *   [WIP: Subproject Meta](https://docs.google.com/document/d/1FHauGII5LNVM-dZcNfzYZ-6WRs9RoPctQ4bw5dczrkk/edit#heading=h.2nslsje41be1)
        *   [WIP: Charter FAQ (the "Why"s)](https://github.com/kubernetes/community/pull/1908)
    *   Reminder: [Contributor Summit](https://github.com/kubernetes/community/tree/master/events/2018/05-contributor-summit), 1 May, day before KubeCon/CloudNativeCon
    *   CNCF would like feedback on the draft blog post for 1.10 beta:
        *   [https://kubernetes.io/blog/2018/03/first-beta-version-of-kubernetes-1-10/](https://kubernetes.io/blog/2018/03/first-beta-version-of-kubernetes-1-10/)
        *   Please contact [Natasha Woods](mailto:nwoods@linuxfoundation.org) with your feedback
    *   Shoutouts this week
        *   See someone doing something great for the community? Mention them in #shoutouts on slack.
        *   Maru Newby (@marun) for [https://github.com/kubernetes/test-infra/pull/7083](https://github.com/kubernetes/test-infra/pull/7083) also Cole Wagner (@cjwagner) and Benjamin Elder (@bentheelder) who have all be super helpful getting this release moving forward.
        *   Meet our Contributors (1st weds of the month - AMA kubernetes.io/community for cal invite) - Aaron Crickenberger (@spiffxp), Davanum Srinivas (@dims), Ilya Dmitrichenko (@errordeveloper), Jennifer Rondeau ( @jrondeau), Kris Nova (@kris-nova), Solly Ross (@directxman12), Jeff Grafton (@ixdy) and Jorge Castro (@jorge)


## March 1, 2018 - ([recording](https://youtu.be/mpfqSBcdSHI))



*   **Moderators**:  Solly Ross [SIG]
*   **Note Taker**: Chris Short - chris@chrisshort.net
*   [ 0:00 ]**  Demo **-- Sonobuoy - diagnostic tool for k8s - Chuck Ha ([chuck@heptio.com](mailto:chuck@heptio.com))
    *   [Link to slides](https://docs.google.com/presentation/d/1aiCdN5RY-mCZdqav5RuVpJfgzW9UGo1-Zlq3aITZ830/edit#slide=id.g3378780fb7_0_65)
    *   [Link to repository](https://github.com/heptio/sonobuoy)
    *   #sonobuoy on the k8s slack
    *   "We love feedback"
*   [ 0:00 ]** Release Updates**
    *   1.10 Release Update [Jaice Singer DuMars ~ Release lead]
        *   Week 9 of 12, full schedule and some important information is [here](https://github.com/kubernetes/sig-release/blob/master/releases/release-1.10/release-1.10.md)
        *   We're in code freeze
            *   Why do we have code freeze?
                *   Provides opportunity for SIGs to focus on bugs/cleanup
                *   Allows time for technical debt elimination while we have a quiet submit queue, e.g. [this](https://groups.google.com/d/msg/kubernetes-dev/UFCzs-Zjj9E/b3_qjj71AwAJ)
                *   
            *   How does it work?
                *   Isolate release-relevant items (issues and PRs) by use of the **<code>status/approved-for-milestone</code></strong> label
                *   If it's in the milestone with that label, we pay close attention, otherwise it is assumed not impacting the release
                *   Bot will nag you about labels so we can focus on what's really important
            *   Check out that [beta](https://github.com/kubernetes/kubernetes/releases/tag/v1.10.0-beta.1), <strong>please</strong>
        *   Collecting [known issues](https://github.com/kubernetes/kubernetes/issues/59764) in a single place so we can properly document it as part of the release notes
        *   Release notes and user-facing documentation should be close to complete ~ <strong>PRs for docs need to be ready for review by tomorrow (6 PM PT unless otherwise stated)</strong>
    *   1.8.9 should be out today
    *   1.7.13 is out
*   [ 0:00 ] <strong>Graph o' the Week </strong>[Josh Berkus]
    *   Weekly update on data from devstats.k8s.io
    *   [Issues and PRs by Milestone](https://k8s.devstats.cncf.io/d/IIUa5kezk/open-issues-prs-by-milestone?orgId=1&from=now-7d&to=now&var-sig_name=All&var-sig=all&var-milestone_name=v1.10&var-milestone=v1_10&var-repo_name=kubernetes%2Fkubernetes&var-repo=kubernetes_kubernetes&var-full_name=Kubernetes)
    *   Allows us to compare workload/readiness with prior releases:
        *   [1.10 at code freeze](https://k8s.devstats.cncf.io/d/IIUa5kezk/open-issues-prs-by-milestone?orgId=1&from=1517101031268&to=1519779431269&var-sig_name=All&var-sig=all&var-milestone_name=v1.10&var-milestone=v1_10&var-repo_name=kubernetes%2Fkubernetes&var-repo=kubernetes_kubernetes&var-full_name=Kubernetes)
        *   [1.9 at code freeze](https://k8s.devstats.cncf.io/d/IIUa5kezk/open-issues-prs-by-milestone?orgId=1&from=1509407831268&to=1511830631269&var-sig_name=All&var-sig=all&var-milestone_name=v1.9&var-milestone=v1_9&var-repo_name=kubernetes%2Fkubernetes&var-repo=kubernetes_kubernetes&var-full_name=Kubernetes)
        *   [1.8 at code freeze](https://k8s.devstats.cncf.io/d/IIUa5kezk/open-issues-prs-by-milestone?orgId=1&from=1501804631268&to=1504483031269&var-sig_name=All&var-sig=all&var-milestone_name=v1.8&var-milestone=v1_8&var-repo_name=kubernetes%2Fkubernetes&var-repo=kubernetes_kubernetes&var-full_name=Kubernetes)
    *   Overcounting issue due to github bug
*   [ 0:00 ] <strong>SIG Updates</strong>
    *   SIG Big Data [Anirudh Ramanathan]
        *   Acting as a bridge to external projects, sometimes work on forks, and then upstream it.
        *   Apache Spark - (now tracked in [JIRA](https://issues.apache.org/jira/browse/SPARK-23529?jql=project%20%3D%20SPARK%20AND%20component%20%3D%20Kubernetes) & [mailing lists](https://spark.apache.org/community.html))
            *   Graduated from a fork and merged back into project - released as Spark 2.3 yesterday!
            *   Top billed feature - k8s support ([link](https://spark.apache.org/releases/spark-release-2-3-0.html))
            *   3 new ASF committers from our SIG
            *   Spark-submit with operator semantics ([WIP](https://github.com/GoogleCloudPlatform/spark-on-k8s-operator))
        *   Apache Airflow ([link](http://incubator.apache.org/projects/airflow.html))
            *   Trying to make a k8s-native DAG scheduler
            *   Currently upstreaming the k8s executor.
            *   Airflow v1.10 will be first release with k8s constructs in it.
        *   HDFS ([link](https://github.com/apache-spark-on-k8s/kubernetes-HDFS/))
            *   Hardening existing work
            *   Added HA namenode and fault tolerance to running within k8s/containers
            *   Demo coming soon in SIG Apps
            *   Success = performant and secure HDFS.
        *   Kube-arbitrator ([link](https://github.com/kubernetes-incubator/kube-arbitrator))
            *   Joint work with sig-scheduling.
            *   Ongoing discussion, MVP coming soon.
    *   SIG Storage [Saad Ali]  - [Slides](https://docs.google.com/presentation/d/1VNQQ9Lzn6ahy9zHm0aE6KOzS6YzCREzNE2Jla2ESbC8/edit?usp=sharing)
        *   Primary work for 1.10 is moving alpha features to beta including: CSI, Local storage, mount propagation, volume protection, ephemeral storage, etc.
        *   In 1.11 will be working on topology-aware storage and moving beta features to GA
    *   SIG Multicluster [Christian Bell]
        *   FederationV1: Low interest in fixing open bugs; bots are auto closing issues
            *   Top-level docs need to update/reflect reality that at the current pace of development, FederationV1 will not reach the maturity as single-cluster Kubernetes.
        *   Most work on federation has moved to a WG FederationV2
            *   Currently in "brainstorming" phase
            *   Moving away from having a completely consistent api with non-federated Kubernetes API
        *   [Cluster Registry](https://github.com/kubernetes/cluster-registry): An API for maintaining a list of clusters and associated metadata. Move to beta this quarter. Being consumed by Kubernetes (Kubemci) and non-Kubernetes projects (istio multi-cluster).
        *   [Kubemci](https://github.com/GoogleCloudPlatform/k8s-multicluster-ingress): A command-line tool (and eventually controller) to configure ingress across multiple clusters.
*   [ 0:00 ] <strong>Announcements</strong>
    *   [Owner/Maintainer ](https://github.com/kubernetes/community/pull/1861/files)[pwittrock]
        *   Maintainer is folding into Owner
    *   Reminder: Contributor Summit happens 1 May, day before KubeCon/CloudNativeCon
        *   [https://github.com/kubernetes/community/tree/master/events/2018/05-contributor-summit](https://github.com/kubernetes/community/tree/master/events/2018/05-contributor-summit)
    *   KubeCon/CloudNativeCon price increase March 9
        *   [https://events.linuxfoundation.org/events/kubecon-cloudnativecon-europe-2018/](https://events.linuxfoundation.org/events/kubecon-cloudnativecon-europe-2018/)
        *   Copenhagen May 2-4, 2018
    *   [Meet Our Contributors is next Weds!](https://github.com/kubernetes/community/blob/master/mentoring/meet-our-contributors.md)
        *   Two times! 330p and 9p UTC
        *   Ask current contributors anything on slack #meet-our-contributors - testing infra, how to make first time contribution, how did they get involved in k8s
    *   Shoutouts!
        *   None on slack this week, thank someone in #shoutouts!
        *   Top 5 in the Kubernetes StackOverflow tag for the week: Radek "Goblin" Pieczonka, aerokite, Vikram Hosakote, Jonah Benton, and fiunchinho


## February 22, 2018 - ([recording](https://www.youtube.com/watch?v=7pN0xdiFqPE))



*   **Moderators**:  Jorge Castro [SIG ContribEx]
*   **Note Taker**: Jaice Singer DuMars
*   [ 0:00 ]** Demo  **No demo this week
*   [ 0:01 ]** Release Updates**
    *   1.10 [Jaice Singer DuMars ~ Release lead]
        *   Week 8 of 12, full schedule and some important information is [here](https://github.com/kubernetes/sig-release/blob/master/releases/release-1.10/release-1.10.md)
        *   We're in code slush (Josh Berkus)
            *   Trying to filter out issues that are not critical to 1.10 release health
            *   SIGs must take ownership of [issues](https://github.com/kubernetes/sig-release/issues/86) with "Approved-For-Milestone" labels + Priority + type + status
        *   **CODE FREEZE ON MONDAY ~ 6PM PST**
            *   No new commits to the release-1.10 branch unless:
                *   SIG approved and tied to an existing issue
                *   Bug fixes (critical)
                *   Test flake fixes (or other release-specific code)
            *   If you have an exception, please contact @jdumars / [jdumars@gmail.com](mailto:jdumars@gmail.com)
        *   Release appears on track for 3/21 delivery, but don't count on it
        *   [https://github.com/kubernetes/sig-release/issues/86](https://github.com/kubernetes/sig-release/issues/86)
        *   If a release team member asks for something (e.g. docs), please make your best effort to help them.
*   [ 0:07 ] **Graph o' the Week **[Josh Berkus]
    *   Weekly update on data from devstats.k8s.io
    *   New and occasional contributors: [Issues](https://k8s.devstats.cncf.io/d/ey0DOdqzz/new-and-episodic-issues?orgId=1)  and [PRs](https://k8s.devstats.cncf.io/d/rCYj6D3kz/new-and-episodic-contributors?orgId=1)
        *   Shows volume of contributions/contributors from community members who are NOT full-time
        *   Important measure of how "welcoming" a community is.
        *   Also important because we've added a lot of process and need to be sure that's not a major blocker for new contributors.
*   [ 0:00 ] **SIG Updates**
    *   SIG Cluster Lifecycle [First Last]
        *   Not happening
*   [ 0:00 ] **Announcements**
    *   Reminder: Contributor Summit happens 1 May, day before KubeCon/CloudNativeCon
        *   [https://github.com/kubernetes/community/tree/master/events/2018/05-contributor-summit](https://github.com/kubernetes/community/tree/master/events/2018/05-contributor-summit)
    *   Shoutouts this week
        *   Zhonghu Xu - @hzxuzhonghu for many high quality apiserver APIs PRs
        *   Mario & Ilya Dmitrichenko for helping out with [user office hours](https://github.com/kubernetes/community/blob/master/events/office-hours.md) this week!
            *   Volunteers still needed!
        *   Someone doing great work? Give them a shoutout in #shoutouts and we'll mention them during the community meeting.
        *   Meet Our Contributors!!
            *   March 7th at two times
            *   #meet-our-contributors in slack for questions and more details
            *   We'd like to do live peer code reviews, too!
    *   Need SIG volunteers to say what is new in 1.10 in a webinar. Minimal time investment. -- contact [nchase@mirantis.com](mailto:nchase@mirantis.com) or @nickchase


## February 15, 2018 -[ (recording](https://www.youtube.com/watch?v=eqg5P81zPbs))



*   **Moderators**:  Josh Berkus [Contribex]
*   **Note Taker**: Tim Pepper [VMWare/Contribex]
*   [ 0:00 ]**  Demo **-- AppZ by Cloudbourne [Rejith Krishnan rkrishnan@cloudbourne.com] (c)
    *   Link to slides: n/a, see live demo in recording
    *   Youtube channel: [https://www.youtube.com/c/Cloudbourne](https://www.youtube.com/c/Cloudbourne)
    *   [https://github.com/rejith/tomcat-loadgen](https://github.com/rejith/tomcat-loadgen)
    *   Platform integrates SCM (GitHub), build (Maven, Gradle, Jenkins), and deploys/monitors app in k8s.  Builds on demand in response to commits in SCM.
    *   Dev, test, prod would use separate yaml files (example for a synthetic load generator using tomcat in [appz.yml](https://github.com/rejith/tomcat-loadgen/blob/master/appz.yml)), each describing the build/deploy/monitor needs for the app
*   [ 0:00 ]** Release Updates**
    *   1.10  [Jaice Singer DuMars ~ Release lead](c)
        *   week 7 of 12 of the Kubernetes 1.10 release cycle
        *   full schedule and some important information is [here](https://github.com/kubernetes/sig-release/blob/master/releases/release-1.10/release-1.10.md)
        *   we're looking on track to meet our release date of March 21st
        *   This week. we're cutting a beta, and setting up the release branch ~ nothing has changed in terms of merges
        *   Next week, we will be entering [code slush](https://github.com/kubernetes/sig-release/blob/master/releases/release-1.10/release-1.10.md#code-slush)  We use this time prior to [Code Freeze](https://github.com/kubernetes/sig-release/blob/master/releases/release-1.10/release-1.10.md#code-freeze) to help reduce issue noise from miscellaneous changes that aren't related to issues that SIGs have approved for the milestone. SIGs are the keepers of this label, not the release team, although we can add the label at the request of a SIG if needed
        *   **All issues in the milestone are considered release-blocking**. That allows any SIG to pull the "stop chain" on the release to prevent defects from impacting our community
        *   **code freeze** begins on** February 26th**
    *   1.9
        *   1.9.3 is out
        *   1.9.4 in 2 weeks (planned)
    *   1.8
        *   1.8.8 is out
*   [ 0:00 ] **SIG Updates**
    *   SIG Testing [Aaron Crickenberger](c)
        *   Held F2F meeting in Seattle Jan 26, thanks to EKS team for hosting (Bob Wise/AWS offers similar to any other SIG needing periodic f2f space in Seattle)
        *   Jenkins is no more, all jobs kicked off by prow now run natively on Kubernetes
        *   Testing commons subproject ([agenda](https://docs.google.com/document/d/1TOC8vnmlkWw6HRNHoe5xSv5-qv7LelX6XK3UVCHuwb0/edit#heading=h.tnoevy5f439o), [wed bi-weekly 7:30am pt](https://zoom.us/my/k8s.sig.testing)) is a forum for discussing how and what to test
        *   Testgrid in go as of jan 16 loads pages WAAAAAY faster
            *   [https://k8s-testgrid.appspot.com/](https://k8s-testgrid.appspot.com/)
            *   [Python vs go](https://www.dropbox.com/s/q1s98e1d1re3wdb/testgrid-go-vs-python.png)
        *   Prow UI updates
            *   [https://go.k8s.io/bot-commands](https://go.k8s.io/bot-commands)
            *   [https://prow.k8s.io/plugins](https://prow.k8s.io/plugins)
        *   Where are we at with tide these days
            *   [Umbrella issue](https://github.com/kubernetes/test-infra/issues/3866)
            *   [Discussion on how best to represent tide status](https://github.com/kubernetes/test-infra/issues/6145)
        *   [Implement bazel remote caching for faster builds](https://github.com/kubernetes/test-infra/issues/6808)
        *   [Proposal: upload conformance results to testgrid](https://docs.google.com/document/d/1lGvP89_DdeNO84I86BVAU4qY3h2VCRll45tGrpyx90A/edit)
        *   [Proposal: release-blocking and merge-blocking criteria](https://docs.google.com/document/d/1kCDdmlpTnHPQt5z8JzODdFCc3T2D4MKR53twsDZu20c/edit)
        *   Docker in Docker  / local e2e
            *   Get to a world where e2e doesn't require full blown cloud
            *   For local dev and for PR's, save on cluster standup/teardown time
            *   [https://github.com/kubernetes/kubernetes/pull/51661](https://github.com/kubernetes/kubernetes/pull/51661)
            *   [Discussed feb 6 meeting](https://docs.google.com/document/d/1z8MQpr_jTwhmjLMUaqQyBk1EYG_Y_3D4y4YdMJ7V1Kk/edit#heading=h.4quubo30kopo)
        *   Misc
            *   [Setting up automation for kubernetes-sigs org](https://github.com/kubernetes/test-infra/pull/6623)
            *   [The label_sync tool](https://github.com/kubernetes/test-infra/tree/master/label_sync)
    *   SIG Contribex [Paris Pittman](c)
        *   Charter
            *   Draft is being socialized in group now
            *   Using tl;dr template from SC (not approved yet but getting ahead of curve)
        *   Contributor Guide
            *   Solving for discoverability, holes in process/documentation, better flow
            *   New area in [k/community ](https://github.com/kubernetes/community/tree/master/contributors/guide)
            *   New on [Kubernetes.io; will be done by 1.10 release ](https://kubernetes.io/docs/imported/community/guide/)
        *   [Mentoring](https://github.com/kubernetes/community/tree/master/mentoring)
            *   Focusing on contributor membership growth; very important to reduce possibilities of burnout, learning and development of current contributors, etc
            *   Testing phase and learning a lot - Group Mentoring, Google Summer of Code, Outreachy, Meet Our Contributors, Proposed "Buddy" Guide Program
            *   Need help!
                *   [issue/1753](https://github.com/kubernetes/community/issues/1753), [issue/1803](https://github.com/kubernetes/community/issues/1803); need more outreachy organization sponsors and sig/wg mentors/projects; need more SIGs/WGs interested in group mentoring
                *   Building skills workshops for group mentoring and future k8s learning and dev. Examples in the [mentee guide](https://github.com/kubernetes/community/blob/master/mentoring/group-mentee-guide.md).
                    *   Examples: communication, code review, writing docs/rel notes, testing, etc.
        *   DevStats
            *   Working on User guide -> [https://github.com/cncf/devstats/issues/35](https://github.com/cncf/devstats/issues/35)
            *   Shoutout to docs for adoption and creating SLOs
            *   What questions do you want answered about the project?
        *   Documenting and improving [communication platforms](https://github.com/kubernetes/community/tree/master/communication)
            *   [Slack guidelines](https://github.com/kubernetes/community/blob/master/communication/slack-guidelines.md)
            *   Working on calendar solutions
        *   Issue Triage and Labels
            *   Proposed and created [Triage Guidelines](https://github.com/kubernetes/community/blob/master/contributors/guide/issue-triage.md) to quickly close issues and clearly define the scope of triage or issues management
            *   Proposed new labels to identify issues that are candidates for close so that issues can be closed quickly manually or with automation with a reasoning for a clear statistic and measurement triage efforts. It is WIP with positive feedbacks from community.
        *   Misc
            *   Roadshow!
            *   F2F this Tuesday @ INDEX
            *   Contributor Summit in Copenhagen
            *   May 1; registration will be on KubeCon/CloudNativeCon site this week
            *   New weekly meeting (from bi-weekly) same day / time (Weds @ 5pUTC)
    *   SIG API Machinery [Daniel Smith](c)
        *   Reminder: SIG-API doesn't own the API (that's SIG-architecture), but rather mechanics in API server, registry and discovery
        *   Design proposal: [https://goo.gl/UbCRuf](https://goo.gl/UbCRuf)
        *   Seeking feedback on webhook mechanism (slack, sig meeting, email to list) and considering some action on it next quarter
        *   Go Contexts: considering addition to go client (recent similar change going into cloud provider).  Should be a very mechanical code change.
*   [ 0:00 ] **Announcements**
    *   Office hours next week!
        *   [https://github.com/kubernetes/community/blob/master/events/office-hours.md](https://github.com/kubernetes/community/blob/master/events/office-hours.md)
    *   Reminder: Contributor Summit will be 1 May, the day before KubeCon/CloudNativeCon EU: [https://github.com/kubernetes/community/tree/master/events/2018/05-contributor-summit](https://github.com/kubernetes/community/tree/master/events/2018/05-contributor-summit)
    *   /lgtm, /approve and the principle of least surprise
        *   [https://github.com/kubernetes/test-infra/issues/6589](https://github.com/kubernetes/test-infra/issues/6589)
        *   Do we all need to use [the exact same code review process](https://github.com/kubernetes/community/blob/master/contributors/guide/owners.md#the-code-review-process)?
        *   How could we make the existing process clearer and better understood?


## February 8, 2018 - ([recording](https://www.youtube.com/watch?v=L1Mk__ddbBg))



*   **Moderators**:  Paris Pittman [SIG ContribEx]
*   **Note Taker**: Josh Berkus [Red Hat/SIG-Release]
*   [ 0:00 ]**  Demo **-- [stork](https://github.com/libopenstorage/stork), storage orchestration runtime for Kubernetes - Dinesh Israni ([disrani@portworx.com](mailto:disrani@portworx.com))
    *   Administrative support for Hyperconverged Storage (that is, Kube storage running in pods).
    *   Has health monitor for storage nodes with automated failover.
    *   Supports the (currently alpha) snapshot provisioner.
    *   Live demo involving Stork storage with MySQL on top.
    *   Currently only supports Portworx storage, but want contributions from other storage drivers.
    *   [Slide](https://docs.google.com/a/portworx.com/presentation/d/e/2PACX-1vQz3SddQVZFvvymniqeOwgUTO9Yb54YqIVLJzL4eM7TU45zjlPOvdyVTjDl7MyuCwDRpKS8lVtUAYiY/pub?start=false&loop=false&delayms=3000)
*   [ 0:14 ]** Release Updates**
    *   **1.10 **[Jaice Singer DuMars ~ Release lead]
        *   Follow along in the official [schedule](https://github.com/kubernetes/sig-release/blob/master/releases/release-1.10/release-1.10.md), we're in week 6 of 12
            *   Please follow the schedule, we will enforce it
        *   **Feature Freeze has passed**, but we are accepting exceptions on a case-by-case basis using the process found [here](https://github.com/kubernetes/features/blob/master/EXCEPTIONS.md)
        *    If you're curious about what is on deck for 1.10 planned work, pay a visit to the [Features Tracking Spreadsheet](https://docs.google.com/spreadsheets/d/17bZrKTk8dOx5nomLrD1-93uBfajK5JS-v1o-nCLJmzE/edit?usp=sharing)
        *   Next week, we will be cutting our first beta release, as well as assembling the release branch
        *   Release team meetings are in the weekly cadence on Mondays at 10am Pacific until February 26th. If you join either the [SIG-Release](https://groups.google.com/forum/#!forum/kubernetes-sig-release) or [kubernetes-milestone-burndown](https://groups.google.com/forum/#!forum/kubernetes-milestone-burndown) groups, you should get an invite
            *   Meetings may or may not include puppies
        *   please remember that **code freeze** begins on** <blink>February 26th</blink>**
            *   Trying to make sure that nobody is surprised by the schedule!
    *   
*   [ 0:18 ] **Graph o' the Week [Tim Pepper; 1.10 release issue triage shadow]**
    *   Release cadence and feature activity:  what is an "issue" or "bug" or "feature" in kubernetes and Github?  We all know GitHub has an "issue" object and there are lots of k8s labels in them, but for feature it's complicated:
        *   project specific labels represent release activities and also "kind/feature"
        *   plus a GitHub object "project"
        *   plus a GitHub object "milestone""
        *   _and_ https://github.com/kubernetes/features
    *   Does the "complicated" show up in the stats?  Are features efficiently created earlier in the cycle and closed as the cycle progresses?
        *   [SIG Issues - 7 Days Moving Average - SIG Release - Kind All - With Release markers](https://k8s.devstats.cncf.io/d/000000031/sig-issues?orgId=1&var-period=d7&var-sig=release&var-kind=All&from=now-1y&to=now)
            *   Doesn't show clear trends, it's a bit chaotic.
        *   [SIG Issues - 7 Days Moving Average - SIG All - Kind feature - With Release markers](https://k8s.devstats.cncf.io/d/000000031/sig-issues?orgId=1&from=now-1y&to=now&var-period=d7&var-sig=All&var-kind=feature)
            *   You can see phases for issues tagged feature.  However, the numbers are increasing towards the end of the release instead of being front-loaded.
            *   For 1.10 we had one SIG which was notably late, trying to figure out if this was a general problem.
        *   You can see roughly 4 phases:  open beginning, feature frozen, code frozen, final stabilization.
        *   And a bonus chart of [features ages](https://k8s.devstats.cncf.io/d/000000002/issues-age?orgId=1&var-period=d7&var-repogroup_name=All&var-repogroup=all&var-sig_name=All&var-kind_name=feature&var-prio_name=All&var-sig=all&var-kind=feature&var-prio=all&from=now-1y&to=now)
        *   We'd like to see people incoming to devstats and looking for answers, trying to figure out how to improve the release.
*   [ 0:23 ] **SIG Updates**
    *   SIG Architecture - [ Brian Grant, co-lead ]
        *   Working on identifying subproject owners
            *   Related to SC decision
            *   Need to map existing code subprojects to SIGs
            *   A lot of the work done by Aaron Crickenberger
            *   Have done a few directories, such as for Workloads APIs
            *   Check SIGs.yaml for what's been identified, make sure that it's correct.
        *   Reviewing architectural issues as they arise
        *   Still working on implementing the KEP
            *   A formalization of the design proposal process
    *   SIG Scalability - [Bob Wise]
        *   [Slides](https://docs.google.com/presentation/d/1QunsQVGe4Ky570dI3hwBPH-BdD65wHkMz-g0S_fPYww/edit#slide=id.p )
        *   Moved meeting to 30min later, biweekly to not conflict with SIG-Arch
        *   They believe that the "[https://docs.google.com/presentation/d/1QunsQVGe4Ky570dI3hwBPH-BdD65wHkMz-g0S_fPYww/edit#slide=id.p](https://docs.google.com/presentation/d/1QunsQVGe4Ky570dI3hwBPH-BdD65wHkMz-g0S_fPYww/edit#slide=id.p)<span style="text-decoration:underline;"> </span>bigger clusters" problem is not interesting to existing members right now, clusters are big (5000 nodes) and stable.  If you want bigger than that, join the SIG.
        *   Mainly about avoiding regressions now.
        *   [They have a new charter](https://github.com/kubernetes/community/pull/1607 ):
        *   Primary work on tooling/monitoring to detect & avoid scaling regressions this year.
        *   They have lots of interest from users in running big clusters, not sure that they're explaining things to those users.
    *   SIG Scheduling - [Bobby Salamat]
        *   A major work item for 1.10 is to move priority and preemption to Beta
            *   New feature in 1.9 alpha.
            *   Very useful for multiple very different workloads
            *   Example: allow production workload to push dev workload aside
        *   Performance improvements
            *   Enable equivalence cache and move it to Beta
            *   Optimize Affinity/Anti-affinity
                *   These are much slower than other predicates (like 10X)
        *   Design a new extension model for scheduler and build a scheduling framework
            *   Currently, the extension model communication is too slow for some plugins.
        *   Three incubators
            *   Kube-arbitrator
                *   "Gang scheduling" so that all pods of a group get scheduled, or not.
                *   Supporting quota for hierarchical namespaces
            *   Cluster capacity tool
                *   Checks if a pod can be scheduled in a certain cluster based on resources
            *   Descheduler
                *   Automated removal of pods to free up resources
*   [ 0:41 ] **Announcements**
    *   Steering Committee Update [Brendan Burns]
        *   [New Repository structure proposal](https://github.com/kubernetes/community/pull/1752)
            *   Sunsetting kubernetes-incubator, won't accept any new projects
            *   3 classes of repositories: (a) associated repos, (b) sig-owned repositories, (c) kubernetes repositories (approved by sig-arch)
            *   This is about new repos going forward, <span style="text-decoration:underline;">not</span> a mandate for existing repos anytime soon.
                *   Except for a few things like Owners files
                *   SIGs who want to move from incubator to sigs repos, stay tuned for details
            *   Question [Matt Farina]: associated repos: who would want one of these?  Why do this?
                *   Brendan: as a prerequisite for submitting a feature to Kubernetes.  CLA is gateway here.  Also, even for external things a consistent process is good and makes it easier for contributors.
                *   Example: Kube-sanity project.  Doesn't belong to a SIG, but all Kube contributors.
    *   Intel Intro [Jose Palafox]
        *   Jose is the program manager for Intel's Kubernetes efforts.
        *   They have a team of 16 engineers on Kube.
        *   Reach out to Jose if you want to collaborate with them.
    *   #shoutouts [Jorge Castro]
        *   Duffie Cooley, Stefan Schimanski, Craig Tracey, Timothy St. Clair, Chuck Ha, Liz Frost, Nikhita Raghunath, Aaron Crickenberger, Ilya Dmitrichenko, Ihor Dvoretski, Ellen Korbes and Tim Pepper
    *   SIG Schedule for this call for the next few months: [SIG Update Schedule](https://docs.google.com/spreadsheets/d/1adztrJ05mQ_cjatYSnvyiy85KjuI6-GuXsRsP-T2R3k/edit#gid=0) (always posted at the top of this document)
        *   Schedule is fixed, please check it.
    *   Meet Our Contributors - first weds of the month
        *   First live-streamed one!
        *   [GH Page](https://github.com/kubernetes/community/blob/master/mentoring/meet-our-contributors.md) ; [Call for volunteers](https://github.com/kubernetes/community/issues/1753)
        *   #meet-our-contributors on slack for questions and code review snips


## February 1, 2018 - [recording](https://www.youtube.com/watch?v=Oj-0l7vdUac)



*   **Moderators**:  Solly Ross  [SIG Autoscaling]
*   **Note Taker**: First Last [Company/SIG]
*   [ 0:00 ]**  Demo **--  [generator-kubegen](https://github.com/sesispla/generator-kubegen), a Kubernetes config generation tool - Sergio Sisternes ([ssistern@everis.com](mailto:ssistern@everis.com))
    *   Link to slides
    *   [Link to repositories](https://github.com/sesispla/generator-kubegen)
    *   Yeoman-based wizard for generating Kubernetes YAML
        *   Asks basic questions, generated Kubernetes object definitions
        *   Can create everything for a basic app, or just individual objects
    *   Questions:
        *   (not a question) Anybody interested in common patterns for should join the App Def WG
*   [ 0:00 ]** Release Updates**
    *   1.10 [Jaice Singer DuMars ~ Release lead/Ihor Dvoretskyi ~ Features lead]
        *   Follow along in the official [schedule](https://github.com/kubernetes/sig-release/blob/master/releases/release-1.10/release-1.10.md), we're in week 5 of 12
        *   Features freeze has passed, and looking good!
        *   [Alphas have been cut successfully](https://github.com/kubernetes/kubernetes/releases/tag/v1.10.0-alpha.2)
        *   Marketing activities (e.g. blog preparation) are beginning
        *   Questions:
            *   (not a question) Lots of issues without labels, please go in and make sure 1.10 issues have labels
            *   How do we go in and try out the latest alphas, e.g. cluster provisioning tool?
                *   Kubeadm supports alphas, as long as they've been pushed to GCS buckets
*   [ 0:00 ] **Graph o' the Week **[Aaron Crickenberger]
    *   Weekly update on data from devstats.k8s.io
    *   Actually a table!  What a twist!
    *   [https://k8s.devstats.cncf.io/dashboard/db/developers-summary?orgId=1](https://k8s.devstats.cncf.io/dashboard/db/developers-summary?orgId=1)
        *   Collects GitHub events (e.g. comments, commits, etc), and associates them with GitHub users
        *   Can break down by releases, time
        *   Not yet broken down by repositories
    *   Questions:
        *   Does it include incubator
            *   Yes, includes all kubernetes-associated repos (kubernetes, kubernetes-incubator, helm, kube-clients)
*   [ 0:00 ] **SIG Updates**
    *   SIG Cluster Ops [Rob Hirschfeld] - [http://bit.ly/k8sclops](http://bit.ly/k8sclops)
        *   New time!  1 hour earlier (12 am Pacific)
            *   Next meeting one hour from now (every two weeks)
        *   Change in format to be more "meetup" style
        *   Looking to bring in more demos & discussions
        *   Specifically want to hear from operators
        *   Fine w/ longer format vendor demos to get feedback
        *   Want to hear more about different deployment patterns
    *   SIG Autoscaling [Solly Ross]
        *   Work continues on VPA (Vertical Pod Autoscaler)
            *   Follow it at [https://github.com/kubernetes/autoscaler](https://github.com/kubernetes/autoscaler)
        *   Investigating minor additions to HPA v2 to improve flexibility with regards to "standalone"/"unassociated" metrics before graduation
        *   Continuing to get feedback on metrics API adapters and identify and improve issues
        *   Meeting times at [https://github.com/kubernetes/community/blob/master/sig-autoscaling/README.md](https://github.com/kubernetes/community/blob/master/sig-autoscaling/README.md)
*   [ 0:00 ] **Announcements**
    *   Steering Committee update
        *   Formalizing subprojects
            *   Part of [proposal from sig architecture](https://docs.google.com/document/d/1FHauGII5LNVM-dZcNfzYZ-6WRs9RoPctQ4bw5dczrkk/edit#heading=h.2nslsje41be1)
            *   Want to make sure everything is owned by some group (SIG or subset thereof), sometimes things are owned by a group within a SIG
            *   [Issue](https://github.com/kubernetes/community/issues/1673)
            *   [Initial implementation PR](https://github.com/kubernetes/community/pull/1674 )
                *   SIG leads should look at PR and sanity check it
            *   This is a non-binding first pass, goal was to make sure every repo had an owning sig, would like help iterating on what subprojects exist and who should own them
            *   Seemed easier to do in one place vs. distributed across all repos
            *   Next steps involve building automation to consume/enforce, making individual repos source of truth via additions to OWNERS files
            *   Examples:
                *   sig-apps owns the "charts" subproject, which corresponds to the "charts" repo
                *   Sig-apps also owns "Workloads API" subproject, corresponding to the API types, clients, etc for the workloads types
                *   Could have a project containing all kubernetes client repos, for instance
            *   Questions
                *   If there's a new subproject that needs a repo, who decides what gets a repo in kubernetes
                    *   Answer: see below :-)
        *   Upcoming
            *   Repositories proposal (aka "the incubator problem")
                *   Moving towards 3 classes of projects (how formal things are)
                    *   Core kubernetes repos (everything in kubernetes/kubernetes + staging, more-or-less), fairly formal, has process, lots of testing, etc
                    *   SIG repos (encourage SIGs to create repos as they see fit, maybe create subproject repos, either stuff outside of core, or stuff that is a prototype before going into core)
                    *   Associated repos (needs CLA bot turned on, code-of-conduct but that's about it)
                *   Doc forthcoming
            *   Expectations of SIG charters and template charters
                *   A checklist/template will come out eventually
                *   Feel free to discuss and submit before then
            *   CNCF graduation
                *   Looking to graduate Kubernetes through the CNCF, making it the first project to do so
    *   Amazon participation update [Bob Wise]
        *   Increasing/ramping up direct involvement
        *   Expect to see more contributions around testing in the short run, AWS experience
        *   Participate in SIG AWS (both EKS and non-EKS)
        *   Can be found on Slack as well, feel free to reach out with feedback, ideas, etc
    *   SIG Testing Commons subproject announced  [Tim St. Clair]
        *   Focus on
            *   tests are written
            *   contributing tests
            *   cleaning up tests
            *   what things are tested
            *   e2e framework
            *   Conformance
        *   Please come participate
    *   Kubernetes Documentation [User Journeys MVP](https://kubernetes.io/docs/home/) launched [Andrew Chen]
        *   Please give SIG Docs for feedback, still adding things later
            *   Can contribute normally (join SIG docs for more information)
        *   New landing page incorporating personas (users, contributors, operators)
        *   Levels of knowledge (foundational, advanced, etc)
        *   Can also just browse docs directly
    *   SIG Arch Announcement [Joe Beda]
        *   Control Plane naming
        *   Feel free to comment offline or on the issue if you have comments
        *   TL;DR: call it the "control plane"
        *   Issue: [https://github.com/kubernetes/website/issues/6525](https://github.com/kubernetes/website/issues/6525)
    *   Contributor Summit for KubeCon/CloudNativeCon EU [Jorge and Paris]
        *   SAVE THE DATE: May 1, 2018
        *   [https://github.com/kubernetes/community/pull/1718](https://github.com/kubernetes/community/pull/1718)
    *   #shoutouts - [Jorge Castro]
        *   New channel in Slack
        *   Someone do something great for the community? Give them a shoutout here and we'll take the time to thank them for their work at the end of each community meeting.
    *   


## January 25, 2018 - ([recording](https://www.youtube.com/watch?v=hAg6aGAG3bs))



*   **Moderators**:  Mario Loria [Meetup Organizer / Liquidweb] (confirmed)
*   **Note Taker(s)**: Jorge Castro [SIG Contributor Experience]
*   [ 0:00 ]**  Demo **-- [kube-toolkit](https://github.com/radu-matei/kube-toolkit) - toolkit for creating gRPC-based CLI and web tools for Kubernetes - Radu Matei ( [radu@radu-matei.com](mailto:radu@radu-matei.com) ) (confirmed)
    *   kube-exec [https://github.com/radu-matei/kube-exec](https://github.com/radu-matei/kube-exec)  - os/exec for remote K8S pods
*   [ 0:00] **INDEX Conference** [Jeff Borek]
    *   [https://developer.ibm.com/indexconf/](https://developer.ibm.com/indexconf/)
        *   Offering space at the Moscone on 20 Feb for communities to hold face to face meetings. Attendance on the 20th is COMPLETELY FREE, but you must register, you'll get 50% off the rest of the conference if you want to stay!
        *   Please contact **jborek@us.ibm.com** if you're interested in claiming some space for your SIG.
        *   "Meet the SIGs" community day will be on the 20th, with Sarah Novotny delivering the keynote.
        *   SIGs interested in participating:
            *   SIG Contributor Experience
            *   SIG Docs
            *   … add yours
*   [ 0:00 ]** Release Updates**
    *   1.10 [Jaice Singer DuMars ~ Release leader]
        *   Follow along in the official [schedule](https://github.com/kubernetes/sig-release/blob/master/releases/release-1.10/release-1.10.md), we're in week 4 of 12
        *   Feature freeze pushed to 1/29 [Ihor Dvoretskyi, Features Lead]
            *   [Feature summary](https://docs.google.com/spreadsheets/d/17bZrKTk8dOx5nomLrD1-93uBfajK5JS-v1o-nCLJmzE/edit#gid=0)
        *   Shooting for Alpha next week on Tuesday
    *   1.9.2 is out (thank you Mehdy!)
    *   1.8.8 - no ETA
*   [ 0:00 ] **Graph o' the Week **[Aaron Crickenberger]
    *   Weekly update on data from devstats.k8s.io
    *   [https://k8s.devstats.cncf.io/dashboard/db/sig-mentions-categories?orgId=1](https://k8s.devstats.cncf.io/dashboard/db/sig-mentions-categories?orgId=1)
    *   Pick a sig, for example SIG-CLI, look into [https://github.com/kubernetes/community/tree/master/sig-cli](https://github.com/kubernetes/community/tree/master/sig-cli) and there are 8 different github teams for SIG-CLI which can be used issue and PR in mentions.  In the devstats then you can view: [https://k8s.devstats.cncf.io/dashboard/db/sig-mentions-categories?orgId=1](https://k8s.devstats.cncf.io/dashboard/db/sig-mentions-categories?orgId=1) se lecting the SIG at the top an you see the usage of the git teams in mentions for that SIG.  Are all the github teams across all the sigs actually in use?  There's a TONNE of them (30 sigs x 8 teams).  Which does a person use to ping the right set of people?  This devstats graph tunnel down is the current best way to find out.  Alternatively need to discuss if these could be simplified into a smaller number of subteams, or do SIG's find the separation into subteams useful?
*   [ 0:00 ] **SIG Updates**
    *   SIG Service Catalog [Paul Morie]  (confirmed)
        *   
    *   SIG CLI [Sean Sullivan] (confirmed)
        *   Moving apply and merge to the server side
        *   Breaking up the monolithic kubectl.
        *   
*   [ 0:00 ] **Announcements**
    *   SIG leads: register to offer intros and deep dives in SIG track at KubeCon/CloudNativeCon Copenhagen (May 2-4): [overview](https://groups.google.com/forum/#!searchin/kubernetes-dev/kohn%7Csort:date/kubernetes-dev/5U-eNRBav2Q/g71MW47ZAgAJ), [signup](https://docs.google.com/forms/d/e/1FAIpQLSedSif6MwGfdI1-Rb33NRjTYwotQtIhNL7-ebtYQoDARPB2Tw/viewform) (1/31 deadline)
    *   [SIG Contributor Experience news: new lead, new meeting](https://groups.google.com/forum/#!topic/kubernetes-dev/65S1Y3IK8PQ)
    *   [Meet Our Contributors ](https://github.com/kubernetes/community/blob/master/mentoring/meet-our-contributors.md)- Feb 7th [Paris]
        *   730a PST/ 3:30 pm UTC & 1pm PST / 9pm UTC
        *   Need contributor volunteers for 1pmPST/9pmUTC session -> [Sign Up [WIP]](https://docs.google.com/spreadsheets/d/1OKc4h-0QLKCbncSloRf_gYklpHYVyNxZmfEM9hlK1xQ/edit?usp=sharing) on m-o-c tab
        *   30 mins AMA; 30 mins live peer code review
        *   Part of larger mentoring initiatives -> [GH repo](https://github.com/kubernetes/community/tree/master/mentoring) ;[ Issue](https://github.com/kubernetes/community/issues/1672)
    *   New SIG-Release lead, Jaice Singer DuMars replacing Phil Wittrock
    *   A new SIG Scheduling lead has been scheduled via Bobby Salamat


## January 18, 2018 - ([recording](https://www.youtube.com/watch?v=x67RK7W-BnM))



*   **Moderators**:  Tim Pepper [VMware/SIG Contrib-Ex]
*   **Note Taker**: Jaice Singer DuMars  [Microsoft/SIG-breakfast]
*   **Chat Transcript**
*   [ 0:00 ]**  Demo **-- Kubernetes on Docker for Mac - Jenny Burcio (jenny@docker.com), Arun Gupta ([arun.gupta@gmail.com](mailto:arun.gupta@gmail.com)) (confirmed)
    *   Link to slides
    *   Link to repositories
    *   Not slated for Linux atm, but Windows might be sooner than later
    *   Uses a CRD where Docker command talks to CRD, creates a new service objects to translate compose files
*   [ 0:00 ]** Release Updates**
    *   1.10 [Jaice Singer DuMars ~ Release lead]
        *   Follow along with the [schedule](https://github.com/kubernetes/sig-release/blob/master/releases/release-1.10/release-1.10.md)
        *   Alpha delayed - requires individual access on Google side so waiting for Caleb's onboarding, also bicycling can be dangerous
        *   [Features](https://github.com/kubernetes/features/issues?utf8=%E2%9C%93&q=is%3Aissue+is%3Aopen+milestone%3Av1.10) update
        *   **Feature freeze is coming, 1/22 [Ihor]**
            *   **Please, ensure that your feature is targeting [1.10 milestone](https://github.com/kubernetes/features/issues?utf8=%E2%9C%93&q=is%3Aissue+is%3Aopen+milestone%3Av1.10) on GitHub**
            *   **Add your feature to the [spreadsheet](https://docs.google.com/spreadsheets/d/17bZrKTk8dOx5nomLrD1-93uBfajK5JS-v1o-nCLJmzE/edit#gid=0)**
    *   1.9.2 Should be today (thanks Mehdy!)
    *   1.8.7 is out (thanks David!)
*   [ 0:00 ] **Graph o' the Week **[Jorge Castro]
    *   Weekly update on data from devstats.k8s.io
    *   [Approvers](https://k8s.devstats.cncf.io/dashboard/db/approvers?orgId=1&var-period=q&var-repogroups=All)
    *   [Approvers Histogram](https://k8s.devstats.cncf.io/dashboard/db/approvers-histogram?orgId=1&var-period_name=Last%20month&var-period=m&var-repogroup_name=All&var-repogroup=all)
*   [ 0:20 ] **SIG Updates**
    *   SIG Docs [Devin Donnelly, Andrew Chen]
        *   K8s Docs Structure
        *   Link to slides (see video for narrative)
        *   
    *   SIG Network [Dan Williams]
        *   1.10 features and work
            *   IPv6 single and dual-stack
            *   IPVS proxy to Beta/GA
            *   Move more kubenet to CNI
            *   Windows proxy and networking
            *   Multiple pod IP addresses (necessary for IPv6 dual-stack)
            *   Continue exploring new Service API
            *   Topology aware ingress and proxies
        *   Continued work on more flexible pod networking through informal Network Plumbing Working Group
    *   SIG Service Catalog [Paul Morie]
        *   Paul is out sick today and sends his regrets, he'll do an update next week
        *   Feel better Paul! +1
*   [0:37] **(Steering Committee)** **Sig Governance Update** [Phillip Wittrock]
    *   Goals: help community to self organize by providing a template charter for SIGs
    *   Complete: reached out to SIG leads with a long form questionnaire to get detailed insight into how various SIGs are structured and function
    *   In progress: developing a template which defines SIG structure and governance within a SIG charter
    *   **Important:** to contribute your insight and experiences to the process, answer these questions [https://goo.gl/Zm81Ly](https://goo.gl/Zm81Ly)
*   [ 0:39 ] **Announcements**
    *   GSoC [Ihor D]
        *   [https://github.com/cncf/soc](https://github.com/cncf/soc); [k8s gh](https://github.com/kubernetes/community/blob/master/mentoring/google-summer-of-code.md)
        *   nikhita has volunteered to drive this program for Kubernetes
    *   SIG Intros & Deep Dives sessions registration at KubeCon/CloudNativeCon & CloudNativeCon will be announced shortly (stay tuned!)
    *   Changes to this meeting's format [Jorge Castro]
        *   SIGs scheduled per cycle instead of adhoc
        *   Demo changes
        *   Note takers (+1)
    *   Meet our Contributors - Ask Us Anything  [Paris]
        *   Feb 7th - 8:30am
        *   [adding link]


## January 11, 2018 - recording



*   **Moderators**: Swarna Podila [SIG Awesome]
*   **Note Taker**: First Last [Company/SIG]
*   **Chat Transcript**
*   [ 0:00 ]**  Demo **--01/11 [KQueen](https://github.com/mirantis/KQueen) Kubernetes cluster manager demo:  Jakub Pavlik ( [jpavlik@mirantis.com](mailto:jpavlik@mirantis.com) )
*   
    *   Tech issues, we'll reschedule this demo at a later date.
*   [ 0:00 ]** Release Updates**
    *   1.10 [Jaice Singer DuMars ~ Release lead]
        *   It's week 2 of 12 of the release (full schedule and some important information is [here](https://github.com/kubernetes/sig-release/blob/master/releases/release-1.10/release-1.10.md))
        *   The [release team](http://bit.ly/k8s110-team) is fully staffed!
        *   we have added a role guide for the [Communications Coordinator](https://github.com/kubernetes/sig-release/blob/master/release-process-documentation/release-team-guides/communications_coordinator.md)
        *   [PR](https://github.com/kubernetes/sig-release/pull/69) out for the Release Leader role documentation
        *   The next major deadline is Feature Freeze on January 22nd [Ihor]
            *   If you're targeting a feature for 1.10 - please, ensure that it's updated in the [Features repo under 1.10 Milestone](https://github.com/kubernetes/features/issues?q=is%3Aopen+is%3Aissue+milestone%3Av1.10);
                *   Ihor will start grooming the 1.10 backlog in the features repo on Monday; now it's a call for feature owners to update the features in the repo
            *   [Features tracking spreadsheet is ready for your contributions](https://docs.google.com/spreadsheets/d/17bZrKTk8dOx5nomLrD1-93uBfajK5JS-v1o-nCLJmzE/edit#gid=0) (shared r/w with kubernetes-dev);
    *   1.9.2 - due Thursday
    *   1.8.7 - Early next week, with PR deadline on Friday
*   [ 0:00 ] **Graph o' the Week **[Aaron Crickenberger / @spiffxp]
    *   Weekly update on data from devstats.k8s.io
    *   Today's graph: [https://k8s.devstats.cncf.io/dashboard/db/need-rebase-prs](https://k8s.devstats.cncf.io/dashboard/db/need-rebase-prs)
        *   [What are the repo groups?](https://github.com/cncf/devstats/blob/master/scripts/kubernetes/repo_groups.sql)
        *   IMO this graph is more useful than [All Need Rebase PR's](https://k8s.devstats.cncf.io/dashboard/db/all-need-rebase-prs?orgId=1)
        *   [I'm asking that we remove the "All… " dashboards](https://github.com/cncf/devstats/issues/41) and merge their graphs into the more detailed dashboards
    *   [Github search: all open kubernetes PR's](https://github.com/pulls?utf8=%E2%9C%93&q=is%3Apr+is%3Aopen+user%3Akubernetes)
    *   [Github search: all open kubernetes PR's with label:needs-rebase](https://github.com/pulls?utf8=%E2%9C%93&q=is%3Apr+is%3Aopen+user%3Akubernetes+label%3Aneeds-rebase+)
    *   [The munger responsible for applying needs-rebase](https://github.com/kubernetes/test-infra/blob/master/mungegithub/mungers/needs_rebase.go)
    *   [We are migrating mungers out of github](https://github.com/kubernetes/test-infra/issues/3331)
    *   [We've added a prow plugin for needs-rebase](https://github.com/kubernetes/test-infra/pull/6121)
    *   [We plan on migrating mungegithub to github.com/kubernetes-retired](https://github.com/kubernetes/test-infra/issues/6104)
*   [ 0:00 ] **SIG Updates ([List of SIGs](https://github.com/kubernetes/community/blob/master/sig-list.md))**
    *   SIG Azure [Jaice Singer DuMars] (confirmed)
        *   Moving Microsoft upstream planning 100% to the SIG
        *   Cloud provider breakout work continues
        *   1.10 planning was yesterday, we're going to try out the new features process
    *   SIG Node[Dawn Chen] (confirmed)
        *   1.10 planning
            *   [https://docs.google.com/document/d/15F3nWPPG3keP0pzxgucPjA7UBj3C31VsFElO7KkDU04/edit?userstoinvite=doug.maceachern@gmail.com&ts=5a579cac](https://docs.google.com/document/d/15F3nWPPG3keP0pzxgucPjA7UBj3C31VsFElO7KkDU04/edit?userstoinvite=doug.maceachern@gmail.com&ts=5a579cac)
*   What's the latest on combining provider work into a single SIG vs breaking it out?
    *   Please see [https://github.com/kubernetes/community/tree/master/wg-cloud-provider](https://github.com/kubernetes/community/tree/master/wg-cloud-provider)
    *   ([Notes](https://docs.google.com/document/d/1OZE-ub-v6B8y-GuaWejL-vU_f9jsjBbrim4LtTfxssw/edit#heading=h.w7i4ksrweimp))
*   [ 0:00 ] **Announcements**
    *   The final call for [KubeCon/CloudNativeCon EU 2018](https://events.linuxfoundation.org/events/kubecon-cloudnativecon-europe-2018/) CFP submissions! [Ihor]
        *   CFP form will close on Jan 12 at 11:59 PT
        *   Early pricing deadline is the same ^
    *   IBM is organizing [INDEX](https://developer.ibm.com/indexconf/) conference in San Francisco late February. There's a free space for f2f meetings at Moscone Center on February 20. [Ihor]
        *   If your SIG (or a different group of contributors) would like to organize a f2f meeting on February 20 in SF, please reach Ihor Dvoretskyi about the details
    *   Setting up merge automation for all github.com/kubernetes repos [Aaron Crickenberger / @spiffxp]
        *   [https://github.com/kubernetes/test-infra/issues/6227](https://github.com/kubernetes/test-infra/issues/6227)
        *   [https://groups.google.com/d/msg/kubernetes-dev/h-0hGFJ8x1E/g4UuGr5zDAAJ](https://groups.google.com/d/msg/kubernetes-dev/h-0hGFJ8x1E/g4UuGr5zDAAJ)
    *   k8s Office hours this Wednesday! Ping Jorge Castro (@jorge on slack) if you want to volunteer [https://git.k8s.io/community/events/office-hours.md](https://git.k8s.io/community/events/office-hours.md)


## January 04, 2018 - ([recording](https://www.youtube.com/watch?v=fdXS-mSX7F8))



*   **Moderators**:  Chris Short []
*   **Note Taker**: Jaice Singer DuMars [SIG-kiwi]
*   **Chat Transcript**
*   [ 0:02 ]**  Demo [ 10 minutes ] **--  [kube-arbitrator](https://github.com/kubernetes-incubator/kube-arbitrator) demo: Klaus Ma (@k82cn, [madaxa@cn.ibm.com](mailto:madaxa@cn.ibm.com) / @jinzhejz, [jinzhej@cn.ibm.com](mailto:jinzhej@cn.ibm.com) )
    *   [design doc](https://docs.google.com/document/d/1-H2hnZap7gQivcSU-9j4ZrJ8wE_WwcfOkTeAGjzUyLA/edit#heading=h.uedqgav5zc53)
    *   [https://github.com/kubernetes-incubator/kube-arbitrator](https://github.com/kubernetes-incubator/kube-arbitrator)
*   [ 0:10 ]** Release Updates [ 5 minutes ]**
    *   1.10 [Jaice Singer DuMars ~ Release lead]
        *   [Schedule](https://github.com/kubernetes/sig-release/blob/master/releases/release-1.10/release-1.10.md) is set ([http://bit.ly/k8s110-schedule](http://bit.ly/k8s110-schedule))
        *   [Team](https://github.com/kubernetes/sig-release/blob/master/releases/release-1.10/release_team.md) is forming ([http://bit.ly/k8s110-team](http://bit.ly/k8s110-team))
            *   **Need** CI Signal, Bug Triage shadow, and Branch Manager (Google only)
        *   Key dates:
            *   **release day**, Wednesday, March 21st
            *   **Feature freeze** is Monday, January 22nd
            *   **Code freeze** begins Monday February 26th and ends Wednesday, March 14th
            *   **Docs** must be completed and reviewed by Friday, March 9th
*   [ 0:15 ] **Graph o' the Week **[Aaron Crickenberger / @spiffxp] [** 5 minutes** ]
    *   Weekly update on data from devstats.k8s.io
    *   [https://k8s.devstats.cncf.io/dashboard/db/bot-commands](https://k8s.devstats.cncf.io/dashboard/db/bot-commands)
    *   Manually updated list of bot commands: [https://go.k8s.io/bot-commands](https://go.k8s.io/bot-commands)
    *   Auto-generated help page with examples: [https://prow.k8s.io/plugin-help.html](https://prow.k8s.io/plugin-help.html)
    *   Mapping of repository groups: [https://github.com/cncf/devstats/blob/master/scripts/kubernetes/repo_groups.sql](https://github.com/cncf/devstats/blob/master/scripts/kubernetes/repo_groups.sql)
    *   Most used command: /cc (to be used for reviews)
    *   Recent command that's seen growth: /hold
    *   Nit: /lgtm can count as /approve, devstats can't tell the difference
    *   Some recent spikes of interest: /priority, /lifecycle, /close
*   [ 0:20 ] **SIG Updates [ 5 minutes ] **
    *   SIG Instrumentation [Frederic Branczyk, CoreOS ]
        *   Presentation link
        *   multiple kube-state-metrics releases - many new metrics, stability, and features ~ making metrics actionable
        *   1.0 has been released
        *   Core and Custom metrics APIs have been promoted to beta
            *   Formally defined APIs for metrics, across all workload types
            *   Can be anything arbitrary your system can capture
            *   Beta as of 1.8, just remember that they are aggregated API servers, so the implementations may be in flux or different
            *   Custom Prometheus Adapter (need GH link from DirectXMan12/k8s-prometheus-adapter
            *   Can autoscale on arbitrary metrics collected in Prometheus
        *   Removing heapster dependencies
            *   Heapster maintainers have come into the SIG
        *   Meet every thursday 6PM European time
*   [ 0:26 ] **Announcements [ 5 minutes ] **
    *   [tstclair] - Socialize proposal to move from 4 -> 3 release cycles a year to reduce
        *   [https://groups.google.com/forum/#!topic/kubernetes-dev/nvEMOYKF8Kk](https://groups.google.com/forum/#!topic/kubernetes-dev/nvEMOYKF8Kk)
    *   Check with your cloud providers wrt. Meltdown/Spectre: [https://meltdownattack.com/](https://meltdownattack.com/)
        *   (Too much info to cover here)
    *   Office Hours is back, 17 Jan! [https://git.k8s.io/community/events/office-hours.md](https://git.k8s.io/community/events/office-hours.md)
    *   New slack guidelines -> [https://github.com/kubernetes/community/blob/master/communication/slack-guidelines.md](https://github.com/kubernetes/community/blob/master/communication/slack-guidelines.md)
    *   Group mentoring cohort #1 kicked off today! \o/
        *   Current members to reviewers
        *   [https://goo.gl/forms/nAWxAWpVBdNQbyWy1](https://goo.gl/forms/nAWxAWpVBdNQbyWy1)
        *   Contributor Office Hours (will be renamed) coming at the end of the month - date TBA; doc to be created in the mentoring docs folder -> [https://github.com/kubernetes/community/tree/master/mentoring](https://github.com/kubernetes/community/tree/master/mentoring)
    *   KubeCon/CloudNativeCon EU 2018 CFP closes on Jan 12 - [https://events.linuxfoundation.org/events/kubecon-cloudnativecon-europe-2018/](https://events.linuxfoundation.org/events/kubecon-cloudnativecon-europe-2018/)
*   [ 0:00 ] **1.9 Release Retrospective Part 2 [ 30 minutes ] **
    *   NOTE: Please add your retro items to the document below the [ part 2 ] line
    *   The [retro doc](http://bit.ly/kube19retro)
    *   Part One [recording](https://youtu.be/oagLX--fdDs)
