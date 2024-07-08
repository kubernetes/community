# Writing Controllers

A Kubernetes controller is an active reconciliation process. That is, it watches some object for the world's desired state, and it watches the world's actual state, too. Then, it sends instructions to try and make the world's current state be more like the desired state.

The simplest implementation of this is a loop:

```go
for {
  desired := getDesiredState()
  current := getCurrentState()
  makeChanges(desired, current)
}
```

Watches, etc, are all merely optimizations of this logic.

## Guidelines

When you're writing controllers, there are few guidelines that will help make sure you get the results and performance you're looking for.

1. Operate on one item at a time.  If you use a `workqueue.Interface`, you'll be able to queue references to particular objects and later pop them in multiple “worker” goroutines with a guarantee that no two goroutines will work on the same item at the same time.

   Many controllers must trigger off multiple resources (I need to "check X if Y changes"), but nearly all controllers can collapse those into a queue of “check this X” based on relationships.  For instance, a ReplicaSet controller needs to react to a pod being deleted, but it does that by finding the related ReplicaSets and queuing references to those.

1. Random ordering between resources. When controllers queue off multiple types of resources, there is no guarantee of ordering amongst those resources.

   Distinct watches are updated independently.  Even with an objective ordering of “created resourceA/X” and “created resourceB/Y”, your controller could observe “created resourceB/Y” and “created resourceA/X”.

1. Level driven, not edge driven.  Just like having a shell script that isn't running all the time, your controller may be off for an indeterminate amount of time before running again.

   If an API object appears with a marker value of `true`, you can't count on having seen it turn from `false` to `true`, only that you now observe it being `true`.  Even an API watch suffers from this problem, so be sure that you're not counting on seeing a change unless your controller is also marking the information it last made the decision on in the object's status.

1. Use `SharedInformers`.  `SharedInformers` provide hooks to receive notifications of adds, updates, and deletes for a particular resource.  They also provide convenience functions for accessing shared caches and determining when a cache is primed.

   Use the factory methods down in https://git.k8s.io/kubernetes/staging/src/k8s.io/client-go/informers/factory.go to ensure that you are sharing the same instance of the cache as everyone else.

   This saves us connections against the API server, duplicate serialization costs server-side, duplicate deserialization costs controller-side, and duplicate caching costs controller-side.

   You may see other mechanisms like reflectors and deltafifos driving controllers.  Those were older mechanisms that we later used to build the `SharedInformers`.  You should avoid using them in new controllers.

1. Never mutate original objects!  Caches are shared across controllers, this means that if you mutate your "copy" (actually a reference or shallow copy) of an object, you'll mess up other controllers (not just your own).

   The most common point of failure is making a shallow copy, then mutating a map, like `Annotations`.  Use `api.Scheme.Copy` to make a deep copy.

1. Wait for your secondary caches.  Many controllers have primary and secondary resources.  Primary resources are the resources that you'll be updating `Status` for.  Secondary resources are resources that you'll be managing (creating/deleting) or using for lookups.

   Use the `cache.WaitForCacheSync` function to wait for your secondary caches before starting your primary sync functions.  This will make sure that things like a Pod count for a ReplicaSet isn't working off of known out of date information that results in thrashing.

1. There are other actors in the system.  Just because you haven't changed an object doesn't mean that somebody else hasn't.

   Don't forget that the current state may change at any moment--it's not sufficient to just watch the desired state. If you use the absence of objects in the desired state to indicate that things in the current state should be deleted, make sure you don't have a bug in your observation code (e.g., act before your cache has filled).

1. Failures happen, and their detection in Kubernetes is imperfect. Run multiple copies of your controller pod (e.g., via a `Deployment`), with one copy active and the other(s) ready to take over at a moment's notice. Use `k8s.io/client-go/tools/leader-election` for that. Even leader election is imperfect; even when using leader election it is still possible --- although very unlikely --- that multiple copies of your controller may be active.

1. Percolate errors to the top level for consistent re-queuing.  We have a  `workqueue.RateLimitingInterface` to allow simple requeuing with reasonable backoffs.

   Your main controller func should return an error when requeuing is necessary.  When it isn't, it should use `utilruntime.HandleErrorWithContext` and return nil instead.  This makes it very easy for reviewers to inspect error handling cases and to be confident that your controller doesn't accidentally lose things it should retry for.

1. Informers can periodically “resync”.  This means to call the Update event handler on every object in the informer's local cache.  This is good for cases where you may need to take additional action on the object, but sometimes you know there won't be more work to do.

   In cases where you are *certain* that you don't need to requeue items when there are no new changes, you can compare the resource version of the old and new objects.  If they are the same, you skip requeuing the work.  Be careful when you do this.  If you ever skip requeuing your item on failures, you could fail, not requeue, and then never retry that item again.

