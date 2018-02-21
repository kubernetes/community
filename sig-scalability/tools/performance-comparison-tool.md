# Performance Comparison Tool

_by Shyam JVS, Google Inc (reviewed by Marek Grabowski & Wojciech Tyczysnki)_

## BACKGROUND

Kubemark is a performance testing tool which we use to run simulated kubernetes clusters. The primary use case is scalability testing, as creating simulated clusters is faster and requires less resources than creating real ones. For more information about kubemark, take a look into the [doc](https://github.com/kubernetes/community/blob/master/contributors/devel/kubemark-guide.md).

## OBJECTIVE

After recent updates, kubemark caught up with the setup of real clusters w.r.t. having all core cluster components, some add-ons and some daemons. Now we want to be able to say confidently that kubemark really reflects performance problems/bottlenecks in real clusters. Thus, our goals are to:

- Make kubemark mimic real clusters in performance enough to allow reasoning about meeting performance SLOs using kubemark runs.
- Formalize the notion of “similar performance” and set up a tool for doing the comparison.

## DESIGN OVERVIEW

We assume that we want to benchmark a test T across two variants A and B. For the benchmarking to be meaningful, these two variants should be running in a similar environment (eg. one in real-cluster and one in kubemark) and at identical scale (eg. both run 2k nodes). At a high-level, the tool should:

- *choose the set of runs* of tests T executed on both A and B environments to use for comparison,
- *obtain the relevant metrics* for the runs chosen for comparison,
- *compute the similarity* for each individual metric, across both the samples,
- *compute overall similarity* of A and B, using similarity values of all metrics.

Final output of the tool will be the answer to the question "are environments A and B similar enough with respect to chosen metrics" given some notion of similarity. The result will contain similarity measure for each metric and a similarity measure for the whole test. E.g.

```
Performance comparison results:
API call latencies:
GET Pod: 0.95
PUT Pod: 0.92
...
E2e Pod startup latencies: 0.99
Total similarity measure: 0.95
```

## DESIGN DETAILS

Performance Comparison Tool's infrastructure is designed to be easily extensible and portable. It'll allow for easy modification/extension of default logic and it'll be possible to run it on any environment that can build go binaries, and have access to relevant data.

It'll consist of a single binary that will be able to read series of test results either from Google Cloud Storage buckets or from local disk, extract relevant metrics from those results, compute given similarity function for all metrics, and finally combine those similarities in the final result.

Moving parts of the system are:

- tests to compare (including the source: local or GCS)
- set of metrics to compare
- definition of similarity measure for single metrics
- definition of similarity measure for whole test (combined metrics)

Below we discuss default choices we made.

### Choosing tests to compare

When running the comparison we need to decide on which tests to include and how to get their data. In the first version of our comparison tool we support only GCS and local sources for data with a well defined structure. We expect to have a bucket/directory with results for each run of the test. Each of those subdirectories need to have dedicated files for metrics for those runs in some well-defined format (like json). We'll expose a flag that'll allow for choosing only a subset of runs (subdirectories) to read.

By default we'll use GCS source and the last ‘n’ (TBD) runs of either tests for comparing. ‘n’ will be a configurable parameter of the tool as it could vary depending on the requirements of various tests.

### Choosing set of metrics to compare

User will be able to support a set of metrics to include into comparison. The only requirement for those metrics is to have a single numeric value.

In the initial version we'll default to the following metrics that are most directly visible in k8s performance:

- Percentiles (90%, 95%, 99%) of pod startup latency
- Percentiles (90%, 95%, 99%) of api request serving latency (split by resource and verb)

The framework itself will be extensible, so it'll be easy to add new metrics. In the future we plan to add:

- etcd request latencies
- pod startup throughput
- resource usage of control-plane components

Because performance results tend to vary a lot, especially when metrics are small (e.g. API call latencies of low tens of milliseconds) due to various processes that happen on the machine (most notably go garbage collector running), before doing any comparison we need to reduce the noise in the data. Those normalization procedures will be defined in the code for each supported metric. For initial ones we're going to set a cutoff threshold and substitute all values smaller than it with the threshold:

- for API call latencies it'll be 50ms
- for e2e Pod startup latency it'll be 1s

### What do we mean by similarity between single metric series?

For each metric we're considering we'll get a series of results, which we'll treat as a series of experiments from a single probability distribution. We have one such series for either tests we want to compare. The question we want to answer is whether their underlying distributions are "similar enough".

Initially we’ll use a simple test to determine if the metrics are similar. We find the ratio of the metric’s averages from either series and check if that ratio is in the interval \[0.66, 1.50\] (i.e. one does not differ from the other by more than 33%). We deem the metric as matched if and only if It lies in the interval. We can switch to more advanced statistical tests on distributions of these metrics in future if needed.

### What do we mean by similarity for whole test?

Once we have calculated the similarity measures for all the metrics in the metrics set, we need to decide how to compute combined similarity score.

We classify each metric as matched/mismatched based on the above test. We could have more than just binary classification in future if needed. Finally, the overall comparison result would be computed as:

- PASS, if at least 90% of the metrics in our set matched
- FAIL, otherwise

(Note: If a metric consistently mismatches across multiple rounds of comparison, it needs fixing)

## RELEVANCE & SCOPE

- This tool can benefit the community in the following ways:
  - Having this tool as open-source would make the process of testing on simulated clusters and claims about performance on real clusters using performance on simulated clusters more clear and transparent.
  - Since performance on simulated clusters indicates the kubernetes side of performance rather than that on the side of the underlying provider infra, it can help the community / kubernetes providers be assured that there indeed are no scalability problems on the side of kubernetes.
- This tool can be extended in future for other use cases like:
  - Compare two different samples of runs from the same test to see which metrics have improved / degraded over time.
  - Run comparison with more advanced statistical tests that validate hypotheses about similarity of the underlying distributions of the metric series and see if the distributions follow some known family of distribution functions.

--------------------------

**NOTES FOR INTERESTED CONTRIBUTORS**

This tool has been implemented and the code for it lies [here](https://github.com/kubernetes/perf-tests/tree/master/benchmark). Further, we have setup an [automated CI job](https://k8s-testgrid.appspot.com/perf-tests#kubemark-100-benchmark) that runs this benchmark periodically and compares the metrics across our 100-node kubemark and 100-node real-cluster runs from the last 24 hrs.

If you want to contribute to this tool, file bugs or help with understanding/resolving differences we’re currently observing across kubemark and real-cluster (e.g [#44701](https://github.com/kubernetes/kubernetes/issues/44701)), ping us on “sig-scale” kubernetes slack channel and/or write an email to `kubernetes-sig-scale@googlegroups.com`.

We have some interesting challenges in store for you, that span multiple parts of the system.
