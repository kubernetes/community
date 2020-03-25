# Namespace Sameness - SIG Multicluster Position Statement

Author: Jeremy Olmsted-Thompson (**[@jeremyot](https://github.com/jeremyot)**), Google  
Last Edit: 2020/03/24  
Status: IN REVIEW  

## Goal
To establish a normative statement for multi-cluster namespace semantics and
governance as a building block for further development which will require
specifying behaviors across clusters.

## Context
Users are reaching for multi-cluster deployments for a
[variety of reasons](http://bit.ly/k8s-multicluster-conversation-starter-doc).
However, Kubernetes treats the cluster boundary as the edge of the universe.
There are currently no standard practices for how to extend the Kubernetes
resource model across multiple clusters. Without common patterns we can’t build
portable tooling to facilitate multi-cluster capabilities and know that behavior
will be consistent for each user.

## Scope
A single organization may need multiple, disjoint sets of clusters. They may,
for example, represent different phases of the development lifecycle (dev,
staging, prod) or support unrelated projects. Each organization governs their
own clusters in isolation, so the scope of a namespace can only reasonably be
declared within the organizational boundary. The scope of namespace identity is
defined as the union of clusters, governed by a single authority, that are
expected to work together.  An authority is a company, organization, team,
individual, or other entity which is entrusted to manage these clusters and, in
particular, to create namespaces in them.

## Position
**For a set of related clusters governed by a single authority, all namespaces of
a given name are considered to be the same namespace. A single namespace should
have a consistent owner across the set of clusters.**