1. If the primary resource your controller is reconciling supports ObservedGeneration in its status, make sure you correctly set it to metadata.Generation whenever the values between the two fields mismatches.

    This lets clients know that the controller has processed a resource. Make sure that your controller is the main controller that is responsible for that resource, otherwise if you need to communicate observation via your own controller, you will need to create a different kind of ObservedGeneration in the Status of the resource.

1. Consider using owner references for resources that result in the creation of other resources (eg. a ReplicaSet results in creating Pods). Thus you ensure that children resources are going to be garbage-collected once a resource managed by your controller is deleted. For more information on owner references, read more [here](https://git.k8s.io/design-proposals-archive/api-machinery/controller-ref.md).

    Pay special attention in the way you are doing adoption. You shouldn't adopt children for a resource when either the parent or the children are marked for deletion. If you are using a cache for your resources, you will likely need to bypass it with a direct API read in case you observe that an owner reference has been updated for one of the children. Thus, you ensure your controller is not racing with the garbage collector.

 See [k8s.io/kubernetes/pull/42938](https://github.com/kubernetes/kubernetes/pull/42938) for more information.

## Rough Structure

Overall, your controller should look something like this:

```go
type Controller struct {
	// pods gives cached access to pods.
	pods informers.PodLister
	podsSynced cache.InformerSynced

	// queue is where incoming work is placed to de-dup and to allow "easy"
	// rate limited requeues on errors
	queue workqueue.TypedRateLimitingInterface[cache.ObjectName]
}

func NewController(pods informers.PodInformer) *Controller {
	c := &Controller{
		pods: pods.Lister(),
		podsSynced: pods.Informer().HasSynced,
		queue: workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "controller-name"),
	}
	
	// register event handlers to fill the queue with pod creations, updates and deletions
	pods.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			ref, err := cache.ObjectToName(obj)
			if err == nil {
				c.queue.Add(ref)
			}
		},
		UpdateFunc: func(old interface{}, new interface{}) {
			ref, err := cache.ObjectToName(new)
			if err == nil {
				c.queue.Add(ref)
			}
		},
		DeleteFunc: func(obj interface{}) {
			// IndexerInformer uses a delta nodeQueue, therefore for deletes we have to use this
			// key function.
			ref, err := cache.DeletionHandlingObjectToName(obj)
			if err == nil {
				c.queue.Add(ref)
			}
		},
	},)
	
	return c
}

func (c *Controller) Run(ctx context.Context, threadiness int) error {
	// don't let panics crash the process
	defer utilruntime.HandleCrash()
	// make sure the work queue is shutdown which will trigger workers to end
	defer c.queue.ShutDown()
	logger := klog.FromContext(ctx)

	logger.Info("Starting <NAME> controller")

	// wait for your secondary caches to fill before starting your work
	if !cache.WaitForCacheSync(ctx.Done(), c.podsSynced) {
		return fmt.Errorf("failed to wait for caches to sync")
	}

	// start up your worker threads based on threadiness.  Some controllers
	// have multiple kinds of workers
	for i := 0; i < threadiness; i++ {
		// runWorker will loop until "something bad" happens.  The .Until will
		// then rekick the worker after one second
		go wait.UntilWithContext(ctx, c.runWorker, time.Second)
	}
	logger.Info("Started workers")

	// wait until we're told to stop
	<-ctx.Done()
	logger.Info("Shutting down <NAME> controller")

	return nil
}

func (c *Controller) runWorker(ctx context.Context) {
	// hot loop until we're told to stop.  processNextWorkItem will
	// automatically wait until there's work available, so we don't worry
	// about secondary waits
	for c.processNextWorkItem(ctx) {
	}
}

// processNextWorkItem deals with one item off the queue.  It returns false
// when it's time to quit.
func (c *Controller) processNextWorkItem(ctx context.Context) bool {
	// Pull the next work item from queue.  It will be an object reference that we use to lookup
	// something in a cache
	ref, shutdown := c.queue.Get()
	if shutdown {
		return false
	}
	// you always have to indicate to the queue that you've completed a piece of
	// work
	defer c.queue.Done(ref)

	// Process the object reference.  This method will contains your "do stuff" logic
	err := c.syncHandler(ref)
	if err == nil {
		// if you had no error, tell the queue to stop tracking history for your
		// item. This will reset things like failure counts for per-item rate
		// limiting
		c.queue.Forget(ref)
		return true
	}

	// there was a failure so be sure to report it.  This method allows for
	// pluggable error handling which can be used for things like
	// cluster-monitoring
	utilruntime.HandleErrorWithContext(ctx, err, "Error syncing; requeuing for later retry", "objectReference", ref))

	// since we failed, we should requeue the item to work on later.  This
	// method will add a backoff to avoid hotlooping on particular items
	// (they're probably still not going to work right away) and overall
	// controller protection (everything I've done is broken, this controller
	// needs to calm down or it can starve other useful work) cases.
	c.queue.AddRateLimited(ref)

	return true
}
```
