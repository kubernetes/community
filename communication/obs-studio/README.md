# OBS Config for Kubernetes

This is the [Open Broadcaster Software](https://obsproject.com/) configuration used to stream meetings and events to YouTube.

## Warning

These files containe the YouTube streaming key in `service.json`, make sure you're not pushing that key to GitHub. If you do accidentally, you will need to [generate a new one](https://support.google.com/youtube/answer/2853818) and ensure all streamers update their OBS to use the new key. 

This key is what allows a streamer to stream as "Kubernetes", so guard it carefully!

## File locations

This config should be copied into: 

Windows: `%APPDATA%\obs-studio`
macOS: `~/Library/Application Support/obs-studio`
Linux: `~/.config/obs-studio` or `~/.var/app/com.obsproject.Studio/config/obs-studio`
