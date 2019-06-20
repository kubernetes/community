# Scheduler Benchmarking

Kubernetes scheduler has integration and e2e benchmarks. It is recommended to
run integration benchmarks at the time of submitting PRs that could potentially
impact performance of the scheduler and give before and after results with your
PR submission.

## Running integration benchmarks

To run integration benchmarks use the following command from inside a Kubernetes
directory. 

```sh
make test-integration WHAT=./test/integration/scheduler_perf KUBE_TEST_VMODULE="''" KUBE_TEST_ARGS="-run=^$$ -bench=."
```

You can also provide a benchmark name in order to run a specific set of
benchmarks. Please refer to [Go documentation on benchmarks](https://golang.org/pkg/testing/#hdr-Benchmarks)
for more information.

```sh
make test-integration WHAT=./test/integration/scheduler_perf KUBE_TEST_VMODULE="''" KUBE_TEST_ARGS="-run=^$$ -bench=BenchmarkScheduling"
```

> To display benchmark output only, you can append `-alsologtostderr=false -logtostderr=false` to `KUBE_TEST_ARGS`.

These benchmarks are located in `./test/integration/scheduler_perf/scheduler_bench_test.go`. 
The function names start with `BenchmarkScheduling`. At the beginning of each
function there are a few lines in the form of:

```go
	tests := []struct{ nodes, existingPods, minPods int }{
		{nodes: 100, existingPods: 1000, minPods: 100},
		{nodes: 1000, existingPods: 1000, minPods: 100},
		{nodes: 5000, existingPods: 1000, minPods: 1000},
	}
```

Each line indicates a test. `nodes` specifies the number of nodes in the test
cluster. `existingPods` specifies how many pods should be created and scheduled
as the initialization phase of the benchmark. `minPods` specifies how many pods
must be created and scheduled as the actual part of benchmarking.
You may add other items to the above array and run the benchmarks again. For 
example, if you want to measure performance in a 5000 node cluster when scheduling
10000 pods, you can add the following to the `tests` array.

```go
{nodes: 5000, existingPods: 1000, minPods: 10000},
```

You can also run a particular configuration from the above `tests` by specifying
it in the `-bench` argument. For example, the following will run only those
benchmarks with 5000 nodes and 1000 existing pods.

```sh
make test-integration WHAT=./test/integration/scheduler_perf KUBE_TEST_VMODULE="''" KUBE_TEST_ARGS="-run=^$$ -bench=BenchmarkScheduling/5000Nodes/1000Pods"
```

## Profiling the scheduler

You can get CPU profiling information from the benchmarks by adding a `-cpuprofile`
to the command above. Example:

```sh
make test-integration WHAT=./test/integration/scheduler_perf KUBE_TEST_VMODULE="''" KUBE_TEST_ARGS="-run=^$$ -bench=BenchmarkScheduling -cpuprofile cpu.out"
```

After obtaining the CPU profile, you can use `pprof` to view the results. For
example, you can use the following command to see the output in your browser: 

```sh
go tool pprof -web cpu.out
```

You may also want to read more about profiling [here](../sig-scalability/profiling.md).

## End to end performance numbers

The scheduler has end to end real cluster benchmarks that run regularly. These
benchmarks are available on our [perf dashboard](http://perf-dash.k8s.io/). Once
on the perf dashboard, you can choose various sizes of clusters to see the
scheduler's performance. For example, you can choose
`gce-5000Nodes > Scheduler > SchedulingThroughput` to see the scheduler's
throughput for a 5000-node cluster.

Note that the API server rate limitation is set to 100 requests per second. As
a result, the scheduler's throughput is capped at 100 pods/s.
