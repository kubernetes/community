# Extending Kubernetes
Note Taker: Clayton Coleman ([smarterclayton](https://github.com/smarterclayton))

* Questions

    * Do we have enough extension mechanisms?

        * See below

        * Implementing network injection that isn’t a CNI injection of some form is hard

            * e.g. adding in arbitrary network devices, for example

    * Are Flex Volumes enough?

        * Maybe?

    * Are we doing ok on kubectl extensions?

        * Yes, we’re heading in the right direction with plugins

        * Kubectl itself should be developed using its own mechanisms

        * Extension points:

            * OpenAPI metadata (operate on object w/o knowing about it)

            * Subresources (generic API-level interfaces for certain facets)

            * Plugins (git-style)

        * Can we do custom validation in kubectl?

            * Do it via admission webhooks (beta in 1.9)

            * Can run validation asynchronously, and report it (put a condition in)

            * Client-side validation is iffy

    * Should we have phased hooks?

    * Should more complex application lifecycle layer on top?

        * Are we consistent enough across our hooks to enable users to build something sane?

    * Can we do referential integrity

        * Technically, with a webhook

        * Generally not the best idea (causes issues with components that don’t expect it)

        * Can maybe do with initializers on a namespace

        * Generally go for eventual consistency (e.g. wait for the secret to exist before starting the pod)

        * Reasons

            * Surface errors to users synchronously

                * Do it on the client-side

            * Avoid dealing with the full matrix of failure modes

    * How do we not re-invent a service mesh with auth webhooks?

        * It’s an open question

    * Can we do conversions on CRDs

        * Maybe with a webhook (ongoing discussion)

        * Why aren’t Custom API servers easy enough

            * Need OpenShift certificate signing mechanisms in Kube, or similar (also exists similarly in Istio)

            * Storage

                * Re-use existing etcd

                * Use your own etcd

                * Planned backing custom API servers with CRD storage

    * Why aren’t they more uniform

        * Because they came at different times

        * It’s hard to fix them now (need more versioning)

        * Different layers have different needs

            * Declarative is API

            * Below Kubelet is gRPC

            * gRPC for KMS, CSI

            * CNI is shell execution (mainly for legacy reasons)

    * Can we make custom controllers easiers?

        * OperatorKit

        * Need better docs (being worked on)

        * Untyped vs generated typed informers and listers

            * No Go generics exists

            * Can use interfaces

            * Can use conversions (may be harder than it needs to be)

            * Can wrap in a generic type (e.g. RBAC types)

        * Generating clients for CRDs (look for stts’s blog post on generators and informers)

* Existing extension mechanisms

    * API Extensions

        * External APIs (Aggregated APIs)

        * Custom Resources

    * Admission/Early-Phase ext points

        * Initializers

            * Can’t do in updates

        * Webhooks

            * Need to make them easier to implement (and a good example)

        * Pod Presets (ish)

        * Webhooks have to be fast, initializers are more async (normal controller style, with restraints)

    * Finalizers (late-phase ext points)

    * Flex Volumes

        * Being used to prototype container identity work, too

    * CSI (soon™)

    * (grpc) CRI

    * (binaries) CNI

    * (webhook) Auth plugins

        * Authz

        * Authn

        * Groups from authn hook ???

    * (API server) Custom metrics API servers

    * (grpc) Device plugins

    * (grpc) KMS (Key Management) 

    * (git style) kubectl plugins

    * (API usage) Any controller

    * External Cloud Provider

    * Pod lifecycle isn’t extensible

        * Object phase

        * Init containers (do exist)

        * defer containers

        * Lee’s debug container 

    * Logging plugins

        * Was a proposal for logging volumes, didn’t get follow-up

* New Extension type things

    * Enumerations in the API

    * Conditions (new condition types)

        * Loose coordination between multiple controllers

        * Signal to end users

        * Exists until formalized as a field

    * runc hooks

        * Could be a CRI concern

    * Optional CRI sub-interfaces?

    * Identity injection (custom CA certs in every pod, etc)

    * Simplifying the creation of controllers with controller generation ie. metacontroller: [https://github.com/GoogleCloudPlatform/kube-metacontroller](https://github.com/GoogleCloudPlatform/kube-metacontroller)

    * API server builder: [https://github.com/kubernetes-incubator/apiserver-builder](https://github.com/kubernetes-incubator/apiserver-builder)

    * Operator pattern toolkits:

        * Rook Team: [https://github.com/rook/operator-kit](https://github.com/rook/operator-kit) 

        * GiantSwarm: [https://github.com/giantswarm/operatorkit](https://github.com/giantswarm/operatorkit)

