# Kubernetes New Contributor Workshop - KubeCon/CloudNativeCon EU 2018 - Notes

Joining in the beginning was onboarding on a yacht
Now is more onboarding a BIG cruise ship.

Will be a Hard schedule, and let's hope we can achieve everything
Sig-contributor-experience -> from Non-member contributors to Owner

## SIG presentation

- SIG-docs & SIG-contributor-experience: **Docs and website** contribution
- SIG-testing: **Testing** contribution
- SIG-\* (*depends on the area to contribute on*): **Code** contribution

**=> Find your first topics**: bug, feature, learning, community development and documentation

Table exercise: Introduce yourself and give a tip on where you want to contribute in Kubernetes


## Communication in the community

Kubernetes community is like a Capybara: community members are really cool with everyone and they are from a lot of different horizons.

- Tech question on Slack and Stack Overflow, not on Github
- A lot of discussion will be involve when GH issues and PR are opened. Don't be frustrated
- Stay patient because there is a lot of contribution

When in doubt, **ask on Slack**

Other communication channels:

- Community meetings
- Mailing lists
- @ on Github
- Office Hour
- Kubernetes meetups https://www.meetup.com/topics/kubernetes

on https://kubernetes.io/community,  there is the schedule for all the SIG/Working group meeting.
If you want to join or create a meetup. Go to **slack#sig-contribex**

## SIG - Special Interest Group

Semi-autonomous teams:
- Own leaders & charteers
- Code, Github repo, Slack, mailing, meeting responsibility

### Types

