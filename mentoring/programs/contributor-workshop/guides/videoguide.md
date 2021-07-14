<!-- omit in toc -->
# Video Recording Guide for Contributor Workshop Segments

This guide covers, how to structure, record, upload and update video guides associated with workshop segments.

- [Format / Video Style](#format--video-style)
- [Encoding Guidelines](#encoding-guidelines)
  - [Zoom Recording](#zoom-recording)
  - [OBS](#obs)
  - [Streamyard](#streamyard)
  - [Adobe Media Encoder](#adobe-media-encoder)
- [Recording Guides](#recording-guides)
  - [Record Using Zoom](#record-using-zoom)
    - [Install Windowed extension (optional)](#install-windowed-extension-optional)
    - [Start the meeting](#start-the-meeting)
    - [Setting up slides for recording](#setting-up-slides-for-recording)
      - [Without browser extension](#without-browser-extension)
      - [With browser extension](#with-browser-extension)
        - [Note : Exiting the windowed mode](#note--exiting-the-windowed-mode)
    - [Adding your facecam](#adding-your-facecam)
    - [Unmute your audio](#unmute-your-audio)
    - [Recording in Zoom](#recording-in-zoom)
      - [Note : Do a test recording first](#note--do-a-test-recording-first)
      - [How to record](#how-to-record)
  - [Recording in OBS](#recording-in-obs)
  - [Recording in Streamyard](#recording-in-streamyard)
- [Updating / Adding video to a segment](#updating--adding-video-to-a-segment)
  - [Rename your recording](#rename-your-recording)
  - [Upload your video to Google Drive](#upload-your-video-to-google-drive)
  - [Updating Youtube](#updating-youtube)
- [Update the Segment page with video](#update-the-segment-page-with-video)
- [Best Practices & Recommendations](#best-practices--recommendations)
  - [Video](#video)
    - [Camera Type](#camera-type)
  - [Audio](#audio)
    - [Microphone Types](#microphone-types)
    - [Sound Environment](#sound-environment)
      - [Background Noise](#background-noise)
        - [Removing Background noise in Zoom](#removing-background-noise-in-zoom)
        - [Removing Background noise with Krisp](#removing-background-noise-with-krisp)
        - [Removing Background noise with NVIDIA Broadcast or RTX Voice](#removing-background-noise-with-nvidia-broadcast-or-rtx-voice)
      - [Other Tips](#other-tips)
  - [Screensharing](#screensharing)
  - [Hardware](#hardware)
    - [Headset Microphones](#headset-microphones)
    - [External Microphones](#external-microphones)
    - [Webcams](#webcams)
    - [Presentation remotes](#presentation-remotes)

## Format / Video Style

You can use any format/style as long as the same information in the guide is captured in the video. Slides, A screenshare, green screen yourself into the repo even! Be creative.

## Encoding Guidelines

### Zoom Recording

Handled by zoom, Can't change these settings.

### OBS

![OBS Settings](obssettings.png)

### Streamyard

Handled by Streamyard, Can't change these settings.

### Adobe Media Encoder

use Youtube 1080p Full HD

## Recording Guides

Guides on how to record a presentation style video (slides/screen,facecam)

### Record Using Zoom

#### Install Windowed extension (optional)

This makes it easier for you to share the slides to zoom. Allows you to pop out the fullscreen view to a single window.

[GitHub Link](https://github.com/dralletje/Windowed)

#### Start the meeting

Start a meeting in zoom (you can start and record a meeting with a single participant)

#### Setting up slides for recording

Open your slides in Presenter View

![Presenter View](presenterview.png)

It should open a presenter view, with a timer, slide notes, and change the tab from editor to a slideshow presentation

![Presenter View and Slides](presenterviewandslides.png)

You will use the presenter view to control the slides, and the slide view will be cast to the zoom call.

##### Without browser extension

Share the window which is displaying the slides view.

![Slides with your tabs](slideswithtabs.png)

##### With browser extension

When moving your mouse on the slide view a small dialog with the slide number, and triple dots will appear. Select the triple dots and enter full screen

![Enter Full Screen on Presentation](presenterfullscreen.png)

After selecting the fullscreen option, you will be presented further options on how to display the slides.

Select the windowed option

![Windowed Fullscreen](presenterwindowed.png)

A new window with only the slides will appear, Share this window in zoom.

![Slides with no tabs](slidesnotabs.png)

###### Note : Exiting the windowed mode

To exit out of this view, and go back the editor, select the window with the slide view back into focus, move your mouse until the triple dots mentioned earlier appear. Select the triple dots again and ``Exit Full Screen``

![Exit Full Screen](presenterexitfullscreen.png)

#### Adding your facecam

If you haven't done so already, start your video in Zoom.

#### Unmute your audio

If you haven't done so already, unmute your audio in Zoom.

#### Recording in Zoom

Now that you've got your slides, video and audio setup, you're ready to record!

##### Note : Do a test recording first

It's good practice to create a test recording of your whole setup in action.
If you change shares between windows, check that changing the shares appears how you want it to in the video.

##### How to record

Select the ``Record`` button in zoom.

If you have the option between saving locally, and in the cloud, select locally.

Once you have finished recording, end the meeting.

A popup should appear with progress on converting the recording.
Once the recording is converted, a file browser window with the recording will be opened. Rename the recording from ``zoom_0.mp4`` to the format described in [Rename your recording](#rename-your-recording)

### Recording in OBS

Coming soon

### Recording in Streamyard

Coming soon

## Updating / Adding video to a segment

### Rename your recording

Rename your recording to ``YYYY-MM-DD-<Segement Name>-<gh-handle>.mp4``

Where

- ``YYYY-MM-DD`` being date, is the date video was recorded
- ``<Segement Name>`` being the name of the markdown file associated with that segment. see [CONTRIBUTING.md](https://github.com/kubernetes/contributor-site/tree/workshop/content/en/workshop/CONTRIBUTING.md#list-of-segments)
- ``<gh-handle>`` being the github handle of the person who recorded the video, if there are multiple presenters just add another dash and the handle. ``<gh-handle>-<gh-handle>``

For example

``2021-05-10-why-contribute-alisondy.mp4``

``2021-05-10-why-contribute-alisondy-mrbobbytables.mp4``

### Upload your video to Google Drive

Upload your recording to the Google Drive folder provided for the Workshop Videos.

### Updating Youtube

Once you have uploaded the video to drive, contact one of the [youtube admins](/communication/moderators.md#youtube-channel) to upload it to the Kubernetes YouTube channel.

> Youtube Admins see [Video Upload Guide](videouploadguide.md) for details about uploading the segment to youtube

## Update the Segment page with video

If adding to a segment page where video doesn't exist, Paste this at the top, after the frontmatter. Fill in the details with the youtube id of the video you just uploaded.

If adding to a segment page where a video already exists, update the youtube id in the shortcode, with the video id of your newly uploaded video.

```txt
## Video Walkthrough

{{< youtube id="<id of your video>" title="Contributor Workshop : <Segment Name" >}}
```

## Best Practices & Recommendations

### Video

#### Camera Type

Use an external camera or webcam to capture your video.

Built in webcams restrict the angle of you webcam, to the viewing angle of your display, and often provide a low quality image.

- Use an external webcam. 1080p preferred or 720p at a minimum. Typically, either will be better quality than the built-in camera on your laptop.
- Lighting - Adjusting your lighting may help you get a better camera feed in less than optimal lighting conditions.
  - Optimize the ambient lighting within your space first.
    - There should be no window or strong light source behind you that darkens your face
    - There should be no strong light source from one side that casts a shadow on the opposite side of your face.
    - Good ambient lighting from your ceiling lights or windows may be enough as long as they are in front of you.
    - Diffused lighting is better than harsh, direct lighting.
- Background – keep it clean and uncluttered. It’s great to show off your personality/style in your home, just be sure to keep everything appropriate for a general audience and compliant with the Code of Conduct. A plain wall is also just fine, no need to overdo it.
- Camera Position - as close as possible at your eye level, not pointing up towards your nose or the ceiling, or down towards the floor. It can be as simple as raising your laptop via a few books. Or if you have an external webcam, position it on the top edge or close to the top of the screen from which you are advancing your slides/notes.
- Keep eye contact with the camera just like you would look at your audience during an in-person presentation. It's okay to look away for brief moments, just make sure you look towards the camera for the majority of the presentation.
- Video with more lighting tips [https://youtu.be/rGcAM1CNEU0](https://youtu.be/rGcAM1CNEU0)

### Audio

#### Microphone Types

**Best:** External microphone (condenser recommended over dynamic*) via a USB connection or DAC. This type of microphone will pick up audio from a distance as well as up close really well. This means you can position it a fair distance away from yourself when recording. If using one of these microphones, try to position it out of frame from your video.

**Good:** Headset Microphone, this could be a microphone located on the wires like on mobile earphones, the mic on your wireless earphones as well as the usual fold down microphone on a headset. All of these assist with isolating microphone pick-up from ambient noises within your recording/streaming space. Wired headsets will give you better quality audio than wireless.

**Acceptable:** External microphone built into your webcam. Often better quality than your built-in laptop microphone.

**Worst:** Your built-in laptop microphone. These are generally low-quality microphones, and they have little or no ability to filter out background noise.

#### Sound Environment

##### Background Noise

It’s best to be in a controlled environment with little to no background noise.
If this is not possible you can use your recording software and other software to remove background noise.
Another option to filtering out background noise or being in a controlled environment is using push to talk, or unmuting your mic whenever you need to speak.

###### Removing Background noise in Zoom

In zoom settings under ``Audio``
Change the Suppress background noise setting to ``High``

###### Removing Background noise with Krisp

Krisp provides AI powered background noise cancellation, see [krisp.ai](https://krisp.ai/) for more information.

###### Removing Background noise with NVIDIA Broadcast or RTX Voice

If you're using Windows, and are using certain models of NVIDIA GPU, you can install either NVIDIA Broadcast or RTX Voice to filter out background noise. See [RTX Voice setup guide](https://www.nvidia.com/en-us/geforce/guides/nvidia-rtx-voice-setup-guide/) and [Broadcast app setup guide](https://www.nvidia.com/en-gb/geforce/guides/broadcast-app-setup-guide/) for more information.

##### Other Tips

- Mute all notifications and programs not being recorded
- Disable any unnecessary audio devices

### Screensharing

- Increase the font size when sharing terminals, text editors
- Ensure you have all the windows you need opened before recording.
- Export any sensitive variables needed in your terminals before starting your presentation
- If using a passworded ssh key, Start a ssh-agent in the terminal beforehand.

### Hardware

Here are some equipment recommendations from the CNCF:

#### Headset Microphones

Logitech / Plantronics – i.e. Logitech H111 / Logitech H390 / Plantronics C225

#### External Microphones

Blue / Samson – i.e. Blue Snowball / Blue Yeti / Samson Meteor / Samson G-Track

#### Webcams

Good value – Logitech C920

More info – [https://www.tomsguide.com/news/where-to-buy-webcams-now-these-retailers-have-stock](https://www.tomsguide.com/news/where-to-buy-webcams-now-these-retailers-have-stock)

#### Presentation remotes

Logitech / Canon / Kensington / SMK

Tested and recommended for PC or Mac – Logitech R400, R800, Spotlight / Canon PR500, PR100
