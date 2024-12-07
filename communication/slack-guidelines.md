---
title: "Slack Guidelines"
description: |
  These are the policies, procedures, and escalation mechanisms for the Slack platform,
  such as requesting channels, tokens, and how to report inappropriate content.
---

Slack serves as the main communication platform for the Kubernetes community
outside of the mailing lists. It's important that conversations stay on topic in
each channel, and that everyone abides by the [Code of Conduct][coc]. There are
over 150,000 members who should all expect to have a positive experience. You
are a part of building and keeping a positive community.

Chat is searchable and public. It's best not to make comments that you would not
say on a video recording or in another public space. Please be courteous to
others.

## Code of Conduct

Kubernetes adheres to the [Kubernetes Code of Conduct][coc] throughout the
project, and includes all communication mediums.

## Admins

-   Check the [centralized list of administrators][admins] for contact information.

Slack Admins will have listed that they are a Slack admin in their Slack
profile, along with their specific timezone.

To connect: please reach out in the `#slack-admins` channel, mention an admin
directly in the `#slack-admins` channel when you have a question, or DM (direct
message) one privately.

## General Communication Guidelines

### Workspace Channel History

The Kubernetes Slack Workspace is archived and made available when the
administrators have time. There is no explicit interval.

[Slack Archive Download]

### Slack Usage Rules

In addition to the [Code of Conduct][coc], users of Kubernetes Slack must adhere to
the following behavior rules.  Failure to follow them may result in warnings,
removal of posts, escalating to temporary or permanent deactivation of the user
account.

* **Adhere to Channel Topics**: for the most part, try to stick to the declared
  topic of the channel you're in.
* **No Spam**: you may not post commercial, promotional messages anywhere in
  Kubernetes, whether in public channels or direct messages, except for very
  specific types of messages in specific channels were it is allowed.
* **No Unsolicited DMs**: do not send any user uninvited direct messages (DMs).
* **Don't Discuss Company Business**: Proprietary conversations should
  occur in your company Slack and/or communication platforms.
* **No Surveys**: surveys, focus groups, free trials, and similar calls for
  participation are considered spam and are prohibited.  Exceptions: an official
  Kubernetes community survey can be posted anywhere, and any survey may be
  posted to the Surveys channel.
* **Jobs and Resumes** belong on the `#kubernetes-careers` channel and should not
  be posted elsewhere.
* **Events** may only be posted to channels where that specific event is
  applicable, except for the `#events` channel.  Generally, this only includes
  nonprofit community events.

### Specific Channel Rules

Some channels have specific rules or guidelines. If they do, they will be listed
in the purpose or pinned docs of that channel.  A few channels worth mentioning
include:

-   `#events` - announcements around any Kubernetes ecosystem related events.
-   `#kubernetes-new-contributors` - Questions and discourse around getting
    started contributing to Kubernetes.
-   `kubernetes-org-members` - Discussion by and for organization members and
    established contributors.
-   `#kubernetes-careers` - Job openings for positions working with/on/around
    Kubernetes. These must be postings for specific jobs, not "cattle calls"
    for general tech hiring. You may also post resumes if you are looking
    to be hired. Job postings must include:
    -   A link to the posting or job description.
    -   The business name that will employ the Kubernetes hire.
    -   The location of the role or if remote is OK.
-   `#surveys` - Cloud native community wide surveys. Posts must be ecosystem
    related.

### Escalating and/or Reporting a Problem

The Kubernetes Slack has an integrated tool for reporting issues. It may be
accessed by clicking on "More actions", the "**...**" to the right of a message,
and selecting **Report message**.

This will open a new dialog prompt where you may describe the problem. When
done, it will send the reported message and your comments to **BOTH** the
[Slack admins][admins] and [Code of Conduct Committee (CoCC)][cocc].

A Slack admin or CoCC member will work to resolve the issue or reach out for
more information.

If the issue has not been responded to in a timely manner, Join the
#slack-admins channel and alert the admins to the current issue. Many Slack
admins watch the channel and should respond to you shortly.

As a last resort, or if the issue is private, contact one of the admins in the
[closest timezone][admins] via DM directly and describe the situation. If the
issue can be documented, please take a screenshot to include in your message.

#### What if you have a problem with an admin?