[SIG List](https://github.com/kubernetes/community/blob/master/sig-list.md)

1. Features Area
    - sig-auth
    - sig-apps
    - sig-autoscaling
    - sig-big-data
    - sig-cli
    - sig-multicluster
    - sig-network
    - sig-node
    - sig-scalability
    - sig-scheduling
    - sig-service-catalog
    - sig-storage
    - sig-ui
2. Plumbing
    - sig-cluster-lifecycle
    - sig-api-machinary
    - sig-instrumentation
3. Cloud Providers *(currently working on moving cloudprovider code out of Core)*
    - sig-aws
    - sig-azure
    - sig-gcp
    - sig-ibmcloud
    - sig-openstack
4. Meta
    - sig-architecture: For all general architectural decision
    - sig-contributor-experience: Helping contributor and community experience
    - sig-product-management: Long-term decision
    - sig-release
    - sig-testing: In charge of all the test for Kubernetes
5. Docs
    - sig-docs: for documentation and website

## Working groups and "Subproject"

From working group to "subproject". 

For specific: tools (ex. Helm), goals (ex. Resource Management) or areas (ex. Machine Learning).

Working groups change around more frequently than SIGs, and some might be temporary.

- wg-app-def
- wg-apply
- wg-cloud-provider
- wg-cluster-api
- wg-container-identity
- ...

### Picking the right SIG:
1. Figure out which area you would like to contribute to
2. Find out which SIG / WG / subproject covers that (tip: ask on #sig-contribex Slack channel)
3. Join that SIG / WG / subproject (you should also join the main SIG when joining a WG / subproject)

## Tour des repositories

Everything will be refactored (cleaning, move, merged,...)

### Core repository 
- [kubernetes/kubernetes](https://github.com/kubernetes/kubernetes)

### Project

- [kubernetes/Community](https://github.com/kubernetes/Community): KubeCon/CloudNativeCon, proposition, Code of conduct and Contribution guideline, SIG-list
- [kubernetes/Features](https://github.com/kubernetes/Features): Features proposal for future release
- [kubernetes/Steering](https://github.com/kubernetes/Steering)
- [kubernetes/Test-Infra](https://github.com/kubernetes/Test-Infra): All related to test except Perf
- [kubernetes/Perf-Tests](https://github.com/kubernetes/Perf-Tests): 

### Docs/Website

- website
- kubernetes-cn
- kubernetes-ko

### Developer Tools

- sample-controller*
- sample- apiserver*
- code-generator*
- k8s.io
- kubernetes-template-project: For new github repo

### Staging repositories

Mirror of core part for easy vendoring

### SIG repositories

- release
- federation
- autoscaler

### Cloud Providers

No AWS

### Tools & Products

- kubeadm
- kubectl
- kops
- helm
- charts
- kompose
- ingress-nginx
- minikube
- dashboard
- heapster
- kubernetes-anywhere
- kube-openapi

### 2nd Namespace: Kubernetes-sigs

Too much places for Random/Incubation stuff.
No working path for **promotion/deprecation**

In future:
1. start in Kubernetes-sigs
2. SIGs determine when and how the project will be **promoted/deprecated**

Those repositories can have their own rules:
- Approval
- Ownership
- ...

## Contribution

### First Bug report

```
- Bug or Feature

- What happened

- How to reproduce

```

 ### Issues as specifications


Most of k8s change start with an issue:

- Feature proposal
- API changes proposal
- Specification

### From Issue to Code/Docs

1. Start with an issue
2. Apply all appropriate labels
3. cc SIG leads and concerned devs
4. Raise the issue at a SIG meeting or on mailing list
5. If *Lazy consensus*, submit a PR

### Required labels https://github.com/kubernetes/test-infra/blob/master/label_sync/labels.md

#### On creation
- `sig/\*`: the sig the issue belong too
- `kind/\*`: 
    - bug
    - feature
    - documentation
    - design
    - failing-test

#### For issue closed as port of **triage**

- `triage/duplicate`
- `triage/needs-information`
- `triage/support`
- `triage/unreproduceable`
- `triage/unresolved`

#### Prority

- `priority/critical-urgent`
- `priority/important-soon`
- `priority/important-longtem`
- `priority/backlog`
- `priority/awaiting-evidence`

#### Area

Free for dedicated issue area

- `area/kubectl`
- `area/api`
- `area/dns`
- `area/platform/gcp`

#### help-wanted

Currently mostly complicated things

#### SOON 

`good-first-issue`

## Making a contribution by Pull Request

We will go through the typical PR process on kubernetes repos.

We will play there: [community/contributors/new-contributor-playground at master · kubernetes/community · GitHub](https://github.com/kubernetes/community/tree/master/contributors/new-contributor-playground)

1. When we contribute to any kubernetes repository, **fork it**

2. Do your modification in your fork
```
$ git clone git@github.com:jgsqware/community.git $GOPATH/src/github.com/kubernetes/community
$ git remote add upstream https://github.com/kubernetes/community.git
$ git remote -v
origin  git@github.com:jgsqware/community.git (fetch)
origin  git@github.com:jgsqware/community.git (push)
upstream        git@github.com:kubernetes/community.git (fetch)
upstream        git@github.com:kubernetes/community.git (push)
$ git checkout -b kubecon
Switched to a new branch 'kubecon'

## DO YOUR MODIFCATION IN THE CODE##

$ git add contributors/new-contributor-playground/new-contibutor-playground-xyz.md
$ git commit


### IN YOUR COMMIT EDITOR ###

    Adding a new contributors file

    We are currently experimenting PR process in the kubernetes repository.

$ git push -u origin kubecon
```

3. Create a Pull request via Github
4. If needed, sign the CLA to make valid your contribution
5. Read the `k8s-ci-robot` message and `/assign @reviewer` recommended by the `k8s-ci-robot`
6. wait for a `LTGM` label from one of the `OWNER/reviewers` 
7. wait for approval from one of `OWNER/approvers`
8. `k8s-ci-robot` will automatically merge the PR

`needs-ok-to-test` is used for non-member contributor to validate the pull request

## Test infrastructure

> How bot toll you when you mess up

At the end of a PR there is a bunch of test. 
2 types:
  - required: Always run and needed to pass to validate the PR (eg. end-to-end test)
  - not required: Needed in specific condition (eg. modifying on ly specific part of code)

If something failed, click on `details` and check the test failure logs to see what happened. 
There is `junit-XX.log` with the list of test executed and `e2e-xxxxx` folder with all the component logs.
To check if the test failed because of your PR or another one, you can click on the **TOP** `pull-request-xxx` link and you will see the test-grid and check if your failing test is failing in other PR too.

If you want to retrigger the test manually, you can comment the PR with `/retest` and `k8s-ci-robot` will retrigger the tests.

## SIG-Docs contribution

Anyone can contribute to docs.

### Kubernetes docs

- Websites URL
- Github Repository
- k8s slack: #sig-docs

### Working with docs

Docs use `k8s-ci-robot`. Approval process is the same as for any k8s repo.
In docs, `master` branch is the current version of the docs. So always branch from `master`. It's continuous deployment
For a specific release docs, branch from `release-1.X`.

## Local build and Test

The code: [kubernetes/kubernetes]
The process: [kubernetes/community]

### Dev Env

You need:
- Go
- Docker


- Lot of RAM and CPU and 10 GB of space
- best to use Linux
- place you k8s repo fork in:
  - `$GOPATH/src/k8s.io/kubernetes`
- `cd $GOPATH/src/k8s.io/kubernetes`
- build: `./build/run.sh make`  
  - Build is incremental, keep running `./build/run.sh make` til it works
- To build variant: `make WHAT="kubectl"`
- Building kubectl on Mac for linux: `KUBE_*_PLATFORM="linux/amd64" make WHAT "kubectl"`

There is `build` documentation there: https://git.k8s.io/kubernetes/build

### Testing
There is `test` documentation there: https://git.k8s.io/community/contributor/guide
