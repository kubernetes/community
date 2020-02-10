---
title: "Making your First Contribution"
weight: 3
slug: "first-contribution" 
---

# Your First Contribution

- [Find something to work on](#find-something-to-work-on)
  - [Find a good first topic](#find-a-good-first-topic)
    - [Issue Assignment in Github](#issue-assignment-in-github)
  - [Learn about SIGs](#learn-about-sigs)
    - [SIG structure](#sig-structure)
    - [Find a SIG that is related to your contribution](#find-a-sig-that-is-related-to-your-contribution)
  - [SIG-specific contributing guidelines](#sig-specific-contributing-guidelines)
  - [File an Issue](#file-an-issue)


Have you ever wanted to contribute to the coolest cloud technology?
We will help you understand the organization of the Kubernetes project and direct you to the best places to get started.
You'll be able to pick up issues, write code to fix them, and get your work reviewed and merged.

Please be aware that due to the large number of issues our triage team deals with, we cannot offer technical support in GitHub issues.
If you have questions about the development process, feel free to jump into our [Slack Channel](http://slack.k8s.io/) or join our [mailing list](https://groups.google.com/forum/#!forum/kubernetes-dev).
You can also ask questions on [ServerFault](https://serverfault.com/questions/tagged/kubernetes) or [Stack Overflow](https://stackoverflow.com/questions/tagged/kubernetes).
The Kubernetes team scans Stack Overflow on a regular basis and will try to ensure your questions don't go unanswered.

## Find something to work on

Help is always welcome! For example, documentation (like the text you are reading now) can always use improvement. 
There's always code that can be clarified and variables or functions that can be renamed or commented.
There's always a need for more test coverage.
You get the idea - if you ever see something you think should be fixed, you should own it.
Here is how you get started.
If you have no idea what to start on, you can browse the [Contributor Role Board](https://discuss.kubernetes.io/c/contributors/role-board) to see who is looking for help.
Those interested in contributing without writing code may also find ideas in the [Non-Code Contributions Guide](non-code-contributions.md).

### Find a good first topic

There are [multiple repositories](https://github.com/kubernetes/) within the Kubernetes organization.
Each repository has beginner-friendly issues that provide a good first issue.
For example, [kubernetes/kubernetes](https://git.k8s.io/kubernetes) has [help wanted](https://go.k8s.io/help-wanted) and [good first issue](https://go.k8s.io/good-first-issue) labels for issues that should not need deep knowledge of the system.
The `good first issue` label indicates that members have committed to providing [extra assistance](/contributors/guide/help-wanted.md) for new contributors.
<!-- TODO: review removing this note after 3 months or after the 1.12 release -->
Please note that while several of the repositories in the Kubernetes community have `good first issue` labels already, they are still being applied throughout the community.

Another good strategy is to find a documentation improvement, such as a missing/broken link, which will give you exposure to the code submission/review process without the added complication of technical depth. Please see [Contributing](#contributing) below for the workflow.

#### Issue Assignment in Github

When you are willing to take on an issue, you can assign it to yourself. Just reply with `/assign` or `/assign @yourself` on an issue, 
then the robot will assign the issue to you and your name will present at `Assignees` list.

### Learn about SIGs

#### SIG structure

You may have noticed that some repositories in the Kubernetes Organization are owned by Special Interest Groups, or SIGs.
We organize the community into SIGs in order to improve our workflow and more easily manage what is a very large community project.
The developers within each SIG have autonomy and ownership over that SIG's part of Kubernetes.

A SIG is an open, community effort.
Anybody is welcome to jump into a SIG and begin fixing issues, critiquing design proposals and reviewing code.
SIGs have regular [video meetings](https://kubernetes.io/community/) which everyone is welcome to.
Each SIG has a slack channel that you can join as well.

There is an entire SIG ([sig-contributor-experience](/sig-contributor-experience/README.md)) devoted to improving your experience as a contributor.
Contributing to Kubernetes should be easy.
If you find a rough edge, let us know! Better yet, help us fix it by joining the SIG; just
show up to one of the [bi-weekly meetings](https://docs.google.com/document/d/1qf-02B7EOrItQgwXFxgqZ5qjW0mtfu5qkYIF1Hl4ZLI/edit).

#### Find a SIG that is related to your contribution

Finding the appropriate SIG for your contribution and adding a SIG label will help you ask questions in the correct place and give your contribution higher visibility and a faster community response.

For Pull Requests, the automatically assigned reviewer will add a SIG label if you haven't done so. See [Open A Pull Request](#open-a-pull-request) below.

For Issues, we are still working on a more automated workflow.
Since SIGs do not directly map onto Kubernetes subrepositories, it may be difficult to find which SIG your contribution belongs in.
Here is the [list of SIGs](/sig-list.md) so that you can determine which is most likely related to your contribution.

*Example:* if you are filing a CNI issue (that's [Container Networking Interface](https://github.com/containernetworking/cni)), you should choose the [Network SIG](http://git.k8s.io/community/sig-network). Add the SIG label in a comment like so:
```
/sig network
```

Follow the link in the SIG name column to reach each SIGs README. 
Most SIGs will have a set of GitHub Teams with tags that can be mentioned in a comment on issues and pull requests for higher visibility. 
If you are not sure about the correct SIG for an issue, you can try SIG-contributor-experience [here](/sig-contributor-experience#github-teams), or [ask in Slack](http://slack.k8s.io/).

### SIG-specific contributing guidelines
Some SIGs have their own `CONTRIBUTING.md` files, which may contain extra information or guidelines in addition to these general ones.
These are located in the SIG-specific community directories:

- [`/sig-apps/CONTRIBUTING.md`](/sig-apps/CONTRIBUTING.md)
- [`/sig-cli/CONTRIBUTING.md`](/sig-cli/CONTRIBUTING.md)
- [`/sig-multicluster/CONTRIBUTING.md`](/sig-multicluster/CONTRIBUTING.md)
- [`/sig-storage/CONTRIBUTING.md`](/sig-storage/CONTRIBUTING.md)
- [`/sig-windows/CONTRIBUTING.md`](/sig-windows/CONTRIBUTING.md)

### File an Issue

Not ready to contribute code, but see something that needs work?
While the community encourages everyone to contribute code, it is also appreciated when someone reports an issue (aka problem).
Issues should be filed under the appropriate Kubernetes subrepository.
Check the [issue triage guide](./issue-triage.md) for more information.

*Example:* a documentation issue should be opened to [kubernetes/website](https://github.com/kubernetes/website/issues).

Make sure to adhere to the prompted submission guidelines while opening an issue.