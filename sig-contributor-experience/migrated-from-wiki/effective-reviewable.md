*Or, one weird trick to make Reviewable awesome*

The Kubernetes team is still new to _Reviewable_. As you discover new cool features and workflows, add them here. Once we have built up a good number of tricks we can reorganize this list.

- Hold off on publishing comments (using the "Publish" button) until you have completed your review. [(source)](@pwittrock)
- When leaving comments, select a "disposition" (the button with your profile picture) to indicate whether the comment requires resolution, or is just advisory and hence requires no response. [(source)](@pwittrock)
- Change a comment's "disposition" to "close" those to which the author didn't respond explicitly but did address with satisfactory changes to the code. Otherwise, the comment hangs out there awaiting a response; in contrast to GitHub's review system, _Reviewable_ doesn't consider a change to the target line to be a sufficient indicator of resolution or obsolescence, which is a safer design. Use the <kbd>y</kbd> to acknowledge the current comment, which indicates that no further response is necessary.
- To "collapse" a whole file in the multi-file view, click the rightmost value in the revision range control. This is effectively saying, "Show no diffs."
- Use the red/green "eye" icon to indicate completion of and to keep track of which files you have reviewed. The <kbd>x</kbd> keyboard shortcut toggles the completion status of the file currently focused in the status bar across the top.
- Use the <kbd>p</kbd> and <kbd>n</kbd> keys to navigate to the previous and next unreviewed file—that is, those whose status is a red circle with a crossed white eye icon, meaning incomplete, as opposed to those with a green circle with a white eye, meaning complete.
- Use the <kbd>j</kbd> and <kbd>k</kbd> keys to navigate to the previous and next comment. Use <kbd>S-j</kbd> and <kbd>S-k</kbd> to navigate between the previous and next _unaddressed_ comment. Usually as the reviewer, you use the latter to go back and check on whether your previous suggestions were addressed.
- Reply with `+lgtm` to apply the "LGTM" label directly from _Reviewable_.