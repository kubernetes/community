# WG AI Gateway Charter

This charter adheres to the conventions described in the [Kubernetes Charter
README] and uses the Roles and Organization Management outlined in
[wg-governance].

[wg-governance]:https://github.com/kubernetes/community/blob/master/committee-steering/governance/wg-governance.md
[Kubernetes Charter README]:https://github.com/kubernetes/community/blob/master/committee-steering/governance/README.md

## Background

We’ve seen large growth in the number of “AI Gateways” that have been launched
in the last couple of years which deploy and operate on Kubernetes, often
utilizing Gateway API. This WG aims to determine if the relevant features have
staying power and will be commonly useful to users for years to come, and if we
should expand the Kubernetes standards around this.

In SIG Network we have the Gateway API Inference Extension (GIE) project. The
GIE currently is paired with a Gateway and “schedules” routes according to
capabilities and metrics advertised by model serving platforms. For the purposes
of this document we’ll call this the “model serving use case”, as this currently
mainly covers the use case where models are being hosted on Kubernetes. There
are deployment situations where users won’t host models but still use a Gateway
to control access to 3rd party services (e.g. Gemini, OpenAI, Mistral, Claude,
etc), we’ll call this the “egress use case”. We find that in both the model
serving and egress use cases users want to be able to add more advanced filters,
policies and other plugins that control or modify inference requests.

However, there are many features we haven’t fully explored yet that seem to be
cleanly addable at the HTTPRoute level via filters or policies. Perhaps some
would even be applicable at the Gateway level. For example, it is conceivable
you might add a “semantic routing” at the HTTPRoute level as a filter to
determine which model to route to before the “routing/scheduling” layer. Or
perhaps you need a policy to rate-limit token usage for requests (maybe this
could even apply at the Gateway level). For the purposes of this charter,
we’ll refer to features at this level as “AI Gateway” features.

## Scope

The scope of this WG is to define terms like "AI Gateway" in the context of
Kubernetes and propose deliverables that need to be adopted in order to **manage
AI traffic** on Kubernetes, such as:

* **Prompt Guards** - Define and enforce content safety rules for inference
  content to detect and block sensitive or malicious prompts.
* **Token Rate Limiting** - enforce rate limiting rules based on token usage to
  control usage and cost.
* **Semantic Routing** - making a routing decision for an inference request
  based on semantic similarity of the request body.
* **Semantic Caching** - Provide caching for inference response based on the
  semantic similarity of prompts.
* **Response Risk** - Define and enforce content safety rules with inference
  response content to detect and block sensitive responses from generative AI
  models.
* **Failure Modes** - How inference routing failures should be handled, what
  failure modes we think are important to cover. For instance this may
  encapsulate fallback and retry policies.
* **Observability** - Evaluate mechanisms for observability for “AI Gateways”
  and if there are AI Gateway specific features needed, make suggestions
  according to existing tools.

> **Note**: The above list of features should be considered an example, and
> non-exhaustive. We may not act on all of these, but the purpose is more to
> illustrate the kind of features we will be exploring.

Across features that are explored by this WG, we will also explore the
application of these features to multi-cluster use cases and provide support
for multi-cluster deployment scenarios.

### In Scope

Overall guidance for the WG is to control scope as much as is feasible. The WG
will support model serving via AI networking and traffic management features
(but not working on model serving itself, unless in conjunction with WG
Serving). In particular, the following is in scope:

* Providing definitions for networking related AI terms in a Kubernetes
  context, such as "AI Gateway".

* Defining important use-cases for Kubernetes users, including both single and
  multi-cluster use cases.

* Determining which common features and capabilities in the "AI Gateway" space
  need to be covered by Kubernetes standards and APIs according to user and
  implementation needs.

* Creating proposals for "AI Gateway" features and capabilities to the
  appropriate sub-projects.

* Propose new sub-projects if existing sub-projects are not sufficient.

### Out of Scope

* Developing whole "AI Gateway" solutions. This group will focus on enabling
  existing and new solutions to be more easily deployed and managed on
  Kubernetes, not creating any new Gateways.

* Any specific kind of hardware support is generally out of scope.

* This group will not cover the entire spectrum of networking for AI. For
  instance: RDMA networks are generally out of scope.

* While we serve the "model serving use case", and important distinction is
  that working directly on Model serving, and AI workloads themselves are
  not in scope unless done in collaboration with WG Serving (see below for a
  more complete explanation about this nuance).

* While we may look into ways in which observability may be tailored
  specifically for AI Gateways, We may make suggestions to groups like
  OpenTelemetry, but we are not going to develop new standards for this as part
  of this working group.

### Additional Scope Distinctions

There is a subtle distinction to be made when it comes to the scope of this WG
for load-balancing and routing inference, particular when dealing with inference
_workloads_: When the use case includes local model serving on the cluster, and
routing and load-balancing features _rely on information from the inference
workloads_, this kind of routing falls under the scope of WG Serving.

A good example of this is the [Gateway API Inference Extension (GIE)][gie].
This project came from WG Serving and specifically handles advanced routing and
load-balancing for inference which is informed by metrics and capabilities being
advertised by the model serving platform (e.g. VLLM). In this vein, the GIE is
effectively an alternative to the Kubernetes `Service` API, whereas this WG
means to operate more at the `Gateway` and `HTTPRoute` level.

Use cases which have to interact with the model serving layer for networking
(as described above) are generally out of scope for this WG. If some feature
the WG is working on absolutely must cross this line, the effort MUST be brought
to WG Serving and worked on as a joint effort with them.

[gie]:https://github.com/kubernetes-sigs/gateway-api-inference-extension

## Deliverables

* A compendium of AI related networking definitions (e.g. "AI Gateway") and
  key use-cases for Kubernetes users.

* Provide a space for collaboration and experimentation to determine the most
  viable features and capabilities that Kubernetes should support. If there is
  strong consensus on any particular ideas, the WG will facilitate and
  coordinate the delivery of proposals in the appropriate areas.

## Stakeholders

* SIG Network
* SIG MultiCluster

### Related WGs

* WG Serving - The domain of WG Serving is AI Workloads, which can be served by
  some of the networking support we want to add. When we have proposals that
  are strongly relevant to serving, we will loop them in so they can provide
  feedback.

## Roles and Organization Management

This working group adheres to the Roles and Organization Management outlined in
[wg-governance] and opts-in to updates and modifications to [wg-governance].

[wg-governance]:https://github.com/kubernetes/community/blob/master/committee-steering/governance/wg-governance.md

## Exit Criteria

The WG is done when its deliverables are complete, according to the defined
scope and a list of key use cases and features agreed upon by the group.

Ideally we want the lifecycle of the WG to go something like this:

1. Determine definitions and key use cases for Kubernetes users and
   implementations, and document those.
2. Determine a list of key features that Kubernetes needs to best support the
   defined use cases.
3. For each feature in that list, make proposals which support them to the
   appropriate sub-projects OR propose new sub-projects if deemed necessary.
4. Once the feature list is complete, leave behind some guidance and best
   practices for future implementations and then exit.
