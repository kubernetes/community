---
title: "Discuss Forum Guidelines"
description: |
  This is an overview of Discourse, the software running the Kubernetes
  community forum, its moderation capabilities and how they are applied, along
  with detailed instructions on requesting new categories.
---

Discuss (discuss.kubernetes.io), is the Kubernetes community forum backed by the
[Discourse][] discussion platform. It serves as the primary communication
platform for Kubernetes users, replacing the kubernetes-users mailing list in
September 2018.

Discuss, like other Kubernetes communication platforms, is public and
searchable. Communication should be polite and respectful. Follow the general
guideline of *"be excellent to each other"*.

## Reference Links

-   [A community forum for Kubernetes KEP][]
-   [Archive k-users][] - kubernetes-users mailing list migrated to Discuss

## Code of conduct

Kubernetes adheres to the [Kubernetes Code of Conduct][] throughout the project,
and includes all communication mediums.

## Privacy Policy

Discuss adheres to the [Linux Foundation Privacy Policy][].

## Admins

-   Check the [centralized list of administrators][admins] for contact information.
-   Discuss administrators are listed on the [Discuss 'About page'][].

To connect with Admins, please reach out to them using Discourse's built in
message system. If there is an issue with the platform itself, please use the
[sig contributor experience mailing list][] or visit the #sig-contribex
[Slack Channel][].

## General communication guidelines

### PM (Private Message) conversations

Please do not engage in company specific or proprietary conversations in the
Kubernetes Discourse instance. This space is meant for conversations around
related Kubernetes open source topics and community. Proprietary conversations
should occur in one of your company communication platforms. As with all
communication, please be mindful of appropriateness, professionalism, and
applicability to the Kubernetes community.

### Escalating and/or reporting a problem

Discourse has a [built in system for flagging inappropriate posts][] that will
notify the admins of a potentially bad post or conversation. If the post
occurred during a period where one of the Admins may not be available, reach out
to one of the [moderators][admins] in the closest
timezone directly. As a moderator, they can flag the post which will [unlist][]
it immediately until an Admin is available to review it.

If there is an issue in one of the Regional Boards, engage with one of the
Regional moderators as a first step. They will be able to add context and aid in
the escalation process.

If the problem is with one of the Admins or Moderators, reach out to one of the
other Admins and describe the situation.

If it is a [Code of Conduct] issue, contact
our [Code of Conduct Committee](mailto:conduct@kubernetes.io) and describe the situation.

## Moderation

Discourse has a built-in set of advanced auto-moderation capabilities that rely
on their [user trust system][]. For example, newly created accounts are rate
limited on posting or replying to topics until their trust level increases. A
user's trust level will increase based on a number of factors including time
spent on the forum, posts or replies made, likes received, and several other
metrics.

Moderators, both those for the General Forum and Regional Board, are manually
promoted by an Admin to [Trust Level 4][user trust system]. With that comes the
full responsibilities of a board moderator.

### Moderator expectations and guidelines

Moderators should adhere to the general Kubernetes project 
[moderation guidelines][].

## New category requests

### Requesting a general category

New category requests should be posted to the [Site Feedback and Help][]
section. Proposed Categories should be community focused and must be related to
the Kubernetes project. They **must not** be company specific with the exception
of cloud providers; however their topics **should not** be related to
proprietary information of the provider.

Once a request has been made, you are encouraged to solicit user support for the
category from the community. The [admins][]
will review the request, if two express their support for the category, it will
be created.

Once created, the "About the Category" section should be updated with a brief
description of the newly created category.

### Requesting a SIG, WG, or subproject category

If you are associated with a [SIG, WG or subjproject][] and would like a Discuss
category to collaborate with others asynchronously; post a message with the
category creation request to the [Site Feedback and Help][] section. An admin
will reach out and provide you with a URL and mail address to use for your
discussions.

### Requesting a regional category

The [Regional Discussions Category][] is intended for those users that belong to
a specific region or share a common language to openly interact and connect with
each other in their native language.

The anti-spam and anti-harassment features built into Discourse do not handle
other languages as well as it does English. It can pick up on general spam but
lacks regional context. For this reason, the Regional categories require
additional native-language moderators.

To request the creation of a new Regional board, post the request in the top
level [Regional Discussions Category][]. If possible solicit additional support
from the regional community and propose potential moderators. Before a Regional
Board can be created, there must be at least one moderator, preferably two with
at least one in the Region's primary time zone.

Once moderators have been selected, the Regional category can be created.

The first post for the new Regional category should be titled "About the
Category". The post should contain the following text in both english and the
region's language:

```
Welcome to the <region> category of the Kubernetes Forum! In here you can chat
and discuss topics of interest to you about Kubernetes in [region language].
This is a place to share Kubernetes related news, projects, tools, blogs and
more. This site is governed by the [CNCF Code of Conduct], and we are
committed to making this a welcoming place for all. If you have any
specific questions or concerns, please contact one of the moderators
for the <region> category listed below.

**Moderator Team:**
- <moderator 1>
- <moderator 2>

[CNCF Code of Conduct]: <Translated CNCF Code of Conduct>
```

The "CNCF Code of Conduct" link should be linked to one of the 
[translated versions of the CNCF Code of Conduct][]. If none is available,
create an issue under the [CNCF foundation][] project requesting the new
translation, and link to the English version until a translated version is made
available.

Lastly, update the [discuss admins][centralized list of administrators] section
in the [moderators.md][centralized list of administrators] list with the new
region, the moderators and their timezone.

  [Discourse]: https://discourse.org/
  [A community forum for Kubernetes KEP]: https://github.com/kubernetes/enhancements/tree/master/keps/sig-contributor-experience/0000-community-forum
  [Archive k-users]: https://github.com/kubernetes/community/issues/2492
  [Kubernetes Code of Conduct]: /code-of-conduct.md
  [Linux Foundation Privacy Policy]: https://www.linuxfoundation.org/privacy/
  [admins]: ./moderators.md#discusskubernetesio
  [Discuss 'About page']: https://discuss.kubernetes.io/about
  [sig contributor experience mailing list]: https://groups.google.com/forum/#!forum/kubernetes-sig-contribex
  [Slack Channel]: https://slack.k8s.io/
  [built in system for flagging inappropriate posts]: https://meta.discourse.org/t/what-are-flags-and-how-do-they-work/32783
  [unlist]: https://meta.discourse.org/t/what-is-the-difference-between-closed-unlisted-and-archived-topics/51238
  [user trust system]: https://blog.discourse.org/2018/06/understanding-discourse-trust-levels/
  [moderation guidelines]: ./moderation.md
  [Site Feedback and Help]: https://discuss.kubernetes.io/c/site-feedback
  [SIG, WG or subjproject]: /sig-list.md
  [Regional Discussions Category]: https://discuss.kubernetes.io/c/regional-discussions
  [translated versions of the CNCF Code of Conduct]: https://github.com/cncf/foundation/tree/master/code-of-conduct-languages
  [CNCF foundation]: https://github.com/cncf/foundation
