---
title: "Help Wanted and Good First Issue Labels"
weight: 9
description: |
  This document provides guidance on how and when to use the help wanted and
  good first issue labels. These are used to identify issues that have been
  specially groomed for new contributors.
---

## Overview

We use two labels to identify issues that have been specifically created or selected for new contributors: [help wanted](#help-wanted) and [good first
issue](#good-first-issue). The `good first issue` label is a subset of the `help wanted`
label, indicating that members have committed to providing extra assistance for
new contributors. All `good first issue` items also have the `help wanted`
label.

We also have some [suggestions](#suggestions-for-experienced-community-members) for using these labels to help
grow and improve our community.

## Help Wanted

Items marked with the `help wanted` label need to ensure that they meet these criteria:

- **Clear Task**
  The task is agreed upon and does not require further discussions in the
  community. Call out if that area of code is untested and requires new
  fixtures. Consensus should exist for the high-level approach.

  API and CLI behavior should be decided and included in the OP issue, for example: "The
  new command syntax is `svcat unbind NAME [--orphan] [--timeout 5m]`", with
  expected validations called out.

- **Goldilocks priority**
  The priority should not be so high that a core contributor should do it, but not too low that it
  isn't useful enough for a core contributor to spend time reviewing it, answering
  questions, helping get it into a release, etc.

- **Up-To-Date**
  Often these issues become obsolete and have already been completed, are no longer
  desired, no longer make sense, or have changed priority or difficulty.

A good example of a Help Wanted issue description can be found here: [kubernetes/test-infra#21356 (comment)](https://github.com/kubernetes/test-infra/issues/21356#issuecomment-799972711).

These commands can be used with GitHub issues to manage the `help wanted` label:

- `/help` : Adds the `help wanted` label to an issue.
- `/remove-help` : Removes the `help wanted` label from an issue. If the
  `good first issue` label is present, it is removed as well. 

## Good First Issue

Items marked with the `good first issue` label are intended for _first-time
contributors_. It indicates that members will keep an eye out for these pull
requests and shepherd it through our processes.

**New contributors should not be left to find an approver, ping for reviews,
decipher prow commands, or identify that their build failed due to a flake.**
It is important to make new contributors feel welcome and valued. We should assure them that they
will have an extra level of help with their first contribution.

After a contributor has successfully completed one or two `good first issue` items, they
should be ready to move on to `help wanted` items.

All `good first issue` items need to follow the guidelines for `help wanted`
items in addition to meeting the following criteria:

- **No Barrier to Entry**
  The task is something that a new contributor can tackle without advanced
  setup or domain knowledge.

- **Solution Explained**
  The recommended solution is clearly described in the issue.

- **Provides Context**
  If background knowledge is required, this should be explicitly mentioned and a
  list of suggested readings included.

- **Gives Examples**
  Link to examples of similar implementations so new contributors have a
  reference guide for their changes.

- **Identifies Relevant Code**
  The relevant code and tests to be changed should be linked in the issue.

- **Ready to Test**
  There should be existing tests that can be modified, or existing test cases
  fit to be copied. If the area of code doesn't have tests, before labeling the
  issue, add a test fixture. This prep often makes a great `help wanted` task!

A good example of a `good first issue` description can be found here: [kubernetes/kubernetes#68231](https://github.com/kubernetes/kubernetes/issues/68231).

These commands can be used in the GitHub issue comments to control the `good first issue` label:

- `/good-first-issue` : Adds the `good first issue` label to an issue. Also adds
  the `help wanted` label, if not already present.
- `/remove-good-first-issue` : Removes the `good first issue` label from an issue.

## Suggestions for Experienced Community Members

We encourage our more experienced members to help new contributors, so that the
Kubernetes community can continue to grow and maintain the kind, inclusive
community that we all enjoy today.

The following suggestions go a long way toward preventing "drive-by" PRs, and
ensure that our investment in new contributors is rewarded by returning contributors.

- Provide extra assistance during reviews on `good first issue` pull requests.
- Answer questions and identify useful docs.
- Offer advice such as "One way to reproduce this in a cluster is to do X and
  then you can use kubectl to poke around," or "Did you know that you can
  use fake clients to setup and test this easier?"
- Help new contributors learn enough about the project, setting up their
  environment, running tests, and navigating this area of the code so that they
  can tackle a related `help wanted` issue next time.

If you make someone feel like a part of our community, they will know that it is safe to ask
questions, that people will let them know the rules, and that their
contributions are helpful and appreciated. They will stick around! 🌈

- Encourage new contributors to seek help on the appropriate slack channels,
  introduce them, and include them in your conversations.
- Invite them to the SIG meetings.
- Give credit to new contributors so that others get to know them: "Hey, would
  someone help give a second LGTM on @newperson's first PR on chocolate
  bunnies?" Mention them in the SIG channel and meeting, and thank them on Twitter or
  #shoutouts.
- Use all the emoji in your approve or lgtm comment. 💖 🚀
- Let them know that their `good first issue` is getting extra attention to make
  the first one easier and help them find a follow-up issue.
- Suggest a related `help wanted` so that can build up experience in an area.
- People are more likely to continue contributing when they know what to expect.
  They want to know the acceptable way to ask for people to review a PR, and how to nudge things along
  when a PR is stalled. Show them how we operate by helping move their first PR
  along.
- If you have time, let the contributor know that they can DM you with questions
  that they aren't yet comfortable asking the wider group.
