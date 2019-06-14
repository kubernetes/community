# Discuss guidelines

Discuss (discuss.kubernetes.io), is the Kubernetes community forum backed by
the [Discourse] discussion platform. It serves as the primary communication
platform for Kubernetes users; replacing the kubernetes-users mailing list in
September 2018.

Discuss, like other Kubernetes communication platforms, is public and searchable.
Communication should be polite and respectful. Follow the general guideline of
_"be excellent to each other"_.

**Reference Links:**
- [KEP 0007] - A community forum for Kubernetes
- [Archive k-users] - kubernetes-users mailing list migrated to Discuss


## Code of conduct

Kubernetes adheres to the [Kubernetes Code of Conduct]
throughout the project, and includes all communication mediums.


## Privacy Policy

Discuss adheres to the the [Linux Foundation Privacy Policy].


## Admins

- Check the [centralized list of administrators][admins] for contact information.
- Discuss administrators are listed on [Discuss About page].

To connect: please reach out to them using Discourse's built in message system.
If there is an issue with the platform itself, please use the
[sig contributor experience mailing list] or the `#sig-contribex` slack channel.

---

## General communication guidelines

### PM (Private Message) conversations

Please do not engage in proprietary company specific conversations in the
Kubernetes Discourse instance. This is meant for conversations around related
Kubernetes open source topics and community. Proprietary conversations should
occur in one of your company communication platforms. As with all
communication, please be mindful of appropriateness, professionalism, and
applicability to the Kubernetes community.


### Escalating and/or reporting a problem

Discourse has a [built in system for flagging inappropriate posts] that will
notify the admins of a potentially bad post or conversation. If the post
occurred during a period where one of the Admins may not be available, reach out
to one of the [moderators][admins] in the closest timezone directly. As a
moderator, they can flag the post which will [unlist] it immediately until an
Admin is available to review it.

If there is an issue in one of the Regional Boards, engage with one of the
Regional moderators as a first step. They will be able to add context and aid
in the escalation process.

If the problem is with one of the Admins or Moderators, reach out to one of the
other Admins and describe the situation.

If it is a [Code of Conduct] issue, contact conduct@kubernetes.io and describe
the situation.

---

## Moderation

Discourse has a built in set of advanced auto-moderation capabilities that
rely on their _"[user trust system][user-trust]"_. For example, newly created
accounts are rate limited on posting or replying to topics until their "trust
level" increases. A user's trust level will increase based on a number of
factors including time spent on the forum, posts or replies made, likes
received, or one of several other metric.

Moderators, both those for the General Forum and Regional Board are manually
promoted by an Admin to [Trust Level 4][user-trust]. With that comes the full
 responsibilities of a board moderator.


### Moderator expectations and guidelines

Moderators should adhere to the general Kubernetes project
[moderation guidelines].


### Other moderator responsibilities

#### Ingest queue

Moderators have access to a private category called _"Ingest"_ that has topics
posted automatically from a variety of Kubernetes/CNCF sources such as
Kubernetes releases, Security Announcements, the [kubernetes.io blog], and other
useful sources such as [Last Week in Kubernetes Development (LWKD)][lwkd].
Moderators are encouraged to tag and move these articles to their relevant
category.

---

## New category requests

### Requesting a general category

New category requests should be posted to the [Site Feedback and Help] section.
Proposed Categories should be community focused and must be related to
Kubernetes project. They must **not** be Company specific with the exception of
Cloud providers; however their topics should not be related to proprietary
information of the provider.

Once a request has been made, you are encouraged to solicit user support for
the category from the community. The [admins] will review the request, if two
express their support for category, it will be created.

Once created, the _"About the <topic> Category"_ should be updated with a brief
description of the newly created category.


### Requesting a SIG, WG, or sub-project category

If you are associated with a [SIG, WG or subj-project] and would like a Discuss
category to collaborate with others asynchronously; post a message with the
category creation request to the [Site Feedback and Help] section. An
admin will reach out and provide you with a URL and mail address to use for
your discussions.


### Requesting a regional category

The [Regional Discussions Category] is intended for those users that belong to a
specific region or share a common language to openly interact and connect with
each other in their native language.

The anti-spam and anti-harassment features built into Discourse do not handle
other languages as well as it does English. It can pick up on general spam
but lacks regional context. For this reason, the Regional categories require
additional native-language moderators.

To request the creation of a new Regional board, post the request the top level
[Regional Discussions Category]. If possible solicit additional support from
the regional community and propose potential moderators. Before a Regional
Board can be created, there must be at least one moderator, preferably two with
at least one in the Region's primary time zone.

Once moderators have been selected, the Regional category can be created.

The first post  of the new  board _"About the <region> Category"_ post should
contain the following text in both english and the region's language:
```
Welcome to the <region> category of the Kubernetes Forum! In here you can chat
and discuss topics of interest to you about Kubernetes in [region language].
This is a place to share Kubernetes related news, projects, tools, blogs and
more. This site is governed by the [CNCF Code of Conduct], and we are committed
to making this a welcoming place for all. If you have any specific questions or
concerns, please contact one of the moderators for the <region> category listed
below.

**Moderator Team:**
- <moderator 1>
- <moderator 2>

[CNCF Code of Conduct]: <Translated CNCF Code of Conduct>
```

The _"CNCF Code of Conduct"_ link should be linked to one of the
[translated versions of the CNCF Code of Conduct]. If none is available, create
an issue under the [CNCF foundation] project requesting the new translation,
and link to the English version until a translated version is made available.

Lastly, update the [discuss admins][admins] section in the [moderators.md][admins]
list with the new region, the moderators and their timezone.


[Discourse]: https://discourse.org
[KEP 0007]: https://github.com/kubernetes/enhancements/blob/master/keps/sig-contributor-experience/0007-20180403-community-forum.md
[archive k-users]: https://github.com/kubernetes/community/issues/2492
[Kubernetes Code of Conduct]: /code-of-conduct.md
[Linux Foundation Privacy Policy]: https://www.linuxfoundation.org/privacy/
[admins]: ./moderators.md#discusskubernetesio
[Discuss About page]: https://discuss.kubernetes.io/about
[sig contributor experience mailing list]: https://groups.google.com/forum/#!forum/kubernetes-sig-contribex
[built in system for flagging inappropriate posts]: https://meta.discourse.org/t/what-are-flags-and-how-do-they-work/32783
[unlist]: https://meta.discourse.org/t/what-is-the-difference-between-closed-unlisted-and-archived-topics/51238
[user-trust]: https://blog.discourse.org/2018/06/understanding-discourse-trust-levels/
[moderation guidelines]: https://github.com/kubernetes/community/blob/master/communication/moderation.md
[kubernetes.io blog]: https://kubernetes.io/blog/
[lwkd]: http://lwkd.info/
[Site Feedback and Help]: https://discuss.kubernetes.io/c/site-feedback
[SIG, WG or subj-project]: https://github.com/kubernetes/community/blob/master/sig-list.md
[Regional Discussions Category]: https://discuss.kubernetes.io/c/regional-discussions
[translated versions of the CNCF Code of Conduct]: https://github.com/cncf/foundation/tree/master/code-of-conduct-languages
[CNCF foundation]: https://github.com/cncf/foundation
