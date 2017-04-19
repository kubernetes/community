# SIG Scalability

Responsible for answering scalability related questions such as:

What size clusters do we think that we should support with Kubernetes in the short to 
medium term? How performant do we think that the control system should be at scale? 
What resource overhead should the Kubernetes control system reasonably consume?

For more details about our objectives please review [Scaling And Performance Goals](goals.md)

## Organizers
- Bob Wise (@countspongebob), Samsung-CNCT 
- Joe Beda (@jbeda), Heptio

## Meetings

- **Every Thursday at 9am pacific.**  
- Contact Joe or Bob for invite. 
- [Zoom link](https://zoom.us/j/989573207)
- [Agenda items](https://docs.google.com/a/bobsplanet.com/document/d/1hEpf25qifVWztaeZPFmjNiJvPo-5JX1z0LSvvVY5G2g/edit?usp=drive_web)

## Slack / Google Groups
- [Slack: #sig-scale](https://kubernetes.slack.com/messages/sig-scale/).  
- [Slack Archive](http://kubernetes.slackarchive.io/sig-scale/)

- [kubernetes-sig-scale](https://groups.google.com/forum/#!forum/kubernetes-sig-scale)

## Docs
- [Extending SLO definitions](extending_slo.md)
- [Scaling And Performance Goals](goals.md)

## Scalability SLOs

We officially support two different SLOs:

1. "API-responsiveness":
   99% of all API calls return in less than 1s

1. "Pod startup time:
   99% of pods (with pre-pulled images) start within 5s

This should be valid on appropriate hardware up to a 5000 node cluster with 30 pods/node.  We eventually want to expand that to 100 pods/node.

For more details how do we measure those, you can look at: http://blog.kubernetes.io/2015_09_01_archive.html

We are working on refining existing SLOs and defining more for other areas of the system.
