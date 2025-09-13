# SIG Node KEP Wrangler Program

SIG Node is frequently the front-runner on completed KEP numbers in a given release. While we’ve historically been very productive, working on so many KEPs means things frequently fall between the cracks. In the retrospective for Kubernetes release v1.31.0 and subsequent conversations, the SIG discussed adding additional volunteers in the KEP wrangling process. This document describes this process and serves as a guide for new KEP wranglers to understand their responsibilities.

## Responsibilities

- Help guide contributors through the KEP process, ensuring they hit the various different deadlines (listed below)
- Help communicate blockers/high priority issues (if any) to the SIG leads and chairs to help move them forward
- Ensure someone responsible for the KEP is responding to the release and docs teams

## Important deadlines 

As a KEP wrangler you should make sure that the KEPs assigned to you are abiding the following deadlines, as set by the release team:
- PRR Freeze
- Enhancements Freeze
- Code/Test Freeze
- Docs Freeze

## Signing up to be a wrangler

If you're interested in becoming a wrangler and helping us out, wait for the wrangler signup form in the [SIG Node mailing list](https://groups.google.com/g/kubernetes-sig-node) during the initial weeks of a release cycle.

## Resources

Wranglers should be part of the following channels in the Kubernetes Slack workspace:

- [#sig-node-wranglers](https://kubernetes.slack.com/archives/C092ZDBRU64)
- [#sig-node](https://kubernetes.slack.com/archives/C0BP8PW9G)

## Process

- At the start of the release, the SIG chairs would create a KEP planning board (You can find the v1.34 one [here](https://github.com/orgs/kubernetes/projects/214/views/2) for example)
- Once the board is created and KEPs are added, feel free to assign KEPs to yourself by adding your name in the `Wranglers` column
- Once you've assigned KEPs to yourself, follow up on the status of the KEPs and ping the KEP authors as and when necessary to make sure that all the different deadlines are met

## Reporting

In order to assess the status of KEPs throughout the release, we plan to post a report of the KEPs every week in the [#sig-node-wranglers](https://kubernetes.slack.com/archives/C092ZDBRU64) channel. This can be done on Tuesdays before the weekly SIG Node meetings.

Use the following templates to post the metrics of KEPs:

#### Before enhancements freeze

```md
Metrics:

By Stage
- Alpha:
- Beta:
- Stable:
- Deprecation:

By Status
- Tracked for enhancemnets freeze: 
- At risk for enhancements freeze:
- Removed from milestone: 
```

#### Before code freeze 

```md
Metrics:

By Stage
- Alpha:
- Beta:
- Stable:
- Deprecation:

By Status
- Tracked for code freeze: 
- At risk for code freeze:
- Removed from milestone: 
```

#### Before docs freeze

```md
Metrics:

By Stage
- Alpha:
- Beta:
- Stable:
- Deprecation:

By Status
- Tracked for docs freeze: 
- At risk for docs freeze:
- Removed from milestone: 
```