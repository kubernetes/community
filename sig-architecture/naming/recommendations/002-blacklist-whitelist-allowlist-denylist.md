# Recommendation: replace blacklist/whitelist with allowlist/denylist

**Last Updated**: 2020-12-01

**Status:** Accepted

## Suggested Alternatives

- `allowlist/denylist`

## Context

The underlying assumption of the whitelist/blacklist metaphor is that white = good and black = bad. Because colors in and of themselves have no predetermined meaning, any meaning we assign to them is cultural: for example, the color red in many Southeast Asian countries is lucky, and is often associated with events like marriages, whereas the color white carries the same connotations in many European countries. In the case of whitelist/blacklist, the terms originate in the publishing industry, which was historically dominated by America and England, two countries that participated in slavery and which grapple with their racist legacies to this day.

From a technical communication perspective, using whitelist/blacklist as a naming convention applies metaphor (and, in turn, unintended meaning) when it isn’t needed. Descriptive words like allowlist/denylist enhance understanding. Allowlist/denylist, or simply allowed/denied as an entity prefix, are also easier to localize.

This term measures up against the evaluation framework as follows:

### First-order concerns

In short: is the term overt or identity specific? **Not quite.**

- Is the term overtly racist? **Maybe**. See [this article](https://www.ncbi.nlm.nih.gov/pmc/articles/PMC6148600/) for more.
- Is the term overtly sexist, transphobic, or pejorative about a gender identity? **No**
- Is the term overtly ableist, or pejorative to neurodiverse or disabled people? **No**
- Is the term overtly homophobic? **No**

### Second-order concerns

In short: is the term ambigously harmful, or is it harmful but not to a specific identity? **Yes**. Once again, see [this article](https://www.ncbi.nlm.nih.gov/pmc/articles/PMC6148600/)

In particular, see the following pull-quote:

> In this context, it is worth examining the origins of the term “blacklist” from the Douglas Harper Etymology Dictionary, which states that its origin and history is:
>
  >> n.
  >>
  >> also black-list, black list, “list of persons who have incurred suspicion,” 1610s, from black (adj.), here indicative of disgrace, censure, punishment (attested from 1590s, in black book) + list (n.). Specifically of employers’ list of workers considered troublesome (usually for union activity) is from 1888. As a verb, from 1718. Related: Blacklisted; blacklisting. [32]
>
> It is notable that the first recorded use of the term occurs at the time of mass enslavement and forced deportation of Africans to work in European-held colonies in the Americas.

While it is not directly attacking specific identities, in other context it does have negative connotations.

- Is the term violent? **Partially**
- Is the term militaristic? **No**

### Third-order concerns

- Is the term evocative instead of descriptive? **Yes**
- Is the term ambiguous? **Yes. As mentioned in the above description, "black" and "white" have different meanings in different cultures.**

## Precedents

- [“Blacklists” and “whitelists”: a salutary warning concerning the prevalence of racist language in discussions of predatory publishing](https://www.ncbi.nlm.nih.gov/pmc/articles/PMC6148600/)

- [IETF Network Working Group: Terminology, Power and Oppressive Language](https://tools.ietf.org/id/draft-knodel-terminology-00.html)

- [Android issue changing these terms (screenshot)](https://9to5google.com/wp-content/uploads/sites/4/2020/06/android-aosp-allowlist-explanation.png)

- [Twitter's language changes](https://www.cnet.com/news/twitter-engineers-replace-racially-loaded-tech-terms-like-master-slave/)

## Impact

The majority of uses of blacklist and whitelist occur in `/vendor` directories.

- [Hound search - blacklist](https://cs.k8s.io/?q=blacklist&i=nope&files=&repos=)
- [Hound search - whitelist](https://cs.k8s.io/?q=whitelist&i=nope&files=&repos=)

If we exclude these, areas for replacement narrow. Many repositories have relatively minor uses outside `/vendor` paths, but the following are significant:

- [kubernetes/kops](https://github.com/kubernetes/kops/): Contains breaking changes (`-whitelisted-healthcheck-cidr` flag, among other uses). Contains uses both in documentation & code.
- [kubernetes/test-infra](https://github.com/kubernetes/test-infra/): Used mainly in relation to Prow.
- [kubernetes/kubernetes](https://github.com/kubernetes/kubernetes/): Contains many non-problematic uses of blacklist, and many potentially breaking change uses of whitelist.
- [kubernetes/api](https://github.com/kubernetes/api/): Used mainly in comments/documentation.
- [kubernetes/cloud-provider-openstack](https://github.com/kubernetes/cloud-provider-openstack/blob/master/docs/keystone-auth/using-auth-data-synchronization.md): Contains breaking changes (config options). See linked documentation.
- [kubernetes/ingress-nginx](https://github.com/kubernetes/ingress-nginx/): Contains breaking changes (the `ipwhitelist` package). `kubernetes/ingress-gce` contains similar changes required.
- [kubernetes-sigs/node-feature-discovery](https://github.com/kubernetes-sigs/node-feature-discovery/blob/master/docs/advanced/worker-commandline-reference.md): Contains breaking changes (see linked file, `--label-whitelist=<pattern>` flag).
- [kubernetes/website](https://github.com/kubernetes/website/): Many uses in documentation.
- [kubernetes-sigs/kubespray](https://github.com/kubernetes-sigs/kubespray): `callback_whitelist`
- [kubernetes-sigs/reference-docs](https://github.com/kubernetes-sigs/reference-docs): Many uses in docs (`description`) fields.

In addition to the above, there are multiple repos which use either term once. `kubernetes/community` also contains uses, mainly in meeting notes.

Due to the nature of these changes, we recommend the following:

1. Pursue changes **outside** of `/vendor` directories and other imported packages. (Scope changes to our own code)
2. Open issues in all repositories for breaking changes, starting the deprecation cycle on old names.
3. Do documentation, test and non-"breaking" (external) changes next/concurrently with the above.
4. When ready, implement new names and await deprecation of old names identified in step 2.
