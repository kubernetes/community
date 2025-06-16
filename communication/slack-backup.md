# Backup private channels and DM's of Slack

To Backup DM's and private channels we currently recommend [slackdump](https://github.com/rusq/slackdump). 

It is an easy to use commandline tool to backup your Channels and DM's and store the data locally. No Online Service involved, so all of your data stays with you.

## Installation

Slackdump is supporting various platforms. You can find the latest release [here](https://github.com/rusq/slackdump/releases/tag/v3.1.4)

### MacOS

```
brew install slackdump
```

### Other OS

On other Operating Systems, please follow these steps:

1. Download the latest release for your operating system from the [releases] page.
1. Unpack the archive to any directory.
1. Run the `./slackdump` or `slackdump.exe` executable (see note below).
1. You know the drill:  use arrow keys to select the menu item, and Enter (or
   Return) to confirm.
1. Follow these [quickstart instructions][man-quickstart].

[releases]: https://github.com/rusq/slackdump/releases/

> [!NOTE] 
> On Windows and macOS you may be presented with "Unknown developer" window,
> this is fine.  Reason for this is that the executable hasn't been signed by
> the developer certificate.

  To work around this:

  - **on Windows**: click "more information", and press "Run
    Anyway" button.
  - **on macOS** 14 Sonoma and prior:  open the folder in Finder, hold Option
    and double click the executable, choose Run.
  - **on macOS** 15 Sequoia and later:  start the slackdump, OS will show the
    "Unknown developer" window, then go to System Preferences -> Security and
    Privacy -> General, and press "Open Anyway" button.


## Step by Step Walkthrough



Use the wizard to get started:

![](https://i.imgur.com/Spe3PwS.jpeg)

If you select export you will be asked for a login. 
![](https://i.imgur.com/WtngQG8.jpeg)


You can select any option you prefer. For some users the default browser Login does not work. We recommend the Cookie and Workspace Name Option.

![](https://i.imgur.com/77vMM7o.jpeg)

To retrieve the Cookie, Login into [https://kubernetes.slack.com](https://kubernetes.slack.com) with your credentials and retrieve the Coookie. It will start with the following formate: 

```
xoxd-
```

![](https://i.imgur.com/92nJ2e4.jpeg)

Enter as workspace Name "kubernetes" and then your Session Cookie and make the Checkmark to create credentials:

![](https://i.imgur.com/9jDVBQ5.jpeg)

Now you can dump your private Channels and DMs
The easiest way is to use ChannelId's

To retrieve a Chanel ID of a private Channel, select the channel and click on the 3 dots to choose "Open channel details". The following Popup will display the Channel ID. 

![](https://i.imgur.com/cxtSQOr.jpeg)

![](https://i.imgur.com/Xr9VTIA.jpeg)

For DM's you follow a similar way. Again the three dots and "Open conversation details". Afterwards you can again copy a ChanelID
![](https://i.imgur.com/sbz1M0G.jpeg)
![](https://i.imgur.com/LztpgNS.jpeg)

Now you can export the Data, by entering the ChanelIDs. 
We recommend the "standard" storage type option.

![](https://i.imgur.com/TchQXRA.jpeg)
Once you selected all teh options hit ESC to go back - then you can say start dump

And you will receive your export!:
![](https://i.imgur.com/roamx80.jpeg)



## View the export:

Once the workspace data is dumped, you can run built-in viewer:

```shell
slackdump view <zip or directory>
```

The built-in viewer supports all types of dumps:

1. Slackdump Archive format;
1. Standard and Mattermost Slack Export;
1. Dump mode files
  
The built-in viewer is experimental, any contributions to make it better looking are welcome.

Alternatively, you can use one of the following tools to preview the
export results:

- [SlackLogViewer] - a fast and powerful Slack Export viewer written in C++, works on Export files (images won't be displayed, unless you used an export token flag).
- [Slackdump2Html] - a great Python application that converts Slack Dump to a
  static browsable HTML.  It works on Dump mode files.
- [slack export viewer][slack-export-viewer] - Slack Export Viewer is a well known viewer for
  slack export files. Supports displaying files if saved in the "Standard" file mode.

[SlackLogViewer]: https://github.com/thayakawa-gh/SlackLogViewer/releases
[Slackdump2Html]: https://github.com/kununu/slackdump2html
[slack-export-viewer]: https://github.com/hfaran/slack-export-viewer






