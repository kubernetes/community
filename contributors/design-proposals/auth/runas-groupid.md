# RunAsGroup Proposal

**Author**: krmayankk@

**Status**: Proposal

## Abstract


As a Kubernetes User, we should be able to specify both user id and group id for the containers running 
inside a pod on a per Container basis, similar to how docker allows that using docker run options `-u, 
--user="" Username or UID (format: <name|uid>[:<group|gid>]) format`.

PodSecurityContext allows Kubernetes users to specify RunAsUser which can be overridden by RunAsUser
in SecurityContext on a per Container basis. There is no equivalent field for specifying the primary
Group of the running container.

## Motivation

Enterprise Kubernetes users want to run containers as non root. This means running containers with a 
non zero user id and non zero primary group id. This gives Enterprises, confidence that their customer code
is running with least privilege and if it escapes the container boundary, will still cause least harm
by decreasing the attack surface.

### What is the significance of Primary Group Id ?
Primary Group Id is the group id used when creating files and directories. It is also the default group 
associated with a user, when he/she logins. All groups are defined in `/etc/group` file and are created
with the `groupadd` command. A Process/Container runs with uid/primary gid of the calling user. If no
primary group is specified for a user, 0(root) group is assumed. This means , any files/directories created
by a process running as user with no primary group associated with it, will be owned by group id 0(root).

## Goals

1. Provide the ability to specify the Primary Group id for a container inside a Pod
2. Bring launching of containers using Kubernetes at par with Dockers by supporting the same features.


## Use Cases

### Use case 1:
As a Kubernetes User, I should be able to control both user id and primary group id of containers 
launched using Kubernetes at runtime, so that i can run the container as non root with least possible
privilege.

### Use case 2:
As a Kubernetes User, I should be able to control both user id and primary group id of containers 
launched using Kubernetes at runtime, so that i can override the user id and primary group id specified
in the Dockerfile of the container image, without having to create a new Docker image.

## Design

### Model

Introduce a new API field in SecurityContext and PodSecurityContext called `RunAsGroup`.

#### SecurityContext

```
// SecurityContext holds security configuration that will be applied to a container.
// Some fields are present in both SecurityContext and PodSecurityContext.  When both
// are set, the values in SecurityContext take precedence.
type SecurityContext struct {
     //Other fields not shown for brevity
    ..... 

     // The UID to run the entrypoint of the container process.
     // Defaults to user specified in image metadata if unspecified.
     // May also be set in PodSecurityContext.  If set in both SecurityContext and
     // PodSecurityContext, the value specified in SecurityContext takes precedence.
     // +optional
     RunAsUser *int64
     // The GID to run the entrypoint of the container process.
     // Defaults to group specified in image metadata if unspecified.
     // May also be set in PodSecurityContext.  If set in both SecurityContext and
     // PodSecurityContext, the value specified in SecurityContext takes precedence.
     // +optional
     RunAsGroup *int64
     // Indicates that the container must run as a non-root user.
     // If true, the Kubelet will validate the image at runtime to ensure that it
     // does not run as UID 0 (root) and fail to start the container if it does.
     // If unset or false, no such validation will be performed.
     // May also be set in SecurityContext.  If set in both SecurityContext and
     // PodSecurityContext, the value specified in SecurityContext takes precedence.
     // +optional
     RunAsNonRoot *bool
     // Indicates that the container must run as a non-root group.
     // If true, the Kubelet will validate the image at runtime to ensure that it
     // does not run as GID 0 (root) and fail to start the container if it does.
     // If unset or false, no such validation will be performed.
     // May also be set in SecurityContext.  If set in both SecurityContext and
     // PodSecurityContext, the value specified in SecurityContext takes precedence.
     // +optional
     RunAsNonRootGroup *bool

    .....
 }
```

#### PodSecurityContext 

```
type PodSecurityContext struct {
     //Other fields not shown for brevity
    ..... 

     // The UID to run the entrypoint of the container process.
     // Defaults to user specified in image metadata if unspecified.
     // May also be set in SecurityContext.  If set in both SecurityContext and
     // PodSecurityContext, the value specified in SecurityContext takes precedence
     // for that container.
     // +optional
     RunAsUser *int64
     // The GID to run the entrypoint of the container process.
     // Defaults to group specified in image metadata if unspecified.
     // May also be set in PodSecurityContext.  If set in both SecurityContext and
     // PodSecurityContext, the value specified in SecurityContext takes precedence.
     // +optional
     RunAsGroup *int64
     // Indicates that the container must run as a non-root user.
     // If true, the Kubelet will validate the image at runtime to ensure that it
     // does not run as UID 0 (root) and fail to start the container if it does.
     // If unset or false, no such validation will be performed.
     // May also be set in SecurityContext.  If set in both SecurityContext and
     // PodSecurityContext, the value specified in SecurityContext takes precedence.
     // +optional
     RunAsNonRoot *bool
     // Indicates that the container must run as a non-root group.
     // If true, the Kubelet will validate the image at runtime to ensure that it
     // does not run as GID 0 (root) and fail to start the container if it does.
     // If unset or false, no such validation will be performed.
     // May also be set in SecurityContext.  If set in both SecurityContext and
     // PodSecurityContext, the value specified in SecurityContext takes precedence.
     // +optional
     RunAsNonRootGroup *bool


    .....
 }
```

