---
title: "Notification Management Best Practices"
description: |
  A collection of tips, filters, and best practices for managing Mailing List
  and GitHub notifications.
---

# Mailing list and GitHub notification best practice

The Kubernetes Mailing list or Google Groups functions as the primary means of
asynchronous communication for the project's [Special Interest Groups (SIG)] and
[Working Groups (WG)]. That's why you may want to set your filters in
your email account to attain a good signal-to-noise ratio with regards to the
mailing list messages and GitHub notifications. All the steps below are for Gmail
users, however similar filters can be made in other email clients.

**Note:** Alternatively, we highly encourage people to use [Gubernator] to
view and acknowledge their Pull Request review notifications.

[Special Interest Groups (SIG)]: /sig-list.md#master-sig-list
[Working Groups (WG)]: /sig-list.md#master-working-group-list
[Gubernator]: https://gubernator.k8s.io/pr


## Creating filters for Kubernetes Mailing lists

It depends on the SIG or/and WG you are involved in. You can setup filters for
your Gmail account to be able to categorize emails from  different mailing lists.


Create a filter following the procedure below:
- In your Gmail account click on **Settings**:
  - **Filters and Blocked Addresses** -> Scroll down and click **create a new filter**
  - In the **to** fields write the email of the SIG's Google Group. 
  - **Create filter** -> Check the box ** Apply the label** and create new
    label by choosing **New label...** in the dropdown list.
  - Click on the **Create filter**.
- Create filter directly for lists:
  - **Matches:** list:"dev@kubernetes.io"
  - **Do this:** Apply label "lists/kubernetes-dev"


For more assistance on creating filters, see the Gmail help page on
[Creating rules to filter your email].


## Creating filters for Kubernetes Github notifications

These below suggested Gmail filters can help you organize and obtain better
signal GitHub notification emails.

Before you begin, you must know how to create filters in Gmail. For this
procedure, see the Gmail help page on [Creating rules to filter your email].

- Apply a blue label on anything kubernetes-related:
  - **Matches:** (kubernetes OR kubernetes-client OR kubernetes-sigs OR
    kubernetes-csi)
  - **Do this:** Apply label "k8s", Mark it as important
- Archive your own actions (sending these is an option in Github's settings).
  You can send them but also archive them, so whenever you need to see the history
  of  an issue you can:
  - **Matches:** to:(your_activity@noreply.github.com)
  - **Do this:** Skip Inbox, Mark as read
- Skip bot comments:
  - **Matches:** (from:(notifications@github.com) (from:(k8s-merge-robot) OR
    from:(Kubernetes Prow Robot) OR from:(k8s-ci-robot)))
  - **Do this:** Skip Inbox, Mark as read
- Skip push notifications:
  - **Matches:** to:(push@noreply.github.com)
  - **Do this:** Skip Inbox, Mark as read
- Apply a red label on things assigned to you and/or things request to be reviewed:
  - **Matches:** to:(assign@noreply.github.com)
  - **Do this:** Star it, Apply label "gh/assigned", Mark it as important
  - **Matches:** to:(review_requested@noreply.github.com)
  - **Do this:** Star it, Apply label "gh/requested_review", Mark it as important
- Apply an orange label on things you commented on:
  - **Matches:** to:(comment@noreply.github.com)
  - **Do this:** Star it, Apply label "gh/commented"
- Apply a yellow label on things you have been mentioned on:
  - **Matches:** to:(mention@noreply.github.com)
  - **Do this:** Apply label "gh/mentioned"
- Apply a grey label:
  - **Matches:** to:(team_mention@noreply.github.com)
  - **Do this:** Apply label "gh/team_mention"
  - **Matches:** to:(author@noreply.github.com)
  - **Do this:** Star it, Apply label "gh/authored", Mark it as important
- Skip messages about issues that you are not participating in, but leave them unread:
  - **Matches:** from:(notifications@github.com) to:(subscribed@noreply.github.com)
  - **Do this:** Skip Inbox
- Categorize per repository:
  - **Matches:** list:(community.kubernetes.github.com)
  - **Do this:** Apply label "k8s/community"

These suggestions come largely from an old [kubernetes-dev] mailing list [thread]
on Gmail filters for Kubernetes.

[kubernetes-dev]: https://groups.google.com/g/kubernetes-dev
[thread]: https://groups.google.com/forum/#!topic/kubernetes-dev/5qU8irU7_tE/discussion

<!-- shared links -->
[Creating rules to filter your email]: https://support.google.com/mail/answer/6579?hl=en
