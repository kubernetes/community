# Kubernetes Icons Set

These icons are a way to standardize Kubernetes architecture diagrams for presentation. Having uniform architecture diagrams improve understandibility.

The icon for Kubernetes itself can be found at https://github.com/kubernetes/kubernetes/tree/master/logo .

## Tools

```shell
./tools/update-everything.sh
```

This script generates png from svg, and Visio Stencil from png.

### Generate png icons from svg

```shell
./tools/rasterize.sh
```

This script will convert all svg into png. It's based on a docker container defined [here](./tools/Dockerfile)

### Generate Visio Stencil from png icons

```shell
./tools/update-visio.sh
```

This script will convert png icons into Visio stencil using project from [here](https://github.com/hoveytechllc/visio-stencil-creator)

A Docker image hosted on GCR (`k8s.gcr.io/visio-stencil-creator`), is used to generate the Visio stencil file. The script [./tools/build-visio-stencil-creator.sh](./tools/build-visio-stencil-creator.sh) can be used to update the Docker image if it needs to be built.

## How to use these icons

Each icon can be found in differents formats
* [png](png)
* [svg](svg)

There are 2 types of icons: labeled and unlabeled.

| Kind  | Type       | Icon                             |
|-------|------------|----------------------------------|
|  Pod  | Labeled    | ![](./png/resources/labeled/pod-128.png)       |
|  Pod  | Unlabeled  | ![](./png/resources/unlabeled/pod-128.png)     |


## Control Plane Components icons
![](./png/control_plane_components/labeled/api-128.png)
![](./png/control_plane_components/labeled/c-c-m-128.png)
![](./png/control_plane_components/labeled/c-m-128.png)
![](./png/control_plane_components/labeled/k-proxy-128.png)
![](./png/control_plane_components/labeled/kubelet-128.png)
![](./png/control_plane_components/labeled/sched-128.png)

## Infrastructure Components icons
![](./png/infrastructure_components/labeled/control-plane-128.png)
![](./png/infrastructure_components/labeled/node-128.png)
![](./png/infrastructure_components/labeled/etcd-128.png)

## Kubernetes Resources icons
![](./png/resources/labeled/c-role-128.png)
![](./png/resources/labeled/cm-128.png)
![](./png/resources/labeled/crb-128.png)
![](./png/resources/labeled/crd-128.png)
![](./png/resources/labeled/cronjob-128.png)
![](./png/resources/labeled/deploy-128.png)
![](./png/resources/labeled/ds-128.png)
![](./png/resources/labeled/ep-128.png)
![](./png/resources/labeled/group-128.png)
![](./png/resources/labeled/hpa-128.png)
![](./png/resources/labeled/ing-128.png)
![](./png/resources/labeled/job-128.png)
![](./png/resources/labeled/limits-128.png)
![](./png/resources/labeled/netpol-128.png)
![](./png/resources/labeled/ns-128.png)
![](./png/resources/labeled/pod-128.png)
![](./png/resources/labeled/psp-128.png)
![](./png/resources/labeled/pv-128.png)
![](./png/resources/labeled/pvc-128.png)
![](./png/resources/labeled/quota-128.png)
![](./png/resources/labeled/rb-128.png)
![](./png/resources/labeled/role-128.png)
![](./png/resources/labeled/rs-128.png)
![](./png/resources/labeled/sa-128.png)
![](./png/resources/labeled/sc-128.png)
![](./png/resources/labeled/secret-128.png)
![](./png/resources/labeled/sts-128.png)
![](./png/resources/labeled/svc-128.png)
![](./png/resources/labeled/user-128.png)
![](./png/resources/labeled/vol-128.png)

## Usage Example

#### Exposed Pod with 3 replicas
![](./docs/k8s-exposed-pod.png)

## License
The Kubernetes Icons Set is licensed under a choice of either Apache-2.0
or CC-BY-4.0 (Creative Commons Attribution 4.0 International). The
Kubernetes logo is a registered trademark of The Linux Foundation, and use
of it as a trademark is subject to The Linux Foundation's Trademark Usage
Guidelines at https://www.linuxfoundation.org/trademark-usage/.
