# SIG Usability Charter

This charter adheres to the conventions described in the [Kubernetes Charter README] and uses
the Roles and Organization Management outlined in [sig-governance].

## Scope

The scope of SIG usability is the core end-user usability of the Kubernetes project. This covers
topics like user experience and accessibility. The goal of SIG Usability is to facilitate adoption of
the Kubernetes project by as diverse a community of end users as possible. We do this by ensuring that
each end user’s interaction with Kubernetes, from discovery to successful production use is seamless
and positive. Examples of efforts include user research, internationalization and accessibility.

### In scope

#### Code, Binaries and Services

- usability of end-user facing experiences (error messages, end-to-end tasks, etc)
- accessibility guidelines for Kubernetes community artifacts, examples include:
   - internationalization of documentation in binaries and documentation
   - color choices for people with color blind-ness
   - ensuring compatibility with screen reader technology
- user interface design for core components with user interfaces.
- code for performing internationaization and other accessbility enablement.
- translations - https://github.com/kubernetes/kubernetes/tree/master/translations
- internationalization - https://github.com/kubernetes/kubernetes/tree/master/pkg/kubectl/util/i18n

#### Repos
The SIG owns a `kubernetes-sigs/sig-usability` repo for the purpose of documenting:

- Common end user profiles (aka personas)
- Ongoing list of research studies and findings
- Ongoing list of known needs for user research (ex: top stackoverflow questions)
- Curated research on usability best practices (ex: HCI research literature on interface usability, error message best practices)
- Templates for anybody to be able to conduct user research interviews

#### Cross-cutting and Externally Facing Processes
SIG Usability will facilitate and collaborate with other SIGs on:
- giving end users channels for proving feedback on their experience
- surveys and other feedback gathering around project usability
- documenting common end user profiles (aka personas)
- documenting common end user interactions with Kubernetes (aka user journeys)
- design elements and language guidelines for Kubernetes concepts (e.g. 'reconciliation' or 'service')
- identifying opportunities for user research and best methods to achieve (qualitative and quantitative).
- tracking adherence to required guidelines such as [WCAG which is required to sell products in Europe].
- documentation in the vein of https://www.microsoft.com/en-us/accessibility/

### Out of scope

* Contributor user experience (covered by SIG-Contributor Experience)
* API Design (covered by SIG-Architecture)
* Command line tool infrastructure (covered by SIG-CLI)
* How-to guides for implementing accessible documentation (covered by SIG-Docs)

## Roles and Organization Management

This sig follows adheres to the Roles and Organization Management outlined in [sig-governance]
and opts-in to updates and modifications to [sig-governance].

### Additional responsibilities of Chairs
None

### Deviations from [sig-governance]

No Tech Leads (for now)

### Subproject Creation

Federation of Subprojects

[sig-governance]: https://github.com/kubernetes/community/blob/master/committee-steering/governance/sig-governance.md
[sig-subprojects]: https://github.com/kubernetes/community/blob/master/sig-YOURSIG/README.md#subprojects
[Kubernetes Charter README]: https://github.com/kubernetes/community/blob/master/committee-steering/governance/README.md
[WCAG which is required to sell products in Europe]: https://www.w3.org/WAI/news/2018-09-13/WCAG-21-EN301549/
