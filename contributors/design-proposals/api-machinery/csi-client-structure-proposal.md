# Overall Kubernetes Client Structure

**Status:** Approved by SIG API Machinery on March 29th, 2017

**Authors:** @lavalamp, @mbohlool

**last edit:** 2017-3-22

## Goals

* Users can build production-grade programmatic use of Kubernetes-style APIs in their language of choice.

## New Concept

Today, Kubernetes has the concept of an API Group. Sometimes it makes sense to package multiple groups together in a client, for example, the core APIs we publish today. I’ll call this a "group collection" as it sounds a bit better than “group group.” Group collections will have names. In particular, this document uses “core” as the name for the current set of APIs.

## Repositories

We’ve decomposed the problem into several components. We’d like to make the following repositories under a new `kubernetes-client` github org.

* github.com/kubernetes-client/gen

    * Contents:

        * OpenAPI preprocessing (shared among multiple languages) and client generator(s) scripts

        * The Kubernetes go language client generator, which currently takes as input a tree of types.go files.

        * Future work: Convert OpenAPI into a types.go tree, or modify the input half of the go language generator.

        * Make the client generation completely automated so it can be part of a build process. Less reason for people to create custom repos as many clients for single language (and different api extensions) could be confusing.

    * gen is intended to be used as part of build tool chains.

* github.com/kubernetes-client/go-base

    * Contents:

        * All reusable components of the existing client-go library, including at least:

            * Transport

            * RESTClient

            * Workqueue

            * Informer (not the typed informers)

            * Dynamic and Discovery clients

            * Utility functions (if still necessary)

        * But omitting:

            * API types

            * current generated code.

    * go-base is usable as a client on its own (dynamic client, discovery client)

* github.com/kubernetes-client/core-go

    * Contents:

        * The existing examples.

        * The output (including API types) resulting from running client-gen on the source api types of the core API Group collection.

        * Any hand-written *_expansion.go files.

* github.com/kubernetes-client/python-base

    * Hand-tuned pieces (auth, watch support etc) for python language clients.

* github.com/kubernetes-client/core-python

    * The output of kubernetes-client/gen for the python language.

    * Note: We should provide a packaging script to make sure this is backward compatible with current `pip` package.

* github.com/kubernetes-client/core-{lang} and github.com/kubernetes-client/{lang}-base

    * The output of kubernetes-client/gen for language {lang}, and hand-tuned pieces for that language. [See here](https://docs.google.com/document/d/1hsJNlowIg-u_rz3JBw9hXh6rj2LgeAdMY7OleOL1srw/edit).

Note that the word "core" in the above package names represents the API groups that will be included in the repository. “core” would indicate inclusion of the groups published in the client today. (One can imagine replacing it with “service-catalog” etc.)

## Why this split?

The division of each language into two repositories is intended to allow for composition: generated clients for [multiple different API sources](https://docs.google.com/document/d/1UZyb5sQc-G2Ix4YL6dA9f4xJWtz3VKCiFPNHVHDUq9Y/edit#) can be used together without any code duplication. To clarify further:

* Some who run the generator don't need go-base.

    * I want to publish an API extension client.

* Some who use the kubernetes-client/core-go package don't need the generator.

    * I don’t use any extensions, just vanilla k8s.

* Some who use go-base need neither the generator nor the core-go package.

    * I want to use the dynamic client since I only care about metadata.

* Those who write automation only for their extension need the generator and go-base but possibly not core-go.

That is, there should only be one kubernetes-client/{lang}-base for any given language, but many different API providers may provide a kubernetes-client/{collection}-{lang} for the language (e.g., core Kubernetes, the API registration API, the cluster federation effort, service catalog, heapster/metrics API, OpenShift).

It is preferred to use the generation as part of the user’s build process so we can have fewer kubernetes-client/{collection}-{lang} for custom APIs out in the wild. Users should only expect super popular extensions to host their own client, as there’s otherwise a combinatorial explosion of API Group collections x languages.

Users may want to run the client generator themselves with the particular collection of APIs enabled in their particular cluster, or at a different version of the generator.

## Versioning

`kubernetes-client/gen` must be versioned, so that users can get deterministic, repeatable client interfaces. (Use case: adding a reference to a new API resource to existing slightly out-of-date code.) We will use semver.

The versions of kubernetes-client/gen must correspond to the versions of all kubernetes-client/{lang}-base repositories.

kubernetes-client/{collection}-{lang} repos have their own version. Each release of such repos must clearly state both:

* the version of the OpenAPI source (or go types for the go client)

* the version of kubernetes-client/gen used to generate the client

This will allow users to regenerate any given kubernetes-client/{collection}-{lang} repo, adding custom API resources etc.

## Clients for API Extensions

Providers of API extensions (e.g., service catalog or cluster federation) may choose to publish a generated client for their API types, to make users’ lives easier (it’s not strictly necessary, since end users could run the generator themselves). Publishing a client like this should be as easy as importing the generator execution script from e.g. the main kubernetes-client/core-go repo, and providing a different input source.

## Language-specific considerations

### Go

* The typed informer generator should be included with the client generator.

* We need a "clientset" adaptor concept to make it easy to compose clients from disparate client repos.

* Prior strategy doc [is here](https://docs.google.com/document/d/1h_IBGYPMa8FS0oih4NbVkAMAzM7YTHr76VBcKy1qFbg/edit#heading=h.ve95t2prztno).

### {My Favorite Language}

Read about the process for producing an official client library [here](https://docs.google.com/document/d/1hsJNlowIg-u_rz3JBw9hXh6rj2LgeAdMY7OleOL1srw/edit).

## Remaining Design Work

* Client Release Process

    * Needs definition, i.e., who does what to which repo when.

* Client release note collection mechanism

    * For go, we’ve talked about amending the Kubernetes merge bot to require a client-relnote: `Blah` in PRs that touch a few pre-identified directories.

    * Once we are working on the generator in the generator repo, it becomes easier to assemble release notes: we can grab all changes to the interface by looking at the kubernetes-client/gen and kubernetes-client/{lang}-base repos, and we can switch the release note rule to start tracking client-visible API changes in the main repository.

* Client library documentation

    * Ideally, generated

    * Ideally, both:

        * In the native form for the language, and

        * In a way that can easily be aggregated with the official Kubernetes API docs.

## Timeline Guesstimate

Rough order of changes that need to be made.

1. Begin working towards collecting client release notes.

2. Split client-go into kubernetes-client/core-go and kubernetes-client/go-base

3. Move go client generator into kubernetes-client/gen

    1. kubernetes-client/gen becomes the canonical location for this. It is vendored into the main repository (downloaded at a specific version & invoked directly would be even better).

    2. The client generator is modified to make a copy of the go types specifically for the client. (Either from the source go types, or generated from an OpenAPI spec.)

4. Split client-python into kubernetes-client/core-python and kubernetes-client/python-base

5. Move OpenAPI generator into kubernetes-client/gen

6. Declare 1.0.0 on kubernetes-client/gen, and all kubernetes-client/{lang}-base repositories. (This doesn’t mean they’re stable, but we need a functioning versioning system since the deliverable here is the entire process, not any one particular client.)

7. Instead of publishing kubernetes-client/core-go from its location in the staging directory:

    3. Add a script to it that downloads kubernetes-client/gen and the main repo and generates the client.

    4. Switch the import direction, so that we really just vendor kubernetes-client/core-go in the main repo. (Alternative: main repo can just run the generator itself to avoid having to make multiple PRs)

8. At this point we should have finished the bootstrapping process and we’ll be ready to automate and execute on whatever release process we’ve defined.

