# Kubernetes: New Client Library Procedure

**Status:** Approved by SIG API Machinery on March 29th, 2017

**Authors:** @mbohlool, @lavalamp

**Last Updated:** 2017-03-06

# Background

Kubernetes currently officially supports both Go and [Python client](https://github.com/kubernetes-incubator/client-python) libraries. The go client is developed and extracted from main kubernetes repositories in a complex process. On the other hand, the python client is based on OpenAPI, and is mostly generated code (via [swagger-codegen](https://github.com/swagger-api/swagger-codegen)). By generating the API Operations and Data Models, updating the client and tracking changes from main repositories becomes much more sustainable.

The python client development process can be repeated for other languages. Supporting a basic set of languages would help the community to build more tools and applications based on kubernetes. We may consider adjusting the go client library generation to match, but that is not the goal of this doc.

More background information can be found [here](https://github.com/kubernetes/kubernetes/issues/22405).

# Languages

The proposal is to support *Java*, *PHP*, *Ruby*, *C#*, and *Javascript* in addition to the already supported libraries, Go and Python. There are good clients for each of these languages, but having a basic supported client would even help those client libraries to focus on their interface and delegate transport and config layer to this basic client. For community members willing to do some work producing a client for their favorite language, this doc establishes a procedure for going about this.

# Development process

Development would be based on a generated client using OpenAPI and [swagger-codegen](https://github.com/swagger-api/swagger-codegen). Some basic functionality such as loading config, watch, etc. would be added (i.e., hand-written) on top of this generated client. The idea is to develop transportation and configuration layer, and modify as few generated files (such as API and models) as possible. The clients would be in alpha, beta or stable stages, and may have either bronze, silver, or gold support according to these requirements:

### Client Capabilities

* Bronze Requirements [![Client Capabilities](https://img.shields.io/badge/Kubernetes%20client-Bronze-blue.svg?style=plastic&colorB=cd7f32&colorA=306CE8)](/contributors/design-proposals/api-machinery/csi-new-client-library-procedure.md#client-capabilities)

    * Support loading config from kube config file

        * Basic Auth (username/password) (Add documentation to discourage this and only use for testing.)

        * X509 Client certificate (inline and referenced by file)

        * Bearer tokens (inline and referenced by file)

        * encryption/TLS (inline, referenced by file, insecure)

    * Basic API calls such as list pods should work

    * Works from within the cluster environment.

* Silver Requirements [![Client Capabilities](https://img.shields.io/badge/Kubernetes%20client-Silver-blue.svg?style=plastic&colorB=C0C0C0&colorA=306CE8)](/contributors/design-proposals/api-machinery/csi-new-client-library-procedure.md#client-capabilities)

    * Support watch calls

* Gold Requirements [![Client Capabilities](https://img.shields.io/badge/Kubernetes%20client-Gold-blue.svg?style=plastic&colorB=FFD700&colorA=306CE8)](/contributors/design-proposals/api-machinery/csi-new-client-library-procedure.md#client-capabilities)

    * Support exec, attach, port-forward calls (these are not normally supported out of the box from [swagger-codegen](https://github.com/swagger-api/swagger-codegen))

    * Proto encoding

### Client Support Level

* Alpha [![Client Support Level](https://img.shields.io/badge/kubernetes%20client-alpha-green.svg?style=plastic&colorA=306CE8)](/contributors/design-proposals/api-machinery/csi-new-client-library-procedure.md#client-support-level)

    * Clients don’t even have to meet bronze requirements

* Beta [![Client Support Level](https://img.shields.io/badge/kubernetes%20client-beta-green.svg?style=plastic&colorA=306CE8)](/contributors/design-proposals/api-machinery/csi-new-client-library-procedure.md#client-support-level)

    * Client at least meets bronze standards

    * Reasonably stable releases

    * Installation instructions

    * 2+ individual maintainers/owners of the repository

* Stable [![Client Support Level](https://img.shields.io/badge/kubernetes%20client-stable-green.svg?style=plastic&colorA=306CE8)](/contributors/design-proposals/api-machinery/csi-new-client-library-procedure.md#client-support-level)

    * Support level documented per-platform

    * Library documentation

    * Deprecation policy (backwards compatibility guarantees documented)

        * How fast may the interface change?

    * Versioning procedure well documented

    * Release process well documented

    * N documented users of the library

The API machinery SIG will somewhere (community repo?) host a page listing clients, including their stability and capability level from the above lists.

# Kubernetes client repo

New clients will start as repositories in the [kubernetes client](https://github.com/kubernetes-client/) organization.

We propose to make a `gen` repository to house common functionality such as preprocessing the OpenAPI spec and running the generator, etc.

For each client language, we’ll make a client-[lang]-base and client-[lang] repository (where lang is one of java, csharp, js, php, ruby). The base repo would have all utility and add-ons for the specified language and the main repo will have generated client and reference to base repo.

# Support

These clients will be supported by the Kubernetes [API Machinery special interest group](/sig-api-machinery); however, individual owner(s) will be needed for each client language for them to be considered stable; the SIG won’t be able to handle the support load otherwise. If the generated clients prove as easy to maintain as we hope, then a few individuals may be able to own multiple clients.

