# Slack Guidelines

Slack serves as the main communication platform for the Kubernetes community
outside of the mailing lists. It’s important that conversations stays on topic
in each channel, and that everyone abides by the [Code of Conduct][coc]. There
are over 50,000 members who should all expect to have a positive experience.

Chat is searchable and public. Do not make comments that you would not say on a
video recording or in another public space. Please be courteous to others.


## Code of Conduct

Kubernetes adheres to the [Kubernetes Code of Conduct][coc] throughout the
project, and includes all communication mediums.


## Admins

- Check the [centralized list of administrators][admins] for contact information.

Slack Admins will have listed that they are a Slack admin in their Slack profile,
along with their specific timezone.

To connect: please reach out in the `#slack-admins` channel, mention an admin
directly in the `#slack-admins` channel when you have a question, or DM (Direct
Message) one privately.

---

## General Communication Guidelines

### Workspace Channel History

The Kubernetes Slack Workspace is archived and made available when the
administrators have time. There is no explicit interval.

[Slack Archive Download]


### DM (Direct Message) Conversations

Please do not engage in proprietary company specific conversations in the
Kubernetes Slack instance. This workspace is meant for conversations related to
Kubernetes open source topics and community. Proprietary conversations should
occur in your company Slack and/or communication platforms.  As with all
communication, please be mindful of appropriateness, professionalism, and
applicability to the Kubernetes community.


### Specific Channel Rules

Some channels have specific rules or guidelines. If they do, they will be listed
in the purpose or pinned docs of that channel.

- `#kubernetes-dev` - Questions and discourse around upstream contributions and
development to kubernetes,
- `#kubernetes-careers` - Job openings for positions working with/on/around
  Kubernetes. Post the job once and pin it. Pins expire after 30 days. Postings
  must include:
  - A link to the posting or job description.
  - The business name that will employ the Kubernetes hire.
  - The location of the role or if remote is OK.


### Escalating and/or Reporting a Problem

Join the `#slack-admins` channel or contact one of the admins in the
[closest timezone][admins] via DM directly and describe the situation. If the
issue can be documented, please take a screenshot to include in your message.

**What if you have a problem with an admin?**

Send a DM to another [listed admin][admins] and describe the situation. If it’s
a [code of conduct][coc] issue, please send an email to <conduct@kubernetes.io>
and describe the situation.

---

## Moderation

### Admin Expectations and Guidelines

Admins should adhere to the general Kubernetes project [moderation guidelines].

Admins should additionally make sure to mention they are a Slack admin and their
timezone in their “What I do” section of their Slack profile.

Be mindful of how you handle communication during stressful interactions.
Administrators act as direct representatives of the community, and need to
maintain a very high level of professionalism at all times. If you feel too
involved in the situation to maintain impartiality or professionalism, that’s
a great time to enlist the help of another admin.

Try to take any situations that involve upset or angry members to DM or video
chat. Please document these interactions for other Slack admins to review.

Content will be automatically removed if it violates code of conduct or is a
sales pitch. Admins will take a screenshot of such behavior in order to document
the situation. The community takes such violations extremely seriously, and the
admins are empowered to handle it swiftly.


### Sending Messages to the Channel

`@all`, `@here` and `@channel` should be used rarely. Members will receive
notifications from these commands. Remember Kubernetes is a global project -
please be kind.


### Slack Requests

Admins are tasked with processing requests for channels and other things such as
bots, tokens or webhook.


#### Channel Requests

Channel requests should be reviewed for their relation and relevance to the
Kubernetes community. Typically channels should be dedicated to SIGs, WGs, UGs,
sub-projects, community topics, and other things related to Kubernetes programs
and projects.

For Kubernetes project centric requests, validate them against the [sig-list],
or request a link to a related issue/PR, or mailing list discussion for the
requested Channel.

Small external projects are encouraged to use the channel of the SIG, WG, or UG
most relevant to them. Other things such as programming language specific
channels are discouraged and should in turn be steered to `#kubernetes-client`
or communication avenues commonly used by their specific language.

