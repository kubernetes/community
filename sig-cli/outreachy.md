# Outreachy

Kubernetes, specifically the SIG-CLI, is happy to announce our participation in the Outreachy program, running from December 2017 to March 2018. 
Please see the [main program page](https://www.outreachy.org/) for the general information about the program, such as its purpose, timeline, eligibility requirements, and how to apply.

**Schedule**

* October 23: application deadline
* November 9: selection decisions are made
* December 5 - March 5: internship

## What is Kubernetes?
Kubernetes is an open-source system for automating deployment, scaling, and management of containerized applications.
It groups containers that make up an application into logical units for easy management and discovery. Kubernetes builds upon 15 years of experience of running production workloads at Google, combined with best-of-breed ideas and practices from the community.

[kubernetes.io](https://kubernetes.io/)

## What are SIGs / What is SIG-CLI?
Kubernetes is a set of projects, each shepherded by a special interest group (SIG). To get a grasp of the projects that we work on, check out the complete [list of SIGs](https://github.com/kubernetes/community/blob/master/sig-list.md).

SIG-CLI Covers kubectl and related tools. We focus on the development and standardization of the CLI framework and its dependencies, the establishment of conventions for writing CLI commands, POSIX compliance, and improving the command line tools from a developer and devops user experience and usability perspective.  

**Commuication:**

SIG-CLI Mailing List - kubernetes-sig-sli@googlegroups.com  
Slack - k8s.slack.io for invite; #sig-cli channel

**Mentors**

Philip Wittrock - pwittrock@google.com  
Similar to contacting the coordinators, the mentors can be contacted at any time either by sending messages to the mailing lists or slack channels.

## Contribute

As part of the application process, the Outreachy program recommends that candidates make small contributions to the project they intend to apply for.
To start working on the project, make sure to fill out the CLA and check if you have the right environment with this guide. The README in the [community repo](https://github.com/kubernetes/community) details these things and more.

Check out these specific resources for how to contribute to CLI: 
* SIG-CLI - [How to Contribute](https://github.com/kubernetes/community/blob/master/sig-cli/CONTRIBUTING.md)
* Filter issue search for: `is:open is:issue label:sig/cli label:help-wanted` (or `label:for-new-contributors`)  

## Available tasks

Develop `kubectl create` commands to make it easy to create Kubernetes resources  
Develop `kubectl set` commands to modify Kubernetes resources  
Required Skills: Go  
Mentor: Phillip Wittrock

**Coordination:**

* Paris Pittman - parispittman@google.com
* Josh Berkus
* Elsie Phillips

The coordinators can be contacted at any time. The easiest way is to send a slack message.

Do you have an idea for a task that is suitable for this program? Contact the mentors or coordinators! Or even better, volunteer for mentoring an intern during the work on your idea!

## Code of Conduct
Kubernetes abides by the CNCF [Code of Conduct](https://github.com/cncf/foundation/blob/master/code-of-conduct.md)
