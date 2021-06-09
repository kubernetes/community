# Segment Upload Guide for Youtube Admins

## Upload the new video

In youtube, upload the same video and title it as follows

```txt
Contributor Workshop : <Segment Name> with <Presenter Name> | <Month> <Year>
```

For example

```txt
Contributor Workshop : Why Contribute? with @alisondy | May 2021
```

## Adding the thumbnail image

```markdown
TODO : Instructions on how to generate a thumbnail image, similar to https://gist.github.com/alisondy/e4ff3b483a32330ade0c880b4a4a10b1
```

Set the thumbnail to the image generated in the previous step

## Setting the description

Use the following template for setting the video description

```txt
This is the video guide for the Contributor Workshop segment <Segment Name>

Link to this segments page:
https://kubernetes.dev/docs/workshop/<page-name>

Link to the workshop:
https://kubernetes.dev/docs/workshop/
```

## Update the previous video for the segment if it exists

If you're updating an already existing video, follow the next steps to update the following on the previous video

- title
- thumbnail
- description
- comments section

## Update the title

Update the title on the previous video to match the following format

```txt
Contributor Workshop : <Segment Name> | <Month> <Year> | 🚨🚨 DEPRECATED 🚨🚨
```

For example

```txt
Contributor Workshop : Why Contribute? with @alisondy | May 2021 | 🚨🚨 DEPRECATED 🚨🚨
```

## Update the thumbnail image

Update the Thumbnail image on the previous video to

insert-deprecation-thumbnail-here.jpeg

```txt


THIS VIDEO IS DEPRECATED!
See the comments & description for details on where to find the new video guide!


```

## Update the Description

On the previous video, add the following to the start of the description

```txt
🚨🚨 DEPRECATED , Please see <link to new video> 🚨🚨
```

## Update the Comments

On the previous video, add the following as a comment as Kubernetes YT & pin to the top of the comments

```txt
🚨🚨 DEPRECATED , Please see <link to new video> 🚨🚨
```

Close the comment section of the deprecated video

## Update the Playlist

Update the youtube playlist with the segment recording.
