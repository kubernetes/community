# Calendar Guidelines

Project meetings are a life line of the Kubernetes project. Consistent
calendaring is a challenge with many different clients, corporate policies,
time zones and various iterations of Daylight Savings Time. This guide should
help you navigate some of the common pitfalls and provide some tips & best
practices.

Please feel free to PR in your favorite tips and tricks that may help others.

- [Establishing a New Meeting](#establishing-a-new-meeting)
  - [Testing Permissions](#testing-permissions)
- [Transferring Ownership](#transferring-ownership)
- [Tips](#tips)
  - [Viewing Kubernetes Project Calendars](#viewing-kubernetes-project-calendars)
  - [Adding Events to Your own Calendar](#adding-events-to-your-own-calendar)
- [Troubleshooting](#troubleshooting)
  - [Permissions Impacted After Changing Positions or Role](#permissions-impacted-after-changing-positions-or-role)


## Establishing a New Meeting

_"I'm a chair for a SIG or WG and need to set up a meeting."_

This procedure will create a calendar that allows for you and all your SIG/WG
Chairs or Tech Leads to edit and manage the meetings.

**NOTE:** As of March 2019, this is the current best practice. However with the
addition of [gsuite], this practice may change soon.

1. Use a poll service such as [doodle] that will help you get a good pulse on
   your community and when they can meet.
2. Create a [new shared calendar] titled "[SIG|WG] Foo Shared Calendar" from
   a gmail/google account that will not have problems sharing or posting
   information publicly. This may mean using a personal gmail account if your
   corporate policies restrict sharing. See [Testing Permissions] to validate
   you can share your calendar.
  - Creating a shared calendar is essential. If you change jobs, email
    addresses, or take a break from the project it allows for a smooth transfer
    of ownership.
3. [Configure access permissions and sharing:]
  - Make all event details publicly accessible.
  - Share it with full rights ("make changes and manage sharing”) to your SIG/WG
    lead mailing list and community@kubernetes.io.
  - Let your other chairs and leads know they can accidentally delete a calendar
    while trying to delete it from theirs.
  - Share with view permissions only (“see all event details”) to your SIG/WG
    mailing list.
4. Once you have a time cadence settled from your members, create a calendar
    invite with the shared calendar as the owner. Configure it with the
    following settings:
  - Name it “[SIG|WG] Foo [Time Cadence ex: Biweekly] Meetings”.
  - Set sharing to public. **NOTE:** most gmail accounts will have a "default
    visibility" setting enabled. Default visibility is usually "private" and
    will need to be set to "public".
  - In the calendar invite body - include your meeting notes, zoom information,
    and any other pertinent information that you want your group to know.
  - Invite your SIG/WG mailing list and the Kubernetes Community Public
    Calendar address: `cgnt364vd8s86hr2phapfjc6uk@group.calendar.google.com`


### Testing Permissions

Make sure your work account doesn't have restrictions for public viewing of
calendar invites you create. If you are unsure, test this with other
contributors before sending it to mailing lists. This is applicable for both the
calendar entry itself and the shared calendar if you are the chair creating it.


## Transferring Ownership

If a chair is offboarding, ensure that shared calendar permissions are
configured. Once the calendar has been migrated, send out a new invite to ensure
there are no possible ghost-entries in member's calendars.

## Tips

### Viewing Kubernetes Project Calendars

 _"I'm a contributor and want to see one of/all of the SIG calendar(s)."_

Public Kubernetes Events can be seen on the [Public Community Calendar].

All of the SIGs and WGs have meeting agendas with detailed information at the
top. You can get this information from the [SIG/WG list]. Join their mailing
list for the most up to date calendar invites. Chairs will always invite the
entire mailing list to their events.


### Adding Events to Your own Calendar

Don't copy calendars if you can help it. Copying the calendar onto your calendar
will prevent you from getting updates like a canceled meeting. Join the main
[contributor mailing list] and any  [SIG/WG list] that is of interest.

Accept the invite from the sender and you'll have the updates.


## Troubleshooting

### Permissions Impacted After Changing Positions or Role

_"I'm a chair and the person that created the meeting is either no longer with
the project or no longer at an employer that holds the invite"_

If the calendar was created as a shared calendar with edit rights granted to
other chairs and leads they should be able to edit the invite and migrate
ownership without issue. If you do not have permissions, check first by sending
an email to community@kubernetes.io. Permissions may have been granted that team
and they will be able to facilitate the change in ownership.

If there is no shared calendar and still one owner, ask the person to transfer
it to a shared calendar or you'll need to create a new one. In these cases it
often best to just create a new one to avoid any possible issues with the
previous calendar. It doesn't hurt to recreate a meeting invite every few months
to refresh invites sent to the group.


[gsuite]: https://github.com/kubernetes/community/issues/3362
[doodle]: https://doodle.com
[testing permissions]: #testing-permissions
[new shared calendar]: https://support.google.com/calendar/answer/37095?hl=en
[configure access permissions and sharing:]: https://support.google.com/calendar/answer/37082?hl=en
[SIG/WG list]: /sig-list.md
[Public Community Calendar]: https://calendar.google.com/calendar/embed?src=cgnt364vd8s86hr2phapfjc6uk%40group.calendar.google.com&ctz=America%2FLos_Angeles
[contributor mailing list]: https://groups.google.com/forum/#!forum/kubernetes-dev