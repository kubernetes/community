# 2025 Annual Report: WG LTS

## Current initiatives and Project Health

1. What work did the WG do this year that should be highlighted?

   * Kubernetes regressions / backports with sig-release
     * Tracked the effect of strengthening Kubernetes backport requirements [at the end of 2024](https://github.com/kubernetes/community/issues/7634)
     * Tracked data on [Kubernetes regressions](https://docs.google.com/spreadsheets/d/1LbGKBC4D2sLkcmzY9qDx9u-1D9TKC_ZrM8iA1eHW4Hs/edit)
     * The minor versions in 2024/2025 (1.31+) are the first ones to have [zero patch release regressions](https://docs.google.com/spreadsheets/d/1LbGKBC4D2sLkcmzY9qDx9u-1D9TKC_ZrM8iA1eHW4Hs/edit?gid=751747416#gid=751747416) since tracking began in 1.19
   * Shepherded [compatibility version (KEP-4330)](https://github.com/kubernetes/enhancements/tree/master/keps/sig-architecture/4330-compatibility-versions) to stable in 1.34 with sig-api-machinery
     * Gives 3 additional releases of runtime behavior compatibility with previous minor versions
   * Explored possible proposals for publishing patches for out-of-support minor versions
     * https://docs.google.com/document/d/1MzSNg2hFEFRnkRyCbEKpt6NFJzBD7FXpfRHY8bNjA_g/edit?tab=t.0
     * https://docs.google.com/document/d/1qYoRE-gtjoe-kiBSZOdokEWmMrtwtQHvgLMLnHTwdPg/edit?tab=t.0
   * Kubecon session [Shaping LTS Together: What We’ve Learned the Hard Way](https://kccncna2025.sched.com/event/27FWW/shaping-lts-together-what-weve-learned-the-hard-way-nikhita-raghunath-nikhila-kamath-broadcom-micah-hausler-aws-jeremy-rickard-microsoft-aniket-ponkshe-canonical)
   * [Spun down the working group](https://groups.google.com/a/kubernetes.io/g/wg-lts/c/e4XyeS19BsU)

2. Are there any areas and/or subprojects that your group needs help with (e.g. fewer than 2 active OWNERS)?

   * No

## Operational

Operational tasks in [wg-governance.md]:

- [x] [README.md] reviewed for accuracy and updated if needed
- [x] WG leaders in [sigs.yaml] are accurate and active, and updated if needed
- [x] Meeting notes and recordings for 2025 are linked from [README.md] and updated/uploaded if needed
- [x] Updates provided to sponsoring SIGs in 2025: 
     - Patches proposal broadcast to sig-security, sig-architecture, sig-release
        - https://groups.google.com/a/kubernetes.io/g/wg-lts/c/v6uC6kDZA7w

[wg-governance.md]: https://git.k8s.io/community/committee-steering/governance/wg-governance.md
[README.md]: https://git.k8s.io/community/wg-lts/README.md
[sigs.yaml]: https://git.k8s.io/community/sigs.yaml