Send a DM to another [listed admin][admins] and describe the situation. If it's
a [code of conduct][coc] issue, please send an email to <conduct@kubernetes.io>
and describe the situation.

## Should you have a channel on the Kubernetes Slack?

The primary purpose of the Kubernetes slack is for the coordination of the
Kubernetes project. However it is useful for developers and users to have a
strong ecosystem of channels for related things. Here are some guidelines for
determining if you should request a channel:

-   The channel MUST be Kubernetes related in some way.
    -   Related cloud native projects might be more appropriate on the
        [CNCF Slack].
-   If requesting a project channel, the project MUST be open source.
    -   Open Source a project BEFORE requesting a channel. We cannot accommodate
        every organization's open sourcing launch plans.
    -   The purpose of Slack is to organize an existing community, not seed new
        ones.
    -   Moderators look at contributor activity, adoption and community
        consensus around a project, otherwise Slack becomes a vehicle for
        promotion instead of promulgation.
    -   Establishment and maintenance of a Slack channel should be an inflection
        point in a project's adoption.
    -   Requesting a channel means maintaining it on behalf of the project
        filing the issue. You will be expected to participate and foster a
        healthy discourse.
    -   External projects (ones not owned by a Kubernetes SIG) may have a
        maximum of two channels, usually `#project` or `#project-users`,
        and `#project-dev`.
    -   A second channel for a specific project will not be approved until the
        first channel demonstrates significant traffic.
-   Channels around commercial services built on OSS projects are allowed.
    -   Users love the value of being able to collaborate around various
        services.
    -   Keep it classy, on the Kubernetes Slack we're all on the same team.
-   The channel MUST be public.
    -   In order to ensure all channels adhere to our Code of Conduct, we
        heavily restrict private channels.
    -   If you need private discussion areas for security-sensitive topics, a
        project-specific Slack or the [CNCF Slack] may be a better fit.
-   Ask in `#slack-admins` or file an issue if you're unsure. We're happy to
    discuss it with you.

### Private Channels

We offer private channels on the Kubernetes Slack extremely sparingly.  The only
private channels we currently support involve community-owned software security
or community moderation in some way.  Unless supported by a Kubernetes SIG
or the Steering Committee, requests for a private channel will be treated with
skepticism.

Most private channels will be required to include a member of the Slack Admins
team.

### Requesting a Channel

Channels and User Groups are managed by [Tempelis], a tool that enables
external management of Slack.

To add a channel, open a Pull Request (PR) updating the [slack-config].

-   Add the channel to 'channels.yaml' following the [Channel Documentation]
    -   Channel names must be 21 characters or less in length, limited by Slack
        design.
    -   Channels must not share the same name with a Slack user or user group.
    -   Typical channel naming conventions follow:
        -   `#kubernetes-foo`
        -   `#sig-foo`
        -   `#meetup-foo`
        -   `#location-user`
        -   `#projectname`
-   In the PR comments, include some details regarding the purpose of the
    Channel.
    -   Channels should be dedicated to [SIGs or WGs][sig-list], sub-projects,
        community topics, or related Kubernetes programs/projects.
    -   Linking to resources such as the PR adding the subproject will speed in
        the validation and processing of the channel creation request.
    -   Channels are NOT
        -   **Company specific**; cloud providers are ok with product names as
            the channel. Discourse will be about Kubernetes-related topics and
            not proprietary information of the provider.
        -   **Private channels** with the exception of: code of conduct matters,
            mentoring, security/vulnerabilities, github management, or
            steering committee.
    -   Special accommodations will be made where necessary.

After you submit your request the Slack Admins will review and follow-up with
any questions in the PR itself.
There are two approvals needed. `/lgtm` and `/approve`. Once one moderator give the `/lgtm`, a hold will be placed on the PR using `/hold`. This hold will remain in place until one or more moderators reviews
and add the `/approve` command as well as `/hold cancel`, which will remove the hold on the PR.
Once it is signed off and merged, the channel will be created.

For further information, see the [Slack Config Documentation].

### Delegating Channel Ownership

Channel management can be delegated to other groups, enabling SIG leads or other
members to govern certain sets of channels. This bypasses the need for a Slack
Admins to sign-off on all requests and passes the responsibility to the most
relevant group.

To delegate channel ownership - Open a Pull Request (PR) updating the
[slack-config].

