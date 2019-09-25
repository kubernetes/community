# Kubernetes API Review Process

## Contact
* [Mailing List](mailto:kubernetes-api-reviewers@googlegroups.com)
* [Slack](https://kubernetes.slack.com/messages/api-review)

# Process Overview and Motivations

To preserve usability and consistency in Kubernetes core APIs, changes and additions require oversight.
The API review process is intended to maintain logical and functional integrity of the API over time,
the consistency of user experience and the ability of previously written tools to function with new APIs.
Wherever possible, the API review process should help change submitters follow [established conventions](/contributors/devel/sig-architecture/api-conventions.md),
and not simply reject without cause. 

Because reviewer bandwidth is limited, the process provides a prioritized backlog.
While this means some changes may be delayed in favor of other higher priority ones,
this will help maintain critical project velocity, transparency, and equilibrium.
Ideally, those whose API review priority is shifted in a release-impacting way will be proactively notified by the reviewers.

# Goals of the process

* Provide a transparent, easily-navigable process so all parties understand their roles, responsibilities, and expectations

* Protect Kubernetes APIs from disruptive, inconsistent, or destabilizing changes

* Optimize the experience of users consuming our APIs, balanced with the effort
  required from API authors.

* Respect, maximize, and expand reviewer bandwidth

* Integration with the regular review process, adding as little API-review-specific overhead as possible

# What APIs need to be reviewed?

## Mandatory
The changer is expected to submit the changes to API review, and the change is blocked on the API reviewer approval.  As a general rule of thumb, anything that is considered to be part of "core Kubernetes" requires a mandatory API review.  The list below outlines what are considered Kubernetes core APIs:

* All API implementations (including alpha versions) that are part of core Kubernetes must be reviewed, including CRDs, so user experience across the Kubernetes ecosystem is consistent.

* github.com/kubernetes/kubernetes, and any kubernetes-* org projects it depends on are required to be reviewed.

* Kubernetes-style API that uses a *.k8s.io or *.kubernetes.io name, e.g. "storage.k8s.io".

* "Critical" or other “highly-integrated” APIs, such as our extension points in the node, apiserver, and controllers.  This includes CSI, CNI, CRI, and CPI.

## Voluntary
Voluntary reviews apply towards non-core APIs that do not meet the [mandatory](#mandatory) requirements listed above.  Changer can request a review, or a 3rd party can nominate the API for review.  The review comments are considered recommendations.    

* SIG sponsored CRD based APIs outside of the core that use the "*.x-k8s.io" namespace.

* SIG sponsored subprojects that produce APIs (including CRDs) outside of *.k8s.io or *.kubernetes.io API groups, and are intended to work with kubectl and/or kube-apiserver. (intent is to ensure consistent user experience across the Kubernetes ecosystem)

## What parts of a PR are "API changes"?

* "Resource APIs" include the versioned data definition (pkg/apis/*/v*/types.go or OpenAPI for CRDs), validation (pkg/apis/*/validation.go or OpenAPI for CRDs).

* Configuration files, flags, and command line arguments are all part of our user and script facing APIs and must be reviewed.

* Compiled-in APIs of the kube-apiserver

* Webhooks request/response formats in kube-apiserver

* HTTP APIs in kubelet

* plugins which are not covered by some other standards effort (e.g. CSI and CNI APIs would be deferred to those standards bodies)

# Mechanics

0. Requesters should complete the pre-review checklist:
    * The goal of the proposed change has agreement from the appropriate sub-project owners and/or SIG leads
    * A [KEP](https://github.com/kubernetes/enhancements/blob/master/keps/0001-kubernetes-enhancement-proposal-process.md) and tracking issue in [kubernetes/enhancements](https://github.com/kubernetes/enhancements/) has been created for changes within the kubernetes-* org introducing:
        * Any new resource type
        * Any new version of a stable API
        * Any new functionality added to a stable API as defined by SIG Architecture and the API Reviewers
        * Any change to the meaning, validation, or behavior of a field
    * The existing [API conventions](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md) (and [API change guidelines](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api_changes.md), if applicable) have been read and followed.

1. Request an API review for a PR or issue in the kubernetes org by adding the `api-review` label with a `/label api-review` comment (requests can be cancelled with a `/remove-label api-review` comment)
    * If this is a review of a PR implementing an already-reviewed design/KEP, reference the approved KEP and note any differences between the approved design and the implementation.

2. API reviews are tracked in a project board at https://github.com/orgs/kubernetes/projects/13
    * Github query for requested reviews not yet in the project:
        * [`is:open org:kubernetes label:api-review -project:kubernetes/13`](https://github.com/search?q=is%3Aopen+org%3Akubernetes+label%3Aapi-review+-project%3Akubernetes%2F13)
    * Github query for items in the project no longer requesting review:
        * [`is:open org:kubernetes -label:api-review project:kubernetes/13`](https://github.com/search?q=is%3Aopen+org%3Akubernetes+-label%3Aapi-review+project%3Akubernetes%2F13)
    * Requests are triaged by API approvers/reviewers/moderators [regularly](#review-lifecycle-timing), and added to the project board if prereqs have been completed
    * As requests are added to the project board, that is reflected in the sidebar of the issue or PR, along with the current status (backlog, assigned, in progress, completed)
    * The API review backlog and queue is publicly visible at https://github.com/orgs/kubernetes/projects/13

3. Backlog
    * An approver or moderator will adjust the prioritization of the issue in the backlog. Reviews are prioritized based on a number of factors:
    * Whether the change is targeting the current milestone, a specific future milestone, or is untargeted
    * The maturity level of the change (generally, GA > beta > alpha changes)
    * Feedback from SIG leads / subproject leads on relative priorities
    * Whether this is an initial review or re-review (or a review of a PR implementing an already-reviewed API in a KEP)
    * Size/complexity of the change

4. Assigned
    * An approver or moderator will assign an approver (or potentially aspiring reviewer - see *training reviews* below for the aspiring reviewer workflow)
    * Reviews are assigned based on reviewer capacity and domain knowledge
    * Assignment of reviewers is done on the issue/PR itself using the normal `/assign` method (works seamlessly with existing github/PR dashboard queries)
    * All API reviews assigned to an individual can be viewed in the project board ([example](https://github.com/orgs/kubernetes/projects/13/?card_filter_query=assignee%3Aliggitt)), for visibility on status, order, and reviewer load

5. In Progress / Approved / Changes Requested / Rejected
    * Reviews proceed like a normal KEP or PR review. Possible outcomes:
    * Approval:
      * Implementation PRs are tagged with `/lgtm /approve` and merged normally
      * KEP PRs containing API designs can also be tagged with `/lgtm /approve`, but should explicitly note if API approval is being given. This approval should be linked to when later requesting review of the implementation PR, and should limit the scope of the implementation review to differences between the approved design and final implementation, problems encountered during implementation, and correctness of the implementation.
      * The approved issue is archived in the review project board, and the `api-review` label is removed.
    * Changes requested:
      * Comments or questions are left on the PR or issue, and the reviewer notifies the submitter
      * The reviewer moves the issue to "Changes Requested" in the review project board
      * Once the requested changes are made, or questions are answered, the submitter notifies the reviewer, who moves the issue back to "In Progress"
      * To the degree possible, complete sets of comments/changes should be requested and made, to avoid excessive back-and-forth cycles.
    * Rejected:
      * If completely rejected, e.g. "please do this work outside the Kubernetes org" - an explanation of why the change was rejected - appeals can be requested from the api-approvers mailing list ([kubernetes-api-reviewers@googlegroups.com](mailto:kubernetes-api-reviewers@googlegroups.com)) where the moderator will coordinate a follow-up review.  If that request results in another rejection, there is no further appeal.
      * The rejected issue is archived in the review project board, and the `api-review` label is removed.

## Review lifecycle timing

Ideally, reviews will happen as quickly as possible, but it is highly dependent on reviewer availability and bandwidth. In general, the following timeframe can be expected:

- time t: request review
- time t+1 week: prioritized and queued
- time t+3 weeks: first review complete
- time t+4 weeks: subsequent review complete
- time t+6 weeks: approved or rejected

Timing of API review requests matters.  The larger the change the more time that must be afforded.  New API resources (aka Kinds) may require significantly more thought than single field additions.  API reviews that are requested too late in a release cycle may not complete in time to make the release.  Plan ahead. Also, if you are changing an approved API, you must consult with the [kubernetes-api-reviewers@googlegroups.com](mailto:kubernetes-api-reviewers@googlegroups.com) list to ensure it is still consistent with the approvals already granted. From a process perspective, you would request a new review in that case.

## The Moderator Role

The moderator role is staffed by SIG Architecture and manages the API review backlog on behalf of the reviewer team. They will ensure that reviews are finished within a reasonable time, that information is correct, and that appropriate state labels are applied. They may also help prioritize the backlog, or move cards across the project board. Their mission is to help reviewers spend the majority of their efforts on performing reviews, not doing process administrivia. They may also work with the review team to schedule face-to-face review sessions as needed, or ensure the review is added to the SIG-Architecture meeting agenda.

## Expanding the Reviewer and Approver Pool

There are two levels of authority granted in this process. The reviewer and approver. Reviewers have the expertise to fully assess and make recommendations such that minimal extra effort is required on the part of the approver. Approvers are vested with final decision-making power for the request, and can only be appealed in the manner stated above. 

To become a formal reviewer, you must gain a high level of proficiency in understanding the API conventions, project structure, and the underlying architecture. The best path is to study the work of other aspiring reviewers and approvers. Once familiar with the core concepts, one should contact the moderator to be matched with a reviewer mentor who can help facilitate their eventual inclusion as a formally-recognized reviewer. All issues in the API Reviews backlog can receive an initial pass by a non-approver. These trial run reviews can then be examined and critiqued by other reviewers or approvers. 

The path from reviewer to approver requires significant reviewer experience, endorsement from their mentor, and a body of successfully-completed training reviews.  Aspiring approvers can petition for inclusion in the OWNERS file by sending an email to the SIG-Architecture and Kubernetes-api-reviewers mailing lists with the subject "**Request for API ****Approver**** Status**" with links to all of their completed reviews for the prior six month period. If all of the current approvers LGTM the issue, then the submitter may open a PR against the OWNERS file with their name added.  

To qualify, aspiring approvers should reference at least 5 provisional reviews of new API types and 10 provisional reviews of changes to existing API types with no substantive disagreement or additions by approver. Aspiring approvers must be able to demonstrate competence both for new API review and for compatibility concerns with modifications to existing APIs.

An indication of whether a provisional review was seen as qualifying should be part of the critique/coaching feedback given ("your provisional review caught everything important that I saw, +1")

### Training Reviews

For those wishing to become formal reviewers, the process has an alternate path where API reviews that are not time-sensitive can be initially reviewed, and then collaboratively examined with an approver/mentor. This can be accomplished asynchronously as comments in the review document, or ideally by pairing on the approver examination. 

Once the provisional review has been accepted by the approver, it then goes to the submitters just as it would in the normal process. 

As aspiring reviewers gain proficiency, they may pair with an approver on time-sensitive reviews to better understand the process with less impact on the review duration. 

Aspiring reviewers should reach out the moderator on slack.  The moderator will add them to the list of aspiring reviewers and assign training reviews as they become available. 

## Aspirations for future iterations of this process

* Make a video on this process as guidance for contributors

* Review the API conventions doc for freshness and organization

* Separate and keep separate the process and the API conventions parts of all our docs.  That way, people who are not subject to our process can read a more succinct set of docs on how to make good APIs.

* Consider a new home for our API docs that is prettier than Github-flavored markdown ( a new tab on kubernetes.io or a new site) 

* Given the level of completeness of the conventions this process should generate documentation of undocumented conventions, which might be pre-existing, or might be new.

* Ideally this process should lend itself to some automation. It currently does not. [ TBD: Create an umbrella issue for this ]

    * [We already tag PRs with api-change](https://github.com/kubernetes/test-infra/blob/1611f3be80713f9df933e25e4cf9a2b538302ef5/mungegithub/misc-mungers/deployment/kubernetes/path-label.txt#L7-L17)

        * We should review whether this can be improved

        * We should look at tabulating these changes at the end of the release

    * We should consider autogenerating API changes docs from openapi, or at least making it easy to review the set of changes

    * We should automate mechanical API review checks

        * For example, this PR: [https://github.com/kubernetes/kubernetes/pull/54887](https://github.com/kubernetes/kubernetes/pull/54887) 

* We may want to move toward the idea and process for "shadow reviewers" outlined in reference [1]

* We need to update other docs to point to this document

* [Updating-docs-for-feature-changes.md](https://git.k8s.io/sig-release/release-team/role-handbooks/documentation-guides/updating-docs-for-feature-changes.md#when-making-api-changes)

* [Api_changes.md](/contributors/devel/sig-architecture/api_changes.md)

* [Api-conventions.md](/contributors/devel/sig-architecture/api-conventions.md)

* [Pull-requests.md](https://github.com/kubernetes/community/blob/a74d906f0121c78114d79a3ac105aa2d36e24b57/contributors/devel/pull-requests.md#2-smaller-is-better-small-commits-small-prs) - should be updated to specifically call out API changes as important

* [Owners.md](https://github.com/kubernetes/community/blob/a74d906f0121c78114d79a3ac105aa2d36e24b57/contributors/devel/owners.md#code-review-process) - should be updated

* We should probably add a checkpoint in the release process that covers focused API review

* We should probably re-review all API changes in a release as part of documentation creation as a "last-chance" guarantee of ensuring we don’t ship something we don’t want to support

* We need a special process for "informational" or non-binding reviews

* [https://docs.google.com/document/d/1OkSQngGem7xaENqaO8jzHLDSSIGh2obPUaJGAFDwTUE/edit?disco=AAAAB_-Yzjw](https://docs.google.com/document/d/1OkSQngGem7xaENqaO8jzHLDSSIGh2obPUaJGAFDwTUE/edit?disco=AAAAB_-Yzjw) 

## References

[1] [https://storage.googleapis.com/pub-tools-public-publication-data/pdf/45294.pdf](https://storage.googleapis.com/pub-tools-public-publication-data/pdf/45294.pdf) 

[2] [http://wiki.netbeans.org/APIReviews](http://wiki.netbeans.org/APIReviews) 

[3] [https://cwiki.apache.org/confluence/display/TS/API+Review+Process](https://cwiki.apache.org/confluence/display/TS/API+Review+Process) 

[4] [https://docs.google.com/document/d/135PQSFTDqUmxsBhG7ZwMYtfWnJ3XTirsdAvytlkzOd0/](https://docs.google.com/document/d/135PQSFTDqUmxsBhG7ZwMYtfWnJ3XTirsdAvytlkzOd0/) 

[5] [https://github.com/kubernetes/community/pull/419](https://github.com/kubernetes/community/pull/419) 

