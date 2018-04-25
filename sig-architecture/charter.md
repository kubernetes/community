# SIG Architecture Charter

## Mission

The Architecture SIG maintains and evolves the design principles of Kubernetes, and provides a consistent body of expertise necessary to ensure architectural consistency over time.

## Scope and subprojects

The scope of SIG Architecture covers issues that span all the system's components, how the code is organized, how the components fit together, how they interact, what common conventions they follow, general requirements and restrictions, overall technical scope of the project, etc. SIG Architecture is also responsible for educating members of other SIGs about these policies and principles. 

SIG Architecture is also a good place to identify gaps in the project, collect thinking on how the project should approach those problems and then work to motivate that those problems get addressed. This will often take the form of working to draw attention through an appropriate SIG.

Issues specific to a particular component or functional area, which would be the purview of some other single SIG, are out of scope for SIG Architecture, except where they deviate from project-wide principles and conventions, or are not yet being addressed, as discussed above.

SIG Architecture’s subprojects are as follows. [sigs.yaml](https://github.com/kubernetes/community/blob/master/sigs.yaml) contains links to their OWNERS.

<table>
  <tr>
    <td>Subproject</td>
    <td>Description</td>
    <td>Example Artifacts</td>
  </tr>
  <tr>
    <td>Kubernetes scope</td>
    <td>Defining "what is Kubernetes", approving which repos belong in the kubernetes GitHub org, evaluating, guiding, and curating new components to be added to Kubernetes releases where the implementation is consistent with the defined architectural standards</td>
    <td>What is (and is not) Kubernetes</td>
  </tr>
  <tr>
    <td>Conformance definition</td>
    <td>Reviewing, approving, and driving changes to the conformance test suite; reviewing, guiding, and creating new conformance profiles</td>
    <td>Test list, Test guidelines</td>
  </tr>
  <tr>
    <td>Architectural governance</td>
    <td>Establishing and documenting design principles, documenting and evolving the system architecture, reviewing, curating, and documenting new extension patterns</td>
    <td>Kubernetes Design and Architecture, Design principles</td>
  </tr>
  <tr>
    <td>API governance</td>
    <td>Establishing and documenting conventions for system and user-facing APIs, define and operate the APl review process, final API implementation consistency validation, co-own top-level API directories with API machinery</td>
    <td>API conventions, API linter</td>
  </tr>
  <tr>
    <td>Deprecation policy</td>
    <td>Maintaining, evolving, and enforcing the deprecation policy</td>
    <td>Deprecation policy</td>
  </tr>
  <tr>
    <td>Code organization</td>
    <td>Overall code organization, including github repositories and branching methodology, top-level and pkg OWNERS of kubernetes/kubernetes, kubernetes-template-project</td>
    <td>kubernetes-template-project</td>
  </tr>
  <tr>
    <td>KEP</td>
    <td>Develop and drive technical enhancement review process</td>
    <td>Process, Template</td>
  </tr>
</table>


## Roles

The following roles are required for the SIG to properly function. In the event that any role is unfilled, the SIG will make a best effort to fill it, and any decisions reliant on a missing role will be postponed until the role is filled.

* **Chair**

    * Number: 2-3

    * Run operations and processes governing the SIG

    * An initial set of chairs was established at the time the SIG was founded.

    * Chairs *MAY* decide to step down at anytime and propose a replacement, who must be approved by all of the other chairs. This *SHOULD* be supported by a majority of SIG Members.

    * Chairs *MAY* select additional chairs by consensus. This *SHOULD* be supported by a majority of SIG Members.

    * Chairs *MUST* remain active in the role and are automatically removed from the position if they are unresponsive for > 3 months and *MAY* be removed by consensus of the other Chairs and all of the Technical leads if not proactively working with other Chairs to fulfill responsibilities.

    * A majority of chairs cannot be from a single company. If a chair’s change of company affiliation causes this threshold to be exceeded, one of the chairs must step down and propose a replacement, as per the procedure above.

    * Defined in[ sigs.yaml](https://github.com/kubernetes/community/blob/master/sigs.yaml#L1454)

* SIG Technical Leads

    * Number: 3

    * Establish new subprojects, and retire existing subprojects

    * Resolve cross-subproject technical issues and decisions, and escalations from subprojects

    * Decision-making MUST be by consensus. It’s particularly important for the technical leads to provide cohesive technical guidance for the project as a whole.

    * The initial set of technical leads is a subset of the long-standing group of API approvers: Clayton Coleman, Tim Hockin, and Brian Grant

    * Technical leads MUST have a demonstrated mastery, proficiency, and review history in the project sufficient to assess their ability to arbitrate highly technical and project-wide decisions

    * Technical leads *MUST* remain active in the role and are automatically removed from the position if they are unresponsive for > 3 months

    * Technical leads *MAY* decide to step down at anytime and propose a replacement, who must be approved by all of the other technical leads

    * All technical leads cannot be from a single company

* Subproject Owners

    * Scoped to a subproject defined in[ sigs.yaml](https://github.com/kubernetes/community/blob/master/sigs.yaml#L1454)

    * The initial owners *SHOULD* be established at subproject founding from relevant OWNERS files wherever possible

    * *MUST* be an escalation point for technical discussions and decisions in the subproject

    * *MUST* set milestone priorities or delegate this responsibility

    * *MUST* remain active in the role and are automatically removed from the position if they are unresponsive for > 3 months and *MAY* be removed by consensus of other subproject owners and all of the Technical leads if not proactively working with other Subproject Owners to fulfill responsibilities.

    * *MAY* decide to step down at anytime and propose a replacement. Use[ lazy-consensus](http://communitymgt.wikia.com/wiki/Lazy_consensus) amongst subproject owners with fallback on majority vote to accept proposal. This *SHOULD* be supported by a majority of subproject members (those having some role in the subproject).

    * *MAY* select additional subproject owners through a[ super-majority](https://en.wikipedia.org/wiki/Supermajority#Two-thirds_vote) vote amongst subproject owners. This *SHOULD* be supported by a majority of subproject contributors (through[ lazy-consensus](http://communitymgt.wikia.com/wiki/Lazy_consensus) with fallback on voting).

    * Number: at least 2

    * Defined in[ sigs.yaml](https://github.com/kubernetes/community/blob/master/sigs.yaml#L1454)[ OWNERS](https://github.com/kubernetes/community/blob/master/committee-steering/governance/contributors/devel/owners.md) files

* Members

    * Currently is comprised of reviewers and approvers in[ OWNERS](https://github.com/kubernetes/community/blob/master/committee-steering/governance/contributors/devel/owners.md) files for subprojects

## Organizational management

* Six months after this charter is first ratified, it MUST be reviewed and re-approved by the SIG in order to evaluate the assumptions made in its initial drafting

* SIG meets bi-weekly on zoom with agenda in meeting notes

    * *SHOULD* be facilitated by chairs unless delegated to specific Members

* SIG holds office hours bi-weekly on zoom with agenda in meeting notes

    * *SHOULD* be facilitated by chairs unless delegated to specific Members

* TODO: The SIG MUST make a best effort to provide leadership opportunities to individuals who represent different races, national origins, ethnicities, genders, abilities, sexual preferences, ages, backgrounds, levels of educational achievement, and socioeconomic statuses

### **Project management**

#### **Subproject creation**

The initial set of subprojects owned by the SIG is defined above.

* New subprojects MUST be created by[ KEP](https://github.com/kubernetes/community/blob/master/keps/0000-kep-template.md) proposal and approved by[ ](http://communitymgt.wikia.com/wiki/Lazy_consensus)consensus of SIG Technical Leads. The result *SHOULD* be supported by the majority of SIG members.

    * KEP *MUST* establish subproject owners

    * [sigs.yaml](https://github.com/kubernetes/community/blob/master/sigs.yaml#L1454) *MUST* be updated to include subproject information and[ OWNERS](https://github.com/kubernetes/community/blob/master/committee-steering/governance/contributors/devel/owners.md) files with subproject owners

#### **Subproject retirement**

Subprojects may be retired, where they are archived to the GitHub kubernetes-retired organization, when they are no longer supported based on the following criteria.

* A subproject is no longer supported when there are no active owners with activity on the project for the following time:

    * Subprojects with no known users can be retired after being unsupported for > 3 months

    * Subprojects with known users may be retired after providing at least 6 months notification of retirement

* Use consensus amongst SIG Technical Leads to decide to retire. This *SHOULD* be supported by a majority of SIG Members.

### **Technical processes**

Subprojects of the SIG *MUST* use the following processes unless explicitly following alternatives they have defined.

* Proposals will be sent as[ KEP](https://github.com/kubernetes/community/blob/master/keps/0000-kep-template.md) PRs, and published to the official group mailing list as an announcement

* Proposals, once submitted, SHOULD be placed on the next full meeting agenda

* Decisions within the scope of individual subprojects should be made by lazy consensus by subproject owners, with fallback to majority vote by subproject owners; if a decision can’t be made, it should be escalated to the SIG Technical leads

* Issues impacting multiple subprojects in the SIG should be resolved by consensus of the owners of the involved subprojects; if a decision can’t be made, it should be escalated to the SIG Technical leads
