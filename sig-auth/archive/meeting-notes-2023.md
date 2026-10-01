# Kubernetes SIG-Auth Meeting Agenda

## Dec 20th, 11a \- Noon (Pacific Time)

- Canceled due to winter holidays.

## Dec 6th, 11a \- Noon (Pacific Time)

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
  - \[gcastle\] Can we improve the safety of anonymous auth without breaking most of the existing use cases? At kubecon we’re giving a talk about how users are [accidentally binding cluster-admin to system:anonymous and getting compromised](https://sched.co/1R2tp). On GKE we’ve blocked that specific configuration but turning off anonymous auth is breaking and it’s been the default for a long time. The dependencies we found in our research were [health checks](https://github.com/kubernetes/kubernetes/issues/43784), [kubeadm](https://github.com/kubernetes/design-proposals-archive/blob/main/cluster-lifecycle/bootstrap-discovery.md#implementation-details), [rancher](https://github.com/rancher/rancher/blob/68215bcdde090854ff28b43b01d1aa0b69611920/pkg/data/dashboard/rbac.go), PSP bindings (EOL anyway), and CI/CD metrics. Should [\--anonymous-auth](https://github.com/rancher/rancher/blob/68215bcdde090854ff28b43b01d1aa0b69611920/pkg/data/dashboard/rbac.go) be \[true, false, **status-only**\]? For status-only we’d block all bindings except the default system:public-info-viewer.  
    - \[liggitt\] health checks and metrics are more common.  
      - Read only, targeted scoping  
    - \[gcastle\] bootstrapping cases, volume is low.  
    - \[mo\] AKS runs with anonymous auth off by default, why not do that?  
    - \[liggitt\] specific example is kubelet health checking api server.  
    - \[micah\] secondary healthcheck port?  
    - \[mo\] using structured authentication config for input?  
    - \[liggitt\] let users choose which paths they want to allow for anonymous auth  
    - \[mtaufen\] sounds more like authz config  
    - \[mo\] could ship a default validating policy binding?  
    - \[liggitt\] Provide a way to configure a more scoped default instead of the current default.  
      - Piggybacking on authn or authz config.  
      - \[mo\] Small new KEP  
    - \[gcastle\] AI: followup on who can help with a small KEP. Options discussed were:  
      - 1\. Change auth config to allow access to the request context then make a decision at auth time to only allow you to be system:anonymous if your request is targeting the allow-list of monitoring/status endpoint URLs. Follow similar pattern to how OIDC is allowed (new part is URL plumbing).  
      - 2\. Build in a CEL authorization policy that’s enabled by default that blocks bindings to system:anonymous unless they are the allow-list of public roles. This has the disadvantage of being deeper in the stack for DOS purposes, but perhaps feels more natural a home for authorization policy.  
      - 3\. Build the status-only option described above. I think we ruled this out in favor of auth config. Problem was that we’d have a constant debate and friction about what is or isn’t included in the “status-only” set. At least if it’s stored in configurable authentication people can modify it to match their needs.  
  - \[ahmedtd\] KEP-4317: PodIdentity certificates, PodCertificate volumes, and in-cluster kube-apiserver client certificates  
    - [KEP Document](https://github.com/kubernetes/enhancements/pull/4318)  
    - [Draft Implementation](https://github.com/kubernetes/kubernetes/pull/121596)  
      - \[taahir\] TPM backed keys are out of scope  
      - \[mo\] what about envs that do not support API server client certs?  
      - \[taahir/mo\] is the custom extension critical?  Could be critical to avoid confusion.  
      - \[mo/micah\] revocation?  
        - Could work for 1p use case with api server when pod is deleted  
      - \[micah\] duration of certs?  
      - \[micah\] blocks pod creation, needs to wait for the signer.  
        - \[mo\] needs to be HA  
      - \[taahir\] plan is to get KEP and alpha implementation merged in v1.30  
        - \[mo\] API review in Jan 2024  
  - \[mtaufen\] should we consider an option where K8s refuses to either admit or schedule Pods with nodeselectors that rely on non-node-restricted labels (and consider deprecating the ability to rely on non-node-restricted labels for scheduling over some time period)? This is a bad practice because nodes can self-edit labels without the noderestriction prefix. Users *can* write their own policy controllers for this already (probably easy to do in CEL even) but would be great to eliminate this risk from kube altogether.  
    - Node restriction labels are a deny list, users do not use labels that fit within the deny list, malicious node can change its labels to steer workloads to itself  
    - \[tallclair\] lots of non security reasons to use labels such as nodes self discovering hardware  
    - \[liggitt\] the author of the scheduler policy may not even know that they have different pools of isolated nodes  
    - \[mtaufen\] would like to offer something to paranoid folks that want this  
    - \[liggitt\] might be good to ask sig security for scanning tools that could detect lack of even a single node restricted label (so that it can be in checklists that people can use to make sure they are following best practice)  
    - \[liggitt\] no lever today to limit kubelet permissions via api server if admin does not want to use self registration  
- Action Items  
  -   
- Sweep issues with leftover time  
  - [CI flakes](https://storage.googleapis.com/k8s-gubernator/triage/index.html?sig=auth)  
  - [CI testgrids](https://testgrid.k8s.io/sig-auth)

## Nov 22nd, 11a \- Noon (Pacific Time)

- Canceled due to Thanksgiving.

## Nov 8th, 11a \- Noon (Pacific Time)

- Canceled due to KubeCon NA 2023\.

## Oct 25th, 11a \- Noon (Pacific Time)

- Canceled to give more time for reviews before code freeze.

## Oct 11th, 11a \- Noon (Pacific Time)

- Recording  
- Announcements  
  - \[mikedanese\]: TL role  
    - stepping down from TL role, will continue as chair  
    - nominating mo (@enj) as TL replacement, has been involved in overseeing / leading technical work already (supported by other TLs, deads2k and liggitt)  
    - will email sig-auth / dev@kubernetes.io and start a couple week lazy consensus  
  - \[liggitt\] 15 working days (including today) until code freeze (10/31)  
- Pulls of note  
  - \[mo\] [https://github.com/kubernetes/kubernetes/pull/121120](https://github.com/kubernetes/kubernetes/pull/121120)  
    - deads2k: separate gates for the anonymous/failed-authn mitigation and the heuristic detecting early client close mitigation  
    - mikedanese: implications of closing connections for clusters with a L7 load balancer in front (where the connection from the load balancer could carry a mix of anonymous and authenticated requests), especially if the load balancer was already mitigating the http/2 issue?  
  - \[vinayak\] [https://github.com/kubernetes/kubernetes/pull/120780](https://github.com/kubernetes/kubernetes/pull/120780)   
- Issues of note  
  -   
- Designs of note  
  -   
- Discussion topic  
  - \[mo/aramase\] SS CSI controller split discussion continued  
    - [Secrets Store CSI Driver Sync Secrets](https://docs.google.com/document/d/1Ylwpg-YXNw6kC9-kdHNYD3ZKskj9TTIopwIxz5VUOW4)  
      - "key" in SecretSync object is the map key, not a confidential encryption key  
      - liggitt: would be helpful to highlight the delta between the built-in permissions and the controller permissions  
      - mo: also highlight the controller wouldn't run in the daemonset that runs on all nodes like the current driver  
      - mo: goal is to not give the controller any secret read permissions, just SSA  
        - liggitt: SSA of name/namespace-only patch can succeed and returns existing content  
        - deads2k: no list/watch means periodic blind writes are required, cost concerns  
    - Make sense to be a separate project to break this functionality out  
      - AI: api review with liggitt in the next few weeks  
      -   
  - \[mo\] thoughts on expanding [Kubelet Credential Providers](https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/2133-kubelet-credential-providers)  
    - currently it only supports giving the plugin the image as input  
      - we could enhance it and maybe the pod API to support sending bound SA tokens through the CredentialProviderRequest  
        - prior art for CredentialProvider indicating a desire for an audience-scoped token: CSIDriver  
        - liggitt: consider impact on kubelet scale for additionally requested tokens  
          - deads2k: only happens once at pod init, doesn't need refreshing during pod lifetime  
        - liggitt: if we make use of per-serviceaccount credentials in image pulls implicit / invisible in kubelet CredentialProvider config, having something like node kep 2535 (linked below) in place would avoid making the cross-pod already-pulled image use issue worse  
      - Will require non-trivial changes to the caching logic used by the kubelet  
    - original issue: [Support an image pull credential flow built on bound service account tokens · Issue \#68810](https://github.com/kubernetes/kubernetes/issues/68810)  
    - we deferred this to a later time in the original KEP [https://github.com/kubernetes/enhancements/pull/1406/files\#r371886703](https://github.com/kubernetes/enhancements/pull/1406/files#r371886703)  
    - Depends on [https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/2535-ensure-secret-pulled-images](https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/2535-ensure-secret-pulled-images) to make sure disk cache does not skip authz checks  
      - Or we could require image pull to be set to always  
      - 2535 is targeting alpha in current/next release  
    - Are credential providers configurable in cloud envs?  i.e. can I use a registry from vendor A on a kubelet run by vendor B?  
      - liggitt: not always  
    - Could we isolate this change to kubelet \+ credential providers?  i.e. no change to the Kube REST API?  
- Action Items  
  -   
- Sweep issues with leftover time  
  - [CI flakes](https://storage.googleapis.com/k8s-gubernator/triage/index.html?sig=auth)  
  - [CI testgrids](https://testgrid.k8s.io/sig-auth)

## Sept 27th, 11a \- Noon (Pacific Time)

- Canceled due to multiple leads being out.

## Sept 13th, 11a \- Noon (Pacific Time)

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
  - \[mo/munnerz\] bound SA token improvements KEP  
    - [https://github.com/kubernetes/enhancements/pull/4141](https://github.com/kubernetes/enhancements/pull/4141)  
    - Do we want to include more information beyond immediate requester?  
      - JTI can be used for deeper cross-referencing anyway  
    - Do we want to include non-’system:’ prefixed users?  
      - Potential privacy concerns including username of requester in issued tokens  
    - Would a system: prefix only implementation initially be preferable?  
    - Discussion:  
      - Adding \`node\` field as a first class field like serviceaccount, pod  
        - Similar to serviceaccount, pod, we should also validate that the referenced node still exists  
        - This would be a breaking change, as upon deleting the Node object, all tokens issued by it would become invalid.  
          - Useful for revocation  
      - Not including requestingUser at all to avoid privacy concerns and allow us to properly assess how to do this securely  
- Action Items  
  -   
- Sweep issues with leftover time  
  - [CI flakes](https://storage.googleapis.com/k8s-gubernator/triage/index.html?sig=auth)  
  - [CI testgrids](https://testgrid.k8s.io/sig-auth)

## August 30th, 11a \- Noon (Pacific Time)

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
  - \[aramase/mo\] SS CSI sync as secrets feature split \+ offline support  
    - [Secrets Store CSI Driver Caching](https://docs.google.com/document/d/1zje84bP5bXoJJrn6GwpgLAeAhUoAjnb5zRE5Jzwwv5M)  
    - offline support:  
      - \[liggit\] make it more clear that this is for new pods, not for existing pods as existing pods get a bit of coasting support so the cloud vault does not become a critical path that can impact availability  
      - \[mo\] this should be an opt-in feature as it would skip many provider authentication checks  
      - \[liggit\] can another workload using a different service account get access to the cached data? no  
      - \[mo\] durable storage is a challenge. CRD is better than k8s secrets to reduce access from over privileged components  
      - \[liggit\] how can we make the encryption key less accessible such that the worst case is you lose the cache. less durable.   
      - \[mtaufen\] why not cache the key on the node? need to survive node restart. how about have the provider run on an isolated node and the cache can move to diff nodes  
      - \[taahir\] nodeauthorizer?  
      - \[mtaufen\] we should try our best to preserve node isolation  
      - \[liggit\] are the users that need this feature the same users that use sync secrets? no  
      - \[taahir\] what if every node has TPM? not guaranteed  
      - \[mo\] would PVs be better than CRDs? rpc making the request against the centralized cache; \[mtaufen\] not sure if this is better  
    - split \- csi mount and secrets sync  
      - secrets sync is used with env vars, tls secrets, ingress  
      - external secrets operators does the same thing  
      - split the sync out of csi  
      - after the split, would both projects be sig auth sub projects?  
        - \[liggitt\] we need to sponsor it since the feature is already part of sig auth  
        - for existing users, need to deprecate in the csi driver and document migration  
        - sample to limit the token requests [https://github.com/aramase/secrets-store-controller/blob/main/config/samples/secrets-store.csi.x-k8s.io\_v1\_secretprovider.yaml](https://github.com/aramase/secrets-store-controller/blob/main/config/samples/secrets-store.csi.x-k8s.io_v1_secretprovider.yaml)  
      - single controller has God level permissions  
        - only if it needs access to the service account it needs to sync; hard to know ahead of time  
        - \[taahir\] this can be limited with a validating webhook  
        - \[mo\] ship with a CEL vap policy? but cannot check what is not in the request  
        - \[mo\] disallow token request and the type of secrets the controller can make  
      - deployment needs to be revisited for providers and this new k8s-sigs controller  
- Action Items  
  -   
- Sweep issues with leftover time  
  - [CI flakes](https://storage.googleapis.com/k8s-gubernator/triage/index.html?sig=auth)  
  - [CI testgrids](https://testgrid.k8s.io/sig-auth)

## August 16th, 11a \- Noon (Pacific Time)

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
  - \[mo/rita\]​ ​[Needs KEP / release work \#sig-auth](https://docs.google.com/document/d/1sY8fRyRtk4eG9R439z5ao5i9bFuuxilS03XaNlqoni0/edit)  
    - Planning for v1.29 and KEP backlog grooming for the next few releases  
  - \[mo\] bound SA token enhancements  
    - \[mike\] what about request chaining (nested user info requester)  
    - POC: [https://github.com/kubernetes/kubernetes/pull/119739](https://github.com/kubernetes/kubernetes/pull/119739)  
    - Very rough KEP from munnerz [https://github.com/kubernetes/enhancements/pull/4141](https://github.com/kubernetes/enhancements/pull/4141)  
    - jti claim set to UUID and observable via token request audit annotation and user info extra field  
      - \[liggitt\] \+1 on jti claim, TokenRequest audit log annotation  
    - Automatically include username and uid of requesting user when token request is performed  
      - Should username and UID have some max length to be included, such as \~256 each?  
      - Included in all user info afterwards (so audit logs, authz, etc can see it)? maybe not (or at least not now)  
      - no groups due to token payload size concerns  
      - no extra due to token payload size concerns and to avoid nesting concerns around multiple chained token requests  
    - \[feedback from liggitt to mo\]  
      - Could include the first and last identity in chain to get the bulk of the information we need without worrying about unbounded lists  
      - Requester information is good to include in the audit log  
      - Not convinced it should be visible in user info  
        - something beyond the authn layer rejecting a request from a service account because of who requested the service account credential seems like breaking layers  
        - for comparison, we don't let layers past impersonation make decisions based on the original user, or impersonation would not work consistently  
      - An actor with direct access to the JWT payload can make policy decisions based on all of its content, including the request (\[mo\] unclear to me exactly how this is different from the user info case above, so we should describe our reasoning clearly)  
      - We should consider adding a proper node ref field to the payload, in addition to the request info  
        - \[mo\] would this show up in user info like pod info (I believe the answer is yes since pod info is most commonly used to look up the node)?  Also, why does secret ref info not show up in user info?  Should it?  
        - \[liggitt\] \+1 to make node name/uid show up in user info just like pod  
      - open question about whether including non-system usernames in a claim (visible to token consumers) would be a PII concern  
        - could survey providers we know about  
        - if it is a concern, could make non-system requester inclusion opt-in  
- Action Items  
  -   
- Sweep issues with leftover time  
  - [CI flakes](https://storage.googleapis.com/k8s-gubernator/triage/index.html?sig=auth)  
  - [CI testgrids](https://testgrid.k8s.io/sig-auth)

## August 2nd, 11a \- Noon (Pacific Time)

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
  - \[chrismuellner\] Proposed changes to allow bidirectional mount propagation without privileged container  
    - How powerful is SYS\_ADMIN capability without privileged today?  
      - Answer: Very\!  
    - Would we be giving substantial host/root capabilities to pods with bidirectional mount propagation?  
    - [https://github.com/kubernetes/kubernetes/pull/117812](https://github.com/kubernetes/kubernetes/pull/117812)  
      - specific comment by \[msau42\] asking to discuss these questions: [https://github.com/kubernetes/kubernetes/pull/117812\#discussion\_r1244616195](https://github.com/kubernetes/kubernetes/pull/117812#discussion_r1244616195)   
        - \[mo\] General answer seems to be that this would still be blocked by PSA baseline so this should be fine  
  - \[liggitt\] Security audit architectural issues  
    - [NCC-E003660-PA6: Additive Access Controls · Issue \#118982](https://github.com/kubernetes/kubernetes/issues/118982)  
      - \[deads\] concerns around latent reads for deny authz  
      - \[mo\] concerns around deny ordering, especially for rbac  
        - \[deads\] can a namespace admin deny the cluster admin from deleting pods in their namespace?  
      - \[liggitt\] deny precedence control at the cluster operator level seems okay and useful for tightly controlling how things work and the ordering  
    - [NCC-E003660-DXX: Lack of Cohesion Between Core Access Control Mechanisms · Issue \#118985](https://github.com/kubernetes/kubernetes/issues/118985)  
- Action Items  
  -   
- Sweep issues with leftover time  
  - [CI flakes](https://storage.googleapis.com/k8s-gubernator/triage/index.html?sig=auth)  
  - [CI testgrids](https://testgrid.k8s.io/sig-auth)

## July 19th, 11a \- Noon (Pacific Time)

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
  - \[Carl Braganza/Prasad Ghangal/Ivan Sim\] SIG-storage data protection working group KEP \- Changed Block Tracking:  
    - [https://github.com/kubernetes/enhancements/pull/4082](https://github.com/kubernetes/enhancements/pull/4082)   
    - We will like to provide a high-level intro to this KEP and get some feedback on the proposed security token mechanism  
    - feedback:  
      - Do not create your own authn/z stack, leverage the existing APIs such as token request / token review / subject access review  
      - avoid putting secret data in object names (visible to anyone with API read permission, used in etcd storage path so is not protected by encryption of etcd object content)  
      - avoid persisting credentials in API objects if possible (visible to anyone with API read permission)  
        - If something must be stored, store a hash  
      - backup client  
        - get an audience-scoped serviceaccount token for the backup client's serviceaccount (time-bound credential identifying a serviceaccount to a particular audience/service)  
        - add token volume to backup client pod spec (tokens get managed by the kubelet) or make TokenRequest API call directly  
          - example of requesting an audience-bound token in a pod spec:  
            - [https://kubernetes.io/docs/tasks/configure-pod-container/configure-service-account/\#launch-a-pod-using-service-account-token-projection](https://kubernetes.io/docs/tasks/configure-pod-container/configure-service-account/#launch-a-pod-using-service-account-token-projection)   
          - example of TokenRequest API  
            - [https://kubernetes.io/docs/reference/kubernetes-api/authentication-resources/token-request-v1/](https://kubernetes.io/docs/reference/kubernetes-api/authentication-resources/token-request-v1/)  
            - [https://github.com/kubernetes/kubernetes/blob/release-1.27/staging/src/k8s.io/client-go/kubernetes/typed/core/v1/serviceaccount.go/\#L54](https://github.com/kubernetes/kubernetes/blob/release-1.27/staging/src/k8s.io/client-go/kubernetes/typed/core/v1/serviceaccount.go/#L54)  
          - this is the field to inject token via projected volume with a specific audience: [https://github.com/kubernetes/kubernetes/blob/master/staging/src/k8s.io/api/core/v1/types.go\#L1746](https://github.com/kubernetes/kubernetes/blob/master/staging/src/k8s.io/api/core/v1/types.go#L1746)  
      - backup service / sidecar  
        - Make a TokenReview API call using the backup service's serviceaccount to check the presented token for the specific CSI driver's audience:  
          - [https://github.com/kubernetes/kubernetes/blob/master/staging/src/k8s.io/api/authentication/v1/types.go\#L52](https://github.com/kubernetes/kubernetes/blob/master/staging/src/k8s.io/api/authentication/v1/types.go#L52)  
          - https://kubernetes.io/docs/reference/kubernetes-api/authentication-resources/token-review-v1/  
        - Take the identity returned from the TokenReview request, then make one or more SubjectAccessReview API calls using the backup service's serviceaccount to authorize that identity for the snapshots being requested  
          - [https://kubernetes.io/docs/reference/kubernetes-api/authorization-resources/subject-access-review-v1/](https://kubernetes.io/docs/reference/kubernetes-api/authorization-resources/subject-access-review-v1/)  
- Action Items  
  -   
- Sweep issues with leftover time  
  - [CI flakes](https://storage.googleapis.com/k8s-gubernator/triage/index.html?sig=auth)  
  - [CI testgrids](https://testgrid.k8s.io/sig-auth)

## July 5th, 11a \- Noon (Pacific Time)

- Canceled due to no agenda

## June 21st, 11a \- Noon (Pacific Time)

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
  - \[mo\] [kms v2 performance numbers](https://gist.github.com/enj/4e819ffaa28b1da1770de449a55b24e3)  
    - Next step: look into better buffer re-use for proto decode  
    - Overall, no concerns over the performance numbers  
  - \[mo\] kms v2 should we write the encrypted seed to disk?  
    - Should be safe since info is always random (no nonce reuse concerns)  
      - RH crypto agrees it’s safe.  
    - Could be optional since requires some storage  
    - Would allow us to retain the encrypted seed until the plugin provides a new key ID (prevent fragmentation of seeds, restarts remain fast indefinitely)  
    - We would probably use the Encrypted Object proto as the schema  
      - Next step: probably under its own feature gate (in the same or new KEP)  
      - Look into storing this in etcd (in the same way as master leases or similar) so that the disk used for storage is the same disk as etcd backups  
      - Technically the idea sounds “fine / implementable”  
  - \[Jim, Jaya \- Policy WG\] [PolicyReport API \- SIG API promotion proposal](https://docs.google.com/presentation/d/1IT-3vXU2XnKinglvqICCWUvCaDGlSm5V8ocJLIjqN7w/edit?usp=sharing)  
  - \[Micah\] requester information in ServiceAccount claims  
    - Two models:  
      - 1\. Node binding, like pod binding or secret binding in current service account token implementation. If node is deleted, then token is invalidated.  
      - 2\. Credential\_origin: list of authenticated users that we append some user info to.  
- Action Items  
  -   
- Sweep issues with leftover time  
  - [CI flakes](https://storage.googleapis.com/k8s-gubernator/triage/index.html?sig=auth)  
  - [CI testgrids](https://testgrid.k8s.io/sig-auth)

## June 7th \- CANCELED

Canceled this occurrence due to a light agenda and multiple leads have conflicts.  Reminders:

- v1.28 PRR freeze \- Thursday 8th June 2023  
  - v1.28 enhancements freeze \- 18:00 PDT Thursday 15th June 2023

## May 24th, 11a \- Noon (Pacific Time)

- Recording  
- Announcements  
  - \[mo\] declare KMS v1beta1 as deprecated in v1.28  
    - The feature is not going away, but will only have security bug fixes going forward  
    - Once KMS v2 goes GA (\~v1.29), we will require a deprecated feature gate to use the functionality  
- Demos  
  - \[mo\] cert based signing for SA tokens, [diff](https://github.com/kubernetes/kubernetes/compare/master...enj:kubernetes:enj/f/sa_cert_signing?w=1)  
    - Mo to investigate how common x5c support is in libs  
      - Related: [1393-oidc-discovery](https://github.com/kubernetes/enhancements/tree/master/keps/sig-auth/1393-oidc-discovery#graduation-criteria) GA criteria  
        - The feature has been confirmed compatible with several independent relying parties by federating K8s identities with multiple top cloud providers and ensuring that the most popular OIDC libraries used by relying parties are compatible.  
    - Mo to look into what key ID (kid) semantics apply with x5c  
      - \[mike\] kid MUST NOT be present if jwk or x5c is present (??)  
- Pulls of note  
  -   
- Issues of note  
  -   
- Designs of note  
  -   
- Discussion topic  
  - \[mo/mike/liggitt\] proposed semantics for KMS v2 crypto  
    - AES-GCM-SIV considerations  
      - Nonce collisions do not leak keys and/or ciphertext  
        - AES-GCM-SIV is designed to preserve both privacy and integrity even if nonces are repeated. To accomplish this, encryption is a function of a nonce, the plaintext message, and optional additional associated data (AAD). In the event a nonce is misused (i.e. used more than once), nothing is revealed except in the case that same message is encrypted multiple times with the same nonce. When that happens, an attacker is able to observe repeat encryptions, since encryption is a deterministic function of the nonce and message. However, beyond that, no additional information is revealed to the attacker. For this reason, AES-GCM-SIV is an ideal choice in cases that unique nonces cannot be guaranteed, such as multiple servers or network devices encrypting messages under the same key without coordination.  
      - No std lib implementation  
        - [proposal: x/crypto: add AES-GCM-SIV · Issue \#54364](https://github.com/golang/go/issues/54364) is open  
        - Has a [POC implementation](https://go-review.googlesource.com/c/crypto/+/404398) with assembly code  
      - No known FIPS module, [RFC](https://www.rfc-editor.org/rfc/rfc8452.html) is from 2019 so it will likely be a while until FIPS catches up  
      - Google’s tink lib has an [implementation](https://github.com/google/tink/blob/master/go/aead/subtle/aes_gcm_siv.go#L57)  
        - Written in pure Go, unclear if timing vectors are a concern  
    - Behind a (new) beta off by default feature gate to allow toggling it on early  
    - Add new protobuf field in EncryptedObject to store new format  
    - Rather than maintaining a per process DEK, kube-apiserver maintains a per process “seed” used to derive a DEK per write.  
      - “Seed” is derived once per process on startup and is 32 random bytes.  
      - A nonce is generated once per write consisting of 28 random bytes  
        - secretbox considers a 24 byte nonce “long enough that randomly generated nonces have negligible risk of collision”  
      - The first 16 bytes is used as input to the KDF (see below)  
      - The last 12 bytes is used as the data encryption nonce  
      - Encrypted records are stored with a prefix of { nonce, encrypted seed } followed by ciphertext.  
      - We read 32 bytes from HKDFExpand(hash=SHA256, key=\$seed, info=\$nonce\[:16\]) to use as AES-GCM-256 data encryption key  
        - Note that the Extract step is skipped because we already have a good pseudo random key (thus there is no salt, only info)  
      - Ciphertext is encrypted/decrypted using derived AES-GCM-256 DEK using the last 12 bytes of the nonce as the initialization vector.  
    - On reads:  
      - Use cache  
        - unbounded size but limit max growth via  
          - 10 min TTL  
          - etcd path as the key (requires checking that value is valid)  
      - On cache miss, reconstruct DEK via seed \+ public nonce  
    - \[mo\] benchmark a complete read and/or write from etcd with real kube secret data  
      - [AES-GCM vs HKDF Perf Numbers](https://gist.github.com/enj/079d367704a7649bdbe30e93451b9854) (these may be too low level to be meaningful)  
  - \[helayoty\] Update sig-auth-tools repo name so other sigs can reuse github actions/tools. ( [Issue](https://github.com/kubernetes/org/issues/4218#issuecomment-1548416324), more discussion with Anish.  
    - \[liggitt\] does contribx have any generic tools?  
    - \[helayoty\] to look into creating a new shared repo for GH actions with folders per sig  
- Action Items  
  -   
- Sweep issues with leftover time  
  - [CI flakes](https://storage.googleapis.com/k8s-gubernator/triage/index.html?sig=auth)  
  - [CI testgrids](https://testgrid.k8s.io/sig-auth)

## May 10th \- CANCELED

Canceled this occurrence due to a light agenda and multiple leads have conflicts.

## May 9th

* Pairing session with maks, anish, mo on KEP-3331  
* [https://hackmd.io/@enj/BJqZ0xOVn](https://hackmd.io/@enj/BJqZ0xOVn)

## May 8th

* Pairing session with rita, nabarun, liggitt on KEP-3221  
* [Structured Authorizer: Pairing Sessions](https://docs.google.com/document/d/1AJ9xPOMin98-5lZiF2pt_weRrMGSNTmdENjDDW2JMwo/edit)

## April 26th, 11a \- Noon (Pacific Time)

- Recording  
- Announcements  
  - KubeCon EU Sig Auth Deep Dive: [https://sched.co/1HyTv](https://sched.co/1HyTv) checkout slides and recordings  
  - Several good discussions around specific feature designs at kubecon, notes below (April 19/20 headings)  
  - sig-auth enhancements [targeting implementation in 1.28](https://github.com/kubernetes/enhancements/issues?q=is%3Aopen+is%3Aissue+label%3Asig%2Fauth+milestone%3Av1.28+-label%3Asig%2Fapi-machinery)  
    - Double-check people assigned  
  - sig-auth enhancements targeting design work in 1.28  
    - [ReferenceGrant](https://github.com/kubernetes/enhancements/issues/3766)  
    - [Out-of-process JWT signing](https://github.com/kubernetes/enhancements/issues/3908)  
  - sig-api-machinery enhancements adjacent to sig-auth [targeting implementation in 1.28](https://github.com/kubernetes/enhancements/issues?q=is%3Aopen+is%3Aissue+label%3Asig%2Fauth+milestone%3Av1.28+label%3Asig%2Fapi-machinery+)  
- Demos  
  -   
- Pulls of note  
  -   
- Issues of note  
  -   
- Designs of note  
  -   
- Discussion topic  
  - \[rata, mrunalp, giuseppe\] User namespaces support  
    - We added user namespaces support, still in alpha  
    - We wonder how this should interact with PSS (i.e. restricted policy today disallows host namespaces, except for the user namespace).  
    - \[deads2k\] work is likely needed in PSA to be more lenient, is that something we expect to be in individual enhancements or as some kind of submission to sig-auth?  Suggest in original enhancement with an auth approval.  
    - Default behavior in pod uses host user namespace (runAsUser identifies a host uid)  
    - When in a user namespace, relaxing capabilities is safer  
    - baseline wouldn't require user namespaces, it prevents escalating from default pod specs  
    - Q: could baseline relax restrictions on things like runAsUser if userNamespace:true?  
      - deads2k: once we're sure the user namespace limitation is honored all the way to the node, this is possible  
      - liggitt: field graduates to GA, oldest node supported has the field as GA or fails safe by rejecting the pod, positive detection of support for user namespace in container runtime  
      -   
    - Needs a KEP update to userNamespaces KEP  
      - when/how PodSecurity will relax in a controlled way for pods which set userNamespace:true  
        - early opt-in (likely a specific api server feature gate in alpha for relaxing PodSecurity for userNamespaces) up to the providers to ensure the kubelet/container runtime will honor userNamespaces  
        - relax by default only once we can ensure oldest nodes honor userNamespaces or fail safe (and ensure container runtime honors userNamespace or fails safe)  
      - when/how PodSecurity will tighten restricted policy to require userNamespaces:true  
        - liggitt: at earliest, when the field graduates (can't require setting a non-GA field); probably not until userNamespace is supported in stateful pods (seems weird to require setting userNamespace:true in pods but let them just mount a PVC volume to get around the restriction); maybe not requiring it even then… need evidence requiring this is justified  
    - feature flag? Opt-in?  
    - \[rata\]: So, summary of all:  
      - Changing requirements for restricted policy: to change the restricted policy (disallows host user namespaces) to require userns we should wait until the field is GA AND supports stateful pods too. It is not clear, though, if we want to change the restricted policy or not, we will reflect this in the KEP. We still have several months for these changes.  
      - Relax checks when userns are in use: We will create a new feature gate for apiserver only that will enable relaxing some checks in the current policies when userns are in use. The burden to ensure that nodes will honor userns is on the cluster admin. We will clearly document this in the docs. This will be opt-in, and can be by default only once we know the latest supported node (n-3 releases) will honor this   
      - The userns KEP will be updated to reflect these two points and sig-auth will review.  
        - From sig-auth we can assign dead2k/liggit for the KEP review  
      - example of windows podOS KEP [describing PodSecurity changes](https://github.com/kubernetes/enhancements/tree/master/keps/sig-windows/2802-identify-windows-pods-apiserver-admission#changes-to-podsecurity-standards)  
  - \[ahmedtd, sujithrap\] Automatic reloading of k8s CA cert in k8s client libraries?  
    - Which files in particular, which clients in particular?  
    - Kube-apiserver serving certificate root — basically if we rotate the kube-apiserver serving certificate, do clients need to be restarted today?  
    - Injected at /var/run/secrets/kubernetes.io/serviceaccount/ca.crt  
    - Read once [here](https://github.com/kubernetes/client-go/blob/6c596fdbc11ce7779ee320c6a11df44c5d516d54/rest/config.go#L528), does not appear to be periodically reloaded.  
    - automatic setting of reloadTLSFiles doesn't happen if a CAFile is set [here](https://github.com/kubernetes/kubernetes/blob/928adb27abf755b402be7594844e3c248b0b7670/staging/src/k8s.io/client-go/transport/transport.go#L155-L168)  
    - AI: open issue with these steps:  
      - fix setting reloadTLSFiles when given a CAFile  
      - ensure reloadTLSFiles covers re-reading the CAFile  
      - add a test to cover the scenario  
  - \[hoskeri\] Inconsistent authorization of node/ resources: [Inconsistent authorization of node/ resources](https://docs.google.com/document/d/1HHuh70tNDsaARvfZe6soq6yL3aLTFKJcRMoPP5Mk2xM/)  
    - authorization subdivision seems like only part of the issue… if node intends to provide logs/stats/metrics as stable APIs, providing more definition around those and making actual subresources for those seems more natural, and would resolve the authorization issue by default  
    - next step: talk with sig-node about the specific endpoints desired

## April 20th

* robscott, nickyoung, deads2k, mo, liggitt sync on ReferenceGrant  
  * [\[PUBLIC\] ReferenceGrant Scratch Doc](https://docs.google.com/document/d/1Xa0aNpHIv6udg1IEFiVKIsYZ1WYC4JqoMpmg8XaCRSU/edit)  
* ritazh, deads2k, mo, liggitt, nabarun, maksim sync on structured authorization config  
  * path forward in design for previous blockers  
    * CEL filters  
      * mimic matchConditions approach taken by admission webhooks  
      * CEL has access to SubjectAccessReview-shaped representation of request and user attributes  
    * kubeconfig connection info  
      * start with discriminated union  
        * type: "KubeConfig", kubeConfigFile: "..."  
        * type: "InClusterConfig"  
    * superuser authorizer  
      * continue treating superuser authorizer as implicitly at the beginning of the list  
      * in the future, after resolving how to make loopback authorization within the apiserver never fail, could allow representing/reordering the superuser authorizer  
    * example of API shape: [SIG-Auth Deep Dive CloudNativeCon EU 2023.pptx](https://docs.google.com/presentation/d/1Iy8ShHxbwaSX9p13uISGzJeYSjbzrYAc/edit#slide=id.g22c81de7d86_1_104)  
  * next steps  
    * update KEP:  
      * [pal.nabarun95@gmail.com](mailto:pal.nabarun95@gmail.com)  
    * implementation  
      * nabarun: create API types for config file, loading helper  
      * nabarun: refactor existing authz structs to transform to authz config API type  
      * nabarun: move validation of flag-bound structs to API config validation  
      * liggitt: create hot-swappable Authorizer implementation  
      * ?: implement file reload  
        * use poll or watch file changes  
        * detect changes in config / referenced kubeconfig files  
        * replace active authorizer on successful load \+ passes validation  
        * log/metric on failure (TODO: ask han whether this metric belongs in the smaller subset)  
      * ritazh: CEL integration \- start functions for validating / compiling / evaluating expressions with subjectAccessReview context  
        * sync with [Joe Betz](mailto:jpbetz@google.com) or [Cici Huang](mailto:cicih@google.com) on how to make use of the existing CEL validating / compiling / evaluation helpers  
    * \[mo\] rip out all the file watch stuff and replace with simple poll

## April 19th

* micahhausler, mo, ritazh, liggitt sync on 1.28 work at Kubecon  
* kmsv2 \- mo/anish/rita/nilekh  
  * finalize crypto changes  
  * stay in beta in 1.28  
  * plan to ga in 1.29  
  * mark kms v1beta1 deprecated, plan to leave in tree inert.  
* external token signing \- design \- igor, micah, mo  
  * limit to tokens requested with explicit audiences? means default tokens given to pods would be unaffected by external signer unavailability  
  * table showing behavior for tokens requested with no audience, kube-apiserver audience, non-kube-apiserver audience, verified with keysets fetched from kube-apiserver or external keyset URL  
  * likely work on design during 1.28  
  * external signer could sign with ECDSA, RSA or certs?  
* structured authenticator config \- design/impl \- mo/anish  
  * structured config  
  * start with jwt authenticator entries  
  * oidc claim enforcement defaults true, can be turned off  
  * flags can be translated to equivalent config with identical behavior  
    * caveat: revisit oidc-signing-algs option, consider dropping this and allowing all algorithms other than none / symmetric ones  
  * file-reload  
* structured authorizer config \- design/impl \- rita, liggitt  
  * structured config file / mutually exclusive with authorizer flags  
  * multiple webhooks  
  * authorizer ordering  
  * cel filtering  
    * cel variables match subjectAccessReview spec API (maybe grouped under \`request\` top-level variable)  
    * Q: how does CEL handle chaining of null variables like resourceAttributes?  
  * failure policy  
  * file-reload  
  * design in a way to accommodate future authorizers (system superuser one, cel-one, etc)

## April 12th, 11a \- Noon (Pacific Time)

- Recording  
- Announcements  
  - KubeCon EU Sig Auth Deep Dive: [https://sched.co/1HyTv](https://sched.co/1HyTv)   
- Demos  
  -   
- Pulls of note  
  -   
- Issues of note  
  -   
- Designs of note  
  -   
- Discussion topic  
  - \[ivelichkovich\] Discuss external signing of service account tokens (draft KEP PR [https://github.com/kubernetes/enhancements/pull/3855](https://github.com/kubernetes/enhancements/pull/3855))  
    - Opened tracking issue now and updated motivation  
      - Want to decouple signing out of kube-apiserver to allow for custom signing processes such encrypting with non-exportable keys  
      - Also allows for extension  
        - \[liggitt\] what types of extension?  
    - \[deads2k\] \- would we recommend/require that SA tokens are not usable between different clusters?  
      - Not enforceable to make signing keys unique to a cluster  
      - kube-apiserver / TokenReview validation does check existence / uid match of service account object, but that doesn't help validators that aren't using TokenReview and are just doing public key validation on their own  
    - interaction with kube-controller-manager token generation? legacy tokens explicitly requested in Secret API objects will stick around  
      - Nature of the token will not change (non-expiring, but revocable via deletion of Secret object)  
      - Micah \- kube-apiserver flags could be set to not honor the legacy tokens  
      - Is requesting/using an explicitly requested secret-based token in conformance? \[liggitt\] quick sweep didn't show any conformance tests exercising this  
      - Should it be in conformance? (unagreed)  
    - What are we protecting?  
      - liggitt \- user with level of access to kube-apiserver to export a key could also change server config in ways that compromised server auth (inject a new signing key, change client-ca bundle, etc) and relaunch the process  
        - micah \- this is for systems where tokens issued by kube-apiserver are used against other systems that verify those tokens against a keyset obtained from a more trusted source  
        - deads2k \- I wonder if someone would actually build a fork without a non-HSM flag  
    - liggitt:  
      - think the kcm token issue needs an answer  
      - not in favor of making token structure extensible in user-visible ways  
      - think that sig-auth should prioritize progressing/graduating features in-progress or resolving gaps with existing features  
    - mo: could this be done out of tree intercepting tokenrequest API  
      - liggitt: this is gross :)  
      - Could we do something with KMS v2?  
    - next steps:  
      - think about how to resolve the kcm Secret-based token issue in the KEP  
      - talk about this more at kubecon  
  - \[leads\] what KEPs are planned for v1.28?   
    - [Existing sig-auth enhancements issues](https://github.com/kubernetes/enhancements/issues?q=is%3Aopen+is%3Aissue+label%3Asig%2Fauth)  
    - Indicate whether design / implementation / review work is needed, and who is responsible for each  
    - Implementation work planned for 1.28  
      - [KEP-3325](https://github.com/kubernetes/enhancements/issues/3325) (whoami) to GA  
        - implementation: nabokihms (needs some tests, good candidate for a blog post)  
        - review: liggitt  
        - effort: \~low (mechanical graduation, no changes expected)  
      - [KEP-2799](https://github.com/kubernetes/enhancements/issues/2799) (reduction of secret-based tokens)  
        - legacytokentracking to GA  
          - impl: zshihang  
          - review: liggitt  
          - effort: \~low (mechanical graduation, no changes expected)  
        - legacytokencleanup to Alpha ([PR \#115554](https://github.com/kubernetes/kubernetes/pull/115554))  
          - impl: yt2985  
          - review: liggitt  
          - effort: \~medium (design complete, new controller impl, testing is more difficult)  
      - [KEP-3299 (kmsv2)](https://github.com/kubernetes/enhancements/tree/master/keps/sig-auth/3299-kms-v2-improvements) staying in beta for 1.28  
        - impl: Anish and mo  
        - \[deads2k\] this was using some apimachinery features, there was apimachinery agreement to review changes to get to beta.  Any unexpected blockers (I remember a couple bugs) or just work?  
        - \[mo\] kms v2 crypto changes is a blocker for GA  
        - reviewers:  
          - apimachinery: deads2k  
          - auth/crypto: mikedanese?  
        - effort: \~medium (some crypto changes, review/impl/test work)  
      - [KEP-3257](https://github.com/kubernetes/enhancements/issues/3257) (cluster trust bundles)  
        - ClusterTrustBundle projected volumes to alpha  
          - impl: ahmedtd  
          - reviewers:  
            - liggitt (for API)  
            - node/storage reviewers: ???  
          - effort: \~medium (design complete, new pod field, kubelet/volume impl)  
        - ~~certificatetrustbundle beta?~~ — \[ahmedtd\]: I think it makes more sense to bring both CTB and Workload Certificates to alpha first.  
        - workload certificates?  
          - \[ahmedtd\] Currently drafting KEP content in [this doc](https://docs.google.com/document/d/1epJMmmXZGDlyNtx-bSoehecTfGy5PF8ketynfaWRnzw/edit#heading=h.98pyuk9izozq), will promote to real KEP with number ASAP  
      - [KEP-3221](https://github.com/kubernetes/enhancements/issues/3221) (multiple authorization webhooks)  
        - design: palnabarun  
          - [https://github.com/kubernetes/enhancements/pull/3376](https://github.com/kubernetes/enhancements/pull/3376)   
          - \[tim\] Can we add scoping rules to the multiple authz webhooks wishlist? i.e. only send requests matching the rules to the webhook (no opinion for everything else)  
            - \[mo\] \+1, \[liggitt\] \+1  
          - \[liggitt\] specifically interested in running more than one webhook and being able to configure failure policy to mean "deny" instead of "no opinion"  
          - \[mo\] exposing multi-webhook authz as a REST API?  
            - \[liggitt\] not in favor of that on by default or in conformance; seems like a separate effort from this KEP  
        - impl: palnabarun, ritazh, liggitt  
        - review: liggitt  
      - [KEP-3331](https://github.com/kubernetes/enhancements/issues/3331) (structured config for jwt/oidc authenticator)  
        - Design: [https://github.com/kubernetes/enhancements/pull/3332](https://github.com/kubernetes/enhancements/pull/3332)  
          - AI: mo/liggitt hammer out config surface area  
        - prereq: improving test coverage for existing OIDC impl  
        - impl: Maksim  
        - reviewers: Mo and liggitt  
    - Design work for 1.28  
      - [KEP-3766](https://github.com/kubernetes/enhancements/issues/3766) (referencegrant)?  
        - Design: [https://github.com/kubernetes/enhancements/pull/3832](https://github.com/kubernetes/enhancements/pull/3832)   
        - liggitt: this seems like ⅓ of a solution to a real problem (expressing resource-to-resource access), but relies on controller authz and controller behavior; the likely outcome is overly broad authz grants to controllers and divergent behavior of controllers in making use of those permissions on behalf of resources  
          - \[liggitt\] would like to see us free up bandwidth to actually tackle the whole problem, where user expresses a cross-resource permission grant, that actually takes effect to grant authorization, and a controller can make use of that permission by indicating it is making a request on behalf of that resource  
          - \[mo\] will talk to Rob at KubeCon to brainstorm ideas  
          - \[deads2k\] on-behalf-of is interesting to me too  
        - AI: sync with rob/nick  
          - done, discussion/notes [here](#bookmark=id.c6wxnm1vx2m9)  
      - [KEP-267](https://github.com/kubernetes/enhancements/issues/267) (kubelet serving certificate) to beta?  
        - Design: SergeyKanzhelev  
          - [https://github.com/kubernetes/enhancements/pull/3806](https://github.com/kubernetes/enhancements/pull/3806) (retroactive KEPification of current state)  
        - \[deads2k\] any changes planned or just taking it through to stable?  
          - \[liggitt\] first step was KEPifiying current state in beta and evaluating any gaps  
      - External token signing: [https://github.com/kubernetes/enhancements/pull/3855](https://github.com/kubernetes/enhancements/pull/3855)   
    - FYI, apimachinery work planned for 1.28 adjacent to sig-auth:  
      - [KEP-3488](https://github.com/kubernetes/enhancements/tree/master/keps/sig-api-machinery/3488-cel-admission-control) (CEL admission control) to Beta  
        - impl: jbetz  
        - review: sig apimachinery reviewers  
      - [KEP-3716](https://github.com/kubernetes/enhancements/tree/master/keps/sig-api-machinery/3716-admission-webhook-match-conditions) (webhook match conditions) to Beta  
        - impl: Igor Velichkovich  
        - review: sig apimachinery reviewers  
    - Not planned for 1.28  
      - [~~KEP-3737~~](https://github.com/kubernetes/enhancements/issues/3737) ~~(fine-grained authz)?~~  
        - ~~Design: [https://github.com/kubernetes/enhancements/pull/3617](https://github.com/kubernetes/enhancements/pull/3617)~~   
        - ~~\[deads2k\] \- lavalamp indicated that he wasn’t likely to have time to implement this cycle~~  
        - ~~impl:~~  
        - ~~review:~~  
  - \[mo\] kms v2 crypto changes: [https://hackmd.io/@enj/SyiXCABZn](https://hackmd.io/@enj/SyiXCABZn)  
    - \[mike\] not sure how realistic this concern is as the use of nonce counter is pretty common. e.g. tls  
    -   
  -   
- Action Items  
  -   
- Sweep issues with leftover time  
  - [CI flakes](https://storage.googleapis.com/k8s-gubernator/triage/index.html?sig=auth)  
  - [CI testgrids](https://testgrid.k8s.io/sig-auth)

## March 29th \- CANCELED

Canceled this occurrence since multiple leads have conflicts.

## March 15th \- CANCELED

Canceled due to proximity to release deadlines.

## March 1st, 11a \- Noon (Pacific Time)

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
  - \[ivelichkovich\] Discuss external signing of service account tokens (draft KEP PR [https://github.com/kubernetes/enhancements/pull/3855](https://github.com/kubernetes/enhancements/pull/3855))  
    - still an [open question](https://docs.google.com/document/d/177duVuRwPCBjrFot0aRptw49JDdHDdLNUGMXxSekO9w/edit?disco=AAAAlo1qUck) about how out-of-process signing interacts with kube-controller-manager token issuance  
    - PR indicates motivation is to support getting a new signing key without restarting the API server; would dynamic reloading of signing key to allow rotation without restart be sufficient?  
      - much smaller surface area than a grpc API  
      - plenty of prior art for dynamic reloading CA bundles / client certs / token files / config files  
  - \[ahmedtd\] [Pre-KEP: Workload Certificates](https://docs.google.com/document/d/1epJMmmXZGDlyNtx-bSoehecTfGy5PF8ketynfaWRnzw/edit)  
    -   
  - \[mo\]: KMS v2 DEK re-use  
    - [https://github.com/kubernetes/kubernetes/pull/116155](https://github.com/kubernetes/kubernetes/pull/116155)  
    - bound time from KMS rotation to new key being used (KMS health check returns keyid change, which will trigger getting a new DEK)  
    - mike: could we use a write counter or incrementing IV instead of a time-based approach to spawning a new DEK?  
  - side-bar about how to handle undecryptable items  
    - currently, undecryptable objects in etcd fail gets/lists/updates/deletes  
    - only current remedy is to delete keys directly from etcd (\!)  
    - we \*just\* started logging etcd keys of undecodeable items in [https://github.com/kubernetes/kubernetes/pull/114376](https://github.com/kubernetes/kubernetes/pull/114376) but that only helps someone with kube-apiserver log access, and the problem is still not fixable via the API  
    - Opened this [https://github.com/kubernetes/kubernetes/issues/116194](https://github.com/kubernetes/kubernetes/issues/116194) to track the request  
    - if someone wanted to look into aggregating etcd paths that failed to decode in a list request to include in the error response, that could provide someone the information required to initiate a cleanup  
    - we could consider allowing a delete request of an object where the old object could not be decoded from storage  
      - what object would we send to admission? would we bypass admission?  
      - we wouldn't be able to validate finalizers or resourceVersion / uid preconditions  
  - \[Mahé\] Would need some help [on scdeny warnings](https://github.com/kubernetes/kubernetes/pull/115879#issuecomment-1436076805) and need someone for sig-auth to triage this PR [https://github.com/kubernetes/kubernetes/pull/115879](https://github.com/kubernetes/kubernetes/pull/115879) read that [https://kubernetes.io/blog/2020/09/03/warnings/](https://kubernetes.io/blog/2020/09/03/warnings/) but still wondering.  
    - server log, guard with a new feature gate marked as alpha and disabled by default  
  - \[ahmedtd\] ClusterTrustBundles: Kubelet now supports mounting by a combination of signerName and label selector.  Do we still need support for mounting a single ClusterTrustBundle by name?  
    - Liggitt: Can we get some usecases? Label selectors are flexible, but how does the user know which selector to specify. Ideally, most of the in-tree use cases wouldn't need more than the signerName ("give me all trust bundles for kubernetes.io/kube-apiserver-client")  
- Action Items  
  -   
- Sweep issues with leftover time  
  - [CI flakes](https://storage.googleapis.com/k8s-gubernator/triage/index.html?sig=auth)  
  - [CI testgrids](https://testgrid.k8s.io/sig-auth)

## February 15th \- CANCELED

Canceled due to a light agenda with multiple folks out.

## February 1st, 11a \- Noon (Pacific Time)

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
  - \[mo/nilekh\] okay to only support \*.\* by itself for encryption at rest?  
    - config file indicates resources to encrypt as list of \`\<resource\>.\<group\>\` entries  
    - jordan: user experience of starting at the top of the config file and first matching config for a resource wins seems reasonable to understand and document  
    - david: going in order makes sense, allows carve-outs with earlier entries  
    - mo: early entry of \*.\* would mask later entry of "secrets"?  
      - david: yes  
      - jordan: would be nice to catch/warn masking configurations that make later entries no-ops if possible  
  - [keps/sig-auth/3766-referencegrant](https://github.com/kubernetes/enhancements/tree/master/keps/sig-auth/3766-referencegrant)  
    - merged as provisional with plenty of unresolved / open questions  
    - this design assumes a controller which has access to the granted resources is self-limiting behavior based on inspections of referencegrant objects  
    - mo: are we sure about codifying cross namespace authz stuff?  
      - Skipped over for now.  Gateway API believes so, so does the the storage usage.  But in general?  
      - jordan: if we have a reasonable story around visibility of information and how the actual authorization is granted, the cross-namespace aspect seems fine… Role/RoleBinding can grant cross-namespace access to serviceaccounts in other namespaces  
    - Jordan \- will we be happy with all controllers using referencegrant having read access to the info in all referencegrants?  
      - Different controllers have different permission levels. Should they all be able to see every referencegrant targeted at other (potentially higher privileged) controllers? referencegrant contains info about granting namespaces, granted resource types/names, and (potentially) consuming namespaces/types/names  
    - Jordan \- are we encouraging granting global read access on the referenced object type to controllers, trusting them to self-limit using referencegrant? since referencegrant doesn't actually grant access, some other RBAC grant step to the controller is required, and that's likely to be set up by the admin setting up the controller, not the user creating the reference grant  
      - user doesn't know the controller identity  
      - in theory, it's possible to grant the controller narrow `get` access object-by-object, but global get is way easier and more typical (rob: this is how most gateway controller access to secrets is granted)  
    - deads \- how does a user know that their reference does or doesn’t do a thing?  
      - Potentially status, but it would be the absence of information.  
    - deads \- How do we know which controller uses a particular instance?  
      - There is no way to determine this now  
      - Potentially empty status would indicate ineffective.  
    - Deads \- do specific resources with identical serializations provide value?  
      - Con \- Kubectl is really hard to use in this case.  
      - Con \- There’s no guarantee of consistent serialization  
      - Pro \- resolves controller permissions from above  
      - Pro \- resolves user permissions about which types of reference grants can be used  
      - Pro \- resolves the problem of which controller is expected to act on the reference grant  
    -   
    - open questions from the KEP:  
    - I still find the from/to a bit confusing  
      - "subjects"/"rules" terminology would match RBAC  
      -   
    - Are we sure we cannot model this as authz checks?  
      - Possible if we decide to either rule out label selectors forever or add them to RBAC  
      -   
    - Probably should mention explicitly that all versions of APIs are equivalent  
      - i.e. granting access to secrets v1 means also granting it for v2  
      -   
    - What are the actual SAR checks being run by KAS?  
      - Ties into the open questions around supporting “verbs”  
      -   
    - Should reference grants be targeted at a single controller with a field that can be selected on?  
      - No, at least for Gateway API, we want ReferenceGrant to apply regardless of the underlying controllers. I can’t think of a scenario where that would be different.  
      - It could be valuable to support “verbs” though  
      -   
    - What is the Go interface that controllers will use to check access?  
      - Similar to SAR as starting point  
      - Would ideally have some kind of centralized cache that could inform controllers when cross-NS references had been revoked  
        - Requires controllers to send every update of a resource that includes/included cross-NS reference to this lib  
      -   
    - Will we provide any automation/tooling (maybe a controller) to help migrate from the old grant API to the new grant API?  
      - I think we should provide a CLI tool that converts old resources to new resources (similar to ingress2gateway)  
      -   
    - Reminder that you need both “Describe the mechanism: Enable alpha ReferenceGrant API” and “enable the actual REST API” to enable the feature  
    -   
  - \[krzys/standa\] progress on kube-rbac-proxy and request for second review: [pre-acceptance issue](https://github.com/brancz/kube-rbac-proxy/issues/169)  
- Action Items  
  -   
- Sweep issues with leftover time  
  - [CI flakes](https://storage.googleapis.com/k8s-gubernator/triage/index.html?sig=auth)  
  - [CI testgrids](https://testgrid.k8s.io/sig-auth)

## Jan 24th, 9a (Pacific Time), KMS-plugin meeting

- Recording  
- Agenda:  
  - Review KMSv2 board for beta blockers  
  - Metrics  
- Discussion notes:

## January 18th \- CANCELED

Canceled due to an empty agenda.

## January 17th, 9a (Pacific Time), KMS-plugin meeting \#21

- Recording  
- Agenda:  
  -   
- Discussion notes:  
  - 

## January 4th, 2023, 11a \- Noon (Pacific Time) {#january-4th,-2023,-11a---noon-(pacific-time)}

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
  - \[mo\] any plans for: forward additional request metadata to \*Reviews  
    - [https://github.com/kubernetes/enhancements/pull/2843](https://github.com/kubernetes/enhancements/pull/2843)?  
    - mike: unsure how to get per-request metadata to play nicely with the authentication cache  
    - mo: would an authentication proxy approach be able to access the desired metadata without plumbing more into tokenreviews?  
    - mike: possibly, yes.  
    - mike: unsure of whether this approach is a general need  
      - mo: not aware of needs from azure for this  
      - david: not aware of needs from redhat for this  
    - AI(mike): summarize considerations and discussion on PR and close  
  - \[riaan\] [Write e2e test for SubjectAccessReview & createAuthorizationV1NamespacedLocalSubjectAccessReview \+2 Endpoints \#114345](https://github.com/kubernetes/kubernetes/pull/114345)  
    - Thank you for the guidance Jordan & Mo.  
    - Can we please get some eye and review. \[liggitt: did a review pass\]  
    - merged in 2022  
  - \[mo/Mahé\] Want to have people of the SIG their opinion on the SecurityContextDeny admission plugin removal. Some info on the Github issue [https://github.com/kubernetes/kubernetes/issues/111516](https://github.com/kubernetes/kubernetes/issues/111516), it was firstly discussed [in a Slack thread here](https://kubernetes.slack.com/archives/C019LFTGNQ3/p1658755949311999). I got support of @liggitt and @tallclair from the SIG for now I think. [We discussed a little bit about the KEP](https://kubernetes.slack.com/archives/C019LFTGNQ3/p1665068965083089) for removing this part, @sftim was in favor of writing a retrospective KEP for scdeny (this addition was pre-KEP, it was a design proposal at the time) in order to remove it.  
    - Liggitt: it’s not useful, provides no protection, but it has not required any maintenance. We should sweep github and see if there’s any usage. If we can’t find any use or do find use and reach out, they fix, we should delete. If we find use and there are valid use cases, we can revisit.  
    - liggitt: updated description with action items  
  - \[mo\] thoughts on [https://github.com/kubernetes/kubernetes/issues/111208](https://github.com/kubernetes/kubernetes/issues/111208) ?  
    - Enhancement (not a bug) to increase the reliability of file based audit logs such as f-sync-ing the file, making sure the file is still valid, responding to signals?  
  - \[ahmedtd\] KEP-3257 for 1.27: Should the kube-apiserver and kubelet halves be squished into one giant PR?  
    - \~3-4 PRs make sense. Separate TrustBundle and project volume API changes. Separate admission. Separate kubelet side changes.  
    - ~~AI(liggitt): add ahmedtd to API review meeting on 2023-01-05.~~ done  
  - \[mikedanese\] Client Exec Auth plugin in 1.27  
    - mike: merging [https://github.com/kubernetes/kubernetes/pull/113639](https://github.com/kubernetes/kubernetes/pull/113639) would be simplifying pre-work: collapse spdy transport if possible. what should we look for to build confidence in this refactor?  
      - Liggitt: Check coverage of changes we’ve made in SPDYTransport:  
        - Ping period, long lived exec  
        - HTTP proxy, socks proxy  
        - Look through history of spdy transport and make sure fixed issues are still represented in tests  
    - Liggitt: Exec’ing random commands specified in kubeconfig is scary. Can we make this any safer?  
      - \[mo\] would like to add this as an allowlist config in .kuberc  
  - \[liggitt\]: whoami should go Beta in 1.27. \- [https://github.com/kubernetes/enhancements/issues/3325](https://github.com/kubernetes/enhancements/issues/3325)   
    - ~~AI(liggitt): check with @nabokihms about driving to beta in 1.27.~~ done  
  - \[tallclair\]: Webhook exclusions in 1.27. More in API Machinery, but relevant here: [https://github.com/kubernetes/enhancements/pull/3694](https://github.com/kubernetes/enhancements/pull/3694)  
    - AI(tallclair): propose an alternative KEP using CEL.  
    - Deads: There’s inconsistency between activation/bindings in all our CEL based extensions. Let’s address that in the KEP (clarify what variables/functions are available to the CEL expressions).  
  - \[mikedanese\]: Dynamic Admission integration with TrustBundles was skipped in the last cycle. Let’s see if we can agree on a design for that in 1.27.  
- Action Items  
  -   
- Sweep issues with leftover time  
  - [CI flakes](https://storage.googleapis.com/k8s-gubernator/triage/index.html?sig=auth)  
  - [CI testgrids](https://testgrid.k8s.io/sig-auth)
