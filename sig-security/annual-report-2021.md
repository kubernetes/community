# Kubernetes SIG Security 2021 Annual Report

## Operational

*   **How are you doing with operational tasks in [sig-governance.md](https://github.com/kubernetes/community/blob/master/committee-steering/governance/sig-governance.md)?**
    *   **Is your README accurate? have a CONTRIBUTING.md file?**
        *   Our README is accurate. As of this report, we don’t yet have a CONTRIBUTING.md file. Through the process of building this report, we have figured out more about ourselves and our vision as a newer SIG, and we have come to realize we need one specific to our SIG, which highlights our differences and encourages new contributors more effectively. Accordingly, we will be writing one shortly.
    *   **All subprojects correctly mapped and listed in [sigs.yaml](https://github.com/kubernetes/community/blob/master/sig-list.md)?**
        *   Yes. As a horizontal SIG, we do not naturally have a large number of OWNERS files, but we have created some to ensure our subproject owners are easy to identify by viewing our README.md.
    *   **What’s your meeting culture? Large/small, active/quiet, learnings?**
        *   Our meetings are open, collaborative, and loosely planned. We actively ask for sharing to encourage fuller participation, and give time to foster community discussion. In every meeting we hear reports from our subprojects, review selected security-related issues and PRs together, and hear and discuss thoughts and ideas community members have brought to share. Much of what happens in any given meeting depends on who shows up and what they choose to bring.
        *   **Meeting notes up to date?**
            *   Yes. In keeping with the open, collaborative culture we foster, [we take notes](https://docs.google.com/document/d/1GgmmNYN88IZ2v2NBiO3gdU8Riomm0upge_XNVxEYXp0) collectively during the meetings, led by a volunteer note taker. We ask for a new volunteer note taker at the start of each meeting, which gives people an additional way to participate.
        *   **Are you keeping recordings up to date?**
            *   Yes. SIG chairs upload the recording videos to YouTube, typically within a few days of the meeting.
        *   **Trends in community members watching recordings?**
            *   [Recordings](https://www.youtube.com/playlist?list=PL69nYSiGNLP1mXOLAc9ti0oX8s_ookQCi) mostly serve as a historical record. There have been 18 SIG-related meetings between Jan 1 and April 8 2021, including the main SIG meeting and subproject meetings, which have an average of 2.94 views on YouTube.
*   **How does the group get updates, reports, or feedback from subprojects?**
    *   Subproject owners or designees talk about it in the main SIG meeting as a standing agenda item.
    *   **Are there any springing up or being retired?**
        *   Because we are a newer SIG, there hasn’t been much opportunity for turnover. We have two established and active subprojects: security-docs and security-audit. A third and newer, security-tooling, is just starting to spring up.
    *   **Are OWNERS files up to date in these areas?**
        *   Yes. As a horizontal SIG that mostly works by interacting with and contributing to other SIGs’ code and projects, we don’t have much need for OWNERS files right now. We have directories for both of the more established subprojects to ensure they are recognized.
*   **Same question as above but for working groups.**
    *   We do not currently have any working groups.
*   **When was your last monthly community-wide update? (provide link to deck and/or recording)**
    *   Since our founding in October 2020, we haven’t given a monthly community-wide update yet. We have a maintainer track session coming up [at KubeCon EU 2021](https://kccnceu2021.sched.com/event/iE5u/).

## Membership

*   **Are all listed SIG leaders (chairs, tech leads, and subproject owners) active?**
    *   Yes
*   **How do you measure membership? By mailing list members, OWNERs, or something else?**
    *   We are interested in membership metrics for the purposes of gauging the health and activity of our community. This leads us to two types of measures: measures of reach and measures of collaboration.
    *   We measure reach using the number of members of our Slack channel and mailing list; these are potential future collaborators on SIG projects. As of April 8, 2021, there are 438 members of the main Slack channel and 131 members of the mailing list.
    *   We measure collaboration by noting who regularly contributes to our community activities. We take attendance at our meetings, which helps with this. The 7 main SIG meetings between January 14, 2021 and April 8, 2021 have had an average of 12.4 documented attendees.
    *   A definitive count of collaborating members across all media at any given moment is difficult to calculate, which we are comfortable with at this time.
*   **How does the group measure reviewer and approver bandwidth? Do you need help in any area now? What are you doing about it?**
    *   Our role in the Kubernetes project means that the majority of our reviewer bandwidth is used on triaging other SIGs’ PRs, KEPs, and Issues to bring them to our membership for comment. Note that this is a different concept of “reviewer” than is typical in a code-owning SIG.
    *   We measure our reviewer bandwidth by noting whether everything brought to SIG Security’s attention is getting addressed in a timely fashion, and the feelings of the reviewers about their workload.
    *   Much of this review is currently done by the SIG chairs, and it is starting to become a bottleneck. We are looking to improve our review bandwidth by bringing review of triaged issues to the regular SIG meetings for more open and inclusive participation. 
*   **Is there a healthy onboarding and growth path for contributors in your SIG? What are some activities that the group does to encourage this? What programs are you participating in to grow contributors throughout the contributor ladder?**
    *   We are onboarding and growing contributors continuously through our processes and culture. We actively practice behaviors of inclusion to help new and growing contributors feel more confident to participate. For example, we list appropriate pronouns in our meeting attendance list, and before changing subjects we pause to ask for additional comments.
    *   In addition to participating in SIG discussions, new contributors can also begin by assisting the subprojects with their tasks. Subproject members guide new contributors to where assistance is needed.
    *   Our open environment creates space for interested contributors to become leaders [by sharing their ideas and energy](https://twitter.com/PuDiJoglekar/status/1380205701994147845). If someone brings an idea that other members are excited about and want to help work on, [those members can organize a subproject together](https://twitter.com/coffeeartgirl/status/1331046904306622465), and we can help them grow. 
    *   The chairs help grow emerging leaders by keeping in close contact with them and providing 1:1 mentorship, [encouragement](https://twitter.com/coffeeartgirl/status/1337078959020912642), and concrete help leveling up skills such as improving communication, leading more inclusive meetings, and presenting at conferences.
*   **What programs do you participate in for new contributors?**
    *   Our broadly welcoming culture is a good opportunity for security-minded Kubernetes users who want to get involved in the upstream Kubernetes community. We plan to extend this further by participating in formal “new contributor” programs in the future.
*   **Does the group have contributors from multiple companies/affiliations?**
    *   Yes
    *   **Can end users/companies contribute in some way that they currently are not?**
        *   End users are currently contributing at all levels. Our contributors represent both end user organizations and Kubernetes vendors. Both SIG Security co-chairs work at companies that are Kubernetes end users.

## Current initiatives and project health

*   **What are initiatives that should be highlighted, lauded, shout out, that your group is proud of? Currently underway? What are some of the longer tail projects that your group is working on?**
    *   We’re proud of the work that our security-docs subproject has started to create a Kubernetes hardening guide. They have collected contributions from across the Kubernetes community into a document outline, and now they’re writing it as a team. You may learn more about this effort in the Kubernetes SIG Security Docs Subproject [meeting notes](https://docs.google.com/document/d/11LZn7qWB0OzbpF8va_YYGQE4fuRARCGY9KL87hwBLBI).
    *   Our third-party security audit subproject is moving forward through the complexity of organizing an audit of one of the world’s largest open-source projects. The key challenges are the scale of our codebase and limited auditor experience with the Go language and cloud-native technologies. They are utilizing members’ existing relationships with audit firms, as well as building new ones, to find solutions to these challenges. 
    *   PodSecurityPolicy was deprecated in the Kubernetes 1.21 release, and SIG Security worked closely with SIG Auth to figure out what comes next. We are delighted to have contributed to [KEP 2579](https://github.com/kubernetes/enhancements/issues/2579) by writing and reviewing proposals, adding our members’ voices to the design meetings, and [blogging about it for the broader Kubernetes community](https://kubernetes.io/blog/2021/04/06/podsecuritypolicy-deprecation-past-present-and-future/). We hope to make our close collaboration with SIG Auth a model for future cross-SIG efforts throughout Kubernetes. Thanks again to everyone who has been involved!
*   **Year to date KEP work review: What’s now stable? Beta? Alpha? Road to alpha?**
    *   We have accepted stewardship of KEPs [1753](https://github.com/kubernetes/enhancements/issues/1753) and [1933](https://github.com/kubernetes/enhancements/issues/1933) from SIG Instrumentation. Their owners and progress remain the same, and we look forward to helping them move toward graduation.
    *   KEP [1933](https://github.com/kubernetes/enhancements/issues/1933) - Defend against logging secrets via static analysis - [graduated](https://github.com/kubernetes/test-infra/pull/20836) to General Availability! As a SIG, we will continue to provide review and technical input for future static analysis rule updates.
    *   KEP [2568](https://github.com/kubernetes/enhancements/issues/2568) - Run control-plane as non-root in kubeadm - has been maturing steadily with help from the Kubernetes community. SIG Security has provided a forum to discuss it and feedback to advance it toward implementable status.
    *   KEP [1981](https://github.com/kubernetes/enhancements/issues/1981) - Support for Windows privileged containers - has matured toward alpha, including SIG Security member feedback. Nice work, team!
*   **What areas and/or subprojects does the group need the most help with?**
    *   We always welcome help from anyone aligned with our vision of improving Kubernetes security by learning together, sharing our expertise, and encouraging cross-SIG collaboration. Ideally, we would like to eventually have participating members with crossover to every SIG, providing a natural conduit of ideas and forming a network of helpful hackers throughout Kubernetes.
    *   As a newer SIG with a broad mission, our most pressing needs right now are around outreach and expertise.
        *   Today, people can help broaden our outreach in a few concrete ways:
            *   Look out for one another, so we can be safer together! If you see a Kubernetes PR or Issue (or Enhancement, or whatever!) that looks like it may be security-relevant, you can tag sig/security to bring them to our attention and get SIG Security input. You can also bring them to our meetings as a topic for discussion!
            *   If you spot things happening in other SIGs that may be relevant, or if something that is relevant to security interests is happening in your own SIG, let us know! Post in #sig-security to get community attention on it, or DM the SIG Security chairs on Slack so we can help. 
            *   We don’t always have to be involved; another way you can help is by reaching out directly to other SIGs whose work may be affected by or related to your own. We all benefit when we work together better, especially on security-related topics.
            *   Come to our meetings and get more involved! If you have friends who are interested in Kubernetes security, encourage them to come and bring their energy and ideas too. We’d love to see you and hear what you have to say!
        *   To bring additional expertise to the table and help balance the load for SIG Leads, we would like to grow new Leads in the future, and are determining how best to implement the Tech Lead role in a SIG like ours. We would like to grow one or two interested members into that new role, and will be looking for help with this bootstrapping process when that time comes.
*   **What's the average open days of a PR and Issue in your group?**
    *   As a SIG who doesn’t own code in a project that encourages non-code contributions, we find the question of which PRs and issues belong to us to be difficult to answer, and would like to encourage further inclusivity of different kinds of contribution in questions on future reports.
    *   **What metrics does your group care about and/or measure?**
        *   The metrics we are most immediately interested in are community engagement metrics, as discussed above under “Membership”.
        *   As we grow and mature as a SIG, we seek to better understand how well we are achieving our overall goal of holistically improving Kubernetes security. To do so, we’d like to find metrics that help answer questions like these:
            *   Are the services that we offer to the project achieving their goals?
            *   Are there further services that we should be offering?
        *   As a newer SIG, we know what we care deeply about, and are working toward it. We haven’t entirely figured out how to measure it all yet, but we’ll get there.
