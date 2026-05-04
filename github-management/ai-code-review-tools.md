# AI Code Review Tools

The Kubernetes project may evaluate AI-powered code review tools on a
per-repo opt-in basis. This document describes the process for requesting,
evaluating, and deciding on the use of such tools.

## Scope

This policy covers AI tools that automatically review pull requests, such as
CodeRabbit or GitHub Copilot code review. It does not cover other AI-powered
tooling such as CI/CD, security scanning, or code generation assistants.

## Requesting a New Tool

A subproject lead or an approver listed in the repository's top-level [OWNERS]
file files an issue on [kubernetes/org]. The issue must:

- Identify the tool and link to its documentation
- Describe the use cases and what the subproject is trying to accomplish
- Explain why existing approved tools do not meet their needs
- List the specific repositories for the pilot
- Acknowledge the [AI guidance for pull requests]
- List the GitHub permissions and OAuth scopes the tool requests

## Privacy and Security Assessment

Upon receiving a request, the [GitHub Administration Team] conducts a privacy
and security assessment of the tool. The assessment documents:

- What GitHub permissions and OAuth scopes the tool requires
- What data the tool accesses and where it is sent
- What AI models are used to process the code
- Data retention and deletion policies
- Security certifications (SOC2, etc.)
- Whether access can be scoped to specific repositories

The assessment is documented in the [kubernetes/org] issue for transparency.

## Approval

The [GitHub Administration Team] reviews the request and the privacy and
security assessment, and approves or rejects the request. If approved, the
[GitHub Administration Team] enables the tool on the requested repositories
and applies an org-wide default configuration.

## Pilot Structure

- The pilot runs for a 90-day evaluation period starting from the date the
  tool is enabled
- The tool is enabled only on the specific requested repositories, not org-wide
- An org-wide default configuration is applied; repositories may customize
  within those bounds
- The sponsoring subproject is responsible for collecting feedback from
  contributors and reviewers during the pilot

## Evaluation and Decision

At the end of the pilot period, the sponsoring subproject posts an evaluation
summary to the original [kubernetes/org] tracking issue. The summary should
cover:

- Quality of reviews (signal vs noise, with examples)
- Contributor and reviewer feedback
- Any issues encountered
- Recommendation (continue, modify, or remove)

The [GitHub Administration Team], in consultation with the sponsoring
subproject, decides to:

- Continue and expand to additional repositories
- Continue with modifications
- Remove the tool

Expansion to additional repositories requires a new issue on [kubernetes/org].
If the tool has already been approved and assessed, the privacy and security
assessment does not need to be repeated. Requests to enable an already-approved
tool on additional repositories are evaluated on cost impact and repository
fit. Approvals for tools with per-repository costs may require additional
review.

## Approved Tools

The following tools have completed the evaluation process and are approved for
use in the Kubernetes project. Subprojects may request enablement on their
repositories without repeating the privacy and security assessment.

No tools have completed the evaluation process yet. This section will be
updated as tools are approved.

## Removal

If a pilot is unsuccessful or a tool is no longer desired, the
[GitHub Administration Team] will disable the integration. Subproject leads may
request removal at any time by filing an issue on [kubernetes/org].

[GitHub Administration Team]: /github-management/README.md#github-administration-team
[AI guidance for pull requests]: /contributors/guide/pull-requests.md#ai-guidance
[OWNERS]: /contributors/guide/owners.md
[kubernetes/org]: https://github.com/kubernetes/org/issues