-   Create a sub-directory under the [slack-config] for your sig or group.
-   Update restrictions.yaml with an entry targeting yaml config files in the
    sub-directory you created along with one or more regular expressions that
    match the channel names that should be delegated.
    -   **Example Restrictions Entry:**
        ```yaml
        restrictions:
        - path: "sig-foo/*.yaml"          # path to channel config
          channels:
          - "^kubernetes-foo-[a-z]{1,3}$" # channel regexp - example match: kubernetes-foo-bar
          - "^foo-[a-zA-Z]+$"             # channel regexp - example match: foo-awesomechannel
        ```
-   Create an [OWNERS] file in the sub-directory adding the appropriate
    reviewers and approvers for the desired channels.
-   In the directory create one or more channel configs following the Channel
    Documentation
    -   **Example Channel Config:**
        ```yaml
        channels:
        - name: kubernetes-foo-bar # regexp: "^kubernetes-foo-[a-z]{1,3}$"
        - name: foo-users          # regexp: "^foo-[a-zA-Z]+$"
        - name: foo-dev            # regexp: "^foo-[a-zA-Z]+$"
        ```

After you submit your PR and the Slack Admins sign off on the update, it will be
merged and the group will be able to fully self-manage their own channels.

For further information, see the
[Slack Config Documentation][Channel Documentation].

## Requesting a User Group

Channels and User Groups are managed by [Tempelis], a tool that enables
external management of Slack.

To add a User Group, open a Pull Request (PR) updating the [slack-config].

-   Add the users to users.yaml. **NOTE:** This must be a mapping of their
    GitHub ID to their Slack Member ID.
-   To get a person's Slack Member ID, view their profile. Then click on the
    "**...**" and select **Copy member ID**. It will be a 9 character string of
    uppercase letters and numbers (example: `U1H63D8SZ`).
-   Update [usergroups.yaml] Follow the guidelines for creating a User Group
    in the Slack Config [User Group Documentation].
-   In the PR comments, include details on the User Group and `/cc` the members
    you are adding so that they may sign off and accept being added to the
    group.

After you submit your request, the Slack Admins will review and follow-up with
any questions in the PR itself. Once it is signed off by the members being added
and the Slack Admins, it will be merged, and the User Group will be created.

For further information, see the
[Slack Config Documentation][Channel Documentation].

## Requesting a Bot, Token, or Webhook

**READ BEFORE SUBMITTING A REQUEST**

Bots, tokens, and webhooks are reviewed on a case-by-case basis **with most
requests being rejected due to security, privacy, and usability concerns**. Bots
and the like tend to make a lot of noise in channels. The Kubernetes Slack
instance has over 145,000 members and it is the role of the Slack admins to
ensure everyone has a great experience.

Typically approved requests include: GitHub, CNCF requests, or other
tools/platforms used to aid in the management of Slack itself.

-   Create a [GitHub Issue] using the Slack Request template.
-   In the description, describe the request, its intended purpose and benefit
    to the community. Supplying links to supporting content such as a document
    outlining what OAuth scopes that are requested and why are **STRONGLY
    ENCOURAGED**.

After you submit your request, the Slack admins will review and follow-up with
any questions in the issue. If consensus can be reached among the admins, the
request will be approved and follow-up communication on implementation will be
discussed in Slack itself.

## Moderation

### Admin Expectations and Guidelines

Admins should adhere to the general Kubernetes project
[moderation guidelines].

Additionally, admins should ensure they have 2-factor auth enabled for their
account and mention they are a Slack admin in the "What I do" portion of their
profile. This message should also include the time zone they are representing.

Be mindful of how you handle communication during stressful interactions.
Administrators act as direct representatives of the community, and need to
maintain a very high level of professionalism at all times. If you feel too
involved in the situation to maintain impartiality or professionalism, that's a
great time to enlist the help of another admin.

Try to take any situations that involve upset or angry members to DM or video
chat. Please document these interactions for other Slack admins to review.

Content will be removed if it violates code of conduct or is a sales pitch.
Admins will take a screenshot of such behavior in order to document the
situation. **The community takes such violations extremely seriously**, and the
admins are empowered to handle it swiftly.

### Sending Messages to the Channel

`@all`, `@here` and `@channel` should be used rarely. Members will receive
notifications from these commands. Remember Kubernetes is a global
project---please be kind.

### Processing Slack Requests

