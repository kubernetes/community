# Declarative application management in Kubernetes

> This article was authored by Brian Grant (bgrant0607) on 8/2/2017. The original Google Doc can be found here: [https://goo.gl/T66ZcD](https://goo.gl/T66ZcD)

Most users will deploy a combination of applications they build themselves, also known as **_bespoke_** applications, and **common off-the-shelf (COTS)** components. Bespoke applications are typically stateless application servers, whereas COTS components are typically infrastructure (and frequently stateful) systems, such as databases, key-value stores, caches, and messaging systems.

In the case of the latter, users sometimes have the choice of using hosted SaaS products that are entirely managed by the service provider and are therefore opaque, also known as **_blackbox_** *services*. However, they often run open-source components themselves, and must configure, deploy, scale, secure, monitor, update, and otherwise manage the lifecycles of these **_whitebox_** *COTS applications*.

This document proposes a unified method of managing both bespoke and off-the-shelf applications declaratively using the same tools and application operator workflow, while leveraging developer-friendly CLIs and UIs, streamlining common tasks, and avoiding common pitfalls. The approach is based on observations of several dozen configuration projects and hundreds of configured applications within Google and in the Kubernetes ecosystem, as well as quantitative analysis of Borg configurations and work on the Kubernetes [system architecture](https://github.com/kubernetes/community/blob/master/contributors/design-proposals/architecture.md), [API](/contributors/devel/sig-architecture/api-conventions.md), and command-line tool ([kubectl](https://github.com/kubernetes/community/wiki/Roadmap:-kubectl)).

The central idea is that a toolbox of composable configuration tools should manipulate configuration data in the form of declarative API resource specifications, which serve as a [declarative data model](https://docs.google.com/document/d/1RmHXdLhNbyOWPW_AtnnowaRfGejw-qlKQIuLKQWlwzs/edit#), not express configuration as code or some other representation that is restrictive, non-standard, and/or difficult to manipulate.

## Declarative configuration

Why the heavy emphasis on configuration in Kubernetes? Kubernetes supports declarative control by specifying users’ desired intent. The intent is carried out by asynchronous control loops, which interact through the Kubernetes API. This declarative approach is critical to the system’s self-healing, autonomic capabilities, and application updates. This approach is in contrast to manual imperative operations or flowchart-like orchestration.

This is aligned with the industry trend towards [immutable infrastructure](http://thenewstack.io/a-brief-look-at-immutable-infrastructure-and-why-it-is-such-a-quest/), which facilitates predictability, reversibility, repeatability, scalability, and availability. Repeatability is even more critical for containers than for VMs, because containers typically have lifetimes that are measured in days, hours, even minutes. Production container images are typically built from configurable/scripted processes and have parameters overridden by configuration rather than modifying them interactively.

What form should this configuration take in Kubernetes? The requirements are as follows:

* Perhaps somewhat obviously, it should support **bulk** management operations: creation, deletion, and updates.

* As stated above, it should be **universal**, usable for both bespoke and off-the-shelf applications, for most major workload categories, including stateless and stateful, and for both development and production environments. It also needs to be applicable to use cases outside application definition, such as policy configuration and component configuration.

* It should **expose** the full power of Kubernetes (all CRUD APIs, API fields, API versions, and extensions), be **consistent** with concepts and properties presented by other tools, and should **teach** Kubernetes concepts and API, while providing a **bridge** for application developers that prefer imperative control or that need wizards and other tools to provide an onramp for beginners.

* It should feel **native** to Kubernetes. There is a place for tools that work across multiple platforms but which are native to another platform and for tools that are designed to work across multiple platforms but are native to none, but such non-native solutions would increase complexity for Kubernetes users by not taking full advantage of Kubernetes-specific mechanisms and conventions.

* It should **integrate** with key user tools and workflows, such as continuous deployment pipelines and application-level configuration formats, and **compose** with built-in and third-party API-based automation, such as [admission control](https://kubernetes.io/docs/admin/admission-controllers/), autoscaling, and [Operators](https://coreos.com/operators). In order to do this, it needs to support **separation of concerns** by supporting multiple distinct configuration sources and preserving declarative intent while allowing automatically set attributes.

* In particular, it should be straightforward (but not required) to manage declarative intent under **version control**, which is [standard industry best practice](http://martinfowler.com/bliki/InfrastructureAsCode.html) and what Google does internally. Version control facilitates reproducibility, reversibility, and an audit trail. Unlike generated build artifacts, configuration is primary human-authored, or at least it is desirable to be human-readable, and it is typically changed with a human in the loop, as opposed to fully automated processes, such as autoscaling. Version control enables the use of familiar tools and processes for change control, review, and conflict resolution.

* Users need the ability to **customize** off-the-shelf configurations and to instantiate multiple **variants**, without crossing the [line into the ecosystem](https://docs.google.com/presentation/d/1oPZ4rznkBe86O4rPwD2CWgqgMuaSXguIBHIE7Y0TKVc/edit#slide=id.g21b1f16809_5_86) of [configuration domain-specific languages, platform as a service, functions as a service](https://kubernetes.io/docs/concepts/overview/what-is-kubernetes/#what-kubernetes-is-not), and so on, though users should be able to [layer such tools/systems on top](https://kubernetes.io/blog/2017/02/caas-the-foundation-for-next-gen-paas/) of the mechanism, should they choose to do so.

* We need to develop clear **conventions**, **examples**, and mechanisms that foster **structure**, to help users understand how to combine Kubernetes’s flexible mechanisms in an effective manner.

## Configuration customization and variant generation

The requirement that drives the most complexity in typical configuration solutions is the need to be able to customize configurations of off-the-shelf components and/or to instantiate multiple variants.

Deploying an application generally requires customization of multiple categories of configuration:

* Frequently customized

    * Context: namespaces, [names, labels](https://github.com/kubernetes/kubernetes/issues/1698), inter-component references, identity

    * Image: repository/registry (source), tag (image stream/channel), digest (specific image)

    * Application configuration, overriding default values in images: command/args, env, app config files, static data

    * Resource parameters: replicas, cpu, memory, volume sources

    * Consumed services: coordinates, credentials, and client configuration

* Less frequently customized

    * Management parameters: probe intervals, rollout constraints, utilization targets

* Customized per environment

    * Environmental adapters: lifecycle hooks, helper sidecars for configuration, monitoring, logging, network/auth proxies, etc

    * Infrastructure mapping: scheduling constraints, tolerations

    * Security and other operational policies: RBAC, pod security policy, network policy, image provenance requirements

* Rarely customized

    * Application topology, which makes up the basic structure of the application: new/replaced components

In order to make an application configuration reusable, users need to be able to customize each of those categories of configuration. There are multiple approaches that could be used:

* Fork: simple to understand; supports arbitrary changes and updates via rebasing, but hard to automate in a repeatable fashion to maintain multiple variants

* Overlay / patch: supports composition and useful for standard transformations, such as setting organizational defaults or injecting environment-specific configuration, but can be fragile with respect to changes in the base configuration

* Composition: useful for orthogonal concerns

    * Pull: Kubernetes provides APIs for distribution of application secrets (Secret) and configuration data (ConfigMap), and there is a [proposal open](http://issues.k8s.io/831) to support application data as well

        * the resource identity is fixed, by the object reference, but the contents are decoupled

            * the explicit reference makes it harder to consume a continuously updated stream of such resources, and harder to generate multiple variants

        * can give the PodSpec author some degree of control over the consumption of the data, such as environment variable names and volume paths (though service accounts are at conventional locations rather than configured ones)

    * Push: facilitates separation of concerns and late binding

        * can be explicit, such as with kubectl set or HorizontalPodAutoscaler

        * can be implicit, such as with LimitRange, PodSecurityPolicy, PodPreset, initializers

            * good for attaching policies to selected resources within a scope (namespace and/or label selector)

* Transformation: useful for common cases (e.g., names and labels)

* Generation: useful for static decisions, like "if this is a Java app…", which can be integrated into the declarative specification

* Automation: useful for dynamic adaptation, such as horizontal and vertical auto-scaling, improves ease of use and aids encapsulation (by not exposing those details), and can mitigate phase-ordering problems

* Parameterization: natural for small numbers of choices the user needs to make, but there are many pitfalls, discussed below

Rather than relying upon a single approach, we should combine these techniques such that disadvantages are mitigated.

Tools used to customize configuration [within Google](http://queue.acm.org/detail.cfm?id=2898444) have included:

* Many bespoke domain-specific configuration languages ([DSLs](http://flabbergast.org))

* Python-based configuration DSLs (e.g., [Skylark](https://github.com/google/skylark))

* Transliterate configuration DSLs into structured data models/APIs, layered over and under existing DSLs in order to provide a form that is more amenable to automatic manipulation

* Configuration overlay systems, override mechanisms, and template inheritance

* Configuration generators, manipulation CLIs, IDEs, and wizards

* Runtime config databases and spreadsheets

* Several workflow/push/reconciliation engines

* Autoscaling and resource-planning tools

Note that forking/branching generally isn’t viable in Google’s monorepo.

Despite many projects over the years, some of which have been very widely used, the problem is still considered to be not solved satisfactorily. Our experiences with these tools have informed this proposal, however, as well as the design of Kubernetes itself.

A non-exhaustive list of tools built by the Kubernetes community (see [spreadsheet](https://docs.google.com/spreadsheets/d/1FCgqz1Ci7_VCz_wdh8vBitZ3giBtac_H8SBw4uxnrsE/edit#gid=0) for up-to-date list), in no particular order, follows:

* [Helm](https://github.com/kubernetes/helm)
* [OC new-app](https://docs.openshift.com/online/dev_guide/application_lifecycle/new_app.html)
* [Kompose](https://github.com/kubernetes-incubator/kompose)
* [Spread](https://github.com/redspread/spread)
* [Draft](https://github.com/Azure/draft)
* [Ksonnet](https://github.com/ksonnet/ksonnet-lib)/[Kubecfg](https://github.com/ksonnet/kubecfg)
* [Databricks Jsonnet](https://databricks.com/blog/2017/06/26/declarative-infrastructure-jsonnet-templating-language.html)
* [Kapitan](https://github.com/deepmind/kapitan)
* [Konfd](https://github.com/kelseyhightower/konfd)
* [Templates](https://docs.openshift.com/online/dev_guide/templates.html)/[Ktmpl](https://github.com/InQuicker/ktmpl)
* [Fabric8 client](https://github.com/fabric8io/kubernetes-client)
* [Kubegen](https://github.com/errordeveloper/kubegen)
* [kenv](https://github.com/thisendout/kenv)
* [Ansible](https://docs.ansible.com/ansible/kubernetes_module.html)
* [Puppet](https://forge.puppet.com/garethr/kubernetes/readme)
* [KPM](https://github.com/coreos/kpm)
* [Nulecule](https://github.com/projectatomic/nulecule)
* [Kedge](https://github.com/kedgeproject/kedge) ([OpenCompose](https://github.com/redhat-developer/opencompose) is deprecated)
* [Chartify](https://github.com/appscode/chartify)
* [Podex](https://github.com/kubernetes/contrib/tree/master/podex)
* [k8sec](https://github.com/dtan4/k8sec)
* [kb80r](https://github.com/UKHomeOffice/kb8or)
* [k8s-kotlin-dsl](https://github.com/fkorotkov/k8s-kotlin-dsl)
* [KY](https://github.com/stellaservice/ky)
* [Kploy](https://github.com/kubernauts/kploy)
* [Kdeploy](https://github.com/flexiant/kdeploy)
* [Kubernetes-deploy](https://github.com/Shopify/kubernetes-deploy)
* [Generator-kubegen](https://www.sesispla.net/blog/language/en/2017/07/introducing-generator-kubegen-a-kubernetes-configuration-file-booster-tool/)
* [K8comp](https://github.com/cststack/k8comp)
* [Kontemplate](https://github.com/tazjin/kontemplate)
* [Kexpand](https://github.com/kopeio/kexpand)
* [Forge](https://github.com/datawire/forge/)
* [Psykube](https://github.com/CommercialTribe/psykube)
* [Koki](http://koki.io)
* [Deploymentizer](https://github.com/InVisionApp/kit-deploymentizer)
* [generator-kubegen](https://github.com/sesispla/generator-kubegen)
* [Broadway](https://github.com/namely/broadway)
* [Srvexpand](https://github.com/kubernetes/kubernetes/pull/1980/files)
* [Rok8s-scripts](https://github.com/reactiveops/rok8s-scripts)
* [ERB-Hiera](https://roobert.github.io/2017/08/16/Kubernetes-Manifest-Templating-with-ERB-and-Hiera/)
* [k8s-icl](https://github.com/archipaorg/k8s-icl)
* [sed](https://stackoverflow.com/questions/42618087/how-to-parameterize-image-version-when-passing-yaml-for-container-creation)
* [envsubst](https://github.com/fabric8io/envsubst)
* [Jinja](https://github.com/tensorflow/ecosystem/tree/master/kubernetes)
* [spiff](https://github.com/cloudfoundry-incubator/spiff)

Additionally, a number of continuous deployment systems use their own formats and/or schemas.

The number of tools is a signal of demand for a customization solution, as well as lack of awareness of and/or dissatisfaction with existing tools. [Many prefer](https://news.ycombinator.com/item?id=15029086) to use the simplest tool that meets their needs. Most of these tools support customization via simple parameter substitution or a more complex configuration domain-specific language, while not adequately supporting the other customization strategies. The pitfalls of parameterization and domain-specific languages are discussed below.

### Parameterization pitfalls

After simply forking (or just cut&paste), parameterization is the most commonly used customization approach. We have [previously discussed](https://github.com/kubernetes/kubernetes/issues/11492) requirements for parameterization mechanisms, such as explicit declaration of parameters for easy discovery, documentation, and validation (e.g., for [form generation](https://github.com/kubernetes/kubernetes/issues/6487)).  It should also be straightforward to provide multiple sets of parameter values in support of variants and to manage them under version control, though many tools do not facilitate that.

Some existing template examples:

* [Openshift templates](https://github.com/openshift/library/tree/master/official) ([MariaDB example](https://github.com/luciddreamz/library/blob/master/official/mariadb/templates/mariadb-persistent.json))
* [Helm charts](https://github.com/kubernetes/charts/) ([Jenkins example](https://github.com/kubernetes/charts/blob/master/stable/jenkins/templates/jenkins-master-deployment.yaml))
* not Kubernetes, but a [Kafka Mesosphere Universe example](https://github.com/mesosphere/universe/blob/version-3.x/repo/packages/C/confluent-kafka/5/marathon.json.mustache)

Parameterization solutions are easy to implement and to use at small scale, but parameterized templates tend to become complex and difficult to maintain. Syntax-oblivious macro substitution (e.g., sed, jinja, envsubst) can be fragile, and parameter substitution sites generally have to be identified manually, which is tedious and error-prone, especially for the most common use cases, such as resource name prefixing.

Additionally, performing all customization via template parameters erodes template encapsulation. Some prior configuration-language design efforts made encapsulation a non-goal due to the widespread desire of users to override arbitrary parts of configurations. If used by enough people, someone will want to override each value in a template. Parameterizing every value in a template creates an alternative API schema that contains an out-of-date subset of the full API, and when [every value is a parameter](https://github.com/kubernetes/charts/blob/e002378c13e91bef4a3b0ba718c191ec791ce3f9/stable/artifactory/templates/artifactory-deployment.yaml), a template combined with its parameters is considerably less readable than the expanded result, and less friendly to data-manipulation scripts and tools.

### Pitfalls of configuration domain-specific languages (DSLs)

Since parameterization and file imports are common features of most configuration domain-specific languages (DSLs), they inherit the pitfalls of parameterization. The complex custom syntax (and/or libraries) of more sophisticated languages also tends to be more opaque, hiding information such as application topology from humans. Users generally need to understand the input language, transformations applied, and output generated, which is more complex for users to learn. Furthermore, custom-built languages [typically lack good tools](http://mikehadlow.blogspot.com/2012/05/configuration-complexity-clock.html) for refactoring, validation, testing, debugging, etc., and hard-coded translations are hard to maintain and keep up to date. And such syntax typically isn’t friendly to tools, for example [hiding information](https://github.com/kubernetes/kubernetes/issues/13241#issuecomment-233731291) about parameters and source dependencies, and is hostile to composition with other tools, configuration sources, configuration languages, runtime automation, and so on. The configuration source must be modified in order to customize additional properties or to add additional resources, which fosters closed, monolithic, fat configuration ecosystems and obstructs separation of concerns. This is especially true of tools and libraries that don’t facilitate post-processing of their output between pre-processing the DSL and actuation of the resulting API resources.

Additionally, the more powerful languages make it easy for users to shoot themselves in their feet. For instance, it can be easy to mix computation and data. Among other problems, embedded code renders the configuration unparsable by other tools (e.g., extraction, injection, manipulation, validation, diff, interpretation, reconciliation, conversion) and clients. Such languages also make it easy to reduce boilerplate, which can be useful, but when taken to the extreme, impairs readability and maintainability. Nested/inherited templates are seductive, for those languages that enable them, but very hard to make reusable and maintainable in practice. Finally, it can be tempting to use these capabilities for many purposes, such as changing defaults or introducing new abstractions, but this can create different and surprising behavior compared to direct API usage through CLIs, libraries, UIs, etc., and create accidental pseudo-APIs rather than intentional, actual APIs. If common needs can only be addressed using the configuration language, then the configuration transformer must be invoked by most clients, as opposed to using the API directly, which is contrary to the design of Kubernetes as an API-centric system.

Such languages are powerful and can perform complex transformations, but we found that to be a [mixed blessing within Google](http://research.google.com/pubs/pub44843.html). For instance, there have been many cases where users needed to generate configuration, manipulate configuration, backport altered API field settings into templates, integrate some kind of dynamic automation with declarative configuration, and so on. All of these scenarios were painful to implement with DSL templates in the way. Templates also created new abstractions, changed API default values, and diverged from the API in other ways that disoriented new users.

A few DSLs are in use in the Kubernetes community, including Go templates (used by Helm, discussed more below), [fluent DSLs](https://github.com/fabric8io/kubernetes-client), and [jsonnet](http://jsonnet.org/), which was inspired by [Google’s Borg configuration language](https://research.google.com/pubs/pub43438.html) ([more on its root language, GCL](http://alexandria.tue.nl/extra1/afstversl/wsk-i/bokharouss2008.pdf)). [Ksonnet-lib](https://github.com/ksonnet/ksonnet-lib) is a community project aimed at building Kubernetes-specific jsonnet libraries. Unfortunately, the examples (e.g., [nginx](https://github.com/ksonnet/ksonnet-lib/blob/master/examples/readme/hello-nginx.jsonnet)) appear more complex than the raw Kubernetes API YAML, so while it may provide more expressive power, it is less approachable. Databricks looks like [the biggest success case](https://databricks.com/blog/2017/06/26/declarative-infrastructure-jsonnet-templating-language.html) with jsonnet to date, and uses an approach that is admittedly more readable than ksonnet-lib, as is [Kubecfg](https://github.com/ksonnet/kubecfg). However, they all encourage users to author and manipulate configuration code written in a DSL rather than configuration data written in a familiar and easily manipulated format, and are unnecessarily complex for most use cases.

Helm is discussed below, with package management.

In case it’s not clear from the above, I do not consider configuration schemas expressed using common data formats such as JSON and YAML (sans use of substitution syntax) to be configuration DSLs.

## Configuration using REST API resource specifications

Given the pitfalls of parameterization and configuration DSLs, as mentioned at the beginning of this document, configuration tooling should manipulate configuration **data**, not convert configuration to code nor other marked-up syntax, and, in the case of Kubernetes, this data should primarily contain specifications of the **literal Kubernetes API resources** required to deploy the application in the manner desired by the user. The Kubernetes API and CLI (kubectl) were designed to support this model, and our documentation and examples use this approach.

[Kubernetes’s API](https://github.com/kubernetes/community/blob/master/contributors/design-proposals/architecture.md#cluster-control-plane-aka-master) provides IaaS-like container-centric primitives such as Pods, Services, and Ingress, and also lifecycle controllers to support orchestration (self-healing, scaling, updates, termination) of common types of workloads, such as ReplicaSet (simple fungible/stateless app manager), Deployment (orchestrates updates of stateless apps), Job (batch), CronJob (cron), DaemonSet (cluster services), StatefulSet (stateful apps), and [custom third-party controllers/operators](https://coreos.com/blog/introducing-operators.html). The workload controllers, such as Deployment, support declarative upgrades using production-grade strategies such as rolling update, so that the client doesn’t need to perform complex orchestration in the common case. (And we’re moving [proven kubectl features to controllers](https://github.com/kubernetes/kubernetes/issues/12143), generally.) We also deliberately decoupled service naming/discovery and load balancing from application implementation in order to maximize deployment flexibility, which should be preserved by the configuration mechanism.

[Kubectl apply](https://github.com/kubernetes/kubernetes/issues/15894)  [was designed](https://github.com/kubernetes/kubernetes/issues/1702) ([original proposal](https://github.com/kubernetes/kubernetes/issues/1178)) to support declarative updates without clobbering operationally and/or automatically set desired state. Properties not explicitly specified by the user are free to be changed by automated and other out-of-band mechanisms. Apply is implemented as a 3-way merge of the user’s previous configuration, the new configuration, and the live state.

We [chose this simple approach of using literal API resource specifications](https://github.com/kubernetes/kubernetes/pull/1007/files) for the following reasons:

* KISS: It was simple and natural, given that we designed the API to support CRUD on declarative primitives, and Kubernetes uses the API representation in all scenarios where API resources need to be serialized (e.g., in persistent cluster storage).
* It didn’t require users to learn multiple different schemas, the API and another configuration format. We believe many/most production users will eventually want to use the API, and knowledge of the API transfers to other clients and tools. It doesn’t obfuscate the API, which is relatively easy to read.
* It automatically stays up to date with the API, automatically supports all Kubernetes resources, versions, extensions, etc., and can be automatically converted to new API versions.
* It could share mechanisms with other clients (e.g., Swagger/OpenAPI, which is used for schema validation), which are now supported in several languages: Go, Python, Java, …
* Declarative configuration is only one interface to the system. There are also CLIs (e.g., kubectl), UIs (e.g., dashboard), mobile apps, chat bots, controllers, admission controllers, Operators, deployment pipelines, etc. Those clients will (and should) target the API. The user will need to interact with the system in terms of the API in these other scenarios.
* The API serves as a well defined intermediate representation, pre- and post-creation, with a documented deprecation policy. Tools, libraries, controllers, UI wizards, etc. can be built on top, leaving room for exploration and innovation within the community. Example API-based transformations include:
    * Overlay application: kubectl patch
    * Generic resource tooling: kubectl label, kubectl annotate
    * Common-case tooling: kubectl set image, kubectl set resources
    * Dynamic pod transformations: LimitRange, PodSecurityPolicy, PodPreset
    * Admission controllers and initializers
    * API-based controllers, higher-level APIs, and controllers driven by custom resources
    * Automation: horizontal and [vertical pod autoscaling](https://github.com/kubernetes/community/pull/338)
* It is inherently composable: just add more resource manifests, in the same file or another file. No embedded imports required.

Of course, there are downsides to the approach:

* Users need to learn some API schema details, though we believe operators will want to learn them, anyway.
* The API schema does contain a fair bit of boilerplate, though it could be auto-generated and generally increases clarity.
* The API introduces a significant number of concepts, though they exist for good reasons.
* The API has no direct representation of common generation steps (e.g., generation of ConfigMap or Secret resources from source data), though these can be described in a declarative format using API conventions, as we do with component configuration in Kubernetes.
* It is harder to fix warts in the API than to paper over them. Fixing "bugs" may break compatibility (e.g., as with changing the default imagePullPolicy). However, the API is versioned, so it is not impossible, and fixing the API benefits all clients, tools, UIs, etc.
* JSON is cumbersome and some users find YAML to be error-prone to write. It would also be nice to support a less error-prone data syntax than YAML, such as [Relaxed JSON](https://github.com/phadej/relaxed-json), [HJson](https://hjson.org/), [HCL](https://github.com/hashicorp/hcl), [StrictYAML](https://github.com/crdoconnor/strictyaml/blob/master/FAQ.rst), or [YAML2](https://github.com/yaml/YAML2/wiki/Goals). However, one major disadvantage would be the lack of library support in multiple languages. HCL also wouldn’t directly map to our API schema due to our avoidance of maps. Perhaps there are there YAML conventions that could result in less error-prone specifications.

## What needs to be improved?

While the basic mechanisms for this approach are in place, a number of common use cases could be made easier. Most user complaints are around discovering what features exist (especially annotations), documentation of and examples using those features, generating/finding skeleton resource specifications (including boilerplate and commonly needed features), formatting and validation of resource specifications, and determining appropriate cpu and memory resource requests and limits. Specific user scenarios are discussed below.

### Bespoke application deployment

Deployment of bespoke applications involves multiple steps:

1. Build the container image
2. Generate and/or modify Kubernetes API resource specifications to use the new image
3. Reconcile those resources with a Kubernetes cluster

Step 1, building the image, is out of scope for Kubernetes. Step 3 is covered by kubectl apply. Some tools in the ecosystem, such as [Draft](https://github.com/Azure/draft), combine the 3 steps.

Kubectl contains ["generator" commands](/contributors/devel/sig-cli/kubectl-conventions.md#generators), such as [kubectl run](https://kubernetes.io/docs/user-guide/kubectl/v1.7/#run), expose, various create commands, to create commonly needed Kubernetes resource configurations. However, they also don’t help users understand current best practices and conventions, such as proper label and annotation usage. This is partly a matter of updating them and partly one of making the generated resources suitable for consumption by new users. Options supporting declarative output, such as dry run, local, export, etc., don’t currently produce clean, readable, reusable resource specifications ([example](https://blog.heptio.com/using-kubectl-to-jumpstart-a-yaml-file-heptioprotip-6f5b8a63a3ea))**.** We should clean them up.

Openshift provides a tool, [oc new-app](https://docs.openshift.com/enterprise/3.1/dev_guide/new_app.html), that can pull source-code templates, [detect](https://github.com/kubernetes/kubernetes/issues/14801)[ application types](https://github.com/kubernetes/kubernetes/issues/14801) and create Kubernetes resources for applications from source and from container images. [podex](https://github.com/kubernetes/contrib/tree/master/podex) was built to extract basic information from an image to facilitate creation of default Kubernetes resources, but hasn’t been kept up to date. Similar resource generation tools would be useful for getting started, and even just [validating that the image really exists](https://github.com/kubernetes/kubernetes/issues/12428) would reduce user error.

For updating the image in an existing deployment, kubectl set image works both on the live state and locally. However, we should [make the image optional](https://github.com/kubernetes/kubernetes/pull/47246) in controllers so that the image could be updated independently of kubectl apply, if desired. And, we need to [automate image tag-to-digest translation](https://github.com/kubernetes/kubernetes/issues/33664) ([original issue](https://github.com/kubernetes/kubernetes/issues/1697)), which is the approach we’d expect users to use in production, as opposed to just immediately re-pulling the new image and restarting all existing containers simultaneously. We should keep the original tag in an imageStream annotation, which could eventually become a field.

### Continuous deployment

In addition to PaaSes, such as [Openshift](https://blog.openshift.com/openshift-3-3-pipelines-deep-dive/) and [Deis Workflow](https://github.com/deis/workflow), numerous continuous deployment systems have been integrated with Kubernetes, such as [Google Container Builder](https://github.com/GoogleCloudPlatform/cloud-builders/tree/master/kubectl), [Jenkins](https://github.com/GoogleCloudPlatform/continuous-deployment-on-kubernetes), [Gitlab](https://about.gitlab.com/2016/11/14/idea-to-production/), [Wercker](http://www.wercker.com/integrations/kubernetes), [Drone](https://open.blogs.nytimes.com/2017/01/12/continuous-deployment-to-google-cloud-platform-with-drone/), [Kit](https://invisionapp.github.io/kit/), [Bitbucket Pipelines](https://confluence.atlassian.com/bitbucket/deploy-to-kubernetes-892623297.html), [Codeship](https://blog.codeship.com/continuous-deployment-of-docker-apps-to-kubernetes/), [Shippable](https://www.shippable.com/kubernetes.html), [SemaphoreCI](https://semaphoreci.com/community/tutorials/continuous-deployment-with-google-container-engine-and-kubernetes), [Appscode](https://appscode.com/products/cloud-deployment/), [Kontinuous](https://github.com/AcalephStorage/kontinuous), [ContinuousPipe](https://continuouspipe.io/), [CodeFresh](https://docs.codefresh.io/docs/kubernetes#section-deploy-to-kubernetes), [CloudMunch](https://www.cloudmunch.com/continuous-delivery-for-kubernetes/), [Distelli](https://www.distelli.com/kubernetes/), [AppLariat](https://www.applariat.com/ci-cd-applariat-travis-gke-kubernetes/), [Weave Flux](https://github.com/weaveworks/flux), and [Argo](https://argoproj.github.io/argo-site/#/). Developers usually favor simplicity, whereas operators have more requirements, such as multi-stage deployment pipelines, deployment environment management (e.g., staging and production), and canary analysis. In either case, users need to be able to deploy both updated images and configuration updates, ideally using the same workflow. [Weave Flux](https://github.com/weaveworks/flux) and [Kube-applier](https://blog.box.com/blog/introducing-kube-applier-declarative-configuration-for-kubernetes/) support unified continuous deployment of this style. In other CD systems a unified flow may be achievable by making the image deployment step perform a local kubectl set image (or equivalent) and commit the change to the configuration, and then use another build/deployment trigger on the configuration repository to invoke kubectl apply --prune.

### Migrating from Docker Compose

Some developers like Docker’s Compose format as a simplified all-in-one configuration schema, or are at least already familiar with it. Kubernetes supports the format using the [Kompose tool](https://github.com/kubernetes/kompose), which provides an easy migration path for these developers by translating the format to Kubernetes resource specifications.

The Compose format, even with extensions (e.g., replica counts, pod groupings, controller types), is inherently much more limited in expressivity than Kubernetes-native resource specifications, so users would not want to use it forever in production. But it provides a useful onramp, without introducing [yet another schema](https://github.com/kubernetes/kubernetes/pull/1980#issuecomment-60457567) to the community. We could potentially increase usage by including it in a [client-tool release bundle](https://github.com/kubernetes/release/issues/3).

### Reconciliation of multiple resources and multiple files

Most applications require multiple Kubernetes resources. Although kubectl supports multiple resources in a single file, most users store the resource specifications using one resource per file, for a number of reasons:

* It was the approach used by all of our early application-stack examples
* It provides more control by making it easier to specify which resources to operate on
* It’s inherently composable -- just add more files

The control issue should be addressed by adding support to select resources to mutate by label selector, name, and resource types, which has been planned from the beginning but hasn’t yet been fully implemented. However, we should also [expand and improve kubectl’s support for input from multiple files](https://github.com/kubernetes/kubernetes/issues/24649).

### Declarative updates

Kubectl apply (and strategic merge patch, upon which apply is built) has a [number of bugs and shortcomings](https://github.com/kubernetes/kubernetes/issues/35234), which we are fixing, since it is the underpinning of many things (declarative configuration, add-on management, controller diffs). Eventually we need [true API support](https://github.com/kubernetes/kubernetes/issues/17333) for apply so that clients can simply PUT their resource manifests and it can be used as the fundamental primitive for declarative updates for all clients. One of the trickier issues we should address with apply is how to handle [controller selector changes](https://github.com/kubernetes/kubernetes/issues/26202). We are likely to forbid changes for now, as we do with resource name changes.

Kubectl should also operate on resources in an intelligent order when presented with multiple resources. While we’ve tried to avoid creation-order dependencies, they do exist in a few places, such as with namespaces, custom resource definitions, and ownerReferences.

### ConfigMap and Secret updates

We need a declarative syntax for regenerating [Secrets](https://github.com/kubernetes/kubernetes/issues/24744) and [ConfigMaps](https://github.com/kubernetes/kubernetes/issues/30337) from their source files that could be used with apply, and provide easier ways to [roll out new ConfigMaps and garbage collect unneeded ones](https://github.com/kubernetes/kubernetes/issues/22368). This could be embedded in a manifest file, which we need for "package" metadata (see [Addon manager proposal](https://docs.google.com/document/d/1Laov9RCOPIexxTMACG6Ffkko9sFMrrZ2ClWEecjYYVg/edit) and [Helm chart.yaml](https://github.com/kubernetes/helm/blob/master/docs/charts.md)). There also needs to be an easier way to [generate names of the new resources](https://github.com/kubernetes/kubernetes/pull/49961) and to update references to ConfigMaps and Secrets, such as in env and volumes. This could be done via new kubectl set commands, but users primarily need the “stream” update model, as with images.

### Determining success/failure

The declarative, [asynchronous control-loop-based approach](https://docs.google.com/presentation/d/1oPZ4rznkBe86O4rPwD2CWgqgMuaSXguIBHIE7Y0TKVc/edit#slide=id.g21b1f16809_3_155) makes it more challenging for the user to determine whether the change they made succeeded or failed, or the system is still converging towards the new desired state. Enough status information needs to be reported such that progress and problems are visible to controllers watching the status, and the status needs to be reported in a consistent enough way that a [general-purpose mechanism](https://github.com/kubernetes/kubernetes/issues/34363) can be built that works for arbitrary API types following Kubernetes API conventions. [Third-party attempts](https://github.com/Mirantis/k8s-AppController#dependencies) to monitor the status generally are not implemented correctly, since Kubernetes’s extensible API model requires exposing distributed-system effects to clients. This complexity can be seen all over our [end-to-end tests](https://github.com/kubernetes/kubernetes/blob/master/test/utils/deployment.go#L74), which have been made robust over many thousands of executions. Definitely authors of individual application configurations should not be forced to figure out how to implement such checks, as they currently do in Helm charts (--wait, test).

### Configuration customization

The strategy for customization involves the following main approaches:

1. Fork or simply copy the resource specifications, and then locally modify them, imperatively, declaratively, or manually, in order to reuse off-the-shelf configuration. To facilitate these modifications, we should:
    * Automate common customizations, especially [name prefixing and label injection](https://github.com/kubernetes/kubernetes/issues/1698) (including selectors, pod template labels, and object references), which would address the most common substitutions in existing templates
    * Fix rough edges for local mutation via kubectl get --export and [kubectl set](https://github.com/kubernetes/kubernetes/issues/21648) ([--dry-run](https://github.com/kubernetes/kubernetes/issues/11488), --local, -o yaml), and enable kubectl to directly update files on disk
    * Build fork/branch management tooling for common workflows, such as branch creation, cherrypicking (e.g., to copy configuration changes from a staging to production branch), rebasing, etc., perhaps as a plugin to kubectl.
    * Build/improve structural diff, conflict detection, validation (e.g., [kubeval](https://github.com/garethr/kubeval), [ConfigMap element properties](https://github.com/kubernetes/kubernetes/issues/4210)), and comprehension tools
2. Resource overlays, for instantiating multiple variants. Kubectl patch already works locally using strategic merge patch, so the overlays have the same structure as the base resources. The main feature needed to facilitate that is automatic pairing of overlays with the resources they should patch.

Fork provides one-time customization, which is the most common case. Overlay patches provide deploy-time customization. These techniques can be combined with dynamic customization (PodPreset, other admission controllers, third-party controllers, etc.) and run-time customization (initContainers and entrypoint.sh scripts inside containers).

Benefits of these approaches:

* Easier for app developers and operators to build initial configurations (no special template syntax)
* Compatible with existing project tooling and conventions, and easy to read since it doesn’t obfuscate the API and doesn’t force users to learn a new way to configure their applications
* Supports best practices
* Handles cases the [original configuration author didn’t envision](http://blog.shippable.com/the-new-devops-matrix-from-hell)
* Handles cases where original author changes things that break existing users
* Supports composition by adding resources: secrets, configmaps, autoscaling
* Supports injection of operational concerns, such as node affinity/anti-affinity and tolerations
* Supports selection among alternatives, and multiple simultaneous versions
* Supports canaries and multi-cluster deployment
* Usable for [add-on management](https://github.com/kubernetes/kubernetes/issues/23233), by avoiding [obstacles that Helm has](https://github.com/kubernetes/kubernetes/issues/23233#issuecomment-285524825), and should eliminate the need for the EnsureExists behavior

#### What about parameterization?

An area where more investigation is needed is explicit inline parameter substitution, which, while overused and should be rendered unnecessary by the capabilities described above, is [frequently requested](https://stackoverflow.com/questions/44832085/passing-variables-to-args-field-in-a-yaml-file-kubernetes) and has been reinvented many times by the community.

A [simple parameterization approach derived from Openshift’s design](https://github.com/kubernetes/community/blob/master/contributors/design-proposals/templates.md) was approved because it was constrained in functionality and solved other problems (e.g., instantiation of resource variants by other controllers, [project templates in Openshift](https://github.com/openshift/training/blob/master/content/default-project-template.yaml)). That proposal explains some of the reasoning behind the design tradeoffs, as well as the [use cases](https://github.com/kubernetes/community/blob/master/contributors/design-proposals/templates.md#use-cases). Work started, but was abandoned, though there is an independent [client-based implementation](https://github.com/InQuicker/ktmpl). However, the Template resource wrapped the resource specifications in another object, which is suboptimal, since transformations would then need to be able to deal with standalone resources, Lists of resources, and Templates, or would need to be applied post-instantiation, and it couldn’t be represented using multiple files, as users prefer.

What is more problematic is that our client libraries, schema validators, yaml/json parsers/decoders, initializers, and protobuf encodings all require that all specified fields have valid values, so parameters cannot currently be left in non-string (e.g., int, bool) fields in actual resources. Additionally, the API server requires at least complete/final resource names to be specified, and strategic merge also requires all merge keys to be specified. Therefore, some amount of pre-instantiation (though not necessarily client-side) transformation is necessary to create valid resources, and we may want to explicitly store the output, or the fields should just contain the default values initially. Parameterized fields could be automatically converted to patches to produce valid resources. Such a transformation could be made reversible, unlike traditional substitution approaches, since the patches could be preserved (e.g., using annotations). The Template API supported the declaration of parameter names, display names, descriptions, default values, required/optional, and types (string, int, bool, base64), and both string and raw json substitutions. If we were to update that specification, we could use the same mechanism for both parameter validation and ConfigMap validation, so that the same mechanism could be used for env substitution and substitution of values of other fields. As mentioned in the [env validation issue](https://github.com/kubernetes/kubernetes/issues/4210#issuecomment-305555589), we should consider a subset of [JSON schema](http://json-schema.org/example1.html), which we’ll probably use for CRD. The only [unsupported attribute](https://tools.ietf.org/html/draft-wright-json-schema-validation-00) appears to be the display name, which is non-critical. [Base64 could be represented using media](http://json-schema.org/latest/json-schema-hypermedia.html#rfc.section.5.3.2). That could be useful as a common parameter schema to facilitate parameter discovery and documentation that is independent of the substitution syntax and mechanism ([example from Deployment Manager](https://github.com/GoogleCloudPlatform/deploymentmanager-samples/blob/master/templates/replicated_service.py.schema)).

Without parameters how would we support a click-to-deploy experience? People who are kicking the tires, have undemanding use cases, are learning, etc. are unlikely to know what customization they want to perform initially, if they even need any. The main information users need to provide is the name prefix they want to apply. Otherwise, choosing among a few alternatives would suit their needs better than parameters. The overlay approach should support that pretty well. Beyond that, I suggest kicking users over to a Kubernetes-specific configuration wizard or schema-aware IDE, and/or support a fork workflow.

The other application-definition [use cases mentioned in the Template proposal](https://github.com/kubernetes/community/blob/master/contributors/design-proposals/templates.md#use-cases) are achievable without parameterization, as well.

#### What about application configuration generation?

A number of legacy applications have configuration mechanisms that couple application options and information about the deployment environment. In such cases, a ConfigMap containing the configuration data is not sufficient, since the runtime information (e.g., identities, secrets, service addresses) must be incorporated. There are a [number of tools used for this purpose outside Kubernetes](https://github.com/kubernetes/kubernetes/issues/2068). However, in Kubernetes, they would have to be run as Pod initContainers, sidecar containers, or container [entrypoint.sh init scripts](https://github.com/kubernetes/kubernetes/issues/30716). As this is only a need of some legacy applications, we should not complicate Kubernetes itself to solve it. Instead, we should be prepared to recommend a third-party tool, or provide one, and ensure the downward API provides the information it would need.

#### What about [package management](https://medium.com/@sdboyer/so-you-want-to-write-a-package-manager-4ae9c17d9527) and Helm?

[Helm](https://github.com/kubernetes/helm/blob/master/docs/chart_repository.md), [KPM](https://github.com/coreos/kpm), [App Registry](https://github.com/app-registry), [Kubepack](https://kubepack.com/), and [DCOS](https://docs.mesosphere.com/1.7/usage/managing-services/) (for Mesos) bundle whitebox off-the-shelf application configurations into **_packages_**. However, unlike traditional artifact repositories, which store and serve generated build artifacts, configurations are primarily human-authored. As mentioned above, it is industry best practice to manage such configurations using version control systems, and Helm package repositories are backed by source code repositories. (Example: [MariaDB](https://github.com/kubernetes/charts/tree/master/stable/mariadb).)

Advantages of packages:

1. Package formats add structure to raw Kubernetes primitives, which are deliberately flexible and freeform
    * Starter resource specifications that illustrate API schema and best practices
    * Labels for application topology (e.g., app, role, tier, track, env) -- similar to the goals of [Label Schema](http://label-schema.org/rc1/)
    * File organization and manifest (list of files), to make it easier for users to navigate larger collections of application specifications, to reduce the need for tooling to search for information, and to facilitate segregation of resources from other artifacts (e.g., container sources)
    * Application metadata: name, authors, description, icon, version, source(s), etc.
    * Application lifecycle operations: build, test, debug, up, upgrade, down, etc.
1. [Package registries/repositories](https://github.com/app-registry/spec) facilitate [discovery](https://youtu.be/zGJsXyzE5A8?t=1159) of off-the-shelf applications and of their dependencies
    * Scattered source repos are hard to find
    * Ideally it would be possible to map the format type to a container containing the tool that understands the format.

Helm is probably the most-used configuration tool other than kubectl, many [application charts](https://github.com/kubernetes/charts) have been developed (as with the [Openshift template library](https://github.com/openshift/library)), and there is an ecosystem growing around it (e.g., [chartify](https://github.com/appscode/chartify), [helmfile](https://github.com/roboll/helmfile), [landscaper](https://github.com/Eneco/landscaper), [draughtsman](https://github.com/giantswarm/draughtsman), [chartmuseum](https://github.com/chartmuseum/chartmuseum)). Helm’s users like the familiar analogy to package management and the structure that it provides. However, while Helm is useful and is the most comprehensive tool, it isn’t suitable for all use cases, such as [add-on management](https://github.com/kubernetes/kubernetes/issues/23233). The biggest obstacle is that its [non-Kubernetes-compatible API](https://kubernetes.io/docs/concepts/overview/what-is-kubernetes/#what-kubernetes-is-not)[ and DSL syntax push it out of Kubernetes proper into the Kubernetes ecosystem](https://kubernetes.io/docs/concepts/overview/what-is-kubernetes/#what-kubernetes-is-not). And, as much as Helm is targeting only Kubernetes, it takes little advantage of that. Additionally, scenarios we’d like to support better include chart authoring (prefer simpler syntax and more straightforward management under version control), operational customization (e.g., via scripting, [forking](https://github.com/kubernetes/helm/issues/2554), or patching/injection), deployment pipelines (e.g., [canaries](https://groups.google.com/forum/#!topic/kubernetes-sig-apps/ouqXYXdsPYw)), multi-cluster / [multi-environment](https://groups.google.com/d/msg/kubernetes-users/GPaGOGxCDD8/NbNL-NPhCAAJ) deployment, and multi-tenancy.

Helm provides functionality covering several areas:

* Package conventions: metadata (e.g., name, version, descriptions, icons; Openshift has [something similar](https://github.com/luciddreamz/library/blob/master/official/java/templates/openjdk18-web-basic-s2i.json#L10)), labels, file organization
* Package bundling, unbundling, and hosting
* Package discovery: search and browse
* [Dependency management](https://github.com/kubernetes/helm/blob/master/docs/charts.md#chart-dependencies)
* Application lifecycle management framework: build, install, uninstall, upgrade, test, etc.
    * a non-container-centric example of that would be [ElasticBox](https://www.ctl.io/knowledge-base/cloud-application-manager/automating-deployments/start-stop-and-upgrade-boxes/)
* Kubernetes drivers for creation, update, deletion, etc.
* Template expansion / schema transformation
* (It’s currently lacking a formal parameter schema.)

It's useful for Helm to provide an integrated framework, but the independent functions could be decoupled, and re-bundled into multiple separate tools:

* Package management -- search, browse, bundle, push, and pull of off-the-shelf application packages and their dependencies.
* Application lifecycle management -- install, delete, upgrade, rollback -- and pre- and post- hooks for each of those lifecycle transitions, and success/failure tests.
* Configuration customization via parameter substitution, aka template expansion, aka rendering.

That would enable the package-like structure and conventions to be used with raw declarative management via kubectl or other tool that linked in its [business logic](https://github.com/kubernetes/kubernetes/issues/7311), for the lifecycle management to be used without the template expansion, and the template expansion to be used in declarative workflows without the lifecycle management. Support for both client-only and server-side operation and migration from grpc to Kubernetes API extension mechanisms would further expand the addressable use cases.

([Newer proposal, presented at the Helm Summit](https://docs.google.com/presentation/d/10dp4hKciccincnH6pAFf7t31s82iNvtt_mwhlUbeCDw/edit#slide=id.p).)

#### What about the service broker?

The [Open Service Broker API](https://openservicebrokerapi.org/) provides a standardized way to provision and bind to blackbox services. It enables late binding of clients to service providers and enables usage of higher-level application services (e.g., caches, databases, messaging systems, object stores) portably, mitigating lock-in and facilitating hybrid and multi-cloud usage of these services, extending the portability of cloud-native applications running on Kubernetes. The service broker is not intended to be a solution for whitebox applications that require any level of management by the user. That degree of abstraction/encapsulation requires full automation, essentially creating a software appliance (cf. [autonomic computing](https://en.wikipedia.org/wiki/Autonomic_computing)): autoscaling, auto-repair, auto-update, automatic monitoring / logging / alerting integration, etc. Operators, initializers, autoscalers, and other automation may eventually achieve this, and we need to for [cluster add-ons](https://github.com/kubernetes/kubernetes/issues/23233) and other [self-hosted components](https://github.com/kubernetes/kubernetes/issues/246), but the typical off-the-shelf application template doesn’t achieve that.

#### What about configurations with high cyclomatic complexity or massive numbers of variants?

Consider more automation, such as autoscaling, self-configuration, etc. to reduce the amount of explicit configuration necessary. One could also write a program in some widely used conventional programming language to generate the resource specifications. It’s more likely to have IDE support, test frameworks, documentation generators, etc. than a DSL. Better yet, create composable transformations, applying [the Unix Philosophy](https://en.wikipedia.org/wiki/Unix_philosophy#Eric_Raymond.E2.80.99s_17_Unix_Rules). In any case, don’t look for a silver bullet to solve all configuration-related problems. Decouple solutions instead.

#### What about providing an intentionally restrictive simplified, tailored developer experience to streamline a specific use case, environment, workflow, etc.?

This is essentially a [DIY PaaS](https://kubernetes.io/blog/2017/02/caas-the-foundation-for-next-gen-paas/). Write a configuration generator, either client-side or using CRDs ([example](https://github.com/pearsontechnology/environment-operator/blob/dev/User_Guide.md)). The effort involved to document the format, validate it, test it, etc. is similar to building a new API, but I could imagine someone eventually building a SDK to make that easier.

#### What about more sophisticated deployment orchestration?

Deployment pipelines, [canary deployments](https://groups.google.com/forum/#!topic/kubernetes-sig-apps/ouqXYXdsPYw), [blue-green deployments](https://groups.google.com/forum/#!topic/kubernetes-sig-apps/mwIq9bpwNCA), dependency-based orchestration, event-driven orchestrations, and [workflow-driven orchestration](https://github.com/kubernetes/kubernetes/issues/1704) should be able to use the building blocks discussed in this document. [AppController](https://github.com/Mirantis/k8s-AppController) and [Smith](https://github.com/atlassian/smith) are examples of tools built by the community.

#### What about UI wizards, IDE integration, application frameworks, etc.?

Representing configuration using the literal API types should facilitate programmatic manipulation of the configuration via user-friendly tools, such as UI wizards (e.g., [dashboard](https://kubernetes.io/docs/tasks/access-application-cluster/web-ui-dashboard/#deploying-containerized-applications), [Yipee.io](https://yipee.io/), and many CD tools, such as [Distelli](https://www.distelli.com/docs/k8s/add-container-to-a-project/)) and IDEs (e.g., [VSCode](https://www.youtube.com/watch?v=QfqS9OSVWGs), [IntelliJ](https://github.com/tinselspoon/intellij-kubernetes)), as well as configuration generation and manipulation by application frameworks (e.g., [Spring Cloud](https://github.com/fabric8io/spring-cloud-kubernetes)).
