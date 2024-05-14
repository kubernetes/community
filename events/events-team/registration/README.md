# Registration Handbook


The Registration lead role is charged with handling all aspects of attendee
registration workflow; including before the event and potentially day-of. This
includes:

- Creating the registration system or coordinating with the CNCF Events manager
  to make one.
- Staffing day-of registration needs (if required).
- Steering attendees to the right session track (Active/Current Contributor
  vs New Contributor).
- Being the point person for registration questions.
- Giving regular reports and information to other Event Coordinators.

---

- [Skills and Experience](#skills-and-experience)
- [Registration Process](#registration-process)
  - [Assembling the Questions and Paths](#assembling-the-questions-and-paths)
  - [Selecting the Registration System](#selecting-the-registration-system)
  - [Approval Process](#approval-process)
    - [Helpful Approval Tips and Facts](#helpful-approval-tips-and-facts)
- [Reports](#reports)

---

## Skills and Experience

A successful Registration lead should have the following qualities:

**Strong written and verbal communication skills.**
- You may be the first person an attendee interacts with. Be empathetic, and
  remember that there may be language differences in play.

**Detail oriented.**
- Registration mistakes happen. Did the person register for the event but not
  select any sessions? Did they misspell their github handle? Catching these
  before the event improves everyone's experience.

**Comfortable working with spreadsheets and reports.**
- Registration requests and reports are all done in spreadsheets. Being
  comfortable with queries and syntax is incredibly helpful, both for the
  registration approval process and later report generation.

## Shadow to the Registration Lead Role
This role can include one or more shadows. The Registration Lead can delegate/assign the tasks to the shadows. The expectation from a shadow is to lead one of the events-team roles in an upcoming summit. The shadows to the Registration Lead are responsible for handling all Registration lead related activities in the absence of the Lead.


## Registration Process

### Assembling the Questions and Paths

Well before the event, collect the information needed for the registration form,
and determine the registration workflow. 

All attendees should be asked the following:

| Question                                                       | Type      | Required | Notes                                                                                                          |
|----------------------------------------------------------------|-----------|----------|----------------------------------------------------------------------------------------------------------------|
| Email                                                          | Free Text | Yes      |                                                                                                                |
| First Name                                                     | Free Text | Yes      |                                                                                                                |
| Last Name                                                      | Free Text | Yes      |                                                                                                                |
| Job Title                                                      | Free Text | Yes      |                                                                                                                |
| Company/Organization                                           | Free Text | Yes      |                                                                                                                |
| GitHub Username *If your GitHub Profile URL is https://github.com/username , only enter username not @username. If you don't have a GitHub username, please enter N/A*   | Free Text | Yes      |                                                                                                                |
| Swag Related Question                                          | Dropdown  | Yes      | Swag question is generally shirt size.                                                                         |
| Please specify dietary needs (if any)                          | Dropdown  | No      | Options: `none`, `Gluten Free`, `Vegetarian`, `Vegan`, `Halal`,`Other`                                                 |
| Do you have a disability that we should be mindful of as we try to accommodate everyone for this event?  | Dropdown  | No      | Options: `Yes`, `No`    |
| What email did you use to register for KubeCon + CloudNative Con <NA/Europe> <year>? **KubeCon + CloudNative Con <NA/Europe> registration is required to attend Kubernetes Contributor Summit <region> <year>**                             | Free Text | Yes      |                                                                                                                |   
| Are you a member of one of the [Kubernetes GitHub Orgs](http://git.k8s.io/community/github-management#actively-used-github-organizations)  | Dropdown | Options: `Yes`, `No` [Pop up if no is selected] Attending the Kubernetes Contributor Summit in-person is limited to Kubernetes Org Members and Sponsored Attendees. If you have questions, please email summit-team@kubernetes.io.  Yes      |                                                                                                                |    
| What SIGs or WGs are you most active in? Only list the top three. *This will help us plan content and activities.*                 | Dropdown  | Yes      | Options: <list all SIGs and WGs>                                             |
| What sessions are you most looking forward to? *This will help us plan content and activities.*                 | Dropdown  | No       | Options: `Unconference`, `Prepared Presentations`, `Steering AMA`, `SIG Discussions/Working Sessions`, `Impromptu Discussions`, `Social/Mingling`                                                |
| What other session(s) would you like to see at the summit? What else would make this event valuable to you? *We've changed content based on contributor feedback to this question in the past.*                          | Free Text | Yes      |                                                                                                                |
| Social Event                                                   | Dropdown  | Yes      | Required if there is an independent social event.                                                              |
| Emergency Contact                                              | Free Text | No       | Name and Phone Number. Phone number to be verified.                                                            |
| How can we make this a valuable event for you?                 | Free Text | No       |                                                                                                                |
| If you had a speaking slot, what would you like to talk about? | Free Text | No       |                                                                                                                |
| Where did you hear about the Contributor Summit?               | Checkbox  | No       | Options: `X (formerly Twitter)`, `Mastodon`, `Kubernetes-dev mailing list`, `GitHub`, `Monthly Kubernetes Community Meeting`, `SIG/WG Meeting`, `Slack`, `Kubernetes Blog`, `Other`          |

Additionally, a separate form can be created for registering the `summit staff`.

### Selecting the Registration System

The registration system selected should support the questions and paths that
were determined for the event. Other potential things to consider when selecting
a system include:
- Is the reg system available in the region the event is being hosted. For
  example, Google Forms is not available in certain regions. If it is used an
  alternative registration path should be considered.
- Can you independently generate invite links or invite people by email for
  sending invites to leads or guests.
- Does it support an approval based workflow? If some registration paths have
  restrictions, an approval based workflow should be considered.
- Does it support Single Sign On (SSO) or GitHub integration? This can save time
  both for attendees and as Registration lead if an approval based workflow is
  being used.
- Can an attendee login and update their own information?
- How easy is it to get data out of the Registration System? The easier it is to
  get data out, the easier it will be to generate reports or use the information
  in an approval workflow.
- Can it integrate with [sched]? If it can automatically add an attendee, it
  will make for a less error prone invite process.

In past events, registration systems that have been used were [Google Forms],
[SurveyMonkey], and [cvent]. All have their pluses and minuses.

[Google Forms] is flexible and easy to get data in and out of, as it is
essentially just a form on top of a spreadsheet. However, it lacks useful other
functionality and is blocked in certain regions.

[SurveyMonkey] is flexible, but as of April 2019 should not be considered as
emails from its invite system are frequently blocked or ignored.

[cvent] is the current system used by the [Linux Foundation] (LF) and supports
SSO, along with multiple workflows. However, no direct access can be given to a
non LF employee and it can be difficult to troubleshoot workflows as all changes
require going through the LF.

Other options exist and should be evaluated when making your decision for the
registration system.

### Communications

Work with the Communications Lead to include the link to the registration page and future 
communications efforts to drive registration.

### Approval Process

The Kubernetes Contributor Summit is for active contributors to the Kubernetes project.
Registrants will be approved if they are org members of one of the following Kubernetes Orgs:
kubernetes, kubernetes-client, kubernetes-csi, kubernetes-sigs, etcd-io.

A SIG lead (co-chair, tech lead, subproject lead) can sponsor a non-org member to attend the summit.
Kubernetes Contributor Summit speakers who are non-org members can attend the summit.

<!--
Discussed in the KCSEU 2024 retro, we need to make exception requests less personal
We can use an existing method like use the same or similar criteria of the elections
or have requests go to the Summit Staff who then ask SIG/WG leads.
-->

There will be a  **Kubernetes Contributor Summit Planning Doc** , in the past this 
was a Google Spreadsheet.

As soon as the **Kubernetes Contributor Summit Planning Doc** is available, populate 
the **org - DO NOT DELTE** tab (the name may change in the future but is always the 
tab that compares if the registrant is an org member) in the 
**Kubernetes Contributor Summit Planning Doc**.

There are 2 options to create a list of org members:

```
git clone https://github.com/kubernetes/org.git
cd org
yq '.admins + .members' \
  config/etcd-io/org.yaml \
  config/kubernetes/org.yaml \
  config/kubernetes-client/org.yaml \
  config/kubernetes-csi/org.yaml \
  config/kubernetes-sigs/org.yaml \
  | jq -s 'add | .[] | ascii_downcase' | jq -s 'sort | unique | .[]' \
  | sed -e 's/"//g' > members.txt
```

an alternative is 

```
git clone https://github.com/kubernetes/org.git
cd org
yq '.admins + .members' \
  config/etcd-io/org.yaml \
  config/kubernetes/org.yaml \
  config/kubernetes-client/org.yaml \
  config/kubernetes-csi/org.yaml \
  config/kubernetes-sigs/org.yaml \
  | sed -e 's/---//g' | sed -e 's/- //g' \
  | sed -e 's/"//g' | sed -e '/^$/d' \
  | awk '{print tolower($0)}' | sort | uniq > members.txt
```

Populate what's in members.txt in the **org - DO NOT DELTE** tab in the 
**Kubernetes Contributor Summit Planning Doc**.

Create a list of staff, shadows, volunteers including day-of volunteers. You can 
create this list from the GitHub staffing issue e.g. https://github.com/kubernetes/community/issues/7611
and query the Summit Staff leads for their day-of volunteers. The Day-of Operations
team will likely have volunteers not listed on the GitHub staffing issue.

Once the talks for the Kubernetes Contributor Summit have been selected, create 
a list of speakers and note if any are not org members so they can be approved
when they register.

Review past declined registrants if any are speakers, email them to re-register 
for the Contributor Summit.

#### Exception Request Scenarios

**A person is not an org member and would like to attend the Contributor Summit**

With 1600+ Kubernetes Org members and limited space and budget requirements, we've
made the decision to limit attendance to those that are already actively engaged
with the project. 

Contributor Summit speakers without org membership can attend the Contributor Summit.

A SIG lead (co-chair, tech lead, subproject lead) can sponsor a non-org member to attend the summit,
this process is TBD.

<!--
Discussed in the KCSEU 2024 retro, we need to make exception requests less personal
We can use an existing method like use the same or similar criteria of the elections
or have requests go to the Summit Staff who then ask SIG/WG leads.
-->

**A speaker is not registered for KubeCon + CloudNativeCon**

KubeCon + CloudNativeCon registration is a requirement to attend the Contributor Summit.
Speakers must be registered for KubeCon + CloudNativeCon.

### Dietary Requirements

Confirm with the CNCF staff rep on who will keep track of dietary requirements.
If there are any email requests, add them to the **Dietary Requirements** tab 
in the **Kubernetes Contributor Summit Planning Doc** and notify the CNCF staff
rep.

### Celebration Guests

The CNCF staff may have the ability to keep track of registrants who are bringing
a plus 1 to the social. There may be emails to request to bring a plus 1.
Add any requests via email to the **Celebration Guests** tab in the 
**Kubernetes Contributor Summit Planning Doc**.

#### Helpful Approval Tips and Facts

In the attendee spreadsheet, knowing several facts about the attendee make it
much easier to process. These facts can be gleamed from external sources such
as [devstats], the [sig-list], [voters], the lists of [Kubernetes Org Members]
 and if they happen to be in any [owners files].

<!--
TODO: Link to spreadsheet example of queries and data stored in gsuite shared
space. NOT tied to personal account.

include devstats query in spreadsheet:
curl 'https://k8s.devstats.cncf.io/api/tsdb/query' -H 'Pragma: no-cache' -H 'Accept-Encoding: gzip, deflate, br' -H 'X-Grafana-Org-Id: 1' -H 'Content-Type: application/json;charset=UTF-8' -H 'Accept: application/json, text/plain, */*' -H 'Cache-Control: no-cache' --data-binary $'{"from":"1396302955111","to":"now","queries":[{"refId":"A","intervalMs":43200000,"maxDataPoints":1920,"datasourceId":5,"rawSql":"select name, value from shdev where series = \'hdev_contributionsallall\' and period = \'y10\'","format":"table"}]}' --compressed | jq -r '.[].A.tables[].rows[] | @csv' > devstats.csv
-->

With that information combined with comments and other information in a single
sheet, most attendees can be processed quickly and any errors found.

For attendees that might be questionable, other useful sources of information
can be gleamed from [Slack], the [Mailing Lists], and [Discuss].


## Reports

Once registration has opened, you should give the rest of the Event Coordinators
at least a weekly update regarding registration numbers and statistics. This
information is useful to keep track on the general _"health"_ of the event
leading up to it.

As an example, if the event is close to capacity or over capacity, you should
communicate that fact and let the Marketing Lead know to start getting the word
out. Another example, if the New Contributor Workshop (NCW) session is lacking
attendees, this information should be given to the NCW lead, Marketing lead, and
the Accessibility, Inclusiveness, and Diversity lead so that they may take
appropriate action.

**Example Report:**

| Registration      | Total | Approved | Social | Summit | Leads  | Voters | Owners |
|-------------------|-------|----------|--------|--------|--------|--------|--------|
| Active            | 151   | 143      | 134    | 124    | 30     | 87     | 101    |
| New               | 78    | 74       | 60     | 74     | 0      | 0      | 0      |
| Total             | 229   | 217      | 194    | 198    | 30     | 87     | 101    |
| % Approved/Type   |       |          |        |        | 20.98% | 60.84% | 70.63% |
| % Going to Summit |       |          |        |        | 24.19% | 70.16% | 81.45% |

| NCW           | 101 Session | 201 Session |
|---------------|-------------|-------------|
| Registrations | 30          | 46          |


After the event, the same information is useful for the Event Retrospective and
combined with the post-event survey data, can provide insights in the event
itself.


## Registration at the Contributor Summit

If there is a registration desk / badge pickup at the Contributor Summit, 
it is helpful to have a Registration Team member at the registration desk during 
the morning of the Kubernetes Contributor Summit to assist with questions.


[sched]: http://sched.com
[google forms]: https://www.google.com/forms/about/
[surveymonkey]: https://www.surveymonkey.com/
[cvent]: http://cvent.com
[linux foundation]:https://www.linuxfoundation.org/
[devstats]: https://k8s.devstats.cncf.io/d/13/developer-activity-counts-by-repository-group?orgId=1
[sig-list]: /sig-list.yaml
[voters]: http://git.k8s.io/steering/elections.md#eligibility-for-voting
[kubernetes org members]: https://git.k8s.io/org
[owners files]: https://cs.k8s.io/?q=&i=fosho&files=OWNERS&repos=
[slack]: http://slack.k8s.io
[mailing lists]: /sig-list.md
[discuss]: https://discuss.kubernetes.io
