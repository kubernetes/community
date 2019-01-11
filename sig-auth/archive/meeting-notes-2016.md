# Kubernetes Sig-auth Meeting Agenda

## December 14th, 2016, 11a - Noon (Pacific Time) - [recording](https://youtu.be/XaPDkb5hFww)



*   Presentation on Openshift Service server cert annotations
    *   Signs a serving certificate for an in cluster DNS name: `service-name.namespace.svc`.
    *   Separate distributes a CA bundle to all pods by piggy-backing on the auto-mounted SA secrets.
    *   A stolen key can't be used by a pod in another namespace since they can't control traffic routing.
    *   Does *not* sign client certificates since those could be re-used from anywhere.
    *   Getting pods to reload updated serving certificates is hard.
    *   Notes:
        *   Usage in openshift here: https://github.com/openshift/origin/blob/master/test/extended/cmd.sh#L182-L188
        *   Add an annotation to your deployment
        *   Get a serving cert in your service account secrets
        *   Ingress only does TLS termination, not re-encryption
        *   Rotation is hard.
*   Demo of https://github.com/kelseyhightower/vault-controller  (Kelsey, 10 minutes demo + 10m questions)
    *   Notes:
        *   Prototype.
        *   Kubernetes doesn't have "real" secrets.
        *   Token per pod
            *   Use an init container which talks to the vault controller
            *   Vault controller looks up a pod with that name
            *   Once the vault controller is confident that it's actually talking to a pod, it will give it a token.
                *   E.g. uses pod IP to actually send token to pod
        *   Dynamic secrets that can have short TTLs and renewing
        *   Vault only lets one client decode the token
            *   Raises an alarm if it's been tried to be unwrapped twice
        *   There is no way to secure the use of particular annotations today.
        *   Still have to handle key/cert updates.  Health check?  Rotation is still hard.
        *   How does this handle generic secrets?
        *   Policies determine what can read what
        *   Vault controller needs to be able to have access to all policies it might delegate to pod.


## November 30th, 2016, 11a - Noon (Pacific Time)



*   Sig Self Assessment [15 min?]
    *   Does your SIG own code in kubernetes/kubernetes?  Elsewhere?  Within kubernetes/kubernetes, is it reflected in OWNERS files?  If not do, you see a path to doing that?
        *   /pkg/apis/{abac,authorization,authentication,rbac}
        *   /pkg/auth
        *   /plugin/pkg/auth
        *   No code outside the main repo
        *   OWNERS? Sig-auth leads in many owners files as individuals, but not aliased as sig-auth. Github kubernetes/sig-auth team is not synced with the google group. Should it be?
    *   Does your SIG spend (not enough, the right amount, too much) time responding to user issues in Github?  On stackoverflow?  On your SIG mailing list?
        *   Mailing list not great. Slack channel is amazing!
        *   Document what kind of questions are appropriate for which forum.
            *   E.g. slack isn't searchable.
        *   Move frequent questions/answers to doc
    *   Is the balance between new features work and stabilization good, over the last 2 releases?  Do you feel the need to focus on stabilization in the next release?  Something else?
    *   Overall, how do you feel good about the level of test coverage for the code you are responsible for?  Docs completeness?  Number of open bug reports?
        *   Docs have a lot of catching up to do.
        *   Unit test coverage is generally good
        *   Integration test coverage is less good… difficult to start servers with lots of different configurations
    *   fill out survey, once per sig, after the meeting. https://docs.google.com/a/google.com/forms/d/e/1FAIpQLScR4dkTyfP56cQdKyqHWMJl3glLfyu46EXqCxAH7Bd-xkJ8cw/viewform
