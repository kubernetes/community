

## February 21, 2019 - ([recording](https://youtu.be/DYmZVxtyCN4))



*   **Moderators**:  Jonas Rosland [SIG ContribEx]
*   **Note Taker**: Jorge Castro [SIG ContribEx]
*   [ 0:00 ]**  Demo **-- Kubernetes in Kubernetes [?ukasz Ole?, [loles@mirantis.com](mailto:loles@mirantis.com)] (confirmed)
    *   [https://docs.google.com/presentation/d/1PsB_fU1IxjS2grRdMdBm_fsNlSgVChWQ6r-HwGyBRNw/edit?usp=sharing](https://docs.google.com/presentation/d/1PsB_fU1IxjS2grRdMdBm_fsNlSgVChWQ6r-HwGyBRNw/edit?usp=sharing) 
    *   [https://github.com/lukaszo/cluster-api-provider-virtlet](https://github.com/lukaszo/cluster-api-provider-virtlet) 
*   [ 0:00 ]** Release Updates**
    *   Current Release Development Cycle  [Aaron Crickenberger - Release Manager]
        *   [We are at Week 7 for v1.14](https://github.com/kubernetes/sig-release/tree/master/releases/release-1.14) - [minutes](https://docs.google.com/document/d/1U3jL8Ucruoq9wyzIgxEdyA51MuOIi_gvecVed1kAli0/edit#heading=h.60ptyogm23dd) - [recording](https://youtu.be/SfaBzKPeaLk)
        *   Everything has a KEP, but are they useful
        *   We have [release-1.14-blocking](https://testgrid.k8s.io/sig-release-1.14-blocking) and [release-1.14-all](https://testgrid.k8s.io/sig-release-1.14-all) testgrid dashboards backed by release-1.14 jobs (thanks @amwat, @dbhanushali, @krzyzacy)
        *   Comments / questions / concerns on how this is going? Add them to [https://bit.ly/k8s114-retro](https://bit.ly/k8s114-retro) 
    *   Upcoming milestones:
        *   **[Burndown](https://github.com/kubernetes/sig-release/tree/master/releases/release-1.14#burndown) Monday February 25** (week 8)
        *   **[Code Freeze](https://github.com/kubernetes/sig-release/tree/master/releases/release-1.14#code-freeze) Thursday March 7** (week 9)
    *   Enhancements
        *   [https://bit.ly/k8s114-enhancements](https://bit.ly/k8s114-enhancements)
    *   CI Signal
        *   [http://bit.ly/k8s114-cisignal](http://bit.ly/k8s114-cisignal)
    *   Patch Release Updates
        *   x.x
        *   y.x
*   [ 0:00 ] **Contributor Tip of the Week **[Katharine Berry] 
    *   [Spyglass](https://github.com/kubernetes/test-infra/tree/master/prow/spyglass), and [the deprecation of Gubernator](https://github.com/kubernetes/test-infra/pull/11302) (to view artifacts! We're not touching the PR dashboard)
    *   You might know Gubernator as? that purple and white thing that shows test failures when you click the "Details" link on job results in PRs
    *   Spyglass is the prow component intended to replace this functionality, it's been around for a few months, but we're finally at enough parity that we'd like to switch over
    *   [kuberentes-dev@ thread](https://groups.google.com/forum/#!topic/kubernetes-dev/bzz5HZ3HRoc)
*   [ 0:00 ] **SIG Updates**
    *   SIG Leads, check out this set of [recommended topics to cover](https://github.com/kubernetes/community/blob/master/events/community-meeting.md#sig-updates) during your update
    *   [Slide template if you need it](https://docs.google.com/presentation/d/1-nTvKCiqu9UvFYUeM6p6RIqHS5-H-u3_x-V4xj_eIWo/edit#slide=id.g401c104a3c_0_0)
    *   Please also check the [SIG Update Schedule](https://docs.google.com/spreadsheets/d/1adztrJ05mQ_cjatYSnvyiy85KjuI6-GuXsRsP-T2R3k)!
    *   SIG Contributor Experience [Paris Pittman] (confirmed)
        *   [Slides](https://docs.google.com/presentation/d/14JKmY0DZ1BiZvy21jCSrxqX8EGRhHuZJWseKW3nQZX4/edit?usp=sharing) (open to kdev and contribex)
    *   SIG AWS [Nishi Davidson] (confirmed)
        *   [Slides](https://docs.google.com/presentation/d/1400kj3WXVZTpTi0dhpmb_9CN3_URV608qgcsey0jQwQ/edit#slide=id.p2)
    *   SIG Scheduling [Bobby Salamat]  (confirmed)
        *   [Slides](https://docs.google.com/presentation/d/13zxbEZPIqPB2AN-nDfY3gXW_ojId8f1yrKFCN_NvLoQ/edit?ts=5c6de6ec#slide=id.p)
*   [ 0:00 ] **?Announcements ?**
    *   From SIG Docs: [https://requestbin.com/](https://requestbin.com/) for live API/webhook testing: We recently launched an updated, hosted version of the tool at RequestBin.com with new features including private bins (with Google / Github authentication), ability to pause the event stream and an improved UI.  Thousands of developers are using the new version today and it is ready for public distribution.
    *   [@kubernetes-maintainers and @kubernetes-release-managers are losing direct write access to kubernetes/enhancements](https://github.com/kubernetes/enhancements/issues/590) [spiffxp]

        **? **Shoutouts this week (Check in #shoutouts on slack) **?**

    *   #kind shoutout from @mauilion on #TGIK right now! Thank you to the maintainers @bentheelder @munnerz
    *   to name a few very helpful hands with #kind @neolit123, @fabrizio.pandini, @Jorge, @TaoBeier, @amwat, @krzyzacy :left-shark:
    *   thanks to @bentheelder and @krzyzacy for helping with the creation of a new "kind" based deployer for kubetest (in test-infra)!!
    *   Shoutout to @bentheelder for sharing good insight on how to publish K8s projects!


## February 14, 2019 (recording)



*   **Moderator**: Jeff Sica, SIG UI
*   **Note Taker: **
*   [ 0:00 ] **Demo**: Kube-service-exporter: A way to bring your own load balancer to balance across multiple clusters, using consul - Guinevere Saenger, GitHub
*   **Release updates:**
    *   **marpaia (confirmed)**
*   **SIG Updates: **
    *   **Docs (Zach - confirmed)**
        *   [Presentation](https://docs.google.com/presentation/d/1cEOPr9rxNAviOjru_7kXaTxDG-8sF_DPiYIpwERL22U/edit#slide=id.g401c104a3c_0_0)
    *   **Storage (Saad - confirmed)**
        *   [Presentation](https://docs.google.com/presentation/d/11xKfUsH_ePscMQMH7HnIKI37cOSc0BWDEQg5fSF7yHw/edit?usp=sharing)
*   **Announcements:**
    *   **Slack update: **we are only manually inviting contributors who need access for now until we hear from Slack. If you are in a SIG and have a member that needs access, Ping in #slack-admins and an admin will DM you for the email. Consumer traffic is being routed to discuss.kubernetes.io.
    *   
    *   **Kubecon Shanghai CFP** ends at 11:59PM PT, February 22, 2019
        *   SIG Deep Dives/Intros due: 
    *   **Shoutouts (see slack #shoutouts) **
        *   nikhita - Figured this makes a good shoutout too! @spiffxp @pwittrock :smile:
        *   dbhanushali (The-Wall)- In no particular order @krzyzacy (seen) @amwat (Amit) @cjwagner (Cole) @ixdy (Jeff) @bentheelder (Been). shoutout for their assistance in test-infra release task automation
        *   spiffxp - Shoutouts to @coderanger @mrbobbytables and @kbarnard10 for putting together and posting https://kubernetes.io/blog/2019/02/11/runc-and-cve-2019-5736/ so quickly
        *   nikhita - shoutout to @mrbobbytables and @justaugustus for handling new member requests in k/org in such a timely manner! It feels like there are at least 4-5 requests every day and they still manage to get to each of them, while simultaneously doing so many other things for our community! :100:
        *   paris - thanks Jeff for hosting this call in my absence at the last minute. True team player! 
        *   mrbobbytables - shoutout to @zacharysarah, sig-docs and everyone involved in kick starting the french translation efforts! @sieben @Aurelien Perrier @lledru @yastij @smana @rbenzair @Jean-Yves Gastaud and the others I don't have slack handles for but heres github: awkif, abuisine, rekcah78 and erickhun


## February 7, 2019 - ([recording](https://youtu.be/F4NTezNKusU))



*   **Moderators**:  Josh Berkus [SIG-Release]
*   **Note Taker**: GDoc outage during meeting, notes are incomplete
*   [ 0:00 ]**  Demo **-- Eclipse Che / CodeReady Workspaces [Mario Loriedo [mloriedo@redhat.com](mailto:mloriedo@redhat.com) Stevan Le Meur [slemeur@redhat.com](mailto:slemeur@redhat.com) ] (confirmed)
    *   [http://bit.ly/intro-che-kubecc](http://bit.ly/intro-che-kubecc)
    *   [https://eclipse.org/che](https://eclipse.org/che)
    *   [https://github.com/eclipse/che](https://github.com/eclipse/che)
*   [ 0:00 ]** Release Updates**
    *   Current Release Development Cycle  [Aaron Crickenberger - Release Manager]
        *   [We are at Week 5 for v1.14](https://github.com/kubernetes/sig-release/tree/master/releases/release-1.14) - [minutes](https://docs.google.com/document/d/1U3jL8Ucruoq9wyzIgxEdyA51MuOIi_gvecVed1kAli0/edit#heading=h.wxzklg8jw1a) - [recording](https://youtu.be/SfaBzKPeaLk)
        *   Upcoming milestones:
            *   **[Burndown](https://github.com/kubernetes/sig-release/tree/master/releases/release-1.14#burndown) Monday February 25** (week 8)
            *   **[Code Freeze](https://github.com/kubernetes/sig-release/tree/master/releases/release-1.14#code-freeze) Thursday March 7** (week 9)
        *   Enhancements
            *   [https://bit.ly/k8s114-enhancements](https://bit.ly/k8s114-enhancements)
            *   [20 alpha](https://github.com/kubernetes/enhancements/issues?utf8=%E2%9C%93&q=is%3Aissue+is%3Aopen+milestone%3Av1.14+sort%3Acreated-asc+label%3Astage%2Falpha), [13 beta](https://github.com/kubernetes/enhancements/issues?utf8=%E2%9C%93&q=is%3Aissue+is%3Aopen+milestone%3Av1.14+sort%3Acreated-asc+label%3Astage%2Fbeta), [6 stable](https://github.com/kubernetes/enhancements/issues?utf8=%E2%9C%93&q=is%3Aissue+is%3Aopen+milestone%3Av1.14+sort%3Acreated-asc+label%3Astage%2Fstable)
            *   Comments / questions / concerns on how this went? Add them to [https://bit.ly/k8s114-retro](https://bit.ly/k8s114-retro) 
        *   CI Signal
            *   [http://bit.ly/k8s114-cisignal](http://bit.ly/k8s114-cisignal)
    *   Patch Release Updates
        *   No patch releases planned this week.
*   [ 0:00 ] **SIG Updates**
    *   SIG Instrumentation [Frederic Branczyk] (confirmed) 
        *   Slides: [https://docs.google.com/presentation/d/1KzYZCf5VUMp65H_vlPwg7AD4UIXWTRtiUK34L1GYOqE/edit?usp=sharing](https://docs.google.com/presentation/d/1KzYZCf5VUMp65H_vlPwg7AD4UIXWTRtiUK34L1GYOqE/edit?usp=sharing)
    *   SIG Testing (postponed due to issues)
    *   SIG PM [Stephen Augustus] (confirmed)
        *   Slides: [https://docs.google.com/presentation/d/1A3xvBhMqjuPu183RFVSBZptm9hz-_Qz24XzZu14Qsoc/edit?usp=sharing](https://docs.google.com/presentation/d/1A3xvBhMqjuPu183RFVSBZptm9hz-_Qz24XzZu14Qsoc/edit?usp=sharing)
*   [ 0:00 ] **?Announcements ?**
    *   SIG Cluster Lifecycle has published a [grooming document](https://github.com/kubernetes/community/blob/master/sig-cluster-lifecycle/grooming.md) - "By providing a consistent set of best practices, it allows developers to work across sub-projects fairly easily. More importantly, it signals direction and priority to developers who may be new to the sub-projects." ? might be useful to other SIGs! 
    *   We had a slack bad actor incident on Sunday Evening and the inviter is down. Slack.k8s.io is pointing to discuss.kubernetes.io. Once we have finished our talks with Slack and moderators have discussed, we will alert everyone as to the status of the inviter. 

        **? **Shoutouts this week (Check in #shoutouts on slack) **?**


        Cblecker: OSS Good Feeling of the Day: Thanks to @eduar's diligent sorting of docs and adding owners files, I'm already seeing PRs being approved and merged by the actual owners of the files, rather than needing top-level approval from folks like me. Delegation and empowering people! \
Paris: shoutout to @coderanger who helped us during a slack spam attack this week and created a 404 redirect to slack.k8s.io to help close the loop for people looking to join this slack for kubernetes support and community.  Also thanks to @thockin and @spiffxp for stepping in to help us on a sunday evening


        Nikhita: shoutout to @paris @jdumars @spiffxp @coderanger @bentheelder for their quick response in dealing with spam attacks!


        ** \
**



## January 31, 2019 ([recording](https://youtu.be/YWa5ATDP9Dk))



*   **Moderators**: Jeffrey Sica [SIG UI]
*   **Note Taker**: 
*   [ 0:00 ]**  Demo**
    *   multicluster-scheduler: Running Argo Workflows Across Multiple Kubernetes Clusters ([adrien@admiralty.io](mailto:adrien@admiralty.io))

        Slides: [https://docs.google.com/presentation/d/12UzVDWmwdyWCMttAWIDld2udri5sILIFWWqZzqfjrX4/edit?usp=sharing](https://docs.google.com/presentation/d/12UzVDWmwdyWCMttAWIDld2udri5sILIFWWqZzqfjrX4/edit?usp=sharing) 

*   [ 0:00 ]** Release Updates**
    *   Current Release Development Cycle  [Aaron Crickenberger - Release Manager]
        *   [We are at Week 4 for v1.14](https://github.com/kubernetes/sig-release/tree/master/releases/release-1.14) - [minutes](https://docs.google.com/document/d/1U3jL8Ucruoq9wyzIgxEdyA51MuOIi_gvecVed1kAli0/edit#heading=h.dhwowocfhooj) - [recording](https://www.youtube.com/watch?v=1GY-VnEHqNY&t=9s&list=PL69nYSiGNLP3QKkOsDsO6A0Y1rhgP84iZ&index=3)
        *   Upcoming milestones:
            *   **[KEPs Implementable](https://github.com/kubernetes/sig-release/tree/master/releases/release-1.14#enhancements-freeze)** **Monday February 4** (week 5)
            *   **[Burndown](https://github.com/kubernetes/sig-release/tree/master/releases/release-1.14#burndown) Monday February 25** (week 8)
            *   **[Code Freeze](https://github.com/kubernetes/sig-release/tree/master/releases/release-1.14#code-freeze) Thursday March 7** (week 9)
        *   Builds
            *   [We cut v1.14.0-alpha.2](https://github.com/kubernetes/sig-release/issues/460)
            *   We will cut by [1.14.0-alpha.3](https://github.com/kubernetes/sig-release/issues/474) Tuesday February 12 (week 6)
        *   Enhancements
            *   [https://bit.ly/k8s114-enhancements](https://bit.ly/k8s114-enhancements)
            *   [Enhancements freeze](https://github.com/kubernetes/sig-release/tree/master/releases/release-1.14#enhancements-freeze) was Tuesday
                *   8 alpha, 8 beta, 9 stable = 25 with issues and implementable KEPs
            *   Son of Enhancements Freeze Pt II First Blood This Time It's Serious is Monday
                *   We would have also accepted "Bring out your KEPs"
                *   [https://kubernetes.slack.com/archives/C2C40FMNF/p1548895977819200](https://kubernetes.slack.com/archives/C2C40FMNF/p1548895977819200)
                *   8 alpha, 3 beta, 0 stable = 11 KEPs that need to get to implementable by Monday
            *   [kubernetes-dev@ e-mail](https://groups.google.com/forum/#!topic/kubernetes-dev/G5X_iQOIKQE)
            *   [repo: kubernetes/enhancements milestone:v1.14](https://github.com/kubernetes/enhancements/milestone/13)
            *   If you have any questions about the KEP process, join us at (HEY AARON FILL THIS OUT) meeting. 
        *   CI Signal
            *   [http://bit.ly/k8s114-cisignal](http://bit.ly/k8s114-cisignal)
            *   We watch [release-master-blocking](https://testgrid.k8s.io/sig-release-master-blocking), [release-master-upgrade](https://testgrid.k8s.io/sig-release-master-upgrade)
            *   We are informed by [release-master-informing](https://testgrid.k8s.io/sig-release-master-informing) (still being iterated on)
            *   FYI: [pull-kubernetes-e2e-kops-aws is optional, non-blocking](https://groups.google.com/forum/#!topic/kubernetes-dev/mvzwnWR3ahg)
            *   Implementing [release-blocking job criteria](https://github.com/kubernetes/sig-release/blob/master/release-blocking-jobs.md#release-blocking-criteria): all release-blocking jobs must have owners
                *   Specify "owner" as "an e-mail address that testgrid can send alerts to if the job fails more than N times in a row"
                *   Tracking issue: [https://github.com/kubernetes/sig-release/issues/441](https://github.com/kubernetes/sig-release/issues/441) 
                *   WIP PR: [https://github.com/kubernetes/test-infra/pull/10870](https://github.com/kubernetes/test-infra/pull/10870)
    *   Patch Release Updates
        *   1.13.3: [aiming](https://groups.google.com/d/topic/kubernetes-dev-announce/mxdNPaiJFHA/discussion) for tomorrow
*   [ 0:00 ] **SIG Updates**
    *   SIG Leads, check out this set of [recommended topics to cover](https://github.com/kubernetes/community/blob/master/events/community-meeting.md#sig-updates) during your update
    *   [Slide template if you need it](https://docs.google.com/presentation/d/1-nTvKCiqu9UvFYUeM6p6RIqHS5-H-u3_x-V4xj_eIWo/edit#slide=id.g401c104a3c_0_0)
    *   Please also check the [SIG Update Schedule](https://docs.google.com/spreadsheets/d/1adztrJ05mQ_cjatYSnvyiy85KjuI6-GuXsRsP-T2R3k)!
    *   SIG-Azure [@strebel confirmed]
    *   SIG-Big Data [???]
    *   SIG-Release [@tpepper confirmed]: [slides link](https://docs.google.com/presentation/d/1B3FVf8B21qBMD1FCKvKx2uS1dNPTG9J9BPRynBpYir8/edit?usp=sharing)
*   [ 0:00 ] **?Announcements ?**
    *   Google Summer of Code [Paris, Nikhita, Ihor]:
        *   The deadline to submit project ideas is February 6 20:00 UTC
        *   More details:
            *   [https://github.com/cncf/soc#project-ideas](https://github.com/cncf/soc#project-ideas) 
            *   [https://groups.google.com/d/msg/kubernetes-dev/S7HsqOkiC5g/DLSXB4oDEQAJ](https://groups.google.com/d/msg/kubernetes-dev/S7HsqOkiC5g/DLSXB4oDEQAJ) 
    *   Register to SIG updates (Intro and Deep Dive sessions) in Barcelona:
        *   The deadline to register to offer Intro and Deep Dive sessions for Barcelona is Friday, February 8.
        *   To register an intro and/or deep dive session for Barcelona, please go to: [https://www.surveymonkey.com/r/J5K9JJR](https://www.surveymonkey.com/r/J5K9JJR) 
    *   SIG/WG Chairs and coleads (Jorge): 
        *   We need to complete our move to host-keys on Zoom.
        *   Here [are the instructions](https://docs.google.com/document/d/1fudC_diqhN2TdclGKnQ4Omu4mwom83kYbZ5uzVRI07w/).
        *   This is the only thing blocking us from having the calendar public on the web! It would be nice to have this back, see [issue #2536](https://github.com/kubernetes/community/issues/2536).
        *   Talk to a Jorge or Paris near you if you have questions/need help. 
    *   Kubecon Shanghai CFP deadline: February 22, 2019
    *   [Meet our contributors next week](https://github.com/kubernetes/community/blob/master/mentoring/meet-our-contributors.md)!
        *   Looking for mentors, please see @paris
        *   Especially if you're new, we want to hear about the beginning of your k8s journey.
    *   **? **Shoutouts this week (Check in #shoutouts on slack) **?**
        *   zacharysarah - Shoutout to @bentheelder for proactively seeking to collaborate on subdomain site design! [https://github.com/kubernetes/org/issues/397](https://github.com/kubernetes/org/issues/397) 
        *   neolit123 - big thank you to @spiffxp for helping drive the PR that removes tracking of placeholder reference docs in the `k/k` repository. [https://github.com/kubernetes/kubernetes/pull/70052](https://github.com/kubernetes/kubernetes/pull/70052)  
        *   spiffxp - shoutouts to @liggitt for making me forever associate "KEP" with [https://www.youtube.com/watch?v=vh3tuL_DVsE](https://www.youtube.com/watch?v=vh3tuL_DVsE) 
        *   spiffxp - shoutouts to @stevekuznetsov for getting PR Number Eleven Thousand in [kubernetes/test-infra](https://github.com/kubernetes/test-infra)! [https://github.com/kubernetes/test-infra/pull/11000](https://github.com/kubernetes/test-infra/pull/11000) 
        *   akutz - I wanted to #shoutout @stevekuznetsov. He always takes the time to fill in his issues on the #sig-testing weekly call agenda ahead of time and with great detail and links. This makes taking notes in real time much easier when it's his issues. Thank you Steve!
        *   paris - @mrbobbytables for knocking out a ton of "glue" work in our communication documentation to make our processes more complete and transparent in `kubernetes/community/communication` and closing several issues around them! thanks so much bob!
        *   spiffxp - shoutout to @liggitt for suggesting the KEP extension to release freeze and providing language around it [https://github.com/kubernetes/sig-release/pull/482](https://github.com/kubernetes/sig-release/pull/482) 


## January 24, 2019 ([recording](https://youtu.be/bJ_p9t19xqs))



*   **Moderators**: Jorge Castro [SIG Contributor Experience]
*   **Note Taker**: Bob Killen
*   [ 0:00 ]**  Demo **-- Tracing Pod Startup in Kubernetes -- David Ashpole (@dashpole) (confirmed)
    *   [Link to slides](https://docs.google.com/presentation/d/1-vGwSAHYNFod2WURHLSXZtL-F2_7gvJcCSRXanW16Zs/edit?usp=sharing)
    *   Intern Sam did all the work ([@Monkeyanator](https://github.com/Monkeyanator))
    *   Latency problems in Kubernetes are hard
    *   current tools don't cut it: events, logs, latency metrics
    *   distributed tracing helps solve these problems
    *   Uses OpenCensus - Open Source vendor agnostic tracing library
        *   can push to other tracing backends such as zipkin
    *   Very easy to configure on top of Kubernetes
    *   Adds annotation to pod and can follow all events related to that pod across Kubernetes by referencing that annotation
    *   Very useful for debugging complex problems that span multiple components 
    *   Future:
        *   pass trace context through downward api into containers
        *   add trace mechanisms to other Kubernetes objects and CRDs
        *   Trace other object processes e.g. object updates and deletions
        *   Link form spans in trace interface to logs (needs context-aware logging)
    *   [KEP is in review](https://github.com/kubernetes/enhancements/pull/650)
    *   Link to repositories:
        *   [kubernetes-sigs/mutating-trace-admission-controller](https://github.com/kubernetes-sigs/mutating-trace-admission-controller)
*   [ 0:00 ]** Release Updates**
    *   Current Release Development Cycle  [Aaron Crickenberger - Release Manager]
        *   [We are at Week 3 for v1.14](https://github.com/kubernetes/sig-release/tree/master/releases/release-1.14) - [minutes](https://docs.google.com/document/d/1U3jL8Ucruoq9wyzIgxEdyA51MuOIi_gvecVed1kAli0/edit#heading=h.jhszixku4dgy) - [recording](https://www.youtube.com/watch?v=OJv2i7TMh9E)
        *   [Our release team has shadows](m)
        *   [We have a release notes draft](https://github.com/kubernetes/sig-release/blob/master/releases/release-1.14/release-notes-draft.md)
        *   Upcoming milestones:
            *   **[Enhancements freeze](https://github.com/kubernetes/sig-release/tree/master/releases/release-1.14#enhancements-freeze) Tuesday Jan 29** (week 4)
        *   Enhancements
            *   Everything must have a KEP
            *   [PR to update KEP template](https://github.com/kubernetes/enhancements/pull/703)
            *   [repo: kubernetes/enhancements milestone:v1.14](https://github.com/kubernetes/enhancements/milestone/13)
            *   [https://bit.ly/k8s114-enhancements](https://bit.ly/k8s114-enhancements) (contact Enhancements Lead @claurence if you need to add/modify stuff)
            *   What about after the freeze? [Use the exception process](https://github.com/kubernetes/sig-release/tree/master/releases/release-1.14#exceptions)
            *   41 enhancements, 19 in alpha, 11 in beta and 6 for GA.  Would be good to have less alpha, more stable.
        *   CI Signal
            *   [http://bit.ly/k8s114-cisignal](http://bit.ly/k8s114-cisignal)
            *   We watch [release-master-blocking](https://testgrid.k8s.io/sig-release-master-blocking) 
            *   Currently tracked in google doc, likely to change going forward
            *   Implementing [release-blocking job criteria](https://github.com/kubernetes/sig-release/blob/master/release-blocking-jobs.md#release-blocking-criteria): all release-blocking jobs must have owners
                *   Specify "owner" as "an e-mail address that testgrid can send alerts to if the job fails more than N times in a row"
                *   Tracking issue: [https://github.com/kubernetes/sig-release/issues/441](https://github.com/kubernetes/sig-release/issues/441) 
                *   WIP PR: [https://github.com/kubernetes/test-infra/pull/10870](https://github.com/kubernetes/test-infra/pull/10870)
                *   Will be required "soon" (~week?) (k-dev@, look for e-mail, and SIGs, expect to be contacted)
    *   Patch Release Updates
        *   x.x
        *   y.x
*   [ 0:00 ] **SIG Updates**
    *   SIG Leads, check out this set of [recommended topics to cover](https://github.com/kubernetes/community/blob/master/events/community-meeting.md#sig-updates) during your update
    *   [Slide template if you need it](https://docs.google.com/presentation/d/1-nTvKCiqu9UvFYUeM6p6RIqHS5-H-u3_x-V4xj_eIWo/edit#slide=id.g401c104a3c_0_0)
    *   Please also check the [SIG Update Schedule](https://docs.google.com/spreadsheets/d/1adztrJ05mQ_cjatYSnvyiy85KjuI6-GuXsRsP-T2R3k)!
    *   [SIG Node](https://docs.google.com/presentation/d/17TtCtbZrPaPtv5xlIo0K-VKqooDNb6UIrcFMCwMre1k/edit#slide=id.p) [Dawn Chen] (confirmed)
        *   Revised SIG Scope
            *   Kubelet and its features
            *   Pod API and Pod behaviors (with sig-architecture)
            *   Node API (with sig-architecture)
            *   Node controller
            *   Node level performance and scalability (with sig-scalability)
            *   Node reliability (problem detection and remediation)
            *   Node lifecycle management (with sig-cluster-lifecycle)
            *   Container runtime
                *   Proposed new container shim api with containerd community
                *   Important for working with Windows
            *   Device management
                *   More collaboration with containerd community
            *   Image management
            *   Node-level resource management (with sig-scheduling)
            *   Hardware discovery
            *   Issues related to node, pod, container monitoring (with sig-instrumentation)
            *   Node level security and Pod isolation (with sig-auth)
            *   Host OS and/or kernel interactions (to a limited extent)
        *   Accomplishments in v1.13
            *   RuntimeClass - multi container runtime support
            *   efficient heartbeat for scalability (alpha)
            *   better windows support
            *   process namespace sharing promoted to beta
        *   Q1 Updates
            *   graduating / promoting features to next phase
                *   efficient heartbeat - beta
                *   runtimeClass - beta
                *   node OS/arch labels to GA
                *   HugePages - graduated
            *   Improve node-level isolation: pids, userns, disk quota
            *   work with sig-windows to promote windows node to GA
            *   in-place pod resource updates
            *   cpu/device co-scheduling
            *   improve debugging at node level
                *   debug containers alpha
                *   [tracing](https://github.com/kubernetes/enhancements/pull/650)
    *   [SIG API Machinery](https://docs.google.com/presentation/d/1S9faDXJ_cs5oZlN74ysJVtytiCnBjoZliFBmt7mZEd8/edit#slide=id.g401c104a3c_0_0) [David Eads] (confirmed) 
        *   Last Cycle
            *   CRD webhook conversion - alpha 1.13
                *   Please test it, looking for more testers to ensure it covers all use cases.
            *   Dynamic typed informers and listers - 1.13
                *   Should make it easier to build dynamic controllers
        *   Pans for upcoming cycles
            *   path for admission webhooks to GA
                *   need to finish KEP
            *   Server side apply - alpha
            *   Storage migration tool
            *   Deprecating swagger.json (not the "normal" openapi)
            *   Deprecating initializers, never made it past alpha
            *   Investigating API request fairness
                *   looking for comments on design proposal (insert link here)
*   [ 0:00 ] **?Announcements ?**
    *   [spiffxp] Nikhita Raghunath (@nikhita) has joined the [GitHub Admin Team](https://github.com/kubernetes/community/tree/master/github-management)
        *   Huge thanks to Garrett Rodrigues (@grodrigues3) for his time on the team
    *   [spiffxp] Next week's Steering Committee meeting: [we're going to try doing it publicly](https://groups.google.com/a/kubernetes.io/forum/#!topic/steering/47-0nmW5MqY), stay tuned for details

        **? **Shoutouts this week (Check in #shoutouts on slack) **?**

    *   Aaron Crickenberger would like to thank: 
        *   Ben Elder fixing yesterday's prow outage even though he wasn't on call, and following up with a post-mortem: [https://github.com/kubernetes/test-infra/pull/10911](https://github.com/kubernetes/test-infra/pull/10911)
        *   @eduar for reaching out to all of the SIGs and his continued work on revamping the kubernetes dev guide [https://github.com/kubernetes/community/issues/3064](https://github.com/kubernetes/community/issues/3064)
        *   Shoutouts to Hippie Hacker, Tim Hockin and Brendan Burns for transitioning the first piece of project infrastructure to the community: DNS
            *   Editor's note: tim-as-a-service DNS has been deprecated
        *   shoutouts to Bob Killen (@mrbobbytables) for putting together a community documentation style guide [https://github.com/kubernetes/community/pull/3125](https://github.com/kubernetes/community/pull/3125) 
    *   Aramb? Alarc?n has[ hosting facilities in Mexico City](https://discuss.kubernetes.io/t/looking-for-a-meetup-in-mexico-city-facilities-for-gathering-host/4347), would like to start hosting a k8s meetup group. 
    *   Henning Jacobs is collecting a list of [Kubernetes Failure Stories](https://srcco.de/posts/kubernetes-failure-stories.html).
    *   Call for demos for this call, see the top of this document if you're interested in giving a demo. 
    *   Also if you want to guest host this meeting, ping @castrojo or @paris 
    *   #talk-proposals on slack - Place for people to discuss CFPs, talks, share stories and techniques, get peer reviews, etc. 


## January 17, 2019 ([recording](https://youtu.be/Pm0heuYoPnA))



*   **Moderators**:  Paris Pittman, SIG-Contributor Experience
*   **Note Taker**: Josh Berkus 
*   [ 0:00 ]**Demo: **Kamus (confirmed) (forgot to get name)
    *   [Kamus](https://github.com/Soluto/kamus) - A secret encryption/decryption solution for Kubernetes applications. 
        *   open source, avail on Github
    *   [Slides](https://docs.google.com/presentation/d/1XN_eSUwefkVs8HLv_IeVfIx4giu_qN4tb7JdLdxuFQ0)
    *   For storing all types of secrets (API token, certs, client pwd)
    *   Existing secrets solutions are incomplete
        *   Folks use "sealed secrets", but that has limitations too
    *   Demo of Kamus
        *   5 pods running PHP app with "decryptor"
            *   it's one for the app, 4 for Kamus api: 2 handling encryption and 2 handling decryption
        *   Uses an encryptor exec to encrypt the secrets, and an init container to provide app containers with secrets access
        *   Demo didn't work initially, but worked later
    *   Where are secrets being stored?
        *   Multiple encryption options
        *   Stored in either azure keyvault or gcp (google cloud)-kms(?)
            *   (so, stored in cloud provider secrets store)
*   [ 0:00 ]** Release Updates**
    *   Current Release Development Cycle  [Aaron Crickenberger, @spiffxp] (confirmed)
        *   [We are at Week 2 for v1.14](https://github.com/kubernetes/sig-release/tree/master/releases/release-1.14) - [minutes](https://docs.google.com/document/d/1U3jL8Ucruoq9wyzIgxEdyA51MuOIi_gvecVed1kAli0/edit#heading=h.jhszixku4dgy) - [recording](https://www.youtube.com/watch?v=_brMdDFOGeM)
        *   [We cut 1.14.0-alpha.1](https://github.com/kubernetes/sig-release/issues/447) Tuesday Jan 15
            *   problem: no deb/rpm packages because there's only one release channel
            *   WIP: alpha/beta channels
            *   SIG-Release will be adding release-engineering subproject
        *   Upcoming milestones:
            *   [Finalized team by Friday Jan 18](https://github.com/kubernetes/sig-release/issues/372 ) (week 2)
            *   **[Enhancements freeze](https://github.com/kubernetes/sig-release/tree/master/releases/release-1.14#enhancements-freeze) Tuesday Jan 29** (week 4)
        *   Gathering enhancements
            *   [repo: kubernetes/enhancements milestone:v1.14](https://github.com/kubernetes/enhancements/milestone/13)
            *   [https://bit.ly/k8s114-enhancements](https://bit.ly/k8s114-enhancements)
            *   Contact Clair if you need to add/modify stuff
        *   Everything Must Have a KEP
            *   [Updating the template](https://github.com/kubernetes/enhancements/pull/690)
            *   [Discussed with SIG PM](https://docs.google.com/document/d/13uHgcLf-hcR4a5QbV888fhnVsF3djBEpN8HolwS0kWM/edit) ([recording](https://youtu.be/jHXm10TieUI)) re: owning the [implementation](https://github.com/orgs/kubernetes/projects/5)
            *   [Discussing at SIG Arch](https://docs.google.com/document/d/1BlmHq5uPyBUDlppYqAAzslVbAO8hilgjqZUTaNXUhKM/edit) right after this
            *   Talked to SIG-Apps, Phil Witrock had suggestions in the form of at PR
            *   WIP, but moving design proposals to KEPs
    *   Patch Release Updates
        *   [Schedule](https://github.com/kubernetes/sig-release/blob/master/releases/patch-releases.md)
        *   v1.10.13
            *    [there was much discussion of whether to cut this at SIG-Release](https://github.com/kubernetes/kubernetes/pull/72860#issuecomment-454579674)
            *   it's beyond the 3 supported versions, but
                *   we still have the CI set up for 1.10
                *   fixing a regression introduced in 1.10's last patch release
*   [ 0:00 ] **Contributor Tip of the Week**
    *   **Mentoring!!**
        *   Grow your contributors!
            *   [Meet Our Contributors - monthly youtube series](https://www.youtube.com/playlist?list=PL69nYSiGNLP3QpQrhZq_sLYo77BVKv09F) 
                *   **New** and **current contributors**
                *   Two sessions for global span 3:30pm and 9pm UTC
                *   To ask a question: #meet-our-contributors in slack or DM paris for anonymity 
                *   Current contributor benefits: an entire session with the steering committee to answer questions around: governance, structure of project, how they got involved, current business. 
                    *   During regular mentor panel session: ask why is your test(s) flaking, how to be a subproject owner, what SIGs are looking for more contributors, etc 
            *   [Google Summer of Code!](https://github.com/cncf/soc)
                *   CNCF submitted our application - we are aiming for as many as possible
                *   Have an interesting project and/or can you mentor? Email was sent from Nikhita kubernetes-dev@google.com 
            *   New contributor workshop
                *   Every KubeCon - will be listed in the co-located events section, announced via blog, and kubernetes-dev@googlegroups.com
                *   Check out the playlist from the last one in December
                *   Onboarded ~200 new contributors last year
                *   Also a good way for current contributors to meet new ones
                *   The videos are also useful for online new contribs
            *   Release Team!
                *   Get a unique chance to see a lot of the projects parts and how things work
                *   [https://www.surveymonkey.com/r/k8s-114-rt-shadows](https://www.surveymonkey.com/r/k8s-114-rt-shadows)
            *   Future
                *   remote pair programming
                *   other new ideas
*   **[ **0:00** ] SIG Updates**
    *   **CLI **(@seans3): [SIG CLI Update Slides](https://docs.google.com/presentation/d/1_xV403gyzKZt19fXb7u1_UZCyLFSgesxTZEcn5A1IaQ/edit?usp=sharing)
    *   Subprojects now:
        *   missed stuff here (look at slides for more)
        *   Kustomize - filling in gaps
    *   Current work:
        *   extension mechanisms like plugins and dynamic commands
        *   moved kubectl outside of kubernetes/kubernetes
        *   declarative management of apps with kustomize
        *   plus server-side apply, which will support "diff"
        *   merged their charter
    *   More about plugins
        *   now beta!
        *   plugin is binary prefaced by "kubectl-"
        *   new repo is kubernetes/cli-runtime, including plugin stuff
        *   see Seattle SIG-CLI deep dive
        *   working on krew, early stages as plugin manager
    *   dynamic command extensions
        *   want to make "kubectl create <name>" work, which is data-based
        *   just starting work on this now
    *   Moved out of core
        *   kubernetes/kubectl
        *   pkg/kubectl will move to a staging repo
    *   Better declarative workflow
        *   kustomize will merge into kubectl
        *   lets you alter YAML doing kube-aware patching
        *   see KEPs, kubernetes-sigs/kustomize
        *   server-side apply supports intelligent "diffing"
    *   New charter:
        *   extra roles, Emeritus Lead and Test Health Manager
    *   Sending out a survey to kubectl users about features soon
    *   
    *   UP NEXT WEEK: NODE, CLOUD PROVIDER, AND API MACHINERY
*   [ 0:00 ] **?Announcements ?**
    *   [KubeCon Barcelona cfp submissions due JANUARY 18TH!! ](https://linuxfoundation.smapply.io/prog/kccnceu19/)
        *   SIG sessions due Feb 8th
    *   Placeholder for GitHub Management subproject [spiffxp]
        *   [Automating team management via PRs to kubernetes/org](https://github.com/kubernetes/org/issues/336)
            *   [kubernetes-dev@ notification](https://groups.google.com/forum/#!topic/kubernetes-dev/dwHkzW6QyTU)
            *   [TBD: docs](https://github.com/kubernetes/community/issues/3102)
            *   Gets us closer to the possibility of [using teams in OWNERS files](https://github.com/kubernetes/test-infra/issues/10065#issuecomment-455278672)
        *   Aaron thinks we've turned it on and so far everything is working
            *   The idea is to handle team membership via PRs
            *   Also to allow teams to be self-managing
        *   If you're interested in helping, please speak up!
    *   Placeholder for Steering [spiffxp]
        *   (updates summary not formally approved)
        *   New steering committee, first meeting was very productive
        *   [Product Security Team to become a Committee](https://github.com/kubernetes/steering/issues/89)
            *   now has rules on who's on it and how it relates to the project
        *   WG-k8s-infra now exists, for migrating stuff off of Google ownership mainly
            *   very early stages, don't even know how stuff will work yet
            *   need help building systems & rules
        *   Have to review charter for CoC Committee
        *   Also want a concept for a group which is just for focused user discussion
            *   not a WG, or SIG, or subproject
        *   Considering opening up meetings for public participation
            *   No idea how this will work yet, but thinking about it
            *   Maybe every other week
            *   Watch kubernetes-dev
        *   [Meeting notes](http://bit.ly/k8s-steering-wd)
    *   [Contributor Experience charter is merged!!](https://github.com/kubernetes/community/tree/master/sig-contributor-experience)
    *   WE NEED MENTORS!! Meet Our Contributors could be a one time one hour commitment that will help hundreds!! Reach out to Paris or #sig-contribex on slack.
    *   **Shoutouts** ?
        *   Slack! #shoutouts 
        *   [justagustus] Shoutout to @jberkus for putting together the draft questions (https://github.com/kubernetes/sig-release/issues/368) for our first ever Release Team shadows questionnaire (https://www.surveymonkey.com/r/k8s-114-rt-shadows)!

            ...and everyone who helped review it! @ihor.dvoretskyi @maria @liggitt @mbohlool @spiffxp @marpaia @bentheelder @tpepper @aleksandram @AishSundar

*   [liggitt] All hail @dims for running the 0-length flake to ground
*   [coderanger] Shoutout to the whole ZH docs translation crew, and a special mention for Adam Dang as putting in a ton of work! In total the team has merged 444 PRs over the past two months of Chinese translation!
*   [spiffxp] shoutout to @akutz for stepping to take notes for sig-testing's weekly meetings, we go a mile a minute and it's much appreciated!
*   [spiffxp] shoutout to @nikhita for moving the kubernetes project values to kubernetes/community for more exposure (https://github.com/kubernetes/steering/pull/88) and improving our WG docs generated from sigs.yaml ([https://github.com/kubernetes/community/pull/3069](https://github.com/kubernetes/community/pull/3069))
*   [nikhita] Shoutout to @mspreitz for adding lots of details to the code-generator conversion-gen docs! [https://github.com/kubernetes/kubernetes/pull/71821](https://github.com/kubernetes/kubernetes/pull/71821) 


## January 10, 2019 - ([recording](https://youtu.be/XXZ0ekSX3cg))



*   **Moderators**:  Josh Berkus [SIG-Release]
*   **Note Taker**: Solly Ross [Google]
*   [ 0:00 ]**  Demo **-- Krew: kubectl plugin manager [Ahmet Alp Balkan, @ahmetb] (confirmed)
    *   [Link to repository](https://github.com/GoogleContainerTools/krew)
    *   [Link to slides](https://docs.google.com/presentation/d/1TTSdInmHbchyAK6lzkVQCUpQNDJcxe6SZGEB5IfRq10/edit#slide=id.p)
    *   Kubectl plugins are stable as of 1.12 -- can extend kubectl by adding new commands with kubectl-foo binaries
        *   e.g . bespoke commands for workflow or extensions
    *   Krew -- like homebrew (brew) for kubectl
        *   Easy way to discover/install plugins, keep up to date
        *   Is plugin itself (`kubectl krew install/upgrade/remove`)
        *   Can easily package for multiple platforms (windows, linux, osx) -- just write manifest pointing at hosting location and files
        *   Doesn't support:
            *   External dependencies (e.g. python)
            *   Version skew
            *   Security scanning
    *   Plugin index
        *   Centralized in YAML file for the moment
        *   3rd-party package index support in the works
*   [ 0:12 ]** Release Updates**
    *   Current Release Development Cycle  [Aaron Crickenberger, @spiffxp] (confirmed)
        *   We are at Week 1 for v1.14
        *   Upcoming milestones
            *   [Finalized schedule by Friday Jan 11](https://github.com/kubernetes/sig-release/pull/431)
                *   Currently targeting late march for release (see above link)
            *   [Finalized team by Friday Jan 18](https://github.com/kubernetes/sig-release/issues/372 ) (week 2)
            *   **Enhancements freeze Tuesday Jan 29** (week 4)
        *   What will we do differently?
            *   **Everything must have a KEP** ([discussing at sig arch at 11am PT](https://docs.google.com/document/d/1BlmHq5uPyBUDlppYqAAzslVbAO8hilgjqZUTaNXUhKM/edit#))
                *   Even stuff that didn't have a KEP before
                *   Needs acceptance criteria
                *   Test, upgrade plans desired
            *   [No code slush](https://github.com/kubernetes/sig-release/issues/269#issuecomment-452062051)
                *   Memes to be delivered instead to reduce paperwork
            *   Using [repo:kubernetes/sig-release milestone:v1.14](https://github.com/kubernetes/sig-release/milestone/7) for release team work (tracking blocking stuff, etc)
            *   New year, new kubernetes-milestone-maintainers GitHub team: [we're purging and starting over with current SIG Release chairs, Release Team members, and SIG chairs/TLs](https://groups.google.com/d/msg/kubernetes-dev/BveUpMt-qO4/ODeT9lLrDwAJ)
                *   This group lets you use the `/milestone and /status `commands
                *   **If some else manages these commands in your SIG, please let SIG release know**
    *   Patch Release Updates
        *   [Let's try using a schedule](https://github.com/kubernetes/sig-release/blob/master/releases/patch-releases.md)
        *   1.13.2: Today
        *   1.12.5: [cut planned Thursday, Jan 17th](https://groups.google.com/forum/#!topic/kubernetes-dev/7hLP6bpvrr8)
*   [ 0:21 ] **Contributor Tip of the Week **[Aaron Crickenberger] 
    *   Let's talk about project boards
    *   Creation:
        *   Per repo board: needs repo write access (eg: [kubernetes/kubernetes](https://github.com/kubernetes/kubernetes/projects))
            *   Has issues with scalability for an org of our size
        *   Org-wide board: _Any kubernetes org member _can create a [kubernetes org project board](https://github.com/orgs/kubernetes/projects)
    *   Their [permissions](https://github.com/orgs/kubernetes/projects/1/settings) are very similar to repo permissions
        *   Visibility: public / private
        *   Org members: admin / write / read / none
        *   Teams: admin / write / read
        *   Collaborators: admin / write / read
        *   Higher permissions higher in the list override lower permissions lower in the list
    *   eg: [SIG Contributor Experience](https://github.com/orgs/kubernetes/projects/1)
        *   Recurring standup during their meeting
        *   Permissions are locked to community-admins, community-maintainers
    *   eg: [Deflaking kubernetes e2e tests](https://github.com/orgs/kubernetes/projects/7)
        *   [Umbrella issue](https://github.com/kubernetes/community/issues/3071)
        *   Anyone can add to it, move cards around, etc
    *   Can see activity with activity tab
    *   Columns can have (github-provided) automation presets applied:
        *   To Do
        *   In Progress
        *   Done
    *   Ideas for automation our community is working on
        *   [Add issues via a prow command](https://github.com/kubernetes/test-infra/issues/10514)
        *   [Auto add issues/PRs based on GitHub query](https://github.com/kubernetes/test-infra/issues/9925)
    *   [Brainstorming project management improvements](https://github.com/kubernetes/community/issues/3079)
        *   If you have suggestions, chime in ^
*   [ 0:27 ] **Open KEPs** [Kubernetes Enhancement Proposals]
    *   [Coscheduling](https://github.com/kubernetes/enhancements/blob/master/keps/sig-scheduling/34-20180703-coscheduling.md)  [Klaus Ma SIG-Scheduling] (confirmed)
    *   [slides here]
    *   Motivation: Some workloads (e.g. batch data processing) need all pods to start together
        *   If some don't start/get the right resources, everything should fail
        *   May need some minimum (softer requirement than "everything")
    *   Proposal
        *   Introduce "group name" annotation, scheduler considers all pods a group as needing to start together
        *   Can separately specify minimum start number for a group
        *   Can mark group as "restart entire group if one pod fails"
    *   Quota brings some issues (quota could block things from creating/starting)
        *   Can mark group as reserving some total amount of resources for the group
    *   Status
        *   Support in [kube-batch](https://github.com/kubernetes-sigs/kube-batch) 0.2+
        *   Ongoing work: PodGroupController, Quota support, better starvation behavior
    *   Other Kube-batch features: Queues, preemption, and more
*   [ 0:00 ] **SIG Updates**
    *   SIG Leads, check out this set of [recommended topics to cover](https://github.com/kubernetes/community/blob/master/events/community-meeting.md#sig-updates) during your update
    *   [Slide template if you need it](https://docs.google.com/presentation/d/1-nTvKCiqu9UvFYUeM6p6RIqHS5-H-u3_x-V4xj_eIWo/edit#slide=id.g401c104a3c_0_0)
    *   Please also check the [SIG Update Schedule](https://docs.google.com/spreadsheets/d/1adztrJ05mQ_cjatYSnvyiy85KjuI6-GuXsRsP-T2R3k)!
    *   SIG PM [Stephen Augustus] (confirmed)
        *   Slides: [https://docs.google.com/presentation/d/1IcrTbZCGlZGZKnBO6srYqPZiwGrDdUnacm7OlF5GM1o/edit?usp=sharing](https://docs.google.com/presentation/d/1IcrTbZCGlZGZKnBO6srYqPZiwGrDdUnacm7OlF5GM1o/edit?usp=sharing)
        *   Intro and deep dive from kubecon available on youtube
        *   Last cycle
            *   Survey on how people work with/use Kubernetes
            *   Improved KEP process (KEP-1a -- how do you implement usage of KEPS)
            *   Moved some content/repos to unify how we track/manage multi-release work (KEPs)
                *   k/features ? k/enhancements
                *   KEPs ? k/enhancements
            *   Categorizing KEPs
        *   Next cycle
            *   Revamp SIG PM charter (align with standard charter setups)
            *   KEP
                *   Clean up KEP process documentation to be clearer
                *   Designate who owns different parts of KEPs at different points in lifecycle of the KEP
                *   Continue migrations
                    *   Design proposals ? k/enhancements (maybe)
                        *   Need to start pruning/figuring out what's relevant
                *   Make it easier to work with KEPs:
                    *   KEP CLI tool (easily update keps, scaffold new ones, etc)
                    *   KEPs on contributor site (for easy browsing, consuming)
                *   KEP GA in 1.15
            *   Burn all the spreadsheets
            *   Make it easier to track projects cross-SIG
    *   SIG Autoscaling [@mwielgus] (confirmed)
        *   Responsible for all the components that adjust cluster objects for cluster needs (VPA, Cluster autoscaler, HPA)
        *   Current features
            *   Faster HPA scaling
            *   Resolving problems with pod priorities in Cluster Autoscaler (a couple remain)
            *   Vertical Pod Autoscaler to beta
            *   Alibaba cloud support in Cluster Autoscaler
        *   Upcoming features
            *   API for scale up/down speed in HPA
        *   Meeting every monday (7:00 AM PST)
    *   [SIG Network [Bowei Du]](https://docs.google.com/presentation/d/13l5gb7MtieQkkMwUGAV5mxip_mGmgvI-_b41WLOhShg/edit#slide=id.g401c104a3c_0_0) (confirmed)
        *   In progress (see slides for more info/links):
            *   IPv6
            *   Custom DNS policy
            *   Pod readiness gates
            *   SCTP support
            *   Node-local DNS caching
        *   Please try stuff out and submit feedback!
        *   Upcoming themes (see slides for links)
            *   Revamping Ingress/L7
            *   Dual stack (IPv4 + IPv6)
            *   Topology-aware services (e.g. node local services)
            *   Revamping services and endpoints
            *   Multicast support
            *   Windows support
        *   Meetings every other Thursday
        *   Looking in to contributor on-ramping guide
*   [ 0:00 ] **?Announcements ?**
    *   New year, new kubernetes-milestone-maintainers GitHub team: [we're purging and starting over with current SIG Release chairs, Release Team members, and SIG chairs/TLs](https://groups.google.com/d/msg/kubernetes-dev/BveUpMt-qO4/ODeT9lLrDwAJ) [Stephen Augustus]
    *   [We're going to enable automated github team management](https://groups.google.com/forum/#!topic/kubernetes-dev/dwHkzW6QyTU) on Friday [Cristoph Blecker]
    *   Kubecon! May 20 ? 23, 2019  |  Fira Barcelona, Barcelona, Spain
        *   **CFP Closes on January 18th!**
        *   [https://events.linuxfoundation.org/events/kubecon-cloudnativecon-europe-2019/cfp/](https://events.linuxfoundation.org/events/kubecon-cloudnativecon-europe-2019/cfp/)
    *   Next week SIG updates: CLI and Node!
        *   Also: **we need demos** (see the top of this document for instructions)
    *   Revamping Dev Guide! If you have a suggestion for the persona of upstream developer -> [https://github.com/kubernetes/community/issues/3064](https://github.com/kubernetes/community/issues/3064). Persona #2 - building on top of K8s is the next one up. 
    *   **Shoutouts:** 
        *   [paris, jberkus] Shoutout to @coderanger for making LWKD happen over the holidays, when I was completely unavailable.  @coderanger also has been super supportive in slack fielding contributor questions while folks are out - thank you!


## January 3rd, 2019 - ([recording](https://www.youtube.com/watch?v=gBfGqyn06bk))



*   **Moderators**:  Jorge Castro [SIG Contributor Experience]
*   Happy New Year!
*   **Note Taker**: Bob Killen [SIG Contributor Experience/University of Michigan]
*   [ 0:00 ]**  Demo **-- OpenLab - Melvin Hillsman ([mrhillsman@gmail.com](mailto:mrhillsman@gmail.com)) - OpenLab is curated infrastructure for open source testing [https://openlabtesting.org](https://openlabtesting.org)
    *   [https://docs.google.com/presentation/d/1DDeXWafI2ucRAwKyl9sStjtMb-yz6w-5DFXyOPNs6ks](https://docs.google.com/presentation/d/1DDeXWafI2ucRAwKyl9sStjtMb-yz6w-5DFXyOPNs6ks)
    *   [https://bit.ly/openlabstart](https://bit.ly/openlabstart)
    *   Curated infrastructure for Open Source Testing focusing on 3 things:
        *   Reduce friction of cloud ecosystem tooling integration
        *   Testbed delivery and maintenance efficiency
        *   cross-community collaboration 
    *   Built completely on Open Source tooling
    *   Dedicated Resources include:
        *   dedicated physical servers
        *   virtual machines
        *   network devices
        *   IoT, GPUs FPGAs
        *   Containers
    *   Quite a few projects make use of the project already
    *   chat: #askopenlab on freenode for user support
    *   mailing list: openlab.groups.io
*   [ 0:00 ]** Release Updates**
    *   Current Release Development Cycle  [Aaron Crickenberger - Release Lead]
        *   We are at Week 0 for v1.14, [release team leads finalized](https://github.com/kubernetes/sig-release/blob/master/releases/release-1.14/release_team.md)
        *   v1.14 schedule draft being reviewed by former release leads, current release lead shadows, sig release chairs
        *   Aiming for kickoff next week, all release team shadows finalized by Friday Jan 11th
        *   Modest proposal: to land in this release, you must have a KEP, even if you didn't before, and that KEP must have a test plan, and an upgrade/downgrade plan
            *   will be discussed at length during next week's sig-arch meeting
    *   Patch Release Updates
        *   discussion ongoing on setting up a schedule for patch releases
*   [ 0:00 ] **SIG Updates**
    *   SIG Leads, check out this set of [recommended topics to cover](https://github.com/kubernetes/community/blob/master/events/community-meeting.md#sig-updates) during your update
    *   [Slide template if you need it](https://docs.google.com/presentation/d/1-nTvKCiqu9UvFYUeM6p6RIqHS5-H-u3_x-V4xj_eIWo/edit#slide=id.g401c104a3c_0_0)
    *   Please also check the [SIG Update Schedule](https://docs.google.com/spreadsheets/d/1adztrJ05mQ_cjatYSnvyiy85KjuI6-GuXsRsP-T2R3k)!
    *   SIG Autoscaling, Networking, and PM due next week! 
    *   SIG Apps [Matt Farina] (confirmed)
        *   [Slides](https://docs.google.com/presentation/d/1jbERrg36lsR05xdDmkF42NoWsM7-rMmQD9GzNYdvPXQ/edit?usp=sharing)
        *   Last Cycle
            *   Charter completed and merged
            *   Figured out needs to make CronJobs GA
            *   Started work on Portable Service Definitions
            *   Work on Application Controller
        *   Upcoming Cycle
            *   Looking for lead on CronJob GA
            *   Begin work on Portable Service Definitions
            *   Application Controller Status
            *   Deprecation of Beta APIs
        *   CronJobs to GA
            *   Currently still batch/v1beta1
            *   Has scalability issues
            *   Controller needs to be rewritten
            *   Looking for contributors
        *   Portable Service Definitions
            *   [insert kep link here]
            *   enable an application to be deployed into multiple environments while relying on external services
            *   Will be built on CRDs + controllers
            *   Looking to solve some UX consistency issues
            *   Looking for contributors
        *   Application Controller Status
            *   Bubble up application deployment rollup status
            *   How to get status for multiple components of an application
        *   Deprecation of Beta APIs
            *   Continue to support beta APIs, despite "formal" deprecation some time ago
            *   Turn off in 1.15 with optional flag to re-enable.
        *   How to Contribute
            *   Lots of opportunities with CronJobs etc
    *   SIG UI [Jeffrey Sica] (confirmed)
        *   [Slides](https://docs.google.com/presentation/d/18H51DlgR_WsBhdRcUF1ru0cuDE7i7m1WsPzLrl2cG3M/edit#slide=id.g338ac0a8b6_0_27)
        *   Last Cycle
            *   Finished and merged SIG-UI charter
            *   2 releases including fix for CVE-2018-18264
            *   Metrics server support
            *   Angular Migration branch merged to master (entire front-end rewrite)
                *   versioning schema will change for future releases
            *   Annual Survey of dashboard users
        *   Upcoming Cycle
            *   Formalize metrics server support
                *   Current solution is stop-gap
                *   Will support prometheus and other sinks in the future
            *   Versions will now be 2.x.x
            *   Better OAuth support
        *   How can you contribute
            *   looking for help with metrics
    *   SIG VMWare [Steve Wong] (confirmed) 
        *   [Slides](https://docs.google.com/presentation/d/1eMaclhtHY2llnmLWe6BoJv1xdmXpaMe79xNjH-klOvY/edit?usp=sharing)
        *   Last Cycle
            *   External vSphere Cloud Provider [alpha]
            *   CSI provider for vSphere
            *   Cluster API provider for vSphere
        *   Upcoming Cycle
            *   Bring external vSphere Cloud Provider to stable release status
            *   Bring CSI provider for vSphere to stable release status
            *   Cluster API provider for vSphere
                *   improve e2e tests
        *   Working to provide licenses for Fusion/Workstation to support minikube CI/CD
        *   If there are any licensing issues when working with commercial VMware components, reach out for license and support help
*   [ 0:00 ] **?Announcements ?**
    *   May 20 ? 23, 2019  |  Fira Barcelona, Barcelona, Spain
        *   CFP Closes on January 18th!
        *   [https://events.linuxfoundation.org/events/kubecon-cloudnativecon-europe-2019/cfp/](https://events.linuxfoundation.org/events/kubecon-cloudnativecon-europe-2019/cfp/)
    *   If you attended the Contributor Summit, please fill out the survey (in email). 
    *   Want to demo during the first 10m of this meeting? Check out the top of this document to sign up! 
    *   Slides/Videos of Kubecon Sessions: [https://github.com/cloudyuga/kubecon18-NA](https://github.com/cloudyuga/kubecon18-NA) (Thanks Cloudyuga!) 

        **? **Shoutouts this week  **?**


        See someone doing great work? Let us know in #shoutouts on slack so we can highlight their work!

    *   Jeff Grafton: shoutout to @cblecker for breaking out the godep verification checks into their own job, bringing the rest-of-the-verify-checks job down to 41m
    *   Jeff Grafton: shoutout to @fisherxu for finally fixing our generated code to not include the year ([https://github.com/kubernetes/kubernetes/pull/59172](https://github.com/kubernetes/kubernetes/pull/59172)), thus preventing the build from breaking and needing "happy new year!" PRs like [https://github.com/kubernetes/kubernetes/pull/57735](https://github.com/kubernetes/kubernetes/pull/57735), [https://github.com/kubernetes/kubernetes/pull/39342](https://github.com/kubernetes/kubernetes/pull/39342), and [https://github.com/kubernetes/kubernetes/pull/19222](https://github.com/kubernetes/kubernetes/pull/19222)
    *   @idealhack: Shoutout to the Chinese reviewer team of SIG docs (Chen Rui @rui, Adam Dang, Xiaolong He, Peter Zhao @xiangpengzhao), the blog post about our 2nd New Contributor Workshop at KubeCon (https://kubernetes.io/blog/2018/12/05/new-contributor-workshop-shanghai/) is now available in Chinese [https://kubernetes.io/zh/blog/2018/12/05/??????????](https://kubernetes.io/zh/blog/2018/12/05/)/. PR: [https://github.com/kubernetes/website/pull/11896](https://github.com/kubernetes/website/pull/11896)


## December 6, 2018 (recording)



*   **Moderators**:  Paris Pittman [Google, ContribEx]
*   **Note Taker**: none - retro
*   [ 0:00 ]**  Demo **-- NO DEMO - RETRO TIME :partyk8s:
*   [ 0:00 ] **?Announcements ?**
    *   **Contributor Summit is at capacity! See everyone in Seattle!!**
    *   **Shoutouts!**
        *   **#shoutouts in slack!**
        *   Aish: Shoutouts to @neolit123 @dims @liggitt for your last minute heroics in getting all the external dependencies in place for 1.13 release :wave::skin-tone-2:
        *   Liggit: shoutout to @ibuildthecloud for finding and reporting https://github.com/kubernetes/kubernetes/issues/71411 responsibly
        *   Zach C from docs: Shoutouts to @tfogo, @jimangel, and @jrondeau for finding and fixing https://github.com/kubernetes/website/pull/11503, this time for good.
            *   Jrondeau - don't forget @zacharysarah on that one too!
        *   Josh berkus: @msau42 for figuring out a major cause of e2e test flakiness!
        *   Spiffxp: shout outs to @paris for making sure the contributor summit sessions from kubecon will be recorded!!!!
*   [ 0:00 ]** Release Updates**
    *   **Retro 1.13 -> [http://bit.ly/k8s113-retro](http://bit.ly/k8s113-retro) **


## November 29, 2018 - ([recording](https://youtu.be/e5oHnmUOz3k))



*   **Moderators**:  Josh Berkus [SIG-Release]
*   **Note Taker**: Solly Ross [Google/SIG Autoscaling]
*   [ 0:00 ]**  Demo **-- Docs Modeling Working Group Demo [Andrew Chen, @chenopsis, Dominik (dominik.tornow@sap.com)] (confirmed)
    *   Link to [slides](https://docs.google.com/presentation/d/1Ycs1-PcnctWRC9wvSbGkM3IfL2hcguQ42wbWp15DkeY/edit?usp=sharing)
    *   Modelling how we design and look at documentation
        *   Idea:
            *   Ideally, two people who look at the same system develop the same mental model
            *   Looking a documentation, there may be encoding/decoding loss (writing/reading docs), which leads to different mental models
        *   **F**undamental **M**odeling **C**oncepts
            *   Approach to system modeling with formal models of system's structure and behavior
            *   Diagrams and formal models can help show whole-picture view
            *   Show people how things work without needing to point people at actual source code
    *   Issues with existing docs:
        *   docs are task focused (good for on-demand "how do I" type questions), but can't easily develop a coherent general picture
        *   non-obvious behavior doesn't match general mental model, docs should help fix that
    *   Process:
        *   Ongoing: Discuss models (in SIG Docs), Interview engineers, validate models, create source materiel (e.g. Medium posts) and get feedback
        *   Eventually: fold back into to k8s.io (planned for next year)
*   [ 0:13 ]** Release Updates**
    *   Current Release Development Cycle  [Aish Sundar - Release Manager]
        *   **Code freeze for 1.13 is now lifted!** Code thaw went into effect 11/28, 8pm PST.
        *   Master is now open for 1.14 development.  
        *   _Only the_ _absolute most critically urgent bug fixes_ might be cherry picked back in time for 1.13.0. 
        *   1.13-rc.2 slated to cut tomorrow, 11/30.
        *   The release is on target for** Monday, 12/3/2018**, pending CI signal.
        *   **_If you still have outstanding Docs PR or Release notes, please get -them in ASAP._**
        *   We're targeting our release retrospective for next week's [Community Meeting](https://docs.google.com/document/d/1VQDIAB0OqiSjIHI8AWMvSdceWhnz56jNpZrLs6o7NJY/edit) on 12/6.  Please add any comments you'd like included in discussion for things that worked well and things that should change in our [1.13 retrospective document](http://bit.ly/k8s113-retro).
    *   Patch Release Updates
        *   1.12.3
        *   1.11.5
        *   1.10.11
*   [ 0:16 ] **SIG Updates**
    *   Info for SIGs:
        *   SIG Leads, check out this set of [recommended topics to cover](https://github.com/kubernetes/community/blob/master/events/community-meeting.md#sig-updates) during your update
        *   [Slide template if you need it](https://docs.google.com/presentation/d/1-nTvKCiqu9UvFYUeM6p6RIqHS5-H-u3_x-V4xj_eIWo/edit#slide=id.g401c104a3c_0_0)
        *   Please also check the [SIG Update Schedule](https://docs.google.com/spreadsheets/d/1adztrJ05mQ_cjatYSnvyiy85KjuI6-GuXsRsP-T2R3k)!
    *   SIG Architecture [Matt Farina] (confirmed)
        *   Slides: [https://docs.google.com/presentation/d/1jwk23WLDLAs5PcFbTKC-P-Hknemu30IlL0u9iWcwT9M/edit?usp=sharing](https://docs.google.com/presentation/d/1jwk23WLDLAs5PcFbTKC-P-Hknemu30IlL0u9iWcwT9M/edit?usp=sharing) 
        *   Updates and ongoing work:
            *   Charter was accepted
            *   KEP process work ongoing
            *   Code processes (e.g. Go modules, how to import k8s.io/kubernetes)
            *   Windows GA plan with SIG Windows, conformance profiles to figure out how to handle conformance for mixed clusters
        *   API Review Process Subprojects
            *   There's a [GitHub board](https://github.com/kubernetes-sigs/architecture-tracking/projects/3) to show backlog, status, etc
        *   KEPs subproject
            *   **KEPs are moving to kubernetes/enhancements, merge before tomorrow unless you want to re-file your PRs**
            *   [Board to track KEP reviews](https://github.com/kubernetes-sigs/architecture-tracking/projects/2)
            *   Each dashboard has KEPs related-to particular SIGs, so you can see what KEPs other SIGs have tagged as related to yours
        *   Conformance Subproject
            *   Board to track issues
            *   Develops conformance testing
        *   SIG Arch meetings are right after the community meetings, come join
    *   SIG Release [Tim Pepper] (confirmed)
        *   SIG Release Update
            *   Slides: [https://docs.google.com/presentation/d/1WtmoYP1Ay9hq2XuPEBQRVorH5pilJ6829qyQD5qWmwY/edit?usp=sharing](https://docs.google.com/presentation/d/1WtmoYP1Ay9hq2XuPEBQRVorH5pilJ6829qyQD5qWmwY/edit?usp=sharing) 
            *   Task:
                *   Put out the releases
                *   Mentor release team so that people don't burn out
                *   Work with other SIGs on release tooling and automation (e.g. with SIG Testing)
                *   Help other SIGs to make sure repos can be part of a release process
            *   See the slides for some nice diagrams about process and timelines
            *   Last cycle
                *   Chairs changeup
                *   Features are now called enhancements
                *   1.12: _First non-Google branch manager, _Move to Tide, Shorter code freeze
            *   Next cycles:
                *   1.13: **Flaky/unmaintained tests moved to non-blocking, **Moving stable branch management to a team of people instead of an individual
                *   1.14: better RPMs and Debs, better build tools, continue working on Tide/automation/labels, continue working on code freeze
            *   Subprojects
                *   Working on licensing and compliance
                *   Working on release security processes
                *   WG LTS (see below)
        *   WG LTS - aka investigating ways to improve the Kubernetes support story
            *   PR for wg addition inbound at [https://github.com/kubernetes/community/pull/2911](https://github.com/kubernetes/community/pull/2911)
            *   PR for clarifying what is supported today: [https://github.com/kubernetes/website/pull/11060](https://github.com/kubernetes/website/pull/11060)
            *   Announcing meeting, google group, slack info, links: [https://groups.google.com/d/msg/kubernetes-dev/Fqya3zlt0QQ/K6JapQX3AAAJ](https://groups.google.com/d/msg/kubernetes-dev/Fqya3zlt0QQ/K6JapQX3AAAJ)
        *   Come get involved!  Always looking for volunteers
            *   [https://git.k8s.io/community/sig-release/README.md](https://git.k8s.io/community/sig-release/README.md)
*   Please drop a note in the community meetings doc, or reach out if you want to talk about a KEP in the community meetig
*   [ 0:00 ] **?Announcements ?**
    *   Contributor Summit [Paris and Jorge]
        *   **We are sold out/waitlisted - L A S T   C A L L  if you're a SIG Chair, TL, or subproject owner **
        *   Talks have been added to the community calendar, shortcut: [http://bit.ly/kubernetes-summit](http://bit.ly/kubernetes-summit)
        *   Check out #contributor-summit on slack
        *   [Event information](https://git.k8s.io/community/events/2018/12-contributor-summit)
    *   Community Meeting Schedule - there are no SIG updates for December.
        *   Today is the last "normal" community meeting
        *   12/6 - Release Retro for 1.13 (tentative!)
        *   12/13 - **Kubecon, no community meeting**
        *   12/20 and 12/27 - **No community meetings**
        *   January 1/3 : SIG Apps,  SIG UI, SIG VMWare
    *   [Meet Our Contributors](https://github.com/kubernetes/community/blob/master/mentoring/meet-our-contributors.md) will be 5 December.
        *   Steering Committee AMA @ 730a PT / 330pm UTC
        *   Mentor panel @ 1pm PT / 9pm UTC
        *   Be a mentor to hundreds with one hour of your time! Reach out to parispittman@google.com / "paris" on slack to get scheduled.
    *   No k8s office hours this month - thanks to all the volunteers who helped make the program a success this year. 
    *   **? **Shoutouts this week **?**
        *   Twitterverse shoutouts for our fearless 1.13 Release Team Lead, @AishSundar: https://twitter.com/stephenaugustus/status/1063610123149545472?s=19
        *   Shoutout to @amerai for adding a search bar to Testgrid so that you don't have to dig to find the right dashboard! https://testgrid.k8s.io/
        *   to @mkimuram & @saad-ali & @msau42 for rapid response to multiple storage test issues with new features.
        *   to @mrhohn for fast & insightful help with sig-network test failures
        *   Huge shoutouts to the entire 1.13 Release leads and shadows for their stellar efforts at every stage throughout the cycle, enabling us to stabilize and  hopefully land the release on time -  @kacole2 @jberkus @cjwagner @dougm @nikopen @tfogo @marpaia @kbarnard10 @spiffxp @tpepper@aleksandram!
        *   Special shoutout to contributors "technically" not on the release team, but have been instrumental in getting us unblocked at numerous points this release with their reviews, test fixes and test-infra support - @dims  @liggitt @justinsb @cblecker @bentheelder @justaugustus (edited)
        *   shoutout to you @AishSundar for keeping us all in line this whole cycle! you've been a totally awesome release lead.
        *   Shoutout to @mrbobbytables for significally reducing my admin overhead for the New Contributor Workshop!


## ~~November 22, 2018~~ (_Note - Thanksgiving in US - **cancelled**)_



*   **Moderators**: Ihor Dvoretskyi [SIG PM/CNCF]
*   **Note Taker**: -
*   [ 0:00 ]** Release Updates**
    *   1.13 is in **Code Freeze - 11/16 to 11/28**
    *   Beta.2 targeted EOD, 11/22 
    *   Still targeting 12/3 for release date.
    *   **[Latest CI Signal Report](https://groups.google.com/forum/#!topic/kubernetes-dev/4agE2-5vXQk)**
        *   No unknown blockers at this point
        *   Lots of chronic flaky tests that could impact release if they persist
    *   **Docs - Yellow**
        *   All Docs PR should be in review by now
        *   Final PR merge** deadline is Wednesday, 11/27**
        *   [15 open Docs PR](https://docs.google.com/spreadsheets/d/1umeZ-AHjjD6ntbFv1J2dJOjWd0Ft0WKAV2Wtd8EQ8FM/edit#gid=0)
    *   **Release Notes**
        *   [Release theme and notes ](https://docs.google.com/document/d/1fL_xUwEWbxdRen-2mljio-Yd8Z2fzh2cQjaIOHRSHG0/edit)almost complete
        *   "[Known issue](https://github.com/kubernetes/kubernetes/issues/70955)" Github issue
        *   SIGs please update asap if not done already
    *   **1.13 Retro**
        *   Targeted for 12/6 during Community meeting
        *   [http://bit.ly/k8s113-retro](http://bit.ly/k8s113-retro) 
    *   **1.14 Release team **
        *   Leads nominated and chosen in SIG-Release meeting last week
        *   Shadows selection process 
        *   Use [GitHub issue](https://github.com/kubernetes/sig-release/issues/372) or [PR](https://github.com/kubernetes/sig-release/pull/377#discussion_r235575768) to express interest
*   [ 0:00 ] **SIG Updates**
    *   **SIG PM - **Stephen Augustus
*   [ 0:00 ] **?Announcements ?**
    *   Contributor Summit @ Kubecon
        *   Shanghai: Great turn out! Lots of great pics on Twitter etc.
        *   Seattle: **Chairs and owners, if you haven't confirmed we're running out of time, please let us know**. You do not need a ticket to kubecon/cnc for this. Email community@kubernetes.io
    *   Kubecon US is SOLD OUT. If you register now you'll be waitlisted
    *   Community Meeting Schedule - there are no SIG updates for December.
        *   12/6 - Release Retro for 1.13 (tentative!)
        *   12/13 - **Kubecon, no community meeting**
        *   12/20 and 12/27 - **No community meetings**
        *   January: SIG Apps,  SIG UI, SIG VMWare
*   **? **Shoutouts this week (Check in #shoutouts on slack) **?**
    *   _@justaugustus_ 
        *   Twitterverse shoutouts for our fearless 1.13 Release Team Lead, @AishSundar: [https://twitter.com/stephenaugustus/status/1063610123149545472](https://twitter.com/stephenaugustus/status/1063610123149545472) 
    *   _@cjwagner_ 
        *   Shoutout to @amerai for adding a search bar to Testgrid so that you don't have to dig to find the right dashboard! [https://testgrid.k8s.io/](https://testgrid.k8s.io/) 


## November 15, 2018 - ([recording](https://youtu.be/wkMRB1dalpA))



*   **Moderators**:  Jorge Castro [SIG Contribex]
*   **Note Taker**: Solly Ross (SIG Autoscaling/Google)
*   [ 0:00 ]**  Demo **--Pulumi - an OSS, k8s-native deployment orchestration engine [Alex Clemmer]
    *   Link to slides
    *   Link to repo
    *   Pulumi: open-source tools for managing cloud infrastructure
        *   Declare steady state (like Kubernetes) using programming languages like Python, Javascript, Typescript to manage the cloud repos
        *   Declare desired "outputs" to be saved for easy access (e.g. Service IPs)
        *   Schema is _exactly _kubernetes schema for Kubernetes types, etc (no special other format)
        *   Knows how interact with deployments (has concept of updates, knows that it needs to rollout, wait for rollout to succeed, only delete old objects after)
    *   Workflow for using something like RDS:
        *   Without Pulumi: deploy using one tool (e.g. terraform), then fetch connection string into secret (maybe using something else), then use in kubernetes (e.g. deploy app via Helm)
        *   With Pulumi: Declare steady state in code for everything
    *   Live demo: deploy CosmosDB + Helm chart
        *   Declare CosmosDB, exported connection string
        *   Declare secret (using normal Kubernetes schema) with connection string
        *   Declare Helm chart (deploying Bitnami Node.JS image) using secret to supply external DB
        *   How it works:
            *   Pulumi figures out dependencies automatically to figure out that CosmosDB needs to come before Secret, chart depends on secret
            *   `pulumi up` will show a "plan" of operations + Kubernetes JSON, executes plan on confirmation
        *   Can specify "stack outputs" to save from the objects generated, to fetch programmatically (e.g. IP of serving generated by Helm chart)
    *   Question
        *   How is schema generated, what happens to unknown attrs
            *   Schema is generated via OpenAPI spec based on all available versions
    *   
*   [ 0:00 ]** Release Updates**
    *   Current Release Development Cycle  [Aish Sundar - 1.13 Release Team Lead]
        *   **Code Freeze is tomorrow, 11/16 5pm PST**
            *   Only 1.13 critical fixes will be approved to go into master
            *   Master reopen 11/28 5pm PST for 1.14 
            *   Next **1.13.0-beta.1** release on Friday, 11/16 5pm PST (after code freeze)
            *   A few at risk items, release team is working to assess
        *   **CI Signal**
            *   [Master-blocking](https://k8s-testgrid.appspot.com/sig-release-master-blocking),[ master-upgrade](https://k8s-testgrid.appspot.com/sig-release-master-upgrade) and[ 1.13-blocking](https://k8s-testgrid.appspot.com/sig-release-1.13-blocking#Summary) are mostly Green, with intermittent flakes which could become blocking if they fail days before release 
            *   **Beta.1 Blocking Issues**
                *   [CSI Alpha Test Failures](https://github.com/kubernetes/kubernetes/issues/70760)
                *   [CSI Volumes [Driver: com.google.csi.gcepd] flakes](https://github.com/kubernetes/kubernetes/issues/65246)
                *   [GCE Upgrade jobs Timeout in 1.13 branch](https://github.com/kubernetes/kubernetes/issues/70912)
                *   Consider these to be priority!
        *   **Release Themes** are due by Wednesday 11/21
            *   Preliminary [draft of the Release notes](https://docs.google.com/document/d/1fL_xUwEWbxdRen-2mljio-Yd8Z2fzh2cQjaIOHRSHG0/edit)
            *   Leads to help fill out "Major Themes" section for your SIG
            *   For any "Known issue" please leave a comment of the issue and draft notes in this [GitHub Issue](https://github.com/kubernetes/kubernetes/issues/70955).
        *   **Kubernetes 1.14 Release Team Volunteers** - [https://github.com/kubernetes/sig-release/issues/372](https://github.com/kubernetes/sig-release/issues/372)
            *   Please consider volunteering :-)
*   Patch Release Updates
    *   v1.12.3 cut planned Monday, Nov 26th
    *   V1.10.0 published earlier this week
*   [ 0:00 ] **Contributor Tip of the Week **[Jeffrey Sica] 
    *   cs.k8s.io -- Search all repos in seconds
        *   Can _regex_ search across all Kubernetes repos and orgs
        *   Automatically filters out certain types of files, but that can be configured
    *   Check out the [contributor cheatsheet](https://github.com/kubernetes/community/blob/master/contributors/guide/contributor-cheatsheet.md) for other shortcuts (PRs accepted!)
*   [ 0:00 ] **SIG Updates**
    *   SIG Leads, check out this set of [recommended topics to cover](https://github.com/kubernetes/community/blob/master/events/community-meeting.md#sig-updates) during your update
    *   [Slide template if you need it](https://docs.google.com/presentation/d/1-nTvKCiqu9UvFYUeM6p6RIqHS5-H-u3_x-V4xj_eIWo/edit#slide=id.g401c104a3c_0_0)
    *   Please also check the [SIG Update Schedule](https://docs.google.com/spreadsheets/d/1adztrJ05mQ_cjatYSnvyiy85KjuI6-GuXsRsP-T2R3k)!
    *   SIG IBM Cloud [Richard Theis] (sig confirmed)
        *   Slides: [https://docs.google.com/presentation/d/11-CASs9mJl1RDwG9aN0UOYxIawNqIVtsMXXwyXWCS0Y/edit?usp=sharing](https://docs.google.com/presentation/d/11-CASs9mJl1RDwG9aN0UOYxIawNqIVtsMXXwyXWCS0Y/edit?usp=sharing) 
        *   Upcoming plans:
            *   Upstream IBM cloud provider repo (currently lives in IBM)
            *   Continue testing, bug fixing, etc
        *   [https://github.com/kubernetes/community/tree/master/sig-ibmcloud](https://github.com/kubernetes/community/tree/master/sig-ibmcloud)
    *   SIG Service Catalog [Carolyn Van Slyck?] (sig confirmed)
        *   Link to slides
*   [ 0:00 ] **?Announcements ?**
    *   Contributor Summit @ Kubecon
        *   Shanghai: Great turn out! Lots of great pics on Twitter etc.
        *   Seattle: **Chairs and owners, if you haven't confirmed we're running out of time, please let us know**. You do not need a ticket to kubecon/cnc for this. Email community@kubernetes.io
    *   Kubecon US is SOLD OUT. If you register now you'll be waitlisted. \

    *   Community Meeting Schedule - there are no SIG updates for December.
        *   11/22 (Thanksgiving in the US) - **Meeting is Still on**, Ihor will be your host! 
        *   12/6 - Release Retro for 1.13 (tentative!)
        *   12/13 - **Kubecon, no community meeting**
        *   12/20 and 12/27 - **No community meetings**
        *   January: SIG Apps,  SIG UI, SIG VMWare \

    *   Steering committee not having meeting in one weeks, will have one just before Kubecon
        *   **SIGs: Please try to have your charters in by KubeCon**
            *   Each SIG has been given a steering committee member to review
            *   reach out to steering committee if you have questions
    *   Office Hours next week on YouTube:
        *   See [this](https://github.com/kubernetes/community/blob/master/events/office-hours.md) for more information
        *   Come answer questions about Kubernetes on a livestream!
*   [ 0:00] **? Shoutouts this week (Check in #shoutouts on slack)** **?**
    *   paris - very big shoutout to @jberkus and the entire kubecon shanghai new contributor workshop team! josh built a team and carried out the event plan for this first time, sold out event in a new market to welcome contributors from this region. the event is in a few hours (from this timestamp) - best of luck and have a great time team!!
    *   jberkus - TY!  Let's add all the names: @tpepper @puja @nabrahams @xiangpengzhao @idealhack & Megan Lehn & Jerry Zhang
    *   neolit123 shoutout to @fabrizio.pandini for organizing the transition of phases in kubeadm to GA and also thank you to all the new kubeadm contributors who helped us with this work @yago @yuexiao wang @ereslibre @Rohit
    *   spiffxp Shoutouts to @chenopis @zacharysarah and @bradtopol for organizing and running the docs translation sprint at kubecon Shanghai
    *   Ivan Font @bentheelder and all others who worked on kind: I wanted to give a shout out for the work done to create kind. Nice work! I've experimented to get kind working with multiple clusters so that we can use it to test federation-v2 with multiple clusters for dev and CI and I'm very impressed with it so far! I filed a few issues #110, #111, #112, and #113 that I've stumbled across in the process of doing that, but it is not a reflection of the quality of work that's been done here. Again, thanks for the awesome work! Thanks to @munnerz @neolit123 @Jorgealarcon @Lion-Wei @TaoBeier @amwat


## November 8, 2018 - ([recording](https://youtu.be/NNHfryDY6mw))



*   **Moderators**:  Jorge Castro [SIG Contributor Experience]
*   **Note Taker**: Solly Ross (Google/SIG AUtoscaling)
*   [ 0:00 ]**  Demo **--IngressRoute with Contour - Steve Sloka ([steves@heptio.com](mailto:steves@heptio.com)) 
    *   Link to slides: [https://docs.google.com/presentation/d/1LAbRU7Fx7fofXolw0GckYw0AckdNk156yUdeBvgndkE/edit?usp=sharing](https://docs.google.com/presentation/d/1LAbRU7Fx7fofXolw0GckYw0AckdNk156yUdeBvgndkE/edit?usp=sharing)
    *   GitHub: [https://j.hept.io/contour](https://j.hept.io/contour)
*   [ 0:00 ]** Release Updates**
    *   Current Release Development Cycle  [Aish Sundar - Release Manager]
        *   We cut Beta0 and the [1.13 Release branch](https://testgrid.k8s.io/sig-release-1.13-blocking) yesterday, 11/7 !
        *   Updated to Go 1.11.2 before Beta.
        *   **Code Slush is this Friday (tomorrow), 11/9**
            *   All PRs need /priority, /kind, /sig, /milestone labels to merge post 5pm PST 
        *   **Code Freeze is just a week away - 11/16**
            *   Handful of Enhancements are pending only tests and docs
            *   There are a couple of Enhancements at risk and a few that have had no activity in past week or so.
            *   Owners **please update k/enhancement issues with current status** 
            *   Please reach out to Release team early on if you think you need to move out your enhancement
        *   **CI Signal**
            *   [Latest CI Signal Report](https://groups.google.com/forum/#!topic/kubernetes-dev/R7Jf96b2L-8) 
            *   Status is Yellow
            *   Kudos to Sig **Cluster-Lifecycle** for fixing long stating Upgrade setup issue [#56787](https://github.com/kubernetes/kubernetes/issues/56787), [#70627](https://github.com/kubernetes/kubernetes/issues/70627)
            *   We now have several new ones (in-progress though) that might become** blocking if not addressed soon**
                *   [CSI new Alpha tests failure](https://github.com/kubernetes/kubernetes/issues/70760)
                *   [Mounted Flex Volume expand failure](https://github.com/kubernetes/kubernetes/issues/70774)
                *   [GCE Serial test timeout](https://github.com/kubernetes/kubernetes/issues/70810)
                *   [Increased CPU usage in scale tests due to a Scheduler change ](https://github.com/kubernetes/kubernetes/issues/70708#issuecomment-436761955)
                *   Couple of HPA test fixes awaiting backport to 1.11 and 1.12. ([#70655](https://github.com/kubernetes/kubernetes/issues/70655), [#69444](https://github.com/kubernetes/kubernetes/issues/69444))
        *   **Docs**
            *   Open PRs: 11/22
            *   Completed PRs: 2/22 
            *   **We have 7 outstanding PRs**. We will be pinging owners on issues.
        *   **Release Notes**
            *   Sig Leads expect initial draft of the release notes coming your way for review next Monday 11/12
            *   _Please leave early feedback if you can_
        *   Questions:
            *   Where should we send 1.14 volunteers
                *   Look out for issue for more info, will be linked in next week's update 
    *   Patch Release Updates
        *   x.x
        *   y.x
*   [ 0:00 ] **Contributor Tip of the Week **[Aaron Crickenberger] 
    *   Let's talk about OWNERS, /lgtm and /approve
    *   How do OWNERS files work: [https://go.k8s.io/owners](https://go.k8s.io/owners)
        *   Loosely similar to CODEOWNERS files
        *   Used to both allow and suggest reviewers and approvers (and labels)
        *   Reviewers look for code quality/sane engineering (_/lgtm_)
            *   CANNOT self-LGTM
        *   Approvers look at "does this change make sense", "does it belong here", "does it clash with other functionality" (_/approve_)
            *   CAN self-approve
    *   How to use the commands:
        *   [https://prow.k8s.io/command-help#approve](https://prow.k8s.io/command-help#approve)
        *   [https://prow.k8s.io/command-help#lgtm](https://prow.k8s.io/command-help#lgtm)
    *   Workflow overview in slide form:
        *   [https://schd.ws/hosted_files/kccnceu18/88/kubecon-eu-2018-machines-can-do-the-work.pdf](https://schd.ws/hosted_files/kccnceu18/88/kubecon-eu-2018-machines-can-do-the-work.pdf) 
    *   BIKESHED TIME
        *   [https://github.com/kubernetes/test-infra/blob/master/prow/plugins.yaml#L39-L86](https://github.com/kubernetes/test-infra/blob/master/prow/plugins.yaml#L39-L86) 
*   [ 0:00 ] **SIG Updates**
    *   SIG Leads, check out this set of [recommended topics to cover](https://github.com/kubernetes/community/blob/master/events/community-meeting.md#sig-updates) during your update
    *   [Slide template if you need it](https://docs.google.com/presentation/d/1-nTvKCiqu9UvFYUeM6p6RIqHS5-H-u3_x-V4xj_eIWo/edit#slide=id.g401c104a3c_0_0)
    *   Please also check the [SIG Update Schedule](https://docs.google.com/spreadsheets/d/1adztrJ05mQ_cjatYSnvyiy85KjuI6-GuXsRsP-T2R3k)!
    *   SIG Cluster Lifecycle [Tim St. Clair] (confirmed)
        *   [Slides](https://docs.google.com/presentation/d/1HB9AmwmWHqAtyFeVlBKOYCrZjBk79W5tLWhZRRL6Amg/edit?usp=sharing)
        *   Homepage: [https://contributor.kubernetes.io/sigs/sig-cluster-lifecycle/](https://contributor.kubernetes.io/sigs/sig-cluster-lifecycle/)
        *   Mission: _simplify creation/upgrade/downgrade/teardown of Kubernetes clusters and their components_
        *   Last cycle
            *   Kubeadm: Config changes, improved CRI, HA, cert management, air-gapped support
            *   ClusterAPI: Provider-specific repos, many providers
        *   Upcoming plans
            *   Better meeting times and subproject coordination
            *   Kubeadm to GA (beta config, command line options fully supported)
            *   Cluster API integrating cluster API into kops
            *   Kubespray defaults to kubeadm
            *   ComponentConfig for structured configuration of other Kubernetes components
            *   **Deprecate**: kube-up, kubernetes-anywhere
        *   Reminder: for upgrade testing: _SIG cluster lifecycle owns the framework, others own the actual tests_
        *   Events: Upcoming kubecon talk on the future of addons
        *   New etcd management proposal/tooling (proposal just approved)
        *   Questions:
            *   If kube-up and kubernetes-anywhere are deprecated, what's the standard deployer (what's the kubeadm version of `curl | bash` to install)?
                *   documented well on the kubeadm section of the docs
                *   kubernetes-anywhere was just used for e2e tests
                *   other tooling (e.g. kubespray) builds on top of/orchestrates kubeadm, kubeadm only sees local machine
                *   Cluster API provides the same view of different backends, kubeadm is the base layer, providers in the middle
            *   Why use CRDs instead of aggregated APIs?
                *   maturity, easy of use, portability
    *   SIG OpenStack [Chris Hoge] (confirmed)
        *   [Slides](https://docs.google.com/presentation/d/1wJMj3MAfv2HGBH5xiB5uhZoTTCNjsEpcQPRslb4gpw4/edit?usp=sharing)
        *   Previous work
            *   In-tree driver is **deprecated** and will go away soon
            *   Manilla provision
            *   CSI support for Manilla and Cinder
            *   Magnum is now Kubernetes Certified Installer
            *   Driver for Cluster API in the works
        *   Future Work
            *   Heat & Senlin autoscaling drivers
            *   Storage driver consolidation
            *   Barbican driver for key management
            *   Finish in-tree code removal
        *   Transitioning into a WG under SIG Cloud Provider
        *   Events:
            *   OpenStack Summit, Berlin (Nov 13-15)
            *   Sessions at KubeCon Seattle
    *   SIG Auth [Mo Khan ] (confirmed) 
        *   [Slides](https://docs.google.com/presentation/d/14lhNK-h4W8W65jRtQtRKztnMygy7Qhk7NNm3CbyFmFM)
        *   Homepage: [https://contributor.kubernetes.io/sigs/sig-cluster-auth/](https://contributor.kubernetes.io/sigs/sig-auth/)
        *   Features
            *   Per-pod ephemeral service account tokens (projected volumes instead of secrets)
                *   **If NOT using client-go today**, need to keep reading token off disk
            *   Restricting Kubelet self-applied labels (via an admission plugin)
            *   Dynamic audit configuration (add/remove audit sinks without restart of API server)
        *   Container Identity WG winding down
*   [ 0:00 ] **?Announcements ?**
    *   Contributor Summit @ Kubecon
        *   Shanghai: Josh is getting on a plane, see you all there!
        *   Seattle: Chairs and owners, if you haven't confirmed we're running out of time, please let us know. 
    *   Community Meeting Schedule
        *   11/22 (Thanksgiving in the US) - Meeting is Still on, Ihor will be your host! 
        *   12/6 - Release Retro for 1.13 (tentative!)
        *   12/13 - **Kubecon, no community meeting**
        *   12/20 and 12/27 - **No community meetings**
        *   January, SIG Apps,  SIG UI, SIG VMWare
    *   Steering committee not having meeting in two weeks, will have one just before Kubecon
        *   SIGs: Please try to have your charters in by KubeCon

        **? **Shoutouts this week (Check in #shoutouts on slack) **?**

    *   paris and jdumars: big thanks to @mattfarina who just spent an hour helping organize our project boards
    *   Jberkus: Shanghai Shoutouts for next week: Megan Lehn for doing all the logistics and legwork from thousands of km away, @puja @xiangpengzhao and @idealhack for translating all the New Contributor Summit materials and many other things besides! Also, to @mrbobbytables and our localization volunteers for getting the international forums at dicuss.kubernetes.io launched!
    *   AishSundar: Shoutout to @jberkus and his team of CI Signal shadows @maria and @mortent for staying on top of CI signal failures and flakes every day, opening and following up on test issues and fixes and help maintain stable test health for 1.13 release !
    *   jdumars - Big thanks to @spiffxp @dims and @mattfarina ? all of whom have stepped up and helped with the work in SIG Architecture!
    *   mauilion - shoutout to @jdetiber for always finding time to help dig into the cluster-api stuffs.
    *   AishSundar - shoutout to @justinsb, yet again, for extremely quick turnaround on a long standing Upgrade testing issue (#56787). This helped us get clean e2e CI coverage one of the 1.13 Beta Feature "Taint Based Evcitions"


## November 1 , 2018 - (recording)



*   **Moderators**:  Tim Pepper [SIG Release / Contrib Ex]
*   **Note Taker**: Solly Ross
*   [ 0:00 ]**  Demo **--  Automation Broker - Michael Hrivnak ([mhrivnak@redhat.com](mailto:mhrivnak@redhat.com)) (confirmed)
    *   [Website](http://automationbroker.io/)
    *   Service Bundle: container image/pod that runs to completion to install a service on the cluster
        *   Hooks into service catalog via _automation broker_
        *   **A**nsible **P**laybook **B**undle
            *   Easy way to make a service bundle
            *   Each service catalog action maps to an Ansible playbook in the bundle
        *   Can run other things besides ansible in service bundles (demo on youtube running Helm)
        *   Service catalog UI support
            *   Partial support in Kubeapps
            *   Support in OpenShift
    *   Ansible Operator
        *   Runs ansible roles/playbooks as an operator
    *   Ansible roles exist for manipulating kubernetes objects
*   [ 0:12 ]** Release Updates**
    *   Current Release Development Cycle  [Aish Sundar - Release Manager]
        *   [v1.13-alpha3](https://github.com/kubernetes/kubernetes/releases/tag/v1.13.0-alpha.3) was cut yesterday, 10/31
        *   **v1.13-beta0 and Release branch creation scheduled for Tuesday, 11/6**
            *   Highly dependant on clean CI Signal
            *   Branch fast forwards will happen everyday thereafter
        *   **Code slush is coming up Friday, 11/9**
            *   Enhancement owners evaluate enhancement readiness based on pending work (code, test and docs) 
            *   **Code Freeze** is just 2 weeks away !
            *   If you need to enhancement adjusted, please work with the Release team
            *   Ensure 1.13 PRs are uptodate on labels (sig, kind, priority, milestone)
            *   **_Tide will start enforcing [Code slush merge label requirements](https://github.com/kubernetes/community/blob/master/contributors/devel/release.md#tldr) _**
        *   **CI Signal**
            *   [This week report](https://groups.google.com/forum/#!topic/kubernetes-dev/_7aSN1agv_4)
            *   Kudos to SIG** Autoscaling** and** Cluster-Lifecycle **for closing out some long standing failing tests ([69444](https://github.com/kubernetes/kubernetes/issues/69444), [70058](https://github.com/kubernetes/kubernetes/issues/70058)) !
            *   **<span style="text-decoration:underline;">Beta Blocking tests</span>**
                *   [[sig-scheduling] Scheduler priorities](https://github.com/kubernetes/kubernetes/issues/69989)
            *   [GKE Upgrade Fail](https://github.com/kubernetes/kubernetes/issues/70445). Identified as GKE specific issue, Fix expected to rollout to OSS today (Not a blocker !)
            *   **Shuffling jobs in release blocking dashboards**
                *   We have finalized the[ criteria for Blocking Tests](https://github.com/kubernetes/sig-release/pull/346)
                *   As first step we are moving GKE upgrade tests from [sig-release-master-upgrade](https://k8s-testgrid.appspot.com/sig-release-master-upgrade#Summary&show-stale-tests=) to [sig-release-master-upgrade-optional](https://k8s-testgrid.appspot.com/sig-release-master-upgrade-optional#Summary&show-stale-tests=) dashboard - [PR#9959](https://github.com/kubernetes/test-infra/pull/9959)
                *   Expect to see more jobs moved out of "release-blocking" to "release-informing" dashboard 
    *   Patch Release Updates
        *   1.11.4 went out last week
        *   1.12.2 went out last week
    *   Questions
        *   Are GKE tests considered blockers?
            *   _Correction from CI Signal: they are right now, but expect a proposal to remove the GKE Upgrade tests from blocking before beta, since those tests have been chronically flaky.  Expect a PR to comment on soon._
            *   _https://groups.google.com/forum/#!topic/kubernetes-dev/8Po230FeEIs_
*   [ 0:00 ] **Open KEPs** - [link](https://groups.google.com/d/topic/kubernetes-dev/LIkZoIqCT20/discussion) to Caleb's announcement..repository is moving location
    *   SIG Architecture is working to improve KEP process for community
    *   Extracting KEPs from community repo (see link to the [discussion on kubernetes-dev](https://groups.google.com/d/topic/kubernetes-dev/LIkZoIqCT20/discussion))
    *   Try to have small merges that document consensus rather than waiting for full approval/finalization to merge
    *   Moving towards eventually making KEPs the main way to propose features (as opposed to being optional)
*   [ 0:00 ] **SIG Updates**
    *   SIG Leads, check out this set of [recommended topics to cover](https://github.com/kubernetes/community/blob/master/events/community-meeting.md#sig-updates) during your update
    *   [Slide template if you need it](https://docs.google.com/presentation/d/1-nTvKCiqu9UvFYUeM6p6RIqHS5-H-u3_x-V4xj_eIWo/edit#slide=id.g401c104a3c_0_0)
    *   Please also check the [SIG Update Schedule](https://docs.google.com/spreadsheets/d/1adztrJ05mQ_cjatYSnvyiy85KjuI6-GuXsRsP-T2R3k)!
    *   SIG AWS [Nishi Davidson] (confirmed)
        *   [https://drive.google.com/file/d/1jDR1Esdu2ApnuLrzsGxn7iv1cU3sWc7R/view](https://drive.google.com/file/d/1jDR1Esdu2ApnuLrzsGxn7iv1cU3sWc7R/view) 
        *   We currently host 5 subprojects in SIG-AWS
        *   Subprojects aws-alb-ingress-controller, aws-ebs-csi-driver and out-of-tree ccm will be alpha in k8s v1.13
        *   Cloud Provider status
            *   In-tree
                *   Adding e2e tests
                *   Will be maintained until out-of-tree is GA, 2 release deprecation
            *   Out-of-tree
                *   GA Q3 2019
        *   CI Signal
            *   Added aws-tester plugin
            *   Creates ephemeral EKS cluster to run Kubernetes e2e tests as periodic jobs (not blocking)
            *   Hoping to integrate etcd conformance tests, cluster API tests as well
    *   [SIG Scheduling [Bobby Salamat] (confirmed)](https://docs.google.com/document/d/1Ztwf24XXR6S13pFBth_W86pNWk20b5Kv7mkcFAZCgJI/edit?usp=sharing)
        *   1.12
            *   Scheduler perf improvements
                *   Only score percentage (50%, but configurable) of feasible nodes per pod, properly considered across failure domains
                *   Improved affinity/anti-affinity performance
            *   Graduated TainNodesByCondiion to beta, which creates taints for node conditions automatically
            *   Enable ImageLocalityFunction by default, which prefer nodes which already have the images for a pod, weight set to avoid putting all pods from an RS on the same node
            *   Scheduling framework design finalized (move scheduler features towards plugins, both in-process and out-of-process)
        *   1.13
            *   Finalize design of gang/co-scheduling (more efficient batch job scheduling, e.g. for ML workloads)
            *   Finalize pod scheduling policies (allowing admins to control how pods get scheduled -- e.g. preventing setting tolerations, preventing certain namespaces from getting placed on certain nodes)
            *   **Deprecating the "critical pod" annotation**, in favor of pod priority and preemption
            *   Enable pod resource limit function (prefer nodes that can fit both a pod's request _and_ limit)
            *   Implement extension points for scheduling framework (see above)
            *   Improve equivalence cache (new design to address existing shortcomings)
    *   SIG Contributor Experience [Paris Pittman] (confirmed)
        *   [Update Deck](https://docs.google.com/presentation/d/1jLL5_nEKAHuhcqNOUgXhNJlgNK45gxklAROGzD8Dg8s/edit?usp=sharing)
        *   What was done last cycle
            *   Theme: making your life easier (automation, documentation, mentoring, events, etc)
            *   Performed the contributor survey (graphs on the way!)
                *   Common comments
                    *   Meetups are out of scope, but will pass information on to CNCF
                    *   You can apply "good first issue" labels even if you didn't file the issue
                    *   People liked slack, release team notes in community meetings
                *   Scrubbed data is in the link, take a look
            *   Misc
                *   Communication moderation changes (stay public while dealing with bad actors) -- **SIG chairs should learn how to follow these processes**
                    *   Calendar is private ATM because of bad actors
                    *   Zoom links aren't publicly posted for similar reasons (_please don't tweet them_), but work is being done to solve this with Zoom
                *   Launched discuss.k8s.io as a community forum, please post/take a look!
        *   Upcoming
            *   Revamp developer guide
            *   Move KEPs out kubernetes/kubernetes
            *   Build a contributor site
            *   Upgraded communications guide
            *   Improve SIG Chair processes (e.g. Zoom-to-Youtube automation process) -- **please reach out if you have opinions**
        *   Seattle Contributor Workshop
            *   Waitlisted (**if you're a chair, TL, or subproject owner who hasn't signed up, please reach out!**)
            *   lots of good content planned
            *   Night-before event to hang out and talk
        *   **Consider mentoring, even if it's just 1 hour per quarter**
            *   Only need one merged PR to be a mentor
        *   See slides for a _whole lot _more work, information, links, and sigup information
*   [ 0:00 ] **?Announcements ?**
    *   [KEPs are moving out of k/community](https://groups.google.com/forum/#!topic/kubernetes-dev/LIkZoIqCT20)** **(calebamiles)
*   Shoutouts
    *   Nikhita: shoutout to @dims for being Asia/EU friendly while deciding the meeting time for #k8s-infra-team
    *   Mzee1000: Shout-out to @mrbobbytables for his help with Kubernetes 101 in Bangalore
    *   Jberkus: to: @justinsb for continuing to be the "difficult test fail" resolver.
    *   Jberkus: to @neolit123 for fast turnaround on kubeadm test fails
    *   Fejta: shoutout to @bentheelder for finally creating a @thockin emoji :thockin:
    *   @liz to: @bentheelder for going above and beyond to help me get my KIND tests working!
    *   @paris thanks to @nikhita @roycaihw @brendanburns @dims and many others for answering questions from first time contributors in the outreachy process slack channel #outreachy-apps
    *   @spiffxp thanks to @audreylim for tackling e2e test error messages as her first kubernetes pull-request (https://github.com/kubernetes/kubernetes/pull/69583)
    *   To Solly Ross for taking notes today


## October 25, 2018 - ([recording](https://www.youtube.com/watch?v=RGT5V9kC4y4))



*   **?Moderator**:  Jorge Castro [SIG Contributor Experience]
*   **?Note Taker**: Josh Berkus [Red Hat/SIG Contributor Experience/Release]
*   [ 0:00 ]** Release Updates**
    *   Current Release Development Cycle  [1.13]
        *   [v1.13-alpha2](https://github.com/kubernetes/kubernetes/releases/tag/v1.13.0-alpha.2) was cut yesterday, 10/24
        *   We are currently in Enhancement freeze with [34 tracked enhancements](https://github.com/kubernetes/enhancements/issues?page=1&q=is%3Aissue+is%3Aopen+milestone%3Av1.13+label%3Atracked%2Fyes&utf8=%E2%9C%93) 
        *   Any further additions need to go through a formal [exception process](https://github.com/kubernetes/enhancements/blob/master/EXCEPTIONS.md)
        *   **3 weeks** to Code Freeze (11/16)
        *   Please open **test, docs PRs** and also add about **release notes** to PRs when applicable
        *   **CI Signal  **
            *   **[Test Failure Report for 10/23](https://groups.google.com/forum/#!topic/kubernetes-dev/_jILqo17Suw)** 
            *   Kudos to SIG** Cluster-Lifecycle** and** Apps **for closing out some long standing failing tests ([69475](https://github.com/kubernetes/kubernetes/issues/69475), [69356](https://github.com/kubernetes/kubernetes/issues/69356)) !!
            *   New [Active Spreadsheet](https://docs.google.com/spreadsheets/d/1TfECf8uSVnHaaCn8KTWK-sMdG3PV7jtW7FggHIUpeUU/edit#gid=781125942) for tracking test failures.  Still in development, but will be more up-to-date than the CI reports.
            *   **<span style="text-decoration:underline;">Beta blocking test failures (11/6)</span>**
                *   **Master-blocking **
                    *   [Issue #69444](https://github.com/kubernetes/kubernetes/issues/69444) -  [Horizontal Pod Autoscaling failures](https://k8s-testgrid.appspot.com/sig-release-master-blocking#gce-cos-master-serial). It is now critical-urgent.
                    *   [Issue #70058](https://github.com/kubernetes/kubernetes/issues/70058) - Kubeadm failures
                    *   [Issue #69989](https://github.com/kubernetes/kubernetes/issues/69989) - Pod SchedulerPriorities
                *   **Upgrade Test Board **
                    *   Same issues tracked in Master blocking
                    *   Tests are [generally flaky](https://github.com/kubernetes/kubernetes/issues/70151). We continue monitor for legit failure signal, alongside deflaking attempts.
    *   Patch Release Updates
        *   [1.12.2 planned for today, 10/25](https://groups.google.com/forum/#!topic/kubernetes-dev/br_xm72nSvo/discussion)
        *   Y.x
*   [ 0:00 ]**  Demo **-- Cluster API AWS Provider ([chuck@heptio.com](mailto:chuck@heptio.com))
    *   [Link to repo](https://github.com/kubernetes-sigs/cluster-api-provider-aws)
    *   Link to slides or docs or whatever goes here. 
    *   Demo of using the Cluster API to provision AWS.
    *   CLI tool: clusterawsadm
        *   Creates IAM rules, etc.
    *   Must already have SSH key pair (does not create)
    *   Starts with an existing Kubernetes cluster to create more clusters; you have to create a 1.11 or later cluster on your own (could be minikube)
    *   Create manifests for the CRDs, using "makemanifest"
    *   Clusterctl crd then controls the cluster.
        *   Pass many parameters by switch
        *   Once the new cluster is created, moves the ClusterAPI to that cluster.
    *   Secrets?  In the CRD defintions.
    *   Config file for ClusterCTL?   Not sure.
*   [ 0:10 ] ?**Contributor Tip of the Week** [Aaron Crickenberger] ?
    *   [HODL](https://i.ytimg.com/vi/c56jGlymb0E/maxresdefault.jpg)
    *   [https://prow.k8s.io/command-help#hold](https://prow.k8s.io/command-help#hold)
        *   Prevents merging
    *   [http://go.k8s.io/github-labels#do-not-merge/hold](http://go.k8s.io/github-labels#do-not-merge/hold) 
    *   /hold to add, /hold cancel to remove
    *   [Good idea / Bad idea](https://cdn-images-1.medium.com/max/814/1*VM4QgVeRrlBxGyvk4F1snw.png)
        *   Good idea: explaining why you're putting on the hold
        *   Bad idea: removing a hold in a PR you're not involved in
    *   Reasons to hold
        *   Hang on, I the reviewer, think this needs more discussion
        *   I, the author, am holding this and will remove it when I've heard from the people I want
        *   I think the author should have final say on when this PR merges
    *   Notes:
        *   Anybody can add or remove a hold (don't even need to be an org member)
        *   Can we blacklist?  We can, from the org
            *   We'd have to restrict hold to org members if it was a problem
*   [ 0:20 ] ?**SIG Updates**?
    *   SIG Leads, check out this set of [recommended topics to cover](https://github.com/kubernetes/community/blob/master/events/community-meeting.md#sig-updates) during your update
    *   [Slide template if you need it](https://docs.google.com/presentation/d/1-nTvKCiqu9UvFYUeM6p6RIqHS5-H-u3_x-V4xj_eIWo/edit#slide=id.g401c104a3c_0_0)
    *   Please also check the [SIG Update Schedule](https://docs.google.com/spreadsheets/d/1adztrJ05mQ_cjatYSnvyiy85KjuI6-GuXsRsP-T2R3k)!
    *   SIG Docs [Jennifer Rondeau, Zach Corleissen] (confirmed)
        *   [Slides](https://docs.google.com/presentation/d/1YNnITGt9-1A-_uFCjXWhAqqEYi-8HM7pFFnG8yEYG5g/edit?usp=sharing)
        *   Thanks to zparnold for adding automation foo tracking doc submistions in the Github API
        *   Reorganized localization, now under a consolidated repo
            *   Everything in kubenetes.io
            *   Thanks to Korean translators for making this happen
            *   Also updated guidelines
        *    Better automation for API reference docs (thanks Chi Ming Tang(sp?))
        *   Jennifer Rondeau is new SIG-Docs co-chair
        *   
    *   Upcoming Work:
        *   Renewing our focus on content and quality
            *   Focusing especially on onboarding 
        *   Assessing how to remove user journeys and remodel the onboarding experience
        *   Better onboarding docs through graphical models
            *   Dedicated resource from Dominic Tornow (SAP) to generate models
            *   [https://medium.com/@dominik.tornow/kubernetes-high-availability-d2c9cbbdd864](https://medium.com/@dominik.tornow/kubernetes-high-availability-d2c9cbbdd864) 
        *   Tim Fogarty (@tfogo) is docs meister for 1.13

        Upcoming doc sprints:

    *   Shanghai (localization workflows)
    *   Seattle (TBD)
*   Next: how do we ensure that content remains fresh?
*   Localization subprojects have been consolidated into k/website
*   Now have a WG for SIG-Docs tooling, led by Luc Perkins(sp?)
*   Want to contribute?  We *always* need technical reviewers!
    *   Or just pick an open issue
    *   PRs get more attention than issues, so if you find something wrong, PR a correction.
*   Chairs: Andrew Chen, Zach Corliessen, Jennifer Rondeau
*   SIG Storage [Saad Ali] (confirmed) ?
*   [Slides](https://docs.google.com/presentation/d/11nTnl549maTV-XUBQ_66_t17Nnal56NiTeo6SAIl810/edit?usp=sharing)
*   Last Quarter:
    *   Topology Aware Volume Scheduling
        *   Make scheduler smarter about where storage is
        *   Used to be a per-storage hack, now an expressible constraint for the scheduler
        *   Started in 1.10, added CSI support in 1.12, beta soon
    *   Snapshot & Restore
        *   Started a year ago.  Was a question whether it should be part of the API at all.  But many DB admins would like it.
        *   Mapping declarative to imperative was hard.
    *   CSI to GA this quarter
        *   PV support is primary
        *   Also want to support ephemeral volumes, and block volumes
        *   Now, we need to migrate the in-tree storage to CSI plugins.  This is a blocker for Cloud Provider migration.
            *   Challenge: end-users need to have a smooth transition
        *   Working on reusable libraries for common storage (iSCSI etc.) that can be used as templates
        *   Adding conformance testing for CSI
        *   GA depends on completing Kubelet registration mechanism
        *   We're extrating the mount library in k/k to a separate repo so that CSI driver authors can use it.
    *   To beta this quarter:
        *   Ephemeral volumes
        *   CSI Topology
        *   CRD automated installation
        *   In-tree Block Volume support
*   Catch up with SIG-Storage at their biweekly meeting, or at Kubecon
    *   Sessions in Seattle, also a "Cloud-native Storage Day"
*   As part of moving to GA, they need a more robust mechanism than user-modifiable CRDs.  THis includes addressing the downgrade problem.
*   [ 0:00 ] **?Announcements ?**
    *   Meet Our Contributors - Nov 7th at 230p and 8p UTC 
        *   230pm UTC - 5 Steering Committee Members AMA
        *   8pm UTC - contributor mentors AMA
        *   #meet-our-contributors on slack
        *   [YouTube Playlist](https://www.youtube.com/playlist?list=PL69nYSiGNLP3QpQrhZq_sLYo77BVKv09F)
    *   The final call! CNCF awards nominations are open, details [here](https://www.cncf.io/blog/2018/10/08/annual-cncf-community-awards-nominations-kick-off-winners-to-be-recognized-at-kubecon-cloudnativecon-seattle/)!
    *   Kubernetes Contributor Summit Details
        *   The **Contributor Social for Kubecon Shanghai** has been scheduled.  It will be from 5pm to 7pm, November 13, at the convention center.  The event will feature a panel of Chinese contributors to Kubernetes, discussing obstacles and opportunities.
        *   Seattle - Registration is closed, waitlist is in effect, if you cannot attend please let #contributor-summit (Paris/Jorge/Bob) know so we can free up your slot!  

        **? **Shoutouts this week (Check in #shoutouts on slack) **?**

        *   pwittrock - Shoutout to @alexismp @jeefy and @mrbobbytables for helping me with my All Things Open Kubebuilder workshop.  Thank you so much!
        *   AishSundar - shoutout to @nikopen for automating the Issues and PR spreadsheet for Bug Triage and CI Signal for 1.13 ! Thanks for staying on top of this and accommodating the feature requests
        *   jberkus - to @justinsb for splitting out our long-running upgrade tests so that they actually complete
        *   spiffxp - Congrats to @bentheelder for creating a PR that deletes over 3 million lines of code [https://github.com/kubernetes-sigs/kustomize/pull/503](https://github.com/kubernetes-sigs/kustomize/pull/503)
            *   Shoutouts to @ixdy for setting things up so we can use shorter URL's https://testgrid.k8s.io and [https://gubernator.k8s.io](https://gubernator.k8s.io)
        *   nabrahams - Shoutout to @idealhack for translating a huge pile of slides in preparation for the New Contributor Workshop in Shanghai.
        *   kacole2 - Shoutout to @AishSundar @spiffxp @claurence @gsaenger and @ameukam for their help on getting all the k/features (enhancements) issues in a great spot where everything is now being tracked to a PR in k/k and getting the freeze over the finish line.
        *   AishSundar - @kacole2 right back at you ! Awesome job on doing all the heavy lifting yourself. Staying on top of ~50 incoming enhancements, following up to prune the list and mentoring the shadows at the same time is no easy feat :slightly_smiling_face:
        *   nikhita - Shoutout to @lukaszgryglicki for being extremely responsive to feature requests for DevStats and implementing them and fixing bugs reallyyyyyyy fast!! :tada:
    *   [Stackoverflow Top Users](https://stackoverflow.com/tags/kubernetes/topusers) (Once a month at the end of the month)
        *   [Rico, Praveen Sripati, ](https://stackoverflow.com/users/2989261/rico)Ij[az Khan, Ryan Dawson, samhain1138, VonC, Michael Hausenblas, David Maze, Ignacio Mill?n, Konstantin Vustin](https://stackoverflow.com/users/4550110/ijaz-khan)
    *   Community meeting Nov 22
        *   We traditionally cancel this meeting due to US Holiday (Thanksgiving)
        *   Let's try to be more global, see #sig-contribex if you want to help drive this meeting this week while the US is out.


## October 18, 2018 - ([recording](https://youtu.be/_oOX6OuPZaM))



*   **Moderators**:  Aaron Crickenberger [SIG ~~Beard~~ Testing]
*   **Note Taker**: First Last [Company/SIG]
*   [ 0:00 ]**  Demo **-- Kubetest [Marco Ceppi, marco@ceppi.net] (confirmed)
    *   [Link to slides](https://docs.google.com/presentation/d/1MjUPBihKHhpAFskOcpwafTWehTrGrSxCNtIBPfFcDC0/edit#slide=id.p)
    *   [Link to repositories](https://github.com/vapor-ware/kubetest)
*   [ 0:00 ]** Release Updates**
    *   Current Release Development Cycle  [Aish Sundar ~ Release 1.13 Lead]
        *   **[v1.13-alpha.1](https://github.com/kubernetes/kubernetes/releases/tag/v1.13.0-alpha.1) **was cut on Monday 10/15
        *   alpha.2 slated for Tuesday 10/23
        *   **Enhancement Freeze is next Tuesday - 10/23**. 
            *   [37 enhancements](https://github.com/kubernetes/features/issues?page=1&q=is%3Aissue+is%3Aopen+milestone%3Av1.13+label%3Atracked%2Fyes&utf8=%E2%9C%93) currently tracked for 1.13
            *   Owners please update your issues in k/features repo to indicate confidence level of the committed work
            *   Provide a list of pending PRs and issues for each enhancement to lohelp us track better
            *   SIG Arch will be reviewing the list of planned enhancements in 1.13
        *   **CI Signal  **
            *   **[Test Failure Report for 10/16](https://groups.google.com/forum/#!topic/kubernetes-dev/MXzAw_GH30Q)** - Net open issues remain constant
            *   Kudos to SIG **Scalability, Scheduling, Node, API-Machinery **for fixing failing tests ([69473](https://github.com/kubernetes/kubernetes/issues/69473), [69597](https://github.com/kubernetes/kubernetes/issues/69597), [69786](https://github.com/kubernetes/kubernetes/issues/69786), [69815](https://github.com/kubernetes/kubernetes/issues/69815))
            *   **<span style="text-decoration:underline;">Beta blocking test failures (11/6)</span>**
                *   **Upgrade Test Board is Red**
                    1. [Issue #69475](https://github.com/kubernetes/kubernetes/issues/69475)  - [gce-new-master-upgrade-cluster-new](https://k8s-testgrid.appspot.com/sig-release-master-upgrade#gce-new-master-upgrade-cluster-new) and [gke-gci-new-gci-master-upgrade-cluster-new](https://k8s-testgrid.appspot.com/sig-release-master-upgrade#gke-gci-new-gci-master-upgrade-cluster-new) were failing for weeks earlier. Splitting Jobs has given us new signal !
                    2. [Issue #69356](https://github.com/kubernetes/kubernetes/issues/69356) - **Daemonset** failure is pending resolution
                *   **Master-blocking **
                    3. [Issue #69891](https://github.com/kubernetes/kubernetes/issues/69891) -  Primarily [GKE tests timing out](https://github.com/kubernetes/kubernetes/issues/69891)
    *   Patch Release Updates
        *   1.12.x
        *   1.11.x - @foxish planning for next week sometime.
        *   1.10.x
*   [ 0:00 ] **Graph o' the Week **[Aaron Crickenberger]
    *   Hi I'd like to propose we retire this slot [https://github.com/kubernetes/community/issues/2818](https://github.com/kubernetes/community/issues/2818) 
    *   But one last thing before we do? <seinfeld>What's the deal with repository groups?</seinfeld>
    *   [https://github.com/cncf/devstats/pull/145](https://github.com/cncf/devstats/pull/145) - let's use sigs.yaml instead
    *   Compare and contrast for yourselves
        *   [The prod server has the old repository groups](https://k8s.devstats.cncf.io/d/12/dashboards?refresh=15m&orgId=1)
        *   [The test server has the new repository groups by sig](https://k8s.teststats.cncf.io/d/12/dashboards?refresh=15m&orgId=1)
*   [ 0:00 ] **Contributor Tips** [Aaron Crickenberger]
    *   Hi welcome to what I propose we replace Graph/KEP o' the Week with [https://github.com/kubernetes/community/issues/2818](https://github.com/kubernetes/community/issues/2818)
    *   Let's talk about /help and /good-first-issue
    *   https://prow.k8s.io/command-help#help
*   [ 0:00 ] **SIG Updates**
    *   SIG Leads, check out this set of [recommended topics to cover](https://github.com/kubernetes/community/blob/master/events/community-meeting.md#sig-updates) during your update
    *   [Slide template if you need it](https://docs.google.com/presentation/d/1-nTvKCiqu9UvFYUeM6p6RIqHS5-H-u3_x-V4xj_eIWo/edit#slide=id.g401c104a3c_0_0)
    *   Please also check the [SIG Update Schedule](https://docs.google.com/spreadsheets/d/1adztrJ05mQ_cjatYSnvyiy85KjuI6-GuXsRsP-T2R3k)!
    *   ~~SIG PM [Stephen Augustus] (confirmed)~~
        *   ~~No show this week~~
    *   SIG Testing [Aaron Crickenberger] (confirmed)
        *   [https://docs.google.com/presentation/d/1TD7Z07LlJDprD7m5K4bv_PpYZ-rLNOptmoU3Pab-gY8/edit#slide=id.g439b3a360b_0_25](https://docs.google.com/presentation/d/1TD7Z07LlJDprD7m5K4bv_PpYZ-rLNOptmoU3Pab-gY8/edit#slide=id.g439b3a360b_0_25) 
    *   SIG Instrumentation [Frederic Branczyk] (confirmed)
*   [ 0:00 ] **Announcements**
    *   Want to host the community meeting? Come talk to us on #sig-contribex on slack!
    *   Tickets are really running low on the [Contributor Summit registration for Kubecon](https://github.com/kubernetes/community/tree/master/events/2018/12-contributor-summit). aka last call! 
    *   [Kubernetes Outreachy Internships are live. ](https://www.outreachy.org/communities/cfp/kubernetes/)
        *   There are three: two with client libraries and one for a revamp dev guide. Brendan Burns, Nikhita Raghunath, and Paris Pittman are mentors and program owners. Reach out with questions. #outreachy-apps is now open for all Outreachy questions. 
        *   Interested applicants need to go through the [Outreachy](https://www.outreachy.org/) application process FIRST and then they will be engaged in the rest of the process. 
        *   They will be required to make a first time contribution as part of the application process. Please try to use the good-first-issue where possible or you will see an increase in spelling error PRs.
    *   ?NCF Awards Nominations are open! Please nominate folks - [https://www.cncf.io/blog/2018/10/08/annual-cncf-community-awards-nominations-kick-off-winners-to-be-recognized-at-kubecon-cloudnativecon-seattle/](https://www.cncf.io/blog/2018/10/08/annual-cncf-community-awards-nominations-kick-off-winners-to-be-recognized-at-kubecon-cloudnativecon-seattle/) 
    *   Shoutouts this week (Check in #shoutouts on slack) (or we'll do it live!)
        *   Shoutouts from @jdumars to #sig-testing for all the work they do
        *   Shouts from @guineversaenger to @cjwagner and @fejta for being helpful and responsive to all questions testing related
        *   Shoutouts to @pohly and @matthyx from @spiffxp for all of the great work they've been contributing to #sig-testing lately


## October 11, 2018 - ([recording](https://youtu.be/MBDOLG5OTpQ))



*   **Moderators**: Paris Pittman [SIG-ContribEx]
*   **Note Taker**: ??? 
*   [ 0:00 ] **Demo: **Encrypted Container Images for k8s
    *   Brandon from IBM to upload slides, URLs, materials, etc.
        *   Presentation: [https://ibm.box.com/s/ctzxk6iu25ras975zwu2h4eijqk9ljmh](https://ibm.box.com/s/ctzxk6iu25ras975zwu2h4eijqk9ljmh)
        *   Design Doc: [https://docs.google.com/document/d/146Eaj7_r1B0Q_2KylVHbXhxcuogsnlSbqjwGTORB8iw/edit](https://docs.google.com/document/d/146Eaj7_r1B0Q_2KylVHbXhxcuogsnlSbqjwGTORB8iw/edit) 
        *   Discussion Issue: [https://github.com/opencontainers/image-spec/issues/747](https://github.com/opencontainers/image-spec/issues/747)
        *   Dev Branch: [https://github.com/stefanberger/containerd/tree/image-encryption.v3](https://github.com/stefanberger/containerd/tree/image-encryption.v3)
*   [ 0:00 ]** Release Updates** 
    *   1.13 release
    *   **Enhancement **collection underway since Oct 8th
        *   Currently [37 issues](https://github.com/kubernetes/features/issues?utf8=?&q=is%3Aissue+is%3Aopen+milestone%3Av1.13+label%3Atracked%2Fyes+) under k/features repo. Please keep the labels (kind, priority, sig) up-to-todate on the issues. Please also indicate level of confidence and what's pending in terms of code, test and docs.
        *   Enhancement Freeze - 10/23
        *   For large / risky enhancements, please try to land by early November so we can assess stability from CI signal and get any bug fixes in by Code freeze (Nov 15th)
    *   **CI Signal - Test failures**
        *   **Except weekly CI Signal report to k-devs@**
        *   [First CI Signal Report](https://groups.google.com/forum/#!topic/kubernetes-dev/yc3OUIA3ybk) went out yesterday.
        *   If there are failing-test issues assigned to your SIG please treat it as priority !
    *   1.13 schedule adjusted to accommodate 2 Alphas (10/15 and 10/23) and 2 RC builds (11/27 and 11/30)
    *   This should help us test out the build and release tools and flows roughly once in 2 weeks during the release
    *   Patch Update:
        *   planning 1.10.9 next tuesday (16th)
*   [ 0:00 ] **Graph of the Week** [Josh Berkus]
    *   [Countries stats](https://k8s.devstats.cncf.io/d/50/countries-stats?orgId=1)
        *   [Contributors from China, Korea, and Japan](https://k8s.devstats.cncf.io/d/50/countries-stats?orgId=1&var-period_name=Quarter&var-countries=%22China%22&var-countries=%22Japan%22&var-countries=%22Korea,%20Republic%20of%22&var-repogroup_name=All&var-metric=contributors&var-cum=countries)
    *   [Timezone Stats](https://k8s.devstats.cncf.io/d/51/timezones-stats?orgId=1)
        *   [Timezones for PR Authors to Networking](https://k8s.devstats.cncf.io/d/51/timezones-stats?orgId=1&from=now-90d&to=now&var-period=w&var-tzs=All&var-repogroup_name=Networking&var-metric=prcreators)
*   [ 0:00 ] **SIG Updates**
    *   Network [casey confirmed]
        *   [https://docs.google.com/presentation/d/1AkTp3m6LC8B7UqKPT3mWIyo5DMu6gdLvexFB7WIVCPI/edit?usp=sharing](https://docs.google.com/presentation/d/1AkTp3m6LC8B7UqKPT3mWIyo5DMu6gdLvexFB7WIVCPI/edit?usp=sharing)
*   [ 0:00 ] **?Announcements ?**
    *   KubeCon/CloudNativeCon Shanghai New Contributor Workshop is sold out!** ** ?
        *   We will have a waiting list up soon
        *   Related to this: we have #cn-dev on Slack and will have a Chinese section of discuss.kubernetes.io soon.
    *   Seattle Contributor Summit [Jorge] - this is your low ticket warning. RSVP to a track and watch as [we build out the schedule](https://github.com/kubernetes/community/tree/master/events/2018/12-contributor-summit).
    *   CNCF Awards - nominations are open! Please, check out the blog post (announcement) for details - [https://www.cncf.io/blog/2018/10/08/annual-cncf-community-awards-nominations-kick-off-winners-to-be-recognized-at-kubecon-cloudnativecon-seattle/](https://www.cncf.io/blog/2018/10/08/annual-cncf-community-awards-nominations-kick-off-winners-to-be-recognized-at-kubecon-cloudnativecon-seattle/) 
    *   [Kubernetes Office Hours](https://github.com/kubernetes/community/blob/master/events/office-hours.md) [Jorge] is next week.
        *   We're short two folks forth both the European and Western editions again this week. If you can help ping @jorge on slack. 
    *   Meet Our Contributors is also looking for more mentors to sit in on a one hour monthly Q&A session. [Watch @sttts do a super quick code base tour.](https://youtu.be/yqB_le-N6EE)(clip of the show) [paris]
        *   230p and 8p UTC; first Weds of the month
    *   **?Shoutouts this week (Check in #shoutouts on slack) ?**
        *   M1kola: Shoutout to @soltysh for re-organising kubectl codebase in [https://github.com/kubernetes/kubernetes/pull/6946](https://github.com/kubernetes/kubernetes/pull/6946). I think, this change improves contributor's experience. Now I spend less time waiting for tests & coverage report when I want to check only specific kubectl command. Significant productivity boost for me and hopefully for other people too.
        *   Leigh: **big** ups to @neolit123 for putting in so much hard work to keep the flurry of sig-cluster-lifecycle work organized -- thank you, Lubomir!


## October 4, 2018 - ([recording](https://www.youtube.com/watch?v=QyFEPSJEO48))



*   **Moderators**: Jeffrey Sica [SIG-UI]
*   **Note Taker**: ??? 
*   [ 0:00 ] **?Announcements ?**
    *   [ 0:00 ]** Release Updates** 
        *   1.13 release cycle started this Monday, 10/1: see [https://github.com/kubernetes/sig-release/tree/master/releases/release-1.13](https://github.com/kubernetes/sig-release/tree/master/releases/release-1.13) for detail schedule
            *   Short release (10 weeks)
            *   Code freeze on Nov 15th
            *   Anticipated release on 12/3
            *   SIGs and leads, please plan 1.13 enhancements load accordingly for example prioritizing stability and bug fixing ahead of features. 
            *   Two key items known currently:
                *   removal of etcd2
                *   Move to go 1.11.1
        *   [1.12.1 planned for tomorrow October 5, 2018](https://groups.google.com/d/topic/kubernetes-dev/h_SoRAyzmDg/discussion)
    *   ?Steering Committee Election Announcement [Brendan Burns] ?
        *   Aaron @spiffxp
        *   Tim St Clair @timothysc
        *   Davanum Srinivas @dims
    *   Contributor Summit Seattle - You don't need to sign up for KubeCon first. This will sell out in the next few weeks. Please RSVP now. [https://git.k8s.io/community/events/2018/12-contributor-summit](https://git.k8s.io/community/events/2018/12-contributor-summit)
    *   ?Shoutouts this week (Check in #shoutouts on slack) ?
        *   bentheelder - Shoutout to @mrhohn @neolit123 @justinsb and @liggitt for each helping me in turn to track down https://github.com/kubernetes/kubernetes/issues/69195 This one was tricky to pin down! :sweat_smile:
        *   paris - shouts to @justaugustus @krzyzacy @mrbobbytables and @jeefy (filing in for @jorge) for being mentors in our first episode of our October #meet-our-contributors :k8s: :heart_eyes_k8s:
        *   paris - shouts to @brendanburns @timothysc and @spiffxp for spending time with us answering questions on #meet-our-contributors this week :k8s: :heart_eyes_k8s: 
*   [ 0:00 ] **Release v1.12 Retrospective:** link [http://bit.ly/k8s112-retro](http://bit.ly/k8s112-retro) 


## September 27, 2018 - recording



*   **Moderators**:  Josh Berkus [SIG-Release]
*   **Note Taker**: Tim Pepper [VMware/SIG-Release]
*   [ 0:00 ]** Release Updates**
    *   Current Release Development Cycle  [Tim Pepper ~ Release Lead]
        *   Release today!   ...expect published artifacts by ~5pm Pacific
        *   Pengfei Ni / @feisky (slack) /  [@feisker](https://github.com/feiskyer) (github) is 1.12.y patch manager.
        *   Any issues need targeted for cherry pick to 1.12.1 asap.
        *   Retro next week.  Enter any notes or thoughts or desires relative to release process in the [1.12 retro doc](https://docs.google.com/document/d/1OgylAYqU0YoJz-PTd8uzyHtMcxYSewSq06AGeh1F-A8/edit#heading=h.tw06ll716grh).
    *   Patch Release Updates
        *   No 1.10, 1.11 updates
        *   1.9.11 Mehdy Bohlool
            *   released today, <span style="text-decoration:underline;">final update</span> to 1.9.  
            *   "Finish Him!"
            *   Time to upgrade if you haven't...1.10, 1.11, 1.12 will now be our active releases.
*   [ 0:00 ] **Graph o' the Week **[Aaron Crickenberger]  How to analyze your test failures...
    *   Triage dashboard: [https://go.k8s.io/triage](https://go.k8s.io/triage) clusters for analysis the past week's worth of failures.
        *   A bunch of BigQuery data is created from GCS buckets.   (This is a publicly accessible BigQuery data set, custom analysis charges your personal Google project.)
        *   Various interactive UI elements allow tunneling into more details.  Can search GitHub for already existing issue(s) referencing an issue you're looking at.  Clicking on a failed test takes you to gubernator logs, which include at the top a command to reproduce the specific test job.  
        *   In the past automation tried to open issues, in hopes that they would get fixed, but that hasn't worked well...needs human interaction and inspection.  Toggles allow looking at CI, PR, or both.
        *   See video for live examples of how/where to click on things in the UI for chasing details of a specific failure.
    *   Triage update job: [https://k8s-testgrid.appspot.com/sig-testing-misc#triage](https://k8s-testgrid.appspot.com/sig-testing-misc#triage)
        *   Configurable notifications via email if more than N failures happen
    *   Triage codebase: [https://github.com/kubernetes/test-infra/tree/master/triage](https://github.com/kubernetes/test-infra/tree/master/triage)
    *   Troubleshooting issue: [https://github.com/kubernetes/test-infra/issues/9271](https://github.com/kubernetes/test-infra/issues/9271)
    *   BigQuery Ingress rate: [http://velodrome.k8s.io/dashboard/db/bigquery-metrics?panelId=12&fullscreen&orgId=1](http://velodrome.k8s.io/dashboard/db/bigquery-metrics?panelId=12&fullscreen&orgId=1)
    *   BigQuery freshness: [https://k8s-testgrid.appspot.com/sig-testing-misc#metrics-kettle](https://k8s-testgrid.appspot.com/sig-testing-misc#metrics-kettle)    
*   [ 0:00 ] **SIG Updates**
    *   SIG Leads, check out this set of [recommended topics to cover](https://github.com/kubernetes/community/blob/master/events/community-meeting.md#sig-updates) during your update
    *   Please also check the [SIG Update Schedule](https://docs.google.com/spreadsheets/d/1adztrJ05mQ_cjatYSnvyiy85KjuI6-GuXsRsP-T2R3k)!
        *   SIG leadership should prepare ahead of time
        *   Give notice to the meeting moderator if your SIG will be unavailable
    *   ~~SIG Azure [Stephen Augustus] (confirmed)~~ - Sorry peeps, last minute conflict today that I can't shuffle
    *   SIG Big Data [Yinan Li] (confirmed)
        *   [Slides](https://docs.google.com/presentation/d/1GO8T23fx8NnXW-8QqzvYwXJFfOB5hNuq_hE4dUd84dw)
        *   Apache Spark 2.4 Release coming mid to late October with Kubernetes functionality:
            *   python and R language support
            *   Client mode for spark-shell and notebooks
            *   Mounting volumes, emptyDri, hostPath, PVC
            *   Executor memory request control improvements (fractional values, milli-CPUs)
            *   Scheduler backend robustness improvements
        *   Future work:
            *   Pod templates for driver and executor pod customizations.  There are so many options configurable...explosion.  Rather than adding more, they're adding template support to make usage easier.
            *   Service shuffling
            *   Kerberos authentication
            *   Improving support for local application dependencies on the client machines doing submissions
            *   Driver resilience, checkpoint/restart for streaming applications.
        *   Non-Spark, other projects:
            *   Airflow 1.10: new operator and executor for arbitrary pods to run tasks.  Operator for life cycle management (newly open sourced)
            *   Spark Operator: exporting jvm/driver/executor metrics to Prometheus
    *   SIG Scalability [Shyam Jeedigunta] (tentative)
        *   Governance: charter is in place, giving official guidance on what the SIG can and can't do (eg: in the past there's been a lot of question around what is and is not release blocking)
        *   Test resources: github issue (link?) exists to track work items ahead of shifting underlying test resources to CNCF
        *   SLOs: API call latency and pod startup latency have been the only two for quite some time.  Looking to add multiple new ones for networking.
        *   Cluster loader ([link](https://github.com/kubernetes/perf-tests/blob/master/clusterloader/docs/design.md)): have achieved minimal viable product (MVP) level of readiness.  This should enable other teams to easily run scale tests themselves.
        *   CI health / tests: Working to speed up pre-submits' running time, as these have historically been the slowest and are a merge bottleneck (eg: kubemark 500 is waay faster).  Investing in de-flaking scale jobs like the 5000 node test, which recently had 8 green runs in a row (a record?).
        *   1.12 release: as typical, close to the end of the release some scalability issues came up, this time in:
            *   taint node by condition feature which lead to notably slower start up.
            *   CoreDNS regression was masked by other regressions until very late in the release cycle, leaving insufficient time to debug safely and feature dropped (CoreDNS is not the default on GCE yet..1.13?)
*   [ 0:32 ] **?Announcements ?**
    *   ?Steering Committee Election Announcement ? 
        *   less than one week to vote! Election closes October 3rd @ 6pm PT. 
        *   Check email for CIVS ballot email titled: "Poll: Kubernetes Steering Committee Election for 2018". 
        *   Voters Guide [https://git.k8s.io/community/events/elections/2018 ](https://git.k8s.io/community/events/elections/2018)
        *   229 of 681 people have voted! We had 309 votes last year.
        *   Mail [community@kubernetes.io](mailto:community@kubernetes.io) if you have any issues/concerns. 
    *   [https://www.surveymonkey.com/r/k8s-contributor-2018](https://www.surveymonkey.com/r/k8s-contributor-2018) -> Contributor Experience survey that will shape our direction and close feedback loops. K8s water bottle for your time today!
    *   SIG Chairs/Leads/TLs/Keepers of Zoom and YouTube - check your email for issues relating to those two services. Approx. 40% of our community are not using the correct zoom license.  Many, many meetings missing from YouTube channel, which is a negative on transparency.
    *   Contributor Summit (Kubecon / CloudNativeCon Seattle) - registration will say waitlist, but: **We are not sold out, yet. Please sign up anyway**. We are working through a registration issue. There is a possibility we will pull registration from the KubeCon site into a separate registration process. If you have signed up already, you will not need to sign up again.  If you are waitlisted, you will be contacted to sort out this issue. Join #contributor-summit on slack for real time info and [GH LINK](https://github.com/kubernetes/community/tree/master/events/2018/12-contributor-summit) for published information. Initial content will be listed next week.
        *   [Draft Agenda for Content](https://docs.google.com/document/d/17StTsUSCh1XxPjF-TpCnhXvZalHs7lwhSrVIgNpFwLU/edit) - still looking for feedback on the sessions/content you'd like to see. 
    *   ?Shoutouts this week (Check in #shoutouts on slack) ?
        *   @dims: Big shoutout to @jonasrosland to getting the CNCF meetup off the ground in boston and @abe for his talk on The Kubernetes Release Cycle ( https://www.meetup.com/Cloud-Native-Computing-Boston/ )
        *   @timothysc: Huge shoutout to the whole release team and everyone else whose put in crazy effort to make RC2!
        *   @aish: HUGE shoutout to @justaugustus for helping us fill the [release 1.13] roster out !
        *   @justaugustus Shoutout to everyone who volunteered for the 1.13 Release Team! We staffed a _FULL_ roster of leads and shadows in 16 days and they'll be in super capable hands with @AishSundar and @spiffxp at the helm!
        *   @spiffxp: shoutout to @cjwagner (Cole Wagner) for all of the work he's done over the past year to remove mungegithub from the project and bring tide to kubernetes/kubernetes


## Sep 20, 2018 - ([recording](https://youtu.be/HqePVxsMPDw))



*   **Moderators**:  Jonas Rosland [SIG-ContribEx]
*   **Note Taker**: Josh Berkus, Jaice Singer DuMars and Jorge Castro [SIG-ContribEx]
*   [ 0:01 ]** Release Updates**
    *   Current Release Development Cycle  [Tim Pepper ~ 1.12 Release lead]
        *   We're down to _probably_ 1 issue pending fix and test soak
        *   Key upcoming events barring any new test issues:
            *   Friday Sept 21 "cherry-pick deadline" (non-event given master hasn't yet thawed)
            *   Friday Sept 21 cut RC2 and built rpms/debs
            *   Monday Sept 24 last release-1.12 branch fast-forward from master branch
            *   Monday Sept 24 thaw master branch
            *   ...final soak: cherry picks only for absolutely critical show stopper bugs?
            *   **Thursday Sept 27 release**
    *   Patch Release Updates
        *   1.9.10 released Aug 3
        *   1.10.8 released Sept 15
        *   1.11.3 released Sept 11
*   [ 0:03 ] **SIG Updates**
    *   SIG Cloud Provider [Chris Hoge] (confirmed)
        *   [https://docs.google.com/presentation/d/186rAa3cNCBOA2GBmFdNvBI_Ko4inaQYHEX2CFA4TSqs/edit#slide=id.p](https://docs.google.com/presentation/d/186rAa3cNCBOA2GBmFdNvBI_Ko4inaQYHEX2CFA4TSqs/edit#slide=id.p)
        *   Looking at cloud providers in core code, like Google, AWS, etc.
            *   Don't want more in the upstream code
            *   Want level playing field
            *   In the process of moving providers to plugins
            *   Minimum requirements to add a provider, is it documented? Do they post results to testgrid? So that users have some assurance that they'll have a positive experience. 
        *   In 1.12:
            *   Added new providers to SIG-CP: Alibaba. DigitalOcean and Baidu in progress. 
            *   Started work on removing the in-tree providers
                *   [https://zoom.us/j/9941605205](https://zoom.us/j/9941605205)
                *   [https://docs.google.com/document/d/1nBwl3BmF4IOZwxakJyePJA9b-Gme7al2RyXWy4JNCMc/edit?usp=sharing](https://docs.google.com/document/d/1nBwl3BmF4IOZwxakJyePJA9b-Gme7al2RyXWy4JNCMc/edit?usp=sharing)
                *   Created CP Extraction Working Group, [meetings 11am](https://zoom.us/j/9941605205 ) Pacific Thursdays.
                *   Many in-tree providers use internal APIs, so code needs to be refactored.
                *   Some APIs need to be moved to staging
        *   In 1.13:
            *   Continue work to move in-tree providers
            *   Document external provider usage, need to make sure all requirements etc. are documented.  Installation process is more complicated, for example.
            *   Maybe moving provider SIGs into SIG-CP, likely to take longer.
            *   Make sure all providers have conformance test results
        *   Collaboration
            *   SIG-Docs doesn't want to be in charge of provider docs, that should be up to SIG-CP
            *   Working with Cluster Lifecycle on install/upgrade
        *   Q: should provider repos be in the Kubernetes org?
            *   A: the decision on whether providers should be part of the kubernetes org is an Arch decision, it's a question of what you think Kubernetes is.  We're trying to provide a level playing field.  Right now, they are part of the org.
    *   SIG Architecture [Jaice Singer DuMars] (confirmed)
        *   What we do
            *   Manage and maintain architectural consistency over time
            *   Manage subprojects:
                *   [API review process](https://github.com/kubernetes/community/blob/master/sig-architecture/api-review-process.md) 
                *   [KEP reviews](https://github.com/kubernetes-sigs/architecture-tracking/projects/2) (architectural lens) 
                *   [Conformance test](https://github.com/kubernetes-sigs/architecture-tracking/projects/1) definition and implementation
                *   Charter in progress due to complex processes we manage, so there are not many defined subprojects yet. 
            *   Manage policy and governance
                *   API governance (guidance docs and the review process)
                *   Deprecation policy
                *   Code organization
                *   KEP process
                *   General issues around Kubernetes scope
        *   What are we working on right now?
            *   Take a look at our [meeting notes](https://docs.google.com/document/d/1BlmHq5uPyBUDlppYqAAzslVbAO8hilgjqZUTaNXUhKM/edit)
            *   Our [YouTube Channel](https://www.youtube.com/playlist?list=PL69nYSiGNLP2m6198LaLN6YahX7EEac5g)
    *   SIG API Machinery [Daniel Smith] (confirmed)
        *   Dry run in alpha to see what the predicted outcomes of an action will be
            *   You can test API to see what it looks like when it runs through the webhooks
        *   CRD versioning change
            *   No schema change allowed
            *   Register a webhook to do a schema change
            *   All the various API definitions should have the same feature set
        *   SSH tunnels are going away - this has been deprecated for a year
        *   Re: server side apply
            *   In a feature branch
                *   Allows writing code during freeze
            *   More complete designs on the way
        *   "We don't own your API" - we're not the API reviewers
            *   We do own _some APIs_ like metadata format, CRD API, webhook interface APIs, comms between aggregator & APIs, controllers, RBAC API, controller shell (informer, reflector, shared informers, etc.), controller manager binary
        *   Upcoming:
            *   rate limits, flow control - prevent API quota over-consumption ~ de facto prioritization of API requests
            *   Internal API Server coordination - e.g. how do you know when every API is serving the same version of a CRD
            *   SIG meeting & agenda ~ if it's empty the day before, the meeting will be cancelled (there's an "agenda closed" meeting event)
*   [ 0:00 ] **Announcements**
    *   Calendar information - subscribe, don't copy! [Jonas]
        *   [Instructions here](https://github.com/kubernetes/community/blob/master/events/community-meeting.md)
        *   Mac Calendar:

<p id="gdcalert1" ><span style="color: red; font-weight: bold">>>>>>  gd2md-html alert: inline image link here (to images/Kubernetes-Community0.png). Store image on your image server and adjust path/filename if necessary. </span><br>(<a href="#">Back to top</a>)(<a href="#gdcalert2">Next alert</a>)<br><span style="color: red; font-weight: bold">>>>>> </span></p>


![alt_text](images/Kubernetes-Community0.png "image_tooltip")

        *   Google Calendar:

<p id="gdcalert2" ><span style="color: red; font-weight: bold">>>>>>  gd2md-html alert: inline image link here (to images/Kubernetes-Community1.png). Store image on your image server and adjust path/filename if necessary. </span><br>(<a href="#">Back to top</a>)(<a href="#gdcalert3">Next alert</a>)<br><span style="color: red; font-weight: bold">>>>>> </span></p>


![alt_text](images/Kubernetes-Community1.png "image_tooltip")

    *   Election Update [Jorge]
        *   [All the info you need](https://groups.google.com/forum/#!topic/kubernetes-dev/0gEdp_xdzEI)
        *   Total voters: 677
        *   Actual votes cast thus far: 144
        *   If you are in [voters.md](https://github.com/kubernetes/community/blob/master/events/elections/2018/voters.md) but have not received a ballot please mail [community@kubernetes.io](mailto:community@kubernetes.io)
        *   Election ends on October 3.
*   Kubernetes Contributor Summit is happening [Jorge] 
    *   [Information](https://github.com/kubernetes/community/tree/master/events/2018/12-contributor-summit)
    *   [Content is in draft](https://docs.google.com/document/d/17StTsUSCh1XxPjF-TpCnhXvZalHs7lwhSrVIgNpFwLU/edit), SIG leads and TLs, please review and comment
    *   Register as part of your Kubecon registration (check the box for collocated events); you will get a follow up email about RSVPing for tracks and

                Sunday dinner and fun @ garage in Seattle 

        *   Contributor Social Shanghai: 11/13, 6pm, at the convention center
            *   Will have panel on contributing from China
*   Contributor Survey [Jorge]
    *   [https://www.surveymonkey.com/r/k8s-contributor-2018](https://www.surveymonkey.com/r/k8s-contributor-2018)
*   [Jorge] October 1st marks the start of Hacktoberfest, a month-long celebration of open source software. This is an opportunity to welcome new Kubernetes contributors to the community. Please help by making an extra effort this month to add more issues with the good-first-issue label. If you have any large tasks that could use help from a lot of contributors, now would be the perfect time to create an issue for it. You can learn more about Hacktoberfest here: [https://hacktoberfest.digitalocean.com/](https://hacktoberfest.digitalocean.com/)
    *   ContribEx will be generating a 404 report for new users so they will have a place to go.
    *   SIGs, consider updating your good-first-issue labels. 
    *   SIG Leads, we've put together some recommendations for how to give an [update for this meeting](https://github.com/kubernetes/community/blob/master/events/community-meeting.md#sig-updates), the host will be reminding you from now on before your update. 
    *   Last call for Outreachy intern requests! 
*   #Shoutouts!_ (want to say thanks? Use the #shoutouts channel in slack)_
        *   @vlad Shlosberg: Huge shoutout to @jorge, @mrbobbytables, @paris, @hubt and a bunch of other for helping make @Foqal a success. Working with me on feedback, helping promote the project, submitting helpful answers, and everything else you guys have done!
        *   @bentheelder: Shoutout to @mrhohn for being eternally responsive to networking issues on everything from PR reviews to sig-network test configs, dns images, the network e2es, and answering questions related to network issues in the infra and helping debug! Thanks for helping get the kube-dns manifest images out the door for 1.12! Zihong is always fixing things for us over in #sig-testing :slightly_smiling_face: and now over in #sig-release
        *   @nikhita: shout out to @carolynvs for creating lots of help-wanted issues on service catalog ([https://github.com/kubernetes-incubator/service-catalog/issues?q=is%3Aissue+is%3Aopen+label%3A%22help+wanted%22](https://github.com/kubernetes-incubator/service-catalog/issues?q=is%3Aissue+is%3Aopen+label%3A%22help+wanted%22))....and having an excellent strategy of doing that! - [https://twitter.com/carolynvs/status/1042061098101485580](https://twitter.com/carolynvs/status/1042061098101485580)
        *   @dims: big shoutout to @fabrizio.pandini for testing all the things! (v1.12 RC1 kubeadm under various scenarios)


## Sep 13, 2018 - ([recording](https://youtu.be/FYJsqqCWRs4))



*   **Moderators**:  Arun Gupta [SIG AWS/Amazon]
*   **Note Taker**: Solly Ross [SIG Autoscaling]
*   [ 0:00 ]**  Demo **-- Answering questions on k8s Slack w/ Foqal [Vlad Shlosberg, [vlad@foqal.io](mailto:vlad@foqal.io)] (confirmed)
    *   [https://docs.google.com/presentation/d/19RNjayF59WanE8Q9ug4sftFXniGQP4PRRXsRC4X7dd4](https://docs.google.com/presentation/d/19RNjayF59WanE8Q9ug4sftFXniGQP4PRRXsRC4X7dd4)
    *   [https://foqal.io/oss](https://foqal.io/oss)
    *   Goals
        *   Improve UX
        *   Focus Contributor Times
    *   Core Idea
        *   Automatically respond to common questions without any special interaction
    *   Functionality
        *   Upon asking a question (without special syntax), Foqal sends answer, marked as just to you
        *   Can rate question, if marked as helpful, the answer is sent to entire channel
    *   Sources
        *   StackOverflow
        *   Docs (divided into small sections)
        *   Slack conversations
            *   Upon detecting question, looks for answers sent afterwards
                *   Sends message to answerer, asking if it's appropriate to store
                *   Can edit answers before storing them
    *   Results
        *   3 months, 2 active channels, 37 helpful autoresponses in past 2 weeks
        *   Slack conversations and Kubernetes docs provide most useful answers
    *   Currently talking to docs folks to use Foqal responses to improve docs content, searchability, and examples
    *   **Invite Foqal bot to your channel in Kube slack**
        *   `/invite @Foqal`
        *   Both SIG channels and more user-facing channels
        *   Add context before storing
        *   Can manually add to Foqal using the elipsis meu on any slack message
    *   Talk to Foqal about...
        *   importing other docs sources
        *   Partitioning (SIG meeting times might not be useful to kubernetes user channels)
        *   Ask Vlad if you have questions
    *   Can also run on private Slack instances
*   [ 0:00 ]** Release Updates**
    *   Current Release Development Cycle  [Tim Pepper ~ 1.12 Release lead]
        *   Still in Code Freeze.  See [here](https://github.com/kubernetes/community/blob/master/contributors/devel/release.md#tldr) for TLDR what do I do to get a merge.
        *   Beta 2			-  Sept. 11
        *   RC			-  Sept. 18
        *   Release target 	- ** Sept. 25: AT RISK**
            *   We are making progress on CI Signal, but slowly.
            *   Depending on today/tomorrow improvements merging and test results showing up by Monday Sept. 17
            *   ...**potential to delay release toward Sept. 27**
        *   Tide: we moved k/k to it on Monday. Worked through a few minor issues.  Seems to be working reasonably now.
*   [ 0:00 ] **SIG Updates**
    *   SIG Windows [Michael Michael] (confirmed)
        *   Finished a bunch of functionality required for moving to stable
            *   Not moving to stable until 1.13, due to conformance, perf, stability hiccups
        *   Want to finalize docs, how-to guides, etc for GA
        *   Stopping feature development to focus on stabilization
    *   SIG Node [Dawn Chen] (confirmed)
        *   Slide: [https://docs.google.com/presentation/d/1G034FTqXeXO5Gf1H-ufTkAMgJKOcx6HCIzB6krkO6zY/edit?usp=sharing](https://docs.google.com/presentation/d/1G034FTqXeXO5Gf1H-ufTkAMgJKOcx6HCIzB6krkO6zY/edit?usp=sharing)
        *   Finished charter
            *   Meetings weekly Tuesday at 10AM PT, Resource Management WG Wednesday 11AM PT, on-demand meetings for Asia times
            *   Revised/categorized SIG scope (see slide 3, large list)
        *   Recent work
            *   Sandbox Pods
                *   RuntimeClass proposal, alpha feature, CRD
                *   Working to integrate with Kata and other sandbox solutions, containerd shim
            *   Windows Container Support (with SIG Windows)
                *   GA in 1.13
                *   Kubelet stats for Windows system containers
                *   Fixes for network, eviction manager bugs
                *   In-review PRs for
                    *   DNS capabilities for Windows CNI (with sig-network)
                    *   Windows CNI support (with sig-network)
                    *   Testing frameworks (with sig-testing)
            *   Testing
                *   Changes in Node E2E (see slide 6 for link)
                    *   Reorganized tests to more easily track results
                    *   New tests need to be tagged to run in normal test suites
                *   CRI Testing dashboard (see slide 6 for link)
                    *   One place to view node conformance test results and features for CRI implementations
            *   Misc
                *   User NS support in progress
                *   ResourceClass API under discussion (beyond just GPU support)
                *   Efficient heartbeat for scalability in progress
                *   PID NS sharing in beta
                *   Updated debug container API, accepted proposal, implementation in progress
*   [ 0:00 ] **Announcements**
    *   Steering Committee Election update: [paris/ihor/jorge]
        *   **Tomorrow! Is the deadline for all nominations** (entire process including bios uploaded) and [voter eligibility forms](https://www.surveymonkey.com/r/k8s-sc-election-2018) (if you are not on [voters.md](https://github.com/kubernetes/community/blob/master/events/elections/2018/voters.md) and want to vote).
            *   Voter eligibility is normally based on contributions in the past 12 months, but you can make a request to be added if you've made non-GitHub contributions and you think you should be eligible
        *   Next? **CIVS polling ballots go out on Wednesday, September 19th** to emails we have on file. If you do not receive an email by Thursday (please check spam/bulk), contact community@kubernetes.io. We will remind everyone on this call next week as well as our regular channels (k-dev ML, discuss.k8s.io, slack, etc.)
    *   #Shoutouts!_ (want to say thanks? Use the #shoutouts channel in slack)_
        *   @Mzee1000: Shout-out to @AishSundar and @gsaenger for incredible help with CI signal
        *   @AishSundar: Huge shoutout to @gsaenger for lighting up the right fires when and where needed for 1.12 !! Way to go
        *   @Justaugustus: Shoutout to @dougm, @dims, @bentheelder, @sttts, and anyone I might've missed for working the weekend to test our Release Engineering tooling ahead of the next beta cut!
        *   @misty: @lucperkins for adding per-heading anchor links to the docs so people can share an in-page section at any level, without having to go back to the TOC to find the link!
        *   @neolit123: thanks to @timothysc and @fabrizio.pandini who helper with debugging a release blocking e2e test for sig-cluster-lifecycle!
        *   @mkumatag: Now we have `v1.12.0-beta.2` release images are all fat manifest.. This made all other architectures first class citizens.. Thanks @dims @dougm @ixdy @luxas @calebamiles @tpepper @bentheelder 
        *   @paris: shout to @ameukam for helping contribex with our communication platform discovery and doing the hard work. perfect example of chopping wood and carrying water.
        *   @tpepper: huge shout out to @bentheelder for working late late last night and right back to it this morning on diagnosing/resolving build pipeline issues in support of 1.12 release


## September 6, 2018 - ([recording](https://youtu.be/xTEBQjnKLi4))



*   **Moderators**:  Jorge Castro [SIG-Contribex / Heptio]
*   **Note Taker**: First Last [Company/SIG]
*   [ 0:00 ]**  Demo **--  [Kubernetes GSoC 2018 project demo - etcdproxy-controller](https://github.com/xmudrii/etcdproxy-controller) - [marko@loodse.com](mailto:marko@loodse.com)
    *   Project URL: [https://github.com/xmudrii/etcdproxy-controller](https://github.com/xmudrii/etcdproxy-controller)
    *   [Google Slides](https://docs.google.com/presentation/d/1KSqGj3AwpFDKdRLk678zSXkdTvmjzut6scZhSwLr0gw/edit?usp=sharing)
*   [ 0:00 ]** Release Updates**
    *   Current Release Development Cycle  [Tim Pepper ~ 1.12 Release lead]
        *   Code Freeze arrived this week.  See [here](https://github.com/kubernetes/community/blob/master/contributors/devel/release.md#tldr) for TLDR what do I do to get a merge
        *   Release target Sept. 25
        *   Release at risk currently due to poor CI Signal ([sig-release-1.12-all](https://k8s-testgrid.appspot.com/sig-release-1.12-all) & [sig-release-1.12-blocking](https://k8s-testgrid.appspot.com/sig-release-1.12-blocking) & [sig-release-master-blocking](https://k8s-testgrid.appspot.com/sig-release-master-blocking) & [sig-release-master-upgrade](https://k8s-testgrid.appspot.com/sig-release-master-upgrade)).  Considering kicking out all GKE blocking tests.  Also significant GCE issues.  Blurs the signal when the hosting is unstable.
        *   Tide...looking to move to it on Monday. 
        *   Patch Release Updates
            *   X.x
            *   Y.x
*   **Moving from Submit Queue to Tide** [spiffxp]
    *   Tide is used on the [majority of our repos](https://github.com/kubernetes/test-infra/issues/6227) now, kubernetes/kubernetes is the only repo using submit-queue.  We plan on changing that next week.
    *   Stop talking SIG Beard, just show me the docs because...
        *   [I just want my PR to merge](https://github.com/kubernetes/test-infra/blob/master/prow/cmd/tide/pr-authors.md) ([using whatever labels the release team tells me to](https://github.com/kubernetes/community/blob/master/contributors/devel/release.md#tldr))
        *   [I just want my repo added to tide](https://github.com/kubernetes/test-infra/blob/master/prow/cmd/tide/config.md)
    *   Ways that tide differs from the submit queue
        *   Tide uses GitHub queries to select PR's into "tide pools", runs as many in a batch as it can ("tide comes in"), and merges them ("tide goes out")
        *   Since there is no queue, there is no [ordering determined by one of 8 different things](https://submit-queue.k8s.io/#/info), ie: applying queue/fix has no effect
        *   Tide will always rerun tests prior to merge, we have an [open issue around ways to support merging without tests](https://github.com/kubernetes/test-infra/issues/5334), ie: applying retest-not-required has no effect
    *   If you used the submit queue UI
        *   [https://submit-queue.k8s.io/#/prs](https://submit-queue.k8s.io/#/prs) -> [https://prow.k8s.io/pr](https://prow.k8s.io/pr) 
        *   [https://submit-queue.k8s.io/#/queue](https://submit-queue.k8s.io/#/queue) -> [https://prow.k8s.io/tide.html](https://prow.k8s.io/tide.html) 
        *   [https://submit-queue.k8s.io/#/history](https://submit-queue.k8s.io/#/history) -> n/a
        *   [https://submit-queue.k8s.io/#/ci](https://submit-queue.k8s.io/#/ci) -> [http://velodrome.k8s.io/dashboard/db/monitoring?orgId=1](http://velodrome.k8s.io/dashboard/db/monitoring?orgId=1)
    *   Rollout plan
        *   [Move all munger functionality out of mungegithub](https://github.com/kubernetes/test-infra/issues/3331)
        *   [Create tracking issue](https://github.com/kubernetes/test-infra/issues/3866)
        *   Propose to release team
        *   Propose to contribex
        *   Propose to community ? we are here
        *   Send something like this out to kubernetes-dev
        *   Put together a PR with queries for k/k master and release branches
        *   Turn down misc-munger and submit-queue instances Monday, remove cherrypick-queue munger from cherrypick instance, check in with release team during daily burndown meetings
        *   If we need to rollback, we will revert any PR's that errantly merged, and turn up the misc-munger and submit-queue instances again
*   **Graph o' the Week **[jberkus]
    *   Re-organized Devstats
    *   Fewer charts, [main dashboard](https://k8s.devstats.cncf.io/d/12/dashboards?refresh=15m&orgId=1) organized into groups
    *   [Consolidated Github stats](https://k8s.devstats.cncf.io/d/49/github-stats-by-repository?orgId=1) for Repos and Repo Groups
        *   [PRs Merged in Dashboard](https://k8s.devstats.cncf.io/d/49/github-stats-by-repository?orgId=1&from=now-2y&to=now&var-period=w&var-repos=%22kubernetes%2Fdashboard%22&var-stat=prmerged), shows dropoff in developmentanyway, 
        *   [Issues Opened in Storage](https://k8s.devstats.cncf.io/d/48/github-stats-by-repository-group?orgId=1&from=now-1y&to=now&var-period=w&var-repogroups=%22CSI%22&var-repogroups=%22Storage%22&var-stat=iopened), shows how increasingly issues are CSI-related
*   [ 0:00 ] **SIG Updates** (estimate ~ 7 minutes)
    *   SIG VMware [ Steve Wong confirmed, Loc Nguyen will co-present  ] 
        *   Deck [link](https://docs.google.com/presentation/d/13Zn2nvd2nyQ2WG3xETxw4V0wHsvPyIPfC8WSaEIBFQU/edit#slide=id.p)
*   [ 0:00 ] **Announcements**
    *   Bug Bounty Program [Jess]
        *   [https://github.com/kubernetes/community/pull/2620](https://github.com/kubernetes/community/pull/2620) 
    *   Election Update [Paris/Jorge/Ihor]
        *   8 days until next important deadline - Sept 14th (attn nominations and eligible voting forms)
    *   [If you are not on this list](https://github.com/kubernetes/community/blob/master/events/elections/2018/voters.md), you are not eligible to vote. Eligible voters are those with 50+ contributions to the project in the last year. If you feel like you have done other ways to contribute upstream in the last year that is outside of GitHub events, [please fill out this form](https://www.surveymonkey.com/r/k8s-sc-election-2018).
    *   Resources:[1] [Steering Committee](https://github.com/kubernetes/steering) - who sits on the committee and terms, their projects and meetings info
        *   [2] [Steering Committee Charter](https://git.k8s.io/steering/charter.md) - this is a great read if you're interested in running (or assessing for the best candidates!)
        *   [3] [Election Process](https://git.k8s.io/steering/elections.md)
        *   [4] **[Voters Guide!](https://git.k8s.io/community/events/elections/2018)** - Updated on a rolling basis. This guide will always have the latest information throughout the election cycle. The complete schedule of events and candidate bios will be housed here. 
    *   [etcd minimum versions updated](https://discuss.kubernetes.io/t/etcd-minimum-versions-3-1-11-3-2-10-3-3-0/2637) - please read if you're running etcd in production, this one's important! 
    *   [Contributor Role Board](https://discuss.kubernetes.io/c/contributors/role-board) [Jorge] 
        *   A place for SIGs and volunteers to link up.
        *   [Example role posting](https://discuss.kubernetes.io/t/kubernetes-needs-a-stack-overflow-shepharding-team/2558)
        *   [Example volunteer posting](https://discuss.kubernetes.io/t/experienced-front-end-developer-looking-for-front-end-work/1417)
        *   SIGs, please start posting available roles, especially shadow roles, we've sketched out [some guidelines here](https://discuss.kubernetes.io/t/about-the-contributor-role-board/1336) and the board has a template, all feedback welcome.
        *   Try to think of meaty roles, easy light roles are welcome, but it would be a shame if it was only just "notetakers". 
    *   Kubernetes 1.13 Release Team is forming! Contact @AishSundar or @tpepper if you want to know more. Watch this GitHub issue for updates: [https://github.com/kubernetes/sig-release/issues/280](https://github.com/kubernetes/sig-release/issues/280)
    *   Outreachy call for mentors and projects. [Check k-dev@ for complete info](https://groups.google.com/d/msgid/kubernetes-dev/CAJRwiqR6Oz6L_jAUh-AMq7syvbb4Q3oitpAqUoq70koAZPBhNQ%40mail.gmail.com.). 
    *   Shoutouts this week (Check in #shoutouts on slack) 
        *   chrishein - Big shoutout to @liggitt @sttts @directxman12 for helping to get support for HPA and the metrics-server in EKS 
        *   zparnold - Shout out to @june.yi @claudiajkang for taking me in like family on a long layover in Seoul. Your amazing Korean translations (also shoutoutable) are only matched by your generous hospitality. I love this Kubermunity even more! (edited)
        *   jdumars - Big shout out to @paris and the Meet our Contributors AUA/Group mentoring!!  Yet another way Kubernetes is setting the standard of excellence in open source communities!  
        *   bentheelder - @paris! Meet our Contributors is awesome \
Also @jorge for all the live stream work! You two rock
        *   jimangel - Shout out to @justaugustus for helping us out with wrangling docs for 1.12 before the freeze - very awesome work!
        *   Jeefy - Shout out to @jorge for taking over hosting duties with three minutes to spare as I was thwarted by contractors :)
        *   AishSundar- Huge shoutout to @tpepper for being an incredible and patient Release lead, striving non-stop to herd multiple issues, pilot and generate new branch manager playbook and keeping all documentation updated as much as possible.
    *   Call for Demos for this call, see the top of this document!


## August 30, 2018 - ([recording](https://youtu.be/i-18dbVB-ao))



*   **Moderators**:  Jorge Castro [SIG Contributor Experience]
*   **Note Taker**: First Last [Company/SIG]
*   [ 0:00 ]**  Demo **--  [Monitoring Kubernetes with Elasticsearch autodiscover](https://github.com/DanRoscigno/katacoda-scenarios/tree/master/hints-based-discovery) - [dan.roscigno@elastic.co](mailto:dan.roscigno@elastic.co)
*   GitHub form the demo: [https://github.com/DanRoscigno/katacoda-scenarios/tree/master/hints-based-discovery](https://github.com/DanRoscigno/katacoda-scenarios/tree/master/hints-based-discovery)
*   Katacoda:[ https://katacoda.com/dan_roscigno_dev/scenarios/hints-based-discovery](https://katacoda.com/dan_roscigno_dev/scenarios/hints-based-discovery)
*   [ 0:00 ]** Release Updates**
    *   Current Release Development Cycle  [Tim Pepper ~ 1.12 Release lead]
        *   Code slush arrived this week
        *   Code freeze begins next week, Tuesday Sept. 4
        *   Extremely important over next weeks for all SIGs leadership and all code committers to monitor [sig-release-1.12-all](https://k8s-testgrid.appspot.com/sig-release-1.12-all) & [sig-release-1.12-blocking](https://k8s-testgrid.appspot.com/sig-release-1.12-blocking) & [sig-release-master-blocking](https://k8s-testgrid.appspot.com/sig-release-master-blocking) & [sig-release-master-upgrade](https://k8s-testgrid.appspot.com/sig-release-master-upgrade) for de-stablization
        *   Release target Sept. 25
    *   Patch Release Updates
        *   x.x
        *   y.x
*   [ 0:00 ] **KEP o' the Week [Juan Vallejo and Maciej Szulik]**
    *   Weekly update on data from devstats or KEP
    *    [https://github.com/kubernetes/community/blob/master/keps/sig-cli/0024-kubectl-plugins.md](https://github.com/kubernetes/community/blob/master/keps/sig-cli/0024-kubectl-plugins.md)
    *   Demo: Sample CLI Plugin [https://github.com/kubernetes/kubernetes/tree/master/staging/src/k8s.io/sample-cli-plugin](https://github.com/kubernetes/kubernetes/tree/master/staging/src/k8s.io/sample-cli-plugin)
*   [ 0:00 ] **SIG Updates**
    *   SIG Cluster Lifecycle [Tim St. Clair] (confirmed)
        *   A number of new sub-projets have been brought into the fold 
*   Kubeadm-dind-cluster
*   [Cluster-api ](https://github.com/kubernetes-sigs/cluster-api-provider-aws)
                *   The Cluster API is a programming framework to rally different providers around a core set of primitives and concepts for declaratively managing api-driven cluster deployments.
                *   Planning for alpha release ~1.12, and is decoupled.
                *   Switching from aggregation api to CRDs
*   Shared Ownership of cluster api implementations for multiple providers (and SIGs)
    *   aws, gcp, openstack and vmware
    *   Kubeadm 
        *   Bug fixes, 1.11 was a big shift 
            *   Kubelet config 
            *   Kubeadm v1alpha2 config changes 
            *   HA instructions
        *   HA configurations are being updated control-plane join 
        *   Certs cleanup & renewal
        *   Config changes for v1alpha3 goal is to hit beta in 1.13 
    *   Our charter has been reviewed and LGTM'd by a number of parties we're only waiting on final approval from bgrant 
        *   For folks wanting to know more we recommend looking at 
            *   [https://github.com/kubernetes/community/tree/master/sig-cluster-lifecycle](https://github.com/kubernetes/community/tree/master/sig-cluster-lifecycle) 
*   [ 0:00 ] **Announcements**
    *   Shoutouts this week (Check in #shoutouts on slack) 
        *   neolit123 - big shoutout to @bentheelder for helping newcomers to understand how the Kubernetes test infrastructure works.
        *   bentheelder - shoutout to @neolit123 for keeping after fixes for kubeadm e2e testing, Lubomir's work should help stabilize this a bit :tada:
        *   nikhita: shoutout to @cblecker for EVERYTHING that he does for the community and especially for keeping the project sane with :bash_fire:
        *   paris and cblecker: shouts to @dims, @neolit123 and @nikhita for [adding labels to OWNERS files](https://github.com/kubernetes/kubernetes/pull/67672)
        *   Jorge: Thanks to everyone involved in our infrastructure teams transitioning CI/CD over to the CNCF. 
    *   Election Update [Paris/Jorge/Ihor]
        *   15 days until next important deadline - Sept 14th (attn nominations and eligible voting forms)
        *   6 people are officially nominated, have accepted, and need to have bios checked in to the voters guide by the above deadline. 
        *   [If you are not on this list](https://github.com/kubernetes/community/blob/master/events/elections/2018/voters.md), you are not eligible to vote. Eligible voters are those with 50+ contributions to the project in the last year. If you feel like you have done other ways to contribute upstream in the last year that is outside of GitHub events, [please fill out this form](https://www.surveymonkey.com/r/k8s-sc-election-2018).
        *   Resources:
            *   [1] [Steering Committee](https://github.com/kubernetes/steering) - who sits on the committee and terms, their projects and meetings info
            *   [2] [Steering Committee Charter](https://git.k8s.io/steering/charter.md) - this is a great read if you're interested in running (or assessing for the best candidates!)
            *   [3] [Election Process](https://git.k8s.io/steering/elections.md)
            *   [4] **[Voters Guide!](https://git.k8s.io/community/events/elections/2018)** - Updated on a rolling basis. This guide will always have the latest information throughout the election cycle. The complete schedule of events and candidate bios will be housed here. 
    *   [Contributor Role Board](https://discuss.kubernetes.io/c/contributors/role-board) [Jorge] 
        *   A place for SIGs and volunteers to link up.
        *   [Example role posting](https://discuss.kubernetes.io/t/kubernetes-needs-a-stack-overflow-shepharding-team/2558)
        *   [Example volunteer posting](https://discuss.kubernetes.io/t/experienced-front-end-developer-looking-for-front-end-work/1417)
        *   SIGs, please start posting available roles, especially shadow roles, we've sketched out [some guidelines here](https://discuss.kubernetes.io/t/about-the-contributor-role-board/1336) and the board has a template, all feedback welcome.
        *   Try to think of meaty roles, easy light roles are welcome, but it would be a shame if it was only just "notetakers". 
    *   [Stackoverflow Top Users](https://stackoverflow.com/tags/kubernetes/topusers) for August
        *   [Matthew L Daniel, Ryan Dawson, Marcin Romaszewicz, Radek 'Goblin' Pieczonka, Kun Li, Nicola Ben, jaxxstorm, VAS, Harshal Shah](https://stackoverflow.com/users/225016/matthew-l-daniel), and [mikejoh](https://stackoverflow.com/users/1547081/mikejoh)
    *   Kubernetes Slide Deck Template [Ihor]
        *   The Kubernetes-branded Slide Deck Template revealed and available to everyone at http://bit.ly/k8s-slide-template.
        *   Details: [https://groups.google.com/forum/#!topic/kubernetes-dev/CzpfB9QK-  7g](https://groups.google.com/forum/#!topic/kubernetes-dev/CzpfB9QK-7g) 


## Aug 23, 2018 - (recording)



*   **Moderators**: Paris Pittman (SIG Contributor Experience)
*   **Note Taker**: Josh Berkus and Danny Rosen
*   [ 0:00 ]**  Demo **-- [KeyCloak](https://www.keycloak.org/) - [bdawidow@redhat.com](mailto:bdawidow@redhat.com), [stian@redhat.com](mailto:stian@redhat.com) (confirmed)
    *   Keycloak is an open source IAM (Identity Access Management) solution
    *   Demo involving Ingress
        *   Set up "realm" for credentials
        *   Then set up security for Ingress endpoints
        *   Supports bearer tokens
        *   Only keycloak sees the credentials, applications only know what's authenticated by access token
        *   Handles managing multiple roles per user, with different levels of permissions by role
        *   Support for multiple identity providers (Github example)
        *   Libraries for auth for javascript, Java.  Supports general SAML libraries for other languages, also working on a goal-based proxy provider.
        *   Support for external user stores (LDAP, Kerberos, Custom)
        *   Multiple identity providers per Realm, can also have database-backed identity database locally.
        *   Keycloak can be used for authentication for Kubernetes itself
        *   Used at U Michigan
        *   Similar to OpenAM but has more features
*   [ 0:11 ]** Release Updates - **
    *   Current Release Development Cycle  [Tim Pepper ~ 1.12 Release lead]
        *   Proposal [in flight](https://groups.google.com/d/topic/kubernetes-dev/MjyJzhBEgkM/discussion) to drop "status/approved-for-milestone" from list of merge required labels during code freeze, with lazy consensus target Aug 27
        *   Code Slush:		Aug. 28
        *   Code Freeze:		Sept. 4
        *   Release Target:		Sept.25
        *   ...one month to go.  Your feature work should be wrapping up ahead of code freeze.  Docs PR's are due.  Test cases should be in place.
        *   Continuous Integration:
            *   active on 1.12 branch: [sig-release-1.12-all](https://k8s-testgrid.appspot.com/sig-release-1.12-all) & [sig-release-1.12-blocking](https://k8s-testgrid.appspot.com/sig-release-1.12-blocking)
            *   Also master branch: [sig-release-master-blocking](https://k8s-testgrid.appspot.com/sig-release-master-blocking) & [sig-release-master-upgrade](https://k8s-testgrid.appspot.com/sig-release-master-upgrade)
            *   Signal: overall health is worrying need attention from SIG-GCP (GKE?), SIG-Auth, SIG-Auth, SIG-Apps, SIG-Storage, SIG-Scalability,  SIG-Cluster-Lifecycle
    *   Patch Release Updates
        *   1.9.10 (20 days ago) - Mehdy Bohlool (@mbohlool)
        *   1.10.7 (3 days ago) - Maciek Pytel (@MaciekPytel)
        *   1.11.2 (15 days ago) - Anirudh Ramanathan (@foxish)
*   [ 0:15 ] **Graph o' the Week **[spiffxp]
    *   Let's talk about our automation's GitHub API Token usage
    *   We get: 5,000 requests per hour
    *   We used to work around this in mungegithub by:
        *   keeping an in-memory cache
        *   tuning munger polling frequency
        *   separating into SQ/misc-mungers instances
    *   Switching to prow to do things on demand vs. a polling loop helped, for a bit
    *   Now, we're using [ghproxy](https://github.com/kubernetes/test-infra/tree/master/ghproxy) (thanks @cjwagner!)
        *   Implemented by our own Cole Wagner
    *   Hero charts: last 6 months of [cache](http://velodrome.k8s.io/dashboard/db/github-cache?refresh=1m&orgId=1&from=now-6M&to=now) and [github token usage](http://velodrome.k8s.io/dashboard/db/monitoring?orgId=1&from=now-6M&to=now)
        *   See population of the cache, how many api tokens we didn't have to use over time
        *   Turned in on mid-May
        *   Prior to turning the cache on, we often hit max tokens, esp. At the end of code freeze
        *   Now usage is much more stable/lower, can go through the backlog faster
        *   We're moving away from mungegithub so you won't see this much more, [moving to Tide for merging](https://github.com/kubernetes/test-infra/issues/3866).
*   [ 0:22 ] **KEP o' the Week powered by SIG PM<sup>?</sup>**
    *   tallclair@ - [KEP](https://github.com/kubernetes/community/blob/master/keps/sig-node/0014-runtime-class.md) 0014-runtime-class
    *   RuntimeClass - Define a generic way for a runtime to be defined, where in the past it was opaque to the control plane beyond kubelet 
    *   Motivation is to support new runtimes, like katacontainers, GVisor and maybe future stuff like serverless runtimes or GPUs
    *   There's a podspec for the RunTimeClass, to decouple the configuration and node-level implementation from the name users need to use
        *   We could end up with more than one class spec for the same runtime
    *   See list of Non-Goals, we're trying to keep the mechanism simple. They do have a list of future extenions, though, such as:
        *   PodOverhead, so that you can account for resources outside those used for the container, like for Kata.
        *   Policies for abstract runtimeclasses in podspec, such as a requirement for a "sandbox" runtime or "unix" (pod doesn't care which specifically they get)
    *   Want to make it consistent to express supported/unsupported features (including mutually exclusive ones on a node like SELInux vs. Apparmor).
    *   Leave Comments:
        *   [https://github.com/kubernetes/community/pull/2489](https://github.com/kubernetes/community/pull/2489)
        *   [SIG-Node](https://github.com/kubernetes/community/tree/master/sig-node)
*   [ 0:00 ] **SIG Updates**
    *   **OpenStack** (Chris Hoge, confirmed)
        *   [https://docs.google.com/presentation/d/1fdq0X-UPN-8xc_3bpvvrwIic_UGTTDyKRt-Cjtgp9io/edit?usp=sharing](https://docs.google.com/presentation/d/1fdq0X-UPN-8xc_3bpvvrwIic_UGTTDyKRt-Cjtgp9io/edit?usp=sharing)
        *   Completed in the last cycle:
            *   CloudProvider Openstack, added conformance testing, lots of bug fixes, sync'd with in-tree provider
            *   Planned to remove the in-tree provider in 1.12, but has been delayed to 1.13 to give users time to move to external provider.
            *   Added Manilla Storage Provisioner for shared storage (NFS) 
            *   Added keystone authenticator for mapping multiple projects to accounts
            *   Added extensive documentation, including general docs for Cloud Providers
            *   Began work for transitioning to WG Openstack of SIG Cloud Provider
        *   Upcoming Work
            *   [Magnum](https://wiki.openstack.org/wiki/Magnum) (OpenStack's service for container orchestrators) conformance & cert testing toward getting it certified as a k8s installer
            *   Driver work: autoscaling drivers, [barbican](https://wiki.openstack.org/wiki/Barbican) driver for key management
    *   **Storage **(Saad Ali, confirmed)
        *   [https://docs.google.com/presentation/d/1TFX6BDCod6E0PJRusQ1zntOX36kDyuO5iycpSfH8pL4/edit?usp=sharing](https://docs.google.com/presentation/d/1TFX6BDCod6E0PJRusQ1zntOX36kDyuO5iycpSfH8pL4/edit?usp=sharing) 
        *   For 1.12:
            *   Topology-aware volume scheduling, since not all volumes work on all nodes, old version was based only on cloud providers.  Moved it to a generic interface both in Kubernetes and in CSI.
            *   This quarter moving in-tree storage to topology, and for all CSI plugins.
            *   We can have volumes provisioned in a smarter way.
            *   First Kubernetes storage features that could not be part of core.
            *   Snapshots / restore functionality (CSI, Kubernetes internal & external)
            *   Drive CSI to GA/Stable
        *   Preparing for CSI (Out of tree volume extension mechanism) for GA / Stable Q4
        *   This Quarter: Support of ephemeral volumes (eg: secret volume, configmap volume). 
        *   Moving Kubelet Device Registration to beta
        *   Adding conformance testing for storage to kubernetes storage suite
        *   Block volume support moving to Beta 
    *   **Apps (**Matt Farina, confirmed**)**
        *   [https://docs.google.com/presentation/d/1jbEDX4GDeCssT4D42Q1iajDSLU3sz_RQgPwDCkR2J1c/edit?usp=sharing](https://docs.google.com/presentation/d/1jbEDX4GDeCssT4D42Q1iajDSLU3sz_RQgPwDCkR2J1c/edit?usp=sharing)** **
        *   Active projects: 
            *   Application CRD & Controller
            *   Workload API
            *   Kompose
            *   Examples
        *   SIG Apps Charter: WIP, should be ready for review soon
        *   Recently merged: [Recommended labels](https://kubernetes.io/docs/concepts/overview/working-with-objects/common-labels/) merged into Helm documentation as well.
        *   [Application CRD & Controller](https://github.com/kubernetes-sigs/application): Cross tool way to describe an application.
        *   Workloads API: Looking at Lifecycle Hooks, Pod disruption budget & Deployments, Jobs with deterministic pod names. 
        *   Time split between Workloads API & Developer tooling week by week.
        *   Kompose:  Converts Docker Compose to Kubernetes objects, actively being worked on
        *   Helm moved to CNCF - Everything from kubernetes-helm has moved to the Helm org. Charts is still using prow/tide automation
        *   
*   [ 0:00 ] **Announcements**
    *   Shoutouts this week (pulled from #shoutouts in slack weekly)
        *   Nikhita: shoutout to @jorge for his video on how to use discuss.kubernetes.io: https://www.youtube.com/watch?v=7wTLgeM25Pk. It's incredibly detailed and helpful in gaining a quick understanding on how to use the website! :raised_hands:
        *   Paris: shouts to @dims for improving the contributor experience. Script dims created will reduce some of the manual assignment of sigs to pr(s) - [https://github.com/kubernetes/kubernetes/pull/67672](https://github.com/kubernetes/kubernetes/pull/67672) 
        *   Cblecker: also to @neolit123 and @nikhita who have been doing the same ^^
        *   Nikhita: shoutout to @cblecker for EVERYTHING that he does for the community and especially for keeping the project sane with fires
    *   [kubernetes-client/typescript has been moved to kubernetes-retired](https://groups.google.com/forum/#!topic/kubernetes-dev/nlc7SKJc9Bg) [spiffxp]
    *   Automating all the things update [spiffxp]
        *   ~~[All repos in orgs we manage should be in sigs.yaml](https://github.com/kubernetes/community/issues/2464)~~ DONE
        *   [All repos in orgs we manage should have OWNERS](https://github.com/kubernetes/community/issues/1721) one last repo
        *   ~~[All repos in orgs we manage should have at least these labels](https://github.com/kubernetes/test-infra/pull/9054)~~ DONE
        *   [All repos in orgs we manage should use the same merge automation](https://github.com/kubernetes/test-infra/issues/6227) kubernetes/* and kubernetes-incubator/* repos remain
        *   [Moving kubernetes/kubernetes to tide](https://github.com/kubernetes/test-infra/issues/3866) (we are close but changing so close to code slush / code freeze is probably unwise)
        *   Should all repos have the /retest commenter enabled?
        *   Should all repos have the /lifecycle stale|rotten|/close commenter enabled?
    *   Seattle Contributor Summit is now a part of the KubeCon registration process. Add as a co-located event. Dec 9th and 10th.
    *   Steering Committee Election Announcement went out to k-dev on Aug 21 (or 22nd depending on where you are in the world!)
        *   Next deadline: Nominations and exception eligible voter forms due on Sept 14th
    *   [Contributor Role Board](https://discuss.kubernetes.io/c/contributors/role-board) [castrojo] (will show you next time due to time constraints, in the mean time check it out!)
        *   A place for volunteers to declare interest
        *   A place for SIGs/WGs/others to post roles for volunteers.
        *   Pairs volunteers with mentors.
        *   SIGs, we'd love to get some postings from you! 
    *   We will have a Contributor Discussion Social at Kubecon Shanghai, on the evening of November 13th.  This will include drinks, snacks, and a panel Q&A on contributing to Kubernetes from China /Asia.  Anyone who contributes to Kubernetes and is at Kubecon Shanghai is invited.  Venue/schedule details TBA.
        *   If you are a Chinese contributor to Kubernetes, we are still looking for panelists.
        *   This is in addition to the New Contributor Workshop and the Doc Sprints during the day, which you can register for with your Shanghai registration.


## Aug 16, 2018 - ([recording](https://youtu.be/PpeGxZRDbII))



*   **Moderators**:  Aaron Crickenberger (@spiffxp, Google, SIG Beard)
*   **Note Taker**: Solly Ross (@directxman12, Red Hat, SIG Autoscaling)
*   [ 0:00 ]**  Demo **-- Kubernetes Ingress Controller for Kong [Harry Bagdi, harry@konghq.com] (confirmed)
    *   Links/contact
        *   [Blog](https://konghq.com/blog/kubernetes-ingress-controller-for-kong/)
        *   [https://docs.google.com/presentation/d/1U4DpjQON33UukAx7Ws5__nOhEO8TWK0EKEse2Fn7CLA/edit?usp=sharing](https://docs.google.com/presentation/d/1U4DpjQON33UukAx7Ws5__nOhEO8TWK0EKEse2Fn7CLA/edit?usp=sharing)
        *   [http://github.com/kong/kubernetes-ingress-controller](http://github.com/kong/kubernetes-ingress-controller)
        *   Email (harry@konghq.com) or use Kubernetes slack(hbagdi) to contact
    *   Kong is an open source API gateway built on nginx
        *   Performance and features from nginx
        *   flexible routing
            *   Hash-based
            *   Cookie-based
            *   client-based
        *   dynamic configuration
        *   plugins for custom logic common to your microservices
    *   Ingress Deployment
        *   Dataplane mode does the proxying, pulling config from the database
        *   Controlplane mode configures things, writing them to a database
        *   Runs in a single namespace, but serves ingresses for all namespaces
        *   Data is proxied directly to pods, skipping kube-proxy
            *   Enables things like sticky sessions in Kong
        *   Custom resource for extending normal Ingress with additional Kong functionality (KongIngress)
            *   Proxy configuration
            *   Routing methods, regex priority, etc
            *   Active and passive health checks
        *   Plugins for custom logic
            *   Use CRDs set up different plugin configurations
            *   For example, rate-limitting
            *   Apply configured plugins to ingresses with annotations specifying the name of an instance of the custom resource
            *   Have many plugins, all opensource
        *   Supports multiple services
        *   Supports TLS upstream and termination
    *   Inspection
        *   Can inject headers for info
            *   Via
            *   Latency
            *   Rate-limitting information
        *   can also be inspected using an HTTP API to check underlying Kong configuration
    *   Questions
        *   Q: How are websockets handled?
            *   Kong can forward websocket traffic directly (you can upgrade connections to websockets as normal)
            *   Can't actively manipulate traffic on websockets
*   [ 0:00 ]** Release Updates**
    *   Current Release Development Cycle  [Tim Pepper ~ 1.12 Release lead]
        *   ~2.5 weeks to code freeze!!!  Yes already!!
        *   40 days to release
        *   release-1.12 branch created Tuesday; fast forwarding daily to master
            *   Fast-forward for next couple of weeks
        *   Branch CI on track to arrive this week
        *   CI signal mostly OK for [release master blocking](https://k8s-testgrid.appspot.com/sig-release-master-blocking) and [release master upgrade](https://k8s-testgrid.appspot.com/sig-release-master-upgrade), but a number of issues being worked
    *   Patch Release Updates
        *   1.9.10 (14 days ago) - Mehdy Bohlool (@mbohlool)
        *   1.10.6 (21 days ago) - Maciek Pytel (@MaciekPytel)
        *   1.11.2 (9 days ago) - Anirudh Ramanathan (@foxish)
*   [ 0:00 ] **Graph o' the Week **[spiffxp]
    *   Let's talk about flaky and failing tests
    *   Testgrid - [presubmits-kubernetes-blocking#Summary](https://k8s-testgrid.appspot.com/presubmits-kubernetes-blocking#Summary)
        *   Show's blocking tests
        *   Also a dashboard for non-blocking tests
        *   Can click to see history of job runs in a grid, where they succeeded and failed
        *   Tests are considered failing until it sees a pass in some particular window
    *   Velodrome - [BigQuery Metrics](http://velodrome.k8s.io/dashboard/db/bigquery-metrics?orgId=1) - [Presubmit Failure Rate](http://velodrome.k8s.io/dashboard/db/bigquery-metrics?panelId=11&fullscreen&orgId=1)
        *   Grafana instance looking at test failures
        *   Can see which suites are failing over time
            *   E.g. kops spiked, integration built over time, but has been fixed (thanks @janetkuo!)
    *   GitHub Query -[ is:open label:kind/flake org:kubernetes](https://github.com/issues?q=is%3Aopen+label%3Akind%2Fflake+org%3Akubernetes+archived%3Afalse+sort%3Aupdated-desc)
        *   Can use this query to find flaky tests (intermittently failing and succeeded)
    *   GitHub Query - [is:open label:kind/failing-test org:kubernetes](https://github.com/issues?utf8=%E2%9C%93&q=is%3Aopen+label%3Akind%2Ffailing-test+org%3Akubernetes+archived%3Afalse+sort%3Aupdated-desc+)
        *   Can use this query to find tests that are failing all the time (as opposed to "just" being flaky)
    *   Who should be helping fix these?
        *   1. Who owns the test?
            *   <code>[<strong>sig-foo</strong>] thing should not explode</code>
        *   2. Who owns the job?
            *   <code>test-infra/config/jobs/kubernetes/<strong>sig-foo</strong>/OWNERS</code>
        *   3. Who owns the infra?
            *   <code>#test-infra</code>
            *   <strong>If you skip steps 1 & 2 and go directly to 3, you will be sent to the back of the line</strong>
*   [ 0:00 ] <strong>KEP o' the Week </strong>[Chris Hoge, @hogepodge, on behalf of Nishi Davidson, @d-nishi]
    *   Part of SIG Cloud Provider
        *   Coordinates stuff among all cloud providers
    *   [https://github.com/kubernetes/community/blob/master/keps/sig-cloud-provider/0019-cloud-provider-documentation.md](https://github.com/kubernetes/community/blob/master/keps/sig-cloud-provider/0019-cloud-provider-documentation.md)  - Accepted 
        *   Transfer responsibility of maintaining docs to cloud providers
        *   Provide documentation on how to activate any out-of-tree cloud provider
        *   Set minimum standards for cloud provider documentation
        *   Maintain docs for how to write a new out-of-tree cloud provider
    *   Follow up discussion in SIG-Cloud-Provider and SIG-AWS
    *   QuestionsC
        *   Q: Working with Cluster Lifecycle to improve workflow in kubeadm?
            *   Yes, working on docs to start out with
*   [ 0:00 ] <strong>SIG Updates</strong>
    *   SIG Docs [Andrew Chen]
        *   [[slide link](https://docs.google.com/presentation/d/1d2DPGphVERniJBT011ZLDwc7nttUv7sqVukw362ozY4/edit?usp=sharing)]
        *   Ongoing/upcoming work
            *   1.12 is under (@zparnold is docs lead)
            *   [Docs contributor guide](https://kubernetes.io/docs/contribute/) has been refactored (@mistyhacks)
            *   Considering alternative search engines for China [PR#9845](https://github.com/kubernetes/website/pull/9845)
            *   Figuring out generated docs (working group) -- e.g. for kubelet [PR#66034](https://github.com/kubernetes/kubernetes/pull/66034)
            *   Proposal for fundamental concepts of Kubernetes (modeling, architecture) [[slides](https://docs.google.com/presentation/d/1vUAkRP-MjNqusqDHBptycdSbC_HSTBKyEHAFZ5OdbQA/edit?usp=sharing)]
                *   Need more/helpful diagrams
        *   PR bash and docs sprint at Write the Docs in Cincinnati
        *   Search outage postmortem [[doc](https://docs.google.com/document/d/1WxrincD0K_IW6VazR4YhMFAC4OKAEdsIJ5L4QXHsJdQ/edit?usp=sharing)]
            *   Kubernetes.io dropping off of search results
            *   Version docs aren't indexed (via X-Robots-Tag: noindex)
            *   Noindex header got added to main site as well by accident, causing no search engine results
            *   What to do going forward
                *   Hand off infra to CNCF, document mechanisms and processes
                *   Adding testing and monitoring, notify on abnormalities
                *   Have better failsafe default state
                    *   master was the exception before, default state was "nothing gets indexed"
                    *   default state should have been "everything got indexed"
    *   SIG IBMCloud [Sahdev Zala]
        *   [Slide deck](https://docs.google.com/presentation/d/1B1UDsHKnFDa3WvOdEOQADrhs18yOxkQkVbGuT4H8NeQ/edit?usp=sharing)
        *   Relatively new SIG for building/maintaining/using Kubernetes with IBM public and private clouds
        *   Meets every other week (Wednesdays at 14:00 EST)
            *   Start with presentations about IBM Cloud Kubernetes Service, IBM Cloud Private (recorded)
                *   IKS supports 3 concurrent releases, multi-az clusters
                *   IBM Cloud Private 2.1.0.3 releaed in May, certified for up to 1000 nodes, scalability work ongoing
        *   Ongoing discussions/work
            *   SIG cloud provider integration
            *   Public repo for IBM cloud provider code
            *   SIG Charter
        *   Future discussions (see SIG agenda)
            *   Hybrid clouds (IKS <-> ICP)
            *   Performance
        *   Community Collaboration
            *   Networking
                *   Working with Red Hat & Tigera
                    *   Move Egres/IPBlock network policy to GA in 1.12
            *   Scalability
                *   Etcd changes to improve cluster creation, improve monitoring overhead
            *   Storage
                *   Flex volume resize and metrics
                *   IBM Cloud object store plugins
    *   SIG Autoscaling [Solly Ross]
        *   SIG is in charge of anything related to automatic scaling both of pods, cluster components themselves, and the cluster (VMs) itself
        *   Horizontal Pod Autoscaler
            *   Removing scale limits in favor of more sophisticated behavior (looking at metric data point timestamps and pod launch timestamps)
            *   Brainstorming further algorithmic improvements (looking at more than one data point, etc) for flexibility around additional use cases and custom metrics
            *   HPA v2beta2 landing in 1.12 release
                *   Specify labels to further scope metrics
                *   Target average values on object metrics (divide value by number of pods)
                *   API consistency improvements
        *   Cluster Autoscaler
            *   Focusing on some large known issues (scaling around GPUs, local persistent volume scaling)
            *   Investigating steps to integrate cluster autoscaler with cluster API (may require some changes to the cluster API instead of custom logic in the autoscaler)
*   [ 0:00 ] <strong>Announcements</strong>
    *   Shoutouts this week
        *   shoutout to Di Xu (@dixudx) for being such an active reviewer and reviewing LOTS of incoming PRs so quickly!!!
        *   shoutout to Arnaud Meukam (@ameukam) and Jeremy Rickard (@jerickar) for being awesome bug triage shadows and handling the job wonderfully while I was out last week!
        *   Mistyhacks: Shoutout to @ianychoi, who has just become a k8s org member in order to work on Korean localization, and is already providing great feedback, as evidenced in this PR: [https://github.com/kubernetes/website/pull/9643/comment#issuecomment-411886340](https://github.com/kubernetes/website/pull/9643/comment#issuecomment-411886340)
        *   @jdumars for creating :testgrid: (Slack emoji)
    *   Steering Committee Elections are coming! Announcements will go out next week on multiple platforms but k-dev@ will be the main communication channel.
        *   [Elections](https://github.com/kubernetes/community/tree/master/events/elections) are coming!
        *   Next week, email will go out with eligibility, etc information on kubernetes-dev ML
        *   There will be a voters guide checked into GitHub as a single source of truth
    *   Changing how we do GitHub membership - file an issue instead of send an e-mail [spiffxp]
        *   [https://github.com/kubernetes/community/pull/2521](https://github.com/kubernetes/community/pull/2521) 
        *   Path to Kubernetes membership used to involve sending an email
        *   Now moving to issue-based system (file issue from template to [kubernetes/org](https://github.com/kubernetes/org) repo) : [Membership issue template](https://github.com/kubernetes/org/blob/master/.github/ISSUE_TEMPLATE/membership.md)
        *   Currently manual, may be automated in the future with filing a PR
    *   Brace yourselves, automation is coming [spiffxp]
        *   I am pushing an agenda as follows, and want help figuring what is common and appreciated, vs. what should be configurable and opt-in
        *   Want to have some common processes across org
            *   Everything uses tide except for k/k, looking to move k/k to tide as well
            *   Common label colors to easily recognize things
        *   [All repos in orgs we manage should be in sigs.yaml](https://github.com/kubernetes/community/issues/2464)
        *   [All repos in orgs we manage should have OWNERS](https://github.com/kubernetes/community/issues/1721) (what are OWNERS and how do we use them? [https://go.k8s.io/owners](https://go.k8s.io/owners) )
        *   [All repos in orgs we manage should have at least these labels](https://github.com/kubernetes/test-infra/pull/9054) (what labels? [https://go.k8s.io/github-labels](https://go.k8s.io/github-labels) 
        *   [All repos in orgs we manage should use the same merge automation](https://github.com/kubernetes/test-infra/issues/6227) (with some configurable settings)
        *   Should all repos have the /retest commenter enabled?
        *   Should all repos have the /lifecycle stale|rotten|/close commenter enabled?
    *   Heapster deprecation reminder [directxman12]
        *   Bug-fix only mode on 1.12, completely deprecated & retired in 1.13
        *   Please start the process of migrating away from Heapster if you haven't already (look at [metrics-server](https://github.com/kubernetes-incubator/metrics-server) and/or third-party monitoring solutions, such as Prometheus)
    *   


## Aug 9, 2018 - recording



*   **Moderators**:  Arun Gupta [Amazon]
*   **Note Taker**:  Tim Pepper [VMWare/SIG Release and Jorge Castro [Heptio/SIG Contribex] and Josh Berkus [Red Hat/SIG Release etc.]
*   [ 0:00 ]**  Demo **-- No demo this week
*   [ 0:00 ]** Release Updates**
    *   Current Release Development Cycle  [Tim Pepper ~ 1.12 Release lead]
        *   We are roughly halfway through the ~12-13 week release cycle for 1.12, but almost ? of the way through our open development phase:
            *   It's been ~50 days since master branch reopened from 1.11's freeze
            *   It is only 26 days to 1.12's code freeze!
        *   1.12.0-beta0 is Aug. 14:  We are validating a new build/publish mechanism and its documentation.  Beta should be cut from a newly created 1.12 release branch next week, CI will be enabled on the branch, and the branch will fast-forward regularly pulling master branch's content for the next weeks.
        *   Looking for high SIG attention toward keeping CI signal green for [release master blocking](https://k8s-testgrid.appspot.com/sig-release-master-blocking) and [release master upgrade](https://k8s-testgrid.appspot.com/sig-release-master-upgrade)
        *   **Code Freeze:		September 4**		(26 days from today)
        *   **Release Target:	September 25**		(47 days from today)
    *   Patch Release Updates
        *   1.9.10 (5 days ago)
        *   1.10.6 (12 days ago)
        *   1.11.2 (1 day ago)
*   [ 0:00 ] **SIG Updates**
    *   SIG Scalability [Shyam Jeedigunta] (confirmed)
        *   Recent work toward improving tools for scale testing:
            *   "Clusterloader" rewrite (in [perf-tests](https://github.com/kubernetes/perf-tests) repo) for declarative cluster setup and measurements. [[Discussion](https://groups.google.com/forum/#!topic/kubernetes-sig-scale/2G6lNidNuaU)] Thanks @sejug for the awesome initial version of the tool.
            *   "perfdash" (code also in [perf-tests](https://github.com/kubernetes/perf-tests) repo) dashboard output to [http://perf-dash.k8s.io/](http://perf-dash.k8s.io/)
        *   For 1.12 kubelet watches for secrets instead of polling, making a big perf win, can scale to a 100k namespaces currently. 
        *   Kubelet heartbeat changes to reduce etcd interactions (see [KEP 0009 node heartbeat](https://github.com/kubernetes/community/blob/master/keps/sig-node/0009-node-heartbeat.md))
            *   Moving node heartbeat to another API
            *   Current node heartbeat produces a LOT of etcd version history, bloating the etcd database
        *   CI Testing
            *   Deflaking our jobs
            *   Solving 1.12 regression
    *   SIG Architecture [Brian Grant] (confirmed)
        *   Meeting time has move to 11am Pacific, immediately following this community meeting
        *   Work tracking methods are changing: 
            *   [API Reviews](https://github.com/kubernetes-sigs/architecture-tracking/projects/3)
            *   [KEP Tracking](https://github.com/kubernetes-sigs/architecture-tracking/projects/2)
            *   [ tracking boards](https://github.com/kubernetes-sigs/architecture-tracking/projects) 

        GitHub notifications don't work for most and slack also is lossy.  Use mailing list.

        *   [Tracking boards ](https://github.com/kubernetes-sigs/architecture-tracking/projects)- if you want to get on the SIG Arch radar, please get onto the project board so you can get on the agenda. Feel free to use the sig-architecture mailing list to reach out to us. (Slack is too ephemeral, please use the list as the primary point of contact.) 
        *   Pushing back on newly compiled-in APIs, reviewing those more closely.
        *   Will post to k-dev on the engagement model for interacting with API changes ? **important**
    *   SIG CLI [Sean Sullivan] (confirmed)
        *   [Slide Deck](https://docs.google.com/presentation/d/1k-dqABfPPVUtppGOIMi0PWSWPwUDbri0pWVu7D5MoFA/edit?usp=sharing)
        *   kubectl update: ongoing work to simplify client and move more into the api server,, improving extensibility with plugins ([KEP PR link](https://github.com/kubernetes/community/pull/2437)), "KREW" plugin management (link?)
        *   [kustomize](https://github.com/kubernetes-sigs/kustomize) update: patch based customization (instead of Helm style templates/yaml)
        *   Charter draft [out for feedback](https://github.com/kubernetes/community/pull/2453)
        *   [SIG-CLI](https://github.com/kubernetes/community/pull/2453) Links:
            *   [Agenda Notes](https://docs.google.com/document/d/1r0YElcXt6G5mOWxwZiXgGu_X6he3F--wKwg-9UBc29I/edit?usp=sharing)
            *   Slack Channel [#sig-cli](https://kubernetes.slack.com/messages/sig-cli)
            *   [Google Group](https://groups.google.com/forum/#!forum/kubernetes-sig-cli)
            *   [Testing Playbook](https://docs.google.com/document/d/1Z3teqtOLvjAtE-eo0G9tjyZbgNc6bMhYGZmOx76v6oM/edit#) 
            *   [Test Grid](https://k8s-testgrid.appspot.com/sig-cli-master)
        *   kustomize repo is under [kubernetes-sigs](https://github.com/kubernetes/community/blob/master/github-management/kubernetes-repositories.md)
            *   Project-wide tip, kubernetes-sigs is the repo that holds things for SIGs that shouldn't be in core but need to be managed by SIGs. It's the evolution of the old kubernetes incubator module. 
    *   SIG AWS [Nishi Davidson] (confirmed)
        *   Slides link
        *   Looking to upstream more, especially documentation and testing
        *   Repos now in kubernetes-sigs namespace
        *   Giving an overview of subprojects:
            *   [Aws-iam-authenticator](https://github.com/kubernetes-sigs/aws-iam-authenticator), allows authentication against IAM credentials for kubernetes running on AWS. Renamed from heptio-authenticator.
            *   [Aws-alb-ingress-controller](https://github.com/kubernetes-sigs/aws-alb-ingress-controller), created by CoreOS and Ticketmaster & donated, watches for ingress events on kubernetes and creates AWS ALBs.  It's in production at Ticketmaster (also used by Bluejeans & Freshworks).  At some point will be added to Amazon EKS.
            *   [Aws-encryption-provider](https://github.com/kubernetes-sigs/aws-encryption-provider) provides envelope encryption for Etcd, still an alpha project where they are debating design elements.
            *   [Aws-csi-driver-ebs](https://github.com/bertinatto/ebs-csi-driver/) allows the CSI driver to work with EBS for PVs.  Collab with Red Hat.  Hope to make stable in 1.13/1.14 and replace the current EBS driver.
            *   Pod-identity-access: just a proposal right now.  Would like to have identity injection inside the pod for IAM credentials.  Target for 1.13/1.14 work.
            *   Cloud-provider-aws:  project to move AWS cloud provider to the cloud provider API (as per [KEP 0019](https://github.com/kubernetes/community/blob/master/keps/sig-cloud-provider/0019-cloud-provider-documentation.md)).  Added a documentation KEP for it.
    *   Cluster API [Kris Nova]
        *   Repository: [https://github.com/kubernetes-sigs/cluster-api-provider-aws](https://github.com/kubernetes-sigs/cluster-api-provider-aws)
        *   Doodles and Docs email: [https://groups.google.com/forum/#!topic/kubernetes-sig-cluster-lifecycle/__XIKigkxkA](https://groups.google.com/forum/#!topic/kubernetes-sig-cluster-lifecycle/__XIKigkxkA)
*   [ 0:00 ] **Steering Committee Updates **[Aaron @spiffxp]
    *   [Steering Committee Elections 2018](https://github.com/kubernetes/steering/issues/63)
    *   Walked through how a meeting works:
        *   [kubernetes/steering project board](https://github.com/kubernetes/steering/projects/1)
        *   They start with a kanban board and look at all of the things they were supposed to have done
        *   Right now they're supposed to be having elections, but there are pending tasks that weren't done a year ago, like deciding who is a "member of standing".  
        *   Went over criteria for member of standing.  Right now they're planning to use Devstats criteria for contributions by contributor ([rolling window 1year](https://k8s.devstats.cncf.io/d/13/developer-activity-counts-by-repository-group?orgId=1&var-period_name=Last%20year&var-metric=contributions&var-repogroup_name=All)), requiring 60 contributions.
        *   Need to [codify SIG liaisons](https://github.com/kubernetes/steering/issues/64) from SC.  This is partly for the charter process.  Have at least 2 people assigned to each SIG.
    *   Code of Conduct Committee (CoCC): open candidates, closed voting -> [set of members added in community repo](https://github.com/kubernetes/community/pull/2498).  See [committee readme](https://github.com/kubernetes/community/tree/master/committee-code-of-conduct) for more info.
    *   Charters: lots of activity but also slow progress.  WIP, lots to do, tracked in [meta issue](https://github.com/kubernetes/steering/issues/31).
    *   Meet Our Contributors - Steering Committee edition
        *   Sep 5th at 1pm PST -> sarahnovotny, bgrant, quinton-hoole, spiffxp, jbeda, michelleN, derekwaynecarr
        *   Join #meet-our-contributors on kubernetes.slack.com for more info
        *   [Last edition with bdburns, pwittroc, and philips](https://youtu.be/BuJhzJriaNY)
        *   [Meet Our Contributors ](https://github.com/kubernetes/community/blob/master/mentoring/meet-our-contributors.md)info
    *   Non SC participation: Would like to allow non-SC members to join the meetings by invitation (meetings are recorded though and posted to the youtube channel for community review), such as Jaice who has been auditing the meetings and asking questions.  Another example is cblecker querying the SC about GH permissions management, and made a proposal for it.  Not suggesting making the meetings open, joining would be by invitation, usually based on a proposal to the SC.
*   [ 0:00 ] **Announcements**
    *   [Kubernetes Office Hours](https://github.com/kubernetes/community/blob/master/events/office-hours.md) is next week! [Jorge]
    *   SIG Update Schedule for this meeting [is updated](https://docs.google.com/spreadsheets/d/1adztrJ05mQ_cjatYSnvyiy85KjuI6-GuXsRsP-T2R3k/edit#gid=1543199895) through October [Jorge]
        *   It is always linked to from the top of this document
        *   SIGs, it is your responsibility to ensure that you can make this update, if not, let someone in SIG Contrib-Ex know so we can schedule you.
    *   Demo section is finally caught up! If you want to demo something during this meeting see the top of this document. [Jorge]
        *   If you've demo'ed over a year ago consider submitting again so we can check out your progress!
    *   GitHub Management subproject [Aaron @spiffxp]
        *   [The responsibilities](https://github.com/kubernetes/community/blob/master/github-management/subproject-responsibilites.md)
        *   [The team](https://github.com/kubernetes/community/tree/master/github-management#github-administration-team)
        *   [The GitHub orgs we manage](https://github.com/kubernetes/community/tree/master/github-management#actively-used-github-organizations)
        *   [We have removed GitHub Owner privileges from most people who had it](https://github.com/kubernetes/community/issues/2465)
            *   If you need things done that require Owner privileges, [please file an issue on kubernetes/org](https://github.com/kubernetes/org/issues/new/choose)
            *   If this broke your workflow, [please file an issue on kubernetes/org](https://github.com/kubernetes/org/issues/new/choose)
        *   [We setup domain verification for the kubernetes GitHub org](https://github.com/kubernetes/org/issues/4)
        *   If you need something done by GH admins, then you file an issue on the kubernetes/org GH repo.  There's an SLO for responses.
    *   Subprojects [Aaron @spiffxp]
        *   A subproject is 1 or N repos or packages 
        *   [All repos must have OWNERS](https://github.com/kubernetes/community/issues/1721)
        *   [All repos must be called out in sigs.yaml](https://github.com/kubernetes/community/issues/2464)
        *   I will eventually be asking why repos in kubernetes-incubator can't instead live in kubernetes-sigs
    *   Sunsetting Kubernetes SIG service accounts [Ihor]
        *   We are not enforcing SIG's to create&use the service accounts anymore, and sunsetting the existing procedure;
        *   If you are an owner of the credentials for `k8s-mirror-foo-*` GitHub users and/or `kubernetes-sig-foo-*` mailing lists, please, work on deleting them
        *   Details - [https://groups.google.com/d/msg/kubernetes-dev/Fg_dWVV0eIQ/bvv64v46CwAJ](https://groups.google.com/d/msg/kubernetes-dev/Fg_dWVV0eIQ/bvv64v46CwAJ) 
    *   Shoutouts this week (Check in #shoutouts on slack) 
        *   paris: thanks to @tpepper (Tim Pepper) @jeefy (Jeffrey Sica) @bentheelder (Benjamin Elder) @rdodev (Ruben Orduz) for great responses and their time on #meet-our-contributors yesterday! :tada:  solid examples of good mentors
        *   spiffxp: thanks to @mhb (Morgan Bauer) for his efforts in working with #sig-testing to get service-catalog all hooked up to prow :prow: and tide
        *   Jerickar (Jeremy Rickard): what @spiffxp said!  tide and prow are dope and we love using the now
        *   tpepper: shoutout to @jorge , @paris , zoom, and any others who've been working for _months_ to improve our meeting moderation abilities and best practices to better insure our collaborations are constructive and resilient in the face of potential abuse
        *   spiffxp: shoutout to @matthyx (Matthias Bertschy) for adding per-repo label support to our `label_sync` bot, so you can add labels to your repo by PR'ing a file instead of making the change manually with admin access
        *   jorge: shoutout to @chenopis (Andrew Chen) for sorting out netlify for the contributor site!
        *   spiffxp: shoutout to @mkumatag (Manjunath Kumatagi) and @dims (Davanum Srinivas) for their push on multi-arch e2e test images, ppc64le is now passing node conformance (https://k8s-testgrid.appspot.com/sig-node-ppc64le#conformance)
    *   


## August 2, 2018 - recording



*   **Moderators**:  Solly Ross
*   **Note Taker**: Bob Killen [Company/SIG]
*   [ 0:00 ]**  Demo **-- Kritis Overview [[aprindle@google.com](mailto:aprindle@google.com)]
    *   [Slides](https://docs.google.com/presentation/d/14CaYbhH_mhrhLg2ULm4fXMfwWOYS_A_EjZwl6olUKWc/edit?usp=sharing )
    *   Build off of grafeas
    *   Test assertions before deploying containers
    *   Can validate / do vulnerability scanning 
    *   Cron schedule that is constantly monitoring to ensure images are never fall out of sync
    *   CRD based configuration
        *   Supports whitelisting images
        *   Can define things such as maximum CVE severity
        *   Can deny images that are usings tags such as 'latest'
    *   Helm Chart available for deployment
    *   When attempting to deploy an image with a vulnerability, user will be given a denied error
    *   Blog post incoming on August 13th
    *   Initial v0.1.0 release coming soon
    *   Custom attestation policies in the future
    *   Questions:
        *   Is like [Portieris](https://github.com/IBM/portieris)? Unknown, will look / follow up
        *   Does it support [Notary](https://github.com/theupdateframework/notary)?  Auth piece has similar goal
            *   Notary support should be possible, both designed for build provenance
    *   [[slides](https://docs.google.com/presentation/d/14CaYbhH_mhrhLg2ULm4fXMfwWOYS_A_EjZwl6olUKWc/edit#slide=id.p)]
    *   [https://github.com/grafeas/kritis](https://github.com/grafeas/kritis)
*   [ 0:00 ]** Release Updates**
    *   Current Release Development Cycle  [Tim Pepper ~ 1.12 Release lead]
        *   v1.12 Feature Freeze: was Tuesday July 31
            *   This was feature definition (not implementation) deadline.
            *   ~50 features captured
            *   Implementation (code, test cases, docs drafts)deadline is "Code Freeze" on Sept. 4
        *   v1.12.0-alpha.1 release cut yesterday Aug. 1.  Is a major milestone in that along with 1.11.1 we are transitioning from Google employees running the build/release mechanism to community members.  The transition has had a few issues, but is rapidly improving.  Expecting first beta to be smooth.
        *   More details and links at [http://bit.ly/k8s112-release-info](http://bit.ly/k8s112-release-info) 
    *   Patch Release Updates
        *   1.11.2 Scheduled for the Aug 11th (cherry picks should be up by Friday, August 3rd)
*   [ 0:00 ] **Open KEPs**
    *   Dynamic Audit Configuration [https://github.com/kubernetes/community/blob/master/keps/sig-auth/0014-dynamic-audit-configuration.md](https://github.com/kubernetes/community/blob/master/keps/sig-auth/0014-dynamic-audit-configuration.md)
        *   Advanced auditing is still difficult to configure
        *   Working on making it similar to dynamic admission control
        *   Support both static runtime configuration via flag and new dynamic method
        *   Moving into alpha in 1.12, beta in 1.13
        *   Will be Feature Gated
        *   Can be used to compute API coverage on a running cluster.  Previously it was not possible to alter the audit config of a running cluster.  Dynamic audit config allows you to turn on API coverage calculator and compute the API usage for a period of time.
*   [ 0:00 ] **SIG Updates**
    *   SIG UI [Jeffrey Sica] (confirmed) \
[https://docs.google.com/presentation/d/1f6dI2mP_5SZeuJd9i3e6y6jx44i6ouFZGvFAfYT1BsA/edit?usp=sharing](https://docs.google.com/presentation/d/1f6dI2mP_5SZeuJd9i3e6y6jx44i6ouFZGvFAfYT1BsA/edit?usp=sharing) 
        *   New release coming soon (2-3 weeks)
            *   Many bug fixes
            *   Will use 1.8.10 client-go
        *   Angular Migration in progress
            *   Migrating from version 1 to version 6
            *   Requires a complete rewrite
        *   Upcoming features
            *   oauth2 integration
            *   multi-arch manifests
            *   security enhancements
                *   inform users when running as admin or with other insecure configuration
            *   Will support multiple themes
            *   Customized CSS (branding etc)
        *   Looking for more contributors
            *   angular js migration
            *   bug triage
            *   feature discovery
    *   ~~SIG AWS [Nishi Davidson] (confirmed)~~ (had to move to later week)
    *   SIG Service Catalog [Jeremy Rickard] (confirmed)
        *   SIG Charter recently approved
        *   SIG Chairs have changed recently (insert names later)
        *   Working actively on improving contributor experience 
            *   active in labeling issues
            *   improving contributor guide
        *   Moving to prow
        *   Service Catalog now supports namespace
        *   Catalog restrictions on a per namespace basis
        *   Working towards providing default types for services
*   [ 0:00 ] **Announcements**
    *   [Kubecon CFP deadline](https://events.linuxfoundation.org/events/kubecon-cloudnativecon-north-america-2018/program/call-for-proposals-cfp/)
    *   Save the Date: Kubernetes Contributor Summit, 10 December, right before Kubecon.
        *   Sunday, 9 December will likely  
    *   Shoutouts this week (Check in #shoutouts on slack) 
        *   thanks to @mhb for his efforts in working with #sig-testing to get service-catalog all hooked up to prow :prow: and tide
        *   thanks to @tpepper @jeefy @bentheelder @rdodev for great responses and their time on #meet-our-contributors yesterday! :tada:  solid examples of good mentors
        *   Shout out to @neolit123 for quick responses and status updates to failing ci tests in cluster lifecycle
        *   Many thanks to @ahmet for quick reviews of changes to kubernetes/examples repo!
    *   [Stackoverflow Top Users](https://stackoverflow.com/tags/kubernetes/topusers) (Once a month at the end of the month)
        *   [Jaxxstorm](https://stackoverflow.com/users/645002/jaxxstorm), [Matthew L Daniel, Nicola Ben, David Maze, Konstantin Vustin,](https://stackoverflow.com/users/225016/matthew-l-daniel) [VAS, VonC, Michael Hausenblas, Const, Marcin Romaszewicz](https://stackoverflow.com/users/9521610/vas)
    *   Turning off bot for 1.12 release, last artifact of munge github (missed 1st part of this)
    *   Contributor Experience is looking for new contributors
    *   SIG leads have an email regarding zoom

