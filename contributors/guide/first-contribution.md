---
title: "Making your First Contribution"
weight: 3
description: |
  Not sure where to make your first contribution? This doc has some tips and
  ideas to help get you started.
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

## Find something to work on

The first step to getting starting contributing to Kubernetes is to find something
to work on. Help is always welcome, and no contribution is too small! 

Here are some things you can do today to get started contributing:

* Help improve the Kubernetes documentation
* Clarify code, variables, or functions that can be renamed or commented on
* Write test coverage
* Help triage issues

If the above suggestions don't appeal to you, you can browse the 
[Contributor Role Board] to see who is looking for help. Those interested 
in contributing without writing code can also find ideas in the 
[Non-Code Contributions Guide].

### Find a good first topic

There are [multiple repositories] within the Kubernetes organization.
Each repository has beginner-friendly issues that are a great place to 
get started on your contributor journey. For example, [kubernetes/kubernetes] has 
[help wanted] and [good first issue] labels for issues that don't need high-level 
Kubernetes knowledge to contribute to. The `good first issue` label also indicates 
that Kubernetes Members have committed to providing [extra assistance] for new 
contributors. Another way to get started is to find a documentation improvement, 
such as a missing/broken link, which will give you exposure to the code 
submission/review process without the added complication of technical depth. 

### Issue Assignment in Github

When you've found an issue to work on, you can assign it to yourself.

* Reply with `/assign` or `/assign @yourself` on the issue you'd like to work on 
* The [K8s-ci-robot] will automatically assign the issue to you. 
* Your name will then be listed under, `Assignees`.

## Learn about SIGs

Some repositories in the Kubernetes Organization are owned by 
[Special Interest Groups], or SIGs.

The Kubernetes community is broken out into SIGs in order to improve its workflow,
and more easily manage what is a very large community project. The developers 
within each SIG have autonomy and ownership over that SIG's part of Kubernetes. 
Understanding how to interact with SIGs is an important part of contributing 
to Kubernetes. Check out the [list of SIGs][sl] for contact information.

### SIG structure

A SIG is an open, community effort.

Anybody is welcome to jump into a SIG and begin fixing issues, critiqe design 
proposals, and review code. SIGs have regular [video meetings] which everyone 
is welcome to attend. Each SIG has a Slack channel, meeting notes, and their own 
documentation that is useful to read and understand. There is an entire SIG 
([sig-contributor-experience]) devoted to improving your experience as a contributor. 
If you have an idea for  improving the contributor experience, please consider 
attending one of the Contributor Experience SIG's [weekly meetings].

### Find a SIG that is related to your contribution

Finding the appropriate SIG for your contribution and adding a SIG label will 
help you ask questions in the correct place and give your contribution higher 
visibility and a faster community response.

For Pull Requests, the automatically assigned reviewer will add a SIG label 
if you haven't already done so. 

For Issues, please note that the community is working on a more automated workflow.
Since SIGs do not directly map onto Kubernetes subrepositories, it may be 
difficult to find which SIG your contribution belongs in. Review the 
[list of SIGs][sl] to determine which SIG is most likely related to your 
contribution.

*Example:* if you are filing a CNI issue (that's [Container Networking Interface]) 
you'd choose the [Network SIG]. Add the SIG label in a new comment on GitHub 
by typing the following:
```
/sig network
```

Follow the link in the SIG name column to reach each SIG'ss README. 

Most SIGs will have a set of GitHub Teams with tags that can be mentioned in a 
comment on issues and pull requests for higher visibility.  If you are not sure 
about the correct SIG for an issue, you can try [SIG-contributor-experience], 
or [ask in Slack].

### SIG-specific contributing guidelines

Some SIGs have their own `CONTRIBUTING.md` files, which may contain extra information 
or guidelines in addition to these general ones. These are located in the SIG-specific 
community directories:

- [`/sig-apps/CONTRIBUTING.md`](/sig-apps/CONTRIBUTING.md)
- [`/sig-cli/CONTRIBUTING.md`](/sig-cli/CONTRIBUTING.md)
- [`/sig-multicluster/CONTRIBUTING.md`](/sig-multicluster/CONTRIBUTING.md)
- [`/sig-storage/CONTRIBUTING.md`](/sig-storage/CONTRIBUTING.md)
- [`/sig-windows/CONTRIBUTING.md`](/sig-windows/CONTRIBUTING.md)

### File an Issue

Not ready to contribute code, but see something that needs work?
While the community encourages everyone to contribute code, it is also appreciated 
when someone reports an issue. Issues should be filed under the appropriate Kubernetes 
subrepository. For example, a documentation issue should be opened in 
[kubernetes/website]. Make sure to adhere to the prompted submission guidelines 
while opening an issue. Check the [issue triage guide] for more information.

[Contributor Role Board]: https://discuss.kubernetes.io/c/contributors/role-board
[k8s-ci-robot]: https://github.com/k8s-ci-robot
[Non-Code Contributions Guide]: ./non-code-contributions.md
[multiple repositories]: https://github.com/kubernetes/
[kubernetes/kubernetes]: https://git.k8s.io/kubernetes
[help wanted]: https://go.k8s.io/help-wanted
[good first issue]: https://go.k8s.io/good-first-issue
[extra assistance]:./help-wanted.md
[sl]: /sig-list.md
[video meetings]: https://kubernetes.io/community/
[sig-contributor-experience]: /sig-contributor-experience/README.md
[weekly meetings]: https://docs.google.com/document/d/1qf-02B7EOrItQgwXFxgqZ5qjW0mtfu5qkYIF1Hl4ZLI/edit
[container networking interface]: https://github.com/containernetworking/cni
[network SIG]: http://git.k8s.io/community/sig-network
[ask in Slack]: http://slack.k8s.io/
[issue triage guide]: ./issue-triage.md
[kubernetes/website]: https://github.com/kubernetes/website/issues
[SIG Contributor Experience]: /sig-contributor-experience#contact






