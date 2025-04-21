# SIG Doc builder

This folder contains scripts to automatically generate documentation about the
different Special Interest Groups (SIGs), Working Groups (WGs)
and Committees of Kubernetes. The authoritative
source for SIG information is the [`sigs.yaml`](/sigs.yaml) file in the project root.
All updates must be done there.

The schema for this file should be self explanatory. However, if you need to see all the options, check out the generator code in `app.go`.

The documentation follows a template and uses the values from [`sigs.yaml`](/sigs.yaml):

- Header: [`header.tmpl`](header.tmpl)
- List: [`list.tmpl`](list.tmpl)
- SIG README: [`sig_readme.tmpl`](sig_readme.tmpl)
- WG README: [`wg_readme.tmpl`](wg_readme.tmpl)
- Committee README: [`committee_readme.tmpl`](committee_readme.tmpl)

**Time Zone gotcha**:
Time zones make everything complicated.
And Daylight Saving time makes it even more complicated.
Meetings are specified with a time zone and we generate a link to http://www.thetimezoneconverter.com/ so people can easily convert it to their local time zone.
To make this work you need to specify the time zone in a way that the web site recognizes.
Practically, that means US pacific time must be `PT (Pacific Time)`.
`PT` isn't good enough, unfortunately.

When an update happens to the this file, the next step is to generate the
accompanying documentation. This takes the format of the following types of doc files:

```
sig-<sig-name>/README.md
wg-<working-group-name>/README.md
committee-<committee-name>/README.md
sig-list.md
```

For example, if a contributor has updated `sig-cluster-lifecycle`, the
following files will be generated:

```
sig-cluster-lifecycle/README.md
sig-list.md
```

## How to use

To (re)build documentation for all the SIGs in a go environment, run:

```bash
make generate
```
or to run this inside a container:
```bash
make generate-containerized
```

To build docs for one SIG, run one of these commands:

```bash
make WHAT=sig-apps
make WHAT=cluster-lifecycle
make WHAT=wg-resource-management
make WHAT=container-identity
```

where the `WHAT` var refers to the directory being built.


To generate the annual report template for a specific year:

```bash
make ANNUAL_REPORT=true
```

This will generate the annual report template for the previous year, as well as
drop GitHub issue templates into the `generator/generated/` directory.

You can generate the issues from these templates by running:

```bash
for i in $(ls -1 generator/generated/*.md); do gh issue create --repo kubernetes/community --title="$(head -n 1 $i)" --body-file $i && rm $i; done
```

 You may run into rate limiting issues, which is why this command removes the
 files after an issue has been successfully created.

<!--TODO: we probably won't need maintainers.txt longterm-->
To generate the maintainers.txt file for updating with the CNCF re:
https://github.com/kubernetes/steering/issues/281

```bash
make MAINTAINERS_LIST=true
```

This will generate an untracked (not saved in git) maintainers.txt file with a 
table in the format requested by the CNCF.
Most contributors will never need to do this.
For more details see the linked steering issue.
<!--END-TODO: we probably won't need maintainers.txt longterm-->

## Adding custom content

### README

If your SIG, WG or Committee wishes to add custom content, you can do so by placing it within
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

### OWNERS_ALIASES

Similarly, custom aliases can be added in the `OWNERS_ALIASES` file by placing
it within the following code comments:

```yaml
## BEGIN CUSTOM CONTENT

## END CUSTOM CONTENT
```
