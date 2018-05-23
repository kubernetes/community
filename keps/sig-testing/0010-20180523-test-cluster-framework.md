---
kep-number: 10
title: Test Cluster Framework
authors:
  - "@timothysc"  # Timothy St. Clair
  - "@marun"      # Maru Newby
  - "Liz Frost"
  - "@frankgreco" # Frank B Greco Jr,
  - "@hoegaarden" # Hannes Hoerl
  - "@totherme"   # Gareth Smith
owning-sig: sig-testing
participating-sigs:
reviewers:
  - TBD
approvers:
  - TBD
editor: TBD
creation-date: 2018-05-23
last-updated: 2018-05-23
status: provisional
see-also:
replaces:
superseded-by:
---

# Test Cluster Framework

## Table of Contents

* [Test Cluster Framework](#test-cluster-framework)
   * [Table of Contents](#table-of-contents)
   * [Summary](#summary)
   * [Motivation](#motivation)
      * [Goals](#goals)
      * [Non-Goals](#non-goals)
   * [Proposal](#proposal)
      * [User Stories](#user-stories)
      * [Implementation Details/Notes/Constraints](#implementation-detailsnotesconstraints)
      * [Risks and Mitigations](#risks-and-mitigations)
   * [Graduation Criteria](#graduation-criteria)
   * [Implementation History](#implementation-history)

<!--
[Tools for generating]: https://github.com/ekalinin/github-markdown-toc
-->

## Summary

Alice, an experienced CLI Dev (see Key Personas in [the linked
document][the_doc]), finds that she can’t easily onboard new contributors like
Bob (see Key Personas in [the linked document][the_doc]), in part because her
CLI and tests are too tightly coupled with heavyweight infrastructure in
`k8s.io/kubernetes`. To solve this problem, Alice sponsors the development of a
new lightweight testing infrastructure outside `k8s.io/kubernetes`.

Meanwhile, Carol, an API/Controller dev (see Key Personas in [the linked
document][the_doc]), is frustrated that there are already too many ways to spin
up test clusters.

There is a clear tension between Alice and Carol’s frustrations. Left to their
own devices, Alice is likely to contribute to the proliferation of test
infrastructure that frustrates Carol, and Carol may find herself advocating
that Alice be prevented from acting to resolve her own frustration with the
status quo.

Most of the content of this KEP has already [been worked on and fleshed
out][the_doc] mostly by people from sig-testing, more specifically people
working on the [testing-commons subproject][testing_commons_notes].

[testing_commons_notes]: https://docs.google.com/document/d/1TOC8vnmlkWw6HRNHoe5xSv5-qv7LelX6XK3UVCHuwb0/edit#
[the_doc]: https://docs.google.com/document/d/13bMjmWpsdkgbY-JayrcU-e_QNwRJCP-rHjtqdeeoQHo/edit#

## Motivation

With this proposal we would like to address these three points of friction:
- For Carol, we introduce an abstraction which is intended to unify existing
  and future methods of spinning up test clusters.
- For Alice’s frustration, we build a lightweight testing framework with no
  dependency on `k8s.io/kubernetes`
- For the tension between Alice and Carol, we ensure that Alice’s framework is
  fully accessible through Carol’s abstraction.


<!--
TODO: Add note that we did interviews and talked to different people, noted in
      the "Validation" section of the personas?

[experience reports]: https://github.com/golang/go/wiki/ExperienceReports
-->

### Goals

1. To design an abstraction of “standing up a kubernetes cluster” which is:
    - expressive enough that most or all e2e tests ought to be able to use it,
      rather than standing things up directly
    - abstract enough that it could in principle be wrapped around many existing
      methods of standing up clusters
    - equally easy to use interactively and in CI
2. To build a lightweight testing framework for CLI use which
    - is not dependant on `k8s.io/kubernetes`
    - is easy for newcomers to use
    - makes use of the abstraction above
3. To build a heavier weight testing framework, for API/Controller use which
    - is fully functional
    - makes use of the abstraction above


### Non-Goals

- To wrap any existing “ways of standing up a cluster” in our abstraction
   - Of course, wrapping existing technology may well be the most efficient way
     to meet goals 2 and 3 (building CLI and API/Controller testing frameworks). 
- To port large numbers of existing tests to use our abstraction.

## Proposal

### User Stories

- [A lightweight framework](https://docs.google.com/document/d/13bMjmWpsdkgbY-JayrcU-e_QNwRJCP-rHjtqdeeoQHo/edit#heading=h.pwzs2vde6z0h)
- [A heavier weight framework](https://docs.google.com/document/d/13bMjmWpsdkgbY-JayrcU-e_QNwRJCP-rHjtqdeeoQHo/edit#heading=h.jk3ttegjtexg)
- [Bridging the two frameworks](https://docs.google.com/document/d/13bMjmWpsdkgbY-JayrcU-e_QNwRJCP-rHjtqdeeoQHo/edit#heading=h.g6pbp8y2hrj)
- [Newcomer-friendly lightweight bootstrapping](https://docs.google.com/document/d/13bMjmWpsdkgbY-JayrcU-e_QNwRJCP-rHjtqdeeoQHo/edit#heading=h.19lu9pyt334a)
- [Proper debugging support](https://docs.google.com/document/d/13bMjmWpsdkgbY-JayrcU-e_QNwRJCP-rHjtqdeeoQHo/edit#heading=h.mzdpiav3x5tz)

There are more [user stories][additional_userstories] and
[personas][additional_personas] which are probably out of scope for this KEP
but paint a potential future direction we could imagine for this test cluster
framework.

[additional_personas]: https://docs.google.com/document/d/13bMjmWpsdkgbY-JayrcU-e_QNwRJCP-rHjtqdeeoQHo/edit#heading=h.7k68eja3mgiu
[additional_userstories]: https://docs.google.com/document/d/13bMjmWpsdkgbY-JayrcU-e_QNwRJCP-rHjtqdeeoQHo/edit#heading=h.egbobak5q21m

### Implementation Details/Notes/Constraints

TBD

### Risks and Mitigations

- [Increasing the number of ways to stand up clusters for testing/investigation](https://docs.google.com/document/d/13bMjmWpsdkgbY-JayrcU-e_QNwRJCP-rHjtqdeeoQHo/edit#heading=h.89dtjdkjrit0)
- [Over-specialising our abstraction](https://docs.google.com/document/d/13bMjmWpsdkgbY-JayrcU-e_QNwRJCP-rHjtqdeeoQHo/edit#heading=h.5gypeylotxzj)

## Graduation Criteria

TBD

## Implementation History

<!--
Major milestones in the life cycle of a KEP should be tracked in `Implementation History`.
Major milestones might include

- the `Summary` and `Motivation` sections being merged signaling SIG acceptance
- the `Proposal` section being merged signaling agreement on a proposed design
- the date implementation started
- the first Kubernetes release where an initial version of the KEP was available
- the version of Kubernetes where the KEP graduated to general availability
- when the KEP was retired or superseded
-->
