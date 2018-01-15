# sig-scalability charter
Author: @porridge (with help from @wasylkowski, @wojtek-t, @shyamjvs, @gmarek)

*This document describes the role, duties, and rights of sig-scalability.*

## Who are we?

sig-scalability helps define kubernetes scalability goals, and makes sure that every kubernetes release meets them by measuring performance/scalability indicators and publishing the results. We also coordinate and contribute to general system-wide scalability/performance improvements (ones which do not fall into the charter of another individual SIG).

We provide performance measurement tools, processes for ensuring continued scalability, support and consultation to kubernetes developers and sig-release members in order to meet the above goals.

## SIG Values
* We are not firefighters, we are fire-prevention specialists.
* We promote deep technical understanding of the kubernetes system and our tools.
* We strive to [eliminate toil](https://landing.google.com/sre/book/chapters/eliminating-toil.html).
* We work towards building a scalable kubernetes even in face of superlinear growth of number of contributions.

## How do we invest our time
Each SIG member should strive to spend _no more than_:
* 25% of time on **Debugging**:
   * Investigate performance regressions and scalability bottlenecks.
   * Keeping a cap on how much time we spend in this area:
      * prevents the risk that regressions overwhelm us,
      * will reliably give us time to invest in other areas (for a long-term improvement),
      * when we hit the cap, it will provide visibility outside of the SIG about problems caused by insufficiently tested features/PRs.
   * 25% is temporary until situation stabilizes. We have a long-term plan to lower this to 10% by the end of 2018.
* 40% of time on **Tools and Processes**:
   * Define and maintain scalability SLI/SLO definition
   * Design, develop and maintain testing infrastructure that makes it very easy to write scalability tests to evaluate the above SLI
   * Improve tools and infrastructure to make it reliable and easy for the community members to run performance tests of their own features before they hit the repo
   * Provide a "scalability validation report" for each release which:
      * describes test environment and enumerates used configuration parameters,
      * enumerates which kubernetes features are known not to scale.
* 25% of time on **Consulting**:
   * Provide advice, consultation and support for kubernetes developers regarding scalability and performance bottlenecks.
* 10% of time on **Performance improvements**:
   * Actually designing and implementing performance improvements to the system.


### Notes about the time allocation split above:
* the numbers are meant as a reminder that balance is necessary to make sure that we are not completely overwhelmed by debugging, you do NOT need to do per-minute accounting of your time,
* the split is expected to change as our tools improve and flakiness decreases - we hope to devote less time to fire-fighting and more time to performance improvements in the long term.

## What can we do/require from other SIGs

*By regression below we mean a regression identified by the set of release-blocking scalability/performance tests (as defined by [the sig-release-x.y-blocking dashboards](https://github.com/kubernetes/test-infra/blob/master/testgrid/config/config.yaml)).*

* We require that significant changes (in terms of impact on the whole system, not in terms of line count), such as: update of Go version, update of etcd, major impactful architectural changes etc. may only be merged:
   * with an explicit approval from a SIG-scalability approver, and
   * after having passed performance testing on biggest supported clusters (unless the scalability approver explicitly states this is not necessary).
* We can block a feature from transitioning to beta status if (when turned on) it causes a significant degradation of other components in terms of scalability/performance. (To judge what is "significant" ideally we would want an objective way such as "breaks an SLO" or "causes SLI degradation of more than XX%". But we are not there yet, so initially this will be sig-scalability team's call.)
* We can block a feature from transitioning to GA status if it cannot be used at scale.
* We can roll back any PR merged to master if it has been identified as a cause of an SLO regression. The offending PR should only be merged again after it has been shown to pass tests at scale.
* We can pause the merge queue in case a regression on master has been observed, until a particular PR has been *identified* as cause of the regression *and* the regression has been *mitigated* (see below for what it means).
* We can require a SIG to introduce a regression-catching benchmark test for a scalability-critical functionality.

"Rules of engagement" for pausing merge queue:
1. observe a scalability regression on master,
2. pause the merge queue,
3. identify the PR which caused the regression
   1. this can be done by reading code changes, bisecting, running A/B tests, etc,
   2. we say a PR is identified as the cause when we are reasonably confident that it indeed caused a regression, *even if the mechanism is not 100% understood* - this is because we want to minimize the time of merge queue pause
4. mitigate the regression; this can potentially be different things, e.g.:
   1. reverting the PR,
   2. switching some feature off (preferably by default, at last resort in the tests only),
   3. downgrading Go version,
   4. etc
5. once the mitigation is merged to master, unpause the merge queue

### Rationale
The above are quite drastic measures, but we believe they are justified if we want kubernetes to maintain scalability SLOs. The reasoning is:
   * reliably testing for regressions takes a lot of time:
     * key scalability e2e tests take too long to execute to be a prerequisite for merging all PRs, this is an inherent characteristic of testing at scale,
     * end-to-end tests are flaky (even when not at scale) requiring retries,
   * we need to prevent regression pile-ups:
     * once a regression is merged, and no other action is taken, it is only a matter of time until another regression is merged on top of it,
     * debugging the cause of two simultaneous (piled-up) regressions is exponentially harder, see [issue 53255](https://github.com/kubernetes/kubernetes/issues/53255) which links to past experience
   * regarding benchmarks, there were several scalability issues in the past caught by (costly) large-scale e2e tests, which could have been caught and fixed earlier and with far less human effort if we had benchmark-like tests. Examples include:
     * scheduler anti-affinity affecting kube-dns,
     * kubelet network plugin increasing pod-startup latency,
     * large responses from apiserver violating gRPC MTU.

As [explained in detail in an issue](https://github.com/kubernetes/kubernetes/issues/53255#issue-261658836), not being able to maintain passing scalability tests adversely affect:
   * release quality
   * release schedule
   * engineer productivity
