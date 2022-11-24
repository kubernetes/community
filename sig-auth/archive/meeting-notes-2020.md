# Kubernetes Sig-auth Meeting Agenda

## December 23 - CANCELLED

Cancelled for December holidays.


## December 9, 11a - Noon (Pacific Time)



*   [Recording](https://youtu.be/beu92Bjble8)
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
    *   [ankeesler] Next steps for getting client-go credential plugins to GA
        *   [https://github.com/kubernetes/enhancements/tree/master/keps/sig-auth/541-external-credential-providers#beta---ga-graduation](https://github.com/kubernetes/enhancements/tree/master/keps/sig-auth/541-external-credential-providers#beta---ga-graduation)
        *   track/resolve known issues, e.g. [https://github.com/kubernetes/kubernetes/issues/91913](https://github.com/kubernetes/kubernetes/issues/91913)
        *   ensure we're on sig-release's tracking list before feature freeze (timeline will be posted in [https://github.com/kubernetes/sig-release/tree/master/releases/](https://github.com/kubernetes/sig-release/tree/master/releases/) once available)
        *   ensure v1 promotion is on API review board (`/label api-review` PRs once open, open early enough to get review bandwidth)
    *   [ahmedtd] SPIFFE client certificate authentication to kube-apiserver: any existing designs / KEPs / work?
        *   SPIFFE puts cert UID in CN, puts username in SAN
        *   questions about what the mapping should be, questions about whether this is a thing that kube-apiserver should be aware of
        *   suggest mailing [https://groups.google.com/g/kubernetes-sig-auth](https://groups.google.com/g/kubernetes-sig-auth) in case there are people working in the area not in the meeting
    *   [zshihang] Disable TokenController/Provide a per service account knob to disable.
        *   can we have a per-service-account way to opt out of controller-created tokens for specific service accounts in lieu of bound service account tokens
        *   general positive attitude toward making it possible to opt out; could be via a field on serviceaccount (or annotation, if there's hesitance on introducing API surface around a controller behavior we'd like to see dropped in the far future)
        *   possible users:
            *   Could allow internal platform providers to roll out a opinionated default in a gradual way via admission control (defined in a site-specific way)
            *   kube-controller-manager opting out of secret-based tokens for controller manager service accounts
            *   wonderful security-conscious users
            *   Tim
    *   [micahhausler] Knob to prevent nodes from creating mirror pods
        *   for clusters that don't make use of mirror pods, would close a gap around nodes creating pods with labels that route traffic to them
        *   services routing to pods can be a feature or a vulnerability, depending on the cluster deployer's perspective
        *   static pod label restriction would narrowly+optionally fix the service hijacking issue ([https://github.com/kubernetes/enhancements/blob/master/keps/sig-auth/20190916-noderestriction-pods.md#label-restrictions](https://github.com/kubernetes/enhancements/blob/master/keps/sig-auth/20190916-noderestriction-pods.md#label-restrictions))
    *   [nckturner] client request signing ([#92535](https://github.com/kubernetes/kubernetes/issues/92535))
        *   consider looking into local authenticated proxy (kubectl -> local proxy -> apiserver)
        *   concerns around exec plugin: performance, inventing a request signing API/standard
    *   [tsable] Future of PodSecurityPolicy
        *   Work toward initial proposal from Tabitha Sable, Ian Coldwater, et al ongoing
        *   liggitt: would recommend starting public discussion as early as possible to make sure progress can be made early
        *   liggitt: since the existing PodSecurityPolicy API is not graduating as-is, suggest marking deprecation in 1.21 (ahead of planned 1.22) to communicate to users they should not start using it if they aren't already. Would keep the planned removal at 1.25.
        *   tsable: would like to reserve the right to move PSP removal out past 1.25 to allow seamless transition to new thing
        *   liggitt: 1.25 is a significant distance away... don't want to commit to moving out removal; instead, would use that endpoint to motivate progress on the replacement
        *
        *   tallclair: while working on the design for the replacement, would be helpful to make clearer what threat models are being addressed by the replacement (container breakout, node compromise, etc) to avoid unbounded API surface
*   Action Items
    *
*   Sweep issues with leftover time
    *   [CI flakes](https://storage.googleapis.com/k8s-gubernator/triage/index.html?sig=auth)
    *   [CI testgrids](https://testgrid.k8s.io/sig-auth)
    *   [unprioritized+unassigned (-needs-information) issues](https://github.com/kubernetes/kubernetes/issues?q=is%3Aissue+is%3Aopen+label%3Asig%2Fauth+-label%3Apriority%2Fawaiting-more-evidence+-label%3Apriority%2Fimportant-longterm+-label%3Apriority%2Fimportant-soon+-label%3Apriority%2Fcritical-urgent+-label%3Apriority%2Fbacklog+no%3Aassignee+-label%3Atriage%2Fneeds-information)
    *   [unprioritized+assigned (-needs-information) issues](https://github.com/kubernetes/kubernetes/issues?q=is%3Aissue+is%3Aopen+label%3Asig%2Fauth+-label%3Apriority%2Fawaiting-more-evidence+-label%3Apriority%2Fimportant-longterm+-label%3Apriority%2Fimportant-soon+-label%3Apriority%2Fcritical-urgent+-label%3Apriority%2Fbacklog+-no%3Aassignee+-label%3Atriage%2Fneeds-information)
    *   [open bugs](https://github.com/kubernetes/kubernetes/issues?utf8=✓&q=is%3Aissue+is%3Aopen+label%3Asig%2Fauth+label%3Akind%2Fbug)


## November 25 - CANCELLED

Cancelled for US Thanksgiving holiday.


## November 11, 11a - Noon (Pacific Time)



*   [Recording](https://www.youtube.com/watch?v=HFc-XX8a1Lo)
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
    *   Future of PodSecurityPolicy
        *   [Options doc](https://docs.google.com/document/d/1VKqjUlpU888OYtIrBwidL43FOLhbmOD5tesYwmjzO4E/edit#)
    *   [ankeesler] Next steps for getting client-go credential plugins to GA
    *   [ahmedtd] SPIFFE client certificate authentication to kube-apiserver: any existing designs / KEPs / work?
*   Action Items
    *
*   Sweep issues with leftover time
    *   [CI flakes](https://storage.googleapis.com/k8s-gubernator/triage/index.html?sig=auth)
    *   [CI testgrids](https://testgrid.k8s.io/sig-auth)
    *   [unprioritized+unassigned (-needs-information) issues](https://github.com/kubernetes/kubernetes/issues?q=is%3Aissue+is%3Aopen+label%3Asig%2Fauth+-label%3Apriority%2Fawaiting-more-evidence+-label%3Apriority%2Fimportant-longterm+-label%3Apriority%2Fimportant-soon+-label%3Apriority%2Fcritical-urgent+-label%3Apriority%2Fbacklog+no%3Aassignee+-label%3Atriage%2Fneeds-information)
    *   [unprioritized+assigned (-needs-information) issues](https://github.com/kubernetes/kubernetes/issues?q=is%3Aissue+is%3Aopen+label%3Asig%2Fauth+-label%3Apriority%2Fawaiting-more-evidence+-label%3Apriority%2Fimportant-longterm+-label%3Apriority%2Fimportant-soon+-label%3Apriority%2Fcritical-urgent+-label%3Apriority%2Fbacklog+-no%3Aassignee+-label%3Atriage%2Fneeds-information)
    *   [open bugs](https://github.com/kubernetes/kubernetes/issues?utf8=✓&q=is%3Aissue+is%3Aopen+label%3Asig%2Fauth+label%3Akind%2Fbug)


## October 28, 11a - Noon (Pacific Time)



*   Recording
*   Announcements
    *   Code freeze for 1.20 is Nov 12th
*   Demos
    *
*   Pulls of note
    *   1.20-targeted PRs
        *   Kubelet credential provider externalization - [#94196](https://github.com/kubernetes/kubernetes/pull/94196)
        *   client-go exec auth plugin - [#95489](https://github.com/kubernetes/kubernetes/pull/95489)
        *   csi service account token - [#93130](https://github.com/kubernetes/kubernetes/pull/93130)
*   Issues of note
    *
*   Designs of note
    *
*   Discussion topic
    *   Close 1.20 loose ends:
        *   Token* Beta/GA graduation
            *   [#95896](https://github.com/kubernetes/kubernetes/pull/95896): make the flags required for TokenRequest so that presubmits e2e tests can pass because the prow job is using kubeadm.
                *   This may be deferred to post-GA.
                *   AI: Reach out to owners of deployment automation + clouds. Ask if this would be painful.
            *   [#95886](https://github.com/kubernetes/kubernetes/pull/95886): add a general test and later will change it to conformance one in the TokenRequest GA pr.
                *   We want this even if the flags aren’t made required in 1.20.
            *   [#93258](https://github.com/kubernetes/kubernetes/pull/93258): mv TokenRequest to GA. this will be reworked after the first two merged.
            *   [#95667](https://github.com/kubernetes/kubernetes/pull/95667): mv BoundServiceAccountToken beta. this will also be reworked after the first three merged.
    *   FYI - label/metadata policy discussion happening in sig-arch
        *   mailing list thread
        *   summary doc
        *   planned for discussion at next sig-arch meeting (11/5)
*   Action Items
    *
*   Sweep issues with leftover time
    *   [CI flakes](https://storage.googleapis.com/k8s-gubernator/triage/index.html?sig=auth)
    *   [CI testgrids](https://testgrid.k8s.io/sig-auth)
    *   [unprioritized+unassigned (-needs-information) issues](https://github.com/kubernetes/kubernetes/issues?q=is%3Aissue+is%3Aopen+label%3Asig%2Fauth+-label%3Apriority%2Fawaiting-more-evidence+-label%3Apriority%2Fimportant-longterm+-label%3Apriority%2Fimportant-soon+-label%3Apriority%2Fcritical-urgent+-label%3Apriority%2Fbacklog+no%3Aassignee+-label%3Atriage%2Fneeds-information)
    *   [unprioritized+assigned (-needs-information) issues](https://github.com/kubernetes/kubernetes/issues?q=is%3Aissue+is%3Aopen+label%3Asig%2Fauth+-label%3Apriority%2Fawaiting-more-evidence+-label%3Apriority%2Fimportant-longterm+-label%3Apriority%2Fimportant-soon+-label%3Apriority%2Fcritical-urgent+-label%3Apriority%2Fbacklog+-no%3Aassignee+-label%3Atriage%2Fneeds-information)
    *   [open bugs](https://github.com/kubernetes/kubernetes/issues?utf8=✓&q=is%3Aissue+is%3Aopen+label%3Asig%2Fauth+label%3Akind%2Fbug)


## October 14, 11a - Noon (Pacific Time)



*   Meeting cancelled - empty agenda.


## September 30, 11a - Noon (Pacific Time)



*   Meeting cancelled - empty agenda.


## September 16, 11a - Noon (Pacific Time)



*   Recording
*   Announcements
    *
*   Demos
    *
*   Pulls of note
    *   [@andrewsykim] initial implementation of docker credentials plugins [https://github.com/kubernetes/kubernetes/pull/94196](https://github.com/kubernetes/kubernetes/pull/94196)
        *   based on KEP approved by SIG Auth [https://github.com/kubernetes/enhancements/pull/1406](https://github.com/kubernetes/enhancements/pull/1406)
            *   mikedanese: Send an update to the KEP so that it reflects the current proposal
            *   liggitt: someone from sig-node and sig-auth should review the PR.
                *   Some concerns around caching.
                *   Liggitt offered to review from a sig-auth perspective.
*   Issues of note
    *   MTB: [kubernetes-sigs/multi-tenancy/issues/1104](https://github.com/kubernetes-sigs/multi-tenancy/issues/1104) (Jim)
        *   Can we use MTB for other types of benchmarks e.g. Pod Security Profiles.
        *   There’s an effort to refactor MTB to be independent of the benchmark.
        *   [https://github.com/kubernetes-sigs/multi-tenancy/tree/master/benchmarks](https://github.com/kubernetes-sigs/multi-tenancy/tree/master/benchmarks)
*   Designs of note
    *   Windows privileged containers proposal [@marosset, @immuzz, @amberguo]
        *   Windows privileged container support ([KEP](https://docs.google.com/document/d/12EUtMdWFxhTCfFrqhlBGWV70MkZZPOgxw0X-LTR0VAo/edit?usp=sharing))
        *   [enhancement issue](https://github.com/kubernetes/enhancements/issues/1981) for tracking
        *   Looking for feedback from SIG-auth on proposal
        *   AI: make sure that changes you make are sync’d to Pod Security Profiles.
            *   E.g. [github.com/open-policy-agent/gatekeeper-library/library/pod-security-policy](https://github.com/open-policy-agent/gatekeeper-library/tree/master/library/pod-security-policy)
            *
*   Discussion topic
    *
*   Action Items
    *
*   Sweep issues with leftover time
    *   [CI flakes](https://storage.googleapis.com/k8s-gubernator/triage/index.html?sig=auth)
    *   [CI testgrids](https://testgrid.k8s.io/sig-auth)
    *   [unprioritized+unassigned (-needs-information) issues](https://github.com/kubernetes/kubernetes/issues?q=is%3Aissue+is%3Aopen+label%3Asig%2Fauth+-label%3Apriority%2Fawaiting-more-evidence+-label%3Apriority%2Fimportant-longterm+-label%3Apriority%2Fimportant-soon+-label%3Apriority%2Fcritical-urgent+-label%3Apriority%2Fbacklog+no%3Aassignee+-label%3Atriage%2Fneeds-information)
    *   [unprioritized+assigned (-needs-information) issues](https://github.com/kubernetes/kubernetes/issues?q=is%3Aissue+is%3Aopen+label%3Asig%2Fauth+-label%3Apriority%2Fawaiting-more-evidence+-label%3Apriority%2Fimportant-longterm+-label%3Apriority%2Fimportant-soon+-label%3Apriority%2Fcritical-urgent+-label%3Apriority%2Fbacklog+-no%3Aassignee+-label%3Atriage%2Fneeds-information)
    *   [open bugs](https://github.com/kubernetes/kubernetes/issues?utf8=✓&q=is%3Aissue+is%3Aopen+label%3Asig%2Fauth+label%3Akind%2Fbug)


## September 2, 11a - Noon (Pacific Time)



*   Recording
*   Announcements
    *   sig-security formation: https://groups.google.com/forum/#!forum/kubernetes-sig-security
*   Demos
    *   [Multi-Tenancy Benchmarks (MTB)](https://github.com/kubernetes-sigs/multi-tenancy/tree/master/benchmarks/kubectl-mtb): Jim
*
*   Pulls of note
    *
*   Issues of note
    *
*   Designs of note
    *
*   Discussion topic
    *   CI health [@liggitt]
        *   [flakes triage](https://storage.googleapis.com/k8s-gubernator/triage/index.html?sig=auth) (having issues at the moment, see [#19078](https://github.com/kubernetes/test-infra/issues/19078))
        *   [recorded training for deflaking tests](https://youtu.be/Ewp8LNY_qTg)
        *   [tightening passing requirements for unit tests](https://github.com/kubernetes/kubernetes/pull/93605)
        *   Upgrade/skew tests for critical features (example: [switching to bound service account token admission](https://github.com/kubernetes/enhancements/pull/1912#discussion_r474122280))
            *   Jordan’s recommendation is to implement this upgrade test in the old suite (and help get the old suite running again... bringup was broken in 1.18 by [#868](https://github.com/kubernetes/release/pull/868#issuecomment-673727541)).
        *   [reliability working group proposal](https://docs.google.com/document/d/1IZnIG26FOJXWvW5n0Yu6zL89pH4cO8NYHKMiESc2fV0/edit)
*   Action Items
    *
*   Sweep issues with leftover time
    *   [CI flakes](https://storage.googleapis.com/k8s-gubernator/triage/index.html?sig=auth)
    *   [CI testgrids](https://testgrid.k8s.io/sig-auth)
    *   [unprioritized+unassigned (-needs-information) issues](https://github.com/kubernetes/kubernetes/issues?q=is%3Aissue+is%3Aopen+label%3Asig%2Fauth+-label%3Apriority%2Fawaiting-more-evidence+-label%3Apriority%2Fimportant-longterm+-label%3Apriority%2Fimportant-soon+-label%3Apriority%2Fcritical-urgent+-label%3Apriority%2Fbacklog+no%3Aassignee+-label%3Atriage%2Fneeds-information)
    *   [unprioritized+assigned (-needs-information) issues](https://github.com/kubernetes/kubernetes/issues?q=is%3Aissue+is%3Aopen+label%3Asig%2Fauth+-label%3Apriority%2Fawaiting-more-evidence+-label%3Apriority%2Fimportant-longterm+-label%3Apriority%2Fimportant-soon+-label%3Apriority%2Fcritical-urgent+-label%3Apriority%2Fbacklog+-no%3Aassignee+-label%3Atriage%2Fneeds-information)
    *   [open bugs](https://github.com/kubernetes/kubernetes/issues?utf8=✓&q=is%3Aissue+is%3Aopen+label%3Asig%2Fauth+label%3Akind%2Fbug)


## August 19, 11a - Noon (Pacific Time)



*   Meeting cancelled - empty agenda, people busy with kubecon and the last week of 1.19


## August 5, 11a - Noon (Pacific Time)



*   Recording
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
    *   [tallclair] Core label policy
        *   Labels don’t necessarily grant or restrict permission.
        *   UID policy in openshift
        *   Delegation
            *   Tim creates replicaset
            *   Can Tim use restricted labels?
            *   Pod OwnerRef/ControllerRef is set by replicasetcontroller
            *   RC has restricted labels, so pod is either allowed or required to have restricted labels.
            *   What about horizontal controllers? Something that swaps out an image on the replicaset?
            *   Tie the object back to the subject that kicked off the chain of events.
        *   This wouldn’t be enforced in authorization, probably admission.
            *   There are other use cases where cluster scoped objects need to be subdivided.
                *   CSR kind of needed this with signer name.
                *   Might be useful with CRDs.
        *   Policy attachment point:
            *   RBAC? Labels aren’t known.
            *   Gatekeeper?
            *   Network policy?
        *   Big win is to apply this to cluster scoped objects.
        *   Mike: If I were to re-review I KEP:
            *   I would like to understand how inheritance works. OwnerRef? What about inheritance from the namespace?
            *   What is the authorization model for the super user that can add or remove labels at specific points?
            *   What are the user stories that this enables? Gatekeeper? RBAC? Network Policy?
        *   Greg: We need an understandable authorization model.
        *   Deads: Namespace subdivision is hard to understand.
        *   Tim: Maybe the hierarchical namespace controller will address some of the use cases that make people want namespace subdivision.
    *   [zshihang] BoundServiceAccountTokenVolume Beta
        *   [https://github.com/kubernetes/enhancements/pull/1912](https://github.com/kubernetes/enhancements/pull/1912)
        *   What to do about warn after is still an open question?
            *   Do we pick a default silently or do we make warn after required?
        *   Warn after was added in 1.19.
        *   Deads: no choice is going to make it easy for people.
        *   Jordan: We can also enable warn after for X releases by default, then disable it for X releases by default, then remove? We start getting some better security properties, people have time to migrate.
        *   Mike: What do we do about token controller?
        *   Liggitt: Leave it running until BoundServiceAccountTokenVolume is GA. We should consider this as a separate deprecation.
            *   One option is to migrate token controller to use the TokenRequest API to populate secrets containing tokens (we thought about this at one point, which is why TokenRequest supports binding to a Secret object)
            *   We may have a lot of users who are creating service accounts, watching for secrets, extracting the token. We don’t have good visibility into this. We want them to ultimately switch to TokenRequest API. Probably announce auto-created service account secrets as deprecated once TokenRequest GAs and point people to that.
        *   Liggitt: TokenRequest API is probably ready for GA in 1.20.
            *   [https://github.com/kubernetes/kubernetes/pull/93258](https://github.com/kubernetes/kubernetes/pull/93258)
*   Action Items
    *
*   Sweep issues with leftover time
    *   [unprioritized+unassigned (-needs-information) issues](https://github.com/kubernetes/kubernetes/issues?q=is%3Aissue+is%3Aopen+label%3Asig%2Fauth+-label%3Apriority%2Fawaiting-more-evidence+-label%3Apriority%2Fimportant-longterm+-label%3Apriority%2Fimportant-soon+-label%3Apriority%2Fcritical-urgent+-label%3Apriority%2Fbacklog+no%3Aassignee+-label%3Atriage%2Fneeds-information)
    *   [unprioritized+assigned (-needs-information) issues](https://github.com/kubernetes/kubernetes/issues?q=is%3Aissue+is%3Aopen+label%3Asig%2Fauth+-label%3Apriority%2Fawaiting-more-evidence+-label%3Apriority%2Fimportant-longterm+-label%3Apriority%2Fimportant-soon+-label%3Apriority%2Fcritical-urgent+-label%3Apriority%2Fbacklog+-no%3Aassignee+-label%3Atriage%2Fneeds-information)
    *   [open bugs](https://github.com/kubernetes/kubernetes/issues?utf8=✓&q=is%3Aissue+is%3Aopen+label%3Asig%2Fauth+label%3Akind%2Fbug)


## July 22, 11a - Noon (Pacific Time)



*   [Recording](https://www.youtube.com/watch?v=-_J18AQ0opM)
*   Announcements
    *
*   Demos
    *
*   Pulls of note
    *
*   Issues of note
    *   [CSR API: cluster signing duration is not flexible and has unsafe defaulting #92678](https://github.com/kubernetes/kubernetes/issues/92678)
        *   Notes from discussion below
        *   A single new .spec field targeted at the user, two possible approaches:
            *   durationHint (when does the duration start?)
            *   notAfterHint (exact time the cert should expire)
        *   The CSR .spec field could be targeted via admission to enforce policy (though self hosted environments would need to be careful with Kubelet CSR flows by using static pods or hosting the webhook externally)
        *   No new field targeted at the approver or the signer
        *   Signer would choose to honor the field if it is able to (it is a hint, not a requirement or guarantee)
        *   Next steps:
            *   Create PR to update the existing CSR KEP with proposed changes
            *   Use that PR to get consensus on the specifics of the API changes
            *   Ping James Munnely for a review on the proposed changes since he missed this discussion and we want to ensure that cert-manager could use this API
            *   After agreement on the API changes:
                *   Create PR to update API (similar to signer name changes)
                *   Update the built-in signers to honor the new API
        *   Future:
            *   A kubelet that is aware of this field and its own rotation could use the field to set the duration to a shorter amount of time
*   Designs of note
    *
*   Discussion topic
    *   [jaybeale, tallclair] SIG-Security - [https://github.com/kubernetes/community/pull/4962](https://github.com/kubernetes/community/pull/4962) [Charter]
    *   [https://github.com/kubernetes/community/pull/4976](https://github.com/kubernetes/community/pull/4976) [Letter to Steering]
*   Action Items
    *
*   Sweep issues with leftover time
    *   [unprioritized+unassigned (-needs-information) issues](https://github.com/kubernetes/kubernetes/issues?q=is%3Aissue+is%3Aopen+label%3Asig%2Fauth+-label%3Apriority%2Fawaiting-more-evidence+-label%3Apriority%2Fimportant-longterm+-label%3Apriority%2Fimportant-soon+-label%3Apriority%2Fcritical-urgent+-label%3Apriority%2Fbacklog+no%3Aassignee+-label%3Atriage%2Fneeds-information)
    *   [unprioritized+assigned (-needs-information) issues](https://github.com/kubernetes/kubernetes/issues?q=is%3Aissue+is%3Aopen+label%3Asig%2Fauth+-label%3Apriority%2Fawaiting-more-evidence+-label%3Apriority%2Fimportant-longterm+-label%3Apriority%2Fimportant-soon+-label%3Apriority%2Fcritical-urgent+-label%3Apriority%2Fbacklog+-no%3Aassignee+-label%3Atriage%2Fneeds-information)
    *   [open bugs](https://github.com/kubernetes/kubernetes/issues?utf8=✓&q=is%3Aissue+is%3Aopen+label%3Asig%2Fauth+label%3Akind%2Fbug)


## July 8, 11a - Noon (Pacific Time)



*   Canceled due to empty agenda
*   Code freeze is July 9th


## June 24, 11a - Noon (Pacific Time)



*   [Recording](https://www.youtube.com/watch?v=Phco26zYEYs)
*   Announcements
    *   @jackkleeman -> Apple, expect delays on intermediate cert work while I figure the OSS process out
*   Demos
    *   Seccomp Operator - Paulo [[slides](https://docs.google.com/presentation/d/1tkz8Zgd4nzrTPaR1jDBkJgBu3jc6jy-4Vd9F1AyCyKE/edit#slide=id.g88d31c8bbc_0_63), [issue](https://github.com/kubernetes/org/issues/1873)]
    *   Policy Report CRD ([PR](https://github.com/kubernetes-sigs/wg-policy-prototypes/pull/4), [doc](https://docs.google.com/document/d/1nICYLkYS1RE3gJzuHOfHeAC25QIkFZfgymFjgOzMDVw/edit#heading=h.wn242twwl364)) - Jim, Erica
*   Pulls of note
    *   FYI: [https://github.com/kubernetes/kubernetes/pull/92006](https://github.com/kubernetes/kubernetes/pull/92006)
*   Issues of note
    *
*   Designs of note
    *
*   Discussion topic
    *   move “k8s.io/kubernetes/pkg/registry/rbac/reconciliation” to staging
        *   sig-cli: _"We are planning to move kubectl to staging([tracking issue](https://github.com/kubernetes/enhancements/issues/1020)), but now it is being blocked as “kubectl auth” still relies on “k8s.io/kubernetes/pkg/registry/rbac/reconciliation”"_
        *   Potential obstacle: the reconciliation package imports other k/k code, such as “pkg/registry/rbac/validation” and “pkg/apis/core/helper”, and these packages import others. We might need to figure out how to make the reconciliation package self-included.
        *   liggitt: initial effort at [https://github.com/liggitt/kubernetes/commits/extract-rbac](https://github.com/liggitt/kubernetes/commits/extract-rbac), got stuck on interface imports that currently live in k8s.io/apiserver (user.Info) or in kube-apiserver (RBAC validation)
            *   would welcome someone picking this up and working through the import issues
        *   deads2k: question about the destination repo/granularity (rbac-specific? auth-specific? "things that depend only on client+api+apimachinery" shared between
    *   [micahhausler] Discussion on adding request signing support (a la AWS SigV4)
        *   Tracking issue with thoughts: [#92535](https://github.com/kubernetes/kubernetes/issues/92535)
        *   On server-side, can do it w/ front-proxy
        *   Client side, can we be better than exposing the full http.RoundTripper interface to an external plugin?
        *   Does per-cluster HTTP proxy work (supported in kubeconfig now via proxy-url field under cluster config)? Is it easy enough? Is kubectl proxy any better?
        *   Would be nice to to have authentication options e.g. limited to “only get pods” or “only in this time”, etc.
    *   [mo] next steps for [External TLS certificate authenticator #1749](https://github.com/kubernetes/enhancements/pull/1749)?
        *   liggitt: similar questions to the previous topic... is externalizing TLS the right approach, or is routing through a client-side proxy sufficient to allow custom TLS behavior?
        *   AI:
*   Action Items
    *
*   Sweep issues with leftover time
    *   [unprioritized+unassigned (-needs-information) issues](https://github.com/kubernetes/kubernetes/issues?q=is%3Aissue+is%3Aopen+label%3Asig%2Fauth+-label%3Apriority%2Fawaiting-more-evidence+-label%3Apriority%2Fimportant-longterm+-label%3Apriority%2Fimportant-soon+-label%3Apriority%2Fcritical-urgent+-label%3Apriority%2Fbacklog+no%3Aassignee+-label%3Atriage%2Fneeds-information)
    *   [unprioritized+assigned (-needs-information) issues](https://github.com/kubernetes/kubernetes/issues?q=is%3Aissue+is%3Aopen+label%3Asig%2Fauth+-label%3Apriority%2Fawaiting-more-evidence+-label%3Apriority%2Fimportant-longterm+-label%3Apriority%2Fimportant-soon+-label%3Apriority%2Fcritical-urgent+-label%3Apriority%2Fbacklog+-no%3Aassignee+-label%3Atriage%2Fneeds-information)
    *   [open bugs](https://github.com/kubernetes/kubernetes/issues?utf8=✓&q=is%3Aissue+is%3Aopen+label%3Asig%2Fauth+label%3Akind%2Fbug)


## June 10, 11a - Noon (Pacific Time)



*   [Recording](https://youtu.be/nzWKyKXS48o)
*   Announcements
    *
*   Demos
    *
*   Pulls of note
    *   [CSR API PRs](https://github.com/orgs/kubernetes/projects/42?card_filter_query=is%3Apr+is%3Aopen+-label%3Ado-not-merge%2Fwork-in-progress)
    *   [Exec PRs](https://github.com/kubernetes/kubernetes/issues/90817)
        *   Cluster info [#91192](https://github.com/kubernetes/kubernetes/pull/91192)
        *   Install hint [#91305](https://github.com/kubernetes/kubernetes/pull/91305)
        *   Bearer token interaction [#91745](https://github.com/kubernetes/kubernetes/pull/91745)
    *
*   Issues of note
    *
*   Designs of note
    *
*   Discussion topic
    *   [mo] CSR API questions
        *   Custom MaxTTLHint?
            *   i.e. for the built-in signers, can a user request a cert that is shorter lived than the 1 year default (which is problematic since we lack any form of revocation)
            *   We state “Expiration/cert lifetime - minimum of CSR signer or request. Sanity of the time is the concern of the signer.”
            *   But we do not expose any per signer config, just a single “--cluster-signing-duration”
            *   Safe duration of client certs vs serving certs - long lived serving certs are probably okay but irrevocable client certs that assert an identity against the Kube API for a year by default seems like a bad idea
        *   Support flow via public key and CSR template without requiring a signed CSR?
            *   i.e. I want to issue a cert for a specific client without need the client’s private key
            *   Currently if you do not want the client to directly be able to create CSRs, you have to give the CSR template to the client, ask it to sign it, and then submit the CSR on the client’s behalf
        *   Support turning off individual signers to make it easier not to have to deploy a custom signer
            *   Writing custom signers for built-in signer names is hard
            *   Enable/disable built-in signers individually
    *   Kubelet serving certificate rotation questions
        *   feature not planned to GA for 1.19 because of open questions
        *   little-to-no guidance for approving SANs
        *   possible kubelet misbehavior when issued certs that don't match SANs (experience report from mhausler)
        *   liggitt: in 1.20, strengthen documentation of requirements, consider ways to improve kubelet behavior to avoid many unapproved/unissued CSRs being created
    *   [ao] Secret Type questions
        *   Does anyone know how the Secret types (tls, basic-auth, etc.) came to be?
            *   We haven’t seen much usage in the community.
            *   They are not documented anywhere that we found, only seen by looking in the code.
        *   What’s their  historical context?
            *   What were they built to solve?
        *   side discussion about encryption at rest in custom resources - api-machinery discussion and general agreement in principle, but has not been prioritized [https://docs.google.com/document/d/1x9RNaaysyO0gXHIr1y50QFbiL1x8OWnk2v3XnrdkT5Y/edit#bookmark=id.hggxaq94ueiq](https://docs.google.com/document/d/1x9RNaaysyO0gXHIr1y50QFbiL1x8OWnk2v3XnrdkT5Y/edit#bookmark=id.hggxaq94ueiq)
    *   [mvladev] Accessing OIDC discovery endpoint of the apiserver when ServiceAccountIssuerDiscovery feature gate is enabled and --anonymous-auth=false is set on the API server question
        *   for anonymous discovery clients to fetch the data from the API server, it must allow anonymous requests
        *   the issuer URL can be set to a different endpoint that is available anonymously (like a public storage bucket), and you could extract discovery/keyset info from the apiserver with a credentialed client and push the discovery/keyset data to that bucket. This also has the benefit of making the discovery data available even when the API server is not available.
*   Action Items
    *
*   Sweep issues with leftover time
    *   [unprioritized+unassigned (-needs-information) issues](https://github.com/kubernetes/kubernetes/issues?q=is%3Aissue+is%3Aopen+label%3Asig%2Fauth+-label%3Apriority%2Fawaiting-more-evidence+-label%3Apriority%2Fimportant-longterm+-label%3Apriority%2Fimportant-soon+-label%3Apriority%2Fcritical-urgent+-label%3Apriority%2Fbacklog+no%3Aassignee+-label%3Atriage%2Fneeds-information)
    *   [unprioritized+assigned (-needs-information) issues](https://github.com/kubernetes/kubernetes/issues?q=is%3Aissue+is%3Aopen+label%3Asig%2Fauth+-label%3Apriority%2Fawaiting-more-evidence+-label%3Apriority%2Fimportant-longterm+-label%3Apriority%2Fimportant-soon+-label%3Apriority%2Fcritical-urgent+-label%3Apriority%2Fbacklog+-no%3Aassignee+-label%3Atriage%2Fneeds-information)
    *   [open bugs](https://github.com/kubernetes/kubernetes/issues?utf8=✓&q=is%3Aissue+is%3Aopen+label%3Asig%2Fauth+label%3Akind%2Fbug)


## May 27, 11a - Noon (Pacific Time)



*   [Recording](https://youtu.be/-JEwkthLaM4)
*   Announcements
    *
*   Demos
    *
*   Pulls of note
    *   [https://kubernetes.io/docs/concepts/security/pod-security-standards/](https://kubernetes.io/docs/concepts/security/pod-security-standards/)
        *   liggitt: tests to validate conformance of implementations
        *   todo: highlight that PSP is still beta, without plans for GA
*   Issues of note
    *
*   Designs of note
    *
*   Discussion topic
    *   API Server per-request mTLS: [#90533](https://github.com/kubernetes/kubernetes/issues/90533)
    *   How to move forward with [External TLS certificate authenticator KEP](https://github.com/kubernetes/enhancements/pull/1749) [Slides](https://docs.google.com/presentation/d/1BDIBMPkFH_V_fB7VGPAysuqGkSf_ffgovhtGwyhvN8o/edit?usp=sharing)
        *   set up External TLS KEP reading
    *   Next steps for Managed CA Bundles: [doc](https://docs.google.com/document/d/1zC7BQdMMFRPuZpMd9Efd0nNNsTKETRuV6xDMZ2EDYr4/edit#heading=h.fg6pi87ho1ie), [#63726](https://github.com/kubernetes/kubernetes/issues/63726)
        *   Follow up in doc comments, and once high-level agreement is reached open a formal KEP
    *   Removing the deprecated insecure port
        *   [Deprecated since 1.10](https://github.com/kubernetes/kubernetes/commit/2f175bc43279a1a4552610a42af0a006f4c6fba2)
        *   Blocked on [insecure healthz](https://github.com/kubernetes/kubernetes/issues/43784)
        *   Follow up on [https://github.com/kubernetes/kubernetes/issues/91506](https://github.com/kubernetes/kubernetes/issues/91506)
*   Action Items
    *
*   Sweep issues with leftover time
    *   [unprioritized+unassigned (-needs-information) issues](https://github.com/kubernetes/kubernetes/issues?q=is%3Aissue+is%3Aopen+label%3Asig%2Fauth+-label%3Apriority%2Fawaiting-more-evidence+-label%3Apriority%2Fimportant-longterm+-label%3Apriority%2Fimportant-soon+-label%3Apriority%2Fcritical-urgent+-label%3Apriority%2Fbacklog+no%3Aassignee+-label%3Atriage%2Fneeds-information)
    *   [unprioritized+assigned (-needs-information) issues](https://github.com/kubernetes/kubernetes/issues?q=is%3Aissue+is%3Aopen+label%3Asig%2Fauth+-label%3Apriority%2Fawaiting-more-evidence+-label%3Apriority%2Fimportant-longterm+-label%3Apriority%2Fimportant-soon+-label%3Apriority%2Fcritical-urgent+-label%3Apriority%2Fbacklog+-no%3Aassignee+-label%3Atriage%2Fneeds-information)
    *   [open bugs](https://github.com/kubernetes/kubernetes/issues?utf8=✓&q=is%3Aissue+is%3Aopen+label%3Asig%2Fauth+label%3Akind%2Fbug)


## May 13, 11a - Noon (Pacific Time)



*   [Recording](https://www.youtube.com/watch?v=wGiyi74iBsw)
*   Announcements
    *
*   Demos
    *
*   Pulls of note
    *
*   Issues of note
    *
*   Designs of note
    *   Managed CA Bundles
        *   Exploration: [doc](https://docs.google.com/document/d/1zC7BQdMMFRPuZpMd9Efd0nNNsTKETRuV6xDMZ2EDYr4/edit#heading=h.fg6pi87ho1ie)
        *   Issue: [#63726](https://github.com/kubernetes/kubernetes/issues/63726)
*   Discussion topic
    *   **[start with this] **Continue discussion around next steps for dynamic audit webhooks
        *   [Dynamic audit proxy webhook design](https://docs.google.com/document/d/16cy_ZD94ooBAvlH-rFOel8RPDWRiGFg4Cz11l4sfEII/edit#)
        *   [Dynamic webhook sinks with static policy](https://docs.google.com/document/d/1MqA-RR_wUrMNMbPB6eDyghn9z3z6CDgKK2lsQcciSE8)
        *   [mo] Invert the sink model from push to pull, rough sketch:
            *   external sink scrapes API server audit logs similar to prometheus metrics / watch
            *   Similar designs: [#53455](https://github.com/kubernetes/kubernetes/issues/53455), [#64494](https://github.com/kubernetes/kubernetes/pull/64494), [community#2241](https://github.com/kubernetes/community/pull/2241)
            *   API server could just write event stream to a file and expose the file through a streaming API
                *   IO bound
                *   Limited CPU use, just stream (gzipped) bytes over network
                *   File would be non-human readable and structured for fast processing by the API server, i.e. limited metadata (i.e timestamp) for any filtering and the associated bytes of the audit event
                *   File could be encrypted using some transient in-memory key - the file would be “lost” on every restart of the API server (this would prevent concerns around storing sensitive data on disk)
                *   File could be stored on tmpfs
                *   File would be automatically truncated after it reached some “reasonable” size such as 1 GB
            *   HA API servers seems to be biggest issue
                *   API server serving the request would need to contact the other servers and multiplex their stream through to the client
        *   [mo] does it make sense to support a binary format for audit events?
        *   Enhancements freeze for v1.19 is May 19 so it would be great to have a decision on the future roadmap from dynamic audit by then
        *   Thoughts on progressing current sink-only impl to beta/ga
            *   liggitt: seems like 1/3 of a solution: provided dynamic sinks, but without understanding of what events will be sent to the sink (no policy), and with at least some existing distributors not planning to enable it as-is. There's been a lot of effort put into trying to figure out the policy and security aspects to define a path to GA, without much success.
            *   deads2k: configuring dynamic sinks to receive unknown events (since there's no policy) is difficult to reason about and hard to support
            *   micah: performance concerns around increasing audit event destinations, little customer demand to date for dynamic sinks
            *   mikedanese: security concerns about sending unfiltered audit events to arbitrary sinks in GKE
        *   AI: tallclair to start draft of email to (sig-auth, kubernetes-developers?, ...) describing intent to remove in-tree alpha sink API/impl in 1.19
    *   (time permitting) [kubernetes/enhancements#1689 Dynamic Authentication Config](https://github.com/kubernetes/enhancements/pull/1689)
    *   [mo] should the credential exec plugin API include an “installHint” help string in the ExecConfig struct
        *   Unable to connect to the server: getting credentials: exec: exec: "exec-binary-foo": executable file not found in $PATH
        *   Could be improved to “you need to install exec-binary-foo via brew install exec-binary-foo …”
        *   Not sure if this quality of life improvement is worth an API change
*   Action Items
    *
*   Sweep issues with leftover time
    *   [unprioritized+unassigned (-needs-information) issues](https://github.com/kubernetes/kubernetes/issues?q=is%3Aissue+is%3Aopen+label%3Asig%2Fauth+-label%3Apriority%2Fawaiting-more-evidence+-label%3Apriority%2Fimportant-longterm+-label%3Apriority%2Fimportant-soon+-label%3Apriority%2Fcritical-urgent+-label%3Apriority%2Fbacklog+no%3Aassignee+-label%3Atriage%2Fneeds-information)
    *   [unprioritized+assigned (-needs-information) issues](https://github.com/kubernetes/kubernetes/issues?q=is%3Aissue+is%3Aopen+label%3Asig%2Fauth+-label%3Apriority%2Fawaiting-more-evidence+-label%3Apriority%2Fimportant-longterm+-label%3Apriority%2Fimportant-soon+-label%3Apriority%2Fcritical-urgent+-label%3Apriority%2Fbacklog+-no%3Aassignee+-label%3Atriage%2Fneeds-information)
    *   [open bugs](https://github.com/kubernetes/kubernetes/issues?utf8=✓&q=is%3Aissue+is%3Aopen+label%3Asig%2Fauth+label%3Akind%2Fbug)


## April 29, 11a - Noon (Pacific Time)



*   [Recording](https://youtu.be/Ez1pt4TqliM)
*   Announcements
    *
*   Demos
    *
*   Pulls of note
    *   [micahhausler] [website#19351](https://github.com/kubernetes/website/pull/19351) Documentation PR of cluster CA rotation
        *   deads to review
*   Issues of note
    *
*   Designs of note
    *   [liggitt] CSR v1 KEP update - [#1691](https://github.com/kubernetes/enhancements/pull/1691)
        *   Liggitt: intent to wrap up feedback this week.
    *   [mo] external credential providers KEP to GA [#1711](https://github.com/kubernetes/enhancements/pull/1711)
        *   Awly: concern with dropping TLS bits from GA
        *   Mo: folding more of the current behavior into the behavior documented in the KEP.
        *   Please review! We want agreement on GA requirements for 1.19.
        *   Follow-up on conformance discussion with conformance group
            *   Suitability for inclusion into conformance is required for GA
            *   But inclusion in conformance is not required as part of GA
            *   Good test coverage is what we care about
*   Discussion topic
    *   Feedback on [kubernetes/enhancements#1689 Dynamic Authentication Config](https://github.com/kubernetes/enhancements/pull/1689)
        *   AI: Tim, Mike: provide feedback on KEP
    *   Continue discussion around next steps for dynamic audit webhooks
        *   [Dynamic audit proxy webhook design](https://docs.google.com/document/d/16cy_ZD94ooBAvlH-rFOel8RPDWRiGFg4Cz11l4sfEII/edit#)
        *   [Dynamic webhook sinks with static policy](https://docs.google.com/document/d/1MqA-RR_wUrMNMbPB6eDyghn9z3z6CDgKK2lsQcciSE8)
*   Action Items
    *
*   Sweep issues with leftover time
    *   [unprioritized+unassigned (-needs-information) issues](https://github.com/kubernetes/kubernetes/issues?q=is%3Aissue+is%3Aopen+label%3Asig%2Fauth+-label%3Apriority%2Fawaiting-more-evidence+-label%3Apriority%2Fimportant-longterm+-label%3Apriority%2Fimportant-soon+-label%3Apriority%2Fcritical-urgent+-label%3Apriority%2Fbacklog+no%3Aassignee+-label%3Atriage%2Fneeds-information)
    *   [unprioritized+assigned (-needs-information) issues](https://github.com/kubernetes/kubernetes/issues?q=is%3Aissue+is%3Aopen+label%3Asig%2Fauth+-label%3Apriority%2Fawaiting-more-evidence+-label%3Apriority%2Fimportant-longterm+-label%3Apriority%2Fimportant-soon+-label%3Apriority%2Fcritical-urgent+-label%3Apriority%2Fbacklog+-no%3Aassignee+-label%3Atriage%2Fneeds-information)
    *   [open bugs](https://github.com/kubernetes/kubernetes/issues?utf8=✓&q=is%3Aissue+is%3Aopen+label%3Asig%2Fauth+label%3Akind%2Fbug)


## April 15, 11a - Noon (Pacific Time)



*   [Recording](https://youtu.be/Fwln2YyNyBs)
*   Announcements
    *
*   Demos
    *
*   Pulls of note
    *
*   Issues of note
    *
*   Designs of note
    *   [kubernetes/enhancements#1689 Dynamic Authentication Config](https://github.com/kubernetes/enhancements/pull/1689)
*   Discussion topic
    *   [jim] [multi-tenancy benchmarks project details ](https://drive.google.com/file/d/1y3Cs79V7pUECJ_j40obb7-PzXaYRQAtE/view?usp=sharing)[15 mins]
    *   [ryan bez] Hierarchical Namespaces subproject proposal [15 mins]
        *   KEP: [https://github.com/kubernetes/enhancements/pull/1686](https://github.com/kubernetes/enhancements/pull/1686)
        *   AI: read KEP and provide guidance on next steps
        *   AI: possibly update requirements for having a repo endorsed by sig-auth
    *   [tallclair] Future of dynamic audit [25 mins]
        *   A few performance issues:
            *   [https://github.com/kubernetes/kubernetes/issues/85489](https://github.com/kubernetes/kubernetes/issues/85489)
            *   [https://github.com/kubernetes/kubernetes/issues/88042](https://github.com/kubernetes/kubernetes/issues/88042)
        *   **Action Items: **Follow up rough 1-page proposals, to be reviewed at the next meeting.
            *   [tallclair] [Dynamic audit MUX webhook proxy](https://docs.google.com/document/d/16cy_ZD94ooBAvlH-rFOel8RPDWRiGFg4Cz11l4sfEII/edit#)
            *   [pbarker] [Dynamic webhook sinks with static policy](https://docs.google.com/document/d/1MqA-RR_wUrMNMbPB6eDyghn9z3z6CDgKK2lsQcciSE8)
            *   (other alternatives welcome)
*   Action Items
    *
*   Sweep issues with leftover time
    *   [unprioritized+unassigned (-needs-information) issues](https://github.com/kubernetes/kubernetes/issues?q=is%3Aissue+is%3Aopen+label%3Asig%2Fauth+-label%3Apriority%2Fawaiting-more-evidence+-label%3Apriority%2Fimportant-longterm+-label%3Apriority%2Fimportant-soon+-label%3Apriority%2Fcritical-urgent+-label%3Apriority%2Fbacklog+no%3Aassignee+-label%3Atriage%2Fneeds-information)
    *   [unprioritized+assigned (-needs-information) issues](https://github.com/kubernetes/kubernetes/issues?q=is%3Aissue+is%3Aopen+label%3Asig%2Fauth+-label%3Apriority%2Fawaiting-more-evidence+-label%3Apriority%2Fimportant-longterm+-label%3Apriority%2Fimportant-soon+-label%3Apriority%2Fcritical-urgent+-label%3Apriority%2Fbacklog+-no%3Aassignee+-label%3Atriage%2Fneeds-information)
    *   [open bugs](https://github.com/kubernetes/kubernetes/issues?utf8=✓&q=is%3Aissue+is%3Aopen+label%3Asig%2Fauth+label%3Akind%2Fbug)


## April 1, 11a - Noon (Pacific Time)



*   [Recording](https://youtu.be/jLTU4RTShYM)
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
    *   [tallclair] Deprecation timeline for RBAC
    *   [jack] Intermediate certificates in certificates API [#1615](https://github.com/kubernetes/enhancements/pull/1615)
        *   [mike]
            *   Intermediates across different cluster’s nodes?
        *   [deads]
            *   how to expose information (structure)
            *   How is this information normally exposed?
            *   Support more intermediates than immediate signer in chain
            *   Support for multiple signers (different private keys) via the built in controller-manager
        *   [jack]
            *   Want to avoid magic/inference to determine what signed the cert
            *   AI(jack): describe the problem around file format/intermediates so we can get folks with more x509 experience to see what makes sense
    *   [mo] [remove openstack auth plugin](https://github.com/kubernetes/kubernetes/pull/89014)
    *   [mo] Deprecate all kubectl authn plugins except oidc and exec -> general feeling was that on the same release we GA this API, we deprecate the gcp and azure plugins which starts the one year timeline for removal from kubectl
        *   gcp
            *   [CJ] Slow python binary that relies on caching in the kubeconfig because it has a 500 ms start time
        *   azure
            *   It takes a long time for bugs to get fixed, ex: [#86481](https://github.com/kubernetes/kubernetes/pull/86481)
            *   Hard to determine bugs vs features, ex: [#87985](https://github.com/kubernetes/kubernetes/pull/87985)
            *   Replacement: [Azure/kubelogin: A Kubernetes credential (exec) plugin implementing azure authentication](https://github.com/Azure/kubelogin)
    *   [mo] Related to above, get exec to GA
        *   Per the [KEP](https://github.com/kubernetes/enhancements/blob/master/keps/sig-auth/20190711-external-credential-providers.md), it went beta in 1.11
            *   This is listed as a GA requirement “support for remote TLS handshakes (e.g. TPM/KMS-hosted keys)”
                *   This requirement was deemed unrelated to the rest of the functionality
        *   Feedback from users
            *   Matt Moyer noted that it is hard to pass information about the cluster to the exec plugin (hostname, CA, audience, etc)
        *   [Implementations](https://grep.app/search?q=.ExecCredential%7B&case=true) - there seem to be 10ish
            *   [mo] We should update the kube docs to include the list of known implementations to make them easier to find
                *   Puts sig-auth in a gatekeeper role?
                *   Security review implications of “blessed” plugins
                *   kubectl / sig-cli owns this list?
                *   Part of [kpt](https://opensource.googleblog.com/2020/03/kpt-packaging-up-your-kubernetes.html)?
        *   Open issues? (need to be address before GA)
            *   [#87369](https://github.com/kubernetes/kubernetes/issues/87369)
            *   [tilt/#2702](https://github.com/windmilleng/tilt/issues/2702) / [#87329 (comment)](https://github.com/kubernetes/kubernetes/pull/87329#discussion_r390369075)
            *   [#89114](https://github.com/kubernetes/kubernetes/issues/89114)
    *   Audit annotations from authentication [#89305](https://github.com/kubernetes/kubernetes/pull/89305)
        *   We should not worry about outbound annotations (i.e. defer in the authentication stack)
        *   Per Tim, Alex’s PR on metrics from authenticators [#88777](https://github.com/kubernetes/kubernetes/pull/88777) should use audit annotations





*   Action Items
    *
*   Sweep issues with leftover time
    *   [unprioritized+unassigned (-needs-information) issues](https://github.com/kubernetes/kubernetes/issues?q=is%3Aissue+is%3Aopen+label%3Asig%2Fauth+-label%3Apriority%2Fawaiting-more-evidence+-label%3Apriority%2Fimportant-longterm+-label%3Apriority%2Fimportant-soon+-label%3Apriority%2Fcritical-urgent+-label%3Apriority%2Fbacklog+no%3Aassignee+-label%3Atriage%2Fneeds-information)
    *   [unprioritized+assigned (-needs-information) issues](https://github.com/kubernetes/kubernetes/issues?q=is%3Aissue+is%3Aopen+label%3Asig%2Fauth+-label%3Apriority%2Fawaiting-more-evidence+-label%3Apriority%2Fimportant-longterm+-label%3Apriority%2Fimportant-soon+-label%3Apriority%2Fcritical-urgent+-label%3Apriority%2Fbacklog+-no%3Aassignee+-label%3Atriage%2Fneeds-information)
    *   [open bugs](https://github.com/kubernetes/kubernetes/issues?utf8=✓&q=is%3Aissue+is%3Aopen+label%3Asig%2Fauth+label%3Akind%2Fbug)


## March 18, 11a - Noon (Pacific Time)



*   [Recording](https://youtu.be/HMwpR9DD5uk)
*   Announcements
    *
*   Demos
    *   [Kyverno - policy management](https://github.com/nirmata/kyverno/blob/master/README.md), Jim
    *
*   Pulls of note
    *
*   Issues of note
    *
*   Designs of note
    *
*   Discussion topic
    *   [Policy WG Git repo request](https://github.com/kubernetes/org/issues/1669) [Jim Bugwadia]
        *   xref multi-tenancy repo request to clarify experimental nature and requirements that code in a shared repository should not be used a dependency - [#56](https://github.com/kubernetes-sigs/multi-tenancy/issues/56#issuecomment-531892603)
    *   [ritazh] subproject maturity level (CNCF maturity model)
        *   [https://github.com/cncf/toc/blob/master/process/graduation_criteria.adoc](https://github.com/cncf/toc/blob/master/process/graduation_criteria.adoc)
        *   dimensions: API surface (alpha/beta/ga), level of testing (scale tested, audited, etc)
        *   default to kubernetes process? KEP outline asks a lot of the questions you'd want answered on the way to being a mature feature
        *   [https://github.com/kubernetes/enhancements/tree/master/keps/sig-auth](https://github.com/kubernetes/enhancements/tree/master/keps/sig-auth)
        *   Examples of retroactive designs moving into KEP format for future iteration:
            *   [CSR API](https://github.com/kubernetes/enhancements/pull/1097)
            *   [service account tokens](https://github.com/kubernetes/enhancements/pull/1594)
        *   AI(ritazh): retroactive KEP for secrets-store-csi-driver
    *   [mo] [Remove basic auth support](https://github.com/kubernetes/kubernetes/pull/89069)
        *   AI(mo): see if anyone is using this on GH / probably send mail out to kube lists
*   Action Items
    *
*   Sweep issues with leftover time
    *   [unprioritized+unassigned (-needs-information) issues](https://github.com/kubernetes/kubernetes/issues?q=is%3Aissue+is%3Aopen+label%3Asig%2Fauth+-label%3Apriority%2Fawaiting-more-evidence+-label%3Apriority%2Fimportant-longterm+-label%3Apriority%2Fimportant-soon+-label%3Apriority%2Fcritical-urgent+-label%3Apriority%2Fbacklog+no%3Aassignee+-label%3Atriage%2Fneeds-information)
    *   [unprioritized+assigned (-needs-information) issues](https://github.com/kubernetes/kubernetes/issues?q=is%3Aissue+is%3Aopen+label%3Asig%2Fauth+-label%3Apriority%2Fawaiting-more-evidence+-label%3Apriority%2Fimportant-longterm+-label%3Apriority%2Fimportant-soon+-label%3Apriority%2Fcritical-urgent+-label%3Apriority%2Fbacklog+-no%3Aassignee+-label%3Atriage%2Fneeds-information)
    *   [open bugs](https://github.com/kubernetes/kubernetes/issues?utf8=✓&q=is%3Aissue+is%3Aopen+label%3Asig%2Fauth+label%3Akind%2Fbug)


## March 4, 11a - Noon (Pacific Time)



*   [Recording](https://youtu.be/ogsxnq_eXpU)
*   Announcements
    *   Migration to kubernetes-sigs/secrets-store-csi-driver is complete!
*   Demos
    *   [cristiklein] PKCS#11 support in kubectl [#64783](https://github.com/kubernetes/kubernetes/issues/64783)
        *   Mo + folks had some conversations with Cristian [on slack](https://kubernetes.slack.com/archives/C0EN96KUY/p1582313541102300?thread_ts=1582308815.101400&cid=C0EN96KUY) in regards to avoiding cgo dependencies, hopefully the demo will cover that
*   Pulls of note
    *   [enhancements#1564](https://github.com/kubernetes/enhancements/pull/1564) External token signer graduation criteria
    *   [External registry credential provider extraction](https://github.com/kubernetes/kubernetes/pull/88813)
*   Issues of note
    *
*   Designs of note
    *
*   Discussion topic
    *   [mo] Thoughts on [#88344](https://github.com/kubernetes/kubernetes/pull/88344)
    *   [Multi-tenancy Benchmarks](https://docs.google.com/document/d/1OuMwRdQbUQi02G7j4jyvscbv-6Mfa0VWjP2U_ghngHg/edit) [Jim Bugwadia]
    *   [Policy Violation CRD](https://docs.google.com/document/d/1QJWcaJdo8w88tIixiZp9zBMHKghTMrHTyHyk6xcoVhk/edit)  [Jim]
    *   [Policy WG Git repo request](https://github.com/kubernetes/org/issues/1669) [Jim]
*   Action Items
    *
*   Sweep issues with leftover time
    *   [unprioritized+unassigned (-needs-information) issues](https://github.com/kubernetes/kubernetes/issues?q=is%3Aissue+is%3Aopen+label%3Asig%2Fauth+-label%3Apriority%2Fawaiting-more-evidence+-label%3Apriority%2Fimportant-longterm+-label%3Apriority%2Fimportant-soon+-label%3Apriority%2Fcritical-urgent+-label%3Apriority%2Fbacklog+no%3Aassignee+-label%3Atriage%2Fneeds-information)
    *   [unprioritized+assigned (-needs-information) issues](https://github.com/kubernetes/kubernetes/issues?q=is%3Aissue+is%3Aopen+label%3Asig%2Fauth+-label%3Apriority%2Fawaiting-more-evidence+-label%3Apriority%2Fimportant-longterm+-label%3Apriority%2Fimportant-soon+-label%3Apriority%2Fcritical-urgent+-label%3Apriority%2Fbacklog+-no%3Aassignee+-label%3Atriage%2Fneeds-information)
    *   [open bugs](https://github.com/kubernetes/kubernetes/issues?utf8=✓&q=is%3Aissue+is%3Aopen+label%3Asig%2Fauth+label%3Akind%2Fbug)


## February 19, 11a - Noon (Pacific Time)



*   [Recording](https://youtu.be/KOHr85mPDuM)
*   Discussion topic
    *   Request to move /version to not be exposed to unauthenticated users by default policy because of information leakage / fingerprinting - [#84040](https://github.com/kubernetes/kubernetes/issues/84040)
        *   Document how to remove
        *   Do not add new endpoints to unauthenticated access going forward
    *   Request to allow specifying different service accounts per container in a pod ([#66020](https://github.com/kubernetes/kubernetes/issues/66020))
        *   Do not support multiple SAs per pod (for each container)
        *   Describe hacks to accomplish only mounting the SA token into some containers
        *   Sub-pod isolation is not a (fully) supported security boundary
            *   [Sandbox Isolation Level Decision](https://docs.google.com/document/d/1fe7lQUjYKR0cijRmSbH_y0_l3CYPkwtQa5ViywuNo8Q/edit#)
            *   [Security Container Isolation - Sandbox Layer](https://docs.google.com/document/d/1QQ5u1RBDLXWvC8K3pscTtTRThsOeBSts_imYEoRyw8A/edit#heading=h.y62q4rre7oeb)
    *   Is [#83588](https://github.com/kubernetes/kubernetes/pull/83588) safe/desirable?
        *   We need to enumerate failure modes of KMS and see which ones would benefit (or be made worse) from a restart of kube-apiserver.
        *   Alex to follow up on issue in regards to why we do not want to hard code this right now
    *   Standard Pod security "profile" definitions - [[PUBLIC] Standardized Pod Security Profiles ](https://docs.google.com/document/d/1d9c4BaDzRw1B5fAcf7gLOMZSVEvrpSutivjfNOwIqT0/edit?usp=sharing)
*   Action Items
    *
*   Sweep issues with leftover time
    *   [unprioritized+unassigned (-needs-information) issues](https://github.com/kubernetes/kubernetes/issues?q=is%3Aissue+is%3Aopen+label%3Asig%2Fauth+-label%3Apriority%2Fawaiting-more-evidence+-label%3Apriority%2Fimportant-longterm+-label%3Apriority%2Fimportant-soon+-label%3Apriority%2Fcritical-urgent+-label%3Apriority%2Fbacklog+no%3Aassignee+-label%3Atriage%2Fneeds-information)
    *   [unprioritized+assigned (-needs-information) issues](https://github.com/kubernetes/kubernetes/issues?q=is%3Aissue+is%3Aopen+label%3Asig%2Fauth+-label%3Apriority%2Fawaiting-more-evidence+-label%3Apriority%2Fimportant-longterm+-label%3Apriority%2Fimportant-soon+-label%3Apriority%2Fcritical-urgent+-label%3Apriority%2Fbacklog+-no%3Aassignee+-label%3Atriage%2Fneeds-information)
    *   [open bugs](https://github.com/kubernetes/kubernetes/issues?utf8=✓&q=is%3Aissue+is%3Aopen+label%3Asig%2Fauth+label%3Akind%2Fbug)


## February 5, 11a - Noon (Pacific Time)



*   ~~Recording~~ Forgot to hit record 😭
*   Announcements
    *
*   Demos
    *
*   Pulls of note
    *   Please review CSR signer name changes
        *   [https://github.com/kubernetes/kubernetes/pull/86476](https://github.com/kubernetes/kubernetes/pull/86476)
        *   [https://github.com/kubernetes/kubernetes/pull/86933](https://github.com/kubernetes/kubernetes/pull/86933)
        *   AI(Mo): review new changes
        *   AI(Jordan): review PRs
        *   AI(Tim): review PRs if time permits
*   Issues of note
    *   Thoughts on [https://github.com/kubernetes/kubernetes/issues/87433](https://github.com/kubernetes/kubernetes/issues/87433)
        *   Can be built out of tree, transitive permissions have rough edges that we would like to avoid in core (e.g. revoking a particular binding doesn't guarantee no actions were taken in the meantime that grant continued presence in or access to the cluster)
        *   AI(Jordan) summarize and [close out issue](https://github.com/kubernetes/kubernetes/issues/87433#issuecomment-582618496)
*   Designs of note
    *
*   Discussion topic
    *   Migrate deislabs/secrets-store-csi-driver to kubernetes-sigs
        *   [https://github.com/kubernetes/org/issues/1245](https://github.com/kubernetes/org/issues/1245)
        *   AI(Rita):
            *   criteria for being added as supported provider
            *   criteria for removal from supported provider list
                *   ex: flakey tests
            *   PR addressing this issue: [https://github.com/deislabs/secrets-store-csi-driver/pull/152](https://github.com/deislabs/secrets-store-csi-driver/pull/152)
        *   AI(Jordan): ack issue [[1](https://groups.google.com/g/kubernetes-sig-auth/c/HeNMNFbwQig/m/o8IPe1hJAwAJ)][[2](https://github.com/kubernetes/org/issues/1245#issuecomment-582638555)]
    *   KMS healthz config?
        *   [https://github.com/kubernetes/kubernetes/pull/86368](https://github.com/kubernetes/kubernetes/pull/86368) (Alex)
        *   [https://github.com/kubernetes/kubernetes/pull/86449](https://github.com/kubernetes/kubernetes/pull/86449) (Mo)
        *   Consider: adding a health check API to the KMS plugin, deferring configuration to plugin
    *   KMS DEK caching
        *   [https://github.com/kubernetes/kubernetes/pull/87438](https://github.com/kubernetes/kubernetes/pull/87438) (Alex)
        *   AI(mike): Schedule follow up conversation with Alex, Mo, Micah, and others to discuss various KMS issues (encoding, algorithm, DEK reuse, interrogating KMS for whether a key id is still valid)
        *   Possible mechanisms for avoiding exceeding encrypt count limits (e.g. for GCM): per apiserver instance DEK + incrementing nonce or max use count with random nonce.
        *   Possible mechanism for preventing a cached DEK from being used after the KEK for it has been removed from the KMS
            *   Updating the encoding format to include the KEK ID would make it easier for providers to scan for in-use KEKs
            *   Could also force a storage migration first (before the KEK is removed)
            *   Does having automatic daily storage migration of secrets reduce some of the concerns around this?
                *   The storage migrator’s state can be “reset” to migrate resources
    *   Question was asked around “is there an operator for KMS”
        *   It is hard to make a generic operator for KMS as the config state of the API servers is not exposed via the Kube API (HA environments require a very specific rotation dance to work correctly)
        *   OpenShift’s implementation which relies on state information of static pods being exposed via the Kube API: [https://github.com/openshift/library-go/commits/master/pkg/operator/encryption](https://github.com/openshift/library-go/commits/master/pkg/operator/encryption)
    *
*   Action Items
    *   AI(Mo): [https://github.com/kubernetes/kubernetes/issues/87369](https://github.com/kubernetes/kubernetes/issues/87369)
    *   KMS: After discussing the various related KMS fixes we want to make, create a unified KEP to distill the changes and motivation
*   Sweep issues with leftover time
    *   [unprioritized+unassigned (-needs-information) issues](https://github.com/kubernetes/kubernetes/issues?q=is%3Aissue+is%3Aopen+label%3Asig%2Fauth+-label%3Apriority%2Fawaiting-more-evidence+-label%3Apriority%2Fimportant-longterm+-label%3Apriority%2Fimportant-soon+-label%3Apriority%2Fcritical-urgent+-label%3Apriority%2Fbacklog+no%3Aassignee+-label%3Atriage%2Fneeds-information)
    *   [unprioritized+assigned (-needs-information) issues](https://github.com/kubernetes/kubernetes/issues?q=is%3Aissue+is%3Aopen+label%3Asig%2Fauth+-label%3Apriority%2Fawaiting-more-evidence+-label%3Apriority%2Fimportant-longterm+-label%3Apriority%2Fimportant-soon+-label%3Apriority%2Fcritical-urgent+-label%3Apriority%2Fbacklog+-no%3Aassignee+-label%3Atriage%2Fneeds-information)
    *   [open bugs](https://github.com/kubernetes/kubernetes/issues?utf8=✓&q=is%3Aissue+is%3Aopen+label%3Asig%2Fauth+label%3Akind%2Fbug)


## January 22, 11a - Noon (Pacific Time)



*   Canceled due to empty agenda


## January 8, 11a - Noon (Pacific Time)



*   [Recording](https://youtu.be/9T1IaXU-L5o)
*   Discussion topic
    *   KubeCon recap - SIG auth discussion
        *   Goals for 2020
        *   [sig auth kubecon na 2019](https://static.sched.com/hosted_files/kccncna19/2c/sig_auth_kubecon_na_2019.pdf)
    *   Changing default envelope encryption from AES-CBC to AES-GCM
        *   [https://github.com/kubernetes/kubernetes/pull/85922](https://github.com/kubernetes/kubernetes/pull/85922)
        *   [https://github.com/kubernetes/kubernetes/pull/86124](https://github.com/kubernetes/kubernetes/pull/86124)
        *   mikedanese: take opportunity to revisit encrypted encoding format
        *   consider including key id info in etcd data to enable selective reencryption (either as part of rotation, or before removal of old keys) or selective force delete (to try to clean up after uncontrolled removal of a kms key that is still needed)
    *   [nckturner] - following up on a sig cloud provider discussion with Mike and Tim at KubeCon about extracting the credential provider.  [https://github.com/kubernetes/enhancements/pull/1406](https://github.com/kubernetes/enhancements/pull/1406)
        *   Micah relay with Nick conversation during meeting
        *   Need point person from sig-node -> Tim Allclair
        *   Intermediate step could be to remove cloud provider SDK and implement some of the light functionality we use
    *   SIG-auth endorsed qualitative pod security policies
        *   Discussion at KC around a small set (3) of policies
            *   Starting point could be Kube docs @tim
            *   Could have pointers to various implementations
        *   Pod “profile” for various policy mechanisms such as PSP, OPA, SCC, etc
        *   Related: best practice audit configuration
    *   [deads2k] CSR multiple signers [https://github.com/kubernetes/enhancements/pull/1400](https://github.com/kubernetes/enhancements/pull/1400)
        *   Release notes for disabling the admission plugin and the RBAC for the existing signers
    *   [ritazh] address questions on secrets-store csi driver
*   Action Items
    *
*   Sweep issues with leftover time
    *   [unprioritized+unassigned (-needs-information) issues](https://github.com/kubernetes/kubernetes/issues?q=is%3Aissue+is%3Aopen+label%3Asig%2Fauth+-label%3Apriority%2Fawaiting-more-evidence+-label%3Apriority%2Fimportant-longterm+-label%3Apriority%2Fimportant-soon+-label%3Apriority%2Fcritical-urgent+-label%3Apriority%2Fbacklog+no%3Aassignee+-label%3Atriage%2Fneeds-information)
    *   [unprioritized+assigned (-needs-information) issues](https://github.com/kubernetes/kubernetes/issues?q=is%3Aissue+is%3Aopen+label%3Asig%2Fauth+-label%3Apriority%2Fawaiting-more-evidence+-label%3Apriority%2Fimportant-longterm+-label%3Apriority%2Fimportant-soon+-label%3Apriority%2Fcritical-urgent+-label%3Apriority%2Fbacklog+-no%3Aassignee+-label%3Atriage%2Fneeds-information)
    *   [open bugs](https://github.com/kubernetes/kubernetes/issues?utf8=✓&q=is%3Aissue+is%3Aopen+label%3Asig%2Fauth+label%3Akind%2Fbug)
