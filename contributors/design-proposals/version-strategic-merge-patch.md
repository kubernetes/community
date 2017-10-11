## Background

Strategic Merge Patch (SMP) is the default patch strategy that is used by`kubectl apply` and `kubectl edit`.
We often need to introduce changes to SMP to fix some issues.
But sometimes it is hard to keep fully backward compitibility.
So we want to have a way to version the SMP.

## Goal

Version the SMP without breaking the API.

## Proposed change

When we want to introduce backward incompitable changes to SMP, we will bump the SMP version.

We support json patch, json merge patch and SMP for PATCH.
Different patch types are distinguished by the Content-Type field in the request header.
We can give the new SMP version a new Content-Type. e.g. strategic-merge-patch-v2.

### Version skew

We always keep at least 2 latest versions.

- When a new kubectl sends a new version SMP to an old server, the server will reject it. kubectl will fall back to use the old version SMP.
- When an old kubectl sends an old version SMP, the server will behave as before since it knows the Content-Type in the header.
