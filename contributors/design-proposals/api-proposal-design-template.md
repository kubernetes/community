# Description of your API

Initial API version: v1alpha1

## Owners

SIG owner: list of SIG that will own the design, implementation, and support
Individual owners: list of individual GitHub users that will own the design, implementation, and support

## Abstract

4-8 sentence description of the problem and solution.  Bullet points welcome.

## Motivation

4-8 sentence of the current state and why it is insufficient.

## Use cases

Bullet point list of use cases this proposal is meant to address.

- As a cluster administrator...
- As a application developer...
- As a Kubernetes extension developer...

## Requirements

Bullet point list of new API requirements.  Requirements will in part be driven by use
cases, but may have additional engineering or interoperability requirements:

- User requirement
- Engineering requirement
- Interoperability requirement

## Overview subsequent versions

Subsequent versions will require their own API proposals, however it is important to include
the plans for them here to provide a broader picture.

### Beta

*Use cases*

- Use cases for transitioning to Beta

*Requirements*

- Additional requirements for transitioning to Beta

### GA

*Use cases*

- Use cases for transitioning to Beta

*Requirements*

- Additional requirements for transitioning to Beta

## Dependencies and deadlines

List of features that will depend on these changes.

List of features that these changes will depend upon.

**Include any deadlines that are related.** e.g. another feature depends on this and it is slated for 1.x.

## Proposed changes for new types or existing types

Declare the proposed Group / Version / Kind for any new types created as part of this proposal.

Declare any new fields for existing types and their version.  Will the fields be added as first class fields or
as annotations?

Describe how any new types or fields will be used by the end user.  Include sample yaml config and full
walkthrough examples of 1 or more usecases.

### Defaulting

Describe defaulting that will be done

### Validation

Describe validation the will be done

### Patching lists - strategic merge keys

For any list fields, describe whether they will be replaced when patched or merged.  If merged,
describe the merge key that will be used.

**mergeKeys must be unique and required for all elements in the list!**  Consequences of using a
non-unique or optional merge key may be severe.

### Controllers

Describe any controllers that will be added or updated.

### Subresources

Describe any subresources that will be added.

## Mandatory API pre-flight check list

The following items must be considered and explained when adding or changing and API.

### API convention deviations

List and justify any deviations from [API conventions].  *None* is acceptable if
there are no deviations.

### Impact on backwards / forwards compatibility

Describe any impact on backwards / forwards [compatibility].

Known future incompatibility risks and mitigation.

### Security, PII and authz

Describe any security concerns or PII that is managed by the API.  Describe
any special considerations that must be made for authentication or
authorization.

### Cli and client library considerations.

Describe any interactions with the cli or client libraries.

Is functionality in this API already present in the cli?  If so, was is the interaction and
path forward?

Is server side garbage collection needed / enabled?

## Alternatives considered

List alternative solutions to the use cases.  Include an analysis of any of
the following that apply.

- Client side solutions
- Existing APIs or solutions that are similar
- Existing workarounds - include solutions that compose existing APIs
- Using ThirdPartyResources

[API conventions]: https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md
[compatibility]: https://github.com/kubernetes/community/blob/master/contributors/devel/api_changes.md#on-compatibility