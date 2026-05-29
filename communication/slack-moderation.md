# Slack Moderation

## Admin Expectations and Guidelines

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

## Sending Messages to the Channel

`@all`, `@here` and `@channel` should be used rarely. Members will receive
notifications from these commands. Remember Kubernetes is a global
project---please be kind.

## Processing Slack Requests

Admins are tasked with processing requests for channels and other things such as
bots, tokens or webhooks. Please see the processes outlined below.

### Processing Channel Requests

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

### Processing User Group Requests

User Group requests should be reviewed for their relation and relevance to the
Kubernetes community along with their importance to the requesting group. They
are a useful alias, but can also easily be spammed or abused.

Before signing off on a User Group PR request, ensure all members of the User
Group have signed off acknowledging they will be added to the group.

After all the User Group members have accepted being added to the group and two
Slack Admins have signed off on the request, the PR will be merged. Once merged,
the User Group will be created.

### Processing Bot, Token, or Webhook Requests

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

### Reasons to inactivate an account

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
[community inviter]: https://slack.k8s.io
[Slack Archive Download]: https://drive.google.com/drive/folders/1Xnkwsxis3tu0pT7rwp-crRq4IciZ5b1o?usp=sharing
[cocc]: /committee-code-of-conduct/README.md
[CNCF Slack]: https://slack.cncf.io/
[Tempelis]: https://sigs.k8s.io/slack-infra/tempelis
[slack-config]: ./slack-config/
[Channel Documentation]: ./slack-config/README.md
[sig-list]: /community/community-groups
[Slack Config Documentation]: ./slack-config/README.md
[OWNERS]: /contributors/guide/OWNERS
[usergroups.yaml]: ./slack-config/usergroups.yaml
[User Group Documentation]: ./slack-config/README.md#usergroups
[GitHub Issue]: https://github.com/kubernetes/community/issues/new/choose
[moderation guidelines]: ./moderation.md
[Slack's policy on inactivated accounts]: https://get.slack.help/hc/en-us/articles/204475027-Deactivate-a-member-s-account
[Message Reporter]: /docs/comms/slack/#reporting-a-problem
