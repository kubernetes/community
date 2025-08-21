# WG API Expression Charter

This charter adheres to the conventions described in the [Kubernetes Charter README] and uses
the Roles and Organization Management outlined in [sig-governance].

## Scope

Enable declarative expression of Kubernetes APIs (metadata, validation rules, etc.) to replace imperative, handwritten logic. This working group focuses on developing frameworks and tools that allow API authors to express API rules declaratively as part of the schema itself, improving development velocity, consistency, and maintainability across both native Kubernetes types and CRDs.

### In scope

#### Code, Binaries and Services

- Declarative validation code generator (validation-gen) that generates Go validation code from declarative tags in API type definitions
- Kube API Linter (KAL) for checking API types against Kubernetes API Conventions and best practices
- Integration frameworks between declarative tools (validation-gen, kubebuilder, KAL, and other API tools)
- Libraries and components for type discovery and declarative API expression that can be reused across code generators
- Tools and processes for generating CEL or OpenAPI validation rules from declarative tags
- Documentation generation tools that create API documentation from declarative expressions

#### Cross-cutting and Externally Facing Processes

- Development and maintenance of KEPs for declarative API expression features (eg: KEP-5073 Declarative Validation)
- Migration strategies and guides for SIGs to adopt the declarative framework
- Establishing patterns for ratcheting, validation, defaulting, and immutability expressed declaratively
- Ensuring consistency between Kubernetes native types and CRD validation experiences
- Automated API quality checks and linting processes integrated into the standard API review workflow

### Out of scope

- General API design decisions unrelated to declarative expression
- Runtime validation logic that cannot be expressed declaratively

## Deliverables

The working group will deliver:

- **validation-gen code generator**: Core tool for generating Go validation code from declarative tags
- **Kube API Linter (KAL)**: Automated API convention checker integrated into the review process
- **Integration strategy**: Cohesive approach for integrating validation-gen, kubebuilder, KAL, and other tools
- **KEPs**: Series of enhancement proposals introducing and evolving the framework
- **CRD equivalence strategy**: Plan to bridge the gap between native types and CRDs
- **Documentation generation strategy**: Automated documentation from declarative API expressions
- **Migration guides**: Documentation for SIGs to adopt the framework and migrate existing APIs

## Stakeholders

- **SIG API Machinery**: Owner of API server, gengo, kube-openapi, and validation-gen tool
- **SIG Architecture**: Owner of Kubernetes API conventions and architectural decisions
- **SIG CLI**: Consumer of more expressive APIs, providing feedback on kubectl interaction and user experience

## Roles and Organization Management

This sig adheres to the Roles and Organization Management outlined in [sig-governance]
and opts-in to updates and modifications to [sig-governance].

### Meeting Mechanics

- **Frequency**: Bi-weekly (every 2 weeks)
- **Duration**: 60 minutes
- **Communication**:
  - Slack: #wg-api-expression
  - Mailing List: kubernetes-wg-api-expression@googlegroups.com

### Chairs

- Aaron Prindle (@aaron-prindle)
- Joel Speed (@JoelSpeed)

## Exit Criteria

The working group will have fulfilled its charter and can be disbanded when:

1. The core framework, including validation-gen and associated libraries, is stable and integrated into the main Kubernetes development workflow
2. Foundational KEPs (e.g., Declarative Validation) are implemented and graduated to GA
3. Kube API Linter (KAL) is mature, widely adopted, and integrated into the standard API development and review process
4. A clear, documented process exists for SIGs to migrate existing types and use the framework, with successful migrations of key types (eg: PodSpec) demonstrating the process
5. Ongoing maintenance of the framework has been successfully transitioned to owner SIGs, primarily SIG API Machinery

[sig-governance]: https://github.com/kubernetes/community/blob/master/committee-steering/governance/sig-governance.md
[sig-subprojects]: https://github.com/kubernetes/community/blob/master/sig-YOURSIG/README.md#subprojects
[Kubernetes Charter README]: https://github.com/kubernetes/community/blob/master/committee-steering/governance/README.md