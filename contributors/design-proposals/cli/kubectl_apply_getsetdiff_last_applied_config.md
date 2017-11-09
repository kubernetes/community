# Kubectl apply subcommands for last-config

## Abstract

`kubectl apply` uses the `last-applied-config` annotation to compute
the removal of fields from local object configuration files and then
send patches to delete those fields from the live object.  Reading or
updating the `last-applied-config` is complex as it requires parsing
out and writing to the annotation.  Here we propose a set of porcelain
commands for users to better understand what is going on in the system
and make updates.

## Motivation

What is going on behind the scenes with `kubectl apply` is opaque.  Users
have to interact directly with annotations on the object to view
and make changes.  In order to stop having `apply` manage a field on
an object, it must be manually removed from the annotation and then be removed
from the local object configuration.  Users should be able to simply edit
the local object configuration and set it as the last-applied-config
to be used for the next diff base.  Storing the last-applied-config
in an annotation adds black magic to `kubectl apply`, and it would
help users learn and understand if the value was exposed in a discoverable
manner.

## Use Cases

1. As a user, I want to be able to diff the last-applied-configuration
   against the current local configuration to see which changes the command is seeing
2. As a user, I want to remove fields from being managed by the local
   object configuration by removing them from the local object configuration
   and setting the last-applied-configuration to match.
3. As a user, I want to be able to view the last-applied-configuration
   on the live object that will be used to calculate the diff patch
   to update the live object from the configuration file.

## Naming and Format possibilities

### Naming

1. *cmd*-last-applied

Rejected alternatives:

2. ~~last-config~~
3. ~~last-applied-config~~
4. ~~last-configuration~~
5. ~~last-applied-configuration~~
6. ~~last~~

### Formats

1. Apply subcommands
  - `kubectl apply set-last-applied/view-last-applied/diff-last-applied
  - a little bit odd to have 2 verbs in a row
  - improves discoverability to have these as subcommands so they are tied to apply

Rejected alternatives:

2. ~~Set/View subcommands~~
  - `kubectl set/view/diff last-applied
  - consistent with other set/view commands
  - clutters discoverability of set/view commands since these are only for apply
  - clutters discoverability for last-applied commands since they are for apply
3. ~~Apply flags~~
  - `kubectl apply [--set-last-applied | --view-last-applied | --diff-last-applied]
  - Not a fan of these

## view last-applied

Porcelain command that retrieves the object and prints the annotation value as yaml or json.

Prints an error message if the object is not managed by `apply`.

1. Get the last-applied by type/name

```sh
kubectl apply view-last-applied deployment/nginx
```

```yaml
apiVersion: extensions/v1beta1
kind: Deployment
metadata:
  name: nginx
spec:
  replicas: 1
  template:
    metadata:
      labels:
        run: nginx
    spec:
      containers:
      - image: nginx
        name: nginx
```

2. Get the last-applied by file, print as json

```sh
kubectl apply view-last-applied -f deployment_nginx.yaml -o json
```

Same as above, but in json

## diff last-applied

Porcelain command that retrieves the object and displays a diff against
the local configuration

1. Diff the last-applied

```sh
kubectl apply diff-last-applied -f deployment_nginx.yaml
```

Opens up a 2-way diff in the default diff viewer.  This should
follow the same semantics as `git diff`.  It should accept either a
flag `--diff-viewer=meld` or check the environment variable
`KUBECTL_EXTERNAL_DIFF=meld`.  If neither is specified, the `diff`
command should be used.

This is meant to show the user what they changed in the configuration,
since it was last applied, but not show what has changed in the server.

The supported output formats should be `yaml` and `json`, as specified
by the `-o` flag.

A future goal is to provide a 3-way diff with `kubectl apply diff -f deployment_nginx.yaml`.
Together these tools would give the user the ability to see what is going
on and compare changes made to the configuration file vs other
changes made to the server independent of the configuration file.

## set last-applied

Porcelain command that sets the last-applied-config annotation to as
if the local configuration file had just been applied.

1. Set the last-applied-config

```sh
kubectl apply set-last-applied -f deployment_nginx.yaml
```

Sends a Patch request to set the last-applied-config as if
the configuration had just been applied.

## edit last-applied

1. Open the last-applied-config in an editor

```sh
kubectl apply edit-last-applied -f deployment_nginx.yaml
```

Since the last-applied-configuration annotation exists only
on the live object, this command can alternatively take the
kind/name.

```sh
kubectl apply edit-last-applied deployment/nginx
```

Sends a Patch request to set the last-applied-config to
the value saved in the editor.

## Example workflow to stop managing a field with apply - using get/set

As a user, I want to have the replicas on a Deployment managed by an autoscaler
instead of by the configuration.

1. Check to make sure the live object is up-to-date
  - `kubectl apply diff-last-applied -f deployment_nginx.yaml`
  - Expect no changes
2. Update the deployment_nginx.yaml by removing the replicas field
3. Diff the last-applied-config to make sure the only change is the removal of the replicas field
4. Remove the replicas field from the last-applied-config so it doesn't get deleted next apply
  - `kubectl apply set-last-applied -f deployment_nginx.yaml`
5. Verify the last-applied-config has been updated
  - `kubectl apply view-last-applied -f deployment_nginx.yaml`

## Example workflow to stop managing a field with apply - using edit

1. Check to make sure the live object is up-to-date
  - `kubectl apply diff-last-applied -f deployment_nginx.yaml`
  - Expect no changes
2. Update the deployment_nginx.yaml by removing the replicas field
3. Edit the last-applied-config and remove the replicas field
  - `kubectl apply edit-last-applied deployment/nginx`
4. Verify the last-applied-config has been updated
  - `kubectl apply view-last-applied -f deployment_nginx.yaml`

