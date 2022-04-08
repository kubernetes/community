# 2021 Annual Report: WG Multitenancy

## Current initiatives

1. What work did the WG do this year that should be highlighted?
   For example, artifacts, reports, white papers produced this year.

   - Hierarchical Namespace Controller: work continued (reached 1.0 in April 2022), moved into its own k8s repo out of WG incubation in a major project milestone
   -- https://github.com/kubernetes-sigs/hierarchical-namespaces 
   - Vertical Cluster Project: moved out of WG incubation and into its own k8s repo in a major milestone 
   -- Now called "Cluster API Provider Nested" 
   -- https://github.com/kubernetes-sigs/cluster-api-provider-nested/tree/main/virtualcluster 
   - Multi-tenancy Benchmark Project: continued 
   -- https://github.com/kubernetes-sigs/multi-tenancy/tree/master/benchmarks
   - Kicked off multi-tenancy documentation project (in March 2022)
   -- https://docs.google.com/document/d/192aPEDsoJ-DWsy1GYvmQt_7tKP5wXh9MN9totE81Dx4/edit# 

2. What initiatives are you working on that aren't being tracked in KEPs?

   - Hierarchical Namespace Controller
   - Virtual Cluster Project
   - Multi-tenancy Benchmark Project
   - Multi-tenancy Documentation Project

## Project health

1. What's the current roadmap until completion of the working group?

   - We had a healthy discussion in December 2021 about whether it was time to spin the working group down or move it to another SIG, or petition for it to become its own SIG, given the long lived nature of its work. From that discussion the leaders of the WG agreed that our next obvious milestone is to add robust documentation to the Kubernetes' website on how to achieve multi-tenancy. We consistently get asked about how to do this, and the upstream documentation has no docs. So we agreed to pick this up as our next major milestone, and are executing on it. We also had some new projects start to emerge, so we are staying open minded about whether we should continue to provide a space for new projects, or if we should wind down once our documentation project is complete.

2. Does the group have contributors from multiple companies/affiliations?

   - Yes, we have leaders from VMware, Google, Alibaba, Nirmata, and Twilio. We also have a revolving series of contributors from many different companies in addition to these.

3. Are there ways end users/companies can contribute that they currently are not?
   If one of those ways is more full time support, what would they work on and why?

   - No, I think we're well supported with attention and concrete work at this time. 
   

## Membership

- Primary slack channel member count: 1168
- Primary mailing list member count: 738
- Primary meeting attendee count (estimated, if needed): 10
- Primary meeting participant count (estimated, if needed): 10

Include any other ways you measure group membership

## Operational

Operational tasks in [wg-governance.md]:

- [ x] [README.md] reviewed for accuracy and updated if needed
- [ x] WG leaders in [sigs.yaml] are accurate and active, and updated if needed
- [ x] Meeting notes and recordings for 2021 are linked from [README.md] and updated/uploaded if needed
- [ x] Updates provided to sponsoring SIGs in 2021
    

[wg-governance.md]: https://git.k8s.io/community/committee-steering/governance/wg-governance.md
[README.md]: https://git.k8s.io/community/wg-multitenancy/README.md
[sigs.yaml]: https://git.k8s.io/community/sigs.yaml

