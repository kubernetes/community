# Resource-dependent Pod Preset

<!-- vscode-markdown-toc -->
* 1. [Abstract](#Abstract)
* 2. [Motivation](#Motivation)
* 3. [Use Cases](#UseCases)
* 4. [Proposed Changes](#ProposedChanges)
	* 4.1. [API objects](#APIobjects)
* 5. [Open Issues](#OpenIssues)
* 6. [Alternatives](#Alternatives)
* 7. [Examples](#Examples)
	* 7.1. [Tensorflow](#Tensorflow)

<!-- vscode-markdown-toc-config
	numbering=true
	autoSave=true
	/vscode-markdown-toc-config -->
<!-- /vscode-markdown-toc -->

##  1. <a name='Abstract'></a>Abstract

This document proposes a policy object that lets Pods which use specific runtime
resources to abstract away some of their details and defer them until creation
time. For example, a Pod using GPU resources might need to access low-level
libraries that are inherently host-specific (sometimes kernel-specific, even).
Currently, such host volume mounts are not very portable, as different hosts
might run different Linux distributions. The problem is thus better tackled at
the cluster and namespace level.

##  2. <a name='Motivation'></a>Motivation

This proposal closely tracks and aligns with the [Pod Preset
proposal](https://github.com/therc/community/blob/master/contributors/design-proposals/pod-preset.md).
The motivations are virtually identical. The only nuance is the matching being
performed based on resources, not labels. Refer to that document for a fuller
picture.

##  3. <a name='UseCases'></a>Use Cases

Just looking at Nvidia hardware alone:

 - As an user, I want to run a plain Tensorflow image with a simple YAML file,
 without having to hardcode or worry about the vagaries of where on the host
 Nvidia's low-level libraries reside. The same pod definition should work both
 on my workstation and on a hosted Kubernetes cluster.
 - As an engineer, I, too, want to spend my time writing CUDA code, not dealing
 with low-level details. I would like to run the standard `nvidia-smi`
 monitoring tool in my container without shipping it myself (for licensing and
 compatibility issues) or hardcoding its location.
 - As a developer, I want to publish my application as a Helm chart, knowing
 that it'll work after the user or the admin have put the right policy in place.
 Helm could prompt for the policy to be created or even validate it as working.
 - As a cluster admin, I want the flexibility to migrate clusters to new base
 images where low-level libraries and binaries might live in different
 directories, without forcing my users to think about the issue or, worse,
 create different templates for different clusters.
 - As a cluster admin, I want to increase transparently the `CUDA_CACHE_MAXSIZE`
 variable from the default for all new pods, because our existing Docker images
 ship with pre-built code for older cards, but for newer hardware the JIT
 compiler gets called at runtime and its output is never cached.
 - As a cluster admin, I want to expose a host directory with low-level GPU
 debugging tools only to pods in the `dev` namespace.

For other resources:
 - As an admin, I want all pods using the opaque integer resource `localssd` to
 automatically mount a host directory of my choice, presumably where my SSD is
 mounted.
 - As an user, I just want to use the fast, local SSD disk without having to
   know where it lives.

##  4. <a name='ProposedChanges'></a>Proposed Changes

###  4.1. <a name='APIobjects'></a>API objects

Two basic options are available: augment `PodPresetSpec` or create a parallel
`ResourceAwarePodPreset`/`ResourceAwarePodPresetSpec` combination.

`ProdPresetSpec` currently lives in the `settings` group and looks like this:

```go
type PodPresetSpec struct {
    // Selector is a label query over a set of resources, in this case pods.
    // Required.
    Selector      unversioned.LabelSelector
    // Env defines the collection of EnvVar to inject into containers.
    // +optional
    Env           []EnvVar
    // EnvFrom defines the collection of EnvFromSource to inject into
    // containers.
    // +optional
    EnvFrom       []EnvFromSource
    // Volumes defines the collection of Volume to inject into the pod.
    // +optional
    Volumes       []Volume `json:omitempty`
    // VolumeMounts defines the collection of VolumeMount to inject into
    // containers.
    // +optional
    VolumeMounts  []VolumeMount
}
```

Ideally, the straighforward change would be a new kind of selector, perhaps
named `ResourceSelector`. This is simpler both in terms of implementation and
users' mental model: one object and one concept, rather than two, similar
ones. The LabelSelector is presently required, but it might have to become
optional, as long as a ResourceSelector is present. TBD: does this cause
migration issues? If both selectors are present, they would be `AND`ed together.
If an admin wanted to `OR` them, they would create two separate policies.

Alternatively, an entirely new object (and controller) could be created. This
would require maintaining parallel API objects, as well as associated logic.

A `ResourceSelector` could be a simple list of resource names:

```go
ResourceSelector   []ResourceName
```

The policy would apply if all the resources in the list are requested by the
pod.

##  5. <a name='OpenIssues'></a>Open Issues

- It's not clear if there's any value in more complex selectors, such as
  `alpha.kubernetes.io/nvidia-gpu > 2` or `not alpha.kubernetes.io/nvidia-gpu`.

- There is no automatic way to detect the lack of an appropriate policy, except
  if handled by a tool like Helm as described above.

- There is no clear semantics when multiple policies apply to the same pod. This
  is especially true when two policies want to e.g. set the same `PATH` or
  `LD_LIBRARY_PATH` environment variables. As with the pod preset proposal,
  during the alpha period, policies could be applied from oldest to newest. For
  the cases just mentioned, the user would probably expect the variable to be set to
  `/path1:/path2` (assuming an Unix-like container), instead of `path2` clobbering
  `path1`.

##  6. <a name='Alternatives'></a>Alternatives

An obvious alternative would use a plain `PodPreset` and forcing users to add a
new label (`use-gpu`?) on top of the existing GPU resource requests/limits. Such
an approach works, but is obviously error-prone. Omitting the label would result
in crash loops, which might or might not be easy to investigate. Such a solution
would lead to a proliferation across the Kubernetes ecosystem of similar, but
not identical labels, too. Perhaps projects like Helm could define a standard
set of such labels.

##  7. <a name='Examples'></a>Examples

###  7.1. <a name='Tensorflow'></a>Tensorflow

**Pod Preset:**

```yaml
kind: PodPreset
apiVersion: settings/v1alpha1
metadata:
  name: gpu-setup
  namespace: myns
spec:
  resourceSelector:
    - alpha.kubernetes.io/nvidia-gpu
  volumeMounts:
    - name: binaries
      mountPath: /opt/bin
      readOnly: true
    - name: libraries
      mountPath: /usr/local/lib/nvidia
      readOnly: true
  volumes:
    - name: binaries
      hostPath:
        path: /opt/bin
    - name: libraries
      hostPath:
        path: /opt/lib64
  env:
    - name: LD_LIBRARY_PATH
      value: /usr/local/lib/nvidia
```


