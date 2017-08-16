# SLO: Kubernetes cluster of size at least X is able to start Y Pods in Z minutes
**This is a WIP SLO doc - something that we want to meet, but we may not be there yet**

## Burst Pod Startup Throughput SLO
### User Stories
- User is running a workload of X total pods and wants to ensure that it can be started in Y time. 
- User is running a system that exhibits very bursty behavior (e.g. shop during Black Friday Sale) and wants to understand how quickly they can react to a dramatic change in workload profile.
- User is running a huge serving app on a huge cluster. He wants to know how quickly he can recreate his whole setup in case of a serious disaster which will bring the whole cluster down.

Current steady state SLOs are do not provide enough data to make these assessments about burst behavior. 
## SLO definition (full)
### Test setup
Standard performance test kubernetes setup, as describe in [the doc](../extending_slo.md#environment).
### Test scenario is following:
- Start with a healthy (all nodes ready, all cluster addons already running) cluster with N (>0) running pause Pods/Node.
- Create a number of Deployments that run X Pods and Namespaces necessary to create them.
- All namespaces should be isomorphic, possibly excluding last one which should run all Pods that didn't fit in the previous ones.
- Single Namespace should run at most 5000 Pods in the following configuration:
  - one big Deployment running 1/3 of all Pods from this Namespace (1667 for 5000 Pod Namespace)
  - medium Deployments, each of which is not running more than 120 Pods, running in total 1/3 of all Pods from this Namespace (14 Deployments with 119 Pods each for 5000 Pod Namespace)
  - small Deployments, each of which is not running more than 10 Pods, running in total 1/3 of all Pods from this Namespace (238 Deployments with 7 Pods each for 5000 Pod Namespace)
- Each Deployment is covered by a single Service.
- Each Pod in any Deployment contains two pause containers, one secret other than ServiceAccount and one ConfigMap, has resource request set and doesn't use any advanced scheduling features (Affinities, etc.) or init containers.
- Measure the time between starting the test and moment when last Pod is started according to it's Kubelet. Note that pause container is ready just after it's started, which may not be true for more complex containers that use nontrivial readiness probes.
### Definition
Kubernetes cluster of size at least X adhering to the environment definition, when running the specified test, 99th percentile of time necessary to start Y pods from the time when user created all controllers to the time when Kubelet starts the last Pod from the set is no greater than Z minutes, assuming that all images are already present on all Nodes.