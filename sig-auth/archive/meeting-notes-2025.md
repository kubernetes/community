# Kubernetes SIG-Auth Meeting Agenda

## Dec 31, 2025 11:00 AM PST

## Canceled for holiday

## Dec 17, 2025 11:00 GMT-8

Canceled, empty agenda

## Dec 3, 2025 11:00 AM PST

- **Recording:** \<link\>  
- **Announcements**  
  - Docs should be merged by today  
- **Demos**  
  -   
- **Pulls of note**  
  -   
- **Issues of note**  
  -   
- **Designs of note**  
  -   
- **Discussion topic**  
  - \[pmengelbert\] Possible deprecation and removal of KMSv2 key\_id\_hash Metric  
    - [https://github.com/kubernetes/kubernetes/issues/134002](https://github.com/kubernetes/kubernetes/issues/134002)  
    - [https://github.com/kubernetes/kubernetes/issues/127772\#issuecomment-2386433795](https://github.com/kubernetes/kubernetes/issues/127772#issuecomment-2386433795)  
    - KeyID values are unbounded  
    - Intention was to allow discovery of which KeyIDs are in use, but old information is reported  
      - We had envisioned a scenario where the old metrics could be cleared out, and some action could be taken to repopulate the metrics, thereby reporting the relevant KeyIDs  
      - In practice, when there is encryption on everything, this becomes very expensive due to the amount of work involved in repopulating the metrics  
        - List calls no longer reach the encryption at rest layer  
      - The only other way is to restart the API server, but this is impractical and not a good recommendation  
    - Due to the above, there is little of value provided by this metric  
    - The way the metric is written causes a performance bottleneck due to lock contention  
    - Because there is little value, and fixing the bottleneck is a fairly large task, I propose deprecating and removing the KeyID metric.  
    - liggitt: the metrics are labeled as “ALPHA”, so we \*can\* remove them … do we know if people are using them in practice, and if so, what we would recommend they use instead for the ways these are being used?  
    - \[deads2k\] supportive of the removal: The current functionality doesn’t achieve the existing goals because the kube-apiserver doesn’t know when the keys are no longer used.  Those who must have this information can scan etcd, same as the apiserver. **Peter** send an email to the mailing list of this removal and rationale  
  - \[ahmedtd\] Go 1.24 ML-KEM Post-Quantum Key Exchange.  Do we need to do anything specific to enable it?  
    - liggitt: [https://go.dev/doc/go1.24\#cryptotlspkgcryptotls](https://go.dev/doc/go1.24#cryptotlspkgcryptotls) / [https://words.filippo.io/2025-state/\#gcus25slide-10.jpeg](https://words.filippo.io/2025-state/#gcus25slide-10.jpeg) seem to indicate no action is required (“The new post-quantum X25519MLKEM768 key exchange mechanism is now supported and is enabled by default when Config.CurvePreferences is nil.”, and Kubernetes doesn’t set CurvePreferences)  
    - \[resolved offline\]  
  - \[Alan Diaz\] When using a custom authorization webhook, I’m looking to confirm if the current function [ConfirmNoEscalation()](https://github.com/kubernetes/kubernetes/blob/master/pkg/registry/rbac/validation/rule.go#L53C6-L53C25) works when using a custom webhook authorizer?  
    - “ConfirmNoEscalation determines if the roles for a given user in a given namespace encompass the provided role.”   
    - Source code in question: [https://github.com/kubernetes/kubernetes/blob/master/staging/src/k8s.io/apiserver/plugin/pkg/authorizer/webhook/webhook.go\#L404](https://github.com/kubernetes/kubernetes/blob/master/staging/src/k8s.io/apiserver/plugin/pkg/authorizer/webhook/webhook.go#L404)  
    - [https://kubernetes.io/docs/reference/access-authn-authz/rbac/\#restrictions-on-role-creation-or-update](https://kubernetes.io/docs/reference/access-authn-authz/rbac/#restrictions-on-role-creation-or-update)  
    - liggitt: that method only considers permissions granted via RBAC ... an external authorizer can allow a user to make modifications by authorizing custom verbs  
      - Deads got there first on the call :)  
      - [https://github.com/deadliggitt](https://github.com/deadliggitt) is amused  
  - \[Bryce Palmer\] structured authn, support for advanced JWT context?   
    - \[mo\] Need to write a webhook  
    - \[deads2k\] start with an out of tree authn webhook to show it works and rally support  
    - [https://openid.net/specs/openid-connect-core-1\_0.html\#ClaimTypes](https://openid.net/specs/openid-connect-core-1_0.html#ClaimTypes) for distributed claims info  
  - \[Bryce Palmer\] Token attenuation  
    - Biscuit ( [https://www.biscuitsec.org/](https://www.biscuitsec.org/) ) vs JWT attenuation  
      - Biscuit would add another expression language (Datalog), worth it?  
      - Biscuit would add a new token format to handle  
      - JWT attenuation method information: [https://docs.rs/attenuable-jwt/latest/attenuable\_jwt/](https://docs.rs/attenuable-jwt/latest/attenuable_jwt/) \- can be done client-side if an attenuation key is provided as part of the initially issued JWT  
        - Would likely require defining some claims \+ values format for attenuation (?)  
    - What would token exchange look like? OK with requiring a token exchange?  
        
- **Action Items**  
  -   
- **Sweep issues with leftover time**  
  - [CI flakes](https://go.k8s.io/triage?sig=auth)  
  - [CI testgrids](https://testgrid.k8s.io/sig-auth)

## Nov 19, 2025 11:00 AM PST

- **Recording:** \<link\>  
- **Announcements**  
  - Docs freeze / blog freeze  
- **Demos**  
  -   
- **Pulls of note**  
  -   
- **Issues of note**  
  -   
- **Designs of note**  
  -   
- **Discussion topic**  
  - \[luxas\] Conditional Authorization KEP published  
    - [https://github.com/kubernetes/enhancements/pull/5684](https://github.com/kubernetes/enhancements/pull/5684)  
    -   
  - \[yue9944882\] How to safely disable the flag “service-account-extend-token-expiration” from kube-apiserver?  
    - The `serviceaccount_stale_tokens_total` metric will indicate extended token use  
    - The audit log will include [authentication.k8s.io/stale-token](http://authentication.k8s.io/stale-token) annotations indicating requests using   
    - [https://github.com/kubernetes/kubernetes/blob/master/pkg/serviceaccount/claims.go\#L265-L276](https://github.com/kubernetes/kubernetes/blob/master/pkg/serviceaccount/claims.go#L265-L276)  
  - \[everettraven ([Bryce Palmer](mailto:bpalmer@redhat.com))\] Looking for general thoughts on adding Demonstrated Proof of Possession (DPoP) support for authentication. Was chatting with David around MCP usage with kube and mentioned how this might be useful for a remote MCP server scenario where an organization wants LLMs/MCP servers to fall under the “zero-trust” bucket and prevent human-users from sharing their direct access tokens.  
    - [https://curity.io/resources/learn/dpop-overview/](https://curity.io/resources/learn/dpop-overview/)   
    - [https://datatracker.ietf.org/doc/html/rfc9449](https://datatracker.ietf.org/doc/html/rfc9449)  
    - \[luxas\] Me and Micah looked at DPoP and similar message signature protection mechanisms through in this [talk in London](https://speakerdeck.com/luxas/end-to-end-message-authenticity-in-cloud-native-systems) ([video](https://www.youtube.com/watch?v=rJacyDygVi0&ab_channel=CloudNativeRejekts)) ([RFC-9421](https://www.rfc-editor.org/rfc/rfc9421.html))  
  - \[[Taahir Ahmed](mailto:taahm@google.com)\] Temperature check — TPM-based node registration  
    - We currently solve this via a combination of exec plugins and supporting daemon for issuing certs  
    - There are only a few ways to use TPMs that make sense.  
    - Is there appetite for figuring out how to get this into Kubelet?   
    - \[luxas\] SPIFFE did some interesting stuff [https://github.com/spiffe/k8s-spiffe-workload-auth-config](https://github.com/spiffe/k8s-spiffe-workload-auth-config)  
- **Action Items**  
  -   
- **Sweep issues with leftover time**  
  - [CI flakes](https://go.k8s.io/triage?sig=auth)  
  - [CI testgrids](https://testgrid.k8s.io/sig-auth)

Nov 5, 2025 11:00 AM PST

- Cancelled, busy with release deadlines

## Oct 23, 2025 11:00 AM PDT

- Dedicated discussion for constrained impersonation, round 2  
- **Recording**  
- **Discussion topic**  
  - [https://github.com/kubernetes/kubernetes/pull/134803](https://github.com/kubernetes/kubernetes/pull/134803)  
- **Action Items**  
  - 

## Oct 22, 2025 11:00 AM PDT

- **Recording:** \<link\>  
- **Announcements**  
  - tomorrow is placeholder docs PR deadline for 1.35  
- **Demos**  
  -   
- **Pulls of note**  
  -   
- **Issues of note**  
  -   
- **Designs of note**  
  -   
- **Discussion topic**  
  - [Alec Henninger](mailto:ahenning@redhat.com) \- [ahenning@redhat.com](mailto:ahenning@redhat.com) \- In discussing conditional authorization with [David Eads](mailto:deads@redhat.com), I'd like to get the sig's opinion on possibly evaluating partially-evaluated expressions in the REST storage layer. For example, label or field selectors (or perhaps full on CEL) passed on from the authorizer all the way into the storage layer. This would theoretically support more expressive policies without requiring admission.  
    - thought experiment: instead of admission doing condition evaluation, what if storage did?  
    - e.g. list / watch / delete only the things you have access to?  
    - mo: had suggested expanding which verbs support label/field selectors, and then scope down conditions to just things expressible via field/label selectors. using field/label selection on read at authz time assumes handler will honor the selectors.   
    - mo: what is the mechanism for authorization to understand the selector in play for a particular request that storage should enforce at write time?  
      - alec: authz would evaluate existing request data; if unconditional authz was not possible, could accumulate conditions (possibly selectors) that would get passed on to storage and evaluated there before a write would be allowed  
      - mo: would the conditions get applied to new or old object?  
        - alec: not sure, probably existing object?  
    - mo: have we thought about how conditional authorization interacts with spots we do secondary authz checks and never do admission  
    - liggitt: separation between admission and storage was less important than the inputs/outputs of the admission layer being better defined and already having hook points to external policy evaluation  
  - [Radostin Stoyanov](mailto:radostin@stoyanov.io) Overview of the Checkpoint/Restore WG use-cases  
    - [https://github.com/kubernetes/community/pull/8508](https://github.com/kubernetes/community/pull/8508)   
    - Slides: [https://docs.google.com/presentation/d/15HM4gFUhix7wJa2-Ycc6JBYpIF\_aoLMM/edit?usp=sharing\&ouid=109608300285541751140\&rtpof=true\&sd=true](https://docs.google.com/presentation/d/15HM4gFUhix7wJa2-Ycc6JBYpIF_aoLMM/edit?usp=sharing&ouid=109608300285541751140&rtpof=true&sd=true)  
    - liggitt: novel aspects  
      - kubelet pushing checkpoint images \- how does the kubelet authenticate? how is the destination determined? how is the pushed image protected?  
      - pod authors indicating they opt-into their pod state being able to be exfiltrated  
      - is there any verifiable connection between original container images and checkpoint images? (base layers, shas, etc)?  
      - make the different personas clear (node admin, cluster admin, author of pods used as snapshot sources, author of pods used as snapshot restore targets, kubelet pushing checkpoint images for reuse, kubelet pulling checkpoint images for restore, etc), as well as the data flow between them  
      - address those questions before exposing new APIs anyone other than a node root user can make use of  
  - [Taahir Ahmed](mailto:taahm@google.com) \- Reviews for pod certificates PRs / agnhost rebuild  
    - [https://github.com/kubernetes/kubernetes/pull/134624](https://github.com/kubernetes/kubernetes/pull/134624) ?  
    - Other stacked PRs (links?)  
      - Metrics PR  
      - Agnhost update PR to handle mtls client server  
      - E2E test PR  
    - Mo and David to review example signers  
      - Set aside time to do this. This needs to be done before beta.  
    - liggitt: need the test PR open so we can look at that and then review the beta API changes  
- **Action Items**  
  -   
- **Sweep issues with leftover time**  
  - [CI flakes](https://go.k8s.io/triage?sig=auth)  
  - [CI testgrids](https://testgrid.k8s.io/sig-auth)

## Oct 8, 2025 11:00 AM PDT

- Cancelled, busy with release deadlines

## Sep 24, 2025 11:00 AM PDT

- Cancelled, empty agenda

## Sep 11, 2025 7:00 AM PDT

- Dedicated discussion for constrained impersonation  
- **Recording**  
- **Discussion topic**  
  - \[mo\] Do we have performance concerns regarding extra authorization checks in constrained impersonation?  
    - Depending on the number of node agents and how many requests those agents make, the overall volume could be significant  
    - Also, need to decide on the exact authz checks per [this discussion](https://github.com/kubernetes/enhancements/pull/5285#discussion_r2157558190)  
  - \[mo\] Should we enhance RBAC wildcards to support escalation checks better?  Is there a way to prevent them from being abused in ways that aren’t intended?  
    - CSR signer names would need domain/\* [https://github.com/kubernetes/kubernetes/issues/122154](https://github.com/kubernetes/kubernetes/issues/122154)   
    - New impersonation stuff likely needs something on verbs \- should we be able to support delegation on SAs? [this discussion](https://github.com/kubernetes/enhancements/pull/5285#discussion_r2145089605)  
- **Action Items**  
  - 

## Sep 10, 2025 11:00 AM PDT

- **Recording**  
- **Announcements**  
  -   
- **Demos**  
  -   
- **Pulls of note**  
  -   
- **Issues of note**  
  -   
- **Designs of note**  
  -   
- **Discussion topic**  
  - \[ahmedtd\] Key-Value pairs in PodCertificate projected volumes?  
    - \[mo\] I was actually thinking it would be a raw extension but that is beside the point  
    - \[ahmedtd\] map\[string\]string where the keys are domain-qualified  
  - \[dheeraj-coding\] \[[X.509 Certificate Authentication in StructuredAuthenticationConfiguration](https://github.com/dheeraj-coding/enhancements/tree/47a5b0020b25e7b4e4f40d04ca72dc16f3144f31/keps/sig-auth/NNN-x509-certificate-authentication-in-structured-authentication-configuration)\]  
    - Waiting to raise the KEP and wanted to get some attention and review before formally raising an issue.  
    - Working prototype: [link](https://github.com/dheeraj-coding/kubernetes/commits/kep)  
    - \[mo\] I am a \-1 to this approach as it is overly coupled, IP based restrictions are not related to x509 authn.  We have also never exposed this aspect of the request to any authenticator before.  If we do want to support something like this, I would suggest a generic approach that allows any kind of authentication to be restricted based on CEL validation rules.  
    - \[dheeraj-coding\] Thanks mo for taking a look, the StructuredAuthenticationConfiguration KEP work done by aramase is to enable CEL validation rules for all authentication types, currently we have CEL validation for jwt token authenticator and anonymous authenticator. I don’t think request level attributes are exposed to jwt token authenticator. X509 is the first new authentication mechanism other than jwt to be added. Are you suggesting we add requestValidation rules to all existing authentication types so it is generic?  
    - \[dheeraj-coding\] Updated to a new KEP for more generic request validation \[[link](https://github.com/dheeraj-coding/enhancements/tree/master/keps/sig-auth/NNN-request-validation-in-structured-authentication-configuration)\]  
    - \[dheeraj-coding\] Working on making a Google Doc to help facilitate discussion  
    - \[dheeraj-coding\] A basic doc to kickstart discussions: [link](https://docs.google.com/document/d/1PcFTXulex5aUeDX4RauYbjXVFU1V44Z2pj7xFIAfhr0/edit?usp=sharing)  
  - \[pmengelbert\] Authentication to Webhooks  
    - April 9th SIG-Auth meeting discussed this  
    - 2 previous attempts failed: [KEP-2510](https://github.com/kubernetes/enhancements/pull/2512), [KEP-658](https://github.com/kubernetes/enhancements/pull/658)  
    - Next steps: Document   
      - Summarize past attempts  
      - Summarize future paths forward  
  - \[pmengelbert\] Credential Plugin Allowlist; in kuberc? client-go? elsewhere?  
  - \[yue9944882\] [\[KMSv2\] Replace DEK seed cache with non-expiring LRU cache · Issue \#132967](https://github.com/kubernetes/kubernetes/issues/132967)  
    - \[mo\] I am unconcerned about this particular failure mode, and generally prefer an approach that does not require hard coding a max cache size (and I definitely do not want to expose any cache size config)  
    - \[mo\] useSeed has to stay forever since we don’t have automatic storage migration  
    - \[yue9944882\] Alternative change: [https://github.com/kubernetes/kubernetes/pull/133007](https://github.com/kubernetes/kubernetes/pull/133007)  
  -   
- Action Items  
  -   
- Sweep issues with leftover time  
  - [CI flakes](https://go.k8s.io/triage?sig=auth)  
  - [CI testgrids](https://testgrid.k8s.io/sig-auth)

## Sep 4, 2025 7:00 AM PDT

* Dedicated demo / discussion of conditional authorization exploration work  
* **SIG-Auth: Conditional Authorization Technical Deep Dive**  
* Thursday, September 4th: 10-11:30am ET/ 7-8:30am PT/ 14:00-15:30 UTC, [https://zoom.us/j/97828052433](https://zoom.us/j/97828052433)  
* Lucas and Micah will be presenting a 90min technical deep dive with Q\&A on their authorization work  
* References  
  * [https://github.com/upbound/kubernetes-cedar-authorizer/](https://github.com/upbound/kubernetes-cedar-authorizer/)  
  * Lucas's proof-of-concept change to K8s in [https://github.com/luxas/kubernetes/tree/conditional\_authz\_2](https://github.com/luxas/kubernetes/tree/conditional_authz_2)   
  * Slides: [Conditional Authorization, SIG Auth Deep Dive](https://docs.google.com/presentation/d/1Hj3n6jtmlSWsqWA2qVbkciJaYSNccUeQ1zWfkczapIs/edit?usp=sharing)  
* Overview  
  * Change to authorization response possibilities  
    * Currently: Allow, Deny, NoOpinion  
    * New: Allow, Deny, NoOpinion, Conditional  
  * Addition of admission-phase check of conditions that must be satisfied  
* Questions  
  * liggitt:  
    * What changes in SubjectAccessReviewStatus and AdmissionReview.{Request,Response} types are needed?  
      * SubjectAccessReviewStatus includes list of conditions. The condition data type:  
        * Effect: Allow, DenyNoOpinion, DenyRequest. (naming TBD)  
          * true for Allow will satisfy that the authorizer’s “at least one allow policy must be matched”  
          * true for DenyNoOpinion means that *this* authorizer does not authorize the request by returning NoOpinion (even if there was also a matching Allow condition. A deny policy matching takes precedence over an allow policy matching)  
          * true for DenyRequest means that the authorizer returns “Deny”, which short-circuits the request by denying. No authorizers afterwards are consulted. Deny is returned even if there was a matching allow rule.  
        * ID: Opaque string, used for e.g. debugging, logging, reason messages  
        * Condition: Opaque string, written e.g. in Cedar, CEL or an arbitrary extension language  
        * ConditionType: Indication of what is format of the condition.  
      * lucas: I had not planned any changes to AdmissionReview, but as per the discussion, I’m ok to add the conditions to AdmissionReview as well, to let opaque conditions pass through, also for conditions that are not written in something k8s would enforce natively (if we add that)  
    * Composition of multiple authorizers, especially in-tree and external webhook authorizers. How does Conditional interact with short-circuiting / unioning of the authorization chain?  
      * Currently once an authorizer allows, no further authorizers are consulted.  
      * Would a Conditional response continue checking for later authorizers? How would results from 4 authorizers of \[1. \[NoOpinion\], 2\. \[Conditional(X)\], 3\. \[Conditional(Y)\], 4\. \[Deny\]\] behave?  
        * Lucas: The first NoOpinion response is skipped, as normal. Because the request *can become* authorized. Responses 2-4 are passed from the authorizer chain with the request context, all the way to the admission phase.  
          * In the conditions enforcer, first the condition set X is evaluated. If those X evaluates to a concrete Allow or Deny, short-circuiting happens.  
          * If X evaluates to NoOpinion, condition set Y is evaluated.   
          * If Y is concrete (Allow/Deny), that response is returned.  
          * If Y is NoOpinion, the algorithm moves on to response 4\.  
          * Because response 4 is a concrete Deny, the request is in this case denied.  
        * Summary: once an unconditional allow or deny is encountered, that ends the chain checking. Order of earlier conditional allows/denies is preserved and at condition evaluation time, if one is satisfied, authorization completes at that point.  
    * How does Conditional interact with short-circuiting of a single authorizer?  
      * Example: RBAC stops checking roles/bindings once an allow is found. If an “allow with condition X” is found, it needs to keep checking, could find another “allow with condition Y” or “allow unconditionally”.  
        * lucas: If both a conditional and unconditional allow is encountered *within the same authorizer*, then that should be simplified into an unconditional allow. However, if an authorizer returns a conditional response, we need to still consider later authorizers in the chain, as per the above example.  
      * Does the authorization response have to be able to include multiple possible conditional branches?  
        * deads: If we choose opaque policies and/or multiple conditions, does this concern go away or is the question directed specifically to CEL/Cedar impls?  
          * lucas: No matter the enforcement language, multiple authorizers might return conditional responses after each other.  
    * mo: cost of not being able to short-circuit and accumulating all conditional branches that allow or deny is a concern  
      * lucas: most policies will be able to be decided to evaluate to true or false at authorization time, but indeed people might be tempted to use write conditional policies more.  
      * liggitt: the easier it is to write policies that span authz/admission, especially transparently, the more policies like that we’ll get; it’s a new capability.  
      * lucas: Indeed it would be nice to be able to decide these conditions earlier. One crazy idea would be to lazily evaluate later authorizers in the chain, if we think that is better.   
    * bpalmer: is it impossible to include the full body / existing data in SubjectAccessReview to avoid needing deferred evaluation?  
      * deads: effectively impossible, yes; reading the incoming data from the body / reading the existing data from etcd at authorization time  
      * mo: if we constrained to field / label selection addressable portions, and add support for field/label selection to writes, that would bound the set of what parts of the object can be used for conditions  
        * liggitt: are you wanting to evaluate at authorization time? doesn’t that require writers to duplicate all field-selectable fields and all labels for their writes in the URL?  
        * deads: this heavily limits the flexibility that conditional authz/admission pairs have.  
  * deads: opaque vs analyzable policies? the condition returned has to be understood by the thing enforcing the condition, but should the plumbing from authorization to admission be opaque?  
    * pro: opaque lets authorizers pair with admission that understands any condition language  
    * pro: allows evolution of new types of evaluators  
    * pro: allows linking via lookups when service providers don’t want to expose attributes for policy decisions to kube-apiserver  
    * pro: forces consideration of how different authorizers can specify different values  
    * con: apiserver can’t do intersection / covers / semantic equality checks of two opaque conditions  
      * Would this be needed when permissions themselves are managed via an external system?  Covers is used for permissions mutation in kube today.  
    * con: need to figure out how an validating admission plugin can “ack” certain conditions sets. Basically, it means that we need to provide a way for the plugin to say “the conditions of authorizer with name foo (or index i) concretely evaluate to Allow/Deny/NoOpinion”.   
      * Then the admission chain is responsible for making sure that the concretized authorizer chain responses turn into a single authorizer response according to the union authorizer logic (e.g. \[Deny,Allow \=\> Deny\], \[Allow,Deny \=\> Allow\], \[NoOpinion,Allow \=\> Allow\]), and enforce that an Allow response is returned. After concretiziation, any un-evaluated conditions are either folded to a Deny (this is always safe to do), or possibly, a NoOpinion (only safe if there were no conditions with effect DenyRequest, but TBD if we want to do)  
      * What to do if one plugin is saying “conditions of authorizer foo evaluate to Allow” and another plugin says “conditions of authorizer foo evaluate to Deny”? In that case, we should probably be fail closed and deny. However, if two plugins ack the same authorizer’s conditions as Allow and NoOpinion, I guess it’s ok to fold that to an allow?  
      * Granularity/condition type should probably be enforced on a authorizer level, i.e. one authorizer can only support exactly one condition type. Is that a fair assumption?  
    * Another possible (but less consistent and intuitive) approach would be to correlate the SubjectAccessReview and AdmissionReview webhooks with some shared, per-request key.  
      * However, this might “break” as authorization responses are cached, so the admission webhook might get an AdmissionReview with a request ID it did not see before in the authorization stage. The API server could even send the request ID of the cached authorization evaluation as well.  
      * On this note: is it a concern that the API server would be storing the conditions in memory as long as authorization responses are cached? What should the total condition and per-condition max lengths be?  
    * liggitt: Can individual conditions be “satisfied” by different admission plugins (including out of tree ones), and all must be satisfied by the end of the admission chain?  
      * lucas: It’s probably only reasonable to do that on a per-authorizer level, but yes, we can make this possible.  
  * mo:  
    * (related to Jordan’s questions above), how does this compare to the functionality available in other authorizers?  (I will describe relevant things that come to mind from Azure)  
    * lucas: The litmus test for this would probably be something like if we can re-implement the Node admission plugin as a conditional node authorizer.  
  * liggitt: comment on cedar where omitted constraints implicitly allow (RBAC requires granting resources:\[“\*”\], in cedar omitting a restriction allows all in that dimension): ABAC did the same thing and it led to scary accidental overgranting  
    * example: does allowing get on “core::nodes” also allow all subresources (like get nodes/\$name/proxy/exec/…) since they were not explicitly specified?  
      * lucas: No, there is a type for each resource+subresource combination. “core::nodes\_proxy” gives access separately to proxying functionality. Also, the language only exposes resource+subresource together, encoded like RBAC does, to avoid this footgun.  
    * \[mo\] this seems to be a fundamental problem with these type of “match on conditions” based languages  
      * lucas: Yes. To make this usable and safe, I’d require specific opt-in annotations (like Go/Rust/etc. linters) that mark “yes, I really meant to keep the resource dimension unrestricted, I did not forget it”. And only in this case the authorizer would consider the policy “valid”. Then a specific environment could even have meta-linting, restricting use of this “allow unrestricted resource” lint directive.  
  * liggitt: if we’re unifying evaluation for reads / writes, need to make sure users understand whether something is evaluated against existing or incoming (reads and deletes only have existing, creates only have incoming, updates have existing and incoming)  
    * lucas: Yes. One can statically know whether a policy author “made a mistake”, by e.g. assuming that a non-fieldselectable field is selectable. However, I was thinking whether it makes sense to start to restrict only field selectable fields in writes as well to keep parity. However, that’d restrict expressiveness a lot, and put even more pressure on making more fields selectable (which might be both a good and bad thing).  
  * Follow-up actions / questions (put your name, your questions or what you want to see the POC explore or compare to guide the next discussion)  
    * liggitt: I expected AdmissionReview.Request to include the unsatisfied conditions, and AdmissionReview.Response to indicate which conditions were satisfied, so that any admission plugin (in-tree or out-of-tree) could participate in satisfying the conditions  
    * david/mo: compare what could be constrained if the conditions were limited to label/field selectors

## Aug 27, 2025 11:00 AM PDT

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
  - 1.35 KEP planning [Needs KEP / release work \#sig-auth](https://docs.google.com/document/d/1sY8fRyRtk4eG9R439z5ao5i9bFuuxilS03XaNlqoni0)  
  - \[yue9944882\] [\[KMSv2\] Replace DEK seed cache with non-expiring LRU cache · Issue \#132967](https://github.com/kubernetes/kubernetes/issues/132967)  
    - \[mo\] I am unconcerned about this particular failure mode, and generally prefer an approach that does not require hard coding a max cache size (and I definitely do not want to expose any cache size config)  
    - \[mo\] useSeed has to stay forever since we don’t have automatic storage migration  
    - \[yue9944882\] Alternative change: [https://github.com/kubernetes/kubernetes/pull/133007](https://github.com/kubernetes/kubernetes/pull/133007)  
  - \[mo\] is it time for kubelet to start watching SAs in the same way that it does secrets/config maps?  We have three current use cases:  
    - SA tokens for image pulls  
    - Pod cert requests  
    - SA token cache invalidation when SA is deleted for running pod  
    - Verify the existing single-item watch cache manager being used for secret/configmap doesn’t have a latent bug we’d be importing to serviceaccounts [https://github.com/kubernetes/kubernetes/issues/124701](https://github.com/kubernetes/kubernetes/issues/124701)   
  - \[ritazh\] Thoughts on [https://github.com/kubernetes/kubernetes/issues/133515](https://github.com/kubernetes/kubernetes/issues/133515)   
    - moved to slack at [https://kubernetes.slack.com/archives/C0EN96KUY/p1756321182236879](https://kubernetes.slack.com/archives/C0EN96KUY/p1756321182236879)   
  - — remaining items moved to next meeting’s agenda  
- Action Items  
  -   
- Sweep issues with leftover time  
  - [CI flakes](https://go.k8s.io/triage?sig=auth)  
  - [CI testgrids](https://testgrid.k8s.io/sig-auth)

## Aug 13, 2025 11:00 AM PDT

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
  - \[aramase\] [Introduce WG Checkpoint Restore](https://github.com/kubernetes/community/pull/8508)  
    - Adding this for discussion based on the SIG Auth triage meeting. SIG Auth is added as a stakeholder.  
    - [Jordan Liggitt](mailto:liggitt@google.com) replied at [https://github.com/kubernetes/community/pull/8508/files\#r2274401485](https://github.com/kubernetes/community/pull/8508/files#r2274401485)   
  - \[rata/mrunalp/haircommander\] User Namespace PSA requirement?  
    - Open to relaxing restrictions that no longer make sense for pods which opt into user namespaces, if there’s compelling reasons to relax and relaxing doesn’t jeopardize defense-in-depth (examples: runAsRoot may be reasonable to relax, capabilities maybe less so)  
    - Likely would not require pods to opt into user namespaces for the restricted PSA profile as long as it could prevent pods from running successfully on valid current nodes  
  - \[mo\] I think we should drop the API server client cert signer from the pod cert KEP and replace it with a signer that manages serving certs for kube services  
    - \[ahmedtd\] I'd like to pull it all into a separate KEP and get the machinery to GA ASAP.  
    - \[ahmedtd\] Will share repo where I am working on mesh certs (client SPIFFE certs)  
    - .  We can add a service certificate controller there.  
  - \[mo\] secondary authorization checks do not show in the audit log, should they?  
    - Enhance audit to optionally track this for some requests if desired  
  - \[yt2985\] [https://github.com/kubernetes/kubernetes/pull/132922](https://github.com/kubernetes/kubernetes/pull/132922)   
    - The fix implemented a CA controller in the transport to periodically read the CA files for the client  
    - We need a stop channel to make the controller stop in the integration test.  
      - Use the global stop channel (wati.NeverStop, like what the certificate rotation controller does) and override it in the integration test with the context.Done() (have issue for the test/integration/apiserver/oidc/oidc\_test.go)  
      - Add a field context in the rest.Config and transport.Config and pass in the test context to these configs in the integration. This way will need to add the context field to many intermediate configs.  
    - (outcome of discussion) — we should do this via a wrapper transport, so that every n minutes, we pause the request and check for updated CA, key, and client cert.  
- Action Items  
  -   
- Sweep issues with leftover time  
  - [CI flakes](https://go.k8s.io/triage?sig=auth)  
  - [CI testgrids](https://testgrid.k8s.io/sig-auth)

## Jul 30, 2025 11:00 AM PDT

- Cancelled due to leads availability

## Jul 16, 2025 11:00 AM PDT

- Cancelled due to leads availability

## Jul 2, 2025 11:00 AM PDT

- Cancelled due to leads availability

## Jun 18, 2025 11:00 AM PDT

- Cancelled due to leads availability

## Jun 4, 2025 11:00 GMT-7

- Recording  
- Announcements  
  -   
- Demos  
  - Lucas Käldström to demonstrate deeper Kubernetes and Cedar integration, for e.g. unified authorization and admission policies, and the ability to uniformly express field/label restrictions for both reads and writes. After this (or in another meeting later), it would be great to discuss if/how we want conditional authorization in SubjectAccessReview to work  
    - Google Slides: [Conditional Authorization for Kubernetes, SIG Auth presentation](https://docs.google.com/presentation/d/1aWFaVtC_f5cKbG49AtNlmaEBbEfyI4pDDlyRZLF8MbE/edit?slide=id.g35c3f1109c7_0_52#slide=id.g35c3f1109c7_0_52)  
    - Speakerdeck:   
- Pulls of note  
  -   
- Issues of note  
  -   
- Designs of note  
  -   
- Discussion topic  
  - \[micahhausler/Jyoti\] [Expand structured authentication for additional authentication types](https://github.com/kubernetes/kubernetes/issues/131851)  
- Action Items  
  -   
- Sweep issues with leftover time  
  - [CI flakes](https://go.k8s.io/triage?sig=auth)  
  - [CI testgrids](https://testgrid.k8s.io/sig-auth)

## May 21, 2025 11:00 AM PDT

- Recording  
- Announcements  
  - \[siyuanfoundation\] [Kubernetes Emulation Version Update](https://docs.google.com/document/d/1Qgnwo24vYgmx5D01bGhMqZPW8bP8dbfQqJGygMqLsEU/edit?usp=sharing)  
  - Tentative 1.34 milestones [https://github.com/kubernetes/sig-release/pull/2780](https://github.com/kubernetes/sig-release/pull/2780)  
    - Enhancements Freeze: Friday, June 20th, 2025  
    - Code Freeze: Thursday, July 24th, 2025  
- Demos  
  -   
- Pulls of note  
  -   
- Issues of note  
  -   
- Designs of note

  - # 

- Discussion topic  
  - \[mo\] [https://github.com/kubernetes/website/pull/50687](https://github.com/kubernetes/website/pull/50687) \- How to treat proposals to add 3rd party projects to this list of multi-tenancy projects?  
    - \[liggitt\] drop the links from the page  
    - \[deads2k\] for external projects (non-kubernetes, non-kubernetes-sigs), it’s fine to drop the links  
    - \[liggitt\] for kubernetes-sigs or CNCF projects, update current state (level of stability) and keep references  
  - \[ritazh\] [https://github.com/kubernetes/kubernetes/issues/131336](https://github.com/kubernetes/kubernetes/issues/131336) \- Proposal to remove the API validation requirement that prevents using allowPrivilegeEscalation=False together with CAP\_SYS\_ADMIN  
    - Confirm with SIG node folks that Jordan’s last comment is correct  
    - If it is correct, next step would be to add warnings  
    - We recommend PSA  
  - \[palnabarun\] Is it possible to have authentication mechanisms enabled for conversion webhooks, like we have for Validating and Mutating Webhooks? Is a proposal to enable it likely to be considered? Related: [https://github.com/kubernetes/kubernetes/issues/131810](https://github.com/kubernetes/kubernetes/issues/131810) but I don’t think the approach in the issue is correct. We need to find the right place to pass the webhook authentication config file to apiextensions-apiserver. AdmissionConfiguration is not the right place.  
    - Conversion is meant to be stateless and only depend on the input, unlike admission webhooks that are based on external state  
    - Confidentiality and integrity not concerns, only availability, so the need for auth to conversion webhook is minimal  
    - Not opposed to a change here, but not high priority for each SIG Auth or API Machinery  
  - \[aramase\] [Does it make sense to use ServiceAccounts for custom resources? · Issue \#131740](https://github.com/kubernetes/kubernetes/issues/131740)  
    - We have a similar model currently with Secrets Store Sync Controller and there are many more projects out there.  
    - \[aramase\] For the sync controller, the service account is tied to the Secret Sync custom resource. This will not work based on comment \- “Future iterations of workload identity between AKS \<-\> Azure will be far more tightly constrained to pods, nodes, networks, TPM, etc so I would not recommend relying on the current lax STS semantics”.  
    - \[aramase\] One option is to make the controller namespace scoped and run an instance of the controller (packaged with provider) in every namespace the user wants to sync Kubernetes secret. The controller can only act on custom resources in the namespace. This solves the workload identity problem because the namespaced controller service account would be used.  
    - Another alternative is allowing Kubernetes to be configured in a way that other claims can be added to the JWT with restrictions. Then the cloud providers trust policy can be configured to be very strict about what claims are on the token.  
      - \[mo\] this idea has been rejected multiple times, SA tokens are a Kube API, they behave the same way everywhere.  Cloud providers don’t get a say here.  
    - \[everettraven\] OLM “v1” has this model of using ServiceAccounts for managing operators it installs  
    - \[stealthybox\] Flux controllers have a Kube identity and potentially a Cloud identity.  
      - when a controller is reconciling a namespaced resource, we want to drop privileges or switch cloud roles constrained to that namespace, resource’s serviceaccount, or the resource itself  
      - Using kubernetes TokenRequest delegates AuthN to Kubernetes via Flux’s Kube identity  
      - We’re open to using cloud specific impersonation features / session-tags instead using Flux’s Cloud identity  
- Action Items  
  -   
- Sweep issues with leftover time  
  - [CI flakes](https://go.k8s.io/triage?sig=auth)  
  - [CI testgrids](https://testgrid.k8s.io/sig-auth)

## May 7, 2025 11:00 AM PDT

- Cancelled due to light agenda, please review items planned for 1.34 such as [KEP-5284: Limit impersonate permissions](https://github.com/kubernetes/enhancements/pull/5285)

## Apr 23, 2025 11:00 AM PDT

- Recording  
- Announcements  
  - [REQUEST: Archive repo kubernetes-sigs/hierarchical-namespaces · Issue \#5484](https://github.com/kubernetes/org/issues/5484)  
  - [REQUEST: Archive repo kubernetes-sigs/pspmigrator · Issue \#5485](https://github.com/kubernetes/org/issues/5485)  
- Demos  
  -   
- Pulls of note  
  -   
- Issues of note  
  -   
- Designs of note  
  -   
- Discussion topic  
  - 1.33 released  
  - 1.34 KEP planning [Needs KEP / release work \#sig-auth](https://docs.google.com/document/d/1sY8fRyRtk4eG9R439z5ao5i9bFuuxilS03XaNlqoni0)  
  - \[mo\] [Use httptrace to validate kubelet serving cert CN](https://github.com/kubernetes/kubernetes/pull/131337)  
    - Another approach could be a ["preflight probe" request](https://github.com/kubernetes/kubernetes/pull/131337#discussion_r2056499010)  
- Action Items  
  -   
- Sweep issues with leftover time  
  - [CI flakes](https://storage.googleapis.com/k8s-gubernator/triage/index.html?sig=auth)  
  - [CI testgrids](https://testgrid.k8s.io/sig-auth)

## Apr 9, 2025 11:00 AM PDT

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
  - \[mo\] maybe it is time for us to prioritize authentication to webhooks by default  
    - [CVE-2025-1974: ingress-nginx admission controller RCE escalation](https://github.com/kubernetes/kubernetes/issues/131009)  
    - CVSS:3.1/AV:N/AC:L/PR:N/UI:N/S:U/C:H/I:H/A:H (Score: 9.8, Critical)  
    - \[liggitt\] first step is to find the old issues that talk about this; what we want to do for authentication (tls, tokens that’s audience scoped, etc.)  
    - \[liggitt\] baseline requirement is this should be generic enough such that webhooks do not need custom ways to support different providers  
    - [https://kubernetes.io/docs/reference/access-authn-authz/extensible-admission-controllers/\#authenticate-apiservers](https://kubernetes.io/docs/reference/access-authn-authz/extensible-admission-controllers/#authenticate-apiservers)   
    - [https://github.com/kubernetes/enhancements/issues/2510](https://github.com/kubernetes/enhancements/issues/2510)  
    - [https://github.com/kubernetes/enhancements/pull/658](https://github.com/kubernetes/enhancements/pull/658)  
    - [https://open-policy-agent.github.io/gatekeeper/website/docs/externaldata/\#authenticate-the-api-server-against-webhook-self-managed-k8s-cluster-only](https://open-policy-agent.github.io/gatekeeper/website/docs/externaldata/#authenticate-the-api-server-against-webhook-self-managed-k8s-cluster-only)   
    - **Need a volunteer for this**  
    - Next step is to start a doc to gather information (requirements, issues, what validations we need)  
      - Would be ideal to do this in 1.34 if we want to pick it up in 1.35  
  - \[ahmedtd\] [RFC 8693](https://www.rfc-editor.org/rfc/rfc8693.html) for kubelet to do a token exchange  
    - GCP/jfrog implement this, but Azure/AWS do not  
- Action Items  
  -   
- Sweep issues with leftover time  
  - [CI flakes](https://storage.googleapis.com/k8s-gubernator/triage/index.html?sig=auth)  
  - [CI testgrids](https://testgrid.k8s.io/sig-auth)

## Mar 26, 2025 11:00 AM PDT

- Cancelled due to leads availability

## Mar 12, 2025 11:00 AM PDT

- Recording  
- Announcements  
  - [SIG Auth Leadership Changes](https://groups.google.com/g/kubernetes-sig-auth/c/NzoDhrHL3VE)  
- Demos  
  -   
- Pulls of note  
  -   
- Issues of note  
  -   
- Designs of note  
  -   
- Discussion topic  
  - \[jian/deads2k\] act-as: request accepted with intersection of requesting user’s authorization and act-as user’s authorization  
    - [Original 2016 discussion (record?) with Eric Tune](https://github.com/kubernetes/kubernetes/issues/27152#issuecomment-229836331)  
    - [Discussion](https://groups.google.com/g/kubernetes-sig-auth/c/gCJ95N3wyZ0/m/0atDhG-fAAAJ?utm_medium=email&utm_source=footer)  
    - [PR showing how it works](https://github.com/qiujian16/kubernetes/pull/1/files) and [separate demo script](https://github.com/qiujian16/validationpolicy/blob/main/actas-test/test.sh)  
    - Micah: Similar in concept to AWS's permission boundaries [https://docs.aws.amazon.com/IAM/latest/UserGuide/access\_policies\_boundaries.html](https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies_boundaries.html)   
      - AWS's implementation is essentially to stuff the intersection policy into the AWS\_SESSION\_TOKEN, so at authZ time, all the information is present  
    - Mo: Azure on behalf of, service never has the permission  
      - Deads: how does the service prove it can act on behalf of user/X?  
        - Mo: authorization is constrained, the proof is part of an azure request context  
      - For Azure, https://learn.microsoft.com/en-us/entra/identity-platform/access-token-claims-reference  
      - “The application ID of the client using the token. The application can act as itself or on behalf of a user. The application ID typically represents an application object, but it can also represent a service principal object in Microsoft Entra ID.  
      - appid may be used in authorization decisions.”  
    - Mo: could a subjectaccessreview in such a client be good enough?  
    - Mo: admission and the rest of the call chain would use the controller identity, not the identity of the user.  Which one should be used?  Why or why not?  
      - What does azure do?  
        - Mo: constrained by permissions of the user and the on-behalf of.  This sounds like both?  
      - It would make sense for admission to have access to both act-as and requesting user  
    - Consider the next agenda topic.  
      - If it were generic and we could say, “controller can impersonate David only when trying to stop/start a VM”  
    - Goal of discussion: anyone have concerns about the idea?  KEP would be next step  
      - Use case: controller that can start/stop VMs but wants to make sure that the user it is acting for can also stop that VM  
  - \[mo\] inverting the node restricted service accounts check ([recording](https://www.youtube.com/watch?v=mlEfjGfDC_g&t=1465s))  
    - Service account impersonate scheduled node POC: [impersonate\_scheduled\_node](https://github.com/kubernetes/kubernetes/compare/master...enj:kubernetes:enj/f/impersonate_scheduled_node)  
    - Relates to the above act-as topic since it can be used as a generic means to constrain when impersonation can be used  
    - \[deads\] wasn’t the goal of this one to force every request from a serviceaccount to be bounded by the node it is on versus having a particular request opt-in to an intersection?  
  - \[ahmedtd\] Should we consider changing automounting service account tokens to be opt-in?  
    - Probably \>90% of workloads deployed on Kubernetes have no need to talk to kube-apiserver.  
    - The default permissions assigned to a serviceaccount don't permit much except some reflection, so workloads who actually want to talk to kube-apiserver still need to be given an appropriate RoleBinding / ClusterRoleBinding.  
      - \[ahmedtd\] existing default behavior puts load on external SA token signer that is wasteful  
      - \[liggitt\] we already have SA / pod level flags for opting out of tokens, making it easier to set via mutating admission  
      - \[ahmedtd\] will look into GKE using audit logs to recommend when pods are not using their tokens per JTI  
- Action Items  
  -   
- Sweep issues with leftover time  
  - [CI flakes](https://storage.googleapis.com/k8s-gubernator/triage/index.html?sig=auth)  
  - [CI testgrids](https://testgrid.k8s.io/sig-auth)

## Feb 26, 2025 11:00 AM PST

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
  - IETF WIMSE \- Workload Identity Practices (Arndt Schwenkschuster; [arndts.ietf@gmail.com](mailto:arndts.ietf@gmail.com))  
    - [WIMSE Workload Identity Practices - Kubernetes SIG Auth](https://docs.google.com/presentation/d/1Q7zZFuYEh9Jh4JXWBci6Uuu32uUjnArByCulJgDYgDs/edit?usp=sharing)  
    - [https://ietf-wg-wimse.github.io/draft-ietf-wimse-workload-identity-practices/draft-ietf-wimse-workload-identity-practices.html](https://ietf-wg-wimse.github.io/draft-ietf-wimse-workload-identity-practices/draft-ietf-wimse-workload-identity-practices.html)  
  - [https://github.com/kubernetes-sigs/hierarchical-namespaces](https://github.com/kubernetes-sigs/hierarchical-namespaces)  
    - Decision: archive it and start a new one (David/Mo/Jordan in agreement)  
    - Rationale: does not seem feasible to try and revive it and transfer to new maintainers  
  - [https://github.com/kubernetes/kubernetes/pull/130120](https://github.com/kubernetes/kubernetes/pull/130120)  
    - liggitt was not too concerned because it looked like a standard cluster scoped resource  
    - Neither Jordan or David had time for API review  
- Action Items  
  -   
- Sweep issues with leftover time  
  - [CI flakes](https://storage.googleapis.com/k8s-gubernator/triage/index.html?sig=auth)  
  - [CI testgrids](https://testgrid.k8s.io/sig-auth)

## Feb 12, 2025 11:00 AM PST

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
  - \[mo\] health of [https://github.com/kubernetes-sigs/hierarchical-namespaces](https://github.com/kubernetes-sigs/hierarchical-namespaces)  
    - Future plans?  
    - Is there any usage?  
      - [https://github.com/kubernetes-sigs/hierarchical-namespaces/issues/397](https://github.com/kubernetes-sigs/hierarchical-namespaces/issues/397)  
    - Should we archive it?  
      - [https://github.com/kubernetes-sigs/hierarchical-namespaces/graphs/contributors](https://github.com/kubernetes-sigs/hierarchical-namespaces/graphs/contributors)   
    - email thread notes  
      - \[rjbez17\] Adrian had to take a step back last year and I’ve been doing a poor job of maintaining it myself. I’ve been actively trying to get some more maintainers and it’s been difficult to keep people engaged. We do not have any current features planned for 2025 but do plan to get into a better cadence on patch releases. Right now we are undergoing a migration of new major versions of our dependencies (with multiple backwards incomparable changes).  
    - \[mo\] meta question: seems like we need a regular review of our sub projects?  
      - \[liggitt\] if we have folks willing to maintain it, then there is less concern about adoption, but without maintainers, the adoption numbers are more relevant  
    - \[Fleury Renaud\] interested in maintaining it  
    - \[liggitt\] how to climb the contributor ladder when a project is inactive and existing maintainers are no longer maintaining it; who in the sig with context for this project to help maintain/mentor new contributors?  
      - \[AI for SIG leads\] need to figure out current state of maintainers  
  - \[vinayakankugoyal\] Looks like NodeAuthenticator tests under test/e2e/auth aren’t running in any job. [https://kubernetes.slack.com/archives/C09QZ4DQB/p1738627076049969](https://kubernetes.slack.com/archives/C09QZ4DQB/p1738627076049969)   
    - \[AI for vinayakankugoyal\] Follow up with sig testing to ensure tests are running  
  - \[vpnachev\] Continue feature request discussion [from slack](https://kubernetes.slack.com/archives/C0EN96KUY/p1738920093068769): Enhance [audit.k8s.io/v1.Event](http://audit.k8s.io/v1.Event) with additional optional field holding cluster identifier.  
    - Questions:  
      - Is a cluster ID even known to an individual cluster?  
        - \[deads\] individual kube clusters have no concept of cluster ID.  Adding such an ID is a sig-arch decision we decided against multiple times, but that’s the venue.  
          - \[mo\] don’t we have the API server ID feature as beta?  
      - \[mo\] could use the webhook sink to edit the events as needed?  
        - \[liggitt\] you can add audit annotations during ingesting of events  
      - xref [https://github.com/kubernetes/enhancements/tree/master/keps/sig-multicluster/2149-clusterid](https://github.com/kubernetes/enhancements/tree/master/keps/sig-multicluster/2149-clusterid)  
      - does it make sense to add fields to audit events an individual cluster wouldn’t populate?  
  - \[vinayakankugoyal\] KEP-4601: Are there any plans to do labelSelector/fieldSelector-based policy in RBAC   
    - \[liggitt\] not in-scope for the current KEP  
    - \[mo\] would require something like RBAC++  
- Action Items  
  -   
- Sweep issues with leftover time  
  - [CI flakes](https://storage.googleapis.com/k8s-gubernator/triage/index.html?sig=auth)  
  - [CI testgrids](https://testgrid.k8s.io/sig-auth)

## Jan 29, 2025 11:00 AM PST

- Canceled due to empty agenda

## Dec 18, 2024 11:00 AM PST

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
  - \[[idunbarh](https://github.com/idunbarh)\] [https://github.com/kubernetes/kubernetes/issues/129075](https://github.com/kubernetes/kubernetes/issues/129075) FIPS compliant k8s discussion  
    - avoid special releases if possible (sounds like this is possible since the go behavior can be opted into at runtime)  
    - work with sig security to add ways to enable tests to provide signal for FIPs-enabled  
  - [Needs KEP / release work \#sig-auth](https://docs.google.com/document/d/1sY8fRyRtk4eG9R439z5ao5i9bFuuxilS03XaNlqoni0) for 1.33  
  - \[vinayakankugoyal\] KEP-4412 status? Did it make ALPHA. I am very interested and can help drive it.  
    - [Projected service account tokens for Kubelet image credential providers](https://kep.k8s.io/4412)  
    - \[liggitt\] was close, but didn’t make 1.32; picking up review intending to land alpha early in 1.33  
- Action Items  
  -   
- Sweep issues with leftover time  
  - [CI flakes](https://storage.googleapis.com/k8s-gubernator/triage/index.html?sig=auth)  
  - [CI testgrids](https://testgrid.k8s.io/sig-auth)

## Jan 1, 2025 11:00 AM PST

- Canceled due to holiday
