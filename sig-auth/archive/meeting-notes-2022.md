# Kubernetes SIG-Auth Meeting Agenda

## December 21st \- CANCELED

Canceled due to US holidays.

## December 7th, 11a \- Noon (Pacific Time)

- Recording  
- Announcements  
  - Last SIG Auth meeting for 2022  
  - New [\#sig-auth-triage](https://kubernetes.slack.com/archives/C04DVNR25NG) slack channel  
- Demos  
  -   
- Pulls of note  
  -   
- Issues of note  
  - \[aramase\] [Issue with k8s.io/docs/tasks/configure-pod-container/enforce-standards-admission-controller/](https://github.com/kubernetes/website/issues/37269)  
    - request for improvements and cross-linking the existing doc  
    - separate question about tightening default for unlabeled namespaces: can't do this compatibly  
    - separate question about observability of cluster defaults: this seems worthwhile to investigate  
    - Mutating admission with CEL may make “admin wants less then permissive by default”  
  - \[aramase/pacoxu\] [cluster/namespace wide environment variables inject into every container](https://github.com/kubernetes/enhancements/issues/3610)  
    - KEP PR: [https://github.com/kubernetes/enhancements/pull/3612](https://github.com/kubernetes/enhancements/pull/3612)  
- Discussion topic  
  - \[ivelichkovich\] [KEP (740)](https://github.com/kubernetes/enhancements/tree/master/keps/sig-auth/740-service-account-external-signing): [External Signing of Service Account Tokens](https://docs.google.com/document/d/177duVuRwPCBjrFot0aRptw49JDdHDdLNUGMXxSekO9w/edit#heading=h.kpy0p3xu2vb9)   
    - mike: in favor of external signing  
    - mo: concerned about complexity/reliability implications; want to make sure we apply lessons learned from kmsv1 where the shape of the API constrains the implementation  
    - micah: motivation similar to original KEP (zero-downtime rotation, no access to private key)  
    - mo: could accomplish rotation with less complexity with file-reloading  
    - mo: JWT can include certificates in signing, wondering if that should factor into the solution  
    - micah: API server is given a signing key and verifying public keyset, the API exposed by this is similar  
    - liggitt: is this enabling significantly different signing algorithms / types?  
      - micah: possibly, but should be opaque to the API server  
      - liggitt: signature blobs in the JWTs are user-facing and possibly verified client-side, so switching from very typical signatures to exotic signatures (even ones that are in-spec for JWTs) may not work everywhere service account tokens currently work  
    - mo: care less about the signature aspect and more that the API server controls the payload (because if external consumers don’t understand your signature it will just fail verification)  
    - mike: less opinionated on retaining complete payload control  
    - mike: timecheck, next step to open PR proposing change, handle questions and PRR in that PR  
  - \[JimBugwadia, Robert Ficcalgia, Jaya Ramanathan\] Policy WG update [Policy WG - KubeCon NA 2022.pptx](https://docs.google.com/presentation/d/1Se3FM5LTILuLdZaMZkcPNKivj8nnOUBl/edit#slide=id.p2)  
    - question about where to move the policy CRD used as an integration point between scanners and dashboards/UIs  
  - \[kms v2\] beta requirements  
    - tests: [https://github.com/kubernetes/kubernetes/issues/114188](https://github.com/kubernetes/kubernetes/issues/114188)   
    - On by default once beta?  
      - David and Jordan confirmed that it seemed fine for you to be able to use the KMS v2 feature without setting the feature flag once it reaches beta  
    - make it clearer which of [https://github.com/kubernetes/kubernetes/issues?q=is%3Aissue+is%3Aopen+kmsv2+](https://github.com/kubernetes/kubernetes/issues?q=is%3Aissue+is%3Aopen+kmsv2+) are required for beta  
    - Add graduation criteria to KEP  
  - \[lavalamp\] [fine-grained-authz KEP](https://github.com/kubernetes/enhancements/pull/3617/files) pre-review (followup to Sep 28 discussion)  
    - Mo: maybe it's time to consider building a multi SAR API so we can do all these checks at once  
  - \[SergeyKanzhelev\] RotateKubeletServerCertificate feature gate is one of those perma betas. In sig node we discussed if this can be GA’d in 1.27 if possible. Not clear what work is needed and who can help.  
    - last discussed in sig-auth [2021-01-06](#bookmark=id.9oa71ha4v5s6)  
    - [https://github.com/kubernetes/enhancements/issues/267\#issuecomment-755765107](https://github.com/kubernetes/enhancements/issues/267#issuecomment-755765107) were the high-level open questions remaining  
    - liggitt: functionality that exists is stable, in use, working successfully, but requires bringing your own CSR approver; it's a little weird to have a GA feature with no project-provided approver, but since kubernetes is agnostic about how nodes get IPs/DNS names, it also currently has to be agnostic about how to verify a given node owns a given IP/DNS name; I would \+1 marking the current functionality stable and deferring a project-provided node address validation / serving CSR approver to a separate effort; would be good to capture the design and production implications of the current approach in a KEP and note the remaining/future possible related work  
  - — cutoff for time  
  - \[mo\] any plans for: forward additional request metadata to \*Reviews  
    - [https://github.com/kubernetes/enhancements/pull/2843](https://github.com/kubernetes/enhancements/pull/2843)?  
  - \[riaan\] [Write e2e test for SubjectAccessReview & createAuthorizationV1NamespacedLocalSubjectAccessReview \+2 Endpoints \#114345](https://github.com/kubernetes/kubernetes/pull/114345)  
    - Thank you for the guidance Jordan & Mo.  
    - Can we please get some eye and review. \[liggitt: did a review pass\]  
- Action Items  
  -   
- Sweep issues with leftover time  
  - [CI flakes](https://storage.googleapis.com/k8s-gubernator/triage/index.html?sig=auth)  
  - [CI testgrids](https://testgrid.k8s.io/sig-auth)

## November 29th, 9a (Pacific Time), KMS-plugin meeting \#20

- Recording  
- Agenda:  
  -   
- Discussion notes:  
  - 

## November 23rd \- CANCELED

Canceled due to US holidays.

## November 9th, 11a \- Noon (Pacific Time)

- Recording  
- Announcements  
  - Code freeze was yesterday at 5 pm PDT (November 8th 2022\)  
- Demos  
  -   
- Pulls of note  
  -   
- Issues of note

  \[riaan\] The Conformance sub project have been hacking away at the K8s-Conformance technical debt for a while now, and we are near the end, with less that 2.5% remaining to be tested (10 endpoints)

  We discussed all the remaining endpoints at the last [SIG Arch meeting](https://docs.google.com/document/d/1BlmHq5uPyBUDlppYqAAzslVbAO8hilgjqZUTaNXUhKM/edit).

  2 of the endpoints are Authorization Endpoints:

* createAuthorizationV1NamespacedLocalSubjectAccessReview  
* createAuthorizationV1SubjectAccessReview

  [@tallclair](https://kubernetes.slack.com/team/U64VCBURE) made the point that it might be challenging to effectively test these endpoints without RBAC, which is not part of Conformance.

  We would like to check if there are any other thoughts / ideas to help close these endpoints as soon as possible.

  deads2k suggestion \- create a serviceaccount, issue a GET request as that serviceaccount, track 403 or success, issue a SAR for the serviceaccount, the result should match

- Designs of note  
  - \[vinayakankugoyal\][Fine grained Kubelet API authorization](https://docs.google.com/document/d/1izCeuXAZ_W6uec-Uj2sLqQodOD5HBtjrhJWT9pbIWiQ/edit?usp=sharing)   
- Discussion topic  
  - \[lavalamp\] [fine-grained-authz KEP](https://github.com/kubernetes/enhancements/pull/3617/files) pre-review (followup to Sep 28 discussion)  
- Action Items  
  -   
- Sweep issues with leftover time  
  - [CI flakes](https://storage.googleapis.com/k8s-gubernator/triage/index.html?sig=auth)  
  - [CI testgrids](https://testgrid.k8s.io/sig-auth)

## October 26th, 11a \- CANCELED

Canceled due to conflicting with KubeCon NA 2022  
Code freeze is at 5 pm PDT on November 8th 2022

## October 18th, 9a (Pacific Time), KMS-plugin meeting \#19

- Recording  
- Agenda:  
  -   
- Discussion notes:  
  - 

## October 12th, 11a \- Noon (Pacific Time)

Canceled due to empty agenda

## October 4th, 9a (Pacific Time), KMS-plugin meeting \#18

- Recording  
- Agenda:  
  -   
- Discussion notes:  
  - 

## September 28th, 11a \- Noon (Pacific Time)

- Recording  
- Announcements  
  -   
- Demos  
  -   
- Pulls of note  
  -   
- Issues of note  
  -   
- Designs of note  
  -   
- Discussion topic  
  - \[lavalamp\] subresources vs fine grained permissions.  
    - See [doc](https://docs.google.com/document/d/11g9nnoRFcOoeNJDUGAWjlKthowEVM3YGrJA3gLzhpf4/edit?resourcekey=0-OOL_NZaFGfPwnRwrx0CBwA#) (shared with sig auth mailing list).  
    - See discussion in [last sig-api-machinery meeting](https://docs.google.com/document/d/1x9RNaaysyO0gXHIr1y50QFbiL1x8OWnk2v3XnrdkT5Y/edit#bookmark=id.9w1dvjde0szd) ([recording](https://www.youtube.com/watch?v=Z-BqxaML_JE&t=29m03s))  
    - doc lists this motivation: controller operating on one aspect of an object we want to disallow from changing most other things (like limiting a controller to finalizers, liens, conditions, etc)  
      - jordan: does this cover the case where a user is otherwise allowed to create/update arbitrary fields and we want to require specific permission to touch a new field (like liens)  
    - tallclair: why can't this be addressed with admission controller  
      - complex to write, especially for end users  
      - daniel thinks of permission to modify a specific field as distinct from policy decisions which are more what admission is oriented toward  
      - mo: also interested in having these protections in place by default  
      - mike: *we* could write the admission plugin to cover in-tree field permissions (examples like gc references and csr signers)  
      - requires granting broad write permission (so the request reaches the admission plugin in the first place), then scoping the write down in admission  
    - thockin: interested in this from api review perspective trying to give people good advice to not open up new holes; not *that* bothered by N subresource approach, think that granting broad access to spec is likely to be less acceptable over time  
    - proposed solution:  
      - at authz layer, check both total object write permission AND per-field write permission, requests allowed in with per-field write permission must get re-authorized in detail once specific fields being set/modified are known  
      - checks live in apiserver code, no admission changes needed  
      - could maybe have "negative permissions" (though deads and liggitt are skeptical of this :)  
    - thockin: do field-level permission grants lock into versioned permissions (if the same field has a different field path in two different API versions)?  
      - daniel: in theory, yes… in practice, most of these fields are metadata and don't change per version; server-side apply jumped through lots of hoops to do field path ownership calculations across versions  
    - mo: we versioned PodSecurity policies; trying to understand if the meaning of granted permissions is changing over time with this proposal ("users with update can't update every field now that special liens field got added")  
      - daniel: to cover the scenario where we give a user permission to "update X except for special field Y"  
    - liggitt: update only or create as well?  
      - daniel: probably applies on create to be complete  
      - liggitt: trying to envision what it looks like for every field to be set   
    - liggitt: worried about users monkeying with fields with system implications (dropping finalizers, setting liens)  
    - david/thockin: worried about controllers with overly broad permissions getting total spec or status write permissions in order to touch a single field like finalizers or a set of annotations, conditions, etc  
      - daniel: this design covers fieldpaths to keyed lists like conditions, it doesn't cover things like annotation prefixes  
    -   
  - \[nckturner\] :45:00 Status update client exec proxy [\#2718](https://github.com/kubernetes/enhancements/issues/2718)  
    - What remains to get it approved for alpha in 1.26?  
      - mike was happy with previous shape  
      - mo comments to be addressed  
      - mo/mike take one more pass  
    - ensuring alpha status is clear (gate/apiVersion/…) is necessary, specific mechanism to do that on the client-side can be pinned down during impl review  
  - \[ahmedtd\] 00:50 [ClusterTrustBundle](https://github.com/kubernetes/enhancements/pull/3258#issuecomment-1260468546) (previously TrustAnchorSet).  
    - What remains to get it approved for alpha in 1.26?  
    - mike/david: to sweep by end of week  
  - \[robscott\] 00:56 [ReferenceGrant](https://gateway-api.sigs.k8s.io/api-types/referencegrant/) to neutral home?  
    - Last discussed in sig-auth on [Sep 15, 2021](#bookmark=id.tyjg7bln4b9s), was named ReferencePolicy at the time  
    - Used within [Gateway API](https://gateway-api.sigs.k8s.io/) to enable:  
      - Cross-namespace references from Gateways to TLS certs  
      - Cross-namespace references from Routes to backends  
    - Scheduled to graduate to beta within Gateway API in \~3 weeks  
    - SIG-Storage has decided to [use the same resource for Cross-Namespace snapshots](https://github.com/kubernetes/enhancements/pull/3295)  
      - They are [interested in moving this resource to a neutral home](https://github.com/kubernetes/enhancements/pull/3295#discussion_r977052862)  
    - questions from @liggitt  
      - is ReferenceGrant advisory? i.e. are controllers expected to already be granted broad read permission to the target resource and self-check/limit by also checking against ReferenceGrant objects?  
        - Yes, that is the current state. If there were a way to automatically enforce this that would also be great, but I think an important distinction between ReferenceGrant and RBAC is that we want grants to be revocable for references that have previously occurred.  
          - is the meaning of deleting/revoking a ReferenceGrant well-defined? are controllers expected to watch updates/deletions and break cross-namespace links/access when the grants go away?  
            - Not well documented, but yes.  
      - are all references assumed to be read-only references?  
        - Although they are read only in terms of Kubernetes resources and access, they could conceivably be used to grant write access to some underlying concept. IE in Gateway API they could allow write requests to a backend (Pod/Service) in a different namespace.  
      - being unable to scope the grant *to* a specific object in the local namespace seems problematic for a generic grant API  
        - rob: "to" has an optional "name" field: [https://gateway-api.sigs.k8s.io/references/spec/\#gateway.networking.k8s.io%2fv1alpha2.ReferenceGrantTo](https://gateway-api.sigs.k8s.io/references/spec/#gateway.networking.k8s.io%2fv1alpha2.ReferenceGrantTo), will update [https://gateway-api.sigs.k8s.io/api-types/referencegrant/\#api-design-decisions](https://gateway-api.sigs.k8s.io/api-types/referencegrant/#api-design-decisions) to clarify that  
      - if the same kind can be referenced in two different ways from another API, is group/kind granularity sufficient for the \`from\` allowlist? As a made up example, if Volumes can reference VolumeSnapshots as a source (readonly) or a destination (write), could I allow my snapshots to be used as a source but not a write destination?  
        - We have not run into that case yet, but it does seem valid. This could likely be solved by adding an additional optional field in the “from” to designate the type of reference that was being allowed, but very unclear what the structure of that would look like.  
      - more an implementation question, but is the result of the following scenarios observably different? (spoiler: I don't think it should be observably different or we leak information about existence of namespaces/objects):  
        - reference to namespace that doesn't exist  
        - reference to namespace that exists, instance that doesn't exist  
        - reference to namespace that exists, instance that exists, covering ReferenceGrant doesn't exist (looks like a RefNotPermitted condition is added?)  
          - These are not intended to be different, but we have not clearly documented that.  
      - as other areas start to consider using the same API / pattern, is there a library they can use to get the details of checking references without leaking information, honoring revocation of grants, and after-the-fact creation of grants consistently correct?  
      - does anything ensure the user creating the ReferenceGrant has permissions (read? write?) on the object they are granting access to? Translating the existing ReferenceGrant into an authz check means translating from Kind to Resource, which is unfortunate, and requires an admission check, which would need to be a webhook admission for a CRD-based ReferenceGrant resource.  
      - is "allow from all namespaces" possible/desirable to model? I tend to agree namespace selectors make it possible to overgrant accidentally, but for places that intend to allow cluster-wide reference (e.g. reference secret or config map containing trust roots from any route), is something like \`namespace: "\*"\` plausible  
  - \<meeting ended over time here, copied remaining item(s) to next agenda\>  
  - \[100mik\] Next steps for [KEP-3221](https://github.com/kubernetes/enhancements/pull/3376)  
    - We were working towards configuring webhook client auth via the new format the KEP proposes. We were considering embedding existing structs used by kubeconfig to achieve the same. Questions around it elaborated [here](https://github.com/kubernetes/enhancements/pull/3376#issuecomment-1260789023). If we reuse a struct which accepts base64 data values (which we do not want to allow). Is it sufficiently secure to disallow this using validation?  
    - CEL usage seems to be pretty coupled with CRD validation today. Would it be a blocker if we worked towards pre-filtering in a subsequent alpha release (alpha2)?  
- Action Items  
  -   
- Sweep issues with leftover time  
  - [CI flakes](https://storage.googleapis.com/k8s-gubernator/triage/index.html?sig=auth)  
  - [CI testgrids](https://testgrid.k8s.io/sig-auth)

## September 14th, 11a \- Noon (Pacific Time)

- Recording  
- Announcements  
  -   
- Demos  
  -   
- Pulls of note  
  -   
- Issues of note  
  -   
- Designs of note  
  -   
- Discussion topic  
  - We have a lot of stuff that we are proposing to do in [v1.26](https://github.com/orgs/kubernetes/projects/98?filterQuery=label%3A"sig%2Fauth") (7 KEPs)  
    - Can we make sure each of these has at least  
      - An assigned implementer  
      - An assigned approver  
  - kmsv2 ([KEP-3299](https://github.com/kubernetes/enhancements/blob/master/keps/sig-auth/3299-kms-v2-improvements/)) \- alpha 2?  
    - @enj, @ritazh, @aramase  
    - open questions  
      - Thoughts on [multi KMS encryption provider support · Issue \#111405](https://github.com/kubernetes/kubernetes/issues/111405) and [remove provider list ordering dependency within EncryptionConfiguration · Issue \#111532](https://github.com/kubernetes/kubernetes/issues/111532)  
        - Liggitt: looks like this is proposing letting you target specific encryption providers  
        - Mikedanese: not clear what the end goal is.  
      - [\[KMSv2\] Allow encryption for all resources \*/\* · Issue \#111977](https://github.com/kubernetes/kubernetes/issues/111977)  
      - hot reload EncryptionConfiguration default behavior, feature flag? [\[KMSv2\] Add support for hot reload of the \`EncryptionConfiguration\` · Issue \#111919](https://github.com/kubernetes/kubernetes/issues/111919)   
        - Nothing objectionable, but we need to figure out how to handle switch over between clients  
    - rotation [https://github.com/kubernetes/enhancements/pull/3486](https://github.com/kubernetes/enhancements/pull/3486)  
      - StorageVersion API is alpha but seems like a good spot to put encryption state. Mo has concerns about whether that API will go to Beta anytime soon.  
      - Aggregated APIs makes this a bit more complicated.  
    - [https://github.com/kubernetes/kubernetes/issues/112160](https://github.com/kubernetes/kubernetes/issues/112160) this could break existing things  
      - we can provide “action required” as part of the release note and describe migration process  
- Action Items  
  -   
- Sweep issues with leftover time  
  - [CI flakes](https://storage.googleapis.com/k8s-gubernator/triage/index.html?sig=auth)  
  - [CI testgrids](https://testgrid.k8s.io/sig-auth)  
    

## August 31st, 11a \- Noon (Pacific Time)

- Recording  
- Announcements  
  -   
- Demos  
  -   
- Pulls of note  
  -   
- Issues of note  
  -   
- Designs of note  
  -   
- Discussion topic  
  - 1.26 plans  
    - bugs  
      - [admission chain is given authorizer that doesn't include superuser fallback (\#109211)](https://github.com/kubernetes/kubernetes/issues/109211)  
        - [PR open](https://github.com/kubernetes/kubernetes/pull/111558)  
        - Should backport the fix when ready to 1.22 assuming diff is not too large  
      - [KeyEncipherment usage should not be required for non-RSA kubelet client/serving CSRs (\#109077)](https://github.com/kubernetes/kubernetes/issues/109077)  
        - [PR open](https://github.com/kubernetes/kubernetes/pull/111660)  
        - Cannot backport since not backwards compatible  
      - [exec auth breaks TLS caching (\#111911)](https://github.com/kubernetes/kubernetes/issues/111911)  
        - [PR open](https://github.com/kubernetes/kubernetes/pull/112017)  
        - Will try to backport to 1.22  
      - [reduce goroutine leakage in \`test/integration/controlplane/transformation\` (\#111674)](https://github.com/kubernetes/kubernetes/issues/111674)  
        - [PR open](https://github.com/kubernetes/kubernetes/pull/111986) \- review by aojea, wojtek-t?  
      - … any other bugs we should prioritize fixing over new features?  
    - features you are planning on working on  
      - whoami API ([KEP-3325](https://github.com/kubernetes/enhancements/blob/master/keps/sig-auth/3325-self-subject-attributes-review-api/), [PR \#111333](https://github.com/kubernetes/kubernetes/pull/111333))  
        - @nabokihms implementing, @deads2k/@liggitt reviewed  
        - already implemented, just needs rebase and release mechanics  
      - legacy token tracking alpha ([KEP-2799](https://github.com/kubernetes/enhancements/tree/master/keps/sig-auth/2799-reduction-of-secret-based-service-account-token), [PR \#108858](https://github.com/kubernetes/kubernetes/pull/108858))  
        - @zshihang implementing, @liggitt reviewer  
        - PR looks close  
      - ~~KMS observability ([KEP-3130](https://github.com/kubernetes/enhancements/blob/master/keps/sig-auth/3130-kms-observability/)) \- This one has been replaced by [KEP-3299](https://github.com/kubernetes/enhancements/blob/master/keps/sig-auth/3299-kms-v2-improvements/)~~  
      - Structured OIDC config ([KEP-3331](https://github.com/kubernetes/enhancements/pull/3332))  
        - AI: mo to ping author to see if progressing in 1.26  
          - Seems like “yes”  
      - Structured authorization config ([KEP-3376](https://github.com/kubernetes/enhancements/pull/3376))  
        - AI: mo to ping author to see if progressing in 1.26  
          - Seems like “yes”  
      - [client exec proxy](https://github.com/kubernetes/enhancements/pull/2693) (KEP-2718)  
        - mike reviewed, mo had open comments  
        - AI: mo to review responses to comments  
      - [Fine grained Kubelet API authorization · Issue \#2862 · kubernetes/enhancements · GitHub](https://github.com/kubernetes/enhancements/issues/2862)  
        - Anyone interested in trying to take this forward?  
        - AI: tallclair to look for someone to drive this  
      - … other non-GA features we should make progress on?  
    -   
  - CEL for Admission Control KEP overview (jpbetz)  
    - 1.25 promoted CEL validation of CRDs to beta  
    - draft KEP for using CEL in validation admission[CEL for Admission and Policy: Draft](https://docs.google.com/document/d/1IkK-5bXzYv2xEbrTwXsHh22sXxd_0NFF-M_rUAT3vjc/edit?resourcekey=0-b0A7F9Wz5lbXvMcqf1SFhw)  
  - [Progress on kube-rbac-proxy pre-acceptance](https://github.com/brancz/kube-rbac-proxy/issues/169)  
    - Done:  
      - Compare defaults to k8s  
      - More explicit code / api usage  
      - v1: not merged into master  
        - Remove insecure listen address  
        - Use k8s cert reloader  
    - In Progress:  
      - v1: won’t be merged into master  
        - Clean up mux logic  
        - Use apiserver filter logic  
        - Using apiserver serving options  
    - To be done:  
      - v1: won’t be merged into master  
        - Add client certs for kube-rbac-proxy to upstream  
        - Consider using upgrade-aware proxy  
          - Try to use the stb lib one if we can attest that it is not vulnerable  
- Action Items  
  -   
- Sweep issues with leftover time  
  - [CI flakes](https://storage.googleapis.com/k8s-gubernator/triage/index.html?sig=auth)  
  - [CI testgrids](https://testgrid.k8s.io/sig-auth)

## August 23rd, 9a (Pacific Time), KMS-plugin meeting \#17

- Recording  
- Agenda:  
  - Discuss KMS items for Kubernetes v1.26  
- Discussion notes:  
  - Open a tracking issue for encrypting all resources [https://github.com/kubernetes/kubernetes/issues/111977](https://github.com/kubernetes/kubernetes/issues/111977) 

## August 17th, 11a \- Noon (Pacific Time)

- Canceled, no agenda

## August 3rd, 11a \- Noon (Pacific Time)

- Recording  
- Announcements  
  -   
- Demos  
  - [Integrate Pod Security admission with Kyverno](https://github.com/kyverno/KDP/blob/main/proposals/extend_pod_security_admission.md): \[Hyokil Kim [@ToLToL](https://github.com/ToLToL), [@ChipZoller](https://github.com/chipzoller), @[JimBugwadia](https://github.com/JimBugwadia)\]   
    - Next step: open an issue on k/k to add new field to return the forbidden paths and its values.   
- Pulls of note  
  -   
- Issues of note  
  -   
- Designs of note  
  -   
- Discussion topic  
  -   
- Action Items  
  -   
- Sweep issues with leftover time  
  - [CI flakes](https://storage.googleapis.com/k8s-gubernator/triage/index.html?sig=auth)  
  - [CI testgrids](https://testgrid.k8s.io/sig-auth)

## July 20th, 11a \- Noon (Pacific Time)

- Recording  
- Announcements  
  - Code freeze coming soon  
- Demos  
  -   
- Issues of note  
  - [https://github.com/kubernetes/kubernetes/issues/109211](https://github.com/kubernetes/kubernetes/issues/109211) \- help wanted  
  - [https://github.com/kubernetes/kubernetes/issues/109077](https://github.com/kubernetes/kubernetes/issues/109077) \- PR opened  
- Pulls of note  
  - [https://github.com/kubernetes/kubernetes/pull/111061](https://github.com/kubernetes/kubernetes/pull/111061)  
- Designs of note  
  -   
- Discussion topic  
  - [https://github.com/kubernetes/kubernetes/pull/111126](https://github.com/kubernetes/kubernetes/pull/111126)   
    - Call for more reviews  
  - [https://github.com/kubernetes/kubernetes/pull/111119\#issuecomment-1184274371](https://github.com/kubernetes/kubernetes/pull/111119#issuecomment-1184274371)  
    - Discuss about KEP comment   
- Action Items  
  -   
- Sweep issues with leftover time  
  - [CI flakes](https://storage.googleapis.com/k8s-gubernator/triage/index.html?sig=auth)  
  - [CI testgrids](https://testgrid.k8s.io/sig-auth)

## July 6th, 11a \- Noon (Pacific Time)

- Recording [https://youtu.be/DVNRckZl3M4](https://youtu.be/DVNRckZl3M4)   
- Announcements  
  -   
- Demos  
  -   
- Pulls of note  
  -   
- Issues of note  
  -   
- Designs of note  
  -   
- Discussion topic  
  - \[nckturner\] Client Proxy for 1.26 [https://github.com/kubernetes/enhancements/pull/2693](https://github.com/kubernetes/enhancements/pull/2693)   
    - Mke to provide another round of review and will follow up with Mo for review when he’s back to the office  
  -   
- Action Items  
  -   
- Sweep issues with leftover time  
  - [CI flakes](https://storage.googleapis.com/k8s-gubernator/triage/index.html?sig=auth)  
    - Mike to follow up on this error from multiple jobs: [https://github.com/kubernetes/kubernetes/issues/110985](https://github.com/kubernetes/kubernetes/issues/110985)  
  - [CI testgrids](https://testgrid.k8s.io/sig-auth)

## June 22nd, 11a \- Noon (Pacific Time)

- Recording  
- Announcements  
  -   
- Demos  
  -   
- Pulls of note  
  -   
- Issues of note  
  -   
- Designs of note  
  -   
- Discussion topic  
  - Review open sig-auth v1.25 KEPs that are not merged  
    - \[ahmedtd\] [https://github.com/kubernetes/enhancements/pull/3258](https://github.com/kubernetes/enhancements/pull/3258) (Trust Anchor Sets)  
      - System:authenticated generally wants \[read,list\] on ClusterTrustBundle objects, but there exist use-cases without, so that RBAC entry should be in the documentation but not default  
      - PRR on how ClusterTrustBundles are used.  
      - Should ControllerManager reconcile the CA bundles for Client CAs, etc, into ClusterTrustBundle objects?  
      - Status fields in ClusterTrustBundles may be unnecessary and confusing \- Configmaps/secrets don’t have them. Needs explicit definition if they’re intended for use.  
      - Overall \- Punted to 1.2.6  
    -   
- Action Items  
  -   
- Sweep issues with leftover time  
  - [CI flakes](https://storage.googleapis.com/k8s-gubernator/triage/index.html?sig=auth)  
  - [CI testgrids](https://testgrid.k8s.io/sig-auth)

## June 8th, 11a \- Noon (Pacific Time)

- Recording [https://youtu.be/SNnvZVvk5VQ](https://youtu.be/SNnvZVvk5VQ)   
- Announcements  
  -   
- Demos  
  - \[stoelinga, tallclair\] [sig-auth endorsement for pspmigrator](https://groups.google.com/g/kubernetes-sig-auth/c/_mYME08jTz8)   
    - tallclair: Write a migration guide [https://kubernetes.io/docs/tasks/configure-pod-container/migrate-from-psp](https://kubernetes.io/docs/tasks/configure-pod-container/migrate-from-psp)  
      - Can’t be fully automated because of things like PSP mutating policies  
    - stoelinga: Demo of [https://github.com/samos123/pspmigrator](https://github.com/samos123/pspmigrator)  
      - Facilitates migration by reading the PSP and applying heuristics in migration guide. In some cases where the migration is simple, it prompts user to enable equivalent PSA policy (namespace labels).  
      - Rita: Is it production ready? Let’s add deprecation and risk messages at the top of the repo before ithe repo is added to the org  
      - Tallclair: The idea is that we err on the side of caution. But it’s something we will keep in mind as we develop? Add limitations sections Bare pods not supported for mutation detection  
      - If we like it we should approve the repo request here:  
        - [https://github.com/kubernetes/org/issues/3441](https://github.com/kubernetes/org/issues/3441)  
- Pulls of note  
  -   
- Issues of note  
  -   
- Designs of note  
  -   
- Discussion topic  
  - \[rita\] [https://github.com/kubernetes/enhancements/pull/3302](https://github.com/kubernetes/enhancements/pull/3302)  
    - First issue: uid, mikedanese defers to KEP authors  
    - Second issue: magic number, mikedanese to follow up  
    - Other than that KEP is LGTM  
  - \[ahmedtd\] [https://github.com/kubernetes/enhancements/pull/3258](https://github.com/kubernetes/enhancements/pull/3258) (Trust Anchor Sets)  
    - Rita: should we add this to the 1.25 enhancement tracking sheet?  
    - AI: ahmedtd to followup with sig auth leads if he thinks the KEP is ready to be tracked v1.25 release  
- Action Items  
  -   
- Sweep issues with leftover time  
  - [CI flakes](https://storage.googleapis.com/k8s-gubernator/triage/index.html?sig=auth)  
  - [CI testgrids](https://testgrid.k8s.io/sig-auth)

## May 25th, 11a \- Noon (Pacific Time)

- Recording  
- Announcements  
  -   
- Demos  
  -   
- Pulls of note  
  -   
- Issues of note  
  -   
- Designs of note  
  -   
- Discussion topic  
  - \[jpbetz\] Informational session: Overview of CEL (SIG API Machinery visitor)  
    - [SIG API Machinery- CEL meetings](https://docs.google.com/document/d/1Z7p5183OsJ1enJA4-WTk98g7agwSILqKUYxwVHZHilE/edit#heading=h.s78ck7wi4qyz) and slack: \#sig-api-machinery-cel-dev for follow-ups  
  - \[rita\] [https://github.com/kubernetes/enhancements/pull/3302](https://github.com/kubernetes/enhancements/pull/3302)  
- Action Items  
  -   
- Sweep issues with leftover time  
  - [CI flakes](https://storage.googleapis.com/k8s-gubernator/triage/index.html?sig=auth)  
  - [CI testgrids](https://testgrid.k8s.io/sig-auth)

## May 24th, 9a (Pacific Time), KMS-plugin meeting \#16

- Recording  
- Agenda:  
  -   
- Discussion notes:  
  - 

## May 11th, 11a \- Noon (Pacific Time)

- Recording  
- Announcements  
  - 1.25 enhancements freeze [June 16](https://github.com/kubernetes/sig-release/tree/master/releases/release-1.25#tldr)  
- Demos  
  -   
- Pulls of note  
  - [https://github.com/kubernetes/kubernetes/pull/109798](https://github.com/kubernetes/kubernetes/pull/109798)  
- Issues of note  
  -   
- Designs of note  
  -   
- Discussion topic  
  - \[nabokihms\] kubectl auth who-am-i  
    - Mikedanese: Google requires userinfo scope to access email, no practical concerns though  
    - Mo, deads2k: Similar APIs exist in openshift and tanzu  
    - Mo: Ideally, we create a new kind under authentication.k8s.io that lives next to TokenReview. That way we can support all other authenticators.  
    - Mikedanese: Next steps, propose changes in a KEP.  
    - Deads2k: Will this ever be required for conformance? Not sure if it should. The KEP should cover this.  
  - \[ritazh\] kms-v2-improvements KEP: [https://github.com/kubernetes/enhancements/pull/3302](https://github.com/kubernetes/enhancements/pull/3302)  
  - \[mo\] client-go TLS config KEP: [https://github.com/kubernetes/enhancements/pull/3309](https://github.com/kubernetes/enhancements/pull/3309)  
    - Deads2k: two concerns:  
      - It’s going direct to stable  
      - Not exposed in kubeconfig. That may be necessary for things like kube-apiserver webhooks.  
  - \[tallclair\] pod security conformance ([https://github.com/kubernetes/enhancements/pull/3310](https://github.com/kubernetes/enhancements/pull/3310))  
- Action Items  
  -   
- Sweep issues with leftover time  
  - [CI flakes](https://storage.googleapis.com/k8s-gubernator/triage/index.html?sig=auth)  
  - [CI testgrids](https://testgrid.k8s.io/sig-auth)

## May 10th, 9a (Pacific Time), KMS-plugin meeting \#15

- Recording  
- Agenda:  
  -   
- Discussion notes:  
  - 

## May 3rd, 9a (Pacific Time), KMS-plugin meeting \#14

- Recording  
- Agenda:  
  -   
- Discussion notes:  
  - [https://docs.google.com/document/d/1WWQH6CukWvz33wlWD5rWw4ji47g8MQW3qmkNoa6hQwM](https://docs.google.com/document/d/1WWQH6CukWvz33wlWD5rWw4ji47g8MQW3qmkNoa6hQwM)

## April 27th, 11a \- Noon (Pacific Time)

- Recording  
- Announcements  
  - 1.25 draft schedule: [https://github.com/kubernetes/sig-release/pull/1885](https://github.com/kubernetes/sig-release/pull/1885)  
- Demos  
  -   
- Pulls of note  
  -   
- Issues of note  
  -   
- Designs of note  
  -   
- Discussion topic  
  - \[tallclair\] PodSecurity GA in v1.25  
    - Experience reports from transitions would be great  
      - Openshift  
        - TODO ask Standa to link the design  
      - gke  
      - AKS \- ritazh can follow up on this  
      - others?  
    - a couple usability tweaks in 1.25 (warning if a user tries to label an exempt namespace, etc), but no fundamental changes  
    - Project board: [https://github.com/orgs/kubernetes/projects/57](https://github.com/orgs/kubernetes/projects/57)  
    - Tim: still opportunity to automate some of the migration tasks mentioned in the migration doc  
    - Tim: have concerns about the apparmor rules allowing all loaded apparmor (it is still beta and set via annotations) policies  
      - Could limit via prefix or do a deeper change \- link issue  
      - Does not block GA because we can change the meaning of this enforcement in the latest level on a new release  
    - Pod OS \- let windows pods omit fields that are linux only on the restricted policy (a future change that does not block GA)  
    - Internal code structure changes if we want the Go APIs to be stable for consumption  
  - \[mo\] thoughts on [Change apiserver healthiness check in KCM · Issue \#108014](https://github.com/kubernetes/kubernetes/issues/108014)  
    - Seems like a general problem with any component that wants to do partial healthz checks  
    - \[liggitt\] is there a way we can improve this so we do not need the knob  
    - original issue prompting switch from simple successfully discovery call to /healthz (1.10 era): [https://github.com/kubernetes/kubernetes/issues/60288](https://github.com/kubernetes/kubernetes/issues/60288)   
  - \[margo\] client-go credential plugin forceRefresh API change  
    - [ExecCredential force refresh](https://docs.google.com/document/d/1jxyIAuUKpgoQg1pzW7zIIq1TrR7r50InDTWJgjYK1CQ)  
    - seems reasonable to give exec plugins info they need to do their job  
    - credentials with a natural expiration time, but possible early revocation makes sense to inform the plugin about  
    - similar to what we had in v1alpha1 exec plugin invocation that [plumbed response info on a retry](https://github.com/kubernetes/kubernetes/blob/release-1.23/staging/src/k8s.io/client-go/pkg/apis/clientauthentication/v1alpha1/types.go#L40-L78) (so the plugin could see 401 responses and headers)  
    - boolean may not be sufficient (might need some of the info like timestamps or response code), suggest a lightweight KEP to lay out use cases and get consensus on design  
    - Taahir: would also like to see caching recommendations or capabilities added to kubectl so exec auth plugin providers don't all *have* to roll their own  
      - Jordan: could make sense, but seems separate from the force refresh request  
  - \[rohitagarwal003\] cert rotation doc minor PR review: [https://github.com/kubernetes/website/pull/32942](https://github.com/kubernetes/website/pull/32942)  
  - \[krzys\] kube-rbac-proxy review progress  
    - Pre-acceptance checklist: [https://github.com/brancz/kube-rbac-proxy/issues/169](https://github.com/brancz/kube-rbac-proxy/issues/169)  
    - Post-acceptance checklist:  
      [https://github.com/brancz/kube-rbac-proxy/issues/168](https://github.com/brancz/kube-rbac-proxy/issues/168)  
    - Help wanted to review a few specific parts related to the actual proxy and the transport  
      - [https://github.com/brancz/kube-rbac-proxy/pull/162\#discussion\_r840644641](https://github.com/brancz/kube-rbac-proxy/pull/162#discussion_r840644641)  
      - [https://github.com/brancz/kube-rbac-proxy/pull/162/files\#diff-2873f79a86c0d8b3335cd7731b0ecf7dd4301eb19a82ef7a1cba7589b5252261L238](https://github.com/brancz/kube-rbac-proxy/pull/162/files#diff-2873f79a86c0d8b3335cd7731b0ecf7dd4301eb19a82ef7a1cba7589b5252261L238)  
    -   
- Action Items  
  -   
- Sweep issues with leftover time  
  - [CI flakes](https://storage.googleapis.com/k8s-gubernator/triage/index.html?sig=auth)  
  - [CI testgrids](https://testgrid.k8s.io/sig-auth)

## April 13th, 11a \- Noon (Pacific Time)

- [Recording](https://youtu.be/bp-Nj3kXm5k)  
- Announcements  
  -   
- Demos  
  -   
- Pulls of note  
  -   
- Issues of note  
  -   
- Designs of note  
  -   
- Discussion topic  
  - \[ahmedtd\] [Trust Anchor Sets KEP](https://github.com/kubernetes/enhancements/issues/3257)  
    - will update the KEP based on feedback  
  - \[liggitt\]  
    - [admission chain is given authorizer that doesn't include superuser fallback \- Issue \#109211](https://github.com/kubernetes/kubernetes/issues/109211)  
    - [KeyEncipherment usage should not be required for non-RSA kubelet client/serving CSRs · Issue \#109077](https://github.com/kubernetes/kubernetes/issues/109077)   
  - \[tallclair\] [\[FR\] Include raw request body in audit logs · Issue \#84571](https://github.com/kubernetes/kubernetes/issues/84571)  
    - [WIP draft KEP](https://github.com/kubernetes/enhancements/pull/3277)  
    - not convinced from security perspectives it’s a good enhancement  
    - does it help with debugging for ignored fields?  
    - ignored fields are addressed by [KEP \#2885](https://github.com/kubernetes/enhancements/blob/master/keps/sig-api-machinery/2885-server-side-unknown-field-validation/README.md)  
- Action Items  
  -   
- Sweep issues with leftover time  
  - [CI flakes](https://storage.googleapis.com/k8s-gubernator/triage/index.html?sig=auth)  
  - [CI testgrids](https://testgrid.k8s.io/sig-auth)

## April 12th, 9a (Pacific Time), KMS-plugin meeting \#13

- Recording  
- Agenda:  
  -   
- Discussion notes:  
  - [https://github.com/ibihim/kms/commit/2fa84b011d8792dedf12c9d56dd8beff6eb0a2ef](https://github.com/ibihim/kms/commit/2fa84b011d8792dedf12c9d56dd8beff6eb0a2ef)

## March 30th, 11a \- Noon (Pacific Time)

- Recording  
- Announcements  
  -   
- Demos  
  -   
- Pulls of note  
  -   
- Issues of note  
  -   
- Designs of note  
  -   
- Discussion topic  
  - kubectl auth who-am-i (whoami service) [Slack thread](https://kubernetes.slack.com/archives/C0EN96KUY/p1647544566605239) (nabokihms)  
    - david/mo fine with it  
    - nabokihms to reach out to mike/jordan and see if they have concerns  
      - Mo: suggested that we code defer bootstrap RBAC policy to the cluster admin if folks are concerned  
  - webhook authorization error policy [Slack thread](https://kubernetes.slack.com/archives/C0EN96KUY/p1646799992375459) (nabokihms)  
    - Mo: same response as [k/k\#104089\_issuecomment-975763186](https://github.com/kubernetes/kubernetes/pull/104089#issuecomment-975763186)  
  - \[mo\] [k8s.io/client-go/plugin/pkg/client/auth/oidc/oidc.go](https://github.com/kubernetes/kubernetes/blob/master/staging/src/k8s.io/client-go/plugin/pkg/client/auth/oidc/oidc.go)  
    - Thoughts on refactoring this package to use a separate file to hold credentials \+ and flock on said file (and no mutexes or global cache)  
    - File would be in \~/.kube/cache/oidc/ and would honor \--cache-dir  
    - hAside from a bug fix Mo made in Dec 2019, this code has been static since Sep 2017 with known bugs / issues  
  - \[ahmedtd\] [Trust Anchor Sets KEP](https://github.com/kubernetes/enhancements/issues/3257)  
- Action Items  
  -   
- Sweep issues with leftover time  
  - [CI flakes](https://storage.googleapis.com/k8s-gubernator/triage/index.html?sig=auth)  
  - [CI testgrids](https://testgrid.k8s.io/sig-auth)

## March 29th, 9a (Pacific Time), KMS-plugin meeting \#12

- Recording  
- Agenda:  
  -   
- Discussion notes:  
  - kubectl get \--raw /api/v1 | jq .resources\[3\]  
  - 

## March 22nd, 9a (Pacific Time), KMS-plugin meeting \#11

- Recording  
- Agenda:  
  -   
- Discussion notes:  
  - 

## March 16th, 11a \- CANCELED

[Code freeze](https://groups.google.com/g/kubernetes-sig-auth/c/WJ3jbBAMVqA/m/gBRKBg4uAgAJ)

## March 15th, 9a (Pacific Time), KMS-plugin meeting \#10

- Recording  
- Agenda:  
  -   
- Discussion notes:  
  - 

## March 8th, 9a (Pacific Time), KMS-plugin meeting \#9

- Recording  
- Agenda:  
  -   
- Discussion notes:  
  - [https://github.com/ibihim/kms-proxy](https://github.com/ibihim/kms-proxy)  
  - [https://github.com/aramase/kubernetes/blob/kms-v2alpha1/staging/src/k8s.io/apiserver/pkg/storage/value/encrypt/envelope/v2alpha1/service.proto](https://github.com/aramase/kubernetes/blob/kms-v2alpha1/staging/src/k8s.io/apiserver/pkg/storage/value/encrypt/envelope/v2alpha1/service.proto)

## March 2nd, 11a \- Noon (Pacific Time)

- Recording  
- Announcements  
  -   
- Demos  
  -   
- Pulls of note  
  -   
- Issues of note  
  -   
- Designs of note  
  - \[taahm\] extracting the TrustAnchorSet object to its own KEP.  I'll try to have an open PR on the enhancements repo before next sig auth  
- Discussion topic  
  - Envelope encryption, should the DEK encryption mode be configurable or should we just use GCM  
    - [https://github.com/kubernetes/kubernetes/pull/86124](https://github.com/kubernetes/kubernetes/pull/86124)  
    - [https://github.com/kubernetes/kubernetes/pull/85922](https://github.com/kubernetes/kubernetes/pull/85922)  
  - \[micahhausler\] Conditions in RBAC  
    - Prototype: [https://github.com/kubernetes/kubernetes/compare/master...micahhausler:feature/rbac-conditions](https://github.com/kubernetes/kubernetes/compare/master...micahhausler:feature/rbac-conditions)  
    - \[tim\] would be helpful to collect use cases  
    - \[deads\] would prefer to use a different resource outside of the existing RBAC resources  
    - \[mike\] would be interesting to have a CEL based authorizer that has access to the node graph so that node restrictions could be applied to CRDs and such  
      - i.e. do not hard code the edges to the node  
  -   
- Action Items  
  -   
- Sweep issues with leftover time  
  - [CI flakes](https://storage.googleapis.com/k8s-gubernator/triage/index.html?sig=auth)  
  - [CI testgrids](https://testgrid.k8s.io/sig-auth)

## March 1st, 9a (Pacific Time), KMS-plugin meeting \#8

- Recording  
- Agenda:  
  -   
- Discussion notes:  
  - 

## Feb 16th, 11a \- Noon (Pacific Time)

- [Recording](https://youtu.be/A0U7_i4aEzM)  
- Announcements  
  -   
- Demos  
  - \[liggitt\] `kubectl create token`  
    - Want to preserve good UX for people exporting secret based tokens.  
- Pulls of note  
  -   
- Issues of note  
  - [kube-apiserver incorrectly parses tokens starting with ' ' \#106142](https://github.com/kubernetes/kubernetes/issues/106142) \- what is the expected behavior here?  
    - Anyone want to volunteer to send an email on the ML about this?  
  - Random question: is it expected that bound service account tokens stop being refreshed for pods in terminationgraceperiod? (taahm)  
    - No this is not expected. Two places where we could have an issue: (1) Kubelet sync loop is not remounting projected volume, (2) node authorizer deauthorizes kubelet TokenRequests.  
- Designs of note  
  -   
- Discussion topic  
  - Continue [Kubelet Certificate Provisioner Projected Volume Draft KEP](https://docs.google.com/document/d/1fc5ZxHpC0QieVeouS0V6Jd3NB_-xGPI0dD-rPECQYBk)  
- Action Items  
  -   
- Sweep issues with leftover time  
  - [CI flakes](https://storage.googleapis.com/k8s-gubernator/triage/index.html?sig=auth)  
    - [https://github.com/kubernetes/kubernetes/issues/105942\#issuecomment-1009416647](https://github.com/kubernetes/kubernetes/issues/105942#issuecomment-1009416647) \- OIDC discovery test failing in upgrade   
  - [CI testgrids](https://testgrid.k8s.io/sig-auth)

## Feb 15th, 9a (Pacific Time), KMS-plugin meeting \#7

- [Recording](https://youtu.be/L5ROLjd9ftw)  
- Agenda:  
  -   
- Discussion notes:  
  - 

## Feb 8th, 9a (Pacific Time), KMS-plugin meeting \#6

- [Recording](https://youtu.be/DrJILisQj1Q)  
- Agenda:  
  -   
- Discussion notes:  
  - 

## Feb 2nd, 11a \- Noon (Pacific Time)

- [Recording](https://youtu.be/VuBLtW0waC0)  
- Announcements  
  -   
- Demos  
  -   
- Pulls of note  
  - kubectl command to request bound service account tokens ([\#107880](https://github.com/kubernetes/kubernetes/pull/107880))  
- Issues of note  
  -   
- Designs of note  
  -   
- Discussion topic  
  - KMS KEP [https://github.com/kubernetes/enhancements/pull/3133](https://github.com/kubernetes/enhancements/pull/3133)  
  - [kube-rbac-proxy](https://github.com/brancz/kube-rbac-proxy): transfer ownership to sig-auth  
    - Blog post: [https://brancz.com/2018/02/27/using-kube-rbac-proxy-to-secure-kubernetes-workloads/](https://brancz.com/2018/02/27/using-kube-rbac-proxy-to-secure-kubernetes-workloads/)  
    - previously discussed [2021-03-17](#bookmark=id.9vt2e8bnbuca)  
    - similar discussion as last time  
      - no particular objections  
      - question about name  
        - 'rbac-proxy' might be a confusing name if it is actually doing subject access reviews and honors all cluster-configured authz types, not just rbac  
      - check if cluster-api/kube-builder are still using this and whether they need to (if they are using k8s.io/apisever, there's a delegating authn/authz built-in)  
    - Sergiusz/Krzysztof will champion getting this into kube this time, David will review from SIG Auth perspective, David has asked for one more reviewer  
      - Need review of exposed API surface (flags/configfile, any network-exposed APIs?)  
      - Code review?  
      - Highlight/emphasize existing audience support in examples/docs \-\> probably need docs around sending your token to this proxy, example of pod getting a projected SA token with an audience  
    - mikedanese: defaulting the audience to the serviceaccount the pod (behind the proxy) runs as would be nicer than defaulting to kube-apiserver's  
    - AI: create tracking issue, copy in discussed review items, tag drivers and reviewers  
  - [Kubelet Certificate Provisioner Projected Volume Draft KEP](https://docs.google.com/document/d/1fc5ZxHpC0QieVeouS0V6Jd3NB_-xGPI0dD-rPECQYBk) (don't need to discuss in depth, but requesting review on the doc)  
    - mikedanese: trust anchors sound like this feature request: [https://github.com/kubernetes/kubernetes/issues/63726](https://github.com/kubernetes/kubernetes/issues/63726)  
    - Important questions regarding kubelet being able to approve the CSR and the signer validating the CSR is okay  
      - Mike: notes that create CSR is not sensitive, Clayton was concerned about DOS against the cluster via CSR creation  
    - Clayton: support of use cases where the pod does not see the private key  
      - Ask: a set of use cases, who else may need to use this private key  
        - Taahir: mTLS, serving certs  
- Action Items  
  -   
- Sweep issues with leftover time  
  - [CI flakes](https://storage.googleapis.com/k8s-gubernator/triage/index.html?sig=auth)  
  - [CI testgrids](https://testgrid.k8s.io/sig-auth)

## Jan 25th, 9a (Pacific Time), KMS-plugin meeting \#5

- [Recording](https://youtu.be/yDMIuFtgW1o)  
- Agenda:  
  -   
- Discussion notes:  
  - \[Mo\] The metrics are not yet documented. We should document this.  
  - \[Mo\] We should document the UID enhancement and metrics  
  - Performance enhancements  
    - \[Damien\] Not all resources are namespaced, so we will not be able to use a DEK per namespace  
    - \[Mo\] a reference library implementation using the hierarchy in KMS plugin  
    - 

## Jan 19th, 11a \- Noon (Pacific Time)

- [Recording](https://youtu.be/lAY3ozEfU9o)  
- Announcements  
  - 1.24 Enhancements Freeze on 18:00 PT Thursday February 3rd with PRR soft freeze on Thursday January 27th.  
    - ping in \#sig-auth slack to get things targeting 1.24 added to the [feature tracking sheet](https://docs.google.com/spreadsheets/d/1T21mUTvPh70NB2eseHjCyD4LgRjyxWI9Bd1SoP8zAwA/edit#gid=1954476102) as well as KEPs reviewed/merged  
- Demos  
  -   
- Pulls of note  
  -   
- Issues of note  
  -   
- Designs of note  
  -   
- Discussion topic  
  - \[ritazh, aramase\] KMS observability KEP [https://github.com/kubernetes/enhancements/pull/3133](https://github.com/kubernetes/enhancements/pull/3133)   
    - \[mo\] kms feature is beta but encryptionconfig api is v1, how does graduation work for this KEP  
    - \[liggitt\]adding data to audit and backend, feature gate provides opportunity to report issues/feedback; may need to add grpc exemptions to api graduation; but this is a low risk change; 1 month turnaround if there is an issue. Reach out to Clayton for review and oversight of all encryption related enhancements  
    - \[alex alten\] what is the generator for UID  
    - \[mo\] KEP references [https://github.com/kubernetes/kubernetes/blob/e9e669aa6037c380469b45200e59cff9b52d6d68/staging/src/k8s.io/apiserver/pkg/admission/plugin/webhook/request/admissionreview.go\#L137](https://github.com/kubernetes/kubernetes/blob/e9e669aa6037c380469b45200e59cff9b52d6d68/staging/src/k8s.io/apiserver/pkg/admission/plugin/webhook/request/admissionreview.go#L137)   
    - \[mo\] wrapper for transformation  
  - \[Jan Šafránek, Jonathan Dobson, sig storage\] In-line CSI ephemeral volumes KEP ([PR](https://github.com/kubernetes/enhancements/pull/3158))  
    - \[Jan\] We will update the documentation to include the security aspects of inline CSI volumes and recommend CSI driver vendors not implement inline volumes for persistent storage unless they also provide a 3rd party pod admission plugin. Any concerns here? We can give guidance but not enforce.  
    - \[liggitt\] each driver can potentially expose unsafe operations. It’s up to each driver implementation  
    - \[dead2k\] hard for operators to know all the vulnerabilities a driver can be exposing  
    - \[liggitt\] would be good to have docs outlining the security concerns for a driver using pv/pvc vs inline to provide guidance  
    - [https://kubernetes-csi.github.io/docs/ephemeral-local-volumes.html](https://kubernetes-csi.github.io/docs/ephemeral-local-volumes.html)   
    - Engaged with drivers that have these known concerns  
  - \[tallclair, liggitt\] v1.24 PodSecurity Plans \- to GA or not to GA?  
    - [Tracking project](https://github.com/orgs/kubernetes/projects/57)  
    - 1.23 beta  
    - scenarios:  
      - GA in 1.24  
        - pro: allows PSP → PodSecurity migration in 1.24 timeframe to be onto a GA thing  
        - pro: allows use in clusters which disable beta features (which isn't common based on PRR survey data)  
        - con: minimal current signal from 1.23 use  
      - GA in 1.25 (preferred based on discussion)  
        - pro: gives more time for signal from 1.23+ use  
        - neutral: PSP → PodSecurity migration in 1.24 timeframe is from one beta thing to another beta thing  
        - con?: super conservative users might wait for PodSecurity to GA before enabling or starting to use it  
    - 1.24 priorities:  
      - make Kubernetes e2e tests coexist with PodSecurity (GA blocker)  
      - seek out feedback on migration/policy coexistence scenarios from users (openshift, … others?)  
      - docs review  
    - \[deads2k\] Have not gotten a lot of user feedback. Worrying for removal  
    - \[liggitt\] should beta features be turned on by default?  
    - \[dead2k\] how will both enforcements running together impact users; how do you have both enforcement mechanisms at the same time  
    - \[ritazh\] how does PodSecurity GA timeline impact PSP removal? Because PSP is beta, they are independent of each other.  
    - outstanding issues:  
      - [dedupe overlapping forbidden messages](https://github.com/kubernetes/kubernetes/issues/106129)  
      - docs on library code  
      - e2e test adjustments  
    - related timelines:  
      - 1.23: PodSecurity just now reaching end-users  
      - 1.24: users must migrate off PSP (which is beta-level)  
      - 1.25: PSP removal (independent of PodSecurity GA)  
      - 1.25: PodOS GA (planned), PodSecurity updates for PodOS (when oldest kubelet is 1.23, which enforces PodOS if present)  
- Action Items  
  -   
- Sweep issues with leftover time  
  - [CI flakes](https://storage.googleapis.com/k8s-gubernator/triage/index.html?sig=auth)  
  - [CI testgrids](https://testgrid.k8s.io/sig-auth)

## Jan 18th, 9a (Pacific Time), KMS-plugin meeting \#4

- [Recording](https://youtu.be/BIFSdUeUd2Q)  
- Agenda:  
  -   
- Discussion notes:  
  - We did not take notes but had some questions to discuss on the next SIG meeting

## Jan 11th, 9a (Pacific Time), KMS-plugin meeting \#3

- [Recording](https://youtu.be/EjmQVYEqHrI)  
- Agenda:  
  -   
- Discussion notes:  
  - Reference KMS library implementation  
  - KMS observability KEP  
    - UID generated from the API server and passed to KMS  
    - All open tracing support  
    - Current status:  
      - Working on the KEP before Jan 20  
      - Focus on newly generated UID for now as a result of last sig auth call, once we figure out how to leverage the k8s audit-id, we could replace the UID with the k8s audit-id  
        - Audit id can be configured by end users, which might not be consistent enough to chain  
        - Cache hits no guarantee every request will hit kms plugin  
        - Focus on api server request to kms plugin using a UID  
      - How do we make sure the UID is always logged in the api server log  
      - Will open the PR this week  
      - Add this to topics for 1/19 sig auth call  
  - Figuring out the history of Audit-ID and why it is under the control of end-users  
    - Had some original desire to have kubectl reuse audit-id when following redirect but that was dropped as a bad idea  
    - We could reach out to MLs to see if anyone is relying on it  
    - Then write KEP to describe the change, only trusted actors can set audit-id. We could introduce this with a new flag to introduce the behavior.  
      - How do we make sure an actor is trusted?  
  - Recovery  
    - Can we use observability data to alert users a secret cannot be decrypted anymore? Then user can initiate delete from etcd.  
    - Hold a lease  
  - Rotation  
    - Can this be automated/orchestrated by the kms plugin via the grpc api rotation policy?  
  - Action item for all: review observability KEP once it is available. Add questions and proposals for performance so we can discuss asyncly and in next weekly’s call

## Jan 5th, 11a \- Noon (Pacific Time)

- [Recording](https://youtu.be/KxX4BLHTcLU)  
- Announcements  
  - Reminder: KEPs targeting 1.24 should be ready for review by mid-January. Feature freeze is expected by end of month, so anything we want to get in needs to be making progress towards approval in the last two weeks.  
- Demos  
  -   
- Pulls of note  
  -   
- Issues of note  
  -   
- Designs of note  
  -   
- Discussion topic  
  - KMS observability related questions \- goal: make it possible to correlate kubernetes requests with KMS plugin actions  
    - \[Rita\] why is admission request generating its own uid instead of using audit ID? [https://github.com/kubernetes/kubernetes/blob/master/staging/src/k8s.io/apiserver/pkg/admission/plugin/webhook/request/admissionreview.go\#L137](https://github.com/kubernetes/kubernetes/blob/master/staging/src/k8s.io/apiserver/pkg/admission/plugin/webhook/request/admissionreview.go#L137)  
      - \[liggitt\] Wanted to ensure the request has gone thru admission review, not a 1:1 mapping  
      - \[Mo\] do we want to correlate k8s audit ID with invocation of admission webhooks  
      - \[mo\] want to understand why kms was invoked and the chain of the request; could we use the k8s audit id? But reliably without user modifications  
    - Is tracing / OpenTelemetry the way forward for tracking requests? [https://kubernetes.io/blog/2021/09/03/api-server-tracing/](https://kubernetes.io/blog/2021/09/03/api-server-tracing/)   
      - \[liggitt\] tracing is done in sample size, prob not what we want  
      -   
    - Why are we okay with a user setting audit ID on the request with no verification?  Could we make this better?  
      - Relevant issue: [Audit ID Chain · Issue \#101597](https://github.com/kubernetes/kubernetes/issues/101597)  
      - \[mo\] would we want a new uid for kms requests?  
      - [https://github.com/kubernetes/kubernetes/issues/95306\#issuecomment-703951126](https://github.com/kubernetes/kubernetes/issues/95306#issuecomment-703951126)   
      - \[liggitt\] do we want to push down the storage layer to correlate with audit id  
      - \[mike\] Not all calls to the KMS correspond to an audit log and vice versa. For reads:  
        - The watch cache loads using the loopback client and reads within the RV window are served from cache. Initial load is not audit logged but does trigger KMS calls. Subsequent reads are audit logged but do not trigger calls to KMS.  
        - Same with DEK cache. For linearized reads (default) that go to etcd, if we get back a DEK we know about, we won’t go to KMS.  
        - Writes should generally be positioned to trigger both audit events and KMS calls.  
    - Open questions:  
      - Why was audit ID allowed to be under user control?  Need to ask Tim about the history behind this  
      - If audit ID can be tightened, can/should we use it for KMS?  If not, what should KMS use instead?  
      - Does the DEK caching from the initial watch cache load of secrets via the loopback client limit the usefulness of this tracking?  
    -   
  - \[JamesL\] 1.24 release team hello  
    -   
  - \[Adam Kaplan\] Follow up on [CSI Ephemeral Volume PodSecurityAdmission Control](https://docs.google.com/document/d/1FSHyYrb0_zWL3d1KZRA_6aXfhF8N-v9D-ggYAidHkwE)  
    - \[deak2k\] want to see the list of csi drivers that are impacted   
    - \[liggitt\]   
      - inline CSI drivers that are currently exposing unsafe parameters to pod authors are a problem to address, but I expected the problem to be addressed by the drivers stopping exposing unsafe params to pod authors, and giving guidance to new drivers to make inline drivers safe, not adding built-in policy to prevent use of unsafe drivers  
      - Want to push drivers to stop exposing unsafe params to inline volumes instead of building in an admission control.   
      - GenericEphemeralVolume is GA, and limits use to what is exposed to the same options exposed by the PVC API… this allows CSI drivers to support inline use by pods without exposing unsafe parameters, and should be preferred for general storage  
      - CSIInlineVolume is still beta. My expectation is that it should only be used for parameters/inputs that are safe to expose to restricted namespaces. It sounds like this is not documented prominently.  
      -   
- Action Items  
  -   
- Sweep issues with leftover time  
  - [CI flakes](https://storage.googleapis.com/k8s-gubernator/triage/index.html?sig=auth)  
  - Liggitt is following up with folks on ​​[https://github.com/kubernetes/kubernetes/issues/105942](https://github.com/kubernetes/kubernetes/issues/105942)   
    - emailed [Walter Fender](mailto:wfender@google.com), will sync up next week  
  - [CI testgrids](https://testgrid.k8s.io/sig-auth)

## Jan 4th, 9a (Pacific Time), KMS-plugin meeting \#2

- [Recording](https://youtu.be/AxPHZ-rABIc)  
- Agenda:  
  -   
- Discussion notes:  
  - We did not take notes but had some questions to discuss on the next SIG meeting  
  - Links from the meeting:  
    - [https://github.com/kubernetes/kubernetes/blob/master/staging/src/k8s.io/apiserver/pkg/storage/value/encrypt/envelope/grpc\_service.go\#L153](https://github.com/kubernetes/kubernetes/blob/master/staging/src/k8s.io/apiserver/pkg/storage/value/encrypt/envelope/grpc_service.go#L153)  
    - [https://github.com/kubernetes/kubernetes/blob/master/staging/src/k8s.io/apiserver/pkg/server/options/encryptionconfig/config.go\#L368](https://github.com/kubernetes/kubernetes/blob/master/staging/src/k8s.io/apiserver/pkg/server/options/encryptionconfig/config.go#L368)  
    - [https://github.com/kubernetes/kubernetes/blob/master/staging/src/k8s.io/apiserver/pkg/apis/audit/v1/types.go\#L79](https://github.com/kubernetes/kubernetes/blob/master/staging/src/k8s.io/apiserver/pkg/apis/audit/v1/types.go#L79)  
    - [https://github.com/kubernetes/kubernetes/blob/master/staging/src/k8s.io/apiserver/pkg/audit/request.go\#L48](https://github.com/kubernetes/kubernetes/blob/master/staging/src/k8s.io/apiserver/pkg/audit/request.go#L48)  
    - [https://github.com/kubernetes/kubernetes/blob/master/staging/src/k8s.io/apiserver/pkg/admission/plugin/webhook/request/admissionreview.go\#L137](https://github.com/kubernetes/kubernetes/blob/master/staging/src/k8s.io/apiserver/pkg/admission/plugin/webhook/request/admissionreview.go#L137)  
    - [https://github.com/kubernetes/kubernetes/blob/master/staging/src/k8s.io/apiserver/pkg/endpoints/filters/with\_auditid.go\#L47](https://github.com/kubernetes/kubernetes/blob/master/staging/src/k8s.io/apiserver/pkg/endpoints/filters/with_auditid.go#L47)  
    - [https://kubernetes.io/blog/2021/09/03/api-server-tracing](https://kubernetes.io/blog/2021/09/03/api-server-tracing)
