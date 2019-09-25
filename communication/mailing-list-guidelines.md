# Mailing list guidelines

The Kubernetes Mailing list or Google Groups functions as the primary means of
asynchronous communication for the project's
[Special Interest Groups (SIG)][sig-list] and [Working Groups (WG)][sig-list].

### ATTENTION: SIG/WG Mailing list owners

If you are currently a moderator of a SIG or WG Mailing List. See the new policy
requirements here:

- [Mailing list annual review](#annual-permissions-review)
- [Mailing list moderation queue](#new-user-posting-queue)
  - [Creating moderation queue](#create-moderation-queue)


## Code of conduct

The Kubernetes project adheres to the community [Code of Conduct] throughout all
platforms and includes all communication mediums.

## Admins

Check the [centralized list of administrators][admins] for contact information.

To connect: Reach out to one of the listed moderators,[Mailing list owners],
the [sig contributor experience Mailing list] or the `#sig-contribex` slack
channel.

### Mailing list owners

Mailing list owners should include the Chairs for your [SIG or WG][sig-list] and
the below contacts:

- parispittman[at]google.com
- jorgec[at]vmware.com
- ihor[at]cncf.io

---

## Moderation

SIG and Working Group Mailing lists should have the [Mailing list owners] as
co-owners to the list so that administrative functions can be managed centrally
across the project.

Moderation of the SIG/WG lists is up to that individual SIG/WG. The admins
are there to help facilitate leadership changes, or various other administrative
functions.

The group management settings for Google Groups have been changed in order
to make groups simpler to manage. This has caused some breaks in certain groups 
visibility settings related to SIG and WG Google Groups.
The instructions on how to fix from Google Groups for owners of the list:
Near the top right, click **Manage group**.
- **Informtation** -> **Directory** -> **Edit the setting to set the desired 
visibility for your group.** -> **Save**.
- This [link] have all the details related to these changes.

Users who are violating the [Code of Conduct] or other negative activities
(like spamming) should be moderated.
- [Lock the thread immediately] so that people cannot reply to the thread.
- [Delete the post].
- In some cases you might need to ban a user from the group, follow
  [these instructions] on how stop a member from being able to post to the group.
  For more technical help on how to use Google Groups, check the [Groups Help]
  page.


### Moderator expectations and guidelines

Moderators should adhere to the general Kubernetes project
[moderation guidelines].


#### New user posting queue

New members who post to the Mailing list will automatically have their messages
put in the [moderation queue]. Moderators of the list will receive a
notification of their message and should process them accordingly.


#### Annual permissions review

SIG and WG Moderators must establish an annual review of their Mailing lists
to ensure their Moderator list is current and includes [Mailing List owners].
Many of the SIG and WG Mailing lists pre-date current communication policy and
an annual review ensures ownership is up to date.

This review does not need to occur at a specific recurring date and can be
combined with other actions such as SIG/WG leadership changes or sub-project
additions.


---

## Mailing list creation

Create a Google Group at https://groups.google.com/forum/#!creategroup,
following the below procedure:
- Each SIG must have two discussion groups with the following settings.
  - `kubernetes-sig-<foo>` (the discussion group):
    - Anyone can view content.
    - Anyone can join.
    - Moderate messages from non-members of the group.
    - Only members can view the list of members.
  - `kubernetes-sig-<foo>-leads` (list for the leads, to be used with Zoom and
    Calendars)
    - Only members can view group content.
    - Anyone can apply to join.
    - Moderate messages from non-members of the group.
    - Only members can view the list of members.
- Groups should be created as e-mail lists with at least three owners and must
  include the [Mailing list owners](#mailing-list-owners).
-  To add the owners, visit the **Group Settings** (drop-down menu on the right
   side), select **Direct Add Members** on the left side and add them via their
   email address (with a suitable welcome message).
- In **Members/All Members** select the [Mailing list owners] and assign them
  to the **owner role**.
- Set the following permissions to **Public**:
  - **View topics**
  - **Post**
  - **Join the Group**
- Create and share your _"meeting notes"_ Google doc with the following
  permissions settings:
  - **Can edit** for members of the newly created Mailing List.
  - **Can comment** for `kubernetes-dev@googlegroups.com`
  - **View only** for anyone with the link. **NOTE:** Depending on
    employer organization policy, this may not be possible to configure. The
    document should be copied over to an account without the restriction and 
    include the owner reference at the top of the document.

Familiarize yourself with the [moderation guidelines] for the project and create
a [moderation queue]. Chairs should be cognizant that a new group will require
an initial time investment moderation-wise as the group establishes itself.


### Create moderation queue

The moderation queue will direct all new user messages to the a moderation
queue before being posted to the Mailing List.

- From the Google Groups management page goto **Settings** -> **Moderation**.
- Configure the following settings:
  - **Post:**
    - Owners of the Group
    - Manager of the Group
    - All Members of the Group
  - **Post as the Group:**
    - Owners of the Group
    - Managers of the Group
  - **Approve Members:**
    - Managers of the Group
  - **Lock Topics:**
    - Owners of the Group
    - Managers of the Group
  - **Modify Members:** (should be Greyed out)
    - Owners of the Group
    - Managers of the Group
  - **New Member Restrictions:**
    - New member posts are moderated
  - **Reject author notification:**
    - Notify authors when moderators reject their posts -> **checked**
    - Message:
      ```
      Since you're a new subscriber you're in a moderation queue, sorry for the inconvenience, a moderator will check your message shortly.
      ```
  - **Spam messages:**
    - Send them to moderation queue and send notification to moderators.

### Archive a mailing list

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
[link]: https://support.google.com/a/answer/9191148?hl=en
