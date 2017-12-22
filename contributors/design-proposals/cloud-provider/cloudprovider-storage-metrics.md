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
does introduces hooks into kubernetes code which can be used to add additional metrics
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
as defined in [Kubernetes Monitoring Architecture](/contributors/design-proposals/instrumentation/monitoring_architecture.md).


The metrics will be emitted using [Prometheus format](https://prometheus.io/docs/instrumenting/exposition_formats/) and available for collection
from `/metrics` HTTP endpoint of kubelet, controller etc. All Kubernetes core components already emit
metrics on `/metrics` HTTP endpoint. This proposal merely extends available metrics to include Cloud provider metrics as well.


Any collector which can parse Prometheus metric format should be able to collect
metrics from these endpoints.

A more detailed description of monitoring pipeline can be found in [Monitoring architecture] (/contributors/design-proposals/instrumentation/monitoring_architecture.md#monitoring-pipeline) document.


#### Metric Types

Since we are interested in count(or rate) and latency percentile metrics of API calls Kubernetes is making to
the external Cloud Provider - we will use [Histogram](https://prometheus.io/docs/practices/histograms/) type for
emitting these metrics.

We will be using `HistogramVec` type so as we can attach dimensions at runtime. All metrics will contain API action
being taken as a dimension. The cloudprovider maintainer may choose to add additional dimensions as needed. If a
dimension is not available at point of emission sentinel value `<n/a>` should be emitted as a placeholder.

We are also interested in counter of cloudprovider API errors. `NewCounterVec` type will be used for keeping
track of API errors.

### GCE Implementation

To begin with we will start emitting following metrics for GCE. Because these metrics are of type
`Histogram` - both count and latency will be automatically calculated.

#### GCE Latency metrics

All gce latency metrics will be named - `cloudprovider_gce_api_request_duration_seconds`. api request
being made will be reported as dimensions.


To begin we will start emitting following metrics:

```
cloudprovider_gce_api_request_duration_seconds { request = "instance_list"}
cloudprovider_gce_api_request_duration_seconds { request = "disk_insert"}
cloudprovider_gce_api_request_duration_seconds { request = "disk_delete"}
cloudprovider_gce_api_request_duration_seconds { request = "attach_disk"}
cloudprovider_gce_api_request_duration_seconds { request = "detach_disk"}
cloudprovider_gce_api_request_duration_seconds { request = "list_disk"}
```

#### GCE API error metrics.

All gce error metrics will be named `cloudprovider_gce_api_request_errors`. api request being made will be
reported as a dimension.

To begin with we expect to report following error metrics:

```
cloudprovider_gce_api_request_errors { request = "instance_list"}
cloudprovider_gce_api_request_errors { request = "disk_insert"}
cloudprovider_gce_api_request_errors { request = "disk_delete"}
cloudprovider_gce_api_request_errors { request = "attach_disk"}
cloudprovider_gce_api_request_errors { request = "detach_disk"}
cloudprovider_gce_api_request_errors { request = "list_disk"}
```


### AWS Implementation

For AWS currently we will use wrapper type `awsSdkEC2` to intercept all storage API calls and
emit metric datapoints.  The reason we are not using approach used for `aws/log_handler` is - because AWS SDK doesn't uses Contexts and hence we can't pass custom information such as API call name or namespace to record with metrics.


#### AWS Latency metrics

All aws API metrics will be named - `cloudprovider_aws_api_request_duration_seconds`. `request` will be reported as dimensions.
AWS maintainer may choose to add additional dimensions as needed.

To begin with we will start emitting following metrics for AWS:

```
cloudprovider_aws_api_request_duration_seconds { request = "attach_volume"}
cloudprovider_aws_api_request_duration_seconds { request = "detach_volume"}
cloudprovider_aws_api_request_duration_seconds { request = "create_tags"}
cloudprovider_aws_api_request_duration_seconds { request = "create_volume"}
cloudprovider_aws_api_request_duration_seconds { request = "delete_volume"}
cloudprovider_aws_api_request_duration_seconds { request = "describe_instance"}
cloudprovider_aws_api_request_duration_seconds { request = "describe_volume"}
```

#### AWS Error metrics

All aws error metrics will be named `cloudprovider_aws_api_request_errors`. api request being made will be
reported as a dimension.

To begin with we expect to report following error metrics:

```
cloudprovider_aws_api_request_errors { request = "attach_volume"}
cloudprovider_aws_api_request_errors { request = "detach_volume"}
cloudprovider_aws_api_request_errors { request = "create_tags"}
cloudprovider_aws_api_request_errors { request = "create_volume"}
cloudprovider_aws_api_request_errors { request = "delete_volume"}
cloudprovider_aws_api_request_errors { request = "describe_instance"}
cloudprovider_aws_api_request_errors { request = "describe_volume"}
```
