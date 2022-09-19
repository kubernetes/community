# Writing Good Conformance Tests for Kubernetes #

The overarching goal of conformance tests is to exercise core
Kubernetes functionality in (as much as possible) a provider-agnostic
way.

Conformance tests are promoted from existing e2e tests which have
been proven to be stable and "non-flaky".

Please read the ["Writing good e2e tests"](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-testing/writing-good-e2e-tests.md) guide as a supplement to this document.

### Requirements  ###

In addition to the "Writing good e2e tests" guide, there are specific
requirements that all Conformance tests must adhere to. You can find
the definitive list of requirements in the [SIG-Arch community
docs](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/conformance-tests.md#conformance-test-requirements).

Generally speaking, tests should not target internal behaviors, but
instead target exposed system behaviors. Care must also be taken to ensure that
the tests do not implicitly test features that are not already subject to
conformance. If a feature is needed for a test, make sure that an existing test
already covers that feature directly.

### Existing Tests ###

When looking for an existing e2e test which exercises a behaviour
which is useful for Conformance, you may find that the test does not
meet all of the [requirements](#Requirements). In these cases, you
might either modify the test to meet requirements, but sometimes doing
so will mean losing some of the intent of the test.

Instead, a better approach may be to duplicate the test, and make your
modifications, thus preserving the original test as-is.

Any existing e2e test should have gone through a standard review
process by a respective [SIG](https://github.com/kubernetes/community/blob/master/sig-list.md)
in order to have been merged in, which means that the test was most likely
reviewed by domain-specific experts. This can be useful to note when
modifying tests to better suit conformance concerns, or meet
[requirements](#Requirements). That said, existings tests are of varying levels
of quality. It is important to verify that the test actually does validate the
expected behavior, as described in any API documentation, and as specified in
the test description you add for promotion.

### New Tests ###

New tests must go through the standard process for e2e tests even when
intended for Conformance from the onset. This means working with
SIG-specific reviewers in order to get your PR accepted. After which,
the test must be in the test cycle for 2 weeks, to prove its
stability. If the test is not slow or flaky, then you may begin the
promotion process by following the steps detailed in ["Promoting
Tests to
Conformance"](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/conformance-tests.md#promoting-tests-to-conformance).