#### PodSecurityPolicy

PodSecurityPolicy defines strategies or conditions that a pod must run with in order to be accepted
into the system. Two of the relevant strategies are RunAsUser and SupplementalGroups. We introduce 
a new strategy called RunAsGroup which will support the following options:
- MustRunAs
- MustRunAsNonRoot
- RunAsAny

```
// PodSecurityPolicySpec defines the policy enforced.
 type PodSecurityPolicySpec struct {
     //Other fields not shown for brevity
    ..... 
  // RunAsUser is the strategy that will dictate the allowable RunAsUser values that may be set.
  RunAsUser RunAsUserStrategyOptions
  // SupplementalGroups is the strategy that will dictate what supplemental groups are used by the SecurityContext.
  SupplementalGroups SupplementalGroupsStrategyOptions


  // RunAsGroup is the strategy that will dictate the allowable RunAsGroup values that may be set.
  RunAsGroup RunAsGroupStrategyOptions
   .....
}

// RunAsGroupStrategyOptions defines the strategy type and any options used to create the strategy.
 type RunAsUserStrategyOptions struct {
     // Rule is the strategy that will dictate the allowable RunAsGroup values that may be set.
     Rule RunAsGroupStrategy
     // Ranges are the allowed ranges of gids that may be used.
     // +optional
     Ranges []GroupIDRange
 }

// RunAsGroupStrategy denotes strategy types for generating RunAsGroup values for a
 // SecurityContext.
 type RunAsGroupStrategy string
 
 const (
     // container must run as a particular gid.
     RunAsGroupStrategyMustRunAs RunAsGroupStrategy = "MustRunAs"
     // container must run as a non-root gid
     RunAsGroupStrategyMustRunAsNonRoot RunAsGroupStrategy = "MustRunAsNonRoot"
     // container may make requests for any gid.
     RunAsGroupStrategyRunAsAny RunAsGroupStrategy = "RunAsAny"
 )
```

## Behavior

Following points should be noted:

- `FSGroup` and `SupplementalGroups` will continue to have their old meanings and would be untouched.  
- The `RunAsGroup` In the SecurityContext will override the `RunAsGroup` in the PodSecurityContext.
- If both `RunAsUser` and `RunAsGroup` are NOT provided, the USER field in Dockerfile is used
- If both `RunAsUser` and `RunAsGroup` are specified, that is passed directly as User.
- If only one of `RunAsUser` or `RunAsGroup` is specified, the remaining value is decided by the Runtime,
  where the Runtime behavior is to make it run with uid or gid as 0.
- If a non numeric Group is specified in the Dockerfile and `RunAsNonRootGroup` is set, this will be 
  treated as error, similar to the behavior of `RunAsNonRoot` for non numeric User in Dockerfile.

Basically, we guarantee to set the values provided by user, and the runtime dictates the rest.

Here is an example of what gets passed to docker User
- runAsUser set to 9999, runAsGroup set to 9999 -> Config.User set to 9999:9999
- runAsUser set to 9999, runAsGroup unset -> Config.User set to 9999 -> docker runs you with 9999:0
- runAsUser unset, runAsGroup set to 9999 -> Config.User set to :9999 -> docker runs you with 0:9999 
- runAsUser unset, runAsGroup unset -> Config.User set to whatever is present in Dockerfile
This is to keep the behavior backward compatible and as expected.

## Summary of Changes needed

At a high level, the changes classify into:
1. API
2. Validation
3. CRI
4. Runtime for Docker and rkt
5. Swagger
6. DockerShim
7. Admission
8. Registry

- plugin/pkg/admission/security/podsecuritypolicy
- plugin/pkg/admission/securitycontext
- pkg/securitycontext/util.go
- pkg/security/podsecuritypolicy/selinux
- pkg/security/podsecuritypolicy/user
- pkg/security/podsecuritypolicy/group
- pkg/registry/extensions/podsecuritypolicy/storage
- pkg/kubelet/rkt
- pkg/kubelet/kuberuntime
- pkg/kubelet/dockershim/
- pkg/kubelet/apis/cri/v1alpha1/runtime
- pkg/apis/extensions/validation/
- pkg/api/validation/
- api/swagger-spec/
- api/openapi-spec/swagger.json