Admins are tasked with processing requests for channels and other things such as
bots, tokens or webhooks. Please see the processes outlined below.

#### Processing Channel Requests

Channel requests should be reviewed for their relation and relevance to the
Kubernetes community. Typically channels should be dedicated to SIGs, WGs,
sub-projects, community topics, and other things related to Kubernetes programs
and projects.

For Kubernetes project centric requests, validate them against the
[sig-list], or request a link to a related issue/PR, or mailing
list discussion for the requested Channel.

Small external projects are encouraged to use the channel of the SIG or WG
most relevant to them. Other things such as programming language-specific
channels are discouraged and should in turn be steered to `#kubernetes-client`
or communication avenues commonly used by their specific language.

In general, use your best judgment.

Once two Slack admins have reviewed and agreed to sponsor the channel, they will
sign off on the Channel Request PR. Once merged, the channel will be created.

Channels managed by [Tempelis] will automatically have default messages
pinned. For any manually-provisioned channels, such as private channels, add the
below message and pin it.

```
This channel abides to the Kubernetes Code of Conduct -
https://git.k8s.io/community/code-of-conduct.md
Contact conduct@kubernetes.io or an admin in the #slack-admins channel if there
is a problem.
```

#### Processing User Group Requests

User Group requests should be reviewed for their relation and relevance to the
Kubernetes community along with their importance to the requesting group. They
are a useful alias, but can also easily be spammed or abused.

Before signing off on a User Group PR request, ensure all members of the User
Group have signed off acknowledging they will be added to the group.

After all the User Group members have accepted being added to the group and two
Slack Admins have signed off on the request, the PR will be merged. Once merged,
the User Group will be created.

#### Processing Bot, Token, or Webhook Requests

Requests should first be evaluated for their relevance to the project. Typically
approved requests are related to: GitHub, the CNCF, or other tools and platforms
used to aid in the management of Slack. Requests outside of this scope should be
heavily scrutinized and reviewed for **ANY** potential security, privacy, or
usability concerns.

It is best to err on the side of not allowing a bot, token, or webhook request
rather than allowing one.

If the admins come to consensus and agree to the request, an admin should assign
the issue to themselves and reach out to the request contact regarding next
steps.

### Inactivating Accounts

For the reasons listed below, admins may inactivate individual Slack accounts.
Due to Slack's framework, it does not allow for an account to be banned or
suspended in the traditional sense, merely inactivated.
See [Slack's policy on inactivated accounts] for more information.

#### Reasons to inactivate an account

-   Spreading spam content in DMs and/or channels.
-   Not adhering to the code of conduct set forth in DMs and/or channels.
-   Overtly selling products, related or unrelated to Kubernetes.

**BE CAREFUL**

To inactivate a user, and optionally remove their content (spam):

-   First, double check you have the correct user by verifying their Slack
    Member ID.
-   Spammers may try and fake or assume the identity of another user.

Once verified, find a message from the offending user. Then select **More
actions**, the "**...**" to the right of a message from the offending user. Then
select **Report message**.

This will open a contextually-aware prompt only available to Slack Admins with
the options to deactivate the user and remove all content from them over the
past "X" minutes/hours.

Report any actions taken to the other slack admins, and if needed the
[Code of Conduct Committee][cocc].

[coc]: /code-of-conduct.md
[admins]: ./moderators.md#Slack
[Slack Archive Download]: https://drive.google.com/drive/folders/1Xnkwsxis3tu0pT7rwp-crRq4IciZ5b1o?usp=sharing
[cocc]: /committee-code-of-conduct/README.md
[CNCF Slack]: https://slack.cncf.io/
[Tempelis]: https://sigs.k8s.io/slack-infra/tempelis
[slack-config]: ./slack-config/
[Channel Documentation]: ./slack-config/README.md
[sig-list]: https://www.kubernetes.dev/community/community-groups
[Slack Config Documentation]: ./slack-config/README.md
[OWNERS]: /contributors/guide/OWNERS
[usergroups.yaml]: ./slack-config/usergroups.yaml
[User Group Documentation]: ./slack-config/README.md#usergroups
[GitHub Issue]: https://github.com/kubernetes/community/issues/new/choose
[moderation guidelines]: ./moderation.md
[Slack's policy on inactivated accounts]: https://get.slack.help/hc/en-us/articles/204475027-Deactivate-a-member-s-account
