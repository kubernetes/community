# Slack Config

YAML files in this directory control the state of Slack, via
[Tempelis]. Updating a configuration file will result in Slack being updated
once the change merges. If a change is not legal, a presubmit will fail.

## Users

There is no safe, stable way to derive a specific Slack user from any
human-readable identifier. Instead of using Slack IDs everywhere, a single
mapping from GitHub usernames to Slack IDs is given in `users.yaml`. To reference
a user, they must first be added to `users.yaml`.

## Channels

Channels can be created by adding a new channel in `channels.yaml`. Channels
should be sorted alphabetically. New channels will be created in accordance
with the template specified in `template.yaml`.

Deleting channels is not permitted, but a channel can be archived by specifying
`archived: true`, or unarchived by removing it (or specifying `false`).

To rename a channel, set its `id` property to its current Slack ID, then change
the name.

A fully-specified channel looks like this:

```yaml
- name: slack-admins # mandatory
  id: C4M06S5HS      # optional except when renaming
  archived: false    # optional for unarchived channels
```

## Usergroups

Usergroups are pingable Slack groups. All members of a usergroup can be
automatically added to certain channels. A usergroup must have at least one
member. A usergroup can be removed by deleting it from the configuration.

Some usergroups (e.g. `@test-infra-oncall`) are managed by other tooling. To
prevent Tempelis from trying to deactivate these usergroups, they can be included
on the list and marked as `external: true`. Other usergroups should look like
this:

```yaml
- name: slack-admins               # mandatory, the pingable handle
  long_name: Slack Admins          # mandatory, the human-readable name
  description: Slack Admin Group   # mandatory, a description
  channels:                        # optional, a list of channels for members to auto-join
    - slack-admins
  members:                         # mandatory, a list of at least one member.
    - castrojo                     # member names must be listed in users.yaml.
    - katharine
    - jeefy
    - mrbobbytables
    - alejandrox1
    - jdumars
    - parispitmann
    - coderanger
    - idvoretskyi
    - idealhack
```

[Tempelis]: https://github.com/kubernetes-sigs/slack-infra/tree/master/tempelis
