---
title: "Help Wanted and Good First Issue Labels"
weight: 1
slug: "help-wanted"
---

# Overview

We use two labels [help wanted](#help-wanted) and [good first
issue](#good-first-issue) to identify issues that have been specially groomed
for new contributors. The `good first issue` label is a subset of `help wanted`
label, indicating that members have committed to providing extra assistance for
new contributors. All `good first issue` items also have the `help wanted`
label.

We also have some [suggestions](#suggestions) for using these labels to help
grow and improve our community.

## Help Wanted

Items marked with the `help wanted` label need to ensure that they are:

- **Low Barrier to Entry**

  It should be tractable for new contributors. Documentation on how that type of
  change should be made should already exist.

- **Clear Task**

  The task is agreed upon and does not require further discussions in the
  community. Call out if that area of code is untested and requires new
  fixtures.

  API / CLI behavior is decided and included in the OP issue, for example: _"The
  new command syntax is `svcat unbind NAME [--orphan] [--timeout 5m]`"_, with
  expected validations called out.

- **Goldilocks priority**

  Not too high that a core contributor should do it, but not too low that it
  isn't useful enough for a core contributor to spend time to review it, answer
  questions, help get it into a release, etc.

- **Up-To-Date**

  Often these issues become obsolete and have already been done, are no longer
  desired, no longer make sense, have changed priority or difficulty , etc.

Related commands:

- `/help` : Adds the `help wanted` label to an issue.
- `/remove-help` : Removes the `help wanted` label from an issue. If the
  `good first issue` label is present, it is removed as well.

## Good First Issue

Items marked with the `good first issue` label are intended for _first-time
contributors_. It indicates that members will keep an eye out for these pull
requests and shepherd it through our processes.

**New contributors should not be left to find an approver, ping for reviews,
decipher prow commands, or identify that their build failed due to a flake.**
This makes new contributors feel welcome, valued, and assures them that they
will have an extra level of help with their first contribution.

After a contributor has successfully completed 1-2 `good first issue`'s, they
should be ready to move on to `help wanted` items, saving remaining `good first
issue`'s for other new contributors.

These items need to ensure that they follow the guidelines for `help wanted`
labels (above) in addition to meeting the following criteria:

- **No Barrier to Entry**

  The task is something that a new contributor can tackle without advanced
  setup, or domain knowledge.

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

Related commands:

- `/good-first-issue` : Adds the `good first issue` label to an issue. Also adds
  the `help wanted` label, if not already present.
- `/remove-good-first-issue` : Removes the `good first issue` label from an
  issue.

# Suggestions

We encourage our more experienced members to help new contributors, so that the
Kubernetes community can continue to grow and maintain the kind, inclusive
community that we all enjoy today.

The following suggestions go a long way toward preventing "drive-by" PRs, and
ensure that our investment in new contributors is rewarded by them coming back
and becoming regulars.

Provide extra assistance during reviews on `good first issue` pull requests:
- Answer questions and identify useful docs.
- Offer advice such as _"One way to reproduce this in a cluster is to do X and
  then you can use kubectl to    poke around"_, or _"Did you know that you can
  use fake clients to setup and test this easier?"_.
- Help new contributors learn enough about the project, setting up their
  environment, running tests, and navigating this area of the code so that they
  can tackle a related `help wanted` issue next time.

If you make someone feel like a part of our community, that it's safe to ask
questions, that people will let them know the rules/norms, that their
contributions are helpful and appreciated... they will stick around! 🌈
- Encourage new contributors to seek help on the appropriate slack channels,
  introduce them, and include them in your conversations.
- Invite them to the SIG meetings.
- Give credit to new contributors so that others get to know them, _"Hey, would
  someone help give a second LGTM on @newperson's first PR on chocolate
  bunnies?"_. Mention them in the SIG channel/meeting, thank them on twitter or
  #shoutouts.
- Use all the emoji in your approve or lgtm comment. 💖 🚀
- Let them know that their `good first issue` is getting extra attention to make
  the first one easier and help them find a follow-up issue.
- Suggest a related `help wanted` so that can build up experience in an area.
- People are more likely to continue contributing when they know what to expect,
  what's the acceptable way to ask for people to review a PR, nudge things along
  when a PR is stalled. Show them how we operate by helping move their first PR
  along.
- If you have time, let the contributor know that they can DM you with questions
  that they aren't yet comfortable asking the wider group.
