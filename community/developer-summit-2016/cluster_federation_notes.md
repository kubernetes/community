# Cluster Federation

There's a whole bunch of reasons why federation is interesting.  There's HA, there's geographic locality, there's just managing very large clusters.  Use cases:

* HA
* Hybrid Cloud
* Geo/latency
* Scalability (many large clusters instead of one gigantic one)
* visibility of multiple clusters

You don't actually need federation for geo-location now, but it helps.  The mental model for this is kind of like Amazon AZ or Google zones.  Sometimes we don't care where a resource is but sometimes we do. Sometimes you want specific policy control, like regulatory constraints about what can run where.

From the enterprise point of view, central IT is in control and knowledge of where stuff gets deployed.  Bob thinks it would be a very bad idea for us to try to solve complex policy ideas and enable them, it's a tar pit.  We should just have the primitives of having different regions and be able to say what goes where.

Currently, you either do node labelling which ends up being complex and dependent on discipline.  Or you have different clusters and you don't have common namespaces.  Some discussion of Intel proposal for cluster metadata.  

Bob's mental model is AWS regions and AZs.  For example, if we're building a big cassandra cluster, and you want to make sure that nodes aren't all in the same zone.

Quinton went over a WIP implementation for applying policies, with a tool which applies policy before resource requests go to the scheduler.  It uses an open-source policy language, and labels on the request.

Notes interrupted here, hopefully other members will fill in.
