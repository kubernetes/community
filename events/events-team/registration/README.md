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
  comfortable with queries and syntax is incredibly helpful, both for for the
  registration approval process and later report generation.


## Registration Process

### Assembling the Questions and Paths

Well before the event, collect the information needed for the registration form,
and determine the registration workflow. In general this boils down to are they
an _"Active/Current Contributor"_ vs a _"New Contributor"_. With a few
independent questions for each registration path.

For the _"New Contributor"_ path there may be additional paths or questions such
as a _"101"_ or _"201"_ track along with other questions that will aid the
workshop organizers.

In addition to any path specific questions, all attendees should be asked the
following:

| Question                                                       | Required | Notes                                                         |
|----------------------------------------------------------------|----------|---------------------------------------------------------------|
| Name                                                           | yes      |                                                               |
| Email                                                          | yes      |                                                               |
| GitHub Handle                                                  | yes      |                                                               |
| Signed the CLA                                                 | yes      | Provide link to CLA                                           |
| Food Restrictions                                              | yes      | values: `none`, `Gluten Free`, `Vegetarian`, `Vegan`, `Other` |
| Swag Related Question                                          | yes*     | Swag question is generally shirt size.                        |
| Social Event                                                   | yes*     | Required if there is an independent social event.             |
| Emergency Contact                                              | No       |                                                               |
| How can we make this a valuable event for you?                 | No       |                                                               |
| If you had a speaking slot, what would you like to talk about? | No       |                                                               |


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

<!--
TODO: insert link/note to reg system used for KubeCon CN
-->
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
require going through through the LF.

Other options exist and should be evaluated when making your decision for the
registration system.


### Approval Process

An approval process for registration is commonly used for multiple reasons, but
really boils down to ensuring the right people are in the right track or room.
The goal is for all contributors to get the best experience out of the event and
that may mean asking them to switch tracks.

A few scenarios that have occurred:

**A person has signed up for the New Contributor Workshop, but in their
comments have mentioned they are unfamiliar with Kubernetes.** 

They may have interpreted the event as an _"Introduction to Kubernetes"_ from an
end-user perspective. In this case it is best to reach out to them and provide
more context for the event. They will likely apologize and ask to be removed.

**A person who looks to be an Active Contributor has signed up for the New 
Contributor Track.**

This could be accidental, or in some scenarios a non-code contributor is looking
for a better introduction to the code base branch into code-based contributions.

**A person registered for the Active or Current Contributor track, but has not
participated with the project.**

They may have accidentally selected the wrong track, or believe they should go
for other reasons. It is best to reach out let them know about the intended
audience and that they may get more out of the other track. Use your best
judgement. If you are unsure, reach out to other Event organizers or other
members of SIG-Contributor Experience for their opinion on the matter.


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