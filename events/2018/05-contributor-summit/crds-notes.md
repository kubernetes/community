# CRDs - future and painpoints
**Lead:** sttts  
**Slides:** combined with the client-go session [here](https://www.dropbox.com/s/n2fczhlbnoabug0/API%20extensions%20contributor%20summit.pdf?dl=0)  
**Thanks to our notetakers:** mrbobbytales, kragniz, tpepper, and onyiny-ang

## outlook - aggregation
* API stable since 1.10. There is a lack of tools and library support.
* GSoC project with @xmudrii: share etcd storage
  * `kubectl create etcdstorage your api-server`
* Store custom data in etcd

## outlook custom resources

1.11:
* alpha: multiplier versions with/without conversion
* alpha: pruning - blocker for GA - unspecified fields are removed
  * deep change of semantics of custom resources
  * from JSON blob store to schema based storage
* alpha: defaulting - defaults from openapi validation schema are applied
* alpha: graceful deletion - (maybe? PR exists)
* alpha: server side printing columns for `kubectl get` customization
* beta: subresources - alpha in 1.10
* will have additionalProperties with extensible string map
  * mutually exclusive with properties

1.12
* multiple versions with declarative field renames
* strict create mode (issue #5889)

Missing from Roadmap:
   - Additional Properties: Forbid additional fields
        - Unknown fields are silently dropped instead of erroring
    - Istio used CRD extensively: proto requires some kind of verification and CRDs are JSON
        - currently planning to go to GA without proto support
        - possibly in the longer term to plan
    - Resource Quotas for Custom Resources
        - doable, we know how but not currently implemented
    - Defaulting: mutating webhook will default things when they are written
    - Is Validation going to be required in the future
        - poll the audience!
        - gauging general sense of validation requirements (who wants them, what's missing?)
        - missing: references to core types aren't allowed/can't be defined -- this can lead to versioning complications
    - limit CRDs clusterwide such that the don't affect all namespaces
        - no good discussion about how to improve this yet
        - feel free to start one!
    - Server side printing columns, per resource type needs to come from server -- client could be in different version than server and highlight wrong columns

Autoscaling is alpha today hopefully beta in 1.11

## The Future: Versioning
* Most asked feature, coming..but slowly
* two types, "noConversion" and "Declarative Conversion"
* "NoConversion" versioning
  * maybe in 1.11
  * ONLY change is apiGroup
  * Run multiple versions at the same time, they are not converted

* "Declarative Conversion" 1.12
* declarative rename e.g
```
spec:
  group: kubecon.io
  version: v1
  conversions:
    declarative:
      renames:
        from: v1pha1
        to: v1
        old: spec.foo
        new: bar
```
* Support for webhook?
  *  not currently, very hard to implement
  * complex problem for end user
  * current need is really only changing for single fields
  * Trying to avoid complexity by adding a lot of conversions

## Questions:
* When should someone move to their own API Server
  * At the moment, telling people to start with CRDs. If you need an aggregated API server for custom versioning or other specific use-cases.
* How do I update everything to a new object version?
  * Have to touch every object.
* are protobuf support in the future?
  * possibly, likely yes
* update on resource quotas for CRDs
  * PoC PR current out, it's doable just not quite done
* Is validation field going to be required?
  *  Eventually, yes? Some work being done to make CRDs work well with `kubectl apply`
* Can CRDs be cluster wide but viewable to only some users.
  * It's been discussed, but hasn't been tackled.
* Is there support for CRDs in kubectl output?
  * server side printing columns will make things easier for client tooling output. Versioning is important for client vs server versioning.
