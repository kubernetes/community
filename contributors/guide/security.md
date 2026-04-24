# Kubernetes Security Processes

This document describes what the Kubernetes project fixes, on what timeline,
and what we decline to treat as a vulnerability. It is written for
reporters, users, distributors, contributors, and community members.

## Scope and assumptions

The properties we defend — authentication, authorization, admission, tenant
isolation, secret handling, component integrity — assume a cluster where
the operator has applied documented secure defaults. Operator-chosen
configuration is a trust boundary: nodes the operator has admitted to the
cluster are trusted; clusters that run with RBAC disabled or with
authorization bypassed are not defended; in-cluster administrators
exercising their granted permissions are not attackers.
For more details on how SRC assesses severity of a reported vulnerability: https://github.com/kubernetes/committee-security-response/blob/main/severity-ratings.md
Components outside the `kubernetes/kubernetes` repository — out-of-tree
CSI, CNI, and cloud-provider implementations, the Kubernetes SIG
repositories, ingress-nginx, image-builder, and similar — are owned by
their respective maintainers. The Kubernetes SRC coordinates
with those maintainers when they consider it necessary.

## Reporting

See:
https://kubernetes.io/docs/reference/issues-security/security/

Report privately through HackerOne (`https://hackerone.com/kubernetes`)
or by email to `security@kubernetes.io`. Include a reproducible proof of
concept exploit code/script, symbol-level reachability evidence, or a code trace from an
attacker-reachable entry point to the flaw. The SRC acknowledges reports
within three working days and begins triage within seven calendar days,
but does not commit to a deterministic investigation or fix deadline.
Complex vulnerabilities in infrastructure this large take weeks of
careful work, and we will not rush a fix we are not confident in.

## What we fix

A finding receives a Kubernetes CVE and follows the release process
below when it breaks one of the properties above — against a currently
supported release and a configuration the project documents as
supported. See: https://github.com/kubernetes/committee-security-response/blob/main/security-release-process.md#patch-release-and-public-communication

Supported releases are the most recent three minors (N, N-1, N-2), each
receiving roughly fourteen months of patch support. See: https://kubernetes.io/releases/patch-releases/#support-period 

We backport security fixes to all three. We do not patch end-of-life 
releases; commercial distributions often provide extended support 
for older versions. 
## What we do not treat as a vulnerability

The following are not Kubernetes vulnerabilities in themselves. They may
still be legitimate bugs, and we welcome those reports through the
normal issue tracker. They will not receive a CVE from the Kubernetes
CNA and will not drive out-of-band releases.

- **Misconfigurations:** Deviating from documented secure defaults —
  anonymous auth, missing RBAC, privileged containers for untrusted
  workloads — is an operator choice, not a product flaw.
- **Out-of-support versions:** We do not investigate reports against
  end-of-life minor releases.
- **Transitive dependency CVEs with no reachable symbol.** Reachability
  is verified by symbol-level analysis (for example, `govulncheck`) or
  a code trace. Version-matching scanner output is not by itself
  evidence of exploitability.
- **Actions available to a caller who already holds the equivalent
  permission.** A cluster admin doing variants of admin is not a
  privilege escalation.
- **Alpha-gated features off by default.** Experimental on purpose; we
  fix the bugs but do not cut emergency releases for them.
- **Findings reachable only on platforms we do not ship** (for example,
  CVEs in a dependency guarded behind a BSD build tag).
- **Denial of service through pathological input from a caller who
  already has write access to the object being parsed.**
- **Crashes where graceful restart is part of the documented
  operational model.**
- **Weak algorithms, missing headers, or deprecated settings an
  operator opted into by configuration.**
- **Scanner output presented as the primary evidence of a flaw.**
  Scanners are useful; they are not authoritative.
- **AI-generated reports whose code paths, symbols, or traces cannot
  be reproduced.**
- **Compliance-framework findings that do not describe a concrete
  flaw.** That conversation belongs between the compliance vendor, distributor and operator, not with the SRC.

## Severity and release path

