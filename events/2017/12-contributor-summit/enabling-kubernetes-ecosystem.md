# Kubernetes Ecosystem

Notes by @directxman12

* How do we e.g. have an object validator

    * Went looking, found 2, second b/c first didn’t work well enough

* How do we enable people to build tools that consume and assist with working with Kubernetes

    * Kube-fuse, validatators, etc

    * How do we make those discoverable

        * Difficult to find via GitHub search

        * Difficult to find via search engines (find blog posts, etc, and not the tools)

* Proposal: structured content (tags, categories, search) for registering and discoverability

* Are there examples of ecosystems that we can follow

    * Package managers (PyPI, crates.io, NPM)

        * Doesn’t quite fit one-to-one b/c some stuff is best practices, etc

            * Docs, stuff like CNI plugins, etc doesn’t actually work well the same view as things that run on Kubernetes

            * At the basic point, just need to find what’s there

        * Traditional package managers focus on consuming the packages

            * If you don’t have packaging, it’s hard to approach (see also: Freshmeat problems)

    * Hadoop: classes that map the Hadoop ecosystem, infographs

    * Wordpress, Drupal are good examples

    * Ansible Galaxy

    * Chrome Extensions/Eclipse plugins

        * Look for things tagged, if comments say "broken" and last update is old, don’t use it	

    * Packagist (PHP package repository)

        * Integrates with GitHub, pulls in details about updates, README, extensibility, etc from GitHub

        * Just helps with discoverability

    * Steam

        * Has curated lists by users

* Opinion: End users need focused, distributable bundles

    * Most people don’t need to do all of everything

    * Different systems for apps vs critical infrastructure

        * Critical infrastructure doesn’t change much

        * Still need to discover when initially building out your system

* Issue: people get overwhelmed with choice

    * We don’t want to endorse things -- users should choose

    * We could let people rank/vote/etc

        * For example, what’s wrong with an "awesome list"

    * Need to look at human-consumable media, not necessarily machine-consumable

* Question: do we curate

    * Do we require "discoverable" things to be maintained/use a CLA/etc?

    * Opinion: no, we can’t possibly curate everything ourselves

* It’s problematic if you can discover new things, but they’re not supported by your Kubernetes distros

    * Not as much a problem for apps stuff, but harder for infrastructure

* Doesn’t GitHub have labels, stars, etc

    * Yes!

    * We could just say "always label your GitHub repo with a XYZ label if you’re a CRI plugin"

    * Comes a bit back to curration to discover benefits of each

    * Enterprises, gitlab make this infeasible

* Core infrastructure is a bit of an edge case, perhaps focus on addons like logging, monitoring, etc

    * Still comes back to: "if distro doesn’t support well, will it still work"

* Issue: there’s things people don’t know they can choose

    * E.g. external DNS providers

* Have a partially curated list of topics, but not curate actual content

    * Maybe leave that up to distros (have different collections of options -- e.g. open-source only, etc)

    * Have "awesome-kubernetes"

* Let SIGs curate their lists?

    * Having all the SIG be different is difficult and confusing

    * SIG leads don’t necessarily want to be the gatekeepers

    * We don’t necessarily want to tempt SIG leads with being the gatekeepers

* If we have something "official", people will assume that stuff is tested, even if it’s not

* Can we just have distributions that have way more options (a la Linux distros)

    * There are currently 34 conformant distros

* If we have ecosystem.k8s.io it’s really easy for people to find how to find things, otherwise could be hard

    * E.g. people don’t necessarily know awesome lists are a thing to search for

* Someone should do a prototype, and then we can have the conversation

* Question: where should this convo continue

    * SIG Apps?

    * Breakout group from SIG Apps?

