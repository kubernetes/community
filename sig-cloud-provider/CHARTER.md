# SIG Cloud Provider Charter

## Mission
The Cloud Provider SIG ensures that the Kubernetes ecosystem is evolving in a way that is neutral to all (public and private) cloud providers. It will be responsible for establishing standards and requirements that must be met by all providers to ensure optimal integration with Kubernetes.

## Subprojects & Areas of Focus

* Maintaining parts of the Kubernetes project that allows Kubernetes to integrate with the underlying provider. This includes but are not limited to:
    * [cloud provider interface](https://github.com/kubernetes/kubernetes/blob/master/pkg/cloudprovider/cloud.go)
    * [cloud-controller-manager](https://github.com/kubernetes/kubernetes/tree/master/cmd/cloud-controller-manager)
    * Deployment tooling which has historically resided under [cluster/](https://github.com/kubernetes/kubernetes/tree/release-1.11/cluster)
* Code ownership for all cloud providers that fall under the kubernetes organization and have opted to be subprojects of SIG Cloud Provider. Following the guidelines around subprojects we anticipate providers will have full autonomy to maintain their own repositories, however, official code ownership will still belong to SIG Cloud Provider.
    * [cloud-provider-azure](https://github.com/kubernetes/cloud-provider-azure)
    * [cloud-provider-gcp](https://github.com/kubernetes/cloud-provider-gcp)
    * [cloud-provider-openstack](https://github.com/kubernetes/cloud-provider-openstack)
    * [cloud-provider-vsphere](https://github.com/kubernetes/cloud-provider-vsphere)
* Standards for documentation that should be included by all providers.
* Defining processes/standards for E2E tests that should be reported by all providers
* Developing future functionality in Kubernetes to support use cases common to all providers while also allowing custom and pluggable implementations when required, some examples include but are not limited to:
    * Extendable node status’ and machine states based on provider
    * Extendable node address types based on provider
    * See also [Cloud Controller Manager KEP](https://github.com/kubernetes/community/blob/master/keps/0002-controller-manager.md)
* The collection of user experience reports from Kubernetes operators running on provider subprojects; and the delivery of roadmap information to SIG PM

## Organizational Management

* Six months after this charter is first ratified, it MUST be reviewed and re-approved by the SIG in order to evaluate the assumptions made in its initial drafting
* SIG meets bi-weekly on zoom with agenda in meeting notes.
    * SHOULD be facilitated by chairs unless delegated to specific Members
* The SIG MUST make a best effort to provide leadership opportunities to individuals who represent different races, national origins, ethnicities, genders, abilities, sexual preferences, ages, backgrounds, levels of educational achievement, and socioeconomic statuses

## Subproject Creation

Each Kubernetes provider will (eventually) be a subproject under SIG Cloud Provider. To add new sub projects (providers), SIG Cloud Provider will maintain an open list of requirements that must be satisfied.
The current requirements can be seen [here](https://github.com/kubernetes/community/blob/master/keps/0002-controller-manager.md#repository-requirements). Each provider subproject is entitled to create 1..N repositories related to cluster turn up or operation on their platform, subject to technical standards set by SIG Cloud Provider.
Creation of a repository SHOULD follow the KEP process to preserve the motivation for the repository and any additional instructions for how other SIGs (e.g SIG Documentation and SIG Release) should interact with the repository

Subprojects that fall under SIG Cloud Provider may also be features in Kubernetes that is requested or needed by all, or at least a large majority of providers. The creation process for these subprojects will follow the usual KEP process.

## Subproject Retirement

Subprojects representing Kubernetes providers may be retired given they do not satisfy requirements for more than 6 months. Final decisions for retirement should be supported by a majority of SIG members using [lazy consensus](http://communitymgt.wikia.com/wiki/Lazy_consensus). Once retired any code related to that provider will be archived into the kubernetes-retired organization.

Subprojects representing Kubernetes features may be retired at any point given a lack of development or a lack of demand. Final decisions for retirement should be supported by a majority of SIG members, ideally from every provider. Once retired, any code related to that subproject will be archived into the kubernetes-retired organization.


## Technical Processes
Subprojects (providers) of the SIG MUST use the following processes unless explicitly following alternatives they have defined.

* Proposals will be sent as [KEP](https://github.com/kubernetes/community/blob/master/keps/0000-kep-template.md) PRs, and published to the official group mailing list as an announcement
* Proposals, once submitted, SHOULD be placed on the next full meeting agenda
* Decisions within the scope of individual subprojects should be made by lazy consensus by subproject owners, with fallback to majority vote by subproject owners; if a decision can’t be made, it should be escalated to the SIG Chairs
* Issues impacting multiple subprojects in the SIG should be resolved by consensus of the owners of the involved subprojects; if a decision can’t be made, it should be escalated to the SIG Chairs

## Roles
The following roles are required for the SIG to function properly. In the event that any role is unfilled, the SIG will make a best effort to fill it. Any decisions reliant on a missing role will be postponed until the role is filled.


### Chairs
* 3 chairs are required
* Run operations and processes governing the SIG
* An initial set of chairs was established at the time the SIG was founded.
* Chairs MAY decide to step down at anytime and propose a replacement, who must be approved by all of the other chairs. This SHOULD be supported by a majority of SIG Members.
* Chairs MAY select additional chairs using lazy consensus amongst SIG Members.
* Chairs MUST remain active in the role and are automatically removed from the position if they are unresponsive for > 3 months and MAY be removed by consensus of the other Chairs and members if not proactively working with other Chairs to fulfill responsibilities.
* Chairs WILL be asked to step down if there is inappropriate behavior or code of conduct issues
* SIG Cloud Provider cannot have more than one chair from any one company.

### Subproject/Provider Owners
* There should be at least 1 representative per subproject/provider (though 3 is recommended to avoid deadlock) as specified in the OWNERS file of each cloud provider repository.
* MUST be an escalation point for technical discussions and decisions in the subproject/provider
* MUST set milestone priorities or delegate this responsibility
* MUST remain active in the role and are automatically removed from the position if they are unresponsive for > 3 months and MAY be removed by consensus of other subproject owners and Chairs if not proactively working with other Subproject Owners to fulfill responsibilities.
* MAY decide to step down at anytime and propose a replacement. This can be done by updating the OWNERS file for any subprojects.
* MAY select additional subproject owners by updating the OWNERs file.
* WILL be asked to step down if there is inappropriate behavior or code of conduct issues

### SIG Members

Approvers and reviewers in the OWNERS file of all subprojects under SIG Cloud Provider.

## Long Term Goals

The long term goal of SIG Cloud Provider is to promote a vendor neutral ecosystem for our community.  Vendors wanting to support Kubernetes should feel equally empowered to do so
as any of today’s existing cloud providers; but more importantly ensuring a high quality user experience across providers. The SIG will act as a central group for developing
the Kubernetes project in a way that ensures all providers share common privileges and responsibilities. Below are some concrete goals on how SIG Cloud Provider plans to accomplish this.

### Consolidating Existing Cloud SIGs

SIG Cloud Provider will aim to eventually consolidate existing cloud provider SIGs and have each provider instead form a subproject under it. The subprojects would drive the development of
individual providers and work closely with SIG Cloud Provider to ensure compatibility with Kubernetes. With this model, code ownership for new and existing providers will belong to SIG Cloud Provider,
limiting SIG sprawl as more providers support Kubernetes. Existing SIGs representing cloud providers are highly encouraged to opt-in as sub-projects under SIG Cloud Provider but are not required to do.
As a SIG opts-in, it will operate to ensure a smooth transition, typically over the course of 3 release cycles.

### Supporting New Cloud Providers

One of the primary goals of SIG Cloud Provider is to become an entrypoint for new providers wishing to support Kubernetes on their platform and ensuring technical excellence from each of those providers.
SIG Cloud Provider will accomplish this by maintaining documentation around how new providers can get started and managing the set of requirements that must be met to onboard them. In addition to
onboarding new providers, the entire lifecycle of providers would also fall under the responsibility of SIG Cloud Provider, which may involve clean up work if a provider decides to no longer support Kubernetes.

