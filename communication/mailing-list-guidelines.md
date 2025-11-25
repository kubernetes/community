---
title: "Mailing List Guidelines"
description: |
  Guidelines for Mailing list (Google Group) creation, sharing, archival and
  moderation.
---

<!-- omit in toc -->
# Mailing list guidelines

The Kubernetes mailing list or Google Groups functions as the primary means of
asynchronous communication for the project's
[Special Interest Groups (SIG)][sig-list], [Working Groups (WG)][sig-list], and 
large subprojects.


- [Code of conduct](#code-of-conduct)
- [Admins](#admins)
  - [Mailing list owners](#mailing-list-owners)
- [Moderation](#moderation)
  - [Moderator expectations and guidelines](#moderator-expectations-and-guidelines)
    - [New user posting queue](#new-user-posting-queue)
    - [Annual permissions review](#annual-permissions-review)
- [Mailing list creation](#mailing-list-creation)
  - [Prerequisites for creating a mailing list](#prerequisites-for-creating-a-mailing-list)
  - [Create the leads and members mailing lists](#create-the-leads-and-members-mailing-lists)
- [Set up shared calendars and meeting with a mailing list](#set-up-shared-calendars-and-meeting-with-a-mailing-list)
  - [Prerequisites for sharing a calendar and meeting notes](#prerequisites-for-sharing-a-calendar-and-meeting-notes)
  - [Sharing the calendar with the Google Group](#sharing-the-calendar-with-the-google-group)
  - [Sharing the meeting notes with the Google Group](#sharing-the-meeting-notes-with-the-google-group)
- [Archive a mailing list](#archive-a-mailing-list)


## Code of conduct

The Kubernetes project adheres to the community [Code of Conduct] throughout all
platforms and includes all communication mediums.

## Admins

Check the [centralized list of administrators][admins] for contact information.

To connect: Reach out to one of the listed moderators, [mailing list owners],
the [SIG Contributor Experience mailing list] or the `#sig-contribex` slack
channel.

### Mailing list owners

Mailing list owners should include the Chairs for your [SIG or WG][sig-list], 
or the leads for your subproject, and the below contacts:

- contributors@kubernetes.io

---

## Moderation

SIG, Working Group, and subproject mailing lists should have the 
[mailing list owners] as co-owners to the list so that administrative functions 
can be managed centrally across the project.

Moderation of the SIG/WG/subproject lists is up to that individual 
SIG/WG/subproject. The admins are there to help facilitate leadership changes, 
or various other administrative functions.

Users who are violating the [Code of Conduct] or other negative activities
(like spamming) should be moderated.
- [Lock the thread immediately] so that people cannot reply to the thread.
- [Delete the post].
- In some cases you might need to ban a user from the group, follow
  [these instructions] on how stop a member from being able to post to the 
  group.
  For more technical help on how to use Google Groups, check the [Groups Help]
  page.


### Moderator expectations and guidelines

Moderators should adhere to the general Kubernetes project
[moderation guidelines].


#### New user posting queue

New members who post to the mailing list will automatically have their messages
put in the [moderation queue]. Moderators of the list will receive a
notification of their message and should process them accordingly.


#### Annual permissions review

SIG, WG, and subproject Moderators must establish an annual review of their 
mailing lists to ensure their Moderator list is current and includes 
[mailing list owners]. Many of the SIG and WG mailing lists pre-date current 
communication policy and an annual review ensures ownership is up to date.

This review does not need to occur at a specific recurring date and can be
combined with other actions such as SIG/WG/subproject leadership changes or 
sub-project additions.


---

## Mailing list creation

All SIGs and WGs require two discussion groups: one for leads/chairs, and one 
for members. Subprojects that opt to have a mailing list only require one for 
members.

### Prerequisites for creating a mailing list

- An email account that can create Google Groups and add members external to 
your organization to a Google Group mailing list. **This might not be possible 
with your employer's email account**. You might need to use a personal email 
account.
- At least 3 mailing list owners (leads), in addition to 
contributors@kubernetes.io
- Familiarity with the [moderation guidelines] for the project and 
[moderation queue]s. Chairs should be cognizant that a new group will require
an initial time investment moderation-wise as the group establishes itself.


### Create the leads and members mailing lists

> **Note:** You will need follow these steps twice! Once for the leads mailing 
list, and again for the members mailing list.

1. Navigate to https://groups.google.com/forum/#!creategroup and fill out the 
**Enter group info** form as follows:

  | Field | Leads ML value | Members ML value | 
  | --- | --- | --- |
  | **Group name** | SIGs: `kubernetes-sig-<foo>-leads`<br>WGs: `kubernetes-wg-<foo>-leads` | SIGs: `kubernetes-sig-<foo>`<br>WGs: `kubernetes-wg-<foo>`<br>Subprojects: `kubernetes-<foo>` | 
  | **Group email address** | Leave as-is | Leave as-is
  | **Group description** | Leads ML for Kubernetes [SIG/WG] Foo | Members ML for Kubernetes [SIG/WG/subproject] Foo |

  Click **Next**.
  
2. Fill out the **Choose privacy settings** with these options:
  
  | Field | Leads ML value | Members ML value | 
  | --- | --- | --- |
  | **Who can see the group** | Group members | Anyone on the web | 
  | **Who can join group** | Invited users only | Anyone on the web | 
  | **Who can view conversations** | Group members | Anyone on the web | 
  | **Who can post** | Anyone on the web | Anyone on the web |
  | **Who can view members** | Group members | Group members | 

  Click **Next**. 
  
3. Fill out the **Add members** form as follows:

  | Field | Leads ML value | Members ML value | 
  | --- | --- | --- |
  | **Group owners** | All SIG/WG leads and contributors@kubernetes.io | All SIG/WG/subproject leads and contributors@kubernetes.io | 
  
  > **Note:** You can add new owners to a mailing list at any time in the 
  **People > Members** screen.
  
  Leave all other fields as-is. Click **Next.**
  
4. Once the group is created, navigate to your group in the Google Groups UI and
 go to **Group settings** to continue setting up permissions. Set the following
  settings:
  
  **Member Privacy**
  
  | Field | Leads ML value | Members ML value | 
  | --- | --- | --- |
  | **Identification required for new members** | Either display name or Google profile | Either display name or Google profile |
  | **Who can view the member's email addresses?** | Group managers | Group managers |
  
  **Posting policies** 
  
  | Field | Leads ML value | Members ML value | 
  | --- | --- | --- |
  | **Conversation history** | On | On |
  | **Who can moderate content** | Group managers | Group managers | 
  | **Who can moderate metadata** | Group members | Group members |
  | **Who can post as the group** | Group owners | Group owners | 
  | **Message moderation** | No moderation | Moderate messages from non-members | 
  | **New member restrictions** | No posting restriction for new members | New member posts are moderated |

  **Email options** 
  
  | Field | Leads ML value | Members ML value | 
  | --- | --- | --- |
  | **Subject prefix** | SIGs: `[k8s-sig-<foo>-leads]`<br>WGs: `[k8s-wg-<foo>-leads]` | SIGs: `[k8s-sig-<foo>]`<br>WGs: `[k8s-wg-<foo>]`<br>Subprojects: `[k8s-<foo>]` | 
  | **Email footer** | Include the standard Groups footer | Include the standard Groups footer |
  | **Group email language** | English (or your group's default language) | English (or your group's default language) |
  
  **Member moderation** 
  
  | Field | Leads ML value | Members ML value | 
  | --- | --- | --- |
  | **Who can manage members** | Group managers | Group managers |
  | **Who can adjust roles** | Group managers | Group managers

5.  Click **Save changes**. 
  Once your mailing list is created, it should also be added to the [sigs.yaml] 
  file. For subprojects, it should be added like:
  ```yaml
  - name: Foo
    contact:
      mailing_list: [link to Google Group]
  ```

## Set up shared calendars and meeting with a mailing list  

Once you've set up your SIG/WG mailing list, you'll need to: 
- Share a calendar with meeting invites on it with the mailing list 
- Share a meeting notes google doc with the mailing list 

### Prerequisites for sharing a calendar and meeting notes

- A member's Google Group.
- A shared calendar.
  > **Note:** Like with mailing lists, your organization's permissions might not
   let you share calendars with the correct permissions. You might need to use a
    personal email address.
  
### Sharing the calendar with the Google Group

You must share the meeting calendar with the following people:
- All leads (individually)
- The kubernetes-[sig-/wg-]foo-leads mailing list
- contributors@kubernetes.io 
- The kubernetes-[sig-/wg-]foo (members) mailing list

1. In Google Calendar, click on the calendar's **...** menu and select 
**Settings and sharing**.
2. In **Access permissions**, check **Make available to public**.
3. Under **Share with specific people, do the following:**
  - For each lead, contributors@kubernetes.io, and 
    sig-foo-leads@kubernetes.io or kubernetes-sig-foo-leads@googlegroups.com (depending which group is active):
    1. Add their email
    2. Give them the permission **Make changes and manage sharing**. 
  - For sig-foo@kubernetes.io or kubernetes-sig-foo@googlegroups.com (depending which group is active), add them and give them the 
  permission **See all event details**.

> **Note:** You need to add the member's mailing list as a guest to any meeting 
invites on the shared calendar for an invite to be sent to members of the group.

### Sharing the meeting notes with the Google Group

- Create and share your _"meeting notes"_ Google doc with the following
  permissions settings:
  - **Can edit** for members of the newly created Mailing List.
  - **Can comment** for `dev@kubernetes.io`
  - **View only** for anyone with the link. **NOTE:** Depending on
    employer organization policy, this may not be possible to configure. The
    document should be copied over to an account without the restriction and 
    include the owner reference at the top of the document.

## Archive a mailing list

To archive a mailing list, use the below procedure.

- Send a final notice to the mailing list that it is closed. This notice should
  include a brief description as to why and include links to any other relevant
  information.
- From the Google Groups management page goto **Information** ->
  **General Information**.
  - Configure the following settings:
    - **Group description** -> Set to the same message used for the final 
      mailing list notice.
- From the Google Groups management page goto **Information** ->
  **Content Control**.
  - Configure the following settings:
    - **Archive messages to the group** -> **checked**
- From the Google Groups management page goto **Permissions** ->
  **Basic permissions**.
  - Configure the following settings:
    - **View Topics:**
      - Managers of the group
      - All members of the group
      - Anyone on the web
    - **Post:** -> uncheck all options
    - **Join the group:**
      - Only invited users

[mailing list owners]: #mailing-list-owners
[moderation queue]: #create-moderation-queue
[sig-list]: /sig-list.md
[Code of Conduct]: /code-of-conduct.md
[admins]: ./moderators.md#mailing-lists
[sig contributor experience mailing list]: https://groups.google.com/forum/#!forum/kubernetes-sig-contribex
[moderation guidelines]: ./moderation.md
[lock the thread immediately]: https://support.google.com/groups/answer/2466386?hl=en#
[delete the post]: https://support.google.com/groups/answer/1046523?hl=en
[these instructions]: https://support.google.com/groups/answer/2646833?hl=en&ref_topic=2458761#
[groups help]: https://support.google.com/groups/answer/2466386?hl=en&ref_topic=2458761
[sigs.yaml]: /sigs.yaml
