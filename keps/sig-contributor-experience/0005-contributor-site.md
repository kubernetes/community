---
kep-number: 5
title: Contributor Site
authors:
  - "@jbeda"
owning-sig: sig-contributor-experience
participating-sigs:
  - sig-architecture
  - sig-docs
reviewers:
  - "@castrojo"
approvers:
  - "@parispittman"
editor: TBD
creation-date: "2018-02-19"
last-updated: "2018-03-07"
status: provisional
---

# Contributor Site

## Table of Contents

* [Table of Contents](#table-of-contents)
* [Summary](#summary)
* [Motivation](#motivation)
    * [Goals](#goals)
    * [Non-Goals](#non-goals)
* [Proposal](#proposal)
    * [Risks and Mitigations](#risks-and-mitigations)
* [Graduation Criteria](#graduation-criteria)
* [Implementation History](#implementation-history)
* [Drawbacks](#drawbacks)
* [Alternatives](#alternatives)

## Summary

We need a way to organize and publish information targeted at contributors.
In order to continue to scale the Kubernetes contributor community we need a convenient, scalable and findable way to publish information.

## Motivation

While the current kubernetes.io site is great for end users, it isn't often used by or aimed at project contributors.
Instead, most contributors look at documentation in markdown files that are spread throughout a wide set of repos and orgs.
It is difficult for users to find this documentation.

Furthermore, this documentation is often duplicated and out of date.
The fact that it isn't collected in one place and presented as a whole leads to fragmentation.
Often times documentation will be duplicated because the authors themselves can't find the relevant docs.

This site will also serve as a starting point for those that are looking to contribute.
This site (and the contributor guide) can provide a soft introduction to the main processes and groups.


Finally, some simple domain specific indexing could go a long way to make it easier to discover and cross link information.
Specifically, building a site that can take advantage of the KEP metadata will both make KEPs more discoverable and encourage those in the community to publish information in a way that *can* be discovered.

### Goals

* A contributor community facing portal to collect information for those actively working on upstream Kubernetes.
* An easy to remember URL. (`contrib.kubernetes.io`? `contributors.kubernetes.io`? `c.kubernetes.io`?)
* A streamlined process to update and share this information.
  Ownership should be delegated using the existing OWNERS mechanisms.
* A site that will be indexed well on Google to collect markdown files from the smattering of repos that we currently have.
  This includes information that is currently in the [community repo](https://github.com/kubernetes/community).
* Provide a place to launch and quickly evolve the contributor handbook.
* Build some simple tools to enhance discoverability within the site.
  This could include features such as automatically linking KEP and SIG names.
* Over time, add an index of events, meetups, and other forums for those that are actively contributing to k8s.

### Non-Goals

* Actively migrate information from multiple orgs/repos.
  This should be a place that people in the contributor community choose to use to communicate vs. being forced.
* Create a super dynamic back end.  This is most likely served best with a static site.
* Other extended community functions like a job board or a list of vendors.

## Proposal

We will build a new static site out of the [community repo](https://github.com/kubernetes/community).

This site will be focused on communicating with and being a place to publish information for those that are looking to contribute to Kubernetes.

We will use Hugo and netlify to build and host the site, respectively. (Details TBD)

The main parts of the site that will be built out first:
* A main landing page describing the purpose of the site.
* A guide on how to contribute/update the site.
* A list and index of KEPs
* A place to start publishing and building the contributor guide.

### Risks and Mitigations

The main risk here is abandonment and rot.
If the automation for updating the site breaks then someone will have to fix it.
If the people or the skillset doesn't exist to do so then the site will get out of sync with the source and create more confusion.

To mitigate this we will (a) ensure that SIG-contributor-experience is signed up to own this site moving forward and (b) keep it simple.
By relying on off the shelf tooling with many users (Hugo and Netlify) we can ensure that there are fewer custom processes and code to break.
The current generation scripts in the community repo haven't proven to be too much for us to handle.

## Graduation Criteria

This effort will have succeeded if:

* The contributor site becomes the de-facto way to publish information for the community.
* People consistently refer to the contributor site when answering questions about "how do I do X" or "what is the status of X".
* The amount of confusion over where to find information is reduced.
* Others in the contributor community actively look to expand the information on the contributor site and move information from islands to this site.

## Implementation History

## Drawbacks

The biggest drawback is that this is yet another thing to keep running.
Currently the markdown files are workable but not super discoverable.
However, they utilize the familiar mechanisms and do not require extra effort or understanding to publish.

The current mechanisms also scale across orgs and repos.
This is a strength as the information is close to the code but also a big disadvantage as it ends up being much less discoverable.

## Alternatives

One alternative is to do nothing.
However, the smattering of markdown through many repos is not scaling and is not discoverable via Google or for other members of the contributor community.

The main alternative here is to build something that is integrated into the user facing kubernetes.io site.
This is not preferred for a variety of reasons.

* **Workflow.** Currently there is quite a bit of process for getting things merged into the main site.
  That process involves approval from someone on SIG-Docs from an editorial point of view along with approval for technical accuracy.
  The two stage approval slows down contributions and creates a much larger barrier than the current markdown based flow.
  In addition, SIG-Docs is already stretched thin dealing with the (more important) user facing content that is their main charter.
* **Quality standards.** The bar for the user facing site is higher than that of the contributor site.
  Speed and openness of communication dominates for the contributor facing site.
  Our bar here is the current pile of Markdown.
* **Different tooling.** We may want to create specialized preprocessors as part of the contributor site build process.
  This could include integrating our current expansion of sigs.yaml into Markdown files.
  It may also include recognizing specific patterns (KEP-N) and creating automatic linkages.
  Applying these to a part of a site or validating them across a larger site will slow creation of these tools.

An alternative to building directly into the website repo is to build in some other repo and do some sort of import into the main website repo.
There are serious downsides to this approach.

* **No pre-commit visualization.** Netlifies capability to show a preview per PR won't work with a custom cross repo workflow.
* **Higher latency for changes.** If the merges are batched and manually approved then there could be a significant time gap between when something is changed and when it is published.
  This is a significant change from the current "pile of markdown in github" process.
* **Opportunity for more complex build breaks.** If something is checked into a satellite repo it may pass all of the presubmit tests there but then fail presubmits on the parent repo.
  This creates a situation where manual intervention is required.
  Complicated pre-submit tests could be built for the satellite repo but those need to be maintained and debugged themselves.
* **New tooling.** New tooling would need to be built that doesn't directly benefit the target audience.
  This tooling will have to be documented and supported vs. using an off the shelf service like netlify.

