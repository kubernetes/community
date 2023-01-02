# Writing good e2e tests for Kubernetes #

## Patterns and Anti-Patterns ##

### Goals of e2e tests ###

Beyond the obvious goal of providing end-to-end system test coverage,
there are a few less obvious goals that you should bear in mind when
designing, writing and debugging your end-to-end tests.  In
particular, "flaky" tests, which pass most of the time but fail
intermittently for difficult-to-diagnose reasons are extremely costly
in terms of blurring our regression signals and slowing down our
automated merge velocity.  Up-front time and effort designing your test
to be reliable is very well spent.  Bear in mind that we have hundreds
of tests, each running in dozens of different environments, and if any
test in any test environment fails, we have to assume that we
potentially have some sort of regression. So if a significant number
of tests fail even only 1% of the time, basic statistics dictates that
we will almost never have a "green" regression indicator.  Stated
another way, writing a test that is only 99% reliable is just about
useless in the harsh reality of a CI environment.  In fact it's worse
than useless, because not only does it not provide a reliable
regression indicator, but it also costs a lot of subsequent debugging
time, and delayed merges.

#### Debuggability ####

If your test fails, it should provide as detailed as possible reasons for the
failure in its failure message. The failure message is the string that gets
passed (directly or indirectly) to `ginkgo.Fail(f)`. That text is what gets
shown in the overview of failed tests for a Prow job and what gets aggregated
by https://go.k8s.io/triage.

A good failure message:
- identifies the test failure
- has enough details to provide some initial understanding of what went wrong

