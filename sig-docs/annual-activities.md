# List of annual activities to be performed by SIG Docs leadership

This document lists and details the activities to be performed annually by [SIG Docs leadership](./sig-docs/README.md#leadership)

The recommended timeline for these activities is usually November - January due to lower levels of activity.

## Prune OWNERS_ALIASES under k/website

At the end of each year, leadership needs to ensure that folks listed in the [OWNERS_ALIASES file](https://github.com/kubernetes/website/blob/main/OWNERS_ALIASES) 
are willing to continue contributing.

This can be done by using the [maintainers tool](https://github.com/kubernetes-sigs/maintainers) which helps parse the file for 
activity levels of those listed from GitHub and [devstats](https://k8s.devstats.cncf.io).

### Steps

- Set up the [maintainers tool](https://github.com/kubernetes-sigs/maintainers) along with its prerequisites locally by following the instructions.
- Clone the [k/website](https://github.com/kubernetes/website) repository on your machine.
- Create a new branch locally. This is required since there are nested OWNERS_ALIASES files within the repository.
- Switch to the newly created branch and delete all nested OWNERS_ALIASES files. Retain the OWNERS_ALIASES file at the top of the repository root only.
Without this, the maintainers tool will not produce the required output.
- Run the maintainers tool against this new copy of the website repo locally.
  - Tip: While executing the command, ensure you exclude the following folks by using the `--exclude` flag:
      - Security Response Committee members,
      - Release engineering approvers and reviewers,
      - Co-chairs and tech leads of SIG Docs, SIG Release, and SIG Security
- You'll get a list of SIG Docs approvers/reviewers with no contributions and with low reviews/approvals.
- Submit a separate PR each for the list with no contributions and with low reviews/approvals and seek feedback. The timeframe for this activity should be ~2 weeks per PR.
  - Tip: The lists will likely be long. It'd be easier if you raise the second one after the first PR is merged.
  - Example PRs:
    - https://github.com/kubernetes/website/pull/38719
    - https://github.com/kubernetes/website/pull/38853
- Socialize the activity within the [SIG Docs google group](https://groups.google.com/u/1/g/kubernetes-sig-docs) and the [#sig-docs slack channel](https://kubernetes.slack.com/messages/sig-docs). Add it as an item to the upcoming meeting agenda.
- Work with all stakeholders to get the PRs merged, ideally before the end of the year.
- In case of issues with merging due to conflicts with localization initiatives, such as lack of minimum number of approvers/reviewers,
seek consensus with the localization subproject and wider SIG Docs community in the next year.
  
## Create new PR Wrangler schedule

Preferably, before the second week of December, [the PR Wrangler schedule](https://github.com/kubernetes/website/wiki/PR-Wranglers) for the upcoming year needs to be released. 

This needs to be done along with or after the pruning of the OWNERS_ALIASES file detailed above.

SIG co-chairs, tech leads, and localization subproject leads have access to edit the [k/website wiki](https://github.com/kubernetes/website/wiki/).

### Steps:

- Create a new page under [PR Wranglers](https://github.com/kubernetes/website/wiki/PR-Wranglers) for the upcoming year.
- Copy the table from previous year's page and amend the dates for the upcoming year.
- Ensure approvers/reviewers pruned per the cleanup detailed above aren't reflecting in the list.
- Socialize the new list on the [#sig-docs slack channel](https://kubernetes.slack.com/messages/sig-docs).

## Create new Issue Wrangler schedule

Preferably, before the second week of December, [the Issue Wrangler schedule](https://github.com/kubernetes/website/wiki/Issue-Wranglers) for the upcoming year needs to be released. 

This needs to be done along with or after the pruning of the OWNERS_ALIASES file detailed above.

SIG co-chairs, tech leads, and localization subproject leads have access to edit the [k/website wiki](https://github.com/kubernetes/website/wiki/).

### Steps:

- Send out a call for nominations for the Issue Wrangler role during the last quarter of the year on the the [SIG Docs google group](https://groups.google.com/u/1/g/kubernetes-sig-docs). 
- Create a new page under [Issue Wranglers](https://github.com/kubernetes/website/wiki/Issue-Wranglers) for the upcoming year.
- Copy the table from previous year's page and amend the dates for the upcoming year.
- Add nominations received to the table.
  - Tip: Aim at populating data for Q1 of the upcoming year by December.
- Socialize the new list on the [#sig-docs slack channel](https://kubernetes.slack.com/messages/sig-docs).

## Prepare annual report

SIG Docs co-chairs and tech leads are responsible for submission of annual report projecting SIG health during the first quarter of every year. 

The annual reports for previous years can be viewed [here](https://github.com/kubernetes/community/tree/master/sig-docs).

General steps and timelines are outlined in [this document](https://github.com/kubernetes/community/blob/master/committee-steering/governance/annual-reports.md).

Specific steps related to SIG Docs are outlined below.

### Steps:

- The [site analytics dashboard](https://datastudio.google.com/u/0/reporting/fede2672-b2fd-402a-91d2-7473bdb10f04/page/567IC) is featured as part of metrics we care about.
  - Specifically, we highlight number of page views and top pages for the year.
- Normally, we do not have KEPs as part of release cycles. Any work on an ongoing KEP or a subproject formalization needs to be highlighted explicitly. 
