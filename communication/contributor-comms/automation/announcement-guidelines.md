# K8s Announcement Bot User guide

The source code for this bot is present in https://github.com/kubernetes-sigs/slack-infra/tree/main/slack-post-message and it is deployed using https://github.com/kubernetes/k8s.io/tree/main/apps/slack-infra

## Prerequisites

### Permissions

Reach out to the 'sig-contribex-comms' channel in kubernetes slack for permissions or for posting something on your behalf.

### Adding bot to a channel

In order for the bot to send message to a channel, it needs to be part of the channel first and before you add it, you need to join the channel. Although we have it added in most of the channels, it still might be required to be added to a new channel. Follow the below steps for doing this


1. Go to the 'More' option in the top left side and then go to 'Apps'
![](https://i.imgur.com/Dk26HwP.png)


2. That will take you to the apps page & choose 'Announcement bot' in this page
![](https://i.imgur.com/x3x3TQ6.png)


3. Now, slack will take you to a DM page with the bot. Click on the title of the bot at the top of this page
![](https://i.imgur.com/x7umv6O.png)


4. In the pop-up click on 'Add this app to a channel' & enter the channel name. This will have add the bot to the specified channel
![](https://i.imgur.com/It5K1sH.png)



## Pushing announcements through the bot

1. From any page on slack, find the '+' button called as "Attachments & shortcuts" in the text box
![](https://i.imgur.com/njN9Xqq.png)

2. In the options displayed, click on 'Browse all shortcuts'
![](https://i.imgur.com/pmVMWzS.png)

3. In the following page choose 'Announcement bot'
![](https://i.imgur.com/1eQVBWk.png)

4. In the options displayed for 'Announcement bot' click on 'Post message' which open the message form
![](https://i.imgur.com/M3mI3Jg.png)

5. In the form, choose the channel(s), write the message that needs to be sent and click on 'submit' buttom.
The keyboard here might not support slack emoji search (with `:`) The workaround for this is to copy emojis 
![](https://i.imgur.com/FlZdGgi.png)
