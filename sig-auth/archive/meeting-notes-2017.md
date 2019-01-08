# Kubernetes Sig-auth Meeting Agenda

## December 27, 2017 - No meeting, Christmas holiday


## December 13, 2017, 11a - Noon (Pacific Time)



*   [Recording](https://youtu.be/hvbPwNpsflE)
*   Announcements
    *   Code freeze lifts today
    *   kubernetes-wg-multitenancy - [google group](https://groups.google.com/forum/#!forum/kubernetes-wg-multitenancy)
*   Demos
*   Pulls of note
    *   User-facing RBAC policies for NetworkPolicy - [#56650](https://github.com/kubernetes/kubernetes/pull/56650) (already approved)
*   Designs of note
    *   Proposal: bound service account tokens - [community#1460](https://github.com/kubernetes/community/pull/1460)
    *   Proposal: self-hosted webhook authorizers - [community#1458](https://github.com/kubernetes/community/pull/1458)
*   Needs review:
    *   Encryption at rest: KMS gRPC plugin - [#55684](https://github.com/kubernetes/kubernetes/pull/55684)
        *   AI: sig-architecture office hours
    *   Rules review: express denies and add webhook support - [#53922](https://github.com/kubernetes/kubernetes/pull/53922)
    *   Reducing scope of node on node actions - [community#911](https://github.com/kubernetes/community/pull/911)


## November 29, 2017, 11a - Noon (Pacific Time)


## Canceled due to release crunch and lack of agenda items.


## November 15, 2017, 11a - Noon (Pacific Time)



*   [Recording](https://www.youtube.com/watch?v=7EPCh8kymrM&list=PL69nYSiGNLP0VMOZ-V7-5AchXTHAQFzJw&index=1)
*   Announcements
    *   Code freeze is coming Nov 22nd.
*   Demos
*   Pulls of note
    *   Kubectl OpenID Connect login (1.10) - [#55514](https://github.com/kubernetes/kubernetes/pull/55514)
        *   Separate login and rotation
        *   Figure out out of tree login/rotation
    *   KMS gRPC plugin (1.10) - [#55684](https://github.com/kubernetes/kubernetes/pull/55684)
        *   Concerns about another plugin mechanism
        *   Make sure that API machinery/API reviewers see this
    *   Allow nodes to impersonate service accounts of pods - [#55019](https://github.com/kubernetes/kubernetes/pull/55019)
        *   Part of a prototype for pod identity
*   Designs of note
*   Discussions
    *   RBAC: way to provide extra context like the node authorizer?
        *   E.g. field selector and label selector
        *   Node authorizer doesn't get additional data about the request, still requires admission controller.
        *   Still possible to create something externally that builds a graph and creates RBAC roles. E.g. ingress secrets for an ingress controller
        *   Help solve the ingress problem: https://github.com/kubernetes/ingress-nginx/issues/816
        *   Action item: Review secrets roadmap for 1.10


## November 1, 2017, 11a - Noon (Pacific Time)



*   [Recording](https://youtu.be/uDHNcZ0eGHI)
*   Announcements
*   Demos
    *    [SPIFFE SPIRE](https://groups.google.com/forum/?utm_medium=email&utm_source=footer#!msg/kubernetes-sig-auth/s3EX5Wmew_Q/1QaOLDb0AwAJ) - Evan Gilman (evan2645)
        *   https://github.com/spiffe/spiffe
        *   https://github.com/spiffe/spire
*   Pulls of note
    *   Google Cloud KMS removal! - [54759](https://github.com/kubernetes/kubernetes/pull/54759) - from [54439](https://github.com/kubernetes/kubernetes/pull/54439) .  This kills [51760](https://github.com/kubernetes/kubernetes/pull/51760)
    *   Short-circuit deny - [#53273](https://github.com/kubernetes/kubernetes/pull/53273) (looks ready to merge, but affects selfsubjectrulesreview)
    *   SelfSubjectRulesReview support for webhook authorizer - [#53922](https://github.com/kubernetes/kubernetes/pull/53922) - what do we do with short-circuit denies?
    *   Aggregated RBAC roles - [#54005](https://github.com/kubernetes/kubernetes/pull/54005) - needs review cycles
*   Designs of note
    *   [External KMS](https://github.com/kubernetes/kubernetes/issues/51965#issuecomment-340850795) - Homework for next time
        *   [Strawman](https://docs.google.com/document/d/1S_Wgn-psI0Z7SYGvp-83ePte5oUNMr4244uanGLYUmw/edit?ts=59f965e1#) for discussion.
*   Discussions
    *   Admission plugins - PSP mutate/validate split in particular
    *   kube-apiserver insecure port deprecation and eventual removal
        *   API aggregation requires secured port
        *   Doesn't support admission plugins wanting authorizers
        *   Going to be difficult in the cluster scripts, but we'll hit and resolve bumps there. - https://kubernetes.io/docs/tasks/access-kubernetes-api/http-proxy-access-api/
    *   For 1.10, external SA token creation and mounting
        *   Clayton/MikeDanese? to send sketch.


## October 18, 2017, 11a - Noon (Pacific Time)



*   [Recording](https://youtu.be/P6oY2FdqYvY)
*   Announcements
    *   New experimental [AWS IAM webhook authenticator](https://github.com/heptiolabs/kubernetes-aws-authenticator) @mattmoyer
*   Demos
*   Pulls of note
    *   Authorizer support for unequivocal deny - [#53273](https://github.com/kubernetes/kubernetes/pull/53273)
    *   Aggregated RBAC roles - [#54005](https://github.com/kubernetes/kubernetes/pull/54005)
    *   SelfSubjectRulesReview support for webhook authorizer - [#53922](https://github.com/kubernetes/kubernetes/pull/53922)
*   Designs of note
    *   [Trustworthy Workload JWTs (v2 of previous doc)](https://docs.google.com/a/google.com/document/d/1jyQjfIOJxOMpyuSV5EBpJY97TjOet6nRJ95zroeKO7k/edit?usp=sharing)
*   Discussions


## October 4, 2017, 11a - Noon (Pacific Time)



*   [Recording](https://youtu.be/qA-UrUsqujw)
*   Announcements
    *   1.8 release retrospective this week, 10/6 at 10a PST - [agenda](https://docs.google.com/document/d/1pZ_hOxrwwPHA2lEWaJizfkvf80eMluVCBw9PDwj8oMA/edit#)
*   Demos
    *   [audit2rbac](https://github.com/liggitt/audit2rbac) - @liggitt
        *   Potential [solution for using this on GKE](https://github.com/kubernetes/kubernetes/issues/53455).
*   Pulls of note
    *   CSR Cleanup Controller - #[51840](https://github.com/kubernetes/kubernetes/pull/51840)
    *   PodSecurityPolicy ordering, non-mutation - #[52849](https://github.com/kubernetes/kubernetes/pull/52849)
*   Designs of note:
    *   Reduce scope of node on node object ([kubernetes/community#911](https://github.com/kubernetes/community/pull/911))
        *   Needed to prevent a node from getting a serving cert for another node
        *   Need to make sure that version skewed kubelets don't self-destruct (n-2)
        *   Could use a cloud provider initializer/controller to provide the information about serving cert.
    *   Vault KMS provider ([kubernetes/community#888](https://github.com/kubernetes/community/pull/888))
    *   [Verifying Workload JWTs](https://docs.google.com/document/d/1pyauN6hffhVbc3xdgi6fGsp0p189xgGaX9k4poOHK9s/edit)
        *   Minting short lived JWTs with audiences
        *   Can't be used to talk to Kubernetes API, e.g. you could verify your identify without giving your access
        *   Two possible models:
            *   Decentralized (built up from container identity)
            *   Centralized (request to API server for a short lived JWT)
        *   Request: feedback on which model we'd prefer
*   1.9 planning
    *   PodSecurityPolicy - @liggitt, @tallclair
        *   ordering/priority - [#52849](https://github.com/kubernetes/kubernetes/pull/52849), [#45358](https://github.com/kubernetes/kubernetes/pull/45358) @liggitt
        *   GCE/GKE/CI enablement - [#52367](https://github.com/kubernetes/kubernetes/pull/52367) @tallclair
        *   bootstrapping/default policy - [#52367](https://github.com/kubernetes/kubernetes/pull/52367) @tallclair?
        *   Transitioning out of extensions API group to which group? Possibly "policy" group, alongside other pod-related policies (PodDisruptionBudget). Requires more thought.
        *   AllowedFlexVolumes - [community#723](https://github.com/kubernetes/community/pull/723) @php-coder?
    *   Node label/taint restriction - [features#279](https://github.com/kubernetes/features/issues/279), [community#911](https://github.com/kubernetes/community/pull/911) @mikedanese, @liggitt
        *   Required to prevent nodes from steering pods
        *   Backwards compatibility issues with existing labels/taints
        *   Coordinate with sig-node (@???), sig-scheduling (@davidopp)
    *   Node status address restriction - [community#911](https://github.com/kubernetes/community/pull/911) @mikedanese, @jcbsmpsn
        *   Required for nodes to obtain serving certs via CSR -
        *   Split kubelet startup flow to obtain addresses from node status
        *   Move authoritative source of addresses to node controller
        *   Coordinate with sig-node/sig-cloud-provider - @???
    *   ClusterRoles for API extensions (CRD and UAS) - @deads2k
        *   https://groups.google.com/forum/?utm_medium=email&utm_source=footer#!msg/kubernetes-sig-auth/lbnMMXq7TyA/qUHNNJSNBwAJ
    *   Design for out-of-process KMS provider API - ~~@jcbsmpsn~~
        *   https://github.com/kubernetes/kubernetes/issues/51965
    *   Vault as KMS provider - @kk, @deads2k
        *   https://github.com/kubernetes/community/pull/888
        *   https://github.com/kubernetes/kubernetes/pull/51760
        *   This is our second implementation for sorting out what an out of process extension point would look like.
        *   The hope is this work provides more information to facilitate a general integration mechanism - https://github.com/kubernetes/kubernetes/issues/51965
    *   AWS as KMS provider - @mattmoyer (tentative)
        *   https://github.com/kubernetes/kubernetes/issues/52916
        *   Blocked on out-of-process extension point (WIP/discussion:???)
        *   Bug for adding an interface to allow out of process integration of KMS providers - https://github.com/kubernetes/kubernetes/issues/51965
    *   Container Isolation (strategy & design) - @tallclair
        *   Holistic look at the future of (Pod)SecurityContext, PodSecurityPolicy, Authz (for pods/PSP), VM based pods, etc.
    *   Continued client cert rotation work - @jcbsmpsn, @mikedanese
        *   Better define failure modes, fallback behaviors
        *   Better testing coverage (enabled in e2e with short expire times)
        *   Progress on bootstrap/attestation work
            *   Would like to change auto-approve to add checks for attestation material in the CSR in environments where that is available (could allow "bootstrap" token to be long-lived)
    *   [ add items you plan to work on ]
    *
    *
*   Discussions
    *   (if we have time) [jbeda] Appropriate usage of TokenReview API
        *   Designed for delegated token authentication _in the same auth domain_ (API -> webhook, kubelet -> API, extension server -> core API server)
        *   Not intended to promote using kube API tokens as authentication to systems outside the API auth domain

Attendees:



*   Joe Beda, Matt Moyer -- Heptio
*   Steven E. Harris — IBM
*   Others?


## September 20, 2017, 11a - Noon (Pacific Time)



*   [Recording](https://youtu.be/jRvlrRKDSs4)
*   Pulls of note
    *   [#52367](https://github.com/kubernetes/kubernetes/pull/52367) - PSP config, looks dependent on [50169](https://github.com/kubernetes/kubernetes/pull/50169)
    *   [#52654](https://github.com/kubernetes/kubernetes/pull/52654) - Add pod disruption budgets to admin/edit/view roles
*   Designs of note:
        *   [Vault KMS Provider](https://github.com/kubernetes/community/pull/888) - Likely for 1.9.  Comments so far look minor, request for reviews
        *   Needs [the plugin mechanism to work](https://github.com/kubernetes/kubernetes/issues/51965) since vault isn't a cloud provider, see [#48574](https://github.com/kubernetes/kubernetes/pull/48574) for background discussion
        *   deads to follow up with api machinery
    *   [Verifying Workload JWTs](https://docs.google.com/document/d/1pyauN6hffhVbc3xdgi6fGsp0p189xgGaX9k4poOHK9s/edit)
        *   Minting short lived JWTs with audiences
        *   Can't be used to talk to Kubernetes API, e.g. you could verify your identify without giving your access
        *   Two possible models:
            *   Decentralized (built up from container identity)
            *   Centralized (request to API server for a short lived JWT)
        *   Request: feedback on which model we'd prefer
*   Discussions
    *   cert-manager [project](https://github.com/jetstack-experimental/cert-manager) for automatic acquisition and renewal of x509 certificates for TLS in Kubernetes from a variety of issuing sources / [incubator proposal](https://github.com/jetstack-experimental/cert-manager/pull/50) - @munnerz
        *   Note: incubator was recently [frozen until after the election](https://groups.google.com/forum/#!topic/kubernetes-dev/iaHlJnA1T4Y)
        *   Several existing projects were focused on ingress TLS through ACME
            *   kube-lego, kube-cert-manager, kubernetes-letsencrypt
        *   Similar API to the certificates API through CRDs
        *   Acts as a ACME style issuer
    *   [support unequivocal DENY in union authorizer](https://github.com/kubernetes/kubernetes/issues/51862)
        *   General agreement on supporting this
        *   Would like to see how it's actually hooked up
        *   Questions about just switching authorizer to AND instead of OR


## September 6th, 2017, 11a - Noon (Pacific Time)


## Canceled due to release crunch and lack of agenda items.


## August 23th, 2017, 11a - Noon (Pacific Time)



*   [Recording](https://youtu.be/t6O8MZ5tgAg)
*   Announcements
    *   Container identity working group meets [Friday 11am EST](https://groups.google.com/forum/#!topic/kubernetes-wg-container-identity/IeKC18iFf9w)
    *   Code freeze on September 1st
*   Demos
    *   SPIFFE + SPIRE overview (Evan Gilman, [Scytale](https://www.scytale.io/)) ~20 min
        *   SPIFFE (https://spiffe.io/, https://github.com/spiffe/spiffe)
            *   Standard for service identity and credentials (x509)
            *   spiffe://example.com/foo/bar <- SVID
            *   API endpoint to get private keys and public certs (ca bundle a caller should trust).  Doesn't specify auth-n/z of API.
            *   Istio uses SVID as identity, doesn't implement workload API
        *   SPIRE (will be open sourced in the future)
            *   Implementation of SPIFFE
            *   Node agent that has to attest to a central "control plane"
            *   Node agent can then request credentials for workloads
        *   SPIRE + kube
            *   Few possible thoughts for "this pod is running on this node"
                *   SPIRE shim for workload registration
                *   Admission controller
                *   Kubelet "plugin"
            *   Node attestation
                *   SPIRE-based, SPIRE attests for kubelet
                *   Kubernetes-based, kubelet attests for SPIRE agent
            *   Workload attestation
        *   SPIFFE + kube
        *   Source: https://docs.google.com/presentation/d/1C49epueLEjsv8QsQ36EoWbK9EuIHvbocGMu1vGmIGNo/edit?usp=sharing
*   Pulls of note
    *   Bootstrap token extra groups ([#50933](https://github.com/kubernetes/kubernetes/pull/50933))
    *   Client cert rotation to beta ([#51045](https://github.com/kubernetes/kubernetes/pull/51045))
    *   Configurable OIDC value prefixes ([#50875](https://github.com/kubernetes/kubernetes/pull/50875))
    *   Google cloud KMS provider ([#48574](https://github.com/kubernetes/kubernetes/pull/48574#issuecomment-324082506))
*   Designs of note:
    *   Reduce scope of node on node object ([kubernetes/community#911](https://github.com/kubernetes/community/pull/911))
        *   Needed to prevent a node from getting a serving cert for another node
        *   Need to make sure that version skewed kubelets don't self-destruct (n-2)
        *   Could use a cloud provider initializer/controller to provide the information about serving cert.
    *   Vault KMS provider ([kubernetes/community#888](https://github.com/kubernetes/community/pull/888))
        *   Greg read it and had a question about authentication backends.  How do you plug in different authentication backends? Kube changes needed?
*   Discussion
    *   SelfSubjectRulesReview going into a v1 API group? (@ericchiang) ([#48051](https://github.com/kubernetes/kubernetes/pull/48051))
        *   Known and learned from openshift
        *   Never persisted (eliminates the migration concern)
        *   Only difference is partial results because of multiple authorizers
        *   Should decide this week ahead of freeze


## August 9th, 2017, 11a - Noon (Pacific Time)



*   [Recording](https://youtu.be/DJDuDNALcMo)
*   Pulls of note
    *   Enable short caching of successful token authentication ([#50258](https://github.com/kubernetes/kubernetes/pull/50258))
    *   E2E node tests ([#49869](https://github.com/kubernetes/kubernetes/pull/49869)), blocked on moving sig-auth e2e tests ([#48966](https://github.com/kubernetes/kubernetes/pull/48966))
    *   CSR approver, adding resync ([#49788](https://github.com/kubernetes/kubernetes/pull/49788))
    *   Node authorizer on by default in local-up-cluster ([#49812](https://github.com/kubernetes/kubernetes/pull/49812)) (merged)
*   Designs of note:
    *   Rules review API proposal ([kubernetes/community#887](https://github.com/kubernetes/community/pull/887))
    *   Out-of-tree keystone authenticator/authorizer ([link](https://groups.google.com/forum/#!topic/kubernetes-sig-openstack/mJQ2jjXRZno))
*   Discussions
    *   [Generating RBAC profiles](https://docs.google.com/document/d/1jyNDETzIUDJWo2ZNscCM7C4scB7U73lSvdrgDy79wyw/edit#) (@tallclair)
    *   Request to cherry pick an OIDC server fix ([#48907](https://github.com/kubernetes/kubernetes/pull/48907))
    *   Container identity working group kickoff summary
        *   [Recording](https://www.youtube.com/watch?v=jcQ0mY7PQ8c&feature=youtu.be)
        *   Talked through the proposal and boundaries
        *   [Initial proposal](https://docs.google.com/document/d/1bCK-1_Zy2WfsrMBJkdaV72d2hidaxZBhS5YQHAgscPI/edit) for the wg
        *   Greg Castle's [design discussion](https://docs.google.com/document/d/1no8rYJ_nzhMeXYLL6JLjVSDtrj7pd6CBBfE3cPGOv8g/edit#heading=h.xgjl2srtytjt)
        *   [Mailing list ](https://groups.google.com/forum/#!forum/kubernetes-wg-container-identity)
    *   SIG apps charts + RBAC discussion
        *   [Recording](https://www.youtube.com/watch?v=eaQlsRoZVqs)
        *   [Notes](https://docs.google.com/document/d/1LZLBGW2wRDwAfdBNHJjFfk9CFoyZPcIYGWU7R1PQ3ng/edit#heading=h.9hy2gp3abcm5)


## July 26th, 2017, 11a - Noon (Pacific Time)



*   [Recording](https://youtu.be/0ENcstWC06k)
*   Pulls of note
    *   Allow nodes to evict their own pods ([#48707](https://github.com/kubernetes/kubernetes/pull/48707))
    *   Best practices for secrets ([#4380](https://github.com/kubernetes/kubernetes.github.io/pull/4380))
    *   Selfsubjectrulesreview ([#48051](https://github.com/kubernetes/kubernetes/pull/48051))
        *   Namespaced or cluster-scoped?
            *   At cluster scope make sure this can't expose information about what namespaces exist
        *   Include control over "extra" data or not?
            *   Used to plumb information between authentication and authorization (example: oauth scopes, keystone project role info)
            *   Use case: give a scoped credential to a third party that is only allowed to submit selfsubjectrulesreviews, and allow that third party to ask about unscoped permissions
            *   If "extra" is only used to limit permissions,
        *   Open feature/issue to discuss a method for webhook authorizers to support this API @ericchiang
    *   Remove automatic binding of system:node role ([#49638](https://github.com/kubernetes/kubernetes/pull/49638))
    *   RBAC v1 ([#49642](https://github.com/kubernetes/kubernetes/pull/49642))
*   Designs of note
    *   Giving nodes write access to pvc/status ([proposal](https://github.com/kubernetes/community/pull/657/files))
    *   kubeadm (alpha) HA implementation plan ([doc link](https://docs.google.com/document/d/1ff70as-CXWeRov8MCUO7UwT-MwIQ_2A0gNNDKpF39U4/edit#))
    *   Vault based KMS provider for secrets in etcd3 ([doc link](https://docs.google.com/document/d/15-baW4i7qws1yxxIYjHXqKpk259ebauQbECQCpPD308/edit#heading=h.67bgmqyjswzf), kk.sriram)
        *   This design is based off the in-flight implementation for Google Cloud KMS integration sakshamsharma has been writing (issue [#48522](https://github.com/kubernetes/kubernetes/issues/48522)), PRs: [#48574](https://github.com/kubernetes/kubernetes/pull/48574), [#49350](https://github.com/kubernetes/kubernetes/pull/49350),[ #49742](https://github.com/kubernetes/kubernetes/pull/49742)
*   Discussions
    *   Dynamic authz and audit webhook configuration (along the lines of [Dynamic Admission Control](https://kubernetes.io/docs/admin/extensible-admission-controllers/) in 1.7) (Matt Moyer)
        *   Admission is purely limiting and has the use case of dynamic aggregated API servers
        *   Clear use case? Could already do this by making your webhook handle the dynamic part.
        *   Overlap with component config? https://github.com/kubernetes/kubernetes/issues/12245
    *   Over next several releases, migrate webhook authn/authz to use v1 APIs (@liggitt to open issue)


## July 12th, 2017, 11a - Noon (Pacific Time)



*   [Recording](https://www.youtube.com/watch?v=9XAQriZ06CI&list=PL69nYSiGNLP0VMOZ-V7-5AchXTHAQFzJw&index=1)
*   Demo
    *   Liz Rice - [kube-bench](https://github.com/aquasecurity/kube-bench) demo - July 12th
        *   Get involved in CIS benchmarks: https://groups.google.com/forum/#!searchin/kubernetes-sig-auth/cis%7Csort:relevance/kubernetes-sig-auth/mJbmwzKTMqY/jDgZWM26DAAJ
        *   Will be updated with the 1.7 benchmark
        *   Has issues with different install tools
            *   Could use help with those default locations
*   Discussions:
*   API resources with Secret information: embed or reference secrets? (Paul Morie)
    *   Needed for [service-catalog work](https://github.com/kubernetes-incubator/service-catalog/issues/621)
    *   Original solution:
        *   Referencing a Secret and key containing confidential information
        *   svc-cat has a need to store secret data that is not going to be injected into Pod, but used as parameters to an external API
        *   Initial design had similar issues as ingress (watch all secrets) https://github.com/kubernetes/ingress/issues/816
    *   Access pattern: individual gets, looking forward to [bulk watch](https://github.com/kubernetes/community/pull/443)
        *   **Action item:** Need to formalize this "don't list/watch all secrets" opinion.
    *   Usability of managing permissions individually:
        *   How to express "grant access to secrets referenced by object X to subject Y"
        *   [Issues with label-aware RBAC](https://github.com/kubernetes/ingress/issues/816#issuecomment-309630254)
*   Additional TLS validation in the kubeadm bootstrap/discovery API ([WIP proposal](https://docs.google.com/document/d/1SP4P7LJWSA8vUXj27UvKdVEdhpo5Fp0QHNo4TXvLQbw/edit#)) (Matt Moyer)
    *   Proposal to add a CA signature pin to kubeadm join
        *   No specific objections to the pinning
        *   Discussion of auth token
*   1.8 planning/coordination
    *   [RBAC to v1](https://github.com/kubernetes/features/issues/2) - @liggitt
        *   Work with sig-apps to improve out-of-box experience with first-class helm charts, etc @ericchiang
        *   Disable legacy ABAC policy in GCE/GKE by default - @cjcullen?
        *   Documentation improvements (e.g. https://github.com/kubernetes/kubernetes.github.io/issues/2792) - @liggitt, @ericchiang, …?
    *   [NodeRestriction control over pod steering](https://github.com/kubernetes/features/issues/279) (prevent adding labels, prevent removing taints / deleting while tainted) - @liggitt
    *   [Audit logging beta](https://github.com/kubernetes/features/issues/22) - @tallclair, @ericchiang, …?
    *   TLS rotation ([client](https://github.com/kubernetes/features/issues/266), [server](https://github.com/kubernetes/features/issues/267)) beta @jcbsmpsn
    *   Container identity working group - @smarterclayton, @destijl
    *   PodSecurityPolicy enablement - @tallclair
    *   [Kubernetes Secrets: Roadmap](https://docs.google.com/document/d/1JAwPuZg47UhfRVlof-lMw08OJztunW8pvTNxDK3rCF8/edit#) work - @destijl, @smarterclayton, …?
        *   Add [support for storing secrets encryption key in Google KMS](https://github.com/kubernetes/kubernetes/issues/48522) @sakshamsharma, @jcbsmpsn
        *   Controller secret access patterns
        *   Generated secret injection into pods/containers, possibly via volume plugins
            *   Probably overlaps with container identity (service account token is attesting by ownership of the token that the agent can act as the service account, but could be much more finely scoped)
    *   SecurityContext improvements (entitlements?) roadmap - @tallclair


## June 28st, 2017, 11a - Noon (Pacific Time)



*   Canceled due to release crunch and lack of agenda items.


## June 14th, 2017, 11a - Noon (Pacific Time)



*   Not recorded
*   Pull requests of note:
    *   Azure client auth plugin merged -[ #43987](https://github.com/kubernetes/kubernetes/pull/43987)
*   Designs of note:
    *   Secrets roadmap? - gcastle?
        *   Didn't happen. Maybe next meeting?
*   Release Issues:
    *   [All issues marked for 1.7](https://github.com/kubernetes/kubernetes/issues?q=is%3Aopen+is%3Aissue+milestone%3Av1.7+label%3Asig%2Fauth)
    *   Update base images - [#47386](https://github.com/kubernetes/kubernetes/issues/47386)
    *   E2E testing of alpha features is broken - [#47283](https://github.com/kubernetes/kubernetes/issues/47283)
    *   PSP paths allow ".." - [#47107](https://github.com/kubernetes/kubernetes/issues/47107)
    *   GCE Load-Balancer Controller using insecure port - [#46983](https://github.com/kubernetes/kubernetes/issues/46983)
*   Demo
    *   Aqua Secrets - Amir Jerbi
        *   https://www.aquasec.com/products/aqua-container-security-platform/
        *   Federated secret store (is backed by vault, kms, etc.)
        *   Injects secrets into containers using a sidecar.
        *   http://blog.aquasec.com/injecting-secrets-kubernetes-hashicorp-vault-and-aqua-on-azure
        *   How does this fit into secrets roadmap?
            *   https://docs.google.com/document/d/1JAwPuZg47UhfRVlof-lMw08OJztunW8pvTNxDK3rCF8/edit
        *   Parts possibly going to be open source?
*   Questions
    *   Openstack auth
        *   Many of the current PRs/Issues for Authn/z backed by Keystone are either abandoned, experimental, or have issues with security.
        *   Openstack provider uses Keystone backed by MySQL. 250+ users, with 2,000+ tenants (tenant per app). Looking for the current recommendations on using Keystone as the source for user auth and roles.
        *   Would prefer if Keystone was backed by LDAP, but this is not possible in this case.
        *   Looking to discuss authn against Keystone, along with status of RBAC role mapping to keystone domains and groups.
    *   Openstack authn/z plugins should be developed out of tree and use the token webhook, authz webhook, and/or RBAC syncing (read policy from keystone and create associated RBAC rules)
        *   https://kubernetes.io/docs/admin/authentication/#webhook-token-authentication
        *   https://kubernetes.io/docs/admin/authorization/webhook/
        *   _Action item_: add better examples of webhook servers?


### Attending:



*   Quintin Lee, CJ Cullen, Saad Ali, Google
*   Luke Heidecke, Solinea
*   Eric Chiang, CoreOS
*   Chris Hoge, OpenStack Foundation (@hogepodge)
*   David Eads, Red Hat
*   Amir Jerbi, Aqua Security


## May 31st, 2017, 11a - Noon (Pacific Time)



*   [Recorded meeting](https://youtu.be/E_3L8dFaha0)
*   Pull requests of note:
    *   Kubelet serving cert creation/rotation merged - [#45059](https://github.com/kubernetes/kubernetes/pull/45059)
    *   Node authorizer merged - [#46076](https://github.com/kubernetes/kubernetes/pull/46076)
    *   [S](https://github.com/kubernetes/kubernetes/pull/40760)A[R](https://github.com/kubernetes/kubernetes/pull/40760) [A](https://github.com/kubernetes/kubernetes/pull/40760)p[p](https://github.com/kubernetes/kubernetes/pull/40760)r[o](https://github.com/kubernetes/kubernetes/pull/40760)v[e](https://github.com/kubernetes/kubernetes/pull/40760)r[ ](https://github.com/kubernetes/kubernetes/pull/40760)m[e](https://github.com/kubernetes/kubernetes/pull/40760)r[g](https://github.com/kubernetes/kubernetes/pull/40760)e[d](https://github.com/kubernetes/kubernetes/pull/40760) [-](https://github.com/kubernetes/kubernetes/pull/40760) [#45619](https://github.com/kubernetes/kubernetes/pull/40760)
    *   Kubelet client cert rotation - [#41912](https://github.com/kubernetes/kubernetes/pull/41912)
    *   Kubelet TLS bootstrap in GCE/GKE - [#40760](https://github.com/kubernetes/kubernetes/pull/40760)
*   Designs of note:
    *   Nothing new
*   End of release stuff:
    *   [Bugs](https://github.com/kubernetes/kubernetes/issues?q=is%3Aopen+is%3Aissue+label%3Asig%2Fauth+label%3Akind%2Fbug) (1)
    *   [Flakes](https://github.com/kubernetes/kubernetes/issues?q=is%3Aopen+is%3Aissue+label%3Asig%2Fauth+label%3Akind%2Fflake) (0)
    *   [PRs](https://github.com/kubernetes/kubernetes/pulls?q=is%3Aopen+is%3Apr+label%3Asig%2Fauth+milestone%3Av1.7)
*   Secrets roadmap
    *   Spectrum of use cases
        *   Value injected into container (easiest to solve, few API implications)
        *   "First class" use pulling from API (e.g. kubelet)
        *   High-privilege controller (e.g. cloud provider, less subdivision needed)
        *   Low-privilege controller (e.g. namespaced ingress controller)
    *   Dimensions to slice secret authorization by:
        *   Available today: Name, namespace
        *   Potential generic mechanisms: Secret type? Labels?
        *   Desired, but difficult to represent generically: relationships/references from other objects ("let the ingress controller get any secret referenced as a TLS cert from an Ingress object", etc)
    *   Mechanisms for authorizing based on object relationships:
        *   Bespoke in-tree authorizer (e.g. node authorizer)
            *   Pros: fast, low write impact to etcd, can be used in combination with any other authorization mode
            *   Cons: approach doesn't scale to wide variety of controllers/relationships (ingress, service catalog, etc, etc, etc)
        *   "Programming RBAC" based on object relationships, e.g. an "ingress role binder" that watches nIngress objects and grants get permissions on referenced TLS secrets to the ingress cotroller
            *   Pros: no core changes, policy is inspectable
            *   Cons: higher write traffic (may not matter for low change relationships), one more component to consider
    *   Problematic API access patterns for secret subdivision:
        *   list/watch
            *   would not be compatible with per-secret authorization strategies
            *   Straw-man guidance: avoid list/watch, prefer individual gets
        *   Lack of separation between metadata and confidential content
            *   Ideally could see a secret exists without seeing secret data
            *   Straw-man proposal:
                *   key/value field that is write-only (e.g. `superSecretData`)
                *   never included in get/list/watch responses
                *   only accessible via a subresource (e.g. `/api/v1/namespaces/myns/secrets/mysecret/data`), which means it is separately authorized.
                *   possible plug-point for external secret management integration (on write, plugin stores secure data, on subresource read, how the data is retrieved is determined by the plugin)


### Attending:



*   Saad Ali, Tim St. Clair, CJ Cullen, Mike Danese, Eric Tune, Jacob Simpson, Quintin Lee, Google
*   Andrew Jessup, Emiliano Berenbaum, Scytale
*   Steven E. Harris, IBM
*   Jordan Liggitt, Derek Carr, Paul Morie, Red Hat
*   Fred Vong, VMware
*   Alexander Brand, Apprenda
*   Joe Beda, Heptio


## May 17th, 2017, 11a - Noon (Pacific Time)



*   [Recorded meeting](https://youtu.be/9MdFd_9CwVw)
*   Pull requests of note:
    *   Mirror pod restriction/validation [#45775](https://github.com/kubernetes/kubernetes/pull/45775)
    *   NodeRestriction admission plugin [#45929](https://github.com/kubernetes/kubernetes/pull/45929)
        *   Proposal: https://github.com/kubernetes/community/blob/master/contributors/design-proposals/kubelet-authorizer.md
    *   Migrate group approver to use subject access review [#45619](https://github.com/kubernetes/kubernetes/pull/45619)
        *   TLS certificate request approving controller
        *   Types of requests: kubelet client cert (bootstrap), kubelet client cert (renew/rotate), kubelet serving cert
        *   Splits signing controller and autoapproving controller to be individually runnable
        *   Optional controller to auto-approve certain types of certificate requests (not required, can use out-of-band methods for approving requests)
    *   Kubeadm, put system certificates in kube-system secrets for self-hosted clusters: [#42548](https://github.com/kubernetes/kubernetes/pull/42548)
        *   Getting secrets in kube-system already game over today?
        *   Intrinsic tradeoff between self-hosted and isolation?
        *   Might help putting them in a different namespace?
        *   Would be good to support both models.
        *   Secrets API not compatible with chain of trust?
        *   Maybe not worse than the current situation? Just know that we're going to have to change this later.
        *   More secure today to shove it in a TPR?
        *   _Action item _(Greg Castle)_:_ Update secrets design doc to describe what secrets are for.
        *   _Action item _(Eric Tune): Setup call around scoping controller privs.
            *   Doc: https://docs.google.com/document/d/1KQNZmNMvViE4us44jIj8k-eOkhTGV21Sq42rNrizlGg/edit?ts=591caaeb#
*   Designs of note (may kubernetes/community PRs discussed in previous meetings):
    *   [Kubernetes Secrets: Updated Roadmap](https://docs.google.com/document/d/1JAwPuZg47UhfRVlof-lMw08OJztunW8pvTNxDK3rCF8/edit#).
    *   TL;DR - Try to get identity so we can use external stores.
*   Suggested working group on container identity
    *   Ensuring containers can assert their identity to other system has been discussed as necessary for a number of efforts, in order to try to get discussion going across the multiple impacted sigs it has been suggested that we try to do a working group.
    *   Try to follow the example of the resource management working group and solicit participants from other sigs and community
    *   Figure out what is out of scope.


### Attendees



*   Joe Beda, Heptio
*   Saad Ali, Google
*   Eric Tune, Google
*   CJ Cullen, Google
*   Yang Guan, Google
*   Quintin Lee, Google
*   Mike Danese, Google
*   Eric Chiang, CoreOS
*   Rithu John, CoreOS
*   Clayton Coleman, Red Hat
*   David Eads, Red Hat
*   Jordan Liggitt, Red Hat
*   Jacob Simpson, Google


## May 3rd, 2017, 11a - Noon (Pacific Time)



*   [Recorded meeting](https://www.youtube.com/watch?v=dLXz1T1yxXw)
*   Pull request of note:
    *   Certificate rotation for kubelet serving certs https://github.com/kubernetes/kubernetes/pull/45059
        *   Need to document the plan for having separate approval controllers
        *   Can platform dependent extensions be made accessible for proving who requested the signature.
        *   Categorization of a request: bootstrap client cert, kubelet requesting serving cert, kubelet requesting renewed client cert is generic.
    *   Issue for expanding approval mechanisms: https://github.com/kubernetes/kubernetes/issues/45030
    *
*   Designs of note (may kubernetes/community PRs discussed in previous meetings)
*   Individual identity in Docker Swarm ~20m - (Ying Li)
    *   Ying didn't join the hangout.
    *   Eric Chiang will follow up with the Swarm people and send something to sig-auth if they're going to present at another sig meeting.
*   Scoped kubelet API access - [#585](https://github.com/kubernetes/community/pull/585) - 10m - @liggitt
    *   Include other areas we'd like to extend it further.  For instance, PersistentVolumes.
    *   Should a kubelet be able to list all Pods in the system, or just Pods on itself?
    *   Ditto for every other Resource in the system (including TPR?)
*   ~~Greg castle: if clayton joins I'd be interested in talking about [#454](https://github.com/kubernetes/community/pull/454), encrypting secrets to see where it is at.~~
    *   Clayton is at Red Hat Summit.  I wouldn't anticipate him here.
*   [jbeda] If we have extra time and folks are interested, I'm happy to give a quick overview of [SPIFFE](https://docs.google.com/document/d/1GjurNK2ROw4rXz-k-l68JtpGRkGj2fZcWqP6gksEriQ/edit#heading=h.pq1kki84bhak).
    *   deads - Looks like it's worth a read, but I don't think we're ready to talk about it today.
*   Attendees:
*   Joe Beda, Heptio
*   Alex Mohr, Google
*   CJ Cullen, Google
*   Greg Castle, Google
*   Jacob Simpson, Google
*   Jacob Beacham, Google
*   Saad Ali, Google
*   Steve Sloka, UPMC Enterprises
*   David Eads, Red Hat
*   Jordan Liggitt, Red Hat
*   Eric Chiang, CoreOS
*   Rithu John, CoreOS
*   Tim St. Clair, Google
*   Quintin Lee, Google
*   Yang Guan, Google
*   Vipin Jain, SPIFFE
*   Emiliano Berenbaum, SPIFFE/Scytale


## April 19th, 2017, 11a - Noon (Pacific Time)



*   [Recorded meeting](https://www.youtube.com/watch?v=eT35tfepKGE)
*   Pull requests of note
    *   Reduce allocations by RBAC authorizer - [#44449](https://github.com/kubernetes/kubernetes/pull/44449) (@liggitt)
*   Designs of note
    *   [Auth setup for federation control plane](https://docs.google.com/document/d/1ioBg6Nn7PwyAAY48S42VQ0Zr1LXY5IAnMbBsftmE8DI/edit?ts=58f56a16) - (Nikhil Jindal)
        *   Per-cluster creds vs per-cluster-per-namespace vs per-user and per-cluster-per-user
        *   Tracking users adds complexity, but would allow federation controller to impersonate users.
    *   [Enterprise control for K8s](https://docs.google.com/document/d/1pH9jn4mj0Dap2LSEYFJDNP7JCE8ScIQoSGysfdzfv9s/edit#heading=h.gw8roqrqd77p) 10m, @erictune, Rae W (rae42), Ray C (rcolline)
        *   Not specifically tied to LDAP (just a strawman)
        *   Hierarchical set of namespaces that can be controlled, resource-quota'd, from a central database
        *   Had issues with "shadow namespaces" (developers who create untracked namespaces)
        *   Possibly no need to alter core kubernetes?
        *   Admission control?
            *   E.g. want to check LDAP before running a pod in a namespace
            *   Maybe can use pod security policies, etc. for these kind of things?
    *   Security model for user facing services - [#532](https://github.com/kubernetes/community/pull/532) (@Q-Lee)
*   Q: Is there a place for make working with policies easier? (e.g. RBAC, networking policy, pod security policy, etc.)
    *   Audit logging https://github.com/kubernetes/community/pull/145
    *   View the namespaces as your security boundary.
    *   Discussion around svc 2 svc auth.  Joe brought up [SPIFFE](https://spiffe.io) (again).


## April 5th, 2017, 11a - Noon (Pacific Time)



*   [Recorded meeting](https://www.youtube.com/watch?v=OfgnBOlSj_A&index=1&list=PL69nYSiGNLP0VMOZ-V7-5AchXTHAQFzJw)
*   Pull requests of note
    *   Brainstorm role name for read access to storageclasses - [#40881](https://github.com/kubernetes/kubernetes/pull/40881) (@seh)
    *   PodSecurityPolicy: limit which host paths are allowed - [#43946](https://github.com/kubernetes/kubernetes/pull/43946) (@jhorwit2)
    *   PodSecurityPolicy: allow within particular namespaces - [#42360](https://github.com/kubernetes/kubernetes/pull/42360) (@liggitt)
    *   Enable service account token lookup/revocation - [#24167](https://github.com/kubernetes/kubernetes/issues/24167), [#44071](https://github.com/kubernetes/kubernetes/pull/44071) (@liggitt)
        *   Performance implications of etcd lookups (can use cache if needed)
        *   HA/Federation implications if service account token use based solely on signature is desired (can always opt out of lookup)
        *   What is the least disruptive way to change this default?
            *   What matches user expectations best?
            *   What could changing the default cause to stop working?
    *   Include system:authenticated group in impersonated requests - [#44076](https://github.com/kubernetes/kubernetes/pull/44076) (@liggitt)
    *   Remove special-case handling for unsecured port - [#43888](https://github.com/kubernetes/kubernetes/pull/43888) (@liggitt)
*   Designs of note
    *   Threat models WIP - [#504](https://github.com/kubernetes/community/pull/504) (@smarterclayton)
    *   Security proposals/designs refresh - [#500](https://github.com/kubernetes/community/pull/500) (@smarterclayton)
*   Kubectl client auth plugins
    *   Azure, [#43987](https://github.com/kubernetes/kubernetes/pull/43987) (@cosmincojocar, @colemickens, @ericchiang)
        *   access_token vs id_token concerns
        *   can we delegate obtaining initial token to external tool, and just handle refreshing the token internally?
    *   Openstack, [#39587](https://github.com/kubernetes/kubernetes/pull/39587) (@zhouhaibing089)
        *   Retrying transport implementation is questionable
    *   Likely common patterns that could be factored out
        *   Respond to 401 API response by trying to obtain new credentials
*   CIS Kubernetes 1.6 Benchmark update (ericchiang)
    *   https://groups.google.com/forum/#!topic/kubernetes-sig-auth/mJbmwzKTMqY
    *   Lots of suggested settings (tickets) that could use input from sig-auth members
    *   Timeframe is 6-8 weeks to completion
*   SIG business item (@erictune)
    *   @erictune stepping down as sig-auth lead, focusing more on sig-apps
    *   No immediate plans to seek a replacement, [most sigs](https://github.com/kubernetes/community/blob/master/sig-list.md) have 2-3 leads (though feedback on that is welcome)


## March 22nd, 2017, 11a - Noon (Pacific Time)



*   [Recorded meeting](https://www.youtube.com/watch?v=SyU6aerXUKo&list=PL69nYSiGNLP0VMOZ-V7-5AchXTHAQFzJw&index=1)
*   1.6 close-out
*   Discussion re https://groups.google.com/forum/#!topic/kubernetes-sig-auth/diw9LH67iIs
    *   Keep kube-up.sh running with legacy ABAC policy by default ([#43544](https://github.com/kubernetes/kubernetes/pull/43544) cjcullen, liggitt)
    *   Continue running some CI jobs with RBAC-only to ensure coverage
    *   Run upgrade CI jobs with legacy ABAC policy to ensure coverage ([#2330](https://github.com/kubernetes/test-infra/pull/2330), liggitt)
    *   Follow ups:
        *   Document intended "secure" profile? (clayton)
        *   Warning about starting in insecure mode (no good place to surface in kube-up today)
        *   "security profile" options ("insecure", "secured", "multitenant", etc)?
*   Release notes
    *   Kubeadm's release notes need more details? (ericchiang will follow up)
    *   kube-up.sh on GCE enables rbac and insecure abac policy by default. Needs to document option to start kube-up.sh without legacy ABAC policy (liggitt/cjcullen)
*   Doc PRs
    *   https://github.com/kubernetes/kubernetes.github.io/pull/2940/
    *   https://deploy-preview-2940--kubernetes-io-master-staging.netlify.com/docs/admin/authorization/rbac/#service-account-permissions


## March 8th, 2017, 11a - Noon (Pacific Time)



*   [Recorded meeting](https://www.youtube.com/watch?v=ri6BVhCMKB8)
*   1.6 close-out
    *   RBAC
        *   [documentation available in preview](https://kubernetes-io-vnext-staging.netlify.com/docs/admin/authorization/rbac/)
        *   Enabled by default in kubeadm, kube-up, bootkube
        *   kops is [proposing](https://github.com/kubernetes/kops/pull/2039/files#diff-38574c080d4e2eb38c49b86e6588ad98R8) to enable RBAC by default (their 1.6 lags k8s)
    *   Secured Kubelet API
        *   Enabled by default in kubeadm, kube-up, bootkube
    *   [Flakes](https://github.com/kubernetes/kubernetes/issues?q=is%3Aissue+is%3Aopen+label%3Asig%2Fauth+label%3Akind%2Fflake+milestone%3Av1.6): 0
    *   [Bugs](https://github.com/kubernetes/kubernetes/issues?utf8=%E2%9C%93&q=%20is%3Aopen%20label%3Asig%2Fauth%20label%3Akind%2Fbug%20milestone%3Av1.6%20): 1
*   1.7 plans / priorities?
    *   Documentation
        *   Threat model @timstclair (node), ...
        *   Hardening guide
        *   User-focused guides @rithujohn191
    *   Deployments
        *   Need sig-auth members to work with each deployment project
        *   Help implement best practices (setup, hardening, etc)
        *   kube-up
        *   kubeadm
        *   kops
        *   …
    *   Add-ons
        *   Almost all add-ons run as default kube-system service account
        *   Need to work with add-on manager devs, and add-on devs to define best practices - [#38541](https://github.com/kubernetes/kubernetes/issues/38541)
        *   Partitioned namespaces, service accounts, accompanying roles, etc
    *   Implementation work
        *   Node authorizer - [#40476](https://github.com/kubernetes/kubernetes/issues/40476), @liggitt
        *   PodSecurityPolicy improvements - [#23217](https://github.com/kubernetes/kubernetes/issues/23217), @liggitt
        *   CSR auto-approval switch from hard-coded group to authz check using SubjectAccessReview - @liggitt
        *   Encrypt secrets at rest - [#12742](https://github.com/kubernetes/kubernetes/issues/12742), @erictune, @smarterclayton
        *   Groupification - [#23720](https://github.com/kubernetes/kubernetes/issues/23720#issuecomment-284032814), @ericchiang
        *   OIDC plugin maintenance - [#42654](https://github.com/kubernetes/kubernetes/issues/42654), @ericchiang
        *   Kubelet server cert bootstrapping @jcbsmpsn
        *   Kubelet client certificate rotation [#4672](https://github.com/kubernetes/kubernetes/issues/4672), @jcbsmpsn
        *   Audit Logging [features#22](https://github.com/kubernetes/features/issues/22), @timstclair, @cjcullen, @ericchiang, others?
        *   Seccomp enabled by default [#39845](https://github.com/kubernetes/kubernetes/issues/39845) , [#145](https://github.com/kubernetes/community/pull/145) @timstclair
    *   Design/Roadmap work
        *   Bulk namespace access control [#40403](https://github.com/kubernetes/kubernetes/issues/40403) @cmluciano, @smarterclayton
        *   [Roadmap](https://docs.google.com/document/d/1T2y-9geg9EfHHtCDYTXptCa-F4kQ0RyiH-c_M1SyD0s/edit?ts=58bf234e#) for secrets, and integration with external secret stores @erictune
        *   [Extensible admission](https://github.com/kubernetes/community/pull/132) (in concert with sig-api-machinery)
    *   Backlog
        *   CRL support - [#33519](https://github.com/kubernetes/kubernetes/pull/33519) (might be mitigated by rotation / shorter expirys)


## February 22nd, 2017, 11a - Noon (Pacific Time)



*   Meeting cancelled


## February 8th, 2017, 11a - Noon (Pacific Time)



*   [Recorded meeting](https://www.youtube.com/watch?v=DZShqi2eDOc)
*   Representative from sig-auth to PM group (5 minutes, cjcullen)
*   CIS security benchmark for kubernetes (15 minutes, praving5)
    *   https://groups.google.com/d/msgid/kubernetes-sig-auth/da0aba5b-9285-487c-b236-34ebddfd0d8a%40googlegroups.com
*   SubjectAccessReview, TokenReview promotion to v1 (5 minutes, liggitt)
    *   https://github.com/kubernetes/kubernetes/pull/40709
    *   https://github.com/kubernetes/kubernetes/pull/41058
*   Auto-reconciliation of bootstrap roles/rolebindings (5 minutes, liggitt)
    *   https://github.com/kubernetes/kubernetes/pull/41155
    *   Add missing permissions to bootstrap roles on server start
    *   Add missing subjects to bootstrap rolebindings on server start
    *   No auto-removal of permissions or subjects
    *   Can opt out of auto-reconciliation by annotating roles/rolebindings (become responsible for manually updating with new permissions/subjects)
    *   Eventually could be moved to a controller loop


## January 25th, 2017, 11a - Noon (Pacific Time)



*   [Recorded meeting](https://www.youtube.com/watch?v=5mYdRXR9ZfM&list=PL69nYSiGNLP0VMOZ-V7-5AchXTHAQFzJw)
*   1.5 cherry picks
    *   OIDC client-side performance fix
        *   https://github.com/kubernetes/kubernetes/pull/40191
    *   Fix for requestheader proxy client cert prompt
        *   https://github.com/kubernetes/kubernetes/pull/40301
*   RBAC v1beta1 released, API enabled by default (liggitt)
    *   https://github.com/kubernetes/kubernetes/pull/39625
*   RBAC rollout for 1.6 (liggitt, cjcullen)
    *   Doc overview - https://kubernetes.io/docs/admin/authorization/#rbac-mode
    *   Built-in role documentation (in progress) - https://github.com/kubernetes/kubernetes.github.io/pull/2169/files?short_path=48e69c0#diff-48e69c0b942ef9dcc93b90046d09f9e6
    *   kubeadm enablement
        *   https://github.com/kubernetes/kubernetes/pull/39846
    *   kops enablement
        *   https://github.com/kubernetes/kops/pull/1357/files
    *   kube-up.sh with GCE provider enables RBAC by default
        *   ABAC enablement on upgrade in progress
    *   GKE enablement for 1.6 in-progress
        *   Should new clusters provide an option to start with legacy ABAC?
        *   Should new clusters provide an option to start with permissive RBAC?
    *   Default roles and role-bindings being added in-tree:
        *   https://github.com/kubernetes/kubernetes/tree/master/plugin/pkg/auth/authorizer/rbac/bootstrappolicy/testdata
    *   Docs in progress (working with docs team to find a home for 1.6 docs)
*   Multidimensional multi-tenancy (mspreitz, clayton, liggitt)
    *   Per-namespace DNS: https://github.com/kubernetes/community/pull/269
    *   API isolation (API authorization)
    *   Service isolation (DNS and/or network policy, explicit service references)
    *   Network isolation (network policy)
    *   Interested sigs:
        *   apimachinery (for watching across a subset of namespaces)
        *   networking (for dns and kube-proxy implications, API shape)
        *   service-catalog (for explicit service references)
    *   many:many visibility
    *   Problem space:
        *   One namespace per cluster
        *   Possibly have a shared one that all namespaces need to hit.
        *   How many "tenants", how big?
            *   Performance issues with kube-dns if you have a lot of tenants?
    *   Policy aware DNS server vs per-namespace DNS server.
    *   In openshift, kube-proxy is the DNS server.
        *   Have to go through kube-proxy to talk to service, extends nicely to DNS.
        *   Kube-proxy would have to understand policy too.
    *   Network policy is only pod selection and only ingress today.
        *   Maybe some sort of API that declares "I consume these service"?
    *   Are service accounts a good proxy for a pod identity?
        *   Could we say something like "if a pod's service account can see a service it can access the service"?
*   Efficient subdivision of namespaces into groups (tenancy of organizations)
    *   https://github.com/kubernetes/kubernetes/issues/40403
    *   efficient multi-namespace watch and authorization
*   SIG Auth update to Kubernetes Community
    *   liggitt will present RBAC status for 1.6
*   Discussions about secrets
    *   Encryption at rest (in etcd).
        *   https://github.com/kubernetes/kubernetes/issues/12742
        *   https://github.com/kubernetes/features/issues/92
    *   Subdivision access (ingress vs build vs serviceaccount)
        *   https://github.com/kubernetes/kubernetes/issues/18725
        *   https://github.com/kubernetes/kubernetes/issues/4957
    *   External injection without visibility in the API (e.g. Vault)
        *   https://github.com/kubernetes/kubernetes/issues/10439
        *   https://github.com/kubernetes/kubernetes/issues/28538
*   Future work: subdivided node access
    *   Currently able to read all secrets, modify all pods
    *   Should limit to resources needed by pods bound to the node
    *   https://github.com/kubernetes/kubernetes/issues/40476
    *   Ericchiang: maybe related to [40403](https://github.com/kubernetes/kubernetes/issues/40403)?


## January 11th, 2017, 11a - Noon (Pacific Time)



*   RBAC v1beta1
    *   https://github.com/kubernetes/kubernetes/pull/39625
    *   attributeRestrictions removed.
    *   Doc pull with default role explanation, command grant explanations: https://github.com/kubernetes/kubernetes.github.io/pull/2169
*   RBAC rollout (deads, cjcullen, liggitt)
    *   Existing bring-up scripts used for CI and production
        *   https://github.com/kubernetes/kubernetes/issues/39532
        *   https://github.com/kubernetes/kubernetes/pull/39537
    *   Upgrades of existing clusters using ABAC
        *   Optionally include existing ABAC policy files
        *   Define permissive RBAC policy matching legacy defaults (but do not use in CI/e2e)
    *   Flushing out required permissions for system components
        *   https://github.com/kubernetes/kubernetes/issues/39639
        *   https://github.com/kubernetes/kubernetes/pull/39641
    *   Example and doc updates
        *   Cluster admin oriented
            *   Delegating permissions per namespace
            *   Giving permissions to a particular service account in a namespace
            *   Giving permissions to a particular service account cluster-wide
            *   Creating a custom role
        *   Developer oriented
            *   Consider permissions required by pod
            *   Describe default RBAC roles (cluster-admin, admin, edit, view) - https://github.com/kubernetes/kubernetes.github.io/pull/2169
            *   Encourage documenting which RBAC role is required by your component
            *   If no default role fits, include a tailored role definition
    *   Deployment methods (kops, kubeadm, ansible, …, …)
        *   Compatible upgrades
        *   Default mode for new installs
        *   Options to load legacy permissive policies?
    *   For e2e troubles, here's a thread with helpful links
        *   Controller permissions, e2e new namespace permissions, kubelet permissions, etc
        *   https://groups.google.com/forum/#!msg/kubernetes-dev/gASoXHp52R8/3sEBdwGGFQAJ;context-place=forum/kubernetes-dev
    *   Want to come up with an education plan.
        *   In some cases, this exposes problems that already existed.  Bad extensions as a for instance, lack of safe multi-tenancy as another.
*   RBAC role granting allowed via "bind" permission from any authorizer
    *   https://github.com/kubernetes/kubernetes/pull/39383
