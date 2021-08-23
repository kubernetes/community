### October 15, 2020 ([recording](https://youtu.be/h7jth3Js-Vg))



*   **Moderators:** POP (@danpopsd)
*   **Note Taker:** Cody Crudgington (@thecrudge1)
*   **SIG-Instrumentation** (Han Kang [https://github.com/logicalhan](https://github.com/logicalhan)) (Instrumentation would like to first as they have a hard stop)
    *   **Slides/Bullets** [https://docs.google.com/presentation/d/1kEG6wnnCdRzpfuDl8GRPzCtKnjhBvh84ESHfif86KZ0/edit#slide=id.g401c104a3c_0_0](https://docs.google.com/presentation/d/1kEG6wnnCdRzpfuDl8GRPzCtKnjhBvh84ESHfif86KZ0/edit#slide=id.g401c104a3c_0_0)
    *   3 pillars of observability for sig-instrumentation: Metrics, logs, distributed tracing
    *   KEP Status:
        *   V2 events api is now graduated.
        *   Beta  
            *   Metrics Stability Framework. SLO type metrics
        *   Alpha  
            *   Efforts being made for structured logging. Went through PRR, links in slides
        *   Road to Alpha
            *   Dynamic cardinality enforcement
            *   Distributed tracing
            *   Log sanitation
                *   Dynamic log sanitation
                *   Sanitization via static analysis
    *   Plans for upcoming cycles
        *   Metrics stability to GA
        *   Kube-state-metrics releases
        *   Structured logging to beta
        *   Alpha distributed tracing
        *   Alpha log sanitation
        *   Alpha dynamic cardinality enforcement
        *   Instrumentation tools
            *   Promq
            *   Metrics-server
    *   NEED NEW CONTRIBUTORS! PLEASE HELP OUT!
*   **Steering Committee** (Paris Pittman [https://github.com/parispittman](https://github.com/parispittman)) (Confirmed)
    *   [https://github.com/kubernetes/steering](https://github.com/kubernetes/steering)
    *   Thank you so much to our contributors and community for your kindness and camaraderie during trying times
    *   We had an election! Welcome - liggit@ (Jordan), mrbobbytables@ (Bob), and wb dims@(Davanum). spiffxp@ and lachie@ thanks for serving! Every year we are so impressed with how many folks raise their hands for this type of contribution - it’s a true testament to our chop wood carry water community.
        *   Thank you to josh, jaice, and ihor for election officer duties!!!
    *   Actually - there were TWO elections!
        *   Many thank yous to the new code of conduct committee! Karen Chu, Tim Pepper, and Celeste Hogan join Aeva and Tasha. Thanks to Jaice and Carolyn for their service!
    *   New:
        *   Unconscious Bias training for Chairs, TLs, and WG organizers - almost done!
            *   https://github.com/kubernetes/community/issues/4959
        *   Annual report process with working groups is underway.
            *   wg-security audit -> sig security
                *   HONK!
            *   Look out for reports to get filed into k/community groups areas
            *   SIGs are next year
        *   Speaking of working groups, a new one was formed: wg-reliability
        *   Still doing our funding thing :) https://github.com/kubernetes/funding
    *   On the schedule:
        *   term limits for steering?
        *   Untangling various uses of the term ‘member’,
    *   Our next public meeting: November 2
    *   Please talk to us! #steering-committee, [steering@kubernetes.io](mailto:steering@kubernetes.io), kubernetes/steering with issues
*   **Release Updates** (Nabarun Pal [https://github.com/palnabarun](https://github.com/palnabarun))  (CONFIRMED)
    *   **Timeline**
        *   We are in Week 5 of the release (almost halfway through)
        *   [Code Freeze](https://github.com/kubernetes/sig-release/blob/master/releases/release_phases.md#code-freeze) begins on November 12.
        *   [Test Freeze](https://github.com/kubernetes/sig-release/blob/master/releases/release_phases.md#test-freeze) is on November 23.
        *   Docs PRs should be ready to be reviewed by November 23.
        *   Docs Freeze is on November 30.
        *   v1.20.0 is planned to be shipped on December 8.
        *   You can see the full 1.20 release calendar [here](https://github.com/kubernetes/sig-release/tree/master/releases/release-1.20)
    *   **New Meeting for Europe + APAC contributors**
        *   We now conduct a parallel Release Team meeting for Europe + APAC focused contributors.
        *   The first meeting was this week.
        *   Monday 12:00 UTC/17:30 IST/05:00 PDT
    *   **Patch Releases**
        *   v1.19.3 were released yesterday on October 14
        *   v1.17.13, v1.18.10 are to be released soon
    *   **Current Release Branch cuts**
        *   1.20.0-alpha.2 was released on October 13
        *   1.20.0-alpha.3, 1.20.0-beta.0 and so on will be released later this month
    *   **Enhancements**
        *   Enhancements Freeze in effect since October 6
        *   Enhancements Being Tracked: 56
            *   Alpha: 23
            *   Beta: 17
            *   Stable: 16
        *   There are 5 exceptions raised as of today for the enhancements freeze.
    *   **Shadow Program**
        *   Shadow Orientation is complete
        *   We have 29 shadows in the team this cycle!!!
    *   We are always welcoming new contributors in SIG Release. Ping us at #sig-release on the Kubernetes Slack.
    *   [Stephen] **[Discussing/modifying the kubernetes/kubernetes release cadence](https://github.com/kubernetes/sig-release/issues/1290)**
        *   k-dev email: [https://groups.google.com/g/kubernetes-dev/c/luW11nhLds4](https://groups.google.com/g/kubernetes-dev/c/luW11nhLds4)
        *   1.19 Release was extended release. Feedback on minor release cadence. Issue opened in sig-release repo. Please provide feedback!
            *   [https://github.com/kubernetes/sig-release/issues/1290](https://github.com/kubernetes/sig-release/issues/1290)
*   **Sig Updates**
    *   **Storage** (Xing Yang [https://github.com/xing-yang](https://github.com/xing-yang)) (CONFIRMED)
        *   **Slides/Bullets: **[https://docs.google.com/presentation/d/1uzS6Q1OmttV-0hzlu_Eublx24nRmWiJQ__J65Sr3hIs/edit?usp=sharing](https://docs.google.com/presentation/d/1uzS6Q1OmttV-0hzlu_Eublx24nRmWiJQ__J65Sr3hIs/edit?usp=sharing)
        *   TL;DR
            *   1.19 Release updates
                *   Beta
                    *   Azure disk and vsphere CSI Migration
                    *   CSI Windows
                    *   Immutable secrets and configmaps
                    *   Volume expansion will stay in beta with improvements
                        *   Offline support detection
                    *   Volume snapshot will state in beta with improvements
                *   Alpha
                    *   CSI storage capacity tracking
                        *   Enables local volume dynamic provisioning
                    *   Generic ephemeral volumes
                    *   CSI volume health
                    *   CSIDrive policy fo fsgroup
            *   Plans for upcoming cycles
                *   CSI volume expansion - Target GA 1.21
                    *   Planning to deprecate online/offline volume expansion plugin
                *   CSI vol.,ume snapshot - Target GA 1.20
                *   Non-recursive volume ownership (fsgroup) - Target beta in 1.20
                *   CSI driver policy group
            *   New features
                *   Pass pod service account token to CSI Target 1.20
            *   In design / proto
                *   Recovery from volume expansion failure
                *   Non recursive selinux and csi drive config
                *   Ceph csi migration
                *   Volume groups
                    *   Snapshot consistency, failure domain spreading
                *   Generic data populator
                *   Container Object Storage Interface
            *   Cross sig WG’s - Data protection WG sponsored by sig-storage and sig-apps
                *   How to quiesce / unquiesce for application snapshots
                    *   containerNotifier
            *   Sig-Apps
                *   Volume expansion for statefulsets
                *   Pvc cleanup on statefulset deletion / scale down
            *   Sig-Scheduling
                *   See slides,
    *   **Service Catalog** (Jonathan Berkhahn [https://github.com/jberkhahn](https://github.com/jberkhahn)**) **(CONFIRMED)
        *   **Slides/Bullets**
            *   **General updates**
                *   Maintenance mode the last 6 months.
                *   No new major features atm
            *   **Helm chart repo moved**
                *   What is the repo link?
            *   **Mszostok stepping down as chair**
        *   Updates coming for spec
*   ** Open discussions **
    *   Shoutout for inclusive training!
*   ** Announcements **
    *   Follow @k8scontributors on twitter
    *   [KubeCon NA 2020 Schedule](https://events.linuxfoundation.org/kubecon-cloudnativecon-north-america/program/schedule/)
    *   Next Community Meeting  - January 21st
    *   
*    **  Shoutouts this month **
    *   @[RobKielty](https://app.slack.com/team/UBEGAFN2V) Shoutout to [@Lauri Apple](https://kubernetes.slack.com/team/U011C07244F) who is doing fantastic work to manage the collection of CI Signal experience and document formally all of the work we do in a goal-directed way.  When we're done, this will be super useful for both new CI Signalers and anyone interested in navigating the CI tooling to find, report and eliminate non-deterministic test results.  Thank you [@Lauri Apple](https://kubernetes.slack.com/team/U011C07244F)!
    *   [Lauri Apple](https://app.slack.com/team/U011C07244F)  [1:02 PM](https://kubernetes.slack.com/archives/C92G08FGD/p1602003731030900) Shoutout in response to [@RobKielty](https://kubernetes.slack.com/team/UBEGAFN2V) for his fantastic work onboarding our 1.20 CI Signal shadows, creating a mechanism for automating CI Signal weekly updates (toil reduction FTW) ([https://github.com/kubernetes/release/issues/1546](https://github.com/kubernetes/release/issues/1546)) and for providing us with many great complexity-distilling analogies that explain what we do! (edited)
    *   @pop [@jeefy](https://kubernetes.slack.com/team/U5MCFK468) a shoutout for stepping up with the stream help on today's steering committee call.?  good job!
    *   [Mrbobbytables](https://app.slack.com/team/U511ZSKHD) shoutout to [@oikiki](https://kubernetes.slack.com/team/U9HFFRFT2) for wrangling Enhancements <3 There are FIFTY-TWO tracked enhancements
    *   @[mrbobbytables](https://app.slack.com/team/U511ZSKHD) shoutout to [@jberkus](https://kubernetes.slack.com/team/U0UKM380M) [@jaice](https://kubernetes.slack.com/team/U0YJS6LHL) and [@ihor.dvoretskyi](https://kubernetes.slack.com/team/U0CBHE6GM) for serving the community by running the 2020 steering election.

**The esteemed Mr. Augustus needs leads to provide any content for Kubecon! Please have it to him by the end of the day!**


### September 17, 2020 - ([recording](https://youtu.be/6Wn_dIEg0E8))


*   **Moderators**:  Eddie Zaneski [AWS/SIG-CLI]
*   **Note Taker**: Josh Berkus [Red Hat/Contribex]
    *   Subscribe to [this thread](https://discuss.kubernetes.io/t/kubernetes-weekly-community-meeting-notes/35) to get these notes in your inbox
*   [ 0:00 ]**  Demo **-- Warning: Helpful Warnings Ahead [Jordan Liggitt]
    *   [https://kubernetes.io/blog/2020/09/03/warnings/](https://kubernetes.io/blog/2020/09/03/warnings/)
    *   It's hard to keep up with kubernetes
    *   Release notes are long
    *   Start with a manifest file -- example v1beta1/RBAC api file
        *   Is that deprecated yet?
        *   When is it going away?
    *   Added ability of the server to return warnings to users as of 1.19
    *   Server sends back headers to the clients (update your client)
    *   This works for any kubectl command
    *   In 1.19, you'll get warnings about all deprecated APIs
    *   In the future, we'll give warnings about other things
        *   Even problematic configurations, like 4millibytes of memory
    *   Also: metrics about upgradability of cluster.  Like, are people actually *using* the things that are deprecated?
        *   Check # of calls to deprecated API
    *   output is to STDERR.
    *   Kubectl does not have an option to silence warnings, but client-go does.
    *   Details and examples in the [blog post](https://kubernetes.io/blog/2020/09/03/warnings/)
*   [ 0:10 ]** Release Updates**
    *   Current Release Development Cycle  [Jeremy Rickard | @jerickar]
    *   Patch Release Updates [https://git.k8s.io/sig-release/releases/patch-releases.md](https://git.k8s.io/sig-release/releases/patch-releases.md)
        *   v1.19.2 planned for Wednesday 2020-09-16
    *   [v1.20 release timeline](https://github.com/kubernetes/sig-release/blob/master/releases/release-1.20/README.md#timeline)
    *   Kicked off the release this week, it is now Week 1
    *   Enhancements freeze is Oct 6
    *   Pushed the docs deadline out a little bit because of conflict with Kubecon (Nov 30)
    *   1.20 will be 13 weeks
    *   Planned release date: Dec 8
    *   Shadows are mostly selected
*   [ 0:00 ] **SIG Updates**
    *   SIG Windows [Mark Rossetti + Michael Michael]
        *   [Slides](https://docs.google.com/presentation/d/1FkS6FIes9opMR1B4qXyVzwGfUBnX0A-sh27Kcnyi-e4/edit#slide=id.g97ab5232db_0_0)
        *   Did a bunch of work on containerd in the last cycle
        *   Aligning windows with the CRI
        *   CNI is now supported with containerd
        *   containerd has Windows features that you don't get with Docker, like graceful shutdown and namespace mapping
        *   now working on HyperV isolated containers
        *   Also developed CSI proxy, windows service to support CSI
            *   Currently supports Azure disks, GCE
            *   Currently in beta
            *   Adding more providers for 1.20 release
        *   Networking work
            *   Support for endpointslices
            *   sticky sessions
            *   dual-stack support
            *   Antrea and Calico CNIs are supported
            *   Need to promote DSR kube-proxy to stable
        *   Shared tech matrix of windows and kubernetes versions
        *   More plans
            *   More CRI-containerd work
            *   GPU support
            *   Add more storage providers, add Valero support for backup
            *   Dump the in-tree Windows extensions
            *   Add privileged container support, which is a big deal, KEP awaiting approval
        *   If you're reviewing PRs and KEPS that might impact Windows, please ping the SIG
    *   SIG Multicluster [Jeremy Olmsted-Thompson + Paul Morie]
        *   [Slides](https://docs.google.com/presentation/d/1WKtsiSQn0sQ3IaSql4pGnH8Qt9ibBjHXMFBYZBhWeII)
        *   Name for a group of clusters: ClusterSet
        *   Are writing guidelines, starting with Namespace Sameness Doctrine
        *   Next: defining what a ClusterSet means
            *   Membership & registration
            *   Workload placement
        *   Move APIs to beta
        *   Kubefed is working on adding "pull reconciliation"
            *   Currently Kubefed does push
            *   Team is working on multicluster reconciler
        *   [KEP 1645](https://github.com/kubernetes/enhancements/tree/master/keps/sig-multicluster/1645-multi-cluster-services-api)
            *   Multicluster DNS
            *   Define Cluster ID
            *   Please give feedback
        *   WorkAPI
            *   Different API type than Kubefed
            *   Fed has you create each resource in each cluster
            *   WorkAPI has a single "work" that deploys resources across the multicluster, with status information
            *   But scheduling is still undefined
            *   Initial KEP being drafted, early prototypes
        *   SIG-MC needs input!
            *   What else should be in a CLusterSet?
            *   Have you built an MC thing that you like?  Share?
            *   Need help from SIG-network, SIG-apps
    *   Code of Conduct Committee [Aeva Black]
        *   [Slides](https://docs.google.com/presentation/d/1yDhuK0DOt2sespWDpWEf-X7zAIyRsIT7pPhbxfFIlLk/edit?usp=sharing)
        *   Recently had elections, have 5 members
        *   Folks don't know what the COCC does
        *   Distinct from SC/Contribex
        *   We're here as a resource to help out
            *   We handle incidents
            *   But, you can also ask for info, doesn't need to be an incident
        *   Focused on harm reduction
        *   Were CoCC for Kubecon SD, tested model there for helping with events
        *   Also working on transparency reports
        *   Did some conflict resolution, including between SIGs
        *   Have offered coaching for problematic content
        *   Working on:
            *   SLA for incident response
            *   Report/Triage docs
            *   Formalize relationship with all admins
*   [ 0:00 ] ** Open discussions **
    *   
*   [ 0:00 ] ** Announcements **
    *   Kubernetes  Steering Committee Elections!! VOTE!! [Josh]
        *   Didn't get your ballot?  [Request a replacement](https://www.surveymonkey.com/r/kubernetes-sc-2020-ballot)
    *   We have passcodes for zoom now!! (zoom’s rules!) 77777 - pass it on.
    *   PK will host next month
*   [ 0:00 ] **  Shoutouts this month **
    *   See someone going above and beyond? Mention them in [#shoutouts](https://kubernetes.slack.com/archives/C92G08FGD) channel so that we can recognize their work in front of the community.
    *   [@sraghunathan](https://kubernetes.slack.com/team/UC8U2V3BM) shoutout to [@karenb](https://kubernetes.slack.com/team/UCLQ9GKSP) for her support with 1.19 k/website release process so far. She is dependable and goes an extra mile to help and constantly on the lookout for improvement. Thanks for your help [@karenb](https://kubernetes.slack.com/team/UCLQ9GKSP)
    *   [@liggitt](https://kubernetes.slack.com/team/U0BGPQ6DS) shoutout to [@knight42](https://kubernetes.slack.com/team/UL9NDRCBV) for tireless work tracking down and fixing test flakes (and real bugs the flakes were trying to alert us about!)
    *   [@aojea](https://kubernetes.slack.com/team/U7CK9A960) big, huge shoutout to [@liggitt](https://kubernetes.slack.com/team/U0BGPQ6DS) , not only due to his work, also for [sharing his knowledge and expertise](https://kubernetes.slack.com/archives/C09QZ4DQB/p1598397716042600)
    *   [@inductor](https://kubernetes.slack.com/team/UD179SDGA) shout out to [@kakts](https://kubernetes.slack.com/team/UJN56KH0A) and [@oke-py](https://kubernetes.slack.com/team/UEQ21UVMF) for creating good first issues in Japanese localisation
    *   [@spiffxp](https://kubernetes.slack.com/team/U09R2FL93) Big big shoutout to [@aojea](https://kubernetes.slack.com/team/U7CK9A960) who helped track down DNS as the culprit of [recent kubernetes/kubernetes CI woes](https://github.com/kubernetes/test-infra/issues/19080#issuecomment-685141853)
    *   [@Lauri Apple](https://kubernetes.slack.com/team/U011C07244F) Shoutout to [@xmudrii](https://kubernetes.slack.com/team/U4Q2TNGVD) for driving efforts to firm up the roles and responsibilities for our Release Manager Associates!
    *   [@jerickar](https://kubernetes.slack.com/team/U72ESU398) K8s 1.19 retro is happening, shout out to [@Lauri Apple](https://kubernetes.slack.com/team/U011C07244F) for digging into previous retro to-do items and  prioritizing things that had fallen by the wayside and helping to make releases better going forward!
    *   [@sftim](https://kubernetes.slack.com/team/UGBUYDQR2) Shoutout to [@zacharysarah](https://kubernetes.slack.com/team/U5WQMKJEA) for a wealth of positive, valuable contributions to SIG Docs
    *   [@ehashman](https://kubernetes.slack.com/team/U9X5ARSLS) Shoutout to [@jvanz](https://kubernetes.slack.com/team/U1F5Y59GB) for this excellent [first PR](https://github.com/kubernetes/kubernetes/pull/92878). I helped backport the fix to 1.17-1.19. Cross-checking this against the production clusters I support, I discovered this bug is producing a full 2% of our production logs, so the fix will save a lot of wasted volume! :)
    *   [@dholbach](https://kubernetes.slack.com/team/U8NHWBFDW) [@SomtochiAma](https://kubernetes.slack.com/team/USPCDTELF) has done great work in her [GSoC internship this year](https://kubernetes.io/blog/2020/09/16/gsoc20-building-operators-for-cluster-addons/)
    *   [@SomtochiAma](https://kubernetes.slack.com/team/USPCDTELF) Shout out to my awesome mentors throughout the GSoC internship [@justinsb](https://kubernetes.slack.com/team/U0A6A01FG) [@stealthybox](https://kubernetes.slack.com/team/U3W18JQLF) and the [#cluster-addons](https://kubernetes.slack.com/archives/CH21PG8SW) community! They were really amazing


### August 20, 2020 - ([recording](https://youtu.be/oDL3Kp5-9eM))


*   **Moderators**:  Jenny Warnke [SIG Contribex]
*   **Note Taker**: Jorge Castro [SIG Contribex]
    *   Subscribe to [this thread](https://discuss.kubernetes.io/t/kubernetes-weekly-community-meeting-notes/35) to get these notes in your inbox
*   [ 1:30 ]** Release Updates**
    *   Current Release Development Cycle  [Jeremy Rickard]
        *   1.19 Is Coming! Aspirationally releasing on Aug 25th
        *   Release Retrospective on Aug 27th
        *   CI Signal Updates/Improvements
        *   Get Ready for 1.20. Question from the audience, is 1.20 a specific stability release? Augustus: No, stability is a property of all releases, we won’t do a “stability release” for only one series.
            *   1.20 Release Team Formation Issue: [https://github.com/kubernetes/sig-release/issues/1185](https://github.com/kubernetes/sig-release/issues/1185)
    *   Patch Release Updates [https://git.k8s.io/sig-release/releases/patch-releases.md](https://git.k8s.io/sig-release/releases/patch-releases.md)
        *   1.18.8, 1.17.11, 1.16.14 released August 13
        *   Next Patches should be Sept 16th
*   [ 5:50 ] **SIG Updates**
    *   SIG Scalability [Matt Matejczyk]
        *   [Slides](https://docs.google.com/presentation/d/1-NOmFvwBqIGkZNQPe4R44hLvv0BuxbZy3DLleipKPHg/edit?usp=sharing)
    *   SIG Autoscaling [Marcin Wielgus]
        *   [Slides](https://docs.google.com/presentation/d/16tPynx13cFk2vmCap5wF8nnhvX5lAK3fu2UHZQR0WUU/edit#slide=id.g401c104a3c_0_0)
    *   SIG Scheduling [Wei Huang]
        *   [Slides](https://docs.google.com/presentation/d/1H27SDMqkzq8zCRveWWtK5g9hCAomKbrzTTVZ5r4h6Xo/edit)
    *   Naming WG [Celeste Horgan]
        *   [Slides](https://docs.google.com/presentation/d/1ImUMm4U9LSVPDcp37gyhG5ztO4zmWc7Go7eOs_kAHxc/edit?usp=sharing)
*   [ 38:00 ] ** Announcements **
    *   KubeCon + CloudNativeCon is wrapping up, thanks everyone for participating!
    *   You can host this meeting, ping @castrojo on slack if you’d like to volunteer!
    *   Follow us on [@k8scontributors](https://twitter.com/k8scontributors)
*   **  Shoutouts this month **
    *   See someone going above and beyond? Mention them in  [#shoutouts](https://kubernetes.slack.com/archives/C92G08FGD) so that we can recognize their work in front of the community.
    *   **[bentheelder](https://app.slack.com/team/U1P7T516X)** shoutout to [@aojea](https://kubernetes.slack.com/team/U7CK9A960) for some **_deep_** debugging and the biggest little flakiness fix in a long time. 3 lines of code and CI went from red  => green, [https://github.com/containernetworking/plugins/pull/509](https://github.com/containernetworking/plugins/pull/509), [https://prow.k8s.io/?job=pull-kubernetes-e2e-kind](https://prow.k8s.io/?job=pull-kubernetes-e2e-kind)
    *   **[Celeste Horgan](https://app.slack.com/team/USF2T2W78)** Shoutout to [@bentheelder](https://kubernetes.slack.com/team/U1P7T516X) for jumping in on docs container issues in such a huge way!! [https://github.com/kubernetes/website/issues/22515](https://github.com/kubernetes/website/issues/22515)
    *   **[annajung](https://app.slack.com/team/U8SLB1P2Q)**  Big shoutout to [@Celeste Horgan](https://kubernetes.slack.com/team/USF2T2W78) [@zacharysarah](https://kubernetes.slack.com/team/U5WQMKJEA) [@sftim](https://kubernetes.slack.com/team/UGBUYDQR2) [@karenb](https://kubernetes.slack.com/team/UCLQ9GKSP) for all their help on getting the docs all reviewed/merged for the 1.19 release! There were so many that needed a review right before the deadline, and they all went the extra mile to help out the docs release team! and shout out to docs lead [@sraghunathan](https://kubernetes.slack.com/team/UC8U2V3BM) for being an awesome lead through out this process! Thanks everyone!
    *   **[Celeste Horgan](https://app.slack.com/team/USF2T2W78)**  Shoutout to [@shuuji3](https://kubernetes.slack.com/team/UFKB9GXUZ) who took on a huge PR to clean up case study images in `kubernetes/website`! Clean all the things!!!
    *   **[onlydole](https://app.slack.com/team/U1DD4AZND)** I want to give a GIANT shoutout to [@Lauri Apple](https://kubernetes.slack.com/team/U011C07244F), [@tpepper](https://kubernetes.slack.com/team/U6UB5V4TX), [@justaugustus](https://kubernetes.slack.com/team/U0E0E78AK), [@jerickar](https://kubernetes.slack.com/team/U72ESU398), [@hasheddan](https://kubernetes.slack.com/team/ULLQEF30C), and [@mrbobbytables](https://kubernetes.slack.com/team/U511ZSKHD) for being such an incredible help today with knocking out 1.19 Release blockers! Teamwork makes the dream work
    *   **[paris](https://app.slack.com/team/U5SB22BBQ)** shoutout to [@dims](https://kubernetes.slack.com/team/U0Y7A2MME) for being an amazing mentor and advocate; dims is currently participating in several mentoring programs including acting as a group lead for 5 folks going from member to reviewer. doing the good work!!
    *   **[kaslin](https://app.slack.com/team/U5ENKU0AE)** shout out to [@rajula96reddy](https://kubernetes.slack.com/team/U7K9EK1HC) for giving an awesome demo of our new twitter repo automation in the contrib-ex marketing team meeting!
    *   **[thockin](https://app.slack.com/team/U0AH4GABW)** Shoutout to [@listx](https://kubernetes.slack.com/team/UFCU8S8P3), who has worked TIRELESSLY for several quarters to get the GCR vanity name flipped to the new community-owned infrastructure and promoter.  This is a big moment in k8s community history - one of the biggest single infrastructural pieces is now properly community owned.
    *   **[paris](https://app.slack.com/team/U5SB22BBQ)** shoutout to [@kaslin](https://kubernetes.slack.com/team/U5ENKU0AE) for building and growing the new @k8scontributors twitter account. kaslin has a lot more planned in the future so be sure to follow.
    *   **[liggitt](https://app.slack.com/team/U0BGPQ6DS)** shoutout to [@hasheddan](https://kubernetes.slack.com/team/ULLQEF30C) for doing the grunge work of helping make CI tests faster and more reliable ([https://github.com/kubernetes/kubernetes/pull/93448](https://github.com/kubernetes/kubernetes/pull/93448) takes a verify check from ~6 minutes to ~6 seconds) (edited)
    *   **[justaugustus](https://app.slack.com/team/U0E0E78AK)**  [https://twitter.com/stephenaugustus/status/1289370393526931456?s=19](https://twitter.com/stephenaugustus/status/1289370393526931456?s=19)
        *   Kudos to everyone diving in on #Kubernetes test-infra/CI Signal discussions/actions this week!
        *   [@BenTheElder](https://twitter.com/BenTheElder) [@spiffxp](https://twitter.com/spiffxp) [@Lauri_Apple](https://twitter.com/Lauri_Apple) [@hasheddan](https://twitter.com/hasheddan) [@RobKielty](https://twitter.com/RobKielty) [@alejandrox135@saschagrunert](https://twitter.com/alejandrox135) [@liggitt](https://twitter.com/liggitt) [@pythomit](https://twitter.com/pythomit)... just to name a few. :)
    *   **[dims](https://app.slack.com/team/U0Y7A2MME)** Big shoutout to [@briangrant](https://kubernetes.slack.com/team/U09R2JFE3) for all his work over the years and making room as emeritus from root OWNERS!
    *   **[inductor](https://app.slack.com/team/UD179SDGA)** Shoutout to [@shuuji3](https://kubernetes.slack.com/team/UFKB9GXUZ) for accelerating Japanese doc localization. His contributions are really quick and that has helped us a lot!
    *   **[Inductor](https://app.slack.com/team/UD179SDGA)** Shoutout to [@oke-py](https://kubernetes.slack.com/team/UEQ21UVMF) for managing the bunch of issues and the whole milestone in our document localization. This could not have done without him!
    *   **[dims](https://app.slack.com/team/U0Y7A2MME)** Another Big shoutout to [@brendanburns](https://kubernetes.slack.com/team/U0BC5M36Y) for his work on k8s. Thanks Brendan! [https://github.com/kubernetes/kubernetes/pull/93683](https://github.com/kubernetes/kubernetes/pull/93683)
    *   **[nikhita](https://app.slack.com/team/U2PQHGMLN)** Figured this was apt here. This is really awesome! Great work, [@michaelkolber](https://kubernetes.slack.com/team/U01531EUK1B)
        *   cc [@bentheelder](https://kubernetes.slack.com/team/U1P7T516X)
    *   **[listx](https://app.slack.com/team/UFCU8S8P3)** belated shoutout to [@justaugustus](https://kubernetes.slack.com/team/U0E0E78AK) for the work in very quickly adding support for building debian-base v1 (stretch) images [https://github.com/kubernetes/release/issues/1472](https://github.com/kubernetes/release/issues/1472)!
    *   **[Lauri Apple](https://app.slack.com/team/U011C07244F)** Shoutout to [@RobKielty](https://kubernetes.slack.com/team/UBEGAFN2V) for helping to push CI Policy Improvements work forward! You're digging in, framing questions and steps, and helping us to project-manage — all essential to our progress.
    *   **[RobKielty](https://app.slack.com/team/UBEGAFN2V)** Aww thanks [@nikhita](https://kubernetes.slack.com/team/U2PQHGMLN) thanks [@cpanato](https://kubernetes.slack.com/team/U8DFY4TTK) and thank you [@Lauri Apple](https://kubernetes.slack.com/team/U011C07244F)  y'all are the best!
    *   **[Celeste Horgan](https://app.slack.com/team/USF2T2W78)** Shoutout to [@jimangel](https://kubernetes.slack.com/team/U4HSVFA5U) for going deep down the commit history rabbit hole to 2015 to figure out whether we could remove a small section of code from the website. You rock!! [https://github.com/kubernetes/website/pull/22998#issuecomment-673124327](https://github.com/kubernetes/website/pull/22998#issuecomment-673124327)


### July 16, 2020 - [(recording)](https://www.youtube.com/watch?v=J3O8fXTm3HE&t=3106s)

*   **Moderators**:  Sushmita Amarnath [SIG Contribex/SIG Release]
*   **Note Taker**: Alison Dowdney [SIG Contribex]
    *   Subscribe to [this thread](https://discuss.kubernetes.io/t/kubernetes-weekly-community-meeting-notes/35) to get these notes in your inbox
*   [ 26:00 ]** Release Updates**
    *   Current Release Development Cycle  [Taylor Dolezal]
    *   Patch Release Updates [https://git.k8s.io/sig-release/releases/patch-releases.md](https://git.k8s.io/sig-release/releases/patch-releases.md)
        *   1.18.6, 1.17.9, 1.16.13 released July 15
        *   Next scheduled release August 12; will likely be last for 1.16 branch
*   [ 0:00 ] **SIG Updates**
    *   SIG Contributor Experience [Jorge Castro / Bob Killen]
        *   [Slides](https://docs.google.com/presentation/d/1AwqjZHmLmpZP5GcRGybV0HtTLB2pFCZsotLzcz69C1Y/edit#slide=id.g8c233a2db1_0_0)
    *   SIG Cluster Lifecycle [Fabrizio Pandini]
        *   [Slides](https://docs.google.com/presentation/d/1ZEDeF6lqxP-LmxCRa2EBmDS1sZFAv3RmrdQOUyd6IAc/edit?usp=sharing)
    *   [27:30] SIG Cloud Provider [Andrew Sy Kim]
        *   [Slides](https://docs.google.com/presentation/d/1wKOEuyFvdEDnrqaJP0ni0geB2tCAW92jIsHsB95zD4M/edit?usp=sharing)
    *   [37:40] SIG Network [Rob Scott]
        *   [Slides](https://docs.google.com/presentation/d/1k4uzqWCQgz8by3ZNUeXb1A5aeOpufOt4UTdGMLd6rjc/edit?usp=sharing)
*   [ 50:00 ] **Announcements **
    *   New Contributor Workshop Kickoff to be announced in mailing list today on k-dev - keep a lookout
*   **Shout Outs!**

    **[paris](https://app.slack.com/team/U5SB22BBQ)** shoutout to our amazing upstream mentor panel for julys edition of [#meet-our-contributors](https://kubernetes.slack.com/archives/C8WRR2BB9) [@jerickar](https://kubernetes.slack.com/team/U72ESU398) [@cecile](https://kubernetes.slack.com/team/U98JPHB2M) [@alejandrox1](https://kubernetes.slack.com/team/U6AS37R50). thank you for your time and contributions to the community! (all: recording will be live soon - check it out if you are a current or aspiring contributor)


    **[sraghunathan](https://app.slack.com/team/UC8U2V3BM)**  Shoutout to [@Divya](https://kubernetes.slack.com/team/UV4J7K97Z) [@annajung](https://kubernetes.slack.com/team/U8SLB1P2Q) (1.19 RT docs shadows) for stepping up without being asked and helping with balancing the load across the team :)  They are both proactive and encourage me to be better at what I do :) Thank you [@annajung](https://kubernetes.slack.com/team/U8SLB1P2Q) [@Divya](https://kubernetes.slack.com/team/UV4J7K97Z) for inspiring me to be a better lead! (edited)


    **[pohly](https://app.slack.com/team/U91901TMF)**  Shoutout for helping with review, debugging and getting "generic ephemeral volumes" and "CSI storage capacity tracking" ready for 1.19: [@jsafrane](https://kubernetes.slack.com/team/U0F49KUHE), [@thockin](https://kubernetes.slack.com/team/U0AH4GABW), [@claytonc](https://kubernetes.slack.com/team/U09SJ9AH0),[@msau42](https://kubernetes.slack.com/team/U3CB0SFMF)


    **[mrbobbytables](https://app.slack.com/team/U511ZSKHD)** shoutout to [@palnabarun](https://kubernetes.slack.com/team/UBH9NTMBM) [@oikiki](https://kubernetes.slack.com/team/U9HFFRFT2) [@johnbelamaric](https://kubernetes.slack.com/team/U246A1A0N) [@Harsha Narayana](https://kubernetes.slack.com/team/UH6AY4WDC) and [@Miroslaw Sedzinski](https://kubernetes.slack.com/team/U011WNA7VSB) for wrangling all the Enhancements around code freeze \o/


    **[Joel Barker](https://app.slack.com/team/U014AHSTBPA)**  Shoutout to [@Celeste Horgan](https://kubernetes.slack.com/team/USF2T2W78) for taking on Hugo-in-container.



### June 18, 2020 - ([recording](https://youtu.be/ObqQxRRl9RQ))

*   **Moderators**:  Lauri Apple [SIG Contribex/SIG Release]
*   **Note Taker**: Lachlan Evanson [Microsoft/SIG Breadsticks]
    *   Subscribe to [this thread](https://discuss.kubernetes.io/t/kubernetes-weekly-community-meeting-notes/35) to get these notes in your inbox
*   [ 0:01 ]** Release Updates**
    *   Current Release Development Cycle  [Taylor Dolezal - Release Lead]
    *   Patch Release Updates [https://git.k8s.io/sig-release/releases/patch-releases.md](https://git.k8s.io/sig-release/releases/patch-releases.md)
        *   June 17th, 2020 patch releases are as follows:
            *   1.18.4: known issue under triage regarding kubernetes-cni conflicts/obsoletes in yum repository
            *   1.17.7
            *   1.16.11
        *   Next scheduled is July 2020:
            *   Cherry pick deadline: 2020-07-10
            *   Release target: 2020-07-15
        *   August through December 2020 monthly patch release target dates are listed in [https://git.k8s.io/sig-release/releases/patch-releases.md](https://git.k8s.io/sig-release/releases/patch-releases.md)
        *   1.19 Release schedule shifts
            *   [https://groups.google.com/forum/#!topic/kubernetes-dev/TVXhcNO3SPU](https://groups.google.com/forum/#!topic/kubernetes-dev/TVXhcNO3SPU)
*   [ 0:03 ]** Contributor Tip of the Month: **Triage-party update/demo [Thomas Stromberg]
    *   Tool for tracking issue triage and collates responses in a logical, actionable layout in a WebUI
    *   Ability to split the workload amongst all participants
    *   Kanban and milestone views available
    *   Can be run on Kubernetes or locally
    *   [https://github.com/google/triage-party](https://github.com/google/triage-party)
    *   [tinyurl.com/minikubeparty](http://tinyurl.com/minikubeparty) &lt;- Minikube instance of triage-party
    *   Q: Are you aware of usage?
    *   A: Not sure but have done some GH searches and there is some usage however it’s only been open sourced about 2 months
    *   Q: How long does it take to stand up?
    *   A: 15m to get started using example config for Kubernetes. Storage options for github cache can be a little more complicated.
    *   Q: Are there video resources?
    *   A: No yet. Just written documentation
    *   Q: What has been the effect on morale and contributor relationships on projects using triage-party?
    *   A: We do a better job keeping an open line of communication with our users. Also helpful for prioritization and having people from the community stepping up to help with triage.
    *   **Next steps:** Leader boards for positive behaviors
    *   If you have any questions, reach out to tstromberg@ on Kubernetes Slack
*   [ 0:19 ] **SIG Updates**
    *   SIG Leads, check out this set of [recommended topics to cover](https://github.com/kubernetes/community/blob/master/events/community-meeting.md#sig-updates) during your update. [Slide template if you need it](https://docs.google.com/presentation/d/1-nTvKCiqu9UvFYUeM6p6RIqHS5-H-u3_x-V4xj_eIWo/edit#slide=id.g401c104a3c_0_0), please also check the [SIG Update Schedule](https://docs.google.com/spreadsheets/d/1adztrJ05mQ_cjatYSnvyiy85KjuI6-GuXsRsP-T2R3k)!
    *   SIG Release [Stephen Augustus]
        *   [Slides](https://docs.google.com/presentation/d/1VuS06RJWccA8ceSlu5a3hIIsVb2mge6Zz5YDpxHiBFY/edit#slide=id.g401c104a3c_0_0)
    *   SIG Apps  [Ken Owens]
        *   [Slides](https://docs.google.com/presentation/d/18UcJQs3ThW6Vdgl_mdc1984uU16GwuIhuD0Pujl42xU/edit?usp=sharing)
    *   Security Response Committee [Tim Allclair]
        *   [Slides](https://docs.google.com/presentation/d/1TJQevF8wLRsjVRNQuUiRgwI5EO2mdt03RfifYwDoGrc/preview)
    *   SIG Architecture [John Belamaric]
        *   [Slides](https://docs.google.com/presentation/d/1NytMrpVYKzFo7rLcEEHnFl8zOx05fnjs3xBSZXVE0nI/edit#slide=id.g401c104a3c_0_0)
*   [ 0:58 ] **Announcements **


### June 18, 2020 - ([recording](https://youtu.be/ObqQxRRl9RQ))



*   **Moderators**:  Lauri Apple [SIG Contribex/SIG Release]
*   **Note Taker**: Lachlan Evanson [Microsoft/SIG Breadsticks]
    *   Subscribe to [this thread](https://discuss.kubernetes.io/t/kubernetes-weekly-community-meeting-notes/35) to get these notes in your inbox
*   [ 0:01 ]** Release Updates**
    *   Current Release Development Cycle  [Taylor Dolezal - Release Lead]
    *   Patch Release Updates [https://git.k8s.io/sig-release/releases/patch-releases.md](https://git.k8s.io/sig-release/releases/patch-releases.md)
        *   June 17th, 2020 patch releases are as follows:
            *   1.18.4: known issue under triage regarding kubernetes-cni the conflicts/obsoletes in yum repository
            *   1.17.7
            *   1.16.11
        *   Next scheduled is July 2020:
            *   Cherry pick deadline: 2020-07-10
            *   Release target: 2020-07-15
        *   August through December 2020 monthly patch release target dates are listed in [https://git.k8s.io/sig-release/releases/patch-releases.md](https://git.k8s.io/sig-release/releases/patch-releases.md)
        *   1.19 Release schedule shifts
            *   [https://groups.google.com/forum/#!topic/kubernetes-dev/TVXhcNO3SPU](https://groups.google.com/forum/#!topic/kubernetes-dev/TVXhcNO3SPU)
*   [ 0:03 ]** Contributor Tip of the Month: **Triage-party update/demo [Thomas Stromberg]
    *   Tool for tracking issue triage and collates responses in a logical, actionable layout in a WebUI
    *   Ability to split the workload amongst all participants
    *   Kanban and milestone views available
    *   Can be run on Kubernetes or locally
    *   [https://github.com/google/triage-party](https://github.com/google/triage-party)
    *   [tinyurl.com/minikubeparty](http://tinyurl.com/minikubeparty) &lt;- Minikube instance of triage-party
    *   Q: Are you aware of usage?
    *   A: Not sure but have done some GH searches and there is some usage however it’s only been open sourced about 2 months
    *   Q: How long does it take to stand up?
    *   A: 15m to get started using example config for Kubernetes. Storage options for github cache can be a little more complicated.
    *   Q: Are there video resources?
    *   A: No yet. Just written documentation
    *   Q: What has been the effect on morale and contributor relationships on projects using triage-party?
    *   A: We do a better job keeping an open line of communication with our users. Also helpful for prioritization and having people from the community stepping up to help with triage.
    *   **Next steps:** Leader boards for positive behaviors
    *   If you have any questions, reach out to tstromberg@ on Kubernetes Slack
*   [ 0:19 ] **SIG Updates**
    *   SIG Leads, check out this set of [recommended topics to cover](https://github.com/kubernetes/community/blob/master/events/community-meeting.md#sig-updates) during your update. [Slide template if you need it](https://docs.google.com/presentation/d/1-nTvKCiqu9UvFYUeM6p6RIqHS5-H-u3_x-V4xj_eIWo/edit#slide=id.g401c104a3c_0_0), please also check the [SIG Update Schedule](https://docs.google.com/spreadsheets/d/1adztrJ05mQ_cjatYSnvyiy85KjuI6-GuXsRsP-T2R3k)!
    *   SIG Release [Stephen Augustus]
        *   [Slides](https://docs.google.com/presentation/d/1VuS06RJWccA8ceSlu5a3hIIsVb2mge6Zz5YDpxHiBFY/edit#slide=id.g401c104a3c_0_0)
    *   SIG Apps  [Ken Owens]
        *   [Slides](https://docs.google.com/presentation/d/18UcJQs3ThW6Vdgl_mdc1984uU16GwuIhuD0Pujl42xU/edit?usp=sharing)
    *   Security Response Committee [Tim Allclair]
        *   [Slides](https://docs.google.com/presentation/d/1TJQevF8wLRsjVRNQuUiRgwI5EO2mdt03RfifYwDoGrc/preview)
    *   SIG Architecture [John Belamaric]
        *   [Slides](https://docs.google.com/presentation/d/1NytMrpVYKzFo7rLcEEHnFl8zOx05fnjs3xBSZXVE0nI/edit#slide=id.g401c104a3c_0_0)
*   [ 0:58 ] **Announcements **
        *   New twitter - follow us: @k8scontributors. An official announcement will be on kubernetes-dev@ mailing list today.
        *   Next month’s host will be Sushmita Amarnath! We’re always looking for new contributors to host this meeting, ping us in #sig-contribex if you’re interested.
    *   ** Shoutouts this month (Check in #shoutouts on slack) **

**[Paris](https://app.slack.com/team/U5SB22BBQ)**: Shoutouts to our brand new upstream marketing team as part of [#sig-contribex](https://kubernetes.slack.com/archives/C1TU9EB9S) working hard to make sure contributors are in the know. You can reach them if you need something amplified through a github issue template on Kubernetes/community or @contributor-comms tag in the #Sig-contribex channel. Thank you to our volunteers for building this!

**[Zacharysarah](https://app.slack.com/team/U5WQMKJEA)**: Shoutout to everyone who contributes to the [kubectl cheat sheet](https://slack-redir.net/link?url=https%3A%2F%2Fkubernetes.io%2Fdocs%2Freference%2Fkubectl%2Fcheatsheet%2F). It's consistently the website's most-visited page, and it helps people a lot. [https://github.com/kubernetes/website/issues/21458](https://slack-redir.net/link?url=https%3A%2F%2Fgithub.com%2Fkubernetes%2Fwebsite%2Fissues%2F21458)

**[Lauri Apple](https://app.slack.com/team/U011C07244F)**: Want to call out [@Divya](https://kubernetes.slack.com/team/UV4J7K97Z) here for her extreme resilience—despite facing electricity issues around Cyclone Nisgara she still showed up for a release team prioritisation session yesterday and added so much value to the team's efforts to simplify release cycle processes. I shared this with my day-job colleagues, they were super-impressed.

**[Bentheelder](https://app.slack.com/team/U1P7T516X)**: shoutout to everyone involved in the [kubernetes.io](https://slack-redir.net/link?url=http%3A%2F%2Fkubernetes.io) announcement

**[sftim](https://app.slack.com/team/UGBUYDQR2)**: Important work, carefully done: [https://github.com/kubernetes/website/pull/21359](https://slack-redir.net/link?url=https%3A%2F%2Fgithub.com%2Fkubernetes%2Fwebsite%2Fpull%2F21359)

I want to call out [@karenb](https://kubernetes.slack.com/team/UCLQ9GKSP) dealing with the aftermath of a tricky parser upgrade for the website, and helping enable the transition to a new Docsy theme.

**[Neolit123](https://app.slack.com/team/U83J4CS3S)**: shout out to [@nicksantos](https://kubernetes.slack.com/team/UBJETRVFY) from Tilt for writing a proposal for standardizing local cluster registries via SIG Cluster Lifecycle: [https://github.com/kubernetes/enhancements/pull/1757](https://slack-redir.net/link?url=https%3A%2F%2Fgithub.com%2Fkubernetes%2Fenhancements%2Fpull%2F1757). Nick coordinated with multiple parties, planned the specification and responded to reviewer comments.

**[Jdetiber](https://app.slack.com/team/U0UV07D8T)**: Huge shout out to [@Katie Gamanji](https://kubernetes.slack.com/team/UC6PGGXS4) and [@naadir](https://kubernetes.slack.com/team/U6RDFQAF5) for delivering an amazing Cluster API Webinar

**[Paris](https://app.slack.com/team/U5SB22BBQ)**: shoutout to [@nikhita](https://kubernetes.slack.com/team/U2PQHGMLN) for being a dependable, amazing, empathetic contributor :kubernetes-heart-eyes: thank you for all that you do.

**[Neolit123](https://app.slack.com/team/U83J4CS3S)**: shout out to [@Abhishek Tamrakar](https://kubernetes.slack.com/team/UMQGX19JS) for enduring the review for a much required documentation page for manual CA rotation:



*   [https://github.com/kubernetes/website/pull/19351](https://slack-redir.net/link?url=https%3A%2F%2Fgithub.com%2Fkubernetes%2Fwebsite%2Fpull%2F19351)
*   [https://github.com/kubernetes/website/pull/21651](https://slack-redir.net/link?url=https%3A%2F%2Fgithub.com%2Fkubernetes%2Fwebsite%2Fpull%2F21651)

**[Justaugustus](https://app.slack.com/team/U0E0E78AK)**: Thanks again to [@Lauri Apple](https://kubernetes.slack.com/team/U011C07244F) for her tireless optimization efforts in SIG Release (and across the project, in general). Seeing a lot of stuff start to shake loose, that we haven't able to prioritize, and I'm really appreciative!


### May 21, 2020 - ([recording](https://youtu.be/ZyUQiN3S6TE))

*   **Moderators**:  Marko Mudrinić [Loodse / SIG-Release]
*   **Note Taker**: First Last [Company/SIG]
    *   Subscribe to [this thread](https://discuss.kubernetes.io/t/kubernetes-weekly-community-meeting-notes/35) to get these notes in your inbox
*   [ 0:00 ]** Release Updates**
    *   Current Release Development Cycle  [Taylor Dolezal - Release Lead]
        *   1.19.0-beta.0 went out on Tues, May 19th 2020
        *   Enhancements FREEZE as of EOD Tues, May 19th 2020
        *   Please add items to the 1.19 retro as you think of them: [https://bit.ly/k8s119-retro](https://bit.ly/k8s119-retro)
    *   Patch Release Updates [https://git.k8s.io/sig-release/releases/patch-releases.md](https://git.k8s.io/sig-release/releases/patch-releases.md)
        *    Patch releases on all branches (1.18, 1.17, 1.16) yesterday
        *   Next patch releases likely mid-June
*   [ 0:00 ] **SIG Updates**
    *   SIG Leads, check out this set of [recommended topics to cover](https://github.com/kubernetes/community/blob/master/events/community-meeting.md#sig-updates) during your update. [Slide template if you need it](https://docs.google.com/presentation/d/1-nTvKCiqu9UvFYUeM6p6RIqHS5-H-u3_x-V4xj_eIWo/edit#slide=id.g401c104a3c_0_0), please also check the [SIG Update Schedule](https://docs.google.com/spreadsheets/d/1adztrJ05mQ_cjatYSnvyiy85KjuI6-GuXsRsP-T2R3k)!
    *   SIG Testing [Aaron Crickenberger, @spiffxp]
        *   [Slides](https://docs.google.com/presentation/d/1H-MLhKJJVsQG2eDCEv48M_WAzMc66dKaYMgfOSGQRJM/edit#slide=id.g401c104a3c_0_0)
    *   SIG UI [Jeffrey Sica]
        *   [Slides](https://docs.google.com/presentation/d/1W4NioOkAF2VFiu-5t80p2vlu3_OznpugiyiViFuitaM/edit#slide=id.g338ac0a8b6_0_27)
    *   SIG API Machinery  [Federico Bongiovanni]
        *   [Slides](https://docs.google.com/presentation/d/1UWRaMVtTD3yVhJ3MGBpt7LRIaRHTaQZoGlDT7Bl7jLE/edit#slide=id.g401c104a3c_0_0)
    *   SIG Usability [Vallery Lancey]
        *   [Slides](https://docs.google.com/presentation/d/18ISAYsExgoxk8ER47rrcv-89ZJBCZeUEdXLjT8q7HL8/)
*   [ 0:00 ] **Announcements **
    *   Reminder: Kubecon Virtual Contributor Summit: Canceled [Jeffrey Sica]
        *   K-Dev email: [https://groups.google.com/forum/#!topic/kubernetes-dev/jnPv42i2ACw](https://groups.google.com/forum/#!topic/kubernetes-dev/jnPv42i2ACw)
    *   Next month’s host will be Lauri Apple! We’re always looking for new contributors to host this meeting, ping us in #sig-contribex if you’re interested.
    *   ** **Shoutouts this month (Check in #shoutouts on slack) ** **
        *   Help wanted section can go here, for example: “SIG Foo is looking for shadows for 1.18 cycle”, etc.
        *   Someone out there making a difference? Give them a shoutout in #shoutouts so the community can celebrate their actions:
*   [Tpepper](https://app.slack.com/team/U6UB5V4TX): Shoutout to [@spiffxp](https://kubernetes.slack.com/team/U09R2FL93) for some test heroics in [https://kubernetes.slack.com/archives/C0BP8PW9G/p1587053606339800](https://kubernetes.slack.com/archives/C0BP8PW9G/p1587053606339800) to fix up a swatch of failing SIG Node test signal on kubelet
*   [mbbroberg](https://app.slack.com/team/U18JTHMDY)  Huge thanks to [@sftim](https://kubernetes.slack.com/team/UGBUYDQR2) for his attention to detail on posts (he notes them as nits, which is kind. I call them gifts <3 )
*   [bartsmykla](https://app.slack.com/team/U9VDVCXGU):  Huge shoutout to [@ameukam](https://kubernetes.slack.com/team/U68KPQ448) for managing to move with [@nikhita](https://kubernetes.slack.com/team/U2PQHGMLN) publishing-bot to our new infrastructure, and not lose motivation, even after a lot of comments and changes ([the process took more than 4 months](https://github.com/kubernetes/k8s.io/pull/520))
*   [Mrbobbytables](https://app.slack.com/team/U511ZSKHD): shoutout to [@zacharysarah](https://kubernetes.slack.com/team/U5WQMKJEA) for helping me chase down a hugo templating problem with the contributor site <3 <3
*   [nikhita](https://app.slack.com/team/U2PQHGMLN)  - Big shoutout to [@markyjackson](https://kubernetes.slack.com/team/U19TKJ64E)! He's been consistently helping new folks get started in contribex, answering questions in [#sig-contribex](https://kubernetes.slack.com/archives/C1TU9EB9S), responding and triaging to incoming PRs and issues in k/community (our response times have reduced so much), and leading our mentoring subproject meetings! This doesn't even begin to scratch the surface with how much work he's been doing tbh. Thanks so much, Marky! :kubernetes-heart-eyes:
*   [bartsmykla](https://app.slack.com/team/U9VDVCXGU) - Shoutout to [@Eric Lemieux](https://kubernetes.slack.com/team/UJ8NKGEP6) for merging his first contribution to k/k8s.io! Great job, and thank you for your work! (ref. [#786](https://github.com/kubernetes/k8s.io/pull/786))
*   [mbbroberg](https://app.slack.com/team/U18JTHMDY) - I finally got my head wrapped around Gubernator today. Serious thanks to the team behind it and [@mrbobbytables](https://kubernetes.slack.com/team/U511ZSKHD) / [@justaugustus](https://kubernetes.slack.com/team/U0E0E78AK) for recommending it! [https://gubernator.k8s.io/](https://gubernator.k8s.io/)
*   [Fabrizio.pandini](https://app.slack.com/team/U3LBAU3LN) - Shoutout to [@fale](https://kubernetes.slack.com/team/U7YKA40JF) for bringing the Italian localisation of the web-site to life! Amazing Job!
*   [markyjackson](https://app.slack.com/team/U19TKJ64E)  - Shoutout to [@cpanato](https://kubernetes.slack.com/team/U8DFY4TTK) and [@veronica](https://kubernetes.slack.com/team/U7NNE57PU) for mentoring and calming effects while I cut the 1.19.0-alpha.3 release for the 1st time
*   [Cpanato](https://app.slack.com/team/U8DFY4TTK) - Huge shoutout to [@xmudrii](https://kubernetes.slack.com/team/U4Q2TNGVD) for setup all we need to start testing with Digital ocean. You bet!!
*   [markyjackson](https://app.slack.com/team/U19TKJ64E)  - I would like to give [@Pierre Humberdroz](https://kubernetes.slack.com/team/UFCT44U86) and [@carlisia](https://kubernetes.slack.com/team/UBDSF40G2) shouts outs for helping me fix a bug in slack infra that was discovered this weekend. Really appreciate your guidance
*   [markyjackson](https://app.slack.com/team/U19TKJ64E)  - I would also like to give a shoutout to [@mrbobbytables](https://kubernetes.slack.com/team/U511ZSKHD) for also helping me this  weekend. Super appreciative of you
*   Hasheddan - Hey folks! If you are interacting with any of [@mhb](https://kubernetes.slack.com/team/U11H6PJUB) [@vpickard](https://kubernetes.slack.com/team/UFGJMKHK8) [@Ed](https://kubernetes.slack.com/team/U8VJC20R3) [@spiffxp](https://kubernetes.slack.com/team/U09R2FL93) this week please tell them thank you for the great work they did today to fix node-kubelet-master, which is now green after they identified a particularly hairy issue. Special shout out to [@vpickard](https://kubernetes.slack.com/team/UFGJMKHK8) who has been leading a general effort to get more insight into the sig-node tests! Thank y’all! :raised-hands x3: [https://testgrid.k8s.io/sig-release-master-blocking#node-kubelet-master](https://testgrid.k8s.io/sig-release-master-blocking#node-kubelet-master)


### April 16, 2020 - ([recording](https://www.youtube.com/watch?v=Y3z2grPHRh4))


*   **Moderators**:  Taylor Dolezal [SIG Release, SIG Docs]
*   **Note Taker**: Robert Kielty
    *   Subscribe to [this thread](https://discuss.kubernetes.io/t/kubernetes-weekly-community-meeting-notes/35) to get these notes in your inbox
*   [ 0:00 ]** Release Updates**
    *   When reviewing updates from SIGs from this meeting refer to the 1.18 release notes to  **[https://kubernetes.io/docs/setup/release/notes/](https://kubernetes.io/docs/setup/release/notes/)**
    *   Current Release Development Cycle  [Taylor Dolezal - 1.19 Release Team Lead]
    *   1.19 is upon us! The shadow selection process is underway and should be finished by early next week. 1.19 will be a release focused on process improvement around the release itself and reducing overall work in progress with what’s currently going on in the world.
    *   Patch Release Updates [https://git.k8s.io/sig-release/releases/patch-releases.md](https://git.k8s.io/sig-release/releases/patch-releases.md)
        *   1.18.2, 1.17.5, 1.16.9, 1.15.12  - Cherry Pick deadline was 04/13/20 - Release **target** today: 04/16/20
        *   SIG Node issue under discussion at [https://kubernetes.slack.com/archives/C0BP8PW9G/p1587053606339800](https://kubernetes.slack.com/archives/C0BP8PW9G/p1587053606339800) may block today’s patch releases
*   [ 0:00 ] **SIG Updates**
    *   SIG Docs [Zach Corleissen]
        *   Slides - [https://docs.google.com/presentation/d/12WnYz8SbjWRZbK4k2qlc1Ab7Z-f2F7kk-Tb2eRTfTy8/edit#slide=id.g401c104a3c_0_0](https://docs.google.com/presentation/d/12WnYz8SbjWRZbK4k2qlc1Ab7Z-f2F7kk-Tb2eRTfTy8/edit#slide=id.g401c104a3c_0_0)
            *   Hugo Docsy Theme
                *   [https://themes.gohugo.io/docsy/](https://themes.gohugo.io/docsy/)
                *   [https://github.com/kubernetes/website/issues/20344](https://github.com/kubernetes/website/issues/20344)
            *   Submitting a docs PR? Be sure to Squash Your Commits, find out more here  via Jim, credited to Celeste, [https://github.com/kubernetes/community/blob/master/contributors/guide/github-workflow.md](https://github.com/kubernetes/community/blob/master/contributors/guide/github-workflow.md)
            *   Please familiaize yourself with the content guide for docs [https://kubernetes.io/docs/contribute/style/content-guide/#what-s-allowed](https://kubernetes.io/docs/contribute/style/content-guide/#what-s-allowed)
    *   SIG Apps - Skipping
    *   SIG CLI [Maciej Szulik and Sean Sullivan]
        *   Slides - [https://docs.google.com/presentation/d/1Y8SHFz6yyYS6rvRCgYUrSgF-moPk_YB6A-7Ykw5eWnU/edit#slide=id.g401c104a3c_0_0](https://docs.google.com/presentation/d/1Y8SHFz6yyYS6rvRCgYUrSgF-moPk_YB6A-7Ykw5eWnU/edit#slide=id.g401c104a3c_0_0)
            *   Kubectl is being migrated out of the core k8s repo via the staging repo (the staging folder being a holding area for parts of the core k8s repo that have a separate new repo as an ultimate final home)
            *   There are good first commit opportunities see slides above
            *   Look at [kustomise](https://github.com/kubernetes-sigs/kustomize) and kubectl plugin management via [krew](https://github.com/kubernetes-sigs/krew)  
    *   SIG Node  [Derek Wayne Carr]
        *   Slides -
        *   [https://docs.google.com/presentation/d/1XQReOXaZP1KOfU3oyMOPvHokSLZSpeIfZneObkSn8aI/edit#slide=id.g401c104a3c_0_](https://docs.google.com/presentation/d/1XQReOXaZP1KOfU3oyMOPvHokSLZSpeIfZneObkSn8aI/edit#slide=id.g401c104a3c_0_0)
            *   [Kubernetes Topology Manager Moves to Beta - Align Up!](https://kubernetes.io/blog/2020/04/01/kubernetes-1-18-feature-topoloy-manager-beta/)
            *   TaintBasedEvictions ([#87487](https://github.com/kubernetes/kubernetes/pull/87487), [@skilxn-go](https://github.com/skilxn-go)) [SIG API Machinery, Apps, Node, Scheduling and Testing]
*   [ 0:00 ] **Announcements **
    *   Launching CommunityBridge Mentorships Q2 2020:** **[https://www.cncf.io/blog/2020/04/16/launching-communitybridge-mentorships-q2-2020/](https://www.cncf.io/blog/2020/04/16/launching-communitybridge-mentorships-q2-2020/)
    *   Next month, [Marko Mudrinić](https://twitter.com/xmudrii) will be hosting
        *   Software Developer at Loodse, project maintainer of Kubicorn, and many more!
    *   ** **Shoutouts this month (Check in #shoutouts on slack) ** **
        *   **Taylor Dolezal:** I would like to give a very huge shoutout to @mrbobbytables!!!!! He has been beyond attentive to all things Blog, Docs, Release, and SO MUCH MORE! I know that without his help, this week would’ve gone by a lot more slowly, and would’ve been a lot more difficult. You are the best, Bob!


### March 19, 2020 - ([recording](https://youtu.be/04FiOd-bG8c))



*   **Moderators**:  Matt Broberg [Red Hat /SIG Contribex]
*   **Note Taker**: Tim Pepper [ VMware / SIG Release]
    *   Subscribe to [this thread](https://discuss.kubernetes.io/t/kubernetes-weekly-community-meeting-notes/35) to get these notes in your inbox
*   [ 12:10 (time in fluid, amirite)]** Release Updates**
    *   Current Release Development Cycle  [Jorge Alarcon - Release Manager]
        *   This is the last week of the 1.18 release cycle, release target is Tuesday next week March 24
        *   **Cherry Pick Deadline** (EOD PST) Thu March 19
        *   **v1.18.0 released Tue March 24**
        *   Now that we are past code-freeze, any PR targeting 1.18 must be [cherry-picked](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-release/cherry-picks.md) from the master branch to the release-1.18 branch.  Please use the cherry pick script described in the link.
        *   1.18 retrospective will be scheduled for April 2, with possible follow on discussion April 6 during the SIG Release biweekly meeting
    *   [Patch Release Updates and Schedule https://git.k8s.io/sig-release/releases/patch-releases.md](https://git.k8s.io/sig-release/releases/patch-releases.md)
        *   March 12, 2020 1.17.4, 1.16.8, and 1.15.11 were released.
        *   Next releases are targeted for April 16, with a cherry pick deadline of April 13.
        *   Next month’s 1.15.12 release is likely to be the final patch release for the release-1.15 branch.
*   [ 12:03 ] **SIG Updates**
    *   SIG Leads, check out this set of [recommended topics to cover](https://github.com/kubernetes/community/blob/master/events/community-meeting.md#sig-updates) during your update. [Slide template if you need it](https://docs.google.com/presentation/d/1-nTvKCiqu9UvFYUeM6p6RIqHS5-H-u3_x-V4xj_eIWo/edit#slide=id.g401c104a3c_0_0), please also check the [SIG Update Schedule](https://docs.google.com/spreadsheets/d/1adztrJ05mQ_cjatYSnvyiy85KjuI6-GuXsRsP-T2R3k)!
    *   Steering committee [Nikhita Raghunath]
        *   [Slides](https://docs.google.com/presentation/d/14sfm543msjQOUTsRrC8pbhHxvd0BTunGEkeE4FAiydk/edit?usp=sharing)
        *   Community Health Check initiative: [https://github.com/kubernetes/steering/issues/153](https://github.com/kubernetes/steering/issues/153)
        *   Kubernetes org repos cleanup: [https://github.com/kubernetes/steering/issues/136](https://github.com/kubernetes/steering/issues/136)
        *   More consistent “help wanted” policy: [https://github.com/kubernetes/steering/issues/141](https://github.com/kubernetes/steering/issues/141)
        *   Contributor travel funding assistance work underway for whenever KubeCon EU 2020 eventually happens: [https://github.com/kubernetes/steering/issues/109](https://github.com/kubernetes/steering/issues/109)
    *   SIG Instrumentation [Elana Hashman]
        *   [Slides](https://docs.google.com/presentation/d/1tWWWsnZZPcoMAYj60jL2541cNA1GyLCR4i-EQzaI9y8/edit#)
        *   New SIG leads, tech leads, and emeritus leads listed in slides and official [SIG List](https://github.com/kubernetes/steering/issues/109)
        *   Metrics stability framework graduating in 1.18 from alpha to beta (/metrics/resource endpoint)
        *   [Structured logging KEP](https://github.com/kubernetes/enhancements/pull/1367) moved to “implementable” with alpha target in 1.19, new contributors wanted (contact @serathius)
        *   [Tracing KEP merged,](https://github.com/kubernetes/enhancements/pull/650) SIG API Machinery doing provisional work
        *   Subprojects:
            *   Kube-state-metrics: 1.9.x released and in bugfix only mode, new contributors wanted (contact @lilic)
            *   Metrics-server: docs and installation improvements, new contributors wanted
    *   SIG Service Catalog [Jonathan Berkhahn]
        *   &lt; no slides >
        *   SIG handles an open service broker implementation
        *   Things quiet lately, mostly maintenance mode, seeking new contributors
        *   CRD based catalog released
        *   If you’re interested in helping with this SIG exploring collaboration with others, and you’re a cloud service user, join in
    *   SIG Storage  [Saad Ali]
        *   [Slides](https://docs.google.com/presentation/d/1riA61cmvdPZiZcY3cwKK0drRZUHS7nAdmNBazDIeYyY/edit?usp=sharing)
        *   New chairs, tech leads
        *   1.17 highlights:
            *   CSI Topology
            *   Volume Snapshot moves to Beta
            *   CSI Migration in Beta
        *   1.18 highlights:
            *   Raw block volumes exposed as /dev node instead of mounted filesystem
            *   Volume cloning moves to GA duplicating PVC’s, if underlying CSI implementation supports it
            *   CSIDriver Kubernetes API Object moves to GA simplifying CSI driver discovery
            *   Windows CSI support introduced in Alpha
            *   Recursive volume ownership OnRootMismatch option introduced in Alpha to speed volume mount time when ownerships are changed at mount time
    *   Code of conduct committee [Tasha Drew]
        *   &lt; Missed meeting, reschedule >
*   [ 0:00 ] **Announcements **
    *   KubeCon update
        *   _Last note from March 4_
        *   [KubeCon + CloudNativeCon EU has been delayed](https://events.linuxfoundation.org/kubecon-cloudnativecon-europe/attend/novel-coronavirus-update/) until July/August
        *   There will be an in-person Contributor Summit there
        *   More virtual meetings are TBD
        *   Thank yous to all our ecosystem tech events organizers who are dealing with very complicated logistics right now
    *   Next month [Taylor Dolezal](https://twitter.com/onlydole) host
        *   SRE at Disney, book fan, (and the 1.19 release lead!)
    *   ** **Shoutouts this month (Check in #shoutouts on slack) ** **
        *   **Paris Pittman**: MEGA shout outs to everyone planning an upcoming event, most notably for our crowd: the #contributor-summit and #kubecon teams. An immeasurable amount of work is going on behind the scenes to figure things out. Spread the word to be patient and kind.
        *   **Jeremy Rickard**: Shout out to all the enhancements 1.18 shadows for being super attentive and on top of things and helping me out this release! @oikiki @Heba @palnabarun @johnbelamaric
        *   **Bartsmykla**: Shout out to @listx for all of his work at container image promoter which takes us much closer to moving infrastructure to community owned. Thank you @listx for all of your hard work on that!
        *   **Antonio Ojea**: Shout out to @bentheelder and @aramase for keeping a quality high bar for the project, and as an example of community work, detecting an issue in one CI job in twitter and fixing it the day after
        *   **Jason DeTiberus**: Huge shoutout to @jayunit100 for getting the forward port and parallelization of the e2e tests for cluster-api-provider-aws across the finish line, especially for fixing all of the bugs I left him with my initial parallelization PoC . Also huge thanks to @T V KUTUMBA RAO SIDRALA @Bhargav Madduru for their work building out the test suite and their initial forward porting efforts.  Additional props to @naadir for his assistance with the forward porting efforts, reviews, and testing.
        *   **Vishakha Nihore**: Huge shoutout to @paris @mrbobbytables & @markyjackson for the constant support and help during my Outreachy Internship. Considering it was my first time working with one of the largest community, I had not imagined it would be so much awesome. I couldn't have imagined working with Kubernetes can be this much fun TBH. All the lgtm  labels I asked from @markyjackson and all the questions (most of them quite simple and absurd) I asked from @mrbobbytables and the constant help I have got from @paris is just priceless.
        *   **Marko Mudrinić**: HUGE shoutout to @jdetiber for helping me understand what is new in the cluster-api v1alpha2 and answering every single question I had! He helped us a lot to understand how things should be done and how to bootstrap the project!


### February 20, 2020 - recording



*   **Moderators**:  Vamshi Samudrala [American Airlines/SIG Contribex]
*   **Note Takers**: Laura Santamaria and Jorge Castro [SIG Contribex]
    *   Subscribe to [this thread](https://discuss.kubernetes.io/t/kubernetes-weekly-community-meeting-notes/35) to get these notes in your inbox
*   [ 0:00 ] Release Updates
    *   Current Release Cycle  [Jorge Alarcon - Release Team Lead]
    *   We are at week 7 (out of 12)!
    *   Currently we are tracking 50 enhancements
        *   Alpha: 18
        *   Beta: 16
        *   Stable: 16
    *   v1.18.0-alpha.5 was released this week.
    *   Release branch 1.18 was created (1.14 CI jobs were deleted)
    *   Code freeze scheduled for Thursday, March 05, 2020 [sig-release/releases/release-1.18](https://github.com/kubernetes/sig-release/tree/master/releases/release-1.18)
    *   Monday, March 16: Week 11 - Docs must be completed and reviewed
    *   Tuesday, March 24: Week 12 - Kubernetes v1.18.0 released
    *   Patch Release Updates [https://git.k8s.io/sig-release/releases/patch-releases.md](https://git.k8s.io/sig-release/releases/patch-releases.md)
        *   1.15.10, 1.16.7, and 1.17.3 were released on 2020-02-11
*   [ 0:00 ] **SIG Updates**
    *   SIG Leads, check out this set of [recommended topics to cover](https://github.com/kubernetes/community/blob/master/events/community-meeting.md#sig-updates) during your update. [Slide template if you need it](https://docs.google.com/presentation/d/1-nTvKCiqu9UvFYUeM6p6RIqHS5-H-u3_x-V4xj_eIWo/edit#slide=id.g401c104a3c_0_0), please also check the [SIG Update Schedule](https://docs.google.com/spreadsheets/d/1adztrJ05mQ_cjatYSnvyiy85KjuI6-GuXsRsP-T2R3k)!
    *   SIG Windows[Michael Michael / Patrick Lang / Deep Debroy] Confirmed
        *   Big investments coming that needed to be stabilized
        *   Windows identity:
            *   Graduating active directory and kube managed service accounts from beta to stable
            *   runAsUserName also going from beta to stable. Will allow apps and workloads to run in k8s well
        *   Lifecycle management and deployment:
            *   Kubeadm support going to beta, Cluster API support initial experimental support landed for Azure
            *   Looking to fork WINS to work on functionality. Should allow run join workflows a bit cleanly
            *   These all should help things run a bit closer to how things run on Linux
            *   Hoping to graduate kubeadm soon if all goes well
            *   Also hoping to provide experimental support for the Azure provider
        *   Trends
            *   Working on CRI-ContainerD to provide a path to run containers on CRI-ContainerD over Docker
            *   The key work that's needed to run containers through CRI expected to release with ContainerD 1.4
            *   Looking forward to keep this consistent across Linux and Windows going forward.
            *   Working also on developing a CSI proxy, which should allow using the proxy to handle privilege limitations
            *   CSI proxy work is being done out of tree to enable the CSI migration initiative to move out
            *   For 1.18, looking at scaling issues:
                *   With horizontal pod autoscaling, there were some issues with stats coming back from Docker. Working on that so it works better with Prometheus and others
                *   CPU limit honoring has been adjusted so it runs better under load
            *   There's been a trend that Windows tests have been failing for a couple days. Working with SIG testing to remove the manual testing, still working on getting some automation finished. [📣CONTRIBUTOR ACTION] Please review the PRs linked in the slides. All help encouraged!
                *   [#85599](https://github.com/kubernetes/kubernetes/pull/85599) [#76838](https://github.com/kubernetes/kubernetes/pull/76838) [#77398](https://github.com/kubernetes/kubernetes/pull/77398) [#77269](https://github.com/kubernetes/kubernetes/pull/77269) [#81226](https://github.com/kubernetes/kubernetes/pull/81226) [#88248](https://github.com/kubernetes/kubernetes/pull/88248) [#88249](https://github.com/kubernetes/kubernetes/pull/88249)
            *   [LogMonitor open-sourced by Microsoft](https://github.com/microsoft/windows-container-tools/tree/master/LogMonitor) - makes it easy to take logs coming from a few places and copy them to stdout to allow kubectl logs, FluentD and others to scrape the logs properly
                *   Demo at Kubecon ([deck](https://sched.co/UagU), [video](https://youtu.be/t_e8SSkpdxU?t=1471))
        *   Notables from 1.17:
            *   Introduced RuntimeClass scheduler. Should make things easier for Windows devs to define the aspects of the workload you're working on with the podspec ([doc](https://kubernetes.io/docs/setup/production-environment/windows/user-guide-windows-containers/#simplifying-with-runtimeclass))
            *   Also added new labels for Windows nodes, allowing major, minor, and patch builds. Much more compatible with OS.
        *   Plans
            *   Continuing major investments: kubeadm and cluster API support.
            *   Investing in more storage options
            *   Compatibility at runtime level with Linux
        *   Slides

            [https://docs.google.com/presentation/d/1nSBVDp7IuyzpakvLvJYtQUsOAJd54iZuXP1pxJR1Pq8/edit?usp=sharing](https://docs.google.com/presentation/d/1nSBVDp7IuyzpakvLvJYtQUsOAJd54iZuXP1pxJR1Pq8/edit?usp=sharing)

    *   SIG MultiCluster [Paul Morie] Confirmed
        *   Last cycle:
            *   Discussing future of the SIG and areas to collaborate. [CONTRIBUTOR ACTION] Please read the multicluster services API proposal (linked in the slides)
            *   [CONTRIBUTOR ACTION] Also hunting for Kubefed maintainers! Open need!
        *   Participation is key to determine the right problems. [CONTRIBUTOR ACTION] Talk to us! Let us know what you're working on outside of the community, and show us your demos.
        *   Kubefed status
            *   Seeking maintainers
            *   How to get involved:
                *   Please feel welcome to reach out to pmorie on Slack. He'll help!
                *   There are biweekly meetings; please join!
        *   Slides

            [https://docs.google.com/presentation/d/1zjeLm_KskJwn60guai0ZofNH5OJhq4rRdasWQRLo3Kw/edit](https://docs.google.com/presentation/d/1zjeLm_KskJwn60guai0ZofNH5OJhq4rRdasWQRLo3Kw/edit)

    *   SIG Auth  [Mikedanese / Mo Khan] Confirmed
        *   Last cycle
            *   Adopted a new subproject: secrets store CSI driver
            *   Donated by {???}, integrates with external secret stores. Alt to secrets volume.
            *   A lot of users wanted deeper integration with Vault and others. So new subproject!
            *   Also doing a lot of work around certs. Certs API has been in beta for a long time; hasn't made much progress, so working on migrating it to GA. Wrote a retroactive KEP to help get to GA, and using that to organize the GA path. Includes support for multiple signers now. (Needed to reimplement the CA to include multiple IDs, so designed support for multiple signers.) Migrating clients to enable dynamic rotation of certs.
            *   Better performance! There's a whole list of improvements on the slide. Interesting ones according to the presentation: Token caching improvements to make the node authorizer (critical) much faster. Also added monitoring around latency, cache performance, authenticator use. Also added to k8s scalability prow scalability limits for auth.
        *   Upcoming cycles
            *   Keep improving scalability of storage encryption. Identified some problems with architecture that they want to address.
            *   Continue work on Certs API GA work. Number of PRs open.
            *   PodSecurityPolicy is on the radar. It's a hodpodge of features, so discussing ways to better contain constraints. See proposal linked in slides. Still discussing.
            *   Working on new service token support out to GA
                *   There were some issues with legacy tokens and compatibility issues that they are addressing.
        *   How this affects you
            *   Better performance. Better security.
            *   [CONTRIBUTOR ACTION] If any of this breaks things for anyone, please let the SIG know!
        *   [CONTRIBUTOR ACTION] Feel free to join the meetings! File bugs! Help improve monitoring! There's a lovely list of first-issue needs open. Also, if you use any of the clients, please take a look at those issues and help contribute. The SIG would love help and would love to hear from you.
        *   Slides

            [https://docs.google.com/presentation/d/1HBMqr5V79S8BSrSMAxPdQiyyCL9byBBWj2D4WrR3hPY/edit#slide=id.g401c104a3c_0_0](https://docs.google.com/presentation/d/1HBMqr5V79S8BSrSMAxPdQiyyCL9byBBWj2D4WrR3hPY/edit#slide=id.g401c104a3c_0_0)

*   [ 0:33 ] **Announcements **
    *   [Contributor Summit for Amsterdam Schedule Announced](https://kubernetes.io/blog/2020/02/18/contributor-summit-amsterdam-schedule-announced/)
    *   Next month expect updates from SIGs: Instrumentation, Storage, Service Catalog, Steering Committee, and hopefully Code of Conduct Committee - hosted by the inimitable [Matt Broberg](https://twitter.com/mbbroberg) (woot woot!)
    *   ** **Shoutouts this month (Check in #shoutouts on slack) ** **
        *   Laura Santamaria- Shoutouts to @castrojo and @marky.jackson for all the help getting up and running for hosting my first community meeting today! Appreciate all y'all do
        *   Markyjackson- Shout out to @nimbinatus for hosting her first community meeting!
        *   Jeremy Rickard- Shout out to @oikiki and @palnabarun for last minute quality control checks on the 1.18 enhancements tracking sheet
        *   Samudrala Vamshi- shoutout to @castrojo and @markyjackson for helping on my first PR to this  community and appreciate all they do
        *   Markyjackson - I would like to give @Vishakha Nihore a shout out for the amazing work she is doing on the contributor experience side. She comes to this project via #outreachy-apps and is absolutely amazing! Thank you @Vishakha Nihore
        *   Vishaka Nihore - Shout out to @mrbobbytables to all the help he provided me whenever I bugged him, even when the mistake was just a flake. Also not to mention his great in depth reviews on my PRs !
        *   Taylor Dolezal- @mrbobbytables - you're a champ through and through
        *   Benjamin Elder- shoutout to @amwat for implementing podman support in #kind, I know this will make some people happy
        *   Guinevere Saenger - shoutout ot @alisondy for KILLING it on the new contributor workshop content!
        *     Benjamin Elder- thanks @pohly for fixing blockfs test flakiness! ([https://github.com/kubernetes/kubernetes/issues/87953](https://github.com/kubernetes/kubernetes/issues/87953))
        *   Codyc - I just want to give a big hug, honk and shoutout to @markyjackson for all he is doing for the community.. I love you brotha
        *   Jason - Huge shoutout to @naadir who has been helping guide and shepard multiple PRs related to the webhook changes needed for multi-tenancy in cluster-api providers for v1alpha3
        *   Antonio- Thanks @bentheelder for your great efforts on keeping a healthy CI , impressive work, huge shoutout to  him [https://prow.k8s.io/?job=pull-kubernetes-e2e-kind](https://prow.k8s.io/?job=pull-kubernetes-e2e-kind)


### January 16, 2020 - [recording](https://youtu.be/Wp7DPvmosu0)



*   **Moderators**:  Laura Santamaria [LogDNA/SIG Contribex]
*   **Note Taker**: Bob Killen [University of Michigan/Contribex]
    *   Subscribe to [this thread](https://discuss.kubernetes.io/t/kubernetes-weekly-community-meeting-notes/35) to get these notes in your inbox
*   [ 0:00 ]** Release Updates **
    *   Current Release Development Cycle  [Bob Killen]
        *   Tuesday, January 28: Week 4 - [Enhancements Freeze](https://github.com/kubernetes/sig-release/blob/master/releases/release_phases.md#enhancements-freeze)
            *   Implementable state
            *   Have a test plan
        *   Thursday, March 05: Week 9 - [Code Freeze](https://github.com/kubernetes/sig-release/blob/master/releases/release_phases.md#code-freeze)
        *   Monday, March 16: Week 11 - Docs must be completed and reviewed
        *   Tuesday, March 24: Week 12 - Kubernetes v1.18.0 released
    *   Patch Release Updates [https://git.k8s.io/sig-release/releases/patch-releases.md](https://git.k8s.io/sig-release/releases/patch-releases.md)
        *   1.17.1 released Jan. 14
        *   1.16.5 coming today Jan. 16
        *   1.15.8 coming today Jan. 16
        *   1.14.11 coming today Jan.16 (to fix an upgrade scenario for 1.15)
        *   A series of bugs have been identified in how the next beta tag is applied on these branches.  For example when “v1.17.1” is tagged and released we also mark the branch with a tag “v1.17.2-beta.0".    The bugs root cause goes back many years in the design and implementation of the “anago” tool used to build and release, but are partially corrected now.  A complete fix likely will come first at the point we replace the “anago” tool.
        *   Next patch releases target Feb. 11 (see: [https://github.com/kubernetes/sig-release/pull/954](https://github.com/kubernetes/sig-release/pull/954))
*   [ 0:00 ] **SIG Updates**
    *   SIG Leads, check out this set of [recommended topics to cover](https://github.com/kubernetes/community/blob/master/events/community-meeting.md#sig-updates) during your update. [Slide template if you need it](https://docs.google.com/presentation/d/1-nTvKCiqu9UvFYUeM6p6RIqHS5-H-u3_x-V4xj_eIWo/edit#slide=id.g401c104a3c_0_0), please also check the [SIG Update Schedule](https://docs.google.com/spreadsheets/d/1adztrJ05mQ_cjatYSnvyiy85KjuI6-GuXsRsP-T2R3k)!
    *   SIG Cloud Provider [Walter Fender]
        *   [Slides](https://docs.google.com/presentation/d/1NX2TnKcqGm_Pg54n690gmd-HCYxsk6agnQsBIrDBNiU/edit#slide=id.g401c104a3c_0_0)
        *   Promoted Node Zone/Region Topology Labels to GA
            *   failure-domain.beta.kubernetes.io/* deprecated
            *   beta.kubernetes.io/instance-type deprecated
        *   Upcoming Cycles
            *   API Server Network Proxy alpa with goal to promote to GA in the upcoming cycles
            *   Extract cloud provider dependencies from the core repo
            *   Generate a controller migration lock mechanism for moving controllers safely between controller managers
            *   Better support for providerless builds for cloud providers who are working out of tree
            *   Improve the tooling/documentation around cloud controller manager and per cloud repos
            *   Targeting removal of in-tree cloud providers by the 1.21 release
        *   What impacts you:
            *   In-tree cloud provider e2e tests are strong test signal, need to figure out how best to transition to out-of-tree
        *   New KEP template to add new cloud providers.
        *   Cloud Provider Extraction WG
            *   Slated for extraction with 1.21
            *   Cloud controller manager is green
    *   SIG Autoscaling [Marcin Wielgus]
        *   Slides
        *   Cluster Autoscaler:
            *   Switching from using raw scheduler predicates to Scheduling Framework. This will improve behavior of CA in various corner cases related to zone-specific storage and affinity/anti-affinity.
            *   Added support for Packet
            *   Improved performance/scalability.
        *   Vertical Pod Autoscaler:
            *   Graduating VPA api to GA soon
        *   Horizontal Pod Autoscaler:
            *   Expanded API to control how fast and how many pods are added on scale-up and scale-down.
            *   Added support for scale to 0 (currently flag-gated).
    *   SIG Scheduling  [Abdullah Gharaibeh]
        *   [Slides](https://docs.google.com/presentation/d/1H27SDMqkzq8zCRveWWtK5g9hCAomKbrzTTVZ5r4h6Xo/edit#slide=id.g401c104a3c_0_0)
        *   What we did last cycle:
            *   [Scheduling Framework has hit milestone 1](https://github.com/kubernetes/kubernetes/issues/83554)
                *   Finalized framework implementation
                *   Wrapped existing predicates and priorities functions in plugins
                *   Added a translation layer from predicate/priority “policies” into Plugin configurations
            *   Performance improvements
                *   [2x improvements for pod scheduling latency (excluding bindings) on 5k clusters](https://bit.ly/2uGalKr)
                *   4x improvements for preferred pod (anti)affinity
            *   Improved observability: new latency, traffic and saturation metrics
                *   Scheduling latency
            *   Features graduated to GA in 1.17
                *   Schedule DaemonSet Pods
                *   Taint nodes by condition
            *   Performance Improvements
                *   Large Scale Clusters
                *   4x improvement in preferred pod affinity
                *   Scheduling latency
                *   How many pods via qued
        *   Plans for upcoming cycles
            *   [Scheduling framework milestone 2](https://github.com/kubernetes/kubernetes/issues/85822)
                *   Move predicates and priorities code to run natively as plugins
                *   Remove external dependencies on predicates (DS controller, Kubelet and Cluster Autoscaler)
                *   Declare Policy API deprecated, the new Plugins API in ComponentConfig is the replacement
            *   [Support schedulers with multiple plugin configurations](https://github.com/kubernetes/kubernetes/issues/85737)
                *   Pod.Spec.SchedulerName will be used to inform the scheduler which canned configuration to use
            *   [Simplify integration with the Cluster Autoscaler](https://github.com/kubernetes/autoscaler/issues/2645)
            *   Further improve scheduling performance
            *   Graduate pod topology spread to Beta (in Alpha since 1.16)
        *   Leadership position changes
            *   Bobby Salamat stepped down as co-chair
            *   Abdullah Gharaibeh new sig co-chair
    *   SIG Scalability [Matt Matejczyk]
        *   [Slides](https://docs.google.com/presentation/d/1T_et57l52gueQSWEKBamy9jDcXVx0Vle6dbt4jIX2SU/edit?usp=sharing)
        *   What we did last cycle
            *   Improved Scalability and Performance Tess
                *   Add support for more kubernetes concepts such as DaemonSets, StatefulSets, Secrets etc.
                *   ClusterLoader2:
                    *   improved testsuite
                    *   Better crashloop detection
                    *   HA support
                *   Build more scale tests into the release branches
                *   Pod throughput tests (containerD vs Docker)
            *   GuardingAgainst Performance Regressions
                *   [Scalability Approval Process](https://github.com/kubernetes/community/blob/master/sig-architecture/production-readiness.md)
                *   [Mitigated Golang 1.13 Performance regression](https://github.com/golang/go/issues/32828)
                *   [Scalability Regressions & Bugs Documents](https://docs.google.com/document/d/1_mqv_T7i5k7_HgcQihEuFdq7ZCIf3AAGyAo9axzdAGI/edit)
            *   Performance Improvements
                *   Watch Serialization Mechanism Improvements
                *   Core Components Improvements:
                    *   NodeLifeCycleController
                    *   GC Controller
                    *   TaintManager
                *   Watch Bookmarks went to GA
                *   KEP for immutable secrets
        *   Plans for upcoming cycles
            *   Kubernetes Scalability Definition
                *   Finalizing existing WIP scalability SLI/SLOs
                *   Updating scalability envelope (thresholds)
                *   Work on hardening and extending the scalability definition
            *   Scalability & Performance Tests
                *   Covering more Kubernetes concepts
                *   Work on Kubemark v2: better cluster simulations
                *   Add other tests: HA, upgrade, chaos etc
            *   Bottleneck Detection & Performance Improvements
                *   [Golang 1.14 release verification](https://github.com/kubernetes/enhancements/pull/1369)
                *   [Immutable Secret Implementation](https://github.com/kubernetes/enhancements/pull/1369)
                *   [Implement Consistent Reads from Cache in etcd](https://github.com/kubernetes/enhancements/pull/1404)
            *   How these plans affect you
                *   **Scalability approval process**
                    *   Will need to work with KEP owners to validate new features
                *   Extending SLI/SLO Coverage
                    *   We’ll be reaching out to help us understand what is important to the users and community
                *   Notable Regressions
                    *   Kubernetes v1.17.0 is vulnerable to [#86483](https://github.com/kubernetes/kubernetes/issues/86483) that can break large clusters on master restart
*   [ 0:00 ] **Announcements **
    *   Contributor Survey: [https://www.surveymonkey.com/r/VYRJZ5G](https://www.surveymonkey.com/r/VYRJZ5G)
    *   Let SIG Contribex know if this new format worked for you by pinging us in Slack.
    *   ** **Shoutouts this month (Check in #shoutouts on slack) ** **
        *   Rawkode -  Awesome props to @alculquicondor for jumping in at very late notice and getting us help with the release blogs for 1.17 :tada:
        *   Sascha - Big pre-release shoutout to @macintoshprime regarding his release notes efforts! That’s a lot of work, big kudos to you and your team!
        *   Zacharysarah - Shoutouts to @mrbobbytables and @gsaenger today for resolving a particularly thorny docs release PR!
        *   Gsaenger - And @Damini Satya !
        *   Nikhita - shoutout to @liggitt @sttts and @dims for tirelessly going through the back and forth the past week to get v0.17.0 tags shipped for published (staging) repos :tada:
        *   @vincepri - Shoutout to @ncdc for the great high quality effort to improve Cluster API documentation book!
        *   Bentheelder - shoutout to @timothysc for all your work in sig-testing, particularly in #testing-commons, and for your leadership in stepping down when you needed to
        *   Bentheelder - shoutout to @yasker for all the help and patience with PVCs in sigs.k8s.io/kind and your work on github.com/rancher/local-path-provisioner, looking forward to github.com/kubernetes-sigs/kind/pull/1157
        *   Paris - shouts to @jberkus @idealhack @cblecker @maria @markyjackson @mrbobbytables @spzala and many others in contribex for their thoughtful review of the upcoming contributor experience survey
        *   Paris - shoutouts to the kubernetes blog team (#sig-docs-blog) for all of their work reviewing PRs and working with contributors on that workflow so our end users and other community members can have great content on the blog.
        *   Markyjackson - Shout out to @mrbobbytables for patiently helping me fix a git problem. Really appreciate you
        *   Nimbinatus/Laura - Shoutouts to @castrojo and @marky.jackson for all the help getting up and running for hosting my first community meeting today! Appreciate all y'all do
        *   Jorge Castro - Huge shoutout to @parispittman for 2 years of service as cochair of SIG Contributor Experience!


### December 12, 2019 ([recording](https://www.youtube.com/watch?v=XQaC5ke9SHc))



*   **Moderators**: Chris Short (Shadow: Vamshi Samudrala)
*   **Note Taker:**
*   **[1.17 Retro](https://docs.google.com/document/d/1AtZ_81F3E4y_04Gx31mnG5w6AG3E0AXDOuLU7RcUIso/edit?usp=sharing)**


### December 5, 2019 (recording)



*   **Moderators**: Jeffrey Sica [Red Hat, SIG-Contribex/Release/UI]
*   **Note Taker**: Jordan Liggitt / Bob Killen
*   [ 0:00 ]** Release Updates **[Guinevere Saenger - Release Lead]
    *   1.17 release
        *   1.17.0 targeting Monday, December 9th
        *   Generally looking good, might have one bugfix in progress
        *   Primary need is for SIG review of release notes
    *   Patch releases ([schedule](https://github.com/kubernetes/sig-release/blob/master/releases/patch-releases.md)):
        *   Cherry pick deadline tomorrow, Dec. 6 ahead of:
            *   [1.14.10](https://groups.google.com/forum/#!topic/kubernetes-dev/hESFCjjeqWA) -- LAST 1.14 PATCH RELEASE
            *   [1.15.7](https://groups.google.com/forum/#!topic/kubernetes-dev/n62_NQcSveY)
            *   [1.16.4](https://groups.google.com/forum/#!topic/kubernetes-dev/hbNwT260eIU)
        *   Release target Wed. Dec. 11
*   [ 0:00 ] **SIG Updates**
    *   SIG-PM Punted (Per Augustus)
    *   SIG-Arch [Liggitt][[Slides](https://docs.google.com/presentation/d/185-xT9ytws4SZ6wShe3-PzPudc7EbJzdjD1Wjema8YQ/edit#slide=id.g338ac0a8b6_0_27)]
        *   Brian Grant, Jaice SInger DuMars and Matt Farina have moved to emeritus
        *   Davanum Srinivas, Derek Carr, and John Belamaric are the current chairs
        *   Focus on Quality / Stability
            *   [Started production readiness review subproject](git.k8s.io/community/sig-architecture/production-readiness.md)
            *   [Started surveying cross-cutting technical debt issues](github.com/orgs/kubernetes/projects/35)
            *   [Proposal to limit beta APIs](https://github.com/kubernetes/enhancements/pull/1266)
                *   Beta apis have had an issue being stuck in beta for extended periods of time.
                *   People start to depend on non-ga beta things
                *   Timebox amount of time a feature is allowed to stay in beta.
            *   [Plan to set up CI to ensure conformance coverage](https://github.com/kubernetes/enhancements/pull/1306)
            *   [Plan to eliminate dependencies on non-GA features in conformance](http://git.k8s.io/enhancements/keps/sig-architecture/20191023-conformance-without-beta.md)
        *   Started a “KEP reading group”
            *   Team collectively reads KEPs and discusses them in meetings
        *   Code organization
            *   Improve usability of versioned go modules
                *   [github.com/kubernetes/enhancements/pull/1350](https://github.com/kubernetes/enhancements/pull/1350)
                *   [github.com/kubernetes/publishing-bot/pull/210](https://github.com/kubernetes/publishing-bot/pull/210)
            *   Continued reduction of external dependendencies
        *   How these plans affect you
            *   Stability / Quality
                *   Prefer completing (or removing) languishing non-GA features in your SIG
                *   Ensure complete test coverage of features during beta stage of development
                *   Plan to promote tests to conformance as part of feature GA
            *   Help Wanted
                *   Help identify areas of technical debt
                *   Give feedback on areas that should be considered for production readiness review
                    *   Reach out to John Belamaric
    *   Security Response Committee (PST) [Liggitt][[Slides](https://docs.google.com/presentation/d/1tWwBOEm66pIiYuQ_XPhvZL3Qt2pSq5loFZFgjerrUzk/edit#slide=id.g6c0dff12cb_0_0)]
        *   SRC responds to security notifications
        *   Members include:
            *   CJ Cullen (@cjcullen)
            *   Joel Smith (@joelsmith)
            *   Jonathan Pulsifer (@jonpulsifer)
            *   Jordan Liggitt (@liggitt)
            *   Luke Hinds (@lukehinds) (new!)
            *   Brandon Philips (@philips)
            *   Tim Allclair (@tallclair)
        *   Recent Efforts
            *   Continuing to setup a bug bounty program with HackerOne
            *   Evaluated 29 reports of security related issues resulting in 5 CVEs
                *   Medium
                    *   CVE-2019-11248
                    *   CVE-2019-11249
                    *   CVE-2019-11250
                    *   CVE-2019-11255
                *   High
                    *   CVE-2019-11253
*   [ 0:00 ] **Announcements **
    *   **_One more Community Meeting before EOY!_**
    *   **_Kubecon EU CFPs closed YESTERDAY GOOD LUCK_**
    *   
    *   ** **Shoutouts this week (Check in #shoutouts on slack) ** **
        *   Jeremy - Shoutout to @bentheelder for the real time kind troubleshooting for the new contributor workshop
            *   Chris Short - You’re my hero @bentheelder
        *   Elana - shout out to @jeefy and @mrbobbytables for the best goose game ever. Honk.
        *   Ben - Shoutout to @cblecker for adding #kind to homebrew! Thank you Christoph!
        *   Paris - Shoutout to the fabulous kubernetes contributor summit team! Thanks for making the show in San Diego a memorable one. Can't wait to see what's next for Amsterdam!


### November 7, 2019 ([recording](https://youtu.be/JP9k9bcl6_c))



*   **Moderators**: Marky Jackson [Sysdig/SIG Contributor Experience/SIG Release]
*   **Note Taker**: [Jorge Castro/SIG Contributor Experience/VMware]
*   [ 0:00 ]** Release Updates **[Guinevere Saenger - Release Lead]
    *   1.17 release
        *   “Calm before the storm” - KubeCon prep, lots of meetings
        *   Everyone filing an enhancement MUST file a docs PR for it by TOMORROW
        *   Code freeze next week, 14 Nov, everything afterwards will be a cherry pick
        *   This tuesday, first beta of 1.17
    *   Patch releases ([schedule](https://github.com/kubernetes/sig-release/blob/master/releases/patch-releases.md)):
        *   Cherry pick deadline tomorrow, Nov. 8 ahead of:
            *   [1.14.9](https://groups.google.com/d/topic/kubernetes-dev/iPWVGsVP4iQ/discussion)
            *   [1.15.6](https://groups.google.com/d/topic/kubernetes-dev/_QL4KyVsCac/discussion)
            *   [1.16.3](https://groups.google.com/d/topic/kubernetes-dev/9oXvqwVbeU0/discussion)
        *   Release target Wed. Nov. 13
*   [ 0:00 ] **SIG Updates**
    *   wg-LTS [@tpepper]: [slides](https://docs.google.com/presentation/d/12tzP3scecY-r-c7GItcOGAC41ZpMBXdBcuT5a7cl-n0/edit?usp=sharing)
    *   wg-k8s-infra [@bartsmykla]: [slides](https://docs.google.com/presentation/d/1-sjO6SiyKoWp5KMFHoTqi5uZNUpN6FBeEcELgws6BhU)
*   [ 0:00 ] **Announcements **
    *   **_This is the last community meeting until December 5th_**
    *   **_Happy Kubecon and happy thanksgiving_**
    *   **_Don’t forget to [register for the contributor summit](https://events19.linuxfoundation.org/events/kubernetes-contributor-summit-north-america-2019/register/)!_**
    *   ** **Shoutouts this week (Check in #shoutouts on slack) ** **
        *   Chris Short gave a huge shoutout to[ @castrojo](https://kubernetes.slack.com/team/U1W1Q6PRQ) and[ @jeefy](https://kubernetes.slack.com/team/U5MCFK468) for getting me all set to stream community meetings. So helpful and kind (even when I forget things)!
        *   Chris Blecker gave a shoutout to [@liggitt](https://kubernetes.slack.com/team/U0BGPQ6DS) and [@bentheelder](https://kubernetes.slack.com/team/U1P7T516X) for their help in getting us upgraded to go1.13. It was a huge effort!
        *   Paris gave a shoutout to everyone on the kubecon planning stretch especially the wonderful contributor summit events team
