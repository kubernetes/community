# Event style guide

Status: During Review

Author: Marek Grabowski (gmarek@)

## Why the guide?

The Event API change proposal is the first step towards having useful Events in the system. Another step is to formalize the Event style guide, i.e. set of properties that developers need to ensure when adding new Events to the system. This is necessary to ensure that we have a system in which all components emit consistently structured Events.

## When to emit an Event?

Events are expected to provide important insights for the application developer/operator on the state of their application. Events relevant to cluster administrators are acceptable, as well, though they usually also have the option of looking at component logs. Events are much more expensive than logs, thus they're not expected to provide in-depth system debugging information. Instead concentrate on things that are important from the application developer's perspective. Events need to be either actionable, or be useful to understand past or future system's behavior. Events are not intended to drive automation. Watching resource status should be sufficient for controllers.

Following are the guidelines for adding Events to the system. Those are not hard-and-fast rules, but should be considered by all contributors adding new Events and members doing reviews.
1. Emit events only when state of the system changes/attempts to change. Events "it's still running" are not interesting. Also, changes that do not add information beyond what is observable by watching the altered resources should not be duplicated as events. Note that adding a reason for some action that can't be inferred from the state change is considered additional information.
1. Limit Events to no more than one per change/attempt. There's no need for Events on "About to do X" AND "Did X"/"Failed to do X". Result is more interesting and implies an attempt.
	1. It may give impression that this gets tricky with scale events, e.g. Deployment scales ReplicaSet which creates/deletes Pods. For us those are 3 (or more) separate Events (3 different objects are affected) so it's fine to emit multiple Events.
1. When an error occurs that prevents a user application from starting or from enacting other normal system behavior, such as object creation, an Event should be emitted (e.g. invalid image).
	1. Note that Events are garbage collected so every user-actionable error needs to be surfaced via resource status as well.
	1. It's usually OK to emit failure Events for each failure. Dedup mechanism will deal with that. The exception is failures that are frequent but typically ephemeral and automatically repairable/recoverable, such as broken socket connections, in which case they should only be reported if persistent and unrepairable, in order to mitigate event spam.
1. When a user application stops running for any reason, an Event should be emitted (e.g. Pod evicted because Node is under memory pressure)
1. If it's a system-wide change of state that may impact currently running applications or have an may have severe impact on future workload schedulability, an Event should be emitted (e.g. Node became unreachable, 1. Failed to create route for Node).
1. If it doesn't fit any of above scenarios you should consider not emitting Event.

## How to structure an Event?
New Event API tries to use more descriptive field names to influence how Events are structured. Event has following fields:
* Regarding
* Related
* ReportingController
* ReportingInstance
* Action
* Reason
* Type
* Note

The Event should be structured in a way that following sentence "makes sense":
"Regarding <Event.Regarding>: <Event.Action> <Event.Related> - <Event.Reason>", e.g.
* Regarding Node X: BecameNotReady - NodeUnreachable
* Regarding Pod X: ScheduledOnNode Node Y - <nil>
* Regarding PVC X: BoundToNode Node Y - <nil>
* Regarding Pod X: KilledContainer Container Y - NodeMemoryPressure

1. ReportingController is a type of a Controller reporting an Event, e.g. k8s.io/node-controller, k8s.io/kubelet. There will be a standard list for controller names for Kubernetes components. Third-party components must namespace themselves in the same manner as label keys. Validation ensures it's a proper qualified name. This shouldn’t be needed in order for users to understand the event, but is provided in case the controller’s logs need to be accessed for further debugging.
1. ReportingInstance is an identifier of the instance of the ReportingController which needs to uniquely identify it. I.e. host name can be used only for controllers that are guaranteed to be unique on the host. This requirement isn't met e.g. for scheduler, so it may need a secondary index. For singleton controllers use Node name (or hostname if controller is not running on the Node). Can have at most 128 alpha-numeric characters.
1. Regarding and Related are ObjectReferences. Regarding should represent the object that's implemented by the ReportingController, Related can contain additional information about another object that takes part in or is affected by the Action (see examples).
1. Action is a low-cardinality (meaning that there's a restricted, predefined set of values allowed) CamelCase string field (i.e. its value has to be determined at compile time) that explains what happened with Regarding/what action did the ReportingController take in Regarding's name.  The tuple of {ReportingController, Action, Reason} must be unique, such that a user could look up documentation. Can have at most 128 characters.
1. Reason is a low-cardinality CamelCase string field (i.e. its value has to be determined at compile time) that explains why ReportingController took Action. Can have at most 128 characters.
1. Type can be either "Normal" or "Warning". "Warning" types are reserved for Events that represent a situation that's not expected in a healthy cluster and/or healthy workload: something unexpected and/or undesirable, at least if it occurs frequently enough and/or for a long enough duration.
1. Note can contain an arbitrary, high-cardinality, user readable summary of the Event. This field can lose data if deduplication is triggered. Can have at most 1024 characters.

