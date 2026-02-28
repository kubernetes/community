---
title: YouTube Channel Guidelines
description: |
  This is an overview of community YouTube practices and admin responsibilities.
---

YouTube serves as the primary means of distribution for recorded Kubernetes
community content including Zoom recordings, official project workshops and
Contributor Summit sessions.

## Code of Conduct

Kubernetes adheres to the [Kubernetes Code of Conduct][coc] throughout the project,
and includes all communications such as YouTube.

## Admins

-   Check the [centralized list of administrators][admins] for contact information.
-   To contact the admin group in Slack, ping `@youtube-admins` in the
    `#sig-contribex` Slack channel.

## Meeting Playlists

The [Kubernetes YouTube Channel] has separate playlists for each SIG or WG 
meeting recordings, as well as recordings of other recurring events such as
the [New Contributor Orientation], and others.

[Subprojects], in addition to SIGs and WGs may request their own playlists
to better target their contributors and increase general discoverability.

To better serve the community, [collaboration] has been enabled to share the
management of the playlists. Anyone with the appropriate link to the particular
playlist can upload videos *to that particular playlist* (links & playlists are
one-to-one).

Each group's playlist link will be shared with the group's leadership via Slack
and the group leadership Google Group. Other playlists links, will be shared
with the appropriate point(s) of contact.

### Uploading Guidelines for Collaborators

**NOTE:** If you're using a Google Workspace account (formerly known as G Suite)
you may need to [update the permissions in your YouTube settings]. If you have
any questions, reach out to the [YouTube admins] or
[SIG Contributor Experience]. You may need to reach out to someone at your
organization if you do not have access to Google Workspace Admin permissions.

**NOTE:** Both public and private steering meeting recordings should be made
public.

With collaboration comes great responsibility. Playlist collaborators in the
community must use it responsibly and are subject to the following guidelines:

-   Group leaders or other appropriate point(s) of contact are the primary
    managers for the playlist, once collaboration is configured. YouTube admins
    should **only** be contacted if the issue cannot be resolved by one of the
    playlist owners.

-   Upload responsibilities belong to the group leaders or other appropriate
    contacts. YouTube admins should **only** be contacted if the issue cannot be
    resolved by the playlist owners.

-   Please post only related content, for example meeting recordings, in the
    appropriate playlists.
    -   Posting of any inappropriate content (i.e. NSFW content)
        will result in ***immediate*** suspension of privileges.

-   All posted videos should use the naming convention: [Name of Playlist's
    Group] Meeting Name for YYYYMMDD
    -   **Example:** [SIG Contributor Experience] Biweekly Meeting for 20240816

-   Playlists should be organized chronologically for ease of use. This can be
    done by updating the default ordering of the Playlist:
    -   From within the Playlist settings, click on the **Basic** Tab.
    -   From the **Ordering** dropdown select "Date added (newest)".
    -   Save the changes and the order should automatically be updated.

-   Please do not remove any already-published content from the playlists
    without checking with the YouTube admins.

-   For any small issues that arise, for example improper naming or ordering,
    you may be asked by the YouTube admins to attempt to resolve the issue
    yourself.

-   Any egregious or habitual violations (3 or more per quarter) of the above
    rules will result in suspension of collaboration privileges for the
    particular individual or for the entire playlist if the individual can't be
    identified.
    -   If an individual is suspended, the playlist link will be remade and the
        new link will be shared with the non-offending individuals.
    -   If playlist collaboration is suspended, the uploading and management of
        the playlist will be handled by the YouTube admins. Uploading the
        problematic group's playlist will not be considered a priority, and
        delays in uploading should be expected.

## Admin Responsibilities

The role of the Youtube Admins is to help maintain a robust YouTube channel that
is valuable to contributors and upholds our transparency goals as laid out by
our governance documents.

### Moderator Expectations and Guidelines

Moderators should adhere to the general Kubernetes project
[moderation guidelines].

Moderation responsibilities for YouTube admins is minimal and is centered around
checking and removing any potential comments that would violate the
[Code of Conduct][coc]. Any potential violations should be sent to
<conduct@kubernetes.io>.

### Trimming and Editing Recordings

YouTube admins are asked to help [trim] and [edit] recordings that come into
the video queue.

#### Examples:

Certain events such as the Contributor Summits are not uploaded directly to
YouTube and require editing.

A Zoom recording may have significant dead-space leading the meeting itself and
also at the end as we end the stream.

A SIG Meeting needs to be edited to make sure it's clear (ie "Kubernetes Special
Interest Group ContribEx 20220131 Meeting").

Make sure to save a copy of the video first before making adjustments as this
can delete the original if not careful. When in doubt, ask.

### Automation

**Note:** There is always room for improvement! As such, the community is open
to trying various integration features or other productivity tools that might
improve the job of admins and help make things more streamlined. Please, reach
out if you can help or have any ideas.

