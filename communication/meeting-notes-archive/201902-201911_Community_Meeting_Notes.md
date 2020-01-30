## November 7, 2019 ([recording](https://youtu.be/JP9k9bcl6_c))


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
*   [ 0:00 ] **📣Announcements 📣**
    *   **_This is the last community meeting until December 5th_**
    *   **_Happy Kubecon and happy thanksgiving_**
    *   **_Don’t forget to [register for the contributor summit](https://events19.linuxfoundation.org/events/kubernetes-contributor-summit-north-america-2019/register/)!_**
    *   **👏 **Shoutouts this week (Check in #shoutouts on slack) **👏**
        *   Chris Short gave a huge shoutout to[ @castrojo](https://kubernetes.slack.com/team/U1W1Q6PRQ) and[ @jeefy](https://kubernetes.slack.com/team/U5MCFK468) for getting me all set to stream community meetings. So helpful and kind (even when I forget things)!
        *   Chris Blecker gave a shoutout to [@liggitt](https://kubernetes.slack.com/team/U0BGPQ6DS) and [@bentheelder](https://kubernetes.slack.com/team/U1P7T516X) for their help in getting us upgraded to go1.13. It was a huge effort!
        *   Paris gave a shoutout to everyone on the kubecon planning stretch especially the wonderful contributor summit events team 


## October 31, 2019



*   **Moderators**: Marky’s assistant Hammy ;-) [SIG Contributor Experience/Release]
*   **Note Taker**: Chris Short
*   [ 0:00 ]** Release Updates **[Guinevere Saenger - Release Lead]
    *   1.17.0-beta.0 released this Tuesday 10/29
    *   1.17 release branch created
        *   All changes to master will be fast forwarded nightly into the 1.17 branch
    *   **CODE FREEZE IS COMING NOVEMBER 14**
        *   after Code Freeze, all approved enhancements work will need to follow cherry-pick process to be merged into the 1.17 branch
    *   1.13 jobs are being removed
*   Patch Release Updates
    *   Next target date 11/13 for all supported branches [https://github.com/kubernetes/sig-release/blob/master/releases/patch-releases.md](https://github.com/kubernetes/sig-release/blob/master/releases/patch-releases.md)
*   
*   [ 0:00 ] **SIG Updates**
    *   SIG Release [Stephen Augustus]
        *   Improved feedback loops between SIG Release and SIG Scalability
        *   Emeritus advisor is awesome
        *   More diversity of all kinds in the Release Teams
        *   Improvements in automation across the board
        *   SIG Release needs more shadows
        *   People are improving test coverage on their features
        *   Release Engineering subproject has started in earnest
        *   Test cleanup and deletion continues
        *   Release Managers Group
            *   [https://github.com/kubernetes/sig-release/blob/master/release-managers.md](https://github.com/kubernetes/sig-release/blob/master/release-managers.md)
            *   Dotted line to Product Security Committee
        *   Release Engineering
            *   Onboarding process improvements
            *   Wiring Release Engineering jobs in CI
            *   Doc cleanups
            *   Working on getting staging/release process into CI
            *   Viewer access to GCP
            *   k/release tooling is getting rewritten in Go and one tool has already been deployed
            *   deb/rpm packaging tools are being built and awesome-ified
            *   Hyperkube out-of-tree in progress
            *   Codebase walkthroughs!!!
        *   Watch for announcements
        *   Pay attention to CI Signal
        *   Be mindful of 1.17 schedule dates
        *   We'll be at KubeCon!
*   [ 0:00 ] **📣Announcements 📣**
    *   **_Don’t forget to [register for the contributor summit](https://events19.linuxfoundation.org/events/kubernetes-contributor-summit-north-america-2019/register/)!_**
    *   **👏 **Shoutouts this week (Check in #shoutouts on slack) **👏**
        *   @dims gave a shoutout to @bartsmykla for setting up / running the on-boarding call for 70+ folks for wg-k8s-infra
        *   Daniel Lipovetsky [@dlipovetaky] gave a shoutout saying,Thank you, thank you, thank you to [@neolit123](https://kubernetes.slack.com/team/U83J4CS3S) [Lubomir Ivanov] for always taking the time to help and mentor. You have been there for me and for many others on what seems like everywhere from k/k, to kubeadm, to docs, and everything in between.
        *   @markyjackson gave a shout out to [@gsaenger](https://kubernetes.slack.com/team/U4H2QU3DW) [@chrisshort](https://kubernetes.slack.com/team/U2YGXSD9B) and [@rael](https://kubernetes.slack.com/team/UHCJ61V2T) for getting together to make the NCW awesome and for being such fine peoples to work with 


## October 24, 2019 ([recording](https://youtu.be/cnkxd0_MpJg))



*   **Moderators**:  Jonas Rosland [VMware/SIG Contributor Experience/Release]
*   **Note Taker**: Thiscould B. You [Company/SIG]
    *   Subscribe to [this thread](https://discuss.kubernetes.io/t/kubernetes-weekly-community-meeting-notes/35) to get these notes in your inbox
*   [ 0:00 ]** Release Updates**
    *   Current Release Development Cycle  [Guinevere Saenger - Release Manager]
        *   1/17.alpha.3 released this Tuesday 10/22
        *   All Enhancement exceptions are merged and tracked
        *   Begin turnaround for release branch creation next week: removal of 1.13 jobs, create 1.17 jobs, create 1.17 release branch, cut the first 1.17 beta
        *   Lots of work from SIG scalability and the CI Signal team to capture scale job flakes early and find causes
    *   Patch Release Updates
        *   [https://github.com/kubernetes/sig-release/blob/master/releases/patch-releases.md](https://github.com/kubernetes/sig-release/blob/master/releases/patch-releases.md)
*   [ 0:00 ] **SIG Updates**
    *   SIG Usability [Tasha Drew @tasha]
        *   [Slides](https://docs.google.com/presentation/d/10BeDObYa2haF3d5k_BxXzzZcpI7ytkaLqCDZJPFkJ-c/edit#slide=id.g401c104a3c_0_0)
    *   WG Multitenancy [Tasha Drew @tasha]
        *   [Slides](https://docs.google.com/presentation/d/1OmhUwagXbO5SLRJglymYPIoo8goIvO096-SqpO9e5Zk/edit#slide=id.g401c104a3c_0_0)
    *   WG Apply Working Group [Jenny Buckley @jennybuckley]
        *   [Slides](https://docs.google.com/presentation/d/1pxtZlWwTQNXcdZ1VhzLpeiV01HB1fThsY3kMDh1tSg4)
    *   WG Machine Learning [punt till next week]
*   [ 0:00 ] **📣Announcements 📣**
    *   **_Don’t forget to [register for the contributor summit](https://events19.linuxfoundation.org/events/kubernetes-contributor-summit-north-america-2019/register/)!_**
    *   Please start [populating your schedules](https://twitter.com/Microwavables/status/1187130618368446464) on sched.org for KubeCon + CloudNativeCon
        *   This helps the planners select the right room size!
        *   You can [export this schedule and add it to your work calendar](https://kccncna19.sched.com/mobile-site) - makes for easier scheduling. Thanks to @dankohn for this tip!
    *   The format of this meeting is changing! 
        *   Moving to every other week in 2020, 2019 remains unchanged
    *   **👏 **Shoutouts this week (Check in #shoutouts on slack) **👏**
        *   [nikhita] shoutout to @bartsmykla for going above and beyond to make sure that all new folks in #wg-k8s-infra have an amazing onboarding experience! 
        *   [mrbobbytables] Big shoutout to @directxman12 for creating a contributor ladder for kubebuilder :heart: [https://github.com/kubernetes-sigs/kubebuilder/blob/master/docs/CONTRIBUTING-ROLES.md](https://github.com/kubernetes-sigs/kubebuilder/blob/master/docs/CONTRIBUTING-ROLES.md)
        *   [gsaenger] Shoutout to @castrojo for the world's most efficient community meeting!
    *   SIG Release and WG Machine Learning will be giving updates next week!


## October 17, 2019



*   **Moderators**:  Jorge Castro [VMware/SIG Contributor Experience]
    *   No video available, Jorge hit the wrong button on OBS. :(
*   **Note Taker**: First Last [Company/SIG]
    *   Subscribe to [this thread](https://discuss.kubernetes.io/t/kubernetes-weekly-community-meeting-notes/35) to get these notes in your inbox
*   [ 0:00 ]** Release Updates**
    *   Current Release Development Cycle  [Guinevere Saenger - Release Manager]
        *   Enhancements Freeze was this past Tuesday, 15 October
        *   Two exceptions filed
        *   We have 44 enhancements tracked: alpha: 11, beta: 13, stable: 20
        *   1.17.0-alpha.2 released on Oct.15
        *   1.17.0-alpha.3 planned for Oct.22
    *   Patch Release Updates
        *   All branches released 15 October
*   [ 0:00 ] **SIG Updates**
    *   SIG Leads, check out this set of [recommended topics to cover](https://github.com/kubernetes/community/blob/master/events/community-meeting.md#sig-updates) during your update. [Slide template if you need it](https://docs.google.com/presentation/d/1-nTvKCiqu9UvFYUeM6p6RIqHS5-H-u3_x-V4xj_eIWo/edit#slide=id.g401c104a3c_0_0), please also check the [SIG Update Schedule](https://docs.google.com/spreadsheets/d/1adztrJ05mQ_cjatYSnvyiy85KjuI6-GuXsRsP-T2R3k)!
    *   SIG API Machinery [David Eads] (confirmed)
        *   [Slides](https://docs.google.com/presentation/d/17Nc5jmoIYyCIvKhfXlP6jESb28Q4E7bL-EA-zMEgB00/edit#slide=id.g439b3a360b_0_1) 
*   [ 0:00 ] **📣Announcements 📣**
    *   Don’t forget to [register for the contributor summit](https://events19.linuxfoundation.org/events/kubernetes-contributor-summit-north-america-2019/register/)!
    *   **👏 **Shoutouts this week (Check in #shoutouts on slack) **👏**
        *   **@jdetiber - **shoutout to @cblecker for adding a `/honk` command to prow
        *   **@gsaenger  -** shoutout to @markyjackson for being such a friendly community meeting host!
    *   SIG Usability, WG Apply, and WG Machine Learning will be giving updates next week!


## October 10, 2019 - ([recording](https://www.youtube.com/watch?v=ggCY_PqOdyA&feature=youtu.be))



*   **Moderators**:  Marky Jackson [ Sysdig/SIG Contribex]
*   **Note Taker**: Bob Killen
*   [ 0:00 ]** Release Updates**
    *   Current Release Development Cycle  [Guinevere Saenger - Release Manager]
        *   We’re in Week 3…
        *   Enhancement Freeze is next Tuesday (Oct. 15). Enhancements must be in by 5PM PT.
        *   1.17.0-alpha.2 scheduled release Tuesday Oct.15
    *   Patch Release Updates
        *   1.16.1 released 1 October
        *   Next patch releases (all branches) scheduled for 15 October
        *   LAST release of 1.13.x
*   [ 0:00 ]** SIG Updates**
    *   WG Security Audit [Jay Beale]
        *   Slides:
        *   [https://docs.google.com/presentation/d/1z4voQDTejEdU2gwNM30gdRxF0M9dAw-BhCKx5VitCWk/edit#slide=id.p1](https://docs.google.com/presentation/d/1z4voQDTejEdU2gwNM30gdRxF0M9dAw-BhCKx5VitCWk/edit#slide=id.p1) 
        *   What we did last cycle
            *   Led the first in a series of Kubernetes security audits
                *   Choose vendors
                *   Gave direction to focus effort
                *   Participated in the threat modeling work that will be used for future releases of Kubernetes
                *   Performed technical editing on the report
                *   Worked on producing reusable artifacts
            *   Complementary efforts to the bug bounty program
            *   Threat model breakdown
                *   Focus on 8 critical components
                    *   Kube-apiserver
                    *   Etcd
                    *   Kube-scheduler
                    *   Kube-controller-manager
                    *   Cloud-controller-manager
                    *   Kubelet
                    *   Kube-proxy
                    *   Container Runtime Interface
            *   Threat model highlighted recommendations
                *   Provide auditing information in a unified fashion to allow a trace of the user’s actions through the system
                *   Warn users who configure a security control that will not be enforced
                    *   Network policies and pod security policies can silently fail.
                *   Require transport encryption w/cert verification
                    *   Multiple components use http
                    *   Multiple components elect not to verify cert validity
                *   Prevent node compromises from leading to cluster-compromises
                    *   Host access gives access to cli arguments, logs etc
                *   Separate privilege levels among controllers
            *   Vulnerability research during cycle
                *   Discovered 37 vulnerabilities
            *   Vulnerability highlights
                *   Non authenticated HTTPS connections
                *   Cert revocation unsupported
                *   PSP Bypass (hostPath va PVs)
                *   TOCTOU Race condition in Kubelet
                *   Kubectl cp directory traversal
                *   System logs containing secrets
            *   Recommendation Highlights
                *   Replace the many cases of logic reimplementation with central libraries
                *   Ease security configuration (particularly defaults)
                *   Improve code documentation around external dependencies
                *   Continue development of security features
            *   Security Audit report [link from report in k/community]
        *   Next cycle:
            *   Plan next security audit
            *   Move towards more secure defaults
    *   SIG Testing [fejta]
        *   [https://docs.google.com/document/d/1uTcLhxM2HwDgtGOiIvlFfRWzQDTvii6qd_XASAubHlk/edit?ts=5d9e6825](https://docs.google.com/presentation/d/1agkyIyyPW-gJGMTF_F_uc0tZkuUFSLl1PUOq9SadNNQ/edit#slide=id.g401c104a3c_0_0)
        *   Last Cycle
            *   Testgrid configs now live alongside their associated prow jobs
            *   Automated the creation of jobs for the test-infra release team role
            *   Deployed new and improved monitoring/alerting stack ([monitoring.prow.k8s.io](https://monitoring.prow.k8s.io))
            *   Reusable verify checks in bazel rules
            *   KinD
                *   Smaller images from providerless kubernetes builds
                *   Release blocking IPv4 and IPv6 test coverage
                *   Provides 75% of pull-kubernetes-e2e-gce coverage without any cloud resources
            *   [TestGrid partially open sourced](https://github.com/GoogleCloudPLatform/testgrid)
        *   Next Cycle
            *   Establish test-infra SLOs
            *   Improve test-infra alerting to better detect and recover from outages
            *   Make KinD a blocking presubmit in k/k
            *   Automate image pushing on merge with a git-ops based promotion to prod method (working with #wg-k8s-infra)
            *   Help repos with preexisting bazel rules adopt reusable verify checks.
            *   Move prow out of test-ifnra into its own repo
            *   Enable in repo prowjob configurations
        *   How these upcoming changes affect you
            *   Help define more reusable verify checks
            *   Start thinking about how/whether your sig can move cloud provider dependencies out of k/k testing to release blocking postsubmits
    *   [ 0:00 ] **📣Announcements 📣**
        *   Announcement Foo #1
        *   **👏 **Shoutouts this week (Check in #shoutouts on slack) **👏**
            *   @jdetiber** **gave a shout to @dims for building out the e2e conformance tests using Cluster API and the GCP Provider
            *   @mrbobbytables gave a shoutout to the other Steering Election committee officials  [@briangrant](https://kubernetes.slack.com/team/U09R2JFE3) [@castrojo](https://kubernetes.slack.com/team/U1W1Q6PRQ) [@ihor.dvoretskyi](https://kubernetes.slack.com/team/U0CBHE6GM)   for putting in the work to make this year’s election possible! 
            *   @ihor.dvoretskyi gave a huge SHOUTOUT to [@mrbobbytables](https://kubernetes.slack.com/team/U511ZSKHD) - another election official!
            *   @cblecker gave a** **shout out to [@bentheelder](https://kubernetes.slack.com/team/U1P7T516X) and [@krzyzacy](https://kubernetes.slack.com/team/U22Q65CTG) for late night debugging on GCE test infra failures
    *   


## October 3, 2019 - ([recording](https://www.youtube.com/watch?v=yFKJgd1S7Zg&feature=youtu.be))



*   **Moderators**:  Jonas Rosland [VMware/SIG Contribex]
*   **Note Taker**: First Last [Company/SIG]
    *   Subscribe to [this thread](https://discuss.kubernetes.io/t/kubernetes-weekly-community-meeting-notes/35) to get these notes in your inbox
*   [ 0:00 ]** Steering Committee Election Results **[Dims]
    *   The following candidates will be joining @dims, @tstclair, and @spiffxp on the Steering Committee (in github handle order): 
        *   Christoph Blecker (@cblecker), Red Hat
        *   Derek Carr (@derekwaynecarr), Red Hat
        *   Nikhita Raghunath (@nikhita), Loodse
        *   Paris Pittman (@parispittman), Google
    *   See [the blog post](https://kubernetes.io/blog/2019/10/03/2019-steering-committee-election-results/) for more information
*   [ 0:00 ]** Release Updates**
    *   Current Release Development Cycle  [Guinevere Saenger - Release Manager]
        *   We’re in Week 2! Shadow selection is 99% complete - congratulations and thanks to all of our hardworking team members
        *   Enhancements Freeze is 15 October!
        *   1.17.0-alpha-1 was released today
        *   Next alpha scheduled for 15 October
    *   Patch Release Updates
        *   1.16.1 released 1 October
        *   Next patch releases scheduled for 15 October
        *   y.x
*   [ 0:00 ] **Contributor Tip of the Week **[First Last] 
    *   A fun graph, contribex info, CI tips, etc.
    *   [Link to a chart, a guide, a tool, etc]
    *   Reach out to #sig-contribex in slack if there is no tip on the agenda yet. Backlog is pinned to the chat. 
*   [ 0:00 ] **SIG Updates**
    *   SIG Leads, check out this set of [recommended topics to cover](https://github.com/kubernetes/community/blob/master/events/community-meeting.md#sig-updates) during your update. [Slide template if you need it](https://docs.google.com/presentation/d/1-nTvKCiqu9UvFYUeM6p6RIqHS5-H-u3_x-V4xj_eIWo/edit#slide=id.g401c104a3c_0_0), please also check the [SIG Update Schedule](https://docs.google.com/spreadsheets/d/1adztrJ05mQ_cjatYSnvyiy85KjuI6-GuXsRsP-T2R3k)!
    *   WG IoT Edge [Dejan Bosanac]
        *   [https://github.com/kubernetes/community/tree/master/wg-iot-edge](https://github.com/kubernetes/community/tree/master/wg-iot-edge)
        *   [https://docs.google.com/presentation/d/1Ozby628UmK91fGOoo_10YDSRJ0BVlUIx-YAF4iwlyOA/](https://docs.google.com/presentation/d/1Ozby628UmK91fGOoo_10YDSRJ0BVlUIx-YAF4iwlyOA/) 
    *   WG Resource Management is spinning down
    *   WG Machine Learning will be spinning down too
*   [ 0:00 ] **📣Announcements 📣**
    *   Contributor Summit Registration is live and we already have 260 RSVPs (~60%). Please do this ASAP. We can always correct, modify, or cancel your registration at a later time. https://events.linuxfoundation.org/events/kubernetes-contributor-summit-north-america-2019/ #contributor-summit for conversation
*   
    *   **👏 **Shoutouts this week (Check in #shoutouts on slack) **👏**
    *   tpepper:
        *   shoutout to @nikhita for a PR description and commit messages in https://github.com/kubernetes/kubernetes/pull/82410 which makes a potentially daunting code review MUCH easier, and to @liggitt for similarly making the cherry-pick review MUCH easier with a stellar PR description text.  Super time saving when there’s a diffstat of “+2,537 −59” but the “why” text focuses the reviewer in on two key lines of code and the associated bugs tracking the problem report.
    *   jdetiber:
        *   Shoutout to @dims for building out the e2e conformance tests using Cluster API and the GCP Provider


## September 26, 2019 - ([recording](https://youtu.be/ghDHHKcQdlo))



*   **Moderators**:  Tim Pepper [VMware/SIG Release]
*   **Note Taker**: Lachlan Evenson [Microsoft/SIG PM]
    *   Subscribe to [this thread](https://discuss.kubernetes.io/t/kubernetes-weekly-community-meeting-notes/35) to get these notes in your inbox
*   [ 0:01 ]**  Demo **-- [Octant](https://github.com/vmware/octant): A web-based, highly extensible platform for developers to better understand the complexity of Kubernetes cluster [Bryan Liles, @bryanl; Wayne Witzel, @wwitzel3]
    *   Web-based, but runs local, using your credentials (simplifies security)
    *   Demo application troubleshooting via the Octant UI
        *   Web app working
        *   Kubectl apply updated app
        *   Web app no longer working
        *   Use Octact to determine the cause
    *   Introduces the concept of “Application” which is a set of consistent labels “app.kubernetes.io/name:httpbin”
    *   Visualization of dependency graph between Kubernetes resources. Detects that the Ingress is pointing to an invalid backend
    *   Drill down into service via the visualization graph and we notice that are no endpoints.
    *   Determine that it’s a bad selector and update and check that the graph is green again.
    *   If you’re on a Mac you can install via `brew install octant`
*   [ 0:14 ]** Release Updates**
    *   1.17 Release Development Cycle  [Guinevere Saenger - Release Manager]
        *   Week 1
        *   Shadow selection happening (application deadline yesterday)
        *   Please be aware that this is a short release
        *   Enhancements freeze 10/15 5pm Pacific
    *   Patch Release Updates
        *   [UPCOMING RELEASE SCHEDULE](https://git.k8s.io/sig-release/releases/patch-releases.md) link
        *   Patch Release   Cherry-picks deadline   Target date 
*   1.16.2  	2019-10-11  		2019-10-15 
*   1.16.1  	2019-09-27  		2019-10-02 
*   1.15.5  	2019-10-11  		2019-10-15 
*   1.14.8  	2019-10-11  		2019-10-15 
*   1.13.12 	2019-10-11  		2019-10-15  (final release of 1.13) 
*   ...as always subject to change for critical-urgent security issues
*   [ 0:17 ] **Contributor Tip of the Week **[Bob Killen] 
    *   What am I an OWNERs of? 
    *   [https://go.k8s.io/owners/mrbobbytables](https://go.k8s.io/owners/mrbobbytables)
        *   Based on [https://cs.k8s.io/](https://cs.k8s.io/), you should use that too! 
*   [ 0:19 ] **SIG Updates**
    *   SIG Leads, check out this set of [recommended topics to cover](https://github.com/kubernetes/community/blob/master/events/community-meeting.md#sig-updates) during your update. [Slide template if you need it](https://docs.google.com/presentation/d/1-nTvKCiqu9UvFYUeM6p6RIqHS5-H-u3_x-V4xj_eIWo/edit#slide=id.g401c104a3c_0_0), please also check the [SIG Update Schedule](https://docs.google.com/spreadsheets/d/1adztrJ05mQ_cjatYSnvyiy85KjuI6-GuXsRsP-T2R3k)!
    *   [ 0:20 ] SIG CLI [Maciej Szulik; confirmed]- [slides](https://docs.google.com/presentation/d/1iEUEC8ZFZgDZYyyjL1kyLBIEEYOIelLz61z3wWCL_Vo/edit)
    *   [ 0:29 ] WG Policy [Zhipeng Huang; confirmed] - [slides](https://docs.google.com/presentation/d/1m53iKvaedB4vPNACB39SWL3XNLDxb-u2J3Iq1rlXMXE/edit?usp=sharing)
    *   [ 0:36 ] WG Component Standard [Michael Taufen; confirmed] - [slides](https://docs.google.com/presentation/d/19jKD8OfahyY8evikouT8m-NXn7Yxt4i2haM9xAjMhyk/edit) 
*   [ 0:43 ] **📣Announcements 📣**
    *   Election Status [Jorge/Bob/Ihor/Brian]
        *   302 of 858 of you have voted.
        *   The next deadline is October 2, that's one week from today! You have until then to complete your ballot. If you have any questions, let us know.
        *   You must be in [voters.md](https://www.google.com/url?q=https%3A%2F%2Fgithub.com%2Fkubernetes%2Fcommunity%2Fblob%2Fmaster%2Fevents%2Felections%2F2019%2Fvoters.md&sa=D&sntz=1&usg=AFQjCNGSDCBbO4BTooEbaj77dC5fAvp4Kw), if you’re not, you cannot vote in this election.
    *   **👏 **Shoutouts this week (Check in #shoutouts on slack) **👏**
        *   **Lots of thanks and reflection as the 1.16 release cycle came to an end.  Great teamwork folks!**
        *   [Lachlan Evenson](https://app.slack.com/team/U0BFDEV1S)
            *   Massive thanks to [@gsaenger](https://kubernetes.slack.com/team/U4H2QU3DW) [@jeefy](https://kubernetes.slack.com/team/U5MCFK468) [@nikopen](https://kubernetes.slack.com/team/U8W5JK630) for being a fantastic lead team for Kubernetes 1.16.
            *   

<p id="gdcalert1" ><span style="color: red; font-weight: bold">>>>>>  gd2md-html alert: inline image link here (to images/Kubernetes-Community0.png). Store image on your image server and adjust path/filename if necessary. </span><br>(<a href="#">Back to top</a>)(<a href="#gdcalert2">Next alert</a>)<br><span style="color: red; font-weight: bold">>>>>> </span></p>


![alt_text](images/Kubernetes-Community0.png "image_tooltip")
6

<p id="gdcalert2" ><span style="color: red; font-weight: bold">>>>>>  gd2md-html alert: inline image link here (to images/Kubernetes-Community1.png). Store image on your image server and adjust path/filename if necessary. </span><br>(<a href="#">Back to top</a>)(<a href="#gdcalert3">Next alert</a>)<br><span style="color: red; font-weight: bold">>>>>> </span></p>


![alt_text](images/Kubernetes-Community1.png "image_tooltip")
9

<p id="gdcalert3" ><span style="color: red; font-weight: bold">>>>>>  gd2md-html alert: inline image link here (to images/Kubernetes-Community2.png). Store image on your image server and adjust path/filename if necessary. </span><br>(<a href="#">Back to top</a>)(<a href="#gdcalert4">Next alert</a>)<br><span style="color: red; font-weight: bold">>>>>> </span></p>


![alt_text](images/Kubernetes-Community2.png "image_tooltip")
3
        *   [Lachlan Evenson](https://app.slack.com/team/U0BFDEV1S)
            *   [@liggitt](https://kubernetes.slack.com/team/U0BGPQ6DS) [@cblecker](https://kubernetes.slack.com/team/U3EDWR9FV) [@dims](https://kubernetes.slack.com/team/U0Y7A2MME) for their tireless work during the 1.16 release getting bugs triaged, following up on PRs, shepherding them through to MERGE, grooming the tide pool, watching the issue queue and just about everything else you can think of. On behalf of the 1.16 release-team, THANK YOU!
            *   

<p id="gdcalert4" ><span style="color: red; font-weight: bold">>>>>>  gd2md-html alert: inline image link here (to images/Kubernetes-Community3.png). Store image on your image server and adjust path/filename if necessary. </span><br>(<a href="#">Back to top</a>)(<a href="#gdcalert5">Next alert</a>)<br><span style="color: red; font-weight: bold">>>>>> </span></p>


![alt_text](images/Kubernetes-Community3.png "image_tooltip")
6

<p id="gdcalert5" ><span style="color: red; font-weight: bold">>>>>>  gd2md-html alert: inline image link here (to images/Kubernetes-Community4.png). Store image on your image server and adjust path/filename if necessary. </span><br>(<a href="#">Back to top</a>)(<a href="#gdcalert6">Next alert</a>)<br><span style="color: red; font-weight: bold">>>>>> </span></p>


![alt_text](images/Kubernetes-Community4.png "image_tooltip")
8

<p id="gdcalert6" ><span style="color: red; font-weight: bold">>>>>>  gd2md-html alert: inline image link here (to images/Kubernetes-Community5.png). Store image on your image server and adjust path/filename if necessary. </span><br>(<a href="#">Back to top</a>)(<a href="#gdcalert7">Next alert</a>)<br><span style="color: red; font-weight: bold">>>>>> </span></p>


![alt_text](images/Kubernetes-Community5.png "image_tooltip")
3
        *   [Stephen Augustus](https://app.slack.com/team/U0E0E78AK)
            *   More shoutouts for SIG Release: [https://twitter.com/stephenaugustus/status/1174797710043430913?s=20](https://twitter.com/stephenaugustus/status/1174797710043430913?s=20)

                    

<p id="gdcalert7" ><span style="color: red; font-weight: bold">>>>>>  gd2md-html alert: inline image link here (to images/Kubernetes-Community6.jpg). Store image on your image server and adjust path/filename if necessary. </span><br>(<a href="#">Back to top</a>)(<a href="#gdcalert8">Next alert</a>)<br><span style="color: red; font-weight: bold">>>>>> </span></p>


![alt_text](images/Kubernetes-Community6.jpg "image_tooltip")
[Stephen Augustus](https://twitter.com/@stephenaugustus) [@stephenaugustus](https://twitter.com/@stephenaugustus)


                    I'm fully aware we're all adults, but I can't help but feel weird dad pride when my "kids" are crushing it.


                    Thank you to all of the SIG Release contributors for the hard work that you do to keep #Kubernetes rolling!


                    Some shoutouts...


                    

<p id="gdcalert8" ><span style="color: red; font-weight: bold">>>>>>  gd2md-html alert: inline image link here (to images/Kubernetes-Community7.png). Store image on your image server and adjust path/filename if necessary. </span><br>(<a href="#">Back to top</a>)(<a href="#gdcalert9">Next alert</a>)<br><span style="color: red; font-weight: bold">>>>>> </span></p>


![alt_text](images/Kubernetes-Community7.png "image_tooltip")
Twitter | [Sep 19th](https://twitter.com/stephenaugustus/status/1174797710043430913?s=20)

            *   

<p id="gdcalert9" ><span style="color: red; font-weight: bold">>>>>>  gd2md-html alert: inline image link here (to images/Kubernetes-Community8.png). Store image on your image server and adjust path/filename if necessary. </span><br>(<a href="#">Back to top</a>)(<a href="#gdcalert10">Next alert</a>)<br><span style="color: red; font-weight: bold">>>>>> </span></p>


![alt_text](images/Kubernetes-Community8.png "image_tooltip")
6

<p id="gdcalert10" ><span style="color: red; font-weight: bold">>>>>>  gd2md-html alert: inline image link here (to images/Kubernetes-Community9.png). Store image on your image server and adjust path/filename if necessary. </span><br>(<a href="#">Back to top</a>)(<a href="#gdcalert11">Next alert</a>)<br><span style="color: red; font-weight: bold">>>>>> </span></p>


![alt_text](images/Kubernetes-Community9.png "image_tooltip")
4
        *   [Jason DeTiberus](https://app.slack.com/team/U0UV07D8T)
            *   Shoutout to [@joonas](https://kubernetes.slack.com/team/U0A2WAE4F) [@alainroy](https://kubernetes.slack.com/team/U0PQCFJ7P) [@vincepri](https://kubernetes.slack.com/team/UCD11GCET) [@tasha](https://kubernetes.slack.com/team/U98L57XTL) [@a_sykim](https://kubernetes.slack.com/team/U1NCJCTFC) [@ncdc](https://kubernetes.slack.com/team/U0A4MJ62V) [@cha](https://kubernetes.slack.com/team/U2R5S3S77) [@Amy Chen@zjs](https://kubernetes.slack.com/team/U20M13D8C) [@justaugustus](https://kubernetes.slack.com/team/U0E0E78AK) [@Vivek Goyal](https://kubernetes.slack.com/team/UFMGKKU82) [@noamran](https://kubernetes.slack.com/team/UNG623S07) [@naadir](https://kubernetes.slack.com/team/U6RDFQAF5) [@timothysc](https://kubernetes.slack.com/team/U09R2P666) [@samba](https://kubernetes.slack.com/team/U7J4ZCS9H) [@dgoel@jieyu](https://kubernetes.slack.com/team/U65PMSC9J) [@tamal](https://kubernetes.slack.com/team/U283HFU8M) [@winnie](https://kubernetes.slack.com/team/UDPR7TZDJ) [@michaelgugino](https://kubernetes.slack.com/team/U6LBB5YN8) [@mrunalp](https://kubernetes.slack.com/team/U0ANSKQ72) [@cecile](https://kubernetes.slack.com/team/U98JPHB2M) [@ritazh](https://kubernetes.slack.com/team/U5CMBA9RD) [@sozercan](https://kubernetes.slack.com/team/U3ML9CKL0) [@jpang@lachie83](https://kubernetes.slack.com/team/U7NM46SN5) [@dlipovetsky](https://kubernetes.slack.com/team/U0LL3DXA9) [@justinsb](https://kubernetes.slack.com/team/U0A6A01FG) [@rudeboy](https://kubernetes.slack.com/team/UE5U1R97H) [@Matt Dennison](https://kubernetes.slack.com/team/U5NB9Q5FW) [@jenny](https://kubernetes.slack.com/team/UJ5AP5665) [@moshloop@hardikdr](https://kubernetes.slack.com/team/U82FU0535) for coming out to the cluster api f2f planning session and helping us define and scope the next release of the project
            *   

<p id="gdcalert11" ><span style="color: red; font-weight: bold">>>>>>  gd2md-html alert: inline image link here (to images/Kubernetes-Community10.png). Store image on your image server and adjust path/filename if necessary. </span><br>(<a href="#">Back to top</a>)(<a href="#gdcalert12">Next alert</a>)<br><span style="color: red; font-weight: bold">>>>>> </span></p>


![alt_text](images/Kubernetes-Community10.png "image_tooltip")
12

<p id="gdcalert12" ><span style="color: red; font-weight: bold">>>>>>  gd2md-html alert: inline image link here (to images/Kubernetes-Community11.gif). Store image on your image server and adjust path/filename if necessary. </span><br>(<a href="#">Back to top</a>)(<a href="#gdcalert13">Next alert</a>)<br><span style="color: red; font-weight: bold">>>>>> </span></p>


![alt_text](images/Kubernetes-Community11.gif "image_tooltip")
9

<p id="gdcalert13" ><span style="color: red; font-weight: bold">>>>>>  gd2md-html alert: inline image link here (to images/Kubernetes-Community12.gif). Store image on your image server and adjust path/filename if necessary. </span><br>(<a href="#">Back to top</a>)(<a href="#gdcalert14">Next alert</a>)<br><span style="color: red; font-weight: bold">>>>>> </span></p>


![alt_text](images/Kubernetes-Community12.gif "image_tooltip")
8

<p id="gdcalert14" ><span style="color: red; font-weight: bold">>>>>>  gd2md-html alert: inline image link here (to images/Kubernetes-Community13.png). Store image on your image server and adjust path/filename if necessary. </span><br>(<a href="#">Back to top</a>)(<a href="#gdcalert15">Next alert</a>)<br><span style="color: red; font-weight: bold">>>>>> </span></p>


![alt_text](images/Kubernetes-Community13.png "image_tooltip")
7

<p id="gdcalert15" ><span style="color: red; font-weight: bold">>>>>>  gd2md-html alert: inline image link here (to images/Kubernetes-Community14.gif). Store image on your image server and adjust path/filename if necessary. </span><br>(<a href="#">Back to top</a>)(<a href="#gdcalert16">Next alert</a>)<br><span style="color: red; font-weight: bold">>>>>> </span></p>


![alt_text](images/Kubernetes-Community14.gif "image_tooltip")
4
        *   [Stephen Augustus](https://app.slack.com/team/U0E0E78AK)
            *   And thanks to you, [@jdetiber](https://kubernetes.slack.com/team/U0UV07D8T), Andy, and TSC for keeping us on task for the week!
            *   

<p id="gdcalert16" ><span style="color: red; font-weight: bold">>>>>>  gd2md-html alert: inline image link here (to images/Kubernetes-Community15.png). Store image on your image server and adjust path/filename if necessary. </span><br>(<a href="#">Back to top</a>)(<a href="#gdcalert17">Next alert</a>)<br><span style="color: red; font-weight: bold">>>>>> </span></p>


![alt_text](images/Kubernetes-Community15.png "image_tooltip")
2

<p id="gdcalert17" ><span style="color: red; font-weight: bold">>>>>>  gd2md-html alert: inline image link here (to images/Kubernetes-Community16.png). Store image on your image server and adjust path/filename if necessary. </span><br>(<a href="#">Back to top</a>)(<a href="#gdcalert18">Next alert</a>)<br><span style="color: red; font-weight: bold">>>>>> </span></p>


![alt_text](images/Kubernetes-Community16.png "image_tooltip")
3

<p id="gdcalert18" ><span style="color: red; font-weight: bold">>>>>>  gd2md-html alert: inline image link here (to images/Kubernetes-Community17.png). Store image on your image server and adjust path/filename if necessary. </span><br>(<a href="#">Back to top</a>)(<a href="#gdcalert19">Next alert</a>)<br><span style="color: red; font-weight: bold">>>>>> </span></p>


![alt_text](images/Kubernetes-Community17.png "image_tooltip")
1
        *   [Josh Berkus](https://app.slack.com/team/U0UKM380M)
            *   huge shoutout to [@thockin](https://kubernetes.slack.com/team/U0AH4GABW) for coming up with tons of ideas for contributor summit sessions
            *   

<p id="gdcalert19" ><span style="color: red; font-weight: bold">>>>>>  gd2md-html alert: inline image link here (to images/Kubernetes-Community18.png). Store image on your image server and adjust path/filename if necessary. </span><br>(<a href="#">Back to top</a>)(<a href="#gdcalert20">Next alert</a>)<br><span style="color: red; font-weight: bold">>>>>> </span></p>


![alt_text](images/Kubernetes-Community18.png "image_tooltip")
4

<p id="gdcalert20" ><span style="color: red; font-weight: bold">>>>>>  gd2md-html alert: inline image link here (to images/Kubernetes-Community19.png). Store image on your image server and adjust path/filename if necessary. </span><br>(<a href="#">Back to top</a>)(<a href="#gdcalert21">Next alert</a>)<br><span style="color: red; font-weight: bold">>>>>> </span></p>


![alt_text](images/Kubernetes-Community19.png "image_tooltip")
4

<p id="gdcalert21" ><span style="color: red; font-weight: bold">>>>>>  gd2md-html alert: inline image link here (to images/Kubernetes-Community20.png). Store image on your image server and adjust path/filename if necessary. </span><br>(<a href="#">Back to top</a>)(<a href="#gdcalert22">Next alert</a>)<br><span style="color: red; font-weight: bold">>>>>> </span></p>


![alt_text](images/Kubernetes-Community20.png "image_tooltip")
3

<p id="gdcalert22" ><span style="color: red; font-weight: bold">>>>>>  gd2md-html alert: inline image link here (to images/Kubernetes-Community21.gif). Store image on your image server and adjust path/filename if necessary. </span><br>(<a href="#">Back to top</a>)(<a href="#gdcalert23">Next alert</a>)<br><span style="color: red; font-weight: bold">>>>>> </span></p>


![alt_text](images/Kubernetes-Community21.gif "image_tooltip")
2

<p id="gdcalert23" ><span style="color: red; font-weight: bold">>>>>>  gd2md-html alert: inline image link here (to images/Kubernetes-Community22.png). Store image on your image server and adjust path/filename if necessary. </span><br>(<a href="#">Back to top</a>)(<a href="#gdcalert24">Next alert</a>)<br><span style="color: red; font-weight: bold">>>>>> </span></p>


![alt_text](images/Kubernetes-Community22.png "image_tooltip")
4
        *   [Tim Hockin](https://app.slack.com/team/U0AH4GABW)
            *   And right back to everyone running that event!
            *   

<p id="gdcalert24" ><span style="color: red; font-weight: bold">>>>>>  gd2md-html alert: inline image link here (to images/Kubernetes-Community23.png). Store image on your image server and adjust path/filename if necessary. </span><br>(<a href="#">Back to top</a>)(<a href="#gdcalert25">Next alert</a>)<br><span style="color: red; font-weight: bold">>>>>> </span></p>


![alt_text](images/Kubernetes-Community23.png "image_tooltip")
3

<p id="gdcalert25" ><span style="color: red; font-weight: bold">>>>>>  gd2md-html alert: inline image link here (to images/Kubernetes-Community24.png). Store image on your image server and adjust path/filename if necessary. </span><br>(<a href="#">Back to top</a>)(<a href="#gdcalert26">Next alert</a>)<br><span style="color: red; font-weight: bold">>>>>> </span></p>


![alt_text](images/Kubernetes-Community24.png "image_tooltip")
4

<p id="gdcalert26" ><span style="color: red; font-weight: bold">>>>>>  gd2md-html alert: inline image link here (to images/Kubernetes-Community25.gif). Store image on your image server and adjust path/filename if necessary. </span><br>(<a href="#">Back to top</a>)(<a href="#gdcalert27">Next alert</a>)<br><span style="color: red; font-weight: bold">>>>>> </span></p>


![alt_text](images/Kubernetes-Community25.gif "image_tooltip")
2
        *   [Vince Prignano](https://app.slack.com/team/UCD11GCET)
            *   Shoutout to [@cecile](https://kubernetes.slack.com/team/U98JPHB2M) for putting a huge amount of work pushing forward CAPZ (Cluster API Azure) to v1alpha2 

<p id="gdcalert27" ><span style="color: red; font-weight: bold">>>>>>  gd2md-html alert: inline image link here (to images/Kubernetes-Community26.png). Store image on your image server and adjust path/filename if necessary. </span><br>(<a href="#">Back to top</a>)(<a href="#gdcalert28">Next alert</a>)<br><span style="color: red; font-weight: bold">>>>>> </span></p>


![alt_text](images/Kubernetes-Community26.png "image_tooltip")
 

<p id="gdcalert28" ><span style="color: red; font-weight: bold">>>>>>  gd2md-html alert: inline image link here (to images/Kubernetes-Community27.png). Store image on your image server and adjust path/filename if necessary. </span><br>(<a href="#">Back to top</a>)(<a href="#gdcalert29">Next alert</a>)<br><span style="color: red; font-weight: bold">>>>>> </span></p>


![alt_text](images/Kubernetes-Community27.png "image_tooltip")
!
            *   

<p id="gdcalert29" ><span style="color: red; font-weight: bold">>>>>>  gd2md-html alert: inline image link here (to images/Kubernetes-Community28.gif). Store image on your image server and adjust path/filename if necessary. </span><br>(<a href="#">Back to top</a>)(<a href="#gdcalert30">Next alert</a>)<br><span style="color: red; font-weight: bold">>>>>> </span></p>


![alt_text](images/Kubernetes-Community28.gif "image_tooltip")
7

<p id="gdcalert30" ><span style="color: red; font-weight: bold">>>>>>  gd2md-html alert: inline image link here (to images/Kubernetes-Community29.png). Store image on your image server and adjust path/filename if necessary. </span><br>(<a href="#">Back to top</a>)(<a href="#gdcalert31">Next alert</a>)<br><span style="color: red; font-weight: bold">>>>>> </span></p>


![alt_text](images/Kubernetes-Community29.png "image_tooltip")
1


## September 12, 2019 - [recording](https://www.youtube.com/watch?v=89pfL0i8BhU)



*   **Moderators**:  Jorge Castro [VMware/SIG ContribEx]
*   [ 0:00 ]** Release Updates**
    *   Current Release Development Cycle  [Jeffrey Sica - Release Manager]
    *   1.16 upcoming milestones
        *   9/10 1.16.0-rc.1
        *   9/10 code thaw
        *   9/12 cherry-pick deadline
        *   9/16 1.16.0 scheduled release date
    *   Patch Release Updates
        *   [UPCOMING RELEASE SCHEDULE](https://git.k8s.io/sig-release/releases/patch-releases.md)
            *   9/13 - cherry pick deadline for 1.13.11, 1.14.7, and 1.15.4
            *   9/18 - release target for 1.13.11, 1.14.7, and 1.15.4
        *   Reminder these pending dates are announced on:
            *   **[https://groups.google.com/forum/#!forum/kubernetes-dev](https://groups.google.com/forum/#!forum/kubernetes-dev)**
*   [ 0:00 ] **SIG Updates**
    *   SIG Leads, check out this set of [recommended topics to cover](https://github.com/kubernetes/community/blob/master/events/community-meeting.md#sig-updates) during your update. [Slide template if you need it](https://docs.google.com/presentation/d/1-nTvKCiqu9UvFYUeM6p6RIqHS5-H-u3_x-V4xj_eIWo/edit#slide=id.g401c104a3c_0_0), please also check the [SIG Update Schedule](https://docs.google.com/spreadsheets/d/1adztrJ05mQ_cjatYSnvyiy85KjuI6-GuXsRsP-T2R3k)!
    *   SIG Scalability [Wojciech Tyczynski] (confirmed)
        *   Slides
    *   SIG Network  [[bowei] (confirmed)](https://docs.google.com/presentation/d/1wKxuuN9wZp0gRRaTA2rJvbu7rPcwuk0JlXvwbxyWBXk/edit#slide=id.g401c104a3c_0_0)
        *   [Slides](https://docs.google.com/presentation/d/1wKxuuN9wZp0gRRaTA2rJvbu7rPcwuk0JlXvwbxyWBXk/edit#slide=id.g401c104a3c_0_0)
    *   SIG Multicluster
        *   Sends regrets, will update at a future date. 
*   [ 0:00 ] **📣Announcements 📣**
    *   Election Update [Jorge/Bob/Ihor/bgrant]
        *   [Candidate platforms are here](https://github.com/kubernetes/community/tree/master/events/elections/2019)
    *   Call for demos for this meeting! See the top of  this document for more information. 
    *   CONTRIBUTOR SUMMIT REGISTRATION IS LIVE! PS - it’s already at 50% capacity. Will be sold out by the end of September. 
        *   [https://events.linuxfoundation.org/events/kubernetes-contributor-summit-north-america-2019/](https://events.linuxfoundation.org/events/kubernetes-contributor-summit-north-america-2019/;)
    *   Release Retrospective is next week with Lachlan Evenson, so not a normal meeting
    *   New users need h[elp on discuss.k8s.io](https://discuss.kubernetes.io/t/help-someone-out-here-win-some-prizes/7877), help out, win a tshirt!
        *   [Please RT this](https://twitter.com/castrojo/status/1172152453300723712)
    *   SIG CLI, WG Policy, and WG Component Standard will be giving a status on 9/26 with your host, Tim Pepper!
    *   **👏 **Shoutouts this week (Check in #shoutouts on slack) **👏**
        *   **@**[vincepr](https://app.slack.com/team/UCD11GCET)i - Huge shoutout and thank you to [@Katharine](https://kubernetes.slack.com/team/UBTBNJ6GL) [@thockin](https://kubernetes.slack.com/team/U0AH4GABW) [@jdetiber](https://kubernetes.slack.com/team/U0UV07D8T) for spending hours together to get Cluster API container build automation up and running
        *   @z[acharysarah](https://app.slack.com/team/U5WQMKJEA) - Shout to [@june.yi](https://kubernetes.slack.com/team/U498UMJ3F) for adding a helper script for localization teams to discover changes to source docs between development branches! [https://github.com/kubernetes/website/pull/15789](https://github.com/kubernetes/website/pull/15789)
        *   @[Matthyx](https://app.slack.com/team/U4A7DM2QZ) - Shout to [@jbeda](https://kubernetes.slack.com/team/U09QZ63DX) [@brendanburns](https://kubernetes.slack.com/team/U0BC5M36Y) [@thockin](https://kubernetes.slack.com/team/U0AH4GABW) and [@aparna](https://kubernetes.slack.com/team/U15SSQ58T) for the nice TechCrunch panel session! [https://youtu.be/OPrh5t24iy4](https://youtu.be/OPrh5t24iy4)
        *   [bentheelder](https://app.slack.com/team/U1P7T516X)

<p id="gdcalert31" ><span style="color: red; font-weight: bold">>>>>>  gd2md-html alert: inline image link here (to images/Kubernetes-Community30.png). Store image on your image server and adjust path/filename if necessary. </span><br>(<a href="#">Back to top</a>)(<a href="#gdcalert32">Next alert</a>)<br><span style="color: red; font-weight: bold">>>>>> </span></p>


![alt_text](images/Kubernetes-Community30.png "image_tooltip")
 - shoutout to [@liggitt](https://kubernetes.slack.com/team/U0BGPQ6DS) for reviewing all of the things 


## September 5, 2019 - ([recording](https://www.youtube.com/watch?v=41Le9wRzbf8))



*   **Moderators**:  Marky [Sysdig/sig-contribex]
*   **Note Taker**: Chris Short [Red Hat/SIG-ContribEx]
    *   Subscribe to [this thread](https://discuss.kubernetes.io/t/kubernetes-weekly-community-meeting-notes/35) to get these notes in your inbox
    *   [Please give us feedback](https://github.com/kubernetes/community/issues/4019) on the usefulness of this meeting to you!
*   [ 0:00 ]**  Demo **-- GSoC Student, Elijah Oyekunle [[@eloyekunle](https://kubernetes.slack.com/team/UBQT0REHG)] - confirmed
        *   [Kubernetes Dashboard to add support for CRDs](https://github.com/eloyekunle/gsoc-2019-meta-k8s)
        *   [CNCF Blog Post](https://www.cncf.io/blog/2019/08/23/cncf-joins-google-summer-of-code-2019-with-17-interns-projects-for-containerd-coredns-kubernetes-opa-prometheus-rook-and-more/)
        *   [Kubernetes](https://github.com/kubernetes/community/blob/master/mentoring/google-summer-of-code.md) blog post 
        *   [Demo slides](https://docs.google.com/presentation/d/1SONr2MgrPhY5P2o4qGG-YH2CsxlzgzSd4XMT49m0dXo/edit#slide=id.gc6fa3c898_0_0)
        *   The k8s dashboard did not support CRDs
            *   Needed to be able to display details about each CRD and its objects.
            *   Support for editing and deletion
            *   Support Scale subresource
            *   Support pinning CRDs to dashboard sidebar
        *   Elijah now works at replex.io
*   [ 0:00 ]** Release Updates**
    *   Current Release Development Cycle  [Lachlan Evenson - Release Manager]
    *   1.16 upcoming milestones
        *   9/9 Docs complete
        *   9/10 1.16.0-rc.1
        *   9/10 code thaw
    *   Patch Release Updates
        *   [UPCOMING RELEASE SCHEDULE](https://git.k8s.io/sig-release/releases/patch-releases.md)
            *   Targeting 1.13.11, 1.14.7, and 1.15.4 for just after 1.16.0
            *   Cherry pick deadline Friday Sept. 13
            *   Release date target Wednesday Sept. 18
    *   Reminder these pending dates are announced on:
        *   **[https://groups.google.com/forum/#!forum/kubernetes-dev](https://groups.google.com/forum/#!forum/kubernetes-dev)**
*   [ 0:00 ] **SIG Updates**
    *   SIG Leads, check out this set of [recommended topics to cover](https://github.com/kubernetes/community/blob/master/events/community-meeting.md#sig-updates) during your update. [Slide template if you need it](https://docs.google.com/presentation/d/1-nTvKCiqu9UvFYUeM6p6RIqHS5-H-u3_x-V4xj_eIWo/edit#slide=id.g401c104a3c_0_0), please also check the [SIG Update Schedule](https://docs.google.com/spreadsheets/d/1adztrJ05mQ_cjatYSnvyiy85KjuI6-GuXsRsP-T2R3k)!
    *   SIG Autoscaling [mwielgus] (confirmed)
        *   Slides
    *   SIG Networking [Bowei]
        *   Moving to next week, we’ll see you there!
    *   SIG Windows  [Michael Michael]  (confirmed)
        *   Slides
*   [ 0:00 ] **📣Announcements 📣**
    *   [Election](https://github.com/kubernetes/community/tree/master/events/elections/2019) Update [Jorge/Ihor/Bob/bgrant]
        *   Three reminders this week: 
        *   You should check to see if you are in voters.md. Your github handle MUST BE IN THIS DOCUMENT to vote: [https://github.com/kubernetes/community/blob/master/events/elections/2019/voters.md](https://github.com/kubernetes/community/blob/master/events/elections/2019/voters.md)
        *   If you feel you have made enough contributions to be able to vote but you are NOT in voters.md, you need to fill out this form by September 11: [https://www.surveymonkey.com/r/k8s-sc-election-2019](https://www.surveymonkey.com/r/k8s-sc-election-2019)
        *   If you are planning on running for Steering Committee please see this section of the documentation, you need to PR your biography into the repo by September 11: [https://github.com/kubernetes/community/tree/master/events/elections/2019#candidacy-process](https://github.com/kubernetes/community/tree/master/events/elections/2019#candidacy-process)
    *   Proposals for sessions at the San Diego Contributor Summit are open until Sept. 9th: [https://forms.gle/BgRAPqLn6W5FHHYN7](https://forms.gle/BgRAPqLn6W5FHHYN7)
    *   **👏 **Shoutouts this week (Check in #shoutouts on slack) **👏**
        *   Shoutout to [@liggitt](https://kubernetes.slack.com/team/U0BGPQ6DS) for tirelessly deflaking our tests! [https://github.com/kubernetes/kubernetes/pull/82200](https://github.com/kubernetes/kubernetes/pull/82200)
        *    shoutout to [@bentheelder](https://kubernetes.slack.com/team/U1P7T516X) and [@Katharine](https://kubernetes.slack.com/team/UBTBNJ6GL) for their work unblocking the queue for merge
        *   Thanks to [@markyjackson](https://kubernetes.slack.com/team/U19TKJ64E) for helping on Jenkins credential issue and sharing his thoughts on Jenkins automation


## August 29, 2019 - ([recording](https://www.youtube.com/watch?v=c1GzVlofSm4))



*   **Moderators**:  Dawn Foster [Pivotal/ContribEx]
*   **Note Taker**: Craig Peters [Microsoft/SIG-x]
    *   Subscribe to [this thread](https://discuss.kubernetes.io/t/kubernetes-weekly-community-meeting-notes/35) to get these notes in your inbox
*   [ 0:00 ]** Release Updates**
    *   Current Release Development Cycle  [Lachlan Evenson - Release Manager]
    *   1.16 Upcoming Milestones
        *   8/29 - 1.16 Code Freeze - label your PRs appropriately! The backlog is big and you don’t want to miss the train
        *   9/3 - Docs PRs ready for review - next Tuesday
        *   9/4 - 1.16.0-beta.2
    *   Patch Release Updates.  
        *   [UPCOMING RELEASE SCHEDULE](https://git.k8s.io/sig-release/releases/patch-releases.md)
        *   TBD for next release on all current supported releases
    *   Reminder these pending dates are announced on:
        *   [https://groups.google.com/forum/#!forum/kubernetes-dev](https://groups.google.com/forum/#!forum/kubernetes-dev)
*   [ 0:00 ]**  Demo **-- Ignite [@luxas] - confirmed
    *   [Slides](https://docs.google.com/presentation/d/1sQHGW93t-LnZZFTTAUyPX80oDfoeTcmADiuZR8bnw90/edit#slide=id.p)
    *   Simplified firecracker UX using the GitOps management model
    *   Questions 
        *   Use of Virtual Kubelet vs CRI (easier development and UX), and 
        *   Difference from kata + kubevirt (full VMs instead of containers)
*   [ 0:00 ] **SIG Updates**
    *   SIG Leads, check out this set of [recommended topics to cover](https://github.com/kubernetes/community/blob/master/events/community-meeting.md#sig-updates) during your update. [Slide template if you need it](https://docs.google.com/presentation/d/1-nTvKCiqu9UvFYUeM6p6RIqHS5-H-u3_x-V4xj_eIWo/edit#slide=id.g401c104a3c_0_0), please also check the [SIG Update Schedule](https://docs.google.com/spreadsheets/d/1adztrJ05mQ_cjatYSnvyiy85KjuI6-GuXsRsP-T2R3k)!
    *   SIG PM [Stephen Augustus]
        *   [Slides](https://docs.google.com/presentation/d/1Pe3Iz4uzMy5gq1alvq7disnFj95MBareK3Hq3M8BHqU/edit?usp=sharing) 
        *   Question: intent is to clean up non- k/k issues
    *   SIG Architecture  [Davanum Srinivas]
        *   Cancelled due to technical difficulties - will be pushed out to a future meeting.
*   [ 0:00 ] **📣Announcements 📣**
    *   Election Update [Jorge]
        *   Three reminders this week: 
        *   You should check to see if you are in voters.md. Your github handle MUST BE IN THIS DOCUMENT to vote: [https://github.com/kubernetes/community/blob/master/events/elections/2019/voters.md](https://github.com/kubernetes/community/blob/master/events/elections/2019/voters.md)
        *   If you feel you have made enough contributions to be able to vote but you are NOT in voters.md, you need to fill out this form by September 11: [https://www.surveymonkey.com/r/k8s-sc-election-2019](https://www.surveymonkey.com/r/k8s-sc-election-2019)
        *   If you are planning on running for Steering Committee please see this section of the documentation, you need to PR your biography into the repo by September 11: [https://github.com/kubernetes/community/tree/master/events/elections/2019#candidacy-process](https://github.com/kubernetes/community/tree/master/events/elections/2019#candidacy-process)
    *   **👏 **Shoutouts this week (Check in #shoutouts on slack) **👏**
        *   From Antonio Ojea - a shoutout to @bentheelder and @Katharine for their work unblocking the queue for merge
        *   To Jorge and the rest putting on the election from Joe Beda
        *   To Release team from Stephen Augustus


## August 22, 2019 - ([recording](https://youtu.be/ECcUlatY64o))



*   **Moderators**:  Josh Berkus [Red Hat/SIG-Release]
*   **Note Taker**: First Last [Company/SIG]
    *   Subscribe to [this thread](https://discuss.kubernetes.io/t/kubernetes-weekly-community-meeting-notes/35) to get these notes in your inbox
*   [ 0:00 ]**  Demo **-- KUDO [Gerred Dillon - [gdillon@mesosphere.com](mailto:gdillon@mesosphere.com)](confirmed)
    *   Slides - [https://docs.google.com/presentation/d/1Ggd4BsOzvJVUngeQ8g3RPPthGAIk9cNlBKewiJKyeuA/edit?usp=sharing](https://docs.google.com/presentation/d/1Ggd4BsOzvJVUngeQ8g3RPPthGAIk9cNlBKewiJKyeuA/edit?usp=sharing)
    *   Website - [https://kudo.dev](https://kudo.dev)
    *   Slack - [https://kubernetes.slack.com/messages/kudo](https://kubernetes.slack.com/messages/kudo)
    *   Github - [https://github.com/kudobuilder/kudo](https://github.com/kudobuilder/kudo)
*   [ 0:00 ]** Release Updates**
    *   Current Release Development Cycle  [Lachlan Evenson - Release Manager]
        *   8/20 - v1.16.0-beta.1
    *   1.16 Upcoming Milestones
        *   8/23 - Docs Placeholder PR deadline
        *   8/29 - 1.16 Code Freeze
    *   Patch Release Updates.  
        *   [UPCOMING RELEASE SCHEDULE](https://git.k8s.io/sig-release/releases/patch-releases.md) 
        *   8/19 - 1.15.3, 1.14.6, 1.13.10 - contain fixes for [CVE-2019-9512](https://nvd.nist.gov/vuln/detail/CVE-2019-9512), [CVE-2019-9514](https://nvd.nist.gov/vuln/detail/CVE-2019-9514) ping and reset floods.
    *   Reminder these pending dates are announced on:
        *   [https://groups.google.com/forum/#!forum/kubernetes-dev](https://groups.google.com/forum/#!forum/kubernetes-dev)
*   [ 0:00 ] **SIG Updates**
    *   SIG Leads, check out this set of [recommended topics to cover](https://github.com/kubernetes/community/blob/master/events/community-meeting.md#sig-updates) during your update. [Slide template if you need it](https://docs.google.com/presentation/d/1-nTvKCiqu9UvFYUeM6p6RIqHS5-H-u3_x-V4xj_eIWo/edit#slide=id.g401c104a3c_0_0), please also check the [SIG Update Schedule](https://docs.google.com/spreadsheets/d/1adztrJ05mQ_cjatYSnvyiy85KjuI6-GuXsRsP-T2R3k)!
    *   SIG Scalability -- postponed
    *   SIG Service Catalog [Mateusz Szostok]
        *   [Slides](https://docs.google.com/presentation/d/13_0OoJyTdfOOu_cPJSRVZUwi0-_DraZspvtDBvhjWP0/edit?usp=sharing) 
        *   Milestone 0.3.0 is moving whole implementation to CRDs
            *   API server version will have last release next week, then go into maintenance
            *   Working on compliance with New Open Service Broker API
            *   Going to move SC resources to Kubernetes domains
            *   Decide what to do with PodPresets
            *   Moving repos:
                *   Minibroker subproject is OSB-API broker implementation (for development & testing)
                *   Go-open-service-broker-client subproject - golang client for communicating with service brokers
            *   Special thanks to Nikhita for migrating repos!
            *   SIG Meeting has changed to 9am Pacific
*   [ 0:00 ] **📣Announcements 📣**
    *   [Steering Committee Election](https://groups.google.com/forum/#!topic/kubernetes-dev/huGXZ2KXBzo)
    *   **👏 **Shoutouts this week (Check in #shoutouts on slack) **👏**
        *   spiffxp: huge thanks to @tpepper for taking notes during #wg-k8s-infra meeting this morning, I can’t keep up with this group without the notes
        *   tpepper: shoutout to @cblecker, @bentheelder and @hhorl for their work across the past five days to get the golang CVE updates accounted for and Kubernetes 1.15.3, 1.14.6, and 1.13.10 out the door today


## August 15, 2019 - ([recording](https://www.youtube.com/watch?v=ZckVULU9sYc))



*   **Moderators**: Chris Short [SIG-ContribEx]
*   **Note Taker**: VOLUNTEER []
*   [ 0:00 ]**  Demo **-- [KubePlus](https://github.com/cloud-ark/kubeplus) - Cluster Add-on to simplify discovery, use and binding of Customer Resources - Devdatta Kulkarni - [devdatta@cloudark.io](mailto:devdatta@cloudark.io) (confirmed)
    *   [https://github.com/cloud-ark/kubeplus](https://github.com/cloud-ark/kubeplus)
    *   [Link to slides](https://drive.google.com/open?id=1fzRLBpCLYBZoMPQhKMQDM4KE5xUh6-xU)
    *   Request for Feedback
        *   [https://github.com/cloud-ark/kubeplus/issues/320](https://github.com/cloud-ark/kubeplus/issues/320)
        *   [https://github.com/cloud-ark/kubeplus/issues/319](https://github.com/cloud-ark/kubeplus/issues/319)
*   [ 0:00 ]** Release Updates**
    *   Current Release Development Cycle  [Lachlan Evenson -1.16 release lead]
    *   8/13 - v1.16.0-beta.0 (1.16 branch cut)
    *   [1.16 Upcoming Milestones](https://git.k8s.io/sig-release/releases/release-1.16/README.md)
        *   Week 8 - Burndown begins!
        *   8/20 - v1.16.0-beta.1
        *   8/29 - 1.16 Code Freeze
    *   Patch Release Updates
        *   [UPCOMING RELEASE SCHEDULE](https://git.k8s.io/sig-release/releases/patch-releases.md) 
            *   target 8/19 - 1.15.3, 1.14.6, 1.13.10
            *   cherry-pick deadline 8/15
            *   Reminder these pending dates are announced on:
                *   [https://groups.google.com/forum/#!forum/kubernetes-dev](https://groups.google.com/forum/#!forum/kubernetes-dev)
*   [ 0:00 ] **Contributor Tip of the Week **[Aaron Crickenberger]
    *   New velodrome dashboards for job health
        *   [http://velodrome.k8s.io/dashboard/db/job-health-release-blocking](http://velodrome.k8s.io/dashboard/db/job-health-release-blocking)
        *   [http://velodrome.k8s.io/dashboard/db/job-health-merge-blocking](http://velodrome.k8s.io/dashboard/db/job-health-merge-blocking)
    *   Data gathered by
        *   [https://github.com/kubernetes/test-infra/tree/master/metrics](https://github.com/kubernetes/test-infra/tree/master/metrics)
    *   Metrics involved
        *   [https://github.com/kubernetes/test-infra/blob/master/metrics/configs/job-health.yaml](https://github.com/kubernetes/test-infra/blob/master/metrics/configs/job-health.yaml)
        *   [https://github.com/kubernetes/test-infra/blob/master/metrics/configs/flakes-daily-config.yaml](https://github.com/kubernetes/test-infra/blob/master/metrics/configs/flakes-daily-config.yaml)
        *   [https://github.com/kubernetes/test-infra/blob/master/metrics/configs/flakes-config.yaml](https://github.com/kubernetes/test-infra/blob/master/metrics/configs/flakes-config.yaml)
        *   [https://github.com/kubernetes/test-infra/blob/master/metrics/configs/failures-config.yaml](https://github.com/kubernetes/test-infra/blob/master/metrics/configs/failures-config.yaml)
*   [ 0:00 ] **SIG Updates**
    *   SIG Leads, check out this set of [recommended topics to cover](https://github.com/kubernetes/community/blob/master/events/community-meeting.md#sig-updates) during your update
    *   [Slide template if you need it](https://docs.google.com/presentation/d/1-nTvKCiqu9UvFYUeM6p6RIqHS5-H-u3_x-V4xj_eIWo/edit#slide=id.g401c104a3c_0_0)
    *   Please also check the [SIG Update Schedule](https://docs.google.com/spreadsheets/d/1adztrJ05mQ_cjatYSnvyiy85KjuI6-GuXsRsP-T2R3k)!
    *   SIG Auth [Mo Khan] (confirmed)
        *   [Link to slides](https://docs.google.com/document/d/1DZrxk-iT6QbxNdLXJN3sTmlfEq1f5FR6SnzSluzxe3A/)
    *   SIG Cluster Lifecycle  [Justin Santa Barbara] (confirmed)
        *   [Link to slides](https://docs.google.com/presentation/d/1gcuz26EEWYVef4gStz1ghLhuTNrTYxfJPxoBAPS-1p8/edit?usp=sharing)
*   [ 0:00 ] **📣Announcements 📣**
    *   We are [now soliciting session proposals](https://forms.gle/L9a1FMrXspVuJCUs7) for the 2019 San Diego Contributor Summit.  If you plan to be at the CS, please consider leading -- or at least requesting -- a session!
    *   SIG Intro and Deep Dive session proposals are due August 16th.  Your SIG should decide on these and submit them soon.
    *   CNCF has [published results of the Kubernetes Security Audit](https://www.cncf.io/blog/2019/08/06/open-sourcing-the-kubernetes-security-audit/) including a set of findings documents:
        *   [Kubernetes Security Review](https://github.com/kubernetes/community/blob/master/wg-security-audit/findings/Kubernetes%20Final%20Report.pdf)
        *   [Attacking and Defending Kubernetes Installations](https://github.com/kubernetes/community/blob/master/wg-security-audit/findings/AtredisPartners_Attacking_Kubernetes-v1.0.pdf)
        *   [Whitepaper](https://github.com/kubernetes/community/blob/master/wg-security-audit/findings/Kubernetes%20White%20Paper.pdf)
        *   [Threat Model](https://github.com/kubernetes/community/blob/master/wg-security-audit/findings/Kubernetes%20Threat%20Model.pdf)
*   **👏 **Shoutouts this week (Check in #shoutouts on slack) **👏**
    *   spiffxp — shouts to @a_sykim for putting in the work reviewing kubernetes/kubernetes test PRs (promoted to approver https://github.com/kubernetes/kubernetes/pull/81176)
    *   spiffxp — shouts to @mrbobbytables for the many kubernetes/community PRs authored and reviewed (promoted to approver https://github.com/kubernetes/community/pull/3986) (edited)
    *   chrisshort — Shoutout to @castrojo for reminding me at the right time to do the right thing for the community meeting next week: https://kubernetes.slack.com/archives/C1TU9EB9S/p1565362165442700
    *   ehashman — shoutout to @gsaenger -- I'm helping some folks on my team get started with upstream contributions and when I hunt through the docs and associated issues I keep running into all sorts of awesome improvements she led :sparkles:


## August 8, 2019 - ([recording](https://youtu.be/EyW2m5Xa_BQ))



*   **Moderators**:  Josh Berkus [SIG-Release]
*   **Note Taker**: Lachlan Evenson [SIG-Release, SIG-PM]
*   [ 0:00 ]**  Demo **-- Pulumi: Managed Kubernetes Clusters as Code [Levi, levi@pulumi.com] (confirmed)
    *   [Link to slides](https://drive.google.com/file/d/15Aw9B3gL9SJFXquFwfVfIvD2aoU4KMiF/view?usp=sharing)
    *   [https://github.com/pulumi/examples/tree/master/kubernetes-ts-multicloud](https://github.com/pulumi/examples/tree/master/kubernetes-ts-multicloud)
*   [ 0:00 ]** Release Updates**
    *   Current Release Development Cycle  [Jeffrey Sica -1.16 Lead shadow]
    *   8/7 - v1.16.0-alpha.3 (final alpha release)
    *   [1.16 Upcoming Milestones](https://git.k8s.io/sig-release/releases/release-1.16/README.md)
        *   8/13 - v1.16.0-beta.0 (1.16 branch cut)
        *   8/29 - 1.16 Code Freeze
    *   Patch Release Updates
        *   Aug. 5, 2019 - 1.15.2, 1.14.5, 1.13.9
            *   [CVE-2019-11249](https://github.com/kubernetes/kubernetes/pull/80436): Incomplete fixes for CVE-2019-1002101 and CVE-2019-11246, kubectl cp potential directory traversal
            *   [CVE-2019-11247](https://github.com/kubernetes/kubernetes/pull/80750): API server allows access to custom resources via wrong scope
            *   Reminder these are announced on:
                *   [https://groups.google.com/forum/#!forum/kubernetes-security-announce](https://groups.google.com/forum/#!forum/kubernetes-security-announce)
                *   [https://groups.google.com/forum/#!forum/kubernetes-announce](https://groups.google.com/forum/#!forum/kubernetes-announce)
        *   [UPCOMING RELEASE SCHEDULE](https://git.k8s.io/sig-release/releases/patch-releases.md) 
            *   target 8/19 - 1.15.3, 1.14.6, 1.13.10
            *   cherry-pick deadline 8/15
            *   Reminder these pending dates are announced on:
                *   [https://groups.google.com/forum/#!forum/kubernetes-dev](https://groups.google.com/forum/#!forum/kubernetes-dev)
*   [ 0:00 ] **SIG Updates**
    *   SIG Leads, check out this set of [recommended topics to cover](https://github.com/kubernetes/community/blob/master/events/community-meeting.md#sig-updates) during your update
    *   [Slide template if you need it](https://docs.google.com/presentation/d/1-nTvKCiqu9UvFYUeM6p6RIqHS5-H-u3_x-V4xj_eIWo/edit#slide=id.g401c104a3c_0_0)
    *   Please also check the [SIG Update Schedule](https://docs.google.com/spreadsheets/d/1adztrJ05mQ_cjatYSnvyiy85KjuI6-GuXsRsP-T2R3k)!
    *   SIG Contributor Experience [Paris Pittman] (confirmed)
        *   [Deck](https://docs.google.com/presentation/d/1_bwA827BVykr8n32B9QoxcBxaFhFUCb0acBLebuzQK0/edit?usp=sharing) - shared with contribex and k-dev mailing lists
        *   New subproject meetings
        *   SIG meeting has been moved to bi-weekly
        *   Mentoring subproject had first meeting this morning (8/8)
    *   SIG Scheduling  [Bobby Salamat](confirmed)
        *   [Link to Preso/materials](https://docs.google.com/presentation/d/1rt2grGgJg96yLrHhM7OOo1U89T0sS-IbiKDhIkioVKk/)
*   [ 0:00 ] **📣Announcements 📣**
    *   We are [now soliciting session proposals](https://forms.gle/L9a1FMrXspVuJCUs7) for the 2019 San Diego Contributor Summit.  If you plan to be at the CS, please consider leading -- or at least requesting -- a session!
    *   SIG Intro and Deep Dive session proposals are due August 16th.  Your SIG should decide on these and submit them soon.
    *   CNCF has [published results of the Kubernetes Security Audit](https://www.cncf.io/blog/2019/08/06/open-sourcing-the-kubernetes-security-audit/) including a set of findings documents:
        *   [Kubernetes Security Review](https://github.com/kubernetes/community/blob/master/wg-security-audit/findings/Kubernetes%20Final%20Report.pdf)
        *   [Attacking and Defending Kubernetes Installations](https://github.com/kubernetes/community/blob/master/wg-security-audit/findings/AtredisPartners_Attacking_Kubernetes-v1.0.pdf)
        *   [Whitepaper](https://github.com/kubernetes/community/blob/master/wg-security-audit/findings/Kubernetes%20White%20Paper.pdf)
        *   [Threat Model](https://github.com/kubernetes/community/blob/master/wg-security-audit/findings/Kubernetes%20Threat%20Model.pdf)
*   **👏 **Shoutouts this week (Check in #shoutouts on slack) **👏**
    *   @j[detiber](https://app.slack.com/team/U0UV07D8T) - Shoutout to [@justinsb](https://kubernetes.slack.com/team/U0A6A01FG) and [@amerai](https://kubernetes.slack.com/team/U3TRY5WV7) for getting AWS Credentials created and updated for the cluster-api-provider-aws testing using Boskos
    *   @parispittman - thanks to our august [#meet-our-contributors](https://kubernetes.slack.com/archives/C8WRR2BB9) mentors! elana ([@ehashman](https://kubernetes.slack.com/team/U9X5ARSLS)) and nikhita ([@nikhita](https://kubernetes.slack.com/team/U2PQHGMLN)). they gave great advice on KEPs, getting started, how to communicate within the project, and so much more. and thanks to [@jeefy](https://kubernetes.slack.com/team/U5MCFK468) for making the youtube magic happen. (edited)
    *   @[jimangel](https://app.slack.com/team/U4HSVFA5U) Shoutout to [@dims](https://kubernetes.slack.com/team/U0Y7A2MME) & [@codyc](https://kubernetes.slack.com/team/UG7C9377V) for assisting new contributors navigate git issues in [#sig-docshttps://kubernetes.slack.com/archives/C1J0BPD2M/p1565222189149200](https://kubernetes.slack.com/archives/C1J0BPD2M) 


## August 1, 2019 - ([recording](https://youtu.be/HDf-yIQQdRY))



*   **Moderators**:  Jorge Castro [SIG Contributor Experience]
*   **Note Taker**: This could be you!  [Company/SIG]
*   [ 0:00 ]**  Demo **--  [Garden](https://garden.io/) - Development orchestrator for Kubernetes - Jon Edvald
    *   [https://github.com/garden-io/garden](https://github.com/garden-io/garden)
    *   #garden on the Kubernetes Slack
    *   A blog post on our remote Kubernetes dev features: [https://medium.com/garden-io/you-dont-need-kubernetes-on-your-laptop-37653cbb28c9](https://medium.com/garden-io/you-dont-need-kubernetes-on-your-laptop-37653cbb28c9)
*   [ 0:00 ]** Release Updates**
    *   Current Release Development Cycle  [Lachlan Evenson - Release Manager]
        *   This week in 1.16
            *   7/30 - Enhancements Freeze
                *   Tracking 39 enhancements
                    *   Alpha: 17
                    *   Beta: 12
                    *   Stable: 10
                    *   Enhancement spotlight - [Sidecar containers (ALPHA)](https://github.com/kubernetes/enhancements/issues/753)
            *   7/30 - v1.16.0-alpha.2
        *   Upcoming Milestones
            *   8/6 - v1.16.0-alpha.3
    *   Patch Release Updates
        *   1.13.9, 1.14.5, 1.15.2 all TBD (Mid-August..)
*   [ 0:00 ] **SIG Updates**
    *   SIG Leads, check out this set of [recommended topics to cover](https://github.com/kubernetes/community/blob/master/events/community-meeting.md#sig-updates) during your update
    *   [Slide template if you need it](https://docs.google.com/presentation/d/1-nTvKCiqu9UvFYUeM6p6RIqHS5-H-u3_x-V4xj_eIWo/edit#slide=id.g401c104a3c_0_0)
    *   Please also check the [SIG Update Schedule](https://docs.google.com/spreadsheets/d/1adztrJ05mQ_cjatYSnvyiy85KjuI6-GuXsRsP-T2R3k)!
    *   Product Security Committee [Jonathan Pulsifer]
        *   [https://docs.google.com/presentation/d/1ym-eIi0ZxIyuP3WccVbmgjIDCFimfVV8Riimb04VWvs/edit#](https://docs.google.com/presentation/d/1ym-eIi0ZxIyuP3WccVbmgjIDCFimfVV8Riimb04VWvs/edit#) 
    *   SIG Instrumentation [Piotr Szczesniak] (confirmed) 
        *   [Slides](https://docs.google.com/presentation/d/1F3gmnDQ1uetEWRu5ex1S_phDcbELrtu3QqlySgWDGcs/edit)
    *   SIG Docs [Jim Angel] (confirmed)
        *   [https://docs.google.com/presentation/d/1OWOE7mBAyAr7EYm90B1VMtCnC6xiPf64N8H5PKk0OMM/edit?usp=sharing](https://docs.google.com/presentation/d/1OWOE7mBAyAr7EYm90B1VMtCnC6xiPf64N8H5PKk0OMM/edit?usp=sharing)
    *   SIG Storage [Saad Ali]
        *   [Slides](https://docs.google.com/presentation/d/1yQI-j8tDK5zWdW3u8Kp3RPL0n0vXLQJjtNDpGIpHyTg/edit?usp=sharing)
*   [ 0:00 ] **📣Announcements 📣**
    *   Conferences: 
        *   Cloud Native Rejekts is happening for San Diego: [https://cloud-native.rejekts.io/](https://cloud-native.rejekts.io/)
        *   [https://events.linuxfoundation.org/events/kubernetes-summit-sydney-2019/](https://events.linuxfoundation.org/events/kubernetes-summit-sydney-2019/)
        *   [https://events.linuxfoundation.org/events/kubernetes-summit-seoul-2019/](https://events.linuxfoundation.org/events/kubernetes-summit-seoul-2019/)
    *   SAVE THE DATE! **Contributor** Summit (not to be confused with Kubernetes Summit!) San Diego will have something for everyone! November 17-18th. Registration will go live at the beginning of September. Updates will come through here and kubernetes-dev@googlegroups.com.
    *   Next week’s updates, SIG Cloud Provider, Contributor Experience, and Scheduling
    *   **👏 **Shoutouts this week (Check in #shoutouts on slack) **👏**
        *   [nikhita](https://app.slack.com/team/U2PQHGMLN) would like to shoutout to Alison Dowdney [@alisondy](https://kubernetes.slack.com/team/U9CBCBLCV) for running (her first) SIG Contribex APAC meeting sooo smoothly today! 

<p id="gdcalert32" ><span style="color: red; font-weight: bold">>>>>>  gd2md-html alert: inline image link here (to images/Kubernetes-Community31.png). Store image on your image server and adjust path/filename if necessary. </span><br>(<a href="#">Back to top</a>)(<a href="#gdcalert33">Next alert</a>)<br><span style="color: red; font-weight: bold">>>>>> </span></p>


![alt_text](images/Kubernetes-Community31.png "image_tooltip")



## July 23, 2019 - ([recording](https://youtu.be/6uZScaWEb08))



*   **Moderators**:  Jorge Castro  [SIG Contributor Experience]
*   **Note Taker**: Josh Berkus  [RH/Release]
*   [ 0:00 ]**  Demo **-- [Conftest](https://github.com/instrumenta/conftest) - (7/25) using Open Policy Agent to write unit tests for Kubernetes configs - [[gareth@morethanseven.net](mailto:gareth@morethanseven.net)] (confirmed)
*   
    *   [Link to slides](https://speakerdeck.com/garethr/unit-testing-kubernetes-configs-using-open-policy-agent-and-conftest)
    *   [https://github.com/instrumenta/conftest](https://github.com/instrumenta/conftest)
    *   Lots of us have written bad kubernetes configs -- it would be good to validate them before deployment.
    *   Write policies for Open Policy Agent using Rego, OPA's DSL
    *   Then point it at a config file and it will unit test it
    *   Can also validate arbitrary JSON docs (YAML, etc.)
    *   Did several demos, including validating a MySQL Helm chart
    *   #conftest channel on slack.openpolicyagent.org
*   [ 0:00 ]** Release Updates**
    *   Current Release Development Cycle  [Guinevere Saenger - Release Manager Shadow]
        *   Enhancements Freeze Tuesday July 30th.  
        *   We will also release Alpha 2 that day
    *   Patch Release Updates
        *   none this week
*   [ 0:00 ] **SIG Updates**
    *   SIG Leads, check out this set of [recommended topics to cover](https://github.com/kubernetes/community/blob/master/events/community-meeting.md#sig-updates) during your update
    *   [Slide template if you need it](https://docs.google.com/presentation/d/1-nTvKCiqu9UvFYUeM6p6RIqHS5-H-u3_x-V4xj_eIWo/edit#slide=id.g401c104a3c_0_0)
    *   Please also check the [SIG Update Schedule](https://docs.google.com/spreadsheets/d/1adztrJ05mQ_cjatYSnvyiy85KjuI6-GuXsRsP-T2R3k)!
    *   SIG Testing  [Aaron Crickenberger]  (confirmed) 
        *   [Slides](https://docs.google.com/presentation/d/1Uha4uxmglm6g7e25DpIn_nVsobJHkSYqIRyu8GqjRVY/edit#slide=id.g5df97ebb00_3_5)
        *   SIG-testing creates infrastructure, they dont' write the tests.  
        *   Subprojects:
            *   KIND (kubernetes-in-docker)
                *   Now has support for IPv6
                *   Only deployment of Kube currently passing* Conformance
                *   Much faster than it used to be
                *   Looking ahead to "road to 1.0" 
                *   Going to remain focused on core feature set
                *   Need contributors!
                *   Would like to support more runtimes
                *   Also want to support more E2E tests
            *   Prow (github automation)
                *   New plugin: Nikita added auto-milestone-add for PRs (would be nice to backfill for this, anyone want to write it?)
                *   Spyglass shows the Prow job results, you can now link to specific log lines for failed jobs.
                *   Prow now works with Bugzilla and Gerrit
                *   Beta support for Tekton pipelines (as well as existing support for Podspecs and Build CRDs)
                *   is now an active project that is distinct from Kubernetes, needs a roadmap (help wanted)
                *   Several KEPs in progress
                *   We also need unit testing for Prow (help wanted)
            *   Test-Infra
                *   go test bench creates junit test results
                *   working on better local testing of Prow jobs
                *   trying to break up Testgrid config file instead of having One File To Rule Them All so that folks can make their own changes
                *   need to measure unit test coverage
                *   triage tool needs rewriting in go (help wanted)
                *   existing python tooling needs to be upgraded to Python3 (help wanted)
            *   Testing-Commons (making repeatable testing frameworks)
                *   Trying to shrink the body of 40+ Kube test images down to just 1-2
                *   Move E2E tests out of tree, maybe migrate to new framework instead of ginko
            *   Workgroup: wg-k8s-infra
                *   Takes all of SIG-testing stuff and implements it on Google Cloud so that we can actually run testing
            *   We are also open sourcing TestGrid!  (help wanted)
            *   SIG is re-thinking meeting schedule, to accommodate other time zones
            *   Have lots of Good First Issues for you to help with
    *   [ 0:00 ] **📣Announcements 📣**
        *   Don’t forget about the [API deprecations](https://kubernetes.io/blog/2019/07/18/api-deprecations-in-1-16/)!
        *   Protip - book your Kubecon travel if you’re planning to attend. :D
        *   SIG instrumentation, SIG Storage, SIG Docs, and the Product Security Committee will be giving their updates next week. 
        *   Want to help host this meeting? Ping @castrojo, we’re always looking for new people to help run this meeting!
    *   **👏 **Shoutouts this week (Check in #shoutouts on slack) **👏**
*   **[Bentheelder](https://app.slack.com/team/U1P7T516X)** (Benjamin Elder) - Shoutout again to [@aojea](https://kubernetes.slack.com/team/U7CK9A960) (Antonio Ojea), thanks to his work we finally have CI passing all conformance tests with an IPv6 [#kind](https://kubernetes.slack.com/archives/CEKK1KTN2) cluster!
    *   [https://testgrid.k8s.io/conformance-kind#Summary&width=20](https://testgrid.k8s.io/conformance-kind#Summary&width=20)
    *   [https://testgrid.k8s.io/conformance-kind#kind%20(IPv6),%20master%20(dev)&width=20](https://testgrid.k8s.io/conformance-kind#kind%20(IPv6),%20master%20(dev)&width=20)
*   **[June.yi](https://app.slack.com/team/U498UMJ3F)** (June Yi) Shoutout to [@seungkyua](https://kubernetes.slack.com/team/U2LTMGVJ7) (Seungkuu Ahn), [@ianychoi](https://kubernetes.slack.com/team/U2M2X370B) (Ian Choi), [@Jesang](https://kubernetes.slack.com/team/UC323SDK3) (Jesang Myung) and [@Seokho](https://kubernetes.slack.com/team/UDBR3JL67) (Seokho Son) for encouraging docs localization as an event host, a session speaker or an attendee at the local community event, Open Infrastructure & Cloud Native Days Korea 2019.
*   **[Detiber](https://app.slack.com/team/U0UV07D8T)** (Jason Detiberus): Shoutout to [@thockin](https://kubernetes.slack.com/team/U0AH4GABW) (Tim Hockin) for helping with troubleshooting and fixing a head scratching permissions issue related to the image promotion process


## July 18, 2019 - ([recording](https://www.youtube.com/watch?v=RWbNg4Wjwpg))



*   **Moderators**:  Jeffrey Sica [SIG UI/ContribEx]
*   **Note Taker**: Bob Killen / Chris Short - Contribex 
*   [ 0:00 ]**  Demo **-- Cluster API Docker Provider - Chuck Ha ([chuckh@vmware.com](mailto:chuckh@vmware.com))
    *   GitHub: [https://github.com/kubernetes-sigs/cluster-api-provider-docker](https://github.com/kubernetes-sigs/cluster-api-provider-docker)
    *   Cluster API has been built extensible enough to be able to provide a generic interface for multiple providers.
    *   A bootstrap or management cluster is required to host the CRDs and configs for the desired clusters.
    *   Docker provider backend was built for fast local testing.
    *   cli-tool - capdctl
        *   Uses [KinD](https://github.com/kubernetes-sigs/kind) as a backend.
        *   Strips out some cloud service bits that aren't necessary from clusterctl
    *   Only requirement is an “external” load balancer.
    *   Can modify clusters after initial provisioning.
    *   Provisioned clusters pass standard conformance tests.
*   [ 0:00 ]** Release Updates**
    *   Current Release Development Cycle  [Jeffrey Sica - Release Manager Shadow]
        *   7/16 - 1.16.0-alpha.1
        *   7/30 - Enhancements freeze
    *   Patch Release Updates
        *   7/18 - 1.15.1
*   [ 0:00 ] **Contributor Tip of the Week **[???] 
    *   ???
*   [ 0:00 ] **KEP of the Week** [Kubernetes Enhancement Proposals]	
    *   [Link to KEP or PR] - [Status] - where to follow up discussion
*   [ 0:00 ] **SIG Updates**
    *   SIG Leads, check out this set of [recommended topics to cover](https://github.com/kubernetes/community/blob/master/events/community-meeting.md#sig-updates) during your update
    *   [Slide template if you need it](https://docs.google.com/presentation/d/1-nTvKCiqu9UvFYUeM6p6RIqHS5-H-u3_x-V4xj_eIWo/edit#slide=id.g401c104a3c_0_0)
    *   Please also check the [SIG Update Schedule](https://docs.google.com/spreadsheets/d/1adztrJ05mQ_cjatYSnvyiy85KjuI6-GuXsRsP-T2R3k)!
    *   **SIG Azure** [ Stephen Augustus ] (Confirmed)
        *   [Slides](https://docs.google.com/presentation/d/16SxXEYQBXKCeCj4E2ODzRkOvZfKLESe-5jV9UxSRhoE/edit?usp=sharing)
        *   Aiming to move out of tree by the 1.18 release.
        *   Last Cycle
            *   Improving testing for out-of-tree cluster providers
            *   New SIG Azure Chair - [Craig Peters](https://github.com/craiglpeters)
        *   Upcoming Cycle (1.16):
            *   continue working on moving azure cloud provider out of tree
            *   move Azure availability zones to GA
            *   move Azure Cross-resource group nodes to GA
            *   Cluster API Azure
                *   VMSS integration
                *   Better AZ Support
                *   Work on v1alpha2 implementation
            *   Complete Administrative work related to SIG Cloud Provider consolidation.
        *   Looking for help/contributors for out-of-tree Azure provider 
    *   **SIG Release** [ Tim Pepper ] (Confirmed) [slides](https://docs.google.com/presentation/d/1t-bOgt6IfHW-TrdMfE3oopleSEqRgvECybZTq7GqVHU/edit?usp=sharing)
        *   Last Cycle
            *   Improved Shadow process
            *   Made improvements to documentation and automation
            *   “test-infra” role has been automated completely
            *   [New release notes website](https://relnotes.k8s.io).
            *   Last scalability issues in 1.15 release almost derailed release, but only caused a slight delay.
                *   
            *   Patch release team has been grown and documentation improved
            *   Release Engineering subproject has been kicked off.
        *   Upcoming Cycle (1.16)
            *   Release Engineering Subproject along with the WG-K8s-infra group
                *   Audit and improve publishing of artifacts
                    *   cleaning up packages (RPMs/Debs)
                *   [Gathering information / Brainstorming](https://docs.google.com/document/d/1Js_36K51Q6AjEsVUjRBMISTA4D7cnjZmoSkn43_Jmxo/edit) on current processes before coming together as a “OMNI” KEP
            *   Release Team
                *   refine release blocking criteria
                *   improve testgrid blocking/informing dashboards
                *   branch management role shifting to “release manager” team
                *   work closer with sig-scalability
        *   Things needed from community
            *   Ongoing attention to CI Signal
                *   deflake tests
                *   make sure tests are owned and get notified of failures.
                *   Keep tests green.
        *   Licensing subproject:
            *   looking for more contributors (reach out ot nikhita)
        *   Release Team:
            *   Big shoutout to Josh Berkus as emeritus lead and keeping things going.
        *   Release Managers / Release Engineering subproject:
            *   [Release Managers doc](https://github.com/kubernetes/sig-release/blob/master/release-managers.md)
            *   Driving down tech debt in release process
            *   Composed of members from:
                *   Patch Release Team
                *   Branch Managers
                *   Release Manager Associates
                *   Build Admins
                *   SIG Chairs
                *   PSC
            *   Building contributor ladder for Release Manager group
            *   Need help moving from :bashfire: to go
                *   background info: [https://github.com/kubernetes/release/issues/816](https://github.com/kubernetes/release/issues/816)
        *   Related Working Group Status
            *   WG LTS
                *   Improve conformance
                *   Move more APIs to stable
*   [ 0:00 ] **📣Announcements 📣**
    *   ???
    *   **👏 **Shoutouts this week (Check in #shoutouts on slack) **👏**
        *   


## July 11, 2019 - ([recording](https://www.youtube.com/watch?v=B6uvfOiFE30))



*   **Moderators**:  Jonas Rosland [SIG ContribEx]
*   **Note Taker**: First Last [Company/SIG]
*   [ 0:00 ]**  Demo **-- [Volcano ](http://github.com/volcano-sh/volcano)- A Kubernetes native batch system [Klaus Ma, [klaus1982.cn@gmail.com](mailto:klaus1982.cn@gmail.com)] (confirmed)
    *   [Link to slides](https://docs.google.com/presentation/d/1AadtRuC-abikWVvz01PW6RdP-GQz1Vz17X9IxqR-l7s/edit?usp=sharing)
    *   [https://github.com/volcano-sh/volcano](https://github.com/volcano-sh/volcano)
*   [ 0:00 ]** Release Updates**
    *   Current Release Development Cycle  [Lachlan Evenson - Release Manager]
        *   Week 2 of a 12 week release cycle (9/16 is the release date)
        *   [1.16 Release team](https://github.com/kubernetes/sig-release/blob/master/releases/release-1.16/release_team.md) has been assembled 
        *   7/16 - 1.16.0-alpha.1
        *   7/30 - Enhancements freeze
    *   Patch Release Updates
        *   7/18 - 1.15.1
        *   7/8 - 1.12.10, 1.13.8, 1.14.4 released
        *   1.12.10 - was the last patch release of 1.12
*   [ 0:00 ] **Contributor Tip of the Week **[Aaron Crickenberger] 
    *   [https://github.com/kubernetes/test-infra](https://github.com/kubernetes/test-infra) - updated docs
    *   This One Weird Trick to make your testgrid changes merge faster
    *   [https://github.com/kubernetes/test-infra/tree/master/config/jobs#job-examples](https://github.com/kubernetes/test-infra/tree/master/config/jobs#job-examples)
    *   [https://github.com/kubernetes/test-infra/blob/master/testgrid/config.md#prow-job-configuration](https://github.com/kubernetes/test-infra/blob/master/testgrid/config.md#prow-job-configuration)
*   [ 0:00 ] **KEP of the Week** [Kubernetes Enhancement Proposals]	
    *   [Link to KEP or PR] - [Status] - where to follow up discussion
*   [ 0:00 ] **SIG Updates**
    *   SIG Leads, check out this set of [recommended topics to cover](https://github.com/kubernetes/community/blob/master/events/community-meeting.md#sig-updates) during your update
    *   [Slide template if you need it](https://docs.google.com/presentation/d/1-nTvKCiqu9UvFYUeM6p6RIqHS5-H-u3_x-V4xj_eIWo/edit#slide=id.g401c104a3c_0_0)
    *   Please also check the [SIG Update Schedule](https://docs.google.com/spreadsheets/d/1adztrJ05mQ_cjatYSnvyiy85KjuI6-GuXsRsP-T2R3k)!
    *   SIG Cloud Provider [Andrew Sy Kim] (confirmed)
        *   [Slides](https://docs.google.com/presentation/d/18FlGrOM1VIPZw2fYy1_t4Z6zGXBeOyWwJvH1dEHWOjs/edit?usp=sharing)
    *   SIG API Machinery  [David Eads] (confirmed)
        *   [Slides](https://docs.google.com/presentation/d/18IPl6uLmd2WWg11VDXeJDl1u5qTMu21cNZlvrtjyE5s/edit?usp=sharing)
    *   Steering [Aaron Crickenberger] (confirmed)
        *   [Slides](https://docs.google.com/presentation/d/1GWRHyKiNeWZYItvNFz6Vm9rOwqQghVa_m5djlgsTxMc/edit#slide=id.g401c104a3c_0_0)
*   [ 0:00 ] **📣Announcements 📣**
    *   More testgrid stuff I forgot if we have time (prefer shoutouts if we don’t) [Aaron Crickenberger]
        *   [We moved a lot of existing job configs over to annotations for you](https://github.com/kubernetes/test-infra/pull/13205)
        *   [Testgrid changes (annotations or not) must pass tests](https://github.com/kubernetes/test-infra/blob/31b2a35/testgrid/cmd/configurator/main_test.go)

        **👏 **Shoutouts this week (Check in #shoutouts on slack) **👏**

    *   @bentheelder (Benjamin Elder): Shoutout again to @aojea (Antonio Ojea), thanks to his work we finally have CI passing all conformance tests with an IPv6 #kind cluster!
    *   
    *   
    *   


## July 4th, 2019 - (Cancelled!)



*   **~~Moderators:  Ihor Dvoretskyi (CNCF)~~**
*   **~~Note Taker: ~~**
*   ~~[ 0:00 ]**  Demo**~~
    *   ~~No demo this week~~
*   ~~[ 0:00 ]** Release Updates**.~~
*   ~~[ 0:00 ] **Contributor Tip of the Week **~~
*   ~~[ 0:00 ] **SIG Updates**~~
    *   **~~SIG CLI~~**
        *   **~~To be rescheduled~~**
    *   **~~SIG Node~~**
        *   **~~To be rescheduled~~**
*   ~~[ 0:00 ] **📣Announcements **~~


## June 20, 2019 - Release Retrospective for 1.15 ([recording](https://www.youtube.com/watch?v=UmJikYIpKLY&feature=youtu.be))



*   **Moderators:** Christine Pukropski (@christine)
*   **See Retrospective Doc:** [https://bit.ly/115-retro](https://bit.ly/115-retro) 
*   Remainder of retro will happen in July 2, 2019 SIG Release meeting: [https://git.k8s.io/community/sig-release](https://git.k8s.io/community/sig-release) 
*   Normal Community Meeting next week!


## June 13, 2019 - ([recording](https://www.youtube.com/watch?v=UQ4yFCeCdRs))



*   **Moderators**:  Lachlan Evenson [sig-pm]
*   **Note Taker**: Jorge Castro [VMware/SIG Contributor Experience]
*   [ 0:00 ]**  Demo **-- [Kyverno](https://kyverno.io/) [[jim@nirmata.com](mailto:jim@nirmata.com)] - Kubernetes native policy management  (confirmed)
    *   [Link to slides](https://docs.google.com/presentation/d/1Y-xgcRN_HTJdcu7x2RNvpBDlzUYMj8uu7RAGWf5fcgw/edit#slide=id.p2)
    *   [Link to repositories](https://github.com/nirmata/kyverno/)
*   [ 0:00 ]** Release Updates**
    *   Current Release Development Cycle  [Claire Laurence - Release Manager]
    *   [Release Team Shadow applications](https://forms.gle/q69PuRM8V9ByVzEA6) are now open.
    *   1.12, 1.13, 1.14 next patch releases probably late June...TBD based on need.
*   [ 0:00 ] **Contributor Tip of the Week **[Jeffrey Sica] 
    *   Emeritus statuses
    *   [relnotes.k8s.io](http://relnotes.k8s.io)
*   [ 0:00 ] **SIG Updates**
    *   SIG Leads, check out this set of [recommended topics to cover](https://github.com/kubernetes/community/blob/master/events/community-meeting.md#sig-updates) during your update
    *   [Slide template if you need it](https://docs.google.com/presentation/d/1-nTvKCiqu9UvFYUeM6p6RIqHS5-H-u3_x-V4xj_eIWo/edit#slide=id.g401c104a3c_0_0)
    *   Please also check the [SIG Update Schedule](https://docs.google.com/spreadsheets/d/1adztrJ05mQ_cjatYSnvyiy85KjuI6-GuXsRsP-T2R3k)!
    *   Sig VMWare (@steve-wong)
        *    [Slides](https://docs.google.com/presentation/d/1KrqbtD8_RlcKBsjTnCGPvsFUNb9JqgjHeWE8uE0g5M8/edit?usp=sharing)
    *   Sig Apps (@janetkuo)
        *   [Slides](https://docs.google.com/presentation/d/1nP_1taxcqUF8mjhzUnZ3F3VpprlJWDAFPylUbqzeGyo/edit?usp=sharing)
*   [ 0:00 ] **📣Announcements **
    *   [Congrats to Bob Killen for joining GitHub Admin Team](https://github.com/kubernetes/community/pull/3787) [spiffxp]
        *   Thanks to Caleb Miles who is now emeritus
    *   Office Hours next week! [Livestream here](https://www.youtube.com/watch?v=eUWczGnsIAk&feature=youtu.be), click the bell for a reminder
        *   Help us out by [retweeting this](https://twitter.com/castrojo/status/1139168083061460993).
        *   Looking for a west coast streamer so we can do a western session, ping @castrojo if you want to help
    *   **👏 **Shoutouts this week (Check in #shoutouts on slack) **👏**
    *   @vincepri - Huge shoutout to @dhellmann and @dwat for taking 1+ hours today to give great feedback about the Cluster API Bootstrap proposal and helping move the project forward!
    *   @aojea - Big shoutout to @bentheelder for having a working IPv6 CI in #kind
    *   @jberkus - huge shoutout to @Katharine for automating a Release Team role out of existence!  (plus all the Test-Infra folks who helped).
    *   @Jdetiber - Shoutout to @justinsb for cutting the cluster-api v0.1.2 bugfix release


## June 6, 2019 - ([recording](https://youtu.be/6caz2sCVRUI))



*   **Moderators**:  Vallery Lancy [Lyft]
*   **Note Taker**: Jorge Castro [VMware]
*   [ 0:00 ]**  Demo **-- [KubeOne](https://github.com/kubermatic/kubeone) Lifecycle management tool for Kubernetes HA clusters - [Marko] [marko@loodse.com](mailto:marko@loodse.com)] (confirmed)
    *   [Introduction to KubeOne slides](https://docs.google.com/presentation/d/1GYtBRLhkvNiYBbxaaDTQ6W2J0HrDa_VcCZJlaNVYojk/edit?usp=sharing)
    *   [KubeOne asciinema](https://asciinema.org/a/244104)
    *   [KubeOne repository](https://github.com/kubermatic/kubeone)
*   [ 0:00 ]** Release Updates**
    *   Current Release Development Cycle  [Claire Laurence - Release Manager]
        *   Doc PRs merged on tuesday
        *   Cut our first beta yesterday, June 13 cherry pick deadline, 
        *   Release is on for June 17 
        *   Current release status is yellow due to some issues (3)
        *   SIGs, please give the release team your release themes if you have not done so already
        *   Lachlan Evenson will be your 1.16 release lead. 
    *   1.13.7 and 1.14.3 releases coming today (June 6)
*   [ 0:00 ] **SIG Updates**
    *   SIG Leads, check out this set of [recommended topics to cover](https://github.com/kubernetes/community/blob/master/events/community-meeting.md#sig-updates) during your update
    *   [Slide template if you need it](https://docs.google.com/presentation/d/1-nTvKCiqu9UvFYUeM6p6RIqHS5-H-u3_x-V4xj_eIWo/edit#slide=id.g401c104a3c_0_0)
    *   Please also check the [SIG Update Schedule](https://docs.google.com/spreadsheets/d/1adztrJ05mQ_cjatYSnvyiy85KjuI6-GuXsRsP-T2R3k)!
    *   SIG Multicluster (@pmorie - [deck](https://docs.google.com/presentation/d/1ljcaOFKutPO21u1lxtuEDY2LFoIVNXBR3xPbKrfRhfE/edit?usp=sharing))
    *   SIG Windows (@patricklang - [deck](https://docs.google.com/presentation/d/14iWRsVsOld9BuQCj3w0h0On-bwZ6BuskFIVmUDvOAQQ/edit#slide=id.g5b1c1c1f54_0_0))
*   [ 0:00 ] **📣Announcements 📣**
    *   Announcement Foo #1

        **👏 **Shoutouts this week (Check in #shoutouts on slack) **👏**

    *   Stefan Schimanski - shoutout to @liggitt (Jordan Liggitt) for having done insanely many reviews—again and again, in super high quality—this release cycle for CRD+webhook-admission related topics.
    *   Andrew Sy-Kim - Big shoutout to @khenidak (Khaled Henidak) for driving the IPv6 dual stack effort! Some great progress made there this release!
    *   Tim Pepper - Shoutout to @msau42 (Michelle Au) …pretty much every time over the past year I’ve gone to look at a release blocking test failing on storage, @msau42’s a couple hours ahead of me, has the issue triaged and line of site on potential fix if not fix already in test.  Way to represent SIG Storage!!
*   
    *   @stealthybox (Leigh Capili)  and @vincepri (Vince Prignano) props to @Leah (Leah Hanson) for taking stellar notes at lightning speed for Cluster Lifecycle and cluster-api meetings
*   
    *   @vllry (Vallery Lancey) - Thanks to all the contribex folks for all their onboarding/growth resources. and just keeping things running :heart:


## May 30, 2019 - ([recording](https://youtu.be/JEQYdEeRrD8))



*   **Moderators**:  Paris Pittman [SIG Contributor Experience/Google]
*   **Note Taker**: [your name here]
*   [ 0:00 ]**  Demo **-No Demo this week!!
*   [ 0:00 ] **Contributor Tip of the Week **
    *   SIGs are doing live bug scrubs, review how tos, and more - just ask!
        *   API Machinery is Friday! Join their mailing list to get the invite  
*   [ 0:00 ]** Release Updates**
    *   Current Release Development Cycle  [Claire Laurence - Release Manager]
        *   Only 29 folks have responded about docs! Get those docs PRs in!!
        *   Starting daily burn downs next week
        *   1.15 retro doc - talk about timelines/deadlines/opinions there
*   [ 0:00 ] **SIG Updates**
        *   SIG Leads, check out this set of [recommended topics to cover](https://github.com/kubernetes/community/blob/master/events/community-meeting.md#sig-updates) during your update
        *   [Slide template if you need it](https://docs.google.com/presentation/d/1-nTvKCiqu9UvFYUeM6p6RIqHS5-H-u3_x-V4xj_eIWo/edit#slide=id.g401c104a3c_0_0)
        *   Please also check the [SIG Update Schedule](https://docs.google.com/spreadsheets/d/1adztrJ05mQ_cjatYSnvyiy85KjuI6-GuXsRsP-T2R3k)!
        *   Service Catalog (Jonathan B, confirmed)
            *   Moving to kubernetes-sigs from incubator
            *   Team re-org 
            *   Supporting api server version 9 months 
            *   Rewriting docs; has a doc website but most is outdated at this point
            *   New folks from SAP participating; looking for new contributors and a new chair - get in contact with Jonathan (current chair)
        *   IBM Cloud (Sahdev, confirmed)
            *   [Slides](https://docs.google.com/presentation/d/1tGk9Opa4-5fr4kng36NfF-xYw_LRzAdphZWTt9XBVs4/edit?usp=sharing)
*   [ 0:00 ] **📣Announcements 📣**
    *   The Shanghai Contributor Summit Committee is looking for experienced contributors and SIG Leads to lead sessions for the current contributor track.  If you might be available for this, please contact @jberkus or @puja on Slack, or email [jberkus@redhat.com](mailto:jberkus@redhat.com).
    *   Meet Our Contributors is next Wednesday!! 
        *   On demand mentoring from another contributor
        *   [Watch past episodes here](https://www.youtube.com/playlist?list=PL69nYSiGNLP3QpQrhZq_sLYo77BVKv09F)
        *   Yes - you can ask for a live code review (we need advance notice)
        *   Yes - you can ask for a code base tour (we need advance notice)
        *   Join #meet-our-contributors to ask questions and find out more
    *   **#👏 **Shoutouts this week (Check in #shoutouts on slack) **👏**
    *   **hhorl shoutout to:**
        *   @tpepper for consistently being a great context-giver, helper, recruiter, describer, mentor, 
        *   @sumi for publishing our packages, especially on not-so-convenient occasions
        *   @Katharine for jumping in, figuring out, and fixing our testgrid (or testgrid config or something else -- still not sure what the exact problem was :wink:) issue
    *   **Bentheelder Shout-out to **
        *   @aojea @Olav @pbnj (and anyone I missed!) for helping answer lots of questions in #kind. I especially appreciate it now while we have an influx of new members and I'm jetlagged out after KubeCon. Thank you all :slightly_smiling_face:
    *   **bentheelder shout-out to **
        *   @paris @tpepper for organizing and @castrojo for hosting the Networking + Mentoring sessions at KubeCon, really awesome experience :slightly_smiling_face:
    *   **gsaenger huge thanks for **
        *   @Deb Giles for making the contributor summit run smooth like butter, especially given some unique challenges with the location!
    *   **Parispittman shout out to:**
        *   Diversity Lunch participants, leads
        *   Mentoring Session participants, leads


## May 16, 2019 - (recording)



*   **Moderators**:  Dawn Foster [SIG Contributor Experience/Pivotal]
*   **Note Taker**: Jorge Castro [SIG Contributor Experience/VMware]
*   [ 0:00 ]**  Demo **-- Metal<sup>3</sup>: Bare metal host management for Kubernetes backed by OpenStack Ironic [Chris Hoge, chris@openstack.org] (confirmed)
    *   [Link to video](https://www.dropbox.com/s/03z0ezgnfppguwz/metal3.mp4?dl=0) - (time lapsed as it’s on real bare metal) 
        *   [A similar demo done real-time and live](https://www.youtube.com/watch?v=Nzq2S53nk9U), only shows BareMetalHosts and not the Machine definition. Also skips cleaning and inspection.
    *   [Metal3 Repo](https://github.com/metal3-io/metal3-docs) - pronounced “metal kubed”
    *   Ironic controlling the infra, small set of services, running in podman in this example, but can run in k8s. 
    *   
*   [ 0:00 ]** Release Updates**
    *   Current Release Development Cycle  [Claire Laurence - Release Manager]
    *   1.15 - No change at # of enhancements being tracker, 46. 23 alpha, 19 beta, 4 stable. 
    *   We cut the 1.15 branch and first 1.15 beta. 
    *   1.15 jobs created, 1.11 jobs removed
    *   For next week:
        *   F2F session during the contributor summit
        *   No major milestones
        *   Burndown starts May .28
        *   Patch Release Updates
            *   None this week
*   [ 0:00 ] **Contributor Tip of the Week **[Nikhita Raghunath or Christoph Blecker] 
    *   A reminder to set your GitHub status to “Busy” only if you are _really_ busy, since this will now prevent automatic PR review requests. Please take care in how you use your busy status to avoid overloading other reviewers. See [thread for details](https://groups.google.com/forum/#!topic/kubernetes-dev/bXHK8l3D6l0). 
*   [ 0:00 ] **KEP of the Week** [Kubernetes Enhancement Proposals]
    *   [Even Pods Spreading](https://github.com/kubernetes/enhancements/blob/master/keps/sig-scheduling/20190221-even-pods-spreading.md) - [Implementable] - [SIG Scheduling](https://github.com/kubernetes/community/tree/master/sig-scheduling), Bobby (Babak) Salamat (@bsalamat) 
        *   Allow users to specify what topology domain a pod can be spread over.
        *   Spread a pod “Among zones, or among nodes” or any arbitrary thing.
        *   Interpod-anti-affinity works, but limited to only 1 pod per topology domain.
        *   This allows you to spread as many pods as you want across all your topology domains. 
        *   API bandwidth is a problem, this feature is at risk for this release.
            *   Jordan Liggitt has gone above and beyond trying to help fix this problem. 
            *   API review is complex, takes people a long time to become a competent API reviewer. 
*   [ 0:00 ] **SIG Updates**
    *   SIG Leads, check out this set of [recommended topics to cover](https://github.com/kubernetes/community/blob/master/events/community-meeting.md#sig-updates) during your update
    *   [Slide template if you need it](https://docs.google.com/presentation/d/1-nTvKCiqu9UvFYUeM6p6RIqHS5-H-u3_x-V4xj_eIWo/edit#slide=id.g401c104a3c_0_0)
    *   Please also check the [SIG Update Schedule](https://docs.google.com/spreadsheets/d/1adztrJ05mQ_cjatYSnvyiy85KjuI6-GuXsRsP-T2R3k)!
    *   SIG Cluster Lifecycle [Lucas Käldström] (confirmed)
        *   [Link to Slides](https://docs.google.com/presentation/d/1QUOsQxfEfHlMq4lPjlK2ewQHsr9peEKymDw5_XwZm8Q/edit)
    *   SIG OpenStack [Chris Hoge] (confirmed)
        *   [Link to Slides](https://docs.google.com/presentation/d/1XmFIoGBsRoEaitVpdl3oNHsv1jhgTO1opITeHaZ-1PA/edit?usp=sharing) 
    *   SIG Auth [Tim Allclair] (confirmed) 
*   [ 0:00 ] **📣Announcements 📣**
    *   SIG Meet and Greet and Contributor Summit Update - Paris PIttman

        **#👏 **Shoutouts this week (Check in #shoutouts on slack) **👏**

*   mrbobbytables: shoutout to @claudiajkang, @Felipe, and @irvifa for localizing the contributor cheatsheet to Korean, Portuguese and Bahasa Indonesian, and @rui for organizing the effort!
*   gsaenger: Shoutout to @jonasrosland for tireless slide edit magic
*   Bentheelder: shoutout to @mrbobbytables for driving home the subproject site hosting process, it's almost done document all the things!! 
*   Jonasrosland: *HUGE SHOUTOUT* to @paris @Dawn Foster @castrojo @Deb Giles @ihor.dvoretskyi @coderanger @mrbobbytables for an amazing job planning out the Kubernetes Contributor Summit in Barcelona these past months!
*   Jonasrosland: And an *enormous shoutout* to @tpepper and @gsaenger for updating and taking on the role of workshop leads for Kubernetes Contributor Summit BCN!


## May 9, 2019 - [recording](https://www.youtube.com/watch?v=9JPPLnFG3fs)



*   **Moderators**:  Jorge Castro  [SIG Contributor Experience]
*   **Note Taker**: First Last [Company/SIG]
*   [ 0:00 ]** Release Updates (Going first this week)**
    *   Current Release Development Cycle  [Claire Laurence - Release Manager]
        *   Third alpha cut this week
        *   Next week will be the first beta
            *   1.15 branch cut
            *   1.15 jobs created
            *   1.11 jobs removed
        *   Tracking for 47 enhancements, but we’ll see how that changes closer to code freeze (May 30th?)
            *   Alin the next day or two
            *   Important for communicationscommunications and blog posts around the 1.15 release
        *   SIG leads:
            *   Start thinking about different themes for your SIGs
            *   If you haven’t heard from SIG Release, you will 
        *   For those at Kubecon EU
            *   Meetup on day 1
    *   [Patch Release Schedule](https://git.k8s.io/sig-release/releases/patch-releases.md) Updates
        *   v1.14.2 coming soon cherry pick merge deadline 5/10 ahead of 5/14 release
        *   v1.13.6 released yesterday...5/8
        *   v1.12.8 released 4/24, next TBD May?
*   [ 0:00 ]**  Demo **-- Stefan Prodan, [Flagger](https://github.com/weaveworks/flagger) (confirmed)
    *   Link to slides
    *   Link to repositories
        *   [GitHub](https://github.com/weaveworks/flagger)
    *   Overview
        *   A kubernetes operator that automates promotion of canary deployments in order to route traffic
        *   Goal is to make deployments observable (plugins for slack, pagerduty, etc.)
        *   Workflow is driven using git (leveraging reviews before applying changes to infrastructure for example)
        *   Grafana dashboard and alerting is included
        *   Gracefully promotes or rolls back deployments based on configurable success rates
        *   Also supports A/B testing, based on specific HTTP headers or cookies
    *   Questions
        *   Is the plan to offer this to the K8s community, CNCF, or some other upstream location?
            *   open source project under weaveworks
            *   plan to submit to CNCF sandbox at some point
            *   API is still alpha
*   [ 0:00 ] **Contributor Tip of the Week **[First Last][ ](https://groups.google.com/forum/#!topic/kubernetes-sig-contribex/jrvAoOjdOtA)
    *   [Github Status review stuff](https://groups.google.com/forum/#!topic/kubernetes-sig-contribex/jrvAoOjdOtA)
    *   Looking for feedback by Wednesday (2019-05-15)
*   [ 0:00 ] **KEP of the Week** [Kubernetes Enhancement Proposals]
    *   [Andrew Kutz] - Kubeadm Machine/Structured Output
        *   WIP: [https://github.com/kubernetes/enhancements/pull/1054](https://github.com/kubernetes/enhancements/pull/1054)
        *   Looking for feedback
        *   Aiming for alpha in 1.16
        *   Looking to add structured output to kubeadm for better tooling/integration.
            *   Need to be able to parse in deterministic way.
        *   support json, yaml, and go-templates
        *   Will be updating the KEP to emit versioned objects
        *   Looking to promote to beta in 1.17 if people are happy with it
*   [ 0:00 ] **SIG Updates**
    *   SIG Leads, check out this set of [recommended topics to cover](https://github.com/kubernetes/community/blob/master/events/community-meeting.md#sig-updates) during your update
    *   [Slide template if you need it](https://docs.google.com/presentation/d/1-nTvKCiqu9UvFYUeM6p6RIqHS5-H-u3_x-V4xj_eIWo/edit#slide=id.g401c104a3c_0_0)
    *   Please also check the [SIG Update Schedule](https://docs.google.com/spreadsheets/d/1adztrJ05mQ_cjatYSnvyiy85KjuI6-GuXsRsP-T2R3k)!
    *   SIG [AWS](https://docs.google.com/presentation/d/1ybq1YrPf45ww069wYVFDfJd6uTEXaFVVWv5540ZQDEo/edit#slide=id.p1) [Nishi Davidson/Justin Santa Barbara] (confirmed)
        *   [Slides](https://docs.google.com/presentation/d/1ybq1YrPf45ww069wYVFDfJd6uTEXaFVVWv5540ZQDEo/edit#slide=id.p1)
        *   [SIG Intro session](https://sched.co/MPhs) @ Kubecon
    *   SIG [Contributor Experience](https://docs.google.com/presentation/d/1mlIdjFYC0ZXhKvDAyCCvVnBwWMddupu3RHVdyE1UuKQ/edit?usp=sharing) [Paris Pittman] (confirmed) 
        *   [SIG Intro session @ KubeCon](https://kccnceu19.sched.com/event/MPhR/intro-contributor-experience-sig-elsie-phillips-red-hat-paris-pittman-google)
        *   [SIG Deep Dive @ KubeCon](https://kccnceu19.sched.com/event/MPjE/deep-dive-contributor-experience-sig-elsie-phillips-red-hat-paris-pittman-google)
        *   new contact: [contributors@kubernetes.io](mailto:contributors@kubernetes.io)
        *   building teams:
            *   triage team
            *   events team
            *   marketing team
        *   looking for apac coordinator
        *   mentoring / succession planning / contributor growth
            *   building programs around personas
            *   succession plans for roles (e.g., shadows?)
            *   different learning styles (videos, documentation etc)
            *   [speed mentoring at Barcelona](https://kccnceu19.sched.com/overview/type/Networking+%2B+Mentoring)
            *   SIG Meet and Greet [link]
            *   GSOC - 6 interns for Kubernetes project
            *   [Meet our Contributors](https://www.youtube.com/playlist?list=PL69nYSiGNLP3QpQrhZq_sLYo77BVKv09F)
            *   Exploring [Community Bridge](https://people.communitybridge.org/) from LF
            *   Expanding mentoring @ scale initiatives
                *   [live API / code reviews](https://www.youtube.com/watch?v=yqB_le-N6EE)
                *   [code base tours](https://github.com/kubernetes/community/issues/2812)
                *   new contributor office hours
        *   Automation / GitHub Management
            *   Improve automation around OWNERS files
            *   Audit of inactive owners
            *   remove inactive reviewers/approvers from owners
            *   Emeritus is for domain exports who may not be active in area, and will be ignored by prow, but can still be referenced if needed.
            *   Improved owners file hygiene (only members added to owners files)
            *   fejta-bot issue lifecycle automation enabled on all orgs
            *   needs-rebase plugin enabled on all orgs
            *   trigger plugin adds needs ‘ok-to-test’ label
            *   restricted @mentions and other messages in commit messages
            *   30 new repos created from last community update
            *   team membership managed through k/org
            *   process for adding subproject sites (netlify) in flight
        *   slack infra
            *   “gitops” for slack management
            *   shoutout to @katharine for being awesome
            *   report message feature added
        *   community site relaunched [link]
        *   Community management
            *   improving training etc for sig chairs and TLs
            *   Assist in bootstrapping and disbanding working groups
        *   events
            *   Barcelona Contributor Summit
                *   Seats still available for New Contributors
        *   communication
            *   Moderators have more than doubled in size since last update
        *   Contributor documentation
            *   contributor/developer guide improved
        *   Future
            *   tie sigs.yaml to everything
            *   build more mentoring programs
            *   more training
    *   [SIG Scheduling](https://docs.google.com/presentation/d/1_q3I2grii7_E9KIxKZ3PMy7Zj2-NiUCXnkJKqRU0QR0/edit#slide=id.g401c104a3c_0_0)  [Bobby Salamat] (confirmed) 
        *   [SIG Intro session @ KubeCon](https://kccnceu19.sched.com/event/MPiV/intro-scheduling-sig-da-ma-shivram-srivastava-huawei)
        *   [SIG Deep Dive @ KubeCon](https://kccnceu19.sched.com/event/MPkg/deep-dive-scheduling-sig-babak-salamat-google)
        *   last cycle (1.14)
            *   improve performance and stability of scheduler
            *   3x performance improvement
                *   100/pods/second in 5000 node clusters
            *   pod priority and preemption graduated to stable
            *   improve scheduling fairness
                *   add back-off mechanism to unschedulable pods
            *   fixed a few race conditions
        *   future (1.15)
            *   improve workload reliability
                *   new feature: even pod spreading [link]
                    *   how many pods / arbitrary failure domain
                    *   deprecate some inter-pod anti-affinity capabilities in the future
            *   improve extensibility of the scheduler
                *   pluggable scheduler [link]
                *   alpha KEP
            *   better pod priority for batch workloads
                *   support non-preempting priority for batch workloads
                *   goes to head of queue
            *   supporting Lt / Gt operators for affinity
        *   How these plans affect you
            *   Generally backwards compatible
            *   Cluster autoscaler may have issues with new scheduling framework
*   [ 0:00 ] **📣Announcements 📣**
    *   testgrid alerts for release-blocking jobs, [deadline Tuesday May 14th](https://github.com/kubernetes/sig-release/issues/441#issuecomment-490560650)
    *   [KubeCon + CloudNativeCon CFP](https://events.linuxfoundation.org/events/kubecon-cloudnativecon-north-america-2019/cfp/) for NA 2019 San Diego is open!
    *   Announcement Foo #1

        **👏 **Shoutouts this week (Check in #shoutouts on slack) **👏**

*   dims - @Damini Satya @jimangel @zparnold @sbezverk @jrondeau Congrats on your Google Open Source Peer Bonus win for your work on Kubernetes! [https://opensource.googleblog.com/2019/04/google-open-source-peer-bonus-winners.html](https://opensource.googleblog.com/2019/04/google-open-source-peer-bonus-winners.html)
*   gsaenger - Shoutout and :sparkles: to @soltysh for the most amazing codebase walkthrough to get me ready to share with new contributors at KubeCon! Thank you so much, I learned a lot!
*   JeremyWx - Big shoutout to @atuvenie for helping me with an aks-engine problem!  After banging my head on my desk for most of the week she pointed out I was using an version with a bug.  My head and my desk, thank you very much!!


## May 2, 2019 - ([recording](https://youtu.be/R8oLTmyZCrU))



*   **Moderators**:  Lachlan Evenson (SIG-PM, 1.15 release team)
*   **Note Taker**: Solly Ross (Google / Kubebuilder)
*   [ 0:00 ]**  Demo **-- k8dash [Eric Herbrandson ([eric@herbrandson.com](mailto:eric@herbrandson.com))] (confirmed)
    *   [Link to slides](https://docs.google.com/presentation/d/1XWGPVeWD_eUtwksghRQAv9VUzv-3oUf73WA6UEyY0ks/edit?usp=sharing)
    *   [Link to repo](https://github.com/herbrandson/k8dash)
    *   Alternative k8s dashboard
    *   Native OIDC integration (no proxy)
    *   Uses watch APIs to update in real-time (no refreshing)
    *   Filterable, sortable views for
        *   Metrics: Resource usage using websockets API -- lots of graphs integrated into other views
        *   Pods
        *   Workloads (see live rollouts, etc)
        *   Storage
        *   Secrets (blurred so still easy to copy)
        *   RBAC
    *   Editing
        *   YAML editor
            *   Context-aware documentation in YAML editor
            *   Can kubectl-apply via UI
        *   Scale
        *   Delete
    *   Views are response (works fine in mobile, nicely resizes to fit)
        *   Debug pods on the go!
    *   **Looking for feedback on:**
        *   **What’s missing for your team?**
        *   How to promote within the community
    *   Questions
        *   Q: What’s the difference between k8dash and kubernetes/dashboard
            *   A: Real-time updates are the big difference (no refreshes, easy to see live updates)
            *   A: OIDC integration
            *   A: uses metrics-server for stats, not heapster (which is deprecated)
*   [ 0:00 ]** Release Updates**
    *   Current Release Development Cycle  [Claire Laurence - Release Manager]
    *   V1.15.0-alpha.2
        *   Week 4 of release cycle (past/current week)
            *   2nd alpha release (Monday)
            *   Enhancements freeze was Tuesday
                *   43 enhancements for 1.15 before freeze
                *   35 enhancements for 1.15 after (including 5 approved exceptions)
        *   Week 5 (upcoming week)
            *   1.11.0 jobs removed (May 7)
            *   3rd alpha (May 7th)
    *   Patch Release Updates
        *   v1.14.2 tentative 5/14
        *   v1.13.6 coming 5/8
        *   v1.12.8 released 4/24
        *   v1.11.10 released 5/1 - this was the final 1.11 patch release
*   [ 0:00 ] **KEP of the Week** [Kubernetes Enhancement Proposals]
    *   [Add revised IPv4/IPv6 dual stack KEP](https://github.com/kubernetes/enhancements/pull/808/files) - [Provisional seeking implementable] 
    *   follow up #k8s-dual-stack on slack or [k8s discuss](https://discuss.kubernetes.io/t/kubernetes-ipv4-ipv6-dual-stack-support-status/4974)
    *   Motivation: enable dual stack support in kubernetes -- pods with ipv4 and ipv6 support addresses side-by-side
        *   IPv6-only has existed for a while in Kubernetes
        *   Dual stack is common migration path
    *   Multi-release KEP
        *   1.15 target is to get multiple IP addresses on a pod, all nodes to have multiple CIDRs
*   [ 0:00 ] **SIG Updates**
    *   SIG Storage [Saad Ali] (confirmed) 
        *   [[slides here](https://docs.google.com/presentation/d/1IApIw92SNEt-iE8YZTnP0cxecrj0LQ-3kbDh70E661c/edit?usp=sharing)]
        *   Kubernetes 1.14
            *   Local PVs moved to GA (local-to-node disk as PV, like hostPath but with scheduler support, blog on k8s.io [link here])
            *   CSI improvements: moving towards feature parity with in-tree volumes
                *   [beta] Raw block volumes (block device in container instead of FS)
                *   [beta] topology (support for expressing that volume is only available to certain nodes for scheduling)
                *   [alpha] resizing (request more size on volume)
            *   [alpha] in-tree → CSI migration (adapters to point in-tree plugins to CSI so we can remove third-party code without breaking users)
            *   Pluggable e2e test framework to make writing tests for all volume plugins easier (lots of volume plugins were untested because tests were very specific)
        *   Kubernetes 1.15
            *   [beta] in-tree → CSI migration (may end up staying alpha)
            *   CSI features
                *   [beta] resizing (may end up staying alpha depending on KEP)
                *   [alpha] ephemeral inline volumes (better support for local, ephemeral volumes like secrets or configmaps in CSI without needing to create a PVC first, inline in pod definition instead)
                *   Volume capacity and usage metrics (exists for in-tree volumes, need support for CSI)
            *   Snapshots
                *   CSI-only feature
                *   [alpha] pause/resume hooks for application-level consistency (instead of just crash consistency)
                *   [in design] volume consistency groups -- multi-volume snapshots
            *   [alpha] Cloning (immediately duplicate volume copy-on-write style if supported by plugin)
            *   [redesign] volume attach limits (most storage systems have limits about how many volumes can be attached to a node, scheduler needs to be aware of this, needs improvement for CSI)
        *   Come learn/participate
            *   Lots of presentations at KubeCon EU 
            *   [https://git.k8s.io/community/sig-storage](https://git.k8s.io/community/sig-storage) for more info, meetings
    *   SIG Docs  [Zach Corleissen/Jennifer Rondeau] (confirmed) 
        *   [[slides here]](https://docs.google.com/presentation/d/12gVJH1kr825YTkjTGT3SwxlfaiQKQXJWgfoBve327Zk/edit?ts=5cca20d7#slide=id.g401c104a3c_0_0)
        *   Last cycle
            *   1.14 docs :-)
            *   New meta-documentation on docs release lead
            *   +6-7 more localizations
                *   Starting more meta-documentation on localization
                *   Lots of good fixes to english docs when translation issues are encountered as well
            *   More roles, mentoring support -- help bring new folks on and get them contributing faster/more easily
            *   **WG-ish group about how to organize security content in docs (talk to @zparnold)**
                *   **Get involved: #sig-docs-security**
            *   Figuring out subdomain-hosting for subprojects (e.g. kind.k8s.io)
            *   Getting more tech writers for pain points in the docs (e.g. “pick the right solution”)
        *   Upcoming plans
            *   Mentorship -- path to approver, new contributor ambassador
                *   Better path for first issue → merged PRs
            *   1.15 docs :-)
            *   Better issue triage
        *   **Upcoming doc sprints**
            *   KubeCon EU (not WriteTheDocs, since it conflicts with KubeCon EU)
            *   KubeCon Shanghai
            *   OpenSource Summit Tokyo
        *   Using shadows for leads due to lead visibility (comes with a good pun, see the recording)
        *   Kubernetes Blog is officially subproject of SIG Docs
            *   @kbarnard10 is forming a team
        *   How to contribute:
            *   See [https://k8s.io/docs/contribute](https://k8s.io/docs/contribute) (thanks Misty :-) )
            *   [https://github.com/kubernetes/website/projects/3](https://github.com/kubernetes/website/projects/3)
            *   [https://git.k8s.io/community/sig-docs](https://git.k8s.io/community/sig-docs) 
*   [ 0:00 ] **📣Announcements 📣**
    *   

        👏 Shoutouts this week (Check in #shoutouts on slack) 👏

    *   paris - thanks to @deads2k and @soltysh for joining us today for the first meet our contributors session. tons of great answers to API and CLI contributing questions - thanks for being mentors!
    *   Soltysh - big thanks to @paris and @castrojo for organizing meet our contributors


## April 25, 2019 - ([recording](https://youtu.be/eXM4lNeV2D4))



*   **Moderators**:  Jorge Castro
*   **Note Taker**: This Could be You! [Company/SIG]
*   [ 0:00 ]**  Demo **-- Daniel Messer, Product Manager, [dmesser@redhat.com](mailto:dmesser@redhat.com) and Diane Mueller, Director, Community Development  [dmueller@redhat.com](mailto:dmueller@redhat.com) - Automated Day 2 Operations on Kubernetes using Operators
    *   Link to slides: [https://docs.google.com/presentation/d/1sz3PkfDu7-FhM8qwjfJXnJCJgdC-4pTUif4KjrcFdyQ/](https://docs.google.com/presentation/d/1sz3PkfDu7-FhM8qwjfJXnJCJgdC-4pTUif4KjrcFdyQ/)
    *   [https://github.com/operator-framework/getting-started - Overview](https://github.com/operator-framework/getting-started)
    *   [https://operatorhub.io/ - Community Operator Hub](https://operatorhub.io/)
    *   [https://www.katacoda.com/openshift/courses/operatorframework - Browser-based tutorials](https://www.katacoda.com/openshift/courses/operatorframework)
    *   [https://groups.google.com/forum/#!forum/operator-framework - Discussion Forum](https://groups.google.com/forum/#!forum/operator-framework)
    *   [https://kubernetes.slack.com/messages/kubernetes-operators - Slack Channel](https://kubernetes.slack.com/messages/kubernetes-operators)
    *   
*   [ 0:00 ]** Release Updates**
    *   Current Release Development Cycle  [Claire Laurence - Release Manager]
    *   
    *   Patch Release Updates
        *   1.11.10 April 30, 2019; cherry pick deadline Friday April 26
        *   1.12.8 April 24, 2019
        *   1.13.6 April ??, 2019; cherry pick deadline April ??
        *   1.14.3 May 2019 TBD
*   [ 0:00 ] **Contributor Tip of the Week **[Jorge Castro]
    *   [Contributor Cheatsheet](https://github.com/kubernetes/community/blob/master/contributors/guide/contributor-cheatsheet/README.md)
    *   Help [translate this](https://github.com/kubernetes/website/issues/13989) into another language
*   [ 0:00 ] **SIG Updates**
    *   SIG Leads, check out this set of [recommended topics to cover](https://github.com/kubernetes/community/blob/master/events/community-meeting.md#sig-updates) during your update
    *   [Slide template if you need it](https://docs.google.com/presentation/d/1-nTvKCiqu9UvFYUeM6p6RIqHS5-H-u3_x-V4xj_eIWo/edit#slide=id.g401c104a3c_0_0)
    *   Please also check the [SIG Update Schedule](https://docs.google.com/spreadsheets/d/1adztrJ05mQ_cjatYSnvyiy85KjuI6-GuXsRsP-T2R3k)!
    *   SIG PM [Stephen Augustus] 
        *   Slideeeesssss: [https://docs.google.com/presentation/d/1ULPn-UH_SnIH5wO9qqTQFytW_UUiKEaRqdf-burte1g/edit?usp=sharing](https://docs.google.com/presentation/d/1ULPn-UH_SnIH5wO9qqTQFytW_UUiKEaRqdf-burte1g/edit?usp=sharing)
    *   SIG Testing  [Erick Fejta] 
        *   Making sure owners are org members. 
        *   Prow updates itself every day now, when there’s a commit fejta-bot will update itself (That’s not creepy at all! --jorge)
        *   Tackle, inside of prow, tackle utility will deploy prow into your cluster and configure github hooks. 
        *   Thano - converts a prow command into a docker run command so you can run it locally. 
        *   Trying to create a roadmap for prow - improving monitoring, use more of the standard k8s APIs, moving configs from jobs into the repo itself, still collecting info, so if you have an idea, please find us on #sig-testing and we’ll point you to the doc.
        *   Building - experimenting with remote execution in bazel, smooth out the CPU spikiness by moving those to remote clusters, so we can run more things. 
        *   Test infra is using go modules - trying to auto patch and upgrade over modules. Minor upgrades are a hassle due to Golang semver issues.

[ 0:00 ] **📣Announcements 📣**



    *   Announcements
    *   WG LTS Survey ending on April 26th.
        *   This survey was created by the[ LTS Working Group](http://git.k8s.io/community/wg-lts) of the Kubernetes project. The purpose of this survey is to understand the challenges faced by various types of stakeholders with respect to the current release cadence of Kubernetes project. The survey questions are classified based on the stakeholder category.
        *   [https://www.surveymonkey.com/r/kubernetes-support-survey-2019](https://www.surveymonkey.com/r/kubernetes-support-survey-2019)
    *   We now have a `#pr-reviews` slack channel as a last resort if your PR is stuck. We would love to have folks who can triage/review as well to join the channel to 	help wither fellow contributors. The idea is to help get someone get a PR “ready” and get the right SIGs/Reviewers/Approvers involved.
    *   **👏 **Shoutouts this week (Check in #shoutouts on slack) **👏**
        *   Nihita would like to “Shoutout to @eduar for expanding the contributor cheatsheet, and for following up on it even after his Outreachy internship ended :tada:”
        *   zacharysarah would like to “Shoutout to @bradtopol for his willingness to step out of a Barcelona speaking slot in order to make room for a local leader to step in. That’s some generous leadership! :partyk8s:
        *   lukaszgryglicki would like to “Shoutout to @Eugene Glotov for helping me with AWS/ELB/EKS stuff - T H A N K S !


## April 18, 2019 - ([recording](https://youtu.be/xOsGC9m164g))



*   **Moderators**:  Bob Killen (sub for Paris Pittman) [SIG-Contributor Experience]
*   **Note Taker**:  [Company/SIG]
*   [ 0:00 ]**  Demo **-- Dan Lorenc - Tekton Pipeline CRD - A K8s-native Pipeline resource.(confirmed)
    *   Link to slides
    *   Link to repositories
    *   Declarative CI/CD system making native use of Kubernetes resources
    *   Contributors: Google, Pivotal, CloudBees, Red Hat, IBM and more
    *   Has Task CRD that defines sequence of steps inside a pod
    *   “Type Safe” Ci/CD system
    *   Pipeline CRD builds a graph/DAG of the tasks
*   [ 0:00 ]** Release Updates**
    *   Current Release Development Cycle  [Claire Laurence - Release Manager]
    *   Patch Release Update
        *   enhancements freeze next week
*   [ 0:00 ] **Contributor Tip of the Week **[Katharine Berry] 
    *   Slack GitOps
    *   Docs: [ kubernetes/community/communication/slack-guidelines](https://github.com/kubernetes/community/blob/master/communication/slack-guidelines.md)
    *   Config: [kubernetes/community/communication/slack-config](https://github.com/kubernetes/community/tree/master/communication/slack-config)
    *   Tooling Repo: [kubernetes-sigs/slack-infra](https://github.com/kubernetes-sigs/slack-infra)
    *   Channel management can be delegated
    *   Possibly rename a channel if something changes outside of git / source of truth.
*   [ 0:00 ] **KEP of the Week** [Kubernetes Enhancement Proposals]
    *   [SIG Cluster Lifecycle] kubeadm config v1beta2] - [WIP / provisional] - [https://github.com/kubernetes/enhancements/pull/969](https://github.com/kubernetes/enhancements/pull/969)
*   [ 0:00 ] **SIG Updates**
    *   SIG Azure [Stephen Augustus]  (confirmed)
        *   Slides: [https://docs.google.com/presentation/d/12CaK1vkjGXZ_nadyFWwwG2XzS_2J8QozUTqPjtWbwZM/edit?usp=sharing](https://docs.google.com/presentation/d/12CaK1vkjGXZ_nadyFWwwG2XzS_2J8QozUTqPjtWbwZM/edit?usp=sharing)
        *   last cycle
            *   working on testing out-of-tree cloud provider
            *   MVP of cluster api for azure
        *   next cycle
            *   continue work on out-of-tree providers
            *   prep work for sig cloud provider consolidation
            *   OPA/Gatekeeper
            *   Large scale cluster support and test on Azure (5,000 nodes)
            *   build prod-ready clyster api implementation.
        *   Looking for more contributors
    *   SIG Release  [Stephen Augustus] (confirmed) 
        *   Slides: [https://docs.google.com/presentation/d/1znI75gT1shim_CiPA2uDhhqhxA76AdnAxlp78eTV_cA/edit?usp=sharing](https://docs.google.com/presentation/d/1znI75gT1shim_CiPA2uDhhqhxA76AdnAxlp78eTV_cA/edit?usp=sharing)
        *   last cycle
            *   release 1.14
            *   keps are now a requirement for in-tree kubernetes enhancements
            *   improved KEP template
            *   introduced a questionnaire for release team shadow process
                *   no longer first come first served shadow selection process.
                *   iterative improvements being made to questionnaire for future releases.
            *   improving release engineering (branch and patch release management)
                *   now patch-release team
            *   licensing subproject team
                *   Everything related to tracking licenses to ensure compliance with CNCF/LF.
        *   next cycle
            *   Improve KEP tracking process.
            *   Staff Release Engineering and Licensing teams
            *   Improve feedback loop for KEPs with sig-pm
            *   Ensure there is concrete membership criteria for the Patch Release/Release Team
            *   Establish policy for tracking out-of-tree enhancements
            *   Establish policy for release artifacts
            *   work with wg k8s-infra-team on creating visible / community managed artifacts.
            *   revisit charter to define in/out-of-scope
            *   build process around orgt-wide license management
    *   SIG Big Data  [First Last] 
*   [ 0:00 ] **📣Announcements 📣**
    *   [Barcelona Contributor Summit](https://events.linuxfoundation.org/events/contributor-summit-europe-2019/) schedule locked in. ([https://contsummiteu19.sched.com/](https://contsummiteu19.sched.com/)) - Invites to sched will be sent out soon (**Must **be registered through contrib summit site to get sched invite).
        *   Reminder that current contributor content is SIG/subproject F2F only.
            *   Kubebuilder subproject
            *   Release Team meeting
            *   SIG-CLI
            *   SIG-Cloud Provider
            *   SIG-Cluster Lifecycle
            *   SIG-IBM Cloud
            *   SIG-Networking
            *   SIG-PM
            *   SIG-Scheduling
            *   SIG-UI
            *   SIG-VMware
            *   SIG-Windows
*   **👏 **Shoutouts this week (Check in #shoutouts on slack) **👏**
    *   lachie83 - shoutout to @hogepodge @lavalamp for providing awesome sig updates today during the community meeting
    *   lachie83 - shoutout to @patricklang for the awesome windows on k8s demo during the community meeting today (this note is from last week --ed.)
    *   zacharysarah - Shoutout to @remyleone for not only leading the French localization of docs, but helping the Spanish and Indonesian projects launch as well.
    *   jdetiber - Shoutout to @vincepri for the great work on getting the Cluster API project scope and objectives documented and working through seemingly unending feedback to get us to the point that the document can be merged. (edited) 


## April  11, 2019 - ([recording](https://youtu.be/zVze4sDNI_E))



*   **Moderators**:  Lachlan Evenson [SIG-PM]
*   **Note Taker**: Solly Ross [Google/KubeBuilder]
*   [ 0:00 ]**  Demo **-- What’s New in Windows Containers in K8s [SIG-Windows], Patrick Lang, Senior Software Engineer, SIG-Windows co-chair - @patricklang
    *   _[no slides, just a fun demo]_
        *   Kube 1.14 cluster with both linux and windows server nodes
            *   Deploying container using IIS (windows app)
            *   Using node selector to make sure workloads land on particular nodes, can also use taints on the windows nodes
            *   Can scale normally, can use standard services (e.g. loadbalancer), can use ingress as well
            *   Existing linux workloads still running fine (e.g. kube-system services)
        *   Building a .NET app on a windows VM, package in container from windows VM
            *   Deployed from Helm chart from same windows VM
            *   Running SQL server on Linux node, run app on Windows node
            *   Can run kubectl, etc from the Windows node
            *   The app works! :-)
    *   Additional Resources
        *   Kubernetes documentation for Windows [https://kubernetes.io/docs/setup/windows/](https://kubernetes.io/docs/setup/windows/)
        *   Sample deployed with Helm [https://github.com/PatrickLang/fabrikamfiber/tree/helm-2019-mssql-linux](https://github.com/PatrickLang/fabrikamfiber/tree/helm-2019-mssql-linux) 
        *   More samples: [https://github.com/PatrickLang/Windows-K8s-Samples](https://github.com/PatrickLang/Windows-K8s-Samples) 
        *   More in-depth Presentation at Kubecon [https://kccna18.sched.com/event/GrRU/tutorial-deploying-windows-apps-with-draft-helm-and-kubernetes-patrick-lang-jessica-deen-microsoft-limited-seating-available-see-description-for-details](https://kccna18.sched.com/event/GrRU/tutorial-deploying-windows-apps-with-draft-helm-and-kubernetes-patrick-lang-jessica-deen-microsoft-limited-seating-available-see-description-for-details) 
            *   Recording on YouTube - [link](https://www.youtube.com/watch?v=i4SrVONbghA&t=1s) 
            *   Sources - [https://github.com/PatrickLang/KubernetesForWindowsTutorial](https://github.com/PatrickLang/KubernetesForWindowsTutorial) 
    *   Setup steps for my Azure cluster using aks-engine - [link](https://github.com/Azure/aks-engine/blob/master/docs/topics/windows.md)
    *   Windows nodes are GA as of 1.14
*   [ 0:00 ]** Release Updates**
    *   Current Release Development Cycle  [Claire Laurence - Release Manager] (Confirmed)
        *   1.15 release cycle began Monday (April 8)
        *   Working on getting schedule finalized (README link here)
        *   11-week cycle (GA on June 17, to avoid releasing during KubeCon Shanghai)
        *   Enhancements tracking starting next week (April 15th-ish)
            *   File issues for including in the 1.15
            *   Enhancements freeze on April 30
                *   Must have open issue in the 1.15 milestone
            *   Please have test plans and graduation criteria
        *   1.15 alpha 1 next tuesday
        *   Release notes starting on April 23rd
    *   Patch Release Updates
        *   1.14.1 released 2019-04-08
        *   1.13.6 TBD April 2019
        *   1.12.8 cherry pick deadline 2019-04-19, release target 2019-04-22
        *   1.11.10 ...officially past 9 months typical support window, but possibly one last release to come TBD April 2019
*   [ 0:00 ] **Contributor Tip of the Week **[Jorge Castro] 
    *   What do I OWNERs and how do I check? 
    *   [https://cs.k8s.io](https://cs.k8s.io) (hound) allows code search through the k8s codebase
    *   Search accepts RegEx
    *   Can search for yourself and make sure you’re in the appropriate OWNERS files and make sure you’re not still listed for things you’re not working on
    *   [https://cs.k8s.io/?q=castrojo&files=OWNERS](https://cs.k8s.io/?q=castrojo&files=OWNERS) - sub in your github name to see which files you’re in. 
*   [ 0:00 ] **KEP of the Week** [Kubernetes Enhancement Proposals]
    *   [KEP for release notes improvements](https://github.com/kubernetes/enhancements/pull/928) - [provisional] - sig-release
        *   New site for end users to consume release note data
        *   Filters, etc
        *   https://k8s-relnotes.netlify.com/ -- demo with subset of 1.14 release notes
    *   [https://git.k8s.io/enhancements](https://git.k8s.io/enhancements) for KEPs, enhancement tracking, etc
*   [ 0:00 ] **SIG Updates**
    *   Info for SIG Leads
        *   SIG Leads, check out this set of [recommended topics to cover](https://github.com/kubernetes/community/blob/master/events/community-meeting.md#sig-updates) during your update
        *   [Slide template if you need it](https://docs.google.com/presentation/d/1-nTvKCiqu9UvFYUeM6p6RIqHS5-H-u3_x-V4xj_eIWo/edit#slide=id.g401c104a3c_0_0)
        *   Please also check the [SIG Update Schedule](https://docs.google.com/spreadsheets/d/1adztrJ05mQ_cjatYSnvyiy85KjuI6-GuXsRsP-T2R3k)!
    *   SIG Cloud Provider [Chris Hoge @hogepodge] (confirmed)
        *   [Link to slides](https://docs.google.com/presentation/d/18FlGrOM1VIPZw2fYy1_t4Z6zGXBeOyWwJvH1dEHWOjs/edit#slide=id.g401c104a3c_0_0)
        *   Work on stuff in k8s that’s common across all cloud providers (e.g. loading of cloud providers)
        *   Current biggest area of work is cloud provider extraction
            *   Want to have out-of-tree cloud providers instead of having to have everything baked in
            *   Several cloud providers are integrated into k8s codebase, need to extract them
                *   Step 1: have interface for providers to interact with
                    *   Done: Cloud Controller Manager is complete
                *   Step 2: unwind built-in cloud providers
                    *   many in-tree cloud providers are deeply integrated
                    *   Moving those to staging
                    *   Don’t want to break existing users, but still want to start breaking out of codebase
        *   New areas of work
            *   Restructuring cloud provider SIGs into SIG Cloud Provider subprojects
                *   [Proposal](https://docs.google.com/document/d/16IAOspFVbGWd3RZg6h7uE3WXlRuvWDNnjBPxwITAda4/edit#heading=h.ooec3l4zsr39)
                *   Will still have quarterly updates for cloud providers, still have event time at KubeCon, etc
                *   Plan to have full implementation for KubeCon San Diego (KubeCon NA 2019)
            *   Replacing SSH tunnel system with API server network proxy
            *   Out-of-tree image credential providers
            *   Better docs
            *   HA clusters with cloud controller manager
                *   e.g. leader election is tricky
        *   **In-tree cloud providers going away by December 2019 -- start using the external providers**
            *   Ovirt, cloud stack, and photon cloud providers are deprecated and will be removed
        *   To help: see issues on [https://git.k8s.io/cloud-provider](https://git.k8s.io/cloud-provider)
            *   Links to Slack and Mailing List in slides
    *   SIG API Machinery  [Daniel Smith @lavalamp] (confirmed)
        *   [Link to slides](https://docs.google.com/presentation/d/1MoY1MybEqk7EDmdgdpk3z5KmgmDrC95x7FzE7FbrfUI/edit)
        *   Current work
            *   Server-side apply is in alpha \o/
                *   Demos in SIG API Machinery meetings
            *   CRD schemas now published into OpenAPI (in alpha)
            *   Storage migration work progressing (updating existing objects in etcd to new schema on upgrade)
        *   Upcoming plans
            *   Move extensibility features to GA by end-of-year (CRDs, webhooks, etc)
            *   Apply to beta in 1.15
            *   KEP for better handling of union types (e.g. VolumeSource)
            *   API server traffic classification/proxying support (can’t talk to etcd if you’re trying to talk to a webhook, for example)
                *   KEP posted in enhancements repo
            *   Better server-side rate limiting
                *   KEP posted
        *   **Finalizing CRD, webhook plans soon, provide feedback soon if you have it**
        *   Lots of subprojects, see [slides](https://docs.google.com/presentation/d/1MoY1MybEqk7EDmdgdpk3z5KmgmDrC95x7FzE7FbrfUI/edit) or README above for details
        *   How to help
            *   Trying bug screen meetings, join mailing list if interested
            *   [https://github.com/kubernetes/community/tree/master/sig-api-machinery](https://github.com/kubernetes/community/tree/master/sig-api-machinery) for more details
*   [ 0:00 ] **📣Announcements 📣**
    *   **[Office Hours next week! ](https://github.com/kubernetes/community/blob/master/events/office-hours.md)**Next Wednesday! Ping @castrojo if you want to get involved. 
    *   [Windows containers in Kubernetes Poll](https://pollev.com/michaelmicha980)
        *   SIG Windows wants feedback on Windows use cases in Kubernetes from users
    *   Cluster API now has a discuss.k8s.io [category for discussions](https://discuss.kubernetes.io/c/contributors/cluster-api) if you want to join in. 

        **👏 **Shoutouts this week (Check in #shoutouts on slack) **👏**

    *   vllry - shoutout to @a_sykim for helping me get a kube-proxy bugfix out the door!
    *   jdetiber - shoutout to @castrojo for getting us setup with a Cluster API discourse topic in no time at all
    *   jdetiber - shoutout to @Katharine for helping out with the recent boskos deployments we've needed for wiring up automated e2e tests for the Cluster API subproject
    *   Top 10 Stackoverflow answerers in the Kubernetes Tag for the last week:
        *   Frank Yucheng Gu, Eduardo Baitello, Rico, cookiedough, Janos Lenart, P Ekambaram, Harsh Manvar, 4c74356b41, A_Suh, Leandro Donizetti Soares
        *   Thanks for helping out! 


## April  4, 2019 - ([recording](https://youtu.be/hD-CCAuXu6Y))



*   **Moderators**:  Vallery Lancey [Lyft / k8s contributor]
*   **Note Taker**: Jorge Castro [SIG Contribex]
*   **Demo **-- k3s  [Darren Shepherd]
    *   [Link to slides](https://docs.google.com/presentation/d/1ZxvuT0kBeQSa9m0WETtxzJUzgJLqZzUlmPhsd0bazlc/edit?usp=sharing)
    *   [Link to repositories](https://github.com/rancher/k3s)
*   **Demo **-- BotKube  [Sanket Sudake([sanket@infracloud.io](mailto:sanket@infracloud.io)) & Prasad ([prasad@infracloud.io](mailto:prasad@infracloud.io))]  (confirmed) 
    *   [Link to slides](https://docs.google.com/presentation/d/1ZxvuT0kBeQSa9m0WETtxzJUzgJLqZzUlmPhsd0bazlc/edit?usp=sharing)
    *   [https://github.com/infracloudio/botkube](https://github.com/infracloudio/botkube)
*   [ 0:00 ]** Release Updates**
    *   Current Release Development Cycle  [Claire Laurence - Release Manager]
    *   Patch Release Updates
        *   1.14.1 cherry pick deadline Friday Apr. 5; target release Monday Apr. 8 
        *   
    *   [ 0:00 ] **SIG Updates**
    *   SIG Leads, check out this set of [recommended topics to cover](https://github.com/kubernetes/community/blob/master/events/community-meeting.md#sig-updates) during your update
    *   [Slide template if you need it](https://docs.google.com/presentation/d/1-nTvKCiqu9UvFYUeM6p6RIqHS5-H-u3_x-V4xj_eIWo/edit#slide=id.g401c104a3c_0_0)
    *   Please also check the [SIG Update Schedule](https://docs.google.com/spreadsheets/d/1adztrJ05mQ_cjatYSnvyiy85KjuI6-GuXsRsP-T2R3k)!
    *   SIG Node [Derek Carr] 
        *   [Link to slides](https://docs.google.com/presentation/d/1J9B17qqSqDdajJGB4DavQF0_sv4MWbMUaJj7UnYihw8/edit#slide=id.g401c104a3c_0_0)
    *   SIG CLI [Sean Sullivan]
        *   Link to [[SIG CLI Slides](https://docs.google.com/presentation/d/14078OyhMJJdMKxK7TMoAkkvE39RkR4_Y9_RJiFo9MEU/edit#slide=id.g4045119028_2_268)] 
    *   [SIG network [Bowei Du] ](https://docs.google.com/presentation/d/1Okf6ol-_Tzr9mtQ105Mfjf9lsUAw_JtlBuUhkv7JlQc/edit?usp=sharing)
        *   Link to slides
*   [ 0:00 ] **📣Announcements 📣**
    *   **No major announcements this week**
        *   SIG Chairs/TLs, remember to check your inboxes and #chairs-and-techleads for Paris’ newsletter

        **👏 **Shoutouts this week (Check in #shoutouts on slack) **👏**

    *   nikhita - shoutout to @rlenferink for consistently stepping up to review PRs in kubernetes/org and other contribex repos :)
    *   Jdetiber - shoutout to @justinsb for helping us get the initial v1alpha1 release of cluster-api out the door and the image published!
    *   Vllry - Huge shoutout to @liggitt for the amount of contributor questions he’s been answering… feels like he’s explaining the codebase and community everywhere I go.
    *   Jdetiber - shoutout to @leah for the great work around enumerating and documenting use cases for the Cluster API subproject
    *   Top 10 Stackoverflow users in the Kubernetes Tag for March:
        *   Jexrael, Dmide, Gordon Linoff, Wiktor Stribizew, Martijn Pieters, Wen-Ben, TJ Crowder, and akrun
        *   Thanks for helping out! 


## March 28, 2019 - Release Retrospective for 1.14 ([recording](https://youtu.be/mPNpcJPZuXw))



*   Moderators: Jaice Singer DuMars and Aaron Crickenberger
*   See Retrospective Doc: [https://docs.google.com/document/d/1he2axf3adOIk3gA3vxFAewejtE2tm3Wl1NA1p-ooXpo/edit#](https://docs.google.com/document/d/1he2axf3adOIk3gA3vxFAewejtE2tm3Wl1NA1p-ooXpo/edit#)
*   Normal Community Meeting next week!


## March 21, 2019 - ([recording](https://www.youtube.com/watch?v=Aqp6Rk1J5Jw))



*   **Moderators**:  Jonas Rosland [SIG Contributor Experience]
*   **Note Taker**: First Last [Company/SIG]
*   [ 0:00 ]**  Demo **-- Sonobuoy: End-to-end K8s conformance testing [John Schnake, jschnake@vmware.com] (confirmed)
    *   Link to slides
    *   [https://github.com/heptio/sonobuoy](https://github.com/heptio/sonobuoy)
    *   Diagnostic Tool to understand state of k8s cluster
        *   Uses plugins and conformance tests
        *   Certified Kubernetes is achievable with Sonobuoy
    *   Plugins can be written as run once or run as DaemonSets
    *   Sonobuoy run --mode quick
    *   #sonobuoy Slack
*   [ 0:00 ]** Release Updates**
    *   Current Release Development Cycle  [Aaron Crickenberger - Release Manager]
        *   We are at Week 11 - Code Thaw ([minutes](https://bit.ly/k8s114-minutes)) ([videos](https://www.youtube.com/playlist?list=PL69nYSiGNLP3QKkOsDsO6A0Y1rhgP84iZ&disable_polymer=true))
            *   **Cherry-pick deadline is EOD PT today**
            *   **Go/no-go decision is 8am PT Monday March 25**
            *   Link to [v1.14.0 Known Issues](https://github.com/kubernetes/kubernetes/issues/74425) to help our release notes authors
            *   Come talk to us in #sig-release
            *   [Our CI Signal report for this week](https://groups.google.com/forum/#!topic/kubernetes-dev/G9tkPYOuXcM)
        *   We are taking over the community meeting next week for our retro
            *   Jaice Singer DuMars will be moderator
            *   Anyone may add questions, opinions, celebrations etc. to [https://bit.ly/k8s114-retro](https://bit.ly/k8s114-retro) 
            *   Show up here next week [to speak about them](https://i.imgur.com/Gz1J1LA.jpg)
        *   Start thinking about 1.15
            *   [We chewed through the PR backlog in under 24 hours](http://velodrome.k8s.io/dashboard/db/monitoring?orgId=1)
            *   [We have a 1.15 release team](https://github.com/kubernetes/sig-release/blob/master/releases/release-1.15/release_team.md) (shadow selection TBD)
            *   [We have a v1.15 milestone in kubernetes/enhancements](https://github.com/kubernetes/enhancements/milestone/16)
    *   Patch Release Updates
        *   v1.14.1 depends on need, for now let’s assume Monday April 8th
        *   [v1.13.5 cut planned for Monday March 25](https://groups.google.com/forum/#!topic/kubernetes-dev/ldJFYzwOMKo)
        *   [v1.12.7 cut planned for Monday March 25](https://groups.google.com/forum/#!topic/kubernetes-dev/v_fQpn5_EDA)
        *   [v1.11.9 cut planned for Monday March 25](https://groups.google.com/forum/#!topic/kubernetes-dev/miKocPuZ_Qg)
*   [ 0:00 ] **Contributor Tip of the Week **[Jonas Rosland] 
    *   ~~If you have an issue with the CNCF CLA (changing jobs, loss of account access etc) contact the LF helpdesk at: [helpdesk@rt.linuxfoundation.org](mailto:helpdesk@rt.linuxfoundation.org)~~
        *   ~~Thank you Bob Killen (@mrbobbytables) !~~
    *   UPD. The Linux Foundation HelpDesk email is not valid anymore. Instead, please visit [https://support.linuxfoundation.org/](https://support.linuxfoundation.org/).  
*   [ 0:00 ] **SIG Updates**
    *   SIG Leads, check out this set of [recommended topics to cover](https://github.com/kubernetes/community/blob/master/events/community-meeting.md#sig-updates) during your update
    *   [Slide template if you need it](https://docs.google.com/presentation/d/1-nTvKCiqu9UvFYUeM6p6RIqHS5-H-u3_x-V4xj_eIWo/edit#slide=id.g401c104a3c_0_0)
    *   Please also check the [SIG Update Schedule](https://docs.google.com/spreadsheets/d/1adztrJ05mQ_cjatYSnvyiy85KjuI6-GuXsRsP-T2R3k)!
    *   SIG UI [Jeffrey Sica] (confirmed)
        *   [https://docs.google.com/presentation/d/1wL6GQY9tUCbYRoJB9aXs-D437K_Rz4nHR2nKDAkzcuY/edit?usp=sharing](https://docs.google.com/presentation/d/1wL6GQY9tUCbYRoJB9aXs-D437K_Rz4nHR2nKDAkzcuY/edit?usp=sharing)
        *    
    *   SIG Apps [Adnan Abdulhussein] (confirmed)
        *   [https://docs.google.com/presentation/d/1x0RvC_Rm5vOUcGKFXSGfW--EpPdd9bmIrvSYDK0CHbg/edit](https://docs.google.com/presentation/d/1x0RvC_Rm5vOUcGKFXSGfW--EpPdd9bmIrvSYDK0CHbg/edit)
    *   SIG Windows [Michael Michael]  (confirmed)
        *   [https://docs.google.com/presentation/d/1P9ygrb5i5zcl84ncZdVjgpB4EB7F8rCxjitnrdfjAfc/edit?usp=sharing](https://docs.google.com/presentation/d/1P9ygrb5i5zcl84ncZdVjgpB4EB7F8rCxjitnrdfjAfc/edit?usp=sharing)
*   [ 0:00 ] **📣Announcements 📣**
    *   Contributor Summit for Barcelona is live!
        *   New contributor workshops, and SIG F2F
            *   101 workshop for newcomers to K8s and open source
            *   201 workshop for newcomers to K8s
        *   Check out the event site here and register: [https://events.linuxfoundation.org/events/contributor-summit-europe-2019/](https://events.linuxfoundation.org/events/contributor-summit-europe-2019/)
        *   Also read up on everything around previous and future Contributor Summits in the latest blog post by Paris Pittman and Jonas Rosland: [https://kubernetes.io/blog/2019/03/20/a-look-back-and-whats-in-store-for-kubernetes-contributor-summits/](https://kubernetes.io/blog/2019/03/20/a-look-back-and-whats-in-store-for-kubernetes-contributor-summits/)
    *   castrojo (Jorge Castro): Shoutout to all of you publishing your meetings, we've crossed 10,000k subscribers on the youtube channel! (10,602 to be exact) - let this also be a reminder to catch up on your uploads if you're behind. :smile:
        *   190,961 views, 881,553 minutes of watch time, 4,187 videos! 

        **👏 **Shoutouts this week (Check in #shoutouts on slack) **👏**

    *   akutz (Andrew Kutz): #shoutout to @stevekuznetsov and @krzyzacy (Sen Lu) for helping me the last couple of days get things straightened out for some upcoming, foundational changes to the way VMware is running e2e on our CCM - Oh, and a shout-out to @bentheelder (Benjamin Elder) since I straight-up stole this bit of a Dockerfile from him
    *   justaugustus (Stephen Augustus): Shout-out to @bentheelder for my new profile pic!
    *   zacharysarah (Zach Corleissen): Shoutout to @paris (Paris Pittman) for her Need to Know emails for K8s chairs. I learn so much! 
    *   bentheelder - And helping review #kind! Thanks for all the hard work @neolit123!
        *   timothysc - might want to poke @neolit123 on reviews, b/c he's done so much test infra this last cycle he deserves some sort of metal.
    *   spiffxp (Aaron Crickenberger): shoutout to @vllry (Vallery Lancey) for a bot that tries to auto-label issues with the right `sig/` label: https://github.com/athenabot/k8s-issues (I noticed it labeling issues while triaging CI signal issues, really neat even if a work in progress!)
    *   jimangel: Shoutout to @m2 (Michael Michael) and @craiglpeters for writing a massive amount of Windows documentation needed for the v1.14 release! Also, thanks for involving the #sig-docs crew very early in the process! Please tag anyone I missed.
    *   tpepper (Tim Pepper): Shoutouts to @spiffxp and @maria (Maria Ntalla) and @jberkus (Josh Berkus) for work across recent months toward a simplified/cleaner https://testgrid.k8s.io/sig-release-master-blocking config.  And to all the SIGs who’ve responded with improvements to known flaky test cases.  It is hard to state the importance of stable/green CI…. :green_heart: :testgrid:
    *   zacharysarah: Shout to @rui for replacing a ton of absolute links in the docs with relative paths. A ton of work for largely invisible results that make localizations easier.
    *   spiffxp: Shout to @nikhita (Nikhita Raghunath) for making sigs.yaml (and the generator app) the grand unified theory of kubernetes community, representing SIGs, WGs, and now thanks to her efforts, User Groups and Committees.  Organizational stuff can be boring but it’s important, so really great stuff to offload to machines!


## March 14, 2019- ([recording](https://youtu.be/qzKntMfd7IM))



*   **Moderators**:  Jorge Castro [SIG Contributor Experience]
*   **Note Taker**: First Last [Company/SIG]
*   [ 0:00 ]**  Demo **-- ~~[dmesser@redhat.com](mailto:dmesser@redhat.com) - Automated Day 2 Operations on Kubernetes using Operators Demo Title [Daniel Messer, [dmesser@redhat.com](mailto:dmesser@redhat.com)] (confirmed) ~~
    *   Didn’t happen - timezone snafu, Jorge will reschedule them as soon as possible.
*   [ 0:00 ]** Release Updates**
    *   Current Release Development Cycle  [Aaron Crickenberger - Release Manager]
        *   We are at Week 10 - Code Freeze ([minutes](https://bit.ly/k8s114-minutes)) ([videos](https://www.youtube.com/playlist?list=PL69nYSiGNLP3QKkOsDsO6A0Y1rhgP84iZ&disable_polymer=true))
            *   It’s time to come talk to us if you feel a PR needs to land (use the exception process)
            *   Daily burndown meetings
            *   #sig-release
            *   kubernetes-release-team@
        *   Upcoming milestones:
            *   **[Code Thaw](https://github.com/kubernetes/sig-release/tree/master/releases/release-1.14#code-thaw) - Tuesday March 19** (week 11)
            *   **Cherry Pick Deadline - Thursday March 21** (week 11)
        *   30 Enhancements ([https://bit.ly/k8s114-enhancements](https://bit.ly/k8s114-enhancements))
        *   CI Signal
            *   [http://bit.ly/k8s114-cisignal](http://bit.ly/k8s114-cisignal)
            *   [https://github.com/orgs/kubernetes/projects/11](https://github.com/orgs/kubernetes/projects/11)
            *   1 failing job in master-blocking, 4 flaking
            *   2 failing jobs in 1.14-blocking, 5 flaking
            *   Let’s talk golang 1.12 for a moment
                *   [Migrate to go 1.12.1](https://github.com/kubernetes/kubernetes/issues/75372)?
                *   … or revert to 1.11.5?
                *   Maybe end up pushing back Code Thaw date to support this, watch kubernetes-dev@ for notification
            *   (Top) Monitoring PRs:
                *   [https://github.com/kubernetes/kubernetes/pull/75366](https://github.com/kubernetes/kubernetes/pull/75366) (expected to stabilize subpath unmount flakes in upgrade tests [#75196](https://github.com/kubernetes/kubernetes/issues/75196); master-upgrade, awaiting /approve, /lgtm)
                *   [https://github.com/kubernetes/kubernetes/pull/75341](https://github.com/kubernetes/kubernetes/pull/75341) (expected to stabilize affinity flakes in upgrade tests [#71423](https://github.com/kubernetes/kubernetes/issues/71423) [#72493](https://github.com/kubernetes/kubernetes/issues/72493); master-upgrade, awaiting release-note)
                *   [https://github.com/kubernetes/ingress-gce/pull/678](https://github.com/kubernetes/ingress-gce/pull/678) (expected to resolve ingress failing tests in master-blocking [#75186](https://github.com/kubernetes/kubernetes/issues/75186), 1.14-blocking, awaiting /lgtm)
        *   Issue Triage
            *   [6 milestone v1.14 PR’s](https://github.com/kubernetes/kubernetes/pulls?utf8=%E2%9C%93&q=repo%3Akubernetes%2Fkubernetes+is%3Aopen+milestone%3Av1.14+is%3Apr+sort%3Aupdated-asc)
    *   Patch Release Updates
        *   x.x
        *   y.x
*   [ 0:00 ] **SIG Updates**
    *   SIG Leads, check out this set of [recommended topics to cover](https://github.com/kubernetes/community/blob/master/events/community-meeting.md#sig-updates) during your update
    *   [Slide template if you need it](https://docs.google.com/presentation/d/1-nTvKCiqu9UvFYUeM6p6RIqHS5-H-u3_x-V4xj_eIWo/edit#slide=id.g401c104a3c_0_0)
    *   Please also check the [SIG Update Schedule](https://docs.google.com/spreadsheets/d/1adztrJ05mQ_cjatYSnvyiy85KjuI6-GuXsRsP-T2R3k)!
    *   SIG Architecture [Matt Farina] (confirmed) [Deck link](https://docs.google.com/presentation/d/1vWmzyKyw7I5n1bTPNwurfmIyY0EjxEjGAmN_Frb_RFU/edit?usp=sharing)
        *   Making changes to better scale and avoid burnout
            *   Focus on:
                *   documenting guidance as opposed to one-off decisions
                *   Move discussions to mailing lists to better include others
                *   more on delegating to OWNERS
                *   Ensuring subprojects cultivate new leaders, make sure subprojects are staffed
            *   Done so far
                *   KEP process → SIG PM
                *   _Meetings every other week now_
                *   API Review shadowing (teach new people to be API reviewers
        *   Other things done:
            *   Add guidance that everything should be tested
            *   KEPs for 1.14+ should now have upgrade/downgrade/test plan
            *   [Documented Kubernetes Scope](https://github.com/kubernetes/community/blob/master/contributors/design-proposals/architecture/scope.md)
            *   Windows nodes GA
        *   How to help
            *   Arch & API Governance subproject: look at process, sign up for shadowing
            *   Conformance Definition: come talk to them to onboard
            *   Code Organization: need people to help kickstart untangling kubernetes/kubernetes, figure out dependency handling (currently mostly unstaffed, needs people to start the effort)
            *   [https://git.k8s.io/community/sig-architecture/README.md](https://git.k8s.io/community/sig-architecture/README.md)
        *   Questions/Additions
            *   SIG arch is always welcomes people who want to come in and get work done
            *   Goal of changes is to avoid SIG Arch being the single-point-of-failure for technical decisions, and instead be a last resort
    *   SIG VMware  [Steven Wong] (confirmed - ~3 min) [Deck link](https://docs.google.com/presentation/d/1vjjPBpAjODXrqK79WKummvhH1P6xztfscBX570S86lI/edit?usp=sharing)
        *   Last cycle
            *   Out-of-tree cloud provider
            *   CSI provider improvements
            *   Cluster API provider
            *   Minikube for Fusion/Workstation
        *   Upcoming cycle
            *   Stabilize cloud provider, CSI provider
            *   Cluster API management cluster pattern
        *   How to help
            *   Help wanted bugs (see slides)
            *   [https://git.k8s.io/kubernetes/community/sig-vmware/README.md](https://git.k8s.io/kubernetes/community/sig-vmware/README.md)
    *   SIG Multicluster [Irfan ur Redhman] (confirmed) 
        *   Time zone problem, Jorge to reschedule asap.
*   [ 0:00 ] **📣Announcements 📣**
    *   New Contributor Summit site is live at [https://events.linuxfoundation.org/events/contributor-summit-europe-2019/](https://events.linuxfoundation.org/events/contributor-summit-europe-2019/), and registration will be live be EOD March 14th
    *   Slack signup is back!
    *   1.15 Release Team
        *   @claurence is 1.15 lead
        *   Issue for rest of roles is yet to be created
    *   **PSA: we have a new channel on slack #pr-reviews to help people find reviewers. Help get eyeballs on PRs.** 
    *   Aaron has a last minute announcement

        **👏 **Shoutouts this week

*   coderanger - Shoutout to @Katharine for kicking butt on Slack automation to help out the admin team!
    *   (e.g. more actions → report for bad messages)
*   mrbobbytables - Just seconding @coderanger seriously big shoutout to @Katharine  for the stuff shes doing to make Slack a better place for all of us :heart: :heart: :heart:
*   spiffxp - Big shouts to @neolit123 for his investigative work on what appears to be a golang 1.12 bug. And @liggitt @justinsb @dims for the continued assists. And the golang team for trying to get us go1.12.1 in time
*   Shout-out to @marpaia for setting up and running the EU friendly release team meetings, stepping in so much whenever Aaron isn't available, and keeping on top of all the things :pray:


## March 7, 2019 - ([recording](https://youtu.be/GeB50xG-gmc))



*   **Moderators**:  Chris Short [SIG-ContribEx]
*   **Note Taker**: Bob Killen - University of Michigan
*   [ 0:00 ]**  Demo **-- Argo CD — Enterprise scale open source GitOps solution to deploy 100s of apps in prod - [Jesse_Suen@intuit.com](mailto:Jesse_Suen@intuit.com) &  [Alexander_Matyushentsev@intuit.com](mailto:Alexander_Matyushentsev@intuit.com) (confirmed)
    *   Slides
    *   [https://github.com/argoproj/argo-cd](https://github.com/argoproj/argo-cd)
    *   Collection of controllers and tools for workflow processing.
    *   Well known for their batch job engine
    *   Supports SSO via dex
    *   Has “GitOps” style workflow capability by auto-syncing with git repos.
    *   Has built-in health checks for native kubernetes objects for use with deploying.
        *   Can add custom checks via lua checks
        *   Can rollback in the event of a degraded deployment
    *   Has pre and post sync hooks that can be used to inject custom logic
        *   Useful for db migrations etc
    *   Emits events for auditing along with full prometheus metrics
*   [ 0:00 ]** Release Updates**
    *   Current Release Development Cycle  [Aaron Crickenberger - Release Manager]
        *   We are at Week 9 - Code Freeze EOD PST Today ([minutes](https://bit.ly/k8s114-minutes)) ([videos](https://www.youtube.com/playlist?list=PL69nYSiGNLP3QKkOsDsO6A0Y1rhgP84iZ&disable_polymer=true))
            *   Still in Burndown 
            *   Daily Burndown meetings begin next week
            *   I made [slides](https://docs.google.com/presentation/d/1Ex1FfgC4e6gFoMzDfi_oqyDF-OFDhRUZYdpzCmv5qFU/edit#slide=id.p)
            *   @jeefy and @Katharine [made something better](https://twitter.com/spiffxp/status/1103717028232290305)
        *   Upcoming milestones:
            *   **[Code Freeze](https://github.com/kubernetes/sig-release/tree/master/releases/release-1.14#code-freeze) - Thursday March 7** (week 9)
                *   [https://twitter.com/spiffxp/status/1103717028232290305](https://twitter.com/spiffxp/status/1103717028232290305)
            *   [Docs Ready For Review](https://github.com/kubernetes/website/pulls?utf8=%E2%9C%93&q=is%3Apr+is%3Aopen+base%3Adev-1.14+label%3Ado-not-merge%2Fwork-in-progress) - Monday March 11
                *   9 WIP PRs in flight 
            *   [Code Thaw](https://github.com/kubernetes/sig-release/tree/master/releases/release-1.14#code-thaw) - Tuesday March 19 (week 11)
        *   Enhancements
            *   [https://bit.ly/k8s114-enhancements](https://bit.ly/k8s114-enhancements)
            *   Aggressive pruning has started
            *   16 / 31 enhancements at risk
        *   CI Signal
            *   [http://bit.ly/k8s114-cisignal](http://bit.ly/k8s114-cisignal)
            *   [https://github.com/orgs/kubernetes/projects/11](https://github.com/orgs/kubernetes/projects/11)
            *   Large number of flakey tests
            *   Prioritizing release blocking jobs
        *   Issue Triage
            *   [45 milestone v1.14 PR’s](https://github.com/kubernetes/kubernetes/pulls?utf8=%E2%9C%93&q=repo%3Akubernetes%2Fkubernetes+is%3Aopen+milestone%3Av1.14+is%3Apr+sort%3Aupdated-asc)
            *   [8 size/XXL PR’s](https://github.com/kubernetes/kubernetes/pulls?utf8=%E2%9C%93&q=repo%3Akubernetes%2Fkubernetes+is%3Aopen+milestone%3Av1.14+is%3Apr+sort%3Aupdated-asc+label%3Asize%2FXXL)
    *   Patch Release Updates
        *   x.x
        *   y.x
*   [ 0:00 ] **Contributor Tip of the Week **[Jorge Castro] 
    *   [Contributor Playground](https://github.com/kubernetes-sigs/contributor-playground/pull/229#issuecomment-466711312) is now live! 
        *   Home for new contributors to learn how to use bots and practice with PRs.
    *   Nice work @gsaenger!
*   [ 0:00 ] **SIG Updates**
    *   SIG Leads, check out this set of [recommended topics to cover](https://github.com/kubernetes/community/blob/master/events/community-meeting.md#sig-updates) during your update
    *   [Slide template if you need it](https://docs.google.com/presentation/d/1-nTvKCiqu9UvFYUeM6p6RIqHS5-H-u3_x-V4xj_eIWo/edit#slide=id.g401c104a3c_0_0)
    *   Please also check the [SIG Update Schedule](https://docs.google.com/spreadsheets/d/1adztrJ05mQ_cjatYSnvyiy85KjuI6-GuXsRsP-T2R3k)!
    *   SIG IBM Cloud [Sahdev Zala and Richard Theis] (confirmed)
        *   [Slides](https://goo.gl/Hqgnae )
        *   Last Cycle
            *   Charter merged
            *   Started internal process to open source cloud provider code
            *   [Report on SIG-IBM Cloud @ KubeCon](https://developer.ibm.com/blogs/recap-kubecon-north-america-2018-sig-ibm-cloud-activities/)
        *   Upcoming Cycle
            *   Move internal cloud-provider repo public
                *   Refactor to use cloud-controller manager design
                *   Working on internal build process changes
    *   SIG Service Catalog [Jonathan Berkhahn] (confirmed)
        *   [Slides](https://docs.google.com/presentation/d/1qvxwtIBVm6ZH_EBxCqL7wdPCEliFDk5JHOUYwouweM4/edit?usp=sharing)
        *   Last Cycle
            *   Continued work on namespaced resources
            *   Began work on transitioning to CRDs
        *   Upcoming Cycle
            *   Namespaced resources
            *   Move to GA
                *   CRDs
                *   Bug fixes
            *   Finish moving tests to prow
            *   Recruiting new contributors and maintainers
*   [ 0:00 ] **📣Announcements 📣**
    *   Working Group LTS has a survey: [https://www.surveymonkey.com/r/kubernetes-support-survey-2019](https://www.surveymonkey.com/r/kubernetes-support-survey-2019) 
    *   Plan to stop serving deprecated extensions/v1beta1, apps/v1beta1, apps/v1beta2 APIs in v1.16+ (tracking issue [#43214](https://github.com/kubernetes/kubernetes/issues/43214), v1.16 PR [#70672](https://github.com/kubernetes/kubernetes/pull/70672)) [liggitt]
        *   `daemonsets`, `deployments`, `replicasets` resources: use `apps/v1`
        *   `networkpolicies` resources: use `networking.k8s.io/v1`
        *   `podsecuritypolicies` resources: use `policy/v1beta1`
    *   Need a few more Slack moderators! [Especially for EU and APAC](https://github.com/kubernetes/community/blob/master/communication/moderation.md). [jorge]

    **👏 Shoutouts this week (Check in #shoutouts on slack) 👏**

*   Aaron Crickenberger: shouts to @oomichi for his continued review of kubernetes/kubernetes test/e2e PR’s, he’s helped land over 30 PR’s this release cycle (including some conformance tests)
*   Lachlan Evenson: shouts to @bentheelder @munnerz and the kind team for all the great work that’s gone into the kind tool and docs. It’s radically changed my inner loop for testing code changes in Kubernetes. Thanks!
*   Lachlan Evenson: shouts to @claurence for being an awesome enhancements lead on the 1.14 release team. She’s taken the time to coach all the shadows and has been diligently grooming the 33 features in the hopper for 1.14. Cheers from your fellow enhancements shadows!
*   Paris Pittman: Thanks to our upstream mentoring panelists on #meet-our-contributors today!! @mike.splain @carolynvs @dims @directxman12 @a_sykim
*   Dims: Thank you to the #meet-our-contributors hosts always excellent hosts @paris and @jorge!
*   Chris Short: Thank you to @mrbobbytables for stepping in to take notes the past two Kubernetes Community meetings. I really appreciate it! :khanparrot::khanparrot::khanparrot:
*   Nikhita: shoutout to @jeefy and @Katharine for [https://twitter.com/spiffxp/status/1103717028232290305](https://twitter.com/spiffxp/status/1103717028232290305)


## February 28, 2019 ([recording](https://youtu.be/fjZ5l8gZrcw))



*   **Moderators**: Chris Short [SIG ContribEx]
*   **Note Taker**: Bob Killen
*   [ 0:00 ]**  Demo **-- Feb 28 - Kubernetes Policy Controller with OPA/KPC - Dave Strebel [strebeld@gmail.com](mailto:strebeld@gmail.com) (confirmed)
    *   Gatekeeper project and Ku[https://youtu.be/fjZ5l8gZrcw](https://youtu.be/fjZ5l8gZrcw)bernetes policy controller merged to become [Open Policy Agent (OPA) Gatekeeper](https://github.com/open-policy-agent/gatekeeper).
    *   OPA - General purpose policy engine using [declarative policy language (rego)](https://www.openpolicyagent.org/docs/how-do-i-write-policies.html)
    *   Focused on protecting the Kubernetes API
    *   Augments Admission / Authorization and Audit capabilities of Kubernetes.
        *   Adds more granular policies and can mutate requests.
    *   Partially backed by CRDs, working on full configuration through CRD.
    *   Can audit current environment against policies without enforcing policies.
    *   Examples: 
        *   Can enforce policies such as images must be pulled from specific registries.
        *   Can add additional annotations on matching criteria.
    *   [CNCF Blog](https://www.cncf.io/blog/2018/03/29/cncf-to-host-open-policy-agent-opa/)
*   [ 0:00 ]** Release Updates**
    *   Current Release Development Cycle  [Aaron Crickenberger - Release Manager] (confirmed)
        *   We are at Week 8 - Burndown ([minutes](https://bit.ly/k8s114-minutes)) ([videos](https://www.youtube.com/playlist?list=PL69nYSiGNLP3QKkOsDsO6A0Y1rhgP84iZ&disable_polymer=true))
        *   [I made some slides](https://docs.google.com/presentation/d/1EIiUqw28SYAtPshsCm4mMyb0yRtv9zCJRaAfstyzkcs/edit#slide=id.p)
        *   Upcoming milestones:
            *   **Docs [Placeholder](https://groups.google.com/d/msg/kubernetes-dev/rMYqtEF7emk/BgYkVugqBgAJ) PRs Friday March 1 **(week 8)
            *   **[Code Freeze](https://github.com/kubernetes/sig-release/tree/master/releases/release-1.14#code-freeze) Thursday March 7** (week 9)
                *   **Reminder: **No code slush leading to freeze.
        *   Enhancements
            *   [https://bit.ly/k8s114-enhancements](https://bit.ly/k8s114-enhancements)
            *   Are you sure you are ready for your enhancement to land
            *   Do you have tests our CI signal team can use to confirm your enhancement is working
            *   14 enhancements at risk
        *   CI Signal
            *   [http://bit.ly/k8s114-cisignal](http://bit.ly/k8s114-cisignal)
            *   [https://github.com/orgs/kubernetes/projects/11](https://github.com/orgs/kubernetes/projects/11)
            *   Starting to assign priority/critical urgent 
            *   All release-blocking flakes or failures are getting /milestone v1.14 /priority critical/urgent
        *   Release notes should well..be about the release.
            *   Create better documentation for end users for release.
            *   [PoC website](https://github.com/kubernetes/sig-release/issues/529) built by @jeefy 
    *   Patch Release Updates
        *   x.x
        *   y.x
*   [ 0:00 ] **Contributor Tip of the Week**
    *   Jordan Liggitt
    *   API review process
    *   What needs a review: [http://go.k8s.io/api-review#what-apis-need-to-be-reviewed](http://go.k8s.io/api-review#what-apis-need-to-be-reviewed) 
    *   How do I get a review: [http://go.k8s.io/api-review#mechanics](http://go.k8s.io/api-review#mechanics)
    *   Tracked reviews: [https://github.com/orgs/kubernetes/projects/13](https://github.com/orgs/kubernetes/projects/13)
    *   [Sample bot hint comment](https://github.com/kubernetes/kubernetes/pull/74477#issuecomment-466823631)
*   [ 0:00 ] **SIG Updates**
    *   SIG Leads, check out this set of [recommended topics to cover](https://github.com/kubernetes/community/blob/master/events/community-meeting.md#sig-updates) during your update
    *   [Slide template if you need it](https://docs.google.com/presentation/d/1-nTvKCiqu9UvFYUeM6p6RIqHS5-H-u3_x-V4xj_eIWo/edit#slide=id.g401c104a3c_0_0)
    *   Please also check the [SIG Update Schedule](https://docs.google.com/spreadsheets/d/1adztrJ05mQ_cjatYSnvyiy85KjuI6-GuXsRsP-T2R3k)!
    *   SIG Cluster Lifecycle [Tim Allclair and Robert Bailey] (confirmed)
        *   kubeadm v1.14
            *   p0 - working on better test automation
                *   starting to use KinD  as a tool to replace kubernetes anywhere
                *   upgrade testing
            *   p1 - improve HA lifecycle
            *   p2 - improve join action
        *   cluster-api
            *   Plan to release v1alapha1 around the v1.14 release
            *   Added support for cascading deletion
            *   Improve machine object deletion
            *   NEED testers
        *   minikube
            *   v0.34 update released
            *   Working towards a v1 release in March
        *   Kops
            *   upgraded to etcd3
            *   working through current CVE issue
            *   building roadmap for cluster-api
        *   KinD
            *   Offline support
            *   Upgraded to v1.13.3
            *   Goal: Use KinD to replace local cluster-up
        *   PSAs
            *   Component config working group started
            *   Working on addon management
                *   CRD lifecycle
        *   Upcoming planning session.
    *   SIG OpenStack [Chris Hoge](confirmed)
        *   [Slides](https://docs.google.com/presentation/d/152oNn6lURkr3cWp6wVh0FttO3gpER_1wV3YuN6aibmM/edit#slide=id.g288eb86093_1_248) 
        *   Moving in-tree provider from kubernetes/kubernetes
        *   Working on building better deployment tooling
            *   Magnum - integrated K8s deployment
            *   Self service through Kops
            *   Cluster-API implementation for OpenStack cloud and bare metal.
        *   Bare Metal Cluster-API implementation will use ironic with multiple deployment options:
            *   Standalone single tenant
            *   Integrated, multi-tenant with cloud-link services.
        *   Kops OpenStack provider is alpha
        *   Plan to work closer with CNCF regarding CI/CD testing and working closer with sig-testing.
        *   
    *   SIG Auth [Mike Danese](confirmed)
        *   [Slides](https://docs.google.com/presentation/d/1PvuhGIKhgOTiXXcGbxlZG7xGWRDSMR8XpREs836_6wU/edit?usp=sharing)
        *   Working on roll out of better service account tokens
        *   Dynamic auditing with per sink policy
        *   Refining approach to the different policy types in Kubernetes
            *   Dynamic admission ecosystem including a general purpose policy engine (e.g. OPA)
            *   New domain-specific policies: scheduling and images
            *   rethinking PodSecurityPolicies due to usability issues
        *   Improve API Server authentication
            *   dynamic webhooks have become popular, need a better way to authenticate the api server as a client.
            *   Webhooks can accept sensitive data and return sensitive data, need to identify the caller
        *   Organization:
            *   Identified and defined subprojects and TLs.
            *   More proactively engage with subprojects
            *   Absorbed wg-container-identity.
*   [ 0:00 ] **📣Announcements 📣**
    *   We’re in need of more slack moderators, [apply here](https://github.com/kubernetes/community/issues/new/choose) and click the moderator request button.
        *   Must be a k8s org member already
        *   APAC and EU moderators needed the most
    *   Shoutouts
        *   Akutz - Major #shoutout to Katharine Berry (@Katharine) for fixing a UX bug within just a few hours of me mentioning it - [https://kubernetes.slack.com/archives/C09QZ4DQB/p1550863111161400](https://kubernetes.slack.com/archives/C09QZ4DQB/p1550863111161400). Ain’t no service like SIG-Testing service, cause SIG-Testing service don’t stop! 
        *   spiffxp - shoutout to Thomas Runyon (@runyontr) for [https://github.com/kubernetes/kubernetes/pull/72939](https://github.com/kubernetes/kubernetes/pull/72939) allowing us to run eg: `make test-cmd WHAT=deployment` to run just the the deployment cli tests, same sort of thing you can do with `make test` and `make test-integration` 
        *   bentheelder - shoutout to @gsaenger for writing a wonderful new message for the welcome bot! looking forward to seeing this in more places - ([https://github.com/kubernetes-sigs/contributor-playground/pull/229#issuecomment-466711312](https://github.com/kubernetes-sigs/contributor-playground/pull/229#issuecomment-466711312))
        *   bentheelder - Shoutout to @jeefy for the really shiny and useful looking structured Kubernetes release notes viewer demo in today's #sig-release meeting, looking forward to seeing more about this! 
        *   Maria - shoutout to Silvia Moura Pina (@smourapina) for putting together a workflow to enable the CI signal subteam of the release team keep on top of flagging issues from e2e tests and coordinate follow-ups and to Jorge Alarcon (@Jorge) for spotting an opportunity to offer broader transparency to what the CI signal team is working on, suggesting a structure and kicking off implementation (find current version at [https://github.com/orgs/kubernetes/projects/11](https://github.com/orgs/kubernetes/projects/11))
        *   spiffxp - shouts to Josh Berkus (@jberkus) for taking notes during today’s steering committee meeting, our google doc clearly becomes way more difficult to use with so many people looking at it, and the written record is invaluable
        *   spiffxp - shoutout to Katharine Berry (@Katharine) for moving us from gubernator to spyglass for all of our test result viewing needs! [https://git.k8s.io/test-infra/prow/spyglass ](https://git.k8s.io/test-infra/prow/spyglass)
        *   codenrhoden - Shoutout to Michelle Au (@msau42) for her patience and helpful guidance in getting a very large PR merged!
        *   @strebel - @jeefy and @onyiny-ang for all their awesome work on the Release-Notes team for 1.14. Especially @jeefy’s work on the Release-Notes website concept