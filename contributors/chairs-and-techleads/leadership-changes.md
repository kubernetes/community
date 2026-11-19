# Leadership Changes

This document covers the steps required to propose and execute changes to SIG, WG, and UG leadership (onboarding incoming leads and offboarding outgoing leads).

When initiating a leadership change, open a tracking issue in [kubernetes/community] using the [Leadership Change issue template](https://github.com/kubernetes/community/issues/new?template=leadership-change.yml) and follow the step-by-step procedures detailed below.

---

## 1. Prerequisites for Nominated Leads

Before nominating an individual for a leadership position (Chair, Technical Lead, or Subproject Owner):

- [ ] **Discuss the proposed changes** with the current group leadership prior to any public announcement.
- [ ] Ensure the nominee is an active [Kubernetes GitHub Org Member][member].
- [ ] Ensure the nominee has completed the Linux Foundation [Inclusive Open Source Community Orientation course][Inclusive Open Source Community Orientation course].

---

## 2. Community Announcement & Lazy Consensus

To announce an intent to step down or nominate a new lead:

- [ ] Send an email to the SIG/WG/UG mailing list and CC the [kubernetes-dev] mailing list (`dev@kubernetes.io`). At a minimum, the email should contain:
  - Intent to step down as the current lead.
  - If nominating another lead:
    - 1–2 sentences explaining why they are being nominated.
    - Links to meeting notes or discussions where the nomination was discussed.
    - Contacts to privately reach out to for questions (current leads) or concerns (current leads + [steering-private]).
  - A **lazy consensus deadline of at least one week** (7 days).

*Note: If multiple candidates are running for an open lead position and lazy consensus cannot be achieved, an election should be held. SIG Contributor Experience (`#sig-contribex` on Slack) should be contacted to assist with the administration of the election.*

---

## 3. Post-Consensus Tasks & Repository Updates

Once lazy consensus has been achieved, complete the following updates across the respective repositories:

### Update `kubernetes/community`
- [ ] Update [`sigs.yaml`]:
  - Add incoming leads under `leadership.chairs` or `leadership.technical_leads`.
  - For offboarding leads, remove them from active leadership and add them under `emeritus_leads` (preserving their name and GitHub handle).
- [ ] Run `make` (or `./generator`) from the repository root to regenerate all SIG `README.md` files and `OWNERS_ALIASES` based on the [generator documentation][generator doc]. Do not edit generated markdown files manually.

### Update `kubernetes/org` (GitHub Teams & Permissions)
- [ ] Update root `OWNERS_ALIASES` in [kubernetes/org] if the leadership role grants top-level organization aliases.
- [ ] Update [milestone-maintainers team] in `config/kubernetes/sig-release/teams.yaml` to ensure incoming leads can manage milestone assignments for PRs/issues.
- [ ] Update relevant team configuration files under `config/kubernetes/` or `config/kubernetes-sigs/` ([team configs]) (for example, `sig-<name>-leads` or `sig-<name>-maintainers`) to grant appropriate repository permissions and review notifications.

### Update `kubernetes/enhancements` (KEP Approvals)
- [ ] Update `OWNERS_ALIASES` in [kubernetes/enhancements] to ensure incoming Chairs and Technical Leads can review and approve Kubernetes Enhancement Proposals (KEPs) for their area.

### Update `kubernetes/kubernetes` (Feature Approvers)
- [ ] Update `feature-approvers` in root `OWNERS_ALIASES` in [kubernetes/kubernetes] (when applicable) to allow leads to approve enhancements during release cycles.

### Update `kubernetes/k8s.io` (Mailing Lists & Google Groups)
- [ ] Update `OWNERS_ALIASES` in [kubernetes/k8s.io] if lead privileges are required in that repository.
- [ ] In [`groups/groups.yaml`](https://github.com/kubernetes/k8s.io/blob/main/groups/groups.yaml), add the new lead to the global [leads mailing list] (`name: leads`).
- [ ] In `groups/groups.yaml` (or the respective group directory under `groups/`), update the membership of the SIG-specific leads mailing list (for example, `sig-<name>-leads@kubernetes.io`).
- [ ] Transfer Google Groups manager / owner permissions to incoming leads on the Google Groups web interface.

### Communication Channels, Access, & Projects
- [ ] **Slack**: Ensure the new lead is invited to and joins `#chairs-and-techleads` on Kubernetes Slack. Offboarding leads should gracefully transition or leave lead-only channels.
- [ ] **Shared Calendar**: Update ownership and editor permissions on the SIG's public community Google Calendar.
- [ ] **1Password**: Contact SIG Contributor Experience leads in `#sig-contribex` to grant the new lead access to the group's shared 1Password vault (used for Zoom hosting credentials and service accounts).
- [ ] **YouTube Playlists**: Contact [YouTube moderators] to update permissions for managing and editing the SIG's community meeting recordings playlist on the official Kubernetes YouTube channel.
- [ ] **GitHub Projects**: Ensure the new lead is added with write/admin access to any active GitHub Project boards (and `.projects` metadata) used by the group for tracking deliverables.

---

## 4. Related Resources

- [SIG, WG, and UG Governance](https://git.k8s.io/community/committee-steering/governance/sig-governance.md)
- [SIG/WG Lifecycle Guide][sig-wg-lifecycle.md]
- [Role of a Technical Lead](technical-lead.md)
- [Resources for Community Group Leads](README.md)
- [Kubernetes Community Membership][member]

[kubernetes-dev]: https://groups.google.com/a/kubernetes.io/g/dev
[steering-private]: mailto:steering-private@kubernetes.io
[member]: /community-membership.md#member
[`sigs.yaml`]: /sigs.yaml
[generator doc]: /generator
[kubernetes/community]: https://github.com/kubernetes/community
[kubernetes/org]: https://github.com/kubernetes/org
[kubernetes/enhancements]: https://github.com/kubernetes/enhancements
[kubernetes/kubernetes]: https://github.com/kubernetes/kubernetes
[kubernetes/k8s.io]: https://github.com/kubernetes/k8s.io
[milestone-maintainers team]: https://git.k8s.io/org/config/kubernetes/sig-release/teams.yaml
[team configs]: https://git.k8s.io/org/config
[Inclusive Open Source Community Orientation course]: https://training.linuxfoundation.org/training/inclusive-open-source-community-orientation-lfc102/
[sig-wg-lifecycle.md]: /sig-wg-lifecycle.md
[leads mailing list]: https://groups.google.com/a/kubernetes.io/g/leads
[YouTube moderators]: https://github.com/kubernetes/community/blob/master/communication/moderators.md#youtube-channel
