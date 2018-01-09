# Volume operation metrics

## Goal

Capture high level metrics for various volume operations in Kubernetes.

## Motivation

Currently we don't have high level metrics that captures time taken
and success/failures rates of various volume operations.

This proposal aims to implement capturing of these metrics at a level
higher than individual volume plugins.

## Implementation

### Metric format and collection

Volume metrics emitted will fall under category of service metrics
as defined in [Kubernetes Monitoring Architecture](/contributors/design-proposals/instrumentation/monitoring_architecture.md).


The metrics will be emitted using [Prometheus format](https://prometheus.io/docs/instrumenting/exposition_formats/) and available for collection
from `/metrics` HTTP endpoint of kubelet and controller-manager.


Any collector which can parse Prometheus metric format should be able to collect
metrics from these endpoints.

A more detailed description of monitoring pipeline can be found in [Monitoring architecture](/contributors/design-proposals/instrumentation/monitoring_architecture.md#monitoring-pipeline) document.

### Metric Types

Since we are interested in count(or rate) and time it takes to perform certain volume operation - we will use [Histogram](https://prometheus.io/docs/practices/histograms/) type for
emitting these metrics.

We will be using `HistogramVec` type so as we can attach dimensions at runtime. All
the volume operation metrics will be named `storage_operation_duration_seconds`.
Name of operation and volume plugin's name will be emitted as dimensions. If for some reason
volume plugin's name is not available when operation is performed - label's value can be set
to `<n/a>`.


We are also interested in count of volume operation failures and hence a metric of type `NewCounterVec`
will be used for keeping track of errors. The error metric will be similarly named `storage_operation_errors_total`.

Following is a sample of metrics (not exhaustive) that will be added by this proposal:


```
storage_operation_duration_seconds { volume_plugin = "aws-ebs", operation_name = "volume_attach" }
storage_operation_duration_seconds { volume_plugin = "aws-ebs", operation_name = "volume_detach" }
storage_operation_duration_seconds { volume_plugin = "glusterfs", operation_name = "volume_provision" }
storage_operation_duration_seconds { volume_plugin = "gce-pd", operation_name = "volume_delete" }
storage_operation_duration_seconds { volume_plugin = "vsphere", operation_name = "volume_mount" }
storage_operation_duration_seconds { volume_plugin = "iscsi" , operation_name = "volume_unmount" }
storage_operation_duration_seconds { volume_plugin = "aws-ebs", operation_name = "unmount_device" }
storage_operation_duration_seconds { volume_plugin = "cinder" , operation_name = "verify_volumes_are_attached" }
storage_operation_duration_seconds { volume_plugin = "<n/a>" , operation_name = "verify_volumes_are_attached_per_node" }
```

Similarly errors will be named:

```
storage_operation_errors_total { volume_plugin = "aws-ebs", operation_name = "volume_attach" }
storage_operation_errors_total { volume_plugin = "aws-ebs", operation_name = "volume_detach" }
storage_operation_errors_total { volume_plugin = "glusterfs", operation_name = "volume_provision" }
storage_operation_errors_total { volume_plugin = "gce-pd", operation_name = "volume_delete" }
storage_operation_errors_total { volume_plugin = "vsphere", operation_name = "volume_mount" }
storage_operation_errors_total { volume_plugin = "iscsi" , operation_name = "volume_unmount" }
storage_operation_errors_total { volume_plugin = "aws-ebs", operation_name = "unmount_device" }
storage_operation_errors_total { volume_plugin = "cinder" , operation_name = "verify_volumes_are_attached" }
storage_operation_errors_total { volume_plugin = "<n/a>" , operation_name = "verify_volumes_are_attached_per_node" }
```

### Implementation Detail

We propose following changes as part of implementation details.

1. All volume operations are executed via `goroutinemap.Run` or `nestedpendingoperations.Run`.
`Run` function interface of these two types can be changed to include a `operationComplete` callback argument.

   For example:

   ```go
   // nestedpendingoperations.go
   Run(v1.UniqueVolumeName, types.UniquePodName, func() error, opComplete func(error)) error
   // goroutinemap
   Run(string, func() error, opComplete func(error)) error
   ```

   This will enable us to know when a volume operation is complete.

2. All `GenXXX` functions in `operation_generator.go` should return plugin name in addition to function and error.

   for example:

   ```go
   GenerateMountVolumeFunc(waitForAttachTimeout time.Duration,
       volumeToMount VolumeToMount,
       actualStateOfWorldMounterUpdater
       ActualStateOfWorldMounterUpdater, isRemount bool) (func() error, pluginName string, err error)
   ```

   Similarly `pv_controller.scheduleOperation` will take plugin name as additional parameter:

   ```go
   func (ctrl *PersistentVolumeController) scheduleOperation(
           operationName string,
           pluginName string,
           operation func() error)
   ```

3. Above changes will enable us to gather required metrics in `operation_executor` or when scheduling a operation in
pv controller.

   For example, metrics for time it takes to attach Volume can be captured via:

   ```go
   func operationExecutorHook(plugin, operationName string) func(error) {
       requestTime := time.Now()
       opComplete := func(err error) {
         timeTaken := time.Since(requestTime).Seconds()
         // Create metric with operation name and plugin name
       }
       return onComplete
    }
   attachFunc, plugin, err :=
           oe.operationGenerator.GenerateAttachVolumeFunc(volumeToAttach, actualStateOfWorld)
   opCompleteFunc := operationExecutorHook(plugin, "volume_attach")
   return oe.pendingOperations.Run(
           volumeToAttach.VolumeName, "" /* podName */, attachFunc, opCompleteFunc)
   ```

   `operationExecutorHook` function is a hook that is registered in operation_executor and it will
   initialize necessary metric params and will return a function. This will be called when
   operation is complete and will finalize metric creation and finally emit the metrics.

### Conclusion

Collection of metrics at operation level ensures almost no code change to volume plugin interface and a very minimum change to controllers.
