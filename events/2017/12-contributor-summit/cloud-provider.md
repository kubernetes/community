# CloudProvider Update 
*Lead: Walter Fender / cheftako*  
*Notetaker: Jago Macleod / jagosan*
4pm 

Goals: 

Make it easy for any new cloud provider to be able to add support for new cloud provider without having to submit code to k/k repo. [good shape!]
For currently in-tree cloud providers to be able to build and release without having to ‘wait’ for upstream kubernetes. (e.g. security patches and features)
Also suboptimal that for technical reasons, other cloudprovider’s ‘init’s are being run in the hosting cloud provider’s environments. 

[TODO: name?]: From a user’s perspective, you might like a more modular cloudprovider. You might want ELB but not the storage from that cloudprovider. 

cheftako: where are we: 

There is a PR out that is a WIP that proves that in GCE it is possible to bring up a cluster 

Kube-controller-manager is disabled. cloud-controller-manager is working there 

Kube-mark tests are a challenge. No support for conditional behavior in test-infra. Now we need an n1-standard2 master. 

Other outstanding issues: 

Who has seen the following in the tests: 

If !GCE -> skip. 

Many raise hands. This is recognized as a gap that needs to be addressed. 

Also need volunteers for each of the cloud providers. Lot of work to be done. 

Thockin would like to minimize the amount of code that lives outside of k/k repo. If we were to take all of lifecycle and move it out of k/k, there is a lot of functionality that could diverge, or could be missing for a user in certain cloud providers. 

Take a page out of API Server book and make a generic controller manager and use for kube-controller-manager and cloud-controller-manager. 

Comment: spent a couple of days looking through the GCE work and found ~4 bugs. So there are a bunch of other 

Broke node controller out into 4 segments. 

Overview on where we are in the overall project: external cloud provicer

We have a long way to go before all the in-tree cloud providers are moved out. 

One of the interesting questions is that eventually there are no in-tree cloud providers; there is now a distributed CI issue where all cloud providers have a mechanism through which to set up their own continuous testing and post results back to testgrid, to expose when changes to one area of the code inadvertently break one or many or all cloud providers. 

MS has a proof of concept basically done. Some questions: legally, what does this mean? E2E tests - how to make blocking. OpenStack -- in e2e tests, there is no tag to ensure nothing was broken in e2e. 

Jaice: to complicate matters, there may be primitives that make it necessary to make the testing much broader. There is a greater chance that the flakes occur in this world. 

thockin - on indemnity, we can make a dedicated repo that is still CNCF owned but controlled by cloud provider 

Exercise the top level interface 

Multiple levels of test here. In the final state, there is no ‘cloudprovider’ interface. You can test that binary / or binaries to test specific cloud implementations; then there is another level which is the e2e tests of basically everything that is the kubernetes functionality. 

When it comes to timing: the original plan was to depend on Flex Volumes. Subsequent discussions have tended toward depending instead on CSI as the volume solution. This means the timing is no earlier than Q3. 

Victory looks like this: 

Major cloud providers shipping and running cloud-controller-manager. Volumes can remain behind until CSI is production ready. 

Also, when we can eventually delete all the in-tree cloud provider code, thockin will buy the cake and the keg. 

Today, when you deal with CCM, and it says ‘what is my cloud provider’ it turns off cloud plugin, and volumes doesn’t work. There is, therefore, a new flag. Set cloudprovider=external and some other flag, can’t remember off top of head. Intention is that this will be a temporary second flag and will be removed as soon as possible. 

A lot of really good work has been done. Need to do a better job of organizing code in KCM & CCM. 

Open Question: will we ship cloud provider specific CCM’s as part of the kubernetes release? 

Or, what cloud providers are included or easily available with the binaries distributed in the Kubernetes release?

Can we define better criteria than ‘you showed up before date xxxxxx’. 

Perhaps there is something that will work for a ‘vanilla’ system or locally, and some way to discover and install CCMs for available cloud providers. Likely the vendor / distributor will be packaging up some other process by which kubernetes is installed and configured. 

Thockin - think about what you do when you compile the linux kernel. Choose options during initialization and customize to hearts content. At the end, you get an artifact. Kbuild. Did a hack 18 months ago and proved it’s possible. 

thockin - do we have concrete list? Not a great list. We are going through the KEP process. AI: call to action. 

We are committed; we are making progress; we meet wednesdays at 10am Pacific Time. 

Fejta: think about a local provider. Move as much as possible onto something that is like minikube, e.g.  Great enthusiasm for this concept. 

AI: fejta to share information on how to submit test results to testgrid and how to leverage the tooling for external cloud providers. 

Summary - yes!