It's okay for it to contain information that changes during each test
run. Aggregation [simplifies the failure message with regular
expressions](https://github.com/kubernetes/test-infra/blob/d56bc333ae8acf176887a3249f750e7a8e0377f0/triage/summarize/text.go#L39-L69)
before looking for similar failures.

"Timeout" is not a useful error
message. "Timed out after 60 seconds waiting for pod xxx to enter
running state, still in pending state" is much more useful to someone
trying to figure out why your test failed and what to do about it.
Specifically,
[assertion](https://onsi.github.io/gomega/#making-assertions) code
like the following generates rather useless errors:

```
Expect(err).NotTo(HaveOccurred())
```

Errors returned by client-go can be very detailed. A better way to test for
errors is with `framework.ExpectNoError` because it logs the full error and
then includes only the shorter `err.Error()` in the failure message. To make
that failure message more informative,
[annotate](https://onsi.github.io/gomega/#annotating-assertions) your assertion with something like this:

```
framework.ExpectNoError(err, "tried creating %d foobars, only created %d", foobarsReqd, foobarsCreated)
```

On the other hand, overly verbose logging, particularly of non-error conditions, can make
it unnecessarily difficult to figure out whether a test failed and if
so why?  So don't log lots of irrelevant stuff either.

Except for this special case, using Gomega assertions directly is
encouraged. They are more versatile than the (few) wrappers that were added to
the E2E framework. Use assertions that match the check in the test. Using Go
code to evaluate some condition and then checking the result often isn't
informative. For example this check should be avoided:

```
gomega.Expect(strings.Contains(actualStr, expectedSubStr)).To(gomega.Equal(true))
```

Better pass both values to Gomega, which will automatically include them in the
failure message. Add an annotation that explains what the assertion is about:

```
gomega.Expect(actualStr).To(gomega.ContainSubstring("xyz"), "checking log output")
```

This produces the following failure message:
```
  [FAILED] checking log output
  Expected
      <string>: hello world
  to contain substring
      <string>: xyz
```

If there is no suitable Gomega assertion, call `ginkgo.Failf` directly:
```
import "github.com/onsi/gomega/format"

ok := someCustomCheck(abc)
if !ok {
    ginkgo.Failf("check xyz failed for object:\n%s", format.Object(abc))
}
```

[Comparing a boolean](https://github.com/kubernetes/kubernetes/issues/105678)
like this against `true` or `false` with `gomega.Equal` or
`framework.ExpectEqual` is not useful because dumping the actual and expected
value just distracts from the underlying failure reason.

Dumping structs with `format.Object` is recommended. Starting with Kubernetes
1.26, `format.Object` will pretty-print Kubernetes API objects or structs [as
YAML and omit unset
fields](https://github.com/kubernetes/kubernetes/pull/113384), which is more
readable than other alternatives like `fmt.Sprintf("%+v")`.

    import (
        "fmt"
        "k8s.io/api/core/v1"
        "k8s.io/kubernetes/test/utils/format"
    )
    
    var pod v1.Pod
    fmt.Printf("Printf: %+v\n\n", pod)
    fmt.Printf("format.Object:\n%s", format.Object(pod, 1 /* indent one level */))
    
    =>
    
    Printf: {TypeMeta:{Kind: APIVersion:} ObjectMeta:{Name: GenerateName: Namespace: SelfLink: UID: ResourceVersion: Generation:0 CreationTimestamp:0001-01-01 00:00:00 +0000 UTC DeletionTimestamp:<nil> DeletionGracePeriodSeconds:<nil> Labels:map[] Annotations:map[] OwnerReferences:[] Finalizers:[] ManagedFields:[]} Spec:{Volumes:[] InitContainers:[] Containers:[] EphemeralContainers:[] RestartPolicy: TerminationGracePeriodSeconds:<nil> ActiveDeadlineSeconds:<nil> DNSPolicy: NodeSelector:map[] ServiceAccountName: DeprecatedServiceAccount: AutomountServiceAccountToken:<nil> NodeName: HostNetwork:false HostPID:false HostIPC:false ShareProcessNamespace:<nil> SecurityContext:nil ImagePullSecrets:[] Hostname: Subdomain: Affinity:nil SchedulerName: Tolerations:[] HostAliases:[] PriorityClassName: Priority:<nil> DNSConfig:nil ReadinessGates:[] RuntimeClassName:<nil> EnableServiceLinks:<nil> PreemptionPolicy:<nil> Overhead:map[] TopologySpreadConstraints:[] SetHostnameAsFQDN:<nil> OS:nil HostUsers:<nil> SchedulingGates:[] ResourceClaims:[]} Status:{Phase: Conditions:[] Message: Reason: NominatedNodeName: HostIP: PodIP: PodIPs:[] StartTime:<nil> InitContainerStatuses:[] ContainerStatuses:[] QOSClass: EphemeralContainerStatuses:[] Resize:}}

    format.Object:
        <v1.Pod>: 
            metadata:
              creationTimestamp: null
            spec:
              containers: null
            status: {}

Consider writing a [Gomega
matcher](https://onsi.github.io/gomega/#adding-your-own-matchers). It combines
the custom check and generating the failure messages and can make tests more
readable.

It is good practice to include details like the object that failed some
assertion in the failure message because then a) the information is available
when analyzing a failure that occurred in the CI and b) it only gets logged
when some assertion fails. Always dumping objects via log messages can make the
test output very large and may distract from the relevant information.

#### Recovering from test failures ####

All tests should ensure that a cluster is restored to the state that it was in
before the test ran. [`ginkgo.DeferCleanup`
](https://pkg.go.dev/github.com/onsi/ginkgo/v2#DeferCleanup) is recommended for
this because it can be called similar to `defer` directly after setting up
something. It is better than `defer` because Ginkgo will show additional
details about which cleanup code is running and (if possible) handle timeouts
for that code (see next section). Is is better than `ginkgo.AfterEach` because
it is not necessary to define additional variables and because
`ginkgo.DeferCleanup` executes code in the more useful last-in-first-out order,
i.e. things that get set up first get removed last.

Objects created in the test namespace do not need to be deleted because
deleting the namespace will also delete them. However, if deleting an object
may fail, then explicitly cleaning it up is better because then failures or
timeouts related to it will be more obvious.

In cases where the test may have removed the object, `framework.IgnoreNotFound`
can be used to ignore the "not found" error:
```
podClient := f.ClientSet.CoreV1().Pods(f.Namespace.Name)
pod, err := podClient.Create(ctx, testPod, metav1.CreateOptions{})
framework.ExpectNoError(err, "create test pod")
ginkgo.DeferCleanup(framework.IgnoreNotFound(podClient.Delete), pod.Name, metav1.DeleteOptions{})
```

#### Interrupting tests ####

When aborting a manual `gingko ./test/e2e` invocation with CTRL-C or a signal,
the currently running test(s) should stop immediately. This gets achieved by
accepting a `ctx context.Context` as first parameter in the Ginkgo callback
function and then passing that context through to all code that might
block. When Ginkgo notices that it needs to shut down, it will cancel that
context and all code trying to use it will immediately return with a `context
canceled` error. Cleanup callbacks get a new context which will time out
eventually to ensure that tests don't get stuck. For a detailed description,
see https://onsi.github.io/ginkgo/#interrupting-aborting-and-timing-out-suites.
Most of the E2E tests [were update to use the Ginkgo
context](https://github.com/kubernetes/kubernetes/pull/112923) at the start of
the 1.27 development cycle.

There are some gotchas:

- Don't use the `ctx` passed into `ginkgo.It` in a `ginkgo.DeferCleanup`
  callback because the context will be canceled when the cleanup code
  runs. This is wrong:

        ginkgo.It("something", func(ctx context.Context) {
              ...
              ginkgo.DeferCleanup(func() {
                  // do something with ctx
              })
        })

  Instead, register a function which accepts a new context:

         ginkgo.DeferCleanup(func(ctx context.Context) {
             // do something with the new ctx
         })

  Anonymous functions can be avoided by passing some existing function and its
  parameters directly to `ginkgo.DeferCleanup`. Again, beware to *not* pass the
  wrong `ctx`. This is wrong:

        ginkgo.It("something", func(ctx context.Context) {
              ...
              ginkgo.DeferCleanup(myDeleteFunc, ctx, objName)
        })

  Instead, just pass the other parameters and let `ginkgo.DeferCleanup`
  add a new context:

        ginkgo.DeferCleanup(myDeleteFunc, objName)

- When starting some background goroutine in a `ginkgo.BeforeEach` callback,
  use `context.WithCancel(context.Background())`. The context passed into the
  callback will get canceled when the callback returns, which would cause the
  background goroutine to stop before the test runs. This works:

        backgroundCtx, cancel := context.WithCancel(context.Background())
        ginkgo.DeferCleanup(cancel)
        _, controller = cache.NewInformer( ... )
        go controller.Run(backgroundCtx.Done())

- When adding a timeout to the context for one particular operation,
  beware of overwriting the `ctx` variable. This code here applies
  the timeout to the next call and everything else that follows:

        ctx, cancel := context.WithTimeout(ctx, 5 * time.Second)
        defer cancel()
        someOperation(ctx)
        ...
        anotherOperation(ctx)

  Better use some other variable name:

        timeoutCtx, cancel := context.WithTimeout(ctx, 5 * time.Second)
        defer cancel()
        someOperation(timeoutCtx)

  When the intention is to set a timeout for the entire callback, use
  [`ginkgo.NodeTimeout`](https://pkg.go.dev/github.com/onsi/ginkgo/v2#NodeTimeout):

        ginkgo.It("something", ginkgo.NodeTimeout(30 * time.Second), func(ctx context.Context) {
        })

  There is also a `ginkgo.SpecTimeout`, but that then also includes the time
  taken for `BeforeEach`, `AfterEach` and `DeferCleanup` callbacks.

#### Polling and timeouts ####

When waiting for something to happen, use a reasonable timeout. Without it, a
test might keep running until the entire test suite gets killed by the
CI. Beware that the CI under load may take a lot longer to complete some
operation compared to running the same test locally. On the other hand, a too
long timeout can be annoying when trying to debug tests locally.

The framework provides some [common
timeouts](https://github.com/kubernetes/kubernetes/blob/eba98af1d8b19b120e39f3/test/e2e/framework/timeouts.go#L44-L109)
through the [framework
instance](https://github.com/kubernetes/kubernetes/blob/1e84987baccbccf929eba98af1d8b19b120e39f3/test/e2e/framework/framework.go#L122-L123).
When writing a test, check whether one of those fits before defining a custom
timeout in the test.

Good code that waits for something to happen meets the following criteria:
- accepts a context for test timeouts
- informative during interactive use (i.e. intermediate reports, either
  periodically or on demand)
- little to no output during a CI run except when it fails
- full explanation when it fails: when it observes some state and then
  encounters errors reading the state, then dumping both the latest
  observed state and the latest error is useful
- extension mechanism for writing custom checks
- early abort when condition cannot be reached anymore

[`gomega.Eventually`](https://pkg.go.dev/github.com/onsi/gomega#Eventually)
satisfies all of these criteria and therefore is recommended, but not required.
In https://github.com/kubernetes/kubernetes/pull/113298,
[test/e2e/framework/pods/wait.go](https://github.com/kubernetes/kubernetes/blob/222f65506252354da012c2e9d5457a6944a4e681/test/e2e/framework/pod/wait.go)
and the framework were modified to use gomega. Typically, `Eventually` is
passed a function which gets an object or lists several of them, then `Should`
checks against the expected result. Errors and retries specific to Kubernetes
are handled by [wrapping client-go
functions](https://github.com/kubernetes/kubernetes/blob/master/test/e2e/framework/get.go).

Using gomega assertions in helper packages is problematic for two reasons:
- The stacktrace associated with the failure starts with the helper unless
  extra care is take to pass in a stack offset.
- Additional explanations for a potential failure must be prepared beforehand
  and passed in.

The E2E framework therefore uses a different approach:
- [`framework.Gomega()`](https://github.com/kubernetes/kubernetes/blob/222f65506252354da012c2e9d5457a6944a4e681/test/e2e/framework/expect.go#L80-L101)
  offers similar functions as the `gomega` package, except that they return a
  normal error instead of failing the test.
- That error gets wrapped with `fmt.Errorf("<explanation>: %w)` to
  add additional information, just as in normal Go code.
- Wrapping the error (`%w` instead of `%v`) is important because then
  `framework.ExpectNoError` directly uses the error message as failure without
  additional boiler plate text. It also is able to log the stacktrace where
  the error occurred and not just where it was finally treated as a test
  failure.

Some tips for writing and debugging long-running tests:

- Use `ginkgo.By` to record individual steps. Ginkgo will use that information
  when describing where a test timed out.

- Invoke the `ginkgo` CLI with `--poll-progress-after=30s` or some other
  suitable duration to [be informed
  early](https://onsi.github.io/ginkgo/#getting-visibility-into-long-running-specs)
  why a test doesn't complete and where it is stuck. A SIGINFO or SIGUSR1
  signal can be sent to the CLI and/or e2e.test processes to trigger an
  immediate progress report.

- Use [`gomega.Eventually`](https://pkg.go.dev/github.com/onsi/gomega#Eventually)
  to wait for some condition. When it times out or
  gets stuck, the last failed assertion will be included in the report
  automatically. A good way to invoke it is:

        gomega.Eventually(ctx, func(ctx context.Context) (book Book, err error) {
            // Retrieve book from API server and return it.
            ...
         }).WithPolling(5 * time.Second).WithTimeout(30 * time.Second).
         Should(gomega.HaveField("Author.DOB.Year()", BeNumerically("<", 1900)))

  Avoid testing for some condition inside the callback and returning a boolean
  because then failure messages are not informative (see above). See
  https://github.com/kubernetes/kubernetes/pull/114640 for an example where
  [gomega/gcustom](https://pkg.go.dev/github.com/onsi/gomega@v1.27.2/gcustom)
  was used to write assertions.

  Some of the E2E framework sub-packages have helper functions that wait for
  certain domain-specific conditions. Currently most of these functions don't
  follow best practices (not using gomega.Eventually, error messages not very
  informative). [Work is
  planned](https://github.com/kubernetes/kubernetes/issues/106575) in that
  area, so beware that these APIs may
  change at some point.

- Use `gomega.Consistently` to ensure that some condition is true
  for a while. As with `gomega.Eventually`, make assertions about the
  value instead of checking the value with Go code and then asserting
  that the code returns true.

- Both `gomega.Consistently` and `gomega.Eventually` can be aborted early via
  [`gomega.StopPolling`](https://onsi.github.io/gomega/#bailing-out-early---polling-functions).

- Avoid polling with functions that don't take a context (`wait.Poll`,
  `wait.PollImmediate`, `wait.Until`, ...) and replace with their counterparts
  that do (`wait.PollWithContext`, `wait.PollImmediateWithContext`,
  `wait.UntilWithContext`, ...) or even better, with `gomega.Eventually`.


#### Ability to run in non-dedicated test clusters ####

To reduce end-to-end delay and improve resource utilization when
running e2e tests, we try, where possible, to run large numbers of
tests in parallel against the same test cluster.  This means that:

1. you should avoid making any assumption (implicit or explicit) that
your test is the only thing running against the cluster.  For example,
making the assumption that your test can run a pod on every node in a
cluster is not a safe assumption, as some other tests, running at the
same time as yours, might have saturated one or more nodes in the
cluster.  Similarly, running a pod in the system namespace, and
assuming that will increase the count of pods in the system
namespace by one is not safe, as some other test might be creating or
deleting pods in the system namespace at the same time as your test.
If you do legitimately need to write a test like that, make sure to
label it ["\[Serial\]"](e2e-tests.md#kinds-of-tests) so that it's easy
to identify, and not run in parallel with any other tests.
1. You should avoid doing things to the cluster that make it difficult
for other tests to reliably do what they're trying to do, at the same
time.  For example, rebooting nodes, disconnecting network interfaces,
or upgrading cluster software as part of your test is likely to
violate the assumptions that other tests might have made about a
reasonably stable cluster environment.  If you need to write such
tests, please label them as
["\[Disruptive\]"](e2e-tests.md#kinds-of-tests) so that it's easy to
identify them, and not run them in parallel with other tests.
1. You should avoid making assumptions about the Kubernetes API that
are not part of the API specification, as your tests will break as
soon as these assumptions become invalid.  For example, relying on
specific Events, Event reasons or Event messages will make your tests
very brittle.

#### Speed of execution ####

We have hundreds of e2e tests, some of which we run in serial, one
after the other, in some cases.  If each test takes just a few minutes
to run, that very quickly adds up to many, many hours of total
execution time.  We try to keep such total execution time down to a
few tens of minutes at most.  Therefore, try (very hard) to keep the
execution time of your individual tests below 2 minutes, ideally
shorter than that.  Concretely, adding inappropriately long 'sleep'
statements or other gratuitous waits to tests is a killer.  If under
normal circumstances your pod enters the running state within 10
seconds, and 99.9% of the time within 30 seconds, it would be
gratuitous to wait 5 minutes for this to happen.  Rather just fail
after 30 seconds, with a clear error message as to why your test
failed ("e.g. Pod x failed to become ready after 30 seconds, it
usually takes 10 seconds").  If you do have a truly legitimate reason
for waiting longer than that, or writing a test which takes longer
than 2 minutes to run, comment very clearly in the code why this is
necessary, and label the test as
["\[Slow\]"](e2e-tests.md#kinds-of-tests), so that it's easy to
identify and avoid in test runs that are required to complete
timeously (for example those that are run against every code
submission before it is allowed to be merged).
Note that completing within, say, 2 minutes only when the test
passes is not generally good enough.  Your test should also fail in a
reasonable time.  We have seen tests that, for example, wait up to 10
minutes for each of several pods to become ready.  Under good
conditions these tests might pass within a few seconds, but if the
pods never become ready (e.g. due to a system regression) they take a
very long time to fail and typically cause the entire test run to time
out, so that no results are produced.  Again, this is a lot less
useful than a test that fails reliably within a minute or two when the
system is not working correctly.

#### Resilience to relatively rare, temporary infrastructure glitches or delays ####

Remember that your test will be run many thousands of
times, at different times of day and night, probably on different
cloud providers, under different load conditions.  And often the
underlying state of these systems is stored in eventually consistent
data stores.  So, for example, if a resource creation request is
theoretically asynchronous, even if you observe it to be practically
synchronous most of the time, write your test to assume that it's
asynchronous (e.g. make the "create" call, and poll or watch the
resource until it's in the correct state before proceeding).
Similarly, don't assume that API endpoints are 100% available.
They're not.  Under high load conditions, API calls might temporarily
fail or time-out. In such cases it's appropriate to back off and retry
a few times before failing your test completely (in which case make
the error message very clear about what happened, e.g. "Retried
http://... 3 times - all failed with xxx".  Use the standard
retry mechanisms provided in the libraries detailed below.

### Some concrete tools at your disposal ###

Obviously most of the above goals apply to many tests, not just yours.
So we've developed a set of reusable test infrastructure, libraries
and best practices to help you to do the right thing, or at least do
the same thing as other tests, so that if that turns out to be the
wrong thing, it can be fixed in one place, not hundreds, to be the
right thing.

Here are a few pointers:

+ [E2e Framework](https://git.k8s.io/kubernetes/test/e2e/framework/framework.go):
   Familiarise yourself with this test framework and how to use it.
   Amongst others, it automatically creates uniquely named namespaces
   within which your tests can run to avoid name clashes, and reliably
   automates cleaning up the mess after your test has completed (it
   just deletes everything in the namespace).  This helps to ensure
   that tests do not leak resources. Note that deleting a namespace
   (and by implication everything in it) is currently an expensive
   operation.  So the fewer resources you create, the less cleaning up
   the framework needs to do, and the faster your test (and other
   tests running concurrently with yours) will complete. Your tests
   should always use this framework.  Trying other home-grown
   approaches to avoiding name clashes and resource leaks has proven
   to be a very bad idea.
+ [E2e utils library](https://git.k8s.io/kubernetes/test/e2e/framework/util.go):
   This handy library provides tons of reusable code for a host of
   commonly needed test functionality, including waiting for resources
   to enter specified states, safely and consistently retrying failed
   operations, usefully reporting errors, and much more.  Make sure
   that you're familiar with what's available there, and use it.
   Likewise, if you come across a generally useful mechanism that's
   not yet implemented there, add it so that others can benefit from
   your brilliance.  In particular pay attention to the variety of
   timeout and retry related constants at the top of that file. Always
   try to reuse these constants rather than try to dream up your own
   values.  Even if the values there are not precisely what you would
   like to use (timeout periods, retry counts etc), the benefit of
   having them be consistent and centrally configurable across our
   entire test suite typically outweighs your personal preferences.
+ **Follow the examples of stable, well-written tests:** Some of our
   existing end-to-end tests are better written and more reliable than
   others.  A few examples of well-written tests include:
   [Replication Controllers](https://git.k8s.io/kubernetes/test/e2e/apps/rc.go),
   [Services](https://git.k8s.io/kubernetes/test/e2e/network/service.go),
   [Reboot](https://git.k8s.io/kubernetes/test/e2e/cloud/gcp/reboot.go).
+ [Ginkgo Test Framework](https://github.com/onsi/ginkgo): This is the
   test library and runner upon which our e2e tests are built.  Before
   you write or refactor a test, read the docs and make sure that you
   understand how it works.  In particular be aware that every test is
   uniquely identified and described (e.g. in test reports) by the
   concatenation of its `Describe` clause and nested `It` clauses.
   So for example `Describe("Pods",...).... It(""should be scheduled
   with cpu and memory limits")` produces a sane test identifier and
   descriptor `Pods should be scheduled with cpu and memory limits`,
   which makes it clear what's being tested, and hence what's not
   working if it fails.  Other good examples include:

```
   CAdvisor should be healthy on every node
```

and

```
   Daemon set should run and stop complex daemon
```

   On the contrary
(these are real examples), the following are less good test
descriptors:

```
   KubeProxy should test kube-proxy
```

and

```
Nodes [Disruptive] Network when a node becomes unreachable
[replication controller] recreates pods scheduled on the
unreachable node AND allows scheduling of pods on a node after
it rejoins the cluster
```

An improvement might be

```
Unreachable nodes are evacuated and then repopulated upon rejoining [Disruptive]
```

Note that opening issues for specific better tooling is welcome, and
code implementing that tooling is even more welcome :-).

### Resource usage ###
When writing tests, resources used in the tests should be chosen specifically and sanely quanified.

Therefore it is important to use resources that are:
1. appropriate to the test
2. inexpensive (containing little or no overhead)
3. create just as many an needed, no more
4. should be cleaned up at the end of the tests's run

For example:
1. only using the resource type which is appropriate to the `test/e2e/<AREA>/<SPECIFIC_AREA>`
2. the resource type `ConfigMap` is inexpensive, common, and stateless. It should be used for such things as fetching created resources
3. although the clusters used for testing are generally beefy, an overamount of resources should not be created as it is unnecessary
4. using afterEach, make sure that your test destroys any resources leftover from the test 

### Logging ###
When writing tests, it's often useful to have logging of events which take place (at least during the development of the given test).
To log, import the module `framework`. Once included, in the test, you will be able to call `framework.Logf`.

#### Example usage ####
1. Printing a single string wrapped in a string
```golang
testvar := "Hello World"
framework.Logf("Logf says: '%v'", testvar)
```

2. Printing a string and a number wrapping in a string
```golang
testvar1 := "a string"
testvar2 := 1
framework.Logf("testvar1: %v; testvar2: %v", testvar1, testvar2)
```

For more information, please refer to [the framework documentation](https://godoc.org/k8s.io/kubernetes/test/e2e/framework#Logf)

