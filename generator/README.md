# SIG Doc builder

This folder contains scripts to automatically generate documentation about the
different Special Interest Groups (SIGs) of Kubernetes. The authoritative
source for SIG information is the `sigs.yaml` file in the project root. All
updates must be done there.

When an update happens to the this file, the next step is generate the
accompanying documentation. This takes the format of two types of doc file:

```
./sig-<sig-name>/README.md
./wg-<working-group-name>/README.md
./sig-list.md
```

For example, if a contributor has updated `sig-cluster-lifecycle`, the
following files will be generated:

```
./sig-cluster-lifecycle/README.md
./sig-list.md
```

## How to use

To (re)build documentation for all the SIGs, run these commands:

```bash
make all
```

To build docs for one SIG, run one these commands:

```bash
make WHAT=sig-apps
make WHAT=cluster-lifecycle
make WHAT=wg-resource-management
make WHAT=container-identity
```

where the `WHAT` var refers to the directory being built.

## Adding custom content to your README

If your SIG or WG wishes to add custom content, you can do so by placing it within
the following code comments:

```markdown
<!-- BEGIN CUSTOM CONTENT -->

<!-- END CUSTOM CONTENT -->
```

Anything inside these code comments are saved by the generator and appended
to newly generated content. Updating any content outside this block, however,
will be overwritten the next time the generator runs.

An example might be:

```markdown
<!-- BEGIN CUSTOM CONTENT -->
## Upcoming SIG goals
- Do this
- Do that
<!-- END CUSTOM CONTENT -->
```