*   1.6 objectives. P1 - staffed and the release was unsuccessful if we don't get it.  P2 - really want it, will devote review time to it, but no one has signed up to do it.  P3 - nice to have, but not being actively worked
    *   Do we have a pithy threat model that we're focusing on for 1.6?  E.g., "compromise nginx running on cluster does not give root in cluster w/o further vulns"?  Possible to configure pods that way, then by default?
        *   Explains to users+devs why we're doing the work below.
        *   Pod security context and RBAC are the work being done for the nginx example
        *   Also need docs on how to use them all together.
    *   RBAC to beta - P1,big - deads, sttts, someone from test please?
        *   Default roles for controllers - pulls in progress
        *   Controller, scheduler, kubelet, etc identity - pulls in progress
        *   Local-up-cluster support - pulls in progress
        *   CI test that enforces controller, kubelet, other infrastructure permissions - need assistance here.  We're on the agenda next meeting: https://docs.google.com/document/d/1z8MQpr_jTwhmjLMUaqQyBk1EYG_Y_3D4y4YdMJ7V1Kk/edit  cjcullen says he'll be able to help once we've reached a point where it make sense to wire this up.
        *   ~~Default (minimal) roles for service accounts (don't want compromised pod to be able to wreak havoc).  This is already done (default).~~
        *   Document how to effectively use RBAC for API protection by default.
        *   Tracking issues to enable in deployment mechanisms (aws, gce, kubeadm, kops, etc, etc, etc) (ericchiang can help with this)
    *   Authn/Authz with federated API servers in a single cluster - P1,small - deads, sttts
        *   API machinery has existing plan: https://github.com/kubernetes/kubernetes/blob/master/docs/proposals/federated-api-servers.md
        *   Pulls in progress to make auth proxy available as the standard authentication
        *   POC being brought into kube demonstrating the concept here: https://github.com/openshift/kube-aggregator
    *   Credential rotation {possible | on by default} - P2,big?
        *   Credential revocation
            *   Service account tokens - delete secret, validate against etcd
            *   X509 certs - CRL?
        *   Credential rotation
            *   Service account token
            *   X509 cert
        *   Credential signing/granting rotation
            *   Token signing keys
            *   CA signing certs
    *   Whitelist paths for host dir (erictune, P2?)
    *   Fixing open issues with kubectl/OIDC.  No net new features. (ericchiang, P2?)


## November 16th, 2016



*   No SIG meetings this week


## November 2nd, 2016, 11a - Noon (Pacific Time)



*   Security response team update
    *   Proposal in progress: https://docs.google.com/document/d/1dGzReKkYdRwyj2uYFMeIJbZesnBlhWgCeuSJ9af4gbE/edit
*   1.5 items - any exceptions required?
    *   https://github.com/kubernetes/features/issues?q=is%3Aopen+is%3Aissue+label%3Ateam%2FSIG-Auth
    *   Kubelet authn/z - [in](https://github.com/kubernetes/kubernetes/pull/34381)
    *   Auth proxy - [in](https://github.com/kubernetes/kubernetes/pull/35452)
    *   SA credential rotation - [in](https://github.com/kubernetes/kubernetes/pull/34029)
    *   RBAC
        *   Bootstrapping, discovery role/rolebinding, default roles in
        *   will continue work next release
            *   Needs mutable, secure API server summarization
    *   PSP - user-specific PR pending ([#33080](https://github.com/kubernetes/kubernetes/pull/33080))
    *   CRL - pending review, unsure if it will be ready for freeze (maybe early 1.6)
        *   https://github.com/kubernetes/kubernetes/pull/33519
        *   https://github.com/kubernetes/kubernetes/pull/35698
        *   Short-lived certs (must be easily rotatable) vs easy revocation (performance hit of CRL lookups)
*   F2F at kubecon?
*   FYI: Kelsey's vault integration sketch
    *   https://github.com/kelseyhightower/vault-controller
    *   Feedback on this please.
*   1.6 planning
    *   Threat models (master, node, pod compromise, etc)
    *   Appropriate use of secrets (encryption requirements, who holds keys, type and access segmentation, etc)
    *   Credential rotation (service account tokens, client/server certs)
*   Intent based RBAC API
    *   https://github.com/kubernetes/kubernetes/pull/31441
    *   Declarative vs imperative
    *   Post-1.5, david and clayton to follow up
*   OpenID Connect client updates
    *   Upstream rewrite of coreos library
    *   Better compatibility with arbitrary oauth/oidc servers
    *   External test coverage (in kube usage of OIDC) before switching
    *   ericchiang to open tracking issue to switch, summarize changes


## October 19th, 2016, 11a - Noon (Pacific Time)



*   Expanding the sig-auth charter (15 min - leads)
    *   Sig-security proposal: https://groups.google.com/forum/#!topic/kubernetes-dev/7YvCC3GrvJE
    *   Current charter: https://gist.github.com/erictune/a123f2df7ce8b5fcdd33
    *   Proposed charter: https://gist.github.com/liggitt/008700a8f22f47b8f75f114dfba21540
*   Security response team / vulnerability management team
    *   Working group of sig-auth
    *   Goals:
        *   Small team to privately triage CVE/security reports
        *   Coordinate with appropriate sig leads / feature owners
        *   Coordinate with releng on point releases or backports for security fixes
        *   Coordinate with vendors/distros on distribution of security fixes?
    *   Two phases:
        *   Figure out how to create and structure the security response team.
            *   One or more sig auth leads to be involved initially?
        *   Security response team executes on mission.
    *   Prior Art:
        *   https://security.openstack.org/vmt-process.html (OpenStack Vulnerability Management Team)
*   Networking round two (20 min - deads) with Dan Williams
    *   [Prior discussion](#september-21st-2016-11a---noon-pacific-time)
    *   What is a tenant?
    *   Recap:
        *   OpenShift: Namespace networks start isolated, can become shared.
        *   Potential for N-to-N relationship where users may need access to some subset of networks.
        *   Current network policy in upstream is not trying to solve tenancy.
            *   Does ingress (no egress), you can whitelist other things to send traffic to you.
            *   Contention around cluster administrator vs app developer.
    *   What is networking sig doing to make this easy to use?
        *   Simple policy: Default deny then you can enable the things you want
        *   Expect higher level tools to control these rules
        *   Similar to ec2 security groups.
    *   Who is the audience?
        *   You write an app with a service and want to expose it to some subset of namespaces.
        *   Again, expect high level tools (e.g. http://www.opencontrail.org/)
        *   Intent based? E.g. if I configure a service I clearly expect the pods backing it to expose specific ports.
    *   RBAC intent. I allow access for "david" to access "pods/proxy"
        *   How would you set up policy in human terms. E.g. This thing can access another thing.
        *   Unreliability of source IP.
    *   How to handle cross-namespace resources
        *   Possibly a higher level concept (e.g. a deployment vs pods)
    *   Next steps
        *   Try to help sort out what a tenant is, or at least summarize what others think of it.
        *   David Oppenheimer will help facilitate that document.
            *   What is a tenant?
            *   Use cases


## October 5th, 2016, 11a - Noon (Pacific Time)



*   Kubelet authn/authz proposal (10 min)
    *   https://github.com/kubernetes/kubernetes/pull/32518
    *   GKE considerations
        *   Can continue running in current mode (no authn/authz)
        *   LGTM'd by CJ Cullen, with reservations around SAR for kubelet
        *   Will continue discussion in PR
*   Review 1.5 goals (15 min) (hopefully with erictune)
    *   Month and a half of time left
    *   rbac to beta? (deads2k)
        *   On by default?
        *   Biggest risk is SA controllers on by default
        *   Eric will follow up on benchmark stuff
        *   Intent based API
            *
    *   login (eric chiang)
        *   Not for 1.5
        *   CoreOS de-prioritized, but would not block it.
    *   SA credential rotation (liggitt)
        *   Pull open to allow multiple signer key recognition to tokens.
        *   Currently nothing that decides when and how the old tokens go away.
        *   https://github.com/kubernetes/kubernetes/pull/34029
    *   File CRL (eric chiang)
        *   https://github.com/kubernetes/kubernetes/pull/33519
    *   X-Remote User authenticator (liggitt)
        *   On the TODO list
    *   Kubelet authn/authz (liggitt)
        *   https://github.com/kubernetes/kubernetes/pull/32518
        *   CSR for server cert
            *   Use the client cert to request a server cert
    *   Make insecure port unnecessary (deads/dims)
        *   Now possible, but not easily possible
        *   Needs refinement and probably RBAC on by default
        *   Controllers and the schedulers use it
            *   Give these "system:" identities
        *   Probably end up with possible in 1.5, but not easy until 1.6
    *   No current work
        *   Kubernetes as an Authenticating Proxy
        *   Default Service Account Power Reduction (authz addresses power reduction, no work in progress to omit token entirely)
*   Presentation about Open Policy Agent by Torin Sandall
    *   [ www.openpolicyagent.org](http://www.openpolicyagent.org)
    *   [15 min?]
    *   Some discussion around remote admission controllers
        *   How do we see this going in the future?
        *   Process based plugin?
        *   Implementing a generic callout would mean we now expose every object to our webhooks. Might be hard to reason about how changes to internal objects would impact every outside evaluator.
    *   Proposal to split the admission controller interface into multiple interfaces: https://github.com/kubernetes/kubernetes/pull/34131
    *   Video! https://www.youtube.com/watch?v=34vDoGSi0JY
    *


## September 21st, 2016, 11a - Noon (Pacific Time)



*   Joint meeting with Sig-network (thockin) 40min
    *   Network ingress policy objects (egress not supported yet)
        *   https://godoc.org/k8s.io/kubernetes/pkg/apis/extensions#NetworkPolicy
        *   Determine which pods can talk to each other
        *   Also select a set of namespaces (label selectors) which can talk to each other. Note: assumption that only admins can create namespaces.
            *   Openshift differentiates between "I want to create this object" and "I want an admin to create this for me"
            *   Further policy (whitelist) around what labels you can set.
        *   Policy that determines what traffic others can send to your pod
            *   Either with your namespace
            *   Or on another namespace
        *   Only additive
    *   How does RBAC complicate network policy?
        *   (chris marino) RBAC doesn't map well to network policy today
        *   RBAC review:
            *   Two levels: cluster admin vs namespace admin
            *   Namespace admin shouldn't be able to impact other namespaces
        *   Network policy side
            *   Mimic AWS security groups
                *   Create a group (e.g. I expose port 80, will only accept traffic from IP foobar, etc.), add a VM to that group
            *   Scope by label selector
            *   Goal: Operator policies that take precedence over user policies
    *   Challenges with using labels on namespaces to identify common policy across namespaces
        *   Label values are limited
        *   Openshift had to use annotations
        *   Who controls the labels?
    *   Notes from RBAC: Grant vs Accept
        *   Just because I give you something, doesn't mean you want it (possible phishing attacks by granting someone access to things with misleading names)
    *   How does RBAC, or its model, extend to network policy?
        *   What level of flexibility/complexity does this need to support?
        *   Types of separation
            *   Dev vs. production, service provider vs. customer, admin vs. teams
            *   Can we provide hierarchy on policy? Namespaces are nice, but they can often overlap.
    *   NOTE: sig-auth might be a good place to discuss general policy when other sigs want to implement something


## September 7th, 2016, 11a - Noon (Pacific Time)



*   "What can I do?" endpoint (david) 10min
    1.  Users need to be able to predict which actions they can perform.  SAR and SelfSAR give them one at a time checks, but getting a complete list is very useful.
    1.  Should not be specific to a particular authorizer.
    1.  https://github.com/kubernetes/kubernetes/issues/31292
    1.  Namespaced or cluster scoped?
        *   Yes, but as a secondary objective
    1.  Would like to see an implementation
        *   How to express queries about non-namespaced objects or cluster wide rules
        *   How to deal with stars, do we return wildcards? authorization.Attributes lacks a current rule for dealing with this.
        *   How to indicate incomplete information.
    1.  Interaction with UI is important (do I show a delete button?), but a heavy operation
        *   Batching important
    1.  Next steps
        *   actually review clients that will use this (demo of OpenShift by liggit)
        *   Propose an API and have UI builders review
*   Tighten RBAC rules (david) 10min
    1.  Stop cross-namespace role references
        *   Currently, fail at runtime
        *   Still allow referencing service accounts in different namespaces
    1.  Require `Kind` and `APIGroup` on the roleref (could be defaulted)
        *
    1.  Breaking changes
    1.  Other notes: bootstrapping of roles
*   Call for 1.5 Goals (erictune) 20min
    1.  Kubernetes as an Authenticating Proxy - P2
        *   OpenShift is interested but does not expect to have Resources.
        *   Google is interested not sure about allocating resources.
        *   Dashboard would like this.
        *   Not well designed yet.
        *   OpenShift interested in the "allow proxy within namespace to anyone with namespace access" use case.
        *   Feels like plan but not execute in 1.5 item, or do it in pieces.
        *   Experience: If I am authenticated to the Kubernetes API, I can be authenticated to addon/apps/charts running on the cluster. Addons and Apps don't need to implement Authn.
        *   Example users of Proxy: K8s Dashboard, Federation Apiserver, Grafana, Spark driver dashboard,.
        *   Jordan says likes the idea, but wants to do the right thing with cookies
    1.  Kubernetes API can sit behind an Authenticating Proxy - P1 (deads2k/liggitt)
        *   What: Cluster deployer can put an authenticating proxy in front of Kubernetes APIserver.
        *   Request Header upstream -> Jordan or David is likely to upstream this in 1.5 time frame.
    1.  Authorization - P1 (deads2k) (etune or ericchiang to review)
        *   What: Review contributions to RBAC, ensure goes to Beta.
        *   Why: because OSS needs authorization
        *   David will drive this.
        *   Requirements: accuracy (deads2k), performance (deads2k, but eric chiang will help with testing), default roles (deads2k)
    1.  Default Service Account Power Reduction - P3
        *   Not issue for OpenShift/RBAC.
        *   Why: currently service accounts are "root" by default on GKE, and changing the default is annoying
            *   Turning on any authorizer (RBAC, ABAC, etc) addresses this for OSS, etune needs to work on it for GKE and etune to fix the ABAC example to use groups.
        *   Way to not have service account (or not have a token).  deads: I don't think this bit is a P1 objective for 1.5.
    1.  Credential rotation - P3
        *   Need to document that rotation is possible.
            *   Might need flag on apiserver to allow multiple SA signer certs
        *   Needs more thought on exactly what threats and what keys are involved.
        *   Why: So if someone gets a service account token, we can revoke it.
            *   Turn on that one admission controller (or authorizer option) by default for standard Kubernetes turnups so it can do Service Account Token revocation.
        *   Why: so if someone takes master's key, we can revoke it (assuming root CA still secure).
        *   Consider letting the JWT verifier to take multiple verification keys
    1.  Eliminate unsecured master port  (issue [13598](https://github.com/kubernetes/kubernetes/issues/13598), possible PR [31491](https://github.com/kubernetes/kubernetes/pull/31491)) - P1 (liggitt)
        *   This should work regardless of any other authorization config
    1.  Add authn/authz to kubelet API (issue [11816](https://github.com/kubernetes/kubernetes/issues/11816), starting PR [31562](https://github.com/kubernetes/kubernetes/pull/31562)) - P1 (liggitt)
    1.  (ericchiang): continue to iron out kubectl login PMaybe (deads2k/red hat can review)
*   Keystone auth integration (liggitt) 20min
    1.  https://github.com/kubernetes/kubernetes/pull/25391
    1.  user mapping
        *   User friendly? Keystone <user name>@<domain name>
        *   Unique and programmatic? Keystone user id
    1.  project/role mapping
        *   Could map keystone project/role tuple to groups.  Suggested by deads, liggitt, and marc boorshtein.  It allows clean layering in kube authentication and authorization layers.
    1.  Kubernetes gets a token
        *   Identifies the keystone user
        *   Is scoped to a specific keystone project and role tuple (or roles?)
    1.  Will arrange two meetings, one for getting to know keystone, one for how to map keystone concepts to kube concepts.
    1.  Not discussed: fate of current password auth? Single-domain? Multi-domain?


## August 24th, 2016, 11a - Noon (Pacific Time)

Cancelled.


## August 10th, 2016, 11a - Noon (Pacific Time)



*   Group for x509 clients certs through Organizations or Organizational Units
    1.  Use organizations for parity with OpenShift
    1.  https://github.com/kubernetes/kubernetes/issues/30260
    1.  Interactions with groupification?
        *   Can you say "no more groups can be added?"
    1.  Interactions with impersonation?
    1.  The more we embed in certs, the more we need revocation (https://github.com/kubernetes/kubernetes/pull/29293)
    1.  AI: ericchiang to open PR
*   Node CSR API (liggitt, mikedanese, gtank)
    1.  https://github.com/kubernetes/kubernetes/pull/30153
    1.  What kind of rules do we need to have a safe-ish system
    1.  Goals:
        *   Partitioned node API permissions
        *   Sufficient info to approve node client certs for shared secret bootstrap
        *   Sufficient info to approve node serving certs for bootstrap
    1.  API-enforced info
        *   Who requested (user/uid/groups)
    1.  User-specified info
        *   What they're requesting (CSR content)
        *   What type of cert they're requesting (ExtKeyUsageServerAuth, ExtKeyUsageClientAuth, KeyUsageCertSign, etc)
    1.  Bootstrap cases
        *   Nodes starting from shared secret, getting individual client certs
            *   Limits usage to client cert
            *   Require specific CSR subject shape (e.g. "O=system:nodes,CN=system:node:<nodename>")
            *   Controller auto-approval can be based on a bootstrap group or user flag (allows a shared bootstrap secret to fan out to node client certs)
        *   Auto-approver for node serving certs
            *   Limits usage to server cert
            *   Requires specific CSR subject/SAN shape (CN only, CN must be in SANs?)
            *   Gates on request coming from the node user? Daisy-chains with node client cert auto-approver, but allows more security when individual identities can be distributed to nodes at provision time
    1.  Notes:
        *   Nodes need to be authorized differently based on the particular node requesting access.
            *   a node should only see pods scheduled to it
            *   a node should only see secrets used by pods scheduled on it
            *   a node should only be able to update itself
            *   etc
        *   To do that, you must have authentication provide information that can
            *   indicate that the subject is a node: special group?
            *   indicate that a node subject corresponds to a _particular_ node
            *   Names to match Jordan's suggestion.
        *   Server certs and client certs should be seen as different actions.
        *   Does the PR generate a serving cert rather than a client cert?  PR needs to be updated, it's touching serving certs.
        *   gtank? Inspect the CSR profile to see if we can make it reasonable for API consumers to express their usage intent regardless of profile names.
            *   Any controller that can decide whether something can be signed needs to be aware of who can request what (requires API knowledge).
            *   liggitt, mikedanese, deads2k want this to be based on API intent. Not sure how much CSR policy overlaps with these use cases.
*   Extend impersonation (david)
    1.  Allow complete impersonation including groups and extra (last bits of user.Info interface)
    1.  Needed for proxy/delegation cases (federated clusters) where auth method is non-transitive (like x509 client cert auth)
    1.  Allows authenticating at the edge, forwarding requests as the determined user (gated on the forwarding user's ability to impersonate)
    1.  Interaction with groupification? Proposal: impersonated data replaces
*   ServiceAccount changes? (clayton/david)
    1.  Don't want to have the tokens in secrets?
    1.  Because service accounts store credentials in secrets anyone who can view secrets can grab the credentials and act on their behalf.
    1.  Possibly move service account tokens to their own resource
        *   Maybe a specific solution to a general problem. Do we need a permission system for secrets or general objects?
    1.  Secrets as subresource on serviceaccount.  Only serviceaccounts can access secrets.  User would have to be able to impersonate the service account to get at it.
    1.  https://github.com/kubernetes/kubernetes/issues/16779
    1.  https://github.com/kubernetes/kubernetes/issues/11070
*   Image policy admission
    1.  Default deny (please)
    1.  Tag to immutable mapping?  Image pull policy?  Can we do it without an immutable image ID?
    1.  Particular registries don't have to have immutable image IDs to make decision.


## July 27th, 2016, 11a - Noon (Pacific Time)



*   `kubectl login` (https://github.com/kubernetes/kubernetes/pull/29350)
    1.  "discovery" mechanism
    1.  Possibly don't want to add new paths to the API server.
    1.  Would be preferable if we could use more standard discovery mechanisms.
    1.  We can start kubectl login without discovery mechanism
*   ServiceAccount changes?  (clayton/david)
    1.  Don't want to have the tokens in secrets?
    1.  Because service accounts store credentials in secrets anyone who can view secrets can grab the credentials and act on their behalf.
    1.  Possibly move service account tokens to their own resource
        *   Maybe a specific solution to a general problem. Do we need a permission system for secrets or general objects?
    1.  Secrets as subresource on serviceaccount.  Only serviceaccounts can access secrets.  User would have to be able to impersonate the service account to get at it. \

*   https://github.com/kubernetes/kubernetes/pull/27336 Add a generic admission controller that uses JSONPath for rules
    1.  Concerned about having to restart the server to have the changes take effect
    1.  Concerned about not having different rules per-subject
*   Track resource creator? (david)  https://github.com/kubernetes/kubernetes/pull/13925
    1.  This has been requested a lot.  Some controllers want to know who created something initially.  Could also be used for special authorization rules that allow creators/owners access (mentioned in SA/secret item)
    1.  This feature seems questionable.  In a world with mutable resources tracking just a creator isn't sufficient to make ACL decisions.
    1.  **Clayton**: any other reason to reconsider?
*   Current work
    1.  David
        *   subject access review from openshift to kube (needs api review)
        *   Enabling webhook token endpoint as on by default (rolled back)
    1.  Eric Chiang:
        *   kubectl login proposal


## July 13th, 2016, 11a - Noon (Pacific Time)



*   Jordan to report on webhook vs native OpenStack auth options.
    1.  Authentication (https://github.com/kubernetes/kubernetes/pull/25391)
        *   Adds token auth, remote authn call
        *   Converting to webhook would lose project id and keystone roles
        *   It uses those for authz?  Put them into user.Info.Extra (maybe, but not first)
        *   First class plugin is reasonable
        *   Can we convert roles to groups?  Not immediately straightforward because roles are scoped to projects.
    1.  Authorization (https://github.com/kubernetes/kubernetes/pull/25624)
        *   On hold until we know how keystone roles are being mapped.
*   Set Goals for 1.4?
    1.  Groupification changes (bobbyrullo as I recall)? Picked up by Eric Chiang.  Not for 1.4
    1.  RBAC alpha (target 1.4) - Eric Chiang
        *   GCE test should run with it on
        *   Default set of roles
        *   Default ServiceAccount roles
    1.  Rbac up to beta level (target 1.5):
        *   On by default for auto-set up cluster
        *   API working
        *   Not mixing resource/non-resource rules
    1.  `kubectl login` Eric Chiang
*   (ericchiang): would like to discuss requirements for initial "kubectl login" command: https://github.com/kubernetes/kubernetes/pull/25758
    1.  Include basic auth challenge handling
        *   [Example challenge handler in OpenShift](https://github.com/openshift/origin/blob/master/pkg/cmd/util/tokencmd/)
            *   ChallengeHandler interface (basic, negotiate, and multi impls)
            *   Basic handler can use specified user/pass or prompt from stdin
            *   Negotiate handler can use gssapi libs if present (via dlopen)
*   ServiceAccount changes?  (clayton/david)
    1.  Don't want to have the tokens in secrets?
    1.  Because service accounts store credentials in secrets anyone who can view secrets can grab the credentials and act on their behalf.
    1.  Possibly move service account tokens to their own resource
        *   Maybe a specific solution to a general problem. Do we need a permission system for secrets or general objects?
    1.  Secrets as subresource on serviceaccount.  Only serviceaccounts can access secrets.  User would have to be able to impersonate the service account to get at it. \

*   https://github.com/kubernetes/kubernetes/pull/27336 Add a generic admission controller that uses JSONPath for rules
*   Track Creator?


## June 29, 2016, 11a - Noon (Pacific Time)



*   SIG business
    *   Jordan, David and Eric Chiang have volunteered to help lead.
*   Security - Auth/Design/API dev is the primary focus, but this group will keep an eye to vulnerability management until there is enough more interest in a more vulnerability focused group.
*   RBAC summary and next steps?
    *   API types and authorizer is in
    *   Basic documentation is in
    *   Next steps:
        *   API support for intent-driven operations (bind user to role, remove user from role, etc)?
        *   CLI tools?
*   OpenStack: standalone plugin vs webhook approach, 1.4 direction
    *   Open PRs:
        *   https://github.com/kubernetes/kubernetes/pull/25624
        *   https://github.com/kubernetes/kubernetes/pull/25391
        *   https://github.com/kubernetes/kubernetes/pull/25536
    *   AI: Jordan to review open PRs, discuss webhook vs native with authors, summarize for next SIG
*   Field level policy: https://github.com/kubernetes/kubernetes/pull/27330
    *   AI: Interested members comment on the proposal
*   GCP
*   Kubectl login: https://github.com/kubernetes/kubernetes/pull/25758


## June 15, 2016, 11a - Noon (Pacific Time)

Cancelled


## June 1, 2016, 11a - Noon (Pacific Time)

OpenStack Identity and Kubernetes (https://github.com/kubernetes/kubernetes/issues/25066) 15m.


## May 18, 2016, 11a - Noon (Pacific Time)

Cancelled due to upcoming release


## May 5, 2016, 11a - Noon (Pacific Time)

Group handling (https://github.com/kubernetes/kubernetes/issues/23720) - bobbyrullo

Service Account tokens for ubernetes: deads2k

Kubectl auth / refresh tokens: liggitt

Field level authorization (impacted by cascading delete): deads2k

Auto-loading of ABAC: etune

Suggestion of future agenda topics: everyone


## April 20, 2016, 11a - Noon (Pacific Time)

First meeting, introductions (5m)

Agenda: etune (5m)

Direction: etune(5m)

Suggestion of future agenda topics: everyone (5m)

Scopes vs. Extra on user.Info: deads2k (5m)

Field level authorization (impacted by cascading delete): deads2k (5m)

Service Account tokens for ubernetes: deads2k (5m)

Auto-loading of ABAC: etune (5m)

Kubectl auth / refresh tokens: liggitt (5m)

Intros

Eric Tune's view of priorities:

- P1:  A default authorization system upstream, but allowing others

- P2:  Address Image Security and Pod Security Policy

- P3: System Roles and Impersonation

deads2k plans to upstream just the subjectAccessReview parts of Openshift Auth.

ericchiang offered to help upstream the policy parts, and deads2k offered to review and consult.

Suggested future agenda topics:

- Group handling (https://github.com/kubernetes/kubernetes/issues/23720) - bobbyrullo

Scopes vs. Extra on user.Info: deads2k

- long discussion

- RedHat wants to have a server that gives out oauth tokens with scopes on them that say

  the token was issued by user U, but the token is scoped to further restrict it to only do certain things.

- erictune asked why the user does not create a separate service account for this more narrowly scoped use case.

- Clayton said some of their customers do not want to use service accounts, for some reason.

- Someone (Jordan or David?) said it was easier to revoke  user U's authentication, and then all their tokens are revoked too.

- We decided to allow the Extra info, rather than having a structured place to put scopes (which would then be the responsibility of kubernetes to manage the scopes and map to REST operations).  Instead, individual authenticator and authorizer modules have to agree on how to interpret the Extra info.  Suggest map with namespaced keys.

Field level authorization (impacted by cascading delete): deads2k

deads2k asked how we would handle cascading delete.  If someone has permission to set the parent label of an object, but not delete the object, they can make the parent delete the object.

Or if they have permission to delete a parent object, they can make the controller delete children even if they don't have permission to delete the children (or not if we use Impersonate-User for the controller).

- We didn't get to finish this discussion.

Kubectl auth / refresh tokens: liggitt (5m)

- Long discussion

- Question about whether apiserver handling client secret is okay

- I'd like someone from community to summarize.

deads2k's perspective:

I think that we had discussions in three areas.



1.  Any login command must be flexible enough to handle the different sorts of authentication providers we already have along with other likely standards.  This means it should have plug-points for things like challenges (basic and kerberos use these).  It should not be tied to a particular provider.  I think we had agreement on this point. \

1.  The API server should not be in business of providing authentication endpoints.  This was only tangentially touched.  We started and got distracted by item 3, so I don't know whether we agree or disagree.  In my mind this is related to item 1.  Core kube should be agnostic, but some sort of discovery or structured returns is probably necessary. \

1.  Automatically attaching client secrets to refresh tokens presented to the API server.  The current design has the API server attaching his own client secret to any refresh token submitted to it.  This allows the token submitter to get back an access token without having the client secret to use the refresh token. It changes the oauth spec's requriement for two pieces of information (refresh token and client secret)  to exchange a refresh token for an access token to a system requiring a single piece of information (refresh token). \
 \
This was the big one.  This takes us from a system where confidential clients are confidential, to a system where a confidential client is blindly attaching his secret to any request saying, "yep, that's me".  This is concerns me a lot.  @erictune has an google OIDC expert for a another opinion.  I think I'd be less concerned if the server behaving this way wasn't our API server, but I'd still like to hear from another expert about whether it's a good idea or not.

Deferred for next meeting:

Field level authorization (impacted by cascading delete): deads2k

Service Account tokens for ubernetes: deads2k (5m)

Auto-loading of ABAC: etune (5m)
