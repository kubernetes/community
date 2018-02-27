Onboarding Developers through Better Documentation

@ K8s Contributor Summit, 12/5

Note Taker: Jared Bhatti ([jaredbhatti](https://github.com/jaredbhatti)), Andrew Chen ([chenopis](https://github.com/chenopis)), Solly Ross ([directxman12](https://github.com/DirectXMan12))

[[TOC]]

### Goal

* Understand how we’re currently onboarding developers now

* Understand the "rough edges" of the onboarding experience for new users

* Understand how we can better target the needs of our audience

* Understand if our contributor process will meet the documentation needs. 

Note: the focus is on documentation, so we won’t be able to act on suggestions outside of that (but we will consider them!). 

### Initial thoughts

* Where we were at the beginning of the 2017: ([@1.4](https://v1-4.docs.kubernetes.io/docs/))

* Where we are now: ([@ 1.8](https://kubernetes.io/docs/home/))

    * We are tied into the launch process

    * We have analytics on our docs

    * We have a full time writer and sig docs contrib group

    * Everything is in templates, we have a style guide

    * Better infrastructure (netlify)

### Meeting Notes

* Introductions

    * Jared: SIG docs maintainer

    * Andrew: SIG docs maintainer, techwriter

    * Radika: 1.8 release team, SIG docs member

    * Phil Wittrock: interested b/c it determines how system is used

    * ??

    * Paul Morie: interested b/c hit challenges that could be avoided with better dev docs (patterns used in Kube can’t be easily found on internet)

    * Nikita: work on CRDs

    * ??: docs are a good intro to core Kube codebase

    * Michael Rubin

    * ??: dashboard contributor, onboarding is a key part of the dashboard experience

    * Steve Wong: storage SIG, on-prem SIG, trying to get head around on-prem, docs geared towards cloud

    * Solly Ross (@directxman12): want to avoid changes to fiddly bits confusing people

    * Morgan Bauer: had issues with people changing fundamental bits out from under other contributors

    * Sahdev Zala: docs are important for new contributors, first place they look, contribex

    * Brad Topol: like explaining stuff, kube conformance WG

    * Tomas: working on side project, looking for inspiration on docs for that

    * Garrett: ContribEx, looking to reduce the question "how do I get started"

    * Stefan (@sttts): where do we push docs changes to mechanics

    * Josh Berkus: ContribEx

* contributor docs

    * Avoid the process for finding which of the 49 people have the right info 6 months after it merges

    * Have contributor docs which merge with "internal features" (e.g. API registration changes)

    * Have a flag for internal release notes, bot which checks

* Issue: opinion: we don’t want to write dev docs because they’ll change

    * Issue is it’s hard to capture every little thing

    * Good start is documentation on generalities (80%/20%): e.g building an API server

        * Started out with no docs on how to do that (has gotten better)

* Big concepts don’t change a lot, describe those

    * Ask questions:

        * What are the big nouns

        * What are the big verbs for those nouns

        * How are the nouns relate

    * Don’t fall into trap of (for example) tracepoints (accidentally creating a lot of small little fiddly points)

* Need a way for "heads up this is changing". If you have questions, ask these specific people. 

    * Very SIG dependent (some do meetings, others mailing-list focused, some have a low bus number)

* How is your SIG interacting with docs

    * Service Catalog

        * In docs folder

        * How do we interact with docs

        * Depends on moving into kubernetes org

* Issue: lots of docs with bits and pieces

    * Not organized into "as a developer, you’re interested in these links"

    * Spend a week, come up with a ToC

    * Many docs in different repos

    * Should have guidelines on where to put things

    * Should have templates for how to write different types of docs

* Issue: wrong abstraction layers are exposed

    * Documenting a complex system has less payoff than simplifying

* Issue: extending Kube blurs the line between contributor and user

    * Put more docs into the same place

    * "Crafting User Journeys": user stories buckets: Customer User Journeys ([project](https://github.com/kubernetes/website/projects/5))

    * Need champion for different personas

* Issue: people (e.g. Storage SIG) want to write docs, but they don’t know where to start

    * Both contributors and users, and both (plugin writing, for example, or how to provision storage securely)

    * Templates/guidelines mentioned above may be helpful

    * Talk to SIG docs (SIG docs wants to try and send delegates to SIG, but for now, come to meetings)

* Question: should there even *be* a split between user and contributor docs?

    * No, but we have historically (didn’t used to have person-power to implement it)

* Onboarding docs push for 2018

    * Point people from moment one to start understanding how to contribute (design patterns, etc)

    * How do we organize, make things findable

    * How do we organize allowing people to contribute right now

        * Main site is all user docs right now

        * Starting to look at auto-import docs

        * User journeys project (above)

        * Testing right now

* Lots of people writing blog posts

    * Any way to aggregate to some of that?

    * Wow to avoid "this doesn’t work anymore"?

    * There is a Kubernetes blog, being migrated under k8s.io

        * People can contribute to the blog for ephemeral stuff

    * We guarantee that k8s.io stuff is updated

    * For blogs, put a "sell-by date"

* Request: join dashboard conversation

    * How do we link from dashboard to docs and vice-versa

* Question: where do the docs of tools like kops fit in

    * More broadly hits on Kubernetes "kernel" vs wrapper projects

    * Right now, it’s checked into repo, so it should have docs on the main site

        * Documentation is endorsing, at bit

    * (but why is kops checked into the repo as opposed to other installers)

    * Have a easy way for projects (incubator and main org) to have their own docs sub-sites (e.g. gh-pages-style for the Kubernetes sites)

### Follow-Ups

* Suggestion: Presentation on how to structure site, write docs, user studies done, etc (why are Kube docs the way they are)

    * "I want to get started fast"

    * "I want to build the next thing that gives me value" (task-driven)

    * "Something broke, where’s the reference docs" (troubleshooting)

### Projects In Development

* Customer User Journeys ([project](https://github.com/kubernetes/website/projects/5))

* Revamp of [docs landing page](https://kubernetes.io/docs/home/)

* Building out a glossary ([project](https://github.com/kubernetes/website/issues))

* More doc sprints ([like this one](https://docs.google.com/document/d/1Ar4ploza6zA1JF3YO4e0lzq1RRBWWx8cAZtRQMMEdsw/edit))

### Continue participating!

* Content exists under [Kubernetes/website](https://github.com/kubernetes/website) (feel free to fix an issue by [creating a PR](https://kubernetes.io/docs/home/contribute/create-pull-request/))

* Join the [kubernetes SIG Docs group](https://groups.google.com/forum/#!forum/kubernetes-sig-docs) 

* Attend the next Kubernetes docs community meeting (Tuesdays @ 10:30am, PST)

* And join the kubernetes sig-docs slack channel ([kubernetes.slack.com](http://kubernetes.slack.com/) #sig-docs)