#### Splain.io

One feature used by several SIGs is splain.io. This tool creates a pipeline
between Zoom and youtube for easier workflows. To use splain.io please follow
the steps outlined below.

**Items needed to use splain.io:**

-   UserName and Password for the Zoom license account. You will need a Zoom
    admin or the admin of the license (ex: SIG Chair) to either provide it to
    you or reset it to a new one.
-   Kubernetes YouTube admin permissions
-   A correctly set up recurring meeting with a start and end time (this is
    important) - check [Zoom guidelines] for more details

**Steps:**

1.  Install splain: <https://marketplace.zoom.us/apps/WPKzwuoLQDuj_gPs68AQxw>
2.  Connect the Zoom account
3.  Connect YouTube account
4.  Click the manage tab next to App Info, and then scroll to locate the config
    button.
5.  Click **Config**
6.  The splain dashboard opens: make sure the box for **Make videos private** is
    checked so we can edit before it goes live.
7.  Test that it works.
8.  Enjoy.

The following SIGs and groups are currently running splain.io:

-   [SIG Auth]
-   [SIG Contributor Experience]
-   [SIG Docs]
-   [SIG Network]
-   [SIG Release]
-   [Steering Committee]
-   [WG Data Protection]
-   The main Zoom admin account (if you log in to splain using this account,
    all of the other accounts will be logged here)

### Descriptions and Playlists

Each video that comes into the queue needs to be added to a playlist, set to
public, and have contextual information added to the description.

**Example description:**

Kubernetes Contributor Experience Special Interest Group Weekly Meeting. Check
here (link to sig list) for a complete list of SIGs and when they meet. Join us!

Please note the following items:

-   The date must be in title and description
-   The date format must be YYYYMMDD
-   The section titled "Is this video made for kids?" **should not** be altered
    for search reasons

Below is an example:

-   Incorrect description: `Data Protection WG Bi-Weekly Meeting for 2/26/2020`
-   Correct description: `Data Protection WG Bi-Weekly Meeting for 20200226`

A short video tutorial of the editing of descriptions and playlists can be found
here:

-   <https://youtu.be/IbZ2wnvu_Zs>

### Thumbnails

There have been ongoing conversations about how to create and manage better
thumbnails for videos using some kind of standard. If you're interested in
helping with thumbnails, please reach out to the YouTube Admins. They would love
to hear from you.

### Streaming Events

YouTube admins with a system capable of streaming may be asked to stream public
Kubernetes Community events such as publicly streamed events. For detailed 
information about streaming, see our
[Streaming Config].

### Migrating Content

In certain cases, such as archiving an inactive SIG or Working Group, it might
be useful to grab content from other channels. It is currently NOT POSSIBLE to
move content from one YouTube channel to another, so the content must be
downloaded and then reuploaded to the Kubernetes channel.

1.  Download [youtube-dl], which makes it easier to bulk download videos.
2.  Download the channel or playlist with: `youtube-dl <url>`
3.  Clean up the filenames as they are used to generate new titles. Do this
    locally since it is easier than doing it per video in the YouTube web UI.
4.  Create a new playlist for the content on the Kubernetes channel, set it to
    Private.
5.  Upload the videos to the Kubernetes Channel.
6.  Ensure titles and metadata are accurate, allow time to pass for YouTube to
    generate thumbnails and process the videos.
7.  Once videos are finalized, set the playlist to Public to publish them.

[coc]: /code-of-conduct.md
[admins]: /communication/moderators.md
[Kubernetes YouTube Channel]: https://www.youtube.com/channel/UCZ2bu0qutTOM0tHYa_jkIwg
[New Contributor Orientation]: /mentoring/new-contributor-orientation/README.md
[Subprojects]: /governance.md#subprojects
[collaboration]: https://support.google.com/youtube/answer/6109639
[update the permissions in your YouTube settings]: https://support.google.com/a/answer/6212415
[YouTube admins]: /communication/moderators.md#YouTube-Channel
[SIG Contributor Experience]: /sig-contributor-experience/README.md
[moderation guidelines]: /communication/moderation.md
[trim]: https://support.google.com/youtube/answer/9057455?hl=en
[edit]: https://support.google.com/youtube/topic/9257530?hl=en&ref_topic=9257610
[Zoom guidelines]: /communication/zoom-guidelines.md
[SIG Auth]: /sig-auth/README.md
[SIG Docs]: /sig-docs/README.md
[SIG Network]: /sig-network/README.md
[SIG Release]: /sig-release/README.md
[Steering Committee]: /committee-steering/governance/README.md
[WG Data Protection]: /wg-data-protection/README.md
[Streaming Config]: /communication/youtube/streaming-config.md
[youtube-dl]: https://ytdl-org.github.io/youtube-dl/index.html
