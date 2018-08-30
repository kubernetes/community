# Zoom Guidelines

Zoom is the main video communication platform for Kubernetes. 
It is used for running the [community meeting](https://github.com/kubernetes/community/blob/master/events/community-meeting.md) and SIG meetings. 
Since the Zoom meetings are open to the general public, a Zoom host has to moderate a meeting if a person is in violation of the code of conduct. 

These guidelines are meant as a tool to help Kubernetes members manage their Zoom resources. 
Check the main [moderation](./moderation.md) page for more information on other tools and general moderation guidelines.

## Code of Conduct
Kubernetes adheres to Cloud Native Compute Foundation's [Code of Conduct](https://github.com/cncf/foundation/blob/master/code-of-conduct.md) throughout the project, and includes all communication mediums.

## Moderation

Zoom has documentation on how to use their moderation tools: 

- https://support.zoom.us/hc/en-us/articles/201362603-Host-Controls-in-a-Meeting

Check the "Screen Share Controls" (via the ^ next to Share Screen): Select who can share in your meeting and if you want only the host or any participant to be able to start a new share when someone is sharing.  

You can also put an attendee on hold. This allows the host(s) to put attendee on hold to temporarily remove an attendee from the meeting. 

Unfortunately, Zoom doesn't have the ability to ban or block people from joining - especially if they have the invitation to that channel and the meeting id is publicly known.

It is required that a host be comfortable with how to use these moderation tools. It is strongly encouraged that at least two people in a given SIG are comfortable with the moderation tools. 

## Meeting Archive Videos

If a violation has been addressed by a host and it has been recorded by Zoom, the video should be edited before being posted on the [Kubernetes channel](https://www.youtube.com/c/kubernetescommunity).

Contact [SIG Contributor Experience](https://github.com/kubernetes/community/tree/master/sig-contributor-experience) if you need help to edit a video before posting it to the public. 

## Admins

- Check the [centralized list of administrators](./moderators.md) for contact information.

Each SIG should have at least one person with a paid Zoom account. 
See the [SIG Creation procedure](https://github.com/kubernetes/community/blob/master/sig-governance.md#sig-creation-procedure) document on how to set up an initial account. 

The Zoom licenses are managed by the [CNCF Service Desk](https://github.com/cncf/servicedesk). 

## Escalating and/Reporting a Problem

Issues that cannot be handle via normal moderation can be escalated to the [Kubernetes steering committee](https://github.com/kubernetes/steering). 

## Audio/Video Quality Recommendations

While video conferencing has been a real boon to productivity there are still [lots of things that can go wrong](https://www.youtube.com/watch?v=JMOOG7rWTPg) during a conference video call.

There are some things that are just plain out of your control, but there are some things that you can control. 
Here are some tips if you're just getting into remote meetings. 
Keep in mind that sometimes things just break and sometimes it's just plain bad luck, so these aren't hard rules, more of a set of loose guidelines on how to tip the odds in your favor.  

### Recommended Hardware to Have

- A dedicated microphone - This is the number one upgrade you can do. Sound is one of those things that can immediately change the quality of your call. If you plan on being here for the long haul something like a  [Blue Yeti](https://www.bluedesigns.com/products/yeti/) will work great due to the simplicity of using USB audio and having a hardware mute button. Consider a [pop filter](https://en.wikipedia.org/wiki/Pop_filter) as well if necessary.
- A Video Camera - A bad image can be worked around if the audio is good. Certain models have noise cancelling dual-microphones, which are a great backup for a dedicated microphone or if you are travelling. 
- A decent set of headphones - Personal preference, these cut down on the audio feedback when in larger meetings. 

What about an integrated headset and microphone? This totally depends on the type. We recommend testing it with a friend or asking around for recommendations for which models work best. 

### Hardware we don't Recommend

- Earbuds. Generally speaking they are not ideal, and while they might sound fine to you when 50 people are on a call the ambient noise adds up. Some people join with earbuds and it sounds excellent, some people join and it sounds terrible. Practicing with someone ahead of time can help you determine how well your earbuds work.

### Pro-tips

- [Join on muted audio and video](https://support.zoom.us/hc/en-us/articles/203024649-Video-Or-Microphone-Off-By-Attendee) in order to prevent noise to those already in a call.
- If you don't have anything to say at that moment, MUTE. This is a common problem, you can help out a teammate by mentioning it on Zoom chat or asking them to mute on the call itself. Hopefully the meeting co-host can help mute before this is too disruptive. Don't feel bad if this happens to you, it's a common occurrence.
- Try to find a quiet meeting place to join from; some coworking spaces and coffee shops have a ton of ambient noise that won't be obvious to you but will be to other people in the meeting. When presenting to large groups consider delegating to another person who is in a quieter environment.
- Using your computer's built in microphone and speakers might work in a pinch, but in general won't work as well as a dedicated headset/microphone.
- Consider using visual signals to agree to points so that you don't have to mute/unmute often during a call. This can be an especially useful technique when people are asking for lazy consensus. A simple thumbs up can go a long ways! 
- It is common for people to step on each other when there's an audio delay, and both parties are trying to communicate something, so don't sweat it, just remember to try and pause before speaking, or consider raising your hand (if your video is on) to help the host determine who should speak first.

Thanks for making Kubernetes meetings work great!