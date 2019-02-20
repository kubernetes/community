# Adding a new member to a Kubernetes GitHub Org

This procedure outlines the steps to add someone to a Kubernetes GitHub
organization. These directions are based off of the [community membership]
guidelines. If there is a discrepancy as to the exact membership requirements,
the [community membership] guidelines take precedence.

## Membership Requirements

All members of the Kubernetes GitHub org are required to do the following:

- Have 2FA enabled on their GitHub account

  This is enforced by GitHub. The contributor will not be able to accept the org
  invitation without it, and they will be removed from the org automatically in
  the event they disable it.

- Join the [k-dev] mailing list

  We don't explicitly check this, but we require the contributor to check the
  box saying that they have joined.

- List their active contributions

  This is left purposefully vague. We want to record what those contributions
  are, but we leave it up to the sponsors to determine if they are at a bar that
  warrants membership. We want to make it easy for those who are actively
  contributing to join, but this bar should be higher than, for example, a few
  spelling corrections.

- List the active subprojects they are involved in

  In the same vein as the above, we leave this vague, and as guidance for the
  sponsors. The sponsors should also be active in these same subprojects.

## Sponsor Requirements

One of the most critical pieces in the process is validating the sponsors meet
the requirements to sponsor a new member. With the size of the Kubernetes
project, those involved with processing new memberships do not have the
visibility to determine if each prospective member's contributions meet the
standard needed to become a member. We rely on sponsors to vouch for the work of
the prospective new member, and to validate the work they are doing.

We do however, have to validate that the sponsor's credentials meet the standard
required to be eligible to sponsor a new member. These requirements are:

- Sponsors must be a member of the org they are attempting to sponsor for.

  For example, if the membership being requested is for kubernetes-incubator,
  then the sponsor should be a member of either that org, or main Kubernetes
  org (as members of the main org have implicit membership in other orgs).

- Sponsors must be a reviewer or approver in at least one OWNERS file in
  either the [Kubernetes GitHub org] or the org they are sponsoring for.

- Sponsors must be from multiple member companies to demonstrate integration
  across community

  A single sponsor may be from the same company as the prospective member, but
  no more than one sponsor can be from the same company. If both sponsors on an
  application are from the same company, please request that the applicant
  solicit an additional sponsor from a different company.

- Sponsors must have close interactions with the prospective member

  Sponsors need to be familiar enough with the prospective member's
  contributions to be able to properly vouch for them. This is to ensure
  integrity in the membership process, and a "web of trust" for the privileges
  we afford members.

If there are questions with regard to a sponsor's qualifications, please attempt
to seek clarification or a second opinion before moving forward with processing
the membership.

## Processing the request

Once all the requirements have been validated, we can move forward with
processing the membership.

1. Open a PR against the [kubernetes/org] config, adding the potential member's
GitHub username to the members list. Note, that the list is alpha sorted, but
case insensitive.

1. For clarity, please use one commit per username added (although you may add
that user to multiple orgs in the same commit).

1. In the PR body (and not the commit message), add `fixes #1234, fixes #1234`
with the issue numbers of the membership request issues that this PR is
resolving. One PR can be used to resolve multiple membership requests.

1. Add a note to the original membership request stating that a membership
invite will be sent out once the PR has merged.

1. Wait for a member of the GitHub administration team to approve the PR for
merge.




[community membership]: /community-membership.md
[k-dev]: https://groups.google.com/forum/#!forum/kubernetes-dev
[kubernetes/org]: https://git.k8s.io/org/
[Kubernetes GitHub org]: https://github.com/kubernetes
