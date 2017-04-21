# SLO: Kubernetes cluster of size at least X is able to start Y Pods in Z minutes

## Burst Pod Startup Throughput SLO
### Summary
Kubernetes cluster of size at least X is able to start Y Pods in Z minutes.
### User Stories
- User is running a workload of X total pods and wants to ensure that it can be started in Y time. 
- User wants to understand how quickly they can react to a dramatic change in workload profile - e.g. they need to begin scheduling new pods when QPS meets Z% of maximum throughput of users system - to accommodate a spike due to a Black Friday Sale.
- User is running a huge serving app on a huge cluster. He wants to know how quickly he can recreate his whole setup in case of a serious disaster which will bring the whole cluster down.

Current steady state SLOs are do not provide enough data to make these assessments about burst behavior. 
## SLO definition (full)
### Test setup
Standard performance test kubernetes setup.
### Test scenario is following:
- Start with a healthy (all nodes ready, all cluster addons already running) cluster with N (>0) running pause Pods/Node.
- Create smallest number of Deployments necessary to create X Pods, with one Deployment running at most 5000 Pods.
- Each Deployment runs in it's own Namespace.
- Each Deployment is covered by a single Service.
- Each Pod in any Deployment contains two pause containers, one secret other than ServiceAccount and one ConfigMap, has resource request set and doesn't use any advanced scheduling features (Affinities, etc.) or init containers.
- Measure the time between starting the test and moment when user-side watch saw system reports all Pods Ready. Note that pause container is ready just after it's started, which may not be true for more complex containers that use nontrivial readiness probes.
### Definition
Kubernetes cluster of size at least X adhering to the environment definition, when running the specified test, 99th percentile of time necessary to start Y pods from the time when user created all controllers to the time when user-side watch observes all Pod as Ready is no greater than Z minutes, assuming that all images are already present on all Nodes.