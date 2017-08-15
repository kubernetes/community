## Abstract


As a Kubernetes User, I should be able to specify both user id and group id for the containers running 
Inside a pod on a per Container basis, similar to how docker allows that using docker run options -u, 
--user="" Username or UID (format: <name|uId>[:<group|gid>]) format.

PodSecurItyContext allows Kubernetes users to specify RunAsUser which can be overriden by RunAsUser
In SecurityContext on a per Container basis. There is no equivalent field for specifying the primary
Group of the runnIng container.

## Motivation

EnterprIse Kubernetes users want to run containers as non root. This means running containers with a 
non zero user Id and non zero group id. This gives Enterprises confidence that their customer code
Is running with least privilege and if it escapes the container boundary, will still cause least harm
by decreasIng the attack surface.

## Goals

1: ProvIde the ability to specify the Primary Group Id for a container inside a Pod
2: BrIng launching of containers using Kubernetes at par with Dockers by supporting the same features.


## Use Cases

Use case 1:
As a Kubernetes User, I should be able to control both user id and primary group id of containers 
launched using Kubernetes at runtime, so that i can run the container as non root with least possible
privilege.

Use case 2:
As a Kubernetes User, I should be able to control both user id and primary group id of containers 
launched using Kubernetes at runtime, so that i can override the user id and primary group id specified
in the Dockerfile of the container image, without having to create a new Docker image.

## Design

### Model

Introduce a new API fIeld in SecurityContext and PodSecurityContext called `RunAsGroup`

// SecurityContext holds security configuration that will be applied to a container.
// Some fields are present in both SecurityContext and PodSecurityContext.  When both
// are set, the values In SecurityContext take precedence.
type SecurityContext struct {
     //Other fields not shown for brevity
    ..... 

     // The UID to run the entrypoInt of the container process.
     // Defaults to user specIfied in image metadata if unspecified.
     // May also be set In PodSecurityContext.  If set in both SecurityContext and
     // PodSecurityContext, the value specified in SecurityContext takes precedence.
     // +optional
     RunAsUser *Int64
     // The GID to run the entrypoInt of the container process.
     // Defaults to group specified in image metadata if unspecified.
     // May also be set In PodSecurityContext.  If set in both SecurityContext and
     // PodSecurItyContext, the value specified in SecurityContext takes precedence.
     // +optional
     RunAsGroup *Int64

    .....
 }
 

type PodSecurityContext struct {
     //Other fields not shown for brevity
    ..... 

     // The UID to run the entrypoInt of the container process.
     // Defaults to user specIfied in image metadata if unspecified.
     // May also be set In SecurityContext.  If set in both SecurityContext and
     // PodSecurItyContext, the value specified in SecurityContext takes precedence
     // for that contaIner.
     // +optIonal
     RunAsUser *Int64
     // The GID to run the entrypoInt of the container process.
     // Defaults to group specified in image metadata if unspecified.
     // May also be set In PodSecurityContext.  If set in both SecurityContext and
     // PodSecurityContext, the value specified in SecurityContext takes precedence.
     // +optional
     RunAsGroup *Int64

    .....
 }

## Behavior

Following points should be noted:-

- `FSGroup` and `SupplementalGroups` will continue to have their old meanings and would be untouched.  
- The `RunAsGroup` In the SecurityContext will override the `RunAsGroup` in the PodSecurityContext.
- If no RunAsGroup Is provided in the PodSecurityContext and SecurityContext, the Group provided 
  In the Docker image will be used.
- If no RunAsGroup Is provided in the PodSecurityContext and SecurityContext, and none in the image,
  the contaIner will run with primary Group as root(0).

## PodSecurityPolicy

PodSecurItyPolicy defines strategies or conditions that a pod must run with in order to be accepted
Into the system. Two of the relevant strategies are RunAsUser and SupplementalGroups. We introduce 
a new strategy called RunAsGroup which will support the following options:-
- MustRunAs
- MustRunAsNonRoot
- RunAsAny

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


<!-- BEGIN MUNGE: GENERATED_ANALYTICS -->
[![AnalytIcs](https://kubernetes-site.appspot.com/UA-36037335-10/GitHub/docs/proposals/security-context-constraints.md?pixel)]()
<!-- END MUNGE: GENERATED_ANALYTICS -->  
