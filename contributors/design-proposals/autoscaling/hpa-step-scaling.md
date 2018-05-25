# HPA: Threshold Based Step Scaling Policy Proposal

## Target of HPA

The target of HPA is to dynamically adjusts the amount of allocated resources (pods) to meet changes in
workload demands. Basically, we evaluate a good HPA scaling policy as:

* Can achieve high resource utilization
* Responsive: scale to the desire number of replicas fast

## Problem

The Horizontal Pod Autoscaler (HPA) is implemented as a Kubernetes API resource and a controller.
Current HPA only implement “target tracking based scaling policy”: the controller periodically adjusts 
the number of replicas in a replication controller or deployment to match the observed average CPU 
utilization to the target specified by user.

In a nutshell, the target number of pods is calculated from the following formula:
```
TargetNumOfPods = ceil(sum(CurrentUtilization) / TargetUtilization)
```

Howeever, this kind of “target tracking based scaling policy” does NOT perform well when we want to maintain high 
resource utilization as well as want the HPA to scale fast, especailly when the workload deamds has spikes. 
To solve this problem, we propose to add a “threshold based step scaling policy” to HPA. We will discuss the
mechainsm of “threshold based step scaling policy” and the scenarios where it performs better than current
“target tracking based scaling policy” in the next two sections.


## Threshold Based Step Scaling

With step scaling policy, you can add one or more step adjustments that enable you to scale based on the
threshold breach. Each step adjustment specifies a lower bound for the metric value, an upper bound for 
the metric value, and the amount by which to scale, based on the scaling adjustment type. 

The HPA controller periodically pull utilization information, the it compares the CurrentUtilization 
against the upper and lower bounds defined by the step adjustments to determine which step adjustment 
to perform.

A concrete example for step scaling policy is:

* Lower bound: NULL, Upper bound: 20, Adjustment: -2; 
* Lower bound: 20, Upper bound: 40, Adjustment: -1;
* Lower bound: 60, Upper bound: 80, Adjustment: +1;
* Lower bound: 80, Upper bound: NULL, Adjustment: +2;

With this policy, you are maintaining the Utilization in [40, 60]:

* if utilization is lower than 40% but higher than 20%, decrease the replica by 1
* If utilization is lower than 20%, decrease the replica by 2
* if utilization is higher than 60% but lower than 80%, increase the replica by 1
* if utilization is higher than 80%, increase the replica by 2

There are a few rules for the step adjustments for your policy:

* The ranges of your step adjustments can't overlap or have a gap.
* Only one step adjustment can have a null lower bound (negative infinity). If one step adjustment has 
    a negative lower bound, then there must be a step adjustment with a null lower bound.
* Only one step adjustment can have a null upper bound (positive infinity). If one step adjustment has 
    a positive upper bound, then there must be a step adjustment with a null upper bound.
* The upper and lower bound can't be null in the same step adjustment.
* If the metric value is above the breach threshold, the lower bound is inclusive and the upper bound is exclusive. If the metric value is below the breach threshold, the lower bound is exclusive and the upper bound is inclusive.

## Scenarios where step scaling policy performs better

### Scenrios:

##### Pod capacity

Assume we have an application, one instance of the applicaiton  can serve 10 request per second with 100% CPU
utilization. 

##### Request pattern
At the very beginning, clients send 5 requests per second, then suddently begin to send 80 requests per second

##### HPA parameters
*  min pod num: 1
*  max pod num: 20

##### Target
We want to achieve high resource utilization as well as scale fast.

### Target Tracking Scaling with TargetCPUUtilization = 80%

scaling policy: because we want to have high resources utilization, so we set target cpu utilization as 80%

scaling steps:

* round 1: 5 requests per seconds, CurrentReplicas = 1
* roubd 2: 80 requests per seconds, CurrentReplicas = 1, DesiredReplicas = ceil(100/80) = 2
* round 3: 80 requests per seconds, CurrentReplicas = 2, DesiredReplicas = ceil(200/80) = 3
* round 4: 80 requests per seconds, CurrentReplicas = 3, DesiredReplicas = ceil(300/80) = 4
* round 5: 80 requests per seconds, CurrentReplicas = 4, DesiredReplicas = ceil(400/80) = 5
* round 6: 80 requests per seconds, CurrentReplicas = 5, DesiredReplicas = ceil(500/80) = 7
* round 7: 80 requests per seconds, CurrentReplicas = 7, DesiredReplicas = ceil(700/80) = 9
* round 8: 80 requests per seconds, CurrentReplicas = 9, DesiredReplicas = ceil(800/80) = 10

In other words, "Target Tracking Scaling Policy" need 8 rounds to scale to desire number of replicas.

###  Step Scaling Policy

scaling policy:

* average cpu utilization in range[40, 88]: do nothing
* 88% < average cpu utilization < 95% : add one more pod
* 95% < average cpu utilization : add two more pod
* 20% < average cpu utilization < 44% : remove one pod
* average cpu utilization < 20%	: remove two pod

scaling steps:

* round 1: 5 requests per seconds, CurrentReplicas = 1
* roubd 2: 80 requests per seconds, CurrentReplicas = 1, DesiredReplicas = 3
* roubd 3: 80 requests per seconds, CurrentReplicas = 3, DesiredReplicas = 5
* roubd 4: 80 requests per seconds, CurrentReplicas = 5, DesiredReplicas = 7
* roubd 5: 80 requests per seconds, CurrentReplicas = 7, DesiredReplicas = 9
* roubd 6: 80 requests per seconds, CurrentReplicas = 9, DesiredRe plicas = 10

In other words, "Step Scaling Policy" only need 6 rounds to scale to desire number of replicas.


### Target Tracking Scaling with TargetCPUUtilization = 50%
If we set the TargetCPUutilization as 50%, "Target Tracking Scaling Policy" can scale faster:

* round 1: 5 requests per seconds, CurrentReplicas = 1
* roubd 2: 80 requests per seconds, CurrentReplicas = 1, DesiredReplicas = ceil(100/50) = 2
* round 3: 80 requests per seconds, CurrentReplicas = 2, DesiredReplicas = ceil(200/50) = 4
* round 4: 80 requests per seconds, CurrentReplicas = 4, DesiredReplicas = ceil(400/50) = 8
* round 5: 80 requests per seconds, CurrentReplicas = 8, DesiredReplicas = ceil(800/50) = 16

With TargetCPUUtilization = 50%, we can scale to desire number of replicas in only 5 rounds, however it
desires 16 repliicas and the average cpu utilization is only 50%, on the other hand, "step 
scaling policy" has desire number of repliicas as 10 and average cpu utilization as 80%.

Someone may argue that we can set the max number of pods as 10, so even with TargetCPUUtilization = 50% we 
will have 10 pods at most. However, how about clients send 60 request per seconds? When clients send 60 
request per seconds, target tracking scaling policy (with TargetCPUUtilization = 50%) will desire 2 replicas,
while our step scaling policy only desire 1 replica.

In conclusion, if we want achieve high resource utilization and at the same time want to scale fast, 
the proposed "threshold based step scaling policy" performs better than the current "target tracking 
scaling policy"

