---
title: "YouTube Guidelines"
description: |
  Overview of community YouTube practices and admin responsibilities.
---

<!-- omit in toc -->
# YouTube Channel Guidelines

YouTube serves as primary means of distribution for recorded Kubernetes
community content including Zoom recordings, official project Workshops and
Contributor Summit sessions.

- [Code of Conduct](#code-of-conduct)
- [Admins](#admins)
- [Meeting Playlists](#meeting-playlists)
  - [Uploading Guidelines for Collaborators](#uploading-guidelines-for-collaborators)
- [Admin Responsibilities](#admin-responsibilities)
  - [Moderator Expectations and Guidelines](#moderator-expectations-and-guidelines)
  - [Trimming and Editing Recordings](#trimming-and-editing-recordings)
  - [Automation](#automation)
  - [Descriptions & Playlists](#descriptions--playlists)
  - [Thumbnails](#thumbnails)
  - [Streaming Events](#streaming-events)
  - [Migrating Content](#migrating-content)

## Code of Conduct

Kubernetes adheres to the [Kubernetes Code of Conduct][coc] throughout the
project, and includes all communication mediums.

## Admins

- Check the [centralized list of administrators] for contact information.
- To contact the admin group in Slack, ping `@youtube-admins` in the `#sig-contribex`
  Slack channel.

## Meeting Playlists

The [Kubernetes YouTube Channel] has separate playlists for each SIG, WG, UG
meeting recordings, as well as recordings of other recurring events such as the
Kubernetes [Community meeting], [Office Hours], [Meet our Contributors] and
others.

[Subprojects], in addition to SIGs, WGs, UGs may request their own playlists
to better target their contributors and increase general discoverability.

To better serve the community, [collaboration] has been enabled to share the
management of the playlists. Anyone with the appropriate link to the particular
playlist can upload videos _to that particular playlist_ (links & playlists are
one-to-one).

Each group's playlist link will be shared with the group's leadership via Slack
and the group leadership Google Group. Other playlists links, for example Office
Hours, will be shared with the appropriate point(s) of contact.

### Uploading Guidelines for Collaborators

**NOTE:** If you're using a G Suite account you may need to [loosen the
permissions in your YouTube settings]. If you have any questions reach out to
the [YouTube admins] or [SIG Contributor Experience].

**NOTE:** Both public and private steering meeting recordings should be made public.

With collaboration comes great responsibility. Playlist collaborators in the
community must use it responsibly and are subject to the following guidelines:

- Group leaders or other appropriate point(s) of contact are the primary
  managers for the playlist, once collaboration is configured. YouTube admins
  should **only** be contacted if the issue cannot be resolved by one of the
  playlist owners.

- Upload responsibilities belong to the group leaders or other appropriate
  point(s) of contact. YouTube admins should **only** be contacted if the
  issue cannot be resolved by the playlist owners.

- Please post only related content; for example: meeting recordings, in the
  appropriate playlists.
  - Posting of any exceedingly inappropriate content (i.e. NSFW content) will
	result in ***immediate*** suspension of privileges.

- All posted videos should use the naming convention:
  `Kubernetes [Name of Playlist’s Group] YYYYMMDD`
  - **Example:** `Kubernetes SIG Service Catalog 20161129`

- Playlists should be organized chronologically for ease of use. This can be
  done by updating the default ordering of the Playlist:
  - From within the Playlist settings, click on the **Basic** Tab.
  - From the **Ordering** dropdown select "Date added (newest)".
  - Save the changes and the order should automatically be updated.

- Please do not remove any already-published content from the playlists without
  checking with the YouTube admins.

- For any small issues that arise, for example improper naming or ordering, you
  may be asked by the YouTube admins to attempt to resolve the issue yourself.

- Any egregious or habitual violations (3 or more per quarter) of the above
  rules will result in suspension of collaboration privileges for the particular
  individual or for the entire playlist if the individual can’t be identified.
  - If an individual is suspended, the playlist link will be remade and the new
    link will be shared with the non-offending individuals.
  - If playlist collaboration is suspended, the uploading and management of
    the playlist will be handled by the YouTube admins. Uploading the
	problematic group's playlist will not be considered a priority, and delays
	in uploading should be expected.


## Admin Responsibilities  

Purpose: Help maintain a robust YouTube channel that is valuable to contributors
and upholds our transparency goals as laid out by our governance docs.  

### Moderator Expectations and Guidelines

Moderators should adhere to the general Kubernetes project
[moderation guidelines].

Moderation responsibilities for YouTube admins is minimal and is centered around
checking and removing any potential comments that would violate the [Code of
Conduct][coc]. Any potential violations should sent to <conduct@kubernetes.io>.

### Trimming and Editing Recordings

YouTube admins are asked to help [trim] and [edit] recordings that come into the
video queue.  
Examples:  
Certain events such as the Contributor Summits are not uploaded directly to
YouTube and require editing.
A Zoom recording may have significant dead-space leading the meeting itself
and also at the end as we end the stream.  
A SIG Meeting needs to be edited to to make sure it's clear (ie "Kubernetes
Special Interest Group ContribEx 20190303 Meeting").   

Make sure to save a copy of the video first before making adjustments as this
can delete the original if not careful. When in doubt, ask.

### Automation

We have been playing around with various integration features with our other
productivity tools and would love to do more! Reach out if you can help.

One feature that we've implemented with several SIGs is splain.io. This tool
creates a pipeline between zoom and youtube.

Here's what you need:  
- UserName and Password for the zoom license account. You'll need a zoom admin
or the admin of the license (ex: SIG Chair) to either provide it to you or reset
it to a new one.
- Kubernetes YouTube admin credentials
- A correctly set up recurring meeting with a start and end time (this is
  important) - check [zoom guidelines] for more details 

Steps:
- Install splain: https://marketplace.zoom.us/apps/WPKzwuoLQDuj_gPs68AQxw
- Connect the zoom account
- Connect YouTube account
- Click the `manage` tab next to App Info and scroll to the config button, click
- The splain dashboard will display: make sure the box for make videos private is
checked so we can edit before it goes live.
- Test that it works
- Enjoy  

The following SIGs and groups are currently running splain.io:
- [SIG Auth](/sig-auth/README.md)
- [SIG Contributor Experience](/sig-contributor-experience/README.md)
- [SIG Docs](/sig-docs/README.md)
- [SIG Network](/sig-network/README.md)
- [Steering Committee](/committee-steering/governance/README.md)
- [WG Data Protection](/wg-data-protection/README.md)

The main zoom admin account which holds Meet Our Contributors and others (if
you log in to splain using this account, all of the other accounts will be
logged here)

TODO - look into splain.io's google drive to youtube pipeline. Also look into
using the gsuite contributors@ account to use the API for zoom cloud -> google
drive -> youtube.

### Descriptions & Playlists

Each video that comes into the queue needs to be added to a playlist, set to
public, and added context to the description.

Example description:  
Kubernetes Contributor Experience Special Interest Group Weekly Meeting. Check
here (link to sig list) for a complete list of SIGs and when they meet. Join us!

Please note the following items:
* The date must be in title and description
* The date format must be YYYYMMDD
* The section titled `Is this video made for kids?` should not be altered for search reasons

Below are a good and incorrect example:
* Incorrect description: `Data Protection WG Bi-Weekly Meeting for 2/26/2020`
* Correct description: `Data Protection WG Bi-Weekly Meeting for 20200226`

A short video tutorial of the editing of descriptions and playlists can be found here:
* https://youtu.be/IbZ2wnvu_Zs

### Thumbnails

TODO - someone help us with better thumbnails to lay over the videos!

### Streaming Events

YouTube admins with a system capable of streaming may be asked to stream public
Kubernetes Community events such as the weekly [Community Meeting],
[Office Hours], [Meet our Contributors], or other publicly streamed events. For
detailed information about streaming, see our [Streaming Config]

### Migrating Content

In certain cases, such as archiving an inactive SIG or Working Group it might be
useful to grab content from other channels. It is currently NOT POSSIBLE to move
content from one YouTube channel to another, so the content must be downloaded
and then reuploaded to the Kubernetes channel

1. Download [youtube-dl], which makes it easier to bulk download videos.
2. Download the channel or playlist with: `youtube-dl <url>` 
3. Clean up the filenames as they are used to generate new titles. Do this
   locally since it is easier than doing it per video in the YouTube web UI.
4. Create a new playlist for the content on the Kubernetes channel, set it to Private.
5. Upload the videos to the Kubernetes Channel.
6. Ensure titles and metadata are accurate, allow time to pass for YouTube to
   generate thumbnails and process the videos. 
7. Once videos are finalized, set the playlist to Public to publish them.


[coc]: /code-of-conduct.md
[Kubernetes YouTube Channel]: https://www.youtube.com/channel/UCZ2bu0qutTOM0tHYa_jkIwg
[collaboration]: https://support.google.com/youtube/answer/6109639
[loosen the permissions in your YouTube settings]: https://support.google.com/a/answer/6212415
[SIG Contributor Experience]: /sig-contributor-experience
[centralized list of administrators]: /communication/moderators.md
[YouTube admins]: /communication/moderators.md#YouTube-Channel
[trim]: https://support.google.com/youtube/answer/9057455?hl=en
[edit]: https://support.google.com/youtube/topic/9257530?hl=en&ref_topic=9257610
[Community Meeting]: /events/community-meeting.md
[Office Hours]: /events/office-hours.md
[Meet our Contributors]: /mentoring/programs/meet-our-contributors.md
[Streaming Config]: ./streaming-config.md
[Subprojects]: /governance.md#subprojects
[moderation guidelines]: /communication/moderation.md
[zoom guidelines]:/communication/zoom-guidelines.md
[youtube-dl]: https://ytdl-org.github.io/youtube-dl/index.html