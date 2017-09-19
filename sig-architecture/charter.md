# SIG Architecture Charter

This charter is a WIP.

The Architecture SIG maintains and evolves the design principles of
Kubernetes, and provides a consistent body of expertise necessary to
ensure architectural consistency over time.

The scope covers issues that span all the system's components, how
they fit together, how they interact, etc.

Specific areas of focus include:

* Defining the scope of the Kubernetes project
  * [What is (and is not) Kubernetes](https://kubernetes.io/docs/concepts/overview/what-is-kubernetes/)
* Maintaining, evolving, and enforcing the deprecation policy
  * [Deprecation policy](https://kubernetes.io/docs/reference/deprecation-policy/)
* Documenting and evolving the system architecture
  * [Kubernetes Design and Architecture](../contributors/design-proposals/architecture/architecture.md)
* Defining and driving necessary extensibility points
* Establishing and documenting design principles
  * [Design principles](../contributors/design-proposals/architecture/principles.md)
* Establishing and documenting conventions for system and user-facing APIs
  * [API conventions](https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md)
* Developing necessary technical review processes, such as the proposal and API review processes
* Driving improvement of overall code organization, including github orgs and repositories
* Educating approvers/owners of other SIGs (e.g., by holding office hours)

Out of scope:
* Issues specific to a particular component or functional area, which would be the purview
  of some other SIG, except where they deviate from project-wide principles and conventions.
* [Release support policy](https://github.com/kubernetes/community/blob/master/contributors/design-proposals/release/versioning.md)
  is owned by SIG Release

TODO:
* Formalize decision processes
* Document initial reviewers and approvers
* Clarify criteria for areas out of scope for the SIG
* Document who owns client library, build, and release artifacts
* Document who owns conformance definition, profiles, etc.
