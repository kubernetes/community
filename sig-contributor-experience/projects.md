# Projects and Goals

*note - this is a temporary file until we can figure out a better project management solution.* 

This is a list of the projects and goals currently underway with Contributor Experience. Please submit a PR if you are adding your project to this list. To introduce a new project, attend a weekly meeting or drop a note to us on the mailing list - details can be found on our [README](README.md). 

Want to contribute? Take a look at this list and the "future" at the bottom. Thank you to all of our contributors for your hard work!

## SIG Planning

Project | Owner(s)/Lead(s) | Description | Q1, Q2, Later
---|---|---|---
[Charter](charter.md) | SIG Leads | Create our first Charter iteration | Q1
Projects and Goals (this doc) | @parispittman | Create projects.md | Q1
de-SPOF Community Management | @castrojo | eg: YouTube management | Ongoing

## Mentoring 
Launch a multi-tier strategy test that promote all levels of the [contributor ladder](/community-membership.md) and diversity throughout the project.

Project | Owner(s)/Lead(s) | Description | Q1, Q2, Later
---|---|---|---
[Group Mentoring](/mentoring/group-mentoring.md) | @parispittman | Launch a test cohort that takes 10 current members to reviewer in 3 k8s repos/projects (kops, kubeadm, workload API). Develop the learning and development workshops. | Q1 and ongoing
[Meet Our Contributors](/mentoring/meet-our-contributors.md) | @parispittman | Monthly web series similar to user office hours that allows anyone to ask new and current contributors questions about our process, ecosystem, or their stories in open source  | Q1 - ongoing
[Outreachy](/mentoring/README.md) | @parispittman | Document new features, create new conceptual content, create new user paths | Q1
[Google Summer of Code](/mentoring/google-summer-of-code.md) | @nikhita | Kubernetes participation in Google Summer of Code for students | Q1 - ongoing
["Buddy" Program](https://github.com/kubernetes/community/issues/1803) | @parispittman, @chris-short | 1 hour 1:1 sessions for new and current contributors to have dedicated time; meet our contributors but personal | Q2

## Contributor Documentation
Ensure the contribution process is well documented, discoverable, and consistent across repos to deliver the best contributor experience. 

Project | Owner(s)/Lead(s) | Description | Q1, Q2, Later
---|---|---|---
[Contributor Guide](/contributors/guide) | @castrojo | Make contributor onboarding easier; first version | Q1
[Developer Guide](https://github.com/kubernetes/community/issues/1919) | @ryanj | a comprehensive guide for upstream developers to be a part of the Contributor Guide | Q2
[New Contributor Website](https://github.com/kubernetes/community/issues/1819) | @castrojo | A new home for all things community -documentation, KEPs, Mentoring, + more | Q2 - ongoing
[Label Documentation](https://gist.github.com/spiffxp/24937d8478853054c088ffc298021214) | @spiffxp | GitHub is rolling out a label description feature in the future; this is to document the label descriptions for contributors | Q1
Issue hygiene | @spzala | Produce clear understanding and documentation of issue hygiene | Q1


## Contributor Workflow and Automation
Ensure contributors have a smooth and similar process across repos to deliver the best contributor experience. Provide sufficient automation so that direct write access to repos is no longer required.

Project | Owner(s)/Lead(s) | Description | Q1, Q2, Later
---|---|---|---
Cherry Pick Process | @spiffxp | Design and Prototype | Q2
PR Descriptions | @grodrigues3, @spiffxp | Automation allowing people to edit pr descriptions for tldr or correct issue / pr; release notes based (design).  Define remaining work to be done. | Q1 - ongoing

## Project-wide Communication Channels
Build, curate, moderate, and make project wide communication channels accessible.

Project | Owner(s)/Lead(s) | Description | Q1, Q2, Later
---|---|---|---
[Slack Admin Guidelines](/communication/slack-guidelines.md) | @parispittman | First iteration of guidelines to administer slack for 30k+ users | Q1
[Weekly K8s Community Meeting](/events/community-meeting.md) | @castrojo, @parispittman | Smooth out weekly community meeting and make adjustments | Ongoing
Roadshow | @parispittman | ContribEx to visit all SIG/WG meetings to deliver messages and collect feedback | Ongoing
[Communication Documentation](/communication/README.md) | @parispittman, @castrojo | Create new communication directory and include all information about our communication channels | Q1, Q2

## DevStats
https://k8s.devstats.cncf.io; work devstats into making it the one stop shop for all things upstream Kubernetes data. @phillels is the product owner.

Project | Owner(s)/Lead(s) | Description | Q1, Q2, Later
---|---|---|---
Graph of the Week | @phillels | Present at least two graphs a month with explanation about why we care to the community meeting | Ongoing
[User Guide](https://github.com/cncf/devstats/issues/35) | @jberkus, @parispittman, @tpepper | Create a v1 README for DevStats that explains each chart | Q2
Data validation | @jberkus | Ensure that the data that running through the tool is accurate and aggregating the right information | ongoing

## Events
Curate and produce the best contributor face to face gatherings.

Project | Owner(s)/Lead(s) | Description | Q1, Q2, Later
---|---|---|---
[KubeCon](https://events.linuxfoundation.org/events/kubecon-cloudnativecon-europe-2018/co-located-events/kubernetes-contributor-summit/) | @castrojo, @jberkus, @parispittman | Ensure contributors have great content at the event(s)| Q2, Q4
KubeCon ContribEx Update and Deep Dive | @parispittman, leads | Run the 30 min update and 30 min deep dive at the event | Q2, Q4


## FUTURE
These projects do not have an owner and have not started yet; however, we would like to get them on the slate for sometime this year. 

Project | Description 
---|--- 
Automation of new membership | Current process is emailing a googlegroup and an org owner adds person to GH org
Audit schedule | Create an audit schedule to make sure all communication pipelines are up and running smoothly eg: check that google group mailing lists have back up owners/maintainers, make sure meetings are following the calendar process, etc.
