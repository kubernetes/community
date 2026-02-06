---
title: "Zoom Guidelines"
description: |
  Policies, procedures and best practices for managing Zoom.
---

Zoom is the main video communication platform for Kubernetes. It is used for
running the [SIG/WG/Committee meetings],and many other Kubernetes online events.
Since the Zoom meetings are open to the general public, a Zoom host or co-host
has to moderate a meeting in all senses of the word, from starting and stopping
the meeting to acting on [Kubernetes code of conduct] issues.

These guidelines are meant as a tool to help Kubernetes members manage their
Zoom resources.

Check the main [moderation] page for more information on other tools
and general moderation guidelines.


## Code of conduct

The Kubernetes project adheres to the [Kubernetes Code of Conduct]
throughout all platforms and includes all communication mediums.

## Zoom license management

Zoom licenses are managed by the [CNCF Service Desk] through the
[Zoom Admins] listed in the  [centralized list of administrators].

### Obtaining a Zoom license

Ensure that all SIG/WG leads, chairs, and any other necessary trusted owners
have access to the `k-sig-<foo>-leads@googlegroups.com` account as described in
the [sig creation procedure]. Once done, contact one of the [Zoom Admins] to
obtain a Zoom license.

## Setting up your meeting and moderation

Do **not** share your Zoom link on social media. This will help curtail trolls
and others who would intentionally attempt to disrupt your Zoom call.

To create a meeting with **moderation** enabled, ensure the following:

-   Have the [latest version] of the Zoom client installed.
-   Be logged in as the leads account associated with the meeting **OR** use the
    [host key] to "claim host".
-   Configure a meeting setup through the "Meeting" menu in the leads Zoom
    account. **NOTE:** Do **NOT** use the "Personal Meeting ID". This will
    create an "ad-hoc" meeting that is time-bounded and without moderation
    capability.
-   Set the password to the meeting to "77777"

After the meeting has started:

-   Assign a co-host to help with moderation. It should never be your note taker
    unless it's a very small group.
-   Turn **off** screen sharing for everyone and indicate "only host". If you
    have others that need to share their screen, the host can enable that on
    the fly. (via the `^` menu next to **Share Screen**)

### Moderation

If you're dealing with a troll or bad actor:

-   Put the troll or bad actor on **hold**. The participant will be put into a
    [waiting room] and will not be able to participate in the call until the
    host removes the hold.
    -   **NOTE:** Depending on your client version this will be called "**Put in
        Waiting Room**" instead of on **hold**.
-   Remove the participant. Please be cautious when testing or using this
    feature, as it is **permanent**. They will never be able to come back into
    that meeting ID on that particular device. Do **not** joke around with
    this feature; it's better to put the attendee on "hold" first and then
    remove.
-   After an action has been taken, use the **lock meeting** feature so that no
    one else can come into the meeting. If that fails, end the call
    immediately, and contact the [Zoom Admins] to report the issue.

**NOTE:** You can find these actions when clicking on the **more** or **"..."**
options after scrolling over the participants name/information.

Hosts **must** be comfortable with how to use these moderation tools and the
Zoom settings in general. Make sure whoever is running your meeting is equipped
with the right knowledge and skills. If you have any questions or concerns,
reach out to the [Zoom Admins] and they will be able to provide further
guidance and training.

#### Related moderation documentation

-   Zoom has [documentation on how to use their moderation tools].
-   Members of the _leads@_ group have access to an extensive 
    [best practices doc] with screenshots going over the community Zoom best
    practices.

### Escalating and Reporting a Problem

Issues that cannot be handled via normal moderation, or with the assistance of
the [Zoom Admins] should be escalated to the Kubernetes 
[Code of Conduct Committee] at conduct@kubernetes.io.

To contact the admin group in Slack, ping `@zoom-admins` in the `#sig-contribex`
Slack channel.

## Meeting recordings

Chairs and TLs are responsible for posting all update meetings to their playlist
on YouTube. [Please follow this guideline for more details].

If a violation has been addressed by a host and it has been recorded by Zoom,
the video should be edited before being posted on the [Kubernetes channel].

Contact [SIG Contributor Experience] if you need help to edit a video
before posting it to the public.

## Screen sharing guidelines and recommendations

Zoom has [documentation on how to use their screen sharing feature].

Recommendations:

-   Turn off notification to prevent any interference.
-   Close all sensitive documents and unrelated programs before sharing the
    screen. Email notifications are distracting!
-   Test your presentation beforehand to make sure everything goes smoothly.
-   Keep your computer background desktop clean. Make sure there are no offensive
    or distracting visuals.

