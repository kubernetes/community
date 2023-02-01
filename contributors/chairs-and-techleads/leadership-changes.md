# Leadership Changes

This document covers steps needed to propose changes to SIG/WG/UG leadership.

- [ ] Discuss the proposed changes with the current leadership.

- [ ] Send an email to the SIG/WG/UG mailing list and cc the [kubernetes-dev]
mailing list. At a minimum, the email should contain:
  - [ ] intent to step down as the current lead
  - [ ] if nominating another lead
    - [ ] 1-2 lines about why they are being nominated
    - [ ] contacts to privately reach out to for questions (current leads)
          or concerns (current leads + [steering-private]) about the nomination
  - [ ] if this was discussed in a meeting, link to meeting notes
  - [ ] lazy consensus deadline of at least one week

- [ ] If nominating another lead, ensure that they
  - [ ] are a Kubernetes GitHub org [member]
  - [ ] have completed the [Inclusive Open Source Community Orientation course]

- [ ] Once lazy consensus has been achieved, update the following
      files in the respective repos:
  - [ ] [kubernetes/community]: [`sigs.yaml`] and use the [generator doc]
        to update `README.md` and `OWNERS_ALIASES` files
  - [ ] [kubernetes/org]: `OWNERS_ALIASES`, [milestone-maintainers team],
        kubernetes and kubernetes-sigs [team configs]
  - [ ] [kubernetes/enhancements]: `OWNERS_ALIASES`
  - [ ] [kubernetes/k8s.io](https://github.com/kubernetes/k8s.io): `groups/groups.yaml` - leads email group members

- [ ] Update all communication properties used by the community group.
      See [sig-wg-lifecycle.md] for details.

Note: If multiple candidates are running for an open lead position and
lazy consensus cannot be achieved, an election should be held.
SIG Contributor Experience should be contacted to assist with the
administration of the election.

[kubernetes-dev]: https://groups.google.com/a/kubernetes.io/g/dev
[steering-private]: steering-private@kubernetes.io
[member]: /community-membership.md#member
[`sigs.yaml`]: /sigs.yaml
[generator doc]: /generator
[kubernetes/community]: https://github.com/kubernetes/community
[kubernetes/org]: https://github.com/kubernetes/org
[kubernetes/enhancements]: https://github.com/kubernetes/enhancements
[milestone-maintainers team]: https://git.k8s.io/org/config/kubernetes/sig-release/teams.yaml
[team configs]: https://git.k8s.io/org/config
[Inclusive Open Source Community Orientation course]: https://training.linuxfoundation.org/training/inclusive-open-source-community-orientation-lfc102/
[sig-wg-lifecycle.md]: /sig-wg-lifecycle.md
