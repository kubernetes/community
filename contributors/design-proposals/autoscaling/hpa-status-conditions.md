Horizontal Pod Autoscaler Status Conditions
===========================================

Currently, the HPA status conveys the last scale time, current and desired
replicas, and the last-retrieved values of the metrics used to autoscale.

However, the status field conveys no information about whether or not the
HPA controller encountered difficulties while attempting to fetch metrics,
or to scale.  While this information is generally conveyed via events,
events are difficult to use to determine the current state of the HPA.

Other objects, such as Pods, include a `Conditions` field, which describe
the current condition of the object.  Adding such a field to the HPA
provides clear indications of the current state of the HPA, allowing users
to more easily recognize problems in their setups.

API Change
----------

The status of the HPA object will gain a new field, `Conditions`, of type
`[]HorizontalPodAutoscalerCondition`, defined as follows:

```go
// HorizontalPodAutoscalerConditionType are the valid conditions of
// a HorizontalPodAutoscaler (see later on in the proposal for valid
// values)
type HorizontalPodAutoscalerConditionType string

// HorizontalPodAutoscalerCondition describes the state of
// a HorizontalPodAutoscaler at a certain point.
type HorizontalPodAutoscalerCondition struct {
    // type describes the current condition
    Type HorizontalPodAutoscalerConditionType
    // status is the status of the condition (True, False, Unknown)
    Status ConditionStatus
    // LastTransitionTime is the last time the condition transitioned from
    // one status to another
    // +optional
    LastTransitionTime metav1.Time
    // reason is the reason for the condition's last transition.
    // +optional
    Reason string
    // message is a human-readable explanation containing details about
    // the transition
    Message string
}
```

Current Conditions Conveyed via Events
--------------------------------------

The following is a list of events emitted by the HPA controller (as of the
writing of this proposal), with descriptions of the conditions which they
represent.  All of these events are caused by issues which block scaling
entirely.

- *SelectorRequired*: the target scalable resource's scale is missing
  a selector.

- *InvalidSelector*: the target scalable's selector couldn't be parsed.

- *FailedGet{Object,Pods,Resource}Metric*: the HPA controller was unable
  to fetch one metric.

- *InvalidMetricSourceType*: the HPA controller encountered an unknown
  metric source type.

- *FailedComputeMetricsReplicas*: this is fired in conjunction with one of
  the two previous events.

- *FailedConvertHPA*: the HPA controller was unable to convert the given
  HPA to the v2alpha1 version.

- *FailedGetScale*: the HPA controller was unable to actually fetch the
  scale for the given scalable resource.

- *FailedRescale*: a scale update was needed and the HPA controller was
  unable to actually update the scale subresource of the target scalable.

- *SuccessfulRescale*: a scale update was needed and everything went
  properly.

- *FailedUpdateStatus*: the HPA controller failed to update the status of
  the HPA object.

New Conditions Types
--------------------

The above conditions can be coalesced into several condition types. Each
condition has one or more associated `Reason` values which map back to
some of the events described above.

- *CanAccessScale*: this condition, when false, indicates issues actually
  getting or updating the scale of the target scalable.  Potential
  `Reason` values include `FailedGet`, `FailedUpdate`
- *InBackoff*: this condition, when true, indicates that the HPA is
  currently within a "scale forbidden window", and therefore will not
  perform scale operations in a particular direction.  Potential `Reason`
  values include `BackoffBoth`, `BackoffDownscale`, and `BackoffUpscale`.
- *CanComputeReplicas*: this condition, when false, indicates issues
  computing the desired replica counts.  Potential `Reason` values include
  `FailedGet{Object,Pods,Resource}Metric`, `InvalidMetricSourceType`, and
  `InvalidSelector` (which includes both missing and unparsable selectors,
  which can be detailed in the `Message` field).
- *DesiredOutsideRange*: this condition, when true, indicates that the
  desired scale currently would be outside the range allowed by the HPA
  spec, and is therefore capped.  Potential `Reason` values include
  `TooFewReplicas` and `TooManyReplicas`.

The `FailedUpdateStatus` event is not described here, as a failure to
update the HPA status would preclude actually conveying this information.
`FailedConvertHPA` is also not described, since it exists more as an
implementation detail of how the current mechanics of the HPA are
implemented, and less as part of the inherent functionality of the HPA
controller.

Open Questions
--------------

* Should `CanScale` be split into `CanGetScale` and `CanUpdateScale` or
  something equivalent?