## Audio/Video quality recommendations

While video conferencing has been a real boon to productivity there are still
[lots of things that can go wrong] during a conference video call.

There are some things that are just plain out of your control, but there are
some things that you can control. Here are some tips if you're just getting into
remote meetings. Keep in mind that sometimes things just break. These are not
hard rules, more of a set of loose guidelines on how to tip the odds in your
favor.

### Recommended hardware to have

-   **A dedicated microphone** - This is the number one upgrade you can do.
    Sound is one of those things that can immediately change the quality of
    your call. If you plan on being here for the long haul, something like a
    [Blue Yeti] will work great due to the simplicity of using USB
    audio and having a hardware mute button. Consider a [pop filter]
    as well if necessary.
-   **A Video Camera** - A bad image can be worked around if the audio is good.
    Certain models have noise canceling dual-microphones, which are a great
    backup for a dedicated microphone or if you are traveling.
-   **A decent set of headphones** - These cut down on the audio feedback when
    in larger meetings.

What about an integrated headset and microphone? This totally depends on the
type. We recommend testing it with a friend or asking around for recommendations
for which models work best.

### Hardware we don't recommend

-   **Earbuds** - These are not ideal, and while they might sound fine to you,
    when 50 people are on a call the ambient noise adds up. Some people join
    with earbuds and it sounds excellent, others join and it sounds
    terrible. Practicing with someone ahead of time can help you determine how
    well your earbuds work.

### Pro-tips

-   [Join on muted audio and video] in order to prevent noise to those
    already in a call.
-   If you don't have anything to say at that moment, **MUTE**. This is a common
    problem. You can help out a teammate by mentioning it on Zoom chat or
    asking them to mute on the call itself. The meeting co-host can help with
    muting noisy attendees before it becomes too disruptive. Don't feel bad if
    this happens to you, it's a common occurrence.
-   Try to find a quiet meeting place to join from; some coworking spaces and
    coffee shops have a ton of ambient noise that won't be obvious to you but
    will be to other people in the meeting. When presenting to large groups
    consider delegating to another person who is in a quieter environment.
-   Using your computer's built-in microphone and speakers might work in a
    pinch, but in general won't work as well as a dedicated
    headset/microphone.
-   Consider using visual signals to agree to points so that you don't have to
    mute/unmute often during a call. This can be an especially useful
    technique when people are asking for lazy consensus. A simple thumbs up
    can go a long way!
-   It is common for people to step on each other when there's an audio delay,
    and both parties are trying to communicate something. Don't worry, just
    remember to try and pause before speaking, or consider raising your hand
    (if your video is on) to help the host determine who should speak first.

Thanks for making Kubernetes meetings work great!

[New Contributor Orientation]: /mentoring/new-contributor-orientation/README.md
[SIG/WG/Committee meetings]: /sig-list.md
[Kubernetes code of conduct]: /code-of-conduct.md
[moderation]: ./moderation.md
[CNCF Service Desk]: https://github.com/cncf/servicedesk
[Zoom Admins]: ./moderators.md#zoom
[centralized list of administrators]: ./moderators.md
[sig creation procedure]: /sig-wg-lifecycle.md#communicate
[latest version]: https://zoom.us/download
[host key]: https://support.zoom.us/hc/en-us/articles/205172555-Host-Key
[waiting room]: https://support.zoom.us/hc/en-us/articles/115000332726-Waiting-Room
[documentation on how to use their moderation tools]: https://support.zoom.us/hc/en-us/articles/201362603-Host-Controls-in-a-Meeting
[best practices doc]: https://docs.google.com/document/d/1fudC_diqhN2TdclGKnQ4Omu4mwom83kYbZ5uzVRI07w/edit?usp=sharing
[Code of Conduct Committee]: /committee-code-of-conduct/README.md
[Please follow this guideline for more details]: ./youtube/youtube-guidelines.md
[Kubernetes channel]: https://www.youtube.com/c/kubernetescommunity
[SIG Contributor Experience]: /sig-contributor-experience
[documentation on how to use their screen sharing feature]: https://support.zoom.us/hc/en-us/articles/201362153-How-Do-I-Share-My-Screen
[lots of things that can go wrong]: https://www.youtube.com/watch?v=JMOOG7rWTPg
[Blue Yeti]: https://www.bluedesigns.com/products/yeti/
[pop filter]: https://en.wikipedia.org/wiki/Pop_filter
[Join on muted audio and video]: https://support.zoom.us/hc/en-us/articles/203024649-Video-Or-Microphone-Off-By-Attendee