In general, use your best judgment.

Once two admins have agreed to sponsor the channel, an admin should assign the
issue to themselves, and create the channel. A message should then be pinned
with the below text:
```
This channel abides to the Kubernetes Code of Conduct - http://git.k8s.io/community/code-of-conduct.md
Contact conduct@kubernetes.io or an admin in the #slack-admins channel if there is a problem.
```


#### Bot, Token, or Webhook Requests

Requests should first be evaluated for their relevance to the project. Typically
approved requests are related to: GitHub, the CNCF, or other tools/platforms
used to aid in the management of Slack. Requests outside of this scope should be
heavily scrutinized and reviewed for **ANY** potential security, privacy, or
usability concerns.

It is best to err on the side of not allowing a Bot, Token or Webhook request
than allowing one.

If consensus among the admins can be reached regarding the request; an admin
should assign the issue to themselves, and reach out to the request contact
regarding next steps.


### Inactivating Accounts

For the reasons listed below, admins may inactivate individual Slack accounts.
Due to Slack’s framework, it does not allow for an account to be banned or
suspended in the traditional sense, merely inactivated. See [Slack’s policy on
inactivated accounts] for more information.

- Spreading spam content in DMs and/or channels.
- Not adhering to the code of conduct set forth in DMs and/or channels.
- Overtly selling products, related or unrelated to Kubernetes.

---

## Requesting a Channel

- Create a [GitHub Issue] using the Slack Request template.
- Include the desired name of the channel. Typical naming conventions follow:
  - `#kubernetes-foo`
  - `#sig-foo`
  - `#meetup-foo`
  - `#location-user`
  - `#projectname`
- In the issue describe the purpose of the Channel. Channels should be
  dedicated to [SIGs, WGs, UGs][sig-list], sub-projects, community topics, or
  related Kubernetes programs/projects.
  - Linking to resources such as the PR adding the subproject will speed in the
    validation and processing of the channel creation request.
  - Channels are **NOT**:
    - Company specific; cloud providers are ok with product names as the channel.
      Discourse will be about Kubernetes-related topics and not proprietary
      information of the provider.
    - Private channels with the exception of: code of conduct matters, mentoring,
      security/vulnerabilities, github management, or steering committee.
  - Special accommodations will be made where necessary.

After you submit your request the Slack admins will review and follow-up with
any questions in the issue. If two admins add their support for the creation of
the channel, an admin will assign it to themselves and close the issue once it
has been created.


## Requesting a Bot, Token, or Webhook

**READ BEFORE SUBMITTING A REQUEST**

Bots, tokens, and webhooks are reviewed on a case-by-case basis with most
requests being rejected due to security, privacy, and usability concerns. Bots
and the like tend to make a lot of noise in channels. The Kubernetes Slack
instance has over 50,000 members and it is the role of the Slack admins to
ensure everyone has a great experience.

Typically approved requests include: GitHub, CNCF requests, or other
tools/platforms used to aid in the management of Slack itself.

- Create a [GitHub Issue] using the Slack Request template.
- In the description, describe the request, its intended purpose and benefit to
  the community. Supplying links to supporting content such as a document
  outlining what OAuth scopes that are requested, and why are
  **STRONGLY ENCOURAGED**.

After you submit your request the Slack admins will review and follow-up with
any questions in the issue. If consensus can be reached among the admins, the
request will be approved and follow-up communication on implementation will be
discussed in Slack itself.


[coc]: /code-of-conduct.md
[admins]: ./moderators.md#Slack
[GitHub Issue]: https://github.com/kubernetes/community/issues/new/choose
[sig-list]: /sig-list.md
[Slack’s policy on inactivated accounts]: https://get.Slack.help/hc/en-us/articles/204475027-Deactivate-a-member-s-account
[Slack Archive Download]: https://drive.google.com/drive/folders/1idJkWcDuSfs8nFUm-1BgvzZxCqPMpDCb?usp=sharing
[moderation guidelines]: ./moderation.md