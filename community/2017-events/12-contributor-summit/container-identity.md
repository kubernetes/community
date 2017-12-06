Container identity

Lead: Greg Castle

How are people using k8s service accounts

* Default service accounts?

* User:

    * (Zalando)

        * 55 clusters

        * Only allow pre-set service accounts postgres "operator service accounts"

        * Don’t enable RBAC

        * Namespaces: usually use the default namespace, more sophisticated clients can use namespaces

        * Cluster per product

        * CRD that defines request for an OAuth2 token

            * Q: Would you want to tie it to a Kubernetes service account. E.g. annotation on a service account to get an OAuth2 token

    * (IBM)

        * Strict way of "cutting up cluster"

        * No team has access to the Kubernetes API

        * Humans do, pods don’t

        * Admins create service accounts and role bindings

        * Issues with pod authors in the namespace can run any service account.

            * Is there a missing binding to control the right to use a service account?

        * Namespace boundary

            * Application

    * (VMware)

        * 1 cluster

        * Projects are per-namespace (200 namespaces)

        * Wrote an admission controller between which pod is using which service account. Denies certain kind of changes.

        * Ownership model "I created this object, only I can change it"

        * Namespace but wanted sub-namespace primitives

        * Service accounts have labels, need to match the pod.

        * Q: Why not just different namespaces?

            * Don’t want to give users the ability to create namespaces

    * (Jetstat)

        * Dev, staging, and prod cluster

        * Namespace per team

        * Their developers don’t have permissions to create workloads

        * CI system is the thing that creates workloads, that controls stuff

    * (CoreOS)

        * Customers create namespaces that are controlled by CI/CD

        * Users have read-access to cluster

        * Central catalog controls the service account

            * Authoritative on RBAC

            * Catalog is currently controlled by CoreOS, can be used by users later

        * Namespaces

            * Catalog is enabled on a namespace

        * Also use service accounts to authenticate CI/CD systems

* Container identity

    * Authenticate workloads against each other

    * Does this fit into the current way you use service accounts?

* What about envoy istio?

    * Complementary efforts

    * Services can do service to service auth

    * Maybe not solving "I want to talk to s3, I want to talk to GitHub"

    * Give them a better way to provision x509 certs

        * Istio will be dependent on container identity

    * Istio and SPIFFE?

        * Istio is using SPIFFE conformant x509, not the workloads API

* Q: how does Zalando RBAC CRDs?

    * Trust two engineers that have the permissions to create the CRDs

    * Emergency procedure with time limited API

* Q: multiple clusters working together?

    * IBM: going to face it

    * CoreOS: thinking about users authenticating across multiple clusters through sharing public keys.

* Liggitt: namespaces are the boundaries of trust

    * Originally thought about using pod templates so admins could create specific templates that pods MUST use.

    * Reference pod template in RC or deployment instead of inlining them

        * Never bubbled up to be the most important thing yet

* Liggitt: Kubernetes API isn’t really designed for things that hold credentials

    * Moving the mounting and creation of the service account to the kubelet through flex volumes

    * Helps you leverage the node level integrations

* Q: what about the secrets proposal?

    * KMS is being added

    * No change for encrypting the secret as you read it

* Slides

    * Authorization should apply to the identity, rather than the namespace

    * Is multiple SA per namespace an anti-pattern?

        * Useful for bounding process within pods to certain permissions

        * Problem when a SA from cluster scoped manifests

* Q: kerberos and vault bindings implemented as flex volumes

    * Is it important to have standard identity mechanisms? (Yes)

        * Can we add it to conformance?

        * Avoid lock-in for IAM

    * Cluster identity also becomes an issue too

        * Pod/UID/Namespace/Cluster

* Potentially use OpenID Connect format so others can use that spec

