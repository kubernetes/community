# Kubernetes Sig-auth Meeting Agenda

## December 25, 11a - Noon (Pacific Time)

*   Canceled due to winter holidays


## December 11, 11a - Noon (Pacific Time)



*   [Recording](https://youtu.be/SpVv7wp8moI)
*   Announcements
    *   
*   Demos
    *   
*   Designs of note
    *   CSR multi-signer proposal - https://github.com/kubernetes/enhancements/pull/1400 [@deads2k]
*   Pulls of note
    *   https://github.com/kubernetes/kubernetes/pull/79083 - client certificate rotation [jackkleeman]
*   Issues of note
    *   
*   Discussion topic
    *   [gcastle] Bringing an internal Google discussion to sig-auth audience: can we do better than file-system mounted secrets inside core K8s? Ideally filesystem information leaks (directory traversal bugs) wouldn’t result in secrets lost to bad guys. Christoph Kern to join us as a guest speaker.
        *   Desire to provide applications with an alternate way to obtain secrets without exposing secrets in files easily vulnerable to directory traversal
        *   Potential directions
            *   option to encrypt + store encryption keys in an xattr. https://github.com/kubernetes/kubernetes/issues/81125#issuecomment-521529601
            *   option to remove secret file content after container is ready
    *   [seh] Related: Following [discussion on Slack](https://kubernetes.slack.com/archives/C0EN96KUY/p1575474009160600), how can CSI drivers more securely acquire service account tokens from the pods into which they’re mounting? Mike Danese suggested that the kubelet should fetch attenuated tokens to hand to the CSI driver, rather than either forcing the driver to grovel around on the filesystem looking for tokens or having the driver request the tokens via the _TokenRequest_ API.
        *   seh: look around the filesystem :(
        *   Tim: CSI driver creates TokenRequest directly
        *   Global CSI Driver object. Is that a good spot? Are those objects usually locked down? Kubelet passes CSI Driver token with some globally configured audience.
*   Action Items
    *   
*   Sweep issues with leftover time
    *   [unprioritized+unassigned (-needs-information) issues](https://github.com/kubernetes/kubernetes/issues?q=is%3Aissue+is%3Aopen+label%3Asig%2Fauth+-label%3Apriority%2Fawaiting-more-evidence+-label%3Apriority%2Fimportant-longterm+-label%3Apriority%2Fimportant-soon+-label%3Apriority%2Fcritical-urgent+-label%3Apriority%2Fbacklog+no%3Aassignee+-label%3Atriage%2Fneeds-information)
    *   [unprioritized+assigned (-needs-information) issues](https://github.com/kubernetes/kubernetes/issues?q=is%3Aissue+is%3Aopen+label%3Asig%2Fauth+-label%3Apriority%2Fawaiting-more-evidence+-label%3Apriority%2Fimportant-longterm+-label%3Apriority%2Fimportant-soon+-label%3Apriority%2Fcritical-urgent+-label%3Apriority%2Fbacklog+-no%3Aassignee+-label%3Atriage%2Fneeds-information)
    *   [open bugs](https://github.com/kubernetes/kubernetes/issues?utf8=✓&q=is%3Aissue+is%3Aopen+label%3Asig%2Fauth+label%3Akind%2Fbug)


## November 27, 11a - Noon (Pacific Time)



*   Canceled due to proximity with Thanksgiving


## November 13, 11a - Noon (Pacific Time)



*   Canceled due to proximity with 1.17 code freeze


## October 30, 11a - Noon (Pacific Time)



*   [Recording](https://www.youtube.com/watch?v=DLgAG3cYiSM)
*   Announcements
    *   
*   Demos
    *   
*   Pulls of note
    *   
*   Issues of note
    *   
*   Designs of note
    *   
*   Discussion topic
    *   Future of PSPs, alternatives, migration
        *   What is the 90% use case?
        *   Corner cases / exceptions:
            *   Windows
            *   Sandboxed RuntimeClasses
        *   
*   Action Items
    *   
*   Sweep issues with leftover time
    *   [unprioritized+unassigned (-needs-information) issues](https://github.com/kubernetes/kubernetes/issues?q=is%3Aissue+is%3Aopen+label%3Asig%2Fauth+-label%3Apriority%2Fawaiting-more-evidence+-label%3Apriority%2Fimportant-longterm+-label%3Apriority%2Fimportant-soon+-label%3Apriority%2Fcritical-urgent+-label%3Apriority%2Fbacklog+no%3Aassignee+-label%3Atriage%2Fneeds-information)
    *   [unprioritized+assigned (-needs-information) issues](https://github.com/kubernetes/kubernetes/issues?q=is%3Aissue+is%3Aopen+label%3Asig%2Fauth+-label%3Apriority%2Fawaiting-more-evidence+-label%3Apriority%2Fimportant-longterm+-label%3Apriority%2Fimportant-soon+-label%3Apriority%2Fcritical-urgent+-label%3Apriority%2Fbacklog+-no%3Aassignee+-label%3Atriage%2Fneeds-information)
    *   [open bugs](https://github.com/kubernetes/kubernetes/issues?utf8=✓&q=is%3Aissue+is%3Aopen+label%3Asig%2Fauth+label%3Akind%2Fbug)


## October 16, 11a - Noon (Pacific Time)



*   [Recording](https://youtu.be/6DbmamBBWWE)
*   Announcements
    *   oss-fuzz and test/fuzz/
        *   https://github.com/google/oss-fuzz
*   Demos
    *   
*   Pulls of note
    *   
*   Issues of note
    *   [Dynamic Audit Policy](https://github.com/kubernetes/enhancements/pull/1259)
        *   we've started this design a few times, and ended up with very complex approaches multiple times
        *   Tim wanted to [ask some more fundamental questions](https://github.com/kubernetes/enhancements/pull/1259#issuecomment-542008989) to help focus the design
        *   to make progress, likely need to have dedicated synchronous time to work through the design
            *   Should we schedule some time at KC NA?
        *   dimensions:
            *   noise (e.g. events)
                *   cluster admin filters things they consider noisy (or resource authors?)
                *   webhooks can opt into more noise?
            *   sensitivity (e.g. secret request/response contents?)
                *   indicate sensitivity per resource?
                    *   built-in resources like secrets
                    *   custom resource authors
                *   could inform things like
                    *   encryption at rest
                    *   allowed audit levels for webhook
                *   cluster admin indicates things they consider sensitive that may not be sent to audit webhooks?
                *   webhooks cannot opt into receiving sensitive info the cluster admin has disallowed?
                *   need to be careful not to mislead webhooks (they think they registered to receive events, but don't get them at all because the cluster-admin disabled a specific resource)
        *   areas of focus for debug
            *   everything (interactions between namespaces/users)
            *   namespace-focused
            *   user-focused
        *   Action items:
            *   look at previously discussed use cases in light of the noise/sensitivity dimensions
            *   consider whether use cases could be addressed with a trusted sink that receives everything and filters (identify gaps that would prevent that approach). most likely limit of backend-based filtering would be scale, but if an API and implementation was proved as a backend, it could be brought into the API server to address scale concerns
            *   follow up with api-machinery on possibility of indicating sensitivity at a resource level
*   Designs of note
    *   [Credential Provider Extraction KEP](https://github.com/kubernetes/enhancements/pull/1284)
*   Discussion topic
    *   
*   Action Items
    *   
*   Sweep issues with leftover time
    *   [unprioritized+unassigned (-needs-information) issues](https://github.com/kubernetes/kubernetes/issues?q=is%3Aissue+is%3Aopen+label%3Asig%2Fauth+-label%3Apriority%2Fawaiting-more-evidence+-label%3Apriority%2Fimportant-longterm+-label%3Apriority%2Fimportant-soon+-label%3Apriority%2Fcritical-urgent+-label%3Apriority%2Fbacklog+no%3Aassignee+-label%3Atriage%2Fneeds-information)
    *   [unprioritized+assigned (-needs-information) issues](https://github.com/kubernetes/kubernetes/issues?q=is%3Aissue+is%3Aopen+label%3Asig%2Fauth+-label%3Apriority%2Fawaiting-more-evidence+-label%3Apriority%2Fimportant-longterm+-label%3Apriority%2Fimportant-soon+-label%3Apriority%2Fcritical-urgent+-label%3Apriority%2Fbacklog+-no%3Aassignee+-label%3Atriage%2Fneeds-information)
    *   [open bugs](https://github.com/kubernetes/kubernetes/issues?utf8=✓&q=is%3Aissue+is%3Aopen+label%3Asig%2Fauth+label%3Akind%2Fbug)


## October 2, 11a - Noon (Pacific Time)



*   [Recording](https://youtu.be/idX8G_DbFJU)
*   Announcements
    *   
*   Demos
    *   
*   Pulls of note
    *   
*   Issues of note
    *   
*   Designs of note
    *   [OIDC Issuer URL](https://github.com/kubernetes/enhancements/pull/1190)
    *   [kubectl logs --insecure-skip-tls-verify](https://github.com/kubernetes/enhancements/pull/1261)
*   Discussion topic
    *   Deprecate ABAC [k/k#82540](https://github.com/kubernetes/kubernetes/pull/82540)
        *   AI(mo): send the deprecation email to kubernetes-dev
    *   Need approval to move secrets-store csi driver to kubernetes-sigs. [Issue link](https://github.com/kubernetes/org/issues/1245)
        *   AI(mikedanese): List pros and cons of inclusion, share findings with the list.
            *   Pros:
                *   Community discussion and contribution.
                *   
*   Action Items
    *   
*   Sweep issues with leftover time
    *   [unprioritized+unassigned (-needs-information) issues](https://github.com/kubernetes/kubernetes/issues?q=is%3Aissue+is%3Aopen+label%3Asig%2Fauth+-label%3Apriority%2Fawaiting-more-evidence+-label%3Apriority%2Fimportant-longterm+-label%3Apriority%2Fimportant-soon+-label%3Apriority%2Fcritical-urgent+-label%3Apriority%2Fbacklog+no%3Aassignee+-label%3Atriage%2Fneeds-information)
    *   [unprioritized+assigned (-needs-information) issues](https://github.com/kubernetes/kubernetes/issues?q=is%3Aissue+is%3Aopen+label%3Asig%2Fauth+-label%3Apriority%2Fawaiting-more-evidence+-label%3Apriority%2Fimportant-longterm+-label%3Apriority%2Fimportant-soon+-label%3Apriority%2Fcritical-urgent+-label%3Apriority%2Fbacklog+-no%3Aassignee+-label%3Atriage%2Fneeds-information)
    *   [open bugs](https://github.com/kubernetes/kubernetes/issues?utf8=✓&q=is%3Aissue+is%3Aopen+label%3Asig%2Fauth+label%3Akind%2Fbug)


## September 18, 11a - Noon (Pacific Time)



*   [Recording](https://youtu.be/Gl6-i-yGKXg)
*   Issues of note
    *   [Kubernetes 3rd party audit findings aggregate issue](https://github.com/kubernetes/kubernetes/issues/81146)
*   Designs of note
    *   [Extended NodeRestrictions for Pods](https://github.com/kubernetes/enhancements/pull/1243)
        *   question about whether we should enforce policy on labels outside k8s.io/kubernetes.io prefixes (node self-labeling doesn't enforce policy 
        *   limiting ownerReferences makes sense (mirror pods shouldn't have controller:true references)
        *   limiting annotations seems speculative. should we lean on webhook admission protections instead?
        *   compatibility questions (things like network plugins setting annotations on pod status update using kubelet creds?)
    *   [Certificates KEP](https://github.com/kubernetes/enhancements/pull/1097)
        *   ready for initial merge to capture current state
        *   follow up to work through GA/v1 requirements
*   Discussion topic
    *   1.17
        *   Token request to GA?
            *   Merged KEP
            *   Update docs
                *   API/usage doc
                *   Define structure of SA tokens as not opaque
                *   Key ID docs?
            *   Promote existing e2es to conformance?
                *   Requires API server invocation changes to be conformant
            *   Gather feedback on usage
                *   Istio
                *   AWS
                *   linkerd wants to
            *   Feature requests
                *   arbitrary claims ([#61795](https://github.com/kubernetes/kubernetes/issues/61795))
                *   pluggable signer backend ([#704](https://github.com/kubernetes/enhancements/pull/704))
            *   Vault CSI driver
                *   In progress
            *   Vault auth backend
                *   Change token review url to point at new version
        *   Plans for SA token volume projection?
            *   Performance numbers around large clusters
            *   volume projection issues with permissions/fsGroups interactions
            *   client readiness (refreshing behavior across libraries)
                *   Some numbers around community adoption?
            *   docs/breadcrumbs for people encountering new behavior
                *   improve message from service account token authenticator when using an expired token to point to solution
                *   ensure docs for components needing to update their client libraries that match the more informative error message so they are discoverable
                *   ensure docs for users encountering expired errors second-hand via apps/components they do not control
            *   method for gathering metrics about whether workloads are refreshing tokens correctly?
                *   maybe mint tokens with long lifetimes, but continue refreshing every 10 minutes, expose metrics or audit info when tokens older than 10 minutes are presented (means a particular workload isn't refreshing tokens correctly)
                *   Expose metrics on successful refreshes
        *   Dynamic audit policy
        *   [Extended NodeRestrictions for Pods](https://github.com/kubernetes/enhancements/pull/1243)
        *   [External Projected Token Creation](https://github.com/kubernetes/enhancements/pull/704)
        *   [OIDC Issuer URL](https://github.com/kubernetes/enhancements/pull/1190) (link fixed)
    *   Dynamic audit policy: [call for comments](https://docs.google.com/document/d/12uYuvykipkG96EJ4PsFtvhzgMY6mmsBgX8kNLuOq_RE/edit?usp=sharing), short proposal overview, roadmap /can do that in the slot above too - Georgi
*   Action Item:
    *   Mike to file an issue for things that need to be done for Vault kubernetes-credential-backend
*   Sweep issues with leftover time
    *   [unprioritized+unassigned (-needs-information) issues](https://github.com/kubernetes/kubernetes/issues?q=is%3Aissue+is%3Aopen+label%3Asig%2Fauth+-label%3Apriority%2Fawaiting-more-evidence+-label%3Apriority%2Fimportant-longterm+-label%3Apriority%2Fimportant-soon+-label%3Apriority%2Fcritical-urgent+-label%3Apriority%2Fbacklog+no%3Aassignee+-label%3Atriage%2Fneeds-information)
    *   [unprioritized+assigned (-needs-information) issues](https://github.com/kubernetes/kubernetes/issues?q=is%3Aissue+is%3Aopen+label%3Asig%2Fauth+-label%3Apriority%2Fawaiting-more-evidence+-label%3Apriority%2Fimportant-longterm+-label%3Apriority%2Fimportant-soon+-label%3Apriority%2Fcritical-urgent+-label%3Apriority%2Fbacklog+-no%3Aassignee+-label%3Atriage%2Fneeds-information)
    *   [open bugs](https://github.com/kubernetes/kubernetes/issues?utf8=✓&q=is%3Aissue+is%3Aopen+label%3Asig%2Fauth+label%3Akind%2Fbug)


## September 4, 2019, 11a - Noon (Pacific Time)



*   Cancelled, empty agenda


## August 21, 2019, 11a - Noon (Pacific Time)



*   [Recording](https://youtu.be/VfOCeAnHh7w)
*   Discussion topic
    *   wg-policy update
*   Sweep issues with leftover time
    *   [unprioritized+unassigned (-needs-information) issues](https://github.com/kubernetes/kubernetes/issues?q=is%3Aissue+is%3Aopen+label%3Asig%2Fauth+-label%3Apriority%2Fawaiting-more-evidence+-label%3Apriority%2Fimportant-longterm+-label%3Apriority%2Fimportant-soon+-label%3Apriority%2Fcritical-urgent+-label%3Apriority%2Fbacklog+no%3Aassignee+-label%3Atriage%2Fneeds-information)
    *   [unprioritized+assigned (-needs-information) issues](https://github.com/kubernetes/kubernetes/issues?q=is%3Aissue+is%3Aopen+label%3Asig%2Fauth+-label%3Apriority%2Fawaiting-more-evidence+-label%3Apriority%2Fimportant-longterm+-label%3Apriority%2Fimportant-soon+-label%3Apriority%2Fcritical-urgent+-label%3Apriority%2Fbacklog+-no%3Aassignee+-label%3Atriage%2Fneeds-information)
    *   [open bugs](https://github.com/kubernetes/kubernetes/issues?utf8=✓&q=is%3Aissue+is%3Aopen+label%3Asig%2Fauth+label%3Akind%2Fbug)


## August 7, 2019, 11a - Noon (Pacific Time)



*   [Recording](https://youtu.be/njCUdT5I7sM)
*   Announcements
    *   Contributor Summit! Nov 17th and 18th before KubeCon/CloudNativeCon. We need security and other lovely sig-auth content! What do you need to share with 400 contributors? Teach them? Security audit review? https://forms.gle/Fez7U8UZAzVxykNW9
*   Demos
    *   
*   Pulls of note
    *   
*   Issues of note
    *   
*   Designs of note
    *   [Mike Danese] Retroactive KEPs for TokenRequest and External Auth Plugin
        *   [links]
    *   [cceckman] KEP for OIDC discovery ([link](https://github.com/kubernetes/enhancements/pull/1190))
        *   Proof of concept implementation [here](https://github.com/kubernetes/kubernetes/pull/80724)
*   Discussion topic
    *   [Rita Zhang] Should we add https://github.com/deislabs/secrets-store-csi-driver to kubernetes-sigs org? Sig sponsor?
        *   We need criteria
            *   Identify owners
            *   Make sure that we have alignment with the sig-auth charter
            *   Review the current state of the project
        *   Rita Zhang to start a thread on kubernetes-sig-auth@googlegroups.com
    *   [Georgy Pavlov] [Audit policy use cases & requirements](https://docs.google.com/document/d/1kpAldrEh9T0J8LELNc7r6IMRF3MpwDZVZ_ALtXOeA94/edit?usp=sharing) - take 1
        *   Need to decide if this needs to be included beta for dynamic audit
    *   [Tim Allclair] [WG Security Process Proposal](https://docs.google.com/document/d/1BnCG_JopFD-XRZKPQjNHrG26G8V6dGC_enCnHw_v5bI/edit#heading=h.2buup6evwogx)
        *   sig-auth to sponsor?
*   Action Items
    *   
*   Sweep issues with leftover time
    *   [unprioritized+unassigned (-needs-information) issues](https://github.com/kubernetes/kubernetes/issues?q=is%3Aissue+is%3Aopen+label%3Asig%2Fauth+-label%3Apriority%2Fawaiting-more-evidence+-label%3Apriority%2Fimportant-longterm+-label%3Apriority%2Fimportant-soon+-label%3Apriority%2Fcritical-urgent+-label%3Apriority%2Fbacklog+no%3Aassignee+-label%3Atriage%2Fneeds-information)
    *   [unprioritized+assigned (-needs-information) issues](https://github.com/kubernetes/kubernetes/issues?q=is%3Aissue+is%3Aopen+label%3Asig%2Fauth+-label%3Apriority%2Fawaiting-more-evidence+-label%3Apriority%2Fimportant-longterm+-label%3Apriority%2Fimportant-soon+-label%3Apriority%2Fcritical-urgent+-label%3Apriority%2Fbacklog+-no%3Aassignee+-label%3Atriage%2Fneeds-information)
    *   [open bugs](https://github.com/kubernetes/kubernetes/issues?utf8=✓&q=is%3Aissue+is%3Aopen+label%3Asig%2Fauth+label%3Akind%2Fbug)


## July 24, 2019, 11a - Noon (Pacific Time)



*   [Recording](https://youtu.be/zq81PYjaiVc)
*   Announcements
    *   Requesting feedback on what folks would like to see covered in KubeCon NA 2019 (San Diego) SIG Auth deep dive (please respond to thread on mailing list)
    *   Options
        *   Walkthrough of OPA/gatekeeper?
            *   Refactor of PSP
            *   [Mo] what does policy wg talk about?
            *   PSP might be a good option to talk about at contributor summit.
        *   Authentication (client exec plugin, webhook)
        *   RBAC?
            *   Maybe not a topic by itself
        *   Certificates API, approvers
        *   Audit webhook
            *   Audit policy API
        *   Node authorizer, node admission
*   Demos
    *   
*   Pulls of note
    *   
*   Issues of note
    *   
*   Designs of note
    *   
*   Discussion topic
    *   Audit policy use cases & requirements
        *   Doc going to mailing list
        *   Discussion punted for next meeting
*   Action Items
    *   
*   Sweep issues with leftover time
    *   [unprioritized+unassigned (-needs-information) issues](https://github.com/kubernetes/kubernetes/issues?q=is%3Aissue+is%3Aopen+label%3Asig%2Fauth+-label%3Apriority%2Fawaiting-more-evidence+-label%3Apriority%2Fimportant-longterm+-label%3Apriority%2Fimportant-soon+-label%3Apriority%2Fcritical-urgent+-label%3Apriority%2Fbacklog+no%3Aassignee+-label%3Atriage%2Fneeds-information)
    *   [unprioritized+assigned (-needs-information) issues](https://github.com/kubernetes/kubernetes/issues?q=is%3Aissue+is%3Aopen+label%3Asig%2Fauth+-label%3Apriority%2Fawaiting-more-evidence+-label%3Apriority%2Fimportant-longterm+-label%3Apriority%2Fimportant-soon+-label%3Apriority%2Fcritical-urgent+-label%3Apriority%2Fbacklog+-no%3Aassignee+-label%3Atriage%2Fneeds-information)
    *   [open bugs](https://github.com/kubernetes/kubernetes/issues?utf8=✓&q=is%3Aissue+is%3Aopen+label%3Asig%2Fauth+label%3Akind%2Fbug)


## July 10, 2019, 11a - Noon (Pacific Time)



*   [Recording](https://youtu.be/WP0-uR-BQC0)
*   Announcements
    *   
*   Demos
    *   
*   Pulls of note
    *   
*   Issues of note
    *   
*   Designs of note
    *   
*   Discussion topic
    *   Future of PodSecurityPolicy, and whether (or to what extent) policy belongs in core Kubernetes - [see this comment](https://github.com/kubernetes/enhancements/issues/5#issuecomment-503657710) for context
        *   [Slides](https://docs.google.com/presentation/d/1Kv6BSBNyLCyglMbK7e6tVOaDYe89LV2aHL2Hlb-9HX8/edit?usp=sharing)
*   Action items
    *   
*   Sweep issues with leftover time
    *   [unprioritized issues (-needs-information)](https://github.com/kubernetes/kubernetes/issues?utf8=✓&q=is%3Aissue+is%3Aopen+label%3Asig%2Fauth+-label%3Apriority%2Fawaiting-more-evidence+-label%3Apriority%2Fimportant-longterm+-label%3Apriority%2Fimportant-soon+-label%3Apriority%2Fcritical-urgent+-label%3Apriority%2Fbacklog+-label%3Atriage%2Fneeds-information+)
    *   [unassigned issues (-needs-information)](https://github.com/kubernetes/kubernetes/issues?utf8=✓&q=is%3Aissue+is%3Aopen+label%3Asig%2Fauth+no%3Aassignee+-label%3Atriage%2Fneeds-information+)
    *   [open bugs](https://github.com/kubernetes/kubernetes/issues?utf8=✓&q=is%3Aissue+is%3Aopen+label%3Asig%2Fauth+label%3Akind%2Fbug)


## June 26, 2019, 11a - Noon (Pacific Time)



*   [Recording](https://youtu.be/vxGhFtF7q10)
*   Announcements
    *   Requesting feedback on what folks would like to see covered in KubeCon NA 2019 (San Diego) SIG Auth deep dive (please respond to thread on mailing list)
*   Demos
    *   Demo and feedback for [kube-oidc-proxy](http://github.com/jetstack/kube-oidc-proxy) - Joshua Van Leeuwen (@joshvanl)
*   Pulls of note
    *   
*   Issues of note
    *   
*   Designs of note
    *   
*   Discussion topic
    *   [@shturec] Quick heads up: auditregistration/v1alpha1 api review 

        https://github.com/kubernetes/kubernetes/pull/71230 

*   Action items
    *   As part of the key ID changes to SA, we should consider if all SA tokens should no longer be considered opaque
*   Can sweep unassigned issues with leftover time
    *   [unprioritized+unassigned issues](https://github.com/kubernetes/kubernetes/issues?q=is%3Aissue+is%3Aopen+label%3Asig%2Fauth+-label%3Apriority%2Fawaiting-more-evidence+-label%3Apriority%2Fimportant-longterm+-label%3Apriority%2Fimportant-soon+-label%3Apriority%2Fcritical-urgent+-label%3Apriority%2Fbacklog+no%3Aassignee)
    *   [unprioritized+assigned issues](https://github.com/kubernetes/kubernetes/issues?utf8=%E2%9C%93&q=is%3Aissue+is%3Aopen+label%3Asig%2Fauth+-label%3Apriority%2Fawaiting-more-evidence+-label%3Apriority%2Fimportant-longterm+-label%3Apriority%2Fimportant-soon+-label%3Apriority%2Fcritical-urgent+-label%3Apriority%2Fbacklog+-no%3Aassignee+)


## June 12, 2019, 11a - Noon (Pacific Time)



*   [Recording](https://www.youtube.com/watch?v=tIFooOtL_40)
*   Announcements
    *   
*   Pulls of note
    *   
*   Issues of note
    *   
*   Designs of note
    *   [External Key Signer KEP](https://github.com/kubernetes/enhancements/pull/704)
    *   [Retroactive KEP: Certificates API](https://github.com/kubernetes/enhancements/pull/1097)
*   Discussion topic
    *   [@jackkleeman] certificate rotation for more cluster components
        *   Is a kubelet style approach appropriate where perhaps you provide an initial cert and then the application keeps it fresh using CSR API
        *   Could the controller sign any csr requested by an entity with the exact same username and group, if they have a special role.
        *   Instead, could we perhaps just allow the reload of certs from disk on a signal
            *   this is the preferred first step, more reusable, allows integration with a broader variety of PKIs
    *   [@ahmedtd, @mikedanese] Add Key IDs to access tokens https://github.com/kubernetes/kubernetes/pull/78502
        *   Follow up with @mo
        *   Note it in OpenID Connect discovery doc
*   Can sweep unassigned issues with leftover time
    *   [unprioritized+unassigned issues](https://github.com/kubernetes/kubernetes/issues?q=is%3Aissue+is%3Aopen+label%3Asig%2Fauth+-label%3Apriority%2Fawaiting-more-evidence+-label%3Apriority%2Fimportant-longterm+-label%3Apriority%2Fimportant-soon+-label%3Apriority%2Fcritical-urgent+-label%3Apriority%2Fbacklog+no%3Aassignee)
    *   [unprioritized+assigned issues](https://github.com/kubernetes/kubernetes/issues?utf8=%E2%9C%93&q=is%3Aissue+is%3Aopen+label%3Asig%2Fauth+-label%3Apriority%2Fawaiting-more-evidence+-label%3Apriority%2Fimportant-longterm+-label%3Apriority%2Fimportant-soon+-label%3Apriority%2Fcritical-urgent+-label%3Apriority%2Fbacklog+-no%3Aassignee+)


## May 29, 2019, 11a - Noon (Pacific Time)



*   Cancelled, empty agenda


## May 15, 2019, 11a - Noon (Pacific Time)



*   Cancelled, empty agenda


## May 1, 2019, 11a - Noon (Pacific Time)



*   [Recording](https://youtu.be/CsKaTVzP1sA)
*   Announcements
    *   1.15 feature freeze was yesterday
*   Pulls of note
    *   
*   Issues of note:
*   Designs of note
*   Discussion topic
    *   [cjcullen] Product Security Committee handling of ~Medium severity vulns: https://github.com/kubernetes/security/pull/28
        *   Open thread on [kubernetes-security-discuss](https://groups.google.com/forum/#!forum/kubernetes-security-discuss) (cc sig-auth)
        *   Might be worth adding a sentence on the availability metric of CVSS as well
        *   Ask security audit team if a threat model is being produced
        *   https://github.com/kubernetes/community/pull/504
    *   [haiyanmeng] [Node-scoped DaemonSets](https://github.com/kubernetes/enhancements/pull/944) follow-up
        *   short-term plan, run node daemonsets with kubelet credentials
        *   long-term plan
            *   avoid tangling authorizers together
            *   consider conditionalizing authorization
    *   [@pbarker] Discuss what issues need to be completed for dynamic audit to reach beta https://github.com/kubernetes/kubernetes/issues/70816
*   Can sweep unassigned issues with leftover time
    *   [unprioritized+unassigned issues](https://github.com/kubernetes/kubernetes/issues?q=is%3Aissue+is%3Aopen+label%3Asig%2Fauth+-label%3Apriority%2Fawaiting-more-evidence+-label%3Apriority%2Fimportant-longterm+-label%3Apriority%2Fimportant-soon+-label%3Apriority%2Fcritical-urgent+-label%3Apriority%2Fbacklog+no%3Aassignee)
    *   [unprioritized+assigned issues](https://github.com/kubernetes/kubernetes/issues?utf8=%E2%9C%93&q=is%3Aissue+is%3Aopen+label%3Asig%2Fauth+-label%3Apriority%2Fawaiting-more-evidence+-label%3Apriority%2Fimportant-longterm+-label%3Apriority%2Fimportant-soon+-label%3Apriority%2Fcritical-urgent+-label%3Apriority%2Fbacklog+-no%3Aassignee+)


## April 17, 2019, 11a - Noon (Pacific Time)



*   [Recording](https://www.youtube.com/watch?v=BvVhLut5KSQ)
*   Announcements
    *   
*   Pulls of note
    *   
*   Issues of note:
*   Designs of note
*   Discussion topic
    *   Gatekeeper presentation & status update
        *   [slides](https://docs.google.com/presentation/d/1zjpoS9wmmMRWGzCdLtMxVy8QY_Wqgt4ushmuHa8vItc/edit?usp=sharing)
        *   pain points:
            *   dynamic restmapper, watch management
    *   [haiyanmeng] [KEP](https://github.com/kubernetes/enhancements/pull/944) for Node-scoped DaemonSet 
*   
    *   [liggitt] split RBAC reconcile/evaluation to staging repo
        *   for consumption by `[kubectl auth reconcile](https://github.com/kubernetes/kubernetes/pull/74879#issuecomment-478675341)`, other external consumers
        *   [mailing list thread](https://groups.google.com/forum/#!topic/kubernetes-sig-auth/3OPUfE4_vXk)
*   Can sweep unassigned issues with leftover time
    *   [unprioritized+unassigned issues](https://github.com/kubernetes/kubernetes/issues?q=is%3Aissue+is%3Aopen+label%3Asig%2Fauth+-label%3Apriority%2Fawaiting-more-evidence+-label%3Apriority%2Fimportant-longterm+-label%3Apriority%2Fimportant-soon+-label%3Apriority%2Fcritical-urgent+-label%3Apriority%2Fbacklog+no%3Aassignee)
    *   [unprioritized+assigned issues](https://github.com/kubernetes/kubernetes/issues?utf8=%E2%9C%93&q=is%3Aissue+is%3Aopen+label%3Asig%2Fauth+-label%3Apriority%2Fawaiting-more-evidence+-label%3Apriority%2Fimportant-longterm+-label%3Apriority%2Fimportant-soon+-label%3Apriority%2Fcritical-urgent+-label%3Apriority%2Fbacklog+-no%3Aassignee+)


## April 3, 2019, 11a - Noon (Pacific Time)



*   [Recording](https://youtu.be/OQ-pld8P4FI)
*   Announcements
    *   
*   Pulls of note
    *   
*   Issues of note:
*   Designs of note
*   Discussion topic
    *   [destijl] [Discouraging use of secrets in environment variables](https://docs.google.com/document/d/1oU_9QPtYuDjVJLJRD2JsmC0xiSqTdcZbMU0kShLOtnE/edit#)
        *   for kube
            *   update API docs to indicate security issues
            *   remove any use from actual kube artifacts
            *   when included in examples, accompany with caveats
            *   Beyond docs:
                *   We likely want some way to programmatically discourage and enforce policy for secrets in env vars. It probably needs to be applicable per namespace, rather than per cluster.
                *   Current feeling is that solving this outside K8s using something like OPA would make some sense.
        *   for knative
            *   consider ways to enable but require users to be aware of security issues (naming of API to include "insecure", etc?)
    *   [tallclair] RuntimeClass supported features & policy
        *   "what does the runtime support?" vs "what is the user allowed to do?"
        *   E.g. gvisor doesn’t allow host networking on purpose, they don’t want people thinking they are sandboxed but allowed to poke a dangerous hole. While gvisor *could* support host networking, they don’t want to.
        *   Windows pods: want to fail fast if they request linux features.
        *   If you request apparmor in selinux pod will fail, reverse isn’t true. Some inconsistency there.
        *   Runtimeclass admission control validation that handles these separately? Do plan to add that for: [injecting pod overhead (WIP)](https://github.com/kubernetes/enhancements/issues/688), [scheduling (tolerations, KEP exists).](https://github.com/kubernetes/enhancements/issues/894)
        *   Do I want to write pods that target multiple runtimeclasses? Default runtimeclass so I don’t care as a pod author? Plan to handle defaulting through PSP. Maybe runtimeclass selection: “support for windows”, “support for GPUs” etc.
        *   Should pod authors target by name the runtimeclass? Expect to do nothing? Will windows add a runtimeclass?
        *   [clayton] We may want a completely separate podspec for windows.
    *   [liggitt] starting KEP for Certificate Signing Request (CSR) to v1
        *   will be looking at partitioning identified gaps/wishlist to determine what is a blocker for v1
        *   https://github.com/kubernetes/kubernetes/issues/69836
        *   https://github.com/kubernetes/kubernetes/issues/64547
        *   https://github.com/kubernetes/kubernetes/issues/67436
*   Can sweep unassigned issues with leftover time
    *   [unprioritized+unassigned issues](https://github.com/kubernetes/kubernetes/issues?q=is%3Aissue+is%3Aopen+label%3Asig%2Fauth+-label%3Apriority%2Fawaiting-more-evidence+-label%3Apriority%2Fimportant-longterm+-label%3Apriority%2Fimportant-soon+-label%3Apriority%2Fcritical-urgent+-label%3Apriority%2Fbacklog+no%3Aassignee)
    *   [unprioritized+assigned issues](https://github.com/kubernetes/kubernetes/issues?utf8=%E2%9C%93&q=is%3Aissue+is%3Aopen+label%3Asig%2Fauth+-label%3Apriority%2Fawaiting-more-evidence+-label%3Apriority%2Fimportant-longterm+-label%3Apriority%2Fimportant-soon+-label%3Apriority%2Fcritical-urgent+-label%3Apriority%2Fbacklog+-no%3Aassignee+)


## March 20, 2019, 11a - Noon (Pacific Time)



*   [Recording](https://youtu.be/11azGj-GtNg)
*   Announcements
    *   [1.15 planning](https://docs.google.com/document/d/1oZ24owtP5heL_3VxIHnZz82vSk_M97YHBqwRNe7uhxM/edit?usp=sharing)
*   Pulls of note
    *   [Node self labeling deprecation timeline](https://github.com/kubernetes/enhancements/pull/852)
*   Issues of note:
    *   [first bug filed due to incompatibility of ProjectedTokenVolume](https://github.com/kubernetes/kubernetes/issues/74565)
        *   Alpha is working as intended
    *   Audit e2e marked as flaky
        *   Plan to migrate many of these test cases to integration
        *   Add retries, permit and handle dropped audit events
        *   In 1.15, we can also get more consistent results by migrating to the dynamic audit webhook
    *   [encrypted secrets](https://github.com/kubernetes/kubernetes/issues/73838)
        *   KEP and feature tracking for 1.15 coming shortly.
*   Designs of note
*   Discussion topic
    *   [service account external signer](https://github.com/kubernetes/enhancements/pull/704)
        *   Focus on requirements. Let’s make sure that the requirements are settled then evaluate possible solutions.
        *   Do any existing interfaces exist that already fit this use case?
            *   PKCS#11? Loading an ABI much bigger than needed interface.


## March 6, 2019, 11a - Noon (Pacific Time)



*   [Recording](https://www.youtube.com/watch?v=FLAuwUWVF74)
*   Announcements
    *   extensions/v1beta1 removal plans (1.16) - [#43214](https://github.com/kubernetes/kubernetes/issues/43214), @liggitt
        *   AI: include details about PSP permissions check
    *   Code freeze tomorrow
    *   wg-multitenancy repo
        *   https://groups.google.com/forum/#!topic/kubernetes-sig-auth/IbPS3qWKxVI
*   Demos
    *   Get feedback for [rbac-manager](https://github.com/reactiveops/rbac-manager) and [rbac-lookup](https://github.com/reactiveops/rbac-lookup) - Rob Scott (@robertjscott)
*   Test flakes
    *   [Audit e2e tests](https://github.com/kubernetes/kubernetes/issues/74745)
        *   Somewhat fragile methodology (scrape audit logs from e2e master, gce-specific), affected by log rotation policies, etc
        *   Should possibly be integration tests instead of e2es?
*   Pulls of note
    *   Audit metadata fix for custom resources - [#74617](https://github.com/kubernetes/kubernetes/pull/74617) 
    *   Controller manager using rotatable tokens - [#72179](https://github.com/kubernetes/kubernetes/pull/72179)
    *   Service account issuer discovery - [community#2314](https://github.com/kubernetes/community/pull/2314/)
        *   AI: KEP from comments
*   Designs of note
*   Discussion topic
    *   Use of meetings for grooming/planning/communication
        *   [mailing list thread](https://groups.google.com/forum/#!topic/kubernetes-sig-auth/5ivS2Dr93h8) 
        *   [cluster-lifecycle example](https://github.com/kubernetes/community/blob/master/sig-cluster-lifecycle/grooming.md)
        *   Sweep things in progress
        *   Clearly define exit criteria
        *   Map out timelines for completion, blockers
    *   Example queries:
        *   Recent items:
            *   [Issues opened since 2019-01-01](https://github.com/kubernetes/kubernetes/issues?utf8=%E2%9C%93&q=is%3Aopen+is%3Aissue+label%3Asig%2Fauth+created%3A%3E2019-01-01+)
            *   [Pull requests opened since 2019-01-01](https://github.com/kubernetes/kubernetes/pulls?utf8=%E2%9C%93&q=is%3Apr+is%3Aopen+label%3Asig%2Fauth+created%3A%3E2019-01-01)
        *   [Issues with no assigned priority](https://github.com/kubernetes/kubernetes/issues?utf8=%E2%9C%93&q=is%3Aissue+is%3Aopen+label%3Asig%2Fauth+-label%3Apriority%2Fawaiting-more-evidence+-label%3Apriority%2Fimportant-longterm+-label%3Apriority%2Fimportant-soon+-label%3Apriority%2Fcritical-urgent+) (mostly feature requests, needs triage)
        *   [Issues with kind/bug](https://github.com/kubernetes/kubernetes/issues?utf8=%E2%9C%93&q=is%3Aissue+is%3Aopen+label%3Asig%2Fauth+label%3Akind%2Fbug+) 


## February 20, 2019, 11a - Noon (Pacific Time)



*   Cancelled, empty agenda


## February 6, 2019, 11a - Noon (Pacific Time)



*   Cancelled, empty agenda


## January 23, 2019, 11a - Noon (Pacific Time)



*   [Recording](https://youtu.be/oiaf_xDdb80)
*   Announcements
    *   Seeding tech leads PR open for comment ([mailing list thread](https://groups.google.com/d/msg/kubernetes-sig-auth/v0-Awf0rBOg/Dk3SipCTAQAJ), [PR](https://github.com/kubernetes/community/pull/3080))
*   Demos
    *   [K8s secret flex volume](https://github.com/Azure/kubernetes-keyvault-flexvol) and [K8s secrets store CSI driver](https://github.com/deislabs/secrets-store-csi-driver) - Rita Zhang (@ritazh)
        *   Feedback on PV mechanism vs possible inline CSI volume ([#64984](https://github.com/kubernetes/kubernetes/issues/64984))
        *   Support env variables: https://github.com/Azure/kubernetes-keyvault-flexvol/issues/28
        *   Update secret value after it’s mounted: https://github.com/Azure/kubernetes-keyvault-flexvol/issues/62
        *   
*   Pulls of note
    *   Deprecate DenyEscalatingExec admission plugin - [#72737](https://github.com/kubernetes/kubernetes/pull/72737)
    *   Deprecating --allow-privileged (@tallclair) - https://github.com/kubernetes/kubernetes/pull/71835
*   Designs of note
    *   
*   Discussion topics
    *   Discuss/review [GMSA KEP for Windows](https://github.com/kubernetes/enhancements/pull/666) - Deep Debroy (@ddebroy)/Jeremy Wood (@jeremywx)/Jean Rogue (@jean)
    *   [@pbarker] APIserver authentication to webhooks KEP https://github.com/kubernetes/enhancements/pull/658
        *   Add comparison/discussion of existing kubeconfig-based mechanism for admission webhooks
        *   Add discussion of audience determination
    *   Discuss KEP requirements and timeline for the 1.14 release (@marpaia)
        *   sig-release: all enhancements for 1.14 must have a reviewed/approved/merged (and implementable?) KEP by feature freeze on 1/29
        *   Old enhancements already in progress need a KEP that includes graduation criteria, testing plan (see the KEP template for relevant checklists)
        *   Assistance available from sig-release
    *   1.14 plans/priorities (add your name to items you plan to work on/participate in)
        *   Work through blockers to CSR API promotion - [#69836](https://github.com/kubernetes/kubernetes/issues/69836) (@liggitt, @krmayankk)
            *   Divide items in that issue into required-for-v1 vs possible-for-v2
            *   Work on KEP with roadmap during 1.14, plan to start API updates in 1.15
        *   Kubelet [Client](https://github.com/kubernetes/enhancements/issues/266)/[Serving](https://github.com/kubernetes/enhancements/issues/267) cert rotation graduation (@liggitt, @krmayankk)
            *   Pull kubelet client cert rotation into KEP format for 1.14, ensure testing/docs are sufficient, push to graduate for 1.14
            *   Pull kubelet server cert rotation into KEP format, promote existing alpha CI tests to always run, consider graduating for 1.14 if ready in time
        *   Webhook auth - [doc](https://docs.google.com/document/d/1rtnPnrW2Ws1I8h826oYwSKuqWif5SaRnm2jvwFH9lIk/edit#heading=h.jlrb1d2p584h), [kep](https://github.com/kubernetes/enhancements/pull/658) (@pbarker?, @liggitt, @krmayankk)
        *   CVE-2018-1002105 [post-mortem action items](https://github.com/kubernetes/kubernetes/issues?utf8=%E2%9C%93&q=is%3Aopen+is%3Aissue+CVE-2018-1002105+label%3Asig%2Fauth) (@dekkagaijin)
        *   Transition SA controller clients to TokenRequest API [#71275](https://github.com/kubernetes/kubernetes/issues/71275) (@enj)
        *   [RunAsGroup](https://github.com/kubernetes/enhancements/issues/213) to beta (Currently not clear what blockers we have to make this change) . I opened the following test issues for CI ([#72287](https://github.com/kubernetes/kubernetes/issues/72287), [#72253](https://github.com/kubernetes/kubernetes/issues/72253)) (@krmayankk)
            *   No known blockers, work through the checklist for graduation/test
        *   [External JWT signing](https://github.com/kubernetes/enhancements/pull/704) support ([#73110](https://github.com/kubernetes/kubernetes/pull/73110)) (@micahhausler)
            *   Motivations:
                *   update signing keys without restart
                *   avoid secret material on disk
                *   allow offloading of token signing
        *   Support configurable ProjectedTokenVolume rotation period ([#73221](https://github.com/kubernetes/kubernetes/issues/73221)) (@micahhausler)


## ~~January 9, 2019, 11a - Noon (Pacific Time)~~ 

cancelled for zoom outage :( 
