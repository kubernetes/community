# Kubernetes Testing Flakiness SLA

This document captures the expectations of the community about flakiness in
our tests and our test infrastructure. It sets out an SLA (Service Level 
Agreement) for flakiness in our tests, as well as actions that we will
take when we are out of SLA.

## Definition of "We"

Throughout the document the term _we_ is used. This is intended to refer
to the Kubernetes project as a whole, and any governance structures the
project puts in place. It is not intended to refer to any specific group
of individuals.

## Definition of a "Flake"

We'll start by the definition of a _flake_. _Flakiness_ is defined for a
complete run of one required job in the pre-submit testing infrastructure 
(e.g. pull-kubernetes-e2e-gke-gci) for a given pull request (PR). We will
not measure flakiness SLA for e2e jobs that are not required for code merge.

A pre-submit job's test result is considered to
be a flake according to two criteria:

1) it both fails and passes without any changes in the code
or environment being tested.

2) the PR in question doesn't cause the flake itself.

### Measuring flakiness
There are a number of challenges in monitoring flakiness. We expect that
the metric will be heuristic in nature and we will iterate on it over time.
Identifying all of the potential problems and ways to measure the metric
are out of the scope of the document, but some currently known challenges
are listed at the end of the document.

## Flakiness SLA
We will measure flakiness based on pull requests that are run through pre-submit
PR testing. The metric that we will monitor is: Pre-submit flakes per PR. This metric
will be calculated on a daily and weekly basis using:

Sum(Flakes not caused by the PR) / Total(PRs)

Our current SLA is that this metric will be less than 0.01
(1% of all PRs have flakes that are not caused by the PR itself).

## Activities on SLA violation
When the project is out of SLA for flakiness on either the daily or weekly metric
we will determine an appropriate actions to take to bring the project back
within SLA. For now these specific actions are to be determined. 

## Non-goals: Flakiness at HEAD
We will consider flakiness for PRs only, we will _not_ currently 
measure flakiness for continuous e2e that are run at HEAD independent of PRs.

There a few reasons for this:
   * The volume of testing on PRs is significantly higher than at HEAD, so flakes are more readily apparent.
   * The flakiness of these tests is already being measured in other places.
   * The goal of this proposal is to improve contributor experience, which is
   governed by PR behavior, rather than the comprehensive e2e suite which is
   intended to ensure project stability.

Going forward, if the e2e suite at HEAD is showing increased instability, we
may chose to update our flakiness SLA.
 
## Infrastructure, enforcement and success metrics

### Monitoring and enforcement infrastructure
SIG-Testing are currently responsible for the submit-queue infrastructure
and will be responsible for designing, implementing and deploying the
relevant monitoring and enforcement mechanisms for this proposal.

### Assessing success
Ultimately, the goal of this effort is to decrease flaky tests and
improve the contributor experience. To that end, SIG-contributor-experience
will assess and evaluate if this proposal is successful, and will refine
or eliminate this proposal as new evidence is obtained.

Success for this proposal will be measured in terms of the overall flakiness
of tests. The number of times the SLA is violated in a given time period,
and the number of PRs to fix test flakes. If this proposal is successful
all of these metrics should decrease over time.

## Approaches to measuring the flakiness metric

Currently, measuring this metric is somewhat tricky, since determining flakes
caused by PRs vs. existing flakes is somewhat challenging. We will use a variety
of heuristics to determine this, including looking at PRs that contain changes
that could not cause flakes (e.g. typos, docs). As well as looking at the
past history of test failures (e.g. if a test fails across many different PRs
it's likely a flake)

### Detecting changes in the environment (e.g. GCE, Azure, AWS, etc)
Changes in the environment are notably hard to measure or control for, but we'll do our best.

### Detecting changes in code
Changes in code are obviously quite easy to control. The github API can tell us if the commit SHA
has changed between different runs of the test result.

### Detecting if it is caused by PR or pre-existing?
Measuring whether a flake is caused by a PR or existing problems is a challenge, but will use
observation across multiple PRs to judge the probability that it is an existing problem or caused by the PR.
If the test flakes across multiple PRs it's likely the test. If it is only in a single PR it is likely
that PR.

### Using retest requests as a signal
When a user requests a re-test of a PR, it is a signal that the user believes the test to be flaky.
We will use this as a strong indication that the test suite result is flaky. If a user requests a retest of the suite,
and that test suite passes, that is a strong indication that there is a flaky test suite involved. 
