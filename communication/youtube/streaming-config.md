# OBS Streaming Config

## Disclaimers: 
* These instructions won't work in OSX. The UX there for capturing audio is not great.
* Streaming requires a stable internet connection with at least ~10Mbps up depending on stream quality
* Some settings may be better than others. Everyones' computer is different. This is a solid baseline to start

## Walkthrough

(Link soon)

## Instructions

* Pre-installation
  * For non-YouTube Admins: Obtain a one-time streaming key from a YouTube Admin.
  * For YouTube Admins: See [streaming key instructions]

* Installation
  * To install OBS, visit their [installation wiki] and find your OS
  * Open OBS
  * Go through Wizard. It sets decent baselines. ([Wizard Screenshots])

* Video 
  * Under the Sources area, press the Plus sign, then Screen or Display capture
    * Name it what you want (eg. Left Monitor.)
    * Select which display you wish to capture (If you have one monitor, you'll have one option)
    * Press OK
    * You should now see your display mirrored within OBS.
  * You can create bumpers to put before/after a stream by creating a new scene.
    * Under the Scenes area, press the Plus sign. Name the new scene "Front Bumper"
    * Under the Sources area, press the Plus sign, then Image. Name the image "Front Bumper Image" and click "OK"
    * Click "Browse" and find an appropriate bumper. If you need an appropriate bumper, contact a YouTube Admin.
    * Click "OK"
    * Adjust the image to fill the screen by clicking and dragging the edges.
  * You can toggle between Scenes by clicking on their name. What you see in the preview is what will be streamed.
* Audio
  * OBS uses whatever output your computer uses by default.
    * Test audio by visiting a YouTube video or playing music and seeing the VU meter move
    * To change/set an interface besides the default:
      * Click the gear under the "Desktop Audio" VU meter and click "Properties"
      * Select a different audio interface, then click "OK"
    * Adjust the Desktop Audio by using the slider. A decent rule is to have music audio peaking in the middle of the yellow.
  * If you are streaming as well as participating in discussion, there's added complexity
    * To ensure your preferred microphone is selected by:
      * Click the gear under the "Mic/Aux" VU meter and click "Properties"
      * Select the correct device from the drop-down and click "OK"
    * Adjust the Mic Level so if you loudly say "WHAT" it doesn't go in the red.
    * While streaming, you will have to mute/unmute yourself both in Zoom as well as in OBS
      * Note: A headset or Mic with a hardware mute makes this a lot easier

* Other Settings
  * Click "Settings" on the bottom right. Everything below this will be within the "Settings" Window.
    * To change or update a YouTube streaming key
      * Click "Stream" on the left
      * Update the "Stream Key" field
      * Click "OK"

* Last steps
  * To test sound levels and capture quality, click "start recording" and play some music while talking. This will give you a good idea how things will sound on-stream.  
  * When starting a stream, start a minute or so early with a bumper up and both "Desktop" and "Mic/Aux" muted.
    * Under "Desktop Audio" as well as "Mic/Aux" you can toggle mute by clicking the "Speaker" icon.

That's it! You should be ready to stream content for the community.

[streaming key instructions]: ./obs/streaming-key.md
[installation wiki]:https://obsproject.com/wiki/install-instructions
[Wizard Screenshots]: ./obs/wizard-screenshots/