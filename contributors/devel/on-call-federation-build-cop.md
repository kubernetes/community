# Federation Buildcop Guide and Playbook

Federation runs two classes of tests: CI and Pre-submits.

## CI

* These tests run on the HEADs of master and release branches (starting
  from Kubernetes v1.7).
* As a result, they run on code that's already merged.
* As the name suggests, they run continuously. Currently, they are
  configured to run at least once every 30 minutes.
* Federation CI tests run as periodic jobs on prow.
* CI jobs always run sequentially. In other words, no single CI job
  can have two instances of the job running at the same time.
* Latest build results can be viewed in [testgrid](https://k8s-testgrid.appspot.com/sig-multicluster)

### Configuration

Configuration steps are described in https://github.com/kubernetes/test-infra#create-a-new-job.
Federation CI e2e job names are as below:
* master branch - `ci-federation-e2e-gce` and `ci-federation-e2e-gce-serial`
* 1.8 release branch - `ci-kubernetes-e2e-gce-federation-release-1-8`
* 1.7 release branch - `ci-kubernetes-e2e-gce-federation-release-1-7`

Search for the above job names in various configuration files as below:

* Prow config: https://github.com/kubernetes/test-infra/blob/master/prow/config.yaml
* Test job/bootstrap config: https://github.com/kubernetes/test-infra/blob/master/jobs/config.json
* Test grid config: https://github.com/kubernetes/test-infra/blob/master/testgrid/config/config.yaml
* Job specific config: https://github.com/kubernetes/test-infra/tree/master/jobs/env

### Results

Results of all the federation CI tests are listed in the corresponding 
tabs on the Cluster Federation page in the testgrid.
https://k8s-testgrid.appspot.com/sig-multicluster

### Playbook

#### Triggering a new run

Please ping someone who has access to the prow project and ask
them to click the `rerun` button from, for example 
http://prow.k8s.io/?type=periodic&job=ci-federation-e2e-gce,
and execute the kubectl command.

#### Quota cleanup

Please ping someone who has access to the GCP project. Ask them to
look at the quotas and delete the leaked resources by clicking the
delete button corresponding to those leaked resources on Google Cloud
Console.


## Pre-submit

* The pre-submit test is currently configured to run on the master
  branch and any release branch that's 1.9 or newer.
* Multiple pre-submit jobs could be running in parallel(one per pr).
* Latest build results can be viewed in [testgrid](https://k8s-testgrid.appspot.com/presubmits-federation)
* We have following pre-submit jobs in federation
  * bazel-test - Runs all the bazel test targets in federation.
  * e2e-gce - Runs federation e2e tests on gce.
  * verify - Runs federation unit, integration tests and few verify scripts.

### Configuration

Configuration steps are described in https://github.com/kubernetes/test-infra#create-a-new-job.
Federation pre-submit jobs have following names.
* bazel-test - `pull-federation-bazel-test`
* verify - `pull-federation-verify`
* e2e-gce - `pull-federation-e2e-gce`

Search for the above job names in various configuration files as below:

* Prow config: https://github.com/kubernetes/test-infra/blob/master/prow/config.yaml
* Test job/bootstrap config: https://github.com/kubernetes/test-infra/blob/master/jobs/config.json
* Test grid config: https://github.com/kubernetes/test-infra/blob/master/testgrid/config/config.yaml
* Job specific config: https://github.com/kubernetes/test-infra/tree/master/jobs/env

### Results

Aggregated results are available on the Gubernator dashboard page for
the federation pre-submit tests.

https://k8s-gubernator.appspot.com/builds/kubernetes-jenkins/pr-logs/directory/pull-federation-e2e-gce

### Metrics

We track the flakiness metrics of all the pre-submit jobs and
individual tests that run against PRs in
[kubernetes/federation](https://github.com/kubernetes/federation).

* The metrics that we track are documented in https://github.com/kubernetes/test-infra/blob/master/metrics/README.md#metrics.
* Job-level metrics are available in http://storage.googleapis.com/k8s-metrics/job-flakes-latest.json.

### Playbook

#### Triggering a new run

Use the `/test` command on the PR to re-trigger the test. The exact
incantation is: `/test pull-federation-e2e-gce`

#### Quota cleanup

Please ping someone who has access to `k8s-jkns-pr-bldr-e2e-gce-fdrtn`
GCP project. Ask them to look at the quotas and delete the leaked
resources by clicking the delete button corresponding to those leaked
resources on Google Cloud Console.
