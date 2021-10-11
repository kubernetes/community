# Kubernetes Policy Management Whitepaper

## About

The Kubernetes Policy Management paper is a [Policy Working Group](https://github.com/kubernetes/community/tree/master/wg-policy) project to ensure the cloud native community shares a common view and have access to information about Kubernetes policy management and related topics.

## Updates

The paper is intended to be a living document that is maintained for the community, by its members. The paper is housed in this Git repository.

### Format

The paper is maintained in markdown format. A PDF format is generated for new versions.

To create a PDF version install [pandoc](https://pandoc.org/installing.html) and run the following command:

```sh
pandoc --variable urlcolor=blue --variable geometry:margin=1in --from gfm -o kubernetes-policy-management.pdf kubernetes-policy-management.md
```

### Contributions

Updates to the whitepaper, suggestions for updates, or discussion for updates should initiate with an issue submitted to the [Policy WG repo](https://github.com/kubernetes-sigs/wg-policy-prototypes) and labeled with "suggestion" and "Kubernetes Policy Panagement Paper".





