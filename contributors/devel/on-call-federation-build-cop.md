# Federation Buildcop Guide and Playbook

Federation runs two classes of tests: CI and Presubmits.

## CI

* These tests run on the HEADs of master and release branches (starting
  from Kubernetes v1.6).
* As a result, they run on code that's already merged.
* As the name suggests, they run continuously. Currently, they are
  configured to run
  [at least once every 30 minutes](https://github.com/kubernetes/test-infra/blob/22c38cfb64137086373e1b89d5e7d98766560747/prow/config.yaml#L3686).
* Federation CI tests run as
  [periodic jobs on prow](https://github.com/kubernetes/test-infra/blob/22c38cfb64137086373e1b89d5e7d98766560747/prow/config.yaml#L3686).
* CI jobs always run sequentially. In other words, no single CI job
  can have two instances of the job running at the same time.

### Configuration

Configuration steps are described in https://github.com/kubernetes/test-infra/blob/0c56d2c9d32307c0a0f8fece85ef6919389e77fd/jenkins/README.md#how-to-work-with-jenkins-jobs

The configuration of CI tests are stored in:

* Jenkins config: https://github.com/kubernetes/test-infra/blob/0c56d2c9d32307c0a0f8fece85ef6919389e77fd/jenkins/job-configs/kubernetes-jenkins/bootstrap-ci.yaml
* Test job/bootstrap config: https://github.com/kubernetes/test-infra/blob/0c56d2c9d32307c0a0f8fece85ef6919389e77fd/jobs/config.json
* Test grid config: https://github.com/kubernetes/test-infra/blob/0c56d2c9d32307c0a0f8fece85ef6919389e77fd/testgrid/config/config.yaml
* Job specific config: https://github.com/kubernetes/test-infra/tree/0c56d2c9d32307c0a0f8fece85ef6919389e77fd/jobs

### Results

Results of all the federation CI tests, including the soak tests, are
listed in the corresponding tabs on the Cluster Federation page in the
testgrid.
https://k8s-testgrid.appspot.com/sig-federation

### Playbook

#### Triggering a new run

Please ping someone who has access to the Jenkins UI/dashboard and ask
them to login and click the "Build Now" link on the Jenkins page
corresponding to the CI job you want to manually start.

#### Quota cleanup

Please ping someone who has access to the GCP project. Ask them to
look at the quotas and delete the leaked resources by clicking the
delete button corresponding to those leaked resources on Google Cloud
Console.


## Presubmit

* We only have one presubmit test, but it is configured very
  differently than the CI tests.
* The presubmit test is currently configured to run on the master
  branch and any release branch that's 1.7 or newer.
* Federation presubmit infrastructure is composed of two separate test
  jobs:
  * Deploy job: This job runs in the background and recycles federated
    clusters every time it runs. Although this job supports federation
    presubmit tests, it is configured as a CI/Soak job. More on
    configuration later. Since recycling federated clusters is an
    expensive operation, we do not want to run this often. Hence, this
    job is configured to run once every 24 hours, around midnight
    Pacific time.
  * Test job: This is the job that runs federation presubmit tests on
    every PR in the core repository, i.e.
    [kubernetes/kubernetes](https://github.com/kubernetes/kubernetes).
    These jobs can run in parallel on the PRs in the repository.

### Two-jobs setup

The deploy job runs once every 24 hours at around midnight Pacific
time. It is configured to turn up and tear down 3 federated clusters.
It starts out by downloading the latest Kubernetes release built from
[kubernetes/kubernetes](https://github.com/kubernetes/kubernetes)
master. It then tears down the existing federated clusters and turns
up new ones. As the clusters are created, their kubeconfigs are
written to a local kubeconfig file where the job runs. Once all the
clusters are successfully turned up, the local kubeconfig is then
copied to a pre-configured GCS bucket. Any existing kubeconfig in the
bucket will be overwritten.

The test job on the other hand starts by copying the latest kubeconfig
from the pre-configured GCS bucket. It uses this kubeconfig to deploy
a new federation control plane on one of the clusters in the
kubeconfig. It then joins all the clusters in the kubeconfig, including
the host cluster where federation control plane is deployed, as members
to the newly created federation control plane. The test job then runs
the federation presubmit tests on this control plane and tears down the
control plane in the end.

Since federated clusters are recycled only once every 24 hours, all
presubmit runs in that period share the federated clusters. And since
there could be multiple presubmit tests running in parallel, each
instance of the test gets its own namespace where it deploys the
federation control plane. These federation control planes deployed in
separate namespaces are independent of each other and do not interfere
with other federation control planes in any way.

### Configuration

The two jobs are configured differently.

#### Deploy job

The deploy job is configured as a CI/Soak job in Jenkins.
Configuration steps are described in https://github.com/kubernetes/test-infra/blob/0c56d2c9d32307c0a0f8fece85ef6919389e77fd/jenkins/README.md#how-to-work-with-jenkins-jobs

The configuration of the deploy job is stored in:

* Jenkins config: https://github.com/kubernetes/test-infra/blob/0c56d2c9d32307c0a0f8fece85ef6919389e77fd/jenkins/job-configs/kubernetes-jenkins/bootstrap-ci-soak.yaml#L76
* Test job/bootstrap config: https://github.com/kubernetes/test-infra/blob/0c56d2c9d32307c0a0f8fece85ef6919389e77fd/jobs/config.json#L3996
* Test grid config: https://github.com/kubernetes/test-infra/blob/0c56d2c9d32307c0a0f8fece85ef6919389e77fd/testgrid/config/config.yaml#L152
* Job specific config: https://github.com/kubernetes/test-infra/blob/0c56d2c9d32307c0a0f8fece85ef6919389e77fd/jobs/ci-kubernetes-pull-gce-federation-deploy.env

#### Test job

The test job is
[configured in prow](https://github.com/kubernetes/test-infra/blob/35ceb37e999bb0589218708262634951b79dfe05/prow/config.yaml#L236),
but it runs in Jenkins mode. The configuration steps are described in
https://github.com/kubernetes/test-infra/blob/0c56d2c9d32307c0a0f8fece85ef6919389e77fd/README.md#create-a-new-job

The configuration of the test job is stored in:

* Prow config: https://github.com/kubernetes/test-infra/blob/0c56d2c9d32307c0a0f8fece85ef6919389e77fd/prow/config.yaml#L244
* Test job/bootstrap config: https://github.com/kubernetes/test-infra/blob/0c56d2c9d32307c0a0f8fece85ef6919389e77fd/jobs/config.json#L4691
* Job specific config: https://github.com/kubernetes/test-infra/blob/0c56d2c9d32307c0a0f8fece85ef6919389e77fd/jobs/pull-kubernetes-federation-e2e-gce.env

### Results

Aggregated results are available on the Gubernator dashboard page for
the federation presubmit tests.

https://k8s-gubernator.appspot.com/builds/kubernetes-jenkins/pr-logs/directory/pull-kubernetes-federation-e2e-gce

### Metrics

We track the flakiness metrics of all the presubmit jobs and
individual tests that run against PRs in
[kubernetes/kubernetes](https://github.com/kubernetes/kubernetes).

* The metrics that we track are documented in https://github.com/kubernetes/test-infra/blob/0c56d2c9d32307c0a0f8fece85ef6919389e77fd/metrics/README.md#metrics.
* Job-level metrics are available in - [http://storage.googleapis.com/k8s-metrics/job-flakes-latest.json]().
* As of this writing, federation presubmits have a [success rate of
  93.4%](http://storage.googleapis.com/k8s-metrics/job-flakes-latest.json).

### Playbook

#### Triggering a new deploy job run

Please ping someone who has access to the Jenkins UI/dashboard and ask
them to login and click the "Build Now" link on the Jenkins page
corresponding to the CI job you want to manually start.

#### Triggering a new test run

Use the `/test` command on the PR to retrigger the test. The exact
incantation is: `/test pull-kubernetes-federation-e2e-gce`

#### Quota cleanup

Please ping someone who has access to `k8s-jkns-pr-bldr-e2e-gce-fdrtn`
GCP project. Ask them to look at the quotas and delete the leaked
resources by clicking the delete button corresponding to those leaked
resources on Google Cloud Console.
