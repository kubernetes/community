# Overview

Describes the process and tooling (`find_green_build`) used to find a
binary signal from the Kubernetes testing framework for the purposes of
selecting a release candidate.  Currently this process is used to gate
all Kubernetes releases.

## Motivation

Previously, the guidance in the [(now deprecated) release document](https://github.com/kubernetes/kubernetes/blob/fc3ef9320eb9d8211d85fbc404e4bbdd751f90af/docs/devel/releasing.md)
was to "look for green tests".  That is, of course, decidedly insufficient.

Software releases should have the goal of being primarily automated and
having a gating binary test signal is a key component to that ultimate goal.

## Design

### General

The idea is to capture and automate the existing manual methods of
finding a green signal for testing.

* Identify a green run from the primary job `ci-kubernetes-e2e-gce`
* Identify matching green runs from the secondary jobs

The tooling should also have a simple and common interface whether using it
for a dashboard, to gate a release within anago or for an individual to use it
to check the state of testing at any time.

Output looks like this:

```
$ find_green_build
find_green_build: BEGIN main on djmm Mon Dec 19 16:28:15 PST 2016

Checking for a valid github API token: OK
Checking required system packages: OK
Checking/setting cloud tools: OK

Getting ci-kubernetes-e2e-gce build results from Jenkins...
Getting ci-kubernetes-e2e-gce-serial build results from Jenkins...
Getting ci-kubernetes-e2e-gce-slow build results from Jenkins...
Getting ci-kubernetes-kubemark-5-gce build results from Jenkins...
Getting ci-kubernetes-e2e-gce-reboot build results from Jenkins...
Getting ci-kubernetes-e2e-gce-scalability build results from Jenkins...
Getting ci-kubernetes-test-go build results from Jenkins...
Getting ci-kubernetes-cross-build build results from Jenkins...
Getting ci-kubernetes-e2e-gke-serial build results from Jenkins...
Getting ci-kubernetes-e2e-gke build results from Jenkins...
Getting ci-kubernetes-e2e-gke-slow build results from Jenkins...

(*) Primary job (-) Secondary jobs

Jenkins Job                       Run #   Build # Time/Status
= ================================= ======  ======= ===========
* ci-kubernetes-e2e-gce             #1668   #2347   [14:46 12/19]
* (--buildversion=v1.6.0-alpha.0.2347+9925b68038eacc)
- ci-kubernetes-e2e-gce-serial      --      --      GIVE UP

* ci-kubernetes-e2e-gce             #1666   #2345   [13:23 12/19]
* (--buildversion=v1.6.0-alpha.0.2345+523ff93471b052)
- ci-kubernetes-e2e-gce-serial      --      --      GIVE UP

* ci-kubernetes-e2e-gce             #1664   #2341   [09:38 12/19]
* (--buildversion=v1.6.0-alpha.0.2341+def802272904c0)
- ci-kubernetes-e2e-gce-serial      --      --      GIVE UP

* ci-kubernetes-e2e-gce             #1662   #2339   [08:45 12/19]
* (--buildversion=v1.6.0-alpha.0.2339+ce67a03b81dee5)
- ci-kubernetes-e2e-gce-serial      --      --      GIVE UP

* ci-kubernetes-e2e-gce             #1653   #2335   [07:42 12/19]
* (--buildversion=v1.6.0-alpha.0.2335+d6046aab0e0678)
- ci-kubernetes-e2e-gce-serial      #192    #2335   PASSED
- ci-kubernetes-e2e-gce-slow        #989    #2335   PASSED
- ci-kubernetes-kubemark-5-gce      #2602   #2335   PASSED
- ci-kubernetes-e2e-gce-reboot      #1523   #2335   PASSED
- ci-kubernetes-e2e-gce-scalability #460    #2335   PASSED
- ci-kubernetes-test-go             #1266   #2335   PASSED
- ci-kubernetes-cross-build         --      --      GIVE UP

* ci-kubernetes-e2e-gce             #1651   #2330   [06:43 12/19]
* (--buildversion=v1.6.0-alpha.0.2330+75dfb21018a7c3)
- ci-kubernetes-e2e-gce-serial      #191    #2319   PASSED
- ci-kubernetes-e2e-gce-slow        #988    #2330   PASSED
- ci-kubernetes-kubemark-5-gce      #2599   #2330   PASSED
- ci-kubernetes-e2e-gce-reboot      #1521   #2330   PASSED
- ci-kubernetes-e2e-gce-scalability #459    #2321   PASSED
- ci-kubernetes-test-go             #1264   #2330   PASSED
- ci-kubernetes-cross-build         #320    #2330   PASSED
- ci-kubernetes-e2e-gke-serial      #233    #2319   PASSED
- ci-kubernetes-e2e-gke             #1834   #2330   PASSED
- ci-kubernetes-e2e-gke-slow        #1041   #2330   PASSED

JENKINS_BUILD_VERSION=v1.6.0-alpha.0.2330+75dfb21018a7c3
RELEASE_VERSION[alpha]=v1.6.0-alpha.1
RELEASE_VERSION_PRIME=v1.6.0-alpha.1
```

### v1

The initial release of this analyzer did everything on the client side.
This was slow to grab 100s of individual test results from GCS.
This was mitigated somewhat by building a local cache, but for those that
weren't using it regularly, the cache building step was a significant
(~1 minute) hit when just trying to check the test status.

### v2

Building and storing that local cache on the jenkins server at build time
was the way to speed things up.  Getting the cache from GCS is now consistent
for all users at ~10 seconds.  After that the analyzer is running.


## Uses

`find_green_build` and its functions are used in 3 ways:

1. During the release process itself via `anago`.
1. When creating a pending release notes report via `relnotes --preview`,
   used in creating dashboards
1. By an individual to get a quick check on the binary signal status of jobs

## Future work

1. There may be other ways to improve the performance of this check by
   doing more work server side.
1. Using the `relnotes --preview` output to generate an external dashboard
   will give more real-time visibility to both candidate release notes and
   testing state.