The SRC assigns one of four severity tiers using project-specific
judgment. CVSS vectors accompany every advisory because downstream
tools expect them, but CVSS is not authoritative and we do not endorse
third-party scores (including NVD's). What the tier actually controls
is the release path. See: https://github.com/kubernetes/committee-security-response/blob/main/security-release-process.md#fix-development-process and 

Patch releases ship monthly; a vulnerability introduced and fixed on
`master` before any release tag carries it does not receive a CVE. See: https://kubernetes.io/releases/patch-releases/

The Kubernetes project is a CVE Numbering Authority for components in
the `kubernetes/kubernetes` repository. CVEs in upstream dependencies
(Go, containerd, etcd, runc, kernel) are the responsibility of their
respective CNAs. CVEs in out-of-tree Kubernetes SIG projects are
issued by those projects' CNAs when they have one, or through MITRE
otherwise. The SRC coordinates across these boundaries but does not
re-issue CVE identifiers for code it does not own.

## CVEs in our dependencies

Kubernetes is a large Go program that vendors hundreds of modules.
Kubernetes releases also include container images that have their own
dependencies. A substantial fraction of all CVE-related traffic we
receive is about code we import rather than code we wrote. This
section describes how we decide what to do with such reports and what
that implies for our release cycle.

We classify every upstream CVE into one of three buckets, and each
bucket has a distinct release path:

- **Reachable and exploitable in Kubernetes.** The vulnerable symbol
  is called from a Kubernetes-shipped binary in a way that preserves
  the upstream attack vector. We treat this as a Kubernetes
  vulnerability: the SRC assigns a severity, coordinates the
  dependency bump with SIG Release, and ships it through the path
  above. Critical and High cases may trigger an out-of-band release;
  Medium and Low cases ride the next scheduled monthly patch.
- **Reachable but not exploitable.** The symbol is called but the
  upstream attack vector cannot be reproduced in our usage (for
  example, the callers pass only trusted input, or a wrapping API
  imposes limits the upstream library does not). We bump the
  dependency on `master` on a normal schedule and backport only when
  the upstream fix is low-risk and contained. No CVE is issued by the
  Kubernetes CNA for this case. Publishing machine-readable
  "Not Affected" statements for this class is planned work (see
  "Machine-readable artifacts" below).
- **Not reachable.** The vulnerable symbol is not called from any
  Kubernetes-shipped binary, or it is guarded behind a build tag we
  do not ship (BSD, exotic architectures, debug-only paths). A
  `master` bump happens when the dependency graph organically pulls
  the fixed version; we do not open cherry-picks for release branches
  merely to change a version string.

Historically, first scenario rarely happens whereas Scenario 2 and 3 are much more prevalent.
The monthly patch cadence is a feature of this classification, not a
bug. A Medium-severity, reachable-and-exploitable upstream CVE ships
within roughly thirty days by default — faster than most compliance
frameworks require, slower than the scanner ticket demands. If the
K8s SRC Fix Lead believes users are exposed sooner than that, the case is
reclassified up (to High, or to an out-of-band release); if the
Any amount of external escalation plays no part in the decision making of when a fix will land.

Go toolchain updates are a special case. Each Kubernetes minor pins a
Go minor; we cannot swap Go versions mid-release-cycle without
substantial risk. When a Go standard library CVE lands, SIG Release
evaluates whether a patch-level Go bump covers it (the common case,
handled in the next scheduled patch), whether a minor Go bump is
required (rare; coordinated with the Go team and typically deferred
to the next Kubernetes minor unless exploitability is high), or
whether a targeted dependency bypass makes the Go issue irrelevant
to Kubernetes (case by case). The Kubernetes CNA does not issue separate CVE numbers 
for Go standard library bugs and uses Go Community's already issued CVEs to track this work.

Dependency updates whose only purpose is to silence a scanner, without additional evidence such as but not limited to reachability analysis, are not eligible for cherry-pick to release branches. This rule is
load-bearing. Without it, every noisy scanner becomes a forcing
function on the release calendar, and the release calendar stops
serving users who actually need predictable patching for relevant issues.

End-of-life branches receive no dependency bumps, whether for CVE
reasons or otherwise. Extended-support scenarios (older clusters on
managed services, appliance deployments, regulated environments) are
a concern for commercial vendors, not for the upstream community.

## Project release cadences

Kubernetes is not a single release train. The monthly patch cadence
described above applies to the core components built out of
`kubernetes/kubernetes` — kube-apiserver, kubelet, controller-manager,
scheduler, kube-proxy. Dozens of other projects in the `kubernetes` and
`kubernetes-sigs` GitHub organizations — ingress-nginx, the CSI and CNI
drivers, cluster-api and its providers, kustomize, cri-tools,
external-dns, image-builder, and many more — run their own release
cadences, set by their own maintainer teams and not aligned with the
core. Most of these projects are maintained by very small groups, often
on volunteer time.

Please do not ask individual project maintainers to cut out-of-band
releases because a scanner flagged something in their repository.
Maintainer time spent on compliance-driven releases is time not spent on
actual bug fixes or feature work, and repeated pressure on a small team
is corrosive. The good-citizen pattern is to file the issue (with
reachability evidence, per this document), read the project's release
cadence, and wait for the next scheduled release unless the finding
genuinely warrants acceleration. AI driven PRs without human reviews or involvement, that fix scanner reported vulnerabilities, do not help and in fact waste precious maintainer time. 

If a finding meets the severity bar for an out-of-band release, the SRC
and the project maintainers will coordinate without external prompting;
you do not need to ping every maintainer individually to get that to
happen. If it does not meet that bar, the fix ships in the project's
next scheduled release. Either way, the release calendar is the
project's call, not the scanner's.

## Machine-readable artifacts

**Currently published.** The SRC and SIG Security publishes an official CVE feed in
JSON and RSS on kubernetes.io, derived from GitHub issues labeled
`official-cve-feed`. Release artifacts carry signed SBOMs and SLSA
provenance. An OSV-format mirror lives at
`kubernetes-sigs/cve-feed-osv` as a secondary artifact for tools that
consume OSV natively.

**Planned.** Per-release VEX ("Not Affected") statements covering the
reachability status of common scanner-flagged CVEs — tracked at
`kubernetes/kubernetes#121454` — are the intended mechanical answer to
transitive-dependency false positives, and will become how we
neutralize scanner noise at scale once emission is in place.
Publication of the primary kubernetes.io CVE feed in OSV schema is
also planned, to close the gap with osv.dev ingestion. Annual
aggregate triage statistics (reports received, reports declined, CVEs
issued, qualitative effort summary) are a proposed addition to the
SIG Security annual report, so the cost of scanner noise is visible
rather than implicit.

If your scanner does not consume VEX when it arrives, please ask your
vendor to fix that before filing additional reports against us.

## Maintainer sustainability

The Security Response Committee is a small, part-time group running
on a business-hours oncall rotation. We prioritize issues with
demonstrated real-world impact on Kubernetes users over issues
surfaced only by vulnerability scanners or compliance tooling. We do
not offer SLA-backed response, and if your compliance program
requires one — a FedRAMP-style thirty-day remediation window, a PCI
audit deadline, a customer contract — please engage a commercial
Kubernetes distribution vendor; that is what they exist for.

## Guidance for contributors fielding security-flavored issues

When a public issue arrives that should have been private, ask the
reporter to move it to HackerOne or email and close the issue; do not
amplify it with technical debate in public. When a scanner-driven
issue arrives, post the reachability evidence (or ask for it), link
this document, and close — the threshold for re-opening is a specific
call path, not a louder complaint. When a compliance-driven issue
arrives, be courteous, point at the release cadence and this
document, and close. When a dependency bump request cites a scanner
but no reachability analysis, decline the backport and link this
document. The SRC comms templates include canned responses for each
of these patterns; prefer them over ad-hoc prose. None of this is
rudeness; it is the only way a volunteer committee can stay effective
in a project this size.