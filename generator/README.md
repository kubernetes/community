# SIG Doc builder

This folder contains scripts to automatically generate documentation about the
different Special Interest Groups (SIGs) of Kubernetes. The authoritative
source for SIG information is the `sigs.yaml` file in the project root. All
updates must be done there.

When an update happens to the this file, the next step is generate the
accompanying documentation. This takes the format of two types of doc file:


```
./<sig-name>/README.md
./sig-list.md
```

For example, if a contributor has updated `sig-cluster-lifecycle`, the
following files will be generated:

```
./sig-cluster-lifecycle/README.md
./sig-list.md
```

## How to use

To (re)build documentation for all the SIGs, run these commands from the
project root:

```bash
make all
```

This in turn builds and runs the doc-generator Docker image.
