# 2021 WG Multitenancy Annual Report

### What was the initial mission of the group and if it's changed, how?

**Initial Mission:**

Define the models of multitenancy that Kubernetes will support. Discuss and execute upon any
remaining work that needs to be done to support these models. Create conformance tests that 
will prove that these models can be built and used in production environments.

**Current Mission:**

We are focusing more on projects using Kubernetes than trying to directly change the API machinery, 
which is a huge and potentially intractable problem. We are continuing to work on conformance 
testing as part of one of our projects.

### What’s the current roadmap until completion?

**We have three projects we are incubating:**

* Multi-Tenancy Benchmarks https://github.com/kubernetes-sigs/multi-tenancy/tree/master/benchmarks 
* Virtual Cluster Project  https://github.com/kubernetes-sigs/multi-tenancy/tree/master/incubator/virtualcluster 
* Hierarchical Namespace Controller https://github.com/kubernetes-sigs/multi-tenancy/tree/master/incubator/hnc 

These are all in active development, and we’re making good progress. Google Cloud has adopted HNC in beta. 
MTB is going to be a great platform for the conformance test suite expansion. Virtual Cluster Project is 
graduating out of our incubator and into its own repo! HNC will probably follow soon.

The roadmap is documented in the [working group project plan](https://docs.google.com/document/d/1U8RQQmTUjxgMZY05HG2f7b3KsB94BhK4Ko6aWbLNXcc/edit).

### Have you produced any artifacts, reports, white papers to date?

You can find a bunch of our docs here: https://github.com/kubernetes-sigs/multi-tenancy/tree/master/docs 

### Is the group active? healthy? contributors from multiple companies and/or end user companies?

Yes, it’s one of the most egalitarian working groups I’ve seen. We have active contributors and participants 
from all over the industry, and a lot of drive bys from consumers of Kubernetes who are just trying to 
configure clusters for multi-tenancy and have questions. Our incubating projects are led by people from 
Google, Nirmata, Alibaba, and Medtronic, have reviewers and participants from other companies, and the WG chairs work at VMware and Cisco.
We have a very diverse group of presenters from different companies who are all trying to solve the same problems, and we all have 
the philosophy of learning from each other and sharing. 

### Is everything in your readme accurate? posting meetings on youtube?

Yes and Yes.

### Do you have regular check-ins with your sponsoring SIGs?

We have a huge number of sponsoring SIGs, many of them send representatives to meetings on an adhoc basis.

### Links to the last two community meeting updates the group has given and notable highlights you’d like to share from those.

