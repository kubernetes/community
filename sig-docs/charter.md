# SIG Docs Charter

This charter adheres to the conventions described in the [Kubernetes Charter README] and adheres to the Roles and Organization Management specified in [sig-governance].

## Scope

SIG Docs publishes Kubernetes documentation on kubernetes.io. Kubernetes documentation includes:
  - Documentation of the core Kubernetes APIs
  - Core Kubernetes architecture 
  - CLI tools shipped with the Kubernetes release

Responsibility for creating feature documentation belongs to the developers and SIGs creating each feature. This includes task-driven documentation for the feature itself and conceptual documentation about the feature.

SIG Docs sets standards for feature documentation, provides clear paths for docs contribution, and [coordinates documentation updates][docs-release] during quarterly releases.

SIG Docs maintains the Kubernetes website's infrastructure, tooling, and analytics.

SIG Docs is responsible for maintaining existing content and UX on the site.

SIG Docs creates subprojects as needed to handle specific aspects of Kubernetes documentation.

### In scope

SIG [readme]

#### Code, Binaries and Services

The kubernetes.io website, which includes:
* Site content and documentation
* [Content style guides][standards] and their application to feature documentation and release notes
* Processes for launching feature content and release notes during quarterly releases
* Site infrastructure (Hugo, Netlify) and tooling
* Site analytics (Google Analytics)
* Site styles and CSS
* Generated reference documentation for:
    - The core Kubernetes APIs
    - Kubernetes command-line tools, including but not limited to kubectl
    - Kubernetes setup tools
    - The Federation API

See Cross-cutting for limitations on how the SIG helps with related efforts on other projects.

#### Cross-cutting and Externally Facing Processes

- [Standards for content][standards]: style guide, contributor guide, PR reviews
  SIG Docs is not responsible for branding any content related to the Kubernetes organization. We are sometimes informally invited to consult on branding and design decisions, but we have no formal responsibility for such work.
  
- [Coordinating docs contributions for quarterly releases][docs-release]

- Kubernetes blog:
  The Kubernetes blog is a subproject of SIG Docs. SIG Docs provides tooling and workflow support for publishing the blog, but does not directly review blog posts or work with blog contributors. The blog subproject provides its own approvers and reviewers, and sets independent standards for creation and review of blog content. 

- Site UX:
  SIG Docs organizes and revises technical content to improve UX and technical accuracy for the current and four most previous releases, based on:
  - user-reported issues
  - site analytics
  - content audits performed by SIG members

- Other project documentation:
  SIG Docs is sometimes advises on Kubernetes components outside of the Kubernetes/Kubernetes GitHub repository, or on other Kubernetes subprojects, but is not responsible for the documentation of any of these components or projects.

### Out of scope

- SIG Docs is not responsible for creating new feature documentation.

    SIG Docs reviews and coordinates feature documentation created by community members.
    
    SIG Docs sets standards for content creation, in the form of a style guide; offers feedback and review of new feature documentation; offers advice about information architecture for new features in the docs; and otherwise provides guidance and oversight to make sure that new feature documentation is maximally helpful to developers.
    
- Branding, logos, or brand style guides for Kubernetes 

## Roles and Organization Management

SIG Docs adheres to the standards for roles and organization management as specified by [sig-governance]. This SIG opts in to updates and modifications to [sig-governance].

### Additional responsibilities of Chairs

Chairs also serve as Tech Leads.

### Deviations from [sig-governance]

Per [readme]: 

- Meetings are weekly
- Once per month, weekly meeting time changes for easier attendance from APAC contributors
- Once per quarter, leads and other interested parties meet to discuss quarterly goals and achievements

### Subproject Creation

SIG Chairs can create subprojects without requiring member votes.



[sig-governance]: https://github.com/kubernetes/community/blob/master/committee-steering/governance/sig-governance.md
[readme]: https://github.com/kubernetes/community/tree/master/sig-docs
[Kubernetes Charter README]: https://github.com/kubernetes/community/blob/master/committee-steering/governance/README.md
[standards]: https://kubernetes.io/docs/contribute/
[docs-release]: https://github.com/kubernetes/sig-release/tree/master/release-team/role-handbooks/docs