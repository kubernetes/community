## Introduction
This document serves to understand how to set up kubemark cluster given that a base cluster (to run hollow-node pods) and separate master (to act as master for the hollow nodes) are already present.

## Precondition
You need kubemark master and external cluster to set up a kubemark cluster.

The functions are as follows:

- kubemark master: can be StandAlone or HA, used to be the kubemark cluster's master
- external cluster: used to create hollow nodes for the kubemark cluster

## Steps:
1. Build kubemark image

If you want to build/use your own kubemark image, do as follows. Otherwise skip this section and just use the latest image `staging-k8s.gcr.io/kubemark:latest` from public repo.

- i. pull kubernetes code

```
cd $GOPATH/src/k8s.io/
git clone git@github.com:kubernetes/kubernetes.git
```

- ii. build kubemark binary 

```
./hack/build-go.sh cmd/kubemark/
cp $GOPATH/src/k8s.io/kubernetes/_output/bin/kubemark $GOPATH/src/k8s.io/kubernetes/cluster/images/kubemark/
```

- iii. build kubemark image

```
cd $GOPATH/src/k8s.io/kubernetes/cluster/images/kubemark/
make build
```

Then you can get the image named `staging-k8s.gcr.io/kubemark:latest` locally.

- iv. push kubemark image

```
docker tag staging-k8s.gcr.io/kubemark:latest {{kubemark_image_registry}}/kubemark:{{kubemark_image_tag}}
docker push {{kubemark_image_registry}}/kubemark:{{kubemark_image_tag}}
```

2. Create hollow nodes

- i. create namespace, configmap and secret

Copy kubemark master's kubeconfig which is used to configure access, put it on a master of external cluster, rename it as config.

```
kubectl create ns kubemark
kubectl create configmap node-configmap -n kubemark --from-literal=content.type="test-cluster"
kubectl create secret generic kubeconfig --type=Opaque --namespace=kubemark --from-file=kubelet.kubeconfig=config --from-file=kubeproxy.kubeconfig=config
```

- ii. apply yaml to create hollow nodes

You can use `hollow-node_simplified_template.yaml` in the current directory, or use `hollow-node_template.yaml` in `kubernetes/test/kubemark/resources/`.

Note: 

- the parameters `{{numreplicas}}` means the number of hollow nodes in the kubemark cluster
- the parameters `{{numreplicas}}`, `{{kubemark_image_registry}}` and `{{kubemark_image_tag}}` need to be filled in the simplified template
- your external cluster should have enough resources to be able to run `{{numreplicas}}` no. of hollow-node pods

```
kubectl create -f hollow-node_simplified_template.yaml
```

Waiting for these hollow-node pods to be running. Then you can see these pods register as kubemark master's nodes.

Finally, kubemark master and external cluster set up the kubemark cluster.
