---
title: "Slack Migration FAQ"
description: |
  Explains the details of our upcoming change to Slack account type and eventual migration.
---

## What is Changing About Slack?

**UPDATE**: We’ve received notice from Salesforce that our Slack workspace **WILL NOT BE DOWNGRADED** on June 20th. Stand by for more details, but for now, there is no urgency to back up private channels or direct messages.

~~Kubernetes Slack will lose its special status and will be changing into a standard free Slack on June 20, 2025~~. Sometime later this year, our community may move to a new platform. If you are responsible for a channel or private channel, or a member of a User Group, you will need to take some actions as soon as you can.

## What Actions do Channel Owners and User Group Members Need to Take Soon?

If you are responsible for a channel or a team with a Slack User Group, you'll need to take some actions very soon to deal with the downgrade.

### Before Friday:

**Saved Files and Pinned Posts**: If your channel has important saved files and information in Pinned Posts, you need to copy that information to other media such as GitHub files or GoogleDocs *as soon as possible* and instead have a Bookmark linking to that location. Once our account is downgraded, we will not be able to access pinned posts and files more than 90 days old.

**Private Channels**: Private Channels are *not* backed up to our archives because they can't be. If you have important information in private channels which is more than 90 days old, copy it to another location as soon as possible. We recommend backing up the contents of your private channel [using Slackdump](slack-backup/README.md), an open source tool.

### Within The Next Few Weeks

**User Groups**: Kubernetes Slack has a few dozen [User Groups](https://github.com/kubernetes/community/blob/master/communication/slack-config/usergroups.yaml). These will stop functioning, although we will retain their membership. However, any ability to tag these user groups or message them will go away, and community documentation and processes will need to reflect that they are not contactable. If, due to the loss of a User Group you need a new Slack channel, please [file a request](https://github.com/kubernetes/community/issues).

Additionally, if you have important saved information in Direct Message conversations which are more than 90 days old, you might consider copying that information somewhere else.

## Why now?

Slack has been a long-time supporter of the Kubernetes ecosystem. However, the scale of the Kubernetes community has put strains on Slack's infrastructure. This week, they let the CNCF Projects staff know that they cannot continue to host us the same way.

## What Will Running on a Free Slack Mean?

There are [multiple limitations on free Slacks](https://slack.com/help/articles/27204752526611-Feature-limitations-on-the-free-version-of-Slack). The ones which will most visibly affect the community are:

**Only retaining 90 days of history**: this means that you will not be able to search for messages and statements in your channels that are more than 90 days old. We are not permanently losing community history, since we have archival backups of all Slack messages. However, those archives are not currently searchable by our community. Note that the 90-day limit includes files which were posted in a Slack message, so if you have important files for a channel which are part of a pinned Slack message, make a copy of those files somewhere else as soon as you can.

**No User Groups**: Slack User Groups will go away and not be taggable any more. We will retain information about who the members of those teams were in Github, but they will not have meaning until we migrate to a new platform.

**Disabling all Workflows**: On our Slack, this mostly means that welcome messages for channels and the Slack will stop working.

**Limitation to 10 Apps**: The Slack Admins team will need to choose around 8 applications which are currently in use on our Slack to disable. They will update this FAQ with which ones will be specifically disabled once that decision is made.

**No Custom Sections**: if you have rearranged your channels into multiple Custom Sections in your Slack interface, those sections will be going away and all channels will be listed together.

While there are other limitations to free Slack, most of them pertain to features we were not using on ours, such as Huddles and Slack Connect.

## What Information Will We Lose?

Not much. Our Slack history is backed up to offline archives, which, while presently not searchable, do include the entire time Kubernetes has been on Slack. All of our team membership and channel information is saved [in github](https://github.com/kubernetes/community/tree/master/communication/slack-config).

The things that we will be losing, if not individually backed up, are:

* Private Channel Conversations: any messages more than 90 days old in Private Channels
* Direct Message Conversations: any DM posts more than 90 days old
* Files Posted to Slack: any files more than 90 days old which were shared only on Slack

If you have important information in any of these places, please take action to back it up now.

## Will We Be Migrating to a New Platform?

Most likely, yes. The CNCF Projects Staff, which includes several long-time Kubernetes contributors, has proposed that we move to Discord.

In the next few weeks, the Kubernetes Steering Committee will make a decision on whether we are migrating to a new discussion platform, and if so where.

## Why Discord?

The Project Staff have proposed that we move to Discord because that platform offers us several advantages, including:

* Easily handles communities our size, or even ten times larger
* More robust moderation tooling
* Ability to support deeper integrations with tools we use, like GitHub
* Support for custom permissions schemes, allowing programmatic channel ownership
* Ability to replace our existing bots and workflows
* CNCF Slack is likely to move to Discord as well

The move to Discord is not decided. The Steering Committee will be discussing it soon.

## Why Not Matrix?

Our project's size and level of traffic would be difficult for Matrix, at its current level of development, to support. Whereas we would not even be in the ten largest Discord communities; some gaming communities are ten times our size.

## Why Not Some Other Platform?

Kubernetes Slack, with over 200,000 members and tens of thousands of posts per day, can only be handled by a highly scalable enterprise chat platform. Our chat platform also needs to be as open to all contributors as our project is, which further limits our technology options.

## Could We Just Stay On Free Slack?

Yes, and that's undoubtedly one of the options which the Steering Committee will discuss. However, several of the limitations of free Slack are quite problematic for our community.

## Who Will Lead Migration Efforts?

SIG Contributor Experience (ContribEx) and the Slack Admins team will work with CNCF staff to implement whatever migration plan that the Steering Committee decides on.

## When Will We Migrate Off Slack?

That will be decided by the Steering Committee, working together with the CNCF Staff and the Slack Admins.

## Where Can I Discuss The Slack Downgrade and Migration?

Join the [community discussion thread](https://github.com/kubernetes/community/issues/8490). Our plans will also get discussed at a public Steering Committee meeting (to be scheduled).
