# Federated ResourceQuota Requirements & Design
@weijinxu, @IrfanUrRehman, @JianhuiZ
reviewed by  @quinton-hoole, @deepak-vijSummary

In a single Kubernetes cluster there is a concept of Resource Quota that manages the resources by namespace so that user usages can be limited to their fair share of the resources. We want to expand this concept to the federation level so that users can treat Federation Control Plane as a “Single Pane of Glass” for seamlessly managing the Resource Quota  information across all the underlying federated clusters.

## User scenario

As a federation user and an admin of clusters, I want to be able to define an overall resource quota for some specific or all resources for a given federated namespace. The overall resource quota limit should be applicable across clusters and preferably enforced at federation control plane level.

For example, let us say that we’ve specified 100 pods in the resource quota spec in the federation namespace, with multiple underlying federated clusters. Let us assume that 50 pods have been deployed already across the federation in a previous deployment run. In a subsequent deployment run, if a user requests additional 50+ more pods to be deployed, the deployment will be rejected as a whole by the federation control plane before deployment process actually starts. If user requests less than 50 pods, the deployment should succeed as per the deployment preferences.

## Implementation details
### Create Local Resource Quota in underlying clusters

A straightforward thought is to divide and propagate the federation quota value to the underlying clusters. This kind of distribution may cause divided quota gets too small share which hardly meet users requests, and may end up in a state that a user may have quite a few resources available but not in the cluster he/she preferred. What makes this worse is that the user wants to deploy a big workload but none of the clusters have enough resource quota value left anymore. This means that we may need to somehow adjust the allocated quota of the underlying clusters accordingly. Similarly, this is also true for clusters joining and leaving the federation.
Our proposal is that to not split the quota value but instead give every underlying cluster the entire federation quota. When a user creates Federation Resource Quota, Federation Resource Quota controller will propagate the federation resource quota value to each federated cluster. At the same time, the Federation Resource Quota controller then keeps monitoring underlying Local Resource Quota information and updates the Federation Resource Quota status with already used/hard resource information to be subsequently used by the Federation Resource Quota Admission Controller.
### Additional Quota Controller Functionality
Above proposal allows users to consume resource at their will freely. But user can go to underlying cluster directly and as each cluster has reserved quota amount same as Federation Resource Quota, they can be consumed more than quota limits if multiple deployments are executed against more than one underlying cluster.
The proposal is that to give every underlying cluster the federation quota initially, and then ensure every cluster’s quota to `LRQ.Used + FRQ.Remaining` on each reconciliation. Which means that given any time, any cluster has a remaining quota the same as the federation remaining quota. This also simplifies cluster joining (assign it federation remain quota) and cluster leaving (nothing needed). And, when the user bypass the federation and deploy some workload to the underlying clusters, the federation quota controller should update the quotas of other clusters accordingly.
#### NOTE: This method can prevent over consumption for users accessing underlying local clusters directly, if they only work against local cluster individually. This can’t prevent the over consumption if they deploy to multiple local clusters parallelly.
### Consider this example:
1. Having 3 clusters A, B and C
2. Assigned user (namespace) a quota of 10 pods
3. Federation quota controller set quota of cluster A to 10, B to 10 and C to 10
4. User deploys a replicaset of 5 where 3 go to cluster A and 2 go to cluster B; user now has a remaining quota of 5 pods
5. Federation quota controller set quota of cluster A, B and C to:<br>
    A: 8 (3+5); B: 7 (2+5);  C: 5 (0+5)
6. Cluster D joins, federation quota controller sets quota of cluster D to 5 (0+5)
7. Cluster B leaves, federation quota controller does nothing.
8. 2 replicas in cluster B rescheduled to cluster C (1 replica) and D (1 replica)
9. Federation quota controller updates quota to:<br>
    A: 8 (3+5); C: 6 (1+5); D: 6 (1+5)
10. User creates another deployment with 10 pods to cluster A (5) and C (5)
11. Federation admission controller captures it and see the request amount (10) is more than remaining quota (5), so refused the request.
12. User creates another deployment with 2 pods to cluster A (1) and C (1)
13. Federation quota controller updates quota to:<br>
    A: 7 (4+3); C: 5 (2+3); D: 4 (1+3)
14. User bypasses federation and deploy 1 pod to cluster D
15. Federation quota controller catches this event and updates quotas to:<br>
    A: 6 (4+2); C: 4 (2+2); D: 4 (2+2)
16. User bypasses federation and deploy 1 pods to cluster C
17. Federation quota controller catches this event and updates quotas to:<br>
    A: 5 (4+1); C: 4 (3+1); D: 3 (2+1)
18. User bypasses federation and deploy 1 pod to cluster C, and 1 pod to cluster D simultaneously. Since both clusters have 1 remaining quota which satisfy the request, both pods are deployed successfully.
19. Federation quota controller catches this event and updates quotas to:<br>
    A: 4 (4+0); C: 4 (4+0); D: 3 (3+0)
20. The total used quota count is 11, which is 1 more than the original quota created in Federation Resource Quota object. This is the result of step 18 when user sent deployment request to different clusters parallelly. But since the remaining quota has been trimmed to 0, no more deployment can be executed either from federation level or through local cluster directly.

### Checking admit policy on Federation Control Plane (FCP)
When user sends requests such as deployment to FCP API, it should check the quota validity first before proceeding. The implementation in Kubernetes clusters is through admission control plugin. We will need to do the same for FCP. The validity algorithm is basically checking if the new request can fit its resource requirement into the remaining quota value, which is the difference between the resource spec defined in resource quota object and the used resource amount.
### Delete resource quotas
Deleting a resource quota object deletes all local resource quotas from corresponding underlying clusters which remove all quota constraints for specified namespace. This should follow the existing logic for cascading deletion. Refer to https://github.com/kubernetes/kubernetes/issues/33612
