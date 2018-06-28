# IMPORTANT - The Kubernetes Incubator process is now deprecated and has been superseded by Kubernetes subprojects

For information on creating a repository for a subproject see: [kubernetes-repositories](/github-management/kubernetes-repositories.md)

Each SIG should define the process for sponsoring new subprojects in its charter.  For information on SIG governance
and charters see: [SIG governance](committee-steering/governance/README.md)

# Kubernetes Incubator Process

**Authors:** Brandon Philips <brandon.philips@coreos.com>, Sarah Novotny <sarahnovotny@google.com>, Brian Grant <briangrant@google.com>

## Kubernetes Incubator

The [Kubernetes Incubator](https://github.com/kubernetes-incubator) is where all new Kubernetes projects should start. Once going through the incubation process explained in this doc they can become full Kubernetes Community Projects. The process is designed to ensure that new projects have correct licensing, up-to-date boilerplate documents, a healthy community process, and are developed using accepted Kubernetes Community practices.

### Context

New Kubernetes projects are getting added in an ad-hoc manner leading to confusion; see [this thread as an example](https://groups.google.com/forum/#!topic/kubernetes-dev/o6E1u-orDK8).

### Status 

This process is light on legalese, completely untested, and only works if people act as good neighbors and community members. It will evolve over time and the authors will try keeping the process light, fast, and objective.

## Does my project need to Incubate?

### Existing Code in Kubernetes

New Kubernetes Community Projects can be created by pulling out code from the Kubernetes main repo (except contrib/). For example, the OWNERS of `k8s.io/kubernetes/pkg/cloudprovider/providers/aws` want to start a new project so they can iterate and release faster than k8s itself. To do this the OWNERS of that package can directly graduate to a Kubernetes Community Project with agreement of the existing package OWNERS, agreement of the parent package (e.g. `k8s.io/kubernetes/pkg/cloudprovider`), and an announcement to kubernetes-dev@googlegroups.com.

### Applications on top of Kubernetes

If you are building an Open Source application on top of Kubernetes that is unrelated to the development, operation, monitoring, APIs, lifecycle, integration, or other primary concerns of the Kubernetes project there is no reason to become a Kubernetes project. If you are unsure that your project is a good fit try running it by a few Kubernetes Special Interest Groups or kubernetes-dev@googlegroups.com.

## Entering Incubation

To create a new project for incubation you must follow these steps: write a proposal, find a champion, gain acceptance into a SIG, and finally get approval of a Sponsor. The Sponsor and Champion cannot be the same person.

Your proposal should include two items. First, a README which outlines the problem to be solved, an example use case as-if the project existed, and a rough roadmap with timelines. Second, an OWNERS file that outlines the makeup of the initial team developing the project. Initially this can be one person but ideally has 3 or more initial developers representing a few different companies or groups. You can use whatever tool you want to host and revise the proposal until the project is accepted to the Incubator: copy/paste from your text editor, Google Docs, GitHub gist, etc.

Once the proposal is written you should identify a champion; this person must be listed as either a reviewer or approver in an [OWNERS file](https://git.k8s.io/kubernetes/OWNERS) in the Kubernetes project. Next, reach out to your potential champion via email to ask if they are interested in helping you through the Incubation process. Ideally some significant follow-up discussion happens via email, calls, or chat to improve the proposal before announcing it to the wider community.

The next discussion should be on a relevant Special Interest Group mailing list. You should post the proposal to the SIG mailing list and wait for discussion for a few days. Iterate on the proposal as needed and if there is rough consensus that the project belongs in the chosen SIG then list that SIG in the proposal README. If consensus isn't reached then identify another SIG and try again; repeat until a SIG is identified.

The final process is to email kubernetes-dev@googlegroups.com to announce your intention to form a new Incubator project. Include your entire proposal in the body of the email and prefix the Subject with [Incubator]. Include links to your discussion on the accepted SIG mailing list to guide the discussion.

Acceptance of the project into the Kubernetes Incubator happens once a Sponsor approves. Anyone listed as an approver in the top-level pkg [OWNERS file](https://git.k8s.io/kubernetes/pkg/OWNERS) can sponsor a project by replying to the kubernetes-dev discussion with LGTM.

## Creation of the Incubator Project

Once your project is accepted someone will create a new GitHub repo for your project to use in the [kubernetes-incubator](https://github.com/kubernetes-incubator) organization. The first order of business is to add the following files to the repo:

- a README from the accepted proposal
- an OWNERS from the accepted proposal
- a CONTRIBUTING file based on kubernetes/kubernetes
- a code-of-conduct.md based on kubernetes/code-of-conduct.md
- a LICENSE file with the Apache 2.0 license
- a RELEASE.md file that talks about the process for releases

To start your project with these files can be found here at [template project](https://github.com/kubernetes/kubernetes-template-project)

## Exiting Incubation

An Incubator Project must exit 12 months from the date of acceptance in one of these three ways:

- Graduate as a new Kubernetes Project
- Merge with another Kubernetes Project
- Retirement

### Graduation

Both time and traction are required to graduate to becoming a new Kubernetes Community Project. It is expected that from the time that your project is accepted to a request for graduation the following criteria are met:

- **Documented users:** the project tracks evidence from downloads, mailing lists, Twitter, blogs, GitHub issues, etc that the project has gained a user base
- **Regular releases:** the project is making releases at least once every 3 months; although release should happen more often
- **Incubation time:** a minimum of 6 months has passed since
- **Healthy community:** a healthy number of users, contributors, and/or OWNERS outside of your own company have participated in the project. The Sponsor and Champion can make this call.
- **Roadmap:** A roadmap for the next release is maintained; this can be in the form of a doc, GitHub milestones, or any other tool but it must be documented in the README

When the OWNERS of the Incubator project determine they have met the criteria to graduate they should contact their Champion and Sponsor to discuss. If both the Sponsor and Champion agree that the above criteria have been met the project can graduate to become a Kubernetes Community Project and exit the Kubernetes Incubator.

An announcement of graduation must be sent to the kubernetes-dev@googlegroups.com mailing list by the Champion.

### Merge

At any point during Incubation a project can exit by merging the project with another Incubator Project or Kubernetes Community Project. The process to merge includes:

1. Incubator Project OWNERS notifying the project Champion and Sponsor
2. Once the Champion and Sponsor have been informed then email kubernetes-dev@googlegroups.com about the intention to exit through a merge

### Retirement

If a project doesn't merge or graduate within 12 months it is retired. If a project fails to make a release for 6 months it is retired. Retired repos are moved to the github.com/kubernetes-incubator-retired organization. A warning email will be sent 3 weeks before the move to all people listed in the root OWNERS file, the Champion, and the Sponsor. A project can be re-incubated after retirement but must go through the process of Incubation Entry from the beginning, the same as any new project.

## FAQ

**Q: What is the role of a Champion?**

**A:** Potential Champions come from the set of all Kubernetes approvers and reviewers with the hope that they will be able to teach the Incubator Project about Kubernetes community norms and processes. The Champion is the primary point of contact for the Incubator Project team; and will help guide the team through the process. The majority of the mentorship, review, and advice will come from the Champion. Being a Champion is a significant amount of work and active participation in the sponsored project is encouraged.

**Q: What is the role of the Sponsor?**

**A:** Potential Sponsors come from the very small set of Kubernetes contributors that can approve any PR because they are listed as approvers in the `kubernetes/pkg` OWNERS file or the top-level OWNERS file. The idea is that by relying on this small set of Kubernetes Community members to make a determination on Incubator projects we will ensure that there is consistency around new projects joining the Incubator. Being a Sponsor is a minor advisory role.

## Existing Repos

Based on the above process there are a number of repos under github.com/kubernetes we should handle through either grandfathering, a move to incubation, or retirement. This list is a rough draft, if you feel strongly about something being miscategorized please comment.

**Projects to Make "Kubernetes Community Projects"**

These are grandfathered in as full projects:

- Kubernetes
- Heapster
- cAdvisor (in-progress move to github.com/kubernetes)
- github.com/kubernetes/kubernetes.github.io
- github.com/kubernetes/test-infra
- github.com/kubernetes/community
- github.com/kubernetes/release
- github.com/kubernetes/features
- github.com/kubernetes/kube-state-metrics
- github.com/kubernetes/pr-bot - move from mungebot, etc from contrib, currently running in "prod" on github.com/kubernetes
- github.com/kubernetes/dashboard
- github.com/kubernetes/helm  (Graduated from incubator on Feb 2017)
- github.com/kubernetes/minikube (Graduated from incubator on Feb 2017)
- github.com/kubernetes/kops (Graduated from incubator in Jun 2017)

**Project to Incubate But Not Move**

These projects are young but have significant user facing docs pointing at their current github.com/kubernetes location. Lets put them through incubation process but leave them at github.com/kubernetes.

- github.com/kubernetes/charts
 
**Projects to Move to Incubator**

- github.com/kubernetes/kube2consul
- github.com/kubernetes/frakti
- github.com/kubernetes/kube-deploy
- github.com/kubernetes/kubernetes-anywhere
- github.com/kubernetes/application-images
- github.com/kubernetes/rktlet
- github.com/kubernetes/horizontal-self-scaler
- github.com/kubernetes/node-problem-detector
- github.com/kubernetes/kubernetes/contrib/mesos -> github.com/kubernetes-incubator/kube-mesos-framework

**Projects to Retire**

- github.com/kubernetes/kube-ui
- github.com/kubernetes/contrib - new projects could be created from the code in this repo, issue to move
- github.com/kubernetes/kubedash
- github.com/kubernetes/kubernetes-docs-cn
- github.com/kubernetes/md-check

## Thank You

Large portions of this process and prose are inspired by the Apache Incubator process.

## Original Discussion
https://groups.google.com/d/msg/kubernetes-dev/o6E1u-orDK8/SAqal_CeCgAJ

## Future Work

- Expanding potential sources of champions outside of Kubernetes main repo
