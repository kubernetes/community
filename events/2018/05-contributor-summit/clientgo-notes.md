# Client-go
**Lead:** munnerz with assist from lavalamp  
**Slides:** combined with the CRD session [here](https://www.dropbox.com/s/n2fczhlbnoabug0/API%20extensions%20contributor%20summit.pdf?dl=0) (CRD is first; client-go is after)  
**Thanks to our notetakers:** kragniz, mrbobbytales, directxman12, onyiny-ang

## Goals for the Session

* What is currently painful when building a controller
* Questions around best practices
* As someone new:
  * What is hard to grasp?
* As someone experienced:
  * What important bits of info do you think are critical


## Pain points when building controller
* A lot of boilerplate
  * Work queues
  * HasSynced functions
  * Re-queuing
* Lack of deep documentation in these areas
  * Some documentation exists, bot focused on k/k core
* Securing webhooks & APIServers
* Validation schemas
* TLS, the number of certs is a pain point
  * It is hard right now, the internal k8s CA has been used a bit.
  * OpenShift has a 'serving cert controller' that will generate a cert based on an annotation that might be able to possibly integrate upstream.
  * Election has been problematic and the Scaling API is low-level and hard to use. doesn't work well if resource has multiple meanings of scale (eg multiple pools of nodes)
* Registering CRDs, what's the best way to go about it?
  * No best way to do it, but has been deployed with application
  * Personally, deploy the CRDs first for RBAC reasons
* Declarative API on one end that has to be translated to translated to a transactional API on the other end (e.g. ingress). Controller trying to change quite a few things.
  * You can do locking, but it has to be built.
* Q: how do you deal with "rolling back" if the underlying infrastructure
  that you're describing says no on an operation?
  * A: use validating webhook?
  * A: use status to keep track of things?
  * A: two types of controllers: `kube --> kube` and `kube --> external`,
    they work differently
  * A: Need a record that keeps track of things in progress. e.g. status. Need more info on how to properly tackle this problem.


## Best practices
(discussion may be shown by Q: for question or A: for audience or answer)
* How do you keep external resources up to date with Kubernetes resources?
  * A: the original intention was to use the sync period on the controller if
    you watch external resources, use that
  * Should you set resync period to never if you're not dealing with
    external resources?
    * A: Yes, it's not a bug if watch fails to deliver things right
    * A: controller automatically relists on connection issues, resync
      interval is *only* for external resources
  * maybe should be renamed to make it clear it's for external resources
* how many times to update status per sync?
  * A: use status conditions to communicate "fluffy" status to user
    (messages, what might be blocked, etc, in HPA), use fields to
    communicate "crunchy" status (last numbers we saw, last metrics, state
    I need later).
* How do I generate nice docs (markdown instead of swagger)
  * A: kubebuilder (kubernetes-sigs/kubebuilder) generates docs out of the
    box
  * A: Want to have IDL pipeline that runs on native types to run on CRDs,
    run on docs generator
* Conditions vs fields
  * used to check a pods’ state
  * "don't use conditions too much"; other features require the use of conditions, status is unsure
  * What does condition mean in this context
    * Additional fields that can have `ready` with a msg, represents `state`.
    * Limit on states that the object can be in.
    * Use conditions to reflect the state of the world, is something blocked etc.
    * Conditions were created to allow for mixed mode of clients, old clients can ignore some conditions while new clients can follow them. Designed to make it easier to extend status without breaking clients.
* Validating webhooks vs OpenAPI schema
* Can we write a test that spins up main API server in process?
  * Can do that current in some k/k tests, but not easy to consume
  * vendoring is hard
  * Currently have a bug where you have to serve aggregated APIs on 443,
    so that might complicate things
* How are people testing extensions?
  * Anyone reusing upstream dind cluster?
  * People looking for a good way to test them.
  * kube-builder uses the sig-testing framework to bring up a local control plane and use that to test against. (@pwittrock)
* How do you start cluster for e2es?
  * Spin up a full cluster with kubeadm and run tests against that
  *  integration tests -- pull in packages that will build the clusters
* Q: what CIs are you using?
   * A: Circle CI and then spin up new VMs to host cluster
   * Mirtantis has a tool for a multi-node dind cluster for testing
*  #testing-commons channel on stack. 27 page document on this--link will be put in slides
* Deploying and managing Validating/Mutating webhooks?
    * how complex should they be?
* When to use subresources?
  * Are people switching to api agg to use this today?
  * Really just for status and scale
  * Why not use subresources today with scale?
     * multiple replicas fields
     * doesn't fit polymorphic structure that exists
     * pwittrock@: kubectl side, scale
     * want to push special kubectl verbs into subresources to make kubectl
    more tolerant to version skew

## Other Questions

* Q: Client-go generated listers, what is the reason for two separate interfaces to retrieve from client and cache?
  * A: historical, but some things are better done local vs on the server.
* issues: client-set interface allows you to pass special options that allow you to do interesting stuff on the API server which isn't necessarily possible in the lister.
     * started as same function call and then diverged
     * lister gives you slice of pointers
     * clientset gives you a slice of not pointers
     * a lot of people would take return from clientset and then convert it to a slice of pointers so the listers helped avoid having to do deep copies every time. TLDR: interfaces are not identical
* Where should questions go on this topic for now?
  * A: most goes to sig-api-machinery right now
  * A : Controller related stuff would probably be best for sig-apps
* Q: Staleness of data, how are people dealing with keeping data up to date with external data?
  * A: Specify sync period on your informer, will put everything through the loop and hit external resources.
* Q: With strictly kubernetes resources, should your sync period be never? aka does the watch return everything.
  * A: The watch should return everything and should be used if its strictly k8s in and k8s out, no need to set the sync period.
* Q: What about controllers in other languages than go?
  * A: [metacontroller](https://github.com/GoogleCloudPlatform/metacontroller) There are client libs in other languages, missing piece is work queue,
    informer, etc
* Cluster API controllers cluster, machineset, deployment, have a copy of
  deployment code for machines. Can we move this code into a library?
  * A: it's a lot of work, someone needs to do it
  * A: Janet Kuo is a good person to talk to (worked on getting core workloads
  API to GA) about opinions on all of this
* Node name duplication caused issues with AWS and long-term caches
  * make sure to store UIDs if you cache across reboot

## Moving Forwards
* How do share/disseminate knowledge (SIG PlatformDev?)
  * Most SIGs maintain their own controllers
  * Wiki? Developer Docs working group?
  * Existing docs focus on in-tree development. Dedicated 'extending kubernetes' section?
* Git-book being developed for kubebuilder (book.kubebuilder.io); would appreciate feedback @pwittrock
* API extensions authors meetups?
* How do we communicate this knowledge for core kubernetes controllers
  * Current-day: code review, hallway conversations
* Working group for platform development kit?
* Q: where should we discuss/have real time conversations?
  * A: #sig-apimachinery, or maybe #sig-apps in slack (or mailing lists) for the workloads controllers
