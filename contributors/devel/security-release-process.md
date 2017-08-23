# Security Release Process

Kubernetes is a large growing community of volunteers, users, and vendors. The Kubernetes community has adopted this security disclosures and response policy to ensure we responsibly handle critical issues.

## Product Security Team (PST)

Security vulnerabilities should be handled quickly and sometimes privately. The primary goal of this process is to reduce the total time users are vulnerable to publicly known exploits.

The Product Security Team (PST) is responsible for organizing the entire response including internal communication and external disclosure but will need help from relevant developers and release managers to successfully run this process.

The initial Product Security Team will consist of four volunteers subscribed to the private [Kubernetes Security](https://groups.google.com/forum/#!forum/kubernetes-security) list. These are the people who have been involved in the initial discussion and volunteered:

- Brandon Philips `<brandon.philips@coreos.com>` [4096R/154343260542DF34]
- Jess Frazelle `<jess@linux.com>` [4096R/0x18F3685C0022BFF3]
- CJ Cullen `<cjcullen@google.com>`
- Tim Allclair `<tallclair@google.com>` [4096R/0x5E6F2E2DA760AF51]
- Jordan Liggitt `<jliggitt@redhat.com>` [4096R/0x39928704103C7229]

**Known issues**

- We haven't specified a way to cycle the Product Security Team; but we need this process deployed quickly as our current process isn't working. I (@philips) will put a deadline of March 1st 2017 to sort that.

## Release Manager Role

Also included on the private [Kubernetes Security](https://groups.google.com/forum/#!forum/kubernetes-security) list are all [Release Managers](https://github.com/kubernetes/community/wiki).

It is the responsibility of the PST to add and remove Release Managers as Kubernetes minor releases created and deprecated.

## Disclosures

### Private Disclosure Processes

The Kubernetes Community asks that all suspected vulnerabilities be privately and responsibly disclosed via the Private Disclosure process available at [https://kubernetes.io/security](https://kubernetes.io/security).

### Public Disclosure Processes

If you know of a publicly disclosed security vulnerability please IMMEDIATELY email [kubernetes-security@googlegroups.com](mailto:kubernetes-security@googlegroups.com) to inform the Product Security Team (PST) about the vulnerability so they may start the patch, release, and communication process.

If possible the PST will ask the person making the public report if the issue can be handled via a private disclosure process. If the reporter denies the PST will move swiftly with the fix and release process. In extreme cases you can ask GitHub to delete the issue but this generally isn't necessary and is unlikely to make a public disclosure less damaging.

## Patch, Release, and Public Communication

For each vulnerability a member of the PST will volunteer to lead coordination with the Fix Team, Release Managers and is responsible for sending disclosure emails to the rest of the community. This lead will be referred to as the Fix Lead.

The role of Fix Lead should rotate round-robin across the PST.

All of the timelines below are suggestions and assume a Private Disclosure. The Fix Lead drives the schedule using their best judgment based on severity, development time, and release manager feedback. If the Fix Lead is dealing with a Public Disclosure all timelines become ASAP.

### Fix Team Organization

These steps should be completed within the first 24 hours of Disclosure.

- The Fix Lead will work quickly to identify relevant engineers from the affected projects and packages and CC those engineers into the disclosure thread. This selected developers are the Fix Team. A best guess is to invite all assignees in the OWNERS file from the affected packages.
- The Fix Lead will get the Fix Team access to private security repos to develop the fix.

### Fix Development Process

These steps should be completed within the 1-7 days of Disclosure.

- The Fix Lead and the Fix Team will create a [CVSS](https://www.first.org/cvss/specification-document) using the [CVSS Calculator](https://www.first.org/cvss/calculator/3.0). The Fix Lead makes the final call on the calculated CVSS; it is better to move quickly than make the CVSS perfect.
- The Fix Team will notify the Fix Lead that work on the fix branch is complete once there are LGTMs on all commits in the private repo from one or more relevant assignees in the relevant OWNERS file.

If the CVSS score is under 4.0 ([a low severity score](https://www.first.org/cvss/specification-document#i5)) the Fix Team can decide to slow the release process down in the face of holidays, developer bandwidth, etc. These decisions must be discussed on the kubernetes-security mailing list.

### Fix Disclosure Process

With the Fix Development underway the Fix Lead needs to come up with an overall communication plan for the wider community. This Disclosure process should begin after the Fix Team has developed a Fix or mitigation so that a realistic timeline can be communicated to users.

**Disclosure of Forthcoming Fix to Users** (Completed within 1-7 days of Disclosure)

- The Fix Lead will email [kubernetes-announce@googlegroups.com](https://groups.google.com/forum/#!forum/kubernetes-announce) and [kubernetes-security-announce@googlegroups.com](https://groups.google.com/forum/#!forum/kubernetes-security-announce) informing users that a security vulnerability has been disclosed and that a fix will be made available at YYYY-MM-DD HH:MM UTC in the future via this list. This time is the Release Date.
- The Fix Lead will include any mitigating steps users can take until a fix is available.

The communication to users should be actionable. They should know when to block time to apply patches, understand exact mitigation steps, etc.

**Optional Fix Disclosure to Private Distributors List** (Completed within 1-14 days of Disclosure):

- The Fix Lead will make a determination with the help of the Fix Team if an issue is critical enough to require early disclosure to distributors. Generally this Private Distributor Disclosure process should be reserved for remotely exploitable or privilege escalation issues. Otherwise, this process can be skipped.
- The Fix Lead will email the patches to kubernetes-distributors-announce@googlegroups.com so distributors can prepare builds to be available to users on the day of the issue's announcement. Distributors can ask to be added to this list by emailing kubernetes-security@googlegroups.com and it is up to the Product Security Team's discretion to manage the list.
  - TODO: Figure out process for getting folks onto this list.
- **What if a vendor breaks embargo?** The PST will assess the damage. The Fix Lead will make the call to release earlier or continue with the plan. When in doubt push forward and go public ASAP.

**Fix Release Day** (Completed within 1-21 days of Disclosure)

- The Release Managers will ensure all the binaries are built, publicly available, and functional before the Release Date.
  - TODO: this will require a private security build process.
- The Release Managers will create a new patch release branch from the latest patch release tag + the fix from the security branch. As a practical example if v1.5.3 is the latest patch release in kubernetes.git a new branch will be created called v1.5.4 which includes only patches required to fix the issue.
- The Fix Lead will cherry-pick the patches onto the master branch and all relevant release branches. The Fix Team will LGTM and merge.
- The Release Managers will merge these PRs as quickly as possible. Changes shouldn't be made to the commits even for a typo in the CHANGELOG as this will change the git sha of the already built and commits leading to confusion and potentially conflicts as the fix is cherry-picked around branches.
- The Fix Lead will request a CVE from [DWF](https://github.com/distributedweaknessfiling/DWF-Documentation) and include the CVSS and release details.
- The Fix Lead will email kubernetes-{dev,users,announce,security-announce}@googlegroups.com now that everything is public announcing the new releases, the CVE number, the location of the binaries, and the relevant merged PRs to get wide distribution and user action. As much as possible this email should be actionable and include links how to apply the fix to users environments; this can include links to external distributor documentation.
- The Fix Lead will remove the Fix Team from the private security repo.

### Retrospective

These steps should be completed 1-3 days after the Release Date. The retrospective process [should be blameless](https://landing.google.com/sre/book/chapters/postmortem-culture.html).

- The Fix Lead will send a retrospective of the process to kubernetes-dev@googlegroups.com including details on everyone involved, the timeline of the process, links to relevant PRs that introduced the issue, if relevant, and any critiques of the response and release process.
- The Release Managers and Fix Team are also encouraged to send their own feedback on the process to kubernetes-dev@googlegroups.com. Honest critique is the only way we are going to get good at this as a community.

<!-- BEGIN MUNGE: GENERATED_ANALYTICS -->
[![Analytics](https://kubernetes-site.appspot.com/UA-36037335-10/GitHub/docs/devel/security-release-process.md?pixel)]()
<!-- END MUNGE: GENERATED_ANALYTICS -->
