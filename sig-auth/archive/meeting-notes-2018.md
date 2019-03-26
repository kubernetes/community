# Kubernetes Sig-auth Meeting Agenda

## December 26, 2018, 11a - Noon (Pacific Time)



*   Canceled for holidays


## December 12, 2018, 11a - Noon (Pacific Time)



*   Canceled in lieu of kubecon seattle SIG events:
*   [State of Kubernetes Security: Mon, December 10, 10:00am – 10:30am](http://bit.ly/kubernetes-summit) (contributors summit)
*   [SIG Auth intro: Tuesday, December 11, 10:50am-11:25am](https://sched.co/Grd0)
*   [SIG Auth deep dive: Thursday, December 13, 10:50am-11:25am](https://sched.co/Grez)


## November 28, 2018, 11a - Noon (Pacific Time)



*   Recording
*   Announcements
    *   [1.13 release notes draft](https://github.com/kubernetes/sig-release/pull/386) open for review ([preview](https://github.com/marpaia/sig-release/blob/1.13-release-notes/releases/release-1.13/release-notes-draft.md)), any comments on SIG Auth items welcome (search for "SIG Auth")
*   Demos
    *   [kube-rbac-proxy](https://github.com/brancz/kube-rbac-proxy)
*   Pulls of note
    *
*   Designs of note
    *
*   Discussion topics
    *   sig-auth ownership of [kube-rbac-proxy](https://github.com/brancz/kube-rbac-proxy)
    *   [@pbarker] Outgoing webhook auth https://docs.google.com/document/d/1rtnPnrW2Ws1I8h826oYwSKuqWif5SaRnm2jvwFH9lIk/edit?usp=sharing
    *   [@cjcullen,@tallclair] Request for feedback on contributor summit "State of Security" outline: https://docs.google.com/document/d/1m5tPasZJ_7IZDw_i-iQsEe7bNMa1VGlytHcI_I9jMOg/edit#
    *   [Clayton] Something about bootstrapping
        *   Unblocking static pods to run prior to establishing kubelet/API connection
        *   https://github.com/kubernetes/kubernetes/pull/71174 ?
*   Action items
    *   @brancz will make issue for review of kube-rbac-proxy


## November 14, 2018, 11a - Noon (Pacific Time)



*   Canceled due to kubecon china and code freeze


## October 31, 2018, 11a - Noon (Pacific Time)



*   [Recording](https://youtu.be/xfj8wG4aG_0)
*   Announcements
    *   Code freeze November 15th (middle of kubecon China)
*   Demos
    *   Nikhil Bhatia @ Microsoft - Demo for Kubernetes Policy Controller -- https://github.com/Azure/kubernetes-policy-controller
*   Pulls of note
    *   .
*   Designs of note
    *   .
*   Discussion topics
    *   [klizhentas] Discuss CSR API promotion out of beta https://github.com/kubernetes/kubernetes/issues/69836
        *   API shape/issues
            *   Requires requesters to know all the info about the end certificate
            *   Use for higher-level requests ("give me a node client cert", "give me a serving cert for a pod backing a service")
        *   Approval flow/issues
            *   Cannot limit or add components to request (limit or add SANs, usages, etc)
        *   Signing flow/issues
            *   Method for multiple signers to interact (or approver to indicate what signer should be used)
        *   Guarantees on issued certificates
            *   No (current) guarantee all requested extensions/SANs are issued
            *   No (current) guarantee issued client certificates will be accepted as API client certs
    *   [liggitt] [node self-labeling resolution](https://github.com/kubernetes/community/pull/911#issuecomment-434757411)
        *   NodeRestriction limits self-labeling under kubernetes.io and k8s.io
    *   [mikedanese] moving service account token secret mounts to projected volumes
        *   Breaks in the presence of PSPs that only allow secret volumes
        *   Ugly options:
            *   Change PSP to allow projected volumes if secret volumes are allowed
            *   …
        *   Discussion in [#69848](https://github.com/kubernetes/kubernetes/pull/69848#issuecomment-434158411)
*   Action items
    *   [Mike] CSR ?


## October 17, 2018, 11a - Noon (Pacific Time)



*   [Recording](https://youtu.be/-t680vY95NU)
*   Announcements
    *   s/@liggitt/@enj/ in sig-auth chairs ([thread](https://groups.google.com/forum/?utm_medium=email&utm_source=footer#!msg/kubernetes-sig-auth/kqhoYhrzYLc/HBCOWEYKBQAJ), [PR](https://github.com/kubernetes/community/pull/2749))
*   Pulls of note
    *   NodeRestriction prevents node label updates ([#68267](https://github.com/kubernetes/kubernetes/pull/68267))
        *   Preference for well-known labels and prefixes (zero-config)
            *   Will try to reconcile with proposed approach for CSI labels
        *   Will start thread with sig-node/storage/auth
        *   Possible that adding arbitrary labels on node create (kubelet --node-labels flag) could still be allowed
            *   Depends on nodes not being able to delete themselves
            *   Assumes a new node isn't compromised
            *   Alternative could be to taint nodes that register with unknown labels
            *   Might be useful to survey use of kubelet --node-labels flag
    *   authenticator interface change ([#69582](https://github.com/kubernetes/kubernetes/pull/69582))
        *   Plumbs context to allow cancelation of backend webhook token authn
        *   Add Response which includes Audience (expected audience of authn, needed for token review of non-api server audiences) and user.Info
        *   Old servers will ignore audience and hence not validate (so client side validation will be required for some time)
    *   Migrate service account volume to a projected volume ([#69848](https://github.com/kubernetes/kubernetes/pull/69848))
        *   Requires kubelets 1.6 or newer
        *   Concerns around the implicit API for SA token secret (contains both the token and the CA data, but we would like not to include the CA data and instead have a config map in the namespace/combining the data using a projected volume)
        *   AI: communicate that users that want to obtain service account credentials for use outside of pods should do so explicitly [with a dedicated secret](https://kubernetes.io/docs/reference/access-authn-authz/service-accounts-admin/#to-create-additional-api-tokens)
        *   AI: put together roadmap for:
            *   making it possible to stop auto-generating token secrets
                *   measure impact to kube/e2e/conformance as a sample of the issued users will encounter
            *   timeline for enabling that option by default
            *   timeline for removing the option to auto-generate token secrets
*   Designs of note
    *   [KEP: authority delegation](https://github.com/kubernetes/community/pull/2794)
*   Discussion topics
    *   [Seeding TL thread](https://groups.google.com/forum/#!topic/kubernetes-sig-auth/v0-Awf0rBOg)
    *   Need sub projects to merge, update owner files, and pick initial technical leads based on owner files
    *   Subproject PR - [#2525](https://github.com/kubernetes/community/pull/2525)


## October 3, 2018, 11a - Noon (Pacific Time)



*   [Recording](https://youtu.be/9EA16cWHY7g)
*   Announcements
    *   [liggitt] Kubernetes Third-Party Security Audit Team Formation ([link](https://github.com/kubernetes/sig-release/issues/284#issuecomment-426384132))
    *   [liggitt] Proposed sig-auth chair change ([thread](https://groups.google.com/forum/#!topic/kubernetes-sig-auth/kqhoYhrzYLc))
    *   [aishs] Update on k8s 1.13 timeline and discuss feature load

    <span style="text-decoration:underline;">Features in 1.13 (from 1.12)</span>

        *   1. [600: Dynamic Audit Configuration](https://github.com/kubernetes/features/issues/600) - pbarker
        *   2. [22: API Audit Logging](https://github.com/kubernetes/features/issues/22) - tallclair
        *   3. [542: TokenRequest API and Kubelet integration](https://github.com/kubernetes/features/issues/542) - mikedanese
*   Pulls of note
    *
*   Discussion topics
    *   [pbarker] discuss changes to dynamic audit api https://github.com/kubernetes/community/pull/2738
    *   [enj] Should we allow admins/editors in a namespace to create events in that namespace (change default RBAC policy)?
        *   Audit is the system of record now (so the creation of fake events is less dangerous / they are not trusted sources of cluster information)
        *   Events are still useful for debugging
            *   Namespace scoped controller can use them to report failures
                *   Makes it easier to test on a cluster with limited access
            *   CI system could emit events to show progress or failures
        *   Users could falsely trust events made by a malicious party
    *


## September 19, 2018, 11a - Noon (Pacific Time)



*   [Recording](https://youtu.be/Ayfx3xTFmU4)
*   Announcements
    *   v1.13 development timeline
    *   Subprojects:
        *   https://github.com/kubernetes/community/pull/2525
*   Discussion topics
    *   Support an image pull credential flow built on bound service account tokens
        *   [kubernetes/kubernetes#68810](https://github.com/kubernetes/kubernetes/issues/68810)
    *   Bug bounty scope
        *   https://github.com/kubernetes/community/pull/2620


## September 5, 2018, 11a - Noon (Pacific Time)



*   Cancelled, no agenda items, release is closing out


## August 22, 2018, 11a - Noon (Pacific Time)



*   [Recording](https://youtu.be/VuxJwopD4nE)
*   Announcements
    *   code slush 8/28, code freeze 9/4
    *   subprojects enumeration PR in progress - [#2525](https://github.com/kubernetes/community/pull/2525)
*   Pulls of note
    *   [@pragashj] present and get feedback on [SAFE](https://github.com/cn-security/safe/blob/master/README.md), [WG proposal](https://github.com/cncf/toc/pull/146/files)
        *   co-chairs: Sarah Allen, Dan Shaw joined meeting to present and answer questions
        *   Tim Allclair has volunteered to act as liaison between the two groups
    *   [wteiken] Change to image review API: https://github.com/kubernetes/kubernetes/pull/64597
    *   [@pbarker] Dynamic audit PRs are up https://github.com/kubernetes/kubernetes/pull/67547

        https://github.com/kubernetes/kubernetes/pull/67257


## August 08, 2018, 11a - Noon (Pacific Time)



*   [Recording](https://www.youtube.com/watch?v=dTzDsrtjyT0)
*   Announcements
    *   HashiCorp Vault KMS provider - [repo](https://github.com/oracle/kubernetes-vault-kms-plugin). Feedback welcome.
    *   Audit event type promotion to v1 - [#65891](https://github.com/kubernetes/kubernetes/pull/65891)
*   Pulls of note
    *   New sig-auth charter: https://github.com/kubernetes/community/blob/master/sig-auth/charter.md
        *   Feedback welcome
        *   Follow up: listing subprojects & related responsibilities
    *   Remove basic auditing: https://github.com/kubernetes/kubernetes/pull/65862
        *   Part of advanced auditing going to GA
        *   Advanced auditing is not 100% compatible with the basic auditing log format
*   Discussion
    *   Container-level service account: [kubernetes/community#2497](https://github.com/kubernetes/community/pull/2497)
        *   Concern: pod is the security boundary, separating permissions at the container level could create some unexpected attack vectors
        *   Useful for auditing
    *   Kubelet TLS Bootstrap GA [kubernetes/features#43](https://github.com/kubernetes/features/issues/43)
        *   client bootstrap flow to GA in 1.12
        *   client rotation stays in beta in 1.12
        *   server rotation promoted to beta in 1.12 (kubelet-only, must bring own approver)
        *   CSR API is still v1beta1. Some potential improvements desired before v1:
            *   approver has input to add/remove/modify the cert that is signed
            *   signing profiles
            *   integration to produce certs with service DNS names as SANs


## July 25, 2018, 11a - Noon (Pacific Time)



*   [Recording](https://www.youtube.com/watch?v=PFGramA3F04)
*   Announcements
    *   (kksriram) Reviewing the Vault KMS provider with HashiCorp this week. Goal is to make repo public this week.
    *   feature freeze is 7/31... be sure there are issues filed in [kubernetes/features](https://github.com/kubernetes/features/issues?q=is%3Aopen+is%3Aissue+label%3Asig%2Fauth) for feature work targeted at 1.12 by then
        *   Please check milestone labels are correct for v1.12
        *   [sig-auth features](https://github.com/kubernetes/features/issues?q=is%3Aopen+is%3Aissue+label%3Asig%2Fauth)
        *   Also if somebody might prune out anything that's there which is old/stale or done / not going to be done, that would be awesome.  A lot of open 2016 and 2017 stuff for sig/auth.
*   Discussion
    *   Discuss addition of dynamic policy to audit kep https://github.com/kubernetes/community/pull/2407
        *
    *   Question for KEP process (take to sig-arch?)
        *   How do evolutions/additions to an accepted/implementable KEP fit with the provisional/implementable bit?


## July 11, 2018, 11a - Noon (Pacific Time)



*   [Recording](https://www.youtube.com/watch?v=r2RoQpxXyr0)
*   Discussion
    *   [plan for secrets](https://docs.google.com/document/d/1OVPuXDB759G_TIzTZ2jrKk-3EVJ4FUP9KNj5nMDOWZg/edit?ts=5b444461#) (@mayakacz, @immutableT)
    *   Improvements made:
        *   1.7 - kubelet doesn't have access to arbitrary secrets when using the Node authorizer and NodeRestriction admission plugin (on by default in kubeadm and kube-up deployments)
        *   1.7 - encryption at rest added in alpha state
        *   1.10 - KMS extension point added in alpha state
        *   1.11 - authorization of watches of individual secrets now possible (no longer require global watch authorization for controllers that only need particular secrets)
    *   Next steps?
        *   Promote encryption at rest from alpha
            *   https://github.com/kubernetes/kubernetes/pull/61592
            *   https://github.com/kubernetes/kubernetes/pull/63032
        *   Work with deployments to enable encryption at rest by default
        *   Promote KMS extension point
            *   need feedback from current known consumers ([Google Cloud KMS](https://github.com/GoogleCloudPlatform/k8s-cloudkms-plugin/),[ Microsoft Azure Key Vault](https://github.com/Azure/kubernetes-kms) and[ AWS KMS](https://github.com/kubernetes-sigs/aws-encryption-provider) plugins, one in the works for[ HashiCorp Vault](https://github.com/kubernetes/kubernetes/issues/49817))
            *   need an owner to drive this
        *   Continue work on audience/time-bound service account tokens in 1.12
        *   Stay involved with discussions of CSI volume plugins for secret injection


## June 27, 2018, 11a - Noon (Pacific Time)



*   [Recording](https://www.youtube.com/watch?v=L_fnXbPhsmM)
*   Announcements
    *   [New sig-auth chair](https://groups.google.com/forum/#!topic/kubernetes-sig-auth/R9scqXygqY8)
    *   New zoom meeting link (will update calendar and the link in this document)
*   Demos
*   Pulls of note
    *   [Add OIDC discovery to svcacct token issuer](https://github.com/kubernetes/community/pull/2314) (@mikedanese)
*   Discussion
    *   Authorization length limits (@tallclair) -  Do we want to support abusing the authorization interfaces for things completely unrelated to Kubernetes? For instance, many resources have a 253 character limit on names, but SubjectAccessReviews allow arbitrarily large values for name, resource, subresource, verb.
        *   What would the authorizer's response be if an API URL with a name segment longer than 253 chars was accessed?
        *   @tallclair will open issue proposing abuse limit (~8k, etc)
            *   Should detail behavior at each level if things that exceed that limit are presented
        *   Authorization layer is intentionally agnostic toward meaning of resource/subresource
    *   1.12 planning/discussion deferred to next meeting because of 1.11 wrap-up


## June 13, 2018, 11a - Noon (Pacific Time)



*   [Recording](https://www.youtube.com/watch?v=C5zbTKxGHR4)
*   Announcements
    *   [1.11 release notes for review](https://github.com/kubernetes/sig-release/blob/master/releases/release-1.11/release_notes_draft.md#sig-auth) - suggestions/pulls due by Friday, 6/15
*   Demos
    *
*   Pulls of note
    *
*   Discussion
    *   Last call for [chair nominations](https://groups.google.com/forum/?utm_medium=email&utm_source=footer#!msg/kubernetes-sig-auth/_uzEMhWQDPM/7r8DsiQ2AQAJ) to fill Eric's spot
        *   If you are interested in this role (or would like to recommend someone), please send Tim and Jordan [an email](mailto:kubernetes-sig-auth-leads@googlegroups.com) at with references to existing contributions.
        *   Hope to recommend a candidate and solicit feedback from the SIG by June 15th
    *   Bound Service Account Tokens use case for dynamic webhook configuration (@mattmoyer/@pbarker)
        *   Possible to fetch/use audience-scoped tokens for use with webhooks?
        *   How we determine the audience is questionable
        *   Token could be fetched/requested, or injected
    *   Further discussion on dynamic audit configuration https://github.com/kubernetes/community/pull/2188 (@pbarker)
        *   Static, non-overrideable flag-based config (what we have today)
        *   Dynamic configs (namespaced and cluster-scoped? Harmonize with other policy approaches - [scheduling](https://www.google.com/url?q=https://github.com/kubernetes/community/pull/1937&sa=D&ust=1528917891000000&usg=AFQjCNFqLzgSYw0OSFTS4sqgKk4oRHePeQ), etc. More discussion in [wg-policy](https://docs.google.com/document/d/1ihFfEfgViKlUMbY2NKxaJzBkgHh-Phk5hqKTzK-NEEs/edit))
    *   Node agents want to add taints to themselves (@mikedanese)
        *   [jordan's preferred solutions](https://github.com/kubernetes/kubernetes/pull/63167#issuecomment-385473656)
        *   Opinion: "Conditions are unnecessary indirections; well known key prefix is hard since there will be vendor specific taints; updating node controller is hard but is probably equivalent to updating the admission controller"
        *   We need this for [k8s-node-termination-handler](https://github.com/GoogleCloudPlatform/k8s-node-termination-handler)
    *   How do we track subproject efforts? (@mikedanese)
        *   Project experiment: https://github.com/kubernetes/kubernetes/projects/13
        *   Umbrella issues (e.g. [#60392](https://github.com/kubernetes/kubernetes/issues/60392))
    *   Quick note: [deficiencies of CSI as it relates to pod identity](https://github.com/kubernetes/kubernetes/issues/64984)
    *   Next meeting: 1.12 plans
        *   Mike: make umbrella issue for tokenrequest work
        *   Jordan: make umbrella issue for node restriction work
        *   @x13n: make umbrella issue for audit work
        *   <add your planned work items/areas here for discussion>


## May 30, 2018, 11a - Noon (Pacific Time)



*   [Recording](https://www.youtube.com/watch?v=FFaIUKBVgYw)
*   Demos
    *   [kube-oidc](https://github.com/ericchiang/kube-oidc): OpenID Connect utilities for Kube [ericchiang]
*   Pulls of note
    *   client-go: promote exec plugin support to beta ([#64482](https://github.com/kubernetes/kubernetes/pull/64482)) [ericchiang]
    *   dynamic audit configuration KEP ([#2188](https://github.com/kubernetes/community/pull/2188))
    *   NonRootGroup API Changes and Implementation ([#62216](https://github.com/kubernetes/kubernetes/pull/62216))
    *   For 1.11 [mikedanese]:
        *   Token Volume Source: https://github.com/kubernetes/kubernetes/pull/63819/
            *   Impl https://github.com/kubernetes/kubernetes/pull/62005
        *   TokenReview: https://github.com/kubernetes/kubernetes/pull/62802
*   Discussion
    *   Review bottleneck - how can we add expand OWNERS and add more approvers in some critical areas
        *   https://github.com/kubernetes/kubernetes/blob/master/plugin/pkg/auth/OWNERS
        *   https://github.com/kubernetes/kubernetes/blob/master/pkg/auth/OWNERS
        *   https://github.com/kubernetes/kubernetes/blob/master/staging/src/k8s.io/client-go/OWNERS
        *   https://github.com/kubernetes/kubernetes/blob/master/staging/src/k8s.io/client-go/util/certificate/OWNERS
        *   https://github.com/kubernetes/kubernetes/blob/master/pkg/kubelet/certificate/OWNERS
    *   1.11 Release timeline:
        *   https://github.com/kubernetes/sig-release/blob/master/releases/release-1.11/release-1.11.md


## May 16, 2018, 11a - Noon (Pacific Time)



*   [Recording](https://www.youtube.com/watch?v=qC3tM1XtV8U)
*   Announcements
    *   [ericchiang] Call for secondary for SIG Auth representative in SIG PM
        *   https://github.com/kubernetes/community/issues/2122
        *   https://github.com/kubernetes/community/blob/master/sig-product-management/SIG%20PM%20representatives.md
*   Pulls of note
    *   "Escalating" RBAC grants ([pull/56358](https://github.com/kubernetes/kubernetes/pull/56358))
*   Discussion
    *   [Cluster managed ca-certs](https://github.com/kubernetes/kubernetes/issues/63726) [tallclair]
    *   [destijl] It's fairly easy to be in an insecure state by continuously upgrading your cluster. Whereas if you just created a new cluster the defaults are safer. There's some [back-compat cluster role bindings](https://github.com/kubernetes/kubernetes/blob/master/cluster/addons/rbac/legacy-kubelet-user-disable/kubelet-binding.yaml#L1-L2) you need to delete, as well as ABAC disabling and other traps. How can we make this better? Should we just document the ones we know about? Only the most diligent will find and read those docs.
    *   [Exec auth plugin](https://github.com/kubernetes/community/blob/master/contributors/design-proposals/auth/kubectl-exec-plugins.md) - Beta requirements (awly@)
        *   https://github.com/heptio/authenticator/blob/master/pkg/token/token.go
        *   https://github.com/kubernetes/cloud-provider-gcp/tree/master/cmd/gke-exec-auth-plugin
        *   https://github.com/kubernetes/kubernetes/issues/61796
        *   Eric to document concerns and share, make a collective decision about whether it's ready for Beta.
    *   Cluster managed pull secrets for infrastructure images [clayton]
        *   Image pull secrets for infrastructure images (pod, daemonsets) or multiple namespaces is annoying to deal with
            *   (a fair number of people use protected registries)
        *   Might make sense to have a kubelet credentialprovider for accessing a kube secret(s) for being a pull secret  (kubelet --pull-secret=namespace/name)
            *   We already have other credential providers that deal with "infrastructure" and kubernetes is "infrastructure"
        *   David brought up that it should probably deal with node groups, which is a long standing problem with dynamic config for nodes (nodes that can label themselves can be anyone).  If we had groups of nodes (machineset or other) then we could also have controllers that manage access / rbac group membership for nodes


## May 2, 2018, 11a - Noon (Pacific Time)



*   Meeting cancelled


## April 18, 2018, 11a - Noon (Pacific Time)



*   [Recording](https://www.youtube.com/watch?v=T4vUfCn_h3s)
*   Announcements
    *   sig-auth status update @ contributor summit - [draft](https://docs.google.com/document/d/1w0DR0dMYRXQVz8unb93bN1Fs_PGLLpXtpm1FIeM0GMc/edit) [tallclair]
*   Pulls of note
    *   No new significant pulls, previously mentioned work is still ongoing
*   Designs of note
    *   [Security Profiles](https://docs.google.com/presentation/d/1PmRcgID9-TKX-0x3E3rkKNaijrXIe45KGTIKTLe5FC8/edit#slide=id.g36a26b6cbf_0_0) [yisui]
        *   Unified description of startup options (consumable by installers or manually)
        *   cluster-level policy (applied at setup or at profile change)
        *   namespace policy (continuously applied)
    *   Update on [ServiceAccountTokenVolumeProjection](https://github.com/kubernetes/community/pull/1973) [mikedanese]
        *   Presented to sig-storage
        *   Working through "public" configmap volume source ACL
*   Discussion
    *   [Governance proposal](https://github.com/kubernetes/community/pull/2000)
        *   Lays out three leadership roles (chair, tech lead, subproject owner)
        *   Tech lead and subproject owner roles were made distinct to think through which responsibilities belonged where, even though initially, the same people might be filling both roles
        *   Defines sig-auth membership (contributors to a subproject or sig health area, docs, tests, etc), for purposes of voting if needed
        *   Minimal comments on doc, will send out call for lazy consensus to mailing list
    *   [Subprojects](https://docs.google.com/document/d/1RJvnSPOJ3JC61gerCpCpaCtzQjRcsZ2tXkcyokr6sLY/edit#heading=h.t1pscltwx2kr) (must be a member of the [sig-auth google group](https://groups.google.com/forum/#!forum/kubernetes-sig-auth) to view)
        *   Will convert to community PR and send out to mailing list for comments (and to ensure identified subproject owners agree they are responsible for those areas)
    *   [Guidelines for node-level plugin auth](https://github.com/kubernetes/kubernetes/issues/62747) [tallclair]


## April 4, 2018, 11a - Noon (Pacific Time)



*   [Recording](https://www.youtube.com/watch?v=PI_Ue6WAJHo)
*   Announcements
    *   Calendar invite updated
*   Demo
    *   [citadel](https://github.com/enj/citadel) @enj @npmccallum: Turn an arbitrary command into a KMS GRPC server (~5 to 10 minutes)
        *   Feedback would be greatly appreciated
        *   Many systems you would want to integrate with don't use gRPC
        *   Citadel implements the gRPC service, then fetches the key encryption key from the external KMS
        *   TODO: measure performance
*   Pulls of note
    *   [scheduling policy](https://github.com/kubernetes/community/pull/1937) (@tallclair)
        *   Questions about what policy should look like inside of Kubernetes
            *   How policies are composed, defaulting, binding, etc…
        *   Next policy we need to add as in tree policy
        *   Conversations in [policy-wg](https://github.com/kubernetes/community/tree/master/wg-policy)
        *   On a related note, [PodRestriction](https://github.com/kubernetes/community/pull/1950) won't go forward
    *    Removing node.ExternalID (@mikedanese)
        *   https://github.com/kubernetes/kubernetes/pull/60692
        *   https://github.com/kubernetes/kubernetes/pull/61877
        *   Removes the need for kubelets to self-delete their Node objects \o/
        *   Field is replaced with instanceID field
            *   Mutable, so kubelets never need delete/recreate to change
            *   Use as source-of-truth vs cloud-provider lookup optimization is being debated in other PRs
*   Designs of note
    *   [Kubelet TLS bootstrap via TPM](https://docs.google.com/document/d/12UgErB_46iHBOi0YEKbaFbEXv9E5O6XWQihwPkwB_7s) (@awly)
        *   Data placed in CSR objects is not guaranteed to be confidential, need to be careful to ensure it is tightly bound to the specific CSR
            *   CSR attestation data doesn't have sensitive information. Even if CSR or final approved cert are leaked, they are useless without corresponding private key.
        *   Related to TLS client cert support for exec plugin: https://github.com/kubernetes/kubernetes/pull/61803
            *   How would we make this support non-exportable keys?
            *   Would adding non-exportable key support overload the exec plugin interface? What performance expectations do we have from TLS handshakes?
    *   [service account token volumes](https://github.com/kubernetes/community/pull/1973) (@mikedanese)
        *   Open question: how to get the CA cert to the volume?
            *   Possibly use kube-public namespace?
*   Discussion
    *   [CRI streaming authentication](https://github.com/kubernetes/kubernetes/issues/36666#issuecomment-378440458)
    *   [Governance proposal](https://github.com/kubernetes/community/pull/2000)
        *   Based on the [short-form SIG governance template](https://github.com/kubernetes/community/blob/master/committee-steering/governance/sig-governance-template-short.md)
        *   Proposes leadership roles: chairs, technical leads, subproject owners
        *   Describes requirements for roles, ideals for how the project would run
        *   See also: [subproject list and status](https://docs.google.com/document/d/1RJvnSPOJ3JC61gerCpCpaCtzQjRcsZ2tXkcyokr6sLY/edit#heading=h.t1pscltwx2kr) (in progress)
        *   Will take feedback/comments in the PR or mailing list and discuss in the next sig-auth meeting


## March 21, 2018, 11a - Noon (Pacific Time)



*   [Recording](https://www.youtube.com/watch?v=EooFi_nBA7c)
*   Demo
    *   https://github.com/CaoShuFeng/k8s-audit-collector
        *   Some comments about use of service account
            *   Exporting SA token and using it for user-auth seems like an anti-pattern we should discourage. Where you can't distinguish code taking actions with a credential from users taking actions it complicates access control, auditing, and incident response when evaluating damage of a lost/stolen credential.
            *   Users should only assume service account identities during debugging. Can be achieved without service account credentials using ActAs delegation.
        *   Lack of guidance here?
            *   Action item: write this up!
*   Designs of note
    *   [PodRestriction](https://github.com/kubernetes/community/pull/1950) (@tallclair)
        *   Whitelist and blacklist type of restrictions only.
        *   Scheduling piece of this yet to be decided.
        *   Enforced at namespace level.
        *   Feedback requested on [`BindingMode`](https://github.com/kubernetes/community/pull/1950/files#diff-4e2582430273ddd4c245083aff3b23e1R119) behaviour.
        *   Overlap with [policy wg](https://docs.google.com/document/d/1Ht8wpj4j9YfAA7aVv9Yn3Ci1T_MLMWt0DBr0QmxI2OM/edit#heading=h.6l8mozsiyzb1) and [opa](http://www.openpolicyagent.org/docs/kubernetes-admission-control.html)
        *   How does lack of defaulting working with pods that express no opinion?
            *   Handled through something like pod presets
        *   TODO: sig-scheduling and policy-wg
    *   [Update] [kubernetes/community#911](https://github.com/kubernetes/community/pull/911/) reduce scope of node on taints/labels (sig-node)
        *   Deferring work on moving reporting of node address
*   Discussion
    *   [SIG contribx presentation](https://docs.google.com/document/d/1hESUJdDy6Q6eysNe_Wyjtq8vI5A3AnrfWCcJFP4yU6I/edit) (Paris Pittman)
        *   Feedback wanted
        *   Defining bot commands and labels
            *   How are you using them? Are you using them?
            *   Help wanted issue and beginner issues
            *   How do you find out about these changes?
        *   Contributor and developer guide
            *   Please review
        *   Suggestions for mentoring
    *   Formalizing feature tracking for 1.11 (@ericchiang)
        *   Many features were in an unclear state over 1.10
            *   Not just alpha/beta/stable, but what work there was to do, what work had been done
        *   Considering encouraging more umbrella bugs like [#60392](https://github.com/kubernetes/kubernetes/issues/60392)
        *   Subprojects


## ~~March 7, 2018, 11a - Noon (Pacific Time)~~



*   NO SIG AUTH (busy with release)
*   Announcements
    *   SIG-auth presenting at the Kubernetes community meeting


## February 21, 2018, 11a - Noon (Pacific Time)



*   [Recording](https://www.youtube.com/watch?v=W82dsv9ITNs)
*   Announcements
    *   1.10 code freeze is 2/26
        *   Time to make progress on stalled PRs/designs if desired for 1.10
*   Pulls of note
    *   [TokenRequest Authenticator (pull/59940)](https://github.com/kubernetes/kubernetes/pull/59940)
*   Designs of note
*   Needs Review
    *   Documenting the KMS encryption provider ([Issue](https://github.com/kubernetes/website/issues/7399), [PR](https://github.com/kubernetes/website/pull/7479)).
    *   10 sig/auth PRs in the v1.10 milestone ([PRs](https://github.com/kubernetes/kubernetes/pulls?q=is%3Aopen+is%3Apr+milestone%3Av1.10+label%3Asig%2Fauth))
*   Discussion
    *   [Container Isolation](https://docs.google.com/document/d/1QQ5u1RBDLXWvC8K3pscTtTRThsOeBSts_imYEoRyw8A/edit) (@tallclair)
        *   Great that this is separated from individual implementations
        *   Goal is to have some sort of way to expressing a workload running in a VM
        *   Questions about boundary being the container or the pod. Will be expressed as different APIs. (similar to [kubernetes#55435](https://github.com/kubernetes/kubernetes/issues/55435))
    *   [Kubernetes bug-bounty program](https://docs.google.com/document/d/1dvlQsOGODhY3blKpjTg6UXzRdPzv5y8V55RD_Pbo7ag/edit?ts=5a8ca2dd#heading=h.7t1efwpev42p) (Maya Kaczorowski)
        *   Warnings about spam
        *   Questions about cross-team coordination
            *   Need to give the other SIGs the heads up
            *   Don't do it before freeze
        *   Current triage
            *   Whoever responds first becomes the owner
            *   Explicit rotation
            *   Fix and release process
            *   Requires adjustments
        *   How do we handle scope?
            *   Conformance
            *   Questions around plugins
        *   Out of scope
            *   Outside of Kubernetes org
            *   What's in the Kubernetes incubator?
        *   Want to be programmatic about the scope
            *   We have descriptions about hardening, but not actual tests
        *   Start very narrow to start
            *   E.g. only in core, no plugins
        *   Note: someone to take over the threat model doc https://github.com/kubernetes/community/pull/504
        *   Action items:
            *   Funding
            *   Working with other SIGs
            *   Owners


## February 7, 2018, 11a - Noon (Pacific Time)



*   [Recording](https://www.youtube.com/watch?v=kls1pbC71n8)
*   Announcements
    *   1.10 code freeze is 2/26
        *   Time to make progress on stalled PRs/designs if desired for 1.10
*   Pulls of note
*   Designs of note
    *   serviceaccount secrets
        *   Consumption of just-in-time service account tokens
        *   Plumbing other info for the in-cluster config (CA, etc) to the kubelet
        *   https://github.com/kubernetes/kubernetes/issues/48408
        *   https://docs.google.com/document/d/1Ki7K9rAC5XIAgyYJXUis1UpgwmtR \
q2Xz68mF3_LggEY/edit?ts=5a68cdbc
*   Needs Review
    *   gRPC KMS extension point implementation ([#55684](https://github.com/kubernetes/kubernetes/pull/55684))
        *   Last call for comments this week
        *   External Google implementation in progress
        *   External Vault implementation in progress
    *   External client-go credential providers ([#59495](https://github.com/kubernetes/kubernetes/pull/59495))
        *   Last call for comments this week
        *   Coordinate with api-machinery / Chao as client-go owner
*   Discussion
*   sig lead change
        *   David Eads (@deads2k, Red Hat) stepping down as sig lead
        *   Tim Allclair (@tallclair, Google) nominated as replacement
            *   Long-term contributor to k8s auth/security
            *   Helped drive pod security policy and audit features
            *   Member of kubernetes product security committee
            *   Brings container/node security expertise
            *   Unanimous support from other leads (Jordan Liggitt, Red Hat; Eric Chiang, CoreOS)
        *   Feedback on the change welcome (either public or private)
        *   Lead change to regain some company diversity was desired before starting to work through defining sig membership, subprojects, tech leads, etc, etc
*   Discuss the [sig governance template PR](https://github.com/kubernetes/community/pull/1650) (not merged, but fairly complete) as it applies to sig-auth:
        *   Do we want to split out roles for sig-auth as proposed in the PR: project lead (needs a better name) that is organizational and helps set strategic direction and tech lead (heavy code, design architecture focus).
            *   How many of each role?
            *   Organizational diversity: after the acquisition of coreos all our current sig-auth leads work at the same company.
        *   Agreement that there is value in documenting functional areas sig-auth is involved in, and the people responsible for them
        *   Agreement that having people focused on strategic direction is valuable (complementing per-release work item planning)


## January 24, 2018, 11a - Noon (Pacific Time)



*   [Recording](https://www.youtube.com/watch?v=hsSaYXBEKaI)
*   Pulls of note
    *   RBAC and PSP in audit - [#58143](https://github.com/kubernetes/kubernetes/pull/58143)
    *   OIDC cleanup - [#58544](https://github.com/kubernetes/kubernetes/pull/58544)
*   Designs of note
    *   Proposal: Validated Pod Annotations - [community#1604](https://github.com/kubernetes/community/pull/1604)
*   Needs review:
    *   gRPC KMS versioning ([#55684](https://github.com/kubernetes/kubernetes/pull/55684))
    *   Introduce podsecuritypolicy.admission.k8s.io/v1beta1 [#54933](https://github.com/kubernetes/kubernetes/pull/54933)
*   Discussion:
    *   Status on the client-go auth provider proposal. - [community#1503](https://github.com/kubernetes/community/pull/1503)
        *   Open questions:
            *   Format of client-go <-> plugin communications
            *   Call patterns (cache in memory, on 401 or expiry re-call)
    *   Should Audit Logging API go to v1 in 1.10? [@destijl, @crassirostris, @tallclair]
        *   Google: we think it should, we haven't seen problems in our usage that would require changes.
        *   Require passing scale tests for v1
            *   Issue: https://github.com/kubernetes/kubernetes/issues/53020
            *   Depends on: https://github.com/kubernetes/kubernetes/pull/56126
        *   Require specing out how the rest of the system contributes audit info
        *   Requirements & discussion on https://github.com/kubernetes/kubernetes/issues/58083
    *   [enj] Thoughts on a subject rules review API (in relation to the current self subject rules review API)
        *   Use case being "can x do all of [a, b, c]" or "is x an admin in namespace foo"
        *   Granted it would be nicer to just ask the server to do the coverage check as well
        *   Intent of SSRR API was advisory, not for externalizing policy rules for external authoritative evaluation
        *   Still building out authorizer support (webhook, node)
        *   Still need to reconcile with explicit deny support
        *   Existing PR to express explicit deny in SelfSubjectRulesReview [#53922](https://github.com/kubernetes/kubernetes/pull/53922)
    *   Status of TokenRequest proposal [kubernetes/community#1460](https://github.com/kubernetes/community/pull/1460)
        *   Expressing request for expiring vs non-expiring token
        *   Version of subresource (v1 vs v1alpha1)
    *   Reducing scope of node on node actions [community#911](https://github.com/kubernetes/community/pull/911)
        *   Current proposal to use initializers (fate of initializers unclear)
        *   Could have the approver (or signer) be cloud aware
            *   If signer can fill in SANs, then client doesn't need as much access to figure out its requested SANs.


## January 10, 2018, 11a - Noon (Pacific Time)



*   [Recording](https://www.youtube.com/watch?v=WIVJmUoqCuA)
*   Pulls of note
    *   TokenRequest API [kubernetes/kubernetes#58027](https://github.com/kubernetes/kubernetes/pull/58027)
    *   KMS provider https://github.com/kubernetes/kubernetes/pull/55684 (design from google doc https://github.com/kubernetes/community/pull/1581)
*   Designs of note
    *   Proposal: external client-go auth providers - [community#1503](https://github.com/kubernetes/community/pull/1503)
*   Needs review:
    *
*   Discussion:
    *   Kubelet TLS bootstrapping - GA in 1.10? - [features#43](https://github.com/kubernetes/features/issues/43#issuecomment-340082591) (@mattmoyer)
        *   Note(@ericchiang): some open questions here: [kubernetes#57164](https://github.com/kubernetes/kubernetes/issues/57164)
        *   Client bootstrapping
            *   Could interact with client-go auth provider, obtain bootstrap credential externally
        *   Client rotation
            *   Still have a lot of logic baked into the kubelet, makes it difficult to take advantage of platform-specific CSR-informing features
        *   Server bootstrapping
            *   Blocked on self-reported node address trustworthiness (API-enabling work in 1.10 - @liggitt)
            *   Requires refactor of kubelet startup flow and triggers of serving cert updates
                *   Question: externalize and make the kubelet pick up certs from disk?
    *   Idea for K8s vulnerability reward program [@destijl]
        *   Needs well-defined setup/config demonstrating current best practices
        *   Define boundaries well
        *   Initial spike of activity/responses needed
        *   Funnel into existing security process, use as impetus to make process and timeframes better defined
        *   Ongoing commitment for new releases, responding with fixes (cross-sig)
        *   Long-term, could inform other CNCF projects
    *   Cross-authorizer RBAC escalation check - [#56358](https://github.com/kubernetes/kubernetes/pull/56358) (@liggitt)
        *   Will rethink as specific escalate query rather than superuser subject access review (parallel to bind check)
    *   TokenRequest and key rotation [@mikedanese]
        *   Current tokens are non-expiring, but bound to service account and secret objects (by embedded uid claims)
        *   How to encourage applications to expect/support rotation
            *   Client-go/in-cluster-config libraries support rotation
        *   Option: separate signing keys for expiring/non-expiring
        *   Limits on token lifetime - min of:
            *   Requester expiration time (if specified)
            *   Admin-specified max time (if specified)
            *   Deletion of bound object (pod, secret, service account, etc)
