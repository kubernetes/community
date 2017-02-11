# Scalability SIG

**Leads:** Bob Wise (@countspongebob) and Joe Beda (@jbeda)

**Slack Channel:** [#sig-scale](https://kubernetes.slack.com/messages/sig-scale/).  [Archive](http://kubernetes.slackarchive.io/sig-scale/)

**Mailing List:** [kubernetes-sig-scale](https://groups.google.com/forum/#!forum/kubernetes-sig-scale)

**Meetings:** Thursdays at 9am pacific.  Contact Joe or Bob for invite. [Notes](https://docs.google.com/a/bobsplanet.com/document/d/1hEpf25qifVWztaeZPFmjNiJvPo-5JX1z0LSvvVY5G2g/edit?usp=drive_web
)

**Docs:**
[Scaling And Performance Goals](goals.md)

### Scalability SLAs

We officially support two different SLAs:

1. "API-responsiveness":
   99% of all API calls return in less than 1s

1. "Pod startup time:
   99% of pods (with pre-pulled images) start within 5s

This should be valid on appropriate hardware up to a 1000 node cluster with 30 pods/node.  We eventually want to expand that to 100 pods/node.

For more details how do we measure those, you can look at: http://blog.kubernetes.io/2015_09_01_archive.html

In the future we may want to add more SLAs (e.g. scheduler throughput), but we are not there yet.
