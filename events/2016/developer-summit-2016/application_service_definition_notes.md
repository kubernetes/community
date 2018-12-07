# Service/Application Definition

We think we need to help out the developers in how do we organize our services and how do I define them nicely and deploy on our orchestrator of choice.  Writing the Kube files is a steep learning curve.  So can we have something which is a little bit easier?

Helm solves one purpose for this.  

Helm contrib: one of the things folks as us is they start from a dockerfile, and they want to have a workflow where they go from dockerfile-->imagebuild-->registry-->resource def.

There are different ways to package applications.  There's the potential for a lot of fragmentation in multi-pod application definitions.  Can we create standards here?

We want to build and generate manifests with one tool.  We want "fun in five" that is have it up and running in five minutes or less.

Another issue is testing mode; currently production-quality Helm charts don't really work on minikube,.  There's some issues around this which we know about.  We need dummy PVCs, LoadBalancer, etc.   Also DNS and Ingress.

We need the 80% case, Fabric8 is a good example of this.  We need a good set of boundary conditions so that the new definition doesn't get bigger than the Kube implementation. Affinity/placement is a good example of "other 20%".  

We also need to look at how to get developer feedback on this so that we're building what they need.  Pradeepto did a comparison of Kompose vs. Docker Compose for simplicity/usability.

One of the things we're discussing the Kompose API.  We want to get rid of this and supply something which people can use directly with kubernetes.  A bunch of shops only have developers.  Someone asked though what's so complicated with Kube definitions.  Have we identified what gives people trouble with this?  We push too many concepts on developers too quickly.  We want some high-level abstract types which represent the 95% use case.  Then we could decompose these to the real types.

What's the gap between compose files and the goal?  As an example, say you want to run a webserver pod.  You have to deal with ingress, and service, and replication controller, and a bunch of other things.  What's the equivalent of "docker run" which is easy to get. The critical thing is how fast you can learn it.

We also need to have reversability so that if you use compose you don't have to edit the kube config after deployment, you can still use the simple concepts.  The context of the chart needs to not be lost.

There was discussion of templating applications.  Person argued that it's really a type system.  Erin suggested that it's more like a personal template, like the car seat configuration.  

There's a need to let developers work on "their machine" using the same spec.  Looking through docker-compose, it's about what developers want, not what kubernetes wants.  This needs to focus on what developers know, not the kube objects.

Someone argued that if we use deployments it's really not that complex.  We probably use too much complexity in our examples.  But if we want to do better than docker-compose, what does it look like?  Having difficulty imagining what that is.  

Maybe the best approach is to create a list of what we need for "what is my app" and compare it with current deployment files.  

There was a lot of discussion of what this looks like.

Is this different from what the PAASes already do?  It's not that different, we want something to work with core kubernetes, and also PAASes are opinionated in different ways.

Being able to view an application as a single unifying concept is a major desire.   Want to click "my app" and see all of the objects associated with it.  It would be an overlay on top of Kubernetes, not something in core.

One pending feature is that you can't look up different types of controllers in the API, that's going to be fixed.  Another one is that we can't trace the dependencies; helm doesn't label all of the components deployed with the app.

Need to identify things which are missing in core kubernetes, if there are any.

## Action Items:

* Reduce the verbosity of injecting configmaps.  We want to simplify the main kubernetes API.  For example, there should be a way to map all variables to ENV as one statement.
* Document where things are hard to understand with deployments.
* Document where things don't work with minikube and deployments.
* Document what's the path from minecraft.jar to running it on a kubernetes cluster?
