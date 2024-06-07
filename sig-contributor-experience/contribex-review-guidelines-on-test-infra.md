# Guiding policy for test-infra changes requiring review by SIG Contributor Experience

This document serves as a guiding policy to help identify which changes in the `git.k8s.io/test-infra/` repository require a review from the Technical Leads of SIG Contributor Experience.

It also outlines areas that are outside the scope of SIG Contributor Experience's review.

Note: This document is not a binding or an exhaustive list, but aims to provide general guidance.

---

## Index

1. [When a Review from SIG Contributor Experience is Required](#when-a-review-from-sig-contributor-experience-is-required)
2. [Recommended Areas for ContribEx Review](#recommended-areas-for-contribex-review)
    1. [Changes to `config/plugins.yaml`](#changes-to-configpluginsyaml)
        - [Plugins under ContribEx Review](#plugins-under-contribex-review)
        - [Plugins Configured for Specific Orgs/Repos](#plugins-configured-for-specific-orgsrepos)
        - [External Plugins Configured for Specific Orgs/Repos](#external-plugins-configured-for-specific-orgsrepos)
    2. [Changes to `config/config.yaml`](#changes-to-configconfigyaml)
    3. [Changes to `config/prow/OWNERS` File](#changes-to-configprowowners-file)
    4. [Changes to Prow CI Jobs and Their OWNERS Files](#changes-to-prow-ci-jobs-and-their-owners-files)
    5. [Changes to `Label_sync` Path](#changes-to-label_sync-path)
3. [Out of Scope](#out-of-scope)

---

## When a Review from SIG Contributor Experience is Required

A review from SIG Contributor Experience is necessary if any changes introduced by a pull request to the test-infra repo:

- Introduce changes to how humans interact with Prow and its components (changes that alter contributor user experience)
- Introduce changes to how bot interact with Prow? (TODO: verify!)
- Introduce changes to how Prow interacts with GitHub, or what information it broadcasts back to GitHub (in form of comments, etc, for example)
- Alter the contribution workflows for project contributors.

Examples include – adding new labels like `kind/foo` enabled/required by default, or changes that modify the issue/PR triage workflow used by project maintainers.

---

## Recommended Areas for ContribEx Review

### Changes to [`config/prow/plugins.yaml`](https://github.com/kubernetes/test-infra/blob/52fb24b56def58e3cc36b10802aa6dea230584a7/config/prow/plugins.yaml)

#### Plugins under ContribEx Review

Following plugins fall under ContribEx review:

- trigger
- owners
- approve
- help
- size
- label
- blockades
- slack
- welcome
- require_matching_label

#### Plugins Configured for Specific Orgs/Repos

**Organizations:**
- kubernetes
- kubernetes-client
- kubernetes-sigs
- etcd-io

**Repositories:**
- kubernetes/kubernetes
- kubernetes/community
- kubernetes/org
- kubernetes/publishing-bot
- kubernetes-sigs/contributor-playground

#### External Plugins Configured for Specific Orgs/Repos

**Organizations:**
- kubernetes
- kubernetes-sigs
- kubernetes-csi
- kubernetes-client

---

### Changes to [`config/prow/config.yaml`](https://github.com/kubernetes/test-infra/blob/52fb24b56def58e3cc36b10802aa6dea230584a7/config/prow/config.yaml)

- [slack_reporter_configs](https://github.com/kubernetes/test-infra/blob/52fb24b56def58e3cc36b10802aa6dea230584a7/config/prow/config.yaml#L215-L226)
- [branch-protection](https://github.com/kubernetes/test-infra/blob/52fb24b56def58e3cc36b10802aa6dea230584a7/config/prow/config.yaml#L228-L636)
- [tide](https://github.com/kubernetes/test-infra/blob/master/config/prow/config.yaml#L638-L844)
- [github_reporter](https://github.com/kubernetes/test-infra/blob/52fb24b56def58e3cc36b10802aa6dea230584a7/config/prow/config.yaml#L850-L853)

---

### Changes to [`config/prow/OWNERS`](https://github.com/kubernetes/test-infra/blob/52fb24b56def58e3cc36b10802aa6dea230584a7/config/prow/OWNERS#L1) file

Any changes to the reviewers/approvers or filters listed in the `config/prow/OWNERS` file.

---

### Changes to Prow CI Jobs and Their OWNERS Files

Changes to the definition of the following Prow CI jobs files as well as their respective OWNERS files, owned directly or indirectly by SIG ContribEx:

- All files under [kubernetes/community/](https://github.com/kubernetes/test-infra/tree/master/config/jobs/kubernetes/community) path
- All files under [kubernetes/org/](https://github.com/kubernetes/test-infra/tree/master/config/jobs/kubernetes/org) path
- [kubernetes/sig-k8s-infra/trusted/sig-contribex-k8s-triage-robot.yaml](https://github.com/kubernetes/test-infra/blob/master/config/jobs/kubernetes/sig-k8s-infra/trusted/sig-contribex-k8s-triage-robot.yaml)
- [kubernetes/sig-k8s-infra/trusted/sig-contribex-peribolos.yaml](https://github.com/kubernetes/test-infra/blob/master/config/jobs/kubernetes/sig-k8s-infra/trusted/sig-contribex-peribolos.yaml)
- [kubernetes/sig-k8s-infra/trusted/sig-k8s-infra-groups.yaml](https://github.com/kubernetes/test-infra/blob/master/config/jobs/kubernetes/sig-k8s-infra/trusted/sig-k8s-infra-groups.yaml)
- [config/jobs/kubernetes-sigs/slack-infra/slack-infra-config.yaml](https://github.com/kubernetes/test-infra/blob/master/config/jobs/kubernetes-sigs/slack-infra/slack-infra-config.yaml)

---

### Changes to `Label_sync` Path

Any changes to the files under [`label_sync`](https://github.com/kubernetes/test-infra/tree/master/label_sync) path.

---

## Out of Scope

The following are outside the scope of review by SIG Contributor Experience:

- Technical changes to the source codebase of Prow components (e.g., `prow/`, now moved to `sigs.k8s.io/prow`, `kubetest`, `kettle`, etc.).
- Technical changes to External plugins
