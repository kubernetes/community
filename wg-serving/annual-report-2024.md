# 2024 Annual Report: WG Serving

## Current Initiatives and Project Health

### 1. What work did the WG do this year that should be highlighted?

#### Public talks

##### Talks that focus on WG Serving

KubeCon NA 2024 talks:

* [WG Serving: Accelerating AI/ML Inference Workloads on Kubernetes](https://youtu.be/HWoHcOTKJM4?si=ngPu5MCJ6vONxX0c)  

Kubernetes Podcast:

* [Kubernetes Working Group Serving](https://kubernetespodcast.com/episode/240-wg-serving/)

##### Other talks that mention WG Serving or initiatives from the WG:

* [Navigating Failures in Pods with Devices: Challenges and Solutions](https://youtu.be/-YCnOYTtVO8?si=ncm8YTOw8fM72UMd)  
* [Optimizing LLM Performance in Kubernetes with OpenTelemetry](https://youtu.be/6rdeFACyyYg?si=eexo_RrN4zcZBlpN)  
* [Solving the Kubernetes Networking API Rubik's Cube](https://youtu.be/S5QsqEb8wec?si=ZI9fXizMofGoA2MA)  
* [Distributed Multi-Node Model Inference Using the LeaderWorkerSet API](https://kccncna2024.sched.com/event/1i7rn)  
* [Engaging the KServe Community, the Impact of Integrating a Solutions with Standardized CNCF Projects](https://www.youtube.com/watch?v=S27wzRNsStU)  
* [Optimizing Load Balancing and Autoscaling for Large Language Model (LLM) Inference on Kubernetes](https://www.youtube.com/watch?v=TSEGAh1bs4A)  
* [Unlocking Potential of Large Models in Production](https://www.youtube.com/watch?v=-xEpzaIvor4)  
* [Best Practices for Deploying LLM Inference, RAG and Fine Tuning Pipelines](https://www.youtube.com/watch?v=EmGe_58524g)  
* [Advancing Cloud Native AI Innovation: Through Open Collaboration](https://www.youtube.com/watch?v=kG_wqP2CXUE)  
* [Kubernetes WG Device Management \- Advancing K8s Support for GPUs](https://sched.co/1hovp)  
* [Better Together\! GPU, TPU and NIC Topological Alignment with DRA](https://sched.co/1i7pv)  
* [Incremental GPU Slicing in Action](https://sched.co/1izuH)  
* [A Tale of 2 Drivers: GPU Configuration on the Fly Using DRA](https://sched.co/1i7lw)  
* [Which GPU Sharing Strategy Is Right for You? A Comprehensive Benchmark Study Using DRA](https://sched.co/1i7ol)

#### New subprojects and initiatives

##### [Inference Perf](https://github.com/kubernetes-sigs/inference-perf)

We created a new GenAI benchmarking tool called [inference-perf](https://github.com/kubernetes-sigs/inference-perf) to consolidate and standardize on a benchmarking-as-code tool which can be used to benchmark different Gen AI workloads running on Kubernetes or elsewhere in a model server and infrastructure agnostic way.

* The approved subproject proposal can be found [here](https://github.com/kubernetes-sigs/wg-serving/tree/main/proposals/013-inference-perf).  
* General design for the tool can be found [here](https://github.com/kubernetes-sigs/inference-perf/blob/main/docs/design.md).  
* We have had contributions from different companies so far like Google, Red Hat, IBM and Capital One. NVIDIA and others looking to contribute as well.  
* We have dependency management set up for the tool and it can be shipped as a python package \- [PR \#13](https://github.com/kubernetes-sigs/inference-perf/pull/13).   
* Github [actions](https://github.com/kubernetes-sigs/inference-perf/actions) are set up to run formatting, linting and other checks on PR submissions.  
* Support for a constant time load generator and load generator to send requests using a Poisson distribution in progress \- [PR \#9](https://github.com/kubernetes-sigs/inference-perf/pull/9).  
* Support for model server clients ([PR \#5](https://github.com/kubernetes-sigs/inference-perf/pull/5)) and metrics collection ([PR \#7](https://github.com/kubernetes-sigs/inference-perf/pull/7)) are in progress as well.

##### [Gateway API Inference Extension](https://github.com/kubernetes-sigs/gateway-api-inference-extension)

The Gateway Inference Extension(GIE) has been developing rapidly, with our first release ([v0.1.0](https://github.com/kubernetes-sigs/gateway-api-inference-extension/releases/tag/v0.1.0)) recently published. We have already seen\[15%-60%\] reduction in output token latency when kv cache is close to saturation in this first release. 

Development will continue rapidly, with a focus on: 

* Productionalization  
* Adoption of the latest developments (prefix caching as an example)  
* And driving improvement/development in other areas of inference routing (multi-LoRA, disaggregated serving pools, etc)

GIE has established adoption patterns for both the [Model Server](https://github.com/kubernetes-sigs/gateway-api-inference-extension/tree/main/docs/proposals/003-endpoint-picker-protocol#model-server-protocol), and the [Gateway](https://github.com/kubernetes-sigs/gateway-api-inference-extension/tree/main/docs/proposals/003-endpoint-picker-protocol#proxy-protocol) interfaces. With regards to model servers, GIE already integrates with [vLLM](https://github.com/vllm-project/vllm) well, and will soon support [JetStream](https://github.com/AI-Hypercomputer/JetStream) as well, other model servers need only implement the protocol and will also cleanly integrate into GIE. As a part of sig-net, GIE seeks to build a strong partnership with the gateways in the space. Integration efforts from these orgs are already underway:

* GKE  
* [KGateway](https://kgateway.dev/)   
* [Istio](https://istio.io/)

GIE was developed on top of [ext-proc](https://www.envoyproxy.io/docs/envoy/latest/configuration/http/http_filters/ext_proc_filter) and [Envoy Proxy](https://www.envoyproxy.io/), and so any Proxy that can support ext-proc, can support the GIE protocol.

The GIE team is looking forward to this upcoming year\! With many integrations upcoming and working with new partners (such as KServe) we look forward to where we are headed in 2025\.

##### [Serving Catalog](https://github.com/kubernetes-sigs/wg-serving/tree/main/serving-catalog)

Serving Catalog creates a repository of example K8s templates to deploy popular inference workloads. [Kustomized Blueprints - Serving Catalog](https://docs.google.com/document/d/1dOP0lIn-DK3tmq8gSvyK51J7CfKTz_Z1mqIAeYK8YS0/edit?tab=t.0#heading=h.48xpf8o6yz1j) described an approach for [LLM Serving Catalog](https://docs.google.com/document/d/1yXOhEXltc69_WFX90kMrHs3NcJGsD-U6-RejvK3XQtE/edit#heading=h.uxsw8fcrx2wp) using Kustomize overlays and components to provide a framework for extensible templates.

The current support matrix is available [here](https://github.com/kubernetes-sigs/wg-serving/blob/main/serving-catalog/catalog.md), including support for:

* [Single-host Inference using Deployments](https://github.com/kubernetes-sigs/wg-serving/tree/main/serving-catalog/core/deployment) for vLLM and JetStream  
* [Multi-host Inference using LeaderWorkerSet](https://github.com/kubernetes-sigs/wg-serving/tree/main/serving-catalog/core/lws) for vLLM  
* Components for [HPA stubs for token-latency](https://github.com/kubernetes-sigs/wg-serving/commit/54ee7234800cad53c0a43668b177c704ae704730)

#### Workstream Updates

**Orchestration**: Progress on various initiatives and relevant projects such as the GIE project, Serving Catalog, and [KServe](https://github.com/kserve/kserve). Please refer to the previous section for more details.

**Autoscaling**: Ongoing efforts to integrate custom metrics for autoscaling AI workloads.

One of the directions for autoscaling is a unification of model weights distributions formats. There is not a single distribution mechanism and WG Serving believes that the container images are the best distribution format. WG Serving identified some problems with the OCI large images support and sponsored the Kubernetes Image VolumeSource KEP work.

**Multi-Host Serving**: Improvements in distributed inference across nodes in vLLM, LeaderWorkerSet, and KServe.

LeaderWorkerSet (LWS) continues to evolve as a key component for multi-host inference, addressing the challenges of deploying large-scale AI/ML models across multiple nodes. The v0.3.0 release introduced subgroup support for disaggregated serving, a new start policy API, and improved inter-container communication through leader address injection. It also added a multi-node serving example for LLaMA 70B on GPUs using vLLM. Building on these capabilities, v0.4.0 & v0.5.0 introduced network configuration support, group size as an environment variable, and expanded multi-host inference examples, including llama.cpp for distributed inference and an updated vLLM example for Llama 3.1-405B. These enhancements reinforce LWS’s flexibility in orchestrating increasingly larger models on Kubernetes.

At the same time, WG-Serving is working closely with vLLM developers on the latest P\&D disaggregation feature progress, actively testing the upstream 1P1D functionality to better understand evolving orchestration requirements. This collaboration aims to drive improvements in xPyD capabilities, further unlocking disaggregated serving on Kubernetes by optimizing workload placement and execution strategies. By refining these mechanisms, we aim to enhance inference performance, ensuring more efficient resource utilization and scalability for large-scale AI workloads.

With these iterative improvements, LWS and vLLM continue to refine multi-host inference, making large-scale distributed model deployments on Kubernetes more reliable, efficient, and adaptable.

In addition, KServe also added multi-host serving capability via vLLM serving runtime.

**DRA (Dynamic Resource Allocation)**: Enhancing GPU/accelerator allocation, structured parameters, and resource claim standardization.

The DRA long term vision will enable many serving-related scenarios in future. In 2024, most of the effort was spent on adjusting plans and design to ensure timely GA of the feature and smooth migration from the device plugin architecture. We are working on prioritizing features in DRA needed for serving workloads, however the major push in the first half of 2025 will still be GA-related activities. WG Serving prepared a document listing scenarios and requirements for the DRA with the hope to start working on some of them in 2025\.

Another topic that is actively being discussed is device failure handling and managing workload affected by those failures.

### 2. Are there any areas and/or subprojects that your group needs help with (e.g. fewer than 2 active OWNERS)?

Not at this point.

## Operational

Operational tasks in [wg-governance.md]:

- [ ] [README.md] reviewed for accuracy and updated if needed
- [ ] WG leaders in [sigs.yaml] are accurate and active, and updated if needed
- [ ] Meeting notes and recordings for 2024 are linked from [README.md] and updated/uploaded if needed
- [ ] Updates provided to sponsoring SIGs in 2024
      - [$sig-name](https://git.k8s.io/community/$sig-id/)
        - links to email, meeting notes, slides, or recordings, etc
      - [$sig-name](https://git.k8s.io/community/$sig-id/)
        - links to email, meeting notes, slides, or recordings, etc

[wg-governance.md]: https://git.k8s.io/community/committee-steering/governance/wg-governance.md
[README.md]: https://git.k8s.io/community/wg-serving/README.md
[sigs.yaml]: https://git.k8s.io/community/sigs.yaml
