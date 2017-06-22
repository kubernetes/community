## Abstract


As a Kubernetes User, i should be able to specify both user id and group id for the containers running 
inside a pod on a per Container basis, similar to how docker allows that using docker run options -u, 
--user="" Username or UID (format: <name|uid>[:<group|gid>]) format.

PodSecurityContext allows Kubernetes users to specify RUnAsUser which can be overriden by RunAsUser
in SecurtiyContext on a per Container basis. There is no equivalent field for specifying the primary
Group of the running container.

## Motivation

Enterprise Kubernetes users want to run containers as non root. This means running containers with a 
non zero user id and non zero group id. This gives Enterprises confidence that their customer code
is running with least privilege and if it escapes the container boundary, will still cause least harm
by decreasing the attack surface.

## Goals

1: Provide the ability to specify the Primary Group Id for a container inside a Pod
2: Bring launching of containers using Kubernetes at par with Dockers by supporting the same features.


## Use Cases

Use case 1:
As a cluster operator, i should be able to control both user id and primary group id of containers 
launched using Kubernetes.

Use case 2:
As a cluster operator, i should be able to override user id and primary group id of a container as 
specified in the Dockerfile without having to create a new docker image for the container.

## Design

### Model

Introduce a new API field in SecurityContext and PodSecurityContext called `RunAsGroup`

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
     // Defaults to user specified in image metadata if unspecified.
     // May also be set in PodSecurityContext.  If set in both SecurityContext and
     // PodSecurityContext, the value specified in SecurityContext takes precedence.
     // +optional
     RunAsGroup *int64

    .....
 }
 

type PoddSecurityContext struct {
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
     // Defaults to user specified in image metadata if unspecified.
     // May also be set in PodSecurityContext.  If set in both SecurityContext and
     // PodSecurityContext, the value specified in SecurityContext takes precedence.
     // +optional
     RunAsGroup *int64

    .....
 }

## Behavior

Following points should be noted:-

- `FSGroup` and `SupplementalGroups` will continue to have their old meanings and would be untouched.  
- The `RunAsGroup` in the SecurityContext will ovveride the `RunAsGroup` in the PodSecurityContext.
- If no RunAsGroup is provided in the PodSecurityContext and SecurityContext, the Group provided 
  in the image will be used.
- If no RunAsGroup is provided in the PodSecurityContext and SecurityContext, and none in the image,
  the container will run with primary Group as root(0).

## Not In Scope

PodSecurityPolicy defines strategies or conditions that a pod must run with in order to be accepted
into the system. Some of these stratgies are RunAsUser and SupplementalGroups. After this change we
would introduce a new strategy called RunAsGroup which will support the following options:-
- MustRunAs
- MustRunAsNonRoot
- RunAsAny
 Although these changes are simple, i would like a wider discussion on whether we want this as part 
of this feature or we need some kind of consolidation both RunASUser and RunAsGroup strategy into
one.


<!-- BEGIN MUNGE: GENERATED_ANALYTICS -->
[![Analytics](https://kubernetes-site.appspot.com/UA-36037335-10/GitHub/docs/proposals/security-context-constraints.md?pixel)]()
<!-- END MUNGE: GENERATED_ANALYTICS -->
