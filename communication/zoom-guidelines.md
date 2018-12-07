# Zoom Guidelines

Zoom is the main video communication platform for Kubernetes.
It is used for running the [community meeting](/events/community-meeting.md), [SIG/WG meetings](/sig-list.md), [Office Hours](/events/office-hours.md), [Meet Our Contributors](/mentoring/meet-our-contributors.md) and many other Kubernetes online events.
Since the Zoom meetings are open to the general public, a Zoom host or cohost has to moderate a meeting in all senses of the word from starting and stopping the meeting to acting on code of conduct issues.

These guidelines are meant as a tool to help Kubernetes members manage their Zoom resources.
Check the main [moderation](./moderation.md) page for more information on other tools and general moderation guidelines.

## Current State

Zoom licenses are managed by the [CNCF Service Desk](https://github.com/cncf/servicedesk) through the Zoom Admins listed below. At the time of this update, we have 41 paid pro user licenses with 38 accounted for.

## Code of Conduct

Kubernetes adheres to Cloud Native Compute Foundation's [Code of Conduct](https://github.com/cncf/foundation/blob/master/code-of-conduct.md) throughout the project, and includes all communication mediums.

## Obtaining a Zoom License

Each SIG should have a paid Zoom account that all leads/chairs/necessary trusted owners have access to through their k-sig-foo-leads@googlegroups.com account
See the [SIG Creation procedure](/sig-governance.md#sig-creation-procedure) document on how to set up an initial account.

## Setting Up Your Meeting and Moderation

Do not share your zoom link on social media.

Moderation will not be available if you are not following this list:

- [latest version](https://zoom.us/download)
- logged in as the leads of that meeting OR have the host key. (Example: you need to use the leads account for sig arch if you are running that meeting or have their meeting key with join before host enabled)
- using a meeting that was set up through the "Meeting" tab in the zoom account and NOT the personal meeting ID

After the meeting has started:

- Assign a cohost to help with moderation. It should never be your notetaker unless it's a very small group.
- Turn off screen sharing for everyone and indicate only host. If you have others that need to share their screen, the host can enable that on the fly. (via the ^ next to Share Screen)

If you're dealing with a troll or bad actor:

- You can put an attendee on hold. The participant will be put into a 'waiting room' and not have ability to chat or discuss written or verbally until the host undoes the hold.
- Remove the participant. Please be cautious when testing or using this feature, as it is permanent. They will never be able to come back into that meeting ID on that device. Do not joke around with this feature; it's better to use the hold first and then remove.
- After an action has been taken, use the 'lock meeting' feature so that no others come into the meeting and you can resume. If that fails, end the call immediately. Contact Zoom Admins after the meeting to report.

You can find these actions when clicking on the 'more' or '...' options after scrolling over the participants name/information.

It is required that a host be comfortable with how to use these moderation tools and the zoom settings in general. Make sure whoever is running your meeting is equipped with the right knowledge and skills.

### Other Related Documentation

Zoom has documentation on how to use their moderation tools:

- https://support.zoom.us/hc/en-us/articles/201362603-Host-Controls-in-a-Meeting

We created an extensive [best practices doc](https://docs.google.com/document/d/1fudC_diqhN2TdclGKnQ4Omu4mwom83kYbZ5uzVRI07w/edit?usp=sharing) with screenshots. Those who belong to kubernetes-sig-leads@ have access.

## Meeting Archive Videos

If a violation has been addressed by a host and it has been recorded by Zoom, the video should be edited before being posted on the [Kubernetes channel](https://www.youtube.com/c/kubernetescommunity).

Contact [SIG Contributor Experience](/sig-contributor-experience) if you need help to edit a video before posting it to the public.

Chairs and TLs are responsible for posting all update meetings to their playlist on YouTube. [Please follow this guideline for more details.](K8sYoutubeCollaboration.md)

## Zoom Admins

Check the [centralized list of administrators](./moderators.md) for contact information.

### Escalating and/Reporting a Problem

Issues that cannot be handle via normal moderation with the Zoom Admins above and there has been a clear code of conduct violation, please escalate to the Kubernetes Code of Conduct Committee at conduct@kubernetes.io.

## Screen sharing guidelines and recommendations

Zoom has a documentation on how to use their screen sharing feature:

- https://support.zoom.us/hc/en-us/articles/201362153-How-Do-I-Share-My-Screen-

Recommendations:

- Turn off notification to prevent any interference.
- Close all sensitive documents and unrelated programs before sharing the screen eg. Emails.
- Test your presentation before hand to make sure everything goes smoothly.
- Keep your desktop clean. Make sure there is no offensive or/and distracting background.

## Audio/Video Quality Recommendations

While video conferencing has been a real boon to productivity there are still [lots of things that can go wrong](https://www.youtube.com/watch?v=JMOOG7rWTPg) during a conference video call.

There are some things that are just plain out of your control, but there are some things that you can control.
Here are some tips if you're just getting into remote meetings.
Keep in mind that sometimes things just break and sometimes it's just plain bad luck, so these aren't hard rules, more of a set of loose guidelines on how to tip the odds in your favor.

### Recommended Hardware to Have

- A dedicated microphone - This is the number one upgrade you can do. Sound is one of those things that can immediately change the quality of your call. If you plan on being here for the long haul something like a [Blue Yeti](https://www.bluedesigns.com/products/yeti/) will work great due to the simplicity of using USB audio and having a hardware mute button. Consider a [pop filter](https://en.wikipedia.org/wiki/Pop_filter) as well if necessary.
- A Video Camera - A bad image can be worked around if the audio is good. Certain models have noise cancelling dual-microphones, which are a great backup for a dedicated microphone or if you are travelling.
- A decent set of headphones - Personal preference, these cut down on the audio feedback when in larger meetings.

What about an integrated headset and microphone? This totally depends on the type. We recommend testing it with a friend or asking around for recommendations for which models work best.

### Hardware we don't recommend

- Earbuds. Generally speaking they are not ideal, and while they might sound fine to you when 50 people are on a call the ambient noise adds up. Some people join with earbuds and it sounds excellent, some people join and it sounds terrible. Practicing with someone ahead of time can help you determine how well your earbuds work.

### Pro-tips

- [Join on muted audio and video](https://support.zoom.us/hc/en-us/articles/203024649-Video-Or-Microphone-Off-By-Attendee) in order to prevent noise to those already in a call.
- If you don't have anything to say at that moment, MUTE. This is a common problem, you can help out a teammate by mentioning it on Zoom chat or asking them to mute on the call itself. Hopefully the meeting co-host can help mute before this is too disruptive. Don't feel bad if this happens to you, it's a common occurrence.
- Try to find a quiet meeting place to join from; some coworking spaces and coffee shops have a ton of ambient noise that won't be obvious to you but will be to other people in the meeting. When presenting to large groups consider delegating to another person who is in a quieter environment.
- Using your computer's built in microphone and speakers might work in a pinch, but in general won't work as well as a dedicated headset/microphone.
- Consider using visual signals to agree to points so that you don't have to mute/unmute often during a call. This can be an especially useful technique when people are asking for lazy consensus. A simple thumbs up can go a long ways!
- It is common for people to step on each other when there's an audio delay, and both parties are trying to communicate something, so don't sweat it, just remember to try and pause before speaking, or consider raising your hand (if your video is on) to help the host determine who should speak first.

Thanks for making Kubernetes meetings work great!
