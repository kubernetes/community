## System throughput SLI/SLO details

### User stories
- As a user, I want a guarantee that my workload of X pods can be started
  within a given time
- As a user, I want to understand how quickly I can react to a dramatic
  change in workload profile when my workload exhibits very bursty behavior
  (e.g. shop during Back Friday Sale)
- As a user, I want a guarantee how quickly I can recreate the whole setup
  in case of a serious disaster which brings the whole cluster down.

### Test scenario
- Start with a healthy (all nodes ready, all cluster addons already running)
  cluster with N (>0) running pause pods per node.
- Create a number of `Namespaces` and a number of `Deployments` in each of them.
- All `Namespaces` should be isomorphic, possibly excluding last one which should
  run all pods that didn't fit in the previous ones.
- Single namespace should run 5000 `Pods` in the following configuration:
  - one big `Deployment` running ~1/3 of all `Pods` from this `namespace`
  - medium `Deployments`, each with 120 `Pods`, in total running ~1/3 of all
    `Pods` from this `namespace`
  - small `Deployment`, each with 10 `Pods`, in total running ~1/3 of all `Pods`
    from this `Namespace`
- Each `Deployment` should be covered by a single `Service`.
- Each `Pod` in any `Deployment` contains two pause containers, one `Secret`
  other than default `ServiceAccount` and one `ConfigMap`. Additionally it has
  resource requests set and doesn't use any advanced scheduling features or
  init containers.
