# Cloud Provider (specifically GCE and AWS) metrics for Storage API calls

## Goal

Kubernetes should provide metrics such as - count & latency percentiles
for cloud provider API it uses to provision persistent volumes.

In a ideal world - we would want these metrics for all cloud providers
and for all API calls kubernetes makes but to limit the scope of this feature
we will implement metrics for:

* GCE
* AWS

We will also implement metrics only for storage API calls for now. This feature
does introduces hooks into kubernetes code which can be used to add additonal metrics
but we only focus on storage API calls here.

## Motivation

* Cluster admins should be able to monitor Cloud API usage of Kubernetes. It will help
  them detect problems in certain scenarios which can blow up the API quota of Cloud
  provider.
* Cluster admins should also be able to monitor health and latency of Cloud API on
  which kubernetes depends on.

## Implementation

### Metric format and collection

Metrics emitted from cloud provider will fall under category of service metrics
as defined in [Kubernetes Monitoring Architecture](https://github.com/kubernetes/community/blob/master/contributors/design-proposals/monitoring_architecture.md).


The metrics will be emitted using [Prometheus format](https://prometheus.io/docs/instrumenting/exposition_formats/) and available for collection
from `/metrics` HTTP endpoint of kubelet, controller etc. All Kubernetes core components already emit
metrics on `/metrics` HTTP endpoint. This proposal merely extends available metrics to include Cloud provider metrics as well.


Any collector which can parse Prometheus metric format should be able to collect
metrics from these endpoints.

A more detailed description of monitoring pipeline can be found in [Monitoring architecture] (https://github.com/kubernetes/community/blob/master/contributors/design-proposals/monitoring_architecture.md#monitoring-pipeline) document.


#### Metric Types

Since we are interested in count(or rate) and latency percentile metrics of API calls Kubernetes is making to
the external Cloud Provider - we will use [Histogram](https://prometheus.io/docs/practices/histograms/) type for
emitting these metrics.

We will be using `HistogramVec` type so as we can attach dimensions at runtime. Whenever available
`namespace` will reported as a dimension with the metric.

### GCE Implementation

For GCE we simply use `gensupport.RegisterHook()` to register a function which will be called
when request is made and response returns.

To begin with we will start emitting following metrics for GCE. Because these metrics are of type
`Summary` - both count and latency will be automatically calculated.

1. gce_instance_list
2. gce_disk_insert
3. gce_disk_delete
4. gce_attach_disk
5. gce_detach_disk
6. gce_list_disk

A POC implementation can be found here - https://github.com/kubernetes/kubernetes/pull/40338/files

### AWS Implementation

For AWS currently we will use wrapper type `awsSdkEC2` to intercept all storage API calls and
emit metric datapoints.  The reason we are not using approach used for `aws/log_handler` is - because AWS SDK doesn't uses Contexts and hence we can't pass custom information such as API call name or namespace to record with metrics.

To begin with we will start emitting following metrics for AWS:

1. aws_attach_volume
2. aws_create_tags
3. aws_create_volume
4. aws_delete_volume
5. aws_describe_instance
6. aws_describe_volume
7. aws_detach_volume
