# Updating workshop segments

This guide covers how to update segments in the Contributor Workshop.

- [Updating workshop segments](#updating-workshop-segments)
  - [Prerequisites](#prerequisites)
    - [Fork & Clone the ``contributor-site``](#fork--clone-the-contributor-site)
    - [Checkout a branch off the workshop branch for your changes](#checkout-a-branch-off-the-workshop-branch-for-your-changes)
    - [Run the site locally with your changes](#run-the-site-locally-with-your-changes)
  - [Modifying an existing segment](#modifying-an-existing-segment)
  - [Adding a segment to the workshop](#adding-a-segment-to-the-workshop)
  - [Adding a video guide to accompany a segment](#adding-a-video-guide-to-accompany-a-segment)
  - [Save, Commit, Push, PR](#save-commit-push-pr)

## Prerequisites

### Fork & Clone the ``contributor-site``

Create a fork of the [``contributor-site``](https://github.com/kubernetes/contributor-site/tree/workshop) repo.

Clone your fork locally

### Checkout a branch off the workshop branch for your changes

```bash
git checkout workshop

git checkout -b my-changes
```

### Run the site locally with your changes

Follow the instructions in the top level [``README.md``](https://github.com/kubernetes/contributor-site/tree/master/README.md#running-the-site-locally) on the contributor-site.

You're now ready to update and add segments!

## Modifying an existing segment

Find the file you would like to edit and open it. See [the list of segments](https://github.com/kubernetes/contributor-site/tree/workshop/content/en/workshop/CONTRIBUTING.md#list-of-segments).

## Adding a segment to the workshop

Make sure to update [the list of segments](https://github.com/kubernetes/contributor-site/tree/workshop/content/en/workshop/CONTRIBUTING.md#list-of-segments), with the new page.

If inserting a page between existing pages, make sure to update the page weightings in the frontmatter and [the list of segments](https://github.com/kubernetes/contributor-site/tree/workshop/content/en/workshop/CONTRIBUTING.md#list-of-segments),.

See [the segment template](../templates/segment-template.md), on how to structure your segment.

## Adding a video guide to accompany a segment

With some of the guides, we might want a youtube video to accompany them. You can find the information for adding a video to the guide, in the [video guide](videoguide.md)

## Save, Commit, Push, PR

Once you've made the modifications and additions you want, commit and push them to your fork.

Then make a pull request of your changes into the ``workshop`` branch of the ``contributor-site`` repo.
