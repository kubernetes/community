# SIG Doc builder

This script will generate the following documentation files:

```
sig-*/README.md
wg-*/README.md
sig-list.md
```

Based off the `sigs.yaml` metadata file.

## How to use

To (re)build documentation for all the SIGs, run these commands:

```bash
make all
```

To build docs for one SIG, run these commands:

```bash
make SIG=sig-apps gen-doc
make SIG=sig-testing gen-doc
make WG=resource-management gen-doc
```

where the `SIG` or `WG` var refers to the directory being built.

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
