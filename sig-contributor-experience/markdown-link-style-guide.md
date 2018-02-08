# Markdown Link Style Guide

Markdown provides for a number of different ways to link between documents. There are advantages and disadvantages of each method.

- If you are linking between documents in the same folder, the easiest way to do it is use a relative link.
```
See this other document [here](document-2.md).
```

- If you are linking to a file in the same repo, but not in the same directory, it's usually best to use an absolute link from the root of the repo. The reason for this is if the source document moves, the destination document will still have the correct link. It also allows your links to work properly in your own fork of the repo.
```
For more information, check out [this](/contributor/guide/document-3.md) document.
```

- If you are linking to a file in another kubernetes repo, then use of the git.k8s.io shortener is preferred over a direct github link. This shortener lives in the kubernetes/k8s.io repo.
```
The super cool [prow tool](https://git.k8s.io/test-infra/prow/README.md) resides in the test-infra repo under the kubernetes organization
```

- If you are moving a document and leaving a tombstone file in it's place, please also set an end date for when the tombstone can be removed.
```
This file has moved to https://git.k8s.io/community/contributors/guide/README.md.
<!--
This file is a placeholder to preserve links.  Please remove after 3 months or the release of kubernetes 1.10, whichever comes first.
-->
```
