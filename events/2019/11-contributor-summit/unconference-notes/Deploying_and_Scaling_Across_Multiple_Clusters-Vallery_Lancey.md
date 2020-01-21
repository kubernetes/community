#### Presenter:
 Vallery Lancey

#### Topic:
 Unconference: Deploying and Scaling across Multiple Clsuters

#### Date & time:
 11/18/2019 9:30am

#### Notes

Getting started: what is multicluster? 
* Federation / copy across otherwise independent clusters
* Deployment: rollout changes
  - Hard w/o big CI/CD changes
  - Granularity of rollouts is hard, if you have two clusters how do you launch less
    than 50% at a time
  - Non-native k8s apps often don’t deal with with partial rollout.
* Big variety of use cases & solutions, custom implementations everywhere

What is definition of multicluster? (vs just having a bunch of clusters)
* Deploy same app in multiple clusters

What are reasons to run?
* Big range, very customer-dependent
* Range from east-coast/west coast clusters to a cluster in every cell tower
* Different fault domains
* Avoid canary deployments
* Backup environment for rollbacks

Need tools above deployment layer, service meshes require committing to fearsome
stack (eg Istio)
* Example: how to generate Flagger to multicluster (Weaveworks incremental rollout
  tool based on linkerd)
* Is there load balancing across clusters?
* Is the system aware of multiple clusters--basic things like how many clusters
  to adjust rollout #s

Why multi-cluster?
* Technical reasons (5k node limit, etc) vs customer/user need
* User needs include
  - CI/CD
  - High availability
  - Blast radius (or is this a technical detail?)
  - Isolated environments
  - Data locality/federal regulations
  - Disaster recovery

Service meshes are an additional step for customers to migrate existing applications
to, but federation is needed now. This leads to custom solutions.

What are multi-cluster building blocks?
* Deployment
* Service  mesh. But this is too big and meshes are too opinionated, should be
broken into layers
  - Service discovery
  - Connectivity
  - Ingress
  - Routing
  - Signals: alerts/health/observability

Fundamental problems that define multicluster:
* Resource usage / technical limitations of data centers
* Administrative groups: independently run clusters that use common organizational
  security policies
* Coordination across clusters
* Deployment
* Rollout new app
* Cascading triggers (service X launches which unblocks service Y to launch etc)


#### Key learnings / takeaways

Multi-cluster is being used today, but in ad-hoc and specific ways. Common usage
patterns need to be discovered and standardized in order to learn how k8s needs
to be extended to support multi-cluster well.


#### Action items
* Survey on customer needs
  - Rather than current usage, as that can be driven by legacy accidents
* Start defining building block/concepts
* Actually we also need to survey current solutions in order to identify building blocks
* Focus on defining minimal set of building blocks to avoid the usual k8s explosion.
