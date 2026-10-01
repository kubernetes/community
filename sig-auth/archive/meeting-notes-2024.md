# Kubernetes SIG-Auth Meeting Agenda

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
  - \[ahmedtd or liggitt or mtaufen (whoever's here)\] [https://github.com/kubernetes/kubernetes/pull/128077](https://github.com/kubernetes/kubernetes/pull/128077) was a breaking change  
    - Need to allow overriding the new audience restrictions enforced by kube-apiserver (e.g. kubelet can only ask for token audiences that are explicit in pod volumes). This breaks daemons that reuse kubelet credentials to get noderestriction protection but need to ask for custom token audiences.  
    - Discussion: [https://github.com/kubernetes/kubernetes/issues/128678](https://github.com/kubernetes/kubernetes/issues/128678)  
    - Likely need:  
      - kube-apiserver flag to specify additional audiences that are always allowed  
      - some way for users who don't control kube-apiserver flags to opt-in to additional audiences as needed  
      - \[deads2k\] why a flag vs. a secondary authz check  
        - \[liggitt\] no objection in principle, especially if dynamic audiences are anticipated  
    - \[mo/micah/anish\] follow-up: enumerate what we are trying to protect with this restriction  
  - \[ahmedtd\] Audit log annotations to show which authenticator produced the user information?  
    - see [https://issue.k8s.io/82295\#issuecomment-2489437252](https://issue.k8s.io/82295#issuecomment-2489437252)  
    - \+1 to including a first-class audit authenticator name field  
    - assigning of names to authenticators (hard-coded for in-tree ones, assigned for configured ones like oidc)  
    - Small KEP to cover the edges, especially regarding delegated authentication / logging of authenticator deciding a TokenReview (probably audit annotations)  
  - \[Bryce Palmer/everettraven\] RBAC++ related \- thoughts on a general API for authorization policies? Rough idea I had for the API as examples: [https://gist.github.com/everettraven/8add24549429b4ab1e0402980eeb8f85](https://gist.github.com/everettraven/8add24549429b4ab1e0402980eeb8f85)  
    - General idea is that it is similar to a combination of the ValidatingAdmissionPolicy \+ Ingress APIs and gives a common on-cluster API for policy engines to use. When looking into Cedar to find an example of the policy syntax, it looks like the Cedar Authorizer has a similar API.  
    - Feedback: Openness of the expressions is a bit concerning. Can’t validate expressions, etc. RBAC is popular because it is predictable, this seems like it could be unpredictable. Exploration in this area is good.  
  - \[ritazh\] [https://github.com/kubernetes/kubernetes/issues/128838](https://github.com/kubernetes/kubernetes/issues/128838)  
    - deads2k \- if I recall this one correctly, it is trying to control write authorization on a per-node basis  
      - If we even in the future want to require per-node authorization as a conformance requirement, then we need to enforce the behavior by default as we introduce it.  
      - So if we decide on that, then I think we’re at a point of either REST handler or default-on admission plugin  
      - If we decide against that, then VAP is a possibility  
    - \[ritazh\] \- based on [https://github.com/kubernetes/kubernetes/issues/128838\#issuecomment-2546167975](https://github.com/kubernetes/kubernetes/issues/128838#issuecomment-2546167975) we should recommend these special ResourceClaimTemplate and ResourceClaim resources with the "adminAccess" field to live in "admin" namespaces, and therefore, expecting namespace as the boundary to allow resource claims with admin access is sufficient. In the REST storage layer, we can validate the ResourceClaimTemplate and ResourceClaim resources are in an admin namespace with an DRAAdmin label should be sufficient.  
    - KEP should describe what normal users can do by default, how to validate, on upgrade, what do admins need to do  
  - \[ritazh\] [https://github.com/kubernetes/kubernetes/issues/127801](https://github.com/kubernetes/kubernetes/issues/127801) \- audit-id validation: size limit, printable chars  
    - size limit (\~128 chars? \~256 chars), printable chars, sanity check; because this is a bug, need a feature flag to gate it, but can be on by default because it is a bug  
    - Separate effort to enumerate where we use audit-id in-tree  
- Action Items  
  -   
- Sweep issues with leftover time  
  - [CI flakes](https://storage.googleapis.com/k8s-gubernator/triage/index.html?sig=auth)  
  - [CI testgrids](https://testgrid.k8s.io/sig-auth)

## Dec 4, 2024 11:00 AM PST

- Canceled due to leads availability

## Nov 20, 2024 11:00 AM PST

- Recording  
- Announcements  
  -   
- Demos  
  - \[micahhausler\]: Cedar Authorizer Demo: [https://github.com/awslabs/cedar-access-control-for-k8s](https://github.com/awslabs/cedar-access-control-for-k8s)  
- Pulls of note  
  -   
- Issues of note  
  -   
- Designs of note  
  -   
- Discussion topic  
  - \[ahmedtd\] Authenticating proxy header for enforcing object bindings?  "I'm authenticating as service account x, but pod y needs to exist for this request to be valid".  
    - \[deads2k\] Limits?  Could a user use this to discover things about other namespaces?  
    - \[ahmedtd/liggit\] Only the authenticating proxy user could set this.  
    - \[hoskeri\] Would a user presented this way also set the extra userinfo set by the service account token authenticator.  
    - \[ahmedtd\] Specifically pods, nodes, service accounts, (and I guess secrets)  
    - \[ahmedtd\] Let's take it back and evaluate what the actual efficiency cost of having the auth proxy evaluate the object bindings is.  
  - \[ahmedtd\] Support service account tokens bound to certificates / keys?  ([RFC 7800](https://datatracker.ietf.org/doc/html/rfc7800) / [RFC 8705](https://datatracker.ietf.org/doc/html/rfc8705))  
    - KEP to explore what it would take to authenticate these tokens.  
    - \[hoskeri\] Tried this earlier with a different [RFC (Demonstrating Proof Of Possession)](https://datatracker.ietf.org/doc/html/rfc9449)  
    - Auth proxy is a problem — there's no way for it to forward the fact "this connection was authenticated with this certificate fingerprint" to kube-apiserver.  
-   
- Action Items  
  -   
- Sweep issues with leftover time  
  - [CI flakes](https://storage.googleapis.com/k8s-gubernator/triage/index.html?sig=auth)  
  - [CI testgrids](https://testgrid.k8s.io/sig-auth)

## Nov 6, 2024 11:00 AM PST

- Canceled due to code freeze

## Oct 23, 2024 11:00 AM PDT

- Recording  
- Announcements  
  -   
- Demos  
  - \[mo/anish\] RBAC++  
    - [SIG Auth Deep Dive - KubeCon NA 2024.pptx](https://docs.google.com/presentation/d/1xwbhVtNYpZGd01G60Y8A5O_a02Oeio-Z/edit#slide=id.p4)  
    - [rbac++ PoC code](https://github.com/kubernetes/kubernetes/compare/master...enj:kubernetes:enj/f/rbacpp?w=1)  
    - [rbac++ example YAML](https://gist.github.com/aramase/f6cf33b914300741c11f121623b47647)  
- Pulls of note  
  -   
- Issues of note  
  -   
- Designs of note  
  - \[guillermo\]  [KEP-4872: Harden Kubelet serving cert validation](https://github.com/kubernetes/enhancements/pull/4911)  
- Discussion topic  
  - [\[Public\] Node Restricted Service Accounts](https://docs.google.com/document/d/1faEAFRQlRAVLK91qUjndaCkhP2GP_GmUC_po-AWcjss/edit?usp=sharing)   
- Action Items  
  -   
- Sweep issues with leftover time  
  - [CI flakes](https://storage.googleapis.com/k8s-gubernator/triage/index.html?sig=auth)  
  - [CI testgrids](https://testgrid.k8s.io/sig-auth)

## Oct 9, 2024 11:00 AM PDT

- Canceled due to KEP freeze

## Sep 25, 2024 11:00 AM PDT

- Canceled due to leads availability

## Sep 11, 2024 11:00 AM PDT

- Recording  
- Announcements  
  - Join the new [https://groups.google.com/a/kubernetes.io/g/sig-auth](https://groups.google.com/a/kubernetes.io/g/sig-auth) ML  
- Demos  
  -   
- Pulls of note  
  -   
- Issues of note  
  -   
- Designs of note  
  -   
- Discussion topic  
  - \[anish\] Metrics reset on authentication config reload as part of [Add JWKS fetch metrics for jwt authenticator](https://github.com/kubernetes/kubernetes/pull/123642)  
    - Relevant discussion: [https://github.com/kubernetes/kubernetes/pull/123642\#discussion\_r1688919218](https://github.com/kubernetes/kubernetes/pull/123642#discussion_r1688919218)  
    - \[taahir\] to share some code regarding a custom collector that would make it easier to avoid race conditions around old issuers that we do not care about  
      - Gist: [https://gist.github.com/ahmedtd/ce30522b437bda9b063f67f8750036bc](https://gist.github.com/ahmedtd/ce30522b437bda9b063f67f8750036bc)   
  - \[eminwux\] Secure proxy mechanism for private clusters’ API  
    - Best way to connect kubectl to private API using a proxy with OIDC Authentication (Proxy-Authorization header on CONNECT vs. inspecting Authorization header)  
    - \[mo\] recommendation is to either:  
      - Use a VPN/mesh solution that is transparent to the API server and kubectl to put the client and server on a private network  
      - Use an L7 proxy to terminate TLS and enforce the appropriate restrictions  
- Action Items  
  -   
- Sweep issues with leftover time  
  - [CI flakes](https://storage.googleapis.com/k8s-gubernator/triage/index.html?sig=auth)  
  - [CI testgrids](https://testgrid.k8s.io/sig-auth)

## Aug 28, 2024 11:00 AM PDT

- Recording  
- Announcements  
  - [Proposed](https://github.com/kubernetes/sig-release/pull/2604) 1.32 enhancements freeze is October 10th, code freeze is November 7th  
  - First alpha release of Secrets Store Sync Controller is out \- [Release v0.0.1 · kubernetes-sigs/secrets-store-sync-controller (github.com)](https://github.com/kubernetes-sigs/secrets-store-sync-controller/releases/tag/v0.0.1)  
- Demos  
  -   
- Pulls of note  
  -   
- Issues of note  
  -   
- Designs of note  
  -   
- Discussion topic  
  - \[guillermo\] Discuss API server to kubelet connection hardening: [Harden Kubelet Serving Certificate Validation in API server](https://docs.google.com/document/d/1Im3UO9ifBLARB4g8uIlXe2j3xPjN2DXERf4hke9gBuU/edit)  
    - Builds on top of [https://github.com/kubernetes/kubernetes/pull/126015](https://github.com/kubernetes/kubernetes/pull/126015)  
    - Would have feature gate to handle maturity level, at beta it would default to doing additional validation  
      - A CLI flag will allow explicit config, to allow for long term opt-out  
      - If the feature isn’t on, cannot set flag at all  
      - Targeting 1.32 alpha, needs a small to describe progression and timing and long term opt-out support (we believe maintenance is minimal)  
      - Could involve client-go changes to help others also perform the right kind of node name checks (i.e. metrics scraping)  
      - When the client opts out of TLS verification, this new additive verification will also be skipped  
    - Insecure backend proxy behavior  
      - [https://github.com/kubernetes/enhancements/tree/master/keps/sig-api-machinery/1295-insecure-backend-proxy](https://github.com/kubernetes/enhancements/tree/master/keps/sig-api-machinery/1295-insecure-backend-proxy)   
      - Per request option likely needs to be honored.  Those who wish to deny the behavior could use admission (VAP?)  
  - [KEP-4317 (Pod Certificates)](https://github.com/kubernetes/enhancements/pull/4318): Temperature check on external TLS?  
    - Defer this for now, just focus on giving the workload a key in the filesystem  
- Action Items  
  -   
- Sweep issues with leftover time  
  - [CI flakes](https://storage.googleapis.com/k8s-gubernator/triage/index.html?sig=auth)  
  - [CI testgrids](https://testgrid.k8s.io/sig-auth)

## Aug 14, 2024 11:00 AM PDT

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
  - 1.32 planning [Needs KEP / release work \#sig-auth](https://docs.google.com/document/d/1sY8fRyRtk4eG9R439z5ao5i9bFuuxilS03XaNlqoni0)  
- Action Items  
  -   
- Sweep issues with leftover time  
  - [CI flakes](https://storage.googleapis.com/k8s-gubernator/triage/index.html?sig=auth)  
  - [CI testgrids](https://testgrid.k8s.io/sig-auth)

## Jul 31, 2024 11:00 AM PDT

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
  - \[mo\] should the JWT authenticator set authentication.kubernetes.io/credential-id automatically if the jti claim is present?  
    - Should the user be able to override this with a CEL expression?  
      - no, and disallow setting *all* k8s.io and kubernetes.io namespaced extra info in custom ways for now  
    - I plan to add this to the structured authn KEP under the existing feature gate  
    - \[liggitt\] seems okay to use the standard JTI claim by default  
      - Should probably just warn on reserved usage namespace?  
    - \[deads\] since this is beta, we should lock down the key space that you can use  
    - \[liggitt\] we can expand what users can drive later such as node name, and should lock it down  
    - In 1.32 tighten structured authentication config file validation to not allow use of k8s.io and kubernetes.io extra info namespaces, and start defaulting to JTI being set if it is present as a non-empty string  
    - \[taahir\] x509 credential id sha256 fingerprint ready for ack from mo  
  - \[mo\] as part of the structured authn KEP, would like to expand token caching to not be a static TTL  
    - \[liggitt\] Assuming it is not too complex to implement and the benchmarks show promise, keeping it internal to the JWT authn seems okay.  We need to honor exp claim and key ID rotation  
      - No impact on any other authenticator  
      - While token review via KAS \-\> webhook might be able to take advantage of caching, external caller \-\> KAS token review would not be able to benefit from this because they would have no way to invalidate the cache  
    - Each token should have its own TTL, maybe with an upper limit on TTL and overall cache size (each cache would opt-in to the behavior, nothing would happen automatically)  
      - Cache would need to expire on certain events like signing key rotation  
    - Possible approaches:  
      - Via a new authentication.kubernetes.io/exp extra key (value is unix time in the same way that JWTs set it)  
        - Can be observed via audit logs (but do we want that?)  
        - How does this interact with token webhooks?  Presumably they can set it?  Opt-in via API server flag?  
        - Would be automatically added for JWT authenticator tokens, presumably with no way to override it  
        - Would be automatically added for client certs, but would be purely informational due to no caching  
      - Via an internal only mechanism such as:  
        - As a new field in authenticator Response  
        - As an optional interface on top of user.Info  
        - Would only be set by the JWT authenticator and would only be consumed by its cache  
  - \[vinayakankugoyal\] A while back Tim opened an issue about making the kubelet API Authz a bit more fine grained. I had a discussion about this with Jordan and wrote a proposal [https://github.com/kubernetes/enhancements/pull/4760](https://github.com/kubernetes/enhancements/pull/4760). Would love to see if we can do this in 1.32  
    - \[dead2k\] Is there an easy spot to see all the kubelet endpoints to be sure these are only ones we want to change?  
      - [https://github.com/kubernetes/kubernetes/blob/master/pkg/kubelet/server/auth.go](https://github.com/kubernetes/kubernetes/blob/master/pkg/kubelet/server/auth.go)   
    - \[mo\] can we keep the scope of this only to the node the workload is running on?  
    - \[mo\] do we need tokens to be audience scoped to the node?  What would the string be?  
  - \[danwinship\] Followup on Surya’s kubelet probes PSA discussion from last time  
    - [SIG-Auth: Small KEP for Adding PSA to block .Host fields in ProbeHandler and LifecycleHandler](https://docs.google.com/document/d/1lpVI4OuxRf7-j2jcfHm48MdHyxMR5PlKyWzAO-NUYgU)  
    - \[mo\] liggitt’s suggestion was: 1.32 PSA can have an extra check; would be a small KEP  
    - \[vinayakankugoyal\] should this be baseline?  
      - \[mo\] sound like we only want privileged pods to be able to do this  
- Action Items  
  - Everyone please review and add things to 1.32 planning [Needs KEP / release work \#sig-auth](https://docs.google.com/document/d/1sY8fRyRtk4eG9R439z5ao5i9bFuuxilS03XaNlqoni0)to make sure they are part of 1.32  
- Sweep issues with leftover time  
  - [CI flakes](https://storage.googleapis.com/k8s-gubernator/triage/index.html?sig=auth)  
  - [CI testgrids](https://testgrid.k8s.io/sig-auth)

## Jul 17, 2024 11:00 AM PDT

- Recording  
- Announcements  
  - Code freeze is next Tuesday, July 23rd, 2024  
- Demos  
  -   
- Pulls of note  
  - Authz selector PR  
    - [https://github.com/kubernetes/kubernetes/pull/125571](https://github.com/kubernetes/kubernetes/pull/125571)  
    - combined PR from David / Jordan  
    - reviewed by Mo / Joe  
    - just needs an integration test added  
- Issues of note  
  -   
- Designs of note  
  -   
- Discussion topic  
  - \[surya\] Does [Pod probes lead to blind SSRF from the node · Issue \#99425](https://github.com/kubernetes/kubernetes/issues/99425) need a new KEP??  
    - [k8s slack thread](https://kubernetes.slack.com/archives/C0EN96KUY/p1717670685761639?thread_ts=1717670527.757369&cid=C0EN96KUY)  
    - [https://github.com/kubernetes/kubernetes/pull/125271](https://github.com/kubernetes/kubernetes/pull/125271)  
    - [https://github.com/kubernetes/enhancements/pull/4558\#issuecomment-2156800239](https://github.com/kubernetes/enhancements/pull/4558#issuecomment-2156800239)  
    - \[liggitt\] Need a KEP to describe things like what happens when users upgrade etc; locking down further for restricted and baseline; can keep the scope small so we can iterate faster; can start with Google doc then open a PR for new KEP  
    - feature gates or PSA versioning?  
    - A google-doc \[mo: could be HackMD doc since KEPs are markdown\] can be the first step and then we can copy that to KEP since its faster to iterate on a google doc  
      - Concerns: API details(where are the .Host fields used from); upgrades? Compatibility?  
      - Should this be privileged?  
    - David mentioned something about feature-gates \=\> do we need one for this?  
      - We agreed “no”, but the version must be set properly in the check.  
  - \[ahmedtd\] Credential IDs for X.509 Certificates? [https://github.com/kubernetes/kubernetes/pull/125634](https://github.com/kubernetes/kubernetes/pull/125634)   
    - \[liggitt\] should ask folks that use these IDs what approach for generating the ID would be the best via slack and MLs, for SIG Auth and SIG Security  
    - Can do micro benchmark against the cost of the cert chain validation  
  - \[micahhausler\] Node Admission CSR restrictions \- [https://github.com/kubernetes/kubernetes/pull/126015](https://github.com/kubernetes/kubernetes/pull/126015)  
    - Add deprecated feature gate, on by default, to allow opt out   
    - Future work (separate PR/ potentially KEP) could also validate that API server connection to kubelet serving cert has the correct CN, not just signed by the right CA \+ DNS/IP match  
  - \[Richard Tweed\] \- Authentication or authorization for dynamic admission control, so webhooks know their requests are coming from the right Kubernetes Control plane. Possibly an auth token or mTLS. Context is [https://github.com/RichardoC/kube-audit-rest](https://github.com/RichardoC/kube-audit-rest) using these to make an audit log for managed Kubernetes clusters  
    - Can bake a credential into the URL path  
    - Can use [https://github.com/openshift/generic-admission-server](https://github.com/openshift/generic-admission-server) to workaround  
      - The server being down does break discovery / namespace deletion  
    - \[mo/liggitt\] could be a future design to use SA tokens with audiences scoped to the webhook  
- Action Items  
  -   
- Sweep issues with leftover time  
  - [CI flakes](https://storage.googleapis.com/k8s-gubernator/triage/index.html?sig=auth)  
  - [CI testgrids](https://testgrid.k8s.io/sig-auth)

## Jul 3, 2024 11:00 AM PDT

- Canceled due to leads availability

## Jun 19, 2024 11:00 AM PDT

- Canceled due to leads availability

## Jun 5, 2024 11:00 AM PDT

- Canceled due to leads availability

## May 22, 2024 11:00 AM PDT

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
  - \[mo\] any concerns with [enable kubelet server to dynamically load tls certificate files by zhangweikop · Pull Request \#124574](https://github.com/kubernetes/kubernetes/pull/124574)?  
    - [https://github.com/kubernetes/enhancements/pull/4411](https://github.com/kubernetes/enhancements/pull/4411)  
    - No concerns from SIG Auth  
  - \[mo\] what is the best path forward for DRA REST proxy? [https://github.com/kubernetes/enhancements/pull/4615\#discussion\_r1590950826](https://github.com/kubernetes/enhancements/pull/4615#discussion_r1590950826)   
    - Want to pass bidirectional data between plugins and API server without needing to update the kubelet  
    - \[deads2k\] the proxy sounds a lot like authn/z / API versioning that the API server currently does  
    - \[liggitt\] if we are going to do work in this space, we should work towards making node daemonsets safer instead of the proxy  
    - Recommended path forward: grant broad read access, use VAP to scope writes to the node  
      - This would be mentioned as a doc update  
      - Example DRA will use this approach to show best practice  
  - \[adambkaplan\] Shared Resource CSI Driver \- potential SIG-sponsored subproject [https://github.com/openshift/csi-driver-shared-resource](https://github.com/openshift/csi-driver-shared-resource)   
    - Related/relevant: ReferenceGrant proposal ([KEP-4601](https://github.com/kubernetes/enhancements/pull/4600))  
    - Lots of open questions, especially around how the driver will be granted scoped access to read secrets/config maps  
- Action Items  
  -   
- Sweep issues with leftover time  
  - [CI flakes](https://storage.googleapis.com/k8s-gubernator/triage/index.html?sig=auth)  
  - [CI testgrids](https://testgrid.k8s.io/sig-auth)

## May 8, 2024 11:00 AM PDT

- Canceled due to a light agenda.

## Apr 24, 2024 11:00 AM PDT

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
  - \[nabokihms\] Add an admission webhook ID that denied a request to audit logs  
    - [slack thread](https://kubernetes.slack.com/archives/C0EN96KUY/p1710427600461869)  
    - visibility to which webhook / admission policy denied in audit makes sense  
    - tim: we already have this for ValidatingAdmissionPolicy, [https://kubernetes.io/docs/reference/labels-annotations-taints/audit-annotations/\#validation-policy-admission-k8s-io-validation-failure](https://kubernetes.io/docs/reference/labels-annotations-taints/audit-annotations/#validation-policy-admission-k8s-io-validation-failure)  
    - David: apimachinery would be fine with recording the info from admission, annotation or field?  
      - Jordan: could apply to arbitrary requests, actual field seems reasonable  
      - Mo: would this require a KEP to add a field  
      - Jordan: small KEP to at least consider likely evolution / future admission fields to record  
  - \[mo\] [https://github.com/kubernetes/kubernetes/pull/123871](https://github.com/kubernetes/kubernetes/pull/123871)  
    - Not exactly backwards compatible so I included a temporary feature gate to aid in transition (though I am unsure this is strictly necessary since falling through to webhook would be dependent on the credential and the earlier authenticator config)  
    - Jordan: should this be done globally or per authenticator; seems useful but not sure about opt-in or opt-out  
    - Mo: feature gate with the goal to remove in the future  
    - Jordan: get feedback to see if anyone is using the same issuer for different authenticators  
  - \[mo\] [Needs KEP / release work \#sig-auth](https://docs.google.com/document/d/1sY8fRyRtk4eG9R439z5ao5i9bFuuxilS03XaNlqoni0/edit) for v1.31  
  - \[anish\] Creating a new repo for the secrets-sync-controller in kubernetes-sigs org  
    - Issue: [https://github.com/kubernetes/org/issues/4902](https://github.com/kubernetes/org/issues/4902)  
  - ~~\[vinaygo\] [Fine grained Kubelet API authorization](https://docs.google.com/document/d/1izCeuXAZ_W6uec-Uj2sLqQodOD5HBtjrhJWT9pbIWiQ/edit#heading=h.xgjl2srtytjt) revival of effort. Looks like someone from fluentbit is driving this now, we don’t need to discuss this~~  
  - \[ahmedtd\] KEP for authn proxy headers to communicate bound object requirements?  
    - conditional identity assertion via proxy  
    - today, a token-validating authentication proxy can reproduce signature verification but not bound object validation  
- Action Items  
  -   
- Sweep issues with leftover time  
  - [CI flakes](https://storage.googleapis.com/k8s-gubernator/triage/index.html?sig=auth)  
  - [CI testgrids](https://testgrid.k8s.io/sig-auth)

## Apr 10, 2024 11:00 AM PDT

- Canceled due to few people unavailable during this time.

## Mar 27, 2024 11:00 AM PDT

- Canceled due to no agenda.

## Mar 13, 2024 11:00 AM PDT

- Canceled due to a light agenda.

## Feb 28, 2024 11:00 AM PST

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
  - \[stealthybox\] Structured Authz blog so we can get more usage  
    - [slack thread](https://kubernetes.slack.com/archives/C05EZFX1Z2L/p1707941036208999?thread_ts=1707847659.589629&cid=C05EZFX1Z2L)  
    - [https://github.com/kubernetes/website/pull/45137](https://github.com/kubernetes/website/pull/45137)  
    - [https://github.com/kubernetes/website/pull/45138](https://github.com/kubernetes/website/pull/45138)  
  - \[linxiulei\] Admission Plugins bypassing REST layer for k8s objects [https://github.com/kubernetes/kubernetes/pull/121979](https://github.com/kubernetes/kubernetes/pull/121979)  
    - decoration that happens in REST storage layer that going directly to etcd would miss:  
      - [https://github.com/kubernetes/kubernetes/blob/master/staging/src/k8s.io/apiserver/pkg/registry/generic/registry/store.go\#L164-L170](https://github.com/kubernetes/kubernetes/blob/master/staging/src/k8s.io/apiserver/pkg/registry/generic/registry/store.go#L164-L170)   
      - [https://github.com/kubernetes/kubernetes/blob/e21a2f5d4f010e49cea1b954bd9b31d94e712c5b/pkg/registry/core/persistentvolumeclaim/storage/storage.go\#L68](https://github.com/kubernetes/kubernetes/blob/e21a2f5d4f010e49cea1b954bd9b31d94e712c5b/pkg/registry/core/persistentvolumeclaim/storage/storage.go#L68)  
    - are objects returned from storage shared or safe to modify?  
      - for this efficiency to work, storage would have to return shared objects  
      - any modification of those shared objects causes problems not just for in-process things looking at the same informer cache, but also things downstream, possibly exposed to API clients  
    - risk of storage-backed client behavior differing from REST-backed client behavior  
      - extensive correctness tests at every level (watcher / informer / rest client) would be required, we definitely don't have these today  
  - \[Anusha, Jim, Gus\] [KEP-4447](https://github.com/kubernetes/enhancements/pull/4448): Promote PolicyReport API to a Kubernetes SIG API  
    - Discussed [June 21st](#bookmark=id.ml6bpv5ek0gp)  
    - [PolicyReport API \- SIG API promotion proposal](https://docs.google.com/presentation/d/1IT-3vXU2XnKinglvqICCWUvCaDGlSm5V8ocJLIjqN7w/edit?usp=sharing)  
    - Offloading etcd  
      - [https://github.com/kyverno/reports-server](https://github.com/kyverno/reports-server)  
      - [https://github.com/open-cluster-management-io/enhancements/tree/main/enhancements/sig-policy/98-long-term-compliance-history](https://github.com/open-cluster-management-io/enhancements/tree/main/enhancements/sig-policy/98-long-term-compliance-history)  
    - Describe planned changes (API group, plan to get to stable, etc)  
    - Update api group to \<something\>.x-k8s.io  
    - Take a look at [https://github.com/kubernetes/enhancements/tree/master/keps/sig-auth/2907-secrets-store-csi-driver](https://github.com/kubernetes/enhancements/tree/master/keps/sig-auth/2907-secrets-store-csi-driver) as that subproject went thru the same process  
  - FYI \- ReferenceGrant API Review tomorrow  
- Action Items  
  -   
- Sweep issues with leftover time  
  - [CI flakes](https://storage.googleapis.com/k8s-gubernator/triage/index.html?sig=auth)  
  - [CI testgrids](https://testgrid.k8s.io/sig-auth)

## Feb 14, 2024 11:00 AM PST

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
  - \[mo\] support splitting of trust between OIDC discovery/new verifier and KCM SA tokens/old verifier  
    - mo: This can be partially done externally when hosting the OIDC endpoint off cluster, but that only works for external verifiers  
    - mo: This can be thought of as a step 0 to external signing of SA tokens  
      - liggitt: unclear why this split is related to external token signing  
    - mo:  
      - KCM doesn't issue OIDC-shaped tokens (no exp claim, etc), but kube-apiserver does  
      - kube-apiserver has to be given public keys that validate both kube-apiserver and kcm-issued tokens for auth to work  
      - kube-apiserver reports all the token-validating public keys it is given in the keyset published at its well-known endpoint for reliant parties to use for OIDC-style validation  
      - if distinct private keys were used to mint tokens in kube-apiserver and kcm, both public keys end up in the well-known endpoint even though the one for KCM shouldn't be used to verify tokens  
    - \[liggitt\]  
      - since the KCM-issued tokens aren't actually valid OIDC-style tokens, parties doing OIDC validation \*shouldn't\* be at risk of accepting KCM-issued tokens (they would fail validation)  
      - Okay with (someone else) doing the work to make the change as long as it is not too complicated  
      - If we do external signing of SA tokens, it should include both KCM and API server, i.e. this is not a step 0 to us doing the external signing work because we are saying that separating the trust is not sufficient because we care about in cluster usage as well  
  - \[mo\] any concerns with adding a rate limited/sometimes client side warning when a credential plugin takes more than 100 ms to respond?  
    - Maybe the warnings would be limited to interactive invocations?  
    - And probably would only be issued within the first minute or so of the process starting?  
      - \[liggitt\] trace logging at same log level (v6?) for network request timing  
      - Should not elevate to a warning, but just make it easy to see  
      - \[stealthybox\] slow cred plugins significantly affect the experience of using kubectl  
        - is there somewhere in kubectl we could opt into showing a warning  
        - \[lucas\] I think this is is one of those features that could be implemented (mostly) without touching kubectl, essentially having kubectl push metrics in a OTel-compatible format to a long-running process running besides it, from which a kubectl plugin can query "what is taking so long" and show p95 metrics, etc., as the metrics wrt this already exist in client-go, they are just not exposed anywhere as kubectl is a short-lived process   
        - is there a place in the development flow of a cred plugin that we can warn the developer of the cred pluign?  
      - mo/liggitt have philosophical difference on whose place is it to warn about kubectl being slow because a cred plugin is slow  
  - \[ahmedtd\] Discussion items for KEP-3257 promotion to beta [https://github.com/kubernetes/enhancements/pull/3913](https://github.com/kubernetes/enhancements/pull/3913)   
    - Revocation?  
      - Document that you can put CRL URIs in the root certificates in ClusterTrustBundles  
      - Non-blocking for ClusterTrustBundle feature, future work to make kube-apiserver (or other clients) safely make use of CRLs in CAs could be queued up  
    - Publishing kube-apiserver serving certificate CA root under a well-known signername?  
      - Cluster’s ca.crt  
      - Client cert CA from the CM in kube-system for apiextensions  
    - String together pod certificates and cluster trust bundles to make it easy / zero-config for kube-apiserver to connect to in-cluster admission webhook backends?  
      - mo: should produce clusterTrustBundle, but not consume yet  
      - liggitt: external consuming, but not internal  
      -   
    - Adjustments for beta  
      - reserve \*.k8s.io (at least in docs, possibly require ack on CTB objects)  
      - For each signer in our documentation, export a set of ClusterTrustBundle objects holding the roots.  
    - Does this need to have to be in the kube-apiserver or kcm?  
      - \[lucas\] \--root-ca-file is passed to both KCM and the API server, but the kubelet ca validation cert flag for API server \-\> kubelet comms is only on the API server  
  - \[stealthybox\] ReferenceGrant API Review tomorrow  
    - adding comments and small API changes to the [KEP 4387](https://github.com/kubernetes/enhancements/pull/4387) today  
      - certain fields should be lists to reduce repetition  
    - will start a draft of authorizer requirements to the KEP  
    - [https://meet.google.com/cwu-ekcs-zjd](https://meet.google.com/cwu-ekcs-zjd), 9AM PT, if anyone wants to join and listen in  
- Action Items  
  -   
- Sweep issues with leftover time  
  - [CI flakes](https://storage.googleapis.com/k8s-gubernator/triage/index.html?sig=auth)  
  - [CI testgrids](https://testgrid.k8s.io/sig-auth)

## Jan 31, 2024 11:00 AM PST

- Recording  
- Announcements  
  -   
- Demos  
  - Kevin Conner \- Demo of CEL Playground  
    We would like to give a demo of [CEL Playground](https://playcel.undistro.io/), an open source playground for testing out CEL expressions.  We are looking for feedback on SIG AUTH use cases for CEL, our intention is to extend what is currently available and provide support for these features alongside other k8s use cases.  We gave a presentation to SIG API Machinery back in October and [received valuable feedback](https://github.com/undistro/cel-playground/issues/41) from Jordan, we are looking for something similar from SIG AUTH.  
    - Q\&A:  
      - optional values support: [https://github.com/google/cel-spec/wiki/proposal-246](https://github.com/google/cel-spec/wiki/proposal-246)  
        - Needs k8s libs v1.29+  
      - Future support for reference grant  
- Pulls of note  
  -   
- Issues of note  
  -   
- Designs of note  
  -   
- Discussion topic  
  - \[robscott\] [Referential Authorization KEP](https://github.com/kubernetes/enhancements/pull/4387)  
    - \[stealthybox\] new API shape in the KEP \-- please review  
    - Lucas Käldström (luxas on Github and Slack) Relation-based Access Control using Zanzibar, [https://github.com/luxas/kube-rebac-authorizer](https://github.com/luxas/kube-rebac-authorizer)  
    - deads: specific details on how the authorizer will be built (RBAC vs in-tree etc) \-\> implementation details will be sticky as folks will depend on them (i.e. can directly observe RBAC on the cluster)  
      - mo/david lean towards have an authorizer like the node authorizer, i.e. no RBAC based implementation in tree  
        - deads: errors? status? warnings?  great opportunities here that were not possible back in RBAC   
        - deads: reserve a slice of resource names (suffix/prefix)

        provide some recommendations for names (maybe validate )

        - Authorizers are optional, so we have to come up with how to handle this not being turned on (maybe warn users but let the API be active?) maybe don’t expose the API? Exposing RBAC was potentially a mistake (manifest applies even though it’s not applied)  
        - Is there a path for adoption of this for core things?  Like pods, controllers, etc  
        - bootstrap object for Ingress API, VolumeSnapshot  
        - Self rules review will not work for this API 🙁  
        - When are we trying to get this code in?  v1.30?  
  - \[anish\] audiences in Structured AuthN Config  
    - [https://kubernetes.slack.com/archives/C04UMAUC4UA/p1705511911467009](https://kubernetes.slack.com/archives/C04UMAUC4UA/p1705511911467009)  
    - [https://github.com/kubernetes/enhancements/pull/4430](https://github.com/kubernetes/enhancements/pull/4430)   
  - \[cici\] Explicitly exclude resources from ValidatingAdmissionPolicy([issue](https://github.com/kubernetes/kubernetes/issues/122205))  
    - Admission should not be able to break read requests (i.e. token review)  
    - Probably should not handle virtual resources since they are not persisted  
    - Add integration test that makes sure that virtual resources cannot be intercepted  
- Action Items  
  -   
- Sweep issues with leftover time  
  - [CI flakes](https://storage.googleapis.com/k8s-gubernator/triage/index.html?sig=auth)  
  - [CI testgrids](https://testgrid.k8s.io/sig-auth)

## Jan 17, 2024 11:00 AM PST

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
  - \[liggitt\] initial 1.30 plans  
    - [Needs KEP / release work \#sig-auth](https://docs.google.com/document/d/1sY8fRyRtk4eG9R439z5ao5i9bFuuxilS03XaNlqoni0/edit)  
  - \[ritazh\] moving to beta for structured authn and authz enhancements. Beta criteria:  
    - Address user reviews and iterate (if needed, keep in Alpha until changes stabilize)  
    - Feature flag will be turned on by default  
      - [https://kubernetes.slack.com/archives/C0EN96KUY/p1705510206573939](https://kubernetes.slack.com/archives/C0EN96KUY/p1705510206573939)  
      - \[deads2k\] not significantly easier to make changes to file vs REST API because you still need to be able to convert to support upgrade   
  - \[mo\] [authn: consider expanding oidc API to allow warnings to be returned to the user · Issue \#119834](https://github.com/kubernetes/kubernetes/issues/119834)  
    - \[mo\] ai to summarize discussion on the issue  
  - \[mo\] thoughts on [Projected Service Account Tokens: Ability to specify multiple Audiences · Issue \#122500](https://github.com/kubernetes/kubernetes/issues/122500)  
    - \[ai for mike\] to update the issue with thoughts and close it out  
  - \[mo\] what should we do about [RoleBindings deleted by namespace finalizer before other objects using those permissions finish stopping/finalizing · Issue \#115070](https://github.com/kubernetes/kubernetes/issues/115070)  
    - \[deads2k\] ai to summarize thoughts around leans and close issue  
  - \[anish (*Adding this from sig-auth issue triage meeting)*\] thoughts on [RBAC return 403 Forbidden after restarting APIServer when there is a large amount of rolebindings](https://github.com/kubernetes/kubernetes/issues/121635)  
    - Should the default behavior of informers at startup be changed?  
    - Should we improve the performance of RBAC authorizer?  
    - [https://kubernetes.slack.com/archives/C04UMAUC4UA/p1704742234899109](https://kubernetes.slack.com/archives/C04UMAUC4UA/p1704742234899109)  
      - \[anish\] to summarize and close the issue  
  - \[kannon92 Kevin Hannon\] Bring RotateKubeletServerCertificate to GA  
    - Taking over [https://github.com/kubernetes/enhancements/pull/3806](https://github.com/kubernetes/enhancements/pull/3806)  
    - Would like feedback on what is necessary to get a retrospective KEP so we can start GAing this feature.  
- Action Items  
  -   
- Sweep issues with leftover time  
  - [CI flakes](https://storage.googleapis.com/k8s-gubernator/triage/index.html?sig=auth)  
  - [CI testgrids](https://testgrid.k8s.io/sig-auth)

## Jan 3rd, 2024, 11a \- Noon (Pacific Time)

- Recording  
- Announcements  
  - \[liggitt\] 1.30 enhancements freeze [proposed](https://github.com/kubernetes/sig-release/pull/2403) for Feb 8th  
- Demos  
  -   
- Pulls of note  
  -   
- Issues of note  
  -   
- Designs of note  
  - ReferenceGrant updates, POC, looking for feedback/help  
    - [https://github.com/robscott/referencegrant-poc](https://github.com/robscott/referencegrant-poc)  
    - [\[SIG-Auth\] ReferenceGrant Proposal](https://docs.google.com/document/d/1poQb0uxOkJsebNgTMrpaogcY9vcehGHe1myqvenCXtU/edit)  
- Discussion topic  
  - [Fix NCC-E003660-R44: Authentication Source Not Shown in Audit Logs](https://github.com/kubernetes/kubernetes/pull/119644)  
    - was discussed a bit as part of [https://github.com/kubernetes/kubernetes/pull/118571](https://github.com/kubernetes/kubernetes/pull/118571) but lower in priority and needs someone to drive  
    - \[liggitt\] needs a KEP, will summarize on issue [https://github.com/kubernetes/kubernetes/issues/119626](https://github.com/kubernetes/kubernetes/issues/119626)  and collapse existing WIP PRs  
  - [NCC-E003660-47W: Loopback Token Usable Externally](https://github.com/kubernetes/kubernetes/issues/119628)  
    - seems low priority given there's no exploit path that doesn't also give access to other secrets that allow persistent access (like service account signing key)  
    - if someone wanted to pick up modifying the superuser token authenticator to be a Request authenticator that would have access to the request IP and limit it to loopback that could be fine, but not a top priority (@liggitt perspective)  
  - [NCC-E003660-F9W: Common Certificate Authority Possible for Client CA and Request](https://github.com/kubernetes/kubernetes/issues/119267)  
    - Overlapping client-ca and requestheader-client-ca-file without requestheader-allowed-names is a very problematic configuration. Preventing specifically that configuration by requiring –requestheader-allowed-names when CA bundles overlap seems reasonable.  
  - [Ability to validate kubectl exec/attach requests in the Admission Controller based on requestor source IP address](https://github.com/kubernetes/kubernetes/issues/121014)  
    - Sounds like to [https://github.com/kubernetes/enhancements/pull/2843/files](https://github.com/kubernetes/enhancements/pull/2843/files), last discussed almost exactly [one year ago](#january-4th,-2023,-11a---noon-\(pacific-time\)) (Jan 4, 2023\)  
    - Discussion points last time:  
      - reliability of source IP across all deployments  
      - impact on additional metadata on authn / authz caches  
      - opening the door to a parade of requests for "just one more request attribute" (arbitrary headers, tls level, authentication method, protocol version, proxy IP, etc, etc)  
  - \[ahmedtd\] Hardcoded 10-second timeout for webhook token authentication?  How can we make this more configurable?  And is silently treating the request as unauthenticated the correct behavior?  
    - [https://github.com/kubernetes/kubernetes/blob/0c645922edcc06adff43c70c02fb56751364bbb5/staging/src/k8s.io/apiserver/plugin/pkg/authenticator/token/webhook/webhook.go\#L154](https://github.com/kubernetes/kubernetes/blob/0c645922edcc06adff43c70c02fb56751364bbb5/staging/src/k8s.io/apiserver/plugin/pkg/authenticator/token/webhook/webhook.go#L154)  
    - Related: Our problem is caused by users with thousands of groups; has there been any thought about different ways of querying group membership during authn/authz?  
    - timeout can be added to the structured authn KEP  
  - \[ritazh\] kms v1 removal timeline  
    - \[liggitt\] not in favor of removal and stop write for v1, maintenance shouldn’t be that bad  
    - \[Min Ni\] pain points between migration vs v1 gaps  
    - \[deads2k\] if storage migrator is better, should removal be considered; less code to maintain the better; in favor of stop write and removal  
    - \[AI/ritazh\] update KEP to describe the motivation for why we want to stop write and removal; remove hard dates for those and requirement for stop write and removal e.g. better storage migrator  
- Action Items  
  -   
- Sweep issues with leftover time  
  - [CI flakes](https://storage.googleapis.com/k8s-gubernator/triage/index.html?sig=auth)  
  - [CI testgrids](https://testgrid.k8s.io/sig-auth)
