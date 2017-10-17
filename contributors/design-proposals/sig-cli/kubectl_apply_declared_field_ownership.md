# Enforcing mutually exclusive ownership of fields between apply and imperative writes

Status: Pending

## Motivation

When using `kubectl apply` to manage Kubernetes objects, the
ownership of individual fields of objects by `kubectl apply`
is mutually exclusive with imperative writes - such as `kubectl edit`,
`kubectl set`, `kubectl scale`, `kubectl replace`, programmatic clients,
and direct api usage.  Imperative writes are writes that are not
reflected back in the object configuration file used by apply.
This constraint is not enforced anywhere, nor is any warning given to
the user if they attempt to violate it.  This results in a poor user experience where
imperative changes to the object are silently dropped the
next time apply is run.

## Proposal

Enforce mutual exclusivity of ownership of fields between `kubectl apply`
and imperative writes through `kubectl`.

- a field set by `kubectl apply` will print a warning when using `kubectl set` until ownership is transferred
- a field set by `kubectl set` will print a warning when using `kubectl apply`.  warning is printed only once until after the next time it is set.
- a field set by `kubectl set` can be changed by `kubectl edit` without a warning
- a field set by `kubectl edit` can be changed by `kubectl set` without a warning
- `kubectl replace` must set the annotation if the object was created with `kubectl apply`.  A warning should be printed to stderr that the object is managed with apply.
- `kubectl apply` must set the annotation if the object does not have the `last-applied-config` annotation present.  A warning should be printed to stderr that the object was not created with apply.

Enforcement to be implemented in every command that writes to the object.  Ownership is defined
using the existing `last-applied-config` annotation on an object without any changes to the format
of the annotation.

## User Experience

### Use Cases

- *As a user, I want some pieces of a **required** field to be managed
  by apply, and others to be managed imperatively.*

**Example:** I want to create a Deployment managed by `apply`, but allow
the replica count to be managed imperatively.

1. Replica count is omitted from the configuration used by apply.
   User changes replica count using one of `kubectl edit`, `kubectl scale`,
   or `kubectl patch`.  If user adds replica count to configuration
   used by `kubectl apply` and runs the apply command, a warning is given
   to the user that the field was already set imperatively, and overwriting
   the value will change ownership of the field to kubectl apply.
2. Replica count is included in the configuration used by apply.
  User attempts to change replica count using one of `kubectl edit`, `kubectl scale`,
  or `kubectl patch`.  A warning is given
  to the user that the field was already set declaratively using kubectl
  apply, and that the user's changes will be lost the next time apply
  is run.

- *As a user, I want some pieces of an **optional** field to be managed
  by apply, and others to be managed imperatively.*

Super set of the cases for *required* fields with the addition of:

1. Clearing the optional field from apply (by deleting it from the
   configuration file or setting it to null in the configuration file)
   is successful and gives no warnings.
2. Clearing the optional field imperatively gives a warning that the field
   is managed declaratively and will be recreated by apply the next
   time apply is run.

- *As a user, I want some elements of a list to be managed by apply, and
  others to be managed imperatively.*

**Example:** I want to create a Deployment with one or more Containers in the PodTemplate.
I want apply to manage the Deployment, but imperative writes to add Containers without
having apply delete them.

Supported.  User runs `kubectl edit` and adds a Container to the PodTemplate.  Running
`kubectl apply` retains the new containers.  The containers are not present in the
local configuration used by apply.

**Example:** I want to create a Deployment with one or more Containers in the PodTemplate.
I want apply to manage all of the Containers, but I want imperative writers to
manage the Container arguments.

Supported.  User runs `kubectl edit` and adds the args field to the Container.
Running `kubectl apply` retains the args field (omitted from the configuration file).

**NOT SUPPORTED Example:** I want to create a Deployment with one or more Containers in the PodTemplate.
I want apply to manage all of the Containers and supply arguments, but I want imperative writers to
add additional arguments.

This is not supported.

1. Lists without a mergeStrategy of stategicmerge cannot have shared ownership
2. Deleting an element of a strategicmerge list from apply is always successful.
  - Elements created imperatively will not appear in the configuration file, and not easily deleted using apply.
3. Deleting an element of a strategicmerge list imperatively that is owned by apply gives a warning
  - Elements created/owned imperatively can be safely deleted
4. Adding an element to a strategicmerge list through apply is successful IFF the element is not already present (as defined by mergekey)
  - If the element is already created imperatively, give a warning
5. Adding an element to a strategicmerge list imperatively is successful IFF the element is not already present (as defined by mergekey)
6. Modifying an element in a strategicmerge list obeys the same rules as modifying an optional field (the synthetic field name is defined by the mergekey).

## Implementation

Elements appearing in the `last-applied-config` annotation are owned
by kubectl apply.

1. `kubectl apply` writes

When calculating the patch, detect conflicting field values - e.g. any
fields whose live value matches neither the value in the last-applied-config
nor the value in the configuration file being applied.  Give a warning about any conflicts.
If adding a field not present in the last-applied-config, give a strong warning
that the field maybe owned by another writert.  If updating
a field that appears in the last-applied-config to a new value, but its live value does not match the
value in the last-applied-config, give a soft warning indicating that the
value has been changed since the last time apply was run.

Example strong warning message:

> WARNING: `kubectl apply` is changing the resource <group/kind/name>
field <jsonpath> value from <x> to <y>.  This field's value does not
appear in the last configuration applied, was previously set by another
source.  This field may have been set by a controller outside of apply.

Example soft warning message:

> Warning: `kubectl apply` is changing the resource <group/kind/name>
field <jsonpath> value from <x> to <y>.  This field's value does not
match the last value set by `apply` (<z>), and another source may have
updated this field.

2. `kubectl edit` writes

When calculating the patch, check if we are patching fields that appear in
the last-applied-config.  If so, fail the edit with an error message and
provide a --force option.

3. `kubectl patch`, `kubectl set`, `kubectl scale` writes

When calculating the patch, check if we are patching fields that appear in
the last-applied-config.  If so, write a warning to stderr.

> Warning: `<command>` changed the resource <group/kind/name>
field <jsonpath> value from <x> to <y>.  This field is managed
by `kubectl apply` (last set the field to <z>), and may be overwritten or
cleared the next time `kubectl apply` is run.

### Client/Server Backwards/Forwards compatibility

TBD

## Other considerations

- What is the right level of enforcement while providing backward compatibility.
  I'd prefer to have the commands fail if they are violating constraints, but this
  will break any scripts using the existing behavior.  Introducing a command
  prompt has a similar issue.  The least invasive thing to do is provide
  a warning on stderr with rollback instructions and succeed the command anyway.
  This at least keeps backward compatibility and allows the user to undo
  their mess.  For things like edit, we can probably fail the command since
  a user is present.  Maybe we can start with a warning for 1.6 and target
  and failing the commands in 1.7?  This will give users time to root
  out issues before failing commands.

- It would be helpful to tell the user what imperative command wrote
  a field when giving a warning from apply that the field is already
  managed imperatively.  We could possibly do this through an annotation
  including the field jsonpath and kubectl command.  This would be
  set by the kubectl client.

## Alternatives considered
