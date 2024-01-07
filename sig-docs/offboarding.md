# Offboarding a SIG Co-chair (Emeritus)

Soon after becoming a co-chair, it's important to prioritize finding successors, and moving to emeritus, a normal part of fostering a healthy SIG docs community.

This process is generally completed the co-chair moving to emeritus.

## Checklist

- `k/website` PR
  - Remove your github username from k/website [OWNERS_ALIASES](https://github.com/kubernetes/website/blob/main/OWNERS_ALIASES)
  - Add your github username, in alphabetical order, to emeritus_approvers in k/website [OWNERS](https://github.com/kubernetes/website/blob/main/OWNERS) and comment out to disable PR assignments
- `k/community` PR
  - Move your github/name fields from `leadership:` to `emeritus_leads:` in k/community's [sigs.yaml](https://github.com/kubernetes/community/blob/master/sigs.yaml) and regenerate the docs README.md with `make WHAT=sig-docs`
    - This also removes your github handle from k/community's `OWNER_ALIASES`
- `k/org` PR
  - Remove your github username from [OWNERS_ALIASES](https://github.com/kubernetes/org/blob/main/OWNERS_ALIASES) under `sig-docs-leads:`
  - Remove your github account, if listed, in `/config/kubernetes-sigs/sig-docs/teams.yaml`
- Other
  - Update your `k/website` PR with reference link(s) to the other PRs for completeness
  - Work with another co-chair to manually revoke `k8s-sig-docs-leads@` Google Group membership
  - Work with another co-chair to manually revoke Google Analytics access and add the new co-chairs email for refreshing the dashboards.
  - Work with another co-chair to manually revoke Netlify access

Note: You may move to emeritus and remain an approver if desired. Alternately, emeritus co-chairs are welcome to open a PR to re-establish an approver role at a later date.
