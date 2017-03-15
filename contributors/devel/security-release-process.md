# Security Release Process

Kubernetes is a large growing community of volunteers, users, and vendors. The Kubernetes community has adopted this security disclosures and response policy to ensure we responsibly handle critical issues.

## Product Security Team (PST)

Security vulnerabilities should be handled quickly and sometimes privately. The primary goal of this process is to reduce the total time users are vulnerable to publicly known exploits.

The Product Security Team (PST) is responsible for organizing the entire response including internal communication and external disclosure but will need help from relevant developers and release managers to successfully run this process.

The initial Product Security Team will consist of four volunteers subscribed to the private [Kubernetes Security](https://groups.google.com/forum/#!forum/kubernetes-security) list. These are the people who have been involved in the initial discussion and volunteered:

- Brandon Philips `<brandon.philips@coreos.com>` [4096R/154343260542DF34]
- Jess Frazelle `<jessfraz@google.com>` [4096R/0x18F3685C0022BFF3]
- CJ Cullen `<cjcullen@google.com>`
- Tim St. Clair `<stclair@google.com>` [4096R/0x5E6F2E2DA760AF51]
- Jordan Liggitt `<jliggitt@redhat.com>`

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

- The Fix Lead and the Fix Team will create a [CVSS](https://www.first.org/cvss/specification-document) using the [CVSS Calculator](https://www.first.org/cvss/calculator/3.0). The Fix Lead makes the final call on the calculated CVSS; it is better to move quickly than make the CVSS prefect.
- To encourage fast adoption of the security fix, the new patched release should
  contain only security fixes.
- To facilitate this goal, the Fix Lead will [create a patch release branch](https://github.com/kubernetes/release/blob/master/docs/branching.md#branching-from-a-tag)
  in a private fork of the public kubernetes.git repository, branched from the
  most recent released tag for the release being fixed.
  - As a hypothetical example, if `v2.4.7` is the latest public patch release for
   `v2.4`, a new branch will be created in the private fork called
   `release-2.4.7`, branching at the `v2.4.7` tag. If `HEAD` were to be used
   instead, the security release might include non-security fixes.
   - Alternatively, this branch could be named something like `issue-N`,
     referencing the GitHub issue tracking the vulnerability and fix.
  - TODO: continuous integration testing should be set up for the branch,
    ideally with only a very small configuration change.
- If a security incident affects multiple supported releases, then patch
  release branches will be created for all affected releases. Vulnerabilities
  affecting `master` will be fixed via cherry-picks after fixes to the release
  branches have been publicized.
- If multiple security issues are reported simultaneously, they will be combined
  into a single security release target, unless it is deemed infeasible to fix
  all vulnerabilities at once (e.g. if one vulnerability will take significantly
  longer to fix).
- The Fix Team will merge fixes directly to the private branch(es), rather than
  following the normal cherry-pick process
  - As is feasible, the normal PR review process (including LGTM and OWNERS
    approval) should be followed for all fixes, though the submit queue
    infrastructure may not be operational on the private branches.
- The Fix Team will notify the Fix Lead that work on the fix branch(es) is
  complete once all relevant changes have been merged into the private
  branch(es).
- Release Managers are encouraged to refrain from creating new public patch
  releases while security fixes are ongoing. Should it be necessary to create a
  new public patch release before the security fixes are complete, the private
  security branch will need to be rebased (and renamed?) on the new latest
  patch release. Only one security release will be built.

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

- The Release Managers will ensure all the binaries are built,
  publicly available, and functional before the Release Date.
    - Note: we ship source with our binary releases, so as soon as binary
      artifacts are publicly available, the fixes (and vulnerabilities) are
      effectively public.
  - CI testing of the private patch release branch should provide confidence in
    the release.
- The Release Manager will use release tooling (i.e. `anago`) to
  - Compute the version for the new security patch release; in the previous
    example, the `release-2.4.7` branch would produce a `v2.4.8` release.
  - Verify that no public patch release has occurred since the security branch
    was cut (in this example, make sure that `v2.4.8` has not been tagged
    publicly).
  - Perform a normal release build from the private release branch, staging
    binary artifacts locally (or possibly in a private bucket).
  - Push the new patch release tag to the private repo.
  - The tooling should **not** upload any artifacts to public storage buckets or
    create a draft release on the public GitHub repo.
- As soon as we are ready to make the fix known publicly (TODO: when?),
  the Release Manager will use tooling to
  - Publish results.
    - Upload binary artifacts to GCS, gcr.io, and GitHub. (TODO: are these
      prebuilt from the previous step?)
    - Update the `CHANGELOG.md` in the **public** `master` branch
    - Create a new release on the **public** kubernetes GitHub
    - Send email announcing the release
  - Merge the private patch release branch into the **public** release branch
    including both the commits and release tag.
    - In our continuing example, this would merge the private `release-2.4.7`
      branch into the public `release-2.4` branch, effectively advancing the
      latter to `2.4.9-beta.0+`. The `git describe --tags` will update to
      '2.4.9-beta.0+` and any changes in the `release-2.4` branch will
      now ship with `v2.4.9`.
- The Fix Lead will request a CVE from [DWF](https://github.com/distributedweaknessfiling/DWF-Documentation) and include the CVSS and release details.
- The Fix Lead will email kubernetes-{dev,users,announce,security-announce}@googlegroups.com now that everything is public announcing the new releases, the CVE number, the location of the binaries, and the relevant merged PRs to get wide distribution and user action. As much as possible this email should be actionable and include links how to apply the fix to users environments; this can include links to external distributor documentation.
- The Fix Lead will remove the Fix Team from the private security repo.
- The Fix Lead will delete the patch release branch from the private security
  repo.
- If fixes are needed for `master`, the Fix Lead (or chosen delegate) will open
  PRs against the public `master` branch containing cherry-picks of the commits
  now merged into the release branch(es). No release will be built or published
  for `master`.

### Retrospective

These steps should be completed 1-3 days after the Release Date. The retrospective process [should be blameless](https://landing.google.com/sre/book/chapters/postmortem-culture.html).

- The Fix Lead will send a retrospective of the process to kubernetes-dev@googlegroups.com including details on everyone involved, the timeline of the process, links to relevant PRs that introduced the issue, if relevant, and any critiques of the response and release process.
- The Release Managers and Fix Team are also encouraged to send their own feedback on the process to kubernetes-dev@googlegroups.com. Honest critique is the only way we are going to get good at this as a community.

<!-- BEGIN MUNGE: GENERATED_ANALYTICS -->
[![Analytics](https://kubernetes-site.appspot.com/UA-36037335-10/GitHub/docs/devel/security-release-process.md?pixel)]()
<!-- END MUNGE: GENERATED_ANALYTICS -->
