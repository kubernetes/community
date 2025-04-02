# SIG Node CI testing subgroup charters

- **Created:**  Jul 24, 2020
- **Last Updated:** Mar 27, 2025

## Overview

SIG Node CI Group is proposed and formed by volunteers from the [SIG Node](charter.md) community.
SIG Node owns a broad range of [code, binaries, and services](charter.md#code-binaries-and-services)
which are fundamental to the success of Kubernetes.
To ensure high quality of these components, SIG Node community has developed a large number of
[end-to-end (e2e) test cases](https://testgrid.k8s.io/sig-node) over the past several years.
With new PR requests merged every day, OS distros being updated,
it has become more and more challenging to maintain these tests and to keep them up to date with new changes.
In order to maintain tests and keep up with the incoming bugs triage, the CI subproject was formed.
The progress of this group can be viewed in the [meeting notes](https://docs.google.com/document/d/1fb-ugvgdSVIkkuJ388_nhp2pBTy_4HEVg5848Xy7n5U/edit#heading=h.badwgoqn6j9e).

## Goal

The goal of this group is to keep a high quality bar for SIG Node and to enable continuous integration of node features of Kubernetes.
More specifically, this group drives all aspects of [SIG Node e2e tests](https://testgrid.k8s.io/sig-node),
and aims to achieve the following goals:

- Maintaining existing node e2e tests healthiness
  - Triage and fix failing tests (especially release-blocking tests) in a timely manner
  - Remove deprecated tests
  - Enhanced test failure troubleshooting and diagnostics tooling and documentation

- Improving test code quality
  - Identify and reduce flaky behavior in tests
  - Review new e2e test code

- Improving test coverage
  - Identify areas missing test coverage, and work with owners to add e2e tests

- OS Image and dependencies support
  - Make sure the representative set of OS images are being tested
  - Make sure the representative set of container runtime versions are being tested

- Test resource utilization and cost optimization
  - Increase ROI of SIG Node tests
  - Help spread tests for different clouds for better usage of cloud credits

- Make sure timely response for critical bugs
  - Triage bugs regularly
  - Identify critical issue and ensure timely follow up
  - Review high priority bugs before each release

## Execution Plan

- Regular meet-ups for discussion and triage
  - Triage test failures
  - Review test related Pull requests and issues backlog
  - Triage bugs
  - Periodically summarize test coverage (flakes, failures etc) in sig-node meeting for broader visibility

- Image Support
  - Maintain owners for supported OS images.
  - Process/documentation on lifecycle of OS images used by SIG Node tests.  
  - Ensure a process and/or automation exists for OS image upgrading process.

- Group Sustainability
  - Identify, document release-blocking tests  
  - Grow node e2e test reviewers and approvers

## History

- Original [charter document](https://docs.google.com/document/d/1yS-XoUl6GjZdjrwxInEZVHhxxLXlTIX2CeWOARmD8tY/)
