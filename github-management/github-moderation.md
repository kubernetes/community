# GitHub Moderation

Contributions to Kubernetes GitHub repositories must meet legal, technical, and behavioral standards. Submissions that fail to meet these expectations, whether due to copyright violations, low-quality automated generation, spam-like behavior, or bad-faith engagement may be moderated, restricted, or escalated according to established project processes.

GitHub moderation is handled at the repository level by maintainers, with escalation to **[@kubernetes/owners](https://github.com/kubernetes/org/blob/main/OWNERS)** and the [Kubernetes Code of Conduct Committee](./committee-code-of-conduct) when necessary.

All actions must follow the [Kubernetes Code of Conduct](./committee-code-of-conduct) and the escalation procedures defined in the [moderation documentation](/communication/moderation.md).

## Comment Moderation

Maintainers may edit or remove comments that:

- Violate the [CNCF Code of Conduct](https://github.com/cncf/foundation/blob/main/code-of-conduct.md).
- Contain harassment or personal attacks
- Include sensitive personal information
- Are spam or malicious

## Legal and Copyright Compliance

All contributions must:

- Be covered by a signed [Contributor License Agreement (CLA)](https://git.k8s.io/community/CLA.md).
- Comply with copyright requirements.
- Not introduce code copied from third-party sources in violation of license terms.

This applies equally to:

- Human-authored content
- Copy-pasted content
- AI-generated content

Use of generative tools does not exempt contributors from CLA or copyright obligations. Contributors are responsible for ensuring they have the right to submit all material.

See [Linux Foundation guidance](https://www.linuxfoundation.org/legal/generative-ai) on generative AI and licensing expectations.

## Quality and Engineering Standards

Each pull request requires reviewer time and project resources.
Submissions must:

- Demonstrate understanding of the change being proposed.
- Be technically sound.
- Meet project quality standards.
- Be responsive to review feedback.

Contributions that appear automated, mechanically generated, or submitted without understanding may be closed if:

- The submitter cannot meaningfully respond to review.
- The change introduces avoidable defects.
- The review burden is disproportionate to the value of the change.

This standard applies regardless of whether content was human-written or AI-generated. Contributors must understand and be able to explain and modify the changes they submit, regardless of how they were authored.

## Spam, Automation, and Bad-Faith Contributions

Kubernetes has historically received automated and low-quality submissions. Repeated spammy or obviously automated behavior may result in:

- PR closure
- Issue locking
- Participation restriction
- Escalation to organization-level moderation
- Account banning where appropriate

Excessive low-quality automated submissions are treated as disruptive behavior.

Moderation decisions will consider:

- Volume of submissions
- Responsiveness to feedback
- Pattern of behavior across repositories

## Trivial or Fragmented Edits

Contributors are encouraged to batch related fixes. Trivial single-line edits that create unnecessary review overhead may be closed with a request to consolidate improvements into a single submission.

Automation reduces cost, but review remains an engineering effort. Contributors should respect reviewer time.

See [guidance on trivial edits](https://www.kubernetes.dev/docs/guide/pull-requests/#trivial-edits) in the pull request guidelines.