# Conventions for Cloud Provider Repositories

This purpose of this document is to define a common structure for the cloud
provider repositories owned by current and future cloud provider SIGs[1]. In
accordance with the WG-Cloud-Provider Charter[2] to "define a set of common
expected behaviors across cloud providers", this proposal defines the location
and structure of commonly expected code.

As each provider can and will have additional features that go beyond expected
common code, this document is only prescriptive to the location of the
following code:

* Cloud Controller Manager implementations.
* Documentation.

This document may be amended with additional locations that relate to enabling
consistent upstream testing, independent storage drivers, and other code with
common integration hooks may be added

## Motivation

The development of the Cloud Controller Manager[3] and Cloud Provider
Interface[4] has enabled the provider SIGs to develop external providers that
capture the core functionality of the upstream providers. By defining the
expected locations and naming conventions of where the external provider code
is, we continue in creating a consistent experience for:

* Users of the providers, who will have easily understandable conventions for
  discovering and using all of the providers.
* SIG-Docs, who will have a common hook for building or linking to externally
  managed documentation
* SIG-Testing, who will be able to use common entry points for enabling
  provider-specific e2e testing.
* Future cloud provider authors, who will have a common framework and examples
  from which to build and share their code base.

## Requirements

Each cloud provider hosted within the `kubernetes` organization shall have a
single repository named `kubernetes/cloud-provider-<provider_name>`. Those
repositories shall have the following structure:

* A `cloud-controller-manager` subdirectory that contains the implementation
  of the provider-specific cloud controller.
* A `docs` subdirectory.
* A `docs/cloud-controller-manager.md` file that describes the options and
  usage of the cloud controller manager code.
* A `tests` subdirectory that contains testing code.

Additionally, the repository should have:

* A `docs/getting-started.md` file that describes the installation and basic
  operation of the cloud controller manager code.

Where the provider has additional capabilities, the repository should have
the following subdirectories that contain the common features:

* `dns` for DNS provider code.
* `cni` for the Container Network Interface (CNI) driver.
* `flex` for the Flex Volume driver.
* `installer` for custom installer code.

Each repository may have additional directories and files that are used for
additional feature that include but are not limited to:

* Other provider specific testing.
* Additional documentation, including examples and developer documentation.
* Dependencies on provider-hosted or other external code.

## Timeline

To facilitate community development, providers named in the `Make SIGs
responsible for implementations of CloudProvider` patch[1] can immediately
migrate their external provider work into their named repositories.

Each provider will work to implement the required structure during the
Kubernetes 1.11 development cycle, with conformance by the 1.11 release.

After the 1.11 release all current and new provider implementations must
conform with the requirements outlined in this document.

## References

1. [Makes SIGs responsible for implementations of `CloudProvider`](https://github.com/kubernetes/community/pull/1862)
2. [Cloud Provider Working Group Proposal](https://docs.google.com/document/d/1m4Kvnh_u_9cENEE9n1ifYowQEFSgiHnbw43urGJMB64/edit#)
3. [Cloud Controller Manager](https://github.com/kubernetes/kubernetes/tree/master/cmd/cloud-controller-manager)
4. [Cloud Provider Interface](https://github.com/kubernetes/kubernetes/blob/master/pkg/cloudprovider/cloud.go)
