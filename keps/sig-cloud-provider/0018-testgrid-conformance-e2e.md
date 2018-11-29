---
kep-number: 0018
title: Reporting Conformance Test Results to Testgrid
authors:
  - "@andrewsykim"
owning-sig: sig-cloud-provider
participating-sigs:
  - sig-testing
  - sig-release
  - sig-aws
  - sig-azure
  - sig-gcp
  - sig-ibmcloud
  - sig-openstack
  - sig-vmware
reviewers:
  - TBD
approvers:
  - TBD
editor: TBD
creation-date: 2018-06-06
last-updated: 2018-11-16
status: implementable

---

# Reporting Conformance Test Results to Testgrid

## Table of Contents

* [Summary](#summary)
* [Motivation](#motivation)
    * [Goals](#goals)
    * [Non-Goals](#non-goals)
* [Proposal](#proposal)
    * [Implementation Details/Notes/Constraints](#implementation-detailsnotesconstraints)
    * [Risks and Mitigations](#risks-and-mitigations)
* [Graduation Criteria](#graduation-criteria)
* [Implementation History](#implementation-history)

## Summary

This is a KEP outlining the motivation behind why cloud providers should periodically upload E2E conformance test results to [Testgrid](https://github.com/kubernetes/test-infra/tree/master/testgrid) and how a cloud provider can go about doing this.

## Motivation

The primary motivation behind collecting conformance test results from various cloud providers on a regular basis is to inform sig-release of any critical bugs. It's important to collect results from various cloud providers to increase coverage and inform other cloud providers of bugs they may be impacted by.

### Goals

* All SIGs / Subprojects owners representing a cloud provider is aware of the importance of frequently uploading conformance test results to testgrid
* There is a clear and detailed process in place for any cloud provider to upload conformance test results to Testgrid.

### Non-Goals

* Test coverage - increasing test coverage is outside the scope of this KEP
* CI/CD - what CI/CD tool is used to run E2E tests and upload results is outside the scope of this KEP. It is up to the cloud provider to decide where/how tests are run.
* Cluster Provisioning - how a cluster is provisioned is outside the scope of this KEP.

## Proposal

We would like to propose that every Kubernetes cloud provider reports conformance test results for every patch version of Kubernetes at the minimum. Running conformance tests against master and on pre-release versions are highly encouraged but will not be a requirement.

### Implementation Details/Notes/Constraints

Before continuing, it is highly recommended you read the following documentation provided by sig-testing:
* [Testgrid Configuration](https://github.com/kubernetes/test-infra/tree/master/testgrid#testgrid)
* [Display Conformance Tests with Testgrid](https://github.com/kubernetes/test-infra/blob/master/testgrid/conformance/README.md)

#### How to Run E2E Conformance Tests

This KEP outlines two ways of running conformance tests, the first using [Sonobuoy](https://github.com/heptio/sonobuoy) and the second using [kubetest](https://github.com/kubernetes/test-infra/tree/master/kubetest). Though Sonobuoy is easier to setup, it does not guarantee that you will run the latest set of conformance tests. Kubetest though requiring a bit more work to setup, can ensure that you are running the latest set of conformance tests.

At this point we will assume that you have a running cluster and that `kubectl` is configured to point to that cluster with admin access.

#### Sonobuoy

You should use Sonobuoy if you would like to run the standard set of CNCF conformance tests. This may exclude any new tests added by the latest versions of Kubernetes.

##### Installing Sonobuoy

The following is mostly a [copy of the sonobuoy documentation](https://github.com/heptio/sonobuoy#download-and-run).

First install sonobuoy. The following command adds it to the `$GOBIN` environment variable which is expected to be part of your `$PATH` environment variable.
```
$ go get -u github.com/heptio/sonobuoy
```

##### Running Conformance Tests with Sonobuoy

You can then start e2e tests by simply running:
```
$ sonobuoy run
Running plugins: e2e, systemd-logs
INFO[0001] created object                                name=heptio-sonobuoy namespace= resource=namespaces
INFO[0001] created object                                name=sonobuoy-serviceaccount namespace=heptio-sonobuoy resource=serviceaccounts
INFO[0001] created object                                name=sonobuoy-serviceaccount-heptio-sonobuoy namespace= resource=clusterrolebindings
INFO[0001] created object                                name=sonobuoy-serviceaccount namespace= resource=clusterroles
INFO[0001] created object                                name=sonobuoy-config-cm namespace=heptio-sonobuoy resource=configmaps
INFO[0002] created object                                name=sonobuoy-plugins-cm namespace=heptio-sonobuoy resource=configmaps
INFO[0002] created object                                name=sonobuoy namespace=heptio-sonobuoy resource=pods
INFO[0002] created object                                name=sonobuoy-master namespace=heptio-sonobuoy resource=services
```

You can then check the status of your e2e tests by running
```
$ sonobuoy status
PLUGIN		STATUS	COUNT
e2e		running	1
systemd_logs	running	3

Sonobuoy is still running. Runs can take up to 60 minutes.
```

E2E tests can take up to an hour. Once your tests are done, you can download a snapshot of your results like so:
```
$ sonobuoy retrieve .
```

Once you have the following snapshot, extract it's contents like so
```
$ mkdir ./results; tar xzf *.tar.gz -C ./results
```

At this point you should have the log file and JUnit results from your tests:
```
results/plugins/e2e/results/e2e.log
results/plugins/e2e/results/junit_01.xml
```

#### Kubetest

You should use kubetest if you want to run the latest set of tests in upstream Kubernetes. This is highly recommended by sig-testing so that new tests can be accounted for in new releases.

##### Installing Kubetest

Install kubetest using the following command which adds it to your `$GOBIN` environment variable which is expected to be part of your `$PATH` environment variable.
```
go get -u k8s.io/test-infra/kubetest
```

##### Running Conformance Test with kubetest

Now you can run conformance test with the following:
```
cd /path/to/k8s.io/kubernetes
export KUBERNETES_CONFORMANCE_TEST=y
kubetest \
    # conformance tests aren't supposed to be aware of providers
    --provider=skeleton \
    # tell ginkgo to only run conformance tests
    --test --test_args="--ginkgo.focus=\[Conformance\]" \
    # grab the most recent CI tarball of kubernetes 1.10, including the tests
    --extract=ci/latest-1.10 \
    # directory to store junit results
    --dump=$(pwd)/_artifacts | tee ./e2e.log
```

Note that `--extract=ci/latest-1.10` indicates that we want to use the binaries/tests on the latest version of 1.10. You can use `--extract=ci/latest` to run the latest set of conformance tests from master.

Once the tests have finished (takes about an hour) you should have the log file and JUnit results from your tests:
```
e2e.log
_artifacts/junit_01.xml
```

#### How to Upload Conformance Test Results to Testgrid

##### Requesting a GCS Bucket

Testgrid requires that you store results in a publicly readable GCS bucket. If for whatever reason you cannot set up a GCS bucket, please contact @BenTheElder or more generally the [gke-kubernetes-engprod](mailto:gke-kubernetes-engprod@google.com) team to arrange for a Google [GKE](https://cloud.google.com/kubernetes-engine/) EngProd provided / maintained bucket for hosting your results.

##### Authenticating to your Testgrid Bucket

Assuming that you have a publicly readable bucket provided by the GKE team, you should have been provided a service account JSON file which you can use with [gcloud](https://cloud.google.com/sdk/downloads) to authenticate with your GCS bucket.

```
$ gcloud auth activate-service-account --key-file /path/to/k8s-conformance-serivce-accout.json
Activated service account credentials for: [demo-bucket-upload@k8s-federated-conformance.iam.gserviceaccount.com]
```

##### Uploading results to Testgrid

At this point you should be able to upload your testgrid results to your GCS bucket. You can do so by running a python script availabile [here](https://github.com/kubernetes/test-infra/tree/master/testgrid/conformance). For this example, we upload results for v1.10 into it's own GCS prefix.
```
git clone https://github.com/kubernetes/test-infra
cd test-infra/testgrid/conformance
./upload_e2e.py --junit /path/to/junit_01.xml \
       --log /path/to/e2e.log \
       --bucket=gs://k8s-conformance-demo/cloud-provider-demo/e2e-conformance-release-v1.10
Uploading entry to: gs://k8s-conformance-demo/cloud-provider-demo/e2e-conformance-release-v1.10/1528333637
Run: ['gsutil', '-q', '-h', 'Content-Type:text/plain', 'cp', '-', 'gs://k8s-conformance-demo/cloud-provider-demo/e2e-conformance-release-v1.10/1528333637/started.json'] stdin={"timestamp": 1528333637}
Run: ['gsutil', '-q', '-h', 'Content-Type:text/plain', 'cp', '-', 'gs://k8s-conformance-demo/cloud-provider-demo/e2e-conformance-release-v1.10/1528333637/finished.json'] stdin={"timestamp": 1528337316, "result": "SUCCESS"}
Run: ['gsutil', '-q', '-h', 'Content-Type:text/plain', 'cp', '~/go/src/k8s.io/kubernetes/results/plugins/e2e/results/e2e.log', 'gs://k8s-conformance-demo/cloud-provider-demo/e2e-conformance-release-v1.10/1528333637/build-log.txt']
Run: ['gsutil', '-q', '-h', 'Content-Type:text/plain', 'cp', '~/go/src/k8s.io/kubernetes/results/plugins/e2e/results/e2e.log', 'gs://k8s-conformance-demo/cloud-provider-demo/e2e-conformance-release-v1.10/1528333637/artifacts/e2e.log']
Done.
```

##### Testgrid Configuration

Next thing you want to do is configure testgrid to read results from your GCS bucket. There are two [configuration](https://github.com/kubernetes/test-infra/tree/master/testgrid#configuration) steps required. One for your [test group](https://github.com/kubernetes/test-infra/tree/master/testgrid#test-groups) and one for your [dashboard](https://github.com/kubernetes/test-infra/tree/master/testgrid#dashboards).

To add a test group update [config.yaml](https://github.com/kubernetes/test-infra/blob/master/testgrid/config.yaml) with something like the following:
```
test_groups:
...
...
- name: cloud-provider-demo-e2e-conformance-release-v1.10
  gcs_prefix: k8s-conformance-demo/cloud-provider-demo/e2e-conformance-release-v1.10
```

To add a link to your results in the testgrid dashboard, update [config.yaml](https://github.com/kubernetes/test-infra/blob/master/testgrid/config.yaml) with something like the following:
```
dashboards:
...
...
- name: conformance-demo-cloud-provider
  dashboard_tab:
  - name: "Demo Cloud Provider, v1.10"
    description: "Runs conformance tests for cloud provier demo on release v1.10"
    test_group_name: cloud-provider-demo-e2e-conformance-release-v1.10
```

Once you've made the following changes, open a PR against the test-infra repo adding the sig testing label (`/sig testing`) and cc'ing @kubernetes/sig-testing-pr-reviews. Once your PR merges you should be able to view your results on https://k8s-testgrid.appspot.com/ which should be ready to be consumed by the necessary stakeholders (sig-release, sig-testing, etc).

#### Lifecycle of Test Results

You can configure the lifecycle of testgrid results by specifying fields like `days_of_results` on your test group configuration. More details about this in the [Testgrid Advanced Configuration](https://github.com/kubernetes/test-infra/tree/master/testgrid#advanced-configuration) docs. If for whatever reason you urgently need to delete testgrid results, you can contact someone from sig-testing.

#### Examples

Here are some more concrete examples of how other cloud providers are running conformance tests and uploading results to testgrid:
* Open Stack
    * [OpenLab zuul job for running/uploading testgrid results](https://github.com/theopenlab/openlab-zuul-jobs/tree/master/playbooks/cloud-provider-openstack-acceptance-test-e2e-conformance)
    * [OpenStack testgrid config](https://github.com/kubernetes/test-infra/pull/7670)
    * [OpenStack conformance tests dashboard](https://github.com/kubernetes/test-infra/pull/8154)


### Risks and Mitigations

#### Operational Overhead

Operating CI/CD system to run conformance tests on a regular basis may incur extra work from every cloud provider. Though we anticipate the benefits of running conformance tests to outweight the operational overhead, in some cases it may not.

Mitigation: TODO

#### Misconfigured Tests

There are various scenarios where cloud providers may mistakenly upload incorrect conformance tests results. One example being uploading results for the wrong Kubernetes version.

Mitigation: TODO

#### Flaky Tests

Tests can fail for various reasons in any cloud environment and may raise false negatives for the release team.

Mitigation: TODO


## Graduation Criteria

All providers are periodically uploading conformance test results in at least one of the methods outlined in this KEP.


[umbrella issues]: TODO

## Implementation History

- Jun 6th 2018: KEP is merged as a signal of acceptance. Cloud providers should now be looking to report their conformance test results to testgrid.
- Nov 19th 2018: KEP has been in implementation stage for roughly 5 months with Alibaba Cloud, Baidu Cloud, DigitalOcean, GCE, OpenStack and vSphere reporting conformance test results to testgrid.

