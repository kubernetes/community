# API Extensions (SIG API Machinery position statement)

Authors: Daniel Smith, David Eads (SIG API Machinery co-leads)  
Last edit: Feb 23  
Status: RELEASED


## Background
We have observed a lot of confusion in the community around the general topic
of ThirdPartyResources (TPRs) and apiserver aggregation (AA). We want to
document the current position of the API Machinery SIG.

Extremely briefly, TPR is a mechanism for lightweight, easy extension of the
kubernetes API, which has collected a [significant userbase](https://gist.github.com/philips/a97a143546c87b86b870a82a753db14c).
AA is a heavier-weight mechanism for accomplishing a similar task; it is
targeted at allowing the Kubernetes project to move away from a monolithic
apiserver, and as a consequence, it will support PaaSes or other users that
need the complete set of server-side kubernetes API semantics.


## Positions
### Q: Do we need two extension mechanisms, or should we provide a single
extension mechanism with multiple opt-in features for users to grow into?
(Binary vs gradient)

We think there is both room in this space and a necessity for both approaches.
TPR is clearly useful to users. In its current state, TPR lacks some features
and has some bugs which limit it. We believe TPR bugs should be fixed and some
features should be added to it (as long as it maintains its ease-of-use, which
we think is its primary feature). We think TPR’s competitive advantage is its
low barrier-to-entry and ease of use.

However, even in the limit where we have added all the features to TPR that
make sense, there’s still a need for apiserver aggregation. Here are two use
cases that TPR cannot address while maintaining its ease of use.
* Heapster / metrics API. The metrics API is going to be data assembled at read
time, which is extremely high churn and should not be stored in an etcd
instance. Heapster needs to use custom storage.
* Full-featured extension APIs (pieces of Kubernetes itself; PaaSes).
  * OpenShift is an example of a full-featured API server that makes use of the
  apimachinery and apiserver features (API versioning, conversion, defaulting,
  serialization (including protocol buffer encoding), storage, security, custom
  subresource handlers, and admission).
  * Integrators who wish to provide this level of features and expect this
  level of API traffic volume are unlikely to be satisfied by webhooks, but
  should still be able to integrate. 
  * If Kubernetes developers could create new APIs in new apiservers instead
  of modifying the core apiserver, it would make life better for everyone:
    * Easier to shard reviews
    * Easier to experiment with APIs
    * No more accidentally enabling a half-baked API
    * Code freeze/release train less disruptive
  * It would be great if it were possible to run these extensions (including
  OpenShift, other PaaSes, and various optional extensions such as the service
  catalog) directly on an existing kubernetes cluster; in fact, we think that
  the alternative to this is a multiplication of forks, which will be really
  bad for the ecosystem as a whole. With ecosystem unification in mind, it
  would be infeasible to ask any consumer with both many users and an
  extensive codebase (such as OpenShift) to rewrite their stack in terms of
  TPRs and webhooks. We have to give such users a path to straight consumption
  as opposed to the current fork-and-modify approach, which has been the only
  feasible one for far too long.

This is not to say that TPR should stay in its current form. The API Machinery
SIG is committed to finishing TPR, making it usable, and maintaining it (but we
need volunteers to step up, or it’s going to take a long time).

The big table in [Eric’s comparison doc](https://docs.google.com/document/d/1y16jKL2hMjQO0trYBJJSczPAWj8vAgNFrdTZeCincmI/edit#heading=h.xugwibxye5f0)
is a good place to learn the current and possible future feature sets of TPRs
and AA. The fact that TPR has been languishing is due to lack of an owner and
lack of people willing to work on it, not lack of belief that it ought to be
fixed and perfected. Eric and Anirudh have agreed to take on this role.

### Q: Should there be a single API object that programs either TPR or AA as appropriate, or should each of these have their own registration object?
We think that the configuration of these two objects is distinct enough that
two API resources are appropriate.

We do need to take care to provide a good user experience, as the API groups
users enter in both AA and TPR come out of the same global namespace. E.g., a
user should not have to make both a TPR registration and an AA registration to
start up a TPR--this would break current users of TPRs.

### Q: Should TPRs be fixed up and extended in-place, or should a replacement be built in a separate TPR apiserver?
TPR is implemented currently with a variety of special cases sprinkled
throughout kube-apiserver code. It would greatly simplify kube-apiserver code
and the TPR implementation if this were separated, and TPR constructed as its
own HTTP server (but still run from kube-apiserver; see bottom Q). However, we
will not block safe, targeted TPR fixes on completion of this split.

### Q: Should TPR maintain compatibility, or should we break compatibility to fix and extend it?
There are two dozen open-source projects that use TPR, and we also know of
private users of TPR, and at least some people consider it to be beta. However,
we may have to implement fixes in a way that requires breaking backward
compatibility. If we do that, we will at a minimum provide migration
instructions and go through a one-release deprecation cycle to give users time
to switch over to the new version. We think this decision is probably best made
by the people actually working on this (currently: @deads2k, @erictune,
@foxish). [Some thoughts here](https://docs.google.com/document/d/1Gg158jO1cRBq-8RrWRAWA2IRF9avscuRaWFmY2Wb6qw/edit).

### Q: Should kube-aggregator be a separate binary/process than kube-apiserver?
For code health reasons, it is very convenient to totally separate the
aggregation layer from apiserver. However, operationally, it is extremely
inconvenient to set up and run an additional binary. Additionally, it is
crucial that all extensibility functionality be in every cluster, because users
need to be able to depend on it; this argues that kube-aggregator can’t be
optional.

Our current plan is to host several logical apiservers (the existing
kube-apiserver, kube-aggregator, and perhaps a hypothetical kube-tprserver,
see above) in a single binary, and launch them in a single process (a drop-in
replacement for the existing kube-apiserver). There are several candidate
mechanisms for accomplishing this and we won’t design this in this document. :)
