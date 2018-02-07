# Kubernetes Github Organization Guide

The Kubernetes project leverages multiple GitHub organizations to store and
organize code. This guide contains the details on how to run those organizations
for CNCF compliance and for the guidelines of the community.

## Organization Naming

Kubernetes managed organizations should be in the form of `kubernetes-[thing]`.
For example, [kubernetes-client](https://github.com/kubernetes-client) where the
API clients are housed.

Prior to creating an organization please contact the steering committee for
direction and approval.

Note: The CNCF, as part of the Linux Foundation, holds the trademark on the
Kubernetes name. All GitHub organizations with Kubernetes in the name should be
managed by the Kubernetes project or use a different name.

## Transferring Outside Code Into A Kubernetes Organization

Due to licensing and CLA issues, prior to transferring software into a Kubernetes
managed organization there is some due diligence that needs to occur. Please
contact the steering committee and CNCF prior to moving any code in.

It is easier to start new code in a Kubernetes organization than it is to
transfer in existing code.

## Current Organizations In Use

The following organizations are currently known to be part of the Kubernetes
project:

* [kubernetes](https://github.com/kubernetes)
* [kubernetes-client](https://github.com/kubernetes-client)
* [kubernetes-csi](https://github.com/kubernetes-csi)
* [kubernetes-graveyard](https://github.com/kubernetes-graveyard) †
* [kubernetes-helm](https://github.com/kubernetes-helm)
* [kubernetes-incubator](https://github.com/kubernetes-incubator)
* [kubernetes-incubator-retired](https://github.com/kubernetes-incubator-retired)
* [kubernetes-retired](https://github.com/kubernetes-retired)
* [kubernetes-security](https://github.com/kubernetes-security)
* [kubernetes-sig-testing](https://github.com/kubernetes-sig-testing)
* [kubernetes-providers](https://github.com/kubernetes-providers)
* [kubernetes-addons](https://github.com/kubernetes-addons)
* [kubernetes-charts](https://github.com/kubernetes-charts)
* [kubernetes-extensions](https://github.com/kubernetes-extensions)
* [kubernetes-federation](https://github.com/kubernetes-federation)
* [kubernetes-sidecars](https://github.com/kubernetes-sidecars)
* [kubernetes-tools](https://github.com/kubernetes-tools)
* [kubernetes-test](https://github.com/kubernetes-test)

† kubernetes-retired should be used instead of kubernetes-graveyard going forward.

Note, this list is subject to change.

There are more organization names that we are squatting on with possible future
intentions. [For more details please see community issue #1407](https://github.com/kubernetes/community/issues/1407).

## Repository Guidance

Repositories have additional guidelines and requirements, such as the use of
CLA checking on all contributions. For more details on those please see the
[Kubernetes Template Project](https://github.com/kubernetes/kubernetes-template-project).
